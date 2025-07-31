// This file is auto-generated, don't edit it. Thanks.
package assistant_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddDomainWordsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddDomainWordsHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddDomainWordsHeaders) GoString() string {
	return s.String()
}

func (s *AddDomainWordsHeaders) SetCommonHeaders(v map[string]*string) *AddDomainWordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddDomainWordsHeaders) SetXAcsDingtalkAccessToken(v string) *AddDomainWordsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddDomainWordsRequest struct {
	// This parameter is required.
	AssistantId *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	// This parameter is required.
	DomainWords []*AddDomainWordsRequestDomainWords `json:"domainWords,omitempty" xml:"domainWords,omitempty" type:"Repeated"`
}

func (s AddDomainWordsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDomainWordsRequest) GoString() string {
	return s.String()
}

func (s *AddDomainWordsRequest) SetAssistantId(v string) *AddDomainWordsRequest {
	s.AssistantId = &v
	return s
}

func (s *AddDomainWordsRequest) SetDomainWords(v []*AddDomainWordsRequestDomainWords) *AddDomainWordsRequest {
	s.DomainWords = v
	return s
}

type AddDomainWordsRequestDomainWords struct {
	Description *string   `json:"description,omitempty" xml:"description,omitempty"`
	DomainWord  *string   `json:"domainWord,omitempty" xml:"domainWord,omitempty"`
	FormalWords []*string `json:"formalWords,omitempty" xml:"formalWords,omitempty" type:"Repeated"`
}

func (s AddDomainWordsRequestDomainWords) String() string {
	return tea.Prettify(s)
}

func (s AddDomainWordsRequestDomainWords) GoString() string {
	return s.String()
}

func (s *AddDomainWordsRequestDomainWords) SetDescription(v string) *AddDomainWordsRequestDomainWords {
	s.Description = &v
	return s
}

func (s *AddDomainWordsRequestDomainWords) SetDomainWord(v string) *AddDomainWordsRequestDomainWords {
	s.DomainWord = &v
	return s
}

func (s *AddDomainWordsRequestDomainWords) SetFormalWords(v []*string) *AddDomainWordsRequestDomainWords {
	s.FormalWords = v
	return s
}

type AddDomainWordsResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddDomainWordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDomainWordsResponseBody) GoString() string {
	return s.String()
}

func (s *AddDomainWordsResponseBody) SetSuccess(v bool) *AddDomainWordsResponseBody {
	s.Success = &v
	return s
}

type AddDomainWordsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDomainWordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDomainWordsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDomainWordsResponse) GoString() string {
	return s.String()
}

func (s *AddDomainWordsResponse) SetHeaders(v map[string]*string) *AddDomainWordsResponse {
	s.Headers = v
	return s
}

func (s *AddDomainWordsResponse) SetStatusCode(v int32) *AddDomainWordsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDomainWordsResponse) SetBody(v *AddDomainWordsResponseBody) *AddDomainWordsResponse {
	s.Body = v
	return s
}

type AddToOrgSkillRepositoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddToOrgSkillRepositoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddToOrgSkillRepositoryHeaders) GoString() string {
	return s.String()
}

func (s *AddToOrgSkillRepositoryHeaders) SetCommonHeaders(v map[string]*string) *AddToOrgSkillRepositoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddToOrgSkillRepositoryHeaders) SetXAcsDingtalkAccessToken(v string) *AddToOrgSkillRepositoryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddToOrgSkillRepositoryRequest struct {
	// This parameter is required.
	ActionId *string `json:"actionId,omitempty" xml:"actionId,omitempty"`
	// This parameter is required.
	ActionVersion *string `json:"actionVersion,omitempty" xml:"actionVersion,omitempty"`
	// This parameter is required.
	OperatorUnionId *string `json:"operatorUnionId,omitempty" xml:"operatorUnionId,omitempty"`
}

func (s AddToOrgSkillRepositoryRequest) String() string {
	return tea.Prettify(s)
}

func (s AddToOrgSkillRepositoryRequest) GoString() string {
	return s.String()
}

func (s *AddToOrgSkillRepositoryRequest) SetActionId(v string) *AddToOrgSkillRepositoryRequest {
	s.ActionId = &v
	return s
}

func (s *AddToOrgSkillRepositoryRequest) SetActionVersion(v string) *AddToOrgSkillRepositoryRequest {
	s.ActionVersion = &v
	return s
}

func (s *AddToOrgSkillRepositoryRequest) SetOperatorUnionId(v string) *AddToOrgSkillRepositoryRequest {
	s.OperatorUnionId = &v
	return s
}

type AddToOrgSkillRepositoryResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddToOrgSkillRepositoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddToOrgSkillRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *AddToOrgSkillRepositoryResponseBody) SetSuccess(v bool) *AddToOrgSkillRepositoryResponseBody {
	s.Success = &v
	return s
}

type AddToOrgSkillRepositoryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddToOrgSkillRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddToOrgSkillRepositoryResponse) String() string {
	return tea.Prettify(s)
}

func (s AddToOrgSkillRepositoryResponse) GoString() string {
	return s.String()
}

func (s *AddToOrgSkillRepositoryResponse) SetHeaders(v map[string]*string) *AddToOrgSkillRepositoryResponse {
	s.Headers = v
	return s
}

func (s *AddToOrgSkillRepositoryResponse) SetStatusCode(v int32) *AddToOrgSkillRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *AddToOrgSkillRepositoryResponse) SetBody(v *AddToOrgSkillRepositoryResponseBody) *AddToOrgSkillRepositoryResponse {
	s.Body = v
	return s
}

type AssistantMeResponseHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AssistantMeResponseHeaders) String() string {
	return tea.Prettify(s)
}

func (s AssistantMeResponseHeaders) GoString() string {
	return s.String()
}

func (s *AssistantMeResponseHeaders) SetCommonHeaders(v map[string]*string) *AssistantMeResponseHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AssistantMeResponseHeaders) SetXAcsDingtalkAccessToken(v string) *AssistantMeResponseHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AssistantMeResponseRequest struct {
	// This parameter is required.
	Input        *string                `json:"input,omitempty" xml:"input,omitempty"`
	Instructions *string                `json:"instructions,omitempty" xml:"instructions,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	Stream       *bool                  `json:"stream,omitempty" xml:"stream,omitempty"`
}

func (s AssistantMeResponseRequest) String() string {
	return tea.Prettify(s)
}

func (s AssistantMeResponseRequest) GoString() string {
	return s.String()
}

func (s *AssistantMeResponseRequest) SetInput(v string) *AssistantMeResponseRequest {
	s.Input = &v
	return s
}

func (s *AssistantMeResponseRequest) SetInstructions(v string) *AssistantMeResponseRequest {
	s.Instructions = &v
	return s
}

func (s *AssistantMeResponseRequest) SetMetadata(v map[string]interface{}) *AssistantMeResponseRequest {
	s.Metadata = v
	return s
}

func (s *AssistantMeResponseRequest) SetStream(v bool) *AssistantMeResponseRequest {
	s.Stream = &v
	return s
}

type AssistantMeResponseResponseBody struct {
	CreatedAt *int64                                   `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Error     *string                                  `json:"error,omitempty" xml:"error,omitempty"`
	Id        *string                                  `json:"id,omitempty" xml:"id,omitempty"`
	Metadata  map[string]interface{}                   `json:"metadata,omitempty" xml:"metadata,omitempty"`
	Model     *string                                  `json:"model,omitempty" xml:"model,omitempty"`
	Object    *string                                  `json:"object,omitempty" xml:"object,omitempty"`
	Output    []*AssistantMeResponseResponseBodyOutput `json:"output,omitempty" xml:"output,omitempty" type:"Repeated"`
	Status    *string                                  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s AssistantMeResponseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssistantMeResponseResponseBody) GoString() string {
	return s.String()
}

func (s *AssistantMeResponseResponseBody) SetCreatedAt(v int64) *AssistantMeResponseResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *AssistantMeResponseResponseBody) SetError(v string) *AssistantMeResponseResponseBody {
	s.Error = &v
	return s
}

func (s *AssistantMeResponseResponseBody) SetId(v string) *AssistantMeResponseResponseBody {
	s.Id = &v
	return s
}

func (s *AssistantMeResponseResponseBody) SetMetadata(v map[string]interface{}) *AssistantMeResponseResponseBody {
	s.Metadata = v
	return s
}

func (s *AssistantMeResponseResponseBody) SetModel(v string) *AssistantMeResponseResponseBody {
	s.Model = &v
	return s
}

func (s *AssistantMeResponseResponseBody) SetObject(v string) *AssistantMeResponseResponseBody {
	s.Object = &v
	return s
}

func (s *AssistantMeResponseResponseBody) SetOutput(v []*AssistantMeResponseResponseBodyOutput) *AssistantMeResponseResponseBody {
	s.Output = v
	return s
}

func (s *AssistantMeResponseResponseBody) SetStatus(v string) *AssistantMeResponseResponseBody {
	s.Status = &v
	return s
}

type AssistantMeResponseResponseBodyOutput struct {
	Content []*AssistantMeResponseResponseBodyOutputContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	Id      *string                                         `json:"id,omitempty" xml:"id,omitempty"`
	Role    *string                                         `json:"role,omitempty" xml:"role,omitempty"`
	Type    *string                                         `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AssistantMeResponseResponseBodyOutput) String() string {
	return tea.Prettify(s)
}

func (s AssistantMeResponseResponseBodyOutput) GoString() string {
	return s.String()
}

func (s *AssistantMeResponseResponseBodyOutput) SetContent(v []*AssistantMeResponseResponseBodyOutputContent) *AssistantMeResponseResponseBodyOutput {
	s.Content = v
	return s
}

func (s *AssistantMeResponseResponseBodyOutput) SetId(v string) *AssistantMeResponseResponseBodyOutput {
	s.Id = &v
	return s
}

func (s *AssistantMeResponseResponseBodyOutput) SetRole(v string) *AssistantMeResponseResponseBodyOutput {
	s.Role = &v
	return s
}

func (s *AssistantMeResponseResponseBodyOutput) SetType(v string) *AssistantMeResponseResponseBodyOutput {
	s.Type = &v
	return s
}

type AssistantMeResponseResponseBodyOutputContent struct {
	Annotations []interface{} `json:"annotations,omitempty" xml:"annotations,omitempty" type:"Repeated"`
	Text        *string       `json:"text,omitempty" xml:"text,omitempty"`
	Type        *string       `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AssistantMeResponseResponseBodyOutputContent) String() string {
	return tea.Prettify(s)
}

func (s AssistantMeResponseResponseBodyOutputContent) GoString() string {
	return s.String()
}

func (s *AssistantMeResponseResponseBodyOutputContent) SetAnnotations(v []interface{}) *AssistantMeResponseResponseBodyOutputContent {
	s.Annotations = v
	return s
}

func (s *AssistantMeResponseResponseBodyOutputContent) SetText(v string) *AssistantMeResponseResponseBodyOutputContent {
	s.Text = &v
	return s
}

func (s *AssistantMeResponseResponseBodyOutputContent) SetType(v string) *AssistantMeResponseResponseBodyOutputContent {
	s.Type = &v
	return s
}

type AssistantMeResponseResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssistantMeResponseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssistantMeResponseResponse) String() string {
	return tea.Prettify(s)
}

func (s AssistantMeResponseResponse) GoString() string {
	return s.String()
}

func (s *AssistantMeResponseResponse) SetHeaders(v map[string]*string) *AssistantMeResponseResponse {
	s.Headers = v
	return s
}

func (s *AssistantMeResponseResponse) SetStatusCode(v int32) *AssistantMeResponseResponse {
	s.StatusCode = &v
	return s
}

func (s *AssistantMeResponseResponse) SetBody(v *AssistantMeResponseResponseBody) *AssistantMeResponseResponse {
	s.Body = v
	return s
}

type AssistantResponseHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AssistantResponseHeaders) String() string {
	return tea.Prettify(s)
}

func (s AssistantResponseHeaders) GoString() string {
	return s.String()
}

func (s *AssistantResponseHeaders) SetCommonHeaders(v map[string]*string) *AssistantResponseHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AssistantResponseHeaders) SetXAcsDingtalkAccessToken(v string) *AssistantResponseHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AssistantResponseRequest struct {
	// This parameter is required.
	Input        *string                `json:"input,omitempty" xml:"input,omitempty"`
	Instructions *string                `json:"instructions,omitempty" xml:"instructions,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	Stream       *bool                  `json:"stream,omitempty" xml:"stream,omitempty"`
}

func (s AssistantResponseRequest) String() string {
	return tea.Prettify(s)
}

func (s AssistantResponseRequest) GoString() string {
	return s.String()
}

func (s *AssistantResponseRequest) SetInput(v string) *AssistantResponseRequest {
	s.Input = &v
	return s
}

func (s *AssistantResponseRequest) SetInstructions(v string) *AssistantResponseRequest {
	s.Instructions = &v
	return s
}

func (s *AssistantResponseRequest) SetMetadata(v map[string]interface{}) *AssistantResponseRequest {
	s.Metadata = v
	return s
}

func (s *AssistantResponseRequest) SetStream(v bool) *AssistantResponseRequest {
	s.Stream = &v
	return s
}

type AssistantResponseResponseBody struct {
	CreatedAt *int64                                 `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Error     *string                                `json:"error,omitempty" xml:"error,omitempty"`
	Id        *string                                `json:"id,omitempty" xml:"id,omitempty"`
	Metadata  map[string]interface{}                 `json:"metadata,omitempty" xml:"metadata,omitempty"`
	Model     *string                                `json:"model,omitempty" xml:"model,omitempty"`
	Object    *string                                `json:"object,omitempty" xml:"object,omitempty"`
	Output    []*AssistantResponseResponseBodyOutput `json:"output,omitempty" xml:"output,omitempty" type:"Repeated"`
	Status    *string                                `json:"status,omitempty" xml:"status,omitempty"`
}

func (s AssistantResponseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssistantResponseResponseBody) GoString() string {
	return s.String()
}

func (s *AssistantResponseResponseBody) SetCreatedAt(v int64) *AssistantResponseResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *AssistantResponseResponseBody) SetError(v string) *AssistantResponseResponseBody {
	s.Error = &v
	return s
}

func (s *AssistantResponseResponseBody) SetId(v string) *AssistantResponseResponseBody {
	s.Id = &v
	return s
}

func (s *AssistantResponseResponseBody) SetMetadata(v map[string]interface{}) *AssistantResponseResponseBody {
	s.Metadata = v
	return s
}

func (s *AssistantResponseResponseBody) SetModel(v string) *AssistantResponseResponseBody {
	s.Model = &v
	return s
}

func (s *AssistantResponseResponseBody) SetObject(v string) *AssistantResponseResponseBody {
	s.Object = &v
	return s
}

func (s *AssistantResponseResponseBody) SetOutput(v []*AssistantResponseResponseBodyOutput) *AssistantResponseResponseBody {
	s.Output = v
	return s
}

func (s *AssistantResponseResponseBody) SetStatus(v string) *AssistantResponseResponseBody {
	s.Status = &v
	return s
}

type AssistantResponseResponseBodyOutput struct {
	Content []*AssistantResponseResponseBodyOutputContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	Id      *string                                       `json:"id,omitempty" xml:"id,omitempty"`
	Role    *string                                       `json:"role,omitempty" xml:"role,omitempty"`
	Type    *string                                       `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AssistantResponseResponseBodyOutput) String() string {
	return tea.Prettify(s)
}

func (s AssistantResponseResponseBodyOutput) GoString() string {
	return s.String()
}

func (s *AssistantResponseResponseBodyOutput) SetContent(v []*AssistantResponseResponseBodyOutputContent) *AssistantResponseResponseBodyOutput {
	s.Content = v
	return s
}

func (s *AssistantResponseResponseBodyOutput) SetId(v string) *AssistantResponseResponseBodyOutput {
	s.Id = &v
	return s
}

func (s *AssistantResponseResponseBodyOutput) SetRole(v string) *AssistantResponseResponseBodyOutput {
	s.Role = &v
	return s
}

func (s *AssistantResponseResponseBodyOutput) SetType(v string) *AssistantResponseResponseBodyOutput {
	s.Type = &v
	return s
}

type AssistantResponseResponseBodyOutputContent struct {
	Annotations []interface{} `json:"annotations,omitempty" xml:"annotations,omitempty" type:"Repeated"`
	Text        *string       `json:"text,omitempty" xml:"text,omitempty"`
	Type        *string       `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AssistantResponseResponseBodyOutputContent) String() string {
	return tea.Prettify(s)
}

func (s AssistantResponseResponseBodyOutputContent) GoString() string {
	return s.String()
}

func (s *AssistantResponseResponseBodyOutputContent) SetAnnotations(v []interface{}) *AssistantResponseResponseBodyOutputContent {
	s.Annotations = v
	return s
}

func (s *AssistantResponseResponseBodyOutputContent) SetText(v string) *AssistantResponseResponseBodyOutputContent {
	s.Text = &v
	return s
}

func (s *AssistantResponseResponseBodyOutputContent) SetType(v string) *AssistantResponseResponseBodyOutputContent {
	s.Type = &v
	return s
}

type AssistantResponseResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssistantResponseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssistantResponseResponse) String() string {
	return tea.Prettify(s)
}

func (s AssistantResponseResponse) GoString() string {
	return s.String()
}

func (s *AssistantResponseResponse) SetHeaders(v map[string]*string) *AssistantResponseResponse {
	s.Headers = v
	return s
}

func (s *AssistantResponseResponse) SetStatusCode(v int32) *AssistantResponseResponse {
	s.StatusCode = &v
	return s
}

func (s *AssistantResponseResponse) SetBody(v *AssistantResponseResponseBody) *AssistantResponseResponse {
	s.Body = v
	return s
}

type BatchGetAICreditsRecordHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchGetAICreditsRecordHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchGetAICreditsRecordHeaders) GoString() string {
	return s.String()
}

func (s *BatchGetAICreditsRecordHeaders) SetCommonHeaders(v map[string]*string) *BatchGetAICreditsRecordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchGetAICreditsRecordHeaders) SetXAcsDingtalkAccessToken(v string) *BatchGetAICreditsRecordHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchGetAICreditsRecordRequest struct {
	AssistantId *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	EndTime     *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	PageNumber  *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize    *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StartTime   *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	UnionId     *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s BatchGetAICreditsRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetAICreditsRecordRequest) GoString() string {
	return s.String()
}

func (s *BatchGetAICreditsRecordRequest) SetAssistantId(v string) *BatchGetAICreditsRecordRequest {
	s.AssistantId = &v
	return s
}

func (s *BatchGetAICreditsRecordRequest) SetEndTime(v string) *BatchGetAICreditsRecordRequest {
	s.EndTime = &v
	return s
}

func (s *BatchGetAICreditsRecordRequest) SetPageNumber(v int32) *BatchGetAICreditsRecordRequest {
	s.PageNumber = &v
	return s
}

func (s *BatchGetAICreditsRecordRequest) SetPageSize(v int32) *BatchGetAICreditsRecordRequest {
	s.PageSize = &v
	return s
}

func (s *BatchGetAICreditsRecordRequest) SetStartTime(v string) *BatchGetAICreditsRecordRequest {
	s.StartTime = &v
	return s
}

func (s *BatchGetAICreditsRecordRequest) SetUnionId(v string) *BatchGetAICreditsRecordRequest {
	s.UnionId = &v
	return s
}

type BatchGetAICreditsRecordResponseBody struct {
	HasMore    *bool                                      `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*BatchGetAICreditsRecordResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	TotalCount *int32                                     `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s BatchGetAICreditsRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchGetAICreditsRecordResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetAICreditsRecordResponseBody) SetHasMore(v bool) *BatchGetAICreditsRecordResponseBody {
	s.HasMore = &v
	return s
}

func (s *BatchGetAICreditsRecordResponseBody) SetList(v []*BatchGetAICreditsRecordResponseBodyList) *BatchGetAICreditsRecordResponseBody {
	s.List = v
	return s
}

func (s *BatchGetAICreditsRecordResponseBody) SetTotalCount(v int32) *BatchGetAICreditsRecordResponseBody {
	s.TotalCount = &v
	return s
}

type BatchGetAICreditsRecordResponseBodyList struct {
	ActionNames   *string  `json:"actionNames,omitempty" xml:"actionNames,omitempty"`
	AssistantId   *string  `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	AssistantName *string  `json:"assistantName,omitempty" xml:"assistantName,omitempty"`
	DeptId        *int64   `json:"deptId,omitempty" xml:"deptId,omitempty"`
	DeptName      *string  `json:"deptName,omitempty" xml:"deptName,omitempty"`
	RequestId     *string  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Time          *string  `json:"time,omitempty" xml:"time,omitempty"`
	UnionId       *string  `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UsedNum       *float64 `json:"usedNum,omitempty" xml:"usedNum,omitempty"`
	UserName      *string  `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s BatchGetAICreditsRecordResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s BatchGetAICreditsRecordResponseBodyList) GoString() string {
	return s.String()
}

func (s *BatchGetAICreditsRecordResponseBodyList) SetActionNames(v string) *BatchGetAICreditsRecordResponseBodyList {
	s.ActionNames = &v
	return s
}

func (s *BatchGetAICreditsRecordResponseBodyList) SetAssistantId(v string) *BatchGetAICreditsRecordResponseBodyList {
	s.AssistantId = &v
	return s
}

func (s *BatchGetAICreditsRecordResponseBodyList) SetAssistantName(v string) *BatchGetAICreditsRecordResponseBodyList {
	s.AssistantName = &v
	return s
}

func (s *BatchGetAICreditsRecordResponseBodyList) SetDeptId(v int64) *BatchGetAICreditsRecordResponseBodyList {
	s.DeptId = &v
	return s
}

func (s *BatchGetAICreditsRecordResponseBodyList) SetDeptName(v string) *BatchGetAICreditsRecordResponseBodyList {
	s.DeptName = &v
	return s
}

func (s *BatchGetAICreditsRecordResponseBodyList) SetRequestId(v string) *BatchGetAICreditsRecordResponseBodyList {
	s.RequestId = &v
	return s
}

func (s *BatchGetAICreditsRecordResponseBodyList) SetTime(v string) *BatchGetAICreditsRecordResponseBodyList {
	s.Time = &v
	return s
}

func (s *BatchGetAICreditsRecordResponseBodyList) SetUnionId(v string) *BatchGetAICreditsRecordResponseBodyList {
	s.UnionId = &v
	return s
}

func (s *BatchGetAICreditsRecordResponseBodyList) SetUsedNum(v float64) *BatchGetAICreditsRecordResponseBodyList {
	s.UsedNum = &v
	return s
}

func (s *BatchGetAICreditsRecordResponseBodyList) SetUserName(v string) *BatchGetAICreditsRecordResponseBodyList {
	s.UserName = &v
	return s
}

type BatchGetAICreditsRecordResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetAICreditsRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetAICreditsRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchGetAICreditsRecordResponse) GoString() string {
	return s.String()
}

func (s *BatchGetAICreditsRecordResponse) SetHeaders(v map[string]*string) *BatchGetAICreditsRecordResponse {
	s.Headers = v
	return s
}

func (s *BatchGetAICreditsRecordResponse) SetStatusCode(v int32) *BatchGetAICreditsRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetAICreditsRecordResponse) SetBody(v *BatchGetAICreditsRecordResponseBody) *BatchGetAICreditsRecordResponse {
	s.Body = v
	return s
}

type CreateAssistantHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateAssistantHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateAssistantHeaders) GoString() string {
	return s.String()
}

func (s *CreateAssistantHeaders) SetCommonHeaders(v map[string]*string) *CreateAssistantHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateAssistantHeaders) SetXAcsDingtalkAccessToken(v string) *CreateAssistantHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateAssistantRequest struct {
	// This parameter is required.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// This parameter is required.
	Instructions *string `json:"instructions,omitempty" xml:"instructions,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	OperatorUnionId  *string   `json:"operatorUnionId,omitempty" xml:"operatorUnionId,omitempty"`
	RecommendPrompts []*string `json:"recommendPrompts,omitempty" xml:"recommendPrompts,omitempty" type:"Repeated"`
	// This parameter is required.
	WelcomeContent *string `json:"welcomeContent,omitempty" xml:"welcomeContent,omitempty"`
}

func (s CreateAssistantRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAssistantRequest) GoString() string {
	return s.String()
}

func (s *CreateAssistantRequest) SetDescription(v string) *CreateAssistantRequest {
	s.Description = &v
	return s
}

func (s *CreateAssistantRequest) SetIcon(v string) *CreateAssistantRequest {
	s.Icon = &v
	return s
}

func (s *CreateAssistantRequest) SetInstructions(v string) *CreateAssistantRequest {
	s.Instructions = &v
	return s
}

func (s *CreateAssistantRequest) SetName(v string) *CreateAssistantRequest {
	s.Name = &v
	return s
}

func (s *CreateAssistantRequest) SetOperatorUnionId(v string) *CreateAssistantRequest {
	s.OperatorUnionId = &v
	return s
}

func (s *CreateAssistantRequest) SetRecommendPrompts(v []*string) *CreateAssistantRequest {
	s.RecommendPrompts = v
	return s
}

func (s *CreateAssistantRequest) SetWelcomeContent(v string) *CreateAssistantRequest {
	s.WelcomeContent = &v
	return s
}

type CreateAssistantResponseBody struct {
	ActionNames        []*string `json:"actionNames,omitempty" xml:"actionNames,omitempty" type:"Repeated"`
	AssistantId        *string   `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	AssistantUnionId   *string   `json:"assistantUnionId,omitempty" xml:"assistantUnionId,omitempty"`
	CreatedAt          *int64    `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatorUnionId     *string   `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	Description        *string   `json:"description,omitempty" xml:"description,omitempty"`
	FallbackContent    *string   `json:"fallbackContent,omitempty" xml:"fallbackContent,omitempty"`
	Icon               *string   `json:"icon,omitempty" xml:"icon,omitempty"`
	Instructions       *string   `json:"instructions,omitempty" xml:"instructions,omitempty"`
	KnowledgeFileNames []*string `json:"knowledgeFileNames,omitempty" xml:"knowledgeFileNames,omitempty" type:"Repeated"`
	Model              *string   `json:"model,omitempty" xml:"model,omitempty"`
	Name               *string   `json:"name,omitempty" xml:"name,omitempty"`
	RecommendPrompts   []*string `json:"recommendPrompts,omitempty" xml:"recommendPrompts,omitempty" type:"Repeated"`
	UnifiedAppId       *string   `json:"unifiedAppId,omitempty" xml:"unifiedAppId,omitempty"`
	WelcomeContent     *string   `json:"welcomeContent,omitempty" xml:"welcomeContent,omitempty"`
}

func (s CreateAssistantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAssistantResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAssistantResponseBody) SetActionNames(v []*string) *CreateAssistantResponseBody {
	s.ActionNames = v
	return s
}

func (s *CreateAssistantResponseBody) SetAssistantId(v string) *CreateAssistantResponseBody {
	s.AssistantId = &v
	return s
}

func (s *CreateAssistantResponseBody) SetAssistantUnionId(v string) *CreateAssistantResponseBody {
	s.AssistantUnionId = &v
	return s
}

func (s *CreateAssistantResponseBody) SetCreatedAt(v int64) *CreateAssistantResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateAssistantResponseBody) SetCreatorUnionId(v string) *CreateAssistantResponseBody {
	s.CreatorUnionId = &v
	return s
}

func (s *CreateAssistantResponseBody) SetDescription(v string) *CreateAssistantResponseBody {
	s.Description = &v
	return s
}

func (s *CreateAssistantResponseBody) SetFallbackContent(v string) *CreateAssistantResponseBody {
	s.FallbackContent = &v
	return s
}

func (s *CreateAssistantResponseBody) SetIcon(v string) *CreateAssistantResponseBody {
	s.Icon = &v
	return s
}

func (s *CreateAssistantResponseBody) SetInstructions(v string) *CreateAssistantResponseBody {
	s.Instructions = &v
	return s
}

func (s *CreateAssistantResponseBody) SetKnowledgeFileNames(v []*string) *CreateAssistantResponseBody {
	s.KnowledgeFileNames = v
	return s
}

func (s *CreateAssistantResponseBody) SetModel(v string) *CreateAssistantResponseBody {
	s.Model = &v
	return s
}

func (s *CreateAssistantResponseBody) SetName(v string) *CreateAssistantResponseBody {
	s.Name = &v
	return s
}

func (s *CreateAssistantResponseBody) SetRecommendPrompts(v []*string) *CreateAssistantResponseBody {
	s.RecommendPrompts = v
	return s
}

func (s *CreateAssistantResponseBody) SetUnifiedAppId(v string) *CreateAssistantResponseBody {
	s.UnifiedAppId = &v
	return s
}

func (s *CreateAssistantResponseBody) SetWelcomeContent(v string) *CreateAssistantResponseBody {
	s.WelcomeContent = &v
	return s
}

type CreateAssistantResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAssistantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAssistantResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAssistantResponse) GoString() string {
	return s.String()
}

func (s *CreateAssistantResponse) SetHeaders(v map[string]*string) *CreateAssistantResponse {
	s.Headers = v
	return s
}

func (s *CreateAssistantResponse) SetStatusCode(v int32) *CreateAssistantResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAssistantResponse) SetBody(v *CreateAssistantResponseBody) *CreateAssistantResponse {
	s.Body = v
	return s
}

type CreateAssistantMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateAssistantMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateAssistantMessageHeaders) GoString() string {
	return s.String()
}

func (s *CreateAssistantMessageHeaders) SetCommonHeaders(v map[string]*string) *CreateAssistantMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateAssistantMessageHeaders) SetXAcsDingtalkAccessToken(v string) *CreateAssistantMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateAssistantMessageRequest struct {
	// This parameter is required.
	Content   *string                `json:"content,omitempty" xml:"content,omitempty"`
	Extension map[string]*string     `json:"extension,omitempty" xml:"extension,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// This parameter is required.
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s CreateAssistantMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAssistantMessageRequest) GoString() string {
	return s.String()
}

func (s *CreateAssistantMessageRequest) SetContent(v string) *CreateAssistantMessageRequest {
	s.Content = &v
	return s
}

func (s *CreateAssistantMessageRequest) SetExtension(v map[string]*string) *CreateAssistantMessageRequest {
	s.Extension = v
	return s
}

func (s *CreateAssistantMessageRequest) SetMetadata(v map[string]interface{}) *CreateAssistantMessageRequest {
	s.Metadata = v
	return s
}

func (s *CreateAssistantMessageRequest) SetRole(v string) *CreateAssistantMessageRequest {
	s.Role = &v
	return s
}

type CreateAssistantMessageResponseBody struct {
	AssistantId *string                `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	Content     []interface{}          `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	CreatedAt   *int64                 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Id          *string                `json:"id,omitempty" xml:"id,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	Object      *string                `json:"object,omitempty" xml:"object,omitempty"`
	Role        *string                `json:"role,omitempty" xml:"role,omitempty"`
	RunId       *string                `json:"runId,omitempty" xml:"runId,omitempty"`
	ThreadId    *string                `json:"threadId,omitempty" xml:"threadId,omitempty"`
}

func (s CreateAssistantMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAssistantMessageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAssistantMessageResponseBody) SetAssistantId(v string) *CreateAssistantMessageResponseBody {
	s.AssistantId = &v
	return s
}

func (s *CreateAssistantMessageResponseBody) SetContent(v []interface{}) *CreateAssistantMessageResponseBody {
	s.Content = v
	return s
}

func (s *CreateAssistantMessageResponseBody) SetCreatedAt(v int64) *CreateAssistantMessageResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateAssistantMessageResponseBody) SetId(v string) *CreateAssistantMessageResponseBody {
	s.Id = &v
	return s
}

func (s *CreateAssistantMessageResponseBody) SetMetadata(v map[string]interface{}) *CreateAssistantMessageResponseBody {
	s.Metadata = v
	return s
}

func (s *CreateAssistantMessageResponseBody) SetObject(v string) *CreateAssistantMessageResponseBody {
	s.Object = &v
	return s
}

func (s *CreateAssistantMessageResponseBody) SetRole(v string) *CreateAssistantMessageResponseBody {
	s.Role = &v
	return s
}

func (s *CreateAssistantMessageResponseBody) SetRunId(v string) *CreateAssistantMessageResponseBody {
	s.RunId = &v
	return s
}

func (s *CreateAssistantMessageResponseBody) SetThreadId(v string) *CreateAssistantMessageResponseBody {
	s.ThreadId = &v
	return s
}

type CreateAssistantMessageResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAssistantMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAssistantMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAssistantMessageResponse) GoString() string {
	return s.String()
}

func (s *CreateAssistantMessageResponse) SetHeaders(v map[string]*string) *CreateAssistantMessageResponse {
	s.Headers = v
	return s
}

func (s *CreateAssistantMessageResponse) SetStatusCode(v int32) *CreateAssistantMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAssistantMessageResponse) SetBody(v *CreateAssistantMessageResponseBody) *CreateAssistantMessageResponse {
	s.Body = v
	return s
}

type CreateAssistantRunHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateAssistantRunHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateAssistantRunHeaders) GoString() string {
	return s.String()
}

func (s *CreateAssistantRunHeaders) SetCommonHeaders(v map[string]*string) *CreateAssistantRunHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateAssistantRunHeaders) SetXAcsDingtalkAccessToken(v string) *CreateAssistantRunHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateAssistantRunRequest struct {
	// This parameter is required.
	AssistantId  *string                `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	Instructions *string                `json:"instructions,omitempty" xml:"instructions,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	Stream       *bool                  `json:"stream,omitempty" xml:"stream,omitempty"`
}

func (s CreateAssistantRunRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAssistantRunRequest) GoString() string {
	return s.String()
}

func (s *CreateAssistantRunRequest) SetAssistantId(v string) *CreateAssistantRunRequest {
	s.AssistantId = &v
	return s
}

func (s *CreateAssistantRunRequest) SetInstructions(v string) *CreateAssistantRunRequest {
	s.Instructions = &v
	return s
}

func (s *CreateAssistantRunRequest) SetMetadata(v map[string]interface{}) *CreateAssistantRunRequest {
	s.Metadata = v
	return s
}

func (s *CreateAssistantRunRequest) SetStream(v bool) *CreateAssistantRunRequest {
	s.Stream = &v
	return s
}

type CreateAssistantRunResponseBody struct {
	AssistantId  *string                `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	CancelledAt  *int64                 `json:"cancelledAt,omitempty" xml:"cancelledAt,omitempty"`
	CompletedAt  *int64                 `json:"completedAt,omitempty" xml:"completedAt,omitempty"`
	CreatedAt    *int64                 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	ExpiresAt    *int64                 `json:"expiresAt,omitempty" xml:"expiresAt,omitempty"`
	FailedAt     *int64                 `json:"failedAt,omitempty" xml:"failedAt,omitempty"`
	Id           *string                `json:"id,omitempty" xml:"id,omitempty"`
	LastErrorMsg *string                `json:"lastErrorMsg,omitempty" xml:"lastErrorMsg,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	Object       *string                `json:"object,omitempty" xml:"object,omitempty"`
	StartedAt    *int64                 `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	Status       *string                `json:"status,omitempty" xml:"status,omitempty"`
	ThreadId     *string                `json:"threadId,omitempty" xml:"threadId,omitempty"`
}

func (s CreateAssistantRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAssistantRunResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAssistantRunResponseBody) SetAssistantId(v string) *CreateAssistantRunResponseBody {
	s.AssistantId = &v
	return s
}

func (s *CreateAssistantRunResponseBody) SetCancelledAt(v int64) *CreateAssistantRunResponseBody {
	s.CancelledAt = &v
	return s
}

func (s *CreateAssistantRunResponseBody) SetCompletedAt(v int64) *CreateAssistantRunResponseBody {
	s.CompletedAt = &v
	return s
}

func (s *CreateAssistantRunResponseBody) SetCreatedAt(v int64) *CreateAssistantRunResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateAssistantRunResponseBody) SetExpiresAt(v int64) *CreateAssistantRunResponseBody {
	s.ExpiresAt = &v
	return s
}

func (s *CreateAssistantRunResponseBody) SetFailedAt(v int64) *CreateAssistantRunResponseBody {
	s.FailedAt = &v
	return s
}

func (s *CreateAssistantRunResponseBody) SetId(v string) *CreateAssistantRunResponseBody {
	s.Id = &v
	return s
}

func (s *CreateAssistantRunResponseBody) SetLastErrorMsg(v string) *CreateAssistantRunResponseBody {
	s.LastErrorMsg = &v
	return s
}

func (s *CreateAssistantRunResponseBody) SetMetadata(v map[string]interface{}) *CreateAssistantRunResponseBody {
	s.Metadata = v
	return s
}

func (s *CreateAssistantRunResponseBody) SetObject(v string) *CreateAssistantRunResponseBody {
	s.Object = &v
	return s
}

func (s *CreateAssistantRunResponseBody) SetStartedAt(v int64) *CreateAssistantRunResponseBody {
	s.StartedAt = &v
	return s
}

func (s *CreateAssistantRunResponseBody) SetStatus(v string) *CreateAssistantRunResponseBody {
	s.Status = &v
	return s
}

func (s *CreateAssistantRunResponseBody) SetThreadId(v string) *CreateAssistantRunResponseBody {
	s.ThreadId = &v
	return s
}

type CreateAssistantRunResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAssistantRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAssistantRunResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAssistantRunResponse) GoString() string {
	return s.String()
}

func (s *CreateAssistantRunResponse) SetHeaders(v map[string]*string) *CreateAssistantRunResponse {
	s.Headers = v
	return s
}

func (s *CreateAssistantRunResponse) SetStatusCode(v int32) *CreateAssistantRunResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAssistantRunResponse) SetBody(v *CreateAssistantRunResponseBody) *CreateAssistantRunResponse {
	s.Body = v
	return s
}

type CreateAssistantThreadHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateAssistantThreadHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateAssistantThreadHeaders) GoString() string {
	return s.String()
}

func (s *CreateAssistantThreadHeaders) SetCommonHeaders(v map[string]*string) *CreateAssistantThreadHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateAssistantThreadHeaders) SetXAcsDingtalkAccessToken(v string) *CreateAssistantThreadHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateAssistantThreadRequest struct {
	// if can be null:
	// true
	Metadata map[string]*string `json:"metadata,omitempty" xml:"metadata,omitempty"`
}

func (s CreateAssistantThreadRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAssistantThreadRequest) GoString() string {
	return s.String()
}

func (s *CreateAssistantThreadRequest) SetMetadata(v map[string]*string) *CreateAssistantThreadRequest {
	s.Metadata = v
	return s
}

type CreateAssistantThreadResponseBody struct {
	CreatedAt *int64                 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Id        *string                `json:"id,omitempty" xml:"id,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	Object    *string                `json:"object,omitempty" xml:"object,omitempty"`
}

func (s CreateAssistantThreadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAssistantThreadResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAssistantThreadResponseBody) SetCreatedAt(v int64) *CreateAssistantThreadResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateAssistantThreadResponseBody) SetId(v string) *CreateAssistantThreadResponseBody {
	s.Id = &v
	return s
}

func (s *CreateAssistantThreadResponseBody) SetMetadata(v map[string]interface{}) *CreateAssistantThreadResponseBody {
	s.Metadata = v
	return s
}

func (s *CreateAssistantThreadResponseBody) SetObject(v string) *CreateAssistantThreadResponseBody {
	s.Object = &v
	return s
}

type CreateAssistantThreadResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAssistantThreadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAssistantThreadResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAssistantThreadResponse) GoString() string {
	return s.String()
}

func (s *CreateAssistantThreadResponse) SetHeaders(v map[string]*string) *CreateAssistantThreadResponse {
	s.Headers = v
	return s
}

func (s *CreateAssistantThreadResponse) SetStatusCode(v int32) *CreateAssistantThreadResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAssistantThreadResponse) SetBody(v *CreateAssistantThreadResponseBody) *CreateAssistantThreadResponse {
	s.Body = v
	return s
}

type DeleteAssistantHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteAssistantHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteAssistantHeaders) GoString() string {
	return s.String()
}

func (s *DeleteAssistantHeaders) SetCommonHeaders(v map[string]*string) *DeleteAssistantHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteAssistantHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteAssistantHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteAssistantRequest struct {
	// This parameter is required.
	AssistantId *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	// This parameter is required.
	OperatorUnionId *string `json:"operatorUnionId,omitempty" xml:"operatorUnionId,omitempty"`
}

func (s DeleteAssistantRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAssistantRequest) GoString() string {
	return s.String()
}

func (s *DeleteAssistantRequest) SetAssistantId(v string) *DeleteAssistantRequest {
	s.AssistantId = &v
	return s
}

func (s *DeleteAssistantRequest) SetOperatorUnionId(v string) *DeleteAssistantRequest {
	s.OperatorUnionId = &v
	return s
}

type DeleteAssistantResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteAssistantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAssistantResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAssistantResponseBody) SetSuccess(v bool) *DeleteAssistantResponseBody {
	s.Success = &v
	return s
}

type DeleteAssistantResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAssistantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAssistantResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAssistantResponse) GoString() string {
	return s.String()
}

func (s *DeleteAssistantResponse) SetHeaders(v map[string]*string) *DeleteAssistantResponse {
	s.Headers = v
	return s
}

func (s *DeleteAssistantResponse) SetStatusCode(v int32) *DeleteAssistantResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAssistantResponse) SetBody(v *DeleteAssistantResponseBody) *DeleteAssistantResponse {
	s.Body = v
	return s
}

type DeleteAssistantMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteAssistantMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteAssistantMessageHeaders) GoString() string {
	return s.String()
}

func (s *DeleteAssistantMessageHeaders) SetCommonHeaders(v map[string]*string) *DeleteAssistantMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteAssistantMessageHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteAssistantMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteAssistantMessageResponseBody struct {
	Deleted *bool   `json:"deleted,omitempty" xml:"deleted,omitempty"`
	Id      *string `json:"id,omitempty" xml:"id,omitempty"`
	Object  *string `json:"object,omitempty" xml:"object,omitempty"`
}

func (s DeleteAssistantMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAssistantMessageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAssistantMessageResponseBody) SetDeleted(v bool) *DeleteAssistantMessageResponseBody {
	s.Deleted = &v
	return s
}

func (s *DeleteAssistantMessageResponseBody) SetId(v string) *DeleteAssistantMessageResponseBody {
	s.Id = &v
	return s
}

func (s *DeleteAssistantMessageResponseBody) SetObject(v string) *DeleteAssistantMessageResponseBody {
	s.Object = &v
	return s
}

type DeleteAssistantMessageResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAssistantMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAssistantMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAssistantMessageResponse) GoString() string {
	return s.String()
}

func (s *DeleteAssistantMessageResponse) SetHeaders(v map[string]*string) *DeleteAssistantMessageResponse {
	s.Headers = v
	return s
}

func (s *DeleteAssistantMessageResponse) SetStatusCode(v int32) *DeleteAssistantMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAssistantMessageResponse) SetBody(v *DeleteAssistantMessageResponseBody) *DeleteAssistantMessageResponse {
	s.Body = v
	return s
}

type DeleteAssistantThreadHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteAssistantThreadHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteAssistantThreadHeaders) GoString() string {
	return s.String()
}

func (s *DeleteAssistantThreadHeaders) SetCommonHeaders(v map[string]*string) *DeleteAssistantThreadHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteAssistantThreadHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteAssistantThreadHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteAssistantThreadResponseBody struct {
	Deleted *bool   `json:"deleted,omitempty" xml:"deleted,omitempty"`
	Id      *string `json:"id,omitempty" xml:"id,omitempty"`
	Object  *string `json:"object,omitempty" xml:"object,omitempty"`
}

func (s DeleteAssistantThreadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAssistantThreadResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAssistantThreadResponseBody) SetDeleted(v bool) *DeleteAssistantThreadResponseBody {
	s.Deleted = &v
	return s
}

func (s *DeleteAssistantThreadResponseBody) SetId(v string) *DeleteAssistantThreadResponseBody {
	s.Id = &v
	return s
}

func (s *DeleteAssistantThreadResponseBody) SetObject(v string) *DeleteAssistantThreadResponseBody {
	s.Object = &v
	return s
}

type DeleteAssistantThreadResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAssistantThreadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAssistantThreadResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAssistantThreadResponse) GoString() string {
	return s.String()
}

func (s *DeleteAssistantThreadResponse) SetHeaders(v map[string]*string) *DeleteAssistantThreadResponse {
	s.Headers = v
	return s
}

func (s *DeleteAssistantThreadResponse) SetStatusCode(v int32) *DeleteAssistantThreadResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAssistantThreadResponse) SetBody(v *DeleteAssistantThreadResponseBody) *DeleteAssistantThreadResponse {
	s.Body = v
	return s
}

type DeleteDomainWordsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteDomainWordsHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainWordsHeaders) GoString() string {
	return s.String()
}

func (s *DeleteDomainWordsHeaders) SetCommonHeaders(v map[string]*string) *DeleteDomainWordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteDomainWordsHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteDomainWordsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteDomainWordsRequest struct {
	// This parameter is required.
	AssistantId *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	// This parameter is required.
	DomainWords []*DeleteDomainWordsRequestDomainWords `json:"domainWords,omitempty" xml:"domainWords,omitempty" type:"Repeated"`
}

func (s DeleteDomainWordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainWordsRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainWordsRequest) SetAssistantId(v string) *DeleteDomainWordsRequest {
	s.AssistantId = &v
	return s
}

func (s *DeleteDomainWordsRequest) SetDomainWords(v []*DeleteDomainWordsRequestDomainWords) *DeleteDomainWordsRequest {
	s.DomainWords = v
	return s
}

type DeleteDomainWordsRequestDomainWords struct {
	Description *string   `json:"description,omitempty" xml:"description,omitempty"`
	DomainWord  *string   `json:"domainWord,omitempty" xml:"domainWord,omitempty"`
	FormalWords []*string `json:"formalWords,omitempty" xml:"formalWords,omitempty" type:"Repeated"`
}

func (s DeleteDomainWordsRequestDomainWords) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainWordsRequestDomainWords) GoString() string {
	return s.String()
}

func (s *DeleteDomainWordsRequestDomainWords) SetDescription(v string) *DeleteDomainWordsRequestDomainWords {
	s.Description = &v
	return s
}

func (s *DeleteDomainWordsRequestDomainWords) SetDomainWord(v string) *DeleteDomainWordsRequestDomainWords {
	s.DomainWord = &v
	return s
}

func (s *DeleteDomainWordsRequestDomainWords) SetFormalWords(v []*string) *DeleteDomainWordsRequestDomainWords {
	s.FormalWords = v
	return s
}

type DeleteDomainWordsResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteDomainWordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainWordsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainWordsResponseBody) SetSuccess(v bool) *DeleteDomainWordsResponseBody {
	s.Success = &v
	return s
}

type DeleteDomainWordsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDomainWordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDomainWordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainWordsResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainWordsResponse) SetHeaders(v map[string]*string) *DeleteDomainWordsResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainWordsResponse) SetStatusCode(v int32) *DeleteDomainWordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainWordsResponse) SetBody(v *DeleteDomainWordsResponseBody) *DeleteDomainWordsResponse {
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
	// This parameter is required.
	AssistantId *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	// This parameter is required.
	StudyId *string `json:"studyId,omitempty" xml:"studyId,omitempty"`
}

func (s DeleteKnowledgeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeRequest) SetAssistantId(v string) *DeleteKnowledgeRequest {
	s.AssistantId = &v
	return s
}

func (s *DeleteKnowledgeRequest) SetStudyId(v string) *DeleteKnowledgeRequest {
	s.StudyId = &v
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

type DeployAssistantHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeployAssistantHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeployAssistantHeaders) GoString() string {
	return s.String()
}

func (s *DeployAssistantHeaders) SetCommonHeaders(v map[string]*string) *DeployAssistantHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeployAssistantHeaders) SetXAcsDingtalkAccessToken(v string) *DeployAssistantHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeployAssistantRequest struct {
	Action           *DeployAssistantRequestAction    `json:"action,omitempty" xml:"action,omitempty" type:"Struct"`
	AiAssistantId    *string                          `json:"aiAssistantId,omitempty" xml:"aiAssistantId,omitempty"`
	AppScopes        *DeployAssistantRequestAppScopes `json:"appScopes,omitempty" xml:"appScopes,omitempty" type:"Struct"`
	Description      *string                          `json:"description,omitempty" xml:"description,omitempty"`
	Fallback         *DeployAssistantRequestFallback  `json:"fallback,omitempty" xml:"fallback,omitempty" type:"Struct"`
	Icon             *string                          `json:"icon,omitempty" xml:"icon,omitempty"`
	Instructions     *string                          `json:"instructions,omitempty" xml:"instructions,omitempty"`
	IsPublic         *int32                           `json:"isPublic,omitempty" xml:"isPublic,omitempty"`
	Name             *string                          `json:"name,omitempty" xml:"name,omitempty"`
	OperateUserId    *string                          `json:"operateUserId,omitempty" xml:"operateUserId,omitempty"`
	RecommendPrompts []*string                        `json:"recommendPrompts,omitempty" xml:"recommendPrompts,omitempty" type:"Repeated"`
	ShareRecipient   *string                          `json:"shareRecipient,omitempty" xml:"shareRecipient,omitempty"`
	ToneStyle        *string                          `json:"toneStyle,omitempty" xml:"toneStyle,omitempty"`
	Uuid             *string                          `json:"uuid,omitempty" xml:"uuid,omitempty"`
	WelcomeContent   *string                          `json:"welcomeContent,omitempty" xml:"welcomeContent,omitempty"`
	WelcomeTitle     *string                          `json:"welcomeTitle,omitempty" xml:"welcomeTitle,omitempty"`
}

func (s DeployAssistantRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployAssistantRequest) GoString() string {
	return s.String()
}

func (s *DeployAssistantRequest) SetAction(v *DeployAssistantRequestAction) *DeployAssistantRequest {
	s.Action = v
	return s
}

func (s *DeployAssistantRequest) SetAiAssistantId(v string) *DeployAssistantRequest {
	s.AiAssistantId = &v
	return s
}

func (s *DeployAssistantRequest) SetAppScopes(v *DeployAssistantRequestAppScopes) *DeployAssistantRequest {
	s.AppScopes = v
	return s
}

func (s *DeployAssistantRequest) SetDescription(v string) *DeployAssistantRequest {
	s.Description = &v
	return s
}

func (s *DeployAssistantRequest) SetFallback(v *DeployAssistantRequestFallback) *DeployAssistantRequest {
	s.Fallback = v
	return s
}

func (s *DeployAssistantRequest) SetIcon(v string) *DeployAssistantRequest {
	s.Icon = &v
	return s
}

func (s *DeployAssistantRequest) SetInstructions(v string) *DeployAssistantRequest {
	s.Instructions = &v
	return s
}

func (s *DeployAssistantRequest) SetIsPublic(v int32) *DeployAssistantRequest {
	s.IsPublic = &v
	return s
}

func (s *DeployAssistantRequest) SetName(v string) *DeployAssistantRequest {
	s.Name = &v
	return s
}

func (s *DeployAssistantRequest) SetOperateUserId(v string) *DeployAssistantRequest {
	s.OperateUserId = &v
	return s
}

func (s *DeployAssistantRequest) SetRecommendPrompts(v []*string) *DeployAssistantRequest {
	s.RecommendPrompts = v
	return s
}

func (s *DeployAssistantRequest) SetShareRecipient(v string) *DeployAssistantRequest {
	s.ShareRecipient = &v
	return s
}

func (s *DeployAssistantRequest) SetToneStyle(v string) *DeployAssistantRequest {
	s.ToneStyle = &v
	return s
}

func (s *DeployAssistantRequest) SetUuid(v string) *DeployAssistantRequest {
	s.Uuid = &v
	return s
}

func (s *DeployAssistantRequest) SetWelcomeContent(v string) *DeployAssistantRequest {
	s.WelcomeContent = &v
	return s
}

func (s *DeployAssistantRequest) SetWelcomeTitle(v string) *DeployAssistantRequest {
	s.WelcomeTitle = &v
	return s
}

type DeployAssistantRequestAction struct {
	ActionAuthInfo *DeployAssistantRequestActionActionAuthInfo `json:"actionAuthInfo,omitempty" xml:"actionAuthInfo,omitempty" type:"Struct"`
	ActionName     *string                                     `json:"actionName,omitempty" xml:"actionName,omitempty"`
	Description    *string                                     `json:"description,omitempty" xml:"description,omitempty"`
	Schema         *string                                     `json:"schema,omitempty" xml:"schema,omitempty"`
}

func (s DeployAssistantRequestAction) String() string {
	return tea.Prettify(s)
}

func (s DeployAssistantRequestAction) GoString() string {
	return s.String()
}

func (s *DeployAssistantRequestAction) SetActionAuthInfo(v *DeployAssistantRequestActionActionAuthInfo) *DeployAssistantRequestAction {
	s.ActionAuthInfo = v
	return s
}

func (s *DeployAssistantRequestAction) SetActionName(v string) *DeployAssistantRequestAction {
	s.ActionName = &v
	return s
}

func (s *DeployAssistantRequestAction) SetDescription(v string) *DeployAssistantRequestAction {
	s.Description = &v
	return s
}

func (s *DeployAssistantRequestAction) SetSchema(v string) *DeployAssistantRequestAction {
	s.Schema = &v
	return s
}

type DeployAssistantRequestActionActionAuthInfo struct {
	AuthConfig         *string `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	AuthenticationType *string `json:"authenticationType,omitempty" xml:"authenticationType,omitempty"`
}

func (s DeployAssistantRequestActionActionAuthInfo) String() string {
	return tea.Prettify(s)
}

func (s DeployAssistantRequestActionActionAuthInfo) GoString() string {
	return s.String()
}

func (s *DeployAssistantRequestActionActionAuthInfo) SetAuthConfig(v string) *DeployAssistantRequestActionActionAuthInfo {
	s.AuthConfig = &v
	return s
}

func (s *DeployAssistantRequestActionActionAuthInfo) SetAuthenticationType(v string) *DeployAssistantRequestActionActionAuthInfo {
	s.AuthenticationType = &v
	return s
}

type DeployAssistantRequestAppScopes struct {
	DeptVisibleScopes  *string `json:"deptVisibleScopes,omitempty" xml:"deptVisibleScopes,omitempty"`
	DynamicGroupScopes *string `json:"dynamicGroupScopes,omitempty" xml:"dynamicGroupScopes,omitempty"`
	IsHidden           *bool   `json:"isHidden,omitempty" xml:"isHidden,omitempty"`
	RoleVisibleScopes  *string `json:"roleVisibleScopes,omitempty" xml:"roleVisibleScopes,omitempty"`
	UserVisibleScopes  *string `json:"userVisibleScopes,omitempty" xml:"userVisibleScopes,omitempty"`
}

func (s DeployAssistantRequestAppScopes) String() string {
	return tea.Prettify(s)
}

func (s DeployAssistantRequestAppScopes) GoString() string {
	return s.String()
}

func (s *DeployAssistantRequestAppScopes) SetDeptVisibleScopes(v string) *DeployAssistantRequestAppScopes {
	s.DeptVisibleScopes = &v
	return s
}

func (s *DeployAssistantRequestAppScopes) SetDynamicGroupScopes(v string) *DeployAssistantRequestAppScopes {
	s.DynamicGroupScopes = &v
	return s
}

func (s *DeployAssistantRequestAppScopes) SetIsHidden(v bool) *DeployAssistantRequestAppScopes {
	s.IsHidden = &v
	return s
}

func (s *DeployAssistantRequestAppScopes) SetRoleVisibleScopes(v string) *DeployAssistantRequestAppScopes {
	s.RoleVisibleScopes = &v
	return s
}

func (s *DeployAssistantRequestAppScopes) SetUserVisibleScopes(v string) *DeployAssistantRequestAppScopes {
	s.UserVisibleScopes = &v
	return s
}

type DeployAssistantRequestFallback struct {
	ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	DefaultMsg *string `json:"defaultMsg,omitempty" xml:"defaultMsg,omitempty"`
}

func (s DeployAssistantRequestFallback) String() string {
	return tea.Prettify(s)
}

func (s DeployAssistantRequestFallback) GoString() string {
	return s.String()
}

func (s *DeployAssistantRequestFallback) SetActionType(v string) *DeployAssistantRequestFallback {
	s.ActionType = &v
	return s
}

func (s *DeployAssistantRequestFallback) SetDefaultMsg(v string) *DeployAssistantRequestFallback {
	s.DefaultMsg = &v
	return s
}

type DeployAssistantResponseBody struct {
	AssistantId *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
}

func (s DeployAssistantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployAssistantResponseBody) GoString() string {
	return s.String()
}

func (s *DeployAssistantResponseBody) SetAssistantId(v string) *DeployAssistantResponseBody {
	s.AssistantId = &v
	return s
}

type DeployAssistantResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployAssistantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployAssistantResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployAssistantResponse) GoString() string {
	return s.String()
}

func (s *DeployAssistantResponse) SetHeaders(v map[string]*string) *DeployAssistantResponse {
	s.Headers = v
	return s
}

func (s *DeployAssistantResponse) SetStatusCode(v int32) *DeployAssistantResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployAssistantResponse) SetBody(v *DeployAssistantResponseBody) *DeployAssistantResponse {
	s.Body = v
	return s
}

type GetAskDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAskDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAskDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetAskDetailHeaders) SetCommonHeaders(v map[string]*string) *GetAskDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAskDetailHeaders) SetXAcsDingtalkAccessToken(v string) *GetAskDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAskDetailRequest struct {
	// This parameter is required.
	AssistantId *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	EndTime     *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	Offset *int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// This parameter is required.
	PageSize  *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetAskDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAskDetailRequest) GoString() string {
	return s.String()
}

func (s *GetAskDetailRequest) SetAssistantId(v string) *GetAskDetailRequest {
	s.AssistantId = &v
	return s
}

func (s *GetAskDetailRequest) SetEndTime(v int64) *GetAskDetailRequest {
	s.EndTime = &v
	return s
}

func (s *GetAskDetailRequest) SetOffset(v int64) *GetAskDetailRequest {
	s.Offset = &v
	return s
}

func (s *GetAskDetailRequest) SetPageSize(v int32) *GetAskDetailRequest {
	s.PageSize = &v
	return s
}

func (s *GetAskDetailRequest) SetStartTime(v int64) *GetAskDetailRequest {
	s.StartTime = &v
	return s
}

type GetAskDetailResponseBody struct {
	Result  *GetAskDetailResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetAskDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetAskDetailResponseBody) SetResult(v *GetAskDetailResponseBodyResult) *GetAskDetailResponseBody {
	s.Result = v
	return s
}

func (s *GetAskDetailResponseBody) SetSuccess(v bool) *GetAskDetailResponseBody {
	s.Success = &v
	return s
}

type GetAskDetailResponseBodyResult struct {
	HasMore    *bool                                 `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*GetAskDetailResponseBodyResultList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextCursor *int64                                `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	TotalCount *int32                                `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetAskDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetAskDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAskDetailResponseBodyResult) SetHasMore(v bool) *GetAskDetailResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *GetAskDetailResponseBodyResult) SetList(v []*GetAskDetailResponseBodyResultList) *GetAskDetailResponseBodyResult {
	s.List = v
	return s
}

func (s *GetAskDetailResponseBodyResult) SetNextCursor(v int64) *GetAskDetailResponseBodyResult {
	s.NextCursor = &v
	return s
}

func (s *GetAskDetailResponseBodyResult) SetTotalCount(v int32) *GetAskDetailResponseBodyResult {
	s.TotalCount = &v
	return s
}

type GetAskDetailResponseBodyResultList struct {
	Answer         *string                                         `json:"answer,omitempty" xml:"answer,omitempty"`
	AnswerResult   *string                                         `json:"answerResult,omitempty" xml:"answerResult,omitempty"`
	CommentTags    []*string                                       `json:"commentTags,omitempty" xml:"commentTags,omitempty" type:"Repeated"`
	IsMarkResolved *bool                                           `json:"isMarkResolved,omitempty" xml:"isMarkResolved,omitempty"`
	Nick           *string                                         `json:"nick,omitempty" xml:"nick,omitempty"`
	Question       *string                                         `json:"question,omitempty" xml:"question,omitempty"`
	References     []*GetAskDetailResponseBodyResultListReferences `json:"references,omitempty" xml:"references,omitempty" type:"Repeated"`
	Time           *int64                                          `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetAskDetailResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s GetAskDetailResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *GetAskDetailResponseBodyResultList) SetAnswer(v string) *GetAskDetailResponseBodyResultList {
	s.Answer = &v
	return s
}

func (s *GetAskDetailResponseBodyResultList) SetAnswerResult(v string) *GetAskDetailResponseBodyResultList {
	s.AnswerResult = &v
	return s
}

func (s *GetAskDetailResponseBodyResultList) SetCommentTags(v []*string) *GetAskDetailResponseBodyResultList {
	s.CommentTags = v
	return s
}

func (s *GetAskDetailResponseBodyResultList) SetIsMarkResolved(v bool) *GetAskDetailResponseBodyResultList {
	s.IsMarkResolved = &v
	return s
}

func (s *GetAskDetailResponseBodyResultList) SetNick(v string) *GetAskDetailResponseBodyResultList {
	s.Nick = &v
	return s
}

func (s *GetAskDetailResponseBodyResultList) SetQuestion(v string) *GetAskDetailResponseBodyResultList {
	s.Question = &v
	return s
}

func (s *GetAskDetailResponseBodyResultList) SetReferences(v []*GetAskDetailResponseBodyResultListReferences) *GetAskDetailResponseBodyResultList {
	s.References = v
	return s
}

func (s *GetAskDetailResponseBodyResultList) SetTime(v int64) *GetAskDetailResponseBodyResultList {
	s.Time = &v
	return s
}

type GetAskDetailResponseBodyResultListReferences struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Url  *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetAskDetailResponseBodyResultListReferences) String() string {
	return tea.Prettify(s)
}

func (s GetAskDetailResponseBodyResultListReferences) GoString() string {
	return s.String()
}

func (s *GetAskDetailResponseBodyResultListReferences) SetName(v string) *GetAskDetailResponseBodyResultListReferences {
	s.Name = &v
	return s
}

func (s *GetAskDetailResponseBodyResultListReferences) SetUrl(v string) *GetAskDetailResponseBodyResultListReferences {
	s.Url = &v
	return s
}

type GetAskDetailResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAskDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAskDetailResponse) GoString() string {
	return s.String()
}

func (s *GetAskDetailResponse) SetHeaders(v map[string]*string) *GetAskDetailResponse {
	s.Headers = v
	return s
}

func (s *GetAskDetailResponse) SetStatusCode(v int32) *GetAskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAskDetailResponse) SetBody(v *GetAskDetailResponseBody) *GetAskDetailResponse {
	s.Body = v
	return s
}

type GetAssistantActionInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAssistantActionInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAssistantActionInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetAssistantActionInfoHeaders) SetCommonHeaders(v map[string]*string) *GetAssistantActionInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAssistantActionInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetAssistantActionInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAssistantActionInfoRequest struct {
	// This parameter is required.
	AssistantId *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
}

func (s GetAssistantActionInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAssistantActionInfoRequest) GoString() string {
	return s.String()
}

func (s *GetAssistantActionInfoRequest) SetAssistantId(v string) *GetAssistantActionInfoRequest {
	s.AssistantId = &v
	return s
}

type GetAssistantActionInfoResponseBody struct {
	ActionList  []*GetAssistantActionInfoResponseBodyActionList `json:"actionList,omitempty" xml:"actionList,omitempty" type:"Repeated"`
	AssistantId *string                                         `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	CorpId      *string                                         `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s GetAssistantActionInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAssistantActionInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAssistantActionInfoResponseBody) SetActionList(v []*GetAssistantActionInfoResponseBodyActionList) *GetAssistantActionInfoResponseBody {
	s.ActionList = v
	return s
}

func (s *GetAssistantActionInfoResponseBody) SetAssistantId(v string) *GetAssistantActionInfoResponseBody {
	s.AssistantId = &v
	return s
}

func (s *GetAssistantActionInfoResponseBody) SetCorpId(v string) *GetAssistantActionInfoResponseBody {
	s.CorpId = &v
	return s
}

type GetAssistantActionInfoResponseBodyActionList struct {
	ActionId      *string `json:"actionId,omitempty" xml:"actionId,omitempty"`
	ActionName    *string `json:"actionName,omitempty" xml:"actionName,omitempty"`
	ActionVersion *string `json:"actionVersion,omitempty" xml:"actionVersion,omitempty"`
	Description   *string `json:"description,omitempty" xml:"description,omitempty"`
	Icon          *string `json:"icon,omitempty" xml:"icon,omitempty"`
}

func (s GetAssistantActionInfoResponseBodyActionList) String() string {
	return tea.Prettify(s)
}

func (s GetAssistantActionInfoResponseBodyActionList) GoString() string {
	return s.String()
}

func (s *GetAssistantActionInfoResponseBodyActionList) SetActionId(v string) *GetAssistantActionInfoResponseBodyActionList {
	s.ActionId = &v
	return s
}

func (s *GetAssistantActionInfoResponseBodyActionList) SetActionName(v string) *GetAssistantActionInfoResponseBodyActionList {
	s.ActionName = &v
	return s
}

func (s *GetAssistantActionInfoResponseBodyActionList) SetActionVersion(v string) *GetAssistantActionInfoResponseBodyActionList {
	s.ActionVersion = &v
	return s
}

func (s *GetAssistantActionInfoResponseBodyActionList) SetDescription(v string) *GetAssistantActionInfoResponseBodyActionList {
	s.Description = &v
	return s
}

func (s *GetAssistantActionInfoResponseBodyActionList) SetIcon(v string) *GetAssistantActionInfoResponseBodyActionList {
	s.Icon = &v
	return s
}

type GetAssistantActionInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAssistantActionInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAssistantActionInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAssistantActionInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAssistantActionInfoResponse) SetHeaders(v map[string]*string) *GetAssistantActionInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAssistantActionInfoResponse) SetStatusCode(v int32) *GetAssistantActionInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAssistantActionInfoResponse) SetBody(v *GetAssistantActionInfoResponseBody) *GetAssistantActionInfoResponse {
	s.Body = v
	return s
}

type GetDomainWordsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDomainWordsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDomainWordsHeaders) GoString() string {
	return s.String()
}

func (s *GetDomainWordsHeaders) SetCommonHeaders(v map[string]*string) *GetDomainWordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDomainWordsHeaders) SetXAcsDingtalkAccessToken(v string) *GetDomainWordsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDomainWordsRequest struct {
	// This parameter is required.
	AssistantId *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
}

func (s GetDomainWordsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDomainWordsRequest) GoString() string {
	return s.String()
}

func (s *GetDomainWordsRequest) SetAssistantId(v string) *GetDomainWordsRequest {
	s.AssistantId = &v
	return s
}

type GetDomainWordsResponseBody struct {
	Result  []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success *bool     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDomainWordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDomainWordsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDomainWordsResponseBody) SetResult(v []*string) *GetDomainWordsResponseBody {
	s.Result = v
	return s
}

func (s *GetDomainWordsResponseBody) SetSuccess(v bool) *GetDomainWordsResponseBody {
	s.Success = &v
	return s
}

type GetDomainWordsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDomainWordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDomainWordsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDomainWordsResponse) GoString() string {
	return s.String()
}

func (s *GetDomainWordsResponse) SetHeaders(v map[string]*string) *GetDomainWordsResponse {
	s.Headers = v
	return s
}

func (s *GetDomainWordsResponse) SetStatusCode(v int32) *GetDomainWordsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDomainWordsResponse) SetBody(v *GetDomainWordsResponseBody) *GetDomainWordsResponse {
	s.Body = v
	return s
}

type GetKnowledgeListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetKnowledgeListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetKnowledgeListHeaders) GoString() string {
	return s.String()
}

func (s *GetKnowledgeListHeaders) SetCommonHeaders(v map[string]*string) *GetKnowledgeListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetKnowledgeListHeaders) SetXAcsDingtalkAccessToken(v string) *GetKnowledgeListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetKnowledgeListRequest struct {
	AssistantId *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
}

func (s GetKnowledgeListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetKnowledgeListRequest) GoString() string {
	return s.String()
}

func (s *GetKnowledgeListRequest) SetAssistantId(v string) *GetKnowledgeListRequest {
	s.AssistantId = &v
	return s
}

type GetKnowledgeListResponseBody struct {
	Result  []*GetKnowledgeListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success *bool                                 `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetKnowledgeListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetKnowledgeListResponseBody) GoString() string {
	return s.String()
}

func (s *GetKnowledgeListResponseBody) SetResult(v []*GetKnowledgeListResponseBodyResult) *GetKnowledgeListResponseBody {
	s.Result = v
	return s
}

func (s *GetKnowledgeListResponseBody) SetSuccess(v bool) *GetKnowledgeListResponseBody {
	s.Success = &v
	return s
}

type GetKnowledgeListResponseBodyResult struct {
	DocFormat *string `json:"docFormat,omitempty" xml:"docFormat,omitempty"`
	DocName   *string `json:"docName,omitempty" xml:"docName,omitempty"`
	DocUrl    *string `json:"docUrl,omitempty" xml:"docUrl,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
	StudyId   *string `json:"studyId,omitempty" xml:"studyId,omitempty"`
}

func (s GetKnowledgeListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetKnowledgeListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetKnowledgeListResponseBodyResult) SetDocFormat(v string) *GetKnowledgeListResponseBodyResult {
	s.DocFormat = &v
	return s
}

func (s *GetKnowledgeListResponseBodyResult) SetDocName(v string) *GetKnowledgeListResponseBodyResult {
	s.DocName = &v
	return s
}

func (s *GetKnowledgeListResponseBodyResult) SetDocUrl(v string) *GetKnowledgeListResponseBodyResult {
	s.DocUrl = &v
	return s
}

func (s *GetKnowledgeListResponseBodyResult) SetStatus(v string) *GetKnowledgeListResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetKnowledgeListResponseBodyResult) SetStudyId(v string) *GetKnowledgeListResponseBodyResult {
	s.StudyId = &v
	return s
}

type GetKnowledgeListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKnowledgeListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKnowledgeListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetKnowledgeListResponse) GoString() string {
	return s.String()
}

func (s *GetKnowledgeListResponse) SetHeaders(v map[string]*string) *GetKnowledgeListResponse {
	s.Headers = v
	return s
}

func (s *GetKnowledgeListResponse) SetStatusCode(v int32) *GetKnowledgeListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKnowledgeListResponse) SetBody(v *GetKnowledgeListResponseBody) *GetKnowledgeListResponse {
	s.Body = v
	return s
}

type InstallAssistantHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InstallAssistantHeaders) String() string {
	return tea.Prettify(s)
}

func (s InstallAssistantHeaders) GoString() string {
	return s.String()
}

func (s *InstallAssistantHeaders) SetCommonHeaders(v map[string]*string) *InstallAssistantHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InstallAssistantHeaders) SetXAcsDingtalkAccessToken(v string) *InstallAssistantHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InstallAssistantRequest struct {
	AssistantId           *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	IsAllOrgMemberVisible *bool   `json:"isAllOrgMemberVisible,omitempty" xml:"isAllOrgMemberVisible,omitempty"`
}

func (s InstallAssistantRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallAssistantRequest) GoString() string {
	return s.String()
}

func (s *InstallAssistantRequest) SetAssistantId(v string) *InstallAssistantRequest {
	s.AssistantId = &v
	return s
}

func (s *InstallAssistantRequest) SetIsAllOrgMemberVisible(v bool) *InstallAssistantRequest {
	s.IsAllOrgMemberVisible = &v
	return s
}

type InstallAssistantResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s InstallAssistantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallAssistantResponseBody) GoString() string {
	return s.String()
}

func (s *InstallAssistantResponseBody) SetSuccess(v bool) *InstallAssistantResponseBody {
	s.Success = &v
	return s
}

type InstallAssistantResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallAssistantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallAssistantResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallAssistantResponse) GoString() string {
	return s.String()
}

func (s *InstallAssistantResponse) SetHeaders(v map[string]*string) *InstallAssistantResponse {
	s.Headers = v
	return s
}

func (s *InstallAssistantResponse) SetStatusCode(v int32) *InstallAssistantResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallAssistantResponse) SetBody(v *InstallAssistantResponseBody) *InstallAssistantResponse {
	s.Body = v
	return s
}

type LearnKnowledgeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s LearnKnowledgeHeaders) String() string {
	return tea.Prettify(s)
}

func (s LearnKnowledgeHeaders) GoString() string {
	return s.String()
}

func (s *LearnKnowledgeHeaders) SetCommonHeaders(v map[string]*string) *LearnKnowledgeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *LearnKnowledgeHeaders) SetXAcsDingtalkAccessToken(v string) *LearnKnowledgeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type LearnKnowledgeRequest struct {
	AssistantId *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	DocUrl      *string `json:"docUrl,omitempty" xml:"docUrl,omitempty"`
}

func (s LearnKnowledgeRequest) String() string {
	return tea.Prettify(s)
}

func (s LearnKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *LearnKnowledgeRequest) SetAssistantId(v string) *LearnKnowledgeRequest {
	s.AssistantId = &v
	return s
}

func (s *LearnKnowledgeRequest) SetDocUrl(v string) *LearnKnowledgeRequest {
	s.DocUrl = &v
	return s
}

type LearnKnowledgeResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s LearnKnowledgeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LearnKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *LearnKnowledgeResponseBody) SetSuccess(v bool) *LearnKnowledgeResponseBody {
	s.Success = &v
	return s
}

type LearnKnowledgeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LearnKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LearnKnowledgeResponse) String() string {
	return tea.Prettify(s)
}

func (s LearnKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *LearnKnowledgeResponse) SetHeaders(v map[string]*string) *LearnKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *LearnKnowledgeResponse) SetStatusCode(v int32) *LearnKnowledgeResponse {
	s.StatusCode = &v
	return s
}

func (s *LearnKnowledgeResponse) SetBody(v *LearnKnowledgeResponseBody) *LearnKnowledgeResponse {
	s.Body = v
	return s
}

type ListAssistantHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListAssistantHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListAssistantHeaders) GoString() string {
	return s.String()
}

func (s *ListAssistantHeaders) SetCommonHeaders(v map[string]*string) *ListAssistantHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAssistantHeaders) SetXAcsDingtalkAccessToken(v string) *ListAssistantHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListAssistantRequest struct {
	Cursor *int64 `json:"cursor,omitempty" xml:"cursor,omitempty"`
	// This parameter is required.
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s ListAssistantRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAssistantRequest) GoString() string {
	return s.String()
}

func (s *ListAssistantRequest) SetCursor(v int64) *ListAssistantRequest {
	s.Cursor = &v
	return s
}

func (s *ListAssistantRequest) SetPageSize(v int32) *ListAssistantRequest {
	s.PageSize = &v
	return s
}

func (s *ListAssistantRequest) SetUnionId(v string) *ListAssistantRequest {
	s.UnionId = &v
	return s
}

type ListAssistantResponseBody struct {
	HasMore    *bool                            `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*ListAssistantResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextCursor *int64                           `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	TotalCount *int32                           `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListAssistantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAssistantResponseBody) GoString() string {
	return s.String()
}

func (s *ListAssistantResponseBody) SetHasMore(v bool) *ListAssistantResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListAssistantResponseBody) SetList(v []*ListAssistantResponseBodyList) *ListAssistantResponseBody {
	s.List = v
	return s
}

func (s *ListAssistantResponseBody) SetNextCursor(v int64) *ListAssistantResponseBody {
	s.NextCursor = &v
	return s
}

func (s *ListAssistantResponseBody) SetTotalCount(v int32) *ListAssistantResponseBody {
	s.TotalCount = &v
	return s
}

type ListAssistantResponseBodyList struct {
	AssistantId    *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	CreatedAt      *int64  `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatorUnionId *string `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	Icon           *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListAssistantResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListAssistantResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListAssistantResponseBodyList) SetAssistantId(v string) *ListAssistantResponseBodyList {
	s.AssistantId = &v
	return s
}

func (s *ListAssistantResponseBodyList) SetCreatedAt(v int64) *ListAssistantResponseBodyList {
	s.CreatedAt = &v
	return s
}

func (s *ListAssistantResponseBodyList) SetCreatorUnionId(v string) *ListAssistantResponseBodyList {
	s.CreatorUnionId = &v
	return s
}

func (s *ListAssistantResponseBodyList) SetDescription(v string) *ListAssistantResponseBodyList {
	s.Description = &v
	return s
}

func (s *ListAssistantResponseBodyList) SetIcon(v string) *ListAssistantResponseBodyList {
	s.Icon = &v
	return s
}

func (s *ListAssistantResponseBodyList) SetName(v string) *ListAssistantResponseBodyList {
	s.Name = &v
	return s
}

type ListAssistantResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAssistantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAssistantResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAssistantResponse) GoString() string {
	return s.String()
}

func (s *ListAssistantResponse) SetHeaders(v map[string]*string) *ListAssistantResponse {
	s.Headers = v
	return s
}

func (s *ListAssistantResponse) SetStatusCode(v int32) *ListAssistantResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAssistantResponse) SetBody(v *ListAssistantResponseBody) *ListAssistantResponse {
	s.Body = v
	return s
}

type ListAssistantMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListAssistantMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListAssistantMessageHeaders) GoString() string {
	return s.String()
}

func (s *ListAssistantMessageHeaders) SetCommonHeaders(v map[string]*string) *ListAssistantMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAssistantMessageHeaders) SetXAcsDingtalkAccessToken(v string) *ListAssistantMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListAssistantMessageRequest struct {
	Limit *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
}

func (s ListAssistantMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAssistantMessageRequest) GoString() string {
	return s.String()
}

func (s *ListAssistantMessageRequest) SetLimit(v int32) *ListAssistantMessageRequest {
	s.Limit = &v
	return s
}

func (s *ListAssistantMessageRequest) SetOrder(v string) *ListAssistantMessageRequest {
	s.Order = &v
	return s
}

func (s *ListAssistantMessageRequest) SetRunId(v string) *ListAssistantMessageRequest {
	s.RunId = &v
	return s
}

type ListAssistantMessageResponseBody struct {
	Data   []*ListAssistantMessageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Object *string                                 `json:"object,omitempty" xml:"object,omitempty"`
}

func (s ListAssistantMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAssistantMessageResponseBody) GoString() string {
	return s.String()
}

func (s *ListAssistantMessageResponseBody) SetData(v []*ListAssistantMessageResponseBodyData) *ListAssistantMessageResponseBody {
	s.Data = v
	return s
}

func (s *ListAssistantMessageResponseBody) SetObject(v string) *ListAssistantMessageResponseBody {
	s.Object = &v
	return s
}

type ListAssistantMessageResponseBodyData struct {
	AssistantId *string                `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	Content     []interface{}          `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	CreatedAt   *int64                 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Id          *string                `json:"id,omitempty" xml:"id,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	Object      *string                `json:"object,omitempty" xml:"object,omitempty"`
	Role        *string                `json:"role,omitempty" xml:"role,omitempty"`
	RunId       *string                `json:"runId,omitempty" xml:"runId,omitempty"`
	ThreadId    *string                `json:"threadId,omitempty" xml:"threadId,omitempty"`
}

func (s ListAssistantMessageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAssistantMessageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAssistantMessageResponseBodyData) SetAssistantId(v string) *ListAssistantMessageResponseBodyData {
	s.AssistantId = &v
	return s
}

func (s *ListAssistantMessageResponseBodyData) SetContent(v []interface{}) *ListAssistantMessageResponseBodyData {
	s.Content = v
	return s
}

func (s *ListAssistantMessageResponseBodyData) SetCreatedAt(v int64) *ListAssistantMessageResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *ListAssistantMessageResponseBodyData) SetId(v string) *ListAssistantMessageResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAssistantMessageResponseBodyData) SetMetadata(v map[string]interface{}) *ListAssistantMessageResponseBodyData {
	s.Metadata = v
	return s
}

func (s *ListAssistantMessageResponseBodyData) SetObject(v string) *ListAssistantMessageResponseBodyData {
	s.Object = &v
	return s
}

func (s *ListAssistantMessageResponseBodyData) SetRole(v string) *ListAssistantMessageResponseBodyData {
	s.Role = &v
	return s
}

func (s *ListAssistantMessageResponseBodyData) SetRunId(v string) *ListAssistantMessageResponseBodyData {
	s.RunId = &v
	return s
}

func (s *ListAssistantMessageResponseBodyData) SetThreadId(v string) *ListAssistantMessageResponseBodyData {
	s.ThreadId = &v
	return s
}

type ListAssistantMessageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAssistantMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAssistantMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAssistantMessageResponse) GoString() string {
	return s.String()
}

func (s *ListAssistantMessageResponse) SetHeaders(v map[string]*string) *ListAssistantMessageResponse {
	s.Headers = v
	return s
}

func (s *ListAssistantMessageResponse) SetStatusCode(v int32) *ListAssistantMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAssistantMessageResponse) SetBody(v *ListAssistantMessageResponseBody) *ListAssistantMessageResponse {
	s.Body = v
	return s
}

type ListAssistantRunHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListAssistantRunHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListAssistantRunHeaders) GoString() string {
	return s.String()
}

func (s *ListAssistantRunHeaders) SetCommonHeaders(v map[string]*string) *ListAssistantRunHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAssistantRunHeaders) SetXAcsDingtalkAccessToken(v string) *ListAssistantRunHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListAssistantRunRequest struct {
	Limit *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
}

func (s ListAssistantRunRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAssistantRunRequest) GoString() string {
	return s.String()
}

func (s *ListAssistantRunRequest) SetLimit(v int32) *ListAssistantRunRequest {
	s.Limit = &v
	return s
}

func (s *ListAssistantRunRequest) SetOrder(v string) *ListAssistantRunRequest {
	s.Order = &v
	return s
}

type ListAssistantRunResponseBody struct {
	Data   []*ListAssistantRunResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Object *string                             `json:"object,omitempty" xml:"object,omitempty"`
}

func (s ListAssistantRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAssistantRunResponseBody) GoString() string {
	return s.String()
}

func (s *ListAssistantRunResponseBody) SetData(v []*ListAssistantRunResponseBodyData) *ListAssistantRunResponseBody {
	s.Data = v
	return s
}

func (s *ListAssistantRunResponseBody) SetObject(v string) *ListAssistantRunResponseBody {
	s.Object = &v
	return s
}

type ListAssistantRunResponseBodyData struct {
	AssistantId  *string                `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	CancelledAt  *int64                 `json:"cancelledAt,omitempty" xml:"cancelledAt,omitempty"`
	CompletedAt  *int64                 `json:"completedAt,omitempty" xml:"completedAt,omitempty"`
	CreatedAt    *int64                 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	ExpiresAt    *int64                 `json:"expiresAt,omitempty" xml:"expiresAt,omitempty"`
	FailedAt     *int64                 `json:"failedAt,omitempty" xml:"failedAt,omitempty"`
	Id           *string                `json:"id,omitempty" xml:"id,omitempty"`
	LastErrorMsg *string                `json:"lastErrorMsg,omitempty" xml:"lastErrorMsg,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	Object       *string                `json:"object,omitempty" xml:"object,omitempty"`
	StartedAt    *int64                 `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	Status       *string                `json:"status,omitempty" xml:"status,omitempty"`
	ThreadId     *string                `json:"threadId,omitempty" xml:"threadId,omitempty"`
}

func (s ListAssistantRunResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAssistantRunResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAssistantRunResponseBodyData) SetAssistantId(v string) *ListAssistantRunResponseBodyData {
	s.AssistantId = &v
	return s
}

func (s *ListAssistantRunResponseBodyData) SetCancelledAt(v int64) *ListAssistantRunResponseBodyData {
	s.CancelledAt = &v
	return s
}

func (s *ListAssistantRunResponseBodyData) SetCompletedAt(v int64) *ListAssistantRunResponseBodyData {
	s.CompletedAt = &v
	return s
}

func (s *ListAssistantRunResponseBodyData) SetCreatedAt(v int64) *ListAssistantRunResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *ListAssistantRunResponseBodyData) SetExpiresAt(v int64) *ListAssistantRunResponseBodyData {
	s.ExpiresAt = &v
	return s
}

func (s *ListAssistantRunResponseBodyData) SetFailedAt(v int64) *ListAssistantRunResponseBodyData {
	s.FailedAt = &v
	return s
}

func (s *ListAssistantRunResponseBodyData) SetId(v string) *ListAssistantRunResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAssistantRunResponseBodyData) SetLastErrorMsg(v string) *ListAssistantRunResponseBodyData {
	s.LastErrorMsg = &v
	return s
}

func (s *ListAssistantRunResponseBodyData) SetMetadata(v map[string]interface{}) *ListAssistantRunResponseBodyData {
	s.Metadata = v
	return s
}

func (s *ListAssistantRunResponseBodyData) SetObject(v string) *ListAssistantRunResponseBodyData {
	s.Object = &v
	return s
}

func (s *ListAssistantRunResponseBodyData) SetStartedAt(v int64) *ListAssistantRunResponseBodyData {
	s.StartedAt = &v
	return s
}

func (s *ListAssistantRunResponseBodyData) SetStatus(v string) *ListAssistantRunResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListAssistantRunResponseBodyData) SetThreadId(v string) *ListAssistantRunResponseBodyData {
	s.ThreadId = &v
	return s
}

type ListAssistantRunResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAssistantRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAssistantRunResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAssistantRunResponse) GoString() string {
	return s.String()
}

func (s *ListAssistantRunResponse) SetHeaders(v map[string]*string) *ListAssistantRunResponse {
	s.Headers = v
	return s
}

func (s *ListAssistantRunResponse) SetStatusCode(v int32) *ListAssistantRunResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAssistantRunResponse) SetBody(v *ListAssistantRunResponseBody) *ListAssistantRunResponse {
	s.Body = v
	return s
}

type ListInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceHeaders) GoString() string {
	return s.String()
}

func (s *ListInstanceHeaders) SetCommonHeaders(v map[string]*string) *ListInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *ListInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListInstanceRequest struct {
	// This parameter is required.
	PrototypeAssistantId *string `json:"prototypeAssistantId,omitempty" xml:"prototypeAssistantId,omitempty"`
}

func (s ListInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceRequest) SetPrototypeAssistantId(v string) *ListInstanceRequest {
	s.PrototypeAssistantId = &v
	return s
}

type ListInstanceResponseBody struct {
	Result []*ListInstanceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBody) SetResult(v []*ListInstanceResponseBodyResult) *ListInstanceResponseBody {
	s.Result = v
	return s
}

type ListInstanceResponseBodyResult struct {
	Description          *string `json:"description,omitempty" xml:"description,omitempty"`
	Icon                 *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Name                 *string `json:"name,omitempty" xml:"name,omitempty"`
	PrototypeAssistantId *string `json:"prototypeAssistantId,omitempty" xml:"prototypeAssistantId,omitempty"`
	TenantAssistantId    *string `json:"tenantAssistantId,omitempty" xml:"tenantAssistantId,omitempty"`
}

func (s ListInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyResult) SetDescription(v string) *ListInstanceResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetIcon(v string) *ListInstanceResponseBodyResult {
	s.Icon = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetName(v string) *ListInstanceResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetPrototypeAssistantId(v string) *ListInstanceResponseBodyResult {
	s.PrototypeAssistantId = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetTenantAssistantId(v string) *ListInstanceResponseBodyResult {
	s.TenantAssistantId = &v
	return s
}

type ListInstanceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceResponse) SetHeaders(v map[string]*string) *ListInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceResponse) SetStatusCode(v int32) *ListInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceResponse) SetBody(v *ListInstanceResponseBody) *ListInstanceResponse {
	s.Body = v
	return s
}

type ListVisibleAssistantHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListVisibleAssistantHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListVisibleAssistantHeaders) GoString() string {
	return s.String()
}

func (s *ListVisibleAssistantHeaders) SetCommonHeaders(v map[string]*string) *ListVisibleAssistantHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListVisibleAssistantHeaders) SetXAcsDingtalkAccessToken(v string) *ListVisibleAssistantHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListVisibleAssistantRequest struct {
	Cursor *int64  `json:"cursor,omitempty" xml:"cursor,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s ListVisibleAssistantRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVisibleAssistantRequest) GoString() string {
	return s.String()
}

func (s *ListVisibleAssistantRequest) SetCursor(v int64) *ListVisibleAssistantRequest {
	s.Cursor = &v
	return s
}

func (s *ListVisibleAssistantRequest) SetName(v string) *ListVisibleAssistantRequest {
	s.Name = &v
	return s
}

func (s *ListVisibleAssistantRequest) SetPageSize(v int32) *ListVisibleAssistantRequest {
	s.PageSize = &v
	return s
}

func (s *ListVisibleAssistantRequest) SetUnionId(v string) *ListVisibleAssistantRequest {
	s.UnionId = &v
	return s
}

type ListVisibleAssistantResponseBody struct {
	HasMore    *bool                                   `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*ListVisibleAssistantResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextCursor *int64                                  `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	TotalCount *int32                                  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListVisibleAssistantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVisibleAssistantResponseBody) GoString() string {
	return s.String()
}

func (s *ListVisibleAssistantResponseBody) SetHasMore(v bool) *ListVisibleAssistantResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListVisibleAssistantResponseBody) SetList(v []*ListVisibleAssistantResponseBodyList) *ListVisibleAssistantResponseBody {
	s.List = v
	return s
}

func (s *ListVisibleAssistantResponseBody) SetNextCursor(v int64) *ListVisibleAssistantResponseBody {
	s.NextCursor = &v
	return s
}

func (s *ListVisibleAssistantResponseBody) SetTotalCount(v int32) *ListVisibleAssistantResponseBody {
	s.TotalCount = &v
	return s
}

type ListVisibleAssistantResponseBodyList struct {
	AssistantId    *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	CreatedAt      *int64  `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatorUnionId *string `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	Icon           *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListVisibleAssistantResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListVisibleAssistantResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListVisibleAssistantResponseBodyList) SetAssistantId(v string) *ListVisibleAssistantResponseBodyList {
	s.AssistantId = &v
	return s
}

func (s *ListVisibleAssistantResponseBodyList) SetCreatedAt(v int64) *ListVisibleAssistantResponseBodyList {
	s.CreatedAt = &v
	return s
}

func (s *ListVisibleAssistantResponseBodyList) SetCreatorUnionId(v string) *ListVisibleAssistantResponseBodyList {
	s.CreatorUnionId = &v
	return s
}

func (s *ListVisibleAssistantResponseBodyList) SetDescription(v string) *ListVisibleAssistantResponseBodyList {
	s.Description = &v
	return s
}

func (s *ListVisibleAssistantResponseBodyList) SetIcon(v string) *ListVisibleAssistantResponseBodyList {
	s.Icon = &v
	return s
}

func (s *ListVisibleAssistantResponseBodyList) SetName(v string) *ListVisibleAssistantResponseBodyList {
	s.Name = &v
	return s
}

type ListVisibleAssistantResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVisibleAssistantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVisibleAssistantResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVisibleAssistantResponse) GoString() string {
	return s.String()
}

func (s *ListVisibleAssistantResponse) SetHeaders(v map[string]*string) *ListVisibleAssistantResponse {
	s.Headers = v
	return s
}

func (s *ListVisibleAssistantResponse) SetStatusCode(v int32) *ListVisibleAssistantResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVisibleAssistantResponse) SetBody(v *ListVisibleAssistantResponseBody) *ListVisibleAssistantResponse {
	s.Body = v
	return s
}

type LogListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s LogListHeaders) String() string {
	return tea.Prettify(s)
}

func (s LogListHeaders) GoString() string {
	return s.String()
}

func (s *LogListHeaders) SetCommonHeaders(v map[string]*string) *LogListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *LogListHeaders) SetXAcsDingtalkAccessToken(v string) *LogListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type LogListRequest struct {
	// This parameter is required.
	AssistantId *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	EndTime     *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	PageNumber  *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize    *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StartTime   *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s LogListRequest) String() string {
	return tea.Prettify(s)
}

func (s LogListRequest) GoString() string {
	return s.String()
}

func (s *LogListRequest) SetAssistantId(v string) *LogListRequest {
	s.AssistantId = &v
	return s
}

func (s *LogListRequest) SetEndTime(v int64) *LogListRequest {
	s.EndTime = &v
	return s
}

func (s *LogListRequest) SetPageNumber(v int32) *LogListRequest {
	s.PageNumber = &v
	return s
}

func (s *LogListRequest) SetPageSize(v int32) *LogListRequest {
	s.PageSize = &v
	return s
}

func (s *LogListRequest) SetStartTime(v int64) *LogListRequest {
	s.StartTime = &v
	return s
}

type LogListResponseBody struct {
	Result  *LogListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                      `json:"success,omitempty" xml:"success,omitempty"`
}

func (s LogListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LogListResponseBody) GoString() string {
	return s.String()
}

func (s *LogListResponseBody) SetResult(v *LogListResponseBodyResult) *LogListResponseBody {
	s.Result = v
	return s
}

func (s *LogListResponseBody) SetSuccess(v bool) *LogListResponseBody {
	s.Success = &v
	return s
}

type LogListResponseBodyResult struct {
	HasMore    *bool                            `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*LogListResponseBodyResultList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	TotalCount *int32                           `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s LogListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s LogListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *LogListResponseBodyResult) SetHasMore(v bool) *LogListResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *LogListResponseBodyResult) SetList(v []*LogListResponseBodyResultList) *LogListResponseBodyResult {
	s.List = v
	return s
}

func (s *LogListResponseBodyResult) SetTotalCount(v int32) *LogListResponseBodyResult {
	s.TotalCount = &v
	return s
}

type LogListResponseBodyResultList struct {
	ActionNames   *string `json:"actionNames,omitempty" xml:"actionNames,omitempty"`
	CustomChannel *string `json:"customChannel,omitempty" xml:"customChannel,omitempty"`
	Input         *string `json:"input,omitempty" xml:"input,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	Output        *string `json:"output,omitempty" xml:"output,omitempty"`
	Result        *string `json:"result,omitempty" xml:"result,omitempty"`
	Scene         *string `json:"scene,omitempty" xml:"scene,omitempty"`
	Time          *int64  `json:"time,omitempty" xml:"time,omitempty"`
	UnionId       *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId        *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s LogListResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s LogListResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *LogListResponseBodyResultList) SetActionNames(v string) *LogListResponseBodyResultList {
	s.ActionNames = &v
	return s
}

func (s *LogListResponseBodyResultList) SetCustomChannel(v string) *LogListResponseBodyResultList {
	s.CustomChannel = &v
	return s
}

func (s *LogListResponseBodyResultList) SetInput(v string) *LogListResponseBodyResultList {
	s.Input = &v
	return s
}

func (s *LogListResponseBodyResultList) SetName(v string) *LogListResponseBodyResultList {
	s.Name = &v
	return s
}

func (s *LogListResponseBodyResultList) SetOutput(v string) *LogListResponseBodyResultList {
	s.Output = &v
	return s
}

func (s *LogListResponseBodyResultList) SetResult(v string) *LogListResponseBodyResultList {
	s.Result = &v
	return s
}

func (s *LogListResponseBodyResultList) SetScene(v string) *LogListResponseBodyResultList {
	s.Scene = &v
	return s
}

func (s *LogListResponseBodyResultList) SetTime(v int64) *LogListResponseBodyResultList {
	s.Time = &v
	return s
}

func (s *LogListResponseBodyResultList) SetUnionId(v string) *LogListResponseBodyResultList {
	s.UnionId = &v
	return s
}

func (s *LogListResponseBodyResultList) SetUserId(v string) *LogListResponseBodyResultList {
	s.UserId = &v
	return s
}

type LogListResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LogListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LogListResponse) String() string {
	return tea.Prettify(s)
}

func (s LogListResponse) GoString() string {
	return s.String()
}

func (s *LogListResponse) SetHeaders(v map[string]*string) *LogListResponse {
	s.Headers = v
	return s
}

func (s *LogListResponse) SetStatusCode(v int32) *LogListResponse {
	s.StatusCode = &v
	return s
}

func (s *LogListResponse) SetBody(v *LogListResponseBody) *LogListResponse {
	s.Body = v
	return s
}

type RelearnKnowledgeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RelearnKnowledgeHeaders) String() string {
	return tea.Prettify(s)
}

func (s RelearnKnowledgeHeaders) GoString() string {
	return s.String()
}

func (s *RelearnKnowledgeHeaders) SetCommonHeaders(v map[string]*string) *RelearnKnowledgeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RelearnKnowledgeHeaders) SetXAcsDingtalkAccessToken(v string) *RelearnKnowledgeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RelearnKnowledgeRequest struct {
	// This parameter is required.
	AssistantId *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
}

func (s RelearnKnowledgeRequest) String() string {
	return tea.Prettify(s)
}

func (s RelearnKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *RelearnKnowledgeRequest) SetAssistantId(v string) *RelearnKnowledgeRequest {
	s.AssistantId = &v
	return s
}

type RelearnKnowledgeResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RelearnKnowledgeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RelearnKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *RelearnKnowledgeResponseBody) SetSuccess(v bool) *RelearnKnowledgeResponseBody {
	s.Success = &v
	return s
}

type RelearnKnowledgeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RelearnKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RelearnKnowledgeResponse) String() string {
	return tea.Prettify(s)
}

func (s RelearnKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *RelearnKnowledgeResponse) SetHeaders(v map[string]*string) *RelearnKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *RelearnKnowledgeResponse) SetStatusCode(v int32) *RelearnKnowledgeResponse {
	s.StatusCode = &v
	return s
}

func (s *RelearnKnowledgeResponse) SetBody(v *RelearnKnowledgeResponseBody) *RelearnKnowledgeResponse {
	s.Body = v
	return s
}

type RemoveAssistantHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RemoveAssistantHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveAssistantHeaders) GoString() string {
	return s.String()
}

func (s *RemoveAssistantHeaders) SetCommonHeaders(v map[string]*string) *RemoveAssistantHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveAssistantHeaders) SetXAcsDingtalkAccessToken(v string) *RemoveAssistantHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RemoveAssistantRequest struct {
	AssistantId     *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	OperatorUnionId *string `json:"operatorUnionId,omitempty" xml:"operatorUnionId,omitempty"`
}

func (s RemoveAssistantRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveAssistantRequest) GoString() string {
	return s.String()
}

func (s *RemoveAssistantRequest) SetAssistantId(v string) *RemoveAssistantRequest {
	s.AssistantId = &v
	return s
}

func (s *RemoveAssistantRequest) SetOperatorUnionId(v string) *RemoveAssistantRequest {
	s.OperatorUnionId = &v
	return s
}

type RemoveAssistantResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RemoveAssistantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveAssistantResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAssistantResponseBody) SetSuccess(v bool) *RemoveAssistantResponseBody {
	s.Success = &v
	return s
}

type RemoveAssistantResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveAssistantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveAssistantResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveAssistantResponse) GoString() string {
	return s.String()
}

func (s *RemoveAssistantResponse) SetHeaders(v map[string]*string) *RemoveAssistantResponse {
	s.Headers = v
	return s
}

func (s *RemoveAssistantResponse) SetStatusCode(v int32) *RemoveAssistantResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveAssistantResponse) SetBody(v *RemoveAssistantResponseBody) *RemoveAssistantResponse {
	s.Body = v
	return s
}

type RemoveFromOrgSkillRepositoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RemoveFromOrgSkillRepositoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveFromOrgSkillRepositoryHeaders) GoString() string {
	return s.String()
}

func (s *RemoveFromOrgSkillRepositoryHeaders) SetCommonHeaders(v map[string]*string) *RemoveFromOrgSkillRepositoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveFromOrgSkillRepositoryHeaders) SetXAcsDingtalkAccessToken(v string) *RemoveFromOrgSkillRepositoryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RemoveFromOrgSkillRepositoryRequest struct {
	// This parameter is required.
	ActionId *string `json:"actionId,omitempty" xml:"actionId,omitempty"`
	// This parameter is required.
	OperatorUnionId *string `json:"operatorUnionId,omitempty" xml:"operatorUnionId,omitempty"`
}

func (s RemoveFromOrgSkillRepositoryRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveFromOrgSkillRepositoryRequest) GoString() string {
	return s.String()
}

func (s *RemoveFromOrgSkillRepositoryRequest) SetActionId(v string) *RemoveFromOrgSkillRepositoryRequest {
	s.ActionId = &v
	return s
}

func (s *RemoveFromOrgSkillRepositoryRequest) SetOperatorUnionId(v string) *RemoveFromOrgSkillRepositoryRequest {
	s.OperatorUnionId = &v
	return s
}

type RemoveFromOrgSkillRepositoryResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RemoveFromOrgSkillRepositoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveFromOrgSkillRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveFromOrgSkillRepositoryResponseBody) SetSuccess(v bool) *RemoveFromOrgSkillRepositoryResponseBody {
	s.Success = &v
	return s
}

type RemoveFromOrgSkillRepositoryResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveFromOrgSkillRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveFromOrgSkillRepositoryResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveFromOrgSkillRepositoryResponse) GoString() string {
	return s.String()
}

func (s *RemoveFromOrgSkillRepositoryResponse) SetHeaders(v map[string]*string) *RemoveFromOrgSkillRepositoryResponse {
	s.Headers = v
	return s
}

func (s *RemoveFromOrgSkillRepositoryResponse) SetStatusCode(v int32) *RemoveFromOrgSkillRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveFromOrgSkillRepositoryResponse) SetBody(v *RemoveFromOrgSkillRepositoryResponseBody) *RemoveFromOrgSkillRepositoryResponse {
	s.Body = v
	return s
}

type RetrieveAssistantBasicInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RetrieveAssistantBasicInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s RetrieveAssistantBasicInfoHeaders) GoString() string {
	return s.String()
}

func (s *RetrieveAssistantBasicInfoHeaders) SetCommonHeaders(v map[string]*string) *RetrieveAssistantBasicInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RetrieveAssistantBasicInfoHeaders) SetXAcsDingtalkAccessToken(v string) *RetrieveAssistantBasicInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RetrieveAssistantBasicInfoRequest struct {
	// This parameter is required.
	AssistantId *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s RetrieveAssistantBasicInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s RetrieveAssistantBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *RetrieveAssistantBasicInfoRequest) SetAssistantId(v string) *RetrieveAssistantBasicInfoRequest {
	s.AssistantId = &v
	return s
}

func (s *RetrieveAssistantBasicInfoRequest) SetUnionId(v string) *RetrieveAssistantBasicInfoRequest {
	s.UnionId = &v
	return s
}

type RetrieveAssistantBasicInfoResponseBody struct {
	ActionNames        []*string `json:"actionNames,omitempty" xml:"actionNames,omitempty" type:"Repeated"`
	AssistantId        *string   `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	AssistantUnionId   *string   `json:"assistantUnionId,omitempty" xml:"assistantUnionId,omitempty"`
	CreatedAt          *int64    `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatorUnionId     *string   `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	Description        *string   `json:"description,omitempty" xml:"description,omitempty"`
	FallbackContent    *string   `json:"fallbackContent,omitempty" xml:"fallbackContent,omitempty"`
	Icon               *string   `json:"icon,omitempty" xml:"icon,omitempty"`
	Instructions       *string   `json:"instructions,omitempty" xml:"instructions,omitempty"`
	KnowledgeFileNames []*string `json:"knowledgeFileNames,omitempty" xml:"knowledgeFileNames,omitempty" type:"Repeated"`
	Model              *string   `json:"model,omitempty" xml:"model,omitempty"`
	Name               *string   `json:"name,omitempty" xml:"name,omitempty"`
	RecommendPrompts   []*string `json:"recommendPrompts,omitempty" xml:"recommendPrompts,omitempty" type:"Repeated"`
	UnifiedAppId       *string   `json:"unifiedAppId,omitempty" xml:"unifiedAppId,omitempty"`
	WelcomeContent     *string   `json:"welcomeContent,omitempty" xml:"welcomeContent,omitempty"`
}

func (s RetrieveAssistantBasicInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetrieveAssistantBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *RetrieveAssistantBasicInfoResponseBody) SetActionNames(v []*string) *RetrieveAssistantBasicInfoResponseBody {
	s.ActionNames = v
	return s
}

func (s *RetrieveAssistantBasicInfoResponseBody) SetAssistantId(v string) *RetrieveAssistantBasicInfoResponseBody {
	s.AssistantId = &v
	return s
}

func (s *RetrieveAssistantBasicInfoResponseBody) SetAssistantUnionId(v string) *RetrieveAssistantBasicInfoResponseBody {
	s.AssistantUnionId = &v
	return s
}

func (s *RetrieveAssistantBasicInfoResponseBody) SetCreatedAt(v int64) *RetrieveAssistantBasicInfoResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *RetrieveAssistantBasicInfoResponseBody) SetCreatorUnionId(v string) *RetrieveAssistantBasicInfoResponseBody {
	s.CreatorUnionId = &v
	return s
}

func (s *RetrieveAssistantBasicInfoResponseBody) SetDescription(v string) *RetrieveAssistantBasicInfoResponseBody {
	s.Description = &v
	return s
}

func (s *RetrieveAssistantBasicInfoResponseBody) SetFallbackContent(v string) *RetrieveAssistantBasicInfoResponseBody {
	s.FallbackContent = &v
	return s
}

func (s *RetrieveAssistantBasicInfoResponseBody) SetIcon(v string) *RetrieveAssistantBasicInfoResponseBody {
	s.Icon = &v
	return s
}

func (s *RetrieveAssistantBasicInfoResponseBody) SetInstructions(v string) *RetrieveAssistantBasicInfoResponseBody {
	s.Instructions = &v
	return s
}

func (s *RetrieveAssistantBasicInfoResponseBody) SetKnowledgeFileNames(v []*string) *RetrieveAssistantBasicInfoResponseBody {
	s.KnowledgeFileNames = v
	return s
}

func (s *RetrieveAssistantBasicInfoResponseBody) SetModel(v string) *RetrieveAssistantBasicInfoResponseBody {
	s.Model = &v
	return s
}

func (s *RetrieveAssistantBasicInfoResponseBody) SetName(v string) *RetrieveAssistantBasicInfoResponseBody {
	s.Name = &v
	return s
}

func (s *RetrieveAssistantBasicInfoResponseBody) SetRecommendPrompts(v []*string) *RetrieveAssistantBasicInfoResponseBody {
	s.RecommendPrompts = v
	return s
}

func (s *RetrieveAssistantBasicInfoResponseBody) SetUnifiedAppId(v string) *RetrieveAssistantBasicInfoResponseBody {
	s.UnifiedAppId = &v
	return s
}

func (s *RetrieveAssistantBasicInfoResponseBody) SetWelcomeContent(v string) *RetrieveAssistantBasicInfoResponseBody {
	s.WelcomeContent = &v
	return s
}

type RetrieveAssistantBasicInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetrieveAssistantBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetrieveAssistantBasicInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s RetrieveAssistantBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *RetrieveAssistantBasicInfoResponse) SetHeaders(v map[string]*string) *RetrieveAssistantBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *RetrieveAssistantBasicInfoResponse) SetStatusCode(v int32) *RetrieveAssistantBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *RetrieveAssistantBasicInfoResponse) SetBody(v *RetrieveAssistantBasicInfoResponseBody) *RetrieveAssistantBasicInfoResponse {
	s.Body = v
	return s
}

type RetrieveAssistantMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RetrieveAssistantMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s RetrieveAssistantMessageHeaders) GoString() string {
	return s.String()
}

func (s *RetrieveAssistantMessageHeaders) SetCommonHeaders(v map[string]*string) *RetrieveAssistantMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RetrieveAssistantMessageHeaders) SetXAcsDingtalkAccessToken(v string) *RetrieveAssistantMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RetrieveAssistantMessageResponseBody struct {
	AssisantId *string                `json:"assisantId,omitempty" xml:"assisantId,omitempty"`
	Content    []interface{}          `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	CreatedAt  *int64                 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Id         *string                `json:"id,omitempty" xml:"id,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	Object     *string                `json:"object,omitempty" xml:"object,omitempty"`
	Role       *string                `json:"role,omitempty" xml:"role,omitempty"`
	RunId      *string                `json:"runId,omitempty" xml:"runId,omitempty"`
	ThreadId   *string                `json:"threadId,omitempty" xml:"threadId,omitempty"`
}

func (s RetrieveAssistantMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetrieveAssistantMessageResponseBody) GoString() string {
	return s.String()
}

func (s *RetrieveAssistantMessageResponseBody) SetAssisantId(v string) *RetrieveAssistantMessageResponseBody {
	s.AssisantId = &v
	return s
}

func (s *RetrieveAssistantMessageResponseBody) SetContent(v []interface{}) *RetrieveAssistantMessageResponseBody {
	s.Content = v
	return s
}

func (s *RetrieveAssistantMessageResponseBody) SetCreatedAt(v int64) *RetrieveAssistantMessageResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *RetrieveAssistantMessageResponseBody) SetId(v string) *RetrieveAssistantMessageResponseBody {
	s.Id = &v
	return s
}

func (s *RetrieveAssistantMessageResponseBody) SetMetadata(v map[string]interface{}) *RetrieveAssistantMessageResponseBody {
	s.Metadata = v
	return s
}

func (s *RetrieveAssistantMessageResponseBody) SetObject(v string) *RetrieveAssistantMessageResponseBody {
	s.Object = &v
	return s
}

func (s *RetrieveAssistantMessageResponseBody) SetRole(v string) *RetrieveAssistantMessageResponseBody {
	s.Role = &v
	return s
}

func (s *RetrieveAssistantMessageResponseBody) SetRunId(v string) *RetrieveAssistantMessageResponseBody {
	s.RunId = &v
	return s
}

func (s *RetrieveAssistantMessageResponseBody) SetThreadId(v string) *RetrieveAssistantMessageResponseBody {
	s.ThreadId = &v
	return s
}

type RetrieveAssistantMessageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetrieveAssistantMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetrieveAssistantMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s RetrieveAssistantMessageResponse) GoString() string {
	return s.String()
}

func (s *RetrieveAssistantMessageResponse) SetHeaders(v map[string]*string) *RetrieveAssistantMessageResponse {
	s.Headers = v
	return s
}

func (s *RetrieveAssistantMessageResponse) SetStatusCode(v int32) *RetrieveAssistantMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *RetrieveAssistantMessageResponse) SetBody(v *RetrieveAssistantMessageResponseBody) *RetrieveAssistantMessageResponse {
	s.Body = v
	return s
}

type RetrieveAssistantRunHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RetrieveAssistantRunHeaders) String() string {
	return tea.Prettify(s)
}

func (s RetrieveAssistantRunHeaders) GoString() string {
	return s.String()
}

func (s *RetrieveAssistantRunHeaders) SetCommonHeaders(v map[string]*string) *RetrieveAssistantRunHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RetrieveAssistantRunHeaders) SetXAcsDingtalkAccessToken(v string) *RetrieveAssistantRunHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RetrieveAssistantRunResponseBody struct {
	AssistantId  *string                `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	CancelledAt  *int64                 `json:"cancelledAt,omitempty" xml:"cancelledAt,omitempty"`
	CompletedAt  *int64                 `json:"completedAt,omitempty" xml:"completedAt,omitempty"`
	CreatedAt    *int64                 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	ExpiresAt    *int64                 `json:"expiresAt,omitempty" xml:"expiresAt,omitempty"`
	FailedAt     *int64                 `json:"failedAt,omitempty" xml:"failedAt,omitempty"`
	Id           *string                `json:"id,omitempty" xml:"id,omitempty"`
	LastErrorMsg *string                `json:"lastErrorMsg,omitempty" xml:"lastErrorMsg,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	Object       *string                `json:"object,omitempty" xml:"object,omitempty"`
	StartedAt    *int64                 `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	Status       *string                `json:"status,omitempty" xml:"status,omitempty"`
	ThreadId     *string                `json:"threadId,omitempty" xml:"threadId,omitempty"`
}

func (s RetrieveAssistantRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetrieveAssistantRunResponseBody) GoString() string {
	return s.String()
}

func (s *RetrieveAssistantRunResponseBody) SetAssistantId(v string) *RetrieveAssistantRunResponseBody {
	s.AssistantId = &v
	return s
}

func (s *RetrieveAssistantRunResponseBody) SetCancelledAt(v int64) *RetrieveAssistantRunResponseBody {
	s.CancelledAt = &v
	return s
}

func (s *RetrieveAssistantRunResponseBody) SetCompletedAt(v int64) *RetrieveAssistantRunResponseBody {
	s.CompletedAt = &v
	return s
}

func (s *RetrieveAssistantRunResponseBody) SetCreatedAt(v int64) *RetrieveAssistantRunResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *RetrieveAssistantRunResponseBody) SetExpiresAt(v int64) *RetrieveAssistantRunResponseBody {
	s.ExpiresAt = &v
	return s
}

func (s *RetrieveAssistantRunResponseBody) SetFailedAt(v int64) *RetrieveAssistantRunResponseBody {
	s.FailedAt = &v
	return s
}

func (s *RetrieveAssistantRunResponseBody) SetId(v string) *RetrieveAssistantRunResponseBody {
	s.Id = &v
	return s
}

func (s *RetrieveAssistantRunResponseBody) SetLastErrorMsg(v string) *RetrieveAssistantRunResponseBody {
	s.LastErrorMsg = &v
	return s
}

func (s *RetrieveAssistantRunResponseBody) SetMetadata(v map[string]interface{}) *RetrieveAssistantRunResponseBody {
	s.Metadata = v
	return s
}

func (s *RetrieveAssistantRunResponseBody) SetObject(v string) *RetrieveAssistantRunResponseBody {
	s.Object = &v
	return s
}

func (s *RetrieveAssistantRunResponseBody) SetStartedAt(v int64) *RetrieveAssistantRunResponseBody {
	s.StartedAt = &v
	return s
}

func (s *RetrieveAssistantRunResponseBody) SetStatus(v string) *RetrieveAssistantRunResponseBody {
	s.Status = &v
	return s
}

func (s *RetrieveAssistantRunResponseBody) SetThreadId(v string) *RetrieveAssistantRunResponseBody {
	s.ThreadId = &v
	return s
}

type RetrieveAssistantRunResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetrieveAssistantRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetrieveAssistantRunResponse) String() string {
	return tea.Prettify(s)
}

func (s RetrieveAssistantRunResponse) GoString() string {
	return s.String()
}

func (s *RetrieveAssistantRunResponse) SetHeaders(v map[string]*string) *RetrieveAssistantRunResponse {
	s.Headers = v
	return s
}

func (s *RetrieveAssistantRunResponse) SetStatusCode(v int32) *RetrieveAssistantRunResponse {
	s.StatusCode = &v
	return s
}

func (s *RetrieveAssistantRunResponse) SetBody(v *RetrieveAssistantRunResponseBody) *RetrieveAssistantRunResponse {
	s.Body = v
	return s
}

type RetrieveAssistantScopeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RetrieveAssistantScopeHeaders) String() string {
	return tea.Prettify(s)
}

func (s RetrieveAssistantScopeHeaders) GoString() string {
	return s.String()
}

func (s *RetrieveAssistantScopeHeaders) SetCommonHeaders(v map[string]*string) *RetrieveAssistantScopeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RetrieveAssistantScopeHeaders) SetXAcsDingtalkAccessToken(v string) *RetrieveAssistantScopeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RetrieveAssistantScopeRequest struct {
	// This parameter is required.
	AssistantId *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
}

func (s RetrieveAssistantScopeRequest) String() string {
	return tea.Prettify(s)
}

func (s RetrieveAssistantScopeRequest) GoString() string {
	return s.String()
}

func (s *RetrieveAssistantScopeRequest) SetAssistantId(v string) *RetrieveAssistantScopeRequest {
	s.AssistantId = &v
	return s
}

type RetrieveAssistantScopeResponseBody struct {
	Result  *RetrieveAssistantScopeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RetrieveAssistantScopeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetrieveAssistantScopeResponseBody) GoString() string {
	return s.String()
}

func (s *RetrieveAssistantScopeResponseBody) SetResult(v *RetrieveAssistantScopeResponseBodyResult) *RetrieveAssistantScopeResponseBody {
	s.Result = v
	return s
}

func (s *RetrieveAssistantScopeResponseBody) SetSuccess(v bool) *RetrieveAssistantScopeResponseBody {
	s.Success = &v
	return s
}

type RetrieveAssistantScopeResponseBodyResult struct {
	AssistantId *string                                         `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	Scopes      *RetrieveAssistantScopeResponseBodyResultScopes `json:"scopes,omitempty" xml:"scopes,omitempty" type:"Struct"`
	Sharing     *bool                                           `json:"sharing,omitempty" xml:"sharing,omitempty"`
}

func (s RetrieveAssistantScopeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s RetrieveAssistantScopeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RetrieveAssistantScopeResponseBodyResult) SetAssistantId(v string) *RetrieveAssistantScopeResponseBodyResult {
	s.AssistantId = &v
	return s
}

func (s *RetrieveAssistantScopeResponseBodyResult) SetScopes(v *RetrieveAssistantScopeResponseBodyResultScopes) *RetrieveAssistantScopeResponseBodyResult {
	s.Scopes = v
	return s
}

func (s *RetrieveAssistantScopeResponseBodyResult) SetSharing(v bool) *RetrieveAssistantScopeResponseBodyResult {
	s.Sharing = &v
	return s
}

type RetrieveAssistantScopeResponseBodyResultScopes struct {
	DeptVisibleScopes  []*string `json:"deptVisibleScopes,omitempty" xml:"deptVisibleScopes,omitempty" type:"Repeated"`
	DynamicGroupScopes []*string `json:"dynamicGroupScopes,omitempty" xml:"dynamicGroupScopes,omitempty" type:"Repeated"`
	IsAdmin            *bool     `json:"isAdmin,omitempty" xml:"isAdmin,omitempty"`
	RoleVisibleScopes  []*string `json:"roleVisibleScopes,omitempty" xml:"roleVisibleScopes,omitempty" type:"Repeated"`
	UserVisibleScopes  []*string `json:"userVisibleScopes,omitempty" xml:"userVisibleScopes,omitempty" type:"Repeated"`
}

func (s RetrieveAssistantScopeResponseBodyResultScopes) String() string {
	return tea.Prettify(s)
}

func (s RetrieveAssistantScopeResponseBodyResultScopes) GoString() string {
	return s.String()
}

func (s *RetrieveAssistantScopeResponseBodyResultScopes) SetDeptVisibleScopes(v []*string) *RetrieveAssistantScopeResponseBodyResultScopes {
	s.DeptVisibleScopes = v
	return s
}

func (s *RetrieveAssistantScopeResponseBodyResultScopes) SetDynamicGroupScopes(v []*string) *RetrieveAssistantScopeResponseBodyResultScopes {
	s.DynamicGroupScopes = v
	return s
}

func (s *RetrieveAssistantScopeResponseBodyResultScopes) SetIsAdmin(v bool) *RetrieveAssistantScopeResponseBodyResultScopes {
	s.IsAdmin = &v
	return s
}

func (s *RetrieveAssistantScopeResponseBodyResultScopes) SetRoleVisibleScopes(v []*string) *RetrieveAssistantScopeResponseBodyResultScopes {
	s.RoleVisibleScopes = v
	return s
}

func (s *RetrieveAssistantScopeResponseBodyResultScopes) SetUserVisibleScopes(v []*string) *RetrieveAssistantScopeResponseBodyResultScopes {
	s.UserVisibleScopes = v
	return s
}

type RetrieveAssistantScopeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetrieveAssistantScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetrieveAssistantScopeResponse) String() string {
	return tea.Prettify(s)
}

func (s RetrieveAssistantScopeResponse) GoString() string {
	return s.String()
}

func (s *RetrieveAssistantScopeResponse) SetHeaders(v map[string]*string) *RetrieveAssistantScopeResponse {
	s.Headers = v
	return s
}

func (s *RetrieveAssistantScopeResponse) SetStatusCode(v int32) *RetrieveAssistantScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *RetrieveAssistantScopeResponse) SetBody(v *RetrieveAssistantScopeResponseBody) *RetrieveAssistantScopeResponse {
	s.Body = v
	return s
}

type RetrieveAssistantThreadHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RetrieveAssistantThreadHeaders) String() string {
	return tea.Prettify(s)
}

func (s RetrieveAssistantThreadHeaders) GoString() string {
	return s.String()
}

func (s *RetrieveAssistantThreadHeaders) SetCommonHeaders(v map[string]*string) *RetrieveAssistantThreadHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RetrieveAssistantThreadHeaders) SetXAcsDingtalkAccessToken(v string) *RetrieveAssistantThreadHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RetrieveAssistantThreadResponseBody struct {
	CreatedAt *int64                 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Id        *string                `json:"id,omitempty" xml:"id,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	Object    *string                `json:"object,omitempty" xml:"object,omitempty"`
}

func (s RetrieveAssistantThreadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetrieveAssistantThreadResponseBody) GoString() string {
	return s.String()
}

func (s *RetrieveAssistantThreadResponseBody) SetCreatedAt(v int64) *RetrieveAssistantThreadResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *RetrieveAssistantThreadResponseBody) SetId(v string) *RetrieveAssistantThreadResponseBody {
	s.Id = &v
	return s
}

func (s *RetrieveAssistantThreadResponseBody) SetMetadata(v map[string]interface{}) *RetrieveAssistantThreadResponseBody {
	s.Metadata = v
	return s
}

func (s *RetrieveAssistantThreadResponseBody) SetObject(v string) *RetrieveAssistantThreadResponseBody {
	s.Object = &v
	return s
}

type RetrieveAssistantThreadResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetrieveAssistantThreadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetrieveAssistantThreadResponse) String() string {
	return tea.Prettify(s)
}

func (s RetrieveAssistantThreadResponse) GoString() string {
	return s.String()
}

func (s *RetrieveAssistantThreadResponse) SetHeaders(v map[string]*string) *RetrieveAssistantThreadResponse {
	s.Headers = v
	return s
}

func (s *RetrieveAssistantThreadResponse) SetStatusCode(v int32) *RetrieveAssistantThreadResponse {
	s.StatusCode = &v
	return s
}

func (s *RetrieveAssistantThreadResponse) SetBody(v *RetrieveAssistantThreadResponseBody) *RetrieveAssistantThreadResponse {
	s.Body = v
	return s
}

type UpdateAssistantBasicInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateAssistantBasicInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateAssistantBasicInfoHeaders) GoString() string {
	return s.String()
}

func (s *UpdateAssistantBasicInfoHeaders) SetCommonHeaders(v map[string]*string) *UpdateAssistantBasicInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateAssistantBasicInfoHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateAssistantBasicInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateAssistantBasicInfoRequest struct {
	// This parameter is required.
	AssistantId     *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	FallbackContent *string `json:"fallbackContent,omitempty" xml:"fallbackContent,omitempty"`
	Icon            *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Instructions    *string `json:"instructions,omitempty" xml:"instructions,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	OperatorUnionId  *string   `json:"operatorUnionId,omitempty" xml:"operatorUnionId,omitempty"`
	RecommendPrompts []*string `json:"recommendPrompts,omitempty" xml:"recommendPrompts,omitempty" type:"Repeated"`
	WelcomeContent   *string   `json:"welcomeContent,omitempty" xml:"welcomeContent,omitempty"`
}

func (s UpdateAssistantBasicInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAssistantBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateAssistantBasicInfoRequest) SetAssistantId(v string) *UpdateAssistantBasicInfoRequest {
	s.AssistantId = &v
	return s
}

func (s *UpdateAssistantBasicInfoRequest) SetDescription(v string) *UpdateAssistantBasicInfoRequest {
	s.Description = &v
	return s
}

func (s *UpdateAssistantBasicInfoRequest) SetFallbackContent(v string) *UpdateAssistantBasicInfoRequest {
	s.FallbackContent = &v
	return s
}

func (s *UpdateAssistantBasicInfoRequest) SetIcon(v string) *UpdateAssistantBasicInfoRequest {
	s.Icon = &v
	return s
}

func (s *UpdateAssistantBasicInfoRequest) SetInstructions(v string) *UpdateAssistantBasicInfoRequest {
	s.Instructions = &v
	return s
}

func (s *UpdateAssistantBasicInfoRequest) SetName(v string) *UpdateAssistantBasicInfoRequest {
	s.Name = &v
	return s
}

func (s *UpdateAssistantBasicInfoRequest) SetOperatorUnionId(v string) *UpdateAssistantBasicInfoRequest {
	s.OperatorUnionId = &v
	return s
}

func (s *UpdateAssistantBasicInfoRequest) SetRecommendPrompts(v []*string) *UpdateAssistantBasicInfoRequest {
	s.RecommendPrompts = v
	return s
}

func (s *UpdateAssistantBasicInfoRequest) SetWelcomeContent(v string) *UpdateAssistantBasicInfoRequest {
	s.WelcomeContent = &v
	return s
}

type UpdateAssistantBasicInfoResponseBody struct {
	ActionNames        []*string `json:"actionNames,omitempty" xml:"actionNames,omitempty" type:"Repeated"`
	AssistantId        *string   `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	AssistantUnionId   *string   `json:"assistantUnionId,omitempty" xml:"assistantUnionId,omitempty"`
	CreatedAt          *int64    `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatorUnionId     *string   `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	Description        *string   `json:"description,omitempty" xml:"description,omitempty"`
	FallbackContent    *string   `json:"fallbackContent,omitempty" xml:"fallbackContent,omitempty"`
	Icon               *string   `json:"icon,omitempty" xml:"icon,omitempty"`
	Instructions       *string   `json:"instructions,omitempty" xml:"instructions,omitempty"`
	KnowledgeFileNames []*string `json:"knowledgeFileNames,omitempty" xml:"knowledgeFileNames,omitempty" type:"Repeated"`
	Model              *string   `json:"model,omitempty" xml:"model,omitempty"`
	Name               *string   `json:"name,omitempty" xml:"name,omitempty"`
	RecommendPrompts   []*string `json:"recommendPrompts,omitempty" xml:"recommendPrompts,omitempty" type:"Repeated"`
	UnifiedAppId       *string   `json:"unifiedAppId,omitempty" xml:"unifiedAppId,omitempty"`
	WelcomeContent     *string   `json:"welcomeContent,omitempty" xml:"welcomeContent,omitempty"`
}

func (s UpdateAssistantBasicInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAssistantBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAssistantBasicInfoResponseBody) SetActionNames(v []*string) *UpdateAssistantBasicInfoResponseBody {
	s.ActionNames = v
	return s
}

func (s *UpdateAssistantBasicInfoResponseBody) SetAssistantId(v string) *UpdateAssistantBasicInfoResponseBody {
	s.AssistantId = &v
	return s
}

func (s *UpdateAssistantBasicInfoResponseBody) SetAssistantUnionId(v string) *UpdateAssistantBasicInfoResponseBody {
	s.AssistantUnionId = &v
	return s
}

func (s *UpdateAssistantBasicInfoResponseBody) SetCreatedAt(v int64) *UpdateAssistantBasicInfoResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *UpdateAssistantBasicInfoResponseBody) SetCreatorUnionId(v string) *UpdateAssistantBasicInfoResponseBody {
	s.CreatorUnionId = &v
	return s
}

func (s *UpdateAssistantBasicInfoResponseBody) SetDescription(v string) *UpdateAssistantBasicInfoResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateAssistantBasicInfoResponseBody) SetFallbackContent(v string) *UpdateAssistantBasicInfoResponseBody {
	s.FallbackContent = &v
	return s
}

func (s *UpdateAssistantBasicInfoResponseBody) SetIcon(v string) *UpdateAssistantBasicInfoResponseBody {
	s.Icon = &v
	return s
}

func (s *UpdateAssistantBasicInfoResponseBody) SetInstructions(v string) *UpdateAssistantBasicInfoResponseBody {
	s.Instructions = &v
	return s
}

func (s *UpdateAssistantBasicInfoResponseBody) SetKnowledgeFileNames(v []*string) *UpdateAssistantBasicInfoResponseBody {
	s.KnowledgeFileNames = v
	return s
}

func (s *UpdateAssistantBasicInfoResponseBody) SetModel(v string) *UpdateAssistantBasicInfoResponseBody {
	s.Model = &v
	return s
}

func (s *UpdateAssistantBasicInfoResponseBody) SetName(v string) *UpdateAssistantBasicInfoResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateAssistantBasicInfoResponseBody) SetRecommendPrompts(v []*string) *UpdateAssistantBasicInfoResponseBody {
	s.RecommendPrompts = v
	return s
}

func (s *UpdateAssistantBasicInfoResponseBody) SetUnifiedAppId(v string) *UpdateAssistantBasicInfoResponseBody {
	s.UnifiedAppId = &v
	return s
}

func (s *UpdateAssistantBasicInfoResponseBody) SetWelcomeContent(v string) *UpdateAssistantBasicInfoResponseBody {
	s.WelcomeContent = &v
	return s
}

type UpdateAssistantBasicInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAssistantBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAssistantBasicInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAssistantBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateAssistantBasicInfoResponse) SetHeaders(v map[string]*string) *UpdateAssistantBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateAssistantBasicInfoResponse) SetStatusCode(v int32) *UpdateAssistantBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAssistantBasicInfoResponse) SetBody(v *UpdateAssistantBasicInfoResponseBody) *UpdateAssistantBasicInfoResponse {
	s.Body = v
	return s
}

type UpdateAssistantScopeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateAssistantScopeHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateAssistantScopeHeaders) GoString() string {
	return s.String()
}

func (s *UpdateAssistantScopeHeaders) SetCommonHeaders(v map[string]*string) *UpdateAssistantScopeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateAssistantScopeHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateAssistantScopeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateAssistantScopeRequest struct {
	// This parameter is required.
	AssistantId *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	// This parameter is required.
	OperatorUnionId *string                            `json:"operatorUnionId,omitempty" xml:"operatorUnionId,omitempty"`
	Scopes          *UpdateAssistantScopeRequestScopes `json:"scopes,omitempty" xml:"scopes,omitempty" type:"Struct"`
	Sharing         *bool                              `json:"sharing,omitempty" xml:"sharing,omitempty"`
}

func (s UpdateAssistantScopeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAssistantScopeRequest) GoString() string {
	return s.String()
}

func (s *UpdateAssistantScopeRequest) SetAssistantId(v string) *UpdateAssistantScopeRequest {
	s.AssistantId = &v
	return s
}

func (s *UpdateAssistantScopeRequest) SetOperatorUnionId(v string) *UpdateAssistantScopeRequest {
	s.OperatorUnionId = &v
	return s
}

func (s *UpdateAssistantScopeRequest) SetScopes(v *UpdateAssistantScopeRequestScopes) *UpdateAssistantScopeRequest {
	s.Scopes = v
	return s
}

func (s *UpdateAssistantScopeRequest) SetSharing(v bool) *UpdateAssistantScopeRequest {
	s.Sharing = &v
	return s
}

type UpdateAssistantScopeRequestScopes struct {
	DeptVisibleScopes  []*string `json:"deptVisibleScopes,omitempty" xml:"deptVisibleScopes,omitempty" type:"Repeated"`
	DynamicGroupScopes []*string `json:"dynamicGroupScopes,omitempty" xml:"dynamicGroupScopes,omitempty" type:"Repeated"`
	IsAdmin            *bool     `json:"isAdmin,omitempty" xml:"isAdmin,omitempty"`
	RoleVisibleScopes  []*string `json:"roleVisibleScopes,omitempty" xml:"roleVisibleScopes,omitempty" type:"Repeated"`
	UserVisibleScopes  []*string `json:"userVisibleScopes,omitempty" xml:"userVisibleScopes,omitempty" type:"Repeated"`
}

func (s UpdateAssistantScopeRequestScopes) String() string {
	return tea.Prettify(s)
}

func (s UpdateAssistantScopeRequestScopes) GoString() string {
	return s.String()
}

func (s *UpdateAssistantScopeRequestScopes) SetDeptVisibleScopes(v []*string) *UpdateAssistantScopeRequestScopes {
	s.DeptVisibleScopes = v
	return s
}

func (s *UpdateAssistantScopeRequestScopes) SetDynamicGroupScopes(v []*string) *UpdateAssistantScopeRequestScopes {
	s.DynamicGroupScopes = v
	return s
}

func (s *UpdateAssistantScopeRequestScopes) SetIsAdmin(v bool) *UpdateAssistantScopeRequestScopes {
	s.IsAdmin = &v
	return s
}

func (s *UpdateAssistantScopeRequestScopes) SetRoleVisibleScopes(v []*string) *UpdateAssistantScopeRequestScopes {
	s.RoleVisibleScopes = v
	return s
}

func (s *UpdateAssistantScopeRequestScopes) SetUserVisibleScopes(v []*string) *UpdateAssistantScopeRequestScopes {
	s.UserVisibleScopes = v
	return s
}

type UpdateAssistantScopeResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       interface{}        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAssistantScopeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAssistantScopeResponse) GoString() string {
	return s.String()
}

func (s *UpdateAssistantScopeResponse) SetHeaders(v map[string]*string) *UpdateAssistantScopeResponse {
	s.Headers = v
	return s
}

func (s *UpdateAssistantScopeResponse) SetStatusCode(v int32) *UpdateAssistantScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAssistantScopeResponse) SetBody(v interface{}) *UpdateAssistantScopeResponse {
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
// 助理添加专业词汇
//
// @param request - AddDomainWordsRequest
//
// @param headers - AddDomainWordsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDomainWordsResponse
func (client *Client) AddDomainWordsWithOptions(request *AddDomainWordsRequest, headers *AddDomainWordsHeaders, runtime *util.RuntimeOptions) (_result *AddDomainWordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssistantId)) {
		body["assistantId"] = request.AssistantId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainWords)) {
		body["domainWords"] = request.DomainWords
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
		Action:      tea.String("AddDomainWords"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/domainWords"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddDomainWordsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 助理添加专业词汇
//
// @param request - AddDomainWordsRequest
//
// @return AddDomainWordsResponse
func (client *Client) AddDomainWords(request *AddDomainWordsRequest) (_result *AddDomainWordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddDomainWordsHeaders{}
	_result = &AddDomainWordsResponse{}
	_body, _err := client.AddDomainWordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加技能到组织技能库
//
// @param request - AddToOrgSkillRepositoryRequest
//
// @param headers - AddToOrgSkillRepositoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddToOrgSkillRepositoryResponse
func (client *Client) AddToOrgSkillRepositoryWithOptions(request *AddToOrgSkillRepositoryRequest, headers *AddToOrgSkillRepositoryHeaders, runtime *util.RuntimeOptions) (_result *AddToOrgSkillRepositoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionId)) {
		body["actionId"] = request.ActionId
	}

	if !tea.BoolValue(util.IsUnset(request.ActionVersion)) {
		body["actionVersion"] = request.ActionVersion
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
		Action:      tea.String("AddToOrgSkillRepository"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/orgActionRepositories"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddToOrgSkillRepositoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加技能到组织技能库
//
// @param request - AddToOrgSkillRepositoryRequest
//
// @return AddToOrgSkillRepositoryResponse
func (client *Client) AddToOrgSkillRepository(request *AddToOrgSkillRepositoryRequest) (_result *AddToOrgSkillRepositoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddToOrgSkillRepositoryHeaders{}
	_result = &AddToOrgSkillRepositoryResponse{}
	_body, _err := client.AddToOrgSkillRepositoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 助理响应接口-委托权限
//
// @param request - AssistantMeResponseRequest
//
// @param headers - AssistantMeResponseHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssistantMeResponseResponse
func (client *Client) AssistantMeResponseWithOptions(assistantId *string, request *AssistantMeResponseRequest, headers *AssistantMeResponseHeaders, runtime *util.RuntimeOptions) (_result *AssistantMeResponseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Input)) {
		body["input"] = request.Input
	}

	if !tea.BoolValue(util.IsUnset(request.Instructions)) {
		body["instructions"] = request.Instructions
	}

	if !tea.BoolValue(util.IsUnset(request.Metadata)) {
		body["metadata"] = request.Metadata
	}

	if !tea.BoolValue(util.IsUnset(request.Stream)) {
		body["stream"] = request.Stream
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
		Action:      tea.String("AssistantMeResponse"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/" + tea.StringValue(assistantId) + "/me/compatible-mode/responses"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AssistantMeResponseResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 助理响应接口-委托权限
//
// @param request - AssistantMeResponseRequest
//
// @return AssistantMeResponseResponse
func (client *Client) AssistantMeResponse(assistantId *string, request *AssistantMeResponseRequest) (_result *AssistantMeResponseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AssistantMeResponseHeaders{}
	_result = &AssistantMeResponseResponse{}
	_body, _err := client.AssistantMeResponseWithOptions(assistantId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 助理响应接口-应用权限
//
// @param request - AssistantResponseRequest
//
// @param headers - AssistantResponseHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssistantResponseResponse
func (client *Client) AssistantResponseWithOptions(assistantId *string, request *AssistantResponseRequest, headers *AssistantResponseHeaders, runtime *util.RuntimeOptions) (_result *AssistantResponseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Input)) {
		body["input"] = request.Input
	}

	if !tea.BoolValue(util.IsUnset(request.Instructions)) {
		body["instructions"] = request.Instructions
	}

	if !tea.BoolValue(util.IsUnset(request.Metadata)) {
		body["metadata"] = request.Metadata
	}

	if !tea.BoolValue(util.IsUnset(request.Stream)) {
		body["stream"] = request.Stream
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
		Action:      tea.String("AssistantResponse"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/" + tea.StringValue(assistantId) + "/compatible-mode/responses"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AssistantResponseResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 助理响应接口-应用权限
//
// @param request - AssistantResponseRequest
//
// @return AssistantResponseResponse
func (client *Client) AssistantResponse(assistantId *string, request *AssistantResponseRequest) (_result *AssistantResponseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AssistantResponseHeaders{}
	_result = &AssistantResponseResponse{}
	_body, _err := client.AssistantResponseWithOptions(assistantId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询算粒的使用记录
//
// @param request - BatchGetAICreditsRecordRequest
//
// @param headers - BatchGetAICreditsRecordHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetAICreditsRecordResponse
func (client *Client) BatchGetAICreditsRecordWithOptions(request *BatchGetAICreditsRecordRequest, headers *BatchGetAICreditsRecordHeaders, runtime *util.RuntimeOptions) (_result *BatchGetAICreditsRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssistantId)) {
		query["assistantId"] = request.AssistantId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
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
		Action:      tea.String("BatchGetAICreditsRecord"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/aiCredits/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchGetAICreditsRecordResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询算粒的使用记录
//
// @param request - BatchGetAICreditsRecordRequest
//
// @return BatchGetAICreditsRecordResponse
func (client *Client) BatchGetAICreditsRecord(request *BatchGetAICreditsRecordRequest) (_result *BatchGetAICreditsRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchGetAICreditsRecordHeaders{}
	_result = &BatchGetAICreditsRecordResponse{}
	_body, _err := client.BatchGetAICreditsRecordWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建AI助理
//
// @param request - CreateAssistantRequest
//
// @param headers - CreateAssistantHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAssistantResponse
func (client *Client) CreateAssistantWithOptions(request *CreateAssistantRequest, headers *CreateAssistantHeaders, runtime *util.RuntimeOptions) (_result *CreateAssistantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.Instructions)) {
		body["instructions"] = request.Instructions
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUnionId)) {
		body["operatorUnionId"] = request.OperatorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.RecommendPrompts)) {
		body["recommendPrompts"] = request.RecommendPrompts
	}

	if !tea.BoolValue(util.IsUnset(request.WelcomeContent)) {
		body["welcomeContent"] = request.WelcomeContent
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
		Action:      tea.String("CreateAssistant"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/basicInfo"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAssistantResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建AI助理
//
// @param request - CreateAssistantRequest
//
// @return CreateAssistantResponse
func (client *Client) CreateAssistant(request *CreateAssistantRequest) (_result *CreateAssistantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateAssistantHeaders{}
	_result = &CreateAssistantResponse{}
	_body, _err := client.CreateAssistantWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建AI助理的消息体
//
// @param request - CreateAssistantMessageRequest
//
// @param headers - CreateAssistantMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAssistantMessageResponse
func (client *Client) CreateAssistantMessageWithOptions(threadId *string, request *CreateAssistantMessageRequest, headers *CreateAssistantMessageHeaders, runtime *util.RuntimeOptions) (_result *CreateAssistantMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Extension)) {
		body["extension"] = request.Extension
	}

	if !tea.BoolValue(util.IsUnset(request.Metadata)) {
		body["metadata"] = request.Metadata
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		body["role"] = request.Role
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
		Action:      tea.String("CreateAssistantMessage"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/threads/" + tea.StringValue(threadId) + "/messages"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAssistantMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建AI助理的消息体
//
// @param request - CreateAssistantMessageRequest
//
// @return CreateAssistantMessageResponse
func (client *Client) CreateAssistantMessage(threadId *string, request *CreateAssistantMessageRequest) (_result *CreateAssistantMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateAssistantMessageHeaders{}
	_result = &CreateAssistantMessageResponse{}
	_body, _err := client.CreateAssistantMessageWithOptions(threadId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建AI助理的运行任务
//
// @param request - CreateAssistantRunRequest
//
// @param headers - CreateAssistantRunHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAssistantRunResponse
func (client *Client) CreateAssistantRunWithOptions(threadId *string, request *CreateAssistantRunRequest, headers *CreateAssistantRunHeaders, runtime *util.RuntimeOptions) (_result *CreateAssistantRunResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssistantId)) {
		body["assistantId"] = request.AssistantId
	}

	if !tea.BoolValue(util.IsUnset(request.Instructions)) {
		body["instructions"] = request.Instructions
	}

	if !tea.BoolValue(util.IsUnset(request.Metadata)) {
		body["metadata"] = request.Metadata
	}

	if !tea.BoolValue(util.IsUnset(request.Stream)) {
		body["stream"] = request.Stream
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
		Action:      tea.String("CreateAssistantRun"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/threads/" + tea.StringValue(threadId) + "/runs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAssistantRunResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建AI助理的运行任务
//
// @param request - CreateAssistantRunRequest
//
// @return CreateAssistantRunResponse
func (client *Client) CreateAssistantRun(threadId *string, request *CreateAssistantRunRequest) (_result *CreateAssistantRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateAssistantRunHeaders{}
	_result = &CreateAssistantRunResponse{}
	_body, _err := client.CreateAssistantRunWithOptions(threadId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建AI助理线程实例
//
// @param request - CreateAssistantThreadRequest
//
// @param headers - CreateAssistantThreadHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAssistantThreadResponse
func (client *Client) CreateAssistantThreadWithOptions(request *CreateAssistantThreadRequest, headers *CreateAssistantThreadHeaders, runtime *util.RuntimeOptions) (_result *CreateAssistantThreadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Metadata)) {
		body["metadata"] = request.Metadata
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
		Action:      tea.String("CreateAssistantThread"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/threads"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAssistantThreadResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建AI助理线程实例
//
// @param request - CreateAssistantThreadRequest
//
// @return CreateAssistantThreadResponse
func (client *Client) CreateAssistantThread(request *CreateAssistantThreadRequest) (_result *CreateAssistantThreadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateAssistantThreadHeaders{}
	_result = &CreateAssistantThreadResponse{}
	_body, _err := client.CreateAssistantThreadWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除AI助理
//
// @param request - DeleteAssistantRequest
//
// @param headers - DeleteAssistantHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAssistantResponse
func (client *Client) DeleteAssistantWithOptions(request *DeleteAssistantRequest, headers *DeleteAssistantHeaders, runtime *util.RuntimeOptions) (_result *DeleteAssistantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssistantId)) {
		query["assistantId"] = request.AssistantId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUnionId)) {
		query["operatorUnionId"] = request.OperatorUnionId
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
		Action:      tea.String("DeleteAssistant"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/basicInfo"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAssistantResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除AI助理
//
// @param request - DeleteAssistantRequest
//
// @return DeleteAssistantResponse
func (client *Client) DeleteAssistant(request *DeleteAssistantRequest) (_result *DeleteAssistantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteAssistantHeaders{}
	_result = &DeleteAssistantResponse{}
	_body, _err := client.DeleteAssistantWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除AI助理的消息体
//
// @param headers - DeleteAssistantMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAssistantMessageResponse
func (client *Client) DeleteAssistantMessageWithOptions(threadId *string, messageId *string, headers *DeleteAssistantMessageHeaders, runtime *util.RuntimeOptions) (_result *DeleteAssistantMessageResponse, _err error) {
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
		Action:      tea.String("DeleteAssistantMessage"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/threads/" + tea.StringValue(threadId) + "/messages/" + tea.StringValue(messageId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAssistantMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除AI助理的消息体
//
// @return DeleteAssistantMessageResponse
func (client *Client) DeleteAssistantMessage(threadId *string, messageId *string) (_result *DeleteAssistantMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteAssistantMessageHeaders{}
	_result = &DeleteAssistantMessageResponse{}
	_body, _err := client.DeleteAssistantMessageWithOptions(threadId, messageId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除AI助理线程实例
//
// @param headers - DeleteAssistantThreadHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAssistantThreadResponse
func (client *Client) DeleteAssistantThreadWithOptions(threadId *string, headers *DeleteAssistantThreadHeaders, runtime *util.RuntimeOptions) (_result *DeleteAssistantThreadResponse, _err error) {
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
		Action:      tea.String("DeleteAssistantThread"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/threads/" + tea.StringValue(threadId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAssistantThreadResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除AI助理线程实例
//
// @return DeleteAssistantThreadResponse
func (client *Client) DeleteAssistantThread(threadId *string) (_result *DeleteAssistantThreadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteAssistantThreadHeaders{}
	_result = &DeleteAssistantThreadResponse{}
	_body, _err := client.DeleteAssistantThreadWithOptions(threadId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 助理删除专业词汇
//
// @param request - DeleteDomainWordsRequest
//
// @param headers - DeleteDomainWordsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainWordsResponse
func (client *Client) DeleteDomainWordsWithOptions(request *DeleteDomainWordsRequest, headers *DeleteDomainWordsHeaders, runtime *util.RuntimeOptions) (_result *DeleteDomainWordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssistantId)) {
		body["assistantId"] = request.AssistantId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainWords)) {
		body["domainWords"] = request.DomainWords
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
		Action:      tea.String("DeleteDomainWords"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/domainWords/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDomainWordsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 助理删除专业词汇
//
// @param request - DeleteDomainWordsRequest
//
// @return DeleteDomainWordsResponse
func (client *Client) DeleteDomainWords(request *DeleteDomainWordsRequest) (_result *DeleteDomainWordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteDomainWordsHeaders{}
	_result = &DeleteDomainWordsResponse{}
	_body, _err := client.DeleteDomainWordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除助理知识
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssistantId)) {
		query["assistantId"] = request.AssistantId
	}

	if !tea.BoolValue(util.IsUnset(request.StudyId)) {
		query["studyId"] = request.StudyId
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
		Action:      tea.String("DeleteKnowledge"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/knowledges/items"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
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
// 删除助理知识
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
// 一键部署AI助理
//
// @param request - DeployAssistantRequest
//
// @param headers - DeployAssistantHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployAssistantResponse
func (client *Client) DeployAssistantWithOptions(request *DeployAssistantRequest, headers *DeployAssistantHeaders, runtime *util.RuntimeOptions) (_result *DeployAssistantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		body["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.AiAssistantId)) {
		body["aiAssistantId"] = request.AiAssistantId
	}

	if !tea.BoolValue(util.IsUnset(request.AppScopes)) {
		body["appScopes"] = request.AppScopes
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Fallback)) {
		body["fallback"] = request.Fallback
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.Instructions)) {
		body["instructions"] = request.Instructions
	}

	if !tea.BoolValue(util.IsUnset(request.IsPublic)) {
		body["isPublic"] = request.IsPublic
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OperateUserId)) {
		body["operateUserId"] = request.OperateUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RecommendPrompts)) {
		body["recommendPrompts"] = request.RecommendPrompts
	}

	if !tea.BoolValue(util.IsUnset(request.ShareRecipient)) {
		body["shareRecipient"] = request.ShareRecipient
	}

	if !tea.BoolValue(util.IsUnset(request.ToneStyle)) {
		body["toneStyle"] = request.ToneStyle
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
	}

	if !tea.BoolValue(util.IsUnset(request.WelcomeContent)) {
		body["welcomeContent"] = request.WelcomeContent
	}

	if !tea.BoolValue(util.IsUnset(request.WelcomeTitle)) {
		body["welcomeTitle"] = request.WelcomeTitle
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
		Action:      tea.String("DeployAssistant"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/deploy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeployAssistantResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 一键部署AI助理
//
// @param request - DeployAssistantRequest
//
// @return DeployAssistantResponse
func (client *Client) DeployAssistant(request *DeployAssistantRequest) (_result *DeployAssistantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeployAssistantHeaders{}
	_result = &DeployAssistantResponse{}
	_body, _err := client.DeployAssistantWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取助理问答明细
//
// @param request - GetAskDetailRequest
//
// @param headers - GetAskDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAskDetailResponse
func (client *Client) GetAskDetailWithOptions(request *GetAskDetailRequest, headers *GetAskDetailHeaders, runtime *util.RuntimeOptions) (_result *GetAskDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssistantId)) {
		query["assistantId"] = request.AssistantId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
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
		Action:      tea.String("GetAskDetail"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/askDetails"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAskDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取助理问答明细
//
// @param request - GetAskDetailRequest
//
// @return GetAskDetailResponse
func (client *Client) GetAskDetail(request *GetAskDetailRequest) (_result *GetAskDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAskDetailHeaders{}
	_result = &GetAskDetailResponse{}
	_body, _err := client.GetAskDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取AI助理技能列表信息
//
// @param request - GetAssistantActionInfoRequest
//
// @param headers - GetAssistantActionInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAssistantActionInfoResponse
func (client *Client) GetAssistantActionInfoWithOptions(request *GetAssistantActionInfoRequest, headers *GetAssistantActionInfoHeaders, runtime *util.RuntimeOptions) (_result *GetAssistantActionInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssistantId)) {
		query["assistantId"] = request.AssistantId
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
		Action:      tea.String("GetAssistantActionInfo"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/actionLists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAssistantActionInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取AI助理技能列表信息
//
// @param request - GetAssistantActionInfoRequest
//
// @return GetAssistantActionInfoResponse
func (client *Client) GetAssistantActionInfo(request *GetAssistantActionInfoRequest) (_result *GetAssistantActionInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAssistantActionInfoHeaders{}
	_result = &GetAssistantActionInfoResponse{}
	_body, _err := client.GetAssistantActionInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取助理专业词汇
//
// @param request - GetDomainWordsRequest
//
// @param headers - GetDomainWordsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDomainWordsResponse
func (client *Client) GetDomainWordsWithOptions(request *GetDomainWordsRequest, headers *GetDomainWordsHeaders, runtime *util.RuntimeOptions) (_result *GetDomainWordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssistantId)) {
		query["assistantId"] = request.AssistantId
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
		Action:      tea.String("GetDomainWords"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/domainWords"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDomainWordsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取助理专业词汇
//
// @param request - GetDomainWordsRequest
//
// @return GetDomainWordsResponse
func (client *Client) GetDomainWords(request *GetDomainWordsRequest) (_result *GetDomainWordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDomainWordsHeaders{}
	_result = &GetDomainWordsResponse{}
	_body, _err := client.GetDomainWordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取助理知识列表
//
// @param request - GetKnowledgeListRequest
//
// @param headers - GetKnowledgeListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKnowledgeListResponse
func (client *Client) GetKnowledgeListWithOptions(request *GetKnowledgeListRequest, headers *GetKnowledgeListHeaders, runtime *util.RuntimeOptions) (_result *GetKnowledgeListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssistantId)) {
		query["assistantId"] = request.AssistantId
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
		Action:      tea.String("GetKnowledgeList"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/knowledges/items"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetKnowledgeListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取助理知识列表
//
// @param request - GetKnowledgeListRequest
//
// @return GetKnowledgeListResponse
func (client *Client) GetKnowledgeList(request *GetKnowledgeListRequest) (_result *GetKnowledgeListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetKnowledgeListHeaders{}
	_result = &GetKnowledgeListResponse{}
	_body, _err := client.GetKnowledgeListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 安装助理
//
// @param request - InstallAssistantRequest
//
// @param headers - InstallAssistantHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallAssistantResponse
func (client *Client) InstallAssistantWithOptions(request *InstallAssistantRequest, headers *InstallAssistantHeaders, runtime *util.RuntimeOptions) (_result *InstallAssistantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssistantId)) {
		body["assistantId"] = request.AssistantId
	}

	if !tea.BoolValue(util.IsUnset(request.IsAllOrgMemberVisible)) {
		body["isAllOrgMemberVisible"] = request.IsAllOrgMemberVisible
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
		Action:      tea.String("InstallAssistant"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/install"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallAssistantResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 安装助理
//
// @param request - InstallAssistantRequest
//
// @return InstallAssistantResponse
func (client *Client) InstallAssistant(request *InstallAssistantRequest) (_result *InstallAssistantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InstallAssistantHeaders{}
	_result = &InstallAssistantResponse{}
	_body, _err := client.InstallAssistantWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 助理学习知识
//
// @param request - LearnKnowledgeRequest
//
// @param headers - LearnKnowledgeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LearnKnowledgeResponse
func (client *Client) LearnKnowledgeWithOptions(request *LearnKnowledgeRequest, headers *LearnKnowledgeHeaders, runtime *util.RuntimeOptions) (_result *LearnKnowledgeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssistantId)) {
		body["assistantId"] = request.AssistantId
	}

	if !tea.BoolValue(util.IsUnset(request.DocUrl)) {
		body["docUrl"] = request.DocUrl
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
		Action:      tea.String("LearnKnowledge"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/knowledges/items"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &LearnKnowledgeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 助理学习知识
//
// @param request - LearnKnowledgeRequest
//
// @return LearnKnowledgeResponse
func (client *Client) LearnKnowledge(request *LearnKnowledgeRequest) (_result *LearnKnowledgeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &LearnKnowledgeHeaders{}
	_result = &LearnKnowledgeResponse{}
	_body, _err := client.LearnKnowledgeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取AI助理列表
//
// @param request - ListAssistantRequest
//
// @param headers - ListAssistantHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAssistantResponse
func (client *Client) ListAssistantWithOptions(request *ListAssistantRequest, headers *ListAssistantHeaders, runtime *util.RuntimeOptions) (_result *ListAssistantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		query["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
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
		Action:      tea.String("ListAssistant"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAssistantResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取AI助理列表
//
// @param request - ListAssistantRequest
//
// @return ListAssistantResponse
func (client *Client) ListAssistant(request *ListAssistantRequest) (_result *ListAssistantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListAssistantHeaders{}
	_result = &ListAssistantResponse{}
	_body, _err := client.ListAssistantWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取AI助理消息列表
//
// @param request - ListAssistantMessageRequest
//
// @param headers - ListAssistantMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAssistantMessageResponse
func (client *Client) ListAssistantMessageWithOptions(threadId *string, request *ListAssistantMessageRequest, headers *ListAssistantMessageHeaders, runtime *util.RuntimeOptions) (_result *ListAssistantMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.RunId)) {
		query["runId"] = request.RunId
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
		Action:      tea.String("ListAssistantMessage"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/threads/" + tea.StringValue(threadId) + "/messages"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAssistantMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取AI助理消息列表
//
// @param request - ListAssistantMessageRequest
//
// @return ListAssistantMessageResponse
func (client *Client) ListAssistantMessage(threadId *string, request *ListAssistantMessageRequest) (_result *ListAssistantMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListAssistantMessageHeaders{}
	_result = &ListAssistantMessageResponse{}
	_body, _err := client.ListAssistantMessageWithOptions(threadId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取AI助理的运行任务的列表
//
// @param request - ListAssistantRunRequest
//
// @param headers - ListAssistantRunHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAssistantRunResponse
func (client *Client) ListAssistantRunWithOptions(threadId *string, request *ListAssistantRunRequest, headers *ListAssistantRunHeaders, runtime *util.RuntimeOptions) (_result *ListAssistantRunResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["order"] = request.Order
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
		Action:      tea.String("ListAssistantRun"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/threads/" + tea.StringValue(threadId) + "/runs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAssistantRunResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取AI助理的运行任务的列表
//
// @param request - ListAssistantRunRequest
//
// @return ListAssistantRunResponse
func (client *Client) ListAssistantRun(threadId *string, request *ListAssistantRunRequest) (_result *ListAssistantRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListAssistantRunHeaders{}
	_result = &ListAssistantRunResponse{}
	_body, _err := client.ListAssistantRunWithOptions(threadId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定助理在组织下已安装的实例信息列表
//
// @param request - ListInstanceRequest
//
// @param headers - ListInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceResponse
func (client *Client) ListInstanceWithOptions(request *ListInstanceRequest, headers *ListInstanceHeaders, runtime *util.RuntimeOptions) (_result *ListInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PrototypeAssistantId)) {
		query["prototypeAssistantId"] = request.PrototypeAssistantId
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
		Action:      tea.String("ListInstance"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/instances/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定助理在组织下已安装的实例信息列表
//
// @param request - ListInstanceRequest
//
// @return ListInstanceResponse
func (client *Client) ListInstance(request *ListInstanceRequest) (_result *ListInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListInstanceHeaders{}
	_result = &ListInstanceResponse{}
	_body, _err := client.ListInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户可见范围的AI助理列表
//
// @param request - ListVisibleAssistantRequest
//
// @param headers - ListVisibleAssistantHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVisibleAssistantResponse
func (client *Client) ListVisibleAssistantWithOptions(request *ListVisibleAssistantRequest, headers *ListVisibleAssistantHeaders, runtime *util.RuntimeOptions) (_result *ListVisibleAssistantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		query["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
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
		Action:      tea.String("ListVisibleAssistant"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/visibleList"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVisibleAssistantResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户可见范围的AI助理列表
//
// @param request - ListVisibleAssistantRequest
//
// @return ListVisibleAssistantResponse
func (client *Client) ListVisibleAssistant(request *ListVisibleAssistantRequest) (_result *ListVisibleAssistantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListVisibleAssistantHeaders{}
	_result = &ListVisibleAssistantResponse{}
	_body, _err := client.ListVisibleAssistantWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取对话明细列表
//
// @param request - LogListRequest
//
// @param headers - LogListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LogListResponse
func (client *Client) LogListWithOptions(request *LogListRequest, headers *LogListHeaders, runtime *util.RuntimeOptions) (_result *LogListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssistantId)) {
		query["assistantId"] = request.AssistantId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
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
		Action:      tea.String("LogList"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/logs/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &LogListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取对话明细列表
//
// @param request - LogListRequest
//
// @return LogListResponse
func (client *Client) LogList(request *LogListRequest) (_result *LogListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &LogListHeaders{}
	_result = &LogListResponse{}
	_body, _err := client.LogListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 助理学习增量知识
//
// @param request - RelearnKnowledgeRequest
//
// @param headers - RelearnKnowledgeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RelearnKnowledgeResponse
func (client *Client) RelearnKnowledgeWithOptions(request *RelearnKnowledgeRequest, headers *RelearnKnowledgeHeaders, runtime *util.RuntimeOptions) (_result *RelearnKnowledgeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssistantId)) {
		body["assistantId"] = request.AssistantId
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
		Action:      tea.String("RelearnKnowledge"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/knowledges/incrLearning"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RelearnKnowledgeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 助理学习增量知识
//
// @param request - RelearnKnowledgeRequest
//
// @return RelearnKnowledgeResponse
func (client *Client) RelearnKnowledge(request *RelearnKnowledgeRequest) (_result *RelearnKnowledgeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RelearnKnowledgeHeaders{}
	_result = &RelearnKnowledgeResponse{}
	_body, _err := client.RelearnKnowledgeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 卸载助理
//
// @param request - RemoveAssistantRequest
//
// @param headers - RemoveAssistantHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveAssistantResponse
func (client *Client) RemoveAssistantWithOptions(request *RemoveAssistantRequest, headers *RemoveAssistantHeaders, runtime *util.RuntimeOptions) (_result *RemoveAssistantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssistantId)) {
		body["assistantId"] = request.AssistantId
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
		Action:      tea.String("RemoveAssistant"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/uninstall"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveAssistantResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 卸载助理
//
// @param request - RemoveAssistantRequest
//
// @return RemoveAssistantResponse
func (client *Client) RemoveAssistant(request *RemoveAssistantRequest) (_result *RemoveAssistantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveAssistantHeaders{}
	_result = &RemoveAssistantResponse{}
	_body, _err := client.RemoveAssistantWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移除组织技能库技能
//
// @param request - RemoveFromOrgSkillRepositoryRequest
//
// @param headers - RemoveFromOrgSkillRepositoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveFromOrgSkillRepositoryResponse
func (client *Client) RemoveFromOrgSkillRepositoryWithOptions(request *RemoveFromOrgSkillRepositoryRequest, headers *RemoveFromOrgSkillRepositoryHeaders, runtime *util.RuntimeOptions) (_result *RemoveFromOrgSkillRepositoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionId)) {
		query["actionId"] = request.ActionId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUnionId)) {
		query["operatorUnionId"] = request.OperatorUnionId
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
		Action:      tea.String("RemoveFromOrgSkillRepository"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/orgActionRepositories"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveFromOrgSkillRepositoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移除组织技能库技能
//
// @param request - RemoveFromOrgSkillRepositoryRequest
//
// @return RemoveFromOrgSkillRepositoryResponse
func (client *Client) RemoveFromOrgSkillRepository(request *RemoveFromOrgSkillRepositoryRequest) (_result *RemoveFromOrgSkillRepositoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveFromOrgSkillRepositoryHeaders{}
	_result = &RemoveFromOrgSkillRepositoryResponse{}
	_body, _err := client.RemoveFromOrgSkillRepositoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询 AI 助理的基本信息
//
// @param request - RetrieveAssistantBasicInfoRequest
//
// @param headers - RetrieveAssistantBasicInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetrieveAssistantBasicInfoResponse
func (client *Client) RetrieveAssistantBasicInfoWithOptions(request *RetrieveAssistantBasicInfoRequest, headers *RetrieveAssistantBasicInfoHeaders, runtime *util.RuntimeOptions) (_result *RetrieveAssistantBasicInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssistantId)) {
		query["assistantId"] = request.AssistantId
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
		Action:      tea.String("RetrieveAssistantBasicInfo"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/basicInfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RetrieveAssistantBasicInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询 AI 助理的基本信息
//
// @param request - RetrieveAssistantBasicInfoRequest
//
// @return RetrieveAssistantBasicInfoResponse
func (client *Client) RetrieveAssistantBasicInfo(request *RetrieveAssistantBasicInfoRequest) (_result *RetrieveAssistantBasicInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RetrieveAssistantBasicInfoHeaders{}
	_result = &RetrieveAssistantBasicInfoResponse{}
	_body, _err := client.RetrieveAssistantBasicInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取AI助理的消息体
//
// @param headers - RetrieveAssistantMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetrieveAssistantMessageResponse
func (client *Client) RetrieveAssistantMessageWithOptions(threadId *string, messageId *string, headers *RetrieveAssistantMessageHeaders, runtime *util.RuntimeOptions) (_result *RetrieveAssistantMessageResponse, _err error) {
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
		Action:      tea.String("RetrieveAssistantMessage"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/threads/" + tea.StringValue(threadId) + "/messages/" + tea.StringValue(messageId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RetrieveAssistantMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取AI助理的消息体
//
// @return RetrieveAssistantMessageResponse
func (client *Client) RetrieveAssistantMessage(threadId *string, messageId *string) (_result *RetrieveAssistantMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RetrieveAssistantMessageHeaders{}
	_result = &RetrieveAssistantMessageResponse{}
	_body, _err := client.RetrieveAssistantMessageWithOptions(threadId, messageId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 检索AI助理的运行任务
//
// @param headers - RetrieveAssistantRunHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetrieveAssistantRunResponse
func (client *Client) RetrieveAssistantRunWithOptions(threadId *string, runId *string, headers *RetrieveAssistantRunHeaders, runtime *util.RuntimeOptions) (_result *RetrieveAssistantRunResponse, _err error) {
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
		Action:      tea.String("RetrieveAssistantRun"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/threads/" + tea.StringValue(threadId) + "/runs/" + tea.StringValue(runId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RetrieveAssistantRunResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检索AI助理的运行任务
//
// @return RetrieveAssistantRunResponse
func (client *Client) RetrieveAssistantRun(threadId *string, runId *string) (_result *RetrieveAssistantRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RetrieveAssistantRunHeaders{}
	_result = &RetrieveAssistantRunResponse{}
	_body, _err := client.RetrieveAssistantRunWithOptions(threadId, runId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取助理的使用范围
//
// @param request - RetrieveAssistantScopeRequest
//
// @param headers - RetrieveAssistantScopeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetrieveAssistantScopeResponse
func (client *Client) RetrieveAssistantScopeWithOptions(request *RetrieveAssistantScopeRequest, headers *RetrieveAssistantScopeHeaders, runtime *util.RuntimeOptions) (_result *RetrieveAssistantScopeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssistantId)) {
		query["assistantId"] = request.AssistantId
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
		Action:      tea.String("RetrieveAssistantScope"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/scope"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RetrieveAssistantScopeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取助理的使用范围
//
// @param request - RetrieveAssistantScopeRequest
//
// @return RetrieveAssistantScopeResponse
func (client *Client) RetrieveAssistantScope(request *RetrieveAssistantScopeRequest) (_result *RetrieveAssistantScopeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RetrieveAssistantScopeHeaders{}
	_result = &RetrieveAssistantScopeResponse{}
	_body, _err := client.RetrieveAssistantScopeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 检索AI助理线程实例
//
// @param headers - RetrieveAssistantThreadHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetrieveAssistantThreadResponse
func (client *Client) RetrieveAssistantThreadWithOptions(threadId *string, headers *RetrieveAssistantThreadHeaders, runtime *util.RuntimeOptions) (_result *RetrieveAssistantThreadResponse, _err error) {
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
		Action:      tea.String("RetrieveAssistantThread"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/threads/" + tea.StringValue(threadId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RetrieveAssistantThreadResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检索AI助理线程实例
//
// @return RetrieveAssistantThreadResponse
func (client *Client) RetrieveAssistantThread(threadId *string) (_result *RetrieveAssistantThreadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RetrieveAssistantThreadHeaders{}
	_result = &RetrieveAssistantThreadResponse{}
	_body, _err := client.RetrieveAssistantThreadWithOptions(threadId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新AI助理基础信息
//
// @param request - UpdateAssistantBasicInfoRequest
//
// @param headers - UpdateAssistantBasicInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAssistantBasicInfoResponse
func (client *Client) UpdateAssistantBasicInfoWithOptions(request *UpdateAssistantBasicInfoRequest, headers *UpdateAssistantBasicInfoHeaders, runtime *util.RuntimeOptions) (_result *UpdateAssistantBasicInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssistantId)) {
		body["assistantId"] = request.AssistantId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FallbackContent)) {
		body["fallbackContent"] = request.FallbackContent
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.Instructions)) {
		body["instructions"] = request.Instructions
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUnionId)) {
		body["operatorUnionId"] = request.OperatorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.RecommendPrompts)) {
		body["recommendPrompts"] = request.RecommendPrompts
	}

	if !tea.BoolValue(util.IsUnset(request.WelcomeContent)) {
		body["welcomeContent"] = request.WelcomeContent
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
		Action:      tea.String("UpdateAssistantBasicInfo"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/basicInfo"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAssistantBasicInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新AI助理基础信息
//
// @param request - UpdateAssistantBasicInfoRequest
//
// @return UpdateAssistantBasicInfoResponse
func (client *Client) UpdateAssistantBasicInfo(request *UpdateAssistantBasicInfoRequest) (_result *UpdateAssistantBasicInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateAssistantBasicInfoHeaders{}
	_result = &UpdateAssistantBasicInfoResponse{}
	_body, _err := client.UpdateAssistantBasicInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新 AI 助理使用范围
//
// @param request - UpdateAssistantScopeRequest
//
// @param headers - UpdateAssistantScopeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAssistantScopeResponse
func (client *Client) UpdateAssistantScopeWithOptions(request *UpdateAssistantScopeRequest, headers *UpdateAssistantScopeHeaders, runtime *util.RuntimeOptions) (_result *UpdateAssistantScopeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssistantId)) {
		body["assistantId"] = request.AssistantId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUnionId)) {
		body["operatorUnionId"] = request.OperatorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.Scopes)) {
		body["scopes"] = request.Scopes
	}

	if !tea.BoolValue(util.IsUnset(request.Sharing)) {
		body["sharing"] = request.Sharing
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
		Action:      tea.String("UpdateAssistantScope"),
		Version:     tea.String("assistant_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/assistant/scope"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("any"),
	}
	_result = &UpdateAssistantScopeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新 AI 助理使用范围
//
// @param request - UpdateAssistantScopeRequest
//
// @return UpdateAssistantScopeResponse
func (client *Client) UpdateAssistantScope(request *UpdateAssistantScopeRequest) (_result *UpdateAssistantScopeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateAssistantScopeHeaders{}
	_result = &UpdateAssistantScopeResponse{}
	_body, _err := client.UpdateAssistantScopeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
