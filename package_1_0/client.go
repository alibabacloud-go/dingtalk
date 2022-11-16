// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package package_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetUploadTokenHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUploadTokenHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUploadTokenHeaders) GoString() string {
	return s.String()
}

func (s *GetUploadTokenHeaders) SetCommonHeaders(v map[string]*string) *GetUploadTokenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUploadTokenHeaders) SetXAcsDingtalkAccessToken(v string) *GetUploadTokenHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUploadTokenRequest struct {
	// 离线包ID
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
}

func (s GetUploadTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUploadTokenRequest) GoString() string {
	return s.String()
}

func (s *GetUploadTokenRequest) SetMiniAppId(v string) *GetUploadTokenRequest {
	s.MiniAppId = &v
	return s
}

type GetUploadTokenResponseBody struct {
	// 阿里云OSS SDK初始化配置项
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// 阿里云OSS SDK初始化配置项
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	// 阿里云OSS SDK初始化配置项
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// 阿里云OSS SDK初始化配置项
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// 阿里云OSS SDK初始化配置项
	Expiration *string `json:"expiration,omitempty" xml:"expiration,omitempty"`
	// 阿里云OSS SDK初始化配置项
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 阿里云OSS SDK初始化配置项
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// 阿里云OSS SDK初始化配置项
	StsToken *string `json:"stsToken,omitempty" xml:"stsToken,omitempty"`
}

func (s GetUploadTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUploadTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadTokenResponseBody) SetAccessKeyId(v string) *GetUploadTokenResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *GetUploadTokenResponseBody) SetAccessKeySecret(v string) *GetUploadTokenResponseBody {
	s.AccessKeySecret = &v
	return s
}

func (s *GetUploadTokenResponseBody) SetBucket(v string) *GetUploadTokenResponseBody {
	s.Bucket = &v
	return s
}

func (s *GetUploadTokenResponseBody) SetEndpoint(v string) *GetUploadTokenResponseBody {
	s.Endpoint = &v
	return s
}

func (s *GetUploadTokenResponseBody) SetExpiration(v string) *GetUploadTokenResponseBody {
	s.Expiration = &v
	return s
}

func (s *GetUploadTokenResponseBody) SetName(v string) *GetUploadTokenResponseBody {
	s.Name = &v
	return s
}

func (s *GetUploadTokenResponseBody) SetRegion(v string) *GetUploadTokenResponseBody {
	s.Region = &v
	return s
}

func (s *GetUploadTokenResponseBody) SetStsToken(v string) *GetUploadTokenResponseBody {
	s.StsToken = &v
	return s
}

type GetUploadTokenResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUploadTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUploadTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUploadTokenResponse) GoString() string {
	return s.String()
}

func (s *GetUploadTokenResponse) SetHeaders(v map[string]*string) *GetUploadTokenResponse {
	s.Headers = v
	return s
}

func (s *GetUploadTokenResponse) SetBody(v *GetUploadTokenResponseBody) *GetUploadTokenResponse {
	s.Body = v
	return s
}

type HPublishPackageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HPublishPackageHeaders) String() string {
	return tea.Prettify(s)
}

func (s HPublishPackageHeaders) GoString() string {
	return s.String()
}

func (s *HPublishPackageHeaders) SetCommonHeaders(v map[string]*string) *HPublishPackageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HPublishPackageHeaders) SetXAcsDingtalkAccessToken(v string) *HPublishPackageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HPublishPackageRequest struct {
	// 离线包ID
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	// 离线包版本号
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s HPublishPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s HPublishPackageRequest) GoString() string {
	return s.String()
}

func (s *HPublishPackageRequest) SetMiniAppId(v string) *HPublishPackageRequest {
	s.MiniAppId = &v
	return s
}

func (s *HPublishPackageRequest) SetVersion(v string) *HPublishPackageRequest {
	s.Version = &v
	return s
}

type HPublishPackageResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HPublishPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HPublishPackageResponseBody) GoString() string {
	return s.String()
}

func (s *HPublishPackageResponseBody) SetSuccess(v bool) *HPublishPackageResponseBody {
	s.Success = &v
	return s
}

type HPublishPackageResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *HPublishPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HPublishPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s HPublishPackageResponse) GoString() string {
	return s.String()
}

func (s *HPublishPackageResponse) SetHeaders(v map[string]*string) *HPublishPackageResponse {
	s.Headers = v
	return s
}

func (s *HPublishPackageResponse) SetBody(v *HPublishPackageResponseBody) *HPublishPackageResponse {
	s.Body = v
	return s
}

type HUploadPackageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HUploadPackageHeaders) String() string {
	return tea.Prettify(s)
}

func (s HUploadPackageHeaders) GoString() string {
	return s.String()
}

func (s *HUploadPackageHeaders) SetCommonHeaders(v map[string]*string) *HUploadPackageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HUploadPackageHeaders) SetXAcsDingtalkAccessToken(v string) *HUploadPackageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HUploadPackageRequest struct {
	// 离线包ID
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	// 离线包资源OSS Key
	OssObjectKey *string `json:"ossObjectKey,omitempty" xml:"ossObjectKey,omitempty"`
}

func (s HUploadPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s HUploadPackageRequest) GoString() string {
	return s.String()
}

func (s *HUploadPackageRequest) SetMiniAppId(v string) *HUploadPackageRequest {
	s.MiniAppId = &v
	return s
}

func (s *HUploadPackageRequest) SetOssObjectKey(v string) *HUploadPackageRequest {
	s.OssObjectKey = &v
	return s
}

type HUploadPackageResponseBody struct {
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s HUploadPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HUploadPackageResponseBody) GoString() string {
	return s.String()
}

func (s *HUploadPackageResponseBody) SetTaskId(v string) *HUploadPackageResponseBody {
	s.TaskId = &v
	return s
}

type HUploadPackageResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *HUploadPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HUploadPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s HUploadPackageResponse) GoString() string {
	return s.String()
}

func (s *HUploadPackageResponse) SetHeaders(v map[string]*string) *HUploadPackageResponse {
	s.Headers = v
	return s
}

func (s *HUploadPackageResponse) SetBody(v *HUploadPackageResponseBody) *HUploadPackageResponse {
	s.Body = v
	return s
}

type HUploadPackageStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HUploadPackageStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s HUploadPackageStatusHeaders) GoString() string {
	return s.String()
}

func (s *HUploadPackageStatusHeaders) SetCommonHeaders(v map[string]*string) *HUploadPackageStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HUploadPackageStatusHeaders) SetXAcsDingtalkAccessToken(v string) *HUploadPackageStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HUploadPackageStatusRequest struct {
	// 离线包ID
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	// 上传任务ID
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s HUploadPackageStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s HUploadPackageStatusRequest) GoString() string {
	return s.String()
}

func (s *HUploadPackageStatusRequest) SetMiniAppId(v string) *HUploadPackageStatusRequest {
	s.MiniAppId = &v
	return s
}

func (s *HUploadPackageStatusRequest) SetTaskId(v string) *HUploadPackageStatusRequest {
	s.TaskId = &v
	return s
}

type HUploadPackageStatusResponseBody struct {
	// 创建时间
	BuildTime *int64 `json:"buildTime,omitempty" xml:"buildTime,omitempty"`
	// 任务是否已结束
	Finished *bool `json:"finished,omitempty" xml:"finished,omitempty"`
	// H5离线包体积，单位Byte
	PackageSize *int64 `json:"packageSize,omitempty" xml:"packageSize,omitempty"`
	// 任务状态。1：构建中；2：成功；3：失败；5：超时。
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 创建离线包接口返回的taskId
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// H5离线包版本号
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s HUploadPackageStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HUploadPackageStatusResponseBody) GoString() string {
	return s.String()
}

func (s *HUploadPackageStatusResponseBody) SetBuildTime(v int64) *HUploadPackageStatusResponseBody {
	s.BuildTime = &v
	return s
}

func (s *HUploadPackageStatusResponseBody) SetFinished(v bool) *HUploadPackageStatusResponseBody {
	s.Finished = &v
	return s
}

func (s *HUploadPackageStatusResponseBody) SetPackageSize(v int64) *HUploadPackageStatusResponseBody {
	s.PackageSize = &v
	return s
}

func (s *HUploadPackageStatusResponseBody) SetStatus(v string) *HUploadPackageStatusResponseBody {
	s.Status = &v
	return s
}

func (s *HUploadPackageStatusResponseBody) SetTaskId(v string) *HUploadPackageStatusResponseBody {
	s.TaskId = &v
	return s
}

func (s *HUploadPackageStatusResponseBody) SetVersion(v string) *HUploadPackageStatusResponseBody {
	s.Version = &v
	return s
}

type HUploadPackageStatusResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *HUploadPackageStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HUploadPackageStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s HUploadPackageStatusResponse) GoString() string {
	return s.String()
}

func (s *HUploadPackageStatusResponse) SetHeaders(v map[string]*string) *HUploadPackageStatusResponse {
	s.Headers = v
	return s
}

func (s *HUploadPackageStatusResponse) SetBody(v *HUploadPackageStatusResponseBody) *HUploadPackageStatusResponse {
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

func (client *Client) GetUploadToken(request *GetUploadTokenRequest) (_result *GetUploadTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUploadTokenHeaders{}
	_result = &GetUploadTokenResponse{}
	_body, _err := client.GetUploadTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUploadTokenWithOptions(request *GetUploadTokenRequest, headers *GetUploadTokenHeaders, runtime *util.RuntimeOptions) (_result *GetUploadTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		query["miniAppId"] = request.MiniAppId
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
	_result = &GetUploadTokenResponse{}
	_body, _err := client.DoROARequest(tea.String("GetUploadToken"), tea.String("package_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/package/uploadTokens"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HPublishPackage(request *HPublishPackageRequest) (_result *HPublishPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HPublishPackageHeaders{}
	_result = &HPublishPackageResponse{}
	_body, _err := client.HPublishPackageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HPublishPackageWithOptions(request *HPublishPackageRequest, headers *HPublishPackageHeaders, runtime *util.RuntimeOptions) (_result *HPublishPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		body["miniAppId"] = request.MiniAppId
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
	_result = &HPublishPackageResponse{}
	_body, _err := client.DoROARequest(tea.String("HPublishPackage"), tea.String("package_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/package/h5/publish"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HUploadPackage(request *HUploadPackageRequest) (_result *HUploadPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HUploadPackageHeaders{}
	_result = &HUploadPackageResponse{}
	_body, _err := client.HUploadPackageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HUploadPackageWithOptions(request *HUploadPackageRequest, headers *HUploadPackageHeaders, runtime *util.RuntimeOptions) (_result *HUploadPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		body["miniAppId"] = request.MiniAppId
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
	_result = &HUploadPackageResponse{}
	_body, _err := client.DoROARequest(tea.String("HUploadPackage"), tea.String("package_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/package/h5/asyncUpload"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HUploadPackageStatus(request *HUploadPackageStatusRequest) (_result *HUploadPackageStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HUploadPackageStatusHeaders{}
	_result = &HUploadPackageStatusResponse{}
	_body, _err := client.HUploadPackageStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HUploadPackageStatusWithOptions(request *HUploadPackageStatusRequest, headers *HUploadPackageStatusHeaders, runtime *util.RuntimeOptions) (_result *HUploadPackageStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		query["miniAppId"] = request.MiniAppId
	}

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
	_result = &HUploadPackageStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("HUploadPackageStatus"), tea.String("package_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/package/h5/uploadStatus"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
