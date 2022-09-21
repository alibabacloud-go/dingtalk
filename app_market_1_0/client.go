// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package app_market_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateAppGoodsServiceConversationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateAppGoodsServiceConversationHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGoodsServiceConversationHeaders) GoString() string {
	return s.String()
}

func (s *CreateAppGoodsServiceConversationHeaders) SetCommonHeaders(v map[string]*string) *CreateAppGoodsServiceConversationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateAppGoodsServiceConversationHeaders) SetXAcsDingtalkAccessToken(v string) *CreateAppGoodsServiceConversationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateAppGoodsServiceConversationRequest struct {
	IsvUserId *string `json:"isvUserId,omitempty" xml:"isvUserId,omitempty"`
	OrderId   *int64  `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s CreateAppGoodsServiceConversationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGoodsServiceConversationRequest) GoString() string {
	return s.String()
}

func (s *CreateAppGoodsServiceConversationRequest) SetIsvUserId(v string) *CreateAppGoodsServiceConversationRequest {
	s.IsvUserId = &v
	return s
}

func (s *CreateAppGoodsServiceConversationRequest) SetOrderId(v int64) *CreateAppGoodsServiceConversationRequest {
	s.OrderId = &v
	return s
}

type CreateAppGoodsServiceConversationResponseBody struct {
	// 群名称
	ConversationName *string `json:"conversationName,omitempty" xml:"conversationName,omitempty"`
	// 是否新群
	NewConversation *bool `json:"newConversation,omitempty" xml:"newConversation,omitempty"`
}

func (s CreateAppGoodsServiceConversationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGoodsServiceConversationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppGoodsServiceConversationResponseBody) SetConversationName(v string) *CreateAppGoodsServiceConversationResponseBody {
	s.ConversationName = &v
	return s
}

func (s *CreateAppGoodsServiceConversationResponseBody) SetNewConversation(v bool) *CreateAppGoodsServiceConversationResponseBody {
	s.NewConversation = &v
	return s
}

type CreateAppGoodsServiceConversationResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAppGoodsServiceConversationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppGoodsServiceConversationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGoodsServiceConversationResponse) GoString() string {
	return s.String()
}

func (s *CreateAppGoodsServiceConversationResponse) SetHeaders(v map[string]*string) *CreateAppGoodsServiceConversationResponse {
	s.Headers = v
	return s
}

func (s *CreateAppGoodsServiceConversationResponse) SetBody(v *CreateAppGoodsServiceConversationResponseBody) *CreateAppGoodsServiceConversationResponse {
	s.Body = v
	return s
}

type GetCoolAppAccessStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCoolAppAccessStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCoolAppAccessStatusHeaders) GoString() string {
	return s.String()
}

func (s *GetCoolAppAccessStatusHeaders) SetCommonHeaders(v map[string]*string) *GetCoolAppAccessStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCoolAppAccessStatusHeaders) SetXAcsDingtalkAccessToken(v string) *GetCoolAppAccessStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCoolAppAccessStatusRequest struct {
	// 免登授权码
	AuthCode *string `json:"authCode,omitempty" xml:"authCode,omitempty"`
	// 酷应用的code
	CoolAppCode *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	// 加密的场域业务code
	EncFieldBizCode *string `json:"encFieldBizCode,omitempty" xml:"encFieldBizCode,omitempty"`
}

func (s GetCoolAppAccessStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCoolAppAccessStatusRequest) GoString() string {
	return s.String()
}

func (s *GetCoolAppAccessStatusRequest) SetAuthCode(v string) *GetCoolAppAccessStatusRequest {
	s.AuthCode = &v
	return s
}

func (s *GetCoolAppAccessStatusRequest) SetCoolAppCode(v string) *GetCoolAppAccessStatusRequest {
	s.CoolAppCode = &v
	return s
}

func (s *GetCoolAppAccessStatusRequest) SetEncFieldBizCode(v string) *GetCoolAppAccessStatusRequest {
	s.EncFieldBizCode = &v
	return s
}

type GetCoolAppAccessStatusResponseBody struct {
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetCoolAppAccessStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCoolAppAccessStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetCoolAppAccessStatusResponseBody) SetStatus(v string) *GetCoolAppAccessStatusResponseBody {
	s.Status = &v
	return s
}

type GetCoolAppAccessStatusResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCoolAppAccessStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCoolAppAccessStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCoolAppAccessStatusResponse) GoString() string {
	return s.String()
}

func (s *GetCoolAppAccessStatusResponse) SetHeaders(v map[string]*string) *GetCoolAppAccessStatusResponse {
	s.Headers = v
	return s
}

func (s *GetCoolAppAccessStatusResponse) SetBody(v *GetCoolAppAccessStatusResponseBody) *GetCoolAppAccessStatusResponse {
	s.Body = v
	return s
}

type GetPersonalExperienceInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPersonalExperienceInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPersonalExperienceInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetPersonalExperienceInfoHeaders) SetCommonHeaders(v map[string]*string) *GetPersonalExperienceInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPersonalExperienceInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetPersonalExperienceInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPersonalExperienceInfoRequest struct {
	// A short description of struct
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetPersonalExperienceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPersonalExperienceInfoRequest) GoString() string {
	return s.String()
}

func (s *GetPersonalExperienceInfoRequest) SetUserId(v string) *GetPersonalExperienceInfoRequest {
	s.UserId = &v
	return s
}

type GetPersonalExperienceInfoResponseBody struct {
	// 数据对象
	Result *GetPersonalExperienceInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetPersonalExperienceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPersonalExperienceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetPersonalExperienceInfoResponseBody) SetResult(v *GetPersonalExperienceInfoResponseBodyResult) *GetPersonalExperienceInfoResponseBody {
	s.Result = v
	return s
}

type GetPersonalExperienceInfoResponseBodyResult struct {
	// 主组织corpId
	MainCorpId *string `json:"mainCorpId,omitempty" xml:"mainCorpId,omitempty"`
}

func (s GetPersonalExperienceInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetPersonalExperienceInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetPersonalExperienceInfoResponseBodyResult) SetMainCorpId(v string) *GetPersonalExperienceInfoResponseBodyResult {
	s.MainCorpId = &v
	return s
}

type GetPersonalExperienceInfoResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPersonalExperienceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPersonalExperienceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPersonalExperienceInfoResponse) GoString() string {
	return s.String()
}

func (s *GetPersonalExperienceInfoResponse) SetHeaders(v map[string]*string) *GetPersonalExperienceInfoResponse {
	s.Headers = v
	return s
}

func (s *GetPersonalExperienceInfoResponse) SetBody(v *GetPersonalExperienceInfoResponseBody) *GetPersonalExperienceInfoResponse {
	s.Body = v
	return s
}

type QueryMarketOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMarketOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMarketOrderHeaders) GoString() string {
	return s.String()
}

func (s *QueryMarketOrderHeaders) SetCommonHeaders(v map[string]*string) *QueryMarketOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMarketOrderHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMarketOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMarketOrderResponseBody struct {
	// 订单ID
	BizOrderId *int64 `json:"bizOrderId,omitempty" xml:"bizOrderId,omitempty"`
	// 组织ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 创建时间戳
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// 生效结束时间
	EndTimestamp *int64 `json:"endTimestamp,omitempty" xml:"endTimestamp,omitempty"`
	// 商品Code
	GoodsCode *string `json:"goodsCode,omitempty" xml:"goodsCode,omitempty"`
	// 商品名称
	GoodsName *string `json:"goodsName,omitempty" xml:"goodsName,omitempty"`
	// 是否内购订单
	InAppOrder *bool `json:"inAppOrder,omitempty" xml:"inAppOrder,omitempty"`
	// 规格编码
	ItemCode *string `json:"itemCode,omitempty" xml:"itemCode,omitempty"`
	// 规格名称
	ItemName *string `json:"itemName,omitempty" xml:"itemName,omitempty"`
	// 支付时间戳
	PaidTimestamp *int64 `json:"paidTimestamp,omitempty" xml:"paidTimestamp,omitempty"`
	// 购买数量
	Quantity *int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 开始生效时间
	StartTimestamp *int64 `json:"startTimestamp,omitempty" xml:"startTimestamp,omitempty"`
	// 订单状态(0:订单关闭； 3：订单支付；4：订单创建)
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 订单实付金额(单位分)
	TotalActualPayFee *int64 `json:"totalActualPayFee,omitempty" xml:"totalActualPayFee,omitempty"`
}

func (s QueryMarketOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMarketOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMarketOrderResponseBody) SetBizOrderId(v int64) *QueryMarketOrderResponseBody {
	s.BizOrderId = &v
	return s
}

func (s *QueryMarketOrderResponseBody) SetCorpId(v string) *QueryMarketOrderResponseBody {
	s.CorpId = &v
	return s
}

func (s *QueryMarketOrderResponseBody) SetCreateTimestamp(v int64) *QueryMarketOrderResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *QueryMarketOrderResponseBody) SetEndTimestamp(v int64) *QueryMarketOrderResponseBody {
	s.EndTimestamp = &v
	return s
}

func (s *QueryMarketOrderResponseBody) SetGoodsCode(v string) *QueryMarketOrderResponseBody {
	s.GoodsCode = &v
	return s
}

func (s *QueryMarketOrderResponseBody) SetGoodsName(v string) *QueryMarketOrderResponseBody {
	s.GoodsName = &v
	return s
}

func (s *QueryMarketOrderResponseBody) SetInAppOrder(v bool) *QueryMarketOrderResponseBody {
	s.InAppOrder = &v
	return s
}

func (s *QueryMarketOrderResponseBody) SetItemCode(v string) *QueryMarketOrderResponseBody {
	s.ItemCode = &v
	return s
}

func (s *QueryMarketOrderResponseBody) SetItemName(v string) *QueryMarketOrderResponseBody {
	s.ItemName = &v
	return s
}

func (s *QueryMarketOrderResponseBody) SetPaidTimestamp(v int64) *QueryMarketOrderResponseBody {
	s.PaidTimestamp = &v
	return s
}

func (s *QueryMarketOrderResponseBody) SetQuantity(v int64) *QueryMarketOrderResponseBody {
	s.Quantity = &v
	return s
}

func (s *QueryMarketOrderResponseBody) SetStartTimestamp(v int64) *QueryMarketOrderResponseBody {
	s.StartTimestamp = &v
	return s
}

func (s *QueryMarketOrderResponseBody) SetStatus(v int64) *QueryMarketOrderResponseBody {
	s.Status = &v
	return s
}

func (s *QueryMarketOrderResponseBody) SetTotalActualPayFee(v int64) *QueryMarketOrderResponseBody {
	s.TotalActualPayFee = &v
	return s
}

type QueryMarketOrderResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMarketOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMarketOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMarketOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryMarketOrderResponse) SetHeaders(v map[string]*string) *QueryMarketOrderResponse {
	s.Headers = v
	return s
}

func (s *QueryMarketOrderResponse) SetBody(v *QueryMarketOrderResponseBody) *QueryMarketOrderResponse {
	s.Body = v
	return s
}

type UserTaskReportHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UserTaskReportHeaders) String() string {
	return tea.Prettify(s)
}

func (s UserTaskReportHeaders) GoString() string {
	return s.String()
}

func (s *UserTaskReportHeaders) SetCommonHeaders(v map[string]*string) *UserTaskReportHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UserTaskReportHeaders) SetXAcsDingtalkAccessToken(v string) *UserTaskReportHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UserTaskReportRequest struct {
	// 业务的幂等ID
	BizNo *string `json:"bizNo,omitempty" xml:"bizNo,omitempty"`
	// operateDate
	OperateDate *string `json:"operateDate,omitempty" xml:"operateDate,omitempty"`
	// taskTag
	TaskTag *string `json:"taskTag,omitempty" xml:"taskTag,omitempty"`
	// staffId
	Userid *string `json:"userid,omitempty" xml:"userid,omitempty"`
}

func (s UserTaskReportRequest) String() string {
	return tea.Prettify(s)
}

func (s UserTaskReportRequest) GoString() string {
	return s.String()
}

func (s *UserTaskReportRequest) SetBizNo(v string) *UserTaskReportRequest {
	s.BizNo = &v
	return s
}

func (s *UserTaskReportRequest) SetOperateDate(v string) *UserTaskReportRequest {
	s.OperateDate = &v
	return s
}

func (s *UserTaskReportRequest) SetTaskTag(v string) *UserTaskReportRequest {
	s.TaskTag = &v
	return s
}

func (s *UserTaskReportRequest) SetUserid(v string) *UserTaskReportRequest {
	s.Userid = &v
	return s
}

type UserTaskReportResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *bool              `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UserTaskReportResponse) String() string {
	return tea.Prettify(s)
}

func (s UserTaskReportResponse) GoString() string {
	return s.String()
}

func (s *UserTaskReportResponse) SetHeaders(v map[string]*string) *UserTaskReportResponse {
	s.Headers = v
	return s
}

func (s *UserTaskReportResponse) SetBody(v bool) *UserTaskReportResponse {
	s.Body = &v
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

func (client *Client) CreateAppGoodsServiceConversation(request *CreateAppGoodsServiceConversationRequest) (_result *CreateAppGoodsServiceConversationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateAppGoodsServiceConversationHeaders{}
	_result = &CreateAppGoodsServiceConversationResponse{}
	_body, _err := client.CreateAppGoodsServiceConversationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppGoodsServiceConversationWithOptions(request *CreateAppGoodsServiceConversationRequest, headers *CreateAppGoodsServiceConversationHeaders, runtime *util.RuntimeOptions) (_result *CreateAppGoodsServiceConversationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsvUserId)) {
		body["isvUserId"] = request.IsvUserId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		body["orderId"] = request.OrderId
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
	_result = &CreateAppGoodsServiceConversationResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateAppGoodsServiceConversation"), tea.String("appMarket_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/appMarket/orders/serviceGroups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCoolAppAccessStatus(request *GetCoolAppAccessStatusRequest) (_result *GetCoolAppAccessStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCoolAppAccessStatusHeaders{}
	_result = &GetCoolAppAccessStatusResponse{}
	_body, _err := client.GetCoolAppAccessStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCoolAppAccessStatusWithOptions(request *GetCoolAppAccessStatusRequest, headers *GetCoolAppAccessStatusHeaders, runtime *util.RuntimeOptions) (_result *GetCoolAppAccessStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		body["authCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.CoolAppCode)) {
		body["coolAppCode"] = request.CoolAppCode
	}

	if !tea.BoolValue(util.IsUnset(request.EncFieldBizCode)) {
		body["encFieldBizCode"] = request.EncFieldBizCode
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
	_result = &GetCoolAppAccessStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("GetCoolAppAccessStatus"), tea.String("appMarket_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/appMarket/coolApps/accessions/statuses/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPersonalExperienceInfo(request *GetPersonalExperienceInfoRequest) (_result *GetPersonalExperienceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPersonalExperienceInfoHeaders{}
	_result = &GetPersonalExperienceInfoResponse{}
	_body, _err := client.GetPersonalExperienceInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPersonalExperienceInfoWithOptions(request *GetPersonalExperienceInfoRequest, headers *GetPersonalExperienceInfoHeaders, runtime *util.RuntimeOptions) (_result *GetPersonalExperienceInfoResponse, _err error) {
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &GetPersonalExperienceInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetPersonalExperienceInfo"), tea.String("appMarket_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/appMarket/personalExperiences"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMarketOrder(orderId *string) (_result *QueryMarketOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMarketOrderHeaders{}
	_result = &QueryMarketOrderResponse{}
	_body, _err := client.QueryMarketOrderWithOptions(orderId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMarketOrderWithOptions(orderId *string, headers *QueryMarketOrderHeaders, runtime *util.RuntimeOptions) (_result *QueryMarketOrderResponse, _err error) {
	orderId = openapiutil.GetEncodeParam(orderId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &QueryMarketOrderResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryMarketOrder"), tea.String("appMarket_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/appMarket/orders/"+tea.StringValue(orderId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UserTaskReport(request *UserTaskReportRequest) (_result *UserTaskReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UserTaskReportHeaders{}
	_result = &UserTaskReportResponse{}
	_body, _err := client.UserTaskReportWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UserTaskReportWithOptions(request *UserTaskReportRequest, headers *UserTaskReportHeaders, runtime *util.RuntimeOptions) (_result *UserTaskReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizNo)) {
		body["bizNo"] = request.BizNo
	}

	if !tea.BoolValue(util.IsUnset(request.OperateDate)) {
		body["operateDate"] = request.OperateDate
	}

	if !tea.BoolValue(util.IsUnset(request.TaskTag)) {
		body["taskTag"] = request.TaskTag
	}

	if !tea.BoolValue(util.IsUnset(request.Userid)) {
		body["userid"] = request.Userid
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
	_result = &UserTaskReportResponse{}
	_body, _err := client.DoROARequest(tea.String("UserTaskReport"), tea.String("appMarket_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/appMarket/tasks"), tea.String("boolean"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
