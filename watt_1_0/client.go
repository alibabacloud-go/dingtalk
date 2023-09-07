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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckInCrowdsByMobileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendBannerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendPopupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendSearchShadeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
