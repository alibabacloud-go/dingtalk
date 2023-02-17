// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package chengfeng_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type OpenAnalyzeDataDTO struct {
	// 部门总人数
	DeptCount *int32 `json:"deptCount,omitempty" xml:"deptCount,omitempty"`
	// 无对齐O人数
	NoAlignObjectiveCount *int32 `json:"noAlignObjectiveCount,omitempty" xml:"noAlignObjectiveCount,omitempty"`
	// 未关联关键行动人数
	NoKeyActionCount *int32 `json:"noKeyActionCount,omitempty" xml:"noKeyActionCount,omitempty"`
	// 目标对齐率
	ObjectiveAlignRate *float64 `json:"objectiveAlignRate,omitempty" xml:"objectiveAlignRate,omitempty"`
	// 目标未设定人数
	ObjectiveNoSetCount *int32 `json:"objectiveNoSetCount,omitempty" xml:"objectiveNoSetCount,omitempty"`
	// 有风险O的人数
	ObjectiveRiskCount *int32 `json:"objectiveRiskCount,omitempty" xml:"objectiveRiskCount,omitempty"`
	// 目标设定率
	ObjectiveSetRate *float64 `json:"objectiveSetRate,omitempty" xml:"objectiveSetRate,omitempty"`
	// 只有1个KR的人数
	OnlyOneKeyResultCount *int32 `json:"onlyOneKeyResultCount,omitempty" xml:"onlyOneKeyResultCount,omitempty"`
	// 只有1个O的人数
	OnlyOneObjectiveCount *int32 `json:"onlyOneObjectiveCount,omitempty" xml:"onlyOneObjectiveCount,omitempty"`
	// 近15天进展更新率
	ProgressUpdateRateLast15Days *float64 `json:"progressUpdateRateLast15Days,omitempty" xml:"progressUpdateRateLast15Days,omitempty"`
	// 近30天进展更新率
	ProgressUpdateRateLast30Days *float64 `json:"progressUpdateRateLast30Days,omitempty" xml:"progressUpdateRateLast30Days,omitempty"`
	// 近7天进展更新率
	ProgressUpdateRateLast7Days *float64 `json:"progressUpdateRateLast7Days,omitempty" xml:"progressUpdateRateLast7Days,omitempty"`
}

func (s OpenAnalyzeDataDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenAnalyzeDataDTO) GoString() string {
	return s.String()
}

func (s *OpenAnalyzeDataDTO) SetDeptCount(v int32) *OpenAnalyzeDataDTO {
	s.DeptCount = &v
	return s
}

func (s *OpenAnalyzeDataDTO) SetNoAlignObjectiveCount(v int32) *OpenAnalyzeDataDTO {
	s.NoAlignObjectiveCount = &v
	return s
}

func (s *OpenAnalyzeDataDTO) SetNoKeyActionCount(v int32) *OpenAnalyzeDataDTO {
	s.NoKeyActionCount = &v
	return s
}

func (s *OpenAnalyzeDataDTO) SetObjectiveAlignRate(v float64) *OpenAnalyzeDataDTO {
	s.ObjectiveAlignRate = &v
	return s
}

func (s *OpenAnalyzeDataDTO) SetObjectiveNoSetCount(v int32) *OpenAnalyzeDataDTO {
	s.ObjectiveNoSetCount = &v
	return s
}

func (s *OpenAnalyzeDataDTO) SetObjectiveRiskCount(v int32) *OpenAnalyzeDataDTO {
	s.ObjectiveRiskCount = &v
	return s
}

func (s *OpenAnalyzeDataDTO) SetObjectiveSetRate(v float64) *OpenAnalyzeDataDTO {
	s.ObjectiveSetRate = &v
	return s
}

func (s *OpenAnalyzeDataDTO) SetOnlyOneKeyResultCount(v int32) *OpenAnalyzeDataDTO {
	s.OnlyOneKeyResultCount = &v
	return s
}

func (s *OpenAnalyzeDataDTO) SetOnlyOneObjectiveCount(v int32) *OpenAnalyzeDataDTO {
	s.OnlyOneObjectiveCount = &v
	return s
}

func (s *OpenAnalyzeDataDTO) SetProgressUpdateRateLast15Days(v float64) *OpenAnalyzeDataDTO {
	s.ProgressUpdateRateLast15Days = &v
	return s
}

func (s *OpenAnalyzeDataDTO) SetProgressUpdateRateLast30Days(v float64) *OpenAnalyzeDataDTO {
	s.ProgressUpdateRateLast30Days = &v
	return s
}

func (s *OpenAnalyzeDataDTO) SetProgressUpdateRateLast7Days(v float64) *OpenAnalyzeDataDTO {
	s.ProgressUpdateRateLast7Days = &v
	return s
}

type OpenKeyResultDTO struct {
	// 主键
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// KR进度
	Progress *int32 `json:"progress,omitempty" xml:"progress,omitempty"`
	// KR的状态:1:正常 3:风险
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// 标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// “@”对象列表
	TitleMentions []*TitleMention `json:"titleMentions,omitempty" xml:"titleMentions,omitempty" type:"Repeated"`
	// KR类型
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s OpenKeyResultDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenKeyResultDTO) GoString() string {
	return s.String()
}

func (s *OpenKeyResultDTO) SetId(v string) *OpenKeyResultDTO {
	s.Id = &v
	return s
}

func (s *OpenKeyResultDTO) SetProgress(v int32) *OpenKeyResultDTO {
	s.Progress = &v
	return s
}

func (s *OpenKeyResultDTO) SetStatus(v int32) *OpenKeyResultDTO {
	s.Status = &v
	return s
}

func (s *OpenKeyResultDTO) SetTitle(v string) *OpenKeyResultDTO {
	s.Title = &v
	return s
}

func (s *OpenKeyResultDTO) SetTitleMentions(v []*TitleMention) *OpenKeyResultDTO {
	s.TitleMentions = v
	return s
}

func (s *OpenKeyResultDTO) SetType(v int32) *OpenKeyResultDTO {
	s.Type = &v
	return s
}

type OpenObjectiveDTO struct {
	// 负责人信息
	Executor *OpenUserDTO `json:"executor,omitempty" xml:"executor,omitempty"`
	// 主键
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// KR列表
	KeyResults []*OpenKeyResultDTO `json:"keyResults,omitempty" xml:"keyResults,omitempty" type:"Repeated"`
	// 周期信息
	Period *OpenPeriodDTO `json:"period,omitempty" xml:"period,omitempty"`
	// 进度
	Progress *int32 `json:"progress,omitempty" xml:"progress,omitempty"`
	// 状态
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// 团队信息列表
	Teams []*OpenTeamDTO `json:"teams,omitempty" xml:"teams,omitempty" type:"Repeated"`
	// 目标标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s OpenObjectiveDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenObjectiveDTO) GoString() string {
	return s.String()
}

func (s *OpenObjectiveDTO) SetExecutor(v *OpenUserDTO) *OpenObjectiveDTO {
	s.Executor = v
	return s
}

func (s *OpenObjectiveDTO) SetId(v string) *OpenObjectiveDTO {
	s.Id = &v
	return s
}

func (s *OpenObjectiveDTO) SetKeyResults(v []*OpenKeyResultDTO) *OpenObjectiveDTO {
	s.KeyResults = v
	return s
}

func (s *OpenObjectiveDTO) SetPeriod(v *OpenPeriodDTO) *OpenObjectiveDTO {
	s.Period = v
	return s
}

func (s *OpenObjectiveDTO) SetProgress(v int32) *OpenObjectiveDTO {
	s.Progress = &v
	return s
}

func (s *OpenObjectiveDTO) SetStatus(v int32) *OpenObjectiveDTO {
	s.Status = &v
	return s
}

func (s *OpenObjectiveDTO) SetTeams(v []*OpenTeamDTO) *OpenObjectiveDTO {
	s.Teams = v
	return s
}

func (s *OpenObjectiveDTO) SetTitle(v string) *OpenObjectiveDTO {
	s.Title = &v
	return s
}

type OpenPeriodDTO struct {
	// 结束日期
	EndDate *int64 `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// 周期id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 周期名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 周期类型
	PeriodBizType *string `json:"periodBizType,omitempty" xml:"periodBizType,omitempty"`
	// 开始日期
	StartDate *int64 `json:"startDate,omitempty" xml:"startDate,omitempty"`
}

func (s OpenPeriodDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenPeriodDTO) GoString() string {
	return s.String()
}

func (s *OpenPeriodDTO) SetEndDate(v int64) *OpenPeriodDTO {
	s.EndDate = &v
	return s
}

func (s *OpenPeriodDTO) SetId(v string) *OpenPeriodDTO {
	s.Id = &v
	return s
}

func (s *OpenPeriodDTO) SetName(v string) *OpenPeriodDTO {
	s.Name = &v
	return s
}

func (s *OpenPeriodDTO) SetPeriodBizType(v string) *OpenPeriodDTO {
	s.PeriodBizType = &v
	return s
}

func (s *OpenPeriodDTO) SetStartDate(v int64) *OpenPeriodDTO {
	s.StartDate = &v
	return s
}

type OpenProgressDTO struct {
	// 创建时间戳
	Created *int64 `json:"created,omitempty" xml:"created,omitempty"`
	// 创建人信息
	Creator *OpenUserDTO `json:"creator,omitempty" xml:"creator,omitempty"`
	// 进展内容
	HtmlContent *string `json:"htmlContent,omitempty" xml:"htmlContent,omitempty"`
	// 主键
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 更新人信息
	Modifier *OpenUserDTO `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// 更新时间戳
	Updated *int64 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s OpenProgressDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenProgressDTO) GoString() string {
	return s.String()
}

func (s *OpenProgressDTO) SetCreated(v int64) *OpenProgressDTO {
	s.Created = &v
	return s
}

func (s *OpenProgressDTO) SetCreator(v *OpenUserDTO) *OpenProgressDTO {
	s.Creator = v
	return s
}

func (s *OpenProgressDTO) SetHtmlContent(v string) *OpenProgressDTO {
	s.HtmlContent = &v
	return s
}

func (s *OpenProgressDTO) SetId(v string) *OpenProgressDTO {
	s.Id = &v
	return s
}

func (s *OpenProgressDTO) SetModifier(v *OpenUserDTO) *OpenProgressDTO {
	s.Modifier = v
	return s
}

func (s *OpenProgressDTO) SetUpdated(v int64) *OpenProgressDTO {
	s.Updated = &v
	return s
}

type OpenTeamDTO struct {
	// 部门id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 部门名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 钉钉对应部门编号
	OpenId *string `json:"openId,omitempty" xml:"openId,omitempty"`
}

func (s OpenTeamDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenTeamDTO) GoString() string {
	return s.String()
}

func (s *OpenTeamDTO) SetId(v string) *OpenTeamDTO {
	s.Id = &v
	return s
}

func (s *OpenTeamDTO) SetName(v string) *OpenTeamDTO {
	s.Name = &v
	return s
}

func (s *OpenTeamDTO) SetOpenId(v string) *OpenTeamDTO {
	s.OpenId = &v
	return s
}

type OpenUserDTO struct {
	// 用户id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 用户名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 钉钉用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s OpenUserDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenUserDTO) GoString() string {
	return s.String()
}

func (s *OpenUserDTO) SetId(v string) *OpenUserDTO {
	s.Id = &v
	return s
}

func (s *OpenUserDTO) SetName(v string) *OpenUserDTO {
	s.Name = &v
	return s
}

func (s *OpenUserDTO) SetUserId(v string) *OpenUserDTO {
	s.UserId = &v
	return s
}

type TitleMention struct {
	// 结束位置
	Length *int32 `json:"length,omitempty" xml:"length,omitempty"`
	// 开始位置
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// “@人员”对象信息
	User *OpenUserDTO `json:"user,omitempty" xml:"user,omitempty"`
}

func (s TitleMention) String() string {
	return tea.Prettify(s)
}

func (s TitleMention) GoString() string {
	return s.String()
}

func (s *TitleMention) SetLength(v int32) *TitleMention {
	s.Length = &v
	return s
}

func (s *TitleMention) SetOffset(v int32) *TitleMention {
	s.Offset = &v
	return s
}

func (s *TitleMention) SetUser(v *OpenUserDTO) *TitleMention {
	s.User = v
	return s
}

type GetAnalyzeDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAnalyzeDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAnalyzeDataHeaders) GoString() string {
	return s.String()
}

func (s *GetAnalyzeDataHeaders) SetCommonHeaders(v map[string]*string) *GetAnalyzeDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAnalyzeDataHeaders) SetXAcsDingtalkAccessToken(v string) *GetAnalyzeDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAnalyzeDataRequest struct {
	// 周期ID列表
	PeriodIds []*string `json:"periodIds,omitempty" xml:"periodIds,omitempty" type:"Repeated"`
	// 部门编号(钉钉部门号)
	DeptId *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s GetAnalyzeDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAnalyzeDataRequest) GoString() string {
	return s.String()
}

func (s *GetAnalyzeDataRequest) SetPeriodIds(v []*string) *GetAnalyzeDataRequest {
	s.PeriodIds = v
	return s
}

func (s *GetAnalyzeDataRequest) SetDeptId(v string) *GetAnalyzeDataRequest {
	s.DeptId = &v
	return s
}

type GetAnalyzeDataResponseBody struct {
	Content *OpenAnalyzeDataDTO `json:"content,omitempty" xml:"content,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetAnalyzeDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAnalyzeDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetAnalyzeDataResponseBody) SetContent(v *OpenAnalyzeDataDTO) *GetAnalyzeDataResponseBody {
	s.Content = v
	return s
}

func (s *GetAnalyzeDataResponseBody) SetRequestId(v string) *GetAnalyzeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAnalyzeDataResponseBody) SetSuccess(v bool) *GetAnalyzeDataResponseBody {
	s.Success = &v
	return s
}

type GetAnalyzeDataResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAnalyzeDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAnalyzeDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAnalyzeDataResponse) GoString() string {
	return s.String()
}

func (s *GetAnalyzeDataResponse) SetHeaders(v map[string]*string) *GetAnalyzeDataResponse {
	s.Headers = v
	return s
}

func (s *GetAnalyzeDataResponse) SetBody(v *GetAnalyzeDataResponseBody) *GetAnalyzeDataResponse {
	s.Body = v
	return s
}

type GetEmployeeInfoByWorkNoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetEmployeeInfoByWorkNoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetEmployeeInfoByWorkNoHeaders) GoString() string {
	return s.String()
}

func (s *GetEmployeeInfoByWorkNoHeaders) SetCommonHeaders(v map[string]*string) *GetEmployeeInfoByWorkNoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetEmployeeInfoByWorkNoHeaders) SetXAcsDingtalkAccessToken(v string) *GetEmployeeInfoByWorkNoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetEmployeeInfoByWorkNoRequest struct {
	// 工号
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s GetEmployeeInfoByWorkNoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEmployeeInfoByWorkNoRequest) GoString() string {
	return s.String()
}

func (s *GetEmployeeInfoByWorkNoRequest) SetWorkNo(v string) *GetEmployeeInfoByWorkNoRequest {
	s.WorkNo = &v
	return s
}

type GetEmployeeInfoByWorkNoResponseBody struct {
	// 请求返回数据对象
	Content *GetEmployeeInfoByWorkNoResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	// 接口请求成功标识,成功为true,失败为false
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetEmployeeInfoByWorkNoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEmployeeInfoByWorkNoResponseBody) GoString() string {
	return s.String()
}

func (s *GetEmployeeInfoByWorkNoResponseBody) SetContent(v *GetEmployeeInfoByWorkNoResponseBodyContent) *GetEmployeeInfoByWorkNoResponseBody {
	s.Content = v
	return s
}

func (s *GetEmployeeInfoByWorkNoResponseBody) SetSuccess(v bool) *GetEmployeeInfoByWorkNoResponseBody {
	s.Success = &v
	return s
}

type GetEmployeeInfoByWorkNoResponseBodyContent struct {
	// 员工姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 员工工号
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s GetEmployeeInfoByWorkNoResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetEmployeeInfoByWorkNoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetEmployeeInfoByWorkNoResponseBodyContent) SetName(v string) *GetEmployeeInfoByWorkNoResponseBodyContent {
	s.Name = &v
	return s
}

func (s *GetEmployeeInfoByWorkNoResponseBodyContent) SetWorkNo(v string) *GetEmployeeInfoByWorkNoResponseBodyContent {
	s.WorkNo = &v
	return s
}

type GetEmployeeInfoByWorkNoResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEmployeeInfoByWorkNoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEmployeeInfoByWorkNoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEmployeeInfoByWorkNoResponse) GoString() string {
	return s.String()
}

func (s *GetEmployeeInfoByWorkNoResponse) SetHeaders(v map[string]*string) *GetEmployeeInfoByWorkNoResponse {
	s.Headers = v
	return s
}

func (s *GetEmployeeInfoByWorkNoResponse) SetBody(v *GetEmployeeInfoByWorkNoResponseBody) *GetEmployeeInfoByWorkNoResponse {
	s.Body = v
	return s
}

type GetUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserHeaders) GoString() string {
	return s.String()
}

func (s *GetUserHeaders) SetCommonHeaders(v map[string]*string) *GetUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserRequest struct {
	// OKR系统内部用户id
	OkrUserId *string `json:"okrUserId,omitempty" xml:"okrUserId,omitempty"`
	// 钉钉UserId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) SetOkrUserId(v string) *GetUserRequest {
	s.OkrUserId = &v
	return s
}

func (s *GetUserRequest) SetUserId(v string) *GetUserRequest {
	s.UserId = &v
	return s
}

type GetUserResponseBody struct {
	Content *OpenUserDTO `json:"content,omitempty" xml:"content,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetContent(v *OpenUserDTO) *GetUserResponseBody {
	s.Content = v
	return s
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResponseBody) SetSuccess(v bool) *GetUserResponseBody {
	s.Success = &v
	return s
}

type GetUserResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type ListAnalyzePeriodsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListAnalyzePeriodsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListAnalyzePeriodsHeaders) GoString() string {
	return s.String()
}

func (s *ListAnalyzePeriodsHeaders) SetCommonHeaders(v map[string]*string) *ListAnalyzePeriodsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAnalyzePeriodsHeaders) SetXAcsDingtalkAccessToken(v string) *ListAnalyzePeriodsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListAnalyzePeriodsResponseBody struct {
	Content []*OpenPeriodDTO `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListAnalyzePeriodsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAnalyzePeriodsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAnalyzePeriodsResponseBody) SetContent(v []*OpenPeriodDTO) *ListAnalyzePeriodsResponseBody {
	s.Content = v
	return s
}

func (s *ListAnalyzePeriodsResponseBody) SetRequestId(v string) *ListAnalyzePeriodsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAnalyzePeriodsResponseBody) SetSuccess(v bool) *ListAnalyzePeriodsResponseBody {
	s.Success = &v
	return s
}

type ListAnalyzePeriodsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAnalyzePeriodsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAnalyzePeriodsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAnalyzePeriodsResponse) GoString() string {
	return s.String()
}

func (s *ListAnalyzePeriodsResponse) SetHeaders(v map[string]*string) *ListAnalyzePeriodsResponse {
	s.Headers = v
	return s
}

func (s *ListAnalyzePeriodsResponse) SetBody(v *ListAnalyzePeriodsResponseBody) *ListAnalyzePeriodsResponse {
	s.Body = v
	return s
}

type ListObjectiveByIdsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListObjectiveByIdsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListObjectiveByIdsHeaders) GoString() string {
	return s.String()
}

func (s *ListObjectiveByIdsHeaders) SetCommonHeaders(v map[string]*string) *ListObjectiveByIdsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListObjectiveByIdsHeaders) SetXAcsDingtalkAccessToken(v string) *ListObjectiveByIdsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListObjectiveByIdsRequest struct {
	// 目标ID列表
	ObjectiveIds []*string `json:"objectiveIds,omitempty" xml:"objectiveIds,omitempty" type:"Repeated"`
}

func (s ListObjectiveByIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListObjectiveByIdsRequest) GoString() string {
	return s.String()
}

func (s *ListObjectiveByIdsRequest) SetObjectiveIds(v []*string) *ListObjectiveByIdsRequest {
	s.ObjectiveIds = v
	return s
}

type ListObjectiveByIdsResponseBody struct {
	Content []*OpenObjectiveDTO `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListObjectiveByIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListObjectiveByIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListObjectiveByIdsResponseBody) SetContent(v []*OpenObjectiveDTO) *ListObjectiveByIdsResponseBody {
	s.Content = v
	return s
}

func (s *ListObjectiveByIdsResponseBody) SetRequestId(v string) *ListObjectiveByIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListObjectiveByIdsResponseBody) SetSuccess(v bool) *ListObjectiveByIdsResponseBody {
	s.Success = &v
	return s
}

type ListObjectiveByIdsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListObjectiveByIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListObjectiveByIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListObjectiveByIdsResponse) GoString() string {
	return s.String()
}

func (s *ListObjectiveByIdsResponse) SetHeaders(v map[string]*string) *ListObjectiveByIdsResponse {
	s.Headers = v
	return s
}

func (s *ListObjectiveByIdsResponse) SetBody(v *ListObjectiveByIdsResponseBody) *ListObjectiveByIdsResponse {
	s.Body = v
	return s
}

type ListObjectiveByUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListObjectiveByUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListObjectiveByUserHeaders) GoString() string {
	return s.String()
}

func (s *ListObjectiveByUserHeaders) SetCommonHeaders(v map[string]*string) *ListObjectiveByUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListObjectiveByUserHeaders) SetXAcsDingtalkAccessToken(v string) *ListObjectiveByUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListObjectiveByUserRequest struct {
	// 页数
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 页大小
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListObjectiveByUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListObjectiveByUserRequest) GoString() string {
	return s.String()
}

func (s *ListObjectiveByUserRequest) SetPageNumber(v int32) *ListObjectiveByUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListObjectiveByUserRequest) SetPageSize(v int32) *ListObjectiveByUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListObjectiveByUserRequest) SetUserId(v string) *ListObjectiveByUserRequest {
	s.UserId = &v
	return s
}

type ListObjectiveByUserResponseBody struct {
	// 请求返回数据对象
	Content *ListObjectiveByUserResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 接口请求是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListObjectiveByUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListObjectiveByUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListObjectiveByUserResponseBody) SetContent(v *ListObjectiveByUserResponseBodyContent) *ListObjectiveByUserResponseBody {
	s.Content = v
	return s
}

func (s *ListObjectiveByUserResponseBody) SetRequestId(v string) *ListObjectiveByUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListObjectiveByUserResponseBody) SetSuccess(v bool) *ListObjectiveByUserResponseBody {
	s.Success = &v
	return s
}

type ListObjectiveByUserResponseBodyContent struct {
	// 总数
	Count      *int32              `json:"count,omitempty" xml:"count,omitempty"`
	Objectives []*OpenObjectiveDTO `json:"objectives,omitempty" xml:"objectives,omitempty" type:"Repeated"`
}

func (s ListObjectiveByUserResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s ListObjectiveByUserResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListObjectiveByUserResponseBodyContent) SetCount(v int32) *ListObjectiveByUserResponseBodyContent {
	s.Count = &v
	return s
}

func (s *ListObjectiveByUserResponseBodyContent) SetObjectives(v []*OpenObjectiveDTO) *ListObjectiveByUserResponseBodyContent {
	s.Objectives = v
	return s
}

type ListObjectiveByUserResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListObjectiveByUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListObjectiveByUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListObjectiveByUserResponse) GoString() string {
	return s.String()
}

func (s *ListObjectiveByUserResponse) SetHeaders(v map[string]*string) *ListObjectiveByUserResponse {
	s.Headers = v
	return s
}

func (s *ListObjectiveByUserResponse) SetBody(v *ListObjectiveByUserResponseBody) *ListObjectiveByUserResponse {
	s.Body = v
	return s
}

type ListProgressByIdsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListProgressByIdsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListProgressByIdsHeaders) GoString() string {
	return s.String()
}

func (s *ListProgressByIdsHeaders) SetCommonHeaders(v map[string]*string) *ListProgressByIdsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListProgressByIdsHeaders) SetXAcsDingtalkAccessToken(v string) *ListProgressByIdsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListProgressByIdsRequest struct {
	// 进展ID列表
	ProgressIds []*string `json:"progressIds,omitempty" xml:"progressIds,omitempty" type:"Repeated"`
}

func (s ListProgressByIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProgressByIdsRequest) GoString() string {
	return s.String()
}

func (s *ListProgressByIdsRequest) SetProgressIds(v []*string) *ListProgressByIdsRequest {
	s.ProgressIds = v
	return s
}

type ListProgressByIdsResponseBody struct {
	Content []*OpenProgressDTO `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListProgressByIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProgressByIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProgressByIdsResponseBody) SetContent(v []*OpenProgressDTO) *ListProgressByIdsResponseBody {
	s.Content = v
	return s
}

func (s *ListProgressByIdsResponseBody) SetRequestId(v string) *ListProgressByIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProgressByIdsResponseBody) SetSuccess(v bool) *ListProgressByIdsResponseBody {
	s.Success = &v
	return s
}

type ListProgressByIdsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProgressByIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProgressByIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProgressByIdsResponse) GoString() string {
	return s.String()
}

func (s *ListProgressByIdsResponse) SetHeaders(v map[string]*string) *ListProgressByIdsResponse {
	s.Headers = v
	return s
}

func (s *ListProgressByIdsResponse) SetBody(v *ListProgressByIdsResponseBody) *ListProgressByIdsResponse {
	s.Body = v
	return s
}

type PageListObjectiveProgressHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PageListObjectiveProgressHeaders) String() string {
	return tea.Prettify(s)
}

func (s PageListObjectiveProgressHeaders) GoString() string {
	return s.String()
}

func (s *PageListObjectiveProgressHeaders) SetCommonHeaders(v map[string]*string) *PageListObjectiveProgressHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PageListObjectiveProgressHeaders) SetXAcsDingtalkAccessToken(v string) *PageListObjectiveProgressHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PageListObjectiveProgressRequest struct {
	// 目标id
	ObjectiveId *string `json:"objectiveId,omitempty" xml:"objectiveId,omitempty"`
	// 页数
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 页大小
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s PageListObjectiveProgressRequest) String() string {
	return tea.Prettify(s)
}

func (s PageListObjectiveProgressRequest) GoString() string {
	return s.String()
}

func (s *PageListObjectiveProgressRequest) SetObjectiveId(v string) *PageListObjectiveProgressRequest {
	s.ObjectiveId = &v
	return s
}

func (s *PageListObjectiveProgressRequest) SetPageNumber(v int32) *PageListObjectiveProgressRequest {
	s.PageNumber = &v
	return s
}

func (s *PageListObjectiveProgressRequest) SetPageSize(v int32) *PageListObjectiveProgressRequest {
	s.PageSize = &v
	return s
}

type PageListObjectiveProgressResponseBody struct {
	Content *PageListObjectiveProgressResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PageListObjectiveProgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageListObjectiveProgressResponseBody) GoString() string {
	return s.String()
}

func (s *PageListObjectiveProgressResponseBody) SetContent(v *PageListObjectiveProgressResponseBodyContent) *PageListObjectiveProgressResponseBody {
	s.Content = v
	return s
}

func (s *PageListObjectiveProgressResponseBody) SetRequestId(v string) *PageListObjectiveProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageListObjectiveProgressResponseBody) SetSuccess(v bool) *PageListObjectiveProgressResponseBody {
	s.Success = &v
	return s
}

type PageListObjectiveProgressResponseBodyContent struct {
	Count        *int32             `json:"count,omitempty" xml:"count,omitempty"`
	ProgressList []*OpenProgressDTO `json:"progressList,omitempty" xml:"progressList,omitempty" type:"Repeated"`
}

func (s PageListObjectiveProgressResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s PageListObjectiveProgressResponseBodyContent) GoString() string {
	return s.String()
}

func (s *PageListObjectiveProgressResponseBodyContent) SetCount(v int32) *PageListObjectiveProgressResponseBodyContent {
	s.Count = &v
	return s
}

func (s *PageListObjectiveProgressResponseBodyContent) SetProgressList(v []*OpenProgressDTO) *PageListObjectiveProgressResponseBodyContent {
	s.ProgressList = v
	return s
}

type PageListObjectiveProgressResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PageListObjectiveProgressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PageListObjectiveProgressResponse) String() string {
	return tea.Prettify(s)
}

func (s PageListObjectiveProgressResponse) GoString() string {
	return s.String()
}

func (s *PageListObjectiveProgressResponse) SetHeaders(v map[string]*string) *PageListObjectiveProgressResponse {
	s.Headers = v
	return s
}

func (s *PageListObjectiveProgressResponse) SetBody(v *PageListObjectiveProgressResponseBody) *PageListObjectiveProgressResponse {
	s.Body = v
	return s
}

type TransferUserObjectiveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TransferUserObjectiveHeaders) String() string {
	return tea.Prettify(s)
}

func (s TransferUserObjectiveHeaders) GoString() string {
	return s.String()
}

func (s *TransferUserObjectiveHeaders) SetCommonHeaders(v map[string]*string) *TransferUserObjectiveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TransferUserObjectiveHeaders) SetXAcsDingtalkAccessToken(v string) *TransferUserObjectiveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TransferUserObjectiveRequest struct {
	// 目标ID
	ObjectiveId *string `json:"objectiveId,omitempty" xml:"objectiveId,omitempty"`
	// 目标钉钉userId
	TargetUserId *string `json:"targetUserId,omitempty" xml:"targetUserId,omitempty"`
}

func (s TransferUserObjectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferUserObjectiveRequest) GoString() string {
	return s.String()
}

func (s *TransferUserObjectiveRequest) SetObjectiveId(v string) *TransferUserObjectiveRequest {
	s.ObjectiveId = &v
	return s
}

func (s *TransferUserObjectiveRequest) SetTargetUserId(v string) *TransferUserObjectiveRequest {
	s.TargetUserId = &v
	return s
}

type TransferUserObjectiveResponseBody struct {
	// 转交是否成功
	Content *bool `json:"content,omitempty" xml:"content,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 返回结果
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TransferUserObjectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferUserObjectiveResponseBody) GoString() string {
	return s.String()
}

func (s *TransferUserObjectiveResponseBody) SetContent(v bool) *TransferUserObjectiveResponseBody {
	s.Content = &v
	return s
}

func (s *TransferUserObjectiveResponseBody) SetRequestId(v string) *TransferUserObjectiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferUserObjectiveResponseBody) SetSuccess(v bool) *TransferUserObjectiveResponseBody {
	s.Success = &v
	return s
}

type TransferUserObjectiveResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TransferUserObjectiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransferUserObjectiveResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferUserObjectiveResponse) GoString() string {
	return s.String()
}

func (s *TransferUserObjectiveResponse) SetHeaders(v map[string]*string) *TransferUserObjectiveResponse {
	s.Headers = v
	return s
}

func (s *TransferUserObjectiveResponse) SetBody(v *TransferUserObjectiveResponseBody) *TransferUserObjectiveResponse {
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

func (client *Client) GetAnalyzeData(request *GetAnalyzeDataRequest) (_result *GetAnalyzeDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAnalyzeDataHeaders{}
	_result = &GetAnalyzeDataResponse{}
	_body, _err := client.GetAnalyzeDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAnalyzeDataWithOptions(request *GetAnalyzeDataRequest, headers *GetAnalyzeDataHeaders, runtime *util.RuntimeOptions) (_result *GetAnalyzeDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PeriodIds)) {
		body["periodIds"] = request.PeriodIds
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
	_result = &GetAnalyzeDataResponse{}
	_body, _err := client.DoROARequest(tea.String("GetAnalyzeData"), tea.String("chengfeng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/chengfeng/okr/analyses/datas/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEmployeeInfoByWorkNo(request *GetEmployeeInfoByWorkNoRequest) (_result *GetEmployeeInfoByWorkNoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetEmployeeInfoByWorkNoHeaders{}
	_result = &GetEmployeeInfoByWorkNoResponse{}
	_body, _err := client.GetEmployeeInfoByWorkNoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEmployeeInfoByWorkNoWithOptions(request *GetEmployeeInfoByWorkNoRequest, headers *GetEmployeeInfoByWorkNoHeaders, runtime *util.RuntimeOptions) (_result *GetEmployeeInfoByWorkNoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkNo)) {
		query["workNo"] = request.WorkNo
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
	_result = &GetEmployeeInfoByWorkNoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetEmployeeInfoByWorkNo"), tea.String("chengfeng_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/chengfeng/workNumbers/employees"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserHeaders{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserWithOptions(request *GetUserRequest, headers *GetUserHeaders, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OkrUserId)) {
		query["okrUserId"] = request.OkrUserId
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
	_result = &GetUserResponse{}
	_body, _err := client.DoROARequest(tea.String("GetUser"), tea.String("chengfeng_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/chengfeng/okr/users"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAnalyzePeriods() (_result *ListAnalyzePeriodsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListAnalyzePeriodsHeaders{}
	_result = &ListAnalyzePeriodsResponse{}
	_body, _err := client.ListAnalyzePeriodsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAnalyzePeriodsWithOptions(headers *ListAnalyzePeriodsHeaders, runtime *util.RuntimeOptions) (_result *ListAnalyzePeriodsResponse, _err error) {
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
	_result = &ListAnalyzePeriodsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListAnalyzePeriods"), tea.String("chengfeng_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/chengfeng/okr/analyses/periods"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListObjectiveByIds(request *ListObjectiveByIdsRequest) (_result *ListObjectiveByIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListObjectiveByIdsHeaders{}
	_result = &ListObjectiveByIdsResponse{}
	_body, _err := client.ListObjectiveByIdsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListObjectiveByIdsWithOptions(request *ListObjectiveByIdsRequest, headers *ListObjectiveByIdsHeaders, runtime *util.RuntimeOptions) (_result *ListObjectiveByIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ObjectiveIds)) {
		body["objectiveIds"] = request.ObjectiveIds
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
	_result = &ListObjectiveByIdsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListObjectiveByIds"), tea.String("chengfeng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/chengfeng/okr/objectives/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListObjectiveByUser(request *ListObjectiveByUserRequest) (_result *ListObjectiveByUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListObjectiveByUserHeaders{}
	_result = &ListObjectiveByUserResponse{}
	_body, _err := client.ListObjectiveByUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListObjectiveByUserWithOptions(request *ListObjectiveByUserRequest, headers *ListObjectiveByUserHeaders, runtime *util.RuntimeOptions) (_result *ListObjectiveByUserResponse, _err error) {
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
	_result = &ListObjectiveByUserResponse{}
	_body, _err := client.DoROARequest(tea.String("ListObjectiveByUser"), tea.String("chengfeng_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/chengfeng/okr/users/objectives"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProgressByIds(request *ListProgressByIdsRequest) (_result *ListProgressByIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListProgressByIdsHeaders{}
	_result = &ListProgressByIdsResponse{}
	_body, _err := client.ListProgressByIdsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProgressByIdsWithOptions(request *ListProgressByIdsRequest, headers *ListProgressByIdsHeaders, runtime *util.RuntimeOptions) (_result *ListProgressByIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProgressIds)) {
		body["progressIds"] = request.ProgressIds
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
	_result = &ListProgressByIdsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListProgressByIds"), tea.String("chengfeng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/chengfeng/okr/objectives/progresses/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PageListObjectiveProgress(request *PageListObjectiveProgressRequest) (_result *PageListObjectiveProgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PageListObjectiveProgressHeaders{}
	_result = &PageListObjectiveProgressResponse{}
	_body, _err := client.PageListObjectiveProgressWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PageListObjectiveProgressWithOptions(request *PageListObjectiveProgressRequest, headers *PageListObjectiveProgressHeaders, runtime *util.RuntimeOptions) (_result *PageListObjectiveProgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ObjectiveId)) {
		query["objectiveId"] = request.ObjectiveId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
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
	_result = &PageListObjectiveProgressResponse{}
	_body, _err := client.DoROARequest(tea.String("PageListObjectiveProgress"), tea.String("chengfeng_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/chengfeng/okr/objectives/progresses/records"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferUserObjective(request *TransferUserObjectiveRequest) (_result *TransferUserObjectiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TransferUserObjectiveHeaders{}
	_result = &TransferUserObjectiveResponse{}
	_body, _err := client.TransferUserObjectiveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferUserObjectiveWithOptions(request *TransferUserObjectiveRequest, headers *TransferUserObjectiveHeaders, runtime *util.RuntimeOptions) (_result *TransferUserObjectiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ObjectiveId)) {
		query["objectiveId"] = request.ObjectiveId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetUserId)) {
		query["targetUserId"] = request.TargetUserId
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
	_result = &TransferUserObjectiveResponse{}
	_body, _err := client.DoROARequest(tea.String("TransferUserObjective"), tea.String("chengfeng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/chengfeng/okr/objectives/transfer"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
