// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package occupationauth_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CheckUserTaskStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CheckUserTaskStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s CheckUserTaskStatusHeaders) GoString() string {
	return s.String()
}

func (s *CheckUserTaskStatusHeaders) SetCommonHeaders(v map[string]*string) *CheckUserTaskStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CheckUserTaskStatusHeaders) SetXAcsDingtalkAccessToken(v string) *CheckUserTaskStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CheckUserTaskStatusRequest struct {
	ProvinceCode *string `json:"provinceCode,omitempty" xml:"provinceCode,omitempty"`
}

func (s CheckUserTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckUserTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckUserTaskStatusRequest) SetProvinceCode(v string) *CheckUserTaskStatusRequest {
	s.ProvinceCode = &v
	return s
}

type CheckUserTaskStatusResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CheckUserTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckUserTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUserTaskStatusResponseBody) SetResult(v bool) *CheckUserTaskStatusResponseBody {
	s.Result = &v
	return s
}

type CheckUserTaskStatusResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckUserTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckUserTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckUserTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckUserTaskStatusResponse) SetHeaders(v map[string]*string) *CheckUserTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckUserTaskStatusResponse) SetBody(v *CheckUserTaskStatusResponseBody) *CheckUserTaskStatusResponse {
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

func (client *Client) CheckUserTaskStatus(request *CheckUserTaskStatusRequest) (_result *CheckUserTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CheckUserTaskStatusHeaders{}
	_result = &CheckUserTaskStatusResponse{}
	_body, _err := client.CheckUserTaskStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckUserTaskStatusWithOptions(request *CheckUserTaskStatusRequest, headers *CheckUserTaskStatusHeaders, runtime *util.RuntimeOptions) (_result *CheckUserTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProvinceCode)) {
		body["provinceCode"] = request.ProvinceCode
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CheckUserTaskStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("CheckUserTaskStatus"), tea.String("occupationauth_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/occupationauth/auths/userTasks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
