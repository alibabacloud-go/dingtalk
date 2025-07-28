// This file is auto-generated, don't edit it. Thanks.
package ai_paa_s_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ExclusiveModelCompleteServiceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ExclusiveModelCompleteServiceHeaders) String() string {
	return tea.Prettify(s)
}

func (s ExclusiveModelCompleteServiceHeaders) GoString() string {
	return s.String()
}

func (s *ExclusiveModelCompleteServiceHeaders) SetCommonHeaders(v map[string]*string) *ExclusiveModelCompleteServiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ExclusiveModelCompleteServiceHeaders) SetXAcsDingtalkAccessToken(v string) *ExclusiveModelCompleteServiceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ExclusiveModelCompleteServiceRequest struct {
	EnableSearch *bool  `json:"enable_search,omitempty" xml:"enable_search,omitempty"`
	MaxTokens    *int32 `json:"max_tokens,omitempty" xml:"max_tokens,omitempty"`
	// This parameter is required.
	Messages []*ExclusiveModelCompleteServiceRequestMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// This parameter is required.
	Model       *string  `json:"model,omitempty" xml:"model,omitempty"`
	Stream      *bool    `json:"stream,omitempty" xml:"stream,omitempty"`
	Temperature *float64 `json:"temperature,omitempty" xml:"temperature,omitempty"`
	TopP        *float64 `json:"top_p,omitempty" xml:"top_p,omitempty"`
}

func (s ExclusiveModelCompleteServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s ExclusiveModelCompleteServiceRequest) GoString() string {
	return s.String()
}

func (s *ExclusiveModelCompleteServiceRequest) SetEnableSearch(v bool) *ExclusiveModelCompleteServiceRequest {
	s.EnableSearch = &v
	return s
}

func (s *ExclusiveModelCompleteServiceRequest) SetMaxTokens(v int32) *ExclusiveModelCompleteServiceRequest {
	s.MaxTokens = &v
	return s
}

func (s *ExclusiveModelCompleteServiceRequest) SetMessages(v []*ExclusiveModelCompleteServiceRequestMessages) *ExclusiveModelCompleteServiceRequest {
	s.Messages = v
	return s
}

func (s *ExclusiveModelCompleteServiceRequest) SetModel(v string) *ExclusiveModelCompleteServiceRequest {
	s.Model = &v
	return s
}

func (s *ExclusiveModelCompleteServiceRequest) SetStream(v bool) *ExclusiveModelCompleteServiceRequest {
	s.Stream = &v
	return s
}

func (s *ExclusiveModelCompleteServiceRequest) SetTemperature(v float64) *ExclusiveModelCompleteServiceRequest {
	s.Temperature = &v
	return s
}

func (s *ExclusiveModelCompleteServiceRequest) SetTopP(v float64) *ExclusiveModelCompleteServiceRequest {
	s.TopP = &v
	return s
}

type ExclusiveModelCompleteServiceRequestMessages struct {
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s ExclusiveModelCompleteServiceRequestMessages) String() string {
	return tea.Prettify(s)
}

func (s ExclusiveModelCompleteServiceRequestMessages) GoString() string {
	return s.String()
}

func (s *ExclusiveModelCompleteServiceRequestMessages) SetContent(v string) *ExclusiveModelCompleteServiceRequestMessages {
	s.Content = &v
	return s
}

func (s *ExclusiveModelCompleteServiceRequestMessages) SetRole(v string) *ExclusiveModelCompleteServiceRequestMessages {
	s.Role = &v
	return s
}

type ExclusiveModelCompleteServiceResponseBody struct {
	Choices []*ExclusiveModelCompleteServiceResponseBodyChoices `json:"choices,omitempty" xml:"choices,omitempty" type:"Repeated"`
	Created *int64                                              `json:"created,omitempty" xml:"created,omitempty"`
	Id      *string                                             `json:"id,omitempty" xml:"id,omitempty"`
	Model   *string                                             `json:"model,omitempty" xml:"model,omitempty"`
	Usage   *ExclusiveModelCompleteServiceResponseBodyUsage     `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s ExclusiveModelCompleteServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExclusiveModelCompleteServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ExclusiveModelCompleteServiceResponseBody) SetChoices(v []*ExclusiveModelCompleteServiceResponseBodyChoices) *ExclusiveModelCompleteServiceResponseBody {
	s.Choices = v
	return s
}

func (s *ExclusiveModelCompleteServiceResponseBody) SetCreated(v int64) *ExclusiveModelCompleteServiceResponseBody {
	s.Created = &v
	return s
}

func (s *ExclusiveModelCompleteServiceResponseBody) SetId(v string) *ExclusiveModelCompleteServiceResponseBody {
	s.Id = &v
	return s
}

func (s *ExclusiveModelCompleteServiceResponseBody) SetModel(v string) *ExclusiveModelCompleteServiceResponseBody {
	s.Model = &v
	return s
}

func (s *ExclusiveModelCompleteServiceResponseBody) SetUsage(v *ExclusiveModelCompleteServiceResponseBodyUsage) *ExclusiveModelCompleteServiceResponseBody {
	s.Usage = v
	return s
}

type ExclusiveModelCompleteServiceResponseBodyChoices struct {
	FinishReason *string                                                  `json:"finishReason,omitempty" xml:"finishReason,omitempty"`
	Message      *ExclusiveModelCompleteServiceResponseBodyChoicesMessage `json:"message,omitempty" xml:"message,omitempty" type:"Struct"`
}

func (s ExclusiveModelCompleteServiceResponseBodyChoices) String() string {
	return tea.Prettify(s)
}

func (s ExclusiveModelCompleteServiceResponseBodyChoices) GoString() string {
	return s.String()
}

func (s *ExclusiveModelCompleteServiceResponseBodyChoices) SetFinishReason(v string) *ExclusiveModelCompleteServiceResponseBodyChoices {
	s.FinishReason = &v
	return s
}

func (s *ExclusiveModelCompleteServiceResponseBodyChoices) SetMessage(v *ExclusiveModelCompleteServiceResponseBodyChoicesMessage) *ExclusiveModelCompleteServiceResponseBodyChoices {
	s.Message = v
	return s
}

type ExclusiveModelCompleteServiceResponseBodyChoicesMessage struct {
	Content          *string `json:"content,omitempty" xml:"content,omitempty"`
	ReasoningContent *string `json:"reasoning_content,omitempty" xml:"reasoning_content,omitempty"`
	Role             *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s ExclusiveModelCompleteServiceResponseBodyChoicesMessage) String() string {
	return tea.Prettify(s)
}

func (s ExclusiveModelCompleteServiceResponseBodyChoicesMessage) GoString() string {
	return s.String()
}

func (s *ExclusiveModelCompleteServiceResponseBodyChoicesMessage) SetContent(v string) *ExclusiveModelCompleteServiceResponseBodyChoicesMessage {
	s.Content = &v
	return s
}

func (s *ExclusiveModelCompleteServiceResponseBodyChoicesMessage) SetReasoningContent(v string) *ExclusiveModelCompleteServiceResponseBodyChoicesMessage {
	s.ReasoningContent = &v
	return s
}

func (s *ExclusiveModelCompleteServiceResponseBodyChoicesMessage) SetRole(v string) *ExclusiveModelCompleteServiceResponseBodyChoicesMessage {
	s.Role = &v
	return s
}

type ExclusiveModelCompleteServiceResponseBodyUsage struct {
	CompletionTokens *int32 `json:"completion_tokens,omitempty" xml:"completion_tokens,omitempty"`
	PromptTokens     *int32 `json:"prompt_tokens,omitempty" xml:"prompt_tokens,omitempty"`
	TotalTokens      *int32 `json:"total_tokens,omitempty" xml:"total_tokens,omitempty"`
}

func (s ExclusiveModelCompleteServiceResponseBodyUsage) String() string {
	return tea.Prettify(s)
}

func (s ExclusiveModelCompleteServiceResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *ExclusiveModelCompleteServiceResponseBodyUsage) SetCompletionTokens(v int32) *ExclusiveModelCompleteServiceResponseBodyUsage {
	s.CompletionTokens = &v
	return s
}

func (s *ExclusiveModelCompleteServiceResponseBodyUsage) SetPromptTokens(v int32) *ExclusiveModelCompleteServiceResponseBodyUsage {
	s.PromptTokens = &v
	return s
}

func (s *ExclusiveModelCompleteServiceResponseBodyUsage) SetTotalTokens(v int32) *ExclusiveModelCompleteServiceResponseBodyUsage {
	s.TotalTokens = &v
	return s
}

type ExclusiveModelCompleteServiceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExclusiveModelCompleteServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExclusiveModelCompleteServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s ExclusiveModelCompleteServiceResponse) GoString() string {
	return s.String()
}

func (s *ExclusiveModelCompleteServiceResponse) SetHeaders(v map[string]*string) *ExclusiveModelCompleteServiceResponse {
	s.Headers = v
	return s
}

func (s *ExclusiveModelCompleteServiceResponse) SetStatusCode(v int32) *ExclusiveModelCompleteServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ExclusiveModelCompleteServiceResponse) SetBody(v *ExclusiveModelCompleteServiceResponseBody) *ExclusiveModelCompleteServiceResponse {
	s.Body = v
	return s
}

type ExecuteAgentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ExecuteAgentHeaders) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAgentHeaders) GoString() string {
	return s.String()
}

func (s *ExecuteAgentHeaders) SetCommonHeaders(v map[string]*string) *ExecuteAgentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ExecuteAgentHeaders) SetXAcsDingtalkAccessToken(v string) *ExecuteAgentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ExecuteAgentRequest struct {
	// This parameter is required.
	AgentCode *string `json:"agentCode,omitempty" xml:"agentCode,omitempty"`
	// This parameter is required.
	Inputs             *ExecuteAgentRequestInputs `json:"inputs,omitempty" xml:"inputs,omitempty" type:"Struct"`
	ScenarioCode       *string                    `json:"scenarioCode,omitempty" xml:"scenarioCode,omitempty"`
	ScenarioInstanceId *string                    `json:"scenarioInstanceId,omitempty" xml:"scenarioInstanceId,omitempty"`
	SkillId            *string                    `json:"skillId,omitempty" xml:"skillId,omitempty"`
}

func (s ExecuteAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAgentRequest) GoString() string {
	return s.String()
}

func (s *ExecuteAgentRequest) SetAgentCode(v string) *ExecuteAgentRequest {
	s.AgentCode = &v
	return s
}

func (s *ExecuteAgentRequest) SetInputs(v *ExecuteAgentRequestInputs) *ExecuteAgentRequest {
	s.Inputs = v
	return s
}

func (s *ExecuteAgentRequest) SetScenarioCode(v string) *ExecuteAgentRequest {
	s.ScenarioCode = &v
	return s
}

func (s *ExecuteAgentRequest) SetScenarioInstanceId(v string) *ExecuteAgentRequest {
	s.ScenarioInstanceId = &v
	return s
}

func (s *ExecuteAgentRequest) SetSkillId(v string) *ExecuteAgentRequest {
	s.SkillId = &v
	return s
}

type ExecuteAgentRequestInputs struct {
	CardData       interface{} `json:"cardData,omitempty" xml:"cardData,omitempty"`
	CardTemplateId *string     `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	Input          *string     `json:"input,omitempty" xml:"input,omitempty"`
}

func (s ExecuteAgentRequestInputs) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAgentRequestInputs) GoString() string {
	return s.String()
}

func (s *ExecuteAgentRequestInputs) SetCardData(v interface{}) *ExecuteAgentRequestInputs {
	s.CardData = v
	return s
}

func (s *ExecuteAgentRequestInputs) SetCardTemplateId(v string) *ExecuteAgentRequestInputs {
	s.CardTemplateId = &v
	return s
}

func (s *ExecuteAgentRequestInputs) SetInput(v string) *ExecuteAgentRequestInputs {
	s.Input = &v
	return s
}

type ExecuteAgentResponseBody struct {
	Result *ExecuteAgentResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ExecuteAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAgentResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteAgentResponseBody) SetResult(v *ExecuteAgentResponseBodyResult) *ExecuteAgentResponseBody {
	s.Result = v
	return s
}

type ExecuteAgentResponseBodyResult struct {
	ExecuteResult *string `json:"executeResult,omitempty" xml:"executeResult,omitempty"`
	SkillId       *string `json:"skillId,omitempty" xml:"skillId,omitempty"`
}

func (s ExecuteAgentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAgentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ExecuteAgentResponseBodyResult) SetExecuteResult(v string) *ExecuteAgentResponseBodyResult {
	s.ExecuteResult = &v
	return s
}

func (s *ExecuteAgentResponseBodyResult) SetSkillId(v string) *ExecuteAgentResponseBodyResult {
	s.SkillId = &v
	return s
}

type ExecuteAgentResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAgentResponse) GoString() string {
	return s.String()
}

func (s *ExecuteAgentResponse) SetHeaders(v map[string]*string) *ExecuteAgentResponse {
	s.Headers = v
	return s
}

func (s *ExecuteAgentResponse) SetStatusCode(v int32) *ExecuteAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteAgentResponse) SetBody(v *ExecuteAgentResponseBody) *ExecuteAgentResponse {
	s.Body = v
	return s
}

type LiandanTextImageGetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s LiandanTextImageGetHeaders) String() string {
	return tea.Prettify(s)
}

func (s LiandanTextImageGetHeaders) GoString() string {
	return s.String()
}

func (s *LiandanTextImageGetHeaders) SetCommonHeaders(v map[string]*string) *LiandanTextImageGetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *LiandanTextImageGetHeaders) SetXAcsDingtalkAccessToken(v string) *LiandanTextImageGetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type LiandanTextImageGetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// IMAGE
	Module *string `json:"module,omitempty" xml:"module,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s LiandanTextImageGetRequest) String() string {
	return tea.Prettify(s)
}

func (s LiandanTextImageGetRequest) GoString() string {
	return s.String()
}

func (s *LiandanTextImageGetRequest) SetModule(v string) *LiandanTextImageGetRequest {
	s.Module = &v
	return s
}

func (s *LiandanTextImageGetRequest) SetTaskId(v string) *LiandanTextImageGetRequest {
	s.TaskId = &v
	return s
}

func (s *LiandanTextImageGetRequest) SetUserId(v string) *LiandanTextImageGetRequest {
	s.UserId = &v
	return s
}

type LiandanTextImageGetResponseBody struct {
	Result []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s LiandanTextImageGetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LiandanTextImageGetResponseBody) GoString() string {
	return s.String()
}

func (s *LiandanTextImageGetResponseBody) SetResult(v []map[string]interface{}) *LiandanTextImageGetResponseBody {
	s.Result = v
	return s
}

func (s *LiandanTextImageGetResponseBody) SetSuccess(v bool) *LiandanTextImageGetResponseBody {
	s.Success = &v
	return s
}

type LiandanTextImageGetResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LiandanTextImageGetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LiandanTextImageGetResponse) String() string {
	return tea.Prettify(s)
}

func (s LiandanTextImageGetResponse) GoString() string {
	return s.String()
}

func (s *LiandanTextImageGetResponse) SetHeaders(v map[string]*string) *LiandanTextImageGetResponse {
	s.Headers = v
	return s
}

func (s *LiandanTextImageGetResponse) SetStatusCode(v int32) *LiandanTextImageGetResponse {
	s.StatusCode = &v
	return s
}

func (s *LiandanTextImageGetResponse) SetBody(v *LiandanTextImageGetResponseBody) *LiandanTextImageGetResponse {
	s.Body = v
	return s
}

type LiandanluExclusiveModelHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s LiandanluExclusiveModelHeaders) String() string {
	return tea.Prettify(s)
}

func (s LiandanluExclusiveModelHeaders) GoString() string {
	return s.String()
}

func (s *LiandanluExclusiveModelHeaders) SetCommonHeaders(v map[string]*string) *LiandanluExclusiveModelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *LiandanluExclusiveModelHeaders) SetXAcsDingtalkAccessToken(v string) *LiandanluExclusiveModelHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type LiandanluExclusiveModelRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// maas1234
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// GENERAL
	Module *string `json:"module,omitempty" xml:"module,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OKR是什么
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 使用该功能的用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s LiandanluExclusiveModelRequest) String() string {
	return tea.Prettify(s)
}

func (s LiandanluExclusiveModelRequest) GoString() string {
	return s.String()
}

func (s *LiandanluExclusiveModelRequest) SetModelId(v string) *LiandanluExclusiveModelRequest {
	s.ModelId = &v
	return s
}

func (s *LiandanluExclusiveModelRequest) SetModule(v string) *LiandanluExclusiveModelRequest {
	s.Module = &v
	return s
}

func (s *LiandanluExclusiveModelRequest) SetPrompt(v string) *LiandanluExclusiveModelRequest {
	s.Prompt = &v
	return s
}

func (s *LiandanluExclusiveModelRequest) SetUserId(v string) *LiandanluExclusiveModelRequest {
	s.UserId = &v
	return s
}

type LiandanluExclusiveModelResponseBody struct {
	// example:
	//
	// requestId_123
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// {        "content":"OKR 全称为 Objective and Key Results，即目标与关键结果法，是一套明确和跟踪目标及其完成情况的管理工具和方法。"   }
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s LiandanluExclusiveModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LiandanluExclusiveModelResponseBody) GoString() string {
	return s.String()
}

func (s *LiandanluExclusiveModelResponseBody) SetRequestId(v string) *LiandanluExclusiveModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *LiandanluExclusiveModelResponseBody) SetResult(v map[string]interface{}) *LiandanluExclusiveModelResponseBody {
	s.Result = v
	return s
}

type LiandanluExclusiveModelResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LiandanluExclusiveModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LiandanluExclusiveModelResponse) String() string {
	return tea.Prettify(s)
}

func (s LiandanluExclusiveModelResponse) GoString() string {
	return s.String()
}

func (s *LiandanluExclusiveModelResponse) SetHeaders(v map[string]*string) *LiandanluExclusiveModelResponse {
	s.Headers = v
	return s
}

func (s *LiandanluExclusiveModelResponse) SetStatusCode(v int32) *LiandanluExclusiveModelResponse {
	s.StatusCode = &v
	return s
}

func (s *LiandanluExclusiveModelResponse) SetBody(v *LiandanluExclusiveModelResponseBody) *LiandanluExclusiveModelResponse {
	s.Body = v
	return s
}

type LiandanluTextToImageModelHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s LiandanluTextToImageModelHeaders) String() string {
	return tea.Prettify(s)
}

func (s LiandanluTextToImageModelHeaders) GoString() string {
	return s.String()
}

func (s *LiandanluTextToImageModelHeaders) SetCommonHeaders(v map[string]*string) *LiandanluTextToImageModelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *LiandanluTextToImageModelHeaders) SetXAcsDingtalkAccessToken(v string) *LiandanluTextToImageModelHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type LiandanluTextToImageModelRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// IMAGE
	Module *string `json:"module,omitempty" xml:"module,omitempty"`
	// example:
	//
	// 1
	Number     *int64             `json:"number,omitempty" xml:"number,omitempty"`
	Parameters map[string]*string `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 画一副风景画
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s LiandanluTextToImageModelRequest) String() string {
	return tea.Prettify(s)
}

func (s LiandanluTextToImageModelRequest) GoString() string {
	return s.String()
}

func (s *LiandanluTextToImageModelRequest) SetModule(v string) *LiandanluTextToImageModelRequest {
	s.Module = &v
	return s
}

func (s *LiandanluTextToImageModelRequest) SetNumber(v int64) *LiandanluTextToImageModelRequest {
	s.Number = &v
	return s
}

func (s *LiandanluTextToImageModelRequest) SetParameters(v map[string]*string) *LiandanluTextToImageModelRequest {
	s.Parameters = v
	return s
}

func (s *LiandanluTextToImageModelRequest) SetPrompt(v string) *LiandanluTextToImageModelRequest {
	s.Prompt = &v
	return s
}

func (s *LiandanluTextToImageModelRequest) SetUserId(v string) *LiandanluTextToImageModelRequest {
	s.UserId = &v
	return s
}

type LiandanluTextToImageModelResponseBody struct {
	// This parameter is required.
	Result *LiandanluTextToImageModelResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s LiandanluTextToImageModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LiandanluTextToImageModelResponseBody) GoString() string {
	return s.String()
}

func (s *LiandanluTextToImageModelResponseBody) SetResult(v *LiandanluTextToImageModelResponseBodyResult) *LiandanluTextToImageModelResponseBody {
	s.Result = v
	return s
}

func (s *LiandanluTextToImageModelResponseBody) SetSuccess(v bool) *LiandanluTextToImageModelResponseBody {
	s.Success = &v
	return s
}

type LiandanluTextToImageModelResponseBodyResult struct {
	// example:
	//
	// 0112_1222
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 123
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// SUCCEEDED
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s LiandanluTextToImageModelResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s LiandanluTextToImageModelResponseBodyResult) GoString() string {
	return s.String()
}

func (s *LiandanluTextToImageModelResponseBodyResult) SetRequestId(v string) *LiandanluTextToImageModelResponseBodyResult {
	s.RequestId = &v
	return s
}

func (s *LiandanluTextToImageModelResponseBodyResult) SetTaskId(v string) *LiandanluTextToImageModelResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *LiandanluTextToImageModelResponseBodyResult) SetTaskStatus(v string) *LiandanluTextToImageModelResponseBodyResult {
	s.TaskStatus = &v
	return s
}

type LiandanluTextToImageModelResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LiandanluTextToImageModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LiandanluTextToImageModelResponse) String() string {
	return tea.Prettify(s)
}

func (s LiandanluTextToImageModelResponse) GoString() string {
	return s.String()
}

func (s *LiandanluTextToImageModelResponse) SetHeaders(v map[string]*string) *LiandanluTextToImageModelResponse {
	s.Headers = v
	return s
}

func (s *LiandanluTextToImageModelResponse) SetStatusCode(v int32) *LiandanluTextToImageModelResponse {
	s.StatusCode = &v
	return s
}

func (s *LiandanluTextToImageModelResponse) SetBody(v *LiandanluTextToImageModelResponseBody) *LiandanluTextToImageModelResponse {
	s.Body = v
	return s
}

type NLToFrameServiceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s NLToFrameServiceHeaders) String() string {
	return tea.Prettify(s)
}

func (s NLToFrameServiceHeaders) GoString() string {
	return s.String()
}

func (s *NLToFrameServiceHeaders) SetCommonHeaders(v map[string]*string) *NLToFrameServiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *NLToFrameServiceHeaders) SetXAcsDingtalkAccessToken(v string) *NLToFrameServiceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type NLToFrameServiceRequest struct {
	ExtensionStr *string `json:"extensionStr,omitempty" xml:"extensionStr,omitempty"`
	IsNewModel   *bool   `json:"isNewModel,omitempty" xml:"isNewModel,omitempty"`
	// This parameter is required.
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// This parameter is required.
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	UserId    *int64  `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s NLToFrameServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s NLToFrameServiceRequest) GoString() string {
	return s.String()
}

func (s *NLToFrameServiceRequest) SetExtensionStr(v string) *NLToFrameServiceRequest {
	s.ExtensionStr = &v
	return s
}

func (s *NLToFrameServiceRequest) SetIsNewModel(v bool) *NLToFrameServiceRequest {
	s.IsNewModel = &v
	return s
}

func (s *NLToFrameServiceRequest) SetModelId(v string) *NLToFrameServiceRequest {
	s.ModelId = &v
	return s
}

func (s *NLToFrameServiceRequest) SetModelName(v string) *NLToFrameServiceRequest {
	s.ModelName = &v
	return s
}

func (s *NLToFrameServiceRequest) SetUserId(v int64) *NLToFrameServiceRequest {
	s.UserId = &v
	return s
}

type NLToFrameServiceResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s NLToFrameServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NLToFrameServiceResponseBody) GoString() string {
	return s.String()
}

func (s *NLToFrameServiceResponseBody) SetResult(v string) *NLToFrameServiceResponseBody {
	s.Result = &v
	return s
}

type NLToFrameServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NLToFrameServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NLToFrameServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s NLToFrameServiceResponse) GoString() string {
	return s.String()
}

func (s *NLToFrameServiceResponse) SetHeaders(v map[string]*string) *NLToFrameServiceResponse {
	s.Headers = v
	return s
}

func (s *NLToFrameServiceResponse) SetStatusCode(v int32) *NLToFrameServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *NLToFrameServiceResponse) SetBody(v *NLToFrameServiceResponseBody) *NLToFrameServiceResponse {
	s.Body = v
	return s
}

type QueryBaymaxSkillLogHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryBaymaxSkillLogHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryBaymaxSkillLogHeaders) GoString() string {
	return s.String()
}

func (s *QueryBaymaxSkillLogHeaders) SetCommonHeaders(v map[string]*string) *QueryBaymaxSkillLogHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryBaymaxSkillLogHeaders) SetXAcsDingtalkAccessToken(v string) *QueryBaymaxSkillLogHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryBaymaxSkillLogRequest struct {
	From *int32 `json:"from,omitempty" xml:"from,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 14da****2760
	LogLevel *string `json:"logLevel,omitempty" xml:"logLevel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 14da****2760
	SkillExecuteId *string `json:"skillExecuteId,omitempty" xml:"skillExecuteId,omitempty"`
	To             *int32  `json:"to,omitempty" xml:"to,omitempty"`
}

func (s QueryBaymaxSkillLogRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBaymaxSkillLogRequest) GoString() string {
	return s.String()
}

func (s *QueryBaymaxSkillLogRequest) SetFrom(v int32) *QueryBaymaxSkillLogRequest {
	s.From = &v
	return s
}

func (s *QueryBaymaxSkillLogRequest) SetLogLevel(v string) *QueryBaymaxSkillLogRequest {
	s.LogLevel = &v
	return s
}

func (s *QueryBaymaxSkillLogRequest) SetSkillExecuteId(v string) *QueryBaymaxSkillLogRequest {
	s.SkillExecuteId = &v
	return s
}

func (s *QueryBaymaxSkillLogRequest) SetTo(v int32) *QueryBaymaxSkillLogRequest {
	s.To = &v
	return s
}

type QueryBaymaxSkillLogResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 14da****2760
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s QueryBaymaxSkillLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBaymaxSkillLogResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBaymaxSkillLogResponseBody) SetResult(v string) *QueryBaymaxSkillLogResponseBody {
	s.Result = &v
	return s
}

type QueryBaymaxSkillLogResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBaymaxSkillLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBaymaxSkillLogResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBaymaxSkillLogResponse) GoString() string {
	return s.String()
}

func (s *QueryBaymaxSkillLogResponse) SetHeaders(v map[string]*string) *QueryBaymaxSkillLogResponse {
	s.Headers = v
	return s
}

func (s *QueryBaymaxSkillLogResponse) SetStatusCode(v int32) *QueryBaymaxSkillLogResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBaymaxSkillLogResponse) SetBody(v *QueryBaymaxSkillLogResponseBody) *QueryBaymaxSkillLogResponse {
	s.Body = v
	return s
}

type QueryConversationMessageForAIHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryConversationMessageForAIHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryConversationMessageForAIHeaders) GoString() string {
	return s.String()
}

func (s *QueryConversationMessageForAIHeaders) SetCommonHeaders(v map[string]*string) *QueryConversationMessageForAIHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryConversationMessageForAIHeaders) SetXAcsDingtalkAccessToken(v string) *QueryConversationMessageForAIHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryConversationMessageForAIRequest struct {
	OpenMsgIds  []*string `json:"openMsgIds,omitempty" xml:"openMsgIds,omitempty" type:"Repeated"`
	RecentDays  *int32    `json:"recentDays,omitempty" xml:"recentDays,omitempty"`
	RecentHours *int32    `json:"recentHours,omitempty" xml:"recentHours,omitempty"`
	RecentN     *int32    `json:"recentN,omitempty" xml:"recentN,omitempty"`
}

func (s QueryConversationMessageForAIRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryConversationMessageForAIRequest) GoString() string {
	return s.String()
}

func (s *QueryConversationMessageForAIRequest) SetOpenMsgIds(v []*string) *QueryConversationMessageForAIRequest {
	s.OpenMsgIds = v
	return s
}

func (s *QueryConversationMessageForAIRequest) SetRecentDays(v int32) *QueryConversationMessageForAIRequest {
	s.RecentDays = &v
	return s
}

func (s *QueryConversationMessageForAIRequest) SetRecentHours(v int32) *QueryConversationMessageForAIRequest {
	s.RecentHours = &v
	return s
}

func (s *QueryConversationMessageForAIRequest) SetRecentN(v int32) *QueryConversationMessageForAIRequest {
	s.RecentN = &v
	return s
}

type QueryConversationMessageForAIShrinkRequest struct {
	OpenMsgIdsShrink *string `json:"openMsgIds,omitempty" xml:"openMsgIds,omitempty"`
	RecentDays       *int32  `json:"recentDays,omitempty" xml:"recentDays,omitempty"`
	RecentHours      *int32  `json:"recentHours,omitempty" xml:"recentHours,omitempty"`
	RecentN          *int32  `json:"recentN,omitempty" xml:"recentN,omitempty"`
}

func (s QueryConversationMessageForAIShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryConversationMessageForAIShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryConversationMessageForAIShrinkRequest) SetOpenMsgIdsShrink(v string) *QueryConversationMessageForAIShrinkRequest {
	s.OpenMsgIdsShrink = &v
	return s
}

func (s *QueryConversationMessageForAIShrinkRequest) SetRecentDays(v int32) *QueryConversationMessageForAIShrinkRequest {
	s.RecentDays = &v
	return s
}

func (s *QueryConversationMessageForAIShrinkRequest) SetRecentHours(v int32) *QueryConversationMessageForAIShrinkRequest {
	s.RecentHours = &v
	return s
}

func (s *QueryConversationMessageForAIShrinkRequest) SetRecentN(v int32) *QueryConversationMessageForAIShrinkRequest {
	s.RecentN = &v
	return s
}

type QueryConversationMessageForAIResponseBody struct {
	Messages []*QueryConversationMessageForAIResponseBodyMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
}

func (s QueryConversationMessageForAIResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryConversationMessageForAIResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConversationMessageForAIResponseBody) SetMessages(v []*QueryConversationMessageForAIResponseBodyMessages) *QueryConversationMessageForAIResponseBody {
	s.Messages = v
	return s
}

type QueryConversationMessageForAIResponseBodyMessages struct {
	AtAll      *bool                                                       `json:"atAll,omitempty" xml:"atAll,omitempty"`
	AtUsers    []*QueryConversationMessageForAIResponseBodyMessagesAtUsers `json:"atUsers,omitempty" xml:"atUsers,omitempty" type:"Repeated"`
	MsgContent *string                                                     `json:"msgContent,omitempty" xml:"msgContent,omitempty"`
	MsgType    *string                                                     `json:"msgType,omitempty" xml:"msgType,omitempty"`
	SendTime   *string                                                     `json:"sendTime,omitempty" xml:"sendTime,omitempty"`
	Sender     *QueryConversationMessageForAIResponseBodyMessagesSender    `json:"sender,omitempty" xml:"sender,omitempty" type:"Struct"`
	Summary    *string                                                     `json:"summary,omitempty" xml:"summary,omitempty"`
}

func (s QueryConversationMessageForAIResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s QueryConversationMessageForAIResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *QueryConversationMessageForAIResponseBodyMessages) SetAtAll(v bool) *QueryConversationMessageForAIResponseBodyMessages {
	s.AtAll = &v
	return s
}

func (s *QueryConversationMessageForAIResponseBodyMessages) SetAtUsers(v []*QueryConversationMessageForAIResponseBodyMessagesAtUsers) *QueryConversationMessageForAIResponseBodyMessages {
	s.AtUsers = v
	return s
}

func (s *QueryConversationMessageForAIResponseBodyMessages) SetMsgContent(v string) *QueryConversationMessageForAIResponseBodyMessages {
	s.MsgContent = &v
	return s
}

func (s *QueryConversationMessageForAIResponseBodyMessages) SetMsgType(v string) *QueryConversationMessageForAIResponseBodyMessages {
	s.MsgType = &v
	return s
}

func (s *QueryConversationMessageForAIResponseBodyMessages) SetSendTime(v string) *QueryConversationMessageForAIResponseBodyMessages {
	s.SendTime = &v
	return s
}

func (s *QueryConversationMessageForAIResponseBodyMessages) SetSender(v *QueryConversationMessageForAIResponseBodyMessagesSender) *QueryConversationMessageForAIResponseBodyMessages {
	s.Sender = v
	return s
}

func (s *QueryConversationMessageForAIResponseBodyMessages) SetSummary(v string) *QueryConversationMessageForAIResponseBodyMessages {
	s.Summary = &v
	return s
}

type QueryConversationMessageForAIResponseBodyMessagesAtUsers struct {
	AgentCode *string `json:"agentCode,omitempty" xml:"agentCode,omitempty"`
	Nick      *string `json:"nick,omitempty" xml:"nick,omitempty"`
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
	UnionId   *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryConversationMessageForAIResponseBodyMessagesAtUsers) String() string {
	return tea.Prettify(s)
}

func (s QueryConversationMessageForAIResponseBodyMessagesAtUsers) GoString() string {
	return s.String()
}

func (s *QueryConversationMessageForAIResponseBodyMessagesAtUsers) SetAgentCode(v string) *QueryConversationMessageForAIResponseBodyMessagesAtUsers {
	s.AgentCode = &v
	return s
}

func (s *QueryConversationMessageForAIResponseBodyMessagesAtUsers) SetNick(v string) *QueryConversationMessageForAIResponseBodyMessagesAtUsers {
	s.Nick = &v
	return s
}

func (s *QueryConversationMessageForAIResponseBodyMessagesAtUsers) SetType(v string) *QueryConversationMessageForAIResponseBodyMessagesAtUsers {
	s.Type = &v
	return s
}

func (s *QueryConversationMessageForAIResponseBodyMessagesAtUsers) SetUnionId(v string) *QueryConversationMessageForAIResponseBodyMessagesAtUsers {
	s.UnionId = &v
	return s
}

func (s *QueryConversationMessageForAIResponseBodyMessagesAtUsers) SetUserId(v string) *QueryConversationMessageForAIResponseBodyMessagesAtUsers {
	s.UserId = &v
	return s
}

type QueryConversationMessageForAIResponseBodyMessagesSender struct {
	AgentCode *string `json:"agentCode,omitempty" xml:"agentCode,omitempty"`
	Nick      *string `json:"nick,omitempty" xml:"nick,omitempty"`
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
	UnionId   *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryConversationMessageForAIResponseBodyMessagesSender) String() string {
	return tea.Prettify(s)
}

func (s QueryConversationMessageForAIResponseBodyMessagesSender) GoString() string {
	return s.String()
}

func (s *QueryConversationMessageForAIResponseBodyMessagesSender) SetAgentCode(v string) *QueryConversationMessageForAIResponseBodyMessagesSender {
	s.AgentCode = &v
	return s
}

func (s *QueryConversationMessageForAIResponseBodyMessagesSender) SetNick(v string) *QueryConversationMessageForAIResponseBodyMessagesSender {
	s.Nick = &v
	return s
}

func (s *QueryConversationMessageForAIResponseBodyMessagesSender) SetType(v string) *QueryConversationMessageForAIResponseBodyMessagesSender {
	s.Type = &v
	return s
}

func (s *QueryConversationMessageForAIResponseBodyMessagesSender) SetUnionId(v string) *QueryConversationMessageForAIResponseBodyMessagesSender {
	s.UnionId = &v
	return s
}

func (s *QueryConversationMessageForAIResponseBodyMessagesSender) SetUserId(v string) *QueryConversationMessageForAIResponseBodyMessagesSender {
	s.UserId = &v
	return s
}

type QueryConversationMessageForAIResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryConversationMessageForAIResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryConversationMessageForAIResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryConversationMessageForAIResponse) GoString() string {
	return s.String()
}

func (s *QueryConversationMessageForAIResponse) SetHeaders(v map[string]*string) *QueryConversationMessageForAIResponse {
	s.Headers = v
	return s
}

func (s *QueryConversationMessageForAIResponse) SetStatusCode(v int32) *QueryConversationMessageForAIResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryConversationMessageForAIResponse) SetBody(v *QueryConversationMessageForAIResponseBody) *QueryConversationMessageForAIResponse {
	s.Body = v
	return s
}

type QueryMemoryLearningTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMemoryLearningTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMemoryLearningTaskHeaders) GoString() string {
	return s.String()
}

func (s *QueryMemoryLearningTaskHeaders) SetCommonHeaders(v map[string]*string) *QueryMemoryLearningTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMemoryLearningTaskHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMemoryLearningTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMemoryLearningTaskRequest struct {
	// This parameter is required.
	AgentCode *string `json:"agentCode,omitempty" xml:"agentCode,omitempty"`
	// This parameter is required.
	LearningCode *string `json:"learningCode,omitempty" xml:"learningCode,omitempty"`
}

func (s QueryMemoryLearningTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMemoryLearningTaskRequest) GoString() string {
	return s.String()
}

func (s *QueryMemoryLearningTaskRequest) SetAgentCode(v string) *QueryMemoryLearningTaskRequest {
	s.AgentCode = &v
	return s
}

func (s *QueryMemoryLearningTaskRequest) SetLearningCode(v string) *QueryMemoryLearningTaskRequest {
	s.LearningCode = &v
	return s
}

type QueryMemoryLearningTaskResponseBody struct {
	Result *QueryMemoryLearningTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryMemoryLearningTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMemoryLearningTaskResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMemoryLearningTaskResponseBody) SetResult(v *QueryMemoryLearningTaskResponseBodyResult) *QueryMemoryLearningTaskResponseBody {
	s.Result = v
	return s
}

type QueryMemoryLearningTaskResponseBodyResult struct {
	Status  *string `json:"status,omitempty" xml:"status,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryMemoryLearningTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryMemoryLearningTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryMemoryLearningTaskResponseBodyResult) SetStatus(v string) *QueryMemoryLearningTaskResponseBodyResult {
	s.Status = &v
	return s
}

func (s *QueryMemoryLearningTaskResponseBodyResult) SetSuccess(v bool) *QueryMemoryLearningTaskResponseBodyResult {
	s.Success = &v
	return s
}

type QueryMemoryLearningTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMemoryLearningTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMemoryLearningTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMemoryLearningTaskResponse) GoString() string {
	return s.String()
}

func (s *QueryMemoryLearningTaskResponse) SetHeaders(v map[string]*string) *QueryMemoryLearningTaskResponse {
	s.Headers = v
	return s
}

func (s *QueryMemoryLearningTaskResponse) SetStatusCode(v int32) *QueryMemoryLearningTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMemoryLearningTaskResponse) SetBody(v *QueryMemoryLearningTaskResponseBody) *QueryMemoryLearningTaskResponse {
	s.Body = v
	return s
}

type SmartFormulaResultServiceRequest struct {
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SmartFormulaResultServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s SmartFormulaResultServiceRequest) GoString() string {
	return s.String()
}

func (s *SmartFormulaResultServiceRequest) SetTaskId(v string) *SmartFormulaResultServiceRequest {
	s.TaskId = &v
	return s
}

type SmartFormulaResultServiceResponseBody struct {
	Result  *SmartFormulaResultServiceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                        `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SmartFormulaResultServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SmartFormulaResultServiceResponseBody) GoString() string {
	return s.String()
}

func (s *SmartFormulaResultServiceResponseBody) SetResult(v *SmartFormulaResultServiceResponseBodyResult) *SmartFormulaResultServiceResponseBody {
	s.Result = v
	return s
}

func (s *SmartFormulaResultServiceResponseBody) SetSuccess(v bool) *SmartFormulaResultServiceResponseBody {
	s.Success = &v
	return s
}

type SmartFormulaResultServiceResponseBodyResult struct {
	Response *string `json:"response,omitempty" xml:"response,omitempty"`
	Status   *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SmartFormulaResultServiceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SmartFormulaResultServiceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SmartFormulaResultServiceResponseBodyResult) SetResponse(v string) *SmartFormulaResultServiceResponseBodyResult {
	s.Response = &v
	return s
}

func (s *SmartFormulaResultServiceResponseBodyResult) SetStatus(v string) *SmartFormulaResultServiceResponseBodyResult {
	s.Status = &v
	return s
}

type SmartFormulaResultServiceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SmartFormulaResultServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SmartFormulaResultServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s SmartFormulaResultServiceResponse) GoString() string {
	return s.String()
}

func (s *SmartFormulaResultServiceResponse) SetHeaders(v map[string]*string) *SmartFormulaResultServiceResponse {
	s.Headers = v
	return s
}

func (s *SmartFormulaResultServiceResponse) SetStatusCode(v int32) *SmartFormulaResultServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *SmartFormulaResultServiceResponse) SetBody(v *SmartFormulaResultServiceResponseBody) *SmartFormulaResultServiceResponse {
	s.Body = v
	return s
}

type SmartFormulaTriggerServiceRequest struct {
	Request *string `json:"request,omitempty" xml:"request,omitempty"`
}

func (s SmartFormulaTriggerServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s SmartFormulaTriggerServiceRequest) GoString() string {
	return s.String()
}

func (s *SmartFormulaTriggerServiceRequest) SetRequest(v string) *SmartFormulaTriggerServiceRequest {
	s.Request = &v
	return s
}

type SmartFormulaTriggerServiceResponseBody struct {
	Result  *SmartFormulaTriggerServiceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SmartFormulaTriggerServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SmartFormulaTriggerServiceResponseBody) GoString() string {
	return s.String()
}

func (s *SmartFormulaTriggerServiceResponseBody) SetResult(v *SmartFormulaTriggerServiceResponseBodyResult) *SmartFormulaTriggerServiceResponseBody {
	s.Result = v
	return s
}

func (s *SmartFormulaTriggerServiceResponseBody) SetSuccess(v bool) *SmartFormulaTriggerServiceResponseBody {
	s.Success = &v
	return s
}

type SmartFormulaTriggerServiceResponseBodyResult struct {
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SmartFormulaTriggerServiceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SmartFormulaTriggerServiceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SmartFormulaTriggerServiceResponseBodyResult) SetTaskId(v string) *SmartFormulaTriggerServiceResponseBodyResult {
	s.TaskId = &v
	return s
}

type SmartFormulaTriggerServiceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SmartFormulaTriggerServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SmartFormulaTriggerServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s SmartFormulaTriggerServiceResponse) GoString() string {
	return s.String()
}

func (s *SmartFormulaTriggerServiceResponse) SetHeaders(v map[string]*string) *SmartFormulaTriggerServiceResponse {
	s.Headers = v
	return s
}

func (s *SmartFormulaTriggerServiceResponse) SetStatusCode(v int32) *SmartFormulaTriggerServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *SmartFormulaTriggerServiceResponse) SetBody(v *SmartFormulaTriggerServiceResponseBody) *SmartFormulaTriggerServiceResponse {
	s.Body = v
	return s
}

type SmartQuoteBatchQueryResultServiceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SmartQuoteBatchQueryResultServiceHeaders) String() string {
	return tea.Prettify(s)
}

func (s SmartQuoteBatchQueryResultServiceHeaders) GoString() string {
	return s.String()
}

func (s *SmartQuoteBatchQueryResultServiceHeaders) SetCommonHeaders(v map[string]*string) *SmartQuoteBatchQueryResultServiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SmartQuoteBatchQueryResultServiceHeaders) SetXAcsDingtalkAccessToken(v string) *SmartQuoteBatchQueryResultServiceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SmartQuoteBatchQueryResultServiceRequest struct {
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SmartQuoteBatchQueryResultServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s SmartQuoteBatchQueryResultServiceRequest) GoString() string {
	return s.String()
}

func (s *SmartQuoteBatchQueryResultServiceRequest) SetTaskId(v string) *SmartQuoteBatchQueryResultServiceRequest {
	s.TaskId = &v
	return s
}

type SmartQuoteBatchQueryResultServiceResponseBody struct {
	Result  *SmartQuoteBatchQueryResultServiceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                                `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SmartQuoteBatchQueryResultServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SmartQuoteBatchQueryResultServiceResponseBody) GoString() string {
	return s.String()
}

func (s *SmartQuoteBatchQueryResultServiceResponseBody) SetResult(v *SmartQuoteBatchQueryResultServiceResponseBodyResult) *SmartQuoteBatchQueryResultServiceResponseBody {
	s.Result = v
	return s
}

func (s *SmartQuoteBatchQueryResultServiceResponseBody) SetSuccess(v bool) *SmartQuoteBatchQueryResultServiceResponseBody {
	s.Success = &v
	return s
}

type SmartQuoteBatchQueryResultServiceResponseBodyResult struct {
	Response *string `json:"response,omitempty" xml:"response,omitempty"`
	Status   *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SmartQuoteBatchQueryResultServiceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SmartQuoteBatchQueryResultServiceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SmartQuoteBatchQueryResultServiceResponseBodyResult) SetResponse(v string) *SmartQuoteBatchQueryResultServiceResponseBodyResult {
	s.Response = &v
	return s
}

func (s *SmartQuoteBatchQueryResultServiceResponseBodyResult) SetStatus(v string) *SmartQuoteBatchQueryResultServiceResponseBodyResult {
	s.Status = &v
	return s
}

type SmartQuoteBatchQueryResultServiceResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SmartQuoteBatchQueryResultServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SmartQuoteBatchQueryResultServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s SmartQuoteBatchQueryResultServiceResponse) GoString() string {
	return s.String()
}

func (s *SmartQuoteBatchQueryResultServiceResponse) SetHeaders(v map[string]*string) *SmartQuoteBatchQueryResultServiceResponse {
	s.Headers = v
	return s
}

func (s *SmartQuoteBatchQueryResultServiceResponse) SetStatusCode(v int32) *SmartQuoteBatchQueryResultServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *SmartQuoteBatchQueryResultServiceResponse) SetBody(v *SmartQuoteBatchQueryResultServiceResponseBody) *SmartQuoteBatchQueryResultServiceResponse {
	s.Body = v
	return s
}

type SmartQuoteBatchQueryServiceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SmartQuoteBatchQueryServiceHeaders) String() string {
	return tea.Prettify(s)
}

func (s SmartQuoteBatchQueryServiceHeaders) GoString() string {
	return s.String()
}

func (s *SmartQuoteBatchQueryServiceHeaders) SetCommonHeaders(v map[string]*string) *SmartQuoteBatchQueryServiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SmartQuoteBatchQueryServiceHeaders) SetXAcsDingtalkAccessToken(v string) *SmartQuoteBatchQueryServiceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SmartQuoteBatchQueryServiceRequest struct {
	Request *string `json:"request,omitempty" xml:"request,omitempty"`
}

func (s SmartQuoteBatchQueryServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s SmartQuoteBatchQueryServiceRequest) GoString() string {
	return s.String()
}

func (s *SmartQuoteBatchQueryServiceRequest) SetRequest(v string) *SmartQuoteBatchQueryServiceRequest {
	s.Request = &v
	return s
}

type SmartQuoteBatchQueryServiceResponseBody struct {
	Result  *SmartQuoteBatchQueryServiceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SmartQuoteBatchQueryServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SmartQuoteBatchQueryServiceResponseBody) GoString() string {
	return s.String()
}

func (s *SmartQuoteBatchQueryServiceResponseBody) SetResult(v *SmartQuoteBatchQueryServiceResponseBodyResult) *SmartQuoteBatchQueryServiceResponseBody {
	s.Result = v
	return s
}

func (s *SmartQuoteBatchQueryServiceResponseBody) SetSuccess(v bool) *SmartQuoteBatchQueryServiceResponseBody {
	s.Success = &v
	return s
}

type SmartQuoteBatchQueryServiceResponseBodyResult struct {
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SmartQuoteBatchQueryServiceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SmartQuoteBatchQueryServiceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SmartQuoteBatchQueryServiceResponseBodyResult) SetTaskId(v string) *SmartQuoteBatchQueryServiceResponseBodyResult {
	s.TaskId = &v
	return s
}

type SmartQuoteBatchQueryServiceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SmartQuoteBatchQueryServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SmartQuoteBatchQueryServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s SmartQuoteBatchQueryServiceResponse) GoString() string {
	return s.String()
}

func (s *SmartQuoteBatchQueryServiceResponse) SetHeaders(v map[string]*string) *SmartQuoteBatchQueryServiceResponse {
	s.Headers = v
	return s
}

func (s *SmartQuoteBatchQueryServiceResponse) SetStatusCode(v int32) *SmartQuoteBatchQueryServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *SmartQuoteBatchQueryServiceResponse) SetBody(v *SmartQuoteBatchQueryServiceResponseBody) *SmartQuoteBatchQueryServiceResponse {
	s.Body = v
	return s
}

type SmartQuoteDataServiceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SmartQuoteDataServiceHeaders) String() string {
	return tea.Prettify(s)
}

func (s SmartQuoteDataServiceHeaders) GoString() string {
	return s.String()
}

func (s *SmartQuoteDataServiceHeaders) SetCommonHeaders(v map[string]*string) *SmartQuoteDataServiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SmartQuoteDataServiceHeaders) SetXAcsDingtalkAccessToken(v string) *SmartQuoteDataServiceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SmartQuoteDataServiceRequest struct {
	Request *string `json:"request,omitempty" xml:"request,omitempty"`
}

func (s SmartQuoteDataServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s SmartQuoteDataServiceRequest) GoString() string {
	return s.String()
}

func (s *SmartQuoteDataServiceRequest) SetRequest(v string) *SmartQuoteDataServiceRequest {
	s.Request = &v
	return s
}

type SmartQuoteDataServiceResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SmartQuoteDataServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SmartQuoteDataServiceResponseBody) GoString() string {
	return s.String()
}

func (s *SmartQuoteDataServiceResponseBody) SetResult(v bool) *SmartQuoteDataServiceResponseBody {
	s.Result = &v
	return s
}

type SmartQuoteDataServiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SmartQuoteDataServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SmartQuoteDataServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s SmartQuoteDataServiceResponse) GoString() string {
	return s.String()
}

func (s *SmartQuoteDataServiceResponse) SetHeaders(v map[string]*string) *SmartQuoteDataServiceResponse {
	s.Headers = v
	return s
}

func (s *SmartQuoteDataServiceResponse) SetStatusCode(v int32) *SmartQuoteDataServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *SmartQuoteDataServiceResponse) SetBody(v *SmartQuoteDataServiceResponseBody) *SmartQuoteDataServiceResponse {
	s.Body = v
	return s
}

type SmartQuoteQueryResultServiceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SmartQuoteQueryResultServiceHeaders) String() string {
	return tea.Prettify(s)
}

func (s SmartQuoteQueryResultServiceHeaders) GoString() string {
	return s.String()
}

func (s *SmartQuoteQueryResultServiceHeaders) SetCommonHeaders(v map[string]*string) *SmartQuoteQueryResultServiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SmartQuoteQueryResultServiceHeaders) SetXAcsDingtalkAccessToken(v string) *SmartQuoteQueryResultServiceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SmartQuoteQueryResultServiceRequest struct {
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SmartQuoteQueryResultServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s SmartQuoteQueryResultServiceRequest) GoString() string {
	return s.String()
}

func (s *SmartQuoteQueryResultServiceRequest) SetTaskId(v string) *SmartQuoteQueryResultServiceRequest {
	s.TaskId = &v
	return s
}

type SmartQuoteQueryResultServiceResponseBody struct {
	Result  *SmartQuoteQueryResultServiceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SmartQuoteQueryResultServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SmartQuoteQueryResultServiceResponseBody) GoString() string {
	return s.String()
}

func (s *SmartQuoteQueryResultServiceResponseBody) SetResult(v *SmartQuoteQueryResultServiceResponseBodyResult) *SmartQuoteQueryResultServiceResponseBody {
	s.Result = v
	return s
}

func (s *SmartQuoteQueryResultServiceResponseBody) SetSuccess(v bool) *SmartQuoteQueryResultServiceResponseBody {
	s.Success = &v
	return s
}

type SmartQuoteQueryResultServiceResponseBodyResult struct {
	Response *string `json:"response,omitempty" xml:"response,omitempty"`
	Status   *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SmartQuoteQueryResultServiceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SmartQuoteQueryResultServiceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SmartQuoteQueryResultServiceResponseBodyResult) SetResponse(v string) *SmartQuoteQueryResultServiceResponseBodyResult {
	s.Response = &v
	return s
}

func (s *SmartQuoteQueryResultServiceResponseBodyResult) SetStatus(v string) *SmartQuoteQueryResultServiceResponseBodyResult {
	s.Status = &v
	return s
}

type SmartQuoteQueryResultServiceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SmartQuoteQueryResultServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SmartQuoteQueryResultServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s SmartQuoteQueryResultServiceResponse) GoString() string {
	return s.String()
}

func (s *SmartQuoteQueryResultServiceResponse) SetHeaders(v map[string]*string) *SmartQuoteQueryResultServiceResponse {
	s.Headers = v
	return s
}

func (s *SmartQuoteQueryResultServiceResponse) SetStatusCode(v int32) *SmartQuoteQueryResultServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *SmartQuoteQueryResultServiceResponse) SetBody(v *SmartQuoteQueryResultServiceResponseBody) *SmartQuoteQueryResultServiceResponse {
	s.Body = v
	return s
}

type SmartQuoteQueryServiceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SmartQuoteQueryServiceHeaders) String() string {
	return tea.Prettify(s)
}

func (s SmartQuoteQueryServiceHeaders) GoString() string {
	return s.String()
}

func (s *SmartQuoteQueryServiceHeaders) SetCommonHeaders(v map[string]*string) *SmartQuoteQueryServiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SmartQuoteQueryServiceHeaders) SetXAcsDingtalkAccessToken(v string) *SmartQuoteQueryServiceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SmartQuoteQueryServiceRequest struct {
	Request *string `json:"request,omitempty" xml:"request,omitempty"`
}

func (s SmartQuoteQueryServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s SmartQuoteQueryServiceRequest) GoString() string {
	return s.String()
}

func (s *SmartQuoteQueryServiceRequest) SetRequest(v string) *SmartQuoteQueryServiceRequest {
	s.Request = &v
	return s
}

type SmartQuoteQueryServiceResponseBody struct {
	Result  *SmartQuoteQueryServiceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SmartQuoteQueryServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SmartQuoteQueryServiceResponseBody) GoString() string {
	return s.String()
}

func (s *SmartQuoteQueryServiceResponseBody) SetResult(v *SmartQuoteQueryServiceResponseBodyResult) *SmartQuoteQueryServiceResponseBody {
	s.Result = v
	return s
}

func (s *SmartQuoteQueryServiceResponseBody) SetSuccess(v bool) *SmartQuoteQueryServiceResponseBody {
	s.Success = &v
	return s
}

type SmartQuoteQueryServiceResponseBodyResult struct {
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SmartQuoteQueryServiceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SmartQuoteQueryServiceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SmartQuoteQueryServiceResponseBodyResult) SetTaskId(v string) *SmartQuoteQueryServiceResponseBodyResult {
	s.TaskId = &v
	return s
}

type SmartQuoteQueryServiceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SmartQuoteQueryServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SmartQuoteQueryServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s SmartQuoteQueryServiceResponse) GoString() string {
	return s.String()
}

func (s *SmartQuoteQueryServiceResponse) SetHeaders(v map[string]*string) *SmartQuoteQueryServiceResponse {
	s.Headers = v
	return s
}

func (s *SmartQuoteQueryServiceResponse) SetStatusCode(v int32) *SmartQuoteQueryServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *SmartQuoteQueryServiceResponse) SetBody(v *SmartQuoteQueryServiceResponseBody) *SmartQuoteQueryServiceResponse {
	s.Body = v
	return s
}

type SubmitMemoryLearningTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SubmitMemoryLearningTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s SubmitMemoryLearningTaskHeaders) GoString() string {
	return s.String()
}

func (s *SubmitMemoryLearningTaskHeaders) SetCommonHeaders(v map[string]*string) *SubmitMemoryLearningTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SubmitMemoryLearningTaskHeaders) SetXAcsDingtalkAccessToken(v string) *SubmitMemoryLearningTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SubmitMemoryLearningTaskRequest struct {
	// This parameter is required.
	AgentCode *string `json:"agentCode,omitempty" xml:"agentCode,omitempty"`
	// This parameter is required.
	Content *SubmitMemoryLearningTaskRequestContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	// This parameter is required.
	LearningMode *string `json:"learningMode,omitempty" xml:"learningMode,omitempty"`
	// This parameter is required.
	MemoryKey *string `json:"memoryKey,omitempty" xml:"memoryKey,omitempty"`
}

func (s SubmitMemoryLearningTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitMemoryLearningTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitMemoryLearningTaskRequest) SetAgentCode(v string) *SubmitMemoryLearningTaskRequest {
	s.AgentCode = &v
	return s
}

func (s *SubmitMemoryLearningTaskRequest) SetContent(v *SubmitMemoryLearningTaskRequestContent) *SubmitMemoryLearningTaskRequest {
	s.Content = v
	return s
}

func (s *SubmitMemoryLearningTaskRequest) SetLearningMode(v string) *SubmitMemoryLearningTaskRequest {
	s.LearningMode = &v
	return s
}

func (s *SubmitMemoryLearningTaskRequest) SetMemoryKey(v string) *SubmitMemoryLearningTaskRequest {
	s.MemoryKey = &v
	return s
}

type SubmitMemoryLearningTaskRequestContent struct {
	KnowledgeBaseUrl *string `json:"knowledgeBaseUrl,omitempty" xml:"knowledgeBaseUrl,omitempty"`
	Type             *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SubmitMemoryLearningTaskRequestContent) String() string {
	return tea.Prettify(s)
}

func (s SubmitMemoryLearningTaskRequestContent) GoString() string {
	return s.String()
}

func (s *SubmitMemoryLearningTaskRequestContent) SetKnowledgeBaseUrl(v string) *SubmitMemoryLearningTaskRequestContent {
	s.KnowledgeBaseUrl = &v
	return s
}

func (s *SubmitMemoryLearningTaskRequestContent) SetType(v string) *SubmitMemoryLearningTaskRequestContent {
	s.Type = &v
	return s
}

type SubmitMemoryLearningTaskShrinkRequest struct {
	// This parameter is required.
	AgentCode *string `json:"agentCode,omitempty" xml:"agentCode,omitempty"`
	// This parameter is required.
	ContentShrink *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	LearningMode *string `json:"learningMode,omitempty" xml:"learningMode,omitempty"`
	// This parameter is required.
	MemoryKey *string `json:"memoryKey,omitempty" xml:"memoryKey,omitempty"`
}

func (s SubmitMemoryLearningTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitMemoryLearningTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitMemoryLearningTaskShrinkRequest) SetAgentCode(v string) *SubmitMemoryLearningTaskShrinkRequest {
	s.AgentCode = &v
	return s
}

func (s *SubmitMemoryLearningTaskShrinkRequest) SetContentShrink(v string) *SubmitMemoryLearningTaskShrinkRequest {
	s.ContentShrink = &v
	return s
}

func (s *SubmitMemoryLearningTaskShrinkRequest) SetLearningMode(v string) *SubmitMemoryLearningTaskShrinkRequest {
	s.LearningMode = &v
	return s
}

func (s *SubmitMemoryLearningTaskShrinkRequest) SetMemoryKey(v string) *SubmitMemoryLearningTaskShrinkRequest {
	s.MemoryKey = &v
	return s
}

type SubmitMemoryLearningTaskResponseBody struct {
	Result *SubmitMemoryLearningTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s SubmitMemoryLearningTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitMemoryLearningTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitMemoryLearningTaskResponseBody) SetResult(v *SubmitMemoryLearningTaskResponseBodyResult) *SubmitMemoryLearningTaskResponseBody {
	s.Result = v
	return s
}

type SubmitMemoryLearningTaskResponseBodyResult struct {
	LearningCode *string `json:"learningCode,omitempty" xml:"learningCode,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SubmitMemoryLearningTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SubmitMemoryLearningTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SubmitMemoryLearningTaskResponseBodyResult) SetLearningCode(v string) *SubmitMemoryLearningTaskResponseBodyResult {
	s.LearningCode = &v
	return s
}

func (s *SubmitMemoryLearningTaskResponseBodyResult) SetStatus(v string) *SubmitMemoryLearningTaskResponseBodyResult {
	s.Status = &v
	return s
}

func (s *SubmitMemoryLearningTaskResponseBodyResult) SetSuccess(v bool) *SubmitMemoryLearningTaskResponseBodyResult {
	s.Success = &v
	return s
}

type SubmitMemoryLearningTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitMemoryLearningTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitMemoryLearningTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitMemoryLearningTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitMemoryLearningTaskResponse) SetHeaders(v map[string]*string) *SubmitMemoryLearningTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitMemoryLearningTaskResponse) SetStatusCode(v int32) *SubmitMemoryLearningTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitMemoryLearningTaskResponse) SetBody(v *SubmitMemoryLearningTaskResponseBody) *SubmitMemoryLearningTaskResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

// Summary:
//
// 炼丹炉专属模型推理服务
//
// @param request - ExclusiveModelCompleteServiceRequest
//
// @param headers - ExclusiveModelCompleteServiceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExclusiveModelCompleteServiceResponse
func (client *Client) ExclusiveModelCompleteServiceWithOptions(request *ExclusiveModelCompleteServiceRequest, headers *ExclusiveModelCompleteServiceHeaders, runtime *util.RuntimeOptions) (_result *ExclusiveModelCompleteServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnableSearch)) {
		body["enable_search"] = request.EnableSearch
	}

	if !tea.BoolValue(util.IsUnset(request.MaxTokens)) {
		body["max_tokens"] = request.MaxTokens
	}

	if !tea.BoolValue(util.IsUnset(request.Messages)) {
		body["messages"] = request.Messages
	}

	if !tea.BoolValue(util.IsUnset(request.Model)) {
		body["model"] = request.Model
	}

	if !tea.BoolValue(util.IsUnset(request.Stream)) {
		body["stream"] = request.Stream
	}

	if !tea.BoolValue(util.IsUnset(request.Temperature)) {
		body["temperature"] = request.Temperature
	}

	if !tea.BoolValue(util.IsUnset(request.TopP)) {
		body["top_p"] = request.TopP
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
		Action:      tea.String("ExclusiveModelCompleteService"),
		Version:     tea.String("aiPaaS_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiPaaS/ai/complete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ExclusiveModelCompleteServiceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 炼丹炉专属模型推理服务
//
// @param request - ExclusiveModelCompleteServiceRequest
//
// @return ExclusiveModelCompleteServiceResponse
func (client *Client) ExclusiveModelCompleteService(request *ExclusiveModelCompleteServiceRequest) (_result *ExclusiveModelCompleteServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ExclusiveModelCompleteServiceHeaders{}
	_result = &ExclusiveModelCompleteServiceResponse{}
	_body, _err := client.ExclusiveModelCompleteServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 执行AI技能
//
// @param request - ExecuteAgentRequest
//
// @param headers - ExecuteAgentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAgentResponse
func (client *Client) ExecuteAgentWithOptions(request *ExecuteAgentRequest, headers *ExecuteAgentHeaders, runtime *util.RuntimeOptions) (_result *ExecuteAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentCode)) {
		body["agentCode"] = request.AgentCode
	}

	if !tea.BoolValue(util.IsUnset(request.Inputs)) {
		body["inputs"] = request.Inputs
	}

	if !tea.BoolValue(util.IsUnset(request.ScenarioCode)) {
		body["scenarioCode"] = request.ScenarioCode
	}

	if !tea.BoolValue(util.IsUnset(request.ScenarioInstanceId)) {
		body["scenarioInstanceId"] = request.ScenarioInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SkillId)) {
		body["skillId"] = request.SkillId
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
		Action:      tea.String("ExecuteAgent"),
		Version:     tea.String("aiPaaS_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiPaaS/me/agents/execute"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteAgentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行AI技能
//
// @param request - ExecuteAgentRequest
//
// @return ExecuteAgentResponse
func (client *Client) ExecuteAgent(request *ExecuteAgentRequest) (_result *ExecuteAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ExecuteAgentHeaders{}
	_result = &ExecuteAgentResponse{}
	_body, _err := client.ExecuteAgentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 炼丹炉文生图任务结果获取
//
// @param request - LiandanTextImageGetRequest
//
// @param headers - LiandanTextImageGetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LiandanTextImageGetResponse
func (client *Client) LiandanTextImageGetWithOptions(request *LiandanTextImageGetRequest, headers *LiandanTextImageGetHeaders, runtime *util.RuntimeOptions) (_result *LiandanTextImageGetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Module)) {
		body["module"] = request.Module
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
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
		Action:      tea.String("LiandanTextImageGet"),
		Version:     tea.String("aiPaaS_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiPaaS/ai/textToImage/results/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &LiandanTextImageGetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 炼丹炉文生图任务结果获取
//
// @param request - LiandanTextImageGetRequest
//
// @return LiandanTextImageGetResponse
func (client *Client) LiandanTextImageGet(request *LiandanTextImageGetRequest) (_result *LiandanTextImageGetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &LiandanTextImageGetHeaders{}
	_result = &LiandanTextImageGetResponse{}
	_body, _err := client.LiandanTextImageGetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 炼丹炉专属模型接口
//
// @param request - LiandanluExclusiveModelRequest
//
// @param headers - LiandanluExclusiveModelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LiandanluExclusiveModelResponse
func (client *Client) LiandanluExclusiveModelWithOptions(request *LiandanluExclusiveModelRequest, headers *LiandanluExclusiveModelHeaders, runtime *util.RuntimeOptions) (_result *LiandanluExclusiveModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["modelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.Module)) {
		body["module"] = request.Module
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["prompt"] = request.Prompt
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
		Action:      tea.String("LiandanluExclusiveModel"),
		Version:     tea.String("aiPaaS_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiPaaS/ai/generate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &LiandanluExclusiveModelResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 炼丹炉专属模型接口
//
// @param request - LiandanluExclusiveModelRequest
//
// @return LiandanluExclusiveModelResponse
func (client *Client) LiandanluExclusiveModel(request *LiandanluExclusiveModelRequest) (_result *LiandanluExclusiveModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &LiandanluExclusiveModelHeaders{}
	_result = &LiandanluExclusiveModelResponse{}
	_body, _err := client.LiandanluExclusiveModelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 炼丹炉通过提示词生成图片
//
// @param request - LiandanluTextToImageModelRequest
//
// @param headers - LiandanluTextToImageModelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LiandanluTextToImageModelResponse
func (client *Client) LiandanluTextToImageModelWithOptions(request *LiandanluTextToImageModelRequest, headers *LiandanluTextToImageModelHeaders, runtime *util.RuntimeOptions) (_result *LiandanluTextToImageModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Module)) {
		body["module"] = request.Module
	}

	if !tea.BoolValue(util.IsUnset(request.Number)) {
		body["number"] = request.Number
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		body["parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["prompt"] = request.Prompt
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
		Action:      tea.String("LiandanluTextToImageModel"),
		Version:     tea.String("aiPaaS_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiPaaS/ai/textToImage/generate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &LiandanluTextToImageModelResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 炼丹炉通过提示词生成图片
//
// @param request - LiandanluTextToImageModelRequest
//
// @return LiandanluTextToImageModelResponse
func (client *Client) LiandanluTextToImageModel(request *LiandanluTextToImageModelRequest) (_result *LiandanluTextToImageModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &LiandanluTextToImageModelHeaders{}
	_result = &LiandanluTextToImageModelResponse{}
	_body, _err := client.LiandanluTextToImageModelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过配置的指令，连接用户和系统，训练大模型
//
// @param request - NLToFrameServiceRequest
//
// @param headers - NLToFrameServiceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return NLToFrameServiceResponse
func (client *Client) NLToFrameServiceWithOptions(request *NLToFrameServiceRequest, headers *NLToFrameServiceHeaders, runtime *util.RuntimeOptions) (_result *NLToFrameServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtensionStr)) {
		body["extensionStr"] = request.ExtensionStr
	}

	if !tea.BoolValue(util.IsUnset(request.IsNewModel)) {
		body["isNewModel"] = request.IsNewModel
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["modelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelName)) {
		body["modelName"] = request.ModelName
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
		Action:      tea.String("NLToFrameService"),
		Version:     tea.String("aiPaaS_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiPaaS/ai/nl2frame"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &NLToFrameServiceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过配置的指令，连接用户和系统，训练大模型
//
// @param request - NLToFrameServiceRequest
//
// @return NLToFrameServiceResponse
func (client *Client) NLToFrameService(request *NLToFrameServiceRequest) (_result *NLToFrameServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &NLToFrameServiceHeaders{}
	_result = &NLToFrameServiceResponse{}
	_body, _err := client.NLToFrameServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Baymax技能执行日志
//
// @param request - QueryBaymaxSkillLogRequest
//
// @param headers - QueryBaymaxSkillLogHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBaymaxSkillLogResponse
func (client *Client) QueryBaymaxSkillLogWithOptions(request *QueryBaymaxSkillLogRequest, headers *QueryBaymaxSkillLogHeaders, runtime *util.RuntimeOptions) (_result *QueryBaymaxSkillLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.From)) {
		query["from"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.LogLevel)) {
		query["logLevel"] = request.LogLevel
	}

	if !tea.BoolValue(util.IsUnset(request.SkillExecuteId)) {
		query["skillExecuteId"] = request.SkillExecuteId
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		query["to"] = request.To
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
		Action:      tea.String("QueryBaymaxSkillLog"),
		Version:     tea.String("aiPaaS_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiPaaS/skills/logs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBaymaxSkillLogResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Baymax技能执行日志
//
// @param request - QueryBaymaxSkillLogRequest
//
// @return QueryBaymaxSkillLogResponse
func (client *Client) QueryBaymaxSkillLog(request *QueryBaymaxSkillLogRequest) (_result *QueryBaymaxSkillLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryBaymaxSkillLogHeaders{}
	_result = &QueryBaymaxSkillLogResponse{}
	_body, _err := client.QueryBaymaxSkillLogWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询会话消息并以大模型友好的协议返回
//
// @param tmpReq - QueryConversationMessageForAIRequest
//
// @param headers - QueryConversationMessageForAIHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryConversationMessageForAIResponse
func (client *Client) QueryConversationMessageForAIWithOptions(cid *string, tmpReq *QueryConversationMessageForAIRequest, headers *QueryConversationMessageForAIHeaders, runtime *util.RuntimeOptions) (_result *QueryConversationMessageForAIResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryConversationMessageForAIShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.OpenMsgIds)) {
		request.OpenMsgIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OpenMsgIds, tea.String("openMsgIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenMsgIdsShrink)) {
		query["openMsgIds"] = request.OpenMsgIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RecentDays)) {
		query["recentDays"] = request.RecentDays
	}

	if !tea.BoolValue(util.IsUnset(request.RecentHours)) {
		query["recentHours"] = request.RecentHours
	}

	if !tea.BoolValue(util.IsUnset(request.RecentN)) {
		query["recentN"] = request.RecentN
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
		Action:      tea.String("QueryConversationMessageForAI"),
		Version:     tea.String("aiPaaS_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiPaaS/me/memory/im/" + tea.StringValue(cid) + "/messages"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryConversationMessageForAIResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询会话消息并以大模型友好的协议返回
//
// @param request - QueryConversationMessageForAIRequest
//
// @return QueryConversationMessageForAIResponse
func (client *Client) QueryConversationMessageForAI(cid *string, request *QueryConversationMessageForAIRequest) (_result *QueryConversationMessageForAIResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryConversationMessageForAIHeaders{}
	_result = &QueryConversationMessageForAIResponse{}
	_body, _err := client.QueryConversationMessageForAIWithOptions(cid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询记忆学习进度
//
// @param request - QueryMemoryLearningTaskRequest
//
// @param headers - QueryMemoryLearningTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMemoryLearningTaskResponse
func (client *Client) QueryMemoryLearningTaskWithOptions(request *QueryMemoryLearningTaskRequest, headers *QueryMemoryLearningTaskHeaders, runtime *util.RuntimeOptions) (_result *QueryMemoryLearningTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentCode)) {
		query["agentCode"] = request.AgentCode
	}

	if !tea.BoolValue(util.IsUnset(request.LearningCode)) {
		query["learningCode"] = request.LearningCode
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
		Action:      tea.String("QueryMemoryLearningTask"),
		Version:     tea.String("aiPaaS_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiPaaS/me/memory/learningTask/get"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMemoryLearningTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询记忆学习进度
//
// @param request - QueryMemoryLearningTaskRequest
//
// @return QueryMemoryLearningTaskResponse
func (client *Client) QueryMemoryLearningTask(request *QueryMemoryLearningTaskRequest) (_result *QueryMemoryLearningTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMemoryLearningTaskHeaders{}
	_result = &QueryMemoryLearningTaskResponse{}
	_body, _err := client.QueryMemoryLearningTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 中信金属智能配料任务结果
//
// @param request - SmartFormulaResultServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SmartFormulaResultServiceResponse
func (client *Client) SmartFormulaResultServiceWithOptions(request *SmartFormulaResultServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SmartFormulaResultServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SmartFormulaResultService"),
		Version:     tea.String("aiPaaS_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiPaaS/nl2x/smartFormulas/results/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SmartFormulaResultServiceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 中信金属智能配料任务结果
//
// @param request - SmartFormulaResultServiceRequest
//
// @return SmartFormulaResultServiceResponse
func (client *Client) SmartFormulaResultService(request *SmartFormulaResultServiceRequest) (_result *SmartFormulaResultServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SmartFormulaResultServiceResponse{}
	_body, _err := client.SmartFormulaResultServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 中信金属智能配料任务触发
//
// @param request - SmartFormulaTriggerServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SmartFormulaTriggerServiceResponse
func (client *Client) SmartFormulaTriggerServiceWithOptions(request *SmartFormulaTriggerServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SmartFormulaTriggerServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Request)) {
		body["request"] = request.Request
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SmartFormulaTriggerService"),
		Version:     tea.String("aiPaaS_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiPaaS/nl2x/smartFormulas/trigger"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SmartFormulaTriggerServiceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 中信金属智能配料任务触发
//
// @param request - SmartFormulaTriggerServiceRequest
//
// @return SmartFormulaTriggerServiceResponse
func (client *Client) SmartFormulaTriggerService(request *SmartFormulaTriggerServiceRequest) (_result *SmartFormulaTriggerServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SmartFormulaTriggerServiceResponse{}
	_body, _err := client.SmartFormulaTriggerServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询智能报价结果
//
// @param request - SmartQuoteBatchQueryResultServiceRequest
//
// @param headers - SmartQuoteBatchQueryResultServiceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SmartQuoteBatchQueryResultServiceResponse
func (client *Client) SmartQuoteBatchQueryResultServiceWithOptions(request *SmartQuoteBatchQueryResultServiceRequest, headers *SmartQuoteBatchQueryResultServiceHeaders, runtime *util.RuntimeOptions) (_result *SmartQuoteBatchQueryResultServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
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
		Action:      tea.String("SmartQuoteBatchQueryResultService"),
		Version:     tea.String("aiPaaS_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiPaaS/nl2x/smartQuotations/results/batchQuery"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SmartQuoteBatchQueryResultServiceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询智能报价结果
//
// @param request - SmartQuoteBatchQueryResultServiceRequest
//
// @return SmartQuoteBatchQueryResultServiceResponse
func (client *Client) SmartQuoteBatchQueryResultService(request *SmartQuoteBatchQueryResultServiceRequest) (_result *SmartQuoteBatchQueryResultServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SmartQuoteBatchQueryResultServiceHeaders{}
	_result = &SmartQuoteBatchQueryResultServiceResponse{}
	_body, _err := client.SmartQuoteBatchQueryResultServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询智能报价
//
// @param request - SmartQuoteBatchQueryServiceRequest
//
// @param headers - SmartQuoteBatchQueryServiceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SmartQuoteBatchQueryServiceResponse
func (client *Client) SmartQuoteBatchQueryServiceWithOptions(request *SmartQuoteBatchQueryServiceRequest, headers *SmartQuoteBatchQueryServiceHeaders, runtime *util.RuntimeOptions) (_result *SmartQuoteBatchQueryServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Request)) {
		body["request"] = request.Request
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
		Action:      tea.String("SmartQuoteBatchQueryService"),
		Version:     tea.String("aiPaaS_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiPaaS/nl2x/smartQuotations/batchQuery"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SmartQuoteBatchQueryServiceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询智能报价
//
// @param request - SmartQuoteBatchQueryServiceRequest
//
// @return SmartQuoteBatchQueryServiceResponse
func (client *Client) SmartQuoteBatchQueryService(request *SmartQuoteBatchQueryServiceRequest) (_result *SmartQuoteBatchQueryServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SmartQuoteBatchQueryServiceHeaders{}
	_result = &SmartQuoteBatchQueryServiceResponse{}
	_body, _err := client.SmartQuoteBatchQueryServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加智能报价数据
//
// @param request - SmartQuoteDataServiceRequest
//
// @param headers - SmartQuoteDataServiceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SmartQuoteDataServiceResponse
func (client *Client) SmartQuoteDataServiceWithOptions(request *SmartQuoteDataServiceRequest, headers *SmartQuoteDataServiceHeaders, runtime *util.RuntimeOptions) (_result *SmartQuoteDataServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Request)) {
		body["request"] = request.Request
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
		Action:      tea.String("SmartQuoteDataService"),
		Version:     tea.String("aiPaaS_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiPaaS/nl2x/smartQuotations/datas"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SmartQuoteDataServiceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加智能报价数据
//
// @param request - SmartQuoteDataServiceRequest
//
// @return SmartQuoteDataServiceResponse
func (client *Client) SmartQuoteDataService(request *SmartQuoteDataServiceRequest) (_result *SmartQuoteDataServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SmartQuoteDataServiceHeaders{}
	_result = &SmartQuoteDataServiceResponse{}
	_body, _err := client.SmartQuoteDataServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询智能报价结果
//
// @param request - SmartQuoteQueryResultServiceRequest
//
// @param headers - SmartQuoteQueryResultServiceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SmartQuoteQueryResultServiceResponse
func (client *Client) SmartQuoteQueryResultServiceWithOptions(request *SmartQuoteQueryResultServiceRequest, headers *SmartQuoteQueryResultServiceHeaders, runtime *util.RuntimeOptions) (_result *SmartQuoteQueryResultServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
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
		Action:      tea.String("SmartQuoteQueryResultService"),
		Version:     tea.String("aiPaaS_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiPaaS/nl2x/smartQuotations/results/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SmartQuoteQueryResultServiceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询智能报价结果
//
// @param request - SmartQuoteQueryResultServiceRequest
//
// @return SmartQuoteQueryResultServiceResponse
func (client *Client) SmartQuoteQueryResultService(request *SmartQuoteQueryResultServiceRequest) (_result *SmartQuoteQueryResultServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SmartQuoteQueryResultServiceHeaders{}
	_result = &SmartQuoteQueryResultServiceResponse{}
	_body, _err := client.SmartQuoteQueryResultServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询智能报价
//
// @param request - SmartQuoteQueryServiceRequest
//
// @param headers - SmartQuoteQueryServiceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SmartQuoteQueryServiceResponse
func (client *Client) SmartQuoteQueryServiceWithOptions(request *SmartQuoteQueryServiceRequest, headers *SmartQuoteQueryServiceHeaders, runtime *util.RuntimeOptions) (_result *SmartQuoteQueryServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Request)) {
		body["request"] = request.Request
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
		Action:      tea.String("SmartQuoteQueryService"),
		Version:     tea.String("aiPaaS_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiPaaS/nl2x/smartQuotations/triggerQuery"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SmartQuoteQueryServiceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询智能报价
//
// @param request - SmartQuoteQueryServiceRequest
//
// @return SmartQuoteQueryServiceResponse
func (client *Client) SmartQuoteQueryService(request *SmartQuoteQueryServiceRequest) (_result *SmartQuoteQueryServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SmartQuoteQueryServiceHeaders{}
	_result = &SmartQuoteQueryServiceResponse{}
	_body, _err := client.SmartQuoteQueryServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交记忆学习任务
//
// @param tmpReq - SubmitMemoryLearningTaskRequest
//
// @param headers - SubmitMemoryLearningTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitMemoryLearningTaskResponse
func (client *Client) SubmitMemoryLearningTaskWithOptions(tmpReq *SubmitMemoryLearningTaskRequest, headers *SubmitMemoryLearningTaskHeaders, runtime *util.RuntimeOptions) (_result *SubmitMemoryLearningTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SubmitMemoryLearningTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Content)) {
		request.ContentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Content, tea.String("content"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentCode)) {
		query["agentCode"] = request.AgentCode
	}

	if !tea.BoolValue(util.IsUnset(request.ContentShrink)) {
		query["content"] = request.ContentShrink
	}

	if !tea.BoolValue(util.IsUnset(request.LearningMode)) {
		query["learningMode"] = request.LearningMode
	}

	if !tea.BoolValue(util.IsUnset(request.MemoryKey)) {
		query["memoryKey"] = request.MemoryKey
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
		Action:      tea.String("SubmitMemoryLearningTask"),
		Version:     tea.String("aiPaaS_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiPaaS/me/memory/learningTask/put"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitMemoryLearningTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交记忆学习任务
//
// @param request - SubmitMemoryLearningTaskRequest
//
// @return SubmitMemoryLearningTaskResponse
func (client *Client) SubmitMemoryLearningTask(request *SubmitMemoryLearningTaskRequest) (_result *SubmitMemoryLearningTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SubmitMemoryLearningTaskHeaders{}
	_result = &SubmitMemoryLearningTaskResponse{}
	_body, _err := client.SubmitMemoryLearningTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
