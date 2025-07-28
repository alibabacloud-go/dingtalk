// This file is auto-generated, don't edit it. Thanks.
package okr_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type OpenKeyResultDTO struct {
	// example:
	//
	// 65222713d0e8b868f9f9ae51
	KrId *string `json:"krId,omitempty" xml:"krId,omitempty"`
	// example:
	//
	// 80
	Progress *int64 `json:"progress,omitempty" xml:"progress,omitempty"`
	// example:
	//
	// 1
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 这是一个KR
	Title         *string         `json:"title,omitempty" xml:"title,omitempty"`
	TitleMentions []*TitleMention `json:"titleMentions,omitempty" xml:"titleMentions,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 10.00
	Weight *float64 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s OpenKeyResultDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenKeyResultDTO) GoString() string {
	return s.String()
}

func (s *OpenKeyResultDTO) SetKrId(v string) *OpenKeyResultDTO {
	s.KrId = &v
	return s
}

func (s *OpenKeyResultDTO) SetProgress(v int64) *OpenKeyResultDTO {
	s.Progress = &v
	return s
}

func (s *OpenKeyResultDTO) SetStatus(v int64) *OpenKeyResultDTO {
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

func (s *OpenKeyResultDTO) SetType(v int64) *OpenKeyResultDTO {
	s.Type = &v
	return s
}

func (s *OpenKeyResultDTO) SetWeight(v float64) *OpenKeyResultDTO {
	s.Weight = &v
	return s
}

type OpenObjectiveDTO struct {
	Executor   *OpenUserDTO        `json:"executor,omitempty" xml:"executor,omitempty"`
	KeyResults []*OpenKeyResultDTO `json:"keyResults,omitempty" xml:"keyResults,omitempty" type:"Repeated"`
	// example:
	//
	// 65222640d0e8b868f9f9ae3c
	ObjectiveId *string        `json:"objectiveId,omitempty" xml:"objectiveId,omitempty"`
	Period      *OpenPeriodDTO `json:"period,omitempty" xml:"period,omitempty"`
	// example:
	//
	// 80
	Progress *int64 `json:"progress,omitempty" xml:"progress,omitempty"`
	// example:
	//
	// 1
	Status *int64         `json:"status,omitempty" xml:"status,omitempty"`
	Teams  []*OpenTeamDTO `json:"teams,omitempty" xml:"teams,omitempty" type:"Repeated"`
	// example:
	//
	// 这是一个O的标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 10.00
	Weight *float64 `json:"weight,omitempty" xml:"weight,omitempty"`
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

func (s *OpenObjectiveDTO) SetKeyResults(v []*OpenKeyResultDTO) *OpenObjectiveDTO {
	s.KeyResults = v
	return s
}

func (s *OpenObjectiveDTO) SetObjectiveId(v string) *OpenObjectiveDTO {
	s.ObjectiveId = &v
	return s
}

func (s *OpenObjectiveDTO) SetPeriod(v *OpenPeriodDTO) *OpenObjectiveDTO {
	s.Period = v
	return s
}

func (s *OpenObjectiveDTO) SetProgress(v int64) *OpenObjectiveDTO {
	s.Progress = &v
	return s
}

func (s *OpenObjectiveDTO) SetStatus(v int64) *OpenObjectiveDTO {
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

func (s *OpenObjectiveDTO) SetWeight(v float64) *OpenObjectiveDTO {
	s.Weight = &v
	return s
}

type OpenPeriodDTO struct {
	EndDate *int64  `json:"endDate,omitempty" xml:"endDate,omitempty"`
	NameCn  *string `json:"nameCn,omitempty" xml:"nameCn,omitempty"`
	NameEn  *string `json:"nameEn,omitempty" xml:"nameEn,omitempty"`
	// This parameter is required.
	PeriodId  *string `json:"periodId,omitempty" xml:"periodId,omitempty"`
	StartDate *int64  `json:"startDate,omitempty" xml:"startDate,omitempty"`
	Status    *int32  `json:"status,omitempty" xml:"status,omitempty"`
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

func (s *OpenPeriodDTO) SetNameCn(v string) *OpenPeriodDTO {
	s.NameCn = &v
	return s
}

func (s *OpenPeriodDTO) SetNameEn(v string) *OpenPeriodDTO {
	s.NameEn = &v
	return s
}

func (s *OpenPeriodDTO) SetPeriodId(v string) *OpenPeriodDTO {
	s.PeriodId = &v
	return s
}

func (s *OpenPeriodDTO) SetStartDate(v int64) *OpenPeriodDTO {
	s.StartDate = &v
	return s
}

func (s *OpenPeriodDTO) SetStatus(v int32) *OpenPeriodDTO {
	s.Status = &v
	return s
}

type OpenTeamDTO struct {
	// example:
	//
	// 0342464558835299
	DeptUid *string `json:"deptUid,omitempty" xml:"deptUid,omitempty"`
	// example:
	//
	// 64cd2e9bb80efb17244c650d
	DingDeptId *string `json:"dingDeptId,omitempty" xml:"dingDeptId,omitempty"`
	// example:
	//
	// xx部门
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s OpenTeamDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenTeamDTO) GoString() string {
	return s.String()
}

func (s *OpenTeamDTO) SetDeptUid(v string) *OpenTeamDTO {
	s.DeptUid = &v
	return s
}

func (s *OpenTeamDTO) SetDingDeptId(v string) *OpenTeamDTO {
	s.DingDeptId = &v
	return s
}

func (s *OpenTeamDTO) SetName(v string) *OpenTeamDTO {
	s.Name = &v
	return s
}

type OpenUserDTO struct {
	// example:
	//
	// 0342464558835299
	DingUserId *string `json:"dingUserId,omitempty" xml:"dingUserId,omitempty"`
	// example:
	//
	// 小明
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 64cd2e9bb80efb17244c650d
	UserUid *string `json:"userUid,omitempty" xml:"userUid,omitempty"`
	// example:
	//
	// 2639402752-1812711657
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s OpenUserDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenUserDTO) GoString() string {
	return s.String()
}

func (s *OpenUserDTO) SetDingUserId(v string) *OpenUserDTO {
	s.DingUserId = &v
	return s
}

func (s *OpenUserDTO) SetName(v string) *OpenUserDTO {
	s.Name = &v
	return s
}

func (s *OpenUserDTO) SetUserUid(v string) *OpenUserDTO {
	s.UserUid = &v
	return s
}

func (s *OpenUserDTO) SetWorkNo(v string) *OpenUserDTO {
	s.WorkNo = &v
	return s
}

type TitleMention struct {
	// example:
	//
	// 20
	Length *int64 `json:"length,omitempty" xml:"length,omitempty"`
	// example:
	//
	// 1
	Offset *int64       `json:"offset,omitempty" xml:"offset,omitempty"`
	User   *OpenUserDTO `json:"user,omitempty" xml:"user,omitempty"`
}

func (s TitleMention) String() string {
	return tea.Prettify(s)
}

func (s TitleMention) GoString() string {
	return s.String()
}

func (s *TitleMention) SetLength(v int64) *TitleMention {
	s.Length = &v
	return s
}

func (s *TitleMention) SetOffset(v int64) *TitleMention {
	s.Offset = &v
	return s
}

func (s *TitleMention) SetUser(v *OpenUserDTO) *TitleMention {
	s.User = v
	return s
}

type AlignObjectiveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AlignObjectiveHeaders) String() string {
	return tea.Prettify(s)
}

func (s AlignObjectiveHeaders) GoString() string {
	return s.String()
}

func (s *AlignObjectiveHeaders) SetCommonHeaders(v map[string]*string) *AlignObjectiveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AlignObjectiveHeaders) SetXAcsDingtalkAccessToken(v string) *AlignObjectiveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AlignObjectiveRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1006
	PeriodId *string `json:"periodId,omitempty" xml:"periodId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 59YD
	TargetId *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0115396701752283
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AlignObjectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s AlignObjectiveRequest) GoString() string {
	return s.String()
}

func (s *AlignObjectiveRequest) SetPeriodId(v string) *AlignObjectiveRequest {
	s.PeriodId = &v
	return s
}

func (s *AlignObjectiveRequest) SetTargetId(v string) *AlignObjectiveRequest {
	s.TargetId = &v
	return s
}

func (s *AlignObjectiveRequest) SetUserId(v string) *AlignObjectiveRequest {
	s.UserId = &v
	return s
}

type AlignObjectiveResponseBody struct {
	Data *AlignObjectiveResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AlignObjectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AlignObjectiveResponseBody) GoString() string {
	return s.String()
}

func (s *AlignObjectiveResponseBody) SetData(v *AlignObjectiveResponseBodyData) *AlignObjectiveResponseBody {
	s.Data = v
	return s
}

func (s *AlignObjectiveResponseBody) SetSuccess(v bool) *AlignObjectiveResponseBody {
	s.Success = &v
	return s
}

type AlignObjectiveResponseBodyData struct {
	// example:
	//
	// 59YD
	AlignId io.Reader `json:"alignId,omitempty" xml:"alignId,omitempty"`
	// example:
	//
	// 5dAX8
	Id io.Reader `json:"id,omitempty" xml:"id,omitempty"`
}

func (s AlignObjectiveResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AlignObjectiveResponseBodyData) GoString() string {
	return s.String()
}

func (s *AlignObjectiveResponseBodyData) SetAlignId(v io.Reader) *AlignObjectiveResponseBodyData {
	s.AlignId = v
	return s
}

func (s *AlignObjectiveResponseBodyData) SetId(v io.Reader) *AlignObjectiveResponseBodyData {
	s.Id = v
	return s
}

type AlignObjectiveResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AlignObjectiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AlignObjectiveResponse) String() string {
	return tea.Prettify(s)
}

func (s AlignObjectiveResponse) GoString() string {
	return s.String()
}

func (s *AlignObjectiveResponse) SetHeaders(v map[string]*string) *AlignObjectiveResponse {
	s.Headers = v
	return s
}

func (s *AlignObjectiveResponse) SetStatusCode(v int32) *AlignObjectiveResponse {
	s.StatusCode = &v
	return s
}

func (s *AlignObjectiveResponse) SetBody(v *AlignObjectiveResponseBody) *AlignObjectiveResponse {
	s.Body = v
	return s
}

type BatchAddPermissionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchAddPermissionHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchAddPermissionHeaders) GoString() string {
	return s.String()
}

func (s *BatchAddPermissionHeaders) SetCommonHeaders(v map[string]*string) *BatchAddPermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchAddPermissionHeaders) SetXAcsDingtalkAccessToken(v string) *BatchAddPermissionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchAddPermissionRequest struct {
	// This parameter is required.
	List []*BatchAddPermissionRequestList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// This parameter is required.
	TargetId *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
	// This parameter is required.
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 当前用户 userId。
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchAddPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchAddPermissionRequest) GoString() string {
	return s.String()
}

func (s *BatchAddPermissionRequest) SetList(v []*BatchAddPermissionRequestList) *BatchAddPermissionRequest {
	s.List = v
	return s
}

func (s *BatchAddPermissionRequest) SetTargetId(v string) *BatchAddPermissionRequest {
	s.TargetId = &v
	return s
}

func (s *BatchAddPermissionRequest) SetTargetType(v string) *BatchAddPermissionRequest {
	s.TargetType = &v
	return s
}

func (s *BatchAddPermissionRequest) SetUserId(v string) *BatchAddPermissionRequest {
	s.UserId = &v
	return s
}

type BatchAddPermissionRequestList struct {
	// This parameter is required.
	Member *BatchAddPermissionRequestListMember `json:"member,omitempty" xml:"member,omitempty" type:"Struct"`
	// This parameter is required.
	PolicyType *int64 `json:"policyType,omitempty" xml:"policyType,omitempty"`
}

func (s BatchAddPermissionRequestList) String() string {
	return tea.Prettify(s)
}

func (s BatchAddPermissionRequestList) GoString() string {
	return s.String()
}

func (s *BatchAddPermissionRequestList) SetMember(v *BatchAddPermissionRequestListMember) *BatchAddPermissionRequestList {
	s.Member = v
	return s
}

func (s *BatchAddPermissionRequestList) SetPolicyType(v int64) *BatchAddPermissionRequestList {
	s.PolicyType = &v
	return s
}

type BatchAddPermissionRequestListMember struct {
	// This parameter is required.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s BatchAddPermissionRequestListMember) String() string {
	return tea.Prettify(s)
}

func (s BatchAddPermissionRequestListMember) GoString() string {
	return s.String()
}

func (s *BatchAddPermissionRequestListMember) SetId(v string) *BatchAddPermissionRequestListMember {
	s.Id = &v
	return s
}

func (s *BatchAddPermissionRequestListMember) SetType(v string) *BatchAddPermissionRequestListMember {
	s.Type = &v
	return s
}

type BatchAddPermissionResponseBody struct {
	// This parameter is required.
	Data *BatchAddPermissionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchAddPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchAddPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *BatchAddPermissionResponseBody) SetData(v *BatchAddPermissionResponseBodyData) *BatchAddPermissionResponseBody {
	s.Data = v
	return s
}

func (s *BatchAddPermissionResponseBody) SetSuccess(v bool) *BatchAddPermissionResponseBody {
	s.Success = &v
	return s
}

type BatchAddPermissionResponseBodyData struct {
	// This parameter is required.
	HasInvalidUser *bool `json:"hasInvalidUser,omitempty" xml:"hasInvalidUser,omitempty"`
	// This parameter is required.
	PermissionTree *BatchAddPermissionResponseBodyDataPermissionTree `json:"permissionTree,omitempty" xml:"permissionTree,omitempty" type:"Struct"`
}

func (s BatchAddPermissionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchAddPermissionResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchAddPermissionResponseBodyData) SetHasInvalidUser(v bool) *BatchAddPermissionResponseBodyData {
	s.HasInvalidUser = &v
	return s
}

func (s *BatchAddPermissionResponseBodyData) SetPermissionTree(v *BatchAddPermissionResponseBodyDataPermissionTree) *BatchAddPermissionResponseBodyData {
	s.PermissionTree = v
	return s
}

type BatchAddPermissionResponseBodyDataPermissionTree struct {
	// This parameter is required.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	PolicyList []*BatchAddPermissionResponseBodyDataPermissionTreePolicyList `json:"policyList,omitempty" xml:"policyList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// public
	Privacy *string `json:"privacy,omitempty" xml:"privacy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// period
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s BatchAddPermissionResponseBodyDataPermissionTree) String() string {
	return tea.Prettify(s)
}

func (s BatchAddPermissionResponseBodyDataPermissionTree) GoString() string {
	return s.String()
}

func (s *BatchAddPermissionResponseBodyDataPermissionTree) SetId(v string) *BatchAddPermissionResponseBodyDataPermissionTree {
	s.Id = &v
	return s
}

func (s *BatchAddPermissionResponseBodyDataPermissionTree) SetPolicyList(v []*BatchAddPermissionResponseBodyDataPermissionTreePolicyList) *BatchAddPermissionResponseBodyDataPermissionTree {
	s.PolicyList = v
	return s
}

func (s *BatchAddPermissionResponseBodyDataPermissionTree) SetPrivacy(v string) *BatchAddPermissionResponseBodyDataPermissionTree {
	s.Privacy = &v
	return s
}

func (s *BatchAddPermissionResponseBodyDataPermissionTree) SetType(v string) *BatchAddPermissionResponseBodyDataPermissionTree {
	s.Type = &v
	return s
}

type BatchAddPermissionResponseBodyDataPermissionTreePolicyList struct {
	// This parameter is required.
	MemberList []*BatchAddPermissionResponseBodyDataPermissionTreePolicyListMemberList `json:"memberList,omitempty" xml:"memberList,omitempty" type:"Repeated"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s BatchAddPermissionResponseBodyDataPermissionTreePolicyList) String() string {
	return tea.Prettify(s)
}

func (s BatchAddPermissionResponseBodyDataPermissionTreePolicyList) GoString() string {
	return s.String()
}

func (s *BatchAddPermissionResponseBodyDataPermissionTreePolicyList) SetMemberList(v []*BatchAddPermissionResponseBodyDataPermissionTreePolicyListMemberList) *BatchAddPermissionResponseBodyDataPermissionTreePolicyList {
	s.MemberList = v
	return s
}

func (s *BatchAddPermissionResponseBodyDataPermissionTreePolicyList) SetName(v string) *BatchAddPermissionResponseBodyDataPermissionTreePolicyList {
	s.Name = &v
	return s
}

func (s *BatchAddPermissionResponseBodyDataPermissionTreePolicyList) SetType(v int64) *BatchAddPermissionResponseBodyDataPermissionTreePolicyList {
	s.Type = &v
	return s
}

type BatchAddPermissionResponseBodyDataPermissionTreePolicyListMemberList struct {
	// This parameter is required.
	Id       *string `json:"id,omitempty" xml:"id,omitempty"`
	Nickname *string `json:"nickname,omitempty" xml:"nickname,omitempty"`
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s BatchAddPermissionResponseBodyDataPermissionTreePolicyListMemberList) String() string {
	return tea.Prettify(s)
}

func (s BatchAddPermissionResponseBodyDataPermissionTreePolicyListMemberList) GoString() string {
	return s.String()
}

func (s *BatchAddPermissionResponseBodyDataPermissionTreePolicyListMemberList) SetId(v string) *BatchAddPermissionResponseBodyDataPermissionTreePolicyListMemberList {
	s.Id = &v
	return s
}

func (s *BatchAddPermissionResponseBodyDataPermissionTreePolicyListMemberList) SetNickname(v string) *BatchAddPermissionResponseBodyDataPermissionTreePolicyListMemberList {
	s.Nickname = &v
	return s
}

func (s *BatchAddPermissionResponseBodyDataPermissionTreePolicyListMemberList) SetType(v string) *BatchAddPermissionResponseBodyDataPermissionTreePolicyListMemberList {
	s.Type = &v
	return s
}

type BatchAddPermissionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchAddPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchAddPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchAddPermissionResponse) GoString() string {
	return s.String()
}

func (s *BatchAddPermissionResponse) SetHeaders(v map[string]*string) *BatchAddPermissionResponse {
	s.Headers = v
	return s
}

func (s *BatchAddPermissionResponse) SetStatusCode(v int32) *BatchAddPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchAddPermissionResponse) SetBody(v *BatchAddPermissionResponseBody) *BatchAddPermissionResponse {
	s.Body = v
	return s
}

type BatchQueryObjectiveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchQueryObjectiveHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryObjectiveHeaders) GoString() string {
	return s.String()
}

func (s *BatchQueryObjectiveHeaders) SetCommonHeaders(v map[string]*string) *BatchQueryObjectiveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchQueryObjectiveHeaders) SetXAcsDingtalkAccessToken(v string) *BatchQueryObjectiveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchQueryObjectiveRequest struct {
	// This parameter is required.
	ObjectiveIds []*string `json:"objectiveIds,omitempty" xml:"objectiveIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 10056
	PeriodId *string `json:"periodId,omitempty" xml:"periodId,omitempty"`
	// example:
	//
	// false
	WithAlign *bool `json:"withAlign,omitempty" xml:"withAlign,omitempty"`
	// example:
	//
	// false
	WithKr *bool `json:"withKr,omitempty" xml:"withKr,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	WithProgress *bool `json:"withProgress,omitempty" xml:"withProgress,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0115396701752283
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchQueryObjectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryObjectiveRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryObjectiveRequest) SetObjectiveIds(v []*string) *BatchQueryObjectiveRequest {
	s.ObjectiveIds = v
	return s
}

func (s *BatchQueryObjectiveRequest) SetPeriodId(v string) *BatchQueryObjectiveRequest {
	s.PeriodId = &v
	return s
}

func (s *BatchQueryObjectiveRequest) SetWithAlign(v bool) *BatchQueryObjectiveRequest {
	s.WithAlign = &v
	return s
}

func (s *BatchQueryObjectiveRequest) SetWithKr(v bool) *BatchQueryObjectiveRequest {
	s.WithKr = &v
	return s
}

func (s *BatchQueryObjectiveRequest) SetWithProgress(v bool) *BatchQueryObjectiveRequest {
	s.WithProgress = &v
	return s
}

func (s *BatchQueryObjectiveRequest) SetUserId(v string) *BatchQueryObjectiveRequest {
	s.UserId = &v
	return s
}

type BatchQueryObjectiveResponseBody struct {
	Data []*BatchQueryObjectiveResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchQueryObjectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryObjectiveResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQueryObjectiveResponseBody) SetData(v []*BatchQueryObjectiveResponseBodyData) *BatchQueryObjectiveResponseBody {
	s.Data = v
	return s
}

func (s *BatchQueryObjectiveResponseBody) SetSuccess(v bool) *BatchQueryObjectiveResponseBody {
	s.Success = &v
	return s
}

type BatchQueryObjectiveResponseBodyData struct {
	AlignFromIds []io.Reader `json:"alignFromIds,omitempty" xml:"alignFromIds,omitempty" type:"Repeated"`
	AlignToIds   []io.Reader `json:"alignToIds,omitempty" xml:"alignToIds,omitempty" type:"Repeated"`
	// example:
	//
	// Objective demo
	Content io.Reader `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 1648625407694
	GmtCreate *float32 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1648625407694
	GmtModified *float32 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 5dAX8
	Id     io.Reader                                    `json:"id,omitempty" xml:"id,omitempty"`
	KrList []*BatchQueryObjectiveResponseBodyDataKrList `json:"krList,omitempty" xml:"krList,omitempty" type:"Repeated"`
	Owner  *BatchQueryObjectiveResponseBodyDataOwner    `json:"owner,omitempty" xml:"owner,omitempty" type:"Struct"`
	// example:
	//
	// 1006
	PeriodId   io.Reader  `json:"periodId,omitempty" xml:"periodId,omitempty"`
	Permission []*float32 `json:"permission,omitempty" xml:"permission,omitempty" type:"Repeated"`
	// example:
	//
	// 3021332
	Position *int32                                       `json:"position,omitempty" xml:"position,omitempty"`
	Progress *BatchQueryObjectiveResponseBodyDataProgress `json:"progress,omitempty" xml:"progress,omitempty" type:"Struct"`
	// example:
	//
	// 100
	ProgressPercent *float32 `json:"progressPercent,omitempty" xml:"progressPercent,omitempty"`
	// example:
	//
	// true
	Published *bool `json:"published,omitempty" xml:"published,omitempty"`
	// example:
	//
	// 20
	Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// s34d
	UserId io.Reader `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// 50
	Weight *float32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s BatchQueryObjectiveResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryObjectiveResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchQueryObjectiveResponseBodyData) SetAlignFromIds(v []io.Reader) *BatchQueryObjectiveResponseBodyData {
	s.AlignFromIds = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetAlignToIds(v []io.Reader) *BatchQueryObjectiveResponseBodyData {
	s.AlignToIds = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetContent(v io.Reader) *BatchQueryObjectiveResponseBodyData {
	s.Content = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetGmtCreate(v float32) *BatchQueryObjectiveResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetGmtModified(v float32) *BatchQueryObjectiveResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetId(v io.Reader) *BatchQueryObjectiveResponseBodyData {
	s.Id = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetKrList(v []*BatchQueryObjectiveResponseBodyDataKrList) *BatchQueryObjectiveResponseBodyData {
	s.KrList = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetOwner(v *BatchQueryObjectiveResponseBodyDataOwner) *BatchQueryObjectiveResponseBodyData {
	s.Owner = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetPeriodId(v io.Reader) *BatchQueryObjectiveResponseBodyData {
	s.PeriodId = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetPermission(v []*float32) *BatchQueryObjectiveResponseBodyData {
	s.Permission = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetPosition(v int32) *BatchQueryObjectiveResponseBodyData {
	s.Position = &v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetProgress(v *BatchQueryObjectiveResponseBodyDataProgress) *BatchQueryObjectiveResponseBodyData {
	s.Progress = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetProgressPercent(v float32) *BatchQueryObjectiveResponseBodyData {
	s.ProgressPercent = &v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetPublished(v bool) *BatchQueryObjectiveResponseBodyData {
	s.Published = &v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetScore(v float32) *BatchQueryObjectiveResponseBodyData {
	s.Score = &v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetStatus(v int32) *BatchQueryObjectiveResponseBodyData {
	s.Status = &v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetUserId(v io.Reader) *BatchQueryObjectiveResponseBodyData {
	s.UserId = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetWeight(v float32) *BatchQueryObjectiveResponseBodyData {
	s.Weight = &v
	return s
}

type BatchQueryObjectiveResponseBodyDataKrList struct {
	// example:
	//
	// 你好
	Content io.Reader `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 1648625407694
	GmtCreate *float32 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1648625407694
	GmtModified *float32 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5w9f
	Id io.Reader `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 5wf8
	ObjectiveId io.Reader  `json:"objectiveId,omitempty" xml:"objectiveId,omitempty"`
	Permission  []*float32 `json:"permission,omitempty" xml:"permission,omitempty" type:"Repeated"`
	// example:
	//
	// 35614536
	Position *int64                                             `json:"position,omitempty" xml:"position,omitempty"`
	Progress *BatchQueryObjectiveResponseBodyDataKrListProgress `json:"progress,omitempty" xml:"progress,omitempty" type:"Struct"`
	// example:
	//
	// 44
	Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
	// example:
	//
	// 44
	Weight *float32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s BatchQueryObjectiveResponseBodyDataKrList) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryObjectiveResponseBodyDataKrList) GoString() string {
	return s.String()
}

func (s *BatchQueryObjectiveResponseBodyDataKrList) SetContent(v io.Reader) *BatchQueryObjectiveResponseBodyDataKrList {
	s.Content = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataKrList) SetGmtCreate(v float32) *BatchQueryObjectiveResponseBodyDataKrList {
	s.GmtCreate = &v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataKrList) SetGmtModified(v float32) *BatchQueryObjectiveResponseBodyDataKrList {
	s.GmtModified = &v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataKrList) SetId(v io.Reader) *BatchQueryObjectiveResponseBodyDataKrList {
	s.Id = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataKrList) SetObjectiveId(v io.Reader) *BatchQueryObjectiveResponseBodyDataKrList {
	s.ObjectiveId = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataKrList) SetPermission(v []*float32) *BatchQueryObjectiveResponseBodyDataKrList {
	s.Permission = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataKrList) SetPosition(v int64) *BatchQueryObjectiveResponseBodyDataKrList {
	s.Position = &v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataKrList) SetProgress(v *BatchQueryObjectiveResponseBodyDataKrListProgress) *BatchQueryObjectiveResponseBodyDataKrList {
	s.Progress = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataKrList) SetScore(v float32) *BatchQueryObjectiveResponseBodyDataKrList {
	s.Score = &v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataKrList) SetWeight(v float32) *BatchQueryObjectiveResponseBodyDataKrList {
	s.Weight = &v
	return s
}

type BatchQueryObjectiveResponseBodyDataKrListProgress struct {
	// example:
	//
	// 30
	Percent *int32 `json:"percent,omitempty" xml:"percent,omitempty"`
}

func (s BatchQueryObjectiveResponseBodyDataKrListProgress) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryObjectiveResponseBodyDataKrListProgress) GoString() string {
	return s.String()
}

func (s *BatchQueryObjectiveResponseBodyDataKrListProgress) SetPercent(v int32) *BatchQueryObjectiveResponseBodyDataKrListProgress {
	s.Percent = &v
	return s
}

type BatchQueryObjectiveResponseBodyDataOwner struct {
	// example:
	//
	// @lADPDh0cQ_j4Mi_NBULNBUA
	AvatarMediaId io.Reader `json:"avatarMediaId,omitempty" xml:"avatarMediaId,omitempty"`
	// example:
	//
	// ding4d1c8883ff63ee8124f2f5cc6abecb85
	CorpId io.Reader `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// K1AMgq
	Id io.Reader `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 你好
	Nickname io.Reader `json:"nickname,omitempty" xml:"nickname,omitempty"`
	// example:
	//
	// 06186238011033616
	UserId io.Reader `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchQueryObjectiveResponseBodyDataOwner) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryObjectiveResponseBodyDataOwner) GoString() string {
	return s.String()
}

func (s *BatchQueryObjectiveResponseBodyDataOwner) SetAvatarMediaId(v io.Reader) *BatchQueryObjectiveResponseBodyDataOwner {
	s.AvatarMediaId = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataOwner) SetCorpId(v io.Reader) *BatchQueryObjectiveResponseBodyDataOwner {
	s.CorpId = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataOwner) SetId(v io.Reader) *BatchQueryObjectiveResponseBodyDataOwner {
	s.Id = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataOwner) SetNickname(v io.Reader) *BatchQueryObjectiveResponseBodyDataOwner {
	s.Nickname = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataOwner) SetUserId(v io.Reader) *BatchQueryObjectiveResponseBodyDataOwner {
	s.UserId = v
	return s
}

type BatchQueryObjectiveResponseBodyDataProgress struct {
	// example:
	//
	// 100
	Percent *int32 `json:"percent,omitempty" xml:"percent,omitempty"`
}

func (s BatchQueryObjectiveResponseBodyDataProgress) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryObjectiveResponseBodyDataProgress) GoString() string {
	return s.String()
}

func (s *BatchQueryObjectiveResponseBodyDataProgress) SetPercent(v int32) *BatchQueryObjectiveResponseBodyDataProgress {
	s.Percent = &v
	return s
}

type BatchQueryObjectiveResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchQueryObjectiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchQueryObjectiveResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryObjectiveResponse) GoString() string {
	return s.String()
}

func (s *BatchQueryObjectiveResponse) SetHeaders(v map[string]*string) *BatchQueryObjectiveResponse {
	s.Headers = v
	return s
}

func (s *BatchQueryObjectiveResponse) SetStatusCode(v int32) *BatchQueryObjectiveResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchQueryObjectiveResponse) SetBody(v *BatchQueryObjectiveResponseBody) *BatchQueryObjectiveResponse {
	s.Body = v
	return s
}

type BatchQueryUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchQueryUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryUserHeaders) GoString() string {
	return s.String()
}

func (s *BatchQueryUserHeaders) SetCommonHeaders(v map[string]*string) *BatchQueryUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchQueryUserHeaders) SetXAcsDingtalkAccessToken(v string) *BatchQueryUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchQueryUserRequest struct {
	OkrUserIds []*string `json:"okrUserIds,omitempty" xml:"okrUserIds,omitempty" type:"Repeated"`
	UserIds    []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s BatchQueryUserRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryUserRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryUserRequest) SetOkrUserIds(v []*string) *BatchQueryUserRequest {
	s.OkrUserIds = v
	return s
}

func (s *BatchQueryUserRequest) SetUserIds(v []*string) *BatchQueryUserRequest {
	s.UserIds = v
	return s
}

type BatchQueryUserResponseBody struct {
	// This parameter is required.
	Data []*BatchQueryUserResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchQueryUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryUserResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQueryUserResponseBody) SetData(v []*BatchQueryUserResponseBodyData) *BatchQueryUserResponseBody {
	s.Data = v
	return s
}

func (s *BatchQueryUserResponseBody) SetSuccess(v bool) *BatchQueryUserResponseBody {
	s.Success = &v
	return s
}

type BatchQueryUserResponseBodyData struct {
	// example:
	//
	// @lADPDh0cQ_j4Mi_NBULNBUA
	AvatarMediaId io.Reader `json:"avatarMediaId,omitempty" xml:"avatarMediaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://static.dingtalk.com/media/lADPEj_RiGhUdKjNC9TNC9A_3024_3028.jpg_620x10000q90.jpg
	AvatarUrl io.Reader `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// ding4d1c8883ff63ee8124f2f5cc6abecb85
	CorpId io.Reader `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// K1AMgq
	Id io.Reader `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 你好
	Nickname io.Reader `json:"nickname,omitempty" xml:"nickname,omitempty"`
	// example:
	//
	// 06186238011033616
	UserId io.Reader `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchQueryUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchQueryUserResponseBodyData) SetAvatarMediaId(v io.Reader) *BatchQueryUserResponseBodyData {
	s.AvatarMediaId = v
	return s
}

func (s *BatchQueryUserResponseBodyData) SetAvatarUrl(v io.Reader) *BatchQueryUserResponseBodyData {
	s.AvatarUrl = v
	return s
}

func (s *BatchQueryUserResponseBodyData) SetCorpId(v io.Reader) *BatchQueryUserResponseBodyData {
	s.CorpId = v
	return s
}

func (s *BatchQueryUserResponseBodyData) SetId(v io.Reader) *BatchQueryUserResponseBodyData {
	s.Id = v
	return s
}

func (s *BatchQueryUserResponseBodyData) SetNickname(v io.Reader) *BatchQueryUserResponseBodyData {
	s.Nickname = v
	return s
}

func (s *BatchQueryUserResponseBodyData) SetUserId(v io.Reader) *BatchQueryUserResponseBodyData {
	s.UserId = v
	return s
}

type BatchQueryUserResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchQueryUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchQueryUserResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryUserResponse) GoString() string {
	return s.String()
}

func (s *BatchQueryUserResponse) SetHeaders(v map[string]*string) *BatchQueryUserResponse {
	s.Headers = v
	return s
}

func (s *BatchQueryUserResponse) SetStatusCode(v int32) *BatchQueryUserResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchQueryUserResponse) SetBody(v *BatchQueryUserResponseBody) *BatchQueryUserResponse {
	s.Body = v
	return s
}

type CreateKeyResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateKeyResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyResultHeaders) GoString() string {
	return s.String()
}

func (s *CreateKeyResultHeaders) SetCommonHeaders(v map[string]*string) *CreateKeyResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateKeyResultHeaders) SetXAcsDingtalkAccessToken(v string) *CreateKeyResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateKeyResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 我的内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 58Y4
	ObjectiveId *string `json:"objectiveId,omitempty" xml:"objectiveId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1006
	PeriodId *string `json:"periodId,omitempty" xml:"periodId,omitempty"`
	// example:
	//
	// 234631
	PrevPosition *int64 `json:"prevPosition,omitempty" xml:"prevPosition,omitempty"`
	// example:
	//
	// 100
	Weight *int64 `json:"weight,omitempty" xml:"weight,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 06186238011033616
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateKeyResultRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyResultRequest) GoString() string {
	return s.String()
}

func (s *CreateKeyResultRequest) SetContent(v string) *CreateKeyResultRequest {
	s.Content = &v
	return s
}

func (s *CreateKeyResultRequest) SetObjectiveId(v string) *CreateKeyResultRequest {
	s.ObjectiveId = &v
	return s
}

func (s *CreateKeyResultRequest) SetPeriodId(v string) *CreateKeyResultRequest {
	s.PeriodId = &v
	return s
}

func (s *CreateKeyResultRequest) SetPrevPosition(v int64) *CreateKeyResultRequest {
	s.PrevPosition = &v
	return s
}

func (s *CreateKeyResultRequest) SetWeight(v int64) *CreateKeyResultRequest {
	s.Weight = &v
	return s
}

func (s *CreateKeyResultRequest) SetUserId(v string) *CreateKeyResultRequest {
	s.UserId = &v
	return s
}

type CreateKeyResultResponseBody struct {
	// This parameter is required.
	Data *CreateKeyResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateKeyResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyResultResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKeyResultResponseBody) SetData(v *CreateKeyResultResponseBodyData) *CreateKeyResultResponseBody {
	s.Data = v
	return s
}

func (s *CreateKeyResultResponseBody) SetSuccess(v bool) *CreateKeyResultResponseBody {
	s.Success = &v
	return s
}

type CreateKeyResultResponseBodyData struct {
	// This parameter is required.
	//
	// example:
	//
	// R45Y
	Id io.Reader `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 420983
	Position *int64 `json:"position,omitempty" xml:"position,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	Weight *int64 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s CreateKeyResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateKeyResultResponseBodyData) SetId(v io.Reader) *CreateKeyResultResponseBodyData {
	s.Id = v
	return s
}

func (s *CreateKeyResultResponseBodyData) SetPosition(v int64) *CreateKeyResultResponseBodyData {
	s.Position = &v
	return s
}

func (s *CreateKeyResultResponseBodyData) SetWeight(v int64) *CreateKeyResultResponseBodyData {
	s.Weight = &v
	return s
}

type CreateKeyResultResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKeyResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKeyResultResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyResultResponse) GoString() string {
	return s.String()
}

func (s *CreateKeyResultResponse) SetHeaders(v map[string]*string) *CreateKeyResultResponse {
	s.Headers = v
	return s
}

func (s *CreateKeyResultResponse) SetStatusCode(v int32) *CreateKeyResultResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKeyResultResponse) SetBody(v *CreateKeyResultResponseBody) *CreateKeyResultResponse {
	s.Body = v
	return s
}

type CreateObjectiveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateObjectiveHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateObjectiveHeaders) GoString() string {
	return s.String()
}

func (s *CreateObjectiveHeaders) SetCommonHeaders(v map[string]*string) *CreateObjectiveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateObjectiveHeaders) SetXAcsDingtalkAccessToken(v string) *CreateObjectiveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateObjectiveRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 我是内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1006
	PeriodId *string `json:"periodId,omitempty" xml:"periodId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1334543
	PrevPosition *string `json:"prevPosition,omitempty" xml:"prevPosition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 06186238011033616
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateObjectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateObjectiveRequest) GoString() string {
	return s.String()
}

func (s *CreateObjectiveRequest) SetContent(v string) *CreateObjectiveRequest {
	s.Content = &v
	return s
}

func (s *CreateObjectiveRequest) SetPeriodId(v string) *CreateObjectiveRequest {
	s.PeriodId = &v
	return s
}

func (s *CreateObjectiveRequest) SetPrevPosition(v string) *CreateObjectiveRequest {
	s.PrevPosition = &v
	return s
}

func (s *CreateObjectiveRequest) SetUserId(v string) *CreateObjectiveRequest {
	s.UserId = &v
	return s
}

type CreateObjectiveResponseBody struct {
	Data *CreateObjectiveResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateObjectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateObjectiveResponseBody) GoString() string {
	return s.String()
}

func (s *CreateObjectiveResponseBody) SetData(v *CreateObjectiveResponseBodyData) *CreateObjectiveResponseBody {
	s.Data = v
	return s
}

func (s *CreateObjectiveResponseBody) SetSuccess(v bool) *CreateObjectiveResponseBody {
	s.Success = &v
	return s
}

type CreateObjectiveResponseBodyData struct {
	// This parameter is required.
	//
	// example:
	//
	// 58YD
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 33453
	Position *string `json:"position,omitempty" xml:"position,omitempty"`
}

func (s CreateObjectiveResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateObjectiveResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateObjectiveResponseBodyData) SetId(v string) *CreateObjectiveResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateObjectiveResponseBodyData) SetPosition(v string) *CreateObjectiveResponseBodyData {
	s.Position = &v
	return s
}

type CreateObjectiveResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateObjectiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateObjectiveResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateObjectiveResponse) GoString() string {
	return s.String()
}

func (s *CreateObjectiveResponse) SetHeaders(v map[string]*string) *CreateObjectiveResponse {
	s.Headers = v
	return s
}

func (s *CreateObjectiveResponse) SetStatusCode(v int32) *CreateObjectiveResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateObjectiveResponse) SetBody(v *CreateObjectiveResponseBody) *CreateObjectiveResponse {
	s.Body = v
	return s
}

type DeleteKeyResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteKeyResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyResultHeaders) GoString() string {
	return s.String()
}

func (s *DeleteKeyResultHeaders) SetCommonHeaders(v map[string]*string) *DeleteKeyResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteKeyResultHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteKeyResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteKeyResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4d2d
	KrId *string `json:"krId,omitempty" xml:"krId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 06186238011033616
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeleteKeyResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyResultRequest) GoString() string {
	return s.String()
}

func (s *DeleteKeyResultRequest) SetKrId(v string) *DeleteKeyResultRequest {
	s.KrId = &v
	return s
}

func (s *DeleteKeyResultRequest) SetUserId(v string) *DeleteKeyResultRequest {
	s.UserId = &v
	return s
}

type DeleteKeyResultResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// success
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteKeyResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyResultResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKeyResultResponseBody) SetData(v bool) *DeleteKeyResultResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteKeyResultResponseBody) SetSuccess(v bool) *DeleteKeyResultResponseBody {
	s.Success = &v
	return s
}

type DeleteKeyResultResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKeyResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKeyResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyResultResponse) GoString() string {
	return s.String()
}

func (s *DeleteKeyResultResponse) SetHeaders(v map[string]*string) *DeleteKeyResultResponse {
	s.Headers = v
	return s
}

func (s *DeleteKeyResultResponse) SetStatusCode(v int32) *DeleteKeyResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKeyResultResponse) SetBody(v *DeleteKeyResultResponseBody) *DeleteKeyResultResponse {
	s.Body = v
	return s
}

type DeleteObjectiveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteObjectiveHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteObjectiveHeaders) GoString() string {
	return s.String()
}

func (s *DeleteObjectiveHeaders) SetCommonHeaders(v map[string]*string) *DeleteObjectiveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteObjectiveHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteObjectiveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteObjectiveRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 06186238011033616
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeleteObjectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteObjectiveRequest) GoString() string {
	return s.String()
}

func (s *DeleteObjectiveRequest) SetUserId(v string) *DeleteObjectiveRequest {
	s.UserId = &v
	return s
}

type DeleteObjectiveResponseBody struct {
	Data *DeleteObjectiveResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteObjectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteObjectiveResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteObjectiveResponseBody) SetData(v *DeleteObjectiveResponseBodyData) *DeleteObjectiveResponseBody {
	s.Data = v
	return s
}

func (s *DeleteObjectiveResponseBody) SetSuccess(v bool) *DeleteObjectiveResponseBody {
	s.Success = &v
	return s
}

type DeleteObjectiveResponseBodyData struct {
	// This parameter is required.
	//
	// example:
	//
	// 58YD
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s DeleteObjectiveResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteObjectiveResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteObjectiveResponseBodyData) SetId(v string) *DeleteObjectiveResponseBodyData {
	s.Id = &v
	return s
}

type DeleteObjectiveResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteObjectiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteObjectiveResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteObjectiveResponse) GoString() string {
	return s.String()
}

func (s *DeleteObjectiveResponse) SetHeaders(v map[string]*string) *DeleteObjectiveResponse {
	s.Headers = v
	return s
}

func (s *DeleteObjectiveResponse) SetStatusCode(v int32) *DeleteObjectiveResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteObjectiveResponse) SetBody(v *DeleteObjectiveResponseBody) *DeleteObjectiveResponse {
	s.Body = v
	return s
}

type DeletePermissionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeletePermissionHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeletePermissionHeaders) GoString() string {
	return s.String()
}

func (s *DeletePermissionHeaders) SetCommonHeaders(v map[string]*string) *DeletePermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeletePermissionHeaders) SetXAcsDingtalkAccessToken(v string) *DeletePermissionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeletePermissionRequest struct {
	// This parameter is required.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	PolicyType *int64 `json:"policyType,omitempty" xml:"policyType,omitempty"`
	// This parameter is required.
	TargetId *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
	// This parameter is required.
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0115396701752283
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeletePermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePermissionRequest) GoString() string {
	return s.String()
}

func (s *DeletePermissionRequest) SetId(v string) *DeletePermissionRequest {
	s.Id = &v
	return s
}

func (s *DeletePermissionRequest) SetPolicyType(v int64) *DeletePermissionRequest {
	s.PolicyType = &v
	return s
}

func (s *DeletePermissionRequest) SetTargetId(v string) *DeletePermissionRequest {
	s.TargetId = &v
	return s
}

func (s *DeletePermissionRequest) SetTargetType(v string) *DeletePermissionRequest {
	s.TargetType = &v
	return s
}

func (s *DeletePermissionRequest) SetType(v string) *DeletePermissionRequest {
	s.Type = &v
	return s
}

func (s *DeletePermissionRequest) SetUserId(v string) *DeletePermissionRequest {
	s.UserId = &v
	return s
}

type DeletePermissionResponseBody struct {
	// This parameter is required.
	Data *DeletePermissionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeletePermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePermissionResponseBody) SetData(v *DeletePermissionResponseBodyData) *DeletePermissionResponseBody {
	s.Data = v
	return s
}

func (s *DeletePermissionResponseBody) SetSuccess(v bool) *DeletePermissionResponseBody {
	s.Success = &v
	return s
}

type DeletePermissionResponseBodyData struct {
	// This parameter is required.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	PolicyList []*DeletePermissionResponseBodyDataPolicyList `json:"policyList,omitempty" xml:"policyList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// public
	Privacy *string `json:"privacy,omitempty" xml:"privacy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// period
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DeletePermissionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeletePermissionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeletePermissionResponseBodyData) SetId(v string) *DeletePermissionResponseBodyData {
	s.Id = &v
	return s
}

func (s *DeletePermissionResponseBodyData) SetPolicyList(v []*DeletePermissionResponseBodyDataPolicyList) *DeletePermissionResponseBodyData {
	s.PolicyList = v
	return s
}

func (s *DeletePermissionResponseBodyData) SetPrivacy(v string) *DeletePermissionResponseBodyData {
	s.Privacy = &v
	return s
}

func (s *DeletePermissionResponseBodyData) SetType(v string) *DeletePermissionResponseBodyData {
	s.Type = &v
	return s
}

type DeletePermissionResponseBodyDataPolicyList struct {
	// This parameter is required.
	MemberList []*DeletePermissionResponseBodyDataPolicyListMemberList `json:"memberList,omitempty" xml:"memberList,omitempty" type:"Repeated"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DeletePermissionResponseBodyDataPolicyList) String() string {
	return tea.Prettify(s)
}

func (s DeletePermissionResponseBodyDataPolicyList) GoString() string {
	return s.String()
}

func (s *DeletePermissionResponseBodyDataPolicyList) SetMemberList(v []*DeletePermissionResponseBodyDataPolicyListMemberList) *DeletePermissionResponseBodyDataPolicyList {
	s.MemberList = v
	return s
}

func (s *DeletePermissionResponseBodyDataPolicyList) SetName(v string) *DeletePermissionResponseBodyDataPolicyList {
	s.Name = &v
	return s
}

func (s *DeletePermissionResponseBodyDataPolicyList) SetType(v int64) *DeletePermissionResponseBodyDataPolicyList {
	s.Type = &v
	return s
}

type DeletePermissionResponseBodyDataPolicyListMemberList struct {
	// This parameter is required.
	Id       *string `json:"id,omitempty" xml:"id,omitempty"`
	Nickname *string `json:"nickname,omitempty" xml:"nickname,omitempty"`
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DeletePermissionResponseBodyDataPolicyListMemberList) String() string {
	return tea.Prettify(s)
}

func (s DeletePermissionResponseBodyDataPolicyListMemberList) GoString() string {
	return s.String()
}

func (s *DeletePermissionResponseBodyDataPolicyListMemberList) SetId(v string) *DeletePermissionResponseBodyDataPolicyListMemberList {
	s.Id = &v
	return s
}

func (s *DeletePermissionResponseBodyDataPolicyListMemberList) SetNickname(v string) *DeletePermissionResponseBodyDataPolicyListMemberList {
	s.Nickname = &v
	return s
}

func (s *DeletePermissionResponseBodyDataPolicyListMemberList) SetType(v string) *DeletePermissionResponseBodyDataPolicyListMemberList {
	s.Type = &v
	return s
}

type DeletePermissionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePermissionResponse) GoString() string {
	return s.String()
}

func (s *DeletePermissionResponse) SetHeaders(v map[string]*string) *DeletePermissionResponse {
	s.Headers = v
	return s
}

func (s *DeletePermissionResponse) SetStatusCode(v int32) *DeletePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePermissionResponse) SetBody(v *DeletePermissionResponseBody) *DeletePermissionResponse {
	s.Body = v
	return s
}

type GetPeriodListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPeriodListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPeriodListHeaders) GoString() string {
	return s.String()
}

func (s *GetPeriodListHeaders) SetCommonHeaders(v map[string]*string) *GetPeriodListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPeriodListHeaders) SetXAcsDingtalkAccessToken(v string) *GetPeriodListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPeriodListResponseBody struct {
	Data    *GetPeriodListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Success *bool                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetPeriodListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPeriodListResponseBody) GoString() string {
	return s.String()
}

func (s *GetPeriodListResponseBody) SetData(v *GetPeriodListResponseBodyData) *GetPeriodListResponseBody {
	s.Data = v
	return s
}

func (s *GetPeriodListResponseBody) SetSuccess(v bool) *GetPeriodListResponseBody {
	s.Success = &v
	return s
}

type GetPeriodListResponseBodyData struct {
	// This parameter is required.
	PeriodList []*GetPeriodListResponseBodyDataPeriodList `json:"periodList,omitempty" xml:"periodList,omitempty" type:"Repeated"`
}

func (s GetPeriodListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPeriodListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPeriodListResponseBodyData) SetPeriodList(v []*GetPeriodListResponseBodyDataPeriodList) *GetPeriodListResponseBodyData {
	s.PeriodList = v
	return s
}

type GetPeriodListResponseBodyDataPeriodList struct {
	// This parameter is required.
	EndTime *float32 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	Id io.Reader `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	IsCurrent *bool `json:"isCurrent,omitempty" xml:"isCurrent,omitempty"`
	// This parameter is required.
	IsYearly *bool `json:"isYearly,omitempty" xml:"isYearly,omitempty"`
	// This parameter is required.
	Name io.Reader `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	StartTime *float32 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetPeriodListResponseBodyDataPeriodList) String() string {
	return tea.Prettify(s)
}

func (s GetPeriodListResponseBodyDataPeriodList) GoString() string {
	return s.String()
}

func (s *GetPeriodListResponseBodyDataPeriodList) SetEndTime(v float32) *GetPeriodListResponseBodyDataPeriodList {
	s.EndTime = &v
	return s
}

func (s *GetPeriodListResponseBodyDataPeriodList) SetId(v io.Reader) *GetPeriodListResponseBodyDataPeriodList {
	s.Id = v
	return s
}

func (s *GetPeriodListResponseBodyDataPeriodList) SetIsCurrent(v bool) *GetPeriodListResponseBodyDataPeriodList {
	s.IsCurrent = &v
	return s
}

func (s *GetPeriodListResponseBodyDataPeriodList) SetIsYearly(v bool) *GetPeriodListResponseBodyDataPeriodList {
	s.IsYearly = &v
	return s
}

func (s *GetPeriodListResponseBodyDataPeriodList) SetName(v io.Reader) *GetPeriodListResponseBodyDataPeriodList {
	s.Name = v
	return s
}

func (s *GetPeriodListResponseBodyDataPeriodList) SetStartTime(v float32) *GetPeriodListResponseBodyDataPeriodList {
	s.StartTime = &v
	return s
}

type GetPeriodListResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPeriodListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPeriodListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPeriodListResponse) GoString() string {
	return s.String()
}

func (s *GetPeriodListResponse) SetHeaders(v map[string]*string) *GetPeriodListResponse {
	s.Headers = v
	return s
}

func (s *GetPeriodListResponse) SetStatusCode(v int32) *GetPeriodListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPeriodListResponse) SetBody(v *GetPeriodListResponseBody) *GetPeriodListResponse {
	s.Body = v
	return s
}

type GetPermissionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPermissionHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionHeaders) GoString() string {
	return s.String()
}

func (s *GetPermissionHeaders) SetCommonHeaders(v map[string]*string) *GetPermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPermissionHeaders) SetXAcsDingtalkAccessToken(v string) *GetPermissionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPermissionRequest struct {
	// This parameter is required.
	TargetId *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
	// This parameter is required.
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 当前用户 userId。
	UserId        *string `json:"userId,omitempty" xml:"userId,omitempty"`
	WithKr        *bool   `json:"withKr,omitempty" xml:"withKr,omitempty"`
	WithObjective *bool   `json:"withObjective,omitempty" xml:"withObjective,omitempty"`
}

func (s GetPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionRequest) GoString() string {
	return s.String()
}

func (s *GetPermissionRequest) SetTargetId(v string) *GetPermissionRequest {
	s.TargetId = &v
	return s
}

func (s *GetPermissionRequest) SetTargetType(v string) *GetPermissionRequest {
	s.TargetType = &v
	return s
}

func (s *GetPermissionRequest) SetUserId(v string) *GetPermissionRequest {
	s.UserId = &v
	return s
}

func (s *GetPermissionRequest) SetWithKr(v bool) *GetPermissionRequest {
	s.WithKr = &v
	return s
}

func (s *GetPermissionRequest) SetWithObjective(v bool) *GetPermissionRequest {
	s.WithObjective = &v
	return s
}

type GetPermissionResponseBody struct {
	// This parameter is required.
	Data *GetPermissionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *GetPermissionResponseBody) SetData(v *GetPermissionResponseBodyData) *GetPermissionResponseBody {
	s.Data = v
	return s
}

func (s *GetPermissionResponseBody) SetSuccess(v bool) *GetPermissionResponseBody {
	s.Success = &v
	return s
}

type GetPermissionResponseBodyData struct {
	// This parameter is required.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	PolicyList []*GetPermissionResponseBodyDataPolicyList `json:"policyList,omitempty" xml:"policyList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// public
	Privacy *string `json:"privacy,omitempty" xml:"privacy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// period
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetPermissionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPermissionResponseBodyData) SetId(v string) *GetPermissionResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetPermissionResponseBodyData) SetPolicyList(v []*GetPermissionResponseBodyDataPolicyList) *GetPermissionResponseBodyData {
	s.PolicyList = v
	return s
}

func (s *GetPermissionResponseBodyData) SetPrivacy(v string) *GetPermissionResponseBodyData {
	s.Privacy = &v
	return s
}

func (s *GetPermissionResponseBodyData) SetType(v string) *GetPermissionResponseBodyData {
	s.Type = &v
	return s
}

type GetPermissionResponseBodyDataPolicyList struct {
	// This parameter is required.
	MemberList []*GetPermissionResponseBodyDataPolicyListMemberList `json:"memberList,omitempty" xml:"memberList,omitempty" type:"Repeated"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetPermissionResponseBodyDataPolicyList) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionResponseBodyDataPolicyList) GoString() string {
	return s.String()
}

func (s *GetPermissionResponseBodyDataPolicyList) SetMemberList(v []*GetPermissionResponseBodyDataPolicyListMemberList) *GetPermissionResponseBodyDataPolicyList {
	s.MemberList = v
	return s
}

func (s *GetPermissionResponseBodyDataPolicyList) SetName(v string) *GetPermissionResponseBodyDataPolicyList {
	s.Name = &v
	return s
}

func (s *GetPermissionResponseBodyDataPolicyList) SetType(v int64) *GetPermissionResponseBodyDataPolicyList {
	s.Type = &v
	return s
}

type GetPermissionResponseBodyDataPolicyListMemberList struct {
	// This parameter is required.
	Id       *string `json:"id,omitempty" xml:"id,omitempty"`
	Nickname *string `json:"nickname,omitempty" xml:"nickname,omitempty"`
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetPermissionResponseBodyDataPolicyListMemberList) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionResponseBodyDataPolicyListMemberList) GoString() string {
	return s.String()
}

func (s *GetPermissionResponseBodyDataPolicyListMemberList) SetId(v string) *GetPermissionResponseBodyDataPolicyListMemberList {
	s.Id = &v
	return s
}

func (s *GetPermissionResponseBodyDataPolicyListMemberList) SetNickname(v string) *GetPermissionResponseBodyDataPolicyListMemberList {
	s.Nickname = &v
	return s
}

func (s *GetPermissionResponseBodyDataPolicyListMemberList) SetType(v string) *GetPermissionResponseBodyDataPolicyListMemberList {
	s.Type = &v
	return s
}

type GetPermissionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionResponse) GoString() string {
	return s.String()
}

func (s *GetPermissionResponse) SetHeaders(v map[string]*string) *GetPermissionResponse {
	s.Headers = v
	return s
}

func (s *GetPermissionResponse) SetStatusCode(v int32) *GetPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPermissionResponse) SetBody(v *GetPermissionResponseBody) *GetPermissionResponse {
	s.Body = v
	return s
}

type GetUserOkrHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserOkrHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserOkrHeaders) GoString() string {
	return s.String()
}

func (s *GetUserOkrHeaders) SetCommonHeaders(v map[string]*string) *GetUserOkrHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserOkrHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserOkrHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserOkrRequest struct {
	// example:
	//
	// 2
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1005
	PeriodId *string `json:"periodId,omitempty" xml:"periodId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 011539670175223
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserOkrRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserOkrRequest) GoString() string {
	return s.String()
}

func (s *GetUserOkrRequest) SetPageNumber(v int64) *GetUserOkrRequest {
	s.PageNumber = &v
	return s
}

func (s *GetUserOkrRequest) SetPageSize(v int64) *GetUserOkrRequest {
	s.PageSize = &v
	return s
}

func (s *GetUserOkrRequest) SetPeriodId(v string) *GetUserOkrRequest {
	s.PeriodId = &v
	return s
}

func (s *GetUserOkrRequest) SetUserId(v string) *GetUserOkrRequest {
	s.UserId = &v
	return s
}

type GetUserOkrResponseBody struct {
	Data *GetUserOkrResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetUserOkrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserOkrResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserOkrResponseBody) SetData(v *GetUserOkrResponseBodyData) *GetUserOkrResponseBody {
	s.Data = v
	return s
}

func (s *GetUserOkrResponseBody) SetSuccess(v bool) *GetUserOkrResponseBody {
	s.Success = &v
	return s
}

type GetUserOkrResponseBodyData struct {
	List []*GetUserOkrResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetUserOkrResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUserOkrResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserOkrResponseBodyData) SetList(v []*GetUserOkrResponseBodyDataList) *GetUserOkrResponseBodyData {
	s.List = v
	return s
}

func (s *GetUserOkrResponseBodyData) SetPageNumber(v int64) *GetUserOkrResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetUserOkrResponseBodyData) SetPageSize(v int64) *GetUserOkrResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetUserOkrResponseBodyData) SetTotalCount(v int64) *GetUserOkrResponseBodyData {
	s.TotalCount = &v
	return s
}

type GetUserOkrResponseBodyDataList struct {
	AlignFromIds []io.Reader `json:"alignFromIds,omitempty" xml:"alignFromIds,omitempty" type:"Repeated"`
	AlignToIds   []io.Reader `json:"alignToIds,omitempty" xml:"alignToIds,omitempty" type:"Repeated"`
	// example:
	//
	// Objective demo
	Content io.Reader `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 1648625407694
	GmtCreate *float32 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1648625407694
	GmtModified *float32 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 5dAX8
	Id     io.Reader                               `json:"id,omitempty" xml:"id,omitempty"`
	KrList []*GetUserOkrResponseBodyDataListKrList `json:"krList,omitempty" xml:"krList,omitempty" type:"Repeated"`
	Owner  *GetUserOkrResponseBodyDataListOwner    `json:"owner,omitempty" xml:"owner,omitempty" type:"Struct"`
	// example:
	//
	// 1006
	PeriodId   io.Reader  `json:"periodId,omitempty" xml:"periodId,omitempty"`
	Permission []*float32 `json:"permission,omitempty" xml:"permission,omitempty" type:"Repeated"`
	// example:
	//
	// 3021332
	Position *int32                                  `json:"position,omitempty" xml:"position,omitempty"`
	Progress *GetUserOkrResponseBodyDataListProgress `json:"progress,omitempty" xml:"progress,omitempty" type:"Struct"`
	// example:
	//
	// 100
	ProgressPercent *float32 `json:"progressPercent,omitempty" xml:"progressPercent,omitempty"`
	// example:
	//
	// true
	Published *bool `json:"published,omitempty" xml:"published,omitempty"`
	// example:
	//
	// 20
	Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// s34d
	UserId io.Reader `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// 50
	Weight *float32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GetUserOkrResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetUserOkrResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetUserOkrResponseBodyDataList) SetAlignFromIds(v []io.Reader) *GetUserOkrResponseBodyDataList {
	s.AlignFromIds = v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetAlignToIds(v []io.Reader) *GetUserOkrResponseBodyDataList {
	s.AlignToIds = v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetContent(v io.Reader) *GetUserOkrResponseBodyDataList {
	s.Content = v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetGmtCreate(v float32) *GetUserOkrResponseBodyDataList {
	s.GmtCreate = &v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetGmtModified(v float32) *GetUserOkrResponseBodyDataList {
	s.GmtModified = &v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetId(v io.Reader) *GetUserOkrResponseBodyDataList {
	s.Id = v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetKrList(v []*GetUserOkrResponseBodyDataListKrList) *GetUserOkrResponseBodyDataList {
	s.KrList = v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetOwner(v *GetUserOkrResponseBodyDataListOwner) *GetUserOkrResponseBodyDataList {
	s.Owner = v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetPeriodId(v io.Reader) *GetUserOkrResponseBodyDataList {
	s.PeriodId = v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetPermission(v []*float32) *GetUserOkrResponseBodyDataList {
	s.Permission = v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetPosition(v int32) *GetUserOkrResponseBodyDataList {
	s.Position = &v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetProgress(v *GetUserOkrResponseBodyDataListProgress) *GetUserOkrResponseBodyDataList {
	s.Progress = v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetProgressPercent(v float32) *GetUserOkrResponseBodyDataList {
	s.ProgressPercent = &v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetPublished(v bool) *GetUserOkrResponseBodyDataList {
	s.Published = &v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetScore(v float32) *GetUserOkrResponseBodyDataList {
	s.Score = &v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetStatus(v int32) *GetUserOkrResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetUserId(v io.Reader) *GetUserOkrResponseBodyDataList {
	s.UserId = v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetWeight(v float32) *GetUserOkrResponseBodyDataList {
	s.Weight = &v
	return s
}

type GetUserOkrResponseBodyDataListKrList struct {
	// example:
	//
	// 你好
	Content io.Reader `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 1648625407694
	GmtCreate *float32 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1648625407694
	GmtModified *float32 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5w9f
	Id io.Reader `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 5wf8
	ObjectiveId io.Reader  `json:"objectiveId,omitempty" xml:"objectiveId,omitempty"`
	Permission  []*float32 `json:"permission,omitempty" xml:"permission,omitempty" type:"Repeated"`
	// example:
	//
	// 35614536
	Position *int64                                        `json:"position,omitempty" xml:"position,omitempty"`
	Progress *GetUserOkrResponseBodyDataListKrListProgress `json:"progress,omitempty" xml:"progress,omitempty" type:"Struct"`
	// example:
	//
	// 44
	Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
	// example:
	//
	// 44
	Weight *float32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GetUserOkrResponseBodyDataListKrList) String() string {
	return tea.Prettify(s)
}

func (s GetUserOkrResponseBodyDataListKrList) GoString() string {
	return s.String()
}

func (s *GetUserOkrResponseBodyDataListKrList) SetContent(v io.Reader) *GetUserOkrResponseBodyDataListKrList {
	s.Content = v
	return s
}

func (s *GetUserOkrResponseBodyDataListKrList) SetGmtCreate(v float32) *GetUserOkrResponseBodyDataListKrList {
	s.GmtCreate = &v
	return s
}

func (s *GetUserOkrResponseBodyDataListKrList) SetGmtModified(v float32) *GetUserOkrResponseBodyDataListKrList {
	s.GmtModified = &v
	return s
}

func (s *GetUserOkrResponseBodyDataListKrList) SetId(v io.Reader) *GetUserOkrResponseBodyDataListKrList {
	s.Id = v
	return s
}

func (s *GetUserOkrResponseBodyDataListKrList) SetObjectiveId(v io.Reader) *GetUserOkrResponseBodyDataListKrList {
	s.ObjectiveId = v
	return s
}

func (s *GetUserOkrResponseBodyDataListKrList) SetPermission(v []*float32) *GetUserOkrResponseBodyDataListKrList {
	s.Permission = v
	return s
}

func (s *GetUserOkrResponseBodyDataListKrList) SetPosition(v int64) *GetUserOkrResponseBodyDataListKrList {
	s.Position = &v
	return s
}

func (s *GetUserOkrResponseBodyDataListKrList) SetProgress(v *GetUserOkrResponseBodyDataListKrListProgress) *GetUserOkrResponseBodyDataListKrList {
	s.Progress = v
	return s
}

func (s *GetUserOkrResponseBodyDataListKrList) SetScore(v float32) *GetUserOkrResponseBodyDataListKrList {
	s.Score = &v
	return s
}

func (s *GetUserOkrResponseBodyDataListKrList) SetWeight(v float32) *GetUserOkrResponseBodyDataListKrList {
	s.Weight = &v
	return s
}

type GetUserOkrResponseBodyDataListKrListProgress struct {
	// example:
	//
	// 30
	Percent *int32 `json:"percent,omitempty" xml:"percent,omitempty"`
}

func (s GetUserOkrResponseBodyDataListKrListProgress) String() string {
	return tea.Prettify(s)
}

func (s GetUserOkrResponseBodyDataListKrListProgress) GoString() string {
	return s.String()
}

func (s *GetUserOkrResponseBodyDataListKrListProgress) SetPercent(v int32) *GetUserOkrResponseBodyDataListKrListProgress {
	s.Percent = &v
	return s
}

type GetUserOkrResponseBodyDataListOwner struct {
	// example:
	//
	// @lADPDh0cQ_j4Mi_NBULNBUA
	AvatarMediaId io.Reader `json:"avatarMediaId,omitempty" xml:"avatarMediaId,omitempty"`
	// example:
	//
	// ding4d1c8883ff63ee8124f2f5cc6abecb85
	CorpId io.Reader `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// K1AMgq
	Id io.Reader `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 你好
	Nickname io.Reader `json:"nickname,omitempty" xml:"nickname,omitempty"`
	// example:
	//
	// 06186238011033616
	UserId io.Reader `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserOkrResponseBodyDataListOwner) String() string {
	return tea.Prettify(s)
}

func (s GetUserOkrResponseBodyDataListOwner) GoString() string {
	return s.String()
}

func (s *GetUserOkrResponseBodyDataListOwner) SetAvatarMediaId(v io.Reader) *GetUserOkrResponseBodyDataListOwner {
	s.AvatarMediaId = v
	return s
}

func (s *GetUserOkrResponseBodyDataListOwner) SetCorpId(v io.Reader) *GetUserOkrResponseBodyDataListOwner {
	s.CorpId = v
	return s
}

func (s *GetUserOkrResponseBodyDataListOwner) SetId(v io.Reader) *GetUserOkrResponseBodyDataListOwner {
	s.Id = v
	return s
}

func (s *GetUserOkrResponseBodyDataListOwner) SetNickname(v io.Reader) *GetUserOkrResponseBodyDataListOwner {
	s.Nickname = v
	return s
}

func (s *GetUserOkrResponseBodyDataListOwner) SetUserId(v io.Reader) *GetUserOkrResponseBodyDataListOwner {
	s.UserId = v
	return s
}

type GetUserOkrResponseBodyDataListProgress struct {
	// example:
	//
	// 100
	Percent *int32 `json:"percent,omitempty" xml:"percent,omitempty"`
}

func (s GetUserOkrResponseBodyDataListProgress) String() string {
	return tea.Prettify(s)
}

func (s GetUserOkrResponseBodyDataListProgress) GoString() string {
	return s.String()
}

func (s *GetUserOkrResponseBodyDataListProgress) SetPercent(v int32) *GetUserOkrResponseBodyDataListProgress {
	s.Percent = &v
	return s
}

type GetUserOkrResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserOkrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserOkrResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserOkrResponse) GoString() string {
	return s.String()
}

func (s *GetUserOkrResponse) SetHeaders(v map[string]*string) *GetUserOkrResponse {
	s.Headers = v
	return s
}

func (s *GetUserOkrResponse) SetStatusCode(v int32) *GetUserOkrResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserOkrResponse) SetBody(v *GetUserOkrResponseBody) *GetUserOkrResponse {
	s.Body = v
	return s
}

type OkrObjectivesBatchHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s OkrObjectivesBatchHeaders) String() string {
	return tea.Prettify(s)
}

func (s OkrObjectivesBatchHeaders) GoString() string {
	return s.String()
}

func (s *OkrObjectivesBatchHeaders) SetCommonHeaders(v map[string]*string) *OkrObjectivesBatchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OkrObjectivesBatchHeaders) SetXAcsDingtalkAccessToken(v string) *OkrObjectivesBatchHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type OkrObjectivesBatchRequest struct {
	// example:
	//
	// dingOKR
	GoodsCode    *string   `json:"goodsCode,omitempty" xml:"goodsCode,omitempty"`
	ObjectiveIds []*string `json:"objectiveIds,omitempty" xml:"objectiveIds,omitempty" type:"Repeated"`
}

func (s OkrObjectivesBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s OkrObjectivesBatchRequest) GoString() string {
	return s.String()
}

func (s *OkrObjectivesBatchRequest) SetGoodsCode(v string) *OkrObjectivesBatchRequest {
	s.GoodsCode = &v
	return s
}

func (s *OkrObjectivesBatchRequest) SetObjectiveIds(v []*string) *OkrObjectivesBatchRequest {
	s.ObjectiveIds = v
	return s
}

type OkrObjectivesBatchResponseBody struct {
	Content []*OpenObjectiveDTO `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	Success *bool               `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OkrObjectivesBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OkrObjectivesBatchResponseBody) GoString() string {
	return s.String()
}

func (s *OkrObjectivesBatchResponseBody) SetContent(v []*OpenObjectiveDTO) *OkrObjectivesBatchResponseBody {
	s.Content = v
	return s
}

func (s *OkrObjectivesBatchResponseBody) SetSuccess(v bool) *OkrObjectivesBatchResponseBody {
	s.Success = &v
	return s
}

type OkrObjectivesBatchResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OkrObjectivesBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OkrObjectivesBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s OkrObjectivesBatchResponse) GoString() string {
	return s.String()
}

func (s *OkrObjectivesBatchResponse) SetHeaders(v map[string]*string) *OkrObjectivesBatchResponse {
	s.Headers = v
	return s
}

func (s *OkrObjectivesBatchResponse) SetStatusCode(v int32) *OkrObjectivesBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *OkrObjectivesBatchResponse) SetBody(v *OkrObjectivesBatchResponseBody) *OkrObjectivesBatchResponse {
	s.Body = v
	return s
}

type OkrObjectivesByUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s OkrObjectivesByUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s OkrObjectivesByUserHeaders) GoString() string {
	return s.String()
}

func (s *OkrObjectivesByUserHeaders) SetCommonHeaders(v map[string]*string) *OkrObjectivesByUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OkrObjectivesByUserHeaders) SetXAcsDingtalkAccessToken(v string) *OkrObjectivesByUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type OkrObjectivesByUserRequest struct {
	// example:
	//
	// dingOKR
	GoodsCode  *string `json:"goodsCode,omitempty" xml:"goodsCode,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s OkrObjectivesByUserRequest) String() string {
	return tea.Prettify(s)
}

func (s OkrObjectivesByUserRequest) GoString() string {
	return s.String()
}

func (s *OkrObjectivesByUserRequest) SetGoodsCode(v string) *OkrObjectivesByUserRequest {
	s.GoodsCode = &v
	return s
}

func (s *OkrObjectivesByUserRequest) SetPageNumber(v int32) *OkrObjectivesByUserRequest {
	s.PageNumber = &v
	return s
}

func (s *OkrObjectivesByUserRequest) SetPageSize(v int32) *OkrObjectivesByUserRequest {
	s.PageSize = &v
	return s
}

type OkrObjectivesByUserResponseBody struct {
	Content   *OkrObjectivesByUserResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	RequestId *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OkrObjectivesByUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OkrObjectivesByUserResponseBody) GoString() string {
	return s.String()
}

func (s *OkrObjectivesByUserResponseBody) SetContent(v *OkrObjectivesByUserResponseBodyContent) *OkrObjectivesByUserResponseBody {
	s.Content = v
	return s
}

func (s *OkrObjectivesByUserResponseBody) SetRequestId(v string) *OkrObjectivesByUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *OkrObjectivesByUserResponseBody) SetSuccess(v bool) *OkrObjectivesByUserResponseBody {
	s.Success = &v
	return s
}

type OkrObjectivesByUserResponseBodyContent struct {
	Result []*OpenObjectiveDTO `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// 15
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s OkrObjectivesByUserResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s OkrObjectivesByUserResponseBodyContent) GoString() string {
	return s.String()
}

func (s *OkrObjectivesByUserResponseBodyContent) SetResult(v []*OpenObjectiveDTO) *OkrObjectivesByUserResponseBodyContent {
	s.Result = v
	return s
}

func (s *OkrObjectivesByUserResponseBodyContent) SetTotalCount(v int64) *OkrObjectivesByUserResponseBodyContent {
	s.TotalCount = &v
	return s
}

type OkrObjectivesByUserResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OkrObjectivesByUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OkrObjectivesByUserResponse) String() string {
	return tea.Prettify(s)
}

func (s OkrObjectivesByUserResponse) GoString() string {
	return s.String()
}

func (s *OkrObjectivesByUserResponse) SetHeaders(v map[string]*string) *OkrObjectivesByUserResponse {
	s.Headers = v
	return s
}

func (s *OkrObjectivesByUserResponse) SetStatusCode(v int32) *OkrObjectivesByUserResponse {
	s.StatusCode = &v
	return s
}

func (s *OkrObjectivesByUserResponse) SetBody(v *OkrObjectivesByUserResponseBody) *OkrObjectivesByUserResponse {
	s.Body = v
	return s
}

type OkrPeriodsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s OkrPeriodsHeaders) String() string {
	return tea.Prettify(s)
}

func (s OkrPeriodsHeaders) GoString() string {
	return s.String()
}

func (s *OkrPeriodsHeaders) SetCommonHeaders(v map[string]*string) *OkrPeriodsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OkrPeriodsHeaders) SetXAcsDingtalkAccessToken(v string) *OkrPeriodsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type OkrPeriodsRequest struct {
	// example:
	//
	// dingOKR
	GoodsCode *string `json:"goodsCode,omitempty" xml:"goodsCode,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 0
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s OkrPeriodsRequest) String() string {
	return tea.Prettify(s)
}

func (s OkrPeriodsRequest) GoString() string {
	return s.String()
}

func (s *OkrPeriodsRequest) SetGoodsCode(v string) *OkrPeriodsRequest {
	s.GoodsCode = &v
	return s
}

func (s *OkrPeriodsRequest) SetPageNumber(v int64) *OkrPeriodsRequest {
	s.PageNumber = &v
	return s
}

func (s *OkrPeriodsRequest) SetPageSize(v int64) *OkrPeriodsRequest {
	s.PageSize = &v
	return s
}

func (s *OkrPeriodsRequest) SetStatus(v int64) *OkrPeriodsRequest {
	s.Status = &v
	return s
}

type OkrPeriodsResponseBody struct {
	Content   *OkrPeriodsResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	RequestId *string                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OkrPeriodsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OkrPeriodsResponseBody) GoString() string {
	return s.String()
}

func (s *OkrPeriodsResponseBody) SetContent(v *OkrPeriodsResponseBodyContent) *OkrPeriodsResponseBody {
	s.Content = v
	return s
}

func (s *OkrPeriodsResponseBody) SetRequestId(v string) *OkrPeriodsResponseBody {
	s.RequestId = &v
	return s
}

func (s *OkrPeriodsResponseBody) SetSuccess(v bool) *OkrPeriodsResponseBody {
	s.Success = &v
	return s
}

type OkrPeriodsResponseBodyContent struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32           `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Result   []*OpenPeriodDTO `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// 15
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s OkrPeriodsResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s OkrPeriodsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *OkrPeriodsResponseBodyContent) SetPageNumber(v int32) *OkrPeriodsResponseBodyContent {
	s.PageNumber = &v
	return s
}

func (s *OkrPeriodsResponseBodyContent) SetPageSize(v int32) *OkrPeriodsResponseBodyContent {
	s.PageSize = &v
	return s
}

func (s *OkrPeriodsResponseBodyContent) SetResult(v []*OpenPeriodDTO) *OkrPeriodsResponseBodyContent {
	s.Result = v
	return s
}

func (s *OkrPeriodsResponseBodyContent) SetTotalCount(v int64) *OkrPeriodsResponseBodyContent {
	s.TotalCount = &v
	return s
}

type OkrPeriodsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OkrPeriodsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OkrPeriodsResponse) String() string {
	return tea.Prettify(s)
}

func (s OkrPeriodsResponse) GoString() string {
	return s.String()
}

func (s *OkrPeriodsResponse) SetHeaders(v map[string]*string) *OkrPeriodsResponse {
	s.Headers = v
	return s
}

func (s *OkrPeriodsResponse) SetStatusCode(v int32) *OkrPeriodsResponse {
	s.StatusCode = &v
	return s
}

func (s *OkrPeriodsResponse) SetBody(v *OkrPeriodsResponseBody) *OkrPeriodsResponse {
	s.Body = v
	return s
}

type UnAlignObjectiveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UnAlignObjectiveHeaders) String() string {
	return tea.Prettify(s)
}

func (s UnAlignObjectiveHeaders) GoString() string {
	return s.String()
}

func (s *UnAlignObjectiveHeaders) SetCommonHeaders(v map[string]*string) *UnAlignObjectiveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UnAlignObjectiveHeaders) SetXAcsDingtalkAccessToken(v string) *UnAlignObjectiveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UnAlignObjectiveRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1006
	PeriodId *string `json:"periodId,omitempty" xml:"periodId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 59YD
	TargetId *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0115396701752283
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UnAlignObjectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s UnAlignObjectiveRequest) GoString() string {
	return s.String()
}

func (s *UnAlignObjectiveRequest) SetPeriodId(v string) *UnAlignObjectiveRequest {
	s.PeriodId = &v
	return s
}

func (s *UnAlignObjectiveRequest) SetTargetId(v string) *UnAlignObjectiveRequest {
	s.TargetId = &v
	return s
}

func (s *UnAlignObjectiveRequest) SetUserId(v string) *UnAlignObjectiveRequest {
	s.UserId = &v
	return s
}

type UnAlignObjectiveResponseBody struct {
	Data *UnAlignObjectiveResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UnAlignObjectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnAlignObjectiveResponseBody) GoString() string {
	return s.String()
}

func (s *UnAlignObjectiveResponseBody) SetData(v *UnAlignObjectiveResponseBodyData) *UnAlignObjectiveResponseBody {
	s.Data = v
	return s
}

func (s *UnAlignObjectiveResponseBody) SetSuccess(v bool) *UnAlignObjectiveResponseBody {
	s.Success = &v
	return s
}

type UnAlignObjectiveResponseBodyData struct {
	// example:
	//
	// 59YD
	AlignId io.Reader `json:"alignId,omitempty" xml:"alignId,omitempty"`
	// example:
	//
	// 5dAX8
	Id io.Reader `json:"id,omitempty" xml:"id,omitempty"`
}

func (s UnAlignObjectiveResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UnAlignObjectiveResponseBodyData) GoString() string {
	return s.String()
}

func (s *UnAlignObjectiveResponseBodyData) SetAlignId(v io.Reader) *UnAlignObjectiveResponseBodyData {
	s.AlignId = v
	return s
}

func (s *UnAlignObjectiveResponseBodyData) SetId(v io.Reader) *UnAlignObjectiveResponseBodyData {
	s.Id = v
	return s
}

type UnAlignObjectiveResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnAlignObjectiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnAlignObjectiveResponse) String() string {
	return tea.Prettify(s)
}

func (s UnAlignObjectiveResponse) GoString() string {
	return s.String()
}

func (s *UnAlignObjectiveResponse) SetHeaders(v map[string]*string) *UnAlignObjectiveResponse {
	s.Headers = v
	return s
}

func (s *UnAlignObjectiveResponse) SetStatusCode(v int32) *UnAlignObjectiveResponse {
	s.StatusCode = &v
	return s
}

func (s *UnAlignObjectiveResponse) SetBody(v *UnAlignObjectiveResponseBody) *UnAlignObjectiveResponse {
	s.Body = v
	return s
}

type UpdateKROfContentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateKROfContentHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfContentHeaders) GoString() string {
	return s.String()
}

func (s *UpdateKROfContentHeaders) SetCommonHeaders(v map[string]*string) *UpdateKROfContentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateKROfContentHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateKROfContentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateKROfContentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 更新内容
	Content         *string   `json:"content,omitempty" xml:"content,omitempty"`
	UpdateQuoteList []*string `json:"updateQuoteList,omitempty" xml:"updateQuoteList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 46GM2
	KrId *string `json:"krId,omitempty" xml:"krId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0115396701752283
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateKROfContentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfContentRequest) GoString() string {
	return s.String()
}

func (s *UpdateKROfContentRequest) SetContent(v string) *UpdateKROfContentRequest {
	s.Content = &v
	return s
}

func (s *UpdateKROfContentRequest) SetUpdateQuoteList(v []*string) *UpdateKROfContentRequest {
	s.UpdateQuoteList = v
	return s
}

func (s *UpdateKROfContentRequest) SetKrId(v string) *UpdateKROfContentRequest {
	s.KrId = &v
	return s
}

func (s *UpdateKROfContentRequest) SetUserId(v string) *UpdateKROfContentRequest {
	s.UserId = &v
	return s
}

type UpdateKROfContentResponseBody struct {
	// This parameter is required.
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateKROfContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfContentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKROfContentResponseBody) SetData(v bool) *UpdateKROfContentResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateKROfContentResponseBody) SetSuccess(v bool) *UpdateKROfContentResponseBody {
	s.Success = &v
	return s
}

type UpdateKROfContentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKROfContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKROfContentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfContentResponse) GoString() string {
	return s.String()
}

func (s *UpdateKROfContentResponse) SetHeaders(v map[string]*string) *UpdateKROfContentResponse {
	s.Headers = v
	return s
}

func (s *UpdateKROfContentResponse) SetStatusCode(v int32) *UpdateKROfContentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKROfContentResponse) SetBody(v *UpdateKROfContentResponseBody) *UpdateKROfContentResponse {
	s.Body = v
	return s
}

type UpdateKROfScoreHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateKROfScoreHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfScoreHeaders) GoString() string {
	return s.String()
}

func (s *UpdateKROfScoreHeaders) SetCommonHeaders(v map[string]*string) *UpdateKROfScoreHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateKROfScoreHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateKROfScoreHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateKROfScoreRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	Score *int64 `json:"score,omitempty" xml:"score,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 46GM2
	KrId *string `json:"krId,omitempty" xml:"krId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0115396701752283
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateKROfScoreRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfScoreRequest) GoString() string {
	return s.String()
}

func (s *UpdateKROfScoreRequest) SetScore(v int64) *UpdateKROfScoreRequest {
	s.Score = &v
	return s
}

func (s *UpdateKROfScoreRequest) SetKrId(v string) *UpdateKROfScoreRequest {
	s.KrId = &v
	return s
}

func (s *UpdateKROfScoreRequest) SetUserId(v string) *UpdateKROfScoreRequest {
	s.UserId = &v
	return s
}

type UpdateKROfScoreResponseBody struct {
	// This parameter is required.
	Data *UpdateKROfScoreResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateKROfScoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfScoreResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKROfScoreResponseBody) SetData(v *UpdateKROfScoreResponseBodyData) *UpdateKROfScoreResponseBody {
	s.Data = v
	return s
}

func (s *UpdateKROfScoreResponseBody) SetSuccess(v bool) *UpdateKROfScoreResponseBody {
	s.Success = &v
	return s
}

type UpdateKROfScoreResponseBodyData struct {
	// This parameter is required.
	//
	// example:
	//
	// 50
	ObjectiveScore *int64 `json:"objectiveScore,omitempty" xml:"objectiveScore,omitempty"`
}

func (s UpdateKROfScoreResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfScoreResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateKROfScoreResponseBodyData) SetObjectiveScore(v int64) *UpdateKROfScoreResponseBodyData {
	s.ObjectiveScore = &v
	return s
}

type UpdateKROfScoreResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKROfScoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKROfScoreResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfScoreResponse) GoString() string {
	return s.String()
}

func (s *UpdateKROfScoreResponse) SetHeaders(v map[string]*string) *UpdateKROfScoreResponse {
	s.Headers = v
	return s
}

func (s *UpdateKROfScoreResponse) SetStatusCode(v int32) *UpdateKROfScoreResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKROfScoreResponse) SetBody(v *UpdateKROfScoreResponseBody) *UpdateKROfScoreResponse {
	s.Body = v
	return s
}

type UpdateKROfWeightHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateKROfWeightHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfWeightHeaders) GoString() string {
	return s.String()
}

func (s *UpdateKROfWeightHeaders) SetCommonHeaders(v map[string]*string) *UpdateKROfWeightHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateKROfWeightHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateKROfWeightHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateKROfWeightRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 100
	Weight *int64 `json:"weight,omitempty" xml:"weight,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 46GM2
	KrId *string `json:"krId,omitempty" xml:"krId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0115396701752283
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateKROfWeightRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfWeightRequest) GoString() string {
	return s.String()
}

func (s *UpdateKROfWeightRequest) SetWeight(v int64) *UpdateKROfWeightRequest {
	s.Weight = &v
	return s
}

func (s *UpdateKROfWeightRequest) SetKrId(v string) *UpdateKROfWeightRequest {
	s.KrId = &v
	return s
}

func (s *UpdateKROfWeightRequest) SetUserId(v string) *UpdateKROfWeightRequest {
	s.UserId = &v
	return s
}

type UpdateKROfWeightResponseBody struct {
	// This parameter is required.
	Data *UpdateKROfWeightResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateKROfWeightResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfWeightResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKROfWeightResponseBody) SetData(v *UpdateKROfWeightResponseBodyData) *UpdateKROfWeightResponseBody {
	s.Data = v
	return s
}

func (s *UpdateKROfWeightResponseBody) SetSuccess(v bool) *UpdateKROfWeightResponseBody {
	s.Success = &v
	return s
}

type UpdateKROfWeightResponseBodyData struct {
	// This parameter is required.
	ObjectiveProgress *UpdateKROfWeightResponseBodyDataObjectiveProgress `json:"objectiveProgress,omitempty" xml:"objectiveProgress,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	ObjectiveScore *int64 `json:"objectiveScore,omitempty" xml:"objectiveScore,omitempty"`
}

func (s UpdateKROfWeightResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfWeightResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateKROfWeightResponseBodyData) SetObjectiveProgress(v *UpdateKROfWeightResponseBodyDataObjectiveProgress) *UpdateKROfWeightResponseBodyData {
	s.ObjectiveProgress = v
	return s
}

func (s *UpdateKROfWeightResponseBodyData) SetObjectiveScore(v int64) *UpdateKROfWeightResponseBodyData {
	s.ObjectiveScore = &v
	return s
}

type UpdateKROfWeightResponseBodyDataObjectiveProgress struct {
	// This parameter is required.
	//
	// example:
	//
	// 0
	Percent *int64 `json:"percent,omitempty" xml:"percent,omitempty"`
}

func (s UpdateKROfWeightResponseBodyDataObjectiveProgress) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfWeightResponseBodyDataObjectiveProgress) GoString() string {
	return s.String()
}

func (s *UpdateKROfWeightResponseBodyDataObjectiveProgress) SetPercent(v int64) *UpdateKROfWeightResponseBodyDataObjectiveProgress {
	s.Percent = &v
	return s
}

type UpdateKROfWeightResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKROfWeightResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKROfWeightResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfWeightResponse) GoString() string {
	return s.String()
}

func (s *UpdateKROfWeightResponse) SetHeaders(v map[string]*string) *UpdateKROfWeightResponse {
	s.Headers = v
	return s
}

func (s *UpdateKROfWeightResponse) SetStatusCode(v int32) *UpdateKROfWeightResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKROfWeightResponse) SetBody(v *UpdateKROfWeightResponseBody) *UpdateKROfWeightResponse {
	s.Body = v
	return s
}

type UpdateObjectiveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateObjectiveHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateObjectiveHeaders) GoString() string {
	return s.String()
}

func (s *UpdateObjectiveHeaders) SetCommonHeaders(v map[string]*string) *UpdateObjectiveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateObjectiveHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateObjectiveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateObjectiveRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 更新的内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 06186238011033616
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateObjectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateObjectiveRequest) GoString() string {
	return s.String()
}

func (s *UpdateObjectiveRequest) SetContent(v string) *UpdateObjectiveRequest {
	s.Content = &v
	return s
}

func (s *UpdateObjectiveRequest) SetUserId(v string) *UpdateObjectiveRequest {
	s.UserId = &v
	return s
}

type UpdateObjectiveResponseBody struct {
	Data *UpdateObjectiveResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateObjectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateObjectiveResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateObjectiveResponseBody) SetData(v *UpdateObjectiveResponseBodyData) *UpdateObjectiveResponseBody {
	s.Data = v
	return s
}

func (s *UpdateObjectiveResponseBody) SetSuccess(v bool) *UpdateObjectiveResponseBody {
	s.Success = &v
	return s
}

type UpdateObjectiveResponseBodyData struct {
	// This parameter is required.
	//
	// example:
	//
	// 58YD
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 33453
	Position *float32 `json:"position,omitempty" xml:"position,omitempty"`
}

func (s UpdateObjectiveResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateObjectiveResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateObjectiveResponseBodyData) SetId(v string) *UpdateObjectiveResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateObjectiveResponseBodyData) SetPosition(v float32) *UpdateObjectiveResponseBodyData {
	s.Position = &v
	return s
}

type UpdateObjectiveResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateObjectiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateObjectiveResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateObjectiveResponse) GoString() string {
	return s.String()
}

func (s *UpdateObjectiveResponse) SetHeaders(v map[string]*string) *UpdateObjectiveResponse {
	s.Headers = v
	return s
}

func (s *UpdateObjectiveResponse) SetStatusCode(v int32) *UpdateObjectiveResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateObjectiveResponse) SetBody(v *UpdateObjectiveResponseBody) *UpdateObjectiveResponse {
	s.Body = v
	return s
}

type UpdatePrivacyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdatePrivacyHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivacyHeaders) GoString() string {
	return s.String()
}

func (s *UpdatePrivacyHeaders) SetCommonHeaders(v map[string]*string) *UpdatePrivacyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdatePrivacyHeaders) SetXAcsDingtalkAccessToken(v string) *UpdatePrivacyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdatePrivacyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// public
	Privacy *string `json:"privacy,omitempty" xml:"privacy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3RF5
	TargetId *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0115396701752283
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdatePrivacyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivacyRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrivacyRequest) SetPrivacy(v string) *UpdatePrivacyRequest {
	s.Privacy = &v
	return s
}

func (s *UpdatePrivacyRequest) SetTargetId(v string) *UpdatePrivacyRequest {
	s.TargetId = &v
	return s
}

func (s *UpdatePrivacyRequest) SetTargetType(v string) *UpdatePrivacyRequest {
	s.TargetType = &v
	return s
}

func (s *UpdatePrivacyRequest) SetUserId(v string) *UpdatePrivacyRequest {
	s.UserId = &v
	return s
}

type UpdatePrivacyResponseBody struct {
	// This parameter is required.
	Data *UpdatePrivacyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdatePrivacyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivacyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrivacyResponseBody) SetData(v *UpdatePrivacyResponseBodyData) *UpdatePrivacyResponseBody {
	s.Data = v
	return s
}

func (s *UpdatePrivacyResponseBody) SetSuccess(v bool) *UpdatePrivacyResponseBody {
	s.Success = &v
	return s
}

type UpdatePrivacyResponseBodyData struct {
	// This parameter is required.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	PolicyList []*UpdatePrivacyResponseBodyDataPolicyList `json:"policyList,omitempty" xml:"policyList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// public
	Privacy *string `json:"privacy,omitempty" xml:"privacy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// period
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdatePrivacyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivacyResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdatePrivacyResponseBodyData) SetId(v string) *UpdatePrivacyResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdatePrivacyResponseBodyData) SetPolicyList(v []*UpdatePrivacyResponseBodyDataPolicyList) *UpdatePrivacyResponseBodyData {
	s.PolicyList = v
	return s
}

func (s *UpdatePrivacyResponseBodyData) SetPrivacy(v string) *UpdatePrivacyResponseBodyData {
	s.Privacy = &v
	return s
}

func (s *UpdatePrivacyResponseBodyData) SetType(v string) *UpdatePrivacyResponseBodyData {
	s.Type = &v
	return s
}

type UpdatePrivacyResponseBodyDataPolicyList struct {
	// This parameter is required.
	MemberList []*UpdatePrivacyResponseBodyDataPolicyListMemberList `json:"memberList,omitempty" xml:"memberList,omitempty" type:"Repeated"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdatePrivacyResponseBodyDataPolicyList) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivacyResponseBodyDataPolicyList) GoString() string {
	return s.String()
}

func (s *UpdatePrivacyResponseBodyDataPolicyList) SetMemberList(v []*UpdatePrivacyResponseBodyDataPolicyListMemberList) *UpdatePrivacyResponseBodyDataPolicyList {
	s.MemberList = v
	return s
}

func (s *UpdatePrivacyResponseBodyDataPolicyList) SetName(v string) *UpdatePrivacyResponseBodyDataPolicyList {
	s.Name = &v
	return s
}

func (s *UpdatePrivacyResponseBodyDataPolicyList) SetType(v int64) *UpdatePrivacyResponseBodyDataPolicyList {
	s.Type = &v
	return s
}

type UpdatePrivacyResponseBodyDataPolicyListMemberList struct {
	// This parameter is required.
	Id       *string `json:"id,omitempty" xml:"id,omitempty"`
	Nickname *string `json:"nickname,omitempty" xml:"nickname,omitempty"`
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdatePrivacyResponseBodyDataPolicyListMemberList) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivacyResponseBodyDataPolicyListMemberList) GoString() string {
	return s.String()
}

func (s *UpdatePrivacyResponseBodyDataPolicyListMemberList) SetId(v string) *UpdatePrivacyResponseBodyDataPolicyListMemberList {
	s.Id = &v
	return s
}

func (s *UpdatePrivacyResponseBodyDataPolicyListMemberList) SetNickname(v string) *UpdatePrivacyResponseBodyDataPolicyListMemberList {
	s.Nickname = &v
	return s
}

func (s *UpdatePrivacyResponseBodyDataPolicyListMemberList) SetType(v string) *UpdatePrivacyResponseBodyDataPolicyListMemberList {
	s.Type = &v
	return s
}

type UpdatePrivacyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePrivacyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePrivacyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivacyResponse) GoString() string {
	return s.String()
}

func (s *UpdatePrivacyResponse) SetHeaders(v map[string]*string) *UpdatePrivacyResponse {
	s.Headers = v
	return s
}

func (s *UpdatePrivacyResponse) SetStatusCode(v int32) *UpdatePrivacyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePrivacyResponse) SetBody(v *UpdatePrivacyResponseBody) *UpdatePrivacyResponse {
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
// 增加对齐目标
//
// @param request - AlignObjectiveRequest
//
// @param headers - AlignObjectiveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AlignObjectiveResponse
func (client *Client) AlignObjectiveWithOptions(objectiveId *string, request *AlignObjectiveRequest, headers *AlignObjectiveHeaders, runtime *util.RuntimeOptions) (_result *AlignObjectiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PeriodId)) {
		body["periodId"] = request.PeriodId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		body["targetId"] = request.TargetId
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
		Action:      tea.String("AlignObjective"),
		Version:     tea.String("okr_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/okr/objectives/" + tea.StringValue(objectiveId) + "/alignments"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AlignObjectiveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增加对齐目标
//
// @param request - AlignObjectiveRequest
//
// @return AlignObjectiveResponse
func (client *Client) AlignObjective(objectiveId *string, request *AlignObjectiveRequest) (_result *AlignObjectiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AlignObjectiveHeaders{}
	_result = &AlignObjectiveResponse{}
	_body, _err := client.AlignObjectiveWithOptions(objectiveId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量添加权限信息
//
// @param request - BatchAddPermissionRequest
//
// @param headers - BatchAddPermissionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchAddPermissionResponse
func (client *Client) BatchAddPermissionWithOptions(request *BatchAddPermissionRequest, headers *BatchAddPermissionHeaders, runtime *util.RuntimeOptions) (_result *BatchAddPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.List)) {
		body["list"] = request.List
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		body["targetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		body["targetType"] = request.TargetType
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
		Action:      tea.String("BatchAddPermission"),
		Version:     tea.String("okr_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/okr/permissions/batch"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchAddPermissionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量添加权限信息
//
// @param request - BatchAddPermissionRequest
//
// @return BatchAddPermissionResponse
func (client *Client) BatchAddPermission(request *BatchAddPermissionRequest) (_result *BatchAddPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchAddPermissionHeaders{}
	_result = &BatchAddPermissionResponse{}
	_body, _err := client.BatchAddPermissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询目标
//
// @param request - BatchQueryObjectiveRequest
//
// @param headers - BatchQueryObjectiveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchQueryObjectiveResponse
func (client *Client) BatchQueryObjectiveWithOptions(request *BatchQueryObjectiveRequest, headers *BatchQueryObjectiveHeaders, runtime *util.RuntimeOptions) (_result *BatchQueryObjectiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ObjectiveIds)) {
		body["objectiveIds"] = request.ObjectiveIds
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodId)) {
		body["periodId"] = request.PeriodId
	}

	if !tea.BoolValue(util.IsUnset(request.WithAlign)) {
		body["withAlign"] = request.WithAlign
	}

	if !tea.BoolValue(util.IsUnset(request.WithKr)) {
		body["withKr"] = request.WithKr
	}

	if !tea.BoolValue(util.IsUnset(request.WithProgress)) {
		body["withProgress"] = request.WithProgress
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
		Action:      tea.String("BatchQueryObjective"),
		Version:     tea.String("okr_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/okr/objectives/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchQueryObjectiveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询目标
//
// @param request - BatchQueryObjectiveRequest
//
// @return BatchQueryObjectiveResponse
func (client *Client) BatchQueryObjective(request *BatchQueryObjectiveRequest) (_result *BatchQueryObjectiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchQueryObjectiveHeaders{}
	_result = &BatchQueryObjectiveResponse{}
	_body, _err := client.BatchQueryObjectiveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询用户信息
//
// @param request - BatchQueryUserRequest
//
// @param headers - BatchQueryUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchQueryUserResponse
func (client *Client) BatchQueryUserWithOptions(request *BatchQueryUserRequest, headers *BatchQueryUserHeaders, runtime *util.RuntimeOptions) (_result *BatchQueryUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OkrUserIds)) {
		body["okrUserIds"] = request.OkrUserIds
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
		Action:      tea.String("BatchQueryUser"),
		Version:     tea.String("okr_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/okr/users/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchQueryUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询用户信息
//
// @param request - BatchQueryUserRequest
//
// @return BatchQueryUserResponse
func (client *Client) BatchQueryUser(request *BatchQueryUserRequest) (_result *BatchQueryUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchQueryUserHeaders{}
	_result = &BatchQueryUserResponse{}
	_body, _err := client.BatchQueryUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建keyResult
//
// @param request - CreateKeyResultRequest
//
// @param headers - CreateKeyResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKeyResultResponse
func (client *Client) CreateKeyResultWithOptions(request *CreateKeyResultRequest, headers *CreateKeyResultHeaders, runtime *util.RuntimeOptions) (_result *CreateKeyResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectiveId)) {
		body["objectiveId"] = request.ObjectiveId
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodId)) {
		body["periodId"] = request.PeriodId
	}

	if !tea.BoolValue(util.IsUnset(request.PrevPosition)) {
		body["prevPosition"] = request.PrevPosition
	}

	if !tea.BoolValue(util.IsUnset(request.Weight)) {
		body["weight"] = request.Weight
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
		Action:      tea.String("CreateKeyResult"),
		Version:     tea.String("okr_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/okr/keyResults"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateKeyResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建keyResult
//
// @param request - CreateKeyResultRequest
//
// @return CreateKeyResultResponse
func (client *Client) CreateKeyResult(request *CreateKeyResultRequest) (_result *CreateKeyResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateKeyResultHeaders{}
	_result = &CreateKeyResultResponse{}
	_body, _err := client.CreateKeyResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建目标
//
// @param request - CreateObjectiveRequest
//
// @param headers - CreateObjectiveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateObjectiveResponse
func (client *Client) CreateObjectiveWithOptions(request *CreateObjectiveRequest, headers *CreateObjectiveHeaders, runtime *util.RuntimeOptions) (_result *CreateObjectiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodId)) {
		body["periodId"] = request.PeriodId
	}

	if !tea.BoolValue(util.IsUnset(request.PrevPosition)) {
		body["prevPosition"] = request.PrevPosition
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
		Action:      tea.String("CreateObjective"),
		Version:     tea.String("okr_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/okr/objectives"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateObjectiveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建目标
//
// @param request - CreateObjectiveRequest
//
// @return CreateObjectiveResponse
func (client *Client) CreateObjective(request *CreateObjectiveRequest) (_result *CreateObjectiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateObjectiveHeaders{}
	_result = &CreateObjectiveResponse{}
	_body, _err := client.CreateObjectiveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除keyresult的方法
//
// @param request - DeleteKeyResultRequest
//
// @param headers - DeleteKeyResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteKeyResultResponse
func (client *Client) DeleteKeyResultWithOptions(request *DeleteKeyResultRequest, headers *DeleteKeyResultHeaders, runtime *util.RuntimeOptions) (_result *DeleteKeyResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KrId)) {
		query["krId"] = request.KrId
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
		Action:      tea.String("DeleteKeyResult"),
		Version:     tea.String("okr_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/okr/keyResults"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteKeyResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除keyresult的方法
//
// @param request - DeleteKeyResultRequest
//
// @return DeleteKeyResultResponse
func (client *Client) DeleteKeyResult(request *DeleteKeyResultRequest) (_result *DeleteKeyResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteKeyResultHeaders{}
	_result = &DeleteKeyResultResponse{}
	_body, _err := client.DeleteKeyResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除目标
//
// @param request - DeleteObjectiveRequest
//
// @param headers - DeleteObjectiveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteObjectiveResponse
func (client *Client) DeleteObjectiveWithOptions(objectiveId *string, request *DeleteObjectiveRequest, headers *DeleteObjectiveHeaders, runtime *util.RuntimeOptions) (_result *DeleteObjectiveResponse, _err error) {
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
		Action:      tea.String("DeleteObjective"),
		Version:     tea.String("okr_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/okr/objectives/" + tea.StringValue(objectiveId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteObjectiveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除目标
//
// @param request - DeleteObjectiveRequest
//
// @return DeleteObjectiveResponse
func (client *Client) DeleteObjective(objectiveId *string, request *DeleteObjectiveRequest) (_result *DeleteObjectiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteObjectiveHeaders{}
	_result = &DeleteObjectiveResponse{}
	_body, _err := client.DeleteObjectiveWithOptions(objectiveId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除权限信息
//
// @param request - DeletePermissionRequest
//
// @param headers - DeletePermissionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePermissionResponse
func (client *Client) DeletePermissionWithOptions(request *DeletePermissionRequest, headers *DeletePermissionHeaders, runtime *util.RuntimeOptions) (_result *DeletePermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyType)) {
		query["policyType"] = request.PolicyType
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["targetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["targetType"] = request.TargetType
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
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
		Action:      tea.String("DeletePermission"),
		Version:     tea.String("okr_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/okr/permissions/delete"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePermissionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除权限信息
//
// @param request - DeletePermissionRequest
//
// @return DeletePermissionResponse
func (client *Client) DeletePermission(request *DeletePermissionRequest) (_result *DeletePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeletePermissionHeaders{}
	_result = &DeletePermissionResponse{}
	_body, _err := client.DeletePermissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取周期列表
//
// @param headers - GetPeriodListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPeriodListResponse
func (client *Client) GetPeriodListWithOptions(headers *GetPeriodListHeaders, runtime *util.RuntimeOptions) (_result *GetPeriodListResponse, _err error) {
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
		Action:      tea.String("GetPeriodList"),
		Version:     tea.String("okr_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/okr/periods"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPeriodListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取周期列表
//
// @return GetPeriodListResponse
func (client *Client) GetPeriodList() (_result *GetPeriodListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPeriodListHeaders{}
	_result = &GetPeriodListResponse{}
	_body, _err := client.GetPeriodListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取权限信息
//
// @param request - GetPermissionRequest
//
// @param headers - GetPermissionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPermissionResponse
func (client *Client) GetPermissionWithOptions(request *GetPermissionRequest, headers *GetPermissionHeaders, runtime *util.RuntimeOptions) (_result *GetPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["targetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["targetType"] = request.TargetType
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.WithKr)) {
		query["withKr"] = request.WithKr
	}

	if !tea.BoolValue(util.IsUnset(request.WithObjective)) {
		query["withObjective"] = request.WithObjective
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
		Action:      tea.String("GetPermission"),
		Version:     tea.String("okr_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/okr/permissions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPermissionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取权限信息
//
// @param request - GetPermissionRequest
//
// @return GetPermissionResponse
func (client *Client) GetPermission(request *GetPermissionRequest) (_result *GetPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPermissionHeaders{}
	_result = &GetPermissionResponse{}
	_body, _err := client.GetPermissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户当前周期下的全部 OKR 内容
//
// @param request - GetUserOkrRequest
//
// @param headers - GetUserOkrHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserOkrResponse
func (client *Client) GetUserOkrWithOptions(request *GetUserOkrRequest, headers *GetUserOkrHeaders, runtime *util.RuntimeOptions) (_result *GetUserOkrResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.PeriodId)) {
		query["periodId"] = request.PeriodId
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
		Action:      tea.String("GetUserOkr"),
		Version:     tea.String("okr_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/okr/users/okrs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserOkrResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户当前周期下的全部 OKR 内容
//
// @param request - GetUserOkrRequest
//
// @return GetUserOkrResponse
func (client *Client) GetUserOkr(request *GetUserOkrRequest) (_result *GetUserOkrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserOkrHeaders{}
	_result = &GetUserOkrResponse{}
	_body, _err := client.GetUserOkrWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询OKR
//
// @param request - OkrObjectivesBatchRequest
//
// @param headers - OkrObjectivesBatchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OkrObjectivesBatchResponse
func (client *Client) OkrObjectivesBatchWithOptions(request *OkrObjectivesBatchRequest, headers *OkrObjectivesBatchHeaders, runtime *util.RuntimeOptions) (_result *OkrObjectivesBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GoodsCode)) {
		body["goodsCode"] = request.GoodsCode
	}

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
		Action:      tea.String("OkrObjectivesBatch"),
		Version:     tea.String("okr_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/okr/pro/objectives/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &OkrObjectivesBatchResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询OKR
//
// @param request - OkrObjectivesBatchRequest
//
// @return OkrObjectivesBatchResponse
func (client *Client) OkrObjectivesBatch(request *OkrObjectivesBatchRequest) (_result *OkrObjectivesBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OkrObjectivesBatchHeaders{}
	_result = &OkrObjectivesBatchResponse{}
	_body, _err := client.OkrObjectivesBatchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询单个用户的OKR
//
// @param request - OkrObjectivesByUserRequest
//
// @param headers - OkrObjectivesByUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OkrObjectivesByUserResponse
func (client *Client) OkrObjectivesByUserWithOptions(dingUserId *string, request *OkrObjectivesByUserRequest, headers *OkrObjectivesByUserHeaders, runtime *util.RuntimeOptions) (_result *OkrObjectivesByUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GoodsCode)) {
		query["goodsCode"] = request.GoodsCode
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
		Action:      tea.String("OkrObjectivesByUser"),
		Version:     tea.String("okr_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/okr/pro/users/" + tea.StringValue(dingUserId) + "/objectives"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &OkrObjectivesByUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询单个用户的OKR
//
// @param request - OkrObjectivesByUserRequest
//
// @return OkrObjectivesByUserResponse
func (client *Client) OkrObjectivesByUser(dingUserId *string, request *OkrObjectivesByUserRequest) (_result *OkrObjectivesByUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OkrObjectivesByUserHeaders{}
	_result = &OkrObjectivesByUserResponse{}
	_body, _err := client.OkrObjectivesByUserWithOptions(dingUserId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取 OKR 周期
//
// @param request - OkrPeriodsRequest
//
// @param headers - OkrPeriodsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OkrPeriodsResponse
func (client *Client) OkrPeriodsWithOptions(request *OkrPeriodsRequest, headers *OkrPeriodsHeaders, runtime *util.RuntimeOptions) (_result *OkrPeriodsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GoodsCode)) {
		query["goodsCode"] = request.GoodsCode
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
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
		Action:      tea.String("OkrPeriods"),
		Version:     tea.String("okr_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/okr/pro/periods"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &OkrPeriodsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取 OKR 周期
//
// @param request - OkrPeriodsRequest
//
// @return OkrPeriodsResponse
func (client *Client) OkrPeriods(request *OkrPeriodsRequest) (_result *OkrPeriodsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OkrPeriodsHeaders{}
	_result = &OkrPeriodsResponse{}
	_body, _err := client.OkrPeriodsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消对齐Objective
//
// @param request - UnAlignObjectiveRequest
//
// @param headers - UnAlignObjectiveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnAlignObjectiveResponse
func (client *Client) UnAlignObjectiveWithOptions(objectiveId *string, request *UnAlignObjectiveRequest, headers *UnAlignObjectiveHeaders, runtime *util.RuntimeOptions) (_result *UnAlignObjectiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PeriodId)) {
		body["periodId"] = request.PeriodId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		body["targetId"] = request.TargetId
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
		Action:      tea.String("UnAlignObjective"),
		Version:     tea.String("okr_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/okr/objectives/" + tea.StringValue(objectiveId) + "/alignments/cancel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UnAlignObjectiveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消对齐Objective
//
// @param request - UnAlignObjectiveRequest
//
// @return UnAlignObjectiveResponse
func (client *Client) UnAlignObjective(objectiveId *string, request *UnAlignObjectiveRequest) (_result *UnAlignObjectiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UnAlignObjectiveHeaders{}
	_result = &UnAlignObjectiveResponse{}
	_body, _err := client.UnAlignObjectiveWithOptions(objectiveId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更改KR内容
//
// @param request - UpdateKROfContentRequest
//
// @param headers - UpdateKROfContentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKROfContentResponse
func (client *Client) UpdateKROfContentWithOptions(request *UpdateKROfContentRequest, headers *UpdateKROfContentHeaders, runtime *util.RuntimeOptions) (_result *UpdateKROfContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KrId)) {
		query["krId"] = request.KrId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateQuoteList)) {
		body["updateQuoteList"] = request.UpdateQuoteList
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
		Action:      tea.String("UpdateKROfContent"),
		Version:     tea.String("okr_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/okr/keyResults/contents"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateKROfContentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更改KR内容
//
// @param request - UpdateKROfContentRequest
//
// @return UpdateKROfContentResponse
func (client *Client) UpdateKROfContent(request *UpdateKROfContentRequest) (_result *UpdateKROfContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateKROfContentHeaders{}
	_result = &UpdateKROfContentResponse{}
	_body, _err := client.UpdateKROfContentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更改KR分数
//
// @param request - UpdateKROfScoreRequest
//
// @param headers - UpdateKROfScoreHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKROfScoreResponse
func (client *Client) UpdateKROfScoreWithOptions(request *UpdateKROfScoreRequest, headers *UpdateKROfScoreHeaders, runtime *util.RuntimeOptions) (_result *UpdateKROfScoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KrId)) {
		query["krId"] = request.KrId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Score)) {
		body["score"] = request.Score
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
		Action:      tea.String("UpdateKROfScore"),
		Version:     tea.String("okr_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/okr/keyResults/scores"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateKROfScoreResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更改KR分数
//
// @param request - UpdateKROfScoreRequest
//
// @return UpdateKROfScoreResponse
func (client *Client) UpdateKROfScore(request *UpdateKROfScoreRequest) (_result *UpdateKROfScoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateKROfScoreHeaders{}
	_result = &UpdateKROfScoreResponse{}
	_body, _err := client.UpdateKROfScoreWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更改 KR 权重
//
// @param request - UpdateKROfWeightRequest
//
// @param headers - UpdateKROfWeightHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKROfWeightResponse
func (client *Client) UpdateKROfWeightWithOptions(request *UpdateKROfWeightRequest, headers *UpdateKROfWeightHeaders, runtime *util.RuntimeOptions) (_result *UpdateKROfWeightResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KrId)) {
		query["krId"] = request.KrId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Weight)) {
		body["weight"] = request.Weight
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
		Action:      tea.String("UpdateKROfWeight"),
		Version:     tea.String("okr_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/okr/keyResults/weights"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateKROfWeightResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更改 KR 权重
//
// @param request - UpdateKROfWeightRequest
//
// @return UpdateKROfWeightResponse
func (client *Client) UpdateKROfWeight(request *UpdateKROfWeightRequest) (_result *UpdateKROfWeightResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateKROfWeightHeaders{}
	_result = &UpdateKROfWeightResponse{}
	_body, _err := client.UpdateKROfWeightWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新目标
//
// @param request - UpdateObjectiveRequest
//
// @param headers - UpdateObjectiveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateObjectiveResponse
func (client *Client) UpdateObjectiveWithOptions(objectiveId *string, request *UpdateObjectiveRequest, headers *UpdateObjectiveHeaders, runtime *util.RuntimeOptions) (_result *UpdateObjectiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateObjective"),
		Version:     tea.String("okr_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/okr/objectives/" + tea.StringValue(objectiveId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateObjectiveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新目标
//
// @param request - UpdateObjectiveRequest
//
// @return UpdateObjectiveResponse
func (client *Client) UpdateObjective(objectiveId *string, request *UpdateObjectiveRequest) (_result *UpdateObjectiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateObjectiveHeaders{}
	_result = &UpdateObjectiveResponse{}
	_body, _err := client.UpdateObjectiveWithOptions(objectiveId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新资源隐私策略
//
// @param request - UpdatePrivacyRequest
//
// @param headers - UpdatePrivacyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrivacyResponse
func (client *Client) UpdatePrivacyWithOptions(request *UpdatePrivacyRequest, headers *UpdatePrivacyHeaders, runtime *util.RuntimeOptions) (_result *UpdatePrivacyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Privacy)) {
		body["privacy"] = request.Privacy
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		body["targetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		body["targetType"] = request.TargetType
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
		Action:      tea.String("UpdatePrivacy"),
		Version:     tea.String("okr_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/okr/permissions/privacies"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePrivacyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新资源隐私策略
//
// @param request - UpdatePrivacyRequest
//
// @return UpdatePrivacyResponse
func (client *Client) UpdatePrivacy(request *UpdatePrivacyRequest) (_result *UpdatePrivacyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdatePrivacyHeaders{}
	_result = &UpdatePrivacyResponse{}
	_body, _err := client.UpdatePrivacyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
