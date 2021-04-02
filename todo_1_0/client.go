// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package todo_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	// id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 标题
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 截止时间
	DueTime *int64 `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	// 完成时间
	FinishTime *int64 `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// 完成状态
	Done *bool `json:"done,omitempty" xml:"done,omitempty"`
	// 执行者列表
	ExecutorIds []*string `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	// 参与者列表
	ParticipantIds []*string `json:"participantIds,omitempty" xml:"participantIds,omitempty" type:"Repeated"`
	// 提醒规则
	Reminder *GetTodoTaskResponseBodyReminder `json:"reminder,omitempty" xml:"reminder,omitempty" type:"Struct"`
	// 重复规则
	Recurrence *string `json:"recurrence,omitempty" xml:"recurrence,omitempty"`
	// 自定义详情页跳转配置
	DetailUrl *GetTodoTaskResponseBodyDetailUrl `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	// 业务来源id
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 业务来源
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 创建时间
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 更新时间
	ModifiedTime *int64 `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// 创建者id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 更新者id
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// 租户id
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 接入业务应用标识
	BizTag *string `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetTodoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTodoTaskResponseBody) SetId(v string) *GetTodoTaskResponseBody {
	s.Id = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetSubject(v string) *GetTodoTaskResponseBody {
	s.Subject = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetDescription(v string) *GetTodoTaskResponseBody {
	s.Description = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetStartTime(v int64) *GetTodoTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetDueTime(v int64) *GetTodoTaskResponseBody {
	s.DueTime = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetFinishTime(v int64) *GetTodoTaskResponseBody {
	s.FinishTime = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetDone(v bool) *GetTodoTaskResponseBody {
	s.Done = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetExecutorIds(v []*string) *GetTodoTaskResponseBody {
	s.ExecutorIds = v
	return s
}

func (s *GetTodoTaskResponseBody) SetParticipantIds(v []*string) *GetTodoTaskResponseBody {
	s.ParticipantIds = v
	return s
}

func (s *GetTodoTaskResponseBody) SetReminder(v *GetTodoTaskResponseBodyReminder) *GetTodoTaskResponseBody {
	s.Reminder = v
	return s
}

func (s *GetTodoTaskResponseBody) SetRecurrence(v string) *GetTodoTaskResponseBody {
	s.Recurrence = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetDetailUrl(v *GetTodoTaskResponseBodyDetailUrl) *GetTodoTaskResponseBody {
	s.DetailUrl = v
	return s
}

func (s *GetTodoTaskResponseBody) SetSourceId(v string) *GetTodoTaskResponseBody {
	s.SourceId = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetSource(v string) *GetTodoTaskResponseBody {
	s.Source = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetCreatedTime(v int64) *GetTodoTaskResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetModifiedTime(v int64) *GetTodoTaskResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetCreatorId(v string) *GetTodoTaskResponseBody {
	s.CreatorId = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetModifierId(v string) *GetTodoTaskResponseBody {
	s.ModifierId = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetTenantId(v string) *GetTodoTaskResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetBizTag(v string) *GetTodoTaskResponseBody {
	s.BizTag = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetRequestId(v string) *GetTodoTaskResponseBody {
	s.RequestId = &v
	return s
}

type GetTodoTaskResponseBodyReminder struct {
	// 提醒通道
	Channel *int32 `json:"channel,omitempty" xml:"channel,omitempty"`
	// 提醒规则配置
	Rules []*GetTodoTaskResponseBodyReminderRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
}

func (s GetTodoTaskResponseBodyReminder) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTaskResponseBodyReminder) GoString() string {
	return s.String()
}

func (s *GetTodoTaskResponseBodyReminder) SetChannel(v int32) *GetTodoTaskResponseBodyReminder {
	s.Channel = &v
	return s
}

func (s *GetTodoTaskResponseBodyReminder) SetRules(v []*GetTodoTaskResponseBodyReminderRules) *GetTodoTaskResponseBodyReminder {
	s.Rules = v
	return s
}

type GetTodoTaskResponseBodyReminderRules struct {
	// startTime:相对开始时间 // dueTime:相对截止时间 // customTime: 绝对时间
	BaseTime *string `json:"baseTime,omitempty" xml:"baseTime,omitempty"`
	// baseTime 为 startTime 或者 dueTime 时表分钟；比如截止前15分钟为 -15，截止前3小时为 -180;basetime 为 customTime 时为时间戳
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
}

func (s GetTodoTaskResponseBodyReminderRules) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTaskResponseBodyReminderRules) GoString() string {
	return s.String()
}

func (s *GetTodoTaskResponseBodyReminderRules) SetBaseTime(v string) *GetTodoTaskResponseBodyReminderRules {
	s.BaseTime = &v
	return s
}

func (s *GetTodoTaskResponseBodyReminderRules) SetOffset(v int32) *GetTodoTaskResponseBodyReminderRules {
	s.Offset = &v
	return s
}

type GetTodoTaskResponseBodyDetailUrl struct {
	// pc端详情页地址
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	// app端详情页地址
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
}

func (s GetTodoTaskResponseBodyDetailUrl) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTaskResponseBodyDetailUrl) GoString() string {
	return s.String()
}

func (s *GetTodoTaskResponseBodyDetailUrl) SetPcUrl(v string) *GetTodoTaskResponseBodyDetailUrl {
	s.PcUrl = &v
	return s
}

func (s *GetTodoTaskResponseBodyDetailUrl) SetAppUrl(v string) *GetTodoTaskResponseBodyDetailUrl {
	s.AppUrl = &v
	return s
}

type GetTodoTaskResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTodoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetTodoTaskResponse) SetBody(v *GetTodoTaskResponseBody) *GetTodoTaskResponse {
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
	// 操作者id
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
	// 删除结果
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteTodoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTodoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTodoTaskResponseBody) SetResult(v bool) *DeleteTodoTaskResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteTodoTaskResponseBody) SetRequestId(v string) *DeleteTodoTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTodoTaskResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTodoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteTodoTaskResponse) SetBody(v *DeleteTodoTaskResponseBody) *DeleteTodoTaskResponse {
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
	// 待办标题
	Sucject *string `json:"sucject,omitempty" xml:"sucject,omitempty"`
	// 待办描述备注
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 截止时间
	DueTime *int64 `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	// 完成状态
	Done *bool `json:"done,omitempty" xml:"done,omitempty"`
	// 执行者列表
	ExecutorIds []*string `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	// 参与者列表
	ParticipantIds []*string `json:"participantIds,omitempty" xml:"participantIds,omitempty" type:"Repeated"`
	// 详情页url跳转地址
	DetailUrl *UpdateTodoTaskRequestDetailUrl `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	// 待办重复规则配置
	Recurrence *string `json:"recurrence,omitempty" xml:"recurrence,omitempty"`
	// 待办提醒规则配置
	Reminder *UpdateTodoTaskRequestReminder `json:"reminder,omitempty" xml:"reminder,omitempty" type:"Struct"`
	// 当前操作者id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s UpdateTodoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskRequest) SetSucject(v string) *UpdateTodoTaskRequest {
	s.Sucject = &v
	return s
}

func (s *UpdateTodoTaskRequest) SetDescription(v string) *UpdateTodoTaskRequest {
	s.Description = &v
	return s
}

func (s *UpdateTodoTaskRequest) SetDueTime(v int64) *UpdateTodoTaskRequest {
	s.DueTime = &v
	return s
}

func (s *UpdateTodoTaskRequest) SetDone(v bool) *UpdateTodoTaskRequest {
	s.Done = &v
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

func (s *UpdateTodoTaskRequest) SetDetailUrl(v *UpdateTodoTaskRequestDetailUrl) *UpdateTodoTaskRequest {
	s.DetailUrl = v
	return s
}

func (s *UpdateTodoTaskRequest) SetRecurrence(v string) *UpdateTodoTaskRequest {
	s.Recurrence = &v
	return s
}

func (s *UpdateTodoTaskRequest) SetReminder(v *UpdateTodoTaskRequestReminder) *UpdateTodoTaskRequest {
	s.Reminder = v
	return s
}

func (s *UpdateTodoTaskRequest) SetOperatorId(v string) *UpdateTodoTaskRequest {
	s.OperatorId = &v
	return s
}

type UpdateTodoTaskRequestDetailUrl struct {
	// app端详情页url
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
	// pc端详情页url
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
}

func (s UpdateTodoTaskRequestDetailUrl) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskRequestDetailUrl) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskRequestDetailUrl) SetAppUrl(v string) *UpdateTodoTaskRequestDetailUrl {
	s.AppUrl = &v
	return s
}

func (s *UpdateTodoTaskRequestDetailUrl) SetPcUrl(v string) *UpdateTodoTaskRequestDetailUrl {
	s.PcUrl = &v
	return s
}

type UpdateTodoTaskRequestReminder struct {
	// 提醒通道，1 弹框，2 短信，3 电话
	Channel *int32 `json:"channel,omitempty" xml:"channel,omitempty"`
	// 提醒规则列表
	Rules []*UpdateTodoTaskRequestReminderRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
}

func (s UpdateTodoTaskRequestReminder) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskRequestReminder) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskRequestReminder) SetChannel(v int32) *UpdateTodoTaskRequestReminder {
	s.Channel = &v
	return s
}

func (s *UpdateTodoTaskRequestReminder) SetRules(v []*UpdateTodoTaskRequestReminderRules) *UpdateTodoTaskRequestReminder {
	s.Rules = v
	return s
}

type UpdateTodoTaskRequestReminderRules struct {
	// startTime:相对开始时间  											//  					dueTime:相对截止时间 											//						customTime: 绝对时间
	BaseTime *string `json:"baseTime,omitempty" xml:"baseTime,omitempty"`
	// 必须，baseTime 为 startTime 或者 dueTime 时表分钟；比如截止前15分钟为 -15，截止前3小时为 -180;basetime 为 customTime 时为时间戳
	Offset *int64 `json:"offset,omitempty" xml:"offset,omitempty"`
}

func (s UpdateTodoTaskRequestReminderRules) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskRequestReminderRules) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskRequestReminderRules) SetBaseTime(v string) *UpdateTodoTaskRequestReminderRules {
	s.BaseTime = &v
	return s
}

func (s *UpdateTodoTaskRequestReminderRules) SetOffset(v int64) *UpdateTodoTaskRequestReminderRules {
	s.Offset = &v
	return s
}

type UpdateTodoTaskResponseBody struct {
	// 更新结果
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTodoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateTodoTaskResponse) SetBody(v *UpdateTodoTaskResponseBody) *UpdateTodoTaskResponse {
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
	// 来源id，接入业务系统侧的唯一标识id
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 待办标题
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// 创建者id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 待办备注描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 截止时间
	DueTime *int64 `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	// 执行者列表
	ExecutorIds []*string `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	// 参与者列表
	ParticipantIds []*string `json:"participantIds,omitempty" xml:"participantIds,omitempty" type:"Repeated"`
	// 详情页url跳转地址
	DetailUrl *CreateTodoTaskRequestDetailUrl `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	// 待办重复规则配置
	Recurrence *string `json:"recurrence,omitempty" xml:"recurrence,omitempty"`
	// 待办提醒规则配置
	Reminder *CreateTodoTaskRequestReminder `json:"reminder,omitempty" xml:"reminder,omitempty" type:"Struct"`
	// 待办通知配置（包含单聊卡片、ding通知、群聊卡片、同步日历、同步系统消息等通知能力）
	NotifyConfigs *CreateTodoTaskRequestNotifyConfigs `json:"notifyConfigs,omitempty" xml:"notifyConfigs,omitempty" type:"Struct"`
	// 当前操作者id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s CreateTodoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskRequest) SetSourceId(v string) *CreateTodoTaskRequest {
	s.SourceId = &v
	return s
}

func (s *CreateTodoTaskRequest) SetSubject(v string) *CreateTodoTaskRequest {
	s.Subject = &v
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

func (s *CreateTodoTaskRequest) SetDueTime(v int64) *CreateTodoTaskRequest {
	s.DueTime = &v
	return s
}

func (s *CreateTodoTaskRequest) SetExecutorIds(v []*string) *CreateTodoTaskRequest {
	s.ExecutorIds = v
	return s
}

func (s *CreateTodoTaskRequest) SetParticipantIds(v []*string) *CreateTodoTaskRequest {
	s.ParticipantIds = v
	return s
}

func (s *CreateTodoTaskRequest) SetDetailUrl(v *CreateTodoTaskRequestDetailUrl) *CreateTodoTaskRequest {
	s.DetailUrl = v
	return s
}

func (s *CreateTodoTaskRequest) SetRecurrence(v string) *CreateTodoTaskRequest {
	s.Recurrence = &v
	return s
}

func (s *CreateTodoTaskRequest) SetReminder(v *CreateTodoTaskRequestReminder) *CreateTodoTaskRequest {
	s.Reminder = v
	return s
}

func (s *CreateTodoTaskRequest) SetNotifyConfigs(v *CreateTodoTaskRequestNotifyConfigs) *CreateTodoTaskRequest {
	s.NotifyConfigs = v
	return s
}

func (s *CreateTodoTaskRequest) SetOperatorId(v string) *CreateTodoTaskRequest {
	s.OperatorId = &v
	return s
}

type CreateTodoTaskRequestDetailUrl struct {
	// app端详情页url
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
	// pc端详情页url
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
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

type CreateTodoTaskRequestReminder struct {
	// 提醒通道，1 弹框，2 短信，3 电话
	Channel *int32 `json:"channel,omitempty" xml:"channel,omitempty"`
	// 提醒规则列表
	Rules []*CreateTodoTaskRequestReminderRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
}

func (s CreateTodoTaskRequestReminder) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskRequestReminder) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskRequestReminder) SetChannel(v int32) *CreateTodoTaskRequestReminder {
	s.Channel = &v
	return s
}

func (s *CreateTodoTaskRequestReminder) SetRules(v []*CreateTodoTaskRequestReminderRules) *CreateTodoTaskRequestReminder {
	s.Rules = v
	return s
}

type CreateTodoTaskRequestReminderRules struct {
	// startTime:相对开始时间  											//  					dueTime:相对截止时间 											//						customTime: 绝对时间
	BaseTime *string `json:"baseTime,omitempty" xml:"baseTime,omitempty"`
	// 必须，baseTime 为 startTime 或者 dueTime 时表分钟；比如截止前15分钟为 -15，截止前3小时为 -180;basetime 为 customTime 时为时间戳
	Offset *int64 `json:"offset,omitempty" xml:"offset,omitempty"`
}

func (s CreateTodoTaskRequestReminderRules) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskRequestReminderRules) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskRequestReminderRules) SetBaseTime(v string) *CreateTodoTaskRequestReminderRules {
	s.BaseTime = &v
	return s
}

func (s *CreateTodoTaskRequestReminderRules) SetOffset(v int64) *CreateTodoTaskRequestReminderRules {
	s.Offset = &v
	return s
}

type CreateTodoTaskRequestNotifyConfigs struct {
	// 是否发送单聊卡片：value:"true/false" （发送则传true）
	SingleChat *string `json:"singleChat,omitempty" xml:"singleChat,omitempty"`
	// 是否发送群聊卡片：value:"groupId"（发送则传对应群聊id）
	GroupChat *string `json:"groupChat,omitempty" xml:"groupChat,omitempty"`
	// ding通知配置：value:"channel"（1钉弹框通知，2钉短信通知，3钉电话通知）
	DingNotify *string `json:"dingNotify,omitempty" xml:"dingNotify,omitempty"`
	// 是否同步到日历：value:"true/false"（同步则传true）
	Canlender *string `json:"canlender,omitempty" xml:"canlender,omitempty"`
}

func (s CreateTodoTaskRequestNotifyConfigs) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskRequestNotifyConfigs) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskRequestNotifyConfigs) SetSingleChat(v string) *CreateTodoTaskRequestNotifyConfigs {
	s.SingleChat = &v
	return s
}

func (s *CreateTodoTaskRequestNotifyConfigs) SetGroupChat(v string) *CreateTodoTaskRequestNotifyConfigs {
	s.GroupChat = &v
	return s
}

func (s *CreateTodoTaskRequestNotifyConfigs) SetDingNotify(v string) *CreateTodoTaskRequestNotifyConfigs {
	s.DingNotify = &v
	return s
}

func (s *CreateTodoTaskRequestNotifyConfigs) SetCanlender(v string) *CreateTodoTaskRequestNotifyConfigs {
	s.Canlender = &v
	return s
}

type CreateTodoTaskResponseBody struct {
	// id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 标题
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 截止时间
	DueTime *int64 `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	// 完成时间
	FinishTime *int64 `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// 完成状态
	Done *bool `json:"done,omitempty" xml:"done,omitempty"`
	// 执行者列表
	ExecutorIds []*string `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	// 参与者列表
	ParticipantIds []*string `json:"participantIds,omitempty" xml:"participantIds,omitempty" type:"Repeated"`
	// 提醒规则
	Reminder *CreateTodoTaskResponseBodyReminder `json:"reminder,omitempty" xml:"reminder,omitempty" type:"Struct"`
	// 待办通知配置（包含单聊卡片、ding通知、群聊卡片、同步日历、同步系统消息等通知能力）
	NotifyConfigs *CreateTodoTaskResponseBodyNotifyConfigs `json:"notifyConfigs,omitempty" xml:"notifyConfigs,omitempty" type:"Struct"`
	// 自定义详情页跳转配置
	DetailUrl *CreateTodoTaskResponseBodyDetailUrl `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	// 重复规则
	Recurrence *string `json:"recurrence,omitempty" xml:"recurrence,omitempty"`
	// 业务来源
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 业务来源id
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 创建时间
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 更新时间
	ModifiedTime *int64 `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// 创建者
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 更新者
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// 租户id
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 接入应用标识
	BizTag *string `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateTodoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskResponseBody) SetId(v string) *CreateTodoTaskResponseBody {
	s.Id = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetSubject(v string) *CreateTodoTaskResponseBody {
	s.Subject = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetDescription(v string) *CreateTodoTaskResponseBody {
	s.Description = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetStartTime(v int64) *CreateTodoTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetDueTime(v int64) *CreateTodoTaskResponseBody {
	s.DueTime = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetFinishTime(v int64) *CreateTodoTaskResponseBody {
	s.FinishTime = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetDone(v bool) *CreateTodoTaskResponseBody {
	s.Done = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetExecutorIds(v []*string) *CreateTodoTaskResponseBody {
	s.ExecutorIds = v
	return s
}

func (s *CreateTodoTaskResponseBody) SetParticipantIds(v []*string) *CreateTodoTaskResponseBody {
	s.ParticipantIds = v
	return s
}

func (s *CreateTodoTaskResponseBody) SetReminder(v *CreateTodoTaskResponseBodyReminder) *CreateTodoTaskResponseBody {
	s.Reminder = v
	return s
}

func (s *CreateTodoTaskResponseBody) SetNotifyConfigs(v *CreateTodoTaskResponseBodyNotifyConfigs) *CreateTodoTaskResponseBody {
	s.NotifyConfigs = v
	return s
}

func (s *CreateTodoTaskResponseBody) SetDetailUrl(v *CreateTodoTaskResponseBodyDetailUrl) *CreateTodoTaskResponseBody {
	s.DetailUrl = v
	return s
}

func (s *CreateTodoTaskResponseBody) SetRecurrence(v string) *CreateTodoTaskResponseBody {
	s.Recurrence = &v
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

func (s *CreateTodoTaskResponseBody) SetCreatedTime(v int64) *CreateTodoTaskResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetModifiedTime(v int64) *CreateTodoTaskResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetCreatorId(v string) *CreateTodoTaskResponseBody {
	s.CreatorId = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetModifierId(v string) *CreateTodoTaskResponseBody {
	s.ModifierId = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetTenantId(v string) *CreateTodoTaskResponseBody {
	s.TenantId = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetBizTag(v string) *CreateTodoTaskResponseBody {
	s.BizTag = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetRequestId(v string) *CreateTodoTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateTodoTaskResponseBodyReminder struct {
	// 提醒通道
	Channel *int32 `json:"channel,omitempty" xml:"channel,omitempty"`
	// 提醒规则配置
	Rules *CreateTodoTaskResponseBodyReminderRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Struct"`
}

func (s CreateTodoTaskResponseBodyReminder) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskResponseBodyReminder) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskResponseBodyReminder) SetChannel(v int32) *CreateTodoTaskResponseBodyReminder {
	s.Channel = &v
	return s
}

func (s *CreateTodoTaskResponseBodyReminder) SetRules(v *CreateTodoTaskResponseBodyReminderRules) *CreateTodoTaskResponseBodyReminder {
	s.Rules = v
	return s
}

type CreateTodoTaskResponseBodyReminderRules struct {
	// 目前支持三种类型：tartDate: 相对开始时间；dueDate: 相对截止时间；customDate: 绝对时间
	BaseTime *string `json:"baseTime,omitempty" xml:"baseTime,omitempty"`
	// 偏移值：baseTime 为 startDate 或者 dueDate 时，offset 为相对分钟的偏移值；baseTime 为 customDate 时，offset 为毫秒时间戳
	Offset *int64 `json:"offset,omitempty" xml:"offset,omitempty"`
}

func (s CreateTodoTaskResponseBodyReminderRules) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskResponseBodyReminderRules) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskResponseBodyReminderRules) SetBaseTime(v string) *CreateTodoTaskResponseBodyReminderRules {
	s.BaseTime = &v
	return s
}

func (s *CreateTodoTaskResponseBodyReminderRules) SetOffset(v int64) *CreateTodoTaskResponseBodyReminderRules {
	s.Offset = &v
	return s
}

type CreateTodoTaskResponseBodyNotifyConfigs struct {
	// 是否发送单聊卡片：value:"true/false" （发送则传true）
	SingleChat *string `json:"singleChat,omitempty" xml:"singleChat,omitempty"`
	// 是否发送群聊卡片：value:"groupId"（发送则传对应群聊id）
	GroupChat *string `json:"groupChat,omitempty" xml:"groupChat,omitempty"`
	// ding通知配置：value:"channel"（1钉弹框通知，2钉短信通知，3钉电话通知）
	DingNotify *string `json:"dingNotify,omitempty" xml:"dingNotify,omitempty"`
	// 是否同步到日历：value:"true/false"（同步则传true）
	Canlender *string `json:"canlender,omitempty" xml:"canlender,omitempty"`
}

func (s CreateTodoTaskResponseBodyNotifyConfigs) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskResponseBodyNotifyConfigs) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskResponseBodyNotifyConfigs) SetSingleChat(v string) *CreateTodoTaskResponseBodyNotifyConfigs {
	s.SingleChat = &v
	return s
}

func (s *CreateTodoTaskResponseBodyNotifyConfigs) SetGroupChat(v string) *CreateTodoTaskResponseBodyNotifyConfigs {
	s.GroupChat = &v
	return s
}

func (s *CreateTodoTaskResponseBodyNotifyConfigs) SetDingNotify(v string) *CreateTodoTaskResponseBodyNotifyConfigs {
	s.DingNotify = &v
	return s
}

func (s *CreateTodoTaskResponseBodyNotifyConfigs) SetCanlender(v string) *CreateTodoTaskResponseBodyNotifyConfigs {
	s.Canlender = &v
	return s
}

type CreateTodoTaskResponseBodyDetailUrl struct {
	// pc端详情页地址
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	// app端详情页地址
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
}

func (s CreateTodoTaskResponseBodyDetailUrl) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskResponseBodyDetailUrl) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskResponseBodyDetailUrl) SetPcUrl(v string) *CreateTodoTaskResponseBodyDetailUrl {
	s.PcUrl = &v
	return s
}

func (s *CreateTodoTaskResponseBodyDetailUrl) SetAppUrl(v string) *CreateTodoTaskResponseBodyDetailUrl {
	s.AppUrl = &v
	return s
}

type CreateTodoTaskResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTodoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateTodoTaskResponse) SetBody(v *CreateTodoTaskResponseBody) *CreateTodoTaskResponse {
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

func (client *Client) GetTodoTask(userId *string, taskId *string) (_result *GetTodoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTodoTaskHeaders{}
	_result = &GetTodoTaskResponse{}
	_body, _err := client.GetTodoTaskWithOptions(userId, taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTodoTaskWithOptions(userId *string, taskId *string, headers *GetTodoTaskHeaders, runtime *util.RuntimeOptions) (_result *GetTodoTaskResponse, _err error) {
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
	_result = &GetTodoTaskResponse{}
	_body, _err := client.DoROARequest(tea.String("GetTodoTask"), tea.String("todo_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/todo/users/"+tea.StringValue(userId)+"/tasks/"+tea.StringValue(taskId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTodoTask(userId *string, taskId *string, request *DeleteTodoTaskRequest) (_result *DeleteTodoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteTodoTaskHeaders{}
	_result = &DeleteTodoTaskResponse{}
	_body, _err := client.DeleteTodoTaskWithOptions(userId, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTodoTaskWithOptions(userId *string, taskId *string, request *DeleteTodoTaskRequest, headers *DeleteTodoTaskHeaders, runtime *util.RuntimeOptions) (_result *DeleteTodoTaskResponse, _err error) {
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
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteTodoTaskResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteTodoTask"), tea.String("todo_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/todo/users/"+tea.StringValue(userId)+"/tasks/"+tea.StringValue(taskId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTodoTask(userId *string, taskId *string, request *UpdateTodoTaskRequest) (_result *UpdateTodoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTodoTaskHeaders{}
	_result = &UpdateTodoTaskResponse{}
	_body, _err := client.UpdateTodoTaskWithOptions(userId, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTodoTaskWithOptions(userId *string, taskId *string, request *UpdateTodoTaskRequest, headers *UpdateTodoTaskHeaders, runtime *util.RuntimeOptions) (_result *UpdateTodoTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Sucject)) {
		body["sucject"] = request.Sucject
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DueTime)) {
		body["dueTime"] = request.DueTime
	}

	if !tea.BoolValue(util.IsUnset(request.Done)) {
		body["done"] = request.Done
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorIds)) {
		body["executorIds"] = request.ExecutorIds
	}

	if !tea.BoolValue(util.IsUnset(request.ParticipantIds)) {
		body["participantIds"] = request.ParticipantIds
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.DetailUrl))) {
		body["detailUrl"] = request.DetailUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Recurrence)) {
		body["recurrence"] = request.Recurrence
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Reminder))) {
		body["reminder"] = request.Reminder
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateTodoTaskResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateTodoTask"), tea.String("todo_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/todo/users/"+tea.StringValue(userId)+"/tasks/"+tea.StringValue(taskId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTodoTask(userId *string, request *CreateTodoTaskRequest) (_result *CreateTodoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTodoTaskHeaders{}
	_result = &CreateTodoTaskResponse{}
	_body, _err := client.CreateTodoTaskWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTodoTaskWithOptions(userId *string, request *CreateTodoTaskRequest, headers *CreateTodoTaskHeaders, runtime *util.RuntimeOptions) (_result *CreateTodoTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		body["sourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		body["subject"] = request.Subject
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorId)) {
		body["creatorId"] = request.CreatorId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
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

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.DetailUrl))) {
		body["detailUrl"] = request.DetailUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Recurrence)) {
		body["recurrence"] = request.Recurrence
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Reminder))) {
		body["reminder"] = request.Reminder
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.NotifyConfigs))) {
		body["notifyConfigs"] = request.NotifyConfigs
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateTodoTaskResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateTodoTask"), tea.String("todo_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/todo/users/"+tea.StringValue(userId)+"/tasks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}