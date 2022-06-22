// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package hrm_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddHrmPreentryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddHrmPreentryHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryHeaders) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryHeaders) SetCommonHeaders(v map[string]*string) *AddHrmPreentryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddHrmPreentryHeaders) SetXAcsDingtalkAccessToken(v string) *AddHrmPreentryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddHrmPreentryRequest struct {
	AgentId      *int64                         `json:"agentId,omitempty" xml:"agentId,omitempty"`
	Groups       []*AddHrmPreentryRequestGroups `json:"groups,omitempty" xml:"groups,omitempty" type:"Repeated"`
	Mobile       *string                        `json:"mobile,omitempty" xml:"mobile,omitempty"`
	Name         *string                        `json:"name,omitempty" xml:"name,omitempty"`
	PreEntryTime *int64                         `json:"preEntryTime,omitempty" xml:"preEntryTime,omitempty"`
}

func (s AddHrmPreentryRequest) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryRequest) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryRequest) SetAgentId(v int64) *AddHrmPreentryRequest {
	s.AgentId = &v
	return s
}

func (s *AddHrmPreentryRequest) SetGroups(v []*AddHrmPreentryRequestGroups) *AddHrmPreentryRequest {
	s.Groups = v
	return s
}

func (s *AddHrmPreentryRequest) SetMobile(v string) *AddHrmPreentryRequest {
	s.Mobile = &v
	return s
}

func (s *AddHrmPreentryRequest) SetName(v string) *AddHrmPreentryRequest {
	s.Name = &v
	return s
}

func (s *AddHrmPreentryRequest) SetPreEntryTime(v int64) *AddHrmPreentryRequest {
	s.PreEntryTime = &v
	return s
}

type AddHrmPreentryRequestGroups struct {
	GroupId  *string                                `json:"groupId,omitempty" xml:"groupId,omitempty"`
	Sections []*AddHrmPreentryRequestGroupsSections `json:"sections,omitempty" xml:"sections,omitempty" type:"Repeated"`
}

func (s AddHrmPreentryRequestGroups) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryRequestGroups) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryRequestGroups) SetGroupId(v string) *AddHrmPreentryRequestGroups {
	s.GroupId = &v
	return s
}

func (s *AddHrmPreentryRequestGroups) SetSections(v []*AddHrmPreentryRequestGroupsSections) *AddHrmPreentryRequestGroups {
	s.Sections = v
	return s
}

type AddHrmPreentryRequestGroupsSections struct {
	EmpFieldVOList []*AddHrmPreentryRequestGroupsSectionsEmpFieldVOList `json:"empFieldVOList,omitempty" xml:"empFieldVOList,omitempty" type:"Repeated"`
	OldIndex       *int32                                               `json:"oldIndex,omitempty" xml:"oldIndex,omitempty"`
}

func (s AddHrmPreentryRequestGroupsSections) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryRequestGroupsSections) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryRequestGroupsSections) SetEmpFieldVOList(v []*AddHrmPreentryRequestGroupsSectionsEmpFieldVOList) *AddHrmPreentryRequestGroupsSections {
	s.EmpFieldVOList = v
	return s
}

func (s *AddHrmPreentryRequestGroupsSections) SetOldIndex(v int32) *AddHrmPreentryRequestGroupsSections {
	s.OldIndex = &v
	return s
}

type AddHrmPreentryRequestGroupsSectionsEmpFieldVOList struct {
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	Value     *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AddHrmPreentryRequestGroupsSectionsEmpFieldVOList) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryRequestGroupsSectionsEmpFieldVOList) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryRequestGroupsSectionsEmpFieldVOList) SetFieldCode(v string) *AddHrmPreentryRequestGroupsSectionsEmpFieldVOList {
	s.FieldCode = &v
	return s
}

func (s *AddHrmPreentryRequestGroupsSectionsEmpFieldVOList) SetValue(v string) *AddHrmPreentryRequestGroupsSectionsEmpFieldVOList {
	s.Value = &v
	return s
}

type AddHrmPreentryResponseBody struct {
	// result
	TmpUserId *string `json:"tmpUserId,omitempty" xml:"tmpUserId,omitempty"`
}

func (s AddHrmPreentryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryResponseBody) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryResponseBody) SetTmpUserId(v string) *AddHrmPreentryResponseBody {
	s.TmpUserId = &v
	return s
}

type AddHrmPreentryResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddHrmPreentryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddHrmPreentryResponse) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryResponse) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryResponse) SetHeaders(v map[string]*string) *AddHrmPreentryResponse {
	s.Headers = v
	return s
}

func (s *AddHrmPreentryResponse) SetBody(v *AddHrmPreentryResponseBody) *AddHrmPreentryResponse {
	s.Body = v
	return s
}

type ECertQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ECertQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s ECertQueryHeaders) GoString() string {
	return s.String()
}

func (s *ECertQueryHeaders) SetCommonHeaders(v map[string]*string) *ECertQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ECertQueryHeaders) SetXAcsDingtalkAccessToken(v string) *ECertQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ECertQueryRequest struct {
	// 用户ID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ECertQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s ECertQueryRequest) GoString() string {
	return s.String()
}

func (s *ECertQueryRequest) SetUserId(v string) *ECertQueryRequest {
	s.UserId = &v
	return s
}

type ECertQueryResponseBody struct {
	// 身份证号码
	CertNO *string `json:"certNO,omitempty" xml:"certNO,omitempty"`
	// 职务ID
	EmployJobId *string `json:"employJobId,omitempty" xml:"employJobId,omitempty"`
	// 职务名称
	EmployJobIdLabel *string `json:"employJobIdLabel,omitempty" xml:"employJobIdLabel,omitempty"`
	// 职位ID
	EmployPositionId *string `json:"employPositionId,omitempty" xml:"employPositionId,omitempty"`
	// 职位名称
	EmployPositionIdLabel *string `json:"employPositionIdLabel,omitempty" xml:"employPositionIdLabel,omitempty"`
	// 职级ID
	EmployPositionRankId *string `json:"employPositionRankId,omitempty" xml:"employPositionRankId,omitempty"`
	// 职级名称
	EmployPositionRankIdLabel *string `json:"employPositionRankIdLabel,omitempty" xml:"employPositionRankIdLabel,omitempty"`
	// 入职日期
	HiredDate *string `json:"hiredDate,omitempty" xml:"hiredDate,omitempty"`
	// 离职日期
	LastWorkDay *string `json:"lastWorkDay,omitempty" xml:"lastWorkDay,omitempty"`
	// 主部门ID
	MainDeptId *int64 `json:"mainDeptId,omitempty" xml:"mainDeptId,omitempty"`
	// 主部门
	MainDeptName *string `json:"mainDeptName,omitempty" xml:"mainDeptName,omitempty"`
	// 姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 身份证姓名
	RealName *string `json:"realName,omitempty" xml:"realName,omitempty"`
	// 被动离职原因
	TerminationReasonPassive []*string `json:"terminationReasonPassive,omitempty" xml:"terminationReasonPassive,omitempty" type:"Repeated"`
	// 主动离职原因
	TerminationReasonVoluntary []*string `json:"terminationReasonVoluntary,omitempty" xml:"terminationReasonVoluntary,omitempty" type:"Repeated"`
}

func (s ECertQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ECertQueryResponseBody) GoString() string {
	return s.String()
}

func (s *ECertQueryResponseBody) SetCertNO(v string) *ECertQueryResponseBody {
	s.CertNO = &v
	return s
}

func (s *ECertQueryResponseBody) SetEmployJobId(v string) *ECertQueryResponseBody {
	s.EmployJobId = &v
	return s
}

func (s *ECertQueryResponseBody) SetEmployJobIdLabel(v string) *ECertQueryResponseBody {
	s.EmployJobIdLabel = &v
	return s
}

func (s *ECertQueryResponseBody) SetEmployPositionId(v string) *ECertQueryResponseBody {
	s.EmployPositionId = &v
	return s
}

func (s *ECertQueryResponseBody) SetEmployPositionIdLabel(v string) *ECertQueryResponseBody {
	s.EmployPositionIdLabel = &v
	return s
}

func (s *ECertQueryResponseBody) SetEmployPositionRankId(v string) *ECertQueryResponseBody {
	s.EmployPositionRankId = &v
	return s
}

func (s *ECertQueryResponseBody) SetEmployPositionRankIdLabel(v string) *ECertQueryResponseBody {
	s.EmployPositionRankIdLabel = &v
	return s
}

func (s *ECertQueryResponseBody) SetHiredDate(v string) *ECertQueryResponseBody {
	s.HiredDate = &v
	return s
}

func (s *ECertQueryResponseBody) SetLastWorkDay(v string) *ECertQueryResponseBody {
	s.LastWorkDay = &v
	return s
}

func (s *ECertQueryResponseBody) SetMainDeptId(v int64) *ECertQueryResponseBody {
	s.MainDeptId = &v
	return s
}

func (s *ECertQueryResponseBody) SetMainDeptName(v string) *ECertQueryResponseBody {
	s.MainDeptName = &v
	return s
}

func (s *ECertQueryResponseBody) SetName(v string) *ECertQueryResponseBody {
	s.Name = &v
	return s
}

func (s *ECertQueryResponseBody) SetRealName(v string) *ECertQueryResponseBody {
	s.RealName = &v
	return s
}

func (s *ECertQueryResponseBody) SetTerminationReasonPassive(v []*string) *ECertQueryResponseBody {
	s.TerminationReasonPassive = v
	return s
}

func (s *ECertQueryResponseBody) SetTerminationReasonVoluntary(v []*string) *ECertQueryResponseBody {
	s.TerminationReasonVoluntary = v
	return s
}

type ECertQueryResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ECertQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ECertQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s ECertQueryResponse) GoString() string {
	return s.String()
}

func (s *ECertQueryResponse) SetHeaders(v map[string]*string) *ECertQueryResponse {
	s.Headers = v
	return s
}

func (s *ECertQueryResponse) SetBody(v *ECertQueryResponseBody) *ECertQueryResponse {
	s.Body = v
	return s
}

type HrmProcessTransferHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrmProcessTransferHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessTransferHeaders) GoString() string {
	return s.String()
}

func (s *HrmProcessTransferHeaders) SetCommonHeaders(v map[string]*string) *HrmProcessTransferHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrmProcessTransferHeaders) SetXAcsDingtalkAccessToken(v string) *HrmProcessTransferHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrmProcessTransferRequest struct {
	// 员工调岗后的部门id列表
	DeptIdsAfterTransfer []*int64 `json:"deptIdsAfterTransfer,omitempty" xml:"deptIdsAfterTransfer,omitempty" type:"Repeated"`
	// 员工调岗后的职务id
	JobIdAfterTransfer *string `json:"jobIdAfterTransfer,omitempty" xml:"jobIdAfterTransfer,omitempty"`
	// 员工调岗后的人事主部门id
	MainDeptIdAfterTransfer *int64 `json:"mainDeptIdAfterTransfer,omitempty" xml:"mainDeptIdAfterTransfer,omitempty"`
	// 操作人
	OperateUserId *string `json:"operateUserId,omitempty" xml:"operateUserId,omitempty"`
	// 员工调岗后的职位id，参数同时有职位名称以及id，以id为准
	PositionIdAfterTransfer *string `json:"positionIdAfterTransfer,omitempty" xml:"positionIdAfterTransfer,omitempty"`
	// 员工调岗后的职级名称，长度不超过64，参数同时有职级名称以及id，以id为准
	PositionLevelAfterTransfer *string `json:"positionLevelAfterTransfer,omitempty" xml:"positionLevelAfterTransfer,omitempty"`
	// 员工调岗后的职位名称，长度不超过124，参数同时有职位名称以及id，以id为准
	PositionNameAfterTransfer *string `json:"positionNameAfterTransfer,omitempty" xml:"positionNameAfterTransfer,omitempty"`
	// 员工调岗后的职级id，参数同时有职级名称以及id，以id为准
	RankIdAfterTransfer *string `json:"rankIdAfterTransfer,omitempty" xml:"rankIdAfterTransfer,omitempty"`
	// 被调岗员工userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s HrmProcessTransferRequest) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessTransferRequest) GoString() string {
	return s.String()
}

func (s *HrmProcessTransferRequest) SetDeptIdsAfterTransfer(v []*int64) *HrmProcessTransferRequest {
	s.DeptIdsAfterTransfer = v
	return s
}

func (s *HrmProcessTransferRequest) SetJobIdAfterTransfer(v string) *HrmProcessTransferRequest {
	s.JobIdAfterTransfer = &v
	return s
}

func (s *HrmProcessTransferRequest) SetMainDeptIdAfterTransfer(v int64) *HrmProcessTransferRequest {
	s.MainDeptIdAfterTransfer = &v
	return s
}

func (s *HrmProcessTransferRequest) SetOperateUserId(v string) *HrmProcessTransferRequest {
	s.OperateUserId = &v
	return s
}

func (s *HrmProcessTransferRequest) SetPositionIdAfterTransfer(v string) *HrmProcessTransferRequest {
	s.PositionIdAfterTransfer = &v
	return s
}

func (s *HrmProcessTransferRequest) SetPositionLevelAfterTransfer(v string) *HrmProcessTransferRequest {
	s.PositionLevelAfterTransfer = &v
	return s
}

func (s *HrmProcessTransferRequest) SetPositionNameAfterTransfer(v string) *HrmProcessTransferRequest {
	s.PositionNameAfterTransfer = &v
	return s
}

func (s *HrmProcessTransferRequest) SetRankIdAfterTransfer(v string) *HrmProcessTransferRequest {
	s.RankIdAfterTransfer = &v
	return s
}

func (s *HrmProcessTransferRequest) SetUserId(v string) *HrmProcessTransferRequest {
	s.UserId = &v
	return s
}

type HrmProcessTransferResponseBody struct {
	// 是否转岗成功
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s HrmProcessTransferResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessTransferResponseBody) GoString() string {
	return s.String()
}

func (s *HrmProcessTransferResponseBody) SetResult(v bool) *HrmProcessTransferResponseBody {
	s.Result = &v
	return s
}

type HrmProcessTransferResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *HrmProcessTransferResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HrmProcessTransferResponse) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessTransferResponse) GoString() string {
	return s.String()
}

func (s *HrmProcessTransferResponse) SetHeaders(v map[string]*string) *HrmProcessTransferResponse {
	s.Headers = v
	return s
}

func (s *HrmProcessTransferResponse) SetBody(v *HrmProcessTransferResponseBody) *HrmProcessTransferResponse {
	s.Body = v
	return s
}

type MasterDataQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MasterDataQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryHeaders) GoString() string {
	return s.String()
}

func (s *MasterDataQueryHeaders) SetCommonHeaders(v map[string]*string) *MasterDataQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MasterDataQueryHeaders) SetXAcsDingtalkAccessToken(v string) *MasterDataQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MasterDataQueryRequest struct {
	// 数据唯一键
	BizUK *string `json:"bizUK,omitempty" xml:"bizUK,omitempty"`
	// 分页查询每页数据条数
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 分页查询的游标
	NextToken *int32 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 当前操作人userId
	OptUserId *string `json:"optUserId,omitempty" xml:"optUserId,omitempty"`
	// 其他查询条件
	QueryParams []*MasterDataQueryRequestQueryParams `json:"queryParams,omitempty" xml:"queryParams,omitempty" type:"Repeated"`
	// 关联id列表，一般为userId
	RelationIds []*string `json:"relationIds,omitempty" xml:"relationIds,omitempty" type:"Repeated"`
	// 领域code 由钉钉分配
	ScopeCode *string `json:"scopeCode,omitempty" xml:"scopeCode,omitempty"`
	// 数据生产方的租户id，由钉钉分配
	TenantId *int64 `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 实体code
	ViewEntityCode *string `json:"viewEntityCode,omitempty" xml:"viewEntityCode,omitempty"`
}

func (s MasterDataQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryRequest) GoString() string {
	return s.String()
}

func (s *MasterDataQueryRequest) SetBizUK(v string) *MasterDataQueryRequest {
	s.BizUK = &v
	return s
}

func (s *MasterDataQueryRequest) SetMaxResults(v int32) *MasterDataQueryRequest {
	s.MaxResults = &v
	return s
}

func (s *MasterDataQueryRequest) SetNextToken(v int32) *MasterDataQueryRequest {
	s.NextToken = &v
	return s
}

func (s *MasterDataQueryRequest) SetOptUserId(v string) *MasterDataQueryRequest {
	s.OptUserId = &v
	return s
}

func (s *MasterDataQueryRequest) SetQueryParams(v []*MasterDataQueryRequestQueryParams) *MasterDataQueryRequest {
	s.QueryParams = v
	return s
}

func (s *MasterDataQueryRequest) SetRelationIds(v []*string) *MasterDataQueryRequest {
	s.RelationIds = v
	return s
}

func (s *MasterDataQueryRequest) SetScopeCode(v string) *MasterDataQueryRequest {
	s.ScopeCode = &v
	return s
}

func (s *MasterDataQueryRequest) SetTenantId(v int64) *MasterDataQueryRequest {
	s.TenantId = &v
	return s
}

func (s *MasterDataQueryRequest) SetViewEntityCode(v string) *MasterDataQueryRequest {
	s.ViewEntityCode = &v
	return s
}

type MasterDataQueryRequestQueryParams struct {
	// 筛选条件
	ConditionList []*MasterDataQueryRequestQueryParamsConditionList `json:"conditionList,omitempty" xml:"conditionList,omitempty" type:"Repeated"`
	// 需要筛选的字段
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	// 筛选条件连接类型
	JoinType *string `json:"joinType,omitempty" xml:"joinType,omitempty"`
}

func (s MasterDataQueryRequestQueryParams) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryRequestQueryParams) GoString() string {
	return s.String()
}

func (s *MasterDataQueryRequestQueryParams) SetConditionList(v []*MasterDataQueryRequestQueryParamsConditionList) *MasterDataQueryRequestQueryParams {
	s.ConditionList = v
	return s
}

func (s *MasterDataQueryRequestQueryParams) SetFieldCode(v string) *MasterDataQueryRequestQueryParams {
	s.FieldCode = &v
	return s
}

func (s *MasterDataQueryRequestQueryParams) SetJoinType(v string) *MasterDataQueryRequestQueryParams {
	s.JoinType = &v
	return s
}

type MasterDataQueryRequestQueryParamsConditionList struct {
	// 字段关系符
	Operate *string `json:"operate,omitempty" xml:"operate,omitempty"`
	// 操作值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s MasterDataQueryRequestQueryParamsConditionList) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryRequestQueryParamsConditionList) GoString() string {
	return s.String()
}

func (s *MasterDataQueryRequestQueryParamsConditionList) SetOperate(v string) *MasterDataQueryRequestQueryParamsConditionList {
	s.Operate = &v
	return s
}

func (s *MasterDataQueryRequestQueryParamsConditionList) SetValue(v string) *MasterDataQueryRequestQueryParamsConditionList {
	s.Value = &v
	return s
}

type MasterDataQueryResponseBody struct {
	// 是否还有更多
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 分页游标
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 结果
	Result []*MasterDataQueryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 总条目数
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s MasterDataQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryResponseBody) GoString() string {
	return s.String()
}

func (s *MasterDataQueryResponseBody) SetHasMore(v bool) *MasterDataQueryResponseBody {
	s.HasMore = &v
	return s
}

func (s *MasterDataQueryResponseBody) SetNextToken(v int64) *MasterDataQueryResponseBody {
	s.NextToken = &v
	return s
}

func (s *MasterDataQueryResponseBody) SetResult(v []*MasterDataQueryResponseBodyResult) *MasterDataQueryResponseBody {
	s.Result = v
	return s
}

func (s *MasterDataQueryResponseBody) SetSuccess(v bool) *MasterDataQueryResponseBody {
	s.Success = &v
	return s
}

func (s *MasterDataQueryResponseBody) SetTotal(v int64) *MasterDataQueryResponseBody {
	s.Total = &v
	return s
}

type MasterDataQueryResponseBodyResult struct {
	// 唯一id
	OuterId *string `json:"outerId,omitempty" xml:"outerId,omitempty"`
	// 关联id列表，一般为userId
	RelationId *string `json:"relationId,omitempty" xml:"relationId,omitempty"`
	// 领域
	ScopeCode *string `json:"scopeCode,omitempty" xml:"scopeCode,omitempty"`
	// 编码
	ViewEntityCode *string `json:"viewEntityCode,omitempty" xml:"viewEntityCode,omitempty"`
	// 字段列表
	ViewEntityFieldVOList []*MasterDataQueryResponseBodyResultViewEntityFieldVOList `json:"viewEntityFieldVOList,omitempty" xml:"viewEntityFieldVOList,omitempty" type:"Repeated"`
}

func (s MasterDataQueryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *MasterDataQueryResponseBodyResult) SetOuterId(v string) *MasterDataQueryResponseBodyResult {
	s.OuterId = &v
	return s
}

func (s *MasterDataQueryResponseBodyResult) SetRelationId(v string) *MasterDataQueryResponseBodyResult {
	s.RelationId = &v
	return s
}

func (s *MasterDataQueryResponseBodyResult) SetScopeCode(v string) *MasterDataQueryResponseBodyResult {
	s.ScopeCode = &v
	return s
}

func (s *MasterDataQueryResponseBodyResult) SetViewEntityCode(v string) *MasterDataQueryResponseBodyResult {
	s.ViewEntityCode = &v
	return s
}

func (s *MasterDataQueryResponseBodyResult) SetViewEntityFieldVOList(v []*MasterDataQueryResponseBodyResultViewEntityFieldVOList) *MasterDataQueryResponseBodyResult {
	s.ViewEntityFieldVOList = v
	return s
}

type MasterDataQueryResponseBodyResultViewEntityFieldVOList struct {
	// 字段code
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	// 字段值
	FieldDataVO *MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO `json:"fieldDataVO,omitempty" xml:"fieldDataVO,omitempty" type:"Struct"`
	// 字段名称
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// 字段类型
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
}

func (s MasterDataQueryResponseBodyResultViewEntityFieldVOList) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryResponseBodyResultViewEntityFieldVOList) GoString() string {
	return s.String()
}

func (s *MasterDataQueryResponseBodyResultViewEntityFieldVOList) SetFieldCode(v string) *MasterDataQueryResponseBodyResultViewEntityFieldVOList {
	s.FieldCode = &v
	return s
}

func (s *MasterDataQueryResponseBodyResultViewEntityFieldVOList) SetFieldDataVO(v *MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO) *MasterDataQueryResponseBodyResultViewEntityFieldVOList {
	s.FieldDataVO = v
	return s
}

func (s *MasterDataQueryResponseBodyResultViewEntityFieldVOList) SetFieldName(v string) *MasterDataQueryResponseBodyResultViewEntityFieldVOList {
	s.FieldName = &v
	return s
}

func (s *MasterDataQueryResponseBodyResultViewEntityFieldVOList) SetFieldType(v string) *MasterDataQueryResponseBodyResultViewEntityFieldVOList {
	s.FieldType = &v
	return s
}

type MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO struct {
	// 字段值的key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 字段值的文本
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO) GoString() string {
	return s.String()
}

func (s *MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO) SetKey(v string) *MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO {
	s.Key = &v
	return s
}

func (s *MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO) SetValue(v string) *MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO {
	s.Value = &v
	return s
}

type MasterDataQueryResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MasterDataQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MasterDataQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryResponse) GoString() string {
	return s.String()
}

func (s *MasterDataQueryResponse) SetHeaders(v map[string]*string) *MasterDataQueryResponse {
	s.Headers = v
	return s
}

func (s *MasterDataQueryResponse) SetBody(v *MasterDataQueryResponseBody) *MasterDataQueryResponse {
	s.Body = v
	return s
}

type MasterDataSaveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MasterDataSaveHeaders) String() string {
	return tea.Prettify(s)
}

func (s MasterDataSaveHeaders) GoString() string {
	return s.String()
}

func (s *MasterDataSaveHeaders) SetCommonHeaders(v map[string]*string) *MasterDataSaveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MasterDataSaveHeaders) SetXAcsDingtalkAccessToken(v string) *MasterDataSaveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MasterDataSaveRequest struct {
	// 主数据
	Body []*MasterDataSaveRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// 租户id
	TenantId *int64 `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s MasterDataSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s MasterDataSaveRequest) GoString() string {
	return s.String()
}

func (s *MasterDataSaveRequest) SetBody(v []*MasterDataSaveRequestBody) *MasterDataSaveRequest {
	s.Body = v
	return s
}

func (s *MasterDataSaveRequest) SetTenantId(v int64) *MasterDataSaveRequest {
	s.TenantId = &v
	return s
}

type MasterDataSaveRequestBody struct {
	// 数据变更时间戳，用以保证更新操作的顺序性
	BizTime *int64 `json:"bizTime,omitempty" xml:"bizTime,omitempty"`
	// 数据流水唯一标识，如流水号，用以唯一确认一条写入数据
	BizUk *string `json:"bizUk,omitempty" xml:"bizUk,omitempty"`
	// 业务域下的细分领域实体
	EntityCode *string `json:"entityCode,omitempty" xml:"entityCode,omitempty"`
	// 数据字段列表
	FieldList []*MasterDataSaveRequestBodyFieldList `json:"fieldList,omitempty" xml:"fieldList,omitempty" type:"Repeated"`
	// 业务域描述，系统分配
	Scope *MasterDataSaveRequestBodyScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
	// 员工id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s MasterDataSaveRequestBody) String() string {
	return tea.Prettify(s)
}

func (s MasterDataSaveRequestBody) GoString() string {
	return s.String()
}

func (s *MasterDataSaveRequestBody) SetBizTime(v int64) *MasterDataSaveRequestBody {
	s.BizTime = &v
	return s
}

func (s *MasterDataSaveRequestBody) SetBizUk(v string) *MasterDataSaveRequestBody {
	s.BizUk = &v
	return s
}

func (s *MasterDataSaveRequestBody) SetEntityCode(v string) *MasterDataSaveRequestBody {
	s.EntityCode = &v
	return s
}

func (s *MasterDataSaveRequestBody) SetFieldList(v []*MasterDataSaveRequestBodyFieldList) *MasterDataSaveRequestBody {
	s.FieldList = v
	return s
}

func (s *MasterDataSaveRequestBody) SetScope(v *MasterDataSaveRequestBodyScope) *MasterDataSaveRequestBody {
	s.Scope = v
	return s
}

func (s *MasterDataSaveRequestBody) SetUserId(v string) *MasterDataSaveRequestBody {
	s.UserId = &v
	return s
}

type MasterDataSaveRequestBodyFieldList struct {
	// 字段名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 字段string值
	ValueStr *string `json:"valueStr,omitempty" xml:"valueStr,omitempty"`
}

func (s MasterDataSaveRequestBodyFieldList) String() string {
	return tea.Prettify(s)
}

func (s MasterDataSaveRequestBodyFieldList) GoString() string {
	return s.String()
}

func (s *MasterDataSaveRequestBodyFieldList) SetName(v string) *MasterDataSaveRequestBodyFieldList {
	s.Name = &v
	return s
}

func (s *MasterDataSaveRequestBodyFieldList) SetValueStr(v string) *MasterDataSaveRequestBodyFieldList {
	s.ValueStr = &v
	return s
}

type MasterDataSaveRequestBodyScope struct {
	// 业务域code，如PERFORMANCE，系统分配
	ScopeCode *string `json:"scopeCode,omitempty" xml:"scopeCode,omitempty"`
	// 业务域版本，接入时系统分配，默认传1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s MasterDataSaveRequestBodyScope) String() string {
	return tea.Prettify(s)
}

func (s MasterDataSaveRequestBodyScope) GoString() string {
	return s.String()
}

func (s *MasterDataSaveRequestBodyScope) SetScopeCode(v string) *MasterDataSaveRequestBodyScope {
	s.ScopeCode = &v
	return s
}

func (s *MasterDataSaveRequestBodyScope) SetVersion(v int32) *MasterDataSaveRequestBodyScope {
	s.Version = &v
	return s
}

type MasterDataSaveResponseBody struct {
	// 是否全部保存成功
	AllSuccess *bool `json:"allSuccess,omitempty" xml:"allSuccess,omitempty"`
	// 保存失败的结果，全部保存成功时为空
	FailResult []*MasterDataSaveResponseBodyFailResult `json:"failResult,omitempty" xml:"failResult,omitempty" type:"Repeated"`
}

func (s MasterDataSaveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MasterDataSaveResponseBody) GoString() string {
	return s.String()
}

func (s *MasterDataSaveResponseBody) SetAllSuccess(v bool) *MasterDataSaveResponseBody {
	s.AllSuccess = &v
	return s
}

func (s *MasterDataSaveResponseBody) SetFailResult(v []*MasterDataSaveResponseBodyFailResult) *MasterDataSaveResponseBody {
	s.FailResult = v
	return s
}

type MasterDataSaveResponseBodyFailResult struct {
	// 业务流水唯一标识，和入参一致
	BizUk *string `json:"bizUk,omitempty" xml:"bizUk,omitempty"`
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s MasterDataSaveResponseBodyFailResult) String() string {
	return tea.Prettify(s)
}

func (s MasterDataSaveResponseBodyFailResult) GoString() string {
	return s.String()
}

func (s *MasterDataSaveResponseBodyFailResult) SetBizUk(v string) *MasterDataSaveResponseBodyFailResult {
	s.BizUk = &v
	return s
}

func (s *MasterDataSaveResponseBodyFailResult) SetErrorCode(v string) *MasterDataSaveResponseBodyFailResult {
	s.ErrorCode = &v
	return s
}

func (s *MasterDataSaveResponseBodyFailResult) SetErrorMsg(v string) *MasterDataSaveResponseBodyFailResult {
	s.ErrorMsg = &v
	return s
}

func (s *MasterDataSaveResponseBodyFailResult) SetSuccess(v bool) *MasterDataSaveResponseBodyFailResult {
	s.Success = &v
	return s
}

type MasterDataSaveResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MasterDataSaveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MasterDataSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s MasterDataSaveResponse) GoString() string {
	return s.String()
}

func (s *MasterDataSaveResponse) SetHeaders(v map[string]*string) *MasterDataSaveResponse {
	s.Headers = v
	return s
}

func (s *MasterDataSaveResponse) SetBody(v *MasterDataSaveResponseBody) *MasterDataSaveResponse {
	s.Body = v
	return s
}

type MasterDataTenantQueyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MasterDataTenantQueyHeaders) String() string {
	return tea.Prettify(s)
}

func (s MasterDataTenantQueyHeaders) GoString() string {
	return s.String()
}

func (s *MasterDataTenantQueyHeaders) SetCommonHeaders(v map[string]*string) *MasterDataTenantQueyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MasterDataTenantQueyHeaders) SetXAcsDingtalkAccessToken(v string) *MasterDataTenantQueyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MasterDataTenantQueyRequest struct {
	// 实体 code
	EntityCode *string `json:"entityCode,omitempty" xml:"entityCode,omitempty"`
	// isv的业务领域
	ScopeCode *string `json:"scopeCode,omitempty" xml:"scopeCode,omitempty"`
}

func (s MasterDataTenantQueyRequest) String() string {
	return tea.Prettify(s)
}

func (s MasterDataTenantQueyRequest) GoString() string {
	return s.String()
}

func (s *MasterDataTenantQueyRequest) SetEntityCode(v string) *MasterDataTenantQueyRequest {
	s.EntityCode = &v
	return s
}

func (s *MasterDataTenantQueyRequest) SetScopeCode(v string) *MasterDataTenantQueyRequest {
	s.ScopeCode = &v
	return s
}

type MasterDataTenantQueyResponseBody struct {
	Result []*MasterDataTenantQueyResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s MasterDataTenantQueyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MasterDataTenantQueyResponseBody) GoString() string {
	return s.String()
}

func (s *MasterDataTenantQueyResponseBody) SetResult(v []*MasterDataTenantQueyResponseBodyResult) *MasterDataTenantQueyResponseBody {
	s.Result = v
	return s
}

type MasterDataTenantQueyResponseBodyResult struct {
	// 该租户是否已向主数据同步数据
	HasData *bool `json:"hasData,omitempty" xml:"hasData,omitempty"`
	// 该租户是否有向主数据写数据的权限
	IntegrateDataAuth *bool `json:"integrateDataAuth,omitempty" xml:"integrateDataAuth,omitempty"`
	// 租户名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 调用方是否有读该租户数据的权限
	ReadAuth *bool `json:"readAuth,omitempty" xml:"readAuth,omitempty"`
	// 租户 id
	TenantId *int64 `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 租户类型
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s MasterDataTenantQueyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s MasterDataTenantQueyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *MasterDataTenantQueyResponseBodyResult) SetHasData(v bool) *MasterDataTenantQueyResponseBodyResult {
	s.HasData = &v
	return s
}

func (s *MasterDataTenantQueyResponseBodyResult) SetIntegrateDataAuth(v bool) *MasterDataTenantQueyResponseBodyResult {
	s.IntegrateDataAuth = &v
	return s
}

func (s *MasterDataTenantQueyResponseBodyResult) SetName(v string) *MasterDataTenantQueyResponseBodyResult {
	s.Name = &v
	return s
}

func (s *MasterDataTenantQueyResponseBodyResult) SetReadAuth(v bool) *MasterDataTenantQueyResponseBodyResult {
	s.ReadAuth = &v
	return s
}

func (s *MasterDataTenantQueyResponseBodyResult) SetTenantId(v int64) *MasterDataTenantQueyResponseBodyResult {
	s.TenantId = &v
	return s
}

func (s *MasterDataTenantQueyResponseBodyResult) SetType(v int32) *MasterDataTenantQueyResponseBodyResult {
	s.Type = &v
	return s
}

type MasterDataTenantQueyResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MasterDataTenantQueyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MasterDataTenantQueyResponse) String() string {
	return tea.Prettify(s)
}

func (s MasterDataTenantQueyResponse) GoString() string {
	return s.String()
}

func (s *MasterDataTenantQueyResponse) SetHeaders(v map[string]*string) *MasterDataTenantQueyResponse {
	s.Headers = v
	return s
}

func (s *MasterDataTenantQueyResponse) SetBody(v *MasterDataTenantQueyResponseBody) *MasterDataTenantQueyResponse {
	s.Body = v
	return s
}

type QueryCustomEntryProcessesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCustomEntryProcessesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomEntryProcessesHeaders) GoString() string {
	return s.String()
}

func (s *QueryCustomEntryProcessesHeaders) SetCommonHeaders(v map[string]*string) *QueryCustomEntryProcessesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCustomEntryProcessesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCustomEntryProcessesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCustomEntryProcessesRequest struct {
	// 最大值
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 偏移量
	NextToken *int32 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 操作人id
	OperateUserId *string `json:"operateUserId,omitempty" xml:"operateUserId,omitempty"`
}

func (s QueryCustomEntryProcessesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomEntryProcessesRequest) GoString() string {
	return s.String()
}

func (s *QueryCustomEntryProcessesRequest) SetMaxResults(v int32) *QueryCustomEntryProcessesRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryCustomEntryProcessesRequest) SetNextToken(v int32) *QueryCustomEntryProcessesRequest {
	s.NextToken = &v
	return s
}

func (s *QueryCustomEntryProcessesRequest) SetOperateUserId(v string) *QueryCustomEntryProcessesRequest {
	s.OperateUserId = &v
	return s
}

type QueryCustomEntryProcessesResponseBody struct {
	// 是否有更多
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 表单信息列表
	List []*QueryCustomEntryProcessesResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// 下次获取数据的起始游标
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryCustomEntryProcessesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomEntryProcessesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCustomEntryProcessesResponseBody) SetHasMore(v bool) *QueryCustomEntryProcessesResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryCustomEntryProcessesResponseBody) SetList(v []*QueryCustomEntryProcessesResponseBodyList) *QueryCustomEntryProcessesResponseBody {
	s.List = v
	return s
}

func (s *QueryCustomEntryProcessesResponseBody) SetNextToken(v int64) *QueryCustomEntryProcessesResponseBody {
	s.NextToken = &v
	return s
}

type QueryCustomEntryProcessesResponseBodyList struct {
	FormDesc *string `json:"formDesc,omitempty" xml:"formDesc,omitempty"`
	FormId   *string `json:"formId,omitempty" xml:"formId,omitempty"`
	FormName *string `json:"formName,omitempty" xml:"formName,omitempty"`
	ShortUrl *string `json:"shortUrl,omitempty" xml:"shortUrl,omitempty"`
}

func (s QueryCustomEntryProcessesResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomEntryProcessesResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryCustomEntryProcessesResponseBodyList) SetFormDesc(v string) *QueryCustomEntryProcessesResponseBodyList {
	s.FormDesc = &v
	return s
}

func (s *QueryCustomEntryProcessesResponseBodyList) SetFormId(v string) *QueryCustomEntryProcessesResponseBodyList {
	s.FormId = &v
	return s
}

func (s *QueryCustomEntryProcessesResponseBodyList) SetFormName(v string) *QueryCustomEntryProcessesResponseBodyList {
	s.FormName = &v
	return s
}

func (s *QueryCustomEntryProcessesResponseBodyList) SetShortUrl(v string) *QueryCustomEntryProcessesResponseBodyList {
	s.ShortUrl = &v
	return s
}

type QueryCustomEntryProcessesResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCustomEntryProcessesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCustomEntryProcessesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomEntryProcessesResponse) GoString() string {
	return s.String()
}

func (s *QueryCustomEntryProcessesResponse) SetHeaders(v map[string]*string) *QueryCustomEntryProcessesResponse {
	s.Headers = v
	return s
}

func (s *QueryCustomEntryProcessesResponse) SetBody(v *QueryCustomEntryProcessesResponseBody) *QueryCustomEntryProcessesResponse {
	s.Body = v
	return s
}

type QueryHrmEmployeeDismissionInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryHrmEmployeeDismissionInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryHrmEmployeeDismissionInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryHrmEmployeeDismissionInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryHrmEmployeeDismissionInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryHrmEmployeeDismissionInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryHrmEmployeeDismissionInfoRequest struct {
	// 员工 ids
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s QueryHrmEmployeeDismissionInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHrmEmployeeDismissionInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryHrmEmployeeDismissionInfoRequest) SetUserIdList(v []*string) *QueryHrmEmployeeDismissionInfoRequest {
	s.UserIdList = v
	return s
}

type QueryHrmEmployeeDismissionInfoShrinkRequest struct {
	// 员工 ids
	UserIdListShrink *string `json:"userIdList,omitempty" xml:"userIdList,omitempty"`
}

func (s QueryHrmEmployeeDismissionInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHrmEmployeeDismissionInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryHrmEmployeeDismissionInfoShrinkRequest) SetUserIdListShrink(v string) *QueryHrmEmployeeDismissionInfoShrinkRequest {
	s.UserIdListShrink = &v
	return s
}

type QueryHrmEmployeeDismissionInfoResponseBody struct {
	// 名称列表
	Result []*QueryHrmEmployeeDismissionInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryHrmEmployeeDismissionInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHrmEmployeeDismissionInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHrmEmployeeDismissionInfoResponseBody) SetResult(v []*QueryHrmEmployeeDismissionInfoResponseBodyResult) *QueryHrmEmployeeDismissionInfoResponseBody {
	s.Result = v
	return s
}

type QueryHrmEmployeeDismissionInfoResponseBodyResult struct {
	// 离职部门列表
	DeptList []*QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList `json:"deptList,omitempty" xml:"deptList,omitempty" type:"Repeated"`
	// 离职交接人
	HandoverUserId *string `json:"handoverUserId,omitempty" xml:"handoverUserId,omitempty"`
	// 最后工作日
	LastWorkDay *int64 `json:"lastWorkDay,omitempty" xml:"lastWorkDay,omitempty"`
	// 离职前主部门id
	MainDeptId *int64 `json:"mainDeptId,omitempty" xml:"mainDeptId,omitempty"`
	// 离职前主部门名称
	MainDeptName *string `json:"mainDeptName,omitempty" xml:"mainDeptName,omitempty"`
	// 离职原因-被动
	PassiveReason []*string `json:"passiveReason,omitempty" xml:"passiveReason,omitempty" type:"Repeated"`
	// 离职前工作状态：1，待入职；2，试用期；3，正式
	PreStatus *int32 `json:"preStatus,omitempty" xml:"preStatus,omitempty"`
	// 离职原因备注
	ReasonMemo *string `json:"reasonMemo,omitempty" xml:"reasonMemo,omitempty"`
	// 离职状态：1，待离职；2，已离职
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// 员工id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 离职原因-主动
	VoluntaryReason []*string `json:"voluntaryReason,omitempty" xml:"voluntaryReason,omitempty" type:"Repeated"`
}

func (s QueryHrmEmployeeDismissionInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryHrmEmployeeDismissionInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetDeptList(v []*QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.DeptList = v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetHandoverUserId(v string) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.HandoverUserId = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetLastWorkDay(v int64) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.LastWorkDay = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetMainDeptId(v int64) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.MainDeptId = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetMainDeptName(v string) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.MainDeptName = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetPassiveReason(v []*string) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.PassiveReason = v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetPreStatus(v int32) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.PreStatus = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetReasonMemo(v string) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.ReasonMemo = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetStatus(v int32) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.Status = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetUserId(v string) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetVoluntaryReason(v []*string) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.VoluntaryReason = v
	return s
}

type QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList struct {
	// 部门id
	DeptId *int64 `json:"dept_id,omitempty" xml:"dept_id,omitempty"`
	// 部门路径
	DeptPath *string `json:"dept_path,omitempty" xml:"dept_path,omitempty"`
}

func (s QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList) String() string {
	return tea.Prettify(s)
}

func (s QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList) GoString() string {
	return s.String()
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList) SetDeptId(v int64) *QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList {
	s.DeptId = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList) SetDeptPath(v string) *QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList {
	s.DeptPath = &v
	return s
}

type QueryHrmEmployeeDismissionInfoResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryHrmEmployeeDismissionInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryHrmEmployeeDismissionInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHrmEmployeeDismissionInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryHrmEmployeeDismissionInfoResponse) SetHeaders(v map[string]*string) *QueryHrmEmployeeDismissionInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponse) SetBody(v *QueryHrmEmployeeDismissionInfoResponseBody) *QueryHrmEmployeeDismissionInfoResponse {
	s.Body = v
	return s
}

type QueryJobRanksHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryJobRanksHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryJobRanksHeaders) GoString() string {
	return s.String()
}

func (s *QueryJobRanksHeaders) SetCommonHeaders(v map[string]*string) *QueryJobRanksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryJobRanksHeaders) SetXAcsDingtalkAccessToken(v string) *QueryJobRanksHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryJobRanksRequest struct {
	// 本次读取的最大数据记录数量
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 标记当前开始读取的位置
	NextToken *int32 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 职级序列
	RankCategoryId *string `json:"rankCategoryId,omitempty" xml:"rankCategoryId,omitempty"`
	// 职级编码
	RankCode *string `json:"rankCode,omitempty" xml:"rankCode,omitempty"`
	// 职级名称
	RankName *string `json:"rankName,omitempty" xml:"rankName,omitempty"`
}

func (s QueryJobRanksRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryJobRanksRequest) GoString() string {
	return s.String()
}

func (s *QueryJobRanksRequest) SetMaxResults(v int32) *QueryJobRanksRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryJobRanksRequest) SetNextToken(v int32) *QueryJobRanksRequest {
	s.NextToken = &v
	return s
}

func (s *QueryJobRanksRequest) SetRankCategoryId(v string) *QueryJobRanksRequest {
	s.RankCategoryId = &v
	return s
}

func (s *QueryJobRanksRequest) SetRankCode(v string) *QueryJobRanksRequest {
	s.RankCode = &v
	return s
}

func (s *QueryJobRanksRequest) SetRankName(v string) *QueryJobRanksRequest {
	s.RankName = &v
	return s
}

type QueryJobRanksResponseBody struct {
	// 是否有更多数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 职级列表
	List []*QueryJobRanksResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryJobRanksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryJobRanksResponseBody) GoString() string {
	return s.String()
}

func (s *QueryJobRanksResponseBody) SetHasMore(v bool) *QueryJobRanksResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryJobRanksResponseBody) SetList(v []*QueryJobRanksResponseBodyList) *QueryJobRanksResponseBody {
	s.List = v
	return s
}

func (s *QueryJobRanksResponseBody) SetNextToken(v int64) *QueryJobRanksResponseBody {
	s.NextToken = &v
	return s
}

type QueryJobRanksResponseBodyList struct {
	// 最大等级
	MaxJobGrade *int32 `json:"maxJobGrade,omitempty" xml:"maxJobGrade,omitempty"`
	// 最小等级
	MinJobGrade *int32 `json:"minJobGrade,omitempty" xml:"minJobGrade,omitempty"`
	// 职级序列ID
	RankCategoryId *string `json:"rankCategoryId,omitempty" xml:"rankCategoryId,omitempty"`
	// 职级编码
	RankCode *string `json:"rankCode,omitempty" xml:"rankCode,omitempty"`
	// 职级描述
	RankDescription *string `json:"rankDescription,omitempty" xml:"rankDescription,omitempty"`
	// 职级ID
	RankId *string `json:"rankId,omitempty" xml:"rankId,omitempty"`
	// 职级名称
	RankName *string `json:"rankName,omitempty" xml:"rankName,omitempty"`
}

func (s QueryJobRanksResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryJobRanksResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryJobRanksResponseBodyList) SetMaxJobGrade(v int32) *QueryJobRanksResponseBodyList {
	s.MaxJobGrade = &v
	return s
}

func (s *QueryJobRanksResponseBodyList) SetMinJobGrade(v int32) *QueryJobRanksResponseBodyList {
	s.MinJobGrade = &v
	return s
}

func (s *QueryJobRanksResponseBodyList) SetRankCategoryId(v string) *QueryJobRanksResponseBodyList {
	s.RankCategoryId = &v
	return s
}

func (s *QueryJobRanksResponseBodyList) SetRankCode(v string) *QueryJobRanksResponseBodyList {
	s.RankCode = &v
	return s
}

func (s *QueryJobRanksResponseBodyList) SetRankDescription(v string) *QueryJobRanksResponseBodyList {
	s.RankDescription = &v
	return s
}

func (s *QueryJobRanksResponseBodyList) SetRankId(v string) *QueryJobRanksResponseBodyList {
	s.RankId = &v
	return s
}

func (s *QueryJobRanksResponseBodyList) SetRankName(v string) *QueryJobRanksResponseBodyList {
	s.RankName = &v
	return s
}

type QueryJobRanksResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryJobRanksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryJobRanksResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryJobRanksResponse) GoString() string {
	return s.String()
}

func (s *QueryJobRanksResponse) SetHeaders(v map[string]*string) *QueryJobRanksResponse {
	s.Headers = v
	return s
}

func (s *QueryJobRanksResponse) SetBody(v *QueryJobRanksResponseBody) *QueryJobRanksResponse {
	s.Body = v
	return s
}

type QueryJobsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryJobsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryJobsHeaders) GoString() string {
	return s.String()
}

func (s *QueryJobsHeaders) SetCommonHeaders(v map[string]*string) *QueryJobsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryJobsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryJobsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryJobsRequest struct {
	// 职务名称
	JobName *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
	// 最大值
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 偏移量
	NextToken *int32 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryJobsRequest) GoString() string {
	return s.String()
}

func (s *QueryJobsRequest) SetJobName(v string) *QueryJobsRequest {
	s.JobName = &v
	return s
}

func (s *QueryJobsRequest) SetMaxResults(v int32) *QueryJobsRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryJobsRequest) SetNextToken(v int32) *QueryJobsRequest {
	s.NextToken = &v
	return s
}

type QueryJobsResponseBody struct {
	// 是否有更多数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 职务列表
	List []*QueryJobsResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// 下次获取数据的起始游标
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryJobsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryJobsResponseBody) SetHasMore(v bool) *QueryJobsResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryJobsResponseBody) SetList(v []*QueryJobsResponseBodyList) *QueryJobsResponseBody {
	s.List = v
	return s
}

func (s *QueryJobsResponseBody) SetNextToken(v int64) *QueryJobsResponseBody {
	s.NextToken = &v
	return s
}

type QueryJobsResponseBodyList struct {
	// 职务描述
	JobDescription *string `json:"jobDescription,omitempty" xml:"jobDescription,omitempty"`
	// 职务ID
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// 职务名称
	JobName *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
}

func (s QueryJobsResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryJobsResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryJobsResponseBodyList) SetJobDescription(v string) *QueryJobsResponseBodyList {
	s.JobDescription = &v
	return s
}

func (s *QueryJobsResponseBodyList) SetJobId(v string) *QueryJobsResponseBodyList {
	s.JobId = &v
	return s
}

func (s *QueryJobsResponseBodyList) SetJobName(v string) *QueryJobsResponseBodyList {
	s.JobName = &v
	return s
}

type QueryJobsResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryJobsResponse) GoString() string {
	return s.String()
}

func (s *QueryJobsResponse) SetHeaders(v map[string]*string) *QueryJobsResponse {
	s.Headers = v
	return s
}

func (s *QueryJobsResponse) SetBody(v *QueryJobsResponseBody) *QueryJobsResponse {
	s.Body = v
	return s
}

type QueryPositionsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryPositionsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryPositionsHeaders) GoString() string {
	return s.String()
}

func (s *QueryPositionsHeaders) SetCommonHeaders(v map[string]*string) *QueryPositionsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryPositionsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryPositionsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryPositionsRequest struct {
	// 部门id
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 职位类别列表
	InCategoryIds []*string `json:"inCategoryIds,omitempty" xml:"inCategoryIds,omitempty" type:"Repeated"`
	// 职位id列表
	InPositionIds []*string `json:"inPositionIds,omitempty" xml:"inPositionIds,omitempty" type:"Repeated"`
	// 职位名称
	PositionName *string `json:"positionName,omitempty" xml:"positionName,omitempty"`
	// 一次查询获取记录数
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 偏移量
	NextToken *int32 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryPositionsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPositionsRequest) GoString() string {
	return s.String()
}

func (s *QueryPositionsRequest) SetDeptId(v int64) *QueryPositionsRequest {
	s.DeptId = &v
	return s
}

func (s *QueryPositionsRequest) SetInCategoryIds(v []*string) *QueryPositionsRequest {
	s.InCategoryIds = v
	return s
}

func (s *QueryPositionsRequest) SetInPositionIds(v []*string) *QueryPositionsRequest {
	s.InPositionIds = v
	return s
}

func (s *QueryPositionsRequest) SetPositionName(v string) *QueryPositionsRequest {
	s.PositionName = &v
	return s
}

func (s *QueryPositionsRequest) SetMaxResults(v int32) *QueryPositionsRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryPositionsRequest) SetNextToken(v int32) *QueryPositionsRequest {
	s.NextToken = &v
	return s
}

type QueryPositionsResponseBody struct {
	// 是否有更多数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 职位列表
	List []*QueryPositionsResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryPositionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPositionsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPositionsResponseBody) SetHasMore(v bool) *QueryPositionsResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryPositionsResponseBody) SetList(v []*QueryPositionsResponseBodyList) *QueryPositionsResponseBody {
	s.List = v
	return s
}

func (s *QueryPositionsResponseBody) SetNextToken(v int64) *QueryPositionsResponseBody {
	s.NextToken = &v
	return s
}

type QueryPositionsResponseBodyList struct {
	// 所属职务ID
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// 职位类别ID
	PositionCategoryId *string `json:"positionCategoryId,omitempty" xml:"positionCategoryId,omitempty"`
	// 职位描述
	PositionDes *string `json:"positionDes,omitempty" xml:"positionDes,omitempty"`
	// 职位ID
	PositionId *string `json:"positionId,omitempty" xml:"positionId,omitempty"`
	// 职位名称
	PositionName *string `json:"positionName,omitempty" xml:"positionName,omitempty"`
	// 职位对应职级列表
	RankIdList []*string `json:"rankIdList,omitempty" xml:"rankIdList,omitempty" type:"Repeated"`
	// 职位状态-0，启用；1，停用
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryPositionsResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryPositionsResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryPositionsResponseBodyList) SetJobId(v string) *QueryPositionsResponseBodyList {
	s.JobId = &v
	return s
}

func (s *QueryPositionsResponseBodyList) SetPositionCategoryId(v string) *QueryPositionsResponseBodyList {
	s.PositionCategoryId = &v
	return s
}

func (s *QueryPositionsResponseBodyList) SetPositionDes(v string) *QueryPositionsResponseBodyList {
	s.PositionDes = &v
	return s
}

func (s *QueryPositionsResponseBodyList) SetPositionId(v string) *QueryPositionsResponseBodyList {
	s.PositionId = &v
	return s
}

func (s *QueryPositionsResponseBodyList) SetPositionName(v string) *QueryPositionsResponseBodyList {
	s.PositionName = &v
	return s
}

func (s *QueryPositionsResponseBodyList) SetRankIdList(v []*string) *QueryPositionsResponseBodyList {
	s.RankIdList = v
	return s
}

func (s *QueryPositionsResponseBodyList) SetStatus(v int32) *QueryPositionsResponseBodyList {
	s.Status = &v
	return s
}

type QueryPositionsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryPositionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPositionsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPositionsResponse) GoString() string {
	return s.String()
}

func (s *QueryPositionsResponse) SetHeaders(v map[string]*string) *QueryPositionsResponse {
	s.Headers = v
	return s
}

func (s *QueryPositionsResponse) SetBody(v *QueryPositionsResponseBody) *QueryPositionsResponse {
	s.Body = v
	return s
}

type SolutionTaskInitHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SolutionTaskInitHeaders) String() string {
	return tea.Prettify(s)
}

func (s SolutionTaskInitHeaders) GoString() string {
	return s.String()
}

func (s *SolutionTaskInitHeaders) SetCommonHeaders(v map[string]*string) *SolutionTaskInitHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SolutionTaskInitHeaders) SetXAcsDingtalkAccessToken(v string) *SolutionTaskInitHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SolutionTaskInitRequest struct {
	// 任务业务模块，如training, performance等
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 任务要求的截止时间
	ClaimTime *int64 `json:"claimTime,omitempty" xml:"claimTime,omitempty"`
	// 任务描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 任务完成时间
	FinishTime *int64 `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// 外部的任务唯一标识
	OuterId *string `json:"outerId,omitempty" xml:"outerId,omitempty"`
	// 任务状态，如running,finished
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 任务名称
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 任务执行人userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 解决方案类型
	SolutionType *string `json:"solutionType,omitempty" xml:"solutionType,omitempty"`
}

func (s SolutionTaskInitRequest) String() string {
	return tea.Prettify(s)
}

func (s SolutionTaskInitRequest) GoString() string {
	return s.String()
}

func (s *SolutionTaskInitRequest) SetCategory(v string) *SolutionTaskInitRequest {
	s.Category = &v
	return s
}

func (s *SolutionTaskInitRequest) SetClaimTime(v int64) *SolutionTaskInitRequest {
	s.ClaimTime = &v
	return s
}

func (s *SolutionTaskInitRequest) SetDescription(v string) *SolutionTaskInitRequest {
	s.Description = &v
	return s
}

func (s *SolutionTaskInitRequest) SetFinishTime(v int64) *SolutionTaskInitRequest {
	s.FinishTime = &v
	return s
}

func (s *SolutionTaskInitRequest) SetOuterId(v string) *SolutionTaskInitRequest {
	s.OuterId = &v
	return s
}

func (s *SolutionTaskInitRequest) SetStatus(v string) *SolutionTaskInitRequest {
	s.Status = &v
	return s
}

func (s *SolutionTaskInitRequest) SetTitle(v string) *SolutionTaskInitRequest {
	s.Title = &v
	return s
}

func (s *SolutionTaskInitRequest) SetUserId(v string) *SolutionTaskInitRequest {
	s.UserId = &v
	return s
}

func (s *SolutionTaskInitRequest) SetSolutionType(v string) *SolutionTaskInitRequest {
	s.SolutionType = &v
	return s
}

type SolutionTaskInitResponseBody struct {
	// 数据是否初始化成功
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SolutionTaskInitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SolutionTaskInitResponseBody) GoString() string {
	return s.String()
}

func (s *SolutionTaskInitResponseBody) SetResult(v bool) *SolutionTaskInitResponseBody {
	s.Result = &v
	return s
}

type SolutionTaskInitResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SolutionTaskInitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SolutionTaskInitResponse) String() string {
	return tea.Prettify(s)
}

func (s SolutionTaskInitResponse) GoString() string {
	return s.String()
}

func (s *SolutionTaskInitResponse) SetHeaders(v map[string]*string) *SolutionTaskInitResponse {
	s.Headers = v
	return s
}

func (s *SolutionTaskInitResponse) SetBody(v *SolutionTaskInitResponseBody) *SolutionTaskInitResponse {
	s.Body = v
	return s
}

type SolutionTaskSaveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SolutionTaskSaveHeaders) String() string {
	return tea.Prettify(s)
}

func (s SolutionTaskSaveHeaders) GoString() string {
	return s.String()
}

func (s *SolutionTaskSaveHeaders) SetCommonHeaders(v map[string]*string) *SolutionTaskSaveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SolutionTaskSaveHeaders) SetXAcsDingtalkAccessToken(v string) *SolutionTaskSaveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SolutionTaskSaveRequest struct {
	// 任务要求的截止时间
	ClaimTime *int64 `json:"claimTime,omitempty" xml:"claimTime,omitempty"`
	// 任务描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 任务完成时间
	FinishTime *int64 `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// 外部的任务唯一标识
	OuterId            *string `json:"outerId,omitempty" xml:"outerId,omitempty"`
	SolutionInstanceId *string `json:"solutionInstanceId,omitempty" xml:"solutionInstanceId,omitempty"`
	StartTime          *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 任务状态，如running,finished
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 任务业务模块，如training, performance等
	TaskType        *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	TemplateOuterId *string `json:"templateOuterId,omitempty" xml:"templateOuterId,omitempty"`
	// 任务名称
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 任务执行人userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 解决方案类型
	SolutionType *string `json:"solutionType,omitempty" xml:"solutionType,omitempty"`
}

func (s SolutionTaskSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s SolutionTaskSaveRequest) GoString() string {
	return s.String()
}

func (s *SolutionTaskSaveRequest) SetClaimTime(v int64) *SolutionTaskSaveRequest {
	s.ClaimTime = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetDescription(v string) *SolutionTaskSaveRequest {
	s.Description = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetFinishTime(v int64) *SolutionTaskSaveRequest {
	s.FinishTime = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetOuterId(v string) *SolutionTaskSaveRequest {
	s.OuterId = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetSolutionInstanceId(v string) *SolutionTaskSaveRequest {
	s.SolutionInstanceId = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetStartTime(v int64) *SolutionTaskSaveRequest {
	s.StartTime = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetStatus(v string) *SolutionTaskSaveRequest {
	s.Status = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetTaskType(v string) *SolutionTaskSaveRequest {
	s.TaskType = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetTemplateOuterId(v string) *SolutionTaskSaveRequest {
	s.TemplateOuterId = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetTitle(v string) *SolutionTaskSaveRequest {
	s.Title = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetUserId(v string) *SolutionTaskSaveRequest {
	s.UserId = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetSolutionType(v string) *SolutionTaskSaveRequest {
	s.SolutionType = &v
	return s
}

type SolutionTaskSaveResponseBody struct {
	// 数据是否保存成功
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SolutionTaskSaveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SolutionTaskSaveResponseBody) GoString() string {
	return s.String()
}

func (s *SolutionTaskSaveResponseBody) SetResult(v bool) *SolutionTaskSaveResponseBody {
	s.Result = &v
	return s
}

type SolutionTaskSaveResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SolutionTaskSaveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SolutionTaskSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s SolutionTaskSaveResponse) GoString() string {
	return s.String()
}

func (s *SolutionTaskSaveResponse) SetHeaders(v map[string]*string) *SolutionTaskSaveResponse {
	s.Headers = v
	return s
}

func (s *SolutionTaskSaveResponse) SetBody(v *SolutionTaskSaveResponseBody) *SolutionTaskSaveResponse {
	s.Body = v
	return s
}

type SyncTaskTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SyncTaskTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncTaskTemplateHeaders) GoString() string {
	return s.String()
}

func (s *SyncTaskTemplateHeaders) SetCommonHeaders(v map[string]*string) *SyncTaskTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncTaskTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *SyncTaskTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SyncTaskTemplateRequest struct {
	// 是否删除任务模版，true删除，false不删除
	Delete *bool `json:"delete,omitempty" xml:"delete,omitempty"`
	// 任务模板描述
	Des *string `json:"des,omitempty" xml:"des,omitempty"`
	// 扩展信息，json串
	Ext *string `json:"ext,omitempty" xml:"ext,omitempty"`
	// 模版名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 任务模版创建人staffId
	OptUserId *string `json:"optUserId,omitempty" xml:"optUserId,omitempty"`
	// isv对应的任务模版唯一键
	OuterId *string `json:"outerId,omitempty" xml:"outerId,omitempty"`
	// 圈人规则
	TaskScopeVO *SyncTaskTemplateRequestTaskScopeVO `json:"taskScopeVO,omitempty" xml:"taskScopeVO,omitempty" type:"Struct"`
	// 任务模版类型：TRAIN_TASK、PERFORMANCE_TASK
	TaskType     *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	SolutionType *string `json:"solutionType,omitempty" xml:"solutionType,omitempty"`
}

func (s SyncTaskTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncTaskTemplateRequest) GoString() string {
	return s.String()
}

func (s *SyncTaskTemplateRequest) SetDelete(v bool) *SyncTaskTemplateRequest {
	s.Delete = &v
	return s
}

func (s *SyncTaskTemplateRequest) SetDes(v string) *SyncTaskTemplateRequest {
	s.Des = &v
	return s
}

func (s *SyncTaskTemplateRequest) SetExt(v string) *SyncTaskTemplateRequest {
	s.Ext = &v
	return s
}

func (s *SyncTaskTemplateRequest) SetName(v string) *SyncTaskTemplateRequest {
	s.Name = &v
	return s
}

func (s *SyncTaskTemplateRequest) SetOptUserId(v string) *SyncTaskTemplateRequest {
	s.OptUserId = &v
	return s
}

func (s *SyncTaskTemplateRequest) SetOuterId(v string) *SyncTaskTemplateRequest {
	s.OuterId = &v
	return s
}

func (s *SyncTaskTemplateRequest) SetTaskScopeVO(v *SyncTaskTemplateRequestTaskScopeVO) *SyncTaskTemplateRequest {
	s.TaskScopeVO = v
	return s
}

func (s *SyncTaskTemplateRequest) SetTaskType(v string) *SyncTaskTemplateRequest {
	s.TaskType = &v
	return s
}

func (s *SyncTaskTemplateRequest) SetSolutionType(v string) *SyncTaskTemplateRequest {
	s.SolutionType = &v
	return s
}

type SyncTaskTemplateRequestTaskScopeVO struct {
	// 按照部门圈人
	DeptIds []*int64 `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	// 按照职位圈人
	PositionIds []*string `json:"positionIds,omitempty" xml:"positionIds,omitempty" type:"Repeated"`
	// 按照角色圈人
	RoleIds []*string `json:"roleIds,omitempty" xml:"roleIds,omitempty" type:"Repeated"`
	// 按照员工userId圈人
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s SyncTaskTemplateRequestTaskScopeVO) String() string {
	return tea.Prettify(s)
}

func (s SyncTaskTemplateRequestTaskScopeVO) GoString() string {
	return s.String()
}

func (s *SyncTaskTemplateRequestTaskScopeVO) SetDeptIds(v []*int64) *SyncTaskTemplateRequestTaskScopeVO {
	s.DeptIds = v
	return s
}

func (s *SyncTaskTemplateRequestTaskScopeVO) SetPositionIds(v []*string) *SyncTaskTemplateRequestTaskScopeVO {
	s.PositionIds = v
	return s
}

func (s *SyncTaskTemplateRequestTaskScopeVO) SetRoleIds(v []*string) *SyncTaskTemplateRequestTaskScopeVO {
	s.RoleIds = v
	return s
}

func (s *SyncTaskTemplateRequestTaskScopeVO) SetUserIds(v []*string) *SyncTaskTemplateRequestTaskScopeVO {
	s.UserIds = v
	return s
}

type SyncTaskTemplateResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SyncTaskTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncTaskTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *SyncTaskTemplateResponseBody) SetResult(v bool) *SyncTaskTemplateResponseBody {
	s.Result = &v
	return s
}

type SyncTaskTemplateResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SyncTaskTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncTaskTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncTaskTemplateResponse) GoString() string {
	return s.String()
}

func (s *SyncTaskTemplateResponse) SetHeaders(v map[string]*string) *SyncTaskTemplateResponse {
	s.Headers = v
	return s
}

func (s *SyncTaskTemplateResponse) SetBody(v *SyncTaskTemplateResponseBody) *SyncTaskTemplateResponse {
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

func (client *Client) AddHrmPreentry(request *AddHrmPreentryRequest) (_result *AddHrmPreentryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddHrmPreentryHeaders{}
	_result = &AddHrmPreentryResponse{}
	_body, _err := client.AddHrmPreentryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddHrmPreentryWithOptions(request *AddHrmPreentryRequest, headers *AddHrmPreentryHeaders, runtime *util.RuntimeOptions) (_result *AddHrmPreentryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		body["agentId"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.Groups)) {
		body["groups"] = request.Groups
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		body["mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PreEntryTime)) {
		body["preEntryTime"] = request.PreEntryTime
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
	_result = &AddHrmPreentryResponse{}
	_body, _err := client.DoROARequest(tea.String("AddHrmPreentry"), tea.String("hrm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/hrm/preentries"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ECertQuery(request *ECertQueryRequest) (_result *ECertQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ECertQueryHeaders{}
	_result = &ECertQueryResponse{}
	_body, _err := client.ECertQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ECertQueryWithOptions(request *ECertQueryRequest, headers *ECertQueryHeaders, runtime *util.RuntimeOptions) (_result *ECertQueryResponse, _err error) {
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
	_result = &ECertQueryResponse{}
	_body, _err := client.DoROARequest(tea.String("ECertQuery"), tea.String("hrm_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/hrm/eCerts"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HrmProcessTransfer(request *HrmProcessTransferRequest) (_result *HrmProcessTransferResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrmProcessTransferHeaders{}
	_result = &HrmProcessTransferResponse{}
	_body, _err := client.HrmProcessTransferWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HrmProcessTransferWithOptions(request *HrmProcessTransferRequest, headers *HrmProcessTransferHeaders, runtime *util.RuntimeOptions) (_result *HrmProcessTransferResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptIdsAfterTransfer)) {
		body["deptIdsAfterTransfer"] = request.DeptIdsAfterTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.JobIdAfterTransfer)) {
		body["jobIdAfterTransfer"] = request.JobIdAfterTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.MainDeptIdAfterTransfer)) {
		body["mainDeptIdAfterTransfer"] = request.MainDeptIdAfterTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.OperateUserId)) {
		body["operateUserId"] = request.OperateUserId
	}

	if !tea.BoolValue(util.IsUnset(request.PositionIdAfterTransfer)) {
		body["positionIdAfterTransfer"] = request.PositionIdAfterTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.PositionLevelAfterTransfer)) {
		body["positionLevelAfterTransfer"] = request.PositionLevelAfterTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.PositionNameAfterTransfer)) {
		body["positionNameAfterTransfer"] = request.PositionNameAfterTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.RankIdAfterTransfer)) {
		body["rankIdAfterTransfer"] = request.RankIdAfterTransfer
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
	_result = &HrmProcessTransferResponse{}
	_body, _err := client.DoROARequest(tea.String("HrmProcessTransfer"), tea.String("hrm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/hrm/processes/transfer"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MasterDataQuery(request *MasterDataQueryRequest) (_result *MasterDataQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MasterDataQueryHeaders{}
	_result = &MasterDataQueryResponse{}
	_body, _err := client.MasterDataQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MasterDataQueryWithOptions(request *MasterDataQueryRequest, headers *MasterDataQueryHeaders, runtime *util.RuntimeOptions) (_result *MasterDataQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizUK)) {
		body["bizUK"] = request.BizUK
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OptUserId)) {
		body["optUserId"] = request.OptUserId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryParams)) {
		body["queryParams"] = request.QueryParams
	}

	if !tea.BoolValue(util.IsUnset(request.RelationIds)) {
		body["relationIds"] = request.RelationIds
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeCode)) {
		body["scopeCode"] = request.ScopeCode
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["tenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.ViewEntityCode)) {
		body["viewEntityCode"] = request.ViewEntityCode
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
	_result = &MasterDataQueryResponse{}
	_body, _err := client.DoROARequest(tea.String("MasterDataQuery"), tea.String("hrm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/hrm/masters/datas/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MasterDataSave(request *MasterDataSaveRequest) (_result *MasterDataSaveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MasterDataSaveHeaders{}
	_result = &MasterDataSaveResponse{}
	_body, _err := client.MasterDataSaveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MasterDataSaveWithOptions(request *MasterDataSaveRequest, headers *MasterDataSaveHeaders, runtime *util.RuntimeOptions) (_result *MasterDataSaveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
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
		Body:    util.ToArray(request.Body),
	}
	_result = &MasterDataSaveResponse{}
	_body, _err := client.DoROARequest(tea.String("MasterDataSave"), tea.String("hrm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/hrm/masters/datas/save"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MasterDataTenantQuey(request *MasterDataTenantQueyRequest) (_result *MasterDataTenantQueyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MasterDataTenantQueyHeaders{}
	_result = &MasterDataTenantQueyResponse{}
	_body, _err := client.MasterDataTenantQueyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MasterDataTenantQueyWithOptions(request *MasterDataTenantQueyRequest, headers *MasterDataTenantQueyHeaders, runtime *util.RuntimeOptions) (_result *MasterDataTenantQueyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityCode)) {
		query["entityCode"] = request.EntityCode
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeCode)) {
		query["scopeCode"] = request.ScopeCode
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
	_result = &MasterDataTenantQueyResponse{}
	_body, _err := client.DoROARequest(tea.String("MasterDataTenantQuey"), tea.String("hrm_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/hrm/masters/tenants"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCustomEntryProcesses(request *QueryCustomEntryProcessesRequest) (_result *QueryCustomEntryProcessesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCustomEntryProcessesHeaders{}
	_result = &QueryCustomEntryProcessesResponse{}
	_body, _err := client.QueryCustomEntryProcessesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCustomEntryProcessesWithOptions(request *QueryCustomEntryProcessesRequest, headers *QueryCustomEntryProcessesHeaders, runtime *util.RuntimeOptions) (_result *QueryCustomEntryProcessesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OperateUserId)) {
		query["operateUserId"] = request.OperateUserId
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
	_result = &QueryCustomEntryProcessesResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryCustomEntryProcesses"), tea.String("hrm_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/hrm/customEntryProcesses"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryHrmEmployeeDismissionInfo(request *QueryHrmEmployeeDismissionInfoRequest) (_result *QueryHrmEmployeeDismissionInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryHrmEmployeeDismissionInfoHeaders{}
	_result = &QueryHrmEmployeeDismissionInfoResponse{}
	_body, _err := client.QueryHrmEmployeeDismissionInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryHrmEmployeeDismissionInfoWithOptions(tmpReq *QueryHrmEmployeeDismissionInfoRequest, headers *QueryHrmEmployeeDismissionInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryHrmEmployeeDismissionInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryHrmEmployeeDismissionInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UserIdList)) {
		request.UserIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIdList, tea.String("userIdList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserIdListShrink)) {
		query["userIdList"] = request.UserIdListShrink
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
	_result = &QueryHrmEmployeeDismissionInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryHrmEmployeeDismissionInfo"), tea.String("hrm_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/hrm/employees/dimissionInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryJobRanks(request *QueryJobRanksRequest) (_result *QueryJobRanksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryJobRanksHeaders{}
	_result = &QueryJobRanksResponse{}
	_body, _err := client.QueryJobRanksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryJobRanksWithOptions(request *QueryJobRanksRequest, headers *QueryJobRanksHeaders, runtime *util.RuntimeOptions) (_result *QueryJobRanksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RankCategoryId)) {
		query["rankCategoryId"] = request.RankCategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.RankCode)) {
		query["rankCode"] = request.RankCode
	}

	if !tea.BoolValue(util.IsUnset(request.RankName)) {
		query["rankName"] = request.RankName
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
	_result = &QueryJobRanksResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryJobRanks"), tea.String("hrm_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/hrm/jobRanks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryJobs(request *QueryJobsRequest) (_result *QueryJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryJobsHeaders{}
	_result = &QueryJobsResponse{}
	_body, _err := client.QueryJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryJobsWithOptions(request *QueryJobsRequest, headers *QueryJobsHeaders, runtime *util.RuntimeOptions) (_result *QueryJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobName)) {
		query["jobName"] = request.JobName
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
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
	_result = &QueryJobsResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryJobs"), tea.String("hrm_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/hrm/jobs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPositions(request *QueryPositionsRequest) (_result *QueryPositionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryPositionsHeaders{}
	_result = &QueryPositionsResponse{}
	_body, _err := client.QueryPositionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPositionsWithOptions(request *QueryPositionsRequest, headers *QueryPositionsHeaders, runtime *util.RuntimeOptions) (_result *QueryPositionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.InCategoryIds)) {
		body["inCategoryIds"] = request.InCategoryIds
	}

	if !tea.BoolValue(util.IsUnset(request.InPositionIds)) {
		body["inPositionIds"] = request.InPositionIds
	}

	if !tea.BoolValue(util.IsUnset(request.PositionName)) {
		body["positionName"] = request.PositionName
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &QueryPositionsResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryPositions"), tea.String("hrm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/hrm/positions/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SolutionTaskInit(request *SolutionTaskInitRequest) (_result *SolutionTaskInitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SolutionTaskInitHeaders{}
	_result = &SolutionTaskInitResponse{}
	_body, _err := client.SolutionTaskInitWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SolutionTaskInitWithOptions(request *SolutionTaskInitRequest, headers *SolutionTaskInitHeaders, runtime *util.RuntimeOptions) (_result *SolutionTaskInitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SolutionType)) {
		query["solutionType"] = request.SolutionType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.ClaimTime)) {
		body["claimTime"] = request.ClaimTime
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FinishTime)) {
		body["finishTime"] = request.FinishTime
	}

	if !tea.BoolValue(util.IsUnset(request.OuterId)) {
		body["outerId"] = request.OuterId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
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
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &SolutionTaskInitResponse{}
	_body, _err := client.DoROARequest(tea.String("SolutionTaskInit"), tea.String("hrm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/hrm/solutions/tasks/init"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SolutionTaskSave(request *SolutionTaskSaveRequest) (_result *SolutionTaskSaveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SolutionTaskSaveHeaders{}
	_result = &SolutionTaskSaveResponse{}
	_body, _err := client.SolutionTaskSaveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SolutionTaskSaveWithOptions(request *SolutionTaskSaveRequest, headers *SolutionTaskSaveHeaders, runtime *util.RuntimeOptions) (_result *SolutionTaskSaveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SolutionType)) {
		query["solutionType"] = request.SolutionType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClaimTime)) {
		body["claimTime"] = request.ClaimTime
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FinishTime)) {
		body["finishTime"] = request.FinishTime
	}

	if !tea.BoolValue(util.IsUnset(request.OuterId)) {
		body["outerId"] = request.OuterId
	}

	if !tea.BoolValue(util.IsUnset(request.SolutionInstanceId)) {
		body["solutionInstanceId"] = request.SolutionInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		body["taskType"] = request.TaskType
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateOuterId)) {
		body["templateOuterId"] = request.TemplateOuterId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
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
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &SolutionTaskSaveResponse{}
	_body, _err := client.DoROARequest(tea.String("SolutionTaskSave"), tea.String("hrm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/hrm/solutions/tasks/save"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncTaskTemplate(request *SyncTaskTemplateRequest) (_result *SyncTaskTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncTaskTemplateHeaders{}
	_result = &SyncTaskTemplateResponse{}
	_body, _err := client.SyncTaskTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncTaskTemplateWithOptions(request *SyncTaskTemplateRequest, headers *SyncTaskTemplateHeaders, runtime *util.RuntimeOptions) (_result *SyncTaskTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SolutionType)) {
		query["solutionType"] = request.SolutionType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Delete)) {
		body["delete"] = request.Delete
	}

	if !tea.BoolValue(util.IsUnset(request.Des)) {
		body["des"] = request.Des
	}

	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OptUserId)) {
		body["optUserId"] = request.OptUserId
	}

	if !tea.BoolValue(util.IsUnset(request.OuterId)) {
		body["outerId"] = request.OuterId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TaskScopeVO))) {
		body["taskScopeVO"] = request.TaskScopeVO
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		body["taskType"] = request.TaskType
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &SyncTaskTemplateResponse{}
	_body, _err := client.DoROARequest(tea.String("SyncTaskTemplate"), tea.String("hrm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/hrm/solutions/tasks/templates/sync"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
