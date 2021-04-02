// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package project_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetTbProjectSourceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	DingOrgId               *string            `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingIsvOrgId            *string            `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingCorpId              *string            `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	DingSuiteKey            *string            `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingAccessTokenType     *string            `json:"dingAccessTokenType,omitempty" xml:"dingAccessTokenType,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTbProjectSourceHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTbProjectSourceHeaders) GoString() string {
	return s.String()
}

func (s *GetTbProjectSourceHeaders) SetCommonHeaders(v map[string]*string) *GetTbProjectSourceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTbProjectSourceHeaders) SetDingOrgId(v string) *GetTbProjectSourceHeaders {
	s.DingOrgId = &v
	return s
}

func (s *GetTbProjectSourceHeaders) SetDingIsvOrgId(v string) *GetTbProjectSourceHeaders {
	s.DingIsvOrgId = &v
	return s
}

func (s *GetTbProjectSourceHeaders) SetDingCorpId(v string) *GetTbProjectSourceHeaders {
	s.DingCorpId = &v
	return s
}

func (s *GetTbProjectSourceHeaders) SetDingSuiteKey(v string) *GetTbProjectSourceHeaders {
	s.DingSuiteKey = &v
	return s
}

func (s *GetTbProjectSourceHeaders) SetDingAccessTokenType(v string) *GetTbProjectSourceHeaders {
	s.DingAccessTokenType = &v
	return s
}

func (s *GetTbProjectSourceHeaders) SetXAcsDingtalkAccessToken(v string) *GetTbProjectSourceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTbProjectSourceResponseBody struct {
	// 应用安装来源，"0"：来自应用中心，”6“：预安装
	InstallSource *string `json:"installSource,omitempty" xml:"installSource,omitempty"`
}

func (s GetTbProjectSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTbProjectSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetTbProjectSourceResponseBody) SetInstallSource(v string) *GetTbProjectSourceResponseBody {
	s.InstallSource = &v
	return s
}

type GetTbProjectSourceResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTbProjectSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTbProjectSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTbProjectSourceResponse) GoString() string {
	return s.String()
}

func (s *GetTbProjectSourceResponse) SetHeaders(v map[string]*string) *GetTbProjectSourceResponse {
	s.Headers = v
	return s
}

func (s *GetTbProjectSourceResponse) SetBody(v *GetTbProjectSourceResponseBody) *GetTbProjectSourceResponse {
	s.Body = v
	return s
}

type GetTbProjectGrayHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	DingAccessTokenType     *string            `json:"dingAccessTokenType,omitempty" xml:"dingAccessTokenType,omitempty"`
	DingSuiteKey            *string            `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingIsvOrgId            *string            `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId               *string            `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingCorpId              *string            `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTbProjectGrayHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTbProjectGrayHeaders) GoString() string {
	return s.String()
}

func (s *GetTbProjectGrayHeaders) SetCommonHeaders(v map[string]*string) *GetTbProjectGrayHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTbProjectGrayHeaders) SetDingAccessTokenType(v string) *GetTbProjectGrayHeaders {
	s.DingAccessTokenType = &v
	return s
}

func (s *GetTbProjectGrayHeaders) SetDingSuiteKey(v string) *GetTbProjectGrayHeaders {
	s.DingSuiteKey = &v
	return s
}

func (s *GetTbProjectGrayHeaders) SetDingIsvOrgId(v string) *GetTbProjectGrayHeaders {
	s.DingIsvOrgId = &v
	return s
}

func (s *GetTbProjectGrayHeaders) SetDingOrgId(v string) *GetTbProjectGrayHeaders {
	s.DingOrgId = &v
	return s
}

func (s *GetTbProjectGrayHeaders) SetDingCorpId(v string) *GetTbProjectGrayHeaders {
	s.DingCorpId = &v
	return s
}

func (s *GetTbProjectGrayHeaders) SetXAcsDingtalkAccessToken(v string) *GetTbProjectGrayHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTbProjectGrayRequest struct {
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
}

func (s GetTbProjectGrayRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTbProjectGrayRequest) GoString() string {
	return s.String()
}

func (s *GetTbProjectGrayRequest) SetLabel(v string) *GetTbProjectGrayRequest {
	s.Label = &v
	return s
}

type GetTbProjectGrayResponseBody struct {
	// 是否灰度
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetTbProjectGrayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTbProjectGrayResponseBody) GoString() string {
	return s.String()
}

func (s *GetTbProjectGrayResponseBody) SetResult(v bool) *GetTbProjectGrayResponseBody {
	s.Result = &v
	return s
}

func (s *GetTbProjectGrayResponseBody) SetRequestId(v string) *GetTbProjectGrayResponseBody {
	s.RequestId = &v
	return s
}

type GetTbProjectGrayResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTbProjectGrayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTbProjectGrayResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTbProjectGrayResponse) GoString() string {
	return s.String()
}

func (s *GetTbProjectGrayResponse) SetHeaders(v map[string]*string) *GetTbProjectGrayResponse {
	s.Headers = v
	return s
}

func (s *GetTbProjectGrayResponse) SetBody(v *GetTbProjectGrayResponseBody) *GetTbProjectGrayResponse {
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

func (client *Client) GetTbProjectSource() (_result *GetTbProjectSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTbProjectSourceHeaders{}
	_result = &GetTbProjectSourceResponse{}
	_body, _err := client.GetTbProjectSourceWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTbProjectSourceWithOptions(headers *GetTbProjectSourceHeaders, runtime *util.RuntimeOptions) (_result *GetTbProjectSourceResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.DingOrgId)) {
		realHeaders["dingOrgId"] = headers.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(headers.DingIsvOrgId)) {
		realHeaders["dingIsvOrgId"] = headers.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(headers.DingCorpId)) {
		realHeaders["dingCorpId"] = headers.DingCorpId
	}

	if !tea.BoolValue(util.IsUnset(headers.DingSuiteKey)) {
		realHeaders["dingSuiteKey"] = headers.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(headers.DingAccessTokenType)) {
		realHeaders["dingAccessTokenType"] = headers.DingAccessTokenType
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &GetTbProjectSourceResponse{}
	_body, _err := client.DoROARequest(tea.String("GetTbProjectSource"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/project/projects/source"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTbProjectGray(request *GetTbProjectGrayRequest) (_result *GetTbProjectGrayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTbProjectGrayHeaders{}
	_result = &GetTbProjectGrayResponse{}
	_body, _err := client.GetTbProjectGrayWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTbProjectGrayWithOptions(request *GetTbProjectGrayRequest, headers *GetTbProjectGrayHeaders, runtime *util.RuntimeOptions) (_result *GetTbProjectGrayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Label)) {
		body["label"] = request.Label
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.DingAccessTokenType)) {
		realHeaders["dingAccessTokenType"] = headers.DingAccessTokenType
	}

	if !tea.BoolValue(util.IsUnset(headers.DingSuiteKey)) {
		realHeaders["dingSuiteKey"] = headers.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(headers.DingIsvOrgId)) {
		realHeaders["dingIsvOrgId"] = headers.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(headers.DingOrgId)) {
		realHeaders["dingOrgId"] = headers.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(headers.DingCorpId)) {
		realHeaders["dingCorpId"] = headers.DingCorpId
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &GetTbProjectGrayResponse{}
	_body, _err := client.DoROARequest(tea.String("GetTbProjectGray"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/project/projects/gray"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
