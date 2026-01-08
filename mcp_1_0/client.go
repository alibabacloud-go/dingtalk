// This file is auto-generated, don't edit it. Thanks.
package mcp_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ActivateMcpHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ActivateMcpHeaders) String() string {
	return tea.Prettify(s)
}

func (s ActivateMcpHeaders) GoString() string {
	return s.String()
}

func (s *ActivateMcpHeaders) SetCommonHeaders(v map[string]*string) *ActivateMcpHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ActivateMcpHeaders) SetXAcsDingtalkAccessToken(v string) *ActivateMcpHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ActivateMcpRequest struct {
	// This parameter is required.
	McpId  *int64  `json:"mcpId,omitempty" xml:"mcpId,omitempty"`
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s ActivateMcpRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivateMcpRequest) GoString() string {
	return s.String()
}

func (s *ActivateMcpRequest) SetMcpId(v int64) *ActivateMcpRequest {
	s.McpId = &v
	return s
}

func (s *ActivateMcpRequest) SetSource(v string) *ActivateMcpRequest {
	s.Source = &v
	return s
}

type ActivateMcpResponseBody struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	JsonConfig *string `json:"jsonConfig,omitempty" xml:"jsonConfig,omitempty"`
	Url        *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ActivateMcpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActivateMcpResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateMcpResponseBody) SetInstanceId(v string) *ActivateMcpResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ActivateMcpResponseBody) SetJsonConfig(v string) *ActivateMcpResponseBody {
	s.JsonConfig = &v
	return s
}

func (s *ActivateMcpResponseBody) SetUrl(v string) *ActivateMcpResponseBody {
	s.Url = &v
	return s
}

type ActivateMcpResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActivateMcpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateMcpResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivateMcpResponse) GoString() string {
	return s.String()
}

func (s *ActivateMcpResponse) SetHeaders(v map[string]*string) *ActivateMcpResponse {
	s.Headers = v
	return s
}

func (s *ActivateMcpResponse) SetStatusCode(v int32) *ActivateMcpResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateMcpResponse) SetBody(v *ActivateMcpResponseBody) *ActivateMcpResponse {
	s.Body = v
	return s
}

type DeleteMcpHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteMcpHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteMcpHeaders) GoString() string {
	return s.String()
}

func (s *DeleteMcpHeaders) SetCommonHeaders(v map[string]*string) *DeleteMcpHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteMcpHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteMcpHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteMcpRequest struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s DeleteMcpRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMcpRequest) GoString() string {
	return s.String()
}

func (s *DeleteMcpRequest) SetInstanceId(v string) *DeleteMcpRequest {
	s.InstanceId = &v
	return s
}

type DeleteMcpResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteMcpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMcpResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMcpResponseBody) SetSuccess(v bool) *DeleteMcpResponseBody {
	s.Success = &v
	return s
}

type DeleteMcpResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMcpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMcpResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMcpResponse) GoString() string {
	return s.String()
}

func (s *DeleteMcpResponse) SetHeaders(v map[string]*string) *DeleteMcpResponse {
	s.Headers = v
	return s
}

func (s *DeleteMcpResponse) SetStatusCode(v int32) *DeleteMcpResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMcpResponse) SetBody(v *DeleteMcpResponseBody) *DeleteMcpResponse {
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
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

// Summary:
//
// 激活MCP
//
// @param request - ActivateMcpRequest
//
// @param headers - ActivateMcpHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateMcpResponse
func (client *Client) ActivateMcpWithOptions(request *ActivateMcpRequest, headers *ActivateMcpHeaders, runtime *util.RuntimeOptions) (_result *ActivateMcpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.McpId)) {
		body["mcpId"] = request.McpId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
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
		Action:      tea.String("ActivateMcp"),
		Version:     tea.String("mcp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/mcp/activate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ActivateMcpResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 激活MCP
//
// @param request - ActivateMcpRequest
//
// @return ActivateMcpResponse
func (client *Client) ActivateMcp(request *ActivateMcpRequest) (_result *ActivateMcpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ActivateMcpHeaders{}
	_result = &ActivateMcpResponse{}
	_body, _err := client.ActivateMcpWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除MCP实例
//
// @param request - DeleteMcpRequest
//
// @param headers - DeleteMcpHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMcpResponse
func (client *Client) DeleteMcpWithOptions(request *DeleteMcpRequest, headers *DeleteMcpHeaders, runtime *util.RuntimeOptions) (_result *DeleteMcpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
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
		Action:      tea.String("DeleteMcp"),
		Version:     tea.String("mcp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/mcp/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMcpResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除MCP实例
//
// @param request - DeleteMcpRequest
//
// @return DeleteMcpResponse
func (client *Client) DeleteMcp(request *DeleteMcpRequest) (_result *DeleteMcpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteMcpHeaders{}
	_result = &DeleteMcpResponse{}
	_body, _err := client.DeleteMcpWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
