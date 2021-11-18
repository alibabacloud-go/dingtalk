// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package workflow_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type QueryFormInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryFormInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryFormInstanceHeaders) GoString() string {
	return s.String()
}

func (s *QueryFormInstanceHeaders) SetCommonHeaders(v map[string]*string) *QueryFormInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryFormInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *QueryFormInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryFormInstanceRequest struct {
	// 表单实例id
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// 表单模板Code
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// 应用搭建id
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
}

func (s QueryFormInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFormInstanceRequest) GoString() string {
	return s.String()
}

func (s *QueryFormInstanceRequest) SetFormInstanceId(v string) *QueryFormInstanceRequest {
	s.FormInstanceId = &v
	return s
}

func (s *QueryFormInstanceRequest) SetFormCode(v string) *QueryFormInstanceRequest {
	s.FormCode = &v
	return s
}

func (s *QueryFormInstanceRequest) SetAppUuid(v string) *QueryFormInstanceRequest {
	s.AppUuid = &v
	return s
}

type QueryFormInstanceResponseBody struct {
	// 实例id
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// 表单数据
	FormInstDataList []*QueryFormInstanceResponseBodyFormInstDataList `json:"formInstDataList,omitempty" xml:"formInstDataList,omitempty" type:"Repeated"`
	// 应用搭建id
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// 表单模板id
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// 表单标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 创建人
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 修改人
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// 实例创建时间戳
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// 实例最近修改时间戳
	ModifyTimestamp *int64 `json:"modifyTimestamp,omitempty" xml:"modifyTimestamp,omitempty"`
	// 外联业务实例id
	OutInstanceId *string `json:"outInstanceId,omitempty" xml:"outInstanceId,omitempty"`
	// 外联业务code
	OutBizCode *string `json:"outBizCode,omitempty" xml:"outBizCode,omitempty"`
	// 扩展信息
	Attributes map[string]interface{} `json:"attributes,omitempty" xml:"attributes,omitempty"`
}

func (s QueryFormInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFormInstanceResponseBody) SetFormInstanceId(v string) *QueryFormInstanceResponseBody {
	s.FormInstanceId = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetFormInstDataList(v []*QueryFormInstanceResponseBodyFormInstDataList) *QueryFormInstanceResponseBody {
	s.FormInstDataList = v
	return s
}

func (s *QueryFormInstanceResponseBody) SetAppUuid(v string) *QueryFormInstanceResponseBody {
	s.AppUuid = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetFormCode(v string) *QueryFormInstanceResponseBody {
	s.FormCode = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetTitle(v string) *QueryFormInstanceResponseBody {
	s.Title = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetCreator(v string) *QueryFormInstanceResponseBody {
	s.Creator = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetModifier(v string) *QueryFormInstanceResponseBody {
	s.Modifier = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetCreateTimestamp(v int64) *QueryFormInstanceResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetModifyTimestamp(v int64) *QueryFormInstanceResponseBody {
	s.ModifyTimestamp = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetOutInstanceId(v string) *QueryFormInstanceResponseBody {
	s.OutInstanceId = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetOutBizCode(v string) *QueryFormInstanceResponseBody {
	s.OutBizCode = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetAttributes(v map[string]interface{}) *QueryFormInstanceResponseBody {
	s.Attributes = v
	return s
}

type QueryFormInstanceResponseBodyFormInstDataList struct {
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	BizAlias      *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	ExtendValue   *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	Label         *string `json:"label,omitempty" xml:"label,omitempty"`
	Value         *string `json:"value,omitempty" xml:"value,omitempty"`
	Key           *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s QueryFormInstanceResponseBodyFormInstDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryFormInstanceResponseBodyFormInstDataList) GoString() string {
	return s.String()
}

func (s *QueryFormInstanceResponseBodyFormInstDataList) SetComponentType(v string) *QueryFormInstanceResponseBodyFormInstDataList {
	s.ComponentType = &v
	return s
}

func (s *QueryFormInstanceResponseBodyFormInstDataList) SetBizAlias(v string) *QueryFormInstanceResponseBodyFormInstDataList {
	s.BizAlias = &v
	return s
}

func (s *QueryFormInstanceResponseBodyFormInstDataList) SetExtendValue(v string) *QueryFormInstanceResponseBodyFormInstDataList {
	s.ExtendValue = &v
	return s
}

func (s *QueryFormInstanceResponseBodyFormInstDataList) SetLabel(v string) *QueryFormInstanceResponseBodyFormInstDataList {
	s.Label = &v
	return s
}

func (s *QueryFormInstanceResponseBodyFormInstDataList) SetValue(v string) *QueryFormInstanceResponseBodyFormInstDataList {
	s.Value = &v
	return s
}

func (s *QueryFormInstanceResponseBodyFormInstDataList) SetKey(v string) *QueryFormInstanceResponseBodyFormInstDataList {
	s.Key = &v
	return s
}

type QueryFormInstanceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryFormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryFormInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFormInstanceResponse) GoString() string {
	return s.String()
}

func (s *QueryFormInstanceResponse) SetHeaders(v map[string]*string) *QueryFormInstanceResponse {
	s.Headers = v
	return s
}

func (s *QueryFormInstanceResponse) SetBody(v *QueryFormInstanceResponseBody) *QueryFormInstanceResponse {
	s.Body = v
	return s
}

type QueryFormByBizTypeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryFormByBizTypeHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryFormByBizTypeHeaders) GoString() string {
	return s.String()
}

func (s *QueryFormByBizTypeHeaders) SetCommonHeaders(v map[string]*string) *QueryFormByBizTypeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryFormByBizTypeHeaders) SetXAcsDingtalkAccessToken(v string) *QueryFormByBizTypeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryFormByBizTypeRequest struct {
	// 应用搭建id
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// 表单业务标识
	BizTypes []*string `json:"bizTypes,omitempty" xml:"bizTypes,omitempty" type:"Repeated"`
}

func (s QueryFormByBizTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFormByBizTypeRequest) GoString() string {
	return s.String()
}

func (s *QueryFormByBizTypeRequest) SetAppUuid(v string) *QueryFormByBizTypeRequest {
	s.AppUuid = &v
	return s
}

func (s *QueryFormByBizTypeRequest) SetBizTypes(v []*string) *QueryFormByBizTypeRequest {
	s.BizTypes = v
	return s
}

type QueryFormByBizTypeResponseBody struct {
	// 模板列表
	Result []*QueryFormByBizTypeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryFormByBizTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFormByBizTypeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFormByBizTypeResponseBody) SetResult(v []*QueryFormByBizTypeResponseBodyResult) *QueryFormByBizTypeResponseBody {
	s.Result = v
	return s
}

type QueryFormByBizTypeResponseBodyResult struct {
	// 创建人
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 应用搭建id
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// 模板code
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// 表单uuid
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 模板名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 模板描述
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 数据归属id
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// 表单类型，0为流程表单，1为数据表单
	AppType *int32 `json:"appType,omitempty" xml:"appType,omitempty"`
	// 业务标识
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 模板状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 创建时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 修改时间
	ModifedTime *int64 `json:"modifedTime,omitempty" xml:"modifedTime,omitempty"`
	// 表单控件描述
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s QueryFormByBizTypeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryFormByBizTypeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryFormByBizTypeResponseBodyResult) SetCreator(v string) *QueryFormByBizTypeResponseBodyResult {
	s.Creator = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetAppUuid(v string) *QueryFormByBizTypeResponseBodyResult {
	s.AppUuid = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetFormCode(v string) *QueryFormByBizTypeResponseBodyResult {
	s.FormCode = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetFormUuid(v string) *QueryFormByBizTypeResponseBodyResult {
	s.FormUuid = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetName(v string) *QueryFormByBizTypeResponseBodyResult {
	s.Name = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetMemo(v string) *QueryFormByBizTypeResponseBodyResult {
	s.Memo = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetOwnerId(v string) *QueryFormByBizTypeResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetAppType(v int32) *QueryFormByBizTypeResponseBodyResult {
	s.AppType = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetBizType(v string) *QueryFormByBizTypeResponseBodyResult {
	s.BizType = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetStatus(v string) *QueryFormByBizTypeResponseBodyResult {
	s.Status = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetCreateTime(v int64) *QueryFormByBizTypeResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetModifedTime(v int64) *QueryFormByBizTypeResponseBodyResult {
	s.ModifedTime = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetContent(v string) *QueryFormByBizTypeResponseBodyResult {
	s.Content = &v
	return s
}

type QueryFormByBizTypeResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryFormByBizTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryFormByBizTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFormByBizTypeResponse) GoString() string {
	return s.String()
}

func (s *QueryFormByBizTypeResponse) SetHeaders(v map[string]*string) *QueryFormByBizTypeResponse {
	s.Headers = v
	return s
}

func (s *QueryFormByBizTypeResponse) SetBody(v *QueryFormByBizTypeResponseBody) *QueryFormByBizTypeResponse {
	s.Body = v
	return s
}

type QueryAllProcessInstancesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllProcessInstancesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesHeaders) SetCommonHeaders(v map[string]*string) *QueryAllProcessInstancesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllProcessInstancesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllProcessInstancesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllProcessInstancesRequest struct {
	// 分页起始值
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 分页大小
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 开始时间
	StartTimeInMills *int64 `json:"startTimeInMills,omitempty" xml:"startTimeInMills,omitempty"`
	// 结束时间
	EndTimeInMills *int64 `json:"endTimeInMills,omitempty" xml:"endTimeInMills,omitempty"`
	// 模板编码
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// 应用编码
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
}

func (s QueryAllProcessInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesRequest) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesRequest) SetNextToken(v string) *QueryAllProcessInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *QueryAllProcessInstancesRequest) SetMaxResults(v int64) *QueryAllProcessInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryAllProcessInstancesRequest) SetStartTimeInMills(v int64) *QueryAllProcessInstancesRequest {
	s.StartTimeInMills = &v
	return s
}

func (s *QueryAllProcessInstancesRequest) SetEndTimeInMills(v int64) *QueryAllProcessInstancesRequest {
	s.EndTimeInMills = &v
	return s
}

func (s *QueryAllProcessInstancesRequest) SetProcessCode(v string) *QueryAllProcessInstancesRequest {
	s.ProcessCode = &v
	return s
}

func (s *QueryAllProcessInstancesRequest) SetAppUuid(v string) *QueryAllProcessInstancesRequest {
	s.AppUuid = &v
	return s
}

type QueryAllProcessInstancesResponseBody struct {
	// result
	Result *QueryAllProcessInstancesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryAllProcessInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesResponseBody) SetResult(v *QueryAllProcessInstancesResponseBodyResult) *QueryAllProcessInstancesResponseBody {
	s.Result = v
	return s
}

type QueryAllProcessInstancesResponseBodyResult struct {
	// 下次获取数据的游标
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 是否有更多数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 总数
	MaxResults *int64                                            `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	List       []*QueryAllProcessInstancesResponseBodyResultList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s QueryAllProcessInstancesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesResponseBodyResult) SetNextToken(v string) *QueryAllProcessInstancesResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResult) SetHasMore(v bool) *QueryAllProcessInstancesResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResult) SetMaxResults(v int64) *QueryAllProcessInstancesResponseBodyResult {
	s.MaxResults = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResult) SetList(v []*QueryAllProcessInstancesResponseBodyResultList) *QueryAllProcessInstancesResponseBodyResult {
	s.List = v
	return s
}

type QueryAllProcessInstancesResponseBodyResultList struct {
	// 流程实例ID
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// 主单实例Id
	MainProcessInstanceId *string `json:"mainProcessInstanceId,omitempty" xml:"mainProcessInstanceId,omitempty"`
	// 审批结束时间
	FinishTime *int64 `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// 附属单信息
	AttachedProcessInstanceIds *string `json:"attachedProcessInstanceIds,omitempty" xml:"attachedProcessInstanceIds,omitempty"`
	// 审批单编号
	BusinessId *string `json:"businessId,omitempty" xml:"businessId,omitempty"`
	// 审批单标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 发起人部门id
	OriginatorDeptId *string `json:"originatorDeptId,omitempty" xml:"originatorDeptId,omitempty"`
	// 审批结果
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// 审批单创建时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 发起者userId
	OriginatorUserid *string `json:"originatorUserid,omitempty" xml:"originatorUserid,omitempty"`
	// 审批单状态
	Status              *string                                                              `json:"status,omitempty" xml:"status,omitempty"`
	FormComponentValues []*QueryAllProcessInstancesResponseBodyResultListFormComponentValues `json:"formComponentValues,omitempty" xml:"formComponentValues,omitempty" type:"Repeated"`
}

func (s QueryAllProcessInstancesResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetProcessInstanceId(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.ProcessInstanceId = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetMainProcessInstanceId(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.MainProcessInstanceId = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetFinishTime(v int64) *QueryAllProcessInstancesResponseBodyResultList {
	s.FinishTime = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetAttachedProcessInstanceIds(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.AttachedProcessInstanceIds = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetBusinessId(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.BusinessId = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetTitle(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.Title = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetOriginatorDeptId(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.OriginatorDeptId = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetResult(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.Result = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetCreateTime(v int64) *QueryAllProcessInstancesResponseBodyResultList {
	s.CreateTime = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetOriginatorUserid(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.OriginatorUserid = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetStatus(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.Status = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetFormComponentValues(v []*QueryAllProcessInstancesResponseBodyResultListFormComponentValues) *QueryAllProcessInstancesResponseBodyResultList {
	s.FormComponentValues = v
	return s
}

type QueryAllProcessInstancesResponseBodyResultListFormComponentValues struct {
	// 控件名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 控件id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 控件数据
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 控件扩展数据
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
}

func (s QueryAllProcessInstancesResponseBodyResultListFormComponentValues) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesResponseBodyResultListFormComponentValues) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesResponseBodyResultListFormComponentValues) SetName(v string) *QueryAllProcessInstancesResponseBodyResultListFormComponentValues {
	s.Name = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListFormComponentValues) SetId(v string) *QueryAllProcessInstancesResponseBodyResultListFormComponentValues {
	s.Id = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListFormComponentValues) SetValue(v string) *QueryAllProcessInstancesResponseBodyResultListFormComponentValues {
	s.Value = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListFormComponentValues) SetExtValue(v string) *QueryAllProcessInstancesResponseBodyResultListFormComponentValues {
	s.ExtValue = &v
	return s
}

type QueryAllProcessInstancesResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAllProcessInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllProcessInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesResponse) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesResponse) SetHeaders(v map[string]*string) *QueryAllProcessInstancesResponse {
	s.Headers = v
	return s
}

func (s *QueryAllProcessInstancesResponse) SetBody(v *QueryAllProcessInstancesResponseBody) *QueryAllProcessInstancesResponse {
	s.Body = v
	return s
}

type QueryAllFormInstancesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllFormInstancesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllFormInstancesHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllFormInstancesHeaders) SetCommonHeaders(v map[string]*string) *QueryAllFormInstancesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllFormInstancesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllFormInstancesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllFormInstancesRequest struct {
	// 分页游标，第一次调用传空或者null
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 翻页size
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 应用搭建id
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// 表单模板id
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
}

func (s QueryAllFormInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllFormInstancesRequest) GoString() string {
	return s.String()
}

func (s *QueryAllFormInstancesRequest) SetNextToken(v string) *QueryAllFormInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *QueryAllFormInstancesRequest) SetMaxResults(v int32) *QueryAllFormInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryAllFormInstancesRequest) SetAppUuid(v string) *QueryAllFormInstancesRequest {
	s.AppUuid = &v
	return s
}

func (s *QueryAllFormInstancesRequest) SetFormCode(v string) *QueryAllFormInstancesRequest {
	s.FormCode = &v
	return s
}

type QueryAllFormInstancesResponseBody struct {
	// 分页结果
	Result *QueryAllFormInstancesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryAllFormInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllFormInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllFormInstancesResponseBody) SetResult(v *QueryAllFormInstancesResponseBodyResult) *QueryAllFormInstancesResponseBody {
	s.Result = v
	return s
}

type QueryAllFormInstancesResponseBodyResult struct {
	// 下一页的游标
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 是否有更多数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 分页大小
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 表单列表
	Values []*QueryAllFormInstancesResponseBodyResultValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s QueryAllFormInstancesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryAllFormInstancesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryAllFormInstancesResponseBodyResult) SetNextToken(v string) *QueryAllFormInstancesResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResult) SetHasMore(v bool) *QueryAllFormInstancesResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResult) SetMaxResults(v int64) *QueryAllFormInstancesResponseBodyResult {
	s.MaxResults = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResult) SetValues(v []*QueryAllFormInstancesResponseBodyResultValues) *QueryAllFormInstancesResponseBodyResult {
	s.Values = v
	return s
}

type QueryAllFormInstancesResponseBodyResultValues struct {
	// 表单实例id
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// 应用搭建id
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// 表单模板code
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// 标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 创建人
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 修改人
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// 创建时间
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// 修改时间
	ModifyTimestamp *int64 `json:"modifyTimestamp,omitempty" xml:"modifyTimestamp,omitempty"`
	// 外部实例编码
	OutInstanceId *string `json:"outInstanceId,omitempty" xml:"outInstanceId,omitempty"`
	// 外部业务编码
	OutBizCode *string `json:"outBizCode,omitempty" xml:"outBizCode,omitempty"`
	// 扩展信息
	Attributes map[string]interface{} `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 表单实例数据
	FormInstDataList []*QueryAllFormInstancesResponseBodyResultValuesFormInstDataList `json:"formInstDataList,omitempty" xml:"formInstDataList,omitempty" type:"Repeated"`
}

func (s QueryAllFormInstancesResponseBodyResultValues) String() string {
	return tea.Prettify(s)
}

func (s QueryAllFormInstancesResponseBodyResultValues) GoString() string {
	return s.String()
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetFormInstanceId(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.FormInstanceId = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetAppUuid(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.AppUuid = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetFormCode(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.FormCode = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetTitle(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.Title = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetCreator(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.Creator = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetModifier(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.Modifier = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetCreateTimestamp(v int64) *QueryAllFormInstancesResponseBodyResultValues {
	s.CreateTimestamp = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetModifyTimestamp(v int64) *QueryAllFormInstancesResponseBodyResultValues {
	s.ModifyTimestamp = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetOutInstanceId(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.OutInstanceId = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetOutBizCode(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.OutBizCode = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetAttributes(v map[string]interface{}) *QueryAllFormInstancesResponseBodyResultValues {
	s.Attributes = v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetFormInstDataList(v []*QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) *QueryAllFormInstancesResponseBodyResultValues {
	s.FormInstDataList = v
	return s
}

type QueryAllFormInstancesResponseBodyResultValuesFormInstDataList struct {
	// 控件类型
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// 控件别名
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// 表单控件扩展数据
	ExtendValue *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	// 控件名称
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// 控件填写的数据
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 控件唯一id
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) GoString() string {
	return s.String()
}

func (s *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) SetComponentType(v string) *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList {
	s.ComponentType = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) SetBizAlias(v string) *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList {
	s.BizAlias = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) SetExtendValue(v string) *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList {
	s.ExtendValue = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) SetLabel(v string) *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList {
	s.Label = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) SetValue(v string) *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList {
	s.Value = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) SetKey(v string) *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList {
	s.Key = &v
	return s
}

type QueryAllFormInstancesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAllFormInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllFormInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllFormInstancesResponse) GoString() string {
	return s.String()
}

func (s *QueryAllFormInstancesResponse) SetHeaders(v map[string]*string) *QueryAllFormInstancesResponse {
	s.Headers = v
	return s
}

func (s *QueryAllFormInstancesResponse) SetBody(v *QueryAllFormInstancesResponseBody) *QueryAllFormInstancesResponse {
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

func (client *Client) QueryFormInstance(request *QueryFormInstanceRequest) (_result *QueryFormInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryFormInstanceHeaders{}
	_result = &QueryFormInstanceResponse{}
	_body, _err := client.QueryFormInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFormInstanceWithOptions(request *QueryFormInstanceRequest, headers *QueryFormInstanceHeaders, runtime *util.RuntimeOptions) (_result *QueryFormInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FormInstanceId)) {
		query["formInstanceId"] = request.FormInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.FormCode)) {
		query["formCode"] = request.FormCode
	}

	if !tea.BoolValue(util.IsUnset(request.AppUuid)) {
		query["appUuid"] = request.AppUuid
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
	_result = &QueryFormInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryFormInstance"), tea.String("workflow_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/workflow/forms/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFormByBizType(request *QueryFormByBizTypeRequest) (_result *QueryFormByBizTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryFormByBizTypeHeaders{}
	_result = &QueryFormByBizTypeResponse{}
	_body, _err := client.QueryFormByBizTypeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFormByBizTypeWithOptions(request *QueryFormByBizTypeRequest, headers *QueryFormByBizTypeHeaders, runtime *util.RuntimeOptions) (_result *QueryFormByBizTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUuid)) {
		body["appUuid"] = request.AppUuid
	}

	if !tea.BoolValue(util.IsUnset(request.BizTypes)) {
		body["bizTypes"] = request.BizTypes
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
	_result = &QueryFormByBizTypeResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryFormByBizType"), tea.String("workflow_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/workflow/forms/forminfos/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllProcessInstances(request *QueryAllProcessInstancesRequest) (_result *QueryAllProcessInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllProcessInstancesHeaders{}
	_result = &QueryAllProcessInstancesResponse{}
	_body, _err := client.QueryAllProcessInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllProcessInstancesWithOptions(request *QueryAllProcessInstancesRequest, headers *QueryAllProcessInstancesHeaders, runtime *util.RuntimeOptions) (_result *QueryAllProcessInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimeInMills)) {
		query["startTimeInMills"] = request.StartTimeInMills
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimeInMills)) {
		query["endTimeInMills"] = request.EndTimeInMills
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		query["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.AppUuid)) {
		query["appUuid"] = request.AppUuid
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
	_result = &QueryAllProcessInstancesResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAllProcessInstances"), tea.String("workflow_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/workflow/processes/pages/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllFormInstances(request *QueryAllFormInstancesRequest) (_result *QueryAllFormInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllFormInstancesHeaders{}
	_result = &QueryAllFormInstancesResponse{}
	_body, _err := client.QueryAllFormInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllFormInstancesWithOptions(request *QueryAllFormInstancesRequest, headers *QueryAllFormInstancesHeaders, runtime *util.RuntimeOptions) (_result *QueryAllFormInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.AppUuid)) {
		query["appUuid"] = request.AppUuid
	}

	if !tea.BoolValue(util.IsUnset(request.FormCode)) {
		query["formCode"] = request.FormCode
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
	_result = &QueryAllFormInstancesResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAllFormInstances"), tea.String("workflow_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/workflow/forms/pages/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
