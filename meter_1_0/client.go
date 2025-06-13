// This file is auto-generated, don't edit it. Thanks.
package meter_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetResourceUseInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetResourceUseInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetResourceUseInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetResourceUseInfoHeaders) SetCommonHeaders(v map[string]*string) *GetResourceUseInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetResourceUseInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetResourceUseInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetResourceUseInfoRequest struct {
	// This parameter is required.
	BenefitCodeList []*string `json:"benefitCodeList,omitempty" xml:"benefitCodeList,omitempty" type:"Repeated"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s GetResourceUseInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceUseInfoRequest) GoString() string {
	return s.String()
}

func (s *GetResourceUseInfoRequest) SetBenefitCodeList(v []*string) *GetResourceUseInfoRequest {
	s.BenefitCodeList = v
	return s
}

func (s *GetResourceUseInfoRequest) SetPeriod(v string) *GetResourceUseInfoRequest {
	s.Period = &v
	return s
}

type GetResourceUseInfoResponseBody struct {
	Result []*GetResourceUseInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetResourceUseInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceUseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceUseInfoResponseBody) SetResult(v []*GetResourceUseInfoResponseBodyResult) *GetResourceUseInfoResponseBody {
	s.Result = v
	return s
}

type GetResourceUseInfoResponseBodyResult struct {
	QuotaNum *int64 `json:"quotaNum,omitempty" xml:"quotaNum,omitempty"`
	UsedNum  *int64 `json:"usedNum,omitempty" xml:"usedNum,omitempty"`
}

func (s GetResourceUseInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetResourceUseInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetResourceUseInfoResponseBodyResult) SetQuotaNum(v int64) *GetResourceUseInfoResponseBodyResult {
	s.QuotaNum = &v
	return s
}

func (s *GetResourceUseInfoResponseBodyResult) SetUsedNum(v int64) *GetResourceUseInfoResponseBodyResult {
	s.UsedNum = &v
	return s
}

type GetResourceUseInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceUseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceUseInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceUseInfoResponse) GoString() string {
	return s.String()
}

func (s *GetResourceUseInfoResponse) SetHeaders(v map[string]*string) *GetResourceUseInfoResponse {
	s.Headers = v
	return s
}

func (s *GetResourceUseInfoResponse) SetStatusCode(v int32) *GetResourceUseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceUseInfoResponse) SetBody(v *GetResourceUseInfoResponseBody) *GetResourceUseInfoResponse {
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
// 获取开放平台当月资源用量
//
// @param request - GetResourceUseInfoRequest
//
// @param headers - GetResourceUseInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceUseInfoResponse
func (client *Client) GetResourceUseInfoWithOptions(request *GetResourceUseInfoRequest, headers *GetResourceUseInfoHeaders, runtime *util.RuntimeOptions) (_result *GetResourceUseInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BenefitCodeList)) {
		body["benefitCodeList"] = request.BenefitCodeList
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		body["period"] = request.Period
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
		Action:      tea.String("GetResourceUseInfo"),
		Version:     tea.String("meter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/meter/resources/useInfos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceUseInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取开放平台当月资源用量
//
// @param request - GetResourceUseInfoRequest
//
// @return GetResourceUseInfoResponse
func (client *Client) GetResourceUseInfo(request *GetResourceUseInfoRequest) (_result *GetResourceUseInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetResourceUseInfoHeaders{}
	_result = &GetResourceUseInfoResponse{}
	_body, _err := client.GetResourceUseInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
