// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package rcs_call_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type RunCallUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RunCallUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s RunCallUserHeaders) GoString() string {
	return s.String()
}

func (s *RunCallUserHeaders) SetCommonHeaders(v map[string]*string) *RunCallUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RunCallUserHeaders) SetXAcsDingtalkAccessToken(v string) *RunCallUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RunCallUserRequest struct {
	AuthorizeCorpId *string `json:"authorizeCorpId,omitempty" xml:"authorizeCorpId,omitempty"`
	AuthorizeUserId *string `json:"authorizeUserId,omitempty" xml:"authorizeUserId,omitempty"`
	OrderId         *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	UserId          *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s RunCallUserRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCallUserRequest) GoString() string {
	return s.String()
}

func (s *RunCallUserRequest) SetAuthorizeCorpId(v string) *RunCallUserRequest {
	s.AuthorizeCorpId = &v
	return s
}

func (s *RunCallUserRequest) SetAuthorizeUserId(v string) *RunCallUserRequest {
	s.AuthorizeUserId = &v
	return s
}

func (s *RunCallUserRequest) SetOrderId(v string) *RunCallUserRequest {
	s.OrderId = &v
	return s
}

func (s *RunCallUserRequest) SetUserId(v string) *RunCallUserRequest {
	s.UserId = &v
	return s
}

type RunCallUserResponseBody struct {
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RunCallUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunCallUserResponseBody) GoString() string {
	return s.String()
}

func (s *RunCallUserResponseBody) SetSuccess(v string) *RunCallUserResponseBody {
	s.Success = &v
	return s
}

type RunCallUserResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunCallUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunCallUserResponse) String() string {
	return tea.Prettify(s)
}

func (s RunCallUserResponse) GoString() string {
	return s.String()
}

func (s *RunCallUserResponse) SetHeaders(v map[string]*string) *RunCallUserResponse {
	s.Headers = v
	return s
}

func (s *RunCallUserResponse) SetStatusCode(v int32) *RunCallUserResponse {
	s.StatusCode = &v
	return s
}

func (s *RunCallUserResponse) SetBody(v *RunCallUserResponseBody) *RunCallUserResponse {
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

func (client *Client) RunCallUserWithOptions(request *RunCallUserRequest, headers *RunCallUserHeaders, runtime *util.RuntimeOptions) (_result *RunCallUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthorizeCorpId)) {
		query["authorizeCorpId"] = request.AuthorizeCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizeUserId)) {
		query["authorizeUserId"] = request.AuthorizeUserId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["orderId"] = request.OrderId
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
		Action:      tea.String("RunCallUser"),
		Version:     tea.String("rcsCall_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rcsCall/users/call"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RunCallUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunCallUser(request *RunCallUserRequest) (_result *RunCallUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RunCallUserHeaders{}
	_result = &RunCallUserResponse{}
	_body, _err := client.RunCallUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
