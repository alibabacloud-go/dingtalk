// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package org_culture_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AssignOrgHoldingToEmpHoldingBatchHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AssignOrgHoldingToEmpHoldingBatchHeaders) String() string {
	return tea.Prettify(s)
}

func (s AssignOrgHoldingToEmpHoldingBatchHeaders) GoString() string {
	return s.String()
}

func (s *AssignOrgHoldingToEmpHoldingBatchHeaders) SetCommonHeaders(v map[string]*string) *AssignOrgHoldingToEmpHoldingBatchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchHeaders) SetXAcsDingtalkAccessToken(v string) *AssignOrgHoldingToEmpHoldingBatchHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AssignOrgHoldingToEmpHoldingBatchRequest struct {
	Remark               *string                                                   `json:"remark,omitempty" xml:"remark,omitempty"`
	SendOrgCultureInform *bool                                                     `json:"sendOrgCultureInform,omitempty" xml:"sendOrgCultureInform,omitempty"`
	SingleAmount         *int64                                                    `json:"singleAmount,omitempty" xml:"singleAmount,omitempty"`
	SourceUsage          *string                                                   `json:"sourceUsage,omitempty" xml:"sourceUsage,omitempty"`
	TargetUsage          *string                                                   `json:"targetUsage,omitempty" xml:"targetUsage,omitempty"`
	TargetUserList       []*AssignOrgHoldingToEmpHoldingBatchRequestTargetUserList `json:"targetUserList,omitempty" xml:"targetUserList,omitempty" type:"Repeated"`
}

func (s AssignOrgHoldingToEmpHoldingBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s AssignOrgHoldingToEmpHoldingBatchRequest) GoString() string {
	return s.String()
}

func (s *AssignOrgHoldingToEmpHoldingBatchRequest) SetRemark(v string) *AssignOrgHoldingToEmpHoldingBatchRequest {
	s.Remark = &v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchRequest) SetSendOrgCultureInform(v bool) *AssignOrgHoldingToEmpHoldingBatchRequest {
	s.SendOrgCultureInform = &v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchRequest) SetSingleAmount(v int64) *AssignOrgHoldingToEmpHoldingBatchRequest {
	s.SingleAmount = &v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchRequest) SetSourceUsage(v string) *AssignOrgHoldingToEmpHoldingBatchRequest {
	s.SourceUsage = &v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchRequest) SetTargetUsage(v string) *AssignOrgHoldingToEmpHoldingBatchRequest {
	s.TargetUsage = &v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchRequest) SetTargetUserList(v []*AssignOrgHoldingToEmpHoldingBatchRequestTargetUserList) *AssignOrgHoldingToEmpHoldingBatchRequest {
	s.TargetUserList = v
	return s
}

type AssignOrgHoldingToEmpHoldingBatchRequestTargetUserList struct {
	OutId        *string `json:"outId,omitempty" xml:"outId,omitempty"`
	TargetUserId *string `json:"targetUserId,omitempty" xml:"targetUserId,omitempty"`
}

func (s AssignOrgHoldingToEmpHoldingBatchRequestTargetUserList) String() string {
	return tea.Prettify(s)
}

func (s AssignOrgHoldingToEmpHoldingBatchRequestTargetUserList) GoString() string {
	return s.String()
}

func (s *AssignOrgHoldingToEmpHoldingBatchRequestTargetUserList) SetOutId(v string) *AssignOrgHoldingToEmpHoldingBatchRequestTargetUserList {
	s.OutId = &v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchRequestTargetUserList) SetTargetUserId(v string) *AssignOrgHoldingToEmpHoldingBatchRequestTargetUserList {
	s.TargetUserId = &v
	return s
}

type AssignOrgHoldingToEmpHoldingBatchResponseBody struct {
	Result  *AssignOrgHoldingToEmpHoldingBatchResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                                `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AssignOrgHoldingToEmpHoldingBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssignOrgHoldingToEmpHoldingBatchResponseBody) GoString() string {
	return s.String()
}

func (s *AssignOrgHoldingToEmpHoldingBatchResponseBody) SetResult(v *AssignOrgHoldingToEmpHoldingBatchResponseBodyResult) *AssignOrgHoldingToEmpHoldingBatchResponseBody {
	s.Result = v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchResponseBody) SetSuccess(v bool) *AssignOrgHoldingToEmpHoldingBatchResponseBody {
	s.Success = &v
	return s
}

type AssignOrgHoldingToEmpHoldingBatchResponseBodyResult struct {
	OpenPointInvokeResultDTOS []*AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS `json:"openPointInvokeResultDTOS,omitempty" xml:"openPointInvokeResultDTOS,omitempty" type:"Repeated"`
}

func (s AssignOrgHoldingToEmpHoldingBatchResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AssignOrgHoldingToEmpHoldingBatchResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AssignOrgHoldingToEmpHoldingBatchResponseBodyResult) SetOpenPointInvokeResultDTOS(v []*AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS) *AssignOrgHoldingToEmpHoldingBatchResponseBodyResult {
	s.OpenPointInvokeResultDTOS = v
	return s
}

type AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS struct {
	Code         *string `json:"code,omitempty" xml:"code,omitempty"`
	InvokeStatus *string `json:"invokeStatus,omitempty" xml:"invokeStatus,omitempty"`
	Msg          *string `json:"msg,omitempty" xml:"msg,omitempty"`
	OutId        *string `json:"outId,omitempty" xml:"outId,omitempty"`
	UserId       *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS) String() string {
	return tea.Prettify(s)
}

func (s AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS) GoString() string {
	return s.String()
}

func (s *AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS) SetCode(v string) *AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS {
	s.Code = &v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS) SetInvokeStatus(v string) *AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS {
	s.InvokeStatus = &v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS) SetMsg(v string) *AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS {
	s.Msg = &v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS) SetOutId(v string) *AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS {
	s.OutId = &v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS) SetUserId(v string) *AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS {
	s.UserId = &v
	return s
}

type AssignOrgHoldingToEmpHoldingBatchResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssignOrgHoldingToEmpHoldingBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssignOrgHoldingToEmpHoldingBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s AssignOrgHoldingToEmpHoldingBatchResponse) GoString() string {
	return s.String()
}

func (s *AssignOrgHoldingToEmpHoldingBatchResponse) SetHeaders(v map[string]*string) *AssignOrgHoldingToEmpHoldingBatchResponse {
	s.Headers = v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchResponse) SetStatusCode(v int32) *AssignOrgHoldingToEmpHoldingBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchResponse) SetBody(v *AssignOrgHoldingToEmpHoldingBatchResponseBody) *AssignOrgHoldingToEmpHoldingBatchResponse {
	s.Body = v
	return s
}

type ConsumeUserPointsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ConsumeUserPointsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ConsumeUserPointsHeaders) GoString() string {
	return s.String()
}

func (s *ConsumeUserPointsHeaders) SetCommonHeaders(v map[string]*string) *ConsumeUserPointsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ConsumeUserPointsHeaders) SetXAcsDingtalkAccessToken(v string) *ConsumeUserPointsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ConsumeUserPointsRequest struct {
	Amount *int64  `json:"amount,omitempty" xml:"amount,omitempty"`
	OutId  *string `json:"outId,omitempty" xml:"outId,omitempty"`
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	Usage  *string `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s ConsumeUserPointsRequest) String() string {
	return tea.Prettify(s)
}

func (s ConsumeUserPointsRequest) GoString() string {
	return s.String()
}

func (s *ConsumeUserPointsRequest) SetAmount(v int64) *ConsumeUserPointsRequest {
	s.Amount = &v
	return s
}

func (s *ConsumeUserPointsRequest) SetOutId(v string) *ConsumeUserPointsRequest {
	s.OutId = &v
	return s
}

func (s *ConsumeUserPointsRequest) SetRemark(v string) *ConsumeUserPointsRequest {
	s.Remark = &v
	return s
}

func (s *ConsumeUserPointsRequest) SetUsage(v string) *ConsumeUserPointsRequest {
	s.Usage = &v
	return s
}

type ConsumeUserPointsResponseBody struct {
	Result  *ConsumeUserPointsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ConsumeUserPointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConsumeUserPointsResponseBody) GoString() string {
	return s.String()
}

func (s *ConsumeUserPointsResponseBody) SetResult(v *ConsumeUserPointsResponseBodyResult) *ConsumeUserPointsResponseBody {
	s.Result = v
	return s
}

func (s *ConsumeUserPointsResponseBody) SetSuccess(v bool) *ConsumeUserPointsResponseBody {
	s.Success = &v
	return s
}

type ConsumeUserPointsResponseBodyResult struct {
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

func (s ConsumeUserPointsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ConsumeUserPointsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ConsumeUserPointsResponseBodyResult) SetAmount(v int64) *ConsumeUserPointsResponseBodyResult {
	s.Amount = &v
	return s
}

type ConsumeUserPointsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConsumeUserPointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConsumeUserPointsResponse) String() string {
	return tea.Prettify(s)
}

func (s ConsumeUserPointsResponse) GoString() string {
	return s.String()
}

func (s *ConsumeUserPointsResponse) SetHeaders(v map[string]*string) *ConsumeUserPointsResponse {
	s.Headers = v
	return s
}

func (s *ConsumeUserPointsResponse) SetStatusCode(v int32) *ConsumeUserPointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ConsumeUserPointsResponse) SetBody(v *ConsumeUserPointsResponseBody) *ConsumeUserPointsResponse {
	s.Body = v
	return s
}

type CreateOrgHonorHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateOrgHonorHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateOrgHonorHeaders) GoString() string {
	return s.String()
}

func (s *CreateOrgHonorHeaders) SetCommonHeaders(v map[string]*string) *CreateOrgHonorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateOrgHonorHeaders) SetXAcsDingtalkAccessToken(v string) *CreateOrgHonorHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateOrgHonorRequest struct {
	AvatarFrameMediaId *string `json:"avatarFrameMediaId,omitempty" xml:"avatarFrameMediaId,omitempty"`
	DefaultBgColor     *string `json:"defaultBgColor,omitempty" xml:"defaultBgColor,omitempty"`
	MedalDesc          *string `json:"medalDesc,omitempty" xml:"medalDesc,omitempty"`
	MedalMediaId       *string `json:"medalMediaId,omitempty" xml:"medalMediaId,omitempty"`
	MedalName          *string `json:"medalName,omitempty" xml:"medalName,omitempty"`
	UserId             *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateOrgHonorRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrgHonorRequest) GoString() string {
	return s.String()
}

func (s *CreateOrgHonorRequest) SetAvatarFrameMediaId(v string) *CreateOrgHonorRequest {
	s.AvatarFrameMediaId = &v
	return s
}

func (s *CreateOrgHonorRequest) SetDefaultBgColor(v string) *CreateOrgHonorRequest {
	s.DefaultBgColor = &v
	return s
}

func (s *CreateOrgHonorRequest) SetMedalDesc(v string) *CreateOrgHonorRequest {
	s.MedalDesc = &v
	return s
}

func (s *CreateOrgHonorRequest) SetMedalMediaId(v string) *CreateOrgHonorRequest {
	s.MedalMediaId = &v
	return s
}

func (s *CreateOrgHonorRequest) SetMedalName(v string) *CreateOrgHonorRequest {
	s.MedalName = &v
	return s
}

func (s *CreateOrgHonorRequest) SetUserId(v string) *CreateOrgHonorRequest {
	s.UserId = &v
	return s
}

type CreateOrgHonorResponseBody struct {
	Result  *CreateOrgHonorResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateOrgHonorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrgHonorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrgHonorResponseBody) SetResult(v *CreateOrgHonorResponseBodyResult) *CreateOrgHonorResponseBody {
	s.Result = v
	return s
}

func (s *CreateOrgHonorResponseBody) SetSuccess(v bool) *CreateOrgHonorResponseBody {
	s.Success = &v
	return s
}

type CreateOrgHonorResponseBodyResult struct {
	HonorId *string `json:"honorId,omitempty" xml:"honorId,omitempty"`
}

func (s CreateOrgHonorResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateOrgHonorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateOrgHonorResponseBodyResult) SetHonorId(v string) *CreateOrgHonorResponseBodyResult {
	s.HonorId = &v
	return s
}

type CreateOrgHonorResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrgHonorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrgHonorResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrgHonorResponse) GoString() string {
	return s.String()
}

func (s *CreateOrgHonorResponse) SetHeaders(v map[string]*string) *CreateOrgHonorResponse {
	s.Headers = v
	return s
}

func (s *CreateOrgHonorResponse) SetStatusCode(v int32) *CreateOrgHonorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrgHonorResponse) SetBody(v *CreateOrgHonorResponseBody) *CreateOrgHonorResponse {
	s.Body = v
	return s
}

type DeductionPointBatchHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeductionPointBatchHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeductionPointBatchHeaders) GoString() string {
	return s.String()
}

func (s *DeductionPointBatchHeaders) SetCommonHeaders(v map[string]*string) *DeductionPointBatchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeductionPointBatchHeaders) SetXAcsDingtalkAccessToken(v string) *DeductionPointBatchHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeductionPointBatchRequest struct {
	DeductionAmount      *int64                                      `json:"deductionAmount,omitempty" xml:"deductionAmount,omitempty"`
	Remark               *string                                     `json:"remark,omitempty" xml:"remark,omitempty"`
	SendOrgCultureInform *bool                                       `json:"sendOrgCultureInform,omitempty" xml:"sendOrgCultureInform,omitempty"`
	TargetUserList       []*DeductionPointBatchRequestTargetUserList `json:"targetUserList,omitempty" xml:"targetUserList,omitempty" type:"Repeated"`
	UserId               *string                                     `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeductionPointBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s DeductionPointBatchRequest) GoString() string {
	return s.String()
}

func (s *DeductionPointBatchRequest) SetDeductionAmount(v int64) *DeductionPointBatchRequest {
	s.DeductionAmount = &v
	return s
}

func (s *DeductionPointBatchRequest) SetRemark(v string) *DeductionPointBatchRequest {
	s.Remark = &v
	return s
}

func (s *DeductionPointBatchRequest) SetSendOrgCultureInform(v bool) *DeductionPointBatchRequest {
	s.SendOrgCultureInform = &v
	return s
}

func (s *DeductionPointBatchRequest) SetTargetUserList(v []*DeductionPointBatchRequestTargetUserList) *DeductionPointBatchRequest {
	s.TargetUserList = v
	return s
}

func (s *DeductionPointBatchRequest) SetUserId(v string) *DeductionPointBatchRequest {
	s.UserId = &v
	return s
}

type DeductionPointBatchRequestTargetUserList struct {
	OutId        *string `json:"outId,omitempty" xml:"outId,omitempty"`
	TargetUserId *string `json:"targetUserId,omitempty" xml:"targetUserId,omitempty"`
}

func (s DeductionPointBatchRequestTargetUserList) String() string {
	return tea.Prettify(s)
}

func (s DeductionPointBatchRequestTargetUserList) GoString() string {
	return s.String()
}

func (s *DeductionPointBatchRequestTargetUserList) SetOutId(v string) *DeductionPointBatchRequestTargetUserList {
	s.OutId = &v
	return s
}

func (s *DeductionPointBatchRequestTargetUserList) SetTargetUserId(v string) *DeductionPointBatchRequestTargetUserList {
	s.TargetUserId = &v
	return s
}

type DeductionPointBatchResponseBody struct {
	Result  *DeductionPointBatchResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeductionPointBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeductionPointBatchResponseBody) GoString() string {
	return s.String()
}

func (s *DeductionPointBatchResponseBody) SetResult(v *DeductionPointBatchResponseBodyResult) *DeductionPointBatchResponseBody {
	s.Result = v
	return s
}

func (s *DeductionPointBatchResponseBody) SetSuccess(v bool) *DeductionPointBatchResponseBody {
	s.Success = &v
	return s
}

type DeductionPointBatchResponseBodyResult struct {
	OpenPointInvokeResultDTOS []*DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS `json:"openPointInvokeResultDTOS,omitempty" xml:"openPointInvokeResultDTOS,omitempty" type:"Repeated"`
}

func (s DeductionPointBatchResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeductionPointBatchResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeductionPointBatchResponseBodyResult) SetOpenPointInvokeResultDTOS(v []*DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS) *DeductionPointBatchResponseBodyResult {
	s.OpenPointInvokeResultDTOS = v
	return s
}

type DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS struct {
	Code         *string `json:"code,omitempty" xml:"code,omitempty"`
	InvokeStatus *string `json:"invokeStatus,omitempty" xml:"invokeStatus,omitempty"`
	Msg          *string `json:"msg,omitempty" xml:"msg,omitempty"`
	OutId        *string `json:"outId,omitempty" xml:"outId,omitempty"`
	UserId       *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS) String() string {
	return tea.Prettify(s)
}

func (s DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS) GoString() string {
	return s.String()
}

func (s *DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS) SetCode(v string) *DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS {
	s.Code = &v
	return s
}

func (s *DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS) SetInvokeStatus(v string) *DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS {
	s.InvokeStatus = &v
	return s
}

func (s *DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS) SetMsg(v string) *DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS {
	s.Msg = &v
	return s
}

func (s *DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS) SetOutId(v string) *DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS {
	s.OutId = &v
	return s
}

func (s *DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS) SetUserId(v string) *DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS {
	s.UserId = &v
	return s
}

type DeductionPointBatchResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeductionPointBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeductionPointBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s DeductionPointBatchResponse) GoString() string {
	return s.String()
}

func (s *DeductionPointBatchResponse) SetHeaders(v map[string]*string) *DeductionPointBatchResponse {
	s.Headers = v
	return s
}

func (s *DeductionPointBatchResponse) SetStatusCode(v int32) *DeductionPointBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *DeductionPointBatchResponse) SetBody(v *DeductionPointBatchResponseBody) *DeductionPointBatchResponse {
	s.Body = v
	return s
}

type ExportPointOpenHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ExportPointOpenHeaders) String() string {
	return tea.Prettify(s)
}

func (s ExportPointOpenHeaders) GoString() string {
	return s.String()
}

func (s *ExportPointOpenHeaders) SetCommonHeaders(v map[string]*string) *ExportPointOpenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ExportPointOpenHeaders) SetXAcsDingtalkAccessToken(v string) *ExportPointOpenHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ExportPointOpenRequest struct {
	ExportDate *string `json:"exportDate,omitempty" xml:"exportDate,omitempty"`
	ExportType *int64  `json:"exportType,omitempty" xml:"exportType,omitempty"`
	UserId     *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExportPointOpenRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportPointOpenRequest) GoString() string {
	return s.String()
}

func (s *ExportPointOpenRequest) SetExportDate(v string) *ExportPointOpenRequest {
	s.ExportDate = &v
	return s
}

func (s *ExportPointOpenRequest) SetExportType(v int64) *ExportPointOpenRequest {
	s.ExportType = &v
	return s
}

func (s *ExportPointOpenRequest) SetUserId(v string) *ExportPointOpenRequest {
	s.UserId = &v
	return s
}

type ExportPointOpenResponseBody struct {
	Result  *bool `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExportPointOpenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportPointOpenResponseBody) GoString() string {
	return s.String()
}

func (s *ExportPointOpenResponseBody) SetResult(v bool) *ExportPointOpenResponseBody {
	s.Result = &v
	return s
}

func (s *ExportPointOpenResponseBody) SetSuccess(v bool) *ExportPointOpenResponseBody {
	s.Success = &v
	return s
}

type ExportPointOpenResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportPointOpenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportPointOpenResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportPointOpenResponse) GoString() string {
	return s.String()
}

func (s *ExportPointOpenResponse) SetHeaders(v map[string]*string) *ExportPointOpenResponse {
	s.Headers = v
	return s
}

func (s *ExportPointOpenResponse) SetStatusCode(v int32) *ExportPointOpenResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportPointOpenResponse) SetBody(v *ExportPointOpenResponseBody) *ExportPointOpenResponse {
	s.Body = v
	return s
}

type GrantHonorHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GrantHonorHeaders) String() string {
	return tea.Prettify(s)
}

func (s GrantHonorHeaders) GoString() string {
	return s.String()
}

func (s *GrantHonorHeaders) SetCommonHeaders(v map[string]*string) *GrantHonorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GrantHonorHeaders) SetXAcsDingtalkAccessToken(v string) *GrantHonorHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GrantHonorRequest struct {
	ExpirationTime      *int64    `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
	GrantReason         *string   `json:"grantReason,omitempty" xml:"grantReason,omitempty"`
	GranterName         *string   `json:"granterName,omitempty" xml:"granterName,omitempty"`
	NoticeAnnouncer     *bool     `json:"noticeAnnouncer,omitempty" xml:"noticeAnnouncer,omitempty"`
	NoticeSingle        *bool     `json:"noticeSingle,omitempty" xml:"noticeSingle,omitempty"`
	OpenConversationIds []*string `json:"openConversationIds,omitempty" xml:"openConversationIds,omitempty" type:"Repeated"`
	ReceiverUserIds     []*string `json:"receiverUserIds,omitempty" xml:"receiverUserIds,omitempty" type:"Repeated"`
	SenderUserId        *string   `json:"senderUserId,omitempty" xml:"senderUserId,omitempty"`
}

func (s GrantHonorRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantHonorRequest) GoString() string {
	return s.String()
}

func (s *GrantHonorRequest) SetExpirationTime(v int64) *GrantHonorRequest {
	s.ExpirationTime = &v
	return s
}

func (s *GrantHonorRequest) SetGrantReason(v string) *GrantHonorRequest {
	s.GrantReason = &v
	return s
}

func (s *GrantHonorRequest) SetGranterName(v string) *GrantHonorRequest {
	s.GranterName = &v
	return s
}

func (s *GrantHonorRequest) SetNoticeAnnouncer(v bool) *GrantHonorRequest {
	s.NoticeAnnouncer = &v
	return s
}

func (s *GrantHonorRequest) SetNoticeSingle(v bool) *GrantHonorRequest {
	s.NoticeSingle = &v
	return s
}

func (s *GrantHonorRequest) SetOpenConversationIds(v []*string) *GrantHonorRequest {
	s.OpenConversationIds = v
	return s
}

func (s *GrantHonorRequest) SetReceiverUserIds(v []*string) *GrantHonorRequest {
	s.ReceiverUserIds = v
	return s
}

func (s *GrantHonorRequest) SetSenderUserId(v string) *GrantHonorRequest {
	s.SenderUserId = &v
	return s
}

type GrantHonorResponseBody struct {
	Result  *GrantHonorResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GrantHonorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantHonorResponseBody) GoString() string {
	return s.String()
}

func (s *GrantHonorResponseBody) SetResult(v *GrantHonorResponseBodyResult) *GrantHonorResponseBody {
	s.Result = v
	return s
}

func (s *GrantHonorResponseBody) SetSuccess(v bool) *GrantHonorResponseBody {
	s.Success = &v
	return s
}

type GrantHonorResponseBodyResult struct {
	FailedUserIds  []*string `json:"failedUserIds,omitempty" xml:"failedUserIds,omitempty" type:"Repeated"`
	SuccessUserIds []*string `json:"successUserIds,omitempty" xml:"successUserIds,omitempty" type:"Repeated"`
}

func (s GrantHonorResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GrantHonorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GrantHonorResponseBodyResult) SetFailedUserIds(v []*string) *GrantHonorResponseBodyResult {
	s.FailedUserIds = v
	return s
}

func (s *GrantHonorResponseBodyResult) SetSuccessUserIds(v []*string) *GrantHonorResponseBodyResult {
	s.SuccessUserIds = v
	return s
}

type GrantHonorResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantHonorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantHonorResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantHonorResponse) GoString() string {
	return s.String()
}

func (s *GrantHonorResponse) SetHeaders(v map[string]*string) *GrantHonorResponse {
	s.Headers = v
	return s
}

func (s *GrantHonorResponse) SetStatusCode(v int32) *GrantHonorResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantHonorResponse) SetBody(v *GrantHonorResponseBody) *GrantHonorResponse {
	s.Body = v
	return s
}

type QueryCorpPointsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCorpPointsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCorpPointsHeaders) GoString() string {
	return s.String()
}

func (s *QueryCorpPointsHeaders) SetCommonHeaders(v map[string]*string) *QueryCorpPointsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCorpPointsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCorpPointsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCorpPointsRequest struct {
	OptUserId *string `json:"optUserId,omitempty" xml:"optUserId,omitempty"`
}

func (s QueryCorpPointsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCorpPointsRequest) GoString() string {
	return s.String()
}

func (s *QueryCorpPointsRequest) SetOptUserId(v string) *QueryCorpPointsRequest {
	s.OptUserId = &v
	return s
}

type QueryCorpPointsResponseBody struct {
	Result  *QueryCorpPointsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryCorpPointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCorpPointsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCorpPointsResponseBody) SetResult(v *QueryCorpPointsResponseBodyResult) *QueryCorpPointsResponseBody {
	s.Result = v
	return s
}

func (s *QueryCorpPointsResponseBody) SetSuccess(v bool) *QueryCorpPointsResponseBody {
	s.Success = &v
	return s
}

type QueryCorpPointsResponseBodyResult struct {
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

func (s QueryCorpPointsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryCorpPointsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryCorpPointsResponseBodyResult) SetAmount(v int64) *QueryCorpPointsResponseBodyResult {
	s.Amount = &v
	return s
}

type QueryCorpPointsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCorpPointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCorpPointsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCorpPointsResponse) GoString() string {
	return s.String()
}

func (s *QueryCorpPointsResponse) SetHeaders(v map[string]*string) *QueryCorpPointsResponse {
	s.Headers = v
	return s
}

func (s *QueryCorpPointsResponse) SetStatusCode(v int32) *QueryCorpPointsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCorpPointsResponse) SetBody(v *QueryCorpPointsResponseBody) *QueryCorpPointsResponse {
	s.Body = v
	return s
}

type QueryEmpPointDetailsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryEmpPointDetailsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryEmpPointDetailsHeaders) GoString() string {
	return s.String()
}

func (s *QueryEmpPointDetailsHeaders) SetCommonHeaders(v map[string]*string) *QueryEmpPointDetailsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryEmpPointDetailsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryEmpPointDetailsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryEmpPointDetailsRequest struct {
	PageNumber *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	UserId     *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryEmpPointDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEmpPointDetailsRequest) GoString() string {
	return s.String()
}

func (s *QueryEmpPointDetailsRequest) SetPageNumber(v int64) *QueryEmpPointDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryEmpPointDetailsRequest) SetPageSize(v int64) *QueryEmpPointDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryEmpPointDetailsRequest) SetUserId(v string) *QueryEmpPointDetailsRequest {
	s.UserId = &v
	return s
}

type QueryEmpPointDetailsResponseBody struct {
	Result  *QueryEmpPointDetailsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryEmpPointDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEmpPointDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEmpPointDetailsResponseBody) SetResult(v *QueryEmpPointDetailsResponseBodyResult) *QueryEmpPointDetailsResponseBody {
	s.Result = v
	return s
}

func (s *QueryEmpPointDetailsResponseBody) SetSuccess(v bool) *QueryEmpPointDetailsResponseBody {
	s.Success = &v
	return s
}

type QueryEmpPointDetailsResponseBodyResult struct {
	Details []*QueryEmpPointDetailsResponseBodyResultDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
	HasMore *bool                                            `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
}

func (s QueryEmpPointDetailsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryEmpPointDetailsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryEmpPointDetailsResponseBodyResult) SetDetails(v []*QueryEmpPointDetailsResponseBodyResultDetails) *QueryEmpPointDetailsResponseBodyResult {
	s.Details = v
	return s
}

func (s *QueryEmpPointDetailsResponseBodyResult) SetHasMore(v bool) *QueryEmpPointDetailsResponseBodyResult {
	s.HasMore = &v
	return s
}

type QueryEmpPointDetailsResponseBodyResultDetails struct {
	Amount                         *int64                                                                       `json:"amount,omitempty" xml:"amount,omitempty"`
	GmtCreate                      *int64                                                                       `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	OutId                          *string                                                                      `json:"outId,omitempty" xml:"outId,omitempty"`
	PointOperateFeatureResponseDTO *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO `json:"pointOperateFeatureResponseDTO,omitempty" xml:"pointOperateFeatureResponseDTO,omitempty" type:"Struct"`
	SourceBizCode                  *string                                                                      `json:"sourceBizCode,omitempty" xml:"sourceBizCode,omitempty"`
}

func (s QueryEmpPointDetailsResponseBodyResultDetails) String() string {
	return tea.Prettify(s)
}

func (s QueryEmpPointDetailsResponseBodyResultDetails) GoString() string {
	return s.String()
}

func (s *QueryEmpPointDetailsResponseBodyResultDetails) SetAmount(v int64) *QueryEmpPointDetailsResponseBodyResultDetails {
	s.Amount = &v
	return s
}

func (s *QueryEmpPointDetailsResponseBodyResultDetails) SetGmtCreate(v int64) *QueryEmpPointDetailsResponseBodyResultDetails {
	s.GmtCreate = &v
	return s
}

func (s *QueryEmpPointDetailsResponseBodyResultDetails) SetOutId(v string) *QueryEmpPointDetailsResponseBodyResultDetails {
	s.OutId = &v
	return s
}

func (s *QueryEmpPointDetailsResponseBodyResultDetails) SetPointOperateFeatureResponseDTO(v *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) *QueryEmpPointDetailsResponseBodyResultDetails {
	s.PointOperateFeatureResponseDTO = v
	return s
}

func (s *QueryEmpPointDetailsResponseBodyResultDetails) SetSourceBizCode(v string) *QueryEmpPointDetailsResponseBodyResultDetails {
	s.SourceBizCode = &v
	return s
}

type QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO struct {
	AccountSource *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource `json:"accountSource,omitempty" xml:"accountSource,omitempty" type:"Struct"`
	AccountTarget *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget `json:"accountTarget,omitempty" xml:"accountTarget,omitempty" type:"Struct"`
	Remark        *string                                                                                   `json:"remark,omitempty" xml:"remark,omitempty"`
	Usage         *string                                                                                   `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) String() string {
	return tea.Prettify(s)
}

func (s QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) GoString() string {
	return s.String()
}

func (s *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) SetAccountSource(v *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource) *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO {
	s.AccountSource = v
	return s
}

func (s *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) SetAccountTarget(v *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget) *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO {
	s.AccountTarget = v
	return s
}

func (s *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) SetRemark(v string) *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO {
	s.Remark = &v
	return s
}

func (s *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) SetUsage(v string) *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO {
	s.Usage = &v
	return s
}

type QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource struct {
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	EmpName     *string `json:"empName,omitempty" xml:"empName,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource) String() string {
	return tea.Prettify(s)
}

func (s QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource) GoString() string {
	return s.String()
}

func (s *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource) SetAccountType(v string) *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource {
	s.AccountType = &v
	return s
}

func (s *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource) SetEmpName(v string) *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource {
	s.EmpName = &v
	return s
}

func (s *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource) SetUserId(v string) *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource {
	s.UserId = &v
	return s
}

type QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget struct {
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	EmpName     *string `json:"empName,omitempty" xml:"empName,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget) String() string {
	return tea.Prettify(s)
}

func (s QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget) GoString() string {
	return s.String()
}

func (s *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget) SetAccountType(v string) *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget {
	s.AccountType = &v
	return s
}

func (s *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget) SetEmpName(v string) *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget {
	s.EmpName = &v
	return s
}

func (s *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget) SetUserId(v string) *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget {
	s.UserId = &v
	return s
}

type QueryEmpPointDetailsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEmpPointDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEmpPointDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEmpPointDetailsResponse) GoString() string {
	return s.String()
}

func (s *QueryEmpPointDetailsResponse) SetHeaders(v map[string]*string) *QueryEmpPointDetailsResponse {
	s.Headers = v
	return s
}

func (s *QueryEmpPointDetailsResponse) SetStatusCode(v int32) *QueryEmpPointDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEmpPointDetailsResponse) SetBody(v *QueryEmpPointDetailsResponseBody) *QueryEmpPointDetailsResponse {
	s.Body = v
	return s
}

type QueryOrgHonorsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOrgHonorsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgHonorsHeaders) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsHeaders) SetCommonHeaders(v map[string]*string) *QueryOrgHonorsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOrgHonorsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOrgHonorsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOrgHonorsRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryOrgHonorsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgHonorsRequest) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsRequest) SetMaxResults(v int32) *QueryOrgHonorsRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryOrgHonorsRequest) SetNextToken(v string) *QueryOrgHonorsRequest {
	s.NextToken = &v
	return s
}

type QueryOrgHonorsResponseBody struct {
	Result  *QueryOrgHonorsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryOrgHonorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgHonorsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsResponseBody) SetResult(v *QueryOrgHonorsResponseBodyResult) *QueryOrgHonorsResponseBody {
	s.Result = v
	return s
}

func (s *QueryOrgHonorsResponseBody) SetSuccess(v bool) *QueryOrgHonorsResponseBody {
	s.Success = &v
	return s
}

type QueryOrgHonorsResponseBodyResult struct {
	NextToken  *string                                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OpenHonors []*QueryOrgHonorsResponseBodyResultOpenHonors `json:"openHonors,omitempty" xml:"openHonors,omitempty" type:"Repeated"`
}

func (s QueryOrgHonorsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgHonorsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsResponseBodyResult) SetNextToken(v string) *QueryOrgHonorsResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *QueryOrgHonorsResponseBodyResult) SetOpenHonors(v []*QueryOrgHonorsResponseBodyResultOpenHonors) *QueryOrgHonorsResponseBodyResult {
	s.OpenHonors = v
	return s
}

type QueryOrgHonorsResponseBodyResultOpenHonors struct {
	HonorDesc          *string `json:"honorDesc,omitempty" xml:"honorDesc,omitempty"`
	HonorId            *int64  `json:"honorId,omitempty" xml:"honorId,omitempty"`
	HonorImgUrl        *string `json:"honorImgUrl,omitempty" xml:"honorImgUrl,omitempty"`
	HonorName          *string `json:"honorName,omitempty" xml:"honorName,omitempty"`
	HonorPendantImgUrl *string `json:"honorPendantImgUrl,omitempty" xml:"honorPendantImgUrl,omitempty"`
}

func (s QueryOrgHonorsResponseBodyResultOpenHonors) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgHonorsResponseBodyResultOpenHonors) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsResponseBodyResultOpenHonors) SetHonorDesc(v string) *QueryOrgHonorsResponseBodyResultOpenHonors {
	s.HonorDesc = &v
	return s
}

func (s *QueryOrgHonorsResponseBodyResultOpenHonors) SetHonorId(v int64) *QueryOrgHonorsResponseBodyResultOpenHonors {
	s.HonorId = &v
	return s
}

func (s *QueryOrgHonorsResponseBodyResultOpenHonors) SetHonorImgUrl(v string) *QueryOrgHonorsResponseBodyResultOpenHonors {
	s.HonorImgUrl = &v
	return s
}

func (s *QueryOrgHonorsResponseBodyResultOpenHonors) SetHonorName(v string) *QueryOrgHonorsResponseBodyResultOpenHonors {
	s.HonorName = &v
	return s
}

func (s *QueryOrgHonorsResponseBodyResultOpenHonors) SetHonorPendantImgUrl(v string) *QueryOrgHonorsResponseBodyResultOpenHonors {
	s.HonorPendantImgUrl = &v
	return s
}

type QueryOrgHonorsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOrgHonorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOrgHonorsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgHonorsResponse) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsResponse) SetHeaders(v map[string]*string) *QueryOrgHonorsResponse {
	s.Headers = v
	return s
}

func (s *QueryOrgHonorsResponse) SetStatusCode(v int32) *QueryOrgHonorsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrgHonorsResponse) SetBody(v *QueryOrgHonorsResponseBody) *QueryOrgHonorsResponse {
	s.Body = v
	return s
}

type QueryOrgPointDetailsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOrgPointDetailsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgPointDetailsHeaders) GoString() string {
	return s.String()
}

func (s *QueryOrgPointDetailsHeaders) SetCommonHeaders(v map[string]*string) *QueryOrgPointDetailsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOrgPointDetailsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOrgPointDetailsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOrgPointDetailsRequest struct {
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	PageNumber  *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize    *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryOrgPointDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgPointDetailsRequest) GoString() string {
	return s.String()
}

func (s *QueryOrgPointDetailsRequest) SetAccountType(v string) *QueryOrgPointDetailsRequest {
	s.AccountType = &v
	return s
}

func (s *QueryOrgPointDetailsRequest) SetPageNumber(v int64) *QueryOrgPointDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryOrgPointDetailsRequest) SetPageSize(v int64) *QueryOrgPointDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryOrgPointDetailsRequest) SetUserId(v string) *QueryOrgPointDetailsRequest {
	s.UserId = &v
	return s
}

type QueryOrgPointDetailsResponseBody struct {
	Result *QueryOrgPointDetailsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryOrgPointDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgPointDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrgPointDetailsResponseBody) SetResult(v *QueryOrgPointDetailsResponseBodyResult) *QueryOrgPointDetailsResponseBody {
	s.Result = v
	return s
}

type QueryOrgPointDetailsResponseBodyResult struct {
	Details []*QueryOrgPointDetailsResponseBodyResultDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
	HasMore *bool                                            `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	Success *bool                                            `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryOrgPointDetailsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgPointDetailsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryOrgPointDetailsResponseBodyResult) SetDetails(v []*QueryOrgPointDetailsResponseBodyResultDetails) *QueryOrgPointDetailsResponseBodyResult {
	s.Details = v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResult) SetHasMore(v bool) *QueryOrgPointDetailsResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResult) SetSuccess(v bool) *QueryOrgPointDetailsResponseBodyResult {
	s.Success = &v
	return s
}

type QueryOrgPointDetailsResponseBodyResultDetails struct {
	Amount                         *int64                                                                       `json:"amount,omitempty" xml:"amount,omitempty"`
	GmtCreate                      *int64                                                                       `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	OutId                          *string                                                                      `json:"outId,omitempty" xml:"outId,omitempty"`
	PointOperateFeatureResponseDTO *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO `json:"pointOperateFeatureResponseDTO,omitempty" xml:"pointOperateFeatureResponseDTO,omitempty" type:"Struct"`
	SourceBizCode                  *string                                                                      `json:"sourceBizCode,omitempty" xml:"sourceBizCode,omitempty"`
}

func (s QueryOrgPointDetailsResponseBodyResultDetails) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgPointDetailsResponseBodyResultDetails) GoString() string {
	return s.String()
}

func (s *QueryOrgPointDetailsResponseBodyResultDetails) SetAmount(v int64) *QueryOrgPointDetailsResponseBodyResultDetails {
	s.Amount = &v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResultDetails) SetGmtCreate(v int64) *QueryOrgPointDetailsResponseBodyResultDetails {
	s.GmtCreate = &v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResultDetails) SetOutId(v string) *QueryOrgPointDetailsResponseBodyResultDetails {
	s.OutId = &v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResultDetails) SetPointOperateFeatureResponseDTO(v *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) *QueryOrgPointDetailsResponseBodyResultDetails {
	s.PointOperateFeatureResponseDTO = v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResultDetails) SetSourceBizCode(v string) *QueryOrgPointDetailsResponseBodyResultDetails {
	s.SourceBizCode = &v
	return s
}

type QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO struct {
	AccountSource *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource `json:"accountSource,omitempty" xml:"accountSource,omitempty" type:"Struct"`
	AccountTarget *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget `json:"accountTarget,omitempty" xml:"accountTarget,omitempty" type:"Struct"`
	Remark        *string                                                                                   `json:"remark,omitempty" xml:"remark,omitempty"`
	Usage         *string                                                                                   `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) GoString() string {
	return s.String()
}

func (s *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) SetAccountSource(v *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource) *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO {
	s.AccountSource = v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) SetAccountTarget(v *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget) *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO {
	s.AccountTarget = v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) SetRemark(v string) *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO {
	s.Remark = &v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) SetUsage(v string) *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO {
	s.Usage = &v
	return s
}

type QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource struct {
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	EmpName     *string `json:"empName,omitempty" xml:"empName,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource) GoString() string {
	return s.String()
}

func (s *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource) SetAccountType(v string) *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource {
	s.AccountType = &v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource) SetEmpName(v string) *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource {
	s.EmpName = &v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource) SetUserId(v string) *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource {
	s.UserId = &v
	return s
}

type QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget struct {
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	EmpName     *string `json:"empName,omitempty" xml:"empName,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget) GoString() string {
	return s.String()
}

func (s *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget) SetAccountType(v string) *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget {
	s.AccountType = &v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget) SetEmpName(v string) *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget {
	s.EmpName = &v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget) SetUserId(v string) *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget {
	s.UserId = &v
	return s
}

type QueryOrgPointDetailsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOrgPointDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOrgPointDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgPointDetailsResponse) GoString() string {
	return s.String()
}

func (s *QueryOrgPointDetailsResponse) SetHeaders(v map[string]*string) *QueryOrgPointDetailsResponse {
	s.Headers = v
	return s
}

func (s *QueryOrgPointDetailsResponse) SetStatusCode(v int32) *QueryOrgPointDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrgPointDetailsResponse) SetBody(v *QueryOrgPointDetailsResponseBody) *QueryOrgPointDetailsResponse {
	s.Body = v
	return s
}

type QueryPointActionAutoAssignRuleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryPointActionAutoAssignRuleHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryPointActionAutoAssignRuleHeaders) GoString() string {
	return s.String()
}

func (s *QueryPointActionAutoAssignRuleHeaders) SetCommonHeaders(v map[string]*string) *QueryPointActionAutoAssignRuleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryPointActionAutoAssignRuleHeaders) SetXAcsDingtalkAccessToken(v string) *QueryPointActionAutoAssignRuleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryPointActionAutoAssignRuleResponseBody struct {
	Result  *QueryPointActionAutoAssignRuleResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryPointActionAutoAssignRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPointActionAutoAssignRuleResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPointActionAutoAssignRuleResponseBody) SetResult(v *QueryPointActionAutoAssignRuleResponseBodyResult) *QueryPointActionAutoAssignRuleResponseBody {
	s.Result = v
	return s
}

func (s *QueryPointActionAutoAssignRuleResponseBody) SetSuccess(v bool) *QueryPointActionAutoAssignRuleResponseBody {
	s.Success = &v
	return s
}

type QueryPointActionAutoAssignRuleResponseBodyResult struct {
	QueryPointRuleResponseDTOS []*QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS `json:"queryPointRuleResponseDTOS,omitempty" xml:"queryPointRuleResponseDTOS,omitempty" type:"Repeated"`
}

func (s QueryPointActionAutoAssignRuleResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryPointActionAutoAssignRuleResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryPointActionAutoAssignRuleResponseBodyResult) SetQueryPointRuleResponseDTOS(v []*QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS) *QueryPointActionAutoAssignRuleResponseBodyResult {
	s.QueryPointRuleResponseDTOS = v
	return s
}

type QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS struct {
	AwardScore    *int64  `json:"awardScore,omitempty" xml:"awardScore,omitempty"`
	Code          *string `json:"code,omitempty" xml:"code,omitempty"`
	DayLimitTimes *int64  `json:"dayLimitTimes,omitempty" xml:"dayLimitTimes,omitempty"`
	Description   *string `json:"description,omitempty" xml:"description,omitempty"`
	Status        *int64  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS) String() string {
	return tea.Prettify(s)
}

func (s QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS) GoString() string {
	return s.String()
}

func (s *QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS) SetAwardScore(v int64) *QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS {
	s.AwardScore = &v
	return s
}

func (s *QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS) SetCode(v string) *QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS {
	s.Code = &v
	return s
}

func (s *QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS) SetDayLimitTimes(v int64) *QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS {
	s.DayLimitTimes = &v
	return s
}

func (s *QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS) SetDescription(v string) *QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS {
	s.Description = &v
	return s
}

func (s *QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS) SetStatus(v int64) *QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS {
	s.Status = &v
	return s
}

type QueryPointActionAutoAssignRuleResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPointActionAutoAssignRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPointActionAutoAssignRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPointActionAutoAssignRuleResponse) GoString() string {
	return s.String()
}

func (s *QueryPointActionAutoAssignRuleResponse) SetHeaders(v map[string]*string) *QueryPointActionAutoAssignRuleResponse {
	s.Headers = v
	return s
}

func (s *QueryPointActionAutoAssignRuleResponse) SetStatusCode(v int32) *QueryPointActionAutoAssignRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPointActionAutoAssignRuleResponse) SetBody(v *QueryPointActionAutoAssignRuleResponseBody) *QueryPointActionAutoAssignRuleResponse {
	s.Body = v
	return s
}

type QueryPointAutoIssueSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryPointAutoIssueSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryPointAutoIssueSettingHeaders) GoString() string {
	return s.String()
}

func (s *QueryPointAutoIssueSettingHeaders) SetCommonHeaders(v map[string]*string) *QueryPointAutoIssueSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryPointAutoIssueSettingHeaders) SetXAcsDingtalkAccessToken(v string) *QueryPointAutoIssueSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryPointAutoIssueSettingResponseBody struct {
	Result  *QueryPointAutoIssueSettingResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryPointAutoIssueSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPointAutoIssueSettingResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPointAutoIssueSettingResponseBody) SetResult(v *QueryPointAutoIssueSettingResponseBodyResult) *QueryPointAutoIssueSettingResponseBody {
	s.Result = v
	return s
}

func (s *QueryPointAutoIssueSettingResponseBody) SetSuccess(v bool) *QueryPointAutoIssueSettingResponseBody {
	s.Success = &v
	return s
}

type QueryPointAutoIssueSettingResponseBodyResult struct {
	PointAutoNum   *int64 `json:"pointAutoNum,omitempty" xml:"pointAutoNum,omitempty"`
	PointAutoState *bool  `json:"pointAutoState,omitempty" xml:"pointAutoState,omitempty"`
	PointAutoTime  *int64 `json:"pointAutoTime,omitempty" xml:"pointAutoTime,omitempty"`
}

func (s QueryPointAutoIssueSettingResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryPointAutoIssueSettingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryPointAutoIssueSettingResponseBodyResult) SetPointAutoNum(v int64) *QueryPointAutoIssueSettingResponseBodyResult {
	s.PointAutoNum = &v
	return s
}

func (s *QueryPointAutoIssueSettingResponseBodyResult) SetPointAutoState(v bool) *QueryPointAutoIssueSettingResponseBodyResult {
	s.PointAutoState = &v
	return s
}

func (s *QueryPointAutoIssueSettingResponseBodyResult) SetPointAutoTime(v int64) *QueryPointAutoIssueSettingResponseBodyResult {
	s.PointAutoTime = &v
	return s
}

type QueryPointAutoIssueSettingResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPointAutoIssueSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPointAutoIssueSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPointAutoIssueSettingResponse) GoString() string {
	return s.String()
}

func (s *QueryPointAutoIssueSettingResponse) SetHeaders(v map[string]*string) *QueryPointAutoIssueSettingResponse {
	s.Headers = v
	return s
}

func (s *QueryPointAutoIssueSettingResponse) SetStatusCode(v int32) *QueryPointAutoIssueSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPointAutoIssueSettingResponse) SetBody(v *QueryPointAutoIssueSettingResponseBody) *QueryPointAutoIssueSettingResponse {
	s.Body = v
	return s
}

type QueryUserHonorsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserHonorsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserHonorsHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsHeaders) SetCommonHeaders(v map[string]*string) *QueryUserHonorsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserHonorsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserHonorsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserHonorsRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryUserHonorsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserHonorsRequest) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsRequest) SetMaxResults(v int32) *QueryUserHonorsRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryUserHonorsRequest) SetNextToken(v string) *QueryUserHonorsRequest {
	s.NextToken = &v
	return s
}

type QueryUserHonorsResponseBody struct {
	Result  *QueryUserHonorsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryUserHonorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserHonorsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsResponseBody) SetResult(v *QueryUserHonorsResponseBodyResult) *QueryUserHonorsResponseBody {
	s.Result = v
	return s
}

func (s *QueryUserHonorsResponseBody) SetSuccess(v bool) *QueryUserHonorsResponseBody {
	s.Success = &v
	return s
}

type QueryUserHonorsResponseBodyResult struct {
	Honors    []*QueryUserHonorsResponseBodyResultHonors `json:"honors,omitempty" xml:"honors,omitempty" type:"Repeated"`
	NextToken *string                                    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryUserHonorsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryUserHonorsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsResponseBodyResult) SetHonors(v []*QueryUserHonorsResponseBodyResultHonors) *QueryUserHonorsResponseBodyResult {
	s.Honors = v
	return s
}

func (s *QueryUserHonorsResponseBodyResult) SetNextToken(v string) *QueryUserHonorsResponseBodyResult {
	s.NextToken = &v
	return s
}

type QueryUserHonorsResponseBodyResultHonors struct {
	ExpirationTime *int64                                                 `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
	GrantHistory   []*QueryUserHonorsResponseBodyResultHonorsGrantHistory `json:"grantHistory,omitempty" xml:"grantHistory,omitempty" type:"Repeated"`
	HonorDesc      *string                                                `json:"honorDesc,omitempty" xml:"honorDesc,omitempty"`
	HonorId        *string                                                `json:"honorId,omitempty" xml:"honorId,omitempty"`
	HonorName      *string                                                `json:"honorName,omitempty" xml:"honorName,omitempty"`
}

func (s QueryUserHonorsResponseBodyResultHonors) String() string {
	return tea.Prettify(s)
}

func (s QueryUserHonorsResponseBodyResultHonors) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsResponseBodyResultHonors) SetExpirationTime(v int64) *QueryUserHonorsResponseBodyResultHonors {
	s.ExpirationTime = &v
	return s
}

func (s *QueryUserHonorsResponseBodyResultHonors) SetGrantHistory(v []*QueryUserHonorsResponseBodyResultHonorsGrantHistory) *QueryUserHonorsResponseBodyResultHonors {
	s.GrantHistory = v
	return s
}

func (s *QueryUserHonorsResponseBodyResultHonors) SetHonorDesc(v string) *QueryUserHonorsResponseBodyResultHonors {
	s.HonorDesc = &v
	return s
}

func (s *QueryUserHonorsResponseBodyResultHonors) SetHonorId(v string) *QueryUserHonorsResponseBodyResultHonors {
	s.HonorId = &v
	return s
}

func (s *QueryUserHonorsResponseBodyResultHonors) SetHonorName(v string) *QueryUserHonorsResponseBodyResultHonors {
	s.HonorName = &v
	return s
}

type QueryUserHonorsResponseBodyResultHonorsGrantHistory struct {
	GrantTime    *int64  `json:"grantTime,omitempty" xml:"grantTime,omitempty"`
	SenderUserid *string `json:"senderUserid,omitempty" xml:"senderUserid,omitempty"`
}

func (s QueryUserHonorsResponseBodyResultHonorsGrantHistory) String() string {
	return tea.Prettify(s)
}

func (s QueryUserHonorsResponseBodyResultHonorsGrantHistory) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsResponseBodyResultHonorsGrantHistory) SetGrantTime(v int64) *QueryUserHonorsResponseBodyResultHonorsGrantHistory {
	s.GrantTime = &v
	return s
}

func (s *QueryUserHonorsResponseBodyResultHonorsGrantHistory) SetSenderUserid(v string) *QueryUserHonorsResponseBodyResultHonorsGrantHistory {
	s.SenderUserid = &v
	return s
}

type QueryUserHonorsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserHonorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserHonorsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserHonorsResponse) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsResponse) SetHeaders(v map[string]*string) *QueryUserHonorsResponse {
	s.Headers = v
	return s
}

func (s *QueryUserHonorsResponse) SetStatusCode(v int32) *QueryUserHonorsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserHonorsResponse) SetBody(v *QueryUserHonorsResponseBody) *QueryUserHonorsResponse {
	s.Body = v
	return s
}

type QueryUserPointsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserPointsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserPointsHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserPointsHeaders) SetCommonHeaders(v map[string]*string) *QueryUserPointsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserPointsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserPointsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserPointsResponseBody struct {
	Result  *QueryUserPointsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryUserPointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserPointsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserPointsResponseBody) SetResult(v *QueryUserPointsResponseBodyResult) *QueryUserPointsResponseBody {
	s.Result = v
	return s
}

func (s *QueryUserPointsResponseBody) SetSuccess(v bool) *QueryUserPointsResponseBody {
	s.Success = &v
	return s
}

type QueryUserPointsResponseBodyResult struct {
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

func (s QueryUserPointsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryUserPointsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserPointsResponseBodyResult) SetAmount(v int64) *QueryUserPointsResponseBodyResult {
	s.Amount = &v
	return s
}

type QueryUserPointsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserPointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserPointsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserPointsResponse) GoString() string {
	return s.String()
}

func (s *QueryUserPointsResponse) SetHeaders(v map[string]*string) *QueryUserPointsResponse {
	s.Headers = v
	return s
}

func (s *QueryUserPointsResponse) SetStatusCode(v int32) *QueryUserPointsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserPointsResponse) SetBody(v *QueryUserPointsResponseBody) *QueryUserPointsResponse {
	s.Body = v
	return s
}

type RecallHonorHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RecallHonorHeaders) String() string {
	return tea.Prettify(s)
}

func (s RecallHonorHeaders) GoString() string {
	return s.String()
}

func (s *RecallHonorHeaders) SetCommonHeaders(v map[string]*string) *RecallHonorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RecallHonorHeaders) SetXAcsDingtalkAccessToken(v string) *RecallHonorHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RecallHonorRequest struct {
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s RecallHonorRequest) String() string {
	return tea.Prettify(s)
}

func (s RecallHonorRequest) GoString() string {
	return s.String()
}

func (s *RecallHonorRequest) SetUserId(v string) *RecallHonorRequest {
	s.UserId = &v
	return s
}

type RecallHonorResponseBody struct {
	Result  *RecallHonorResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RecallHonorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecallHonorResponseBody) GoString() string {
	return s.String()
}

func (s *RecallHonorResponseBody) SetResult(v *RecallHonorResponseBodyResult) *RecallHonorResponseBody {
	s.Result = v
	return s
}

func (s *RecallHonorResponseBody) SetSuccess(v bool) *RecallHonorResponseBody {
	s.Success = &v
	return s
}

type RecallHonorResponseBodyResult struct {
	HonorId *string `json:"honorId,omitempty" xml:"honorId,omitempty"`
}

func (s RecallHonorResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s RecallHonorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RecallHonorResponseBodyResult) SetHonorId(v string) *RecallHonorResponseBodyResult {
	s.HonorId = &v
	return s
}

type RecallHonorResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecallHonorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecallHonorResponse) String() string {
	return tea.Prettify(s)
}

func (s RecallHonorResponse) GoString() string {
	return s.String()
}

func (s *RecallHonorResponse) SetHeaders(v map[string]*string) *RecallHonorResponse {
	s.Headers = v
	return s
}

func (s *RecallHonorResponse) SetStatusCode(v int32) *RecallHonorResponse {
	s.StatusCode = &v
	return s
}

func (s *RecallHonorResponse) SetBody(v *RecallHonorResponseBody) *RecallHonorResponse {
	s.Body = v
	return s
}

type UpdateAutoIssuePointHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateAutoIssuePointHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateAutoIssuePointHeaders) GoString() string {
	return s.String()
}

func (s *UpdateAutoIssuePointHeaders) SetCommonHeaders(v map[string]*string) *UpdateAutoIssuePointHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateAutoIssuePointHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateAutoIssuePointHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateAutoIssuePointRequest struct {
	PointAutoNum   *int64  `json:"pointAutoNum,omitempty" xml:"pointAutoNum,omitempty"`
	PointAutoState *bool   `json:"pointAutoState,omitempty" xml:"pointAutoState,omitempty"`
	PointAutoTime  *int64  `json:"pointAutoTime,omitempty" xml:"pointAutoTime,omitempty"`
	UserId         *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateAutoIssuePointRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAutoIssuePointRequest) GoString() string {
	return s.String()
}

func (s *UpdateAutoIssuePointRequest) SetPointAutoNum(v int64) *UpdateAutoIssuePointRequest {
	s.PointAutoNum = &v
	return s
}

func (s *UpdateAutoIssuePointRequest) SetPointAutoState(v bool) *UpdateAutoIssuePointRequest {
	s.PointAutoState = &v
	return s
}

func (s *UpdateAutoIssuePointRequest) SetPointAutoTime(v int64) *UpdateAutoIssuePointRequest {
	s.PointAutoTime = &v
	return s
}

func (s *UpdateAutoIssuePointRequest) SetUserId(v string) *UpdateAutoIssuePointRequest {
	s.UserId = &v
	return s
}

type UpdateAutoIssuePointResponseBody struct {
	Result  *UpdateAutoIssuePointResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateAutoIssuePointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAutoIssuePointResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAutoIssuePointResponseBody) SetResult(v *UpdateAutoIssuePointResponseBodyResult) *UpdateAutoIssuePointResponseBody {
	s.Result = v
	return s
}

func (s *UpdateAutoIssuePointResponseBody) SetSuccess(v bool) *UpdateAutoIssuePointResponseBody {
	s.Success = &v
	return s
}

type UpdateAutoIssuePointResponseBodyResult struct {
	NextAutoIssuePointTime *int64 `json:"nextAutoIssuePointTime,omitempty" xml:"nextAutoIssuePointTime,omitempty"`
}

func (s UpdateAutoIssuePointResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateAutoIssuePointResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateAutoIssuePointResponseBodyResult) SetNextAutoIssuePointTime(v int64) *UpdateAutoIssuePointResponseBodyResult {
	s.NextAutoIssuePointTime = &v
	return s
}

type UpdateAutoIssuePointResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAutoIssuePointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAutoIssuePointResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAutoIssuePointResponse) GoString() string {
	return s.String()
}

func (s *UpdateAutoIssuePointResponse) SetHeaders(v map[string]*string) *UpdateAutoIssuePointResponse {
	s.Headers = v
	return s
}

func (s *UpdateAutoIssuePointResponse) SetStatusCode(v int32) *UpdateAutoIssuePointResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAutoIssuePointResponse) SetBody(v *UpdateAutoIssuePointResponseBody) *UpdateAutoIssuePointResponse {
	s.Body = v
	return s
}

type UpdatePointActionAutoAssignRuleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdatePointActionAutoAssignRuleHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdatePointActionAutoAssignRuleHeaders) GoString() string {
	return s.String()
}

func (s *UpdatePointActionAutoAssignRuleHeaders) SetCommonHeaders(v map[string]*string) *UpdatePointActionAutoAssignRuleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdatePointActionAutoAssignRuleHeaders) SetXAcsDingtalkAccessToken(v string) *UpdatePointActionAutoAssignRuleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdatePointActionAutoAssignRuleRequest struct {
	UpdatePointRuleRequestDTOList []*UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList `json:"updatePointRuleRequestDTOList,omitempty" xml:"updatePointRuleRequestDTOList,omitempty" type:"Repeated"`
	UserId                        *string                                                                `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdatePointActionAutoAssignRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePointActionAutoAssignRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdatePointActionAutoAssignRuleRequest) SetUpdatePointRuleRequestDTOList(v []*UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList) *UpdatePointActionAutoAssignRuleRequest {
	s.UpdatePointRuleRequestDTOList = v
	return s
}

func (s *UpdatePointActionAutoAssignRuleRequest) SetUserId(v string) *UpdatePointActionAutoAssignRuleRequest {
	s.UserId = &v
	return s
}

type UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList struct {
	AwardScore    *int64  `json:"awardScore,omitempty" xml:"awardScore,omitempty"`
	Code          *string `json:"code,omitempty" xml:"code,omitempty"`
	DayLimitTimes *int64  `json:"dayLimitTimes,omitempty" xml:"dayLimitTimes,omitempty"`
	Status        *int64  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList) String() string {
	return tea.Prettify(s)
}

func (s UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList) GoString() string {
	return s.String()
}

func (s *UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList) SetAwardScore(v int64) *UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList {
	s.AwardScore = &v
	return s
}

func (s *UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList) SetCode(v string) *UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList {
	s.Code = &v
	return s
}

func (s *UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList) SetDayLimitTimes(v int64) *UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList {
	s.DayLimitTimes = &v
	return s
}

func (s *UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList) SetStatus(v int64) *UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList {
	s.Status = &v
	return s
}

type UpdatePointActionAutoAssignRuleResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdatePointActionAutoAssignRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePointActionAutoAssignRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePointActionAutoAssignRuleResponseBody) SetSuccess(v bool) *UpdatePointActionAutoAssignRuleResponseBody {
	s.Success = &v
	return s
}

type UpdatePointActionAutoAssignRuleResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePointActionAutoAssignRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePointActionAutoAssignRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePointActionAutoAssignRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdatePointActionAutoAssignRuleResponse) SetHeaders(v map[string]*string) *UpdatePointActionAutoAssignRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdatePointActionAutoAssignRuleResponse) SetStatusCode(v int32) *UpdatePointActionAutoAssignRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePointActionAutoAssignRuleResponse) SetBody(v *UpdatePointActionAutoAssignRuleResponseBody) *UpdatePointActionAutoAssignRuleResponse {
	s.Body = v
	return s
}

type WearOrgHonorHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s WearOrgHonorHeaders) String() string {
	return tea.Prettify(s)
}

func (s WearOrgHonorHeaders) GoString() string {
	return s.String()
}

func (s *WearOrgHonorHeaders) SetCommonHeaders(v map[string]*string) *WearOrgHonorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *WearOrgHonorHeaders) SetXAcsDingtalkAccessToken(v string) *WearOrgHonorHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type WearOrgHonorRequest struct {
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	Wear   *bool   `json:"wear,omitempty" xml:"wear,omitempty"`
}

func (s WearOrgHonorRequest) String() string {
	return tea.Prettify(s)
}

func (s WearOrgHonorRequest) GoString() string {
	return s.String()
}

func (s *WearOrgHonorRequest) SetUserId(v string) *WearOrgHonorRequest {
	s.UserId = &v
	return s
}

func (s *WearOrgHonorRequest) SetWear(v bool) *WearOrgHonorRequest {
	s.Wear = &v
	return s
}

type WearOrgHonorResponseBody struct {
	Result  *WearOrgHonorResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s WearOrgHonorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WearOrgHonorResponseBody) GoString() string {
	return s.String()
}

func (s *WearOrgHonorResponseBody) SetResult(v *WearOrgHonorResponseBodyResult) *WearOrgHonorResponseBody {
	s.Result = v
	return s
}

func (s *WearOrgHonorResponseBody) SetSuccess(v bool) *WearOrgHonorResponseBody {
	s.Success = &v
	return s
}

type WearOrgHonorResponseBodyResult struct {
	HonorId *string `json:"honorId,omitempty" xml:"honorId,omitempty"`
}

func (s WearOrgHonorResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s WearOrgHonorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *WearOrgHonorResponseBodyResult) SetHonorId(v string) *WearOrgHonorResponseBodyResult {
	s.HonorId = &v
	return s
}

type WearOrgHonorResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WearOrgHonorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WearOrgHonorResponse) String() string {
	return tea.Prettify(s)
}

func (s WearOrgHonorResponse) GoString() string {
	return s.String()
}

func (s *WearOrgHonorResponse) SetHeaders(v map[string]*string) *WearOrgHonorResponse {
	s.Headers = v
	return s
}

func (s *WearOrgHonorResponse) SetStatusCode(v int32) *WearOrgHonorResponse {
	s.StatusCode = &v
	return s
}

func (s *WearOrgHonorResponse) SetBody(v *WearOrgHonorResponseBody) *WearOrgHonorResponse {
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

func (client *Client) AssignOrgHoldingToEmpHoldingBatchWithOptions(request *AssignOrgHoldingToEmpHoldingBatchRequest, headers *AssignOrgHoldingToEmpHoldingBatchHeaders, runtime *util.RuntimeOptions) (_result *AssignOrgHoldingToEmpHoldingBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.SendOrgCultureInform)) {
		body["sendOrgCultureInform"] = request.SendOrgCultureInform
	}

	if !tea.BoolValue(util.IsUnset(request.SingleAmount)) {
		body["singleAmount"] = request.SingleAmount
	}

	if !tea.BoolValue(util.IsUnset(request.SourceUsage)) {
		body["sourceUsage"] = request.SourceUsage
	}

	if !tea.BoolValue(util.IsUnset(request.TargetUsage)) {
		body["targetUsage"] = request.TargetUsage
	}

	if !tea.BoolValue(util.IsUnset(request.TargetUserList)) {
		body["targetUserList"] = request.TargetUserList
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
		Action:      tea.String("AssignOrgHoldingToEmpHoldingBatch"),
		Version:     tea.String("orgCulture_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/orgCulture/organizations/points/assign"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AssignOrgHoldingToEmpHoldingBatchResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssignOrgHoldingToEmpHoldingBatch(request *AssignOrgHoldingToEmpHoldingBatchRequest) (_result *AssignOrgHoldingToEmpHoldingBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AssignOrgHoldingToEmpHoldingBatchHeaders{}
	_result = &AssignOrgHoldingToEmpHoldingBatchResponse{}
	_body, _err := client.AssignOrgHoldingToEmpHoldingBatchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConsumeUserPointsWithOptions(userId *string, request *ConsumeUserPointsRequest, headers *ConsumeUserPointsHeaders, runtime *util.RuntimeOptions) (_result *ConsumeUserPointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		body["amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		body["outId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Usage)) {
		body["usage"] = request.Usage
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
		Action:      tea.String("ConsumeUserPoints"),
		Version:     tea.String("orgCulture_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/orgCulture/users/" + tea.StringValue(userId) + "/points/deduct"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ConsumeUserPointsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConsumeUserPoints(userId *string, request *ConsumeUserPointsRequest) (_result *ConsumeUserPointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ConsumeUserPointsHeaders{}
	_result = &ConsumeUserPointsResponse{}
	_body, _err := client.ConsumeUserPointsWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrgHonorWithOptions(request *CreateOrgHonorRequest, headers *CreateOrgHonorHeaders, runtime *util.RuntimeOptions) (_result *CreateOrgHonorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AvatarFrameMediaId)) {
		body["avatarFrameMediaId"] = request.AvatarFrameMediaId
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultBgColor)) {
		body["defaultBgColor"] = request.DefaultBgColor
	}

	if !tea.BoolValue(util.IsUnset(request.MedalDesc)) {
		body["medalDesc"] = request.MedalDesc
	}

	if !tea.BoolValue(util.IsUnset(request.MedalMediaId)) {
		body["medalMediaId"] = request.MedalMediaId
	}

	if !tea.BoolValue(util.IsUnset(request.MedalName)) {
		body["medalName"] = request.MedalName
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
		Action:      tea.String("CreateOrgHonor"),
		Version:     tea.String("orgCulture_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/orgCulture/honors/templates"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOrgHonorResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOrgHonor(request *CreateOrgHonorRequest) (_result *CreateOrgHonorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateOrgHonorHeaders{}
	_result = &CreateOrgHonorResponse{}
	_body, _err := client.CreateOrgHonorWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeductionPointBatchWithOptions(request *DeductionPointBatchRequest, headers *DeductionPointBatchHeaders, runtime *util.RuntimeOptions) (_result *DeductionPointBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeductionAmount)) {
		body["deductionAmount"] = request.DeductionAmount
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.SendOrgCultureInform)) {
		body["sendOrgCultureInform"] = request.SendOrgCultureInform
	}

	if !tea.BoolValue(util.IsUnset(request.TargetUserList)) {
		body["targetUserList"] = request.TargetUserList
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
		Action:      tea.String("DeductionPointBatch"),
		Version:     tea.String("orgCulture_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/orgCulture/users/points/deduct"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeductionPointBatchResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeductionPointBatch(request *DeductionPointBatchRequest) (_result *DeductionPointBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeductionPointBatchHeaders{}
	_result = &DeductionPointBatchResponse{}
	_body, _err := client.DeductionPointBatchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExportPointOpenWithOptions(request *ExportPointOpenRequest, headers *ExportPointOpenHeaders, runtime *util.RuntimeOptions) (_result *ExportPointOpenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExportDate)) {
		body["exportDate"] = request.ExportDate
	}

	if !tea.BoolValue(util.IsUnset(request.ExportType)) {
		body["exportType"] = request.ExportType
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
		Action:      tea.String("ExportPointOpen"),
		Version:     tea.String("orgCulture_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/orgCulture/users/points/export"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportPointOpenResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExportPointOpen(request *ExportPointOpenRequest) (_result *ExportPointOpenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ExportPointOpenHeaders{}
	_result = &ExportPointOpenResponse{}
	_body, _err := client.ExportPointOpenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrantHonorWithOptions(honorId *string, request *GrantHonorRequest, headers *GrantHonorHeaders, runtime *util.RuntimeOptions) (_result *GrantHonorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExpirationTime)) {
		body["expirationTime"] = request.ExpirationTime
	}

	if !tea.BoolValue(util.IsUnset(request.GrantReason)) {
		body["grantReason"] = request.GrantReason
	}

	if !tea.BoolValue(util.IsUnset(request.GranterName)) {
		body["granterName"] = request.GranterName
	}

	if !tea.BoolValue(util.IsUnset(request.NoticeAnnouncer)) {
		body["noticeAnnouncer"] = request.NoticeAnnouncer
	}

	if !tea.BoolValue(util.IsUnset(request.NoticeSingle)) {
		body["noticeSingle"] = request.NoticeSingle
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationIds)) {
		body["openConversationIds"] = request.OpenConversationIds
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUserIds)) {
		body["receiverUserIds"] = request.ReceiverUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.SenderUserId)) {
		body["senderUserId"] = request.SenderUserId
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
		Action:      tea.String("GrantHonor"),
		Version:     tea.String("orgCulture_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/orgCulture/honors/" + tea.StringValue(honorId) + "/grant"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GrantHonorResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrantHonor(honorId *string, request *GrantHonorRequest) (_result *GrantHonorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GrantHonorHeaders{}
	_result = &GrantHonorResponse{}
	_body, _err := client.GrantHonorWithOptions(honorId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCorpPointsWithOptions(request *QueryCorpPointsRequest, headers *QueryCorpPointsHeaders, runtime *util.RuntimeOptions) (_result *QueryCorpPointsResponse, _err error) {
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
		Action:      tea.String("QueryCorpPoints"),
		Version:     tea.String("orgCulture_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/orgCulture/organizations/points"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCorpPointsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCorpPoints(request *QueryCorpPointsRequest) (_result *QueryCorpPointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCorpPointsHeaders{}
	_result = &QueryCorpPointsResponse{}
	_body, _err := client.QueryCorpPointsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryEmpPointDetailsWithOptions(request *QueryEmpPointDetailsRequest, headers *QueryEmpPointDetailsHeaders, runtime *util.RuntimeOptions) (_result *QueryEmpPointDetailsResponse, _err error) {
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
		Action:      tea.String("QueryEmpPointDetails"),
		Version:     tea.String("orgCulture_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/orgCulture/points/empDetails"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryEmpPointDetailsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryEmpPointDetails(request *QueryEmpPointDetailsRequest) (_result *QueryEmpPointDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryEmpPointDetailsHeaders{}
	_result = &QueryEmpPointDetailsResponse{}
	_body, _err := client.QueryEmpPointDetailsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrgHonorsWithOptions(request *QueryOrgHonorsRequest, headers *QueryOrgHonorsHeaders, runtime *util.RuntimeOptions) (_result *QueryOrgHonorsResponse, _err error) {
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
		Action:      tea.String("QueryOrgHonors"),
		Version:     tea.String("orgCulture_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/orgCulture/organizations/honors"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOrgHonorsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrgHonors(request *QueryOrgHonorsRequest) (_result *QueryOrgHonorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOrgHonorsHeaders{}
	_result = &QueryOrgHonorsResponse{}
	_body, _err := client.QueryOrgHonorsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrgPointDetailsWithOptions(request *QueryOrgPointDetailsRequest, headers *QueryOrgPointDetailsHeaders, runtime *util.RuntimeOptions) (_result *QueryOrgPointDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountType)) {
		query["accountType"] = request.AccountType
	}

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
		Action:      tea.String("QueryOrgPointDetails"),
		Version:     tea.String("orgCulture_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/orgCulture/points/orgDetails"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOrgPointDetailsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrgPointDetails(request *QueryOrgPointDetailsRequest) (_result *QueryOrgPointDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOrgPointDetailsHeaders{}
	_result = &QueryOrgPointDetailsResponse{}
	_body, _err := client.QueryOrgPointDetailsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPointActionAutoAssignRuleWithOptions(headers *QueryPointActionAutoAssignRuleHeaders, runtime *util.RuntimeOptions) (_result *QueryPointActionAutoAssignRuleResponse, _err error) {
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
		Action:      tea.String("QueryPointActionAutoAssignRule"),
		Version:     tea.String("orgCulture_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/orgCulture/users/points/actionRules"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPointActionAutoAssignRuleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPointActionAutoAssignRule() (_result *QueryPointActionAutoAssignRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryPointActionAutoAssignRuleHeaders{}
	_result = &QueryPointActionAutoAssignRuleResponse{}
	_body, _err := client.QueryPointActionAutoAssignRuleWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPointAutoIssueSettingWithOptions(headers *QueryPointAutoIssueSettingHeaders, runtime *util.RuntimeOptions) (_result *QueryPointAutoIssueSettingResponse, _err error) {
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
		Action:      tea.String("QueryPointAutoIssueSetting"),
		Version:     tea.String("orgCulture_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/orgCulture/users/points"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPointAutoIssueSettingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPointAutoIssueSetting() (_result *QueryPointAutoIssueSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryPointAutoIssueSettingHeaders{}
	_result = &QueryPointAutoIssueSettingResponse{}
	_body, _err := client.QueryPointAutoIssueSettingWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserHonorsWithOptions(userId *string, request *QueryUserHonorsRequest, headers *QueryUserHonorsHeaders, runtime *util.RuntimeOptions) (_result *QueryUserHonorsResponse, _err error) {
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
		Action:      tea.String("QueryUserHonors"),
		Version:     tea.String("orgCulture_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/orgCulture/honors/users/" + tea.StringValue(userId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserHonorsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserHonors(userId *string, request *QueryUserHonorsRequest) (_result *QueryUserHonorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserHonorsHeaders{}
	_result = &QueryUserHonorsResponse{}
	_body, _err := client.QueryUserHonorsWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserPointsWithOptions(userId *string, headers *QueryUserPointsHeaders, runtime *util.RuntimeOptions) (_result *QueryUserPointsResponse, _err error) {
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
		Action:      tea.String("QueryUserPoints"),
		Version:     tea.String("orgCulture_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/orgCulture/users/" + tea.StringValue(userId) + "/points"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserPointsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserPoints(userId *string) (_result *QueryUserPointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserPointsHeaders{}
	_result = &QueryUserPointsResponse{}
	_body, _err := client.QueryUserPointsWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecallHonorWithOptions(honorId *string, request *RecallHonorRequest, headers *RecallHonorHeaders, runtime *util.RuntimeOptions) (_result *RecallHonorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
		Action:      tea.String("RecallHonor"),
		Version:     tea.String("orgCulture_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/orgCulture/honors/" + tea.StringValue(honorId) + "/recall"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RecallHonorResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecallHonor(honorId *string, request *RecallHonorRequest) (_result *RecallHonorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RecallHonorHeaders{}
	_result = &RecallHonorResponse{}
	_body, _err := client.RecallHonorWithOptions(honorId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAutoIssuePointWithOptions(request *UpdateAutoIssuePointRequest, headers *UpdateAutoIssuePointHeaders, runtime *util.RuntimeOptions) (_result *UpdateAutoIssuePointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PointAutoNum)) {
		body["pointAutoNum"] = request.PointAutoNum
	}

	if !tea.BoolValue(util.IsUnset(request.PointAutoState)) {
		body["pointAutoState"] = request.PointAutoState
	}

	if !tea.BoolValue(util.IsUnset(request.PointAutoTime)) {
		body["pointAutoTime"] = request.PointAutoTime
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
		Action:      tea.String("UpdateAutoIssuePoint"),
		Version:     tea.String("orgCulture_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/orgCulture/users/points/set"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAutoIssuePointResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAutoIssuePoint(request *UpdateAutoIssuePointRequest) (_result *UpdateAutoIssuePointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateAutoIssuePointHeaders{}
	_result = &UpdateAutoIssuePointResponse{}
	_body, _err := client.UpdateAutoIssuePointWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePointActionAutoAssignRuleWithOptions(request *UpdatePointActionAutoAssignRuleRequest, headers *UpdatePointActionAutoAssignRuleHeaders, runtime *util.RuntimeOptions) (_result *UpdatePointActionAutoAssignRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UpdatePointRuleRequestDTOList)) {
		body["updatePointRuleRequestDTOList"] = request.UpdatePointRuleRequestDTOList
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
		Action:      tea.String("UpdatePointActionAutoAssignRule"),
		Version:     tea.String("orgCulture_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/orgCulture/users/points/actionRules"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePointActionAutoAssignRuleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePointActionAutoAssignRule(request *UpdatePointActionAutoAssignRuleRequest) (_result *UpdatePointActionAutoAssignRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdatePointActionAutoAssignRuleHeaders{}
	_result = &UpdatePointActionAutoAssignRuleResponse{}
	_body, _err := client.UpdatePointActionAutoAssignRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) WearOrgHonorWithOptions(honorId *string, request *WearOrgHonorRequest, headers *WearOrgHonorHeaders, runtime *util.RuntimeOptions) (_result *WearOrgHonorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Wear)) {
		body["wear"] = request.Wear
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
		Action:      tea.String("WearOrgHonor"),
		Version:     tea.String("orgCulture_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/orgCulture/honors/" + tea.StringValue(honorId) + "/wear"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &WearOrgHonorResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) WearOrgHonor(honorId *string, request *WearOrgHonorRequest) (_result *WearOrgHonorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &WearOrgHonorHeaders{}
	_result = &WearOrgHonorResponse{}
	_body, _err := client.WearOrgHonorWithOptions(honorId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
