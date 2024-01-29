// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package badge_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateBadgeCodeUserInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateBadgeCodeUserInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateBadgeCodeUserInstanceHeaders) GoString() string {
	return s.String()
}

func (s *CreateBadgeCodeUserInstanceHeaders) SetCommonHeaders(v map[string]*string) *CreateBadgeCodeUserInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateBadgeCodeUserInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *CreateBadgeCodeUserInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateBadgeCodeUserInstanceRequest struct {
	AvailableTimes       []*CreateBadgeCodeUserInstanceRequestAvailableTimes `json:"availableTimes,omitempty" xml:"availableTimes,omitempty" type:"Repeated"`
	CodeIdentity         *string                                             `json:"codeIdentity,omitempty" xml:"codeIdentity,omitempty"`
	CodeValue            *string                                             `json:"codeValue,omitempty" xml:"codeValue,omitempty"`
	CodeValueType        *string                                             `json:"codeValueType,omitempty" xml:"codeValueType,omitempty"`
	CorpId               *string                                             `json:"corpId,omitempty" xml:"corpId,omitempty"`
	ExtInfo              map[string]interface{}                              `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	GmtExpired           *string                                             `json:"gmtExpired,omitempty" xml:"gmtExpired,omitempty"`
	RequestId            *string                                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Status               *string                                             `json:"status,omitempty" xml:"status,omitempty"`
	UserCorpRelationType *string                                             `json:"userCorpRelationType,omitempty" xml:"userCorpRelationType,omitempty"`
	UserIdentity         *string                                             `json:"userIdentity,omitempty" xml:"userIdentity,omitempty"`
}

func (s CreateBadgeCodeUserInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBadgeCodeUserInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateBadgeCodeUserInstanceRequest) SetAvailableTimes(v []*CreateBadgeCodeUserInstanceRequestAvailableTimes) *CreateBadgeCodeUserInstanceRequest {
	s.AvailableTimes = v
	return s
}

func (s *CreateBadgeCodeUserInstanceRequest) SetCodeIdentity(v string) *CreateBadgeCodeUserInstanceRequest {
	s.CodeIdentity = &v
	return s
}

func (s *CreateBadgeCodeUserInstanceRequest) SetCodeValue(v string) *CreateBadgeCodeUserInstanceRequest {
	s.CodeValue = &v
	return s
}

func (s *CreateBadgeCodeUserInstanceRequest) SetCodeValueType(v string) *CreateBadgeCodeUserInstanceRequest {
	s.CodeValueType = &v
	return s
}

func (s *CreateBadgeCodeUserInstanceRequest) SetCorpId(v string) *CreateBadgeCodeUserInstanceRequest {
	s.CorpId = &v
	return s
}

func (s *CreateBadgeCodeUserInstanceRequest) SetExtInfo(v map[string]interface{}) *CreateBadgeCodeUserInstanceRequest {
	s.ExtInfo = v
	return s
}

func (s *CreateBadgeCodeUserInstanceRequest) SetGmtExpired(v string) *CreateBadgeCodeUserInstanceRequest {
	s.GmtExpired = &v
	return s
}

func (s *CreateBadgeCodeUserInstanceRequest) SetRequestId(v string) *CreateBadgeCodeUserInstanceRequest {
	s.RequestId = &v
	return s
}

func (s *CreateBadgeCodeUserInstanceRequest) SetStatus(v string) *CreateBadgeCodeUserInstanceRequest {
	s.Status = &v
	return s
}

func (s *CreateBadgeCodeUserInstanceRequest) SetUserCorpRelationType(v string) *CreateBadgeCodeUserInstanceRequest {
	s.UserCorpRelationType = &v
	return s
}

func (s *CreateBadgeCodeUserInstanceRequest) SetUserIdentity(v string) *CreateBadgeCodeUserInstanceRequest {
	s.UserIdentity = &v
	return s
}

type CreateBadgeCodeUserInstanceRequestAvailableTimes struct {
	GmtEnd   *string `json:"gmtEnd,omitempty" xml:"gmtEnd,omitempty"`
	GmtStart *string `json:"gmtStart,omitempty" xml:"gmtStart,omitempty"`
}

func (s CreateBadgeCodeUserInstanceRequestAvailableTimes) String() string {
	return tea.Prettify(s)
}

func (s CreateBadgeCodeUserInstanceRequestAvailableTimes) GoString() string {
	return s.String()
}

func (s *CreateBadgeCodeUserInstanceRequestAvailableTimes) SetGmtEnd(v string) *CreateBadgeCodeUserInstanceRequestAvailableTimes {
	s.GmtEnd = &v
	return s
}

func (s *CreateBadgeCodeUserInstanceRequestAvailableTimes) SetGmtStart(v string) *CreateBadgeCodeUserInstanceRequestAvailableTimes {
	s.GmtStart = &v
	return s
}

type CreateBadgeCodeUserInstanceResponseBody struct {
	CodeDetailUrl *string `json:"codeDetailUrl,omitempty" xml:"codeDetailUrl,omitempty"`
	CodeId        *string `json:"codeId,omitempty" xml:"codeId,omitempty"`
}

func (s CreateBadgeCodeUserInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBadgeCodeUserInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBadgeCodeUserInstanceResponseBody) SetCodeDetailUrl(v string) *CreateBadgeCodeUserInstanceResponseBody {
	s.CodeDetailUrl = &v
	return s
}

func (s *CreateBadgeCodeUserInstanceResponseBody) SetCodeId(v string) *CreateBadgeCodeUserInstanceResponseBody {
	s.CodeId = &v
	return s
}

type CreateBadgeCodeUserInstanceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBadgeCodeUserInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBadgeCodeUserInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBadgeCodeUserInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateBadgeCodeUserInstanceResponse) SetHeaders(v map[string]*string) *CreateBadgeCodeUserInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateBadgeCodeUserInstanceResponse) SetStatusCode(v int32) *CreateBadgeCodeUserInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBadgeCodeUserInstanceResponse) SetBody(v *CreateBadgeCodeUserInstanceResponseBody) *CreateBadgeCodeUserInstanceResponse {
	s.Body = v
	return s
}

type CreateBadgeNotifyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateBadgeNotifyHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateBadgeNotifyHeaders) GoString() string {
	return s.String()
}

func (s *CreateBadgeNotifyHeaders) SetCommonHeaders(v map[string]*string) *CreateBadgeNotifyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateBadgeNotifyHeaders) SetXAcsDingtalkAccessToken(v string) *CreateBadgeNotifyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateBadgeNotifyRequest struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	MsgId   *string `json:"msgId,omitempty" xml:"msgId,omitempty"`
	MsgType *string `json:"msgType,omitempty" xml:"msgType,omitempty"`
	UserId  *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateBadgeNotifyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBadgeNotifyRequest) GoString() string {
	return s.String()
}

func (s *CreateBadgeNotifyRequest) SetContent(v string) *CreateBadgeNotifyRequest {
	s.Content = &v
	return s
}

func (s *CreateBadgeNotifyRequest) SetMsgId(v string) *CreateBadgeNotifyRequest {
	s.MsgId = &v
	return s
}

func (s *CreateBadgeNotifyRequest) SetMsgType(v string) *CreateBadgeNotifyRequest {
	s.MsgType = &v
	return s
}

func (s *CreateBadgeNotifyRequest) SetUserId(v string) *CreateBadgeNotifyRequest {
	s.UserId = &v
	return s
}

type CreateBadgeNotifyResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateBadgeNotifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBadgeNotifyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBadgeNotifyResponseBody) SetResult(v bool) *CreateBadgeNotifyResponseBody {
	s.Result = &v
	return s
}

type CreateBadgeNotifyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBadgeNotifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBadgeNotifyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBadgeNotifyResponse) GoString() string {
	return s.String()
}

func (s *CreateBadgeNotifyResponse) SetHeaders(v map[string]*string) *CreateBadgeNotifyResponse {
	s.Headers = v
	return s
}

func (s *CreateBadgeNotifyResponse) SetStatusCode(v int32) *CreateBadgeNotifyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBadgeNotifyResponse) SetBody(v *CreateBadgeNotifyResponseBody) *CreateBadgeNotifyResponse {
	s.Body = v
	return s
}

type DecodeBadgeCodeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DecodeBadgeCodeHeaders) String() string {
	return tea.Prettify(s)
}

func (s DecodeBadgeCodeHeaders) GoString() string {
	return s.String()
}

func (s *DecodeBadgeCodeHeaders) SetCommonHeaders(v map[string]*string) *DecodeBadgeCodeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DecodeBadgeCodeHeaders) SetXAcsDingtalkAccessToken(v string) *DecodeBadgeCodeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DecodeBadgeCodeRequest struct {
	PayCode   *string `json:"payCode,omitempty" xml:"payCode,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DecodeBadgeCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s DecodeBadgeCodeRequest) GoString() string {
	return s.String()
}

func (s *DecodeBadgeCodeRequest) SetPayCode(v string) *DecodeBadgeCodeRequest {
	s.PayCode = &v
	return s
}

func (s *DecodeBadgeCodeRequest) SetRequestId(v string) *DecodeBadgeCodeRequest {
	s.RequestId = &v
	return s
}

type DecodeBadgeCodeResponseBody struct {
	AlipayCode           *string `json:"alipayCode,omitempty" xml:"alipayCode,omitempty"`
	CodeId               *string `json:"codeId,omitempty" xml:"codeId,omitempty"`
	CodeIdentity         *string `json:"codeIdentity,omitempty" xml:"codeIdentity,omitempty"`
	CodeType             *string `json:"codeType,omitempty" xml:"codeType,omitempty"`
	CorpId               *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	ExtInfo              *string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	OutBizId             *string `json:"outBizId,omitempty" xml:"outBizId,omitempty"`
	UserCorpRelationType *string `json:"userCorpRelationType,omitempty" xml:"userCorpRelationType,omitempty"`
	UserId               *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DecodeBadgeCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DecodeBadgeCodeResponseBody) GoString() string {
	return s.String()
}

func (s *DecodeBadgeCodeResponseBody) SetAlipayCode(v string) *DecodeBadgeCodeResponseBody {
	s.AlipayCode = &v
	return s
}

func (s *DecodeBadgeCodeResponseBody) SetCodeId(v string) *DecodeBadgeCodeResponseBody {
	s.CodeId = &v
	return s
}

func (s *DecodeBadgeCodeResponseBody) SetCodeIdentity(v string) *DecodeBadgeCodeResponseBody {
	s.CodeIdentity = &v
	return s
}

func (s *DecodeBadgeCodeResponseBody) SetCodeType(v string) *DecodeBadgeCodeResponseBody {
	s.CodeType = &v
	return s
}

func (s *DecodeBadgeCodeResponseBody) SetCorpId(v string) *DecodeBadgeCodeResponseBody {
	s.CorpId = &v
	return s
}

func (s *DecodeBadgeCodeResponseBody) SetExtInfo(v string) *DecodeBadgeCodeResponseBody {
	s.ExtInfo = &v
	return s
}

func (s *DecodeBadgeCodeResponseBody) SetOutBizId(v string) *DecodeBadgeCodeResponseBody {
	s.OutBizId = &v
	return s
}

func (s *DecodeBadgeCodeResponseBody) SetUserCorpRelationType(v string) *DecodeBadgeCodeResponseBody {
	s.UserCorpRelationType = &v
	return s
}

func (s *DecodeBadgeCodeResponseBody) SetUserId(v string) *DecodeBadgeCodeResponseBody {
	s.UserId = &v
	return s
}

type DecodeBadgeCodeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DecodeBadgeCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DecodeBadgeCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DecodeBadgeCodeResponse) GoString() string {
	return s.String()
}

func (s *DecodeBadgeCodeResponse) SetHeaders(v map[string]*string) *DecodeBadgeCodeResponse {
	s.Headers = v
	return s
}

func (s *DecodeBadgeCodeResponse) SetStatusCode(v int32) *DecodeBadgeCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DecodeBadgeCodeResponse) SetBody(v *DecodeBadgeCodeResponseBody) *DecodeBadgeCodeResponse {
	s.Body = v
	return s
}

type NotifyBadgeCodePayResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s NotifyBadgeCodePayResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s NotifyBadgeCodePayResultHeaders) GoString() string {
	return s.String()
}

func (s *NotifyBadgeCodePayResultHeaders) SetCommonHeaders(v map[string]*string) *NotifyBadgeCodePayResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *NotifyBadgeCodePayResultHeaders) SetXAcsDingtalkAccessToken(v string) *NotifyBadgeCodePayResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type NotifyBadgeCodePayResultRequest struct {
	Amount               *string                                                `json:"amount,omitempty" xml:"amount,omitempty"`
	ChargeAmount         *string                                                `json:"chargeAmount,omitempty" xml:"chargeAmount,omitempty"`
	CorpId               *string                                                `json:"corpId,omitempty" xml:"corpId,omitempty"`
	ExtInfo              *string                                                `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	GmtTradeCreate       *string                                                `json:"gmtTradeCreate,omitempty" xml:"gmtTradeCreate,omitempty"`
	GmtTradeFinish       *string                                                `json:"gmtTradeFinish,omitempty" xml:"gmtTradeFinish,omitempty"`
	MerchantName         *string                                                `json:"merchantName,omitempty" xml:"merchantName,omitempty"`
	PayChannelDetailList []*NotifyBadgeCodePayResultRequestPayChannelDetailList `json:"payChannelDetailList,omitempty" xml:"payChannelDetailList,omitempty" type:"Repeated"`
	PayCode              *string                                                `json:"payCode,omitempty" xml:"payCode,omitempty"`
	PromotionAmount      *string                                                `json:"promotionAmount,omitempty" xml:"promotionAmount,omitempty"`
	Remark               *string                                                `json:"remark,omitempty" xml:"remark,omitempty"`
	Title                *string                                                `json:"title,omitempty" xml:"title,omitempty"`
	TradeErrorCode       *string                                                `json:"tradeErrorCode,omitempty" xml:"tradeErrorCode,omitempty"`
	TradeErrorMsg        *string                                                `json:"tradeErrorMsg,omitempty" xml:"tradeErrorMsg,omitempty"`
	TradeNo              *string                                                `json:"tradeNo,omitempty" xml:"tradeNo,omitempty"`
	TradeStatus          *string                                                `json:"tradeStatus,omitempty" xml:"tradeStatus,omitempty"`
	UserId               *string                                                `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s NotifyBadgeCodePayResultRequest) String() string {
	return tea.Prettify(s)
}

func (s NotifyBadgeCodePayResultRequest) GoString() string {
	return s.String()
}

func (s *NotifyBadgeCodePayResultRequest) SetAmount(v string) *NotifyBadgeCodePayResultRequest {
	s.Amount = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequest) SetChargeAmount(v string) *NotifyBadgeCodePayResultRequest {
	s.ChargeAmount = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequest) SetCorpId(v string) *NotifyBadgeCodePayResultRequest {
	s.CorpId = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequest) SetExtInfo(v string) *NotifyBadgeCodePayResultRequest {
	s.ExtInfo = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequest) SetGmtTradeCreate(v string) *NotifyBadgeCodePayResultRequest {
	s.GmtTradeCreate = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequest) SetGmtTradeFinish(v string) *NotifyBadgeCodePayResultRequest {
	s.GmtTradeFinish = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequest) SetMerchantName(v string) *NotifyBadgeCodePayResultRequest {
	s.MerchantName = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequest) SetPayChannelDetailList(v []*NotifyBadgeCodePayResultRequestPayChannelDetailList) *NotifyBadgeCodePayResultRequest {
	s.PayChannelDetailList = v
	return s
}

func (s *NotifyBadgeCodePayResultRequest) SetPayCode(v string) *NotifyBadgeCodePayResultRequest {
	s.PayCode = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequest) SetPromotionAmount(v string) *NotifyBadgeCodePayResultRequest {
	s.PromotionAmount = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequest) SetRemark(v string) *NotifyBadgeCodePayResultRequest {
	s.Remark = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequest) SetTitle(v string) *NotifyBadgeCodePayResultRequest {
	s.Title = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequest) SetTradeErrorCode(v string) *NotifyBadgeCodePayResultRequest {
	s.TradeErrorCode = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequest) SetTradeErrorMsg(v string) *NotifyBadgeCodePayResultRequest {
	s.TradeErrorMsg = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequest) SetTradeNo(v string) *NotifyBadgeCodePayResultRequest {
	s.TradeNo = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequest) SetTradeStatus(v string) *NotifyBadgeCodePayResultRequest {
	s.TradeStatus = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequest) SetUserId(v string) *NotifyBadgeCodePayResultRequest {
	s.UserId = &v
	return s
}

type NotifyBadgeCodePayResultRequestPayChannelDetailList struct {
	Amount             *string                                                                  `json:"amount,omitempty" xml:"amount,omitempty"`
	FundToolDetailList []*NotifyBadgeCodePayResultRequestPayChannelDetailListFundToolDetailList `json:"fundToolDetailList,omitempty" xml:"fundToolDetailList,omitempty" type:"Repeated"`
	GmtCreate          *string                                                                  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtFinish          *string                                                                  `json:"gmtFinish,omitempty" xml:"gmtFinish,omitempty"`
	PayChannelName     *string                                                                  `json:"payChannelName,omitempty" xml:"payChannelName,omitempty"`
	PayChannelOrderNo  *string                                                                  `json:"payChannelOrderNo,omitempty" xml:"payChannelOrderNo,omitempty"`
	PayChannelType     *string                                                                  `json:"payChannelType,omitempty" xml:"payChannelType,omitempty"`
	PromotionAmount    *string                                                                  `json:"promotionAmount,omitempty" xml:"promotionAmount,omitempty"`
}

func (s NotifyBadgeCodePayResultRequestPayChannelDetailList) String() string {
	return tea.Prettify(s)
}

func (s NotifyBadgeCodePayResultRequestPayChannelDetailList) GoString() string {
	return s.String()
}

func (s *NotifyBadgeCodePayResultRequestPayChannelDetailList) SetAmount(v string) *NotifyBadgeCodePayResultRequestPayChannelDetailList {
	s.Amount = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequestPayChannelDetailList) SetFundToolDetailList(v []*NotifyBadgeCodePayResultRequestPayChannelDetailListFundToolDetailList) *NotifyBadgeCodePayResultRequestPayChannelDetailList {
	s.FundToolDetailList = v
	return s
}

func (s *NotifyBadgeCodePayResultRequestPayChannelDetailList) SetGmtCreate(v string) *NotifyBadgeCodePayResultRequestPayChannelDetailList {
	s.GmtCreate = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequestPayChannelDetailList) SetGmtFinish(v string) *NotifyBadgeCodePayResultRequestPayChannelDetailList {
	s.GmtFinish = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequestPayChannelDetailList) SetPayChannelName(v string) *NotifyBadgeCodePayResultRequestPayChannelDetailList {
	s.PayChannelName = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequestPayChannelDetailList) SetPayChannelOrderNo(v string) *NotifyBadgeCodePayResultRequestPayChannelDetailList {
	s.PayChannelOrderNo = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequestPayChannelDetailList) SetPayChannelType(v string) *NotifyBadgeCodePayResultRequestPayChannelDetailList {
	s.PayChannelType = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequestPayChannelDetailList) SetPromotionAmount(v string) *NotifyBadgeCodePayResultRequestPayChannelDetailList {
	s.PromotionAmount = &v
	return s
}

type NotifyBadgeCodePayResultRequestPayChannelDetailListFundToolDetailList struct {
	Amount            *string `json:"amount,omitempty" xml:"amount,omitempty"`
	ExtInfo           *string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	FundToolName      *string `json:"fundToolName,omitempty" xml:"fundToolName,omitempty"`
	GmtCreate         *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtFinish         *string `json:"gmtFinish,omitempty" xml:"gmtFinish,omitempty"`
	PromotionFundTool *bool   `json:"promotionFundTool,omitempty" xml:"promotionFundTool,omitempty"`
}

func (s NotifyBadgeCodePayResultRequestPayChannelDetailListFundToolDetailList) String() string {
	return tea.Prettify(s)
}

func (s NotifyBadgeCodePayResultRequestPayChannelDetailListFundToolDetailList) GoString() string {
	return s.String()
}

func (s *NotifyBadgeCodePayResultRequestPayChannelDetailListFundToolDetailList) SetAmount(v string) *NotifyBadgeCodePayResultRequestPayChannelDetailListFundToolDetailList {
	s.Amount = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequestPayChannelDetailListFundToolDetailList) SetExtInfo(v string) *NotifyBadgeCodePayResultRequestPayChannelDetailListFundToolDetailList {
	s.ExtInfo = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequestPayChannelDetailListFundToolDetailList) SetFundToolName(v string) *NotifyBadgeCodePayResultRequestPayChannelDetailListFundToolDetailList {
	s.FundToolName = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequestPayChannelDetailListFundToolDetailList) SetGmtCreate(v string) *NotifyBadgeCodePayResultRequestPayChannelDetailListFundToolDetailList {
	s.GmtCreate = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequestPayChannelDetailListFundToolDetailList) SetGmtFinish(v string) *NotifyBadgeCodePayResultRequestPayChannelDetailListFundToolDetailList {
	s.GmtFinish = &v
	return s
}

func (s *NotifyBadgeCodePayResultRequestPayChannelDetailListFundToolDetailList) SetPromotionFundTool(v bool) *NotifyBadgeCodePayResultRequestPayChannelDetailListFundToolDetailList {
	s.PromotionFundTool = &v
	return s
}

type NotifyBadgeCodePayResultResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s NotifyBadgeCodePayResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NotifyBadgeCodePayResultResponseBody) GoString() string {
	return s.String()
}

func (s *NotifyBadgeCodePayResultResponseBody) SetResult(v string) *NotifyBadgeCodePayResultResponseBody {
	s.Result = &v
	return s
}

type NotifyBadgeCodePayResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NotifyBadgeCodePayResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NotifyBadgeCodePayResultResponse) String() string {
	return tea.Prettify(s)
}

func (s NotifyBadgeCodePayResultResponse) GoString() string {
	return s.String()
}

func (s *NotifyBadgeCodePayResultResponse) SetHeaders(v map[string]*string) *NotifyBadgeCodePayResultResponse {
	s.Headers = v
	return s
}

func (s *NotifyBadgeCodePayResultResponse) SetStatusCode(v int32) *NotifyBadgeCodePayResultResponse {
	s.StatusCode = &v
	return s
}

func (s *NotifyBadgeCodePayResultResponse) SetBody(v *NotifyBadgeCodePayResultResponseBody) *NotifyBadgeCodePayResultResponse {
	s.Body = v
	return s
}

type NotifyBadgeCodeRefundResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s NotifyBadgeCodeRefundResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s NotifyBadgeCodeRefundResultHeaders) GoString() string {
	return s.String()
}

func (s *NotifyBadgeCodeRefundResultHeaders) SetCommonHeaders(v map[string]*string) *NotifyBadgeCodeRefundResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *NotifyBadgeCodeRefundResultHeaders) SetXAcsDingtalkAccessToken(v string) *NotifyBadgeCodeRefundResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type NotifyBadgeCodeRefundResultRequest struct {
	CorpId                *string                                                   `json:"corpId,omitempty" xml:"corpId,omitempty"`
	GmtRefund             *string                                                   `json:"gmtRefund,omitempty" xml:"gmtRefund,omitempty"`
	PayChannelDetailList  []*NotifyBadgeCodeRefundResultRequestPayChannelDetailList `json:"payChannelDetailList,omitempty" xml:"payChannelDetailList,omitempty" type:"Repeated"`
	PayCode               *string                                                   `json:"payCode,omitempty" xml:"payCode,omitempty"`
	RefundAmount          *string                                                   `json:"refundAmount,omitempty" xml:"refundAmount,omitempty"`
	RefundOrderNo         *string                                                   `json:"refundOrderNo,omitempty" xml:"refundOrderNo,omitempty"`
	RefundPromotionAmount *string                                                   `json:"refundPromotionAmount,omitempty" xml:"refundPromotionAmount,omitempty"`
	Remark                *string                                                   `json:"remark,omitempty" xml:"remark,omitempty"`
	TradeNo               *string                                                   `json:"tradeNo,omitempty" xml:"tradeNo,omitempty"`
	UserId                *string                                                   `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s NotifyBadgeCodeRefundResultRequest) String() string {
	return tea.Prettify(s)
}

func (s NotifyBadgeCodeRefundResultRequest) GoString() string {
	return s.String()
}

func (s *NotifyBadgeCodeRefundResultRequest) SetCorpId(v string) *NotifyBadgeCodeRefundResultRequest {
	s.CorpId = &v
	return s
}

func (s *NotifyBadgeCodeRefundResultRequest) SetGmtRefund(v string) *NotifyBadgeCodeRefundResultRequest {
	s.GmtRefund = &v
	return s
}

func (s *NotifyBadgeCodeRefundResultRequest) SetPayChannelDetailList(v []*NotifyBadgeCodeRefundResultRequestPayChannelDetailList) *NotifyBadgeCodeRefundResultRequest {
	s.PayChannelDetailList = v
	return s
}

func (s *NotifyBadgeCodeRefundResultRequest) SetPayCode(v string) *NotifyBadgeCodeRefundResultRequest {
	s.PayCode = &v
	return s
}

func (s *NotifyBadgeCodeRefundResultRequest) SetRefundAmount(v string) *NotifyBadgeCodeRefundResultRequest {
	s.RefundAmount = &v
	return s
}

func (s *NotifyBadgeCodeRefundResultRequest) SetRefundOrderNo(v string) *NotifyBadgeCodeRefundResultRequest {
	s.RefundOrderNo = &v
	return s
}

func (s *NotifyBadgeCodeRefundResultRequest) SetRefundPromotionAmount(v string) *NotifyBadgeCodeRefundResultRequest {
	s.RefundPromotionAmount = &v
	return s
}

func (s *NotifyBadgeCodeRefundResultRequest) SetRemark(v string) *NotifyBadgeCodeRefundResultRequest {
	s.Remark = &v
	return s
}

func (s *NotifyBadgeCodeRefundResultRequest) SetTradeNo(v string) *NotifyBadgeCodeRefundResultRequest {
	s.TradeNo = &v
	return s
}

func (s *NotifyBadgeCodeRefundResultRequest) SetUserId(v string) *NotifyBadgeCodeRefundResultRequest {
	s.UserId = &v
	return s
}

type NotifyBadgeCodeRefundResultRequestPayChannelDetailList struct {
	Amount                  *string                                                                     `json:"amount,omitempty" xml:"amount,omitempty"`
	FundToolDetailList      []*NotifyBadgeCodeRefundResultRequestPayChannelDetailListFundToolDetailList `json:"fundToolDetailList,omitempty" xml:"fundToolDetailList,omitempty" type:"Repeated"`
	PayChannelName          *string                                                                     `json:"payChannelName,omitempty" xml:"payChannelName,omitempty"`
	PayChannelOrderNo       *string                                                                     `json:"payChannelOrderNo,omitempty" xml:"payChannelOrderNo,omitempty"`
	PayChannelRefundOrderNo *string                                                                     `json:"payChannelRefundOrderNo,omitempty" xml:"payChannelRefundOrderNo,omitempty"`
	PayChannelType          *string                                                                     `json:"payChannelType,omitempty" xml:"payChannelType,omitempty"`
	PromotionAmount         *string                                                                     `json:"promotionAmount,omitempty" xml:"promotionAmount,omitempty"`
}

func (s NotifyBadgeCodeRefundResultRequestPayChannelDetailList) String() string {
	return tea.Prettify(s)
}

func (s NotifyBadgeCodeRefundResultRequestPayChannelDetailList) GoString() string {
	return s.String()
}

func (s *NotifyBadgeCodeRefundResultRequestPayChannelDetailList) SetAmount(v string) *NotifyBadgeCodeRefundResultRequestPayChannelDetailList {
	s.Amount = &v
	return s
}

func (s *NotifyBadgeCodeRefundResultRequestPayChannelDetailList) SetFundToolDetailList(v []*NotifyBadgeCodeRefundResultRequestPayChannelDetailListFundToolDetailList) *NotifyBadgeCodeRefundResultRequestPayChannelDetailList {
	s.FundToolDetailList = v
	return s
}

func (s *NotifyBadgeCodeRefundResultRequestPayChannelDetailList) SetPayChannelName(v string) *NotifyBadgeCodeRefundResultRequestPayChannelDetailList {
	s.PayChannelName = &v
	return s
}

func (s *NotifyBadgeCodeRefundResultRequestPayChannelDetailList) SetPayChannelOrderNo(v string) *NotifyBadgeCodeRefundResultRequestPayChannelDetailList {
	s.PayChannelOrderNo = &v
	return s
}

func (s *NotifyBadgeCodeRefundResultRequestPayChannelDetailList) SetPayChannelRefundOrderNo(v string) *NotifyBadgeCodeRefundResultRequestPayChannelDetailList {
	s.PayChannelRefundOrderNo = &v
	return s
}

func (s *NotifyBadgeCodeRefundResultRequestPayChannelDetailList) SetPayChannelType(v string) *NotifyBadgeCodeRefundResultRequestPayChannelDetailList {
	s.PayChannelType = &v
	return s
}

func (s *NotifyBadgeCodeRefundResultRequestPayChannelDetailList) SetPromotionAmount(v string) *NotifyBadgeCodeRefundResultRequestPayChannelDetailList {
	s.PromotionAmount = &v
	return s
}

type NotifyBadgeCodeRefundResultRequestPayChannelDetailListFundToolDetailList struct {
	Amount            *string `json:"amount,omitempty" xml:"amount,omitempty"`
	ExtInfo           *string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	FundToolName      *string `json:"fundToolName,omitempty" xml:"fundToolName,omitempty"`
	GmtCreate         *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtFinish         *string `json:"gmtFinish,omitempty" xml:"gmtFinish,omitempty"`
	PromotionFundTool *bool   `json:"promotionFundTool,omitempty" xml:"promotionFundTool,omitempty"`
}

func (s NotifyBadgeCodeRefundResultRequestPayChannelDetailListFundToolDetailList) String() string {
	return tea.Prettify(s)
}

func (s NotifyBadgeCodeRefundResultRequestPayChannelDetailListFundToolDetailList) GoString() string {
	return s.String()
}

func (s *NotifyBadgeCodeRefundResultRequestPayChannelDetailListFundToolDetailList) SetAmount(v string) *NotifyBadgeCodeRefundResultRequestPayChannelDetailListFundToolDetailList {
	s.Amount = &v
	return s
}

func (s *NotifyBadgeCodeRefundResultRequestPayChannelDetailListFundToolDetailList) SetExtInfo(v string) *NotifyBadgeCodeRefundResultRequestPayChannelDetailListFundToolDetailList {
	s.ExtInfo = &v
	return s
}

func (s *NotifyBadgeCodeRefundResultRequestPayChannelDetailListFundToolDetailList) SetFundToolName(v string) *NotifyBadgeCodeRefundResultRequestPayChannelDetailListFundToolDetailList {
	s.FundToolName = &v
	return s
}

func (s *NotifyBadgeCodeRefundResultRequestPayChannelDetailListFundToolDetailList) SetGmtCreate(v string) *NotifyBadgeCodeRefundResultRequestPayChannelDetailListFundToolDetailList {
	s.GmtCreate = &v
	return s
}

func (s *NotifyBadgeCodeRefundResultRequestPayChannelDetailListFundToolDetailList) SetGmtFinish(v string) *NotifyBadgeCodeRefundResultRequestPayChannelDetailListFundToolDetailList {
	s.GmtFinish = &v
	return s
}

func (s *NotifyBadgeCodeRefundResultRequestPayChannelDetailListFundToolDetailList) SetPromotionFundTool(v bool) *NotifyBadgeCodeRefundResultRequestPayChannelDetailListFundToolDetailList {
	s.PromotionFundTool = &v
	return s
}

type NotifyBadgeCodeRefundResultResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s NotifyBadgeCodeRefundResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NotifyBadgeCodeRefundResultResponseBody) GoString() string {
	return s.String()
}

func (s *NotifyBadgeCodeRefundResultResponseBody) SetResult(v string) *NotifyBadgeCodeRefundResultResponseBody {
	s.Result = &v
	return s
}

type NotifyBadgeCodeRefundResultResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NotifyBadgeCodeRefundResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NotifyBadgeCodeRefundResultResponse) String() string {
	return tea.Prettify(s)
}

func (s NotifyBadgeCodeRefundResultResponse) GoString() string {
	return s.String()
}

func (s *NotifyBadgeCodeRefundResultResponse) SetHeaders(v map[string]*string) *NotifyBadgeCodeRefundResultResponse {
	s.Headers = v
	return s
}

func (s *NotifyBadgeCodeRefundResultResponse) SetStatusCode(v int32) *NotifyBadgeCodeRefundResultResponse {
	s.StatusCode = &v
	return s
}

func (s *NotifyBadgeCodeRefundResultResponse) SetBody(v *NotifyBadgeCodeRefundResultResponseBody) *NotifyBadgeCodeRefundResultResponse {
	s.Body = v
	return s
}

type NotifyBadgeCodeVerifyResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s NotifyBadgeCodeVerifyResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s NotifyBadgeCodeVerifyResultHeaders) GoString() string {
	return s.String()
}

func (s *NotifyBadgeCodeVerifyResultHeaders) SetCommonHeaders(v map[string]*string) *NotifyBadgeCodeVerifyResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *NotifyBadgeCodeVerifyResultHeaders) SetXAcsDingtalkAccessToken(v string) *NotifyBadgeCodeVerifyResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type NotifyBadgeCodeVerifyResultRequest struct {
	CorpId               *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	PayCode              *string `json:"payCode,omitempty" xml:"payCode,omitempty"`
	Remark               *string `json:"remark,omitempty" xml:"remark,omitempty"`
	UserCorpRelationType *string `json:"userCorpRelationType,omitempty" xml:"userCorpRelationType,omitempty"`
	UserIdentity         *string `json:"userIdentity,omitempty" xml:"userIdentity,omitempty"`
	VerifyEvent          *string `json:"verifyEvent,omitempty" xml:"verifyEvent,omitempty"`
	VerifyLocation       *string `json:"verifyLocation,omitempty" xml:"verifyLocation,omitempty"`
	VerifyNo             *string `json:"verifyNo,omitempty" xml:"verifyNo,omitempty"`
	VerifyResult         *bool   `json:"verifyResult,omitempty" xml:"verifyResult,omitempty"`
	VerifyTime           *string `json:"verifyTime,omitempty" xml:"verifyTime,omitempty"`
}

func (s NotifyBadgeCodeVerifyResultRequest) String() string {
	return tea.Prettify(s)
}

func (s NotifyBadgeCodeVerifyResultRequest) GoString() string {
	return s.String()
}

func (s *NotifyBadgeCodeVerifyResultRequest) SetCorpId(v string) *NotifyBadgeCodeVerifyResultRequest {
	s.CorpId = &v
	return s
}

func (s *NotifyBadgeCodeVerifyResultRequest) SetPayCode(v string) *NotifyBadgeCodeVerifyResultRequest {
	s.PayCode = &v
	return s
}

func (s *NotifyBadgeCodeVerifyResultRequest) SetRemark(v string) *NotifyBadgeCodeVerifyResultRequest {
	s.Remark = &v
	return s
}

func (s *NotifyBadgeCodeVerifyResultRequest) SetUserCorpRelationType(v string) *NotifyBadgeCodeVerifyResultRequest {
	s.UserCorpRelationType = &v
	return s
}

func (s *NotifyBadgeCodeVerifyResultRequest) SetUserIdentity(v string) *NotifyBadgeCodeVerifyResultRequest {
	s.UserIdentity = &v
	return s
}

func (s *NotifyBadgeCodeVerifyResultRequest) SetVerifyEvent(v string) *NotifyBadgeCodeVerifyResultRequest {
	s.VerifyEvent = &v
	return s
}

func (s *NotifyBadgeCodeVerifyResultRequest) SetVerifyLocation(v string) *NotifyBadgeCodeVerifyResultRequest {
	s.VerifyLocation = &v
	return s
}

func (s *NotifyBadgeCodeVerifyResultRequest) SetVerifyNo(v string) *NotifyBadgeCodeVerifyResultRequest {
	s.VerifyNo = &v
	return s
}

func (s *NotifyBadgeCodeVerifyResultRequest) SetVerifyResult(v bool) *NotifyBadgeCodeVerifyResultRequest {
	s.VerifyResult = &v
	return s
}

func (s *NotifyBadgeCodeVerifyResultRequest) SetVerifyTime(v string) *NotifyBadgeCodeVerifyResultRequest {
	s.VerifyTime = &v
	return s
}

type NotifyBadgeCodeVerifyResultResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s NotifyBadgeCodeVerifyResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NotifyBadgeCodeVerifyResultResponseBody) GoString() string {
	return s.String()
}

func (s *NotifyBadgeCodeVerifyResultResponseBody) SetResult(v string) *NotifyBadgeCodeVerifyResultResponseBody {
	s.Result = &v
	return s
}

type NotifyBadgeCodeVerifyResultResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NotifyBadgeCodeVerifyResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NotifyBadgeCodeVerifyResultResponse) String() string {
	return tea.Prettify(s)
}

func (s NotifyBadgeCodeVerifyResultResponse) GoString() string {
	return s.String()
}

func (s *NotifyBadgeCodeVerifyResultResponse) SetHeaders(v map[string]*string) *NotifyBadgeCodeVerifyResultResponse {
	s.Headers = v
	return s
}

func (s *NotifyBadgeCodeVerifyResultResponse) SetStatusCode(v int32) *NotifyBadgeCodeVerifyResultResponse {
	s.StatusCode = &v
	return s
}

func (s *NotifyBadgeCodeVerifyResultResponse) SetBody(v *NotifyBadgeCodeVerifyResultResponseBody) *NotifyBadgeCodeVerifyResultResponse {
	s.Body = v
	return s
}

type SaveBadgeCodeCorpInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveBadgeCodeCorpInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveBadgeCodeCorpInstanceHeaders) GoString() string {
	return s.String()
}

func (s *SaveBadgeCodeCorpInstanceHeaders) SetCommonHeaders(v map[string]*string) *SaveBadgeCodeCorpInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveBadgeCodeCorpInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *SaveBadgeCodeCorpInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveBadgeCodeCorpInstanceRequest struct {
	CodeIdentity *string            `json:"codeIdentity,omitempty" xml:"codeIdentity,omitempty"`
	CorpId       *string            `json:"corpId,omitempty" xml:"corpId,omitempty"`
	ExtInfo      map[string]*string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	Status       *string            `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SaveBadgeCodeCorpInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBadgeCodeCorpInstanceRequest) GoString() string {
	return s.String()
}

func (s *SaveBadgeCodeCorpInstanceRequest) SetCodeIdentity(v string) *SaveBadgeCodeCorpInstanceRequest {
	s.CodeIdentity = &v
	return s
}

func (s *SaveBadgeCodeCorpInstanceRequest) SetCorpId(v string) *SaveBadgeCodeCorpInstanceRequest {
	s.CorpId = &v
	return s
}

func (s *SaveBadgeCodeCorpInstanceRequest) SetExtInfo(v map[string]*string) *SaveBadgeCodeCorpInstanceRequest {
	s.ExtInfo = v
	return s
}

func (s *SaveBadgeCodeCorpInstanceRequest) SetStatus(v string) *SaveBadgeCodeCorpInstanceRequest {
	s.Status = &v
	return s
}

type SaveBadgeCodeCorpInstanceResponseBody struct {
	CodeIdentity *string            `json:"codeIdentity,omitempty" xml:"codeIdentity,omitempty"`
	CorpId       *string            `json:"corpId,omitempty" xml:"corpId,omitempty"`
	ExtInfo      map[string]*string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	Status       *string            `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SaveBadgeCodeCorpInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveBadgeCodeCorpInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBadgeCodeCorpInstanceResponseBody) SetCodeIdentity(v string) *SaveBadgeCodeCorpInstanceResponseBody {
	s.CodeIdentity = &v
	return s
}

func (s *SaveBadgeCodeCorpInstanceResponseBody) SetCorpId(v string) *SaveBadgeCodeCorpInstanceResponseBody {
	s.CorpId = &v
	return s
}

func (s *SaveBadgeCodeCorpInstanceResponseBody) SetExtInfo(v map[string]*string) *SaveBadgeCodeCorpInstanceResponseBody {
	s.ExtInfo = v
	return s
}

func (s *SaveBadgeCodeCorpInstanceResponseBody) SetStatus(v string) *SaveBadgeCodeCorpInstanceResponseBody {
	s.Status = &v
	return s
}

type SaveBadgeCodeCorpInstanceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveBadgeCodeCorpInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveBadgeCodeCorpInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBadgeCodeCorpInstanceResponse) GoString() string {
	return s.String()
}

func (s *SaveBadgeCodeCorpInstanceResponse) SetHeaders(v map[string]*string) *SaveBadgeCodeCorpInstanceResponse {
	s.Headers = v
	return s
}

func (s *SaveBadgeCodeCorpInstanceResponse) SetStatusCode(v int32) *SaveBadgeCodeCorpInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBadgeCodeCorpInstanceResponse) SetBody(v *SaveBadgeCodeCorpInstanceResponseBody) *SaveBadgeCodeCorpInstanceResponse {
	s.Body = v
	return s
}

type UpdateBadgeCodeUserInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateBadgeCodeUserInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateBadgeCodeUserInstanceHeaders) GoString() string {
	return s.String()
}

func (s *UpdateBadgeCodeUserInstanceHeaders) SetCommonHeaders(v map[string]*string) *UpdateBadgeCodeUserInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateBadgeCodeUserInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateBadgeCodeUserInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateBadgeCodeUserInstanceRequest struct {
	AvailableTimes       []*UpdateBadgeCodeUserInstanceRequestAvailableTimes `json:"availableTimes,omitempty" xml:"availableTimes,omitempty" type:"Repeated"`
	CodeId               *string                                             `json:"codeId,omitempty" xml:"codeId,omitempty"`
	CodeIdentity         *string                                             `json:"codeIdentity,omitempty" xml:"codeIdentity,omitempty"`
	CodeValue            *string                                             `json:"codeValue,omitempty" xml:"codeValue,omitempty"`
	CorpId               *string                                             `json:"corpId,omitempty" xml:"corpId,omitempty"`
	ExtInfo              map[string]interface{}                              `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	GmtExpired           *string                                             `json:"gmtExpired,omitempty" xml:"gmtExpired,omitempty"`
	Status               *string                                             `json:"status,omitempty" xml:"status,omitempty"`
	UserCorpRelationType *string                                             `json:"userCorpRelationType,omitempty" xml:"userCorpRelationType,omitempty"`
	UserIdentity         *string                                             `json:"userIdentity,omitempty" xml:"userIdentity,omitempty"`
}

func (s UpdateBadgeCodeUserInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBadgeCodeUserInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateBadgeCodeUserInstanceRequest) SetAvailableTimes(v []*UpdateBadgeCodeUserInstanceRequestAvailableTimes) *UpdateBadgeCodeUserInstanceRequest {
	s.AvailableTimes = v
	return s
}

func (s *UpdateBadgeCodeUserInstanceRequest) SetCodeId(v string) *UpdateBadgeCodeUserInstanceRequest {
	s.CodeId = &v
	return s
}

func (s *UpdateBadgeCodeUserInstanceRequest) SetCodeIdentity(v string) *UpdateBadgeCodeUserInstanceRequest {
	s.CodeIdentity = &v
	return s
}

func (s *UpdateBadgeCodeUserInstanceRequest) SetCodeValue(v string) *UpdateBadgeCodeUserInstanceRequest {
	s.CodeValue = &v
	return s
}

func (s *UpdateBadgeCodeUserInstanceRequest) SetCorpId(v string) *UpdateBadgeCodeUserInstanceRequest {
	s.CorpId = &v
	return s
}

func (s *UpdateBadgeCodeUserInstanceRequest) SetExtInfo(v map[string]interface{}) *UpdateBadgeCodeUserInstanceRequest {
	s.ExtInfo = v
	return s
}

func (s *UpdateBadgeCodeUserInstanceRequest) SetGmtExpired(v string) *UpdateBadgeCodeUserInstanceRequest {
	s.GmtExpired = &v
	return s
}

func (s *UpdateBadgeCodeUserInstanceRequest) SetStatus(v string) *UpdateBadgeCodeUserInstanceRequest {
	s.Status = &v
	return s
}

func (s *UpdateBadgeCodeUserInstanceRequest) SetUserCorpRelationType(v string) *UpdateBadgeCodeUserInstanceRequest {
	s.UserCorpRelationType = &v
	return s
}

func (s *UpdateBadgeCodeUserInstanceRequest) SetUserIdentity(v string) *UpdateBadgeCodeUserInstanceRequest {
	s.UserIdentity = &v
	return s
}

type UpdateBadgeCodeUserInstanceRequestAvailableTimes struct {
	GmtEnd   *string `json:"gmtEnd,omitempty" xml:"gmtEnd,omitempty"`
	GmtStart *string `json:"gmtStart,omitempty" xml:"gmtStart,omitempty"`
}

func (s UpdateBadgeCodeUserInstanceRequestAvailableTimes) String() string {
	return tea.Prettify(s)
}

func (s UpdateBadgeCodeUserInstanceRequestAvailableTimes) GoString() string {
	return s.String()
}

func (s *UpdateBadgeCodeUserInstanceRequestAvailableTimes) SetGmtEnd(v string) *UpdateBadgeCodeUserInstanceRequestAvailableTimes {
	s.GmtEnd = &v
	return s
}

func (s *UpdateBadgeCodeUserInstanceRequestAvailableTimes) SetGmtStart(v string) *UpdateBadgeCodeUserInstanceRequestAvailableTimes {
	s.GmtStart = &v
	return s
}

type UpdateBadgeCodeUserInstanceResponseBody struct {
	CodeId *string `json:"codeId,omitempty" xml:"codeId,omitempty"`
}

func (s UpdateBadgeCodeUserInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBadgeCodeUserInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBadgeCodeUserInstanceResponseBody) SetCodeId(v string) *UpdateBadgeCodeUserInstanceResponseBody {
	s.CodeId = &v
	return s
}

type UpdateBadgeCodeUserInstanceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBadgeCodeUserInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBadgeCodeUserInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBadgeCodeUserInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateBadgeCodeUserInstanceResponse) SetHeaders(v map[string]*string) *UpdateBadgeCodeUserInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateBadgeCodeUserInstanceResponse) SetStatusCode(v int32) *UpdateBadgeCodeUserInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBadgeCodeUserInstanceResponse) SetBody(v *UpdateBadgeCodeUserInstanceResponseBody) *UpdateBadgeCodeUserInstanceResponse {
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

func (client *Client) CreateBadgeCodeUserInstanceWithOptions(request *CreateBadgeCodeUserInstanceRequest, headers *CreateBadgeCodeUserInstanceHeaders, runtime *util.RuntimeOptions) (_result *CreateBadgeCodeUserInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AvailableTimes)) {
		body["availableTimes"] = request.AvailableTimes
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

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["extInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.GmtExpired)) {
		body["gmtExpired"] = request.GmtExpired
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserCorpRelationType)) {
		body["userCorpRelationType"] = request.UserCorpRelationType
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdentity)) {
		body["userIdentity"] = request.UserIdentity
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
		Action:      tea.String("CreateBadgeCodeUserInstance"),
		Version:     tea.String("badge_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/badge/codes/userInstances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBadgeCodeUserInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBadgeCodeUserInstance(request *CreateBadgeCodeUserInstanceRequest) (_result *CreateBadgeCodeUserInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateBadgeCodeUserInstanceHeaders{}
	_result = &CreateBadgeCodeUserInstanceResponse{}
	_body, _err := client.CreateBadgeCodeUserInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBadgeNotifyWithOptions(request *CreateBadgeNotifyRequest, headers *CreateBadgeNotifyHeaders, runtime *util.RuntimeOptions) (_result *CreateBadgeNotifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.MsgId)) {
		body["msgId"] = request.MsgId
	}

	if !tea.BoolValue(util.IsUnset(request.MsgType)) {
		body["msgType"] = request.MsgType
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
		Action:      tea.String("CreateBadgeNotify"),
		Version:     tea.String("badge_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/badge/notices"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBadgeNotifyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBadgeNotify(request *CreateBadgeNotifyRequest) (_result *CreateBadgeNotifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateBadgeNotifyHeaders{}
	_result = &CreateBadgeNotifyResponse{}
	_body, _err := client.CreateBadgeNotifyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DecodeBadgeCodeWithOptions(request *DecodeBadgeCodeRequest, headers *DecodeBadgeCodeHeaders, runtime *util.RuntimeOptions) (_result *DecodeBadgeCodeResponse, _err error) {
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
		Action:      tea.String("DecodeBadgeCode"),
		Version:     tea.String("badge_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/badge/codes/decode"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DecodeBadgeCodeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DecodeBadgeCode(request *DecodeBadgeCodeRequest) (_result *DecodeBadgeCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DecodeBadgeCodeHeaders{}
	_result = &DecodeBadgeCodeResponse{}
	_body, _err := client.DecodeBadgeCodeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) NotifyBadgeCodePayResultWithOptions(request *NotifyBadgeCodePayResultRequest, headers *NotifyBadgeCodePayResultHeaders, runtime *util.RuntimeOptions) (_result *NotifyBadgeCodePayResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		body["amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeAmount)) {
		body["chargeAmount"] = request.ChargeAmount
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["extInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.GmtTradeCreate)) {
		body["gmtTradeCreate"] = request.GmtTradeCreate
	}

	if !tea.BoolValue(util.IsUnset(request.GmtTradeFinish)) {
		body["gmtTradeFinish"] = request.GmtTradeFinish
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantName)) {
		body["merchantName"] = request.MerchantName
	}

	if !tea.BoolValue(util.IsUnset(request.PayChannelDetailList)) {
		body["payChannelDetailList"] = request.PayChannelDetailList
	}

	if !tea.BoolValue(util.IsUnset(request.PayCode)) {
		body["payCode"] = request.PayCode
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionAmount)) {
		body["promotionAmount"] = request.PromotionAmount
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.TradeErrorCode)) {
		body["tradeErrorCode"] = request.TradeErrorCode
	}

	if !tea.BoolValue(util.IsUnset(request.TradeErrorMsg)) {
		body["tradeErrorMsg"] = request.TradeErrorMsg
	}

	if !tea.BoolValue(util.IsUnset(request.TradeNo)) {
		body["tradeNo"] = request.TradeNo
	}

	if !tea.BoolValue(util.IsUnset(request.TradeStatus)) {
		body["tradeStatus"] = request.TradeStatus
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
		Action:      tea.String("NotifyBadgeCodePayResult"),
		Version:     tea.String("badge_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/badge/codes/payResults"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &NotifyBadgeCodePayResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) NotifyBadgeCodePayResult(request *NotifyBadgeCodePayResultRequest) (_result *NotifyBadgeCodePayResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &NotifyBadgeCodePayResultHeaders{}
	_result = &NotifyBadgeCodePayResultResponse{}
	_body, _err := client.NotifyBadgeCodePayResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) NotifyBadgeCodeRefundResultWithOptions(request *NotifyBadgeCodeRefundResultRequest, headers *NotifyBadgeCodeRefundResultHeaders, runtime *util.RuntimeOptions) (_result *NotifyBadgeCodeRefundResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.GmtRefund)) {
		body["gmtRefund"] = request.GmtRefund
	}

	if !tea.BoolValue(util.IsUnset(request.PayChannelDetailList)) {
		body["payChannelDetailList"] = request.PayChannelDetailList
	}

	if !tea.BoolValue(util.IsUnset(request.PayCode)) {
		body["payCode"] = request.PayCode
	}

	if !tea.BoolValue(util.IsUnset(request.RefundAmount)) {
		body["refundAmount"] = request.RefundAmount
	}

	if !tea.BoolValue(util.IsUnset(request.RefundOrderNo)) {
		body["refundOrderNo"] = request.RefundOrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.RefundPromotionAmount)) {
		body["refundPromotionAmount"] = request.RefundPromotionAmount
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.TradeNo)) {
		body["tradeNo"] = request.TradeNo
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
		Action:      tea.String("NotifyBadgeCodeRefundResult"),
		Version:     tea.String("badge_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/badge/codes/refundResults"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &NotifyBadgeCodeRefundResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) NotifyBadgeCodeRefundResult(request *NotifyBadgeCodeRefundResultRequest) (_result *NotifyBadgeCodeRefundResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &NotifyBadgeCodeRefundResultHeaders{}
	_result = &NotifyBadgeCodeRefundResultResponse{}
	_body, _err := client.NotifyBadgeCodeRefundResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) NotifyBadgeCodeVerifyResultWithOptions(request *NotifyBadgeCodeVerifyResultRequest, headers *NotifyBadgeCodeVerifyResultHeaders, runtime *util.RuntimeOptions) (_result *NotifyBadgeCodeVerifyResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.PayCode)) {
		body["payCode"] = request.PayCode
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.UserCorpRelationType)) {
		body["userCorpRelationType"] = request.UserCorpRelationType
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdentity)) {
		body["userIdentity"] = request.UserIdentity
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyEvent)) {
		body["verifyEvent"] = request.VerifyEvent
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyLocation)) {
		body["verifyLocation"] = request.VerifyLocation
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyNo)) {
		body["verifyNo"] = request.VerifyNo
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyResult)) {
		body["verifyResult"] = request.VerifyResult
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyTime)) {
		body["verifyTime"] = request.VerifyTime
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
		Action:      tea.String("NotifyBadgeCodeVerifyResult"),
		Version:     tea.String("badge_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/badge/codes/verifyResults"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &NotifyBadgeCodeVerifyResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) NotifyBadgeCodeVerifyResult(request *NotifyBadgeCodeVerifyResultRequest) (_result *NotifyBadgeCodeVerifyResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &NotifyBadgeCodeVerifyResultHeaders{}
	_result = &NotifyBadgeCodeVerifyResultResponse{}
	_body, _err := client.NotifyBadgeCodeVerifyResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveBadgeCodeCorpInstanceWithOptions(request *SaveBadgeCodeCorpInstanceRequest, headers *SaveBadgeCodeCorpInstanceHeaders, runtime *util.RuntimeOptions) (_result *SaveBadgeCodeCorpInstanceResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["extInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
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
		Action:      tea.String("SaveBadgeCodeCorpInstance"),
		Version:     tea.String("badge_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/badge/codes/corpInstances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveBadgeCodeCorpInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveBadgeCodeCorpInstance(request *SaveBadgeCodeCorpInstanceRequest) (_result *SaveBadgeCodeCorpInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveBadgeCodeCorpInstanceHeaders{}
	_result = &SaveBadgeCodeCorpInstanceResponse{}
	_body, _err := client.SaveBadgeCodeCorpInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateBadgeCodeUserInstanceWithOptions(request *UpdateBadgeCodeUserInstanceRequest, headers *UpdateBadgeCodeUserInstanceHeaders, runtime *util.RuntimeOptions) (_result *UpdateBadgeCodeUserInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AvailableTimes)) {
		body["availableTimes"] = request.AvailableTimes
	}

	if !tea.BoolValue(util.IsUnset(request.CodeId)) {
		body["codeId"] = request.CodeId
	}

	if !tea.BoolValue(util.IsUnset(request.CodeIdentity)) {
		body["codeIdentity"] = request.CodeIdentity
	}

	if !tea.BoolValue(util.IsUnset(request.CodeValue)) {
		body["codeValue"] = request.CodeValue
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["extInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.GmtExpired)) {
		body["gmtExpired"] = request.GmtExpired
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserCorpRelationType)) {
		body["userCorpRelationType"] = request.UserCorpRelationType
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdentity)) {
		body["userIdentity"] = request.UserIdentity
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
		Action:      tea.String("UpdateBadgeCodeUserInstance"),
		Version:     tea.String("badge_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/badge/codes/userInstances"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateBadgeCodeUserInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateBadgeCodeUserInstance(request *UpdateBadgeCodeUserInstanceRequest) (_result *UpdateBadgeCodeUserInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateBadgeCodeUserInstanceHeaders{}
	_result = &UpdateBadgeCodeUserInstanceResponse{}
	_body, _err := client.UpdateBadgeCodeUserInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
