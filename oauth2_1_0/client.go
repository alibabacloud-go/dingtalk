// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package oauth2_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetUserTokenRequest struct {
	// 应用id
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// 应用密码
	ClientSecret *string `json:"clientSecret,omitempty" xml:"clientSecret,omitempty"`
	// OAuth 2.0 临时授权码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// OAuth 2.0 刷新令牌
	RefreshToken *string `json:"refreshToken,omitempty" xml:"refreshToken,omitempty"`
	// 分为authorization_code和refresh_token。使用授权码换token，传authorization_code；使用刷新token换用户token，传refresh_token
	GrantType *string `json:"grantType,omitempty" xml:"grantType,omitempty"`
}

func (s GetUserTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserTokenRequest) GoString() string {
	return s.String()
}

func (s *GetUserTokenRequest) SetClientId(v string) *GetUserTokenRequest {
	s.ClientId = &v
	return s
}

func (s *GetUserTokenRequest) SetClientSecret(v string) *GetUserTokenRequest {
	s.ClientSecret = &v
	return s
}

func (s *GetUserTokenRequest) SetCode(v string) *GetUserTokenRequest {
	s.Code = &v
	return s
}

func (s *GetUserTokenRequest) SetRefreshToken(v string) *GetUserTokenRequest {
	s.RefreshToken = &v
	return s
}

func (s *GetUserTokenRequest) SetGrantType(v string) *GetUserTokenRequest {
	s.GrantType = &v
	return s
}

type GetUserTokenResponseBody struct {
	// accessToken
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// refreshToken
	RefreshToken *string `json:"refreshToken,omitempty" xml:"refreshToken,omitempty"`
	// 超时时间
	ExpireIn *int64 `json:"expireIn,omitempty" xml:"expireIn,omitempty"`
	// 所选企业corpId
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s GetUserTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserTokenResponseBody) SetAccessToken(v string) *GetUserTokenResponseBody {
	s.AccessToken = &v
	return s
}

func (s *GetUserTokenResponseBody) SetRefreshToken(v string) *GetUserTokenResponseBody {
	s.RefreshToken = &v
	return s
}

func (s *GetUserTokenResponseBody) SetExpireIn(v int64) *GetUserTokenResponseBody {
	s.ExpireIn = &v
	return s
}

func (s *GetUserTokenResponseBody) SetCorpId(v string) *GetUserTokenResponseBody {
	s.CorpId = &v
	return s
}

type GetUserTokenResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserTokenResponse) GoString() string {
	return s.String()
}

func (s *GetUserTokenResponse) SetHeaders(v map[string]*string) *GetUserTokenResponse {
	s.Headers = v
	return s
}

func (s *GetUserTokenResponse) SetBody(v *GetUserTokenResponseBody) *GetUserTokenResponse {
	s.Body = v
	return s
}

type GetAccessTokenRequest struct {
	// 应用id
	AppKey *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	// 应用密码
	AppSecret *string `json:"appSecret,omitempty" xml:"appSecret,omitempty"`
}

func (s GetAccessTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *GetAccessTokenRequest) SetAppKey(v string) *GetAccessTokenRequest {
	s.AppKey = &v
	return s
}

func (s *GetAccessTokenRequest) SetAppSecret(v string) *GetAccessTokenRequest {
	s.AppSecret = &v
	return s
}

type GetAccessTokenResponseBody struct {
	// accessToken
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// 超时时间
	ExpireIn *int64 `json:"expireIn,omitempty" xml:"expireIn,omitempty"`
}

func (s GetAccessTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessTokenResponseBody) SetAccessToken(v string) *GetAccessTokenResponseBody {
	s.AccessToken = &v
	return s
}

func (s *GetAccessTokenResponseBody) SetExpireIn(v int64) *GetAccessTokenResponseBody {
	s.ExpireIn = &v
	return s
}

type GetAccessTokenResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetAccessTokenResponse) SetBody(v *GetAccessTokenResponseBody) *GetAccessTokenResponse {
	s.Body = v
	return s
}

type GetSuiteAccessTokenRequest struct {
	// 应用id
	SuiteKey *string `json:"suiteKey,omitempty" xml:"suiteKey,omitempty"`
	// 应用密码
	SuiteSecret *string `json:"suiteSecret,omitempty" xml:"suiteSecret,omitempty"`
	// suiteTicket
	SuiteTicket *string `json:"suiteTicket,omitempty" xml:"suiteTicket,omitempty"`
}

func (s GetSuiteAccessTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSuiteAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *GetSuiteAccessTokenRequest) SetSuiteKey(v string) *GetSuiteAccessTokenRequest {
	s.SuiteKey = &v
	return s
}

func (s *GetSuiteAccessTokenRequest) SetSuiteSecret(v string) *GetSuiteAccessTokenRequest {
	s.SuiteSecret = &v
	return s
}

func (s *GetSuiteAccessTokenRequest) SetSuiteTicket(v string) *GetSuiteAccessTokenRequest {
	s.SuiteTicket = &v
	return s
}

type GetSuiteAccessTokenResponseBody struct {
	// accessToken
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// 超时时间
	ExpireIn *int64 `json:"expireIn,omitempty" xml:"expireIn,omitempty"`
}

func (s GetSuiteAccessTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSuiteAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetSuiteAccessTokenResponseBody) SetAccessToken(v string) *GetSuiteAccessTokenResponseBody {
	s.AccessToken = &v
	return s
}

func (s *GetSuiteAccessTokenResponseBody) SetExpireIn(v int64) *GetSuiteAccessTokenResponseBody {
	s.ExpireIn = &v
	return s
}

type GetSuiteAccessTokenResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSuiteAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSuiteAccessTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSuiteAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *GetSuiteAccessTokenResponse) SetHeaders(v map[string]*string) *GetSuiteAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *GetSuiteAccessTokenResponse) SetBody(v *GetSuiteAccessTokenResponseBody) *GetSuiteAccessTokenResponse {
	s.Body = v
	return s
}

type GetCorpAccessTokenRequest struct {
	// 应用id
	SuiteKey *string `json:"suiteKey,omitempty" xml:"suiteKey,omitempty"`
	// 应用密码
	SuiteSecret *string `json:"suiteSecret,omitempty" xml:"suiteSecret,omitempty"`
	// OAuth 2.0 临时授权码
	AuthCorpId *string `json:"authCorpId,omitempty" xml:"authCorpId,omitempty"`
	// suiteTicket
	SuiteTicket *string `json:"suiteTicket,omitempty" xml:"suiteTicket,omitempty"`
}

func (s GetCorpAccessTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCorpAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *GetCorpAccessTokenRequest) SetSuiteKey(v string) *GetCorpAccessTokenRequest {
	s.SuiteKey = &v
	return s
}

func (s *GetCorpAccessTokenRequest) SetSuiteSecret(v string) *GetCorpAccessTokenRequest {
	s.SuiteSecret = &v
	return s
}

func (s *GetCorpAccessTokenRequest) SetAuthCorpId(v string) *GetCorpAccessTokenRequest {
	s.AuthCorpId = &v
	return s
}

func (s *GetCorpAccessTokenRequest) SetSuiteTicket(v string) *GetCorpAccessTokenRequest {
	s.SuiteTicket = &v
	return s
}

type GetCorpAccessTokenResponseBody struct {
	// accessToken
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// 超时时间
	ExpireIn *int64 `json:"expireIn,omitempty" xml:"expireIn,omitempty"`
}

func (s GetCorpAccessTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCorpAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetCorpAccessTokenResponseBody) SetAccessToken(v string) *GetCorpAccessTokenResponseBody {
	s.AccessToken = &v
	return s
}

func (s *GetCorpAccessTokenResponseBody) SetExpireIn(v int64) *GetCorpAccessTokenResponseBody {
	s.ExpireIn = &v
	return s
}

type GetCorpAccessTokenResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCorpAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCorpAccessTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCorpAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *GetCorpAccessTokenResponse) SetHeaders(v map[string]*string) *GetCorpAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *GetCorpAccessTokenResponse) SetBody(v *GetCorpAccessTokenResponseBody) *GetCorpAccessTokenResponse {
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

func (client *Client) GetUserToken(request *GetUserTokenRequest) (_result *GetUserTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserTokenResponse{}
	_body, _err := client.GetUserTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserTokenWithOptions(request *GetUserTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetUserTokenResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.RefreshToken)) {
		body["refreshToken"] = request.RefreshToken
	}

	if !tea.BoolValue(util.IsUnset(request.GrantType)) {
		body["grantType"] = request.GrantType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &GetUserTokenResponse{}
	_body, _err := client.DoROARequest(tea.String("GetUserToken"), tea.String("oauth2_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/oauth2/userAccessToken"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccessToken(request *GetAccessTokenRequest) (_result *GetAccessTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAccessTokenResponse{}
	_body, _err := client.GetAccessTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAccessTokenWithOptions(request *GetAccessTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAccessTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["appKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.AppSecret)) {
		body["appSecret"] = request.AppSecret
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &GetAccessTokenResponse{}
	_body, _err := client.DoROARequest(tea.String("GetAccessToken"), tea.String("oauth2_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/oauth2/accessToken"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSuiteAccessToken(request *GetSuiteAccessTokenRequest) (_result *GetSuiteAccessTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSuiteAccessTokenResponse{}
	_body, _err := client.GetSuiteAccessTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSuiteAccessTokenWithOptions(request *GetSuiteAccessTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSuiteAccessTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SuiteKey)) {
		body["suiteKey"] = request.SuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.SuiteSecret)) {
		body["suiteSecret"] = request.SuiteSecret
	}

	if !tea.BoolValue(util.IsUnset(request.SuiteTicket)) {
		body["suiteTicket"] = request.SuiteTicket
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &GetSuiteAccessTokenResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSuiteAccessToken"), tea.String("oauth2_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/oauth2/suiteAccessToken"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCorpAccessToken(request *GetCorpAccessTokenRequest) (_result *GetCorpAccessTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCorpAccessTokenResponse{}
	_body, _err := client.GetCorpAccessTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCorpAccessTokenWithOptions(request *GetCorpAccessTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCorpAccessTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SuiteKey)) {
		body["suiteKey"] = request.SuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.SuiteSecret)) {
		body["suiteSecret"] = request.SuiteSecret
	}

	if !tea.BoolValue(util.IsUnset(request.AuthCorpId)) {
		body["authCorpId"] = request.AuthCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.SuiteTicket)) {
		body["suiteTicket"] = request.SuiteTicket
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &GetCorpAccessTokenResponse{}
	_body, _err := client.DoROARequest(tea.String("GetCorpAccessToken"), tea.String("oauth2_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/oauth2/corpAccessToken"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
