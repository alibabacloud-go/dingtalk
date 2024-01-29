// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package contract_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ContractBenefitConsumeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ContractBenefitConsumeHeaders) String() string {
	return tea.Prettify(s)
}

func (s ContractBenefitConsumeHeaders) GoString() string {
	return s.String()
}

func (s *ContractBenefitConsumeHeaders) SetCommonHeaders(v map[string]*string) *ContractBenefitConsumeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ContractBenefitConsumeHeaders) SetXAcsDingtalkAccessToken(v string) *ContractBenefitConsumeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ContractBenefitConsumeRequest struct {
	BenefitPoint *string            `json:"benefitPoint,omitempty" xml:"benefitPoint,omitempty"`
	BizRequestId *string            `json:"bizRequestId,omitempty" xml:"bizRequestId,omitempty"`
	ConsumeQuota *int64             `json:"consumeQuota,omitempty" xml:"consumeQuota,omitempty"`
	CorpId       *string            `json:"corpId,omitempty" xml:"corpId,omitempty"`
	ExtParams    map[string]*string `json:"extParams,omitempty" xml:"extParams,omitempty"`
	IsvCorpId    *string            `json:"isvCorpId,omitempty" xml:"isvCorpId,omitempty"`
	OptUnionId   *string            `json:"optUnionId,omitempty" xml:"optUnionId,omitempty"`
}

func (s ContractBenefitConsumeRequest) String() string {
	return tea.Prettify(s)
}

func (s ContractBenefitConsumeRequest) GoString() string {
	return s.String()
}

func (s *ContractBenefitConsumeRequest) SetBenefitPoint(v string) *ContractBenefitConsumeRequest {
	s.BenefitPoint = &v
	return s
}

func (s *ContractBenefitConsumeRequest) SetBizRequestId(v string) *ContractBenefitConsumeRequest {
	s.BizRequestId = &v
	return s
}

func (s *ContractBenefitConsumeRequest) SetConsumeQuota(v int64) *ContractBenefitConsumeRequest {
	s.ConsumeQuota = &v
	return s
}

func (s *ContractBenefitConsumeRequest) SetCorpId(v string) *ContractBenefitConsumeRequest {
	s.CorpId = &v
	return s
}

func (s *ContractBenefitConsumeRequest) SetExtParams(v map[string]*string) *ContractBenefitConsumeRequest {
	s.ExtParams = v
	return s
}

func (s *ContractBenefitConsumeRequest) SetIsvCorpId(v string) *ContractBenefitConsumeRequest {
	s.IsvCorpId = &v
	return s
}

func (s *ContractBenefitConsumeRequest) SetOptUnionId(v string) *ContractBenefitConsumeRequest {
	s.OptUnionId = &v
	return s
}

type ContractBenefitConsumeResponseBody struct {
	Result  *ContractBenefitConsumeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ContractBenefitConsumeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ContractBenefitConsumeResponseBody) GoString() string {
	return s.String()
}

func (s *ContractBenefitConsumeResponseBody) SetResult(v *ContractBenefitConsumeResponseBodyResult) *ContractBenefitConsumeResponseBody {
	s.Result = v
	return s
}

func (s *ContractBenefitConsumeResponseBody) SetSuccess(v bool) *ContractBenefitConsumeResponseBody {
	s.Success = &v
	return s
}

type ContractBenefitConsumeResponseBodyResult struct {
	ConsumeResult *bool `json:"consumeResult,omitempty" xml:"consumeResult,omitempty"`
}

func (s ContractBenefitConsumeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ContractBenefitConsumeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ContractBenefitConsumeResponseBodyResult) SetConsumeResult(v bool) *ContractBenefitConsumeResponseBodyResult {
	s.ConsumeResult = &v
	return s
}

type ContractBenefitConsumeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ContractBenefitConsumeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContractBenefitConsumeResponse) String() string {
	return tea.Prettify(s)
}

func (s ContractBenefitConsumeResponse) GoString() string {
	return s.String()
}

func (s *ContractBenefitConsumeResponse) SetHeaders(v map[string]*string) *ContractBenefitConsumeResponse {
	s.Headers = v
	return s
}

func (s *ContractBenefitConsumeResponse) SetStatusCode(v int32) *ContractBenefitConsumeResponse {
	s.StatusCode = &v
	return s
}

func (s *ContractBenefitConsumeResponse) SetBody(v *ContractBenefitConsumeResponseBody) *ContractBenefitConsumeResponse {
	s.Body = v
	return s
}

type EsignQueryApprovalInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EsignQueryApprovalInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryApprovalInfoHeaders) GoString() string {
	return s.String()
}

func (s *EsignQueryApprovalInfoHeaders) SetCommonHeaders(v map[string]*string) *EsignQueryApprovalInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EsignQueryApprovalInfoHeaders) SetXAcsDingtalkAccessToken(v string) *EsignQueryApprovalInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EsignQueryApprovalInfoRequest struct {
	CorpId      *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	EsignFlowId *string `json:"esignFlowId,omitempty" xml:"esignFlowId,omitempty"`
	UnionId     *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s EsignQueryApprovalInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryApprovalInfoRequest) GoString() string {
	return s.String()
}

func (s *EsignQueryApprovalInfoRequest) SetCorpId(v string) *EsignQueryApprovalInfoRequest {
	s.CorpId = &v
	return s
}

func (s *EsignQueryApprovalInfoRequest) SetEsignFlowId(v string) *EsignQueryApprovalInfoRequest {
	s.EsignFlowId = &v
	return s
}

func (s *EsignQueryApprovalInfoRequest) SetUnionId(v string) *EsignQueryApprovalInfoRequest {
	s.UnionId = &v
	return s
}

type EsignQueryApprovalInfoResponseBody struct {
	Result  *EsignQueryApprovalInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s EsignQueryApprovalInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryApprovalInfoResponseBody) GoString() string {
	return s.String()
}

func (s *EsignQueryApprovalInfoResponseBody) SetResult(v *EsignQueryApprovalInfoResponseBodyResult) *EsignQueryApprovalInfoResponseBody {
	s.Result = v
	return s
}

func (s *EsignQueryApprovalInfoResponseBody) SetSuccess(v bool) *EsignQueryApprovalInfoResponseBody {
	s.Success = &v
	return s
}

type EsignQueryApprovalInfoResponseBodyResult struct {
	BpmsProcessBusinessId  *string `json:"bpmsProcessBusinessId,omitempty" xml:"bpmsProcessBusinessId,omitempty"`
	BpmsProcessInstanceId  *string `json:"bpmsProcessInstanceId,omitempty" xml:"bpmsProcessInstanceId,omitempty"`
	BpmsProcessInstanceUrl *string `json:"bpmsProcessInstanceUrl,omitempty" xml:"bpmsProcessInstanceUrl,omitempty"`
}

func (s EsignQueryApprovalInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryApprovalInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *EsignQueryApprovalInfoResponseBodyResult) SetBpmsProcessBusinessId(v string) *EsignQueryApprovalInfoResponseBodyResult {
	s.BpmsProcessBusinessId = &v
	return s
}

func (s *EsignQueryApprovalInfoResponseBodyResult) SetBpmsProcessInstanceId(v string) *EsignQueryApprovalInfoResponseBodyResult {
	s.BpmsProcessInstanceId = &v
	return s
}

func (s *EsignQueryApprovalInfoResponseBodyResult) SetBpmsProcessInstanceUrl(v string) *EsignQueryApprovalInfoResponseBodyResult {
	s.BpmsProcessInstanceUrl = &v
	return s
}

type EsignQueryApprovalInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EsignQueryApprovalInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EsignQueryApprovalInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryApprovalInfoResponse) GoString() string {
	return s.String()
}

func (s *EsignQueryApprovalInfoResponse) SetHeaders(v map[string]*string) *EsignQueryApprovalInfoResponse {
	s.Headers = v
	return s
}

func (s *EsignQueryApprovalInfoResponse) SetStatusCode(v int32) *EsignQueryApprovalInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *EsignQueryApprovalInfoResponse) SetBody(v *EsignQueryApprovalInfoResponseBody) *EsignQueryApprovalInfoResponse {
	s.Body = v
	return s
}

type EsignQueryGrantInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EsignQueryGrantInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryGrantInfoHeaders) GoString() string {
	return s.String()
}

func (s *EsignQueryGrantInfoHeaders) SetCommonHeaders(v map[string]*string) *EsignQueryGrantInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EsignQueryGrantInfoHeaders) SetXAcsDingtalkAccessToken(v string) *EsignQueryGrantInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EsignQueryGrantInfoRequest struct {
	CorpId    *string            `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Extension map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
	UnionId   *string            `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s EsignQueryGrantInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryGrantInfoRequest) GoString() string {
	return s.String()
}

func (s *EsignQueryGrantInfoRequest) SetCorpId(v string) *EsignQueryGrantInfoRequest {
	s.CorpId = &v
	return s
}

func (s *EsignQueryGrantInfoRequest) SetExtension(v map[string]*string) *EsignQueryGrantInfoRequest {
	s.Extension = v
	return s
}

func (s *EsignQueryGrantInfoRequest) SetUnionId(v string) *EsignQueryGrantInfoRequest {
	s.UnionId = &v
	return s
}

type EsignQueryGrantInfoResponseBody struct {
	Result  *EsignQueryGrantInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s EsignQueryGrantInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryGrantInfoResponseBody) GoString() string {
	return s.String()
}

func (s *EsignQueryGrantInfoResponseBody) SetResult(v *EsignQueryGrantInfoResponseBodyResult) *EsignQueryGrantInfoResponseBody {
	s.Result = v
	return s
}

func (s *EsignQueryGrantInfoResponseBody) SetSuccess(v bool) *EsignQueryGrantInfoResponseBody {
	s.Success = &v
	return s
}

type EsignQueryGrantInfoResponseBodyResult struct {
	LegalPerson         *string `json:"legalPerson,omitempty" xml:"legalPerson,omitempty"`
	MobilePhoneNumber   *string `json:"mobilePhoneNumber,omitempty" xml:"mobilePhoneNumber,omitempty"`
	OrgName             *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	StateCode           *string `json:"stateCode,omitempty" xml:"stateCode,omitempty"`
	UnifiedSocialCredit *string `json:"unifiedSocialCredit,omitempty" xml:"unifiedSocialCredit,omitempty"`
	UserName            *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s EsignQueryGrantInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryGrantInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *EsignQueryGrantInfoResponseBodyResult) SetLegalPerson(v string) *EsignQueryGrantInfoResponseBodyResult {
	s.LegalPerson = &v
	return s
}

func (s *EsignQueryGrantInfoResponseBodyResult) SetMobilePhoneNumber(v string) *EsignQueryGrantInfoResponseBodyResult {
	s.MobilePhoneNumber = &v
	return s
}

func (s *EsignQueryGrantInfoResponseBodyResult) SetOrgName(v string) *EsignQueryGrantInfoResponseBodyResult {
	s.OrgName = &v
	return s
}

func (s *EsignQueryGrantInfoResponseBodyResult) SetStateCode(v string) *EsignQueryGrantInfoResponseBodyResult {
	s.StateCode = &v
	return s
}

func (s *EsignQueryGrantInfoResponseBodyResult) SetUnifiedSocialCredit(v string) *EsignQueryGrantInfoResponseBodyResult {
	s.UnifiedSocialCredit = &v
	return s
}

func (s *EsignQueryGrantInfoResponseBodyResult) SetUserName(v string) *EsignQueryGrantInfoResponseBodyResult {
	s.UserName = &v
	return s
}

type EsignQueryGrantInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EsignQueryGrantInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EsignQueryGrantInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryGrantInfoResponse) GoString() string {
	return s.String()
}

func (s *EsignQueryGrantInfoResponse) SetHeaders(v map[string]*string) *EsignQueryGrantInfoResponse {
	s.Headers = v
	return s
}

func (s *EsignQueryGrantInfoResponse) SetStatusCode(v int32) *EsignQueryGrantInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *EsignQueryGrantInfoResponse) SetBody(v *EsignQueryGrantInfoResponseBody) *EsignQueryGrantInfoResponse {
	s.Body = v
	return s
}

type EsignQueryIdentityByTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EsignQueryIdentityByTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryIdentityByTicketHeaders) GoString() string {
	return s.String()
}

func (s *EsignQueryIdentityByTicketHeaders) SetCommonHeaders(v map[string]*string) *EsignQueryIdentityByTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EsignQueryIdentityByTicketHeaders) SetXAcsDingtalkAccessToken(v string) *EsignQueryIdentityByTicketHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EsignQueryIdentityByTicketRequest struct {
	CorpId    *string            `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Extension map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
	Ticket    *string            `json:"ticket,omitempty" xml:"ticket,omitempty"`
}

func (s EsignQueryIdentityByTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryIdentityByTicketRequest) GoString() string {
	return s.String()
}

func (s *EsignQueryIdentityByTicketRequest) SetCorpId(v string) *EsignQueryIdentityByTicketRequest {
	s.CorpId = &v
	return s
}

func (s *EsignQueryIdentityByTicketRequest) SetExtension(v map[string]*string) *EsignQueryIdentityByTicketRequest {
	s.Extension = v
	return s
}

func (s *EsignQueryIdentityByTicketRequest) SetTicket(v string) *EsignQueryIdentityByTicketRequest {
	s.Ticket = &v
	return s
}

type EsignQueryIdentityByTicketResponseBody struct {
	Result  *EsignQueryIdentityByTicketResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s EsignQueryIdentityByTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryIdentityByTicketResponseBody) GoString() string {
	return s.String()
}

func (s *EsignQueryIdentityByTicketResponseBody) SetResult(v *EsignQueryIdentityByTicketResponseBodyResult) *EsignQueryIdentityByTicketResponseBody {
	s.Result = v
	return s
}

func (s *EsignQueryIdentityByTicketResponseBody) SetSuccess(v bool) *EsignQueryIdentityByTicketResponseBody {
	s.Success = &v
	return s
}

type EsignQueryIdentityByTicketResponseBodyResult struct {
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s EsignQueryIdentityByTicketResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryIdentityByTicketResponseBodyResult) GoString() string {
	return s.String()
}

func (s *EsignQueryIdentityByTicketResponseBodyResult) SetUnionId(v string) *EsignQueryIdentityByTicketResponseBodyResult {
	s.UnionId = &v
	return s
}

type EsignQueryIdentityByTicketResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EsignQueryIdentityByTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EsignQueryIdentityByTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryIdentityByTicketResponse) GoString() string {
	return s.String()
}

func (s *EsignQueryIdentityByTicketResponse) SetHeaders(v map[string]*string) *EsignQueryIdentityByTicketResponse {
	s.Headers = v
	return s
}

func (s *EsignQueryIdentityByTicketResponse) SetStatusCode(v int32) *EsignQueryIdentityByTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *EsignQueryIdentityByTicketResponse) SetBody(v *EsignQueryIdentityByTicketResponseBody) *EsignQueryIdentityByTicketResponse {
	s.Body = v
	return s
}

type EsignSyncEventHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EsignSyncEventHeaders) String() string {
	return tea.Prettify(s)
}

func (s EsignSyncEventHeaders) GoString() string {
	return s.String()
}

func (s *EsignSyncEventHeaders) SetCommonHeaders(v map[string]*string) *EsignSyncEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EsignSyncEventHeaders) SetXAcsDingtalkAccessToken(v string) *EsignSyncEventHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EsignSyncEventRequest struct {
	Action    *string            `json:"action,omitempty" xml:"action,omitempty"`
	CorpId    *string            `json:"corpId,omitempty" xml:"corpId,omitempty"`
	EsignData *string            `json:"esignData,omitempty" xml:"esignData,omitempty"`
	Extension map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
	UnionId   *string            `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s EsignSyncEventRequest) String() string {
	return tea.Prettify(s)
}

func (s EsignSyncEventRequest) GoString() string {
	return s.String()
}

func (s *EsignSyncEventRequest) SetAction(v string) *EsignSyncEventRequest {
	s.Action = &v
	return s
}

func (s *EsignSyncEventRequest) SetCorpId(v string) *EsignSyncEventRequest {
	s.CorpId = &v
	return s
}

func (s *EsignSyncEventRequest) SetEsignData(v string) *EsignSyncEventRequest {
	s.EsignData = &v
	return s
}

func (s *EsignSyncEventRequest) SetExtension(v map[string]*string) *EsignSyncEventRequest {
	s.Extension = v
	return s
}

func (s *EsignSyncEventRequest) SetUnionId(v string) *EsignSyncEventRequest {
	s.UnionId = &v
	return s
}

type EsignSyncEventResponseBody struct {
	Result  *EsignSyncEventResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s EsignSyncEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EsignSyncEventResponseBody) GoString() string {
	return s.String()
}

func (s *EsignSyncEventResponseBody) SetResult(v *EsignSyncEventResponseBodyResult) *EsignSyncEventResponseBody {
	s.Result = v
	return s
}

func (s *EsignSyncEventResponseBody) SetSuccess(v bool) *EsignSyncEventResponseBody {
	s.Success = &v
	return s
}

type EsignSyncEventResponseBodyResult struct {
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s EsignSyncEventResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s EsignSyncEventResponseBodyResult) GoString() string {
	return s.String()
}

func (s *EsignSyncEventResponseBodyResult) SetMessage(v string) *EsignSyncEventResponseBodyResult {
	s.Message = &v
	return s
}

type EsignSyncEventResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EsignSyncEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EsignSyncEventResponse) String() string {
	return tea.Prettify(s)
}

func (s EsignSyncEventResponse) GoString() string {
	return s.String()
}

func (s *EsignSyncEventResponse) SetHeaders(v map[string]*string) *EsignSyncEventResponse {
	s.Headers = v
	return s
}

func (s *EsignSyncEventResponse) SetStatusCode(v int32) *EsignSyncEventResponse {
	s.StatusCode = &v
	return s
}

func (s *EsignSyncEventResponse) SetBody(v *EsignSyncEventResponseBody) *EsignSyncEventResponse {
	s.Body = v
	return s
}

type QueryAdvancedContractVersionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAdvancedContractVersionHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAdvancedContractVersionHeaders) GoString() string {
	return s.String()
}

func (s *QueryAdvancedContractVersionHeaders) SetCommonHeaders(v map[string]*string) *QueryAdvancedContractVersionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAdvancedContractVersionHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAdvancedContractVersionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAdvancedContractVersionRequest struct {
	CorpId    *string            `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Extension map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
}

func (s QueryAdvancedContractVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAdvancedContractVersionRequest) GoString() string {
	return s.String()
}

func (s *QueryAdvancedContractVersionRequest) SetCorpId(v string) *QueryAdvancedContractVersionRequest {
	s.CorpId = &v
	return s
}

func (s *QueryAdvancedContractVersionRequest) SetExtension(v map[string]*string) *QueryAdvancedContractVersionRequest {
	s.Extension = v
	return s
}

type QueryAdvancedContractVersionResponseBody struct {
	Result  *QueryAdvancedContractVersionResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryAdvancedContractVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAdvancedContractVersionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAdvancedContractVersionResponseBody) SetResult(v *QueryAdvancedContractVersionResponseBodyResult) *QueryAdvancedContractVersionResponseBody {
	s.Result = v
	return s
}

func (s *QueryAdvancedContractVersionResponseBody) SetSuccess(v bool) *QueryAdvancedContractVersionResponseBody {
	s.Success = &v
	return s
}

type QueryAdvancedContractVersionResponseBodyResult struct {
	EnableEsignAttachmentSign *bool              `json:"enableEsignAttachmentSign,omitempty" xml:"enableEsignAttachmentSign,omitempty"`
	Extension                 map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
	Version                   *string            `json:"version,omitempty" xml:"version,omitempty"`
}

func (s QueryAdvancedContractVersionResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryAdvancedContractVersionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryAdvancedContractVersionResponseBodyResult) SetEnableEsignAttachmentSign(v bool) *QueryAdvancedContractVersionResponseBodyResult {
	s.EnableEsignAttachmentSign = &v
	return s
}

func (s *QueryAdvancedContractVersionResponseBodyResult) SetExtension(v map[string]*string) *QueryAdvancedContractVersionResponseBodyResult {
	s.Extension = v
	return s
}

func (s *QueryAdvancedContractVersionResponseBodyResult) SetVersion(v string) *QueryAdvancedContractVersionResponseBodyResult {
	s.Version = &v
	return s
}

type QueryAdvancedContractVersionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAdvancedContractVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAdvancedContractVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAdvancedContractVersionResponse) GoString() string {
	return s.String()
}

func (s *QueryAdvancedContractVersionResponse) SetHeaders(v map[string]*string) *QueryAdvancedContractVersionResponse {
	s.Headers = v
	return s
}

func (s *QueryAdvancedContractVersionResponse) SetStatusCode(v int32) *QueryAdvancedContractVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAdvancedContractVersionResponse) SetBody(v *QueryAdvancedContractVersionResponseBody) *QueryAdvancedContractVersionResponse {
	s.Body = v
	return s
}

type SendContractCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendContractCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendContractCardHeaders) GoString() string {
	return s.String()
}

func (s *SendContractCardHeaders) SetCommonHeaders(v map[string]*string) *SendContractCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendContractCardHeaders) SetXAcsDingtalkAccessToken(v string) *SendContractCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendContractCardRequest struct {
	CardType          *string                              `json:"cardType,omitempty" xml:"cardType,omitempty"`
	ContractInfo      *SendContractCardRequestContractInfo `json:"contractInfo,omitempty" xml:"contractInfo,omitempty" type:"Struct"`
	CorpId            *string                              `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Extension         map[string]*string                   `json:"extension,omitempty" xml:"extension,omitempty"`
	ProcessInstanceId *string                              `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	ReceiveGroups     []*string                            `json:"receiveGroups,omitempty" xml:"receiveGroups,omitempty" type:"Repeated"`
	Receivers         []*SendContractCardRequestReceivers  `json:"receivers,omitempty" xml:"receivers,omitempty" type:"Repeated"`
	Sender            *SendContractCardRequestSender       `json:"sender,omitempty" xml:"sender,omitempty" type:"Struct"`
	SyncSingleChat    *bool                                `json:"syncSingleChat,omitempty" xml:"syncSingleChat,omitempty"`
}

func (s SendContractCardRequest) String() string {
	return tea.Prettify(s)
}

func (s SendContractCardRequest) GoString() string {
	return s.String()
}

func (s *SendContractCardRequest) SetCardType(v string) *SendContractCardRequest {
	s.CardType = &v
	return s
}

func (s *SendContractCardRequest) SetContractInfo(v *SendContractCardRequestContractInfo) *SendContractCardRequest {
	s.ContractInfo = v
	return s
}

func (s *SendContractCardRequest) SetCorpId(v string) *SendContractCardRequest {
	s.CorpId = &v
	return s
}

func (s *SendContractCardRequest) SetExtension(v map[string]*string) *SendContractCardRequest {
	s.Extension = v
	return s
}

func (s *SendContractCardRequest) SetProcessInstanceId(v string) *SendContractCardRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *SendContractCardRequest) SetReceiveGroups(v []*string) *SendContractCardRequest {
	s.ReceiveGroups = v
	return s
}

func (s *SendContractCardRequest) SetReceivers(v []*SendContractCardRequestReceivers) *SendContractCardRequest {
	s.Receivers = v
	return s
}

func (s *SendContractCardRequest) SetSender(v *SendContractCardRequestSender) *SendContractCardRequest {
	s.Sender = v
	return s
}

func (s *SendContractCardRequest) SetSyncSingleChat(v bool) *SendContractCardRequest {
	s.SyncSingleChat = &v
	return s
}

type SendContractCardRequestContractInfo struct {
	ContractCode *string `json:"contractCode,omitempty" xml:"contractCode,omitempty"`
	ContractName *string `json:"contractName,omitempty" xml:"contractName,omitempty"`
	CreateTime   *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	SignUserName *string `json:"signUserName,omitempty" xml:"signUserName,omitempty"`
}

func (s SendContractCardRequestContractInfo) String() string {
	return tea.Prettify(s)
}

func (s SendContractCardRequestContractInfo) GoString() string {
	return s.String()
}

func (s *SendContractCardRequestContractInfo) SetContractCode(v string) *SendContractCardRequestContractInfo {
	s.ContractCode = &v
	return s
}

func (s *SendContractCardRequestContractInfo) SetContractName(v string) *SendContractCardRequestContractInfo {
	s.ContractName = &v
	return s
}

func (s *SendContractCardRequestContractInfo) SetCreateTime(v int64) *SendContractCardRequestContractInfo {
	s.CreateTime = &v
	return s
}

func (s *SendContractCardRequestContractInfo) SetSignUserName(v string) *SendContractCardRequestContractInfo {
	s.SignUserName = &v
	return s
}

type SendContractCardRequestReceivers struct {
	CorpId   *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	UserId   *string `json:"userId,omitempty" xml:"userId,omitempty"`
	UserType *string `json:"userType,omitempty" xml:"userType,omitempty"`
}

func (s SendContractCardRequestReceivers) String() string {
	return tea.Prettify(s)
}

func (s SendContractCardRequestReceivers) GoString() string {
	return s.String()
}

func (s *SendContractCardRequestReceivers) SetCorpId(v string) *SendContractCardRequestReceivers {
	s.CorpId = &v
	return s
}

func (s *SendContractCardRequestReceivers) SetUserId(v string) *SendContractCardRequestReceivers {
	s.UserId = &v
	return s
}

func (s *SendContractCardRequestReceivers) SetUserType(v string) *SendContractCardRequestReceivers {
	s.UserType = &v
	return s
}

type SendContractCardRequestSender struct {
	CorpId   *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	UserId   *string `json:"userId,omitempty" xml:"userId,omitempty"`
	UserType *string `json:"userType,omitempty" xml:"userType,omitempty"`
}

func (s SendContractCardRequestSender) String() string {
	return tea.Prettify(s)
}

func (s SendContractCardRequestSender) GoString() string {
	return s.String()
}

func (s *SendContractCardRequestSender) SetCorpId(v string) *SendContractCardRequestSender {
	s.CorpId = &v
	return s
}

func (s *SendContractCardRequestSender) SetUserId(v string) *SendContractCardRequestSender {
	s.UserId = &v
	return s
}

func (s *SendContractCardRequestSender) SetUserType(v string) *SendContractCardRequestSender {
	s.UserType = &v
	return s
}

type SendContractCardResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SendContractCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendContractCardResponseBody) GoString() string {
	return s.String()
}

func (s *SendContractCardResponseBody) SetSuccess(v bool) *SendContractCardResponseBody {
	s.Success = &v
	return s
}

type SendContractCardResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendContractCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendContractCardResponse) String() string {
	return tea.Prettify(s)
}

func (s SendContractCardResponse) GoString() string {
	return s.String()
}

func (s *SendContractCardResponse) SetHeaders(v map[string]*string) *SendContractCardResponse {
	s.Headers = v
	return s
}

func (s *SendContractCardResponse) SetStatusCode(v int32) *SendContractCardResponse {
	s.StatusCode = &v
	return s
}

func (s *SendContractCardResponse) SetBody(v *SendContractCardResponseBody) *SendContractCardResponse {
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

func (client *Client) ContractBenefitConsumeWithOptions(request *ContractBenefitConsumeRequest, headers *ContractBenefitConsumeHeaders, runtime *util.RuntimeOptions) (_result *ContractBenefitConsumeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BenefitPoint)) {
		body["benefitPoint"] = request.BenefitPoint
	}

	if !tea.BoolValue(util.IsUnset(request.BizRequestId)) {
		body["bizRequestId"] = request.BizRequestId
	}

	if !tea.BoolValue(util.IsUnset(request.ConsumeQuota)) {
		body["consumeQuota"] = request.ConsumeQuota
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtParams)) {
		body["extParams"] = request.ExtParams
	}

	if !tea.BoolValue(util.IsUnset(request.IsvCorpId)) {
		body["isvCorpId"] = request.IsvCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.OptUnionId)) {
		body["optUnionId"] = request.OptUnionId
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
		Action:      tea.String("ContractBenefitConsume"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/benefits/consume"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ContractBenefitConsumeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ContractBenefitConsume(request *ContractBenefitConsumeRequest) (_result *ContractBenefitConsumeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ContractBenefitConsumeHeaders{}
	_result = &ContractBenefitConsumeResponse{}
	_body, _err := client.ContractBenefitConsumeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EsignQueryApprovalInfoWithOptions(request *EsignQueryApprovalInfoRequest, headers *EsignQueryApprovalInfoHeaders, runtime *util.RuntimeOptions) (_result *EsignQueryApprovalInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.EsignFlowId)) {
		body["esignFlowId"] = request.EsignFlowId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("EsignQueryApprovalInfo"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/esign/approvalInfos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &EsignQueryApprovalInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EsignQueryApprovalInfo(request *EsignQueryApprovalInfoRequest) (_result *EsignQueryApprovalInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EsignQueryApprovalInfoHeaders{}
	_result = &EsignQueryApprovalInfoResponse{}
	_body, _err := client.EsignQueryApprovalInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EsignQueryGrantInfoWithOptions(request *EsignQueryGrantInfoRequest, headers *EsignQueryGrantInfoHeaders, runtime *util.RuntimeOptions) (_result *EsignQueryGrantInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Extension)) {
		body["extension"] = request.Extension
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("EsignQueryGrantInfo"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/esign/anthInfos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &EsignQueryGrantInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EsignQueryGrantInfo(request *EsignQueryGrantInfoRequest) (_result *EsignQueryGrantInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EsignQueryGrantInfoHeaders{}
	_result = &EsignQueryGrantInfoResponse{}
	_body, _err := client.EsignQueryGrantInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EsignQueryIdentityByTicketWithOptions(request *EsignQueryIdentityByTicketRequest, headers *EsignQueryIdentityByTicketHeaders, runtime *util.RuntimeOptions) (_result *EsignQueryIdentityByTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Extension)) {
		body["extension"] = request.Extension
	}

	if !tea.BoolValue(util.IsUnset(request.Ticket)) {
		body["ticket"] = request.Ticket
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
		Action:      tea.String("EsignQueryIdentityByTicket"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/esign/tickets/identities/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &EsignQueryIdentityByTicketResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EsignQueryIdentityByTicket(request *EsignQueryIdentityByTicketRequest) (_result *EsignQueryIdentityByTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EsignQueryIdentityByTicketHeaders{}
	_result = &EsignQueryIdentityByTicketResponse{}
	_body, _err := client.EsignQueryIdentityByTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EsignSyncEventWithOptions(request *EsignSyncEventRequest, headers *EsignSyncEventHeaders, runtime *util.RuntimeOptions) (_result *EsignSyncEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		body["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.EsignData)) {
		body["esignData"] = request.EsignData
	}

	if !tea.BoolValue(util.IsUnset(request.Extension)) {
		body["extension"] = request.Extension
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("EsignSyncEvent"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/esign/events/sync"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &EsignSyncEventResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EsignSyncEvent(request *EsignSyncEventRequest) (_result *EsignSyncEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EsignSyncEventHeaders{}
	_result = &EsignSyncEventResponse{}
	_body, _err := client.EsignSyncEventWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAdvancedContractVersionWithOptions(request *QueryAdvancedContractVersionRequest, headers *QueryAdvancedContractVersionHeaders, runtime *util.RuntimeOptions) (_result *QueryAdvancedContractVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Extension)) {
		body["extension"] = request.Extension
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
		Action:      tea.String("QueryAdvancedContractVersion"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/versions/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAdvancedContractVersionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAdvancedContractVersion(request *QueryAdvancedContractVersionRequest) (_result *QueryAdvancedContractVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAdvancedContractVersionHeaders{}
	_result = &QueryAdvancedContractVersionResponse{}
	_body, _err := client.QueryAdvancedContractVersionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendContractCardWithOptions(request *SendContractCardRequest, headers *SendContractCardHeaders, runtime *util.RuntimeOptions) (_result *SendContractCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CardType)) {
		body["cardType"] = request.CardType
	}

	if !tea.BoolValue(util.IsUnset(request.ContractInfo)) {
		body["contractInfo"] = request.ContractInfo
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Extension)) {
		body["extension"] = request.Extension
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		body["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiveGroups)) {
		body["receiveGroups"] = request.ReceiveGroups
	}

	if !tea.BoolValue(util.IsUnset(request.Receivers)) {
		body["receivers"] = request.Receivers
	}

	if !tea.BoolValue(util.IsUnset(request.Sender)) {
		body["sender"] = request.Sender
	}

	if !tea.BoolValue(util.IsUnset(request.SyncSingleChat)) {
		body["syncSingleChat"] = request.SyncSingleChat
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
		Action:      tea.String("SendContractCard"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/cards/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendContractCardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendContractCard(request *SendContractCardRequest) (_result *SendContractCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendContractCardHeaders{}
	_result = &SendContractCardResponse{}
	_body, _err := client.SendContractCardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
