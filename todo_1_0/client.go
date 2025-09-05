// This file is auto-generated, don't edit it. Thanks.
package todo_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CountTodoTasksHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CountTodoTasksHeaders) String() string {
	return tea.Prettify(s)
}

func (s CountTodoTasksHeaders) GoString() string {
	return s.String()
}

func (s *CountTodoTasksHeaders) SetCommonHeaders(v map[string]*string) *CountTodoTasksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CountTodoTasksHeaders) SetXAcsDingtalkAccessToken(v string) *CountTodoTasksHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CountTodoTasksRequest struct {
	FromDueTime *int64      `json:"fromDueTime,omitempty" xml:"fromDueTime,omitempty"`
	IsDone      *bool       `json:"isDone,omitempty" xml:"isDone,omitempty"`
	IsRecycled  *bool       `json:"isRecycled,omitempty" xml:"isRecycled,omitempty"`
	RoleTypes   [][]*string `json:"roleTypes,omitempty" xml:"roleTypes,omitempty" type:"Repeated"`
	ToDueTime   *int64      `json:"toDueTime,omitempty" xml:"toDueTime,omitempty"`
}

func (s CountTodoTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s CountTodoTasksRequest) GoString() string {
	return s.String()
}

func (s *CountTodoTasksRequest) SetFromDueTime(v int64) *CountTodoTasksRequest {
	s.FromDueTime = &v
	return s
}

func (s *CountTodoTasksRequest) SetIsDone(v bool) *CountTodoTasksRequest {
	s.IsDone = &v
	return s
}

func (s *CountTodoTasksRequest) SetIsRecycled(v bool) *CountTodoTasksRequest {
	s.IsRecycled = &v
	return s
}

func (s *CountTodoTasksRequest) SetRoleTypes(v [][]*string) *CountTodoTasksRequest {
	s.RoleTypes = v
	return s
}

func (s *CountTodoTasksRequest) SetToDueTime(v int64) *CountTodoTasksRequest {
	s.ToDueTime = &v
	return s
}

type CountTodoTasksResponseBody struct {
	// example:
	//
	// 99
	Result *int32 `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CountTodoTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountTodoTasksResponseBody) GoString() string {
	return s.String()
}

func (s *CountTodoTasksResponseBody) SetResult(v int32) *CountTodoTasksResponseBody {
	s.Result = &v
	return s
}

type CountTodoTasksResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CountTodoTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CountTodoTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s CountTodoTasksResponse) GoString() string {
	return s.String()
}

func (s *CountTodoTasksResponse) SetHeaders(v map[string]*string) *CountTodoTasksResponse {
	s.Headers = v
	return s
}

func (s *CountTodoTasksResponse) SetStatusCode(v int32) *CountTodoTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *CountTodoTasksResponse) SetBody(v *CountTodoTasksResponseBody) *CountTodoTasksResponse {
	s.Body = v
	return s
}

type CreatePersonalTodoTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreatePersonalTodoTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreatePersonalTodoTaskHeaders) GoString() string {
	return s.String()
}

func (s *CreatePersonalTodoTaskHeaders) SetCommonHeaders(v map[string]*string) *CreatePersonalTodoTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreatePersonalTodoTaskHeaders) SetXAcsDingtalkAccessToken(v string) *CreatePersonalTodoTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreatePersonalTodoTaskRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	DueTime     *int64  `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	// This parameter is required.
	ExecutorIds       []*string                                   `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	NotifyConfigs     *CreatePersonalTodoTaskRequestNotifyConfigs `json:"notifyConfigs,omitempty" xml:"notifyConfigs,omitempty" type:"Struct"`
	ParticipantIds    []*string                                   `json:"participantIds,omitempty" xml:"participantIds,omitempty" type:"Repeated"`
	ReminderTimeStamp *int64                                      `json:"reminderTimeStamp,omitempty" xml:"reminderTimeStamp,omitempty"`
	// This parameter is required.
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
}

func (s CreatePersonalTodoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePersonalTodoTaskRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalTodoTaskRequest) SetDescription(v string) *CreatePersonalTodoTaskRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalTodoTaskRequest) SetDueTime(v int64) *CreatePersonalTodoTaskRequest {
	s.DueTime = &v
	return s
}

func (s *CreatePersonalTodoTaskRequest) SetExecutorIds(v []*string) *CreatePersonalTodoTaskRequest {
	s.ExecutorIds = v
	return s
}

func (s *CreatePersonalTodoTaskRequest) SetNotifyConfigs(v *CreatePersonalTodoTaskRequestNotifyConfigs) *CreatePersonalTodoTaskRequest {
	s.NotifyConfigs = v
	return s
}

func (s *CreatePersonalTodoTaskRequest) SetParticipantIds(v []*string) *CreatePersonalTodoTaskRequest {
	s.ParticipantIds = v
	return s
}

func (s *CreatePersonalTodoTaskRequest) SetReminderTimeStamp(v int64) *CreatePersonalTodoTaskRequest {
	s.ReminderTimeStamp = &v
	return s
}

func (s *CreatePersonalTodoTaskRequest) SetSubject(v string) *CreatePersonalTodoTaskRequest {
	s.Subject = &v
	return s
}

type CreatePersonalTodoTaskRequestNotifyConfigs struct {
	DingNotify *string `json:"dingNotify,omitempty" xml:"dingNotify,omitempty"`
}

func (s CreatePersonalTodoTaskRequestNotifyConfigs) String() string {
	return tea.Prettify(s)
}

func (s CreatePersonalTodoTaskRequestNotifyConfigs) GoString() string {
	return s.String()
}

func (s *CreatePersonalTodoTaskRequestNotifyConfigs) SetDingNotify(v string) *CreatePersonalTodoTaskRequestNotifyConfigs {
	s.DingNotify = &v
	return s
}

type CreatePersonalTodoTaskResponseBody struct {
	CreatedTime *int64  `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	TaskId      *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreatePersonalTodoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePersonalTodoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersonalTodoTaskResponseBody) SetCreatedTime(v int64) *CreatePersonalTodoTaskResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CreatePersonalTodoTaskResponseBody) SetTaskId(v string) *CreatePersonalTodoTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreatePersonalTodoTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePersonalTodoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePersonalTodoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePersonalTodoTaskResponse) GoString() string {
	return s.String()
}

func (s *CreatePersonalTodoTaskResponse) SetHeaders(v map[string]*string) *CreatePersonalTodoTaskResponse {
	s.Headers = v
	return s
}

func (s *CreatePersonalTodoTaskResponse) SetStatusCode(v int32) *CreatePersonalTodoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePersonalTodoTaskResponse) SetBody(v *CreatePersonalTodoTaskResponseBody) *CreatePersonalTodoTaskResponse {
	s.Body = v
	return s
}

type CreateTodoTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateTodoTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskHeaders) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskHeaders) SetCommonHeaders(v map[string]*string) *CreateTodoTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTodoTaskHeaders) SetXAcsDingtalkAccessToken(v string) *CreateTodoTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateTodoTaskRequest struct {
	ActionList          []*CreateTodoTaskRequestActionList        `json:"actionList,omitempty" xml:"actionList,omitempty" type:"Repeated"`
	BizCategoryId       *string                                   `json:"bizCategoryId,omitempty" xml:"bizCategoryId,omitempty"`
	ContentFieldList    []*CreateTodoTaskRequestContentFieldList  `json:"contentFieldList,omitempty" xml:"contentFieldList,omitempty" type:"Repeated"`
	CreatorId           *string                                   `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description         *string                                   `json:"description,omitempty" xml:"description,omitempty"`
	DetailUrl           *CreateTodoTaskRequestDetailUrl           `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	DueTime             *int64                                    `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	ExecutorIds         []*string                                 `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	IsOnlyShowExecutor  *bool                                     `json:"isOnlyShowExecutor,omitempty" xml:"isOnlyShowExecutor,omitempty"`
	NotifyConfigs       *CreateTodoTaskRequestNotifyConfigs       `json:"notifyConfigs,omitempty" xml:"notifyConfigs,omitempty" type:"Struct"`
	ParticipantIds      []*string                                 `json:"participantIds,omitempty" xml:"participantIds,omitempty" type:"Repeated"`
	Priority            *int32                                    `json:"priority,omitempty" xml:"priority,omitempty"`
	RemindNotifyConfigs *CreateTodoTaskRequestRemindNotifyConfigs `json:"remindNotifyConfigs,omitempty" xml:"remindNotifyConfigs,omitempty" type:"Struct"`
	ReminderTimeStamp   *int64                                    `json:"reminderTimeStamp,omitempty" xml:"reminderTimeStamp,omitempty"`
	SourceId            *string                                   `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// This parameter is required.
	Subject        *string                `json:"subject,omitempty" xml:"subject,omitempty"`
	ThirdExtension map[string]interface{} `json:"thirdExtension,omitempty" xml:"thirdExtension,omitempty"`
	// example:
	//
	// TODO
	TodoType   *string `json:"todoType,omitempty" xml:"todoType,omitempty"`
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s CreateTodoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskRequest) SetActionList(v []*CreateTodoTaskRequestActionList) *CreateTodoTaskRequest {
	s.ActionList = v
	return s
}

func (s *CreateTodoTaskRequest) SetBizCategoryId(v string) *CreateTodoTaskRequest {
	s.BizCategoryId = &v
	return s
}

func (s *CreateTodoTaskRequest) SetContentFieldList(v []*CreateTodoTaskRequestContentFieldList) *CreateTodoTaskRequest {
	s.ContentFieldList = v
	return s
}

func (s *CreateTodoTaskRequest) SetCreatorId(v string) *CreateTodoTaskRequest {
	s.CreatorId = &v
	return s
}

func (s *CreateTodoTaskRequest) SetDescription(v string) *CreateTodoTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateTodoTaskRequest) SetDetailUrl(v *CreateTodoTaskRequestDetailUrl) *CreateTodoTaskRequest {
	s.DetailUrl = v
	return s
}

func (s *CreateTodoTaskRequest) SetDueTime(v int64) *CreateTodoTaskRequest {
	s.DueTime = &v
	return s
}

func (s *CreateTodoTaskRequest) SetExecutorIds(v []*string) *CreateTodoTaskRequest {
	s.ExecutorIds = v
	return s
}

func (s *CreateTodoTaskRequest) SetIsOnlyShowExecutor(v bool) *CreateTodoTaskRequest {
	s.IsOnlyShowExecutor = &v
	return s
}

func (s *CreateTodoTaskRequest) SetNotifyConfigs(v *CreateTodoTaskRequestNotifyConfigs) *CreateTodoTaskRequest {
	s.NotifyConfigs = v
	return s
}

func (s *CreateTodoTaskRequest) SetParticipantIds(v []*string) *CreateTodoTaskRequest {
	s.ParticipantIds = v
	return s
}

func (s *CreateTodoTaskRequest) SetPriority(v int32) *CreateTodoTaskRequest {
	s.Priority = &v
	return s
}

func (s *CreateTodoTaskRequest) SetRemindNotifyConfigs(v *CreateTodoTaskRequestRemindNotifyConfigs) *CreateTodoTaskRequest {
	s.RemindNotifyConfigs = v
	return s
}

func (s *CreateTodoTaskRequest) SetReminderTimeStamp(v int64) *CreateTodoTaskRequest {
	s.ReminderTimeStamp = &v
	return s
}

func (s *CreateTodoTaskRequest) SetSourceId(v string) *CreateTodoTaskRequest {
	s.SourceId = &v
	return s
}

func (s *CreateTodoTaskRequest) SetSubject(v string) *CreateTodoTaskRequest {
	s.Subject = &v
	return s
}

func (s *CreateTodoTaskRequest) SetThirdExtension(v map[string]interface{}) *CreateTodoTaskRequest {
	s.ThirdExtension = v
	return s
}

func (s *CreateTodoTaskRequest) SetTodoType(v string) *CreateTodoTaskRequest {
	s.TodoType = &v
	return s
}

func (s *CreateTodoTaskRequest) SetOperatorId(v string) *CreateTodoTaskRequest {
	s.OperatorId = &v
	return s
}

type CreateTodoTaskRequestActionList struct {
	ActionKey  *string                               `json:"actionKey,omitempty" xml:"actionKey,omitempty"`
	ActionType *int32                                `json:"actionType,omitempty" xml:"actionType,omitempty"`
	Param      *CreateTodoTaskRequestActionListParam `json:"param,omitempty" xml:"param,omitempty" type:"Struct"`
	PcUrl      *string                               `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	Title      *string                               `json:"title,omitempty" xml:"title,omitempty"`
	Url        *string                               `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateTodoTaskRequestActionList) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskRequestActionList) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskRequestActionList) SetActionKey(v string) *CreateTodoTaskRequestActionList {
	s.ActionKey = &v
	return s
}

func (s *CreateTodoTaskRequestActionList) SetActionType(v int32) *CreateTodoTaskRequestActionList {
	s.ActionType = &v
	return s
}

func (s *CreateTodoTaskRequestActionList) SetParam(v *CreateTodoTaskRequestActionListParam) *CreateTodoTaskRequestActionList {
	s.Param = v
	return s
}

func (s *CreateTodoTaskRequestActionList) SetPcUrl(v string) *CreateTodoTaskRequestActionList {
	s.PcUrl = &v
	return s
}

func (s *CreateTodoTaskRequestActionList) SetTitle(v string) *CreateTodoTaskRequestActionList {
	s.Title = &v
	return s
}

func (s *CreateTodoTaskRequestActionList) SetUrl(v string) *CreateTodoTaskRequestActionList {
	s.Url = &v
	return s
}

type CreateTodoTaskRequestActionListParam struct {
	Body   *string                `json:"body,omitempty" xml:"body,omitempty"`
	Header map[string]interface{} `json:"header,omitempty" xml:"header,omitempty"`
}

func (s CreateTodoTaskRequestActionListParam) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskRequestActionListParam) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskRequestActionListParam) SetBody(v string) *CreateTodoTaskRequestActionListParam {
	s.Body = &v
	return s
}

func (s *CreateTodoTaskRequestActionListParam) SetHeader(v map[string]interface{}) *CreateTodoTaskRequestActionListParam {
	s.Header = v
	return s
}

type CreateTodoTaskRequestContentFieldList struct {
	FieldKey   *string `json:"fieldKey,omitempty" xml:"fieldKey,omitempty"`
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
}

func (s CreateTodoTaskRequestContentFieldList) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskRequestContentFieldList) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskRequestContentFieldList) SetFieldKey(v string) *CreateTodoTaskRequestContentFieldList {
	s.FieldKey = &v
	return s
}

func (s *CreateTodoTaskRequestContentFieldList) SetFieldValue(v string) *CreateTodoTaskRequestContentFieldList {
	s.FieldValue = &v
	return s
}

type CreateTodoTaskRequestDetailUrl struct {
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
	PcUrl  *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
}

func (s CreateTodoTaskRequestDetailUrl) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskRequestDetailUrl) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskRequestDetailUrl) SetAppUrl(v string) *CreateTodoTaskRequestDetailUrl {
	s.AppUrl = &v
	return s
}

func (s *CreateTodoTaskRequestDetailUrl) SetPcUrl(v string) *CreateTodoTaskRequestDetailUrl {
	s.PcUrl = &v
	return s
}

type CreateTodoTaskRequestNotifyConfigs struct {
	DingNotify        *string `json:"dingNotify,omitempty" xml:"dingNotify,omitempty"`
	SendAssistantChat *string `json:"sendAssistantChat,omitempty" xml:"sendAssistantChat,omitempty"`
	SendTodoApn       *string `json:"sendTodoApn,omitempty" xml:"sendTodoApn,omitempty"`
}

func (s CreateTodoTaskRequestNotifyConfigs) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskRequestNotifyConfigs) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskRequestNotifyConfigs) SetDingNotify(v string) *CreateTodoTaskRequestNotifyConfigs {
	s.DingNotify = &v
	return s
}

func (s *CreateTodoTaskRequestNotifyConfigs) SetSendAssistantChat(v string) *CreateTodoTaskRequestNotifyConfigs {
	s.SendAssistantChat = &v
	return s
}

func (s *CreateTodoTaskRequestNotifyConfigs) SetSendTodoApn(v string) *CreateTodoTaskRequestNotifyConfigs {
	s.SendTodoApn = &v
	return s
}

type CreateTodoTaskRequestRemindNotifyConfigs struct {
	DingNotify  *string `json:"dingNotify,omitempty" xml:"dingNotify,omitempty"`
	SendTodoApn *string `json:"sendTodoApn,omitempty" xml:"sendTodoApn,omitempty"`
}

func (s CreateTodoTaskRequestRemindNotifyConfigs) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskRequestRemindNotifyConfigs) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskRequestRemindNotifyConfigs) SetDingNotify(v string) *CreateTodoTaskRequestRemindNotifyConfigs {
	s.DingNotify = &v
	return s
}

func (s *CreateTodoTaskRequestRemindNotifyConfigs) SetSendTodoApn(v string) *CreateTodoTaskRequestRemindNotifyConfigs {
	s.SendTodoApn = &v
	return s
}

type CreateTodoTaskResponseBody struct {
	BizTag           *string                                       `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	ContentFieldList []*CreateTodoTaskResponseBodyContentFieldList `json:"contentFieldList,omitempty" xml:"contentFieldList,omitempty" type:"Repeated"`
	// This parameter is required.
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// This parameter is required.
	CreatorId   *string                              `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description *string                              `json:"description,omitempty" xml:"description,omitempty"`
	DetailUrl   *CreateTodoTaskResponseBodyDetailUrl `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	Done        *bool                                `json:"done,omitempty" xml:"done,omitempty"`
	DueTime     *int64                               `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	ExecutorIds []*string                            `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	FinishTime  *int64                               `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// This parameter is required.
	Id                 *string                                  `json:"id,omitempty" xml:"id,omitempty"`
	IsOnlyShowExecutor *bool                                    `json:"isOnlyShowExecutor,omitempty" xml:"isOnlyShowExecutor,omitempty"`
	ModifiedTime       *int64                                   `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId         *string                                  `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	NotifyConfigs      *CreateTodoTaskResponseBodyNotifyConfigs `json:"notifyConfigs,omitempty" xml:"notifyConfigs,omitempty" type:"Struct"`
	ParticipantIds     []*string                                `json:"participantIds,omitempty" xml:"participantIds,omitempty" type:"Repeated"`
	Priority           *int32                                   `json:"priority,omitempty" xml:"priority,omitempty"`
	RequestId          *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Source             *string                                  `json:"source,omitempty" xml:"source,omitempty"`
	SourceId           *string                                  `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	StartTime          *int64                                   `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// This parameter is required.
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
}

func (s CreateTodoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskResponseBody) SetBizTag(v string) *CreateTodoTaskResponseBody {
	s.BizTag = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetContentFieldList(v []*CreateTodoTaskResponseBodyContentFieldList) *CreateTodoTaskResponseBody {
	s.ContentFieldList = v
	return s
}

func (s *CreateTodoTaskResponseBody) SetCreatedTime(v int64) *CreateTodoTaskResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetCreatorId(v string) *CreateTodoTaskResponseBody {
	s.CreatorId = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetDescription(v string) *CreateTodoTaskResponseBody {
	s.Description = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetDetailUrl(v *CreateTodoTaskResponseBodyDetailUrl) *CreateTodoTaskResponseBody {
	s.DetailUrl = v
	return s
}

func (s *CreateTodoTaskResponseBody) SetDone(v bool) *CreateTodoTaskResponseBody {
	s.Done = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetDueTime(v int64) *CreateTodoTaskResponseBody {
	s.DueTime = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetExecutorIds(v []*string) *CreateTodoTaskResponseBody {
	s.ExecutorIds = v
	return s
}

func (s *CreateTodoTaskResponseBody) SetFinishTime(v int64) *CreateTodoTaskResponseBody {
	s.FinishTime = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetId(v string) *CreateTodoTaskResponseBody {
	s.Id = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetIsOnlyShowExecutor(v bool) *CreateTodoTaskResponseBody {
	s.IsOnlyShowExecutor = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetModifiedTime(v int64) *CreateTodoTaskResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetModifierId(v string) *CreateTodoTaskResponseBody {
	s.ModifierId = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetNotifyConfigs(v *CreateTodoTaskResponseBodyNotifyConfigs) *CreateTodoTaskResponseBody {
	s.NotifyConfigs = v
	return s
}

func (s *CreateTodoTaskResponseBody) SetParticipantIds(v []*string) *CreateTodoTaskResponseBody {
	s.ParticipantIds = v
	return s
}

func (s *CreateTodoTaskResponseBody) SetPriority(v int32) *CreateTodoTaskResponseBody {
	s.Priority = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetRequestId(v string) *CreateTodoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetSource(v string) *CreateTodoTaskResponseBody {
	s.Source = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetSourceId(v string) *CreateTodoTaskResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetStartTime(v int64) *CreateTodoTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetSubject(v string) *CreateTodoTaskResponseBody {
	s.Subject = &v
	return s
}

type CreateTodoTaskResponseBodyContentFieldList struct {
	FieldKey   *string `json:"fieldKey,omitempty" xml:"fieldKey,omitempty"`
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
}

func (s CreateTodoTaskResponseBodyContentFieldList) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskResponseBodyContentFieldList) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskResponseBodyContentFieldList) SetFieldKey(v string) *CreateTodoTaskResponseBodyContentFieldList {
	s.FieldKey = &v
	return s
}

func (s *CreateTodoTaskResponseBodyContentFieldList) SetFieldValue(v string) *CreateTodoTaskResponseBodyContentFieldList {
	s.FieldValue = &v
	return s
}

type CreateTodoTaskResponseBodyDetailUrl struct {
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
	PcUrl  *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
}

func (s CreateTodoTaskResponseBodyDetailUrl) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskResponseBodyDetailUrl) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskResponseBodyDetailUrl) SetAppUrl(v string) *CreateTodoTaskResponseBodyDetailUrl {
	s.AppUrl = &v
	return s
}

func (s *CreateTodoTaskResponseBodyDetailUrl) SetPcUrl(v string) *CreateTodoTaskResponseBodyDetailUrl {
	s.PcUrl = &v
	return s
}

type CreateTodoTaskResponseBodyNotifyConfigs struct {
	DingNotify *string `json:"dingNotify,omitempty" xml:"dingNotify,omitempty"`
}

func (s CreateTodoTaskResponseBodyNotifyConfigs) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskResponseBodyNotifyConfigs) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskResponseBodyNotifyConfigs) SetDingNotify(v string) *CreateTodoTaskResponseBodyNotifyConfigs {
	s.DingNotify = &v
	return s
}

type CreateTodoTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTodoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTodoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskResponse) SetHeaders(v map[string]*string) *CreateTodoTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateTodoTaskResponse) SetStatusCode(v int32) *CreateTodoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTodoTaskResponse) SetBody(v *CreateTodoTaskResponseBody) *CreateTodoTaskResponse {
	s.Body = v
	return s
}

type CreateTodoTypeConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateTodoTypeConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTypeConfigHeaders) GoString() string {
	return s.String()
}

func (s *CreateTodoTypeConfigHeaders) SetCommonHeaders(v map[string]*string) *CreateTodoTypeConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTodoTypeConfigHeaders) SetXAcsDingtalkAccessToken(v string) *CreateTodoTypeConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateTodoTypeConfigRequest struct {
	ActionList []*CreateTodoTypeConfigRequestActionList `json:"actionList,omitempty" xml:"actionList,omitempty" type:"Repeated"`
	// This parameter is required.
	CardType         *int32                                         `json:"cardType,omitempty" xml:"cardType,omitempty"`
	ContentFieldList []*CreateTodoTypeConfigRequestContentFieldList `json:"contentFieldList,omitempty" xml:"contentFieldList,omitempty" type:"Repeated"`
	Description      *string                                        `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// This parameter is required.
	PcDetailUrlOpenMode *string `json:"pcDetailUrlOpenMode,omitempty" xml:"pcDetailUrlOpenMode,omitempty"`
	OperatorId          *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s CreateTodoTypeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTypeConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateTodoTypeConfigRequest) SetActionList(v []*CreateTodoTypeConfigRequestActionList) *CreateTodoTypeConfigRequest {
	s.ActionList = v
	return s
}

func (s *CreateTodoTypeConfigRequest) SetCardType(v int32) *CreateTodoTypeConfigRequest {
	s.CardType = &v
	return s
}

func (s *CreateTodoTypeConfigRequest) SetContentFieldList(v []*CreateTodoTypeConfigRequestContentFieldList) *CreateTodoTypeConfigRequest {
	s.ContentFieldList = v
	return s
}

func (s *CreateTodoTypeConfigRequest) SetDescription(v string) *CreateTodoTypeConfigRequest {
	s.Description = &v
	return s
}

func (s *CreateTodoTypeConfigRequest) SetIcon(v string) *CreateTodoTypeConfigRequest {
	s.Icon = &v
	return s
}

func (s *CreateTodoTypeConfigRequest) SetPcDetailUrlOpenMode(v string) *CreateTodoTypeConfigRequest {
	s.PcDetailUrlOpenMode = &v
	return s
}

func (s *CreateTodoTypeConfigRequest) SetOperatorId(v string) *CreateTodoTypeConfigRequest {
	s.OperatorId = &v
	return s
}

type CreateTodoTypeConfigRequestActionList struct {
	ActionKey       *string                `json:"actionKey,omitempty" xml:"actionKey,omitempty"`
	ActionType      *int32                 `json:"actionType,omitempty" xml:"actionType,omitempty"`
	ButtonStyleType *int32                 `json:"buttonStyleType,omitempty" xml:"buttonStyleType,omitempty"`
	NameI18n        map[string]interface{} `json:"nameI18n,omitempty" xml:"nameI18n,omitempty"`
	Url             *string                `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateTodoTypeConfigRequestActionList) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTypeConfigRequestActionList) GoString() string {
	return s.String()
}

func (s *CreateTodoTypeConfigRequestActionList) SetActionKey(v string) *CreateTodoTypeConfigRequestActionList {
	s.ActionKey = &v
	return s
}

func (s *CreateTodoTypeConfigRequestActionList) SetActionType(v int32) *CreateTodoTypeConfigRequestActionList {
	s.ActionType = &v
	return s
}

func (s *CreateTodoTypeConfigRequestActionList) SetButtonStyleType(v int32) *CreateTodoTypeConfigRequestActionList {
	s.ButtonStyleType = &v
	return s
}

func (s *CreateTodoTypeConfigRequestActionList) SetNameI18n(v map[string]interface{}) *CreateTodoTypeConfigRequestActionList {
	s.NameI18n = v
	return s
}

func (s *CreateTodoTypeConfigRequestActionList) SetUrl(v string) *CreateTodoTypeConfigRequestActionList {
	s.Url = &v
	return s
}

type CreateTodoTypeConfigRequestContentFieldList struct {
	FieldKey  *string                `json:"fieldKey,omitempty" xml:"fieldKey,omitempty"`
	FieldType *string                `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	NameI18n  map[string]interface{} `json:"nameI18n,omitempty" xml:"nameI18n,omitempty"`
}

func (s CreateTodoTypeConfigRequestContentFieldList) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTypeConfigRequestContentFieldList) GoString() string {
	return s.String()
}

func (s *CreateTodoTypeConfigRequestContentFieldList) SetFieldKey(v string) *CreateTodoTypeConfigRequestContentFieldList {
	s.FieldKey = &v
	return s
}

func (s *CreateTodoTypeConfigRequestContentFieldList) SetFieldType(v string) *CreateTodoTypeConfigRequestContentFieldList {
	s.FieldType = &v
	return s
}

func (s *CreateTodoTypeConfigRequestContentFieldList) SetNameI18n(v map[string]interface{}) *CreateTodoTypeConfigRequestContentFieldList {
	s.NameI18n = v
	return s
}

type CreateTodoTypeConfigResponseBody struct {
	ActionList       []*CreateTodoTypeConfigResponseBodyActionList       `json:"actionList,omitempty" xml:"actionList,omitempty" type:"Repeated"`
	BizTag           *string                                             `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	CardType         *int32                                              `json:"cardType,omitempty" xml:"cardType,omitempty"`
	ContentFieldList []*CreateTodoTypeConfigResponseBodyContentFieldList `json:"contentFieldList,omitempty" xml:"contentFieldList,omitempty" type:"Repeated"`
	// This parameter is required.
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// This parameter is required.
	CreatorId   *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Icon        *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// This parameter is required.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	ModifiedTime *int64 `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// This parameter is required.
	ModifierId          *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	PcDetailUrlOpenMode *string `json:"pcDetailUrlOpenMode,omitempty" xml:"pcDetailUrlOpenMode,omitempty"`
	RequestId           *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateTodoTypeConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTypeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTodoTypeConfigResponseBody) SetActionList(v []*CreateTodoTypeConfigResponseBodyActionList) *CreateTodoTypeConfigResponseBody {
	s.ActionList = v
	return s
}

func (s *CreateTodoTypeConfigResponseBody) SetBizTag(v string) *CreateTodoTypeConfigResponseBody {
	s.BizTag = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBody) SetCardType(v int32) *CreateTodoTypeConfigResponseBody {
	s.CardType = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBody) SetContentFieldList(v []*CreateTodoTypeConfigResponseBodyContentFieldList) *CreateTodoTypeConfigResponseBody {
	s.ContentFieldList = v
	return s
}

func (s *CreateTodoTypeConfigResponseBody) SetCreatedTime(v int64) *CreateTodoTypeConfigResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBody) SetCreatorId(v string) *CreateTodoTypeConfigResponseBody {
	s.CreatorId = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBody) SetDescription(v string) *CreateTodoTypeConfigResponseBody {
	s.Description = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBody) SetIcon(v string) *CreateTodoTypeConfigResponseBody {
	s.Icon = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBody) SetId(v string) *CreateTodoTypeConfigResponseBody {
	s.Id = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBody) SetModifiedTime(v int64) *CreateTodoTypeConfigResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBody) SetModifierId(v string) *CreateTodoTypeConfigResponseBody {
	s.ModifierId = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBody) SetPcDetailUrlOpenMode(v string) *CreateTodoTypeConfigResponseBody {
	s.PcDetailUrlOpenMode = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBody) SetRequestId(v string) *CreateTodoTypeConfigResponseBody {
	s.RequestId = &v
	return s
}

type CreateTodoTypeConfigResponseBodyActionList struct {
	ActionKey       *string `json:"actionKey,omitempty" xml:"actionKey,omitempty"`
	ActionType      *int32  `json:"actionType,omitempty" xml:"actionType,omitempty"`
	ButtonStyleType *int32  `json:"buttonStyleType,omitempty" xml:"buttonStyleType,omitempty"`
	// This parameter is required.
	NameI18n map[string]interface{} `json:"nameI18n,omitempty" xml:"nameI18n,omitempty"`
	Url      *string                `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateTodoTypeConfigResponseBodyActionList) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTypeConfigResponseBodyActionList) GoString() string {
	return s.String()
}

func (s *CreateTodoTypeConfigResponseBodyActionList) SetActionKey(v string) *CreateTodoTypeConfigResponseBodyActionList {
	s.ActionKey = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBodyActionList) SetActionType(v int32) *CreateTodoTypeConfigResponseBodyActionList {
	s.ActionType = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBodyActionList) SetButtonStyleType(v int32) *CreateTodoTypeConfigResponseBodyActionList {
	s.ButtonStyleType = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBodyActionList) SetNameI18n(v map[string]interface{}) *CreateTodoTypeConfigResponseBodyActionList {
	s.NameI18n = v
	return s
}

func (s *CreateTodoTypeConfigResponseBodyActionList) SetUrl(v string) *CreateTodoTypeConfigResponseBodyActionList {
	s.Url = &v
	return s
}

type CreateTodoTypeConfigResponseBodyContentFieldList struct {
	FieldKey  *string                `json:"fieldKey,omitempty" xml:"fieldKey,omitempty"`
	FieldType *string                `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	NameI18n  map[string]interface{} `json:"nameI18n,omitempty" xml:"nameI18n,omitempty"`
}

func (s CreateTodoTypeConfigResponseBodyContentFieldList) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTypeConfigResponseBodyContentFieldList) GoString() string {
	return s.String()
}

func (s *CreateTodoTypeConfigResponseBodyContentFieldList) SetFieldKey(v string) *CreateTodoTypeConfigResponseBodyContentFieldList {
	s.FieldKey = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBodyContentFieldList) SetFieldType(v string) *CreateTodoTypeConfigResponseBodyContentFieldList {
	s.FieldType = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBodyContentFieldList) SetNameI18n(v map[string]interface{}) *CreateTodoTypeConfigResponseBodyContentFieldList {
	s.NameI18n = v
	return s
}

type CreateTodoTypeConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTodoTypeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTodoTypeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTypeConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateTodoTypeConfigResponse) SetHeaders(v map[string]*string) *CreateTodoTypeConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateTodoTypeConfigResponse) SetStatusCode(v int32) *CreateTodoTypeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTodoTypeConfigResponse) SetBody(v *CreateTodoTypeConfigResponseBody) *CreateTodoTypeConfigResponse {
	s.Body = v
	return s
}

type DeleteTodoTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteTodoTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteTodoTaskHeaders) GoString() string {
	return s.String()
}

func (s *DeleteTodoTaskHeaders) SetCommonHeaders(v map[string]*string) *DeleteTodoTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteTodoTaskHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteTodoTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteTodoTaskRequest struct {
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s DeleteTodoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTodoTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteTodoTaskRequest) SetOperatorId(v string) *DeleteTodoTaskRequest {
	s.OperatorId = &v
	return s
}

type DeleteTodoTaskResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteTodoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTodoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTodoTaskResponseBody) SetRequestId(v string) *DeleteTodoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTodoTaskResponseBody) SetResult(v bool) *DeleteTodoTaskResponseBody {
	s.Result = &v
	return s
}

type DeleteTodoTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTodoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTodoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTodoTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteTodoTaskResponse) SetHeaders(v map[string]*string) *DeleteTodoTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteTodoTaskResponse) SetStatusCode(v int32) *DeleteTodoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTodoTaskResponse) SetBody(v *DeleteTodoTaskResponseBody) *DeleteTodoTaskResponse {
	s.Body = v
	return s
}

type GetTodoTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTodoTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTaskHeaders) GoString() string {
	return s.String()
}

func (s *GetTodoTaskHeaders) SetCommonHeaders(v map[string]*string) *GetTodoTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTodoTaskHeaders) SetXAcsDingtalkAccessToken(v string) *GetTodoTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTodoTaskResponseBody struct {
	BizTag             *string                           `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	CardTypeId         *string                           `json:"cardTypeId,omitempty" xml:"cardTypeId,omitempty"`
	CreatedTime        *int64                            `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	CreatorId          *string                           `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description        *string                           `json:"description,omitempty" xml:"description,omitempty"`
	DetailUrl          *GetTodoTaskResponseBodyDetailUrl `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	Done               *bool                             `json:"done,omitempty" xml:"done,omitempty"`
	DueTime            *int64                            `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	ExecutorIds        []*string                         `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	FinishTime         *int64                            `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	Id                 *string                           `json:"id,omitempty" xml:"id,omitempty"`
	IsOnlyShowExecutor *bool                             `json:"isOnlyShowExecutor,omitempty" xml:"isOnlyShowExecutor,omitempty"`
	ModifiedTime       *int64                            `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId         *string                           `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	ParticipantIds     []*string                         `json:"participantIds,omitempty" xml:"participantIds,omitempty" type:"Repeated"`
	Priority           *int32                            `json:"priority,omitempty" xml:"priority,omitempty"`
	RequestId          *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Source             *string                           `json:"source,omitempty" xml:"source,omitempty"`
	SourceId           *string                           `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	StartTime          *int64                            `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Subject            *string                           `json:"subject,omitempty" xml:"subject,omitempty"`
	TenantId           *string                           `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	TenantType         *string                           `json:"tenantType,omitempty" xml:"tenantType,omitempty"`
}

func (s GetTodoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTodoTaskResponseBody) SetBizTag(v string) *GetTodoTaskResponseBody {
	s.BizTag = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetCardTypeId(v string) *GetTodoTaskResponseBody {
	s.CardTypeId = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetCreatedTime(v int64) *GetTodoTaskResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetCreatorId(v string) *GetTodoTaskResponseBody {
	s.CreatorId = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetDescription(v string) *GetTodoTaskResponseBody {
	s.Description = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetDetailUrl(v *GetTodoTaskResponseBodyDetailUrl) *GetTodoTaskResponseBody {
	s.DetailUrl = v
	return s
}

func (s *GetTodoTaskResponseBody) SetDone(v bool) *GetTodoTaskResponseBody {
	s.Done = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetDueTime(v int64) *GetTodoTaskResponseBody {
	s.DueTime = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetExecutorIds(v []*string) *GetTodoTaskResponseBody {
	s.ExecutorIds = v
	return s
}

func (s *GetTodoTaskResponseBody) SetFinishTime(v int64) *GetTodoTaskResponseBody {
	s.FinishTime = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetId(v string) *GetTodoTaskResponseBody {
	s.Id = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetIsOnlyShowExecutor(v bool) *GetTodoTaskResponseBody {
	s.IsOnlyShowExecutor = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetModifiedTime(v int64) *GetTodoTaskResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetModifierId(v string) *GetTodoTaskResponseBody {
	s.ModifierId = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetParticipantIds(v []*string) *GetTodoTaskResponseBody {
	s.ParticipantIds = v
	return s
}

func (s *GetTodoTaskResponseBody) SetPriority(v int32) *GetTodoTaskResponseBody {
	s.Priority = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetRequestId(v string) *GetTodoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetSource(v string) *GetTodoTaskResponseBody {
	s.Source = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetSourceId(v string) *GetTodoTaskResponseBody {
	s.SourceId = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetStartTime(v int64) *GetTodoTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetSubject(v string) *GetTodoTaskResponseBody {
	s.Subject = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetTenantId(v string) *GetTodoTaskResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetTenantType(v string) *GetTodoTaskResponseBody {
	s.TenantType = &v
	return s
}

type GetTodoTaskResponseBodyDetailUrl struct {
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
	PcUrl  *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
}

func (s GetTodoTaskResponseBodyDetailUrl) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTaskResponseBodyDetailUrl) GoString() string {
	return s.String()
}

func (s *GetTodoTaskResponseBodyDetailUrl) SetAppUrl(v string) *GetTodoTaskResponseBodyDetailUrl {
	s.AppUrl = &v
	return s
}

func (s *GetTodoTaskResponseBodyDetailUrl) SetPcUrl(v string) *GetTodoTaskResponseBodyDetailUrl {
	s.PcUrl = &v
	return s
}

type GetTodoTaskResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTodoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTodoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTodoTaskResponse) SetHeaders(v map[string]*string) *GetTodoTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTodoTaskResponse) SetStatusCode(v int32) *GetTodoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTodoTaskResponse) SetBody(v *GetTodoTaskResponseBody) *GetTodoTaskResponse {
	s.Body = v
	return s
}

type GetTodoTaskBySourceIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTodoTaskBySourceIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTaskBySourceIdHeaders) GoString() string {
	return s.String()
}

func (s *GetTodoTaskBySourceIdHeaders) SetCommonHeaders(v map[string]*string) *GetTodoTaskBySourceIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTodoTaskBySourceIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetTodoTaskBySourceIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTodoTaskBySourceIdResponseBody struct {
	BizTag             *string                                     `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	CreatedTime        *int64                                      `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	CreatorId          *string                                     `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description        *string                                     `json:"description,omitempty" xml:"description,omitempty"`
	DetailUrl          *GetTodoTaskBySourceIdResponseBodyDetailUrl `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	Done               *bool                                       `json:"done,omitempty" xml:"done,omitempty"`
	DueTime            *int64                                      `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	ExecutorIds        []*string                                   `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	FinishTime         *int64                                      `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	Id                 *string                                     `json:"id,omitempty" xml:"id,omitempty"`
	IsOnlyShowExecutor *bool                                       `json:"isOnlyShowExecutor,omitempty" xml:"isOnlyShowExecutor,omitempty"`
	ModifiedTime       *int64                                      `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId         *string                                     `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	ParticipantIds     []*string                                   `json:"participantIds,omitempty" xml:"participantIds,omitempty" type:"Repeated"`
	Priority           *int32                                      `json:"priority,omitempty" xml:"priority,omitempty"`
	RequestId          *string                                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Source             *string                                     `json:"source,omitempty" xml:"source,omitempty"`
	SourceId           *string                                     `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	StartTime          *int64                                      `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Subject            *string                                     `json:"subject,omitempty" xml:"subject,omitempty"`
	TenantId           *string                                     `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	TenantType         *string                                     `json:"tenantType,omitempty" xml:"tenantType,omitempty"`
}

func (s GetTodoTaskBySourceIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTaskBySourceIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetTodoTaskBySourceIdResponseBody) SetBizTag(v string) *GetTodoTaskBySourceIdResponseBody {
	s.BizTag = &v
	return s
}

func (s *GetTodoTaskBySourceIdResponseBody) SetCreatedTime(v int64) *GetTodoTaskBySourceIdResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetTodoTaskBySourceIdResponseBody) SetCreatorId(v string) *GetTodoTaskBySourceIdResponseBody {
	s.CreatorId = &v
	return s
}

func (s *GetTodoTaskBySourceIdResponseBody) SetDescription(v string) *GetTodoTaskBySourceIdResponseBody {
	s.Description = &v
	return s
}

func (s *GetTodoTaskBySourceIdResponseBody) SetDetailUrl(v *GetTodoTaskBySourceIdResponseBodyDetailUrl) *GetTodoTaskBySourceIdResponseBody {
	s.DetailUrl = v
	return s
}

func (s *GetTodoTaskBySourceIdResponseBody) SetDone(v bool) *GetTodoTaskBySourceIdResponseBody {
	s.Done = &v
	return s
}

func (s *GetTodoTaskBySourceIdResponseBody) SetDueTime(v int64) *GetTodoTaskBySourceIdResponseBody {
	s.DueTime = &v
	return s
}

func (s *GetTodoTaskBySourceIdResponseBody) SetExecutorIds(v []*string) *GetTodoTaskBySourceIdResponseBody {
	s.ExecutorIds = v
	return s
}

func (s *GetTodoTaskBySourceIdResponseBody) SetFinishTime(v int64) *GetTodoTaskBySourceIdResponseBody {
	s.FinishTime = &v
	return s
}

func (s *GetTodoTaskBySourceIdResponseBody) SetId(v string) *GetTodoTaskBySourceIdResponseBody {
	s.Id = &v
	return s
}

func (s *GetTodoTaskBySourceIdResponseBody) SetIsOnlyShowExecutor(v bool) *GetTodoTaskBySourceIdResponseBody {
	s.IsOnlyShowExecutor = &v
	return s
}

func (s *GetTodoTaskBySourceIdResponseBody) SetModifiedTime(v int64) *GetTodoTaskBySourceIdResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetTodoTaskBySourceIdResponseBody) SetModifierId(v string) *GetTodoTaskBySourceIdResponseBody {
	s.ModifierId = &v
	return s
}

func (s *GetTodoTaskBySourceIdResponseBody) SetParticipantIds(v []*string) *GetTodoTaskBySourceIdResponseBody {
	s.ParticipantIds = v
	return s
}

func (s *GetTodoTaskBySourceIdResponseBody) SetPriority(v int32) *GetTodoTaskBySourceIdResponseBody {
	s.Priority = &v
	return s
}

func (s *GetTodoTaskBySourceIdResponseBody) SetRequestId(v string) *GetTodoTaskBySourceIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTodoTaskBySourceIdResponseBody) SetSource(v string) *GetTodoTaskBySourceIdResponseBody {
	s.Source = &v
	return s
}

func (s *GetTodoTaskBySourceIdResponseBody) SetSourceId(v string) *GetTodoTaskBySourceIdResponseBody {
	s.SourceId = &v
	return s
}

func (s *GetTodoTaskBySourceIdResponseBody) SetStartTime(v int64) *GetTodoTaskBySourceIdResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetTodoTaskBySourceIdResponseBody) SetSubject(v string) *GetTodoTaskBySourceIdResponseBody {
	s.Subject = &v
	return s
}

func (s *GetTodoTaskBySourceIdResponseBody) SetTenantId(v string) *GetTodoTaskBySourceIdResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetTodoTaskBySourceIdResponseBody) SetTenantType(v string) *GetTodoTaskBySourceIdResponseBody {
	s.TenantType = &v
	return s
}

type GetTodoTaskBySourceIdResponseBodyDetailUrl struct {
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
	PcUrl  *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
}

func (s GetTodoTaskBySourceIdResponseBodyDetailUrl) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTaskBySourceIdResponseBodyDetailUrl) GoString() string {
	return s.String()
}

func (s *GetTodoTaskBySourceIdResponseBodyDetailUrl) SetAppUrl(v string) *GetTodoTaskBySourceIdResponseBodyDetailUrl {
	s.AppUrl = &v
	return s
}

func (s *GetTodoTaskBySourceIdResponseBodyDetailUrl) SetPcUrl(v string) *GetTodoTaskBySourceIdResponseBodyDetailUrl {
	s.PcUrl = &v
	return s
}

type GetTodoTaskBySourceIdResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTodoTaskBySourceIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTodoTaskBySourceIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTaskBySourceIdResponse) GoString() string {
	return s.String()
}

func (s *GetTodoTaskBySourceIdResponse) SetHeaders(v map[string]*string) *GetTodoTaskBySourceIdResponse {
	s.Headers = v
	return s
}

func (s *GetTodoTaskBySourceIdResponse) SetStatusCode(v int32) *GetTodoTaskBySourceIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTodoTaskBySourceIdResponse) SetBody(v *GetTodoTaskBySourceIdResponseBody) *GetTodoTaskBySourceIdResponse {
	s.Body = v
	return s
}

type GetTodoTaskDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTodoTaskDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTaskDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetTodoTaskDetailHeaders) SetCommonHeaders(v map[string]*string) *GetTodoTaskDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTodoTaskDetailHeaders) SetXAcsDingtalkAccessToken(v string) *GetTodoTaskDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTodoTaskDetailResponseBody struct {
	BizTag             *string                                        `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	Category           *string                                        `json:"category,omitempty" xml:"category,omitempty"`
	CreatedTime        *int64                                         `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	CreatorId          *string                                        `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description        *string                                        `json:"description,omitempty" xml:"description,omitempty"`
	DetailUrl          *GetTodoTaskDetailResponseBodyDetailUrl        `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	Done               *bool                                          `json:"done,omitempty" xml:"done,omitempty"`
	DueTime            *int64                                         `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	ExecutorIds        []*string                                      `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	ExecutorStatus     []*GetTodoTaskDetailResponseBodyExecutorStatus `json:"executorStatus,omitempty" xml:"executorStatus,omitempty" type:"Repeated"`
	FinishTime         *int64                                         `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	Id                 *string                                        `json:"id,omitempty" xml:"id,omitempty"`
	IsOnlyShowExecutor *bool                                          `json:"isOnlyShowExecutor,omitempty" xml:"isOnlyShowExecutor,omitempty"`
	ModifiedTime       *int64                                         `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId         *string                                        `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	OrgInfo            *GetTodoTaskDetailResponseBodyOrgInfo          `json:"orgInfo,omitempty" xml:"orgInfo,omitempty" type:"Struct"`
	ParticipantIds     []*string                                      `json:"participantIds,omitempty" xml:"participantIds,omitempty" type:"Repeated"`
	Priority           *int32                                         `json:"priority,omitempty" xml:"priority,omitempty"`
	RequestId          *string                                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Source             *string                                        `json:"source,omitempty" xml:"source,omitempty"`
	SourceId           *string                                        `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	StartTime          *int64                                         `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Subject            *string                                        `json:"subject,omitempty" xml:"subject,omitempty"`
	TenantId           *string                                        `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	TenantType         *string                                        `json:"tenantType,omitempty" xml:"tenantType,omitempty"`
	TodoCardView       *GetTodoTaskDetailResponseBodyTodoCardView     `json:"todoCardView,omitempty" xml:"todoCardView,omitempty" type:"Struct"`
}

func (s GetTodoTaskDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetTodoTaskDetailResponseBody) SetBizTag(v string) *GetTodoTaskDetailResponseBody {
	s.BizTag = &v
	return s
}

func (s *GetTodoTaskDetailResponseBody) SetCategory(v string) *GetTodoTaskDetailResponseBody {
	s.Category = &v
	return s
}

func (s *GetTodoTaskDetailResponseBody) SetCreatedTime(v int64) *GetTodoTaskDetailResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetTodoTaskDetailResponseBody) SetCreatorId(v string) *GetTodoTaskDetailResponseBody {
	s.CreatorId = &v
	return s
}

func (s *GetTodoTaskDetailResponseBody) SetDescription(v string) *GetTodoTaskDetailResponseBody {
	s.Description = &v
	return s
}

func (s *GetTodoTaskDetailResponseBody) SetDetailUrl(v *GetTodoTaskDetailResponseBodyDetailUrl) *GetTodoTaskDetailResponseBody {
	s.DetailUrl = v
	return s
}

func (s *GetTodoTaskDetailResponseBody) SetDone(v bool) *GetTodoTaskDetailResponseBody {
	s.Done = &v
	return s
}

func (s *GetTodoTaskDetailResponseBody) SetDueTime(v int64) *GetTodoTaskDetailResponseBody {
	s.DueTime = &v
	return s
}

func (s *GetTodoTaskDetailResponseBody) SetExecutorIds(v []*string) *GetTodoTaskDetailResponseBody {
	s.ExecutorIds = v
	return s
}

func (s *GetTodoTaskDetailResponseBody) SetExecutorStatus(v []*GetTodoTaskDetailResponseBodyExecutorStatus) *GetTodoTaskDetailResponseBody {
	s.ExecutorStatus = v
	return s
}

func (s *GetTodoTaskDetailResponseBody) SetFinishTime(v int64) *GetTodoTaskDetailResponseBody {
	s.FinishTime = &v
	return s
}

func (s *GetTodoTaskDetailResponseBody) SetId(v string) *GetTodoTaskDetailResponseBody {
	s.Id = &v
	return s
}

func (s *GetTodoTaskDetailResponseBody) SetIsOnlyShowExecutor(v bool) *GetTodoTaskDetailResponseBody {
	s.IsOnlyShowExecutor = &v
	return s
}

func (s *GetTodoTaskDetailResponseBody) SetModifiedTime(v int64) *GetTodoTaskDetailResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetTodoTaskDetailResponseBody) SetModifierId(v string) *GetTodoTaskDetailResponseBody {
	s.ModifierId = &v
	return s
}

func (s *GetTodoTaskDetailResponseBody) SetOrgInfo(v *GetTodoTaskDetailResponseBodyOrgInfo) *GetTodoTaskDetailResponseBody {
	s.OrgInfo = v
	return s
}

func (s *GetTodoTaskDetailResponseBody) SetParticipantIds(v []*string) *GetTodoTaskDetailResponseBody {
	s.ParticipantIds = v
	return s
}

func (s *GetTodoTaskDetailResponseBody) SetPriority(v int32) *GetTodoTaskDetailResponseBody {
	s.Priority = &v
	return s
}

func (s *GetTodoTaskDetailResponseBody) SetRequestId(v string) *GetTodoTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTodoTaskDetailResponseBody) SetSource(v string) *GetTodoTaskDetailResponseBody {
	s.Source = &v
	return s
}

func (s *GetTodoTaskDetailResponseBody) SetSourceId(v string) *GetTodoTaskDetailResponseBody {
	s.SourceId = &v
	return s
}

func (s *GetTodoTaskDetailResponseBody) SetStartTime(v int64) *GetTodoTaskDetailResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetTodoTaskDetailResponseBody) SetSubject(v string) *GetTodoTaskDetailResponseBody {
	s.Subject = &v
	return s
}

func (s *GetTodoTaskDetailResponseBody) SetTenantId(v string) *GetTodoTaskDetailResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetTodoTaskDetailResponseBody) SetTenantType(v string) *GetTodoTaskDetailResponseBody {
	s.TenantType = &v
	return s
}

func (s *GetTodoTaskDetailResponseBody) SetTodoCardView(v *GetTodoTaskDetailResponseBodyTodoCardView) *GetTodoTaskDetailResponseBody {
	s.TodoCardView = v
	return s
}

type GetTodoTaskDetailResponseBodyDetailUrl struct {
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
	PcUrl  *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
}

func (s GetTodoTaskDetailResponseBodyDetailUrl) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTaskDetailResponseBodyDetailUrl) GoString() string {
	return s.String()
}

func (s *GetTodoTaskDetailResponseBodyDetailUrl) SetAppUrl(v string) *GetTodoTaskDetailResponseBodyDetailUrl {
	s.AppUrl = &v
	return s
}

func (s *GetTodoTaskDetailResponseBodyDetailUrl) SetPcUrl(v string) *GetTodoTaskDetailResponseBodyDetailUrl {
	s.PcUrl = &v
	return s
}

type GetTodoTaskDetailResponseBodyExecutorStatus struct {
	IsDone *bool   `json:"isDone,omitempty" xml:"isDone,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetTodoTaskDetailResponseBodyExecutorStatus) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTaskDetailResponseBodyExecutorStatus) GoString() string {
	return s.String()
}

func (s *GetTodoTaskDetailResponseBodyExecutorStatus) SetIsDone(v bool) *GetTodoTaskDetailResponseBodyExecutorStatus {
	s.IsDone = &v
	return s
}

func (s *GetTodoTaskDetailResponseBodyExecutorStatus) SetUserId(v string) *GetTodoTaskDetailResponseBodyExecutorStatus {
	s.UserId = &v
	return s
}

type GetTodoTaskDetailResponseBodyOrgInfo struct {
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetTodoTaskDetailResponseBodyOrgInfo) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTaskDetailResponseBodyOrgInfo) GoString() string {
	return s.String()
}

func (s *GetTodoTaskDetailResponseBodyOrgInfo) SetCorpId(v string) *GetTodoTaskDetailResponseBodyOrgInfo {
	s.CorpId = &v
	return s
}

func (s *GetTodoTaskDetailResponseBodyOrgInfo) SetName(v string) *GetTodoTaskDetailResponseBodyOrgInfo {
	s.Name = &v
	return s
}

type GetTodoTaskDetailResponseBodyTodoCardView struct {
	ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// example:
	//
	// standard , nonStandard, 标准卡片和非标准卡片，非标准卡片由第三方接入方自定义
	CardType *string `json:"cardType,omitempty" xml:"cardType,omitempty"`
	// example:
	//
	// 是 icon, 或者checkbox 类型的
	CircleELType        *string                                                         `json:"circleELType,omitempty" xml:"circleELType,omitempty"`
	ContentType         *string                                                         `json:"contentType,omitempty" xml:"contentType,omitempty"`
	Icon                *string                                                         `json:"icon,omitempty" xml:"icon,omitempty"`
	TodoCardContentList []*GetTodoTaskDetailResponseBodyTodoCardViewTodoCardContentList `json:"todoCardContentList,omitempty" xml:"todoCardContentList,omitempty" type:"Repeated"`
	TodoCardTitle       *string                                                         `json:"todoCardTitle,omitempty" xml:"todoCardTitle,omitempty"`
}

func (s GetTodoTaskDetailResponseBodyTodoCardView) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTaskDetailResponseBodyTodoCardView) GoString() string {
	return s.String()
}

func (s *GetTodoTaskDetailResponseBodyTodoCardView) SetActionType(v string) *GetTodoTaskDetailResponseBodyTodoCardView {
	s.ActionType = &v
	return s
}

func (s *GetTodoTaskDetailResponseBodyTodoCardView) SetCardType(v string) *GetTodoTaskDetailResponseBodyTodoCardView {
	s.CardType = &v
	return s
}

func (s *GetTodoTaskDetailResponseBodyTodoCardView) SetCircleELType(v string) *GetTodoTaskDetailResponseBodyTodoCardView {
	s.CircleELType = &v
	return s
}

func (s *GetTodoTaskDetailResponseBodyTodoCardView) SetContentType(v string) *GetTodoTaskDetailResponseBodyTodoCardView {
	s.ContentType = &v
	return s
}

func (s *GetTodoTaskDetailResponseBodyTodoCardView) SetIcon(v string) *GetTodoTaskDetailResponseBodyTodoCardView {
	s.Icon = &v
	return s
}

func (s *GetTodoTaskDetailResponseBodyTodoCardView) SetTodoCardContentList(v []*GetTodoTaskDetailResponseBodyTodoCardViewTodoCardContentList) *GetTodoTaskDetailResponseBodyTodoCardView {
	s.TodoCardContentList = v
	return s
}

func (s *GetTodoTaskDetailResponseBodyTodoCardView) SetTodoCardTitle(v string) *GetTodoTaskDetailResponseBodyTodoCardView {
	s.TodoCardTitle = &v
	return s
}

type GetTodoTaskDetailResponseBodyTodoCardViewTodoCardContentList struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetTodoTaskDetailResponseBodyTodoCardViewTodoCardContentList) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTaskDetailResponseBodyTodoCardViewTodoCardContentList) GoString() string {
	return s.String()
}

func (s *GetTodoTaskDetailResponseBodyTodoCardViewTodoCardContentList) SetName(v string) *GetTodoTaskDetailResponseBodyTodoCardViewTodoCardContentList {
	s.Name = &v
	return s
}

func (s *GetTodoTaskDetailResponseBodyTodoCardViewTodoCardContentList) SetValue(v string) *GetTodoTaskDetailResponseBodyTodoCardViewTodoCardContentList {
	s.Value = &v
	return s
}

type GetTodoTaskDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTodoTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTodoTaskDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *GetTodoTaskDetailResponse) SetHeaders(v map[string]*string) *GetTodoTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *GetTodoTaskDetailResponse) SetStatusCode(v int32) *GetTodoTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTodoTaskDetailResponse) SetBody(v *GetTodoTaskDetailResponseBody) *GetTodoTaskDetailResponse {
	s.Body = v
	return s
}

type GetTodoTypeConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTodoTypeConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTypeConfigHeaders) GoString() string {
	return s.String()
}

func (s *GetTodoTypeConfigHeaders) SetCommonHeaders(v map[string]*string) *GetTodoTypeConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTodoTypeConfigHeaders) SetXAcsDingtalkAccessToken(v string) *GetTodoTypeConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTodoTypeConfigResponseBody struct {
	ActionList       []*GetTodoTypeConfigResponseBodyActionList       `json:"actionList,omitempty" xml:"actionList,omitempty" type:"Repeated"`
	BizTag           *string                                          `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	CardType         *int32                                           `json:"cardType,omitempty" xml:"cardType,omitempty"`
	ContentFieldList []*GetTodoTypeConfigResponseBodyContentFieldList `json:"contentFieldList,omitempty" xml:"contentFieldList,omitempty" type:"Repeated"`
	// This parameter is required.
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// This parameter is required.
	CreatorId   *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Icon        *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// This parameter is required.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	ModifiedTime *int64 `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// This parameter is required.
	ModifierId          *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	PcDetailUrlOpenMode *string `json:"pcDetailUrlOpenMode,omitempty" xml:"pcDetailUrlOpenMode,omitempty"`
	RequestId           *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetTodoTypeConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTypeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetTodoTypeConfigResponseBody) SetActionList(v []*GetTodoTypeConfigResponseBodyActionList) *GetTodoTypeConfigResponseBody {
	s.ActionList = v
	return s
}

func (s *GetTodoTypeConfigResponseBody) SetBizTag(v string) *GetTodoTypeConfigResponseBody {
	s.BizTag = &v
	return s
}

func (s *GetTodoTypeConfigResponseBody) SetCardType(v int32) *GetTodoTypeConfigResponseBody {
	s.CardType = &v
	return s
}

func (s *GetTodoTypeConfigResponseBody) SetContentFieldList(v []*GetTodoTypeConfigResponseBodyContentFieldList) *GetTodoTypeConfigResponseBody {
	s.ContentFieldList = v
	return s
}

func (s *GetTodoTypeConfigResponseBody) SetCreatedTime(v int64) *GetTodoTypeConfigResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetTodoTypeConfigResponseBody) SetCreatorId(v string) *GetTodoTypeConfigResponseBody {
	s.CreatorId = &v
	return s
}

func (s *GetTodoTypeConfigResponseBody) SetDescription(v string) *GetTodoTypeConfigResponseBody {
	s.Description = &v
	return s
}

func (s *GetTodoTypeConfigResponseBody) SetIcon(v string) *GetTodoTypeConfigResponseBody {
	s.Icon = &v
	return s
}

func (s *GetTodoTypeConfigResponseBody) SetId(v string) *GetTodoTypeConfigResponseBody {
	s.Id = &v
	return s
}

func (s *GetTodoTypeConfigResponseBody) SetModifiedTime(v int64) *GetTodoTypeConfigResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetTodoTypeConfigResponseBody) SetModifierId(v string) *GetTodoTypeConfigResponseBody {
	s.ModifierId = &v
	return s
}

func (s *GetTodoTypeConfigResponseBody) SetPcDetailUrlOpenMode(v string) *GetTodoTypeConfigResponseBody {
	s.PcDetailUrlOpenMode = &v
	return s
}

func (s *GetTodoTypeConfigResponseBody) SetRequestId(v string) *GetTodoTypeConfigResponseBody {
	s.RequestId = &v
	return s
}

type GetTodoTypeConfigResponseBodyActionList struct {
	ActionKey       *string `json:"actionKey,omitempty" xml:"actionKey,omitempty"`
	ActionType      *int32  `json:"actionType,omitempty" xml:"actionType,omitempty"`
	ButtonStyleType *int32  `json:"buttonStyleType,omitempty" xml:"buttonStyleType,omitempty"`
	// This parameter is required.
	NameI18n map[string]interface{} `json:"nameI18n,omitempty" xml:"nameI18n,omitempty"`
	Url      *string                `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetTodoTypeConfigResponseBodyActionList) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTypeConfigResponseBodyActionList) GoString() string {
	return s.String()
}

func (s *GetTodoTypeConfigResponseBodyActionList) SetActionKey(v string) *GetTodoTypeConfigResponseBodyActionList {
	s.ActionKey = &v
	return s
}

func (s *GetTodoTypeConfigResponseBodyActionList) SetActionType(v int32) *GetTodoTypeConfigResponseBodyActionList {
	s.ActionType = &v
	return s
}

func (s *GetTodoTypeConfigResponseBodyActionList) SetButtonStyleType(v int32) *GetTodoTypeConfigResponseBodyActionList {
	s.ButtonStyleType = &v
	return s
}

func (s *GetTodoTypeConfigResponseBodyActionList) SetNameI18n(v map[string]interface{}) *GetTodoTypeConfigResponseBodyActionList {
	s.NameI18n = v
	return s
}

func (s *GetTodoTypeConfigResponseBodyActionList) SetUrl(v string) *GetTodoTypeConfigResponseBodyActionList {
	s.Url = &v
	return s
}

type GetTodoTypeConfigResponseBodyContentFieldList struct {
	FieldKey  *string                `json:"fieldKey,omitempty" xml:"fieldKey,omitempty"`
	FieldType *string                `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	NameI18n  map[string]interface{} `json:"nameI18n,omitempty" xml:"nameI18n,omitempty"`
}

func (s GetTodoTypeConfigResponseBodyContentFieldList) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTypeConfigResponseBodyContentFieldList) GoString() string {
	return s.String()
}

func (s *GetTodoTypeConfigResponseBodyContentFieldList) SetFieldKey(v string) *GetTodoTypeConfigResponseBodyContentFieldList {
	s.FieldKey = &v
	return s
}

func (s *GetTodoTypeConfigResponseBodyContentFieldList) SetFieldType(v string) *GetTodoTypeConfigResponseBodyContentFieldList {
	s.FieldType = &v
	return s
}

func (s *GetTodoTypeConfigResponseBodyContentFieldList) SetNameI18n(v map[string]interface{}) *GetTodoTypeConfigResponseBodyContentFieldList {
	s.NameI18n = v
	return s
}

type GetTodoTypeConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTodoTypeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTodoTypeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTypeConfigResponse) GoString() string {
	return s.String()
}

func (s *GetTodoTypeConfigResponse) SetHeaders(v map[string]*string) *GetTodoTypeConfigResponse {
	s.Headers = v
	return s
}

func (s *GetTodoTypeConfigResponse) SetStatusCode(v int32) *GetTodoTypeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTodoTypeConfigResponse) SetBody(v *GetTodoTypeConfigResponseBody) *GetTodoTypeConfigResponse {
	s.Body = v
	return s
}

type ListAllBizCategoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListAllBizCategoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListAllBizCategoryHeaders) GoString() string {
	return s.String()
}

func (s *ListAllBizCategoryHeaders) SetCommonHeaders(v map[string]*string) *ListAllBizCategoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAllBizCategoryHeaders) SetXAcsDingtalkAccessToken(v string) *ListAllBizCategoryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListAllBizCategoryResponseBody struct {
	BizCategories []*ListAllBizCategoryResponseBodyBizCategories `json:"bizCategories,omitempty" xml:"bizCategories,omitempty" type:"Repeated"`
}

func (s ListAllBizCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAllBizCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllBizCategoryResponseBody) SetBizCategories(v []*ListAllBizCategoryResponseBodyBizCategories) *ListAllBizCategoryResponseBody {
	s.BizCategories = v
	return s
}

type ListAllBizCategoryResponseBodyBizCategories struct {
	BizCategoryId *string `json:"bizCategoryId,omitempty" xml:"bizCategoryId,omitempty"`
	CreatedTime   *int64  `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Description   *string `json:"description,omitempty" xml:"description,omitempty"`
	ModifiedTime  *int64  `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListAllBizCategoryResponseBodyBizCategories) String() string {
	return tea.Prettify(s)
}

func (s ListAllBizCategoryResponseBodyBizCategories) GoString() string {
	return s.String()
}

func (s *ListAllBizCategoryResponseBodyBizCategories) SetBizCategoryId(v string) *ListAllBizCategoryResponseBodyBizCategories {
	s.BizCategoryId = &v
	return s
}

func (s *ListAllBizCategoryResponseBodyBizCategories) SetCreatedTime(v int64) *ListAllBizCategoryResponseBodyBizCategories {
	s.CreatedTime = &v
	return s
}

func (s *ListAllBizCategoryResponseBodyBizCategories) SetDescription(v string) *ListAllBizCategoryResponseBodyBizCategories {
	s.Description = &v
	return s
}

func (s *ListAllBizCategoryResponseBodyBizCategories) SetModifiedTime(v int64) *ListAllBizCategoryResponseBodyBizCategories {
	s.ModifiedTime = &v
	return s
}

func (s *ListAllBizCategoryResponseBodyBizCategories) SetName(v string) *ListAllBizCategoryResponseBodyBizCategories {
	s.Name = &v
	return s
}

type ListAllBizCategoryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllBizCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllBizCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAllBizCategoryResponse) GoString() string {
	return s.String()
}

func (s *ListAllBizCategoryResponse) SetHeaders(v map[string]*string) *ListAllBizCategoryResponse {
	s.Headers = v
	return s
}

func (s *ListAllBizCategoryResponse) SetStatusCode(v int32) *ListAllBizCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllBizCategoryResponse) SetBody(v *ListAllBizCategoryResponseBody) *ListAllBizCategoryResponse {
	s.Body = v
	return s
}

type QueryOrgConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOrgConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgConfigHeaders) GoString() string {
	return s.String()
}

func (s *QueryOrgConfigHeaders) SetCommonHeaders(v map[string]*string) *QueryOrgConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOrgConfigHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOrgConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOrgConfigResponseBody struct {
	AppDisplayStyle  *string   `json:"appDisplayStyle,omitempty" xml:"appDisplayStyle,omitempty"`
	CreatedTime      *int64    `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	HomepageCatalogs []*string `json:"homepageCatalogs,omitempty" xml:"homepageCatalogs,omitempty" type:"Repeated"`
	ModifiedTime     *int64    `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	Status           *int32    `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryOrgConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrgConfigResponseBody) SetAppDisplayStyle(v string) *QueryOrgConfigResponseBody {
	s.AppDisplayStyle = &v
	return s
}

func (s *QueryOrgConfigResponseBody) SetCreatedTime(v int64) *QueryOrgConfigResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *QueryOrgConfigResponseBody) SetHomepageCatalogs(v []*string) *QueryOrgConfigResponseBody {
	s.HomepageCatalogs = v
	return s
}

func (s *QueryOrgConfigResponseBody) SetModifiedTime(v int64) *QueryOrgConfigResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *QueryOrgConfigResponseBody) SetStatus(v int32) *QueryOrgConfigResponseBody {
	s.Status = &v
	return s
}

type QueryOrgConfigResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOrgConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOrgConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryOrgConfigResponse) SetHeaders(v map[string]*string) *QueryOrgConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryOrgConfigResponse) SetStatusCode(v int32) *QueryOrgConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrgConfigResponse) SetBody(v *QueryOrgConfigResponseBody) *QueryOrgConfigResponse {
	s.Body = v
	return s
}

type QueryOrgTodoByUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOrgTodoByUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTodoByUserHeaders) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoByUserHeaders) SetCommonHeaders(v map[string]*string) *QueryOrgTodoByUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOrgTodoByUserHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOrgTodoByUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOrgTodoByUserRequest struct {
	FromDueTime      *int64      `json:"fromDueTime,omitempty" xml:"fromDueTime,omitempty"`
	IsDone           *bool       `json:"isDone,omitempty" xml:"isDone,omitempty"`
	MaxResults       *int32      `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NeedPersonalTodo *bool       `json:"needPersonalTodo,omitempty" xml:"needPersonalTodo,omitempty"`
	NextToken        *string     `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OrderBy          *string     `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection   *string     `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	RoleTypes        [][]*string `json:"roleTypes,omitempty" xml:"roleTypes,omitempty" type:"Repeated"`
	Subject          *string     `json:"subject,omitempty" xml:"subject,omitempty"`
	ToDueTime        *int64      `json:"toDueTime,omitempty" xml:"toDueTime,omitempty"`
	// example:
	//
	// TODO
	TodoType *string `json:"todoType,omitempty" xml:"todoType,omitempty"`
}

func (s QueryOrgTodoByUserRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTodoByUserRequest) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoByUserRequest) SetFromDueTime(v int64) *QueryOrgTodoByUserRequest {
	s.FromDueTime = &v
	return s
}

func (s *QueryOrgTodoByUserRequest) SetIsDone(v bool) *QueryOrgTodoByUserRequest {
	s.IsDone = &v
	return s
}

func (s *QueryOrgTodoByUserRequest) SetMaxResults(v int32) *QueryOrgTodoByUserRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryOrgTodoByUserRequest) SetNeedPersonalTodo(v bool) *QueryOrgTodoByUserRequest {
	s.NeedPersonalTodo = &v
	return s
}

func (s *QueryOrgTodoByUserRequest) SetNextToken(v string) *QueryOrgTodoByUserRequest {
	s.NextToken = &v
	return s
}

func (s *QueryOrgTodoByUserRequest) SetOrderBy(v string) *QueryOrgTodoByUserRequest {
	s.OrderBy = &v
	return s
}

func (s *QueryOrgTodoByUserRequest) SetOrderDirection(v string) *QueryOrgTodoByUserRequest {
	s.OrderDirection = &v
	return s
}

func (s *QueryOrgTodoByUserRequest) SetRoleTypes(v [][]*string) *QueryOrgTodoByUserRequest {
	s.RoleTypes = v
	return s
}

func (s *QueryOrgTodoByUserRequest) SetSubject(v string) *QueryOrgTodoByUserRequest {
	s.Subject = &v
	return s
}

func (s *QueryOrgTodoByUserRequest) SetToDueTime(v int64) *QueryOrgTodoByUserRequest {
	s.ToDueTime = &v
	return s
}

func (s *QueryOrgTodoByUserRequest) SetTodoType(v string) *QueryOrgTodoByUserRequest {
	s.TodoType = &v
	return s
}

type QueryOrgTodoByUserResponseBody struct {
	// This parameter is required.
	MaxResults *int32                                     `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string                                    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	TodoCards  []*QueryOrgTodoByUserResponseBodyTodoCards `json:"todoCards,omitempty" xml:"todoCards,omitempty" type:"Repeated"`
}

func (s QueryOrgTodoByUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTodoByUserResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoByUserResponseBody) SetMaxResults(v int32) *QueryOrgTodoByUserResponseBody {
	s.MaxResults = &v
	return s
}

func (s *QueryOrgTodoByUserResponseBody) SetNextToken(v string) *QueryOrgTodoByUserResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryOrgTodoByUserResponseBody) SetTodoCards(v []*QueryOrgTodoByUserResponseBodyTodoCards) *QueryOrgTodoByUserResponseBody {
	s.TodoCards = v
	return s
}

type QueryOrgTodoByUserResponseBodyTodoCards struct {
	BizTag       *string                                           `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	CreatedTime  *int64                                            `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	CreatorId    *string                                           `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	DetailUrl    *QueryOrgTodoByUserResponseBodyTodoCardsDetailUrl `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	DueTime      *int64                                            `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	IsDone       *bool                                             `json:"isDone,omitempty" xml:"isDone,omitempty"`
	ModifiedTime *int64                                            `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	Priority     *int32                                            `json:"priority,omitempty" xml:"priority,omitempty"`
	// example:
	//
	// 一个业务自解释的json格式的字符串
	SourceExt *string `json:"sourceExt,omitempty" xml:"sourceExt,omitempty"`
	SourceId  *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	Subject   *string `json:"subject,omitempty" xml:"subject,omitempty"`
	TaskId    *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// TODO
	TodoType *string `json:"todoType,omitempty" xml:"todoType,omitempty"`
}

func (s QueryOrgTodoByUserResponseBodyTodoCards) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTodoByUserResponseBodyTodoCards) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoByUserResponseBodyTodoCards) SetBizTag(v string) *QueryOrgTodoByUserResponseBodyTodoCards {
	s.BizTag = &v
	return s
}

func (s *QueryOrgTodoByUserResponseBodyTodoCards) SetCreatedTime(v int64) *QueryOrgTodoByUserResponseBodyTodoCards {
	s.CreatedTime = &v
	return s
}

func (s *QueryOrgTodoByUserResponseBodyTodoCards) SetCreatorId(v string) *QueryOrgTodoByUserResponseBodyTodoCards {
	s.CreatorId = &v
	return s
}

func (s *QueryOrgTodoByUserResponseBodyTodoCards) SetDetailUrl(v *QueryOrgTodoByUserResponseBodyTodoCardsDetailUrl) *QueryOrgTodoByUserResponseBodyTodoCards {
	s.DetailUrl = v
	return s
}

func (s *QueryOrgTodoByUserResponseBodyTodoCards) SetDueTime(v int64) *QueryOrgTodoByUserResponseBodyTodoCards {
	s.DueTime = &v
	return s
}

func (s *QueryOrgTodoByUserResponseBodyTodoCards) SetIsDone(v bool) *QueryOrgTodoByUserResponseBodyTodoCards {
	s.IsDone = &v
	return s
}

func (s *QueryOrgTodoByUserResponseBodyTodoCards) SetModifiedTime(v int64) *QueryOrgTodoByUserResponseBodyTodoCards {
	s.ModifiedTime = &v
	return s
}

func (s *QueryOrgTodoByUserResponseBodyTodoCards) SetPriority(v int32) *QueryOrgTodoByUserResponseBodyTodoCards {
	s.Priority = &v
	return s
}

func (s *QueryOrgTodoByUserResponseBodyTodoCards) SetSourceExt(v string) *QueryOrgTodoByUserResponseBodyTodoCards {
	s.SourceExt = &v
	return s
}

func (s *QueryOrgTodoByUserResponseBodyTodoCards) SetSourceId(v string) *QueryOrgTodoByUserResponseBodyTodoCards {
	s.SourceId = &v
	return s
}

func (s *QueryOrgTodoByUserResponseBodyTodoCards) SetSubject(v string) *QueryOrgTodoByUserResponseBodyTodoCards {
	s.Subject = &v
	return s
}

func (s *QueryOrgTodoByUserResponseBodyTodoCards) SetTaskId(v string) *QueryOrgTodoByUserResponseBodyTodoCards {
	s.TaskId = &v
	return s
}

func (s *QueryOrgTodoByUserResponseBodyTodoCards) SetTodoType(v string) *QueryOrgTodoByUserResponseBodyTodoCards {
	s.TodoType = &v
	return s
}

type QueryOrgTodoByUserResponseBodyTodoCardsDetailUrl struct {
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
	PcUrl  *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
}

func (s QueryOrgTodoByUserResponseBodyTodoCardsDetailUrl) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTodoByUserResponseBodyTodoCardsDetailUrl) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoByUserResponseBodyTodoCardsDetailUrl) SetAppUrl(v string) *QueryOrgTodoByUserResponseBodyTodoCardsDetailUrl {
	s.AppUrl = &v
	return s
}

func (s *QueryOrgTodoByUserResponseBodyTodoCardsDetailUrl) SetPcUrl(v string) *QueryOrgTodoByUserResponseBodyTodoCardsDetailUrl {
	s.PcUrl = &v
	return s
}

type QueryOrgTodoByUserResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOrgTodoByUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOrgTodoByUserResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTodoByUserResponse) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoByUserResponse) SetHeaders(v map[string]*string) *QueryOrgTodoByUserResponse {
	s.Headers = v
	return s
}

func (s *QueryOrgTodoByUserResponse) SetStatusCode(v int32) *QueryOrgTodoByUserResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrgTodoByUserResponse) SetBody(v *QueryOrgTodoByUserResponseBody) *QueryOrgTodoByUserResponse {
	s.Body = v
	return s
}

type QueryOrgTodoTasksHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOrgTodoTasksHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTodoTasksHeaders) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksHeaders) SetCommonHeaders(v map[string]*string) *QueryOrgTodoTasksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOrgTodoTasksHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOrgTodoTasksHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOrgTodoTasksRequest struct {
	IsDone    *bool       `json:"isDone,omitempty" xml:"isDone,omitempty"`
	NextToken *string     `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RoleTypes [][]*string `json:"roleTypes,omitempty" xml:"roleTypes,omitempty" type:"Repeated"`
	// example:
	//
	// TODO
	TodoType *string `json:"todoType,omitempty" xml:"todoType,omitempty"`
}

func (s QueryOrgTodoTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTodoTasksRequest) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksRequest) SetIsDone(v bool) *QueryOrgTodoTasksRequest {
	s.IsDone = &v
	return s
}

func (s *QueryOrgTodoTasksRequest) SetNextToken(v string) *QueryOrgTodoTasksRequest {
	s.NextToken = &v
	return s
}

func (s *QueryOrgTodoTasksRequest) SetRoleTypes(v [][]*string) *QueryOrgTodoTasksRequest {
	s.RoleTypes = v
	return s
}

func (s *QueryOrgTodoTasksRequest) SetTodoType(v string) *QueryOrgTodoTasksRequest {
	s.TodoType = &v
	return s
}

type QueryOrgTodoTasksResponseBody struct {
	NextToken *string                                   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	TodoCards []*QueryOrgTodoTasksResponseBodyTodoCards `json:"todoCards,omitempty" xml:"todoCards,omitempty" type:"Repeated"`
}

func (s QueryOrgTodoTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTodoTasksResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksResponseBody) SetNextToken(v string) *QueryOrgTodoTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBody) SetTodoCards(v []*QueryOrgTodoTasksResponseBodyTodoCards) *QueryOrgTodoTasksResponseBody {
	s.TodoCards = v
	return s
}

type QueryOrgTodoTasksResponseBodyTodoCards struct {
	BizTag       *string                                          `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	CreatedTime  *int64                                           `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	CreatorId    *string                                          `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	DetailUrl    *QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	DueTime      *int64                                           `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	IsDone       *bool                                            `json:"isDone,omitempty" xml:"isDone,omitempty"`
	ModifiedTime *int64                                           `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	Priority     *int32                                           `json:"priority,omitempty" xml:"priority,omitempty"`
	SourceId     *string                                          `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	Subject      *string                                          `json:"subject,omitempty" xml:"subject,omitempty"`
	TaskId       *string                                          `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// TODO
	TodoType *string `json:"todoType,omitempty" xml:"todoType,omitempty"`
}

func (s QueryOrgTodoTasksResponseBodyTodoCards) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTodoTasksResponseBodyTodoCards) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetBizTag(v string) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.BizTag = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetCreatedTime(v int64) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.CreatedTime = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetCreatorId(v string) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.CreatorId = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetDetailUrl(v *QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.DetailUrl = v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetDueTime(v int64) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.DueTime = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetIsDone(v bool) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.IsDone = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetModifiedTime(v int64) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.ModifiedTime = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetPriority(v int32) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.Priority = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetSourceId(v string) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.SourceId = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetSubject(v string) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.Subject = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetTaskId(v string) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.TaskId = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetTodoType(v string) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.TodoType = &v
	return s
}

type QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl struct {
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
	PcUrl  *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
}

func (s QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl) SetAppUrl(v string) *QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl {
	s.AppUrl = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl) SetPcUrl(v string) *QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl {
	s.PcUrl = &v
	return s
}

type QueryOrgTodoTasksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOrgTodoTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOrgTodoTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTodoTasksResponse) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksResponse) SetHeaders(v map[string]*string) *QueryOrgTodoTasksResponse {
	s.Headers = v
	return s
}

func (s *QueryOrgTodoTasksResponse) SetStatusCode(v int32) *QueryOrgTodoTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrgTodoTasksResponse) SetBody(v *QueryOrgTodoTasksResponseBody) *QueryOrgTodoTasksResponse {
	s.Body = v
	return s
}

type QueryTodoTasksHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryTodoTasksHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoTasksHeaders) GoString() string {
	return s.String()
}

func (s *QueryTodoTasksHeaders) SetCommonHeaders(v map[string]*string) *QueryTodoTasksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryTodoTasksHeaders) SetXAcsDingtalkAccessToken(v string) *QueryTodoTasksHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryTodoTasksRequest struct {
	Category       *string     `json:"category,omitempty" xml:"category,omitempty"`
	FromDueTime    *int64      `json:"fromDueTime,omitempty" xml:"fromDueTime,omitempty"`
	IsDone         *bool       `json:"isDone,omitempty" xml:"isDone,omitempty"`
	IsRecycled     *bool       `json:"isRecycled,omitempty" xml:"isRecycled,omitempty"`
	NextToken      *string     `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OrderBy        *string     `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string     `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	RoleTypes      [][]*string `json:"roleTypes,omitempty" xml:"roleTypes,omitempty" type:"Repeated"`
	ToDueTime      *int64      `json:"toDueTime,omitempty" xml:"toDueTime,omitempty"`
}

func (s QueryTodoTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoTasksRequest) GoString() string {
	return s.String()
}

func (s *QueryTodoTasksRequest) SetCategory(v string) *QueryTodoTasksRequest {
	s.Category = &v
	return s
}

func (s *QueryTodoTasksRequest) SetFromDueTime(v int64) *QueryTodoTasksRequest {
	s.FromDueTime = &v
	return s
}

func (s *QueryTodoTasksRequest) SetIsDone(v bool) *QueryTodoTasksRequest {
	s.IsDone = &v
	return s
}

func (s *QueryTodoTasksRequest) SetIsRecycled(v bool) *QueryTodoTasksRequest {
	s.IsRecycled = &v
	return s
}

func (s *QueryTodoTasksRequest) SetNextToken(v string) *QueryTodoTasksRequest {
	s.NextToken = &v
	return s
}

func (s *QueryTodoTasksRequest) SetOrderBy(v string) *QueryTodoTasksRequest {
	s.OrderBy = &v
	return s
}

func (s *QueryTodoTasksRequest) SetOrderDirection(v string) *QueryTodoTasksRequest {
	s.OrderDirection = &v
	return s
}

func (s *QueryTodoTasksRequest) SetRoleTypes(v [][]*string) *QueryTodoTasksRequest {
	s.RoleTypes = v
	return s
}

func (s *QueryTodoTasksRequest) SetToDueTime(v int64) *QueryTodoTasksRequest {
	s.ToDueTime = &v
	return s
}

type QueryTodoTasksResponseBody struct {
	NextToken  *string                                `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	TodoCards  []*QueryTodoTasksResponseBodyTodoCards `json:"todoCards,omitempty" xml:"todoCards,omitempty" type:"Repeated"`
	TotalCount *int32                                 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryTodoTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoTasksResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTodoTasksResponseBody) SetNextToken(v string) *QueryTodoTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryTodoTasksResponseBody) SetTodoCards(v []*QueryTodoTasksResponseBodyTodoCards) *QueryTodoTasksResponseBody {
	s.TodoCards = v
	return s
}

func (s *QueryTodoTasksResponseBody) SetTotalCount(v int32) *QueryTodoTasksResponseBody {
	s.TotalCount = &v
	return s
}

type QueryTodoTasksResponseBodyTodoCards struct {
	BizTag         *string                                            `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	Category       *string                                            `json:"category,omitempty" xml:"category,omitempty"`
	CreatedTime    *int64                                             `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	CreatorId      *string                                            `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	DetailUrl      *QueryTodoTasksResponseBodyTodoCardsDetailUrl      `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	DueTime        *int64                                             `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	IsDone         *bool                                              `json:"isDone,omitempty" xml:"isDone,omitempty"`
	ModifiedTime   *int64                                             `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	OrgInfo        *QueryTodoTasksResponseBodyTodoCardsOrgInfo        `json:"orgInfo,omitempty" xml:"orgInfo,omitempty" type:"Struct"`
	OriginalSource *QueryTodoTasksResponseBodyTodoCardsOriginalSource `json:"originalSource,omitempty" xml:"originalSource,omitempty" type:"Struct"`
	Priority       *int32                                             `json:"priority,omitempty" xml:"priority,omitempty"`
	SourceId       *string                                            `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	Subject        *string                                            `json:"subject,omitempty" xml:"subject,omitempty"`
	TaskId         *string                                            `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TodoCardView   *QueryTodoTasksResponseBodyTodoCardsTodoCardView   `json:"todoCardView,omitempty" xml:"todoCardView,omitempty" type:"Struct"`
	TodoStatus     *string                                            `json:"todoStatus,omitempty" xml:"todoStatus,omitempty"`
}

func (s QueryTodoTasksResponseBodyTodoCards) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoTasksResponseBodyTodoCards) GoString() string {
	return s.String()
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetBizTag(v string) *QueryTodoTasksResponseBodyTodoCards {
	s.BizTag = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetCategory(v string) *QueryTodoTasksResponseBodyTodoCards {
	s.Category = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetCreatedTime(v int64) *QueryTodoTasksResponseBodyTodoCards {
	s.CreatedTime = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetCreatorId(v string) *QueryTodoTasksResponseBodyTodoCards {
	s.CreatorId = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetDetailUrl(v *QueryTodoTasksResponseBodyTodoCardsDetailUrl) *QueryTodoTasksResponseBodyTodoCards {
	s.DetailUrl = v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetDueTime(v int64) *QueryTodoTasksResponseBodyTodoCards {
	s.DueTime = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetIsDone(v bool) *QueryTodoTasksResponseBodyTodoCards {
	s.IsDone = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetModifiedTime(v int64) *QueryTodoTasksResponseBodyTodoCards {
	s.ModifiedTime = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetOrgInfo(v *QueryTodoTasksResponseBodyTodoCardsOrgInfo) *QueryTodoTasksResponseBodyTodoCards {
	s.OrgInfo = v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetOriginalSource(v *QueryTodoTasksResponseBodyTodoCardsOriginalSource) *QueryTodoTasksResponseBodyTodoCards {
	s.OriginalSource = v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetPriority(v int32) *QueryTodoTasksResponseBodyTodoCards {
	s.Priority = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetSourceId(v string) *QueryTodoTasksResponseBodyTodoCards {
	s.SourceId = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetSubject(v string) *QueryTodoTasksResponseBodyTodoCards {
	s.Subject = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetTaskId(v string) *QueryTodoTasksResponseBodyTodoCards {
	s.TaskId = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetTodoCardView(v *QueryTodoTasksResponseBodyTodoCardsTodoCardView) *QueryTodoTasksResponseBodyTodoCards {
	s.TodoCardView = v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetTodoStatus(v string) *QueryTodoTasksResponseBodyTodoCards {
	s.TodoStatus = &v
	return s
}

type QueryTodoTasksResponseBodyTodoCardsDetailUrl struct {
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
	PcUrl  *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
}

func (s QueryTodoTasksResponseBodyTodoCardsDetailUrl) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoTasksResponseBodyTodoCardsDetailUrl) GoString() string {
	return s.String()
}

func (s *QueryTodoTasksResponseBodyTodoCardsDetailUrl) SetAppUrl(v string) *QueryTodoTasksResponseBodyTodoCardsDetailUrl {
	s.AppUrl = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCardsDetailUrl) SetPcUrl(v string) *QueryTodoTasksResponseBodyTodoCardsDetailUrl {
	s.PcUrl = &v
	return s
}

type QueryTodoTasksResponseBodyTodoCardsOrgInfo struct {
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryTodoTasksResponseBodyTodoCardsOrgInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoTasksResponseBodyTodoCardsOrgInfo) GoString() string {
	return s.String()
}

func (s *QueryTodoTasksResponseBodyTodoCardsOrgInfo) SetCorpId(v string) *QueryTodoTasksResponseBodyTodoCardsOrgInfo {
	s.CorpId = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCardsOrgInfo) SetName(v string) *QueryTodoTasksResponseBodyTodoCardsOrgInfo {
	s.Name = &v
	return s
}

type QueryTodoTasksResponseBodyTodoCardsOriginalSource struct {
	SourceTitle *string `json:"sourceTitle,omitempty" xml:"sourceTitle,omitempty"`
}

func (s QueryTodoTasksResponseBodyTodoCardsOriginalSource) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoTasksResponseBodyTodoCardsOriginalSource) GoString() string {
	return s.String()
}

func (s *QueryTodoTasksResponseBodyTodoCardsOriginalSource) SetSourceTitle(v string) *QueryTodoTasksResponseBodyTodoCardsOriginalSource {
	s.SourceTitle = &v
	return s
}

type QueryTodoTasksResponseBodyTodoCardsTodoCardView struct {
	ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// example:
	//
	// standard , nonStandard, 标准卡片和非标准卡片，非标准卡片由第三方接入方自定义
	CardType *string `json:"cardType,omitempty" xml:"cardType,omitempty"`
	// example:
	//
	// 是 icon, 或者checkbox 类型的
	CircleELType        *string                                                               `json:"circleELType,omitempty" xml:"circleELType,omitempty"`
	ContentType         *string                                                               `json:"contentType,omitempty" xml:"contentType,omitempty"`
	Icon                *string                                                               `json:"icon,omitempty" xml:"icon,omitempty"`
	TodoCardContentList []*QueryTodoTasksResponseBodyTodoCardsTodoCardViewTodoCardContentList `json:"todoCardContentList,omitempty" xml:"todoCardContentList,omitempty" type:"Repeated"`
	TodoCardTitle       *string                                                               `json:"todoCardTitle,omitempty" xml:"todoCardTitle,omitempty"`
}

func (s QueryTodoTasksResponseBodyTodoCardsTodoCardView) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoTasksResponseBodyTodoCardsTodoCardView) GoString() string {
	return s.String()
}

func (s *QueryTodoTasksResponseBodyTodoCardsTodoCardView) SetActionType(v string) *QueryTodoTasksResponseBodyTodoCardsTodoCardView {
	s.ActionType = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCardsTodoCardView) SetCardType(v string) *QueryTodoTasksResponseBodyTodoCardsTodoCardView {
	s.CardType = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCardsTodoCardView) SetCircleELType(v string) *QueryTodoTasksResponseBodyTodoCardsTodoCardView {
	s.CircleELType = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCardsTodoCardView) SetContentType(v string) *QueryTodoTasksResponseBodyTodoCardsTodoCardView {
	s.ContentType = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCardsTodoCardView) SetIcon(v string) *QueryTodoTasksResponseBodyTodoCardsTodoCardView {
	s.Icon = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCardsTodoCardView) SetTodoCardContentList(v []*QueryTodoTasksResponseBodyTodoCardsTodoCardViewTodoCardContentList) *QueryTodoTasksResponseBodyTodoCardsTodoCardView {
	s.TodoCardContentList = v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCardsTodoCardView) SetTodoCardTitle(v string) *QueryTodoTasksResponseBodyTodoCardsTodoCardView {
	s.TodoCardTitle = &v
	return s
}

type QueryTodoTasksResponseBodyTodoCardsTodoCardViewTodoCardContentList struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s QueryTodoTasksResponseBodyTodoCardsTodoCardViewTodoCardContentList) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoTasksResponseBodyTodoCardsTodoCardViewTodoCardContentList) GoString() string {
	return s.String()
}

func (s *QueryTodoTasksResponseBodyTodoCardsTodoCardViewTodoCardContentList) SetName(v string) *QueryTodoTasksResponseBodyTodoCardsTodoCardViewTodoCardContentList {
	s.Name = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCardsTodoCardViewTodoCardContentList) SetValue(v string) *QueryTodoTasksResponseBodyTodoCardsTodoCardViewTodoCardContentList {
	s.Value = &v
	return s
}

type QueryTodoTasksResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTodoTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTodoTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoTasksResponse) GoString() string {
	return s.String()
}

func (s *QueryTodoTasksResponse) SetHeaders(v map[string]*string) *QueryTodoTasksResponse {
	s.Headers = v
	return s
}

func (s *QueryTodoTasksResponse) SetStatusCode(v int32) *QueryTodoTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTodoTasksResponse) SetBody(v *QueryTodoTasksResponseBody) *QueryTodoTasksResponse {
	s.Body = v
	return s
}

type RemoveBizCategoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RemoveBizCategoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveBizCategoryHeaders) GoString() string {
	return s.String()
}

func (s *RemoveBizCategoryHeaders) SetCommonHeaders(v map[string]*string) *RemoveBizCategoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveBizCategoryHeaders) SetXAcsDingtalkAccessToken(v string) *RemoveBizCategoryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RemoveBizCategoryRequest struct {
	// This parameter is required.
	BizCategoryId *string `json:"bizCategoryId,omitempty" xml:"bizCategoryId,omitempty"`
	// This parameter is required.
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s RemoveBizCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveBizCategoryRequest) GoString() string {
	return s.String()
}

func (s *RemoveBizCategoryRequest) SetBizCategoryId(v string) *RemoveBizCategoryRequest {
	s.BizCategoryId = &v
	return s
}

func (s *RemoveBizCategoryRequest) SetOperatorId(v string) *RemoveBizCategoryRequest {
	s.OperatorId = &v
	return s
}

type RemoveBizCategoryResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RemoveBizCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveBizCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveBizCategoryResponseBody) SetSuccess(v bool) *RemoveBizCategoryResponseBody {
	s.Success = &v
	return s
}

type RemoveBizCategoryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveBizCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveBizCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveBizCategoryResponse) GoString() string {
	return s.String()
}

func (s *RemoveBizCategoryResponse) SetHeaders(v map[string]*string) *RemoveBizCategoryResponse {
	s.Headers = v
	return s
}

func (s *RemoveBizCategoryResponse) SetStatusCode(v int32) *RemoveBizCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveBizCategoryResponse) SetBody(v *RemoveBizCategoryResponseBody) *RemoveBizCategoryResponse {
	s.Body = v
	return s
}

type SetBizCategoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetBizCategoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetBizCategoryHeaders) GoString() string {
	return s.String()
}

func (s *SetBizCategoryHeaders) SetCommonHeaders(v map[string]*string) *SetBizCategoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetBizCategoryHeaders) SetXAcsDingtalkAccessToken(v string) *SetBizCategoryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetBizCategoryRequest struct {
	// This parameter is required.
	BizCategoryId *string `json:"bizCategoryId,omitempty" xml:"bizCategoryId,omitempty"`
	Description   *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s SetBizCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s SetBizCategoryRequest) GoString() string {
	return s.String()
}

func (s *SetBizCategoryRequest) SetBizCategoryId(v string) *SetBizCategoryRequest {
	s.BizCategoryId = &v
	return s
}

func (s *SetBizCategoryRequest) SetDescription(v string) *SetBizCategoryRequest {
	s.Description = &v
	return s
}

func (s *SetBizCategoryRequest) SetName(v string) *SetBizCategoryRequest {
	s.Name = &v
	return s
}

func (s *SetBizCategoryRequest) SetOperatorId(v string) *SetBizCategoryRequest {
	s.OperatorId = &v
	return s
}

type SetBizCategoryResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SetBizCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetBizCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *SetBizCategoryResponseBody) SetSuccess(v bool) *SetBizCategoryResponseBody {
	s.Success = &v
	return s
}

type SetBizCategoryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetBizCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetBizCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s SetBizCategoryResponse) GoString() string {
	return s.String()
}

func (s *SetBizCategoryResponse) SetHeaders(v map[string]*string) *SetBizCategoryResponse {
	s.Headers = v
	return s
}

func (s *SetBizCategoryResponse) SetStatusCode(v int32) *SetBizCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *SetBizCategoryResponse) SetBody(v *SetBizCategoryResponseBody) *SetBizCategoryResponse {
	s.Body = v
	return s
}

type SetOrgConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetOrgConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetOrgConfigHeaders) GoString() string {
	return s.String()
}

func (s *SetOrgConfigHeaders) SetCommonHeaders(v map[string]*string) *SetOrgConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetOrgConfigHeaders) SetXAcsDingtalkAccessToken(v string) *SetOrgConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetOrgConfigRequest struct {
	// if can be null:
	// true
	AppDisplayStyle *int32 `json:"appDisplayStyle,omitempty" xml:"appDisplayStyle,omitempty"`
	// if can be null:
	// true
	HomepageCatalogs []*string `json:"homepageCatalogs,omitempty" xml:"homepageCatalogs,omitempty" type:"Repeated"`
	// This parameter is required.
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// This parameter is required.
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SetOrgConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetOrgConfigRequest) GoString() string {
	return s.String()
}

func (s *SetOrgConfigRequest) SetAppDisplayStyle(v int32) *SetOrgConfigRequest {
	s.AppDisplayStyle = &v
	return s
}

func (s *SetOrgConfigRequest) SetHomepageCatalogs(v []*string) *SetOrgConfigRequest {
	s.HomepageCatalogs = v
	return s
}

func (s *SetOrgConfigRequest) SetOperatorId(v string) *SetOrgConfigRequest {
	s.OperatorId = &v
	return s
}

func (s *SetOrgConfigRequest) SetStatus(v int32) *SetOrgConfigRequest {
	s.Status = &v
	return s
}

type SetOrgConfigResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SetOrgConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetOrgConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetOrgConfigResponseBody) SetSuccess(v bool) *SetOrgConfigResponseBody {
	s.Success = &v
	return s
}

type SetOrgConfigResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetOrgConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetOrgConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetOrgConfigResponse) GoString() string {
	return s.String()
}

func (s *SetOrgConfigResponse) SetHeaders(v map[string]*string) *SetOrgConfigResponse {
	s.Headers = v
	return s
}

func (s *SetOrgConfigResponse) SetStatusCode(v int32) *SetOrgConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetOrgConfigResponse) SetBody(v *SetOrgConfigResponseBody) *SetOrgConfigResponse {
	s.Body = v
	return s
}

type UpdateTodoTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateTodoTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskHeaders) SetCommonHeaders(v map[string]*string) *UpdateTodoTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTodoTaskHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateTodoTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateTodoTaskRequest struct {
	Description    *string   `json:"description,omitempty" xml:"description,omitempty"`
	Done           *bool     `json:"done,omitempty" xml:"done,omitempty"`
	DueTime        *int64    `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	ExecutorIds    []*string `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	ParticipantIds []*string `json:"participantIds,omitempty" xml:"participantIds,omitempty" type:"Repeated"`
	Subject        *string   `json:"subject,omitempty" xml:"subject,omitempty"`
	OperatorId     *string   `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s UpdateTodoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskRequest) SetDescription(v string) *UpdateTodoTaskRequest {
	s.Description = &v
	return s
}

func (s *UpdateTodoTaskRequest) SetDone(v bool) *UpdateTodoTaskRequest {
	s.Done = &v
	return s
}

func (s *UpdateTodoTaskRequest) SetDueTime(v int64) *UpdateTodoTaskRequest {
	s.DueTime = &v
	return s
}

func (s *UpdateTodoTaskRequest) SetExecutorIds(v []*string) *UpdateTodoTaskRequest {
	s.ExecutorIds = v
	return s
}

func (s *UpdateTodoTaskRequest) SetParticipantIds(v []*string) *UpdateTodoTaskRequest {
	s.ParticipantIds = v
	return s
}

func (s *UpdateTodoTaskRequest) SetSubject(v string) *UpdateTodoTaskRequest {
	s.Subject = &v
	return s
}

func (s *UpdateTodoTaskRequest) SetOperatorId(v string) *UpdateTodoTaskRequest {
	s.OperatorId = &v
	return s
}

type UpdateTodoTaskResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateTodoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskResponseBody) SetResult(v bool) *UpdateTodoTaskResponseBody {
	s.Result = &v
	return s
}

type UpdateTodoTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTodoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTodoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskResponse) SetHeaders(v map[string]*string) *UpdateTodoTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateTodoTaskResponse) SetStatusCode(v int32) *UpdateTodoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTodoTaskResponse) SetBody(v *UpdateTodoTaskResponseBody) *UpdateTodoTaskResponse {
	s.Body = v
	return s
}

type UpdateTodoTaskExecutorStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateTodoTaskExecutorStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusHeaders) SetCommonHeaders(v map[string]*string) *UpdateTodoTaskExecutorStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTodoTaskExecutorStatusHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateTodoTaskExecutorStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateTodoTaskExecutorStatusRequest struct {
	ExecutorStatusList []*UpdateTodoTaskExecutorStatusRequestExecutorStatusList `json:"executorStatusList,omitempty" xml:"executorStatusList,omitempty" type:"Repeated"`
	OperatorId         *string                                                  `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s UpdateTodoTaskExecutorStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusRequest) SetExecutorStatusList(v []*UpdateTodoTaskExecutorStatusRequestExecutorStatusList) *UpdateTodoTaskExecutorStatusRequest {
	s.ExecutorStatusList = v
	return s
}

func (s *UpdateTodoTaskExecutorStatusRequest) SetOperatorId(v string) *UpdateTodoTaskExecutorStatusRequest {
	s.OperatorId = &v
	return s
}

type UpdateTodoTaskExecutorStatusRequestExecutorStatusList struct {
	Id     *string `json:"id,omitempty" xml:"id,omitempty"`
	IsDone *bool   `json:"isDone,omitempty" xml:"isDone,omitempty"`
}

func (s UpdateTodoTaskExecutorStatusRequestExecutorStatusList) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusRequestExecutorStatusList) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusRequestExecutorStatusList) SetId(v string) *UpdateTodoTaskExecutorStatusRequestExecutorStatusList {
	s.Id = &v
	return s
}

func (s *UpdateTodoTaskExecutorStatusRequestExecutorStatusList) SetIsDone(v bool) *UpdateTodoTaskExecutorStatusRequestExecutorStatusList {
	s.IsDone = &v
	return s
}

type UpdateTodoTaskExecutorStatusResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateTodoTaskExecutorStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusResponseBody) SetResult(v bool) *UpdateTodoTaskExecutorStatusResponseBody {
	s.Result = &v
	return s
}

type UpdateTodoTaskExecutorStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTodoTaskExecutorStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTodoTaskExecutorStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusResponse) SetHeaders(v map[string]*string) *UpdateTodoTaskExecutorStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateTodoTaskExecutorStatusResponse) SetStatusCode(v int32) *UpdateTodoTaskExecutorStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTodoTaskExecutorStatusResponse) SetBody(v *UpdateTodoTaskExecutorStatusResponseBody) *UpdateTodoTaskExecutorStatusResponse {
	s.Body = v
	return s
}

type UpdateTodoTypeConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateTodoTypeConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTypeConfigHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTodoTypeConfigHeaders) SetCommonHeaders(v map[string]*string) *UpdateTodoTypeConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTodoTypeConfigHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateTodoTypeConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateTodoTypeConfigRequest struct {
	ActionList          []*UpdateTodoTypeConfigRequestActionList       `json:"actionList,omitempty" xml:"actionList,omitempty" type:"Repeated"`
	CardType            *int32                                         `json:"cardType,omitempty" xml:"cardType,omitempty"`
	ContentFieldList    []*UpdateTodoTypeConfigRequestContentFieldList `json:"contentFieldList,omitempty" xml:"contentFieldList,omitempty" type:"Repeated"`
	Description         *string                                        `json:"description,omitempty" xml:"description,omitempty"`
	Icon                *string                                        `json:"icon,omitempty" xml:"icon,omitempty"`
	PcDetailUrlOpenMode *string                                        `json:"pcDetailUrlOpenMode,omitempty" xml:"pcDetailUrlOpenMode,omitempty"`
	OperatorId          *string                                        `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s UpdateTodoTypeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTypeConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateTodoTypeConfigRequest) SetActionList(v []*UpdateTodoTypeConfigRequestActionList) *UpdateTodoTypeConfigRequest {
	s.ActionList = v
	return s
}

func (s *UpdateTodoTypeConfigRequest) SetCardType(v int32) *UpdateTodoTypeConfigRequest {
	s.CardType = &v
	return s
}

func (s *UpdateTodoTypeConfigRequest) SetContentFieldList(v []*UpdateTodoTypeConfigRequestContentFieldList) *UpdateTodoTypeConfigRequest {
	s.ContentFieldList = v
	return s
}

func (s *UpdateTodoTypeConfigRequest) SetDescription(v string) *UpdateTodoTypeConfigRequest {
	s.Description = &v
	return s
}

func (s *UpdateTodoTypeConfigRequest) SetIcon(v string) *UpdateTodoTypeConfigRequest {
	s.Icon = &v
	return s
}

func (s *UpdateTodoTypeConfigRequest) SetPcDetailUrlOpenMode(v string) *UpdateTodoTypeConfigRequest {
	s.PcDetailUrlOpenMode = &v
	return s
}

func (s *UpdateTodoTypeConfigRequest) SetOperatorId(v string) *UpdateTodoTypeConfigRequest {
	s.OperatorId = &v
	return s
}

type UpdateTodoTypeConfigRequestActionList struct {
	ActionKey       *string                `json:"actionKey,omitempty" xml:"actionKey,omitempty"`
	ActionType      *int32                 `json:"actionType,omitempty" xml:"actionType,omitempty"`
	ButtonStyleType *int32                 `json:"buttonStyleType,omitempty" xml:"buttonStyleType,omitempty"`
	NameI18n        map[string]interface{} `json:"nameI18n,omitempty" xml:"nameI18n,omitempty"`
	Url             *string                `json:"url,omitempty" xml:"url,omitempty"`
}

func (s UpdateTodoTypeConfigRequestActionList) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTypeConfigRequestActionList) GoString() string {
	return s.String()
}

func (s *UpdateTodoTypeConfigRequestActionList) SetActionKey(v string) *UpdateTodoTypeConfigRequestActionList {
	s.ActionKey = &v
	return s
}

func (s *UpdateTodoTypeConfigRequestActionList) SetActionType(v int32) *UpdateTodoTypeConfigRequestActionList {
	s.ActionType = &v
	return s
}

func (s *UpdateTodoTypeConfigRequestActionList) SetButtonStyleType(v int32) *UpdateTodoTypeConfigRequestActionList {
	s.ButtonStyleType = &v
	return s
}

func (s *UpdateTodoTypeConfigRequestActionList) SetNameI18n(v map[string]interface{}) *UpdateTodoTypeConfigRequestActionList {
	s.NameI18n = v
	return s
}

func (s *UpdateTodoTypeConfigRequestActionList) SetUrl(v string) *UpdateTodoTypeConfigRequestActionList {
	s.Url = &v
	return s
}

type UpdateTodoTypeConfigRequestContentFieldList struct {
	FieldKey  *string                `json:"fieldKey,omitempty" xml:"fieldKey,omitempty"`
	FieldType *string                `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	NameI18n  map[string]interface{} `json:"nameI18n,omitempty" xml:"nameI18n,omitempty"`
}

func (s UpdateTodoTypeConfigRequestContentFieldList) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTypeConfigRequestContentFieldList) GoString() string {
	return s.String()
}

func (s *UpdateTodoTypeConfigRequestContentFieldList) SetFieldKey(v string) *UpdateTodoTypeConfigRequestContentFieldList {
	s.FieldKey = &v
	return s
}

func (s *UpdateTodoTypeConfigRequestContentFieldList) SetFieldType(v string) *UpdateTodoTypeConfigRequestContentFieldList {
	s.FieldType = &v
	return s
}

func (s *UpdateTodoTypeConfigRequestContentFieldList) SetNameI18n(v map[string]interface{}) *UpdateTodoTypeConfigRequestContentFieldList {
	s.NameI18n = v
	return s
}

type UpdateTodoTypeConfigResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateTodoTypeConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTypeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTodoTypeConfigResponseBody) SetResult(v bool) *UpdateTodoTypeConfigResponseBody {
	s.Result = &v
	return s
}

type UpdateTodoTypeConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTodoTypeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTodoTypeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTypeConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateTodoTypeConfigResponse) SetHeaders(v map[string]*string) *UpdateTodoTypeConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateTodoTypeConfigResponse) SetStatusCode(v int32) *UpdateTodoTypeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTodoTypeConfigResponse) SetBody(v *UpdateTodoTypeConfigResponseBody) *UpdateTodoTypeConfigResponse {
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
// 查询用户待办计数
//
// @param request - CountTodoTasksRequest
//
// @param headers - CountTodoTasksHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CountTodoTasksResponse
func (client *Client) CountTodoTasksWithOptions(unionId *string, request *CountTodoTasksRequest, headers *CountTodoTasksHeaders, runtime *util.RuntimeOptions) (_result *CountTodoTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FromDueTime)) {
		body["fromDueTime"] = request.FromDueTime
	}

	if !tea.BoolValue(util.IsUnset(request.IsDone)) {
		body["isDone"] = request.IsDone
	}

	if !tea.BoolValue(util.IsUnset(request.IsRecycled)) {
		body["isRecycled"] = request.IsRecycled
	}

	if !tea.BoolValue(util.IsUnset(request.RoleTypes)) {
		body["roleTypes"] = request.RoleTypes
	}

	if !tea.BoolValue(util.IsUnset(request.ToDueTime)) {
		body["toDueTime"] = request.ToDueTime
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
		Action:      tea.String("CountTodoTasks"),
		Version:     tea.String("todo_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todo/users/" + tea.StringValue(unionId) + "/tasks/count"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CountTodoTasksResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户待办计数
//
// @param request - CountTodoTasksRequest
//
// @return CountTodoTasksResponse
func (client *Client) CountTodoTasks(unionId *string, request *CountTodoTasksRequest) (_result *CountTodoTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CountTodoTasksHeaders{}
	_result = &CountTodoTasksResponse{}
	_body, _err := client.CountTodoTasksWithOptions(unionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 以用户个人身份创建个人待办
//
// @param request - CreatePersonalTodoTaskRequest
//
// @param headers - CreatePersonalTodoTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePersonalTodoTaskResponse
func (client *Client) CreatePersonalTodoTaskWithOptions(request *CreatePersonalTodoTaskRequest, headers *CreatePersonalTodoTaskHeaders, runtime *util.RuntimeOptions) (_result *CreatePersonalTodoTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
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

	if !tea.BoolValue(util.IsUnset(request.ParticipantIds)) {
		body["participantIds"] = request.ParticipantIds
	}

	if !tea.BoolValue(util.IsUnset(request.ReminderTimeStamp)) {
		body["reminderTimeStamp"] = request.ReminderTimeStamp
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		body["subject"] = request.Subject
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
		Action:      tea.String("CreatePersonalTodoTask"),
		Version:     tea.String("todo_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todo/users/me/personalTasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePersonalTodoTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 以用户个人身份创建个人待办
//
// @param request - CreatePersonalTodoTaskRequest
//
// @return CreatePersonalTodoTaskResponse
func (client *Client) CreatePersonalTodoTask(request *CreatePersonalTodoTaskRequest) (_result *CreatePersonalTodoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreatePersonalTodoTaskHeaders{}
	_result = &CreatePersonalTodoTaskResponse{}
	_body, _err := client.CreatePersonalTodoTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建待办
//
// @param request - CreateTodoTaskRequest
//
// @param headers - CreateTodoTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTodoTaskResponse
func (client *Client) CreateTodoTaskWithOptions(unionId *string, request *CreateTodoTaskRequest, headers *CreateTodoTaskHeaders, runtime *util.RuntimeOptions) (_result *CreateTodoTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionList)) {
		body["actionList"] = request.ActionList
	}

	if !tea.BoolValue(util.IsUnset(request.BizCategoryId)) {
		body["bizCategoryId"] = request.BizCategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.ContentFieldList)) {
		body["contentFieldList"] = request.ContentFieldList
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorId)) {
		body["creatorId"] = request.CreatorId
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

	if !tea.BoolValue(util.IsUnset(request.IsOnlyShowExecutor)) {
		body["isOnlyShowExecutor"] = request.IsOnlyShowExecutor
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyConfigs)) {
		body["notifyConfigs"] = request.NotifyConfigs
	}

	if !tea.BoolValue(util.IsUnset(request.ParticipantIds)) {
		body["participantIds"] = request.ParticipantIds
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.RemindNotifyConfigs)) {
		body["remindNotifyConfigs"] = request.RemindNotifyConfigs
	}

	if !tea.BoolValue(util.IsUnset(request.ReminderTimeStamp)) {
		body["reminderTimeStamp"] = request.ReminderTimeStamp
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		body["sourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		body["subject"] = request.Subject
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdExtension)) {
		body["thirdExtension"] = request.ThirdExtension
	}

	if !tea.BoolValue(util.IsUnset(request.TodoType)) {
		body["todoType"] = request.TodoType
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
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTodoTask"),
		Version:     tea.String("todo_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todo/users/" + tea.StringValue(unionId) + "/tasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTodoTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建待办
//
// @param request - CreateTodoTaskRequest
//
// @return CreateTodoTaskResponse
func (client *Client) CreateTodoTask(unionId *string, request *CreateTodoTaskRequest) (_result *CreateTodoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTodoTaskHeaders{}
	_result = &CreateTodoTaskResponse{}
	_body, _err := client.CreateTodoTaskWithOptions(unionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建待办卡片类型配置
//
// @param request - CreateTodoTypeConfigRequest
//
// @param headers - CreateTodoTypeConfigHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTodoTypeConfigResponse
func (client *Client) CreateTodoTypeConfigWithOptions(unionId *string, request *CreateTodoTypeConfigRequest, headers *CreateTodoTypeConfigHeaders, runtime *util.RuntimeOptions) (_result *CreateTodoTypeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionList)) {
		body["actionList"] = request.ActionList
	}

	if !tea.BoolValue(util.IsUnset(request.CardType)) {
		body["cardType"] = request.CardType
	}

	if !tea.BoolValue(util.IsUnset(request.ContentFieldList)) {
		body["contentFieldList"] = request.ContentFieldList
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.PcDetailUrlOpenMode)) {
		body["pcDetailUrlOpenMode"] = request.PcDetailUrlOpenMode
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
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTodoTypeConfig"),
		Version:     tea.String("todo_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todo/users/" + tea.StringValue(unionId) + "/configs/types"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTodoTypeConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建待办卡片类型配置
//
// @param request - CreateTodoTypeConfigRequest
//
// @return CreateTodoTypeConfigResponse
func (client *Client) CreateTodoTypeConfig(unionId *string, request *CreateTodoTypeConfigRequest) (_result *CreateTodoTypeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTodoTypeConfigHeaders{}
	_result = &CreateTodoTypeConfigResponse{}
	_body, _err := client.CreateTodoTypeConfigWithOptions(unionId, request, headers, runtime)
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
// @param request - DeleteTodoTaskRequest
//
// @param headers - DeleteTodoTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTodoTaskResponse
func (client *Client) DeleteTodoTaskWithOptions(unionId *string, taskId *string, request *DeleteTodoTaskRequest, headers *DeleteTodoTaskHeaders, runtime *util.RuntimeOptions) (_result *DeleteTodoTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
		Action:      tea.String("DeleteTodoTask"),
		Version:     tea.String("todo_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todo/users/" + tea.StringValue(unionId) + "/tasks/" + tea.StringValue(taskId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTodoTaskResponse{}
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
// @param request - DeleteTodoTaskRequest
//
// @return DeleteTodoTaskResponse
func (client *Client) DeleteTodoTask(unionId *string, taskId *string, request *DeleteTodoTaskRequest) (_result *DeleteTodoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteTodoTaskHeaders{}
	_result = &DeleteTodoTaskResponse{}
	_body, _err := client.DeleteTodoTaskWithOptions(unionId, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询待办
//
// @param headers - GetTodoTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTodoTaskResponse
func (client *Client) GetTodoTaskWithOptions(unionId *string, taskId *string, headers *GetTodoTaskHeaders, runtime *util.RuntimeOptions) (_result *GetTodoTaskResponse, _err error) {
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
		Action:      tea.String("GetTodoTask"),
		Version:     tea.String("todo_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todo/users/" + tea.StringValue(unionId) + "/tasks/" + tea.StringValue(taskId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTodoTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询待办
//
// @return GetTodoTaskResponse
func (client *Client) GetTodoTask(unionId *string, taskId *string) (_result *GetTodoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTodoTaskHeaders{}
	_result = &GetTodoTaskResponse{}
	_body, _err := client.GetTodoTaskWithOptions(unionId, taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据sourceId查询待办详情
//
// @param headers - GetTodoTaskBySourceIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTodoTaskBySourceIdResponse
func (client *Client) GetTodoTaskBySourceIdWithOptions(unionId *string, sourceId *string, headers *GetTodoTaskBySourceIdHeaders, runtime *util.RuntimeOptions) (_result *GetTodoTaskBySourceIdResponse, _err error) {
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
		Action:      tea.String("GetTodoTaskBySourceId"),
		Version:     tea.String("todo_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todo/users/" + tea.StringValue(unionId) + "/tasks/sources/" + tea.StringValue(sourceId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTodoTaskBySourceIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据sourceId查询待办详情
//
// @return GetTodoTaskBySourceIdResponse
func (client *Client) GetTodoTaskBySourceId(unionId *string, sourceId *string) (_result *GetTodoTaskBySourceIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTodoTaskBySourceIdHeaders{}
	_result = &GetTodoTaskBySourceIdResponse{}
	_body, _err := client.GetTodoTaskBySourceIdWithOptions(unionId, sourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 专属钉根据待办ID查询待办详情
//
// @param headers - GetTodoTaskDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTodoTaskDetailResponse
func (client *Client) GetTodoTaskDetailWithOptions(taskId *string, unionId *string, headers *GetTodoTaskDetailHeaders, runtime *util.RuntimeOptions) (_result *GetTodoTaskDetailResponse, _err error) {
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
		Action:      tea.String("GetTodoTaskDetail"),
		Version:     tea.String("todo_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todo/exclusive/users/" + tea.StringValue(unionId) + "/tasks/" + tea.StringValue(taskId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTodoTaskDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 专属钉根据待办ID查询待办详情
//
// @return GetTodoTaskDetailResponse
func (client *Client) GetTodoTaskDetail(taskId *string, unionId *string) (_result *GetTodoTaskDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTodoTaskDetailHeaders{}
	_result = &GetTodoTaskDetailResponse{}
	_body, _err := client.GetTodoTaskDetailWithOptions(taskId, unionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据id获取待办卡片类型配置
//
// @param headers - GetTodoTypeConfigHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTodoTypeConfigResponse
func (client *Client) GetTodoTypeConfigWithOptions(unionId *string, cardTypeId *string, headers *GetTodoTypeConfigHeaders, runtime *util.RuntimeOptions) (_result *GetTodoTypeConfigResponse, _err error) {
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
		Action:      tea.String("GetTodoTypeConfig"),
		Version:     tea.String("todo_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todo/users/" + tea.StringValue(unionId) + "/configs/types/" + tea.StringValue(cardTypeId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTodoTypeConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据id获取待办卡片类型配置
//
// @return GetTodoTypeConfigResponse
func (client *Client) GetTodoTypeConfig(unionId *string, cardTypeId *string) (_result *GetTodoTypeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTodoTypeConfigHeaders{}
	_result = &GetTodoTypeConfigResponse{}
	_body, _err := client.GetTodoTypeConfigWithOptions(unionId, cardTypeId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询待办应用下所有二级分类
//
// @param headers - ListAllBizCategoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllBizCategoryResponse
func (client *Client) ListAllBizCategoryWithOptions(headers *ListAllBizCategoryHeaders, runtime *util.RuntimeOptions) (_result *ListAllBizCategoryResponse, _err error) {
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
		Action:      tea.String("ListAllBizCategory"),
		Version:     tea.String("todo_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todo/apps/allBizcategories/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAllBizCategoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询待办应用下所有二级分类
//
// @return ListAllBizCategoryResponse
func (client *Client) ListAllBizCategory() (_result *ListAllBizCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListAllBizCategoryHeaders{}
	_result = &ListAllBizCategoryResponse{}
	_body, _err := client.ListAllBizCategoryWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询企业配置
//
// @param headers - QueryOrgConfigHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOrgConfigResponse
func (client *Client) QueryOrgConfigWithOptions(headers *QueryOrgConfigHeaders, runtime *util.RuntimeOptions) (_result *QueryOrgConfigResponse, _err error) {
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
		Action:      tea.String("QueryOrgConfig"),
		Version:     tea.String("todo_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todo/organizations/configs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOrgConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询企业配置
//
// @return QueryOrgConfigResponse
func (client *Client) QueryOrgConfig() (_result *QueryOrgConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOrgConfigHeaders{}
	_result = &QueryOrgConfigResponse{}
	_body, _err := client.QueryOrgConfigWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户企业类型待办列表，支持查询当前企业的一方应用、三方应用、自建应用产生的工作待办数据
//
// @param request - QueryOrgTodoByUserRequest
//
// @param headers - QueryOrgTodoByUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOrgTodoByUserResponse
func (client *Client) QueryOrgTodoByUserWithOptions(unionId *string, request *QueryOrgTodoByUserRequest, headers *QueryOrgTodoByUserHeaders, runtime *util.RuntimeOptions) (_result *QueryOrgTodoByUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FromDueTime)) {
		body["fromDueTime"] = request.FromDueTime
	}

	if !tea.BoolValue(util.IsUnset(request.IsDone)) {
		body["isDone"] = request.IsDone
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NeedPersonalTodo)) {
		body["needPersonalTodo"] = request.NeedPersonalTodo
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		body["orderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.OrderDirection)) {
		body["orderDirection"] = request.OrderDirection
	}

	if !tea.BoolValue(util.IsUnset(request.RoleTypes)) {
		body["roleTypes"] = request.RoleTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		body["subject"] = request.Subject
	}

	if !tea.BoolValue(util.IsUnset(request.ToDueTime)) {
		body["toDueTime"] = request.ToDueTime
	}

	if !tea.BoolValue(util.IsUnset(request.TodoType)) {
		body["todoType"] = request.TodoType
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
		Action:      tea.String("QueryOrgTodoByUser"),
		Version:     tea.String("todo_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todo/users/" + tea.StringValue(unionId) + "/organizations/tasks/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOrgTodoByUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户企业类型待办列表，支持查询当前企业的一方应用、三方应用、自建应用产生的工作待办数据
//
// @param request - QueryOrgTodoByUserRequest
//
// @return QueryOrgTodoByUserResponse
func (client *Client) QueryOrgTodoByUser(unionId *string, request *QueryOrgTodoByUserRequest) (_result *QueryOrgTodoByUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOrgTodoByUserHeaders{}
	_result = &QueryOrgTodoByUserResponse{}
	_body, _err := client.QueryOrgTodoByUserWithOptions(unionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询企业下用户待办列表
//
// @param request - QueryOrgTodoTasksRequest
//
// @param headers - QueryOrgTodoTasksHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOrgTodoTasksResponse
func (client *Client) QueryOrgTodoTasksWithOptions(unionId *string, request *QueryOrgTodoTasksRequest, headers *QueryOrgTodoTasksHeaders, runtime *util.RuntimeOptions) (_result *QueryOrgTodoTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsDone)) {
		body["isDone"] = request.IsDone
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RoleTypes)) {
		body["roleTypes"] = request.RoleTypes
	}

	if !tea.BoolValue(util.IsUnset(request.TodoType)) {
		body["todoType"] = request.TodoType
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
		Action:      tea.String("QueryOrgTodoTasks"),
		Version:     tea.String("todo_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todo/users/" + tea.StringValue(unionId) + "/org/tasks/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOrgTodoTasksResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询企业下用户待办列表
//
// @param request - QueryOrgTodoTasksRequest
//
// @return QueryOrgTodoTasksResponse
func (client *Client) QueryOrgTodoTasks(unionId *string, request *QueryOrgTodoTasksRequest) (_result *QueryOrgTodoTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOrgTodoTasksHeaders{}
	_result = &QueryOrgTodoTasksResponse{}
	_body, _err := client.QueryOrgTodoTasksWithOptions(unionId, request, headers, runtime)
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
// @param request - QueryTodoTasksRequest
//
// @param headers - QueryTodoTasksHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTodoTasksResponse
func (client *Client) QueryTodoTasksWithOptions(unionId *string, request *QueryTodoTasksRequest, headers *QueryTodoTasksHeaders, runtime *util.RuntimeOptions) (_result *QueryTodoTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.FromDueTime)) {
		body["fromDueTime"] = request.FromDueTime
	}

	if !tea.BoolValue(util.IsUnset(request.IsDone)) {
		body["isDone"] = request.IsDone
	}

	if !tea.BoolValue(util.IsUnset(request.IsRecycled)) {
		body["isRecycled"] = request.IsRecycled
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		body["orderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.OrderDirection)) {
		body["orderDirection"] = request.OrderDirection
	}

	if !tea.BoolValue(util.IsUnset(request.RoleTypes)) {
		body["roleTypes"] = request.RoleTypes
	}

	if !tea.BoolValue(util.IsUnset(request.ToDueTime)) {
		body["toDueTime"] = request.ToDueTime
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
		Action:      tea.String("QueryTodoTasks"),
		Version:     tea.String("todo_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todo/users/" + tea.StringValue(unionId) + "/tasks/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTodoTasksResponse{}
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
// @param request - QueryTodoTasksRequest
//
// @return QueryTodoTasksResponse
func (client *Client) QueryTodoTasks(unionId *string, request *QueryTodoTasksRequest) (_result *QueryTodoTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryTodoTasksHeaders{}
	_result = &QueryTodoTasksResponse{}
	_body, _err := client.QueryTodoTasksWithOptions(unionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除二级分类
//
// @param request - RemoveBizCategoryRequest
//
// @param headers - RemoveBizCategoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveBizCategoryResponse
func (client *Client) RemoveBizCategoryWithOptions(request *RemoveBizCategoryRequest, headers *RemoveBizCategoryHeaders, runtime *util.RuntimeOptions) (_result *RemoveBizCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCategoryId)) {
		query["bizCategoryId"] = request.BizCategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
		Action:      tea.String("RemoveBizCategory"),
		Version:     tea.String("todo_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todo/apps/bizcategories"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveBizCategoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除二级分类
//
// @param request - RemoveBizCategoryRequest
//
// @return RemoveBizCategoryResponse
func (client *Client) RemoveBizCategory(request *RemoveBizCategoryRequest) (_result *RemoveBizCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveBizCategoryHeaders{}
	_result = &RemoveBizCategoryResponse{}
	_body, _err := client.RemoveBizCategoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 保存二级分类
//
// @param request - SetBizCategoryRequest
//
// @param headers - SetBizCategoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetBizCategoryResponse
func (client *Client) SetBizCategoryWithOptions(request *SetBizCategoryRequest, headers *SetBizCategoryHeaders, runtime *util.RuntimeOptions) (_result *SetBizCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCategoryId)) {
		query["bizCategoryId"] = request.BizCategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
		Action:      tea.String("SetBizCategory"),
		Version:     tea.String("todo_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todo/apps/bizcategories/save"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SetBizCategoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存二级分类
//
// @param request - SetBizCategoryRequest
//
// @return SetBizCategoryResponse
func (client *Client) SetBizCategory(request *SetBizCategoryRequest) (_result *SetBizCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetBizCategoryHeaders{}
	_result = &SetBizCategoryResponse{}
	_body, _err := client.SetBizCategoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 保存企业配置
//
// @param request - SetOrgConfigRequest
//
// @param headers - SetOrgConfigHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetOrgConfigResponse
func (client *Client) SetOrgConfigWithOptions(request *SetOrgConfigRequest, headers *SetOrgConfigHeaders, runtime *util.RuntimeOptions) (_result *SetOrgConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppDisplayStyle)) {
		body["appDisplayStyle"] = request.AppDisplayStyle
	}

	if !tea.BoolValue(util.IsUnset(request.HomepageCatalogs)) {
		body["homepageCatalogs"] = request.HomepageCatalogs
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
		Action:      tea.String("SetOrgConfig"),
		Version:     tea.String("todo_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todo/organizations/configs/save"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SetOrgConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存企业配置
//
// @param request - SetOrgConfigRequest
//
// @return SetOrgConfigResponse
func (client *Client) SetOrgConfig(request *SetOrgConfigRequest) (_result *SetOrgConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetOrgConfigHeaders{}
	_result = &SetOrgConfigResponse{}
	_body, _err := client.SetOrgConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新待办
//
// @param request - UpdateTodoTaskRequest
//
// @param headers - UpdateTodoTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTodoTaskResponse
func (client *Client) UpdateTodoTaskWithOptions(unionId *string, taskId *string, request *UpdateTodoTaskRequest, headers *UpdateTodoTaskHeaders, runtime *util.RuntimeOptions) (_result *UpdateTodoTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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

	if !tea.BoolValue(util.IsUnset(request.ParticipantIds)) {
		body["participantIds"] = request.ParticipantIds
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		body["subject"] = request.Subject
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
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTodoTask"),
		Version:     tea.String("todo_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todo/users/" + tea.StringValue(unionId) + "/tasks/" + tea.StringValue(taskId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTodoTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新待办
//
// @param request - UpdateTodoTaskRequest
//
// @return UpdateTodoTaskResponse
func (client *Client) UpdateTodoTask(unionId *string, taskId *string, request *UpdateTodoTaskRequest) (_result *UpdateTodoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTodoTaskHeaders{}
	_result = &UpdateTodoTaskResponse{}
	_body, _err := client.UpdateTodoTaskWithOptions(unionId, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新待办执行者状态
//
// @param request - UpdateTodoTaskExecutorStatusRequest
//
// @param headers - UpdateTodoTaskExecutorStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTodoTaskExecutorStatusResponse
func (client *Client) UpdateTodoTaskExecutorStatusWithOptions(unionId *string, taskId *string, request *UpdateTodoTaskExecutorStatusRequest, headers *UpdateTodoTaskExecutorStatusHeaders, runtime *util.RuntimeOptions) (_result *UpdateTodoTaskExecutorStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExecutorStatusList)) {
		body["executorStatusList"] = request.ExecutorStatusList
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
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTodoTaskExecutorStatus"),
		Version:     tea.String("todo_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todo/users/" + tea.StringValue(unionId) + "/tasks/" + tea.StringValue(taskId) + "/executorStatus"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTodoTaskExecutorStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新待办执行者状态
//
// @param request - UpdateTodoTaskExecutorStatusRequest
//
// @return UpdateTodoTaskExecutorStatusResponse
func (client *Client) UpdateTodoTaskExecutorStatus(unionId *string, taskId *string, request *UpdateTodoTaskExecutorStatusRequest) (_result *UpdateTodoTaskExecutorStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTodoTaskExecutorStatusHeaders{}
	_result = &UpdateTodoTaskExecutorStatusResponse{}
	_body, _err := client.UpdateTodoTaskExecutorStatusWithOptions(unionId, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新待办卡片类型配置
//
// @param request - UpdateTodoTypeConfigRequest
//
// @param headers - UpdateTodoTypeConfigHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTodoTypeConfigResponse
func (client *Client) UpdateTodoTypeConfigWithOptions(unionId *string, cardTypeId *string, request *UpdateTodoTypeConfigRequest, headers *UpdateTodoTypeConfigHeaders, runtime *util.RuntimeOptions) (_result *UpdateTodoTypeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionList)) {
		body["actionList"] = request.ActionList
	}

	if !tea.BoolValue(util.IsUnset(request.CardType)) {
		body["cardType"] = request.CardType
	}

	if !tea.BoolValue(util.IsUnset(request.ContentFieldList)) {
		body["contentFieldList"] = request.ContentFieldList
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.PcDetailUrlOpenMode)) {
		body["pcDetailUrlOpenMode"] = request.PcDetailUrlOpenMode
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
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTodoTypeConfig"),
		Version:     tea.String("todo_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/todo/users/" + tea.StringValue(unionId) + "/configs/types/" + tea.StringValue(cardTypeId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTodoTypeConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新待办卡片类型配置
//
// @param request - UpdateTodoTypeConfigRequest
//
// @return UpdateTodoTypeConfigResponse
func (client *Client) UpdateTodoTypeConfig(unionId *string, cardTypeId *string, request *UpdateTodoTypeConfigRequest) (_result *UpdateTodoTypeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTodoTypeConfigHeaders{}
	_result = &UpdateTodoTypeConfigResponse{}
	_body, _err := client.UpdateTodoTypeConfigWithOptions(unionId, cardTypeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
