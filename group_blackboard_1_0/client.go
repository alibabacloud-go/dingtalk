// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package group_blackboard_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateGroupBlackboardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateGroupBlackboardHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupBlackboardHeaders) GoString() string {
	return s.String()
}

func (s *CreateGroupBlackboardHeaders) SetCommonHeaders(v map[string]*string) *CreateGroupBlackboardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateGroupBlackboardHeaders) SetXAcsDingtalkAccessToken(v string) *CreateGroupBlackboardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateGroupBlackboardRequest struct {
	Content            *string `json:"content,omitempty" xml:"content,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	SendDing           *bool   `json:"sendDing,omitempty" xml:"sendDing,omitempty"`
	Sticky             *bool   `json:"sticky,omitempty" xml:"sticky,omitempty"`
	UniqueId           *string `json:"uniqueId,omitempty" xml:"uniqueId,omitempty"`
	UserId             *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateGroupBlackboardRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupBlackboardRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupBlackboardRequest) SetContent(v string) *CreateGroupBlackboardRequest {
	s.Content = &v
	return s
}

func (s *CreateGroupBlackboardRequest) SetOpenConversationId(v string) *CreateGroupBlackboardRequest {
	s.OpenConversationId = &v
	return s
}

func (s *CreateGroupBlackboardRequest) SetSendDing(v bool) *CreateGroupBlackboardRequest {
	s.SendDing = &v
	return s
}

func (s *CreateGroupBlackboardRequest) SetSticky(v bool) *CreateGroupBlackboardRequest {
	s.Sticky = &v
	return s
}

func (s *CreateGroupBlackboardRequest) SetUniqueId(v string) *CreateGroupBlackboardRequest {
	s.UniqueId = &v
	return s
}

func (s *CreateGroupBlackboardRequest) SetUserId(v string) *CreateGroupBlackboardRequest {
	s.UserId = &v
	return s
}

type CreateGroupBlackboardResponseBody struct {
	DataId  *string `json:"dataId,omitempty" xml:"dataId,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateGroupBlackboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupBlackboardResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupBlackboardResponseBody) SetDataId(v string) *CreateGroupBlackboardResponseBody {
	s.DataId = &v
	return s
}

func (s *CreateGroupBlackboardResponseBody) SetSuccess(v bool) *CreateGroupBlackboardResponseBody {
	s.Success = &v
	return s
}

type CreateGroupBlackboardResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGroupBlackboardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGroupBlackboardResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupBlackboardResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupBlackboardResponse) SetHeaders(v map[string]*string) *CreateGroupBlackboardResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupBlackboardResponse) SetStatusCode(v int32) *CreateGroupBlackboardResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupBlackboardResponse) SetBody(v *CreateGroupBlackboardResponseBody) *CreateGroupBlackboardResponse {
	s.Body = v
	return s
}

type DeleteGroupBlackboardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteGroupBlackboardHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupBlackboardHeaders) GoString() string {
	return s.String()
}

func (s *DeleteGroupBlackboardHeaders) SetCommonHeaders(v map[string]*string) *DeleteGroupBlackboardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteGroupBlackboardHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteGroupBlackboardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteGroupBlackboardRequest struct {
	DataId             *string `json:"dataId,omitempty" xml:"dataId,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	UserId             *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeleteGroupBlackboardRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupBlackboardRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupBlackboardRequest) SetDataId(v string) *DeleteGroupBlackboardRequest {
	s.DataId = &v
	return s
}

func (s *DeleteGroupBlackboardRequest) SetOpenConversationId(v string) *DeleteGroupBlackboardRequest {
	s.OpenConversationId = &v
	return s
}

func (s *DeleteGroupBlackboardRequest) SetUserId(v string) *DeleteGroupBlackboardRequest {
	s.UserId = &v
	return s
}

type DeleteGroupBlackboardResponseBody struct {
	IsDeleted *bool `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	Success   *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteGroupBlackboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupBlackboardResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupBlackboardResponseBody) SetIsDeleted(v bool) *DeleteGroupBlackboardResponseBody {
	s.IsDeleted = &v
	return s
}

func (s *DeleteGroupBlackboardResponseBody) SetSuccess(v bool) *DeleteGroupBlackboardResponseBody {
	s.Success = &v
	return s
}

type DeleteGroupBlackboardResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGroupBlackboardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGroupBlackboardResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupBlackboardResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupBlackboardResponse) SetHeaders(v map[string]*string) *DeleteGroupBlackboardResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupBlackboardResponse) SetStatusCode(v int32) *DeleteGroupBlackboardResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGroupBlackboardResponse) SetBody(v *DeleteGroupBlackboardResponseBody) *DeleteGroupBlackboardResponse {
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

func (client *Client) CreateGroupBlackboardWithOptions(request *CreateGroupBlackboardRequest, headers *CreateGroupBlackboardHeaders, runtime *util.RuntimeOptions) (_result *CreateGroupBlackboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.SendDing)) {
		body["sendDing"] = request.SendDing
	}

	if !tea.BoolValue(util.IsUnset(request.Sticky)) {
		body["sticky"] = request.Sticky
	}

	if !tea.BoolValue(util.IsUnset(request.UniqueId)) {
		body["uniqueId"] = request.UniqueId
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
		Action:      tea.String("CreateGroupBlackboard"),
		Version:     tea.String("groupBlackboard_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/groupBlackboard/blackboards"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGroupBlackboardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGroupBlackboard(request *CreateGroupBlackboardRequest) (_result *CreateGroupBlackboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateGroupBlackboardHeaders{}
	_result = &CreateGroupBlackboardResponse{}
	_body, _err := client.CreateGroupBlackboardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGroupBlackboardWithOptions(request *DeleteGroupBlackboardRequest, headers *DeleteGroupBlackboardHeaders, runtime *util.RuntimeOptions) (_result *DeleteGroupBlackboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataId)) {
		body["dataId"] = request.DataId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
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
		Action:      tea.String("DeleteGroupBlackboard"),
		Version:     tea.String("groupBlackboard_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/groupBlackboard/blackboards/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGroupBlackboardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGroupBlackboard(request *DeleteGroupBlackboardRequest) (_result *DeleteGroupBlackboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteGroupBlackboardHeaders{}
	_result = &DeleteGroupBlackboardResponse{}
	_body, _err := client.DeleteGroupBlackboardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
