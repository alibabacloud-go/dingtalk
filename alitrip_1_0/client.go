// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package alitrip_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ApproveCityCarApplyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ApproveCityCarApplyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ApproveCityCarApplyHeaders) GoString() string {
	return s.String()
}

func (s *ApproveCityCarApplyHeaders) SetCommonHeaders(v map[string]*string) *ApproveCityCarApplyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ApproveCityCarApplyHeaders) SetXAcsDingtalkAccessToken(v string) *ApproveCityCarApplyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ApproveCityCarApplyRequest struct {
	// 第三方企业ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 审批时间
	OperateTime *string `json:"operateTime,omitempty" xml:"operateTime,omitempty"`
	// 审批备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 审批结果：1-同意，2-拒绝
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 第三方审批单ID
	ThirdPartApplyId *string `json:"thirdPartApplyId,omitempty" xml:"thirdPartApplyId,omitempty"`
	// 审批的第三方员工ID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// suiteKey
	DingSuiteKey *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	// account
	DingCorpId *string `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	// tokenGrantType
	DingTokenGrantType *int64 `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
}

func (s ApproveCityCarApplyRequest) String() string {
	return tea.Prettify(s)
}

func (s ApproveCityCarApplyRequest) GoString() string {
	return s.String()
}

func (s *ApproveCityCarApplyRequest) SetCorpId(v string) *ApproveCityCarApplyRequest {
	s.CorpId = &v
	return s
}

func (s *ApproveCityCarApplyRequest) SetOperateTime(v string) *ApproveCityCarApplyRequest {
	s.OperateTime = &v
	return s
}

func (s *ApproveCityCarApplyRequest) SetRemark(v string) *ApproveCityCarApplyRequest {
	s.Remark = &v
	return s
}

func (s *ApproveCityCarApplyRequest) SetStatus(v int64) *ApproveCityCarApplyRequest {
	s.Status = &v
	return s
}

func (s *ApproveCityCarApplyRequest) SetThirdPartApplyId(v string) *ApproveCityCarApplyRequest {
	s.ThirdPartApplyId = &v
	return s
}

func (s *ApproveCityCarApplyRequest) SetUserId(v string) *ApproveCityCarApplyRequest {
	s.UserId = &v
	return s
}

func (s *ApproveCityCarApplyRequest) SetDingSuiteKey(v string) *ApproveCityCarApplyRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *ApproveCityCarApplyRequest) SetDingCorpId(v string) *ApproveCityCarApplyRequest {
	s.DingCorpId = &v
	return s
}

func (s *ApproveCityCarApplyRequest) SetDingTokenGrantType(v int64) *ApproveCityCarApplyRequest {
	s.DingTokenGrantType = &v
	return s
}

type ApproveCityCarApplyResponseBody struct {
	// 审批结果
	ApproveResult *bool `json:"approveResult,omitempty" xml:"approveResult,omitempty"`
}

func (s ApproveCityCarApplyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApproveCityCarApplyResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveCityCarApplyResponseBody) SetApproveResult(v bool) *ApproveCityCarApplyResponseBody {
	s.ApproveResult = &v
	return s
}

type ApproveCityCarApplyResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApproveCityCarApplyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApproveCityCarApplyResponse) String() string {
	return tea.Prettify(s)
}

func (s ApproveCityCarApplyResponse) GoString() string {
	return s.String()
}

func (s *ApproveCityCarApplyResponse) SetHeaders(v map[string]*string) *ApproveCityCarApplyResponse {
	s.Headers = v
	return s
}

func (s *ApproveCityCarApplyResponse) SetBody(v *ApproveCityCarApplyResponseBody) *ApproveCityCarApplyResponse {
	s.Body = v
	return s
}

type BillSettementHotelHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BillSettementHotelHeaders) String() string {
	return tea.Prettify(s)
}

func (s BillSettementHotelHeaders) GoString() string {
	return s.String()
}

func (s *BillSettementHotelHeaders) SetCommonHeaders(v map[string]*string) *BillSettementHotelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BillSettementHotelHeaders) SetXAcsDingtalkAccessToken(v string) *BillSettementHotelHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BillSettementHotelRequest struct {
	// 第三方企业
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 类目：机酒火车 1：机票； 2：酒店； 4：用车 6：商旅火车票
	Category *int64 `json:"category,omitempty" xml:"category,omitempty"`
	// 每页数据量，默认100，最高500
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 记账更新开始日期
	PeriodStart *string `json:"periodStart,omitempty" xml:"periodStart,omitempty"`
	// 页数，从1开始
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 记账更新结束日期
	PeriodEnd *string `json:"periodEnd,omitempty" xml:"periodEnd,omitempty"`
}

func (s BillSettementHotelRequest) String() string {
	return tea.Prettify(s)
}

func (s BillSettementHotelRequest) GoString() string {
	return s.String()
}

func (s *BillSettementHotelRequest) SetCorpId(v string) *BillSettementHotelRequest {
	s.CorpId = &v
	return s
}

func (s *BillSettementHotelRequest) SetCategory(v int64) *BillSettementHotelRequest {
	s.Category = &v
	return s
}

func (s *BillSettementHotelRequest) SetPageSize(v int64) *BillSettementHotelRequest {
	s.PageSize = &v
	return s
}

func (s *BillSettementHotelRequest) SetPeriodStart(v string) *BillSettementHotelRequest {
	s.PeriodStart = &v
	return s
}

func (s *BillSettementHotelRequest) SetPageNumber(v int64) *BillSettementHotelRequest {
	s.PageNumber = &v
	return s
}

func (s *BillSettementHotelRequest) SetPeriodEnd(v string) *BillSettementHotelRequest {
	s.PeriodEnd = &v
	return s
}

type BillSettementHotelResponseBody struct {
	// 结果msg
	ResultMsg *string `json:"resultMsg,omitempty" xml:"resultMsg,omitempty"`
	// module
	Module *BillSettementHotelResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 结果code
	ResultCode *int64 `json:"resultCode,omitempty" xml:"resultCode,omitempty"`
}

func (s BillSettementHotelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BillSettementHotelResponseBody) GoString() string {
	return s.String()
}

func (s *BillSettementHotelResponseBody) SetResultMsg(v string) *BillSettementHotelResponseBody {
	s.ResultMsg = &v
	return s
}

func (s *BillSettementHotelResponseBody) SetModule(v *BillSettementHotelResponseBodyModule) *BillSettementHotelResponseBody {
	s.Module = v
	return s
}

func (s *BillSettementHotelResponseBody) SetSuccess(v bool) *BillSettementHotelResponseBody {
	s.Success = &v
	return s
}

func (s *BillSettementHotelResponseBody) SetResultCode(v int64) *BillSettementHotelResponseBody {
	s.ResultCode = &v
	return s
}

type BillSettementHotelResponseBodyModule struct {
	// 类目
	Category *int64 `json:"category,omitempty" xml:"category,omitempty"`
	// 企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 数据集合
	DataList []*BillSettementHotelResponseBodyModuleDataList `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 记账更新结束日期
	PeriodEnd *string `json:"periodEnd,omitempty" xml:"periodEnd,omitempty"`
	// 记账更新开始日期
	PeriodStart *string `json:"periodStart,omitempty" xml:"periodStart,omitempty"`
	// 总数据量
	TotalNum *int64 `json:"totalNum,omitempty" xml:"totalNum,omitempty"`
}

func (s BillSettementHotelResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s BillSettementHotelResponseBodyModule) GoString() string {
	return s.String()
}

func (s *BillSettementHotelResponseBodyModule) SetCategory(v int64) *BillSettementHotelResponseBodyModule {
	s.Category = &v
	return s
}

func (s *BillSettementHotelResponseBodyModule) SetCorpId(v string) *BillSettementHotelResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *BillSettementHotelResponseBodyModule) SetDataList(v []*BillSettementHotelResponseBodyModuleDataList) *BillSettementHotelResponseBodyModule {
	s.DataList = v
	return s
}

func (s *BillSettementHotelResponseBodyModule) SetPeriodEnd(v string) *BillSettementHotelResponseBodyModule {
	s.PeriodEnd = &v
	return s
}

func (s *BillSettementHotelResponseBodyModule) SetPeriodStart(v string) *BillSettementHotelResponseBodyModule {
	s.PeriodStart = &v
	return s
}

func (s *BillSettementHotelResponseBodyModule) SetTotalNum(v int64) *BillSettementHotelResponseBodyModule {
	s.TotalNum = &v
	return s
}

type BillSettementHotelResponseBodyModuleDataList struct {
	// 交易流水号
	AlipayTradeNo *string `json:"alipayTradeNo,omitempty" xml:"alipayTradeNo,omitempty"`
	// 审批单号
	ApplyId *string `json:"applyId,omitempty" xml:"applyId,omitempty"`
	// 预定时间
	BookTime *string `json:"bookTime,omitempty" xml:"bookTime,omitempty"`
	// 预定人use id
	BookerId *string `json:"bookerId,omitempty" xml:"bookerId,omitempty"`
	// 预订人名称
	BookerName *string `json:"bookerName,omitempty" xml:"bookerName,omitempty"`
	// 资金方向
	CapitalDirection *string `json:"capitalDirection,omitempty" xml:"capitalDirection,omitempty"`
	// 级联部门
	CascadeDepartment *string `json:"cascadeDepartment,omitempty" xml:"cascadeDepartment,omitempty"`
	// 入住时间
	CheckInDate *string `json:"checkInDate,omitempty" xml:"checkInDate,omitempty"`
	// 离店时间
	CheckoutDate *string `json:"checkoutDate,omitempty" xml:"checkoutDate,omitempty"`
	// 入住城市
	City *string `json:"city,omitempty" xml:"city,omitempty"`
	// 城市编码
	CityCode *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	// 企业退款金额
	CorpRefundFee *float64 `json:"corpRefundFee,omitempty" xml:"corpRefundFee,omitempty"`
	// 企业支付金额
	CorpTotalFee *float64 `json:"corpTotalFee,omitempty" xml:"corpTotalFee,omitempty"`
	// 成本中心名称
	CostCenter *string `json:"costCenter,omitempty" xml:"costCenter,omitempty"`
	// 成本中心编码
	CostCenterNumber *string `json:"costCenterNumber,omitempty" xml:"costCenterNumber,omitempty"`
	// 末级部门
	Department *string `json:"department,omitempty" xml:"department,omitempty"`
	// 部门id
	DepartmentId *string `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 费用类型
	FeeType *string `json:"feeType,omitempty" xml:"feeType,omitempty"`
	// 杂费
	Fees *float64 `json:"fees,omitempty" xml:"fees,omitempty"`
	// 福豆支付
	FuPointFee *float64 `json:"fuPointFee,omitempty" xml:"fuPointFee,omitempty"`
	// 酒店名称
	HotelName *string `json:"hotelName,omitempty" xml:"hotelName,omitempty"`
	// 序号
	Index *string `json:"index,omitempty" xml:"index,omitempty"`
	// 发票抬头
	InvoiceTitle *string `json:"invoiceTitle,omitempty" xml:"invoiceTitle,omitempty"`
	// 是否协议价
	IsNegotiation *bool `json:"isNegotiation,omitempty" xml:"isNegotiation,omitempty"`
	// 是否合住
	IsShareStr *string `json:"isShareStr,omitempty" xml:"isShareStr,omitempty"`
	// 入住天数
	Nights *int64 `json:"nights,omitempty" xml:"nights,omitempty"`
	// 订单号
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// 订单金额
	OrderPrice *float64 `json:"orderPrice,omitempty" xml:"orderPrice,omitempty"`
	// 订单类型
	OrderType *string `json:"orderType,omitempty" xml:"orderType,omitempty"`
	// 超标审批单号
	OverApplyId *string `json:"overApplyId,omitempty" xml:"overApplyId,omitempty"`
	// 个人退款金额
	PersonRefundFee *float64 `json:"personRefundFee,omitempty" xml:"personRefundFee,omitempty"`
	// 个人支付金额
	PersonSettlePrice *float64 `json:"personSettlePrice,omitempty" xml:"personSettlePrice,omitempty"`
	// 主键id
	PrimaryId *int64 `json:"primaryId,omitempty" xml:"primaryId,omitempty"`
	// 项目编码
	ProjectCode *string `json:"projectCode,omitempty" xml:"projectCode,omitempty"`
	// 项目名称
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// 优惠券
	PromotionFee *float64 `json:"promotionFee,omitempty" xml:"promotionFee,omitempty"`
	// 房间数
	RoomNumber *int64 `json:"roomNumber,omitempty" xml:"roomNumber,omitempty"`
	// 房价
	RoomPrice *float64 `json:"roomPrice,omitempty" xml:"roomPrice,omitempty"`
	// 房间类型
	RoomType *string `json:"roomType,omitempty" xml:"roomType,omitempty"`
	// 服务费,仅在 feeType 20111、20112中展示
	ServiceFee *float64 `json:"serviceFee,omitempty" xml:"serviceFee,omitempty"`
	// 结算金额
	SettlementFee *int64 `json:"settlementFee,omitempty" xml:"settlementFee,omitempty"`
	// 结算时间
	SettlementTime *string `json:"settlementTime,omitempty" xml:"settlementTime,omitempty"`
	// 结算类型
	SettlementType *string `json:"settlementType,omitempty" xml:"settlementType,omitempty"`
	// 入账状态
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 总间夜数
	TotalNights *int64 `json:"totalNights,omitempty" xml:"totalNights,omitempty"`
	// 出行人use id
	TravelerId *string `json:"travelerId,omitempty" xml:"travelerId,omitempty"`
	// 出行人名称
	TravelerName *string `json:"travelerName,omitempty" xml:"travelerName,omitempty"`
	// 预订人工号
	BookerJobNo *string `json:"bookerJobNo,omitempty" xml:"bookerJobNo,omitempty"`
	// 出行人工号
	TravelerJobNo *string `json:"travelerJobNo,omitempty" xml:"travelerJobNo,omitempty"`
}

func (s BillSettementHotelResponseBodyModuleDataList) String() string {
	return tea.Prettify(s)
}

func (s BillSettementHotelResponseBodyModuleDataList) GoString() string {
	return s.String()
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetAlipayTradeNo(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.AlipayTradeNo = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetApplyId(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.ApplyId = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetBookTime(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.BookTime = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetBookerId(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.BookerId = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetBookerName(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.BookerName = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetCapitalDirection(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.CapitalDirection = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetCascadeDepartment(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.CascadeDepartment = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetCheckInDate(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.CheckInDate = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetCheckoutDate(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.CheckoutDate = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetCity(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.City = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetCityCode(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.CityCode = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetCorpRefundFee(v float64) *BillSettementHotelResponseBodyModuleDataList {
	s.CorpRefundFee = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetCorpTotalFee(v float64) *BillSettementHotelResponseBodyModuleDataList {
	s.CorpTotalFee = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetCostCenter(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.CostCenter = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetCostCenterNumber(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.CostCenterNumber = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetDepartment(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.Department = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetDepartmentId(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.DepartmentId = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetFeeType(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.FeeType = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetFees(v float64) *BillSettementHotelResponseBodyModuleDataList {
	s.Fees = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetFuPointFee(v float64) *BillSettementHotelResponseBodyModuleDataList {
	s.FuPointFee = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetHotelName(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.HotelName = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetIndex(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.Index = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetInvoiceTitle(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.InvoiceTitle = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetIsNegotiation(v bool) *BillSettementHotelResponseBodyModuleDataList {
	s.IsNegotiation = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetIsShareStr(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.IsShareStr = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetNights(v int64) *BillSettementHotelResponseBodyModuleDataList {
	s.Nights = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetOrderId(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.OrderId = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetOrderPrice(v float64) *BillSettementHotelResponseBodyModuleDataList {
	s.OrderPrice = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetOrderType(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.OrderType = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetOverApplyId(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.OverApplyId = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetPersonRefundFee(v float64) *BillSettementHotelResponseBodyModuleDataList {
	s.PersonRefundFee = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetPersonSettlePrice(v float64) *BillSettementHotelResponseBodyModuleDataList {
	s.PersonSettlePrice = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetPrimaryId(v int64) *BillSettementHotelResponseBodyModuleDataList {
	s.PrimaryId = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetProjectCode(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.ProjectCode = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetProjectName(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.ProjectName = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetPromotionFee(v float64) *BillSettementHotelResponseBodyModuleDataList {
	s.PromotionFee = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetRoomNumber(v int64) *BillSettementHotelResponseBodyModuleDataList {
	s.RoomNumber = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetRoomPrice(v float64) *BillSettementHotelResponseBodyModuleDataList {
	s.RoomPrice = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetRoomType(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.RoomType = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetServiceFee(v float64) *BillSettementHotelResponseBodyModuleDataList {
	s.ServiceFee = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetSettlementFee(v int64) *BillSettementHotelResponseBodyModuleDataList {
	s.SettlementFee = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetSettlementTime(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.SettlementTime = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetSettlementType(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.SettlementType = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetStatus(v int64) *BillSettementHotelResponseBodyModuleDataList {
	s.Status = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetTotalNights(v int64) *BillSettementHotelResponseBodyModuleDataList {
	s.TotalNights = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetTravelerId(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.TravelerId = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetTravelerName(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.TravelerName = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetBookerJobNo(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.BookerJobNo = &v
	return s
}

func (s *BillSettementHotelResponseBodyModuleDataList) SetTravelerJobNo(v string) *BillSettementHotelResponseBodyModuleDataList {
	s.TravelerJobNo = &v
	return s
}

type BillSettementHotelResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BillSettementHotelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BillSettementHotelResponse) String() string {
	return tea.Prettify(s)
}

func (s BillSettementHotelResponse) GoString() string {
	return s.String()
}

func (s *BillSettementHotelResponse) SetHeaders(v map[string]*string) *BillSettementHotelResponse {
	s.Headers = v
	return s
}

func (s *BillSettementHotelResponse) SetBody(v *BillSettementHotelResponseBody) *BillSettementHotelResponse {
	s.Body = v
	return s
}

type GetFlightExceedApplyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFlightExceedApplyHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFlightExceedApplyHeaders) GoString() string {
	return s.String()
}

func (s *GetFlightExceedApplyHeaders) SetCommonHeaders(v map[string]*string) *GetFlightExceedApplyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFlightExceedApplyHeaders) SetXAcsDingtalkAccessToken(v string) *GetFlightExceedApplyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFlightExceedApplyRequest struct {
	// 第三方企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 商旅超标审批单id
	ApplyId *string `json:"applyId,omitempty" xml:"applyId,omitempty"`
}

func (s GetFlightExceedApplyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFlightExceedApplyRequest) GoString() string {
	return s.String()
}

func (s *GetFlightExceedApplyRequest) SetCorpId(v string) *GetFlightExceedApplyRequest {
	s.CorpId = &v
	return s
}

func (s *GetFlightExceedApplyRequest) SetApplyId(v string) *GetFlightExceedApplyRequest {
	s.ApplyId = &v
	return s
}

type GetFlightExceedApplyResponseBody struct {
	// 第三方企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 商旅超标审批单id
	ApplyId *int64 `json:"applyId,omitempty" xml:"applyId,omitempty"`
	// 审批单状态 0:审批中 1:已同意 2:已拒绝
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// 出差原因
	BtripCause *string `json:"btripCause,omitempty" xml:"btripCause,omitempty"`
	// 超标类型，1:折扣 2,8,10:时间 3,9,11:折扣和时间
	ExceedType *int32 `json:"exceedType,omitempty" xml:"exceedType,omitempty"`
	// 超标原因
	ExceedReason *string `json:"exceedReason,omitempty" xml:"exceedReason,omitempty"`
	// 原差旅标准
	OriginStandard *string `json:"originStandard,omitempty" xml:"originStandard,omitempty"`
	// 审批单提交时间
	SubmitTime *string `json:"submitTime,omitempty" xml:"submitTime,omitempty"`
	// 第三方用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 意向出行信息
	ApplyIntentionInfoDO *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO `json:"applyIntentionInfoDO,omitempty" xml:"applyIntentionInfoDO,omitempty" type:"Struct"`
	// 第三方出差审批单号
	ThirdpartApplyId *string `json:"thirdpartApplyId,omitempty" xml:"thirdpartApplyId,omitempty"`
}

func (s GetFlightExceedApplyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFlightExceedApplyResponseBody) GoString() string {
	return s.String()
}

func (s *GetFlightExceedApplyResponseBody) SetCorpId(v string) *GetFlightExceedApplyResponseBody {
	s.CorpId = &v
	return s
}

func (s *GetFlightExceedApplyResponseBody) SetApplyId(v int64) *GetFlightExceedApplyResponseBody {
	s.ApplyId = &v
	return s
}

func (s *GetFlightExceedApplyResponseBody) SetStatus(v int32) *GetFlightExceedApplyResponseBody {
	s.Status = &v
	return s
}

func (s *GetFlightExceedApplyResponseBody) SetBtripCause(v string) *GetFlightExceedApplyResponseBody {
	s.BtripCause = &v
	return s
}

func (s *GetFlightExceedApplyResponseBody) SetExceedType(v int32) *GetFlightExceedApplyResponseBody {
	s.ExceedType = &v
	return s
}

func (s *GetFlightExceedApplyResponseBody) SetExceedReason(v string) *GetFlightExceedApplyResponseBody {
	s.ExceedReason = &v
	return s
}

func (s *GetFlightExceedApplyResponseBody) SetOriginStandard(v string) *GetFlightExceedApplyResponseBody {
	s.OriginStandard = &v
	return s
}

func (s *GetFlightExceedApplyResponseBody) SetSubmitTime(v string) *GetFlightExceedApplyResponseBody {
	s.SubmitTime = &v
	return s
}

func (s *GetFlightExceedApplyResponseBody) SetUserId(v string) *GetFlightExceedApplyResponseBody {
	s.UserId = &v
	return s
}

func (s *GetFlightExceedApplyResponseBody) SetApplyIntentionInfoDO(v *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO) *GetFlightExceedApplyResponseBody {
	s.ApplyIntentionInfoDO = v
	return s
}

func (s *GetFlightExceedApplyResponseBody) SetThirdpartApplyId(v string) *GetFlightExceedApplyResponseBody {
	s.ThirdpartApplyId = &v
	return s
}

type GetFlightExceedApplyResponseBodyApplyIntentionInfoDO struct {
	// 到达城市三字码
	ArrCity *string `json:"arrCity,omitempty" xml:"arrCity,omitempty"`
	// 到达城市名称
	ArrCityName *string `json:"arrCityName,omitempty" xml:"arrCityName,omitempty"`
	// 到达时间
	ArrTime *string `json:"arrTime,omitempty" xml:"arrTime,omitempty"`
	// 超标的舱位，F：头等舱 C：商务舱 Y：经济舱 P：超值经济舱
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// 申请超标的舱等 0：头等舱 1：商务舱 2：经济舱 3：超值经济舱
	CabinClass *int32 `json:"cabinClass,omitempty" xml:"cabinClass,omitempty"`
	// 舱等描述，头等舱，商务舱，经济舱，超值经济舱
	CabinClassStr *string `json:"cabinClassStr,omitempty" xml:"cabinClassStr,omitempty"`
	// 出发城市三字码
	DepCity *string `json:"depCity,omitempty" xml:"depCity,omitempty"`
	// 出发城市名称
	DepCityName *string `json:"depCityName,omitempty" xml:"depCityName,omitempty"`
	// 出发时间
	DepTime *string `json:"depTime,omitempty" xml:"depTime,omitempty"`
	// 折扣
	Discount *float64 `json:"discount,omitempty" xml:"discount,omitempty"`
	// 航班号
	FlightNo *string `json:"flightNo,omitempty" xml:"flightNo,omitempty"`
	// 意向航班价格（元）
	Price *int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 超标类型，1:折扣 2,8,10:时间 3,9,11:折扣和时间
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetFlightExceedApplyResponseBodyApplyIntentionInfoDO) String() string {
	return tea.Prettify(s)
}

func (s GetFlightExceedApplyResponseBodyApplyIntentionInfoDO) GoString() string {
	return s.String()
}

func (s *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO) SetArrCity(v string) *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO {
	s.ArrCity = &v
	return s
}

func (s *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO) SetArrCityName(v string) *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO {
	s.ArrCityName = &v
	return s
}

func (s *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO) SetArrTime(v string) *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO {
	s.ArrTime = &v
	return s
}

func (s *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO) SetCabin(v string) *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO {
	s.Cabin = &v
	return s
}

func (s *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO) SetCabinClass(v int32) *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO {
	s.CabinClass = &v
	return s
}

func (s *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO) SetCabinClassStr(v string) *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO {
	s.CabinClassStr = &v
	return s
}

func (s *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO) SetDepCity(v string) *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO {
	s.DepCity = &v
	return s
}

func (s *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO) SetDepCityName(v string) *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO {
	s.DepCityName = &v
	return s
}

func (s *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO) SetDepTime(v string) *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO {
	s.DepTime = &v
	return s
}

func (s *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO) SetDiscount(v float64) *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO {
	s.Discount = &v
	return s
}

func (s *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO) SetFlightNo(v string) *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO {
	s.FlightNo = &v
	return s
}

func (s *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO) SetPrice(v int64) *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO {
	s.Price = &v
	return s
}

func (s *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO) SetType(v int32) *GetFlightExceedApplyResponseBodyApplyIntentionInfoDO {
	s.Type = &v
	return s
}

type GetFlightExceedApplyResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFlightExceedApplyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFlightExceedApplyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFlightExceedApplyResponse) GoString() string {
	return s.String()
}

func (s *GetFlightExceedApplyResponse) SetHeaders(v map[string]*string) *GetFlightExceedApplyResponse {
	s.Headers = v
	return s
}

func (s *GetFlightExceedApplyResponse) SetBody(v *GetFlightExceedApplyResponseBody) *GetFlightExceedApplyResponse {
	s.Body = v
	return s
}

type BillSettementCarHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BillSettementCarHeaders) String() string {
	return tea.Prettify(s)
}

func (s BillSettementCarHeaders) GoString() string {
	return s.String()
}

func (s *BillSettementCarHeaders) SetCommonHeaders(v map[string]*string) *BillSettementCarHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BillSettementCarHeaders) SetXAcsDingtalkAccessToken(v string) *BillSettementCarHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BillSettementCarRequest struct {
	CorpId      *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Category    *int64  `json:"category,omitempty" xml:"category,omitempty"`
	PageSize    *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	PeriodStart *string `json:"periodStart,omitempty" xml:"periodStart,omitempty"`
	PeriodEnd   *string `json:"periodEnd,omitempty" xml:"periodEnd,omitempty"`
	PageNumber  *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
}

func (s BillSettementCarRequest) String() string {
	return tea.Prettify(s)
}

func (s BillSettementCarRequest) GoString() string {
	return s.String()
}

func (s *BillSettementCarRequest) SetCorpId(v string) *BillSettementCarRequest {
	s.CorpId = &v
	return s
}

func (s *BillSettementCarRequest) SetCategory(v int64) *BillSettementCarRequest {
	s.Category = &v
	return s
}

func (s *BillSettementCarRequest) SetPageSize(v int64) *BillSettementCarRequest {
	s.PageSize = &v
	return s
}

func (s *BillSettementCarRequest) SetPeriodStart(v string) *BillSettementCarRequest {
	s.PeriodStart = &v
	return s
}

func (s *BillSettementCarRequest) SetPeriodEnd(v string) *BillSettementCarRequest {
	s.PeriodEnd = &v
	return s
}

func (s *BillSettementCarRequest) SetPageNumber(v int64) *BillSettementCarRequest {
	s.PageNumber = &v
	return s
}

type BillSettementCarResponseBody struct {
	// 结果msg
	ResultMsg *string `json:"resultMsg,omitempty" xml:"resultMsg,omitempty"`
	// module
	Module *BillSettementCarResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 结果code
	ResultCode *int64 `json:"resultCode,omitempty" xml:"resultCode,omitempty"`
}

func (s BillSettementCarResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BillSettementCarResponseBody) GoString() string {
	return s.String()
}

func (s *BillSettementCarResponseBody) SetResultMsg(v string) *BillSettementCarResponseBody {
	s.ResultMsg = &v
	return s
}

func (s *BillSettementCarResponseBody) SetModule(v *BillSettementCarResponseBodyModule) *BillSettementCarResponseBody {
	s.Module = v
	return s
}

func (s *BillSettementCarResponseBody) SetSuccess(v bool) *BillSettementCarResponseBody {
	s.Success = &v
	return s
}

func (s *BillSettementCarResponseBody) SetResultCode(v int64) *BillSettementCarResponseBody {
	s.ResultCode = &v
	return s
}

type BillSettementCarResponseBodyModule struct {
	// 类目
	Category *int64 `json:"category,omitempty" xml:"category,omitempty"`
	// 企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 数据集合
	DataList []*BillSettementCarResponseBodyModuleDataList `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 记账更新开始日期
	PeriodEnd *string `json:"periodEnd,omitempty" xml:"periodEnd,omitempty"`
	// 记账更新结束日期
	PeriodStart *string `json:"periodStart,omitempty" xml:"periodStart,omitempty"`
	// 总数量
	TotalNum *int64 `json:"totalNum,omitempty" xml:"totalNum,omitempty"`
}

func (s BillSettementCarResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s BillSettementCarResponseBodyModule) GoString() string {
	return s.String()
}

func (s *BillSettementCarResponseBodyModule) SetCategory(v int64) *BillSettementCarResponseBodyModule {
	s.Category = &v
	return s
}

func (s *BillSettementCarResponseBodyModule) SetCorpId(v string) *BillSettementCarResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *BillSettementCarResponseBodyModule) SetDataList(v []*BillSettementCarResponseBodyModuleDataList) *BillSettementCarResponseBodyModule {
	s.DataList = v
	return s
}

func (s *BillSettementCarResponseBodyModule) SetPeriodEnd(v string) *BillSettementCarResponseBodyModule {
	s.PeriodEnd = &v
	return s
}

func (s *BillSettementCarResponseBodyModule) SetPeriodStart(v string) *BillSettementCarResponseBodyModule {
	s.PeriodStart = &v
	return s
}

func (s *BillSettementCarResponseBodyModule) SetTotalNum(v int64) *BillSettementCarResponseBodyModule {
	s.TotalNum = &v
	return s
}

type BillSettementCarResponseBodyModuleDataList struct {
	// 支付交易流水号
	AlipayTradeNo *string `json:"alipayTradeNo,omitempty" xml:"alipayTradeNo,omitempty"`
	// 审批单号
	ApplyId *string `json:"applyId,omitempty" xml:"applyId,omitempty"`
	// 到达城市
	ArrCity *string `json:"arrCity,omitempty" xml:"arrCity,omitempty"`
	// 到达日期
	ArrDate *string `json:"arrDate,omitempty" xml:"arrDate,omitempty"`
	// 到达地
	ArrLocation *string `json:"arrLocation,omitempty" xml:"arrLocation,omitempty"`
	// 到达时间
	ArrTime *string `json:"arrTime,omitempty" xml:"arrTime,omitempty"`
	// 预定时间
	BookTime *string `json:"bookTime,omitempty" xml:"bookTime,omitempty"`
	// 预定人use id
	BookerId *string `json:"bookerId,omitempty" xml:"bookerId,omitempty"`
	// 预订人名称
	BookerName *string `json:"bookerName,omitempty" xml:"bookerName,omitempty"`
	// 用车事由
	BusinessCategory *string `json:"businessCategory,omitempty" xml:"businessCategory,omitempty"`
	// 资金方向
	CapitalDirection *string `json:"capitalDirection,omitempty" xml:"capitalDirection,omitempty"`
	// 车型
	CarLevel *string `json:"carLevel,omitempty" xml:"carLevel,omitempty"`
	// 级联部门
	CascadeDepartment *string `json:"cascadeDepartment,omitempty" xml:"cascadeDepartment,omitempty"`
	// 成本中心名称
	CostCenter *string `json:"costCenter,omitempty" xml:"costCenter,omitempty"`
	// 成本中心编号
	CostCenterNumber *string `json:"costCenterNumber,omitempty" xml:"costCenterNumber,omitempty"`
	// 优惠券
	Coupon *float64 `json:"coupon,omitempty" xml:"coupon,omitempty"`
	// 优惠金额
	CouponPrice *float64 `json:"couponPrice,omitempty" xml:"couponPrice,omitempty"`
	// 末级部门
	Department *string `json:"department,omitempty" xml:"department,omitempty"`
	// 部门id
	DepartmentId *string `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 出发城市
	DeptCity *string `json:"deptCity,omitempty" xml:"deptCity,omitempty"`
	// 出发日期
	DeptDate *string `json:"deptDate,omitempty" xml:"deptDate,omitempty"`
	// 出发地
	DeptLocation *string `json:"deptLocation,omitempty" xml:"deptLocation,omitempty"`
	// 出发时间
	DeptTime *string `json:"deptTime,omitempty" xml:"deptTime,omitempty"`
	// 预估行驶距离
	EstimateDriveDistance *string `json:"estimateDriveDistance,omitempty" xml:"estimateDriveDistance,omitempty"`
	// 预估金额
	EstimatePrice *float64 `json:"estimatePrice,omitempty" xml:"estimatePrice,omitempty"`
	// 费用类型
	FeeType *string `json:"feeType,omitempty" xml:"feeType,omitempty"`
	// 序号
	Index *string `json:"index,omitempty" xml:"index,omitempty"`
	// 发票抬头
	InvoiceTitle *string `json:"invoiceTitle,omitempty" xml:"invoiceTitle,omitempty"`
	// 用车事由
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 订单id
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// 订单金额
	OrderPrice *float64 `json:"orderPrice,omitempty" xml:"orderPrice,omitempty"`
	// 超标审批单号
	OverApplyId *string `json:"overApplyId,omitempty" xml:"overApplyId,omitempty"`
	// 个人支付金额
	PersonSettleFee *float64 `json:"personSettleFee,omitempty" xml:"personSettleFee,omitempty"`
	PrimaryId       *string  `json:"primaryId,omitempty" xml:"primaryId,omitempty"`
	// 项目编码
	ProjectCode *string `json:"projectCode,omitempty" xml:"projectCode,omitempty"`
	// 项目名称
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// 供应商
	ProviderName *string `json:"providerName,omitempty" xml:"providerName,omitempty"`
	// 实际行驶距离
	RealDriveDistance *string `json:"realDriveDistance,omitempty" xml:"realDriveDistance,omitempty"`
	// 实际上车点
	RealFromAddr *string `json:"realFromAddr,omitempty" xml:"realFromAddr,omitempty"`
	// 实际下车点
	RealToAddr *string `json:"realToAddr,omitempty" xml:"realToAddr,omitempty"`
	// 服务费，仅在feeType 40111 中展示
	ServiceFee *string `json:"serviceFee,omitempty" xml:"serviceFee,omitempty"`
	// 结算金额
	SettlementFee *float64 `json:"settlementFee,omitempty" xml:"settlementFee,omitempty"`
	// 结算时间
	SettlementTime *string `json:"settlementTime,omitempty" xml:"settlementTime,omitempty"`
	// 结算类型
	SettlementType *string `json:"settlementType,omitempty" xml:"settlementType,omitempty"`
	// 特别关注订单
	SpecialOrder *string `json:"specialOrder,omitempty" xml:"specialOrder,omitempty"`
	// 特别关注原因
	SpecialReason *string `json:"specialReason,omitempty" xml:"specialReason,omitempty"`
	// 入账状态
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 出行人use id
	TravelerId *string `json:"travelerId,omitempty" xml:"travelerId,omitempty"`
	// 出行人名称
	TravelerName *string `json:"travelerName,omitempty" xml:"travelerName,omitempty"`
	// 员工是否认可
	UserConfirmDesc *string `json:"userConfirmDesc,omitempty" xml:"userConfirmDesc,omitempty"`
	// 预订人工号
	BookerJobNo *string `json:"bookerJobNo,omitempty" xml:"bookerJobNo,omitempty"`
	// 出行人工号
	TravelerJobNo *string `json:"travelerJobNo,omitempty" xml:"travelerJobNo,omitempty"`
}

func (s BillSettementCarResponseBodyModuleDataList) String() string {
	return tea.Prettify(s)
}

func (s BillSettementCarResponseBodyModuleDataList) GoString() string {
	return s.String()
}

func (s *BillSettementCarResponseBodyModuleDataList) SetAlipayTradeNo(v string) *BillSettementCarResponseBodyModuleDataList {
	s.AlipayTradeNo = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetApplyId(v string) *BillSettementCarResponseBodyModuleDataList {
	s.ApplyId = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetArrCity(v string) *BillSettementCarResponseBodyModuleDataList {
	s.ArrCity = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetArrDate(v string) *BillSettementCarResponseBodyModuleDataList {
	s.ArrDate = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetArrLocation(v string) *BillSettementCarResponseBodyModuleDataList {
	s.ArrLocation = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetArrTime(v string) *BillSettementCarResponseBodyModuleDataList {
	s.ArrTime = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetBookTime(v string) *BillSettementCarResponseBodyModuleDataList {
	s.BookTime = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetBookerId(v string) *BillSettementCarResponseBodyModuleDataList {
	s.BookerId = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetBookerName(v string) *BillSettementCarResponseBodyModuleDataList {
	s.BookerName = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetBusinessCategory(v string) *BillSettementCarResponseBodyModuleDataList {
	s.BusinessCategory = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetCapitalDirection(v string) *BillSettementCarResponseBodyModuleDataList {
	s.CapitalDirection = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetCarLevel(v string) *BillSettementCarResponseBodyModuleDataList {
	s.CarLevel = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetCascadeDepartment(v string) *BillSettementCarResponseBodyModuleDataList {
	s.CascadeDepartment = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetCostCenter(v string) *BillSettementCarResponseBodyModuleDataList {
	s.CostCenter = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetCostCenterNumber(v string) *BillSettementCarResponseBodyModuleDataList {
	s.CostCenterNumber = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetCoupon(v float64) *BillSettementCarResponseBodyModuleDataList {
	s.Coupon = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetCouponPrice(v float64) *BillSettementCarResponseBodyModuleDataList {
	s.CouponPrice = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetDepartment(v string) *BillSettementCarResponseBodyModuleDataList {
	s.Department = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetDepartmentId(v string) *BillSettementCarResponseBodyModuleDataList {
	s.DepartmentId = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetDeptCity(v string) *BillSettementCarResponseBodyModuleDataList {
	s.DeptCity = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetDeptDate(v string) *BillSettementCarResponseBodyModuleDataList {
	s.DeptDate = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetDeptLocation(v string) *BillSettementCarResponseBodyModuleDataList {
	s.DeptLocation = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetDeptTime(v string) *BillSettementCarResponseBodyModuleDataList {
	s.DeptTime = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetEstimateDriveDistance(v string) *BillSettementCarResponseBodyModuleDataList {
	s.EstimateDriveDistance = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetEstimatePrice(v float64) *BillSettementCarResponseBodyModuleDataList {
	s.EstimatePrice = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetFeeType(v string) *BillSettementCarResponseBodyModuleDataList {
	s.FeeType = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetIndex(v string) *BillSettementCarResponseBodyModuleDataList {
	s.Index = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetInvoiceTitle(v string) *BillSettementCarResponseBodyModuleDataList {
	s.InvoiceTitle = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetMemo(v string) *BillSettementCarResponseBodyModuleDataList {
	s.Memo = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetOrderId(v string) *BillSettementCarResponseBodyModuleDataList {
	s.OrderId = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetOrderPrice(v float64) *BillSettementCarResponseBodyModuleDataList {
	s.OrderPrice = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetOverApplyId(v string) *BillSettementCarResponseBodyModuleDataList {
	s.OverApplyId = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetPersonSettleFee(v float64) *BillSettementCarResponseBodyModuleDataList {
	s.PersonSettleFee = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetPrimaryId(v string) *BillSettementCarResponseBodyModuleDataList {
	s.PrimaryId = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetProjectCode(v string) *BillSettementCarResponseBodyModuleDataList {
	s.ProjectCode = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetProjectName(v string) *BillSettementCarResponseBodyModuleDataList {
	s.ProjectName = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetProviderName(v string) *BillSettementCarResponseBodyModuleDataList {
	s.ProviderName = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetRealDriveDistance(v string) *BillSettementCarResponseBodyModuleDataList {
	s.RealDriveDistance = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetRealFromAddr(v string) *BillSettementCarResponseBodyModuleDataList {
	s.RealFromAddr = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetRealToAddr(v string) *BillSettementCarResponseBodyModuleDataList {
	s.RealToAddr = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetServiceFee(v string) *BillSettementCarResponseBodyModuleDataList {
	s.ServiceFee = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetSettlementFee(v float64) *BillSettementCarResponseBodyModuleDataList {
	s.SettlementFee = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetSettlementTime(v string) *BillSettementCarResponseBodyModuleDataList {
	s.SettlementTime = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetSettlementType(v string) *BillSettementCarResponseBodyModuleDataList {
	s.SettlementType = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetSpecialOrder(v string) *BillSettementCarResponseBodyModuleDataList {
	s.SpecialOrder = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetSpecialReason(v string) *BillSettementCarResponseBodyModuleDataList {
	s.SpecialReason = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetStatus(v int64) *BillSettementCarResponseBodyModuleDataList {
	s.Status = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetTravelerId(v string) *BillSettementCarResponseBodyModuleDataList {
	s.TravelerId = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetTravelerName(v string) *BillSettementCarResponseBodyModuleDataList {
	s.TravelerName = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetUserConfirmDesc(v string) *BillSettementCarResponseBodyModuleDataList {
	s.UserConfirmDesc = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetBookerJobNo(v string) *BillSettementCarResponseBodyModuleDataList {
	s.BookerJobNo = &v
	return s
}

func (s *BillSettementCarResponseBodyModuleDataList) SetTravelerJobNo(v string) *BillSettementCarResponseBodyModuleDataList {
	s.TravelerJobNo = &v
	return s
}

type BillSettementCarResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BillSettementCarResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BillSettementCarResponse) String() string {
	return tea.Prettify(s)
}

func (s BillSettementCarResponse) GoString() string {
	return s.String()
}

func (s *BillSettementCarResponse) SetHeaders(v map[string]*string) *BillSettementCarResponse {
	s.Headers = v
	return s
}

func (s *BillSettementCarResponse) SetBody(v *BillSettementCarResponseBody) *BillSettementCarResponse {
	s.Body = v
	return s
}

type BillSettementBtripTrainHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BillSettementBtripTrainHeaders) String() string {
	return tea.Prettify(s)
}

func (s BillSettementBtripTrainHeaders) GoString() string {
	return s.String()
}

func (s *BillSettementBtripTrainHeaders) SetCommonHeaders(v map[string]*string) *BillSettementBtripTrainHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BillSettementBtripTrainHeaders) SetXAcsDingtalkAccessToken(v string) *BillSettementBtripTrainHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BillSettementBtripTrainRequest struct {
	CorpId      *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Category    *int64  `json:"category,omitempty" xml:"category,omitempty"`
	PageSize    *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	PeriodStart *string `json:"periodStart,omitempty" xml:"periodStart,omitempty"`
	PageNumber  *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PeriodEnd   *string `json:"periodEnd,omitempty" xml:"periodEnd,omitempty"`
}

func (s BillSettementBtripTrainRequest) String() string {
	return tea.Prettify(s)
}

func (s BillSettementBtripTrainRequest) GoString() string {
	return s.String()
}

func (s *BillSettementBtripTrainRequest) SetCorpId(v string) *BillSettementBtripTrainRequest {
	s.CorpId = &v
	return s
}

func (s *BillSettementBtripTrainRequest) SetCategory(v int64) *BillSettementBtripTrainRequest {
	s.Category = &v
	return s
}

func (s *BillSettementBtripTrainRequest) SetPageSize(v int64) *BillSettementBtripTrainRequest {
	s.PageSize = &v
	return s
}

func (s *BillSettementBtripTrainRequest) SetPeriodStart(v string) *BillSettementBtripTrainRequest {
	s.PeriodStart = &v
	return s
}

func (s *BillSettementBtripTrainRequest) SetPageNumber(v int64) *BillSettementBtripTrainRequest {
	s.PageNumber = &v
	return s
}

func (s *BillSettementBtripTrainRequest) SetPeriodEnd(v string) *BillSettementBtripTrainRequest {
	s.PeriodEnd = &v
	return s
}

type BillSettementBtripTrainResponseBody struct {
	// 结果msg
	ResultMsg *string `json:"resultMsg,omitempty" xml:"resultMsg,omitempty"`
	// module
	Module *BillSettementBtripTrainResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 结果code
	ResultCode *int64 `json:"resultCode,omitempty" xml:"resultCode,omitempty"`
}

func (s BillSettementBtripTrainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BillSettementBtripTrainResponseBody) GoString() string {
	return s.String()
}

func (s *BillSettementBtripTrainResponseBody) SetResultMsg(v string) *BillSettementBtripTrainResponseBody {
	s.ResultMsg = &v
	return s
}

func (s *BillSettementBtripTrainResponseBody) SetModule(v *BillSettementBtripTrainResponseBodyModule) *BillSettementBtripTrainResponseBody {
	s.Module = v
	return s
}

func (s *BillSettementBtripTrainResponseBody) SetSuccess(v bool) *BillSettementBtripTrainResponseBody {
	s.Success = &v
	return s
}

func (s *BillSettementBtripTrainResponseBody) SetResultCode(v int64) *BillSettementBtripTrainResponseBody {
	s.ResultCode = &v
	return s
}

type BillSettementBtripTrainResponseBodyModule struct {
	// 类目
	Category *int64 `json:"category,omitempty" xml:"category,omitempty"`
	// 企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 数据集合
	DataList []*BillSettementBtripTrainResponseBodyModuleDataList `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 记账更新开始时间
	PeriodEnd *string `json:"periodEnd,omitempty" xml:"periodEnd,omitempty"`
	// 记账更新结束时间
	PeriodStart *string `json:"periodStart,omitempty" xml:"periodStart,omitempty"`
	// 总数据量
	TotalNum *int64 `json:"totalNum,omitempty" xml:"totalNum,omitempty"`
}

func (s BillSettementBtripTrainResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s BillSettementBtripTrainResponseBodyModule) GoString() string {
	return s.String()
}

func (s *BillSettementBtripTrainResponseBodyModule) SetCategory(v int64) *BillSettementBtripTrainResponseBodyModule {
	s.Category = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModule) SetCorpId(v string) *BillSettementBtripTrainResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModule) SetDataList(v []*BillSettementBtripTrainResponseBodyModuleDataList) *BillSettementBtripTrainResponseBodyModule {
	s.DataList = v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModule) SetPeriodEnd(v string) *BillSettementBtripTrainResponseBodyModule {
	s.PeriodEnd = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModule) SetPeriodStart(v string) *BillSettementBtripTrainResponseBodyModule {
	s.PeriodStart = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModule) SetTotalNum(v int64) *BillSettementBtripTrainResponseBodyModule {
	s.TotalNum = &v
	return s
}

type BillSettementBtripTrainResponseBodyModuleDataList struct {
	// 交易流水号
	AlipayTradeNo *string `json:"alipayTradeNo,omitempty" xml:"alipayTradeNo,omitempty"`
	// 审批单号
	ApplyId *string `json:"applyId,omitempty" xml:"applyId,omitempty"`
	// 到达日期
	ArrDate *string `json:"arrDate,omitempty" xml:"arrDate,omitempty"`
	// 到达站点
	ArrStation *string `json:"arrStation,omitempty" xml:"arrStation,omitempty"`
	// 到达时间
	ArrTime *string `json:"arrTime,omitempty" xml:"arrTime,omitempty"`
	// 预定时间
	BookTime *string `json:"bookTime,omitempty" xml:"bookTime,omitempty"`
	// 预定人use id
	BookerId *string `json:"bookerId,omitempty" xml:"bookerId,omitempty"`
	// 预订人名称
	BookerName *string `json:"bookerName,omitempty" xml:"bookerName,omitempty"`
	// 资金方向
	CapitalDirection *string `json:"capitalDirection,omitempty" xml:"capitalDirection,omitempty"`
	// 级联部门
	CascadeDepartment *string `json:"cascadeDepartment,omitempty" xml:"cascadeDepartment,omitempty"`
	// 改签手续费
	ChangeFee *float64 `json:"changeFee,omitempty" xml:"changeFee,omitempty"`
	// 成本中心名称
	CostCenter *string `json:"costCenter,omitempty" xml:"costCenter,omitempty"`
	// 成本中心编码
	CostCenterNumber *string `json:"costCenterNumber,omitempty" xml:"costCenterNumber,omitempty"`
	// 折扣率
	Coupon *float64 `json:"coupon,omitempty" xml:"coupon,omitempty"`
	// 末级部门
	Department *string `json:"department,omitempty" xml:"department,omitempty"`
	// 部门id
	DepartmentId *string `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 出发日期
	DeptDate *string `json:"deptDate,omitempty" xml:"deptDate,omitempty"`
	// 出发站
	DeptStation *string `json:"deptStation,omitempty" xml:"deptStation,omitempty"`
	// 出发时间
	DeptTime *string `json:"deptTime,omitempty" xml:"deptTime,omitempty"`
	// 费用类型
	FeeType *string `json:"feeType,omitempty" xml:"feeType,omitempty"`
	// 序号
	Index *string `json:"index,omitempty" xml:"index,omitempty"`
	// 发票抬头
	InvoiceTitle *string `json:"invoiceTitle,omitempty" xml:"invoiceTitle,omitempty"`
	// 订单号
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// 订单金额
	OrderPrice *float64 `json:"orderPrice,omitempty" xml:"orderPrice,omitempty"`
	// 超标审批单号
	OverApplyId *string `json:"overApplyId,omitempty" xml:"overApplyId,omitempty"`
	// 主键id
	PrimaryId *int64 `json:"primaryId,omitempty" xml:"primaryId,omitempty"`
	// 项目编号
	ProjectCode *string `json:"projectCode,omitempty" xml:"projectCode,omitempty"`
	// 项目名称
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// 退款手续费
	RefundFee *float64 `json:"refundFee,omitempty" xml:"refundFee,omitempty"`
	// 运行时长
	RunTime *string `json:"runTime,omitempty" xml:"runTime,omitempty"`
	// 座位号
	SeatNo *string `json:"seatNo,omitempty" xml:"seatNo,omitempty"`
	// 坐席
	SeatType *string `json:"seatType,omitempty" xml:"seatType,omitempty"`
	// 服务费，仅在feeType 6007、6008中展示
	ServiceFee *float64 `json:"serviceFee,omitempty" xml:"serviceFee,omitempty"`
	// 结算金额
	SettlementFee *float64 `json:"settlementFee,omitempty" xml:"settlementFee,omitempty"`
	// 结算时间
	SettlementTime *string `json:"settlementTime,omitempty" xml:"settlementTime,omitempty"`
	// 结算类型
	SettlementType *string `json:"settlementType,omitempty" xml:"settlementType,omitempty"`
	// 入账状态
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 票面票号
	TicketNo *string `json:"ticketNo,omitempty" xml:"ticketNo,omitempty"`
	// 票价
	TicketPrice *float64 `json:"ticketPrice,omitempty" xml:"ticketPrice,omitempty"`
	// 车次号
	TrainNo *string `json:"trainNo,omitempty" xml:"trainNo,omitempty"`
	// 车次类型
	TrainType *string `json:"trainType,omitempty" xml:"trainType,omitempty"`
	// 出行人useId
	TravelerId *string `json:"travelerId,omitempty" xml:"travelerId,omitempty"`
	// 出行人名称
	TravelerName *string `json:"travelerName,omitempty" xml:"travelerName,omitempty"`
	// 预订人工号
	BookerJobNo *string `json:"bookerJobNo,omitempty" xml:"bookerJobNo,omitempty"`
	// 出行人工号
	TravelerJobNo *string `json:"travelerJobNo,omitempty" xml:"travelerJobNo,omitempty"`
	// 发票类型
	VoucherType *int64 `json:"voucherType,omitempty" xml:"voucherType,omitempty"`
}

func (s BillSettementBtripTrainResponseBodyModuleDataList) String() string {
	return tea.Prettify(s)
}

func (s BillSettementBtripTrainResponseBodyModuleDataList) GoString() string {
	return s.String()
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetAlipayTradeNo(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.AlipayTradeNo = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetApplyId(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.ApplyId = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetArrDate(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.ArrDate = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetArrStation(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.ArrStation = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetArrTime(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.ArrTime = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetBookTime(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.BookTime = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetBookerId(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.BookerId = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetBookerName(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.BookerName = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetCapitalDirection(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.CapitalDirection = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetCascadeDepartment(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.CascadeDepartment = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetChangeFee(v float64) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.ChangeFee = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetCostCenter(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.CostCenter = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetCostCenterNumber(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.CostCenterNumber = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetCoupon(v float64) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.Coupon = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetDepartment(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.Department = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetDepartmentId(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.DepartmentId = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetDeptDate(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.DeptDate = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetDeptStation(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.DeptStation = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetDeptTime(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.DeptTime = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetFeeType(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.FeeType = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetIndex(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.Index = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetInvoiceTitle(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.InvoiceTitle = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetOrderId(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.OrderId = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetOrderPrice(v float64) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.OrderPrice = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetOverApplyId(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.OverApplyId = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetPrimaryId(v int64) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.PrimaryId = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetProjectCode(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.ProjectCode = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetProjectName(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.ProjectName = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetRefundFee(v float64) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.RefundFee = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetRunTime(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.RunTime = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetSeatNo(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.SeatNo = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetSeatType(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.SeatType = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetServiceFee(v float64) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.ServiceFee = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetSettlementFee(v float64) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.SettlementFee = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetSettlementTime(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.SettlementTime = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetSettlementType(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.SettlementType = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetStatus(v int64) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.Status = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetTicketNo(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.TicketNo = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetTicketPrice(v float64) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.TicketPrice = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetTrainNo(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.TrainNo = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetTrainType(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.TrainType = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetTravelerId(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.TravelerId = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetTravelerName(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.TravelerName = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetBookerJobNo(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.BookerJobNo = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetTravelerJobNo(v string) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.TravelerJobNo = &v
	return s
}

func (s *BillSettementBtripTrainResponseBodyModuleDataList) SetVoucherType(v int64) *BillSettementBtripTrainResponseBodyModuleDataList {
	s.VoucherType = &v
	return s
}

type BillSettementBtripTrainResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BillSettementBtripTrainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BillSettementBtripTrainResponse) String() string {
	return tea.Prettify(s)
}

func (s BillSettementBtripTrainResponse) GoString() string {
	return s.String()
}

func (s *BillSettementBtripTrainResponse) SetHeaders(v map[string]*string) *BillSettementBtripTrainResponse {
	s.Headers = v
	return s
}

func (s *BillSettementBtripTrainResponse) SetBody(v *BillSettementBtripTrainResponseBody) *BillSettementBtripTrainResponse {
	s.Body = v
	return s
}

type SyncExceedApplyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SyncExceedApplyHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncExceedApplyHeaders) GoString() string {
	return s.String()
}

func (s *SyncExceedApplyHeaders) SetCommonHeaders(v map[string]*string) *SyncExceedApplyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncExceedApplyHeaders) SetXAcsDingtalkAccessToken(v string) *SyncExceedApplyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SyncExceedApplyRequest struct {
	// 审批意见
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 商旅超标审批单id
	ApplyId *string `json:"applyId,omitempty" xml:"applyId,omitempty"`
	// 企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 第三方流程实例id
	ThirdpartyFlowId *string `json:"thirdpartyFlowId,omitempty" xml:"thirdpartyFlowId,omitempty"`
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 审批单状态 1同意2拒绝
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SyncExceedApplyRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncExceedApplyRequest) GoString() string {
	return s.String()
}

func (s *SyncExceedApplyRequest) SetRemark(v string) *SyncExceedApplyRequest {
	s.Remark = &v
	return s
}

func (s *SyncExceedApplyRequest) SetApplyId(v string) *SyncExceedApplyRequest {
	s.ApplyId = &v
	return s
}

func (s *SyncExceedApplyRequest) SetCorpId(v string) *SyncExceedApplyRequest {
	s.CorpId = &v
	return s
}

func (s *SyncExceedApplyRequest) SetThirdpartyFlowId(v string) *SyncExceedApplyRequest {
	s.ThirdpartyFlowId = &v
	return s
}

func (s *SyncExceedApplyRequest) SetUserId(v string) *SyncExceedApplyRequest {
	s.UserId = &v
	return s
}

func (s *SyncExceedApplyRequest) SetStatus(v int32) *SyncExceedApplyRequest {
	s.Status = &v
	return s
}

type SyncExceedApplyResponseBody struct {
	// 是否同步成功
	Module *bool `json:"module,omitempty" xml:"module,omitempty"`
}

func (s SyncExceedApplyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncExceedApplyResponseBody) GoString() string {
	return s.String()
}

func (s *SyncExceedApplyResponseBody) SetModule(v bool) *SyncExceedApplyResponseBody {
	s.Module = &v
	return s
}

type SyncExceedApplyResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SyncExceedApplyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncExceedApplyResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncExceedApplyResponse) GoString() string {
	return s.String()
}

func (s *SyncExceedApplyResponse) SetHeaders(v map[string]*string) *SyncExceedApplyResponse {
	s.Headers = v
	return s
}

func (s *SyncExceedApplyResponse) SetBody(v *SyncExceedApplyResponseBody) *SyncExceedApplyResponse {
	s.Body = v
	return s
}

type AddCityCarApplyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddCityCarApplyHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddCityCarApplyHeaders) GoString() string {
	return s.String()
}

func (s *AddCityCarApplyHeaders) SetCommonHeaders(v map[string]*string) *AddCityCarApplyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddCityCarApplyHeaders) SetXAcsDingtalkAccessToken(v string) *AddCityCarApplyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddCityCarApplyRequest struct {
	// 出差事由
	Cause *string `json:"cause,omitempty" xml:"cause,omitempty"`
	// 用车城市
	City *string `json:"city,omitempty" xml:"city,omitempty"`
	// 第三方企业ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 用车时间，按天管控，比如传值2021-03-18 20:26:56表示2021-03-18当天可用车，跨天情况配合finishedDate参数使用
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// 审批单关联的项目code
	ProjectCode *string `json:"projectCode,omitempty" xml:"projectCode,omitempty"`
	// 审批单关联的项目名
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// 审批单状态：0-申请，1-同意，2-拒绝
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 三方审批单ID
	ThirdPartApplyId *string `json:"thirdPartApplyId,omitempty" xml:"thirdPartApplyId,omitempty"`
	// 审批单关联的三方成本中心ID
	ThirdPartCostCenterId *string `json:"thirdPartCostCenterId,omitempty" xml:"thirdPartCostCenterId,omitempty"`
	// 审批单关联的三方发票抬头ID
	ThirdPartInvoiceId *string `json:"thirdPartInvoiceId,omitempty" xml:"thirdPartInvoiceId,omitempty"`
	// 审批单可用总次数
	TimesTotal *int64 `json:"timesTotal,omitempty" xml:"timesTotal,omitempty"`
	// 审批单可用次数类型：1-次数不限制，2-用户可指定次数，3-管理员限制次数；如果企业没有限制审批单使用次数的需求，这个参数传1(次数不限制)，同时times_total和times_used都传0即可
	TimesType *int64 `json:"timesType,omitempty" xml:"timesType,omitempty"`
	// 审批单已用次数
	TimesUsed *int64 `json:"timesUsed,omitempty" xml:"timesUsed,omitempty"`
	// 审批单标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 发起审批的第三方员工ID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// suiteKey
	DingSuiteKey *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	// account
	DingCorpId *string `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	// tokenGrantType
	DingTokenGrantType *int64 `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	// 用车截止时间，按天管控，比如date传值2021-03-18 20:26:56、finished_date传值2021-03-30 20:26:56表示2021-03-18(含)到2021-03-30(含)之间可用车，该参数不传值情况使用date作为用车截止时间；
	FinishedDate *string `json:"finishedDate,omitempty" xml:"finishedDate,omitempty"`
}

func (s AddCityCarApplyRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCityCarApplyRequest) GoString() string {
	return s.String()
}

func (s *AddCityCarApplyRequest) SetCause(v string) *AddCityCarApplyRequest {
	s.Cause = &v
	return s
}

func (s *AddCityCarApplyRequest) SetCity(v string) *AddCityCarApplyRequest {
	s.City = &v
	return s
}

func (s *AddCityCarApplyRequest) SetCorpId(v string) *AddCityCarApplyRequest {
	s.CorpId = &v
	return s
}

func (s *AddCityCarApplyRequest) SetDate(v string) *AddCityCarApplyRequest {
	s.Date = &v
	return s
}

func (s *AddCityCarApplyRequest) SetProjectCode(v string) *AddCityCarApplyRequest {
	s.ProjectCode = &v
	return s
}

func (s *AddCityCarApplyRequest) SetProjectName(v string) *AddCityCarApplyRequest {
	s.ProjectName = &v
	return s
}

func (s *AddCityCarApplyRequest) SetStatus(v int64) *AddCityCarApplyRequest {
	s.Status = &v
	return s
}

func (s *AddCityCarApplyRequest) SetThirdPartApplyId(v string) *AddCityCarApplyRequest {
	s.ThirdPartApplyId = &v
	return s
}

func (s *AddCityCarApplyRequest) SetThirdPartCostCenterId(v string) *AddCityCarApplyRequest {
	s.ThirdPartCostCenterId = &v
	return s
}

func (s *AddCityCarApplyRequest) SetThirdPartInvoiceId(v string) *AddCityCarApplyRequest {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *AddCityCarApplyRequest) SetTimesTotal(v int64) *AddCityCarApplyRequest {
	s.TimesTotal = &v
	return s
}

func (s *AddCityCarApplyRequest) SetTimesType(v int64) *AddCityCarApplyRequest {
	s.TimesType = &v
	return s
}

func (s *AddCityCarApplyRequest) SetTimesUsed(v int64) *AddCityCarApplyRequest {
	s.TimesUsed = &v
	return s
}

func (s *AddCityCarApplyRequest) SetTitle(v string) *AddCityCarApplyRequest {
	s.Title = &v
	return s
}

func (s *AddCityCarApplyRequest) SetUserId(v string) *AddCityCarApplyRequest {
	s.UserId = &v
	return s
}

func (s *AddCityCarApplyRequest) SetDingSuiteKey(v string) *AddCityCarApplyRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *AddCityCarApplyRequest) SetDingCorpId(v string) *AddCityCarApplyRequest {
	s.DingCorpId = &v
	return s
}

func (s *AddCityCarApplyRequest) SetDingTokenGrantType(v int64) *AddCityCarApplyRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *AddCityCarApplyRequest) SetFinishedDate(v string) *AddCityCarApplyRequest {
	s.FinishedDate = &v
	return s
}

type AddCityCarApplyResponseBody struct {
	// 商旅内部审批单ID
	ApplyId *int64 `json:"applyId,omitempty" xml:"applyId,omitempty"`
}

func (s AddCityCarApplyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCityCarApplyResponseBody) GoString() string {
	return s.String()
}

func (s *AddCityCarApplyResponseBody) SetApplyId(v int64) *AddCityCarApplyResponseBody {
	s.ApplyId = &v
	return s
}

type AddCityCarApplyResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddCityCarApplyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddCityCarApplyResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCityCarApplyResponse) GoString() string {
	return s.String()
}

func (s *AddCityCarApplyResponse) SetHeaders(v map[string]*string) *AddCityCarApplyResponse {
	s.Headers = v
	return s
}

func (s *AddCityCarApplyResponse) SetBody(v *AddCityCarApplyResponseBody) *AddCityCarApplyResponse {
	s.Body = v
	return s
}

type BillSettementFlightHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BillSettementFlightHeaders) String() string {
	return tea.Prettify(s)
}

func (s BillSettementFlightHeaders) GoString() string {
	return s.String()
}

func (s *BillSettementFlightHeaders) SetCommonHeaders(v map[string]*string) *BillSettementFlightHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BillSettementFlightHeaders) SetXAcsDingtalkAccessToken(v string) *BillSettementFlightHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BillSettementFlightRequest struct {
	// 第三方企业ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 类目：机酒火车 1：机票； 2：酒店； 4：用车 6：商旅火车票
	Category *int64 `json:"category,omitempty" xml:"category,omitempty"`
	// 每页数据量，默认100，最高500
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 记账更新开始日期
	PeriodStart *string `json:"periodStart,omitempty" xml:"periodStart,omitempty"`
	// 页数，从1开始
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 记账更新结束日期
	PeriodEnd *string `json:"periodEnd,omitempty" xml:"periodEnd,omitempty"`
}

func (s BillSettementFlightRequest) String() string {
	return tea.Prettify(s)
}

func (s BillSettementFlightRequest) GoString() string {
	return s.String()
}

func (s *BillSettementFlightRequest) SetCorpId(v string) *BillSettementFlightRequest {
	s.CorpId = &v
	return s
}

func (s *BillSettementFlightRequest) SetCategory(v int64) *BillSettementFlightRequest {
	s.Category = &v
	return s
}

func (s *BillSettementFlightRequest) SetPageSize(v int64) *BillSettementFlightRequest {
	s.PageSize = &v
	return s
}

func (s *BillSettementFlightRequest) SetPeriodStart(v string) *BillSettementFlightRequest {
	s.PeriodStart = &v
	return s
}

func (s *BillSettementFlightRequest) SetPageNumber(v int64) *BillSettementFlightRequest {
	s.PageNumber = &v
	return s
}

func (s *BillSettementFlightRequest) SetPeriodEnd(v string) *BillSettementFlightRequest {
	s.PeriodEnd = &v
	return s
}

type BillSettementFlightResponseBody struct {
	// 结果msg
	ResultMsg *string `json:"resultMsg,omitempty" xml:"resultMsg,omitempty"`
	// module
	Module *BillSettementFlightResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 结果code
	ResultCode *int64 `json:"resultCode,omitempty" xml:"resultCode,omitempty"`
}

func (s BillSettementFlightResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BillSettementFlightResponseBody) GoString() string {
	return s.String()
}

func (s *BillSettementFlightResponseBody) SetResultMsg(v string) *BillSettementFlightResponseBody {
	s.ResultMsg = &v
	return s
}

func (s *BillSettementFlightResponseBody) SetModule(v *BillSettementFlightResponseBodyModule) *BillSettementFlightResponseBody {
	s.Module = v
	return s
}

func (s *BillSettementFlightResponseBody) SetSuccess(v bool) *BillSettementFlightResponseBody {
	s.Success = &v
	return s
}

func (s *BillSettementFlightResponseBody) SetResultCode(v int64) *BillSettementFlightResponseBody {
	s.ResultCode = &v
	return s
}

type BillSettementFlightResponseBodyModule struct {
	// 类目
	Category *int64 `json:"category,omitempty" xml:"category,omitempty"`
	// 企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 数据集合
	DataList []*BillSettementFlightResponseBodyModuleDataList `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 记账更新开始日期
	PeriodEnd *string `json:"periodEnd,omitempty" xml:"periodEnd,omitempty"`
	// 记账更新结束日期
	PeriodStart *string `json:"periodStart,omitempty" xml:"periodStart,omitempty"`
	// 总数据量
	TotalNum *int64 `json:"totalNum,omitempty" xml:"totalNum,omitempty"`
}

func (s BillSettementFlightResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s BillSettementFlightResponseBodyModule) GoString() string {
	return s.String()
}

func (s *BillSettementFlightResponseBodyModule) SetCategory(v int64) *BillSettementFlightResponseBodyModule {
	s.Category = &v
	return s
}

func (s *BillSettementFlightResponseBodyModule) SetCorpId(v string) *BillSettementFlightResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *BillSettementFlightResponseBodyModule) SetDataList(v []*BillSettementFlightResponseBodyModuleDataList) *BillSettementFlightResponseBodyModule {
	s.DataList = v
	return s
}

func (s *BillSettementFlightResponseBodyModule) SetPeriodEnd(v string) *BillSettementFlightResponseBodyModule {
	s.PeriodEnd = &v
	return s
}

func (s *BillSettementFlightResponseBodyModule) SetPeriodStart(v string) *BillSettementFlightResponseBodyModule {
	s.PeriodStart = &v
	return s
}

func (s *BillSettementFlightResponseBodyModule) SetTotalNum(v int64) *BillSettementFlightResponseBodyModule {
	s.TotalNum = &v
	return s
}

type BillSettementFlightResponseBodyModuleDataList struct {
	// 提前预定天数
	AdvanceDay *int64 `json:"advanceDay,omitempty" xml:"advanceDay,omitempty"`
	// 航司三字码
	AirlineCorpCode *string `json:"airlineCorpCode,omitempty" xml:"airlineCorpCode,omitempty"`
	// 航司名称
	AirlineCorpName *string `json:"airlineCorpName,omitempty" xml:"airlineCorpName,omitempty"`
	// 交易流水号
	AlipayTradeNo *string `json:"alipayTradeNo,omitempty" xml:"alipayTradeNo,omitempty"`
	// 审批单号
	ApplyId *string `json:"applyId,omitempty" xml:"applyId,omitempty"`
	// 到达机场二字码
	ArrAirportCode *string `json:"arrAirportCode,omitempty" xml:"arrAirportCode,omitempty"`
	// 到达城市
	ArrCity *string `json:"arrCity,omitempty" xml:"arrCity,omitempty"`
	// 到达日期
	ArrDate *string `json:"arrDate,omitempty" xml:"arrDate,omitempty"`
	// 到达机场
	ArrStation *string `json:"arrStation,omitempty" xml:"arrStation,omitempty"`
	// 到达时间
	ArrTime *string `json:"arrTime,omitempty" xml:"arrTime,omitempty"`
	// 预定时间
	BookTime *string `json:"bookTime,omitempty" xml:"bookTime,omitempty"`
	// 预订人use id
	BookerId *string `json:"bookerId,omitempty" xml:"bookerId,omitempty"`
	// 预订人名称
	BookerName *string `json:"bookerName,omitempty" xml:"bookerName,omitempty"`
	// 商旅优惠金额
	BtripCouponFee *float64 `json:"btripCouponFee,omitempty" xml:"btripCouponFee,omitempty"`
	// 基建费
	BuildFee *float64 `json:"buildFee,omitempty" xml:"buildFee,omitempty"`
	// 舱位
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// 舱位码
	CabinClass *string `json:"cabinClass,omitempty" xml:"cabinClass,omitempty"`
	// 资金方向
	CapitalDirection *string `json:"capitalDirection,omitempty" xml:"capitalDirection,omitempty"`
	// 级联部门
	CascadeDepartment *string `json:"cascadeDepartment,omitempty" xml:"cascadeDepartment,omitempty"`
	// 改签费用
	ChangeFee *float64 `json:"changeFee,omitempty" xml:"changeFee,omitempty"`
	// 订单金额
	CorpPayOrderFee *float64 `json:"corpPayOrderFee,omitempty" xml:"corpPayOrderFee,omitempty"`
	// 成本中心名称
	CostCenter *string `json:"costCenter,omitempty" xml:"costCenter,omitempty"`
	// 成本中心编号
	CostCenterNumber *string `json:"costCenterNumber,omitempty" xml:"costCenterNumber,omitempty"`
	// 优惠券
	Coupon *float64 `json:"coupon,omitempty" xml:"coupon,omitempty"`
	// 起飞机场二字码
	DepAirportCode *string `json:"depAirportCode,omitempty" xml:"depAirportCode,omitempty"`
	// 末级部门
	Department *string `json:"department,omitempty" xml:"department,omitempty"`
	// 部门id
	DepartmentId *string `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 起飞城市
	DeptCity *string `json:"deptCity,omitempty" xml:"deptCity,omitempty"`
	// 起飞日期
	DeptDate *string `json:"deptDate,omitempty" xml:"deptDate,omitempty"`
	// 起飞机场
	DeptStation *string `json:"deptStation,omitempty" xml:"deptStation,omitempty"`
	// 起飞时间
	DeptTime *string `json:"deptTime,omitempty" xml:"deptTime,omitempty"`
	// 折扣率
	Discount *string `json:"discount,omitempty" xml:"discount,omitempty"`
	// 费用类型
	FeeType *string `json:"feeType,omitempty" xml:"feeType,omitempty"`
	// 航班号
	FlightNo *string `json:"flightNo,omitempty" xml:"flightNo,omitempty"`
	// 序号
	Index *string `json:"index,omitempty" xml:"index,omitempty"`
	// 保险费
	InsuranceFee *float64 `json:"insuranceFee,omitempty" xml:"insuranceFee,omitempty"`
	// 发票抬头
	InvoiceTitle *string `json:"invoiceTitle,omitempty" xml:"invoiceTitle,omitempty"`
	// 行程单打印序号
	ItineraryNum *string `json:"itineraryNum,omitempty" xml:"itineraryNum,omitempty"`
	// 行程单金额
	ItineraryPrice *float64 `json:"itineraryPrice,omitempty" xml:"itineraryPrice,omitempty"`
	// 低价提醒（起飞时间）
	MostDifferenceDeptTime *string `json:"mostDifferenceDeptTime,omitempty" xml:"mostDifferenceDeptTime,omitempty"`
	// 低价提醒（折扣）
	MostDifferenceDiscount *float64 `json:"mostDifferenceDiscount,omitempty" xml:"mostDifferenceDiscount,omitempty"`
	// 低价提醒(航班号)
	MostDifferenceFlightNo *string `json:"mostDifferenceFlightNo,omitempty" xml:"mostDifferenceFlightNo,omitempty"`
	// 低价提醒(与最低价差额)
	MostDifferencePrice *float64 `json:"mostDifferencePrice,omitempty" xml:"mostDifferencePrice,omitempty"`
	// 不选低价原因
	MostDifferenceReason *string `json:"mostDifferenceReason,omitempty" xml:"mostDifferenceReason,omitempty"`
	// 低价航班价格
	MostPrice *float64 `json:"mostPrice,omitempty" xml:"mostPrice,omitempty"`
	// 协议价优惠金额
	NegotiationCouponFee *float64 `json:"negotiationCouponFee,omitempty" xml:"negotiationCouponFee,omitempty"`
	// 燃油费
	OilFee *float64 `json:"oilFee,omitempty" xml:"oilFee,omitempty"`
	// 订单号
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// 超标审批单号
	OverApplyId *string `json:"overApplyId,omitempty" xml:"overApplyId,omitempty"`
	// 主键id
	PrimaryId *int64 `json:"primaryId,omitempty" xml:"primaryId,omitempty"`
	// 项目代码
	ProjectCode *string `json:"projectCode,omitempty" xml:"projectCode,omitempty"`
	// 项目名称
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// 退款手续费
	RefundFee *float64 `json:"refundFee,omitempty" xml:"refundFee,omitempty"`
	// 改签退票手续费
	RefundUpgradeCost *float64 `json:"refundUpgradeCost,omitempty" xml:"refundUpgradeCost,omitempty"`
	// 是否重复退
	RepeatRefund *string `json:"repeatRefund,omitempty" xml:"repeatRefund,omitempty"`
	// 销售价
	SealPrice *float64 `json:"sealPrice,omitempty" xml:"sealPrice,omitempty"`
	// 服务费，仅在feeType  11001、11002中展示
	ServiceFee *float64 `json:"serviceFee,omitempty" xml:"serviceFee,omitempty"`
	// 结算金额
	SettlementFee *float64 `json:"settlementFee,omitempty" xml:"settlementFee,omitempty"`
	// 结算时间
	SettlementTime *string `json:"settlementTime,omitempty" xml:"settlementTime,omitempty"`
	// 结算类型
	SettlementType *string `json:"settlementType,omitempty" xml:"settlementType,omitempty"`
	// 入账状态
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 行程单号
	TicketId *string `json:"ticketId,omitempty" xml:"ticketId,omitempty"`
	// 出行人use id
	TravelerId *string `json:"travelerId,omitempty" xml:"travelerId,omitempty"`
	// 出行人名称
	TravelerName *string `json:"travelerName,omitempty" xml:"travelerName,omitempty"`
	// 改签差价
	UpgradeCost *float64 `json:"upgradeCost,omitempty" xml:"upgradeCost,omitempty"`
	// 预订人工号
	BookerJobNo *string `json:"bookerJobNo,omitempty" xml:"bookerJobNo,omitempty"`
	// 出行人工号
	TravelerJobNo *string `json:"travelerJobNo,omitempty" xml:"travelerJobNo,omitempty"`
}

func (s BillSettementFlightResponseBodyModuleDataList) String() string {
	return tea.Prettify(s)
}

func (s BillSettementFlightResponseBodyModuleDataList) GoString() string {
	return s.String()
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetAdvanceDay(v int64) *BillSettementFlightResponseBodyModuleDataList {
	s.AdvanceDay = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetAirlineCorpCode(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.AirlineCorpCode = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetAirlineCorpName(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.AirlineCorpName = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetAlipayTradeNo(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.AlipayTradeNo = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetApplyId(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.ApplyId = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetArrAirportCode(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.ArrAirportCode = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetArrCity(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.ArrCity = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetArrDate(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.ArrDate = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetArrStation(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.ArrStation = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetArrTime(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.ArrTime = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetBookTime(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.BookTime = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetBookerId(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.BookerId = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetBookerName(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.BookerName = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetBtripCouponFee(v float64) *BillSettementFlightResponseBodyModuleDataList {
	s.BtripCouponFee = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetBuildFee(v float64) *BillSettementFlightResponseBodyModuleDataList {
	s.BuildFee = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetCabin(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.Cabin = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetCabinClass(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.CabinClass = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetCapitalDirection(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.CapitalDirection = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetCascadeDepartment(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.CascadeDepartment = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetChangeFee(v float64) *BillSettementFlightResponseBodyModuleDataList {
	s.ChangeFee = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetCorpPayOrderFee(v float64) *BillSettementFlightResponseBodyModuleDataList {
	s.CorpPayOrderFee = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetCostCenter(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.CostCenter = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetCostCenterNumber(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.CostCenterNumber = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetCoupon(v float64) *BillSettementFlightResponseBodyModuleDataList {
	s.Coupon = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetDepAirportCode(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.DepAirportCode = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetDepartment(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.Department = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetDepartmentId(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.DepartmentId = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetDeptCity(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.DeptCity = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetDeptDate(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.DeptDate = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetDeptStation(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.DeptStation = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetDeptTime(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.DeptTime = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetDiscount(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.Discount = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetFeeType(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.FeeType = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetFlightNo(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.FlightNo = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetIndex(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.Index = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetInsuranceFee(v float64) *BillSettementFlightResponseBodyModuleDataList {
	s.InsuranceFee = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetInvoiceTitle(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.InvoiceTitle = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetItineraryNum(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.ItineraryNum = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetItineraryPrice(v float64) *BillSettementFlightResponseBodyModuleDataList {
	s.ItineraryPrice = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetMostDifferenceDeptTime(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.MostDifferenceDeptTime = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetMostDifferenceDiscount(v float64) *BillSettementFlightResponseBodyModuleDataList {
	s.MostDifferenceDiscount = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetMostDifferenceFlightNo(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.MostDifferenceFlightNo = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetMostDifferencePrice(v float64) *BillSettementFlightResponseBodyModuleDataList {
	s.MostDifferencePrice = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetMostDifferenceReason(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.MostDifferenceReason = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetMostPrice(v float64) *BillSettementFlightResponseBodyModuleDataList {
	s.MostPrice = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetNegotiationCouponFee(v float64) *BillSettementFlightResponseBodyModuleDataList {
	s.NegotiationCouponFee = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetOilFee(v float64) *BillSettementFlightResponseBodyModuleDataList {
	s.OilFee = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetOrderId(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.OrderId = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetOverApplyId(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.OverApplyId = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetPrimaryId(v int64) *BillSettementFlightResponseBodyModuleDataList {
	s.PrimaryId = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetProjectCode(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.ProjectCode = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetProjectName(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.ProjectName = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetRefundFee(v float64) *BillSettementFlightResponseBodyModuleDataList {
	s.RefundFee = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetRefundUpgradeCost(v float64) *BillSettementFlightResponseBodyModuleDataList {
	s.RefundUpgradeCost = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetRepeatRefund(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.RepeatRefund = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetSealPrice(v float64) *BillSettementFlightResponseBodyModuleDataList {
	s.SealPrice = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetServiceFee(v float64) *BillSettementFlightResponseBodyModuleDataList {
	s.ServiceFee = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetSettlementFee(v float64) *BillSettementFlightResponseBodyModuleDataList {
	s.SettlementFee = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetSettlementTime(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.SettlementTime = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetSettlementType(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.SettlementType = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetStatus(v int64) *BillSettementFlightResponseBodyModuleDataList {
	s.Status = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetTicketId(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.TicketId = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetTravelerId(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.TravelerId = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetTravelerName(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.TravelerName = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetUpgradeCost(v float64) *BillSettementFlightResponseBodyModuleDataList {
	s.UpgradeCost = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetBookerJobNo(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.BookerJobNo = &v
	return s
}

func (s *BillSettementFlightResponseBodyModuleDataList) SetTravelerJobNo(v string) *BillSettementFlightResponseBodyModuleDataList {
	s.TravelerJobNo = &v
	return s
}

type BillSettementFlightResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BillSettementFlightResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BillSettementFlightResponse) String() string {
	return tea.Prettify(s)
}

func (s BillSettementFlightResponse) GoString() string {
	return s.String()
}

func (s *BillSettementFlightResponse) SetHeaders(v map[string]*string) *BillSettementFlightResponse {
	s.Headers = v
	return s
}

func (s *BillSettementFlightResponse) SetBody(v *BillSettementFlightResponseBody) *BillSettementFlightResponse {
	s.Body = v
	return s
}

type GetHotelExceedApplyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetHotelExceedApplyHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetHotelExceedApplyHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelExceedApplyHeaders) SetCommonHeaders(v map[string]*string) *GetHotelExceedApplyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelExceedApplyHeaders) SetXAcsDingtalkAccessToken(v string) *GetHotelExceedApplyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetHotelExceedApplyRequest struct {
	// 第三方企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 商旅超标审批单id
	ApplyId *string `json:"applyId,omitempty" xml:"applyId,omitempty"`
}

func (s GetHotelExceedApplyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotelExceedApplyRequest) GoString() string {
	return s.String()
}

func (s *GetHotelExceedApplyRequest) SetCorpId(v string) *GetHotelExceedApplyRequest {
	s.CorpId = &v
	return s
}

func (s *GetHotelExceedApplyRequest) SetApplyId(v string) *GetHotelExceedApplyRequest {
	s.ApplyId = &v
	return s
}

type GetHotelExceedApplyResponseBody struct {
	// 第三方企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 商旅超标审批单id
	ApplyId *int64 `json:"applyId,omitempty" xml:"applyId,omitempty"`
	// 审批单状态 0:审批中 1:已同意 2:已拒绝
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// 出差原因
	BtripCause *string `json:"btripCause,omitempty" xml:"btripCause,omitempty"`
	// 超标类型，32：金额超标
	ExceedType *int32 `json:"exceedType,omitempty" xml:"exceedType,omitempty"`
	// 超标原因
	ExceedReason *string `json:"exceedReason,omitempty" xml:"exceedReason,omitempty"`
	// 原差旅标准
	OriginStandard *string `json:"originStandard,omitempty" xml:"originStandard,omitempty"`
	// 审批单提交时间
	SubmitTime *string `json:"submitTime,omitempty" xml:"submitTime,omitempty"`
	// 第三方用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 意向出行信息
	ApplyIntentionInfoDO *GetHotelExceedApplyResponseBodyApplyIntentionInfoDO `json:"applyIntentionInfoDO,omitempty" xml:"applyIntentionInfoDO,omitempty" type:"Struct"`
	// 第三方出差审批单号
	ThirdpartApplyId *string `json:"thirdpartApplyId,omitempty" xml:"thirdpartApplyId,omitempty"`
}

func (s GetHotelExceedApplyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotelExceedApplyResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelExceedApplyResponseBody) SetCorpId(v string) *GetHotelExceedApplyResponseBody {
	s.CorpId = &v
	return s
}

func (s *GetHotelExceedApplyResponseBody) SetApplyId(v int64) *GetHotelExceedApplyResponseBody {
	s.ApplyId = &v
	return s
}

func (s *GetHotelExceedApplyResponseBody) SetStatus(v int32) *GetHotelExceedApplyResponseBody {
	s.Status = &v
	return s
}

func (s *GetHotelExceedApplyResponseBody) SetBtripCause(v string) *GetHotelExceedApplyResponseBody {
	s.BtripCause = &v
	return s
}

func (s *GetHotelExceedApplyResponseBody) SetExceedType(v int32) *GetHotelExceedApplyResponseBody {
	s.ExceedType = &v
	return s
}

func (s *GetHotelExceedApplyResponseBody) SetExceedReason(v string) *GetHotelExceedApplyResponseBody {
	s.ExceedReason = &v
	return s
}

func (s *GetHotelExceedApplyResponseBody) SetOriginStandard(v string) *GetHotelExceedApplyResponseBody {
	s.OriginStandard = &v
	return s
}

func (s *GetHotelExceedApplyResponseBody) SetSubmitTime(v string) *GetHotelExceedApplyResponseBody {
	s.SubmitTime = &v
	return s
}

func (s *GetHotelExceedApplyResponseBody) SetUserId(v string) *GetHotelExceedApplyResponseBody {
	s.UserId = &v
	return s
}

func (s *GetHotelExceedApplyResponseBody) SetApplyIntentionInfoDO(v *GetHotelExceedApplyResponseBodyApplyIntentionInfoDO) *GetHotelExceedApplyResponseBody {
	s.ApplyIntentionInfoDO = v
	return s
}

func (s *GetHotelExceedApplyResponseBody) SetThirdpartApplyId(v string) *GetHotelExceedApplyResponseBody {
	s.ThirdpartApplyId = &v
	return s
}

type GetHotelExceedApplyResponseBodyApplyIntentionInfoDO struct {
	// 入住日期
	CheckIn *string `json:"checkIn,omitempty" xml:"checkIn,omitempty"`
	// 离店日期
	CheckOut *string `json:"checkOut,omitempty" xml:"checkOut,omitempty"`
	// 入住城市三字码
	CityCode *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	// 入住城市名称
	CityName *string `json:"cityName,omitempty" xml:"cityName,omitempty"`
	// 意向酒店金额（分）
	Price *int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 是否合住
	Together *bool `json:"together,omitempty" xml:"together,omitempty"`
	// 超标类型，32：金额超标
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetHotelExceedApplyResponseBodyApplyIntentionInfoDO) String() string {
	return tea.Prettify(s)
}

func (s GetHotelExceedApplyResponseBodyApplyIntentionInfoDO) GoString() string {
	return s.String()
}

func (s *GetHotelExceedApplyResponseBodyApplyIntentionInfoDO) SetCheckIn(v string) *GetHotelExceedApplyResponseBodyApplyIntentionInfoDO {
	s.CheckIn = &v
	return s
}

func (s *GetHotelExceedApplyResponseBodyApplyIntentionInfoDO) SetCheckOut(v string) *GetHotelExceedApplyResponseBodyApplyIntentionInfoDO {
	s.CheckOut = &v
	return s
}

func (s *GetHotelExceedApplyResponseBodyApplyIntentionInfoDO) SetCityCode(v string) *GetHotelExceedApplyResponseBodyApplyIntentionInfoDO {
	s.CityCode = &v
	return s
}

func (s *GetHotelExceedApplyResponseBodyApplyIntentionInfoDO) SetCityName(v string) *GetHotelExceedApplyResponseBodyApplyIntentionInfoDO {
	s.CityName = &v
	return s
}

func (s *GetHotelExceedApplyResponseBodyApplyIntentionInfoDO) SetPrice(v int64) *GetHotelExceedApplyResponseBodyApplyIntentionInfoDO {
	s.Price = &v
	return s
}

func (s *GetHotelExceedApplyResponseBodyApplyIntentionInfoDO) SetTogether(v bool) *GetHotelExceedApplyResponseBodyApplyIntentionInfoDO {
	s.Together = &v
	return s
}

func (s *GetHotelExceedApplyResponseBodyApplyIntentionInfoDO) SetType(v int32) *GetHotelExceedApplyResponseBodyApplyIntentionInfoDO {
	s.Type = &v
	return s
}

type GetHotelExceedApplyResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHotelExceedApplyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotelExceedApplyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotelExceedApplyResponse) GoString() string {
	return s.String()
}

func (s *GetHotelExceedApplyResponse) SetHeaders(v map[string]*string) *GetHotelExceedApplyResponse {
	s.Headers = v
	return s
}

func (s *GetHotelExceedApplyResponse) SetBody(v *GetHotelExceedApplyResponseBody) *GetHotelExceedApplyResponse {
	s.Body = v
	return s
}

type QueryUnionOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUnionOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUnionOrderHeaders) GoString() string {
	return s.String()
}

func (s *QueryUnionOrderHeaders) SetCommonHeaders(v map[string]*string) *QueryUnionOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUnionOrderHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUnionOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUnionOrderRequest struct {
	// 第三方企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 第三方申请单id
	ThirdPartApplyId *string `json:"thirdPartApplyId,omitempty" xml:"thirdPartApplyId,omitempty"`
	// 关联单号
	UnionNo *string `json:"unionNo,omitempty" xml:"unionNo,omitempty"`
}

func (s QueryUnionOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUnionOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryUnionOrderRequest) SetCorpId(v string) *QueryUnionOrderRequest {
	s.CorpId = &v
	return s
}

func (s *QueryUnionOrderRequest) SetThirdPartApplyId(v string) *QueryUnionOrderRequest {
	s.ThirdPartApplyId = &v
	return s
}

func (s *QueryUnionOrderRequest) SetUnionNo(v string) *QueryUnionOrderRequest {
	s.UnionNo = &v
	return s
}

type QueryUnionOrderResponseBody struct {
	// 飞机订单信息
	FlightList []*QueryUnionOrderResponseBodyFlightList `json:"flightList,omitempty" xml:"flightList,omitempty" type:"Repeated"`
	// 企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 火车订单信息
	TrainList []*QueryUnionOrderResponseBodyTrainList `json:"trainList,omitempty" xml:"trainList,omitempty" type:"Repeated"`
	// 酒店订单信息
	HotelList []*QueryUnionOrderResponseBodyHotelList `json:"hotelList,omitempty" xml:"hotelList,omitempty" type:"Repeated"`
	// 用车订单信息
	VehicleList []*QueryUnionOrderResponseBodyVehicleList `json:"vehicleList,omitempty" xml:"vehicleList,omitempty" type:"Repeated"`
}

func (s QueryUnionOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUnionOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUnionOrderResponseBody) SetFlightList(v []*QueryUnionOrderResponseBodyFlightList) *QueryUnionOrderResponseBody {
	s.FlightList = v
	return s
}

func (s *QueryUnionOrderResponseBody) SetCorpId(v string) *QueryUnionOrderResponseBody {
	s.CorpId = &v
	return s
}

func (s *QueryUnionOrderResponseBody) SetTrainList(v []*QueryUnionOrderResponseBodyTrainList) *QueryUnionOrderResponseBody {
	s.TrainList = v
	return s
}

func (s *QueryUnionOrderResponseBody) SetHotelList(v []*QueryUnionOrderResponseBodyHotelList) *QueryUnionOrderResponseBody {
	s.HotelList = v
	return s
}

func (s *QueryUnionOrderResponseBody) SetVehicleList(v []*QueryUnionOrderResponseBodyVehicleList) *QueryUnionOrderResponseBody {
	s.VehicleList = v
	return s
}

type QueryUnionOrderResponseBodyFlightList struct {
	// 订单id
	FlightOrderId *int64 `json:"flightOrderId,omitempty" xml:"flightOrderId,omitempty"`
	// 订单状态：0待支付,1出票中,2已关闭,3有改签单,4有退票单,5出票成功,6退票申请中,7改签申请中
	FlightOrderStatus *int64 `json:"flightOrderStatus,omitempty" xml:"flightOrderStatus,omitempty"`
}

func (s QueryUnionOrderResponseBodyFlightList) String() string {
	return tea.Prettify(s)
}

func (s QueryUnionOrderResponseBodyFlightList) GoString() string {
	return s.String()
}

func (s *QueryUnionOrderResponseBodyFlightList) SetFlightOrderId(v int64) *QueryUnionOrderResponseBodyFlightList {
	s.FlightOrderId = &v
	return s
}

func (s *QueryUnionOrderResponseBodyFlightList) SetFlightOrderStatus(v int64) *QueryUnionOrderResponseBodyFlightList {
	s.FlightOrderStatus = &v
	return s
}

type QueryUnionOrderResponseBodyTrainList struct {
	// 火车订单号
	TrainOrderId *int64 `json:"trainOrderId,omitempty" xml:"trainOrderId,omitempty"`
	// 订单状态：0待支付,1出票中,2已关闭,3,改签成功,4退票成功,5出票完成,6退票申请中,7改签申请中,8已出票,已发货,9出票失败,10改签失败,11退票失败
	TrainOrderstatus *int64 `json:"trainOrderstatus,omitempty" xml:"trainOrderstatus,omitempty"`
}

func (s QueryUnionOrderResponseBodyTrainList) String() string {
	return tea.Prettify(s)
}

func (s QueryUnionOrderResponseBodyTrainList) GoString() string {
	return s.String()
}

func (s *QueryUnionOrderResponseBodyTrainList) SetTrainOrderId(v int64) *QueryUnionOrderResponseBodyTrainList {
	s.TrainOrderId = &v
	return s
}

func (s *QueryUnionOrderResponseBodyTrainList) SetTrainOrderstatus(v int64) *QueryUnionOrderResponseBodyTrainList {
	s.TrainOrderstatus = &v
	return s
}

type QueryUnionOrderResponseBodyHotelList struct {
	// 酒店订单号
	HotelOrderId *int64 `json:"hotelOrderId,omitempty" xml:"hotelOrderId,omitempty"`
	// 订单状态1:等待确认,2:等待付款,3:预订成功,4:申请退款,5:退款成功,6:已关闭,7:结账成功,8:支付成功
	HotelOrderStatus *int64 `json:"hotelOrderStatus,omitempty" xml:"hotelOrderStatus,omitempty"`
}

func (s QueryUnionOrderResponseBodyHotelList) String() string {
	return tea.Prettify(s)
}

func (s QueryUnionOrderResponseBodyHotelList) GoString() string {
	return s.String()
}

func (s *QueryUnionOrderResponseBodyHotelList) SetHotelOrderId(v int64) *QueryUnionOrderResponseBodyHotelList {
	s.HotelOrderId = &v
	return s
}

func (s *QueryUnionOrderResponseBodyHotelList) SetHotelOrderStatus(v int64) *QueryUnionOrderResponseBodyHotelList {
	s.HotelOrderStatus = &v
	return s
}

type QueryUnionOrderResponseBodyVehicleList struct {
	// 用车订单号
	VehicleOrderId *int64 `json:"vehicleOrderId,omitempty" xml:"vehicleOrderId,omitempty"`
	// 订单状态:0:初始状态,1:已超时,2:派单成功,3:派单失败,4:已退款,5:已支付,6:已取消
	VehicleOrderStatus *int64 `json:"vehicleOrderStatus,omitempty" xml:"vehicleOrderStatus,omitempty"`
}

func (s QueryUnionOrderResponseBodyVehicleList) String() string {
	return tea.Prettify(s)
}

func (s QueryUnionOrderResponseBodyVehicleList) GoString() string {
	return s.String()
}

func (s *QueryUnionOrderResponseBodyVehicleList) SetVehicleOrderId(v int64) *QueryUnionOrderResponseBodyVehicleList {
	s.VehicleOrderId = &v
	return s
}

func (s *QueryUnionOrderResponseBodyVehicleList) SetVehicleOrderStatus(v int64) *QueryUnionOrderResponseBodyVehicleList {
	s.VehicleOrderStatus = &v
	return s
}

type QueryUnionOrderResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUnionOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUnionOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUnionOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryUnionOrderResponse) SetHeaders(v map[string]*string) *QueryUnionOrderResponse {
	s.Headers = v
	return s
}

func (s *QueryUnionOrderResponse) SetBody(v *QueryUnionOrderResponseBody) *QueryUnionOrderResponse {
	s.Body = v
	return s
}

type QueryCityCarApplyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCityCarApplyHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCityCarApplyHeaders) GoString() string {
	return s.String()
}

func (s *QueryCityCarApplyHeaders) SetCommonHeaders(v map[string]*string) *QueryCityCarApplyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCityCarApplyHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCityCarApplyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCityCarApplyRequest struct {
	// 第三方企业ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 审批单创建时间小于值
	CreatedEndAt *string `json:"createdEndAt,omitempty" xml:"createdEndAt,omitempty"`
	// 审批单创建时间大于等于值
	CreatedStartAt *string `json:"createdStartAt,omitempty" xml:"createdStartAt,omitempty"`
	// 页码，要求大于等于1，默认1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页数据量，要求大于等于1，默认20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 三方审批单ID
	ThirdPartApplyId *string `json:"thirdPartApplyId,omitempty" xml:"thirdPartApplyId,omitempty"`
	// 第三方员工ID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryCityCarApplyRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCityCarApplyRequest) GoString() string {
	return s.String()
}

func (s *QueryCityCarApplyRequest) SetCorpId(v string) *QueryCityCarApplyRequest {
	s.CorpId = &v
	return s
}

func (s *QueryCityCarApplyRequest) SetCreatedEndAt(v string) *QueryCityCarApplyRequest {
	s.CreatedEndAt = &v
	return s
}

func (s *QueryCityCarApplyRequest) SetCreatedStartAt(v string) *QueryCityCarApplyRequest {
	s.CreatedStartAt = &v
	return s
}

func (s *QueryCityCarApplyRequest) SetPageNumber(v int64) *QueryCityCarApplyRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryCityCarApplyRequest) SetPageSize(v int64) *QueryCityCarApplyRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCityCarApplyRequest) SetThirdPartApplyId(v string) *QueryCityCarApplyRequest {
	s.ThirdPartApplyId = &v
	return s
}

func (s *QueryCityCarApplyRequest) SetUserId(v string) *QueryCityCarApplyRequest {
	s.UserId = &v
	return s
}

type QueryCityCarApplyResponseBody struct {
	// 审批单列表
	ApplyList []*QueryCityCarApplyResponseBodyApplyList `json:"applyList,omitempty" xml:"applyList,omitempty" type:"Repeated"`
	// 总数
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s QueryCityCarApplyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCityCarApplyResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCityCarApplyResponseBody) SetApplyList(v []*QueryCityCarApplyResponseBodyApplyList) *QueryCityCarApplyResponseBody {
	s.ApplyList = v
	return s
}

func (s *QueryCityCarApplyResponseBody) SetTotal(v int64) *QueryCityCarApplyResponseBody {
	s.Total = &v
	return s
}

type QueryCityCarApplyResponseBodyApplyList struct {
	// 审批单列表
	ApproverList []*QueryCityCarApplyResponseBodyApplyListApproverList `json:"approverList,omitempty" xml:"approverList,omitempty" type:"Repeated"`
	// 员工所在部门ID
	DepartId *string `json:"departId,omitempty" xml:"departId,omitempty"`
	// 员工所在部门名
	DepartName *string `json:"departName,omitempty" xml:"departName,omitempty"`
	// 创建时间
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 最近修改时间
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 审批单关联的行程
	ItineraryList []*QueryCityCarApplyResponseBodyApplyListItineraryList `json:"itineraryList,omitempty" xml:"itineraryList,omitempty" type:"Repeated"`
	// 审批单状态：0-申请，1-同意，2-拒绝
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 审批单状态：0-申请，1-同意，2-拒绝
	StatusDesc *string `json:"statusDesc,omitempty" xml:"statusDesc,omitempty"`
	// 三方审批单ID
	ThirdPartApplyId *string `json:"thirdPartApplyId,omitempty" xml:"thirdPartApplyId,omitempty"`
	// 申请事由
	TripCause *string `json:"tripCause,omitempty" xml:"tripCause,omitempty"`
	// 审批单标题
	TripTitle *string `json:"tripTitle,omitempty" xml:"tripTitle,omitempty"`
	// 发起审批员工ID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 发起审批员工名
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s QueryCityCarApplyResponseBodyApplyList) String() string {
	return tea.Prettify(s)
}

func (s QueryCityCarApplyResponseBodyApplyList) GoString() string {
	return s.String()
}

func (s *QueryCityCarApplyResponseBodyApplyList) SetApproverList(v []*QueryCityCarApplyResponseBodyApplyListApproverList) *QueryCityCarApplyResponseBodyApplyList {
	s.ApproverList = v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyList) SetDepartId(v string) *QueryCityCarApplyResponseBodyApplyList {
	s.DepartId = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyList) SetDepartName(v string) *QueryCityCarApplyResponseBodyApplyList {
	s.DepartName = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyList) SetGmtCreate(v string) *QueryCityCarApplyResponseBodyApplyList {
	s.GmtCreate = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyList) SetGmtModified(v string) *QueryCityCarApplyResponseBodyApplyList {
	s.GmtModified = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyList) SetItineraryList(v []*QueryCityCarApplyResponseBodyApplyListItineraryList) *QueryCityCarApplyResponseBodyApplyList {
	s.ItineraryList = v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyList) SetStatus(v int64) *QueryCityCarApplyResponseBodyApplyList {
	s.Status = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyList) SetStatusDesc(v string) *QueryCityCarApplyResponseBodyApplyList {
	s.StatusDesc = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyList) SetThirdPartApplyId(v string) *QueryCityCarApplyResponseBodyApplyList {
	s.ThirdPartApplyId = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyList) SetTripCause(v string) *QueryCityCarApplyResponseBodyApplyList {
	s.TripCause = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyList) SetTripTitle(v string) *QueryCityCarApplyResponseBodyApplyList {
	s.TripTitle = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyList) SetUserId(v string) *QueryCityCarApplyResponseBodyApplyList {
	s.UserId = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyList) SetUserName(v string) *QueryCityCarApplyResponseBodyApplyList {
	s.UserName = &v
	return s
}

type QueryCityCarApplyResponseBodyApplyListApproverList struct {
	// 审批备注
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
	// 审批时间
	OperateTime *string `json:"operateTime,omitempty" xml:"operateTime,omitempty"`
	// 审批人排序值
	Order *int64 `json:"order,omitempty" xml:"order,omitempty"`
	// 审批状态枚举：审批状态：0-审批中，1-已同意，2-已拒绝
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 审批状态描述
	StatusDesc *string `json:"statusDesc,omitempty" xml:"statusDesc,omitempty"`
	// 审批员工ID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 审批员工名
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s QueryCityCarApplyResponseBodyApplyListApproverList) String() string {
	return tea.Prettify(s)
}

func (s QueryCityCarApplyResponseBodyApplyListApproverList) GoString() string {
	return s.String()
}

func (s *QueryCityCarApplyResponseBodyApplyListApproverList) SetNote(v string) *QueryCityCarApplyResponseBodyApplyListApproverList {
	s.Note = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyListApproverList) SetOperateTime(v string) *QueryCityCarApplyResponseBodyApplyListApproverList {
	s.OperateTime = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyListApproverList) SetOrder(v int64) *QueryCityCarApplyResponseBodyApplyListApproverList {
	s.Order = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyListApproverList) SetStatus(v int64) *QueryCityCarApplyResponseBodyApplyListApproverList {
	s.Status = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyListApproverList) SetStatusDesc(v string) *QueryCityCarApplyResponseBodyApplyListApproverList {
	s.StatusDesc = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyListApproverList) SetUserId(v string) *QueryCityCarApplyResponseBodyApplyListApproverList {
	s.UserId = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyListApproverList) SetUserName(v string) *QueryCityCarApplyResponseBodyApplyListApproverList {
	s.UserName = &v
	return s
}

type QueryCityCarApplyResponseBodyApplyListItineraryList struct {
	// 目的地城市
	ArrCity *string `json:"arrCity,omitempty" xml:"arrCity,omitempty"`
	// 目的地城市三字码
	ArrCityCode *string `json:"arrCityCode,omitempty" xml:"arrCityCode,omitempty"`
	// 到达目的地城市时间
	ArrDate *string `json:"arrDate,omitempty" xml:"arrDate,omitempty"`
	// 商旅内部成本中心ID
	CostCenterId *int64 `json:"costCenterId,omitempty" xml:"costCenterId,omitempty"`
	// 成本中心名称
	CostCenterName *string `json:"costCenterName,omitempty" xml:"costCenterName,omitempty"`
	// 出发城市
	DepCity *string `json:"depCity,omitempty" xml:"depCity,omitempty"`
	// 出发城市三字码
	DepCityCode *string `json:"depCityCode,omitempty" xml:"depCityCode,omitempty"`
	// 出发时间
	DepDate *string `json:"depDate,omitempty" xml:"depDate,omitempty"`
	// 商旅内部发票抬头ID
	InvoiceId *int64 `json:"invoiceId,omitempty" xml:"invoiceId,omitempty"`
	// 发票抬头名称
	InvoiceName *string `json:"invoiceName,omitempty" xml:"invoiceName,omitempty"`
	// 商旅内部行程单ID
	ItineraryId *string `json:"itineraryId,omitempty" xml:"itineraryId,omitempty"`
	// 项目code
	ProjectCode *string `json:"projectCode,omitempty" xml:"projectCode,omitempty"`
	// 项目名称
	ProjectTitle *string `json:"projectTitle,omitempty" xml:"projectTitle,omitempty"`
	// 交通方式：4-市内交通
	TrafficType *int64 `json:"trafficType,omitempty" xml:"trafficType,omitempty"`
}

func (s QueryCityCarApplyResponseBodyApplyListItineraryList) String() string {
	return tea.Prettify(s)
}

func (s QueryCityCarApplyResponseBodyApplyListItineraryList) GoString() string {
	return s.String()
}

func (s *QueryCityCarApplyResponseBodyApplyListItineraryList) SetArrCity(v string) *QueryCityCarApplyResponseBodyApplyListItineraryList {
	s.ArrCity = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyListItineraryList) SetArrCityCode(v string) *QueryCityCarApplyResponseBodyApplyListItineraryList {
	s.ArrCityCode = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyListItineraryList) SetArrDate(v string) *QueryCityCarApplyResponseBodyApplyListItineraryList {
	s.ArrDate = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyListItineraryList) SetCostCenterId(v int64) *QueryCityCarApplyResponseBodyApplyListItineraryList {
	s.CostCenterId = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyListItineraryList) SetCostCenterName(v string) *QueryCityCarApplyResponseBodyApplyListItineraryList {
	s.CostCenterName = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyListItineraryList) SetDepCity(v string) *QueryCityCarApplyResponseBodyApplyListItineraryList {
	s.DepCity = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyListItineraryList) SetDepCityCode(v string) *QueryCityCarApplyResponseBodyApplyListItineraryList {
	s.DepCityCode = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyListItineraryList) SetDepDate(v string) *QueryCityCarApplyResponseBodyApplyListItineraryList {
	s.DepDate = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyListItineraryList) SetInvoiceId(v int64) *QueryCityCarApplyResponseBodyApplyListItineraryList {
	s.InvoiceId = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyListItineraryList) SetInvoiceName(v string) *QueryCityCarApplyResponseBodyApplyListItineraryList {
	s.InvoiceName = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyListItineraryList) SetItineraryId(v string) *QueryCityCarApplyResponseBodyApplyListItineraryList {
	s.ItineraryId = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyListItineraryList) SetProjectCode(v string) *QueryCityCarApplyResponseBodyApplyListItineraryList {
	s.ProjectCode = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyListItineraryList) SetProjectTitle(v string) *QueryCityCarApplyResponseBodyApplyListItineraryList {
	s.ProjectTitle = &v
	return s
}

func (s *QueryCityCarApplyResponseBodyApplyListItineraryList) SetTrafficType(v int64) *QueryCityCarApplyResponseBodyApplyListItineraryList {
	s.TrafficType = &v
	return s
}

type QueryCityCarApplyResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCityCarApplyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCityCarApplyResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCityCarApplyResponse) GoString() string {
	return s.String()
}

func (s *QueryCityCarApplyResponse) SetHeaders(v map[string]*string) *QueryCityCarApplyResponse {
	s.Headers = v
	return s
}

func (s *QueryCityCarApplyResponse) SetBody(v *QueryCityCarApplyResponseBody) *QueryCityCarApplyResponse {
	s.Body = v
	return s
}

type GetTrainExceedApplyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTrainExceedApplyHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTrainExceedApplyHeaders) GoString() string {
	return s.String()
}

func (s *GetTrainExceedApplyHeaders) SetCommonHeaders(v map[string]*string) *GetTrainExceedApplyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTrainExceedApplyHeaders) SetXAcsDingtalkAccessToken(v string) *GetTrainExceedApplyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTrainExceedApplyRequest struct {
	// 第三方企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 商旅超标审批单id
	ApplyId *string `json:"applyId,omitempty" xml:"applyId,omitempty"`
}

func (s GetTrainExceedApplyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTrainExceedApplyRequest) GoString() string {
	return s.String()
}

func (s *GetTrainExceedApplyRequest) SetCorpId(v string) *GetTrainExceedApplyRequest {
	s.CorpId = &v
	return s
}

func (s *GetTrainExceedApplyRequest) SetApplyId(v string) *GetTrainExceedApplyRequest {
	s.ApplyId = &v
	return s
}

type GetTrainExceedApplyResponseBody struct {
	// 第三方企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 商旅超标审批单id
	ApplyId *int64 `json:"applyId,omitempty" xml:"applyId,omitempty"`
	// 审批单状态 0:审批中 1:已同意 2:已拒绝
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// 出差原因
	BtripCause *string `json:"btripCause,omitempty" xml:"btripCause,omitempty"`
	// 超标类型，32：坐席超标
	ExceedType *int32 `json:"exceedType,omitempty" xml:"exceedType,omitempty"`
	// 超标原因
	ExceedReason *string `json:"exceedReason,omitempty" xml:"exceedReason,omitempty"`
	// 原差旅标准
	OriginStandard *string `json:"originStandard,omitempty" xml:"originStandard,omitempty"`
	// 审批单提交时间
	SubmitTime *string `json:"submitTime,omitempty" xml:"submitTime,omitempty"`
	// 第三方用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 意向出行信息
	ApplyIntentionInfoDO *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO `json:"applyIntentionInfoDO,omitempty" xml:"applyIntentionInfoDO,omitempty" type:"Struct"`
	// 第三方出差审批单号
	ThirdpartApplyId *string `json:"thirdpartApplyId,omitempty" xml:"thirdpartApplyId,omitempty"`
}

func (s GetTrainExceedApplyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrainExceedApplyResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrainExceedApplyResponseBody) SetCorpId(v string) *GetTrainExceedApplyResponseBody {
	s.CorpId = &v
	return s
}

func (s *GetTrainExceedApplyResponseBody) SetApplyId(v int64) *GetTrainExceedApplyResponseBody {
	s.ApplyId = &v
	return s
}

func (s *GetTrainExceedApplyResponseBody) SetStatus(v int32) *GetTrainExceedApplyResponseBody {
	s.Status = &v
	return s
}

func (s *GetTrainExceedApplyResponseBody) SetBtripCause(v string) *GetTrainExceedApplyResponseBody {
	s.BtripCause = &v
	return s
}

func (s *GetTrainExceedApplyResponseBody) SetExceedType(v int32) *GetTrainExceedApplyResponseBody {
	s.ExceedType = &v
	return s
}

func (s *GetTrainExceedApplyResponseBody) SetExceedReason(v string) *GetTrainExceedApplyResponseBody {
	s.ExceedReason = &v
	return s
}

func (s *GetTrainExceedApplyResponseBody) SetOriginStandard(v string) *GetTrainExceedApplyResponseBody {
	s.OriginStandard = &v
	return s
}

func (s *GetTrainExceedApplyResponseBody) SetSubmitTime(v string) *GetTrainExceedApplyResponseBody {
	s.SubmitTime = &v
	return s
}

func (s *GetTrainExceedApplyResponseBody) SetUserId(v string) *GetTrainExceedApplyResponseBody {
	s.UserId = &v
	return s
}

func (s *GetTrainExceedApplyResponseBody) SetApplyIntentionInfoDO(v *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO) *GetTrainExceedApplyResponseBody {
	s.ApplyIntentionInfoDO = v
	return s
}

func (s *GetTrainExceedApplyResponseBody) SetThirdpartApplyId(v string) *GetTrainExceedApplyResponseBody {
	s.ThirdpartApplyId = &v
	return s
}

type GetTrainExceedApplyResponseBodyApplyIntentionInfoDO struct {
	// 意向坐席价格（分）
	Price *int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 出发城市名
	DepCityName *string `json:"depCityName,omitempty" xml:"depCityName,omitempty"`
	// 到达城市名
	ArrCityName *string `json:"arrCityName,omitempty" xml:"arrCityName,omitempty"`
	// 出发城市三字码
	DepCity *string `json:"depCity,omitempty" xml:"depCity,omitempty"`
	// 到达城市三字码
	ArrCity *string `json:"arrCity,omitempty" xml:"arrCity,omitempty"`
	// 出发时间
	DepTime *string `json:"depTime,omitempty" xml:"depTime,omitempty"`
	// 到达时间
	ArrTime *string `json:"arrTime,omitempty" xml:"arrTime,omitempty"`
	// 到达站点名称
	ArrStation *string `json:"arrStation,omitempty" xml:"arrStation,omitempty"`
	// 出发站点名称
	DepStation *string `json:"depStation,omitempty" xml:"depStation,omitempty"`
	// 意向车次号
	TrainNo *string `json:"trainNo,omitempty" xml:"trainNo,omitempty"`
	// 意向车次类型
	TrainTypeDesc *string `json:"trainTypeDesc,omitempty" xml:"trainTypeDesc,omitempty"`
	// 意向坐席名称
	SeatName *string `json:"seatName,omitempty" xml:"seatName,omitempty"`
}

func (s GetTrainExceedApplyResponseBodyApplyIntentionInfoDO) String() string {
	return tea.Prettify(s)
}

func (s GetTrainExceedApplyResponseBodyApplyIntentionInfoDO) GoString() string {
	return s.String()
}

func (s *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO) SetPrice(v int64) *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO {
	s.Price = &v
	return s
}

func (s *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO) SetDepCityName(v string) *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO {
	s.DepCityName = &v
	return s
}

func (s *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO) SetArrCityName(v string) *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO {
	s.ArrCityName = &v
	return s
}

func (s *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO) SetDepCity(v string) *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO {
	s.DepCity = &v
	return s
}

func (s *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO) SetArrCity(v string) *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO {
	s.ArrCity = &v
	return s
}

func (s *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO) SetDepTime(v string) *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO {
	s.DepTime = &v
	return s
}

func (s *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO) SetArrTime(v string) *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO {
	s.ArrTime = &v
	return s
}

func (s *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO) SetArrStation(v string) *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO {
	s.ArrStation = &v
	return s
}

func (s *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO) SetDepStation(v string) *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO {
	s.DepStation = &v
	return s
}

func (s *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO) SetTrainNo(v string) *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO {
	s.TrainNo = &v
	return s
}

func (s *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO) SetTrainTypeDesc(v string) *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO {
	s.TrainTypeDesc = &v
	return s
}

func (s *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO) SetSeatName(v string) *GetTrainExceedApplyResponseBodyApplyIntentionInfoDO {
	s.SeatName = &v
	return s
}

type GetTrainExceedApplyResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTrainExceedApplyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTrainExceedApplyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrainExceedApplyResponse) GoString() string {
	return s.String()
}

func (s *GetTrainExceedApplyResponse) SetHeaders(v map[string]*string) *GetTrainExceedApplyResponse {
	s.Headers = v
	return s
}

func (s *GetTrainExceedApplyResponse) SetBody(v *GetTrainExceedApplyResponseBody) *GetTrainExceedApplyResponse {
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

func (client *Client) ApproveCityCarApply(request *ApproveCityCarApplyRequest) (_result *ApproveCityCarApplyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ApproveCityCarApplyHeaders{}
	_result = &ApproveCityCarApplyResponse{}
	_body, _err := client.ApproveCityCarApplyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApproveCityCarApplyWithOptions(request *ApproveCityCarApplyRequest, headers *ApproveCityCarApplyHeaders, runtime *util.RuntimeOptions) (_result *ApproveCityCarApplyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.OperateTime)) {
		body["operateTime"] = request.OperateTime
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartApplyId)) {
		body["thirdPartApplyId"] = request.ThirdPartApplyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.DingCorpId)) {
		body["dingCorpId"] = request.DingCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
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
	_result = &ApproveCityCarApplyResponse{}
	_body, _err := client.DoROARequest(tea.String("ApproveCityCarApply"), tea.String("alitrip_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/alitrip/cityCarApprovals"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BillSettementHotel(request *BillSettementHotelRequest) (_result *BillSettementHotelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BillSettementHotelHeaders{}
	_result = &BillSettementHotelResponse{}
	_body, _err := client.BillSettementHotelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BillSettementHotelWithOptions(request *BillSettementHotelRequest, headers *BillSettementHotelHeaders, runtime *util.RuntimeOptions) (_result *BillSettementHotelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodStart)) {
		query["periodStart"] = request.PeriodStart
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodEnd)) {
		query["periodEnd"] = request.PeriodEnd
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
	_result = &BillSettementHotelResponse{}
	_body, _err := client.DoROARequest(tea.String("BillSettementHotel"), tea.String("alitrip_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/alitrip/billSettlements/hotels"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFlightExceedApply(request *GetFlightExceedApplyRequest) (_result *GetFlightExceedApplyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFlightExceedApplyHeaders{}
	_result = &GetFlightExceedApplyResponse{}
	_body, _err := client.GetFlightExceedApplyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFlightExceedApplyWithOptions(request *GetFlightExceedApplyRequest, headers *GetFlightExceedApplyHeaders, runtime *util.RuntimeOptions) (_result *GetFlightExceedApplyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.ApplyId)) {
		query["applyId"] = request.ApplyId
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
	_result = &GetFlightExceedApplyResponse{}
	_body, _err := client.DoROARequest(tea.String("GetFlightExceedApply"), tea.String("alitrip_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/alitrip/exceedapply/getFlight"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BillSettementCar(request *BillSettementCarRequest) (_result *BillSettementCarResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BillSettementCarHeaders{}
	_result = &BillSettementCarResponse{}
	_body, _err := client.BillSettementCarWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BillSettementCarWithOptions(request *BillSettementCarRequest, headers *BillSettementCarHeaders, runtime *util.RuntimeOptions) (_result *BillSettementCarResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodStart)) {
		query["periodStart"] = request.PeriodStart
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodEnd)) {
		query["periodEnd"] = request.PeriodEnd
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
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
	_result = &BillSettementCarResponse{}
	_body, _err := client.DoROARequest(tea.String("BillSettementCar"), tea.String("alitrip_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/alitrip/billSettlements/cars"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BillSettementBtripTrain(request *BillSettementBtripTrainRequest) (_result *BillSettementBtripTrainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BillSettementBtripTrainHeaders{}
	_result = &BillSettementBtripTrainResponse{}
	_body, _err := client.BillSettementBtripTrainWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BillSettementBtripTrainWithOptions(request *BillSettementBtripTrainRequest, headers *BillSettementBtripTrainHeaders, runtime *util.RuntimeOptions) (_result *BillSettementBtripTrainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodStart)) {
		query["periodStart"] = request.PeriodStart
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodEnd)) {
		query["periodEnd"] = request.PeriodEnd
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
	_result = &BillSettementBtripTrainResponse{}
	_body, _err := client.DoROARequest(tea.String("BillSettementBtripTrain"), tea.String("alitrip_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/alitrip/billSettlements/btripTrains"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncExceedApply(request *SyncExceedApplyRequest) (_result *SyncExceedApplyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncExceedApplyHeaders{}
	_result = &SyncExceedApplyResponse{}
	_body, _err := client.SyncExceedApplyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncExceedApplyWithOptions(request *SyncExceedApplyRequest, headers *SyncExceedApplyHeaders, runtime *util.RuntimeOptions) (_result *SyncExceedApplyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ApplyId)) {
		query["applyId"] = request.ApplyId
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdpartyFlowId)) {
		query["thirdpartyFlowId"] = request.ThirdpartyFlowId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
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
	_result = &SyncExceedApplyResponse{}
	_body, _err := client.DoROARequest(tea.String("SyncExceedApply"), tea.String("alitrip_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/alitrip/exceedapply/sync"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddCityCarApply(request *AddCityCarApplyRequest) (_result *AddCityCarApplyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddCityCarApplyHeaders{}
	_result = &AddCityCarApplyResponse{}
	_body, _err := client.AddCityCarApplyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddCityCarApplyWithOptions(request *AddCityCarApplyRequest, headers *AddCityCarApplyHeaders, runtime *util.RuntimeOptions) (_result *AddCityCarApplyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cause)) {
		body["cause"] = request.Cause
	}

	if !tea.BoolValue(util.IsUnset(request.City)) {
		body["city"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Date)) {
		body["date"] = request.Date
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectCode)) {
		body["projectCode"] = request.ProjectCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["projectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartApplyId)) {
		body["thirdPartApplyId"] = request.ThirdPartApplyId
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartCostCenterId)) {
		body["thirdPartCostCenterId"] = request.ThirdPartCostCenterId
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartInvoiceId)) {
		body["thirdPartInvoiceId"] = request.ThirdPartInvoiceId
	}

	if !tea.BoolValue(util.IsUnset(request.TimesTotal)) {
		body["timesTotal"] = request.TimesTotal
	}

	if !tea.BoolValue(util.IsUnset(request.TimesType)) {
		body["timesType"] = request.TimesType
	}

	if !tea.BoolValue(util.IsUnset(request.TimesUsed)) {
		body["timesUsed"] = request.TimesUsed
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.DingCorpId)) {
		body["dingCorpId"] = request.DingCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.FinishedDate)) {
		body["finishedDate"] = request.FinishedDate
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
	_result = &AddCityCarApplyResponse{}
	_body, _err := client.DoROARequest(tea.String("AddCityCarApply"), tea.String("alitrip_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/alitrip/cityCarApprovals"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BillSettementFlight(request *BillSettementFlightRequest) (_result *BillSettementFlightResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BillSettementFlightHeaders{}
	_result = &BillSettementFlightResponse{}
	_body, _err := client.BillSettementFlightWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BillSettementFlightWithOptions(request *BillSettementFlightRequest, headers *BillSettementFlightHeaders, runtime *util.RuntimeOptions) (_result *BillSettementFlightResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodStart)) {
		query["periodStart"] = request.PeriodStart
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodEnd)) {
		query["periodEnd"] = request.PeriodEnd
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
	_result = &BillSettementFlightResponse{}
	_body, _err := client.DoROARequest(tea.String("BillSettementFlight"), tea.String("alitrip_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/alitrip/billSettlements/flights"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotelExceedApply(request *GetHotelExceedApplyRequest) (_result *GetHotelExceedApplyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetHotelExceedApplyHeaders{}
	_result = &GetHotelExceedApplyResponse{}
	_body, _err := client.GetHotelExceedApplyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotelExceedApplyWithOptions(request *GetHotelExceedApplyRequest, headers *GetHotelExceedApplyHeaders, runtime *util.RuntimeOptions) (_result *GetHotelExceedApplyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.ApplyId)) {
		query["applyId"] = request.ApplyId
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
	_result = &GetHotelExceedApplyResponse{}
	_body, _err := client.DoROARequest(tea.String("GetHotelExceedApply"), tea.String("alitrip_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/alitrip/exceedapply/getHotel"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUnionOrder(request *QueryUnionOrderRequest) (_result *QueryUnionOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUnionOrderHeaders{}
	_result = &QueryUnionOrderResponse{}
	_body, _err := client.QueryUnionOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUnionOrderWithOptions(request *QueryUnionOrderRequest, headers *QueryUnionOrderHeaders, runtime *util.RuntimeOptions) (_result *QueryUnionOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartApplyId)) {
		query["thirdPartApplyId"] = request.ThirdPartApplyId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionNo)) {
		query["unionNo"] = request.UnionNo
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
	_result = &QueryUnionOrderResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryUnionOrder"), tea.String("alitrip_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/alitrip/unionOrders"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCityCarApply(request *QueryCityCarApplyRequest) (_result *QueryCityCarApplyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCityCarApplyHeaders{}
	_result = &QueryCityCarApplyResponse{}
	_body, _err := client.QueryCityCarApplyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCityCarApplyWithOptions(request *QueryCityCarApplyRequest, headers *QueryCityCarApplyHeaders, runtime *util.RuntimeOptions) (_result *QueryCityCarApplyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedEndAt)) {
		query["createdEndAt"] = request.CreatedEndAt
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedStartAt)) {
		query["createdStartAt"] = request.CreatedStartAt
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartApplyId)) {
		query["thirdPartApplyId"] = request.ThirdPartApplyId
	}

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
	_result = &QueryCityCarApplyResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryCityCarApply"), tea.String("alitrip_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/alitrip/cityCarApprovals"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTrainExceedApply(request *GetTrainExceedApplyRequest) (_result *GetTrainExceedApplyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTrainExceedApplyHeaders{}
	_result = &GetTrainExceedApplyResponse{}
	_body, _err := client.GetTrainExceedApplyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTrainExceedApplyWithOptions(request *GetTrainExceedApplyRequest, headers *GetTrainExceedApplyHeaders, runtime *util.RuntimeOptions) (_result *GetTrainExceedApplyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.ApplyId)) {
		query["applyId"] = request.ApplyId
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
	_result = &GetTrainExceedApplyResponse{}
	_body, _err := client.DoROARequest(tea.String("GetTrainExceedApply"), tea.String("alitrip_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/alitrip/exceedapply/getTrain"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
