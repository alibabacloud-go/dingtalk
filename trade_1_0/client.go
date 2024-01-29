// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package trade_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CheckOpportunityResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CheckOpportunityResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s CheckOpportunityResultHeaders) GoString() string {
	return s.String()
}

func (s *CheckOpportunityResultHeaders) SetCommonHeaders(v map[string]*string) *CheckOpportunityResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CheckOpportunityResultHeaders) SetXAcsDingtalkAccessToken(v string) *CheckOpportunityResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CheckOpportunityResultRequest struct {
	BelongToPhoneNum *string `json:"belongToPhoneNum,omitempty" xml:"belongToPhoneNum,omitempty"`
	ContactPhoneNum  *string `json:"contactPhoneNum,omitempty" xml:"contactPhoneNum,omitempty"`
	CorpId           *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeptId           *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	MarketCode       *string `json:"marketCode,omitempty" xml:"marketCode,omitempty"`
}

func (s CheckOpportunityResultRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckOpportunityResultRequest) GoString() string {
	return s.String()
}

func (s *CheckOpportunityResultRequest) SetBelongToPhoneNum(v string) *CheckOpportunityResultRequest {
	s.BelongToPhoneNum = &v
	return s
}

func (s *CheckOpportunityResultRequest) SetContactPhoneNum(v string) *CheckOpportunityResultRequest {
	s.ContactPhoneNum = &v
	return s
}

func (s *CheckOpportunityResultRequest) SetCorpId(v string) *CheckOpportunityResultRequest {
	s.CorpId = &v
	return s
}

func (s *CheckOpportunityResultRequest) SetDeptId(v int64) *CheckOpportunityResultRequest {
	s.DeptId = &v
	return s
}

func (s *CheckOpportunityResultRequest) SetMarketCode(v string) *CheckOpportunityResultRequest {
	s.MarketCode = &v
	return s
}

type CheckOpportunityResultResponseBody struct {
	BizSuccess *bool `json:"bizSuccess,omitempty" xml:"bizSuccess,omitempty"`
}

func (s CheckOpportunityResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckOpportunityResultResponseBody) GoString() string {
	return s.String()
}

func (s *CheckOpportunityResultResponseBody) SetBizSuccess(v bool) *CheckOpportunityResultResponseBody {
	s.BizSuccess = &v
	return s
}

type CheckOpportunityResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckOpportunityResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckOpportunityResultResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckOpportunityResultResponse) GoString() string {
	return s.String()
}

func (s *CheckOpportunityResultResponse) SetHeaders(v map[string]*string) *CheckOpportunityResultResponse {
	s.Headers = v
	return s
}

func (s *CheckOpportunityResultResponse) SetStatusCode(v int32) *CheckOpportunityResultResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckOpportunityResultResponse) SetBody(v *CheckOpportunityResultResponseBody) *CheckOpportunityResultResponse {
	s.Body = v
	return s
}

type CreateOpportunityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateOpportunityHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateOpportunityHeaders) GoString() string {
	return s.String()
}

func (s *CreateOpportunityHeaders) SetCommonHeaders(v map[string]*string) *CreateOpportunityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateOpportunityHeaders) SetXAcsDingtalkAccessToken(v string) *CreateOpportunityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateOpportunityRequest struct {
	BelongToPhoneNum *string `json:"belongToPhoneNum,omitempty" xml:"belongToPhoneNum,omitempty"`
	ContactPhoneNum  *string `json:"contactPhoneNum,omitempty" xml:"contactPhoneNum,omitempty"`
	CorpId           *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeptId           *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	MarketCode       *string `json:"marketCode,omitempty" xml:"marketCode,omitempty"`
}

func (s CreateOpportunityRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOpportunityRequest) GoString() string {
	return s.String()
}

func (s *CreateOpportunityRequest) SetBelongToPhoneNum(v string) *CreateOpportunityRequest {
	s.BelongToPhoneNum = &v
	return s
}

func (s *CreateOpportunityRequest) SetContactPhoneNum(v string) *CreateOpportunityRequest {
	s.ContactPhoneNum = &v
	return s
}

func (s *CreateOpportunityRequest) SetCorpId(v string) *CreateOpportunityRequest {
	s.CorpId = &v
	return s
}

func (s *CreateOpportunityRequest) SetDeptId(v int64) *CreateOpportunityRequest {
	s.DeptId = &v
	return s
}

func (s *CreateOpportunityRequest) SetMarketCode(v string) *CreateOpportunityRequest {
	s.MarketCode = &v
	return s
}

type CreateOpportunityResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateOpportunityResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOpportunityResponse) GoString() string {
	return s.String()
}

func (s *CreateOpportunityResponse) SetHeaders(v map[string]*string) *CreateOpportunityResponse {
	s.Headers = v
	return s
}

func (s *CreateOpportunityResponse) SetStatusCode(v int32) *CreateOpportunityResponse {
	s.StatusCode = &v
	return s
}

type QueryTradeOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryTradeOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeOrderHeaders) GoString() string {
	return s.String()
}

func (s *QueryTradeOrderHeaders) SetCommonHeaders(v map[string]*string) *QueryTradeOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryTradeOrderHeaders) SetXAcsDingtalkAccessToken(v string) *QueryTradeOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryTradeOrderRequest struct {
	OrderId      *int64  `json:"orderId,omitempty" xml:"orderId,omitempty"`
	OuterOrderId *string `json:"outerOrderId,omitempty" xml:"outerOrderId,omitempty"`
}

func (s QueryTradeOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryTradeOrderRequest) SetOrderId(v int64) *QueryTradeOrderRequest {
	s.OrderId = &v
	return s
}

func (s *QueryTradeOrderRequest) SetOuterOrderId(v string) *QueryTradeOrderRequest {
	s.OuterOrderId = &v
	return s
}

type QueryTradeOrderResponseBody struct {
	ArticleCode     *string `json:"articleCode,omitempty" xml:"articleCode,omitempty"`
	ArticleItemCode *string `json:"articleItemCode,omitempty" xml:"articleItemCode,omitempty"`
	ArticleItemName *string `json:"articleItemName,omitempty" xml:"articleItemName,omitempty"`
	ArticleName     *string `json:"articleName,omitempty" xml:"articleName,omitempty"`
	CloseTime       *int64  `json:"closeTime,omitempty" xml:"closeTime,omitempty"`
	CreateTime      *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Fee             *int64  `json:"fee,omitempty" xml:"fee,omitempty"`
	IsvCropId       *string `json:"isvCropId,omitempty" xml:"isvCropId,omitempty"`
	OrderId         *int64  `json:"orderId,omitempty" xml:"orderId,omitempty"`
	OuterOrderId    *string `json:"outerOrderId,omitempty" xml:"outerOrderId,omitempty"`
	PayFee          *int64  `json:"payFee,omitempty" xml:"payFee,omitempty"`
	PayTime         *int64  `json:"payTime,omitempty" xml:"payTime,omitempty"`
	Quantity        *int64  `json:"quantity,omitempty" xml:"quantity,omitempty"`
	RefundTime      *int64  `json:"refundTime,omitempty" xml:"refundTime,omitempty"`
	Status          *int32  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryTradeOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTradeOrderResponseBody) SetArticleCode(v string) *QueryTradeOrderResponseBody {
	s.ArticleCode = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetArticleItemCode(v string) *QueryTradeOrderResponseBody {
	s.ArticleItemCode = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetArticleItemName(v string) *QueryTradeOrderResponseBody {
	s.ArticleItemName = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetArticleName(v string) *QueryTradeOrderResponseBody {
	s.ArticleName = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetCloseTime(v int64) *QueryTradeOrderResponseBody {
	s.CloseTime = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetCreateTime(v int64) *QueryTradeOrderResponseBody {
	s.CreateTime = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetFee(v int64) *QueryTradeOrderResponseBody {
	s.Fee = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetIsvCropId(v string) *QueryTradeOrderResponseBody {
	s.IsvCropId = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetOrderId(v int64) *QueryTradeOrderResponseBody {
	s.OrderId = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetOuterOrderId(v string) *QueryTradeOrderResponseBody {
	s.OuterOrderId = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetPayFee(v int64) *QueryTradeOrderResponseBody {
	s.PayFee = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetPayTime(v int64) *QueryTradeOrderResponseBody {
	s.PayTime = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetQuantity(v int64) *QueryTradeOrderResponseBody {
	s.Quantity = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetRefundTime(v int64) *QueryTradeOrderResponseBody {
	s.RefundTime = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetStatus(v int32) *QueryTradeOrderResponseBody {
	s.Status = &v
	return s
}

type QueryTradeOrderResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTradeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTradeOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryTradeOrderResponse) SetHeaders(v map[string]*string) *QueryTradeOrderResponse {
	s.Headers = v
	return s
}

func (s *QueryTradeOrderResponse) SetStatusCode(v int32) *QueryTradeOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTradeOrderResponse) SetBody(v *QueryTradeOrderResponseBody) *QueryTradeOrderResponse {
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
	interfaceSPI, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = interfaceSPI
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

func (client *Client) CheckOpportunityResultWithOptions(request *CheckOpportunityResultRequest, headers *CheckOpportunityResultHeaders, runtime *util.RuntimeOptions) (_result *CheckOpportunityResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BelongToPhoneNum)) {
		query["belongToPhoneNum"] = request.BelongToPhoneNum
	}

	if !tea.BoolValue(util.IsUnset(request.ContactPhoneNum)) {
		query["contactPhoneNum"] = request.ContactPhoneNum
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.MarketCode)) {
		query["marketCode"] = request.MarketCode
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
		Action:      tea.String("CheckOpportunityResult"),
		Version:     tea.String("trade_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/trade/opportunity/check"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckOpportunityResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckOpportunityResult(request *CheckOpportunityResultRequest) (_result *CheckOpportunityResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CheckOpportunityResultHeaders{}
	_result = &CheckOpportunityResultResponse{}
	_body, _err := client.CheckOpportunityResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOpportunityWithOptions(request *CreateOpportunityRequest, headers *CreateOpportunityHeaders, runtime *util.RuntimeOptions) (_result *CreateOpportunityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BelongToPhoneNum)) {
		body["belongToPhoneNum"] = request.BelongToPhoneNum
	}

	if !tea.BoolValue(util.IsUnset(request.ContactPhoneNum)) {
		body["contactPhoneNum"] = request.ContactPhoneNum
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.MarketCode)) {
		body["marketCode"] = request.MarketCode
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
		Action:      tea.String("CreateOpportunity"),
		Version:     tea.String("trade_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/trade/opportunities"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOpportunityResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOpportunity(request *CreateOpportunityRequest) (_result *CreateOpportunityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateOpportunityHeaders{}
	_result = &CreateOpportunityResponse{}
	_body, _err := client.CreateOpportunityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTradeOrderWithOptions(request *QueryTradeOrderRequest, headers *QueryTradeOrderHeaders, runtime *util.RuntimeOptions) (_result *QueryTradeOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		body["orderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OuterOrderId)) {
		body["outerOrderId"] = request.OuterOrderId
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
		Action:      tea.String("QueryTradeOrder"),
		Version:     tea.String("trade_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/trade/orders/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTradeOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTradeOrder(request *QueryTradeOrderRequest) (_result *QueryTradeOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryTradeOrderHeaders{}
	_result = &QueryTradeOrderResponse{}
	_body, _err := client.QueryTradeOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
