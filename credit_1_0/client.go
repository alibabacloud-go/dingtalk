// This file is auto-generated, don't edit it. Thanks.
package credit_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type QueryScoreHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryScoreHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryScoreHeaders) GoString() string {
	return s.String()
}

func (s *QueryScoreHeaders) SetCommonHeaders(v map[string]*string) *QueryScoreHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryScoreHeaders) SetXAcsDingtalkAccessToken(v string) *QueryScoreHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryScoreRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// MD5
	Encryption *string `json:"encryption,omitempty" xml:"encryption,omitempty"`
	// example:
	//
	// a0fbf479272cd38c220fbf726678d8d6
	FullName *string `json:"fullName,omitempty" xml:"fullName,omitempty"`
	// example:
	//
	// b04a604cf00e64136b386e83444245c3
	IdCardCode *string `json:"idCardCode,omitempty" xml:"idCardCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e10adc3949ba59abbe56e057f20f883e
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// example:
	//
	// aca03c931768ea4b0244531aca9a19ee
	OrgName *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	// example:
	//
	// a57d7bf49b6e44180b21b1fea80eec0a
	UniScCode *string `json:"uniScCode,omitempty" xml:"uniScCode,omitempty"`
}

func (s QueryScoreRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryScoreRequest) GoString() string {
	return s.String()
}

func (s *QueryScoreRequest) SetEncryption(v string) *QueryScoreRequest {
	s.Encryption = &v
	return s
}

func (s *QueryScoreRequest) SetFullName(v string) *QueryScoreRequest {
	s.FullName = &v
	return s
}

func (s *QueryScoreRequest) SetIdCardCode(v string) *QueryScoreRequest {
	s.IdCardCode = &v
	return s
}

func (s *QueryScoreRequest) SetMobile(v string) *QueryScoreRequest {
	s.Mobile = &v
	return s
}

func (s *QueryScoreRequest) SetOrgName(v string) *QueryScoreRequest {
	s.OrgName = &v
	return s
}

func (s *QueryScoreRequest) SetUniScCode(v string) *QueryScoreRequest {
	s.UniScCode = &v
	return s
}

type QueryScoreResponseBody struct {
	Result  *QueryScoreResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryScoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryScoreResponseBody) GoString() string {
	return s.String()
}

func (s *QueryScoreResponseBody) SetResult(v *QueryScoreResponseBodyResult) *QueryScoreResponseBody {
	s.Result = v
	return s
}

func (s *QueryScoreResponseBody) SetSuccess(v bool) *QueryScoreResponseBody {
	s.Success = &v
	return s
}

type QueryScoreResponseBodyResult struct {
	CcocScore               *float64 `json:"ccocScore,omitempty" xml:"ccocScore,omitempty"`
	CityChangeCnt3y         *int64   `json:"cityChangeCnt3y,omitempty" xml:"cityChangeCnt3y,omitempty"`
	CityChangeTrend2y       *float64 `json:"cityChangeTrend2y,omitempty" xml:"cityChangeTrend2y,omitempty"`
	ClassificationOfOrg     *string  `json:"classificationOfOrg,omitempty" xml:"classificationOfOrg,omitempty"`
	GrowthRateLoginDays180d *float64 `json:"growthRateLoginDays180d,omitempty" xml:"growthRateLoginDays180d,omitempty"`
	IndChangeCnt3y          *int64   `json:"indChangeCnt3y,omitempty" xml:"indChangeCnt3y,omitempty"`
	IndChangeTrend2y        *float64 `json:"indChangeTrend2y,omitempty" xml:"indChangeTrend2y,omitempty"`
	JobChangeCnt3y          *int64   `json:"jobChangeCnt3y,omitempty" xml:"jobChangeCnt3y,omitempty"`
	JobTitle                *string  `json:"jobTitle,omitempty" xml:"jobTitle,omitempty"`
	JoinDays                *int64   `json:"joinDays,omitempty" xml:"joinDays,omitempty"`
	LoginDays14dPct         *float64 `json:"loginDays14dPct,omitempty" xml:"loginDays14dPct,omitempty"`
	LoginDays365dPct        *float64 `json:"loginDays365dPct,omitempty" xml:"loginDays365dPct,omitempty"`
	OrgCnt                  *int64   `json:"orgCnt,omitempty" xml:"orgCnt,omitempty"`
	OrgIndustrySubNameNew   *string  `json:"orgIndustrySubNameNew,omitempty" xml:"orgIndustrySubNameNew,omitempty"`
	OrgIndustryUpNameNew    *string  `json:"orgIndustryUpNameNew,omitempty" xml:"orgIndustryUpNameNew,omitempty"`
}

func (s QueryScoreResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryScoreResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryScoreResponseBodyResult) SetCcocScore(v float64) *QueryScoreResponseBodyResult {
	s.CcocScore = &v
	return s
}

func (s *QueryScoreResponseBodyResult) SetCityChangeCnt3y(v int64) *QueryScoreResponseBodyResult {
	s.CityChangeCnt3y = &v
	return s
}

func (s *QueryScoreResponseBodyResult) SetCityChangeTrend2y(v float64) *QueryScoreResponseBodyResult {
	s.CityChangeTrend2y = &v
	return s
}

func (s *QueryScoreResponseBodyResult) SetClassificationOfOrg(v string) *QueryScoreResponseBodyResult {
	s.ClassificationOfOrg = &v
	return s
}

func (s *QueryScoreResponseBodyResult) SetGrowthRateLoginDays180d(v float64) *QueryScoreResponseBodyResult {
	s.GrowthRateLoginDays180d = &v
	return s
}

func (s *QueryScoreResponseBodyResult) SetIndChangeCnt3y(v int64) *QueryScoreResponseBodyResult {
	s.IndChangeCnt3y = &v
	return s
}

func (s *QueryScoreResponseBodyResult) SetIndChangeTrend2y(v float64) *QueryScoreResponseBodyResult {
	s.IndChangeTrend2y = &v
	return s
}

func (s *QueryScoreResponseBodyResult) SetJobChangeCnt3y(v int64) *QueryScoreResponseBodyResult {
	s.JobChangeCnt3y = &v
	return s
}

func (s *QueryScoreResponseBodyResult) SetJobTitle(v string) *QueryScoreResponseBodyResult {
	s.JobTitle = &v
	return s
}

func (s *QueryScoreResponseBodyResult) SetJoinDays(v int64) *QueryScoreResponseBodyResult {
	s.JoinDays = &v
	return s
}

func (s *QueryScoreResponseBodyResult) SetLoginDays14dPct(v float64) *QueryScoreResponseBodyResult {
	s.LoginDays14dPct = &v
	return s
}

func (s *QueryScoreResponseBodyResult) SetLoginDays365dPct(v float64) *QueryScoreResponseBodyResult {
	s.LoginDays365dPct = &v
	return s
}

func (s *QueryScoreResponseBodyResult) SetOrgCnt(v int64) *QueryScoreResponseBodyResult {
	s.OrgCnt = &v
	return s
}

func (s *QueryScoreResponseBodyResult) SetOrgIndustrySubNameNew(v string) *QueryScoreResponseBodyResult {
	s.OrgIndustrySubNameNew = &v
	return s
}

func (s *QueryScoreResponseBodyResult) SetOrgIndustryUpNameNew(v string) *QueryScoreResponseBodyResult {
	s.OrgIndustryUpNameNew = &v
	return s
}

type QueryScoreResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryScoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryScoreResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryScoreResponse) GoString() string {
	return s.String()
}

func (s *QueryScoreResponse) SetHeaders(v map[string]*string) *QueryScoreResponse {
	s.Headers = v
	return s
}

func (s *QueryScoreResponse) SetStatusCode(v int32) *QueryScoreResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryScoreResponse) SetBody(v *QueryScoreResponseBody) *QueryScoreResponse {
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
// 查询用户金融评分数据
//
// @param request - QueryScoreRequest
//
// @param headers - QueryScoreHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryScoreResponse
func (client *Client) QueryScoreWithOptions(request *QueryScoreRequest, headers *QueryScoreHeaders, runtime *util.RuntimeOptions) (_result *QueryScoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Encryption)) {
		query["encryption"] = request.Encryption
	}

	if !tea.BoolValue(util.IsUnset(request.FullName)) {
		query["fullName"] = request.FullName
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardCode)) {
		query["idCardCode"] = request.IdCardCode
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		query["orgName"] = request.OrgName
	}

	if !tea.BoolValue(util.IsUnset(request.UniScCode)) {
		query["uniScCode"] = request.UniScCode
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
		Action:      tea.String("QueryScore"),
		Version:     tea.String("credit_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/credit/scores"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryScoreResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户金融评分数据
//
// @param request - QueryScoreRequest
//
// @return QueryScoreResponse
func (client *Client) QueryScore(request *QueryScoreRequest) (_result *QueryScoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryScoreHeaders{}
	_result = &QueryScoreResponse{}
	_body, _err := client.QueryScoreWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
