// This file is auto-generated, don't edit it. Thanks.
package team_sphere_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AnalysisReportHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AnalysisReportHeaders) String() string {
	return tea.Prettify(s)
}

func (s AnalysisReportHeaders) GoString() string {
	return s.String()
}

func (s *AnalysisReportHeaders) SetCommonHeaders(v map[string]*string) *AnalysisReportHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AnalysisReportHeaders) SetXAcsDingtalkAccessToken(v string) *AnalysisReportHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AnalysisReportRequest struct {
	Filter   *AnalysisReportRequestFilter `json:"filter,omitempty" xml:"filter,omitempty" type:"Struct"`
	ReportId *string                      `json:"reportId,omitempty" xml:"reportId,omitempty"`
}

func (s AnalysisReportRequest) String() string {
	return tea.Prettify(s)
}

func (s AnalysisReportRequest) GoString() string {
	return s.String()
}

func (s *AnalysisReportRequest) SetFilter(v *AnalysisReportRequestFilter) *AnalysisReportRequest {
	s.Filter = v
	return s
}

func (s *AnalysisReportRequest) SetReportId(v string) *AnalysisReportRequest {
	s.ReportId = &v
	return s
}

type AnalysisReportRequestFilter struct {
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
}

func (s AnalysisReportRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s AnalysisReportRequestFilter) GoString() string {
	return s.String()
}

func (s *AnalysisReportRequestFilter) SetCreated(v string) *AnalysisReportRequestFilter {
	s.Created = &v
	return s
}

type AnalysisReportResponseBody struct {
	Result []*AnalysisReportResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s AnalysisReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AnalysisReportResponseBody) GoString() string {
	return s.String()
}

func (s *AnalysisReportResponseBody) SetResult(v []*AnalysisReportResponseBodyResult) *AnalysisReportResponseBody {
	s.Result = v
	return s
}

type AnalysisReportResponseBodyResult struct {
	Cols                  []*AnalysisReportResponseBodyResultCols                `json:"cols,omitempty" xml:"cols,omitempty" type:"Repeated"`
	ListQuery             [][]*AnalysisReportResponseBodyResultListQuery         `json:"listQuery,omitempty" xml:"listQuery,omitempty" type:"Repeated"`
	Name                  *string                                                `json:"name,omitempty" xml:"name,omitempty"`
	Rows                  [][]*string                                            `json:"rows,omitempty" xml:"rows,omitempty" type:"Repeated"`
	Tips                  *string                                                `json:"tips,omitempty" xml:"tips,omitempty"`
	Title                 *string                                                `json:"title,omitempty" xml:"title,omitempty"`
	VisualizationSettings *AnalysisReportResponseBodyResultVisualizationSettings `json:"visualizationSettings,omitempty" xml:"visualizationSettings,omitempty" type:"Struct"`
}

func (s AnalysisReportResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AnalysisReportResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AnalysisReportResponseBodyResult) SetCols(v []*AnalysisReportResponseBodyResultCols) *AnalysisReportResponseBodyResult {
	s.Cols = v
	return s
}

func (s *AnalysisReportResponseBodyResult) SetListQuery(v [][]*AnalysisReportResponseBodyResultListQuery) *AnalysisReportResponseBodyResult {
	s.ListQuery = v
	return s
}

func (s *AnalysisReportResponseBodyResult) SetName(v string) *AnalysisReportResponseBodyResult {
	s.Name = &v
	return s
}

func (s *AnalysisReportResponseBodyResult) SetRows(v [][]*string) *AnalysisReportResponseBodyResult {
	s.Rows = v
	return s
}

func (s *AnalysisReportResponseBodyResult) SetTips(v string) *AnalysisReportResponseBodyResult {
	s.Tips = &v
	return s
}

func (s *AnalysisReportResponseBodyResult) SetTitle(v string) *AnalysisReportResponseBodyResult {
	s.Title = &v
	return s
}

func (s *AnalysisReportResponseBodyResult) SetVisualizationSettings(v *AnalysisReportResponseBodyResultVisualizationSettings) *AnalysisReportResponseBodyResult {
	s.VisualizationSettings = v
	return s
}

type AnalysisReportResponseBodyResultCols struct {
	BaseType *string `json:"baseType,omitempty" xml:"baseType,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Theme    *string `json:"theme,omitempty" xml:"theme,omitempty"`
}

func (s AnalysisReportResponseBodyResultCols) String() string {
	return tea.Prettify(s)
}

func (s AnalysisReportResponseBodyResultCols) GoString() string {
	return s.String()
}

func (s *AnalysisReportResponseBodyResultCols) SetBaseType(v string) *AnalysisReportResponseBodyResultCols {
	s.BaseType = &v
	return s
}

func (s *AnalysisReportResponseBodyResultCols) SetName(v string) *AnalysisReportResponseBodyResultCols {
	s.Name = &v
	return s
}

func (s *AnalysisReportResponseBodyResultCols) SetTheme(v string) *AnalysisReportResponseBodyResultCols {
	s.Theme = &v
	return s
}

type AnalysisReportResponseBodyResultListQuery struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AnalysisReportResponseBodyResultListQuery) String() string {
	return tea.Prettify(s)
}

func (s AnalysisReportResponseBodyResultListQuery) GoString() string {
	return s.String()
}

func (s *AnalysisReportResponseBodyResultListQuery) SetKey(v string) *AnalysisReportResponseBodyResultListQuery {
	s.Key = &v
	return s
}

func (s *AnalysisReportResponseBodyResultListQuery) SetValue(v string) *AnalysisReportResponseBodyResultListQuery {
	s.Value = &v
	return s
}

type AnalysisReportResponseBodyResultVisualizationSettings struct {
	Dimension *int64  `json:"dimension,omitempty" xml:"dimension,omitempty"`
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AnalysisReportResponseBodyResultVisualizationSettings) String() string {
	return tea.Prettify(s)
}

func (s AnalysisReportResponseBodyResultVisualizationSettings) GoString() string {
	return s.String()
}

func (s *AnalysisReportResponseBodyResultVisualizationSettings) SetDimension(v int64) *AnalysisReportResponseBodyResultVisualizationSettings {
	s.Dimension = &v
	return s
}

func (s *AnalysisReportResponseBodyResultVisualizationSettings) SetType(v string) *AnalysisReportResponseBodyResultVisualizationSettings {
	s.Type = &v
	return s
}

type AnalysisReportResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AnalysisReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AnalysisReportResponse) String() string {
	return tea.Prettify(s)
}

func (s AnalysisReportResponse) GoString() string {
	return s.String()
}

func (s *AnalysisReportResponse) SetHeaders(v map[string]*string) *AnalysisReportResponse {
	s.Headers = v
	return s
}

func (s *AnalysisReportResponse) SetStatusCode(v int32) *AnalysisReportResponse {
	s.StatusCode = &v
	return s
}

func (s *AnalysisReportResponse) SetBody(v *AnalysisReportResponseBody) *AnalysisReportResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// 明天12点前完成周报撰写
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// false
	DisableActivity *bool `json:"disableActivity,omitempty" xml:"disableActivity,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// false
	DisableNotification *bool `json:"disableNotification,omitempty" xml:"disableNotification,omitempty"`
	// example:
	//
	// 2021-08-13T07:36:50.318Z
	DueDate *string `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	// example:
	//
	// 173xxxx
	ExecutorId     *string   `json:"executorId,omitempty" xml:"executorId,omitempty"`
	InvolveMembers []*string `json:"involveMembers,omitempty" xml:"involveMembers,omitempty" type:"Repeated"`
	// example:
	//
	// 我是一条任务备注
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// involves
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
	AncestorIds []*string `json:"ancestorIds,omitempty" xml:"ancestorIds,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	AttachmentsCount *int32 `json:"attachmentsCount,omitempty" xml:"attachmentsCount,omitempty"`
	// example:
	//
	// 明天12点前写好周报
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 2021-08-13T07:36:50.318Z
	Created *string                                          `json:"created,omitempty" xml:"created,omitempty"`
	Creator *CreateOrganizationTaskResponseBodyResultCreator `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	// example:
	//
	// 173xxxx
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// 2021-08-13T07:36:50.318Z
	DueDate  *string                                           `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	Executor *CreateOrganizationTaskResponseBodyResultExecutor `json:"executor,omitempty" xml:"executor,omitempty" type:"Struct"`
	// example:
	//
	// 173xxxx
	ExecutorId *string `json:"executorId,omitempty" xml:"executorId,omitempty"`
	// example:
	//
	// false
	HasReminder *bool `json:"hasReminder,omitempty" xml:"hasReminder,omitempty"`
	// example:
	//
	// 62a697c053c2ef5xxxxxx
	Id             *string                                              `json:"id,omitempty" xml:"id,omitempty"`
	InvolveMembers []*string                                            `json:"involveMembers,omitempty" xml:"involveMembers,omitempty" type:"Repeated"`
	Involvers      []*CreateOrganizationTaskResponseBodyResultInvolvers `json:"involvers,omitempty" xml:"involvers,omitempty" type:"Repeated"`
	// example:
	//
	// false
	IsDeleted *bool `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	// example:
	//
	// false
	IsDone *string `json:"isDone,omitempty" xml:"isDone,omitempty"`
	// example:
	//
	// 我是一条备注哦
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
	// example:
	//
	// 2021-08-13T07:36:50.318Z
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
	// example:
	//
	// members
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

func (s *CreateOrganizationTaskResponseBodyResult) SetUpdated(v string) *CreateOrganizationTaskResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *CreateOrganizationTaskResponseBodyResult) SetVisible(v string) *CreateOrganizationTaskResponseBodyResult {
	s.Visible = &v
	return s
}

type CreateOrganizationTaskResponseBodyResultCreator struct {
	// example:
	//
	// https://xxxxxxxxxx
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// 鬼斩
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 173xxxx
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
	// example:
	//
	// https://xxxxxxxxxx
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// 鬼斩
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 173xxxx
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
	// example:
	//
	// httpx://xxx
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// 173xxxx
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 鬼斩
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrganizationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CreateProjectMembersV3Headers struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateProjectMembersV3Headers) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectMembersV3Headers) GoString() string {
	return s.String()
}

func (s *CreateProjectMembersV3Headers) SetCommonHeaders(v map[string]*string) *CreateProjectMembersV3Headers {
	s.CommonHeaders = v
	return s
}

func (s *CreateProjectMembersV3Headers) SetXAcsDingtalkAccessToken(v string) *CreateProjectMembersV3Headers {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateProjectMembersV3Request struct {
	// This parameter is required.
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s CreateProjectMembersV3Request) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectMembersV3Request) GoString() string {
	return s.String()
}

func (s *CreateProjectMembersV3Request) SetUserIds(v []*string) *CreateProjectMembersV3Request {
	s.UserIds = v
	return s
}

type CreateProjectMembersV3ResponseBody struct {
	RequestId *string                                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*CreateProjectMembersV3ResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s CreateProjectMembersV3ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectMembersV3ResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectMembersV3ResponseBody) SetRequestId(v string) *CreateProjectMembersV3ResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectMembersV3ResponseBody) SetResult(v []*CreateProjectMembersV3ResponseBodyResult) *CreateProjectMembersV3ResponseBody {
	s.Result = v
	return s
}

type CreateProjectMembersV3ResponseBodyResult struct {
	BoundToObjectId   *string `json:"boundToObjectId,omitempty" xml:"boundToObjectId,omitempty"`
	BoundToObjectType *string `json:"boundToObjectType,omitempty" xml:"boundToObjectType,omitempty"`
	Joined            *string `json:"joined,omitempty" xml:"joined,omitempty"`
	Role              *int32  `json:"role,omitempty" xml:"role,omitempty"`
	UserId            *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateProjectMembersV3ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectMembersV3ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateProjectMembersV3ResponseBodyResult) SetBoundToObjectId(v string) *CreateProjectMembersV3ResponseBodyResult {
	s.BoundToObjectId = &v
	return s
}

func (s *CreateProjectMembersV3ResponseBodyResult) SetBoundToObjectType(v string) *CreateProjectMembersV3ResponseBodyResult {
	s.BoundToObjectType = &v
	return s
}

func (s *CreateProjectMembersV3ResponseBodyResult) SetJoined(v string) *CreateProjectMembersV3ResponseBodyResult {
	s.Joined = &v
	return s
}

func (s *CreateProjectMembersV3ResponseBodyResult) SetRole(v int32) *CreateProjectMembersV3ResponseBodyResult {
	s.Role = &v
	return s
}

func (s *CreateProjectMembersV3ResponseBodyResult) SetUserId(v string) *CreateProjectMembersV3ResponseBodyResult {
	s.UserId = &v
	return s
}

type CreateProjectMembersV3Response struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProjectMembersV3ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProjectMembersV3Response) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectMembersV3Response) GoString() string {
	return s.String()
}

func (s *CreateProjectMembersV3Response) SetHeaders(v map[string]*string) *CreateProjectMembersV3Response {
	s.Headers = v
	return s
}

func (s *CreateProjectMembersV3Response) SetStatusCode(v int32) *CreateProjectMembersV3Response {
	s.StatusCode = &v
	return s
}

func (s *CreateProjectMembersV3Response) SetBody(v *CreateProjectMembersV3ResponseBody) *CreateProjectMembersV3Response {
	s.Body = v
	return s
}

type CreateProjectV3Headers struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateProjectV3Headers) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectV3Headers) GoString() string {
	return s.String()
}

func (s *CreateProjectV3Headers) SetCommonHeaders(v map[string]*string) *CreateProjectV3Headers {
	s.CommonHeaders = v
	return s
}

func (s *CreateProjectV3Headers) SetXAcsDingtalkAccessToken(v string) *CreateProjectV3Headers {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateProjectV3Request struct {
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s CreateProjectV3Request) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectV3Request) GoString() string {
	return s.String()
}

func (s *CreateProjectV3Request) SetName(v string) *CreateProjectV3Request {
	s.Name = &v
	return s
}

func (s *CreateProjectV3Request) SetOrganizationId(v string) *CreateProjectV3Request {
	s.OrganizationId = &v
	return s
}

type CreateProjectV3ResponseBody struct {
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateProjectV3ResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateProjectV3ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectV3ResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectV3ResponseBody) SetRequestId(v string) *CreateProjectV3ResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectV3ResponseBody) SetResult(v *CreateProjectV3ResponseBodyResult) *CreateProjectV3ResponseBody {
	s.Result = v
	return s
}

type CreateProjectV3ResponseBodyResult struct {
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s CreateProjectV3ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectV3ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateProjectV3ResponseBodyResult) SetId(v string) *CreateProjectV3ResponseBodyResult {
	s.Id = &v
	return s
}

type CreateProjectV3Response struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProjectV3ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProjectV3Response) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectV3Response) GoString() string {
	return s.String()
}

func (s *CreateProjectV3Response) SetHeaders(v map[string]*string) *CreateProjectV3Response {
	s.Headers = v
	return s
}

func (s *CreateProjectV3Response) SetStatusCode(v int32) *CreateProjectV3Response {
	s.StatusCode = &v
	return s
}

func (s *CreateProjectV3Response) SetBody(v *CreateProjectV3ResponseBody) *CreateProjectV3Response {
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
	// This parameter is required.
	//
	// example:
	//
	// 任务标题
	Content             *string                          `json:"content,omitempty" xml:"content,omitempty"`
	Customfields        []*CreateTaskRequestCustomfields `json:"customfields,omitempty" xml:"customfields,omitempty" type:"Repeated"`
	DisableActivity     *bool                            `json:"disableActivity,omitempty" xml:"disableActivity,omitempty"`
	DisableNotification *bool                            `json:"disableNotification,omitempty" xml:"disableNotification,omitempty"`
	// example:
	//
	// 2022-06-13T07:36:50.318Z
	DueDate *string `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	// example:
	//
	// 173xxxx
	ExecutorId *string `json:"executorId,omitempty" xml:"executorId,omitempty"`
	// example:
	//
	// 我是一条任务备注
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 62c25e3b376exxxxxx
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

func (s *CreateTaskRequest) SetDisableActivity(v bool) *CreateTaskRequest {
	s.DisableActivity = &v
	return s
}

func (s *CreateTaskRequest) SetDisableNotification(v bool) *CreateTaskRequest {
	s.DisableNotification = &v
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

func (s *CreateTaskRequest) SetProjectId(v string) *CreateTaskRequest {
	s.ProjectId = &v
	return s
}

type CreateTaskRequestCustomfields struct {
	// example:
	//
	// 62fb0bxxxxxxx
	CustomfieldId *string `json:"customfieldId,omitempty" xml:"customfieldId,omitempty"`
	// example:
	//
	// 自定义字段-文本
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
	Id       *string `json:"id,omitempty" xml:"id,omitempty"`
	ThumbUrl *string `json:"thumbUrl,omitempty" xml:"thumbUrl,omitempty"`
	// example:
	//
	// 我是自定义字段显示值
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateTaskRequestCustomfieldsValue) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskRequestCustomfieldsValue) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestCustomfieldsValue) SetId(v string) *CreateTaskRequestCustomfieldsValue {
	s.Id = &v
	return s
}

func (s *CreateTaskRequestCustomfieldsValue) SetThumbUrl(v string) *CreateTaskRequestCustomfieldsValue {
	s.ThumbUrl = &v
	return s
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
	// example:
	//
	// 任务标题
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 2021-08-13T07:36:50.318Z
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// example:
	//
	// 173xxxxx
	CreatorId    *string                                     `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Customfields []*CreateTaskResponseBodyResultCustomfields `json:"customfields,omitempty" xml:"customfields,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-08-13T07:36:50.318Z
	DueDate *string `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	// example:
	//
	// 173xxxx
	ExecutorId     *string   `json:"executorId,omitempty" xml:"executorId,omitempty"`
	InvolveMembers []*string `json:"involveMembers,omitempty" xml:"involveMembers,omitempty" type:"Repeated"`
	// example:
	//
	// 我是一条备注
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
	// example:
	//
	// -10
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// example:
	//
	// 62c25e3b376ecxxxxxx
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// 62a697c053c2ef5xxxxxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 2021-08-13T07:36:50.318Z
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
	// example:
	//
	// 625bcxdxxxxxx
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
	// example:
	//
	// 我是自定义字段显示值
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DeleteProjectMembersV3Headers struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteProjectMembersV3Headers) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectMembersV3Headers) GoString() string {
	return s.String()
}

func (s *DeleteProjectMembersV3Headers) SetCommonHeaders(v map[string]*string) *DeleteProjectMembersV3Headers {
	s.CommonHeaders = v
	return s
}

func (s *DeleteProjectMembersV3Headers) SetXAcsDingtalkAccessToken(v string) *DeleteProjectMembersV3Headers {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteProjectMembersV3Request struct {
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s DeleteProjectMembersV3Request) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectMembersV3Request) GoString() string {
	return s.String()
}

func (s *DeleteProjectMembersV3Request) SetUserIds(v []*string) *DeleteProjectMembersV3Request {
	s.UserIds = v
	return s
}

type DeleteProjectMembersV3ResponseBody struct {
	Errors    []*DeleteProjectMembersV3ResponseBodyErrors `json:"errors,omitempty" xml:"errors,omitempty" type:"Repeated"`
	RequestId *string                                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*string                                   `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s DeleteProjectMembersV3ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectMembersV3ResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectMembersV3ResponseBody) SetErrors(v []*DeleteProjectMembersV3ResponseBodyErrors) *DeleteProjectMembersV3ResponseBody {
	s.Errors = v
	return s
}

func (s *DeleteProjectMembersV3ResponseBody) SetRequestId(v string) *DeleteProjectMembersV3ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProjectMembersV3ResponseBody) SetResult(v []*string) *DeleteProjectMembersV3ResponseBody {
	s.Result = v
	return s
}

type DeleteProjectMembersV3ResponseBodyErrors struct {
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s DeleteProjectMembersV3ResponseBodyErrors) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectMembersV3ResponseBodyErrors) GoString() string {
	return s.String()
}

func (s *DeleteProjectMembersV3ResponseBodyErrors) SetMessage(v string) *DeleteProjectMembersV3ResponseBodyErrors {
	s.Message = &v
	return s
}

type DeleteProjectMembersV3Response struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProjectMembersV3ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProjectMembersV3Response) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectMembersV3Response) GoString() string {
	return s.String()
}

func (s *DeleteProjectMembersV3Response) SetHeaders(v map[string]*string) *DeleteProjectMembersV3Response {
	s.Headers = v
	return s
}

func (s *DeleteProjectMembersV3Response) SetStatusCode(v int32) *DeleteProjectMembersV3Response {
	s.StatusCode = &v
	return s
}

func (s *DeleteProjectMembersV3Response) SetBody(v *DeleteProjectMembersV3ResponseBody) *DeleteProjectMembersV3Response {
	s.Body = v
	return s
}

type GetFootprintProjectHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFootprintProjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFootprintProjectHeaders) GoString() string {
	return s.String()
}

func (s *GetFootprintProjectHeaders) SetCommonHeaders(v map[string]*string) *GetFootprintProjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFootprintProjectHeaders) SetXAcsDingtalkAccessToken(v string) *GetFootprintProjectHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFootprintProjectResponseBody struct {
	Result []*GetFootprintProjectResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetFootprintProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFootprintProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetFootprintProjectResponseBody) SetResult(v []*GetFootprintProjectResponseBodyResult) *GetFootprintProjectResponseBody {
	s.Result = v
	return s
}

type GetFootprintProjectResponseBodyResult struct {
	// example:
	//
	// 2022-07-04T03:29:34.770Z
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// example:
	//
	// 5f687406f05b283425ea8f6f
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// xxxx
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1234
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// true
	IsDeleted *bool `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	// example:
	//
	// https://things.teambition.net?icon_logo=%F0%9F%92%A5
	Logo *string `json:"logo,omitempty" xml:"logo,omitempty"`
	// example:
	//
	// x项目
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 6139cd1aba5b128516ec8c86
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// 2022-07-04T03:29:34.770Z
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
	// example:
	//
	// project
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
}

func (s GetFootprintProjectResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFootprintProjectResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFootprintProjectResponseBodyResult) SetCreated(v string) *GetFootprintProjectResponseBodyResult {
	s.Created = &v
	return s
}

func (s *GetFootprintProjectResponseBodyResult) SetCreatorId(v string) *GetFootprintProjectResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *GetFootprintProjectResponseBodyResult) SetDescription(v string) *GetFootprintProjectResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetFootprintProjectResponseBodyResult) SetId(v string) *GetFootprintProjectResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetFootprintProjectResponseBodyResult) SetIsDeleted(v bool) *GetFootprintProjectResponseBodyResult {
	s.IsDeleted = &v
	return s
}

func (s *GetFootprintProjectResponseBodyResult) SetLogo(v string) *GetFootprintProjectResponseBodyResult {
	s.Logo = &v
	return s
}

func (s *GetFootprintProjectResponseBodyResult) SetName(v string) *GetFootprintProjectResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetFootprintProjectResponseBodyResult) SetOrganizationId(v string) *GetFootprintProjectResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *GetFootprintProjectResponseBodyResult) SetUpdated(v string) *GetFootprintProjectResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *GetFootprintProjectResponseBodyResult) SetVisibility(v string) *GetFootprintProjectResponseBodyResult {
	s.Visibility = &v
	return s
}

type GetFootprintProjectResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFootprintProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFootprintProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFootprintProjectResponse) GoString() string {
	return s.String()
}

func (s *GetFootprintProjectResponse) SetHeaders(v map[string]*string) *GetFootprintProjectResponse {
	s.Headers = v
	return s
}

func (s *GetFootprintProjectResponse) SetStatusCode(v int32) *GetFootprintProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFootprintProjectResponse) SetBody(v *GetFootprintProjectResponseBody) *GetFootprintProjectResponse {
	s.Body = v
	return s
}

type GetFootprintTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFootprintTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFootprintTaskHeaders) GoString() string {
	return s.String()
}

func (s *GetFootprintTaskHeaders) SetCommonHeaders(v map[string]*string) *GetFootprintTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFootprintTaskHeaders) SetXAcsDingtalkAccessToken(v string) *GetFootprintTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFootprintTaskResponseBody struct {
	Result []*GetFootprintTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetFootprintTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFootprintTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetFootprintTaskResponseBody) SetResult(v []*GetFootprintTaskResponseBodyResult) *GetFootprintTaskResponseBody {
	s.Result = v
	return s
}

type GetFootprintTaskResponseBodyResult struct {
	// example:
	//
	// 2024-09-19T11:07:51.468Z
	Accomplished *string                                           `json:"accomplished,omitempty" xml:"accomplished,omitempty"`
	BasicPos     *string                                           `json:"basicPos,omitempty" xml:"basicPos,omitempty"`
	Content      *string                                           `json:"content,omitempty" xml:"content,omitempty"`
	Created      *string                                           `json:"created,omitempty" xml:"created,omitempty"`
	CreatorId    *string                                           `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Customfields []*GetFootprintTaskResponseBodyResultCustomfields `json:"customfields,omitempty" xml:"customfields,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-09-13T10:00:00.000Z
	DueDate        *string   `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	ExecutorId     *string   `json:"executorId,omitempty" xml:"executorId,omitempty"`
	Id             *string   `json:"id,omitempty" xml:"id,omitempty"`
	InvolveMembers []*string `json:"involveMembers,omitempty" xml:"involveMembers,omitempty" type:"Repeated"`
	// example:
	//
	// false
	IsArchived *bool `json:"isArchived,omitempty" xml:"isArchived,omitempty"`
	IsDeleted  *bool `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	// example:
	//
	// true
	IsDone *bool `json:"isDone,omitempty" xml:"isDone,omitempty"`
	// example:
	//
	// test123
	Note           *string `json:"note,omitempty" xml:"note,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// 0
	Pos *int64 `json:"pos,omitempty" xml:"pos,omitempty"`
	// example:
	//
	// 0
	Priority  *int64  `json:"priority,omitempty" xml:"priority,omitempty"`
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	SfcId     *string `json:"sfcId,omitempty" xml:"sfcId,omitempty"`
	// example:
	//
	// 6639f974916cdb178e7d***
	StageId *string `json:"stageId,omitempty" xml:"stageId,omitempty"`
	// example:
	//
	// 2024-09-13T10:00:00.000Z
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// example:
	//
	// 6639f974916cdb178e7d***
	TasklistId *string `json:"tasklistId,omitempty" xml:"tasklistId,omitempty"`
	// example:
	//
	// 6639f974916cdb178e7****
	TfsId *string `json:"tfsId,omitempty" xml:"tfsId,omitempty"`
	// example:
	//
	// 540
	UniqueId *int64  `json:"uniqueId,omitempty" xml:"uniqueId,omitempty"`
	Updated  *string `json:"updated,omitempty" xml:"updated,omitempty"`
	// example:
	//
	// members
	Visible *string `json:"visible,omitempty" xml:"visible,omitempty"`
}

func (s GetFootprintTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFootprintTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFootprintTaskResponseBodyResult) SetAccomplished(v string) *GetFootprintTaskResponseBodyResult {
	s.Accomplished = &v
	return s
}

func (s *GetFootprintTaskResponseBodyResult) SetBasicPos(v string) *GetFootprintTaskResponseBodyResult {
	s.BasicPos = &v
	return s
}

func (s *GetFootprintTaskResponseBodyResult) SetContent(v string) *GetFootprintTaskResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetFootprintTaskResponseBodyResult) SetCreated(v string) *GetFootprintTaskResponseBodyResult {
	s.Created = &v
	return s
}

func (s *GetFootprintTaskResponseBodyResult) SetCreatorId(v string) *GetFootprintTaskResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *GetFootprintTaskResponseBodyResult) SetCustomfields(v []*GetFootprintTaskResponseBodyResultCustomfields) *GetFootprintTaskResponseBodyResult {
	s.Customfields = v
	return s
}

func (s *GetFootprintTaskResponseBodyResult) SetDueDate(v string) *GetFootprintTaskResponseBodyResult {
	s.DueDate = &v
	return s
}

func (s *GetFootprintTaskResponseBodyResult) SetExecutorId(v string) *GetFootprintTaskResponseBodyResult {
	s.ExecutorId = &v
	return s
}

func (s *GetFootprintTaskResponseBodyResult) SetId(v string) *GetFootprintTaskResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetFootprintTaskResponseBodyResult) SetInvolveMembers(v []*string) *GetFootprintTaskResponseBodyResult {
	s.InvolveMembers = v
	return s
}

func (s *GetFootprintTaskResponseBodyResult) SetIsArchived(v bool) *GetFootprintTaskResponseBodyResult {
	s.IsArchived = &v
	return s
}

func (s *GetFootprintTaskResponseBodyResult) SetIsDeleted(v bool) *GetFootprintTaskResponseBodyResult {
	s.IsDeleted = &v
	return s
}

func (s *GetFootprintTaskResponseBodyResult) SetIsDone(v bool) *GetFootprintTaskResponseBodyResult {
	s.IsDone = &v
	return s
}

func (s *GetFootprintTaskResponseBodyResult) SetNote(v string) *GetFootprintTaskResponseBodyResult {
	s.Note = &v
	return s
}

func (s *GetFootprintTaskResponseBodyResult) SetOrganizationId(v string) *GetFootprintTaskResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *GetFootprintTaskResponseBodyResult) SetPos(v int64) *GetFootprintTaskResponseBodyResult {
	s.Pos = &v
	return s
}

func (s *GetFootprintTaskResponseBodyResult) SetPriority(v int64) *GetFootprintTaskResponseBodyResult {
	s.Priority = &v
	return s
}

func (s *GetFootprintTaskResponseBodyResult) SetProjectId(v string) *GetFootprintTaskResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *GetFootprintTaskResponseBodyResult) SetSfcId(v string) *GetFootprintTaskResponseBodyResult {
	s.SfcId = &v
	return s
}

func (s *GetFootprintTaskResponseBodyResult) SetStageId(v string) *GetFootprintTaskResponseBodyResult {
	s.StageId = &v
	return s
}

func (s *GetFootprintTaskResponseBodyResult) SetStartDate(v string) *GetFootprintTaskResponseBodyResult {
	s.StartDate = &v
	return s
}

func (s *GetFootprintTaskResponseBodyResult) SetTasklistId(v string) *GetFootprintTaskResponseBodyResult {
	s.TasklistId = &v
	return s
}

func (s *GetFootprintTaskResponseBodyResult) SetTfsId(v string) *GetFootprintTaskResponseBodyResult {
	s.TfsId = &v
	return s
}

func (s *GetFootprintTaskResponseBodyResult) SetUniqueId(v int64) *GetFootprintTaskResponseBodyResult {
	s.UniqueId = &v
	return s
}

func (s *GetFootprintTaskResponseBodyResult) SetUpdated(v string) *GetFootprintTaskResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *GetFootprintTaskResponseBodyResult) SetVisible(v string) *GetFootprintTaskResponseBodyResult {
	s.Visible = &v
	return s
}

type GetFootprintTaskResponseBodyResultCustomfields struct {
	// example:
	//
	// 666a472803e46df8ce98a4a5
	CustomfieldId *string `json:"customfieldId,omitempty" xml:"customfieldId,omitempty"`
	// example:
	//
	// date
	Type   *string                  `json:"type,omitempty" xml:"type,omitempty"`
	Value  []map[string]interface{} `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
	Values []*string                `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetFootprintTaskResponseBodyResultCustomfields) String() string {
	return tea.Prettify(s)
}

func (s GetFootprintTaskResponseBodyResultCustomfields) GoString() string {
	return s.String()
}

func (s *GetFootprintTaskResponseBodyResultCustomfields) SetCustomfieldId(v string) *GetFootprintTaskResponseBodyResultCustomfields {
	s.CustomfieldId = &v
	return s
}

func (s *GetFootprintTaskResponseBodyResultCustomfields) SetType(v string) *GetFootprintTaskResponseBodyResultCustomfields {
	s.Type = &v
	return s
}

func (s *GetFootprintTaskResponseBodyResultCustomfields) SetValue(v []map[string]interface{}) *GetFootprintTaskResponseBodyResultCustomfields {
	s.Value = v
	return s
}

func (s *GetFootprintTaskResponseBodyResultCustomfields) SetValues(v []*string) *GetFootprintTaskResponseBodyResultCustomfields {
	s.Values = v
	return s
}

type GetFootprintTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFootprintTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFootprintTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFootprintTaskResponse) GoString() string {
	return s.String()
}

func (s *GetFootprintTaskResponse) SetHeaders(v map[string]*string) *GetFootprintTaskResponse {
	s.Headers = v
	return s
}

func (s *GetFootprintTaskResponse) SetStatusCode(v int32) *GetFootprintTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFootprintTaskResponse) SetBody(v *GetFootprintTaskResponseBody) *GetFootprintTaskResponse {
	s.Body = v
	return s
}

type GetFreeTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFreeTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFreeTaskHeaders) GoString() string {
	return s.String()
}

func (s *GetFreeTaskHeaders) SetCommonHeaders(v map[string]*string) *GetFreeTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFreeTaskHeaders) SetXAcsDingtalkAccessToken(v string) *GetFreeTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFreeTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0517xxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetFreeTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFreeTaskRequest) GoString() string {
	return s.String()
}

func (s *GetFreeTaskRequest) SetUserId(v string) *GetFreeTaskRequest {
	s.UserId = &v
	return s
}

type GetFreeTaskResponseBody struct {
	RequestId *string                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetFreeTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetFreeTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFreeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetFreeTaskResponseBody) SetRequestId(v string) *GetFreeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFreeTaskResponseBody) SetResult(v *GetFreeTaskResponseBodyResult) *GetFreeTaskResponseBody {
	s.Result = v
	return s
}

type GetFreeTaskResponseBodyResult struct {
	Accomplished *string   `json:"accomplished,omitempty" xml:"accomplished,omitempty"`
	AncestorIds  []*string `json:"ancestorIds,omitempty" xml:"ancestorIds,omitempty" type:"Repeated"`
	Content      *string   `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 2022-07-04T03:29:34.770Z
	Created        *string   `json:"created,omitempty" xml:"created,omitempty"`
	CreatorId      *string   `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	DueDate        *string   `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	ExecutorId     *string   `json:"executorId,omitempty" xml:"executorId,omitempty"`
	Id             *string   `json:"id,omitempty" xml:"id,omitempty"`
	InvolveMembers []*string `json:"involveMembers,omitempty" xml:"involveMembers,omitempty" type:"Repeated"`
	IsArchive      *bool     `json:"isArchive,omitempty" xml:"isArchive,omitempty"`
	IsDone         *bool     `json:"isDone,omitempty" xml:"isDone,omitempty"`
	Note           *string   `json:"note,omitempty" xml:"note,omitempty"`
	OrganizationId *string   `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	Priority       *int32    `json:"priority,omitempty" xml:"priority,omitempty"`
	Recurrence     []*string `json:"recurrence,omitempty" xml:"recurrence,omitempty" type:"Repeated"`
	SourceId       *string   `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// 2022-07-04T03:29:34.770Z
	StartDate *string   `json:"startDate,omitempty" xml:"startDate,omitempty"`
	TagIds    []*string `json:"tagIds,omitempty" xml:"tagIds,omitempty" type:"Repeated"`
	UniqueId  *int32    `json:"uniqueId,omitempty" xml:"uniqueId,omitempty"`
	// example:
	//
	// 2022-07-04T03:29:34.770Z
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
	Visible *string `json:"visible,omitempty" xml:"visible,omitempty"`
}

func (s GetFreeTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFreeTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFreeTaskResponseBodyResult) SetAccomplished(v string) *GetFreeTaskResponseBodyResult {
	s.Accomplished = &v
	return s
}

func (s *GetFreeTaskResponseBodyResult) SetAncestorIds(v []*string) *GetFreeTaskResponseBodyResult {
	s.AncestorIds = v
	return s
}

func (s *GetFreeTaskResponseBodyResult) SetContent(v string) *GetFreeTaskResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetFreeTaskResponseBodyResult) SetCreated(v string) *GetFreeTaskResponseBodyResult {
	s.Created = &v
	return s
}

func (s *GetFreeTaskResponseBodyResult) SetCreatorId(v string) *GetFreeTaskResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *GetFreeTaskResponseBodyResult) SetDueDate(v string) *GetFreeTaskResponseBodyResult {
	s.DueDate = &v
	return s
}

func (s *GetFreeTaskResponseBodyResult) SetExecutorId(v string) *GetFreeTaskResponseBodyResult {
	s.ExecutorId = &v
	return s
}

func (s *GetFreeTaskResponseBodyResult) SetId(v string) *GetFreeTaskResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetFreeTaskResponseBodyResult) SetInvolveMembers(v []*string) *GetFreeTaskResponseBodyResult {
	s.InvolveMembers = v
	return s
}

func (s *GetFreeTaskResponseBodyResult) SetIsArchive(v bool) *GetFreeTaskResponseBodyResult {
	s.IsArchive = &v
	return s
}

func (s *GetFreeTaskResponseBodyResult) SetIsDone(v bool) *GetFreeTaskResponseBodyResult {
	s.IsDone = &v
	return s
}

func (s *GetFreeTaskResponseBodyResult) SetNote(v string) *GetFreeTaskResponseBodyResult {
	s.Note = &v
	return s
}

func (s *GetFreeTaskResponseBodyResult) SetOrganizationId(v string) *GetFreeTaskResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *GetFreeTaskResponseBodyResult) SetPriority(v int32) *GetFreeTaskResponseBodyResult {
	s.Priority = &v
	return s
}

func (s *GetFreeTaskResponseBodyResult) SetRecurrence(v []*string) *GetFreeTaskResponseBodyResult {
	s.Recurrence = v
	return s
}

func (s *GetFreeTaskResponseBodyResult) SetSourceId(v string) *GetFreeTaskResponseBodyResult {
	s.SourceId = &v
	return s
}

func (s *GetFreeTaskResponseBodyResult) SetStartDate(v string) *GetFreeTaskResponseBodyResult {
	s.StartDate = &v
	return s
}

func (s *GetFreeTaskResponseBodyResult) SetTagIds(v []*string) *GetFreeTaskResponseBodyResult {
	s.TagIds = v
	return s
}

func (s *GetFreeTaskResponseBodyResult) SetUniqueId(v int32) *GetFreeTaskResponseBodyResult {
	s.UniqueId = &v
	return s
}

func (s *GetFreeTaskResponseBodyResult) SetUpdated(v string) *GetFreeTaskResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *GetFreeTaskResponseBodyResult) SetVisible(v string) *GetFreeTaskResponseBodyResult {
	s.Visible = &v
	return s
}

type GetFreeTaskResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFreeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFreeTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFreeTaskResponse) GoString() string {
	return s.String()
}

func (s *GetFreeTaskResponse) SetHeaders(v map[string]*string) *GetFreeTaskResponse {
	s.Headers = v
	return s
}

func (s *GetFreeTaskResponse) SetStatusCode(v int32) *GetFreeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFreeTaskResponse) SetBody(v *GetFreeTaskResponseBody) *GetFreeTaskResponse {
	s.Body = v
	return s
}

type GetProjectMembersV3Headers struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetProjectMembersV3Headers) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMembersV3Headers) GoString() string {
	return s.String()
}

func (s *GetProjectMembersV3Headers) SetCommonHeaders(v map[string]*string) *GetProjectMembersV3Headers {
	s.CommonHeaders = v
	return s
}

func (s *GetProjectMembersV3Headers) SetXAcsDingtalkAccessToken(v string) *GetProjectMembersV3Headers {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetProjectMembersV3Request struct {
	MaxResults    *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken     *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ProjectRoleId *string `json:"projectRoleId,omitempty" xml:"projectRoleId,omitempty"`
	UserIds       *string `json:"userIds,omitempty" xml:"userIds,omitempty"`
}

func (s GetProjectMembersV3Request) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMembersV3Request) GoString() string {
	return s.String()
}

func (s *GetProjectMembersV3Request) SetMaxResults(v int32) *GetProjectMembersV3Request {
	s.MaxResults = &v
	return s
}

func (s *GetProjectMembersV3Request) SetNextToken(v string) *GetProjectMembersV3Request {
	s.NextToken = &v
	return s
}

func (s *GetProjectMembersV3Request) SetProjectRoleId(v string) *GetProjectMembersV3Request {
	s.ProjectRoleId = &v
	return s
}

func (s *GetProjectMembersV3Request) SetUserIds(v string) *GetProjectMembersV3Request {
	s.UserIds = &v
	return s
}

type GetProjectMembersV3ResponseBody struct {
	// example:
	//
	// f279e812-e431-428d-846d-cxxxxxx
	NextToken *string                                  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*GetProjectMembersV3ResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetProjectMembersV3ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMembersV3ResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectMembersV3ResponseBody) SetNextToken(v string) *GetProjectMembersV3ResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetProjectMembersV3ResponseBody) SetRequestId(v string) *GetProjectMembersV3ResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectMembersV3ResponseBody) SetResult(v []*GetProjectMembersV3ResponseBodyResult) *GetProjectMembersV3ResponseBody {
	s.Result = v
	return s
}

type GetProjectMembersV3ResponseBodyResult struct {
	Role   *int32  `json:"role,omitempty" xml:"role,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetProjectMembersV3ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMembersV3ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetProjectMembersV3ResponseBodyResult) SetRole(v int32) *GetProjectMembersV3ResponseBodyResult {
	s.Role = &v
	return s
}

func (s *GetProjectMembersV3ResponseBodyResult) SetUserId(v string) *GetProjectMembersV3ResponseBodyResult {
	s.UserId = &v
	return s
}

type GetProjectMembersV3Response struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectMembersV3ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectMembersV3Response) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMembersV3Response) GoString() string {
	return s.String()
}

func (s *GetProjectMembersV3Response) SetHeaders(v map[string]*string) *GetProjectMembersV3Response {
	s.Headers = v
	return s
}

func (s *GetProjectMembersV3Response) SetStatusCode(v int32) *GetProjectMembersV3Response {
	s.StatusCode = &v
	return s
}

func (s *GetProjectMembersV3Response) SetBody(v *GetProjectMembersV3ResponseBody) *GetProjectMembersV3Response {
	s.Body = v
	return s
}

type GetProjectRolesV3Headers struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetProjectRolesV3Headers) String() string {
	return tea.Prettify(s)
}

func (s GetProjectRolesV3Headers) GoString() string {
	return s.String()
}

func (s *GetProjectRolesV3Headers) SetCommonHeaders(v map[string]*string) *GetProjectRolesV3Headers {
	s.CommonHeaders = v
	return s
}

func (s *GetProjectRolesV3Headers) SetXAcsDingtalkAccessToken(v string) *GetProjectRolesV3Headers {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetProjectRolesV3Request struct {
	IncludeHidden *bool   `json:"includeHidden,omitempty" xml:"includeHidden,omitempty"`
	Level         *int64  `json:"level,omitempty" xml:"level,omitempty"`
	MaxResults    *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken     *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetProjectRolesV3Request) String() string {
	return tea.Prettify(s)
}

func (s GetProjectRolesV3Request) GoString() string {
	return s.String()
}

func (s *GetProjectRolesV3Request) SetIncludeHidden(v bool) *GetProjectRolesV3Request {
	s.IncludeHidden = &v
	return s
}

func (s *GetProjectRolesV3Request) SetLevel(v int64) *GetProjectRolesV3Request {
	s.Level = &v
	return s
}

func (s *GetProjectRolesV3Request) SetMaxResults(v int32) *GetProjectRolesV3Request {
	s.MaxResults = &v
	return s
}

func (s *GetProjectRolesV3Request) SetNextToken(v string) *GetProjectRolesV3Request {
	s.NextToken = &v
	return s
}

type GetProjectRolesV3ResponseBody struct {
	// example:
	//
	// f279e812-e431-428d-846d-cxxxxxx
	NextToken *string                                `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*GetProjectRolesV3ResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetProjectRolesV3ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectRolesV3ResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectRolesV3ResponseBody) SetNextToken(v string) *GetProjectRolesV3ResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetProjectRolesV3ResponseBody) SetRequestId(v string) *GetProjectRolesV3ResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectRolesV3ResponseBody) SetResult(v []*GetProjectRolesV3ResponseBodyResult) *GetProjectRolesV3ResponseBody {
	s.Result = v
	return s
}

type GetProjectRolesV3ResponseBodyResult struct {
	Display       *bool   `json:"display,omitempty" xml:"display,omitempty"`
	Id            *string `json:"id,omitempty" xml:"id,omitempty"`
	IsDefaultRole *bool   `json:"isDefaultRole,omitempty" xml:"isDefaultRole,omitempty"`
	Level         *int32  `json:"level,omitempty" xml:"level,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	OriginalId    *string `json:"originalId,omitempty" xml:"originalId,omitempty"`
}

func (s GetProjectRolesV3ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetProjectRolesV3ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetProjectRolesV3ResponseBodyResult) SetDisplay(v bool) *GetProjectRolesV3ResponseBodyResult {
	s.Display = &v
	return s
}

func (s *GetProjectRolesV3ResponseBodyResult) SetId(v string) *GetProjectRolesV3ResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetProjectRolesV3ResponseBodyResult) SetIsDefaultRole(v bool) *GetProjectRolesV3ResponseBodyResult {
	s.IsDefaultRole = &v
	return s
}

func (s *GetProjectRolesV3ResponseBodyResult) SetLevel(v int32) *GetProjectRolesV3ResponseBodyResult {
	s.Level = &v
	return s
}

func (s *GetProjectRolesV3ResponseBodyResult) SetName(v string) *GetProjectRolesV3ResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetProjectRolesV3ResponseBodyResult) SetOriginalId(v string) *GetProjectRolesV3ResponseBodyResult {
	s.OriginalId = &v
	return s
}

type GetProjectRolesV3Response struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectRolesV3ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectRolesV3Response) String() string {
	return tea.Prettify(s)
}

func (s GetProjectRolesV3Response) GoString() string {
	return s.String()
}

func (s *GetProjectRolesV3Response) SetHeaders(v map[string]*string) *GetProjectRolesV3Response {
	s.Headers = v
	return s
}

func (s *GetProjectRolesV3Response) SetStatusCode(v int32) *GetProjectRolesV3Response {
	s.StatusCode = &v
	return s
}

func (s *GetProjectRolesV3Response) SetBody(v *GetProjectRolesV3ResponseBody) *GetProjectRolesV3Response {
	s.Body = v
	return s
}

type GetStaredProjectsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetStaredProjectsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetStaredProjectsHeaders) GoString() string {
	return s.String()
}

func (s *GetStaredProjectsHeaders) SetCommonHeaders(v map[string]*string) *GetStaredProjectsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetStaredProjectsHeaders) SetXAcsDingtalkAccessToken(v string) *GetStaredProjectsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetStaredProjectsRequest struct {
	// example:
	//
	// 10
	MaxResults *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// updated:desc
	SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
}

func (s GetStaredProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStaredProjectsRequest) GoString() string {
	return s.String()
}

func (s *GetStaredProjectsRequest) SetMaxResults(v int64) *GetStaredProjectsRequest {
	s.MaxResults = &v
	return s
}

func (s *GetStaredProjectsRequest) SetNextToken(v string) *GetStaredProjectsRequest {
	s.NextToken = &v
	return s
}

func (s *GetStaredProjectsRequest) SetSortBy(v string) *GetStaredProjectsRequest {
	s.SortBy = &v
	return s
}

type GetStaredProjectsResponseBody struct {
	// example:
	//
	// f279e812-e431-428d-846d-cxxxxxx
	NextToken *string   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Result    []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetStaredProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStaredProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *GetStaredProjectsResponseBody) SetNextToken(v string) *GetStaredProjectsResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetStaredProjectsResponseBody) SetResult(v []*string) *GetStaredProjectsResponseBody {
	s.Result = v
	return s
}

type GetStaredProjectsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStaredProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStaredProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStaredProjectsResponse) GoString() string {
	return s.String()
}

func (s *GetStaredProjectsResponse) SetHeaders(v map[string]*string) *GetStaredProjectsResponse {
	s.Headers = v
	return s
}

func (s *GetStaredProjectsResponse) SetStatusCode(v int32) *GetStaredProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStaredProjectsResponse) SetBody(v *GetStaredProjectsResponseBody) *GetStaredProjectsResponse {
	s.Body = v
	return s
}

type GetTbUserIdByDingUserIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTbUserIdByDingUserIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTbUserIdByDingUserIdHeaders) GoString() string {
	return s.String()
}

func (s *GetTbUserIdByDingUserIdHeaders) SetCommonHeaders(v map[string]*string) *GetTbUserIdByDingUserIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTbUserIdByDingUserIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetTbUserIdByDingUserIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTbUserIdByDingUserIdRequest struct {
	// This parameter is required.
	DingUserIds *string `json:"dingUserIds,omitempty" xml:"dingUserIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0517xxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetTbUserIdByDingUserIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTbUserIdByDingUserIdRequest) GoString() string {
	return s.String()
}

func (s *GetTbUserIdByDingUserIdRequest) SetDingUserIds(v string) *GetTbUserIdByDingUserIdRequest {
	s.DingUserIds = &v
	return s
}

func (s *GetTbUserIdByDingUserIdRequest) SetUserId(v string) *GetTbUserIdByDingUserIdRequest {
	s.UserId = &v
	return s
}

type GetTbUserIdByDingUserIdResponseBody struct {
	RequestId *string                                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*GetTbUserIdByDingUserIdResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetTbUserIdByDingUserIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTbUserIdByDingUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetTbUserIdByDingUserIdResponseBody) SetRequestId(v string) *GetTbUserIdByDingUserIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTbUserIdByDingUserIdResponseBody) SetResult(v []*GetTbUserIdByDingUserIdResponseBodyResult) *GetTbUserIdByDingUserIdResponseBody {
	s.Result = v
	return s
}

type GetTbUserIdByDingUserIdResponseBodyResult struct {
	DingtalkUserId *string `json:"dingtalkUserId,omitempty" xml:"dingtalkUserId,omitempty"`
	TbUserId       *string `json:"tbUserId,omitempty" xml:"tbUserId,omitempty"`
}

func (s GetTbUserIdByDingUserIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetTbUserIdByDingUserIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetTbUserIdByDingUserIdResponseBodyResult) SetDingtalkUserId(v string) *GetTbUserIdByDingUserIdResponseBodyResult {
	s.DingtalkUserId = &v
	return s
}

func (s *GetTbUserIdByDingUserIdResponseBodyResult) SetTbUserId(v string) *GetTbUserIdByDingUserIdResponseBodyResult {
	s.TbUserId = &v
	return s
}

type GetTbUserIdByDingUserIdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTbUserIdByDingUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTbUserIdByDingUserIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTbUserIdByDingUserIdResponse) GoString() string {
	return s.String()
}

func (s *GetTbUserIdByDingUserIdResponse) SetHeaders(v map[string]*string) *GetTbUserIdByDingUserIdResponse {
	s.Headers = v
	return s
}

func (s *GetTbUserIdByDingUserIdResponse) SetStatusCode(v int32) *GetTbUserIdByDingUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTbUserIdByDingUserIdResponse) SetBody(v *GetTbUserIdByDingUserIdResponseBody) *GetTbUserIdByDingUserIdResponse {
	s.Body = v
	return s
}

type GetThingOrgIdByDingOrgIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetThingOrgIdByDingOrgIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetThingOrgIdByDingOrgIdHeaders) GoString() string {
	return s.String()
}

func (s *GetThingOrgIdByDingOrgIdHeaders) SetCommonHeaders(v map[string]*string) *GetThingOrgIdByDingOrgIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetThingOrgIdByDingOrgIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetThingOrgIdByDingOrgIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetThingOrgIdByDingOrgIdResponseBody struct {
	// This parameter is required.
	Result *GetThingOrgIdByDingOrgIdResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetThingOrgIdByDingOrgIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetThingOrgIdByDingOrgIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetThingOrgIdByDingOrgIdResponseBody) SetResult(v *GetThingOrgIdByDingOrgIdResponseBodyResult) *GetThingOrgIdByDingOrgIdResponseBody {
	s.Result = v
	return s
}

type GetThingOrgIdByDingOrgIdResponseBodyResult struct {
	// This parameter is required.
	//
	// example:
	//
	// 50c32afae8cf1439xxxx
	TbOrganizationId *string `json:"tbOrganizationId,omitempty" xml:"tbOrganizationId,omitempty"`
}

func (s GetThingOrgIdByDingOrgIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetThingOrgIdByDingOrgIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetThingOrgIdByDingOrgIdResponseBodyResult) SetTbOrganizationId(v string) *GetThingOrgIdByDingOrgIdResponseBodyResult {
	s.TbOrganizationId = &v
	return s
}

type GetThingOrgIdByDingOrgIdResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetThingOrgIdByDingOrgIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetThingOrgIdByDingOrgIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetThingOrgIdByDingOrgIdResponse) GoString() string {
	return s.String()
}

func (s *GetThingOrgIdByDingOrgIdResponse) SetHeaders(v map[string]*string) *GetThingOrgIdByDingOrgIdResponse {
	s.Headers = v
	return s
}

func (s *GetThingOrgIdByDingOrgIdResponse) SetStatusCode(v int32) *GetThingOrgIdByDingOrgIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetThingOrgIdByDingOrgIdResponse) SetBody(v *GetThingOrgIdByDingOrgIdResponseBody) *GetThingOrgIdByDingOrgIdResponse {
	s.Body = v
	return s
}

type GetUserJoinedProjectsV3Headers struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserJoinedProjectsV3Headers) String() string {
	return tea.Prettify(s)
}

func (s GetUserJoinedProjectsV3Headers) GoString() string {
	return s.String()
}

func (s *GetUserJoinedProjectsV3Headers) SetCommonHeaders(v map[string]*string) *GetUserJoinedProjectsV3Headers {
	s.CommonHeaders = v
	return s
}

func (s *GetUserJoinedProjectsV3Headers) SetXAcsDingtalkAccessToken(v string) *GetUserJoinedProjectsV3Headers {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserJoinedProjectsV3Request struct {
	MaxResults        *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken         *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ProjectIds        *string `json:"projectIds,omitempty" xml:"projectIds,omitempty"`
	ProjectRoleLevels *string `json:"projectRoleLevels,omitempty" xml:"projectRoleLevels,omitempty"`
	SortBy            *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
}

func (s GetUserJoinedProjectsV3Request) String() string {
	return tea.Prettify(s)
}

func (s GetUserJoinedProjectsV3Request) GoString() string {
	return s.String()
}

func (s *GetUserJoinedProjectsV3Request) SetMaxResults(v int32) *GetUserJoinedProjectsV3Request {
	s.MaxResults = &v
	return s
}

func (s *GetUserJoinedProjectsV3Request) SetNextToken(v string) *GetUserJoinedProjectsV3Request {
	s.NextToken = &v
	return s
}

func (s *GetUserJoinedProjectsV3Request) SetProjectIds(v string) *GetUserJoinedProjectsV3Request {
	s.ProjectIds = &v
	return s
}

func (s *GetUserJoinedProjectsV3Request) SetProjectRoleLevels(v string) *GetUserJoinedProjectsV3Request {
	s.ProjectRoleLevels = &v
	return s
}

func (s *GetUserJoinedProjectsV3Request) SetSortBy(v string) *GetUserJoinedProjectsV3Request {
	s.SortBy = &v
	return s
}

type GetUserJoinedProjectsV3ResponseBody struct {
	// example:
	//
	// f279e812-e431-428d-846d-cxxxxxx
	NextToken *string   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetUserJoinedProjectsV3ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserJoinedProjectsV3ResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserJoinedProjectsV3ResponseBody) SetNextToken(v string) *GetUserJoinedProjectsV3ResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetUserJoinedProjectsV3ResponseBody) SetRequestId(v string) *GetUserJoinedProjectsV3ResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserJoinedProjectsV3ResponseBody) SetResult(v []*string) *GetUserJoinedProjectsV3ResponseBody {
	s.Result = v
	return s
}

type GetUserJoinedProjectsV3Response struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserJoinedProjectsV3ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserJoinedProjectsV3Response) String() string {
	return tea.Prettify(s)
}

func (s GetUserJoinedProjectsV3Response) GoString() string {
	return s.String()
}

func (s *GetUserJoinedProjectsV3Response) SetHeaders(v map[string]*string) *GetUserJoinedProjectsV3Response {
	s.Headers = v
	return s
}

func (s *GetUserJoinedProjectsV3Response) SetStatusCode(v int32) *GetUserJoinedProjectsV3Response {
	s.StatusCode = &v
	return s
}

func (s *GetUserJoinedProjectsV3Response) SetBody(v *GetUserJoinedProjectsV3ResponseBody) *GetUserJoinedProjectsV3Response {
	s.Body = v
	return s
}

type ListAllTaskViewHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListAllTaskViewHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListAllTaskViewHeaders) GoString() string {
	return s.String()
}

func (s *ListAllTaskViewHeaders) SetCommonHeaders(v map[string]*string) *ListAllTaskViewHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAllTaskViewHeaders) SetXAcsDingtalkAccessToken(v string) *ListAllTaskViewHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListAllTaskViewResponseBody struct {
	Result *ListAllTaskViewResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ListAllTaskViewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAllTaskViewResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllTaskViewResponseBody) SetResult(v *ListAllTaskViewResponseBodyResult) *ListAllTaskViewResponseBody {
	s.Result = v
	return s
}

type ListAllTaskViewResponseBodyResult struct {
	BoundToObjectId *string `json:"boundToObjectId,omitempty" xml:"boundToObjectId,omitempty"`
	// example:
	//
	// user
	BoundToObjectType *string `json:"boundToObjectType,omitempty" xml:"boundToObjectType,omitempty"`
	// example:
	//
	// 2022-07-04T03:29:34.770Z
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// example:
	//
	// 5f687406f05b283425ea8f6f
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// xxxx
	Description *string                                     `json:"description,omitempty" xml:"description,omitempty"`
	Filter      *ListAllTaskViewResponseBodyResultFilter    `json:"filter,omitempty" xml:"filter,omitempty" type:"Struct"`
	GroupType   *ListAllTaskViewResponseBodyResultGroupType `json:"groupType,omitempty" xml:"groupType,omitempty" type:"Struct"`
	Id          *string                                     `json:"id,omitempty" xml:"id,omitempty"`
	IsDeleted   *bool                                       `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	// example:
	//
	// x项目
	Name      *string                                     `json:"name,omitempty" xml:"name,omitempty"`
	OrderType *ListAllTaskViewResponseBodyResultOrderType `json:"orderType,omitempty" xml:"orderType,omitempty" type:"Struct"`
	// example:
	//
	// 6139cd1aba5b128516ec8c86
	OrganizationId *string                                       `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	ShowType       *ListAllTaskViewResponseBodyResultShowType    `json:"showType,omitempty" xml:"showType,omitempty" type:"Struct"`
	ToolbarInfo    *ListAllTaskViewResponseBodyResultToolbarInfo `json:"toolbarInfo,omitempty" xml:"toolbarInfo,omitempty" type:"Struct"`
	Tql            *string                                       `json:"tql,omitempty" xml:"tql,omitempty"`
	// example:
	//
	// 2022-07-04T03:29:34.770Z
	Updated     *string                                       `json:"updated,omitempty" xml:"updated,omitempty"`
	ViewSetting *ListAllTaskViewResponseBodyResultViewSetting `json:"viewSetting,omitempty" xml:"viewSetting,omitempty" type:"Struct"`
}

func (s ListAllTaskViewResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAllTaskViewResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAllTaskViewResponseBodyResult) SetBoundToObjectId(v string) *ListAllTaskViewResponseBodyResult {
	s.BoundToObjectId = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResult) SetBoundToObjectType(v string) *ListAllTaskViewResponseBodyResult {
	s.BoundToObjectType = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResult) SetCreated(v string) *ListAllTaskViewResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResult) SetCreatorId(v string) *ListAllTaskViewResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResult) SetDescription(v string) *ListAllTaskViewResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResult) SetFilter(v *ListAllTaskViewResponseBodyResultFilter) *ListAllTaskViewResponseBodyResult {
	s.Filter = v
	return s
}

func (s *ListAllTaskViewResponseBodyResult) SetGroupType(v *ListAllTaskViewResponseBodyResultGroupType) *ListAllTaskViewResponseBodyResult {
	s.GroupType = v
	return s
}

func (s *ListAllTaskViewResponseBodyResult) SetId(v string) *ListAllTaskViewResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResult) SetIsDeleted(v bool) *ListAllTaskViewResponseBodyResult {
	s.IsDeleted = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResult) SetName(v string) *ListAllTaskViewResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResult) SetOrderType(v *ListAllTaskViewResponseBodyResultOrderType) *ListAllTaskViewResponseBodyResult {
	s.OrderType = v
	return s
}

func (s *ListAllTaskViewResponseBodyResult) SetOrganizationId(v string) *ListAllTaskViewResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResult) SetShowType(v *ListAllTaskViewResponseBodyResultShowType) *ListAllTaskViewResponseBodyResult {
	s.ShowType = v
	return s
}

func (s *ListAllTaskViewResponseBodyResult) SetToolbarInfo(v *ListAllTaskViewResponseBodyResultToolbarInfo) *ListAllTaskViewResponseBodyResult {
	s.ToolbarInfo = v
	return s
}

func (s *ListAllTaskViewResponseBodyResult) SetTql(v string) *ListAllTaskViewResponseBodyResult {
	s.Tql = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResult) SetUpdated(v string) *ListAllTaskViewResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResult) SetViewSetting(v *ListAllTaskViewResponseBodyResultViewSetting) *ListAllTaskViewResponseBodyResult {
	s.ViewSetting = v
	return s
}

type ListAllTaskViewResponseBodyResultFilter struct {
	Conditions          []*ListAllTaskViewResponseBodyResultFilterConditions          `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	InvisibleConditions []*ListAllTaskViewResponseBodyResultFilterInvisibleConditions `json:"invisibleConditions,omitempty" xml:"invisibleConditions,omitempty" type:"Repeated"`
	Logic               *string                                                       `json:"logic,omitempty" xml:"logic,omitempty"`
	LogicDisabled       *bool                                                         `json:"logicDisabled,omitempty" xml:"logicDisabled,omitempty"`
	Q                   *string                                                       `json:"q,omitempty" xml:"q,omitempty"`
}

func (s ListAllTaskViewResponseBodyResultFilter) String() string {
	return tea.Prettify(s)
}

func (s ListAllTaskViewResponseBodyResultFilter) GoString() string {
	return s.String()
}

func (s *ListAllTaskViewResponseBodyResultFilter) SetConditions(v []*ListAllTaskViewResponseBodyResultFilterConditions) *ListAllTaskViewResponseBodyResultFilter {
	s.Conditions = v
	return s
}

func (s *ListAllTaskViewResponseBodyResultFilter) SetInvisibleConditions(v []*ListAllTaskViewResponseBodyResultFilterInvisibleConditions) *ListAllTaskViewResponseBodyResultFilter {
	s.InvisibleConditions = v
	return s
}

func (s *ListAllTaskViewResponseBodyResultFilter) SetLogic(v string) *ListAllTaskViewResponseBodyResultFilter {
	s.Logic = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultFilter) SetLogicDisabled(v bool) *ListAllTaskViewResponseBodyResultFilter {
	s.LogicDisabled = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultFilter) SetQ(v string) *ListAllTaskViewResponseBodyResultFilter {
	s.Q = &v
	return s
}

type ListAllTaskViewResponseBodyResultFilterConditions struct {
	Fixed  *bool                                                      `json:"fixed,omitempty" xml:"fixed,omitempty"`
	Key    *string                                                    `json:"key,omitempty" xml:"key,omitempty"`
	Op     *ListAllTaskViewResponseBodyResultFilterConditionsOp       `json:"op,omitempty" xml:"op,omitempty" type:"Struct"`
	Values []*ListAllTaskViewResponseBodyResultFilterConditionsValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s ListAllTaskViewResponseBodyResultFilterConditions) String() string {
	return tea.Prettify(s)
}

func (s ListAllTaskViewResponseBodyResultFilterConditions) GoString() string {
	return s.String()
}

func (s *ListAllTaskViewResponseBodyResultFilterConditions) SetFixed(v bool) *ListAllTaskViewResponseBodyResultFilterConditions {
	s.Fixed = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultFilterConditions) SetKey(v string) *ListAllTaskViewResponseBodyResultFilterConditions {
	s.Key = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultFilterConditions) SetOp(v *ListAllTaskViewResponseBodyResultFilterConditionsOp) *ListAllTaskViewResponseBodyResultFilterConditions {
	s.Op = v
	return s
}

func (s *ListAllTaskViewResponseBodyResultFilterConditions) SetValues(v []*ListAllTaskViewResponseBodyResultFilterConditionsValues) *ListAllTaskViewResponseBodyResultFilterConditions {
	s.Values = v
	return s
}

type ListAllTaskViewResponseBodyResultFilterConditionsOp struct {
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListAllTaskViewResponseBodyResultFilterConditionsOp) String() string {
	return tea.Prettify(s)
}

func (s ListAllTaskViewResponseBodyResultFilterConditionsOp) GoString() string {
	return s.String()
}

func (s *ListAllTaskViewResponseBodyResultFilterConditionsOp) SetValue(v string) *ListAllTaskViewResponseBodyResultFilterConditionsOp {
	s.Value = &v
	return s
}

type ListAllTaskViewResponseBodyResultFilterConditionsValues struct {
	Deep  *string `json:"deep,omitempty" xml:"deep,omitempty"`
	Id    *string `json:"id,omitempty" xml:"id,omitempty"`
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListAllTaskViewResponseBodyResultFilterConditionsValues) String() string {
	return tea.Prettify(s)
}

func (s ListAllTaskViewResponseBodyResultFilterConditionsValues) GoString() string {
	return s.String()
}

func (s *ListAllTaskViewResponseBodyResultFilterConditionsValues) SetDeep(v string) *ListAllTaskViewResponseBodyResultFilterConditionsValues {
	s.Deep = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultFilterConditionsValues) SetId(v string) *ListAllTaskViewResponseBodyResultFilterConditionsValues {
	s.Id = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultFilterConditionsValues) SetLabel(v string) *ListAllTaskViewResponseBodyResultFilterConditionsValues {
	s.Label = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultFilterConditionsValues) SetName(v string) *ListAllTaskViewResponseBodyResultFilterConditionsValues {
	s.Name = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultFilterConditionsValues) SetValue(v string) *ListAllTaskViewResponseBodyResultFilterConditionsValues {
	s.Value = &v
	return s
}

type ListAllTaskViewResponseBodyResultFilterInvisibleConditions struct {
	Fixed  *bool                                                               `json:"fixed,omitempty" xml:"fixed,omitempty"`
	Key    *string                                                             `json:"key,omitempty" xml:"key,omitempty"`
	Op     *ListAllTaskViewResponseBodyResultFilterInvisibleConditionsOp       `json:"op,omitempty" xml:"op,omitempty" type:"Struct"`
	Values []*ListAllTaskViewResponseBodyResultFilterInvisibleConditionsValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s ListAllTaskViewResponseBodyResultFilterInvisibleConditions) String() string {
	return tea.Prettify(s)
}

func (s ListAllTaskViewResponseBodyResultFilterInvisibleConditions) GoString() string {
	return s.String()
}

func (s *ListAllTaskViewResponseBodyResultFilterInvisibleConditions) SetFixed(v bool) *ListAllTaskViewResponseBodyResultFilterInvisibleConditions {
	s.Fixed = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultFilterInvisibleConditions) SetKey(v string) *ListAllTaskViewResponseBodyResultFilterInvisibleConditions {
	s.Key = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultFilterInvisibleConditions) SetOp(v *ListAllTaskViewResponseBodyResultFilterInvisibleConditionsOp) *ListAllTaskViewResponseBodyResultFilterInvisibleConditions {
	s.Op = v
	return s
}

func (s *ListAllTaskViewResponseBodyResultFilterInvisibleConditions) SetValues(v []*ListAllTaskViewResponseBodyResultFilterInvisibleConditionsValues) *ListAllTaskViewResponseBodyResultFilterInvisibleConditions {
	s.Values = v
	return s
}

type ListAllTaskViewResponseBodyResultFilterInvisibleConditionsOp struct {
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListAllTaskViewResponseBodyResultFilterInvisibleConditionsOp) String() string {
	return tea.Prettify(s)
}

func (s ListAllTaskViewResponseBodyResultFilterInvisibleConditionsOp) GoString() string {
	return s.String()
}

func (s *ListAllTaskViewResponseBodyResultFilterInvisibleConditionsOp) SetValue(v string) *ListAllTaskViewResponseBodyResultFilterInvisibleConditionsOp {
	s.Value = &v
	return s
}

type ListAllTaskViewResponseBodyResultFilterInvisibleConditionsValues struct {
	Deep  *string `json:"deep,omitempty" xml:"deep,omitempty"`
	Id    *string `json:"id,omitempty" xml:"id,omitempty"`
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListAllTaskViewResponseBodyResultFilterInvisibleConditionsValues) String() string {
	return tea.Prettify(s)
}

func (s ListAllTaskViewResponseBodyResultFilterInvisibleConditionsValues) GoString() string {
	return s.String()
}

func (s *ListAllTaskViewResponseBodyResultFilterInvisibleConditionsValues) SetDeep(v string) *ListAllTaskViewResponseBodyResultFilterInvisibleConditionsValues {
	s.Deep = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultFilterInvisibleConditionsValues) SetId(v string) *ListAllTaskViewResponseBodyResultFilterInvisibleConditionsValues {
	s.Id = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultFilterInvisibleConditionsValues) SetLabel(v string) *ListAllTaskViewResponseBodyResultFilterInvisibleConditionsValues {
	s.Label = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultFilterInvisibleConditionsValues) SetName(v string) *ListAllTaskViewResponseBodyResultFilterInvisibleConditionsValues {
	s.Name = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultFilterInvisibleConditionsValues) SetValue(v string) *ListAllTaskViewResponseBodyResultFilterInvisibleConditionsValues {
	s.Value = &v
	return s
}

type ListAllTaskViewResponseBodyResultGroupType struct {
	CanCreateGroup *bool   `json:"canCreateGroup,omitempty" xml:"canCreateGroup,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	Value          *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListAllTaskViewResponseBodyResultGroupType) String() string {
	return tea.Prettify(s)
}

func (s ListAllTaskViewResponseBodyResultGroupType) GoString() string {
	return s.String()
}

func (s *ListAllTaskViewResponseBodyResultGroupType) SetCanCreateGroup(v bool) *ListAllTaskViewResponseBodyResultGroupType {
	s.CanCreateGroup = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultGroupType) SetName(v string) *ListAllTaskViewResponseBodyResultGroupType {
	s.Name = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultGroupType) SetValue(v string) *ListAllTaskViewResponseBodyResultGroupType {
	s.Value = &v
	return s
}

type ListAllTaskViewResponseBodyResultOrderType struct {
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Value     *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListAllTaskViewResponseBodyResultOrderType) String() string {
	return tea.Prettify(s)
}

func (s ListAllTaskViewResponseBodyResultOrderType) GoString() string {
	return s.String()
}

func (s *ListAllTaskViewResponseBodyResultOrderType) SetDirection(v string) *ListAllTaskViewResponseBodyResultOrderType {
	s.Direction = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultOrderType) SetName(v string) *ListAllTaskViewResponseBodyResultOrderType {
	s.Name = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultOrderType) SetValue(v string) *ListAllTaskViewResponseBodyResultOrderType {
	s.Value = &v
	return s
}

type ListAllTaskViewResponseBodyResultShowType struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListAllTaskViewResponseBodyResultShowType) String() string {
	return tea.Prettify(s)
}

func (s ListAllTaskViewResponseBodyResultShowType) GoString() string {
	return s.String()
}

func (s *ListAllTaskViewResponseBodyResultShowType) SetName(v string) *ListAllTaskViewResponseBodyResultShowType {
	s.Name = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultShowType) SetValue(v string) *ListAllTaskViewResponseBodyResultShowType {
	s.Value = &v
	return s
}

type ListAllTaskViewResponseBodyResultToolbarInfo struct {
	GroupTypes []*ListAllTaskViewResponseBodyResultToolbarInfoGroupTypes `json:"groupTypes,omitempty" xml:"groupTypes,omitempty" type:"Repeated"`
	OrderTypes []*ListAllTaskViewResponseBodyResultToolbarInfoOrderTypes `json:"orderTypes,omitempty" xml:"orderTypes,omitempty" type:"Repeated"`
	ShowTypes  []*ListAllTaskViewResponseBodyResultToolbarInfoShowTypes  `json:"showTypes,omitempty" xml:"showTypes,omitempty" type:"Repeated"`
}

func (s ListAllTaskViewResponseBodyResultToolbarInfo) String() string {
	return tea.Prettify(s)
}

func (s ListAllTaskViewResponseBodyResultToolbarInfo) GoString() string {
	return s.String()
}

func (s *ListAllTaskViewResponseBodyResultToolbarInfo) SetGroupTypes(v []*ListAllTaskViewResponseBodyResultToolbarInfoGroupTypes) *ListAllTaskViewResponseBodyResultToolbarInfo {
	s.GroupTypes = v
	return s
}

func (s *ListAllTaskViewResponseBodyResultToolbarInfo) SetOrderTypes(v []*ListAllTaskViewResponseBodyResultToolbarInfoOrderTypes) *ListAllTaskViewResponseBodyResultToolbarInfo {
	s.OrderTypes = v
	return s
}

func (s *ListAllTaskViewResponseBodyResultToolbarInfo) SetShowTypes(v []*ListAllTaskViewResponseBodyResultToolbarInfoShowTypes) *ListAllTaskViewResponseBodyResultToolbarInfo {
	s.ShowTypes = v
	return s
}

type ListAllTaskViewResponseBodyResultToolbarInfoGroupTypes struct {
	CanCreateGroup *bool                                                          `json:"canCreateGroup,omitempty" xml:"canCreateGroup,omitempty"`
	Name           *string                                                        `json:"name,omitempty" xml:"name,omitempty"`
	ServiceName    *string                                                        `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	Setting        *ListAllTaskViewResponseBodyResultToolbarInfoGroupTypesSetting `json:"setting,omitempty" xml:"setting,omitempty" type:"Struct"`
	Value          *string                                                        `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListAllTaskViewResponseBodyResultToolbarInfoGroupTypes) String() string {
	return tea.Prettify(s)
}

func (s ListAllTaskViewResponseBodyResultToolbarInfoGroupTypes) GoString() string {
	return s.String()
}

func (s *ListAllTaskViewResponseBodyResultToolbarInfoGroupTypes) SetCanCreateGroup(v bool) *ListAllTaskViewResponseBodyResultToolbarInfoGroupTypes {
	s.CanCreateGroup = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultToolbarInfoGroupTypes) SetName(v string) *ListAllTaskViewResponseBodyResultToolbarInfoGroupTypes {
	s.Name = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultToolbarInfoGroupTypes) SetServiceName(v string) *ListAllTaskViewResponseBodyResultToolbarInfoGroupTypes {
	s.ServiceName = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultToolbarInfoGroupTypes) SetSetting(v *ListAllTaskViewResponseBodyResultToolbarInfoGroupTypesSetting) *ListAllTaskViewResponseBodyResultToolbarInfoGroupTypes {
	s.Setting = v
	return s
}

func (s *ListAllTaskViewResponseBodyResultToolbarInfoGroupTypes) SetValue(v string) *ListAllTaskViewResponseBodyResultToolbarInfoGroupTypes {
	s.Value = &v
	return s
}

type ListAllTaskViewResponseBodyResultToolbarInfoGroupTypesSetting struct {
	DateType  *string `json:"dateType,omitempty" xml:"dateType,omitempty"`
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
}

func (s ListAllTaskViewResponseBodyResultToolbarInfoGroupTypesSetting) String() string {
	return tea.Prettify(s)
}

func (s ListAllTaskViewResponseBodyResultToolbarInfoGroupTypesSetting) GoString() string {
	return s.String()
}

func (s *ListAllTaskViewResponseBodyResultToolbarInfoGroupTypesSetting) SetDateType(v string) *ListAllTaskViewResponseBodyResultToolbarInfoGroupTypesSetting {
	s.DateType = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultToolbarInfoGroupTypesSetting) SetFieldName(v string) *ListAllTaskViewResponseBodyResultToolbarInfoGroupTypesSetting {
	s.FieldName = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultToolbarInfoGroupTypesSetting) SetFieldType(v string) *ListAllTaskViewResponseBodyResultToolbarInfoGroupTypesSetting {
	s.FieldType = &v
	return s
}

type ListAllTaskViewResponseBodyResultToolbarInfoOrderTypes struct {
	Direction        *string `json:"direction,omitempty" xml:"direction,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	SupportDirection *string `json:"supportDirection,omitempty" xml:"supportDirection,omitempty"`
	Value            *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListAllTaskViewResponseBodyResultToolbarInfoOrderTypes) String() string {
	return tea.Prettify(s)
}

func (s ListAllTaskViewResponseBodyResultToolbarInfoOrderTypes) GoString() string {
	return s.String()
}

func (s *ListAllTaskViewResponseBodyResultToolbarInfoOrderTypes) SetDirection(v string) *ListAllTaskViewResponseBodyResultToolbarInfoOrderTypes {
	s.Direction = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultToolbarInfoOrderTypes) SetName(v string) *ListAllTaskViewResponseBodyResultToolbarInfoOrderTypes {
	s.Name = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultToolbarInfoOrderTypes) SetSupportDirection(v string) *ListAllTaskViewResponseBodyResultToolbarInfoOrderTypes {
	s.SupportDirection = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultToolbarInfoOrderTypes) SetValue(v string) *ListAllTaskViewResponseBodyResultToolbarInfoOrderTypes {
	s.Value = &v
	return s
}

type ListAllTaskViewResponseBodyResultToolbarInfoShowTypes struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListAllTaskViewResponseBodyResultToolbarInfoShowTypes) String() string {
	return tea.Prettify(s)
}

func (s ListAllTaskViewResponseBodyResultToolbarInfoShowTypes) GoString() string {
	return s.String()
}

func (s *ListAllTaskViewResponseBodyResultToolbarInfoShowTypes) SetName(v string) *ListAllTaskViewResponseBodyResultToolbarInfoShowTypes {
	s.Name = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultToolbarInfoShowTypes) SetValue(v string) *ListAllTaskViewResponseBodyResultToolbarInfoShowTypes {
	s.Value = &v
	return s
}

type ListAllTaskViewResponseBodyResultViewSetting struct {
	// example:
	//
	// true
	ShowDoneTask *bool `json:"showDoneTask,omitempty" xml:"showDoneTask,omitempty"`
	// example:
	//
	// true
	ShowSubTask *bool `json:"showSubTask,omitempty" xml:"showSubTask,omitempty"`
}

func (s ListAllTaskViewResponseBodyResultViewSetting) String() string {
	return tea.Prettify(s)
}

func (s ListAllTaskViewResponseBodyResultViewSetting) GoString() string {
	return s.String()
}

func (s *ListAllTaskViewResponseBodyResultViewSetting) SetShowDoneTask(v bool) *ListAllTaskViewResponseBodyResultViewSetting {
	s.ShowDoneTask = &v
	return s
}

func (s *ListAllTaskViewResponseBodyResultViewSetting) SetShowSubTask(v bool) *ListAllTaskViewResponseBodyResultViewSetting {
	s.ShowSubTask = &v
	return s
}

type ListAllTaskViewResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllTaskViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllTaskViewResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAllTaskViewResponse) GoString() string {
	return s.String()
}

func (s *ListAllTaskViewResponse) SetHeaders(v map[string]*string) *ListAllTaskViewResponse {
	s.Headers = v
	return s
}

func (s *ListAllTaskViewResponse) SetStatusCode(v int32) *ListAllTaskViewResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllTaskViewResponse) SetBody(v *ListAllTaskViewResponseBody) *ListAllTaskViewResponse {
	s.Body = v
	return s
}

type ListMyShortcutViewsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListMyShortcutViewsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListMyShortcutViewsHeaders) GoString() string {
	return s.String()
}

func (s *ListMyShortcutViewsHeaders) SetCommonHeaders(v map[string]*string) *ListMyShortcutViewsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListMyShortcutViewsHeaders) SetXAcsDingtalkAccessToken(v string) *ListMyShortcutViewsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListMyShortcutViewsRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListMyShortcutViewsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMyShortcutViewsRequest) GoString() string {
	return s.String()
}

func (s *ListMyShortcutViewsRequest) SetMaxResults(v int32) *ListMyShortcutViewsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMyShortcutViewsRequest) SetNextToken(v string) *ListMyShortcutViewsRequest {
	s.NextToken = &v
	return s
}

type ListMyShortcutViewsResponseBody struct {
	// example:
	//
	// f279e812-e431-428d-846d-cxxxxxx
	NextToken *string                                  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Result    []*ListMyShortcutViewsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListMyShortcutViewsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMyShortcutViewsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMyShortcutViewsResponseBody) SetNextToken(v string) *ListMyShortcutViewsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMyShortcutViewsResponseBody) SetResult(v []*ListMyShortcutViewsResponseBodyResult) *ListMyShortcutViewsResponseBody {
	s.Result = v
	return s
}

type ListMyShortcutViewsResponseBodyResult struct {
	BoundToObjectId *string `json:"boundToObjectId,omitempty" xml:"boundToObjectId,omitempty"`
	// example:
	//
	// user
	BoundToObjectType *string `json:"boundToObjectType,omitempty" xml:"boundToObjectType,omitempty"`
	// example:
	//
	// 2022-07-04T03:29:34.770Z
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// example:
	//
	// 5f687406f05b283425ea8f6f
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// xxxx
	Description *string                                         `json:"description,omitempty" xml:"description,omitempty"`
	Filter      *ListMyShortcutViewsResponseBodyResultFilter    `json:"filter,omitempty" xml:"filter,omitempty" type:"Struct"`
	GroupType   *ListMyShortcutViewsResponseBodyResultGroupType `json:"groupType,omitempty" xml:"groupType,omitempty" type:"Struct"`
	Id          *string                                         `json:"id,omitempty" xml:"id,omitempty"`
	IsDeleted   *bool                                           `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	// example:
	//
	// x项目
	Name      *string                                         `json:"name,omitempty" xml:"name,omitempty"`
	OrderType *ListMyShortcutViewsResponseBodyResultOrderType `json:"orderType,omitempty" xml:"orderType,omitempty" type:"Struct"`
	// example:
	//
	// 6139cd1aba5b128516ec8c86
	OrganizationId *string                                           `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	ShowType       *ListMyShortcutViewsResponseBodyResultShowType    `json:"showType,omitempty" xml:"showType,omitempty" type:"Struct"`
	ToolbarInfo    *ListMyShortcutViewsResponseBodyResultToolbarInfo `json:"toolbarInfo,omitempty" xml:"toolbarInfo,omitempty" type:"Struct"`
	Tql            *string                                           `json:"tql,omitempty" xml:"tql,omitempty"`
	// example:
	//
	// 2022-07-04T03:29:34.770Z
	Updated     *string                                           `json:"updated,omitempty" xml:"updated,omitempty"`
	ViewSetting *ListMyShortcutViewsResponseBodyResultViewSetting `json:"viewSetting,omitempty" xml:"viewSetting,omitempty" type:"Struct"`
}

func (s ListMyShortcutViewsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListMyShortcutViewsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListMyShortcutViewsResponseBodyResult) SetBoundToObjectId(v string) *ListMyShortcutViewsResponseBodyResult {
	s.BoundToObjectId = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResult) SetBoundToObjectType(v string) *ListMyShortcutViewsResponseBodyResult {
	s.BoundToObjectType = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResult) SetCreated(v string) *ListMyShortcutViewsResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResult) SetCreatorId(v string) *ListMyShortcutViewsResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResult) SetDescription(v string) *ListMyShortcutViewsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResult) SetFilter(v *ListMyShortcutViewsResponseBodyResultFilter) *ListMyShortcutViewsResponseBodyResult {
	s.Filter = v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResult) SetGroupType(v *ListMyShortcutViewsResponseBodyResultGroupType) *ListMyShortcutViewsResponseBodyResult {
	s.GroupType = v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResult) SetId(v string) *ListMyShortcutViewsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResult) SetIsDeleted(v bool) *ListMyShortcutViewsResponseBodyResult {
	s.IsDeleted = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResult) SetName(v string) *ListMyShortcutViewsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResult) SetOrderType(v *ListMyShortcutViewsResponseBodyResultOrderType) *ListMyShortcutViewsResponseBodyResult {
	s.OrderType = v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResult) SetOrganizationId(v string) *ListMyShortcutViewsResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResult) SetShowType(v *ListMyShortcutViewsResponseBodyResultShowType) *ListMyShortcutViewsResponseBodyResult {
	s.ShowType = v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResult) SetToolbarInfo(v *ListMyShortcutViewsResponseBodyResultToolbarInfo) *ListMyShortcutViewsResponseBodyResult {
	s.ToolbarInfo = v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResult) SetTql(v string) *ListMyShortcutViewsResponseBodyResult {
	s.Tql = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResult) SetUpdated(v string) *ListMyShortcutViewsResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResult) SetViewSetting(v *ListMyShortcutViewsResponseBodyResultViewSetting) *ListMyShortcutViewsResponseBodyResult {
	s.ViewSetting = v
	return s
}

type ListMyShortcutViewsResponseBodyResultFilter struct {
	Conditions          []*ListMyShortcutViewsResponseBodyResultFilterConditions          `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	InvisibleConditions []*ListMyShortcutViewsResponseBodyResultFilterInvisibleConditions `json:"invisibleConditions,omitempty" xml:"invisibleConditions,omitempty" type:"Repeated"`
	Logic               *string                                                           `json:"logic,omitempty" xml:"logic,omitempty"`
	LogicDisabled       *bool                                                             `json:"logicDisabled,omitempty" xml:"logicDisabled,omitempty"`
	Q                   *string                                                           `json:"q,omitempty" xml:"q,omitempty"`
}

func (s ListMyShortcutViewsResponseBodyResultFilter) String() string {
	return tea.Prettify(s)
}

func (s ListMyShortcutViewsResponseBodyResultFilter) GoString() string {
	return s.String()
}

func (s *ListMyShortcutViewsResponseBodyResultFilter) SetConditions(v []*ListMyShortcutViewsResponseBodyResultFilterConditions) *ListMyShortcutViewsResponseBodyResultFilter {
	s.Conditions = v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultFilter) SetInvisibleConditions(v []*ListMyShortcutViewsResponseBodyResultFilterInvisibleConditions) *ListMyShortcutViewsResponseBodyResultFilter {
	s.InvisibleConditions = v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultFilter) SetLogic(v string) *ListMyShortcutViewsResponseBodyResultFilter {
	s.Logic = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultFilter) SetLogicDisabled(v bool) *ListMyShortcutViewsResponseBodyResultFilter {
	s.LogicDisabled = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultFilter) SetQ(v string) *ListMyShortcutViewsResponseBodyResultFilter {
	s.Q = &v
	return s
}

type ListMyShortcutViewsResponseBodyResultFilterConditions struct {
	Fixed  *bool                                                          `json:"fixed,omitempty" xml:"fixed,omitempty"`
	Key    *string                                                        `json:"key,omitempty" xml:"key,omitempty"`
	Op     *ListMyShortcutViewsResponseBodyResultFilterConditionsOp       `json:"op,omitempty" xml:"op,omitempty" type:"Struct"`
	Values []*ListMyShortcutViewsResponseBodyResultFilterConditionsValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s ListMyShortcutViewsResponseBodyResultFilterConditions) String() string {
	return tea.Prettify(s)
}

func (s ListMyShortcutViewsResponseBodyResultFilterConditions) GoString() string {
	return s.String()
}

func (s *ListMyShortcutViewsResponseBodyResultFilterConditions) SetFixed(v bool) *ListMyShortcutViewsResponseBodyResultFilterConditions {
	s.Fixed = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultFilterConditions) SetKey(v string) *ListMyShortcutViewsResponseBodyResultFilterConditions {
	s.Key = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultFilterConditions) SetOp(v *ListMyShortcutViewsResponseBodyResultFilterConditionsOp) *ListMyShortcutViewsResponseBodyResultFilterConditions {
	s.Op = v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultFilterConditions) SetValues(v []*ListMyShortcutViewsResponseBodyResultFilterConditionsValues) *ListMyShortcutViewsResponseBodyResultFilterConditions {
	s.Values = v
	return s
}

type ListMyShortcutViewsResponseBodyResultFilterConditionsOp struct {
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListMyShortcutViewsResponseBodyResultFilterConditionsOp) String() string {
	return tea.Prettify(s)
}

func (s ListMyShortcutViewsResponseBodyResultFilterConditionsOp) GoString() string {
	return s.String()
}

func (s *ListMyShortcutViewsResponseBodyResultFilterConditionsOp) SetValue(v string) *ListMyShortcutViewsResponseBodyResultFilterConditionsOp {
	s.Value = &v
	return s
}

type ListMyShortcutViewsResponseBodyResultFilterConditionsValues struct {
	Deep  *string `json:"deep,omitempty" xml:"deep,omitempty"`
	Id    *string `json:"id,omitempty" xml:"id,omitempty"`
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListMyShortcutViewsResponseBodyResultFilterConditionsValues) String() string {
	return tea.Prettify(s)
}

func (s ListMyShortcutViewsResponseBodyResultFilterConditionsValues) GoString() string {
	return s.String()
}

func (s *ListMyShortcutViewsResponseBodyResultFilterConditionsValues) SetDeep(v string) *ListMyShortcutViewsResponseBodyResultFilterConditionsValues {
	s.Deep = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultFilterConditionsValues) SetId(v string) *ListMyShortcutViewsResponseBodyResultFilterConditionsValues {
	s.Id = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultFilterConditionsValues) SetLabel(v string) *ListMyShortcutViewsResponseBodyResultFilterConditionsValues {
	s.Label = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultFilterConditionsValues) SetName(v string) *ListMyShortcutViewsResponseBodyResultFilterConditionsValues {
	s.Name = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultFilterConditionsValues) SetValue(v string) *ListMyShortcutViewsResponseBodyResultFilterConditionsValues {
	s.Value = &v
	return s
}

type ListMyShortcutViewsResponseBodyResultFilterInvisibleConditions struct {
	Fixed  *bool                                                                   `json:"fixed,omitempty" xml:"fixed,omitempty"`
	Key    *string                                                                 `json:"key,omitempty" xml:"key,omitempty"`
	Op     *ListMyShortcutViewsResponseBodyResultFilterInvisibleConditionsOp       `json:"op,omitempty" xml:"op,omitempty" type:"Struct"`
	Values []*ListMyShortcutViewsResponseBodyResultFilterInvisibleConditionsValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s ListMyShortcutViewsResponseBodyResultFilterInvisibleConditions) String() string {
	return tea.Prettify(s)
}

func (s ListMyShortcutViewsResponseBodyResultFilterInvisibleConditions) GoString() string {
	return s.String()
}

func (s *ListMyShortcutViewsResponseBodyResultFilterInvisibleConditions) SetFixed(v bool) *ListMyShortcutViewsResponseBodyResultFilterInvisibleConditions {
	s.Fixed = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultFilterInvisibleConditions) SetKey(v string) *ListMyShortcutViewsResponseBodyResultFilterInvisibleConditions {
	s.Key = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultFilterInvisibleConditions) SetOp(v *ListMyShortcutViewsResponseBodyResultFilterInvisibleConditionsOp) *ListMyShortcutViewsResponseBodyResultFilterInvisibleConditions {
	s.Op = v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultFilterInvisibleConditions) SetValues(v []*ListMyShortcutViewsResponseBodyResultFilterInvisibleConditionsValues) *ListMyShortcutViewsResponseBodyResultFilterInvisibleConditions {
	s.Values = v
	return s
}

type ListMyShortcutViewsResponseBodyResultFilterInvisibleConditionsOp struct {
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListMyShortcutViewsResponseBodyResultFilterInvisibleConditionsOp) String() string {
	return tea.Prettify(s)
}

func (s ListMyShortcutViewsResponseBodyResultFilterInvisibleConditionsOp) GoString() string {
	return s.String()
}

func (s *ListMyShortcutViewsResponseBodyResultFilterInvisibleConditionsOp) SetValue(v string) *ListMyShortcutViewsResponseBodyResultFilterInvisibleConditionsOp {
	s.Value = &v
	return s
}

type ListMyShortcutViewsResponseBodyResultFilterInvisibleConditionsValues struct {
	Deep  *string `json:"deep,omitempty" xml:"deep,omitempty"`
	Id    *string `json:"id,omitempty" xml:"id,omitempty"`
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListMyShortcutViewsResponseBodyResultFilterInvisibleConditionsValues) String() string {
	return tea.Prettify(s)
}

func (s ListMyShortcutViewsResponseBodyResultFilterInvisibleConditionsValues) GoString() string {
	return s.String()
}

func (s *ListMyShortcutViewsResponseBodyResultFilterInvisibleConditionsValues) SetDeep(v string) *ListMyShortcutViewsResponseBodyResultFilterInvisibleConditionsValues {
	s.Deep = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultFilterInvisibleConditionsValues) SetId(v string) *ListMyShortcutViewsResponseBodyResultFilterInvisibleConditionsValues {
	s.Id = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultFilterInvisibleConditionsValues) SetLabel(v string) *ListMyShortcutViewsResponseBodyResultFilterInvisibleConditionsValues {
	s.Label = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultFilterInvisibleConditionsValues) SetName(v string) *ListMyShortcutViewsResponseBodyResultFilterInvisibleConditionsValues {
	s.Name = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultFilterInvisibleConditionsValues) SetValue(v string) *ListMyShortcutViewsResponseBodyResultFilterInvisibleConditionsValues {
	s.Value = &v
	return s
}

type ListMyShortcutViewsResponseBodyResultGroupType struct {
	CanCreateGroup *bool   `json:"canCreateGroup,omitempty" xml:"canCreateGroup,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	Value          *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListMyShortcutViewsResponseBodyResultGroupType) String() string {
	return tea.Prettify(s)
}

func (s ListMyShortcutViewsResponseBodyResultGroupType) GoString() string {
	return s.String()
}

func (s *ListMyShortcutViewsResponseBodyResultGroupType) SetCanCreateGroup(v bool) *ListMyShortcutViewsResponseBodyResultGroupType {
	s.CanCreateGroup = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultGroupType) SetName(v string) *ListMyShortcutViewsResponseBodyResultGroupType {
	s.Name = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultGroupType) SetValue(v string) *ListMyShortcutViewsResponseBodyResultGroupType {
	s.Value = &v
	return s
}

type ListMyShortcutViewsResponseBodyResultOrderType struct {
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Value     *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListMyShortcutViewsResponseBodyResultOrderType) String() string {
	return tea.Prettify(s)
}

func (s ListMyShortcutViewsResponseBodyResultOrderType) GoString() string {
	return s.String()
}

func (s *ListMyShortcutViewsResponseBodyResultOrderType) SetDirection(v string) *ListMyShortcutViewsResponseBodyResultOrderType {
	s.Direction = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultOrderType) SetName(v string) *ListMyShortcutViewsResponseBodyResultOrderType {
	s.Name = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultOrderType) SetValue(v string) *ListMyShortcutViewsResponseBodyResultOrderType {
	s.Value = &v
	return s
}

type ListMyShortcutViewsResponseBodyResultShowType struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListMyShortcutViewsResponseBodyResultShowType) String() string {
	return tea.Prettify(s)
}

func (s ListMyShortcutViewsResponseBodyResultShowType) GoString() string {
	return s.String()
}

func (s *ListMyShortcutViewsResponseBodyResultShowType) SetName(v string) *ListMyShortcutViewsResponseBodyResultShowType {
	s.Name = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultShowType) SetValue(v string) *ListMyShortcutViewsResponseBodyResultShowType {
	s.Value = &v
	return s
}

type ListMyShortcutViewsResponseBodyResultToolbarInfo struct {
	GroupTypes []*ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypes `json:"groupTypes,omitempty" xml:"groupTypes,omitempty" type:"Repeated"`
	OrderTypes []*ListMyShortcutViewsResponseBodyResultToolbarInfoOrderTypes `json:"orderTypes,omitempty" xml:"orderTypes,omitempty" type:"Repeated"`
	ShowTypes  []*ListMyShortcutViewsResponseBodyResultToolbarInfoShowTypes  `json:"showTypes,omitempty" xml:"showTypes,omitempty" type:"Repeated"`
}

func (s ListMyShortcutViewsResponseBodyResultToolbarInfo) String() string {
	return tea.Prettify(s)
}

func (s ListMyShortcutViewsResponseBodyResultToolbarInfo) GoString() string {
	return s.String()
}

func (s *ListMyShortcutViewsResponseBodyResultToolbarInfo) SetGroupTypes(v []*ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypes) *ListMyShortcutViewsResponseBodyResultToolbarInfo {
	s.GroupTypes = v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultToolbarInfo) SetOrderTypes(v []*ListMyShortcutViewsResponseBodyResultToolbarInfoOrderTypes) *ListMyShortcutViewsResponseBodyResultToolbarInfo {
	s.OrderTypes = v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultToolbarInfo) SetShowTypes(v []*ListMyShortcutViewsResponseBodyResultToolbarInfoShowTypes) *ListMyShortcutViewsResponseBodyResultToolbarInfo {
	s.ShowTypes = v
	return s
}

type ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypes struct {
	CanCreateGroup *bool                                                              `json:"canCreateGroup,omitempty" xml:"canCreateGroup,omitempty"`
	Name           *string                                                            `json:"name,omitempty" xml:"name,omitempty"`
	ServiceName    *string                                                            `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	Setting        *ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypesSetting `json:"setting,omitempty" xml:"setting,omitempty" type:"Struct"`
	Value          *string                                                            `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypes) String() string {
	return tea.Prettify(s)
}

func (s ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypes) GoString() string {
	return s.String()
}

func (s *ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypes) SetCanCreateGroup(v bool) *ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypes {
	s.CanCreateGroup = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypes) SetName(v string) *ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypes {
	s.Name = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypes) SetServiceName(v string) *ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypes {
	s.ServiceName = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypes) SetSetting(v *ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypesSetting) *ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypes {
	s.Setting = v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypes) SetValue(v string) *ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypes {
	s.Value = &v
	return s
}

type ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypesSetting struct {
	DateType  *string `json:"dateType,omitempty" xml:"dateType,omitempty"`
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
}

func (s ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypesSetting) String() string {
	return tea.Prettify(s)
}

func (s ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypesSetting) GoString() string {
	return s.String()
}

func (s *ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypesSetting) SetDateType(v string) *ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypesSetting {
	s.DateType = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypesSetting) SetFieldName(v string) *ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypesSetting {
	s.FieldName = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypesSetting) SetFieldType(v string) *ListMyShortcutViewsResponseBodyResultToolbarInfoGroupTypesSetting {
	s.FieldType = &v
	return s
}

type ListMyShortcutViewsResponseBodyResultToolbarInfoOrderTypes struct {
	Direction        *string `json:"direction,omitempty" xml:"direction,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	SupportDirection *string `json:"supportDirection,omitempty" xml:"supportDirection,omitempty"`
	Value            *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListMyShortcutViewsResponseBodyResultToolbarInfoOrderTypes) String() string {
	return tea.Prettify(s)
}

func (s ListMyShortcutViewsResponseBodyResultToolbarInfoOrderTypes) GoString() string {
	return s.String()
}

func (s *ListMyShortcutViewsResponseBodyResultToolbarInfoOrderTypes) SetDirection(v string) *ListMyShortcutViewsResponseBodyResultToolbarInfoOrderTypes {
	s.Direction = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultToolbarInfoOrderTypes) SetName(v string) *ListMyShortcutViewsResponseBodyResultToolbarInfoOrderTypes {
	s.Name = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultToolbarInfoOrderTypes) SetSupportDirection(v string) *ListMyShortcutViewsResponseBodyResultToolbarInfoOrderTypes {
	s.SupportDirection = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultToolbarInfoOrderTypes) SetValue(v string) *ListMyShortcutViewsResponseBodyResultToolbarInfoOrderTypes {
	s.Value = &v
	return s
}

type ListMyShortcutViewsResponseBodyResultToolbarInfoShowTypes struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListMyShortcutViewsResponseBodyResultToolbarInfoShowTypes) String() string {
	return tea.Prettify(s)
}

func (s ListMyShortcutViewsResponseBodyResultToolbarInfoShowTypes) GoString() string {
	return s.String()
}

func (s *ListMyShortcutViewsResponseBodyResultToolbarInfoShowTypes) SetName(v string) *ListMyShortcutViewsResponseBodyResultToolbarInfoShowTypes {
	s.Name = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultToolbarInfoShowTypes) SetValue(v string) *ListMyShortcutViewsResponseBodyResultToolbarInfoShowTypes {
	s.Value = &v
	return s
}

type ListMyShortcutViewsResponseBodyResultViewSetting struct {
	// example:
	//
	// true
	ShowDoneTask *bool `json:"showDoneTask,omitempty" xml:"showDoneTask,omitempty"`
	// example:
	//
	// true
	ShowSubTask *bool `json:"showSubTask,omitempty" xml:"showSubTask,omitempty"`
}

func (s ListMyShortcutViewsResponseBodyResultViewSetting) String() string {
	return tea.Prettify(s)
}

func (s ListMyShortcutViewsResponseBodyResultViewSetting) GoString() string {
	return s.String()
}

func (s *ListMyShortcutViewsResponseBodyResultViewSetting) SetShowDoneTask(v bool) *ListMyShortcutViewsResponseBodyResultViewSetting {
	s.ShowDoneTask = &v
	return s
}

func (s *ListMyShortcutViewsResponseBodyResultViewSetting) SetShowSubTask(v bool) *ListMyShortcutViewsResponseBodyResultViewSetting {
	s.ShowSubTask = &v
	return s
}

type ListMyShortcutViewsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMyShortcutViewsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMyShortcutViewsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMyShortcutViewsResponse) GoString() string {
	return s.String()
}

func (s *ListMyShortcutViewsResponse) SetHeaders(v map[string]*string) *ListMyShortcutViewsResponse {
	s.Headers = v
	return s
}

func (s *ListMyShortcutViewsResponse) SetStatusCode(v int32) *ListMyShortcutViewsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMyShortcutViewsResponse) SetBody(v *ListMyShortcutViewsResponseBody) *ListMyShortcutViewsResponse {
	s.Body = v
	return s
}

type QueryAllTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllTaskHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllTaskHeaders) SetCommonHeaders(v map[string]*string) *QueryAllTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllTaskHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllTaskRequest struct {
	// This parameter is required.
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s QueryAllTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllTaskRequest) GoString() string {
	return s.String()
}

func (s *QueryAllTaskRequest) SetTaskId(v string) *QueryAllTaskRequest {
	s.TaskId = &v
	return s
}

type QueryAllTaskResponseBody struct {
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*QueryAllTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryAllTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllTaskResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllTaskResponseBody) SetRequestId(v string) *QueryAllTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAllTaskResponseBody) SetResult(v []*QueryAllTaskResponseBodyResult) *QueryAllTaskResponseBody {
	s.Result = v
	return s
}

type QueryAllTaskResponseBodyResult struct {
	AccomplishTime *string   `json:"accomplishTime,omitempty" xml:"accomplishTime,omitempty"`
	AncestorIds    []*string `json:"ancestorIds,omitempty" xml:"ancestorIds,omitempty" type:"Repeated"`
	Content        *string   `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 2022-07-04T03:29:34.770Z
	Created        *string                                       `json:"created,omitempty" xml:"created,omitempty"`
	CreatorId      *string                                       `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Customfields   []*QueryAllTaskResponseBodyResultCustomfields `json:"customfields,omitempty" xml:"customfields,omitempty" type:"Repeated"`
	DueDate        *string                                       `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	ExecutorId     *string                                       `json:"executorId,omitempty" xml:"executorId,omitempty"`
	Id             *string                                       `json:"id,omitempty" xml:"id,omitempty"`
	InvolveMembers []*string                                     `json:"involveMembers,omitempty" xml:"involveMembers,omitempty" type:"Repeated"`
	IsArchived     *bool                                         `json:"isArchived,omitempty" xml:"isArchived,omitempty"`
	IsDone         *bool                                         `json:"isDone,omitempty" xml:"isDone,omitempty"`
	Note           *string                                       `json:"note,omitempty" xml:"note,omitempty"`
	ParentTaskId   *string                                       `json:"parentTaskId,omitempty" xml:"parentTaskId,omitempty"`
	Priority       *int32                                        `json:"priority,omitempty" xml:"priority,omitempty"`
	ProjectId      *string                                       `json:"projectId,omitempty" xml:"projectId,omitempty"`
	SfcId          *string                                       `json:"sfcId,omitempty" xml:"sfcId,omitempty"`
	StageId        *string                                       `json:"stageId,omitempty" xml:"stageId,omitempty"`
	// example:
	//
	// 2022-07-04T03:29:34.770Z
	StartDate *string   `json:"startDate,omitempty" xml:"startDate,omitempty"`
	TagIds    []*string `json:"tagIds,omitempty" xml:"tagIds,omitempty" type:"Repeated"`
	// Deprecated
	TaskId     *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TasklistId *string `json:"tasklistId,omitempty" xml:"tasklistId,omitempty"`
	TfsId      *string `json:"tfsId,omitempty" xml:"tfsId,omitempty"`
	UniqueId   *string `json:"uniqueId,omitempty" xml:"uniqueId,omitempty"`
	// example:
	//
	// 2022-07-04T03:29:34.770Z
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
	Visible *string `json:"visible,omitempty" xml:"visible,omitempty"`
}

func (s QueryAllTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryAllTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryAllTaskResponseBodyResult) SetAccomplishTime(v string) *QueryAllTaskResponseBodyResult {
	s.AccomplishTime = &v
	return s
}

func (s *QueryAllTaskResponseBodyResult) SetAncestorIds(v []*string) *QueryAllTaskResponseBodyResult {
	s.AncestorIds = v
	return s
}

func (s *QueryAllTaskResponseBodyResult) SetContent(v string) *QueryAllTaskResponseBodyResult {
	s.Content = &v
	return s
}

func (s *QueryAllTaskResponseBodyResult) SetCreated(v string) *QueryAllTaskResponseBodyResult {
	s.Created = &v
	return s
}

func (s *QueryAllTaskResponseBodyResult) SetCreatorId(v string) *QueryAllTaskResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *QueryAllTaskResponseBodyResult) SetCustomfields(v []*QueryAllTaskResponseBodyResultCustomfields) *QueryAllTaskResponseBodyResult {
	s.Customfields = v
	return s
}

func (s *QueryAllTaskResponseBodyResult) SetDueDate(v string) *QueryAllTaskResponseBodyResult {
	s.DueDate = &v
	return s
}

func (s *QueryAllTaskResponseBodyResult) SetExecutorId(v string) *QueryAllTaskResponseBodyResult {
	s.ExecutorId = &v
	return s
}

func (s *QueryAllTaskResponseBodyResult) SetId(v string) *QueryAllTaskResponseBodyResult {
	s.Id = &v
	return s
}

func (s *QueryAllTaskResponseBodyResult) SetInvolveMembers(v []*string) *QueryAllTaskResponseBodyResult {
	s.InvolveMembers = v
	return s
}

func (s *QueryAllTaskResponseBodyResult) SetIsArchived(v bool) *QueryAllTaskResponseBodyResult {
	s.IsArchived = &v
	return s
}

func (s *QueryAllTaskResponseBodyResult) SetIsDone(v bool) *QueryAllTaskResponseBodyResult {
	s.IsDone = &v
	return s
}

func (s *QueryAllTaskResponseBodyResult) SetNote(v string) *QueryAllTaskResponseBodyResult {
	s.Note = &v
	return s
}

func (s *QueryAllTaskResponseBodyResult) SetParentTaskId(v string) *QueryAllTaskResponseBodyResult {
	s.ParentTaskId = &v
	return s
}

func (s *QueryAllTaskResponseBodyResult) SetPriority(v int32) *QueryAllTaskResponseBodyResult {
	s.Priority = &v
	return s
}

func (s *QueryAllTaskResponseBodyResult) SetProjectId(v string) *QueryAllTaskResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *QueryAllTaskResponseBodyResult) SetSfcId(v string) *QueryAllTaskResponseBodyResult {
	s.SfcId = &v
	return s
}

func (s *QueryAllTaskResponseBodyResult) SetStageId(v string) *QueryAllTaskResponseBodyResult {
	s.StageId = &v
	return s
}

func (s *QueryAllTaskResponseBodyResult) SetStartDate(v string) *QueryAllTaskResponseBodyResult {
	s.StartDate = &v
	return s
}

func (s *QueryAllTaskResponseBodyResult) SetTagIds(v []*string) *QueryAllTaskResponseBodyResult {
	s.TagIds = v
	return s
}

func (s *QueryAllTaskResponseBodyResult) SetTaskId(v string) *QueryAllTaskResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *QueryAllTaskResponseBodyResult) SetTasklistId(v string) *QueryAllTaskResponseBodyResult {
	s.TasklistId = &v
	return s
}

func (s *QueryAllTaskResponseBodyResult) SetTfsId(v string) *QueryAllTaskResponseBodyResult {
	s.TfsId = &v
	return s
}

func (s *QueryAllTaskResponseBodyResult) SetUniqueId(v string) *QueryAllTaskResponseBodyResult {
	s.UniqueId = &v
	return s
}

func (s *QueryAllTaskResponseBodyResult) SetUpdated(v string) *QueryAllTaskResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *QueryAllTaskResponseBodyResult) SetVisible(v string) *QueryAllTaskResponseBodyResult {
	s.Visible = &v
	return s
}

type QueryAllTaskResponseBodyResultCustomfields struct {
	CfId  *string                                            `json:"cfId,omitempty" xml:"cfId,omitempty"`
	Type  *string                                            `json:"type,omitempty" xml:"type,omitempty"`
	Value []*QueryAllTaskResponseBodyResultCustomfieldsValue `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s QueryAllTaskResponseBodyResultCustomfields) String() string {
	return tea.Prettify(s)
}

func (s QueryAllTaskResponseBodyResultCustomfields) GoString() string {
	return s.String()
}

func (s *QueryAllTaskResponseBodyResultCustomfields) SetCfId(v string) *QueryAllTaskResponseBodyResultCustomfields {
	s.CfId = &v
	return s
}

func (s *QueryAllTaskResponseBodyResultCustomfields) SetType(v string) *QueryAllTaskResponseBodyResultCustomfields {
	s.Type = &v
	return s
}

func (s *QueryAllTaskResponseBodyResultCustomfields) SetValue(v []*QueryAllTaskResponseBodyResultCustomfieldsValue) *QueryAllTaskResponseBodyResultCustomfields {
	s.Value = v
	return s
}

type QueryAllTaskResponseBodyResultCustomfieldsValue struct {
	Id         *string `json:"id,omitempty" xml:"id,omitempty"`
	MetaString *string `json:"metaString,omitempty" xml:"metaString,omitempty"`
	Title      *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryAllTaskResponseBodyResultCustomfieldsValue) String() string {
	return tea.Prettify(s)
}

func (s QueryAllTaskResponseBodyResultCustomfieldsValue) GoString() string {
	return s.String()
}

func (s *QueryAllTaskResponseBodyResultCustomfieldsValue) SetId(v string) *QueryAllTaskResponseBodyResultCustomfieldsValue {
	s.Id = &v
	return s
}

func (s *QueryAllTaskResponseBodyResultCustomfieldsValue) SetMetaString(v string) *QueryAllTaskResponseBodyResultCustomfieldsValue {
	s.MetaString = &v
	return s
}

func (s *QueryAllTaskResponseBodyResultCustomfieldsValue) SetTitle(v string) *QueryAllTaskResponseBodyResultCustomfieldsValue {
	s.Title = &v
	return s
}

type QueryAllTaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAllTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAllTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllTaskResponse) GoString() string {
	return s.String()
}

func (s *QueryAllTaskResponse) SetHeaders(v map[string]*string) *QueryAllTaskResponse {
	s.Headers = v
	return s
}

func (s *QueryAllTaskResponse) SetStatusCode(v int32) *QueryAllTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAllTaskResponse) SetBody(v *QueryAllTaskResponseBody) *QueryAllTaskResponse {
	s.Body = v
	return s
}

type QueryTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskHeaders) GoString() string {
	return s.String()
}

func (s *QueryTaskHeaders) SetCommonHeaders(v map[string]*string) *QueryTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryTaskHeaders) SetXAcsDingtalkAccessToken(v string) *QueryTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryTaskRequest struct {
	MaxResults *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Tql        *string `json:"tql,omitempty" xml:"tql,omitempty"`
}

func (s QueryTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskRequest) SetMaxResults(v int64) *QueryTaskRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryTaskRequest) SetNextToken(v string) *QueryTaskRequest {
	s.NextToken = &v
	return s
}

func (s *QueryTaskRequest) SetTql(v string) *QueryTaskRequest {
	s.Tql = &v
	return s
}

type QueryTaskResponseBody struct {
	// example:
	//
	// f279e812-e431-428d-846d-cxxxxxx
	NextToken *string                        `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId *string                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*QueryTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTaskResponseBody) SetNextToken(v string) *QueryTaskResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryTaskResponseBody) SetRequestId(v string) *QueryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTaskResponseBody) SetResult(v []*QueryTaskResponseBodyResult) *QueryTaskResponseBody {
	s.Result = v
	return s
}

type QueryTaskResponseBodyResult struct {
	AccomplishTime *string `json:"accomplishTime,omitempty" xml:"accomplishTime,omitempty"`
	Content        *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 2022-07-04T03:29:34.770Z
	Created          *string                                      `json:"created,omitempty" xml:"created,omitempty"`
	CreatorId        *string                                      `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Customfields     []*QueryTaskResponseBodyResultCustomfields   `json:"customfields,omitempty" xml:"customfields,omitempty" type:"Repeated"`
	DueDate          *string                                      `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	ExecutorId       *string                                      `json:"executorId,omitempty" xml:"executorId,omitempty"`
	ExecutorUserInfo *QueryTaskResponseBodyResultExecutorUserInfo `json:"executorUserInfo,omitempty" xml:"executorUserInfo,omitempty" type:"Struct"`
	Id               *string                                      `json:"id,omitempty" xml:"id,omitempty"`
	InvolveMembers   []*string                                    `json:"involveMembers,omitempty" xml:"involveMembers,omitempty" type:"Repeated"`
	IsDone           *bool                                        `json:"isDone,omitempty" xml:"isDone,omitempty"`
	Note             *string                                      `json:"note,omitempty" xml:"note,omitempty"`
	ProjectId        *string                                      `json:"projectId,omitempty" xml:"projectId,omitempty"`
	ProjectInfo      *QueryTaskResponseBodyResultProjectInfo      `json:"projectInfo,omitempty" xml:"projectInfo,omitempty" type:"Struct"`
	// Deprecated
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 2022-07-04T03:29:34.770Z
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s QueryTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryTaskResponseBodyResult) SetAccomplishTime(v string) *QueryTaskResponseBodyResult {
	s.AccomplishTime = &v
	return s
}

func (s *QueryTaskResponseBodyResult) SetContent(v string) *QueryTaskResponseBodyResult {
	s.Content = &v
	return s
}

func (s *QueryTaskResponseBodyResult) SetCreated(v string) *QueryTaskResponseBodyResult {
	s.Created = &v
	return s
}

func (s *QueryTaskResponseBodyResult) SetCreatorId(v string) *QueryTaskResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *QueryTaskResponseBodyResult) SetCustomfields(v []*QueryTaskResponseBodyResultCustomfields) *QueryTaskResponseBodyResult {
	s.Customfields = v
	return s
}

func (s *QueryTaskResponseBodyResult) SetDueDate(v string) *QueryTaskResponseBodyResult {
	s.DueDate = &v
	return s
}

func (s *QueryTaskResponseBodyResult) SetExecutorId(v string) *QueryTaskResponseBodyResult {
	s.ExecutorId = &v
	return s
}

func (s *QueryTaskResponseBodyResult) SetExecutorUserInfo(v *QueryTaskResponseBodyResultExecutorUserInfo) *QueryTaskResponseBodyResult {
	s.ExecutorUserInfo = v
	return s
}

func (s *QueryTaskResponseBodyResult) SetId(v string) *QueryTaskResponseBodyResult {
	s.Id = &v
	return s
}

func (s *QueryTaskResponseBodyResult) SetInvolveMembers(v []*string) *QueryTaskResponseBodyResult {
	s.InvolveMembers = v
	return s
}

func (s *QueryTaskResponseBodyResult) SetIsDone(v bool) *QueryTaskResponseBodyResult {
	s.IsDone = &v
	return s
}

func (s *QueryTaskResponseBodyResult) SetNote(v string) *QueryTaskResponseBodyResult {
	s.Note = &v
	return s
}

func (s *QueryTaskResponseBodyResult) SetProjectId(v string) *QueryTaskResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *QueryTaskResponseBodyResult) SetProjectInfo(v *QueryTaskResponseBodyResultProjectInfo) *QueryTaskResponseBodyResult {
	s.ProjectInfo = v
	return s
}

func (s *QueryTaskResponseBodyResult) SetTaskId(v string) *QueryTaskResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *QueryTaskResponseBodyResult) SetUpdated(v string) *QueryTaskResponseBodyResult {
	s.Updated = &v
	return s
}

type QueryTaskResponseBodyResultCustomfields struct {
	CfId  *string                                         `json:"cfId,omitempty" xml:"cfId,omitempty"`
	Type  *string                                         `json:"type,omitempty" xml:"type,omitempty"`
	Value []*QueryTaskResponseBodyResultCustomfieldsValue `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s QueryTaskResponseBodyResultCustomfields) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskResponseBodyResultCustomfields) GoString() string {
	return s.String()
}

func (s *QueryTaskResponseBodyResultCustomfields) SetCfId(v string) *QueryTaskResponseBodyResultCustomfields {
	s.CfId = &v
	return s
}

func (s *QueryTaskResponseBodyResultCustomfields) SetType(v string) *QueryTaskResponseBodyResultCustomfields {
	s.Type = &v
	return s
}

func (s *QueryTaskResponseBodyResultCustomfields) SetValue(v []*QueryTaskResponseBodyResultCustomfieldsValue) *QueryTaskResponseBodyResultCustomfields {
	s.Value = v
	return s
}

type QueryTaskResponseBodyResultCustomfieldsValue struct {
	Id         *string `json:"id,omitempty" xml:"id,omitempty"`
	MetaString *string `json:"metaString,omitempty" xml:"metaString,omitempty"`
	Title      *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryTaskResponseBodyResultCustomfieldsValue) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskResponseBodyResultCustomfieldsValue) GoString() string {
	return s.String()
}

func (s *QueryTaskResponseBodyResultCustomfieldsValue) SetId(v string) *QueryTaskResponseBodyResultCustomfieldsValue {
	s.Id = &v
	return s
}

func (s *QueryTaskResponseBodyResultCustomfieldsValue) SetMetaString(v string) *QueryTaskResponseBodyResultCustomfieldsValue {
	s.MetaString = &v
	return s
}

func (s *QueryTaskResponseBodyResultCustomfieldsValue) SetTitle(v string) *QueryTaskResponseBodyResultCustomfieldsValue {
	s.Title = &v
	return s
}

type QueryTaskResponseBodyResultExecutorUserInfo struct {
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	Id        *string `json:"id,omitempty" xml:"id,omitempty"`
	MemberId  *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryTaskResponseBodyResultExecutorUserInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskResponseBodyResultExecutorUserInfo) GoString() string {
	return s.String()
}

func (s *QueryTaskResponseBodyResultExecutorUserInfo) SetAvatarUrl(v string) *QueryTaskResponseBodyResultExecutorUserInfo {
	s.AvatarUrl = &v
	return s
}

func (s *QueryTaskResponseBodyResultExecutorUserInfo) SetId(v string) *QueryTaskResponseBodyResultExecutorUserInfo {
	s.Id = &v
	return s
}

func (s *QueryTaskResponseBodyResultExecutorUserInfo) SetMemberId(v string) *QueryTaskResponseBodyResultExecutorUserInfo {
	s.MemberId = &v
	return s
}

func (s *QueryTaskResponseBodyResultExecutorUserInfo) SetName(v string) *QueryTaskResponseBodyResultExecutorUserInfo {
	s.Name = &v
	return s
}

func (s *QueryTaskResponseBodyResultExecutorUserInfo) SetUserId(v string) *QueryTaskResponseBodyResultExecutorUserInfo {
	s.UserId = &v
	return s
}

type QueryTaskResponseBodyResultProjectInfo struct {
	Id             *string `json:"id,omitempty" xml:"id,omitempty"`
	Logo           *string `json:"logo,omitempty" xml:"logo,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s QueryTaskResponseBodyResultProjectInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskResponseBodyResultProjectInfo) GoString() string {
	return s.String()
}

func (s *QueryTaskResponseBodyResultProjectInfo) SetId(v string) *QueryTaskResponseBodyResultProjectInfo {
	s.Id = &v
	return s
}

func (s *QueryTaskResponseBodyResultProjectInfo) SetLogo(v string) *QueryTaskResponseBodyResultProjectInfo {
	s.Logo = &v
	return s
}

func (s *QueryTaskResponseBodyResultProjectInfo) SetName(v string) *QueryTaskResponseBodyResultProjectInfo {
	s.Name = &v
	return s
}

func (s *QueryTaskResponseBodyResultProjectInfo) SetOrganizationId(v string) *QueryTaskResponseBodyResultProjectInfo {
	s.OrganizationId = &v
	return s
}

type QueryTaskResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskResponse) SetHeaders(v map[string]*string) *QueryTaskResponse {
	s.Headers = v
	return s
}

func (s *QueryTaskResponse) SetStatusCode(v int32) *QueryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTaskResponse) SetBody(v *QueryTaskResponseBody) *QueryTaskResponse {
	s.Body = v
	return s
}

type QueryTasksV3Headers struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryTasksV3Headers) String() string {
	return tea.Prettify(s)
}

func (s QueryTasksV3Headers) GoString() string {
	return s.String()
}

func (s *QueryTasksV3Headers) SetCommonHeaders(v map[string]*string) *QueryTasksV3Headers {
	s.CommonHeaders = v
	return s
}

func (s *QueryTasksV3Headers) SetXAcsDingtalkAccessToken(v string) *QueryTasksV3Headers {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryTasksV3Request struct {
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s QueryTasksV3Request) String() string {
	return tea.Prettify(s)
}

func (s QueryTasksV3Request) GoString() string {
	return s.String()
}

func (s *QueryTasksV3Request) SetTaskId(v string) *QueryTasksV3Request {
	s.TaskId = &v
	return s
}

type QueryTasksV3ResponseBody struct {
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*QueryTasksV3ResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryTasksV3ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTasksV3ResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTasksV3ResponseBody) SetRequestId(v string) *QueryTasksV3ResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTasksV3ResponseBody) SetResult(v []*QueryTasksV3ResponseBodyResult) *QueryTasksV3ResponseBody {
	s.Result = v
	return s
}

type QueryTasksV3ResponseBodyResult struct {
	AccomplishTime *string `json:"accomplishTime,omitempty" xml:"accomplishTime,omitempty"`
	Content        *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 2022-07-04T03:29:34.770Z
	Created        *string                                       `json:"created,omitempty" xml:"created,omitempty"`
	CreatorId      *string                                       `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Customfields   []*QueryTasksV3ResponseBodyResultCustomfields `json:"customfields,omitempty" xml:"customfields,omitempty" type:"Repeated"`
	DueDate        *string                                       `json:"dueDate,omitempty" xml:"dueDate,omitempty"`
	ExecutorId     *string                                       `json:"executorId,omitempty" xml:"executorId,omitempty"`
	Id             *string                                       `json:"id,omitempty" xml:"id,omitempty"`
	InvolveMembers []*string                                     `json:"involveMembers,omitempty" xml:"involveMembers,omitempty" type:"Repeated"`
	IsDone         *bool                                         `json:"isDone,omitempty" xml:"isDone,omitempty"`
	Note           *string                                       `json:"note,omitempty" xml:"note,omitempty"`
	ProjectId      *string                                       `json:"projectId,omitempty" xml:"projectId,omitempty"`
	SourceId       *string                                       `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// Deprecated
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 2022-07-04T03:29:34.770Z
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s QueryTasksV3ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryTasksV3ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryTasksV3ResponseBodyResult) SetAccomplishTime(v string) *QueryTasksV3ResponseBodyResult {
	s.AccomplishTime = &v
	return s
}

func (s *QueryTasksV3ResponseBodyResult) SetContent(v string) *QueryTasksV3ResponseBodyResult {
	s.Content = &v
	return s
}

func (s *QueryTasksV3ResponseBodyResult) SetCreated(v string) *QueryTasksV3ResponseBodyResult {
	s.Created = &v
	return s
}

func (s *QueryTasksV3ResponseBodyResult) SetCreatorId(v string) *QueryTasksV3ResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *QueryTasksV3ResponseBodyResult) SetCustomfields(v []*QueryTasksV3ResponseBodyResultCustomfields) *QueryTasksV3ResponseBodyResult {
	s.Customfields = v
	return s
}

func (s *QueryTasksV3ResponseBodyResult) SetDueDate(v string) *QueryTasksV3ResponseBodyResult {
	s.DueDate = &v
	return s
}

func (s *QueryTasksV3ResponseBodyResult) SetExecutorId(v string) *QueryTasksV3ResponseBodyResult {
	s.ExecutorId = &v
	return s
}

func (s *QueryTasksV3ResponseBodyResult) SetId(v string) *QueryTasksV3ResponseBodyResult {
	s.Id = &v
	return s
}

func (s *QueryTasksV3ResponseBodyResult) SetInvolveMembers(v []*string) *QueryTasksV3ResponseBodyResult {
	s.InvolveMembers = v
	return s
}

func (s *QueryTasksV3ResponseBodyResult) SetIsDone(v bool) *QueryTasksV3ResponseBodyResult {
	s.IsDone = &v
	return s
}

func (s *QueryTasksV3ResponseBodyResult) SetNote(v string) *QueryTasksV3ResponseBodyResult {
	s.Note = &v
	return s
}

func (s *QueryTasksV3ResponseBodyResult) SetProjectId(v string) *QueryTasksV3ResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *QueryTasksV3ResponseBodyResult) SetSourceId(v string) *QueryTasksV3ResponseBodyResult {
	s.SourceId = &v
	return s
}

func (s *QueryTasksV3ResponseBodyResult) SetTaskId(v string) *QueryTasksV3ResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *QueryTasksV3ResponseBodyResult) SetUpdated(v string) *QueryTasksV3ResponseBodyResult {
	s.Updated = &v
	return s
}

type QueryTasksV3ResponseBodyResultCustomfields struct {
	CfId  *string                                            `json:"cfId,omitempty" xml:"cfId,omitempty"`
	Type  *string                                            `json:"type,omitempty" xml:"type,omitempty"`
	Value []*QueryTasksV3ResponseBodyResultCustomfieldsValue `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s QueryTasksV3ResponseBodyResultCustomfields) String() string {
	return tea.Prettify(s)
}

func (s QueryTasksV3ResponseBodyResultCustomfields) GoString() string {
	return s.String()
}

func (s *QueryTasksV3ResponseBodyResultCustomfields) SetCfId(v string) *QueryTasksV3ResponseBodyResultCustomfields {
	s.CfId = &v
	return s
}

func (s *QueryTasksV3ResponseBodyResultCustomfields) SetType(v string) *QueryTasksV3ResponseBodyResultCustomfields {
	s.Type = &v
	return s
}

func (s *QueryTasksV3ResponseBodyResultCustomfields) SetValue(v []*QueryTasksV3ResponseBodyResultCustomfieldsValue) *QueryTasksV3ResponseBodyResultCustomfields {
	s.Value = v
	return s
}

type QueryTasksV3ResponseBodyResultCustomfieldsValue struct {
	Id         *string `json:"id,omitempty" xml:"id,omitempty"`
	MetaString *string `json:"metaString,omitempty" xml:"metaString,omitempty"`
	Title      *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryTasksV3ResponseBodyResultCustomfieldsValue) String() string {
	return tea.Prettify(s)
}

func (s QueryTasksV3ResponseBodyResultCustomfieldsValue) GoString() string {
	return s.String()
}

func (s *QueryTasksV3ResponseBodyResultCustomfieldsValue) SetId(v string) *QueryTasksV3ResponseBodyResultCustomfieldsValue {
	s.Id = &v
	return s
}

func (s *QueryTasksV3ResponseBodyResultCustomfieldsValue) SetMetaString(v string) *QueryTasksV3ResponseBodyResultCustomfieldsValue {
	s.MetaString = &v
	return s
}

func (s *QueryTasksV3ResponseBodyResultCustomfieldsValue) SetTitle(v string) *QueryTasksV3ResponseBodyResultCustomfieldsValue {
	s.Title = &v
	return s
}

type QueryTasksV3Response struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTasksV3ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTasksV3Response) String() string {
	return tea.Prettify(s)
}

func (s QueryTasksV3Response) GoString() string {
	return s.String()
}

func (s *QueryTasksV3Response) SetHeaders(v map[string]*string) *QueryTasksV3Response {
	s.Headers = v
	return s
}

func (s *QueryTasksV3Response) SetStatusCode(v int32) *QueryTasksV3Response {
	s.StatusCode = &v
	return s
}

func (s *QueryTasksV3Response) SetBody(v *QueryTasksV3ResponseBody) *QueryTasksV3Response {
	s.Body = v
	return s
}

type SearchAllTasksByTqlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchAllTasksByTqlHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchAllTasksByTqlHeaders) GoString() string {
	return s.String()
}

func (s *SearchAllTasksByTqlHeaders) SetCommonHeaders(v map[string]*string) *SearchAllTasksByTqlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchAllTasksByTqlHeaders) SetXAcsDingtalkAccessToken(v string) *SearchAllTasksByTqlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchAllTasksByTqlRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Tql        *string `json:"tql,omitempty" xml:"tql,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0517xxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchAllTasksByTqlRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchAllTasksByTqlRequest) GoString() string {
	return s.String()
}

func (s *SearchAllTasksByTqlRequest) SetMaxResults(v int32) *SearchAllTasksByTqlRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchAllTasksByTqlRequest) SetNextToken(v string) *SearchAllTasksByTqlRequest {
	s.NextToken = &v
	return s
}

func (s *SearchAllTasksByTqlRequest) SetTql(v string) *SearchAllTasksByTqlRequest {
	s.Tql = &v
	return s
}

func (s *SearchAllTasksByTqlRequest) SetUserId(v string) *SearchAllTasksByTqlRequest {
	s.UserId = &v
	return s
}

type SearchAllTasksByTqlResponseBody struct {
	// example:
	//
	// f279e812-e431-428d-846d-cxxxxxx
	NextToken *string   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalSize *int32    `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s SearchAllTasksByTqlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchAllTasksByTqlResponseBody) GoString() string {
	return s.String()
}

func (s *SearchAllTasksByTqlResponseBody) SetNextToken(v string) *SearchAllTasksByTqlResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchAllTasksByTqlResponseBody) SetRequestId(v string) *SearchAllTasksByTqlResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchAllTasksByTqlResponseBody) SetResult(v []*string) *SearchAllTasksByTqlResponseBody {
	s.Result = v
	return s
}

func (s *SearchAllTasksByTqlResponseBody) SetTotalSize(v int32) *SearchAllTasksByTqlResponseBody {
	s.TotalSize = &v
	return s
}

type SearchAllTasksByTqlResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchAllTasksByTqlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchAllTasksByTqlResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchAllTasksByTqlResponse) GoString() string {
	return s.String()
}

func (s *SearchAllTasksByTqlResponse) SetHeaders(v map[string]*string) *SearchAllTasksByTqlResponse {
	s.Headers = v
	return s
}

func (s *SearchAllTasksByTqlResponse) SetStatusCode(v int32) *SearchAllTasksByTqlResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchAllTasksByTqlResponse) SetBody(v *SearchAllTasksByTqlResponseBody) *SearchAllTasksByTqlResponse {
	s.Body = v
	return s
}

type SearchProjectCustomFiledsV3Headers struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchProjectCustomFiledsV3Headers) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectCustomFiledsV3Headers) GoString() string {
	return s.String()
}

func (s *SearchProjectCustomFiledsV3Headers) SetCommonHeaders(v map[string]*string) *SearchProjectCustomFiledsV3Headers {
	s.CommonHeaders = v
	return s
}

func (s *SearchProjectCustomFiledsV3Headers) SetXAcsDingtalkAccessToken(v string) *SearchProjectCustomFiledsV3Headers {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchProjectCustomFiledsV3Request struct {
	CfIds      *string `json:"cfIds,omitempty" xml:"cfIds,omitempty"`
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	SfcId      *string `json:"sfcId,omitempty" xml:"sfcId,omitempty"`
}

func (s SearchProjectCustomFiledsV3Request) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectCustomFiledsV3Request) GoString() string {
	return s.String()
}

func (s *SearchProjectCustomFiledsV3Request) SetCfIds(v string) *SearchProjectCustomFiledsV3Request {
	s.CfIds = &v
	return s
}

func (s *SearchProjectCustomFiledsV3Request) SetMaxResults(v int32) *SearchProjectCustomFiledsV3Request {
	s.MaxResults = &v
	return s
}

func (s *SearchProjectCustomFiledsV3Request) SetNextToken(v string) *SearchProjectCustomFiledsV3Request {
	s.NextToken = &v
	return s
}

func (s *SearchProjectCustomFiledsV3Request) SetSfcId(v string) *SearchProjectCustomFiledsV3Request {
	s.SfcId = &v
	return s
}

type SearchProjectCustomFiledsV3ResponseBody struct {
	// example:
	//
	// f279e812-e431-428d-846d-cxxxxxx
	NextToken *string                                          `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId *string                                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*SearchProjectCustomFiledsV3ResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// 35
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s SearchProjectCustomFiledsV3ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectCustomFiledsV3ResponseBody) GoString() string {
	return s.String()
}

func (s *SearchProjectCustomFiledsV3ResponseBody) SetNextToken(v string) *SearchProjectCustomFiledsV3ResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchProjectCustomFiledsV3ResponseBody) SetRequestId(v string) *SearchProjectCustomFiledsV3ResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchProjectCustomFiledsV3ResponseBody) SetResult(v []*SearchProjectCustomFiledsV3ResponseBodyResult) *SearchProjectCustomFiledsV3ResponseBody {
	s.Result = v
	return s
}

func (s *SearchProjectCustomFiledsV3ResponseBody) SetTotalCount(v int32) *SearchProjectCustomFiledsV3ResponseBody {
	s.TotalCount = &v
	return s
}

type SearchProjectCustomFiledsV3ResponseBodyResult struct {
	AdvancedCustomfield *SearchProjectCustomFiledsV3ResponseBodyResultAdvancedCustomfield `json:"advancedCustomfield,omitempty" xml:"advancedCustomfield,omitempty" type:"Struct"`
	BoundToObjectId     *string                                                           `json:"boundToObjectId,omitempty" xml:"boundToObjectId,omitempty"`
	BoundToObjectType   *string                                                           `json:"boundToObjectType,omitempty" xml:"boundToObjectType,omitempty"`
	Choices             []*SearchProjectCustomFiledsV3ResponseBodyResultChoices           `json:"choices,omitempty" xml:"choices,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-07-04T03:29:34.770Z
	Created    *string `json:"created,omitempty" xml:"created,omitempty"`
	CreatorId  *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Id         *string `json:"id,omitempty" xml:"id,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	OriginalId *string `json:"originalId,omitempty" xml:"originalId,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SearchProjectCustomFiledsV3ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectCustomFiledsV3ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SearchProjectCustomFiledsV3ResponseBodyResult) SetAdvancedCustomfield(v *SearchProjectCustomFiledsV3ResponseBodyResultAdvancedCustomfield) *SearchProjectCustomFiledsV3ResponseBodyResult {
	s.AdvancedCustomfield = v
	return s
}

func (s *SearchProjectCustomFiledsV3ResponseBodyResult) SetBoundToObjectId(v string) *SearchProjectCustomFiledsV3ResponseBodyResult {
	s.BoundToObjectId = &v
	return s
}

func (s *SearchProjectCustomFiledsV3ResponseBodyResult) SetBoundToObjectType(v string) *SearchProjectCustomFiledsV3ResponseBodyResult {
	s.BoundToObjectType = &v
	return s
}

func (s *SearchProjectCustomFiledsV3ResponseBodyResult) SetChoices(v []*SearchProjectCustomFiledsV3ResponseBodyResultChoices) *SearchProjectCustomFiledsV3ResponseBodyResult {
	s.Choices = v
	return s
}

func (s *SearchProjectCustomFiledsV3ResponseBodyResult) SetCreated(v string) *SearchProjectCustomFiledsV3ResponseBodyResult {
	s.Created = &v
	return s
}

func (s *SearchProjectCustomFiledsV3ResponseBodyResult) SetCreatorId(v string) *SearchProjectCustomFiledsV3ResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *SearchProjectCustomFiledsV3ResponseBodyResult) SetId(v string) *SearchProjectCustomFiledsV3ResponseBodyResult {
	s.Id = &v
	return s
}

func (s *SearchProjectCustomFiledsV3ResponseBodyResult) SetName(v string) *SearchProjectCustomFiledsV3ResponseBodyResult {
	s.Name = &v
	return s
}

func (s *SearchProjectCustomFiledsV3ResponseBodyResult) SetOriginalId(v string) *SearchProjectCustomFiledsV3ResponseBodyResult {
	s.OriginalId = &v
	return s
}

func (s *SearchProjectCustomFiledsV3ResponseBodyResult) SetType(v string) *SearchProjectCustomFiledsV3ResponseBodyResult {
	s.Type = &v
	return s
}

type SearchProjectCustomFiledsV3ResponseBodyResultAdvancedCustomfield struct {
	Id         *string `json:"id,omitempty" xml:"id,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s SearchProjectCustomFiledsV3ResponseBodyResultAdvancedCustomfield) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectCustomFiledsV3ResponseBodyResultAdvancedCustomfield) GoString() string {
	return s.String()
}

func (s *SearchProjectCustomFiledsV3ResponseBodyResultAdvancedCustomfield) SetId(v string) *SearchProjectCustomFiledsV3ResponseBodyResultAdvancedCustomfield {
	s.Id = &v
	return s
}

func (s *SearchProjectCustomFiledsV3ResponseBodyResultAdvancedCustomfield) SetName(v string) *SearchProjectCustomFiledsV3ResponseBodyResultAdvancedCustomfield {
	s.Name = &v
	return s
}

func (s *SearchProjectCustomFiledsV3ResponseBodyResultAdvancedCustomfield) SetObjectType(v string) *SearchProjectCustomFiledsV3ResponseBodyResultAdvancedCustomfield {
	s.ObjectType = &v
	return s
}

type SearchProjectCustomFiledsV3ResponseBodyResultChoices struct {
	Id    *string `json:"id,omitempty" xml:"id,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s SearchProjectCustomFiledsV3ResponseBodyResultChoices) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectCustomFiledsV3ResponseBodyResultChoices) GoString() string {
	return s.String()
}

func (s *SearchProjectCustomFiledsV3ResponseBodyResultChoices) SetId(v string) *SearchProjectCustomFiledsV3ResponseBodyResultChoices {
	s.Id = &v
	return s
}

func (s *SearchProjectCustomFiledsV3ResponseBodyResultChoices) SetValue(v string) *SearchProjectCustomFiledsV3ResponseBodyResultChoices {
	s.Value = &v
	return s
}

type SearchProjectCustomFiledsV3Response struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchProjectCustomFiledsV3ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchProjectCustomFiledsV3Response) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectCustomFiledsV3Response) GoString() string {
	return s.String()
}

func (s *SearchProjectCustomFiledsV3Response) SetHeaders(v map[string]*string) *SearchProjectCustomFiledsV3Response {
	s.Headers = v
	return s
}

func (s *SearchProjectCustomFiledsV3Response) SetStatusCode(v int32) *SearchProjectCustomFiledsV3Response {
	s.StatusCode = &v
	return s
}

func (s *SearchProjectCustomFiledsV3Response) SetBody(v *SearchProjectCustomFiledsV3ResponseBody) *SearchProjectCustomFiledsV3Response {
	s.Body = v
	return s
}

type SearchProjectsV3Headers struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchProjectsV3Headers) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectsV3Headers) GoString() string {
	return s.String()
}

func (s *SearchProjectsV3Headers) SetCommonHeaders(v map[string]*string) *SearchProjectsV3Headers {
	s.CommonHeaders = v
	return s
}

func (s *SearchProjectsV3Headers) SetXAcsDingtalkAccessToken(v string) *SearchProjectsV3Headers {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchProjectsV3Request struct {
	IncludeTemplate *bool   `json:"includeTemplate,omitempty" xml:"includeTemplate,omitempty"`
	MaxResults      *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	NextToken       *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ProjectIds      *string `json:"projectIds,omitempty" xml:"projectIds,omitempty"`
	SourceId        *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0517xxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchProjectsV3Request) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectsV3Request) GoString() string {
	return s.String()
}

func (s *SearchProjectsV3Request) SetIncludeTemplate(v bool) *SearchProjectsV3Request {
	s.IncludeTemplate = &v
	return s
}

func (s *SearchProjectsV3Request) SetMaxResults(v int32) *SearchProjectsV3Request {
	s.MaxResults = &v
	return s
}

func (s *SearchProjectsV3Request) SetName(v string) *SearchProjectsV3Request {
	s.Name = &v
	return s
}

func (s *SearchProjectsV3Request) SetNextToken(v string) *SearchProjectsV3Request {
	s.NextToken = &v
	return s
}

func (s *SearchProjectsV3Request) SetProjectIds(v string) *SearchProjectsV3Request {
	s.ProjectIds = &v
	return s
}

func (s *SearchProjectsV3Request) SetSourceId(v string) *SearchProjectsV3Request {
	s.SourceId = &v
	return s
}

func (s *SearchProjectsV3Request) SetUserId(v string) *SearchProjectsV3Request {
	s.UserId = &v
	return s
}

type SearchProjectsV3ResponseBody struct {
	// example:
	//
	// f279e812-e431-428d-846d-cxxxxxx
	NextToken *string                               `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*SearchProjectsV3ResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s SearchProjectsV3ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectsV3ResponseBody) GoString() string {
	return s.String()
}

func (s *SearchProjectsV3ResponseBody) SetNextToken(v string) *SearchProjectsV3ResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchProjectsV3ResponseBody) SetRequestId(v string) *SearchProjectsV3ResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchProjectsV3ResponseBody) SetResult(v []*SearchProjectsV3ResponseBodyResult) *SearchProjectsV3ResponseBody {
	s.Result = v
	return s
}

type SearchProjectsV3ResponseBodyResult struct {
	// example:
	//
	// 2022-07-04T03:29:34.770Z
	Created        *string `json:"created,omitempty" xml:"created,omitempty"`
	CreatorId      *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	Id             *string `json:"id,omitempty" xml:"id,omitempty"`
	IsArchived     *bool   `json:"isArchived,omitempty" xml:"isArchived,omitempty"`
	IsTemplate     *bool   `json:"isTemplate,omitempty" xml:"isTemplate,omitempty"`
	Logo           *string `json:"logo,omitempty" xml:"logo,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	SourceId       *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// 2022-07-04T03:29:34.770Z
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s SearchProjectsV3ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectsV3ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SearchProjectsV3ResponseBodyResult) SetCreated(v string) *SearchProjectsV3ResponseBodyResult {
	s.Created = &v
	return s
}

func (s *SearchProjectsV3ResponseBodyResult) SetCreatorId(v string) *SearchProjectsV3ResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *SearchProjectsV3ResponseBodyResult) SetDescription(v string) *SearchProjectsV3ResponseBodyResult {
	s.Description = &v
	return s
}

func (s *SearchProjectsV3ResponseBodyResult) SetId(v string) *SearchProjectsV3ResponseBodyResult {
	s.Id = &v
	return s
}

func (s *SearchProjectsV3ResponseBodyResult) SetIsArchived(v bool) *SearchProjectsV3ResponseBodyResult {
	s.IsArchived = &v
	return s
}

func (s *SearchProjectsV3ResponseBodyResult) SetIsTemplate(v bool) *SearchProjectsV3ResponseBodyResult {
	s.IsTemplate = &v
	return s
}

func (s *SearchProjectsV3ResponseBodyResult) SetLogo(v string) *SearchProjectsV3ResponseBodyResult {
	s.Logo = &v
	return s
}

func (s *SearchProjectsV3ResponseBodyResult) SetName(v string) *SearchProjectsV3ResponseBodyResult {
	s.Name = &v
	return s
}

func (s *SearchProjectsV3ResponseBodyResult) SetOrganizationId(v string) *SearchProjectsV3ResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *SearchProjectsV3ResponseBodyResult) SetSourceId(v string) *SearchProjectsV3ResponseBodyResult {
	s.SourceId = &v
	return s
}

func (s *SearchProjectsV3ResponseBodyResult) SetUpdated(v string) *SearchProjectsV3ResponseBodyResult {
	s.Updated = &v
	return s
}

type SearchProjectsV3Response struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchProjectsV3ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchProjectsV3Response) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectsV3Response) GoString() string {
	return s.String()
}

func (s *SearchProjectsV3Response) SetHeaders(v map[string]*string) *SearchProjectsV3Response {
	s.Headers = v
	return s
}

func (s *SearchProjectsV3Response) SetStatusCode(v int32) *SearchProjectsV3Response {
	s.StatusCode = &v
	return s
}

func (s *SearchProjectsV3Response) SetBody(v *SearchProjectsV3ResponseBody) *SearchProjectsV3Response {
	s.Body = v
	return s
}

type UpdateProjectMemberRoleV3Headers struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateProjectMemberRoleV3Headers) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectMemberRoleV3Headers) GoString() string {
	return s.String()
}

func (s *UpdateProjectMemberRoleV3Headers) SetCommonHeaders(v map[string]*string) *UpdateProjectMemberRoleV3Headers {
	s.CommonHeaders = v
	return s
}

func (s *UpdateProjectMemberRoleV3Headers) SetXAcsDingtalkAccessToken(v string) *UpdateProjectMemberRoleV3Headers {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateProjectMemberRoleV3Request struct {
	RoleIds []*string `json:"roleIds,omitempty" xml:"roleIds,omitempty" type:"Repeated"`
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s UpdateProjectMemberRoleV3Request) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectMemberRoleV3Request) GoString() string {
	return s.String()
}

func (s *UpdateProjectMemberRoleV3Request) SetRoleIds(v []*string) *UpdateProjectMemberRoleV3Request {
	s.RoleIds = v
	return s
}

func (s *UpdateProjectMemberRoleV3Request) SetUserIds(v []*string) *UpdateProjectMemberRoleV3Request {
	s.UserIds = v
	return s
}

type UpdateProjectMemberRoleV3ResponseBody struct {
	RequestId *string                                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*UpdateProjectMemberRoleV3ResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s UpdateProjectMemberRoleV3ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectMemberRoleV3ResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectMemberRoleV3ResponseBody) SetRequestId(v string) *UpdateProjectMemberRoleV3ResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectMemberRoleV3ResponseBody) SetResult(v []*UpdateProjectMemberRoleV3ResponseBodyResult) *UpdateProjectMemberRoleV3ResponseBody {
	s.Result = v
	return s
}

type UpdateProjectMemberRoleV3ResponseBodyResult struct {
	Id      *string   `json:"id,omitempty" xml:"id,omitempty"`
	Role    *int32    `json:"role,omitempty" xml:"role,omitempty"`
	RoleIds []*string `json:"roleIds,omitempty" xml:"roleIds,omitempty" type:"Repeated"`
	UserId  *string   `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateProjectMemberRoleV3ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectMemberRoleV3ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateProjectMemberRoleV3ResponseBodyResult) SetId(v string) *UpdateProjectMemberRoleV3ResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateProjectMemberRoleV3ResponseBodyResult) SetRole(v int32) *UpdateProjectMemberRoleV3ResponseBodyResult {
	s.Role = &v
	return s
}

func (s *UpdateProjectMemberRoleV3ResponseBodyResult) SetRoleIds(v []*string) *UpdateProjectMemberRoleV3ResponseBodyResult {
	s.RoleIds = v
	return s
}

func (s *UpdateProjectMemberRoleV3ResponseBodyResult) SetUserId(v string) *UpdateProjectMemberRoleV3ResponseBodyResult {
	s.UserId = &v
	return s
}

type UpdateProjectMemberRoleV3Response struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProjectMemberRoleV3ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProjectMemberRoleV3Response) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectMemberRoleV3Response) GoString() string {
	return s.String()
}

func (s *UpdateProjectMemberRoleV3Response) SetHeaders(v map[string]*string) *UpdateProjectMemberRoleV3Response {
	s.Headers = v
	return s
}

func (s *UpdateProjectMemberRoleV3Response) SetStatusCode(v int32) *UpdateProjectMemberRoleV3Response {
	s.StatusCode = &v
	return s
}

func (s *UpdateProjectMemberRoleV3Response) SetBody(v *UpdateProjectMemberRoleV3ResponseBody) *UpdateProjectMemberRoleV3Response {
	s.Body = v
	return s
}

type UpdateProjectV3Headers struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateProjectV3Headers) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectV3Headers) GoString() string {
	return s.String()
}

func (s *UpdateProjectV3Headers) SetCommonHeaders(v map[string]*string) *UpdateProjectV3Headers {
	s.CommonHeaders = v
	return s
}

func (s *UpdateProjectV3Headers) SetXAcsDingtalkAccessToken(v string) *UpdateProjectV3Headers {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateProjectV3Request struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateProjectV3Request) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectV3Request) GoString() string {
	return s.String()
}

func (s *UpdateProjectV3Request) SetDescription(v string) *UpdateProjectV3Request {
	s.Description = &v
	return s
}

func (s *UpdateProjectV3Request) SetName(v string) *UpdateProjectV3Request {
	s.Name = &v
	return s
}

type UpdateProjectV3ResponseBody struct {
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *UpdateProjectV3ResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateProjectV3ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectV3ResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectV3ResponseBody) SetRequestId(v string) *UpdateProjectV3ResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectV3ResponseBody) SetResult(v *UpdateProjectV3ResponseBodyResult) *UpdateProjectV3ResponseBody {
	s.Result = v
	return s
}

type UpdateProjectV3ResponseBodyResult struct {
	// example:
	//
	// 2022-07-04T03:29:34.770Z
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateProjectV3ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectV3ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateProjectV3ResponseBodyResult) SetUpdated(v string) *UpdateProjectV3ResponseBodyResult {
	s.Updated = &v
	return s
}

type UpdateProjectV3Response struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProjectV3ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProjectV3Response) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectV3Response) GoString() string {
	return s.String()
}

func (s *UpdateProjectV3Response) SetHeaders(v map[string]*string) *UpdateProjectV3Response {
	s.Headers = v
	return s
}

func (s *UpdateProjectV3Response) SetStatusCode(v int32) *UpdateProjectV3Response {
	s.StatusCode = &v
	return s
}

func (s *UpdateProjectV3Response) SetBody(v *UpdateProjectV3ResponseBody) *UpdateProjectV3Response {
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
// 查询任务概览
//
// @param request - AnalysisReportRequest
//
// @param headers - AnalysisReportHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AnalysisReportResponse
func (client *Client) AnalysisReportWithOptions(userId *string, request *AnalysisReportRequest, headers *AnalysisReportHeaders, runtime *util.RuntimeOptions) (_result *AnalysisReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		body["filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.ReportId)) {
		body["reportId"] = request.ReportId
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
		Action:      tea.String("AnalysisReport"),
		Version:     tea.String("teamSphere_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/teamSphere/users/" + tea.StringValue(userId) + "/analyses/report"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AnalysisReportResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询任务概览
//
// @param request - AnalysisReportRequest
//
// @return AnalysisReportResponse
func (client *Client) AnalysisReport(userId *string, request *AnalysisReportRequest) (_result *AnalysisReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AnalysisReportHeaders{}
	_result = &AnalysisReportResponse{}
	_body, _err := client.AnalysisReportWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建自由任务
//
// @param request - CreateOrganizationTaskRequest
//
// @param headers - CreateOrganizationTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrganizationTaskResponse
func (client *Client) CreateOrganizationTaskWithOptions(userId *string, request *CreateOrganizationTaskRequest, headers *CreateOrganizationTaskHeaders, runtime *util.RuntimeOptions) (_result *CreateOrganizationTaskResponse, _err error) {
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
		Version:     tea.String("teamSphere_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/teamSphere/organizations/users/" + tea.StringValue(userId) + "/tasks"),
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

// Summary:
//
// 创建自由任务
//
// @param request - CreateOrganizationTaskRequest
//
// @return CreateOrganizationTaskResponse
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

// Summary:
//
// 创建项目成员
//
// @param request - CreateProjectMembersV3Request
//
// @param headers - CreateProjectMembersV3Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectMembersV3Response
func (client *Client) CreateProjectMembersV3WithOptions(userId *string, projectId *string, request *CreateProjectMembersV3Request, headers *CreateProjectMembersV3Headers, runtime *util.RuntimeOptions) (_result *CreateProjectMembersV3Response, _err error) {
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
		Action:      tea.String("CreateProjectMembersV3"),
		Version:     tea.String("teamSphere_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/teamSphere/users/" + tea.StringValue(userId) + "/projects/" + tea.StringValue(projectId) + "/members"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProjectMembersV3Response{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建项目成员
//
// @param request - CreateProjectMembersV3Request
//
// @return CreateProjectMembersV3Response
func (client *Client) CreateProjectMembersV3(userId *string, projectId *string, request *CreateProjectMembersV3Request) (_result *CreateProjectMembersV3Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateProjectMembersV3Headers{}
	_result = &CreateProjectMembersV3Response{}
	_body, _err := client.CreateProjectMembersV3WithOptions(userId, projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建协作空间。
//
// @param request - CreateProjectV3Request
//
// @param headers - CreateProjectV3Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectV3Response
func (client *Client) CreateProjectV3WithOptions(userId *string, request *CreateProjectV3Request, headers *CreateProjectV3Headers, runtime *util.RuntimeOptions) (_result *CreateProjectV3Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["organizationId"] = request.OrganizationId
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
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProjectV3"),
		Version:     tea.String("teamSphere_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/teamSphere/users/" + tea.StringValue(userId) + "/projects"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProjectV3Response{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建协作空间。
//
// @param request - CreateProjectV3Request
//
// @return CreateProjectV3Response
func (client *Client) CreateProjectV3(userId *string, request *CreateProjectV3Request) (_result *CreateProjectV3Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateProjectV3Headers{}
	_result = &CreateProjectV3Response{}
	_body, _err := client.CreateProjectV3WithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建协作空间任务
//
// @param request - CreateTaskRequest
//
// @param headers - CreateTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTaskResponse
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

	if !tea.BoolValue(util.IsUnset(request.Note)) {
		body["note"] = request.Note
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
	params := &openapi.Params{
		Action:      tea.String("CreateTask"),
		Version:     tea.String("teamSphere_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/teamSphere/users/" + tea.StringValue(userId) + "/tasks"),
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

// Summary:
//
// 创建协作空间任务
//
// @param request - CreateTaskRequest
//
// @return CreateTaskResponse
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

// Summary:
//
// 删除项目成员。
//
// @param request - DeleteProjectMembersV3Request
//
// @param headers - DeleteProjectMembersV3Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectMembersV3Response
func (client *Client) DeleteProjectMembersV3WithOptions(userId *string, projectId *string, request *DeleteProjectMembersV3Request, headers *DeleteProjectMembersV3Headers, runtime *util.RuntimeOptions) (_result *DeleteProjectMembersV3Response, _err error) {
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
		Action:      tea.String("DeleteProjectMembersV3"),
		Version:     tea.String("teamSphere_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/teamSphere/users/" + tea.StringValue(userId) + "/projects/" + tea.StringValue(projectId) + "/members/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProjectMembersV3Response{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除项目成员。
//
// @param request - DeleteProjectMembersV3Request
//
// @return DeleteProjectMembersV3Response
func (client *Client) DeleteProjectMembersV3(userId *string, projectId *string, request *DeleteProjectMembersV3Request) (_result *DeleteProjectMembersV3Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteProjectMembersV3Headers{}
	_result = &DeleteProjectMembersV3Response{}
	_body, _err := client.DeleteProjectMembersV3WithOptions(userId, projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取最近访问的项目
//
// @param headers - GetFootprintProjectHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFootprintProjectResponse
func (client *Client) GetFootprintProjectWithOptions(userId *string, headers *GetFootprintProjectHeaders, runtime *util.RuntimeOptions) (_result *GetFootprintProjectResponse, _err error) {
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
		Action:      tea.String("GetFootprintProject"),
		Version:     tea.String("teamSphere_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/teamSphere/users/" + tea.StringValue(userId) + "/footprints/projects"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFootprintProjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取最近访问的项目
//
// @return GetFootprintProjectResponse
func (client *Client) GetFootprintProject(userId *string) (_result *GetFootprintProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFootprintProjectHeaders{}
	_result = &GetFootprintProjectResponse{}
	_body, _err := client.GetFootprintProjectWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取最近访问的任务
//
// @param headers - GetFootprintTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFootprintTaskResponse
func (client *Client) GetFootprintTaskWithOptions(userId *string, headers *GetFootprintTaskHeaders, runtime *util.RuntimeOptions) (_result *GetFootprintTaskResponse, _err error) {
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
		Action:      tea.String("GetFootprintTask"),
		Version:     tea.String("teamSphere_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/teamSphere/users/" + tea.StringValue(userId) + "/footprints/tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFootprintTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取最近访问的任务
//
// @return GetFootprintTaskResponse
func (client *Client) GetFootprintTask(userId *string) (_result *GetFootprintTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFootprintTaskHeaders{}
	_result = &GetFootprintTaskResponse{}
	_body, _err := client.GetFootprintTaskWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询轻任务详情。
//
// @param request - GetFreeTaskRequest
//
// @param headers - GetFreeTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFreeTaskResponse
func (client *Client) GetFreeTaskWithOptions(taskId *string, request *GetFreeTaskRequest, headers *GetFreeTaskHeaders, runtime *util.RuntimeOptions) (_result *GetFreeTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      tea.String("GetFreeTask"),
		Version:     tea.String("teamSphere_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/teamSphere/organizations/tasks/" + tea.StringValue(taskId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFreeTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询轻任务详情。
//
// @param request - GetFreeTaskRequest
//
// @return GetFreeTaskResponse
func (client *Client) GetFreeTask(taskId *string, request *GetFreeTaskRequest) (_result *GetFreeTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFreeTaskHeaders{}
	_result = &GetFreeTaskResponse{}
	_body, _err := client.GetFreeTaskWithOptions(taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取协作空间成员列表。
//
// @param request - GetProjectMembersV3Request
//
// @param headers - GetProjectMembersV3Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectMembersV3Response
func (client *Client) GetProjectMembersV3WithOptions(userId *string, projectId *string, request *GetProjectMembersV3Request, headers *GetProjectMembersV3Headers, runtime *util.RuntimeOptions) (_result *GetProjectMembersV3Response, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ProjectRoleId)) {
		query["projectRoleId"] = request.ProjectRoleId
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
		Action:      tea.String("GetProjectMembersV3"),
		Version:     tea.String("teamSphere_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/teamSphere/users/" + tea.StringValue(userId) + "/projects/" + tea.StringValue(projectId) + "/members"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectMembersV3Response{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取协作空间成员列表。
//
// @param request - GetProjectMembersV3Request
//
// @return GetProjectMembersV3Response
func (client *Client) GetProjectMembersV3(userId *string, projectId *string, request *GetProjectMembersV3Request) (_result *GetProjectMembersV3Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetProjectMembersV3Headers{}
	_result = &GetProjectMembersV3Response{}
	_body, _err := client.GetProjectMembersV3WithOptions(userId, projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取项目角色列表。
//
// @param request - GetProjectRolesV3Request
//
// @param headers - GetProjectRolesV3Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectRolesV3Response
func (client *Client) GetProjectRolesV3WithOptions(userId *string, projectId *string, request *GetProjectRolesV3Request, headers *GetProjectRolesV3Headers, runtime *util.RuntimeOptions) (_result *GetProjectRolesV3Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncludeHidden)) {
		query["includeHidden"] = request.IncludeHidden
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		query["level"] = request.Level
	}

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
		Action:      tea.String("GetProjectRolesV3"),
		Version:     tea.String("teamSphere_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/teamSphere/users/" + tea.StringValue(userId) + "/projects/" + tea.StringValue(projectId) + "/roles"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectRolesV3Response{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取项目角色列表。
//
// @param request - GetProjectRolesV3Request
//
// @return GetProjectRolesV3Response
func (client *Client) GetProjectRolesV3(userId *string, projectId *string, request *GetProjectRolesV3Request) (_result *GetProjectRolesV3Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetProjectRolesV3Headers{}
	_result = &GetProjectRolesV3Response{}
	_body, _err := client.GetProjectRolesV3WithOptions(userId, projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户星标协作空间
//
// @param request - GetStaredProjectsRequest
//
// @param headers - GetStaredProjectsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStaredProjectsResponse
func (client *Client) GetStaredProjectsWithOptions(userId *string, request *GetStaredProjectsRequest, headers *GetStaredProjectsHeaders, runtime *util.RuntimeOptions) (_result *GetStaredProjectsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["sortBy"] = request.SortBy
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
		Action:      tea.String("GetStaredProjects"),
		Version:     tea.String("teamSphere_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/teamSphere/users/" + tea.StringValue(userId) + "/staredProjects"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStaredProjectsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户星标协作空间
//
// @param request - GetStaredProjectsRequest
//
// @return GetStaredProjectsResponse
func (client *Client) GetStaredProjects(userId *string, request *GetStaredProjectsRequest) (_result *GetStaredProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetStaredProjectsHeaders{}
	_result = &GetStaredProjectsResponse{}
	_body, _err := client.GetStaredProjectsWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 钉钉 userId 查询 24位长 userId。
//
// @param request - GetTbUserIdByDingUserIdRequest
//
// @param headers - GetTbUserIdByDingUserIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTbUserIdByDingUserIdResponse
func (client *Client) GetTbUserIdByDingUserIdWithOptions(request *GetTbUserIdByDingUserIdRequest, headers *GetTbUserIdByDingUserIdHeaders, runtime *util.RuntimeOptions) (_result *GetTbUserIdByDingUserIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingUserIds)) {
		query["dingUserIds"] = request.DingUserIds
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
		Action:      tea.String("GetTbUserIdByDingUserId"),
		Version:     tea.String("teamSphere_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/teamSphere/idmaps/userIds"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTbUserIdByDingUserIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 钉钉 userId 查询 24位长 userId。
//
// @param request - GetTbUserIdByDingUserIdRequest
//
// @return GetTbUserIdByDingUserIdResponse
func (client *Client) GetTbUserIdByDingUserId(request *GetTbUserIdByDingUserIdRequest) (_result *GetTbUserIdByDingUserIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTbUserIdByDingUserIdHeaders{}
	_result = &GetTbUserIdByDingUserIdResponse{}
	_body, _err := client.GetTbUserIdByDingUserIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取快办企业ID
//
// @param headers - GetThingOrgIdByDingOrgIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetThingOrgIdByDingOrgIdResponse
func (client *Client) GetThingOrgIdByDingOrgIdWithOptions(headers *GetThingOrgIdByDingOrgIdHeaders, runtime *util.RuntimeOptions) (_result *GetThingOrgIdByDingOrgIdResponse, _err error) {
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
		Action:      tea.String("GetThingOrgIdByDingOrgId"),
		Version:     tea.String("teamSphere_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/teamSphere/organizations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetThingOrgIdByDingOrgIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取快办企业ID
//
// @return GetThingOrgIdByDingOrgIdResponse
func (client *Client) GetThingOrgIdByDingOrgId() (_result *GetThingOrgIdByDingOrgIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetThingOrgIdByDingOrgIdHeaders{}
	_result = &GetThingOrgIdByDingOrgIdResponse{}
	_body, _err := client.GetThingOrgIdByDingOrgIdWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户参与项目。
//
// @param request - GetUserJoinedProjectsV3Request
//
// @param headers - GetUserJoinedProjectsV3Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserJoinedProjectsV3Response
func (client *Client) GetUserJoinedProjectsV3WithOptions(userId *string, request *GetUserJoinedProjectsV3Request, headers *GetUserJoinedProjectsV3Headers, runtime *util.RuntimeOptions) (_result *GetUserJoinedProjectsV3Response, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ProjectIds)) {
		query["projectIds"] = request.ProjectIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectRoleLevels)) {
		query["projectRoleLevels"] = request.ProjectRoleLevels
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["sortBy"] = request.SortBy
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
		Action:      tea.String("GetUserJoinedProjectsV3"),
		Version:     tea.String("teamSphere_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/teamSphere/users/" + tea.StringValue(userId) + "/projects/userJoined"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserJoinedProjectsV3Response{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户参与项目。
//
// @param request - GetUserJoinedProjectsV3Request
//
// @return GetUserJoinedProjectsV3Response
func (client *Client) GetUserJoinedProjectsV3(userId *string, request *GetUserJoinedProjectsV3Request) (_result *GetUserJoinedProjectsV3Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserJoinedProjectsV3Headers{}
	_result = &GetUserJoinedProjectsV3Response{}
	_body, _err := client.GetUserJoinedProjectsV3WithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取全部任务
//
// @param headers - ListAllTaskViewHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllTaskViewResponse
func (client *Client) ListAllTaskViewWithOptions(userId *string, headers *ListAllTaskViewHeaders, runtime *util.RuntimeOptions) (_result *ListAllTaskViewResponse, _err error) {
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
		Action:      tea.String("ListAllTaskView"),
		Version:     tea.String("teamSphere_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/teamSphere/users/" + tea.StringValue(userId) + "/allTaskViews"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAllTaskViewResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取全部任务
//
// @return ListAllTaskViewResponse
func (client *Client) ListAllTaskView(userId *string) (_result *ListAllTaskViewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListAllTaskViewHeaders{}
	_result = &ListAllTaskViewResponse{}
	_body, _err := client.ListAllTaskViewWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询我的捷径
//
// @param request - ListMyShortcutViewsRequest
//
// @param headers - ListMyShortcutViewsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMyShortcutViewsResponse
func (client *Client) ListMyShortcutViewsWithOptions(userId *string, request *ListMyShortcutViewsRequest, headers *ListMyShortcutViewsHeaders, runtime *util.RuntimeOptions) (_result *ListMyShortcutViewsResponse, _err error) {
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
		Action:      tea.String("ListMyShortcutViews"),
		Version:     tea.String("teamSphere_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/teamSphere/users/" + tea.StringValue(userId) + "/shortcutViews"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMyShortcutViewsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询我的捷径
//
// @param request - ListMyShortcutViewsRequest
//
// @return ListMyShortcutViewsResponse
func (client *Client) ListMyShortcutViews(userId *string, request *ListMyShortcutViewsRequest) (_result *ListMyShortcutViewsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListMyShortcutViewsHeaders{}
	_result = &ListMyShortcutViewsResponse{}
	_body, _err := client.ListMyShortcutViewsWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询自由任务和项目任务详情。
//
// @param request - QueryAllTaskRequest
//
// @param headers - QueryAllTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAllTaskResponse
func (client *Client) QueryAllTaskWithOptions(userId *string, request *QueryAllTaskRequest, headers *QueryAllTaskHeaders, runtime *util.RuntimeOptions) (_result *QueryAllTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      tea.String("QueryAllTask"),
		Version:     tea.String("teamSphere_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/teamSphere/users/" + tea.StringValue(userId) + "/tasks/query"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAllTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询自由任务和项目任务详情。
//
// @param request - QueryAllTaskRequest
//
// @return QueryAllTaskResponse
func (client *Client) QueryAllTask(userId *string, request *QueryAllTaskRequest) (_result *QueryAllTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllTaskHeaders{}
	_result = &QueryAllTaskResponse{}
	_body, _err := client.QueryAllTaskWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询我的任务
//
// @param request - QueryTaskRequest
//
// @param headers - QueryTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTaskResponse
func (client *Client) QueryTaskWithOptions(userId *string, request *QueryTaskRequest, headers *QueryTaskHeaders, runtime *util.RuntimeOptions) (_result *QueryTaskResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Tql)) {
		body["tql"] = request.Tql
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
		Action:      tea.String("QueryTask"),
		Version:     tea.String("teamSphere_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/teamSphere/users/" + tea.StringValue(userId) + "/tasks/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询我的任务
//
// @param request - QueryTaskRequest
//
// @return QueryTaskResponse
func (client *Client) QueryTask(userId *string, request *QueryTaskRequest) (_result *QueryTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryTaskHeaders{}
	_result = &QueryTaskResponse{}
	_body, _err := client.QueryTaskWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询协作空间任务详情。
//
// @param request - QueryTasksV3Request
//
// @param headers - QueryTasksV3Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTasksV3Response
func (client *Client) QueryTasksV3WithOptions(userId *string, request *QueryTasksV3Request, headers *QueryTasksV3Headers, runtime *util.RuntimeOptions) (_result *QueryTasksV3Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      tea.String("QueryTasksV3"),
		Version:     tea.String("teamSphere_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/teamSphere/user/" + tea.StringValue(userId) + "/tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTasksV3Response{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询协作空间任务详情。
//
// @param request - QueryTasksV3Request
//
// @return QueryTasksV3Response
func (client *Client) QueryTasksV3(userId *string, request *QueryTasksV3Request) (_result *QueryTasksV3Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryTasksV3Headers{}
	_result = &QueryTasksV3Response{}
	_body, _err := client.QueryTasksV3WithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过TQL搜索自由任务和协作空间任务ID。
//
// @param request - SearchAllTasksByTqlRequest
//
// @param headers - SearchAllTasksByTqlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchAllTasksByTqlResponse
func (client *Client) SearchAllTasksByTqlWithOptions(request *SearchAllTasksByTqlRequest, headers *SearchAllTasksByTqlHeaders, runtime *util.RuntimeOptions) (_result *SearchAllTasksByTqlResponse, _err error) {
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
		Action:      tea.String("SearchAllTasksByTql"),
		Version:     tea.String("teamSphere_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/teamSphere/taskIds"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchAllTasksByTqlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过TQL搜索自由任务和协作空间任务ID。
//
// @param request - SearchAllTasksByTqlRequest
//
// @return SearchAllTasksByTqlResponse
func (client *Client) SearchAllTasksByTql(request *SearchAllTasksByTqlRequest) (_result *SearchAllTasksByTqlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchAllTasksByTqlHeaders{}
	_result = &SearchAllTasksByTqlResponse{}
	_body, _err := client.SearchAllTasksByTqlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 搜索项目自定义字段。
//
// @param request - SearchProjectCustomFiledsV3Request
//
// @param headers - SearchProjectCustomFiledsV3Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchProjectCustomFiledsV3Response
func (client *Client) SearchProjectCustomFiledsV3WithOptions(userId *string, projectId *string, request *SearchProjectCustomFiledsV3Request, headers *SearchProjectCustomFiledsV3Headers, runtime *util.RuntimeOptions) (_result *SearchProjectCustomFiledsV3Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CfIds)) {
		query["cfIds"] = request.CfIds
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.SfcId)) {
		query["sfcId"] = request.SfcId
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
		Action:      tea.String("SearchProjectCustomFiledsV3"),
		Version:     tea.String("teamSphere_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/teamSphere/users/" + tea.StringValue(userId) + "/projects/" + tea.StringValue(projectId) + "/customFields"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchProjectCustomFiledsV3Response{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索项目自定义字段。
//
// @param request - SearchProjectCustomFiledsV3Request
//
// @return SearchProjectCustomFiledsV3Response
func (client *Client) SearchProjectCustomFiledsV3(userId *string, projectId *string, request *SearchProjectCustomFiledsV3Request) (_result *SearchProjectCustomFiledsV3Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchProjectCustomFiledsV3Headers{}
	_result = &SearchProjectCustomFiledsV3Response{}
	_body, _err := client.SearchProjectCustomFiledsV3WithOptions(userId, projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询协作空间。
//
// @param request - SearchProjectsV3Request
//
// @param headers - SearchProjectsV3Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchProjectsV3Response
func (client *Client) SearchProjectsV3WithOptions(request *SearchProjectsV3Request, headers *SearchProjectsV3Headers, runtime *util.RuntimeOptions) (_result *SearchProjectsV3Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncludeTemplate)) {
		query["includeTemplate"] = request.IncludeTemplate
	}

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
		Action:      tea.String("SearchProjectsV3"),
		Version:     tea.String("teamSphere_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/teamSphere/projects"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchProjectsV3Response{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询协作空间。
//
// @param request - SearchProjectsV3Request
//
// @return SearchProjectsV3Response
func (client *Client) SearchProjectsV3(request *SearchProjectsV3Request) (_result *SearchProjectsV3Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchProjectsV3Headers{}
	_result = &SearchProjectsV3Response{}
	_body, _err := client.SearchProjectsV3WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改项目成员的角色。
//
// @param request - UpdateProjectMemberRoleV3Request
//
// @param headers - UpdateProjectMemberRoleV3Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectMemberRoleV3Response
func (client *Client) UpdateProjectMemberRoleV3WithOptions(userId *string, projectId *string, request *UpdateProjectMemberRoleV3Request, headers *UpdateProjectMemberRoleV3Headers, runtime *util.RuntimeOptions) (_result *UpdateProjectMemberRoleV3Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleIds)) {
		body["roleIds"] = request.RoleIds
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
		Action:      tea.String("UpdateProjectMemberRoleV3"),
		Version:     tea.String("teamSphere_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/teamSphere/users/" + tea.StringValue(userId) + "/projects/" + tea.StringValue(projectId) + "/roles/assign"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProjectMemberRoleV3Response{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改项目成员的角色。
//
// @param request - UpdateProjectMemberRoleV3Request
//
// @return UpdateProjectMemberRoleV3Response
func (client *Client) UpdateProjectMemberRoleV3(userId *string, projectId *string, request *UpdateProjectMemberRoleV3Request) (_result *UpdateProjectMemberRoleV3Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateProjectMemberRoleV3Headers{}
	_result = &UpdateProjectMemberRoleV3Response{}
	_body, _err := client.UpdateProjectMemberRoleV3WithOptions(userId, projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新协作空间。
//
// @param request - UpdateProjectV3Request
//
// @param headers - UpdateProjectV3Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectV3Response
func (client *Client) UpdateProjectV3WithOptions(userId *string, projectId *string, request *UpdateProjectV3Request, headers *UpdateProjectV3Headers, runtime *util.RuntimeOptions) (_result *UpdateProjectV3Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

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
		Action:      tea.String("UpdateProjectV3"),
		Version:     tea.String("teamSphere_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/teamSphere/users/" + tea.StringValue(userId) + "/projects/" + tea.StringValue(projectId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProjectV3Response{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新协作空间。
//
// @param request - UpdateProjectV3Request
//
// @return UpdateProjectV3Response
func (client *Client) UpdateProjectV3(userId *string, projectId *string, request *UpdateProjectV3Request) (_result *UpdateProjectV3Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateProjectV3Headers{}
	_result = &UpdateProjectV3Response{}
	_body, _err := client.UpdateProjectV3WithOptions(userId, projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
