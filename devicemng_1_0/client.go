// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package devicemng_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type RegisterDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RegisterDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceHeaders) GoString() string {
	return s.String()
}

func (s *RegisterDeviceHeaders) SetCommonHeaders(v map[string]*string) *RegisterDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RegisterDeviceHeaders) SetXAcsDingtalkAccessToken(v string) *RegisterDeviceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RegisterDeviceRequest struct {
	// 组织id
	DingCorpId *string `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	// 设备标识
	DeviceKey *string `json:"deviceKey,omitempty" xml:"deviceKey,omitempty"`
	// 设备名称
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// 部门id
	DepartmentId *int64 `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 管理员userId列表
	Managers *string `json:"managers,omitempty" xml:"managers,omitempty"`
	// 协助者userId列表
	Collaborators *string `json:"collaborators,omitempty" xml:"collaborators,omitempty"`
	// 设备描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 创建者userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s RegisterDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceRequest) GoString() string {
	return s.String()
}

func (s *RegisterDeviceRequest) SetDingCorpId(v string) *RegisterDeviceRequest {
	s.DingCorpId = &v
	return s
}

func (s *RegisterDeviceRequest) SetDeviceKey(v string) *RegisterDeviceRequest {
	s.DeviceKey = &v
	return s
}

func (s *RegisterDeviceRequest) SetDeviceName(v string) *RegisterDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *RegisterDeviceRequest) SetDepartmentId(v int64) *RegisterDeviceRequest {
	s.DepartmentId = &v
	return s
}

func (s *RegisterDeviceRequest) SetManagers(v string) *RegisterDeviceRequest {
	s.Managers = &v
	return s
}

func (s *RegisterDeviceRequest) SetCollaborators(v string) *RegisterDeviceRequest {
	s.Collaborators = &v
	return s
}

func (s *RegisterDeviceRequest) SetDescription(v string) *RegisterDeviceRequest {
	s.Description = &v
	return s
}

func (s *RegisterDeviceRequest) SetUserId(v string) *RegisterDeviceRequest {
	s.UserId = &v
	return s
}

type RegisterDeviceResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RegisterDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponseBody) SetResult(v string) *RegisterDeviceResponseBody {
	s.Result = &v
	return s
}

type RegisterDeviceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceResponse) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponse) SetHeaders(v map[string]*string) *RegisterDeviceResponse {
	s.Headers = v
	return s
}

func (s *RegisterDeviceResponse) SetBody(v *RegisterDeviceResponseBody) *RegisterDeviceResponse {
	s.Body = v
	return s
}

type BatchRegisterDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchRegisterDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchRegisterDeviceHeaders) GoString() string {
	return s.String()
}

func (s *BatchRegisterDeviceHeaders) SetCommonHeaders(v map[string]*string) *BatchRegisterDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchRegisterDeviceHeaders) SetXAcsDingtalkAccessToken(v string) *BatchRegisterDeviceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchRegisterDeviceRequest struct {
	// 设备列表
	DeviceList []*BatchRegisterDeviceRequestDeviceList `json:"deviceList,omitempty" xml:"deviceList,omitempty" type:"Repeated"`
	// 组织id
	DingCorpId *string `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	// 创建者userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchRegisterDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchRegisterDeviceRequest) GoString() string {
	return s.String()
}

func (s *BatchRegisterDeviceRequest) SetDeviceList(v []*BatchRegisterDeviceRequestDeviceList) *BatchRegisterDeviceRequest {
	s.DeviceList = v
	return s
}

func (s *BatchRegisterDeviceRequest) SetDingCorpId(v string) *BatchRegisterDeviceRequest {
	s.DingCorpId = &v
	return s
}

func (s *BatchRegisterDeviceRequest) SetUserId(v string) *BatchRegisterDeviceRequest {
	s.UserId = &v
	return s
}

type BatchRegisterDeviceRequestDeviceList struct {
	// 设备标识
	DeviceKey *string `json:"deviceKey,omitempty" xml:"deviceKey,omitempty"`
	// 设备名称
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// 部门id
	DepartmentId *int64 `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 管理员userId列表
	Managers *string `json:"managers,omitempty" xml:"managers,omitempty"`
	// 协助者userId列表
	Collaborators *string `json:"collaborators,omitempty" xml:"collaborators,omitempty"`
	// 设备描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s BatchRegisterDeviceRequestDeviceList) String() string {
	return tea.Prettify(s)
}

func (s BatchRegisterDeviceRequestDeviceList) GoString() string {
	return s.String()
}

func (s *BatchRegisterDeviceRequestDeviceList) SetDeviceKey(v string) *BatchRegisterDeviceRequestDeviceList {
	s.DeviceKey = &v
	return s
}

func (s *BatchRegisterDeviceRequestDeviceList) SetDeviceName(v string) *BatchRegisterDeviceRequestDeviceList {
	s.DeviceName = &v
	return s
}

func (s *BatchRegisterDeviceRequestDeviceList) SetDepartmentId(v int64) *BatchRegisterDeviceRequestDeviceList {
	s.DepartmentId = &v
	return s
}

func (s *BatchRegisterDeviceRequestDeviceList) SetManagers(v string) *BatchRegisterDeviceRequestDeviceList {
	s.Managers = &v
	return s
}

func (s *BatchRegisterDeviceRequestDeviceList) SetCollaborators(v string) *BatchRegisterDeviceRequestDeviceList {
	s.Collaborators = &v
	return s
}

func (s *BatchRegisterDeviceRequestDeviceList) SetDescription(v string) *BatchRegisterDeviceRequestDeviceList {
	s.Description = &v
	return s
}

type BatchRegisterDeviceResponseBody struct {
	// Id of the request
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s BatchRegisterDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchRegisterDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *BatchRegisterDeviceResponseBody) SetResult(v string) *BatchRegisterDeviceResponseBody {
	s.Result = &v
	return s
}

type BatchRegisterDeviceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchRegisterDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchRegisterDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchRegisterDeviceResponse) GoString() string {
	return s.String()
}

func (s *BatchRegisterDeviceResponse) SetHeaders(v map[string]*string) *BatchRegisterDeviceResponse {
	s.Headers = v
	return s
}

func (s *BatchRegisterDeviceResponse) SetBody(v *BatchRegisterDeviceResponseBody) *BatchRegisterDeviceResponse {
	s.Body = v
	return s
}

type CreateDepartmentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateDepartmentHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateDepartmentHeaders) GoString() string {
	return s.String()
}

func (s *CreateDepartmentHeaders) SetCommonHeaders(v map[string]*string) *CreateDepartmentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateDepartmentHeaders) SetXAcsDingtalkAccessToken(v string) *CreateDepartmentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateDepartmentRequest struct {
	// 组织id
	DingCorpId *string `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	// 部门名称
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 部门类型
	DepartmentType *string `json:"departmentType,omitempty" xml:"departmentType,omitempty"`
	// 业务系统地址
	SystemUrl *string `json:"systemUrl,omitempty" xml:"systemUrl,omitempty"`
	// 认证方式
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
	// 认证信息
	AuthInfo *string `json:"authInfo,omitempty" xml:"authInfo,omitempty"`
	// 部门描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 业务扩展
	BizExt *string `json:"bizExt,omitempty" xml:"bizExt,omitempty"`
	// 创建人工号
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDepartmentRequest) GoString() string {
	return s.String()
}

func (s *CreateDepartmentRequest) SetDingCorpId(v string) *CreateDepartmentRequest {
	s.DingCorpId = &v
	return s
}

func (s *CreateDepartmentRequest) SetDepartmentName(v string) *CreateDepartmentRequest {
	s.DepartmentName = &v
	return s
}

func (s *CreateDepartmentRequest) SetDepartmentType(v string) *CreateDepartmentRequest {
	s.DepartmentType = &v
	return s
}

func (s *CreateDepartmentRequest) SetSystemUrl(v string) *CreateDepartmentRequest {
	s.SystemUrl = &v
	return s
}

func (s *CreateDepartmentRequest) SetAuthType(v string) *CreateDepartmentRequest {
	s.AuthType = &v
	return s
}

func (s *CreateDepartmentRequest) SetAuthInfo(v string) *CreateDepartmentRequest {
	s.AuthInfo = &v
	return s
}

func (s *CreateDepartmentRequest) SetDescription(v string) *CreateDepartmentRequest {
	s.Description = &v
	return s
}

func (s *CreateDepartmentRequest) SetBizExt(v string) *CreateDepartmentRequest {
	s.BizExt = &v
	return s
}

func (s *CreateDepartmentRequest) SetUserId(v string) *CreateDepartmentRequest {
	s.UserId = &v
	return s
}

type CreateDepartmentResponseBody struct {
	// Id of the request
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDepartmentResponseBody) SetResult(v string) *CreateDepartmentResponseBody {
	s.Result = &v
	return s
}

type CreateDepartmentResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDepartmentResponse) GoString() string {
	return s.String()
}

func (s *CreateDepartmentResponse) SetHeaders(v map[string]*string) *CreateDepartmentResponse {
	s.Headers = v
	return s
}

func (s *CreateDepartmentResponse) SetBody(v *CreateDepartmentResponseBody) *CreateDepartmentResponse {
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

func (client *Client) RegisterDevice(request *RegisterDeviceRequest) (_result *RegisterDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RegisterDeviceHeaders{}
	_result = &RegisterDeviceResponse{}
	_body, _err := client.RegisterDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterDeviceWithOptions(request *RegisterDeviceRequest, headers *RegisterDeviceHeaders, runtime *util.RuntimeOptions) (_result *RegisterDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingCorpId)) {
		body["dingCorpId"] = request.DingCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceKey)) {
		body["deviceKey"] = request.DeviceKey
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		body["deviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		body["departmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.Managers)) {
		body["managers"] = request.Managers
	}

	if !tea.BoolValue(util.IsUnset(request.Collaborators)) {
		body["collaborators"] = request.Collaborators
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &RegisterDeviceResponse{}
	_body, _err := client.DoROARequest(tea.String("RegisterDevice"), tea.String("devicemng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/devicemng/devices"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchRegisterDevice(request *BatchRegisterDeviceRequest) (_result *BatchRegisterDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchRegisterDeviceHeaders{}
	_result = &BatchRegisterDeviceResponse{}
	_body, _err := client.BatchRegisterDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchRegisterDeviceWithOptions(request *BatchRegisterDeviceRequest, headers *BatchRegisterDeviceHeaders, runtime *util.RuntimeOptions) (_result *BatchRegisterDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceList)) {
		body["deviceList"] = request.DeviceList
	}

	if !tea.BoolValue(util.IsUnset(request.DingCorpId)) {
		body["dingCorpId"] = request.DingCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &BatchRegisterDeviceResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchRegisterDevice"), tea.String("devicemng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/devicemng/devices/batch"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDepartment(request *CreateDepartmentRequest) (_result *CreateDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateDepartmentHeaders{}
	_result = &CreateDepartmentResponse{}
	_body, _err := client.CreateDepartmentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDepartmentWithOptions(request *CreateDepartmentRequest, headers *CreateDepartmentHeaders, runtime *util.RuntimeOptions) (_result *CreateDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingCorpId)) {
		body["dingCorpId"] = request.DingCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentName)) {
		body["departmentName"] = request.DepartmentName
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentType)) {
		body["departmentType"] = request.DepartmentType
	}

	if !tea.BoolValue(util.IsUnset(request.SystemUrl)) {
		body["systemUrl"] = request.SystemUrl
	}

	if !tea.BoolValue(util.IsUnset(request.AuthType)) {
		body["authType"] = request.AuthType
	}

	if !tea.BoolValue(util.IsUnset(request.AuthInfo)) {
		body["authInfo"] = request.AuthInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.BizExt)) {
		body["bizExt"] = request.BizExt
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateDepartmentResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateDepartment"), tea.String("devicemng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/devicemng/departments"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
