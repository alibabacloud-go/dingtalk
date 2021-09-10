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
