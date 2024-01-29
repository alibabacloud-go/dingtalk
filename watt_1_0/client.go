// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package watt_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CheckInCrowdsByMobileRequest struct {
	CrowdIds []byte  `json:"crowdIds,omitempty" xml:"crowdIds,omitempty"`
	Mobile   *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
}

func (s CheckInCrowdsByMobileRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckInCrowdsByMobileRequest) GoString() string {
	return s.String()
}

func (s *CheckInCrowdsByMobileRequest) SetCrowdIds(v []byte) *CheckInCrowdsByMobileRequest {
	s.CrowdIds = v
	return s
}

func (s *CheckInCrowdsByMobileRequest) SetMobile(v string) *CheckInCrowdsByMobileRequest {
	s.Mobile = &v
	return s
}

type CheckInCrowdsByMobileResponseBody struct {
	Data    *bool  `json:"data,omitempty" xml:"data,omitempty"`
	Success *bool  `json:"success,omitempty" xml:"success,omitempty"`
	Total   *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s CheckInCrowdsByMobileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckInCrowdsByMobileResponseBody) GoString() string {
	return s.String()
}

func (s *CheckInCrowdsByMobileResponseBody) SetData(v bool) *CheckInCrowdsByMobileResponseBody {
	s.Data = &v
	return s
}

func (s *CheckInCrowdsByMobileResponseBody) SetSuccess(v bool) *CheckInCrowdsByMobileResponseBody {
	s.Success = &v
	return s
}

func (s *CheckInCrowdsByMobileResponseBody) SetTotal(v int32) *CheckInCrowdsByMobileResponseBody {
	s.Total = &v
	return s
}

type CheckInCrowdsByMobileResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckInCrowdsByMobileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckInCrowdsByMobileResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckInCrowdsByMobileResponse) GoString() string {
	return s.String()
}

func (s *CheckInCrowdsByMobileResponse) SetHeaders(v map[string]*string) *CheckInCrowdsByMobileResponse {
	s.Headers = v
	return s
}

func (s *CheckInCrowdsByMobileResponse) SetStatusCode(v int32) *CheckInCrowdsByMobileResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckInCrowdsByMobileResponse) SetBody(v *CheckInCrowdsByMobileResponseBody) *CheckInCrowdsByMobileResponse {
	s.Body = v
	return s
}

type ConsumePointHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ConsumePointHeaders) String() string {
	return tea.Prettify(s)
}

func (s ConsumePointHeaders) GoString() string {
	return s.String()
}

func (s *ConsumePointHeaders) SetCommonHeaders(v map[string]*string) *ConsumePointHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ConsumePointHeaders) SetXAcsDingtalkAccessToken(v string) *ConsumePointHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ConsumePointRequest struct {
	Body *ConsumePointRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s ConsumePointRequest) String() string {
	return tea.Prettify(s)
}

func (s ConsumePointRequest) GoString() string {
	return s.String()
}

func (s *ConsumePointRequest) SetBody(v *ConsumePointRequestBody) *ConsumePointRequest {
	s.Body = v
	return s
}

type ConsumePointRequestBody struct {
	ConsumeDetail *ConsumePointRequestBodyConsumeDetail `json:"consumeDetail,omitempty" xml:"consumeDetail,omitempty" type:"Struct"`
	PointPoolCode *string                               `json:"pointPoolCode,omitempty" xml:"pointPoolCode,omitempty"`
	Points        *int64                                `json:"points,omitempty" xml:"points,omitempty"`
	RequestId     *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ConsumePointRequestBody) String() string {
	return tea.Prettify(s)
}

func (s ConsumePointRequestBody) GoString() string {
	return s.String()
}

func (s *ConsumePointRequestBody) SetConsumeDetail(v *ConsumePointRequestBodyConsumeDetail) *ConsumePointRequestBody {
	s.ConsumeDetail = v
	return s
}

func (s *ConsumePointRequestBody) SetPointPoolCode(v string) *ConsumePointRequestBody {
	s.PointPoolCode = &v
	return s
}

func (s *ConsumePointRequestBody) SetPoints(v int64) *ConsumePointRequestBody {
	s.Points = &v
	return s
}

func (s *ConsumePointRequestBody) SetRequestId(v string) *ConsumePointRequestBody {
	s.RequestId = &v
	return s
}

type ConsumePointRequestBodyConsumeDetail struct {
	Benefit *ConsumePointRequestBodyConsumeDetailBenefit `json:"benefit,omitempty" xml:"benefit,omitempty" type:"Struct"`
	Type    *string                                      `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ConsumePointRequestBodyConsumeDetail) String() string {
	return tea.Prettify(s)
}

func (s ConsumePointRequestBodyConsumeDetail) GoString() string {
	return s.String()
}

func (s *ConsumePointRequestBodyConsumeDetail) SetBenefit(v *ConsumePointRequestBodyConsumeDetailBenefit) *ConsumePointRequestBodyConsumeDetail {
	s.Benefit = v
	return s
}

func (s *ConsumePointRequestBodyConsumeDetail) SetType(v string) *ConsumePointRequestBodyConsumeDetail {
	s.Type = &v
	return s
}

type ConsumePointRequestBodyConsumeDetailBenefit struct {
	BenefitId    *string `json:"benefitId,omitempty" xml:"benefitId,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	SupplierName *string `json:"supplierName,omitempty" xml:"supplierName,omitempty"`
	UseUrl       *string `json:"useUrl,omitempty" xml:"useUrl,omitempty"`
}

func (s ConsumePointRequestBodyConsumeDetailBenefit) String() string {
	return tea.Prettify(s)
}

func (s ConsumePointRequestBodyConsumeDetailBenefit) GoString() string {
	return s.String()
}

func (s *ConsumePointRequestBodyConsumeDetailBenefit) SetBenefitId(v string) *ConsumePointRequestBodyConsumeDetailBenefit {
	s.BenefitId = &v
	return s
}

func (s *ConsumePointRequestBodyConsumeDetailBenefit) SetName(v string) *ConsumePointRequestBodyConsumeDetailBenefit {
	s.Name = &v
	return s
}

func (s *ConsumePointRequestBodyConsumeDetailBenefit) SetSupplierName(v string) *ConsumePointRequestBodyConsumeDetailBenefit {
	s.SupplierName = &v
	return s
}

func (s *ConsumePointRequestBodyConsumeDetailBenefit) SetUseUrl(v string) *ConsumePointRequestBodyConsumeDetailBenefit {
	s.UseUrl = &v
	return s
}

type ConsumePointShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConsumePointShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ConsumePointShrinkRequest) GoString() string {
	return s.String()
}

func (s *ConsumePointShrinkRequest) SetBodyShrink(v string) *ConsumePointShrinkRequest {
	s.BodyShrink = &v
	return s
}

type ConsumePointResponseBody struct {
	Result  *ConsumePointResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ConsumePointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConsumePointResponseBody) GoString() string {
	return s.String()
}

func (s *ConsumePointResponseBody) SetResult(v *ConsumePointResponseBodyResult) *ConsumePointResponseBody {
	s.Result = v
	return s
}

func (s *ConsumePointResponseBody) SetSuccess(v bool) *ConsumePointResponseBody {
	s.Success = &v
	return s
}

type ConsumePointResponseBodyResult struct {
	ConsumedPoints *int64 `json:"consumedPoints,omitempty" xml:"consumedPoints,omitempty"`
}

func (s ConsumePointResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ConsumePointResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ConsumePointResponseBodyResult) SetConsumedPoints(v int64) *ConsumePointResponseBodyResult {
	s.ConsumedPoints = &v
	return s
}

type ConsumePointResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConsumePointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConsumePointResponse) String() string {
	return tea.Prettify(s)
}

func (s ConsumePointResponse) GoString() string {
	return s.String()
}

func (s *ConsumePointResponse) SetHeaders(v map[string]*string) *ConsumePointResponse {
	s.Headers = v
	return s
}

func (s *ConsumePointResponse) SetStatusCode(v int32) *ConsumePointResponse {
	s.StatusCode = &v
	return s
}

func (s *ConsumePointResponse) SetBody(v *ConsumePointResponseBody) *ConsumePointResponse {
	s.Body = v
	return s
}

type CreateDeliveryPlanHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateDeliveryPlanHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateDeliveryPlanHeaders) GoString() string {
	return s.String()
}

func (s *CreateDeliveryPlanHeaders) SetCommonHeaders(v map[string]*string) *CreateDeliveryPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateDeliveryPlanHeaders) SetXAcsDingtalkAccessToken(v string) *CreateDeliveryPlanHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateDeliveryPlanRequest struct {
	Content    map[string]interface{} `json:"content,omitempty" xml:"content,omitempty"`
	EndTime    *int64                 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ResId      *string                `json:"resId,omitempty" xml:"resId,omitempty"`
	StartTime  *int64                 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	UserIdList []*string              `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s CreateDeliveryPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeliveryPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateDeliveryPlanRequest) SetContent(v map[string]interface{}) *CreateDeliveryPlanRequest {
	s.Content = v
	return s
}

func (s *CreateDeliveryPlanRequest) SetEndTime(v int64) *CreateDeliveryPlanRequest {
	s.EndTime = &v
	return s
}

func (s *CreateDeliveryPlanRequest) SetResId(v string) *CreateDeliveryPlanRequest {
	s.ResId = &v
	return s
}

func (s *CreateDeliveryPlanRequest) SetStartTime(v int64) *CreateDeliveryPlanRequest {
	s.StartTime = &v
	return s
}

func (s *CreateDeliveryPlanRequest) SetUserIdList(v []*string) *CreateDeliveryPlanRequest {
	s.UserIdList = v
	return s
}

type CreateDeliveryPlanResponseBody struct {
	Arguments []interface{} `json:"arguments,omitempty" xml:"arguments,omitempty" type:"Repeated"`
	Success   *bool         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateDeliveryPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeliveryPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeliveryPlanResponseBody) SetArguments(v []interface{}) *CreateDeliveryPlanResponseBody {
	s.Arguments = v
	return s
}

func (s *CreateDeliveryPlanResponseBody) SetSuccess(v bool) *CreateDeliveryPlanResponseBody {
	s.Success = &v
	return s
}

type CreateDeliveryPlanResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDeliveryPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeliveryPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeliveryPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateDeliveryPlanResponse) SetHeaders(v map[string]*string) *CreateDeliveryPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateDeliveryPlanResponse) SetStatusCode(v int32) *CreateDeliveryPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeliveryPlanResponse) SetBody(v *CreateDeliveryPlanResponseBody) *CreateDeliveryPlanResponse {
	s.Body = v
	return s
}

type GetPointInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPointInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPointInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetPointInfoHeaders) SetCommonHeaders(v map[string]*string) *GetPointInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPointInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetPointInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPointInfoRequest struct {
	PointPoolCode *string `json:"pointPoolCode,omitempty" xml:"pointPoolCode,omitempty"`
}

func (s GetPointInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPointInfoRequest) GoString() string {
	return s.String()
}

func (s *GetPointInfoRequest) SetPointPoolCode(v string) *GetPointInfoRequest {
	s.PointPoolCode = &v
	return s
}

type GetPointInfoResponseBody struct {
	Result  *GetPointInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetPointInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPointInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetPointInfoResponseBody) SetResult(v *GetPointInfoResponseBodyResult) *GetPointInfoResponseBody {
	s.Result = v
	return s
}

func (s *GetPointInfoResponseBody) SetSuccess(v bool) *GetPointInfoResponseBody {
	s.Success = &v
	return s
}

type GetPointInfoResponseBodyResult struct {
	UserPoints *int64 `json:"userPoints,omitempty" xml:"userPoints,omitempty"`
}

func (s GetPointInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetPointInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetPointInfoResponseBodyResult) SetUserPoints(v int64) *GetPointInfoResponseBodyResult {
	s.UserPoints = &v
	return s
}

type GetPointInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPointInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPointInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPointInfoResponse) GoString() string {
	return s.String()
}

func (s *GetPointInfoResponse) SetHeaders(v map[string]*string) *GetPointInfoResponse {
	s.Headers = v
	return s
}

func (s *GetPointInfoResponse) SetStatusCode(v int32) *GetPointInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPointInfoResponse) SetBody(v *GetPointInfoResponseBody) *GetPointInfoResponse {
	s.Body = v
	return s
}

type RevertPointHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RevertPointHeaders) String() string {
	return tea.Prettify(s)
}

func (s RevertPointHeaders) GoString() string {
	return s.String()
}

func (s *RevertPointHeaders) SetCommonHeaders(v map[string]*string) *RevertPointHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RevertPointHeaders) SetXAcsDingtalkAccessToken(v string) *RevertPointHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RevertPointRequest struct {
	Body *RevertPointRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s RevertPointRequest) String() string {
	return tea.Prettify(s)
}

func (s RevertPointRequest) GoString() string {
	return s.String()
}

func (s *RevertPointRequest) SetBody(v *RevertPointRequestBody) *RevertPointRequest {
	s.Body = v
	return s
}

type RevertPointRequestBody struct {
	PointPoolCode *string `json:"pointPoolCode,omitempty" xml:"pointPoolCode,omitempty"`
	RequestId     *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RevertPointRequestBody) String() string {
	return tea.Prettify(s)
}

func (s RevertPointRequestBody) GoString() string {
	return s.String()
}

func (s *RevertPointRequestBody) SetPointPoolCode(v string) *RevertPointRequestBody {
	s.PointPoolCode = &v
	return s
}

func (s *RevertPointRequestBody) SetRequestId(v string) *RevertPointRequestBody {
	s.RequestId = &v
	return s
}

type RevertPointShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevertPointShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RevertPointShrinkRequest) GoString() string {
	return s.String()
}

func (s *RevertPointShrinkRequest) SetBodyShrink(v string) *RevertPointShrinkRequest {
	s.BodyShrink = &v
	return s
}

type RevertPointResponseBody struct {
	Result  *RevertPointResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RevertPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevertPointResponseBody) GoString() string {
	return s.String()
}

func (s *RevertPointResponseBody) SetResult(v *RevertPointResponseBodyResult) *RevertPointResponseBody {
	s.Result = v
	return s
}

func (s *RevertPointResponseBody) SetSuccess(v bool) *RevertPointResponseBody {
	s.Success = &v
	return s
}

type RevertPointResponseBodyResult struct {
	RevertedPoints *int64 `json:"revertedPoints,omitempty" xml:"revertedPoints,omitempty"`
}

func (s RevertPointResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s RevertPointResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RevertPointResponseBodyResult) SetRevertedPoints(v int64) *RevertPointResponseBodyResult {
	s.RevertedPoints = &v
	return s
}

type RevertPointResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevertPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevertPointResponse) String() string {
	return tea.Prettify(s)
}

func (s RevertPointResponse) GoString() string {
	return s.String()
}

func (s *RevertPointResponse) SetHeaders(v map[string]*string) *RevertPointResponse {
	s.Headers = v
	return s
}

func (s *RevertPointResponse) SetStatusCode(v int32) *RevertPointResponse {
	s.StatusCode = &v
	return s
}

func (s *RevertPointResponse) SetBody(v *RevertPointResponseBody) *RevertPointResponse {
	s.Body = v
	return s
}

type SendBannerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendBannerHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendBannerHeaders) GoString() string {
	return s.String()
}

func (s *SendBannerHeaders) SetCommonHeaders(v map[string]*string) *SendBannerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendBannerHeaders) SetXAcsDingtalkAccessToken(v string) *SendBannerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendBannerRequest struct {
	Content   map[string]interface{} `json:"content,omitempty" xml:"content,omitempty"`
	EndTime   *int64                 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime *int64                 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	UserId    *string                `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SendBannerRequest) String() string {
	return tea.Prettify(s)
}

func (s SendBannerRequest) GoString() string {
	return s.String()
}

func (s *SendBannerRequest) SetContent(v map[string]interface{}) *SendBannerRequest {
	s.Content = v
	return s
}

func (s *SendBannerRequest) SetEndTime(v int64) *SendBannerRequest {
	s.EndTime = &v
	return s
}

func (s *SendBannerRequest) SetStartTime(v int64) *SendBannerRequest {
	s.StartTime = &v
	return s
}

func (s *SendBannerRequest) SetUserId(v string) *SendBannerRequest {
	s.UserId = &v
	return s
}

type SendBannerResponseBody struct {
	Arguments []interface{} `json:"arguments,omitempty" xml:"arguments,omitempty" type:"Repeated"`
	Success   *bool         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SendBannerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendBannerResponseBody) GoString() string {
	return s.String()
}

func (s *SendBannerResponseBody) SetArguments(v []interface{}) *SendBannerResponseBody {
	s.Arguments = v
	return s
}

func (s *SendBannerResponseBody) SetSuccess(v bool) *SendBannerResponseBody {
	s.Success = &v
	return s
}

type SendBannerResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendBannerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendBannerResponse) String() string {
	return tea.Prettify(s)
}

func (s SendBannerResponse) GoString() string {
	return s.String()
}

func (s *SendBannerResponse) SetHeaders(v map[string]*string) *SendBannerResponse {
	s.Headers = v
	return s
}

func (s *SendBannerResponse) SetStatusCode(v int32) *SendBannerResponse {
	s.StatusCode = &v
	return s
}

func (s *SendBannerResponse) SetBody(v *SendBannerResponseBody) *SendBannerResponse {
	s.Body = v
	return s
}

type SendPopupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendPopupHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendPopupHeaders) GoString() string {
	return s.String()
}

func (s *SendPopupHeaders) SetCommonHeaders(v map[string]*string) *SendPopupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendPopupHeaders) SetXAcsDingtalkAccessToken(v string) *SendPopupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendPopupRequest struct {
	Content   map[string]interface{} `json:"content,omitempty" xml:"content,omitempty"`
	EndTime   *int64                 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime *int64                 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	UserId    *string                `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SendPopupRequest) String() string {
	return tea.Prettify(s)
}

func (s SendPopupRequest) GoString() string {
	return s.String()
}

func (s *SendPopupRequest) SetContent(v map[string]interface{}) *SendPopupRequest {
	s.Content = v
	return s
}

func (s *SendPopupRequest) SetEndTime(v int64) *SendPopupRequest {
	s.EndTime = &v
	return s
}

func (s *SendPopupRequest) SetStartTime(v int64) *SendPopupRequest {
	s.StartTime = &v
	return s
}

func (s *SendPopupRequest) SetUserId(v string) *SendPopupRequest {
	s.UserId = &v
	return s
}

type SendPopupResponseBody struct {
	Arguments []interface{} `json:"arguments,omitempty" xml:"arguments,omitempty" type:"Repeated"`
	Success   *bool         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SendPopupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendPopupResponseBody) GoString() string {
	return s.String()
}

func (s *SendPopupResponseBody) SetArguments(v []interface{}) *SendPopupResponseBody {
	s.Arguments = v
	return s
}

func (s *SendPopupResponseBody) SetSuccess(v bool) *SendPopupResponseBody {
	s.Success = &v
	return s
}

type SendPopupResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendPopupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendPopupResponse) String() string {
	return tea.Prettify(s)
}

func (s SendPopupResponse) GoString() string {
	return s.String()
}

func (s *SendPopupResponse) SetHeaders(v map[string]*string) *SendPopupResponse {
	s.Headers = v
	return s
}

func (s *SendPopupResponse) SetStatusCode(v int32) *SendPopupResponse {
	s.StatusCode = &v
	return s
}

func (s *SendPopupResponse) SetBody(v *SendPopupResponseBody) *SendPopupResponse {
	s.Body = v
	return s
}

type SendSearchShadeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendSearchShadeHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendSearchShadeHeaders) GoString() string {
	return s.String()
}

func (s *SendSearchShadeHeaders) SetCommonHeaders(v map[string]*string) *SendSearchShadeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendSearchShadeHeaders) SetXAcsDingtalkAccessToken(v string) *SendSearchShadeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendSearchShadeRequest struct {
	Content   map[string]interface{} `json:"content,omitempty" xml:"content,omitempty"`
	EndTime   *int64                 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime *int64                 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	UserId    *string                `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SendSearchShadeRequest) String() string {
	return tea.Prettify(s)
}

func (s SendSearchShadeRequest) GoString() string {
	return s.String()
}

func (s *SendSearchShadeRequest) SetContent(v map[string]interface{}) *SendSearchShadeRequest {
	s.Content = v
	return s
}

func (s *SendSearchShadeRequest) SetEndTime(v int64) *SendSearchShadeRequest {
	s.EndTime = &v
	return s
}

func (s *SendSearchShadeRequest) SetStartTime(v int64) *SendSearchShadeRequest {
	s.StartTime = &v
	return s
}

func (s *SendSearchShadeRequest) SetUserId(v string) *SendSearchShadeRequest {
	s.UserId = &v
	return s
}

type SendSearchShadeResponseBody struct {
	Arguments []interface{} `json:"arguments,omitempty" xml:"arguments,omitempty" type:"Repeated"`
	Success   *bool         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SendSearchShadeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendSearchShadeResponseBody) GoString() string {
	return s.String()
}

func (s *SendSearchShadeResponseBody) SetArguments(v []interface{}) *SendSearchShadeResponseBody {
	s.Arguments = v
	return s
}

func (s *SendSearchShadeResponseBody) SetSuccess(v bool) *SendSearchShadeResponseBody {
	s.Success = &v
	return s
}

type SendSearchShadeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendSearchShadeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendSearchShadeResponse) String() string {
	return tea.Prettify(s)
}

func (s SendSearchShadeResponse) GoString() string {
	return s.String()
}

func (s *SendSearchShadeResponse) SetHeaders(v map[string]*string) *SendSearchShadeResponse {
	s.Headers = v
	return s
}

func (s *SendSearchShadeResponse) SetStatusCode(v int32) *SendSearchShadeResponse {
	s.StatusCode = &v
	return s
}

func (s *SendSearchShadeResponse) SetBody(v *SendSearchShadeResponseBody) *SendSearchShadeResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

func (client *Client) CheckInCrowdsByMobileWithOptions(request *CheckInCrowdsByMobileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CheckInCrowdsByMobileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CrowdIds)) {
		query["crowdIds"] = request.CrowdIds
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["mobile"] = request.Mobile
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckInCrowdsByMobile"),
		Version:     tea.String("watt_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/watt/crowdIdentifications/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckInCrowdsByMobileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckInCrowdsByMobile(request *CheckInCrowdsByMobileRequest) (_result *CheckInCrowdsByMobileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckInCrowdsByMobileResponse{}
	_body, _err := client.CheckInCrowdsByMobileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConsumePointWithOptions(tmpReq *ConsumePointRequest, headers *ConsumePointHeaders, runtime *util.RuntimeOptions) (_result *ConsumePointResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ConsumePointShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
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
		Action:      tea.String("ConsumePoint"),
		Version:     tea.String("watt_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/watt/points/consume"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ConsumePointResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConsumePoint(request *ConsumePointRequest) (_result *ConsumePointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ConsumePointHeaders{}
	_result = &ConsumePointResponse{}
	_body, _err := client.ConsumePointWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDeliveryPlanWithOptions(request *CreateDeliveryPlanRequest, headers *CreateDeliveryPlanHeaders, runtime *util.RuntimeOptions) (_result *CreateDeliveryPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ResId)) {
		body["resId"] = request.ResId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdList)) {
		body["userIdList"] = request.UserIdList
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
		Action:      tea.String("CreateDeliveryPlan"),
		Version:     tea.String("watt_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/watt/deliveryPlans/publish"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDeliveryPlanResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDeliveryPlan(request *CreateDeliveryPlanRequest) (_result *CreateDeliveryPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateDeliveryPlanHeaders{}
	_result = &CreateDeliveryPlanResponse{}
	_body, _err := client.CreateDeliveryPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPointInfoWithOptions(request *GetPointInfoRequest, headers *GetPointInfoHeaders, runtime *util.RuntimeOptions) (_result *GetPointInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PointPoolCode)) {
		query["pointPoolCode"] = request.PointPoolCode
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
		Action:      tea.String("GetPointInfo"),
		Version:     tea.String("watt_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/watt/points"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPointInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPointInfo(request *GetPointInfoRequest) (_result *GetPointInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPointInfoHeaders{}
	_result = &GetPointInfoResponse{}
	_body, _err := client.GetPointInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RevertPointWithOptions(tmpReq *RevertPointRequest, headers *RevertPointHeaders, runtime *util.RuntimeOptions) (_result *RevertPointResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RevertPointShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
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
		Action:      tea.String("RevertPoint"),
		Version:     tea.String("watt_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/watt/points/revert"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RevertPointResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevertPoint(request *RevertPointRequest) (_result *RevertPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RevertPointHeaders{}
	_result = &RevertPointResponse{}
	_body, _err := client.RevertPointWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendBannerWithOptions(request *SendBannerRequest, headers *SendBannerHeaders, runtime *util.RuntimeOptions) (_result *SendBannerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
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
		Action:      tea.String("SendBanner"),
		Version:     tea.String("watt_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/watt/banners/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendBannerResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendBanner(request *SendBannerRequest) (_result *SendBannerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendBannerHeaders{}
	_result = &SendBannerResponse{}
	_body, _err := client.SendBannerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendPopupWithOptions(request *SendPopupRequest, headers *SendPopupHeaders, runtime *util.RuntimeOptions) (_result *SendPopupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
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
		Action:      tea.String("SendPopup"),
		Version:     tea.String("watt_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/watt/popups/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendPopupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendPopup(request *SendPopupRequest) (_result *SendPopupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendPopupHeaders{}
	_result = &SendPopupResponse{}
	_body, _err := client.SendPopupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendSearchShadeWithOptions(request *SendSearchShadeRequest, headers *SendSearchShadeHeaders, runtime *util.RuntimeOptions) (_result *SendSearchShadeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
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
		Action:      tea.String("SendSearchShade"),
		Version:     tea.String("watt_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/watt/searchShades/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendSearchShadeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendSearchShade(request *SendSearchShadeRequest) (_result *SendSearchShadeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendSearchShadeHeaders{}
	_result = &SendSearchShadeResponse{}
	_body, _err := client.SendSearchShadeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
