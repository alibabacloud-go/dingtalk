// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package project_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddProjectMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddProjectMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddProjectMemberHeaders) GoString() string {
	return s.String()
}

func (s *AddProjectMemberHeaders) SetCommonHeaders(v map[string]*string) *AddProjectMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddProjectMemberHeaders) SetXAcsDingtalkAccessToken(v string) *AddProjectMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddProjectMemberRequest struct {
	// 用户ID列表，建议一次不超过10个
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s AddProjectMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s AddProjectMemberRequest) GoString() string {
	return s.String()
}

func (s *AddProjectMemberRequest) SetUserIds(v []*string) *AddProjectMemberRequest {
	s.UserIds = v
	return s
}

type AddProjectMemberResponseBody struct {
	Result []*AddProjectMemberResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s AddProjectMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddProjectMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddProjectMemberResponseBody) SetResult(v []*AddProjectMemberResponseBodyResult) *AddProjectMemberResponseBody {
	s.Result = v
	return s
}

type AddProjectMemberResponseBodyResult struct {
	Joined   *string `json:"joined,omitempty" xml:"joined,omitempty"`
	Nickname *string `json:"nickname,omitempty" xml:"nickname,omitempty"`
}

func (s AddProjectMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddProjectMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddProjectMemberResponseBodyResult) SetJoined(v string) *AddProjectMemberResponseBodyResult {
	s.Joined = &v
	return s
}

func (s *AddProjectMemberResponseBodyResult) SetNickname(v string) *AddProjectMemberResponseBodyResult {
	s.Nickname = &v
	return s
}

type AddProjectMemberResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddProjectMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddProjectMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s AddProjectMemberResponse) GoString() string {
	return s.String()
}

func (s *AddProjectMemberResponse) SetHeaders(v map[string]*string) *AddProjectMemberResponse {
	s.Headers = v
	return s
}

func (s *AddProjectMemberResponse) SetBody(v *AddProjectMemberResponseBody) *AddProjectMemberResponse {
	s.Body = v
	return s
}

type CreateOrganizationTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateOrganizationTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateOrganizationTaskHeaders) GoString() string {
	return s.String()
}

func (s *CreateOrganizationTaskHeaders) SetCommonHeaders(v map[string]*string) *CreateOrganizationTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateOrganizationTaskHeaders) SetXAcsDingtalkAccessToken(v string) *CreateOrganizationTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateOrganizationTaskRequest struct {
	// 任务标题
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 任务创建日期
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 是否禁止动态
	DisableActivity *bool `json:"disableActivity,omitempty" xml:"disableActivity,omitempty"`
	// 是否禁止通知
	DisableNotification *bool `json:"disableNotification,omitempty" xml:"disableNotification,omitempty"`
	// 任务截止日期
	DueDate *string `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	// 执行者id
	ExecutorId *string `json:"executorId,omitempty" xml:"executorId,omitempty"`
	// 参与者id
	InvolveMembers []*string `json:"involveMembers,omitempty" xml:"involveMembers,omitempty" type:"Repeated"`
	// 任务备注
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
	// 优先级【-10,0,1,2】中选一个
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// 任务可见性。involves：仅参与者可见。members:所有人可见
	Visible *string `json:"visible,omitempty" xml:"visible,omitempty"`
}

func (s CreateOrganizationTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrganizationTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateOrganizationTaskRequest) SetContent(v string) *CreateOrganizationTaskRequest {
	s.Content = &v
	return s
}

func (s *CreateOrganizationTaskRequest) SetCreateTime(v string) *CreateOrganizationTaskRequest {
	s.CreateTime = &v
	return s
}

func (s *CreateOrganizationTaskRequest) SetDisableActivity(v bool) *CreateOrganizationTaskRequest {
	s.DisableActivity = &v
	return s
}

func (s *CreateOrganizationTaskRequest) SetDisableNotification(v bool) *CreateOrganizationTaskRequest {
	s.DisableNotification = &v
	return s
}

func (s *CreateOrganizationTaskRequest) SetDueDate(v string) *CreateOrganizationTaskRequest {
	s.DueDate = &v
	return s
}

func (s *CreateOrganizationTaskRequest) SetExecutorId(v string) *CreateOrganizationTaskRequest {
	s.ExecutorId = &v
	return s
}

func (s *CreateOrganizationTaskRequest) SetInvolveMembers(v []*string) *CreateOrganizationTaskRequest {
	s.InvolveMembers = v
	return s
}

func (s *CreateOrganizationTaskRequest) SetNote(v string) *CreateOrganizationTaskRequest {
	s.Note = &v
	return s
}

func (s *CreateOrganizationTaskRequest) SetPriority(v int32) *CreateOrganizationTaskRequest {
	s.Priority = &v
	return s
}

func (s *CreateOrganizationTaskRequest) SetVisible(v string) *CreateOrganizationTaskRequest {
	s.Visible = &v
	return s
}

type CreateOrganizationTaskResponseBody struct {
	// 返回结果对象
	Result *CreateOrganizationTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateOrganizationTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrganizationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrganizationTaskResponseBody) SetResult(v *CreateOrganizationTaskResponseBodyResult) *CreateOrganizationTaskResponseBody {
	s.Result = v
	return s
}

type CreateOrganizationTaskResponseBodyResult struct {
	// 父任务Id
	AncestorIds []*string `json:"ancestorIds,omitempty" xml:"ancestorIds,omitempty" type:"Repeated"`
	// 附件数量
	AttachmentsCount *int32 `json:"attachmentsCount,omitempty" xml:"attachmentsCount,omitempty"`
	// 任务标题
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 创建时间
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// 创建者
	Creator *CreateOrganizationTaskResponseBodyResultCreator `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	// 创建者id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 任务截止日期
	DueDate *string `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	// 执行者
	Executor *CreateOrganizationTaskResponseBodyResultExecutor `json:"executor,omitempty" xml:"executor,omitempty" type:"Struct"`
	// 执行者id
	ExecutorId *string `json:"executorId,omitempty" xml:"executorId,omitempty"`
	// 是否有提醒
	HasReminder *bool `json:"hasReminder,omitempty" xml:"hasReminder,omitempty"`
	// 任务id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 参与者id列表
	InvolveMembers []*string `json:"involveMembers,omitempty" xml:"involveMembers,omitempty" type:"Repeated"`
	// 参与者列表
	Involvers []*CreateOrganizationTaskResponseBodyResultInvolvers `json:"involvers,omitempty" xml:"involvers,omitempty" type:"Repeated"`
	// 是否删除
	IsDeleted *bool `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	// 是否完成
	IsDone *string `json:"isDone,omitempty" xml:"isDone,omitempty"`
	// 任务备注
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
	// 优先级【-10,0,1,2】中选一个
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// 更新时间
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
	// 任务可见性。involves：仅参与者可见。members:所有人可见
	Visible *string `json:"visible,omitempty" xml:"visible,omitempty"`
}

func (s CreateOrganizationTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateOrganizationTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateOrganizationTaskResponseBodyResult) SetAncestorIds(v []*string) *CreateOrganizationTaskResponseBodyResult {
	s.AncestorIds = v
	return s
}

func (s *CreateOrganizationTaskResponseBodyResult) SetAttachmentsCount(v int32) *CreateOrganizationTaskResponseBodyResult {
	s.AttachmentsCount = &v
	return s
}

func (s *CreateOrganizationTaskResponseBodyResult) SetContent(v string) *CreateOrganizationTaskResponseBodyResult {
	s.Content = &v
	return s
}

func (s *CreateOrganizationTaskResponseBodyResult) SetCreated(v string) *CreateOrganizationTaskResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateOrganizationTaskResponseBodyResult) SetCreator(v *CreateOrganizationTaskResponseBodyResultCreator) *CreateOrganizationTaskResponseBodyResult {
	s.Creator = v
	return s
}

func (s *CreateOrganizationTaskResponseBodyResult) SetCreatorId(v string) *CreateOrganizationTaskResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *CreateOrganizationTaskResponseBodyResult) SetDueDate(v string) *CreateOrganizationTaskResponseBodyResult {
	s.DueDate = &v
	return s
}

func (s *CreateOrganizationTaskResponseBodyResult) SetExecutor(v *CreateOrganizationTaskResponseBodyResultExecutor) *CreateOrganizationTaskResponseBodyResult {
	s.Executor = v
	return s
}

func (s *CreateOrganizationTaskResponseBodyResult) SetExecutorId(v string) *CreateOrganizationTaskResponseBodyResult {
	s.ExecutorId = &v
	return s
}

func (s *CreateOrganizationTaskResponseBodyResult) SetHasReminder(v bool) *CreateOrganizationTaskResponseBodyResult {
	s.HasReminder = &v
	return s
}

func (s *CreateOrganizationTaskResponseBodyResult) SetId(v string) *CreateOrganizationTaskResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateOrganizationTaskResponseBodyResult) SetInvolveMembers(v []*string) *CreateOrganizationTaskResponseBodyResult {
	s.InvolveMembers = v
	return s
}

func (s *CreateOrganizationTaskResponseBodyResult) SetInvolvers(v []*CreateOrganizationTaskResponseBodyResultInvolvers) *CreateOrganizationTaskResponseBodyResult {
	s.Involvers = v
	return s
}

func (s *CreateOrganizationTaskResponseBodyResult) SetIsDeleted(v bool) *CreateOrganizationTaskResponseBodyResult {
	s.IsDeleted = &v
	return s
}

func (s *CreateOrganizationTaskResponseBodyResult) SetIsDone(v string) *CreateOrganizationTaskResponseBodyResult {
	s.IsDone = &v
	return s
}

func (s *CreateOrganizationTaskResponseBodyResult) SetNote(v string) *CreateOrganizationTaskResponseBodyResult {
	s.Note = &v
	return s
}

func (s *CreateOrganizationTaskResponseBodyResult) SetPriority(v int32) *CreateOrganizationTaskResponseBodyResult {
	s.Priority = &v
	return s
}

func (s *CreateOrganizationTaskResponseBodyResult) SetUpdated(v string) *CreateOrganizationTaskResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *CreateOrganizationTaskResponseBodyResult) SetVisible(v string) *CreateOrganizationTaskResponseBodyResult {
	s.Visible = &v
	return s
}

type CreateOrganizationTaskResponseBodyResultCreator struct {
	// 创建者头像地址
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// 创建者姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 创建者id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateOrganizationTaskResponseBodyResultCreator) String() string {
	return tea.Prettify(s)
}

func (s CreateOrganizationTaskResponseBodyResultCreator) GoString() string {
	return s.String()
}

func (s *CreateOrganizationTaskResponseBodyResultCreator) SetAvatarUrl(v string) *CreateOrganizationTaskResponseBodyResultCreator {
	s.AvatarUrl = &v
	return s
}

func (s *CreateOrganizationTaskResponseBodyResultCreator) SetName(v string) *CreateOrganizationTaskResponseBodyResultCreator {
	s.Name = &v
	return s
}

func (s *CreateOrganizationTaskResponseBodyResultCreator) SetUserId(v string) *CreateOrganizationTaskResponseBodyResultCreator {
	s.UserId = &v
	return s
}

type CreateOrganizationTaskResponseBodyResultExecutor struct {
	// 头像地址
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// 姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 执行者id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateOrganizationTaskResponseBodyResultExecutor) String() string {
	return tea.Prettify(s)
}

func (s CreateOrganizationTaskResponseBodyResultExecutor) GoString() string {
	return s.String()
}

func (s *CreateOrganizationTaskResponseBodyResultExecutor) SetAvatarUrl(v string) *CreateOrganizationTaskResponseBodyResultExecutor {
	s.AvatarUrl = &v
	return s
}

func (s *CreateOrganizationTaskResponseBodyResultExecutor) SetName(v string) *CreateOrganizationTaskResponseBodyResultExecutor {
	s.Name = &v
	return s
}

func (s *CreateOrganizationTaskResponseBodyResultExecutor) SetUserId(v string) *CreateOrganizationTaskResponseBodyResultExecutor {
	s.UserId = &v
	return s
}

type CreateOrganizationTaskResponseBodyResultInvolvers struct {
	// 头像
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// 用户id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 名字
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateOrganizationTaskResponseBodyResultInvolvers) String() string {
	return tea.Prettify(s)
}

func (s CreateOrganizationTaskResponseBodyResultInvolvers) GoString() string {
	return s.String()
}

func (s *CreateOrganizationTaskResponseBodyResultInvolvers) SetAvatarUrl(v string) *CreateOrganizationTaskResponseBodyResultInvolvers {
	s.AvatarUrl = &v
	return s
}

func (s *CreateOrganizationTaskResponseBodyResultInvolvers) SetId(v string) *CreateOrganizationTaskResponseBodyResultInvolvers {
	s.Id = &v
	return s
}

func (s *CreateOrganizationTaskResponseBodyResultInvolvers) SetName(v string) *CreateOrganizationTaskResponseBodyResultInvolvers {
	s.Name = &v
	return s
}

type CreateOrganizationTaskResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateOrganizationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOrganizationTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrganizationTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateOrganizationTaskResponse) SetHeaders(v map[string]*string) *CreateOrganizationTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateOrganizationTaskResponse) SetBody(v *CreateOrganizationTaskResponseBody) *CreateOrganizationTaskResponse {
	s.Body = v
	return s
}

type CreatePlanTimeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreatePlanTimeHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreatePlanTimeHeaders) GoString() string {
	return s.String()
}

func (s *CreatePlanTimeHeaders) SetCommonHeaders(v map[string]*string) *CreatePlanTimeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreatePlanTimeHeaders) SetXAcsDingtalkAccessToken(v string) *CreatePlanTimeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreatePlanTimeRequest struct {
	// 结束时间
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// 执行者userid
	ExecutorId *string `json:"executorId,omitempty" xml:"executorId,omitempty"`
	// 是否包含假期
	IncludesHolidays *bool `json:"includesHolidays,omitempty" xml:"includesHolidays,omitempty"`
	// 是否连续
	IsDuration *bool `json:"isDuration,omitempty" xml:"isDuration,omitempty"`
	// 对象id，比如任务id
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// 对象类型，默认为task
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// 计划工时数（单位：毫秒，1小时即为 3600000）
	PlanTime *int64 `json:"planTime,omitempty" xml:"planTime,omitempty"`
	// 开始时间
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// 工时所属人员userid
	SubmitterId *string `json:"submitterId,omitempty" xml:"submitterId,omitempty"`
	// 接口校验类型，当前默认organization
	TenantType *string `json:"tenantType,omitempty" xml:"tenantType,omitempty"`
}

func (s CreatePlanTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePlanTimeRequest) GoString() string {
	return s.String()
}

func (s *CreatePlanTimeRequest) SetEndDate(v string) *CreatePlanTimeRequest {
	s.EndDate = &v
	return s
}

func (s *CreatePlanTimeRequest) SetExecutorId(v string) *CreatePlanTimeRequest {
	s.ExecutorId = &v
	return s
}

func (s *CreatePlanTimeRequest) SetIncludesHolidays(v bool) *CreatePlanTimeRequest {
	s.IncludesHolidays = &v
	return s
}

func (s *CreatePlanTimeRequest) SetIsDuration(v bool) *CreatePlanTimeRequest {
	s.IsDuration = &v
	return s
}

func (s *CreatePlanTimeRequest) SetObjectId(v string) *CreatePlanTimeRequest {
	s.ObjectId = &v
	return s
}

func (s *CreatePlanTimeRequest) SetObjectType(v string) *CreatePlanTimeRequest {
	s.ObjectType = &v
	return s
}

func (s *CreatePlanTimeRequest) SetPlanTime(v int64) *CreatePlanTimeRequest {
	s.PlanTime = &v
	return s
}

func (s *CreatePlanTimeRequest) SetStartDate(v string) *CreatePlanTimeRequest {
	s.StartDate = &v
	return s
}

func (s *CreatePlanTimeRequest) SetSubmitterId(v string) *CreatePlanTimeRequest {
	s.SubmitterId = &v
	return s
}

func (s *CreatePlanTimeRequest) SetTenantType(v string) *CreatePlanTimeRequest {
	s.TenantType = &v
	return s
}

type CreatePlanTimeResponseBody struct {
	Result *CreatePlanTimeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreatePlanTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePlanTimeResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePlanTimeResponseBody) SetResult(v *CreatePlanTimeResponseBodyResult) *CreatePlanTimeResponseBody {
	s.Result = v
	return s
}

type CreatePlanTimeResponseBodyResult struct {
	Body []*CreatePlanTimeResponseBodyResultBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// 执行结果描述
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	Ok      *bool   `json:"ok,omitempty" xml:"ok,omitempty"`
}

func (s CreatePlanTimeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreatePlanTimeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreatePlanTimeResponseBodyResult) SetBody(v []*CreatePlanTimeResponseBodyResultBody) *CreatePlanTimeResponseBodyResult {
	s.Body = v
	return s
}

func (s *CreatePlanTimeResponseBodyResult) SetMessage(v string) *CreatePlanTimeResponseBodyResult {
	s.Message = &v
	return s
}

func (s *CreatePlanTimeResponseBodyResult) SetOk(v bool) *CreatePlanTimeResponseBodyResult {
	s.Ok = &v
	return s
}

type CreatePlanTimeResponseBodyResultBody struct {
	// 更新工时所属日期
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// 工时关联的数据id
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// 计划工时数
	PlanTime *int64 `json:"planTime,omitempty" xml:"planTime,omitempty"`
}

func (s CreatePlanTimeResponseBodyResultBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePlanTimeResponseBodyResultBody) GoString() string {
	return s.String()
}

func (s *CreatePlanTimeResponseBodyResultBody) SetDate(v string) *CreatePlanTimeResponseBodyResultBody {
	s.Date = &v
	return s
}

func (s *CreatePlanTimeResponseBodyResultBody) SetObjectId(v string) *CreatePlanTimeResponseBodyResultBody {
	s.ObjectId = &v
	return s
}

func (s *CreatePlanTimeResponseBodyResultBody) SetPlanTime(v int64) *CreatePlanTimeResponseBodyResultBody {
	s.PlanTime = &v
	return s
}

type CreatePlanTimeResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatePlanTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePlanTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePlanTimeResponse) GoString() string {
	return s.String()
}

func (s *CreatePlanTimeResponse) SetHeaders(v map[string]*string) *CreatePlanTimeResponse {
	s.Headers = v
	return s
}

func (s *CreatePlanTimeResponse) SetBody(v *CreatePlanTimeResponseBody) *CreatePlanTimeResponse {
	s.Body = v
	return s
}

type CreateProjectByTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateProjectByTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectByTemplateHeaders) GoString() string {
	return s.String()
}

func (s *CreateProjectByTemplateHeaders) SetCommonHeaders(v map[string]*string) *CreateProjectByTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateProjectByTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *CreateProjectByTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateProjectByTemplateRequest struct {
	// 项目名字
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 模板ID
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s CreateProjectByTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectByTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectByTemplateRequest) SetName(v string) *CreateProjectByTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectByTemplateRequest) SetTemplateId(v string) *CreateProjectByTemplateRequest {
	s.TemplateId = &v
	return s
}

type CreateProjectByTemplateResponseBody struct {
	// 返回结果对象
	Result *CreateProjectByTemplateResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateProjectByTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectByTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectByTemplateResponseBody) SetResult(v *CreateProjectByTemplateResponseBodyResult) *CreateProjectByTemplateResponseBody {
	s.Result = v
	return s
}

type CreateProjectByTemplateResponseBodyResult struct {
	// 创建时间
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// 项目ID
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 项目图标地址
	Logo *string `json:"logo,omitempty" xml:"logo,omitempty"`
	// 项目名字
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateProjectByTemplateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectByTemplateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateProjectByTemplateResponseBodyResult) SetCreated(v string) *CreateProjectByTemplateResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateProjectByTemplateResponseBodyResult) SetId(v string) *CreateProjectByTemplateResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateProjectByTemplateResponseBodyResult) SetLogo(v string) *CreateProjectByTemplateResponseBodyResult {
	s.Logo = &v
	return s
}

func (s *CreateProjectByTemplateResponseBodyResult) SetName(v string) *CreateProjectByTemplateResponseBodyResult {
	s.Name = &v
	return s
}

type CreateProjectByTemplateResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateProjectByTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProjectByTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectByTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectByTemplateResponse) SetHeaders(v map[string]*string) *CreateProjectByTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectByTemplateResponse) SetBody(v *CreateProjectByTemplateResponseBody) *CreateProjectByTemplateResponse {
	s.Body = v
	return s
}

type CreateTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskHeaders) GoString() string {
	return s.String()
}

func (s *CreateTaskHeaders) SetCommonHeaders(v map[string]*string) *CreateTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTaskHeaders) SetXAcsDingtalkAccessToken(v string) *CreateTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateTaskRequest struct {
	// 任务标题
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 自定义字段列表
	Customfields []*CreateTaskRequestCustomfields `json:"customfields,omitempty" xml:"customfields,omitempty" type:"Repeated"`
	// 任务截止时间
	DueDate *string `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	// 执行者userId
	ExecutorId *string `json:"executorId,omitempty" xml:"executorId,omitempty"`
	// 任务备注
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
	// 任务优先级
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// 项目id
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s CreateTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskRequest) SetContent(v string) *CreateTaskRequest {
	s.Content = &v
	return s
}

func (s *CreateTaskRequest) SetCustomfields(v []*CreateTaskRequestCustomfields) *CreateTaskRequest {
	s.Customfields = v
	return s
}

func (s *CreateTaskRequest) SetDueDate(v string) *CreateTaskRequest {
	s.DueDate = &v
	return s
}

func (s *CreateTaskRequest) SetExecutorId(v string) *CreateTaskRequest {
	s.ExecutorId = &v
	return s
}

func (s *CreateTaskRequest) SetNote(v string) *CreateTaskRequest {
	s.Note = &v
	return s
}

func (s *CreateTaskRequest) SetPriority(v int32) *CreateTaskRequest {
	s.Priority = &v
	return s
}

func (s *CreateTaskRequest) SetProjectId(v string) *CreateTaskRequest {
	s.ProjectId = &v
	return s
}

type CreateTaskRequestCustomfields struct {
	// 自定义字段id
	CustomfieldId *string `json:"customfieldId,omitempty" xml:"customfieldId,omitempty"`
	// 自定义字段名称
	CustomfieldName *string `json:"customfieldName,omitempty" xml:"customfieldName,omitempty"`
	// 自定义字段值
	Value []*CreateTaskRequestCustomfieldsValue `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s CreateTaskRequestCustomfields) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskRequestCustomfields) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestCustomfields) SetCustomfieldId(v string) *CreateTaskRequestCustomfields {
	s.CustomfieldId = &v
	return s
}

func (s *CreateTaskRequestCustomfields) SetCustomfieldName(v string) *CreateTaskRequestCustomfields {
	s.CustomfieldName = &v
	return s
}

func (s *CreateTaskRequestCustomfields) SetValue(v []*CreateTaskRequestCustomfieldsValue) *CreateTaskRequestCustomfields {
	s.Value = v
	return s
}

type CreateTaskRequestCustomfieldsValue struct {
	// 自定义字段显示值
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateTaskRequestCustomfieldsValue) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskRequestCustomfieldsValue) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestCustomfieldsValue) SetTitle(v string) *CreateTaskRequestCustomfieldsValue {
	s.Title = &v
	return s
}

type CreateTaskResponseBody struct {
	// 返回结果对象
	Result *CreateTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskResponseBody) SetResult(v *CreateTaskResponseBodyResult) *CreateTaskResponseBody {
	s.Result = v
	return s
}

type CreateTaskResponseBodyResult struct {
	// 任务标题
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 创建时间
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// 任务创建者userId
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 自定义字段列表
	Customfields []*CreateTaskResponseBodyResultCustomfields `json:"customfields,omitempty" xml:"customfields,omitempty" type:"Repeated"`
	// 任务截止时间
	DueDate *string `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	// 任务执行者userId
	ExecutorId *string `json:"executorId,omitempty" xml:"executorId,omitempty"`
	// 任务参与者列表
	InvolveMembers []*string `json:"involveMembers,omitempty" xml:"involveMembers,omitempty" type:"Repeated"`
	// 任务备注
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
	// 任务优先级
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// 项目id
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// 任务id
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// 更新时间
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s CreateTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateTaskResponseBodyResult) SetContent(v string) *CreateTaskResponseBodyResult {
	s.Content = &v
	return s
}

func (s *CreateTaskResponseBodyResult) SetCreated(v string) *CreateTaskResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateTaskResponseBodyResult) SetCreatorId(v string) *CreateTaskResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *CreateTaskResponseBodyResult) SetCustomfields(v []*CreateTaskResponseBodyResultCustomfields) *CreateTaskResponseBodyResult {
	s.Customfields = v
	return s
}

func (s *CreateTaskResponseBodyResult) SetDueDate(v string) *CreateTaskResponseBodyResult {
	s.DueDate = &v
	return s
}

func (s *CreateTaskResponseBodyResult) SetExecutorId(v string) *CreateTaskResponseBodyResult {
	s.ExecutorId = &v
	return s
}

func (s *CreateTaskResponseBodyResult) SetInvolveMembers(v []*string) *CreateTaskResponseBodyResult {
	s.InvolveMembers = v
	return s
}

func (s *CreateTaskResponseBodyResult) SetNote(v string) *CreateTaskResponseBodyResult {
	s.Note = &v
	return s
}

func (s *CreateTaskResponseBodyResult) SetPriority(v int32) *CreateTaskResponseBodyResult {
	s.Priority = &v
	return s
}

func (s *CreateTaskResponseBodyResult) SetProjectId(v string) *CreateTaskResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *CreateTaskResponseBodyResult) SetTaskId(v string) *CreateTaskResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *CreateTaskResponseBodyResult) SetUpdated(v string) *CreateTaskResponseBodyResult {
	s.Updated = &v
	return s
}

type CreateTaskResponseBodyResultCustomfields struct {
	// 自定义字段id
	CustomfieldId *string `json:"customfieldId,omitempty" xml:"customfieldId,omitempty"`
	// 自定义字段值
	Value []*CreateTaskResponseBodyResultCustomfieldsValue `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s CreateTaskResponseBodyResultCustomfields) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskResponseBodyResultCustomfields) GoString() string {
	return s.String()
}

func (s *CreateTaskResponseBodyResultCustomfields) SetCustomfieldId(v string) *CreateTaskResponseBodyResultCustomfields {
	s.CustomfieldId = &v
	return s
}

func (s *CreateTaskResponseBodyResultCustomfields) SetValue(v []*CreateTaskResponseBodyResultCustomfieldsValue) *CreateTaskResponseBodyResultCustomfields {
	s.Value = v
	return s
}

type CreateTaskResponseBodyResultCustomfieldsValue struct {
	// 自定义字段显示值
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateTaskResponseBodyResultCustomfieldsValue) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskResponseBodyResultCustomfieldsValue) GoString() string {
	return s.String()
}

func (s *CreateTaskResponseBodyResultCustomfieldsValue) SetTitle(v string) *CreateTaskResponseBodyResultCustomfieldsValue {
	s.Title = &v
	return s
}

type CreateTaskResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateTaskResponse) SetHeaders(v map[string]*string) *CreateTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateTaskResponse) SetBody(v *CreateTaskResponseBody) *CreateTaskResponse {
	s.Body = v
	return s
}

type CreateTaskObjectLinkHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateTaskObjectLinkHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskObjectLinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateTaskObjectLinkHeaders) SetCommonHeaders(v map[string]*string) *CreateTaskObjectLinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTaskObjectLinkHeaders) SetXAcsDingtalkAccessToken(v string) *CreateTaskObjectLinkHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateTaskObjectLinkRequest struct {
	// 关联链接对象
	LinkedData *CreateTaskObjectLinkRequestLinkedData `json:"linkedData,omitempty" xml:"linkedData,omitempty" type:"Struct"`
}

func (s CreateTaskObjectLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskObjectLinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskObjectLinkRequest) SetLinkedData(v *CreateTaskObjectLinkRequestLinkedData) *CreateTaskObjectLinkRequest {
	s.LinkedData = v
	return s
}

type CreateTaskObjectLinkRequestLinkedData struct {
	// 关联对象描述
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 关联对象头像url
	ThumbnailUrl *string `json:"thumbnailUrl,omitempty" xml:"thumbnailUrl,omitempty"`
	// 关联对象标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 关联对象链接url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateTaskObjectLinkRequestLinkedData) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskObjectLinkRequestLinkedData) GoString() string {
	return s.String()
}

func (s *CreateTaskObjectLinkRequestLinkedData) SetContent(v string) *CreateTaskObjectLinkRequestLinkedData {
	s.Content = &v
	return s
}

func (s *CreateTaskObjectLinkRequestLinkedData) SetThumbnailUrl(v string) *CreateTaskObjectLinkRequestLinkedData {
	s.ThumbnailUrl = &v
	return s
}

func (s *CreateTaskObjectLinkRequestLinkedData) SetTitle(v string) *CreateTaskObjectLinkRequestLinkedData {
	s.Title = &v
	return s
}

func (s *CreateTaskObjectLinkRequestLinkedData) SetUrl(v string) *CreateTaskObjectLinkRequestLinkedData {
	s.Url = &v
	return s
}

type CreateTaskObjectLinkResponseBody struct {
	// 返回结果对象
	Result *CreateTaskObjectLinkResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateTaskObjectLinkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskObjectLinkResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskObjectLinkResponseBody) SetResult(v *CreateTaskObjectLinkResponseBodyResult) *CreateTaskObjectLinkResponseBody {
	s.Result = v
	return s
}

type CreateTaskObjectLinkResponseBodyResult struct {
	// 创建时间
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// 关联对象id
	ObjectLinkId *string `json:"objectLinkId,omitempty" xml:"objectLinkId,omitempty"`
}

func (s CreateTaskObjectLinkResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskObjectLinkResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateTaskObjectLinkResponseBodyResult) SetCreated(v string) *CreateTaskObjectLinkResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateTaskObjectLinkResponseBodyResult) SetObjectLinkId(v string) *CreateTaskObjectLinkResponseBodyResult {
	s.ObjectLinkId = &v
	return s
}

type CreateTaskObjectLinkResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTaskObjectLinkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTaskObjectLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskObjectLinkResponse) GoString() string {
	return s.String()
}

func (s *CreateTaskObjectLinkResponse) SetHeaders(v map[string]*string) *CreateTaskObjectLinkResponse {
	s.Headers = v
	return s
}

func (s *CreateTaskObjectLinkResponse) SetBody(v *CreateTaskObjectLinkResponseBody) *CreateTaskObjectLinkResponse {
	s.Body = v
	return s
}

type CreateWorkTimeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateWorkTimeHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkTimeHeaders) GoString() string {
	return s.String()
}

func (s *CreateWorkTimeHeaders) SetCommonHeaders(v map[string]*string) *CreateWorkTimeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateWorkTimeHeaders) SetXAcsDingtalkAccessToken(v string) *CreateWorkTimeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateWorkTimeRequest struct {
	// 结束时间
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// 执行者userid
	ExecutorId *string `json:"executorId,omitempty" xml:"executorId,omitempty"`
	// 是否包含节假日
	IncludesHolidays *bool `json:"includesHolidays,omitempty" xml:"includesHolidays,omitempty"`
	// 是否连续
	IsDuration *bool `json:"isDuration,omitempty" xml:"isDuration,omitempty"`
	// 对象 ID，比如 任务 ID
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// 对象类型，默认为 task
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// 开始时间
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// 工时所属人员userid
	SubmitterId *string `json:"submitterId,omitempty" xml:"submitterId,omitempty"`
	// 实际工时数（单位毫秒，1小时即为3600000）
	WorkTime *int64 `json:"workTime,omitempty" xml:"workTime,omitempty"`
	// 接口校验类型，当前默认organization
	TenantType *string `json:"tenantType,omitempty" xml:"tenantType,omitempty"`
}

func (s CreateWorkTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkTimeRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkTimeRequest) SetEndDate(v string) *CreateWorkTimeRequest {
	s.EndDate = &v
	return s
}

func (s *CreateWorkTimeRequest) SetExecutorId(v string) *CreateWorkTimeRequest {
	s.ExecutorId = &v
	return s
}

func (s *CreateWorkTimeRequest) SetIncludesHolidays(v bool) *CreateWorkTimeRequest {
	s.IncludesHolidays = &v
	return s
}

func (s *CreateWorkTimeRequest) SetIsDuration(v bool) *CreateWorkTimeRequest {
	s.IsDuration = &v
	return s
}

func (s *CreateWorkTimeRequest) SetObjectId(v string) *CreateWorkTimeRequest {
	s.ObjectId = &v
	return s
}

func (s *CreateWorkTimeRequest) SetObjectType(v string) *CreateWorkTimeRequest {
	s.ObjectType = &v
	return s
}

func (s *CreateWorkTimeRequest) SetStartDate(v string) *CreateWorkTimeRequest {
	s.StartDate = &v
	return s
}

func (s *CreateWorkTimeRequest) SetSubmitterId(v string) *CreateWorkTimeRequest {
	s.SubmitterId = &v
	return s
}

func (s *CreateWorkTimeRequest) SetWorkTime(v int64) *CreateWorkTimeRequest {
	s.WorkTime = &v
	return s
}

func (s *CreateWorkTimeRequest) SetTenantType(v string) *CreateWorkTimeRequest {
	s.TenantType = &v
	return s
}

type CreateWorkTimeResponseBody struct {
	Result *CreateWorkTimeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateWorkTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkTimeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkTimeResponseBody) SetResult(v *CreateWorkTimeResponseBodyResult) *CreateWorkTimeResponseBody {
	s.Result = v
	return s
}

type CreateWorkTimeResponseBodyResult struct {
	Body []*CreateWorkTimeResponseBodyResultBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// 执行结果描述
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	Ok      *bool   `json:"ok,omitempty" xml:"ok,omitempty"`
}

func (s CreateWorkTimeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkTimeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateWorkTimeResponseBodyResult) SetBody(v []*CreateWorkTimeResponseBodyResultBody) *CreateWorkTimeResponseBodyResult {
	s.Body = v
	return s
}

func (s *CreateWorkTimeResponseBodyResult) SetMessage(v string) *CreateWorkTimeResponseBodyResult {
	s.Message = &v
	return s
}

func (s *CreateWorkTimeResponseBodyResult) SetOk(v bool) *CreateWorkTimeResponseBodyResult {
	s.Ok = &v
	return s
}

type CreateWorkTimeResponseBodyResultBody struct {
	// 工时所属日期
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// 工时关联的数据 ID
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// 实际工时
	WorkTime *int64 `json:"workTime,omitempty" xml:"workTime,omitempty"`
}

func (s CreateWorkTimeResponseBodyResultBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkTimeResponseBodyResultBody) GoString() string {
	return s.String()
}

func (s *CreateWorkTimeResponseBodyResultBody) SetDate(v string) *CreateWorkTimeResponseBodyResultBody {
	s.Date = &v
	return s
}

func (s *CreateWorkTimeResponseBodyResultBody) SetTaskId(v string) *CreateWorkTimeResponseBodyResultBody {
	s.TaskId = &v
	return s
}

func (s *CreateWorkTimeResponseBodyResultBody) SetWorkTime(v int64) *CreateWorkTimeResponseBodyResultBody {
	s.WorkTime = &v
	return s
}

type CreateWorkTimeResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateWorkTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWorkTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkTimeResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkTimeResponse) SetHeaders(v map[string]*string) *CreateWorkTimeResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkTimeResponse) SetBody(v *CreateWorkTimeResponseBody) *CreateWorkTimeResponse {
	s.Body = v
	return s
}

type GetDeptsByOrgIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	DingAccessTokenType     *string            `json:"dingAccessTokenType,omitempty" xml:"dingAccessTokenType,omitempty"`
	DingOrgId               *string            `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDeptsByOrgIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDeptsByOrgIdHeaders) GoString() string {
	return s.String()
}

func (s *GetDeptsByOrgIdHeaders) SetCommonHeaders(v map[string]*string) *GetDeptsByOrgIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeptsByOrgIdHeaders) SetDingAccessTokenType(v string) *GetDeptsByOrgIdHeaders {
	s.DingAccessTokenType = &v
	return s
}

func (s *GetDeptsByOrgIdHeaders) SetDingOrgId(v string) *GetDeptsByOrgIdHeaders {
	s.DingOrgId = &v
	return s
}

func (s *GetDeptsByOrgIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetDeptsByOrgIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDeptsByOrgIdRequest struct {
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetDeptsByOrgIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeptsByOrgIdRequest) GoString() string {
	return s.String()
}

func (s *GetDeptsByOrgIdRequest) SetMaxResults(v int32) *GetDeptsByOrgIdRequest {
	s.MaxResults = &v
	return s
}

func (s *GetDeptsByOrgIdRequest) SetNextToken(v int64) *GetDeptsByOrgIdRequest {
	s.NextToken = &v
	return s
}

type GetDeptsByOrgIdResponseBody struct {
	// deptList
	DeptList []*GetDeptsByOrgIdResponseBodyDeptList `json:"deptList,omitempty" xml:"deptList,omitempty" type:"Repeated"`
	// hasMore
	HasMore    *bool  `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// nextCursor
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetDeptsByOrgIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeptsByOrgIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeptsByOrgIdResponseBody) SetDeptList(v []*GetDeptsByOrgIdResponseBodyDeptList) *GetDeptsByOrgIdResponseBody {
	s.DeptList = v
	return s
}

func (s *GetDeptsByOrgIdResponseBody) SetHasMore(v bool) *GetDeptsByOrgIdResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetDeptsByOrgIdResponseBody) SetMaxResults(v int32) *GetDeptsByOrgIdResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetDeptsByOrgIdResponseBody) SetNextToken(v int64) *GetDeptsByOrgIdResponseBody {
	s.NextToken = &v
	return s
}

type GetDeptsByOrgIdResponseBodyDeptList struct {
	// id
	DeptId *int64 `json:"dept_id,omitempty" xml:"dept_id,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// parentId
	ParentId *int64 `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
}

func (s GetDeptsByOrgIdResponseBodyDeptList) String() string {
	return tea.Prettify(s)
}

func (s GetDeptsByOrgIdResponseBodyDeptList) GoString() string {
	return s.String()
}

func (s *GetDeptsByOrgIdResponseBodyDeptList) SetDeptId(v int64) *GetDeptsByOrgIdResponseBodyDeptList {
	s.DeptId = &v
	return s
}

func (s *GetDeptsByOrgIdResponseBodyDeptList) SetName(v string) *GetDeptsByOrgIdResponseBodyDeptList {
	s.Name = &v
	return s
}

func (s *GetDeptsByOrgIdResponseBodyDeptList) SetParentId(v int64) *GetDeptsByOrgIdResponseBodyDeptList {
	s.ParentId = &v
	return s
}

type GetDeptsByOrgIdResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDeptsByOrgIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeptsByOrgIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeptsByOrgIdResponse) GoString() string {
	return s.String()
}

func (s *GetDeptsByOrgIdResponse) SetHeaders(v map[string]*string) *GetDeptsByOrgIdResponse {
	s.Headers = v
	return s
}

func (s *GetDeptsByOrgIdResponse) SetBody(v *GetDeptsByOrgIdResponseBody) *GetDeptsByOrgIdResponse {
	s.Body = v
	return s
}

type GetEmpsByOrgIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	DingAccessTokenType     *string            `json:"dingAccessTokenType,omitempty" xml:"dingAccessTokenType,omitempty"`
	DingOrgId               *string            `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetEmpsByOrgIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetEmpsByOrgIdHeaders) GoString() string {
	return s.String()
}

func (s *GetEmpsByOrgIdHeaders) SetCommonHeaders(v map[string]*string) *GetEmpsByOrgIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetEmpsByOrgIdHeaders) SetDingAccessTokenType(v string) *GetEmpsByOrgIdHeaders {
	s.DingAccessTokenType = &v
	return s
}

func (s *GetEmpsByOrgIdHeaders) SetDingOrgId(v string) *GetEmpsByOrgIdHeaders {
	s.DingOrgId = &v
	return s
}

func (s *GetEmpsByOrgIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetEmpsByOrgIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetEmpsByOrgIdRequest struct {
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NeedDept   *bool  `json:"needDept,omitempty" xml:"needDept,omitempty"`
	NextToken  *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetEmpsByOrgIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEmpsByOrgIdRequest) GoString() string {
	return s.String()
}

func (s *GetEmpsByOrgIdRequest) SetMaxResults(v int32) *GetEmpsByOrgIdRequest {
	s.MaxResults = &v
	return s
}

func (s *GetEmpsByOrgIdRequest) SetNeedDept(v bool) *GetEmpsByOrgIdRequest {
	s.NeedDept = &v
	return s
}

func (s *GetEmpsByOrgIdRequest) SetNextToken(v int64) *GetEmpsByOrgIdRequest {
	s.NextToken = &v
	return s
}

type GetEmpsByOrgIdResponseBody struct {
	// empList
	EmpList []*GetEmpsByOrgIdResponseBodyEmpList `json:"empList,omitempty" xml:"empList,omitempty" type:"Repeated"`
	// hasMore
	HasMore   *bool  `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetEmpsByOrgIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEmpsByOrgIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetEmpsByOrgIdResponseBody) SetEmpList(v []*GetEmpsByOrgIdResponseBodyEmpList) *GetEmpsByOrgIdResponseBody {
	s.EmpList = v
	return s
}

func (s *GetEmpsByOrgIdResponseBody) SetHasMore(v bool) *GetEmpsByOrgIdResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetEmpsByOrgIdResponseBody) SetNextToken(v int64) *GetEmpsByOrgIdResponseBody {
	s.NextToken = &v
	return s
}

type GetEmpsByOrgIdResponseBodyEmpList struct {
	// avatar
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// deptIdList
	DeptIdList []*int64 `json:"dept_id_list,omitempty" xml:"dept_id_list,omitempty" type:"Repeated"`
	// dingId
	DingId *string `json:"dingId,omitempty" xml:"dingId,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// nick
	Nick *string `json:"nick,omitempty" xml:"nick,omitempty"`
	// orgId
	OrgId    *int64  `json:"orgId,omitempty" xml:"orgId,omitempty"`
	Position *string `json:"position,omitempty" xml:"position,omitempty"`
	// unionId
	Unionid *string `json:"unionid,omitempty" xml:"unionid,omitempty"`
	// userid
	Userid *string `json:"userid,omitempty" xml:"userid,omitempty"`
}

func (s GetEmpsByOrgIdResponseBodyEmpList) String() string {
	return tea.Prettify(s)
}

func (s GetEmpsByOrgIdResponseBodyEmpList) GoString() string {
	return s.String()
}

func (s *GetEmpsByOrgIdResponseBodyEmpList) SetAvatar(v string) *GetEmpsByOrgIdResponseBodyEmpList {
	s.Avatar = &v
	return s
}

func (s *GetEmpsByOrgIdResponseBodyEmpList) SetDeptIdList(v []*int64) *GetEmpsByOrgIdResponseBodyEmpList {
	s.DeptIdList = v
	return s
}

func (s *GetEmpsByOrgIdResponseBodyEmpList) SetDingId(v string) *GetEmpsByOrgIdResponseBodyEmpList {
	s.DingId = &v
	return s
}

func (s *GetEmpsByOrgIdResponseBodyEmpList) SetName(v string) *GetEmpsByOrgIdResponseBodyEmpList {
	s.Name = &v
	return s
}

func (s *GetEmpsByOrgIdResponseBodyEmpList) SetNick(v string) *GetEmpsByOrgIdResponseBodyEmpList {
	s.Nick = &v
	return s
}

func (s *GetEmpsByOrgIdResponseBodyEmpList) SetOrgId(v int64) *GetEmpsByOrgIdResponseBodyEmpList {
	s.OrgId = &v
	return s
}

func (s *GetEmpsByOrgIdResponseBodyEmpList) SetPosition(v string) *GetEmpsByOrgIdResponseBodyEmpList {
	s.Position = &v
	return s
}

func (s *GetEmpsByOrgIdResponseBodyEmpList) SetUnionid(v string) *GetEmpsByOrgIdResponseBodyEmpList {
	s.Unionid = &v
	return s
}

func (s *GetEmpsByOrgIdResponseBodyEmpList) SetUserid(v string) *GetEmpsByOrgIdResponseBodyEmpList {
	s.Userid = &v
	return s
}

type GetEmpsByOrgIdResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEmpsByOrgIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEmpsByOrgIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEmpsByOrgIdResponse) GoString() string {
	return s.String()
}

func (s *GetEmpsByOrgIdResponse) SetHeaders(v map[string]*string) *GetEmpsByOrgIdResponse {
	s.Headers = v
	return s
}

func (s *GetEmpsByOrgIdResponse) SetBody(v *GetEmpsByOrgIdResponseBody) *GetEmpsByOrgIdResponse {
	s.Body = v
	return s
}

type GetOrganizatioTaskByIdsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOrganizatioTaskByIdsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizatioTaskByIdsHeaders) GoString() string {
	return s.String()
}

func (s *GetOrganizatioTaskByIdsHeaders) SetCommonHeaders(v map[string]*string) *GetOrganizatioTaskByIdsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOrganizatioTaskByIdsHeaders) SetXAcsDingtalkAccessToken(v string) *GetOrganizatioTaskByIdsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOrganizatioTaskByIdsRequest struct {
	// 多个任务id
	TaskIds *string `json:"taskIds,omitempty" xml:"taskIds,omitempty"`
}

func (s GetOrganizatioTaskByIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizatioTaskByIdsRequest) GoString() string {
	return s.String()
}

func (s *GetOrganizatioTaskByIdsRequest) SetTaskIds(v string) *GetOrganizatioTaskByIdsRequest {
	s.TaskIds = &v
	return s
}

type GetOrganizatioTaskByIdsResponseBody struct {
	// 返回结构体
	Result []*GetOrganizatioTaskByIdsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetOrganizatioTaskByIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizatioTaskByIdsResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrganizatioTaskByIdsResponseBody) SetResult(v []*GetOrganizatioTaskByIdsResponseBodyResult) *GetOrganizatioTaskByIdsResponseBody {
	s.Result = v
	return s
}

type GetOrganizatioTaskByIdsResponseBodyResult struct {
	// 父任务id
	AncestorIds []*string `json:"ancestorIds,omitempty" xml:"ancestorIds,omitempty" type:"Repeated"`
	// 任务标题
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 创建时间
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// 创建者id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 任务截止时间
	DueDate *string `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	// 执行者id
	ExecutorId *string `json:"executorId,omitempty" xml:"executorId,omitempty"`
	// 参与者列表
	InvolveMembers []*string `json:"involveMembers,omitempty" xml:"involveMembers,omitempty" type:"Repeated"`
	// 任务是否已删除
	IsDeleted *bool `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	// 任务是否已完成
	IsDone *bool `json:"isDone,omitempty" xml:"isDone,omitempty"`
	// 任务自定义标记
	Labels []*string `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// 任务备注
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
	// 优先级【-10,0,1,2】中选一个
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// 任务开始时间
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// 任务id
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// 更新时间
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
	// 任务可见性。involves：仅参与者可见。members:所有人可见
	Visible *string `json:"visible,omitempty" xml:"visible,omitempty"`
}

func (s GetOrganizatioTaskByIdsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizatioTaskByIdsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetOrganizatioTaskByIdsResponseBodyResult) SetAncestorIds(v []*string) *GetOrganizatioTaskByIdsResponseBodyResult {
	s.AncestorIds = v
	return s
}

func (s *GetOrganizatioTaskByIdsResponseBodyResult) SetContent(v string) *GetOrganizatioTaskByIdsResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetOrganizatioTaskByIdsResponseBodyResult) SetCreated(v string) *GetOrganizatioTaskByIdsResponseBodyResult {
	s.Created = &v
	return s
}

func (s *GetOrganizatioTaskByIdsResponseBodyResult) SetCreatorId(v string) *GetOrganizatioTaskByIdsResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *GetOrganizatioTaskByIdsResponseBodyResult) SetDueDate(v string) *GetOrganizatioTaskByIdsResponseBodyResult {
	s.DueDate = &v
	return s
}

func (s *GetOrganizatioTaskByIdsResponseBodyResult) SetExecutorId(v string) *GetOrganizatioTaskByIdsResponseBodyResult {
	s.ExecutorId = &v
	return s
}

func (s *GetOrganizatioTaskByIdsResponseBodyResult) SetInvolveMembers(v []*string) *GetOrganizatioTaskByIdsResponseBodyResult {
	s.InvolveMembers = v
	return s
}

func (s *GetOrganizatioTaskByIdsResponseBodyResult) SetIsDeleted(v bool) *GetOrganizatioTaskByIdsResponseBodyResult {
	s.IsDeleted = &v
	return s
}

func (s *GetOrganizatioTaskByIdsResponseBodyResult) SetIsDone(v bool) *GetOrganizatioTaskByIdsResponseBodyResult {
	s.IsDone = &v
	return s
}

func (s *GetOrganizatioTaskByIdsResponseBodyResult) SetLabels(v []*string) *GetOrganizatioTaskByIdsResponseBodyResult {
	s.Labels = v
	return s
}

func (s *GetOrganizatioTaskByIdsResponseBodyResult) SetNote(v string) *GetOrganizatioTaskByIdsResponseBodyResult {
	s.Note = &v
	return s
}

func (s *GetOrganizatioTaskByIdsResponseBodyResult) SetPriority(v int32) *GetOrganizatioTaskByIdsResponseBodyResult {
	s.Priority = &v
	return s
}

func (s *GetOrganizatioTaskByIdsResponseBodyResult) SetStartDate(v string) *GetOrganizatioTaskByIdsResponseBodyResult {
	s.StartDate = &v
	return s
}

func (s *GetOrganizatioTaskByIdsResponseBodyResult) SetTaskId(v string) *GetOrganizatioTaskByIdsResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *GetOrganizatioTaskByIdsResponseBodyResult) SetUpdated(v string) *GetOrganizatioTaskByIdsResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *GetOrganizatioTaskByIdsResponseBodyResult) SetVisible(v string) *GetOrganizatioTaskByIdsResponseBodyResult {
	s.Visible = &v
	return s
}

type GetOrganizatioTaskByIdsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOrganizatioTaskByIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOrganizatioTaskByIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizatioTaskByIdsResponse) GoString() string {
	return s.String()
}

func (s *GetOrganizatioTaskByIdsResponse) SetHeaders(v map[string]*string) *GetOrganizatioTaskByIdsResponse {
	s.Headers = v
	return s
}

func (s *GetOrganizatioTaskByIdsResponse) SetBody(v *GetOrganizatioTaskByIdsResponseBody) *GetOrganizatioTaskByIdsResponse {
	s.Body = v
	return s
}

type GetOrganizationPriorityListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOrganizationPriorityListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationPriorityListHeaders) GoString() string {
	return s.String()
}

func (s *GetOrganizationPriorityListHeaders) SetCommonHeaders(v map[string]*string) *GetOrganizationPriorityListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOrganizationPriorityListHeaders) SetXAcsDingtalkAccessToken(v string) *GetOrganizationPriorityListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOrganizationPriorityListResponseBody struct {
	// 优先级列表
	Result []*GetOrganizationPriorityListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetOrganizationPriorityListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationPriorityListResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrganizationPriorityListResponseBody) SetResult(v []*GetOrganizationPriorityListResponseBodyResult) *GetOrganizationPriorityListResponseBody {
	s.Result = v
	return s
}

type GetOrganizationPriorityListResponseBodyResult struct {
	// 颜色
	Color *string `json:"color,omitempty" xml:"color,omitempty"`
	// 名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 优先级值
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// id
	PriorityId *string `json:"priorityId,omitempty" xml:"priorityId,omitempty"`
}

func (s GetOrganizationPriorityListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationPriorityListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetOrganizationPriorityListResponseBodyResult) SetColor(v string) *GetOrganizationPriorityListResponseBodyResult {
	s.Color = &v
	return s
}

func (s *GetOrganizationPriorityListResponseBodyResult) SetName(v string) *GetOrganizationPriorityListResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetOrganizationPriorityListResponseBodyResult) SetPriority(v string) *GetOrganizationPriorityListResponseBodyResult {
	s.Priority = &v
	return s
}

func (s *GetOrganizationPriorityListResponseBodyResult) SetPriorityId(v string) *GetOrganizationPriorityListResponseBodyResult {
	s.PriorityId = &v
	return s
}

type GetOrganizationPriorityListResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOrganizationPriorityListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOrganizationPriorityListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationPriorityListResponse) GoString() string {
	return s.String()
}

func (s *GetOrganizationPriorityListResponse) SetHeaders(v map[string]*string) *GetOrganizationPriorityListResponse {
	s.Headers = v
	return s
}

func (s *GetOrganizationPriorityListResponse) SetBody(v *GetOrganizationPriorityListResponseBody) *GetOrganizationPriorityListResponse {
	s.Body = v
	return s
}

type GetOrganizationTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOrganizationTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationTaskHeaders) GoString() string {
	return s.String()
}

func (s *GetOrganizationTaskHeaders) SetCommonHeaders(v map[string]*string) *GetOrganizationTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOrganizationTaskHeaders) SetXAcsDingtalkAccessToken(v string) *GetOrganizationTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOrganizationTaskResponseBody struct {
	// 返回结构体
	Result *GetOrganizationTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetOrganizationTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrganizationTaskResponseBody) SetResult(v *GetOrganizationTaskResponseBodyResult) *GetOrganizationTaskResponseBody {
	s.Result = v
	return s
}

type GetOrganizationTaskResponseBodyResult struct {
	// 父任务id
	AncestorIds []*string `json:"ancestorIds,omitempty" xml:"ancestorIds,omitempty" type:"Repeated"`
	// 任务标题
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 创建时间
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// 创建者id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 任务截止时间
	DueDate *string `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	// 执行者id
	ExecutorId *string `json:"executorId,omitempty" xml:"executorId,omitempty"`
	// 参与者列表
	InvolveMembers []*string `json:"involveMembers,omitempty" xml:"involveMembers,omitempty" type:"Repeated"`
	// 任务是否已删除
	IsDeleted *bool `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	// 任务是否已完成
	IsDone *bool `json:"isDone,omitempty" xml:"isDone,omitempty"`
	// 任务自定义标记
	Labels []*string `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// 任务备注
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
	// 优先级【-10,0,1,2】中选一个
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// 任务开始时间
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// 任务id
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// 更新时间
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
	// 任务可见性。involves：仅参与者可见。members:所有人可见
	Visible *string `json:"visible,omitempty" xml:"visible,omitempty"`
}

func (s GetOrganizationTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetOrganizationTaskResponseBodyResult) SetAncestorIds(v []*string) *GetOrganizationTaskResponseBodyResult {
	s.AncestorIds = v
	return s
}

func (s *GetOrganizationTaskResponseBodyResult) SetContent(v string) *GetOrganizationTaskResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetOrganizationTaskResponseBodyResult) SetCreated(v string) *GetOrganizationTaskResponseBodyResult {
	s.Created = &v
	return s
}

func (s *GetOrganizationTaskResponseBodyResult) SetCreatorId(v string) *GetOrganizationTaskResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *GetOrganizationTaskResponseBodyResult) SetDueDate(v string) *GetOrganizationTaskResponseBodyResult {
	s.DueDate = &v
	return s
}

func (s *GetOrganizationTaskResponseBodyResult) SetExecutorId(v string) *GetOrganizationTaskResponseBodyResult {
	s.ExecutorId = &v
	return s
}

func (s *GetOrganizationTaskResponseBodyResult) SetInvolveMembers(v []*string) *GetOrganizationTaskResponseBodyResult {
	s.InvolveMembers = v
	return s
}

func (s *GetOrganizationTaskResponseBodyResult) SetIsDeleted(v bool) *GetOrganizationTaskResponseBodyResult {
	s.IsDeleted = &v
	return s
}

func (s *GetOrganizationTaskResponseBodyResult) SetIsDone(v bool) *GetOrganizationTaskResponseBodyResult {
	s.IsDone = &v
	return s
}

func (s *GetOrganizationTaskResponseBodyResult) SetLabels(v []*string) *GetOrganizationTaskResponseBodyResult {
	s.Labels = v
	return s
}

func (s *GetOrganizationTaskResponseBodyResult) SetNote(v string) *GetOrganizationTaskResponseBodyResult {
	s.Note = &v
	return s
}

func (s *GetOrganizationTaskResponseBodyResult) SetPriority(v int32) *GetOrganizationTaskResponseBodyResult {
	s.Priority = &v
	return s
}

func (s *GetOrganizationTaskResponseBodyResult) SetStartDate(v string) *GetOrganizationTaskResponseBodyResult {
	s.StartDate = &v
	return s
}

func (s *GetOrganizationTaskResponseBodyResult) SetTaskId(v string) *GetOrganizationTaskResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *GetOrganizationTaskResponseBodyResult) SetUpdated(v string) *GetOrganizationTaskResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *GetOrganizationTaskResponseBodyResult) SetVisible(v string) *GetOrganizationTaskResponseBodyResult {
	s.Visible = &v
	return s
}

type GetOrganizationTaskResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOrganizationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOrganizationTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationTaskResponse) GoString() string {
	return s.String()
}

func (s *GetOrganizationTaskResponse) SetHeaders(v map[string]*string) *GetOrganizationTaskResponse {
	s.Headers = v
	return s
}

func (s *GetOrganizationTaskResponse) SetBody(v *GetOrganizationTaskResponseBody) *GetOrganizationTaskResponse {
	s.Body = v
	return s
}

type GetProjectGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetProjectGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetProjectGroupHeaders) GoString() string {
	return s.String()
}

func (s *GetProjectGroupHeaders) SetCommonHeaders(v map[string]*string) *GetProjectGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetProjectGroupHeaders) SetXAcsDingtalkAccessToken(v string) *GetProjectGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetProjectGroupRequest struct {
	// 分页大小，最小1，默认10，最大1000
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 查看者ID
	ViewerId *string `json:"viewerId,omitempty" xml:"viewerId,omitempty"`
}

func (s GetProjectGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectGroupRequest) GoString() string {
	return s.String()
}

func (s *GetProjectGroupRequest) SetPageSize(v int32) *GetProjectGroupRequest {
	s.PageSize = &v
	return s
}

func (s *GetProjectGroupRequest) SetViewerId(v string) *GetProjectGroupRequest {
	s.ViewerId = &v
	return s
}

type GetProjectGroupResponseBody struct {
	// 返回结果对象
	Result []*GetProjectGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetProjectGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectGroupResponseBody) SetResult(v []*GetProjectGroupResponseBodyResult) *GetProjectGroupResponseBody {
	s.Result = v
	return s
}

type GetProjectGroupResponseBodyResult struct {
	// 创建时间
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// 分组ID
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 分组名字
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 更新时间
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
	// 分组可见性。organization 或者 involves
	Visible *string `json:"visible,omitempty" xml:"visible,omitempty"`
}

func (s GetProjectGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetProjectGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetProjectGroupResponseBodyResult) SetCreated(v string) *GetProjectGroupResponseBodyResult {
	s.Created = &v
	return s
}

func (s *GetProjectGroupResponseBodyResult) SetId(v string) *GetProjectGroupResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetProjectGroupResponseBodyResult) SetName(v string) *GetProjectGroupResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetProjectGroupResponseBodyResult) SetUpdated(v string) *GetProjectGroupResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *GetProjectGroupResponseBodyResult) SetVisible(v string) *GetProjectGroupResponseBodyResult {
	s.Visible = &v
	return s
}

type GetProjectGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProjectGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProjectGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectGroupResponse) GoString() string {
	return s.String()
}

func (s *GetProjectGroupResponse) SetHeaders(v map[string]*string) *GetProjectGroupResponse {
	s.Headers = v
	return s
}

func (s *GetProjectGroupResponse) SetBody(v *GetProjectGroupResponseBody) *GetProjectGroupResponse {
	s.Body = v
	return s
}

type GetTbOrgIdByDingOrgIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTbOrgIdByDingOrgIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTbOrgIdByDingOrgIdHeaders) GoString() string {
	return s.String()
}

func (s *GetTbOrgIdByDingOrgIdHeaders) SetCommonHeaders(v map[string]*string) *GetTbOrgIdByDingOrgIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTbOrgIdByDingOrgIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetTbOrgIdByDingOrgIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTbOrgIdByDingOrgIdRequest struct {
	// 操作者userId
	OptUserId *string `json:"optUserId,omitempty" xml:"optUserId,omitempty"`
}

func (s GetTbOrgIdByDingOrgIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTbOrgIdByDingOrgIdRequest) GoString() string {
	return s.String()
}

func (s *GetTbOrgIdByDingOrgIdRequest) SetOptUserId(v string) *GetTbOrgIdByDingOrgIdRequest {
	s.OptUserId = &v
	return s
}

type GetTbOrgIdByDingOrgIdResponseBody struct {
	// 结果对象
	Result *GetTbOrgIdByDingOrgIdResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetTbOrgIdByDingOrgIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTbOrgIdByDingOrgIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetTbOrgIdByDingOrgIdResponseBody) SetResult(v *GetTbOrgIdByDingOrgIdResponseBodyResult) *GetTbOrgIdByDingOrgIdResponseBody {
	s.Result = v
	return s
}

type GetTbOrgIdByDingOrgIdResponseBodyResult struct {
	// Teambition企业Id
	TbOrganizationId *string `json:"tbOrganizationId,omitempty" xml:"tbOrganizationId,omitempty"`
}

func (s GetTbOrgIdByDingOrgIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetTbOrgIdByDingOrgIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetTbOrgIdByDingOrgIdResponseBodyResult) SetTbOrganizationId(v string) *GetTbOrgIdByDingOrgIdResponseBodyResult {
	s.TbOrganizationId = &v
	return s
}

type GetTbOrgIdByDingOrgIdResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTbOrgIdByDingOrgIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTbOrgIdByDingOrgIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTbOrgIdByDingOrgIdResponse) GoString() string {
	return s.String()
}

func (s *GetTbOrgIdByDingOrgIdResponse) SetHeaders(v map[string]*string) *GetTbOrgIdByDingOrgIdResponse {
	s.Headers = v
	return s
}

func (s *GetTbOrgIdByDingOrgIdResponse) SetBody(v *GetTbOrgIdByDingOrgIdResponseBody) *GetTbOrgIdByDingOrgIdResponse {
	s.Body = v
	return s
}

type GetTbProjectGrayHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	DingAccessTokenType     *string            `json:"dingAccessTokenType,omitempty" xml:"dingAccessTokenType,omitempty"`
	DingCorpId              *string            `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	DingIsvOrgId            *string            `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId               *string            `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingSuiteKey            *string            `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTbProjectGrayHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTbProjectGrayHeaders) GoString() string {
	return s.String()
}

func (s *GetTbProjectGrayHeaders) SetCommonHeaders(v map[string]*string) *GetTbProjectGrayHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTbProjectGrayHeaders) SetDingAccessTokenType(v string) *GetTbProjectGrayHeaders {
	s.DingAccessTokenType = &v
	return s
}

func (s *GetTbProjectGrayHeaders) SetDingCorpId(v string) *GetTbProjectGrayHeaders {
	s.DingCorpId = &v
	return s
}

func (s *GetTbProjectGrayHeaders) SetDingIsvOrgId(v string) *GetTbProjectGrayHeaders {
	s.DingIsvOrgId = &v
	return s
}

func (s *GetTbProjectGrayHeaders) SetDingOrgId(v string) *GetTbProjectGrayHeaders {
	s.DingOrgId = &v
	return s
}

func (s *GetTbProjectGrayHeaders) SetDingSuiteKey(v string) *GetTbProjectGrayHeaders {
	s.DingSuiteKey = &v
	return s
}

func (s *GetTbProjectGrayHeaders) SetXAcsDingtalkAccessToken(v string) *GetTbProjectGrayHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTbProjectGrayRequest struct {
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
}

func (s GetTbProjectGrayRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTbProjectGrayRequest) GoString() string {
	return s.String()
}

func (s *GetTbProjectGrayRequest) SetLabel(v string) *GetTbProjectGrayRequest {
	s.Label = &v
	return s
}

type GetTbProjectGrayResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 是否灰度
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetTbProjectGrayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTbProjectGrayResponseBody) GoString() string {
	return s.String()
}

func (s *GetTbProjectGrayResponseBody) SetRequestId(v string) *GetTbProjectGrayResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTbProjectGrayResponseBody) SetResult(v bool) *GetTbProjectGrayResponseBody {
	s.Result = &v
	return s
}

type GetTbProjectGrayResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTbProjectGrayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTbProjectGrayResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTbProjectGrayResponse) GoString() string {
	return s.String()
}

func (s *GetTbProjectGrayResponse) SetHeaders(v map[string]*string) *GetTbProjectGrayResponse {
	s.Headers = v
	return s
}

func (s *GetTbProjectGrayResponse) SetBody(v *GetTbProjectGrayResponseBody) *GetTbProjectGrayResponse {
	s.Body = v
	return s
}

type GetTbProjectSourceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	DingAccessTokenType     *string            `json:"dingAccessTokenType,omitempty" xml:"dingAccessTokenType,omitempty"`
	DingCorpId              *string            `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	DingIsvOrgId            *string            `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId               *string            `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingSuiteKey            *string            `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTbProjectSourceHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTbProjectSourceHeaders) GoString() string {
	return s.String()
}

func (s *GetTbProjectSourceHeaders) SetCommonHeaders(v map[string]*string) *GetTbProjectSourceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTbProjectSourceHeaders) SetDingAccessTokenType(v string) *GetTbProjectSourceHeaders {
	s.DingAccessTokenType = &v
	return s
}

func (s *GetTbProjectSourceHeaders) SetDingCorpId(v string) *GetTbProjectSourceHeaders {
	s.DingCorpId = &v
	return s
}

func (s *GetTbProjectSourceHeaders) SetDingIsvOrgId(v string) *GetTbProjectSourceHeaders {
	s.DingIsvOrgId = &v
	return s
}

func (s *GetTbProjectSourceHeaders) SetDingOrgId(v string) *GetTbProjectSourceHeaders {
	s.DingOrgId = &v
	return s
}

func (s *GetTbProjectSourceHeaders) SetDingSuiteKey(v string) *GetTbProjectSourceHeaders {
	s.DingSuiteKey = &v
	return s
}

func (s *GetTbProjectSourceHeaders) SetXAcsDingtalkAccessToken(v string) *GetTbProjectSourceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTbProjectSourceResponseBody struct {
	// 应用安装来源，"0"：来自应用中心，”6“：预安装
	InstallSource *string `json:"installSource,omitempty" xml:"installSource,omitempty"`
}

func (s GetTbProjectSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTbProjectSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetTbProjectSourceResponseBody) SetInstallSource(v string) *GetTbProjectSourceResponseBody {
	s.InstallSource = &v
	return s
}

type GetTbProjectSourceResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTbProjectSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTbProjectSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTbProjectSourceResponse) GoString() string {
	return s.String()
}

func (s *GetTbProjectSourceResponse) SetHeaders(v map[string]*string) *GetTbProjectSourceResponse {
	s.Headers = v
	return s
}

func (s *GetTbProjectSourceResponse) SetBody(v *GetTbProjectSourceResponseBody) *GetTbProjectSourceResponse {
	s.Body = v
	return s
}

type GetTbUserIdByStaffIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTbUserIdByStaffIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTbUserIdByStaffIdHeaders) GoString() string {
	return s.String()
}

func (s *GetTbUserIdByStaffIdHeaders) SetCommonHeaders(v map[string]*string) *GetTbUserIdByStaffIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTbUserIdByStaffIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetTbUserIdByStaffIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTbUserIdByStaffIdRequest struct {
	// 操作者userId
	OptUserId *string `json:"optUserId,omitempty" xml:"optUserId,omitempty"`
	// 用户userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetTbUserIdByStaffIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTbUserIdByStaffIdRequest) GoString() string {
	return s.String()
}

func (s *GetTbUserIdByStaffIdRequest) SetOptUserId(v string) *GetTbUserIdByStaffIdRequest {
	s.OptUserId = &v
	return s
}

func (s *GetTbUserIdByStaffIdRequest) SetUserId(v string) *GetTbUserIdByStaffIdRequest {
	s.UserId = &v
	return s
}

type GetTbUserIdByStaffIdResponseBody struct {
	// 结果对象
	Result *GetTbUserIdByStaffIdResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetTbUserIdByStaffIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTbUserIdByStaffIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetTbUserIdByStaffIdResponseBody) SetResult(v *GetTbUserIdByStaffIdResponseBodyResult) *GetTbUserIdByStaffIdResponseBody {
	s.Result = v
	return s
}

type GetTbUserIdByStaffIdResponseBodyResult struct {
	// Teambition用户id
	TbUserId *string `json:"tbUserId,omitempty" xml:"tbUserId,omitempty"`
}

func (s GetTbUserIdByStaffIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetTbUserIdByStaffIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetTbUserIdByStaffIdResponseBodyResult) SetTbUserId(v string) *GetTbUserIdByStaffIdResponseBodyResult {
	s.TbUserId = &v
	return s
}

type GetTbUserIdByStaffIdResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTbUserIdByStaffIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTbUserIdByStaffIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTbUserIdByStaffIdResponse) GoString() string {
	return s.String()
}

func (s *GetTbUserIdByStaffIdResponse) SetHeaders(v map[string]*string) *GetTbUserIdByStaffIdResponse {
	s.Headers = v
	return s
}

func (s *GetTbUserIdByStaffIdResponse) SetBody(v *GetTbUserIdByStaffIdResponseBody) *GetTbUserIdByStaffIdResponse {
	s.Body = v
	return s
}

type SearchProjectTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchProjectTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectTemplateHeaders) GoString() string {
	return s.String()
}

func (s *SearchProjectTemplateHeaders) SetCommonHeaders(v map[string]*string) *SearchProjectTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchProjectTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *SearchProjectTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchProjectTemplateRequest struct {
	// 项目模板名关键词
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
}

func (s SearchProjectTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectTemplateRequest) GoString() string {
	return s.String()
}

func (s *SearchProjectTemplateRequest) SetKeyword(v string) *SearchProjectTemplateRequest {
	s.Keyword = &v
	return s
}

type SearchProjectTemplateResponseBody struct {
	// 返回结果对象
	Result []*SearchProjectTemplateResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s SearchProjectTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *SearchProjectTemplateResponseBody) SetResult(v []*SearchProjectTemplateResponseBodyResult) *SearchProjectTemplateResponseBody {
	s.Result = v
	return s
}

type SearchProjectTemplateResponseBodyResult struct {
	// 创建时间
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// 模板描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 模板id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 是否已删除
	IsDeleted *bool `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	// 是否demo模板
	IsDemo *bool `json:"isDemo,omitempty" xml:"isDemo,omitempty"`
	// 模板log地址
	Logo *string `json:"logo,omitempty" xml:"logo,omitempty"`
	// 模板名字
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 更新时间
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
	// 模板可见性。organization 或者 involves
	Visible *string `json:"visible,omitempty" xml:"visible,omitempty"`
}

func (s SearchProjectTemplateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectTemplateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SearchProjectTemplateResponseBodyResult) SetCreated(v string) *SearchProjectTemplateResponseBodyResult {
	s.Created = &v
	return s
}

func (s *SearchProjectTemplateResponseBodyResult) SetDescription(v string) *SearchProjectTemplateResponseBodyResult {
	s.Description = &v
	return s
}

func (s *SearchProjectTemplateResponseBodyResult) SetId(v string) *SearchProjectTemplateResponseBodyResult {
	s.Id = &v
	return s
}

func (s *SearchProjectTemplateResponseBodyResult) SetIsDeleted(v bool) *SearchProjectTemplateResponseBodyResult {
	s.IsDeleted = &v
	return s
}

func (s *SearchProjectTemplateResponseBodyResult) SetIsDemo(v bool) *SearchProjectTemplateResponseBodyResult {
	s.IsDemo = &v
	return s
}

func (s *SearchProjectTemplateResponseBodyResult) SetLogo(v string) *SearchProjectTemplateResponseBodyResult {
	s.Logo = &v
	return s
}

func (s *SearchProjectTemplateResponseBodyResult) SetName(v string) *SearchProjectTemplateResponseBodyResult {
	s.Name = &v
	return s
}

func (s *SearchProjectTemplateResponseBodyResult) SetUpdated(v string) *SearchProjectTemplateResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *SearchProjectTemplateResponseBodyResult) SetVisible(v string) *SearchProjectTemplateResponseBodyResult {
	s.Visible = &v
	return s
}

type SearchProjectTemplateResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchProjectTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchProjectTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectTemplateResponse) GoString() string {
	return s.String()
}

func (s *SearchProjectTemplateResponse) SetHeaders(v map[string]*string) *SearchProjectTemplateResponse {
	s.Headers = v
	return s
}

func (s *SearchProjectTemplateResponse) SetBody(v *SearchProjectTemplateResponseBody) *SearchProjectTemplateResponse {
	s.Body = v
	return s
}

type UpdateCustomfieldValueHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateCustomfieldValueHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomfieldValueHeaders) GoString() string {
	return s.String()
}

func (s *UpdateCustomfieldValueHeaders) SetCommonHeaders(v map[string]*string) *UpdateCustomfieldValueHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateCustomfieldValueHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateCustomfieldValueHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateCustomfieldValueRequest struct {
	// 自定义字段id
	CustomfieldId *string `json:"customfieldId,omitempty" xml:"customfieldId,omitempty"`
	// 自定义字段名
	CustomfieldName *string `json:"customfieldName,omitempty" xml:"customfieldName,omitempty"`
	// 自定义字段值列表，最多10个
	Value []*UpdateCustomfieldValueRequestValue `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s UpdateCustomfieldValueRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomfieldValueRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomfieldValueRequest) SetCustomfieldId(v string) *UpdateCustomfieldValueRequest {
	s.CustomfieldId = &v
	return s
}

func (s *UpdateCustomfieldValueRequest) SetCustomfieldName(v string) *UpdateCustomfieldValueRequest {
	s.CustomfieldName = &v
	return s
}

func (s *UpdateCustomfieldValueRequest) SetValue(v []*UpdateCustomfieldValueRequestValue) *UpdateCustomfieldValueRequest {
	s.Value = v
	return s
}

type UpdateCustomfieldValueRequestValue struct {
	// 自定义字段显示值
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s UpdateCustomfieldValueRequestValue) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomfieldValueRequestValue) GoString() string {
	return s.String()
}

func (s *UpdateCustomfieldValueRequestValue) SetTitle(v string) *UpdateCustomfieldValueRequestValue {
	s.Title = &v
	return s
}

type UpdateCustomfieldValueResponseBody struct {
	// 返回对象
	Result *UpdateCustomfieldValueResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateCustomfieldValueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomfieldValueResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomfieldValueResponseBody) SetResult(v *UpdateCustomfieldValueResponseBodyResult) *UpdateCustomfieldValueResponseBody {
	s.Result = v
	return s
}

type UpdateCustomfieldValueResponseBodyResult struct {
	// 自定义字段数组
	Customfields []*UpdateCustomfieldValueResponseBodyResultCustomfields `json:"customfields,omitempty" xml:"customfields,omitempty" type:"Repeated"`
}

func (s UpdateCustomfieldValueResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomfieldValueResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateCustomfieldValueResponseBodyResult) SetCustomfields(v []*UpdateCustomfieldValueResponseBodyResultCustomfields) *UpdateCustomfieldValueResponseBodyResult {
	s.Customfields = v
	return s
}

type UpdateCustomfieldValueResponseBodyResultCustomfields struct {
	// 自定义字段id
	CustomfieldId *string `json:"customfieldId,omitempty" xml:"customfieldId,omitempty"`
	// 自定义字段值对象
	Value []*UpdateCustomfieldValueResponseBodyResultCustomfieldsValue `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s UpdateCustomfieldValueResponseBodyResultCustomfields) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomfieldValueResponseBodyResultCustomfields) GoString() string {
	return s.String()
}

func (s *UpdateCustomfieldValueResponseBodyResultCustomfields) SetCustomfieldId(v string) *UpdateCustomfieldValueResponseBodyResultCustomfields {
	s.CustomfieldId = &v
	return s
}

func (s *UpdateCustomfieldValueResponseBodyResultCustomfields) SetValue(v []*UpdateCustomfieldValueResponseBodyResultCustomfieldsValue) *UpdateCustomfieldValueResponseBodyResultCustomfields {
	s.Value = v
	return s
}

type UpdateCustomfieldValueResponseBodyResultCustomfieldsValue struct {
	// 自定义字段显示值
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s UpdateCustomfieldValueResponseBodyResultCustomfieldsValue) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomfieldValueResponseBodyResultCustomfieldsValue) GoString() string {
	return s.String()
}

func (s *UpdateCustomfieldValueResponseBodyResultCustomfieldsValue) SetTitle(v string) *UpdateCustomfieldValueResponseBodyResultCustomfieldsValue {
	s.Title = &v
	return s
}

type UpdateCustomfieldValueResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCustomfieldValueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCustomfieldValueResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomfieldValueResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomfieldValueResponse) SetHeaders(v map[string]*string) *UpdateCustomfieldValueResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomfieldValueResponse) SetBody(v *UpdateCustomfieldValueResponseBody) *UpdateCustomfieldValueResponse {
	s.Body = v
	return s
}

type UpdateOrganizationTaskContentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateOrganizationTaskContentHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskContentHeaders) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskContentHeaders) SetCommonHeaders(v map[string]*string) *UpdateOrganizationTaskContentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateOrganizationTaskContentHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateOrganizationTaskContentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateOrganizationTaskContentRequest struct {
	// 任务标题
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 是否禁止动态
	DisableActivity *bool `json:"disableActivity,omitempty" xml:"disableActivity,omitempty"`
	// 是否禁止通知
	DisableNotification *bool `json:"disableNotification,omitempty" xml:"disableNotification,omitempty"`
}

func (s UpdateOrganizationTaskContentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskContentRequest) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskContentRequest) SetContent(v string) *UpdateOrganizationTaskContentRequest {
	s.Content = &v
	return s
}

func (s *UpdateOrganizationTaskContentRequest) SetDisableActivity(v bool) *UpdateOrganizationTaskContentRequest {
	s.DisableActivity = &v
	return s
}

func (s *UpdateOrganizationTaskContentRequest) SetDisableNotification(v bool) *UpdateOrganizationTaskContentRequest {
	s.DisableNotification = &v
	return s
}

type UpdateOrganizationTaskContentResponseBody struct {
	// 返回对象
	Result *UpdateOrganizationTaskContentResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateOrganizationTaskContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskContentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskContentResponseBody) SetResult(v *UpdateOrganizationTaskContentResponseBodyResult) *UpdateOrganizationTaskContentResponseBody {
	s.Result = v
	return s
}

type UpdateOrganizationTaskContentResponseBodyResult struct {
	// 任务标题
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 更新时间
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateOrganizationTaskContentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskContentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskContentResponseBodyResult) SetContent(v string) *UpdateOrganizationTaskContentResponseBodyResult {
	s.Content = &v
	return s
}

func (s *UpdateOrganizationTaskContentResponseBodyResult) SetUpdated(v string) *UpdateOrganizationTaskContentResponseBodyResult {
	s.Updated = &v
	return s
}

type UpdateOrganizationTaskContentResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateOrganizationTaskContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateOrganizationTaskContentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskContentResponse) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskContentResponse) SetHeaders(v map[string]*string) *UpdateOrganizationTaskContentResponse {
	s.Headers = v
	return s
}

func (s *UpdateOrganizationTaskContentResponse) SetBody(v *UpdateOrganizationTaskContentResponseBody) *UpdateOrganizationTaskContentResponse {
	s.Body = v
	return s
}

type UpdateOrganizationTaskDueDateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateOrganizationTaskDueDateHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskDueDateHeaders) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskDueDateHeaders) SetCommonHeaders(v map[string]*string) *UpdateOrganizationTaskDueDateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateOrganizationTaskDueDateHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateOrganizationTaskDueDateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateOrganizationTaskDueDateRequest struct {
	// 是否禁止动态
	DisableActivity *bool `json:"disableActivity,omitempty" xml:"disableActivity,omitempty"`
	// 是否禁止通知
	DisableNotification *bool `json:"disableNotification,omitempty" xml:"disableNotification,omitempty"`
	// 任务截止时间
	DueDate *string `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
}

func (s UpdateOrganizationTaskDueDateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskDueDateRequest) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskDueDateRequest) SetDisableActivity(v bool) *UpdateOrganizationTaskDueDateRequest {
	s.DisableActivity = &v
	return s
}

func (s *UpdateOrganizationTaskDueDateRequest) SetDisableNotification(v bool) *UpdateOrganizationTaskDueDateRequest {
	s.DisableNotification = &v
	return s
}

func (s *UpdateOrganizationTaskDueDateRequest) SetDueDate(v string) *UpdateOrganizationTaskDueDateRequest {
	s.DueDate = &v
	return s
}

type UpdateOrganizationTaskDueDateResponseBody struct {
	// 返回对象
	Result *UpdateOrganizationTaskDueDateResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateOrganizationTaskDueDateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskDueDateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskDueDateResponseBody) SetResult(v *UpdateOrganizationTaskDueDateResponseBodyResult) *UpdateOrganizationTaskDueDateResponseBody {
	s.Result = v
	return s
}

type UpdateOrganizationTaskDueDateResponseBodyResult struct {
	// 任务截止时间
	DueDate *string `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	// 更新时间
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s UpdateOrganizationTaskDueDateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskDueDateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskDueDateResponseBodyResult) SetDueDate(v string) *UpdateOrganizationTaskDueDateResponseBodyResult {
	s.DueDate = &v
	return s
}

func (s *UpdateOrganizationTaskDueDateResponseBodyResult) SetUpdateTime(v string) *UpdateOrganizationTaskDueDateResponseBodyResult {
	s.UpdateTime = &v
	return s
}

type UpdateOrganizationTaskDueDateResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateOrganizationTaskDueDateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateOrganizationTaskDueDateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskDueDateResponse) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskDueDateResponse) SetHeaders(v map[string]*string) *UpdateOrganizationTaskDueDateResponse {
	s.Headers = v
	return s
}

func (s *UpdateOrganizationTaskDueDateResponse) SetBody(v *UpdateOrganizationTaskDueDateResponseBody) *UpdateOrganizationTaskDueDateResponse {
	s.Body = v
	return s
}

type UpdateOrganizationTaskExecutorHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateOrganizationTaskExecutorHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskExecutorHeaders) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskExecutorHeaders) SetCommonHeaders(v map[string]*string) *UpdateOrganizationTaskExecutorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateOrganizationTaskExecutorHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateOrganizationTaskExecutorHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateOrganizationTaskExecutorRequest struct {
	DisableActivity     *bool   `json:"disableActivity,omitempty" xml:"disableActivity,omitempty"`
	DisableNotification *bool   `json:"disableNotification,omitempty" xml:"disableNotification,omitempty"`
	ExecutorId          *string `json:"executorId,omitempty" xml:"executorId,omitempty"`
}

func (s UpdateOrganizationTaskExecutorRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskExecutorRequest) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskExecutorRequest) SetDisableActivity(v bool) *UpdateOrganizationTaskExecutorRequest {
	s.DisableActivity = &v
	return s
}

func (s *UpdateOrganizationTaskExecutorRequest) SetDisableNotification(v bool) *UpdateOrganizationTaskExecutorRequest {
	s.DisableNotification = &v
	return s
}

func (s *UpdateOrganizationTaskExecutorRequest) SetExecutorId(v string) *UpdateOrganizationTaskExecutorRequest {
	s.ExecutorId = &v
	return s
}

type UpdateOrganizationTaskExecutorResponseBody struct {
	// 返回对象
	Result *UpdateOrganizationTaskExecutorResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateOrganizationTaskExecutorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskExecutorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskExecutorResponseBody) SetResult(v *UpdateOrganizationTaskExecutorResponseBodyResult) *UpdateOrganizationTaskExecutorResponseBody {
	s.Result = v
	return s
}

type UpdateOrganizationTaskExecutorResponseBodyResult struct {
	// 执行者信息
	Executor *UpdateOrganizationTaskExecutorResponseBodyResultExecutor `json:"executor,omitempty" xml:"executor,omitempty" type:"Struct"`
	// 执行者id
	ExecutorId *string `json:"executorId,omitempty" xml:"executorId,omitempty"`
	// 参与者列表
	Involvers []*UpdateOrganizationTaskExecutorResponseBodyResultInvolvers `json:"involvers,omitempty" xml:"involvers,omitempty" type:"Repeated"`
	// 更新时间
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateOrganizationTaskExecutorResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskExecutorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskExecutorResponseBodyResult) SetExecutor(v *UpdateOrganizationTaskExecutorResponseBodyResultExecutor) *UpdateOrganizationTaskExecutorResponseBodyResult {
	s.Executor = v
	return s
}

func (s *UpdateOrganizationTaskExecutorResponseBodyResult) SetExecutorId(v string) *UpdateOrganizationTaskExecutorResponseBodyResult {
	s.ExecutorId = &v
	return s
}

func (s *UpdateOrganizationTaskExecutorResponseBodyResult) SetInvolvers(v []*UpdateOrganizationTaskExecutorResponseBodyResultInvolvers) *UpdateOrganizationTaskExecutorResponseBodyResult {
	s.Involvers = v
	return s
}

func (s *UpdateOrganizationTaskExecutorResponseBodyResult) SetUpdated(v string) *UpdateOrganizationTaskExecutorResponseBodyResult {
	s.Updated = &v
	return s
}

type UpdateOrganizationTaskExecutorResponseBodyResultExecutor struct {
	// 头像
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// 名字
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户uid
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateOrganizationTaskExecutorResponseBodyResultExecutor) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskExecutorResponseBodyResultExecutor) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskExecutorResponseBodyResultExecutor) SetAvatarUrl(v string) *UpdateOrganizationTaskExecutorResponseBodyResultExecutor {
	s.AvatarUrl = &v
	return s
}

func (s *UpdateOrganizationTaskExecutorResponseBodyResultExecutor) SetName(v string) *UpdateOrganizationTaskExecutorResponseBodyResultExecutor {
	s.Name = &v
	return s
}

func (s *UpdateOrganizationTaskExecutorResponseBodyResultExecutor) SetUserId(v string) *UpdateOrganizationTaskExecutorResponseBodyResultExecutor {
	s.UserId = &v
	return s
}

type UpdateOrganizationTaskExecutorResponseBodyResultInvolvers struct {
	// 头像
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// 名字
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户uid
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateOrganizationTaskExecutorResponseBodyResultInvolvers) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskExecutorResponseBodyResultInvolvers) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskExecutorResponseBodyResultInvolvers) SetAvatarUrl(v string) *UpdateOrganizationTaskExecutorResponseBodyResultInvolvers {
	s.AvatarUrl = &v
	return s
}

func (s *UpdateOrganizationTaskExecutorResponseBodyResultInvolvers) SetName(v string) *UpdateOrganizationTaskExecutorResponseBodyResultInvolvers {
	s.Name = &v
	return s
}

func (s *UpdateOrganizationTaskExecutorResponseBodyResultInvolvers) SetUserId(v string) *UpdateOrganizationTaskExecutorResponseBodyResultInvolvers {
	s.UserId = &v
	return s
}

type UpdateOrganizationTaskExecutorResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateOrganizationTaskExecutorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateOrganizationTaskExecutorResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskExecutorResponse) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskExecutorResponse) SetHeaders(v map[string]*string) *UpdateOrganizationTaskExecutorResponse {
	s.Headers = v
	return s
}

func (s *UpdateOrganizationTaskExecutorResponse) SetBody(v *UpdateOrganizationTaskExecutorResponseBody) *UpdateOrganizationTaskExecutorResponse {
	s.Body = v
	return s
}

type UpdateOrganizationTaskInvolveMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateOrganizationTaskInvolveMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskInvolveMembersHeaders) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskInvolveMembersHeaders) SetCommonHeaders(v map[string]*string) *UpdateOrganizationTaskInvolveMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateOrganizationTaskInvolveMembersHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateOrganizationTaskInvolveMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateOrganizationTaskInvolveMembersRequest struct {
	// 增加的参与者uid
	AddInvolvers []*string `json:"addInvolvers,omitempty" xml:"addInvolvers,omitempty" type:"Repeated"`
	// 删除的参与者uid
	DelInvolvers []*string `json:"delInvolvers,omitempty" xml:"delInvolvers,omitempty" type:"Repeated"`
	// 是否禁用动态
	DisableActivity *bool `json:"disableActivity,omitempty" xml:"disableActivity,omitempty"`
	// 是否禁用通知
	DisableNotification *bool `json:"disableNotification,omitempty" xml:"disableNotification,omitempty"`
	// 所有参与者uid
	InvolveMembers []*string `json:"involveMembers,omitempty" xml:"involveMembers,omitempty" type:"Repeated"`
}

func (s UpdateOrganizationTaskInvolveMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskInvolveMembersRequest) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskInvolveMembersRequest) SetAddInvolvers(v []*string) *UpdateOrganizationTaskInvolveMembersRequest {
	s.AddInvolvers = v
	return s
}

func (s *UpdateOrganizationTaskInvolveMembersRequest) SetDelInvolvers(v []*string) *UpdateOrganizationTaskInvolveMembersRequest {
	s.DelInvolvers = v
	return s
}

func (s *UpdateOrganizationTaskInvolveMembersRequest) SetDisableActivity(v bool) *UpdateOrganizationTaskInvolveMembersRequest {
	s.DisableActivity = &v
	return s
}

func (s *UpdateOrganizationTaskInvolveMembersRequest) SetDisableNotification(v bool) *UpdateOrganizationTaskInvolveMembersRequest {
	s.DisableNotification = &v
	return s
}

func (s *UpdateOrganizationTaskInvolveMembersRequest) SetInvolveMembers(v []*string) *UpdateOrganizationTaskInvolveMembersRequest {
	s.InvolveMembers = v
	return s
}

type UpdateOrganizationTaskInvolveMembersResponseBody struct {
	// 返回对象
	Result *UpdateOrganizationTaskInvolveMembersResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateOrganizationTaskInvolveMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskInvolveMembersResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskInvolveMembersResponseBody) SetResult(v *UpdateOrganizationTaskInvolveMembersResponseBodyResult) *UpdateOrganizationTaskInvolveMembersResponseBody {
	s.Result = v
	return s
}

type UpdateOrganizationTaskInvolveMembersResponseBodyResult struct {
	// 参与者列表
	Involvers []*UpdateOrganizationTaskInvolveMembersResponseBodyResultInvolvers `json:"involvers,omitempty" xml:"involvers,omitempty" type:"Repeated"`
	// 更新时间
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateOrganizationTaskInvolveMembersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskInvolveMembersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskInvolveMembersResponseBodyResult) SetInvolvers(v []*UpdateOrganizationTaskInvolveMembersResponseBodyResultInvolvers) *UpdateOrganizationTaskInvolveMembersResponseBodyResult {
	s.Involvers = v
	return s
}

func (s *UpdateOrganizationTaskInvolveMembersResponseBodyResult) SetUpdated(v string) *UpdateOrganizationTaskInvolveMembersResponseBodyResult {
	s.Updated = &v
	return s
}

type UpdateOrganizationTaskInvolveMembersResponseBodyResultInvolvers struct {
	// 头像
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// 名字
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户uid
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateOrganizationTaskInvolveMembersResponseBodyResultInvolvers) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskInvolveMembersResponseBodyResultInvolvers) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskInvolveMembersResponseBodyResultInvolvers) SetAvatarUrl(v string) *UpdateOrganizationTaskInvolveMembersResponseBodyResultInvolvers {
	s.AvatarUrl = &v
	return s
}

func (s *UpdateOrganizationTaskInvolveMembersResponseBodyResultInvolvers) SetName(v string) *UpdateOrganizationTaskInvolveMembersResponseBodyResultInvolvers {
	s.Name = &v
	return s
}

func (s *UpdateOrganizationTaskInvolveMembersResponseBodyResultInvolvers) SetUserId(v string) *UpdateOrganizationTaskInvolveMembersResponseBodyResultInvolvers {
	s.UserId = &v
	return s
}

type UpdateOrganizationTaskInvolveMembersResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateOrganizationTaskInvolveMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateOrganizationTaskInvolveMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskInvolveMembersResponse) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskInvolveMembersResponse) SetHeaders(v map[string]*string) *UpdateOrganizationTaskInvolveMembersResponse {
	s.Headers = v
	return s
}

func (s *UpdateOrganizationTaskInvolveMembersResponse) SetBody(v *UpdateOrganizationTaskInvolveMembersResponseBody) *UpdateOrganizationTaskInvolveMembersResponse {
	s.Body = v
	return s
}

type UpdateOrganizationTaskNoteHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateOrganizationTaskNoteHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskNoteHeaders) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskNoteHeaders) SetCommonHeaders(v map[string]*string) *UpdateOrganizationTaskNoteHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateOrganizationTaskNoteHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateOrganizationTaskNoteHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateOrganizationTaskNoteRequest struct {
	// 是否禁用动态
	DisableActivity *bool `json:"disableActivity,omitempty" xml:"disableActivity,omitempty"`
	// 是否禁用通知
	DisableNotification *bool `json:"disableNotification,omitempty" xml:"disableNotification,omitempty"`
	// 任务备注
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
}

func (s UpdateOrganizationTaskNoteRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskNoteRequest) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskNoteRequest) SetDisableActivity(v bool) *UpdateOrganizationTaskNoteRequest {
	s.DisableActivity = &v
	return s
}

func (s *UpdateOrganizationTaskNoteRequest) SetDisableNotification(v bool) *UpdateOrganizationTaskNoteRequest {
	s.DisableNotification = &v
	return s
}

func (s *UpdateOrganizationTaskNoteRequest) SetNote(v string) *UpdateOrganizationTaskNoteRequest {
	s.Note = &v
	return s
}

type UpdateOrganizationTaskNoteResponseBody struct {
	Result *UpdateOrganizationTaskNoteResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateOrganizationTaskNoteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskNoteResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskNoteResponseBody) SetResult(v *UpdateOrganizationTaskNoteResponseBodyResult) *UpdateOrganizationTaskNoteResponseBody {
	s.Result = v
	return s
}

type UpdateOrganizationTaskNoteResponseBodyResult struct {
	// 任务备注
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
	// 更新时间
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateOrganizationTaskNoteResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskNoteResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskNoteResponseBodyResult) SetNote(v string) *UpdateOrganizationTaskNoteResponseBodyResult {
	s.Note = &v
	return s
}

func (s *UpdateOrganizationTaskNoteResponseBodyResult) SetUpdated(v string) *UpdateOrganizationTaskNoteResponseBodyResult {
	s.Updated = &v
	return s
}

type UpdateOrganizationTaskNoteResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateOrganizationTaskNoteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateOrganizationTaskNoteResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskNoteResponse) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskNoteResponse) SetHeaders(v map[string]*string) *UpdateOrganizationTaskNoteResponse {
	s.Headers = v
	return s
}

func (s *UpdateOrganizationTaskNoteResponse) SetBody(v *UpdateOrganizationTaskNoteResponseBody) *UpdateOrganizationTaskNoteResponse {
	s.Body = v
	return s
}

type UpdateOrganizationTaskPriorityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateOrganizationTaskPriorityHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskPriorityHeaders) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskPriorityHeaders) SetCommonHeaders(v map[string]*string) *UpdateOrganizationTaskPriorityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateOrganizationTaskPriorityHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateOrganizationTaskPriorityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateOrganizationTaskPriorityRequest struct {
	// 是否禁止动态
	DisableActivity *bool `json:"disableActivity,omitempty" xml:"disableActivity,omitempty"`
	// 是否禁止通知
	DisableNotification *bool `json:"disableNotification,omitempty" xml:"disableNotification,omitempty"`
	// 优先级【-10,0,1,2】中的一个值
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
}

func (s UpdateOrganizationTaskPriorityRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskPriorityRequest) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskPriorityRequest) SetDisableActivity(v bool) *UpdateOrganizationTaskPriorityRequest {
	s.DisableActivity = &v
	return s
}

func (s *UpdateOrganizationTaskPriorityRequest) SetDisableNotification(v bool) *UpdateOrganizationTaskPriorityRequest {
	s.DisableNotification = &v
	return s
}

func (s *UpdateOrganizationTaskPriorityRequest) SetPriority(v int32) *UpdateOrganizationTaskPriorityRequest {
	s.Priority = &v
	return s
}

type UpdateOrganizationTaskPriorityResponseBody struct {
	// 返回对象
	Result *UpdateOrganizationTaskPriorityResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateOrganizationTaskPriorityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskPriorityResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskPriorityResponseBody) SetResult(v *UpdateOrganizationTaskPriorityResponseBodyResult) *UpdateOrganizationTaskPriorityResponseBody {
	s.Result = v
	return s
}

type UpdateOrganizationTaskPriorityResponseBodyResult struct {
	// 优先级【-10,0,1,2】中的一个
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// 更新时间
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateOrganizationTaskPriorityResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskPriorityResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskPriorityResponseBodyResult) SetPriority(v int32) *UpdateOrganizationTaskPriorityResponseBodyResult {
	s.Priority = &v
	return s
}

func (s *UpdateOrganizationTaskPriorityResponseBodyResult) SetUpdated(v string) *UpdateOrganizationTaskPriorityResponseBodyResult {
	s.Updated = &v
	return s
}

type UpdateOrganizationTaskPriorityResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateOrganizationTaskPriorityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateOrganizationTaskPriorityResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskPriorityResponse) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskPriorityResponse) SetHeaders(v map[string]*string) *UpdateOrganizationTaskPriorityResponse {
	s.Headers = v
	return s
}

func (s *UpdateOrganizationTaskPriorityResponse) SetBody(v *UpdateOrganizationTaskPriorityResponseBody) *UpdateOrganizationTaskPriorityResponse {
	s.Body = v
	return s
}

type UpdateOrganizationTaskStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateOrganizationTaskStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskStatusHeaders) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskStatusHeaders) SetCommonHeaders(v map[string]*string) *UpdateOrganizationTaskStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateOrganizationTaskStatusHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateOrganizationTaskStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateOrganizationTaskStatusRequest struct {
	// 是否禁用动态
	DisableActivity *bool `json:"disableActivity,omitempty" xml:"disableActivity,omitempty"`
	// 是否禁用通知
	DisableNotification *bool `json:"disableNotification,omitempty" xml:"disableNotification,omitempty"`
	// true改成完成，false 改成未完成
	IsDone *bool `json:"isDone,omitempty" xml:"isDone,omitempty"`
}

func (s UpdateOrganizationTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskStatusRequest) SetDisableActivity(v bool) *UpdateOrganizationTaskStatusRequest {
	s.DisableActivity = &v
	return s
}

func (s *UpdateOrganizationTaskStatusRequest) SetDisableNotification(v bool) *UpdateOrganizationTaskStatusRequest {
	s.DisableNotification = &v
	return s
}

func (s *UpdateOrganizationTaskStatusRequest) SetIsDone(v bool) *UpdateOrganizationTaskStatusRequest {
	s.IsDone = &v
	return s
}

type UpdateOrganizationTaskStatusResponseBody struct {
	// 返回对象
	Result *UpdateOrganizationTaskStatusResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateOrganizationTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskStatusResponseBody) SetResult(v *UpdateOrganizationTaskStatusResponseBodyResult) *UpdateOrganizationTaskStatusResponseBody {
	s.Result = v
	return s
}

type UpdateOrganizationTaskStatusResponseBodyResult struct {
	// 是否已完成
	IsDone *bool `json:"isDone,omitempty" xml:"isDone,omitempty"`
	// 更新时间
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s UpdateOrganizationTaskStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskStatusResponseBodyResult) SetIsDone(v bool) *UpdateOrganizationTaskStatusResponseBodyResult {
	s.IsDone = &v
	return s
}

func (s *UpdateOrganizationTaskStatusResponseBodyResult) SetUpdateTime(v string) *UpdateOrganizationTaskStatusResponseBodyResult {
	s.UpdateTime = &v
	return s
}

type UpdateOrganizationTaskStatusResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateOrganizationTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateOrganizationTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationTaskStatusResponse) SetHeaders(v map[string]*string) *UpdateOrganizationTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateOrganizationTaskStatusResponse) SetBody(v *UpdateOrganizationTaskStatusResponseBody) *UpdateOrganizationTaskStatusResponse {
	s.Body = v
	return s
}

type UpdateProjectGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateProjectGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectGroupHeaders) GoString() string {
	return s.String()
}

func (s *UpdateProjectGroupHeaders) SetCommonHeaders(v map[string]*string) *UpdateProjectGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateProjectGroupHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateProjectGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateProjectGroupRequest struct {
	// 增加到项目分组的Id列表，最多5个
	AddProjectGroupIds []*string `json:"addProjectGroupIds,omitempty" xml:"addProjectGroupIds,omitempty" type:"Repeated"`
	// 移除项目分组的Id列表，最多5个
	DelProjectGroupIds []*string `json:"delProjectGroupIds,omitempty" xml:"delProjectGroupIds,omitempty" type:"Repeated"`
}

func (s UpdateProjectGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectGroupRequest) SetAddProjectGroupIds(v []*string) *UpdateProjectGroupRequest {
	s.AddProjectGroupIds = v
	return s
}

func (s *UpdateProjectGroupRequest) SetDelProjectGroupIds(v []*string) *UpdateProjectGroupRequest {
	s.DelProjectGroupIds = v
	return s
}

type UpdateProjectGroupResponseBody struct {
	// 结果对象
	Result *UpdateProjectGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateProjectGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectGroupResponseBody) SetResult(v *UpdateProjectGroupResponseBodyResult) *UpdateProjectGroupResponseBody {
	s.Result = v
	return s
}

type UpdateProjectGroupResponseBodyResult struct {
	// 是否成功
	Ok *bool `json:"ok,omitempty" xml:"ok,omitempty"`
}

func (s UpdateProjectGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateProjectGroupResponseBodyResult) SetOk(v bool) *UpdateProjectGroupResponseBodyResult {
	s.Ok = &v
	return s
}

type UpdateProjectGroupResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateProjectGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProjectGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectGroupResponse) SetHeaders(v map[string]*string) *UpdateProjectGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectGroupResponse) SetBody(v *UpdateProjectGroupResponseBody) *UpdateProjectGroupResponse {
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

func (client *Client) AddProjectMember(userId *string, projectId *string, request *AddProjectMemberRequest) (_result *AddProjectMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddProjectMemberHeaders{}
	_result = &AddProjectMemberResponse{}
	_body, _err := client.AddProjectMemberWithOptions(userId, projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddProjectMemberWithOptions(userId *string, projectId *string, request *AddProjectMemberRequest, headers *AddProjectMemberHeaders, runtime *util.RuntimeOptions) (_result *AddProjectMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	projectId = openapiutil.GetEncodeParam(projectId)
	body := map[string]interface{}{}
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
	_result = &AddProjectMemberResponse{}
	_body, _err := client.DoROARequest(tea.String("AddProjectMember"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/project/users/"+tea.StringValue(userId)+"/projects/"+tea.StringValue(projectId)+"/members"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOrganizationTask(userId *string, request *CreateOrganizationTaskRequest) (_result *CreateOrganizationTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateOrganizationTaskHeaders{}
	_result = &CreateOrganizationTaskResponse{}
	_body, _err := client.CreateOrganizationTaskWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrganizationTaskWithOptions(userId *string, request *CreateOrganizationTaskRequest, headers *CreateOrganizationTaskHeaders, runtime *util.RuntimeOptions) (_result *CreateOrganizationTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTime)) {
		body["createTime"] = request.CreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.DisableActivity)) {
		body["disableActivity"] = request.DisableActivity
	}

	if !tea.BoolValue(util.IsUnset(request.DisableNotification)) {
		body["disableNotification"] = request.DisableNotification
	}

	if !tea.BoolValue(util.IsUnset(request.DueDate)) {
		body["dueDate"] = request.DueDate
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorId)) {
		body["executorId"] = request.ExecutorId
	}

	if !tea.BoolValue(util.IsUnset(request.InvolveMembers)) {
		body["involveMembers"] = request.InvolveMembers
	}

	if !tea.BoolValue(util.IsUnset(request.Note)) {
		body["note"] = request.Note
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.Visible)) {
		body["visible"] = request.Visible
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
	_result = &CreateOrganizationTaskResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateOrganizationTask"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/project/organizations/users/"+tea.StringValue(userId)+"/tasks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePlanTime(userId *string, request *CreatePlanTimeRequest) (_result *CreatePlanTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreatePlanTimeHeaders{}
	_result = &CreatePlanTimeResponse{}
	_body, _err := client.CreatePlanTimeWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePlanTimeWithOptions(userId *string, request *CreatePlanTimeRequest, headers *CreatePlanTimeHeaders, runtime *util.RuntimeOptions) (_result *CreatePlanTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TenantType)) {
		query["tenantType"] = request.TenantType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["endDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorId)) {
		body["executorId"] = request.ExecutorId
	}

	if !tea.BoolValue(util.IsUnset(request.IncludesHolidays)) {
		body["includesHolidays"] = request.IncludesHolidays
	}

	if !tea.BoolValue(util.IsUnset(request.IsDuration)) {
		body["isDuration"] = request.IsDuration
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectId)) {
		body["objectId"] = request.ObjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectType)) {
		body["objectType"] = request.ObjectType
	}

	if !tea.BoolValue(util.IsUnset(request.PlanTime)) {
		body["planTime"] = request.PlanTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["startDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.SubmitterId)) {
		body["submitterId"] = request.SubmitterId
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
	_result = &CreatePlanTimeResponse{}
	_body, _err := client.DoROARequest(tea.String("CreatePlanTime"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/project/users/"+tea.StringValue(userId)+"/planTimes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProjectByTemplate(userId *string, request *CreateProjectByTemplateRequest) (_result *CreateProjectByTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateProjectByTemplateHeaders{}
	_result = &CreateProjectByTemplateResponse{}
	_body, _err := client.CreateProjectByTemplateWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProjectByTemplateWithOptions(userId *string, request *CreateProjectByTemplateRequest, headers *CreateProjectByTemplateHeaders, runtime *util.RuntimeOptions) (_result *CreateProjectByTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
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
	_result = &CreateProjectByTemplateResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateProjectByTemplate"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/project/users/"+tea.StringValue(userId)+"/templates/projects"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTask(userId *string, request *CreateTaskRequest) (_result *CreateTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTaskHeaders{}
	_result = &CreateTaskResponse{}
	_body, _err := client.CreateTaskWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTaskWithOptions(userId *string, request *CreateTaskRequest, headers *CreateTaskHeaders, runtime *util.RuntimeOptions) (_result *CreateTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Customfields)) {
		body["customfields"] = request.Customfields
	}

	if !tea.BoolValue(util.IsUnset(request.DueDate)) {
		body["dueDate"] = request.DueDate
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorId)) {
		body["executorId"] = request.ExecutorId
	}

	if !tea.BoolValue(util.IsUnset(request.Note)) {
		body["note"] = request.Note
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["projectId"] = request.ProjectId
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
	_result = &CreateTaskResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateTask"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/project/users/"+tea.StringValue(userId)+"/tasks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTaskObjectLink(userId *string, taskId *string, request *CreateTaskObjectLinkRequest) (_result *CreateTaskObjectLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTaskObjectLinkHeaders{}
	_result = &CreateTaskObjectLinkResponse{}
	_body, _err := client.CreateTaskObjectLinkWithOptions(userId, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTaskObjectLinkWithOptions(userId *string, taskId *string, request *CreateTaskObjectLinkRequest, headers *CreateTaskObjectLinkHeaders, runtime *util.RuntimeOptions) (_result *CreateTaskObjectLinkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	taskId = openapiutil.GetEncodeParam(taskId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.LinkedData))) {
		body["linkedData"] = request.LinkedData
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
	_result = &CreateTaskObjectLinkResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateTaskObjectLink"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/project/users/"+tea.StringValue(userId)+"/tasks/"+tea.StringValue(taskId)+"/objectLinks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWorkTime(userId *string, request *CreateWorkTimeRequest) (_result *CreateWorkTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateWorkTimeHeaders{}
	_result = &CreateWorkTimeResponse{}
	_body, _err := client.CreateWorkTimeWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWorkTimeWithOptions(userId *string, request *CreateWorkTimeRequest, headers *CreateWorkTimeHeaders, runtime *util.RuntimeOptions) (_result *CreateWorkTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TenantType)) {
		query["tenantType"] = request.TenantType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["endDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorId)) {
		body["executorId"] = request.ExecutorId
	}

	if !tea.BoolValue(util.IsUnset(request.IncludesHolidays)) {
		body["includesHolidays"] = request.IncludesHolidays
	}

	if !tea.BoolValue(util.IsUnset(request.IsDuration)) {
		body["isDuration"] = request.IsDuration
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectId)) {
		body["objectId"] = request.ObjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectType)) {
		body["objectType"] = request.ObjectType
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["startDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.SubmitterId)) {
		body["submitterId"] = request.SubmitterId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkTime)) {
		body["workTime"] = request.WorkTime
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
	_result = &CreateWorkTimeResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateWorkTime"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/project/users/"+tea.StringValue(userId)+"/workTimes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeptsByOrgId(request *GetDeptsByOrgIdRequest) (_result *GetDeptsByOrgIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDeptsByOrgIdHeaders{}
	_result = &GetDeptsByOrgIdResponse{}
	_body, _err := client.GetDeptsByOrgIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeptsByOrgIdWithOptions(request *GetDeptsByOrgIdRequest, headers *GetDeptsByOrgIdHeaders, runtime *util.RuntimeOptions) (_result *GetDeptsByOrgIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.DingAccessTokenType)) {
		realHeaders["dingAccessTokenType"] = util.ToJSONString(headers.DingAccessTokenType)
	}

	if !tea.BoolValue(util.IsUnset(headers.DingOrgId)) {
		realHeaders["dingOrgId"] = util.ToJSONString(headers.DingOrgId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &GetDeptsByOrgIdResponse{}
	_body, _err := client.DoROARequest(tea.String("GetDeptsByOrgId"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/project/orgs/depts"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEmpsByOrgId(request *GetEmpsByOrgIdRequest) (_result *GetEmpsByOrgIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetEmpsByOrgIdHeaders{}
	_result = &GetEmpsByOrgIdResponse{}
	_body, _err := client.GetEmpsByOrgIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEmpsByOrgIdWithOptions(request *GetEmpsByOrgIdRequest, headers *GetEmpsByOrgIdHeaders, runtime *util.RuntimeOptions) (_result *GetEmpsByOrgIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NeedDept)) {
		query["needDept"] = request.NeedDept
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.DingAccessTokenType)) {
		realHeaders["dingAccessTokenType"] = util.ToJSONString(headers.DingAccessTokenType)
	}

	if !tea.BoolValue(util.IsUnset(headers.DingOrgId)) {
		realHeaders["dingOrgId"] = util.ToJSONString(headers.DingOrgId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &GetEmpsByOrgIdResponse{}
	_body, _err := client.DoROARequest(tea.String("GetEmpsByOrgId"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/project/orgs/employees"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOrganizatioTaskByIds(userId *string, request *GetOrganizatioTaskByIdsRequest) (_result *GetOrganizatioTaskByIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOrganizatioTaskByIdsHeaders{}
	_result = &GetOrganizatioTaskByIdsResponse{}
	_body, _err := client.GetOrganizatioTaskByIdsWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOrganizatioTaskByIdsWithOptions(userId *string, request *GetOrganizatioTaskByIdsRequest, headers *GetOrganizatioTaskByIdsHeaders, runtime *util.RuntimeOptions) (_result *GetOrganizatioTaskByIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskIds)) {
		query["taskIds"] = request.TaskIds
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
	_result = &GetOrganizatioTaskByIdsResponse{}
	_body, _err := client.DoROARequest(tea.String("GetOrganizatioTaskByIds"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/project/organizations/users/"+tea.StringValue(userId)+"/tasks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOrganizationPriorityList(userId *string) (_result *GetOrganizationPriorityListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOrganizationPriorityListHeaders{}
	_result = &GetOrganizationPriorityListResponse{}
	_body, _err := client.GetOrganizationPriorityListWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOrganizationPriorityListWithOptions(userId *string, headers *GetOrganizationPriorityListHeaders, runtime *util.RuntimeOptions) (_result *GetOrganizationPriorityListResponse, _err error) {
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
	_result = &GetOrganizationPriorityListResponse{}
	_body, _err := client.DoROARequest(tea.String("GetOrganizationPriorityList"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/project/organizations/users/"+tea.StringValue(userId)+"/priorities"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOrganizationTask(taskId *string, userId *string) (_result *GetOrganizationTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOrganizationTaskHeaders{}
	_result = &GetOrganizationTaskResponse{}
	_body, _err := client.GetOrganizationTaskWithOptions(taskId, userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOrganizationTaskWithOptions(taskId *string, userId *string, headers *GetOrganizationTaskHeaders, runtime *util.RuntimeOptions) (_result *GetOrganizationTaskResponse, _err error) {
	taskId = openapiutil.GetEncodeParam(taskId)
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
	_result = &GetOrganizationTaskResponse{}
	_body, _err := client.DoROARequest(tea.String("GetOrganizationTask"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/project/organizations/users/"+tea.StringValue(userId)+"/tasks/"+tea.StringValue(taskId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProjectGroup(userId *string, request *GetProjectGroupRequest) (_result *GetProjectGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetProjectGroupHeaders{}
	_result = &GetProjectGroupResponse{}
	_body, _err := client.GetProjectGroupWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProjectGroupWithOptions(userId *string, request *GetProjectGroupRequest, headers *GetProjectGroupHeaders, runtime *util.RuntimeOptions) (_result *GetProjectGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ViewerId)) {
		query["viewerId"] = request.ViewerId
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
	_result = &GetProjectGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("GetProjectGroup"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/project/organizations/users/"+tea.StringValue(userId)+"/groups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTbOrgIdByDingOrgId(request *GetTbOrgIdByDingOrgIdRequest) (_result *GetTbOrgIdByDingOrgIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTbOrgIdByDingOrgIdHeaders{}
	_result = &GetTbOrgIdByDingOrgIdResponse{}
	_body, _err := client.GetTbOrgIdByDingOrgIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTbOrgIdByDingOrgIdWithOptions(request *GetTbOrgIdByDingOrgIdRequest, headers *GetTbOrgIdByDingOrgIdHeaders, runtime *util.RuntimeOptions) (_result *GetTbOrgIdByDingOrgIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OptUserId)) {
		query["optUserId"] = request.OptUserId
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
	_result = &GetTbOrgIdByDingOrgIdResponse{}
	_body, _err := client.DoROARequest(tea.String("GetTbOrgIdByDingOrgId"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/project/teambition/organizations"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTbProjectGray(request *GetTbProjectGrayRequest) (_result *GetTbProjectGrayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTbProjectGrayHeaders{}
	_result = &GetTbProjectGrayResponse{}
	_body, _err := client.GetTbProjectGrayWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTbProjectGrayWithOptions(request *GetTbProjectGrayRequest, headers *GetTbProjectGrayHeaders, runtime *util.RuntimeOptions) (_result *GetTbProjectGrayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Label)) {
		body["label"] = request.Label
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.DingAccessTokenType)) {
		realHeaders["dingAccessTokenType"] = util.ToJSONString(headers.DingAccessTokenType)
	}

	if !tea.BoolValue(util.IsUnset(headers.DingCorpId)) {
		realHeaders["dingCorpId"] = util.ToJSONString(headers.DingCorpId)
	}

	if !tea.BoolValue(util.IsUnset(headers.DingIsvOrgId)) {
		realHeaders["dingIsvOrgId"] = util.ToJSONString(headers.DingIsvOrgId)
	}

	if !tea.BoolValue(util.IsUnset(headers.DingOrgId)) {
		realHeaders["dingOrgId"] = util.ToJSONString(headers.DingOrgId)
	}

	if !tea.BoolValue(util.IsUnset(headers.DingSuiteKey)) {
		realHeaders["dingSuiteKey"] = util.ToJSONString(headers.DingSuiteKey)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &GetTbProjectGrayResponse{}
	_body, _err := client.DoROARequest(tea.String("GetTbProjectGray"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/project/projects/gray"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTbProjectSource() (_result *GetTbProjectSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTbProjectSourceHeaders{}
	_result = &GetTbProjectSourceResponse{}
	_body, _err := client.GetTbProjectSourceWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTbProjectSourceWithOptions(headers *GetTbProjectSourceHeaders, runtime *util.RuntimeOptions) (_result *GetTbProjectSourceResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.DingAccessTokenType)) {
		realHeaders["dingAccessTokenType"] = util.ToJSONString(headers.DingAccessTokenType)
	}

	if !tea.BoolValue(util.IsUnset(headers.DingCorpId)) {
		realHeaders["dingCorpId"] = util.ToJSONString(headers.DingCorpId)
	}

	if !tea.BoolValue(util.IsUnset(headers.DingIsvOrgId)) {
		realHeaders["dingIsvOrgId"] = util.ToJSONString(headers.DingIsvOrgId)
	}

	if !tea.BoolValue(util.IsUnset(headers.DingOrgId)) {
		realHeaders["dingOrgId"] = util.ToJSONString(headers.DingOrgId)
	}

	if !tea.BoolValue(util.IsUnset(headers.DingSuiteKey)) {
		realHeaders["dingSuiteKey"] = util.ToJSONString(headers.DingSuiteKey)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &GetTbProjectSourceResponse{}
	_body, _err := client.DoROARequest(tea.String("GetTbProjectSource"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/project/projects/source"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTbUserIdByStaffId(request *GetTbUserIdByStaffIdRequest) (_result *GetTbUserIdByStaffIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTbUserIdByStaffIdHeaders{}
	_result = &GetTbUserIdByStaffIdResponse{}
	_body, _err := client.GetTbUserIdByStaffIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTbUserIdByStaffIdWithOptions(request *GetTbUserIdByStaffIdRequest, headers *GetTbUserIdByStaffIdHeaders, runtime *util.RuntimeOptions) (_result *GetTbUserIdByStaffIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OptUserId)) {
		query["optUserId"] = request.OptUserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
	_result = &GetTbUserIdByStaffIdResponse{}
	_body, _err := client.DoROARequest(tea.String("GetTbUserIdByStaffId"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/project/teambition/users"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchProjectTemplate(userId *string, request *SearchProjectTemplateRequest) (_result *SearchProjectTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchProjectTemplateHeaders{}
	_result = &SearchProjectTemplateResponse{}
	_body, _err := client.SearchProjectTemplateWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchProjectTemplateWithOptions(userId *string, request *SearchProjectTemplateRequest, headers *SearchProjectTemplateHeaders, runtime *util.RuntimeOptions) (_result *SearchProjectTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
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
	_result = &SearchProjectTemplateResponse{}
	_body, _err := client.DoROARequest(tea.String("SearchProjectTemplate"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/project/organizations/users/"+tea.StringValue(userId)+"/templates"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCustomfieldValue(userId *string, taskId *string, request *UpdateCustomfieldValueRequest) (_result *UpdateCustomfieldValueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateCustomfieldValueHeaders{}
	_result = &UpdateCustomfieldValueResponse{}
	_body, _err := client.UpdateCustomfieldValueWithOptions(userId, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCustomfieldValueWithOptions(userId *string, taskId *string, request *UpdateCustomfieldValueRequest, headers *UpdateCustomfieldValueHeaders, runtime *util.RuntimeOptions) (_result *UpdateCustomfieldValueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	taskId = openapiutil.GetEncodeParam(taskId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomfieldId)) {
		body["customfieldId"] = request.CustomfieldId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomfieldName)) {
		body["customfieldName"] = request.CustomfieldName
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["value"] = request.Value
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
	_result = &UpdateCustomfieldValueResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateCustomfieldValue"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/project/users/"+tea.StringValue(userId)+"/tasks/"+tea.StringValue(taskId)+"/customFields"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateOrganizationTaskContent(taskId *string, userId *string, request *UpdateOrganizationTaskContentRequest) (_result *UpdateOrganizationTaskContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateOrganizationTaskContentHeaders{}
	_result = &UpdateOrganizationTaskContentResponse{}
	_body, _err := client.UpdateOrganizationTaskContentWithOptions(taskId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateOrganizationTaskContentWithOptions(taskId *string, userId *string, request *UpdateOrganizationTaskContentRequest, headers *UpdateOrganizationTaskContentHeaders, runtime *util.RuntimeOptions) (_result *UpdateOrganizationTaskContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	taskId = openapiutil.GetEncodeParam(taskId)
	userId = openapiutil.GetEncodeParam(userId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DisableActivity)) {
		body["disableActivity"] = request.DisableActivity
	}

	if !tea.BoolValue(util.IsUnset(request.DisableNotification)) {
		body["disableNotification"] = request.DisableNotification
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
	_result = &UpdateOrganizationTaskContentResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateOrganizationTaskContent"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/project/organizations/users/"+tea.StringValue(userId)+"/tasks/"+tea.StringValue(taskId)+"/contents"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateOrganizationTaskDueDate(taskId *string, userId *string, request *UpdateOrganizationTaskDueDateRequest) (_result *UpdateOrganizationTaskDueDateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateOrganizationTaskDueDateHeaders{}
	_result = &UpdateOrganizationTaskDueDateResponse{}
	_body, _err := client.UpdateOrganizationTaskDueDateWithOptions(taskId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateOrganizationTaskDueDateWithOptions(taskId *string, userId *string, request *UpdateOrganizationTaskDueDateRequest, headers *UpdateOrganizationTaskDueDateHeaders, runtime *util.RuntimeOptions) (_result *UpdateOrganizationTaskDueDateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	taskId = openapiutil.GetEncodeParam(taskId)
	userId = openapiutil.GetEncodeParam(userId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisableActivity)) {
		body["disableActivity"] = request.DisableActivity
	}

	if !tea.BoolValue(util.IsUnset(request.DisableNotification)) {
		body["disableNotification"] = request.DisableNotification
	}

	if !tea.BoolValue(util.IsUnset(request.DueDate)) {
		body["dueDate"] = request.DueDate
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
	_result = &UpdateOrganizationTaskDueDateResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateOrganizationTaskDueDate"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/project/organizations/users/"+tea.StringValue(userId)+"/tasks/"+tea.StringValue(taskId)+"/dueDates"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateOrganizationTaskExecutor(taskId *string, userId *string, request *UpdateOrganizationTaskExecutorRequest) (_result *UpdateOrganizationTaskExecutorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateOrganizationTaskExecutorHeaders{}
	_result = &UpdateOrganizationTaskExecutorResponse{}
	_body, _err := client.UpdateOrganizationTaskExecutorWithOptions(taskId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateOrganizationTaskExecutorWithOptions(taskId *string, userId *string, request *UpdateOrganizationTaskExecutorRequest, headers *UpdateOrganizationTaskExecutorHeaders, runtime *util.RuntimeOptions) (_result *UpdateOrganizationTaskExecutorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	taskId = openapiutil.GetEncodeParam(taskId)
	userId = openapiutil.GetEncodeParam(userId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisableActivity)) {
		body["disableActivity"] = request.DisableActivity
	}

	if !tea.BoolValue(util.IsUnset(request.DisableNotification)) {
		body["disableNotification"] = request.DisableNotification
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorId)) {
		body["executorId"] = request.ExecutorId
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
	_result = &UpdateOrganizationTaskExecutorResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateOrganizationTaskExecutor"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/project/organizations/users/"+tea.StringValue(userId)+"/tasks/"+tea.StringValue(taskId)+"/executors"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateOrganizationTaskInvolveMembers(taskId *string, userId *string, request *UpdateOrganizationTaskInvolveMembersRequest) (_result *UpdateOrganizationTaskInvolveMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateOrganizationTaskInvolveMembersHeaders{}
	_result = &UpdateOrganizationTaskInvolveMembersResponse{}
	_body, _err := client.UpdateOrganizationTaskInvolveMembersWithOptions(taskId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateOrganizationTaskInvolveMembersWithOptions(taskId *string, userId *string, request *UpdateOrganizationTaskInvolveMembersRequest, headers *UpdateOrganizationTaskInvolveMembersHeaders, runtime *util.RuntimeOptions) (_result *UpdateOrganizationTaskInvolveMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	taskId = openapiutil.GetEncodeParam(taskId)
	userId = openapiutil.GetEncodeParam(userId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddInvolvers)) {
		body["addInvolvers"] = request.AddInvolvers
	}

	if !tea.BoolValue(util.IsUnset(request.DelInvolvers)) {
		body["delInvolvers"] = request.DelInvolvers
	}

	if !tea.BoolValue(util.IsUnset(request.DisableActivity)) {
		body["disableActivity"] = request.DisableActivity
	}

	if !tea.BoolValue(util.IsUnset(request.DisableNotification)) {
		body["disableNotification"] = request.DisableNotification
	}

	if !tea.BoolValue(util.IsUnset(request.InvolveMembers)) {
		body["involveMembers"] = request.InvolveMembers
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
	_result = &UpdateOrganizationTaskInvolveMembersResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateOrganizationTaskInvolveMembers"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/project/organizations/users/"+tea.StringValue(userId)+"/tasks/"+tea.StringValue(taskId)+"/involveMembers"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateOrganizationTaskNote(taskId *string, userId *string, request *UpdateOrganizationTaskNoteRequest) (_result *UpdateOrganizationTaskNoteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateOrganizationTaskNoteHeaders{}
	_result = &UpdateOrganizationTaskNoteResponse{}
	_body, _err := client.UpdateOrganizationTaskNoteWithOptions(taskId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateOrganizationTaskNoteWithOptions(taskId *string, userId *string, request *UpdateOrganizationTaskNoteRequest, headers *UpdateOrganizationTaskNoteHeaders, runtime *util.RuntimeOptions) (_result *UpdateOrganizationTaskNoteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	taskId = openapiutil.GetEncodeParam(taskId)
	userId = openapiutil.GetEncodeParam(userId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisableActivity)) {
		body["disableActivity"] = request.DisableActivity
	}

	if !tea.BoolValue(util.IsUnset(request.DisableNotification)) {
		body["disableNotification"] = request.DisableNotification
	}

	if !tea.BoolValue(util.IsUnset(request.Note)) {
		body["note"] = request.Note
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
	_result = &UpdateOrganizationTaskNoteResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateOrganizationTaskNote"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/project/organizations/users/"+tea.StringValue(userId)+"/tasks/"+tea.StringValue(taskId)+"/notes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateOrganizationTaskPriority(taskId *string, userId *string, request *UpdateOrganizationTaskPriorityRequest) (_result *UpdateOrganizationTaskPriorityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateOrganizationTaskPriorityHeaders{}
	_result = &UpdateOrganizationTaskPriorityResponse{}
	_body, _err := client.UpdateOrganizationTaskPriorityWithOptions(taskId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateOrganizationTaskPriorityWithOptions(taskId *string, userId *string, request *UpdateOrganizationTaskPriorityRequest, headers *UpdateOrganizationTaskPriorityHeaders, runtime *util.RuntimeOptions) (_result *UpdateOrganizationTaskPriorityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	taskId = openapiutil.GetEncodeParam(taskId)
	userId = openapiutil.GetEncodeParam(userId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisableActivity)) {
		body["disableActivity"] = request.DisableActivity
	}

	if !tea.BoolValue(util.IsUnset(request.DisableNotification)) {
		body["disableNotification"] = request.DisableNotification
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["priority"] = request.Priority
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
	_result = &UpdateOrganizationTaskPriorityResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateOrganizationTaskPriority"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/project/organizations/users/"+tea.StringValue(userId)+"/tasks/"+tea.StringValue(taskId)+"/priorities"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateOrganizationTaskStatus(taskId *string, userId *string, request *UpdateOrganizationTaskStatusRequest) (_result *UpdateOrganizationTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateOrganizationTaskStatusHeaders{}
	_result = &UpdateOrganizationTaskStatusResponse{}
	_body, _err := client.UpdateOrganizationTaskStatusWithOptions(taskId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateOrganizationTaskStatusWithOptions(taskId *string, userId *string, request *UpdateOrganizationTaskStatusRequest, headers *UpdateOrganizationTaskStatusHeaders, runtime *util.RuntimeOptions) (_result *UpdateOrganizationTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	taskId = openapiutil.GetEncodeParam(taskId)
	userId = openapiutil.GetEncodeParam(userId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisableActivity)) {
		body["disableActivity"] = request.DisableActivity
	}

	if !tea.BoolValue(util.IsUnset(request.DisableNotification)) {
		body["disableNotification"] = request.DisableNotification
	}

	if !tea.BoolValue(util.IsUnset(request.IsDone)) {
		body["isDone"] = request.IsDone
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
	_result = &UpdateOrganizationTaskStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateOrganizationTaskStatus"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/project/organizations/users/"+tea.StringValue(userId)+"/tasks/"+tea.StringValue(taskId)+"/states"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProjectGroup(userId *string, projectId *string, request *UpdateProjectGroupRequest) (_result *UpdateProjectGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateProjectGroupHeaders{}
	_result = &UpdateProjectGroupResponse{}
	_body, _err := client.UpdateProjectGroupWithOptions(userId, projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProjectGroupWithOptions(userId *string, projectId *string, request *UpdateProjectGroupRequest, headers *UpdateProjectGroupHeaders, runtime *util.RuntimeOptions) (_result *UpdateProjectGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	projectId = openapiutil.GetEncodeParam(projectId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddProjectGroupIds)) {
		body["addProjectGroupIds"] = request.AddProjectGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.DelProjectGroupIds)) {
		body["delProjectGroupIds"] = request.DelProjectGroupIds
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
	_result = &UpdateProjectGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateProjectGroup"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/project/users/"+tea.StringValue(userId)+"/projects/"+tea.StringValue(projectId)+"/groups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
