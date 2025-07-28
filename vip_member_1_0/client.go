// This file is auto-generated, don't edit it. Thanks.
package vip_member_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DirectRedeemVipMemberByMobileHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DirectRedeemVipMemberByMobileHeaders) String() string {
	return tea.Prettify(s)
}

func (s DirectRedeemVipMemberByMobileHeaders) GoString() string {
	return s.String()
}

func (s *DirectRedeemVipMemberByMobileHeaders) SetCommonHeaders(v map[string]*string) *DirectRedeemVipMemberByMobileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DirectRedeemVipMemberByMobileHeaders) SetXAcsDingtalkAccessToken(v string) *DirectRedeemVipMemberByMobileHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DirectRedeemVipMemberByMobileRequest struct {
	BizRequestId *string `json:"bizRequestId,omitempty" xml:"bizRequestId,omitempty"`
	Channel      *string `json:"channel,omitempty" xml:"channel,omitempty"`
	DingtalkId   *string `json:"dingtalkId,omitempty" xml:"dingtalkId,omitempty"`
	Duration     *int64  `json:"duration,omitempty" xml:"duration,omitempty"`
	Mobile       *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	Uuid         *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s DirectRedeemVipMemberByMobileRequest) String() string {
	return tea.Prettify(s)
}

func (s DirectRedeemVipMemberByMobileRequest) GoString() string {
	return s.String()
}

func (s *DirectRedeemVipMemberByMobileRequest) SetBizRequestId(v string) *DirectRedeemVipMemberByMobileRequest {
	s.BizRequestId = &v
	return s
}

func (s *DirectRedeemVipMemberByMobileRequest) SetChannel(v string) *DirectRedeemVipMemberByMobileRequest {
	s.Channel = &v
	return s
}

func (s *DirectRedeemVipMemberByMobileRequest) SetDingtalkId(v string) *DirectRedeemVipMemberByMobileRequest {
	s.DingtalkId = &v
	return s
}

func (s *DirectRedeemVipMemberByMobileRequest) SetDuration(v int64) *DirectRedeemVipMemberByMobileRequest {
	s.Duration = &v
	return s
}

func (s *DirectRedeemVipMemberByMobileRequest) SetMobile(v string) *DirectRedeemVipMemberByMobileRequest {
	s.Mobile = &v
	return s
}

func (s *DirectRedeemVipMemberByMobileRequest) SetUuid(v string) *DirectRedeemVipMemberByMobileRequest {
	s.Uuid = &v
	return s
}

type DirectRedeemVipMemberByMobileResponseBody struct {
	BizRequestId *string `json:"bizRequestId,omitempty" xml:"bizRequestId,omitempty"`
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg     *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result       *bool   `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DirectRedeemVipMemberByMobileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DirectRedeemVipMemberByMobileResponseBody) GoString() string {
	return s.String()
}

func (s *DirectRedeemVipMemberByMobileResponseBody) SetBizRequestId(v string) *DirectRedeemVipMemberByMobileResponseBody {
	s.BizRequestId = &v
	return s
}

func (s *DirectRedeemVipMemberByMobileResponseBody) SetErrorCode(v string) *DirectRedeemVipMemberByMobileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DirectRedeemVipMemberByMobileResponseBody) SetErrorMsg(v string) *DirectRedeemVipMemberByMobileResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DirectRedeemVipMemberByMobileResponseBody) SetResult(v bool) *DirectRedeemVipMemberByMobileResponseBody {
	s.Result = &v
	return s
}

type DirectRedeemVipMemberByMobileResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DirectRedeemVipMemberByMobileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DirectRedeemVipMemberByMobileResponse) String() string {
	return tea.Prettify(s)
}

func (s DirectRedeemVipMemberByMobileResponse) GoString() string {
	return s.String()
}

func (s *DirectRedeemVipMemberByMobileResponse) SetHeaders(v map[string]*string) *DirectRedeemVipMemberByMobileResponse {
	s.Headers = v
	return s
}

func (s *DirectRedeemVipMemberByMobileResponse) SetStatusCode(v int32) *DirectRedeemVipMemberByMobileResponse {
	s.StatusCode = &v
	return s
}

func (s *DirectRedeemVipMemberByMobileResponse) SetBody(v *DirectRedeemVipMemberByMobileResponseBody) *DirectRedeemVipMemberByMobileResponse {
	s.Body = v
	return s
}

type InvalidRedeemVipMemberByBizRequestIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InvalidRedeemVipMemberByBizRequestIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s InvalidRedeemVipMemberByBizRequestIdHeaders) GoString() string {
	return s.String()
}

func (s *InvalidRedeemVipMemberByBizRequestIdHeaders) SetCommonHeaders(v map[string]*string) *InvalidRedeemVipMemberByBizRequestIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvalidRedeemVipMemberByBizRequestIdHeaders) SetXAcsDingtalkAccessToken(v string) *InvalidRedeemVipMemberByBizRequestIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InvalidRedeemVipMemberByBizRequestIdRequest struct {
	BizRequestId *string `json:"bizRequestId,omitempty" xml:"bizRequestId,omitempty"`
	Channel      *string `json:"channel,omitempty" xml:"channel,omitempty"`
	DingtalkId   *string `json:"dingtalkId,omitempty" xml:"dingtalkId,omitempty"`
	Duration     *int64  `json:"duration,omitempty" xml:"duration,omitempty"`
	Mobile       *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	Uuid         *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s InvalidRedeemVipMemberByBizRequestIdRequest) String() string {
	return tea.Prettify(s)
}

func (s InvalidRedeemVipMemberByBizRequestIdRequest) GoString() string {
	return s.String()
}

func (s *InvalidRedeemVipMemberByBizRequestIdRequest) SetBizRequestId(v string) *InvalidRedeemVipMemberByBizRequestIdRequest {
	s.BizRequestId = &v
	return s
}

func (s *InvalidRedeemVipMemberByBizRequestIdRequest) SetChannel(v string) *InvalidRedeemVipMemberByBizRequestIdRequest {
	s.Channel = &v
	return s
}

func (s *InvalidRedeemVipMemberByBizRequestIdRequest) SetDingtalkId(v string) *InvalidRedeemVipMemberByBizRequestIdRequest {
	s.DingtalkId = &v
	return s
}

func (s *InvalidRedeemVipMemberByBizRequestIdRequest) SetDuration(v int64) *InvalidRedeemVipMemberByBizRequestIdRequest {
	s.Duration = &v
	return s
}

func (s *InvalidRedeemVipMemberByBizRequestIdRequest) SetMobile(v string) *InvalidRedeemVipMemberByBizRequestIdRequest {
	s.Mobile = &v
	return s
}

func (s *InvalidRedeemVipMemberByBizRequestIdRequest) SetUuid(v string) *InvalidRedeemVipMemberByBizRequestIdRequest {
	s.Uuid = &v
	return s
}

type InvalidRedeemVipMemberByBizRequestIdResponseBody struct {
	BizRequestId *string `json:"bizRequestId,omitempty" xml:"bizRequestId,omitempty"`
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg     *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result       *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s InvalidRedeemVipMemberByBizRequestIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvalidRedeemVipMemberByBizRequestIdResponseBody) GoString() string {
	return s.String()
}

func (s *InvalidRedeemVipMemberByBizRequestIdResponseBody) SetBizRequestId(v string) *InvalidRedeemVipMemberByBizRequestIdResponseBody {
	s.BizRequestId = &v
	return s
}

func (s *InvalidRedeemVipMemberByBizRequestIdResponseBody) SetErrorCode(v string) *InvalidRedeemVipMemberByBizRequestIdResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *InvalidRedeemVipMemberByBizRequestIdResponseBody) SetErrorMsg(v string) *InvalidRedeemVipMemberByBizRequestIdResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *InvalidRedeemVipMemberByBizRequestIdResponseBody) SetResult(v string) *InvalidRedeemVipMemberByBizRequestIdResponseBody {
	s.Result = &v
	return s
}

type InvalidRedeemVipMemberByBizRequestIdResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvalidRedeemVipMemberByBizRequestIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvalidRedeemVipMemberByBizRequestIdResponse) String() string {
	return tea.Prettify(s)
}

func (s InvalidRedeemVipMemberByBizRequestIdResponse) GoString() string {
	return s.String()
}

func (s *InvalidRedeemVipMemberByBizRequestIdResponse) SetHeaders(v map[string]*string) *InvalidRedeemVipMemberByBizRequestIdResponse {
	s.Headers = v
	return s
}

func (s *InvalidRedeemVipMemberByBizRequestIdResponse) SetStatusCode(v int32) *InvalidRedeemVipMemberByBizRequestIdResponse {
	s.StatusCode = &v
	return s
}

func (s *InvalidRedeemVipMemberByBizRequestIdResponse) SetBody(v *InvalidRedeemVipMemberByBizRequestIdResponseBody) *InvalidRedeemVipMemberByBizRequestIdResponse {
	s.Body = v
	return s
}

type QueryRedeemVipMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryRedeemVipMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryRedeemVipMemberHeaders) GoString() string {
	return s.String()
}

func (s *QueryRedeemVipMemberHeaders) SetCommonHeaders(v map[string]*string) *QueryRedeemVipMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryRedeemVipMemberHeaders) SetXAcsDingtalkAccessToken(v string) *QueryRedeemVipMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryRedeemVipMemberRequest struct {
	BizRequestId *string `json:"bizRequestId,omitempty" xml:"bizRequestId,omitempty"`
	Channel      *string `json:"channel,omitempty" xml:"channel,omitempty"`
	DingtalkId   *string `json:"dingtalkId,omitempty" xml:"dingtalkId,omitempty"`
	Duration     *int64  `json:"duration,omitempty" xml:"duration,omitempty"`
	Mobile       *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	Uuid         *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s QueryRedeemVipMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRedeemVipMemberRequest) GoString() string {
	return s.String()
}

func (s *QueryRedeemVipMemberRequest) SetBizRequestId(v string) *QueryRedeemVipMemberRequest {
	s.BizRequestId = &v
	return s
}

func (s *QueryRedeemVipMemberRequest) SetChannel(v string) *QueryRedeemVipMemberRequest {
	s.Channel = &v
	return s
}

func (s *QueryRedeemVipMemberRequest) SetDingtalkId(v string) *QueryRedeemVipMemberRequest {
	s.DingtalkId = &v
	return s
}

func (s *QueryRedeemVipMemberRequest) SetDuration(v int64) *QueryRedeemVipMemberRequest {
	s.Duration = &v
	return s
}

func (s *QueryRedeemVipMemberRequest) SetMobile(v string) *QueryRedeemVipMemberRequest {
	s.Mobile = &v
	return s
}

func (s *QueryRedeemVipMemberRequest) SetUuid(v string) *QueryRedeemVipMemberRequest {
	s.Uuid = &v
	return s
}

type QueryRedeemVipMemberResponseBody struct {
	ErrorCode    *string                                         `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg     *string                                         `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	QueryResults []*QueryRedeemVipMemberResponseBodyQueryResults `json:"queryResults,omitempty" xml:"queryResults,omitempty" type:"Repeated"`
	Result       *bool                                           `json:"result,omitempty" xml:"result,omitempty"`
}

func (s QueryRedeemVipMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRedeemVipMemberResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRedeemVipMemberResponseBody) SetErrorCode(v string) *QueryRedeemVipMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryRedeemVipMemberResponseBody) SetErrorMsg(v string) *QueryRedeemVipMemberResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryRedeemVipMemberResponseBody) SetQueryResults(v []*QueryRedeemVipMemberResponseBodyQueryResults) *QueryRedeemVipMemberResponseBody {
	s.QueryResults = v
	return s
}

func (s *QueryRedeemVipMemberResponseBody) SetResult(v bool) *QueryRedeemVipMemberResponseBody {
	s.Result = &v
	return s
}

type QueryRedeemVipMemberResponseBodyQueryResults struct {
	Action     *string `json:"action,omitempty" xml:"action,omitempty"`
	ActionTime *string `json:"actionTime,omitempty" xml:"actionTime,omitempty"`
	DingtalkId *string `json:"dingtalkId,omitempty" xml:"dingtalkId,omitempty"`
	Duration   *int64  `json:"duration,omitempty" xml:"duration,omitempty"`
	Nick       *string `json:"nick,omitempty" xml:"nick,omitempty"`
}

func (s QueryRedeemVipMemberResponseBodyQueryResults) String() string {
	return tea.Prettify(s)
}

func (s QueryRedeemVipMemberResponseBodyQueryResults) GoString() string {
	return s.String()
}

func (s *QueryRedeemVipMemberResponseBodyQueryResults) SetAction(v string) *QueryRedeemVipMemberResponseBodyQueryResults {
	s.Action = &v
	return s
}

func (s *QueryRedeemVipMemberResponseBodyQueryResults) SetActionTime(v string) *QueryRedeemVipMemberResponseBodyQueryResults {
	s.ActionTime = &v
	return s
}

func (s *QueryRedeemVipMemberResponseBodyQueryResults) SetDingtalkId(v string) *QueryRedeemVipMemberResponseBodyQueryResults {
	s.DingtalkId = &v
	return s
}

func (s *QueryRedeemVipMemberResponseBodyQueryResults) SetDuration(v int64) *QueryRedeemVipMemberResponseBodyQueryResults {
	s.Duration = &v
	return s
}

func (s *QueryRedeemVipMemberResponseBodyQueryResults) SetNick(v string) *QueryRedeemVipMemberResponseBodyQueryResults {
	s.Nick = &v
	return s
}

type QueryRedeemVipMemberResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRedeemVipMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRedeemVipMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRedeemVipMemberResponse) GoString() string {
	return s.String()
}

func (s *QueryRedeemVipMemberResponse) SetHeaders(v map[string]*string) *QueryRedeemVipMemberResponse {
	s.Headers = v
	return s
}

func (s *QueryRedeemVipMemberResponse) SetStatusCode(v int32) *QueryRedeemVipMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRedeemVipMemberResponse) SetBody(v *QueryRedeemVipMemberResponseBody) *QueryRedeemVipMemberResponse {
	s.Body = v
	return s
}

type QueryVipMemberInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryVipMemberInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryVipMemberInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryVipMemberInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryVipMemberInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryVipMemberInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryVipMemberInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryVipMemberInfoRequest struct {
	ChannelCode *string `json:"channelCode,omitempty" xml:"channelCode,omitempty"`
}

func (s QueryVipMemberInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryVipMemberInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryVipMemberInfoRequest) SetChannelCode(v string) *QueryVipMemberInfoRequest {
	s.ChannelCode = &v
	return s
}

type QueryVipMemberInfoResponseBody struct {
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	IsVip      *bool   `json:"isVip,omitempty" xml:"isVip,omitempty"`
}

func (s QueryVipMemberInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryVipMemberInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVipMemberInfoResponseBody) SetExpireTime(v string) *QueryVipMemberInfoResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *QueryVipMemberInfoResponseBody) SetIsVip(v bool) *QueryVipMemberInfoResponseBody {
	s.IsVip = &v
	return s
}

type QueryVipMemberInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryVipMemberInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryVipMemberInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryVipMemberInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryVipMemberInfoResponse) SetHeaders(v map[string]*string) *QueryVipMemberInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryVipMemberInfoResponse) SetStatusCode(v int32) *QueryVipMemberInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVipMemberInfoResponse) SetBody(v *QueryVipMemberInfoResponseBody) *QueryVipMemberInfoResponse {
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
// 通过手机号直充365会员
//
// @param request - DirectRedeemVipMemberByMobileRequest
//
// @param headers - DirectRedeemVipMemberByMobileHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DirectRedeemVipMemberByMobileResponse
func (client *Client) DirectRedeemVipMemberByMobileWithOptions(request *DirectRedeemVipMemberByMobileRequest, headers *DirectRedeemVipMemberByMobileHeaders, runtime *util.RuntimeOptions) (_result *DirectRedeemVipMemberByMobileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizRequestId)) {
		body["bizRequestId"] = request.BizRequestId
	}

	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		body["channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.DingtalkId)) {
		body["dingtalkId"] = request.DingtalkId
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		body["duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		body["mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
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
		Action:      tea.String("DirectRedeemVipMemberByMobile"),
		Version:     tea.String("vipMember_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/vipMember/users/directRedeemVip"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DirectRedeemVipMemberByMobileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过手机号直充365会员
//
// @param request - DirectRedeemVipMemberByMobileRequest
//
// @return DirectRedeemVipMemberByMobileResponse
func (client *Client) DirectRedeemVipMemberByMobile(request *DirectRedeemVipMemberByMobileRequest) (_result *DirectRedeemVipMemberByMobileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DirectRedeemVipMemberByMobileHeaders{}
	_result = &DirectRedeemVipMemberByMobileResponse{}
	_body, _err := client.DirectRedeemVipMemberByMobileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过虚拟订单号作废365会员权益
//
// @param request - InvalidRedeemVipMemberByBizRequestIdRequest
//
// @param headers - InvalidRedeemVipMemberByBizRequestIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvalidRedeemVipMemberByBizRequestIdResponse
func (client *Client) InvalidRedeemVipMemberByBizRequestIdWithOptions(request *InvalidRedeemVipMemberByBizRequestIdRequest, headers *InvalidRedeemVipMemberByBizRequestIdHeaders, runtime *util.RuntimeOptions) (_result *InvalidRedeemVipMemberByBizRequestIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizRequestId)) {
		body["bizRequestId"] = request.BizRequestId
	}

	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		body["channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.DingtalkId)) {
		body["dingtalkId"] = request.DingtalkId
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		body["duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		body["mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
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
		Action:      tea.String("InvalidRedeemVipMemberByBizRequestId"),
		Version:     tea.String("vipMember_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/vipMember/users/invalidRedeemVip"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &InvalidRedeemVipMemberByBizRequestIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过虚拟订单号作废365会员权益
//
// @param request - InvalidRedeemVipMemberByBizRequestIdRequest
//
// @return InvalidRedeemVipMemberByBizRequestIdResponse
func (client *Client) InvalidRedeemVipMemberByBizRequestId(request *InvalidRedeemVipMemberByBizRequestIdRequest) (_result *InvalidRedeemVipMemberByBizRequestIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InvalidRedeemVipMemberByBizRequestIdHeaders{}
	_result = &InvalidRedeemVipMemberByBizRequestIdResponse{}
	_body, _err := client.InvalidRedeemVipMemberByBizRequestIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询365会员直充信息
//
// @param request - QueryRedeemVipMemberRequest
//
// @param headers - QueryRedeemVipMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRedeemVipMemberResponse
func (client *Client) QueryRedeemVipMemberWithOptions(request *QueryRedeemVipMemberRequest, headers *QueryRedeemVipMemberHeaders, runtime *util.RuntimeOptions) (_result *QueryRedeemVipMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizRequestId)) {
		body["bizRequestId"] = request.BizRequestId
	}

	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		body["channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.DingtalkId)) {
		body["dingtalkId"] = request.DingtalkId
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		body["duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		body["mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
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
		Action:      tea.String("QueryRedeemVipMember"),
		Version:     tea.String("vipMember_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/vipMember/users/queryRedeemVip"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRedeemVipMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询365会员直充信息
//
// @param request - QueryRedeemVipMemberRequest
//
// @return QueryRedeemVipMemberResponse
func (client *Client) QueryRedeemVipMember(request *QueryRedeemVipMemberRequest) (_result *QueryRedeemVipMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryRedeemVipMemberHeaders{}
	_result = &QueryRedeemVipMemberResponse{}
	_body, _err := client.QueryRedeemVipMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询365会员信息
//
// @param request - QueryVipMemberInfoRequest
//
// @param headers - QueryVipMemberInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryVipMemberInfoResponse
func (client *Client) QueryVipMemberInfoWithOptions(request *QueryVipMemberInfoRequest, headers *QueryVipMemberInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryVipMemberInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelCode)) {
		body["channelCode"] = request.ChannelCode
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
		Action:      tea.String("QueryVipMemberInfo"),
		Version:     tea.String("vipMember_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/vipMember/users/memberInfos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryVipMemberInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询365会员信息
//
// @param request - QueryVipMemberInfoRequest
//
// @return QueryVipMemberInfoResponse
func (client *Client) QueryVipMemberInfo(request *QueryVipMemberInfoRequest) (_result *QueryVipMemberInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryVipMemberInfoHeaders{}
	_result = &QueryVipMemberInfoResponse{}
	_body, _err := client.QueryVipMemberInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
