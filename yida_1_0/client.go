// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package yida_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetFormDataByIDHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFormDataByIDHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDHeaders) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDHeaders) SetCommonHeaders(v map[string]*string) *GetFormDataByIDHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFormDataByIDHeaders) SetXAcsDingtalkAccessToken(v string) *GetFormDataByIDHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFormDataByIDRequest struct {
	// 应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 应用秘钥。在应用数据中获取。
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 语言。可选值：zh_CN/en_US 默认：zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s GetFormDataByIDRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDRequest) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDRequest) SetAppType(v string) *GetFormDataByIDRequest {
	s.AppType = &v
	return s
}

func (s *GetFormDataByIDRequest) SetSystemToken(v string) *GetFormDataByIDRequest {
	s.SystemToken = &v
	return s
}

func (s *GetFormDataByIDRequest) SetUserId(v string) *GetFormDataByIDRequest {
	s.UserId = &v
	return s
}

func (s *GetFormDataByIDRequest) SetLanguage(v string) *GetFormDataByIDRequest {
	s.Language = &v
	return s
}

type GetFormDataByIDResponseBody struct {
	// 发起人详情
	Originator *GetFormDataByIDResponseBodyOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	// 最后修改时间
	ModifiedTimeGMT *string `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	// 表单ID
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 表单实例ID
	FormInstId *string `json:"formInstId,omitempty" xml:"formInstId,omitempty"`
	// 表单数据详情
	FormData map[string]interface{} `json:"formData,omitempty" xml:"formData,omitempty"`
}

func (s GetFormDataByIDResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDResponseBody) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDResponseBody) SetOriginator(v *GetFormDataByIDResponseBodyOriginator) *GetFormDataByIDResponseBody {
	s.Originator = v
	return s
}

func (s *GetFormDataByIDResponseBody) SetModifiedTimeGMT(v string) *GetFormDataByIDResponseBody {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *GetFormDataByIDResponseBody) SetFormUuid(v string) *GetFormDataByIDResponseBody {
	s.FormUuid = &v
	return s
}

func (s *GetFormDataByIDResponseBody) SetFormInstId(v string) *GetFormDataByIDResponseBody {
	s.FormInstId = &v
	return s
}

func (s *GetFormDataByIDResponseBody) SetFormData(v map[string]interface{}) *GetFormDataByIDResponseBody {
	s.FormData = v
	return s
}

type GetFormDataByIDResponseBodyOriginator struct {
	// 用户工号
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户名
	Name *GetFormDataByIDResponseBodyOriginatorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	// 部门名称
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 邮箱
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
}

func (s GetFormDataByIDResponseBodyOriginator) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDResponseBodyOriginator) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDResponseBodyOriginator) SetUserId(v string) *GetFormDataByIDResponseBodyOriginator {
	s.UserId = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginator) SetName(v *GetFormDataByIDResponseBodyOriginatorName) *GetFormDataByIDResponseBodyOriginator {
	s.Name = v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginator) SetDepartmentName(v string) *GetFormDataByIDResponseBodyOriginator {
	s.DepartmentName = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginator) SetEmail(v string) *GetFormDataByIDResponseBodyOriginator {
	s.Email = &v
	return s
}

type GetFormDataByIDResponseBodyOriginatorName struct {
	// 中文名称
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// 英文名称
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	// 国际化
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetFormDataByIDResponseBodyOriginatorName) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDResponseBodyOriginatorName) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDResponseBodyOriginatorName) SetNameInChinese(v string) *GetFormDataByIDResponseBodyOriginatorName {
	s.NameInChinese = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginatorName) SetNameInEnglish(v string) *GetFormDataByIDResponseBodyOriginatorName {
	s.NameInEnglish = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginatorName) SetType(v string) *GetFormDataByIDResponseBodyOriginatorName {
	s.Type = &v
	return s
}

type GetFormDataByIDResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFormDataByIDResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFormDataByIDResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDResponse) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDResponse) SetHeaders(v map[string]*string) *GetFormDataByIDResponse {
	s.Headers = v
	return s
}

func (s *GetFormDataByIDResponse) SetBody(v *GetFormDataByIDResponseBody) *GetFormDataByIDResponse {
	s.Body = v
	return s
}

type SaveFormDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveFormDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveFormDataHeaders) GoString() string {
	return s.String()
}

func (s *SaveFormDataHeaders) SetCommonHeaders(v map[string]*string) *SaveFormDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveFormDataHeaders) SetXAcsDingtalkAccessToken(v string) *SaveFormDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveFormDataRequest struct {
	// 应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 应用秘钥。在应用数据中获取。
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 语言。可选值：zh_CN/en_US 默认：zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 表单ID
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 表单数据
	FormDataJson *string `json:"formDataJson,omitempty" xml:"formDataJson,omitempty"`
}

func (s SaveFormDataRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveFormDataRequest) GoString() string {
	return s.String()
}

func (s *SaveFormDataRequest) SetAppType(v string) *SaveFormDataRequest {
	s.AppType = &v
	return s
}

func (s *SaveFormDataRequest) SetSystemToken(v string) *SaveFormDataRequest {
	s.SystemToken = &v
	return s
}

func (s *SaveFormDataRequest) SetUserId(v string) *SaveFormDataRequest {
	s.UserId = &v
	return s
}

func (s *SaveFormDataRequest) SetLanguage(v string) *SaveFormDataRequest {
	s.Language = &v
	return s
}

func (s *SaveFormDataRequest) SetFormUuid(v string) *SaveFormDataRequest {
	s.FormUuid = &v
	return s
}

func (s *SaveFormDataRequest) SetFormDataJson(v string) *SaveFormDataRequest {
	s.FormDataJson = &v
	return s
}

type SaveFormDataResponseBody struct {
	// 表单实例ID
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SaveFormDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveFormDataResponseBody) GoString() string {
	return s.String()
}

func (s *SaveFormDataResponseBody) SetResult(v string) *SaveFormDataResponseBody {
	s.Result = &v
	return s
}

type SaveFormDataResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveFormDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveFormDataResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveFormDataResponse) GoString() string {
	return s.String()
}

func (s *SaveFormDataResponse) SetHeaders(v map[string]*string) *SaveFormDataResponse {
	s.Headers = v
	return s
}

func (s *SaveFormDataResponse) SetBody(v *SaveFormDataResponseBody) *SaveFormDataResponse {
	s.Body = v
	return s
}

type LoginCodeGenHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s LoginCodeGenHeaders) String() string {
	return tea.Prettify(s)
}

func (s LoginCodeGenHeaders) GoString() string {
	return s.String()
}

func (s *LoginCodeGenHeaders) SetCommonHeaders(v map[string]*string) *LoginCodeGenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *LoginCodeGenHeaders) SetXAcsDingtalkAccessToken(v string) *LoginCodeGenHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type LoginCodeGenRequest struct {
	// userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s LoginCodeGenRequest) String() string {
	return tea.Prettify(s)
}

func (s LoginCodeGenRequest) GoString() string {
	return s.String()
}

func (s *LoginCodeGenRequest) SetUserId(v string) *LoginCodeGenRequest {
	s.UserId = &v
	return s
}

type LoginCodeGenResponseBody struct {
	// result
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s LoginCodeGenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LoginCodeGenResponseBody) GoString() string {
	return s.String()
}

func (s *LoginCodeGenResponseBody) SetResult(v string) *LoginCodeGenResponseBody {
	s.Result = &v
	return s
}

type LoginCodeGenResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *LoginCodeGenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LoginCodeGenResponse) String() string {
	return tea.Prettify(s)
}

func (s LoginCodeGenResponse) GoString() string {
	return s.String()
}

func (s *LoginCodeGenResponse) SetHeaders(v map[string]*string) *LoginCodeGenResponse {
	s.Headers = v
	return s
}

func (s *LoginCodeGenResponse) SetBody(v *LoginCodeGenResponseBody) *LoginCodeGenResponse {
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

func (client *Client) GetFormDataByID(id *string, request *GetFormDataByIDRequest) (_result *GetFormDataByIDResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFormDataByIDHeaders{}
	_result = &GetFormDataByIDResponse{}
	_body, _err := client.GetFormDataByIDWithOptions(id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFormDataByIDWithOptions(id *string, request *GetFormDataByIDRequest, headers *GetFormDataByIDHeaders, runtime *util.RuntimeOptions) (_result *GetFormDataByIDResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
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
		Query:   openapiutil.Query(query),
	}
	_result = &GetFormDataByIDResponse{}
	_body, _err := client.DoROARequest(tea.String("GetFormDataByID"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/forms/instances/"+tea.StringValue(id)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveFormData(request *SaveFormDataRequest) (_result *SaveFormDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveFormDataHeaders{}
	_result = &SaveFormDataResponse{}
	_body, _err := client.SaveFormDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveFormDataWithOptions(request *SaveFormDataRequest, headers *SaveFormDataHeaders, runtime *util.RuntimeOptions) (_result *SaveFormDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.FormDataJson)) {
		body["formDataJson"] = request.FormDataJson
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
	_result = &SaveFormDataResponse{}
	_body, _err := client.DoROARequest(tea.String("SaveFormData"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/forms/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LoginCodeGen(request *LoginCodeGenRequest) (_result *LoginCodeGenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &LoginCodeGenHeaders{}
	_result = &LoginCodeGenResponse{}
	_body, _err := client.LoginCodeGenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LoginCodeGenWithOptions(request *LoginCodeGenRequest, headers *LoginCodeGenHeaders, runtime *util.RuntimeOptions) (_result *LoginCodeGenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
		Query:   openapiutil.Query(query),
	}
	_result = &LoginCodeGenResponse{}
	_body, _err := client.DoROARequest(tea.String("LoginCodeGen"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/authorizations/loginCodes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
