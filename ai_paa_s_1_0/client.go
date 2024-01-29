// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package ai_paa_s_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	AgentCode          *string                    `json:"agentCode,omitempty" xml:"agentCode,omitempty"`
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
	Module *string `json:"module,omitempty" xml:"module,omitempty"`
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
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
	Result  []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success *bool                    `json:"success,omitempty" xml:"success,omitempty"`
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
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	Module  *string `json:"module,omitempty" xml:"module,omitempty"`
	Prompt  *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	UserId  *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
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
	Module     *string            `json:"module,omitempty" xml:"module,omitempty"`
	Number     *int64             `json:"number,omitempty" xml:"number,omitempty"`
	Parameters map[string]*string `json:"parameters,omitempty" xml:"parameters,omitempty"`
	Prompt     *string            `json:"prompt,omitempty" xml:"prompt,omitempty"`
	UserId     *string            `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Result  *LiandanluTextToImageModelResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                        `json:"success,omitempty" xml:"success,omitempty"`
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
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TaskId     *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
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
	From           *int32  `json:"from,omitempty" xml:"from,omitempty"`
	LogLevel       *string `json:"logLevel,omitempty" xml:"logLevel,omitempty"`
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
	AgentCode    *string `json:"agentCode,omitempty" xml:"agentCode,omitempty"`
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
	AgentCode    *string                                 `json:"agentCode,omitempty" xml:"agentCode,omitempty"`
	Content      *SubmitMemoryLearningTaskRequestContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	LearningMode *string                                 `json:"learningMode,omitempty" xml:"learningMode,omitempty"`
	MemoryKey    *string                                 `json:"memoryKey,omitempty" xml:"memoryKey,omitempty"`
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
	AgentCode     *string `json:"agentCode,omitempty" xml:"agentCode,omitempty"`
	ContentShrink *string `json:"content,omitempty" xml:"content,omitempty"`
	LearningMode  *string `json:"learningMode,omitempty" xml:"learningMode,omitempty"`
	MemoryKey     *string `json:"memoryKey,omitempty" xml:"memoryKey,omitempty"`
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
	interfaceSPI, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = interfaceSPI
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

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
