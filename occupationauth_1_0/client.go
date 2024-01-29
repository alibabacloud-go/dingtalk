// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package occupationauth_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckUserTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CheckUserTaskStatusResponse) SetStatusCode(v int32) *CheckUserTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckUserTaskStatusResponse) SetBody(v *CheckUserTaskStatusResponseBody) *CheckUserTaskStatusResponse {
	s.Body = v
	return s
}

type CheckUserTasksStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CheckUserTasksStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s CheckUserTasksStatusHeaders) GoString() string {
	return s.String()
}

func (s *CheckUserTasksStatusHeaders) SetCommonHeaders(v map[string]*string) *CheckUserTasksStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CheckUserTasksStatusHeaders) SetXAcsDingtalkAccessToken(v string) *CheckUserTasksStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CheckUserTasksStatusRequest struct {
	ProvinceCode *string `json:"provinceCode,omitempty" xml:"provinceCode,omitempty"`
}

func (s CheckUserTasksStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckUserTasksStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckUserTasksStatusRequest) SetProvinceCode(v string) *CheckUserTasksStatusRequest {
	s.ProvinceCode = &v
	return s
}

type CheckUserTasksStatusResponseBody struct {
	Status *bool `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CheckUserTasksStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckUserTasksStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUserTasksStatusResponseBody) SetStatus(v bool) *CheckUserTasksStatusResponseBody {
	s.Status = &v
	return s
}

type CheckUserTasksStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckUserTasksStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckUserTasksStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckUserTasksStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckUserTasksStatusResponse) SetHeaders(v map[string]*string) *CheckUserTasksStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckUserTasksStatusResponse) SetStatusCode(v int32) *CheckUserTasksStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckUserTasksStatusResponse) SetBody(v *CheckUserTasksStatusResponseBody) *CheckUserTasksStatusResponse {
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckUserTaskStatus"),
		Version:     tea.String("occupationauth_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/occupationauth/auths/userTasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckUserTaskStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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

func (client *Client) CheckUserTasksStatusWithOptions(request *CheckUserTasksStatusRequest, headers *CheckUserTasksStatusHeaders, runtime *util.RuntimeOptions) (_result *CheckUserTasksStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProvinceCode)) {
		query["provinceCode"] = request.ProvinceCode
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
		Action:      tea.String("CheckUserTasksStatus"),
		Version:     tea.String("occupationauth_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/occupationauth/userTasks/check"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckUserTasksStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckUserTasksStatus(request *CheckUserTasksStatusRequest) (_result *CheckUserTasksStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CheckUserTasksStatusHeaders{}
	_result = &CheckUserTasksStatusResponse{}
	_body, _err := client.CheckUserTasksStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
