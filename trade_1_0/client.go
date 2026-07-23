// This file is auto-generated, don't edit it. Thanks.
package trade_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
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
	// This parameter is required.
	BelongToPhoneNum *string `json:"belongToPhoneNum,omitempty" xml:"belongToPhoneNum,omitempty"`
	// This parameter is required.
	ContactPhoneNum *string `json:"contactPhoneNum,omitempty" xml:"contactPhoneNum,omitempty"`
	// This parameter is required.
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeptId *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// This parameter is required.
	MarketCode *string `json:"marketCode,omitempty" xml:"marketCode,omitempty"`
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
	// example:
	//
	// true
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

type CreateClueTempHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateClueTempHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateClueTempHeaders) GoString() string {
	return s.String()
}

func (s *CreateClueTempHeaders) SetCommonHeaders(v map[string]*string) *CreateClueTempHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateClueTempHeaders) SetXAcsDingtalkAccessToken(v string) *CreateClueTempHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateClueTempRequest struct {
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// This parameter is required.
	ContactName *string `json:"contactName,omitempty" xml:"contactName,omitempty"`
	DeptId      *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// if can be null:
	// false
	Ext *string `json:"ext,omitempty" xml:"ext,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	PhoneNum    *string `json:"phoneNum,omitempty" xml:"phoneNum,omitempty"`
	Position    *string `json:"position,omitempty" xml:"position,omitempty"`
	ProductCode *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
	SalesId     *int64  `json:"salesId,omitempty" xml:"salesId,omitempty"`
	// This parameter is required.
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s CreateClueTempRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClueTempRequest) GoString() string {
	return s.String()
}

func (s *CreateClueTempRequest) SetAddress(v string) *CreateClueTempRequest {
	s.Address = &v
	return s
}

func (s *CreateClueTempRequest) SetContactName(v string) *CreateClueTempRequest {
	s.ContactName = &v
	return s
}

func (s *CreateClueTempRequest) SetDeptId(v int64) *CreateClueTempRequest {
	s.DeptId = &v
	return s
}

func (s *CreateClueTempRequest) SetExt(v string) *CreateClueTempRequest {
	s.Ext = &v
	return s
}

func (s *CreateClueTempRequest) SetName(v string) *CreateClueTempRequest {
	s.Name = &v
	return s
}

func (s *CreateClueTempRequest) SetPhoneNum(v string) *CreateClueTempRequest {
	s.PhoneNum = &v
	return s
}

func (s *CreateClueTempRequest) SetPosition(v string) *CreateClueTempRequest {
	s.Position = &v
	return s
}

func (s *CreateClueTempRequest) SetProductCode(v string) *CreateClueTempRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateClueTempRequest) SetSalesId(v int64) *CreateClueTempRequest {
	s.SalesId = &v
	return s
}

func (s *CreateClueTempRequest) SetSource(v string) *CreateClueTempRequest {
	s.Source = &v
	return s
}

type CreateClueTempResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateClueTempResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClueTempResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClueTempResponseBody) SetSuccess(v bool) *CreateClueTempResponseBody {
	s.Success = &v
	return s
}

type CreateClueTempResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClueTempResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateClueTempResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClueTempResponse) GoString() string {
	return s.String()
}

func (s *CreateClueTempResponse) SetHeaders(v map[string]*string) *CreateClueTempResponse {
	s.Headers = v
	return s
}

func (s *CreateClueTempResponse) SetStatusCode(v int32) *CreateClueTempResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClueTempResponse) SetBody(v *CreateClueTempResponseBody) *CreateClueTempResponse {
	s.Body = v
	return s
}

type CreateNoteForIsvHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateNoteForIsvHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateNoteForIsvHeaders) GoString() string {
	return s.String()
}

func (s *CreateNoteForIsvHeaders) SetCommonHeaders(v map[string]*string) *CreateNoteForIsvHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateNoteForIsvHeaders) SetXAcsDingtalkAccessToken(v string) *CreateNoteForIsvHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateNoteForIsvRequest struct {
	ContactName *string `json:"contactName,omitempty" xml:"contactName,omitempty"`
	// This parameter is required.
	ContactPhoneNum *string `json:"contactPhoneNum,omitempty" xml:"contactPhoneNum,omitempty"`
	ContactTitle    *string `json:"contactTitle,omitempty" xml:"contactTitle,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// This parameter is required.
	InputPhoneNum *string `json:"inputPhoneNum,omitempty" xml:"inputPhoneNum,omitempty"`
}

func (s CreateNoteForIsvRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNoteForIsvRequest) GoString() string {
	return s.String()
}

func (s *CreateNoteForIsvRequest) SetContactName(v string) *CreateNoteForIsvRequest {
	s.ContactName = &v
	return s
}

func (s *CreateNoteForIsvRequest) SetContactPhoneNum(v string) *CreateNoteForIsvRequest {
	s.ContactPhoneNum = &v
	return s
}

func (s *CreateNoteForIsvRequest) SetContactTitle(v string) *CreateNoteForIsvRequest {
	s.ContactTitle = &v
	return s
}

func (s *CreateNoteForIsvRequest) SetContent(v string) *CreateNoteForIsvRequest {
	s.Content = &v
	return s
}

func (s *CreateNoteForIsvRequest) SetCorpId(v string) *CreateNoteForIsvRequest {
	s.CorpId = &v
	return s
}

func (s *CreateNoteForIsvRequest) SetInputPhoneNum(v string) *CreateNoteForIsvRequest {
	s.InputPhoneNum = &v
	return s
}

type CreateNoteForIsvResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateNoteForIsvResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNoteForIsvResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNoteForIsvResponseBody) SetSuccess(v bool) *CreateNoteForIsvResponseBody {
	s.Success = &v
	return s
}

type CreateNoteForIsvResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNoteForIsvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNoteForIsvResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNoteForIsvResponse) GoString() string {
	return s.String()
}

func (s *CreateNoteForIsvResponse) SetHeaders(v map[string]*string) *CreateNoteForIsvResponse {
	s.Headers = v
	return s
}

func (s *CreateNoteForIsvResponse) SetStatusCode(v int32) *CreateNoteForIsvResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNoteForIsvResponse) SetBody(v *CreateNoteForIsvResponseBody) *CreateNoteForIsvResponse {
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
	// This parameter is required.
	BelongToPhoneNum *string `json:"belongToPhoneNum,omitempty" xml:"belongToPhoneNum,omitempty"`
	// This parameter is required.
	ContactPhoneNum *string `json:"contactPhoneNum,omitempty" xml:"contactPhoneNum,omitempty"`
	// This parameter is required.
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeptId *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// This parameter is required.
	MarketCode *string `json:"marketCode,omitempty" xml:"marketCode,omitempty"`
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

type PurchaseMixViewHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PurchaseMixViewHeaders) String() string {
	return tea.Prettify(s)
}

func (s PurchaseMixViewHeaders) GoString() string {
	return s.String()
}

func (s *PurchaseMixViewHeaders) SetCommonHeaders(v map[string]*string) *PurchaseMixViewHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PurchaseMixViewHeaders) SetXAcsDingtalkAccessToken(v string) *PurchaseMixViewHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PurchaseMixViewRequest struct {
	ChannelCode       *string                       `json:"channelCode,omitempty" xml:"channelCode,omitempty"`
	CombineActivityId *int64                        `json:"combineActivityId,omitempty" xml:"combineActivityId,omitempty"`
	CreateOrder       *bool                         `json:"createOrder,omitempty" xml:"createOrder,omitempty"`
	List              []*PurchaseMixViewRequestList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	Memo              *string                       `json:"memo,omitempty" xml:"memo,omitempty"`
	MergeOrderName    *string                       `json:"mergeOrderName,omitempty" xml:"mergeOrderName,omitempty"`
	NeedModelTypeList []*string                     `json:"needModelTypeList,omitempty" xml:"needModelTypeList,omitempty" type:"Repeated"`
	ObjId             *int64                        `json:"objId,omitempty" xml:"objId,omitempty"`
	ObjType           *int64                        `json:"objType,omitempty" xml:"objType,omitempty"`
	OrderParamMap     map[string]interface{}        `json:"orderParamMap,omitempty" xml:"orderParamMap,omitempty"`
	OuterTradeCode    *string                       `json:"outerTradeCode,omitempty" xml:"outerTradeCode,omitempty"`
	OuterTradeType    *string                       `json:"outerTradeType,omitempty" xml:"outerTradeType,omitempty"`
	PlanId            *int64                        `json:"planId,omitempty" xml:"planId,omitempty"`
	RequestId         *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Uid               *int64                        `json:"uid,omitempty" xml:"uid,omitempty"`
	UnPay             *bool                         `json:"unPay,omitempty" xml:"unPay,omitempty"`
}

func (s PurchaseMixViewRequest) String() string {
	return tea.Prettify(s)
}

func (s PurchaseMixViewRequest) GoString() string {
	return s.String()
}

func (s *PurchaseMixViewRequest) SetChannelCode(v string) *PurchaseMixViewRequest {
	s.ChannelCode = &v
	return s
}

func (s *PurchaseMixViewRequest) SetCombineActivityId(v int64) *PurchaseMixViewRequest {
	s.CombineActivityId = &v
	return s
}

func (s *PurchaseMixViewRequest) SetCreateOrder(v bool) *PurchaseMixViewRequest {
	s.CreateOrder = &v
	return s
}

func (s *PurchaseMixViewRequest) SetList(v []*PurchaseMixViewRequestList) *PurchaseMixViewRequest {
	s.List = v
	return s
}

func (s *PurchaseMixViewRequest) SetMemo(v string) *PurchaseMixViewRequest {
	s.Memo = &v
	return s
}

func (s *PurchaseMixViewRequest) SetMergeOrderName(v string) *PurchaseMixViewRequest {
	s.MergeOrderName = &v
	return s
}

func (s *PurchaseMixViewRequest) SetNeedModelTypeList(v []*string) *PurchaseMixViewRequest {
	s.NeedModelTypeList = v
	return s
}

func (s *PurchaseMixViewRequest) SetObjId(v int64) *PurchaseMixViewRequest {
	s.ObjId = &v
	return s
}

func (s *PurchaseMixViewRequest) SetObjType(v int64) *PurchaseMixViewRequest {
	s.ObjType = &v
	return s
}

func (s *PurchaseMixViewRequest) SetOrderParamMap(v map[string]interface{}) *PurchaseMixViewRequest {
	s.OrderParamMap = v
	return s
}

func (s *PurchaseMixViewRequest) SetOuterTradeCode(v string) *PurchaseMixViewRequest {
	s.OuterTradeCode = &v
	return s
}

func (s *PurchaseMixViewRequest) SetOuterTradeType(v string) *PurchaseMixViewRequest {
	s.OuterTradeType = &v
	return s
}

func (s *PurchaseMixViewRequest) SetPlanId(v int64) *PurchaseMixViewRequest {
	s.PlanId = &v
	return s
}

func (s *PurchaseMixViewRequest) SetRequestId(v string) *PurchaseMixViewRequest {
	s.RequestId = &v
	return s
}

func (s *PurchaseMixViewRequest) SetUid(v int64) *PurchaseMixViewRequest {
	s.Uid = &v
	return s
}

func (s *PurchaseMixViewRequest) SetUnPay(v bool) *PurchaseMixViewRequest {
	s.UnPay = &v
	return s
}

type PurchaseMixViewRequestList struct {
	ActivityId         *int64                                          `json:"activityId,omitempty" xml:"activityId,omitempty"`
	ArticleCode        *string                                         `json:"articleCode,omitempty" xml:"articleCode,omitempty"`
	ArticleItemCode    *string                                         `json:"articleItemCode,omitempty" xml:"articleItemCode,omitempty"`
	ArticleItemId      *int64                                          `json:"articleItemId,omitempty" xml:"articleItemId,omitempty"`
	ArticleItemName    *string                                         `json:"articleItemName,omitempty" xml:"articleItemName,omitempty"`
	BizType            *int64                                          `json:"bizType,omitempty" xml:"bizType,omitempty"`
	CouponId           *int64                                          `json:"couponId,omitempty" xml:"couponId,omitempty"`
	CycNum             *int64                                          `json:"cycNum,omitempty" xml:"cycNum,omitempty"`
	CycType            *int64                                          `json:"cycType,omitempty" xml:"cycType,omitempty"`
	CycUnit            *int64                                          `json:"cycUnit,omitempty" xml:"cycUnit,omitempty"`
	Extend1            *int64                                          `json:"extend1,omitempty" xml:"extend1,omitempty"`
	InstanceNum        *int64                                          `json:"instanceNum,omitempty" xml:"instanceNum,omitempty"`
	IsCredit           *bool                                           `json:"isCredit,omitempty" xml:"isCredit,omitempty"`
	ItemComponentList  []*PurchaseMixViewRequestListItemComponentList  `json:"itemComponentList,omitempty" xml:"itemComponentList,omitempty" type:"Repeated"`
	MinorItemParamList []*PurchaseMixViewRequestListMinorItemParamList `json:"minorItemParamList,omitempty" xml:"minorItemParamList,omitempty" type:"Repeated"`
	OrderParamMap      map[string]interface{}                          `json:"orderParamMap,omitempty" xml:"orderParamMap,omitempty"`
	OrderType          *string                                         `json:"orderType,omitempty" xml:"orderType,omitempty"`
	SaleMarketType     *string                                         `json:"saleMarketType,omitempty" xml:"saleMarketType,omitempty"`
	SaleOrgId          *int64                                          `json:"saleOrgId,omitempty" xml:"saleOrgId,omitempty"`
	SubQuantity        *int64                                          `json:"subQuantity,omitempty" xml:"subQuantity,omitempty"`
	TradeModelType     *string                                         `json:"tradeModelType,omitempty" xml:"tradeModelType,omitempty"`
}

func (s PurchaseMixViewRequestList) String() string {
	return tea.Prettify(s)
}

func (s PurchaseMixViewRequestList) GoString() string {
	return s.String()
}

func (s *PurchaseMixViewRequestList) SetActivityId(v int64) *PurchaseMixViewRequestList {
	s.ActivityId = &v
	return s
}

func (s *PurchaseMixViewRequestList) SetArticleCode(v string) *PurchaseMixViewRequestList {
	s.ArticleCode = &v
	return s
}

func (s *PurchaseMixViewRequestList) SetArticleItemCode(v string) *PurchaseMixViewRequestList {
	s.ArticleItemCode = &v
	return s
}

func (s *PurchaseMixViewRequestList) SetArticleItemId(v int64) *PurchaseMixViewRequestList {
	s.ArticleItemId = &v
	return s
}

func (s *PurchaseMixViewRequestList) SetArticleItemName(v string) *PurchaseMixViewRequestList {
	s.ArticleItemName = &v
	return s
}

func (s *PurchaseMixViewRequestList) SetBizType(v int64) *PurchaseMixViewRequestList {
	s.BizType = &v
	return s
}

func (s *PurchaseMixViewRequestList) SetCouponId(v int64) *PurchaseMixViewRequestList {
	s.CouponId = &v
	return s
}

func (s *PurchaseMixViewRequestList) SetCycNum(v int64) *PurchaseMixViewRequestList {
	s.CycNum = &v
	return s
}

func (s *PurchaseMixViewRequestList) SetCycType(v int64) *PurchaseMixViewRequestList {
	s.CycType = &v
	return s
}

func (s *PurchaseMixViewRequestList) SetCycUnit(v int64) *PurchaseMixViewRequestList {
	s.CycUnit = &v
	return s
}

func (s *PurchaseMixViewRequestList) SetExtend1(v int64) *PurchaseMixViewRequestList {
	s.Extend1 = &v
	return s
}

func (s *PurchaseMixViewRequestList) SetInstanceNum(v int64) *PurchaseMixViewRequestList {
	s.InstanceNum = &v
	return s
}

func (s *PurchaseMixViewRequestList) SetIsCredit(v bool) *PurchaseMixViewRequestList {
	s.IsCredit = &v
	return s
}

func (s *PurchaseMixViewRequestList) SetItemComponentList(v []*PurchaseMixViewRequestListItemComponentList) *PurchaseMixViewRequestList {
	s.ItemComponentList = v
	return s
}

func (s *PurchaseMixViewRequestList) SetMinorItemParamList(v []*PurchaseMixViewRequestListMinorItemParamList) *PurchaseMixViewRequestList {
	s.MinorItemParamList = v
	return s
}

func (s *PurchaseMixViewRequestList) SetOrderParamMap(v map[string]interface{}) *PurchaseMixViewRequestList {
	s.OrderParamMap = v
	return s
}

func (s *PurchaseMixViewRequestList) SetOrderType(v string) *PurchaseMixViewRequestList {
	s.OrderType = &v
	return s
}

func (s *PurchaseMixViewRequestList) SetSaleMarketType(v string) *PurchaseMixViewRequestList {
	s.SaleMarketType = &v
	return s
}

func (s *PurchaseMixViewRequestList) SetSaleOrgId(v int64) *PurchaseMixViewRequestList {
	s.SaleOrgId = &v
	return s
}

func (s *PurchaseMixViewRequestList) SetSubQuantity(v int64) *PurchaseMixViewRequestList {
	s.SubQuantity = &v
	return s
}

func (s *PurchaseMixViewRequestList) SetTradeModelType(v string) *PurchaseMixViewRequestList {
	s.TradeModelType = &v
	return s
}

type PurchaseMixViewRequestListItemComponentList struct {
	ComponentCode    *string                                                        `json:"componentCode,omitempty" xml:"componentCode,omitempty"`
	ComponentName    *string                                                        `json:"componentName,omitempty" xml:"componentName,omitempty"`
	ItemPropertyList []*PurchaseMixViewRequestListItemComponentListItemPropertyList `json:"itemPropertyList,omitempty" xml:"itemPropertyList,omitempty" type:"Repeated"`
}

func (s PurchaseMixViewRequestListItemComponentList) String() string {
	return tea.Prettify(s)
}

func (s PurchaseMixViewRequestListItemComponentList) GoString() string {
	return s.String()
}

func (s *PurchaseMixViewRequestListItemComponentList) SetComponentCode(v string) *PurchaseMixViewRequestListItemComponentList {
	s.ComponentCode = &v
	return s
}

func (s *PurchaseMixViewRequestListItemComponentList) SetComponentName(v string) *PurchaseMixViewRequestListItemComponentList {
	s.ComponentName = &v
	return s
}

func (s *PurchaseMixViewRequestListItemComponentList) SetItemPropertyList(v []*PurchaseMixViewRequestListItemComponentListItemPropertyList) *PurchaseMixViewRequestListItemComponentList {
	s.ItemPropertyList = v
	return s
}

type PurchaseMixViewRequestListItemComponentListItemPropertyList struct {
	Code  *string `json:"code,omitempty" xml:"code,omitempty"`
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s PurchaseMixViewRequestListItemComponentListItemPropertyList) String() string {
	return tea.Prettify(s)
}

func (s PurchaseMixViewRequestListItemComponentListItemPropertyList) GoString() string {
	return s.String()
}

func (s *PurchaseMixViewRequestListItemComponentListItemPropertyList) SetCode(v string) *PurchaseMixViewRequestListItemComponentListItemPropertyList {
	s.Code = &v
	return s
}

func (s *PurchaseMixViewRequestListItemComponentListItemPropertyList) SetName(v string) *PurchaseMixViewRequestListItemComponentListItemPropertyList {
	s.Name = &v
	return s
}

func (s *PurchaseMixViewRequestListItemComponentListItemPropertyList) SetValue(v string) *PurchaseMixViewRequestListItemComponentListItemPropertyList {
	s.Value = &v
	return s
}

type PurchaseMixViewRequestListMinorItemParamList struct {
	MinorItemCode      *string `json:"minorItemCode,omitempty" xml:"minorItemCode,omitempty"`
	MinorItemGroupCode *string `json:"minorItemGroupCode,omitempty" xml:"minorItemGroupCode,omitempty"`
}

func (s PurchaseMixViewRequestListMinorItemParamList) String() string {
	return tea.Prettify(s)
}

func (s PurchaseMixViewRequestListMinorItemParamList) GoString() string {
	return s.String()
}

func (s *PurchaseMixViewRequestListMinorItemParamList) SetMinorItemCode(v string) *PurchaseMixViewRequestListMinorItemParamList {
	s.MinorItemCode = &v
	return s
}

func (s *PurchaseMixViewRequestListMinorItemParamList) SetMinorItemGroupCode(v string) *PurchaseMixViewRequestListMinorItemParamList {
	s.MinorItemGroupCode = &v
	return s
}

type PurchaseMixViewResponseBody struct {
	ErrorCode *string                            `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                            `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result    *PurchaseMixViewResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success   *bool                              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PurchaseMixViewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PurchaseMixViewResponseBody) GoString() string {
	return s.String()
}

func (s *PurchaseMixViewResponseBody) SetErrorCode(v string) *PurchaseMixViewResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PurchaseMixViewResponseBody) SetErrorMsg(v string) *PurchaseMixViewResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *PurchaseMixViewResponseBody) SetResult(v *PurchaseMixViewResponseBodyResult) *PurchaseMixViewResponseBody {
	s.Result = v
	return s
}

func (s *PurchaseMixViewResponseBody) SetSuccess(v bool) *PurchaseMixViewResponseBody {
	s.Success = &v
	return s
}

type PurchaseMixViewResponseBodyResult struct {
	AliyunArticleItemViewUnitList      []*PurchaseMixViewResponseBodyResultAliyunArticleItemViewUnitList    `json:"aliyunArticleItemViewUnitList,omitempty" xml:"aliyunArticleItemViewUnitList,omitempty" type:"Repeated"`
	MixPromotionVO                     *PurchaseMixViewResponseBodyResultMixPromotionVO                     `json:"mixPromotionVO,omitempty" xml:"mixPromotionVO,omitempty" type:"Struct"`
	RecommendedMarketingCollocationDTO *PurchaseMixViewResponseBodyResultRecommendedMarketingCollocationDTO `json:"recommendedMarketingCollocationDTO,omitempty" xml:"recommendedMarketingCollocationDTO,omitempty" type:"Struct"`
}

func (s PurchaseMixViewResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PurchaseMixViewResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PurchaseMixViewResponseBodyResult) SetAliyunArticleItemViewUnitList(v []*PurchaseMixViewResponseBodyResultAliyunArticleItemViewUnitList) *PurchaseMixViewResponseBodyResult {
	s.AliyunArticleItemViewUnitList = v
	return s
}

func (s *PurchaseMixViewResponseBodyResult) SetMixPromotionVO(v *PurchaseMixViewResponseBodyResultMixPromotionVO) *PurchaseMixViewResponseBodyResult {
	s.MixPromotionVO = v
	return s
}

func (s *PurchaseMixViewResponseBodyResult) SetRecommendedMarketingCollocationDTO(v *PurchaseMixViewResponseBodyResultRecommendedMarketingCollocationDTO) *PurchaseMixViewResponseBodyResult {
	s.RecommendedMarketingCollocationDTO = v
	return s
}

type PurchaseMixViewResponseBodyResultAliyunArticleItemViewUnitList struct {
	ActualPayFee    *float64  `json:"actualPayFee,omitempty" xml:"actualPayFee,omitempty"`
	ArticleCode     *string   `json:"articleCode,omitempty" xml:"articleCode,omitempty"`
	ArticleItemCode *string   `json:"articleItemCode,omitempty" xml:"articleItemCode,omitempty"`
	BizTagList      []*string `json:"bizTagList,omitempty" xml:"bizTagList,omitempty" type:"Repeated"`
	EndDate         *int64    `json:"endDate,omitempty" xml:"endDate,omitempty"`
	OrderType       *string   `json:"orderType,omitempty" xml:"orderType,omitempty"`
	PayFee          *float64  `json:"payFee,omitempty" xml:"payFee,omitempty"`
	StartDate       *int64    `json:"startDate,omitempty" xml:"startDate,omitempty"`
}

func (s PurchaseMixViewResponseBodyResultAliyunArticleItemViewUnitList) String() string {
	return tea.Prettify(s)
}

func (s PurchaseMixViewResponseBodyResultAliyunArticleItemViewUnitList) GoString() string {
	return s.String()
}

func (s *PurchaseMixViewResponseBodyResultAliyunArticleItemViewUnitList) SetActualPayFee(v float64) *PurchaseMixViewResponseBodyResultAliyunArticleItemViewUnitList {
	s.ActualPayFee = &v
	return s
}

func (s *PurchaseMixViewResponseBodyResultAliyunArticleItemViewUnitList) SetArticleCode(v string) *PurchaseMixViewResponseBodyResultAliyunArticleItemViewUnitList {
	s.ArticleCode = &v
	return s
}

func (s *PurchaseMixViewResponseBodyResultAliyunArticleItemViewUnitList) SetArticleItemCode(v string) *PurchaseMixViewResponseBodyResultAliyunArticleItemViewUnitList {
	s.ArticleItemCode = &v
	return s
}

func (s *PurchaseMixViewResponseBodyResultAliyunArticleItemViewUnitList) SetBizTagList(v []*string) *PurchaseMixViewResponseBodyResultAliyunArticleItemViewUnitList {
	s.BizTagList = v
	return s
}

func (s *PurchaseMixViewResponseBodyResultAliyunArticleItemViewUnitList) SetEndDate(v int64) *PurchaseMixViewResponseBodyResultAliyunArticleItemViewUnitList {
	s.EndDate = &v
	return s
}

func (s *PurchaseMixViewResponseBodyResultAliyunArticleItemViewUnitList) SetOrderType(v string) *PurchaseMixViewResponseBodyResultAliyunArticleItemViewUnitList {
	s.OrderType = &v
	return s
}

func (s *PurchaseMixViewResponseBodyResultAliyunArticleItemViewUnitList) SetPayFee(v float64) *PurchaseMixViewResponseBodyResultAliyunArticleItemViewUnitList {
	s.PayFee = &v
	return s
}

func (s *PurchaseMixViewResponseBodyResultAliyunArticleItemViewUnitList) SetStartDate(v int64) *PurchaseMixViewResponseBodyResultAliyunArticleItemViewUnitList {
	s.StartDate = &v
	return s
}

type PurchaseMixViewResponseBodyResultMixPromotionVO struct {
	ActivityId       *int64                 `json:"activityId,omitempty" xml:"activityId,omitempty"`
	ActivityName     *string                `json:"activityName,omitempty" xml:"activityName,omitempty"`
	ActivityUrl      *string                `json:"activityUrl,omitempty" xml:"activityUrl,omitempty"`
	EndTime          *int64                 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ExtendParam      map[string]interface{} `json:"extendParam,omitempty" xml:"extendParam,omitempty"`
	ForbiddenCoupon  *bool                  `json:"forbiddenCoupon,omitempty" xml:"forbiddenCoupon,omitempty"`
	IsSelected       *bool                  `json:"isSelected,omitempty" xml:"isSelected,omitempty"`
	MarketActivityId *int64                 `json:"marketActivityId,omitempty" xml:"marketActivityId,omitempty"`
	PromotionId      *int64                 `json:"promotionId,omitempty" xml:"promotionId,omitempty"`
	PromotionName    *string                `json:"promotionName,omitempty" xml:"promotionName,omitempty"`
	PromotionType    *string                `json:"promotionType,omitempty" xml:"promotionType,omitempty"`
	Source           *string                `json:"source,omitempty" xml:"source,omitempty"`
	StartTime        *int64                 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Type             *string                `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PurchaseMixViewResponseBodyResultMixPromotionVO) String() string {
	return tea.Prettify(s)
}

func (s PurchaseMixViewResponseBodyResultMixPromotionVO) GoString() string {
	return s.String()
}

func (s *PurchaseMixViewResponseBodyResultMixPromotionVO) SetActivityId(v int64) *PurchaseMixViewResponseBodyResultMixPromotionVO {
	s.ActivityId = &v
	return s
}

func (s *PurchaseMixViewResponseBodyResultMixPromotionVO) SetActivityName(v string) *PurchaseMixViewResponseBodyResultMixPromotionVO {
	s.ActivityName = &v
	return s
}

func (s *PurchaseMixViewResponseBodyResultMixPromotionVO) SetActivityUrl(v string) *PurchaseMixViewResponseBodyResultMixPromotionVO {
	s.ActivityUrl = &v
	return s
}

func (s *PurchaseMixViewResponseBodyResultMixPromotionVO) SetEndTime(v int64) *PurchaseMixViewResponseBodyResultMixPromotionVO {
	s.EndTime = &v
	return s
}

func (s *PurchaseMixViewResponseBodyResultMixPromotionVO) SetExtendParam(v map[string]interface{}) *PurchaseMixViewResponseBodyResultMixPromotionVO {
	s.ExtendParam = v
	return s
}

func (s *PurchaseMixViewResponseBodyResultMixPromotionVO) SetForbiddenCoupon(v bool) *PurchaseMixViewResponseBodyResultMixPromotionVO {
	s.ForbiddenCoupon = &v
	return s
}

func (s *PurchaseMixViewResponseBodyResultMixPromotionVO) SetIsSelected(v bool) *PurchaseMixViewResponseBodyResultMixPromotionVO {
	s.IsSelected = &v
	return s
}

func (s *PurchaseMixViewResponseBodyResultMixPromotionVO) SetMarketActivityId(v int64) *PurchaseMixViewResponseBodyResultMixPromotionVO {
	s.MarketActivityId = &v
	return s
}

func (s *PurchaseMixViewResponseBodyResultMixPromotionVO) SetPromotionId(v int64) *PurchaseMixViewResponseBodyResultMixPromotionVO {
	s.PromotionId = &v
	return s
}

func (s *PurchaseMixViewResponseBodyResultMixPromotionVO) SetPromotionName(v string) *PurchaseMixViewResponseBodyResultMixPromotionVO {
	s.PromotionName = &v
	return s
}

func (s *PurchaseMixViewResponseBodyResultMixPromotionVO) SetPromotionType(v string) *PurchaseMixViewResponseBodyResultMixPromotionVO {
	s.PromotionType = &v
	return s
}

func (s *PurchaseMixViewResponseBodyResultMixPromotionVO) SetSource(v string) *PurchaseMixViewResponseBodyResultMixPromotionVO {
	s.Source = &v
	return s
}

func (s *PurchaseMixViewResponseBodyResultMixPromotionVO) SetStartTime(v int64) *PurchaseMixViewResponseBodyResultMixPromotionVO {
	s.StartTime = &v
	return s
}

func (s *PurchaseMixViewResponseBodyResultMixPromotionVO) SetType(v string) *PurchaseMixViewResponseBodyResultMixPromotionVO {
	s.Type = &v
	return s
}

type PurchaseMixViewResponseBodyResultRecommendedMarketingCollocationDTO struct {
	ActivityId *int64 `json:"activityId,omitempty" xml:"activityId,omitempty"`
	CouponId   *int64 `json:"couponId,omitempty" xml:"couponId,omitempty"`
}

func (s PurchaseMixViewResponseBodyResultRecommendedMarketingCollocationDTO) String() string {
	return tea.Prettify(s)
}

func (s PurchaseMixViewResponseBodyResultRecommendedMarketingCollocationDTO) GoString() string {
	return s.String()
}

func (s *PurchaseMixViewResponseBodyResultRecommendedMarketingCollocationDTO) SetActivityId(v int64) *PurchaseMixViewResponseBodyResultRecommendedMarketingCollocationDTO {
	s.ActivityId = &v
	return s
}

func (s *PurchaseMixViewResponseBodyResultRecommendedMarketingCollocationDTO) SetCouponId(v int64) *PurchaseMixViewResponseBodyResultRecommendedMarketingCollocationDTO {
	s.CouponId = &v
	return s
}

type PurchaseMixViewResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PurchaseMixViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PurchaseMixViewResponse) String() string {
	return tea.Prettify(s)
}

func (s PurchaseMixViewResponse) GoString() string {
	return s.String()
}

func (s *PurchaseMixViewResponse) SetHeaders(v map[string]*string) *PurchaseMixViewResponse {
	s.Headers = v
	return s
}

func (s *PurchaseMixViewResponse) SetStatusCode(v int32) *PurchaseMixViewResponse {
	s.StatusCode = &v
	return s
}

func (s *PurchaseMixViewResponse) SetBody(v *PurchaseMixViewResponseBody) *PurchaseMixViewResponse {
	s.Body = v
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
	// This parameter is required.
	ArticleCode *string `json:"articleCode,omitempty" xml:"articleCode,omitempty"`
	// This parameter is required.
	ArticleItemCode *string `json:"articleItemCode,omitempty" xml:"articleItemCode,omitempty"`
	// This parameter is required.
	ArticleItemName *string `json:"articleItemName,omitempty" xml:"articleItemName,omitempty"`
	// This parameter is required.
	ArticleName *string `json:"articleName,omitempty" xml:"articleName,omitempty"`
	CloseTime   *int64  `json:"closeTime,omitempty" xml:"closeTime,omitempty"`
	// This parameter is required.
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// This parameter is required.
	Fee *int64 `json:"fee,omitempty" xml:"fee,omitempty"`
	// This parameter is required.
	IsvCropId *string `json:"isvCropId,omitempty" xml:"isvCropId,omitempty"`
	// This parameter is required.
	OrderId      *int64  `json:"orderId,omitempty" xml:"orderId,omitempty"`
	OuterOrderId *string `json:"outerOrderId,omitempty" xml:"outerOrderId,omitempty"`
	// This parameter is required.
	PayFee  *int64 `json:"payFee,omitempty" xml:"payFee,omitempty"`
	PayTime *int64 `json:"payTime,omitempty" xml:"payTime,omitempty"`
	// This parameter is required.
	Quantity   *int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	RefundTime *int64 `json:"refundTime,omitempty" xml:"refundTime,omitempty"`
	// This parameter is required.
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
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
// isv检查商机创建是否符合预期
//
// @param request - CheckOpportunityResultRequest
//
// @param headers - CheckOpportunityResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckOpportunityResultResponse
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

// Summary:
//
// isv检查商机创建是否符合预期
//
// @param request - CheckOpportunityResultRequest
//
// @return CheckOpportunityResultResponse
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

// Summary:
//
// 用于客户跟进线索创建
//
// @param request - CreateClueTempRequest
//
// @param headers - CreateClueTempHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClueTempResponse
func (client *Client) CreateClueTempWithOptions(request *CreateClueTempRequest, headers *CreateClueTempHeaders, runtime *util.RuntimeOptions) (_result *CreateClueTempResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		body["address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		body["contactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNum)) {
		body["phoneNum"] = request.PhoneNum
	}

	if !tea.BoolValue(util.IsUnset(request.Position)) {
		body["position"] = request.Position
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["productCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.SalesId)) {
		body["salesId"] = request.SalesId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
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
		Action:      tea.String("CreateClueTemp"),
		Version:     tea.String("trade_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/trade/clueTemps"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateClueTempResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用于客户跟进线索创建
//
// @param request - CreateClueTempRequest
//
// @return CreateClueTempResponse
func (client *Client) CreateClueTemp(request *CreateClueTempRequest) (_result *CreateClueTempResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateClueTempHeaders{}
	_result = &CreateClueTempResponse{}
	_body, _err := client.CreateClueTempWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建小记
//
// @param request - CreateNoteForIsvRequest
//
// @param headers - CreateNoteForIsvHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNoteForIsvResponse
func (client *Client) CreateNoteForIsvWithOptions(request *CreateNoteForIsvRequest, headers *CreateNoteForIsvHeaders, runtime *util.RuntimeOptions) (_result *CreateNoteForIsvResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		body["contactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.ContactPhoneNum)) {
		body["contactPhoneNum"] = request.ContactPhoneNum
	}

	if !tea.BoolValue(util.IsUnset(request.ContactTitle)) {
		body["contactTitle"] = request.ContactTitle
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.InputPhoneNum)) {
		body["inputPhoneNum"] = request.InputPhoneNum
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
		Action:      tea.String("CreateNoteForIsv"),
		Version:     tea.String("trade_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/trade/notes"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNoteForIsvResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建小记
//
// @param request - CreateNoteForIsvRequest
//
// @return CreateNoteForIsvResponse
func (client *Client) CreateNoteForIsv(request *CreateNoteForIsvRequest) (_result *CreateNoteForIsvResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateNoteForIsvHeaders{}
	_result = &CreateNoteForIsvResponse{}
	_body, _err := client.CreateNoteForIsvWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// isv创建商机
//
// @param request - CreateOpportunityRequest
//
// @param headers - CreateOpportunityHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOpportunityResponse
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

// Summary:
//
// isv创建商机
//
// @param request - CreateOpportunityRequest
//
// @return CreateOpportunityResponse
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

// Summary:
//
// # QwenWoker询价接口
//
// @param request - PurchaseMixViewRequest
//
// @param headers - PurchaseMixViewHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PurchaseMixViewResponse
func (client *Client) PurchaseMixViewWithOptions(request *PurchaseMixViewRequest, headers *PurchaseMixViewHeaders, runtime *util.RuntimeOptions) (_result *PurchaseMixViewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelCode)) {
		body["channelCode"] = request.ChannelCode
	}

	if !tea.BoolValue(util.IsUnset(request.CombineActivityId)) {
		body["combineActivityId"] = request.CombineActivityId
	}

	if !tea.BoolValue(util.IsUnset(request.CreateOrder)) {
		body["createOrder"] = request.CreateOrder
	}

	if !tea.BoolValue(util.IsUnset(request.List)) {
		body["list"] = request.List
	}

	if !tea.BoolValue(util.IsUnset(request.Memo)) {
		body["memo"] = request.Memo
	}

	if !tea.BoolValue(util.IsUnset(request.MergeOrderName)) {
		body["mergeOrderName"] = request.MergeOrderName
	}

	if !tea.BoolValue(util.IsUnset(request.NeedModelTypeList)) {
		body["needModelTypeList"] = request.NeedModelTypeList
	}

	if !tea.BoolValue(util.IsUnset(request.ObjId)) {
		body["objId"] = request.ObjId
	}

	if !tea.BoolValue(util.IsUnset(request.ObjType)) {
		body["objType"] = request.ObjType
	}

	if !tea.BoolValue(util.IsUnset(request.OrderParamMap)) {
		body["orderParamMap"] = request.OrderParamMap
	}

	if !tea.BoolValue(util.IsUnset(request.OuterTradeCode)) {
		body["outerTradeCode"] = request.OuterTradeCode
	}

	if !tea.BoolValue(util.IsUnset(request.OuterTradeType)) {
		body["outerTradeType"] = request.OuterTradeType
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		body["planId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		body["uid"] = request.Uid
	}

	if !tea.BoolValue(util.IsUnset(request.UnPay)) {
		body["unPay"] = request.UnPay
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
		Action:      tea.String("PurchaseMixView"),
		Version:     tea.String("trade_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/trade/purchase/mixView"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PurchaseMixViewResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # QwenWoker询价接口
//
// @param request - PurchaseMixViewRequest
//
// @return PurchaseMixViewResponse
func (client *Client) PurchaseMixView(request *PurchaseMixViewRequest) (_result *PurchaseMixViewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PurchaseMixViewHeaders{}
	_result = &PurchaseMixViewResponse{}
	_body, _err := client.PurchaseMixViewWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询订单信息
//
// @param request - QueryTradeOrderRequest
//
// @param headers - QueryTradeOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTradeOrderResponse
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

// Summary:
//
// 查询订单信息
//
// @param request - QueryTradeOrderRequest
//
// @return QueryTradeOrderResponse
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
