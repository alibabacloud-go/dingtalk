// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package project_integration_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddAttendeeToEventGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddAttendeeToEventGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddAttendeeToEventGroupHeaders) GoString() string {
	return s.String()
}

func (s *AddAttendeeToEventGroupHeaders) SetCommonHeaders(v map[string]*string) *AddAttendeeToEventGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddAttendeeToEventGroupHeaders) SetXAcsDingtalkAccessToken(v string) *AddAttendeeToEventGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddAttendeeToEventGroupResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddAttendeeToEventGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAttendeeToEventGroupResponse) GoString() string {
	return s.String()
}

func (s *AddAttendeeToEventGroupResponse) SetHeaders(v map[string]*string) *AddAttendeeToEventGroupResponse {
	s.Headers = v
	return s
}

func (s *AddAttendeeToEventGroupResponse) SetBody(v map[string]interface{}) *AddAttendeeToEventGroupResponse {
	s.Body = v
	return s
}

type CreateEventGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateEventGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateEventGroupHeaders) GoString() string {
	return s.String()
}

func (s *CreateEventGroupHeaders) SetCommonHeaders(v map[string]*string) *CreateEventGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateEventGroupHeaders) SetXAcsDingtalkAccessToken(v string) *CreateEventGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateEventGroupResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEventGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEventGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateEventGroupResponse) SetHeaders(v map[string]*string) *CreateEventGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateEventGroupResponse) SetBody(v map[string]interface{}) *CreateEventGroupResponse {
	s.Body = v
	return s
}

type SendInteractiveCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendInteractiveCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendInteractiveCardHeaders) GoString() string {
	return s.String()
}

func (s *SendInteractiveCardHeaders) SetCommonHeaders(v map[string]*string) *SendInteractiveCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendInteractiveCardHeaders) SetXAcsDingtalkAccessToken(v string) *SendInteractiveCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendInteractiveCardResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendInteractiveCardResponse) String() string {
	return tea.Prettify(s)
}

func (s SendInteractiveCardResponse) GoString() string {
	return s.String()
}

func (s *SendInteractiveCardResponse) SetHeaders(v map[string]*string) *SendInteractiveCardResponse {
	s.Headers = v
	return s
}

func (s *SendInteractiveCardResponse) SetBody(v map[string]interface{}) *SendInteractiveCardResponse {
	s.Body = v
	return s
}

type SendSingleInteractiveCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendSingleInteractiveCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendSingleInteractiveCardHeaders) GoString() string {
	return s.String()
}

func (s *SendSingleInteractiveCardHeaders) SetCommonHeaders(v map[string]*string) *SendSingleInteractiveCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendSingleInteractiveCardHeaders) SetXAcsDingtalkAccessToken(v string) *SendSingleInteractiveCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendSingleInteractiveCardResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendSingleInteractiveCardResponse) String() string {
	return tea.Prettify(s)
}

func (s SendSingleInteractiveCardResponse) GoString() string {
	return s.String()
}

func (s *SendSingleInteractiveCardResponse) SetHeaders(v map[string]*string) *SendSingleInteractiveCardResponse {
	s.Headers = v
	return s
}

func (s *SendSingleInteractiveCardResponse) SetBody(v map[string]interface{}) *SendSingleInteractiveCardResponse {
	s.Body = v
	return s
}

type UpdateInteractiveCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateInteractiveCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateInteractiveCardHeaders) GoString() string {
	return s.String()
}

func (s *UpdateInteractiveCardHeaders) SetCommonHeaders(v map[string]*string) *UpdateInteractiveCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateInteractiveCardHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateInteractiveCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateInteractiveCardResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInteractiveCardResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInteractiveCardResponse) GoString() string {
	return s.String()
}

func (s *UpdateInteractiveCardResponse) SetHeaders(v map[string]*string) *UpdateInteractiveCardResponse {
	s.Headers = v
	return s
}

func (s *UpdateInteractiveCardResponse) SetBody(v map[string]interface{}) *UpdateInteractiveCardResponse {
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

func (client *Client) AddAttendeeToEventGroup(userId *string, groupId *string) (_result *AddAttendeeToEventGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddAttendeeToEventGroupHeaders{}
	_result = &AddAttendeeToEventGroupResponse{}
	_body, _err := client.AddAttendeeToEventGroupWithOptions(userId, groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddAttendeeToEventGroupWithOptions(userId *string, groupId *string, headers *AddAttendeeToEventGroupHeaders, runtime *util.RuntimeOptions) (_result *AddAttendeeToEventGroupResponse, _err error) {
	userId = openapiutil.GetEncodeParam(userId)
	groupId = openapiutil.GetEncodeParam(groupId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &AddAttendeeToEventGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("AddAttendeeToEventGroup"), tea.String("projectIntegration_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/projectIntegration/users/"+tea.StringValue(userId)+"/eventGroups/"+tea.StringValue(groupId)+"/members"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEventGroup(userId *string) (_result *CreateEventGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateEventGroupHeaders{}
	_result = &CreateEventGroupResponse{}
	_body, _err := client.CreateEventGroupWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEventGroupWithOptions(userId *string, headers *CreateEventGroupHeaders, runtime *util.RuntimeOptions) (_result *CreateEventGroupResponse, _err error) {
	userId = openapiutil.GetEncodeParam(userId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &CreateEventGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateEventGroup"), tea.String("projectIntegration_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/projectIntegration/users/"+tea.StringValue(userId)+"/eventGroups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendInteractiveCard(userId *string) (_result *SendInteractiveCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendInteractiveCardHeaders{}
	_result = &SendInteractiveCardResponse{}
	_body, _err := client.SendInteractiveCardWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendInteractiveCardWithOptions(userId *string, headers *SendInteractiveCardHeaders, runtime *util.RuntimeOptions) (_result *SendInteractiveCardResponse, _err error) {
	userId = openapiutil.GetEncodeParam(userId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &SendInteractiveCardResponse{}
	_body, _err := client.DoROARequest(tea.String("SendInteractiveCard"), tea.String("projectIntegration_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/projectIntegration/users/"+tea.StringValue(userId)+"/groupChatCardMessages"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendSingleInteractiveCard(userId *string) (_result *SendSingleInteractiveCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendSingleInteractiveCardHeaders{}
	_result = &SendSingleInteractiveCardResponse{}
	_body, _err := client.SendSingleInteractiveCardWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendSingleInteractiveCardWithOptions(userId *string, headers *SendSingleInteractiveCardHeaders, runtime *util.RuntimeOptions) (_result *SendSingleInteractiveCardResponse, _err error) {
	userId = openapiutil.GetEncodeParam(userId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &SendSingleInteractiveCardResponse{}
	_body, _err := client.DoROARequest(tea.String("SendSingleInteractiveCard"), tea.String("projectIntegration_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/projectIntegration/users/"+tea.StringValue(userId)+"/singleChatCardMessages"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInteractiveCard(userId *string) (_result *UpdateInteractiveCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateInteractiveCardHeaders{}
	_result = &UpdateInteractiveCardResponse{}
	_body, _err := client.UpdateInteractiveCardWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInteractiveCardWithOptions(userId *string, headers *UpdateInteractiveCardHeaders, runtime *util.RuntimeOptions) (_result *UpdateInteractiveCardResponse, _err error) {
	userId = openapiutil.GetEncodeParam(userId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &UpdateInteractiveCardResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateInteractiveCard"), tea.String("projectIntegration_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/projectIntegration/users/"+tea.StringValue(userId)+"/cardMessages"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
