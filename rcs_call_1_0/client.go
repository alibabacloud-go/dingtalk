// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package rcs_call_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
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
	// 授权isv套件企业的corpid
	AuthorizeCorpId *string `json:"authorizeCorpId,omitempty" xml:"authorizeCorpId,omitempty"`
	// 授权isv套件企业的员工userid
	AuthorizeUserId *string `json:"authorizeUserId,omitempty" xml:"authorizeUserId,omitempty"`
	// 订单id
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// isv套件所属企业下的员工userid
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	// 执行结果
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RunCallUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
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
	_result = &RunCallUserResponse{}
	_body, _err := client.DoROARequest(tea.String("RunCallUser"), tea.String("rcsCall_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/rcsCall/users/call"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
