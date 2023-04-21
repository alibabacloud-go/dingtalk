// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package gateway_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type OpenConnectionRequest struct {
	// 企业三方应用为suiteKey
	// 企业自建应用为appKey
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// 企业三方应用为suiteSecret
	// 企业自己应用为appSecret
	ClientSecret  *string                               `json:"clientSecret,omitempty" xml:"clientSecret,omitempty"`
	Subscriptions []*OpenConnectionRequestSubscriptions `json:"subscriptions,omitempty" xml:"subscriptions,omitempty" type:"Repeated"`
}

func (s OpenConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenConnectionRequest) GoString() string {
	return s.String()
}

func (s *OpenConnectionRequest) SetClientId(v string) *OpenConnectionRequest {
	s.ClientId = &v
	return s
}

func (s *OpenConnectionRequest) SetClientSecret(v string) *OpenConnectionRequest {
	s.ClientSecret = &v
	return s
}

func (s *OpenConnectionRequest) SetSubscriptions(v []*OpenConnectionRequestSubscriptions) *OpenConnectionRequest {
	s.Subscriptions = v
	return s
}

type OpenConnectionRequestSubscriptions struct {
	// 订阅的TOPIC
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
	// 订阅类型
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s OpenConnectionRequestSubscriptions) String() string {
	return tea.Prettify(s)
}

func (s OpenConnectionRequestSubscriptions) GoString() string {
	return s.String()
}

func (s *OpenConnectionRequestSubscriptions) SetTopic(v string) *OpenConnectionRequestSubscriptions {
	s.Topic = &v
	return s
}

func (s *OpenConnectionRequestSubscriptions) SetType(v string) *OpenConnectionRequestSubscriptions {
	s.Type = &v
	return s
}

type OpenConnectionResponseBody struct {
	// 长连接接入点
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// 连接ticket
	Ticket *string `json:"ticket,omitempty" xml:"ticket,omitempty"`
}

func (s OpenConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *OpenConnectionResponseBody) SetEndpoint(v string) *OpenConnectionResponseBody {
	s.Endpoint = &v
	return s
}

func (s *OpenConnectionResponseBody) SetTicket(v string) *OpenConnectionResponseBody {
	s.Ticket = &v
	return s
}

type OpenConnectionResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenConnectionResponse) GoString() string {
	return s.String()
}

func (s *OpenConnectionResponse) SetHeaders(v map[string]*string) *OpenConnectionResponse {
	s.Headers = v
	return s
}

func (s *OpenConnectionResponse) SetBody(v *OpenConnectionResponseBody) *OpenConnectionResponse {
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

func (client *Client) OpenConnection(request *OpenConnectionRequest) (_result *OpenConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OpenConnectionResponse{}
	_body, _err := client.OpenConnectionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenConnectionWithOptions(request *OpenConnectionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *OpenConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		body["clientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSecret)) {
		body["clientSecret"] = request.ClientSecret
	}

	if !tea.BoolValue(util.IsUnset(request.Subscriptions)) {
		body["subscriptions"] = request.Subscriptions
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &OpenConnectionResponse{}
	_body, _err := client.DoROARequest(tea.String("OpenConnection"), tea.String("gateway_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/gateway/connections/open"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
