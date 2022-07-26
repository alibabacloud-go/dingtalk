// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package trip_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type SyncSecretKeyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SyncSecretKeyHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncSecretKeyHeaders) GoString() string {
	return s.String()
}

func (s *SyncSecretKeyHeaders) SetCommonHeaders(v map[string]*string) *SyncSecretKeyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncSecretKeyHeaders) SetXAcsDingtalkAccessToken(v string) *SyncSecretKeyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SyncSecretKeyRequest struct {
	// 操作类型，ADD/QUERY/MODIFY/DEL
	ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// 验签加密串
	SecretString *string `json:"secretString,omitempty" xml:"secretString,omitempty"`
	// 钉钉侧对应的组织ID
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
	// 商旅侧appkey
	TripAppKey *string `json:"tripAppKey,omitempty" xml:"tripAppKey,omitempty"`
	// 商旅对接密钥
	TripAppSecurity *string `json:"tripAppSecurity,omitempty" xml:"tripAppSecurity,omitempty"`
	// 商旅侧组织ID
	TripCorpId *string `json:"tripCorpId,omitempty" xml:"tripCorpId,omitempty"`
}

func (s SyncSecretKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncSecretKeyRequest) GoString() string {
	return s.String()
}

func (s *SyncSecretKeyRequest) SetActionType(v string) *SyncSecretKeyRequest {
	s.ActionType = &v
	return s
}

func (s *SyncSecretKeyRequest) SetSecretString(v string) *SyncSecretKeyRequest {
	s.SecretString = &v
	return s
}

func (s *SyncSecretKeyRequest) SetTargetCorpId(v string) *SyncSecretKeyRequest {
	s.TargetCorpId = &v
	return s
}

func (s *SyncSecretKeyRequest) SetTripAppKey(v string) *SyncSecretKeyRequest {
	s.TripAppKey = &v
	return s
}

func (s *SyncSecretKeyRequest) SetTripAppSecurity(v string) *SyncSecretKeyRequest {
	s.TripAppSecurity = &v
	return s
}

func (s *SyncSecretKeyRequest) SetTripCorpId(v string) *SyncSecretKeyRequest {
	s.TripCorpId = &v
	return s
}

type SyncSecretKeyResponseBody struct {
	Result *SyncSecretKeyResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// 是否成功
	//
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SyncSecretKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncSecretKeyResponseBody) GoString() string {
	return s.String()
}

func (s *SyncSecretKeyResponseBody) SetResult(v *SyncSecretKeyResponseBodyResult) *SyncSecretKeyResponseBody {
	s.Result = v
	return s
}

func (s *SyncSecretKeyResponseBody) SetSuccess(v string) *SyncSecretKeyResponseBody {
	s.Success = &v
	return s
}

type SyncSecretKeyResponseBodyResult struct {
	// 验签加密串
	SecretString *string `json:"secretString,omitempty" xml:"secretString,omitempty"`
	// 钉钉侧对应的组织ID
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
	// 商旅侧对接key
	TripAppKey *string `json:"tripAppKey,omitempty" xml:"tripAppKey,omitempty"`
	// 商旅侧对接密钥
	TripAppSecurity *string `json:"tripAppSecurity,omitempty" xml:"tripAppSecurity,omitempty"`
	// 商旅侧对应的组织ID
	TripCorpId *string `json:"tripCorpId,omitempty" xml:"tripCorpId,omitempty"`
}

func (s SyncSecretKeyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SyncSecretKeyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SyncSecretKeyResponseBodyResult) SetSecretString(v string) *SyncSecretKeyResponseBodyResult {
	s.SecretString = &v
	return s
}

func (s *SyncSecretKeyResponseBodyResult) SetTargetCorpId(v string) *SyncSecretKeyResponseBodyResult {
	s.TargetCorpId = &v
	return s
}

func (s *SyncSecretKeyResponseBodyResult) SetTripAppKey(v string) *SyncSecretKeyResponseBodyResult {
	s.TripAppKey = &v
	return s
}

func (s *SyncSecretKeyResponseBodyResult) SetTripAppSecurity(v string) *SyncSecretKeyResponseBodyResult {
	s.TripAppSecurity = &v
	return s
}

func (s *SyncSecretKeyResponseBodyResult) SetTripCorpId(v string) *SyncSecretKeyResponseBodyResult {
	s.TripCorpId = &v
	return s
}

type SyncSecretKeyResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SyncSecretKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncSecretKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncSecretKeyResponse) GoString() string {
	return s.String()
}

func (s *SyncSecretKeyResponse) SetHeaders(v map[string]*string) *SyncSecretKeyResponse {
	s.Headers = v
	return s
}

func (s *SyncSecretKeyResponse) SetBody(v *SyncSecretKeyResponseBody) *SyncSecretKeyResponse {
	s.Body = v
	return s
}

type SyncTripOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SyncTripOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncTripOrderHeaders) GoString() string {
	return s.String()
}

func (s *SyncTripOrderHeaders) SetCommonHeaders(v map[string]*string) *SyncTripOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncTripOrderHeaders) SetXAcsDingtalkAccessToken(v string) *SyncTripOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SyncTripOrderRequest struct {
	// 币种
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 钉钉用户id
	DingUserId *string `json:"dingUserId,omitempty" xml:"dingUserId,omitempty"`
	// 优惠金额
	DiscountAmount *string `json:"discountAmount,omitempty" xml:"discountAmount,omitempty"`
	// 是否是改签单
	EndorseFlag *bool                      `json:"endorseFlag,omitempty" xml:"endorseFlag,omitempty"`
	Event       *SyncTripOrderRequestEvent `json:"event,omitempty" xml:"event,omitempty" type:"Struct"`
	// 下单时间
	GmtOrder *string `json:"gmtOrder,omitempty" xml:"gmtOrder,omitempty"`
	// 付款时间
	GmtPay *string `json:"gmtPay,omitempty" xml:"gmtPay,omitempty"`
	// 退款时间
	GmtRefund *string `json:"gmtRefund,omitempty" xml:"gmtRefund,omitempty"`
	// 发票申请链接
	InvoiceApplyUrl *string `json:"invoiceApplyUrl,omitempty" xml:"invoiceApplyUrl,omitempty"`
	// 行程单号
	JourneyBizNo *string `json:"journeyBizNo,omitempty" xml:"journeyBizNo,omitempty"`
	// 订单详情列表
	OrderDetails []*SyncTripOrderRequestOrderDetails `json:"orderDetails,omitempty" xml:"orderDetails,omitempty" type:"Repeated"`
	// 供应商订单号
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	// 订单详情链接
	OrderUrl *string `json:"orderUrl,omitempty" xml:"orderUrl,omitempty"`
	// 实付金额
	RealAmount *string `json:"realAmount,omitempty" xml:"realAmount,omitempty"`
	// 退款金额
	RefundAmount *string `json:"refundAmount,omitempty" xml:"refundAmount,omitempty"`
	// 供应商关联订单号
	RelativeOrderNo *string `json:"relativeOrderNo,omitempty" xml:"relativeOrderNo,omitempty"`
	// 来源埋点
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 用户组织id
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
	// 总金额
	TotalAmount *string `json:"totalAmount,omitempty" xml:"totalAmount,omitempty"`
	// 订单类型
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SyncTripOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncTripOrderRequest) GoString() string {
	return s.String()
}

func (s *SyncTripOrderRequest) SetCurrency(v string) *SyncTripOrderRequest {
	s.Currency = &v
	return s
}

func (s *SyncTripOrderRequest) SetDingUserId(v string) *SyncTripOrderRequest {
	s.DingUserId = &v
	return s
}

func (s *SyncTripOrderRequest) SetDiscountAmount(v string) *SyncTripOrderRequest {
	s.DiscountAmount = &v
	return s
}

func (s *SyncTripOrderRequest) SetEndorseFlag(v bool) *SyncTripOrderRequest {
	s.EndorseFlag = &v
	return s
}

func (s *SyncTripOrderRequest) SetEvent(v *SyncTripOrderRequestEvent) *SyncTripOrderRequest {
	s.Event = v
	return s
}

func (s *SyncTripOrderRequest) SetGmtOrder(v string) *SyncTripOrderRequest {
	s.GmtOrder = &v
	return s
}

func (s *SyncTripOrderRequest) SetGmtPay(v string) *SyncTripOrderRequest {
	s.GmtPay = &v
	return s
}

func (s *SyncTripOrderRequest) SetGmtRefund(v string) *SyncTripOrderRequest {
	s.GmtRefund = &v
	return s
}

func (s *SyncTripOrderRequest) SetInvoiceApplyUrl(v string) *SyncTripOrderRequest {
	s.InvoiceApplyUrl = &v
	return s
}

func (s *SyncTripOrderRequest) SetJourneyBizNo(v string) *SyncTripOrderRequest {
	s.JourneyBizNo = &v
	return s
}

func (s *SyncTripOrderRequest) SetOrderDetails(v []*SyncTripOrderRequestOrderDetails) *SyncTripOrderRequest {
	s.OrderDetails = v
	return s
}

func (s *SyncTripOrderRequest) SetOrderNo(v string) *SyncTripOrderRequest {
	s.OrderNo = &v
	return s
}

func (s *SyncTripOrderRequest) SetOrderUrl(v string) *SyncTripOrderRequest {
	s.OrderUrl = &v
	return s
}

func (s *SyncTripOrderRequest) SetRealAmount(v string) *SyncTripOrderRequest {
	s.RealAmount = &v
	return s
}

func (s *SyncTripOrderRequest) SetRefundAmount(v string) *SyncTripOrderRequest {
	s.RefundAmount = &v
	return s
}

func (s *SyncTripOrderRequest) SetRelativeOrderNo(v string) *SyncTripOrderRequest {
	s.RelativeOrderNo = &v
	return s
}

func (s *SyncTripOrderRequest) SetSource(v string) *SyncTripOrderRequest {
	s.Source = &v
	return s
}

func (s *SyncTripOrderRequest) SetTargetCorpId(v string) *SyncTripOrderRequest {
	s.TargetCorpId = &v
	return s
}

func (s *SyncTripOrderRequest) SetTotalAmount(v string) *SyncTripOrderRequest {
	s.TotalAmount = &v
	return s
}

func (s *SyncTripOrderRequest) SetType(v string) *SyncTripOrderRequest {
	s.Type = &v
	return s
}

type SyncTripOrderRequestEvent struct {
	// 订单事件
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 事件时间
	GmtAction *string `json:"gmtAction,omitempty" xml:"gmtAction,omitempty"`
}

func (s SyncTripOrderRequestEvent) String() string {
	return tea.Prettify(s)
}

func (s SyncTripOrderRequestEvent) GoString() string {
	return s.String()
}

func (s *SyncTripOrderRequestEvent) SetAction(v string) *SyncTripOrderRequestEvent {
	s.Action = &v
	return s
}

func (s *SyncTripOrderRequestEvent) SetGmtAction(v string) *SyncTripOrderRequestEvent {
	s.GmtAction = &v
	return s
}

type SyncTripOrderRequestOrderDetails struct {
	// 到达时间
	ArrivalTime *string `json:"arrivalTime,omitempty" xml:"arrivalTime,omitempty"`
	// 车辆颜色
	CarColor *string `json:"carColor,omitempty" xml:"carColor,omitempty"`
	// 车辆型号
	CarModel *string `json:"carModel,omitempty" xml:"carModel,omitempty"`
	// 车牌号
	CarNumber *string `json:"carNumber,omitempty" xml:"carNumber,omitempty"`
	// 餐食描述
	CateringType *string `json:"cateringType,omitempty" xml:"cateringType,omitempty"`
	// 入住时间
	CheckInTime *string `json:"checkInTime,omitempty" xml:"checkInTime,omitempty"`
	// 离店时间
	CheckOutTime *string `json:"checkOutTime,omitempty" xml:"checkOutTime,omitempty"`
	// 出发时间
	DepartTime *string `json:"departTime,omitempty" xml:"departTime,omitempty"`
	// 目的地城市码
	DestinationCityCode *string `json:"destinationCityCode,omitempty" xml:"destinationCityCode,omitempty"`
	// 目的站名称
	DestinationStation *string `json:"destinationStation,omitempty" xml:"destinationStation,omitempty"`
	// 酒店地址
	HotelAddress *string `json:"hotelAddress,omitempty" xml:"hotelAddress,omitempty"`
	// 酒店定位信息
	HotelLocation *SyncTripOrderRequestOrderDetailsHotelLocation `json:"hotelLocation,omitempty" xml:"hotelLocation,omitempty" type:"Struct"`
	// 酒店名称
	HotelName *string `json:"hotelName,omitempty" xml:"hotelName,omitempty"`
	// 出发地城市码
	OriginCityCode *string `json:"originCityCode,omitempty" xml:"originCityCode,omitempty"`
	// 出发站名称
	OriginStation *string `json:"originStation,omitempty" xml:"originStation,omitempty"`
	// 房间数
	RoomCount *int32 `json:"roomCount,omitempty" xml:"roomCount,omitempty"`
	// 舱位
	SeatInfo *string `json:"seatInfo,omitempty" xml:"seatInfo,omitempty"`
	// “服务类型”
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	// 下游供应商logo
	SubSupplyLogo *string `json:"subSupplyLogo,omitempty" xml:"subSupplyLogo,omitempty"`
	// 下游供应商名称
	SubSupplyName *string `json:"subSupplyName,omitempty" xml:"subSupplyName,omitempty"`
	// 专车类型
	TaxiType *string `json:"taxiType,omitempty" xml:"taxiType,omitempty"`
	// 联系方式
	Telephone *string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	// 火车/航班班次
	TransportNumber *string `json:"transportNumber,omitempty" xml:"transportNumber,omitempty"`
	// 房型描述
	TypeDescription *string `json:"typeDescription,omitempty" xml:"typeDescription,omitempty"`
}

func (s SyncTripOrderRequestOrderDetails) String() string {
	return tea.Prettify(s)
}

func (s SyncTripOrderRequestOrderDetails) GoString() string {
	return s.String()
}

func (s *SyncTripOrderRequestOrderDetails) SetArrivalTime(v string) *SyncTripOrderRequestOrderDetails {
	s.ArrivalTime = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetCarColor(v string) *SyncTripOrderRequestOrderDetails {
	s.CarColor = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetCarModel(v string) *SyncTripOrderRequestOrderDetails {
	s.CarModel = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetCarNumber(v string) *SyncTripOrderRequestOrderDetails {
	s.CarNumber = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetCateringType(v string) *SyncTripOrderRequestOrderDetails {
	s.CateringType = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetCheckInTime(v string) *SyncTripOrderRequestOrderDetails {
	s.CheckInTime = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetCheckOutTime(v string) *SyncTripOrderRequestOrderDetails {
	s.CheckOutTime = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetDepartTime(v string) *SyncTripOrderRequestOrderDetails {
	s.DepartTime = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetDestinationCityCode(v string) *SyncTripOrderRequestOrderDetails {
	s.DestinationCityCode = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetDestinationStation(v string) *SyncTripOrderRequestOrderDetails {
	s.DestinationStation = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetHotelAddress(v string) *SyncTripOrderRequestOrderDetails {
	s.HotelAddress = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetHotelLocation(v *SyncTripOrderRequestOrderDetailsHotelLocation) *SyncTripOrderRequestOrderDetails {
	s.HotelLocation = v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetHotelName(v string) *SyncTripOrderRequestOrderDetails {
	s.HotelName = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetOriginCityCode(v string) *SyncTripOrderRequestOrderDetails {
	s.OriginCityCode = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetOriginStation(v string) *SyncTripOrderRequestOrderDetails {
	s.OriginStation = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetRoomCount(v int32) *SyncTripOrderRequestOrderDetails {
	s.RoomCount = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetSeatInfo(v string) *SyncTripOrderRequestOrderDetails {
	s.SeatInfo = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetServiceType(v string) *SyncTripOrderRequestOrderDetails {
	s.ServiceType = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetSubSupplyLogo(v string) *SyncTripOrderRequestOrderDetails {
	s.SubSupplyLogo = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetSubSupplyName(v string) *SyncTripOrderRequestOrderDetails {
	s.SubSupplyName = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetTaxiType(v string) *SyncTripOrderRequestOrderDetails {
	s.TaxiType = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetTelephone(v string) *SyncTripOrderRequestOrderDetails {
	s.Telephone = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetTransportNumber(v string) *SyncTripOrderRequestOrderDetails {
	s.TransportNumber = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetTypeDescription(v string) *SyncTripOrderRequestOrderDetails {
	s.TypeDescription = &v
	return s
}

type SyncTripOrderRequestOrderDetailsHotelLocation struct {
	// 纬度
	Lat *string `json:"lat,omitempty" xml:"lat,omitempty"`
	// 经度
	Lon *string `json:"lon,omitempty" xml:"lon,omitempty"`
	// 坐标数据源
	// - BD09：来自百度地图的经纬坐标
	// - GCJ02: 来自高德地图，腾讯地图，Apple地图的坐标
	// - WGS84: 来自GPS的坐标
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 定位url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SyncTripOrderRequestOrderDetailsHotelLocation) String() string {
	return tea.Prettify(s)
}

func (s SyncTripOrderRequestOrderDetailsHotelLocation) GoString() string {
	return s.String()
}

func (s *SyncTripOrderRequestOrderDetailsHotelLocation) SetLat(v string) *SyncTripOrderRequestOrderDetailsHotelLocation {
	s.Lat = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetailsHotelLocation) SetLon(v string) *SyncTripOrderRequestOrderDetailsHotelLocation {
	s.Lon = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetailsHotelLocation) SetSource(v string) *SyncTripOrderRequestOrderDetailsHotelLocation {
	s.Source = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetailsHotelLocation) SetUrl(v string) *SyncTripOrderRequestOrderDetailsHotelLocation {
	s.Url = &v
	return s
}

type SyncTripOrderResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SyncTripOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncTripOrderResponseBody) GoString() string {
	return s.String()
}

func (s *SyncTripOrderResponseBody) SetRequestId(v string) *SyncTripOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncTripOrderResponseBody) SetSuccess(v bool) *SyncTripOrderResponseBody {
	s.Success = &v
	return s
}

type SyncTripOrderResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SyncTripOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncTripOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncTripOrderResponse) GoString() string {
	return s.String()
}

func (s *SyncTripOrderResponse) SetHeaders(v map[string]*string) *SyncTripOrderResponse {
	s.Headers = v
	return s
}

func (s *SyncTripOrderResponse) SetBody(v *SyncTripOrderResponseBody) *SyncTripOrderResponse {
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

func (client *Client) SyncSecretKey(request *SyncSecretKeyRequest) (_result *SyncSecretKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncSecretKeyHeaders{}
	_result = &SyncSecretKeyResponse{}
	_body, _err := client.SyncSecretKeyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncSecretKeyWithOptions(request *SyncSecretKeyRequest, headers *SyncSecretKeyHeaders, runtime *util.RuntimeOptions) (_result *SyncSecretKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionType)) {
		body["actionType"] = request.ActionType
	}

	if !tea.BoolValue(util.IsUnset(request.SecretString)) {
		body["secretString"] = request.SecretString
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		body["targetCorpId"] = request.TargetCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.TripAppKey)) {
		body["tripAppKey"] = request.TripAppKey
	}

	if !tea.BoolValue(util.IsUnset(request.TripAppSecurity)) {
		body["tripAppSecurity"] = request.TripAppSecurity
	}

	if !tea.BoolValue(util.IsUnset(request.TripCorpId)) {
		body["tripCorpId"] = request.TripCorpId
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
	_result = &SyncSecretKeyResponse{}
	_body, _err := client.DoROARequest(tea.String("SyncSecretKey"), tea.String("trip_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/trip/secretKeys/sync"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncTripOrder(request *SyncTripOrderRequest) (_result *SyncTripOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncTripOrderHeaders{}
	_result = &SyncTripOrderResponse{}
	_body, _err := client.SyncTripOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncTripOrderWithOptions(request *SyncTripOrderRequest, headers *SyncTripOrderHeaders, runtime *util.RuntimeOptions) (_result *SyncTripOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Currency)) {
		body["currency"] = request.Currency
	}

	if !tea.BoolValue(util.IsUnset(request.DingUserId)) {
		body["dingUserId"] = request.DingUserId
	}

	if !tea.BoolValue(util.IsUnset(request.DiscountAmount)) {
		body["discountAmount"] = request.DiscountAmount
	}

	if !tea.BoolValue(util.IsUnset(request.EndorseFlag)) {
		body["endorseFlag"] = request.EndorseFlag
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Event))) {
		body["event"] = request.Event
	}

	if !tea.BoolValue(util.IsUnset(request.GmtOrder)) {
		body["gmtOrder"] = request.GmtOrder
	}

	if !tea.BoolValue(util.IsUnset(request.GmtPay)) {
		body["gmtPay"] = request.GmtPay
	}

	if !tea.BoolValue(util.IsUnset(request.GmtRefund)) {
		body["gmtRefund"] = request.GmtRefund
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceApplyUrl)) {
		body["invoiceApplyUrl"] = request.InvoiceApplyUrl
	}

	if !tea.BoolValue(util.IsUnset(request.JourneyBizNo)) {
		body["journeyBizNo"] = request.JourneyBizNo
	}

	if !tea.BoolValue(util.IsUnset(request.OrderDetails)) {
		body["orderDetails"] = request.OrderDetails
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNo)) {
		body["orderNo"] = request.OrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.OrderUrl)) {
		body["orderUrl"] = request.OrderUrl
	}

	if !tea.BoolValue(util.IsUnset(request.RealAmount)) {
		body["realAmount"] = request.RealAmount
	}

	if !tea.BoolValue(util.IsUnset(request.RefundAmount)) {
		body["refundAmount"] = request.RefundAmount
	}

	if !tea.BoolValue(util.IsUnset(request.RelativeOrderNo)) {
		body["relativeOrderNo"] = request.RelativeOrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		body["targetCorpId"] = request.TargetCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.TotalAmount)) {
		body["totalAmount"] = request.TotalAmount
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
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
	_result = &SyncTripOrderResponse{}
	_body, _err := client.DoROARequest(tea.String("SyncTripOrder"), tea.String("trip_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/trip/tripOrders/sync"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
