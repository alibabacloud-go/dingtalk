// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package service_group_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DeleteKnowledgeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteKnowledgeHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteKnowledgeHeaders) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeHeaders) SetCommonHeaders(v map[string]*string) *DeleteKnowledgeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteKnowledgeHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteKnowledgeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteKnowledgeRequest struct {
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 知识库的唯一标识 比如:天工知识库ID
	LibraryKey *string `json:"libraryKey,omitempty" xml:"libraryKey,omitempty"`
	// 知识点来源 CCM:天工知识库
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 知识点唯一标识
	SourcePrimaryKey *string `json:"sourcePrimaryKey,omitempty" xml:"sourcePrimaryKey,omitempty"`
}

func (s DeleteKnowledgeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeRequest) SetDingIsvOrgId(v int64) *DeleteKnowledgeRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *DeleteKnowledgeRequest) SetDingOrgId(v int64) *DeleteKnowledgeRequest {
	s.DingOrgId = &v
	return s
}

func (s *DeleteKnowledgeRequest) SetDingSuiteKey(v string) *DeleteKnowledgeRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *DeleteKnowledgeRequest) SetDingTokenGrantType(v int64) *DeleteKnowledgeRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *DeleteKnowledgeRequest) SetOpenTeamId(v string) *DeleteKnowledgeRequest {
	s.OpenTeamId = &v
	return s
}

func (s *DeleteKnowledgeRequest) SetLibraryKey(v string) *DeleteKnowledgeRequest {
	s.LibraryKey = &v
	return s
}

func (s *DeleteKnowledgeRequest) SetSource(v string) *DeleteKnowledgeRequest {
	s.Source = &v
	return s
}

func (s *DeleteKnowledgeRequest) SetSourcePrimaryKey(v string) *DeleteKnowledgeRequest {
	s.SourcePrimaryKey = &v
	return s
}

type DeleteKnowledgeResponseBody struct {
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteKnowledgeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeResponseBody) SetSuccess(v bool) *DeleteKnowledgeResponseBody {
	s.Success = &v
	return s
}

type DeleteKnowledgeResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteKnowledgeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeResponse) SetHeaders(v map[string]*string) *DeleteKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *DeleteKnowledgeResponse) SetBody(v *DeleteKnowledgeResponseBody) *DeleteKnowledgeResponse {
	s.Body = v
	return s
}

type SearchGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchGroupHeaders) GoString() string {
	return s.String()
}

func (s *SearchGroupHeaders) SetCommonHeaders(v map[string]*string) *SearchGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchGroupHeaders) SetXAcsDingtalkAccessToken(v string) *SearchGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchGroupRequest struct {
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	// 开放群ID
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 群名称
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 开群组ID
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// 用来标记当前开始读取的位置，置空表示从头开始。
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 本次读取的最大数据记录数量，此参数为可选参数，用户传入为空时，应该有默认值。应设置最大值限制，最大不超过100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
}

func (s SearchGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchGroupRequest) GoString() string {
	return s.String()
}

func (s *SearchGroupRequest) SetDingIsvOrgId(v int64) *SearchGroupRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *SearchGroupRequest) SetDingOrgId(v int64) *SearchGroupRequest {
	s.DingOrgId = &v
	return s
}

func (s *SearchGroupRequest) SetDingSuiteKey(v string) *SearchGroupRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *SearchGroupRequest) SetDingTokenGrantType(v int64) *SearchGroupRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *SearchGroupRequest) SetOpenConversationId(v string) *SearchGroupRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SearchGroupRequest) SetGroupName(v string) *SearchGroupRequest {
	s.GroupName = &v
	return s
}

func (s *SearchGroupRequest) SetOpenTeamId(v string) *SearchGroupRequest {
	s.OpenTeamId = &v
	return s
}

func (s *SearchGroupRequest) SetOpenGroupSetId(v string) *SearchGroupRequest {
	s.OpenGroupSetId = &v
	return s
}

func (s *SearchGroupRequest) SetNextToken(v string) *SearchGroupRequest {
	s.NextToken = &v
	return s
}

func (s *SearchGroupRequest) SetMaxResults(v int32) *SearchGroupRequest {
	s.MaxResults = &v
	return s
}

type SearchGroupResponseBody struct {
	// 本次请求条件下的数据总量，此参数为可选参数，默认可不返回。本次请求条件下的数据总量，此参数为可选参数，默认可不返回
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 本次请求所返回的最大记录条数。
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 已读未读信息列表
	Records []*SearchGroupResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
}

func (s SearchGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchGroupResponseBody) GoString() string {
	return s.String()
}

func (s *SearchGroupResponseBody) SetTotalCount(v int32) *SearchGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchGroupResponseBody) SetNextToken(v string) *SearchGroupResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchGroupResponseBody) SetMaxResults(v int32) *SearchGroupResponseBody {
	s.MaxResults = &v
	return s
}

func (s *SearchGroupResponseBody) SetRecords(v []*SearchGroupResponseBodyRecords) *SearchGroupResponseBody {
	s.Records = v
	return s
}

type SearchGroupResponseBodyRecords struct {
	// 开放群ID
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 群名称
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 开放群组ID
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
}

func (s SearchGroupResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s SearchGroupResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *SearchGroupResponseBodyRecords) SetOpenConversationId(v string) *SearchGroupResponseBodyRecords {
	s.OpenConversationId = &v
	return s
}

func (s *SearchGroupResponseBodyRecords) SetGroupName(v string) *SearchGroupResponseBodyRecords {
	s.GroupName = &v
	return s
}

func (s *SearchGroupResponseBodyRecords) SetOpenTeamId(v string) *SearchGroupResponseBodyRecords {
	s.OpenTeamId = &v
	return s
}

func (s *SearchGroupResponseBodyRecords) SetOpenGroupSetId(v string) *SearchGroupResponseBodyRecords {
	s.OpenGroupSetId = &v
	return s
}

type SearchGroupResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchGroupResponse) GoString() string {
	return s.String()
}

func (s *SearchGroupResponse) SetHeaders(v map[string]*string) *SearchGroupResponse {
	s.Headers = v
	return s
}

func (s *SearchGroupResponse) SetBody(v *SearchGroupResponseBody) *SearchGroupResponse {
	s.Body = v
	return s
}

type CreateGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupHeaders) GoString() string {
	return s.String()
}

func (s *CreateGroupHeaders) SetCommonHeaders(v map[string]*string) *CreateGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateGroupHeaders) SetXAcsDingtalkAccessToken(v string) *CreateGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateGroupRequest struct {
	// 业务关联id
	GroupBizId *string `json:"groupBizId,omitempty" xml:"groupBizId,omitempty"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 开放群组ID
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// 群名称
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// 群主员工ID
	OwnerStaffId *string `json:"ownerStaffId,omitempty" xml:"ownerStaffId,omitempty"`
	// 群成员员工ID列表
	MemberStaffIds     []*string `json:"memberStaffIds,omitempty" xml:"memberStaffIds,omitempty" type:"Repeated"`
	DingIsvOrgId       *int64    `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId          *int64    `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingSuiteKey       *string   `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int64    `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
}

func (s CreateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) SetGroupBizId(v string) *CreateGroupRequest {
	s.GroupBizId = &v
	return s
}

func (s *CreateGroupRequest) SetOpenTeamId(v string) *CreateGroupRequest {
	s.OpenTeamId = &v
	return s
}

func (s *CreateGroupRequest) SetOpenGroupSetId(v string) *CreateGroupRequest {
	s.OpenGroupSetId = &v
	return s
}

func (s *CreateGroupRequest) SetGroupName(v string) *CreateGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateGroupRequest) SetOwnerStaffId(v string) *CreateGroupRequest {
	s.OwnerStaffId = &v
	return s
}

func (s *CreateGroupRequest) SetMemberStaffIds(v []*string) *CreateGroupRequest {
	s.MemberStaffIds = v
	return s
}

func (s *CreateGroupRequest) SetDingIsvOrgId(v int64) *CreateGroupRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *CreateGroupRequest) SetDingOrgId(v int64) *CreateGroupRequest {
	s.DingOrgId = &v
	return s
}

func (s *CreateGroupRequest) SetDingSuiteKey(v string) *CreateGroupRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *CreateGroupRequest) SetDingTokenGrantType(v int64) *CreateGroupRequest {
	s.DingTokenGrantType = &v
	return s
}

type CreateGroupResponseBody struct {
	// 开放群ID
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 入群url
	GroupUrl *string `json:"groupUrl,omitempty" xml:"groupUrl,omitempty"`
}

func (s CreateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBody) SetOpenConversationId(v string) *CreateGroupResponseBody {
	s.OpenConversationId = &v
	return s
}

func (s *CreateGroupResponseBody) SetGroupUrl(v string) *CreateGroupResponseBody {
	s.GroupUrl = &v
	return s
}

type CreateGroupResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupResponse) SetHeaders(v map[string]*string) *CreateGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupResponse) SetBody(v *CreateGroupResponseBody) *CreateGroupResponse {
	s.Body = v
	return s
}

type SendServiceGroupMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendServiceGroupMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendServiceGroupMessageHeaders) GoString() string {
	return s.String()
}

func (s *SendServiceGroupMessageHeaders) SetCommonHeaders(v map[string]*string) *SendServiceGroupMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendServiceGroupMessageHeaders) SetXAcsDingtalkAccessToken(v string) *SendServiceGroupMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendServiceGroupMessageRequest struct {
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	// 开放群ID
	TargetOpenConversationId *string `json:"targetOpenConversationId,omitempty" xml:"targetOpenConversationId,omitempty"`
	// 标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 是否 at所有人
	IsAtAll *bool `json:"isAtAll,omitempty" xml:"isAtAll,omitempty"`
	// at 手机号
	AtMobiles []*string `json:"atMobiles,omitempty" xml:"atMobiles,omitempty" type:"Repeated"`
	// at dingtalkId
	AtDingtalkIds []*string `json:"atDingtalkIds,omitempty" xml:"atDingtalkIds,omitempty" type:"Repeated"`
	// at unionIds
	AtUnionIds []*string `json:"atUnionIds,omitempty" xml:"atUnionIds,omitempty" type:"Repeated"`
	// 手机号接收者
	ReceiverMobiles []*string `json:"receiverMobiles,omitempty" xml:"receiverMobiles,omitempty" type:"Repeated"`
	// dingtalkId接收者
	ReceiverDingtalkIds []*string `json:"receiverDingtalkIds,omitempty" xml:"receiverDingtalkIds,omitempty" type:"Repeated"`
	// unionId接收者
	ReceiverUnionIds []*string `json:"receiverUnionIds,omitempty" xml:"receiverUnionIds,omitempty" type:"Repeated"`
	// 消息类型：MARKDOWN，ACTIONCARD
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// 排列方式：0-按钮竖直排列，1-按钮横向排列
	BtnOrientation *string `json:"btnOrientation,omitempty" xml:"btnOrientation,omitempty"`
	// actionCard按钮
	Btns []*SendServiceGroupMessageRequestBtns `json:"btns,omitempty" xml:"btns,omitempty" type:"Repeated"`
}

func (s SendServiceGroupMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendServiceGroupMessageRequest) GoString() string {
	return s.String()
}

func (s *SendServiceGroupMessageRequest) SetDingIsvOrgId(v int64) *SendServiceGroupMessageRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *SendServiceGroupMessageRequest) SetDingOrgId(v int64) *SendServiceGroupMessageRequest {
	s.DingOrgId = &v
	return s
}

func (s *SendServiceGroupMessageRequest) SetDingTokenGrantType(v int64) *SendServiceGroupMessageRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *SendServiceGroupMessageRequest) SetDingSuiteKey(v string) *SendServiceGroupMessageRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *SendServiceGroupMessageRequest) SetTargetOpenConversationId(v string) *SendServiceGroupMessageRequest {
	s.TargetOpenConversationId = &v
	return s
}

func (s *SendServiceGroupMessageRequest) SetTitle(v string) *SendServiceGroupMessageRequest {
	s.Title = &v
	return s
}

func (s *SendServiceGroupMessageRequest) SetContent(v string) *SendServiceGroupMessageRequest {
	s.Content = &v
	return s
}

func (s *SendServiceGroupMessageRequest) SetIsAtAll(v bool) *SendServiceGroupMessageRequest {
	s.IsAtAll = &v
	return s
}

func (s *SendServiceGroupMessageRequest) SetAtMobiles(v []*string) *SendServiceGroupMessageRequest {
	s.AtMobiles = v
	return s
}

func (s *SendServiceGroupMessageRequest) SetAtDingtalkIds(v []*string) *SendServiceGroupMessageRequest {
	s.AtDingtalkIds = v
	return s
}

func (s *SendServiceGroupMessageRequest) SetAtUnionIds(v []*string) *SendServiceGroupMessageRequest {
	s.AtUnionIds = v
	return s
}

func (s *SendServiceGroupMessageRequest) SetReceiverMobiles(v []*string) *SendServiceGroupMessageRequest {
	s.ReceiverMobiles = v
	return s
}

func (s *SendServiceGroupMessageRequest) SetReceiverDingtalkIds(v []*string) *SendServiceGroupMessageRequest {
	s.ReceiverDingtalkIds = v
	return s
}

func (s *SendServiceGroupMessageRequest) SetReceiverUnionIds(v []*string) *SendServiceGroupMessageRequest {
	s.ReceiverUnionIds = v
	return s
}

func (s *SendServiceGroupMessageRequest) SetMessageType(v string) *SendServiceGroupMessageRequest {
	s.MessageType = &v
	return s
}

func (s *SendServiceGroupMessageRequest) SetBtnOrientation(v string) *SendServiceGroupMessageRequest {
	s.BtnOrientation = &v
	return s
}

func (s *SendServiceGroupMessageRequest) SetBtns(v []*SendServiceGroupMessageRequestBtns) *SendServiceGroupMessageRequest {
	s.Btns = v
	return s
}

type SendServiceGroupMessageRequestBtns struct {
	// 跳转地址
	ActionURL *string `json:"actionURL,omitempty" xml:"actionURL,omitempty"`
	// 按钮名称
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SendServiceGroupMessageRequestBtns) String() string {
	return tea.Prettify(s)
}

func (s SendServiceGroupMessageRequestBtns) GoString() string {
	return s.String()
}

func (s *SendServiceGroupMessageRequestBtns) SetActionURL(v string) *SendServiceGroupMessageRequestBtns {
	s.ActionURL = &v
	return s
}

func (s *SendServiceGroupMessageRequestBtns) SetTitle(v string) *SendServiceGroupMessageRequestBtns {
	s.Title = &v
	return s
}

type SendServiceGroupMessageResponseBody struct {
	// 开放消息任务ID
	OpenMsgTaskId *string `json:"openMsgTaskId,omitempty" xml:"openMsgTaskId,omitempty"`
}

func (s SendServiceGroupMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendServiceGroupMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendServiceGroupMessageResponseBody) SetOpenMsgTaskId(v string) *SendServiceGroupMessageResponseBody {
	s.OpenMsgTaskId = &v
	return s
}

type SendServiceGroupMessageResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendServiceGroupMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendServiceGroupMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendServiceGroupMessageResponse) GoString() string {
	return s.String()
}

func (s *SendServiceGroupMessageResponse) SetHeaders(v map[string]*string) *SendServiceGroupMessageResponse {
	s.Headers = v
	return s
}

func (s *SendServiceGroupMessageResponse) SetBody(v *SendServiceGroupMessageResponseBody) *SendServiceGroupMessageResponse {
	s.Body = v
	return s
}

type AddKnowledgeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddKnowledgeHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddKnowledgeHeaders) GoString() string {
	return s.String()
}

func (s *AddKnowledgeHeaders) SetCommonHeaders(v map[string]*string) *AddKnowledgeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddKnowledgeHeaders) SetXAcsDingtalkAccessToken(v string) *AddKnowledgeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddKnowledgeRequest struct {
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 知识库的唯一标识
	LibraryKey *string `json:"libraryKey,omitempty" xml:"libraryKey,omitempty"`
	// 知识点来源
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 知识点唯一标识
	SourcePrimaryKey *string `json:"sourcePrimaryKey,omitempty" xml:"sourcePrimaryKey,omitempty"`
	// 知识点类型 NORMAL：普通型 CARD：卡片 CONDITION：条件
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 知识点名称
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 知识点内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// CCM的知识点外链
	LinkUrl *string `json:"linkUrl,omitempty" xml:"linkUrl,omitempty"`
	// 知识点版本号
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s AddKnowledgeRequest) String() string {
	return tea.Prettify(s)
}

func (s AddKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *AddKnowledgeRequest) SetDingIsvOrgId(v int64) *AddKnowledgeRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *AddKnowledgeRequest) SetDingOrgId(v int64) *AddKnowledgeRequest {
	s.DingOrgId = &v
	return s
}

func (s *AddKnowledgeRequest) SetDingSuiteKey(v string) *AddKnowledgeRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *AddKnowledgeRequest) SetDingTokenGrantType(v int64) *AddKnowledgeRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *AddKnowledgeRequest) SetOpenTeamId(v string) *AddKnowledgeRequest {
	s.OpenTeamId = &v
	return s
}

func (s *AddKnowledgeRequest) SetLibraryKey(v string) *AddKnowledgeRequest {
	s.LibraryKey = &v
	return s
}

func (s *AddKnowledgeRequest) SetSource(v string) *AddKnowledgeRequest {
	s.Source = &v
	return s
}

func (s *AddKnowledgeRequest) SetSourcePrimaryKey(v string) *AddKnowledgeRequest {
	s.SourcePrimaryKey = &v
	return s
}

func (s *AddKnowledgeRequest) SetType(v string) *AddKnowledgeRequest {
	s.Type = &v
	return s
}

func (s *AddKnowledgeRequest) SetTitle(v string) *AddKnowledgeRequest {
	s.Title = &v
	return s
}

func (s *AddKnowledgeRequest) SetContent(v string) *AddKnowledgeRequest {
	s.Content = &v
	return s
}

func (s *AddKnowledgeRequest) SetLinkUrl(v string) *AddKnowledgeRequest {
	s.LinkUrl = &v
	return s
}

func (s *AddKnowledgeRequest) SetVersion(v string) *AddKnowledgeRequest {
	s.Version = &v
	return s
}

type AddKnowledgeResponseBody struct {
	// 开放知识点ID
	OpenKnowledgeId *string `json:"openKnowledgeId,omitempty" xml:"openKnowledgeId,omitempty"`
}

func (s AddKnowledgeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *AddKnowledgeResponseBody) SetOpenKnowledgeId(v string) *AddKnowledgeResponseBody {
	s.OpenKnowledgeId = &v
	return s
}

type AddKnowledgeResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddKnowledgeResponse) String() string {
	return tea.Prettify(s)
}

func (s AddKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *AddKnowledgeResponse) SetHeaders(v map[string]*string) *AddKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *AddKnowledgeResponse) SetBody(v *AddKnowledgeResponseBody) *AddKnowledgeResponse {
	s.Body = v
	return s
}

type QueryServiceGroupMessageReadStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryServiceGroupMessageReadStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceGroupMessageReadStatusHeaders) GoString() string {
	return s.String()
}

func (s *QueryServiceGroupMessageReadStatusHeaders) SetCommonHeaders(v map[string]*string) *QueryServiceGroupMessageReadStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryServiceGroupMessageReadStatusHeaders) SetXAcsDingtalkAccessToken(v string) *QueryServiceGroupMessageReadStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryServiceGroupMessageReadStatusRequest struct {
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 开放群ID
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 开放消息ID
	OpenMsgTaskId *string `json:"openMsgTaskId,omitempty" xml:"openMsgTaskId,omitempty"`
	// 用来标记当前开始读取的位置，置空表示从头开始。
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 本次读取的最大数据记录数量，此参数为可选参数，用户传入为空时，应该有默认值。应设置最大值限制，最大不超过100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
}

func (s QueryServiceGroupMessageReadStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceGroupMessageReadStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryServiceGroupMessageReadStatusRequest) SetDingIsvOrgId(v int64) *QueryServiceGroupMessageReadStatusRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusRequest) SetDingOrgId(v int64) *QueryServiceGroupMessageReadStatusRequest {
	s.DingOrgId = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusRequest) SetDingTokenGrantType(v int64) *QueryServiceGroupMessageReadStatusRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusRequest) SetDingSuiteKey(v string) *QueryServiceGroupMessageReadStatusRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusRequest) SetOpenTeamId(v string) *QueryServiceGroupMessageReadStatusRequest {
	s.OpenTeamId = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusRequest) SetOpenConversationId(v string) *QueryServiceGroupMessageReadStatusRequest {
	s.OpenConversationId = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusRequest) SetOpenMsgTaskId(v string) *QueryServiceGroupMessageReadStatusRequest {
	s.OpenMsgTaskId = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusRequest) SetNextToken(v string) *QueryServiceGroupMessageReadStatusRequest {
	s.NextToken = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusRequest) SetMaxResults(v int32) *QueryServiceGroupMessageReadStatusRequest {
	s.MaxResults = &v
	return s
}

type QueryServiceGroupMessageReadStatusResponseBody struct {
	// 本次请求条件下的数据总量，此参数为可选参数，默认可不返回。本次请求条件下的数据总量，此参数为可选参数，默认可不返回
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 本次请求所返回的最大记录条数。
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 已读未读信息列表
	Records []*QueryServiceGroupMessageReadStatusResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
}

func (s QueryServiceGroupMessageReadStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceGroupMessageReadStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryServiceGroupMessageReadStatusResponseBody) SetTotalCount(v int32) *QueryServiceGroupMessageReadStatusResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusResponseBody) SetNextToken(v string) *QueryServiceGroupMessageReadStatusResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusResponseBody) SetMaxResults(v int32) *QueryServiceGroupMessageReadStatusResponseBody {
	s.MaxResults = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusResponseBody) SetRecords(v []*QueryServiceGroupMessageReadStatusResponseBodyRecords) *QueryServiceGroupMessageReadStatusResponseBody {
	s.Records = v
	return s
}

type QueryServiceGroupMessageReadStatusResponseBodyRecords struct {
	// 已读人员为企业员工则有值
	ReceiverUserId *string `json:"receiverUserId,omitempty" xml:"receiverUserId,omitempty"`
	// 已读人员为非企业员工则有值
	ReceiverUnionId *string `json:"receiverUnionId,omitempty" xml:"receiverUnionId,omitempty"`
	// 状态：已读1/未读0
	ReadStatus *int32 `json:"readStatus,omitempty" xml:"readStatus,omitempty"`
}

func (s QueryServiceGroupMessageReadStatusResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceGroupMessageReadStatusResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *QueryServiceGroupMessageReadStatusResponseBodyRecords) SetReceiverUserId(v string) *QueryServiceGroupMessageReadStatusResponseBodyRecords {
	s.ReceiverUserId = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusResponseBodyRecords) SetReceiverUnionId(v string) *QueryServiceGroupMessageReadStatusResponseBodyRecords {
	s.ReceiverUnionId = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusResponseBodyRecords) SetReadStatus(v int32) *QueryServiceGroupMessageReadStatusResponseBodyRecords {
	s.ReadStatus = &v
	return s
}

type QueryServiceGroupMessageReadStatusResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryServiceGroupMessageReadStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryServiceGroupMessageReadStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceGroupMessageReadStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryServiceGroupMessageReadStatusResponse) SetHeaders(v map[string]*string) *QueryServiceGroupMessageReadStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryServiceGroupMessageReadStatusResponse) SetBody(v *QueryServiceGroupMessageReadStatusResponseBody) *QueryServiceGroupMessageReadStatusResponse {
	s.Body = v
	return s
}

type AddLibraryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddLibraryHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddLibraryHeaders) GoString() string {
	return s.String()
}

func (s *AddLibraryHeaders) SetCommonHeaders(v map[string]*string) *AddLibraryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddLibraryHeaders) SetXAcsDingtalkAccessToken(v string) *AddLibraryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddLibraryRequest struct {
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	// 团队id列表
	OpenTeamIds []*string `json:"openTeamIds,omitempty" xml:"openTeamIds,omitempty" type:"Repeated"`
	// 知识库名称
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 知识库描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 知识库类型 INTERNAL:内部知识库 EXTERNAL:外部知识库
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 知识来源
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 知识库的唯一性标识
	SourcePrimaryKey *string `json:"sourcePrimaryKey,omitempty" xml:"sourcePrimaryKey,omitempty"`
	// 员工ID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AddLibraryRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLibraryRequest) GoString() string {
	return s.String()
}

func (s *AddLibraryRequest) SetDingTokenGrantType(v int64) *AddLibraryRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *AddLibraryRequest) SetDingIsvOrgId(v int64) *AddLibraryRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *AddLibraryRequest) SetDingSuiteKey(v string) *AddLibraryRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *AddLibraryRequest) SetDingOrgId(v int64) *AddLibraryRequest {
	s.DingOrgId = &v
	return s
}

func (s *AddLibraryRequest) SetOpenTeamIds(v []*string) *AddLibraryRequest {
	s.OpenTeamIds = v
	return s
}

func (s *AddLibraryRequest) SetTitle(v string) *AddLibraryRequest {
	s.Title = &v
	return s
}

func (s *AddLibraryRequest) SetDescription(v string) *AddLibraryRequest {
	s.Description = &v
	return s
}

func (s *AddLibraryRequest) SetType(v string) *AddLibraryRequest {
	s.Type = &v
	return s
}

func (s *AddLibraryRequest) SetSource(v string) *AddLibraryRequest {
	s.Source = &v
	return s
}

func (s *AddLibraryRequest) SetSourcePrimaryKey(v string) *AddLibraryRequest {
	s.SourcePrimaryKey = &v
	return s
}

func (s *AddLibraryRequest) SetUserId(v string) *AddLibraryRequest {
	s.UserId = &v
	return s
}

type AddLibraryResponseBody struct {
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddLibraryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *AddLibraryResponseBody) SetSuccess(v bool) *AddLibraryResponseBody {
	s.Success = &v
	return s
}

type AddLibraryResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddLibraryResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLibraryResponse) GoString() string {
	return s.String()
}

func (s *AddLibraryResponse) SetHeaders(v map[string]*string) *AddLibraryResponse {
	s.Headers = v
	return s
}

func (s *AddLibraryResponse) SetBody(v *AddLibraryResponseBody) *AddLibraryResponse {
	s.Body = v
	return s
}

type ListUserTeamsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListUserTeamsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListUserTeamsHeaders) GoString() string {
	return s.String()
}

func (s *ListUserTeamsHeaders) SetCommonHeaders(v map[string]*string) *ListUserTeamsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListUserTeamsHeaders) SetXAcsDingtalkAccessToken(v string) *ListUserTeamsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListUserTeamsResponseBody struct {
	// teams
	Teams []*ListUserTeamsResponseBodyTeams `json:"teams,omitempty" xml:"teams,omitempty" type:"Repeated"`
}

func (s ListUserTeamsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserTeamsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserTeamsResponseBody) SetTeams(v []*ListUserTeamsResponseBodyTeams) *ListUserTeamsResponseBody {
	s.Teams = v
	return s
}

type ListUserTeamsResponseBodyTeams struct {
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 团队名称
	TeamName *string `json:"teamName,omitempty" xml:"teamName,omitempty"`
}

func (s ListUserTeamsResponseBodyTeams) String() string {
	return tea.Prettify(s)
}

func (s ListUserTeamsResponseBodyTeams) GoString() string {
	return s.String()
}

func (s *ListUserTeamsResponseBodyTeams) SetOpenTeamId(v string) *ListUserTeamsResponseBodyTeams {
	s.OpenTeamId = &v
	return s
}

func (s *ListUserTeamsResponseBodyTeams) SetTeamName(v string) *ListUserTeamsResponseBodyTeams {
	s.TeamName = &v
	return s
}

type ListUserTeamsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserTeamsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserTeamsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserTeamsResponse) GoString() string {
	return s.String()
}

func (s *ListUserTeamsResponse) SetHeaders(v map[string]*string) *ListUserTeamsResponse {
	s.Headers = v
	return s
}

func (s *ListUserTeamsResponse) SetBody(v *ListUserTeamsResponseBody) *ListUserTeamsResponse {
	s.Body = v
	return s
}

type QueryGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupHeaders) GoString() string {
	return s.String()
}

func (s *QueryGroupHeaders) SetCommonHeaders(v map[string]*string) *QueryGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGroupHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGroupRequest struct {
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	// 开放群ID
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s QueryGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryGroupRequest) SetDingIsvOrgId(v int64) *QueryGroupRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *QueryGroupRequest) SetDingOrgId(v int64) *QueryGroupRequest {
	s.DingOrgId = &v
	return s
}

func (s *QueryGroupRequest) SetDingSuiteKey(v string) *QueryGroupRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *QueryGroupRequest) SetDingTokenGrantType(v int64) *QueryGroupRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *QueryGroupRequest) SetOpenConversationId(v string) *QueryGroupRequest {
	s.OpenConversationId = &v
	return s
}

type QueryGroupResponseBody struct {
	// 开放群ID
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 群名称
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 开放群组ID
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// 入群URL
	GroupUrl *string `json:"groupUrl,omitempty" xml:"groupUrl,omitempty"`
	// 服务群机器人code
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	// 服务群机器人名称
	RobotName *string `json:"robotName,omitempty" xml:"robotName,omitempty"`
}

func (s QueryGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupResponseBody) SetOpenConversationId(v string) *QueryGroupResponseBody {
	s.OpenConversationId = &v
	return s
}

func (s *QueryGroupResponseBody) SetGroupName(v string) *QueryGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *QueryGroupResponseBody) SetOpenTeamId(v string) *QueryGroupResponseBody {
	s.OpenTeamId = &v
	return s
}

func (s *QueryGroupResponseBody) SetOpenGroupSetId(v string) *QueryGroupResponseBody {
	s.OpenGroupSetId = &v
	return s
}

func (s *QueryGroupResponseBody) SetGroupUrl(v string) *QueryGroupResponseBody {
	s.GroupUrl = &v
	return s
}

func (s *QueryGroupResponseBody) SetRobotCode(v string) *QueryGroupResponseBody {
	s.RobotCode = &v
	return s
}

func (s *QueryGroupResponseBody) SetRobotName(v string) *QueryGroupResponseBody {
	s.RobotName = &v
	return s
}

type QueryGroupResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupResponse) SetHeaders(v map[string]*string) *QueryGroupResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupResponse) SetBody(v *QueryGroupResponseBody) *QueryGroupResponse {
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

func (client *Client) DeleteKnowledge(request *DeleteKnowledgeRequest) (_result *DeleteKnowledgeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteKnowledgeHeaders{}
	_result = &DeleteKnowledgeResponse{}
	_body, _err := client.DeleteKnowledgeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteKnowledgeWithOptions(request *DeleteKnowledgeRequest, headers *DeleteKnowledgeHeaders, runtime *util.RuntimeOptions) (_result *DeleteKnowledgeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryKey)) {
		body["libraryKey"] = request.LibraryKey
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourcePrimaryKey)) {
		body["sourcePrimaryKey"] = request.SourcePrimaryKey
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
	_result = &DeleteKnowledgeResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteKnowledge"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/knowledges/batchDelete"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchGroup(request *SearchGroupRequest) (_result *SearchGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchGroupHeaders{}
	_result = &SearchGroupResponse{}
	_body, _err := client.SearchGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchGroupWithOptions(request *SearchGroupRequest, headers *SearchGroupHeaders, runtime *util.RuntimeOptions) (_result *SearchGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenGroupSetId)) {
		body["openGroupSetId"] = request.OpenGroupSetId
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
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
	_result = &SearchGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("SearchGroup"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/groups/search"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGroup(request *CreateGroupRequest) (_result *CreateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateGroupHeaders{}
	_result = &CreateGroupResponse{}
	_body, _err := client.CreateGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGroupWithOptions(request *CreateGroupRequest, headers *CreateGroupHeaders, runtime *util.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupBizId)) {
		body["groupBizId"] = request.GroupBizId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenGroupSetId)) {
		body["openGroupSetId"] = request.OpenGroupSetId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerStaffId)) {
		body["ownerStaffId"] = request.OwnerStaffId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberStaffIds)) {
		body["memberStaffIds"] = request.MemberStaffIds
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
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
	_result = &CreateGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateGroup"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/groups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendServiceGroupMessage(request *SendServiceGroupMessageRequest) (_result *SendServiceGroupMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendServiceGroupMessageHeaders{}
	_result = &SendServiceGroupMessageResponse{}
	_body, _err := client.SendServiceGroupMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendServiceGroupMessageWithOptions(request *SendServiceGroupMessageRequest, headers *SendServiceGroupMessageHeaders, runtime *util.RuntimeOptions) (_result *SendServiceGroupMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.TargetOpenConversationId)) {
		body["targetOpenConversationId"] = request.TargetOpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.IsAtAll)) {
		body["isAtAll"] = request.IsAtAll
	}

	if !tea.BoolValue(util.IsUnset(request.AtMobiles)) {
		body["atMobiles"] = request.AtMobiles
	}

	if !tea.BoolValue(util.IsUnset(request.AtDingtalkIds)) {
		body["atDingtalkIds"] = request.AtDingtalkIds
	}

	if !tea.BoolValue(util.IsUnset(request.AtUnionIds)) {
		body["atUnionIds"] = request.AtUnionIds
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverMobiles)) {
		body["receiverMobiles"] = request.ReceiverMobiles
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverDingtalkIds)) {
		body["receiverDingtalkIds"] = request.ReceiverDingtalkIds
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUnionIds)) {
		body["receiverUnionIds"] = request.ReceiverUnionIds
	}

	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		body["messageType"] = request.MessageType
	}

	if !tea.BoolValue(util.IsUnset(request.BtnOrientation)) {
		body["btnOrientation"] = request.BtnOrientation
	}

	if !tea.BoolValue(util.IsUnset(request.Btns)) {
		body["btns"] = request.Btns
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
	_result = &SendServiceGroupMessageResponse{}
	_body, _err := client.DoROARequest(tea.String("SendServiceGroupMessage"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/messages/send"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddKnowledge(request *AddKnowledgeRequest) (_result *AddKnowledgeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddKnowledgeHeaders{}
	_result = &AddKnowledgeResponse{}
	_body, _err := client.AddKnowledgeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddKnowledgeWithOptions(request *AddKnowledgeRequest, headers *AddKnowledgeHeaders, runtime *util.RuntimeOptions) (_result *AddKnowledgeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryKey)) {
		body["libraryKey"] = request.LibraryKey
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourcePrimaryKey)) {
		body["sourcePrimaryKey"] = request.SourcePrimaryKey
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.LinkUrl)) {
		body["linkUrl"] = request.LinkUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["version"] = request.Version
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
	_result = &AddKnowledgeResponse{}
	_body, _err := client.DoROARequest(tea.String("AddKnowledge"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/knowledges"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryServiceGroupMessageReadStatus(request *QueryServiceGroupMessageReadStatusRequest) (_result *QueryServiceGroupMessageReadStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryServiceGroupMessageReadStatusHeaders{}
	_result = &QueryServiceGroupMessageReadStatusResponse{}
	_body, _err := client.QueryServiceGroupMessageReadStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryServiceGroupMessageReadStatusWithOptions(request *QueryServiceGroupMessageReadStatusRequest, headers *QueryServiceGroupMessageReadStatusHeaders, runtime *util.RuntimeOptions) (_result *QueryServiceGroupMessageReadStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenMsgTaskId)) {
		body["openMsgTaskId"] = request.OpenMsgTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
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
	_result = &QueryServiceGroupMessageReadStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryServiceGroupMessageReadStatus"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/messages/readStatus/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddLibrary(request *AddLibraryRequest) (_result *AddLibraryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddLibraryHeaders{}
	_result = &AddLibraryResponse{}
	_body, _err := client.AddLibraryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddLibraryWithOptions(request *AddLibraryRequest, headers *AddLibraryHeaders, runtime *util.RuntimeOptions) (_result *AddLibraryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamIds)) {
		body["openTeamIds"] = request.OpenTeamIds
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourcePrimaryKey)) {
		body["sourcePrimaryKey"] = request.SourcePrimaryKey
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
	_result = &AddLibraryResponse{}
	_body, _err := client.DoROARequest(tea.String("AddLibrary"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/librarys"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserTeams(userId *string) (_result *ListUserTeamsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListUserTeamsHeaders{}
	_result = &ListUserTeamsResponse{}
	_body, _err := client.ListUserTeamsWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserTeamsWithOptions(userId *string, headers *ListUserTeamsHeaders, runtime *util.RuntimeOptions) (_result *ListUserTeamsResponse, _err error) {
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
	_result = &ListUserTeamsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListUserTeams"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/serviceGroup/users/"+tea.StringValue(userId)+"/teams"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryGroup(request *QueryGroupRequest) (_result *QueryGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGroupHeaders{}
	_result = &QueryGroupResponse{}
	_body, _err := client.QueryGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryGroupWithOptions(request *QueryGroupRequest, headers *QueryGroupHeaders, runtime *util.RuntimeOptions) (_result *QueryGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
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
	_result = &QueryGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryGroup"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/groups/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
