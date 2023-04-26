// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package hrm_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddHrmPreentryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddHrmPreentryHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryHeaders) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryHeaders) SetCommonHeaders(v map[string]*string) *AddHrmPreentryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddHrmPreentryHeaders) SetXAcsDingtalkAccessToken(v string) *AddHrmPreentryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddHrmPreentryRequest struct {
	AgentId             *int64                         `json:"agentId,omitempty" xml:"agentId,omitempty"`
	Groups              []*AddHrmPreentryRequestGroups `json:"groups,omitempty" xml:"groups,omitempty" type:"Repeated"`
	Mobile              *string                        `json:"mobile,omitempty" xml:"mobile,omitempty"`
	Name                *string                        `json:"name,omitempty" xml:"name,omitempty"`
	NeedSendPreEntryMsg *bool                          `json:"needSendPreEntryMsg,omitempty" xml:"needSendPreEntryMsg,omitempty"`
	PreEntryTime        *int64                         `json:"preEntryTime,omitempty" xml:"preEntryTime,omitempty"`
}

func (s AddHrmPreentryRequest) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryRequest) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryRequest) SetAgentId(v int64) *AddHrmPreentryRequest {
	s.AgentId = &v
	return s
}

func (s *AddHrmPreentryRequest) SetGroups(v []*AddHrmPreentryRequestGroups) *AddHrmPreentryRequest {
	s.Groups = v
	return s
}

func (s *AddHrmPreentryRequest) SetMobile(v string) *AddHrmPreentryRequest {
	s.Mobile = &v
	return s
}

func (s *AddHrmPreentryRequest) SetName(v string) *AddHrmPreentryRequest {
	s.Name = &v
	return s
}

func (s *AddHrmPreentryRequest) SetNeedSendPreEntryMsg(v bool) *AddHrmPreentryRequest {
	s.NeedSendPreEntryMsg = &v
	return s
}

func (s *AddHrmPreentryRequest) SetPreEntryTime(v int64) *AddHrmPreentryRequest {
	s.PreEntryTime = &v
	return s
}

type AddHrmPreentryRequestGroups struct {
	GroupId  *string                                `json:"groupId,omitempty" xml:"groupId,omitempty"`
	Sections []*AddHrmPreentryRequestGroupsSections `json:"sections,omitempty" xml:"sections,omitempty" type:"Repeated"`
}

func (s AddHrmPreentryRequestGroups) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryRequestGroups) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryRequestGroups) SetGroupId(v string) *AddHrmPreentryRequestGroups {
	s.GroupId = &v
	return s
}

func (s *AddHrmPreentryRequestGroups) SetSections(v []*AddHrmPreentryRequestGroupsSections) *AddHrmPreentryRequestGroups {
	s.Sections = v
	return s
}

type AddHrmPreentryRequestGroupsSections struct {
	EmpFieldVOList []*AddHrmPreentryRequestGroupsSectionsEmpFieldVOList `json:"empFieldVOList,omitempty" xml:"empFieldVOList,omitempty" type:"Repeated"`
	OldIndex       *int32                                               `json:"oldIndex,omitempty" xml:"oldIndex,omitempty"`
}

func (s AddHrmPreentryRequestGroupsSections) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryRequestGroupsSections) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryRequestGroupsSections) SetEmpFieldVOList(v []*AddHrmPreentryRequestGroupsSectionsEmpFieldVOList) *AddHrmPreentryRequestGroupsSections {
	s.EmpFieldVOList = v
	return s
}

func (s *AddHrmPreentryRequestGroupsSections) SetOldIndex(v int32) *AddHrmPreentryRequestGroupsSections {
	s.OldIndex = &v
	return s
}

type AddHrmPreentryRequestGroupsSectionsEmpFieldVOList struct {
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	Value     *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AddHrmPreentryRequestGroupsSectionsEmpFieldVOList) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryRequestGroupsSectionsEmpFieldVOList) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryRequestGroupsSectionsEmpFieldVOList) SetFieldCode(v string) *AddHrmPreentryRequestGroupsSectionsEmpFieldVOList {
	s.FieldCode = &v
	return s
}

func (s *AddHrmPreentryRequestGroupsSectionsEmpFieldVOList) SetValue(v string) *AddHrmPreentryRequestGroupsSectionsEmpFieldVOList {
	s.Value = &v
	return s
}

type AddHrmPreentryResponseBody struct {
	TmpUserId *string `json:"tmpUserId,omitempty" xml:"tmpUserId,omitempty"`
}

func (s AddHrmPreentryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryResponseBody) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryResponseBody) SetTmpUserId(v string) *AddHrmPreentryResponseBody {
	s.TmpUserId = &v
	return s
}

type AddHrmPreentryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddHrmPreentryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddHrmPreentryResponse) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryResponse) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryResponse) SetHeaders(v map[string]*string) *AddHrmPreentryResponse {
	s.Headers = v
	return s
}

func (s *AddHrmPreentryResponse) SetStatusCode(v int32) *AddHrmPreentryResponse {
	s.StatusCode = &v
	return s
}

func (s *AddHrmPreentryResponse) SetBody(v *AddHrmPreentryResponseBody) *AddHrmPreentryResponse {
	s.Body = v
	return s
}

type ECertQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ECertQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s ECertQueryHeaders) GoString() string {
	return s.String()
}

func (s *ECertQueryHeaders) SetCommonHeaders(v map[string]*string) *ECertQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ECertQueryHeaders) SetXAcsDingtalkAccessToken(v string) *ECertQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ECertQueryRequest struct {
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ECertQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s ECertQueryRequest) GoString() string {
	return s.String()
}

func (s *ECertQueryRequest) SetUserId(v string) *ECertQueryRequest {
	s.UserId = &v
	return s
}

type ECertQueryResponseBody struct {
	CertNO                     *string   `json:"certNO,omitempty" xml:"certNO,omitempty"`
	EmployJobId                *string   `json:"employJobId,omitempty" xml:"employJobId,omitempty"`
	EmployJobIdLabel           *string   `json:"employJobIdLabel,omitempty" xml:"employJobIdLabel,omitempty"`
	EmployPositionId           *string   `json:"employPositionId,omitempty" xml:"employPositionId,omitempty"`
	EmployPositionIdLabel      *string   `json:"employPositionIdLabel,omitempty" xml:"employPositionIdLabel,omitempty"`
	EmployPositionRankId       *string   `json:"employPositionRankId,omitempty" xml:"employPositionRankId,omitempty"`
	EmployPositionRankIdLabel  *string   `json:"employPositionRankIdLabel,omitempty" xml:"employPositionRankIdLabel,omitempty"`
	HiredDate                  *string   `json:"hiredDate,omitempty" xml:"hiredDate,omitempty"`
	LastWorkDay                *string   `json:"lastWorkDay,omitempty" xml:"lastWorkDay,omitempty"`
	MainDeptId                 *int64    `json:"mainDeptId,omitempty" xml:"mainDeptId,omitempty"`
	MainDeptName               *string   `json:"mainDeptName,omitempty" xml:"mainDeptName,omitempty"`
	Name                       *string   `json:"name,omitempty" xml:"name,omitempty"`
	RealName                   *string   `json:"realName,omitempty" xml:"realName,omitempty"`
	TerminationReasonPassive   []*string `json:"terminationReasonPassive,omitempty" xml:"terminationReasonPassive,omitempty" type:"Repeated"`
	TerminationReasonVoluntary []*string `json:"terminationReasonVoluntary,omitempty" xml:"terminationReasonVoluntary,omitempty" type:"Repeated"`
}

func (s ECertQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ECertQueryResponseBody) GoString() string {
	return s.String()
}

func (s *ECertQueryResponseBody) SetCertNO(v string) *ECertQueryResponseBody {
	s.CertNO = &v
	return s
}

func (s *ECertQueryResponseBody) SetEmployJobId(v string) *ECertQueryResponseBody {
	s.EmployJobId = &v
	return s
}

func (s *ECertQueryResponseBody) SetEmployJobIdLabel(v string) *ECertQueryResponseBody {
	s.EmployJobIdLabel = &v
	return s
}

func (s *ECertQueryResponseBody) SetEmployPositionId(v string) *ECertQueryResponseBody {
	s.EmployPositionId = &v
	return s
}

func (s *ECertQueryResponseBody) SetEmployPositionIdLabel(v string) *ECertQueryResponseBody {
	s.EmployPositionIdLabel = &v
	return s
}

func (s *ECertQueryResponseBody) SetEmployPositionRankId(v string) *ECertQueryResponseBody {
	s.EmployPositionRankId = &v
	return s
}

func (s *ECertQueryResponseBody) SetEmployPositionRankIdLabel(v string) *ECertQueryResponseBody {
	s.EmployPositionRankIdLabel = &v
	return s
}

func (s *ECertQueryResponseBody) SetHiredDate(v string) *ECertQueryResponseBody {
	s.HiredDate = &v
	return s
}

func (s *ECertQueryResponseBody) SetLastWorkDay(v string) *ECertQueryResponseBody {
	s.LastWorkDay = &v
	return s
}

func (s *ECertQueryResponseBody) SetMainDeptId(v int64) *ECertQueryResponseBody {
	s.MainDeptId = &v
	return s
}

func (s *ECertQueryResponseBody) SetMainDeptName(v string) *ECertQueryResponseBody {
	s.MainDeptName = &v
	return s
}

func (s *ECertQueryResponseBody) SetName(v string) *ECertQueryResponseBody {
	s.Name = &v
	return s
}

func (s *ECertQueryResponseBody) SetRealName(v string) *ECertQueryResponseBody {
	s.RealName = &v
	return s
}

func (s *ECertQueryResponseBody) SetTerminationReasonPassive(v []*string) *ECertQueryResponseBody {
	s.TerminationReasonPassive = v
	return s
}

func (s *ECertQueryResponseBody) SetTerminationReasonVoluntary(v []*string) *ECertQueryResponseBody {
	s.TerminationReasonVoluntary = v
	return s
}

type ECertQueryResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ECertQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ECertQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s ECertQueryResponse) GoString() string {
	return s.String()
}

func (s *ECertQueryResponse) SetHeaders(v map[string]*string) *ECertQueryResponse {
	s.Headers = v
	return s
}

func (s *ECertQueryResponse) SetStatusCode(v int32) *ECertQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *ECertQueryResponse) SetBody(v *ECertQueryResponseBody) *ECertQueryResponse {
	s.Body = v
	return s
}

type HrmProcessRegularHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrmProcessRegularHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessRegularHeaders) GoString() string {
	return s.String()
}

func (s *HrmProcessRegularHeaders) SetCommonHeaders(v map[string]*string) *HrmProcessRegularHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrmProcessRegularHeaders) SetXAcsDingtalkAccessToken(v string) *HrmProcessRegularHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrmProcessRegularRequest struct {
	OperationId *string `json:"operationId,omitempty" xml:"operationId,omitempty"`
	RegularDate *int64  `json:"regularDate,omitempty" xml:"regularDate,omitempty"`
	Remark      *string `json:"remark,omitempty" xml:"remark,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s HrmProcessRegularRequest) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessRegularRequest) GoString() string {
	return s.String()
}

func (s *HrmProcessRegularRequest) SetOperationId(v string) *HrmProcessRegularRequest {
	s.OperationId = &v
	return s
}

func (s *HrmProcessRegularRequest) SetRegularDate(v int64) *HrmProcessRegularRequest {
	s.RegularDate = &v
	return s
}

func (s *HrmProcessRegularRequest) SetRemark(v string) *HrmProcessRegularRequest {
	s.Remark = &v
	return s
}

func (s *HrmProcessRegularRequest) SetUserId(v string) *HrmProcessRegularRequest {
	s.UserId = &v
	return s
}

type HrmProcessRegularResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s HrmProcessRegularResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessRegularResponseBody) GoString() string {
	return s.String()
}

func (s *HrmProcessRegularResponseBody) SetResult(v bool) *HrmProcessRegularResponseBody {
	s.Result = &v
	return s
}

type HrmProcessRegularResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *HrmProcessRegularResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HrmProcessRegularResponse) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessRegularResponse) GoString() string {
	return s.String()
}

func (s *HrmProcessRegularResponse) SetHeaders(v map[string]*string) *HrmProcessRegularResponse {
	s.Headers = v
	return s
}

func (s *HrmProcessRegularResponse) SetStatusCode(v int32) *HrmProcessRegularResponse {
	s.StatusCode = &v
	return s
}

func (s *HrmProcessRegularResponse) SetBody(v *HrmProcessRegularResponseBody) *HrmProcessRegularResponse {
	s.Body = v
	return s
}

type HrmProcessTransferHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrmProcessTransferHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessTransferHeaders) GoString() string {
	return s.String()
}

func (s *HrmProcessTransferHeaders) SetCommonHeaders(v map[string]*string) *HrmProcessTransferHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrmProcessTransferHeaders) SetXAcsDingtalkAccessToken(v string) *HrmProcessTransferHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrmProcessTransferRequest struct {
	DeptIdsAfterTransfer       []*int64 `json:"deptIdsAfterTransfer,omitempty" xml:"deptIdsAfterTransfer,omitempty" type:"Repeated"`
	JobIdAfterTransfer         *string  `json:"jobIdAfterTransfer,omitempty" xml:"jobIdAfterTransfer,omitempty"`
	MainDeptIdAfterTransfer    *int64   `json:"mainDeptIdAfterTransfer,omitempty" xml:"mainDeptIdAfterTransfer,omitempty"`
	OperateUserId              *string  `json:"operateUserId,omitempty" xml:"operateUserId,omitempty"`
	PositionIdAfterTransfer    *string  `json:"positionIdAfterTransfer,omitempty" xml:"positionIdAfterTransfer,omitempty"`
	PositionLevelAfterTransfer *string  `json:"positionLevelAfterTransfer,omitempty" xml:"positionLevelAfterTransfer,omitempty"`
	PositionNameAfterTransfer  *string  `json:"positionNameAfterTransfer,omitempty" xml:"positionNameAfterTransfer,omitempty"`
	RankIdAfterTransfer        *string  `json:"rankIdAfterTransfer,omitempty" xml:"rankIdAfterTransfer,omitempty"`
	UserId                     *string  `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s HrmProcessTransferRequest) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessTransferRequest) GoString() string {
	return s.String()
}

func (s *HrmProcessTransferRequest) SetDeptIdsAfterTransfer(v []*int64) *HrmProcessTransferRequest {
	s.DeptIdsAfterTransfer = v
	return s
}

func (s *HrmProcessTransferRequest) SetJobIdAfterTransfer(v string) *HrmProcessTransferRequest {
	s.JobIdAfterTransfer = &v
	return s
}

func (s *HrmProcessTransferRequest) SetMainDeptIdAfterTransfer(v int64) *HrmProcessTransferRequest {
	s.MainDeptIdAfterTransfer = &v
	return s
}

func (s *HrmProcessTransferRequest) SetOperateUserId(v string) *HrmProcessTransferRequest {
	s.OperateUserId = &v
	return s
}

func (s *HrmProcessTransferRequest) SetPositionIdAfterTransfer(v string) *HrmProcessTransferRequest {
	s.PositionIdAfterTransfer = &v
	return s
}

func (s *HrmProcessTransferRequest) SetPositionLevelAfterTransfer(v string) *HrmProcessTransferRequest {
	s.PositionLevelAfterTransfer = &v
	return s
}

func (s *HrmProcessTransferRequest) SetPositionNameAfterTransfer(v string) *HrmProcessTransferRequest {
	s.PositionNameAfterTransfer = &v
	return s
}

func (s *HrmProcessTransferRequest) SetRankIdAfterTransfer(v string) *HrmProcessTransferRequest {
	s.RankIdAfterTransfer = &v
	return s
}

func (s *HrmProcessTransferRequest) SetUserId(v string) *HrmProcessTransferRequest {
	s.UserId = &v
	return s
}

type HrmProcessTransferResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s HrmProcessTransferResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessTransferResponseBody) GoString() string {
	return s.String()
}

func (s *HrmProcessTransferResponseBody) SetResult(v bool) *HrmProcessTransferResponseBody {
	s.Result = &v
	return s
}

type HrmProcessTransferResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *HrmProcessTransferResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HrmProcessTransferResponse) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessTransferResponse) GoString() string {
	return s.String()
}

func (s *HrmProcessTransferResponse) SetHeaders(v map[string]*string) *HrmProcessTransferResponse {
	s.Headers = v
	return s
}

func (s *HrmProcessTransferResponse) SetStatusCode(v int32) *HrmProcessTransferResponse {
	s.StatusCode = &v
	return s
}

func (s *HrmProcessTransferResponse) SetBody(v *HrmProcessTransferResponseBody) *HrmProcessTransferResponse {
	s.Body = v
	return s
}

type HrmProcessUpdateTerminationInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrmProcessUpdateTerminationInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessUpdateTerminationInfoHeaders) GoString() string {
	return s.String()
}

func (s *HrmProcessUpdateTerminationInfoHeaders) SetCommonHeaders(v map[string]*string) *HrmProcessUpdateTerminationInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrmProcessUpdateTerminationInfoHeaders) SetXAcsDingtalkAccessToken(v string) *HrmProcessUpdateTerminationInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrmProcessUpdateTerminationInfoRequest struct {
	DismissionMemo *string `json:"dismissionMemo,omitempty" xml:"dismissionMemo,omitempty"`
	LastWorkDate   *int64  `json:"lastWorkDate,omitempty" xml:"lastWorkDate,omitempty"`
	UserId         *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s HrmProcessUpdateTerminationInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessUpdateTerminationInfoRequest) GoString() string {
	return s.String()
}

func (s *HrmProcessUpdateTerminationInfoRequest) SetDismissionMemo(v string) *HrmProcessUpdateTerminationInfoRequest {
	s.DismissionMemo = &v
	return s
}

func (s *HrmProcessUpdateTerminationInfoRequest) SetLastWorkDate(v int64) *HrmProcessUpdateTerminationInfoRequest {
	s.LastWorkDate = &v
	return s
}

func (s *HrmProcessUpdateTerminationInfoRequest) SetUserId(v string) *HrmProcessUpdateTerminationInfoRequest {
	s.UserId = &v
	return s
}

type HrmProcessUpdateTerminationInfoResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s HrmProcessUpdateTerminationInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessUpdateTerminationInfoResponseBody) GoString() string {
	return s.String()
}

func (s *HrmProcessUpdateTerminationInfoResponseBody) SetResult(v bool) *HrmProcessUpdateTerminationInfoResponseBody {
	s.Result = &v
	return s
}

type HrmProcessUpdateTerminationInfoResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *HrmProcessUpdateTerminationInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HrmProcessUpdateTerminationInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessUpdateTerminationInfoResponse) GoString() string {
	return s.String()
}

func (s *HrmProcessUpdateTerminationInfoResponse) SetHeaders(v map[string]*string) *HrmProcessUpdateTerminationInfoResponse {
	s.Headers = v
	return s
}

func (s *HrmProcessUpdateTerminationInfoResponse) SetStatusCode(v int32) *HrmProcessUpdateTerminationInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *HrmProcessUpdateTerminationInfoResponse) SetBody(v *HrmProcessUpdateTerminationInfoResponseBody) *HrmProcessUpdateTerminationInfoResponse {
	s.Body = v
	return s
}

type MasterDataQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MasterDataQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryHeaders) GoString() string {
	return s.String()
}

func (s *MasterDataQueryHeaders) SetCommonHeaders(v map[string]*string) *MasterDataQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MasterDataQueryHeaders) SetXAcsDingtalkAccessToken(v string) *MasterDataQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MasterDataQueryRequest struct {
	BizUK          *string                              `json:"bizUK,omitempty" xml:"bizUK,omitempty"`
	MaxResults     *int32                               `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken      *int32                               `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OptUserId      *string                              `json:"optUserId,omitempty" xml:"optUserId,omitempty"`
	QueryParams    []*MasterDataQueryRequestQueryParams `json:"queryParams,omitempty" xml:"queryParams,omitempty" type:"Repeated"`
	RelationIds    []*string                            `json:"relationIds,omitempty" xml:"relationIds,omitempty" type:"Repeated"`
	ScopeCode      *string                              `json:"scopeCode,omitempty" xml:"scopeCode,omitempty"`
	TenantId       *int64                               `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	ViewEntityCode *string                              `json:"viewEntityCode,omitempty" xml:"viewEntityCode,omitempty"`
}

func (s MasterDataQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryRequest) GoString() string {
	return s.String()
}

func (s *MasterDataQueryRequest) SetBizUK(v string) *MasterDataQueryRequest {
	s.BizUK = &v
	return s
}

func (s *MasterDataQueryRequest) SetMaxResults(v int32) *MasterDataQueryRequest {
	s.MaxResults = &v
	return s
}

func (s *MasterDataQueryRequest) SetNextToken(v int32) *MasterDataQueryRequest {
	s.NextToken = &v
	return s
}

func (s *MasterDataQueryRequest) SetOptUserId(v string) *MasterDataQueryRequest {
	s.OptUserId = &v
	return s
}

func (s *MasterDataQueryRequest) SetQueryParams(v []*MasterDataQueryRequestQueryParams) *MasterDataQueryRequest {
	s.QueryParams = v
	return s
}

func (s *MasterDataQueryRequest) SetRelationIds(v []*string) *MasterDataQueryRequest {
	s.RelationIds = v
	return s
}

func (s *MasterDataQueryRequest) SetScopeCode(v string) *MasterDataQueryRequest {
	s.ScopeCode = &v
	return s
}

func (s *MasterDataQueryRequest) SetTenantId(v int64) *MasterDataQueryRequest {
	s.TenantId = &v
	return s
}

func (s *MasterDataQueryRequest) SetViewEntityCode(v string) *MasterDataQueryRequest {
	s.ViewEntityCode = &v
	return s
}

type MasterDataQueryRequestQueryParams struct {
	ConditionList []*MasterDataQueryRequestQueryParamsConditionList `json:"conditionList,omitempty" xml:"conditionList,omitempty" type:"Repeated"`
	FieldCode     *string                                           `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	JoinType      *string                                           `json:"joinType,omitempty" xml:"joinType,omitempty"`
}

func (s MasterDataQueryRequestQueryParams) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryRequestQueryParams) GoString() string {
	return s.String()
}

func (s *MasterDataQueryRequestQueryParams) SetConditionList(v []*MasterDataQueryRequestQueryParamsConditionList) *MasterDataQueryRequestQueryParams {
	s.ConditionList = v
	return s
}

func (s *MasterDataQueryRequestQueryParams) SetFieldCode(v string) *MasterDataQueryRequestQueryParams {
	s.FieldCode = &v
	return s
}

func (s *MasterDataQueryRequestQueryParams) SetJoinType(v string) *MasterDataQueryRequestQueryParams {
	s.JoinType = &v
	return s
}

type MasterDataQueryRequestQueryParamsConditionList struct {
	Operate *string `json:"operate,omitempty" xml:"operate,omitempty"`
	Value   *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s MasterDataQueryRequestQueryParamsConditionList) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryRequestQueryParamsConditionList) GoString() string {
	return s.String()
}

func (s *MasterDataQueryRequestQueryParamsConditionList) SetOperate(v string) *MasterDataQueryRequestQueryParamsConditionList {
	s.Operate = &v
	return s
}

func (s *MasterDataQueryRequestQueryParamsConditionList) SetValue(v string) *MasterDataQueryRequestQueryParamsConditionList {
	s.Value = &v
	return s
}

type MasterDataQueryResponseBody struct {
	HasMore   *bool                                `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken *int64                               `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Result    []*MasterDataQueryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success   *bool                                `json:"success,omitempty" xml:"success,omitempty"`
	Total     *int64                               `json:"total,omitempty" xml:"total,omitempty"`
}

func (s MasterDataQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryResponseBody) GoString() string {
	return s.String()
}

func (s *MasterDataQueryResponseBody) SetHasMore(v bool) *MasterDataQueryResponseBody {
	s.HasMore = &v
	return s
}

func (s *MasterDataQueryResponseBody) SetNextToken(v int64) *MasterDataQueryResponseBody {
	s.NextToken = &v
	return s
}

func (s *MasterDataQueryResponseBody) SetResult(v []*MasterDataQueryResponseBodyResult) *MasterDataQueryResponseBody {
	s.Result = v
	return s
}

func (s *MasterDataQueryResponseBody) SetSuccess(v bool) *MasterDataQueryResponseBody {
	s.Success = &v
	return s
}

func (s *MasterDataQueryResponseBody) SetTotal(v int64) *MasterDataQueryResponseBody {
	s.Total = &v
	return s
}

type MasterDataQueryResponseBodyResult struct {
	OuterId               *string                                                   `json:"outerId,omitempty" xml:"outerId,omitempty"`
	RelationId            *string                                                   `json:"relationId,omitempty" xml:"relationId,omitempty"`
	ScopeCode             *string                                                   `json:"scopeCode,omitempty" xml:"scopeCode,omitempty"`
	ViewEntityCode        *string                                                   `json:"viewEntityCode,omitempty" xml:"viewEntityCode,omitempty"`
	ViewEntityFieldVOList []*MasterDataQueryResponseBodyResultViewEntityFieldVOList `json:"viewEntityFieldVOList,omitempty" xml:"viewEntityFieldVOList,omitempty" type:"Repeated"`
}

func (s MasterDataQueryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *MasterDataQueryResponseBodyResult) SetOuterId(v string) *MasterDataQueryResponseBodyResult {
	s.OuterId = &v
	return s
}

func (s *MasterDataQueryResponseBodyResult) SetRelationId(v string) *MasterDataQueryResponseBodyResult {
	s.RelationId = &v
	return s
}

func (s *MasterDataQueryResponseBodyResult) SetScopeCode(v string) *MasterDataQueryResponseBodyResult {
	s.ScopeCode = &v
	return s
}

func (s *MasterDataQueryResponseBodyResult) SetViewEntityCode(v string) *MasterDataQueryResponseBodyResult {
	s.ViewEntityCode = &v
	return s
}

func (s *MasterDataQueryResponseBodyResult) SetViewEntityFieldVOList(v []*MasterDataQueryResponseBodyResultViewEntityFieldVOList) *MasterDataQueryResponseBodyResult {
	s.ViewEntityFieldVOList = v
	return s
}

type MasterDataQueryResponseBodyResultViewEntityFieldVOList struct {
	FieldCode   *string                                                            `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	FieldDataVO *MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO `json:"fieldDataVO,omitempty" xml:"fieldDataVO,omitempty" type:"Struct"`
	FieldName   *string                                                            `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	FieldType   *string                                                            `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
}

func (s MasterDataQueryResponseBodyResultViewEntityFieldVOList) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryResponseBodyResultViewEntityFieldVOList) GoString() string {
	return s.String()
}

func (s *MasterDataQueryResponseBodyResultViewEntityFieldVOList) SetFieldCode(v string) *MasterDataQueryResponseBodyResultViewEntityFieldVOList {
	s.FieldCode = &v
	return s
}

func (s *MasterDataQueryResponseBodyResultViewEntityFieldVOList) SetFieldDataVO(v *MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO) *MasterDataQueryResponseBodyResultViewEntityFieldVOList {
	s.FieldDataVO = v
	return s
}

func (s *MasterDataQueryResponseBodyResultViewEntityFieldVOList) SetFieldName(v string) *MasterDataQueryResponseBodyResultViewEntityFieldVOList {
	s.FieldName = &v
	return s
}

func (s *MasterDataQueryResponseBodyResultViewEntityFieldVOList) SetFieldType(v string) *MasterDataQueryResponseBodyResultViewEntityFieldVOList {
	s.FieldType = &v
	return s
}

type MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO) GoString() string {
	return s.String()
}

func (s *MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO) SetKey(v string) *MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO {
	s.Key = &v
	return s
}

func (s *MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO) SetValue(v string) *MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO {
	s.Value = &v
	return s
}

type MasterDataQueryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MasterDataQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MasterDataQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryResponse) GoString() string {
	return s.String()
}

func (s *MasterDataQueryResponse) SetHeaders(v map[string]*string) *MasterDataQueryResponse {
	s.Headers = v
	return s
}

func (s *MasterDataQueryResponse) SetStatusCode(v int32) *MasterDataQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *MasterDataQueryResponse) SetBody(v *MasterDataQueryResponseBody) *MasterDataQueryResponse {
	s.Body = v
	return s
}

type MasterDataSaveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MasterDataSaveHeaders) String() string {
	return tea.Prettify(s)
}

func (s MasterDataSaveHeaders) GoString() string {
	return s.String()
}

func (s *MasterDataSaveHeaders) SetCommonHeaders(v map[string]*string) *MasterDataSaveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MasterDataSaveHeaders) SetXAcsDingtalkAccessToken(v string) *MasterDataSaveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MasterDataSaveRequest struct {
	Body     []*MasterDataSaveRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	TenantId *int64                       `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s MasterDataSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s MasterDataSaveRequest) GoString() string {
	return s.String()
}

func (s *MasterDataSaveRequest) SetBody(v []*MasterDataSaveRequestBody) *MasterDataSaveRequest {
	s.Body = v
	return s
}

func (s *MasterDataSaveRequest) SetTenantId(v int64) *MasterDataSaveRequest {
	s.TenantId = &v
	return s
}

type MasterDataSaveRequestBody struct {
	BizTime    *int64                                `json:"bizTime,omitempty" xml:"bizTime,omitempty"`
	BizUk      *string                               `json:"bizUk,omitempty" xml:"bizUk,omitempty"`
	EntityCode *string                               `json:"entityCode,omitempty" xml:"entityCode,omitempty"`
	FieldList  []*MasterDataSaveRequestBodyFieldList `json:"fieldList,omitempty" xml:"fieldList,omitempty" type:"Repeated"`
	Scope      *MasterDataSaveRequestBodyScope       `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
	UserId     *string                               `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s MasterDataSaveRequestBody) String() string {
	return tea.Prettify(s)
}

func (s MasterDataSaveRequestBody) GoString() string {
	return s.String()
}

func (s *MasterDataSaveRequestBody) SetBizTime(v int64) *MasterDataSaveRequestBody {
	s.BizTime = &v
	return s
}

func (s *MasterDataSaveRequestBody) SetBizUk(v string) *MasterDataSaveRequestBody {
	s.BizUk = &v
	return s
}

func (s *MasterDataSaveRequestBody) SetEntityCode(v string) *MasterDataSaveRequestBody {
	s.EntityCode = &v
	return s
}

func (s *MasterDataSaveRequestBody) SetFieldList(v []*MasterDataSaveRequestBodyFieldList) *MasterDataSaveRequestBody {
	s.FieldList = v
	return s
}

func (s *MasterDataSaveRequestBody) SetScope(v *MasterDataSaveRequestBodyScope) *MasterDataSaveRequestBody {
	s.Scope = v
	return s
}

func (s *MasterDataSaveRequestBody) SetUserId(v string) *MasterDataSaveRequestBody {
	s.UserId = &v
	return s
}

type MasterDataSaveRequestBodyFieldList struct {
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	ValueStr *string `json:"valueStr,omitempty" xml:"valueStr,omitempty"`
}

func (s MasterDataSaveRequestBodyFieldList) String() string {
	return tea.Prettify(s)
}

func (s MasterDataSaveRequestBodyFieldList) GoString() string {
	return s.String()
}

func (s *MasterDataSaveRequestBodyFieldList) SetName(v string) *MasterDataSaveRequestBodyFieldList {
	s.Name = &v
	return s
}

func (s *MasterDataSaveRequestBodyFieldList) SetValueStr(v string) *MasterDataSaveRequestBodyFieldList {
	s.ValueStr = &v
	return s
}

type MasterDataSaveRequestBodyScope struct {
	ScopeCode *string `json:"scopeCode,omitempty" xml:"scopeCode,omitempty"`
	Version   *int32  `json:"version,omitempty" xml:"version,omitempty"`
}

func (s MasterDataSaveRequestBodyScope) String() string {
	return tea.Prettify(s)
}

func (s MasterDataSaveRequestBodyScope) GoString() string {
	return s.String()
}

func (s *MasterDataSaveRequestBodyScope) SetScopeCode(v string) *MasterDataSaveRequestBodyScope {
	s.ScopeCode = &v
	return s
}

func (s *MasterDataSaveRequestBodyScope) SetVersion(v int32) *MasterDataSaveRequestBodyScope {
	s.Version = &v
	return s
}

type MasterDataSaveResponseBody struct {
	AllSuccess *bool                                   `json:"allSuccess,omitempty" xml:"allSuccess,omitempty"`
	FailResult []*MasterDataSaveResponseBodyFailResult `json:"failResult,omitempty" xml:"failResult,omitempty" type:"Repeated"`
}

func (s MasterDataSaveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MasterDataSaveResponseBody) GoString() string {
	return s.String()
}

func (s *MasterDataSaveResponseBody) SetAllSuccess(v bool) *MasterDataSaveResponseBody {
	s.AllSuccess = &v
	return s
}

func (s *MasterDataSaveResponseBody) SetFailResult(v []*MasterDataSaveResponseBodyFailResult) *MasterDataSaveResponseBody {
	s.FailResult = v
	return s
}

type MasterDataSaveResponseBodyFailResult struct {
	BizUk     *string `json:"bizUk,omitempty" xml:"bizUk,omitempty"`
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s MasterDataSaveResponseBodyFailResult) String() string {
	return tea.Prettify(s)
}

func (s MasterDataSaveResponseBodyFailResult) GoString() string {
	return s.String()
}

func (s *MasterDataSaveResponseBodyFailResult) SetBizUk(v string) *MasterDataSaveResponseBodyFailResult {
	s.BizUk = &v
	return s
}

func (s *MasterDataSaveResponseBodyFailResult) SetErrorCode(v string) *MasterDataSaveResponseBodyFailResult {
	s.ErrorCode = &v
	return s
}

func (s *MasterDataSaveResponseBodyFailResult) SetErrorMsg(v string) *MasterDataSaveResponseBodyFailResult {
	s.ErrorMsg = &v
	return s
}

func (s *MasterDataSaveResponseBodyFailResult) SetSuccess(v bool) *MasterDataSaveResponseBodyFailResult {
	s.Success = &v
	return s
}

type MasterDataSaveResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MasterDataSaveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MasterDataSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s MasterDataSaveResponse) GoString() string {
	return s.String()
}

func (s *MasterDataSaveResponse) SetHeaders(v map[string]*string) *MasterDataSaveResponse {
	s.Headers = v
	return s
}

func (s *MasterDataSaveResponse) SetStatusCode(v int32) *MasterDataSaveResponse {
	s.StatusCode = &v
	return s
}

func (s *MasterDataSaveResponse) SetBody(v *MasterDataSaveResponseBody) *MasterDataSaveResponse {
	s.Body = v
	return s
}

type MasterDataTenantQueyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MasterDataTenantQueyHeaders) String() string {
	return tea.Prettify(s)
}

func (s MasterDataTenantQueyHeaders) GoString() string {
	return s.String()
}

func (s *MasterDataTenantQueyHeaders) SetCommonHeaders(v map[string]*string) *MasterDataTenantQueyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MasterDataTenantQueyHeaders) SetXAcsDingtalkAccessToken(v string) *MasterDataTenantQueyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MasterDataTenantQueyRequest struct {
	EntityCode *string `json:"entityCode,omitempty" xml:"entityCode,omitempty"`
	ScopeCode  *string `json:"scopeCode,omitempty" xml:"scopeCode,omitempty"`
}

func (s MasterDataTenantQueyRequest) String() string {
	return tea.Prettify(s)
}

func (s MasterDataTenantQueyRequest) GoString() string {
	return s.String()
}

func (s *MasterDataTenantQueyRequest) SetEntityCode(v string) *MasterDataTenantQueyRequest {
	s.EntityCode = &v
	return s
}

func (s *MasterDataTenantQueyRequest) SetScopeCode(v string) *MasterDataTenantQueyRequest {
	s.ScopeCode = &v
	return s
}

type MasterDataTenantQueyResponseBody struct {
	Result []*MasterDataTenantQueyResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s MasterDataTenantQueyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MasterDataTenantQueyResponseBody) GoString() string {
	return s.String()
}

func (s *MasterDataTenantQueyResponseBody) SetResult(v []*MasterDataTenantQueyResponseBodyResult) *MasterDataTenantQueyResponseBody {
	s.Result = v
	return s
}

type MasterDataTenantQueyResponseBodyResult struct {
	HasData           *bool   `json:"hasData,omitempty" xml:"hasData,omitempty"`
	IntegrateDataAuth *bool   `json:"integrateDataAuth,omitempty" xml:"integrateDataAuth,omitempty"`
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	ReadAuth          *bool   `json:"readAuth,omitempty" xml:"readAuth,omitempty"`
	TenantId          *int64  `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	Type              *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s MasterDataTenantQueyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s MasterDataTenantQueyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *MasterDataTenantQueyResponseBodyResult) SetHasData(v bool) *MasterDataTenantQueyResponseBodyResult {
	s.HasData = &v
	return s
}

func (s *MasterDataTenantQueyResponseBodyResult) SetIntegrateDataAuth(v bool) *MasterDataTenantQueyResponseBodyResult {
	s.IntegrateDataAuth = &v
	return s
}

func (s *MasterDataTenantQueyResponseBodyResult) SetName(v string) *MasterDataTenantQueyResponseBodyResult {
	s.Name = &v
	return s
}

func (s *MasterDataTenantQueyResponseBodyResult) SetReadAuth(v bool) *MasterDataTenantQueyResponseBodyResult {
	s.ReadAuth = &v
	return s
}

func (s *MasterDataTenantQueyResponseBodyResult) SetTenantId(v int64) *MasterDataTenantQueyResponseBodyResult {
	s.TenantId = &v
	return s
}

func (s *MasterDataTenantQueyResponseBodyResult) SetType(v int32) *MasterDataTenantQueyResponseBodyResult {
	s.Type = &v
	return s
}

type MasterDataTenantQueyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MasterDataTenantQueyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MasterDataTenantQueyResponse) String() string {
	return tea.Prettify(s)
}

func (s MasterDataTenantQueyResponse) GoString() string {
	return s.String()
}

func (s *MasterDataTenantQueyResponse) SetHeaders(v map[string]*string) *MasterDataTenantQueyResponse {
	s.Headers = v
	return s
}

func (s *MasterDataTenantQueyResponse) SetStatusCode(v int32) *MasterDataTenantQueyResponse {
	s.StatusCode = &v
	return s
}

func (s *MasterDataTenantQueyResponse) SetBody(v *MasterDataTenantQueyResponseBody) *MasterDataTenantQueyResponse {
	s.Body = v
	return s
}

type QueryCustomEntryProcessesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCustomEntryProcessesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomEntryProcessesHeaders) GoString() string {
	return s.String()
}

func (s *QueryCustomEntryProcessesHeaders) SetCommonHeaders(v map[string]*string) *QueryCustomEntryProcessesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCustomEntryProcessesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCustomEntryProcessesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCustomEntryProcessesRequest struct {
	MaxResults    *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken     *int32  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OperateUserId *string `json:"operateUserId,omitempty" xml:"operateUserId,omitempty"`
}

func (s QueryCustomEntryProcessesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomEntryProcessesRequest) GoString() string {
	return s.String()
}

func (s *QueryCustomEntryProcessesRequest) SetMaxResults(v int32) *QueryCustomEntryProcessesRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryCustomEntryProcessesRequest) SetNextToken(v int32) *QueryCustomEntryProcessesRequest {
	s.NextToken = &v
	return s
}

func (s *QueryCustomEntryProcessesRequest) SetOperateUserId(v string) *QueryCustomEntryProcessesRequest {
	s.OperateUserId = &v
	return s
}

type QueryCustomEntryProcessesResponseBody struct {
	HasMore   *bool                                        `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List      []*QueryCustomEntryProcessesResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextToken *int64                                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryCustomEntryProcessesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomEntryProcessesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCustomEntryProcessesResponseBody) SetHasMore(v bool) *QueryCustomEntryProcessesResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryCustomEntryProcessesResponseBody) SetList(v []*QueryCustomEntryProcessesResponseBodyList) *QueryCustomEntryProcessesResponseBody {
	s.List = v
	return s
}

func (s *QueryCustomEntryProcessesResponseBody) SetNextToken(v int64) *QueryCustomEntryProcessesResponseBody {
	s.NextToken = &v
	return s
}

type QueryCustomEntryProcessesResponseBodyList struct {
	FormDesc *string `json:"formDesc,omitempty" xml:"formDesc,omitempty"`
	FormId   *string `json:"formId,omitempty" xml:"formId,omitempty"`
	FormName *string `json:"formName,omitempty" xml:"formName,omitempty"`
	ShortUrl *string `json:"shortUrl,omitempty" xml:"shortUrl,omitempty"`
}

func (s QueryCustomEntryProcessesResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomEntryProcessesResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryCustomEntryProcessesResponseBodyList) SetFormDesc(v string) *QueryCustomEntryProcessesResponseBodyList {
	s.FormDesc = &v
	return s
}

func (s *QueryCustomEntryProcessesResponseBodyList) SetFormId(v string) *QueryCustomEntryProcessesResponseBodyList {
	s.FormId = &v
	return s
}

func (s *QueryCustomEntryProcessesResponseBodyList) SetFormName(v string) *QueryCustomEntryProcessesResponseBodyList {
	s.FormName = &v
	return s
}

func (s *QueryCustomEntryProcessesResponseBodyList) SetShortUrl(v string) *QueryCustomEntryProcessesResponseBodyList {
	s.ShortUrl = &v
	return s
}

type QueryCustomEntryProcessesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCustomEntryProcessesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCustomEntryProcessesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomEntryProcessesResponse) GoString() string {
	return s.String()
}

func (s *QueryCustomEntryProcessesResponse) SetHeaders(v map[string]*string) *QueryCustomEntryProcessesResponse {
	s.Headers = v
	return s
}

func (s *QueryCustomEntryProcessesResponse) SetStatusCode(v int32) *QueryCustomEntryProcessesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCustomEntryProcessesResponse) SetBody(v *QueryCustomEntryProcessesResponseBody) *QueryCustomEntryProcessesResponse {
	s.Body = v
	return s
}

type QueryDismissionStaffIdListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDismissionStaffIdListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDismissionStaffIdListHeaders) GoString() string {
	return s.String()
}

func (s *QueryDismissionStaffIdListHeaders) SetCommonHeaders(v map[string]*string) *QueryDismissionStaffIdListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDismissionStaffIdListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDismissionStaffIdListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDismissionStaffIdListRequest struct {
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryDismissionStaffIdListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDismissionStaffIdListRequest) GoString() string {
	return s.String()
}

func (s *QueryDismissionStaffIdListRequest) SetMaxResults(v int32) *QueryDismissionStaffIdListRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryDismissionStaffIdListRequest) SetNextToken(v int64) *QueryDismissionStaffIdListRequest {
	s.NextToken = &v
	return s
}

type QueryDismissionStaffIdListResponseBody struct {
	HasMore    *bool     `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken  *int64    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s QueryDismissionStaffIdListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDismissionStaffIdListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDismissionStaffIdListResponseBody) SetHasMore(v bool) *QueryDismissionStaffIdListResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryDismissionStaffIdListResponseBody) SetNextToken(v int64) *QueryDismissionStaffIdListResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryDismissionStaffIdListResponseBody) SetUserIdList(v []*string) *QueryDismissionStaffIdListResponseBody {
	s.UserIdList = v
	return s
}

type QueryDismissionStaffIdListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDismissionStaffIdListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDismissionStaffIdListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDismissionStaffIdListResponse) GoString() string {
	return s.String()
}

func (s *QueryDismissionStaffIdListResponse) SetHeaders(v map[string]*string) *QueryDismissionStaffIdListResponse {
	s.Headers = v
	return s
}

func (s *QueryDismissionStaffIdListResponse) SetStatusCode(v int32) *QueryDismissionStaffIdListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDismissionStaffIdListResponse) SetBody(v *QueryDismissionStaffIdListResponseBody) *QueryDismissionStaffIdListResponse {
	s.Body = v
	return s
}

type QueryHrmEmployeeDismissionInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryHrmEmployeeDismissionInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryHrmEmployeeDismissionInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryHrmEmployeeDismissionInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryHrmEmployeeDismissionInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryHrmEmployeeDismissionInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryHrmEmployeeDismissionInfoRequest struct {
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s QueryHrmEmployeeDismissionInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHrmEmployeeDismissionInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryHrmEmployeeDismissionInfoRequest) SetUserIdList(v []*string) *QueryHrmEmployeeDismissionInfoRequest {
	s.UserIdList = v
	return s
}

type QueryHrmEmployeeDismissionInfoShrinkRequest struct {
	UserIdListShrink *string `json:"userIdList,omitempty" xml:"userIdList,omitempty"`
}

func (s QueryHrmEmployeeDismissionInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHrmEmployeeDismissionInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryHrmEmployeeDismissionInfoShrinkRequest) SetUserIdListShrink(v string) *QueryHrmEmployeeDismissionInfoShrinkRequest {
	s.UserIdListShrink = &v
	return s
}

type QueryHrmEmployeeDismissionInfoResponseBody struct {
	Result []*QueryHrmEmployeeDismissionInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryHrmEmployeeDismissionInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHrmEmployeeDismissionInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHrmEmployeeDismissionInfoResponseBody) SetResult(v []*QueryHrmEmployeeDismissionInfoResponseBodyResult) *QueryHrmEmployeeDismissionInfoResponseBody {
	s.Result = v
	return s
}

type QueryHrmEmployeeDismissionInfoResponseBodyResult struct {
	DeptList        []*QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList `json:"deptList,omitempty" xml:"deptList,omitempty" type:"Repeated"`
	HandoverUserId  *string                                                     `json:"handoverUserId,omitempty" xml:"handoverUserId,omitempty"`
	LastWorkDay     *int64                                                      `json:"lastWorkDay,omitempty" xml:"lastWorkDay,omitempty"`
	MainDeptId      *int64                                                      `json:"mainDeptId,omitempty" xml:"mainDeptId,omitempty"`
	MainDeptName    *string                                                     `json:"mainDeptName,omitempty" xml:"mainDeptName,omitempty"`
	PassiveReason   []*string                                                   `json:"passiveReason,omitempty" xml:"passiveReason,omitempty" type:"Repeated"`
	PreStatus       *int32                                                      `json:"preStatus,omitempty" xml:"preStatus,omitempty"`
	ReasonMemo      *string                                                     `json:"reasonMemo,omitempty" xml:"reasonMemo,omitempty"`
	Status          *int32                                                      `json:"status,omitempty" xml:"status,omitempty"`
	UserId          *string                                                     `json:"userId,omitempty" xml:"userId,omitempty"`
	VoluntaryReason []*string                                                   `json:"voluntaryReason,omitempty" xml:"voluntaryReason,omitempty" type:"Repeated"`
}

func (s QueryHrmEmployeeDismissionInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryHrmEmployeeDismissionInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetDeptList(v []*QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.DeptList = v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetHandoverUserId(v string) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.HandoverUserId = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetLastWorkDay(v int64) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.LastWorkDay = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetMainDeptId(v int64) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.MainDeptId = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetMainDeptName(v string) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.MainDeptName = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetPassiveReason(v []*string) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.PassiveReason = v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetPreStatus(v int32) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.PreStatus = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetReasonMemo(v string) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.ReasonMemo = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetStatus(v int32) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.Status = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetUserId(v string) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetVoluntaryReason(v []*string) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.VoluntaryReason = v
	return s
}

type QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList struct {
	DeptId   *int64  `json:"dept_id,omitempty" xml:"dept_id,omitempty"`
	DeptPath *string `json:"dept_path,omitempty" xml:"dept_path,omitempty"`
}

func (s QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList) String() string {
	return tea.Prettify(s)
}

func (s QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList) GoString() string {
	return s.String()
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList) SetDeptId(v int64) *QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList {
	s.DeptId = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList) SetDeptPath(v string) *QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList {
	s.DeptPath = &v
	return s
}

type QueryHrmEmployeeDismissionInfoResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryHrmEmployeeDismissionInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryHrmEmployeeDismissionInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHrmEmployeeDismissionInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryHrmEmployeeDismissionInfoResponse) SetHeaders(v map[string]*string) *QueryHrmEmployeeDismissionInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponse) SetStatusCode(v int32) *QueryHrmEmployeeDismissionInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponse) SetBody(v *QueryHrmEmployeeDismissionInfoResponseBody) *QueryHrmEmployeeDismissionInfoResponse {
	s.Body = v
	return s
}

type QueryJobRanksHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryJobRanksHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryJobRanksHeaders) GoString() string {
	return s.String()
}

func (s *QueryJobRanksHeaders) SetCommonHeaders(v map[string]*string) *QueryJobRanksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryJobRanksHeaders) SetXAcsDingtalkAccessToken(v string) *QueryJobRanksHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryJobRanksRequest struct {
	MaxResults     *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken      *int32  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RankCategoryId *string `json:"rankCategoryId,omitempty" xml:"rankCategoryId,omitempty"`
	RankCode       *string `json:"rankCode,omitempty" xml:"rankCode,omitempty"`
	RankName       *string `json:"rankName,omitempty" xml:"rankName,omitempty"`
}

func (s QueryJobRanksRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryJobRanksRequest) GoString() string {
	return s.String()
}

func (s *QueryJobRanksRequest) SetMaxResults(v int32) *QueryJobRanksRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryJobRanksRequest) SetNextToken(v int32) *QueryJobRanksRequest {
	s.NextToken = &v
	return s
}

func (s *QueryJobRanksRequest) SetRankCategoryId(v string) *QueryJobRanksRequest {
	s.RankCategoryId = &v
	return s
}

func (s *QueryJobRanksRequest) SetRankCode(v string) *QueryJobRanksRequest {
	s.RankCode = &v
	return s
}

func (s *QueryJobRanksRequest) SetRankName(v string) *QueryJobRanksRequest {
	s.RankName = &v
	return s
}

type QueryJobRanksResponseBody struct {
	HasMore   *bool                            `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List      []*QueryJobRanksResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextToken *int64                           `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryJobRanksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryJobRanksResponseBody) GoString() string {
	return s.String()
}

func (s *QueryJobRanksResponseBody) SetHasMore(v bool) *QueryJobRanksResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryJobRanksResponseBody) SetList(v []*QueryJobRanksResponseBodyList) *QueryJobRanksResponseBody {
	s.List = v
	return s
}

func (s *QueryJobRanksResponseBody) SetNextToken(v int64) *QueryJobRanksResponseBody {
	s.NextToken = &v
	return s
}

type QueryJobRanksResponseBodyList struct {
	MaxJobGrade     *int32  `json:"maxJobGrade,omitempty" xml:"maxJobGrade,omitempty"`
	MinJobGrade     *int32  `json:"minJobGrade,omitempty" xml:"minJobGrade,omitempty"`
	RankCategoryId  *string `json:"rankCategoryId,omitempty" xml:"rankCategoryId,omitempty"`
	RankCode        *string `json:"rankCode,omitempty" xml:"rankCode,omitempty"`
	RankDescription *string `json:"rankDescription,omitempty" xml:"rankDescription,omitempty"`
	RankId          *string `json:"rankId,omitempty" xml:"rankId,omitempty"`
	RankName        *string `json:"rankName,omitempty" xml:"rankName,omitempty"`
}

func (s QueryJobRanksResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryJobRanksResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryJobRanksResponseBodyList) SetMaxJobGrade(v int32) *QueryJobRanksResponseBodyList {
	s.MaxJobGrade = &v
	return s
}

func (s *QueryJobRanksResponseBodyList) SetMinJobGrade(v int32) *QueryJobRanksResponseBodyList {
	s.MinJobGrade = &v
	return s
}

func (s *QueryJobRanksResponseBodyList) SetRankCategoryId(v string) *QueryJobRanksResponseBodyList {
	s.RankCategoryId = &v
	return s
}

func (s *QueryJobRanksResponseBodyList) SetRankCode(v string) *QueryJobRanksResponseBodyList {
	s.RankCode = &v
	return s
}

func (s *QueryJobRanksResponseBodyList) SetRankDescription(v string) *QueryJobRanksResponseBodyList {
	s.RankDescription = &v
	return s
}

func (s *QueryJobRanksResponseBodyList) SetRankId(v string) *QueryJobRanksResponseBodyList {
	s.RankId = &v
	return s
}

func (s *QueryJobRanksResponseBodyList) SetRankName(v string) *QueryJobRanksResponseBodyList {
	s.RankName = &v
	return s
}

type QueryJobRanksResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryJobRanksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryJobRanksResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryJobRanksResponse) GoString() string {
	return s.String()
}

func (s *QueryJobRanksResponse) SetHeaders(v map[string]*string) *QueryJobRanksResponse {
	s.Headers = v
	return s
}

func (s *QueryJobRanksResponse) SetStatusCode(v int32) *QueryJobRanksResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryJobRanksResponse) SetBody(v *QueryJobRanksResponseBody) *QueryJobRanksResponse {
	s.Body = v
	return s
}

type QueryJobsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryJobsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryJobsHeaders) GoString() string {
	return s.String()
}

func (s *QueryJobsHeaders) SetCommonHeaders(v map[string]*string) *QueryJobsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryJobsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryJobsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryJobsRequest struct {
	JobName    *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int32  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryJobsRequest) GoString() string {
	return s.String()
}

func (s *QueryJobsRequest) SetJobName(v string) *QueryJobsRequest {
	s.JobName = &v
	return s
}

func (s *QueryJobsRequest) SetMaxResults(v int32) *QueryJobsRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryJobsRequest) SetNextToken(v int32) *QueryJobsRequest {
	s.NextToken = &v
	return s
}

type QueryJobsResponseBody struct {
	HasMore   *bool                        `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List      []*QueryJobsResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextToken *int64                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryJobsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryJobsResponseBody) SetHasMore(v bool) *QueryJobsResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryJobsResponseBody) SetList(v []*QueryJobsResponseBodyList) *QueryJobsResponseBody {
	s.List = v
	return s
}

func (s *QueryJobsResponseBody) SetNextToken(v int64) *QueryJobsResponseBody {
	s.NextToken = &v
	return s
}

type QueryJobsResponseBodyList struct {
	JobDescription *string `json:"jobDescription,omitempty" xml:"jobDescription,omitempty"`
	JobId          *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	JobName        *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
}

func (s QueryJobsResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryJobsResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryJobsResponseBodyList) SetJobDescription(v string) *QueryJobsResponseBodyList {
	s.JobDescription = &v
	return s
}

func (s *QueryJobsResponseBodyList) SetJobId(v string) *QueryJobsResponseBodyList {
	s.JobId = &v
	return s
}

func (s *QueryJobsResponseBodyList) SetJobName(v string) *QueryJobsResponseBodyList {
	s.JobName = &v
	return s
}

type QueryJobsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryJobsResponse) GoString() string {
	return s.String()
}

func (s *QueryJobsResponse) SetHeaders(v map[string]*string) *QueryJobsResponse {
	s.Headers = v
	return s
}

func (s *QueryJobsResponse) SetStatusCode(v int32) *QueryJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryJobsResponse) SetBody(v *QueryJobsResponseBody) *QueryJobsResponse {
	s.Body = v
	return s
}

type QueryPositionsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryPositionsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryPositionsHeaders) GoString() string {
	return s.String()
}

func (s *QueryPositionsHeaders) SetCommonHeaders(v map[string]*string) *QueryPositionsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryPositionsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryPositionsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryPositionsRequest struct {
	DeptId        *int64    `json:"deptId,omitempty" xml:"deptId,omitempty"`
	InCategoryIds []*string `json:"inCategoryIds,omitempty" xml:"inCategoryIds,omitempty" type:"Repeated"`
	InPositionIds []*string `json:"inPositionIds,omitempty" xml:"inPositionIds,omitempty" type:"Repeated"`
	PositionName  *string   `json:"positionName,omitempty" xml:"positionName,omitempty"`
	MaxResults    *int32    `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken     *int32    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryPositionsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPositionsRequest) GoString() string {
	return s.String()
}

func (s *QueryPositionsRequest) SetDeptId(v int64) *QueryPositionsRequest {
	s.DeptId = &v
	return s
}

func (s *QueryPositionsRequest) SetInCategoryIds(v []*string) *QueryPositionsRequest {
	s.InCategoryIds = v
	return s
}

func (s *QueryPositionsRequest) SetInPositionIds(v []*string) *QueryPositionsRequest {
	s.InPositionIds = v
	return s
}

func (s *QueryPositionsRequest) SetPositionName(v string) *QueryPositionsRequest {
	s.PositionName = &v
	return s
}

func (s *QueryPositionsRequest) SetMaxResults(v int32) *QueryPositionsRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryPositionsRequest) SetNextToken(v int32) *QueryPositionsRequest {
	s.NextToken = &v
	return s
}

type QueryPositionsResponseBody struct {
	HasMore   *bool                             `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List      []*QueryPositionsResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextToken *int64                            `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryPositionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPositionsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPositionsResponseBody) SetHasMore(v bool) *QueryPositionsResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryPositionsResponseBody) SetList(v []*QueryPositionsResponseBodyList) *QueryPositionsResponseBody {
	s.List = v
	return s
}

func (s *QueryPositionsResponseBody) SetNextToken(v int64) *QueryPositionsResponseBody {
	s.NextToken = &v
	return s
}

type QueryPositionsResponseBodyList struct {
	JobId              *string   `json:"jobId,omitempty" xml:"jobId,omitempty"`
	PositionCategoryId *string   `json:"positionCategoryId,omitempty" xml:"positionCategoryId,omitempty"`
	PositionDes        *string   `json:"positionDes,omitempty" xml:"positionDes,omitempty"`
	PositionId         *string   `json:"positionId,omitempty" xml:"positionId,omitempty"`
	PositionName       *string   `json:"positionName,omitempty" xml:"positionName,omitempty"`
	RankIdList         []*string `json:"rankIdList,omitempty" xml:"rankIdList,omitempty" type:"Repeated"`
	Status             *int32    `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryPositionsResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryPositionsResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryPositionsResponseBodyList) SetJobId(v string) *QueryPositionsResponseBodyList {
	s.JobId = &v
	return s
}

func (s *QueryPositionsResponseBodyList) SetPositionCategoryId(v string) *QueryPositionsResponseBodyList {
	s.PositionCategoryId = &v
	return s
}

func (s *QueryPositionsResponseBodyList) SetPositionDes(v string) *QueryPositionsResponseBodyList {
	s.PositionDes = &v
	return s
}

func (s *QueryPositionsResponseBodyList) SetPositionId(v string) *QueryPositionsResponseBodyList {
	s.PositionId = &v
	return s
}

func (s *QueryPositionsResponseBodyList) SetPositionName(v string) *QueryPositionsResponseBodyList {
	s.PositionName = &v
	return s
}

func (s *QueryPositionsResponseBodyList) SetRankIdList(v []*string) *QueryPositionsResponseBodyList {
	s.RankIdList = v
	return s
}

func (s *QueryPositionsResponseBodyList) SetStatus(v int32) *QueryPositionsResponseBodyList {
	s.Status = &v
	return s
}

type QueryPositionsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryPositionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPositionsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPositionsResponse) GoString() string {
	return s.String()
}

func (s *QueryPositionsResponse) SetHeaders(v map[string]*string) *QueryPositionsResponse {
	s.Headers = v
	return s
}

func (s *QueryPositionsResponse) SetStatusCode(v int32) *QueryPositionsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPositionsResponse) SetBody(v *QueryPositionsResponseBody) *QueryPositionsResponse {
	s.Body = v
	return s
}

type RosterMetaAvailableFieldListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RosterMetaAvailableFieldListHeaders) String() string {
	return tea.Prettify(s)
}

func (s RosterMetaAvailableFieldListHeaders) GoString() string {
	return s.String()
}

func (s *RosterMetaAvailableFieldListHeaders) SetCommonHeaders(v map[string]*string) *RosterMetaAvailableFieldListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RosterMetaAvailableFieldListHeaders) SetXAcsDingtalkAccessToken(v string) *RosterMetaAvailableFieldListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RosterMetaAvailableFieldListRequest struct {
	AppAgentId *int64 `json:"appAgentId,omitempty" xml:"appAgentId,omitempty"`
}

func (s RosterMetaAvailableFieldListRequest) String() string {
	return tea.Prettify(s)
}

func (s RosterMetaAvailableFieldListRequest) GoString() string {
	return s.String()
}

func (s *RosterMetaAvailableFieldListRequest) SetAppAgentId(v int64) *RosterMetaAvailableFieldListRequest {
	s.AppAgentId = &v
	return s
}

type RosterMetaAvailableFieldListResponseBody struct {
	Result []*RosterMetaAvailableFieldListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s RosterMetaAvailableFieldListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RosterMetaAvailableFieldListResponseBody) GoString() string {
	return s.String()
}

func (s *RosterMetaAvailableFieldListResponseBody) SetResult(v []*RosterMetaAvailableFieldListResponseBodyResult) *RosterMetaAvailableFieldListResponseBody {
	s.Result = v
	return s
}

type RosterMetaAvailableFieldListResponseBodyResult struct {
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
}

func (s RosterMetaAvailableFieldListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s RosterMetaAvailableFieldListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RosterMetaAvailableFieldListResponseBodyResult) SetFieldCode(v string) *RosterMetaAvailableFieldListResponseBodyResult {
	s.FieldCode = &v
	return s
}

func (s *RosterMetaAvailableFieldListResponseBodyResult) SetFieldName(v string) *RosterMetaAvailableFieldListResponseBodyResult {
	s.FieldName = &v
	return s
}

func (s *RosterMetaAvailableFieldListResponseBodyResult) SetFieldType(v string) *RosterMetaAvailableFieldListResponseBodyResult {
	s.FieldType = &v
	return s
}

type RosterMetaAvailableFieldListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RosterMetaAvailableFieldListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RosterMetaAvailableFieldListResponse) String() string {
	return tea.Prettify(s)
}

func (s RosterMetaAvailableFieldListResponse) GoString() string {
	return s.String()
}

func (s *RosterMetaAvailableFieldListResponse) SetHeaders(v map[string]*string) *RosterMetaAvailableFieldListResponse {
	s.Headers = v
	return s
}

func (s *RosterMetaAvailableFieldListResponse) SetStatusCode(v int32) *RosterMetaAvailableFieldListResponse {
	s.StatusCode = &v
	return s
}

func (s *RosterMetaAvailableFieldListResponse) SetBody(v *RosterMetaAvailableFieldListResponseBody) *RosterMetaAvailableFieldListResponse {
	s.Body = v
	return s
}

type RosterMetaFieldOptionsUpdateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RosterMetaFieldOptionsUpdateHeaders) String() string {
	return tea.Prettify(s)
}

func (s RosterMetaFieldOptionsUpdateHeaders) GoString() string {
	return s.String()
}

func (s *RosterMetaFieldOptionsUpdateHeaders) SetCommonHeaders(v map[string]*string) *RosterMetaFieldOptionsUpdateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RosterMetaFieldOptionsUpdateHeaders) SetXAcsDingtalkAccessToken(v string) *RosterMetaFieldOptionsUpdateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RosterMetaFieldOptionsUpdateRequest struct {
	AppAgentId *int64    `json:"appAgentId,omitempty" xml:"appAgentId,omitempty"`
	FieldCode  *string   `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	GroupId    *string   `json:"groupId,omitempty" xml:"groupId,omitempty"`
	Labels     []*string `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	ModifyType *string   `json:"modifyType,omitempty" xml:"modifyType,omitempty"`
}

func (s RosterMetaFieldOptionsUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s RosterMetaFieldOptionsUpdateRequest) GoString() string {
	return s.String()
}

func (s *RosterMetaFieldOptionsUpdateRequest) SetAppAgentId(v int64) *RosterMetaFieldOptionsUpdateRequest {
	s.AppAgentId = &v
	return s
}

func (s *RosterMetaFieldOptionsUpdateRequest) SetFieldCode(v string) *RosterMetaFieldOptionsUpdateRequest {
	s.FieldCode = &v
	return s
}

func (s *RosterMetaFieldOptionsUpdateRequest) SetGroupId(v string) *RosterMetaFieldOptionsUpdateRequest {
	s.GroupId = &v
	return s
}

func (s *RosterMetaFieldOptionsUpdateRequest) SetLabels(v []*string) *RosterMetaFieldOptionsUpdateRequest {
	s.Labels = v
	return s
}

func (s *RosterMetaFieldOptionsUpdateRequest) SetModifyType(v string) *RosterMetaFieldOptionsUpdateRequest {
	s.ModifyType = &v
	return s
}

type RosterMetaFieldOptionsUpdateResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RosterMetaFieldOptionsUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RosterMetaFieldOptionsUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *RosterMetaFieldOptionsUpdateResponseBody) SetResult(v bool) *RosterMetaFieldOptionsUpdateResponseBody {
	s.Result = &v
	return s
}

type RosterMetaFieldOptionsUpdateResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RosterMetaFieldOptionsUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RosterMetaFieldOptionsUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s RosterMetaFieldOptionsUpdateResponse) GoString() string {
	return s.String()
}

func (s *RosterMetaFieldOptionsUpdateResponse) SetHeaders(v map[string]*string) *RosterMetaFieldOptionsUpdateResponse {
	s.Headers = v
	return s
}

func (s *RosterMetaFieldOptionsUpdateResponse) SetStatusCode(v int32) *RosterMetaFieldOptionsUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *RosterMetaFieldOptionsUpdateResponse) SetBody(v *RosterMetaFieldOptionsUpdateResponseBody) *RosterMetaFieldOptionsUpdateResponse {
	s.Body = v
	return s
}

type SolutionTaskInitHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SolutionTaskInitHeaders) String() string {
	return tea.Prettify(s)
}

func (s SolutionTaskInitHeaders) GoString() string {
	return s.String()
}

func (s *SolutionTaskInitHeaders) SetCommonHeaders(v map[string]*string) *SolutionTaskInitHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SolutionTaskInitHeaders) SetXAcsDingtalkAccessToken(v string) *SolutionTaskInitHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SolutionTaskInitRequest struct {
	Category     *string `json:"category,omitempty" xml:"category,omitempty"`
	ClaimTime    *int64  `json:"claimTime,omitempty" xml:"claimTime,omitempty"`
	Description  *string `json:"description,omitempty" xml:"description,omitempty"`
	FinishTime   *int64  `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	OuterId      *string `json:"outerId,omitempty" xml:"outerId,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
	Title        *string `json:"title,omitempty" xml:"title,omitempty"`
	UserId       *string `json:"userId,omitempty" xml:"userId,omitempty"`
	SolutionType *string `json:"solutionType,omitempty" xml:"solutionType,omitempty"`
}

func (s SolutionTaskInitRequest) String() string {
	return tea.Prettify(s)
}

func (s SolutionTaskInitRequest) GoString() string {
	return s.String()
}

func (s *SolutionTaskInitRequest) SetCategory(v string) *SolutionTaskInitRequest {
	s.Category = &v
	return s
}

func (s *SolutionTaskInitRequest) SetClaimTime(v int64) *SolutionTaskInitRequest {
	s.ClaimTime = &v
	return s
}

func (s *SolutionTaskInitRequest) SetDescription(v string) *SolutionTaskInitRequest {
	s.Description = &v
	return s
}

func (s *SolutionTaskInitRequest) SetFinishTime(v int64) *SolutionTaskInitRequest {
	s.FinishTime = &v
	return s
}

func (s *SolutionTaskInitRequest) SetOuterId(v string) *SolutionTaskInitRequest {
	s.OuterId = &v
	return s
}

func (s *SolutionTaskInitRequest) SetStatus(v string) *SolutionTaskInitRequest {
	s.Status = &v
	return s
}

func (s *SolutionTaskInitRequest) SetTitle(v string) *SolutionTaskInitRequest {
	s.Title = &v
	return s
}

func (s *SolutionTaskInitRequest) SetUserId(v string) *SolutionTaskInitRequest {
	s.UserId = &v
	return s
}

func (s *SolutionTaskInitRequest) SetSolutionType(v string) *SolutionTaskInitRequest {
	s.SolutionType = &v
	return s
}

type SolutionTaskInitResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SolutionTaskInitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SolutionTaskInitResponseBody) GoString() string {
	return s.String()
}

func (s *SolutionTaskInitResponseBody) SetResult(v bool) *SolutionTaskInitResponseBody {
	s.Result = &v
	return s
}

type SolutionTaskInitResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SolutionTaskInitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SolutionTaskInitResponse) String() string {
	return tea.Prettify(s)
}

func (s SolutionTaskInitResponse) GoString() string {
	return s.String()
}

func (s *SolutionTaskInitResponse) SetHeaders(v map[string]*string) *SolutionTaskInitResponse {
	s.Headers = v
	return s
}

func (s *SolutionTaskInitResponse) SetStatusCode(v int32) *SolutionTaskInitResponse {
	s.StatusCode = &v
	return s
}

func (s *SolutionTaskInitResponse) SetBody(v *SolutionTaskInitResponseBody) *SolutionTaskInitResponse {
	s.Body = v
	return s
}

type SolutionTaskSaveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SolutionTaskSaveHeaders) String() string {
	return tea.Prettify(s)
}

func (s SolutionTaskSaveHeaders) GoString() string {
	return s.String()
}

func (s *SolutionTaskSaveHeaders) SetCommonHeaders(v map[string]*string) *SolutionTaskSaveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SolutionTaskSaveHeaders) SetXAcsDingtalkAccessToken(v string) *SolutionTaskSaveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SolutionTaskSaveRequest struct {
	ClaimTime          *int64  `json:"claimTime,omitempty" xml:"claimTime,omitempty"`
	Description        *string `json:"description,omitempty" xml:"description,omitempty"`
	FinishTime         *int64  `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	OuterId            *string `json:"outerId,omitempty" xml:"outerId,omitempty"`
	SolutionInstanceId *string `json:"solutionInstanceId,omitempty" xml:"solutionInstanceId,omitempty"`
	StartTime          *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status             *string `json:"status,omitempty" xml:"status,omitempty"`
	TaskType           *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	TemplateOuterId    *string `json:"templateOuterId,omitempty" xml:"templateOuterId,omitempty"`
	Title              *string `json:"title,omitempty" xml:"title,omitempty"`
	UserId             *string `json:"userId,omitempty" xml:"userId,omitempty"`
	SolutionType       *string `json:"solutionType,omitempty" xml:"solutionType,omitempty"`
}

func (s SolutionTaskSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s SolutionTaskSaveRequest) GoString() string {
	return s.String()
}

func (s *SolutionTaskSaveRequest) SetClaimTime(v int64) *SolutionTaskSaveRequest {
	s.ClaimTime = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetDescription(v string) *SolutionTaskSaveRequest {
	s.Description = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetFinishTime(v int64) *SolutionTaskSaveRequest {
	s.FinishTime = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetOuterId(v string) *SolutionTaskSaveRequest {
	s.OuterId = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetSolutionInstanceId(v string) *SolutionTaskSaveRequest {
	s.SolutionInstanceId = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetStartTime(v int64) *SolutionTaskSaveRequest {
	s.StartTime = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetStatus(v string) *SolutionTaskSaveRequest {
	s.Status = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetTaskType(v string) *SolutionTaskSaveRequest {
	s.TaskType = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetTemplateOuterId(v string) *SolutionTaskSaveRequest {
	s.TemplateOuterId = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetTitle(v string) *SolutionTaskSaveRequest {
	s.Title = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetUserId(v string) *SolutionTaskSaveRequest {
	s.UserId = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetSolutionType(v string) *SolutionTaskSaveRequest {
	s.SolutionType = &v
	return s
}

type SolutionTaskSaveResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SolutionTaskSaveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SolutionTaskSaveResponseBody) GoString() string {
	return s.String()
}

func (s *SolutionTaskSaveResponseBody) SetResult(v bool) *SolutionTaskSaveResponseBody {
	s.Result = &v
	return s
}

type SolutionTaskSaveResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SolutionTaskSaveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SolutionTaskSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s SolutionTaskSaveResponse) GoString() string {
	return s.String()
}

func (s *SolutionTaskSaveResponse) SetHeaders(v map[string]*string) *SolutionTaskSaveResponse {
	s.Headers = v
	return s
}

func (s *SolutionTaskSaveResponse) SetStatusCode(v int32) *SolutionTaskSaveResponse {
	s.StatusCode = &v
	return s
}

func (s *SolutionTaskSaveResponse) SetBody(v *SolutionTaskSaveResponseBody) *SolutionTaskSaveResponse {
	s.Body = v
	return s
}

type SyncTaskTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SyncTaskTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncTaskTemplateHeaders) GoString() string {
	return s.String()
}

func (s *SyncTaskTemplateHeaders) SetCommonHeaders(v map[string]*string) *SyncTaskTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncTaskTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *SyncTaskTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SyncTaskTemplateRequest struct {
	Delete       *bool                               `json:"delete,omitempty" xml:"delete,omitempty"`
	Des          *string                             `json:"des,omitempty" xml:"des,omitempty"`
	Ext          *string                             `json:"ext,omitempty" xml:"ext,omitempty"`
	Name         *string                             `json:"name,omitempty" xml:"name,omitempty"`
	OptUserId    *string                             `json:"optUserId,omitempty" xml:"optUserId,omitempty"`
	OuterId      *string                             `json:"outerId,omitempty" xml:"outerId,omitempty"`
	TaskScopeVO  *SyncTaskTemplateRequestTaskScopeVO `json:"taskScopeVO,omitempty" xml:"taskScopeVO,omitempty" type:"Struct"`
	TaskType     *string                             `json:"taskType,omitempty" xml:"taskType,omitempty"`
	SolutionType *string                             `json:"solutionType,omitempty" xml:"solutionType,omitempty"`
}

func (s SyncTaskTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncTaskTemplateRequest) GoString() string {
	return s.String()
}

func (s *SyncTaskTemplateRequest) SetDelete(v bool) *SyncTaskTemplateRequest {
	s.Delete = &v
	return s
}

func (s *SyncTaskTemplateRequest) SetDes(v string) *SyncTaskTemplateRequest {
	s.Des = &v
	return s
}

func (s *SyncTaskTemplateRequest) SetExt(v string) *SyncTaskTemplateRequest {
	s.Ext = &v
	return s
}

func (s *SyncTaskTemplateRequest) SetName(v string) *SyncTaskTemplateRequest {
	s.Name = &v
	return s
}

func (s *SyncTaskTemplateRequest) SetOptUserId(v string) *SyncTaskTemplateRequest {
	s.OptUserId = &v
	return s
}

func (s *SyncTaskTemplateRequest) SetOuterId(v string) *SyncTaskTemplateRequest {
	s.OuterId = &v
	return s
}

func (s *SyncTaskTemplateRequest) SetTaskScopeVO(v *SyncTaskTemplateRequestTaskScopeVO) *SyncTaskTemplateRequest {
	s.TaskScopeVO = v
	return s
}

func (s *SyncTaskTemplateRequest) SetTaskType(v string) *SyncTaskTemplateRequest {
	s.TaskType = &v
	return s
}

func (s *SyncTaskTemplateRequest) SetSolutionType(v string) *SyncTaskTemplateRequest {
	s.SolutionType = &v
	return s
}

type SyncTaskTemplateRequestTaskScopeVO struct {
	DeptIds     []*int64  `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	PositionIds []*string `json:"positionIds,omitempty" xml:"positionIds,omitempty" type:"Repeated"`
	RoleIds     []*string `json:"roleIds,omitempty" xml:"roleIds,omitempty" type:"Repeated"`
	UserIds     []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s SyncTaskTemplateRequestTaskScopeVO) String() string {
	return tea.Prettify(s)
}

func (s SyncTaskTemplateRequestTaskScopeVO) GoString() string {
	return s.String()
}

func (s *SyncTaskTemplateRequestTaskScopeVO) SetDeptIds(v []*int64) *SyncTaskTemplateRequestTaskScopeVO {
	s.DeptIds = v
	return s
}

func (s *SyncTaskTemplateRequestTaskScopeVO) SetPositionIds(v []*string) *SyncTaskTemplateRequestTaskScopeVO {
	s.PositionIds = v
	return s
}

func (s *SyncTaskTemplateRequestTaskScopeVO) SetRoleIds(v []*string) *SyncTaskTemplateRequestTaskScopeVO {
	s.RoleIds = v
	return s
}

func (s *SyncTaskTemplateRequestTaskScopeVO) SetUserIds(v []*string) *SyncTaskTemplateRequestTaskScopeVO {
	s.UserIds = v
	return s
}

type SyncTaskTemplateResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SyncTaskTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncTaskTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *SyncTaskTemplateResponseBody) SetResult(v bool) *SyncTaskTemplateResponseBody {
	s.Result = &v
	return s
}

type SyncTaskTemplateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SyncTaskTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncTaskTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncTaskTemplateResponse) GoString() string {
	return s.String()
}

func (s *SyncTaskTemplateResponse) SetHeaders(v map[string]*string) *SyncTaskTemplateResponse {
	s.Headers = v
	return s
}

func (s *SyncTaskTemplateResponse) SetStatusCode(v int32) *SyncTaskTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncTaskTemplateResponse) SetBody(v *SyncTaskTemplateResponseBody) *SyncTaskTemplateResponse {
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

func (client *Client) AddHrmPreentryWithOptions(request *AddHrmPreentryRequest, headers *AddHrmPreentryHeaders, runtime *util.RuntimeOptions) (_result *AddHrmPreentryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		body["agentId"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.Groups)) {
		body["groups"] = request.Groups
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		body["mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NeedSendPreEntryMsg)) {
		body["needSendPreEntryMsg"] = request.NeedSendPreEntryMsg
	}

	if !tea.BoolValue(util.IsUnset(request.PreEntryTime)) {
		body["preEntryTime"] = request.PreEntryTime
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
		Action:      tea.String("AddHrmPreentry"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/preentries"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddHrmPreentryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddHrmPreentry(request *AddHrmPreentryRequest) (_result *AddHrmPreentryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddHrmPreentryHeaders{}
	_result = &AddHrmPreentryResponse{}
	_body, _err := client.AddHrmPreentryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ECertQueryWithOptions(request *ECertQueryRequest, headers *ECertQueryHeaders, runtime *util.RuntimeOptions) (_result *ECertQueryResponse, _err error) {
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
		Action:      tea.String("ECertQuery"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/eCerts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ECertQueryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ECertQuery(request *ECertQueryRequest) (_result *ECertQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ECertQueryHeaders{}
	_result = &ECertQueryResponse{}
	_body, _err := client.ECertQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HrmProcessRegularWithOptions(request *HrmProcessRegularRequest, headers *HrmProcessRegularHeaders, runtime *util.RuntimeOptions) (_result *HrmProcessRegularResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperationId)) {
		body["operationId"] = request.OperationId
	}

	if !tea.BoolValue(util.IsUnset(request.RegularDate)) {
		body["regularDate"] = request.RegularDate
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
		Action:      tea.String("HrmProcessRegular"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/processes/regulars/become"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrmProcessRegularResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HrmProcessRegular(request *HrmProcessRegularRequest) (_result *HrmProcessRegularResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrmProcessRegularHeaders{}
	_result = &HrmProcessRegularResponse{}
	_body, _err := client.HrmProcessRegularWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HrmProcessTransferWithOptions(request *HrmProcessTransferRequest, headers *HrmProcessTransferHeaders, runtime *util.RuntimeOptions) (_result *HrmProcessTransferResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptIdsAfterTransfer)) {
		body["deptIdsAfterTransfer"] = request.DeptIdsAfterTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.JobIdAfterTransfer)) {
		body["jobIdAfterTransfer"] = request.JobIdAfterTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.MainDeptIdAfterTransfer)) {
		body["mainDeptIdAfterTransfer"] = request.MainDeptIdAfterTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.OperateUserId)) {
		body["operateUserId"] = request.OperateUserId
	}

	if !tea.BoolValue(util.IsUnset(request.PositionIdAfterTransfer)) {
		body["positionIdAfterTransfer"] = request.PositionIdAfterTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.PositionLevelAfterTransfer)) {
		body["positionLevelAfterTransfer"] = request.PositionLevelAfterTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.PositionNameAfterTransfer)) {
		body["positionNameAfterTransfer"] = request.PositionNameAfterTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.RankIdAfterTransfer)) {
		body["rankIdAfterTransfer"] = request.RankIdAfterTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
		Action:      tea.String("HrmProcessTransfer"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/processes/transfer"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrmProcessTransferResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HrmProcessTransfer(request *HrmProcessTransferRequest) (_result *HrmProcessTransferResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrmProcessTransferHeaders{}
	_result = &HrmProcessTransferResponse{}
	_body, _err := client.HrmProcessTransferWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HrmProcessUpdateTerminationInfoWithOptions(request *HrmProcessUpdateTerminationInfoRequest, headers *HrmProcessUpdateTerminationInfoHeaders, runtime *util.RuntimeOptions) (_result *HrmProcessUpdateTerminationInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DismissionMemo)) {
		body["dismissionMemo"] = request.DismissionMemo
	}

	if !tea.BoolValue(util.IsUnset(request.LastWorkDate)) {
		body["lastWorkDate"] = request.LastWorkDate
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
		Action:      tea.String("HrmProcessUpdateTerminationInfo"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/processes/employees/terminations"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrmProcessUpdateTerminationInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HrmProcessUpdateTerminationInfo(request *HrmProcessUpdateTerminationInfoRequest) (_result *HrmProcessUpdateTerminationInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrmProcessUpdateTerminationInfoHeaders{}
	_result = &HrmProcessUpdateTerminationInfoResponse{}
	_body, _err := client.HrmProcessUpdateTerminationInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MasterDataQueryWithOptions(request *MasterDataQueryRequest, headers *MasterDataQueryHeaders, runtime *util.RuntimeOptions) (_result *MasterDataQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizUK)) {
		body["bizUK"] = request.BizUK
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OptUserId)) {
		body["optUserId"] = request.OptUserId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryParams)) {
		body["queryParams"] = request.QueryParams
	}

	if !tea.BoolValue(util.IsUnset(request.RelationIds)) {
		body["relationIds"] = request.RelationIds
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeCode)) {
		body["scopeCode"] = request.ScopeCode
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["tenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.ViewEntityCode)) {
		body["viewEntityCode"] = request.ViewEntityCode
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
		Action:      tea.String("MasterDataQuery"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/masters/datas/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &MasterDataQueryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MasterDataQuery(request *MasterDataQueryRequest) (_result *MasterDataQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MasterDataQueryHeaders{}
	_result = &MasterDataQueryResponse{}
	_body, _err := client.MasterDataQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MasterDataSaveWithOptions(request *MasterDataSaveRequest, headers *MasterDataSaveHeaders, runtime *util.RuntimeOptions) (_result *MasterDataSaveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
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
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("MasterDataSave"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/masters/datas/save"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &MasterDataSaveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MasterDataSave(request *MasterDataSaveRequest) (_result *MasterDataSaveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MasterDataSaveHeaders{}
	_result = &MasterDataSaveResponse{}
	_body, _err := client.MasterDataSaveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MasterDataTenantQueyWithOptions(request *MasterDataTenantQueyRequest, headers *MasterDataTenantQueyHeaders, runtime *util.RuntimeOptions) (_result *MasterDataTenantQueyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityCode)) {
		query["entityCode"] = request.EntityCode
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeCode)) {
		query["scopeCode"] = request.ScopeCode
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
		Action:      tea.String("MasterDataTenantQuey"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/masters/tenants"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &MasterDataTenantQueyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MasterDataTenantQuey(request *MasterDataTenantQueyRequest) (_result *MasterDataTenantQueyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MasterDataTenantQueyHeaders{}
	_result = &MasterDataTenantQueyResponse{}
	_body, _err := client.MasterDataTenantQueyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCustomEntryProcessesWithOptions(request *QueryCustomEntryProcessesRequest, headers *QueryCustomEntryProcessesHeaders, runtime *util.RuntimeOptions) (_result *QueryCustomEntryProcessesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.OperateUserId)) {
		query["operateUserId"] = request.OperateUserId
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
		Action:      tea.String("QueryCustomEntryProcesses"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/customEntryProcesses"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCustomEntryProcessesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCustomEntryProcesses(request *QueryCustomEntryProcessesRequest) (_result *QueryCustomEntryProcessesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCustomEntryProcessesHeaders{}
	_result = &QueryCustomEntryProcessesResponse{}
	_body, _err := client.QueryCustomEntryProcessesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDismissionStaffIdListWithOptions(request *QueryDismissionStaffIdListRequest, headers *QueryDismissionStaffIdListHeaders, runtime *util.RuntimeOptions) (_result *QueryDismissionStaffIdListResponse, _err error) {
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
		Action:      tea.String("QueryDismissionStaffIdList"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/employees/dismissions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDismissionStaffIdListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDismissionStaffIdList(request *QueryDismissionStaffIdListRequest) (_result *QueryDismissionStaffIdListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDismissionStaffIdListHeaders{}
	_result = &QueryDismissionStaffIdListResponse{}
	_body, _err := client.QueryDismissionStaffIdListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryHrmEmployeeDismissionInfoWithOptions(tmpReq *QueryHrmEmployeeDismissionInfoRequest, headers *QueryHrmEmployeeDismissionInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryHrmEmployeeDismissionInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryHrmEmployeeDismissionInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UserIdList)) {
		request.UserIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIdList, tea.String("userIdList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserIdListShrink)) {
		query["userIdList"] = request.UserIdListShrink
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
		Action:      tea.String("QueryHrmEmployeeDismissionInfo"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/employees/dimissionInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryHrmEmployeeDismissionInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryHrmEmployeeDismissionInfo(request *QueryHrmEmployeeDismissionInfoRequest) (_result *QueryHrmEmployeeDismissionInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryHrmEmployeeDismissionInfoHeaders{}
	_result = &QueryHrmEmployeeDismissionInfoResponse{}
	_body, _err := client.QueryHrmEmployeeDismissionInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryJobRanksWithOptions(request *QueryJobRanksRequest, headers *QueryJobRanksHeaders, runtime *util.RuntimeOptions) (_result *QueryJobRanksResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.RankCategoryId)) {
		query["rankCategoryId"] = request.RankCategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.RankCode)) {
		query["rankCode"] = request.RankCode
	}

	if !tea.BoolValue(util.IsUnset(request.RankName)) {
		query["rankName"] = request.RankName
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
		Action:      tea.String("QueryJobRanks"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/jobRanks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryJobRanksResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryJobRanks(request *QueryJobRanksRequest) (_result *QueryJobRanksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryJobRanksHeaders{}
	_result = &QueryJobRanksResponse{}
	_body, _err := client.QueryJobRanksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryJobsWithOptions(request *QueryJobsRequest, headers *QueryJobsHeaders, runtime *util.RuntimeOptions) (_result *QueryJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobName)) {
		query["jobName"] = request.JobName
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
		Action:      tea.String("QueryJobs"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/jobs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryJobsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryJobs(request *QueryJobsRequest) (_result *QueryJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryJobsHeaders{}
	_result = &QueryJobsResponse{}
	_body, _err := client.QueryJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPositionsWithOptions(request *QueryPositionsRequest, headers *QueryPositionsHeaders, runtime *util.RuntimeOptions) (_result *QueryPositionsResponse, _err error) {
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

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.InCategoryIds)) {
		body["inCategoryIds"] = request.InCategoryIds
	}

	if !tea.BoolValue(util.IsUnset(request.InPositionIds)) {
		body["inPositionIds"] = request.InPositionIds
	}

	if !tea.BoolValue(util.IsUnset(request.PositionName)) {
		body["positionName"] = request.PositionName
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
		Action:      tea.String("QueryPositions"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/positions/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPositionsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPositions(request *QueryPositionsRequest) (_result *QueryPositionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryPositionsHeaders{}
	_result = &QueryPositionsResponse{}
	_body, _err := client.QueryPositionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RosterMetaAvailableFieldListWithOptions(request *RosterMetaAvailableFieldListRequest, headers *RosterMetaAvailableFieldListHeaders, runtime *util.RuntimeOptions) (_result *RosterMetaAvailableFieldListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppAgentId)) {
		query["appAgentId"] = request.AppAgentId
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
		Action:      tea.String("RosterMetaAvailableFieldList"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/rosters/meta/authorities/fields"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RosterMetaAvailableFieldListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RosterMetaAvailableFieldList(request *RosterMetaAvailableFieldListRequest) (_result *RosterMetaAvailableFieldListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RosterMetaAvailableFieldListHeaders{}
	_result = &RosterMetaAvailableFieldListResponse{}
	_body, _err := client.RosterMetaAvailableFieldListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RosterMetaFieldOptionsUpdateWithOptions(request *RosterMetaFieldOptionsUpdateRequest, headers *RosterMetaFieldOptionsUpdateHeaders, runtime *util.RuntimeOptions) (_result *RosterMetaFieldOptionsUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppAgentId)) {
		query["appAgentId"] = request.AppAgentId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FieldCode)) {
		body["fieldCode"] = request.FieldCode
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["groupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyType)) {
		body["modifyType"] = request.ModifyType
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
		Action:      tea.String("RosterMetaFieldOptionsUpdate"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/rosters/meta/fields/options"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RosterMetaFieldOptionsUpdateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RosterMetaFieldOptionsUpdate(request *RosterMetaFieldOptionsUpdateRequest) (_result *RosterMetaFieldOptionsUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RosterMetaFieldOptionsUpdateHeaders{}
	_result = &RosterMetaFieldOptionsUpdateResponse{}
	_body, _err := client.RosterMetaFieldOptionsUpdateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SolutionTaskInitWithOptions(request *SolutionTaskInitRequest, headers *SolutionTaskInitHeaders, runtime *util.RuntimeOptions) (_result *SolutionTaskInitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SolutionType)) {
		query["solutionType"] = request.SolutionType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.ClaimTime)) {
		body["claimTime"] = request.ClaimTime
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FinishTime)) {
		body["finishTime"] = request.FinishTime
	}

	if !tea.BoolValue(util.IsUnset(request.OuterId)) {
		body["outerId"] = request.OuterId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
		Action:      tea.String("SolutionTaskInit"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/solutions/tasks/init"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SolutionTaskInitResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SolutionTaskInit(request *SolutionTaskInitRequest) (_result *SolutionTaskInitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SolutionTaskInitHeaders{}
	_result = &SolutionTaskInitResponse{}
	_body, _err := client.SolutionTaskInitWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SolutionTaskSaveWithOptions(request *SolutionTaskSaveRequest, headers *SolutionTaskSaveHeaders, runtime *util.RuntimeOptions) (_result *SolutionTaskSaveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SolutionType)) {
		query["solutionType"] = request.SolutionType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClaimTime)) {
		body["claimTime"] = request.ClaimTime
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FinishTime)) {
		body["finishTime"] = request.FinishTime
	}

	if !tea.BoolValue(util.IsUnset(request.OuterId)) {
		body["outerId"] = request.OuterId
	}

	if !tea.BoolValue(util.IsUnset(request.SolutionInstanceId)) {
		body["solutionInstanceId"] = request.SolutionInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		body["taskType"] = request.TaskType
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateOuterId)) {
		body["templateOuterId"] = request.TemplateOuterId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
		Action:      tea.String("SolutionTaskSave"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/solutions/tasks/save"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SolutionTaskSaveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SolutionTaskSave(request *SolutionTaskSaveRequest) (_result *SolutionTaskSaveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SolutionTaskSaveHeaders{}
	_result = &SolutionTaskSaveResponse{}
	_body, _err := client.SolutionTaskSaveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncTaskTemplateWithOptions(request *SyncTaskTemplateRequest, headers *SyncTaskTemplateHeaders, runtime *util.RuntimeOptions) (_result *SyncTaskTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SolutionType)) {
		query["solutionType"] = request.SolutionType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Delete)) {
		body["delete"] = request.Delete
	}

	if !tea.BoolValue(util.IsUnset(request.Des)) {
		body["des"] = request.Des
	}

	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OptUserId)) {
		body["optUserId"] = request.OptUserId
	}

	if !tea.BoolValue(util.IsUnset(request.OuterId)) {
		body["outerId"] = request.OuterId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskScopeVO)) {
		body["taskScopeVO"] = request.TaskScopeVO
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		body["taskType"] = request.TaskType
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
		Action:      tea.String("SyncTaskTemplate"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/solutions/tasks/templates/sync"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncTaskTemplateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncTaskTemplate(request *SyncTaskTemplateRequest) (_result *SyncTaskTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncTaskTemplateHeaders{}
	_result = &SyncTaskTemplateResponse{}
	_body, _err := client.SyncTaskTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
