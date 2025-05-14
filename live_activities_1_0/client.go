// This file is auto-generated, don't edit it. Thanks.
package live_activities_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type PushLiveActivityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PushLiveActivityHeaders) String() string {
	return tea.Prettify(s)
}

func (s PushLiveActivityHeaders) GoString() string {
	return s.String()
}

func (s *PushLiveActivityHeaders) SetCommonHeaders(v map[string]*string) *PushLiveActivityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PushLiveActivityHeaders) SetXAcsDingtalkAccessToken(v string) *PushLiveActivityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PushLiveActivityRequest struct {
	ActivityEventData   *PushLiveActivityRequestActivityEventData   `json:"activityEventData,omitempty" xml:"activityEventData,omitempty" type:"Struct"`
	ActivityEventOption *PushLiveActivityRequestActivityEventOption `json:"activityEventOption,omitempty" xml:"activityEventOption,omitempty" type:"Struct"`
	// example:
	//
	// bizUniqueId
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
}

func (s PushLiveActivityRequest) String() string {
	return tea.Prettify(s)
}

func (s PushLiveActivityRequest) GoString() string {
	return s.String()
}

func (s *PushLiveActivityRequest) SetActivityEventData(v *PushLiveActivityRequestActivityEventData) *PushLiveActivityRequest {
	s.ActivityEventData = v
	return s
}

func (s *PushLiveActivityRequest) SetActivityEventOption(v *PushLiveActivityRequestActivityEventOption) *PushLiveActivityRequest {
	s.ActivityEventOption = v
	return s
}

func (s *PushLiveActivityRequest) SetActivityId(v string) *PushLiveActivityRequest {
	s.ActivityId = &v
	return s
}

type PushLiveActivityRequestActivityEventData struct {
	I18nContentState interface{} `json:"i18nContentState,omitempty" xml:"i18nContentState,omitempty"`
	// example:
	//
	// ride_with_alibtrip
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s PushLiveActivityRequestActivityEventData) String() string {
	return tea.Prettify(s)
}

func (s PushLiveActivityRequestActivityEventData) GoString() string {
	return s.String()
}

func (s *PushLiveActivityRequestActivityEventData) SetI18nContentState(v interface{}) *PushLiveActivityRequestActivityEventData {
	s.I18nContentState = v
	return s
}

func (s *PushLiveActivityRequestActivityEventData) SetTemplateId(v string) *PushLiveActivityRequestActivityEventData {
	s.TemplateId = &v
	return s
}

type PushLiveActivityRequestActivityEventOption struct {
	// example:
	//
	// 1686903998
	DismissalDate *int64 `json:"dismissalDate,omitempty" xml:"dismissalDate,omitempty"`
	// example:
	//
	// update
	PushType *string `json:"pushType,omitempty" xml:"pushType,omitempty"`
	// example:
	//
	// 1686903998
	SendDate *int64 `json:"sendDate,omitempty" xml:"sendDate,omitempty"`
	// example:
	//
	// 1686903998
	StaleDate *int64 `json:"staleDate,omitempty" xml:"staleDate,omitempty"`
}

func (s PushLiveActivityRequestActivityEventOption) String() string {
	return tea.Prettify(s)
}

func (s PushLiveActivityRequestActivityEventOption) GoString() string {
	return s.String()
}

func (s *PushLiveActivityRequestActivityEventOption) SetDismissalDate(v int64) *PushLiveActivityRequestActivityEventOption {
	s.DismissalDate = &v
	return s
}

func (s *PushLiveActivityRequestActivityEventOption) SetPushType(v string) *PushLiveActivityRequestActivityEventOption {
	s.PushType = &v
	return s
}

func (s *PushLiveActivityRequestActivityEventOption) SetSendDate(v int64) *PushLiveActivityRequestActivityEventOption {
	s.SendDate = &v
	return s
}

func (s *PushLiveActivityRequestActivityEventOption) SetStaleDate(v int64) *PushLiveActivityRequestActivityEventOption {
	s.StaleDate = &v
	return s
}

type PushLiveActivityResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s PushLiveActivityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushLiveActivityResponseBody) GoString() string {
	return s.String()
}

func (s *PushLiveActivityResponseBody) SetResult(v string) *PushLiveActivityResponseBody {
	s.Result = &v
	return s
}

type PushLiveActivityResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushLiveActivityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushLiveActivityResponse) String() string {
	return tea.Prettify(s)
}

func (s PushLiveActivityResponse) GoString() string {
	return s.String()
}

func (s *PushLiveActivityResponse) SetHeaders(v map[string]*string) *PushLiveActivityResponse {
	s.Headers = v
	return s
}

func (s *PushLiveActivityResponse) SetStatusCode(v int32) *PushLiveActivityResponse {
	s.StatusCode = &v
	return s
}

func (s *PushLiveActivityResponse) SetBody(v *PushLiveActivityResponseBody) *PushLiveActivityResponse {
	s.Body = v
	return s
}

type SendLiveActivityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendLiveActivityHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendLiveActivityHeaders) GoString() string {
	return s.String()
}

func (s *SendLiveActivityHeaders) SetCommonHeaders(v map[string]*string) *SendLiveActivityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendLiveActivityHeaders) SetXAcsDingtalkAccessToken(v string) *SendLiveActivityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendLiveActivityRequest struct {
	ActivityEventData   *SendLiveActivityRequestActivityEventData   `json:"activityEventData,omitempty" xml:"activityEventData,omitempty" type:"Struct"`
	ActivityEventOption *SendLiveActivityRequestActivityEventOption `json:"activityEventOption,omitempty" xml:"activityEventOption,omitempty" type:"Struct"`
	// example:
	//
	// bizUniqueId
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
}

func (s SendLiveActivityRequest) String() string {
	return tea.Prettify(s)
}

func (s SendLiveActivityRequest) GoString() string {
	return s.String()
}

func (s *SendLiveActivityRequest) SetActivityEventData(v *SendLiveActivityRequestActivityEventData) *SendLiveActivityRequest {
	s.ActivityEventData = v
	return s
}

func (s *SendLiveActivityRequest) SetActivityEventOption(v *SendLiveActivityRequestActivityEventOption) *SendLiveActivityRequest {
	s.ActivityEventOption = v
	return s
}

func (s *SendLiveActivityRequest) SetActivityId(v string) *SendLiveActivityRequest {
	s.ActivityId = &v
	return s
}

type SendLiveActivityRequestActivityEventData struct {
	I18nContentState interface{} `json:"i18nContentState,omitempty" xml:"i18nContentState,omitempty"`
	// example:
	//
	// ride_with_alibtrip
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s SendLiveActivityRequestActivityEventData) String() string {
	return tea.Prettify(s)
}

func (s SendLiveActivityRequestActivityEventData) GoString() string {
	return s.String()
}

func (s *SendLiveActivityRequestActivityEventData) SetI18nContentState(v interface{}) *SendLiveActivityRequestActivityEventData {
	s.I18nContentState = v
	return s
}

func (s *SendLiveActivityRequestActivityEventData) SetTemplateId(v string) *SendLiveActivityRequestActivityEventData {
	s.TemplateId = &v
	return s
}

type SendLiveActivityRequestActivityEventOption struct {
	// example:
	//
	// 1686903998
	DismissalDate *int64 `json:"dismissalDate,omitempty" xml:"dismissalDate,omitempty"`
	// example:
	//
	// update
	PushType *string `json:"pushType,omitempty" xml:"pushType,omitempty"`
	// example:
	//
	// 1686903998
	SendDate *int64 `json:"sendDate,omitempty" xml:"sendDate,omitempty"`
	// example:
	//
	// 1686903998
	StaleDate *int64 `json:"staleDate,omitempty" xml:"staleDate,omitempty"`
}

func (s SendLiveActivityRequestActivityEventOption) String() string {
	return tea.Prettify(s)
}

func (s SendLiveActivityRequestActivityEventOption) GoString() string {
	return s.String()
}

func (s *SendLiveActivityRequestActivityEventOption) SetDismissalDate(v int64) *SendLiveActivityRequestActivityEventOption {
	s.DismissalDate = &v
	return s
}

func (s *SendLiveActivityRequestActivityEventOption) SetPushType(v string) *SendLiveActivityRequestActivityEventOption {
	s.PushType = &v
	return s
}

func (s *SendLiveActivityRequestActivityEventOption) SetSendDate(v int64) *SendLiveActivityRequestActivityEventOption {
	s.SendDate = &v
	return s
}

func (s *SendLiveActivityRequestActivityEventOption) SetStaleDate(v int64) *SendLiveActivityRequestActivityEventOption {
	s.StaleDate = &v
	return s
}

type SendLiveActivityResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SendLiveActivityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendLiveActivityResponseBody) GoString() string {
	return s.String()
}

func (s *SendLiveActivityResponseBody) SetResult(v string) *SendLiveActivityResponseBody {
	s.Result = &v
	return s
}

type SendLiveActivityResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendLiveActivityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendLiveActivityResponse) String() string {
	return tea.Prettify(s)
}

func (s SendLiveActivityResponse) GoString() string {
	return s.String()
}

func (s *SendLiveActivityResponse) SetHeaders(v map[string]*string) *SendLiveActivityResponse {
	s.Headers = v
	return s
}

func (s *SendLiveActivityResponse) SetStatusCode(v int32) *SendLiveActivityResponse {
	s.StatusCode = &v
	return s
}

func (s *SendLiveActivityResponse) SetBody(v *SendLiveActivityResponseBody) *SendLiveActivityResponse {
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
// 实时活动发送接口
//
// @param request - PushLiveActivityRequest
//
// @param headers - PushLiveActivityHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushLiveActivityResponse
func (client *Client) PushLiveActivityWithOptions(request *PushLiveActivityRequest, headers *PushLiveActivityHeaders, runtime *util.RuntimeOptions) (_result *PushLiveActivityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivityEventData)) {
		body["activityEventData"] = request.ActivityEventData
	}

	if !tea.BoolValue(util.IsUnset(request.ActivityEventOption)) {
		body["activityEventOption"] = request.ActivityEventOption
	}

	if !tea.BoolValue(util.IsUnset(request.ActivityId)) {
		body["activityId"] = request.ActivityId
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
		Action:      tea.String("PushLiveActivity"),
		Version:     tea.String("liveActivities_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/liveActivities/push"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PushLiveActivityResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实时活动发送接口
//
// @param request - PushLiveActivityRequest
//
// @return PushLiveActivityResponse
func (client *Client) PushLiveActivity(request *PushLiveActivityRequest) (_result *PushLiveActivityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PushLiveActivityHeaders{}
	_result = &PushLiveActivityResponse{}
	_body, _err := client.PushLiveActivityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发送实时活动
//
// @param request - SendLiveActivityRequest
//
// @param headers - SendLiveActivityHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendLiveActivityResponse
func (client *Client) SendLiveActivityWithOptions(request *SendLiveActivityRequest, headers *SendLiveActivityHeaders, runtime *util.RuntimeOptions) (_result *SendLiveActivityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivityEventData)) {
		body["activityEventData"] = request.ActivityEventData
	}

	if !tea.BoolValue(util.IsUnset(request.ActivityEventOption)) {
		body["activityEventOption"] = request.ActivityEventOption
	}

	if !tea.BoolValue(util.IsUnset(request.ActivityId)) {
		body["activityId"] = request.ActivityId
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
		Action:      tea.String("SendLiveActivity"),
		Version:     tea.String("liveActivities_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/liveActivities/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendLiveActivityResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送实时活动
//
// @param request - SendLiveActivityRequest
//
// @return SendLiveActivityResponse
func (client *Client) SendLiveActivity(request *SendLiveActivityRequest) (_result *SendLiveActivityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendLiveActivityHeaders{}
	_result = &SendLiveActivityResponse{}
	_body, _err := client.SendLiveActivityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
