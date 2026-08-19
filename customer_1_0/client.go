// This file is auto-generated, don't edit it. Thanks.
package customer_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CustomeRpcCallRequest struct {
	MethodName *string                `json:"methodName,omitempty" xml:"methodName,omitempty"`
	Params     map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
}

func (s CustomeRpcCallRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomeRpcCallRequest) GoString() string {
	return s.String()
}

func (s *CustomeRpcCallRequest) SetMethodName(v string) *CustomeRpcCallRequest {
	s.MethodName = &v
	return s
}

func (s *CustomeRpcCallRequest) SetParams(v map[string]interface{}) *CustomeRpcCallRequest {
	s.Params = v
	return s
}

type CustomeRpcCallShrinkRequest struct {
	MethodName   *string `json:"methodName,omitempty" xml:"methodName,omitempty"`
	ParamsShrink *string `json:"params,omitempty" xml:"params,omitempty"`
}

func (s CustomeRpcCallShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomeRpcCallShrinkRequest) GoString() string {
	return s.String()
}

func (s *CustomeRpcCallShrinkRequest) SetMethodName(v string) *CustomeRpcCallShrinkRequest {
	s.MethodName = &v
	return s
}

func (s *CustomeRpcCallShrinkRequest) SetParamsShrink(v string) *CustomeRpcCallShrinkRequest {
	s.ParamsShrink = &v
	return s
}

type CustomeRpcCallResponseBody struct {
	ErrorCode *string                  `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                  `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result    []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success   *string                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CustomeRpcCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomeRpcCallResponseBody) GoString() string {
	return s.String()
}

func (s *CustomeRpcCallResponseBody) SetErrorCode(v string) *CustomeRpcCallResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CustomeRpcCallResponseBody) SetErrorMsg(v string) *CustomeRpcCallResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CustomeRpcCallResponseBody) SetResult(v []map[string]interface{}) *CustomeRpcCallResponseBody {
	s.Result = v
	return s
}

func (s *CustomeRpcCallResponseBody) SetSuccess(v string) *CustomeRpcCallResponseBody {
	s.Success = &v
	return s
}

type CustomeRpcCallResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CustomeRpcCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CustomeRpcCallResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomeRpcCallResponse) GoString() string {
	return s.String()
}

func (s *CustomeRpcCallResponse) SetHeaders(v map[string]*string) *CustomeRpcCallResponse {
	s.Headers = v
	return s
}

func (s *CustomeRpcCallResponse) SetStatusCode(v int32) *CustomeRpcCallResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomeRpcCallResponse) SetBody(v *CustomeRpcCallResponseBody) *CustomeRpcCallResponse {
	s.Body = v
	return s
}

type ProjectSetupResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ProjectSetupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ProjectSetupResponseBody) GoString() string {
	return s.String()
}

func (s *ProjectSetupResponseBody) SetSuccess(v bool) *ProjectSetupResponseBody {
	s.Success = &v
	return s
}

type ProjectSetupResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProjectSetupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ProjectSetupResponse) String() string {
	return tea.Prettify(s)
}

func (s ProjectSetupResponse) GoString() string {
	return s.String()
}

func (s *ProjectSetupResponse) SetHeaders(v map[string]*string) *ProjectSetupResponse {
	s.Headers = v
	return s
}

func (s *ProjectSetupResponse) SetStatusCode(v int32) *ProjectSetupResponse {
	s.StatusCode = &v
	return s
}

func (s *ProjectSetupResponse) SetBody(v *ProjectSetupResponseBody) *ProjectSetupResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

// Summary:
//
// 大客户ltcPRC接口调用
//
// @param tmpReq - CustomeRpcCallRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CustomeRpcCallResponse
func (client *Client) CustomeRpcCallWithOptions(tmpReq *CustomeRpcCallRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CustomeRpcCallResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CustomeRpcCallShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Params)) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, tea.String("params"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MethodName)) {
		query["methodName"] = request.MethodName
	}

	if !tea.BoolValue(util.IsUnset(request.ParamsShrink)) {
		query["params"] = request.ParamsShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CustomeRpcCall"),
		Version:     tea.String("customer_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/customer/rpcCall"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CustomeRpcCallResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 大客户ltcPRC接口调用
//
// @param request - CustomeRpcCallRequest
//
// @return CustomeRpcCallResponse
func (client *Client) CustomeRpcCall(request *CustomeRpcCallRequest) (_result *CustomeRpcCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CustomeRpcCallResponse{}
	_body, _err := client.CustomeRpcCallWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 直签立项审批
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ProjectSetupResponse
func (client *Client) ProjectSetupWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ProjectSetupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ProjectSetup"),
		Version:     tea.String("customer_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/customer/project/setup"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ProjectSetupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 直签立项审批
//
// @return ProjectSetupResponse
func (client *Client) ProjectSetup() (_result *ProjectSetupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ProjectSetupResponse{}
	_body, _err := client.ProjectSetupWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
