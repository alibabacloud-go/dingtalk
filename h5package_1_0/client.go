// This file is auto-generated, don't edit it. Thanks.
package h5package_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreatePackageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreatePackageHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreatePackageHeaders) GoString() string {
	return s.String()
}

func (s *CreatePackageHeaders) SetCommonHeaders(v map[string]*string) *CreatePackageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreatePackageHeaders) SetXAcsDingtalkAccessToken(v string) *CreatePackageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreatePackageRequest struct {
	// example:
	//
	// 1234
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// 1234
	AppId *int64 `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// https://example.com/myapp/index.html
	HomeUrl *string `json:"homeUrl,omitempty" xml:"homeUrl,omitempty"`
	// example:
	//
	// aaaaaa/bbbbbb
	OssObjectKey *string `json:"ossObjectKey,omitempty" xml:"ossObjectKey,omitempty"`
}

func (s CreatePackageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePackageRequest) GoString() string {
	return s.String()
}

func (s *CreatePackageRequest) SetAgentId(v int64) *CreatePackageRequest {
	s.AgentId = &v
	return s
}

func (s *CreatePackageRequest) SetAppId(v int64) *CreatePackageRequest {
	s.AppId = &v
	return s
}

func (s *CreatePackageRequest) SetHomeUrl(v string) *CreatePackageRequest {
	s.HomeUrl = &v
	return s
}

func (s *CreatePackageRequest) SetOssObjectKey(v string) *CreatePackageRequest {
	s.OssObjectKey = &v
	return s
}

type CreatePackageResponseBody struct {
	// example:
	//
	// 1663748308644pjpF
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreatePackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePackageResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePackageResponseBody) SetTaskId(v string) *CreatePackageResponseBody {
	s.TaskId = &v
	return s
}

type CreatePackageResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePackageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePackageResponse) GoString() string {
	return s.String()
}

func (s *CreatePackageResponse) SetHeaders(v map[string]*string) *CreatePackageResponse {
	s.Headers = v
	return s
}

func (s *CreatePackageResponse) SetStatusCode(v int32) *CreatePackageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePackageResponse) SetBody(v *CreatePackageResponseBody) *CreatePackageResponse {
	s.Body = v
	return s
}

type GetAccessTokenHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAccessTokenHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAccessTokenHeaders) GoString() string {
	return s.String()
}

func (s *GetAccessTokenHeaders) SetCommonHeaders(v map[string]*string) *GetAccessTokenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAccessTokenHeaders) SetXAcsDingtalkAccessToken(v string) *GetAccessTokenHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAccessTokenRequest struct {
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	AppId   *int64 `json:"appId,omitempty" xml:"appId,omitempty"`
}

func (s GetAccessTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *GetAccessTokenRequest) SetAgentId(v int64) *GetAccessTokenRequest {
	s.AgentId = &v
	return s
}

func (s *GetAccessTokenRequest) SetAppId(v int64) *GetAccessTokenRequest {
	s.AppId = &v
	return s
}

type GetAccessTokenResponseBody struct {
	// example:
	//
	// STS.NUPjgnMhCVWvo1HSxfftf
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// example:
	//
	// ASviryNDy9tTuS5KiYMA6fCYf81vHg4KdoX7CVHz4CSx
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	// example:
	//
	// dingtalk-bucket
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// example:
	//
	// oss-cn-shanghai.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// example:
	//
	// 2022-09-21T09:32:16Z
	Expiration *string `json:"expiration,omitempty" xml:"expiration,omitempty"`
	// example:
	//
	// 5000000002761167/1663751835956
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// oss-cn-shanghai
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// CAIS0QJ1q6Ft5B2yfSjIr5blId3aoLdi4ZWdbRf5t3gzavt...
	StsToken *string `json:"stsToken,omitempty" xml:"stsToken,omitempty"`
}

func (s GetAccessTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessTokenResponseBody) SetAccessKeyId(v string) *GetAccessTokenResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *GetAccessTokenResponseBody) SetAccessKeySecret(v string) *GetAccessTokenResponseBody {
	s.AccessKeySecret = &v
	return s
}

func (s *GetAccessTokenResponseBody) SetBucket(v string) *GetAccessTokenResponseBody {
	s.Bucket = &v
	return s
}

func (s *GetAccessTokenResponseBody) SetEndpoint(v string) *GetAccessTokenResponseBody {
	s.Endpoint = &v
	return s
}

func (s *GetAccessTokenResponseBody) SetExpiration(v string) *GetAccessTokenResponseBody {
	s.Expiration = &v
	return s
}

func (s *GetAccessTokenResponseBody) SetName(v string) *GetAccessTokenResponseBody {
	s.Name = &v
	return s
}

func (s *GetAccessTokenResponseBody) SetRegion(v string) *GetAccessTokenResponseBody {
	s.Region = &v
	return s
}

func (s *GetAccessTokenResponseBody) SetStsToken(v string) *GetAccessTokenResponseBody {
	s.StsToken = &v
	return s
}

type GetAccessTokenResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *GetAccessTokenResponse) SetHeaders(v map[string]*string) *GetAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *GetAccessTokenResponse) SetStatusCode(v int32) *GetAccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessTokenResponse) SetBody(v *GetAccessTokenResponseBody) *GetAccessTokenResponse {
	s.Body = v
	return s
}

type GetCreateStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCreateStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCreateStatusHeaders) GoString() string {
	return s.String()
}

func (s *GetCreateStatusHeaders) SetCommonHeaders(v map[string]*string) *GetCreateStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCreateStatusHeaders) SetXAcsDingtalkAccessToken(v string) *GetCreateStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCreateStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1663748308644pjpF
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetCreateStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCreateStatusRequest) GoString() string {
	return s.String()
}

func (s *GetCreateStatusRequest) SetTaskId(v string) *GetCreateStatusRequest {
	s.TaskId = &v
	return s
}

type GetCreateStatusResponseBody struct {
	// example:
	//
	// 1663743241146
	BuildTime *int64 `json:"buildTime,omitempty" xml:"buildTime,omitempty"`
	Finished  *bool  `json:"finished,omitempty" xml:"finished,omitempty"`
	// example:
	//
	// 0
	PackageSize *int64 `json:"packageSize,omitempty" xml:"packageSize,omitempty"`
	// example:
	//
	// 2
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1663748308644pjpF
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 0.0.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetCreateStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCreateStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetCreateStatusResponseBody) SetBuildTime(v int64) *GetCreateStatusResponseBody {
	s.BuildTime = &v
	return s
}

func (s *GetCreateStatusResponseBody) SetFinished(v bool) *GetCreateStatusResponseBody {
	s.Finished = &v
	return s
}

func (s *GetCreateStatusResponseBody) SetPackageSize(v int64) *GetCreateStatusResponseBody {
	s.PackageSize = &v
	return s
}

func (s *GetCreateStatusResponseBody) SetStatus(v string) *GetCreateStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetCreateStatusResponseBody) SetTaskId(v string) *GetCreateStatusResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetCreateStatusResponseBody) SetVersion(v string) *GetCreateStatusResponseBody {
	s.Version = &v
	return s
}

type GetCreateStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCreateStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCreateStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCreateStatusResponse) GoString() string {
	return s.String()
}

func (s *GetCreateStatusResponse) SetHeaders(v map[string]*string) *GetCreateStatusResponse {
	s.Headers = v
	return s
}

func (s *GetCreateStatusResponse) SetStatusCode(v int32) *GetCreateStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCreateStatusResponse) SetBody(v *GetCreateStatusResponseBody) *GetCreateStatusResponse {
	s.Body = v
	return s
}

type PublishPackageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PublishPackageHeaders) String() string {
	return tea.Prettify(s)
}

func (s PublishPackageHeaders) GoString() string {
	return s.String()
}

func (s *PublishPackageHeaders) SetCommonHeaders(v map[string]*string) *PublishPackageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PublishPackageHeaders) SetXAcsDingtalkAccessToken(v string) *PublishPackageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PublishPackageRequest struct {
	// example:
	//
	// 1234
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// 1234
	AppId *int64 `json:"appId,omitempty" xml:"appId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.0.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s PublishPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishPackageRequest) GoString() string {
	return s.String()
}

func (s *PublishPackageRequest) SetAgentId(v int64) *PublishPackageRequest {
	s.AgentId = &v
	return s
}

func (s *PublishPackageRequest) SetAppId(v int64) *PublishPackageRequest {
	s.AppId = &v
	return s
}

func (s *PublishPackageRequest) SetVersion(v string) *PublishPackageRequest {
	s.Version = &v
	return s
}

type PublishPackageResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PublishPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishPackageResponseBody) GoString() string {
	return s.String()
}

func (s *PublishPackageResponseBody) SetSuccess(v bool) *PublishPackageResponseBody {
	s.Success = &v
	return s
}

type PublishPackageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishPackageResponse) GoString() string {
	return s.String()
}

func (s *PublishPackageResponse) SetHeaders(v map[string]*string) *PublishPackageResponse {
	s.Headers = v
	return s
}

func (s *PublishPackageResponse) SetStatusCode(v int32) *PublishPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishPackageResponse) SetBody(v *PublishPackageResponseBody) *PublishPackageResponse {
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
// 上传H5离线包
//
// @param request - CreatePackageRequest
//
// @param headers - CreatePackageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePackageResponse
func (client *Client) CreatePackageWithOptions(request *CreatePackageRequest, headers *CreatePackageHeaders, runtime *util.RuntimeOptions) (_result *CreatePackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		body["agentId"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.HomeUrl)) {
		body["homeUrl"] = request.HomeUrl
	}

	if !tea.BoolValue(util.IsUnset(request.OssObjectKey)) {
		body["ossObjectKey"] = request.OssObjectKey
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
		Action:      tea.String("CreatePackage"),
		Version:     tea.String("h5package_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/h5package/asyncUpload"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePackageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传H5离线包
//
// @param request - CreatePackageRequest
//
// @return CreatePackageResponse
func (client *Client) CreatePackage(request *CreatePackageRequest) (_result *CreatePackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreatePackageHeaders{}
	_result = &CreatePackageResponse{}
	_body, _err := client.CreatePackageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取包上传一次性AccessToken
//
// @param request - GetAccessTokenRequest
//
// @param headers - GetAccessTokenHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessTokenResponse
func (client *Client) GetAccessTokenWithOptions(request *GetAccessTokenRequest, headers *GetAccessTokenHeaders, runtime *util.RuntimeOptions) (_result *GetAccessTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		query["agentId"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["appId"] = request.AppId
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
		Action:      tea.String("GetAccessToken"),
		Version:     tea.String("h5package_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/h5package/uploadTokens"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccessTokenResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取包上传一次性AccessToken
//
// @param request - GetAccessTokenRequest
//
// @return GetAccessTokenResponse
func (client *Client) GetAccessToken(request *GetAccessTokenRequest) (_result *GetAccessTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAccessTokenHeaders{}
	_result = &GetAccessTokenResponse{}
	_body, _err := client.GetAccessTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取H5离线包版本创建状态
//
// @param request - GetCreateStatusRequest
//
// @param headers - GetCreateStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCreateStatusResponse
func (client *Client) GetCreateStatusWithOptions(request *GetCreateStatusRequest, headers *GetCreateStatusHeaders, runtime *util.RuntimeOptions) (_result *GetCreateStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
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
		Action:      tea.String("GetCreateStatus"),
		Version:     tea.String("h5package_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/h5package/uploadStatus"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCreateStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取H5离线包版本创建状态
//
// @param request - GetCreateStatusRequest
//
// @return GetCreateStatusResponse
func (client *Client) GetCreateStatus(request *GetCreateStatusRequest) (_result *GetCreateStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCreateStatusHeaders{}
	_result = &GetCreateStatusResponse{}
	_body, _err := client.GetCreateStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发布离线包
//
// @param request - PublishPackageRequest
//
// @param headers - PublishPackageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishPackageResponse
func (client *Client) PublishPackageWithOptions(request *PublishPackageRequest, headers *PublishPackageHeaders, runtime *util.RuntimeOptions) (_result *PublishPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		body["agentId"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["version"] = request.Version
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
		Action:      tea.String("PublishPackage"),
		Version:     tea.String("h5package_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/h5package/publish"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishPackageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发布离线包
//
// @param request - PublishPackageRequest
//
// @return PublishPackageResponse
func (client *Client) PublishPackage(request *PublishPackageRequest) (_result *PublishPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PublishPackageHeaders{}
	_result = &PublishPackageResponse{}
	_body, _err := client.PublishPackageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
