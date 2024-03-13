// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package ai_interaction_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ReplyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ReplyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ReplyHeaders) GoString() string {
	return s.String()
}

func (s *ReplyHeaders) SetCommonHeaders(v map[string]*string) *ReplyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReplyHeaders) SetXAcsDingtalkAccessToken(v string) *ReplyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ReplyRequest struct {
	Content           *string `json:"content,omitempty" xml:"content,omitempty"`
	ContentType       *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	ConversationToken *string `json:"conversationToken,omitempty" xml:"conversationToken,omitempty"`
}

func (s ReplyRequest) String() string {
	return tea.Prettify(s)
}

func (s ReplyRequest) GoString() string {
	return s.String()
}

func (s *ReplyRequest) SetContent(v string) *ReplyRequest {
	s.Content = &v
	return s
}

func (s *ReplyRequest) SetContentType(v string) *ReplyRequest {
	s.ContentType = &v
	return s
}

func (s *ReplyRequest) SetConversationToken(v string) *ReplyRequest {
	s.ConversationToken = &v
	return s
}

type ReplyResponseBody struct {
	Result *ReplyResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ReplyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReplyResponseBody) GoString() string {
	return s.String()
}

func (s *ReplyResponseBody) SetResult(v *ReplyResponseBodyResult) *ReplyResponseBody {
	s.Result = v
	return s
}

type ReplyResponseBodyResult struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ReplyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ReplyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ReplyResponseBodyResult) SetSuccess(v bool) *ReplyResponseBodyResult {
	s.Success = &v
	return s
}

type ReplyResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReplyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReplyResponse) String() string {
	return tea.Prettify(s)
}

func (s ReplyResponse) GoString() string {
	return s.String()
}

func (s *ReplyResponse) SetHeaders(v map[string]*string) *ReplyResponse {
	s.Headers = v
	return s
}

func (s *ReplyResponse) SetStatusCode(v int32) *ReplyResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplyResponse) SetBody(v *ReplyResponseBody) *ReplyResponse {
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

func (client *Client) ReplyWithOptions(request *ReplyRequest, headers *ReplyHeaders, runtime *util.RuntimeOptions) (_result *ReplyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ContentType)) {
		body["contentType"] = request.ContentType
	}

	if !tea.BoolValue(util.IsUnset(request.ConversationToken)) {
		body["conversationToken"] = request.ConversationToken
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
		Action:      tea.String("Reply"),
		Version:     tea.String("aiInteraction_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiInteraction/reply"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ReplyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Reply(request *ReplyRequest) (_result *ReplyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ReplyHeaders{}
	_result = &ReplyResponse{}
	_body, _err := client.ReplyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
