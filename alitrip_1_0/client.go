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
