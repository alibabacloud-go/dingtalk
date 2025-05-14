// This file is auto-generated, don't edit it. Thanks.
package ding_phone_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddCallConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddCallConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddCallConfigHeaders) GoString() string {
	return s.String()
}

func (s *AddCallConfigHeaders) SetCommonHeaders(v map[string]*string) *AddCallConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddCallConfigHeaders) SetXAcsDingtalkAccessToken(v string) *AddCallConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddCallConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ding3f583b067f2q450c12d
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// token12345
	IsvToken *string `json:"isvToken,omitempty" xml:"isvToken,omitempty"`
	// example:
	//
	// 057112345678
	PhoneNumber *string `json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	// example:
	//
	// call
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
}

func (s AddCallConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCallConfigRequest) GoString() string {
	return s.String()
}

func (s *AddCallConfigRequest) SetCorpId(v string) *AddCallConfigRequest {
	s.CorpId = &v
	return s
}

func (s *AddCallConfigRequest) SetIsvToken(v string) *AddCallConfigRequest {
	s.IsvToken = &v
	return s
}

func (s *AddCallConfigRequest) SetPhoneNumber(v string) *AddCallConfigRequest {
	s.PhoneNumber = &v
	return s
}

func (s *AddCallConfigRequest) SetScopeType(v string) *AddCallConfigRequest {
	s.ScopeType = &v
	return s
}

type AddCallConfigResponseBody struct {
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s AddCallConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCallConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddCallConfigResponseBody) SetToken(v string) *AddCallConfigResponseBody {
	s.Token = &v
	return s
}

type AddCallConfigResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCallConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCallConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCallConfigResponse) GoString() string {
	return s.String()
}

func (s *AddCallConfigResponse) SetHeaders(v map[string]*string) *AddCallConfigResponse {
	s.Headers = v
	return s
}

func (s *AddCallConfigResponse) SetStatusCode(v int32) *AddCallConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCallConfigResponse) SetBody(v *AddCallConfigResponseBody) *AddCallConfigResponse {
	s.Body = v
	return s
}

type DelCallConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DelCallConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s DelCallConfigHeaders) GoString() string {
	return s.String()
}

func (s *DelCallConfigHeaders) SetCommonHeaders(v map[string]*string) *DelCallConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DelCallConfigHeaders) SetXAcsDingtalkAccessToken(v string) *DelCallConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DelCallConfigRequest struct {
	// example:
	//
	// ding3f583b067250d34dd
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// token1233143
	IsvToken *string `json:"isvToken,omitempty" xml:"isvToken,omitempty"`
	// example:
	//
	// 057112345678
	PhoneNumber *string `json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
}

func (s DelCallConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DelCallConfigRequest) GoString() string {
	return s.String()
}

func (s *DelCallConfigRequest) SetCorpId(v string) *DelCallConfigRequest {
	s.CorpId = &v
	return s
}

func (s *DelCallConfigRequest) SetIsvToken(v string) *DelCallConfigRequest {
	s.IsvToken = &v
	return s
}

func (s *DelCallConfigRequest) SetPhoneNumber(v string) *DelCallConfigRequest {
	s.PhoneNumber = &v
	return s
}

type DelCallConfigResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DelCallConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DelCallConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DelCallConfigResponseBody) SetResult(v bool) *DelCallConfigResponseBody {
	s.Result = &v
	return s
}

type DelCallConfigResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DelCallConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DelCallConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DelCallConfigResponse) GoString() string {
	return s.String()
}

func (s *DelCallConfigResponse) SetHeaders(v map[string]*string) *DelCallConfigResponse {
	s.Headers = v
	return s
}

func (s *DelCallConfigResponse) SetStatusCode(v int32) *DelCallConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DelCallConfigResponse) SetBody(v *DelCallConfigResponseBody) *DelCallConfigResponse {
	s.Body = v
	return s
}

type QueryCallConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCallConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCallConfigHeaders) GoString() string {
	return s.String()
}

func (s *QueryCallConfigHeaders) SetCommonHeaders(v map[string]*string) *QueryCallConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCallConfigHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCallConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCallConfigRequest struct {
	// example:
	//
	// ding3f583b0672efc12d
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// token23dafds
	IsvToken *string `json:"isvToken,omitempty" xml:"isvToken,omitempty"`
	// example:
	//
	// 057112345678
	PhoneNumber *string `json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	// example:
	//
	// call
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
}

func (s QueryCallConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCallConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryCallConfigRequest) SetCorpId(v string) *QueryCallConfigRequest {
	s.CorpId = &v
	return s
}

func (s *QueryCallConfigRequest) SetIsvToken(v string) *QueryCallConfigRequest {
	s.IsvToken = &v
	return s
}

func (s *QueryCallConfigRequest) SetPhoneNumber(v string) *QueryCallConfigRequest {
	s.PhoneNumber = &v
	return s
}

func (s *QueryCallConfigRequest) SetScopeType(v string) *QueryCallConfigRequest {
	s.ScopeType = &v
	return s
}

type QueryCallConfigResponseBody struct {
	Result []*QueryCallConfigResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryCallConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCallConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCallConfigResponseBody) SetResult(v []*QueryCallConfigResponseBodyResult) *QueryCallConfigResponseBody {
	s.Result = v
	return s
}

type QueryCallConfigResponseBodyResult struct {
	AccountDomain *string `json:"accountDomain,omitempty" xml:"accountDomain,omitempty"`
	AccountId     *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	CallInType    *int32  `json:"callInType,omitempty" xml:"callInType,omitempty"`
	CallOutType   *int32  `json:"callOutType,omitempty" xml:"callOutType,omitempty"`
	CreateUid     *string `json:"createUid,omitempty" xml:"createUid,omitempty"`
	PhoneNumber   *string `json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	ScopeType     *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	ShowType      *int32  `json:"showType,omitempty" xml:"showType,omitempty"`
	SourceType    *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Status        *int32  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryCallConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryCallConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryCallConfigResponseBodyResult) SetAccountDomain(v string) *QueryCallConfigResponseBodyResult {
	s.AccountDomain = &v
	return s
}

func (s *QueryCallConfigResponseBodyResult) SetAccountId(v string) *QueryCallConfigResponseBodyResult {
	s.AccountId = &v
	return s
}

func (s *QueryCallConfigResponseBodyResult) SetCallInType(v int32) *QueryCallConfigResponseBodyResult {
	s.CallInType = &v
	return s
}

func (s *QueryCallConfigResponseBodyResult) SetCallOutType(v int32) *QueryCallConfigResponseBodyResult {
	s.CallOutType = &v
	return s
}

func (s *QueryCallConfigResponseBodyResult) SetCreateUid(v string) *QueryCallConfigResponseBodyResult {
	s.CreateUid = &v
	return s
}

func (s *QueryCallConfigResponseBodyResult) SetPhoneNumber(v string) *QueryCallConfigResponseBodyResult {
	s.PhoneNumber = &v
	return s
}

func (s *QueryCallConfigResponseBodyResult) SetScopeType(v string) *QueryCallConfigResponseBodyResult {
	s.ScopeType = &v
	return s
}

func (s *QueryCallConfigResponseBodyResult) SetShowType(v int32) *QueryCallConfigResponseBodyResult {
	s.ShowType = &v
	return s
}

func (s *QueryCallConfigResponseBodyResult) SetSourceType(v string) *QueryCallConfigResponseBodyResult {
	s.SourceType = &v
	return s
}

func (s *QueryCallConfigResponseBodyResult) SetStatus(v int32) *QueryCallConfigResponseBodyResult {
	s.Status = &v
	return s
}

type QueryCallConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCallConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCallConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCallConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryCallConfigResponse) SetHeaders(v map[string]*string) *QueryCallConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryCallConfigResponse) SetStatusCode(v int32) *QueryCallConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCallConfigResponse) SetBody(v *QueryCallConfigResponseBody) *QueryCallConfigResponse {
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
// 添加外呼码号配置
//
// @param request - AddCallConfigRequest
//
// @param headers - AddCallConfigHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCallConfigResponse
func (client *Client) AddCallConfigWithOptions(request *AddCallConfigRequest, headers *AddCallConfigHeaders, runtime *util.RuntimeOptions) (_result *AddCallConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.IsvToken)) {
		query["isvToken"] = request.IsvToken
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["phoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeType)) {
		query["scopeType"] = request.ScopeType
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
		Action:      tea.String("AddCallConfig"),
		Version:     tea.String("dingPhone_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/dingPhone/callConfigs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddCallConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加外呼码号配置
//
// @param request - AddCallConfigRequest
//
// @return AddCallConfigResponse
func (client *Client) AddCallConfig(request *AddCallConfigRequest) (_result *AddCallConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddCallConfigHeaders{}
	_result = &AddCallConfigResponse{}
	_body, _err := client.AddCallConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除码号配置
//
// @param request - DelCallConfigRequest
//
// @param headers - DelCallConfigHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DelCallConfigResponse
func (client *Client) DelCallConfigWithOptions(request *DelCallConfigRequest, headers *DelCallConfigHeaders, runtime *util.RuntimeOptions) (_result *DelCallConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.IsvToken)) {
		query["isvToken"] = request.IsvToken
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["phoneNumber"] = request.PhoneNumber
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
		Action:      tea.String("DelCallConfig"),
		Version:     tea.String("dingPhone_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/dingPhone/callConfigs"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DelCallConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除码号配置
//
// @param request - DelCallConfigRequest
//
// @return DelCallConfigResponse
func (client *Client) DelCallConfig(request *DelCallConfigRequest) (_result *DelCallConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DelCallConfigHeaders{}
	_result = &DelCallConfigResponse{}
	_body, _err := client.DelCallConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询外呼码号配置
//
// @param request - QueryCallConfigRequest
//
// @param headers - QueryCallConfigHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCallConfigResponse
func (client *Client) QueryCallConfigWithOptions(request *QueryCallConfigRequest, headers *QueryCallConfigHeaders, runtime *util.RuntimeOptions) (_result *QueryCallConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.IsvToken)) {
		query["isvToken"] = request.IsvToken
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["phoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeType)) {
		query["scopeType"] = request.ScopeType
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
		Action:      tea.String("QueryCallConfig"),
		Version:     tea.String("dingPhone_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/dingPhone/callConfigs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCallConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询外呼码号配置
//
// @param request - QueryCallConfigRequest
//
// @return QueryCallConfigResponse
func (client *Client) QueryCallConfig(request *QueryCallConfigRequest) (_result *QueryCallConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCallConfigHeaders{}
	_result = &QueryCallConfigResponse{}
	_body, _err := client.QueryCallConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
