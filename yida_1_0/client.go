// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package yida_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type LoginCodeGenHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s LoginCodeGenHeaders) String() string {
	return tea.Prettify(s)
}

func (s LoginCodeGenHeaders) GoString() string {
	return s.String()
}

func (s *LoginCodeGenHeaders) SetCommonHeaders(v map[string]*string) *LoginCodeGenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *LoginCodeGenHeaders) SetXAcsDingtalkAccessToken(v string) *LoginCodeGenHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type LoginCodeGenRequest struct {
	// userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s LoginCodeGenRequest) String() string {
	return tea.Prettify(s)
}

func (s LoginCodeGenRequest) GoString() string {
	return s.String()
}

func (s *LoginCodeGenRequest) SetUserId(v string) *LoginCodeGenRequest {
	s.UserId = &v
	return s
}

type LoginCodeGenResponseBody struct {
	// result
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s LoginCodeGenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LoginCodeGenResponseBody) GoString() string {
	return s.String()
}

func (s *LoginCodeGenResponseBody) SetResult(v string) *LoginCodeGenResponseBody {
	s.Result = &v
	return s
}

type LoginCodeGenResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *LoginCodeGenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LoginCodeGenResponse) String() string {
	return tea.Prettify(s)
}

func (s LoginCodeGenResponse) GoString() string {
	return s.String()
}

func (s *LoginCodeGenResponse) SetHeaders(v map[string]*string) *LoginCodeGenResponse {
	s.Headers = v
	return s
}

func (s *LoginCodeGenResponse) SetBody(v *LoginCodeGenResponseBody) *LoginCodeGenResponse {
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

func (client *Client) LoginCodeGen(request *LoginCodeGenRequest) (_result *LoginCodeGenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &LoginCodeGenHeaders{}
	_result = &LoginCodeGenResponse{}
	_body, _err := client.LoginCodeGenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LoginCodeGenWithOptions(request *LoginCodeGenRequest, headers *LoginCodeGenHeaders, runtime *util.RuntimeOptions) (_result *LoginCodeGenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
		Query:   openapiutil.Query(query),
	}
	_result = &LoginCodeGenResponse{}
	_body, _err := client.DoROARequest(tea.String("LoginCodeGen"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/authorizations/loginCodes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
