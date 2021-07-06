// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package industry_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type QueryUserInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryUserInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserInfoResponseBody struct {
	// 人员详情
	Content *QueryUserInfoResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
}

func (s QueryUserInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponseBody) SetContent(v *QueryUserInfoResponseBodyContent) *QueryUserInfoResponseBody {
	s.Content = v
	return s
}

type QueryUserInfoResponseBodyContent struct {
	// 用户Id
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// 用户名称
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// 职称标签
	Job *QueryUserInfoResponseBodyContentJob `json:"job,omitempty" xml:"job,omitempty" type:"Struct"`
	// 工号
	JobNum *string `json:"jobNum,omitempty" xml:"jobNum,omitempty"`
	// 工作状态标签
	JobStatus *QueryUserInfoResponseBodyContentJobStatus `json:"jobStatus,omitempty" xml:"jobStatus,omitempty" type:"Struct"`
	// 人员属性标签
	UserProb *QueryUserInfoResponseBodyContentUserProb `json:"userProb,omitempty" xml:"userProb,omitempty" type:"Struct"`
	// 所在医疗组
	Group []*QueryUserInfoResponseBodyContentGroup `json:"group,omitempty" xml:"group,omitempty" type:"Repeated"`
	// 所在科室
	Dept []*QueryUserInfoResponseBodyContentDept `json:"dept,omitempty" xml:"dept,omitempty" type:"Repeated"`
}

func (s QueryUserInfoResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponseBodyContent) SetUid(v string) *QueryUserInfoResponseBodyContent {
	s.Uid = &v
	return s
}

func (s *QueryUserInfoResponseBodyContent) SetUserName(v string) *QueryUserInfoResponseBodyContent {
	s.UserName = &v
	return s
}

func (s *QueryUserInfoResponseBodyContent) SetJob(v *QueryUserInfoResponseBodyContentJob) *QueryUserInfoResponseBodyContent {
	s.Job = v
	return s
}

func (s *QueryUserInfoResponseBodyContent) SetJobNum(v string) *QueryUserInfoResponseBodyContent {
	s.JobNum = &v
	return s
}

func (s *QueryUserInfoResponseBodyContent) SetJobStatus(v *QueryUserInfoResponseBodyContentJobStatus) *QueryUserInfoResponseBodyContent {
	s.JobStatus = v
	return s
}

func (s *QueryUserInfoResponseBodyContent) SetUserProb(v *QueryUserInfoResponseBodyContentUserProb) *QueryUserInfoResponseBodyContent {
	s.UserProb = v
	return s
}

func (s *QueryUserInfoResponseBodyContent) SetGroup(v []*QueryUserInfoResponseBodyContentGroup) *QueryUserInfoResponseBodyContent {
	s.Group = v
	return s
}

func (s *QueryUserInfoResponseBodyContent) SetDept(v []*QueryUserInfoResponseBodyContentDept) *QueryUserInfoResponseBodyContent {
	s.Dept = v
	return s
}

type QueryUserInfoResponseBodyContentJob struct {
	// 标签Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 标签类型
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 分类
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 展示名称
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s QueryUserInfoResponseBodyContentJob) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponseBodyContentJob) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponseBodyContentJob) SetCode(v string) *QueryUserInfoResponseBodyContentJob {
	s.Code = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentJob) SetBizType(v string) *QueryUserInfoResponseBodyContentJob {
	s.BizType = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentJob) SetCategory(v string) *QueryUserInfoResponseBodyContentJob {
	s.Category = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentJob) SetDisplayName(v string) *QueryUserInfoResponseBodyContentJob {
	s.DisplayName = &v
	return s
}

type QueryUserInfoResponseBodyContentJobStatus struct {
	// 标签Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 标签类型
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 分类
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 展示名称
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s QueryUserInfoResponseBodyContentJobStatus) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponseBodyContentJobStatus) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponseBodyContentJobStatus) SetCode(v string) *QueryUserInfoResponseBodyContentJobStatus {
	s.Code = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentJobStatus) SetBizType(v string) *QueryUserInfoResponseBodyContentJobStatus {
	s.BizType = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentJobStatus) SetCategory(v string) *QueryUserInfoResponseBodyContentJobStatus {
	s.Category = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentJobStatus) SetDisplayName(v string) *QueryUserInfoResponseBodyContentJobStatus {
	s.DisplayName = &v
	return s
}

type QueryUserInfoResponseBodyContentUserProb struct {
	// 标签Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 标签类型
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 分类
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 展示名称
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s QueryUserInfoResponseBodyContentUserProb) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponseBodyContentUserProb) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponseBodyContentUserProb) SetCode(v string) *QueryUserInfoResponseBodyContentUserProb {
	s.Code = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentUserProb) SetBizType(v string) *QueryUserInfoResponseBodyContentUserProb {
	s.BizType = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentUserProb) SetCategory(v string) *QueryUserInfoResponseBodyContentUserProb {
	s.Category = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentUserProb) SetDisplayName(v string) *QueryUserInfoResponseBodyContentUserProb {
	s.DisplayName = &v
	return s
}

type QueryUserInfoResponseBodyContentGroup struct {
	// 医疗组Id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 医疗组名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 医疗组所在科室Id
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 医疗组所在科室名称
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
}

func (s QueryUserInfoResponseBodyContentGroup) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponseBodyContentGroup) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponseBodyContentGroup) SetId(v int64) *QueryUserInfoResponseBodyContentGroup {
	s.Id = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentGroup) SetName(v string) *QueryUserInfoResponseBodyContentGroup {
	s.Name = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentGroup) SetDeptId(v int64) *QueryUserInfoResponseBodyContentGroup {
	s.DeptId = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentGroup) SetDeptName(v string) *QueryUserInfoResponseBodyContentGroup {
	s.DeptName = &v
	return s
}

type QueryUserInfoResponseBodyContentDept struct {
	// 科室Id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 科室名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryUserInfoResponseBodyContentDept) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponseBodyContentDept) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponseBodyContentDept) SetId(v int64) *QueryUserInfoResponseBodyContentDept {
	s.Id = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentDept) SetName(v string) *QueryUserInfoResponseBodyContentDept {
	s.Name = &v
	return s
}

type QueryUserInfoResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponse) SetHeaders(v map[string]*string) *QueryUserInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryUserInfoResponse) SetBody(v *QueryUserInfoResponseBody) *QueryUserInfoResponse {
	s.Body = v
	return s
}

type QueryAllMemberByDeptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllMemberByDeptHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByDeptHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByDeptHeaders) SetCommonHeaders(v map[string]*string) *QueryAllMemberByDeptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllMemberByDeptHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllMemberByDeptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllMemberByDeptRequest struct {
	// 分页查询页容量
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 分页查询页码
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
}

func (s QueryAllMemberByDeptRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByDeptRequest) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByDeptRequest) SetPageSize(v int32) *QueryAllMemberByDeptRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAllMemberByDeptRequest) SetPageNumber(v int32) *QueryAllMemberByDeptRequest {
	s.PageNumber = &v
	return s
}

type QueryAllMemberByDeptResponseBody struct {
	// 人员列表
	Content []*QueryAllMemberByDeptResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// 总页数
	TotalPages *int32 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
	// 数据总量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 当前页码
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
}

func (s QueryAllMemberByDeptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByDeptResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByDeptResponseBody) SetContent(v []*QueryAllMemberByDeptResponseBodyContent) *QueryAllMemberByDeptResponseBody {
	s.Content = v
	return s
}

func (s *QueryAllMemberByDeptResponseBody) SetTotalPages(v int32) *QueryAllMemberByDeptResponseBody {
	s.TotalPages = &v
	return s
}

func (s *QueryAllMemberByDeptResponseBody) SetTotalCount(v int64) *QueryAllMemberByDeptResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryAllMemberByDeptResponseBody) SetCurrentPage(v int32) *QueryAllMemberByDeptResponseBody {
	s.CurrentPage = &v
	return s
}

type QueryAllMemberByDeptResponseBodyContent struct {
	// 用户Id
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// 用户名称
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// 工号
	JobNum *string `json:"jobNum,omitempty" xml:"jobNum,omitempty"`
}

func (s QueryAllMemberByDeptResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByDeptResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByDeptResponseBodyContent) SetUid(v string) *QueryAllMemberByDeptResponseBodyContent {
	s.Uid = &v
	return s
}

func (s *QueryAllMemberByDeptResponseBodyContent) SetUserName(v string) *QueryAllMemberByDeptResponseBodyContent {
	s.UserName = &v
	return s
}

func (s *QueryAllMemberByDeptResponseBodyContent) SetJobNum(v string) *QueryAllMemberByDeptResponseBodyContent {
	s.JobNum = &v
	return s
}

type QueryAllMemberByDeptResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAllMemberByDeptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllMemberByDeptResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByDeptResponse) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByDeptResponse) SetHeaders(v map[string]*string) *QueryAllMemberByDeptResponse {
	s.Headers = v
	return s
}

func (s *QueryAllMemberByDeptResponse) SetBody(v *QueryAllMemberByDeptResponseBody) *QueryAllMemberByDeptResponse {
	s.Body = v
	return s
}

type QueryAllMemberByGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllMemberByGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByGroupHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByGroupHeaders) SetCommonHeaders(v map[string]*string) *QueryAllMemberByGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllMemberByGroupHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllMemberByGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllMemberByGroupRequest struct {
	// 分页查询分页大小
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 分页查询页码
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
}

func (s QueryAllMemberByGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByGroupRequest) SetPageSize(v int32) *QueryAllMemberByGroupRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAllMemberByGroupRequest) SetPageNumber(v int32) *QueryAllMemberByGroupRequest {
	s.PageNumber = &v
	return s
}

type QueryAllMemberByGroupResponseBody struct {
	// 人员列表
	Content []*QueryAllMemberByGroupResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// 总页数
	TotalPages *int32 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
	// 数据总量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 当前页码
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
}

func (s QueryAllMemberByGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByGroupResponseBody) SetContent(v []*QueryAllMemberByGroupResponseBodyContent) *QueryAllMemberByGroupResponseBody {
	s.Content = v
	return s
}

func (s *QueryAllMemberByGroupResponseBody) SetTotalPages(v int32) *QueryAllMemberByGroupResponseBody {
	s.TotalPages = &v
	return s
}

func (s *QueryAllMemberByGroupResponseBody) SetTotalCount(v int64) *QueryAllMemberByGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryAllMemberByGroupResponseBody) SetCurrentPage(v int32) *QueryAllMemberByGroupResponseBody {
	s.CurrentPage = &v
	return s
}

type QueryAllMemberByGroupResponseBodyContent struct {
	// 用户Id
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// 用户名称
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// 工号
	JobNum *string `json:"jobNum,omitempty" xml:"jobNum,omitempty"`
}

func (s QueryAllMemberByGroupResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByGroupResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByGroupResponseBodyContent) SetUid(v string) *QueryAllMemberByGroupResponseBodyContent {
	s.Uid = &v
	return s
}

func (s *QueryAllMemberByGroupResponseBodyContent) SetUserName(v string) *QueryAllMemberByGroupResponseBodyContent {
	s.UserName = &v
	return s
}

func (s *QueryAllMemberByGroupResponseBodyContent) SetJobNum(v string) *QueryAllMemberByGroupResponseBodyContent {
	s.JobNum = &v
	return s
}

type QueryAllMemberByGroupResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAllMemberByGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllMemberByGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByGroupResponse) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByGroupResponse) SetHeaders(v map[string]*string) *QueryAllMemberByGroupResponse {
	s.Headers = v
	return s
}

func (s *QueryAllMemberByGroupResponse) SetBody(v *QueryAllMemberByGroupResponseBody) *QueryAllMemberByGroupResponse {
	s.Body = v
	return s
}

type QueryUserRolesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserRolesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserRolesHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserRolesHeaders) SetCommonHeaders(v map[string]*string) *QueryUserRolesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserRolesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserRolesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserRolesResponseBody struct {
	// 扩展属性
	Content []*QueryUserRolesResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s QueryUserRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserRolesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserRolesResponseBody) SetContent(v []*QueryUserRolesResponseBodyContent) *QueryUserRolesResponseBody {
	s.Content = v
	return s
}

type QueryUserRolesResponseBodyContent struct {
	// 角色编码
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	// 角色名称
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
}

func (s QueryUserRolesResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryUserRolesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryUserRolesResponseBodyContent) SetRoleCode(v string) *QueryUserRolesResponseBodyContent {
	s.RoleCode = &v
	return s
}

func (s *QueryUserRolesResponseBodyContent) SetRoleName(v string) *QueryUserRolesResponseBodyContent {
	s.RoleName = &v
	return s
}

type QueryUserRolesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUserRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserRolesResponse) GoString() string {
	return s.String()
}

func (s *QueryUserRolesResponse) SetHeaders(v map[string]*string) *QueryUserRolesResponse {
	s.Headers = v
	return s
}

func (s *QueryUserRolesResponse) SetBody(v *QueryUserRolesResponseBody) *QueryUserRolesResponse {
	s.Body = v
	return s
}

type QueryAllGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllGroupHeaders) SetCommonHeaders(v map[string]*string) *QueryAllGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllGroupHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllGroupRequest struct {
	// 分页查询页容量
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 分页查询页码
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
}

func (s QueryAllGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryAllGroupRequest) SetPageSize(v int32) *QueryAllGroupRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAllGroupRequest) SetPageNumber(v int32) *QueryAllGroupRequest {
	s.PageNumber = &v
	return s
}

type QueryAllGroupResponseBody struct {
	// Id of the request
	Content []*QueryAllGroupResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// 总页数
	TotalPages *int64 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
	// 数据总量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 当前页码
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
}

func (s QueryAllGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllGroupResponseBody) SetContent(v []*QueryAllGroupResponseBodyContent) *QueryAllGroupResponseBody {
	s.Content = v
	return s
}

func (s *QueryAllGroupResponseBody) SetTotalPages(v int64) *QueryAllGroupResponseBody {
	s.TotalPages = &v
	return s
}

func (s *QueryAllGroupResponseBody) SetTotalCount(v int64) *QueryAllGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryAllGroupResponseBody) SetCurrentPage(v int64) *QueryAllGroupResponseBody {
	s.CurrentPage = &v
	return s
}

type QueryAllGroupResponseBodyContent struct {
	// 医疗组Id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 医疗组名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 所在科室Id
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s QueryAllGroupResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryAllGroupResponseBodyContent) SetId(v int64) *QueryAllGroupResponseBodyContent {
	s.Id = &v
	return s
}

func (s *QueryAllGroupResponseBodyContent) SetName(v string) *QueryAllGroupResponseBodyContent {
	s.Name = &v
	return s
}

func (s *QueryAllGroupResponseBodyContent) SetDeptId(v int64) *QueryAllGroupResponseBodyContent {
	s.DeptId = &v
	return s
}

type QueryAllGroupResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAllGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupResponse) GoString() string {
	return s.String()
}

func (s *QueryAllGroupResponse) SetHeaders(v map[string]*string) *QueryAllGroupResponse {
	s.Headers = v
	return s
}

func (s *QueryAllGroupResponse) SetBody(v *QueryAllGroupResponseBody) *QueryAllGroupResponse {
	s.Body = v
	return s
}

type QueryAllDoctorsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllDoctorsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDoctorsHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllDoctorsHeaders) SetCommonHeaders(v map[string]*string) *QueryAllDoctorsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllDoctorsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllDoctorsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllDoctorsRequest struct {
	// 分页查询页容量
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 分页查询页码
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
}

func (s QueryAllDoctorsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDoctorsRequest) GoString() string {
	return s.String()
}

func (s *QueryAllDoctorsRequest) SetPageSize(v int32) *QueryAllDoctorsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAllDoctorsRequest) SetPageNum(v int32) *QueryAllDoctorsRequest {
	s.PageNum = &v
	return s
}

type QueryAllDoctorsResponseBody struct {
	// 人员列表
	Content []*QueryAllDoctorsResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// 总页数
	TotalPages *int32 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
	// 数据总量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 当前页码
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
}

func (s QueryAllDoctorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDoctorsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllDoctorsResponseBody) SetContent(v []*QueryAllDoctorsResponseBodyContent) *QueryAllDoctorsResponseBody {
	s.Content = v
	return s
}

func (s *QueryAllDoctorsResponseBody) SetTotalPages(v int32) *QueryAllDoctorsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *QueryAllDoctorsResponseBody) SetTotalCount(v int64) *QueryAllDoctorsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryAllDoctorsResponseBody) SetCurrentPage(v int32) *QueryAllDoctorsResponseBody {
	s.CurrentPage = &v
	return s
}

type QueryAllDoctorsResponseBodyContent struct {
	// 用户Id
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// 用户名称
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// 工号
	JobNum *string `json:"jobNum,omitempty" xml:"jobNum,omitempty"`
}

func (s QueryAllDoctorsResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDoctorsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryAllDoctorsResponseBodyContent) SetUid(v string) *QueryAllDoctorsResponseBodyContent {
	s.Uid = &v
	return s
}

func (s *QueryAllDoctorsResponseBodyContent) SetUserName(v string) *QueryAllDoctorsResponseBodyContent {
	s.UserName = &v
	return s
}

func (s *QueryAllDoctorsResponseBodyContent) SetJobNum(v string) *QueryAllDoctorsResponseBodyContent {
	s.JobNum = &v
	return s
}

type QueryAllDoctorsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAllDoctorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllDoctorsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDoctorsResponse) GoString() string {
	return s.String()
}

func (s *QueryAllDoctorsResponse) SetHeaders(v map[string]*string) *QueryAllDoctorsResponse {
	s.Headers = v
	return s
}

func (s *QueryAllDoctorsResponse) SetBody(v *QueryAllDoctorsResponseBody) *QueryAllDoctorsResponse {
	s.Body = v
	return s
}

type QueryAllGroupsInDeptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllGroupsInDeptHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupsInDeptHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllGroupsInDeptHeaders) SetCommonHeaders(v map[string]*string) *QueryAllGroupsInDeptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllGroupsInDeptHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllGroupsInDeptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllGroupsInDeptRequest struct {
	// 分页查询页容量
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 分页查询页码
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
}

func (s QueryAllGroupsInDeptRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupsInDeptRequest) GoString() string {
	return s.String()
}

func (s *QueryAllGroupsInDeptRequest) SetPageSize(v int32) *QueryAllGroupsInDeptRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAllGroupsInDeptRequest) SetPageNumber(v int32) *QueryAllGroupsInDeptRequest {
	s.PageNumber = &v
	return s
}

type QueryAllGroupsInDeptResponseBody struct {
	// Id of the request
	Content []*QueryAllGroupsInDeptResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// 总页数
	TotalPages *int64 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
	// 数据总量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 当前页码
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
}

func (s QueryAllGroupsInDeptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupsInDeptResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllGroupsInDeptResponseBody) SetContent(v []*QueryAllGroupsInDeptResponseBodyContent) *QueryAllGroupsInDeptResponseBody {
	s.Content = v
	return s
}

func (s *QueryAllGroupsInDeptResponseBody) SetTotalPages(v int64) *QueryAllGroupsInDeptResponseBody {
	s.TotalPages = &v
	return s
}

func (s *QueryAllGroupsInDeptResponseBody) SetTotalCount(v int64) *QueryAllGroupsInDeptResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryAllGroupsInDeptResponseBody) SetCurrentPage(v int64) *QueryAllGroupsInDeptResponseBody {
	s.CurrentPage = &v
	return s
}

type QueryAllGroupsInDeptResponseBodyContent struct {
	// 医疗组Id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 医疗组名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 所在科室Id
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s QueryAllGroupsInDeptResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupsInDeptResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryAllGroupsInDeptResponseBodyContent) SetId(v int64) *QueryAllGroupsInDeptResponseBodyContent {
	s.Id = &v
	return s
}

func (s *QueryAllGroupsInDeptResponseBodyContent) SetName(v string) *QueryAllGroupsInDeptResponseBodyContent {
	s.Name = &v
	return s
}

func (s *QueryAllGroupsInDeptResponseBodyContent) SetDeptId(v int64) *QueryAllGroupsInDeptResponseBodyContent {
	s.DeptId = &v
	return s
}

type QueryAllGroupsInDeptResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAllGroupsInDeptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllGroupsInDeptResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupsInDeptResponse) GoString() string {
	return s.String()
}

func (s *QueryAllGroupsInDeptResponse) SetHeaders(v map[string]*string) *QueryAllGroupsInDeptResponse {
	s.Headers = v
	return s
}

func (s *QueryAllGroupsInDeptResponse) SetBody(v *QueryAllGroupsInDeptResponseBody) *QueryAllGroupsInDeptResponse {
	s.Body = v
	return s
}

type QueryBizOptLogHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryBizOptLogHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryBizOptLogHeaders) GoString() string {
	return s.String()
}

func (s *QueryBizOptLogHeaders) SetCommonHeaders(v map[string]*string) *QueryBizOptLogHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryBizOptLogHeaders) SetXAcsDingtalkAccessToken(v string) *QueryBizOptLogHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryBizOptLogRequest struct {
	// 拉取记录的起始位置，默认从上次拉取的最后位置开始
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 每次拉取的数据量，最大200条
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
}

func (s QueryBizOptLogRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBizOptLogRequest) GoString() string {
	return s.String()
}

func (s *QueryBizOptLogRequest) SetNextToken(v int64) *QueryBizOptLogRequest {
	s.NextToken = &v
	return s
}

func (s *QueryBizOptLogRequest) SetMaxResults(v int32) *QueryBizOptLogRequest {
	s.MaxResults = &v
	return s
}

type QueryBizOptLogResponseBody struct {
	// content
	Content []*QueryBizOptLogResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// 下次拉取数据的起始位置
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryBizOptLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBizOptLogResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBizOptLogResponseBody) SetContent(v []*QueryBizOptLogResponseBodyContent) *QueryBizOptLogResponseBody {
	s.Content = v
	return s
}

func (s *QueryBizOptLogResponseBody) SetNextToken(v int64) *QueryBizOptLogResponseBody {
	s.NextToken = &v
	return s
}

type QueryBizOptLogResponseBodyContent struct {
	// 日志ID
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 数据类型
	DataType *int32 `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// 业务类型
	BizType *int32 `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 操作时间 时间戳
	OptTime *int64 `json:"optTime,omitempty" xml:"optTime,omitempty"`
	// 操作用户code
	OptUserCode *string `json:"optUserCode,omitempty" xml:"optUserCode,omitempty"`
	// 操作用户名称
	OptUserName *string `json:"optUserName,omitempty" xml:"optUserName,omitempty"`
	// 操作者工号
	OptJobNumber *string `json:"optJobNumber,omitempty" xml:"optJobNumber,omitempty"`
	// 操作类型
	OptType *int32 `json:"optType,omitempty" xml:"optType,omitempty"`
	// 操作业务类型
	OptBizType *int32 `json:"optBizType,omitempty" xml:"optBizType,omitempty"`
	// 操作对象code，人员code，或者部门code
	OptObjectCode *string `json:"optObjectCode,omitempty" xml:"optObjectCode,omitempty"`
	// 操作对象人员工号
	OptObjectUserJobNo *string `json:"optObjectUserJobNo,omitempty" xml:"optObjectUserJobNo,omitempty"`
	// 操作对象名称
	OptObjectName *string `json:"optObjectName,omitempty" xml:"optObjectName,omitempty"`
	// 操作是否成功
	OptSuccess *int32 `json:"optSuccess,omitempty" xml:"optSuccess,omitempty"`
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 操作前对象数据快照，json格式
	OptBeforeData *string `json:"optBeforeData,omitempty" xml:"optBeforeData,omitempty"`
	// 操作后对象数据快照，json格式
	OptAfterData *string `json:"optAfterData,omitempty" xml:"optAfterData,omitempty"`
	// 扩展信息，map json格式
	OptExtend *string `json:"optExtend,omitempty" xml:"optExtend,omitempty"`
}

func (s QueryBizOptLogResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryBizOptLogResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryBizOptLogResponseBodyContent) SetId(v int64) *QueryBizOptLogResponseBodyContent {
	s.Id = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetDataType(v int32) *QueryBizOptLogResponseBodyContent {
	s.DataType = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetBizType(v int32) *QueryBizOptLogResponseBodyContent {
	s.BizType = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptTime(v int64) *QueryBizOptLogResponseBodyContent {
	s.OptTime = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptUserCode(v string) *QueryBizOptLogResponseBodyContent {
	s.OptUserCode = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptUserName(v string) *QueryBizOptLogResponseBodyContent {
	s.OptUserName = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptJobNumber(v string) *QueryBizOptLogResponseBodyContent {
	s.OptJobNumber = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptType(v int32) *QueryBizOptLogResponseBodyContent {
	s.OptType = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptBizType(v int32) *QueryBizOptLogResponseBodyContent {
	s.OptBizType = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptObjectCode(v string) *QueryBizOptLogResponseBodyContent {
	s.OptObjectCode = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptObjectUserJobNo(v string) *QueryBizOptLogResponseBodyContent {
	s.OptObjectUserJobNo = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptObjectName(v string) *QueryBizOptLogResponseBodyContent {
	s.OptObjectName = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptSuccess(v int32) *QueryBizOptLogResponseBodyContent {
	s.OptSuccess = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetRemark(v string) *QueryBizOptLogResponseBodyContent {
	s.Remark = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptBeforeData(v string) *QueryBizOptLogResponseBodyContent {
	s.OptBeforeData = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptAfterData(v string) *QueryBizOptLogResponseBodyContent {
	s.OptAfterData = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptExtend(v string) *QueryBizOptLogResponseBodyContent {
	s.OptExtend = &v
	return s
}

type QueryBizOptLogResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryBizOptLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBizOptLogResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBizOptLogResponse) GoString() string {
	return s.String()
}

func (s *QueryBizOptLogResponse) SetHeaders(v map[string]*string) *QueryBizOptLogResponse {
	s.Headers = v
	return s
}

func (s *QueryBizOptLogResponse) SetBody(v *QueryBizOptLogResponseBody) *QueryBizOptLogResponse {
	s.Body = v
	return s
}

type QueryUserExtInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserExtInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserExtInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserExtInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryUserExtInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserExtInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserExtInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserExtInfoResponseBody struct {
	// 扩展属性
	Content []*QueryUserExtInfoResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s QueryUserExtInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserExtInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserExtInfoResponseBody) SetContent(v []*QueryUserExtInfoResponseBodyContent) *QueryUserExtInfoResponseBody {
	s.Content = v
	return s
}

type QueryUserExtInfoResponseBodyContent struct {
	// 扩展属性Key
	UserExtendKey *string `json:"userExtendKey,omitempty" xml:"userExtendKey,omitempty"`
	// 扩展属性值
	UserExtendValue *string `json:"userExtendValue,omitempty" xml:"userExtendValue,omitempty"`
	// 扩展属性描述
	UserExtendDisplayName *string `json:"userExtendDisplayName,omitempty" xml:"userExtendDisplayName,omitempty"`
}

func (s QueryUserExtInfoResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryUserExtInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryUserExtInfoResponseBodyContent) SetUserExtendKey(v string) *QueryUserExtInfoResponseBodyContent {
	s.UserExtendKey = &v
	return s
}

func (s *QueryUserExtInfoResponseBodyContent) SetUserExtendValue(v string) *QueryUserExtInfoResponseBodyContent {
	s.UserExtendValue = &v
	return s
}

func (s *QueryUserExtInfoResponseBodyContent) SetUserExtendDisplayName(v string) *QueryUserExtInfoResponseBodyContent {
	s.UserExtendDisplayName = &v
	return s
}

type QueryUserExtInfoResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUserExtInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserExtInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserExtInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryUserExtInfoResponse) SetHeaders(v map[string]*string) *QueryUserExtInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryUserExtInfoResponse) SetBody(v *QueryUserExtInfoResponseBody) *QueryUserExtInfoResponse {
	s.Body = v
	return s
}

type QueryAllDepartmentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllDepartmentHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentHeaders) SetCommonHeaders(v map[string]*string) *QueryAllDepartmentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllDepartmentHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllDepartmentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllDepartmentRequest struct {
	// 分页查询分页大小
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 分页查询页码
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
}

func (s QueryAllDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentRequest) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentRequest) SetPageSize(v int32) *QueryAllDepartmentRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAllDepartmentRequest) SetPageNumber(v int32) *QueryAllDepartmentRequest {
	s.PageNumber = &v
	return s
}

type QueryAllDepartmentResponseBody struct {
	// 科室列表
	Content []*QueryAllDepartmentResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// 总页数
	TotalPages *int32 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
	// 数据总量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 当前页码
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
}

func (s QueryAllDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentResponseBody) SetContent(v []*QueryAllDepartmentResponseBodyContent) *QueryAllDepartmentResponseBody {
	s.Content = v
	return s
}

func (s *QueryAllDepartmentResponseBody) SetTotalPages(v int32) *QueryAllDepartmentResponseBody {
	s.TotalPages = &v
	return s
}

func (s *QueryAllDepartmentResponseBody) SetTotalCount(v int64) *QueryAllDepartmentResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryAllDepartmentResponseBody) SetCurrentPage(v int32) *QueryAllDepartmentResponseBody {
	s.CurrentPage = &v
	return s
}

type QueryAllDepartmentResponseBodyContent struct {
	// 科室Id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 科室名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryAllDepartmentResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentResponseBodyContent) SetId(v int64) *QueryAllDepartmentResponseBodyContent {
	s.Id = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContent) SetName(v string) *QueryAllDepartmentResponseBodyContent {
	s.Name = &v
	return s
}

type QueryAllDepartmentResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAllDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentResponse) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentResponse) SetHeaders(v map[string]*string) *QueryAllDepartmentResponse {
	s.Headers = v
	return s
}

func (s *QueryAllDepartmentResponse) SetBody(v *QueryAllDepartmentResponseBody) *QueryAllDepartmentResponse {
	s.Body = v
	return s
}

type QueryDepartmentInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDepartmentInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryDepartmentInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryDepartmentInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDepartmentInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDepartmentInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDepartmentInfoResponseBody struct {
	// 科室详情
	Content *QueryDepartmentInfoResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
}

func (s QueryDepartmentInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDepartmentInfoResponseBody) SetContent(v *QueryDepartmentInfoResponseBodyContent) *QueryDepartmentInfoResponseBody {
	s.Content = v
	return s
}

type QueryDepartmentInfoResponseBodyContent struct {
	// 科室Id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 科室名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 科室主任
	Leader *QueryDepartmentInfoResponseBodyContentLeader `json:"leader,omitempty" xml:"leader,omitempty" type:"Struct"`
	// 住院总医师
	ResidentLeader *QueryDepartmentInfoResponseBodyContentResidentLeader `json:"residentLeader,omitempty" xml:"residentLeader,omitempty" type:"Struct"`
}

func (s QueryDepartmentInfoResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryDepartmentInfoResponseBodyContent) SetId(v int64) *QueryDepartmentInfoResponseBodyContent {
	s.Id = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContent) SetName(v string) *QueryDepartmentInfoResponseBodyContent {
	s.Name = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContent) SetLeader(v *QueryDepartmentInfoResponseBodyContentLeader) *QueryDepartmentInfoResponseBodyContent {
	s.Leader = v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContent) SetResidentLeader(v *QueryDepartmentInfoResponseBodyContentResidentLeader) *QueryDepartmentInfoResponseBodyContent {
	s.ResidentLeader = v
	return s
}

type QueryDepartmentInfoResponseBodyContentLeader struct {
	// 姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 人员Id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 工号
	JobNumber *string `json:"jobNumber,omitempty" xml:"jobNumber,omitempty"`
	// 工作标签
	Job *QueryDepartmentInfoResponseBodyContentLeaderJob `json:"job,omitempty" xml:"job,omitempty" type:"Struct"`
}

func (s QueryDepartmentInfoResponseBodyContentLeader) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentInfoResponseBodyContentLeader) GoString() string {
	return s.String()
}

func (s *QueryDepartmentInfoResponseBodyContentLeader) SetName(v string) *QueryDepartmentInfoResponseBodyContentLeader {
	s.Name = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentLeader) SetUserId(v string) *QueryDepartmentInfoResponseBodyContentLeader {
	s.UserId = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentLeader) SetJobNumber(v string) *QueryDepartmentInfoResponseBodyContentLeader {
	s.JobNumber = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentLeader) SetJob(v *QueryDepartmentInfoResponseBodyContentLeaderJob) *QueryDepartmentInfoResponseBodyContentLeader {
	s.Job = v
	return s
}

type QueryDepartmentInfoResponseBodyContentLeaderJob struct {
	// 标签Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 业务类型
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 分类
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 展示名称
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s QueryDepartmentInfoResponseBodyContentLeaderJob) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentInfoResponseBodyContentLeaderJob) GoString() string {
	return s.String()
}

func (s *QueryDepartmentInfoResponseBodyContentLeaderJob) SetCode(v string) *QueryDepartmentInfoResponseBodyContentLeaderJob {
	s.Code = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentLeaderJob) SetBizType(v string) *QueryDepartmentInfoResponseBodyContentLeaderJob {
	s.BizType = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentLeaderJob) SetCategory(v string) *QueryDepartmentInfoResponseBodyContentLeaderJob {
	s.Category = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentLeaderJob) SetDisplayName(v string) *QueryDepartmentInfoResponseBodyContentLeaderJob {
	s.DisplayName = &v
	return s
}

type QueryDepartmentInfoResponseBodyContentResidentLeader struct {
	// 姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 人员Id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 工号
	JobNumber *string `json:"jobNumber,omitempty" xml:"jobNumber,omitempty"`
	// 工作标签
	Job *QueryDepartmentInfoResponseBodyContentResidentLeaderJob `json:"job,omitempty" xml:"job,omitempty" type:"Struct"`
}

func (s QueryDepartmentInfoResponseBodyContentResidentLeader) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentInfoResponseBodyContentResidentLeader) GoString() string {
	return s.String()
}

func (s *QueryDepartmentInfoResponseBodyContentResidentLeader) SetName(v string) *QueryDepartmentInfoResponseBodyContentResidentLeader {
	s.Name = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentResidentLeader) SetUserId(v string) *QueryDepartmentInfoResponseBodyContentResidentLeader {
	s.UserId = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentResidentLeader) SetJobNumber(v string) *QueryDepartmentInfoResponseBodyContentResidentLeader {
	s.JobNumber = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentResidentLeader) SetJob(v *QueryDepartmentInfoResponseBodyContentResidentLeaderJob) *QueryDepartmentInfoResponseBodyContentResidentLeader {
	s.Job = v
	return s
}

type QueryDepartmentInfoResponseBodyContentResidentLeaderJob struct {
	// 标签Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 业务类型
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 分类
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 展示名称
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s QueryDepartmentInfoResponseBodyContentResidentLeaderJob) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentInfoResponseBodyContentResidentLeaderJob) GoString() string {
	return s.String()
}

func (s *QueryDepartmentInfoResponseBodyContentResidentLeaderJob) SetCode(v string) *QueryDepartmentInfoResponseBodyContentResidentLeaderJob {
	s.Code = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentResidentLeaderJob) SetBizType(v string) *QueryDepartmentInfoResponseBodyContentResidentLeaderJob {
	s.BizType = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentResidentLeaderJob) SetCategory(v string) *QueryDepartmentInfoResponseBodyContentResidentLeaderJob {
	s.Category = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentResidentLeaderJob) SetDisplayName(v string) *QueryDepartmentInfoResponseBodyContentResidentLeaderJob {
	s.DisplayName = &v
	return s
}

type QueryDepartmentInfoResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDepartmentInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDepartmentInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryDepartmentInfoResponse) SetHeaders(v map[string]*string) *QueryDepartmentInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryDepartmentInfoResponse) SetBody(v *QueryDepartmentInfoResponseBody) *QueryDepartmentInfoResponse {
	s.Body = v
	return s
}

type QueryGroupInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGroupInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryGroupInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGroupInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGroupInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGroupInfoResponseBody struct {
	// 医疗组详情
	Content *QueryGroupInfoResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
}

func (s QueryGroupInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoResponseBody) SetContent(v *QueryGroupInfoResponseBodyContent) *QueryGroupInfoResponseBody {
	s.Content = v
	return s
}

type QueryGroupInfoResponseBodyContent struct {
	// 医疗组Id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 医疗组名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 科室Id
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 医疗组组长
	Leader *QueryGroupInfoResponseBodyContentLeader `json:"leader,omitempty" xml:"leader,omitempty" type:"Struct"`
	// 有效期开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 有效期结束时间
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
}

func (s QueryGroupInfoResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoResponseBodyContent) SetId(v int64) *QueryGroupInfoResponseBodyContent {
	s.Id = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContent) SetName(v string) *QueryGroupInfoResponseBodyContent {
	s.Name = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContent) SetDeptId(v int64) *QueryGroupInfoResponseBodyContent {
	s.DeptId = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContent) SetLeader(v *QueryGroupInfoResponseBodyContentLeader) *QueryGroupInfoResponseBodyContent {
	s.Leader = v
	return s
}

func (s *QueryGroupInfoResponseBodyContent) SetStartTime(v int64) *QueryGroupInfoResponseBodyContent {
	s.StartTime = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContent) SetEndTime(v int64) *QueryGroupInfoResponseBodyContent {
	s.EndTime = &v
	return s
}

type QueryGroupInfoResponseBodyContentLeader struct {
	// 姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 人员Id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 工号
	JobNumber *string `json:"jobNumber,omitempty" xml:"jobNumber,omitempty"`
	// 工作标签
	Job *QueryGroupInfoResponseBodyContentLeaderJob `json:"job,omitempty" xml:"job,omitempty" type:"Struct"`
}

func (s QueryGroupInfoResponseBodyContentLeader) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoResponseBodyContentLeader) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoResponseBodyContentLeader) SetName(v string) *QueryGroupInfoResponseBodyContentLeader {
	s.Name = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentLeader) SetUserId(v string) *QueryGroupInfoResponseBodyContentLeader {
	s.UserId = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentLeader) SetJobNumber(v string) *QueryGroupInfoResponseBodyContentLeader {
	s.JobNumber = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentLeader) SetJob(v *QueryGroupInfoResponseBodyContentLeaderJob) *QueryGroupInfoResponseBodyContentLeader {
	s.Job = v
	return s
}

type QueryGroupInfoResponseBodyContentLeaderJob struct {
	// 标签Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 业务类型
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 分类
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 展示名称
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s QueryGroupInfoResponseBodyContentLeaderJob) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoResponseBodyContentLeaderJob) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoResponseBodyContentLeaderJob) SetCode(v string) *QueryGroupInfoResponseBodyContentLeaderJob {
	s.Code = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentLeaderJob) SetBizType(v string) *QueryGroupInfoResponseBodyContentLeaderJob {
	s.BizType = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentLeaderJob) SetCategory(v string) *QueryGroupInfoResponseBodyContentLeaderJob {
	s.Category = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentLeaderJob) SetDisplayName(v string) *QueryGroupInfoResponseBodyContentLeaderJob {
	s.DisplayName = &v
	return s
}

type QueryGroupInfoResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryGroupInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryGroupInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoResponse) SetHeaders(v map[string]*string) *QueryGroupInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupInfoResponse) SetBody(v *QueryGroupInfoResponseBody) *QueryGroupInfoResponse {
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

func (client *Client) QueryUserInfo(userId *string) (_result *QueryUserInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserInfoHeaders{}
	_result = &QueryUserInfoResponse{}
	_body, _err := client.QueryUserInfoWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserInfoWithOptions(userId *string, headers *QueryUserInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryUserInfoResponse, _err error) {
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
	_result = &QueryUserInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryUserInfo"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/users/"+tea.StringValue(userId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllMemberByDept(deptId *string, request *QueryAllMemberByDeptRequest) (_result *QueryAllMemberByDeptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllMemberByDeptHeaders{}
	_result = &QueryAllMemberByDeptResponse{}
	_body, _err := client.QueryAllMemberByDeptWithOptions(deptId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllMemberByDeptWithOptions(deptId *string, request *QueryAllMemberByDeptRequest, headers *QueryAllMemberByDeptHeaders, runtime *util.RuntimeOptions) (_result *QueryAllMemberByDeptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
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
	_result = &QueryAllMemberByDeptResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAllMemberByDept"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/departments/"+tea.StringValue(deptId)+"/members"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllMemberByGroup(groupId *string, request *QueryAllMemberByGroupRequest) (_result *QueryAllMemberByGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllMemberByGroupHeaders{}
	_result = &QueryAllMemberByGroupResponse{}
	_body, _err := client.QueryAllMemberByGroupWithOptions(groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllMemberByGroupWithOptions(groupId *string, request *QueryAllMemberByGroupRequest, headers *QueryAllMemberByGroupHeaders, runtime *util.RuntimeOptions) (_result *QueryAllMemberByGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
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
	_result = &QueryAllMemberByGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAllMemberByGroup"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/groups/"+tea.StringValue(groupId)+"/members"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserRoles(userId *string) (_result *QueryUserRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserRolesHeaders{}
	_result = &QueryUserRolesResponse{}
	_body, _err := client.QueryUserRolesWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserRolesWithOptions(userId *string, headers *QueryUserRolesHeaders, runtime *util.RuntimeOptions) (_result *QueryUserRolesResponse, _err error) {
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
	_result = &QueryUserRolesResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryUserRoles"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/users/"+tea.StringValue(userId)+"/roles"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllGroup(request *QueryAllGroupRequest) (_result *QueryAllGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllGroupHeaders{}
	_result = &QueryAllGroupResponse{}
	_body, _err := client.QueryAllGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllGroupWithOptions(request *QueryAllGroupRequest, headers *QueryAllGroupHeaders, runtime *util.RuntimeOptions) (_result *QueryAllGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
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
	_result = &QueryAllGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAllGroup"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/groups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllDoctors(request *QueryAllDoctorsRequest) (_result *QueryAllDoctorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllDoctorsHeaders{}
	_result = &QueryAllDoctorsResponse{}
	_body, _err := client.QueryAllDoctorsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllDoctorsWithOptions(request *QueryAllDoctorsRequest, headers *QueryAllDoctorsHeaders, runtime *util.RuntimeOptions) (_result *QueryAllDoctorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
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
	_result = &QueryAllDoctorsResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAllDoctors"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/doctors"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllGroupsInDept(deptId *string, request *QueryAllGroupsInDeptRequest) (_result *QueryAllGroupsInDeptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllGroupsInDeptHeaders{}
	_result = &QueryAllGroupsInDeptResponse{}
	_body, _err := client.QueryAllGroupsInDeptWithOptions(deptId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllGroupsInDeptWithOptions(deptId *string, request *QueryAllGroupsInDeptRequest, headers *QueryAllGroupsInDeptHeaders, runtime *util.RuntimeOptions) (_result *QueryAllGroupsInDeptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
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
	_result = &QueryAllGroupsInDeptResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAllGroupsInDept"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/departments/"+tea.StringValue(deptId)+"/groups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBizOptLog(request *QueryBizOptLogRequest) (_result *QueryBizOptLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryBizOptLogHeaders{}
	_result = &QueryBizOptLogResponse{}
	_body, _err := client.QueryBizOptLogWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBizOptLogWithOptions(request *QueryBizOptLogRequest, headers *QueryBizOptLogHeaders, runtime *util.RuntimeOptions) (_result *QueryBizOptLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
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
	_result = &QueryBizOptLogResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryBizOptLog"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/bizOptLogs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserExtInfo(userId *string) (_result *QueryUserExtInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserExtInfoHeaders{}
	_result = &QueryUserExtInfoResponse{}
	_body, _err := client.QueryUserExtInfoWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserExtInfoWithOptions(userId *string, headers *QueryUserExtInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryUserExtInfoResponse, _err error) {
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
	_result = &QueryUserExtInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryUserExtInfo"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/users/"+tea.StringValue(userId)+"/extInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllDepartment(request *QueryAllDepartmentRequest) (_result *QueryAllDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllDepartmentHeaders{}
	_result = &QueryAllDepartmentResponse{}
	_body, _err := client.QueryAllDepartmentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllDepartmentWithOptions(request *QueryAllDepartmentRequest, headers *QueryAllDepartmentHeaders, runtime *util.RuntimeOptions) (_result *QueryAllDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
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
	_result = &QueryAllDepartmentResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAllDepartment"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/departments"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDepartmentInfo(deptId *string) (_result *QueryDepartmentInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDepartmentInfoHeaders{}
	_result = &QueryDepartmentInfoResponse{}
	_body, _err := client.QueryDepartmentInfoWithOptions(deptId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDepartmentInfoWithOptions(deptId *string, headers *QueryDepartmentInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryDepartmentInfoResponse, _err error) {
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
	_result = &QueryDepartmentInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryDepartmentInfo"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/departments/"+tea.StringValue(deptId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryGroupInfo(groupId *string) (_result *QueryGroupInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGroupInfoHeaders{}
	_result = &QueryGroupInfoResponse{}
	_body, _err := client.QueryGroupInfoWithOptions(groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryGroupInfoWithOptions(groupId *string, headers *QueryGroupInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryGroupInfoResponse, _err error) {
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
	_result = &QueryGroupInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryGroupInfo"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/groups/"+tea.StringValue(groupId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
