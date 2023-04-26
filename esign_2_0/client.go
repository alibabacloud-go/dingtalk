// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package esign_2_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ApprovalListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ServiceGroup            *string            `json:"serviceGroup,omitempty" xml:"serviceGroup,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ApprovalListHeaders) String() string {
	return tea.Prettify(s)
}

func (s ApprovalListHeaders) GoString() string {
	return s.String()
}

func (s *ApprovalListHeaders) SetCommonHeaders(v map[string]*string) *ApprovalListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ApprovalListHeaders) SetServiceGroup(v string) *ApprovalListHeaders {
	s.ServiceGroup = &v
	return s
}

func (s *ApprovalListHeaders) SetXAcsDingtalkAccessToken(v string) *ApprovalListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ApprovalListResponseBody struct {
	Data []*ApprovalListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s ApprovalListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApprovalListResponseBody) GoString() string {
	return s.String()
}

func (s *ApprovalListResponseBody) SetData(v []*ApprovalListResponseBodyData) *ApprovalListResponseBody {
	s.Data = v
	return s
}

type ApprovalListResponseBodyData struct {
	ApprovalName       *string                                      `json:"approvalName,omitempty" xml:"approvalName,omitempty"`
	ApprovalNodes      []*ApprovalListResponseBodyDataApprovalNodes `json:"approvalNodes,omitempty" xml:"approvalNodes,omitempty" type:"Repeated"`
	EndTime            *float32                                     `json:"endTime,omitempty" xml:"endTime,omitempty"`
	RefuseReason       *string                                      `json:"refuseReason,omitempty" xml:"refuseReason,omitempty"`
	SealIdImg          *string                                      `json:"sealIdImg,omitempty" xml:"sealIdImg,omitempty"`
	SponsorAccountName *string                                      `json:"sponsorAccountName,omitempty" xml:"sponsorAccountName,omitempty"`
	StartTime          *float32                                     `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status             *string                                      `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ApprovalListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ApprovalListResponseBodyData) GoString() string {
	return s.String()
}

func (s *ApprovalListResponseBodyData) SetApprovalName(v string) *ApprovalListResponseBodyData {
	s.ApprovalName = &v
	return s
}

func (s *ApprovalListResponseBodyData) SetApprovalNodes(v []*ApprovalListResponseBodyDataApprovalNodes) *ApprovalListResponseBodyData {
	s.ApprovalNodes = v
	return s
}

func (s *ApprovalListResponseBodyData) SetEndTime(v float32) *ApprovalListResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ApprovalListResponseBodyData) SetRefuseReason(v string) *ApprovalListResponseBodyData {
	s.RefuseReason = &v
	return s
}

func (s *ApprovalListResponseBodyData) SetSealIdImg(v string) *ApprovalListResponseBodyData {
	s.SealIdImg = &v
	return s
}

func (s *ApprovalListResponseBodyData) SetSponsorAccountName(v string) *ApprovalListResponseBodyData {
	s.SponsorAccountName = &v
	return s
}

func (s *ApprovalListResponseBodyData) SetStartTime(v float32) *ApprovalListResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ApprovalListResponseBodyData) SetStatus(v string) *ApprovalListResponseBodyData {
	s.Status = &v
	return s
}

type ApprovalListResponseBodyDataApprovalNodes struct {
	ApprovalTime *string  `json:"approvalTime,omitempty" xml:"approvalTime,omitempty"`
	ApproverName *string  `json:"approverName,omitempty" xml:"approverName,omitempty"`
	StartTime    *float32 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status       *string  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ApprovalListResponseBodyDataApprovalNodes) String() string {
	return tea.Prettify(s)
}

func (s ApprovalListResponseBodyDataApprovalNodes) GoString() string {
	return s.String()
}

func (s *ApprovalListResponseBodyDataApprovalNodes) SetApprovalTime(v string) *ApprovalListResponseBodyDataApprovalNodes {
	s.ApprovalTime = &v
	return s
}

func (s *ApprovalListResponseBodyDataApprovalNodes) SetApproverName(v string) *ApprovalListResponseBodyDataApprovalNodes {
	s.ApproverName = &v
	return s
}

func (s *ApprovalListResponseBodyDataApprovalNodes) SetStartTime(v float32) *ApprovalListResponseBodyDataApprovalNodes {
	s.StartTime = &v
	return s
}

func (s *ApprovalListResponseBodyDataApprovalNodes) SetStatus(v string) *ApprovalListResponseBodyDataApprovalNodes {
	s.Status = &v
	return s
}

type ApprovalListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApprovalListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApprovalListResponse) String() string {
	return tea.Prettify(s)
}

func (s ApprovalListResponse) GoString() string {
	return s.String()
}

func (s *ApprovalListResponse) SetHeaders(v map[string]*string) *ApprovalListResponse {
	s.Headers = v
	return s
}

func (s *ApprovalListResponse) SetStatusCode(v int32) *ApprovalListResponse {
	s.StatusCode = &v
	return s
}

func (s *ApprovalListResponse) SetBody(v *ApprovalListResponseBody) *ApprovalListResponse {
	s.Body = v
	return s
}

type CancelCorpAuthHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ServiceGroup            *string            `json:"serviceGroup,omitempty" xml:"serviceGroup,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CancelCorpAuthHeaders) String() string {
	return tea.Prettify(s)
}

func (s CancelCorpAuthHeaders) GoString() string {
	return s.String()
}

func (s *CancelCorpAuthHeaders) SetCommonHeaders(v map[string]*string) *CancelCorpAuthHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CancelCorpAuthHeaders) SetServiceGroup(v string) *CancelCorpAuthHeaders {
	s.ServiceGroup = &v
	return s
}

func (s *CancelCorpAuthHeaders) SetXAcsDingtalkAccessToken(v string) *CancelCorpAuthHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CancelCorpAuthRequest struct {
}

func (s CancelCorpAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelCorpAuthRequest) GoString() string {
	return s.String()
}

type CancelCorpAuthResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CancelCorpAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelCorpAuthResponseBody) GoString() string {
	return s.String()
}

func (s *CancelCorpAuthResponseBody) SetResult(v bool) *CancelCorpAuthResponseBody {
	s.Result = &v
	return s
}

type CancelCorpAuthResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelCorpAuthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelCorpAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelCorpAuthResponse) GoString() string {
	return s.String()
}

func (s *CancelCorpAuthResponse) SetHeaders(v map[string]*string) *CancelCorpAuthResponse {
	s.Headers = v
	return s
}

func (s *CancelCorpAuthResponse) SetStatusCode(v int32) *CancelCorpAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelCorpAuthResponse) SetBody(v *CancelCorpAuthResponseBody) *CancelCorpAuthResponse {
	s.Body = v
	return s
}

type ChannelOrdersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ServiceGroup            *string            `json:"serviceGroup,omitempty" xml:"serviceGroup,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ChannelOrdersHeaders) String() string {
	return tea.Prettify(s)
}

func (s ChannelOrdersHeaders) GoString() string {
	return s.String()
}

func (s *ChannelOrdersHeaders) SetCommonHeaders(v map[string]*string) *ChannelOrdersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ChannelOrdersHeaders) SetServiceGroup(v string) *ChannelOrdersHeaders {
	s.ServiceGroup = &v
	return s
}

func (s *ChannelOrdersHeaders) SetXAcsDingtalkAccessToken(v string) *ChannelOrdersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ChannelOrdersRequest struct {
	ItemCode        *string  `json:"itemCode,omitempty" xml:"itemCode,omitempty"`
	ItemName        *string  `json:"itemName,omitempty" xml:"itemName,omitempty"`
	OrderCreateTime *float32 `json:"orderCreateTime,omitempty" xml:"orderCreateTime,omitempty"`
	OrderId         *string  `json:"orderId,omitempty" xml:"orderId,omitempty"`
	PayFee          *float32 `json:"payFee,omitempty" xml:"payFee,omitempty"`
	Quantity        *float32 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

func (s ChannelOrdersRequest) String() string {
	return tea.Prettify(s)
}

func (s ChannelOrdersRequest) GoString() string {
	return s.String()
}

func (s *ChannelOrdersRequest) SetItemCode(v string) *ChannelOrdersRequest {
	s.ItemCode = &v
	return s
}

func (s *ChannelOrdersRequest) SetItemName(v string) *ChannelOrdersRequest {
	s.ItemName = &v
	return s
}

func (s *ChannelOrdersRequest) SetOrderCreateTime(v float32) *ChannelOrdersRequest {
	s.OrderCreateTime = &v
	return s
}

func (s *ChannelOrdersRequest) SetOrderId(v string) *ChannelOrdersRequest {
	s.OrderId = &v
	return s
}

func (s *ChannelOrdersRequest) SetPayFee(v float32) *ChannelOrdersRequest {
	s.PayFee = &v
	return s
}

func (s *ChannelOrdersRequest) SetQuantity(v float32) *ChannelOrdersRequest {
	s.Quantity = &v
	return s
}

type ChannelOrdersResponseBody struct {
	EsignOrderId *string `json:"esignOrderId,omitempty" xml:"esignOrderId,omitempty"`
}

func (s ChannelOrdersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChannelOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *ChannelOrdersResponseBody) SetEsignOrderId(v string) *ChannelOrdersResponseBody {
	s.EsignOrderId = &v
	return s
}

type ChannelOrdersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ChannelOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChannelOrdersResponse) String() string {
	return tea.Prettify(s)
}

func (s ChannelOrdersResponse) GoString() string {
	return s.String()
}

func (s *ChannelOrdersResponse) SetHeaders(v map[string]*string) *ChannelOrdersResponse {
	s.Headers = v
	return s
}

func (s *ChannelOrdersResponse) SetStatusCode(v int32) *ChannelOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *ChannelOrdersResponse) SetBody(v *ChannelOrdersResponseBody) *ChannelOrdersResponse {
	s.Body = v
	return s
}

type CorpRealnameHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ServiceGroup            *string            `json:"serviceGroup,omitempty" xml:"serviceGroup,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CorpRealnameHeaders) String() string {
	return tea.Prettify(s)
}

func (s CorpRealnameHeaders) GoString() string {
	return s.String()
}

func (s *CorpRealnameHeaders) SetCommonHeaders(v map[string]*string) *CorpRealnameHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CorpRealnameHeaders) SetServiceGroup(v string) *CorpRealnameHeaders {
	s.ServiceGroup = &v
	return s
}

func (s *CorpRealnameHeaders) SetXAcsDingtalkAccessToken(v string) *CorpRealnameHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CorpRealnameRequest struct {
	RedirectUrl *string `json:"redirectUrl,omitempty" xml:"redirectUrl,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CorpRealnameRequest) String() string {
	return tea.Prettify(s)
}

func (s CorpRealnameRequest) GoString() string {
	return s.String()
}

func (s *CorpRealnameRequest) SetRedirectUrl(v string) *CorpRealnameRequest {
	s.RedirectUrl = &v
	return s
}

func (s *CorpRealnameRequest) SetUserId(v string) *CorpRealnameRequest {
	s.UserId = &v
	return s
}

type CorpRealnameResponseBody struct {
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	PcUrl     *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	TaskId    *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CorpRealnameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CorpRealnameResponseBody) GoString() string {
	return s.String()
}

func (s *CorpRealnameResponseBody) SetMobileUrl(v string) *CorpRealnameResponseBody {
	s.MobileUrl = &v
	return s
}

func (s *CorpRealnameResponseBody) SetPcUrl(v string) *CorpRealnameResponseBody {
	s.PcUrl = &v
	return s
}

func (s *CorpRealnameResponseBody) SetTaskId(v string) *CorpRealnameResponseBody {
	s.TaskId = &v
	return s
}

type CorpRealnameResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CorpRealnameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CorpRealnameResponse) String() string {
	return tea.Prettify(s)
}

func (s CorpRealnameResponse) GoString() string {
	return s.String()
}

func (s *CorpRealnameResponse) SetHeaders(v map[string]*string) *CorpRealnameResponse {
	s.Headers = v
	return s
}

func (s *CorpRealnameResponse) SetStatusCode(v int32) *CorpRealnameResponse {
	s.StatusCode = &v
	return s
}

func (s *CorpRealnameResponse) SetBody(v *CorpRealnameResponseBody) *CorpRealnameResponse {
	s.Body = v
	return s
}

type CreateDevelopersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ServiceGroup            *string            `json:"serviceGroup,omitempty" xml:"serviceGroup,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateDevelopersHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateDevelopersHeaders) GoString() string {
	return s.String()
}

func (s *CreateDevelopersHeaders) SetCommonHeaders(v map[string]*string) *CreateDevelopersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateDevelopersHeaders) SetServiceGroup(v string) *CreateDevelopersHeaders {
	s.ServiceGroup = &v
	return s
}

func (s *CreateDevelopersHeaders) SetXAcsDingtalkAccessToken(v string) *CreateDevelopersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateDevelopersRequest struct {
	NoticeUrl *string `json:"noticeUrl,omitempty" xml:"noticeUrl,omitempty"`
}

func (s CreateDevelopersRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDevelopersRequest) GoString() string {
	return s.String()
}

func (s *CreateDevelopersRequest) SetNoticeUrl(v string) *CreateDevelopersRequest {
	s.NoticeUrl = &v
	return s
}

type CreateDevelopersResponseBody struct {
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CreateDevelopersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDevelopersResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDevelopersResponseBody) SetData(v bool) *CreateDevelopersResponseBody {
	s.Data = &v
	return s
}

type CreateDevelopersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDevelopersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDevelopersResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDevelopersResponse) GoString() string {
	return s.String()
}

func (s *CreateDevelopersResponse) SetHeaders(v map[string]*string) *CreateDevelopersResponse {
	s.Headers = v
	return s
}

func (s *CreateDevelopersResponse) SetStatusCode(v int32) *CreateDevelopersResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDevelopersResponse) SetBody(v *CreateDevelopersResponseBody) *CreateDevelopersResponse {
	s.Body = v
	return s
}

type CreateProcessHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateProcessHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessHeaders) GoString() string {
	return s.String()
}

func (s *CreateProcessHeaders) SetCommonHeaders(v map[string]*string) *CreateProcessHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateProcessHeaders) SetXAcsDingtalkAccessToken(v string) *CreateProcessHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateProcessRequest struct {
	Ccs             []*CreateProcessRequestCcs          `json:"ccs,omitempty" xml:"ccs,omitempty" type:"Repeated"`
	Files           []*CreateProcessRequestFiles        `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	InitiatorUserId *string                             `json:"initiatorUserId,omitempty" xml:"initiatorUserId,omitempty"`
	Participants    []*CreateProcessRequestParticipants `json:"participants,omitempty" xml:"participants,omitempty" type:"Repeated"`
	RedirectUrl     *string                             `json:"redirectUrl,omitempty" xml:"redirectUrl,omitempty"`
	SignEndTime     *int64                              `json:"signEndTime,omitempty" xml:"signEndTime,omitempty"`
	SourceInfo      *CreateProcessRequestSourceInfo     `json:"sourceInfo,omitempty" xml:"sourceInfo,omitempty" type:"Struct"`
	TaskName        *string                             `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s CreateProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessRequest) GoString() string {
	return s.String()
}

func (s *CreateProcessRequest) SetCcs(v []*CreateProcessRequestCcs) *CreateProcessRequest {
	s.Ccs = v
	return s
}

func (s *CreateProcessRequest) SetFiles(v []*CreateProcessRequestFiles) *CreateProcessRequest {
	s.Files = v
	return s
}

func (s *CreateProcessRequest) SetInitiatorUserId(v string) *CreateProcessRequest {
	s.InitiatorUserId = &v
	return s
}

func (s *CreateProcessRequest) SetParticipants(v []*CreateProcessRequestParticipants) *CreateProcessRequest {
	s.Participants = v
	return s
}

func (s *CreateProcessRequest) SetRedirectUrl(v string) *CreateProcessRequest {
	s.RedirectUrl = &v
	return s
}

func (s *CreateProcessRequest) SetSignEndTime(v int64) *CreateProcessRequest {
	s.SignEndTime = &v
	return s
}

func (s *CreateProcessRequest) SetSourceInfo(v *CreateProcessRequestSourceInfo) *CreateProcessRequest {
	s.SourceInfo = v
	return s
}

func (s *CreateProcessRequest) SetTaskName(v string) *CreateProcessRequest {
	s.TaskName = &v
	return s
}

type CreateProcessRequestCcs struct {
	Account     *string `json:"account,omitempty" xml:"account,omitempty"`
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	OrgName     *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateProcessRequestCcs) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessRequestCcs) GoString() string {
	return s.String()
}

func (s *CreateProcessRequestCcs) SetAccount(v string) *CreateProcessRequestCcs {
	s.Account = &v
	return s
}

func (s *CreateProcessRequestCcs) SetAccountName(v string) *CreateProcessRequestCcs {
	s.AccountName = &v
	return s
}

func (s *CreateProcessRequestCcs) SetAccountType(v string) *CreateProcessRequestCcs {
	s.AccountType = &v
	return s
}

func (s *CreateProcessRequestCcs) SetOrgName(v string) *CreateProcessRequestCcs {
	s.OrgName = &v
	return s
}

func (s *CreateProcessRequestCcs) SetUserId(v string) *CreateProcessRequestCcs {
	s.UserId = &v
	return s
}

type CreateProcessRequestFiles struct {
	FileId   *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
}

func (s CreateProcessRequestFiles) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessRequestFiles) GoString() string {
	return s.String()
}

func (s *CreateProcessRequestFiles) SetFileId(v string) *CreateProcessRequestFiles {
	s.FileId = &v
	return s
}

func (s *CreateProcessRequestFiles) SetFileName(v string) *CreateProcessRequestFiles {
	s.FileName = &v
	return s
}

func (s *CreateProcessRequestFiles) SetFileType(v int32) *CreateProcessRequestFiles {
	s.FileType = &v
	return s
}

type CreateProcessRequestParticipants struct {
	Account          *string                                        `json:"account,omitempty" xml:"account,omitempty"`
	AccountName      *string                                        `json:"accountName,omitempty" xml:"accountName,omitempty"`
	AccountType      *string                                        `json:"accountType,omitempty" xml:"accountType,omitempty"`
	OrgName          *string                                        `json:"orgName,omitempty" xml:"orgName,omitempty"`
	SignOrder        *int32                                         `json:"signOrder,omitempty" xml:"signOrder,omitempty"`
	SignPosList      []*CreateProcessRequestParticipantsSignPosList `json:"signPosList,omitempty" xml:"signPosList,omitempty" type:"Repeated"`
	SignRequirements *string                                        `json:"signRequirements,omitempty" xml:"signRequirements,omitempty"`
	UserId           *string                                        `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateProcessRequestParticipants) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessRequestParticipants) GoString() string {
	return s.String()
}

func (s *CreateProcessRequestParticipants) SetAccount(v string) *CreateProcessRequestParticipants {
	s.Account = &v
	return s
}

func (s *CreateProcessRequestParticipants) SetAccountName(v string) *CreateProcessRequestParticipants {
	s.AccountName = &v
	return s
}

func (s *CreateProcessRequestParticipants) SetAccountType(v string) *CreateProcessRequestParticipants {
	s.AccountType = &v
	return s
}

func (s *CreateProcessRequestParticipants) SetOrgName(v string) *CreateProcessRequestParticipants {
	s.OrgName = &v
	return s
}

func (s *CreateProcessRequestParticipants) SetSignOrder(v int32) *CreateProcessRequestParticipants {
	s.SignOrder = &v
	return s
}

func (s *CreateProcessRequestParticipants) SetSignPosList(v []*CreateProcessRequestParticipantsSignPosList) *CreateProcessRequestParticipants {
	s.SignPosList = v
	return s
}

func (s *CreateProcessRequestParticipants) SetSignRequirements(v string) *CreateProcessRequestParticipants {
	s.SignRequirements = &v
	return s
}

func (s *CreateProcessRequestParticipants) SetUserId(v string) *CreateProcessRequestParticipants {
	s.UserId = &v
	return s
}

type CreateProcessRequestParticipantsSignPosList struct {
	FileId          *string                                              `json:"fileId,omitempty" xml:"fileId,omitempty"`
	IsCrossPage     *bool                                                `json:"isCrossPage,omitempty" xml:"isCrossPage,omitempty"`
	NeedSignDate    *bool                                                `json:"needSignDate,omitempty" xml:"needSignDate,omitempty"`
	Page            *string                                              `json:"page,omitempty" xml:"page,omitempty"`
	SignDate        *CreateProcessRequestParticipantsSignPosListSignDate `json:"signDate,omitempty" xml:"signDate,omitempty" type:"Struct"`
	SignRequirement *string                                              `json:"signRequirement,omitempty" xml:"signRequirement,omitempty"`
	X               *float64                                             `json:"x,omitempty" xml:"x,omitempty"`
	Y               *float64                                             `json:"y,omitempty" xml:"y,omitempty"`
}

func (s CreateProcessRequestParticipantsSignPosList) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessRequestParticipantsSignPosList) GoString() string {
	return s.String()
}

func (s *CreateProcessRequestParticipantsSignPosList) SetFileId(v string) *CreateProcessRequestParticipantsSignPosList {
	s.FileId = &v
	return s
}

func (s *CreateProcessRequestParticipantsSignPosList) SetIsCrossPage(v bool) *CreateProcessRequestParticipantsSignPosList {
	s.IsCrossPage = &v
	return s
}

func (s *CreateProcessRequestParticipantsSignPosList) SetNeedSignDate(v bool) *CreateProcessRequestParticipantsSignPosList {
	s.NeedSignDate = &v
	return s
}

func (s *CreateProcessRequestParticipantsSignPosList) SetPage(v string) *CreateProcessRequestParticipantsSignPosList {
	s.Page = &v
	return s
}

func (s *CreateProcessRequestParticipantsSignPosList) SetSignDate(v *CreateProcessRequestParticipantsSignPosListSignDate) *CreateProcessRequestParticipantsSignPosList {
	s.SignDate = v
	return s
}

func (s *CreateProcessRequestParticipantsSignPosList) SetSignRequirement(v string) *CreateProcessRequestParticipantsSignPosList {
	s.SignRequirement = &v
	return s
}

func (s *CreateProcessRequestParticipantsSignPosList) SetX(v float64) *CreateProcessRequestParticipantsSignPosList {
	s.X = &v
	return s
}

func (s *CreateProcessRequestParticipantsSignPosList) SetY(v float64) *CreateProcessRequestParticipantsSignPosList {
	s.Y = &v
	return s
}

type CreateProcessRequestParticipantsSignPosListSignDate struct {
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
}

func (s CreateProcessRequestParticipantsSignPosListSignDate) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessRequestParticipantsSignPosListSignDate) GoString() string {
	return s.String()
}

func (s *CreateProcessRequestParticipantsSignPosListSignDate) SetFormat(v string) *CreateProcessRequestParticipantsSignPosListSignDate {
	s.Format = &v
	return s
}

type CreateProcessRequestSourceInfo struct {
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	PcUrl     *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	ShowText  *string `json:"showText,omitempty" xml:"showText,omitempty"`
}

func (s CreateProcessRequestSourceInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessRequestSourceInfo) GoString() string {
	return s.String()
}

func (s *CreateProcessRequestSourceInfo) SetMobileUrl(v string) *CreateProcessRequestSourceInfo {
	s.MobileUrl = &v
	return s
}

func (s *CreateProcessRequestSourceInfo) SetPcUrl(v string) *CreateProcessRequestSourceInfo {
	s.PcUrl = &v
	return s
}

func (s *CreateProcessRequestSourceInfo) SetShowText(v string) *CreateProcessRequestSourceInfo {
	s.ShowText = &v
	return s
}

type CreateProcessResponseBody struct {
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProcessResponseBody) SetTaskId(v string) *CreateProcessResponseBody {
	s.TaskId = &v
	return s
}

type CreateProcessResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateProcessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessResponse) GoString() string {
	return s.String()
}

func (s *CreateProcessResponse) SetHeaders(v map[string]*string) *CreateProcessResponse {
	s.Headers = v
	return s
}

func (s *CreateProcessResponse) SetStatusCode(v int32) *CreateProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProcessResponse) SetBody(v *CreateProcessResponseBody) *CreateProcessResponse {
	s.Body = v
	return s
}

type GetAttachsApprovalHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ServiceGroup            *string            `json:"serviceGroup,omitempty" xml:"serviceGroup,omitempty"`
	TsignOpenAppId          *string            `json:"tsignOpenAppId,omitempty" xml:"tsignOpenAppId,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAttachsApprovalHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAttachsApprovalHeaders) GoString() string {
	return s.String()
}

func (s *GetAttachsApprovalHeaders) SetCommonHeaders(v map[string]*string) *GetAttachsApprovalHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAttachsApprovalHeaders) SetServiceGroup(v string) *GetAttachsApprovalHeaders {
	s.ServiceGroup = &v
	return s
}

func (s *GetAttachsApprovalHeaders) SetTsignOpenAppId(v string) *GetAttachsApprovalHeaders {
	s.TsignOpenAppId = &v
	return s
}

func (s *GetAttachsApprovalHeaders) SetXAcsDingtalkAccessToken(v string) *GetAttachsApprovalHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAttachsApprovalResponseBody struct {
	Data []*GetAttachsApprovalResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s GetAttachsApprovalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAttachsApprovalResponseBody) GoString() string {
	return s.String()
}

func (s *GetAttachsApprovalResponseBody) SetData(v []*GetAttachsApprovalResponseBodyData) *GetAttachsApprovalResponseBody {
	s.Data = v
	return s
}

type GetAttachsApprovalResponseBodyData struct {
	Files  []*GetAttachsApprovalResponseBodyDataFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	FlowId *string                                    `json:"flowId,omitempty" xml:"flowId,omitempty"`
	Status *string                                    `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetAttachsApprovalResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAttachsApprovalResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAttachsApprovalResponseBodyData) SetFiles(v []*GetAttachsApprovalResponseBodyDataFiles) *GetAttachsApprovalResponseBodyData {
	s.Files = v
	return s
}

func (s *GetAttachsApprovalResponseBodyData) SetFlowId(v string) *GetAttachsApprovalResponseBodyData {
	s.FlowId = &v
	return s
}

func (s *GetAttachsApprovalResponseBodyData) SetStatus(v string) *GetAttachsApprovalResponseBodyData {
	s.Status = &v
	return s
}

type GetAttachsApprovalResponseBodyDataFiles struct {
	FileName          *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	OriginalFileUrl   *string `json:"originalFileUrl,omitempty" xml:"originalFileUrl,omitempty"`
	SignFinishFileUrl *string `json:"signFinishFileUrl,omitempty" xml:"signFinishFileUrl,omitempty"`
}

func (s GetAttachsApprovalResponseBodyDataFiles) String() string {
	return tea.Prettify(s)
}

func (s GetAttachsApprovalResponseBodyDataFiles) GoString() string {
	return s.String()
}

func (s *GetAttachsApprovalResponseBodyDataFiles) SetFileName(v string) *GetAttachsApprovalResponseBodyDataFiles {
	s.FileName = &v
	return s
}

func (s *GetAttachsApprovalResponseBodyDataFiles) SetOriginalFileUrl(v string) *GetAttachsApprovalResponseBodyDataFiles {
	s.OriginalFileUrl = &v
	return s
}

func (s *GetAttachsApprovalResponseBodyDataFiles) SetSignFinishFileUrl(v string) *GetAttachsApprovalResponseBodyDataFiles {
	s.SignFinishFileUrl = &v
	return s
}

type GetAttachsApprovalResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAttachsApprovalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAttachsApprovalResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAttachsApprovalResponse) GoString() string {
	return s.String()
}

func (s *GetAttachsApprovalResponse) SetHeaders(v map[string]*string) *GetAttachsApprovalResponse {
	s.Headers = v
	return s
}

func (s *GetAttachsApprovalResponse) SetStatusCode(v int32) *GetAttachsApprovalResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAttachsApprovalResponse) SetBody(v *GetAttachsApprovalResponseBody) *GetAttachsApprovalResponse {
	s.Body = v
	return s
}

type GetAuthUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ServiceGroup            *string            `json:"serviceGroup,omitempty" xml:"serviceGroup,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAuthUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAuthUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetAuthUrlHeaders) SetCommonHeaders(v map[string]*string) *GetAuthUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAuthUrlHeaders) SetServiceGroup(v string) *GetAuthUrlHeaders {
	s.ServiceGroup = &v
	return s
}

func (s *GetAuthUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GetAuthUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAuthUrlRequest struct {
	RedirectUrl *string `json:"redirectUrl,omitempty" xml:"redirectUrl,omitempty"`
}

func (s GetAuthUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAuthUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAuthUrlRequest) SetRedirectUrl(v string) *GetAuthUrlRequest {
	s.RedirectUrl = &v
	return s
}

type GetAuthUrlResponseBody struct {
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	PcUrl     *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	TaskId    *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetAuthUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuthUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthUrlResponseBody) SetMobileUrl(v string) *GetAuthUrlResponseBody {
	s.MobileUrl = &v
	return s
}

func (s *GetAuthUrlResponseBody) SetPcUrl(v string) *GetAuthUrlResponseBody {
	s.PcUrl = &v
	return s
}

func (s *GetAuthUrlResponseBody) SetTaskId(v string) *GetAuthUrlResponseBody {
	s.TaskId = &v
	return s
}

type GetAuthUrlResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAuthUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAuthUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAuthUrlResponse) GoString() string {
	return s.String()
}

func (s *GetAuthUrlResponse) SetHeaders(v map[string]*string) *GetAuthUrlResponse {
	s.Headers = v
	return s
}

func (s *GetAuthUrlResponse) SetStatusCode(v int32) *GetAuthUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuthUrlResponse) SetBody(v *GetAuthUrlResponseBody) *GetAuthUrlResponse {
	s.Body = v
	return s
}

type GetContractMarginHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ServiceGroup            *string            `json:"serviceGroup,omitempty" xml:"serviceGroup,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetContractMarginHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetContractMarginHeaders) GoString() string {
	return s.String()
}

func (s *GetContractMarginHeaders) SetCommonHeaders(v map[string]*string) *GetContractMarginHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetContractMarginHeaders) SetServiceGroup(v string) *GetContractMarginHeaders {
	s.ServiceGroup = &v
	return s
}

func (s *GetContractMarginHeaders) SetXAcsDingtalkAccessToken(v string) *GetContractMarginHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetContractMarginResponseBody struct {
	Margin *float32 `json:"margin,omitempty" xml:"margin,omitempty"`
}

func (s GetContractMarginResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetContractMarginResponseBody) GoString() string {
	return s.String()
}

func (s *GetContractMarginResponseBody) SetMargin(v float32) *GetContractMarginResponseBody {
	s.Margin = &v
	return s
}

type GetContractMarginResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetContractMarginResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetContractMarginResponse) String() string {
	return tea.Prettify(s)
}

func (s GetContractMarginResponse) GoString() string {
	return s.String()
}

func (s *GetContractMarginResponse) SetHeaders(v map[string]*string) *GetContractMarginResponse {
	s.Headers = v
	return s
}

func (s *GetContractMarginResponse) SetStatusCode(v int32) *GetContractMarginResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContractMarginResponse) SetBody(v *GetContractMarginResponseBody) *GetContractMarginResponse {
	s.Body = v
	return s
}

type GetCorpConsoleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ServiceGroup            *string            `json:"serviceGroup,omitempty" xml:"serviceGroup,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCorpConsoleHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCorpConsoleHeaders) GoString() string {
	return s.String()
}

func (s *GetCorpConsoleHeaders) SetCommonHeaders(v map[string]*string) *GetCorpConsoleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCorpConsoleHeaders) SetServiceGroup(v string) *GetCorpConsoleHeaders {
	s.ServiceGroup = &v
	return s
}

func (s *GetCorpConsoleHeaders) SetXAcsDingtalkAccessToken(v string) *GetCorpConsoleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCorpConsoleResponseBody struct {
	OrgConsoleUrl *string `json:"orgConsoleUrl,omitempty" xml:"orgConsoleUrl,omitempty"`
}

func (s GetCorpConsoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCorpConsoleResponseBody) GoString() string {
	return s.String()
}

func (s *GetCorpConsoleResponseBody) SetOrgConsoleUrl(v string) *GetCorpConsoleResponseBody {
	s.OrgConsoleUrl = &v
	return s
}

type GetCorpConsoleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCorpConsoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCorpConsoleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCorpConsoleResponse) GoString() string {
	return s.String()
}

func (s *GetCorpConsoleResponse) SetHeaders(v map[string]*string) *GetCorpConsoleResponse {
	s.Headers = v
	return s
}

func (s *GetCorpConsoleResponse) SetStatusCode(v int32) *GetCorpConsoleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCorpConsoleResponse) SetBody(v *GetCorpConsoleResponseBody) *GetCorpConsoleResponse {
	s.Body = v
	return s
}

type GetCorpInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ServiceGroup            *string            `json:"serviceGroup,omitempty" xml:"serviceGroup,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCorpInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCorpInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetCorpInfoHeaders) SetCommonHeaders(v map[string]*string) *GetCorpInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCorpInfoHeaders) SetServiceGroup(v string) *GetCorpInfoHeaders {
	s.ServiceGroup = &v
	return s
}

func (s *GetCorpInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetCorpInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCorpInfoResponseBody struct {
	IsRealName  *string `json:"isRealName,omitempty" xml:"isRealName,omitempty"`
	OrgRealName *string `json:"orgRealName,omitempty" xml:"orgRealName,omitempty"`
}

func (s GetCorpInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCorpInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetCorpInfoResponseBody) SetIsRealName(v string) *GetCorpInfoResponseBody {
	s.IsRealName = &v
	return s
}

func (s *GetCorpInfoResponseBody) SetOrgRealName(v string) *GetCorpInfoResponseBody {
	s.OrgRealName = &v
	return s
}

type GetCorpInfoResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCorpInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCorpInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCorpInfoResponse) GoString() string {
	return s.String()
}

func (s *GetCorpInfoResponse) SetHeaders(v map[string]*string) *GetCorpInfoResponse {
	s.Headers = v
	return s
}

func (s *GetCorpInfoResponse) SetStatusCode(v int32) *GetCorpInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCorpInfoResponse) SetBody(v *GetCorpInfoResponseBody) *GetCorpInfoResponse {
	s.Body = v
	return s
}

type GetExecuteUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ServiceGroup            *string            `json:"serviceGroup,omitempty" xml:"serviceGroup,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetExecuteUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetExecuteUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetExecuteUrlHeaders) SetCommonHeaders(v map[string]*string) *GetExecuteUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetExecuteUrlHeaders) SetServiceGroup(v string) *GetExecuteUrlHeaders {
	s.ServiceGroup = &v
	return s
}

func (s *GetExecuteUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GetExecuteUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetExecuteUrlRequest struct {
	Account       *string `json:"account,omitempty" xml:"account,omitempty"`
	SignContainer *int32  `json:"signContainer,omitempty" xml:"signContainer,omitempty"`
	TaskId        *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetExecuteUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExecuteUrlRequest) GoString() string {
	return s.String()
}

func (s *GetExecuteUrlRequest) SetAccount(v string) *GetExecuteUrlRequest {
	s.Account = &v
	return s
}

func (s *GetExecuteUrlRequest) SetSignContainer(v int32) *GetExecuteUrlRequest {
	s.SignContainer = &v
	return s
}

func (s *GetExecuteUrlRequest) SetTaskId(v string) *GetExecuteUrlRequest {
	s.TaskId = &v
	return s
}

type GetExecuteUrlResponseBody struct {
	LongUrl   *string `json:"longUrl,omitempty" xml:"longUrl,omitempty"`
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	PcUrl     *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	ShortUrl  *string `json:"shortUrl,omitempty" xml:"shortUrl,omitempty"`
}

func (s GetExecuteUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExecuteUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetExecuteUrlResponseBody) SetLongUrl(v string) *GetExecuteUrlResponseBody {
	s.LongUrl = &v
	return s
}

func (s *GetExecuteUrlResponseBody) SetMobileUrl(v string) *GetExecuteUrlResponseBody {
	s.MobileUrl = &v
	return s
}

func (s *GetExecuteUrlResponseBody) SetPcUrl(v string) *GetExecuteUrlResponseBody {
	s.PcUrl = &v
	return s
}

func (s *GetExecuteUrlResponseBody) SetShortUrl(v string) *GetExecuteUrlResponseBody {
	s.ShortUrl = &v
	return s
}

type GetExecuteUrlResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetExecuteUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetExecuteUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExecuteUrlResponse) GoString() string {
	return s.String()
}

func (s *GetExecuteUrlResponse) SetHeaders(v map[string]*string) *GetExecuteUrlResponse {
	s.Headers = v
	return s
}

func (s *GetExecuteUrlResponse) SetStatusCode(v int32) *GetExecuteUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExecuteUrlResponse) SetBody(v *GetExecuteUrlResponseBody) *GetExecuteUrlResponse {
	s.Body = v
	return s
}

type GetFileInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ServiceGroup            *string            `json:"serviceGroup,omitempty" xml:"serviceGroup,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFileInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFileInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetFileInfoHeaders) SetCommonHeaders(v map[string]*string) *GetFileInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFileInfoHeaders) SetServiceGroup(v string) *GetFileInfoHeaders {
	s.ServiceGroup = &v
	return s
}

func (s *GetFileInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetFileInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFileInfoResponseBody struct {
	DownloadUrl   *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	FileId        *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	PdfTotalPages *int64  `json:"pdfTotalPages,omitempty" xml:"pdfTotalPages,omitempty"`
	Size          *int64  `json:"size,omitempty" xml:"size,omitempty"`
	Status        *int64  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetFileInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFileInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileInfoResponseBody) SetDownloadUrl(v string) *GetFileInfoResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *GetFileInfoResponseBody) SetFileId(v string) *GetFileInfoResponseBody {
	s.FileId = &v
	return s
}

func (s *GetFileInfoResponseBody) SetName(v string) *GetFileInfoResponseBody {
	s.Name = &v
	return s
}

func (s *GetFileInfoResponseBody) SetPdfTotalPages(v int64) *GetFileInfoResponseBody {
	s.PdfTotalPages = &v
	return s
}

func (s *GetFileInfoResponseBody) SetSize(v int64) *GetFileInfoResponseBody {
	s.Size = &v
	return s
}

func (s *GetFileInfoResponseBody) SetStatus(v int64) *GetFileInfoResponseBody {
	s.Status = &v
	return s
}

type GetFileInfoResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFileInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFileInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFileInfoResponse) GoString() string {
	return s.String()
}

func (s *GetFileInfoResponse) SetHeaders(v map[string]*string) *GetFileInfoResponse {
	s.Headers = v
	return s
}

func (s *GetFileInfoResponse) SetStatusCode(v int32) *GetFileInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileInfoResponse) SetBody(v *GetFileInfoResponseBody) *GetFileInfoResponse {
	s.Body = v
	return s
}

type GetFileUploadUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ServiceGroup            *string            `json:"serviceGroup,omitempty" xml:"serviceGroup,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFileUploadUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFileUploadUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetFileUploadUrlHeaders) SetCommonHeaders(v map[string]*string) *GetFileUploadUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFileUploadUrlHeaders) SetServiceGroup(v string) *GetFileUploadUrlHeaders {
	s.ServiceGroup = &v
	return s
}

func (s *GetFileUploadUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GetFileUploadUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFileUploadUrlRequest struct {
	ContentMd5  *string `json:"contentMd5,omitempty" xml:"contentMd5,omitempty"`
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	Convert2Pdf *bool   `json:"convert2Pdf,omitempty" xml:"convert2Pdf,omitempty"`
	FileName    *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileSize    *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
}

func (s GetFileUploadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFileUploadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetFileUploadUrlRequest) SetContentMd5(v string) *GetFileUploadUrlRequest {
	s.ContentMd5 = &v
	return s
}

func (s *GetFileUploadUrlRequest) SetContentType(v string) *GetFileUploadUrlRequest {
	s.ContentType = &v
	return s
}

func (s *GetFileUploadUrlRequest) SetConvert2Pdf(v bool) *GetFileUploadUrlRequest {
	s.Convert2Pdf = &v
	return s
}

func (s *GetFileUploadUrlRequest) SetFileName(v string) *GetFileUploadUrlRequest {
	s.FileName = &v
	return s
}

func (s *GetFileUploadUrlRequest) SetFileSize(v int64) *GetFileUploadUrlRequest {
	s.FileSize = &v
	return s
}

type GetFileUploadUrlResponseBody struct {
	FileId    *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	UploadUrl *string `json:"uploadUrl,omitempty" xml:"uploadUrl,omitempty"`
}

func (s GetFileUploadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFileUploadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileUploadUrlResponseBody) SetFileId(v string) *GetFileUploadUrlResponseBody {
	s.FileId = &v
	return s
}

func (s *GetFileUploadUrlResponseBody) SetUploadUrl(v string) *GetFileUploadUrlResponseBody {
	s.UploadUrl = &v
	return s
}

type GetFileUploadUrlResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFileUploadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFileUploadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFileUploadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetFileUploadUrlResponse) SetHeaders(v map[string]*string) *GetFileUploadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetFileUploadUrlResponse) SetStatusCode(v int32) *GetFileUploadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileUploadUrlResponse) SetBody(v *GetFileUploadUrlResponseBody) *GetFileUploadUrlResponse {
	s.Body = v
	return s
}

type GetFlowDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ServiceGroup            *string            `json:"serviceGroup,omitempty" xml:"serviceGroup,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFlowDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFlowDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetFlowDetailHeaders) SetCommonHeaders(v map[string]*string) *GetFlowDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFlowDetailHeaders) SetServiceGroup(v string) *GetFlowDetailHeaders {
	s.ServiceGroup = &v
	return s
}

func (s *GetFlowDetailHeaders) SetXAcsDingtalkAccessToken(v string) *GetFlowDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFlowDetailResponseBody struct {
	BusinessScene           *string                          `json:"businessScene,omitempty" xml:"businessScene,omitempty"`
	FlowStatus              *float32                         `json:"flowStatus,omitempty" xml:"flowStatus,omitempty"`
	InitiatorAuthorizedName *string                          `json:"initiatorAuthorizedName,omitempty" xml:"initiatorAuthorizedName,omitempty"`
	InitiatorName           *string                          `json:"initiatorName,omitempty" xml:"initiatorName,omitempty"`
	Logs                    []*GetFlowDetailResponseBodyLogs `json:"logs,omitempty" xml:"logs,omitempty" type:"Repeated"`
}

func (s GetFlowDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFlowDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetFlowDetailResponseBody) SetBusinessScene(v string) *GetFlowDetailResponseBody {
	s.BusinessScene = &v
	return s
}

func (s *GetFlowDetailResponseBody) SetFlowStatus(v float32) *GetFlowDetailResponseBody {
	s.FlowStatus = &v
	return s
}

func (s *GetFlowDetailResponseBody) SetInitiatorAuthorizedName(v string) *GetFlowDetailResponseBody {
	s.InitiatorAuthorizedName = &v
	return s
}

func (s *GetFlowDetailResponseBody) SetInitiatorName(v string) *GetFlowDetailResponseBody {
	s.InitiatorName = &v
	return s
}

func (s *GetFlowDetailResponseBody) SetLogs(v []*GetFlowDetailResponseBodyLogs) *GetFlowDetailResponseBody {
	s.Logs = v
	return s
}

type GetFlowDetailResponseBodyLogs struct {
	LogType             *string  `json:"logType,omitempty" xml:"logType,omitempty"`
	OperateDescription  *string  `json:"operateDescription,omitempty" xml:"operateDescription,omitempty"`
	OperateTime         *float32 `json:"operateTime,omitempty" xml:"operateTime,omitempty"`
	OperatorAccountName *string  `json:"operatorAccountName,omitempty" xml:"operatorAccountName,omitempty"`
}

func (s GetFlowDetailResponseBodyLogs) String() string {
	return tea.Prettify(s)
}

func (s GetFlowDetailResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *GetFlowDetailResponseBodyLogs) SetLogType(v string) *GetFlowDetailResponseBodyLogs {
	s.LogType = &v
	return s
}

func (s *GetFlowDetailResponseBodyLogs) SetOperateDescription(v string) *GetFlowDetailResponseBodyLogs {
	s.OperateDescription = &v
	return s
}

func (s *GetFlowDetailResponseBodyLogs) SetOperateTime(v float32) *GetFlowDetailResponseBodyLogs {
	s.OperateTime = &v
	return s
}

func (s *GetFlowDetailResponseBodyLogs) SetOperatorAccountName(v string) *GetFlowDetailResponseBodyLogs {
	s.OperatorAccountName = &v
	return s
}

type GetFlowDetailResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFlowDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFlowDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFlowDetailResponse) GoString() string {
	return s.String()
}

func (s *GetFlowDetailResponse) SetHeaders(v map[string]*string) *GetFlowDetailResponse {
	s.Headers = v
	return s
}

func (s *GetFlowDetailResponse) SetStatusCode(v int32) *GetFlowDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFlowDetailResponse) SetBody(v *GetFlowDetailResponseBody) *GetFlowDetailResponse {
	s.Body = v
	return s
}

type GetFlowDocsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ServiceGroup            *string            `json:"serviceGroup,omitempty" xml:"serviceGroup,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFlowDocsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFlowDocsHeaders) GoString() string {
	return s.String()
}

func (s *GetFlowDocsHeaders) SetCommonHeaders(v map[string]*string) *GetFlowDocsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFlowDocsHeaders) SetServiceGroup(v string) *GetFlowDocsHeaders {
	s.ServiceGroup = &v
	return s
}

func (s *GetFlowDocsHeaders) SetXAcsDingtalkAccessToken(v string) *GetFlowDocsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFlowDocsResponseBody struct {
	Data []*GetFlowDocsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s GetFlowDocsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFlowDocsResponseBody) GoString() string {
	return s.String()
}

func (s *GetFlowDocsResponseBody) SetData(v []*GetFlowDocsResponseBodyData) *GetFlowDocsResponseBody {
	s.Data = v
	return s
}

type GetFlowDocsResponseBodyData struct {
	FileId   *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileUrl  *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
}

func (s GetFlowDocsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFlowDocsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFlowDocsResponseBodyData) SetFileId(v string) *GetFlowDocsResponseBodyData {
	s.FileId = &v
	return s
}

func (s *GetFlowDocsResponseBodyData) SetFileName(v string) *GetFlowDocsResponseBodyData {
	s.FileName = &v
	return s
}

func (s *GetFlowDocsResponseBodyData) SetFileUrl(v string) *GetFlowDocsResponseBodyData {
	s.FileUrl = &v
	return s
}

type GetFlowDocsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFlowDocsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFlowDocsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFlowDocsResponse) GoString() string {
	return s.String()
}

func (s *GetFlowDocsResponse) SetHeaders(v map[string]*string) *GetFlowDocsResponse {
	s.Headers = v
	return s
}

func (s *GetFlowDocsResponse) SetStatusCode(v int32) *GetFlowDocsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFlowDocsResponse) SetBody(v *GetFlowDocsResponseBody) *GetFlowDocsResponse {
	s.Body = v
	return s
}

type GetIsvStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ServiceGroup            *string            `json:"serviceGroup,omitempty" xml:"serviceGroup,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetIsvStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetIsvStatusHeaders) GoString() string {
	return s.String()
}

func (s *GetIsvStatusHeaders) SetCommonHeaders(v map[string]*string) *GetIsvStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetIsvStatusHeaders) SetServiceGroup(v string) *GetIsvStatusHeaders {
	s.ServiceGroup = &v
	return s
}

func (s *GetIsvStatusHeaders) SetXAcsDingtalkAccessToken(v string) *GetIsvStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetIsvStatusResponseBody struct {
	AuthStatus    *string `json:"authStatus,omitempty" xml:"authStatus,omitempty"`
	InstallStatus *string `json:"installStatus,omitempty" xml:"installStatus,omitempty"`
}

func (s GetIsvStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIsvStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetIsvStatusResponseBody) SetAuthStatus(v string) *GetIsvStatusResponseBody {
	s.AuthStatus = &v
	return s
}

func (s *GetIsvStatusResponseBody) SetInstallStatus(v string) *GetIsvStatusResponseBody {
	s.InstallStatus = &v
	return s
}

type GetIsvStatusResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetIsvStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIsvStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIsvStatusResponse) GoString() string {
	return s.String()
}

func (s *GetIsvStatusResponse) SetHeaders(v map[string]*string) *GetIsvStatusResponse {
	s.Headers = v
	return s
}

func (s *GetIsvStatusResponse) SetStatusCode(v int32) *GetIsvStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIsvStatusResponse) SetBody(v *GetIsvStatusResponseBody) *GetIsvStatusResponse {
	s.Body = v
	return s
}

type GetSignDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ServiceGroup            *string            `json:"serviceGroup,omitempty" xml:"serviceGroup,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSignDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSignDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetSignDetailHeaders) SetCommonHeaders(v map[string]*string) *GetSignDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSignDetailHeaders) SetServiceGroup(v string) *GetSignDetailHeaders {
	s.ServiceGroup = &v
	return s
}

func (s *GetSignDetailHeaders) SetXAcsDingtalkAccessToken(v string) *GetSignDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSignDetailResponseBody struct {
	BusinessScene *string                             `json:"businessScene,omitempty" xml:"businessScene,omitempty"`
	FlowStatus    *float32                            `json:"flowStatus,omitempty" xml:"flowStatus,omitempty"`
	Signers       []*GetSignDetailResponseBodySigners `json:"signers,omitempty" xml:"signers,omitempty" type:"Repeated"`
}

func (s GetSignDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSignDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetSignDetailResponseBody) SetBusinessScene(v string) *GetSignDetailResponseBody {
	s.BusinessScene = &v
	return s
}

func (s *GetSignDetailResponseBody) SetFlowStatus(v float32) *GetSignDetailResponseBody {
	s.FlowStatus = &v
	return s
}

func (s *GetSignDetailResponseBody) SetSigners(v []*GetSignDetailResponseBodySigners) *GetSignDetailResponseBody {
	s.Signers = v
	return s
}

type GetSignDetailResponseBodySigners struct {
	SignStatus *float32 `json:"signStatus,omitempty" xml:"signStatus,omitempty"`
	SignerName *string  `json:"signerName,omitempty" xml:"signerName,omitempty"`
}

func (s GetSignDetailResponseBodySigners) String() string {
	return tea.Prettify(s)
}

func (s GetSignDetailResponseBodySigners) GoString() string {
	return s.String()
}

func (s *GetSignDetailResponseBodySigners) SetSignStatus(v float32) *GetSignDetailResponseBodySigners {
	s.SignStatus = &v
	return s
}

func (s *GetSignDetailResponseBodySigners) SetSignerName(v string) *GetSignDetailResponseBodySigners {
	s.SignerName = &v
	return s
}

type GetSignDetailResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSignDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSignDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSignDetailResponse) GoString() string {
	return s.String()
}

func (s *GetSignDetailResponse) SetHeaders(v map[string]*string) *GetSignDetailResponse {
	s.Headers = v
	return s
}

func (s *GetSignDetailResponse) SetStatusCode(v int32) *GetSignDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSignDetailResponse) SetBody(v *GetSignDetailResponseBody) *GetSignDetailResponse {
	s.Body = v
	return s
}

type GetUserInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ServiceGroup            *string            `json:"serviceGroup,omitempty" xml:"serviceGroup,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetUserInfoHeaders) SetCommonHeaders(v map[string]*string) *GetUserInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserInfoHeaders) SetServiceGroup(v string) *GetUserInfoHeaders {
	s.ServiceGroup = &v
	return s
}

func (s *GetUserInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserInfoResponseBody struct {
	IsRealName   *string `json:"isRealName,omitempty" xml:"isRealName,omitempty"`
	UserRealName *string `json:"userRealName,omitempty" xml:"userRealName,omitempty"`
}

func (s GetUserInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserInfoResponseBody) SetIsRealName(v string) *GetUserInfoResponseBody {
	s.IsRealName = &v
	return s
}

func (s *GetUserInfoResponseBody) SetUserRealName(v string) *GetUserInfoResponseBody {
	s.UserRealName = &v
	return s
}

type GetUserInfoResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserInfoResponse) GoString() string {
	return s.String()
}

func (s *GetUserInfoResponse) SetHeaders(v map[string]*string) *GetUserInfoResponse {
	s.Headers = v
	return s
}

func (s *GetUserInfoResponse) SetStatusCode(v int32) *GetUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserInfoResponse) SetBody(v *GetUserInfoResponseBody) *GetUserInfoResponse {
	s.Body = v
	return s
}

type ProcessStartHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ServiceGroup            *string            `json:"serviceGroup,omitempty" xml:"serviceGroup,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ProcessStartHeaders) String() string {
	return tea.Prettify(s)
}

func (s ProcessStartHeaders) GoString() string {
	return s.String()
}

func (s *ProcessStartHeaders) SetCommonHeaders(v map[string]*string) *ProcessStartHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ProcessStartHeaders) SetServiceGroup(v string) *ProcessStartHeaders {
	s.ServiceGroup = &v
	return s
}

func (s *ProcessStartHeaders) SetXAcsDingtalkAccessToken(v string) *ProcessStartHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ProcessStartRequest struct {
	AutoStart       *string                            `json:"autoStart,omitempty" xml:"autoStart,omitempty"`
	Ccs             []*ProcessStartRequestCcs          `json:"ccs,omitempty" xml:"ccs,omitempty" type:"Repeated"`
	Files           []*ProcessStartRequestFiles        `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	InitiatorUserId *string                            `json:"initiatorUserId,omitempty" xml:"initiatorUserId,omitempty"`
	Participants    []*ProcessStartRequestParticipants `json:"participants,omitempty" xml:"participants,omitempty" type:"Repeated"`
	RedirectUrl     *string                            `json:"redirectUrl,omitempty" xml:"redirectUrl,omitempty"`
	SourceInfo      *ProcessStartRequestSourceInfo     `json:"sourceInfo,omitempty" xml:"sourceInfo,omitempty" type:"Struct"`
	TaskName        *string                            `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s ProcessStartRequest) String() string {
	return tea.Prettify(s)
}

func (s ProcessStartRequest) GoString() string {
	return s.String()
}

func (s *ProcessStartRequest) SetAutoStart(v string) *ProcessStartRequest {
	s.AutoStart = &v
	return s
}

func (s *ProcessStartRequest) SetCcs(v []*ProcessStartRequestCcs) *ProcessStartRequest {
	s.Ccs = v
	return s
}

func (s *ProcessStartRequest) SetFiles(v []*ProcessStartRequestFiles) *ProcessStartRequest {
	s.Files = v
	return s
}

func (s *ProcessStartRequest) SetInitiatorUserId(v string) *ProcessStartRequest {
	s.InitiatorUserId = &v
	return s
}

func (s *ProcessStartRequest) SetParticipants(v []*ProcessStartRequestParticipants) *ProcessStartRequest {
	s.Participants = v
	return s
}

func (s *ProcessStartRequest) SetRedirectUrl(v string) *ProcessStartRequest {
	s.RedirectUrl = &v
	return s
}

func (s *ProcessStartRequest) SetSourceInfo(v *ProcessStartRequestSourceInfo) *ProcessStartRequest {
	s.SourceInfo = v
	return s
}

func (s *ProcessStartRequest) SetTaskName(v string) *ProcessStartRequest {
	s.TaskName = &v
	return s
}

type ProcessStartRequestCcs struct {
	Account     *string `json:"account,omitempty" xml:"account,omitempty"`
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	OrgName     *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ProcessStartRequestCcs) String() string {
	return tea.Prettify(s)
}

func (s ProcessStartRequestCcs) GoString() string {
	return s.String()
}

func (s *ProcessStartRequestCcs) SetAccount(v string) *ProcessStartRequestCcs {
	s.Account = &v
	return s
}

func (s *ProcessStartRequestCcs) SetAccountName(v string) *ProcessStartRequestCcs {
	s.AccountName = &v
	return s
}

func (s *ProcessStartRequestCcs) SetAccountType(v string) *ProcessStartRequestCcs {
	s.AccountType = &v
	return s
}

func (s *ProcessStartRequestCcs) SetOrgName(v string) *ProcessStartRequestCcs {
	s.OrgName = &v
	return s
}

func (s *ProcessStartRequestCcs) SetUserId(v string) *ProcessStartRequestCcs {
	s.UserId = &v
	return s
}

type ProcessStartRequestFiles struct {
	FileId   *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s ProcessStartRequestFiles) String() string {
	return tea.Prettify(s)
}

func (s ProcessStartRequestFiles) GoString() string {
	return s.String()
}

func (s *ProcessStartRequestFiles) SetFileId(v string) *ProcessStartRequestFiles {
	s.FileId = &v
	return s
}

func (s *ProcessStartRequestFiles) SetFileName(v string) *ProcessStartRequestFiles {
	s.FileName = &v
	return s
}

type ProcessStartRequestParticipants struct {
	Account          *string `json:"account,omitempty" xml:"account,omitempty"`
	AccountName      *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	AccountType      *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	OrgName          *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	SignRequirements *string `json:"signRequirements,omitempty" xml:"signRequirements,omitempty"`
	UserId           *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ProcessStartRequestParticipants) String() string {
	return tea.Prettify(s)
}

func (s ProcessStartRequestParticipants) GoString() string {
	return s.String()
}

func (s *ProcessStartRequestParticipants) SetAccount(v string) *ProcessStartRequestParticipants {
	s.Account = &v
	return s
}

func (s *ProcessStartRequestParticipants) SetAccountName(v string) *ProcessStartRequestParticipants {
	s.AccountName = &v
	return s
}

func (s *ProcessStartRequestParticipants) SetAccountType(v string) *ProcessStartRequestParticipants {
	s.AccountType = &v
	return s
}

func (s *ProcessStartRequestParticipants) SetOrgName(v string) *ProcessStartRequestParticipants {
	s.OrgName = &v
	return s
}

func (s *ProcessStartRequestParticipants) SetSignRequirements(v string) *ProcessStartRequestParticipants {
	s.SignRequirements = &v
	return s
}

func (s *ProcessStartRequestParticipants) SetUserId(v string) *ProcessStartRequestParticipants {
	s.UserId = &v
	return s
}

type ProcessStartRequestSourceInfo struct {
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	PcUrl     *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	ShowText  *string `json:"showText,omitempty" xml:"showText,omitempty"`
}

func (s ProcessStartRequestSourceInfo) String() string {
	return tea.Prettify(s)
}

func (s ProcessStartRequestSourceInfo) GoString() string {
	return s.String()
}

func (s *ProcessStartRequestSourceInfo) SetMobileUrl(v string) *ProcessStartRequestSourceInfo {
	s.MobileUrl = &v
	return s
}

func (s *ProcessStartRequestSourceInfo) SetPcUrl(v string) *ProcessStartRequestSourceInfo {
	s.PcUrl = &v
	return s
}

func (s *ProcessStartRequestSourceInfo) SetShowText(v string) *ProcessStartRequestSourceInfo {
	s.ShowText = &v
	return s
}

type ProcessStartResponseBody struct {
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	PcUrl     *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	TaskId    *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ProcessStartResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ProcessStartResponseBody) GoString() string {
	return s.String()
}

func (s *ProcessStartResponseBody) SetMobileUrl(v string) *ProcessStartResponseBody {
	s.MobileUrl = &v
	return s
}

func (s *ProcessStartResponseBody) SetPcUrl(v string) *ProcessStartResponseBody {
	s.PcUrl = &v
	return s
}

func (s *ProcessStartResponseBody) SetTaskId(v string) *ProcessStartResponseBody {
	s.TaskId = &v
	return s
}

type ProcessStartResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ProcessStartResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ProcessStartResponse) String() string {
	return tea.Prettify(s)
}

func (s ProcessStartResponse) GoString() string {
	return s.String()
}

func (s *ProcessStartResponse) SetHeaders(v map[string]*string) *ProcessStartResponse {
	s.Headers = v
	return s
}

func (s *ProcessStartResponse) SetStatusCode(v int32) *ProcessStartResponse {
	s.StatusCode = &v
	return s
}

func (s *ProcessStartResponse) SetBody(v *ProcessStartResponseBody) *ProcessStartResponse {
	s.Body = v
	return s
}

type ResaleOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ServiceGroup            *string            `json:"serviceGroup,omitempty" xml:"serviceGroup,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ResaleOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s ResaleOrderHeaders) GoString() string {
	return s.String()
}

func (s *ResaleOrderHeaders) SetCommonHeaders(v map[string]*string) *ResaleOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ResaleOrderHeaders) SetServiceGroup(v string) *ResaleOrderHeaders {
	s.ServiceGroup = &v
	return s
}

func (s *ResaleOrderHeaders) SetXAcsDingtalkAccessToken(v string) *ResaleOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ResaleOrderRequest struct {
	OrderCreateTime  *float32 `json:"orderCreateTime,omitempty" xml:"orderCreateTime,omitempty"`
	OrderId          *string  `json:"orderId,omitempty" xml:"orderId,omitempty"`
	Quantity         *float32 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	ServiceStartTime *float32 `json:"serviceStartTime,omitempty" xml:"serviceStartTime,omitempty"`
	ServiceStopTime  *float32 `json:"serviceStopTime,omitempty" xml:"serviceStopTime,omitempty"`
}

func (s ResaleOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s ResaleOrderRequest) GoString() string {
	return s.String()
}

func (s *ResaleOrderRequest) SetOrderCreateTime(v float32) *ResaleOrderRequest {
	s.OrderCreateTime = &v
	return s
}

func (s *ResaleOrderRequest) SetOrderId(v string) *ResaleOrderRequest {
	s.OrderId = &v
	return s
}

func (s *ResaleOrderRequest) SetQuantity(v float32) *ResaleOrderRequest {
	s.Quantity = &v
	return s
}

func (s *ResaleOrderRequest) SetServiceStartTime(v float32) *ResaleOrderRequest {
	s.ServiceStartTime = &v
	return s
}

func (s *ResaleOrderRequest) SetServiceStopTime(v float32) *ResaleOrderRequest {
	s.ServiceStopTime = &v
	return s
}

type ResaleOrderResponseBody struct {
	EsignOrderId *string `json:"esignOrderId,omitempty" xml:"esignOrderId,omitempty"`
}

func (s ResaleOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResaleOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ResaleOrderResponseBody) SetEsignOrderId(v string) *ResaleOrderResponseBody {
	s.EsignOrderId = &v
	return s
}

type ResaleOrderResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResaleOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResaleOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s ResaleOrderResponse) GoString() string {
	return s.String()
}

func (s *ResaleOrderResponse) SetHeaders(v map[string]*string) *ResaleOrderResponse {
	s.Headers = v
	return s
}

func (s *ResaleOrderResponse) SetStatusCode(v int32) *ResaleOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *ResaleOrderResponse) SetBody(v *ResaleOrderResponseBody) *ResaleOrderResponse {
	s.Body = v
	return s
}

type UsersRealnameHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ServiceGroup            *string            `json:"serviceGroup,omitempty" xml:"serviceGroup,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UsersRealnameHeaders) String() string {
	return tea.Prettify(s)
}

func (s UsersRealnameHeaders) GoString() string {
	return s.String()
}

func (s *UsersRealnameHeaders) SetCommonHeaders(v map[string]*string) *UsersRealnameHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UsersRealnameHeaders) SetServiceGroup(v string) *UsersRealnameHeaders {
	s.ServiceGroup = &v
	return s
}

func (s *UsersRealnameHeaders) SetXAcsDingtalkAccessToken(v string) *UsersRealnameHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UsersRealnameRequest struct {
	RedirectUrl *string `json:"redirectUrl,omitempty" xml:"redirectUrl,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UsersRealnameRequest) String() string {
	return tea.Prettify(s)
}

func (s UsersRealnameRequest) GoString() string {
	return s.String()
}

func (s *UsersRealnameRequest) SetRedirectUrl(v string) *UsersRealnameRequest {
	s.RedirectUrl = &v
	return s
}

func (s *UsersRealnameRequest) SetUserId(v string) *UsersRealnameRequest {
	s.UserId = &v
	return s
}

type UsersRealnameResponseBody struct {
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	PcUrl     *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	TaskId    *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s UsersRealnameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UsersRealnameResponseBody) GoString() string {
	return s.String()
}

func (s *UsersRealnameResponseBody) SetMobileUrl(v string) *UsersRealnameResponseBody {
	s.MobileUrl = &v
	return s
}

func (s *UsersRealnameResponseBody) SetPcUrl(v string) *UsersRealnameResponseBody {
	s.PcUrl = &v
	return s
}

func (s *UsersRealnameResponseBody) SetTaskId(v string) *UsersRealnameResponseBody {
	s.TaskId = &v
	return s
}

type UsersRealnameResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UsersRealnameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UsersRealnameResponse) String() string {
	return tea.Prettify(s)
}

func (s UsersRealnameResponse) GoString() string {
	return s.String()
}

func (s *UsersRealnameResponse) SetHeaders(v map[string]*string) *UsersRealnameResponse {
	s.Headers = v
	return s
}

func (s *UsersRealnameResponse) SetStatusCode(v int32) *UsersRealnameResponse {
	s.StatusCode = &v
	return s
}

func (s *UsersRealnameResponse) SetBody(v *UsersRealnameResponseBody) *UsersRealnameResponse {
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

func (client *Client) ApprovalListWithOptions(taskId *string, headers *ApprovalListHeaders, runtime *util.RuntimeOptions) (_result *ApprovalListResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ServiceGroup)) {
		realHeaders["serviceGroup"] = util.ToJSONString(headers.ServiceGroup)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("ApprovalList"),
		Version:     tea.String("esign_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/esign/approvals/" + tea.StringValue(taskId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ApprovalListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApprovalList(taskId *string) (_result *ApprovalListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ApprovalListHeaders{}
	_result = &ApprovalListResponse{}
	_body, _err := client.ApprovalListWithOptions(taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelCorpAuthWithOptions(request *CancelCorpAuthRequest, headers *CancelCorpAuthHeaders, runtime *util.RuntimeOptions) (_result *CancelCorpAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ServiceGroup)) {
		realHeaders["serviceGroup"] = util.ToJSONString(headers.ServiceGroup)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("CancelCorpAuth"),
		Version:     tea.String("esign_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/esign/auths/cancel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelCorpAuthResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelCorpAuth(request *CancelCorpAuthRequest) (_result *CancelCorpAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CancelCorpAuthHeaders{}
	_result = &CancelCorpAuthResponse{}
	_body, _err := client.CancelCorpAuthWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChannelOrdersWithOptions(request *ChannelOrdersRequest, headers *ChannelOrdersHeaders, runtime *util.RuntimeOptions) (_result *ChannelOrdersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ItemCode)) {
		body["itemCode"] = request.ItemCode
	}

	if !tea.BoolValue(util.IsUnset(request.ItemName)) {
		body["itemName"] = request.ItemName
	}

	if !tea.BoolValue(util.IsUnset(request.OrderCreateTime)) {
		body["orderCreateTime"] = request.OrderCreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		body["orderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.PayFee)) {
		body["payFee"] = request.PayFee
	}

	if !tea.BoolValue(util.IsUnset(request.Quantity)) {
		body["quantity"] = request.Quantity
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ServiceGroup)) {
		realHeaders["serviceGroup"] = util.ToJSONString(headers.ServiceGroup)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ChannelOrders"),
		Version:     tea.String("esign_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/esign/orders/channel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ChannelOrdersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChannelOrders(request *ChannelOrdersRequest) (_result *ChannelOrdersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ChannelOrdersHeaders{}
	_result = &ChannelOrdersResponse{}
	_body, _err := client.ChannelOrdersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CorpRealnameWithOptions(request *CorpRealnameRequest, headers *CorpRealnameHeaders, runtime *util.RuntimeOptions) (_result *CorpRealnameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RedirectUrl)) {
		body["redirectUrl"] = request.RedirectUrl
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ServiceGroup)) {
		realHeaders["serviceGroup"] = util.ToJSONString(headers.ServiceGroup)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CorpRealname"),
		Version:     tea.String("esign_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/esign/corps/realnames"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CorpRealnameResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CorpRealname(request *CorpRealnameRequest) (_result *CorpRealnameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CorpRealnameHeaders{}
	_result = &CorpRealnameResponse{}
	_body, _err := client.CorpRealnameWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDevelopersWithOptions(request *CreateDevelopersRequest, headers *CreateDevelopersHeaders, runtime *util.RuntimeOptions) (_result *CreateDevelopersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NoticeUrl)) {
		body["noticeUrl"] = request.NoticeUrl
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ServiceGroup)) {
		realHeaders["serviceGroup"] = util.ToJSONString(headers.ServiceGroup)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDevelopers"),
		Version:     tea.String("esign_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/esign/developers"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDevelopersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDevelopers(request *CreateDevelopersRequest) (_result *CreateDevelopersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateDevelopersHeaders{}
	_result = &CreateDevelopersResponse{}
	_body, _err := client.CreateDevelopersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProcessWithOptions(request *CreateProcessRequest, headers *CreateProcessHeaders, runtime *util.RuntimeOptions) (_result *CreateProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ccs)) {
		body["ccs"] = request.Ccs
	}

	if !tea.BoolValue(util.IsUnset(request.Files)) {
		body["files"] = request.Files
	}

	if !tea.BoolValue(util.IsUnset(request.InitiatorUserId)) {
		body["initiatorUserId"] = request.InitiatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Participants)) {
		body["participants"] = request.Participants
	}

	if !tea.BoolValue(util.IsUnset(request.RedirectUrl)) {
		body["redirectUrl"] = request.RedirectUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SignEndTime)) {
		body["signEndTime"] = request.SignEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.SourceInfo)) {
		body["sourceInfo"] = request.SourceInfo
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		body["taskName"] = request.TaskName
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
		Action:      tea.String("CreateProcess"),
		Version:     tea.String("esign_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/esign/process/startAtOnce"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProcessResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProcess(request *CreateProcessRequest) (_result *CreateProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateProcessHeaders{}
	_result = &CreateProcessResponse{}
	_body, _err := client.CreateProcessWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAttachsApprovalWithOptions(instanceId *string, headers *GetAttachsApprovalHeaders, runtime *util.RuntimeOptions) (_result *GetAttachsApprovalResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ServiceGroup)) {
		realHeaders["serviceGroup"] = util.ToJSONString(headers.ServiceGroup)
	}

	if !tea.BoolValue(util.IsUnset(headers.TsignOpenAppId)) {
		realHeaders["tsignOpenAppId"] = util.ToJSONString(headers.TsignOpenAppId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetAttachsApproval"),
		Version:     tea.String("esign_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/esign/dingInstances/" + tea.StringValue(instanceId) + "/attachments"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAttachsApprovalResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAttachsApproval(instanceId *string) (_result *GetAttachsApprovalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAttachsApprovalHeaders{}
	_result = &GetAttachsApprovalResponse{}
	_body, _err := client.GetAttachsApprovalWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAuthUrlWithOptions(request *GetAuthUrlRequest, headers *GetAuthUrlHeaders, runtime *util.RuntimeOptions) (_result *GetAuthUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RedirectUrl)) {
		body["redirectUrl"] = request.RedirectUrl
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ServiceGroup)) {
		realHeaders["serviceGroup"] = util.ToJSONString(headers.ServiceGroup)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAuthUrl"),
		Version:     tea.String("esign_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/esign/auths/urls"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAuthUrlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAuthUrl(request *GetAuthUrlRequest) (_result *GetAuthUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAuthUrlHeaders{}
	_result = &GetAuthUrlResponse{}
	_body, _err := client.GetAuthUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetContractMarginWithOptions(headers *GetContractMarginHeaders, runtime *util.RuntimeOptions) (_result *GetContractMarginResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ServiceGroup)) {
		realHeaders["serviceGroup"] = util.ToJSONString(headers.ServiceGroup)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetContractMargin"),
		Version:     tea.String("esign_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/esign/margins"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetContractMarginResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetContractMargin() (_result *GetContractMarginResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetContractMarginHeaders{}
	_result = &GetContractMarginResponse{}
	_body, _err := client.GetContractMarginWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCorpConsoleWithOptions(headers *GetCorpConsoleHeaders, runtime *util.RuntimeOptions) (_result *GetCorpConsoleResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ServiceGroup)) {
		realHeaders["serviceGroup"] = util.ToJSONString(headers.ServiceGroup)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetCorpConsole"),
		Version:     tea.String("esign_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/esign/corps/consoles"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCorpConsoleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCorpConsole() (_result *GetCorpConsoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCorpConsoleHeaders{}
	_result = &GetCorpConsoleResponse{}
	_body, _err := client.GetCorpConsoleWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCorpInfoWithOptions(headers *GetCorpInfoHeaders, runtime *util.RuntimeOptions) (_result *GetCorpInfoResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ServiceGroup)) {
		realHeaders["serviceGroup"] = util.ToJSONString(headers.ServiceGroup)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetCorpInfo"),
		Version:     tea.String("esign_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/esign/corps/infos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCorpInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCorpInfo() (_result *GetCorpInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCorpInfoHeaders{}
	_result = &GetCorpInfoResponse{}
	_body, _err := client.GetCorpInfoWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetExecuteUrlWithOptions(request *GetExecuteUrlRequest, headers *GetExecuteUrlHeaders, runtime *util.RuntimeOptions) (_result *GetExecuteUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		body["account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.SignContainer)) {
		body["signContainer"] = request.SignContainer
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ServiceGroup)) {
		realHeaders["serviceGroup"] = util.ToJSONString(headers.ServiceGroup)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetExecuteUrl"),
		Version:     tea.String("esign_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/esign/process/executeUrls"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExecuteUrlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetExecuteUrl(request *GetExecuteUrlRequest) (_result *GetExecuteUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetExecuteUrlHeaders{}
	_result = &GetExecuteUrlResponse{}
	_body, _err := client.GetExecuteUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFileInfoWithOptions(fileId *string, headers *GetFileInfoHeaders, runtime *util.RuntimeOptions) (_result *GetFileInfoResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ServiceGroup)) {
		realHeaders["serviceGroup"] = util.ToJSONString(headers.ServiceGroup)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetFileInfo"),
		Version:     tea.String("esign_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/esign/files/" + tea.StringValue(fileId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFileInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFileInfo(fileId *string) (_result *GetFileInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFileInfoHeaders{}
	_result = &GetFileInfoResponse{}
	_body, _err := client.GetFileInfoWithOptions(fileId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFileUploadUrlWithOptions(request *GetFileUploadUrlRequest, headers *GetFileUploadUrlHeaders, runtime *util.RuntimeOptions) (_result *GetFileUploadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContentMd5)) {
		body["contentMd5"] = request.ContentMd5
	}

	if !tea.BoolValue(util.IsUnset(request.ContentType)) {
		body["contentType"] = request.ContentType
	}

	if !tea.BoolValue(util.IsUnset(request.Convert2Pdf)) {
		body["convert2Pdf"] = request.Convert2Pdf
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileSize)) {
		body["fileSize"] = request.FileSize
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ServiceGroup)) {
		realHeaders["serviceGroup"] = util.ToJSONString(headers.ServiceGroup)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFileUploadUrl"),
		Version:     tea.String("esign_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/esign/files/uploadUrls"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFileUploadUrlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFileUploadUrl(request *GetFileUploadUrlRequest) (_result *GetFileUploadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFileUploadUrlHeaders{}
	_result = &GetFileUploadUrlResponse{}
	_body, _err := client.GetFileUploadUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFlowDetailWithOptions(taskId *string, headers *GetFlowDetailHeaders, runtime *util.RuntimeOptions) (_result *GetFlowDetailResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ServiceGroup)) {
		realHeaders["serviceGroup"] = util.ToJSONString(headers.ServiceGroup)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetFlowDetail"),
		Version:     tea.String("esign_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/esign/flowTasks/" + tea.StringValue(taskId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFlowDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFlowDetail(taskId *string) (_result *GetFlowDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFlowDetailHeaders{}
	_result = &GetFlowDetailResponse{}
	_body, _err := client.GetFlowDetailWithOptions(taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFlowDocsWithOptions(taskId *string, headers *GetFlowDocsHeaders, runtime *util.RuntimeOptions) (_result *GetFlowDocsResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ServiceGroup)) {
		realHeaders["serviceGroup"] = util.ToJSONString(headers.ServiceGroup)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetFlowDocs"),
		Version:     tea.String("esign_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/esign/flowTasks/" + tea.StringValue(taskId) + "/docs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFlowDocsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFlowDocs(taskId *string) (_result *GetFlowDocsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFlowDocsHeaders{}
	_result = &GetFlowDocsResponse{}
	_body, _err := client.GetFlowDocsWithOptions(taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIsvStatusWithOptions(headers *GetIsvStatusHeaders, runtime *util.RuntimeOptions) (_result *GetIsvStatusResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ServiceGroup)) {
		realHeaders["serviceGroup"] = util.ToJSONString(headers.ServiceGroup)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetIsvStatus"),
		Version:     tea.String("esign_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/esign/corps/appStatus"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIsvStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIsvStatus() (_result *GetIsvStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetIsvStatusHeaders{}
	_result = &GetIsvStatusResponse{}
	_body, _err := client.GetIsvStatusWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSignDetailWithOptions(taskId *string, headers *GetSignDetailHeaders, runtime *util.RuntimeOptions) (_result *GetSignDetailResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ServiceGroup)) {
		realHeaders["serviceGroup"] = util.ToJSONString(headers.ServiceGroup)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetSignDetail"),
		Version:     tea.String("esign_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/esign/signTasks/" + tea.StringValue(taskId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSignDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSignDetail(taskId *string) (_result *GetSignDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSignDetailHeaders{}
	_result = &GetSignDetailResponse{}
	_body, _err := client.GetSignDetailWithOptions(taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserInfoWithOptions(userId *string, headers *GetUserInfoHeaders, runtime *util.RuntimeOptions) (_result *GetUserInfoResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ServiceGroup)) {
		realHeaders["serviceGroup"] = util.ToJSONString(headers.ServiceGroup)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserInfo"),
		Version:     tea.String("esign_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/esign/users/" + tea.StringValue(userId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserInfo(userId *string) (_result *GetUserInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserInfoHeaders{}
	_result = &GetUserInfoResponse{}
	_body, _err := client.GetUserInfoWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ProcessStartWithOptions(request *ProcessStartRequest, headers *ProcessStartHeaders, runtime *util.RuntimeOptions) (_result *ProcessStartResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoStart)) {
		body["autoStart"] = request.AutoStart
	}

	if !tea.BoolValue(util.IsUnset(request.Ccs)) {
		body["ccs"] = request.Ccs
	}

	if !tea.BoolValue(util.IsUnset(request.Files)) {
		body["files"] = request.Files
	}

	if !tea.BoolValue(util.IsUnset(request.InitiatorUserId)) {
		body["initiatorUserId"] = request.InitiatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Participants)) {
		body["participants"] = request.Participants
	}

	if !tea.BoolValue(util.IsUnset(request.RedirectUrl)) {
		body["redirectUrl"] = request.RedirectUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SourceInfo)) {
		body["sourceInfo"] = request.SourceInfo
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		body["taskName"] = request.TaskName
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ServiceGroup)) {
		realHeaders["serviceGroup"] = util.ToJSONString(headers.ServiceGroup)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ProcessStart"),
		Version:     tea.String("esign_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/esign/processes/startUrls"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ProcessStartResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ProcessStart(request *ProcessStartRequest) (_result *ProcessStartResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ProcessStartHeaders{}
	_result = &ProcessStartResponse{}
	_body, _err := client.ProcessStartWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResaleOrderWithOptions(request *ResaleOrderRequest, headers *ResaleOrderHeaders, runtime *util.RuntimeOptions) (_result *ResaleOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderCreateTime)) {
		body["orderCreateTime"] = request.OrderCreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		body["orderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.Quantity)) {
		body["quantity"] = request.Quantity
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceStartTime)) {
		body["serviceStartTime"] = request.ServiceStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceStopTime)) {
		body["serviceStopTime"] = request.ServiceStopTime
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ServiceGroup)) {
		realHeaders["serviceGroup"] = util.ToJSONString(headers.ServiceGroup)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ResaleOrder"),
		Version:     tea.String("esign_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/esign/orders/resale"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ResaleOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResaleOrder(request *ResaleOrderRequest) (_result *ResaleOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ResaleOrderHeaders{}
	_result = &ResaleOrderResponse{}
	_body, _err := client.ResaleOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UsersRealnameWithOptions(request *UsersRealnameRequest, headers *UsersRealnameHeaders, runtime *util.RuntimeOptions) (_result *UsersRealnameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RedirectUrl)) {
		body["redirectUrl"] = request.RedirectUrl
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ServiceGroup)) {
		realHeaders["serviceGroup"] = util.ToJSONString(headers.ServiceGroup)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UsersRealname"),
		Version:     tea.String("esign_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/esign/users/realnames"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UsersRealnameResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UsersRealname(request *UsersRealnameRequest) (_result *UsersRealnameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UsersRealnameHeaders{}
	_result = &UsersRealnameResponse{}
	_body, _err := client.UsersRealnameWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
