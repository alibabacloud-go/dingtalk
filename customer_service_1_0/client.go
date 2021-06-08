// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package customer_service_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketHeaders) GoString() string {
	return s.String()
}

func (s *CreateTicketHeaders) SetCommonHeaders(v map[string]*string) *CreateTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTicketHeaders) SetXAcsDingtalkAccessToken(v string) *CreateTicketHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateTicketRequest struct {
	// 会员来源
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 第三方会员ID
	ForeignId *string `json:"foreignId,omitempty" xml:"foreignId,omitempty"`
	// 第三方会员名称
	ForeignName *string `json:"foreignName,omitempty" xml:"foreignName,omitempty"`
	// 实例ID
	OpenInstanceId *string `json:"openInstanceId,omitempty" xml:"openInstanceId,omitempty"`
	// 智能客服产品
	ProductionType *int32 `json:"productionType,omitempty" xml:"productionType,omitempty"`
	// 工单模板ID
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// 工单标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 工单表单
	Properties []*CreateTicketRequestProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Repeated"`
}

func (s CreateTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketRequest) GoString() string {
	return s.String()
}

func (s *CreateTicketRequest) SetSourceId(v string) *CreateTicketRequest {
	s.SourceId = &v
	return s
}

func (s *CreateTicketRequest) SetForeignId(v string) *CreateTicketRequest {
	s.ForeignId = &v
	return s
}

func (s *CreateTicketRequest) SetForeignName(v string) *CreateTicketRequest {
	s.ForeignName = &v
	return s
}

func (s *CreateTicketRequest) SetOpenInstanceId(v string) *CreateTicketRequest {
	s.OpenInstanceId = &v
	return s
}

func (s *CreateTicketRequest) SetProductionType(v int32) *CreateTicketRequest {
	s.ProductionType = &v
	return s
}

func (s *CreateTicketRequest) SetTemplateId(v string) *CreateTicketRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateTicketRequest) SetTitle(v string) *CreateTicketRequest {
	s.Title = &v
	return s
}

func (s *CreateTicketRequest) SetProperties(v []*CreateTicketRequestProperties) *CreateTicketRequest {
	s.Properties = v
	return s
}

type CreateTicketRequestProperties struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateTicketRequestProperties) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketRequestProperties) GoString() string {
	return s.String()
}

func (s *CreateTicketRequestProperties) SetName(v string) *CreateTicketRequestProperties {
	s.Name = &v
	return s
}

func (s *CreateTicketRequestProperties) SetValue(v string) *CreateTicketRequestProperties {
	s.Value = &v
	return s
}

type CreateTicketResponseBody struct {
	// 新创建工单ID
	TicketId *string `json:"ticketId,omitempty" xml:"ticketId,omitempty"`
}

func (s CreateTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTicketResponseBody) SetTicketId(v string) *CreateTicketResponseBody {
	s.TicketId = &v
	return s
}

type CreateTicketResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketResponse) GoString() string {
	return s.String()
}

func (s *CreateTicketResponse) SetHeaders(v map[string]*string) *CreateTicketResponse {
	s.Headers = v
	return s
}

func (s *CreateTicketResponse) SetBody(v *CreateTicketResponseBody) *CreateTicketResponse {
	s.Body = v
	return s
}

type PageListRobotHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PageListRobotHeaders) String() string {
	return tea.Prettify(s)
}

func (s PageListRobotHeaders) GoString() string {
	return s.String()
}

func (s *PageListRobotHeaders) SetCommonHeaders(v map[string]*string) *PageListRobotHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PageListRobotHeaders) SetXAcsDingtalkAccessToken(v string) *PageListRobotHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PageListRobotRequest struct {
	// 查询的企业Id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 多实例ID
	OpenInstanceId *string `json:"openInstanceId,omitempty" xml:"openInstanceId,omitempty"`
	// 产品类型
	ProductionType *int32 `json:"productionType,omitempty" xml:"productionType,omitempty"`
	// 用来标记当前开始读取的位置，置空表示从头开始
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 本次读取的最大数据记录数量
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
}

func (s PageListRobotRequest) String() string {
	return tea.Prettify(s)
}

func (s PageListRobotRequest) GoString() string {
	return s.String()
}

func (s *PageListRobotRequest) SetCorpId(v string) *PageListRobotRequest {
	s.CorpId = &v
	return s
}

func (s *PageListRobotRequest) SetOpenInstanceId(v string) *PageListRobotRequest {
	s.OpenInstanceId = &v
	return s
}

func (s *PageListRobotRequest) SetProductionType(v int32) *PageListRobotRequest {
	s.ProductionType = &v
	return s
}

func (s *PageListRobotRequest) SetNextToken(v int64) *PageListRobotRequest {
	s.NextToken = &v
	return s
}

func (s *PageListRobotRequest) SetMaxResults(v int32) *PageListRobotRequest {
	s.MaxResults = &v
	return s
}

type PageListRobotResponseBody struct {
	// 查询结果总数
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 下一次查询起始游标
	NextCursor *int64 `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	// 是否有更多结果
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 查询结果列表
	List []*PageListRobotResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s PageListRobotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageListRobotResponseBody) GoString() string {
	return s.String()
}

func (s *PageListRobotResponseBody) SetTotal(v int64) *PageListRobotResponseBody {
	s.Total = &v
	return s
}

func (s *PageListRobotResponseBody) SetNextCursor(v int64) *PageListRobotResponseBody {
	s.NextCursor = &v
	return s
}

func (s *PageListRobotResponseBody) SetHasMore(v bool) *PageListRobotResponseBody {
	s.HasMore = &v
	return s
}

func (s *PageListRobotResponseBody) SetList(v []*PageListRobotResponseBodyList) *PageListRobotResponseBody {
	s.List = v
	return s
}

type PageListRobotResponseBodyList struct {
	// 机器人自增Id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 机器人名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 机器人APPKEY
	AppKey *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	// 机器人所在租户ID
	AccountId *int64 `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// 机器人状态
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PageListRobotResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s PageListRobotResponseBodyList) GoString() string {
	return s.String()
}

func (s *PageListRobotResponseBodyList) SetId(v int64) *PageListRobotResponseBodyList {
	s.Id = &v
	return s
}

func (s *PageListRobotResponseBodyList) SetName(v string) *PageListRobotResponseBodyList {
	s.Name = &v
	return s
}

func (s *PageListRobotResponseBodyList) SetAppKey(v string) *PageListRobotResponseBodyList {
	s.AppKey = &v
	return s
}

func (s *PageListRobotResponseBodyList) SetAccountId(v int64) *PageListRobotResponseBodyList {
	s.AccountId = &v
	return s
}

func (s *PageListRobotResponseBodyList) SetStatus(v int32) *PageListRobotResponseBodyList {
	s.Status = &v
	return s
}

type PageListRobotResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PageListRobotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PageListRobotResponse) String() string {
	return tea.Prettify(s)
}

func (s PageListRobotResponse) GoString() string {
	return s.String()
}

func (s *PageListRobotResponse) SetHeaders(v map[string]*string) *PageListRobotResponse {
	s.Headers = v
	return s
}

func (s *PageListRobotResponse) SetBody(v *PageListRobotResponseBody) *PageListRobotResponse {
	s.Body = v
	return s
}

type PageListActionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PageListActionHeaders) String() string {
	return tea.Prettify(s)
}

func (s PageListActionHeaders) GoString() string {
	return s.String()
}

func (s *PageListActionHeaders) SetCommonHeaders(v map[string]*string) *PageListActionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PageListActionHeaders) SetXAcsDingtalkAccessToken(v string) *PageListActionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PageListActionRequest struct {
	// 实例id
	OpenInstanceId *string `json:"openInstanceId,omitempty" xml:"openInstanceId,omitempty"`
	// 产品类型
	ProductionType *int64 `json:"productionType,omitempty" xml:"productionType,omitempty"`
	// 用来标记当前开始读取的位置，置空表示从头开始。
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 本次读取的最大数据记录数量，此参数为可选参数，用户传入为空时，应该有默认值。应设置最大值限制，最大不超过100
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
}

func (s PageListActionRequest) String() string {
	return tea.Prettify(s)
}

func (s PageListActionRequest) GoString() string {
	return s.String()
}

func (s *PageListActionRequest) SetOpenInstanceId(v string) *PageListActionRequest {
	s.OpenInstanceId = &v
	return s
}

func (s *PageListActionRequest) SetProductionType(v int64) *PageListActionRequest {
	s.ProductionType = &v
	return s
}

func (s *PageListActionRequest) SetNextToken(v string) *PageListActionRequest {
	s.NextToken = &v
	return s
}

func (s *PageListActionRequest) SetMaxResults(v int64) *PageListActionRequest {
	s.MaxResults = &v
	return s
}

type PageListActionResponseBody struct {
	// nextCursor
	NextCursor *int64 `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	// total
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
	// list
	List []*PageListActionResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s PageListActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageListActionResponseBody) GoString() string {
	return s.String()
}

func (s *PageListActionResponseBody) SetNextCursor(v int64) *PageListActionResponseBody {
	s.NextCursor = &v
	return s
}

func (s *PageListActionResponseBody) SetTotal(v int64) *PageListActionResponseBody {
	s.Total = &v
	return s
}

func (s *PageListActionResponseBody) SetList(v []*PageListActionResponseBodyList) *PageListActionResponseBody {
	s.List = v
	return s
}

type PageListActionResponseBodyList struct {
	// operatorId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// operator
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// operatorRole
	OperatorRole *string `json:"operatorRole,omitempty" xml:"operatorRole,omitempty"`
	// actionCode
	ActionCode *string `json:"actionCode,omitempty" xml:"actionCode,omitempty"`
	// actionContent
	ActionContent []*PageListActionResponseBodyListActionContent `json:"actionContent,omitempty" xml:"actionContent,omitempty" type:"Repeated"`
}

func (s PageListActionResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s PageListActionResponseBodyList) GoString() string {
	return s.String()
}

func (s *PageListActionResponseBodyList) SetOperatorId(v string) *PageListActionResponseBodyList {
	s.OperatorId = &v
	return s
}

func (s *PageListActionResponseBodyList) SetOperator(v string) *PageListActionResponseBodyList {
	s.Operator = &v
	return s
}

func (s *PageListActionResponseBodyList) SetOperatorRole(v string) *PageListActionResponseBodyList {
	s.OperatorRole = &v
	return s
}

func (s *PageListActionResponseBodyList) SetActionCode(v string) *PageListActionResponseBodyList {
	s.ActionCode = &v
	return s
}

func (s *PageListActionResponseBodyList) SetActionContent(v []*PageListActionResponseBodyListActionContent) *PageListActionResponseBodyList {
	s.ActionContent = v
	return s
}

type PageListActionResponseBodyListActionContent struct {
	// displayValue
	DisplayValue *string `json:"displayValue,omitempty" xml:"displayValue,omitempty"`
	// displayName
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// valueType
	ValueType *string `json:"valueType,omitempty" xml:"valueType,omitempty"`
}

func (s PageListActionResponseBodyListActionContent) String() string {
	return tea.Prettify(s)
}

func (s PageListActionResponseBodyListActionContent) GoString() string {
	return s.String()
}

func (s *PageListActionResponseBodyListActionContent) SetDisplayValue(v string) *PageListActionResponseBodyListActionContent {
	s.DisplayValue = &v
	return s
}

func (s *PageListActionResponseBodyListActionContent) SetDisplayName(v string) *PageListActionResponseBodyListActionContent {
	s.DisplayName = &v
	return s
}

func (s *PageListActionResponseBodyListActionContent) SetName(v string) *PageListActionResponseBodyListActionContent {
	s.Name = &v
	return s
}

func (s *PageListActionResponseBodyListActionContent) SetValue(v string) *PageListActionResponseBodyListActionContent {
	s.Value = &v
	return s
}

func (s *PageListActionResponseBodyListActionContent) SetValueType(v string) *PageListActionResponseBodyListActionContent {
	s.ValueType = &v
	return s
}

type PageListActionResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PageListActionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PageListActionResponse) String() string {
	return tea.Prettify(s)
}

func (s PageListActionResponse) GoString() string {
	return s.String()
}

func (s *PageListActionResponse) SetHeaders(v map[string]*string) *PageListActionResponse {
	s.Headers = v
	return s
}

func (s *PageListActionResponse) SetBody(v *PageListActionResponseBody) *PageListActionResponse {
	s.Body = v
	return s
}

type ExecuteActivityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ExecuteActivityHeaders) String() string {
	return tea.Prettify(s)
}

func (s ExecuteActivityHeaders) GoString() string {
	return s.String()
}

func (s *ExecuteActivityHeaders) SetCommonHeaders(v map[string]*string) *ExecuteActivityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ExecuteActivityHeaders) SetXAcsDingtalkAccessToken(v string) *ExecuteActivityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ExecuteActivityRequest struct {
	// 来源ID
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 会员ID
	ForeignId *string `json:"foreignId,omitempty" xml:"foreignId,omitempty"`
	// 会员名称
	ForeignName *string `json:"foreignName,omitempty" xml:"foreignName,omitempty"`
	// 动作编码
	ActivityCode *string `json:"activityCode,omitempty" xml:"activityCode,omitempty"`
	// 实例id
	OpenInstanceId *string `json:"openInstanceId,omitempty" xml:"openInstanceId,omitempty"`
	// 产品类型
	ProductionType *int32 `json:"productionType,omitempty" xml:"productionType,omitempty"`
	// 工单表单字段
	Properties []*ExecuteActivityRequestProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Repeated"`
}

func (s ExecuteActivityRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteActivityRequest) GoString() string {
	return s.String()
}

func (s *ExecuteActivityRequest) SetSourceId(v string) *ExecuteActivityRequest {
	s.SourceId = &v
	return s
}

func (s *ExecuteActivityRequest) SetForeignId(v string) *ExecuteActivityRequest {
	s.ForeignId = &v
	return s
}

func (s *ExecuteActivityRequest) SetForeignName(v string) *ExecuteActivityRequest {
	s.ForeignName = &v
	return s
}

func (s *ExecuteActivityRequest) SetActivityCode(v string) *ExecuteActivityRequest {
	s.ActivityCode = &v
	return s
}

func (s *ExecuteActivityRequest) SetOpenInstanceId(v string) *ExecuteActivityRequest {
	s.OpenInstanceId = &v
	return s
}

func (s *ExecuteActivityRequest) SetProductionType(v int32) *ExecuteActivityRequest {
	s.ProductionType = &v
	return s
}

func (s *ExecuteActivityRequest) SetProperties(v []*ExecuteActivityRequestProperties) *ExecuteActivityRequest {
	s.Properties = v
	return s
}

type ExecuteActivityRequestProperties struct {
	// 字段的key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 字段的值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ExecuteActivityRequestProperties) String() string {
	return tea.Prettify(s)
}

func (s ExecuteActivityRequestProperties) GoString() string {
	return s.String()
}

func (s *ExecuteActivityRequestProperties) SetName(v string) *ExecuteActivityRequestProperties {
	s.Name = &v
	return s
}

func (s *ExecuteActivityRequestProperties) SetValue(v string) *ExecuteActivityRequestProperties {
	s.Value = &v
	return s
}

type ExecuteActivityResponseBody struct {
	// 任务id
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ExecuteActivityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteActivityResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteActivityResponseBody) SetTaskId(v string) *ExecuteActivityResponseBody {
	s.TaskId = &v
	return s
}

type ExecuteActivityResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecuteActivityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteActivityResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteActivityResponse) GoString() string {
	return s.String()
}

func (s *ExecuteActivityResponse) SetHeaders(v map[string]*string) *ExecuteActivityResponse {
	s.Headers = v
	return s
}

func (s *ExecuteActivityResponse) SetBody(v *ExecuteActivityResponseBody) *ExecuteActivityResponse {
	s.Body = v
	return s
}

type PageListTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PageListTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s PageListTicketHeaders) GoString() string {
	return s.String()
}

func (s *PageListTicketHeaders) SetCommonHeaders(v map[string]*string) *PageListTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PageListTicketHeaders) SetXAcsDingtalkAccessToken(v string) *PageListTicketHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PageListTicketRequest struct {
	// 实例id
	OpenInstanceId *string `json:"openInstanceId,omitempty" xml:"openInstanceId,omitempty"`
	// 产品类型
	ProductionType *int32 `json:"productionType,omitempty" xml:"productionType,omitempty"`
	// 工单模板
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// 工单ID
	TicketId *string `json:"ticketId,omitempty" xml:"ticketId,omitempty"`
	// 来源
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 第三方用户id
	ForeignId *string `json:"foreignId,omitempty" xml:"foreignId,omitempty"`
	// 工单状态
	TicketStatus *string `json:"ticketStatus,omitempty" xml:"ticketStatus,omitempty"`
	// 开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 结束时间
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 用来标记当前开始读取的位置，置空表示从头开始
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 本次读取的最大数据记录数量
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
}

func (s PageListTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s PageListTicketRequest) GoString() string {
	return s.String()
}

func (s *PageListTicketRequest) SetOpenInstanceId(v string) *PageListTicketRequest {
	s.OpenInstanceId = &v
	return s
}

func (s *PageListTicketRequest) SetProductionType(v int32) *PageListTicketRequest {
	s.ProductionType = &v
	return s
}

func (s *PageListTicketRequest) SetTemplateId(v string) *PageListTicketRequest {
	s.TemplateId = &v
	return s
}

func (s *PageListTicketRequest) SetTicketId(v string) *PageListTicketRequest {
	s.TicketId = &v
	return s
}

func (s *PageListTicketRequest) SetSourceId(v string) *PageListTicketRequest {
	s.SourceId = &v
	return s
}

func (s *PageListTicketRequest) SetForeignId(v string) *PageListTicketRequest {
	s.ForeignId = &v
	return s
}

func (s *PageListTicketRequest) SetTicketStatus(v string) *PageListTicketRequest {
	s.TicketStatus = &v
	return s
}

func (s *PageListTicketRequest) SetStartTime(v int64) *PageListTicketRequest {
	s.StartTime = &v
	return s
}

func (s *PageListTicketRequest) SetEndTime(v int64) *PageListTicketRequest {
	s.EndTime = &v
	return s
}

func (s *PageListTicketRequest) SetNextToken(v string) *PageListTicketRequest {
	s.NextToken = &v
	return s
}

func (s *PageListTicketRequest) SetMaxResults(v int32) *PageListTicketRequest {
	s.MaxResults = &v
	return s
}

type PageListTicketResponseBody struct {
	// nextCursor
	NextCursor *int64 `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	// total
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
	// list
	List []*PageListTicketResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s PageListTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageListTicketResponseBody) GoString() string {
	return s.String()
}

func (s *PageListTicketResponseBody) SetNextCursor(v int64) *PageListTicketResponseBody {
	s.NextCursor = &v
	return s
}

func (s *PageListTicketResponseBody) SetTotal(v int64) *PageListTicketResponseBody {
	s.Total = &v
	return s
}

func (s *PageListTicketResponseBody) SetList(v []*PageListTicketResponseBodyList) *PageListTicketResponseBody {
	s.List = v
	return s
}

type PageListTicketResponseBodyList struct {
	// foreignId
	ForeignId *string `json:"foreignId,omitempty" xml:"foreignId,omitempty"`
	// sourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// foreignName
	ForeignName *string `json:"foreignName,omitempty" xml:"foreignName,omitempty"`
	// templateId
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// title
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// ticketId
	TicketId *string `json:"ticketId,omitempty" xml:"ticketId,omitempty"`
	// ticketStatus
	TicketStatus *string `json:"ticketStatus,omitempty" xml:"ticketStatus,omitempty"`
	// openInstanceId
	OpenInstanceId *string `json:"openInstanceId,omitempty" xml:"openInstanceId,omitempty"`
	// productionType
	ProductionType *int32 `json:"productionType,omitempty" xml:"productionType,omitempty"`
	// gmtCreate
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// gmtModified
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// bizDataMap
	BizDataMap map[string]interface{} `json:"bizDataMap,omitempty" xml:"bizDataMap,omitempty"`
}

func (s PageListTicketResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s PageListTicketResponseBodyList) GoString() string {
	return s.String()
}

func (s *PageListTicketResponseBodyList) SetForeignId(v string) *PageListTicketResponseBodyList {
	s.ForeignId = &v
	return s
}

func (s *PageListTicketResponseBodyList) SetSourceId(v string) *PageListTicketResponseBodyList {
	s.SourceId = &v
	return s
}

func (s *PageListTicketResponseBodyList) SetForeignName(v string) *PageListTicketResponseBodyList {
	s.ForeignName = &v
	return s
}

func (s *PageListTicketResponseBodyList) SetTemplateId(v string) *PageListTicketResponseBodyList {
	s.TemplateId = &v
	return s
}

func (s *PageListTicketResponseBodyList) SetTitle(v string) *PageListTicketResponseBodyList {
	s.Title = &v
	return s
}

func (s *PageListTicketResponseBodyList) SetTicketId(v string) *PageListTicketResponseBodyList {
	s.TicketId = &v
	return s
}

func (s *PageListTicketResponseBodyList) SetTicketStatus(v string) *PageListTicketResponseBodyList {
	s.TicketStatus = &v
	return s
}

func (s *PageListTicketResponseBodyList) SetOpenInstanceId(v string) *PageListTicketResponseBodyList {
	s.OpenInstanceId = &v
	return s
}

func (s *PageListTicketResponseBodyList) SetProductionType(v int32) *PageListTicketResponseBodyList {
	s.ProductionType = &v
	return s
}

func (s *PageListTicketResponseBodyList) SetGmtCreate(v string) *PageListTicketResponseBodyList {
	s.GmtCreate = &v
	return s
}

func (s *PageListTicketResponseBodyList) SetGmtModified(v string) *PageListTicketResponseBodyList {
	s.GmtModified = &v
	return s
}

func (s *PageListTicketResponseBodyList) SetBizDataMap(v map[string]interface{}) *PageListTicketResponseBodyList {
	s.BizDataMap = v
	return s
}

type PageListTicketResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PageListTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PageListTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s PageListTicketResponse) GoString() string {
	return s.String()
}

func (s *PageListTicketResponse) SetHeaders(v map[string]*string) *PageListTicketResponse {
	s.Headers = v
	return s
}

func (s *PageListTicketResponse) SetBody(v *PageListTicketResponseBody) *PageListTicketResponse {
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

func (client *Client) CreateTicket(request *CreateTicketRequest) (_result *CreateTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTicketHeaders{}
	_result = &CreateTicketResponse{}
	_body, _err := client.CreateTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTicketWithOptions(request *CreateTicketRequest, headers *CreateTicketHeaders, runtime *util.RuntimeOptions) (_result *CreateTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		body["sourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ForeignId)) {
		body["foreignId"] = request.ForeignId
	}

	if !tea.BoolValue(util.IsUnset(request.ForeignName)) {
		body["foreignName"] = request.ForeignName
	}

	if !tea.BoolValue(util.IsUnset(request.OpenInstanceId)) {
		body["openInstanceId"] = request.OpenInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductionType)) {
		body["productionType"] = request.ProductionType
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Properties)) {
		body["properties"] = request.Properties
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
	_result = &CreateTicketResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateTicket"), tea.String("customerService_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/customerService/tickets"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PageListRobot(request *PageListRobotRequest) (_result *PageListRobotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PageListRobotHeaders{}
	_result = &PageListRobotResponse{}
	_body, _err := client.PageListRobotWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PageListRobotWithOptions(request *PageListRobotRequest, headers *PageListRobotHeaders, runtime *util.RuntimeOptions) (_result *PageListRobotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenInstanceId)) {
		query["openInstanceId"] = request.OpenInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductionType)) {
		query["productionType"] = request.ProductionType
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
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
	_result = &PageListRobotResponse{}
	_body, _err := client.DoROARequest(tea.String("PageListRobot"), tea.String("customerService_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/customerService/robots"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PageListAction(ticketId *string, request *PageListActionRequest) (_result *PageListActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PageListActionHeaders{}
	_result = &PageListActionResponse{}
	_body, _err := client.PageListActionWithOptions(ticketId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PageListActionWithOptions(ticketId *string, request *PageListActionRequest, headers *PageListActionHeaders, runtime *util.RuntimeOptions) (_result *PageListActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenInstanceId)) {
		query["openInstanceId"] = request.OpenInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductionType)) {
		query["productionType"] = request.ProductionType
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
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
	_result = &PageListActionResponse{}
	_body, _err := client.DoROARequest(tea.String("PageListAction"), tea.String("customerService_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/customerService/tickets/"+tea.StringValue(ticketId)+"/actions"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteActivity(ticketId *string, request *ExecuteActivityRequest) (_result *ExecuteActivityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ExecuteActivityHeaders{}
	_result = &ExecuteActivityResponse{}
	_body, _err := client.ExecuteActivityWithOptions(ticketId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteActivityWithOptions(ticketId *string, request *ExecuteActivityRequest, headers *ExecuteActivityHeaders, runtime *util.RuntimeOptions) (_result *ExecuteActivityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		body["sourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ForeignId)) {
		body["foreignId"] = request.ForeignId
	}

	if !tea.BoolValue(util.IsUnset(request.ForeignName)) {
		body["foreignName"] = request.ForeignName
	}

	if !tea.BoolValue(util.IsUnset(request.ActivityCode)) {
		body["activityCode"] = request.ActivityCode
	}

	if !tea.BoolValue(util.IsUnset(request.OpenInstanceId)) {
		body["openInstanceId"] = request.OpenInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductionType)) {
		body["productionType"] = request.ProductionType
	}

	if !tea.BoolValue(util.IsUnset(request.Properties)) {
		body["properties"] = request.Properties
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
	_result = &ExecuteActivityResponse{}
	_body, _err := client.DoROARequest(tea.String("ExecuteActivity"), tea.String("customerService_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/customerService/tickets/"+tea.StringValue(ticketId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PageListTicket(request *PageListTicketRequest) (_result *PageListTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PageListTicketHeaders{}
	_result = &PageListTicketResponse{}
	_body, _err := client.PageListTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PageListTicketWithOptions(request *PageListTicketRequest, headers *PageListTicketHeaders, runtime *util.RuntimeOptions) (_result *PageListTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenInstanceId)) {
		query["openInstanceId"] = request.OpenInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductionType)) {
		query["productionType"] = request.ProductionType
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["templateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TicketId)) {
		query["ticketId"] = request.TicketId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		query["sourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ForeignId)) {
		query["foreignId"] = request.ForeignId
	}

	if !tea.BoolValue(util.IsUnset(request.TicketStatus)) {
		query["ticketStatus"] = request.TicketStatus
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
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
	_result = &PageListTicketResponse{}
	_body, _err := client.DoROARequest(tea.String("PageListTicket"), tea.String("customerService_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/customerService/tickets"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
