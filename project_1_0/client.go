// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package project_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
