// This file is auto-generated, don't edit it. Thanks.
package todo_e_e_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AppCreateEnterpriseTodoTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AppCreateEnterpriseTodoTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s AppCreateEnterpriseTodoTaskHeaders) GoString() string {
	return s.String()
}

func (s *AppCreateEnterpriseTodoTaskHeaders) SetCommonHeaders(v map[string]*string) *AppCreateEnterpriseTodoTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AppCreateEnterpriseTodoTaskHeaders) SetXAcsDingtalkAccessToken(v string) *AppCreateEnterpriseTodoTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AppCreateEnterpriseTodoTaskRequest struct {
	BizCategoryId      *string                                           `json:"bizCategoryId,omitempty" xml:"bizCategoryId,omitempty"`
	BizCreatedTime     *int64                                            `json:"bizCreatedTime,omitempty" xml:"bizCreatedTime,omitempty"`
	CustomFields       []*AppCreateEnterpriseTodoTaskRequestCustomFields `json:"customFields,omitempty" xml:"customFields,omitempty" type:"Repeated"`
	Description        *string                                           `json:"description,omitempty" xml:"description,omitempty"`
	DetailUrl          *AppCreateEnterpriseTodoTaskRequestDetailUrl      `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	DueTime            *int64                                            `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	ExecutorIds        []*string                                         `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	NotifyConfigs      *AppCreateEnterpriseTodoTaskRequestNotifyConfigs  `json:"notifyConfigs,omitempty" xml:"notifyConfigs,omitempty" type:"Struct"`
	OperatorId         *string                                           `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	Priority           *int32                                            `json:"priority,omitempty" xml:"priority,omitempty"`
	SourceId           *string                                           `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	SourceTitle        *string                                           `json:"sourceTitle,omitempty" xml:"sourceTitle,omitempty"`
	Subject            *string                                           `json:"subject,omitempty" xml:"subject,omitempty"`
	ToolbarTemplateKey *string                                           `json:"toolbarTemplateKey,omitempty" xml:"toolbarTemplateKey,omitempty"`
	Type               *string                                           `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AppCreateEnterpriseTodoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s AppCreateEnterpriseTodoTaskRequest) GoString() string {
	return s.String()
}

func (s *AppCreateEnterpriseTodoTaskRequest) SetBizCategoryId(v string) *AppCreateEnterpriseTodoTaskRequest {
	s.BizCategoryId = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskRequest) SetBizCreatedTime(v int64) *AppCreateEnterpriseTodoTaskRequest {
	s.BizCreatedTime = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskRequest) SetCustomFields(v []*AppCreateEnterpriseTodoTaskRequestCustomFields) *AppCreateEnterpriseTodoTaskRequest {
	s.CustomFields = v
	return s
}

func (s *AppCreateEnterpriseTodoTaskRequest) SetDescription(v string) *AppCreateEnterpriseTodoTaskRequest {
	s.Description = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskRequest) SetDetailUrl(v *AppCreateEnterpriseTodoTaskRequestDetailUrl) *AppCreateEnterpriseTodoTaskRequest {
	s.DetailUrl = v
	return s
}

func (s *AppCreateEnterpriseTodoTaskRequest) SetDueTime(v int64) *AppCreateEnterpriseTodoTaskRequest {
	s.DueTime = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskRequest) SetExecutorIds(v []*string) *AppCreateEnterpriseTodoTaskRequest {
	s.ExecutorIds = v
	return s
}

func (s *AppCreateEnterpriseTodoTaskRequest) SetNotifyConfigs(v *AppCreateEnterpriseTodoTaskRequestNotifyConfigs) *AppCreateEnterpriseTodoTaskRequest {
	s.NotifyConfigs = v
	return s
}

func (s *AppCreateEnterpriseTodoTaskRequest) SetOperatorId(v string) *AppCreateEnterpriseTodoTaskRequest {
	s.OperatorId = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskRequest) SetPriority(v int32) *AppCreateEnterpriseTodoTaskRequest {
	s.Priority = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskRequest) SetSourceId(v string) *AppCreateEnterpriseTodoTaskRequest {
	s.SourceId = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskRequest) SetSourceTitle(v string) *AppCreateEnterpriseTodoTaskRequest {
	s.SourceTitle = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskRequest) SetSubject(v string) *AppCreateEnterpriseTodoTaskRequest {
	s.Subject = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskRequest) SetToolbarTemplateKey(v string) *AppCreateEnterpriseTodoTaskRequest {
	s.ToolbarTemplateKey = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskRequest) SetType(v string) *AppCreateEnterpriseTodoTaskRequest {
	s.Type = &v
	return s
}

type AppCreateEnterpriseTodoTaskRequestCustomFields struct {
	FieldKey   *string `json:"fieldKey,omitempty" xml:"fieldKey,omitempty"`
	FieldLink  *string `json:"fieldLink,omitempty" xml:"fieldLink,omitempty"`
	FieldType  *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
	Icon       *string `json:"icon,omitempty" xml:"icon,omitempty"`
}

func (s AppCreateEnterpriseTodoTaskRequestCustomFields) String() string {
	return tea.Prettify(s)
}

func (s AppCreateEnterpriseTodoTaskRequestCustomFields) GoString() string {
	return s.String()
}

func (s *AppCreateEnterpriseTodoTaskRequestCustomFields) SetFieldKey(v string) *AppCreateEnterpriseTodoTaskRequestCustomFields {
	s.FieldKey = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskRequestCustomFields) SetFieldLink(v string) *AppCreateEnterpriseTodoTaskRequestCustomFields {
	s.FieldLink = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskRequestCustomFields) SetFieldType(v string) *AppCreateEnterpriseTodoTaskRequestCustomFields {
	s.FieldType = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskRequestCustomFields) SetFieldValue(v string) *AppCreateEnterpriseTodoTaskRequestCustomFields {
	s.FieldValue = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskRequestCustomFields) SetIcon(v string) *AppCreateEnterpriseTodoTaskRequestCustomFields {
	s.Icon = &v
	return s
}

type AppCreateEnterpriseTodoTaskRequestDetailUrl struct {
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
	WebUrl *string `json:"webUrl,omitempty" xml:"webUrl,omitempty"`
}

func (s AppCreateEnterpriseTodoTaskRequestDetailUrl) String() string {
	return tea.Prettify(s)
}

func (s AppCreateEnterpriseTodoTaskRequestDetailUrl) GoString() string {
	return s.String()
}

func (s *AppCreateEnterpriseTodoTaskRequestDetailUrl) SetAppUrl(v string) *AppCreateEnterpriseTodoTaskRequestDetailUrl {
	s.AppUrl = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskRequestDetailUrl) SetWebUrl(v string) *AppCreateEnterpriseTodoTaskRequestDetailUrl {
	s.WebUrl = &v
	return s
}

type AppCreateEnterpriseTodoTaskRequestNotifyConfigs struct {
	Assistance *string `json:"assistance,omitempty" xml:"assistance,omitempty"`
	DingNotify *string `json:"dingNotify,omitempty" xml:"dingNotify,omitempty"`
}

func (s AppCreateEnterpriseTodoTaskRequestNotifyConfigs) String() string {
	return tea.Prettify(s)
}

func (s AppCreateEnterpriseTodoTaskRequestNotifyConfigs) GoString() string {
	return s.String()
}

func (s *AppCreateEnterpriseTodoTaskRequestNotifyConfigs) SetAssistance(v string) *AppCreateEnterpriseTodoTaskRequestNotifyConfigs {
	s.Assistance = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskRequestNotifyConfigs) SetDingNotify(v string) *AppCreateEnterpriseTodoTaskRequestNotifyConfigs {
	s.DingNotify = &v
	return s
}

type AppCreateEnterpriseTodoTaskResponseBody struct {
	BizCategoryId *string                                           `json:"bizCategoryId,omitempty" xml:"bizCategoryId,omitempty"`
	CreatedTime   *int64                                            `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	CreatorId     *string                                           `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description   *string                                           `json:"description,omitempty" xml:"description,omitempty"`
	DetailUrl     *AppCreateEnterpriseTodoTaskResponseBodyDetailUrl `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	Done          *bool                                             `json:"done,omitempty" xml:"done,omitempty"`
	DueTime       *int64                                            `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	ExecutorIds   []*string                                         `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	ModifiedTime  *int64                                            `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	Priority      *int32                                            `json:"priority,omitempty" xml:"priority,omitempty"`
	SourceId      *string                                           `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	Subject       *string                                           `json:"subject,omitempty" xml:"subject,omitempty"`
	TaskId        *string                                           `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s AppCreateEnterpriseTodoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppCreateEnterpriseTodoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *AppCreateEnterpriseTodoTaskResponseBody) SetBizCategoryId(v string) *AppCreateEnterpriseTodoTaskResponseBody {
	s.BizCategoryId = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskResponseBody) SetCreatedTime(v int64) *AppCreateEnterpriseTodoTaskResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskResponseBody) SetCreatorId(v string) *AppCreateEnterpriseTodoTaskResponseBody {
	s.CreatorId = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskResponseBody) SetDescription(v string) *AppCreateEnterpriseTodoTaskResponseBody {
	s.Description = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskResponseBody) SetDetailUrl(v *AppCreateEnterpriseTodoTaskResponseBodyDetailUrl) *AppCreateEnterpriseTodoTaskResponseBody {
	s.DetailUrl = v
	return s
}

func (s *AppCreateEnterpriseTodoTaskResponseBody) SetDone(v bool) *AppCreateEnterpriseTodoTaskResponseBody {
	s.Done = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskResponseBody) SetDueTime(v int64) *AppCreateEnterpriseTodoTaskResponseBody {
	s.DueTime = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskResponseBody) SetExecutorIds(v []*string) *AppCreateEnterpriseTodoTaskResponseBody {
	s.ExecutorIds = v
	return s
}

func (s *AppCreateEnterpriseTodoTaskResponseBody) SetModifiedTime(v int64) *AppCreateEnterpriseTodoTaskResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskResponseBody) SetPriority(v int32) *AppCreateEnterpriseTodoTaskResponseBody {
	s.Priority = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskResponseBody) SetSourceId(v string) *AppCreateEnterpriseTodoTaskResponseBody {
	s.SourceId = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskResponseBody) SetSubject(v string) *AppCreateEnterpriseTodoTaskResponseBody {
	s.Subject = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskResponseBody) SetTaskId(v string) *AppCreateEnterpriseTodoTaskResponseBody {
	s.TaskId = &v
	return s
}

type AppCreateEnterpriseTodoTaskResponseBodyDetailUrl struct {
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
	WebUrl *string `json:"webUrl,omitempty" xml:"webUrl,omitempty"`
}

func (s AppCreateEnterpriseTodoTaskResponseBodyDetailUrl) String() string {
	return tea.Prettify(s)
}

func (s AppCreateEnterpriseTodoTaskResponseBodyDetailUrl) GoString() string {
	return s.String()
}

func (s *AppCreateEnterpriseTodoTaskResponseBodyDetailUrl) SetAppUrl(v string) *AppCreateEnterpriseTodoTaskResponseBodyDetailUrl {
	s.AppUrl = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskResponseBodyDetailUrl) SetWebUrl(v string) *AppCreateEnterpriseTodoTaskResponseBodyDetailUrl {
	s.WebUrl = &v
	return s
}

type AppCreateEnterpriseTodoTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AppCreateEnterpriseTodoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AppCreateEnterpriseTodoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s AppCreateEnterpriseTodoTaskResponse) GoString() string {
	return s.String()
}

func (s *AppCreateEnterpriseTodoTaskResponse) SetHeaders(v map[string]*string) *AppCreateEnterpriseTodoTaskResponse {
	s.Headers = v
	return s
}

func (s *AppCreateEnterpriseTodoTaskResponse) SetStatusCode(v int32) *AppCreateEnterpriseTodoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *AppCreateEnterpriseTodoTaskResponse) SetBody(v *AppCreateEnterpriseTodoTaskResponseBody) *AppCreateEnterpriseTodoTaskResponse {
	s.Body = v
	return s
}

type AppDeleteTodoEETaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AppDeleteTodoEETaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s AppDeleteTodoEETaskHeaders) GoString() string {
	return s.String()
}

func (s *AppDeleteTodoEETaskHeaders) SetCommonHeaders(v map[string]*string) *AppDeleteTodoEETaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AppDeleteTodoEETaskHeaders) SetXAcsDingtalkAccessToken(v string) *AppDeleteTodoEETaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AppDeleteTodoEETaskRequest struct {
	OperatorId *string   `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	TaskIds    []*string `json:"taskIds,omitempty" xml:"taskIds,omitempty" type:"Repeated"`
}

func (s AppDeleteTodoEETaskRequest) String() string {
	return tea.Prettify(s)
}

func (s AppDeleteTodoEETaskRequest) GoString() string {
	return s.String()
}

func (s *AppDeleteTodoEETaskRequest) SetOperatorId(v string) *AppDeleteTodoEETaskRequest {
	s.OperatorId = &v
	return s
}

func (s *AppDeleteTodoEETaskRequest) SetTaskIds(v []*string) *AppDeleteTodoEETaskRequest {
	s.TaskIds = v
	return s
}

type AppDeleteTodoEETaskResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AppDeleteTodoEETaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppDeleteTodoEETaskResponseBody) GoString() string {
	return s.String()
}

func (s *AppDeleteTodoEETaskResponseBody) SetSuccess(v bool) *AppDeleteTodoEETaskResponseBody {
	s.Success = &v
	return s
}

type AppDeleteTodoEETaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AppDeleteTodoEETaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AppDeleteTodoEETaskResponse) String() string {
	return tea.Prettify(s)
}

func (s AppDeleteTodoEETaskResponse) GoString() string {
	return s.String()
}

func (s *AppDeleteTodoEETaskResponse) SetHeaders(v map[string]*string) *AppDeleteTodoEETaskResponse {
	s.Headers = v
	return s
}

func (s *AppDeleteTodoEETaskResponse) SetStatusCode(v int32) *AppDeleteTodoEETaskResponse {
	s.StatusCode = &v
	return s
}

func (s *AppDeleteTodoEETaskResponse) SetBody(v *AppDeleteTodoEETaskResponseBody) *AppDeleteTodoEETaskResponse {
	s.Body = v
	return s
}

type AppGetUserTaskListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AppGetUserTaskListHeaders) String() string {
	return tea.Prettify(s)
}

func (s AppGetUserTaskListHeaders) GoString() string {
	return s.String()
}

func (s *AppGetUserTaskListHeaders) SetCommonHeaders(v map[string]*string) *AppGetUserTaskListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AppGetUserTaskListHeaders) SetXAcsDingtalkAccessToken(v string) *AppGetUserTaskListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AppGetUserTaskListRequest struct {
	Done       *bool   `json:"done,omitempty" xml:"done,omitempty"`
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AppGetUserTaskListRequest) String() string {
	return tea.Prettify(s)
}

func (s AppGetUserTaskListRequest) GoString() string {
	return s.String()
}

func (s *AppGetUserTaskListRequest) SetDone(v bool) *AppGetUserTaskListRequest {
	s.Done = &v
	return s
}

func (s *AppGetUserTaskListRequest) SetOperatorId(v string) *AppGetUserTaskListRequest {
	s.OperatorId = &v
	return s
}

func (s *AppGetUserTaskListRequest) SetPageNumber(v int32) *AppGetUserTaskListRequest {
	s.PageNumber = &v
	return s
}

func (s *AppGetUserTaskListRequest) SetPageSize(v int32) *AppGetUserTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *AppGetUserTaskListRequest) SetType(v string) *AppGetUserTaskListRequest {
	s.Type = &v
	return s
}

type AppGetUserTaskListResponseBody struct {
	Data       []*AppGetUserTaskListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	HasMore    *bool                                 `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	PageNumber *int32                                `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32                                `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	TotalCount *int32                                `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s AppGetUserTaskListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppGetUserTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *AppGetUserTaskListResponseBody) SetData(v []*AppGetUserTaskListResponseBodyData) *AppGetUserTaskListResponseBody {
	s.Data = v
	return s
}

func (s *AppGetUserTaskListResponseBody) SetHasMore(v bool) *AppGetUserTaskListResponseBody {
	s.HasMore = &v
	return s
}

func (s *AppGetUserTaskListResponseBody) SetPageNumber(v int32) *AppGetUserTaskListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *AppGetUserTaskListResponseBody) SetPageSize(v int32) *AppGetUserTaskListResponseBody {
	s.PageSize = &v
	return s
}

func (s *AppGetUserTaskListResponseBody) SetTotalCount(v int32) *AppGetUserTaskListResponseBody {
	s.TotalCount = &v
	return s
}

type AppGetUserTaskListResponseBodyData struct {
	BizCategoryId *string                                      `json:"bizCategoryId,omitempty" xml:"bizCategoryId,omitempty"`
	CreatedTime   *int64                                       `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Description   *string                                      `json:"description,omitempty" xml:"description,omitempty"`
	DetailUrl     *AppGetUserTaskListResponseBodyDataDetailUrl `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	Done          *bool                                        `json:"done,omitempty" xml:"done,omitempty"`
	DueTime       *int64                                       `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	ModifiedTime  *int64                                       `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	OperatorId    *string                                      `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	Priority      *int64                                       `json:"priority,omitempty" xml:"priority,omitempty"`
	Subject       *string                                      `json:"subject,omitempty" xml:"subject,omitempty"`
	TaskId        *string                                      `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s AppGetUserTaskListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AppGetUserTaskListResponseBodyData) GoString() string {
	return s.String()
}

func (s *AppGetUserTaskListResponseBodyData) SetBizCategoryId(v string) *AppGetUserTaskListResponseBodyData {
	s.BizCategoryId = &v
	return s
}

func (s *AppGetUserTaskListResponseBodyData) SetCreatedTime(v int64) *AppGetUserTaskListResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *AppGetUserTaskListResponseBodyData) SetDescription(v string) *AppGetUserTaskListResponseBodyData {
	s.Description = &v
	return s
}

func (s *AppGetUserTaskListResponseBodyData) SetDetailUrl(v *AppGetUserTaskListResponseBodyDataDetailUrl) *AppGetUserTaskListResponseBodyData {
	s.DetailUrl = v
	return s
}

func (s *AppGetUserTaskListResponseBodyData) SetDone(v bool) *AppGetUserTaskListResponseBodyData {
	s.Done = &v
	return s
}

func (s *AppGetUserTaskListResponseBodyData) SetDueTime(v int64) *AppGetUserTaskListResponseBodyData {
	s.DueTime = &v
	return s
}

func (s *AppGetUserTaskListResponseBodyData) SetModifiedTime(v int64) *AppGetUserTaskListResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *AppGetUserTaskListResponseBodyData) SetOperatorId(v string) *AppGetUserTaskListResponseBodyData {
	s.OperatorId = &v
	return s
}

func (s *AppGetUserTaskListResponseBodyData) SetPriority(v int64) *AppGetUserTaskListResponseBodyData {
	s.Priority = &v
	return s
}

func (s *AppGetUserTaskListResponseBodyData) SetSubject(v string) *AppGetUserTaskListResponseBodyData {
	s.Subject = &v
	return s
}

func (s *AppGetUserTaskListResponseBodyData) SetTaskId(v string) *AppGetUserTaskListResponseBodyData {
	s.TaskId = &v
	return s
}

type AppGetUserTaskListResponseBodyDataDetailUrl struct {
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
	WebUrl *string `json:"webUrl,omitempty" xml:"webUrl,omitempty"`
}

func (s AppGetUserTaskListResponseBodyDataDetailUrl) String() string {
	return tea.Prettify(s)
}

func (s AppGetUserTaskListResponseBodyDataDetailUrl) GoString() string {
	return s.String()
}

func (s *AppGetUserTaskListResponseBodyDataDetailUrl) SetAppUrl(v string) *AppGetUserTaskListResponseBodyDataDetailUrl {
	s.AppUrl = &v
	return s
}

func (s *AppGetUserTaskListResponseBodyDataDetailUrl) SetWebUrl(v string) *AppGetUserTaskListResponseBodyDataDetailUrl {
	s.WebUrl = &v
	return s
}

type AppGetUserTaskListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AppGetUserTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AppGetUserTaskListResponse) String() string {
	return tea.Prettify(s)
}

func (s AppGetUserTaskListResponse) GoString() string {
	return s.String()
}

func (s *AppGetUserTaskListResponse) SetHeaders(v map[string]*string) *AppGetUserTaskListResponse {
	s.Headers = v
	return s
}

func (s *AppGetUserTaskListResponse) SetStatusCode(v int32) *AppGetUserTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *AppGetUserTaskListResponse) SetBody(v *AppGetUserTaskListResponseBody) *AppGetUserTaskListResponse {
	s.Body = v
	return s
}

type AppUpdateTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AppUpdateTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s AppUpdateTaskHeaders) GoString() string {
	return s.String()
}

func (s *AppUpdateTaskHeaders) SetCommonHeaders(v map[string]*string) *AppUpdateTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AppUpdateTaskHeaders) SetXAcsDingtalkAccessToken(v string) *AppUpdateTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AppUpdateTaskRequest struct {
	BizCreatedTime     *int64    `json:"bizCreatedTime,omitempty" xml:"bizCreatedTime,omitempty"`
	Description        *string   `json:"description,omitempty" xml:"description,omitempty"`
	Done               *bool     `json:"done,omitempty" xml:"done,omitempty"`
	DueTime            *int64    `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	ExecutorIds        []*string `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	OperatorId         *string   `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	Subject            *string   `json:"subject,omitempty" xml:"subject,omitempty"`
	TaskId             *int64    `json:"taskId,omitempty" xml:"taskId,omitempty"`
	ToolbarTemplateKey *string   `json:"toolbarTemplateKey,omitempty" xml:"toolbarTemplateKey,omitempty"`
}

func (s AppUpdateTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s AppUpdateTaskRequest) GoString() string {
	return s.String()
}

func (s *AppUpdateTaskRequest) SetBizCreatedTime(v int64) *AppUpdateTaskRequest {
	s.BizCreatedTime = &v
	return s
}

func (s *AppUpdateTaskRequest) SetDescription(v string) *AppUpdateTaskRequest {
	s.Description = &v
	return s
}

func (s *AppUpdateTaskRequest) SetDone(v bool) *AppUpdateTaskRequest {
	s.Done = &v
	return s
}

func (s *AppUpdateTaskRequest) SetDueTime(v int64) *AppUpdateTaskRequest {
	s.DueTime = &v
	return s
}

func (s *AppUpdateTaskRequest) SetExecutorIds(v []*string) *AppUpdateTaskRequest {
	s.ExecutorIds = v
	return s
}

func (s *AppUpdateTaskRequest) SetOperatorId(v string) *AppUpdateTaskRequest {
	s.OperatorId = &v
	return s
}

func (s *AppUpdateTaskRequest) SetSubject(v string) *AppUpdateTaskRequest {
	s.Subject = &v
	return s
}

func (s *AppUpdateTaskRequest) SetTaskId(v int64) *AppUpdateTaskRequest {
	s.TaskId = &v
	return s
}

func (s *AppUpdateTaskRequest) SetToolbarTemplateKey(v string) *AppUpdateTaskRequest {
	s.ToolbarTemplateKey = &v
	return s
}

type AppUpdateTaskResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AppUpdateTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppUpdateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *AppUpdateTaskResponseBody) SetSuccess(v bool) *AppUpdateTaskResponseBody {
	s.Success = &v
	return s
}

type AppUpdateTaskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AppUpdateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AppUpdateTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s AppUpdateTaskResponse) GoString() string {
	return s.String()
}

func (s *AppUpdateTaskResponse) SetHeaders(v map[string]*string) *AppUpdateTaskResponse {
	s.Headers = v
	return s
}

func (s *AppUpdateTaskResponse) SetStatusCode(v int32) *AppUpdateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *AppUpdateTaskResponse) SetBody(v *AppUpdateTaskResponseBody) *AppUpdateTaskResponse {
	s.Body = v
	return s
}

type AppUpdateUserTaskStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AppUpdateUserTaskStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s AppUpdateUserTaskStatusHeaders) GoString() string {
	return s.String()
}

func (s *AppUpdateUserTaskStatusHeaders) SetCommonHeaders(v map[string]*string) *AppUpdateUserTaskStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AppUpdateUserTaskStatusHeaders) SetXAcsDingtalkAccessToken(v string) *AppUpdateUserTaskStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AppUpdateUserTaskStatusRequest struct {
	OperatorId       *string                                           `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	UserTaskStatuses []*AppUpdateUserTaskStatusRequestUserTaskStatuses `json:"userTaskStatuses,omitempty" xml:"userTaskStatuses,omitempty" type:"Repeated"`
}

func (s AppUpdateUserTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s AppUpdateUserTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *AppUpdateUserTaskStatusRequest) SetOperatorId(v string) *AppUpdateUserTaskStatusRequest {
	s.OperatorId = &v
	return s
}

func (s *AppUpdateUserTaskStatusRequest) SetUserTaskStatuses(v []*AppUpdateUserTaskStatusRequestUserTaskStatuses) *AppUpdateUserTaskStatusRequest {
	s.UserTaskStatuses = v
	return s
}

type AppUpdateUserTaskStatusRequestUserTaskStatuses struct {
	Done *bool `json:"done,omitempty" xml:"done,omitempty"`
	// if can be null:
	// false
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s AppUpdateUserTaskStatusRequestUserTaskStatuses) String() string {
	return tea.Prettify(s)
}

func (s AppUpdateUserTaskStatusRequestUserTaskStatuses) GoString() string {
	return s.String()
}

func (s *AppUpdateUserTaskStatusRequestUserTaskStatuses) SetDone(v bool) *AppUpdateUserTaskStatusRequestUserTaskStatuses {
	s.Done = &v
	return s
}

func (s *AppUpdateUserTaskStatusRequestUserTaskStatuses) SetTaskId(v string) *AppUpdateUserTaskStatusRequestUserTaskStatuses {
	s.TaskId = &v
	return s
}

type AppUpdateUserTaskStatusResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AppUpdateUserTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppUpdateUserTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *AppUpdateUserTaskStatusResponseBody) SetSuccess(v bool) *AppUpdateUserTaskStatusResponseBody {
	s.Success = &v
	return s
}

type AppUpdateUserTaskStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AppUpdateUserTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AppUpdateUserTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s AppUpdateUserTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *AppUpdateUserTaskStatusResponse) SetHeaders(v map[string]*string) *AppUpdateUserTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *AppUpdateUserTaskStatusResponse) SetStatusCode(v int32) *AppUpdateUserTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *AppUpdateUserTaskStatusResponse) SetBody(v *AppUpdateUserTaskStatusResponseBody) *AppUpdateUserTaskStatusResponse {
	s.Body = v
	return s
}

type CreateEnterpriseTodoTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateEnterpriseTodoTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseTodoTaskHeaders) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseTodoTaskHeaders) SetCommonHeaders(v map[string]*string) *CreateEnterpriseTodoTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateEnterpriseTodoTaskHeaders) SetXAcsDingtalkAccessToken(v string) *CreateEnterpriseTodoTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateEnterpriseTodoTaskRequest struct {
	BizCategoryId *string                                        `json:"bizCategoryId,omitempty" xml:"bizCategoryId,omitempty"`
	CustomFields  []*CreateEnterpriseTodoTaskRequestCustomFields `json:"customFields,omitempty" xml:"customFields,omitempty" type:"Repeated"`
	Description   *string                                        `json:"description,omitempty" xml:"description,omitempty"`
	DetailUrl     *CreateEnterpriseTodoTaskRequestDetailUrl      `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	DueTime       *int64                                         `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	ExecutorIds   []*string                                      `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	NotifyConfigs *CreateEnterpriseTodoTaskRequestNotifyConfigs  `json:"notifyConfigs,omitempty" xml:"notifyConfigs,omitempty" type:"Struct"`
	OperatorId    *string                                        `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	Priority      *int32                                         `json:"priority,omitempty" xml:"priority,omitempty"`
	SourceId      *string                                        `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	SourceTitle   *string                                        `json:"sourceTitle,omitempty" xml:"sourceTitle,omitempty"`
	Subject       *string                                        `json:"subject,omitempty" xml:"subject,omitempty"`
	TrackerIds    []*string                                      `json:"trackerIds,omitempty" xml:"trackerIds,omitempty" type:"Repeated"`
	Type          *string                                        `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateEnterpriseTodoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseTodoTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseTodoTaskRequest) SetBizCategoryId(v string) *CreateEnterpriseTodoTaskRequest {
	s.BizCategoryId = &v
	return s
}

func (s *CreateEnterpriseTodoTaskRequest) SetCustomFields(v []*CreateEnterpriseTodoTaskRequestCustomFields) *CreateEnterpriseTodoTaskRequest {
	s.CustomFields = v
	return s
}

func (s *CreateEnterpriseTodoTaskRequest) SetDescription(v string) *CreateEnterpriseTodoTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateEnterpriseTodoTaskRequest) SetDetailUrl(v *CreateEnterpriseTodoTaskRequestDetailUrl) *CreateEnterpriseTodoTaskRequest {
	s.DetailUrl = v
	return s
}

func (s *CreateEnterpriseTodoTaskRequest) SetDueTime(v int64) *CreateEnterpriseTodoTaskRequest {
	s.DueTime = &v
	return s
}

func (s *CreateEnterpriseTodoTaskRequest) SetExecutorIds(v []*string) *CreateEnterpriseTodoTaskRequest {
	s.ExecutorIds = v
	return s
}

func (s *CreateEnterpriseTodoTaskRequest) SetNotifyConfigs(v *CreateEnterpriseTodoTaskRequestNotifyConfigs) *CreateEnterpriseTodoTaskRequest {
	s.NotifyConfigs = v
	return s
}

func (s *CreateEnterpriseTodoTaskRequest) SetOperatorId(v string) *CreateEnterpriseTodoTaskRequest {
	s.OperatorId = &v
	return s
}

func (s *CreateEnterpriseTodoTaskRequest) SetPriority(v int32) *CreateEnterpriseTodoTaskRequest {
	s.Priority = &v
	return s
}

func (s *CreateEnterpriseTodoTaskRequest) SetSourceId(v string) *CreateEnterpriseTodoTaskRequest {
	s.SourceId = &v
	return s
}

func (s *CreateEnterpriseTodoTaskRequest) SetSourceTitle(v string) *CreateEnterpriseTodoTaskRequest {
	s.SourceTitle = &v
	return s
}

func (s *CreateEnterpriseTodoTaskRequest) SetSubject(v string) *CreateEnterpriseTodoTaskRequest {
	s.Subject = &v
	return s
}

func (s *CreateEnterpriseTodoTaskRequest) SetTrackerIds(v []*string) *CreateEnterpriseTodoTaskRequest {
	s.TrackerIds = v
	return s
}

func (s *CreateEnterpriseTodoTaskRequest) SetType(v string) *CreateEnterpriseTodoTaskRequest {
	s.Type = &v
	return s
}

type CreateEnterpriseTodoTaskRequestCustomFields struct {
	FieldKey   *string `json:"fieldKey,omitempty" xml:"fieldKey,omitempty"`
	FieldLink  *string `json:"fieldLink,omitempty" xml:"fieldLink,omitempty"`
	FieldType  *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
	Icon       *string `json:"icon,omitempty" xml:"icon,omitempty"`
}

func (s CreateEnterpriseTodoTaskRequestCustomFields) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseTodoTaskRequestCustomFields) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseTodoTaskRequestCustomFields) SetFieldKey(v string) *CreateEnterpriseTodoTaskRequestCustomFields {
	s.FieldKey = &v
	return s
}

func (s *CreateEnterpriseTodoTaskRequestCustomFields) SetFieldLink(v string) *CreateEnterpriseTodoTaskRequestCustomFields {
	s.FieldLink = &v
	return s
}

func (s *CreateEnterpriseTodoTaskRequestCustomFields) SetFieldType(v string) *CreateEnterpriseTodoTaskRequestCustomFields {
	s.FieldType = &v
	return s
}

func (s *CreateEnterpriseTodoTaskRequestCustomFields) SetFieldValue(v string) *CreateEnterpriseTodoTaskRequestCustomFields {
	s.FieldValue = &v
	return s
}

func (s *CreateEnterpriseTodoTaskRequestCustomFields) SetIcon(v string) *CreateEnterpriseTodoTaskRequestCustomFields {
	s.Icon = &v
	return s
}

type CreateEnterpriseTodoTaskRequestDetailUrl struct {
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
	WebUrl *string `json:"webUrl,omitempty" xml:"webUrl,omitempty"`
}

func (s CreateEnterpriseTodoTaskRequestDetailUrl) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseTodoTaskRequestDetailUrl) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseTodoTaskRequestDetailUrl) SetAppUrl(v string) *CreateEnterpriseTodoTaskRequestDetailUrl {
	s.AppUrl = &v
	return s
}

func (s *CreateEnterpriseTodoTaskRequestDetailUrl) SetWebUrl(v string) *CreateEnterpriseTodoTaskRequestDetailUrl {
	s.WebUrl = &v
	return s
}

type CreateEnterpriseTodoTaskRequestNotifyConfigs struct {
	DingNotify *string `json:"dingNotify,omitempty" xml:"dingNotify,omitempty"`
}

func (s CreateEnterpriseTodoTaskRequestNotifyConfigs) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseTodoTaskRequestNotifyConfigs) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseTodoTaskRequestNotifyConfigs) SetDingNotify(v string) *CreateEnterpriseTodoTaskRequestNotifyConfigs {
	s.DingNotify = &v
	return s
}

type CreateEnterpriseTodoTaskResponseBody struct {
	BizCategoryId *string                                        `json:"bizCategoryId,omitempty" xml:"bizCategoryId,omitempty"`
	CreatedTime   *int64                                         `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	CreatorId     *string                                        `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description   *string                                        `json:"description,omitempty" xml:"description,omitempty"`
	DetailUrl     *CreateEnterpriseTodoTaskResponseBodyDetailUrl `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	Done          *bool                                          `json:"done,omitempty" xml:"done,omitempty"`
	DueTime       *int64                                         `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	ExecutorIds   []*string                                      `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	ModifiedTime  *int64                                         `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	SourceId      *string                                        `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	Subject       *string                                        `json:"subject,omitempty" xml:"subject,omitempty"`
	TaskId        *string                                        `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateEnterpriseTodoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseTodoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseTodoTaskResponseBody) SetBizCategoryId(v string) *CreateEnterpriseTodoTaskResponseBody {
	s.BizCategoryId = &v
	return s
}

func (s *CreateEnterpriseTodoTaskResponseBody) SetCreatedTime(v int64) *CreateEnterpriseTodoTaskResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CreateEnterpriseTodoTaskResponseBody) SetCreatorId(v string) *CreateEnterpriseTodoTaskResponseBody {
	s.CreatorId = &v
	return s
}

func (s *CreateEnterpriseTodoTaskResponseBody) SetDescription(v string) *CreateEnterpriseTodoTaskResponseBody {
	s.Description = &v
	return s
}

func (s *CreateEnterpriseTodoTaskResponseBody) SetDetailUrl(v *CreateEnterpriseTodoTaskResponseBodyDetailUrl) *CreateEnterpriseTodoTaskResponseBody {
	s.DetailUrl = v
	return s
}

func (s *CreateEnterpriseTodoTaskResponseBody) SetDone(v bool) *CreateEnterpriseTodoTaskResponseBody {
	s.Done = &v
	return s
}

func (s *CreateEnterpriseTodoTaskResponseBody) SetDueTime(v int64) *CreateEnterpriseTodoTaskResponseBody {
	s.DueTime = &v
	return s
}

func (s *CreateEnterpriseTodoTaskResponseBody) SetExecutorIds(v []*string) *CreateEnterpriseTodoTaskResponseBody {
	s.ExecutorIds = v
	return s
}

func (s *CreateEnterpriseTodoTaskResponseBody) SetModifiedTime(v int64) *CreateEnterpriseTodoTaskResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *CreateEnterpriseTodoTaskResponseBody) SetSourceId(v string) *CreateEnterpriseTodoTaskResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreateEnterpriseTodoTaskResponseBody) SetSubject(v string) *CreateEnterpriseTodoTaskResponseBody {
	s.Subject = &v
	return s
}

func (s *CreateEnterpriseTodoTaskResponseBody) SetTaskId(v string) *CreateEnterpriseTodoTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateEnterpriseTodoTaskResponseBodyDetailUrl struct {
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
	WebUrl *string `json:"webUrl,omitempty" xml:"webUrl,omitempty"`
}

func (s CreateEnterpriseTodoTaskResponseBodyDetailUrl) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseTodoTaskResponseBodyDetailUrl) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseTodoTaskResponseBodyDetailUrl) SetAppUrl(v string) *CreateEnterpriseTodoTaskResponseBodyDetailUrl {
	s.AppUrl = &v
	return s
}

func (s *CreateEnterpriseTodoTaskResponseBodyDetailUrl) SetWebUrl(v string) *CreateEnterpriseTodoTaskResponseBodyDetailUrl {
	s.WebUrl = &v
	return s
}

type CreateEnterpriseTodoTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEnterpriseTodoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEnterpriseTodoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseTodoTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseTodoTaskResponse) SetHeaders(v map[string]*string) *CreateEnterpriseTodoTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateEnterpriseTodoTaskResponse) SetStatusCode(v int32) *CreateEnterpriseTodoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEnterpriseTodoTaskResponse) SetBody(v *CreateEnterpriseTodoTaskResponseBody) *CreateEnterpriseTodoTaskResponse {
	s.Body = v
	return s
}

type CreateStandardTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateStandardTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateStandardTemplateHeaders) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateHeaders) SetCommonHeaders(v map[string]*string) *CreateStandardTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateStandardTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *CreateStandardTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateStandardTemplateRequest struct {
	Actions     []*CreateStandardTemplateRequestActions `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	Description *string                                 `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string                                 `json:"name,omitempty" xml:"name,omitempty"`
	OperatorId  *string                                 `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	Service     *CreateStandardTemplateRequestService   `json:"service,omitempty" xml:"service,omitempty" type:"Struct"`
}

func (s CreateStandardTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStandardTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateRequest) SetActions(v []*CreateStandardTemplateRequestActions) *CreateStandardTemplateRequest {
	s.Actions = v
	return s
}

func (s *CreateStandardTemplateRequest) SetDescription(v string) *CreateStandardTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateStandardTemplateRequest) SetName(v string) *CreateStandardTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateStandardTemplateRequest) SetOperatorId(v string) *CreateStandardTemplateRequest {
	s.OperatorId = &v
	return s
}

func (s *CreateStandardTemplateRequest) SetService(v *CreateStandardTemplateRequestService) *CreateStandardTemplateRequest {
	s.Service = v
	return s
}

type CreateStandardTemplateRequestActions struct {
	ActionGroup        *string `json:"actionGroup,omitempty" xml:"actionGroup,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	NeedReason         *bool   `json:"needReason,omitempty" xml:"needReason,omitempty"`
	NeedReasonRequired *bool   `json:"needReasonRequired,omitempty" xml:"needReasonRequired,omitempty"`
	Order              *int64  `json:"order,omitempty" xml:"order,omitempty"`
	StyleType          *int64  `json:"styleType,omitempty" xml:"styleType,omitempty"`
}

func (s CreateStandardTemplateRequestActions) String() string {
	return tea.Prettify(s)
}

func (s CreateStandardTemplateRequestActions) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateRequestActions) SetActionGroup(v string) *CreateStandardTemplateRequestActions {
	s.ActionGroup = &v
	return s
}

func (s *CreateStandardTemplateRequestActions) SetName(v string) *CreateStandardTemplateRequestActions {
	s.Name = &v
	return s
}

func (s *CreateStandardTemplateRequestActions) SetNeedReason(v bool) *CreateStandardTemplateRequestActions {
	s.NeedReason = &v
	return s
}

func (s *CreateStandardTemplateRequestActions) SetNeedReasonRequired(v bool) *CreateStandardTemplateRequestActions {
	s.NeedReasonRequired = &v
	return s
}

func (s *CreateStandardTemplateRequestActions) SetOrder(v int64) *CreateStandardTemplateRequestActions {
	s.Order = &v
	return s
}

func (s *CreateStandardTemplateRequestActions) SetStyleType(v int64) *CreateStandardTemplateRequestActions {
	s.StyleType = &v
	return s
}

type CreateStandardTemplateRequestService struct {
	CallbackUrl *string `json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
}

func (s CreateStandardTemplateRequestService) String() string {
	return tea.Prettify(s)
}

func (s CreateStandardTemplateRequestService) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateRequestService) SetCallbackUrl(v string) *CreateStandardTemplateRequestService {
	s.CallbackUrl = &v
	return s
}

type CreateStandardTemplateResponseBody struct {
	Actions     []*CreateStandardTemplateResponseBodyActions `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	Description *string                                      `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string                                      `json:"name,omitempty" xml:"name,omitempty"`
	TemplateKey *string                                      `json:"templateKey,omitempty" xml:"templateKey,omitempty"`
}

func (s CreateStandardTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStandardTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateResponseBody) SetActions(v []*CreateStandardTemplateResponseBodyActions) *CreateStandardTemplateResponseBody {
	s.Actions = v
	return s
}

func (s *CreateStandardTemplateResponseBody) SetDescription(v string) *CreateStandardTemplateResponseBody {
	s.Description = &v
	return s
}

func (s *CreateStandardTemplateResponseBody) SetName(v string) *CreateStandardTemplateResponseBody {
	s.Name = &v
	return s
}

func (s *CreateStandardTemplateResponseBody) SetTemplateKey(v string) *CreateStandardTemplateResponseBody {
	s.TemplateKey = &v
	return s
}

type CreateStandardTemplateResponseBodyActions struct {
	ActionKey          *string `json:"actionKey,omitempty" xml:"actionKey,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	NeedReason         *bool   `json:"needReason,omitempty" xml:"needReason,omitempty"`
	NeedReasonRequired *bool   `json:"needReasonRequired,omitempty" xml:"needReasonRequired,omitempty"`
	Order              *int64  `json:"order,omitempty" xml:"order,omitempty"`
	StyleType          *int64  `json:"styleType,omitempty" xml:"styleType,omitempty"`
}

func (s CreateStandardTemplateResponseBodyActions) String() string {
	return tea.Prettify(s)
}

func (s CreateStandardTemplateResponseBodyActions) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateResponseBodyActions) SetActionKey(v string) *CreateStandardTemplateResponseBodyActions {
	s.ActionKey = &v
	return s
}

func (s *CreateStandardTemplateResponseBodyActions) SetName(v string) *CreateStandardTemplateResponseBodyActions {
	s.Name = &v
	return s
}

func (s *CreateStandardTemplateResponseBodyActions) SetNeedReason(v bool) *CreateStandardTemplateResponseBodyActions {
	s.NeedReason = &v
	return s
}

func (s *CreateStandardTemplateResponseBodyActions) SetNeedReasonRequired(v bool) *CreateStandardTemplateResponseBodyActions {
	s.NeedReasonRequired = &v
	return s
}

func (s *CreateStandardTemplateResponseBodyActions) SetOrder(v int64) *CreateStandardTemplateResponseBodyActions {
	s.Order = &v
	return s
}

func (s *CreateStandardTemplateResponseBodyActions) SetStyleType(v int64) *CreateStandardTemplateResponseBodyActions {
	s.StyleType = &v
	return s
}

type CreateStandardTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStandardTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStandardTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStandardTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateResponse) SetHeaders(v map[string]*string) *CreateStandardTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateStandardTemplateResponse) SetStatusCode(v int32) *CreateStandardTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStandardTemplateResponse) SetBody(v *CreateStandardTemplateResponseBody) *CreateStandardTemplateResponse {
	s.Body = v
	return s
}

type DeleteCategorySourceConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteCategorySourceConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteCategorySourceConfigHeaders) GoString() string {
	return s.String()
}

func (s *DeleteCategorySourceConfigHeaders) SetCommonHeaders(v map[string]*string) *DeleteCategorySourceConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteCategorySourceConfigHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteCategorySourceConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteCategorySourceConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1001
	BizCategoryId *string `json:"bizCategoryId,omitempty" xml:"bizCategoryId,omitempty"`
}

func (s DeleteCategorySourceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCategorySourceConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteCategorySourceConfigRequest) SetBizCategoryId(v string) *DeleteCategorySourceConfigRequest {
	s.BizCategoryId = &v
	return s
}

type DeleteCategorySourceConfigResponseBody struct {
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteCategorySourceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCategorySourceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCategorySourceConfigResponseBody) SetSuccess(v string) *DeleteCategorySourceConfigResponseBody {
	s.Success = &v
	return s
}

type DeleteCategorySourceConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCategorySourceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCategorySourceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCategorySourceConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteCategorySourceConfigResponse) SetHeaders(v map[string]*string) *DeleteCategorySourceConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteCategorySourceConfigResponse) SetStatusCode(v int32) *DeleteCategorySourceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCategorySourceConfigResponse) SetBody(v *DeleteCategorySourceConfigResponseBody) *DeleteCategorySourceConfigResponse {
	s.Body = v
	return s
}

type DeleteTodoEETaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteTodoEETaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteTodoEETaskHeaders) GoString() string {
	return s.String()
}

func (s *DeleteTodoEETaskHeaders) SetCommonHeaders(v map[string]*string) *DeleteTodoEETaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteTodoEETaskHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteTodoEETaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteTodoEETaskRequest struct {
	TaskIds []*string `json:"taskIds,omitempty" xml:"taskIds,omitempty" type:"Repeated"`
}

func (s DeleteTodoEETaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTodoEETaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteTodoEETaskRequest) SetTaskIds(v []*string) *DeleteTodoEETaskRequest {
	s.TaskIds = v
	return s
}

type DeleteTodoEETaskResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteTodoEETaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTodoEETaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTodoEETaskResponseBody) SetResult(v bool) *DeleteTodoEETaskResponseBody {
	s.Result = &v
	return s
}

type DeleteTodoEETaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTodoEETaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTodoEETaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTodoEETaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteTodoEETaskResponse) SetHeaders(v map[string]*string) *DeleteTodoEETaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteTodoEETaskResponse) SetStatusCode(v int32) *DeleteTodoEETaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTodoEETaskResponse) SetBody(v *DeleteTodoEETaskResponseBody) *DeleteTodoEETaskResponse {
	s.Body = v
	return s
}

type GetCategorySourceConfigListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCategorySourceConfigListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCategorySourceConfigListHeaders) GoString() string {
	return s.String()
}

func (s *GetCategorySourceConfigListHeaders) SetCommonHeaders(v map[string]*string) *GetCategorySourceConfigListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCategorySourceConfigListHeaders) SetXAcsDingtalkAccessToken(v string) *GetCategorySourceConfigListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCategorySourceConfigListRequest struct {
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetCategorySourceConfigListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCategorySourceConfigListRequest) GoString() string {
	return s.String()
}

func (s *GetCategorySourceConfigListRequest) SetNextToken(v string) *GetCategorySourceConfigListRequest {
	s.NextToken = &v
	return s
}

type GetCategorySourceConfigListResponseBody struct {
	Configs    []*GetCategorySourceConfigListResponseBodyConfigs `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
	NextToken  *string                                           `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	TotalCount *int32                                            `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetCategorySourceConfigListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCategorySourceConfigListResponseBody) GoString() string {
	return s.String()
}

func (s *GetCategorySourceConfigListResponseBody) SetConfigs(v []*GetCategorySourceConfigListResponseBodyConfigs) *GetCategorySourceConfigListResponseBody {
	s.Configs = v
	return s
}

func (s *GetCategorySourceConfigListResponseBody) SetNextToken(v string) *GetCategorySourceConfigListResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetCategorySourceConfigListResponseBody) SetTotalCount(v int32) *GetCategorySourceConfigListResponseBody {
	s.TotalCount = &v
	return s
}

type GetCategorySourceConfigListResponseBodyConfigs struct {
	BizCategoryId   *string `json:"bizCategoryId,omitempty" xml:"bizCategoryId,omitempty"`
	BizCategoryName *string `json:"bizCategoryName,omitempty" xml:"bizCategoryName,omitempty"`
	ConfigId        *string `json:"configId,omitempty" xml:"configId,omitempty"`
}

func (s GetCategorySourceConfigListResponseBodyConfigs) String() string {
	return tea.Prettify(s)
}

func (s GetCategorySourceConfigListResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *GetCategorySourceConfigListResponseBodyConfigs) SetBizCategoryId(v string) *GetCategorySourceConfigListResponseBodyConfigs {
	s.BizCategoryId = &v
	return s
}

func (s *GetCategorySourceConfigListResponseBodyConfigs) SetBizCategoryName(v string) *GetCategorySourceConfigListResponseBodyConfigs {
	s.BizCategoryName = &v
	return s
}

func (s *GetCategorySourceConfigListResponseBodyConfigs) SetConfigId(v string) *GetCategorySourceConfigListResponseBodyConfigs {
	s.ConfigId = &v
	return s
}

type GetCategorySourceConfigListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCategorySourceConfigListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCategorySourceConfigListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCategorySourceConfigListResponse) GoString() string {
	return s.String()
}

func (s *GetCategorySourceConfigListResponse) SetHeaders(v map[string]*string) *GetCategorySourceConfigListResponse {
	s.Headers = v
	return s
}

func (s *GetCategorySourceConfigListResponse) SetStatusCode(v int32) *GetCategorySourceConfigListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCategorySourceConfigListResponse) SetBody(v *GetCategorySourceConfigListResponseBody) *GetCategorySourceConfigListResponse {
	s.Body = v
	return s
}

type GetTemplateListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTemplateListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateListHeaders) GoString() string {
	return s.String()
}

func (s *GetTemplateListHeaders) SetCommonHeaders(v map[string]*string) *GetTemplateListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTemplateListHeaders) SetXAcsDingtalkAccessToken(v string) *GetTemplateListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTemplateListRequest struct {
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetTemplateListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateListRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateListRequest) SetPageNumber(v int32) *GetTemplateListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetTemplateListRequest) SetPageSize(v int32) *GetTemplateListRequest {
	s.PageSize = &v
	return s
}

type GetTemplateListResponseBody struct {
	Data       []*GetTemplateListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	PageNumber *int32                             `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32                             `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	TotalCount *int32                             `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetTemplateListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateListResponseBody) SetData(v []*GetTemplateListResponseBodyData) *GetTemplateListResponseBody {
	s.Data = v
	return s
}

func (s *GetTemplateListResponseBody) SetPageNumber(v int32) *GetTemplateListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetTemplateListResponseBody) SetPageSize(v int32) *GetTemplateListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetTemplateListResponseBody) SetTotalCount(v int32) *GetTemplateListResponseBody {
	s.TotalCount = &v
	return s
}

type GetTemplateListResponseBodyData struct {
	Actions      []*GetTemplateListResponseBodyDataActions `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	CreateTime   *int64                                    `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId    *string                                   `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description  *string                                   `json:"description,omitempty" xml:"description,omitempty"`
	ModifiedTime *int64                                    `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId   *string                                   `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	Name         *string                                   `json:"name,omitempty" xml:"name,omitempty"`
	TemplateKey  *string                                   `json:"templateKey,omitempty" xml:"templateKey,omitempty"`
}

func (s GetTemplateListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTemplateListResponseBodyData) SetActions(v []*GetTemplateListResponseBodyDataActions) *GetTemplateListResponseBodyData {
	s.Actions = v
	return s
}

func (s *GetTemplateListResponseBodyData) SetCreateTime(v int64) *GetTemplateListResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetTemplateListResponseBodyData) SetCreatorId(v string) *GetTemplateListResponseBodyData {
	s.CreatorId = &v
	return s
}

func (s *GetTemplateListResponseBodyData) SetDescription(v string) *GetTemplateListResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetTemplateListResponseBodyData) SetModifiedTime(v int64) *GetTemplateListResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *GetTemplateListResponseBodyData) SetModifierId(v string) *GetTemplateListResponseBodyData {
	s.ModifierId = &v
	return s
}

func (s *GetTemplateListResponseBodyData) SetName(v string) *GetTemplateListResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetTemplateListResponseBodyData) SetTemplateKey(v string) *GetTemplateListResponseBodyData {
	s.TemplateKey = &v
	return s
}

type GetTemplateListResponseBodyDataActions struct {
	ActionKey          *string `json:"actionKey,omitempty" xml:"actionKey,omitempty"`
	Description        *string `json:"description,omitempty" xml:"description,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	NeedReason         *bool   `json:"needReason,omitempty" xml:"needReason,omitempty"`
	NeedReasonRequired *bool   `json:"needReasonRequired,omitempty" xml:"needReasonRequired,omitempty"`
	Order              *int64  `json:"order,omitempty" xml:"order,omitempty"`
	StyleType          *int64  `json:"styleType,omitempty" xml:"styleType,omitempty"`
}

func (s GetTemplateListResponseBodyDataActions) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateListResponseBodyDataActions) GoString() string {
	return s.String()
}

func (s *GetTemplateListResponseBodyDataActions) SetActionKey(v string) *GetTemplateListResponseBodyDataActions {
	s.ActionKey = &v
	return s
}

func (s *GetTemplateListResponseBodyDataActions) SetDescription(v string) *GetTemplateListResponseBodyDataActions {
	s.Description = &v
	return s
}

func (s *GetTemplateListResponseBodyDataActions) SetName(v string) *GetTemplateListResponseBodyDataActions {
	s.Name = &v
	return s
}

func (s *GetTemplateListResponseBodyDataActions) SetNeedReason(v bool) *GetTemplateListResponseBodyDataActions {
	s.NeedReason = &v
	return s
}

func (s *GetTemplateListResponseBodyDataActions) SetNeedReasonRequired(v bool) *GetTemplateListResponseBodyDataActions {
	s.NeedReasonRequired = &v
	return s
}

func (s *GetTemplateListResponseBodyDataActions) SetOrder(v int64) *GetTemplateListResponseBodyDataActions {
	s.Order = &v
	return s
}

func (s *GetTemplateListResponseBodyDataActions) SetStyleType(v int64) *GetTemplateListResponseBodyDataActions {
	s.StyleType = &v
	return s
}

type GetTemplateListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTemplateListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateListResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateListResponse) SetHeaders(v map[string]*string) *GetTemplateListResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateListResponse) SetStatusCode(v int32) *GetTemplateListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateListResponse) SetBody(v *GetTemplateListResponseBody) *GetTemplateListResponse {
	s.Body = v
	return s
}

type GetUserTaskListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserTaskListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserTaskListHeaders) GoString() string {
	return s.String()
}

func (s *GetUserTaskListHeaders) SetCommonHeaders(v map[string]*string) *GetUserTaskListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserTaskListHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserTaskListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserTaskListRequest struct {
	Done       *bool   `json:"done,omitempty" xml:"done,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetUserTaskListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserTaskListRequest) GoString() string {
	return s.String()
}

func (s *GetUserTaskListRequest) SetDone(v bool) *GetUserTaskListRequest {
	s.Done = &v
	return s
}

func (s *GetUserTaskListRequest) SetPageNumber(v int32) *GetUserTaskListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetUserTaskListRequest) SetPageSize(v int32) *GetUserTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *GetUserTaskListRequest) SetType(v string) *GetUserTaskListRequest {
	s.Type = &v
	return s
}

type GetUserTaskListResponseBody struct {
	Data       []*GetUserTaskListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	PageNumber *int32                             `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32                             `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	TotalCount *int64                             `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetUserTaskListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserTaskListResponseBody) SetData(v []*GetUserTaskListResponseBodyData) *GetUserTaskListResponseBody {
	s.Data = v
	return s
}

func (s *GetUserTaskListResponseBody) SetPageNumber(v int32) *GetUserTaskListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetUserTaskListResponseBody) SetPageSize(v int32) *GetUserTaskListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetUserTaskListResponseBody) SetTotalCount(v int64) *GetUserTaskListResponseBody {
	s.TotalCount = &v
	return s
}

type GetUserTaskListResponseBodyData struct {
	CreatedTime *int64  `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Done        *bool   `json:"done,omitempty" xml:"done,omitempty"`
	DueTime     *int64  `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	Subject     *string `json:"subject,omitempty" xml:"subject,omitempty"`
	TaskId      *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetUserTaskListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUserTaskListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserTaskListResponseBodyData) SetCreatedTime(v int64) *GetUserTaskListResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetUserTaskListResponseBodyData) SetDescription(v string) *GetUserTaskListResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetUserTaskListResponseBodyData) SetDone(v bool) *GetUserTaskListResponseBodyData {
	s.Done = &v
	return s
}

func (s *GetUserTaskListResponseBodyData) SetDueTime(v int64) *GetUserTaskListResponseBodyData {
	s.DueTime = &v
	return s
}

func (s *GetUserTaskListResponseBodyData) SetSubject(v string) *GetUserTaskListResponseBodyData {
	s.Subject = &v
	return s
}

func (s *GetUserTaskListResponseBodyData) SetTaskId(v string) *GetUserTaskListResponseBodyData {
	s.TaskId = &v
	return s
}

type GetUserTaskListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserTaskListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserTaskListResponse) GoString() string {
	return s.String()
}

func (s *GetUserTaskListResponse) SetHeaders(v map[string]*string) *GetUserTaskListResponse {
	s.Headers = v
	return s
}

func (s *GetUserTaskListResponse) SetStatusCode(v int32) *GetUserTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserTaskListResponse) SetBody(v *GetUserTaskListResponseBody) *GetUserTaskListResponse {
	s.Body = v
	return s
}

type QueryTaskExecutionStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryTaskExecutionStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskExecutionStatusHeaders) GoString() string {
	return s.String()
}

func (s *QueryTaskExecutionStatusHeaders) SetCommonHeaders(v map[string]*string) *QueryTaskExecutionStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryTaskExecutionStatusHeaders) SetXAcsDingtalkAccessToken(v string) *QueryTaskExecutionStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryTaskExecutionStatusRequest struct {
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s QueryTaskExecutionStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskExecutionStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskExecutionStatusRequest) SetPageNumber(v int32) *QueryTaskExecutionStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryTaskExecutionStatusRequest) SetPageSize(v int32) *QueryTaskExecutionStatusRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTaskExecutionStatusRequest) SetTaskId(v string) *QueryTaskExecutionStatusRequest {
	s.TaskId = &v
	return s
}

type QueryTaskExecutionStatusResponseBody struct {
	Data       []*QueryTaskExecutionStatusResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	HasMore    *bool                                       `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	PageNumber *int32                                      `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32                                      `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	TotalCount *int32                                      `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryTaskExecutionStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskExecutionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTaskExecutionStatusResponseBody) SetData(v []*QueryTaskExecutionStatusResponseBodyData) *QueryTaskExecutionStatusResponseBody {
	s.Data = v
	return s
}

func (s *QueryTaskExecutionStatusResponseBody) SetHasMore(v bool) *QueryTaskExecutionStatusResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryTaskExecutionStatusResponseBody) SetPageNumber(v int32) *QueryTaskExecutionStatusResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryTaskExecutionStatusResponseBody) SetPageSize(v int32) *QueryTaskExecutionStatusResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTaskExecutionStatusResponseBody) SetTotalCount(v int32) *QueryTaskExecutionStatusResponseBody {
	s.TotalCount = &v
	return s
}

type QueryTaskExecutionStatusResponseBodyData struct {
	Done       *bool   `json:"done,omitempty" xml:"done,omitempty"`
	ExecutorId *string `json:"executorId,omitempty" xml:"executorId,omitempty"`
	FinishDate *int64  `json:"finishDate,omitempty" xml:"finishDate,omitempty"`
}

func (s QueryTaskExecutionStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskExecutionStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTaskExecutionStatusResponseBodyData) SetDone(v bool) *QueryTaskExecutionStatusResponseBodyData {
	s.Done = &v
	return s
}

func (s *QueryTaskExecutionStatusResponseBodyData) SetExecutorId(v string) *QueryTaskExecutionStatusResponseBodyData {
	s.ExecutorId = &v
	return s
}

func (s *QueryTaskExecutionStatusResponseBodyData) SetFinishDate(v int64) *QueryTaskExecutionStatusResponseBodyData {
	s.FinishDate = &v
	return s
}

type QueryTaskExecutionStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTaskExecutionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTaskExecutionStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskExecutionStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskExecutionStatusResponse) SetHeaders(v map[string]*string) *QueryTaskExecutionStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryTaskExecutionStatusResponse) SetStatusCode(v int32) *QueryTaskExecutionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTaskExecutionStatusResponse) SetBody(v *QueryTaskExecutionStatusResponseBody) *QueryTaskExecutionStatusResponse {
	s.Body = v
	return s
}

type RegisterCategorySourceConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RegisterCategorySourceConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s RegisterCategorySourceConfigHeaders) GoString() string {
	return s.String()
}

func (s *RegisterCategorySourceConfigHeaders) SetCommonHeaders(v map[string]*string) *RegisterCategorySourceConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RegisterCategorySourceConfigHeaders) SetXAcsDingtalkAccessToken(v string) *RegisterCategorySourceConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RegisterCategorySourceConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1001
	BizCategoryId *string `json:"bizCategoryId,omitempty" xml:"bizCategoryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 财务_审批_考勤
	BizCategoryName *string `json:"bizCategoryName,omitempty" xml:"bizCategoryName,omitempty"`
	// example:
	//
	// 张三
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s RegisterCategorySourceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterCategorySourceConfigRequest) GoString() string {
	return s.String()
}

func (s *RegisterCategorySourceConfigRequest) SetBizCategoryId(v string) *RegisterCategorySourceConfigRequest {
	s.BizCategoryId = &v
	return s
}

func (s *RegisterCategorySourceConfigRequest) SetBizCategoryName(v string) *RegisterCategorySourceConfigRequest {
	s.BizCategoryName = &v
	return s
}

func (s *RegisterCategorySourceConfigRequest) SetOperatorId(v string) *RegisterCategorySourceConfigRequest {
	s.OperatorId = &v
	return s
}

type RegisterCategorySourceConfigResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RegisterCategorySourceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterCategorySourceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterCategorySourceConfigResponseBody) SetSuccess(v bool) *RegisterCategorySourceConfigResponseBody {
	s.Success = &v
	return s
}

type RegisterCategorySourceConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterCategorySourceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterCategorySourceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterCategorySourceConfigResponse) GoString() string {
	return s.String()
}

func (s *RegisterCategorySourceConfigResponse) SetHeaders(v map[string]*string) *RegisterCategorySourceConfigResponse {
	s.Headers = v
	return s
}

func (s *RegisterCategorySourceConfigResponse) SetStatusCode(v int32) *RegisterCategorySourceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterCategorySourceConfigResponse) SetBody(v *RegisterCategorySourceConfigResponseBody) *RegisterCategorySourceConfigResponse {
	s.Body = v
	return s
}

type UpdateCategorySourceConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateCategorySourceConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateCategorySourceConfigHeaders) GoString() string {
	return s.String()
}

func (s *UpdateCategorySourceConfigHeaders) SetCommonHeaders(v map[string]*string) *UpdateCategorySourceConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateCategorySourceConfigHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateCategorySourceConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateCategorySourceConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1001
	BizCategoryId *string `json:"bizCategoryId,omitempty" xml:"bizCategoryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 考勤_财务
	BizCategoryName *string `json:"bizCategoryName,omitempty" xml:"bizCategoryName,omitempty"`
	OperatorId      *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s UpdateCategorySourceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCategorySourceConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateCategorySourceConfigRequest) SetBizCategoryId(v string) *UpdateCategorySourceConfigRequest {
	s.BizCategoryId = &v
	return s
}

func (s *UpdateCategorySourceConfigRequest) SetBizCategoryName(v string) *UpdateCategorySourceConfigRequest {
	s.BizCategoryName = &v
	return s
}

func (s *UpdateCategorySourceConfigRequest) SetOperatorId(v string) *UpdateCategorySourceConfigRequest {
	s.OperatorId = &v
	return s
}

type UpdateCategorySourceConfigResponseBody struct {
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateCategorySourceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCategorySourceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCategorySourceConfigResponseBody) SetSuccess(v string) *UpdateCategorySourceConfigResponseBody {
	s.Success = &v
	return s
}

type UpdateCategorySourceConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCategorySourceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCategorySourceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCategorySourceConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateCategorySourceConfigResponse) SetHeaders(v map[string]*string) *UpdateCategorySourceConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateCategorySourceConfigResponse) SetStatusCode(v int32) *UpdateCategorySourceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCategorySourceConfigResponse) SetBody(v *UpdateCategorySourceConfigResponseBody) *UpdateCategorySourceConfigResponse {
	s.Body = v
	return s
}

type UpdateStandardTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateStandardTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateStandardTemplateHeaders) GoString() string {
	return s.String()
}

func (s *UpdateStandardTemplateHeaders) SetCommonHeaders(v map[string]*string) *UpdateStandardTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateStandardTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateStandardTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateStandardTemplateRequest struct {
	Actions     []*UpdateStandardTemplateRequestActions `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	OperatorId  *string                                 `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	Service     *UpdateStandardTemplateRequestService   `json:"service,omitempty" xml:"service,omitempty" type:"Struct"`
	TemplateKey *string                                 `json:"templateKey,omitempty" xml:"templateKey,omitempty"`
}

func (s UpdateStandardTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStandardTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateStandardTemplateRequest) SetActions(v []*UpdateStandardTemplateRequestActions) *UpdateStandardTemplateRequest {
	s.Actions = v
	return s
}

func (s *UpdateStandardTemplateRequest) SetOperatorId(v string) *UpdateStandardTemplateRequest {
	s.OperatorId = &v
	return s
}

func (s *UpdateStandardTemplateRequest) SetService(v *UpdateStandardTemplateRequestService) *UpdateStandardTemplateRequest {
	s.Service = v
	return s
}

func (s *UpdateStandardTemplateRequest) SetTemplateKey(v string) *UpdateStandardTemplateRequest {
	s.TemplateKey = &v
	return s
}

type UpdateStandardTemplateRequestActions struct {
	ActionGroup        *string `json:"actionGroup,omitempty" xml:"actionGroup,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	NeedReason         *bool   `json:"needReason,omitempty" xml:"needReason,omitempty"`
	NeedReasonRequired *bool   `json:"needReasonRequired,omitempty" xml:"needReasonRequired,omitempty"`
	Order              *int64  `json:"order,omitempty" xml:"order,omitempty"`
	StyleType          *int64  `json:"styleType,omitempty" xml:"styleType,omitempty"`
}

func (s UpdateStandardTemplateRequestActions) String() string {
	return tea.Prettify(s)
}

func (s UpdateStandardTemplateRequestActions) GoString() string {
	return s.String()
}

func (s *UpdateStandardTemplateRequestActions) SetActionGroup(v string) *UpdateStandardTemplateRequestActions {
	s.ActionGroup = &v
	return s
}

func (s *UpdateStandardTemplateRequestActions) SetName(v string) *UpdateStandardTemplateRequestActions {
	s.Name = &v
	return s
}

func (s *UpdateStandardTemplateRequestActions) SetNeedReason(v bool) *UpdateStandardTemplateRequestActions {
	s.NeedReason = &v
	return s
}

func (s *UpdateStandardTemplateRequestActions) SetNeedReasonRequired(v bool) *UpdateStandardTemplateRequestActions {
	s.NeedReasonRequired = &v
	return s
}

func (s *UpdateStandardTemplateRequestActions) SetOrder(v int64) *UpdateStandardTemplateRequestActions {
	s.Order = &v
	return s
}

func (s *UpdateStandardTemplateRequestActions) SetStyleType(v int64) *UpdateStandardTemplateRequestActions {
	s.StyleType = &v
	return s
}

type UpdateStandardTemplateRequestService struct {
	CallbackUrl *string `json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
}

func (s UpdateStandardTemplateRequestService) String() string {
	return tea.Prettify(s)
}

func (s UpdateStandardTemplateRequestService) GoString() string {
	return s.String()
}

func (s *UpdateStandardTemplateRequestService) SetCallbackUrl(v string) *UpdateStandardTemplateRequestService {
	s.CallbackUrl = &v
	return s
}

type UpdateStandardTemplateResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateStandardTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateStandardTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStandardTemplateResponseBody) SetSuccess(v bool) *UpdateStandardTemplateResponseBody {
	s.Success = &v
	return s
}

type UpdateStandardTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStandardTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStandardTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateStandardTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateStandardTemplateResponse) SetHeaders(v map[string]*string) *UpdateStandardTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateStandardTemplateResponse) SetStatusCode(v int32) *UpdateStandardTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStandardTemplateResponse) SetBody(v *UpdateStandardTemplateResponseBody) *UpdateStandardTemplateResponse {
	s.Body = v
	return s
}

type UpdateTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTaskHeaders) SetCommonHeaders(v map[string]*string) *UpdateTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTaskHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateTaskRequest struct {
	Description *string   `json:"description,omitempty" xml:"description,omitempty"`
	Done        *bool     `json:"done,omitempty" xml:"done,omitempty"`
	DueTime     *int64    `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	ExecutorIds []*string `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	Subject     *string   `json:"subject,omitempty" xml:"subject,omitempty"`
	TaskId      *int64    `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s UpdateTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskRequest) SetDescription(v string) *UpdateTaskRequest {
	s.Description = &v
	return s
}

func (s *UpdateTaskRequest) SetDone(v bool) *UpdateTaskRequest {
	s.Done = &v
	return s
}

func (s *UpdateTaskRequest) SetDueTime(v int64) *UpdateTaskRequest {
	s.DueTime = &v
	return s
}

func (s *UpdateTaskRequest) SetExecutorIds(v []*string) *UpdateTaskRequest {
	s.ExecutorIds = v
	return s
}

func (s *UpdateTaskRequest) SetSubject(v string) *UpdateTaskRequest {
	s.Subject = &v
	return s
}

func (s *UpdateTaskRequest) SetTaskId(v int64) *UpdateTaskRequest {
	s.TaskId = &v
	return s
}

type UpdateTaskResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskResponseBody) SetResult(v bool) *UpdateTaskResponseBody {
	s.Result = &v
	return s
}

type UpdateTaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskResponse) SetHeaders(v map[string]*string) *UpdateTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskResponse) SetStatusCode(v int32) *UpdateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskResponse) SetBody(v *UpdateTaskResponseBody) *UpdateTaskResponse {
	s.Body = v
	return s
}

type UpdateUserTaskStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateUserTaskStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserTaskStatusHeaders) GoString() string {
	return s.String()
}

func (s *UpdateUserTaskStatusHeaders) SetCommonHeaders(v map[string]*string) *UpdateUserTaskStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateUserTaskStatusHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateUserTaskStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateUserTaskStatusRequest struct {
	OperatorId       *string                                        `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	UserTaskStatuses []*UpdateUserTaskStatusRequestUserTaskStatuses `json:"userTaskStatuses,omitempty" xml:"userTaskStatuses,omitempty" type:"Repeated"`
}

func (s UpdateUserTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserTaskStatusRequest) SetOperatorId(v string) *UpdateUserTaskStatusRequest {
	s.OperatorId = &v
	return s
}

func (s *UpdateUserTaskStatusRequest) SetUserTaskStatuses(v []*UpdateUserTaskStatusRequestUserTaskStatuses) *UpdateUserTaskStatusRequest {
	s.UserTaskStatuses = v
	return s
}

type UpdateUserTaskStatusRequestUserTaskStatuses struct {
	Done   *bool   `json:"done,omitempty" xml:"done,omitempty"`
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s UpdateUserTaskStatusRequestUserTaskStatuses) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserTaskStatusRequestUserTaskStatuses) GoString() string {
	return s.String()
}

func (s *UpdateUserTaskStatusRequestUserTaskStatuses) SetDone(v bool) *UpdateUserTaskStatusRequestUserTaskStatuses {
	s.Done = &v
	return s
}

func (s *UpdateUserTaskStatusRequestUserTaskStatuses) SetTaskId(v string) *UpdateUserTaskStatusRequestUserTaskStatuses {
	s.TaskId = &v
	return s
}

type UpdateUserTaskStatusResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateUserTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserTaskStatusResponseBody) SetSuccess(v bool) *UpdateUserTaskStatusResponseBody {
	s.Success = &v
	return s
}

type UpdateUserTaskStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserTaskStatusResponse) SetHeaders(v map[string]*string) *UpdateUserTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserTaskStatusResponse) SetStatusCode(v int32) *UpdateUserTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserTaskStatusResponse) SetBody(v *UpdateUserTaskStatusResponseBody) *UpdateUserTaskStatusResponse {
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
// 创建专属待办
//
// @param request - AppCreateEnterpriseTodoTaskRequest
//
// @param headers - AppCreateEnterpriseTodoTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AppCreateEnterpriseTodoTaskResponse
func (client *Client) AppCreateEnterpriseTodoTaskWithOptions(request *AppCreateEnterpriseTodoTaskRequest, headers *AppCreateEnterpriseTodoTaskHeaders, runtime *util.RuntimeOptions) (_result *AppCreateEnterpriseTodoTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCategoryId)) {
		body["bizCategoryId"] = request.BizCategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.BizCreatedTime)) {
		body["bizCreatedTime"] = request.BizCreatedTime
	}

	if !tea.BoolValue(util.IsUnset(request.CustomFields)) {
		body["customFields"] = request.CustomFields
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DetailUrl)) {
		body["detailUrl"] = request.DetailUrl
	}

	if !tea.BoolValue(util.IsUnset(request.DueTime)) {
		body["dueTime"] = request.DueTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorIds)) {
		body["executorIds"] = request.ExecutorIds
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyConfigs)) {
		body["notifyConfigs"] = request.NotifyConfigs
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		body["sourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceTitle)) {
		body["sourceTitle"] = request.SourceTitle
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		body["subject"] = request.Subject
	}

	if !tea.BoolValue(util.IsUnset(request.ToolbarTemplateKey)) {
		body["toolbarTemplateKey"] = request.ToolbarTemplateKey
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
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
		Action:      tea.String("AppCreateEnterpriseTodoTask"),
		Version:     tea.String("todoEE_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todoEE/apps/users/tasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AppCreateEnterpriseTodoTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建专属待办
//
// @param request - AppCreateEnterpriseTodoTaskRequest
//
// @return AppCreateEnterpriseTodoTaskResponse
func (client *Client) AppCreateEnterpriseTodoTask(request *AppCreateEnterpriseTodoTaskRequest) (_result *AppCreateEnterpriseTodoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AppCreateEnterpriseTodoTaskHeaders{}
	_result = &AppCreateEnterpriseTodoTaskResponse{}
	_body, _err := client.AppCreateEnterpriseTodoTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除专属待办
//
// @param request - AppDeleteTodoEETaskRequest
//
// @param headers - AppDeleteTodoEETaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AppDeleteTodoEETaskResponse
func (client *Client) AppDeleteTodoEETaskWithOptions(request *AppDeleteTodoEETaskRequest, headers *AppDeleteTodoEETaskHeaders, runtime *util.RuntimeOptions) (_result *AppDeleteTodoEETaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskIds)) {
		body["taskIds"] = request.TaskIds
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
		Action:      tea.String("AppDeleteTodoEETask"),
		Version:     tea.String("todoEE_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todoEE/apps/users/tasks/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AppDeleteTodoEETaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除专属待办
//
// @param request - AppDeleteTodoEETaskRequest
//
// @return AppDeleteTodoEETaskResponse
func (client *Client) AppDeleteTodoEETask(request *AppDeleteTodoEETaskRequest) (_result *AppDeleteTodoEETaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AppDeleteTodoEETaskHeaders{}
	_result = &AppDeleteTodoEETaskResponse{}
	_body, _err := client.AppDeleteTodoEETaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户待办列表
//
// @param request - AppGetUserTaskListRequest
//
// @param headers - AppGetUserTaskListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AppGetUserTaskListResponse
func (client *Client) AppGetUserTaskListWithOptions(request *AppGetUserTaskListRequest, headers *AppGetUserTaskListHeaders, runtime *util.RuntimeOptions) (_result *AppGetUserTaskListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Done)) {
		body["done"] = request.Done
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
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
		Action:      tea.String("AppGetUserTaskList"),
		Version:     tea.String("todoEE_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todoEE/apps/users/tasks/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AppGetUserTaskListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户待办列表
//
// @param request - AppGetUserTaskListRequest
//
// @return AppGetUserTaskListResponse
func (client *Client) AppGetUserTaskList(request *AppGetUserTaskListRequest) (_result *AppGetUserTaskListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AppGetUserTaskListHeaders{}
	_result = &AppGetUserTaskListResponse{}
	_body, _err := client.AppGetUserTaskListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新专属待办信息
//
// @param request - AppUpdateTaskRequest
//
// @param headers - AppUpdateTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AppUpdateTaskResponse
func (client *Client) AppUpdateTaskWithOptions(request *AppUpdateTaskRequest, headers *AppUpdateTaskHeaders, runtime *util.RuntimeOptions) (_result *AppUpdateTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCreatedTime)) {
		body["bizCreatedTime"] = request.BizCreatedTime
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Done)) {
		body["done"] = request.Done
	}

	if !tea.BoolValue(util.IsUnset(request.DueTime)) {
		body["dueTime"] = request.DueTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorIds)) {
		body["executorIds"] = request.ExecutorIds
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		body["subject"] = request.Subject
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.ToolbarTemplateKey)) {
		body["toolbarTemplateKey"] = request.ToolbarTemplateKey
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
		Action:      tea.String("AppUpdateTask"),
		Version:     tea.String("todoEE_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todoEE/apps/users/tasks/infos"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AppUpdateTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新专属待办信息
//
// @param request - AppUpdateTaskRequest
//
// @return AppUpdateTaskResponse
func (client *Client) AppUpdateTask(request *AppUpdateTaskRequest) (_result *AppUpdateTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AppUpdateTaskHeaders{}
	_result = &AppUpdateTaskResponse{}
	_body, _err := client.AppUpdateTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新用户的待办状态
//
// @param request - AppUpdateUserTaskStatusRequest
//
// @param headers - AppUpdateUserTaskStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AppUpdateUserTaskStatusResponse
func (client *Client) AppUpdateUserTaskStatusWithOptions(request *AppUpdateUserTaskStatusRequest, headers *AppUpdateUserTaskStatusHeaders, runtime *util.RuntimeOptions) (_result *AppUpdateUserTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.UserTaskStatuses)) {
		body["userTaskStatuses"] = request.UserTaskStatuses
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
		Action:      tea.String("AppUpdateUserTaskStatus"),
		Version:     tea.String("todoEE_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todoEE/apps/users/tasks/statuses"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AppUpdateUserTaskStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新用户的待办状态
//
// @param request - AppUpdateUserTaskStatusRequest
//
// @return AppUpdateUserTaskStatusResponse
func (client *Client) AppUpdateUserTaskStatus(request *AppUpdateUserTaskStatusRequest) (_result *AppUpdateUserTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AppUpdateUserTaskStatusHeaders{}
	_result = &AppUpdateUserTaskStatusResponse{}
	_body, _err := client.AppUpdateUserTaskStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建企业待办
//
// @param request - CreateEnterpriseTodoTaskRequest
//
// @param headers - CreateEnterpriseTodoTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEnterpriseTodoTaskResponse
func (client *Client) CreateEnterpriseTodoTaskWithOptions(request *CreateEnterpriseTodoTaskRequest, headers *CreateEnterpriseTodoTaskHeaders, runtime *util.RuntimeOptions) (_result *CreateEnterpriseTodoTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCategoryId)) {
		body["bizCategoryId"] = request.BizCategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomFields)) {
		body["customFields"] = request.CustomFields
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DetailUrl)) {
		body["detailUrl"] = request.DetailUrl
	}

	if !tea.BoolValue(util.IsUnset(request.DueTime)) {
		body["dueTime"] = request.DueTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorIds)) {
		body["executorIds"] = request.ExecutorIds
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyConfigs)) {
		body["notifyConfigs"] = request.NotifyConfigs
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		body["sourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceTitle)) {
		body["sourceTitle"] = request.SourceTitle
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		body["subject"] = request.Subject
	}

	if !tea.BoolValue(util.IsUnset(request.TrackerIds)) {
		body["trackerIds"] = request.TrackerIds
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
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
		Action:      tea.String("CreateEnterpriseTodoTask"),
		Version:     tea.String("todoEE_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todoEE/users/tasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEnterpriseTodoTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建企业待办
//
// @param request - CreateEnterpriseTodoTaskRequest
//
// @return CreateEnterpriseTodoTaskResponse
func (client *Client) CreateEnterpriseTodoTask(request *CreateEnterpriseTodoTaskRequest) (_result *CreateEnterpriseTodoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateEnterpriseTodoTaskHeaders{}
	_result = &CreateEnterpriseTodoTaskResponse{}
	_body, _err := client.CreateEnterpriseTodoTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建专属待办模板实例
//
// @param request - CreateStandardTemplateRequest
//
// @param headers - CreateStandardTemplateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStandardTemplateResponse
func (client *Client) CreateStandardTemplateWithOptions(request *CreateStandardTemplateRequest, headers *CreateStandardTemplateHeaders, runtime *util.RuntimeOptions) (_result *CreateStandardTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Actions)) {
		body["actions"] = request.Actions
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.Service)) {
		body["service"] = request.Service
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
		Action:      tea.String("CreateStandardTemplate"),
		Version:     tea.String("todoEE_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todoEE/standards/templates"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateStandardTemplateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建专属待办模板实例
//
// @param request - CreateStandardTemplateRequest
//
// @return CreateStandardTemplateResponse
func (client *Client) CreateStandardTemplate(request *CreateStandardTemplateRequest) (_result *CreateStandardTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateStandardTemplateHeaders{}
	_result = &CreateStandardTemplateResponse{}
	_body, _err := client.CreateStandardTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除应用类目信息
//
// @param request - DeleteCategorySourceConfigRequest
//
// @param headers - DeleteCategorySourceConfigHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCategorySourceConfigResponse
func (client *Client) DeleteCategorySourceConfigWithOptions(request *DeleteCategorySourceConfigRequest, headers *DeleteCategorySourceConfigHeaders, runtime *util.RuntimeOptions) (_result *DeleteCategorySourceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCategoryId)) {
		body["bizCategoryId"] = request.BizCategoryId
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
		Action:      tea.String("DeleteCategorySourceConfig"),
		Version:     tea.String("todoEE_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todoEE/apps/categories/sourceConfigs/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCategorySourceConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除应用类目信息
//
// @param request - DeleteCategorySourceConfigRequest
//
// @return DeleteCategorySourceConfigResponse
func (client *Client) DeleteCategorySourceConfig(request *DeleteCategorySourceConfigRequest) (_result *DeleteCategorySourceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteCategorySourceConfigHeaders{}
	_result = &DeleteCategorySourceConfigResponse{}
	_body, _err := client.DeleteCategorySourceConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除待办
//
// @param request - DeleteTodoEETaskRequest
//
// @param headers - DeleteTodoEETaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTodoEETaskResponse
func (client *Client) DeleteTodoEETaskWithOptions(request *DeleteTodoEETaskRequest, headers *DeleteTodoEETaskHeaders, runtime *util.RuntimeOptions) (_result *DeleteTodoEETaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskIds)) {
		body["taskIds"] = request.TaskIds
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
		Action:      tea.String("DeleteTodoEETask"),
		Version:     tea.String("todoEE_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todoEE/users/tasks/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTodoEETaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除待办
//
// @param request - DeleteTodoEETaskRequest
//
// @return DeleteTodoEETaskResponse
func (client *Client) DeleteTodoEETask(request *DeleteTodoEETaskRequest) (_result *DeleteTodoEETaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteTodoEETaskHeaders{}
	_result = &DeleteTodoEETaskResponse{}
	_body, _err := client.DeleteTodoEETaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询应用注册类目信息列表
//
// @param request - GetCategorySourceConfigListRequest
//
// @param headers - GetCategorySourceConfigListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCategorySourceConfigListResponse
func (client *Client) GetCategorySourceConfigListWithOptions(request *GetCategorySourceConfigListRequest, headers *GetCategorySourceConfigListHeaders, runtime *util.RuntimeOptions) (_result *GetCategorySourceConfigListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
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
		Action:      tea.String("GetCategorySourceConfigList"),
		Version:     tea.String("todoEE_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todoEE/apps/categories/sourceConfigs/lists/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCategorySourceConfigListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询应用注册类目信息列表
//
// @param request - GetCategorySourceConfigListRequest
//
// @return GetCategorySourceConfigListResponse
func (client *Client) GetCategorySourceConfigList(request *GetCategorySourceConfigListRequest) (_result *GetCategorySourceConfigListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCategorySourceConfigListHeaders{}
	_result = &GetCategorySourceConfigListResponse{}
	_body, _err := client.GetCategorySourceConfigListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询创建的Template列表
//
// @param request - GetTemplateListRequest
//
// @param headers - GetTemplateListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateListResponse
func (client *Client) GetTemplateListWithOptions(request *GetTemplateListRequest, headers *GetTemplateListHeaders, runtime *util.RuntimeOptions) (_result *GetTemplateListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
		Action:      tea.String("GetTemplateList"),
		Version:     tea.String("todoEE_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todoEE/templates/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTemplateListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询创建的Template列表
//
// @param request - GetTemplateListRequest
//
// @return GetTemplateListResponse
func (client *Client) GetTemplateList(request *GetTemplateListRequest) (_result *GetTemplateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTemplateListHeaders{}
	_result = &GetTemplateListResponse{}
	_body, _err := client.GetTemplateListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户待办列表
//
// @param request - GetUserTaskListRequest
//
// @param headers - GetUserTaskListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserTaskListResponse
func (client *Client) GetUserTaskListWithOptions(request *GetUserTaskListRequest, headers *GetUserTaskListHeaders, runtime *util.RuntimeOptions) (_result *GetUserTaskListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Done)) {
		body["done"] = request.Done
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
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
		Action:      tea.String("GetUserTaskList"),
		Version:     tea.String("todoEE_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todoEE/users/tasks/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserTaskListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户待办列表
//
// @param request - GetUserTaskListRequest
//
// @return GetUserTaskListResponse
func (client *Client) GetUserTaskList(request *GetUserTaskListRequest) (_result *GetUserTaskListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserTaskListHeaders{}
	_result = &GetUserTaskListResponse{}
	_body, _err := client.GetUserTaskListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询任务所有执行人的完成状态
//
// @param request - QueryTaskExecutionStatusRequest
//
// @param headers - QueryTaskExecutionStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTaskExecutionStatusResponse
func (client *Client) QueryTaskExecutionStatusWithOptions(request *QueryTaskExecutionStatusRequest, headers *QueryTaskExecutionStatusHeaders, runtime *util.RuntimeOptions) (_result *QueryTaskExecutionStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
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
		Action:      tea.String("QueryTaskExecutionStatus"),
		Version:     tea.String("todoEE_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todoEE/apps/users/tasks/executionStatuses"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTaskExecutionStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询任务所有执行人的完成状态
//
// @param request - QueryTaskExecutionStatusRequest
//
// @return QueryTaskExecutionStatusResponse
func (client *Client) QueryTaskExecutionStatus(request *QueryTaskExecutionStatusRequest) (_result *QueryTaskExecutionStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryTaskExecutionStatusHeaders{}
	_result = &QueryTaskExecutionStatusResponse{}
	_body, _err := client.QueryTaskExecutionStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 注册应用类目信息
//
// @param request - RegisterCategorySourceConfigRequest
//
// @param headers - RegisterCategorySourceConfigHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterCategorySourceConfigResponse
func (client *Client) RegisterCategorySourceConfigWithOptions(request *RegisterCategorySourceConfigRequest, headers *RegisterCategorySourceConfigHeaders, runtime *util.RuntimeOptions) (_result *RegisterCategorySourceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCategoryId)) {
		body["bizCategoryId"] = request.BizCategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.BizCategoryName)) {
		body["bizCategoryName"] = request.BizCategoryName
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
		Action:      tea.String("RegisterCategorySourceConfig"),
		Version:     tea.String("todoEE_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todoEE/apps/categories/sourceConfigs/register"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RegisterCategorySourceConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 注册应用类目信息
//
// @param request - RegisterCategorySourceConfigRequest
//
// @return RegisterCategorySourceConfigResponse
func (client *Client) RegisterCategorySourceConfig(request *RegisterCategorySourceConfigRequest) (_result *RegisterCategorySourceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RegisterCategorySourceConfigHeaders{}
	_result = &RegisterCategorySourceConfigResponse{}
	_body, _err := client.RegisterCategorySourceConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改应用类目注册信息
//
// @param request - UpdateCategorySourceConfigRequest
//
// @param headers - UpdateCategorySourceConfigHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCategorySourceConfigResponse
func (client *Client) UpdateCategorySourceConfigWithOptions(request *UpdateCategorySourceConfigRequest, headers *UpdateCategorySourceConfigHeaders, runtime *util.RuntimeOptions) (_result *UpdateCategorySourceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCategoryId)) {
		body["bizCategoryId"] = request.BizCategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.BizCategoryName)) {
		body["bizCategoryName"] = request.BizCategoryName
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
		Action:      tea.String("UpdateCategorySourceConfig"),
		Version:     tea.String("todoEE_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todoEE/apps/categories/sourceConfigs"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCategorySourceConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改应用类目注册信息
//
// @param request - UpdateCategorySourceConfigRequest
//
// @return UpdateCategorySourceConfigResponse
func (client *Client) UpdateCategorySourceConfig(request *UpdateCategorySourceConfigRequest) (_result *UpdateCategorySourceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateCategorySourceConfigHeaders{}
	_result = &UpdateCategorySourceConfigResponse{}
	_body, _err := client.UpdateCategorySourceConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新标准模板
//
// @param request - UpdateStandardTemplateRequest
//
// @param headers - UpdateStandardTemplateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStandardTemplateResponse
func (client *Client) UpdateStandardTemplateWithOptions(request *UpdateStandardTemplateRequest, headers *UpdateStandardTemplateHeaders, runtime *util.RuntimeOptions) (_result *UpdateStandardTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Actions)) {
		body["actions"] = request.Actions
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.Service)) {
		body["service"] = request.Service
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateKey)) {
		body["templateKey"] = request.TemplateKey
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
		Action:      tea.String("UpdateStandardTemplate"),
		Version:     tea.String("todoEE_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todoEE/standards/templates/infos"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateStandardTemplateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新标准模板
//
// @param request - UpdateStandardTemplateRequest
//
// @return UpdateStandardTemplateResponse
func (client *Client) UpdateStandardTemplate(request *UpdateStandardTemplateRequest) (_result *UpdateStandardTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateStandardTemplateHeaders{}
	_result = &UpdateStandardTemplateResponse{}
	_body, _err := client.UpdateStandardTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新待办信息
//
// @param request - UpdateTaskRequest
//
// @param headers - UpdateTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskResponse
func (client *Client) UpdateTaskWithOptions(request *UpdateTaskRequest, headers *UpdateTaskHeaders, runtime *util.RuntimeOptions) (_result *UpdateTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Done)) {
		body["done"] = request.Done
	}

	if !tea.BoolValue(util.IsUnset(request.DueTime)) {
		body["dueTime"] = request.DueTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorIds)) {
		body["executorIds"] = request.ExecutorIds
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		body["subject"] = request.Subject
	}

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
		Action:      tea.String("UpdateTask"),
		Version:     tea.String("todoEE_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todoEE/users/tasks/infos"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新待办信息
//
// @param request - UpdateTaskRequest
//
// @return UpdateTaskResponse
func (client *Client) UpdateTask(request *UpdateTaskRequest) (_result *UpdateTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTaskHeaders{}
	_result = &UpdateTaskResponse{}
	_body, _err := client.UpdateTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新用户的待办状态
//
// @param request - UpdateUserTaskStatusRequest
//
// @param headers - UpdateUserTaskStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserTaskStatusResponse
func (client *Client) UpdateUserTaskStatusWithOptions(request *UpdateUserTaskStatusRequest, headers *UpdateUserTaskStatusHeaders, runtime *util.RuntimeOptions) (_result *UpdateUserTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.UserTaskStatuses)) {
		body["userTaskStatuses"] = request.UserTaskStatuses
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
		Action:      tea.String("UpdateUserTaskStatus"),
		Version:     tea.String("todoEE_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todoEE/users/tasks/statuses"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserTaskStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新用户的待办状态
//
// @param request - UpdateUserTaskStatusRequest
//
// @return UpdateUserTaskStatusResponse
func (client *Client) UpdateUserTaskStatus(request *UpdateUserTaskStatusRequest) (_result *UpdateUserTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateUserTaskStatusHeaders{}
	_result = &UpdateUserTaskStatusResponse{}
	_body, _err := client.UpdateUserTaskStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
