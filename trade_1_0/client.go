// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package trade_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	// 外部订单号
	OuterOrderId *string `json:"outerOrderId,omitempty" xml:"outerOrderId,omitempty"`
	// 内部订单号
	OrderId      *int64  `json:"orderId,omitempty" xml:"orderId,omitempty"`
	DingIsvOrgId *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingSuiteKey *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
}

func (s QueryTradeOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryTradeOrderRequest) SetOuterOrderId(v string) *QueryTradeOrderRequest {
	s.OuterOrderId = &v
	return s
}

func (s *QueryTradeOrderRequest) SetOrderId(v int64) *QueryTradeOrderRequest {
	s.OrderId = &v
	return s
}

func (s *QueryTradeOrderRequest) SetDingIsvOrgId(v int64) *QueryTradeOrderRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *QueryTradeOrderRequest) SetDingSuiteKey(v string) *QueryTradeOrderRequest {
	s.DingSuiteKey = &v
	return s
}

type QueryTradeOrderResponseBody struct {
	// ISV的组织ID
	IsvCropId *string `json:"isvCropId,omitempty" xml:"isvCropId,omitempty"`
	// 商品名称
	ArticleName *string `json:"articleName,omitempty" xml:"articleName,omitempty"`
	// 商品编码
	ArticleCode *string `json:"articleCode,omitempty" xml:"articleCode,omitempty"`
	// 规格名称
	ArticleItemName *string `json:"articleItemName,omitempty" xml:"articleItemName,omitempty"`
	// 规格编码
	ArticleItemCode *string `json:"articleItemCode,omitempty" xml:"articleItemCode,omitempty"`
	// 商品数量
	Quantity *int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 外部订单号
	OuterOrderId *string `json:"outerOrderId,omitempty" xml:"outerOrderId,omitempty"`
	// 内部订单号
	OrderId *int64 `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// 原价（单位：分）
	Fee *int64 `json:"fee,omitempty" xml:"fee,omitempty"`
	// 实际支付的价格（单位：分）
	PayFee *int64 `json:"payFee,omitempty" xml:"payFee,omitempty"`
	// 订单创建时间（单位：ms）
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 订单退款时间（单位：ms）
	RefundTime *int64 `json:"refundTime,omitempty" xml:"refundTime,omitempty"`
	// 订单关闭时间（单位：ms）
	CloseTime *int64 `json:"closeTime,omitempty" xml:"closeTime,omitempty"`
	// 订单支付时间（单位：ms）
	PayTime *int64 `json:"payTime,omitempty" xml:"payTime,omitempty"`
	// 订单状态：-1表示已关闭、0表示未支付、1表示已支付、2表示已退款
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryTradeOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTradeOrderResponseBody) SetIsvCropId(v string) *QueryTradeOrderResponseBody {
	s.IsvCropId = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetArticleName(v string) *QueryTradeOrderResponseBody {
	s.ArticleName = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetArticleCode(v string) *QueryTradeOrderResponseBody {
	s.ArticleCode = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetArticleItemName(v string) *QueryTradeOrderResponseBody {
	s.ArticleItemName = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetArticleItemCode(v string) *QueryTradeOrderResponseBody {
	s.ArticleItemCode = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetQuantity(v int64) *QueryTradeOrderResponseBody {
	s.Quantity = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetOuterOrderId(v string) *QueryTradeOrderResponseBody {
	s.OuterOrderId = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetOrderId(v int64) *QueryTradeOrderResponseBody {
	s.OrderId = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetFee(v int64) *QueryTradeOrderResponseBody {
	s.Fee = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetPayFee(v int64) *QueryTradeOrderResponseBody {
	s.PayFee = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetCreateTime(v int64) *QueryTradeOrderResponseBody {
	s.CreateTime = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetRefundTime(v int64) *QueryTradeOrderResponseBody {
	s.RefundTime = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetCloseTime(v int64) *QueryTradeOrderResponseBody {
	s.CloseTime = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetPayTime(v int64) *QueryTradeOrderResponseBody {
	s.PayTime = &v
	return s
}

func (s *QueryTradeOrderResponseBody) SetStatus(v int32) *QueryTradeOrderResponseBody {
	s.Status = &v
	return s
}

type QueryTradeOrderResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTradeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
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

func (client *Client) QueryTradeOrderWithOptions(request *QueryTradeOrderRequest, headers *QueryTradeOrderHeaders, runtime *util.RuntimeOptions) (_result *QueryTradeOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OuterOrderId)) {
		body["outerOrderId"] = request.OuterOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		body["orderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
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
	_result = &QueryTradeOrderResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryTradeOrder"), tea.String("trade_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/trade/orders/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
