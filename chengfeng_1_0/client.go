// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package chengfeng_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CfEmploymentRecordResp struct {
	DeptCode        *string `json:"deptCode,omitempty" xml:"deptCode,omitempty"`
	DeptName        *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	EmployeeStatus  *string `json:"employeeStatus,omitempty" xml:"employeeStatus,omitempty"`
	EndDate         *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	IsLatestRecord  *bool   `json:"isLatestRecord,omitempty" xml:"isLatestRecord,omitempty"`
	JobLevelName    *string `json:"jobLevelName,omitempty" xml:"jobLevelName,omitempty"`
	JobPositionCode *string `json:"jobPositionCode,omitempty" xml:"jobPositionCode,omitempty"`
	JobPositionName *string `json:"jobPositionName,omitempty" xml:"jobPositionName,omitempty"`
	JobPostCode     *string `json:"jobPostCode,omitempty" xml:"jobPostCode,omitempty"`
	JobPostName     *string `json:"jobPostName,omitempty" xml:"jobPostName,omitempty"`
	ServiceStatus   *string `json:"serviceStatus,omitempty" xml:"serviceStatus,omitempty"`
	ServiceType     *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	StartDate       *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	WorkNumbers     *string `json:"workNumbers,omitempty" xml:"workNumbers,omitempty"`
}

func (s CfEmploymentRecordResp) String() string {
	return tea.Prettify(s)
}

func (s CfEmploymentRecordResp) GoString() string {
	return s.String()
}

func (s *CfEmploymentRecordResp) SetDeptCode(v string) *CfEmploymentRecordResp {
	s.DeptCode = &v
	return s
}

func (s *CfEmploymentRecordResp) SetDeptName(v string) *CfEmploymentRecordResp {
	s.DeptName = &v
	return s
}

func (s *CfEmploymentRecordResp) SetEmployeeStatus(v string) *CfEmploymentRecordResp {
	s.EmployeeStatus = &v
	return s
}

func (s *CfEmploymentRecordResp) SetEndDate(v string) *CfEmploymentRecordResp {
	s.EndDate = &v
	return s
}

func (s *CfEmploymentRecordResp) SetIsLatestRecord(v bool) *CfEmploymentRecordResp {
	s.IsLatestRecord = &v
	return s
}

func (s *CfEmploymentRecordResp) SetJobLevelName(v string) *CfEmploymentRecordResp {
	s.JobLevelName = &v
	return s
}

func (s *CfEmploymentRecordResp) SetJobPositionCode(v string) *CfEmploymentRecordResp {
	s.JobPositionCode = &v
	return s
}

func (s *CfEmploymentRecordResp) SetJobPositionName(v string) *CfEmploymentRecordResp {
	s.JobPositionName = &v
	return s
}

func (s *CfEmploymentRecordResp) SetJobPostCode(v string) *CfEmploymentRecordResp {
	s.JobPostCode = &v
	return s
}

func (s *CfEmploymentRecordResp) SetJobPostName(v string) *CfEmploymentRecordResp {
	s.JobPostName = &v
	return s
}

func (s *CfEmploymentRecordResp) SetServiceStatus(v string) *CfEmploymentRecordResp {
	s.ServiceStatus = &v
	return s
}

func (s *CfEmploymentRecordResp) SetServiceType(v string) *CfEmploymentRecordResp {
	s.ServiceType = &v
	return s
}

func (s *CfEmploymentRecordResp) SetStartDate(v string) *CfEmploymentRecordResp {
	s.StartDate = &v
	return s
}

func (s *CfEmploymentRecordResp) SetWorkNumbers(v string) *CfEmploymentRecordResp {
	s.WorkNumbers = &v
	return s
}

type CfJobLevelResp struct {
	Level     *int64  `json:"level,omitempty" xml:"level,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	StopDate  *string `json:"stopDate,omitempty" xml:"stopDate,omitempty"`
}

func (s CfJobLevelResp) String() string {
	return tea.Prettify(s)
}

func (s CfJobLevelResp) GoString() string {
	return s.String()
}

func (s *CfJobLevelResp) SetLevel(v int64) *CfJobLevelResp {
	s.Level = &v
	return s
}

func (s *CfJobLevelResp) SetName(v string) *CfJobLevelResp {
	s.Name = &v
	return s
}

func (s *CfJobLevelResp) SetStartDate(v string) *CfJobLevelResp {
	s.StartDate = &v
	return s
}

func (s *CfJobLevelResp) SetStopDate(v string) *CfJobLevelResp {
	s.StopDate = &v
	return s
}

type CfJobPositionResp struct {
	JobPositionCode *string `json:"jobPositionCode,omitempty" xml:"jobPositionCode,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CfJobPositionResp) String() string {
	return tea.Prettify(s)
}

func (s CfJobPositionResp) GoString() string {
	return s.String()
}

func (s *CfJobPositionResp) SetJobPositionCode(v string) *CfJobPositionResp {
	s.JobPositionCode = &v
	return s
}

func (s *CfJobPositionResp) SetName(v string) *CfJobPositionResp {
	s.Name = &v
	return s
}

type CfJobPostResp struct {
	JobPostCode *string `json:"jobPostCode,omitempty" xml:"jobPostCode,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CfJobPostResp) String() string {
	return tea.Prettify(s)
}

func (s CfJobPostResp) GoString() string {
	return s.String()
}

func (s *CfJobPostResp) SetJobPostCode(v string) *CfJobPostResp {
	s.JobPostCode = &v
	return s
}

func (s *CfJobPostResp) SetName(v string) *CfJobPostResp {
	s.Name = &v
	return s
}

type CfOrgResp struct {
	DeptCode             *string `json:"deptCode,omitempty" xml:"deptCode,omitempty"`
	DeptName             *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	Level                *int64  `json:"level,omitempty" xml:"level,omitempty"`
	OrganizationCodePath *string `json:"organizationCodePath,omitempty" xml:"organizationCodePath,omitempty"`
	OrganizationPath     *string `json:"organizationPath,omitempty" xml:"organizationPath,omitempty"`
	ParentDeptCode       *string `json:"parentDeptCode,omitempty" xml:"parentDeptCode,omitempty"`
}

func (s CfOrgResp) String() string {
	return tea.Prettify(s)
}

func (s CfOrgResp) GoString() string {
	return s.String()
}

func (s *CfOrgResp) SetDeptCode(v string) *CfOrgResp {
	s.DeptCode = &v
	return s
}

func (s *CfOrgResp) SetDeptName(v string) *CfOrgResp {
	s.DeptName = &v
	return s
}

func (s *CfOrgResp) SetLevel(v int64) *CfOrgResp {
	s.Level = &v
	return s
}

func (s *CfOrgResp) SetOrganizationCodePath(v string) *CfOrgResp {
	s.OrganizationCodePath = &v
	return s
}

func (s *CfOrgResp) SetOrganizationPath(v string) *CfOrgResp {
	s.OrganizationPath = &v
	return s
}

func (s *CfOrgResp) SetParentDeptCode(v string) *CfOrgResp {
	s.ParentDeptCode = &v
	return s
}

type CfStaffResp struct {
	DeptCode    *string `json:"deptCode,omitempty" xml:"deptCode,omitempty"`
	DeptName    *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	Email       *string `json:"email,omitempty" xml:"email,omitempty"`
	Mobile      *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	NickName    *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	WorkNumbers *string `json:"workNumbers,omitempty" xml:"workNumbers,omitempty"`
}

func (s CfStaffResp) String() string {
	return tea.Prettify(s)
}

func (s CfStaffResp) GoString() string {
	return s.String()
}

func (s *CfStaffResp) SetDeptCode(v string) *CfStaffResp {
	s.DeptCode = &v
	return s
}

func (s *CfStaffResp) SetDeptName(v string) *CfStaffResp {
	s.DeptName = &v
	return s
}

func (s *CfStaffResp) SetEmail(v string) *CfStaffResp {
	s.Email = &v
	return s
}

func (s *CfStaffResp) SetMobile(v string) *CfStaffResp {
	s.Mobile = &v
	return s
}

func (s *CfStaffResp) SetName(v string) *CfStaffResp {
	s.Name = &v
	return s
}

func (s *CfStaffResp) SetNickName(v string) *CfStaffResp {
	s.NickName = &v
	return s
}

func (s *CfStaffResp) SetWorkNumbers(v string) *CfStaffResp {
	s.WorkNumbers = &v
	return s
}

type OpenAnalyzeDataDTO struct {
	DeptCount                    *int32   `json:"deptCount,omitempty" xml:"deptCount,omitempty"`
	NoAlignObjectiveCount        *int32   `json:"noAlignObjectiveCount,omitempty" xml:"noAlignObjectiveCount,omitempty"`
	NoKeyActionCount             *int32   `json:"noKeyActionCount,omitempty" xml:"noKeyActionCount,omitempty"`
	ObjectiveAlignRate           *float64 `json:"objectiveAlignRate,omitempty" xml:"objectiveAlignRate,omitempty"`
	ObjectiveNoSetCount          *int32   `json:"objectiveNoSetCount,omitempty" xml:"objectiveNoSetCount,omitempty"`
	ObjectiveRiskCount           *int32   `json:"objectiveRiskCount,omitempty" xml:"objectiveRiskCount,omitempty"`
	ObjectiveSetRate             *float64 `json:"objectiveSetRate,omitempty" xml:"objectiveSetRate,omitempty"`
	OnlyOneKeyResultCount        *int32   `json:"onlyOneKeyResultCount,omitempty" xml:"onlyOneKeyResultCount,omitempty"`
	OnlyOneObjectiveCount        *int32   `json:"onlyOneObjectiveCount,omitempty" xml:"onlyOneObjectiveCount,omitempty"`
	ProgressUpdateRateLast15Days *float64 `json:"progressUpdateRateLast15Days,omitempty" xml:"progressUpdateRateLast15Days,omitempty"`
	ProgressUpdateRateLast30Days *float64 `json:"progressUpdateRateLast30Days,omitempty" xml:"progressUpdateRateLast30Days,omitempty"`
	ProgressUpdateRateLast7Days  *float64 `json:"progressUpdateRateLast7Days,omitempty" xml:"progressUpdateRateLast7Days,omitempty"`
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
	Id            *string         `json:"id,omitempty" xml:"id,omitempty"`
	Progress      *int32          `json:"progress,omitempty" xml:"progress,omitempty"`
	Status        *int32          `json:"status,omitempty" xml:"status,omitempty"`
	Title         *string         `json:"title,omitempty" xml:"title,omitempty"`
	TitleMentions []*TitleMention `json:"titleMentions,omitempty" xml:"titleMentions,omitempty" type:"Repeated"`
	Type          *int32          `json:"type,omitempty" xml:"type,omitempty"`
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
	Executor   *OpenUserDTO        `json:"executor,omitempty" xml:"executor,omitempty"`
	Id         *string             `json:"id,omitempty" xml:"id,omitempty"`
	KeyResults []*OpenKeyResultDTO `json:"keyResults,omitempty" xml:"keyResults,omitempty" type:"Repeated"`
	Period     *OpenPeriodDTO      `json:"period,omitempty" xml:"period,omitempty"`
	Progress   *int32              `json:"progress,omitempty" xml:"progress,omitempty"`
	Status     *int32              `json:"status,omitempty" xml:"status,omitempty"`
	Teams      []*OpenTeamDTO      `json:"teams,omitempty" xml:"teams,omitempty" type:"Repeated"`
	Title      *string             `json:"title,omitempty" xml:"title,omitempty"`
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
	EndDate       *int64  `json:"endDate,omitempty" xml:"endDate,omitempty"`
	Id            *string `json:"id,omitempty" xml:"id,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	PeriodBizType *string `json:"periodBizType,omitempty" xml:"periodBizType,omitempty"`
	StartDate     *int64  `json:"startDate,omitempty" xml:"startDate,omitempty"`
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
	Created     *int64       `json:"created,omitempty" xml:"created,omitempty"`
	Creator     *OpenUserDTO `json:"creator,omitempty" xml:"creator,omitempty"`
	HtmlContent *string      `json:"htmlContent,omitempty" xml:"htmlContent,omitempty"`
	Id          *string      `json:"id,omitempty" xml:"id,omitempty"`
	Modifier    *OpenUserDTO `json:"modifier,omitempty" xml:"modifier,omitempty"`
	Updated     *int64       `json:"updated,omitempty" xml:"updated,omitempty"`
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
	Id     *string `json:"id,omitempty" xml:"id,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
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
	Id     *string `json:"id,omitempty" xml:"id,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
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
	Length *int32       `json:"length,omitempty" xml:"length,omitempty"`
	Offset *int32       `json:"offset,omitempty" xml:"offset,omitempty"`
	User   *OpenUserDTO `json:"user,omitempty" xml:"user,omitempty"`
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

type GetAllJobLevelHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAllJobLevelHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAllJobLevelHeaders) GoString() string {
	return s.String()
}

func (s *GetAllJobLevelHeaders) SetCommonHeaders(v map[string]*string) *GetAllJobLevelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAllJobLevelHeaders) SetXAcsDingtalkAccessToken(v string) *GetAllJobLevelHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAllJobLevelResponseBody struct {
	Content   []*CfJobLevelResp `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	RequestId *string           `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetAllJobLevelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAllJobLevelResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllJobLevelResponseBody) SetContent(v []*CfJobLevelResp) *GetAllJobLevelResponseBody {
	s.Content = v
	return s
}

func (s *GetAllJobLevelResponseBody) SetRequestId(v string) *GetAllJobLevelResponseBody {
	s.RequestId = &v
	return s
}

type GetAllJobLevelResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAllJobLevelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAllJobLevelResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllJobLevelResponse) GoString() string {
	return s.String()
}

func (s *GetAllJobLevelResponse) SetHeaders(v map[string]*string) *GetAllJobLevelResponse {
	s.Headers = v
	return s
}

func (s *GetAllJobLevelResponse) SetStatusCode(v int32) *GetAllJobLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAllJobLevelResponse) SetBody(v *GetAllJobLevelResponseBody) *GetAllJobLevelResponse {
	s.Body = v
	return s
}

type GetAllJobPositionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAllJobPositionHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAllJobPositionHeaders) GoString() string {
	return s.String()
}

func (s *GetAllJobPositionHeaders) SetCommonHeaders(v map[string]*string) *GetAllJobPositionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAllJobPositionHeaders) SetXAcsDingtalkAccessToken(v string) *GetAllJobPositionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAllJobPositionResponseBody struct {
	Content   []*CfJobPositionResp `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	RequestId *string              `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetAllJobPositionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAllJobPositionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllJobPositionResponseBody) SetContent(v []*CfJobPositionResp) *GetAllJobPositionResponseBody {
	s.Content = v
	return s
}

func (s *GetAllJobPositionResponseBody) SetRequestId(v string) *GetAllJobPositionResponseBody {
	s.RequestId = &v
	return s
}

type GetAllJobPositionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAllJobPositionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAllJobPositionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllJobPositionResponse) GoString() string {
	return s.String()
}

func (s *GetAllJobPositionResponse) SetHeaders(v map[string]*string) *GetAllJobPositionResponse {
	s.Headers = v
	return s
}

func (s *GetAllJobPositionResponse) SetStatusCode(v int32) *GetAllJobPositionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAllJobPositionResponse) SetBody(v *GetAllJobPositionResponseBody) *GetAllJobPositionResponse {
	s.Body = v
	return s
}

type GetAllJobPostHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAllJobPostHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAllJobPostHeaders) GoString() string {
	return s.String()
}

func (s *GetAllJobPostHeaders) SetCommonHeaders(v map[string]*string) *GetAllJobPostHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAllJobPostHeaders) SetXAcsDingtalkAccessToken(v string) *GetAllJobPostHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAllJobPostResponseBody struct {
	Content   []*CfJobPostResp `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	RequestId *string          `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetAllJobPostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAllJobPostResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllJobPostResponseBody) SetContent(v []*CfJobPostResp) *GetAllJobPostResponseBody {
	s.Content = v
	return s
}

func (s *GetAllJobPostResponseBody) SetRequestId(v string) *GetAllJobPostResponseBody {
	s.RequestId = &v
	return s
}

type GetAllJobPostResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAllJobPostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAllJobPostResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllJobPostResponse) GoString() string {
	return s.String()
}

func (s *GetAllJobPostResponse) SetHeaders(v map[string]*string) *GetAllJobPostResponse {
	s.Headers = v
	return s
}

func (s *GetAllJobPostResponse) SetStatusCode(v int32) *GetAllJobPostResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAllJobPostResponse) SetBody(v *GetAllJobPostResponseBody) *GetAllJobPostResponse {
	s.Body = v
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
	PeriodIds []*string `json:"periodIds,omitempty" xml:"periodIds,omitempty" type:"Repeated"`
	DeptId    *string   `json:"deptId,omitempty" xml:"deptId,omitempty"`
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
	Content   *OpenAnalyzeDataDTO `json:"content,omitempty" xml:"content,omitempty"`
	RequestId *string             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool               `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAnalyzeDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetAnalyzeDataResponse) SetStatusCode(v int32) *GetAnalyzeDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAnalyzeDataResponse) SetBody(v *GetAnalyzeDataResponseBody) *GetAnalyzeDataResponse {
	s.Body = v
	return s
}

type GetChildOrgListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetChildOrgListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetChildOrgListHeaders) GoString() string {
	return s.String()
}

func (s *GetChildOrgListHeaders) SetCommonHeaders(v map[string]*string) *GetChildOrgListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetChildOrgListHeaders) SetXAcsDingtalkAccessToken(v string) *GetChildOrgListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetChildOrgListRequest struct {
	DeptCode *string `json:"deptCode,omitempty" xml:"deptCode,omitempty"`
}

func (s GetChildOrgListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetChildOrgListRequest) GoString() string {
	return s.String()
}

func (s *GetChildOrgListRequest) SetDeptCode(v string) *GetChildOrgListRequest {
	s.DeptCode = &v
	return s
}

type GetChildOrgListResponseBody struct {
	Content   []*CfOrgResp `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	RequestId *string      `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetChildOrgListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetChildOrgListResponseBody) GoString() string {
	return s.String()
}

func (s *GetChildOrgListResponseBody) SetContent(v []*CfOrgResp) *GetChildOrgListResponseBody {
	s.Content = v
	return s
}

func (s *GetChildOrgListResponseBody) SetRequestId(v string) *GetChildOrgListResponseBody {
	s.RequestId = &v
	return s
}

type GetChildOrgListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetChildOrgListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetChildOrgListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetChildOrgListResponse) GoString() string {
	return s.String()
}

func (s *GetChildOrgListResponse) SetHeaders(v map[string]*string) *GetChildOrgListResponse {
	s.Headers = v
	return s
}

func (s *GetChildOrgListResponse) SetStatusCode(v int32) *GetChildOrgListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChildOrgListResponse) SetBody(v *GetChildOrgListResponseBody) *GetChildOrgListResponse {
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
	Content *GetEmployeeInfoByWorkNoResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	Success *bool                                       `json:"success,omitempty" xml:"success,omitempty"`
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
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetEmployeeInfoByWorkNoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetEmployeeInfoByWorkNoResponse) SetStatusCode(v int32) *GetEmployeeInfoByWorkNoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEmployeeInfoByWorkNoResponse) SetBody(v *GetEmployeeInfoByWorkNoResponseBody) *GetEmployeeInfoByWorkNoResponse {
	s.Body = v
	return s
}

type GetEmploymentRecordByWorkNoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetEmploymentRecordByWorkNoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetEmploymentRecordByWorkNoHeaders) GoString() string {
	return s.String()
}

func (s *GetEmploymentRecordByWorkNoHeaders) SetCommonHeaders(v map[string]*string) *GetEmploymentRecordByWorkNoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetEmploymentRecordByWorkNoHeaders) SetXAcsDingtalkAccessToken(v string) *GetEmploymentRecordByWorkNoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetEmploymentRecordByWorkNoResponseBody struct {
	Content   []*CfEmploymentRecordResp `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	RequestId *string                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetEmploymentRecordByWorkNoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEmploymentRecordByWorkNoResponseBody) GoString() string {
	return s.String()
}

func (s *GetEmploymentRecordByWorkNoResponseBody) SetContent(v []*CfEmploymentRecordResp) *GetEmploymentRecordByWorkNoResponseBody {
	s.Content = v
	return s
}

func (s *GetEmploymentRecordByWorkNoResponseBody) SetRequestId(v string) *GetEmploymentRecordByWorkNoResponseBody {
	s.RequestId = &v
	return s
}

type GetEmploymentRecordByWorkNoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetEmploymentRecordByWorkNoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEmploymentRecordByWorkNoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEmploymentRecordByWorkNoResponse) GoString() string {
	return s.String()
}

func (s *GetEmploymentRecordByWorkNoResponse) SetHeaders(v map[string]*string) *GetEmploymentRecordByWorkNoResponse {
	s.Headers = v
	return s
}

func (s *GetEmploymentRecordByWorkNoResponse) SetStatusCode(v int32) *GetEmploymentRecordByWorkNoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEmploymentRecordByWorkNoResponse) SetBody(v *GetEmploymentRecordByWorkNoResponseBody) *GetEmploymentRecordByWorkNoResponse {
	s.Body = v
	return s
}

type GetJobPositionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetJobPositionHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetJobPositionHeaders) GoString() string {
	return s.String()
}

func (s *GetJobPositionHeaders) SetCommonHeaders(v map[string]*string) *GetJobPositionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetJobPositionHeaders) SetXAcsDingtalkAccessToken(v string) *GetJobPositionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetJobPositionRequest struct {
	JobPositionCode *string `json:"jobPositionCode,omitempty" xml:"jobPositionCode,omitempty"`
}

func (s GetJobPositionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobPositionRequest) GoString() string {
	return s.String()
}

func (s *GetJobPositionRequest) SetJobPositionCode(v string) *GetJobPositionRequest {
	s.JobPositionCode = &v
	return s
}

type GetJobPositionResponseBody struct {
	Content   *GetJobPositionResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetJobPositionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobPositionResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobPositionResponseBody) SetContent(v *GetJobPositionResponseBodyContent) *GetJobPositionResponseBody {
	s.Content = v
	return s
}

func (s *GetJobPositionResponseBody) SetRequestId(v string) *GetJobPositionResponseBody {
	s.RequestId = &v
	return s
}

type GetJobPositionResponseBodyContent struct {
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	EstablishDate   *string `json:"establishDate,omitempty" xml:"establishDate,omitempty"`
	JobCode         *string `json:"jobCode,omitempty" xml:"jobCode,omitempty"`
	JobRequirements *string `json:"jobRequirements,omitempty" xml:"jobRequirements,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	StartDate       *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	StopDate        *string `json:"stopDate,omitempty" xml:"stopDate,omitempty"`
}

func (s GetJobPositionResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetJobPositionResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetJobPositionResponseBodyContent) SetDescription(v string) *GetJobPositionResponseBodyContent {
	s.Description = &v
	return s
}

func (s *GetJobPositionResponseBodyContent) SetEstablishDate(v string) *GetJobPositionResponseBodyContent {
	s.EstablishDate = &v
	return s
}

func (s *GetJobPositionResponseBodyContent) SetJobCode(v string) *GetJobPositionResponseBodyContent {
	s.JobCode = &v
	return s
}

func (s *GetJobPositionResponseBodyContent) SetJobRequirements(v string) *GetJobPositionResponseBodyContent {
	s.JobRequirements = &v
	return s
}

func (s *GetJobPositionResponseBodyContent) SetName(v string) *GetJobPositionResponseBodyContent {
	s.Name = &v
	return s
}

func (s *GetJobPositionResponseBodyContent) SetStartDate(v string) *GetJobPositionResponseBodyContent {
	s.StartDate = &v
	return s
}

func (s *GetJobPositionResponseBodyContent) SetStopDate(v string) *GetJobPositionResponseBodyContent {
	s.StopDate = &v
	return s
}

type GetJobPositionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetJobPositionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobPositionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobPositionResponse) GoString() string {
	return s.String()
}

func (s *GetJobPositionResponse) SetHeaders(v map[string]*string) *GetJobPositionResponse {
	s.Headers = v
	return s
}

func (s *GetJobPositionResponse) SetStatusCode(v int32) *GetJobPositionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobPositionResponse) SetBody(v *GetJobPositionResponseBody) *GetJobPositionResponse {
	s.Body = v
	return s
}

type GetJobPostHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetJobPostHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetJobPostHeaders) GoString() string {
	return s.String()
}

func (s *GetJobPostHeaders) SetCommonHeaders(v map[string]*string) *GetJobPostHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetJobPostHeaders) SetXAcsDingtalkAccessToken(v string) *GetJobPostHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetJobPostRequest struct {
	JobPostCode *string `json:"jobPostCode,omitempty" xml:"jobPostCode,omitempty"`
}

func (s GetJobPostRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobPostRequest) GoString() string {
	return s.String()
}

func (s *GetJobPostRequest) SetJobPostCode(v string) *GetJobPostRequest {
	s.JobPostCode = &v
	return s
}

type GetJobPostResponseBody struct {
	Content   *GetJobPostResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	RequestId *string                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetJobPostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobPostResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobPostResponseBody) SetContent(v *GetJobPostResponseBodyContent) *GetJobPostResponseBody {
	s.Content = v
	return s
}

func (s *GetJobPostResponseBody) SetRequestId(v string) *GetJobPostResponseBody {
	s.RequestId = &v
	return s
}

type GetJobPostResponseBodyContent struct {
	Code          *string `json:"code,omitempty" xml:"code,omitempty"`
	EstablishDate *string `json:"establishDate,omitempty" xml:"establishDate,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	StartDate     *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	StopDate      *string `json:"stopDate,omitempty" xml:"stopDate,omitempty"`
}

func (s GetJobPostResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetJobPostResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetJobPostResponseBodyContent) SetCode(v string) *GetJobPostResponseBodyContent {
	s.Code = &v
	return s
}

func (s *GetJobPostResponseBodyContent) SetEstablishDate(v string) *GetJobPostResponseBodyContent {
	s.EstablishDate = &v
	return s
}

func (s *GetJobPostResponseBodyContent) SetName(v string) *GetJobPostResponseBodyContent {
	s.Name = &v
	return s
}

func (s *GetJobPostResponseBodyContent) SetStartDate(v string) *GetJobPostResponseBodyContent {
	s.StartDate = &v
	return s
}

func (s *GetJobPostResponseBodyContent) SetStopDate(v string) *GetJobPostResponseBodyContent {
	s.StopDate = &v
	return s
}

type GetJobPostResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetJobPostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobPostResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobPostResponse) GoString() string {
	return s.String()
}

func (s *GetJobPostResponse) SetHeaders(v map[string]*string) *GetJobPostResponse {
	s.Headers = v
	return s
}

func (s *GetJobPostResponse) SetStatusCode(v int32) *GetJobPostResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobPostResponse) SetBody(v *GetJobPostResponseBody) *GetJobPostResponse {
	s.Body = v
	return s
}

type GetOrgInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOrgInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOrgInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetOrgInfoHeaders) SetCommonHeaders(v map[string]*string) *GetOrgInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOrgInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetOrgInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOrgInfoRequest struct {
	DeptCode *string `json:"deptCode,omitempty" xml:"deptCode,omitempty"`
}

func (s GetOrgInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOrgInfoRequest) GoString() string {
	return s.String()
}

func (s *GetOrgInfoRequest) SetDeptCode(v string) *GetOrgInfoRequest {
	s.DeptCode = &v
	return s
}

type GetOrgInfoResponseBody struct {
	Content   *GetOrgInfoResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	RequestId *string                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetOrgInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrgInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrgInfoResponseBody) SetContent(v *GetOrgInfoResponseBodyContent) *GetOrgInfoResponseBody {
	s.Content = v
	return s
}

func (s *GetOrgInfoResponseBody) SetRequestId(v string) *GetOrgInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetOrgInfoResponseBodyContent struct {
	DeptCode             *string `json:"deptCode,omitempty" xml:"deptCode,omitempty"`
	DeptName             *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	DeptNum              *string `json:"deptNum,omitempty" xml:"deptNum,omitempty"`
	Level                *string `json:"level,omitempty" xml:"level,omitempty"`
	OrganizationCodePath *string `json:"organizationCodePath,omitempty" xml:"organizationCodePath,omitempty"`
	OrganizationPath     *string `json:"organizationPath,omitempty" xml:"organizationPath,omitempty"`
	ParentDeptCode       *string `json:"parentDeptCode,omitempty" xml:"parentDeptCode,omitempty"`
	ShortName            *string `json:"shortName,omitempty" xml:"shortName,omitempty"`
	StartDate            *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	StopDate             *string `json:"stopDate,omitempty" xml:"stopDate,omitempty"`
}

func (s GetOrgInfoResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetOrgInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetOrgInfoResponseBodyContent) SetDeptCode(v string) *GetOrgInfoResponseBodyContent {
	s.DeptCode = &v
	return s
}

func (s *GetOrgInfoResponseBodyContent) SetDeptName(v string) *GetOrgInfoResponseBodyContent {
	s.DeptName = &v
	return s
}

func (s *GetOrgInfoResponseBodyContent) SetDeptNum(v string) *GetOrgInfoResponseBodyContent {
	s.DeptNum = &v
	return s
}

func (s *GetOrgInfoResponseBodyContent) SetLevel(v string) *GetOrgInfoResponseBodyContent {
	s.Level = &v
	return s
}

func (s *GetOrgInfoResponseBodyContent) SetOrganizationCodePath(v string) *GetOrgInfoResponseBodyContent {
	s.OrganizationCodePath = &v
	return s
}

func (s *GetOrgInfoResponseBodyContent) SetOrganizationPath(v string) *GetOrgInfoResponseBodyContent {
	s.OrganizationPath = &v
	return s
}

func (s *GetOrgInfoResponseBodyContent) SetParentDeptCode(v string) *GetOrgInfoResponseBodyContent {
	s.ParentDeptCode = &v
	return s
}

func (s *GetOrgInfoResponseBodyContent) SetShortName(v string) *GetOrgInfoResponseBodyContent {
	s.ShortName = &v
	return s
}

func (s *GetOrgInfoResponseBodyContent) SetStartDate(v string) *GetOrgInfoResponseBodyContent {
	s.StartDate = &v
	return s
}

func (s *GetOrgInfoResponseBodyContent) SetStopDate(v string) *GetOrgInfoResponseBodyContent {
	s.StopDate = &v
	return s
}

type GetOrgInfoResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOrgInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOrgInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrgInfoResponse) GoString() string {
	return s.String()
}

func (s *GetOrgInfoResponse) SetHeaders(v map[string]*string) *GetOrgInfoResponse {
	s.Headers = v
	return s
}

func (s *GetOrgInfoResponse) SetStatusCode(v int32) *GetOrgInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrgInfoResponse) SetBody(v *GetOrgInfoResponseBody) *GetOrgInfoResponse {
	s.Body = v
	return s
}

type GetStaffInfoByWorkNoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetStaffInfoByWorkNoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetStaffInfoByWorkNoHeaders) GoString() string {
	return s.String()
}

func (s *GetStaffInfoByWorkNoHeaders) SetCommonHeaders(v map[string]*string) *GetStaffInfoByWorkNoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetStaffInfoByWorkNoHeaders) SetXAcsDingtalkAccessToken(v string) *GetStaffInfoByWorkNoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetStaffInfoByWorkNoRequest struct {
	WorkNumbers *string `json:"workNumbers,omitempty" xml:"workNumbers,omitempty"`
}

func (s GetStaffInfoByWorkNoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStaffInfoByWorkNoRequest) GoString() string {
	return s.String()
}

func (s *GetStaffInfoByWorkNoRequest) SetWorkNumbers(v string) *GetStaffInfoByWorkNoRequest {
	s.WorkNumbers = &v
	return s
}

type GetStaffInfoByWorkNoResponseBody struct {
	Content   *GetStaffInfoByWorkNoResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetStaffInfoByWorkNoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStaffInfoByWorkNoResponseBody) GoString() string {
	return s.String()
}

func (s *GetStaffInfoByWorkNoResponseBody) SetContent(v *GetStaffInfoByWorkNoResponseBodyContent) *GetStaffInfoByWorkNoResponseBody {
	s.Content = v
	return s
}

func (s *GetStaffInfoByWorkNoResponseBody) SetRequestId(v string) *GetStaffInfoByWorkNoResponseBody {
	s.RequestId = &v
	return s
}

type GetStaffInfoByWorkNoResponseBodyContent struct {
	DeptCode        *string `json:"deptCode,omitempty" xml:"deptCode,omitempty"`
	DeptName        *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	Email           *string `json:"email,omitempty" xml:"email,omitempty"`
	EmployType      *string `json:"employType,omitempty" xml:"employType,omitempty"`
	EmployeeStatus  *string `json:"employeeStatus,omitempty" xml:"employeeStatus,omitempty"`
	JobLevelName    *string `json:"jobLevelName,omitempty" xml:"jobLevelName,omitempty"`
	JobPositionCode *string `json:"jobPositionCode,omitempty" xml:"jobPositionCode,omitempty"`
	JobPositionName *string `json:"jobPositionName,omitempty" xml:"jobPositionName,omitempty"`
	JobPostCode     *string `json:"jobPostCode,omitempty" xml:"jobPostCode,omitempty"`
	JobPostName     *string `json:"jobPostName,omitempty" xml:"jobPostName,omitempty"`
	Mobile          *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	NickName        *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	WorkNumbers     *string `json:"workNumbers,omitempty" xml:"workNumbers,omitempty"`
}

func (s GetStaffInfoByWorkNoResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetStaffInfoByWorkNoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetStaffInfoByWorkNoResponseBodyContent) SetDeptCode(v string) *GetStaffInfoByWorkNoResponseBodyContent {
	s.DeptCode = &v
	return s
}

func (s *GetStaffInfoByWorkNoResponseBodyContent) SetDeptName(v string) *GetStaffInfoByWorkNoResponseBodyContent {
	s.DeptName = &v
	return s
}

func (s *GetStaffInfoByWorkNoResponseBodyContent) SetEmail(v string) *GetStaffInfoByWorkNoResponseBodyContent {
	s.Email = &v
	return s
}

func (s *GetStaffInfoByWorkNoResponseBodyContent) SetEmployType(v string) *GetStaffInfoByWorkNoResponseBodyContent {
	s.EmployType = &v
	return s
}

func (s *GetStaffInfoByWorkNoResponseBodyContent) SetEmployeeStatus(v string) *GetStaffInfoByWorkNoResponseBodyContent {
	s.EmployeeStatus = &v
	return s
}

func (s *GetStaffInfoByWorkNoResponseBodyContent) SetJobLevelName(v string) *GetStaffInfoByWorkNoResponseBodyContent {
	s.JobLevelName = &v
	return s
}

func (s *GetStaffInfoByWorkNoResponseBodyContent) SetJobPositionCode(v string) *GetStaffInfoByWorkNoResponseBodyContent {
	s.JobPositionCode = &v
	return s
}

func (s *GetStaffInfoByWorkNoResponseBodyContent) SetJobPositionName(v string) *GetStaffInfoByWorkNoResponseBodyContent {
	s.JobPositionName = &v
	return s
}

func (s *GetStaffInfoByWorkNoResponseBodyContent) SetJobPostCode(v string) *GetStaffInfoByWorkNoResponseBodyContent {
	s.JobPostCode = &v
	return s
}

func (s *GetStaffInfoByWorkNoResponseBodyContent) SetJobPostName(v string) *GetStaffInfoByWorkNoResponseBodyContent {
	s.JobPostName = &v
	return s
}

func (s *GetStaffInfoByWorkNoResponseBodyContent) SetMobile(v string) *GetStaffInfoByWorkNoResponseBodyContent {
	s.Mobile = &v
	return s
}

func (s *GetStaffInfoByWorkNoResponseBodyContent) SetName(v string) *GetStaffInfoByWorkNoResponseBodyContent {
	s.Name = &v
	return s
}

func (s *GetStaffInfoByWorkNoResponseBodyContent) SetNickName(v string) *GetStaffInfoByWorkNoResponseBodyContent {
	s.NickName = &v
	return s
}

func (s *GetStaffInfoByWorkNoResponseBodyContent) SetWorkNumbers(v string) *GetStaffInfoByWorkNoResponseBodyContent {
	s.WorkNumbers = &v
	return s
}

type GetStaffInfoByWorkNoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetStaffInfoByWorkNoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStaffInfoByWorkNoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStaffInfoByWorkNoResponse) GoString() string {
	return s.String()
}

func (s *GetStaffInfoByWorkNoResponse) SetHeaders(v map[string]*string) *GetStaffInfoByWorkNoResponse {
	s.Headers = v
	return s
}

func (s *GetStaffInfoByWorkNoResponse) SetStatusCode(v int32) *GetStaffInfoByWorkNoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStaffInfoByWorkNoResponse) SetBody(v *GetStaffInfoByWorkNoResponseBody) *GetStaffInfoByWorkNoResponse {
	s.Body = v
	return s
}

type GetStaffPageQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetStaffPageQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetStaffPageQueryHeaders) GoString() string {
	return s.String()
}

func (s *GetStaffPageQueryHeaders) SetCommonHeaders(v map[string]*string) *GetStaffPageQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetStaffPageQueryHeaders) SetXAcsDingtalkAccessToken(v string) *GetStaffPageQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetStaffPageQueryRequest struct {
	DeptCode   *string `json:"deptCode,omitempty" xml:"deptCode,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	WorkNo     *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s GetStaffPageQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStaffPageQueryRequest) GoString() string {
	return s.String()
}

func (s *GetStaffPageQueryRequest) SetDeptCode(v string) *GetStaffPageQueryRequest {
	s.DeptCode = &v
	return s
}

func (s *GetStaffPageQueryRequest) SetName(v string) *GetStaffPageQueryRequest {
	s.Name = &v
	return s
}

func (s *GetStaffPageQueryRequest) SetPageNumber(v int32) *GetStaffPageQueryRequest {
	s.PageNumber = &v
	return s
}

func (s *GetStaffPageQueryRequest) SetPageSize(v int32) *GetStaffPageQueryRequest {
	s.PageSize = &v
	return s
}

func (s *GetStaffPageQueryRequest) SetWorkNo(v string) *GetStaffPageQueryRequest {
	s.WorkNo = &v
	return s
}

type GetStaffPageQueryResponseBody struct {
	Content   *GetStaffPageQueryResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetStaffPageQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStaffPageQueryResponseBody) GoString() string {
	return s.String()
}

func (s *GetStaffPageQueryResponseBody) SetContent(v *GetStaffPageQueryResponseBodyContent) *GetStaffPageQueryResponseBody {
	s.Content = v
	return s
}

func (s *GetStaffPageQueryResponseBody) SetRequestId(v string) *GetStaffPageQueryResponseBody {
	s.RequestId = &v
	return s
}

type GetStaffPageQueryResponseBodyContent struct {
	Data       []*CfStaffResp `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	PageNumber *int32         `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32         `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	TotalCount *int64         `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetStaffPageQueryResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetStaffPageQueryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetStaffPageQueryResponseBodyContent) SetData(v []*CfStaffResp) *GetStaffPageQueryResponseBodyContent {
	s.Data = v
	return s
}

func (s *GetStaffPageQueryResponseBodyContent) SetPageNumber(v int32) *GetStaffPageQueryResponseBodyContent {
	s.PageNumber = &v
	return s
}

func (s *GetStaffPageQueryResponseBodyContent) SetPageSize(v int32) *GetStaffPageQueryResponseBodyContent {
	s.PageSize = &v
	return s
}

func (s *GetStaffPageQueryResponseBodyContent) SetTotalCount(v int64) *GetStaffPageQueryResponseBodyContent {
	s.TotalCount = &v
	return s
}

type GetStaffPageQueryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetStaffPageQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStaffPageQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStaffPageQueryResponse) GoString() string {
	return s.String()
}

func (s *GetStaffPageQueryResponse) SetHeaders(v map[string]*string) *GetStaffPageQueryResponse {
	s.Headers = v
	return s
}

func (s *GetStaffPageQueryResponse) SetStatusCode(v int32) *GetStaffPageQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStaffPageQueryResponse) SetBody(v *GetStaffPageQueryResponseBody) *GetStaffPageQueryResponse {
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
	OkrUserId *string `json:"okrUserId,omitempty" xml:"okrUserId,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Content   *OpenUserDTO `json:"content,omitempty" xml:"content,omitempty"`
	RequestId *string      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool        `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetUserResponse) SetStatusCode(v int32) *GetUserResponse {
	s.StatusCode = &v
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
	Content   []*OpenPeriodDTO `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	RequestId *string          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool            `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAnalyzePeriodsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListAnalyzePeriodsResponse) SetStatusCode(v int32) *ListAnalyzePeriodsResponse {
	s.StatusCode = &v
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
	Content   []*OpenObjectiveDTO `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	RequestId *string             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool               `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListObjectiveByIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListObjectiveByIdsResponse) SetStatusCode(v int32) *ListObjectiveByIdsResponse {
	s.StatusCode = &v
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
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	UserId     *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Content   *ListObjectiveByUserResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	RequestId *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                   `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListObjectiveByUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListObjectiveByUserResponse) SetStatusCode(v int32) *ListObjectiveByUserResponse {
	s.StatusCode = &v
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
	Content   []*OpenProgressDTO `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	RequestId *string            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool              `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProgressByIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListProgressByIdsResponse) SetStatusCode(v int32) *ListProgressByIdsResponse {
	s.StatusCode = &v
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
	ObjectiveId *string `json:"objectiveId,omitempty" xml:"objectiveId,omitempty"`
	PageNumber  *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize    *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
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
	Content   *PageListObjectiveProgressResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	RequestId *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                         `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PageListObjectiveProgressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *PageListObjectiveProgressResponse) SetStatusCode(v int32) *PageListObjectiveProgressResponse {
	s.StatusCode = &v
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
	ObjectiveId  *string `json:"objectiveId,omitempty" xml:"objectiveId,omitempty"`
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
	Content   *bool   `json:"content,omitempty" xml:"content,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TransferUserObjectiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TransferUserObjectiveResponse) SetStatusCode(v int32) *TransferUserObjectiveResponse {
	s.StatusCode = &v
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

func (client *Client) GetAllJobLevelWithOptions(headers *GetAllJobLevelHeaders, runtime *util.RuntimeOptions) (_result *GetAllJobLevelResponse, _err error) {
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
		Action:      tea.String("GetAllJobLevel"),
		Version:     tea.String("chengfeng_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/chengfeng/jobLevels"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAllJobLevelResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAllJobLevel() (_result *GetAllJobLevelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAllJobLevelHeaders{}
	_result = &GetAllJobLevelResponse{}
	_body, _err := client.GetAllJobLevelWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAllJobPositionWithOptions(headers *GetAllJobPositionHeaders, runtime *util.RuntimeOptions) (_result *GetAllJobPositionResponse, _err error) {
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
		Action:      tea.String("GetAllJobPosition"),
		Version:     tea.String("chengfeng_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/chengfeng/jobPositions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAllJobPositionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAllJobPosition() (_result *GetAllJobPositionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAllJobPositionHeaders{}
	_result = &GetAllJobPositionResponse{}
	_body, _err := client.GetAllJobPositionWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAllJobPostWithOptions(headers *GetAllJobPostHeaders, runtime *util.RuntimeOptions) (_result *GetAllJobPostResponse, _err error) {
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
		Action:      tea.String("GetAllJobPost"),
		Version:     tea.String("chengfeng_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/chengfeng/jobPosts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAllJobPostResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAllJobPost() (_result *GetAllJobPostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAllJobPostHeaders{}
	_result = &GetAllJobPostResponse{}
	_body, _err := client.GetAllJobPostWithOptions(headers, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetAnalyzeData"),
		Version:     tea.String("chengfeng_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/chengfeng/okr/analyses/datas/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAnalyzeDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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

func (client *Client) GetChildOrgListWithOptions(request *GetChildOrgListRequest, headers *GetChildOrgListHeaders, runtime *util.RuntimeOptions) (_result *GetChildOrgListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptCode)) {
		query["deptCode"] = request.DeptCode
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
		Action:      tea.String("GetChildOrgList"),
		Version:     tea.String("chengfeng_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/chengfeng/organizations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetChildOrgListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetChildOrgList(request *GetChildOrgListRequest) (_result *GetChildOrgListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetChildOrgListHeaders{}
	_result = &GetChildOrgListResponse{}
	_body, _err := client.GetChildOrgListWithOptions(request, headers, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetEmployeeInfoByWorkNo"),
		Version:     tea.String("chengfeng_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/chengfeng/workNumbers/employees"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEmployeeInfoByWorkNoResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) GetEmploymentRecordByWorkNoWithOptions(workNumbers *string, headers *GetEmploymentRecordByWorkNoHeaders, runtime *util.RuntimeOptions) (_result *GetEmploymentRecordByWorkNoResponse, _err error) {
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
		Action:      tea.String("GetEmploymentRecordByWorkNo"),
		Version:     tea.String("chengfeng_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/chengfeng/users/workNumber/" + tea.StringValue(workNumbers) + "employmentRecords"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEmploymentRecordByWorkNoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEmploymentRecordByWorkNo(workNumbers *string) (_result *GetEmploymentRecordByWorkNoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetEmploymentRecordByWorkNoHeaders{}
	_result = &GetEmploymentRecordByWorkNoResponse{}
	_body, _err := client.GetEmploymentRecordByWorkNoWithOptions(workNumbers, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobPositionWithOptions(request *GetJobPositionRequest, headers *GetJobPositionHeaders, runtime *util.RuntimeOptions) (_result *GetJobPositionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobPositionCode)) {
		query["jobPositionCode"] = request.JobPositionCode
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
		Action:      tea.String("GetJobPosition"),
		Version:     tea.String("chengfeng_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/chengfeng/jobPositions/infos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobPositionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobPosition(request *GetJobPositionRequest) (_result *GetJobPositionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetJobPositionHeaders{}
	_result = &GetJobPositionResponse{}
	_body, _err := client.GetJobPositionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobPostWithOptions(request *GetJobPostRequest, headers *GetJobPostHeaders, runtime *util.RuntimeOptions) (_result *GetJobPostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobPostCode)) {
		query["jobPostCode"] = request.JobPostCode
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
		Action:      tea.String("GetJobPost"),
		Version:     tea.String("chengfeng_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/chengfeng/jobPosts/infos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobPostResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobPost(request *GetJobPostRequest) (_result *GetJobPostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetJobPostHeaders{}
	_result = &GetJobPostResponse{}
	_body, _err := client.GetJobPostWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOrgInfoWithOptions(request *GetOrgInfoRequest, headers *GetOrgInfoHeaders, runtime *util.RuntimeOptions) (_result *GetOrgInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptCode)) {
		query["deptCode"] = request.DeptCode
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
		Action:      tea.String("GetOrgInfo"),
		Version:     tea.String("chengfeng_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/chengfeng/organizations/infos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOrgInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOrgInfo(request *GetOrgInfoRequest) (_result *GetOrgInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOrgInfoHeaders{}
	_result = &GetOrgInfoResponse{}
	_body, _err := client.GetOrgInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStaffInfoByWorkNoWithOptions(request *GetStaffInfoByWorkNoRequest, headers *GetStaffInfoByWorkNoHeaders, runtime *util.RuntimeOptions) (_result *GetStaffInfoByWorkNoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkNumbers)) {
		query["workNumbers"] = request.WorkNumbers
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
		Action:      tea.String("GetStaffInfoByWorkNo"),
		Version:     tea.String("chengfeng_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/chengfeng/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStaffInfoByWorkNoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStaffInfoByWorkNo(request *GetStaffInfoByWorkNoRequest) (_result *GetStaffInfoByWorkNoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetStaffInfoByWorkNoHeaders{}
	_result = &GetStaffInfoByWorkNoResponse{}
	_body, _err := client.GetStaffInfoByWorkNoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStaffPageQueryWithOptions(request *GetStaffPageQueryRequest, headers *GetStaffPageQueryHeaders, runtime *util.RuntimeOptions) (_result *GetStaffPageQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptCode)) {
		query["deptCode"] = request.DeptCode
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

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
	params := &openapi.Params{
		Action:      tea.String("GetStaffPageQuery"),
		Version:     tea.String("chengfeng_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/chengfeng/users/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStaffPageQueryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStaffPageQuery(request *GetStaffPageQueryRequest) (_result *GetStaffPageQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetStaffPageQueryHeaders{}
	_result = &GetStaffPageQueryResponse{}
	_body, _err := client.GetStaffPageQueryWithOptions(request, headers, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetUser"),
		Version:     tea.String("chengfeng_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/chengfeng/okr/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("ListAnalyzePeriods"),
		Version:     tea.String("chengfeng_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/chengfeng/okr/analyses/periods"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAnalyzePeriodsResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("ListObjectiveByIds"),
		Version:     tea.String("chengfeng_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/chengfeng/okr/objectives/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListObjectiveByIdsResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("ListObjectiveByUser"),
		Version:     tea.String("chengfeng_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/chengfeng/okr/users/objectives"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListObjectiveByUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("ListProgressByIds"),
		Version:     tea.String("chengfeng_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/chengfeng/okr/objectives/progresses/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProgressByIdsResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("PageListObjectiveProgress"),
		Version:     tea.String("chengfeng_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/chengfeng/okr/objectives/progresses/records"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PageListObjectiveProgressResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("TransferUserObjective"),
		Version:     tea.String("chengfeng_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/chengfeng/okr/objectives/transfer"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TransferUserObjectiveResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
