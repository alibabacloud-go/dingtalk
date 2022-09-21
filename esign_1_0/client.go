// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package esign_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AuthUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AuthUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s AuthUrlHeaders) GoString() string {
	return s.String()
}

func (s *AuthUrlHeaders) SetCommonHeaders(v map[string]*string) *AuthUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AuthUrlHeaders) SetXAcsDingtalkAccessToken(v string) *AuthUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AuthUrlRequest struct {
	RedirectUrl *string `json:"redirectUrl,omitempty" xml:"redirectUrl,omitempty"`
}

func (s AuthUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthUrlRequest) GoString() string {
	return s.String()
}

func (s *AuthUrlRequest) SetRedirectUrl(v string) *AuthUrlRequest {
	s.RedirectUrl = &v
	return s
}

type AuthUrlResponseBody struct {
	Code    *int32                   `json:"code,omitempty" xml:"code,omitempty"`
	Data    *AuthUrlResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                  `json:"message,omitempty" xml:"message,omitempty"`
}

func (s AuthUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthUrlResponseBody) GoString() string {
	return s.String()
}

func (s *AuthUrlResponseBody) SetCode(v int32) *AuthUrlResponseBody {
	s.Code = &v
	return s
}

func (s *AuthUrlResponseBody) SetData(v *AuthUrlResponseBodyData) *AuthUrlResponseBody {
	s.Data = v
	return s
}

func (s *AuthUrlResponseBody) SetMessage(v string) *AuthUrlResponseBody {
	s.Message = &v
	return s
}

type AuthUrlResponseBodyData struct {
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	PcUrl     *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	TaskId    *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s AuthUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AuthUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *AuthUrlResponseBodyData) SetMobileUrl(v string) *AuthUrlResponseBodyData {
	s.MobileUrl = &v
	return s
}

func (s *AuthUrlResponseBodyData) SetPcUrl(v string) *AuthUrlResponseBodyData {
	s.PcUrl = &v
	return s
}

func (s *AuthUrlResponseBodyData) SetTaskId(v string) *AuthUrlResponseBodyData {
	s.TaskId = &v
	return s
}

type AuthUrlResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AuthUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AuthUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthUrlResponse) GoString() string {
	return s.String()
}

func (s *AuthUrlResponse) SetHeaders(v map[string]*string) *AuthUrlResponse {
	s.Headers = v
	return s
}

func (s *AuthUrlResponse) SetBody(v *AuthUrlResponseBody) *AuthUrlResponse {
	s.Body = v
	return s
}

type CancelCorpAuthHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
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

func (s *CancelCorpAuthHeaders) SetXAcsDingtalkAccessToken(v string) *CancelCorpAuthHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CancelCorpAuthResponseBody struct {
	Code    *int32                          `json:"code,omitempty" xml:"code,omitempty"`
	Data    *CancelCorpAuthResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                         `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CancelCorpAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelCorpAuthResponseBody) GoString() string {
	return s.String()
}

func (s *CancelCorpAuthResponseBody) SetCode(v int32) *CancelCorpAuthResponseBody {
	s.Code = &v
	return s
}

func (s *CancelCorpAuthResponseBody) SetData(v *CancelCorpAuthResponseBodyData) *CancelCorpAuthResponseBody {
	s.Data = v
	return s
}

func (s *CancelCorpAuthResponseBody) SetMessage(v string) *CancelCorpAuthResponseBody {
	s.Message = &v
	return s
}

type CancelCorpAuthResponseBodyData struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CancelCorpAuthResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CancelCorpAuthResponseBodyData) GoString() string {
	return s.String()
}

func (s *CancelCorpAuthResponseBodyData) SetResult(v bool) *CancelCorpAuthResponseBodyData {
	s.Result = &v
	return s
}

type CancelCorpAuthResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelCorpAuthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CancelCorpAuthResponse) SetBody(v *CancelCorpAuthResponseBody) *CancelCorpAuthResponse {
	s.Body = v
	return s
}

type ChannelOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ChannelOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s ChannelOrderHeaders) GoString() string {
	return s.String()
}

func (s *ChannelOrderHeaders) SetCommonHeaders(v map[string]*string) *ChannelOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ChannelOrderHeaders) SetXAcsDingtalkAccessToken(v string) *ChannelOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ChannelOrderRequest struct {
	ItemCode        *string `json:"itemCode,omitempty" xml:"itemCode,omitempty"`
	ItemName        *string `json:"itemName,omitempty" xml:"itemName,omitempty"`
	OrderCreateTime *int64  `json:"orderCreateTime,omitempty" xml:"orderCreateTime,omitempty"`
	OrderId         *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	PayFee          *int64  `json:"payFee,omitempty" xml:"payFee,omitempty"`
	Quantity        *int64  `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

func (s ChannelOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s ChannelOrderRequest) GoString() string {
	return s.String()
}

func (s *ChannelOrderRequest) SetItemCode(v string) *ChannelOrderRequest {
	s.ItemCode = &v
	return s
}

func (s *ChannelOrderRequest) SetItemName(v string) *ChannelOrderRequest {
	s.ItemName = &v
	return s
}

func (s *ChannelOrderRequest) SetOrderCreateTime(v int64) *ChannelOrderRequest {
	s.OrderCreateTime = &v
	return s
}

func (s *ChannelOrderRequest) SetOrderId(v string) *ChannelOrderRequest {
	s.OrderId = &v
	return s
}

func (s *ChannelOrderRequest) SetPayFee(v int64) *ChannelOrderRequest {
	s.PayFee = &v
	return s
}

func (s *ChannelOrderRequest) SetQuantity(v int64) *ChannelOrderRequest {
	s.Quantity = &v
	return s
}

type ChannelOrderResponseBody struct {
	Code    *int32                        `json:"code,omitempty" xml:"code,omitempty"`
	Data    *ChannelOrderResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                       `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ChannelOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChannelOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ChannelOrderResponseBody) SetCode(v int32) *ChannelOrderResponseBody {
	s.Code = &v
	return s
}

func (s *ChannelOrderResponseBody) SetData(v *ChannelOrderResponseBodyData) *ChannelOrderResponseBody {
	s.Data = v
	return s
}

func (s *ChannelOrderResponseBody) SetMessage(v string) *ChannelOrderResponseBody {
	s.Message = &v
	return s
}

type ChannelOrderResponseBodyData struct {
	EsignOrderId *string `json:"esignOrderId,omitempty" xml:"esignOrderId,omitempty"`
}

func (s ChannelOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ChannelOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChannelOrderResponseBodyData) SetEsignOrderId(v string) *ChannelOrderResponseBodyData {
	s.EsignOrderId = &v
	return s
}

type ChannelOrderResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ChannelOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChannelOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s ChannelOrderResponse) GoString() string {
	return s.String()
}

func (s *ChannelOrderResponse) SetHeaders(v map[string]*string) *ChannelOrderResponse {
	s.Headers = v
	return s
}

func (s *ChannelOrderResponse) SetBody(v *ChannelOrderResponseBody) *ChannelOrderResponse {
	s.Body = v
	return s
}

type ContractMarginHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ContractMarginHeaders) String() string {
	return tea.Prettify(s)
}

func (s ContractMarginHeaders) GoString() string {
	return s.String()
}

func (s *ContractMarginHeaders) SetCommonHeaders(v map[string]*string) *ContractMarginHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ContractMarginHeaders) SetXAcsDingtalkAccessToken(v string) *ContractMarginHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ContractMarginResponseBody struct {
	Code    *int32                          `json:"code,omitempty" xml:"code,omitempty"`
	Data    *ContractMarginResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                         `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ContractMarginResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ContractMarginResponseBody) GoString() string {
	return s.String()
}

func (s *ContractMarginResponseBody) SetCode(v int32) *ContractMarginResponseBody {
	s.Code = &v
	return s
}

func (s *ContractMarginResponseBody) SetData(v *ContractMarginResponseBodyData) *ContractMarginResponseBody {
	s.Data = v
	return s
}

func (s *ContractMarginResponseBody) SetMessage(v string) *ContractMarginResponseBody {
	s.Message = &v
	return s
}

type ContractMarginResponseBodyData struct {
	Margin *int64 `json:"margin,omitempty" xml:"margin,omitempty"`
}

func (s ContractMarginResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ContractMarginResponseBodyData) GoString() string {
	return s.String()
}

func (s *ContractMarginResponseBodyData) SetMargin(v int64) *ContractMarginResponseBodyData {
	s.Margin = &v
	return s
}

type ContractMarginResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ContractMarginResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ContractMarginResponse) String() string {
	return tea.Prettify(s)
}

func (s ContractMarginResponse) GoString() string {
	return s.String()
}

func (s *ContractMarginResponse) SetHeaders(v map[string]*string) *ContractMarginResponse {
	s.Headers = v
	return s
}

func (s *ContractMarginResponse) SetBody(v *ContractMarginResponseBody) *ContractMarginResponse {
	s.Body = v
	return s
}

type CorpConsoleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CorpConsoleHeaders) String() string {
	return tea.Prettify(s)
}

func (s CorpConsoleHeaders) GoString() string {
	return s.String()
}

func (s *CorpConsoleHeaders) SetCommonHeaders(v map[string]*string) *CorpConsoleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CorpConsoleHeaders) SetXAcsDingtalkAccessToken(v string) *CorpConsoleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CorpConsoleResponseBody struct {
	Code    *int32                       `json:"code,omitempty" xml:"code,omitempty"`
	Data    *CorpConsoleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                      `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CorpConsoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CorpConsoleResponseBody) GoString() string {
	return s.String()
}

func (s *CorpConsoleResponseBody) SetCode(v int32) *CorpConsoleResponseBody {
	s.Code = &v
	return s
}

func (s *CorpConsoleResponseBody) SetData(v *CorpConsoleResponseBodyData) *CorpConsoleResponseBody {
	s.Data = v
	return s
}

func (s *CorpConsoleResponseBody) SetMessage(v string) *CorpConsoleResponseBody {
	s.Message = &v
	return s
}

type CorpConsoleResponseBodyData struct {
	OrgConsoleUrl *int64 `json:"orgConsoleUrl,omitempty" xml:"orgConsoleUrl,omitempty"`
}

func (s CorpConsoleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CorpConsoleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CorpConsoleResponseBodyData) SetOrgConsoleUrl(v int64) *CorpConsoleResponseBodyData {
	s.OrgConsoleUrl = &v
	return s
}

type CorpConsoleResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CorpConsoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CorpConsoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CorpConsoleResponse) GoString() string {
	return s.String()
}

func (s *CorpConsoleResponse) SetHeaders(v map[string]*string) *CorpConsoleResponse {
	s.Headers = v
	return s
}

func (s *CorpConsoleResponse) SetBody(v *CorpConsoleResponseBody) *CorpConsoleResponse {
	s.Body = v
	return s
}

type CorpInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CorpInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s CorpInfoHeaders) GoString() string {
	return s.String()
}

func (s *CorpInfoHeaders) SetCommonHeaders(v map[string]*string) *CorpInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CorpInfoHeaders) SetXAcsDingtalkAccessToken(v string) *CorpInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CorpInfoResponseBody struct {
	Code    *int32                    `json:"code,omitempty" xml:"code,omitempty"`
	Data    *CorpInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                   `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CorpInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CorpInfoResponseBody) GoString() string {
	return s.String()
}

func (s *CorpInfoResponseBody) SetCode(v int32) *CorpInfoResponseBody {
	s.Code = &v
	return s
}

func (s *CorpInfoResponseBody) SetData(v *CorpInfoResponseBodyData) *CorpInfoResponseBody {
	s.Data = v
	return s
}

func (s *CorpInfoResponseBody) SetMessage(v string) *CorpInfoResponseBody {
	s.Message = &v
	return s
}

type CorpInfoResponseBodyData struct {
	OrgRealName *string `json:"orgRealName,omitempty" xml:"orgRealName,omitempty"`
	RealName    *bool   `json:"realName,omitempty" xml:"realName,omitempty"`
}

func (s CorpInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CorpInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *CorpInfoResponseBodyData) SetOrgRealName(v string) *CorpInfoResponseBodyData {
	s.OrgRealName = &v
	return s
}

func (s *CorpInfoResponseBodyData) SetRealName(v bool) *CorpInfoResponseBodyData {
	s.RealName = &v
	return s
}

type CorpInfoResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CorpInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CorpInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s CorpInfoResponse) GoString() string {
	return s.String()
}

func (s *CorpInfoResponse) SetHeaders(v map[string]*string) *CorpInfoResponse {
	s.Headers = v
	return s
}

func (s *CorpInfoResponse) SetBody(v *CorpInfoResponseBody) *CorpInfoResponse {
	s.Body = v
	return s
}

type CreateDeveloperHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateDeveloperHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateDeveloperHeaders) GoString() string {
	return s.String()
}

func (s *CreateDeveloperHeaders) SetCommonHeaders(v map[string]*string) *CreateDeveloperHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateDeveloperHeaders) SetXAcsDingtalkAccessToken(v string) *CreateDeveloperHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateDeveloperRequest struct {
	RedirectUrl *string `json:"redirectUrl,omitempty" xml:"redirectUrl,omitempty"`
}

func (s CreateDeveloperRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeveloperRequest) GoString() string {
	return s.String()
}

func (s *CreateDeveloperRequest) SetRedirectUrl(v string) *CreateDeveloperRequest {
	s.RedirectUrl = &v
	return s
}

type CreateDeveloperResponseBody struct {
	Code    *int32  `json:"code,omitempty" xml:"code,omitempty"`
	Data    *bool   `json:"data,omitempty" xml:"data,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CreateDeveloperResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeveloperResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeveloperResponseBody) SetCode(v int32) *CreateDeveloperResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDeveloperResponseBody) SetData(v bool) *CreateDeveloperResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDeveloperResponseBody) SetMessage(v string) *CreateDeveloperResponseBody {
	s.Message = &v
	return s
}

type CreateDeveloperResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDeveloperResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDeveloperResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeveloperResponse) GoString() string {
	return s.String()
}

func (s *CreateDeveloperResponse) SetHeaders(v map[string]*string) *CreateDeveloperResponse {
	s.Headers = v
	return s
}

func (s *CreateDeveloperResponse) SetBody(v *CreateDeveloperResponseBody) *CreateDeveloperResponse {
	s.Body = v
	return s
}

type GetCorpRealnameUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCorpRealnameUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCorpRealnameUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetCorpRealnameUrlHeaders) SetCommonHeaders(v map[string]*string) *GetCorpRealnameUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCorpRealnameUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GetCorpRealnameUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCorpRealnameUrlRequest struct {
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetCorpRealnameUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCorpRealnameUrlRequest) GoString() string {
	return s.String()
}

func (s *GetCorpRealnameUrlRequest) SetUserId(v string) *GetCorpRealnameUrlRequest {
	s.UserId = &v
	return s
}

type GetCorpRealnameUrlResponseBody struct {
	Code    *int32                              `json:"code,omitempty" xml:"code,omitempty"`
	Data    *GetCorpRealnameUrlResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                             `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetCorpRealnameUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCorpRealnameUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetCorpRealnameUrlResponseBody) SetCode(v int32) *GetCorpRealnameUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetCorpRealnameUrlResponseBody) SetData(v *GetCorpRealnameUrlResponseBodyData) *GetCorpRealnameUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetCorpRealnameUrlResponseBody) SetMessage(v string) *GetCorpRealnameUrlResponseBody {
	s.Message = &v
	return s
}

type GetCorpRealnameUrlResponseBodyData struct {
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	PcUrl     *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	TaskId    *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetCorpRealnameUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCorpRealnameUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCorpRealnameUrlResponseBodyData) SetMobileUrl(v string) *GetCorpRealnameUrlResponseBodyData {
	s.MobileUrl = &v
	return s
}

func (s *GetCorpRealnameUrlResponseBodyData) SetPcUrl(v string) *GetCorpRealnameUrlResponseBodyData {
	s.PcUrl = &v
	return s
}

func (s *GetCorpRealnameUrlResponseBodyData) SetTaskId(v string) *GetCorpRealnameUrlResponseBodyData {
	s.TaskId = &v
	return s
}

type GetCorpRealnameUrlResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCorpRealnameUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCorpRealnameUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCorpRealnameUrlResponse) GoString() string {
	return s.String()
}

func (s *GetCorpRealnameUrlResponse) SetHeaders(v map[string]*string) *GetCorpRealnameUrlResponse {
	s.Headers = v
	return s
}

func (s *GetCorpRealnameUrlResponse) SetBody(v *GetCorpRealnameUrlResponseBody) *GetCorpRealnameUrlResponse {
	s.Body = v
	return s
}

type GetCropStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCropStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCropStatusHeaders) GoString() string {
	return s.String()
}

func (s *GetCropStatusHeaders) SetCommonHeaders(v map[string]*string) *GetCropStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCropStatusHeaders) SetXAcsDingtalkAccessToken(v string) *GetCropStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCropStatusResponseBody struct {
	Code    *int32                         `json:"code,omitempty" xml:"code,omitempty"`
	Data    *GetCropStatusResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                        `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetCropStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCropStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetCropStatusResponseBody) SetCode(v int32) *GetCropStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetCropStatusResponseBody) SetData(v *GetCropStatusResponseBodyData) *GetCropStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetCropStatusResponseBody) SetMessage(v string) *GetCropStatusResponseBody {
	s.Message = &v
	return s
}

type GetCropStatusResponseBodyData struct {
	AuthStatus    *string `json:"authStatus,omitempty" xml:"authStatus,omitempty"`
	InstallStatus *string `json:"installStatus,omitempty" xml:"installStatus,omitempty"`
}

func (s GetCropStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCropStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCropStatusResponseBodyData) SetAuthStatus(v string) *GetCropStatusResponseBodyData {
	s.AuthStatus = &v
	return s
}

func (s *GetCropStatusResponseBodyData) SetInstallStatus(v string) *GetCropStatusResponseBodyData {
	s.InstallStatus = &v
	return s
}

type GetCropStatusResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCropStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCropStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCropStatusResponse) GoString() string {
	return s.String()
}

func (s *GetCropStatusResponse) SetHeaders(v map[string]*string) *GetCropStatusResponse {
	s.Headers = v
	return s
}

func (s *GetCropStatusResponse) SetBody(v *GetCropStatusResponseBody) *GetCropStatusResponse {
	s.Body = v
	return s
}

type GetFileHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFileHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFileHeaders) GoString() string {
	return s.String()
}

func (s *GetFileHeaders) SetCommonHeaders(v map[string]*string) *GetFileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFileHeaders) SetXAcsDingtalkAccessToken(v string) *GetFileHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFileResponseBody struct {
	Code    *int32                   `json:"code,omitempty" xml:"code,omitempty"`
	Data    *GetFileResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                  `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileResponseBody) SetCode(v int32) *GetFileResponseBody {
	s.Code = &v
	return s
}

func (s *GetFileResponseBody) SetData(v *GetFileResponseBodyData) *GetFileResponseBody {
	s.Data = v
	return s
}

func (s *GetFileResponseBody) SetMessage(v string) *GetFileResponseBody {
	s.Message = &v
	return s
}

type GetFileResponseBodyData struct {
	DownloadUrl   *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	FileId        *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	PdfTotalPages *int32  `json:"pdfTotalPages,omitempty" xml:"pdfTotalPages,omitempty"`
	Size          *int64  `json:"size,omitempty" xml:"size,omitempty"`
	Status        *int32  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFileResponseBodyData) SetDownloadUrl(v string) *GetFileResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *GetFileResponseBodyData) SetFileId(v string) *GetFileResponseBodyData {
	s.FileId = &v
	return s
}

func (s *GetFileResponseBodyData) SetName(v string) *GetFileResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetFileResponseBodyData) SetPdfTotalPages(v int32) *GetFileResponseBodyData {
	s.PdfTotalPages = &v
	return s
}

func (s *GetFileResponseBodyData) SetSize(v int64) *GetFileResponseBodyData {
	s.Size = &v
	return s
}

func (s *GetFileResponseBodyData) SetStatus(v int32) *GetFileResponseBodyData {
	s.Status = &v
	return s
}

type GetFileResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFileResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFileResponse) GoString() string {
	return s.String()
}

func (s *GetFileResponse) SetHeaders(v map[string]*string) *GetFileResponse {
	s.Headers = v
	return s
}

func (s *GetFileResponse) SetBody(v *GetFileResponseBody) *GetFileResponse {
	s.Body = v
	return s
}

type GetFlowDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
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

func (s *GetFlowDetailHeaders) SetXAcsDingtalkAccessToken(v string) *GetFlowDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFlowDetailRequest struct {
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetFlowDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFlowDetailRequest) GoString() string {
	return s.String()
}

func (s *GetFlowDetailRequest) SetTaskId(v string) *GetFlowDetailRequest {
	s.TaskId = &v
	return s
}

type GetFlowDetailResponseBody struct {
	Code    *int32                         `json:"code,omitempty" xml:"code,omitempty"`
	Data    *GetFlowDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                        `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetFlowDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFlowDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetFlowDetailResponseBody) SetCode(v int32) *GetFlowDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetFlowDetailResponseBody) SetData(v *GetFlowDetailResponseBodyData) *GetFlowDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetFlowDetailResponseBody) SetMessage(v string) *GetFlowDetailResponseBody {
	s.Message = &v
	return s
}

type GetFlowDetailResponseBodyData struct {
	BusinessSense           *string                              `json:"businessSense,omitempty" xml:"businessSense,omitempty"`
	FlowStatus              *int32                               `json:"flowStatus,omitempty" xml:"flowStatus,omitempty"`
	InitiatorAuthorizedName *string                              `json:"initiatorAuthorizedName,omitempty" xml:"initiatorAuthorizedName,omitempty"`
	InitiatorName           *string                              `json:"initiatorName,omitempty" xml:"initiatorName,omitempty"`
	Logs                    []*GetFlowDetailResponseBodyDataLogs `json:"logs,omitempty" xml:"logs,omitempty" type:"Repeated"`
}

func (s GetFlowDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFlowDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFlowDetailResponseBodyData) SetBusinessSense(v string) *GetFlowDetailResponseBodyData {
	s.BusinessSense = &v
	return s
}

func (s *GetFlowDetailResponseBodyData) SetFlowStatus(v int32) *GetFlowDetailResponseBodyData {
	s.FlowStatus = &v
	return s
}

func (s *GetFlowDetailResponseBodyData) SetInitiatorAuthorizedName(v string) *GetFlowDetailResponseBodyData {
	s.InitiatorAuthorizedName = &v
	return s
}

func (s *GetFlowDetailResponseBodyData) SetInitiatorName(v string) *GetFlowDetailResponseBodyData {
	s.InitiatorName = &v
	return s
}

func (s *GetFlowDetailResponseBodyData) SetLogs(v []*GetFlowDetailResponseBodyDataLogs) *GetFlowDetailResponseBodyData {
	s.Logs = v
	return s
}

type GetFlowDetailResponseBodyDataLogs struct {
	LogType             *string `json:"logType,omitempty" xml:"logType,omitempty"`
	OperateDescription  *string `json:"operateDescription,omitempty" xml:"operateDescription,omitempty"`
	OperateTime         *int64  `json:"operateTime,omitempty" xml:"operateTime,omitempty"`
	OperatorAccountName *string `json:"operatorAccountName,omitempty" xml:"operatorAccountName,omitempty"`
}

func (s GetFlowDetailResponseBodyDataLogs) String() string {
	return tea.Prettify(s)
}

func (s GetFlowDetailResponseBodyDataLogs) GoString() string {
	return s.String()
}

func (s *GetFlowDetailResponseBodyDataLogs) SetLogType(v string) *GetFlowDetailResponseBodyDataLogs {
	s.LogType = &v
	return s
}

func (s *GetFlowDetailResponseBodyDataLogs) SetOperateDescription(v string) *GetFlowDetailResponseBodyDataLogs {
	s.OperateDescription = &v
	return s
}

func (s *GetFlowDetailResponseBodyDataLogs) SetOperateTime(v int64) *GetFlowDetailResponseBodyDataLogs {
	s.OperateTime = &v
	return s
}

func (s *GetFlowDetailResponseBodyDataLogs) SetOperatorAccountName(v string) *GetFlowDetailResponseBodyDataLogs {
	s.OperatorAccountName = &v
	return s
}

type GetFlowDetailResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFlowDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetFlowDetailResponse) SetBody(v *GetFlowDetailResponseBody) *GetFlowDetailResponse {
	s.Body = v
	return s
}

type GetFlowSignDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFlowSignDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFlowSignDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetFlowSignDetailHeaders) SetCommonHeaders(v map[string]*string) *GetFlowSignDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFlowSignDetailHeaders) SetXAcsDingtalkAccessToken(v string) *GetFlowSignDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFlowSignDetailRequest struct {
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetFlowSignDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFlowSignDetailRequest) GoString() string {
	return s.String()
}

func (s *GetFlowSignDetailRequest) SetTaskId(v string) *GetFlowSignDetailRequest {
	s.TaskId = &v
	return s
}

type GetFlowSignDetailResponseBody struct {
	Code    *int32                             `json:"code,omitempty" xml:"code,omitempty"`
	Data    *GetFlowSignDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                            `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetFlowSignDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFlowSignDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetFlowSignDetailResponseBody) SetCode(v int32) *GetFlowSignDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetFlowSignDetailResponseBody) SetData(v *GetFlowSignDetailResponseBodyData) *GetFlowSignDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetFlowSignDetailResponseBody) SetMessage(v string) *GetFlowSignDetailResponseBody {
	s.Message = &v
	return s
}

type GetFlowSignDetailResponseBodyData struct {
	BusinessSense *string                                     `json:"businessSense,omitempty" xml:"businessSense,omitempty"`
	FlowStatus    *int32                                      `json:"flowStatus,omitempty" xml:"flowStatus,omitempty"`
	Signers       []*GetFlowSignDetailResponseBodyDataSigners `json:"signers,omitempty" xml:"signers,omitempty" type:"Repeated"`
}

func (s GetFlowSignDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFlowSignDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFlowSignDetailResponseBodyData) SetBusinessSense(v string) *GetFlowSignDetailResponseBodyData {
	s.BusinessSense = &v
	return s
}

func (s *GetFlowSignDetailResponseBodyData) SetFlowStatus(v int32) *GetFlowSignDetailResponseBodyData {
	s.FlowStatus = &v
	return s
}

func (s *GetFlowSignDetailResponseBodyData) SetSigners(v []*GetFlowSignDetailResponseBodyDataSigners) *GetFlowSignDetailResponseBodyData {
	s.Signers = v
	return s
}

type GetFlowSignDetailResponseBodyDataSigners struct {
	SignStatus *int32  `json:"signStatus,omitempty" xml:"signStatus,omitempty"`
	SignerName *string `json:"signerName,omitempty" xml:"signerName,omitempty"`
}

func (s GetFlowSignDetailResponseBodyDataSigners) String() string {
	return tea.Prettify(s)
}

func (s GetFlowSignDetailResponseBodyDataSigners) GoString() string {
	return s.String()
}

func (s *GetFlowSignDetailResponseBodyDataSigners) SetSignStatus(v int32) *GetFlowSignDetailResponseBodyDataSigners {
	s.SignStatus = &v
	return s
}

func (s *GetFlowSignDetailResponseBodyDataSigners) SetSignerName(v string) *GetFlowSignDetailResponseBodyDataSigners {
	s.SignerName = &v
	return s
}

type GetFlowSignDetailResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFlowSignDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFlowSignDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFlowSignDetailResponse) GoString() string {
	return s.String()
}

func (s *GetFlowSignDetailResponse) SetHeaders(v map[string]*string) *GetFlowSignDetailResponse {
	s.Headers = v
	return s
}

func (s *GetFlowSignDetailResponse) SetBody(v *GetFlowSignDetailResponseBody) *GetFlowSignDetailResponse {
	s.Body = v
	return s
}

type GetProcessStartUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetProcessStartUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetProcessStartUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetProcessStartUrlHeaders) SetCommonHeaders(v map[string]*string) *GetProcessStartUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetProcessStartUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GetProcessStartUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetProcessStartUrlRequest struct {
	Ccs             []*GetProcessStartUrlRequestCcs          `json:"ccs,omitempty" xml:"ccs,omitempty" type:"Repeated"`
	Files           []*GetProcessStartUrlRequestFiles        `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	InitiatorUserId *string                                  `json:"initiatorUserId,omitempty" xml:"initiatorUserId,omitempty"`
	Participants    []*GetProcessStartUrlRequestParticipants `json:"participants,omitempty" xml:"participants,omitempty" type:"Repeated"`
	RedirectUrl     *string                                  `json:"redirectUrl,omitempty" xml:"redirectUrl,omitempty"`
	SourceInfo      *GetProcessStartUrlRequestSourceInfo     `json:"sourceInfo,omitempty" xml:"sourceInfo,omitempty" type:"Struct"`
	TaskName        *string                                  `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s GetProcessStartUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProcessStartUrlRequest) GoString() string {
	return s.String()
}

func (s *GetProcessStartUrlRequest) SetCcs(v []*GetProcessStartUrlRequestCcs) *GetProcessStartUrlRequest {
	s.Ccs = v
	return s
}

func (s *GetProcessStartUrlRequest) SetFiles(v []*GetProcessStartUrlRequestFiles) *GetProcessStartUrlRequest {
	s.Files = v
	return s
}

func (s *GetProcessStartUrlRequest) SetInitiatorUserId(v string) *GetProcessStartUrlRequest {
	s.InitiatorUserId = &v
	return s
}

func (s *GetProcessStartUrlRequest) SetParticipants(v []*GetProcessStartUrlRequestParticipants) *GetProcessStartUrlRequest {
	s.Participants = v
	return s
}

func (s *GetProcessStartUrlRequest) SetRedirectUrl(v string) *GetProcessStartUrlRequest {
	s.RedirectUrl = &v
	return s
}

func (s *GetProcessStartUrlRequest) SetSourceInfo(v *GetProcessStartUrlRequestSourceInfo) *GetProcessStartUrlRequest {
	s.SourceInfo = v
	return s
}

func (s *GetProcessStartUrlRequest) SetTaskName(v string) *GetProcessStartUrlRequest {
	s.TaskName = &v
	return s
}

type GetProcessStartUrlRequestCcs struct {
	Account     *string `json:"account,omitempty" xml:"account,omitempty"`
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	OrgName     *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetProcessStartUrlRequestCcs) String() string {
	return tea.Prettify(s)
}

func (s GetProcessStartUrlRequestCcs) GoString() string {
	return s.String()
}

func (s *GetProcessStartUrlRequestCcs) SetAccount(v string) *GetProcessStartUrlRequestCcs {
	s.Account = &v
	return s
}

func (s *GetProcessStartUrlRequestCcs) SetAccountName(v string) *GetProcessStartUrlRequestCcs {
	s.AccountName = &v
	return s
}

func (s *GetProcessStartUrlRequestCcs) SetAccountType(v string) *GetProcessStartUrlRequestCcs {
	s.AccountType = &v
	return s
}

func (s *GetProcessStartUrlRequestCcs) SetOrgName(v string) *GetProcessStartUrlRequestCcs {
	s.OrgName = &v
	return s
}

func (s *GetProcessStartUrlRequestCcs) SetUserId(v string) *GetProcessStartUrlRequestCcs {
	s.UserId = &v
	return s
}

type GetProcessStartUrlRequestFiles struct {
	FileId   *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s GetProcessStartUrlRequestFiles) String() string {
	return tea.Prettify(s)
}

func (s GetProcessStartUrlRequestFiles) GoString() string {
	return s.String()
}

func (s *GetProcessStartUrlRequestFiles) SetFileId(v string) *GetProcessStartUrlRequestFiles {
	s.FileId = &v
	return s
}

func (s *GetProcessStartUrlRequestFiles) SetFileName(v string) *GetProcessStartUrlRequestFiles {
	s.FileName = &v
	return s
}

type GetProcessStartUrlRequestParticipants struct {
	Account          *string `json:"account,omitempty" xml:"account,omitempty"`
	AccountName      *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	AccountType      *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	OrgName          *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	SignRequirements *string `json:"signRequirements,omitempty" xml:"signRequirements,omitempty"`
	UserId           *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetProcessStartUrlRequestParticipants) String() string {
	return tea.Prettify(s)
}

func (s GetProcessStartUrlRequestParticipants) GoString() string {
	return s.String()
}

func (s *GetProcessStartUrlRequestParticipants) SetAccount(v string) *GetProcessStartUrlRequestParticipants {
	s.Account = &v
	return s
}

func (s *GetProcessStartUrlRequestParticipants) SetAccountName(v string) *GetProcessStartUrlRequestParticipants {
	s.AccountName = &v
	return s
}

func (s *GetProcessStartUrlRequestParticipants) SetAccountType(v string) *GetProcessStartUrlRequestParticipants {
	s.AccountType = &v
	return s
}

func (s *GetProcessStartUrlRequestParticipants) SetOrgName(v string) *GetProcessStartUrlRequestParticipants {
	s.OrgName = &v
	return s
}

func (s *GetProcessStartUrlRequestParticipants) SetSignRequirements(v string) *GetProcessStartUrlRequestParticipants {
	s.SignRequirements = &v
	return s
}

func (s *GetProcessStartUrlRequestParticipants) SetUserId(v string) *GetProcessStartUrlRequestParticipants {
	s.UserId = &v
	return s
}

type GetProcessStartUrlRequestSourceInfo struct {
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	PcUrl     *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	ShowText  *string `json:"showText,omitempty" xml:"showText,omitempty"`
}

func (s GetProcessStartUrlRequestSourceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetProcessStartUrlRequestSourceInfo) GoString() string {
	return s.String()
}

func (s *GetProcessStartUrlRequestSourceInfo) SetMobileUrl(v string) *GetProcessStartUrlRequestSourceInfo {
	s.MobileUrl = &v
	return s
}

func (s *GetProcessStartUrlRequestSourceInfo) SetPcUrl(v string) *GetProcessStartUrlRequestSourceInfo {
	s.PcUrl = &v
	return s
}

func (s *GetProcessStartUrlRequestSourceInfo) SetShowText(v string) *GetProcessStartUrlRequestSourceInfo {
	s.ShowText = &v
	return s
}

type GetProcessStartUrlResponseBody struct {
	Code    *int32                              `json:"code,omitempty" xml:"code,omitempty"`
	Data    *GetProcessStartUrlResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                             `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetProcessStartUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProcessStartUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetProcessStartUrlResponseBody) SetCode(v int32) *GetProcessStartUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetProcessStartUrlResponseBody) SetData(v *GetProcessStartUrlResponseBodyData) *GetProcessStartUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetProcessStartUrlResponseBody) SetMessage(v string) *GetProcessStartUrlResponseBody {
	s.Message = &v
	return s
}

type GetProcessStartUrlResponseBodyData struct {
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	PcUrl     *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	TaskId    *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetProcessStartUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProcessStartUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProcessStartUrlResponseBodyData) SetMobileUrl(v string) *GetProcessStartUrlResponseBodyData {
	s.MobileUrl = &v
	return s
}

func (s *GetProcessStartUrlResponseBodyData) SetPcUrl(v string) *GetProcessStartUrlResponseBodyData {
	s.PcUrl = &v
	return s
}

func (s *GetProcessStartUrlResponseBodyData) SetTaskId(v string) *GetProcessStartUrlResponseBodyData {
	s.TaskId = &v
	return s
}

type GetProcessStartUrlResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProcessStartUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProcessStartUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProcessStartUrlResponse) GoString() string {
	return s.String()
}

func (s *GetProcessStartUrlResponse) SetHeaders(v map[string]*string) *GetProcessStartUrlResponse {
	s.Headers = v
	return s
}

func (s *GetProcessStartUrlResponse) SetBody(v *GetProcessStartUrlResponseBody) *GetProcessStartUrlResponse {
	s.Body = v
	return s
}

type GetSignNoticeUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSignNoticeUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSignNoticeUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetSignNoticeUrlHeaders) SetCommonHeaders(v map[string]*string) *GetSignNoticeUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSignNoticeUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GetSignNoticeUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSignNoticeUrlRequest struct {
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetSignNoticeUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSignNoticeUrlRequest) GoString() string {
	return s.String()
}

func (s *GetSignNoticeUrlRequest) SetTaskId(v string) *GetSignNoticeUrlRequest {
	s.TaskId = &v
	return s
}

type GetSignNoticeUrlResponseBody struct {
	// 返回错误码
	Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
	// 返回数据
	Data *GetSignNoticeUrlResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 返回结果信息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetSignNoticeUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSignNoticeUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetSignNoticeUrlResponseBody) SetCode(v int32) *GetSignNoticeUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetSignNoticeUrlResponseBody) SetData(v *GetSignNoticeUrlResponseBodyData) *GetSignNoticeUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetSignNoticeUrlResponseBody) SetMessage(v string) *GetSignNoticeUrlResponseBody {
	s.Message = &v
	return s
}

type GetSignNoticeUrlResponseBodyData struct {
	// 移动端URL
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	// PC端URL
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
}

func (s GetSignNoticeUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSignNoticeUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSignNoticeUrlResponseBodyData) SetMobileUrl(v string) *GetSignNoticeUrlResponseBodyData {
	s.MobileUrl = &v
	return s
}

func (s *GetSignNoticeUrlResponseBodyData) SetPcUrl(v string) *GetSignNoticeUrlResponseBodyData {
	s.PcUrl = &v
	return s
}

type GetSignNoticeUrlResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSignNoticeUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSignNoticeUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSignNoticeUrlResponse) GoString() string {
	return s.String()
}

func (s *GetSignNoticeUrlResponse) SetHeaders(v map[string]*string) *GetSignNoticeUrlResponse {
	s.Headers = v
	return s
}

func (s *GetSignNoticeUrlResponse) SetBody(v *GetSignNoticeUrlResponseBody) *GetSignNoticeUrlResponse {
	s.Body = v
	return s
}

type GetUploadUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUploadUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUploadUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetUploadUrlHeaders) SetCommonHeaders(v map[string]*string) *GetUploadUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUploadUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GetUploadUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUploadUrlRequest struct {
	ContentMd5  *string `json:"contentMd5,omitempty" xml:"contentMd5,omitempty"`
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	Convert2Pdf *bool   `json:"convert2Pdf,omitempty" xml:"convert2Pdf,omitempty"`
	FileName    *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileSize    *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
}

func (s GetUploadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUploadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetUploadUrlRequest) SetContentMd5(v string) *GetUploadUrlRequest {
	s.ContentMd5 = &v
	return s
}

func (s *GetUploadUrlRequest) SetContentType(v string) *GetUploadUrlRequest {
	s.ContentType = &v
	return s
}

func (s *GetUploadUrlRequest) SetConvert2Pdf(v bool) *GetUploadUrlRequest {
	s.Convert2Pdf = &v
	return s
}

func (s *GetUploadUrlRequest) SetFileName(v string) *GetUploadUrlRequest {
	s.FileName = &v
	return s
}

func (s *GetUploadUrlRequest) SetFileSize(v int64) *GetUploadUrlRequest {
	s.FileSize = &v
	return s
}

type GetUploadUrlResponseBody struct {
	Code    *int32                        `json:"code,omitempty" xml:"code,omitempty"`
	Data    *GetUploadUrlResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                       `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetUploadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUploadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadUrlResponseBody) SetCode(v int32) *GetUploadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetUploadUrlResponseBody) SetData(v *GetUploadUrlResponseBodyData) *GetUploadUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetUploadUrlResponseBody) SetMessage(v string) *GetUploadUrlResponseBody {
	s.Message = &v
	return s
}

type GetUploadUrlResponseBodyData struct {
	FileId    *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	UploadUrl *string `json:"uploadUrl,omitempty" xml:"uploadUrl,omitempty"`
}

func (s GetUploadUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUploadUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUploadUrlResponseBodyData) SetFileId(v string) *GetUploadUrlResponseBodyData {
	s.FileId = &v
	return s
}

func (s *GetUploadUrlResponseBodyData) SetUploadUrl(v string) *GetUploadUrlResponseBodyData {
	s.UploadUrl = &v
	return s
}

type GetUploadUrlResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUploadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUploadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUploadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetUploadUrlResponse) SetHeaders(v map[string]*string) *GetUploadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetUploadUrlResponse) SetBody(v *GetUploadUrlResponseBody) *GetUploadUrlResponse {
	s.Body = v
	return s
}

type GetUserInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
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

func (s *GetUserInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserInfoResponseBody struct {
	Code    *int32                       `json:"code,omitempty" xml:"code,omitempty"`
	Data    *GetUserInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                      `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetUserInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserInfoResponseBody) SetCode(v int32) *GetUserInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserInfoResponseBody) SetData(v *GetUserInfoResponseBodyData) *GetUserInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetUserInfoResponseBody) SetMessage(v string) *GetUserInfoResponseBody {
	s.Message = &v
	return s
}

type GetUserInfoResponseBodyData struct {
	RealName     *bool   `json:"realName,omitempty" xml:"realName,omitempty"`
	UserRealName *string `json:"userRealName,omitempty" xml:"userRealName,omitempty"`
}

func (s GetUserInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUserInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserInfoResponseBodyData) SetRealName(v bool) *GetUserInfoResponseBodyData {
	s.RealName = &v
	return s
}

func (s *GetUserInfoResponseBodyData) SetUserRealName(v string) *GetUserInfoResponseBodyData {
	s.UserRealName = &v
	return s
}

type GetUserInfoResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetUserInfoResponse) SetBody(v *GetUserInfoResponseBody) *GetUserInfoResponse {
	s.Body = v
	return s
}

type GetUserRealnameUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserRealnameUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserRealnameUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetUserRealnameUrlHeaders) SetCommonHeaders(v map[string]*string) *GetUserRealnameUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserRealnameUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserRealnameUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserRealnameUrlRequest struct {
	RedirectUrl *string `json:"redirectUrl,omitempty" xml:"redirectUrl,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserRealnameUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserRealnameUrlRequest) GoString() string {
	return s.String()
}

func (s *GetUserRealnameUrlRequest) SetRedirectUrl(v string) *GetUserRealnameUrlRequest {
	s.RedirectUrl = &v
	return s
}

func (s *GetUserRealnameUrlRequest) SetUserId(v string) *GetUserRealnameUrlRequest {
	s.UserId = &v
	return s
}

type GetUserRealnameUrlResponseBody struct {
	Code    *int32                              `json:"code,omitempty" xml:"code,omitempty"`
	Data    *GetUserRealnameUrlResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                             `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetUserRealnameUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserRealnameUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserRealnameUrlResponseBody) SetCode(v int32) *GetUserRealnameUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserRealnameUrlResponseBody) SetData(v *GetUserRealnameUrlResponseBodyData) *GetUserRealnameUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetUserRealnameUrlResponseBody) SetMessage(v string) *GetUserRealnameUrlResponseBody {
	s.Message = &v
	return s
}

type GetUserRealnameUrlResponseBodyData struct {
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	PcUrl     *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	TaskId    *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetUserRealnameUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUserRealnameUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserRealnameUrlResponseBodyData) SetMobileUrl(v string) *GetUserRealnameUrlResponseBodyData {
	s.MobileUrl = &v
	return s
}

func (s *GetUserRealnameUrlResponseBodyData) SetPcUrl(v string) *GetUserRealnameUrlResponseBodyData {
	s.PcUrl = &v
	return s
}

func (s *GetUserRealnameUrlResponseBodyData) SetTaskId(v string) *GetUserRealnameUrlResponseBodyData {
	s.TaskId = &v
	return s
}

type GetUserRealnameUrlResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserRealnameUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserRealnameUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserRealnameUrlResponse) GoString() string {
	return s.String()
}

func (s *GetUserRealnameUrlResponse) SetHeaders(v map[string]*string) *GetUserRealnameUrlResponse {
	s.Headers = v
	return s
}

func (s *GetUserRealnameUrlResponse) SetBody(v *GetUserRealnameUrlResponseBody) *GetUserRealnameUrlResponse {
	s.Body = v
	return s
}

type ListFlowDocsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListFlowDocsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListFlowDocsHeaders) GoString() string {
	return s.String()
}

func (s *ListFlowDocsHeaders) SetCommonHeaders(v map[string]*string) *ListFlowDocsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListFlowDocsHeaders) SetXAcsDingtalkAccessToken(v string) *ListFlowDocsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListFlowDocsRequest struct {
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListFlowDocsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowDocsRequest) GoString() string {
	return s.String()
}

func (s *ListFlowDocsRequest) SetTaskId(v string) *ListFlowDocsRequest {
	s.TaskId = &v
	return s
}

type ListFlowDocsResponseBody struct {
	Code    *int32                          `json:"code,omitempty" xml:"code,omitempty"`
	Data    []*ListFlowDocsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Message *string                         `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ListFlowDocsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowDocsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowDocsResponseBody) SetCode(v int32) *ListFlowDocsResponseBody {
	s.Code = &v
	return s
}

func (s *ListFlowDocsResponseBody) SetData(v []*ListFlowDocsResponseBodyData) *ListFlowDocsResponseBody {
	s.Data = v
	return s
}

func (s *ListFlowDocsResponseBody) SetMessage(v string) *ListFlowDocsResponseBody {
	s.Message = &v
	return s
}

type ListFlowDocsResponseBodyData struct {
	FileId   *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileUrl  *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
}

func (s ListFlowDocsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListFlowDocsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFlowDocsResponseBodyData) SetFileId(v string) *ListFlowDocsResponseBodyData {
	s.FileId = &v
	return s
}

func (s *ListFlowDocsResponseBodyData) SetFileName(v string) *ListFlowDocsResponseBodyData {
	s.FileName = &v
	return s
}

func (s *ListFlowDocsResponseBodyData) SetFileUrl(v string) *ListFlowDocsResponseBodyData {
	s.FileUrl = &v
	return s
}

type ListFlowDocsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowDocsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowDocsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowDocsResponse) GoString() string {
	return s.String()
}

func (s *ListFlowDocsResponse) SetHeaders(v map[string]*string) *ListFlowDocsResponse {
	s.Headers = v
	return s
}

func (s *ListFlowDocsResponse) SetBody(v *ListFlowDocsResponseBody) *ListFlowDocsResponse {
	s.Body = v
	return s
}

type ListSealApprovalHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListSealApprovalHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSealApprovalHeaders) GoString() string {
	return s.String()
}

func (s *ListSealApprovalHeaders) SetCommonHeaders(v map[string]*string) *ListSealApprovalHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSealApprovalHeaders) SetXAcsDingtalkAccessToken(v string) *ListSealApprovalHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListSealApprovalRequest struct {
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListSealApprovalRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSealApprovalRequest) GoString() string {
	return s.String()
}

func (s *ListSealApprovalRequest) SetTaskId(v string) *ListSealApprovalRequest {
	s.TaskId = &v
	return s
}

type ListSealApprovalResponseBody struct {
	Code    *int32                              `json:"code,omitempty" xml:"code,omitempty"`
	Data    []*ListSealApprovalResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Message *string                             `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ListSealApprovalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSealApprovalResponseBody) GoString() string {
	return s.String()
}

func (s *ListSealApprovalResponseBody) SetCode(v int32) *ListSealApprovalResponseBody {
	s.Code = &v
	return s
}

func (s *ListSealApprovalResponseBody) SetData(v []*ListSealApprovalResponseBodyData) *ListSealApprovalResponseBody {
	s.Data = v
	return s
}

func (s *ListSealApprovalResponseBody) SetMessage(v string) *ListSealApprovalResponseBody {
	s.Message = &v
	return s
}

type ListSealApprovalResponseBodyData struct {
	ApprovalName       *string                                          `json:"approvalName,omitempty" xml:"approvalName,omitempty"`
	ApprovalNodes      []*ListSealApprovalResponseBodyDataApprovalNodes `json:"approvalNodes,omitempty" xml:"approvalNodes,omitempty" type:"Repeated"`
	EndTime            *int64                                           `json:"endTime,omitempty" xml:"endTime,omitempty"`
	RefuseReason       *string                                          `json:"refuseReason,omitempty" xml:"refuseReason,omitempty"`
	SealIdImg          *string                                          `json:"sealIdImg,omitempty" xml:"sealIdImg,omitempty"`
	SponsorAccountName *string                                          `json:"sponsorAccountName,omitempty" xml:"sponsorAccountName,omitempty"`
	StartTime          *int64                                           `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status             *string                                          `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListSealApprovalResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSealApprovalResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSealApprovalResponseBodyData) SetApprovalName(v string) *ListSealApprovalResponseBodyData {
	s.ApprovalName = &v
	return s
}

func (s *ListSealApprovalResponseBodyData) SetApprovalNodes(v []*ListSealApprovalResponseBodyDataApprovalNodes) *ListSealApprovalResponseBodyData {
	s.ApprovalNodes = v
	return s
}

func (s *ListSealApprovalResponseBodyData) SetEndTime(v int64) *ListSealApprovalResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListSealApprovalResponseBodyData) SetRefuseReason(v string) *ListSealApprovalResponseBodyData {
	s.RefuseReason = &v
	return s
}

func (s *ListSealApprovalResponseBodyData) SetSealIdImg(v string) *ListSealApprovalResponseBodyData {
	s.SealIdImg = &v
	return s
}

func (s *ListSealApprovalResponseBodyData) SetSponsorAccountName(v string) *ListSealApprovalResponseBodyData {
	s.SponsorAccountName = &v
	return s
}

func (s *ListSealApprovalResponseBodyData) SetStartTime(v int64) *ListSealApprovalResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListSealApprovalResponseBodyData) SetStatus(v string) *ListSealApprovalResponseBodyData {
	s.Status = &v
	return s
}

type ListSealApprovalResponseBodyDataApprovalNodes struct {
	ApprovalTime *int64  `json:"approvalTime,omitempty" xml:"approvalTime,omitempty"`
	ApproverName *string `json:"approverName,omitempty" xml:"approverName,omitempty"`
	StartTime    *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListSealApprovalResponseBodyDataApprovalNodes) String() string {
	return tea.Prettify(s)
}

func (s ListSealApprovalResponseBodyDataApprovalNodes) GoString() string {
	return s.String()
}

func (s *ListSealApprovalResponseBodyDataApprovalNodes) SetApprovalTime(v int64) *ListSealApprovalResponseBodyDataApprovalNodes {
	s.ApprovalTime = &v
	return s
}

func (s *ListSealApprovalResponseBodyDataApprovalNodes) SetApproverName(v string) *ListSealApprovalResponseBodyDataApprovalNodes {
	s.ApproverName = &v
	return s
}

func (s *ListSealApprovalResponseBodyDataApprovalNodes) SetStartTime(v int64) *ListSealApprovalResponseBodyDataApprovalNodes {
	s.StartTime = &v
	return s
}

func (s *ListSealApprovalResponseBodyDataApprovalNodes) SetStatus(v string) *ListSealApprovalResponseBodyDataApprovalNodes {
	s.Status = &v
	return s
}

type ListSealApprovalResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSealApprovalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSealApprovalResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSealApprovalResponse) GoString() string {
	return s.String()
}

func (s *ListSealApprovalResponse) SetHeaders(v map[string]*string) *ListSealApprovalResponse {
	s.Headers = v
	return s
}

func (s *ListSealApprovalResponse) SetBody(v *ListSealApprovalResponseBody) *ListSealApprovalResponse {
	s.Body = v
	return s
}

type OrderResaleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s OrderResaleHeaders) String() string {
	return tea.Prettify(s)
}

func (s OrderResaleHeaders) GoString() string {
	return s.String()
}

func (s *OrderResaleHeaders) SetCommonHeaders(v map[string]*string) *OrderResaleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OrderResaleHeaders) SetXAcsDingtalkAccessToken(v string) *OrderResaleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type OrderResaleRequest struct {
	OrderCreateTime  *int64  `json:"orderCreateTime,omitempty" xml:"orderCreateTime,omitempty"`
	OrderId          *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	Quantity         *int64  `json:"quantity,omitempty" xml:"quantity,omitempty"`
	ServiceStartTime *int64  `json:"serviceStartTime,omitempty" xml:"serviceStartTime,omitempty"`
	ServiceStopTime  *int64  `json:"serviceStopTime,omitempty" xml:"serviceStopTime,omitempty"`
}

func (s OrderResaleRequest) String() string {
	return tea.Prettify(s)
}

func (s OrderResaleRequest) GoString() string {
	return s.String()
}

func (s *OrderResaleRequest) SetOrderCreateTime(v int64) *OrderResaleRequest {
	s.OrderCreateTime = &v
	return s
}

func (s *OrderResaleRequest) SetOrderId(v string) *OrderResaleRequest {
	s.OrderId = &v
	return s
}

func (s *OrderResaleRequest) SetQuantity(v int64) *OrderResaleRequest {
	s.Quantity = &v
	return s
}

func (s *OrderResaleRequest) SetServiceStartTime(v int64) *OrderResaleRequest {
	s.ServiceStartTime = &v
	return s
}

func (s *OrderResaleRequest) SetServiceStopTime(v int64) *OrderResaleRequest {
	s.ServiceStopTime = &v
	return s
}

type OrderResaleResponseBody struct {
	Code    *int32                       `json:"code,omitempty" xml:"code,omitempty"`
	Data    *OrderResaleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                      `json:"message,omitempty" xml:"message,omitempty"`
}

func (s OrderResaleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OrderResaleResponseBody) GoString() string {
	return s.String()
}

func (s *OrderResaleResponseBody) SetCode(v int32) *OrderResaleResponseBody {
	s.Code = &v
	return s
}

func (s *OrderResaleResponseBody) SetData(v *OrderResaleResponseBodyData) *OrderResaleResponseBody {
	s.Data = v
	return s
}

func (s *OrderResaleResponseBody) SetMessage(v string) *OrderResaleResponseBody {
	s.Message = &v
	return s
}

type OrderResaleResponseBodyData struct {
	EsignOrderId *string `json:"esignOrderId,omitempty" xml:"esignOrderId,omitempty"`
}

func (s OrderResaleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OrderResaleResponseBodyData) GoString() string {
	return s.String()
}

func (s *OrderResaleResponseBodyData) SetEsignOrderId(v string) *OrderResaleResponseBodyData {
	s.EsignOrderId = &v
	return s
}

type OrderResaleResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OrderResaleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OrderResaleResponse) String() string {
	return tea.Prettify(s)
}

func (s OrderResaleResponse) GoString() string {
	return s.String()
}

func (s *OrderResaleResponse) SetHeaders(v map[string]*string) *OrderResaleResponse {
	s.Headers = v
	return s
}

func (s *OrderResaleResponse) SetBody(v *OrderResaleResponseBody) *OrderResaleResponse {
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

func (client *Client) AuthUrl(request *AuthUrlRequest) (_result *AuthUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AuthUrlHeaders{}
	_result = &AuthUrlResponse{}
	_body, _err := client.AuthUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AuthUrlWithOptions(request *AuthUrlRequest, headers *AuthUrlHeaders, runtime *util.RuntimeOptions) (_result *AuthUrlResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &AuthUrlResponse{}
	_body, _err := client.DoROARequest(tea.String("AuthUrl"), tea.String("esign_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/esign/auths/url"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelCorpAuth() (_result *CancelCorpAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CancelCorpAuthHeaders{}
	_result = &CancelCorpAuthResponse{}
	_body, _err := client.CancelCorpAuthWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelCorpAuthWithOptions(headers *CancelCorpAuthHeaders, runtime *util.RuntimeOptions) (_result *CancelCorpAuthResponse, _err error) {
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
	_result = &CancelCorpAuthResponse{}
	_body, _err := client.DoROARequest(tea.String("CancelCorpAuth"), tea.String("esign_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/esign/corps/auth/cancel"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChannelOrder(request *ChannelOrderRequest) (_result *ChannelOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ChannelOrderHeaders{}
	_result = &ChannelOrderResponse{}
	_body, _err := client.ChannelOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChannelOrderWithOptions(request *ChannelOrderRequest, headers *ChannelOrderHeaders, runtime *util.RuntimeOptions) (_result *ChannelOrderResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &ChannelOrderResponse{}
	_body, _err := client.DoROARequest(tea.String("ChannelOrder"), tea.String("esign_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/esign/orders/channel"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ContractMargin() (_result *ContractMarginResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ContractMarginHeaders{}
	_result = &ContractMarginResponse{}
	_body, _err := client.ContractMarginWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ContractMarginWithOptions(headers *ContractMarginHeaders, runtime *util.RuntimeOptions) (_result *ContractMarginResponse, _err error) {
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
	_result = &ContractMarginResponse{}
	_body, _err := client.DoROARequest(tea.String("ContractMargin"), tea.String("esign_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/esign/contracts/margin"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CorpConsole() (_result *CorpConsoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CorpConsoleHeaders{}
	_result = &CorpConsoleResponse{}
	_body, _err := client.CorpConsoleWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CorpConsoleWithOptions(headers *CorpConsoleHeaders, runtime *util.RuntimeOptions) (_result *CorpConsoleResponse, _err error) {
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
	_result = &CorpConsoleResponse{}
	_body, _err := client.DoROARequest(tea.String("CorpConsole"), tea.String("esign_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/esign/corps/console"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CorpInfo() (_result *CorpInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CorpInfoHeaders{}
	_result = &CorpInfoResponse{}
	_body, _err := client.CorpInfoWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CorpInfoWithOptions(headers *CorpInfoHeaders, runtime *util.RuntimeOptions) (_result *CorpInfoResponse, _err error) {
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
	_result = &CorpInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("CorpInfo"), tea.String("esign_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/esign/corps/info"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDeveloper(request *CreateDeveloperRequest) (_result *CreateDeveloperResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateDeveloperHeaders{}
	_result = &CreateDeveloperResponse{}
	_body, _err := client.CreateDeveloperWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDeveloperWithOptions(request *CreateDeveloperRequest, headers *CreateDeveloperHeaders, runtime *util.RuntimeOptions) (_result *CreateDeveloperResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateDeveloperResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateDeveloper"), tea.String("esign_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/esign/developers/create"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCorpRealnameUrl(request *GetCorpRealnameUrlRequest) (_result *GetCorpRealnameUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCorpRealnameUrlHeaders{}
	_result = &GetCorpRealnameUrlResponse{}
	_body, _err := client.GetCorpRealnameUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCorpRealnameUrlWithOptions(request *GetCorpRealnameUrlRequest, headers *GetCorpRealnameUrlHeaders, runtime *util.RuntimeOptions) (_result *GetCorpRealnameUrlResponse, _err error) {
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
	_result = &GetCorpRealnameUrlResponse{}
	_body, _err := client.DoROARequest(tea.String("GetCorpRealnameUrl"), tea.String("esign_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/esign/corps/realname"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCropStatus() (_result *GetCropStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCropStatusHeaders{}
	_result = &GetCropStatusResponse{}
	_body, _err := client.GetCropStatusWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCropStatusWithOptions(headers *GetCropStatusHeaders, runtime *util.RuntimeOptions) (_result *GetCropStatusResponse, _err error) {
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
	_result = &GetCropStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("GetCropStatus"), tea.String("esign_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/esign/corps/statuses"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFile(fileId *string) (_result *GetFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFileHeaders{}
	_result = &GetFileResponse{}
	_body, _err := client.GetFileWithOptions(fileId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFileWithOptions(fileId *string, headers *GetFileHeaders, runtime *util.RuntimeOptions) (_result *GetFileResponse, _err error) {
	fileId = openapiutil.GetEncodeParam(fileId)
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
	_result = &GetFileResponse{}
	_body, _err := client.DoROARequest(tea.String("GetFile"), tea.String("esign_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/esign/files/"+tea.StringValue(fileId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFlowDetail(request *GetFlowDetailRequest) (_result *GetFlowDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFlowDetailHeaders{}
	_result = &GetFlowDetailResponse{}
	_body, _err := client.GetFlowDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFlowDetailWithOptions(request *GetFlowDetailRequest, headers *GetFlowDetailHeaders, runtime *util.RuntimeOptions) (_result *GetFlowDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
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
	_result = &GetFlowDetailResponse{}
	_body, _err := client.DoROARequest(tea.String("GetFlowDetail"), tea.String("esign_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/esign/flows/detail"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFlowSignDetail(request *GetFlowSignDetailRequest) (_result *GetFlowSignDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFlowSignDetailHeaders{}
	_result = &GetFlowSignDetailResponse{}
	_body, _err := client.GetFlowSignDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFlowSignDetailWithOptions(request *GetFlowSignDetailRequest, headers *GetFlowSignDetailHeaders, runtime *util.RuntimeOptions) (_result *GetFlowSignDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
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
	_result = &GetFlowSignDetailResponse{}
	_body, _err := client.DoROARequest(tea.String("GetFlowSignDetail"), tea.String("esign_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/esign/flows/sign/detail"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProcessStartUrl(request *GetProcessStartUrlRequest) (_result *GetProcessStartUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetProcessStartUrlHeaders{}
	_result = &GetProcessStartUrlResponse{}
	_body, _err := client.GetProcessStartUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProcessStartUrlWithOptions(request *GetProcessStartUrlRequest, headers *GetProcessStartUrlHeaders, runtime *util.RuntimeOptions) (_result *GetProcessStartUrlResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SourceInfo))) {
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
	_result = &GetProcessStartUrlResponse{}
	_body, _err := client.DoROARequest(tea.String("GetProcessStartUrl"), tea.String("esign_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/esign/process/start"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSignNoticeUrl(request *GetSignNoticeUrlRequest) (_result *GetSignNoticeUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSignNoticeUrlHeaders{}
	_result = &GetSignNoticeUrlResponse{}
	_body, _err := client.GetSignNoticeUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSignNoticeUrlWithOptions(request *GetSignNoticeUrlRequest, headers *GetSignNoticeUrlHeaders, runtime *util.RuntimeOptions) (_result *GetSignNoticeUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
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
	_result = &GetSignNoticeUrlResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSignNoticeUrl"), tea.String("esign_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/esign/signs/notice/url"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUploadUrl(request *GetUploadUrlRequest) (_result *GetUploadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUploadUrlHeaders{}
	_result = &GetUploadUrlResponse{}
	_body, _err := client.GetUploadUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUploadUrlWithOptions(request *GetUploadUrlRequest, headers *GetUploadUrlHeaders, runtime *util.RuntimeOptions) (_result *GetUploadUrlResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &GetUploadUrlResponse{}
	_body, _err := client.DoROARequest(tea.String("GetUploadUrl"), tea.String("esign_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/esign/files/getUploadUrl"), tea.String("json"), req, runtime)
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

func (client *Client) GetUserInfoWithOptions(userId *string, headers *GetUserInfoHeaders, runtime *util.RuntimeOptions) (_result *GetUserInfoResponse, _err error) {
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
	_result = &GetUserInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetUserInfo"), tea.String("esign_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/esign/users/"+tea.StringValue(userId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserRealnameUrl(request *GetUserRealnameUrlRequest) (_result *GetUserRealnameUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserRealnameUrlHeaders{}
	_result = &GetUserRealnameUrlResponse{}
	_body, _err := client.GetUserRealnameUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserRealnameUrlWithOptions(request *GetUserRealnameUrlRequest, headers *GetUserRealnameUrlHeaders, runtime *util.RuntimeOptions) (_result *GetUserRealnameUrlResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &GetUserRealnameUrlResponse{}
	_body, _err := client.DoROARequest(tea.String("GetUserRealnameUrl"), tea.String("esign_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/esign/users/realname"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowDocs(request *ListFlowDocsRequest) (_result *ListFlowDocsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListFlowDocsHeaders{}
	_result = &ListFlowDocsResponse{}
	_body, _err := client.ListFlowDocsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowDocsWithOptions(request *ListFlowDocsRequest, headers *ListFlowDocsHeaders, runtime *util.RuntimeOptions) (_result *ListFlowDocsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
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
	_result = &ListFlowDocsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListFlowDocs"), tea.String("esign_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/esign/flows/docs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSealApproval(request *ListSealApprovalRequest) (_result *ListSealApprovalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSealApprovalHeaders{}
	_result = &ListSealApprovalResponse{}
	_body, _err := client.ListSealApprovalWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSealApprovalWithOptions(request *ListSealApprovalRequest, headers *ListSealApprovalHeaders, runtime *util.RuntimeOptions) (_result *ListSealApprovalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
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
	_result = &ListSealApprovalResponse{}
	_body, _err := client.DoROARequest(tea.String("ListSealApproval"), tea.String("esign_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/esign/seals/approval/list"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OrderResale(request *OrderResaleRequest) (_result *OrderResaleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OrderResaleHeaders{}
	_result = &OrderResaleResponse{}
	_body, _err := client.OrderResaleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OrderResaleWithOptions(request *OrderResaleRequest, headers *OrderResaleHeaders, runtime *util.RuntimeOptions) (_result *OrderResaleResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &OrderResaleResponse{}
	_body, _err := client.DoROARequest(tea.String("OrderResale"), tea.String("esign_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/esign/orders/resale"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
