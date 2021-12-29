// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package finance_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type NotifyPayCodePayResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s NotifyPayCodePayResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayCodePayResultHeaders) GoString() string {
	return s.String()
}

func (s *NotifyPayCodePayResultHeaders) SetCommonHeaders(v map[string]*string) *NotifyPayCodePayResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *NotifyPayCodePayResultHeaders) SetXAcsDingtalkAccessToken(v string) *NotifyPayCodePayResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type NotifyPayCodePayResultRequest struct {
	// 付款码值
	PayCode *string `json:"payCode,omitempty" xml:"payCode,omitempty"`
	// 企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 交易开始时间
	GmtTradeCreate *string `json:"gmtTradeCreate,omitempty" xml:"gmtTradeCreate,omitempty"`
	// 交易结束时间
	GmtTradeFinish *string `json:"gmtTradeFinish,omitempty" xml:"gmtTradeFinish,omitempty"`
	// 交易号
	TradeNo *string `json:"tradeNo,omitempty" xml:"tradeNo,omitempty"`
	// 交易状态
	TradeStatus *string `json:"tradeStatus,omitempty" xml:"tradeStatus,omitempty"`
	// 订单标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 订单金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 订单优惠金额
	PromotionAmount *string `json:"promotionAmount,omitempty" xml:"promotionAmount,omitempty"`
	// 收费金额
	ChargeAmount *string `json:"chargeAmount,omitempty" xml:"chargeAmount,omitempty"`
	// 支付渠道明细信息
	PayChannelDetailList []*NotifyPayCodePayResultRequestPayChannelDetailList `json:"payChannelDetailList,omitempty" xml:"payChannelDetailList,omitempty" type:"Repeated"`
	// 支付失败错误码
	TradeErrorCode *string `json:"tradeErrorCode,omitempty" xml:"tradeErrorCode,omitempty"`
	// 支付失败信息
	TradeErrorMsg *string `json:"tradeErrorMsg,omitempty" xml:"tradeErrorMsg,omitempty"`
	// 扩展信息
	ExtInfo *string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// isvOrgId
	DingIsvOrgId *int64 `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	// merchantName
	MerchantName *string `json:"merchantName,omitempty" xml:"merchantName,omitempty"`
}

func (s NotifyPayCodePayResultRequest) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayCodePayResultRequest) GoString() string {
	return s.String()
}

func (s *NotifyPayCodePayResultRequest) SetPayCode(v string) *NotifyPayCodePayResultRequest {
	s.PayCode = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetCorpId(v string) *NotifyPayCodePayResultRequest {
	s.CorpId = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetUserId(v string) *NotifyPayCodePayResultRequest {
	s.UserId = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetGmtTradeCreate(v string) *NotifyPayCodePayResultRequest {
	s.GmtTradeCreate = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetGmtTradeFinish(v string) *NotifyPayCodePayResultRequest {
	s.GmtTradeFinish = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetTradeNo(v string) *NotifyPayCodePayResultRequest {
	s.TradeNo = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetTradeStatus(v string) *NotifyPayCodePayResultRequest {
	s.TradeStatus = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetTitle(v string) *NotifyPayCodePayResultRequest {
	s.Title = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetRemark(v string) *NotifyPayCodePayResultRequest {
	s.Remark = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetAmount(v string) *NotifyPayCodePayResultRequest {
	s.Amount = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetPromotionAmount(v string) *NotifyPayCodePayResultRequest {
	s.PromotionAmount = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetChargeAmount(v string) *NotifyPayCodePayResultRequest {
	s.ChargeAmount = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetPayChannelDetailList(v []*NotifyPayCodePayResultRequestPayChannelDetailList) *NotifyPayCodePayResultRequest {
	s.PayChannelDetailList = v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetTradeErrorCode(v string) *NotifyPayCodePayResultRequest {
	s.TradeErrorCode = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetTradeErrorMsg(v string) *NotifyPayCodePayResultRequest {
	s.TradeErrorMsg = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetExtInfo(v string) *NotifyPayCodePayResultRequest {
	s.ExtInfo = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetDingIsvOrgId(v int64) *NotifyPayCodePayResultRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetMerchantName(v string) *NotifyPayCodePayResultRequest {
	s.MerchantName = &v
	return s
}

type NotifyPayCodePayResultRequestPayChannelDetailList struct {
	// 支付渠道名称
	PayChannelName *string `json:"payChannelName,omitempty" xml:"payChannelName,omitempty"`
	// 开始时间
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 结束时间
	GmtFinish *string `json:"gmtFinish,omitempty" xml:"gmtFinish,omitempty"`
	// 支付渠道类型
	PayChannelType *string `json:"payChannelType,omitempty" xml:"payChannelType,omitempty"`
	// 支付金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 支付渠道单号
	PayChannelOrderNo *string `json:"payChannelOrderNo,omitempty" xml:"payChannelOrderNo,omitempty"`
	// 优惠金额
	PromotionAmount *string `json:"promotionAmount,omitempty" xml:"promotionAmount,omitempty"`
	// 资金工具明细
	FundToolDetailList []*NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList `json:"fundToolDetailList,omitempty" xml:"fundToolDetailList,omitempty" type:"Repeated"`
}

func (s NotifyPayCodePayResultRequestPayChannelDetailList) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayCodePayResultRequestPayChannelDetailList) GoString() string {
	return s.String()
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailList) SetPayChannelName(v string) *NotifyPayCodePayResultRequestPayChannelDetailList {
	s.PayChannelName = &v
	return s
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailList) SetGmtCreate(v string) *NotifyPayCodePayResultRequestPayChannelDetailList {
	s.GmtCreate = &v
	return s
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailList) SetGmtFinish(v string) *NotifyPayCodePayResultRequestPayChannelDetailList {
	s.GmtFinish = &v
	return s
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailList) SetPayChannelType(v string) *NotifyPayCodePayResultRequestPayChannelDetailList {
	s.PayChannelType = &v
	return s
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailList) SetAmount(v string) *NotifyPayCodePayResultRequestPayChannelDetailList {
	s.Amount = &v
	return s
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailList) SetPayChannelOrderNo(v string) *NotifyPayCodePayResultRequestPayChannelDetailList {
	s.PayChannelOrderNo = &v
	return s
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailList) SetPromotionAmount(v string) *NotifyPayCodePayResultRequestPayChannelDetailList {
	s.PromotionAmount = &v
	return s
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailList) SetFundToolDetailList(v []*NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList) *NotifyPayCodePayResultRequestPayChannelDetailList {
	s.FundToolDetailList = v
	return s
}

type NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList struct {
	// 资金渠道名称
	FundToolName *string `json:"fundToolName,omitempty" xml:"fundToolName,omitempty"`
	// 1.00
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 开始时间
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 结束时间
	GmtFinish *string `json:"gmtFinish,omitempty" xml:"gmtFinish,omitempty"`
	// 是否是优惠工具
	PromotionFundTool *bool `json:"promotionFundTool,omitempty" xml:"promotionFundTool,omitempty"`
	// 扩展信息
	ExtInfo *string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
}

func (s NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList) GoString() string {
	return s.String()
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList) SetFundToolName(v string) *NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList {
	s.FundToolName = &v
	return s
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList) SetAmount(v string) *NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList {
	s.Amount = &v
	return s
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList) SetGmtCreate(v string) *NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList {
	s.GmtCreate = &v
	return s
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList) SetGmtFinish(v string) *NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList {
	s.GmtFinish = &v
	return s
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList) SetPromotionFundTool(v bool) *NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList {
	s.PromotionFundTool = &v
	return s
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList) SetExtInfo(v string) *NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList {
	s.ExtInfo = &v
	return s
}

type NotifyPayCodePayResultResponseBody struct {
	// 处理结果
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s NotifyPayCodePayResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayCodePayResultResponseBody) GoString() string {
	return s.String()
}

func (s *NotifyPayCodePayResultResponseBody) SetResult(v string) *NotifyPayCodePayResultResponseBody {
	s.Result = &v
	return s
}

type NotifyPayCodePayResultResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *NotifyPayCodePayResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s NotifyPayCodePayResultResponse) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayCodePayResultResponse) GoString() string {
	return s.String()
}

func (s *NotifyPayCodePayResultResponse) SetHeaders(v map[string]*string) *NotifyPayCodePayResultResponse {
	s.Headers = v
	return s
}

func (s *NotifyPayCodePayResultResponse) SetBody(v *NotifyPayCodePayResultResponseBody) *NotifyPayCodePayResultResponse {
	s.Body = v
	return s
}

type UpateUserCodeInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpateUserCodeInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpateUserCodeInstanceHeaders) GoString() string {
	return s.String()
}

func (s *UpateUserCodeInstanceHeaders) SetCommonHeaders(v map[string]*string) *UpateUserCodeInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpateUserCodeInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *UpateUserCodeInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpateUserCodeInstanceRequest struct {
	// 用户码ID
	CodeId *string `json:"codeId,omitempty" xml:"codeId,omitempty"`
	// 码标识
	CodeIdentity *string `json:"codeIdentity,omitempty" xml:"codeIdentity,omitempty"`
	// 码值
	CodeValue *string `json:"codeValue,omitempty" xml:"codeValue,omitempty"`
	// 状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 企业ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 用户和企业的关系类型，区分内部员工，外部联系人，无关系普通用户
	UserCorpRelationType *string `json:"userCorpRelationType,omitempty" xml:"userCorpRelationType,omitempty"`
	// 用户身份标识，取值和用户企业关系类型相关，如果企业无关，传入手机号
	UserIdentity *string `json:"userIdentity,omitempty" xml:"userIdentity,omitempty"`
	// 临时码，传入过期时间
	GmtExpired *string `json:"gmtExpired,omitempty" xml:"gmtExpired,omitempty"`
	// 有效时间列表，对于连续时间段，只需传入一个对象即可，注意过期时间必须晚于最晚结束时间
	AvailableTimes []*UpateUserCodeInstanceRequestAvailableTimes `json:"availableTimes,omitempty" xml:"availableTimes,omitempty" type:"Repeated"`
	// 扩展参数
	ExtInfo      map[string]interface{} `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	DingOrgId    *int64                 `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingIsvOrgId *int64                 `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
}

func (s UpateUserCodeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpateUserCodeInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpateUserCodeInstanceRequest) SetCodeId(v string) *UpateUserCodeInstanceRequest {
	s.CodeId = &v
	return s
}

func (s *UpateUserCodeInstanceRequest) SetCodeIdentity(v string) *UpateUserCodeInstanceRequest {
	s.CodeIdentity = &v
	return s
}

func (s *UpateUserCodeInstanceRequest) SetCodeValue(v string) *UpateUserCodeInstanceRequest {
	s.CodeValue = &v
	return s
}

func (s *UpateUserCodeInstanceRequest) SetStatus(v string) *UpateUserCodeInstanceRequest {
	s.Status = &v
	return s
}

func (s *UpateUserCodeInstanceRequest) SetCorpId(v string) *UpateUserCodeInstanceRequest {
	s.CorpId = &v
	return s
}

func (s *UpateUserCodeInstanceRequest) SetUserCorpRelationType(v string) *UpateUserCodeInstanceRequest {
	s.UserCorpRelationType = &v
	return s
}

func (s *UpateUserCodeInstanceRequest) SetUserIdentity(v string) *UpateUserCodeInstanceRequest {
	s.UserIdentity = &v
	return s
}

func (s *UpateUserCodeInstanceRequest) SetGmtExpired(v string) *UpateUserCodeInstanceRequest {
	s.GmtExpired = &v
	return s
}

func (s *UpateUserCodeInstanceRequest) SetAvailableTimes(v []*UpateUserCodeInstanceRequestAvailableTimes) *UpateUserCodeInstanceRequest {
	s.AvailableTimes = v
	return s
}

func (s *UpateUserCodeInstanceRequest) SetExtInfo(v map[string]interface{}) *UpateUserCodeInstanceRequest {
	s.ExtInfo = v
	return s
}

func (s *UpateUserCodeInstanceRequest) SetDingOrgId(v int64) *UpateUserCodeInstanceRequest {
	s.DingOrgId = &v
	return s
}

func (s *UpateUserCodeInstanceRequest) SetDingIsvOrgId(v int64) *UpateUserCodeInstanceRequest {
	s.DingIsvOrgId = &v
	return s
}

type UpateUserCodeInstanceRequestAvailableTimes struct {
	// 开始时间
	GmtStart *string `json:"gmtStart,omitempty" xml:"gmtStart,omitempty"`
	// 结束时间
	GmtEnd *string `json:"gmtEnd,omitempty" xml:"gmtEnd,omitempty"`
}

func (s UpateUserCodeInstanceRequestAvailableTimes) String() string {
	return tea.Prettify(s)
}

func (s UpateUserCodeInstanceRequestAvailableTimes) GoString() string {
	return s.String()
}

func (s *UpateUserCodeInstanceRequestAvailableTimes) SetGmtStart(v string) *UpateUserCodeInstanceRequestAvailableTimes {
	s.GmtStart = &v
	return s
}

func (s *UpateUserCodeInstanceRequestAvailableTimes) SetGmtEnd(v string) *UpateUserCodeInstanceRequestAvailableTimes {
	s.GmtEnd = &v
	return s
}

type UpateUserCodeInstanceResponseBody struct {
	// 码ID
	CodeId *string `json:"codeId,omitempty" xml:"codeId,omitempty"`
}

func (s UpateUserCodeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpateUserCodeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpateUserCodeInstanceResponseBody) SetCodeId(v string) *UpateUserCodeInstanceResponseBody {
	s.CodeId = &v
	return s
}

type UpateUserCodeInstanceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpateUserCodeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpateUserCodeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpateUserCodeInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpateUserCodeInstanceResponse) SetHeaders(v map[string]*string) *UpateUserCodeInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpateUserCodeInstanceResponse) SetBody(v *UpateUserCodeInstanceResponseBody) *UpateUserCodeInstanceResponse {
	s.Body = v
	return s
}

type ApplyBatchPayHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ApplyBatchPayHeaders) String() string {
	return tea.Prettify(s)
}

func (s ApplyBatchPayHeaders) GoString() string {
	return s.String()
}

func (s *ApplyBatchPayHeaders) SetCommonHeaders(v map[string]*string) *ApplyBatchPayHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ApplyBatchPayHeaders) SetXAcsDingtalkAccessToken(v string) *ApplyBatchPayHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ApplyBatchPayRequest struct {
	// 支付发起人staffId
	StaffId *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
	// 支付账号唯一id
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// 钉钉订单号(和商户批次号一一对应)
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	// 订单总金额（必填）, 单位为：元
	TransAmount *string `json:"transAmount,omitempty" xml:"transAmount,omitempty"`
	// 回调url
	ReturnUrl *string `json:"returnUrl,omitempty" xml:"returnUrl,omitempty"`
	// 公用回传参数，如果请求时传递了该参数，则异步通知商户时会回传该参数
	PassBackParams map[string]interface{} `json:"passBackParams,omitempty" xml:"passBackParams,omitempty"`
	// 支付终端
	PayTerminal *string `json:"payTerminal,omitempty" xml:"payTerminal,omitempty"`
	// 转账过期时间
	TransExpireTime *string `json:"transExpireTime,omitempty" xml:"transExpireTime,omitempty"`
}

func (s ApplyBatchPayRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyBatchPayRequest) GoString() string {
	return s.String()
}

func (s *ApplyBatchPayRequest) SetStaffId(v string) *ApplyBatchPayRequest {
	s.StaffId = &v
	return s
}

func (s *ApplyBatchPayRequest) SetAccountId(v string) *ApplyBatchPayRequest {
	s.AccountId = &v
	return s
}

func (s *ApplyBatchPayRequest) SetOrderNo(v string) *ApplyBatchPayRequest {
	s.OrderNo = &v
	return s
}

func (s *ApplyBatchPayRequest) SetTransAmount(v string) *ApplyBatchPayRequest {
	s.TransAmount = &v
	return s
}

func (s *ApplyBatchPayRequest) SetReturnUrl(v string) *ApplyBatchPayRequest {
	s.ReturnUrl = &v
	return s
}

func (s *ApplyBatchPayRequest) SetPassBackParams(v map[string]interface{}) *ApplyBatchPayRequest {
	s.PassBackParams = v
	return s
}

func (s *ApplyBatchPayRequest) SetPayTerminal(v string) *ApplyBatchPayRequest {
	s.PayTerminal = &v
	return s
}

func (s *ApplyBatchPayRequest) SetTransExpireTime(v string) *ApplyBatchPayRequest {
	s.TransExpireTime = &v
	return s
}

type ApplyBatchPayResponseBody struct {
	// 支付确认页数据
	PayData *string `json:"payData,omitempty" xml:"payData,omitempty"`
	// 钉钉支付的批次号
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
}

func (s ApplyBatchPayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyBatchPayResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyBatchPayResponseBody) SetPayData(v string) *ApplyBatchPayResponseBody {
	s.PayData = &v
	return s
}

func (s *ApplyBatchPayResponseBody) SetOrderNo(v string) *ApplyBatchPayResponseBody {
	s.OrderNo = &v
	return s
}

type ApplyBatchPayResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApplyBatchPayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyBatchPayResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyBatchPayResponse) GoString() string {
	return s.String()
}

func (s *ApplyBatchPayResponse) SetHeaders(v map[string]*string) *ApplyBatchPayResponse {
	s.Headers = v
	return s
}

func (s *ApplyBatchPayResponse) SetBody(v *ApplyBatchPayResponseBody) *ApplyBatchPayResponse {
	s.Body = v
	return s
}

type QueryUserAlipayAccountHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserAlipayAccountHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserAlipayAccountHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserAlipayAccountHeaders) SetCommonHeaders(v map[string]*string) *QueryUserAlipayAccountHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserAlipayAccountHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserAlipayAccountHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserAlipayAccountResponseBody struct {
	// 支付宝uid
	AlipayUid *string `json:"alipayUid,omitempty" xml:"alipayUid,omitempty"`
}

func (s QueryUserAlipayAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserAlipayAccountResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserAlipayAccountResponseBody) SetAlipayUid(v string) *QueryUserAlipayAccountResponseBody {
	s.AlipayUid = &v
	return s
}

type QueryUserAlipayAccountResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUserAlipayAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserAlipayAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserAlipayAccountResponse) GoString() string {
	return s.String()
}

func (s *QueryUserAlipayAccountResponse) SetHeaders(v map[string]*string) *QueryUserAlipayAccountResponse {
	s.Headers = v
	return s
}

func (s *QueryUserAlipayAccountResponse) SetBody(v *QueryUserAlipayAccountResponseBody) *QueryUserAlipayAccountResponse {
	s.Body = v
	return s
}

type DecodePayCodeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DecodePayCodeHeaders) String() string {
	return tea.Prettify(s)
}

func (s DecodePayCodeHeaders) GoString() string {
	return s.String()
}

func (s *DecodePayCodeHeaders) SetCommonHeaders(v map[string]*string) *DecodePayCodeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DecodePayCodeHeaders) SetXAcsDingtalkAccessToken(v string) *DecodePayCodeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DecodePayCodeRequest struct {
	// payCode
	PayCode *string `json:"payCode,omitempty" xml:"payCode,omitempty"`
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// ISV组织ID
	DingIsvOrgId *int64 `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	// 组织ID
	DingOrgId *int64 `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
}

func (s DecodePayCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s DecodePayCodeRequest) GoString() string {
	return s.String()
}

func (s *DecodePayCodeRequest) SetPayCode(v string) *DecodePayCodeRequest {
	s.PayCode = &v
	return s
}

func (s *DecodePayCodeRequest) SetRequestId(v string) *DecodePayCodeRequest {
	s.RequestId = &v
	return s
}

func (s *DecodePayCodeRequest) SetDingIsvOrgId(v int64) *DecodePayCodeRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *DecodePayCodeRequest) SetDingOrgId(v int64) *DecodePayCodeRequest {
	s.DingOrgId = &v
	return s
}

type DecodePayCodeResponseBody struct {
	// 企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 员工id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户是否还在组织内
	UserInCorp *bool `json:"userInCorp,omitempty" xml:"userInCorp,omitempty"`
	// 码类型
	CodeType *string `json:"codeType,omitempty" xml:"codeType,omitempty"`
	// 支付宝付款码
	AlipayCode *string `json:"alipayCode,omitempty" xml:"alipayCode,omitempty"`
	// 用户和企业关系
	UserCorpRelationType *string `json:"userCorpRelationType,omitempty" xml:"userCorpRelationType,omitempty"`
	// 工牌码：DT_IDENTITY，访客码：DT_VISITOR，会展码：DT_CONFERENCE
	CodeIdentity *string `json:"codeIdentity,omitempty" xml:"codeIdentity,omitempty"`
	// 码ID，对于访客或会展码等静态码值返回
	CodeId *string `json:"codeId,omitempty" xml:"codeId,omitempty"`
	// 外部业务ID,其值为调用创建用户码接口传入的requestId
	OutBizId *string `json:"outBizId,omitempty" xml:"outBizId,omitempty"`
}

func (s DecodePayCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DecodePayCodeResponseBody) GoString() string {
	return s.String()
}

func (s *DecodePayCodeResponseBody) SetCorpId(v string) *DecodePayCodeResponseBody {
	s.CorpId = &v
	return s
}

func (s *DecodePayCodeResponseBody) SetUserId(v string) *DecodePayCodeResponseBody {
	s.UserId = &v
	return s
}

func (s *DecodePayCodeResponseBody) SetUserInCorp(v bool) *DecodePayCodeResponseBody {
	s.UserInCorp = &v
	return s
}

func (s *DecodePayCodeResponseBody) SetCodeType(v string) *DecodePayCodeResponseBody {
	s.CodeType = &v
	return s
}

func (s *DecodePayCodeResponseBody) SetAlipayCode(v string) *DecodePayCodeResponseBody {
	s.AlipayCode = &v
	return s
}

func (s *DecodePayCodeResponseBody) SetUserCorpRelationType(v string) *DecodePayCodeResponseBody {
	s.UserCorpRelationType = &v
	return s
}

func (s *DecodePayCodeResponseBody) SetCodeIdentity(v string) *DecodePayCodeResponseBody {
	s.CodeIdentity = &v
	return s
}

func (s *DecodePayCodeResponseBody) SetCodeId(v string) *DecodePayCodeResponseBody {
	s.CodeId = &v
	return s
}

func (s *DecodePayCodeResponseBody) SetOutBizId(v string) *DecodePayCodeResponseBody {
	s.OutBizId = &v
	return s
}

type DecodePayCodeResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DecodePayCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DecodePayCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DecodePayCodeResponse) GoString() string {
	return s.String()
}

func (s *DecodePayCodeResponse) SetHeaders(v map[string]*string) *DecodePayCodeResponse {
	s.Headers = v
	return s
}

func (s *DecodePayCodeResponse) SetBody(v *DecodePayCodeResponseBody) *DecodePayCodeResponse {
	s.Body = v
	return s
}

type CreateBatchTradeOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateBatchTradeOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchTradeOrderHeaders) GoString() string {
	return s.String()
}

func (s *CreateBatchTradeOrderHeaders) SetCommonHeaders(v map[string]*string) *CreateBatchTradeOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateBatchTradeOrderHeaders) SetXAcsDingtalkAccessToken(v string) *CreateBatchTradeOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateBatchTradeOrderRequest struct {
	// 员工staffId
	StaffId *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
	// 付款账号唯一id
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// 付款账号(注意：用户上送的是脱敏数据)
	AccountNo *string `json:"accountNo,omitempty" xml:"accountNo,omitempty"`
	// 交易抬头
	TradeTitle *string `json:"tradeTitle,omitempty" xml:"tradeTitle,omitempty"`
	// 外部商户批次号
	OutBatchNo *string `json:"outBatchNo,omitempty" xml:"outBatchNo,omitempty"`
	// 批次备注
	BatchRemark *string `json:"batchRemark,omitempty" xml:"batchRemark,omitempty"`
	// 总笔数（必填）
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 总金额（必填，单位：元）
	TotalAmount *string `json:"totalAmount,omitempty" xml:"totalAmount,omitempty"`
	// 交易明细列表
	BatchTradeDetails []*CreateBatchTradeOrderRequestBatchTradeDetails `json:"batchTradeDetails,omitempty" xml:"batchTradeDetails,omitempty" type:"Repeated"`
}

func (s CreateBatchTradeOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchTradeOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateBatchTradeOrderRequest) SetStaffId(v string) *CreateBatchTradeOrderRequest {
	s.StaffId = &v
	return s
}

func (s *CreateBatchTradeOrderRequest) SetAccountId(v string) *CreateBatchTradeOrderRequest {
	s.AccountId = &v
	return s
}

func (s *CreateBatchTradeOrderRequest) SetAccountNo(v string) *CreateBatchTradeOrderRequest {
	s.AccountNo = &v
	return s
}

func (s *CreateBatchTradeOrderRequest) SetTradeTitle(v string) *CreateBatchTradeOrderRequest {
	s.TradeTitle = &v
	return s
}

func (s *CreateBatchTradeOrderRequest) SetOutBatchNo(v string) *CreateBatchTradeOrderRequest {
	s.OutBatchNo = &v
	return s
}

func (s *CreateBatchTradeOrderRequest) SetBatchRemark(v string) *CreateBatchTradeOrderRequest {
	s.BatchRemark = &v
	return s
}

func (s *CreateBatchTradeOrderRequest) SetTotalCount(v int64) *CreateBatchTradeOrderRequest {
	s.TotalCount = &v
	return s
}

func (s *CreateBatchTradeOrderRequest) SetTotalAmount(v string) *CreateBatchTradeOrderRequest {
	s.TotalAmount = &v
	return s
}

func (s *CreateBatchTradeOrderRequest) SetBatchTradeDetails(v []*CreateBatchTradeOrderRequestBatchTradeDetails) *CreateBatchTradeOrderRequest {
	s.BatchTradeDetails = v
	return s
}

type CreateBatchTradeOrderRequestBatchTradeDetails struct {
	// 序号（必填）
	SerialNo *int64 `json:"serialNo,omitempty" xml:"serialNo,omitempty"`
	// 金额（必填，单位：元）
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 收款方户名（必填）
	PayeeAccountName *string `json:"payeeAccountName,omitempty" xml:"payeeAccountName,omitempty"`
	// 收款方账号（必填）
	PayeeAccountNo *string `json:"payeeAccountNo,omitempty" xml:"payeeAccountNo,omitempty"`
	// 收款方账号类型（必填）
	PayeeAccountType *string `json:"payeeAccountType,omitempty" xml:"payeeAccountType,omitempty"`
	// 备注（选填）
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
}

func (s CreateBatchTradeOrderRequestBatchTradeDetails) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchTradeOrderRequestBatchTradeDetails) GoString() string {
	return s.String()
}

func (s *CreateBatchTradeOrderRequestBatchTradeDetails) SetSerialNo(v int64) *CreateBatchTradeOrderRequestBatchTradeDetails {
	s.SerialNo = &v
	return s
}

func (s *CreateBatchTradeOrderRequestBatchTradeDetails) SetAmount(v string) *CreateBatchTradeOrderRequestBatchTradeDetails {
	s.Amount = &v
	return s
}

func (s *CreateBatchTradeOrderRequestBatchTradeDetails) SetPayeeAccountName(v string) *CreateBatchTradeOrderRequestBatchTradeDetails {
	s.PayeeAccountName = &v
	return s
}

func (s *CreateBatchTradeOrderRequestBatchTradeDetails) SetPayeeAccountNo(v string) *CreateBatchTradeOrderRequestBatchTradeDetails {
	s.PayeeAccountNo = &v
	return s
}

func (s *CreateBatchTradeOrderRequestBatchTradeDetails) SetPayeeAccountType(v string) *CreateBatchTradeOrderRequestBatchTradeDetails {
	s.PayeeAccountType = &v
	return s
}

func (s *CreateBatchTradeOrderRequestBatchTradeDetails) SetMemo(v string) *CreateBatchTradeOrderRequestBatchTradeDetails {
	s.Memo = &v
	return s
}

type CreateBatchTradeOrderResponseBody struct {
	// 钉钉批次单号
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	// 商户批次号
	OutBatchNo *string `json:"outBatchNo,omitempty" xml:"outBatchNo,omitempty"`
	// 批次订单状态
	OrderStatus *string `json:"orderStatus,omitempty" xml:"orderStatus,omitempty"`
}

func (s CreateBatchTradeOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchTradeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBatchTradeOrderResponseBody) SetOrderNo(v string) *CreateBatchTradeOrderResponseBody {
	s.OrderNo = &v
	return s
}

func (s *CreateBatchTradeOrderResponseBody) SetOutBatchNo(v string) *CreateBatchTradeOrderResponseBody {
	s.OutBatchNo = &v
	return s
}

func (s *CreateBatchTradeOrderResponseBody) SetOrderStatus(v string) *CreateBatchTradeOrderResponseBody {
	s.OrderStatus = &v
	return s
}

type CreateBatchTradeOrderResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateBatchTradeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBatchTradeOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchTradeOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateBatchTradeOrderResponse) SetHeaders(v map[string]*string) *CreateBatchTradeOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateBatchTradeOrderResponse) SetBody(v *CreateBatchTradeOrderResponseBody) *CreateBatchTradeOrderResponse {
	s.Body = v
	return s
}

type NotifyPayCodeRefundResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s NotifyPayCodeRefundResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayCodeRefundResultHeaders) GoString() string {
	return s.String()
}

func (s *NotifyPayCodeRefundResultHeaders) SetCommonHeaders(v map[string]*string) *NotifyPayCodeRefundResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *NotifyPayCodeRefundResultHeaders) SetXAcsDingtalkAccessToken(v string) *NotifyPayCodeRefundResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type NotifyPayCodeRefundResultRequest struct {
	// 企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 交易订单号
	TradeNo *string `json:"tradeNo,omitempty" xml:"tradeNo,omitempty"`
	// 本次退款订单号
	RefundOrderNo *string `json:"refundOrderNo,omitempty" xml:"refundOrderNo,omitempty"`
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 退款金额
	RefundAmount *string `json:"refundAmount,omitempty" xml:"refundAmount,omitempty"`
	// 退款的优惠金额
	RefundPromotionAmount *string `json:"refundPromotionAmount,omitempty" xml:"refundPromotionAmount,omitempty"`
	// 退款时间
	GmtRefund *string `json:"gmtRefund,omitempty" xml:"gmtRefund,omitempty"`
	// 支付渠道信息
	PayChannelDetailList []*NotifyPayCodeRefundResultRequestPayChannelDetailList `json:"payChannelDetailList,omitempty" xml:"payChannelDetailList,omitempty" type:"Repeated"`
	// isvOrgId
	DingIsvOrgId *int64 `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	// 支付时使用的付款码
	PayCode *string `json:"payCode,omitempty" xml:"payCode,omitempty"`
}

func (s NotifyPayCodeRefundResultRequest) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayCodeRefundResultRequest) GoString() string {
	return s.String()
}

func (s *NotifyPayCodeRefundResultRequest) SetCorpId(v string) *NotifyPayCodeRefundResultRequest {
	s.CorpId = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequest) SetUserId(v string) *NotifyPayCodeRefundResultRequest {
	s.UserId = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequest) SetTradeNo(v string) *NotifyPayCodeRefundResultRequest {
	s.TradeNo = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequest) SetRefundOrderNo(v string) *NotifyPayCodeRefundResultRequest {
	s.RefundOrderNo = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequest) SetRemark(v string) *NotifyPayCodeRefundResultRequest {
	s.Remark = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequest) SetRefundAmount(v string) *NotifyPayCodeRefundResultRequest {
	s.RefundAmount = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequest) SetRefundPromotionAmount(v string) *NotifyPayCodeRefundResultRequest {
	s.RefundPromotionAmount = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequest) SetGmtRefund(v string) *NotifyPayCodeRefundResultRequest {
	s.GmtRefund = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequest) SetPayChannelDetailList(v []*NotifyPayCodeRefundResultRequestPayChannelDetailList) *NotifyPayCodeRefundResultRequest {
	s.PayChannelDetailList = v
	return s
}

func (s *NotifyPayCodeRefundResultRequest) SetDingIsvOrgId(v int64) *NotifyPayCodeRefundResultRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequest) SetPayCode(v string) *NotifyPayCodeRefundResultRequest {
	s.PayCode = &v
	return s
}

type NotifyPayCodeRefundResultRequestPayChannelDetailList struct {
	// 支付渠道名称
	PayChannelName *string `json:"payChannelName,omitempty" xml:"payChannelName,omitempty"`
	// 支付渠道类型
	PayChannelType *string `json:"payChannelType,omitempty" xml:"payChannelType,omitempty"`
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 支付渠道号
	PayChannelOrderNo *string `json:"payChannelOrderNo,omitempty" xml:"payChannelOrderNo,omitempty"`
	// 支付渠道退款号
	PayChannelRefundOrderNo *string `json:"payChannelRefundOrderNo,omitempty" xml:"payChannelRefundOrderNo,omitempty"`
	// 优惠金额
	PromotionAmount *string `json:"promotionAmount,omitempty" xml:"promotionAmount,omitempty"`
	// 支付资金列表
	FundToolDetailList []*NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList `json:"fundToolDetailList,omitempty" xml:"fundToolDetailList,omitempty" type:"Repeated"`
}

func (s NotifyPayCodeRefundResultRequestPayChannelDetailList) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayCodeRefundResultRequestPayChannelDetailList) GoString() string {
	return s.String()
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailList) SetPayChannelName(v string) *NotifyPayCodeRefundResultRequestPayChannelDetailList {
	s.PayChannelName = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailList) SetPayChannelType(v string) *NotifyPayCodeRefundResultRequestPayChannelDetailList {
	s.PayChannelType = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailList) SetAmount(v string) *NotifyPayCodeRefundResultRequestPayChannelDetailList {
	s.Amount = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailList) SetPayChannelOrderNo(v string) *NotifyPayCodeRefundResultRequestPayChannelDetailList {
	s.PayChannelOrderNo = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailList) SetPayChannelRefundOrderNo(v string) *NotifyPayCodeRefundResultRequestPayChannelDetailList {
	s.PayChannelRefundOrderNo = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailList) SetPromotionAmount(v string) *NotifyPayCodeRefundResultRequestPayChannelDetailList {
	s.PromotionAmount = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailList) SetFundToolDetailList(v []*NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList) *NotifyPayCodeRefundResultRequestPayChannelDetailList {
	s.FundToolDetailList = v
	return s
}

type NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList struct {
	// 资金工具名称
	FundToolName *string `json:"fundToolName,omitempty" xml:"fundToolName,omitempty"`
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 创建时间
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 完成时间
	GmtFinish *string `json:"gmtFinish,omitempty" xml:"gmtFinish,omitempty"`
	// 是否是优惠工具
	PromotionFundTool *bool `json:"promotionFundTool,omitempty" xml:"promotionFundTool,omitempty"`
	// 扩展信息
	ExtInfo *string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
}

func (s NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList) GoString() string {
	return s.String()
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList) SetFundToolName(v string) *NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList {
	s.FundToolName = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList) SetAmount(v string) *NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList {
	s.Amount = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList) SetGmtCreate(v string) *NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList {
	s.GmtCreate = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList) SetGmtFinish(v string) *NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList {
	s.GmtFinish = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList) SetPromotionFundTool(v bool) *NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList {
	s.PromotionFundTool = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList) SetExtInfo(v string) *NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList {
	s.ExtInfo = &v
	return s
}

type NotifyPayCodeRefundResultResponseBody struct {
	// 处理结果
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s NotifyPayCodeRefundResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayCodeRefundResultResponseBody) GoString() string {
	return s.String()
}

func (s *NotifyPayCodeRefundResultResponseBody) SetResult(v string) *NotifyPayCodeRefundResultResponseBody {
	s.Result = &v
	return s
}

type NotifyPayCodeRefundResultResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *NotifyPayCodeRefundResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s NotifyPayCodeRefundResultResponse) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayCodeRefundResultResponse) GoString() string {
	return s.String()
}

func (s *NotifyPayCodeRefundResultResponse) SetHeaders(v map[string]*string) *NotifyPayCodeRefundResultResponse {
	s.Headers = v
	return s
}

func (s *NotifyPayCodeRefundResultResponse) SetBody(v *NotifyPayCodeRefundResultResponseBody) *NotifyPayCodeRefundResultResponse {
	s.Body = v
	return s
}

type QueryBatchTradeDetailListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryBatchTradeDetailListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchTradeDetailListHeaders) GoString() string {
	return s.String()
}

func (s *QueryBatchTradeDetailListHeaders) SetCommonHeaders(v map[string]*string) *QueryBatchTradeDetailListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryBatchTradeDetailListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryBatchTradeDetailListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryBatchTradeDetailListRequest struct {
	// 外部商户批次号
	OutBatchNo *string `json:"outBatchNo,omitempty" xml:"outBatchNo,omitempty"`
	// 当前页数
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页记录数
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryBatchTradeDetailListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchTradeDetailListRequest) GoString() string {
	return s.String()
}

func (s *QueryBatchTradeDetailListRequest) SetOutBatchNo(v string) *QueryBatchTradeDetailListRequest {
	s.OutBatchNo = &v
	return s
}

func (s *QueryBatchTradeDetailListRequest) SetPageNumber(v int32) *QueryBatchTradeDetailListRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryBatchTradeDetailListRequest) SetPageSize(v int32) *QueryBatchTradeDetailListRequest {
	s.PageSize = &v
	return s
}

type QueryBatchTradeDetailListResponseBody struct {
	// 总记录数
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
	// 单页条数
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 总页数
	TotalPageNumber *int32 `json:"totalPageNumber,omitempty" xml:"totalPageNumber,omitempty"`
	// 当前页数
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 明细列表
	BatchTradeDetailList []*QueryBatchTradeDetailListResponseBodyBatchTradeDetailList `json:"batchTradeDetailList,omitempty" xml:"batchTradeDetailList,omitempty" type:"Repeated"`
}

func (s QueryBatchTradeDetailListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchTradeDetailListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBatchTradeDetailListResponseBody) SetTotal(v int32) *QueryBatchTradeDetailListResponseBody {
	s.Total = &v
	return s
}

func (s *QueryBatchTradeDetailListResponseBody) SetPageSize(v int32) *QueryBatchTradeDetailListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryBatchTradeDetailListResponseBody) SetTotalPageNumber(v int32) *QueryBatchTradeDetailListResponseBody {
	s.TotalPageNumber = &v
	return s
}

func (s *QueryBatchTradeDetailListResponseBody) SetPageNumber(v int32) *QueryBatchTradeDetailListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryBatchTradeDetailListResponseBody) SetBatchTradeDetailList(v []*QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) *QueryBatchTradeDetailListResponseBody {
	s.BatchTradeDetailList = v
	return s
}

type QueryBatchTradeDetailListResponseBodyBatchTradeDetailList struct {
	// 序号
	SerialNo *int64 `json:"serialNo,omitempty" xml:"serialNo,omitempty"`
	// 明细单号
	DetailNo *string `json:"detailNo,omitempty" xml:"detailNo,omitempty"`
	// 收款人账号
	PayeeAccountNo *string `json:"payeeAccountNo,omitempty" xml:"payeeAccountNo,omitempty"`
	// 收款账号类型
	PayeeAccountType *string `json:"payeeAccountType,omitempty" xml:"payeeAccountType,omitempty"`
	// 状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 收款方电子钱包持有者姓名
	PayeeAccountName *string `json:"payeeAccountName,omitempty" xml:"payeeAccountName,omitempty"`
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 备注
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 订单时间时间
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 支付完成时间
	GmtFinish *string `json:"gmtFinish,omitempty" xml:"gmtFinish,omitempty"`
}

func (s QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) GoString() string {
	return s.String()
}

func (s *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) SetSerialNo(v int64) *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList {
	s.SerialNo = &v
	return s
}

func (s *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) SetDetailNo(v string) *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList {
	s.DetailNo = &v
	return s
}

func (s *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) SetPayeeAccountNo(v string) *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList {
	s.PayeeAccountNo = &v
	return s
}

func (s *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) SetPayeeAccountType(v string) *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList {
	s.PayeeAccountType = &v
	return s
}

func (s *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) SetStatus(v string) *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList {
	s.Status = &v
	return s
}

func (s *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) SetPayeeAccountName(v string) *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList {
	s.PayeeAccountName = &v
	return s
}

func (s *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) SetAmount(v string) *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList {
	s.Amount = &v
	return s
}

func (s *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) SetMemo(v string) *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList {
	s.Memo = &v
	return s
}

func (s *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) SetGmtCreate(v string) *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList {
	s.GmtCreate = &v
	return s
}

func (s *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) SetGmtFinish(v string) *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList {
	s.GmtFinish = &v
	return s
}

type QueryBatchTradeDetailListResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryBatchTradeDetailListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBatchTradeDetailListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchTradeDetailListResponse) GoString() string {
	return s.String()
}

func (s *QueryBatchTradeDetailListResponse) SetHeaders(v map[string]*string) *QueryBatchTradeDetailListResponse {
	s.Headers = v
	return s
}

func (s *QueryBatchTradeDetailListResponse) SetBody(v *QueryBatchTradeDetailListResponseBody) *QueryBatchTradeDetailListResponse {
	s.Body = v
	return s
}

type CreateUserCodeInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateUserCodeInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateUserCodeInstanceHeaders) GoString() string {
	return s.String()
}

func (s *CreateUserCodeInstanceHeaders) SetCommonHeaders(v map[string]*string) *CreateUserCodeInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateUserCodeInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *CreateUserCodeInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateUserCodeInstanceRequest struct {
	// 业务幂等ID
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 码标识，由钉钉颁发
	CodeIdentity *string `json:"codeIdentity,omitempty" xml:"codeIdentity,omitempty"`
	// 码值
	CodeValue *string `json:"codeValue,omitempty" xml:"codeValue,omitempty"`
	// 码值类型，钉钉静态码值：DING_STATIC，访客码或会展码传入
	CodeValueType *string `json:"codeValueType,omitempty" xml:"codeValueType,omitempty"`
	// 状态，传入关闭状态需要用户手动开启后才会渲染二维
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 企业ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 用户和企业的关系类型，区分内部员工，外部联系人，无关系普通用户
	UserCorpRelationType *string `json:"userCorpRelationType,omitempty" xml:"userCorpRelationType,omitempty"`
	// 用户身份标识，取值和用户企业关系类型相关，如果企业无关，传入手机号
	UserIdentity *string `json:"userIdentity,omitempty" xml:"userIdentity,omitempty"`
	// 临时码，传入过期时间
	GmtExpired *string `json:"gmtExpired,omitempty" xml:"gmtExpired,omitempty"`
	// 有效时间列表，对于连续时间段，只需传入一个对象即可，注意过期时间必须晚于最晚结束时间
	AvailableTimes []*CreateUserCodeInstanceRequestAvailableTimes `json:"availableTimes,omitempty" xml:"availableTimes,omitempty" type:"Repeated"`
	// 扩展参数
	ExtInfo      map[string]interface{} `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	DingOrgId    *int64                 `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingIsvOrgId *int64                 `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
}

func (s CreateUserCodeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserCodeInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateUserCodeInstanceRequest) SetRequestId(v string) *CreateUserCodeInstanceRequest {
	s.RequestId = &v
	return s
}

func (s *CreateUserCodeInstanceRequest) SetCodeIdentity(v string) *CreateUserCodeInstanceRequest {
	s.CodeIdentity = &v
	return s
}

func (s *CreateUserCodeInstanceRequest) SetCodeValue(v string) *CreateUserCodeInstanceRequest {
	s.CodeValue = &v
	return s
}

func (s *CreateUserCodeInstanceRequest) SetCodeValueType(v string) *CreateUserCodeInstanceRequest {
	s.CodeValueType = &v
	return s
}

func (s *CreateUserCodeInstanceRequest) SetStatus(v string) *CreateUserCodeInstanceRequest {
	s.Status = &v
	return s
}

func (s *CreateUserCodeInstanceRequest) SetCorpId(v string) *CreateUserCodeInstanceRequest {
	s.CorpId = &v
	return s
}

func (s *CreateUserCodeInstanceRequest) SetUserCorpRelationType(v string) *CreateUserCodeInstanceRequest {
	s.UserCorpRelationType = &v
	return s
}

func (s *CreateUserCodeInstanceRequest) SetUserIdentity(v string) *CreateUserCodeInstanceRequest {
	s.UserIdentity = &v
	return s
}

func (s *CreateUserCodeInstanceRequest) SetGmtExpired(v string) *CreateUserCodeInstanceRequest {
	s.GmtExpired = &v
	return s
}

func (s *CreateUserCodeInstanceRequest) SetAvailableTimes(v []*CreateUserCodeInstanceRequestAvailableTimes) *CreateUserCodeInstanceRequest {
	s.AvailableTimes = v
	return s
}

func (s *CreateUserCodeInstanceRequest) SetExtInfo(v map[string]interface{}) *CreateUserCodeInstanceRequest {
	s.ExtInfo = v
	return s
}

func (s *CreateUserCodeInstanceRequest) SetDingOrgId(v int64) *CreateUserCodeInstanceRequest {
	s.DingOrgId = &v
	return s
}

func (s *CreateUserCodeInstanceRequest) SetDingIsvOrgId(v int64) *CreateUserCodeInstanceRequest {
	s.DingIsvOrgId = &v
	return s
}

type CreateUserCodeInstanceRequestAvailableTimes struct {
	// 开始时间
	GmtStart *string `json:"gmtStart,omitempty" xml:"gmtStart,omitempty"`
	// 结束时间
	GmtEnd *string `json:"gmtEnd,omitempty" xml:"gmtEnd,omitempty"`
}

func (s CreateUserCodeInstanceRequestAvailableTimes) String() string {
	return tea.Prettify(s)
}

func (s CreateUserCodeInstanceRequestAvailableTimes) GoString() string {
	return s.String()
}

func (s *CreateUserCodeInstanceRequestAvailableTimes) SetGmtStart(v string) *CreateUserCodeInstanceRequestAvailableTimes {
	s.GmtStart = &v
	return s
}

func (s *CreateUserCodeInstanceRequestAvailableTimes) SetGmtEnd(v string) *CreateUserCodeInstanceRequestAvailableTimes {
	s.GmtEnd = &v
	return s
}

type CreateUserCodeInstanceResponseBody struct {
	// 码ID
	CodeId *string `json:"codeId,omitempty" xml:"codeId,omitempty"`
	// 码详情跳转地址
	CodeDetailUrl *string `json:"codeDetailUrl,omitempty" xml:"codeDetailUrl,omitempty"`
}

func (s CreateUserCodeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserCodeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserCodeInstanceResponseBody) SetCodeId(v string) *CreateUserCodeInstanceResponseBody {
	s.CodeId = &v
	return s
}

func (s *CreateUserCodeInstanceResponseBody) SetCodeDetailUrl(v string) *CreateUserCodeInstanceResponseBody {
	s.CodeDetailUrl = &v
	return s
}

type CreateUserCodeInstanceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUserCodeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUserCodeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserCodeInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateUserCodeInstanceResponse) SetHeaders(v map[string]*string) *CreateUserCodeInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateUserCodeInstanceResponse) SetBody(v *CreateUserCodeInstanceResponseBody) *CreateUserCodeInstanceResponse {
	s.Body = v
	return s
}

type QueryBatchTradeOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryBatchTradeOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchTradeOrderHeaders) GoString() string {
	return s.String()
}

func (s *QueryBatchTradeOrderHeaders) SetCommonHeaders(v map[string]*string) *QueryBatchTradeOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryBatchTradeOrderHeaders) SetXAcsDingtalkAccessToken(v string) *QueryBatchTradeOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryBatchTradeOrderRequest struct {
	// 外部商户批次号列表
	OutBatchNos []*string `json:"outBatchNos,omitempty" xml:"outBatchNos,omitempty" type:"Repeated"`
}

func (s QueryBatchTradeOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchTradeOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryBatchTradeOrderRequest) SetOutBatchNos(v []*string) *QueryBatchTradeOrderRequest {
	s.OutBatchNos = v
	return s
}

type QueryBatchTradeOrderResponseBody struct {
	// 批量交易订单VO
	BatchTradeOrderVOs []*QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs `json:"batchTradeOrderVOs,omitempty" xml:"batchTradeOrderVOs,omitempty" type:"Repeated"`
}

func (s QueryBatchTradeOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchTradeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBatchTradeOrderResponseBody) SetBatchTradeOrderVOs(v []*QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) *QueryBatchTradeOrderResponseBody {
	s.BatchTradeOrderVOs = v
	return s
}

type QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs struct {
	// 批次号
	OutBatchNo *string `json:"outBatchNo,omitempty" xml:"outBatchNo,omitempty"`
	// 支付宝批次订单号
	AlipayTransId *string `json:"alipayTransId,omitempty" xml:"alipayTransId,omitempty"`
	// 状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 成功笔数
	SuccessCount *int64 `json:"successCount,omitempty" xml:"successCount,omitempty"`
	// 成功金额（元）
	SuccessAmount *string `json:"successAmount,omitempty" xml:"successAmount,omitempty"`
	// 失败笔数
	FailCount *int64 `json:"failCount,omitempty" xml:"failCount,omitempty"`
	// 明细处理失败的支付汇总金额
	FailAmount *string `json:"failAmount,omitempty" xml:"failAmount,omitempty"`
	// 批次的总金额（元）
	TotalAmount *string `json:"totalAmount,omitempty" xml:"totalAmount,omitempty"`
	// 批次完成交易时间
	GmtFinish *string `json:"gmtFinish,omitempty" xml:"gmtFinish,omitempty"`
	// 批次受理交易时间
	GmtSubmit *string `json:"gmtSubmit,omitempty" xml:"gmtSubmit,omitempty"`
	// 失败原因
	FailReason *string `json:"failReason,omitempty" xml:"failReason,omitempty"`
	// 付款方需要支付的金额（元）
	PaymentAmount *string `json:"paymentAmount,omitempty" xml:"paymentAmount,omitempty"`
	// 支付币种
	PaymentCurrency *string `json:"paymentCurrency,omitempty" xml:"paymentCurrency,omitempty"`
	// 付款人staffId
	PayerStaffId *string `json:"payerStaffId,omitempty" xml:"payerStaffId,omitempty"`
}

func (s QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) GoString() string {
	return s.String()
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetOutBatchNo(v string) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.OutBatchNo = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetAlipayTransId(v string) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.AlipayTransId = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetStatus(v string) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.Status = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetSuccessCount(v int64) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.SuccessCount = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetSuccessAmount(v string) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.SuccessAmount = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetFailCount(v int64) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.FailCount = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetFailAmount(v string) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.FailAmount = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetTotalAmount(v string) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.TotalAmount = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetGmtFinish(v string) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.GmtFinish = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetGmtSubmit(v string) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.GmtSubmit = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetFailReason(v string) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.FailReason = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetPaymentAmount(v string) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.PaymentAmount = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetPaymentCurrency(v string) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.PaymentCurrency = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetPayerStaffId(v string) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.PayerStaffId = &v
	return s
}

type QueryBatchTradeOrderResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryBatchTradeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBatchTradeOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchTradeOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryBatchTradeOrderResponse) SetHeaders(v map[string]*string) *QueryBatchTradeOrderResponse {
	s.Headers = v
	return s
}

func (s *QueryBatchTradeOrderResponse) SetBody(v *QueryBatchTradeOrderResponseBody) *QueryBatchTradeOrderResponse {
	s.Body = v
	return s
}

type SaveCorpPayCodeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveCorpPayCodeHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveCorpPayCodeHeaders) GoString() string {
	return s.String()
}

func (s *SaveCorpPayCodeHeaders) SetCommonHeaders(v map[string]*string) *SaveCorpPayCodeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveCorpPayCodeHeaders) SetXAcsDingtalkAccessToken(v string) *SaveCorpPayCodeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveCorpPayCodeRequest struct {
	// 码标识，由钉钉颁发
	CodeIdentity *string `json:"codeIdentity,omitempty" xml:"codeIdentity,omitempty"`
	// 开通的企业ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 状态，OPEN或CLOSED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 扩展参数
	ExtInfo map[string]*string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// 企业orgId
	DingOrgId    *int64 `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingIsvOrgId *int64 `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
}

func (s SaveCorpPayCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveCorpPayCodeRequest) GoString() string {
	return s.String()
}

func (s *SaveCorpPayCodeRequest) SetCodeIdentity(v string) *SaveCorpPayCodeRequest {
	s.CodeIdentity = &v
	return s
}

func (s *SaveCorpPayCodeRequest) SetCorpId(v string) *SaveCorpPayCodeRequest {
	s.CorpId = &v
	return s
}

func (s *SaveCorpPayCodeRequest) SetStatus(v string) *SaveCorpPayCodeRequest {
	s.Status = &v
	return s
}

func (s *SaveCorpPayCodeRequest) SetExtInfo(v map[string]*string) *SaveCorpPayCodeRequest {
	s.ExtInfo = v
	return s
}

func (s *SaveCorpPayCodeRequest) SetDingOrgId(v int64) *SaveCorpPayCodeRequest {
	s.DingOrgId = &v
	return s
}

func (s *SaveCorpPayCodeRequest) SetDingIsvOrgId(v int64) *SaveCorpPayCodeRequest {
	s.DingIsvOrgId = &v
	return s
}

type SaveCorpPayCodeResponseBody struct {
	// 码标识
	CodeIdentity *string `json:"codeIdentity,omitempty" xml:"codeIdentity,omitempty"`
	// 开通的企业ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 扩展参数
	ExtInfo map[string]*string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
}

func (s SaveCorpPayCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveCorpPayCodeResponseBody) GoString() string {
	return s.String()
}

func (s *SaveCorpPayCodeResponseBody) SetCodeIdentity(v string) *SaveCorpPayCodeResponseBody {
	s.CodeIdentity = &v
	return s
}

func (s *SaveCorpPayCodeResponseBody) SetCorpId(v string) *SaveCorpPayCodeResponseBody {
	s.CorpId = &v
	return s
}

func (s *SaveCorpPayCodeResponseBody) SetStatus(v string) *SaveCorpPayCodeResponseBody {
	s.Status = &v
	return s
}

func (s *SaveCorpPayCodeResponseBody) SetExtInfo(v map[string]*string) *SaveCorpPayCodeResponseBody {
	s.ExtInfo = v
	return s
}

type SaveCorpPayCodeResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveCorpPayCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveCorpPayCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveCorpPayCodeResponse) GoString() string {
	return s.String()
}

func (s *SaveCorpPayCodeResponse) SetHeaders(v map[string]*string) *SaveCorpPayCodeResponse {
	s.Headers = v
	return s
}

func (s *SaveCorpPayCodeResponse) SetBody(v *SaveCorpPayCodeResponseBody) *SaveCorpPayCodeResponse {
	s.Body = v
	return s
}

type NotifyVerifyResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s NotifyVerifyResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s NotifyVerifyResultHeaders) GoString() string {
	return s.String()
}

func (s *NotifyVerifyResultHeaders) SetCommonHeaders(v map[string]*string) *NotifyVerifyResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *NotifyVerifyResultHeaders) SetXAcsDingtalkAccessToken(v string) *NotifyVerifyResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type NotifyVerifyResultRequest struct {
	// 码值
	PayCode *string `json:"payCode,omitempty" xml:"payCode,omitempty"`
	// 企业ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 用户和企业的关系类型，区分内部员工，外部联系人，无关系普通用户
	UserCorpRelationType *string `json:"userCorpRelationType,omitempty" xml:"userCorpRelationType,omitempty"`
	// 用户身份标识
	UserIdentity *string `json:"userIdentity,omitempty" xml:"userIdentity,omitempty"`
	// 验证时间
	VerifyTime *string `json:"verifyTime,omitempty" xml:"verifyTime,omitempty"`
	// 验证结果
	VerifyResult *bool `json:"verifyResult,omitempty" xml:"verifyResult,omitempty"`
	// 验证地点，调用时请务必传入，以便生成工牌使用记录
	VerifyLocation *string `json:"verifyLocation,omitempty" xml:"verifyLocation,omitempty"`
	// 验证流水号，长度不超过32位，用户下唯一，调用时请务必传入，以便生成工牌使用记录
	VerifyNo *string `json:"verifyNo,omitempty" xml:"verifyNo,omitempty"`
	// 验证事件，长度不超过8个中文
	VerifyEvent *string `json:"verifyEvent,omitempty" xml:"verifyEvent,omitempty"`
	// 备注信息
	Remark       *string `json:"remark,omitempty" xml:"remark,omitempty"`
	DingOrgId    *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingIsvOrgId *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
}

func (s NotifyVerifyResultRequest) String() string {
	return tea.Prettify(s)
}

func (s NotifyVerifyResultRequest) GoString() string {
	return s.String()
}

func (s *NotifyVerifyResultRequest) SetPayCode(v string) *NotifyVerifyResultRequest {
	s.PayCode = &v
	return s
}

func (s *NotifyVerifyResultRequest) SetCorpId(v string) *NotifyVerifyResultRequest {
	s.CorpId = &v
	return s
}

func (s *NotifyVerifyResultRequest) SetUserCorpRelationType(v string) *NotifyVerifyResultRequest {
	s.UserCorpRelationType = &v
	return s
}

func (s *NotifyVerifyResultRequest) SetUserIdentity(v string) *NotifyVerifyResultRequest {
	s.UserIdentity = &v
	return s
}

func (s *NotifyVerifyResultRequest) SetVerifyTime(v string) *NotifyVerifyResultRequest {
	s.VerifyTime = &v
	return s
}

func (s *NotifyVerifyResultRequest) SetVerifyResult(v bool) *NotifyVerifyResultRequest {
	s.VerifyResult = &v
	return s
}

func (s *NotifyVerifyResultRequest) SetVerifyLocation(v string) *NotifyVerifyResultRequest {
	s.VerifyLocation = &v
	return s
}

func (s *NotifyVerifyResultRequest) SetVerifyNo(v string) *NotifyVerifyResultRequest {
	s.VerifyNo = &v
	return s
}

func (s *NotifyVerifyResultRequest) SetVerifyEvent(v string) *NotifyVerifyResultRequest {
	s.VerifyEvent = &v
	return s
}

func (s *NotifyVerifyResultRequest) SetRemark(v string) *NotifyVerifyResultRequest {
	s.Remark = &v
	return s
}

func (s *NotifyVerifyResultRequest) SetDingOrgId(v int64) *NotifyVerifyResultRequest {
	s.DingOrgId = &v
	return s
}

func (s *NotifyVerifyResultRequest) SetDingIsvOrgId(v int64) *NotifyVerifyResultRequest {
	s.DingIsvOrgId = &v
	return s
}

type NotifyVerifyResultResponseBody struct {
	// 结果
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s NotifyVerifyResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NotifyVerifyResultResponseBody) GoString() string {
	return s.String()
}

func (s *NotifyVerifyResultResponseBody) SetResult(v string) *NotifyVerifyResultResponseBody {
	s.Result = &v
	return s
}

type NotifyVerifyResultResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *NotifyVerifyResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s NotifyVerifyResultResponse) String() string {
	return tea.Prettify(s)
}

func (s NotifyVerifyResultResponse) GoString() string {
	return s.String()
}

func (s *NotifyVerifyResultResponse) SetHeaders(v map[string]*string) *NotifyVerifyResultResponse {
	s.Headers = v
	return s
}

func (s *NotifyVerifyResultResponse) SetBody(v *NotifyVerifyResultResponseBody) *NotifyVerifyResultResponse {
	s.Body = v
	return s
}

type QueryPayAccountListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryPayAccountListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryPayAccountListHeaders) GoString() string {
	return s.String()
}

func (s *QueryPayAccountListHeaders) SetCommonHeaders(v map[string]*string) *QueryPayAccountListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryPayAccountListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryPayAccountListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryPayAccountListResponseBody struct {
	// 账号列表
	PayAccountVOList []*QueryPayAccountListResponseBodyPayAccountVOList `json:"payAccountVOList,omitempty" xml:"payAccountVOList,omitempty" type:"Repeated"`
}

func (s QueryPayAccountListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPayAccountListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPayAccountListResponseBody) SetPayAccountVOList(v []*QueryPayAccountListResponseBodyPayAccountVOList) *QueryPayAccountListResponseBody {
	s.PayAccountVOList = v
	return s
}

type QueryPayAccountListResponseBodyPayAccountVOList struct {
	// 付款账号（脱敏）
	AccountNo *string `json:"accountNo,omitempty" xml:"accountNo,omitempty"`
	// 账号名称
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	// 账户类型
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	// 账户备注
	AccountRemark *string `json:"accountRemark,omitempty" xml:"accountRemark,omitempty"`
	// 账号唯一id
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// 账户分类
	AccountClass *string `json:"accountClass,omitempty" xml:"accountClass,omitempty"`
}

func (s QueryPayAccountListResponseBodyPayAccountVOList) String() string {
	return tea.Prettify(s)
}

func (s QueryPayAccountListResponseBodyPayAccountVOList) GoString() string {
	return s.String()
}

func (s *QueryPayAccountListResponseBodyPayAccountVOList) SetAccountNo(v string) *QueryPayAccountListResponseBodyPayAccountVOList {
	s.AccountNo = &v
	return s
}

func (s *QueryPayAccountListResponseBodyPayAccountVOList) SetAccountName(v string) *QueryPayAccountListResponseBodyPayAccountVOList {
	s.AccountName = &v
	return s
}

func (s *QueryPayAccountListResponseBodyPayAccountVOList) SetAccountType(v string) *QueryPayAccountListResponseBodyPayAccountVOList {
	s.AccountType = &v
	return s
}

func (s *QueryPayAccountListResponseBodyPayAccountVOList) SetAccountRemark(v string) *QueryPayAccountListResponseBodyPayAccountVOList {
	s.AccountRemark = &v
	return s
}

func (s *QueryPayAccountListResponseBodyPayAccountVOList) SetAccountId(v string) *QueryPayAccountListResponseBodyPayAccountVOList {
	s.AccountId = &v
	return s
}

func (s *QueryPayAccountListResponseBodyPayAccountVOList) SetAccountClass(v string) *QueryPayAccountListResponseBodyPayAccountVOList {
	s.AccountClass = &v
	return s
}

type QueryPayAccountListResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryPayAccountListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPayAccountListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPayAccountListResponse) GoString() string {
	return s.String()
}

func (s *QueryPayAccountListResponse) SetHeaders(v map[string]*string) *QueryPayAccountListResponse {
	s.Headers = v
	return s
}

func (s *QueryPayAccountListResponse) SetBody(v *QueryPayAccountListResponseBody) *QueryPayAccountListResponse {
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

func (client *Client) NotifyPayCodePayResult(request *NotifyPayCodePayResultRequest) (_result *NotifyPayCodePayResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &NotifyPayCodePayResultHeaders{}
	_result = &NotifyPayCodePayResultResponse{}
	_body, _err := client.NotifyPayCodePayResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) NotifyPayCodePayResultWithOptions(request *NotifyPayCodePayResultRequest, headers *NotifyPayCodePayResultHeaders, runtime *util.RuntimeOptions) (_result *NotifyPayCodePayResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PayCode)) {
		body["payCode"] = request.PayCode
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.GmtTradeCreate)) {
		body["gmtTradeCreate"] = request.GmtTradeCreate
	}

	if !tea.BoolValue(util.IsUnset(request.GmtTradeFinish)) {
		body["gmtTradeFinish"] = request.GmtTradeFinish
	}

	if !tea.BoolValue(util.IsUnset(request.TradeNo)) {
		body["tradeNo"] = request.TradeNo
	}

	if !tea.BoolValue(util.IsUnset(request.TradeStatus)) {
		body["tradeStatus"] = request.TradeStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		body["amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionAmount)) {
		body["promotionAmount"] = request.PromotionAmount
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeAmount)) {
		body["chargeAmount"] = request.ChargeAmount
	}

	if !tea.BoolValue(util.IsUnset(request.PayChannelDetailList)) {
		body["payChannelDetailList"] = request.PayChannelDetailList
	}

	if !tea.BoolValue(util.IsUnset(request.TradeErrorCode)) {
		body["tradeErrorCode"] = request.TradeErrorCode
	}

	if !tea.BoolValue(util.IsUnset(request.TradeErrorMsg)) {
		body["tradeErrorMsg"] = request.TradeErrorMsg
	}

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["extInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantName)) {
		body["merchantName"] = request.MerchantName
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
	_result = &NotifyPayCodePayResultResponse{}
	_body, _err := client.DoROARequest(tea.String("NotifyPayCodePayResult"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/finance/payCodes/payResults/notify"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpateUserCodeInstance(request *UpateUserCodeInstanceRequest) (_result *UpateUserCodeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpateUserCodeInstanceHeaders{}
	_result = &UpateUserCodeInstanceResponse{}
	_body, _err := client.UpateUserCodeInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpateUserCodeInstanceWithOptions(request *UpateUserCodeInstanceRequest, headers *UpateUserCodeInstanceHeaders, runtime *util.RuntimeOptions) (_result *UpateUserCodeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CodeId)) {
		body["codeId"] = request.CodeId
	}

	if !tea.BoolValue(util.IsUnset(request.CodeIdentity)) {
		body["codeIdentity"] = request.CodeIdentity
	}

	if !tea.BoolValue(util.IsUnset(request.CodeValue)) {
		body["codeValue"] = request.CodeValue
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.UserCorpRelationType)) {
		body["userCorpRelationType"] = request.UserCorpRelationType
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdentity)) {
		body["userIdentity"] = request.UserIdentity
	}

	if !tea.BoolValue(util.IsUnset(request.GmtExpired)) {
		body["gmtExpired"] = request.GmtExpired
	}

	if !tea.BoolValue(util.IsUnset(request.AvailableTimes)) {
		body["availableTimes"] = request.AvailableTimes
	}

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["extInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
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
	_result = &UpateUserCodeInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("UpateUserCodeInstance"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/finance/payCodes/userInstances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyBatchPay(request *ApplyBatchPayRequest) (_result *ApplyBatchPayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ApplyBatchPayHeaders{}
	_result = &ApplyBatchPayResponse{}
	_body, _err := client.ApplyBatchPayWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyBatchPayWithOptions(request *ApplyBatchPayRequest, headers *ApplyBatchPayHeaders, runtime *util.RuntimeOptions) (_result *ApplyBatchPayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StaffId)) {
		body["staffId"] = request.StaffId
	}

	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["accountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNo)) {
		body["orderNo"] = request.OrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.TransAmount)) {
		body["transAmount"] = request.TransAmount
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnUrl)) {
		body["returnUrl"] = request.ReturnUrl
	}

	if !tea.BoolValue(util.IsUnset(request.PassBackParams)) {
		body["passBackParams"] = request.PassBackParams
	}

	if !tea.BoolValue(util.IsUnset(request.PayTerminal)) {
		body["payTerminal"] = request.PayTerminal
	}

	if !tea.BoolValue(util.IsUnset(request.TransExpireTime)) {
		body["transExpireTime"] = request.TransExpireTime
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
	_result = &ApplyBatchPayResponse{}
	_body, _err := client.DoROARequest(tea.String("ApplyBatchPay"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/finance/batchTrades/orders/pay"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserAlipayAccount() (_result *QueryUserAlipayAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserAlipayAccountHeaders{}
	_result = &QueryUserAlipayAccountResponse{}
	_body, _err := client.QueryUserAlipayAccountWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserAlipayAccountWithOptions(headers *QueryUserAlipayAccountHeaders, runtime *util.RuntimeOptions) (_result *QueryUserAlipayAccountResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &QueryUserAlipayAccountResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryUserAlipayAccount"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/finance/userAlipayAccounts"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DecodePayCode(request *DecodePayCodeRequest) (_result *DecodePayCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DecodePayCodeHeaders{}
	_result = &DecodePayCodeResponse{}
	_body, _err := client.DecodePayCodeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DecodePayCodeWithOptions(request *DecodePayCodeRequest, headers *DecodePayCodeHeaders, runtime *util.RuntimeOptions) (_result *DecodePayCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PayCode)) {
		body["payCode"] = request.PayCode
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
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
	_result = &DecodePayCodeResponse{}
	_body, _err := client.DoROARequest(tea.String("DecodePayCode"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/finance/payCodes/decode"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBatchTradeOrder(request *CreateBatchTradeOrderRequest) (_result *CreateBatchTradeOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateBatchTradeOrderHeaders{}
	_result = &CreateBatchTradeOrderResponse{}
	_body, _err := client.CreateBatchTradeOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBatchTradeOrderWithOptions(request *CreateBatchTradeOrderRequest, headers *CreateBatchTradeOrderHeaders, runtime *util.RuntimeOptions) (_result *CreateBatchTradeOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StaffId)) {
		body["staffId"] = request.StaffId
	}

	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["accountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.AccountNo)) {
		body["accountNo"] = request.AccountNo
	}

	if !tea.BoolValue(util.IsUnset(request.TradeTitle)) {
		body["tradeTitle"] = request.TradeTitle
	}

	if !tea.BoolValue(util.IsUnset(request.OutBatchNo)) {
		body["outBatchNo"] = request.OutBatchNo
	}

	if !tea.BoolValue(util.IsUnset(request.BatchRemark)) {
		body["batchRemark"] = request.BatchRemark
	}

	if !tea.BoolValue(util.IsUnset(request.TotalCount)) {
		body["totalCount"] = request.TotalCount
	}

	if !tea.BoolValue(util.IsUnset(request.TotalAmount)) {
		body["totalAmount"] = request.TotalAmount
	}

	if !tea.BoolValue(util.IsUnset(request.BatchTradeDetails)) {
		body["batchTradeDetails"] = request.BatchTradeDetails
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
	_result = &CreateBatchTradeOrderResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateBatchTradeOrder"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/finance/batchTrades/orders"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) NotifyPayCodeRefundResult(request *NotifyPayCodeRefundResultRequest) (_result *NotifyPayCodeRefundResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &NotifyPayCodeRefundResultHeaders{}
	_result = &NotifyPayCodeRefundResultResponse{}
	_body, _err := client.NotifyPayCodeRefundResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) NotifyPayCodeRefundResultWithOptions(request *NotifyPayCodeRefundResultRequest, headers *NotifyPayCodeRefundResultHeaders, runtime *util.RuntimeOptions) (_result *NotifyPayCodeRefundResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.TradeNo)) {
		body["tradeNo"] = request.TradeNo
	}

	if !tea.BoolValue(util.IsUnset(request.RefundOrderNo)) {
		body["refundOrderNo"] = request.RefundOrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.RefundAmount)) {
		body["refundAmount"] = request.RefundAmount
	}

	if !tea.BoolValue(util.IsUnset(request.RefundPromotionAmount)) {
		body["refundPromotionAmount"] = request.RefundPromotionAmount
	}

	if !tea.BoolValue(util.IsUnset(request.GmtRefund)) {
		body["gmtRefund"] = request.GmtRefund
	}

	if !tea.BoolValue(util.IsUnset(request.PayChannelDetailList)) {
		body["payChannelDetailList"] = request.PayChannelDetailList
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PayCode)) {
		body["payCode"] = request.PayCode
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
	_result = &NotifyPayCodeRefundResultResponse{}
	_body, _err := client.DoROARequest(tea.String("NotifyPayCodeRefundResult"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/finance/payCodes/refundResults/notify"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBatchTradeDetailList(request *QueryBatchTradeDetailListRequest) (_result *QueryBatchTradeDetailListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryBatchTradeDetailListHeaders{}
	_result = &QueryBatchTradeDetailListResponse{}
	_body, _err := client.QueryBatchTradeDetailListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBatchTradeDetailListWithOptions(request *QueryBatchTradeDetailListRequest, headers *QueryBatchTradeDetailListHeaders, runtime *util.RuntimeOptions) (_result *QueryBatchTradeDetailListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutBatchNo)) {
		query["outBatchNo"] = request.OutBatchNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
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
	_result = &QueryBatchTradeDetailListResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryBatchTradeDetailList"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/finance/batchTrades/details"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUserCodeInstance(request *CreateUserCodeInstanceRequest) (_result *CreateUserCodeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateUserCodeInstanceHeaders{}
	_result = &CreateUserCodeInstanceResponse{}
	_body, _err := client.CreateUserCodeInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUserCodeInstanceWithOptions(request *CreateUserCodeInstanceRequest, headers *CreateUserCodeInstanceHeaders, runtime *util.RuntimeOptions) (_result *CreateUserCodeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.CodeIdentity)) {
		body["codeIdentity"] = request.CodeIdentity
	}

	if !tea.BoolValue(util.IsUnset(request.CodeValue)) {
		body["codeValue"] = request.CodeValue
	}

	if !tea.BoolValue(util.IsUnset(request.CodeValueType)) {
		body["codeValueType"] = request.CodeValueType
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.UserCorpRelationType)) {
		body["userCorpRelationType"] = request.UserCorpRelationType
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdentity)) {
		body["userIdentity"] = request.UserIdentity
	}

	if !tea.BoolValue(util.IsUnset(request.GmtExpired)) {
		body["gmtExpired"] = request.GmtExpired
	}

	if !tea.BoolValue(util.IsUnset(request.AvailableTimes)) {
		body["availableTimes"] = request.AvailableTimes
	}

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["extInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
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
	_result = &CreateUserCodeInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateUserCodeInstance"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/finance/payCodes/userInstances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBatchTradeOrder(request *QueryBatchTradeOrderRequest) (_result *QueryBatchTradeOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryBatchTradeOrderHeaders{}
	_result = &QueryBatchTradeOrderResponse{}
	_body, _err := client.QueryBatchTradeOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBatchTradeOrderWithOptions(request *QueryBatchTradeOrderRequest, headers *QueryBatchTradeOrderHeaders, runtime *util.RuntimeOptions) (_result *QueryBatchTradeOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutBatchNos)) {
		body["outBatchNos"] = request.OutBatchNos
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
	_result = &QueryBatchTradeOrderResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryBatchTradeOrder"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/finance/batchTrades/orders/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveCorpPayCode(request *SaveCorpPayCodeRequest) (_result *SaveCorpPayCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveCorpPayCodeHeaders{}
	_result = &SaveCorpPayCodeResponse{}
	_body, _err := client.SaveCorpPayCodeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveCorpPayCodeWithOptions(request *SaveCorpPayCodeRequest, headers *SaveCorpPayCodeHeaders, runtime *util.RuntimeOptions) (_result *SaveCorpPayCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CodeIdentity)) {
		body["codeIdentity"] = request.CodeIdentity
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["extInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
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
	_result = &SaveCorpPayCodeResponse{}
	_body, _err := client.DoROARequest(tea.String("SaveCorpPayCode"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/finance/payCodes/corpSettings"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) NotifyVerifyResult(request *NotifyVerifyResultRequest) (_result *NotifyVerifyResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &NotifyVerifyResultHeaders{}
	_result = &NotifyVerifyResultResponse{}
	_body, _err := client.NotifyVerifyResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) NotifyVerifyResultWithOptions(request *NotifyVerifyResultRequest, headers *NotifyVerifyResultHeaders, runtime *util.RuntimeOptions) (_result *NotifyVerifyResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PayCode)) {
		body["payCode"] = request.PayCode
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.UserCorpRelationType)) {
		body["userCorpRelationType"] = request.UserCorpRelationType
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdentity)) {
		body["userIdentity"] = request.UserIdentity
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyTime)) {
		body["verifyTime"] = request.VerifyTime
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyResult)) {
		body["verifyResult"] = request.VerifyResult
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyLocation)) {
		body["verifyLocation"] = request.VerifyLocation
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyNo)) {
		body["verifyNo"] = request.VerifyNo
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyEvent)) {
		body["verifyEvent"] = request.VerifyEvent
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
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
	_result = &NotifyVerifyResultResponse{}
	_body, _err := client.DoROARequest(tea.String("NotifyVerifyResult"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/finance/payCodes/verifyResults/notify"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPayAccountList() (_result *QueryPayAccountListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryPayAccountListHeaders{}
	_result = &QueryPayAccountListResponse{}
	_body, _err := client.QueryPayAccountListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPayAccountListWithOptions(headers *QueryPayAccountListHeaders, runtime *util.RuntimeOptions) (_result *QueryPayAccountListResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &QueryPayAccountListResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryPayAccountList"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/finance/payAccounts"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
