// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package resident_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type UpdateResideceGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateResideceGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateResideceGroupHeaders) GoString() string {
	return s.String()
}

func (s *UpdateResideceGroupHeaders) SetCommonHeaders(v map[string]*string) *UpdateResideceGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateResideceGroupHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateResideceGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateResideceGroupRequest struct {
	// 组长userid
	ManagerUserId *string `json:"managerUserId,omitempty" xml:"managerUserId,omitempty"`
	// 组名字
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 组id
	DepartmentId *int64 `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
}

func (s UpdateResideceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResideceGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateResideceGroupRequest) SetManagerUserId(v string) *UpdateResideceGroupRequest {
	s.ManagerUserId = &v
	return s
}

func (s *UpdateResideceGroupRequest) SetDepartmentName(v string) *UpdateResideceGroupRequest {
	s.DepartmentName = &v
	return s
}

func (s *UpdateResideceGroupRequest) SetDepartmentId(v int64) *UpdateResideceGroupRequest {
	s.DepartmentId = &v
	return s
}

type UpdateResideceGroupResponseBody struct {
	// 是否更新成功
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateResideceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResideceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResideceGroupResponseBody) SetResult(v bool) *UpdateResideceGroupResponseBody {
	s.Result = &v
	return s
}

type UpdateResideceGroupResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateResideceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateResideceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResideceGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateResideceGroupResponse) SetHeaders(v map[string]*string) *UpdateResideceGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateResideceGroupResponse) SetBody(v *UpdateResideceGroupResponseBody) *UpdateResideceGroupResponse {
	s.Body = v
	return s
}

type AddPointHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddPointHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddPointHeaders) GoString() string {
	return s.String()
}

func (s *AddPointHeaders) SetCommonHeaders(v map[string]*string) *AddPointHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddPointHeaders) SetXAcsDingtalkAccessToken(v string) *AddPointHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddPointRequest struct {
	// 是否查询全员圈积分
	IsCircle *bool `json:"isCircle,omitempty" xml:"isCircle,omitempty"`
	// 加积分的唯一幂等标志
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 成员id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 规则代码（可空）,如果不为空的话，score值使用ruleCode对应的score增加分数
	RuleCode *string `json:"ruleCode,omitempty" xml:"ruleCode,omitempty"`
	// 规则名字
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// 增加积分的时间戳毫秒数，如果为空使用系统当前毫秒数
	ActionTime *int64 `json:"actionTime,omitempty" xml:"actionTime,omitempty"`
	// 本次增加积分：正数表示增加/负数表示扣减
	Score *int32 `json:"score,omitempty" xml:"score,omitempty"`
}

func (s AddPointRequest) String() string {
	return tea.Prettify(s)
}

func (s AddPointRequest) GoString() string {
	return s.String()
}

func (s *AddPointRequest) SetIsCircle(v bool) *AddPointRequest {
	s.IsCircle = &v
	return s
}

func (s *AddPointRequest) SetUuid(v string) *AddPointRequest {
	s.Uuid = &v
	return s
}

func (s *AddPointRequest) SetUserId(v string) *AddPointRequest {
	s.UserId = &v
	return s
}

func (s *AddPointRequest) SetRuleCode(v string) *AddPointRequest {
	s.RuleCode = &v
	return s
}

func (s *AddPointRequest) SetRuleName(v string) *AddPointRequest {
	s.RuleName = &v
	return s
}

func (s *AddPointRequest) SetActionTime(v int64) *AddPointRequest {
	s.ActionTime = &v
	return s
}

func (s *AddPointRequest) SetScore(v int32) *AddPointRequest {
	s.Score = &v
	return s
}

type AddPointResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s AddPointResponse) String() string {
	return tea.Prettify(s)
}

func (s AddPointResponse) GoString() string {
	return s.String()
}

func (s *AddPointResponse) SetHeaders(v map[string]*string) *AddPointResponse {
	s.Headers = v
	return s
}

type DeleteResidentDepartmentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteResidentDepartmentHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteResidentDepartmentHeaders) GoString() string {
	return s.String()
}

func (s *DeleteResidentDepartmentHeaders) SetCommonHeaders(v map[string]*string) *DeleteResidentDepartmentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteResidentDepartmentHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteResidentDepartmentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteResidentDepartmentRequest struct {
	// 组/户id
	DepartmentId *int64 `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
}

func (s DeleteResidentDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteResidentDepartmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteResidentDepartmentRequest) SetDepartmentId(v int64) *DeleteResidentDepartmentRequest {
	s.DepartmentId = &v
	return s
}

type DeleteResidentDepartmentResponseBody struct {
	// 是否删除成功
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteResidentDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResidentDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResidentDepartmentResponseBody) SetResult(v bool) *DeleteResidentDepartmentResponseBody {
	s.Result = &v
	return s
}

type DeleteResidentDepartmentResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteResidentDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteResidentDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResidentDepartmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteResidentDepartmentResponse) SetHeaders(v map[string]*string) *DeleteResidentDepartmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteResidentDepartmentResponse) SetBody(v *DeleteResidentDepartmentResponseBody) *DeleteResidentDepartmentResponse {
	s.Body = v
	return s
}

type AddResidentUsersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddResidentUsersHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddResidentUsersHeaders) GoString() string {
	return s.String()
}

func (s *AddResidentUsersHeaders) SetCommonHeaders(v map[string]*string) *AddResidentUsersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddResidentUsersHeaders) SetXAcsDingtalkAccessToken(v string) *AddResidentUsersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddResidentUsersRequest struct {
	// 家庭住址
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// 是否是租客
	IsLeaseholder *bool `json:"isLeaseholder,omitempty" xml:"isLeaseholder,omitempty"`
	// 居民名字
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// 手机号码
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 户/租户部门id
	DepartmentId *int64 `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 扩展字段（包括身份证/性别/民族）
	ExtField []*AddResidentUsersRequestExtField `json:"extField,omitempty" xml:"extField,omitempty" type:"Repeated"`
	// 与户主的关系
	RelateType *string `json:"relateType,omitempty" xml:"relateType,omitempty"`
}

func (s AddResidentUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddResidentUsersRequest) GoString() string {
	return s.String()
}

func (s *AddResidentUsersRequest) SetAddress(v string) *AddResidentUsersRequest {
	s.Address = &v
	return s
}

func (s *AddResidentUsersRequest) SetIsLeaseholder(v bool) *AddResidentUsersRequest {
	s.IsLeaseholder = &v
	return s
}

func (s *AddResidentUsersRequest) SetUserName(v string) *AddResidentUsersRequest {
	s.UserName = &v
	return s
}

func (s *AddResidentUsersRequest) SetMobile(v string) *AddResidentUsersRequest {
	s.Mobile = &v
	return s
}

func (s *AddResidentUsersRequest) SetDepartmentId(v int64) *AddResidentUsersRequest {
	s.DepartmentId = &v
	return s
}

func (s *AddResidentUsersRequest) SetExtField(v []*AddResidentUsersRequestExtField) *AddResidentUsersRequest {
	s.ExtField = v
	return s
}

func (s *AddResidentUsersRequest) SetRelateType(v string) *AddResidentUsersRequest {
	s.RelateType = &v
	return s
}

type AddResidentUsersRequestExtField struct {
	// 扩展字段值
	ItemValue *string `json:"itemValue,omitempty" xml:"itemValue,omitempty"`
	// 扩展字段名字
	ItemName *string `json:"itemName,omitempty" xml:"itemName,omitempty"`
}

func (s AddResidentUsersRequestExtField) String() string {
	return tea.Prettify(s)
}

func (s AddResidentUsersRequestExtField) GoString() string {
	return s.String()
}

func (s *AddResidentUsersRequestExtField) SetItemValue(v string) *AddResidentUsersRequestExtField {
	s.ItemValue = &v
	return s
}

func (s *AddResidentUsersRequestExtField) SetItemName(v string) *AddResidentUsersRequestExtField {
	s.ItemName = &v
	return s
}

type AddResidentUsersResponseBody struct {
	// 创建成功的userId
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s AddResidentUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddResidentUsersResponseBody) GoString() string {
	return s.String()
}

func (s *AddResidentUsersResponseBody) SetResult(v string) *AddResidentUsersResponseBody {
	s.Result = &v
	return s
}

type AddResidentUsersResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddResidentUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddResidentUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s AddResidentUsersResponse) GoString() string {
	return s.String()
}

func (s *AddResidentUsersResponse) SetHeaders(v map[string]*string) *AddResidentUsersResponse {
	s.Headers = v
	return s
}

func (s *AddResidentUsersResponse) SetBody(v *AddResidentUsersResponseBody) *AddResidentUsersResponse {
	s.Body = v
	return s
}

type AddResidentDepartmentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddResidentDepartmentHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddResidentDepartmentHeaders) GoString() string {
	return s.String()
}

func (s *AddResidentDepartmentHeaders) SetCommonHeaders(v map[string]*string) *AddResidentDepartmentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddResidentDepartmentHeaders) SetXAcsDingtalkAccessToken(v string) *AddResidentDepartmentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddResidentDepartmentRequest struct {
	// 是否为组
	IsResidenceGroup *bool `json:"isResidenceGroup,omitempty" xml:"isResidenceGroup,omitempty"`
	// 部门名字
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 父部门id
	ParentDepartmentId *int64 `json:"parentDepartmentId,omitempty" xml:"parentDepartmentId,omitempty"`
}

func (s AddResidentDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s AddResidentDepartmentRequest) GoString() string {
	return s.String()
}

func (s *AddResidentDepartmentRequest) SetIsResidenceGroup(v bool) *AddResidentDepartmentRequest {
	s.IsResidenceGroup = &v
	return s
}

func (s *AddResidentDepartmentRequest) SetDepartmentName(v string) *AddResidentDepartmentRequest {
	s.DepartmentName = &v
	return s
}

func (s *AddResidentDepartmentRequest) SetParentDepartmentId(v int64) *AddResidentDepartmentRequest {
	s.ParentDepartmentId = &v
	return s
}

type AddResidentDepartmentResponseBody struct {
	// 创建成功的deptId
	Result *int64 `json:"result,omitempty" xml:"result,omitempty"`
}

func (s AddResidentDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddResidentDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *AddResidentDepartmentResponseBody) SetResult(v int64) *AddResidentDepartmentResponseBody {
	s.Result = &v
	return s
}

type AddResidentDepartmentResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddResidentDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddResidentDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s AddResidentDepartmentResponse) GoString() string {
	return s.String()
}

func (s *AddResidentDepartmentResponse) SetHeaders(v map[string]*string) *AddResidentDepartmentResponse {
	s.Headers = v
	return s
}

func (s *AddResidentDepartmentResponse) SetBody(v *AddResidentDepartmentResponseBody) *AddResidentDepartmentResponse {
	s.Body = v
	return s
}

type PagePointHistoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PagePointHistoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s PagePointHistoryHeaders) GoString() string {
	return s.String()
}

func (s *PagePointHistoryHeaders) SetCommonHeaders(v map[string]*string) *PagePointHistoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PagePointHistoryHeaders) SetXAcsDingtalkAccessToken(v string) *PagePointHistoryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PagePointHistoryRequest struct {
	// 是否查询全员圈积分
	IsCircle *bool `json:"isCircle,omitempty" xml:"isCircle,omitempty"`
	// 用户userid，可空，不传表示查询组织内所有用户的流水数据
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用来标记当前开始读取的位置
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 本次读取的最大数据记录数量，最大20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 起始时间Unix时间戳，可空
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 结束时间Unix时间戳（不包含），可空
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
}

func (s PagePointHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s PagePointHistoryRequest) GoString() string {
	return s.String()
}

func (s *PagePointHistoryRequest) SetIsCircle(v bool) *PagePointHistoryRequest {
	s.IsCircle = &v
	return s
}

func (s *PagePointHistoryRequest) SetUserId(v string) *PagePointHistoryRequest {
	s.UserId = &v
	return s
}

func (s *PagePointHistoryRequest) SetNextToken(v int64) *PagePointHistoryRequest {
	s.NextToken = &v
	return s
}

func (s *PagePointHistoryRequest) SetMaxResults(v int32) *PagePointHistoryRequest {
	s.MaxResults = &v
	return s
}

func (s *PagePointHistoryRequest) SetStartTime(v int64) *PagePointHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *PagePointHistoryRequest) SetEndTime(v int64) *PagePointHistoryRequest {
	s.EndTime = &v
	return s
}

type PagePointHistoryResponseBody struct {
	// 查询所得积分流水集合
	PointRecordList []*PagePointHistoryResponseBodyPointRecordList `json:"pointRecordList,omitempty" xml:"pointRecordList,omitempty" type:"Repeated"`
	// 是否有下一页
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 下一个游标值
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 总数，如果为-1则不计算总数
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s PagePointHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PagePointHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *PagePointHistoryResponseBody) SetPointRecordList(v []*PagePointHistoryResponseBodyPointRecordList) *PagePointHistoryResponseBody {
	s.PointRecordList = v
	return s
}

func (s *PagePointHistoryResponseBody) SetHasMore(v bool) *PagePointHistoryResponseBody {
	s.HasMore = &v
	return s
}

func (s *PagePointHistoryResponseBody) SetNextToken(v int64) *PagePointHistoryResponseBody {
	s.NextToken = &v
	return s
}

func (s *PagePointHistoryResponseBody) SetTotalCount(v int64) *PagePointHistoryResponseBody {
	s.TotalCount = &v
	return s
}

type PagePointHistoryResponseBodyPointRecordList struct {
	// 组织id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 成员id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 增加或减少的分数（增加为正数，减少为负数）
	Score *int32 `json:"score,omitempty" xml:"score,omitempty"`
	// 创建时间（精确到毫秒数）
	CreateAt *int64 `json:"createAt,omitempty" xml:"createAt,omitempty"`
	// 幂等键
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 对应的行为代码（可空）
	RuleCode *string `json:"ruleCode,omitempty" xml:"ruleCode,omitempty"`
	// 对应的行为名字
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
}

func (s PagePointHistoryResponseBodyPointRecordList) String() string {
	return tea.Prettify(s)
}

func (s PagePointHistoryResponseBodyPointRecordList) GoString() string {
	return s.String()
}

func (s *PagePointHistoryResponseBodyPointRecordList) SetCorpId(v string) *PagePointHistoryResponseBodyPointRecordList {
	s.CorpId = &v
	return s
}

func (s *PagePointHistoryResponseBodyPointRecordList) SetUserId(v string) *PagePointHistoryResponseBodyPointRecordList {
	s.UserId = &v
	return s
}

func (s *PagePointHistoryResponseBodyPointRecordList) SetScore(v int32) *PagePointHistoryResponseBodyPointRecordList {
	s.Score = &v
	return s
}

func (s *PagePointHistoryResponseBodyPointRecordList) SetCreateAt(v int64) *PagePointHistoryResponseBodyPointRecordList {
	s.CreateAt = &v
	return s
}

func (s *PagePointHistoryResponseBodyPointRecordList) SetUuid(v string) *PagePointHistoryResponseBodyPointRecordList {
	s.Uuid = &v
	return s
}

func (s *PagePointHistoryResponseBodyPointRecordList) SetRuleCode(v string) *PagePointHistoryResponseBodyPointRecordList {
	s.RuleCode = &v
	return s
}

func (s *PagePointHistoryResponseBodyPointRecordList) SetRuleName(v string) *PagePointHistoryResponseBodyPointRecordList {
	s.RuleName = &v
	return s
}

type PagePointHistoryResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PagePointHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PagePointHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s PagePointHistoryResponse) GoString() string {
	return s.String()
}

func (s *PagePointHistoryResponse) SetHeaders(v map[string]*string) *PagePointHistoryResponse {
	s.Headers = v
	return s
}

func (s *PagePointHistoryResponse) SetBody(v *PagePointHistoryResponseBody) *PagePointHistoryResponse {
	s.Body = v
	return s
}

type RemoveResidentUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RemoveResidentUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveResidentUserHeaders) GoString() string {
	return s.String()
}

func (s *RemoveResidentUserHeaders) SetCommonHeaders(v map[string]*string) *RemoveResidentUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveResidentUserHeaders) SetXAcsDingtalkAccessToken(v string) *RemoveResidentUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RemoveResidentUserRequest struct {
	// 户/租户部门id
	DepartmentId *int64 `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s RemoveResidentUserRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveResidentUserRequest) GoString() string {
	return s.String()
}

func (s *RemoveResidentUserRequest) SetDepartmentId(v int64) *RemoveResidentUserRequest {
	s.DepartmentId = &v
	return s
}

func (s *RemoveResidentUserRequest) SetUserId(v string) *RemoveResidentUserRequest {
	s.UserId = &v
	return s
}

type RemoveResidentUserResponseBody struct {
	// 是否移除成功
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RemoveResidentUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveResidentUserResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveResidentUserResponseBody) SetResult(v bool) *RemoveResidentUserResponseBody {
	s.Result = &v
	return s
}

type RemoveResidentUserResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveResidentUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveResidentUserResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveResidentUserResponse) GoString() string {
	return s.String()
}

func (s *RemoveResidentUserResponse) SetHeaders(v map[string]*string) *RemoveResidentUserResponse {
	s.Headers = v
	return s
}

func (s *RemoveResidentUserResponse) SetBody(v *RemoveResidentUserResponseBody) *RemoveResidentUserResponse {
	s.Body = v
	return s
}

type UpdateResidenceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateResidenceHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidenceHeaders) GoString() string {
	return s.String()
}

func (s *UpdateResidenceHeaders) SetCommonHeaders(v map[string]*string) *UpdateResidenceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateResidenceHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateResidenceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateResidenceRequest struct {
	// 家庭管理员用户id
	ManagerUserId *string `json:"managerUserId,omitempty" xml:"managerUserId,omitempty"`
	// 户名字
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 组id
	DepartmentId *int64 `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 所属网格
	Grid *string `json:"grid,omitempty" xml:"grid,omitempty"`
	// 家庭电话
	HomeTel *string `json:"homeTel,omitempty" xml:"homeTel,omitempty"`
	// 是否是贫困户
	Destitute *bool `json:"destitute,omitempty" xml:"destitute,omitempty"`
	// 组id
	ParentDepartmentId *int64 `json:"parentDepartmentId,omitempty" xml:"parentDepartmentId,omitempty"`
}

func (s UpdateResidenceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidenceRequest) GoString() string {
	return s.String()
}

func (s *UpdateResidenceRequest) SetManagerUserId(v string) *UpdateResidenceRequest {
	s.ManagerUserId = &v
	return s
}

func (s *UpdateResidenceRequest) SetDepartmentName(v string) *UpdateResidenceRequest {
	s.DepartmentName = &v
	return s
}

func (s *UpdateResidenceRequest) SetDepartmentId(v int64) *UpdateResidenceRequest {
	s.DepartmentId = &v
	return s
}

func (s *UpdateResidenceRequest) SetGrid(v string) *UpdateResidenceRequest {
	s.Grid = &v
	return s
}

func (s *UpdateResidenceRequest) SetHomeTel(v string) *UpdateResidenceRequest {
	s.HomeTel = &v
	return s
}

func (s *UpdateResidenceRequest) SetDestitute(v bool) *UpdateResidenceRequest {
	s.Destitute = &v
	return s
}

func (s *UpdateResidenceRequest) SetParentDepartmentId(v int64) *UpdateResidenceRequest {
	s.ParentDepartmentId = &v
	return s
}

type UpdateResidenceResponseBody struct {
	// 是否更新成功
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateResidenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidenceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResidenceResponseBody) SetResult(v bool) *UpdateResidenceResponseBody {
	s.Result = &v
	return s
}

type UpdateResidenceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateResidenceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateResidenceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidenceResponse) GoString() string {
	return s.String()
}

func (s *UpdateResidenceResponse) SetHeaders(v map[string]*string) *UpdateResidenceResponse {
	s.Headers = v
	return s
}

func (s *UpdateResidenceResponse) SetBody(v *UpdateResidenceResponseBody) *UpdateResidenceResponse {
	s.Body = v
	return s
}

type ListPointRulesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListPointRulesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListPointRulesHeaders) GoString() string {
	return s.String()
}

func (s *ListPointRulesHeaders) SetCommonHeaders(v map[string]*string) *ListPointRulesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListPointRulesHeaders) SetXAcsDingtalkAccessToken(v string) *ListPointRulesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListPointRulesRequest struct {
	// 是否查询全员圈积分
	IsCircle *bool `json:"isCircle,omitempty" xml:"isCircle,omitempty"`
}

func (s ListPointRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPointRulesRequest) GoString() string {
	return s.String()
}

func (s *ListPointRulesRequest) SetIsCircle(v bool) *ListPointRulesRequest {
	s.IsCircle = &v
	return s
}

type ListPointRulesResponseBody struct {
	// 查询所得积分规则集合
	PointRuleList []*ListPointRulesResponseBodyPointRuleList `json:"pointRuleList,omitempty" xml:"pointRuleList,omitempty" type:"Repeated"`
}

func (s ListPointRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPointRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPointRulesResponseBody) SetPointRuleList(v []*ListPointRulesResponseBodyPointRuleList) *ListPointRulesResponseBody {
	s.PointRuleList = v
	return s
}

type ListPointRulesResponseBodyPointRuleList struct {
	// 组织id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 增加或减少的分数（增加为正数，减少为负数）
	Score *int32 `json:"score,omitempty" xml:"score,omitempty"`
	// 单日计次上限，0表示无上限
	DayLimitTimes *int32 `json:"dayLimitTimes,omitempty" xml:"dayLimitTimes,omitempty"`
	// 生效状态 0：不生效，1：生效
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// 对应的行为代码（可空）
	RuleCode *string `json:"ruleCode,omitempty" xml:"ruleCode,omitempty"`
	// 对应的行为名字
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// 扩展字段
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 分组ID, 默认写入为0
	GroupId *int32 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// 排序ID
	OrderId *int32 `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s ListPointRulesResponseBodyPointRuleList) String() string {
	return tea.Prettify(s)
}

func (s ListPointRulesResponseBodyPointRuleList) GoString() string {
	return s.String()
}

func (s *ListPointRulesResponseBodyPointRuleList) SetCorpId(v string) *ListPointRulesResponseBodyPointRuleList {
	s.CorpId = &v
	return s
}

func (s *ListPointRulesResponseBodyPointRuleList) SetScore(v int32) *ListPointRulesResponseBodyPointRuleList {
	s.Score = &v
	return s
}

func (s *ListPointRulesResponseBodyPointRuleList) SetDayLimitTimes(v int32) *ListPointRulesResponseBodyPointRuleList {
	s.DayLimitTimes = &v
	return s
}

func (s *ListPointRulesResponseBodyPointRuleList) SetStatus(v int32) *ListPointRulesResponseBodyPointRuleList {
	s.Status = &v
	return s
}

func (s *ListPointRulesResponseBodyPointRuleList) SetRuleCode(v string) *ListPointRulesResponseBodyPointRuleList {
	s.RuleCode = &v
	return s
}

func (s *ListPointRulesResponseBodyPointRuleList) SetRuleName(v string) *ListPointRulesResponseBodyPointRuleList {
	s.RuleName = &v
	return s
}

func (s *ListPointRulesResponseBodyPointRuleList) SetExtension(v string) *ListPointRulesResponseBodyPointRuleList {
	s.Extension = &v
	return s
}

func (s *ListPointRulesResponseBodyPointRuleList) SetGroupId(v int32) *ListPointRulesResponseBodyPointRuleList {
	s.GroupId = &v
	return s
}

func (s *ListPointRulesResponseBodyPointRuleList) SetOrderId(v int32) *ListPointRulesResponseBodyPointRuleList {
	s.OrderId = &v
	return s
}

type ListPointRulesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPointRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPointRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPointRulesResponse) GoString() string {
	return s.String()
}

func (s *ListPointRulesResponse) SetHeaders(v map[string]*string) *ListPointRulesResponse {
	s.Headers = v
	return s
}

func (s *ListPointRulesResponse) SetBody(v *ListPointRulesResponseBody) *ListPointRulesResponse {
	s.Body = v
	return s
}

type UpdateResidentUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateResidentUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidentUserHeaders) GoString() string {
	return s.String()
}

func (s *UpdateResidentUserHeaders) SetCommonHeaders(v map[string]*string) *UpdateResidentUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateResidentUserHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateResidentUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateResidentUserRequest struct {
	// 家庭住址
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// 是否保留原部门
	IsRetainOldDept *bool `json:"isRetainOldDept,omitempty" xml:"isRetainOldDept,omitempty"`
	// 居民名字
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// 手机号码
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 所在新的户/租户部门id
	DepartmentId *int64 `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 扩展字段（包括身份证/性别/民族）
	ExtField []*UpdateResidentUserRequestExtField `json:"extField,omitempty" xml:"extField,omitempty" type:"Repeated"`
	// 与户主的关系
	RelateType *string `json:"relateType,omitempty" xml:"relateType,omitempty"`
	// 人员userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 原所在部门id
	OldDepartmentId *int64 `json:"oldDepartmentId,omitempty" xml:"oldDepartmentId,omitempty"`
}

func (s UpdateResidentUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidentUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateResidentUserRequest) SetAddress(v string) *UpdateResidentUserRequest {
	s.Address = &v
	return s
}

func (s *UpdateResidentUserRequest) SetIsRetainOldDept(v bool) *UpdateResidentUserRequest {
	s.IsRetainOldDept = &v
	return s
}

func (s *UpdateResidentUserRequest) SetUserName(v string) *UpdateResidentUserRequest {
	s.UserName = &v
	return s
}

func (s *UpdateResidentUserRequest) SetMobile(v string) *UpdateResidentUserRequest {
	s.Mobile = &v
	return s
}

func (s *UpdateResidentUserRequest) SetDepartmentId(v int64) *UpdateResidentUserRequest {
	s.DepartmentId = &v
	return s
}

func (s *UpdateResidentUserRequest) SetExtField(v []*UpdateResidentUserRequestExtField) *UpdateResidentUserRequest {
	s.ExtField = v
	return s
}

func (s *UpdateResidentUserRequest) SetRelateType(v string) *UpdateResidentUserRequest {
	s.RelateType = &v
	return s
}

func (s *UpdateResidentUserRequest) SetUserId(v string) *UpdateResidentUserRequest {
	s.UserId = &v
	return s
}

func (s *UpdateResidentUserRequest) SetOldDepartmentId(v int64) *UpdateResidentUserRequest {
	s.OldDepartmentId = &v
	return s
}

type UpdateResidentUserRequestExtField struct {
	// 扩展字段值
	ItemValue *string `json:"itemValue,omitempty" xml:"itemValue,omitempty"`
	// 扩展字段名字
	ItemName *string `json:"itemName,omitempty" xml:"itemName,omitempty"`
}

func (s UpdateResidentUserRequestExtField) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidentUserRequestExtField) GoString() string {
	return s.String()
}

func (s *UpdateResidentUserRequestExtField) SetItemValue(v string) *UpdateResidentUserRequestExtField {
	s.ItemValue = &v
	return s
}

func (s *UpdateResidentUserRequestExtField) SetItemName(v string) *UpdateResidentUserRequestExtField {
	s.ItemName = &v
	return s
}

type UpdateResidentUserResponseBody struct {
	// 是否更新成功
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateResidentUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidentUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResidentUserResponseBody) SetResult(v bool) *UpdateResidentUserResponseBody {
	s.Result = &v
	return s
}

type UpdateResidentUserResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateResidentUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateResidentUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidentUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateResidentUserResponse) SetHeaders(v map[string]*string) *UpdateResidentUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateResidentUserResponse) SetBody(v *UpdateResidentUserResponseBody) *UpdateResidentUserResponse {
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

func (client *Client) UpdateResideceGroup(request *UpdateResideceGroupRequest) (_result *UpdateResideceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateResideceGroupHeaders{}
	_result = &UpdateResideceGroupResponse{}
	_body, _err := client.UpdateResideceGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateResideceGroupWithOptions(request *UpdateResideceGroupRequest, headers *UpdateResideceGroupHeaders, runtime *util.RuntimeOptions) (_result *UpdateResideceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ManagerUserId)) {
		query["managerUserId"] = request.ManagerUserId
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentName)) {
		query["departmentName"] = request.DepartmentName
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["departmentId"] = request.DepartmentId
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
	_result = &UpdateResideceGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateResideceGroup"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/resident/departments/updateResideceGroup"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddPoint(request *AddPointRequest) (_result *AddPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddPointHeaders{}
	_result = &AddPointResponse{}
	_body, _err := client.AddPointWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddPointWithOptions(request *AddPointRequest, headers *AddPointHeaders, runtime *util.RuntimeOptions) (_result *AddPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsCircle)) {
		query["isCircle"] = request.IsCircle
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["uuid"] = request.Uuid
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleCode)) {
		query["ruleCode"] = request.RuleCode
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["ruleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.ActionTime)) {
		query["actionTime"] = request.ActionTime
	}

	if !tea.BoolValue(util.IsUnset(request.Score)) {
		query["score"] = request.Score
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
	_result = &AddPointResponse{}
	_body, _err := client.DoROARequest(tea.String("AddPoint"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/resident/points"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteResidentDepartment(request *DeleteResidentDepartmentRequest) (_result *DeleteResidentDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteResidentDepartmentHeaders{}
	_result = &DeleteResidentDepartmentResponse{}
	_body, _err := client.DeleteResidentDepartmentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteResidentDepartmentWithOptions(request *DeleteResidentDepartmentRequest, headers *DeleteResidentDepartmentHeaders, runtime *util.RuntimeOptions) (_result *DeleteResidentDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["departmentId"] = request.DepartmentId
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
	_result = &DeleteResidentDepartmentResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteResidentDepartment"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/resident/departments"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddResidentUsers(request *AddResidentUsersRequest) (_result *AddResidentUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddResidentUsersHeaders{}
	_result = &AddResidentUsersResponse{}
	_body, _err := client.AddResidentUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddResidentUsersWithOptions(request *AddResidentUsersRequest, headers *AddResidentUsersHeaders, runtime *util.RuntimeOptions) (_result *AddResidentUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.IsLeaseholder)) {
		query["isLeaseholder"] = request.IsLeaseholder
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["userName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["departmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtField)) {
		query["extField"] = request.ExtField
	}

	if !tea.BoolValue(util.IsUnset(request.RelateType)) {
		query["relateType"] = request.RelateType
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
	_result = &AddResidentUsersResponse{}
	_body, _err := client.DoROARequest(tea.String("AddResidentUsers"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/resident/users"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddResidentDepartment(request *AddResidentDepartmentRequest) (_result *AddResidentDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddResidentDepartmentHeaders{}
	_result = &AddResidentDepartmentResponse{}
	_body, _err := client.AddResidentDepartmentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddResidentDepartmentWithOptions(request *AddResidentDepartmentRequest, headers *AddResidentDepartmentHeaders, runtime *util.RuntimeOptions) (_result *AddResidentDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsResidenceGroup)) {
		query["isResidenceGroup"] = request.IsResidenceGroup
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentName)) {
		query["departmentName"] = request.DepartmentName
	}

	if !tea.BoolValue(util.IsUnset(request.ParentDepartmentId)) {
		query["parentDepartmentId"] = request.ParentDepartmentId
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
	_result = &AddResidentDepartmentResponse{}
	_body, _err := client.DoROARequest(tea.String("AddResidentDepartment"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/resident/departments"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PagePointHistory(request *PagePointHistoryRequest) (_result *PagePointHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PagePointHistoryHeaders{}
	_result = &PagePointHistoryResponse{}
	_body, _err := client.PagePointHistoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PagePointHistoryWithOptions(request *PagePointHistoryRequest, headers *PagePointHistoryHeaders, runtime *util.RuntimeOptions) (_result *PagePointHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsCircle)) {
		query["isCircle"] = request.IsCircle
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
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
	_result = &PagePointHistoryResponse{}
	_body, _err := client.DoROARequest(tea.String("PagePointHistory"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/resident/points/records"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveResidentUser(request *RemoveResidentUserRequest) (_result *RemoveResidentUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveResidentUserHeaders{}
	_result = &RemoveResidentUserResponse{}
	_body, _err := client.RemoveResidentUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveResidentUserWithOptions(request *RemoveResidentUserRequest, headers *RemoveResidentUserHeaders, runtime *util.RuntimeOptions) (_result *RemoveResidentUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["departmentId"] = request.DepartmentId
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
	_result = &RemoveResidentUserResponse{}
	_body, _err := client.DoROARequest(tea.String("RemoveResidentUser"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/resident/users/remove"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateResidence(request *UpdateResidenceRequest) (_result *UpdateResidenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateResidenceHeaders{}
	_result = &UpdateResidenceResponse{}
	_body, _err := client.UpdateResidenceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateResidenceWithOptions(request *UpdateResidenceRequest, headers *UpdateResidenceHeaders, runtime *util.RuntimeOptions) (_result *UpdateResidenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ManagerUserId)) {
		query["managerUserId"] = request.ManagerUserId
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentName)) {
		query["departmentName"] = request.DepartmentName
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["departmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.Grid)) {
		query["grid"] = request.Grid
	}

	if !tea.BoolValue(util.IsUnset(request.HomeTel)) {
		query["homeTel"] = request.HomeTel
	}

	if !tea.BoolValue(util.IsUnset(request.Destitute)) {
		query["destitute"] = request.Destitute
	}

	if !tea.BoolValue(util.IsUnset(request.ParentDepartmentId)) {
		query["parentDepartmentId"] = request.ParentDepartmentId
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
	_result = &UpdateResidenceResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateResidence"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/resident/departments/updateResidece"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPointRules(request *ListPointRulesRequest) (_result *ListPointRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListPointRulesHeaders{}
	_result = &ListPointRulesResponse{}
	_body, _err := client.ListPointRulesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPointRulesWithOptions(request *ListPointRulesRequest, headers *ListPointRulesHeaders, runtime *util.RuntimeOptions) (_result *ListPointRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsCircle)) {
		query["isCircle"] = request.IsCircle
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
	_result = &ListPointRulesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListPointRules"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/resident/points/rules"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateResidentUser(request *UpdateResidentUserRequest) (_result *UpdateResidentUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateResidentUserHeaders{}
	_result = &UpdateResidentUserResponse{}
	_body, _err := client.UpdateResidentUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateResidentUserWithOptions(request *UpdateResidentUserRequest, headers *UpdateResidentUserHeaders, runtime *util.RuntimeOptions) (_result *UpdateResidentUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.IsRetainOldDept)) {
		query["isRetainOldDept"] = request.IsRetainOldDept
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["userName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["departmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtField)) {
		query["extField"] = request.ExtField
	}

	if !tea.BoolValue(util.IsUnset(request.RelateType)) {
		query["relateType"] = request.RelateType
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.OldDepartmentId)) {
		query["oldDepartmentId"] = request.OldDepartmentId
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
	_result = &UpdateResidentUserResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateResidentUser"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/resident/users"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
