// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package todo_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetTodoTypeConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTodoTypeConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTypeConfigHeaders) GoString() string {
	return s.String()
}

func (s *GetTodoTypeConfigHeaders) SetCommonHeaders(v map[string]*string) *GetTodoTypeConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTodoTypeConfigHeaders) SetXAcsDingtalkAccessToken(v string) *GetTodoTypeConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTodoTypeConfigResponseBody struct {
	// id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 创建时间
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 更新时间
	ModifiedTime *int64 `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// 创建者（用户的unionId）
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 更新者（用户的unionId）
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// 接入应用标识
	BizTag *string `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 卡片类型（取值为：1-标准卡片，2-自定义卡片）
	CardType *int32 `json:"cardType,omitempty" xml:"cardType,omitempty"`
	// 卡片类型icon，用于在待办列表展示
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 待办卡片类型描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 详情页链接在PC端的打开方式，取值为：「PC_SLIDE」-PC端侧边栏打开，「PC_BROWSER」-浏览器打开
	PcDetailUrlOpenMode *string `json:"pcDetailUrlOpenMode,omitempty" xml:"pcDetailUrlOpenMode,omitempty"`
	// 待办卡片内容区表单自定义字段配置
	ContentFieldList []*GetTodoTypeConfigResponseBodyContentFieldList `json:"contentFieldList,omitempty" xml:"contentFieldList,omitempty" type:"Repeated"`
	// 待办卡片操作区按钮配置
	ActionList []*GetTodoTypeConfigResponseBodyActionList `json:"actionList,omitempty" xml:"actionList,omitempty" type:"Repeated"`
}

func (s GetTodoTypeConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTypeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetTodoTypeConfigResponseBody) SetId(v string) *GetTodoTypeConfigResponseBody {
	s.Id = &v
	return s
}

func (s *GetTodoTypeConfigResponseBody) SetCreatedTime(v int64) *GetTodoTypeConfigResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetTodoTypeConfigResponseBody) SetModifiedTime(v int64) *GetTodoTypeConfigResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetTodoTypeConfigResponseBody) SetCreatorId(v string) *GetTodoTypeConfigResponseBody {
	s.CreatorId = &v
	return s
}

func (s *GetTodoTypeConfigResponseBody) SetModifierId(v string) *GetTodoTypeConfigResponseBody {
	s.ModifierId = &v
	return s
}

func (s *GetTodoTypeConfigResponseBody) SetBizTag(v string) *GetTodoTypeConfigResponseBody {
	s.BizTag = &v
	return s
}

func (s *GetTodoTypeConfigResponseBody) SetRequestId(v string) *GetTodoTypeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTodoTypeConfigResponseBody) SetCardType(v int32) *GetTodoTypeConfigResponseBody {
	s.CardType = &v
	return s
}

func (s *GetTodoTypeConfigResponseBody) SetIcon(v string) *GetTodoTypeConfigResponseBody {
	s.Icon = &v
	return s
}

func (s *GetTodoTypeConfigResponseBody) SetDescription(v string) *GetTodoTypeConfigResponseBody {
	s.Description = &v
	return s
}

func (s *GetTodoTypeConfigResponseBody) SetPcDetailUrlOpenMode(v string) *GetTodoTypeConfigResponseBody {
	s.PcDetailUrlOpenMode = &v
	return s
}

func (s *GetTodoTypeConfigResponseBody) SetContentFieldList(v []*GetTodoTypeConfigResponseBodyContentFieldList) *GetTodoTypeConfigResponseBody {
	s.ContentFieldList = v
	return s
}

func (s *GetTodoTypeConfigResponseBody) SetActionList(v []*GetTodoTypeConfigResponseBodyActionList) *GetTodoTypeConfigResponseBody {
	s.ActionList = v
	return s
}

type GetTodoTypeConfigResponseBodyContentFieldList struct {
	// 字段唯一标识
	FieldKey *string `json:"fieldKey,omitempty" xml:"fieldKey,omitempty"`
	// 字段类型（取值为：text-文本，url-链接）
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	// 字段的显示名称（支持国际化）
	NameI18n map[string]interface{} `json:"nameI18n,omitempty" xml:"nameI18n,omitempty"`
}

func (s GetTodoTypeConfigResponseBodyContentFieldList) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTypeConfigResponseBodyContentFieldList) GoString() string {
	return s.String()
}

func (s *GetTodoTypeConfigResponseBodyContentFieldList) SetFieldKey(v string) *GetTodoTypeConfigResponseBodyContentFieldList {
	s.FieldKey = &v
	return s
}

func (s *GetTodoTypeConfigResponseBodyContentFieldList) SetFieldType(v string) *GetTodoTypeConfigResponseBodyContentFieldList {
	s.FieldType = &v
	return s
}

func (s *GetTodoTypeConfigResponseBodyContentFieldList) SetNameI18n(v map[string]interface{}) *GetTodoTypeConfigResponseBodyContentFieldList {
	s.NameI18n = v
	return s
}

type GetTodoTypeConfigResponseBodyActionList struct {
	// 操作按钮的唯一标识
	ActionKey *string `json:"actionKey,omitempty" xml:"actionKey,omitempty"`
	// 按钮样式类型（101：蓝色线型主按钮样式，例如「同意」，102：黑色线型副按钮样式，例如「拒绝」）
	ButtonStyleType *int32 `json:"buttonStyleType,omitempty" xml:"buttonStyleType,omitempty"`
	// 按钮类型（1：有操作的，2：直接跳转）
	ActionType *int32 `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// 按钮类型为直接跳转时，对应的跳转url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 按钮操作的显示名称（支持国际化）
	NameI18n map[string]interface{} `json:"nameI18n,omitempty" xml:"nameI18n,omitempty"`
}

func (s GetTodoTypeConfigResponseBodyActionList) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTypeConfigResponseBodyActionList) GoString() string {
	return s.String()
}

func (s *GetTodoTypeConfigResponseBodyActionList) SetActionKey(v string) *GetTodoTypeConfigResponseBodyActionList {
	s.ActionKey = &v
	return s
}

func (s *GetTodoTypeConfigResponseBodyActionList) SetButtonStyleType(v int32) *GetTodoTypeConfigResponseBodyActionList {
	s.ButtonStyleType = &v
	return s
}

func (s *GetTodoTypeConfigResponseBodyActionList) SetActionType(v int32) *GetTodoTypeConfigResponseBodyActionList {
	s.ActionType = &v
	return s
}

func (s *GetTodoTypeConfigResponseBodyActionList) SetUrl(v string) *GetTodoTypeConfigResponseBodyActionList {
	s.Url = &v
	return s
}

func (s *GetTodoTypeConfigResponseBodyActionList) SetNameI18n(v map[string]interface{}) *GetTodoTypeConfigResponseBodyActionList {
	s.NameI18n = v
	return s
}

type GetTodoTypeConfigResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTodoTypeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTodoTypeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTypeConfigResponse) GoString() string {
	return s.String()
}

func (s *GetTodoTypeConfigResponse) SetHeaders(v map[string]*string) *GetTodoTypeConfigResponse {
	s.Headers = v
	return s
}

func (s *GetTodoTypeConfigResponse) SetBody(v *GetTodoTypeConfigResponseBody) *GetTodoTypeConfigResponse {
	s.Body = v
	return s
}

type QueryTodoTasksHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryTodoTasksHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoTasksHeaders) GoString() string {
	return s.String()
}

func (s *QueryTodoTasksHeaders) SetCommonHeaders(v map[string]*string) *QueryTodoTasksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryTodoTasksHeaders) SetXAcsDingtalkAccessToken(v string) *QueryTodoTasksHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryTodoTasksRequest struct {
	// 分页游标。如果一个查询条件一次无法全部返回结果，会返回分页token，下次查询带上该token后会返回后续数据，直到分页token为null表示数据已经全部查询完毕。
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 排序字段。枚举值 默认为截止时间 dueTime。created | modified | finished | startTime | dueTime 创建时间 | 更新时间 | 完成时间 | 开始时间 | 截止时间
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// 排序方向。枚举值asc | desc 默认 asc
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	// 待办完成状态。
	IsDone *bool `json:"isDone,omitempty" xml:"isDone,omitempty"`
	// 查询目标用户角色类型，执行人 | 创建人 | 参与人，可以同时传多个值。如：[["executor"], ["creator"],["participant"]] 或 [["executor", "creator"]]
	RoleTypes [][]*string `json:"roleTypes,omitempty" xml:"roleTypes,omitempty" type:"Repeated"`
	// 查询从计划完成时间开始
	FromDueTime *int64 `json:"fromDueTime,omitempty" xml:"fromDueTime,omitempty"`
	// 查询到计划完成时间结束
	ToDueTime *int64 `json:"toDueTime,omitempty" xml:"toDueTime,omitempty"`
	// 所属分类
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 待办回收状态
	IsRecycled *bool `json:"isRecycled,omitempty" xml:"isRecycled,omitempty"`
}

func (s QueryTodoTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoTasksRequest) GoString() string {
	return s.String()
}

func (s *QueryTodoTasksRequest) SetNextToken(v string) *QueryTodoTasksRequest {
	s.NextToken = &v
	return s
}

func (s *QueryTodoTasksRequest) SetOrderBy(v string) *QueryTodoTasksRequest {
	s.OrderBy = &v
	return s
}

func (s *QueryTodoTasksRequest) SetOrderDirection(v string) *QueryTodoTasksRequest {
	s.OrderDirection = &v
	return s
}

func (s *QueryTodoTasksRequest) SetIsDone(v bool) *QueryTodoTasksRequest {
	s.IsDone = &v
	return s
}

func (s *QueryTodoTasksRequest) SetRoleTypes(v [][]*string) *QueryTodoTasksRequest {
	s.RoleTypes = v
	return s
}

func (s *QueryTodoTasksRequest) SetFromDueTime(v int64) *QueryTodoTasksRequest {
	s.FromDueTime = &v
	return s
}

func (s *QueryTodoTasksRequest) SetToDueTime(v int64) *QueryTodoTasksRequest {
	s.ToDueTime = &v
	return s
}

func (s *QueryTodoTasksRequest) SetCategory(v string) *QueryTodoTasksRequest {
	s.Category = &v
	return s
}

func (s *QueryTodoTasksRequest) SetIsRecycled(v bool) *QueryTodoTasksRequest {
	s.IsRecycled = &v
	return s
}

type QueryTodoTasksResponseBody struct {
	// 翻页token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 待办卡片列表
	TodoCards []*QueryTodoTasksResponseBodyTodoCards `json:"todoCards,omitempty" xml:"todoCards,omitempty" type:"Repeated"`
	// 数据总量
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryTodoTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoTasksResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTodoTasksResponseBody) SetNextToken(v string) *QueryTodoTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryTodoTasksResponseBody) SetTodoCards(v []*QueryTodoTasksResponseBodyTodoCards) *QueryTodoTasksResponseBody {
	s.TodoCards = v
	return s
}

func (s *QueryTodoTasksResponseBody) SetTotalCount(v int32) *QueryTodoTasksResponseBody {
	s.TotalCount = &v
	return s
}

type QueryTodoTasksResponseBodyTodoCards struct {
	// 待办id
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// 待办标题
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// 待办截止时间
	DueTime *int64 `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	// 详情页链接
	DetailUrl *QueryTodoTasksResponseBodyTodoCardsDetailUrl `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	// 待办卡片视图模型
	TodoCardView *QueryTodoTasksResponseBodyTodoCardsTodoCardView `json:"todoCardView,omitempty" xml:"todoCardView,omitempty" type:"Struct"`
	// 优先级
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// 创建时间
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 更新时间
	ModifiedTime *int64 `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// 待办状态
	TodoStatus *string `json:"todoStatus,omitempty" xml:"todoStatus,omitempty"`
	// 创建者id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 来源id
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 所属分类
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 所属应用
	BizTag *string `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	// 业务来源信息
	OriginalSource *QueryTodoTasksResponseBodyTodoCardsOriginalSource `json:"originalSource,omitempty" xml:"originalSource,omitempty" type:"Struct"`
	// 待办完成状态
	IsDone *bool `json:"isDone,omitempty" xml:"isDone,omitempty"`
	// 所属组织信息
	OrgInfo *QueryTodoTasksResponseBodyTodoCardsOrgInfo `json:"orgInfo,omitempty" xml:"orgInfo,omitempty" type:"Struct"`
}

func (s QueryTodoTasksResponseBodyTodoCards) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoTasksResponseBodyTodoCards) GoString() string {
	return s.String()
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetTaskId(v string) *QueryTodoTasksResponseBodyTodoCards {
	s.TaskId = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetSubject(v string) *QueryTodoTasksResponseBodyTodoCards {
	s.Subject = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetDueTime(v int64) *QueryTodoTasksResponseBodyTodoCards {
	s.DueTime = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetDetailUrl(v *QueryTodoTasksResponseBodyTodoCardsDetailUrl) *QueryTodoTasksResponseBodyTodoCards {
	s.DetailUrl = v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetTodoCardView(v *QueryTodoTasksResponseBodyTodoCardsTodoCardView) *QueryTodoTasksResponseBodyTodoCards {
	s.TodoCardView = v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetPriority(v int32) *QueryTodoTasksResponseBodyTodoCards {
	s.Priority = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetCreatedTime(v int64) *QueryTodoTasksResponseBodyTodoCards {
	s.CreatedTime = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetModifiedTime(v int64) *QueryTodoTasksResponseBodyTodoCards {
	s.ModifiedTime = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetTodoStatus(v string) *QueryTodoTasksResponseBodyTodoCards {
	s.TodoStatus = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetCreatorId(v string) *QueryTodoTasksResponseBodyTodoCards {
	s.CreatorId = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetSourceId(v string) *QueryTodoTasksResponseBodyTodoCards {
	s.SourceId = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetCategory(v string) *QueryTodoTasksResponseBodyTodoCards {
	s.Category = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetBizTag(v string) *QueryTodoTasksResponseBodyTodoCards {
	s.BizTag = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetOriginalSource(v *QueryTodoTasksResponseBodyTodoCardsOriginalSource) *QueryTodoTasksResponseBodyTodoCards {
	s.OriginalSource = v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetIsDone(v bool) *QueryTodoTasksResponseBodyTodoCards {
	s.IsDone = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCards) SetOrgInfo(v *QueryTodoTasksResponseBodyTodoCardsOrgInfo) *QueryTodoTasksResponseBodyTodoCards {
	s.OrgInfo = v
	return s
}

type QueryTodoTasksResponseBodyTodoCardsDetailUrl struct {
	// 移动端url地址
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
	// pc端url地址
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
}

func (s QueryTodoTasksResponseBodyTodoCardsDetailUrl) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoTasksResponseBodyTodoCardsDetailUrl) GoString() string {
	return s.String()
}

func (s *QueryTodoTasksResponseBodyTodoCardsDetailUrl) SetAppUrl(v string) *QueryTodoTasksResponseBodyTodoCardsDetailUrl {
	s.AppUrl = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCardsDetailUrl) SetPcUrl(v string) *QueryTodoTasksResponseBodyTodoCardsDetailUrl {
	s.PcUrl = &v
	return s
}

type QueryTodoTasksResponseBodyTodoCardsTodoCardView struct {
	// 卡片类型
	CardType *string `json:"cardType,omitempty" xml:"cardType,omitempty"`
	// 卡片左上角 区域类型是 icon, 或者checkbox 类型的
	CircleELType *string `json:"circleELType,omitempty" xml:"circleELType,omitempty"`
	// icon, name ,内容区域类型是 icon+value, 或者name+value 类型的
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// link, button, 操作区类型，是链接类型，或者按钮类型
	ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// 卡片icon
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 卡片标题
	TodoCardTitle       *string                                                               `json:"todoCardTitle,omitempty" xml:"todoCardTitle,omitempty"`
	TodoCardContentList []*QueryTodoTasksResponseBodyTodoCardsTodoCardViewTodoCardContentList `json:"todoCardContentList,omitempty" xml:"todoCardContentList,omitempty" type:"Repeated"`
}

func (s QueryTodoTasksResponseBodyTodoCardsTodoCardView) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoTasksResponseBodyTodoCardsTodoCardView) GoString() string {
	return s.String()
}

func (s *QueryTodoTasksResponseBodyTodoCardsTodoCardView) SetCardType(v string) *QueryTodoTasksResponseBodyTodoCardsTodoCardView {
	s.CardType = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCardsTodoCardView) SetCircleELType(v string) *QueryTodoTasksResponseBodyTodoCardsTodoCardView {
	s.CircleELType = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCardsTodoCardView) SetContentType(v string) *QueryTodoTasksResponseBodyTodoCardsTodoCardView {
	s.ContentType = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCardsTodoCardView) SetActionType(v string) *QueryTodoTasksResponseBodyTodoCardsTodoCardView {
	s.ActionType = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCardsTodoCardView) SetIcon(v string) *QueryTodoTasksResponseBodyTodoCardsTodoCardView {
	s.Icon = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCardsTodoCardView) SetTodoCardTitle(v string) *QueryTodoTasksResponseBodyTodoCardsTodoCardView {
	s.TodoCardTitle = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCardsTodoCardView) SetTodoCardContentList(v []*QueryTodoTasksResponseBodyTodoCardsTodoCardViewTodoCardContentList) *QueryTodoTasksResponseBodyTodoCardsTodoCardView {
	s.TodoCardContentList = v
	return s
}

type QueryTodoTasksResponseBodyTodoCardsTodoCardViewTodoCardContentList struct {
	// 自定义表单内容名字
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 自定义表单内容值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s QueryTodoTasksResponseBodyTodoCardsTodoCardViewTodoCardContentList) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoTasksResponseBodyTodoCardsTodoCardViewTodoCardContentList) GoString() string {
	return s.String()
}

func (s *QueryTodoTasksResponseBodyTodoCardsTodoCardViewTodoCardContentList) SetName(v string) *QueryTodoTasksResponseBodyTodoCardsTodoCardViewTodoCardContentList {
	s.Name = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCardsTodoCardViewTodoCardContentList) SetValue(v string) *QueryTodoTasksResponseBodyTodoCardsTodoCardViewTodoCardContentList {
	s.Value = &v
	return s
}

type QueryTodoTasksResponseBodyTodoCardsOriginalSource struct {
	// 业务来源展示名称
	SourceTitle *string `json:"sourceTitle,omitempty" xml:"sourceTitle,omitempty"`
}

func (s QueryTodoTasksResponseBodyTodoCardsOriginalSource) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoTasksResponseBodyTodoCardsOriginalSource) GoString() string {
	return s.String()
}

func (s *QueryTodoTasksResponseBodyTodoCardsOriginalSource) SetSourceTitle(v string) *QueryTodoTasksResponseBodyTodoCardsOriginalSource {
	s.SourceTitle = &v
	return s
}

type QueryTodoTasksResponseBodyTodoCardsOrgInfo struct {
	// 组织corpId
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 组织名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryTodoTasksResponseBodyTodoCardsOrgInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoTasksResponseBodyTodoCardsOrgInfo) GoString() string {
	return s.String()
}

func (s *QueryTodoTasksResponseBodyTodoCardsOrgInfo) SetCorpId(v string) *QueryTodoTasksResponseBodyTodoCardsOrgInfo {
	s.CorpId = &v
	return s
}

func (s *QueryTodoTasksResponseBodyTodoCardsOrgInfo) SetName(v string) *QueryTodoTasksResponseBodyTodoCardsOrgInfo {
	s.Name = &v
	return s
}

type QueryTodoTasksResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTodoTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTodoTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoTasksResponse) GoString() string {
	return s.String()
}

func (s *QueryTodoTasksResponse) SetHeaders(v map[string]*string) *QueryTodoTasksResponse {
	s.Headers = v
	return s
}

func (s *QueryTodoTasksResponse) SetBody(v *QueryTodoTasksResponseBody) *QueryTodoTasksResponse {
	s.Body = v
	return s
}

type UpdateTodoTypeConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateTodoTypeConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTypeConfigHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTodoTypeConfigHeaders) SetCommonHeaders(v map[string]*string) *UpdateTodoTypeConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTodoTypeConfigHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateTodoTypeConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateTodoTypeConfigRequest struct {
	// 卡片类型（取值为：1-标准卡片，2-自定义卡片）
	CardType *int32 `json:"cardType,omitempty" xml:"cardType,omitempty"`
	// 卡片类型icon，用于在待办列表展示
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 待办卡片类型描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 详情页链接在PC端的打开方式，取值为：「PC_SLIDE」-PC端侧边栏打开，「PC_BROWSER」-浏览器打开
	PcDetailUrlOpenMode *string `json:"pcDetailUrlOpenMode,omitempty" xml:"pcDetailUrlOpenMode,omitempty"`
	// 待办卡片内容区表单自定义字段配置
	ContentFieldList []*UpdateTodoTypeConfigRequestContentFieldList `json:"contentFieldList,omitempty" xml:"contentFieldList,omitempty" type:"Repeated"`
	// 待办卡片操作区按钮配置
	ActionList []*UpdateTodoTypeConfigRequestActionList `json:"actionList,omitempty" xml:"actionList,omitempty" type:"Repeated"`
	// 当前操作者id，需传用户的unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s UpdateTodoTypeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTypeConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateTodoTypeConfigRequest) SetCardType(v int32) *UpdateTodoTypeConfigRequest {
	s.CardType = &v
	return s
}

func (s *UpdateTodoTypeConfigRequest) SetIcon(v string) *UpdateTodoTypeConfigRequest {
	s.Icon = &v
	return s
}

func (s *UpdateTodoTypeConfigRequest) SetDescription(v string) *UpdateTodoTypeConfigRequest {
	s.Description = &v
	return s
}

func (s *UpdateTodoTypeConfigRequest) SetPcDetailUrlOpenMode(v string) *UpdateTodoTypeConfigRequest {
	s.PcDetailUrlOpenMode = &v
	return s
}

func (s *UpdateTodoTypeConfigRequest) SetContentFieldList(v []*UpdateTodoTypeConfigRequestContentFieldList) *UpdateTodoTypeConfigRequest {
	s.ContentFieldList = v
	return s
}

func (s *UpdateTodoTypeConfigRequest) SetActionList(v []*UpdateTodoTypeConfigRequestActionList) *UpdateTodoTypeConfigRequest {
	s.ActionList = v
	return s
}

func (s *UpdateTodoTypeConfigRequest) SetOperatorId(v string) *UpdateTodoTypeConfigRequest {
	s.OperatorId = &v
	return s
}

type UpdateTodoTypeConfigRequestContentFieldList struct {
	// 字段唯一标识
	FieldKey *string `json:"fieldKey,omitempty" xml:"fieldKey,omitempty"`
	// 字段类型（取值为：text-文本，url-链接）
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	// 字段显示名称(需支持国际化)
	NameI18n map[string]interface{} `json:"nameI18n,omitempty" xml:"nameI18n,omitempty"`
}

func (s UpdateTodoTypeConfigRequestContentFieldList) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTypeConfigRequestContentFieldList) GoString() string {
	return s.String()
}

func (s *UpdateTodoTypeConfigRequestContentFieldList) SetFieldKey(v string) *UpdateTodoTypeConfigRequestContentFieldList {
	s.FieldKey = &v
	return s
}

func (s *UpdateTodoTypeConfigRequestContentFieldList) SetFieldType(v string) *UpdateTodoTypeConfigRequestContentFieldList {
	s.FieldType = &v
	return s
}

func (s *UpdateTodoTypeConfigRequestContentFieldList) SetNameI18n(v map[string]interface{}) *UpdateTodoTypeConfigRequestContentFieldList {
	s.NameI18n = v
	return s
}

type UpdateTodoTypeConfigRequestActionList struct {
	// 操作按钮的唯一标识
	ActionKey *string `json:"actionKey,omitempty" xml:"actionKey,omitempty"`
	// 按钮样式类型（101：蓝色线型主按钮样式，例如「同意」，102：黑色线型副按钮样式，例如「拒绝」）
	ButtonStyleType *int32 `json:"buttonStyleType,omitempty" xml:"buttonStyleType,omitempty"`
	// 按钮类型（1：有操作的，2：直接跳转）
	ActionType *int32 `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// 按钮类型为直接跳转时，对应的跳转url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 按钮操作的显示名称（需支持国际化）
	NameI18n map[string]interface{} `json:"nameI18n,omitempty" xml:"nameI18n,omitempty"`
}

func (s UpdateTodoTypeConfigRequestActionList) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTypeConfigRequestActionList) GoString() string {
	return s.String()
}

func (s *UpdateTodoTypeConfigRequestActionList) SetActionKey(v string) *UpdateTodoTypeConfigRequestActionList {
	s.ActionKey = &v
	return s
}

func (s *UpdateTodoTypeConfigRequestActionList) SetButtonStyleType(v int32) *UpdateTodoTypeConfigRequestActionList {
	s.ButtonStyleType = &v
	return s
}

func (s *UpdateTodoTypeConfigRequestActionList) SetActionType(v int32) *UpdateTodoTypeConfigRequestActionList {
	s.ActionType = &v
	return s
}

func (s *UpdateTodoTypeConfigRequestActionList) SetUrl(v string) *UpdateTodoTypeConfigRequestActionList {
	s.Url = &v
	return s
}

func (s *UpdateTodoTypeConfigRequestActionList) SetNameI18n(v map[string]interface{}) *UpdateTodoTypeConfigRequestActionList {
	s.NameI18n = v
	return s
}

type UpdateTodoTypeConfigResponseBody struct {
	// 更新结果
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateTodoTypeConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTypeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTodoTypeConfigResponseBody) SetResult(v bool) *UpdateTodoTypeConfigResponseBody {
	s.Result = &v
	return s
}

type UpdateTodoTypeConfigResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTodoTypeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTodoTypeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTypeConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateTodoTypeConfigResponse) SetHeaders(v map[string]*string) *UpdateTodoTypeConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateTodoTypeConfigResponse) SetBody(v *UpdateTodoTypeConfigResponseBody) *UpdateTodoTypeConfigResponse {
	s.Body = v
	return s
}

type GetTodoTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTodoTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTaskHeaders) GoString() string {
	return s.String()
}

func (s *GetTodoTaskHeaders) SetCommonHeaders(v map[string]*string) *GetTodoTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTodoTaskHeaders) SetXAcsDingtalkAccessToken(v string) *GetTodoTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTodoTaskResponseBody struct {
	// id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 标题
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 截止时间
	DueTime *int64 `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	// 完成时间
	FinishTime *int64 `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// 完成状态
	Done *bool `json:"done,omitempty" xml:"done,omitempty"`
	// 执行者列表（用户的unionId）
	ExecutorIds []*string `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	// 参与者列表（用户的unionId）
	ParticipantIds []*string `json:"participantIds,omitempty" xml:"participantIds,omitempty" type:"Repeated"`
	// 自定义详情页跳转配置
	DetailUrl *GetTodoTaskResponseBodyDetailUrl `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	// 业务来源id
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 业务来源
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 创建时间
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 更新时间
	ModifiedTime *int64 `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// 创建者id（用户的unionId）
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 更新者id（用户的unionId）
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// 租户id(unionId/orgId/groupId)
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 租户类型（user/org/group）
	TenantType *string `json:"tenantType,omitempty" xml:"tenantType,omitempty"`
	// 接入业务应用标识
	BizTag *string `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 待办卡片类型id
	CardTypeId *string `json:"cardTypeId,omitempty" xml:"cardTypeId,omitempty"`
}

func (s GetTodoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTodoTaskResponseBody) SetId(v string) *GetTodoTaskResponseBody {
	s.Id = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetSubject(v string) *GetTodoTaskResponseBody {
	s.Subject = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetDescription(v string) *GetTodoTaskResponseBody {
	s.Description = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetStartTime(v int64) *GetTodoTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetDueTime(v int64) *GetTodoTaskResponseBody {
	s.DueTime = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetFinishTime(v int64) *GetTodoTaskResponseBody {
	s.FinishTime = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetDone(v bool) *GetTodoTaskResponseBody {
	s.Done = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetExecutorIds(v []*string) *GetTodoTaskResponseBody {
	s.ExecutorIds = v
	return s
}

func (s *GetTodoTaskResponseBody) SetParticipantIds(v []*string) *GetTodoTaskResponseBody {
	s.ParticipantIds = v
	return s
}

func (s *GetTodoTaskResponseBody) SetDetailUrl(v *GetTodoTaskResponseBodyDetailUrl) *GetTodoTaskResponseBody {
	s.DetailUrl = v
	return s
}

func (s *GetTodoTaskResponseBody) SetSourceId(v string) *GetTodoTaskResponseBody {
	s.SourceId = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetSource(v string) *GetTodoTaskResponseBody {
	s.Source = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetCreatedTime(v int64) *GetTodoTaskResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetModifiedTime(v int64) *GetTodoTaskResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetCreatorId(v string) *GetTodoTaskResponseBody {
	s.CreatorId = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetModifierId(v string) *GetTodoTaskResponseBody {
	s.ModifierId = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetTenantId(v string) *GetTodoTaskResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetTenantType(v string) *GetTodoTaskResponseBody {
	s.TenantType = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetBizTag(v string) *GetTodoTaskResponseBody {
	s.BizTag = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetRequestId(v string) *GetTodoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTodoTaskResponseBody) SetCardTypeId(v string) *GetTodoTaskResponseBody {
	s.CardTypeId = &v
	return s
}

type GetTodoTaskResponseBodyDetailUrl struct {
	// pc端详情页地址
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	// app端详情页地址
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
}

func (s GetTodoTaskResponseBodyDetailUrl) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTaskResponseBodyDetailUrl) GoString() string {
	return s.String()
}

func (s *GetTodoTaskResponseBodyDetailUrl) SetPcUrl(v string) *GetTodoTaskResponseBodyDetailUrl {
	s.PcUrl = &v
	return s
}

func (s *GetTodoTaskResponseBodyDetailUrl) SetAppUrl(v string) *GetTodoTaskResponseBodyDetailUrl {
	s.AppUrl = &v
	return s
}

type GetTodoTaskResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTodoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTodoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTodoTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTodoTaskResponse) SetHeaders(v map[string]*string) *GetTodoTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTodoTaskResponse) SetBody(v *GetTodoTaskResponseBody) *GetTodoTaskResponse {
	s.Body = v
	return s
}

type DeleteTodoTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteTodoTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteTodoTaskHeaders) GoString() string {
	return s.String()
}

func (s *DeleteTodoTaskHeaders) SetCommonHeaders(v map[string]*string) *DeleteTodoTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteTodoTaskHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteTodoTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteTodoTaskRequest struct {
	// 操作者id，需传用户的unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s DeleteTodoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTodoTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteTodoTaskRequest) SetOperatorId(v string) *DeleteTodoTaskRequest {
	s.OperatorId = &v
	return s
}

type DeleteTodoTaskResponseBody struct {
	// 删除结果
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteTodoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTodoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTodoTaskResponseBody) SetResult(v bool) *DeleteTodoTaskResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteTodoTaskResponseBody) SetRequestId(v string) *DeleteTodoTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTodoTaskResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTodoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTodoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTodoTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteTodoTaskResponse) SetHeaders(v map[string]*string) *DeleteTodoTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteTodoTaskResponse) SetBody(v *DeleteTodoTaskResponseBody) *DeleteTodoTaskResponse {
	s.Body = v
	return s
}

type UpdateTodoTaskExecutorStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateTodoTaskExecutorStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusHeaders) SetCommonHeaders(v map[string]*string) *UpdateTodoTaskExecutorStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTodoTaskExecutorStatusHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateTodoTaskExecutorStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateTodoTaskExecutorStatusRequest struct {
	// 执行者状态列表，id需传用户的unionId
	ExecutorStatusList []*UpdateTodoTaskExecutorStatusRequestExecutorStatusList `json:"executorStatusList,omitempty" xml:"executorStatusList,omitempty" type:"Repeated"`
	// 当前操作者id，需传用户的unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s UpdateTodoTaskExecutorStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusRequest) SetExecutorStatusList(v []*UpdateTodoTaskExecutorStatusRequestExecutorStatusList) *UpdateTodoTaskExecutorStatusRequest {
	s.ExecutorStatusList = v
	return s
}

func (s *UpdateTodoTaskExecutorStatusRequest) SetOperatorId(v string) *UpdateTodoTaskExecutorStatusRequest {
	s.OperatorId = &v
	return s
}

type UpdateTodoTaskExecutorStatusRequestExecutorStatusList struct {
	// 执行者id，需传用户的unionId
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 执行者完成状态
	IsDone *bool `json:"isDone,omitempty" xml:"isDone,omitempty"`
}

func (s UpdateTodoTaskExecutorStatusRequestExecutorStatusList) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusRequestExecutorStatusList) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusRequestExecutorStatusList) SetId(v string) *UpdateTodoTaskExecutorStatusRequestExecutorStatusList {
	s.Id = &v
	return s
}

func (s *UpdateTodoTaskExecutorStatusRequestExecutorStatusList) SetIsDone(v bool) *UpdateTodoTaskExecutorStatusRequestExecutorStatusList {
	s.IsDone = &v
	return s
}

type UpdateTodoTaskExecutorStatusResponseBody struct {
	// 更新结果
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateTodoTaskExecutorStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusResponseBody) SetResult(v bool) *UpdateTodoTaskExecutorStatusResponseBody {
	s.Result = &v
	return s
}

type UpdateTodoTaskExecutorStatusResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTodoTaskExecutorStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTodoTaskExecutorStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusResponse) SetHeaders(v map[string]*string) *UpdateTodoTaskExecutorStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateTodoTaskExecutorStatusResponse) SetBody(v *UpdateTodoTaskExecutorStatusResponseBody) *UpdateTodoTaskExecutorStatusResponse {
	s.Body = v
	return s
}

type CreateTodoTypeConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateTodoTypeConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTypeConfigHeaders) GoString() string {
	return s.String()
}

func (s *CreateTodoTypeConfigHeaders) SetCommonHeaders(v map[string]*string) *CreateTodoTypeConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTodoTypeConfigHeaders) SetXAcsDingtalkAccessToken(v string) *CreateTodoTypeConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateTodoTypeConfigRequest struct {
	// 卡片类型（取值为：1-标准卡片，2-自定义卡片）
	CardType *int32 `json:"cardType,omitempty" xml:"cardType,omitempty"`
	// 卡片类型icon，用于在待办列表展示
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 待办卡片类型描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 详情页链接在PC端的打开方式，取值为：「PC_SLIDE」-PC端侧边栏打开，「PC_BROWSER」-浏览器打开
	PcDetailUrlOpenMode *string `json:"pcDetailUrlOpenMode,omitempty" xml:"pcDetailUrlOpenMode,omitempty"`
	// 待办卡片内容区表单自定义字段配置
	ContentFieldList []*CreateTodoTypeConfigRequestContentFieldList `json:"contentFieldList,omitempty" xml:"contentFieldList,omitempty" type:"Repeated"`
	// 待办卡片操作区按钮配置
	ActionList []*CreateTodoTypeConfigRequestActionList `json:"actionList,omitempty" xml:"actionList,omitempty" type:"Repeated"`
	// 当前操作者id，需传用户的unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s CreateTodoTypeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTypeConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateTodoTypeConfigRequest) SetCardType(v int32) *CreateTodoTypeConfigRequest {
	s.CardType = &v
	return s
}

func (s *CreateTodoTypeConfigRequest) SetIcon(v string) *CreateTodoTypeConfigRequest {
	s.Icon = &v
	return s
}

func (s *CreateTodoTypeConfigRequest) SetDescription(v string) *CreateTodoTypeConfigRequest {
	s.Description = &v
	return s
}

func (s *CreateTodoTypeConfigRequest) SetPcDetailUrlOpenMode(v string) *CreateTodoTypeConfigRequest {
	s.PcDetailUrlOpenMode = &v
	return s
}

func (s *CreateTodoTypeConfigRequest) SetContentFieldList(v []*CreateTodoTypeConfigRequestContentFieldList) *CreateTodoTypeConfigRequest {
	s.ContentFieldList = v
	return s
}

func (s *CreateTodoTypeConfigRequest) SetActionList(v []*CreateTodoTypeConfigRequestActionList) *CreateTodoTypeConfigRequest {
	s.ActionList = v
	return s
}

func (s *CreateTodoTypeConfigRequest) SetOperatorId(v string) *CreateTodoTypeConfigRequest {
	s.OperatorId = &v
	return s
}

type CreateTodoTypeConfigRequestContentFieldList struct {
	// 字段唯一标识
	FieldKey *string `json:"fieldKey,omitempty" xml:"fieldKey,omitempty"`
	// 字段类型（取值为：text-文本，url-链接）
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	// 字段显示名称(需支持国际化)
	NameI18n map[string]interface{} `json:"nameI18n,omitempty" xml:"nameI18n,omitempty"`
}

func (s CreateTodoTypeConfigRequestContentFieldList) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTypeConfigRequestContentFieldList) GoString() string {
	return s.String()
}

func (s *CreateTodoTypeConfigRequestContentFieldList) SetFieldKey(v string) *CreateTodoTypeConfigRequestContentFieldList {
	s.FieldKey = &v
	return s
}

func (s *CreateTodoTypeConfigRequestContentFieldList) SetFieldType(v string) *CreateTodoTypeConfigRequestContentFieldList {
	s.FieldType = &v
	return s
}

func (s *CreateTodoTypeConfigRequestContentFieldList) SetNameI18n(v map[string]interface{}) *CreateTodoTypeConfigRequestContentFieldList {
	s.NameI18n = v
	return s
}

type CreateTodoTypeConfigRequestActionList struct {
	// 操作按钮的唯一标识
	ActionKey *string `json:"actionKey,omitempty" xml:"actionKey,omitempty"`
	// 按钮样式类型（101：蓝色线型主按钮样式，例如「同意」，102：黑色线型副按钮样式，例如「拒绝」）
	ButtonStyleType *int32 `json:"buttonStyleType,omitempty" xml:"buttonStyleType,omitempty"`
	// 按钮类型（1：有操作的，2：直接跳转）
	ActionType *int32 `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// 按钮类型为直接跳转时，对应的跳转url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 按钮操作的显示名称（需支持国际化）
	NameI18n map[string]interface{} `json:"nameI18n,omitempty" xml:"nameI18n,omitempty"`
}

func (s CreateTodoTypeConfigRequestActionList) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTypeConfigRequestActionList) GoString() string {
	return s.String()
}

func (s *CreateTodoTypeConfigRequestActionList) SetActionKey(v string) *CreateTodoTypeConfigRequestActionList {
	s.ActionKey = &v
	return s
}

func (s *CreateTodoTypeConfigRequestActionList) SetButtonStyleType(v int32) *CreateTodoTypeConfigRequestActionList {
	s.ButtonStyleType = &v
	return s
}

func (s *CreateTodoTypeConfigRequestActionList) SetActionType(v int32) *CreateTodoTypeConfigRequestActionList {
	s.ActionType = &v
	return s
}

func (s *CreateTodoTypeConfigRequestActionList) SetUrl(v string) *CreateTodoTypeConfigRequestActionList {
	s.Url = &v
	return s
}

func (s *CreateTodoTypeConfigRequestActionList) SetNameI18n(v map[string]interface{}) *CreateTodoTypeConfigRequestActionList {
	s.NameI18n = v
	return s
}

type CreateTodoTypeConfigResponseBody struct {
	// id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 创建时间
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 更新时间
	ModifiedTime *int64 `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// 创建者（用户的unionId）
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 更新者（用户的unionId）
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// 接入应用标识
	BizTag *string `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 卡片类型（取值为：1-标准卡片，2-自定义卡片）
	CardType *int32 `json:"cardType,omitempty" xml:"cardType,omitempty"`
	// 卡片类型icon，用于在待办列表展示
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 待办卡片类型描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 详情页链接在PC端的打开方式，取值为：「PC_SLIDE」-PC端侧边栏打开，「PC_BROWSER」-浏览器打开
	PcDetailUrlOpenMode *string `json:"pcDetailUrlOpenMode,omitempty" xml:"pcDetailUrlOpenMode,omitempty"`
	// 待办卡片内容区表单自定义字段配置
	ContentFieldList []*CreateTodoTypeConfigResponseBodyContentFieldList `json:"contentFieldList,omitempty" xml:"contentFieldList,omitempty" type:"Repeated"`
	// 待办卡片操作区按钮配置
	ActionList []*CreateTodoTypeConfigResponseBodyActionList `json:"actionList,omitempty" xml:"actionList,omitempty" type:"Repeated"`
}

func (s CreateTodoTypeConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTypeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTodoTypeConfigResponseBody) SetId(v string) *CreateTodoTypeConfigResponseBody {
	s.Id = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBody) SetCreatedTime(v int64) *CreateTodoTypeConfigResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBody) SetModifiedTime(v int64) *CreateTodoTypeConfigResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBody) SetCreatorId(v string) *CreateTodoTypeConfigResponseBody {
	s.CreatorId = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBody) SetModifierId(v string) *CreateTodoTypeConfigResponseBody {
	s.ModifierId = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBody) SetBizTag(v string) *CreateTodoTypeConfigResponseBody {
	s.BizTag = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBody) SetRequestId(v string) *CreateTodoTypeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBody) SetCardType(v int32) *CreateTodoTypeConfigResponseBody {
	s.CardType = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBody) SetIcon(v string) *CreateTodoTypeConfigResponseBody {
	s.Icon = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBody) SetDescription(v string) *CreateTodoTypeConfigResponseBody {
	s.Description = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBody) SetPcDetailUrlOpenMode(v string) *CreateTodoTypeConfigResponseBody {
	s.PcDetailUrlOpenMode = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBody) SetContentFieldList(v []*CreateTodoTypeConfigResponseBodyContentFieldList) *CreateTodoTypeConfigResponseBody {
	s.ContentFieldList = v
	return s
}

func (s *CreateTodoTypeConfigResponseBody) SetActionList(v []*CreateTodoTypeConfigResponseBodyActionList) *CreateTodoTypeConfigResponseBody {
	s.ActionList = v
	return s
}

type CreateTodoTypeConfigResponseBodyContentFieldList struct {
	// 字段唯一标识
	FieldKey *string `json:"fieldKey,omitempty" xml:"fieldKey,omitempty"`
	// 字段类型（取值为：text-文本，url-链接）
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	// 字段的显示名称（支持国际化）
	NameI18n map[string]interface{} `json:"nameI18n,omitempty" xml:"nameI18n,omitempty"`
}

func (s CreateTodoTypeConfigResponseBodyContentFieldList) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTypeConfigResponseBodyContentFieldList) GoString() string {
	return s.String()
}

func (s *CreateTodoTypeConfigResponseBodyContentFieldList) SetFieldKey(v string) *CreateTodoTypeConfigResponseBodyContentFieldList {
	s.FieldKey = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBodyContentFieldList) SetFieldType(v string) *CreateTodoTypeConfigResponseBodyContentFieldList {
	s.FieldType = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBodyContentFieldList) SetNameI18n(v map[string]interface{}) *CreateTodoTypeConfigResponseBodyContentFieldList {
	s.NameI18n = v
	return s
}

type CreateTodoTypeConfigResponseBodyActionList struct {
	// 操作按钮的唯一标识
	ActionKey *string `json:"actionKey,omitempty" xml:"actionKey,omitempty"`
	// 按钮样式类型（101：蓝色线型主按钮样式，例如「同意」，102：黑色线型副按钮样式，例如「拒绝」）
	ButtonStyleType *int32 `json:"buttonStyleType,omitempty" xml:"buttonStyleType,omitempty"`
	// 按钮类型（1：有操作的，2：直接跳转）
	ActionType *int32 `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// 按钮类型为直接跳转时，对应的跳转url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 按钮操作的显示名称（支持国际化）
	NameI18n map[string]interface{} `json:"nameI18n,omitempty" xml:"nameI18n,omitempty"`
}

func (s CreateTodoTypeConfigResponseBodyActionList) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTypeConfigResponseBodyActionList) GoString() string {
	return s.String()
}

func (s *CreateTodoTypeConfigResponseBodyActionList) SetActionKey(v string) *CreateTodoTypeConfigResponseBodyActionList {
	s.ActionKey = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBodyActionList) SetButtonStyleType(v int32) *CreateTodoTypeConfigResponseBodyActionList {
	s.ButtonStyleType = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBodyActionList) SetActionType(v int32) *CreateTodoTypeConfigResponseBodyActionList {
	s.ActionType = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBodyActionList) SetUrl(v string) *CreateTodoTypeConfigResponseBodyActionList {
	s.Url = &v
	return s
}

func (s *CreateTodoTypeConfigResponseBodyActionList) SetNameI18n(v map[string]interface{}) *CreateTodoTypeConfigResponseBodyActionList {
	s.NameI18n = v
	return s
}

type CreateTodoTypeConfigResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTodoTypeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTodoTypeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTypeConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateTodoTypeConfigResponse) SetHeaders(v map[string]*string) *CreateTodoTypeConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateTodoTypeConfigResponse) SetBody(v *CreateTodoTypeConfigResponseBody) *CreateTodoTypeConfigResponse {
	s.Body = v
	return s
}

type CountTodoTasksHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CountTodoTasksHeaders) String() string {
	return tea.Prettify(s)
}

func (s CountTodoTasksHeaders) GoString() string {
	return s.String()
}

func (s *CountTodoTasksHeaders) SetCommonHeaders(v map[string]*string) *CountTodoTasksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CountTodoTasksHeaders) SetXAcsDingtalkAccessToken(v string) *CountTodoTasksHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CountTodoTasksRequest struct {
	// 待办完成状态。
	IsDone *bool `json:"isDone,omitempty" xml:"isDone,omitempty"`
	// 查询目标用户角色类型，执行人 | 创建人 | 参与人，可以同时传多个值。如：[["executor"], ["creator"],["participant"]] 或 [["executor", "creator"]]
	RoleTypes [][]*string `json:"roleTypes,omitempty" xml:"roleTypes,omitempty" type:"Repeated"`
	// 查询从计划完成时间开始
	FromDueTime *int64 `json:"fromDueTime,omitempty" xml:"fromDueTime,omitempty"`
	// 查询到计划完成时间结束
	ToDueTime *int64 `json:"toDueTime,omitempty" xml:"toDueTime,omitempty"`
	// 所属分类
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 待办回收状态
	IsRecycled *bool `json:"isRecycled,omitempty" xml:"isRecycled,omitempty"`
}

func (s CountTodoTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s CountTodoTasksRequest) GoString() string {
	return s.String()
}

func (s *CountTodoTasksRequest) SetIsDone(v bool) *CountTodoTasksRequest {
	s.IsDone = &v
	return s
}

func (s *CountTodoTasksRequest) SetRoleTypes(v [][]*string) *CountTodoTasksRequest {
	s.RoleTypes = v
	return s
}

func (s *CountTodoTasksRequest) SetFromDueTime(v int64) *CountTodoTasksRequest {
	s.FromDueTime = &v
	return s
}

func (s *CountTodoTasksRequest) SetToDueTime(v int64) *CountTodoTasksRequest {
	s.ToDueTime = &v
	return s
}

func (s *CountTodoTasksRequest) SetCategory(v string) *CountTodoTasksRequest {
	s.Category = &v
	return s
}

func (s *CountTodoTasksRequest) SetIsRecycled(v bool) *CountTodoTasksRequest {
	s.IsRecycled = &v
	return s
}

type CountTodoTasksResponseBody struct {
	// 待办数量
	Result *int32 `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CountTodoTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountTodoTasksResponseBody) GoString() string {
	return s.String()
}

func (s *CountTodoTasksResponseBody) SetResult(v int32) *CountTodoTasksResponseBody {
	s.Result = &v
	return s
}

type CountTodoTasksResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CountTodoTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CountTodoTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s CountTodoTasksResponse) GoString() string {
	return s.String()
}

func (s *CountTodoTasksResponse) SetHeaders(v map[string]*string) *CountTodoTasksResponse {
	s.Headers = v
	return s
}

func (s *CountTodoTasksResponse) SetBody(v *CountTodoTasksResponseBody) *CountTodoTasksResponse {
	s.Body = v
	return s
}

type UpdateTodoTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateTodoTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskHeaders) SetCommonHeaders(v map[string]*string) *UpdateTodoTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTodoTaskHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateTodoTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateTodoTaskRequest struct {
	// 待办标题
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// 待办描述备注
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 截止时间
	DueTime *int64 `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	// 完成状态
	Done *bool `json:"done,omitempty" xml:"done,omitempty"`
	// 执行者列表，需传用户的unionId
	ExecutorIds []*string `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	// 参与者列表，需传用户的unionId
	ParticipantIds []*string `json:"participantIds,omitempty" xml:"participantIds,omitempty" type:"Repeated"`
	// 当前操作者id，需传用户的unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s UpdateTodoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskRequest) SetSubject(v string) *UpdateTodoTaskRequest {
	s.Subject = &v
	return s
}

func (s *UpdateTodoTaskRequest) SetDescription(v string) *UpdateTodoTaskRequest {
	s.Description = &v
	return s
}

func (s *UpdateTodoTaskRequest) SetDueTime(v int64) *UpdateTodoTaskRequest {
	s.DueTime = &v
	return s
}

func (s *UpdateTodoTaskRequest) SetDone(v bool) *UpdateTodoTaskRequest {
	s.Done = &v
	return s
}

func (s *UpdateTodoTaskRequest) SetExecutorIds(v []*string) *UpdateTodoTaskRequest {
	s.ExecutorIds = v
	return s
}

func (s *UpdateTodoTaskRequest) SetParticipantIds(v []*string) *UpdateTodoTaskRequest {
	s.ParticipantIds = v
	return s
}

func (s *UpdateTodoTaskRequest) SetOperatorId(v string) *UpdateTodoTaskRequest {
	s.OperatorId = &v
	return s
}

type UpdateTodoTaskResponseBody struct {
	// 更新结果
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateTodoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskResponseBody) SetResult(v bool) *UpdateTodoTaskResponseBody {
	s.Result = &v
	return s
}

type UpdateTodoTaskResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTodoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTodoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskResponse) SetHeaders(v map[string]*string) *UpdateTodoTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateTodoTaskResponse) SetBody(v *UpdateTodoTaskResponseBody) *UpdateTodoTaskResponse {
	s.Body = v
	return s
}

type CreateTodoTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateTodoTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskHeaders) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskHeaders) SetCommonHeaders(v map[string]*string) *CreateTodoTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTodoTaskHeaders) SetXAcsDingtalkAccessToken(v string) *CreateTodoTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateTodoTaskRequest struct {
	// 来源id，接入业务系统侧的唯一标识id
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 待办标题
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// 创建者id，需传用户的unionId
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 待办备注描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 截止时间
	DueTime *int64 `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	// 执行者列表，需传用户的unionId
	ExecutorIds []*string `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	// 参与者列表，需传用户的unionId
	ParticipantIds []*string `json:"participantIds,omitempty" xml:"participantIds,omitempty" type:"Repeated"`
	// 详情页url跳转地址
	DetailUrl *CreateTodoTaskRequestDetailUrl `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	// 当前操作者id，需传用户的unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s CreateTodoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskRequest) SetSourceId(v string) *CreateTodoTaskRequest {
	s.SourceId = &v
	return s
}

func (s *CreateTodoTaskRequest) SetSubject(v string) *CreateTodoTaskRequest {
	s.Subject = &v
	return s
}

func (s *CreateTodoTaskRequest) SetCreatorId(v string) *CreateTodoTaskRequest {
	s.CreatorId = &v
	return s
}

func (s *CreateTodoTaskRequest) SetDescription(v string) *CreateTodoTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateTodoTaskRequest) SetDueTime(v int64) *CreateTodoTaskRequest {
	s.DueTime = &v
	return s
}

func (s *CreateTodoTaskRequest) SetExecutorIds(v []*string) *CreateTodoTaskRequest {
	s.ExecutorIds = v
	return s
}

func (s *CreateTodoTaskRequest) SetParticipantIds(v []*string) *CreateTodoTaskRequest {
	s.ParticipantIds = v
	return s
}

func (s *CreateTodoTaskRequest) SetDetailUrl(v *CreateTodoTaskRequestDetailUrl) *CreateTodoTaskRequest {
	s.DetailUrl = v
	return s
}

func (s *CreateTodoTaskRequest) SetOperatorId(v string) *CreateTodoTaskRequest {
	s.OperatorId = &v
	return s
}

type CreateTodoTaskRequestDetailUrl struct {
	// app端详情页url
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
	// pc端详情页url
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
}

func (s CreateTodoTaskRequestDetailUrl) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskRequestDetailUrl) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskRequestDetailUrl) SetAppUrl(v string) *CreateTodoTaskRequestDetailUrl {
	s.AppUrl = &v
	return s
}

func (s *CreateTodoTaskRequestDetailUrl) SetPcUrl(v string) *CreateTodoTaskRequestDetailUrl {
	s.PcUrl = &v
	return s
}

type CreateTodoTaskResponseBody struct {
	// id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 标题
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 截止时间
	DueTime *int64 `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	// 完成时间
	FinishTime *int64 `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// 完成状态
	Done *bool `json:"done,omitempty" xml:"done,omitempty"`
	// 执行者列表（用户的unionId）
	ExecutorIds []*string `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	// 参与者列表（用户的unionId）
	ParticipantIds []*string `json:"participantIds,omitempty" xml:"participantIds,omitempty" type:"Repeated"`
	// 自定义详情页跳转配置
	DetailUrl *CreateTodoTaskResponseBodyDetailUrl `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	// 业务来源
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 业务来源id
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 创建时间
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 更新时间
	ModifiedTime *int64 `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// 创建者（用户的unionId）
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 更新者（用户的unionId）
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// 接入应用标识
	BizTag *string `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateTodoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskResponseBody) SetId(v string) *CreateTodoTaskResponseBody {
	s.Id = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetSubject(v string) *CreateTodoTaskResponseBody {
	s.Subject = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetDescription(v string) *CreateTodoTaskResponseBody {
	s.Description = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetStartTime(v int64) *CreateTodoTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetDueTime(v int64) *CreateTodoTaskResponseBody {
	s.DueTime = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetFinishTime(v int64) *CreateTodoTaskResponseBody {
	s.FinishTime = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetDone(v bool) *CreateTodoTaskResponseBody {
	s.Done = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetExecutorIds(v []*string) *CreateTodoTaskResponseBody {
	s.ExecutorIds = v
	return s
}

func (s *CreateTodoTaskResponseBody) SetParticipantIds(v []*string) *CreateTodoTaskResponseBody {
	s.ParticipantIds = v
	return s
}

func (s *CreateTodoTaskResponseBody) SetDetailUrl(v *CreateTodoTaskResponseBodyDetailUrl) *CreateTodoTaskResponseBody {
	s.DetailUrl = v
	return s
}

func (s *CreateTodoTaskResponseBody) SetSource(v string) *CreateTodoTaskResponseBody {
	s.Source = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetSourceId(v string) *CreateTodoTaskResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetCreatedTime(v int64) *CreateTodoTaskResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetModifiedTime(v int64) *CreateTodoTaskResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetCreatorId(v string) *CreateTodoTaskResponseBody {
	s.CreatorId = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetModifierId(v string) *CreateTodoTaskResponseBody {
	s.ModifierId = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetBizTag(v string) *CreateTodoTaskResponseBody {
	s.BizTag = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetRequestId(v string) *CreateTodoTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateTodoTaskResponseBodyDetailUrl struct {
	// pc端详情页地址
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	// app端详情页地址
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
}

func (s CreateTodoTaskResponseBodyDetailUrl) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskResponseBodyDetailUrl) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskResponseBodyDetailUrl) SetPcUrl(v string) *CreateTodoTaskResponseBodyDetailUrl {
	s.PcUrl = &v
	return s
}

func (s *CreateTodoTaskResponseBodyDetailUrl) SetAppUrl(v string) *CreateTodoTaskResponseBodyDetailUrl {
	s.AppUrl = &v
	return s
}

type CreateTodoTaskResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTodoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTodoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskResponse) SetHeaders(v map[string]*string) *CreateTodoTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateTodoTaskResponse) SetBody(v *CreateTodoTaskResponseBody) *CreateTodoTaskResponse {
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

func (client *Client) GetTodoTypeConfig(unionId *string, cardTypeId *string) (_result *GetTodoTypeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTodoTypeConfigHeaders{}
	_result = &GetTodoTypeConfigResponse{}
	_body, _err := client.GetTodoTypeConfigWithOptions(unionId, cardTypeId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTodoTypeConfigWithOptions(unionId *string, cardTypeId *string, headers *GetTodoTypeConfigHeaders, runtime *util.RuntimeOptions) (_result *GetTodoTypeConfigResponse, _err error) {
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
	_result = &GetTodoTypeConfigResponse{}
	_body, _err := client.DoROARequest(tea.String("GetTodoTypeConfig"), tea.String("todo_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/todo/users/"+tea.StringValue(unionId)+"/configs/types/"+tea.StringValue(cardTypeId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTodoTasks(unionId *string, request *QueryTodoTasksRequest) (_result *QueryTodoTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryTodoTasksHeaders{}
	_result = &QueryTodoTasksResponse{}
	_body, _err := client.QueryTodoTasksWithOptions(unionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTodoTasksWithOptions(unionId *string, request *QueryTodoTasksRequest, headers *QueryTodoTasksHeaders, runtime *util.RuntimeOptions) (_result *QueryTodoTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		body["orderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.OrderDirection)) {
		body["orderDirection"] = request.OrderDirection
	}

	if !tea.BoolValue(util.IsUnset(request.IsDone)) {
		body["isDone"] = request.IsDone
	}

	if !tea.BoolValue(util.IsUnset(request.RoleTypes)) {
		body["roleTypes"] = request.RoleTypes
	}

	if !tea.BoolValue(util.IsUnset(request.FromDueTime)) {
		body["fromDueTime"] = request.FromDueTime
	}

	if !tea.BoolValue(util.IsUnset(request.ToDueTime)) {
		body["toDueTime"] = request.ToDueTime
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.IsRecycled)) {
		body["isRecycled"] = request.IsRecycled
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
	_result = &QueryTodoTasksResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryTodoTasks"), tea.String("todo_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/todo/users/"+tea.StringValue(unionId)+"/tasks/list"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTodoTypeConfig(unionId *string, cardTypeId *string, request *UpdateTodoTypeConfigRequest) (_result *UpdateTodoTypeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTodoTypeConfigHeaders{}
	_result = &UpdateTodoTypeConfigResponse{}
	_body, _err := client.UpdateTodoTypeConfigWithOptions(unionId, cardTypeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTodoTypeConfigWithOptions(unionId *string, cardTypeId *string, request *UpdateTodoTypeConfigRequest, headers *UpdateTodoTypeConfigHeaders, runtime *util.RuntimeOptions) (_result *UpdateTodoTypeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CardType)) {
		body["cardType"] = request.CardType
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.PcDetailUrlOpenMode)) {
		body["pcDetailUrlOpenMode"] = request.PcDetailUrlOpenMode
	}

	if !tea.BoolValue(util.IsUnset(request.ContentFieldList)) {
		body["contentFieldList"] = request.ContentFieldList
	}

	if !tea.BoolValue(util.IsUnset(request.ActionList)) {
		body["actionList"] = request.ActionList
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateTodoTypeConfigResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateTodoTypeConfig"), tea.String("todo_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/todo/users/"+tea.StringValue(unionId)+"/configs/types/"+tea.StringValue(cardTypeId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTodoTask(unionId *string, taskId *string) (_result *GetTodoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTodoTaskHeaders{}
	_result = &GetTodoTaskResponse{}
	_body, _err := client.GetTodoTaskWithOptions(unionId, taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTodoTaskWithOptions(unionId *string, taskId *string, headers *GetTodoTaskHeaders, runtime *util.RuntimeOptions) (_result *GetTodoTaskResponse, _err error) {
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
	_result = &GetTodoTaskResponse{}
	_body, _err := client.DoROARequest(tea.String("GetTodoTask"), tea.String("todo_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/todo/users/"+tea.StringValue(unionId)+"/tasks/"+tea.StringValue(taskId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTodoTask(unionId *string, taskId *string, request *DeleteTodoTaskRequest) (_result *DeleteTodoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteTodoTaskHeaders{}
	_result = &DeleteTodoTaskResponse{}
	_body, _err := client.DeleteTodoTaskWithOptions(unionId, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTodoTaskWithOptions(unionId *string, taskId *string, request *DeleteTodoTaskRequest, headers *DeleteTodoTaskHeaders, runtime *util.RuntimeOptions) (_result *DeleteTodoTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
	_result = &DeleteTodoTaskResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteTodoTask"), tea.String("todo_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/todo/users/"+tea.StringValue(unionId)+"/tasks/"+tea.StringValue(taskId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTodoTaskExecutorStatus(unionId *string, taskId *string, request *UpdateTodoTaskExecutorStatusRequest) (_result *UpdateTodoTaskExecutorStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTodoTaskExecutorStatusHeaders{}
	_result = &UpdateTodoTaskExecutorStatusResponse{}
	_body, _err := client.UpdateTodoTaskExecutorStatusWithOptions(unionId, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTodoTaskExecutorStatusWithOptions(unionId *string, taskId *string, request *UpdateTodoTaskExecutorStatusRequest, headers *UpdateTodoTaskExecutorStatusHeaders, runtime *util.RuntimeOptions) (_result *UpdateTodoTaskExecutorStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExecutorStatusList)) {
		body["executorStatusList"] = request.ExecutorStatusList
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateTodoTaskExecutorStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateTodoTaskExecutorStatus"), tea.String("todo_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/todo/users/"+tea.StringValue(unionId)+"/tasks/"+tea.StringValue(taskId)+"/executorStatus"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTodoTypeConfig(unionId *string, request *CreateTodoTypeConfigRequest) (_result *CreateTodoTypeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTodoTypeConfigHeaders{}
	_result = &CreateTodoTypeConfigResponse{}
	_body, _err := client.CreateTodoTypeConfigWithOptions(unionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTodoTypeConfigWithOptions(unionId *string, request *CreateTodoTypeConfigRequest, headers *CreateTodoTypeConfigHeaders, runtime *util.RuntimeOptions) (_result *CreateTodoTypeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CardType)) {
		body["cardType"] = request.CardType
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.PcDetailUrlOpenMode)) {
		body["pcDetailUrlOpenMode"] = request.PcDetailUrlOpenMode
	}

	if !tea.BoolValue(util.IsUnset(request.ContentFieldList)) {
		body["contentFieldList"] = request.ContentFieldList
	}

	if !tea.BoolValue(util.IsUnset(request.ActionList)) {
		body["actionList"] = request.ActionList
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateTodoTypeConfigResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateTodoTypeConfig"), tea.String("todo_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/todo/users/"+tea.StringValue(unionId)+"/configs/types"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CountTodoTasks(unionId *string, request *CountTodoTasksRequest) (_result *CountTodoTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CountTodoTasksHeaders{}
	_result = &CountTodoTasksResponse{}
	_body, _err := client.CountTodoTasksWithOptions(unionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CountTodoTasksWithOptions(unionId *string, request *CountTodoTasksRequest, headers *CountTodoTasksHeaders, runtime *util.RuntimeOptions) (_result *CountTodoTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsDone)) {
		body["isDone"] = request.IsDone
	}

	if !tea.BoolValue(util.IsUnset(request.RoleTypes)) {
		body["roleTypes"] = request.RoleTypes
	}

	if !tea.BoolValue(util.IsUnset(request.FromDueTime)) {
		body["fromDueTime"] = request.FromDueTime
	}

	if !tea.BoolValue(util.IsUnset(request.ToDueTime)) {
		body["toDueTime"] = request.ToDueTime
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.IsRecycled)) {
		body["isRecycled"] = request.IsRecycled
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
	_result = &CountTodoTasksResponse{}
	_body, _err := client.DoROARequest(tea.String("CountTodoTasks"), tea.String("todo_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/todo/users/"+tea.StringValue(unionId)+"/tasks/count"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTodoTask(unionId *string, taskId *string, request *UpdateTodoTaskRequest) (_result *UpdateTodoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTodoTaskHeaders{}
	_result = &UpdateTodoTaskResponse{}
	_body, _err := client.UpdateTodoTaskWithOptions(unionId, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTodoTaskWithOptions(unionId *string, taskId *string, request *UpdateTodoTaskRequest, headers *UpdateTodoTaskHeaders, runtime *util.RuntimeOptions) (_result *UpdateTodoTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		body["subject"] = request.Subject
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DueTime)) {
		body["dueTime"] = request.DueTime
	}

	if !tea.BoolValue(util.IsUnset(request.Done)) {
		body["done"] = request.Done
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorIds)) {
		body["executorIds"] = request.ExecutorIds
	}

	if !tea.BoolValue(util.IsUnset(request.ParticipantIds)) {
		body["participantIds"] = request.ParticipantIds
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateTodoTaskResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateTodoTask"), tea.String("todo_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/todo/users/"+tea.StringValue(unionId)+"/tasks/"+tea.StringValue(taskId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTodoTask(unionId *string, request *CreateTodoTaskRequest) (_result *CreateTodoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTodoTaskHeaders{}
	_result = &CreateTodoTaskResponse{}
	_body, _err := client.CreateTodoTaskWithOptions(unionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTodoTaskWithOptions(unionId *string, request *CreateTodoTaskRequest, headers *CreateTodoTaskHeaders, runtime *util.RuntimeOptions) (_result *CreateTodoTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		body["sourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		body["subject"] = request.Subject
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorId)) {
		body["creatorId"] = request.CreatorId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DueTime)) {
		body["dueTime"] = request.DueTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorIds)) {
		body["executorIds"] = request.ExecutorIds
	}

	if !tea.BoolValue(util.IsUnset(request.ParticipantIds)) {
		body["participantIds"] = request.ParticipantIds
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.DetailUrl))) {
		body["detailUrl"] = request.DetailUrl
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateTodoTaskResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateTodoTask"), tea.String("todo_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/todo/users/"+tea.StringValue(unionId)+"/tasks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
