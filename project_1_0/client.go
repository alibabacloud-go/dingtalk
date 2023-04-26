// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package project_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddProjectMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddProjectMemberResponse) SetStatusCode(v int32) *AddProjectMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *AddProjectMemberResponse) SetBody(v *AddProjectMemberResponseBody) *AddProjectMemberResponse {
	s.Body = v
	return s
}

type ArchiveProjectHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ArchiveProjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s ArchiveProjectHeaders) GoString() string {
	return s.String()
}

func (s *ArchiveProjectHeaders) SetCommonHeaders(v map[string]*string) *ArchiveProjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ArchiveProjectHeaders) SetXAcsDingtalkAccessToken(v string) *ArchiveProjectHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ArchiveProjectResponseBody struct {
	Result *ArchiveProjectResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ArchiveProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ArchiveProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ArchiveProjectResponseBody) SetResult(v *ArchiveProjectResponseBodyResult) *ArchiveProjectResponseBody {
	s.Result = v
	return s
}

type ArchiveProjectResponseBodyResult struct {
	IsArchived *bool   `json:"isArchived,omitempty" xml:"isArchived,omitempty"`
	Updated    *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ArchiveProjectResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ArchiveProjectResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ArchiveProjectResponseBodyResult) SetIsArchived(v bool) *ArchiveProjectResponseBodyResult {
	s.IsArchived = &v
	return s
}

func (s *ArchiveProjectResponseBodyResult) SetUpdated(v string) *ArchiveProjectResponseBodyResult {
	s.Updated = &v
	return s
}

type ArchiveProjectResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ArchiveProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ArchiveProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s ArchiveProjectResponse) GoString() string {
	return s.String()
}

func (s *ArchiveProjectResponse) SetHeaders(v map[string]*string) *ArchiveProjectResponse {
	s.Headers = v
	return s
}

func (s *ArchiveProjectResponse) SetStatusCode(v int32) *ArchiveProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ArchiveProjectResponse) SetBody(v *ArchiveProjectResponseBody) *ArchiveProjectResponse {
	s.Body = v
	return s
}

type ArchiveTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ArchiveTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s ArchiveTaskHeaders) GoString() string {
	return s.String()
}

func (s *ArchiveTaskHeaders) SetCommonHeaders(v map[string]*string) *ArchiveTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ArchiveTaskHeaders) SetXAcsDingtalkAccessToken(v string) *ArchiveTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ArchiveTaskResponseBody struct {
	Result *ArchiveTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ArchiveTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ArchiveTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ArchiveTaskResponseBody) SetResult(v *ArchiveTaskResponseBodyResult) *ArchiveTaskResponseBody {
	s.Result = v
	return s
}

type ArchiveTaskResponseBodyResult struct {
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ArchiveTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ArchiveTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ArchiveTaskResponseBodyResult) SetUpdated(v string) *ArchiveTaskResponseBodyResult {
	s.Updated = &v
	return s
}

type ArchiveTaskResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ArchiveTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ArchiveTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ArchiveTaskResponse) GoString() string {
	return s.String()
}

func (s *ArchiveTaskResponse) SetHeaders(v map[string]*string) *ArchiveTaskResponse {
	s.Headers = v
	return s
}

func (s *ArchiveTaskResponse) SetStatusCode(v int32) *ArchiveTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ArchiveTaskResponse) SetBody(v *ArchiveTaskResponseBody) *ArchiveTaskResponse {
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
	Content             *string   `json:"content,omitempty" xml:"content,omitempty"`
	CreateTime          *string   `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DisableActivity     *bool     `json:"disableActivity,omitempty" xml:"disableActivity,omitempty"`
	DisableNotification *bool     `json:"disableNotification,omitempty" xml:"disableNotification,omitempty"`
	DueDate             *string   `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	ExecutorId          *string   `json:"executorId,omitempty" xml:"executorId,omitempty"`
	InvolveMembers      []*string `json:"involveMembers,omitempty" xml:"involveMembers,omitempty" type:"Repeated"`
	Note                *string   `json:"note,omitempty" xml:"note,omitempty"`
	Priority            *int32    `json:"priority,omitempty" xml:"priority,omitempty"`
	Visible             *string   `json:"visible,omitempty" xml:"visible,omitempty"`
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
	AncestorIds      []*string                                            `json:"ancestorIds,omitempty" xml:"ancestorIds,omitempty" type:"Repeated"`
	AttachmentsCount *int32                                               `json:"attachmentsCount,omitempty" xml:"attachmentsCount,omitempty"`
	Content          *string                                              `json:"content,omitempty" xml:"content,omitempty"`
	Created          *string                                              `json:"created,omitempty" xml:"created,omitempty"`
	Creator          *CreateOrganizationTaskResponseBodyResultCreator     `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	CreatorId        *string                                              `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	DueDate          *string                                              `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	Executor         *CreateOrganizationTaskResponseBodyResultExecutor    `json:"executor,omitempty" xml:"executor,omitempty" type:"Struct"`
	ExecutorId       *string                                              `json:"executorId,omitempty" xml:"executorId,omitempty"`
	HasReminder      *bool                                                `json:"hasReminder,omitempty" xml:"hasReminder,omitempty"`
	Id               *string                                              `json:"id,omitempty" xml:"id,omitempty"`
	InvolveMembers   []*string                                            `json:"involveMembers,omitempty" xml:"involveMembers,omitempty" type:"Repeated"`
	Involvers        []*CreateOrganizationTaskResponseBodyResultInvolvers `json:"involvers,omitempty" xml:"involvers,omitempty" type:"Repeated"`
	IsDeleted        *bool                                                `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	IsDone           *string                                              `json:"isDone,omitempty" xml:"isDone,omitempty"`
	Note             *string                                              `json:"note,omitempty" xml:"note,omitempty"`
	Priority         *int32                                               `json:"priority,omitempty" xml:"priority,omitempty"`
	Updated          *string                                              `json:"updated,omitempty" xml:"updated,omitempty"`
	Visible          *string                                              `json:"visible,omitempty" xml:"visible,omitempty"`
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
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	Id        *string `json:"id,omitempty" xml:"id,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateOrganizationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateOrganizationTaskResponse) SetStatusCode(v int32) *CreateOrganizationTaskResponse {
	s.StatusCode = &v
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
	EndDate          *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	ExecutorId       *string `json:"executorId,omitempty" xml:"executorId,omitempty"`
	IncludesHolidays *bool   `json:"includesHolidays,omitempty" xml:"includesHolidays,omitempty"`
	IsDuration       *bool   `json:"isDuration,omitempty" xml:"isDuration,omitempty"`
	ObjectId         *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	ObjectType       *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	PlanTime         *int64  `json:"planTime,omitempty" xml:"planTime,omitempty"`
	StartDate        *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	SubmitterId      *string `json:"submitterId,omitempty" xml:"submitterId,omitempty"`
	TenantType       *string `json:"tenantType,omitempty" xml:"tenantType,omitempty"`
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
	Body    []*CreatePlanTimeResponseBodyResultBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	Message *string                                 `json:"message,omitempty" xml:"message,omitempty"`
	Ok      *bool                                   `json:"ok,omitempty" xml:"ok,omitempty"`
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
	Date     *string `json:"date,omitempty" xml:"date,omitempty"`
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	PlanTime *int64  `json:"planTime,omitempty" xml:"planTime,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePlanTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreatePlanTimeResponse) SetStatusCode(v int32) *CreatePlanTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePlanTimeResponse) SetBody(v *CreatePlanTimeResponseBody) *CreatePlanTimeResponse {
	s.Body = v
	return s
}

type CreateProjectHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateProjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectHeaders) GoString() string {
	return s.String()
}

func (s *CreateProjectHeaders) SetCommonHeaders(v map[string]*string) *CreateProjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateProjectHeaders) SetXAcsDingtalkAccessToken(v string) *CreateProjectHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateProjectRequest struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetName(v string) *CreateProjectRequest {
	s.Name = &v
	return s
}

type CreateProjectResponseBody struct {
	Result *CreateProjectResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) SetResult(v *CreateProjectResponseBodyResult) *CreateProjectResponseBody {
	s.Result = v
	return s
}

type CreateProjectResponseBodyResult struct {
	Created             *string                                        `json:"created,omitempty" xml:"created,omitempty"`
	CreatorId           *string                                        `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Customfields        []*CreateProjectResponseBodyResultCustomfields `json:"customfields,omitempty" xml:"customfields,omitempty" type:"Repeated"`
	DefaultCollectionId *string                                        `json:"defaultCollectionId,omitempty" xml:"defaultCollectionId,omitempty"`
	IsArchived          *bool                                          `json:"isArchived,omitempty" xml:"isArchived,omitempty"`
	IsSuspended         *bool                                          `json:"isSuspended,omitempty" xml:"isSuspended,omitempty"`
	IsTemplate          *bool                                          `json:"isTemplate,omitempty" xml:"isTemplate,omitempty"`
	Logo                *string                                        `json:"logo,omitempty" xml:"logo,omitempty"`
	Name                *string                                        `json:"name,omitempty" xml:"name,omitempty"`
	NormalType          *string                                        `json:"normalType,omitempty" xml:"normalType,omitempty"`
	ProjectId           *string                                        `json:"projectId,omitempty" xml:"projectId,omitempty"`
	RootCollectionId    *string                                        `json:"rootCollectionId,omitempty" xml:"rootCollectionId,omitempty"`
	SourceId            *string                                        `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	UniqueIdPrefix      *string                                        `json:"uniqueIdPrefix,omitempty" xml:"uniqueIdPrefix,omitempty"`
	Updated             *string                                        `json:"updated,omitempty" xml:"updated,omitempty"`
	Visibility          *string                                        `json:"visibility,omitempty" xml:"visibility,omitempty"`
}

func (s CreateProjectResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBodyResult) SetCreated(v string) *CreateProjectResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateProjectResponseBodyResult) SetCreatorId(v string) *CreateProjectResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *CreateProjectResponseBodyResult) SetCustomfields(v []*CreateProjectResponseBodyResultCustomfields) *CreateProjectResponseBodyResult {
	s.Customfields = v
	return s
}

func (s *CreateProjectResponseBodyResult) SetDefaultCollectionId(v string) *CreateProjectResponseBodyResult {
	s.DefaultCollectionId = &v
	return s
}

func (s *CreateProjectResponseBodyResult) SetIsArchived(v bool) *CreateProjectResponseBodyResult {
	s.IsArchived = &v
	return s
}

func (s *CreateProjectResponseBodyResult) SetIsSuspended(v bool) *CreateProjectResponseBodyResult {
	s.IsSuspended = &v
	return s
}

func (s *CreateProjectResponseBodyResult) SetIsTemplate(v bool) *CreateProjectResponseBodyResult {
	s.IsTemplate = &v
	return s
}

func (s *CreateProjectResponseBodyResult) SetLogo(v string) *CreateProjectResponseBodyResult {
	s.Logo = &v
	return s
}

func (s *CreateProjectResponseBodyResult) SetName(v string) *CreateProjectResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateProjectResponseBodyResult) SetNormalType(v string) *CreateProjectResponseBodyResult {
	s.NormalType = &v
	return s
}

func (s *CreateProjectResponseBodyResult) SetProjectId(v string) *CreateProjectResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectResponseBodyResult) SetRootCollectionId(v string) *CreateProjectResponseBodyResult {
	s.RootCollectionId = &v
	return s
}

func (s *CreateProjectResponseBodyResult) SetSourceId(v string) *CreateProjectResponseBodyResult {
	s.SourceId = &v
	return s
}

func (s *CreateProjectResponseBodyResult) SetUniqueIdPrefix(v string) *CreateProjectResponseBodyResult {
	s.UniqueIdPrefix = &v
	return s
}

func (s *CreateProjectResponseBodyResult) SetUpdated(v string) *CreateProjectResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *CreateProjectResponseBodyResult) SetVisibility(v string) *CreateProjectResponseBodyResult {
	s.Visibility = &v
	return s
}

type CreateProjectResponseBodyResultCustomfields struct {
	CustomfieldId *string                                             `json:"customfieldId,omitempty" xml:"customfieldId,omitempty"`
	Type          *string                                             `json:"type,omitempty" xml:"type,omitempty"`
	Value         []*CreateProjectResponseBodyResultCustomfieldsValue `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s CreateProjectResponseBodyResultCustomfields) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBodyResultCustomfields) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBodyResultCustomfields) SetCustomfieldId(v string) *CreateProjectResponseBodyResultCustomfields {
	s.CustomfieldId = &v
	return s
}

func (s *CreateProjectResponseBodyResultCustomfields) SetType(v string) *CreateProjectResponseBodyResultCustomfields {
	s.Type = &v
	return s
}

func (s *CreateProjectResponseBodyResultCustomfields) SetValue(v []*CreateProjectResponseBodyResultCustomfieldsValue) *CreateProjectResponseBodyResultCustomfields {
	s.Value = v
	return s
}

type CreateProjectResponseBodyResultCustomfieldsValue struct {
	FieldvalueId *string `json:"fieldvalueId,omitempty" xml:"fieldvalueId,omitempty"`
	MetaString   *string `json:"metaString,omitempty" xml:"metaString,omitempty"`
	Title        *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateProjectResponseBodyResultCustomfieldsValue) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBodyResultCustomfieldsValue) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBodyResultCustomfieldsValue) SetFieldvalueId(v string) *CreateProjectResponseBodyResultCustomfieldsValue {
	s.FieldvalueId = &v
	return s
}

func (s *CreateProjectResponseBodyResultCustomfieldsValue) SetMetaString(v string) *CreateProjectResponseBodyResultCustomfieldsValue {
	s.MetaString = &v
	return s
}

func (s *CreateProjectResponseBodyResultCustomfieldsValue) SetTitle(v string) *CreateProjectResponseBodyResultCustomfieldsValue {
	s.Title = &v
	return s
}

type CreateProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectResponse) SetHeaders(v map[string]*string) *CreateProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectResponse) SetStatusCode(v int32) *CreateProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProjectResponse) SetBody(v *CreateProjectResponseBody) *CreateProjectResponse {
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
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
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
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	Id      *string `json:"id,omitempty" xml:"id,omitempty"`
	Logo    *string `json:"logo,omitempty" xml:"logo,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateProjectByTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateProjectByTemplateResponse) SetStatusCode(v int32) *CreateProjectByTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProjectByTemplateResponse) SetBody(v *CreateProjectByTemplateResponseBody) *CreateProjectByTemplateResponse {
	s.Body = v
	return s
}

type CreateProjectCustomfieldStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateProjectCustomfieldStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectCustomfieldStatusHeaders) GoString() string {
	return s.String()
}

func (s *CreateProjectCustomfieldStatusHeaders) SetCommonHeaders(v map[string]*string) *CreateProjectCustomfieldStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateProjectCustomfieldStatusHeaders) SetXAcsDingtalkAccessToken(v string) *CreateProjectCustomfieldStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateProjectCustomfieldStatusRequest struct {
	CustomfieldId         *string                                       `json:"customfieldId,omitempty" xml:"customfieldId,omitempty"`
	CustomfieldInstanceId *string                                       `json:"customfieldInstanceId,omitempty" xml:"customfieldInstanceId,omitempty"`
	CustomfieldName       *string                                       `json:"customfieldName,omitempty" xml:"customfieldName,omitempty"`
	Value                 []*CreateProjectCustomfieldStatusRequestValue `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s CreateProjectCustomfieldStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectCustomfieldStatusRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectCustomfieldStatusRequest) SetCustomfieldId(v string) *CreateProjectCustomfieldStatusRequest {
	s.CustomfieldId = &v
	return s
}

func (s *CreateProjectCustomfieldStatusRequest) SetCustomfieldInstanceId(v string) *CreateProjectCustomfieldStatusRequest {
	s.CustomfieldInstanceId = &v
	return s
}

func (s *CreateProjectCustomfieldStatusRequest) SetCustomfieldName(v string) *CreateProjectCustomfieldStatusRequest {
	s.CustomfieldName = &v
	return s
}

func (s *CreateProjectCustomfieldStatusRequest) SetValue(v []*CreateProjectCustomfieldStatusRequestValue) *CreateProjectCustomfieldStatusRequest {
	s.Value = v
	return s
}

type CreateProjectCustomfieldStatusRequestValue struct {
	FieldvalueId *string `json:"fieldvalueId,omitempty" xml:"fieldvalueId,omitempty"`
	MetaString   *string `json:"metaString,omitempty" xml:"metaString,omitempty"`
	Title        *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateProjectCustomfieldStatusRequestValue) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectCustomfieldStatusRequestValue) GoString() string {
	return s.String()
}

func (s *CreateProjectCustomfieldStatusRequestValue) SetFieldvalueId(v string) *CreateProjectCustomfieldStatusRequestValue {
	s.FieldvalueId = &v
	return s
}

func (s *CreateProjectCustomfieldStatusRequestValue) SetMetaString(v string) *CreateProjectCustomfieldStatusRequestValue {
	s.MetaString = &v
	return s
}

func (s *CreateProjectCustomfieldStatusRequestValue) SetTitle(v string) *CreateProjectCustomfieldStatusRequestValue {
	s.Title = &v
	return s
}

type CreateProjectCustomfieldStatusResponseBody struct {
	Result *CreateProjectCustomfieldStatusResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateProjectCustomfieldStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectCustomfieldStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectCustomfieldStatusResponseBody) SetResult(v *CreateProjectCustomfieldStatusResponseBodyResult) *CreateProjectCustomfieldStatusResponseBody {
	s.Result = v
	return s
}

type CreateProjectCustomfieldStatusResponseBodyResult struct {
	AdvCfObjectType *string                                                  `json:"advCfObjectType,omitempty" xml:"advCfObjectType,omitempty"`
	CustomfieldId   *string                                                  `json:"customfieldId,omitempty" xml:"customfieldId,omitempty"`
	Name            *string                                                  `json:"name,omitempty" xml:"name,omitempty"`
	OriginalId      *string                                                  `json:"originalId,omitempty" xml:"originalId,omitempty"`
	Type            *string                                                  `json:"type,omitempty" xml:"type,omitempty"`
	Value           []*CreateProjectCustomfieldStatusResponseBodyResultValue `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s CreateProjectCustomfieldStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectCustomfieldStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateProjectCustomfieldStatusResponseBodyResult) SetAdvCfObjectType(v string) *CreateProjectCustomfieldStatusResponseBodyResult {
	s.AdvCfObjectType = &v
	return s
}

func (s *CreateProjectCustomfieldStatusResponseBodyResult) SetCustomfieldId(v string) *CreateProjectCustomfieldStatusResponseBodyResult {
	s.CustomfieldId = &v
	return s
}

func (s *CreateProjectCustomfieldStatusResponseBodyResult) SetName(v string) *CreateProjectCustomfieldStatusResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateProjectCustomfieldStatusResponseBodyResult) SetOriginalId(v string) *CreateProjectCustomfieldStatusResponseBodyResult {
	s.OriginalId = &v
	return s
}

func (s *CreateProjectCustomfieldStatusResponseBodyResult) SetType(v string) *CreateProjectCustomfieldStatusResponseBodyResult {
	s.Type = &v
	return s
}

func (s *CreateProjectCustomfieldStatusResponseBodyResult) SetValue(v []*CreateProjectCustomfieldStatusResponseBodyResultValue) *CreateProjectCustomfieldStatusResponseBodyResult {
	s.Value = v
	return s
}

type CreateProjectCustomfieldStatusResponseBodyResultValue struct {
	FieldvalueId *string `json:"fieldvalueId,omitempty" xml:"fieldvalueId,omitempty"`
	MetaString   *string `json:"metaString,omitempty" xml:"metaString,omitempty"`
	Title        *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateProjectCustomfieldStatusResponseBodyResultValue) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectCustomfieldStatusResponseBodyResultValue) GoString() string {
	return s.String()
}

func (s *CreateProjectCustomfieldStatusResponseBodyResultValue) SetFieldvalueId(v string) *CreateProjectCustomfieldStatusResponseBodyResultValue {
	s.FieldvalueId = &v
	return s
}

func (s *CreateProjectCustomfieldStatusResponseBodyResultValue) SetMetaString(v string) *CreateProjectCustomfieldStatusResponseBodyResultValue {
	s.MetaString = &v
	return s
}

func (s *CreateProjectCustomfieldStatusResponseBodyResultValue) SetTitle(v string) *CreateProjectCustomfieldStatusResponseBodyResultValue {
	s.Title = &v
	return s
}

type CreateProjectCustomfieldStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateProjectCustomfieldStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProjectCustomfieldStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectCustomfieldStatusResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectCustomfieldStatusResponse) SetHeaders(v map[string]*string) *CreateProjectCustomfieldStatusResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectCustomfieldStatusResponse) SetStatusCode(v int32) *CreateProjectCustomfieldStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProjectCustomfieldStatusResponse) SetBody(v *CreateProjectCustomfieldStatusResponseBody) *CreateProjectCustomfieldStatusResponse {
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
	Content               *string                          `json:"content,omitempty" xml:"content,omitempty"`
	Customfields          []*CreateTaskRequestCustomfields `json:"customfields,omitempty" xml:"customfields,omitempty" type:"Repeated"`
	DueDate               *string                          `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	ExecutorId            *string                          `json:"executorId,omitempty" xml:"executorId,omitempty"`
	Note                  *string                          `json:"note,omitempty" xml:"note,omitempty"`
	ParentTaskId          *string                          `json:"parentTaskId,omitempty" xml:"parentTaskId,omitempty"`
	Priority              *int32                           `json:"priority,omitempty" xml:"priority,omitempty"`
	ProjectId             *string                          `json:"projectId,omitempty" xml:"projectId,omitempty"`
	ScenariofieldconfigId *string                          `json:"scenariofieldconfigId,omitempty" xml:"scenariofieldconfigId,omitempty"`
	StageId               *string                          `json:"stageId,omitempty" xml:"stageId,omitempty"`
	StartDate             *string                          `json:"startDate,omitempty" xml:"startDate,omitempty"`
	Visible               *string                          `json:"visible,omitempty" xml:"visible,omitempty"`
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

func (s *CreateTaskRequest) SetParentTaskId(v string) *CreateTaskRequest {
	s.ParentTaskId = &v
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

func (s *CreateTaskRequest) SetScenariofieldconfigId(v string) *CreateTaskRequest {
	s.ScenariofieldconfigId = &v
	return s
}

func (s *CreateTaskRequest) SetStageId(v string) *CreateTaskRequest {
	s.StageId = &v
	return s
}

func (s *CreateTaskRequest) SetStartDate(v string) *CreateTaskRequest {
	s.StartDate = &v
	return s
}

func (s *CreateTaskRequest) SetVisible(v string) *CreateTaskRequest {
	s.Visible = &v
	return s
}

type CreateTaskRequestCustomfields struct {
	CustomfieldId   *string                               `json:"customfieldId,omitempty" xml:"customfieldId,omitempty"`
	CustomfieldName *string                               `json:"customfieldName,omitempty" xml:"customfieldName,omitempty"`
	Value           []*CreateTaskRequestCustomfieldsValue `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
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
	Content        *string                                     `json:"content,omitempty" xml:"content,omitempty"`
	Created        *string                                     `json:"created,omitempty" xml:"created,omitempty"`
	CreatorId      *string                                     `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Customfields   []*CreateTaskResponseBodyResultCustomfields `json:"customfields,omitempty" xml:"customfields,omitempty" type:"Repeated"`
	DueDate        *string                                     `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	ExecutorId     *string                                     `json:"executorId,omitempty" xml:"executorId,omitempty"`
	InvolveMembers []*string                                   `json:"involveMembers,omitempty" xml:"involveMembers,omitempty" type:"Repeated"`
	Note           *string                                     `json:"note,omitempty" xml:"note,omitempty"`
	Priority       *int32                                      `json:"priority,omitempty" xml:"priority,omitempty"`
	ProjectId      *string                                     `json:"projectId,omitempty" xml:"projectId,omitempty"`
	TaskId         *string                                     `json:"taskId,omitempty" xml:"taskId,omitempty"`
	Updated        *string                                     `json:"updated,omitempty" xml:"updated,omitempty"`
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
	CustomfieldId *string                                          `json:"customfieldId,omitempty" xml:"customfieldId,omitempty"`
	Value         []*CreateTaskResponseBodyResultCustomfieldsValue `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateTaskResponse) SetStatusCode(v int32) *CreateTaskResponse {
	s.StatusCode = &v
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
	Content      *string `json:"content,omitempty" xml:"content,omitempty"`
	ThumbnailUrl *string `json:"thumbnailUrl,omitempty" xml:"thumbnailUrl,omitempty"`
	Title        *string `json:"title,omitempty" xml:"title,omitempty"`
	Url          *string `json:"url,omitempty" xml:"url,omitempty"`
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
	Created      *string `json:"created,omitempty" xml:"created,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTaskObjectLinkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateTaskObjectLinkResponse) SetStatusCode(v int32) *CreateTaskObjectLinkResponse {
	s.StatusCode = &v
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
	EndDate          *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	ExecutorId       *string `json:"executorId,omitempty" xml:"executorId,omitempty"`
	IncludesHolidays *bool   `json:"includesHolidays,omitempty" xml:"includesHolidays,omitempty"`
	IsDuration       *bool   `json:"isDuration,omitempty" xml:"isDuration,omitempty"`
	ObjectId         *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	ObjectType       *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	StartDate        *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	SubmitterId      *string `json:"submitterId,omitempty" xml:"submitterId,omitempty"`
	WorkTime         *int64  `json:"workTime,omitempty" xml:"workTime,omitempty"`
	TenantType       *string `json:"tenantType,omitempty" xml:"tenantType,omitempty"`
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
	Body    []*CreateWorkTimeResponseBodyResultBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	Message *string                                 `json:"message,omitempty" xml:"message,omitempty"`
	Ok      *bool                                   `json:"ok,omitempty" xml:"ok,omitempty"`
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
	Date     *string `json:"date,omitempty" xml:"date,omitempty"`
	TaskId   *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	WorkTime *int64  `json:"workTime,omitempty" xml:"workTime,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateWorkTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateWorkTimeResponse) SetStatusCode(v int32) *CreateWorkTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkTimeResponse) SetBody(v *CreateWorkTimeResponseBody) *CreateWorkTimeResponse {
	s.Body = v
	return s
}

type CreateWorkTimeApproveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateWorkTimeApproveHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkTimeApproveHeaders) GoString() string {
	return s.String()
}

func (s *CreateWorkTimeApproveHeaders) SetCommonHeaders(v map[string]*string) *CreateWorkTimeApproveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateWorkTimeApproveHeaders) SetXAcsDingtalkAccessToken(v string) *CreateWorkTimeApproveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateWorkTimeApproveRequest struct {
	WorkTimeIds []*string `json:"workTimeIds,omitempty" xml:"workTimeIds,omitempty" type:"Repeated"`
}

func (s CreateWorkTimeApproveRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkTimeApproveRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkTimeApproveRequest) SetWorkTimeIds(v []*string) *CreateWorkTimeApproveRequest {
	s.WorkTimeIds = v
	return s
}

type CreateWorkTimeApproveResponseBody struct {
	Message   *string                                  `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateWorkTimeApproveResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateWorkTimeApproveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkTimeApproveResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkTimeApproveResponseBody) SetMessage(v string) *CreateWorkTimeApproveResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWorkTimeApproveResponseBody) SetRequestId(v string) *CreateWorkTimeApproveResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkTimeApproveResponseBody) SetResult(v *CreateWorkTimeApproveResponseBodyResult) *CreateWorkTimeApproveResponseBody {
	s.Result = v
	return s
}

type CreateWorkTimeApproveResponseBodyResult struct {
	ApproveOpenId  *string   `json:"approveOpenId,omitempty" xml:"approveOpenId,omitempty"`
	CreatedAt      *string   `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatorId      *string   `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	OrganizationId *string   `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	Status         *string   `json:"status,omitempty" xml:"status,omitempty"`
	TaskId         *string   `json:"taskId,omitempty" xml:"taskId,omitempty"`
	Time           *int32    `json:"time,omitempty" xml:"time,omitempty"`
	UpdatedAt      *string   `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	UserId         *string   `json:"userId,omitempty" xml:"userId,omitempty"`
	WorkTimeIds    []*string `json:"workTimeIds,omitempty" xml:"workTimeIds,omitempty" type:"Repeated"`
}

func (s CreateWorkTimeApproveResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkTimeApproveResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateWorkTimeApproveResponseBodyResult) SetApproveOpenId(v string) *CreateWorkTimeApproveResponseBodyResult {
	s.ApproveOpenId = &v
	return s
}

func (s *CreateWorkTimeApproveResponseBodyResult) SetCreatedAt(v string) *CreateWorkTimeApproveResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *CreateWorkTimeApproveResponseBodyResult) SetCreatorId(v string) *CreateWorkTimeApproveResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *CreateWorkTimeApproveResponseBodyResult) SetOrganizationId(v string) *CreateWorkTimeApproveResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *CreateWorkTimeApproveResponseBodyResult) SetStatus(v string) *CreateWorkTimeApproveResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CreateWorkTimeApproveResponseBodyResult) SetTaskId(v string) *CreateWorkTimeApproveResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *CreateWorkTimeApproveResponseBodyResult) SetTime(v int32) *CreateWorkTimeApproveResponseBodyResult {
	s.Time = &v
	return s
}

func (s *CreateWorkTimeApproveResponseBodyResult) SetUpdatedAt(v string) *CreateWorkTimeApproveResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *CreateWorkTimeApproveResponseBodyResult) SetUserId(v string) *CreateWorkTimeApproveResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *CreateWorkTimeApproveResponseBodyResult) SetWorkTimeIds(v []*string) *CreateWorkTimeApproveResponseBodyResult {
	s.WorkTimeIds = v
	return s
}

type CreateWorkTimeApproveResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateWorkTimeApproveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWorkTimeApproveResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkTimeApproveResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkTimeApproveResponse) SetHeaders(v map[string]*string) *CreateWorkTimeApproveResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkTimeApproveResponse) SetStatusCode(v int32) *CreateWorkTimeApproveResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkTimeApproveResponse) SetBody(v *CreateWorkTimeApproveResponseBody) *CreateWorkTimeApproveResponse {
	s.Body = v
	return s
}

type DeleteProjectMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteProjectMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectMemberHeaders) GoString() string {
	return s.String()
}

func (s *DeleteProjectMemberHeaders) SetCommonHeaders(v map[string]*string) *DeleteProjectMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteProjectMemberHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteProjectMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteProjectMemberRequest struct {
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s DeleteProjectMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectMemberRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectMemberRequest) SetUserIds(v []*string) *DeleteProjectMemberRequest {
	s.UserIds = v
	return s
}

type DeleteProjectMemberResponseBody struct {
	Result []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s DeleteProjectMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectMemberResponseBody) SetResult(v []*string) *DeleteProjectMemberResponseBody {
	s.Result = v
	return s
}

type DeleteProjectMemberResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteProjectMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProjectMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectMemberResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectMemberResponse) SetHeaders(v map[string]*string) *DeleteProjectMemberResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectMemberResponse) SetStatusCode(v int32) *DeleteProjectMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProjectMemberResponse) SetBody(v *DeleteProjectMemberResponseBody) *DeleteProjectMemberResponse {
	s.Body = v
	return s
}

type DeleteTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteTaskHeaders) GoString() string {
	return s.String()
}

func (s *DeleteTaskHeaders) SetCommonHeaders(v map[string]*string) *DeleteTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteTaskHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteTaskResponseBody struct {
	Result map[string]*string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTaskResponseBody) SetResult(v map[string]*string) *DeleteTaskResponseBody {
	s.Result = v
	return s
}

type DeleteTaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteTaskResponse) SetHeaders(v map[string]*string) *DeleteTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteTaskResponse) SetStatusCode(v int32) *DeleteTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTaskResponse) SetBody(v *DeleteTaskResponseBody) *DeleteTaskResponse {
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
	DeptList   []*GetDeptsByOrgIdResponseBodyDeptList `json:"deptList,omitempty" xml:"deptList,omitempty" type:"Repeated"`
	HasMore    *bool                                  `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	MaxResults *int32                                 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int64                                 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	DeptId   *int64  `json:"dept_id,omitempty" xml:"dept_id,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	ParentId *int64  `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDeptsByOrgIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetDeptsByOrgIdResponse) SetStatusCode(v int32) *GetDeptsByOrgIdResponse {
	s.StatusCode = &v
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
	EmpList   []*GetEmpsByOrgIdResponseBodyEmpList `json:"empList,omitempty" xml:"empList,omitempty" type:"Repeated"`
	HasMore   *bool                                `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken *int64                               `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	Avatar     *string  `json:"avatar,omitempty" xml:"avatar,omitempty"`
	DeptIdList []*int64 `json:"dept_id_list,omitempty" xml:"dept_id_list,omitempty" type:"Repeated"`
	DingId     *string  `json:"dingId,omitempty" xml:"dingId,omitempty"`
	Name       *string  `json:"name,omitempty" xml:"name,omitempty"`
	Nick       *string  `json:"nick,omitempty" xml:"nick,omitempty"`
	OrgId      *int64   `json:"orgId,omitempty" xml:"orgId,omitempty"`
	Position   *string  `json:"position,omitempty" xml:"position,omitempty"`
	Unionid    *string  `json:"unionid,omitempty" xml:"unionid,omitempty"`
	Userid     *string  `json:"userid,omitempty" xml:"userid,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetEmpsByOrgIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetEmpsByOrgIdResponse) SetStatusCode(v int32) *GetEmpsByOrgIdResponse {
	s.StatusCode = &v
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
	AncestorIds    []*string `json:"ancestorIds,omitempty" xml:"ancestorIds,omitempty" type:"Repeated"`
	Content        *string   `json:"content,omitempty" xml:"content,omitempty"`
	Created        *string   `json:"created,omitempty" xml:"created,omitempty"`
	CreatorId      *string   `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	DueDate        *string   `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	ExecutorId     *string   `json:"executorId,omitempty" xml:"executorId,omitempty"`
	InvolveMembers []*string `json:"involveMembers,omitempty" xml:"involveMembers,omitempty" type:"Repeated"`
	IsDeleted      *bool     `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	IsDone         *bool     `json:"isDone,omitempty" xml:"isDone,omitempty"`
	Labels         []*string `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	Note           *string   `json:"note,omitempty" xml:"note,omitempty"`
	Priority       *int32    `json:"priority,omitempty" xml:"priority,omitempty"`
	StartDate      *string   `json:"startDate,omitempty" xml:"startDate,omitempty"`
	TaskId         *string   `json:"taskId,omitempty" xml:"taskId,omitempty"`
	Updated        *string   `json:"updated,omitempty" xml:"updated,omitempty"`
	Visible        *string   `json:"visible,omitempty" xml:"visible,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOrganizatioTaskByIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetOrganizatioTaskByIdsResponse) SetStatusCode(v int32) *GetOrganizatioTaskByIdsResponse {
	s.StatusCode = &v
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
	Color      *string `json:"color,omitempty" xml:"color,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	Priority   *string `json:"priority,omitempty" xml:"priority,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOrganizationPriorityListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetOrganizationPriorityListResponse) SetStatusCode(v int32) *GetOrganizationPriorityListResponse {
	s.StatusCode = &v
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
	AncestorIds    []*string `json:"ancestorIds,omitempty" xml:"ancestorIds,omitempty" type:"Repeated"`
	Content        *string   `json:"content,omitempty" xml:"content,omitempty"`
	Created        *string   `json:"created,omitempty" xml:"created,omitempty"`
	CreatorId      *string   `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	DueDate        *string   `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	ExecutorId     *string   `json:"executorId,omitempty" xml:"executorId,omitempty"`
	InvolveMembers []*string `json:"involveMembers,omitempty" xml:"involveMembers,omitempty" type:"Repeated"`
	IsDeleted      *bool     `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	IsDone         *bool     `json:"isDone,omitempty" xml:"isDone,omitempty"`
	Labels         []*string `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	Note           *string   `json:"note,omitempty" xml:"note,omitempty"`
	Priority       *int32    `json:"priority,omitempty" xml:"priority,omitempty"`
	StartDate      *string   `json:"startDate,omitempty" xml:"startDate,omitempty"`
	TaskId         *string   `json:"taskId,omitempty" xml:"taskId,omitempty"`
	Updated        *string   `json:"updated,omitempty" xml:"updated,omitempty"`
	Visible        *string   `json:"visible,omitempty" xml:"visible,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOrganizationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetOrganizationTaskResponse) SetStatusCode(v int32) *GetOrganizationTaskResponse {
	s.StatusCode = &v
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
	PageSize *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
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
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	Id      *string `json:"id,omitempty" xml:"id,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetProjectGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetProjectGroupResponse) SetStatusCode(v int32) *GetProjectGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectGroupResponse) SetBody(v *GetProjectGroupResponseBody) *GetProjectGroupResponse {
	s.Body = v
	return s
}

type GetProjectMemebersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetProjectMemebersHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMemebersHeaders) GoString() string {
	return s.String()
}

func (s *GetProjectMemebersHeaders) SetCommonHeaders(v map[string]*string) *GetProjectMemebersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetProjectMemebersHeaders) SetXAcsDingtalkAccessToken(v string) *GetProjectMemebersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetProjectMemebersRequest struct {
	MaxResults    *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	ProjectRoleId *string `json:"projectRoleId,omitempty" xml:"projectRoleId,omitempty"`
	Skip          *int32  `json:"skip,omitempty" xml:"skip,omitempty"`
	UserIds       *string `json:"userIds,omitempty" xml:"userIds,omitempty"`
}

func (s GetProjectMemebersRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMemebersRequest) GoString() string {
	return s.String()
}

func (s *GetProjectMemebersRequest) SetMaxResults(v int32) *GetProjectMemebersRequest {
	s.MaxResults = &v
	return s
}

func (s *GetProjectMemebersRequest) SetProjectRoleId(v string) *GetProjectMemebersRequest {
	s.ProjectRoleId = &v
	return s
}

func (s *GetProjectMemebersRequest) SetSkip(v int32) *GetProjectMemebersRequest {
	s.Skip = &v
	return s
}

func (s *GetProjectMemebersRequest) SetUserIds(v string) *GetProjectMemebersRequest {
	s.UserIds = &v
	return s
}

type GetProjectMemebersResponseBody struct {
	Result []*GetProjectMemebersResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetProjectMemebersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMemebersResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectMemebersResponseBody) SetResult(v []*GetProjectMemebersResponseBodyResult) *GetProjectMemebersResponseBody {
	s.Result = v
	return s
}

type GetProjectMemebersResponseBodyResult struct {
	// Deprecated
	MemberId *string   `json:"memberId,omitempty" xml:"memberId,omitempty"`
	Role     *int32    `json:"role,omitempty" xml:"role,omitempty"`
	RoleIds  []*string `json:"roleIds,omitempty" xml:"roleIds,omitempty" type:"Repeated"`
	UserId   *string   `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetProjectMemebersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMemebersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetProjectMemebersResponseBodyResult) SetMemberId(v string) *GetProjectMemebersResponseBodyResult {
	s.MemberId = &v
	return s
}

func (s *GetProjectMemebersResponseBodyResult) SetRole(v int32) *GetProjectMemebersResponseBodyResult {
	s.Role = &v
	return s
}

func (s *GetProjectMemebersResponseBodyResult) SetRoleIds(v []*string) *GetProjectMemebersResponseBodyResult {
	s.RoleIds = v
	return s
}

func (s *GetProjectMemebersResponseBodyResult) SetUserId(v string) *GetProjectMemebersResponseBodyResult {
	s.UserId = &v
	return s
}

type GetProjectMemebersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetProjectMemebersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProjectMemebersResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMemebersResponse) GoString() string {
	return s.String()
}

func (s *GetProjectMemebersResponse) SetHeaders(v map[string]*string) *GetProjectMemebersResponse {
	s.Headers = v
	return s
}

func (s *GetProjectMemebersResponse) SetStatusCode(v int32) *GetProjectMemebersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectMemebersResponse) SetBody(v *GetProjectMemebersResponseBody) *GetProjectMemebersResponse {
	s.Body = v
	return s
}

type GetProjectStatusListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetProjectStatusListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetProjectStatusListHeaders) GoString() string {
	return s.String()
}

func (s *GetProjectStatusListHeaders) SetCommonHeaders(v map[string]*string) *GetProjectStatusListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetProjectStatusListHeaders) SetXAcsDingtalkAccessToken(v string) *GetProjectStatusListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetProjectStatusListResponseBody struct {
	Result []*GetProjectStatusListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetProjectStatusListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectStatusListResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectStatusListResponseBody) SetResult(v []*GetProjectStatusListResponseBodyResult) *GetProjectStatusListResponseBody {
	s.Result = v
	return s
}

type GetProjectStatusListResponseBodyResult struct {
	Content   *string `json:"content,omitempty" xml:"content,omitempty"`
	Created   *string `json:"created,omitempty" xml:"created,omitempty"`
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Degree    *string `json:"degree,omitempty" xml:"degree,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s GetProjectStatusListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetProjectStatusListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetProjectStatusListResponseBodyResult) SetContent(v string) *GetProjectStatusListResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetProjectStatusListResponseBodyResult) SetCreated(v string) *GetProjectStatusListResponseBodyResult {
	s.Created = &v
	return s
}

func (s *GetProjectStatusListResponseBodyResult) SetCreatorId(v string) *GetProjectStatusListResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *GetProjectStatusListResponseBodyResult) SetDegree(v string) *GetProjectStatusListResponseBodyResult {
	s.Degree = &v
	return s
}

func (s *GetProjectStatusListResponseBodyResult) SetName(v string) *GetProjectStatusListResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetProjectStatusListResponseBodyResult) SetProjectId(v string) *GetProjectStatusListResponseBodyResult {
	s.ProjectId = &v
	return s
}

type GetProjectStatusListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetProjectStatusListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProjectStatusListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectStatusListResponse) GoString() string {
	return s.String()
}

func (s *GetProjectStatusListResponse) SetHeaders(v map[string]*string) *GetProjectStatusListResponse {
	s.Headers = v
	return s
}

func (s *GetProjectStatusListResponse) SetStatusCode(v int32) *GetProjectStatusListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectStatusListResponse) SetBody(v *GetProjectStatusListResponseBody) *GetProjectStatusListResponse {
	s.Body = v
	return s
}

type GetTaskByIdsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTaskByIdsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTaskByIdsHeaders) GoString() string {
	return s.String()
}

func (s *GetTaskByIdsHeaders) SetCommonHeaders(v map[string]*string) *GetTaskByIdsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTaskByIdsHeaders) SetXAcsDingtalkAccessToken(v string) *GetTaskByIdsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTaskByIdsRequest struct {
	ParentTaskId *string `json:"parentTaskId,omitempty" xml:"parentTaskId,omitempty"`
	TaskId       *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetTaskByIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskByIdsRequest) GoString() string {
	return s.String()
}

func (s *GetTaskByIdsRequest) SetParentTaskId(v string) *GetTaskByIdsRequest {
	s.ParentTaskId = &v
	return s
}

func (s *GetTaskByIdsRequest) SetTaskId(v string) *GetTaskByIdsRequest {
	s.TaskId = &v
	return s
}

type GetTaskByIdsResponseBody struct {
	Result []*GetTaskByIdsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetTaskByIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskByIdsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskByIdsResponseBody) SetResult(v []*GetTaskByIdsResponseBodyResult) *GetTaskByIdsResponseBody {
	s.Result = v
	return s
}

type GetTaskByIdsResponseBodyResult struct {
	AccomplishTime        *string                                       `json:"accomplishTime,omitempty" xml:"accomplishTime,omitempty"`
	AncestorIds           []*string                                     `json:"ancestorIds,omitempty" xml:"ancestorIds,omitempty" type:"Repeated"`
	Content               *string                                       `json:"content,omitempty" xml:"content,omitempty"`
	Created               *string                                       `json:"created,omitempty" xml:"created,omitempty"`
	CreatorId             *string                                       `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Customfields          []*GetTaskByIdsResponseBodyResultCustomfields `json:"customfields,omitempty" xml:"customfields,omitempty" type:"Repeated"`
	DueDate               *string                                       `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	ExecutorId            *string                                       `json:"executorId,omitempty" xml:"executorId,omitempty"`
	InvolveMembers        []*string                                     `json:"involveMembers,omitempty" xml:"involveMembers,omitempty" type:"Repeated"`
	IsArchived            *bool                                         `json:"isArchived,omitempty" xml:"isArchived,omitempty"`
	IsDone                *bool                                         `json:"isDone,omitempty" xml:"isDone,omitempty"`
	Note                  *string                                       `json:"note,omitempty" xml:"note,omitempty"`
	ParentTaskId          *string                                       `json:"parentTaskId,omitempty" xml:"parentTaskId,omitempty"`
	Priority              *int32                                        `json:"priority,omitempty" xml:"priority,omitempty"`
	ProjectId             *string                                       `json:"projectId,omitempty" xml:"projectId,omitempty"`
	Recurrence            []*string                                     `json:"recurrence,omitempty" xml:"recurrence,omitempty" type:"Repeated"`
	ScenariofieldconfigId *string                                       `json:"scenariofieldconfigId,omitempty" xml:"scenariofieldconfigId,omitempty"`
	SprintId              *string                                       `json:"sprintId,omitempty" xml:"sprintId,omitempty"`
	StageId               *string                                       `json:"stageId,omitempty" xml:"stageId,omitempty"`
	StartDate             *string                                       `json:"startDate,omitempty" xml:"startDate,omitempty"`
	StoryPoint            *string                                       `json:"storyPoint,omitempty" xml:"storyPoint,omitempty"`
	TagIds                []*string                                     `json:"tagIds,omitempty" xml:"tagIds,omitempty" type:"Repeated"`
	TaskId                *string                                       `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TaskListId            *string                                       `json:"taskListId,omitempty" xml:"taskListId,omitempty"`
	TaskflowstatusId      *string                                       `json:"taskflowstatusId,omitempty" xml:"taskflowstatusId,omitempty"`
	UniqueId              *string                                       `json:"uniqueId,omitempty" xml:"uniqueId,omitempty"`
	Updated               *string                                       `json:"updated,omitempty" xml:"updated,omitempty"`
	Visible               *string                                       `json:"visible,omitempty" xml:"visible,omitempty"`
}

func (s GetTaskByIdsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetTaskByIdsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetTaskByIdsResponseBodyResult) SetAccomplishTime(v string) *GetTaskByIdsResponseBodyResult {
	s.AccomplishTime = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetAncestorIds(v []*string) *GetTaskByIdsResponseBodyResult {
	s.AncestorIds = v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetContent(v string) *GetTaskByIdsResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetCreated(v string) *GetTaskByIdsResponseBodyResult {
	s.Created = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetCreatorId(v string) *GetTaskByIdsResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetCustomfields(v []*GetTaskByIdsResponseBodyResultCustomfields) *GetTaskByIdsResponseBodyResult {
	s.Customfields = v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetDueDate(v string) *GetTaskByIdsResponseBodyResult {
	s.DueDate = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetExecutorId(v string) *GetTaskByIdsResponseBodyResult {
	s.ExecutorId = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetInvolveMembers(v []*string) *GetTaskByIdsResponseBodyResult {
	s.InvolveMembers = v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetIsArchived(v bool) *GetTaskByIdsResponseBodyResult {
	s.IsArchived = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetIsDone(v bool) *GetTaskByIdsResponseBodyResult {
	s.IsDone = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetNote(v string) *GetTaskByIdsResponseBodyResult {
	s.Note = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetParentTaskId(v string) *GetTaskByIdsResponseBodyResult {
	s.ParentTaskId = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetPriority(v int32) *GetTaskByIdsResponseBodyResult {
	s.Priority = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetProjectId(v string) *GetTaskByIdsResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetRecurrence(v []*string) *GetTaskByIdsResponseBodyResult {
	s.Recurrence = v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetScenariofieldconfigId(v string) *GetTaskByIdsResponseBodyResult {
	s.ScenariofieldconfigId = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetSprintId(v string) *GetTaskByIdsResponseBodyResult {
	s.SprintId = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetStageId(v string) *GetTaskByIdsResponseBodyResult {
	s.StageId = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetStartDate(v string) *GetTaskByIdsResponseBodyResult {
	s.StartDate = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetStoryPoint(v string) *GetTaskByIdsResponseBodyResult {
	s.StoryPoint = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetTagIds(v []*string) *GetTaskByIdsResponseBodyResult {
	s.TagIds = v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetTaskId(v string) *GetTaskByIdsResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetTaskListId(v string) *GetTaskByIdsResponseBodyResult {
	s.TaskListId = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetTaskflowstatusId(v string) *GetTaskByIdsResponseBodyResult {
	s.TaskflowstatusId = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetUniqueId(v string) *GetTaskByIdsResponseBodyResult {
	s.UniqueId = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetUpdated(v string) *GetTaskByIdsResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResult) SetVisible(v string) *GetTaskByIdsResponseBodyResult {
	s.Visible = &v
	return s
}

type GetTaskByIdsResponseBodyResultCustomfields struct {
	CustomfieldId *string                                            `json:"customfieldId,omitempty" xml:"customfieldId,omitempty"`
	Type          *string                                            `json:"type,omitempty" xml:"type,omitempty"`
	Value         []*GetTaskByIdsResponseBodyResultCustomfieldsValue `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s GetTaskByIdsResponseBodyResultCustomfields) String() string {
	return tea.Prettify(s)
}

func (s GetTaskByIdsResponseBodyResultCustomfields) GoString() string {
	return s.String()
}

func (s *GetTaskByIdsResponseBodyResultCustomfields) SetCustomfieldId(v string) *GetTaskByIdsResponseBodyResultCustomfields {
	s.CustomfieldId = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResultCustomfields) SetType(v string) *GetTaskByIdsResponseBodyResultCustomfields {
	s.Type = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResultCustomfields) SetValue(v []*GetTaskByIdsResponseBodyResultCustomfieldsValue) *GetTaskByIdsResponseBodyResultCustomfields {
	s.Value = v
	return s
}

type GetTaskByIdsResponseBodyResultCustomfieldsValue struct {
	FieldvalueId *string `json:"fieldvalueId,omitempty" xml:"fieldvalueId,omitempty"`
	MetaString   *string `json:"metaString,omitempty" xml:"metaString,omitempty"`
	Title        *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetTaskByIdsResponseBodyResultCustomfieldsValue) String() string {
	return tea.Prettify(s)
}

func (s GetTaskByIdsResponseBodyResultCustomfieldsValue) GoString() string {
	return s.String()
}

func (s *GetTaskByIdsResponseBodyResultCustomfieldsValue) SetFieldvalueId(v string) *GetTaskByIdsResponseBodyResultCustomfieldsValue {
	s.FieldvalueId = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResultCustomfieldsValue) SetMetaString(v string) *GetTaskByIdsResponseBodyResultCustomfieldsValue {
	s.MetaString = &v
	return s
}

func (s *GetTaskByIdsResponseBodyResultCustomfieldsValue) SetTitle(v string) *GetTaskByIdsResponseBodyResultCustomfieldsValue {
	s.Title = &v
	return s
}

type GetTaskByIdsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTaskByIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTaskByIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskByIdsResponse) GoString() string {
	return s.String()
}

func (s *GetTaskByIdsResponse) SetHeaders(v map[string]*string) *GetTaskByIdsResponse {
	s.Headers = v
	return s
}

func (s *GetTaskByIdsResponse) SetStatusCode(v int32) *GetTaskByIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskByIdsResponse) SetBody(v *GetTaskByIdsResponseBody) *GetTaskByIdsResponse {
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTbOrgIdByDingOrgIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetTbOrgIdByDingOrgIdResponse) SetStatusCode(v int32) *GetTbOrgIdByDingOrgIdResponse {
	s.StatusCode = &v
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
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTbProjectGrayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetTbProjectGrayResponse) SetStatusCode(v int32) *GetTbProjectGrayResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTbProjectSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetTbProjectSourceResponse) SetStatusCode(v int32) *GetTbProjectSourceResponse {
	s.StatusCode = &v
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
	OptUserId *string `json:"optUserId,omitempty" xml:"optUserId,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTbUserIdByStaffIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetTbUserIdByStaffIdResponse) SetStatusCode(v int32) *GetTbUserIdByStaffIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTbUserIdByStaffIdResponse) SetBody(v *GetTbUserIdByStaffIdResponseBody) *GetTbUserIdByStaffIdResponse {
	s.Body = v
	return s
}

type GetUserJoinedProjectHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserJoinedProjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserJoinedProjectHeaders) GoString() string {
	return s.String()
}

func (s *GetUserJoinedProjectHeaders) SetCommonHeaders(v map[string]*string) *GetUserJoinedProjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserJoinedProjectHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserJoinedProjectHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserJoinedProjectRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetUserJoinedProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserJoinedProjectRequest) GoString() string {
	return s.String()
}

func (s *GetUserJoinedProjectRequest) SetMaxResults(v int32) *GetUserJoinedProjectRequest {
	s.MaxResults = &v
	return s
}

func (s *GetUserJoinedProjectRequest) SetNextToken(v string) *GetUserJoinedProjectRequest {
	s.NextToken = &v
	return s
}

type GetUserJoinedProjectResponseBody struct {
	NextToken  *string   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Result     []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int32    `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetUserJoinedProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserJoinedProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserJoinedProjectResponseBody) SetNextToken(v string) *GetUserJoinedProjectResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetUserJoinedProjectResponseBody) SetResult(v []*string) *GetUserJoinedProjectResponseBody {
	s.Result = v
	return s
}

func (s *GetUserJoinedProjectResponseBody) SetTotalCount(v int32) *GetUserJoinedProjectResponseBody {
	s.TotalCount = &v
	return s
}

type GetUserJoinedProjectResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserJoinedProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserJoinedProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserJoinedProjectResponse) GoString() string {
	return s.String()
}

func (s *GetUserJoinedProjectResponse) SetHeaders(v map[string]*string) *GetUserJoinedProjectResponse {
	s.Headers = v
	return s
}

func (s *GetUserJoinedProjectResponse) SetStatusCode(v int32) *GetUserJoinedProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserJoinedProjectResponse) SetBody(v *GetUserJoinedProjectResponseBody) *GetUserJoinedProjectResponse {
	s.Body = v
	return s
}

type QueryProjectHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryProjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectHeaders) GoString() string {
	return s.String()
}

func (s *QueryProjectHeaders) SetCommonHeaders(v map[string]*string) *QueryProjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryProjectHeaders) SetXAcsDingtalkAccessToken(v string) *QueryProjectHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryProjectRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ProjectIds *string `json:"projectIds,omitempty" xml:"projectIds,omitempty"`
	SourceId   *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
}

func (s QueryProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectRequest) GoString() string {
	return s.String()
}

func (s *QueryProjectRequest) SetMaxResults(v int32) *QueryProjectRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryProjectRequest) SetName(v string) *QueryProjectRequest {
	s.Name = &v
	return s
}

func (s *QueryProjectRequest) SetNextToken(v string) *QueryProjectRequest {
	s.NextToken = &v
	return s
}

func (s *QueryProjectRequest) SetProjectIds(v string) *QueryProjectRequest {
	s.ProjectIds = &v
	return s
}

func (s *QueryProjectRequest) SetSourceId(v string) *QueryProjectRequest {
	s.SourceId = &v
	return s
}

type QueryProjectResponseBody struct {
	Result []*QueryProjectResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectResponseBody) GoString() string {
	return s.String()
}

func (s *QueryProjectResponseBody) SetResult(v []*QueryProjectResponseBodyResult) *QueryProjectResponseBody {
	s.Result = v
	return s
}

type QueryProjectResponseBodyResult struct {
	Created        *string                                       `json:"created,omitempty" xml:"created,omitempty"`
	CreatorId      *string                                       `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Customfields   []*QueryProjectResponseBodyResultCustomfields `json:"customfields,omitempty" xml:"customfields,omitempty" type:"Repeated"`
	Description    *string                                       `json:"description,omitempty" xml:"description,omitempty"`
	EndDate        *string                                       `json:"endDate,omitempty" xml:"endDate,omitempty"`
	IsArchived     *bool                                         `json:"isArchived,omitempty" xml:"isArchived,omitempty"`
	IsSuspended    *bool                                         `json:"isSuspended,omitempty" xml:"isSuspended,omitempty"`
	IsTemplate     *bool                                         `json:"isTemplate,omitempty" xml:"isTemplate,omitempty"`
	Logo           *string                                       `json:"logo,omitempty" xml:"logo,omitempty"`
	Name           *string                                       `json:"name,omitempty" xml:"name,omitempty"`
	OrganizationId *string                                       `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	ProjectId      *string                                       `json:"projectId,omitempty" xml:"projectId,omitempty"`
	StartDate      *string                                       `json:"startDate,omitempty" xml:"startDate,omitempty"`
	UniqueIdPrefix *string                                       `json:"uniqueIdPrefix,omitempty" xml:"uniqueIdPrefix,omitempty"`
	Updated        *string                                       `json:"updated,omitempty" xml:"updated,omitempty"`
	Visibility     *string                                       `json:"visibility,omitempty" xml:"visibility,omitempty"`
}

func (s QueryProjectResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryProjectResponseBodyResult) SetCreated(v string) *QueryProjectResponseBodyResult {
	s.Created = &v
	return s
}

func (s *QueryProjectResponseBodyResult) SetCreatorId(v string) *QueryProjectResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *QueryProjectResponseBodyResult) SetCustomfields(v []*QueryProjectResponseBodyResultCustomfields) *QueryProjectResponseBodyResult {
	s.Customfields = v
	return s
}

func (s *QueryProjectResponseBodyResult) SetDescription(v string) *QueryProjectResponseBodyResult {
	s.Description = &v
	return s
}

func (s *QueryProjectResponseBodyResult) SetEndDate(v string) *QueryProjectResponseBodyResult {
	s.EndDate = &v
	return s
}

func (s *QueryProjectResponseBodyResult) SetIsArchived(v bool) *QueryProjectResponseBodyResult {
	s.IsArchived = &v
	return s
}

func (s *QueryProjectResponseBodyResult) SetIsSuspended(v bool) *QueryProjectResponseBodyResult {
	s.IsSuspended = &v
	return s
}

func (s *QueryProjectResponseBodyResult) SetIsTemplate(v bool) *QueryProjectResponseBodyResult {
	s.IsTemplate = &v
	return s
}

func (s *QueryProjectResponseBodyResult) SetLogo(v string) *QueryProjectResponseBodyResult {
	s.Logo = &v
	return s
}

func (s *QueryProjectResponseBodyResult) SetName(v string) *QueryProjectResponseBodyResult {
	s.Name = &v
	return s
}

func (s *QueryProjectResponseBodyResult) SetOrganizationId(v string) *QueryProjectResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *QueryProjectResponseBodyResult) SetProjectId(v string) *QueryProjectResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *QueryProjectResponseBodyResult) SetStartDate(v string) *QueryProjectResponseBodyResult {
	s.StartDate = &v
	return s
}

func (s *QueryProjectResponseBodyResult) SetUniqueIdPrefix(v string) *QueryProjectResponseBodyResult {
	s.UniqueIdPrefix = &v
	return s
}

func (s *QueryProjectResponseBodyResult) SetUpdated(v string) *QueryProjectResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *QueryProjectResponseBodyResult) SetVisibility(v string) *QueryProjectResponseBodyResult {
	s.Visibility = &v
	return s
}

type QueryProjectResponseBodyResultCustomfields struct {
	CustomfieldId *string                                            `json:"customfieldId,omitempty" xml:"customfieldId,omitempty"`
	Type          *string                                            `json:"type,omitempty" xml:"type,omitempty"`
	Value         []*QueryProjectResponseBodyResultCustomfieldsValue `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s QueryProjectResponseBodyResultCustomfields) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectResponseBodyResultCustomfields) GoString() string {
	return s.String()
}

func (s *QueryProjectResponseBodyResultCustomfields) SetCustomfieldId(v string) *QueryProjectResponseBodyResultCustomfields {
	s.CustomfieldId = &v
	return s
}

func (s *QueryProjectResponseBodyResultCustomfields) SetType(v string) *QueryProjectResponseBodyResultCustomfields {
	s.Type = &v
	return s
}

func (s *QueryProjectResponseBodyResultCustomfields) SetValue(v []*QueryProjectResponseBodyResultCustomfieldsValue) *QueryProjectResponseBodyResultCustomfields {
	s.Value = v
	return s
}

type QueryProjectResponseBodyResultCustomfieldsValue struct {
	FieldvalueId *string `json:"fieldvalueId,omitempty" xml:"fieldvalueId,omitempty"`
	MetaString   *string `json:"metaString,omitempty" xml:"metaString,omitempty"`
	Title        *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryProjectResponseBodyResultCustomfieldsValue) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectResponseBodyResultCustomfieldsValue) GoString() string {
	return s.String()
}

func (s *QueryProjectResponseBodyResultCustomfieldsValue) SetFieldvalueId(v string) *QueryProjectResponseBodyResultCustomfieldsValue {
	s.FieldvalueId = &v
	return s
}

func (s *QueryProjectResponseBodyResultCustomfieldsValue) SetMetaString(v string) *QueryProjectResponseBodyResultCustomfieldsValue {
	s.MetaString = &v
	return s
}

func (s *QueryProjectResponseBodyResultCustomfieldsValue) SetTitle(v string) *QueryProjectResponseBodyResultCustomfieldsValue {
	s.Title = &v
	return s
}

type QueryProjectResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectResponse) GoString() string {
	return s.String()
}

func (s *QueryProjectResponse) SetHeaders(v map[string]*string) *QueryProjectResponse {
	s.Headers = v
	return s
}

func (s *QueryProjectResponse) SetStatusCode(v int32) *QueryProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryProjectResponse) SetBody(v *QueryProjectResponseBody) *QueryProjectResponse {
	s.Body = v
	return s
}

type QueryTaskOfProjectHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryTaskOfProjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskOfProjectHeaders) GoString() string {
	return s.String()
}

func (s *QueryTaskOfProjectHeaders) SetCommonHeaders(v map[string]*string) *QueryTaskOfProjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryTaskOfProjectHeaders) SetXAcsDingtalkAccessToken(v string) *QueryTaskOfProjectHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryTaskOfProjectRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Query      *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s QueryTaskOfProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskOfProjectRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskOfProjectRequest) SetMaxResults(v int32) *QueryTaskOfProjectRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryTaskOfProjectRequest) SetNextToken(v string) *QueryTaskOfProjectRequest {
	s.NextToken = &v
	return s
}

func (s *QueryTaskOfProjectRequest) SetQuery(v string) *QueryTaskOfProjectRequest {
	s.Query = &v
	return s
}

type QueryTaskOfProjectResponseBody struct {
	NextToken  *string                                 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Result     []*QueryTaskOfProjectResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int32                                  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryTaskOfProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskOfProjectResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTaskOfProjectResponseBody) SetNextToken(v string) *QueryTaskOfProjectResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryTaskOfProjectResponseBody) SetResult(v []*QueryTaskOfProjectResponseBodyResult) *QueryTaskOfProjectResponseBody {
	s.Result = v
	return s
}

func (s *QueryTaskOfProjectResponseBody) SetTotalCount(v int32) *QueryTaskOfProjectResponseBody {
	s.TotalCount = &v
	return s
}

type QueryTaskOfProjectResponseBodyResult struct {
	Accomplished          *string                                             `json:"accomplished,omitempty" xml:"accomplished,omitempty"`
	AncestorIds           []*string                                           `json:"ancestorIds,omitempty" xml:"ancestorIds,omitempty" type:"Repeated"`
	Content               *string                                             `json:"content,omitempty" xml:"content,omitempty"`
	Created               *string                                             `json:"created,omitempty" xml:"created,omitempty"`
	CreatorId             *string                                             `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Customfields          []*QueryTaskOfProjectResponseBodyResultCustomfields `json:"customfields,omitempty" xml:"customfields,omitempty" type:"Repeated"`
	DueDate               *string                                             `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	ExecutorId            *string                                             `json:"executorId,omitempty" xml:"executorId,omitempty"`
	InvolveMembers        []*string                                           `json:"involveMembers,omitempty" xml:"involveMembers,omitempty" type:"Repeated"`
	IsArchived            *bool                                               `json:"isArchived,omitempty" xml:"isArchived,omitempty"`
	IsDeleted             *bool                                               `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	IsDone                *bool                                               `json:"isDone,omitempty" xml:"isDone,omitempty"`
	Labels                []*string                                           `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	Note                  *string                                             `json:"note,omitempty" xml:"note,omitempty"`
	Priority              *int64                                              `json:"priority,omitempty" xml:"priority,omitempty"`
	Progress              *int32                                              `json:"progress,omitempty" xml:"progress,omitempty"`
	ProjectId             *string                                             `json:"projectId,omitempty" xml:"projectId,omitempty"`
	ScenariofieldconfigId *string                                             `json:"scenariofieldconfigId,omitempty" xml:"scenariofieldconfigId,omitempty"`
	SprintId              *string                                             `json:"sprintId,omitempty" xml:"sprintId,omitempty"`
	StageId               *string                                             `json:"stageId,omitempty" xml:"stageId,omitempty"`
	StartDate             *string                                             `json:"startDate,omitempty" xml:"startDate,omitempty"`
	StoryPoint            *int32                                              `json:"storyPoint,omitempty" xml:"storyPoint,omitempty"`
	TagIds                []*string                                           `json:"tagIds,omitempty" xml:"tagIds,omitempty" type:"Repeated"`
	TaskId                *string                                             `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TaskflowstatusId      *string                                             `json:"taskflowstatusId,omitempty" xml:"taskflowstatusId,omitempty"`
	Updated               *string                                             `json:"updated,omitempty" xml:"updated,omitempty"`
	Visible               *string                                             `json:"visible,omitempty" xml:"visible,omitempty"`
}

func (s QueryTaskOfProjectResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskOfProjectResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryTaskOfProjectResponseBodyResult) SetAccomplished(v string) *QueryTaskOfProjectResponseBodyResult {
	s.Accomplished = &v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetAncestorIds(v []*string) *QueryTaskOfProjectResponseBodyResult {
	s.AncestorIds = v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetContent(v string) *QueryTaskOfProjectResponseBodyResult {
	s.Content = &v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetCreated(v string) *QueryTaskOfProjectResponseBodyResult {
	s.Created = &v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetCreatorId(v string) *QueryTaskOfProjectResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetCustomfields(v []*QueryTaskOfProjectResponseBodyResultCustomfields) *QueryTaskOfProjectResponseBodyResult {
	s.Customfields = v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetDueDate(v string) *QueryTaskOfProjectResponseBodyResult {
	s.DueDate = &v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetExecutorId(v string) *QueryTaskOfProjectResponseBodyResult {
	s.ExecutorId = &v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetInvolveMembers(v []*string) *QueryTaskOfProjectResponseBodyResult {
	s.InvolveMembers = v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetIsArchived(v bool) *QueryTaskOfProjectResponseBodyResult {
	s.IsArchived = &v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetIsDeleted(v bool) *QueryTaskOfProjectResponseBodyResult {
	s.IsDeleted = &v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetIsDone(v bool) *QueryTaskOfProjectResponseBodyResult {
	s.IsDone = &v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetLabels(v []*string) *QueryTaskOfProjectResponseBodyResult {
	s.Labels = v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetNote(v string) *QueryTaskOfProjectResponseBodyResult {
	s.Note = &v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetPriority(v int64) *QueryTaskOfProjectResponseBodyResult {
	s.Priority = &v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetProgress(v int32) *QueryTaskOfProjectResponseBodyResult {
	s.Progress = &v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetProjectId(v string) *QueryTaskOfProjectResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetScenariofieldconfigId(v string) *QueryTaskOfProjectResponseBodyResult {
	s.ScenariofieldconfigId = &v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetSprintId(v string) *QueryTaskOfProjectResponseBodyResult {
	s.SprintId = &v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetStageId(v string) *QueryTaskOfProjectResponseBodyResult {
	s.StageId = &v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetStartDate(v string) *QueryTaskOfProjectResponseBodyResult {
	s.StartDate = &v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetStoryPoint(v int32) *QueryTaskOfProjectResponseBodyResult {
	s.StoryPoint = &v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetTagIds(v []*string) *QueryTaskOfProjectResponseBodyResult {
	s.TagIds = v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetTaskId(v string) *QueryTaskOfProjectResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetTaskflowstatusId(v string) *QueryTaskOfProjectResponseBodyResult {
	s.TaskflowstatusId = &v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetUpdated(v string) *QueryTaskOfProjectResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *QueryTaskOfProjectResponseBodyResult) SetVisible(v string) *QueryTaskOfProjectResponseBodyResult {
	s.Visible = &v
	return s
}

type QueryTaskOfProjectResponseBodyResultCustomfields struct {
	CustomfieldId *string `json:"customfieldId,omitempty" xml:"customfieldId,omitempty"`
}

func (s QueryTaskOfProjectResponseBodyResultCustomfields) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskOfProjectResponseBodyResultCustomfields) GoString() string {
	return s.String()
}

func (s *QueryTaskOfProjectResponseBodyResultCustomfields) SetCustomfieldId(v string) *QueryTaskOfProjectResponseBodyResultCustomfields {
	s.CustomfieldId = &v
	return s
}

type QueryTaskOfProjectResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTaskOfProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTaskOfProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskOfProjectResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskOfProjectResponse) SetHeaders(v map[string]*string) *QueryTaskOfProjectResponse {
	s.Headers = v
	return s
}

func (s *QueryTaskOfProjectResponse) SetStatusCode(v int32) *QueryTaskOfProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTaskOfProjectResponse) SetBody(v *QueryTaskOfProjectResponseBody) *QueryTaskOfProjectResponse {
	s.Body = v
	return s
}

type SeachTaskStageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SeachTaskStageHeaders) String() string {
	return tea.Prettify(s)
}

func (s SeachTaskStageHeaders) GoString() string {
	return s.String()
}

func (s *SeachTaskStageHeaders) SetCommonHeaders(v map[string]*string) *SeachTaskStageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SeachTaskStageHeaders) SetXAcsDingtalkAccessToken(v string) *SeachTaskStageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SeachTaskStageRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Query      *string `json:"query,omitempty" xml:"query,omitempty"`
	StageIds   *string `json:"stageIds,omitempty" xml:"stageIds,omitempty"`
	TaskListId *string `json:"taskListId,omitempty" xml:"taskListId,omitempty"`
}

func (s SeachTaskStageRequest) String() string {
	return tea.Prettify(s)
}

func (s SeachTaskStageRequest) GoString() string {
	return s.String()
}

func (s *SeachTaskStageRequest) SetMaxResults(v int32) *SeachTaskStageRequest {
	s.MaxResults = &v
	return s
}

func (s *SeachTaskStageRequest) SetNextToken(v string) *SeachTaskStageRequest {
	s.NextToken = &v
	return s
}

func (s *SeachTaskStageRequest) SetQuery(v string) *SeachTaskStageRequest {
	s.Query = &v
	return s
}

func (s *SeachTaskStageRequest) SetStageIds(v string) *SeachTaskStageRequest {
	s.StageIds = &v
	return s
}

func (s *SeachTaskStageRequest) SetTaskListId(v string) *SeachTaskStageRequest {
	s.TaskListId = &v
	return s
}

type SeachTaskStageResponseBody struct {
	NextToken  *string                             `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Result     []*SeachTaskStageResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int32                              `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s SeachTaskStageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SeachTaskStageResponseBody) GoString() string {
	return s.String()
}

func (s *SeachTaskStageResponseBody) SetNextToken(v string) *SeachTaskStageResponseBody {
	s.NextToken = &v
	return s
}

func (s *SeachTaskStageResponseBody) SetResult(v []*SeachTaskStageResponseBodyResult) *SeachTaskStageResponseBody {
	s.Result = v
	return s
}

func (s *SeachTaskStageResponseBody) SetTotalCount(v int32) *SeachTaskStageResponseBody {
	s.TotalCount = &v
	return s
}

type SeachTaskStageResponseBodyResult struct {
	Created     *string `json:"created,omitempty" xml:"created,omitempty"`
	CreatorId   *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	ProjectId   *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	TaskListId  *string `json:"taskListId,omitempty" xml:"taskListId,omitempty"`
	TaskStageId *string `json:"taskStageId,omitempty" xml:"taskStageId,omitempty"`
	Updated     *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s SeachTaskStageResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SeachTaskStageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SeachTaskStageResponseBodyResult) SetCreated(v string) *SeachTaskStageResponseBodyResult {
	s.Created = &v
	return s
}

func (s *SeachTaskStageResponseBodyResult) SetCreatorId(v string) *SeachTaskStageResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *SeachTaskStageResponseBodyResult) SetDescription(v string) *SeachTaskStageResponseBodyResult {
	s.Description = &v
	return s
}

func (s *SeachTaskStageResponseBodyResult) SetName(v string) *SeachTaskStageResponseBodyResult {
	s.Name = &v
	return s
}

func (s *SeachTaskStageResponseBodyResult) SetProjectId(v string) *SeachTaskStageResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *SeachTaskStageResponseBodyResult) SetTaskListId(v string) *SeachTaskStageResponseBodyResult {
	s.TaskListId = &v
	return s
}

func (s *SeachTaskStageResponseBodyResult) SetTaskStageId(v string) *SeachTaskStageResponseBodyResult {
	s.TaskStageId = &v
	return s
}

func (s *SeachTaskStageResponseBodyResult) SetUpdated(v string) *SeachTaskStageResponseBodyResult {
	s.Updated = &v
	return s
}

type SeachTaskStageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SeachTaskStageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SeachTaskStageResponse) String() string {
	return tea.Prettify(s)
}

func (s SeachTaskStageResponse) GoString() string {
	return s.String()
}

func (s *SeachTaskStageResponse) SetHeaders(v map[string]*string) *SeachTaskStageResponse {
	s.Headers = v
	return s
}

func (s *SeachTaskStageResponse) SetStatusCode(v int32) *SeachTaskStageResponse {
	s.StatusCode = &v
	return s
}

func (s *SeachTaskStageResponse) SetBody(v *SeachTaskStageResponseBody) *SeachTaskStageResponse {
	s.Body = v
	return s
}

type SearchOranizationCustomfieldHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchOranizationCustomfieldHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchOranizationCustomfieldHeaders) GoString() string {
	return s.String()
}

func (s *SearchOranizationCustomfieldHeaders) SetCommonHeaders(v map[string]*string) *SearchOranizationCustomfieldHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchOranizationCustomfieldHeaders) SetXAcsDingtalkAccessToken(v string) *SearchOranizationCustomfieldHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchOranizationCustomfieldRequest struct {
	CustomfieldIds *string `json:"customfieldIds,omitempty" xml:"customfieldIds,omitempty"`
	InstanceIds    *string `json:"instanceIds,omitempty" xml:"instanceIds,omitempty"`
	MaxResults     *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken      *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ProjectIds     *string `json:"projectIds,omitempty" xml:"projectIds,omitempty"`
	Query          *string `json:"query,omitempty" xml:"query,omitempty"`
	Scope          *string `json:"scope,omitempty" xml:"scope,omitempty"`
}

func (s SearchOranizationCustomfieldRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchOranizationCustomfieldRequest) GoString() string {
	return s.String()
}

func (s *SearchOranizationCustomfieldRequest) SetCustomfieldIds(v string) *SearchOranizationCustomfieldRequest {
	s.CustomfieldIds = &v
	return s
}

func (s *SearchOranizationCustomfieldRequest) SetInstanceIds(v string) *SearchOranizationCustomfieldRequest {
	s.InstanceIds = &v
	return s
}

func (s *SearchOranizationCustomfieldRequest) SetMaxResults(v int32) *SearchOranizationCustomfieldRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchOranizationCustomfieldRequest) SetNextToken(v string) *SearchOranizationCustomfieldRequest {
	s.NextToken = &v
	return s
}

func (s *SearchOranizationCustomfieldRequest) SetProjectIds(v string) *SearchOranizationCustomfieldRequest {
	s.ProjectIds = &v
	return s
}

func (s *SearchOranizationCustomfieldRequest) SetQuery(v string) *SearchOranizationCustomfieldRequest {
	s.Query = &v
	return s
}

func (s *SearchOranizationCustomfieldRequest) SetScope(v string) *SearchOranizationCustomfieldRequest {
	s.Scope = &v
	return s
}

type SearchOranizationCustomfieldResponseBody struct {
	NextToken  *string                                           `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Result     []*SearchOranizationCustomfieldResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int32                                            `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s SearchOranizationCustomfieldResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchOranizationCustomfieldResponseBody) GoString() string {
	return s.String()
}

func (s *SearchOranizationCustomfieldResponseBody) SetNextToken(v string) *SearchOranizationCustomfieldResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchOranizationCustomfieldResponseBody) SetResult(v []*SearchOranizationCustomfieldResponseBodyResult) *SearchOranizationCustomfieldResponseBody {
	s.Result = v
	return s
}

func (s *SearchOranizationCustomfieldResponseBody) SetTotalCount(v int32) *SearchOranizationCustomfieldResponseBody {
	s.TotalCount = &v
	return s
}

type SearchOranizationCustomfieldResponseBodyResult struct {
	AdvancedCustomfield *SearchOranizationCustomfieldResponseBodyResultAdvancedCustomfield `json:"advancedCustomfield,omitempty" xml:"advancedCustomfield,omitempty" type:"Struct"`
	Choices             []*SearchOranizationCustomfieldResponseBodyResultChoices           `json:"choices,omitempty" xml:"choices,omitempty" type:"Repeated"`
	Created             *string                                                            `json:"created,omitempty" xml:"created,omitempty"`
	CreatorId           *string                                                            `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	CustomfieldsId      *string                                                            `json:"customfieldsId,omitempty" xml:"customfieldsId,omitempty"`
	Name                *string                                                            `json:"name,omitempty" xml:"name,omitempty"`
	Payload             map[string]interface{}                                             `json:"payload,omitempty" xml:"payload,omitempty"`
	Type                *string                                                            `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SearchOranizationCustomfieldResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SearchOranizationCustomfieldResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SearchOranizationCustomfieldResponseBodyResult) SetAdvancedCustomfield(v *SearchOranizationCustomfieldResponseBodyResultAdvancedCustomfield) *SearchOranizationCustomfieldResponseBodyResult {
	s.AdvancedCustomfield = v
	return s
}

func (s *SearchOranizationCustomfieldResponseBodyResult) SetChoices(v []*SearchOranizationCustomfieldResponseBodyResultChoices) *SearchOranizationCustomfieldResponseBodyResult {
	s.Choices = v
	return s
}

func (s *SearchOranizationCustomfieldResponseBodyResult) SetCreated(v string) *SearchOranizationCustomfieldResponseBodyResult {
	s.Created = &v
	return s
}

func (s *SearchOranizationCustomfieldResponseBodyResult) SetCreatorId(v string) *SearchOranizationCustomfieldResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *SearchOranizationCustomfieldResponseBodyResult) SetCustomfieldsId(v string) *SearchOranizationCustomfieldResponseBodyResult {
	s.CustomfieldsId = &v
	return s
}

func (s *SearchOranizationCustomfieldResponseBodyResult) SetName(v string) *SearchOranizationCustomfieldResponseBodyResult {
	s.Name = &v
	return s
}

func (s *SearchOranizationCustomfieldResponseBodyResult) SetPayload(v map[string]interface{}) *SearchOranizationCustomfieldResponseBodyResult {
	s.Payload = v
	return s
}

func (s *SearchOranizationCustomfieldResponseBodyResult) SetType(v string) *SearchOranizationCustomfieldResponseBodyResult {
	s.Type = &v
	return s
}

type SearchOranizationCustomfieldResponseBodyResultAdvancedCustomfield struct {
	AdvancedCustomfieldId *string `json:"advancedCustomfieldId,omitempty" xml:"advancedCustomfieldId,omitempty"`
	Name                  *string `json:"name,omitempty" xml:"name,omitempty"`
	ObjectType            *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s SearchOranizationCustomfieldResponseBodyResultAdvancedCustomfield) String() string {
	return tea.Prettify(s)
}

func (s SearchOranizationCustomfieldResponseBodyResultAdvancedCustomfield) GoString() string {
	return s.String()
}

func (s *SearchOranizationCustomfieldResponseBodyResultAdvancedCustomfield) SetAdvancedCustomfieldId(v string) *SearchOranizationCustomfieldResponseBodyResultAdvancedCustomfield {
	s.AdvancedCustomfieldId = &v
	return s
}

func (s *SearchOranizationCustomfieldResponseBodyResultAdvancedCustomfield) SetName(v string) *SearchOranizationCustomfieldResponseBodyResultAdvancedCustomfield {
	s.Name = &v
	return s
}

func (s *SearchOranizationCustomfieldResponseBodyResultAdvancedCustomfield) SetObjectType(v string) *SearchOranizationCustomfieldResponseBodyResultAdvancedCustomfield {
	s.ObjectType = &v
	return s
}

type SearchOranizationCustomfieldResponseBodyResultChoices struct {
	ChoiceId *string `json:"choiceId,omitempty" xml:"choiceId,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s SearchOranizationCustomfieldResponseBodyResultChoices) String() string {
	return tea.Prettify(s)
}

func (s SearchOranizationCustomfieldResponseBodyResultChoices) GoString() string {
	return s.String()
}

func (s *SearchOranizationCustomfieldResponseBodyResultChoices) SetChoiceId(v string) *SearchOranizationCustomfieldResponseBodyResultChoices {
	s.ChoiceId = &v
	return s
}

func (s *SearchOranizationCustomfieldResponseBodyResultChoices) SetValue(v string) *SearchOranizationCustomfieldResponseBodyResultChoices {
	s.Value = &v
	return s
}

type SearchOranizationCustomfieldResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchOranizationCustomfieldResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchOranizationCustomfieldResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchOranizationCustomfieldResponse) GoString() string {
	return s.String()
}

func (s *SearchOranizationCustomfieldResponse) SetHeaders(v map[string]*string) *SearchOranizationCustomfieldResponse {
	s.Headers = v
	return s
}

func (s *SearchOranizationCustomfieldResponse) SetStatusCode(v int32) *SearchOranizationCustomfieldResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchOranizationCustomfieldResponse) SetBody(v *SearchOranizationCustomfieldResponseBody) *SearchOranizationCustomfieldResponse {
	s.Body = v
	return s
}

type SearchProjectCustomfieldHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchProjectCustomfieldHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectCustomfieldHeaders) GoString() string {
	return s.String()
}

func (s *SearchProjectCustomfieldHeaders) SetCommonHeaders(v map[string]*string) *SearchProjectCustomfieldHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchProjectCustomfieldHeaders) SetXAcsDingtalkAccessToken(v string) *SearchProjectCustomfieldHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchProjectCustomfieldRequest struct {
	CustomfieldIds        *string `json:"customfieldIds,omitempty" xml:"customfieldIds,omitempty"`
	InstanceIds           *string `json:"instanceIds,omitempty" xml:"instanceIds,omitempty"`
	MaxResults            *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken             *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Query                 *string `json:"query,omitempty" xml:"query,omitempty"`
	ScenariofieldconfigId *string `json:"scenariofieldconfigId,omitempty" xml:"scenariofieldconfigId,omitempty"`
	Scope                 *string `json:"scope,omitempty" xml:"scope,omitempty"`
}

func (s SearchProjectCustomfieldRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectCustomfieldRequest) GoString() string {
	return s.String()
}

func (s *SearchProjectCustomfieldRequest) SetCustomfieldIds(v string) *SearchProjectCustomfieldRequest {
	s.CustomfieldIds = &v
	return s
}

func (s *SearchProjectCustomfieldRequest) SetInstanceIds(v string) *SearchProjectCustomfieldRequest {
	s.InstanceIds = &v
	return s
}

func (s *SearchProjectCustomfieldRequest) SetMaxResults(v int32) *SearchProjectCustomfieldRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchProjectCustomfieldRequest) SetNextToken(v string) *SearchProjectCustomfieldRequest {
	s.NextToken = &v
	return s
}

func (s *SearchProjectCustomfieldRequest) SetQuery(v string) *SearchProjectCustomfieldRequest {
	s.Query = &v
	return s
}

func (s *SearchProjectCustomfieldRequest) SetScenariofieldconfigId(v string) *SearchProjectCustomfieldRequest {
	s.ScenariofieldconfigId = &v
	return s
}

func (s *SearchProjectCustomfieldRequest) SetScope(v string) *SearchProjectCustomfieldRequest {
	s.Scope = &v
	return s
}

type SearchProjectCustomfieldResponseBody struct {
	NextToken  *string                                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Result     []*SearchProjectCustomfieldResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int32                                        `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s SearchProjectCustomfieldResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectCustomfieldResponseBody) GoString() string {
	return s.String()
}

func (s *SearchProjectCustomfieldResponseBody) SetNextToken(v string) *SearchProjectCustomfieldResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchProjectCustomfieldResponseBody) SetResult(v []*SearchProjectCustomfieldResponseBodyResult) *SearchProjectCustomfieldResponseBody {
	s.Result = v
	return s
}

func (s *SearchProjectCustomfieldResponseBody) SetTotalCount(v int32) *SearchProjectCustomfieldResponseBody {
	s.TotalCount = &v
	return s
}

type SearchProjectCustomfieldResponseBodyResult struct {
	AdvancedCustomfield *SearchProjectCustomfieldResponseBodyResultAdvancedCustomfield `json:"advancedCustomfield,omitempty" xml:"advancedCustomfield,omitempty" type:"Struct"`
	BoundToObjectId     *string                                                        `json:"boundToObjectId,omitempty" xml:"boundToObjectId,omitempty"`
	Choices             []*SearchProjectCustomfieldResponseBodyResultChoices           `json:"choices,omitempty" xml:"choices,omitempty" type:"Repeated"`
	Created             *string                                                        `json:"created,omitempty" xml:"created,omitempty"`
	CreatorId           *string                                                        `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	CustomfiledId       *string                                                        `json:"customfiledId,omitempty" xml:"customfiledId,omitempty"`
	Name                *string                                                        `json:"name,omitempty" xml:"name,omitempty"`
	OriginalId          *string                                                        `json:"originalId,omitempty" xml:"originalId,omitempty"`
	Payload             map[string]interface{}                                         `json:"payload,omitempty" xml:"payload,omitempty"`
	Type                *string                                                        `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SearchProjectCustomfieldResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectCustomfieldResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SearchProjectCustomfieldResponseBodyResult) SetAdvancedCustomfield(v *SearchProjectCustomfieldResponseBodyResultAdvancedCustomfield) *SearchProjectCustomfieldResponseBodyResult {
	s.AdvancedCustomfield = v
	return s
}

func (s *SearchProjectCustomfieldResponseBodyResult) SetBoundToObjectId(v string) *SearchProjectCustomfieldResponseBodyResult {
	s.BoundToObjectId = &v
	return s
}

func (s *SearchProjectCustomfieldResponseBodyResult) SetChoices(v []*SearchProjectCustomfieldResponseBodyResultChoices) *SearchProjectCustomfieldResponseBodyResult {
	s.Choices = v
	return s
}

func (s *SearchProjectCustomfieldResponseBodyResult) SetCreated(v string) *SearchProjectCustomfieldResponseBodyResult {
	s.Created = &v
	return s
}

func (s *SearchProjectCustomfieldResponseBodyResult) SetCreatorId(v string) *SearchProjectCustomfieldResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *SearchProjectCustomfieldResponseBodyResult) SetCustomfiledId(v string) *SearchProjectCustomfieldResponseBodyResult {
	s.CustomfiledId = &v
	return s
}

func (s *SearchProjectCustomfieldResponseBodyResult) SetName(v string) *SearchProjectCustomfieldResponseBodyResult {
	s.Name = &v
	return s
}

func (s *SearchProjectCustomfieldResponseBodyResult) SetOriginalId(v string) *SearchProjectCustomfieldResponseBodyResult {
	s.OriginalId = &v
	return s
}

func (s *SearchProjectCustomfieldResponseBodyResult) SetPayload(v map[string]interface{}) *SearchProjectCustomfieldResponseBodyResult {
	s.Payload = v
	return s
}

func (s *SearchProjectCustomfieldResponseBodyResult) SetType(v string) *SearchProjectCustomfieldResponseBodyResult {
	s.Type = &v
	return s
}

type SearchProjectCustomfieldResponseBodyResultAdvancedCustomfield struct {
	AdvancedCustomfieldId *string `json:"advancedCustomfieldId,omitempty" xml:"advancedCustomfieldId,omitempty"`
	Name                  *string `json:"name,omitempty" xml:"name,omitempty"`
	ObjectType            *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s SearchProjectCustomfieldResponseBodyResultAdvancedCustomfield) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectCustomfieldResponseBodyResultAdvancedCustomfield) GoString() string {
	return s.String()
}

func (s *SearchProjectCustomfieldResponseBodyResultAdvancedCustomfield) SetAdvancedCustomfieldId(v string) *SearchProjectCustomfieldResponseBodyResultAdvancedCustomfield {
	s.AdvancedCustomfieldId = &v
	return s
}

func (s *SearchProjectCustomfieldResponseBodyResultAdvancedCustomfield) SetName(v string) *SearchProjectCustomfieldResponseBodyResultAdvancedCustomfield {
	s.Name = &v
	return s
}

func (s *SearchProjectCustomfieldResponseBodyResultAdvancedCustomfield) SetObjectType(v string) *SearchProjectCustomfieldResponseBodyResultAdvancedCustomfield {
	s.ObjectType = &v
	return s
}

type SearchProjectCustomfieldResponseBodyResultChoices struct {
	ChoiceId *string `json:"choiceId,omitempty" xml:"choiceId,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s SearchProjectCustomfieldResponseBodyResultChoices) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectCustomfieldResponseBodyResultChoices) GoString() string {
	return s.String()
}

func (s *SearchProjectCustomfieldResponseBodyResultChoices) SetChoiceId(v string) *SearchProjectCustomfieldResponseBodyResultChoices {
	s.ChoiceId = &v
	return s
}

func (s *SearchProjectCustomfieldResponseBodyResultChoices) SetValue(v string) *SearchProjectCustomfieldResponseBodyResultChoices {
	s.Value = &v
	return s
}

type SearchProjectCustomfieldResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchProjectCustomfieldResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchProjectCustomfieldResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectCustomfieldResponse) GoString() string {
	return s.String()
}

func (s *SearchProjectCustomfieldResponse) SetHeaders(v map[string]*string) *SearchProjectCustomfieldResponse {
	s.Headers = v
	return s
}

func (s *SearchProjectCustomfieldResponse) SetStatusCode(v int32) *SearchProjectCustomfieldResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchProjectCustomfieldResponse) SetBody(v *SearchProjectCustomfieldResponseBody) *SearchProjectCustomfieldResponse {
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
	Created     *string `json:"created,omitempty" xml:"created,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Id          *string `json:"id,omitempty" xml:"id,omitempty"`
	IsDeleted   *bool   `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	IsDemo      *bool   `json:"isDemo,omitempty" xml:"isDemo,omitempty"`
	Logo        *string `json:"logo,omitempty" xml:"logo,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Updated     *string `json:"updated,omitempty" xml:"updated,omitempty"`
	Visible     *string `json:"visible,omitempty" xml:"visible,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchProjectTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SearchProjectTemplateResponse) SetStatusCode(v int32) *SearchProjectTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchProjectTemplateResponse) SetBody(v *SearchProjectTemplateResponseBody) *SearchProjectTemplateResponse {
	s.Body = v
	return s
}

type SearchTaskFlowHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchTaskFlowHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchTaskFlowHeaders) GoString() string {
	return s.String()
}

func (s *SearchTaskFlowHeaders) SetCommonHeaders(v map[string]*string) *SearchTaskFlowHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchTaskFlowHeaders) SetXAcsDingtalkAccessToken(v string) *SearchTaskFlowHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchTaskFlowRequest struct {
	MaxResults  *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken   *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Query       *string `json:"query,omitempty" xml:"query,omitempty"`
	TaskflowIds *string `json:"taskflowIds,omitempty" xml:"taskflowIds,omitempty"`
}

func (s SearchTaskFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTaskFlowRequest) GoString() string {
	return s.String()
}

func (s *SearchTaskFlowRequest) SetMaxResults(v int32) *SearchTaskFlowRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchTaskFlowRequest) SetNextToken(v string) *SearchTaskFlowRequest {
	s.NextToken = &v
	return s
}

func (s *SearchTaskFlowRequest) SetQuery(v string) *SearchTaskFlowRequest {
	s.Query = &v
	return s
}

func (s *SearchTaskFlowRequest) SetTaskflowIds(v string) *SearchTaskFlowRequest {
	s.TaskflowIds = &v
	return s
}

type SearchTaskFlowResponseBody struct {
	Result []*SearchTaskFlowResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s SearchTaskFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTaskFlowResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTaskFlowResponseBody) SetResult(v []*SearchTaskFlowResponseBodyResult) *SearchTaskFlowResponseBody {
	s.Result = v
	return s
}

type SearchTaskFlowResponseBodyResult struct {
	BoundToObjectId   *string `json:"boundToObjectId,omitempty" xml:"boundToObjectId,omitempty"`
	BoundToObjectType *string `json:"boundToObjectType,omitempty" xml:"boundToObjectType,omitempty"`
	Created           *string `json:"created,omitempty" xml:"created,omitempty"`
	CreatorId         *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	IsDeleted         *bool   `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	TaskflowId        *string `json:"taskflowId,omitempty" xml:"taskflowId,omitempty"`
	Updated           *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s SearchTaskFlowResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SearchTaskFlowResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SearchTaskFlowResponseBodyResult) SetBoundToObjectId(v string) *SearchTaskFlowResponseBodyResult {
	s.BoundToObjectId = &v
	return s
}

func (s *SearchTaskFlowResponseBodyResult) SetBoundToObjectType(v string) *SearchTaskFlowResponseBodyResult {
	s.BoundToObjectType = &v
	return s
}

func (s *SearchTaskFlowResponseBodyResult) SetCreated(v string) *SearchTaskFlowResponseBodyResult {
	s.Created = &v
	return s
}

func (s *SearchTaskFlowResponseBodyResult) SetCreatorId(v string) *SearchTaskFlowResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *SearchTaskFlowResponseBodyResult) SetIsDeleted(v bool) *SearchTaskFlowResponseBodyResult {
	s.IsDeleted = &v
	return s
}

func (s *SearchTaskFlowResponseBodyResult) SetName(v string) *SearchTaskFlowResponseBodyResult {
	s.Name = &v
	return s
}

func (s *SearchTaskFlowResponseBodyResult) SetTaskflowId(v string) *SearchTaskFlowResponseBodyResult {
	s.TaskflowId = &v
	return s
}

func (s *SearchTaskFlowResponseBodyResult) SetUpdated(v string) *SearchTaskFlowResponseBodyResult {
	s.Updated = &v
	return s
}

type SearchTaskFlowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchTaskFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchTaskFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTaskFlowResponse) GoString() string {
	return s.String()
}

func (s *SearchTaskFlowResponse) SetHeaders(v map[string]*string) *SearchTaskFlowResponse {
	s.Headers = v
	return s
}

func (s *SearchTaskFlowResponse) SetStatusCode(v int32) *SearchTaskFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTaskFlowResponse) SetBody(v *SearchTaskFlowResponseBody) *SearchTaskFlowResponse {
	s.Body = v
	return s
}

type SearchTaskListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchTaskListHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchTaskListHeaders) GoString() string {
	return s.String()
}

func (s *SearchTaskListHeaders) SetCommonHeaders(v map[string]*string) *SearchTaskListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchTaskListHeaders) SetXAcsDingtalkAccessToken(v string) *SearchTaskListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchTaskListRequest struct {
	MaxResults  *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken   *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Query       *string `json:"query,omitempty" xml:"query,omitempty"`
	TaskListIds *string `json:"taskListIds,omitempty" xml:"taskListIds,omitempty"`
}

func (s SearchTaskListRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTaskListRequest) GoString() string {
	return s.String()
}

func (s *SearchTaskListRequest) SetMaxResults(v int32) *SearchTaskListRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchTaskListRequest) SetNextToken(v string) *SearchTaskListRequest {
	s.NextToken = &v
	return s
}

func (s *SearchTaskListRequest) SetQuery(v string) *SearchTaskListRequest {
	s.Query = &v
	return s
}

func (s *SearchTaskListRequest) SetTaskListIds(v string) *SearchTaskListRequest {
	s.TaskListIds = &v
	return s
}

type SearchTaskListResponseBody struct {
	NextToken  *string                             `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Result     []*SearchTaskListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int32                              `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s SearchTaskListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTaskListResponseBody) SetNextToken(v string) *SearchTaskListResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchTaskListResponseBody) SetResult(v []*SearchTaskListResponseBodyResult) *SearchTaskListResponseBody {
	s.Result = v
	return s
}

func (s *SearchTaskListResponseBody) SetTotalCount(v int32) *SearchTaskListResponseBody {
	s.TotalCount = &v
	return s
}

type SearchTaskListResponseBodyResult struct {
	Created     *string `json:"created,omitempty" xml:"created,omitempty"`
	CreatorId   *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	ProjectId   *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	TaskId      *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
	Updated     *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s SearchTaskListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SearchTaskListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SearchTaskListResponseBodyResult) SetCreated(v string) *SearchTaskListResponseBodyResult {
	s.Created = &v
	return s
}

func (s *SearchTaskListResponseBodyResult) SetCreatorId(v string) *SearchTaskListResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *SearchTaskListResponseBodyResult) SetDescription(v string) *SearchTaskListResponseBodyResult {
	s.Description = &v
	return s
}

func (s *SearchTaskListResponseBodyResult) SetProjectId(v string) *SearchTaskListResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *SearchTaskListResponseBodyResult) SetTaskId(v string) *SearchTaskListResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *SearchTaskListResponseBodyResult) SetTitle(v string) *SearchTaskListResponseBodyResult {
	s.Title = &v
	return s
}

func (s *SearchTaskListResponseBodyResult) SetUpdated(v string) *SearchTaskListResponseBodyResult {
	s.Updated = &v
	return s
}

type SearchTaskListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchTaskListResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTaskListResponse) GoString() string {
	return s.String()
}

func (s *SearchTaskListResponse) SetHeaders(v map[string]*string) *SearchTaskListResponse {
	s.Headers = v
	return s
}

func (s *SearchTaskListResponse) SetStatusCode(v int32) *SearchTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTaskListResponse) SetBody(v *SearchTaskListResponseBody) *SearchTaskListResponse {
	s.Body = v
	return s
}

type SearchTaskflowStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchTaskflowStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchTaskflowStatusHeaders) GoString() string {
	return s.String()
}

func (s *SearchTaskflowStatusHeaders) SetCommonHeaders(v map[string]*string) *SearchTaskflowStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchTaskflowStatusHeaders) SetXAcsDingtalkAccessToken(v string) *SearchTaskflowStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchTaskflowStatusRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Query      *string `json:"query,omitempty" xml:"query,omitempty"`
	TfIds      *string `json:"tfIds,omitempty" xml:"tfIds,omitempty"`
	TfsIds     *string `json:"tfsIds,omitempty" xml:"tfsIds,omitempty"`
}

func (s SearchTaskflowStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTaskflowStatusRequest) GoString() string {
	return s.String()
}

func (s *SearchTaskflowStatusRequest) SetMaxResults(v int32) *SearchTaskflowStatusRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchTaskflowStatusRequest) SetNextToken(v string) *SearchTaskflowStatusRequest {
	s.NextToken = &v
	return s
}

func (s *SearchTaskflowStatusRequest) SetQuery(v string) *SearchTaskflowStatusRequest {
	s.Query = &v
	return s
}

func (s *SearchTaskflowStatusRequest) SetTfIds(v string) *SearchTaskflowStatusRequest {
	s.TfIds = &v
	return s
}

func (s *SearchTaskflowStatusRequest) SetTfsIds(v string) *SearchTaskflowStatusRequest {
	s.TfsIds = &v
	return s
}

type SearchTaskflowStatusResponseBody struct {
	Result []*SearchTaskflowStatusResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s SearchTaskflowStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTaskflowStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTaskflowStatusResponseBody) SetResult(v []*SearchTaskflowStatusResponseBodyResult) *SearchTaskflowStatusResponseBody {
	s.Result = v
	return s
}

type SearchTaskflowStatusResponseBodyResult struct {
	Created                     *string   `json:"created,omitempty" xml:"created,omitempty"`
	CreatorId                   *string   `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	IsDeleted                   *bool     `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	IsTaskflowstatusruleexector *bool     `json:"isTaskflowstatusruleexector,omitempty" xml:"isTaskflowstatusruleexector,omitempty"`
	Kind                        *string   `json:"kind,omitempty" xml:"kind,omitempty"`
	Name                        *string   `json:"name,omitempty" xml:"name,omitempty"`
	Pos                         *int32    `json:"pos,omitempty" xml:"pos,omitempty"`
	RejectStatusIds             []*string `json:"rejectStatusIds,omitempty" xml:"rejectStatusIds,omitempty" type:"Repeated"`
	TaskflowId                  *string   `json:"taskflowId,omitempty" xml:"taskflowId,omitempty"`
	TaskflowStatusId            *string   `json:"taskflowStatusId,omitempty" xml:"taskflowStatusId,omitempty"`
	Updated                     *string   `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s SearchTaskflowStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SearchTaskflowStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SearchTaskflowStatusResponseBodyResult) SetCreated(v string) *SearchTaskflowStatusResponseBodyResult {
	s.Created = &v
	return s
}

func (s *SearchTaskflowStatusResponseBodyResult) SetCreatorId(v string) *SearchTaskflowStatusResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *SearchTaskflowStatusResponseBodyResult) SetIsDeleted(v bool) *SearchTaskflowStatusResponseBodyResult {
	s.IsDeleted = &v
	return s
}

func (s *SearchTaskflowStatusResponseBodyResult) SetIsTaskflowstatusruleexector(v bool) *SearchTaskflowStatusResponseBodyResult {
	s.IsTaskflowstatusruleexector = &v
	return s
}

func (s *SearchTaskflowStatusResponseBodyResult) SetKind(v string) *SearchTaskflowStatusResponseBodyResult {
	s.Kind = &v
	return s
}

func (s *SearchTaskflowStatusResponseBodyResult) SetName(v string) *SearchTaskflowStatusResponseBodyResult {
	s.Name = &v
	return s
}

func (s *SearchTaskflowStatusResponseBodyResult) SetPos(v int32) *SearchTaskflowStatusResponseBodyResult {
	s.Pos = &v
	return s
}

func (s *SearchTaskflowStatusResponseBodyResult) SetRejectStatusIds(v []*string) *SearchTaskflowStatusResponseBodyResult {
	s.RejectStatusIds = v
	return s
}

func (s *SearchTaskflowStatusResponseBodyResult) SetTaskflowId(v string) *SearchTaskflowStatusResponseBodyResult {
	s.TaskflowId = &v
	return s
}

func (s *SearchTaskflowStatusResponseBodyResult) SetTaskflowStatusId(v string) *SearchTaskflowStatusResponseBodyResult {
	s.TaskflowStatusId = &v
	return s
}

func (s *SearchTaskflowStatusResponseBodyResult) SetUpdated(v string) *SearchTaskflowStatusResponseBodyResult {
	s.Updated = &v
	return s
}

type SearchTaskflowStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchTaskflowStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchTaskflowStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTaskflowStatusResponse) GoString() string {
	return s.String()
}

func (s *SearchTaskflowStatusResponse) SetHeaders(v map[string]*string) *SearchTaskflowStatusResponse {
	s.Headers = v
	return s
}

func (s *SearchTaskflowStatusResponse) SetStatusCode(v int32) *SearchTaskflowStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTaskflowStatusResponse) SetBody(v *SearchTaskflowStatusResponseBody) *SearchTaskflowStatusResponse {
	s.Body = v
	return s
}

type SearchUserTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchUserTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchUserTaskHeaders) GoString() string {
	return s.String()
}

func (s *SearchUserTaskHeaders) SetCommonHeaders(v map[string]*string) *SearchUserTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchUserTaskHeaders) SetXAcsDingtalkAccessToken(v string) *SearchUserTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchUserTaskRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Tql        *string `json:"tql,omitempty" xml:"tql,omitempty"`
}

func (s SearchUserTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchUserTaskRequest) GoString() string {
	return s.String()
}

func (s *SearchUserTaskRequest) SetMaxResults(v int32) *SearchUserTaskRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchUserTaskRequest) SetNextToken(v string) *SearchUserTaskRequest {
	s.NextToken = &v
	return s
}

func (s *SearchUserTaskRequest) SetTql(v string) *SearchUserTaskRequest {
	s.Tql = &v
	return s
}

type SearchUserTaskResponseBody struct {
	NextToken *string   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalSize *int32    `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s SearchUserTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchUserTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SearchUserTaskResponseBody) SetNextToken(v string) *SearchUserTaskResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchUserTaskResponseBody) SetRequestId(v string) *SearchUserTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchUserTaskResponseBody) SetResult(v []*string) *SearchUserTaskResponseBody {
	s.Result = v
	return s
}

func (s *SearchUserTaskResponseBody) SetTotalSize(v int32) *SearchUserTaskResponseBody {
	s.TotalSize = &v
	return s
}

type SearchUserTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchUserTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchUserTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchUserTaskResponse) GoString() string {
	return s.String()
}

func (s *SearchUserTaskResponse) SetHeaders(v map[string]*string) *SearchUserTaskResponse {
	s.Headers = v
	return s
}

func (s *SearchUserTaskResponse) SetStatusCode(v int32) *SearchUserTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchUserTaskResponse) SetBody(v *SearchUserTaskResponseBody) *SearchUserTaskResponse {
	s.Body = v
	return s
}

type SuspendProjectHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SuspendProjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s SuspendProjectHeaders) GoString() string {
	return s.String()
}

func (s *SuspendProjectHeaders) SetCommonHeaders(v map[string]*string) *SuspendProjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SuspendProjectHeaders) SetXAcsDingtalkAccessToken(v string) *SuspendProjectHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SuspendProjectResponseBody struct {
	Result *SuspendProjectResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s SuspendProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SuspendProjectResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendProjectResponseBody) SetResult(v *SuspendProjectResponseBodyResult) *SuspendProjectResponseBody {
	s.Result = v
	return s
}

type SuspendProjectResponseBodyResult struct {
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s SuspendProjectResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SuspendProjectResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SuspendProjectResponseBodyResult) SetUpdated(v string) *SuspendProjectResponseBodyResult {
	s.Updated = &v
	return s
}

type SuspendProjectResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SuspendProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SuspendProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s SuspendProjectResponse) GoString() string {
	return s.String()
}

func (s *SuspendProjectResponse) SetHeaders(v map[string]*string) *SuspendProjectResponse {
	s.Headers = v
	return s
}

func (s *SuspendProjectResponse) SetStatusCode(v int32) *SuspendProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendProjectResponse) SetBody(v *SuspendProjectResponseBody) *SuspendProjectResponse {
	s.Body = v
	return s
}

type UnSuspendProjectHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UnSuspendProjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s UnSuspendProjectHeaders) GoString() string {
	return s.String()
}

func (s *UnSuspendProjectHeaders) SetCommonHeaders(v map[string]*string) *UnSuspendProjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UnSuspendProjectHeaders) SetXAcsDingtalkAccessToken(v string) *UnSuspendProjectHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UnSuspendProjectResponseBody struct {
	Result *UnSuspendProjectResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UnSuspendProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnSuspendProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UnSuspendProjectResponseBody) SetResult(v *UnSuspendProjectResponseBodyResult) *UnSuspendProjectResponseBody {
	s.Result = v
	return s
}

type UnSuspendProjectResponseBodyResult struct {
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UnSuspendProjectResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UnSuspendProjectResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UnSuspendProjectResponseBodyResult) SetUpdated(v string) *UnSuspendProjectResponseBodyResult {
	s.Updated = &v
	return s
}

type UnSuspendProjectResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnSuspendProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnSuspendProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s UnSuspendProjectResponse) GoString() string {
	return s.String()
}

func (s *UnSuspendProjectResponse) SetHeaders(v map[string]*string) *UnSuspendProjectResponse {
	s.Headers = v
	return s
}

func (s *UnSuspendProjectResponse) SetStatusCode(v int32) *UnSuspendProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *UnSuspendProjectResponse) SetBody(v *UnSuspendProjectResponseBody) *UnSuspendProjectResponse {
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
	CustomfieldId   *string                               `json:"customfieldId,omitempty" xml:"customfieldId,omitempty"`
	CustomfieldName *string                               `json:"customfieldName,omitempty" xml:"customfieldName,omitempty"`
	Value           []*UpdateCustomfieldValueRequestValue `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
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
	CustomfieldId *string                                                      `json:"customfieldId,omitempty" xml:"customfieldId,omitempty"`
	Value         []*UpdateCustomfieldValueResponseBodyResultCustomfieldsValue `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateCustomfieldValueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateCustomfieldValueResponse) SetStatusCode(v int32) *UpdateCustomfieldValueResponse {
	s.StatusCode = &v
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
	Content             *string `json:"content,omitempty" xml:"content,omitempty"`
	DisableActivity     *bool   `json:"disableActivity,omitempty" xml:"disableActivity,omitempty"`
	DisableNotification *bool   `json:"disableNotification,omitempty" xml:"disableNotification,omitempty"`
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
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateOrganizationTaskContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateOrganizationTaskContentResponse) SetStatusCode(v int32) *UpdateOrganizationTaskContentResponse {
	s.StatusCode = &v
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
	DisableActivity     *bool   `json:"disableActivity,omitempty" xml:"disableActivity,omitempty"`
	DisableNotification *bool   `json:"disableNotification,omitempty" xml:"disableNotification,omitempty"`
	DueDate             *string `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
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
	DueDate    *string `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateOrganizationTaskDueDateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateOrganizationTaskDueDateResponse) SetStatusCode(v int32) *UpdateOrganizationTaskDueDateResponse {
	s.StatusCode = &v
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
	Executor   *UpdateOrganizationTaskExecutorResponseBodyResultExecutor    `json:"executor,omitempty" xml:"executor,omitempty" type:"Struct"`
	ExecutorId *string                                                      `json:"executorId,omitempty" xml:"executorId,omitempty"`
	Involvers  []*UpdateOrganizationTaskExecutorResponseBodyResultInvolvers `json:"involvers,omitempty" xml:"involvers,omitempty" type:"Repeated"`
	Updated    *string                                                      `json:"updated,omitempty" xml:"updated,omitempty"`
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
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateOrganizationTaskExecutorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateOrganizationTaskExecutorResponse) SetStatusCode(v int32) *UpdateOrganizationTaskExecutorResponse {
	s.StatusCode = &v
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
	AddInvolvers        []*string `json:"addInvolvers,omitempty" xml:"addInvolvers,omitempty" type:"Repeated"`
	DelInvolvers        []*string `json:"delInvolvers,omitempty" xml:"delInvolvers,omitempty" type:"Repeated"`
	DisableActivity     *bool     `json:"disableActivity,omitempty" xml:"disableActivity,omitempty"`
	DisableNotification *bool     `json:"disableNotification,omitempty" xml:"disableNotification,omitempty"`
	InvolveMembers      []*string `json:"involveMembers,omitempty" xml:"involveMembers,omitempty" type:"Repeated"`
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
	Involvers []*UpdateOrganizationTaskInvolveMembersResponseBodyResultInvolvers `json:"involvers,omitempty" xml:"involvers,omitempty" type:"Repeated"`
	Updated   *string                                                            `json:"updated,omitempty" xml:"updated,omitempty"`
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
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateOrganizationTaskInvolveMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateOrganizationTaskInvolveMembersResponse) SetStatusCode(v int32) *UpdateOrganizationTaskInvolveMembersResponse {
	s.StatusCode = &v
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
	DisableActivity     *bool   `json:"disableActivity,omitempty" xml:"disableActivity,omitempty"`
	DisableNotification *bool   `json:"disableNotification,omitempty" xml:"disableNotification,omitempty"`
	Note                *string `json:"note,omitempty" xml:"note,omitempty"`
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
	Note    *string `json:"note,omitempty" xml:"note,omitempty"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateOrganizationTaskNoteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateOrganizationTaskNoteResponse) SetStatusCode(v int32) *UpdateOrganizationTaskNoteResponse {
	s.StatusCode = &v
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
	DisableActivity     *bool  `json:"disableActivity,omitempty" xml:"disableActivity,omitempty"`
	DisableNotification *bool  `json:"disableNotification,omitempty" xml:"disableNotification,omitempty"`
	Priority            *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
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
	Priority *int32  `json:"priority,omitempty" xml:"priority,omitempty"`
	Updated  *string `json:"updated,omitempty" xml:"updated,omitempty"`
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateOrganizationTaskPriorityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateOrganizationTaskPriorityResponse) SetStatusCode(v int32) *UpdateOrganizationTaskPriorityResponse {
	s.StatusCode = &v
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
	DisableActivity     *bool `json:"disableActivity,omitempty" xml:"disableActivity,omitempty"`
	DisableNotification *bool `json:"disableNotification,omitempty" xml:"disableNotification,omitempty"`
	IsDone              *bool `json:"isDone,omitempty" xml:"isDone,omitempty"`
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
	IsDone     *bool   `json:"isDone,omitempty" xml:"isDone,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateOrganizationTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateOrganizationTaskStatusResponse) SetStatusCode(v int32) *UpdateOrganizationTaskStatusResponse {
	s.StatusCode = &v
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
	AddProjectGroupIds []*string `json:"addProjectGroupIds,omitempty" xml:"addProjectGroupIds,omitempty" type:"Repeated"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateProjectGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateProjectGroupResponse) SetStatusCode(v int32) *UpdateProjectGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProjectGroupResponse) SetBody(v *UpdateProjectGroupResponseBody) *UpdateProjectGroupResponse {
	s.Body = v
	return s
}

type UpdateTaskContentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateTaskContentHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskContentHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTaskContentHeaders) SetCommonHeaders(v map[string]*string) *UpdateTaskContentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTaskContentHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateTaskContentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateTaskContentRequest struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s UpdateTaskContentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskContentRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskContentRequest) SetContent(v string) *UpdateTaskContentRequest {
	s.Content = &v
	return s
}

type UpdateTaskContentResponseBody struct {
	Result *UpdateTaskContentResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateTaskContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskContentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskContentResponseBody) SetResult(v *UpdateTaskContentResponseBodyResult) *UpdateTaskContentResponseBody {
	s.Result = v
	return s
}

type UpdateTaskContentResponseBodyResult struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateTaskContentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskContentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateTaskContentResponseBodyResult) SetContent(v string) *UpdateTaskContentResponseBodyResult {
	s.Content = &v
	return s
}

func (s *UpdateTaskContentResponseBodyResult) SetUpdated(v string) *UpdateTaskContentResponseBodyResult {
	s.Updated = &v
	return s
}

type UpdateTaskContentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTaskContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTaskContentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskContentResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskContentResponse) SetHeaders(v map[string]*string) *UpdateTaskContentResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskContentResponse) SetStatusCode(v int32) *UpdateTaskContentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskContentResponse) SetBody(v *UpdateTaskContentResponseBody) *UpdateTaskContentResponse {
	s.Body = v
	return s
}

type UpdateTaskDueDateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateTaskDueDateHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskDueDateHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTaskDueDateHeaders) SetCommonHeaders(v map[string]*string) *UpdateTaskDueDateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTaskDueDateHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateTaskDueDateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateTaskDueDateRequest struct {
	DueDate *string `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
}

func (s UpdateTaskDueDateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskDueDateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskDueDateRequest) SetDueDate(v string) *UpdateTaskDueDateRequest {
	s.DueDate = &v
	return s
}

type UpdateTaskDueDateResponseBody struct {
	Result *UpdateTaskDueDateResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateTaskDueDateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskDueDateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskDueDateResponseBody) SetResult(v *UpdateTaskDueDateResponseBodyResult) *UpdateTaskDueDateResponseBody {
	s.Result = v
	return s
}

type UpdateTaskDueDateResponseBodyResult struct {
	DueDate *string `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateTaskDueDateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskDueDateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateTaskDueDateResponseBodyResult) SetDueDate(v string) *UpdateTaskDueDateResponseBodyResult {
	s.DueDate = &v
	return s
}

func (s *UpdateTaskDueDateResponseBodyResult) SetUpdated(v string) *UpdateTaskDueDateResponseBodyResult {
	s.Updated = &v
	return s
}

type UpdateTaskDueDateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTaskDueDateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTaskDueDateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskDueDateResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskDueDateResponse) SetHeaders(v map[string]*string) *UpdateTaskDueDateResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskDueDateResponse) SetStatusCode(v int32) *UpdateTaskDueDateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskDueDateResponse) SetBody(v *UpdateTaskDueDateResponseBody) *UpdateTaskDueDateResponse {
	s.Body = v
	return s
}

type UpdateTaskExecutorHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateTaskExecutorHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskExecutorHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTaskExecutorHeaders) SetCommonHeaders(v map[string]*string) *UpdateTaskExecutorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTaskExecutorHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateTaskExecutorHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateTaskExecutorRequest struct {
	ExecutorId *string `json:"executorId,omitempty" xml:"executorId,omitempty"`
}

func (s UpdateTaskExecutorRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskExecutorRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskExecutorRequest) SetExecutorId(v string) *UpdateTaskExecutorRequest {
	s.ExecutorId = &v
	return s
}

type UpdateTaskExecutorResponseBody struct {
	Result *UpdateTaskExecutorResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateTaskExecutorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskExecutorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskExecutorResponseBody) SetResult(v *UpdateTaskExecutorResponseBodyResult) *UpdateTaskExecutorResponseBody {
	s.Result = v
	return s
}

type UpdateTaskExecutorResponseBodyResult struct {
	ExecutorId *string `json:"executorId,omitempty" xml:"executorId,omitempty"`
	Updated    *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateTaskExecutorResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskExecutorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateTaskExecutorResponseBodyResult) SetExecutorId(v string) *UpdateTaskExecutorResponseBodyResult {
	s.ExecutorId = &v
	return s
}

func (s *UpdateTaskExecutorResponseBodyResult) SetUpdated(v string) *UpdateTaskExecutorResponseBodyResult {
	s.Updated = &v
	return s
}

type UpdateTaskExecutorResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTaskExecutorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTaskExecutorResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskExecutorResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskExecutorResponse) SetHeaders(v map[string]*string) *UpdateTaskExecutorResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskExecutorResponse) SetStatusCode(v int32) *UpdateTaskExecutorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskExecutorResponse) SetBody(v *UpdateTaskExecutorResponseBody) *UpdateTaskExecutorResponse {
	s.Body = v
	return s
}

type UpdateTaskInvolvemembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateTaskInvolvemembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskInvolvemembersHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTaskInvolvemembersHeaders) SetCommonHeaders(v map[string]*string) *UpdateTaskInvolvemembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTaskInvolvemembersHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateTaskInvolvemembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateTaskInvolvemembersRequest struct {
	AddInvolvers   []*string `json:"addInvolvers,omitempty" xml:"addInvolvers,omitempty" type:"Repeated"`
	DelInvolvers   []*string `json:"delInvolvers,omitempty" xml:"delInvolvers,omitempty" type:"Repeated"`
	InvolveMembers []*string `json:"involveMembers,omitempty" xml:"involveMembers,omitempty" type:"Repeated"`
}

func (s UpdateTaskInvolvemembersRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskInvolvemembersRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskInvolvemembersRequest) SetAddInvolvers(v []*string) *UpdateTaskInvolvemembersRequest {
	s.AddInvolvers = v
	return s
}

func (s *UpdateTaskInvolvemembersRequest) SetDelInvolvers(v []*string) *UpdateTaskInvolvemembersRequest {
	s.DelInvolvers = v
	return s
}

func (s *UpdateTaskInvolvemembersRequest) SetInvolveMembers(v []*string) *UpdateTaskInvolvemembersRequest {
	s.InvolveMembers = v
	return s
}

type UpdateTaskInvolvemembersResponseBody struct {
	Result *UpdateTaskInvolvemembersResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateTaskInvolvemembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskInvolvemembersResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskInvolvemembersResponseBody) SetResult(v *UpdateTaskInvolvemembersResponseBodyResult) *UpdateTaskInvolvemembersResponseBody {
	s.Result = v
	return s
}

type UpdateTaskInvolvemembersResponseBodyResult struct {
	InvolveMembers []*string `json:"involveMembers,omitempty" xml:"involveMembers,omitempty" type:"Repeated"`
	Updated        *string   `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateTaskInvolvemembersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskInvolvemembersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateTaskInvolvemembersResponseBodyResult) SetInvolveMembers(v []*string) *UpdateTaskInvolvemembersResponseBodyResult {
	s.InvolveMembers = v
	return s
}

func (s *UpdateTaskInvolvemembersResponseBodyResult) SetUpdated(v string) *UpdateTaskInvolvemembersResponseBodyResult {
	s.Updated = &v
	return s
}

type UpdateTaskInvolvemembersResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTaskInvolvemembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTaskInvolvemembersResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskInvolvemembersResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskInvolvemembersResponse) SetHeaders(v map[string]*string) *UpdateTaskInvolvemembersResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskInvolvemembersResponse) SetStatusCode(v int32) *UpdateTaskInvolvemembersResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskInvolvemembersResponse) SetBody(v *UpdateTaskInvolvemembersResponseBody) *UpdateTaskInvolvemembersResponse {
	s.Body = v
	return s
}

type UpdateTaskNoteHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateTaskNoteHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskNoteHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTaskNoteHeaders) SetCommonHeaders(v map[string]*string) *UpdateTaskNoteHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTaskNoteHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateTaskNoteHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateTaskNoteRequest struct {
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
}

func (s UpdateTaskNoteRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskNoteRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskNoteRequest) SetNote(v string) *UpdateTaskNoteRequest {
	s.Note = &v
	return s
}

type UpdateTaskNoteResponseBody struct {
	Result *UpdateTaskNoteResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateTaskNoteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskNoteResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskNoteResponseBody) SetResult(v *UpdateTaskNoteResponseBodyResult) *UpdateTaskNoteResponseBody {
	s.Result = v
	return s
}

type UpdateTaskNoteResponseBodyResult struct {
	Note    *string `json:"note,omitempty" xml:"note,omitempty"`
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateTaskNoteResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskNoteResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateTaskNoteResponseBodyResult) SetNote(v string) *UpdateTaskNoteResponseBodyResult {
	s.Note = &v
	return s
}

func (s *UpdateTaskNoteResponseBodyResult) SetUpdated(v string) *UpdateTaskNoteResponseBodyResult {
	s.Updated = &v
	return s
}

type UpdateTaskNoteResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTaskNoteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTaskNoteResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskNoteResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskNoteResponse) SetHeaders(v map[string]*string) *UpdateTaskNoteResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskNoteResponse) SetStatusCode(v int32) *UpdateTaskNoteResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskNoteResponse) SetBody(v *UpdateTaskNoteResponseBody) *UpdateTaskNoteResponse {
	s.Body = v
	return s
}

type UpdateTaskPriorityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateTaskPriorityHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskPriorityHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTaskPriorityHeaders) SetCommonHeaders(v map[string]*string) *UpdateTaskPriorityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTaskPriorityHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateTaskPriorityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateTaskPriorityRequest struct {
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
}

func (s UpdateTaskPriorityRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskPriorityRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskPriorityRequest) SetPriority(v int32) *UpdateTaskPriorityRequest {
	s.Priority = &v
	return s
}

type UpdateTaskPriorityResponseBody struct {
	Result *UpdateTaskPriorityResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateTaskPriorityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskPriorityResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskPriorityResponseBody) SetResult(v *UpdateTaskPriorityResponseBodyResult) *UpdateTaskPriorityResponseBody {
	s.Result = v
	return s
}

type UpdateTaskPriorityResponseBodyResult struct {
	Priority *int32  `json:"priority,omitempty" xml:"priority,omitempty"`
	Updated  *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateTaskPriorityResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskPriorityResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateTaskPriorityResponseBodyResult) SetPriority(v int32) *UpdateTaskPriorityResponseBodyResult {
	s.Priority = &v
	return s
}

func (s *UpdateTaskPriorityResponseBodyResult) SetUpdated(v string) *UpdateTaskPriorityResponseBodyResult {
	s.Updated = &v
	return s
}

type UpdateTaskPriorityResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTaskPriorityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTaskPriorityResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskPriorityResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskPriorityResponse) SetHeaders(v map[string]*string) *UpdateTaskPriorityResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskPriorityResponse) SetStatusCode(v int32) *UpdateTaskPriorityResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskPriorityResponse) SetBody(v *UpdateTaskPriorityResponseBody) *UpdateTaskPriorityResponse {
	s.Body = v
	return s
}

type UpdateTaskStageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateTaskStageHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskStageHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTaskStageHeaders) SetCommonHeaders(v map[string]*string) *UpdateTaskStageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTaskStageHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateTaskStageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateTaskStageRequest struct {
	StageId *string `json:"stageId,omitempty" xml:"stageId,omitempty"`
}

func (s UpdateTaskStageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskStageRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskStageRequest) SetStageId(v string) *UpdateTaskStageRequest {
	s.StageId = &v
	return s
}

type UpdateTaskStageResponseBody struct {
	Result *UpdateTaskStageResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateTaskStageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskStageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskStageResponseBody) SetResult(v *UpdateTaskStageResponseBodyResult) *UpdateTaskStageResponseBody {
	s.Result = v
	return s
}

type UpdateTaskStageResponseBodyResult struct {
	AccomplishTime *string `json:"accomplishTime,omitempty" xml:"accomplishTime,omitempty"`
	IsDone         *bool   `json:"isDone,omitempty" xml:"isDone,omitempty"`
	ProjectId      *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	StageId        *string `json:"stageId,omitempty" xml:"stageId,omitempty"`
	TaskId         *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TaskListId     *string `json:"taskListId,omitempty" xml:"taskListId,omitempty"`
	Updated        *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateTaskStageResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskStageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateTaskStageResponseBodyResult) SetAccomplishTime(v string) *UpdateTaskStageResponseBodyResult {
	s.AccomplishTime = &v
	return s
}

func (s *UpdateTaskStageResponseBodyResult) SetIsDone(v bool) *UpdateTaskStageResponseBodyResult {
	s.IsDone = &v
	return s
}

func (s *UpdateTaskStageResponseBodyResult) SetProjectId(v string) *UpdateTaskStageResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *UpdateTaskStageResponseBodyResult) SetStageId(v string) *UpdateTaskStageResponseBodyResult {
	s.StageId = &v
	return s
}

func (s *UpdateTaskStageResponseBodyResult) SetTaskId(v string) *UpdateTaskStageResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *UpdateTaskStageResponseBodyResult) SetTaskListId(v string) *UpdateTaskStageResponseBodyResult {
	s.TaskListId = &v
	return s
}

func (s *UpdateTaskStageResponseBodyResult) SetUpdated(v string) *UpdateTaskStageResponseBodyResult {
	s.Updated = &v
	return s
}

type UpdateTaskStageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTaskStageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTaskStageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskStageResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskStageResponse) SetHeaders(v map[string]*string) *UpdateTaskStageResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskStageResponse) SetStatusCode(v int32) *UpdateTaskStageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskStageResponse) SetBody(v *UpdateTaskStageResponseBody) *UpdateTaskStageResponse {
	s.Body = v
	return s
}

type UpdateTaskStartdateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateTaskStartdateHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskStartdateHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTaskStartdateHeaders) SetCommonHeaders(v map[string]*string) *UpdateTaskStartdateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTaskStartdateHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateTaskStartdateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateTaskStartdateRequest struct {
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
}

func (s UpdateTaskStartdateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskStartdateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskStartdateRequest) SetStartDate(v string) *UpdateTaskStartdateRequest {
	s.StartDate = &v
	return s
}

type UpdateTaskStartdateResponseBody struct {
	Result *UpdateTaskStartdateResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateTaskStartdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskStartdateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskStartdateResponseBody) SetResult(v *UpdateTaskStartdateResponseBodyResult) *UpdateTaskStartdateResponseBody {
	s.Result = v
	return s
}

type UpdateTaskStartdateResponseBodyResult struct {
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	Updated   *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateTaskStartdateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskStartdateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateTaskStartdateResponseBodyResult) SetStartDate(v string) *UpdateTaskStartdateResponseBodyResult {
	s.StartDate = &v
	return s
}

func (s *UpdateTaskStartdateResponseBodyResult) SetUpdated(v string) *UpdateTaskStartdateResponseBodyResult {
	s.Updated = &v
	return s
}

type UpdateTaskStartdateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTaskStartdateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTaskStartdateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskStartdateResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskStartdateResponse) SetHeaders(v map[string]*string) *UpdateTaskStartdateResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskStartdateResponse) SetStatusCode(v int32) *UpdateTaskStartdateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskStartdateResponse) SetBody(v *UpdateTaskStartdateResponseBody) *UpdateTaskStartdateResponse {
	s.Body = v
	return s
}

type UpdateTaskTaskflowstatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateTaskTaskflowstatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskTaskflowstatusHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTaskTaskflowstatusHeaders) SetCommonHeaders(v map[string]*string) *UpdateTaskTaskflowstatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTaskTaskflowstatusHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateTaskTaskflowstatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateTaskTaskflowstatusRequest struct {
	TaskflowStatusId *string `json:"taskflowStatusId,omitempty" xml:"taskflowStatusId,omitempty"`
	TfsUpdateNote    *string `json:"tfsUpdateNote,omitempty" xml:"tfsUpdateNote,omitempty"`
}

func (s UpdateTaskTaskflowstatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskTaskflowstatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskTaskflowstatusRequest) SetTaskflowStatusId(v string) *UpdateTaskTaskflowstatusRequest {
	s.TaskflowStatusId = &v
	return s
}

func (s *UpdateTaskTaskflowstatusRequest) SetTfsUpdateNote(v string) *UpdateTaskTaskflowstatusRequest {
	s.TfsUpdateNote = &v
	return s
}

type UpdateTaskTaskflowstatusResponseBody struct {
	Result *UpdateTaskTaskflowstatusResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateTaskTaskflowstatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskTaskflowstatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskTaskflowstatusResponseBody) SetResult(v *UpdateTaskTaskflowstatusResponseBodyResult) *UpdateTaskTaskflowstatusResponseBody {
	s.Result = v
	return s
}

type UpdateTaskTaskflowstatusResponseBodyResult struct {
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateTaskTaskflowstatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskTaskflowstatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateTaskTaskflowstatusResponseBodyResult) SetUpdated(v string) *UpdateTaskTaskflowstatusResponseBodyResult {
	s.Updated = &v
	return s
}

type UpdateTaskTaskflowstatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTaskTaskflowstatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTaskTaskflowstatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskTaskflowstatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskTaskflowstatusResponse) SetHeaders(v map[string]*string) *UpdateTaskTaskflowstatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskTaskflowstatusResponse) SetStatusCode(v int32) *UpdateTaskTaskflowstatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskTaskflowstatusResponse) SetBody(v *UpdateTaskTaskflowstatusResponseBody) *UpdateTaskTaskflowstatusResponse {
	s.Body = v
	return s
}

type UpdateWorkTimeApproveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateWorkTimeApproveHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkTimeApproveHeaders) GoString() string {
	return s.String()
}

func (s *UpdateWorkTimeApproveHeaders) SetCommonHeaders(v map[string]*string) *UpdateWorkTimeApproveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateWorkTimeApproveHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateWorkTimeApproveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateWorkTimeApproveRequest struct {
	FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Status     *string `json:"status,omitempty" xml:"status,omitempty"`
	SubmitTime *string `json:"submitTime,omitempty" xml:"submitTime,omitempty"`
	Title      *string `json:"title,omitempty" xml:"title,omitempty"`
	Url        *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s UpdateWorkTimeApproveRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkTimeApproveRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkTimeApproveRequest) SetFinishTime(v string) *UpdateWorkTimeApproveRequest {
	s.FinishTime = &v
	return s
}

func (s *UpdateWorkTimeApproveRequest) SetInstanceId(v string) *UpdateWorkTimeApproveRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateWorkTimeApproveRequest) SetStatus(v string) *UpdateWorkTimeApproveRequest {
	s.Status = &v
	return s
}

func (s *UpdateWorkTimeApproveRequest) SetSubmitTime(v string) *UpdateWorkTimeApproveRequest {
	s.SubmitTime = &v
	return s
}

func (s *UpdateWorkTimeApproveRequest) SetTitle(v string) *UpdateWorkTimeApproveRequest {
	s.Title = &v
	return s
}

func (s *UpdateWorkTimeApproveRequest) SetUrl(v string) *UpdateWorkTimeApproveRequest {
	s.Url = &v
	return s
}

type UpdateWorkTimeApproveResponseBody struct {
	Message   *string                                  `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *UpdateWorkTimeApproveResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateWorkTimeApproveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkTimeApproveResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkTimeApproveResponseBody) SetMessage(v string) *UpdateWorkTimeApproveResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWorkTimeApproveResponseBody) SetRequestId(v string) *UpdateWorkTimeApproveResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkTimeApproveResponseBody) SetResult(v *UpdateWorkTimeApproveResponseBodyResult) *UpdateWorkTimeApproveResponseBody {
	s.Result = v
	return s
}

type UpdateWorkTimeApproveResponseBodyResult struct {
	ApproveOpenId  *string   `json:"approveOpenId,omitempty" xml:"approveOpenId,omitempty"`
	CreatedAt      *string   `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatorId      *string   `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	FinishTime     *string   `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	InstanceId     *string   `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	OrganizationId *string   `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	Status         *string   `json:"status,omitempty" xml:"status,omitempty"`
	SubmitTime     *string   `json:"submitTime,omitempty" xml:"submitTime,omitempty"`
	TaskId         *string   `json:"taskId,omitempty" xml:"taskId,omitempty"`
	Time           *int32    `json:"time,omitempty" xml:"time,omitempty"`
	Title          *string   `json:"title,omitempty" xml:"title,omitempty"`
	UpdatedAt      *string   `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	Url            *string   `json:"url,omitempty" xml:"url,omitempty"`
	UserId         *string   `json:"userId,omitempty" xml:"userId,omitempty"`
	WorkTimeIds    []*string `json:"workTimeIds,omitempty" xml:"workTimeIds,omitempty" type:"Repeated"`
}

func (s UpdateWorkTimeApproveResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkTimeApproveResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateWorkTimeApproveResponseBodyResult) SetApproveOpenId(v string) *UpdateWorkTimeApproveResponseBodyResult {
	s.ApproveOpenId = &v
	return s
}

func (s *UpdateWorkTimeApproveResponseBodyResult) SetCreatedAt(v string) *UpdateWorkTimeApproveResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *UpdateWorkTimeApproveResponseBodyResult) SetCreatorId(v string) *UpdateWorkTimeApproveResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *UpdateWorkTimeApproveResponseBodyResult) SetFinishTime(v string) *UpdateWorkTimeApproveResponseBodyResult {
	s.FinishTime = &v
	return s
}

func (s *UpdateWorkTimeApproveResponseBodyResult) SetInstanceId(v string) *UpdateWorkTimeApproveResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *UpdateWorkTimeApproveResponseBodyResult) SetOrganizationId(v string) *UpdateWorkTimeApproveResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *UpdateWorkTimeApproveResponseBodyResult) SetStatus(v string) *UpdateWorkTimeApproveResponseBodyResult {
	s.Status = &v
	return s
}

func (s *UpdateWorkTimeApproveResponseBodyResult) SetSubmitTime(v string) *UpdateWorkTimeApproveResponseBodyResult {
	s.SubmitTime = &v
	return s
}

func (s *UpdateWorkTimeApproveResponseBodyResult) SetTaskId(v string) *UpdateWorkTimeApproveResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *UpdateWorkTimeApproveResponseBodyResult) SetTime(v int32) *UpdateWorkTimeApproveResponseBodyResult {
	s.Time = &v
	return s
}

func (s *UpdateWorkTimeApproveResponseBodyResult) SetTitle(v string) *UpdateWorkTimeApproveResponseBodyResult {
	s.Title = &v
	return s
}

func (s *UpdateWorkTimeApproveResponseBodyResult) SetUpdatedAt(v string) *UpdateWorkTimeApproveResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateWorkTimeApproveResponseBodyResult) SetUrl(v string) *UpdateWorkTimeApproveResponseBodyResult {
	s.Url = &v
	return s
}

func (s *UpdateWorkTimeApproveResponseBodyResult) SetUserId(v string) *UpdateWorkTimeApproveResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *UpdateWorkTimeApproveResponseBodyResult) SetWorkTimeIds(v []*string) *UpdateWorkTimeApproveResponseBodyResult {
	s.WorkTimeIds = v
	return s
}

type UpdateWorkTimeApproveResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateWorkTimeApproveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateWorkTimeApproveResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkTimeApproveResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkTimeApproveResponse) SetHeaders(v map[string]*string) *UpdateWorkTimeApproveResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkTimeApproveResponse) SetStatusCode(v int32) *UpdateWorkTimeApproveResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkTimeApproveResponse) SetBody(v *UpdateWorkTimeApproveResponseBody) *UpdateWorkTimeApproveResponse {
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

func (client *Client) AddProjectMemberWithOptions(userId *string, projectId *string, request *AddProjectMemberRequest, headers *AddProjectMemberHeaders, runtime *util.RuntimeOptions) (_result *AddProjectMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
	params := &openapi.Params{
		Action:      tea.String("AddProjectMember"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/projects/" + tea.StringValue(projectId) + "/members"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddProjectMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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

func (client *Client) ArchiveProjectWithOptions(userId *string, projectId *string, headers *ArchiveProjectHeaders, runtime *util.RuntimeOptions) (_result *ArchiveProjectResponse, _err error) {
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
		Action:      tea.String("ArchiveProject"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/projects/" + tea.StringValue(projectId) + "/archive"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ArchiveProjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ArchiveProject(userId *string, projectId *string) (_result *ArchiveProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ArchiveProjectHeaders{}
	_result = &ArchiveProjectResponse{}
	_body, _err := client.ArchiveProjectWithOptions(userId, projectId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ArchiveTaskWithOptions(userId *string, taskId *string, headers *ArchiveTaskHeaders, runtime *util.RuntimeOptions) (_result *ArchiveTaskResponse, _err error) {
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
		Action:      tea.String("ArchiveTask"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/tasks/" + tea.StringValue(taskId) + "/archive"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ArchiveTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ArchiveTask(userId *string, taskId *string) (_result *ArchiveTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ArchiveTaskHeaders{}
	_result = &ArchiveTaskResponse{}
	_body, _err := client.ArchiveTaskWithOptions(userId, taskId, headers, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("CreateOrganizationTask"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/organizations/users/" + tea.StringValue(userId) + "/tasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOrganizationTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) CreatePlanTimeWithOptions(userId *string, request *CreatePlanTimeRequest, headers *CreatePlanTimeHeaders, runtime *util.RuntimeOptions) (_result *CreatePlanTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
	params := &openapi.Params{
		Action:      tea.String("CreatePlanTime"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/planTimes"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePlanTimeResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) CreateProjectWithOptions(userId *string, request *CreateProjectRequest, headers *CreateProjectHeaders, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
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
		Action:      tea.String("CreateProject"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/projects"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProject(userId *string, request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateProjectHeaders{}
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(userId, request, headers, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("CreateProjectByTemplate"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/templates/projects"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProjectByTemplateResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) CreateProjectCustomfieldStatusWithOptions(userId *string, projectId *string, request *CreateProjectCustomfieldStatusRequest, headers *CreateProjectCustomfieldStatusHeaders, runtime *util.RuntimeOptions) (_result *CreateProjectCustomfieldStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomfieldId)) {
		body["customfieldId"] = request.CustomfieldId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomfieldInstanceId)) {
		body["customfieldInstanceId"] = request.CustomfieldInstanceId
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
	params := &openapi.Params{
		Action:      tea.String("CreateProjectCustomfieldStatus"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/projects/" + tea.StringValue(projectId) + "/customfields"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProjectCustomfieldStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProjectCustomfieldStatus(userId *string, projectId *string, request *CreateProjectCustomfieldStatusRequest) (_result *CreateProjectCustomfieldStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateProjectCustomfieldStatusHeaders{}
	_result = &CreateProjectCustomfieldStatusResponse{}
	_body, _err := client.CreateProjectCustomfieldStatusWithOptions(userId, projectId, request, headers, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.ParentTaskId)) {
		body["parentTaskId"] = request.ParentTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["projectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ScenariofieldconfigId)) {
		body["scenariofieldconfigId"] = request.ScenariofieldconfigId
	}

	if !tea.BoolValue(util.IsUnset(request.StageId)) {
		body["stageId"] = request.StageId
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["startDate"] = request.StartDate
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
	params := &openapi.Params{
		Action:      tea.String("CreateTask"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/tasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) CreateTaskObjectLinkWithOptions(userId *string, taskId *string, request *CreateTaskObjectLinkRequest, headers *CreateTaskObjectLinkHeaders, runtime *util.RuntimeOptions) (_result *CreateTaskObjectLinkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LinkedData)) {
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
	params := &openapi.Params{
		Action:      tea.String("CreateTaskObjectLink"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/tasks/" + tea.StringValue(taskId) + "/objectLinks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTaskObjectLinkResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) CreateWorkTimeWithOptions(userId *string, request *CreateWorkTimeRequest, headers *CreateWorkTimeHeaders, runtime *util.RuntimeOptions) (_result *CreateWorkTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
	params := &openapi.Params{
		Action:      tea.String("CreateWorkTime"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/workTimes"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWorkTimeResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) CreateWorkTimeApproveWithOptions(userId *string, request *CreateWorkTimeApproveRequest, headers *CreateWorkTimeApproveHeaders, runtime *util.RuntimeOptions) (_result *CreateWorkTimeApproveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkTimeIds)) {
		body["workTimeIds"] = request.WorkTimeIds
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
		Action:      tea.String("CreateWorkTimeApprove"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/workTimes/approvals"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWorkTimeApproveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWorkTimeApprove(userId *string, request *CreateWorkTimeApproveRequest) (_result *CreateWorkTimeApproveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateWorkTimeApproveHeaders{}
	_result = &CreateWorkTimeApproveResponse{}
	_body, _err := client.CreateWorkTimeApproveWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProjectMemberWithOptions(userId *string, projectId *string, request *DeleteProjectMemberRequest, headers *DeleteProjectMemberHeaders, runtime *util.RuntimeOptions) (_result *DeleteProjectMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
	params := &openapi.Params{
		Action:      tea.String("DeleteProjectMember"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/projects/" + tea.StringValue(projectId) + "/members/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProjectMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProjectMember(userId *string, projectId *string, request *DeleteProjectMemberRequest) (_result *DeleteProjectMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteProjectMemberHeaders{}
	_result = &DeleteProjectMemberResponse{}
	_body, _err := client.DeleteProjectMemberWithOptions(userId, projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTaskWithOptions(userId *string, taskId *string, headers *DeleteTaskHeaders, runtime *util.RuntimeOptions) (_result *DeleteTaskResponse, _err error) {
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
		Action:      tea.String("DeleteTask"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/tasks/" + tea.StringValue(taskId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTask(userId *string, taskId *string) (_result *DeleteTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteTaskHeaders{}
	_result = &DeleteTaskResponse{}
	_body, _err := client.DeleteTaskWithOptions(userId, taskId, headers, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetDeptsByOrgId"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/orgs/depts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeptsByOrgIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetEmpsByOrgId"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/orgs/employees"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEmpsByOrgIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) GetOrganizatioTaskByIdsWithOptions(userId *string, request *GetOrganizatioTaskByIdsRequest, headers *GetOrganizatioTaskByIdsHeaders, runtime *util.RuntimeOptions) (_result *GetOrganizatioTaskByIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
	params := &openapi.Params{
		Action:      tea.String("GetOrganizatioTaskByIds"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/organizations/users/" + tea.StringValue(userId) + "/tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOrganizatioTaskByIdsResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) GetOrganizationPriorityListWithOptions(userId *string, headers *GetOrganizationPriorityListHeaders, runtime *util.RuntimeOptions) (_result *GetOrganizationPriorityListResponse, _err error) {
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
		Action:      tea.String("GetOrganizationPriorityList"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/organizations/users/" + tea.StringValue(userId) + "/priorities"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOrganizationPriorityListResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) GetOrganizationTaskWithOptions(taskId *string, userId *string, headers *GetOrganizationTaskHeaders, runtime *util.RuntimeOptions) (_result *GetOrganizationTaskResponse, _err error) {
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
		Action:      tea.String("GetOrganizationTask"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/organizations/users/" + tea.StringValue(userId) + "/tasks/" + tea.StringValue(taskId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOrganizationTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) GetProjectGroupWithOptions(userId *string, request *GetProjectGroupRequest, headers *GetProjectGroupHeaders, runtime *util.RuntimeOptions) (_result *GetProjectGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
	params := &openapi.Params{
		Action:      tea.String("GetProjectGroup"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/organizations/users/" + tea.StringValue(userId) + "/groups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) GetProjectMemebersWithOptions(userId *string, projectId *string, request *GetProjectMemebersRequest, headers *GetProjectMemebersHeaders, runtime *util.RuntimeOptions) (_result *GetProjectMemebersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectRoleId)) {
		query["projectRoleId"] = request.ProjectRoleId
	}

	if !tea.BoolValue(util.IsUnset(request.Skip)) {
		query["skip"] = request.Skip
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		query["userIds"] = request.UserIds
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
		Action:      tea.String("GetProjectMemebers"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/projects/" + tea.StringValue(projectId) + "/members"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectMemebersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProjectMemebers(userId *string, projectId *string, request *GetProjectMemebersRequest) (_result *GetProjectMemebersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetProjectMemebersHeaders{}
	_result = &GetProjectMemebersResponse{}
	_body, _err := client.GetProjectMemebersWithOptions(userId, projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProjectStatusListWithOptions(userId *string, projectId *string, headers *GetProjectStatusListHeaders, runtime *util.RuntimeOptions) (_result *GetProjectStatusListResponse, _err error) {
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
		Action:      tea.String("GetProjectStatusList"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/projects/" + tea.StringValue(projectId) + "/statuses"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectStatusListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProjectStatusList(userId *string, projectId *string) (_result *GetProjectStatusListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetProjectStatusListHeaders{}
	_result = &GetProjectStatusListResponse{}
	_body, _err := client.GetProjectStatusListWithOptions(userId, projectId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTaskByIdsWithOptions(userId *string, request *GetTaskByIdsRequest, headers *GetTaskByIdsHeaders, runtime *util.RuntimeOptions) (_result *GetTaskByIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ParentTaskId)) {
		query["parentTaskId"] = request.ParentTaskId
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
		Action:      tea.String("GetTaskByIds"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskByIdsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTaskByIds(userId *string, request *GetTaskByIdsRequest) (_result *GetTaskByIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTaskByIdsHeaders{}
	_result = &GetTaskByIdsResponse{}
	_body, _err := client.GetTaskByIdsWithOptions(userId, request, headers, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetTbOrgIdByDingOrgId"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/teambition/organizations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTbOrgIdByDingOrgIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetTbProjectGray"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/projects/gray"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTbProjectGrayResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetTbProjectSource"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/projects/source"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTbProjectSourceResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetTbUserIdByStaffId"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/teambition/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTbUserIdByStaffIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) GetUserJoinedProjectWithOptions(userId *string, request *GetUserJoinedProjectRequest, headers *GetUserJoinedProjectHeaders, runtime *util.RuntimeOptions) (_result *GetUserJoinedProjectResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserJoinedProject"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/joinProjects"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserJoinedProjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserJoinedProject(userId *string, request *GetUserJoinedProjectRequest) (_result *GetUserJoinedProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserJoinedProjectHeaders{}
	_result = &GetUserJoinedProjectResponse{}
	_body, _err := client.GetUserJoinedProjectWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryProjectWithOptions(userId *string, request *QueryProjectRequest, headers *QueryProjectHeaders, runtime *util.RuntimeOptions) (_result *QueryProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectIds)) {
		query["projectIds"] = request.ProjectIds
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		query["sourceId"] = request.SourceId
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
		Action:      tea.String("QueryProject"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/projects/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryProjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryProject(userId *string, request *QueryProjectRequest) (_result *QueryProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryProjectHeaders{}
	_result = &QueryProjectResponse{}
	_body, _err := client.QueryProjectWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTaskOfProjectWithOptions(userId *string, projectId *string, request *QueryTaskOfProjectRequest, headers *QueryTaskOfProjectHeaders, runtime *util.RuntimeOptions) (_result *QueryTaskOfProjectResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
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
		Action:      tea.String("QueryTaskOfProject"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/projectIds/" + tea.StringValue(projectId) + "/tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTaskOfProjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTaskOfProject(userId *string, projectId *string, request *QueryTaskOfProjectRequest) (_result *QueryTaskOfProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryTaskOfProjectHeaders{}
	_result = &QueryTaskOfProjectResponse{}
	_body, _err := client.QueryTaskOfProjectWithOptions(userId, projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SeachTaskStageWithOptions(userId *string, projectId *string, request *SeachTaskStageRequest, headers *SeachTaskStageHeaders, runtime *util.RuntimeOptions) (_result *SeachTaskStageResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.StageIds)) {
		query["stageIds"] = request.StageIds
	}

	if !tea.BoolValue(util.IsUnset(request.TaskListId)) {
		query["taskListId"] = request.TaskListId
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
		Action:      tea.String("SeachTaskStage"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/projects/" + tea.StringValue(projectId) + "/taskStages/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SeachTaskStageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SeachTaskStage(userId *string, projectId *string, request *SeachTaskStageRequest) (_result *SeachTaskStageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SeachTaskStageHeaders{}
	_result = &SeachTaskStageResponse{}
	_body, _err := client.SeachTaskStageWithOptions(userId, projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchOranizationCustomfieldWithOptions(userId *string, request *SearchOranizationCustomfieldRequest, headers *SearchOranizationCustomfieldHeaders, runtime *util.RuntimeOptions) (_result *SearchOranizationCustomfieldResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomfieldIds)) {
		query["customfieldIds"] = request.CustomfieldIds
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["instanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectIds)) {
		query["projectIds"] = request.ProjectIds
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["scope"] = request.Scope
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
		Action:      tea.String("SearchOranizationCustomfield"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/organizations/users/" + tea.StringValue(userId) + "/customfields/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchOranizationCustomfieldResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchOranizationCustomfield(userId *string, request *SearchOranizationCustomfieldRequest) (_result *SearchOranizationCustomfieldResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchOranizationCustomfieldHeaders{}
	_result = &SearchOranizationCustomfieldResponse{}
	_body, _err := client.SearchOranizationCustomfieldWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchProjectCustomfieldWithOptions(userId *string, projectId *string, request *SearchProjectCustomfieldRequest, headers *SearchProjectCustomfieldHeaders, runtime *util.RuntimeOptions) (_result *SearchProjectCustomfieldResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomfieldIds)) {
		query["customfieldIds"] = request.CustomfieldIds
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["instanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.ScenariofieldconfigId)) {
		query["scenariofieldconfigId"] = request.ScenariofieldconfigId
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["scope"] = request.Scope
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
		Action:      tea.String("SearchProjectCustomfield"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/projects/" + tea.StringValue(projectId) + "/customfields/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchProjectCustomfieldResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchProjectCustomfield(userId *string, projectId *string, request *SearchProjectCustomfieldRequest) (_result *SearchProjectCustomfieldResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchProjectCustomfieldHeaders{}
	_result = &SearchProjectCustomfieldResponse{}
	_body, _err := client.SearchProjectCustomfieldWithOptions(userId, projectId, request, headers, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("SearchProjectTemplate"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/organizations/users/" + tea.StringValue(userId) + "/templates"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchProjectTemplateResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) SearchTaskFlowWithOptions(userId *string, projectId *string, request *SearchTaskFlowRequest, headers *SearchTaskFlowHeaders, runtime *util.RuntimeOptions) (_result *SearchTaskFlowResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.TaskflowIds)) {
		query["taskflowIds"] = request.TaskflowIds
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
		Action:      tea.String("SearchTaskFlow"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/projects/" + tea.StringValue(projectId) + "/taskflows/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchTaskFlowResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchTaskFlow(userId *string, projectId *string, request *SearchTaskFlowRequest) (_result *SearchTaskFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchTaskFlowHeaders{}
	_result = &SearchTaskFlowResponse{}
	_body, _err := client.SearchTaskFlowWithOptions(userId, projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchTaskListWithOptions(userId *string, projectId *string, request *SearchTaskListRequest, headers *SearchTaskListHeaders, runtime *util.RuntimeOptions) (_result *SearchTaskListResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.TaskListIds)) {
		query["taskListIds"] = request.TaskListIds
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
		Action:      tea.String("SearchTaskList"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/projects/" + tea.StringValue(projectId) + "/taskLists/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchTaskListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchTaskList(userId *string, projectId *string, request *SearchTaskListRequest) (_result *SearchTaskListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchTaskListHeaders{}
	_result = &SearchTaskListResponse{}
	_body, _err := client.SearchTaskListWithOptions(userId, projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchTaskflowStatusWithOptions(userId *string, projectId *string, request *SearchTaskflowStatusRequest, headers *SearchTaskflowStatusHeaders, runtime *util.RuntimeOptions) (_result *SearchTaskflowStatusResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.TfIds)) {
		query["tfIds"] = request.TfIds
	}

	if !tea.BoolValue(util.IsUnset(request.TfsIds)) {
		query["tfsIds"] = request.TfsIds
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
		Action:      tea.String("SearchTaskflowStatus"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/projects/" + tea.StringValue(projectId) + "/taskflowStatuses/search"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchTaskflowStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchTaskflowStatus(userId *string, projectId *string, request *SearchTaskflowStatusRequest) (_result *SearchTaskflowStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchTaskflowStatusHeaders{}
	_result = &SearchTaskflowStatusResponse{}
	_body, _err := client.SearchTaskflowStatusWithOptions(userId, projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchUserTaskWithOptions(userId *string, request *SearchUserTaskRequest, headers *SearchUserTaskHeaders, runtime *util.RuntimeOptions) (_result *SearchUserTaskResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Tql)) {
		query["tql"] = request.Tql
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
		Action:      tea.String("SearchUserTask"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/tasks/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchUserTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchUserTask(userId *string, request *SearchUserTaskRequest) (_result *SearchUserTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchUserTaskHeaders{}
	_result = &SearchUserTaskResponse{}
	_body, _err := client.SearchUserTaskWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SuspendProjectWithOptions(projectId *string, userId *string, headers *SuspendProjectHeaders, runtime *util.RuntimeOptions) (_result *SuspendProjectResponse, _err error) {
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
		Action:      tea.String("SuspendProject"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/projects/" + tea.StringValue(projectId) + "/suspend"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SuspendProjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SuspendProject(projectId *string, userId *string) (_result *SuspendProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SuspendProjectHeaders{}
	_result = &SuspendProjectResponse{}
	_body, _err := client.SuspendProjectWithOptions(projectId, userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnSuspendProjectWithOptions(projectId *string, userId *string, headers *UnSuspendProjectHeaders, runtime *util.RuntimeOptions) (_result *UnSuspendProjectResponse, _err error) {
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
		Action:      tea.String("UnSuspendProject"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/projects/" + tea.StringValue(projectId) + "/unsuspend"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UnSuspendProjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnSuspendProject(projectId *string, userId *string) (_result *UnSuspendProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UnSuspendProjectHeaders{}
	_result = &UnSuspendProjectResponse{}
	_body, _err := client.UnSuspendProjectWithOptions(projectId, userId, headers, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("UpdateCustomfieldValue"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/tasks/" + tea.StringValue(taskId) + "/customFields"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCustomfieldValueResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) UpdateOrganizationTaskContentWithOptions(taskId *string, userId *string, request *UpdateOrganizationTaskContentRequest, headers *UpdateOrganizationTaskContentHeaders, runtime *util.RuntimeOptions) (_result *UpdateOrganizationTaskContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
	params := &openapi.Params{
		Action:      tea.String("UpdateOrganizationTaskContent"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/organizations/users/" + tea.StringValue(userId) + "/tasks/" + tea.StringValue(taskId) + "/contents"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateOrganizationTaskContentResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) UpdateOrganizationTaskDueDateWithOptions(taskId *string, userId *string, request *UpdateOrganizationTaskDueDateRequest, headers *UpdateOrganizationTaskDueDateHeaders, runtime *util.RuntimeOptions) (_result *UpdateOrganizationTaskDueDateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
	params := &openapi.Params{
		Action:      tea.String("UpdateOrganizationTaskDueDate"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/organizations/users/" + tea.StringValue(userId) + "/tasks/" + tea.StringValue(taskId) + "/dueDates"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateOrganizationTaskDueDateResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) UpdateOrganizationTaskExecutorWithOptions(taskId *string, userId *string, request *UpdateOrganizationTaskExecutorRequest, headers *UpdateOrganizationTaskExecutorHeaders, runtime *util.RuntimeOptions) (_result *UpdateOrganizationTaskExecutorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
	params := &openapi.Params{
		Action:      tea.String("UpdateOrganizationTaskExecutor"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/organizations/users/" + tea.StringValue(userId) + "/tasks/" + tea.StringValue(taskId) + "/executors"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateOrganizationTaskExecutorResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) UpdateOrganizationTaskInvolveMembersWithOptions(taskId *string, userId *string, request *UpdateOrganizationTaskInvolveMembersRequest, headers *UpdateOrganizationTaskInvolveMembersHeaders, runtime *util.RuntimeOptions) (_result *UpdateOrganizationTaskInvolveMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
	params := &openapi.Params{
		Action:      tea.String("UpdateOrganizationTaskInvolveMembers"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/organizations/users/" + tea.StringValue(userId) + "/tasks/" + tea.StringValue(taskId) + "/involveMembers"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateOrganizationTaskInvolveMembersResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) UpdateOrganizationTaskNoteWithOptions(taskId *string, userId *string, request *UpdateOrganizationTaskNoteRequest, headers *UpdateOrganizationTaskNoteHeaders, runtime *util.RuntimeOptions) (_result *UpdateOrganizationTaskNoteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
	params := &openapi.Params{
		Action:      tea.String("UpdateOrganizationTaskNote"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/organizations/users/" + tea.StringValue(userId) + "/tasks/" + tea.StringValue(taskId) + "/notes"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateOrganizationTaskNoteResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) UpdateOrganizationTaskPriorityWithOptions(taskId *string, userId *string, request *UpdateOrganizationTaskPriorityRequest, headers *UpdateOrganizationTaskPriorityHeaders, runtime *util.RuntimeOptions) (_result *UpdateOrganizationTaskPriorityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
	params := &openapi.Params{
		Action:      tea.String("UpdateOrganizationTaskPriority"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/organizations/users/" + tea.StringValue(userId) + "/tasks/" + tea.StringValue(taskId) + "/priorities"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateOrganizationTaskPriorityResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) UpdateOrganizationTaskStatusWithOptions(taskId *string, userId *string, request *UpdateOrganizationTaskStatusRequest, headers *UpdateOrganizationTaskStatusHeaders, runtime *util.RuntimeOptions) (_result *UpdateOrganizationTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
	params := &openapi.Params{
		Action:      tea.String("UpdateOrganizationTaskStatus"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/organizations/users/" + tea.StringValue(userId) + "/tasks/" + tea.StringValue(taskId) + "/states"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateOrganizationTaskStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) UpdateProjectGroupWithOptions(userId *string, projectId *string, request *UpdateProjectGroupRequest, headers *UpdateProjectGroupHeaders, runtime *util.RuntimeOptions) (_result *UpdateProjectGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
	params := &openapi.Params{
		Action:      tea.String("UpdateProjectGroup"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/projects/" + tea.StringValue(projectId) + "/groups"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProjectGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) UpdateTaskContentWithOptions(userId *string, taskId *string, request *UpdateTaskContentRequest, headers *UpdateTaskContentHeaders, runtime *util.RuntimeOptions) (_result *UpdateTaskContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
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
		Action:      tea.String("UpdateTaskContent"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/tasks/" + tea.StringValue(taskId) + "/contents"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTaskContentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTaskContent(userId *string, taskId *string, request *UpdateTaskContentRequest) (_result *UpdateTaskContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTaskContentHeaders{}
	_result = &UpdateTaskContentResponse{}
	_body, _err := client.UpdateTaskContentWithOptions(userId, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTaskDueDateWithOptions(userId *string, taskId *string, request *UpdateTaskDueDateRequest, headers *UpdateTaskDueDateHeaders, runtime *util.RuntimeOptions) (_result *UpdateTaskDueDateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
	params := &openapi.Params{
		Action:      tea.String("UpdateTaskDueDate"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/tasks/" + tea.StringValue(taskId) + "/dueDates"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTaskDueDateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTaskDueDate(userId *string, taskId *string, request *UpdateTaskDueDateRequest) (_result *UpdateTaskDueDateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTaskDueDateHeaders{}
	_result = &UpdateTaskDueDateResponse{}
	_body, _err := client.UpdateTaskDueDateWithOptions(userId, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTaskExecutorWithOptions(userId *string, taskId *string, request *UpdateTaskExecutorRequest, headers *UpdateTaskExecutorHeaders, runtime *util.RuntimeOptions) (_result *UpdateTaskExecutorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
	params := &openapi.Params{
		Action:      tea.String("UpdateTaskExecutor"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/tasks/" + tea.StringValue(taskId) + "/executors"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTaskExecutorResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTaskExecutor(userId *string, taskId *string, request *UpdateTaskExecutorRequest) (_result *UpdateTaskExecutorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTaskExecutorHeaders{}
	_result = &UpdateTaskExecutorResponse{}
	_body, _err := client.UpdateTaskExecutorWithOptions(userId, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTaskInvolvemembersWithOptions(userId *string, taskId *string, request *UpdateTaskInvolvemembersRequest, headers *UpdateTaskInvolvemembersHeaders, runtime *util.RuntimeOptions) (_result *UpdateTaskInvolvemembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddInvolvers)) {
		body["addInvolvers"] = request.AddInvolvers
	}

	if !tea.BoolValue(util.IsUnset(request.DelInvolvers)) {
		body["delInvolvers"] = request.DelInvolvers
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
	params := &openapi.Params{
		Action:      tea.String("UpdateTaskInvolvemembers"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/tasks/" + tea.StringValue(taskId) + "/involveMembers"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTaskInvolvemembersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTaskInvolvemembers(userId *string, taskId *string, request *UpdateTaskInvolvemembersRequest) (_result *UpdateTaskInvolvemembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTaskInvolvemembersHeaders{}
	_result = &UpdateTaskInvolvemembersResponse{}
	_body, _err := client.UpdateTaskInvolvemembersWithOptions(userId, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTaskNoteWithOptions(userId *string, taskId *string, request *UpdateTaskNoteRequest, headers *UpdateTaskNoteHeaders, runtime *util.RuntimeOptions) (_result *UpdateTaskNoteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
	params := &openapi.Params{
		Action:      tea.String("UpdateTaskNote"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/tasks/" + tea.StringValue(taskId) + "/notes"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTaskNoteResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTaskNote(userId *string, taskId *string, request *UpdateTaskNoteRequest) (_result *UpdateTaskNoteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTaskNoteHeaders{}
	_result = &UpdateTaskNoteResponse{}
	_body, _err := client.UpdateTaskNoteWithOptions(userId, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTaskPriorityWithOptions(userId *string, taskId *string, request *UpdateTaskPriorityRequest, headers *UpdateTaskPriorityHeaders, runtime *util.RuntimeOptions) (_result *UpdateTaskPriorityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
	params := &openapi.Params{
		Action:      tea.String("UpdateTaskPriority"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/tasks/" + tea.StringValue(taskId) + "/priorities"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTaskPriorityResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTaskPriority(userId *string, taskId *string, request *UpdateTaskPriorityRequest) (_result *UpdateTaskPriorityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTaskPriorityHeaders{}
	_result = &UpdateTaskPriorityResponse{}
	_body, _err := client.UpdateTaskPriorityWithOptions(userId, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTaskStageWithOptions(userId *string, taskId *string, request *UpdateTaskStageRequest, headers *UpdateTaskStageHeaders, runtime *util.RuntimeOptions) (_result *UpdateTaskStageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StageId)) {
		body["stageId"] = request.StageId
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
		Action:      tea.String("UpdateTaskStage"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/tasks/" + tea.StringValue(taskId) + "/stages"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTaskStageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTaskStage(userId *string, taskId *string, request *UpdateTaskStageRequest) (_result *UpdateTaskStageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTaskStageHeaders{}
	_result = &UpdateTaskStageResponse{}
	_body, _err := client.UpdateTaskStageWithOptions(userId, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTaskStartdateWithOptions(userId *string, taskId *string, request *UpdateTaskStartdateRequest, headers *UpdateTaskStartdateHeaders, runtime *util.RuntimeOptions) (_result *UpdateTaskStartdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["startDate"] = request.StartDate
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
		Action:      tea.String("UpdateTaskStartdate"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/tasks/" + tea.StringValue(taskId) + "/startDates"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTaskStartdateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTaskStartdate(userId *string, taskId *string, request *UpdateTaskStartdateRequest) (_result *UpdateTaskStartdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTaskStartdateHeaders{}
	_result = &UpdateTaskStartdateResponse{}
	_body, _err := client.UpdateTaskStartdateWithOptions(userId, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTaskTaskflowstatusWithOptions(userId *string, taskId *string, request *UpdateTaskTaskflowstatusRequest, headers *UpdateTaskTaskflowstatusHeaders, runtime *util.RuntimeOptions) (_result *UpdateTaskTaskflowstatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskflowStatusId)) {
		body["taskflowStatusId"] = request.TaskflowStatusId
	}

	if !tea.BoolValue(util.IsUnset(request.TfsUpdateNote)) {
		body["tfsUpdateNote"] = request.TfsUpdateNote
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
		Action:      tea.String("UpdateTaskTaskflowstatus"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/tasks/" + tea.StringValue(taskId) + "/taskflowStatuses"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTaskTaskflowstatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTaskTaskflowstatus(userId *string, taskId *string, request *UpdateTaskTaskflowstatusRequest) (_result *UpdateTaskTaskflowstatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTaskTaskflowstatusHeaders{}
	_result = &UpdateTaskTaskflowstatusResponse{}
	_body, _err := client.UpdateTaskTaskflowstatusWithOptions(userId, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWorkTimeApproveWithOptions(userId *string, approveOpenId *string, request *UpdateWorkTimeApproveRequest, headers *UpdateWorkTimeApproveHeaders, runtime *util.RuntimeOptions) (_result *UpdateWorkTimeApproveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FinishTime)) {
		body["finishTime"] = request.FinishTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.SubmitTime)) {
		body["submitTime"] = request.SubmitTime
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["url"] = request.Url
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
		Action:      tea.String("UpdateWorkTimeApprove"),
		Version:     tea.String("project_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/project/users/" + tea.StringValue(userId) + "/workTimes/approvals/" + tea.StringValue(approveOpenId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWorkTimeApproveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWorkTimeApprove(userId *string, approveOpenId *string, request *UpdateWorkTimeApproveRequest) (_result *UpdateWorkTimeApproveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateWorkTimeApproveHeaders{}
	_result = &UpdateWorkTimeApproveResponse{}
	_body, _err := client.UpdateWorkTimeApproveWithOptions(userId, approveOpenId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
