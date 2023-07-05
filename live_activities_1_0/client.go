// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package live_activities_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	ActivityId          *string                                     `json:"activityId,omitempty" xml:"activityId,omitempty"`
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
	TemplateId       *string     `json:"templateId,omitempty" xml:"templateId,omitempty"`
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
	DismissalDate *int64  `json:"dismissalDate,omitempty" xml:"dismissalDate,omitempty"`
	PushType      *string `json:"pushType,omitempty" xml:"pushType,omitempty"`
	SendDate      *int64  `json:"sendDate,omitempty" xml:"sendDate,omitempty"`
	StaleDate     *int64  `json:"staleDate,omitempty" xml:"staleDate,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendLiveActivityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
