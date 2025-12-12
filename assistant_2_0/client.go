// This file is auto-generated, don't edit it. Thanks.
package assistant_2_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
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

func (s *CreateAssistantMessageRequest) SetRole(v string) *CreateAssistantMessageRequest {
	s.Role = &v
	return s
}

type CreateAssistantMessageResponseBody struct {
	AssistantId *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	CreatedAt   *int64  `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Id          *string `json:"id,omitempty" xml:"id,omitempty"`
	Object      *string `json:"object,omitempty" xml:"object,omitempty"`
	Role        *string `json:"role,omitempty" xml:"role,omitempty"`
	RunId       *string `json:"runId,omitempty" xml:"runId,omitempty"`
	ThreadId    *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
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

func (s *CreateAssistantMessageResponseBody) SetCreatedAt(v int64) *CreateAssistantMessageResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateAssistantMessageResponseBody) SetId(v string) *CreateAssistantMessageResponseBody {
	s.Id = &v
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

type CreateAssistantRunResponseBody struct {
	AssistantId  *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	CancelledAt  *int64  `json:"cancelledAt,omitempty" xml:"cancelledAt,omitempty"`
	CompletedAt  *int64  `json:"completedAt,omitempty" xml:"completedAt,omitempty"`
	CreatedAt    *int64  `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	ExpiresAt    *int64  `json:"expiresAt,omitempty" xml:"expiresAt,omitempty"`
	FailedAt     *int64  `json:"failedAt,omitempty" xml:"failedAt,omitempty"`
	Id           *string `json:"id,omitempty" xml:"id,omitempty"`
	LastErrorMsg *string `json:"lastErrorMsg,omitempty" xml:"lastErrorMsg,omitempty"`
	Object       *string `json:"object,omitempty" xml:"object,omitempty"`
	StartedAt    *int64  `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
	ThreadId     *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
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

type CreateAssistantThreadResponseBody struct {
	CreatedAt *int64  `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Id        *string `json:"id,omitempty" xml:"id,omitempty"`
	Object    *string `json:"object,omitempty" xml:"object,omitempty"`
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
	AssistantId *string       `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	Content     []interface{} `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	CreatedAt   *int64        `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Id          *string       `json:"id,omitempty" xml:"id,omitempty"`
	Object      *string       `json:"object,omitempty" xml:"object,omitempty"`
	Role        *string       `json:"role,omitempty" xml:"role,omitempty"`
	RunId       *string       `json:"runId,omitempty" xml:"runId,omitempty"`
	ThreadId    *string       `json:"threadId,omitempty" xml:"threadId,omitempty"`
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
	AssisantId *string       `json:"assisantId,omitempty" xml:"assisantId,omitempty"`
	Content    []interface{} `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	CreatedAt  *int64        `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Id         *string       `json:"id,omitempty" xml:"id,omitempty"`
	Object     *string       `json:"object,omitempty" xml:"object,omitempty"`
	Role       *string       `json:"role,omitempty" xml:"role,omitempty"`
	RunId      *string       `json:"runId,omitempty" xml:"runId,omitempty"`
	ThreadId   *string       `json:"threadId,omitempty" xml:"threadId,omitempty"`
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
	CreatedAt *int64  `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Id        *string `json:"id,omitempty" xml:"id,omitempty"`
	Object    *string `json:"object,omitempty" xml:"object,omitempty"`
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		query["role"] = request.Role
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
		Action:      tea.String("CreateAssistantMessage"),
		Version:     tea.String("assistant_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/assistant/threads/" + tea.StringValue(threadId) + "/messages"),
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
// @param headers - CreateAssistantRunHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAssistantRunResponse
func (client *Client) CreateAssistantRunWithOptions(threadId *string, headers *CreateAssistantRunHeaders, runtime *util.RuntimeOptions) (_result *CreateAssistantRunResponse, _err error) {
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
		Action:      tea.String("CreateAssistantRun"),
		Version:     tea.String("assistant_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/assistant/threads/" + tea.StringValue(threadId) + "/runs"),
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
// @return CreateAssistantRunResponse
func (client *Client) CreateAssistantRun(threadId *string) (_result *CreateAssistantRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateAssistantRunHeaders{}
	_result = &CreateAssistantRunResponse{}
	_body, _err := client.CreateAssistantRunWithOptions(threadId, headers, runtime)
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
// @param headers - CreateAssistantThreadHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAssistantThreadResponse
func (client *Client) CreateAssistantThreadWithOptions(headers *CreateAssistantThreadHeaders, runtime *util.RuntimeOptions) (_result *CreateAssistantThreadResponse, _err error) {
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
		Action:      tea.String("CreateAssistantThread"),
		Version:     tea.String("assistant_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/assistant/threads"),
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
// @return CreateAssistantThreadResponse
func (client *Client) CreateAssistantThread() (_result *CreateAssistantThreadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateAssistantThreadHeaders{}
	_result = &CreateAssistantThreadResponse{}
	_body, _err := client.CreateAssistantThreadWithOptions(headers, runtime)
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
		Version:     tea.String("assistant_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/assistant/threads/" + tea.StringValue(threadId) + "/messages/" + tea.StringValue(messageId)),
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
		Version:     tea.String("assistant_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/assistant/threads/" + tea.StringValue(threadId)),
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
		Version:     tea.String("assistant_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/assistant/threads/" + tea.StringValue(threadId) + "/messages"),
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
		Version:     tea.String("assistant_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/assistant/threads/" + tea.StringValue(threadId) + "/messages/" + tea.StringValue(messageId)),
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
		Version:     tea.String("assistant_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/assistant/threads/" + tea.StringValue(threadId)),
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
