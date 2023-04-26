// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package blackboard_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type QueryBlackboardSpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryBlackboardSpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardSpaceHeaders) GoString() string {
	return s.String()
}

func (s *QueryBlackboardSpaceHeaders) SetCommonHeaders(v map[string]*string) *QueryBlackboardSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryBlackboardSpaceHeaders) SetXAcsDingtalkAccessToken(v string) *QueryBlackboardSpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryBlackboardSpaceRequest struct {
	OperationUserId *string `json:"operationUserId,omitempty" xml:"operationUserId,omitempty"`
}

func (s QueryBlackboardSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardSpaceRequest) GoString() string {
	return s.String()
}

func (s *QueryBlackboardSpaceRequest) SetOperationUserId(v string) *QueryBlackboardSpaceRequest {
	s.OperationUserId = &v
	return s
}

type QueryBlackboardSpaceResponseBody struct {
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s QueryBlackboardSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBlackboardSpaceResponseBody) SetSpaceId(v string) *QueryBlackboardSpaceResponseBody {
	s.SpaceId = &v
	return s
}

type QueryBlackboardSpaceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryBlackboardSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBlackboardSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardSpaceResponse) GoString() string {
	return s.String()
}

func (s *QueryBlackboardSpaceResponse) SetHeaders(v map[string]*string) *QueryBlackboardSpaceResponse {
	s.Headers = v
	return s
}

func (s *QueryBlackboardSpaceResponse) SetStatusCode(v int32) *QueryBlackboardSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBlackboardSpaceResponse) SetBody(v *QueryBlackboardSpaceResponseBody) *QueryBlackboardSpaceResponse {
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

func (client *Client) QueryBlackboardSpaceWithOptions(request *QueryBlackboardSpaceRequest, headers *QueryBlackboardSpaceHeaders, runtime *util.RuntimeOptions) (_result *QueryBlackboardSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperationUserId)) {
		query["operationUserId"] = request.OperationUserId
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
		Action:      tea.String("QueryBlackboardSpace"),
		Version:     tea.String("blackboard_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/blackboard/spaces"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBlackboardSpaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBlackboardSpace(request *QueryBlackboardSpaceRequest) (_result *QueryBlackboardSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryBlackboardSpaceHeaders{}
	_result = &QueryBlackboardSpaceResponse{}
	_body, _err := client.QueryBlackboardSpaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
