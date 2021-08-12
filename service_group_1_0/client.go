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
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 工单创建人UnionId
	CreatorUnionId *string `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	// 工单处理人UnionId列表
	ProcessorUnionIds []*string `json:"processorUnionIds,omitempty" xml:"processorUnionIds,omitempty" type:"Repeated"`
	// 工单场景 SG 或 VOC
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// 工单场景信息
	SceneContext *CreateTicketRequestSceneContext `json:"sceneContext,omitempty" xml:"sceneContext,omitempty" type:"Struct"`
	// 工单模板业务ID
	OpenTemplateBizId *string `json:"openTemplateBizId,omitempty" xml:"openTemplateBizId,omitempty"`
	// 工单标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 自定义组件字段值(JSON格式)
	CustomFields *string `json:"customFields,omitempty" xml:"customFields,omitempty"`
	// 通知接收人配置
	Notify *CreateTicketRequestNotify `json:"notify,omitempty" xml:"notify,omitempty" type:"Struct"`
}

func (s CreateTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketRequest) GoString() string {
	return s.String()
}

func (s *CreateTicketRequest) SetDingIsvOrgId(v int64) *CreateTicketRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *CreateTicketRequest) SetDingOrgId(v int64) *CreateTicketRequest {
	s.DingOrgId = &v
	return s
}

func (s *CreateTicketRequest) SetDingSuiteKey(v string) *CreateTicketRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *CreateTicketRequest) SetDingTokenGrantType(v int64) *CreateTicketRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *CreateTicketRequest) SetOpenTeamId(v string) *CreateTicketRequest {
	s.OpenTeamId = &v
	return s
}

func (s *CreateTicketRequest) SetCreatorUnionId(v string) *CreateTicketRequest {
	s.CreatorUnionId = &v
	return s
}

func (s *CreateTicketRequest) SetProcessorUnionIds(v []*string) *CreateTicketRequest {
	s.ProcessorUnionIds = v
	return s
}

func (s *CreateTicketRequest) SetScene(v string) *CreateTicketRequest {
	s.Scene = &v
	return s
}

func (s *CreateTicketRequest) SetSceneContext(v *CreateTicketRequestSceneContext) *CreateTicketRequest {
	s.SceneContext = v
	return s
}

func (s *CreateTicketRequest) SetOpenTemplateBizId(v string) *CreateTicketRequest {
	s.OpenTemplateBizId = &v
	return s
}

func (s *CreateTicketRequest) SetTitle(v string) *CreateTicketRequest {
	s.Title = &v
	return s
}

func (s *CreateTicketRequest) SetCustomFields(v string) *CreateTicketRequest {
	s.CustomFields = &v
	return s
}

func (s *CreateTicketRequest) SetNotify(v *CreateTicketRequestNotify) *CreateTicketRequest {
	s.Notify = v
	return s
}

type CreateTicketRequestSceneContext struct {
	// 服务群openConversationId
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 工单相关人UnionId列表
	RelevantorUnionIds []*string `json:"relevantorUnionIds,omitempty" xml:"relevantorUnionIds,omitempty" type:"Repeated"`
	// 工单相关的群消息列表
	GroupMsgs []*CreateTicketRequestSceneContextGroupMsgs `json:"groupMsgs,omitempty" xml:"groupMsgs,omitempty" type:"Repeated"`
	// VOC类型工单，对应话题ID
	TopicId *string `json:"topicId,omitempty" xml:"topicId,omitempty"`
}

func (s CreateTicketRequestSceneContext) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketRequestSceneContext) GoString() string {
	return s.String()
}

func (s *CreateTicketRequestSceneContext) SetOpenConversationId(v string) *CreateTicketRequestSceneContext {
	s.OpenConversationId = &v
	return s
}

func (s *CreateTicketRequestSceneContext) SetRelevantorUnionIds(v []*string) *CreateTicketRequestSceneContext {
	s.RelevantorUnionIds = v
	return s
}

func (s *CreateTicketRequestSceneContext) SetGroupMsgs(v []*CreateTicketRequestSceneContextGroupMsgs) *CreateTicketRequestSceneContext {
	s.GroupMsgs = v
	return s
}

func (s *CreateTicketRequestSceneContext) SetTopicId(v string) *CreateTicketRequestSceneContext {
	s.TopicId = &v
	return s
}

type CreateTicketRequestSceneContextGroupMsgs struct {
	// 勾选消息openMsgId
	OpenMsgId *string `json:"openMsgId,omitempty" xml:"openMsgId,omitempty"`
	// 是否为锚点消息
	Anchor *bool `json:"anchor,omitempty" xml:"anchor,omitempty"`
}

func (s CreateTicketRequestSceneContextGroupMsgs) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketRequestSceneContextGroupMsgs) GoString() string {
	return s.String()
}

func (s *CreateTicketRequestSceneContextGroupMsgs) SetOpenMsgId(v string) *CreateTicketRequestSceneContextGroupMsgs {
	s.OpenMsgId = &v
	return s
}

func (s *CreateTicketRequestSceneContextGroupMsgs) SetAnchor(v bool) *CreateTicketRequestSceneContextGroupMsgs {
	s.Anchor = &v
	return s
}

type CreateTicketRequestNotify struct {
	// 企业工作通知接收人（钉钉UnionId）
	WorkNoticeReceiverUnionIds []*string `json:"workNoticeReceiverUnionIds,omitempty" xml:"workNoticeReceiverUnionIds,omitempty" type:"Repeated"`
	// 服务群通知接收人（钉钉UnionId）
	GroupNoticeReceiverUnionIds []*string `json:"groupNoticeReceiverUnionIds,omitempty" xml:"groupNoticeReceiverUnionIds,omitempty" type:"Repeated"`
	// 是否向群内推送一个全员可见工单通知卡片
	NoticeAllGroupMember *bool `json:"noticeAllGroupMember,omitempty" xml:"noticeAllGroupMember,omitempty"`
}

func (s CreateTicketRequestNotify) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketRequestNotify) GoString() string {
	return s.String()
}

func (s *CreateTicketRequestNotify) SetWorkNoticeReceiverUnionIds(v []*string) *CreateTicketRequestNotify {
	s.WorkNoticeReceiverUnionIds = v
	return s
}

func (s *CreateTicketRequestNotify) SetGroupNoticeReceiverUnionIds(v []*string) *CreateTicketRequestNotify {
	s.GroupNoticeReceiverUnionIds = v
	return s
}

func (s *CreateTicketRequestNotify) SetNoticeAllGroupMember(v bool) *CreateTicketRequestNotify {
	s.NoticeAllGroupMember = &v
	return s
}

type CreateTicketResponseBody struct {
	// 工单开放ID
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
}

func (s CreateTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTicketResponseBody) SetOpenTicketId(v string) *CreateTicketResponseBody {
	s.OpenTicketId = &v
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

type AssignTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AssignTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s AssignTicketHeaders) GoString() string {
	return s.String()
}

func (s *AssignTicketHeaders) SetCommonHeaders(v map[string]*string) *AssignTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AssignTicketHeaders) SetXAcsDingtalkAccessToken(v string) *AssignTicketHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AssignTicketRequest struct {
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	// 操作人unionId（管理员）
	OperatorUnionId *string `json:"operatorUnionId,omitempty" xml:"operatorUnionId,omitempty"`
	// 工单开放ID
	OpenTicketId      *string   `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	ProcessorUnionIds []*string `json:"processorUnionIds,omitempty" xml:"processorUnionIds,omitempty" type:"Repeated"`
	// 备注
	TicketMemo *AssignTicketRequestTicketMemo `json:"ticketMemo,omitempty" xml:"ticketMemo,omitempty" type:"Struct"`
	Notify     *AssignTicketRequestNotify     `json:"notify,omitempty" xml:"notify,omitempty" type:"Struct"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s AssignTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s AssignTicketRequest) GoString() string {
	return s.String()
}

func (s *AssignTicketRequest) SetDingIsvOrgId(v int64) *AssignTicketRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *AssignTicketRequest) SetDingOrgId(v int64) *AssignTicketRequest {
	s.DingOrgId = &v
	return s
}

func (s *AssignTicketRequest) SetDingSuiteKey(v string) *AssignTicketRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *AssignTicketRequest) SetDingTokenGrantType(v int64) *AssignTicketRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *AssignTicketRequest) SetOperatorUnionId(v string) *AssignTicketRequest {
	s.OperatorUnionId = &v
	return s
}

func (s *AssignTicketRequest) SetOpenTicketId(v string) *AssignTicketRequest {
	s.OpenTicketId = &v
	return s
}

func (s *AssignTicketRequest) SetProcessorUnionIds(v []*string) *AssignTicketRequest {
	s.ProcessorUnionIds = v
	return s
}

func (s *AssignTicketRequest) SetTicketMemo(v *AssignTicketRequestTicketMemo) *AssignTicketRequest {
	s.TicketMemo = v
	return s
}

func (s *AssignTicketRequest) SetNotify(v *AssignTicketRequestNotify) *AssignTicketRequest {
	s.Notify = v
	return s
}

func (s *AssignTicketRequest) SetOpenTeamId(v string) *AssignTicketRequest {
	s.OpenTeamId = &v
	return s
}

type AssignTicketRequestTicketMemo struct {
	// 备注文字
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 备注相关的附件
	Attachments []*AssignTicketRequestTicketMemoAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
}

func (s AssignTicketRequestTicketMemo) String() string {
	return tea.Prettify(s)
}

func (s AssignTicketRequestTicketMemo) GoString() string {
	return s.String()
}

func (s *AssignTicketRequestTicketMemo) SetMemo(v string) *AssignTicketRequestTicketMemo {
	s.Memo = &v
	return s
}

func (s *AssignTicketRequestTicketMemo) SetAttachments(v []*AssignTicketRequestTicketMemoAttachments) *AssignTicketRequestTicketMemo {
	s.Attachments = v
	return s
}

type AssignTicketRequestTicketMemoAttachments struct {
	// 文件名
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s AssignTicketRequestTicketMemoAttachments) String() string {
	return tea.Prettify(s)
}

func (s AssignTicketRequestTicketMemoAttachments) GoString() string {
	return s.String()
}

func (s *AssignTicketRequestTicketMemoAttachments) SetFileName(v string) *AssignTicketRequestTicketMemoAttachments {
	s.FileName = &v
	return s
}

func (s *AssignTicketRequestTicketMemoAttachments) SetKey(v string) *AssignTicketRequestTicketMemoAttachments {
	s.Key = &v
	return s
}

type AssignTicketRequestNotify struct {
	WorkNoticeReceiverUnionIds  []*string `json:"workNoticeReceiverUnionIds,omitempty" xml:"workNoticeReceiverUnionIds,omitempty" type:"Repeated"`
	GroupNoticeReceiverUnionIds []*string `json:"groupNoticeReceiverUnionIds,omitempty" xml:"groupNoticeReceiverUnionIds,omitempty" type:"Repeated"`
	// 是否向群内推送一个全员可见工单通知卡片
	NoticeAllGroupMember *bool `json:"noticeAllGroupMember,omitempty" xml:"noticeAllGroupMember,omitempty"`
}

func (s AssignTicketRequestNotify) String() string {
	return tea.Prettify(s)
}

func (s AssignTicketRequestNotify) GoString() string {
	return s.String()
}

func (s *AssignTicketRequestNotify) SetWorkNoticeReceiverUnionIds(v []*string) *AssignTicketRequestNotify {
	s.WorkNoticeReceiverUnionIds = v
	return s
}

func (s *AssignTicketRequestNotify) SetGroupNoticeReceiverUnionIds(v []*string) *AssignTicketRequestNotify {
	s.GroupNoticeReceiverUnionIds = v
	return s
}

func (s *AssignTicketRequestNotify) SetNoticeAllGroupMember(v bool) *AssignTicketRequestNotify {
	s.NoticeAllGroupMember = &v
	return s
}

type AssignTicketResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s AssignTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s AssignTicketResponse) GoString() string {
	return s.String()
}

func (s *AssignTicketResponse) SetHeaders(v map[string]*string) *AssignTicketResponse {
	s.Headers = v
	return s
}

type FinishTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s FinishTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s FinishTicketHeaders) GoString() string {
	return s.String()
}

func (s *FinishTicketHeaders) SetCommonHeaders(v map[string]*string) *FinishTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FinishTicketHeaders) SetXAcsDingtalkAccessToken(v string) *FinishTicketHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type FinishTicketRequest struct {
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	OpenTeamId         *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 当前工单处理人
	ProcessorUnionId *string `json:"processorUnionId,omitempty" xml:"processorUnionId,omitempty"`
	// 工单开放id
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	// 备注
	TicketMemo *FinishTicketRequestTicketMemo `json:"ticketMemo,omitempty" xml:"ticketMemo,omitempty" type:"Struct"`
	// 工单通知
	Notify *FinishTicketRequestNotify `json:"notify,omitempty" xml:"notify,omitempty" type:"Struct"`
}

func (s FinishTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s FinishTicketRequest) GoString() string {
	return s.String()
}

func (s *FinishTicketRequest) SetDingIsvOrgId(v int64) *FinishTicketRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *FinishTicketRequest) SetDingOrgId(v int64) *FinishTicketRequest {
	s.DingOrgId = &v
	return s
}

func (s *FinishTicketRequest) SetDingSuiteKey(v string) *FinishTicketRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *FinishTicketRequest) SetDingTokenGrantType(v int64) *FinishTicketRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *FinishTicketRequest) SetOpenTeamId(v string) *FinishTicketRequest {
	s.OpenTeamId = &v
	return s
}

func (s *FinishTicketRequest) SetProcessorUnionId(v string) *FinishTicketRequest {
	s.ProcessorUnionId = &v
	return s
}

func (s *FinishTicketRequest) SetOpenTicketId(v string) *FinishTicketRequest {
	s.OpenTicketId = &v
	return s
}

func (s *FinishTicketRequest) SetTicketMemo(v *FinishTicketRequestTicketMemo) *FinishTicketRequest {
	s.TicketMemo = v
	return s
}

func (s *FinishTicketRequest) SetNotify(v *FinishTicketRequestNotify) *FinishTicketRequest {
	s.Notify = v
	return s
}

type FinishTicketRequestTicketMemo struct {
	// 备注文字
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 备注相关的附件
	Attachments []*FinishTicketRequestTicketMemoAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
}

func (s FinishTicketRequestTicketMemo) String() string {
	return tea.Prettify(s)
}

func (s FinishTicketRequestTicketMemo) GoString() string {
	return s.String()
}

func (s *FinishTicketRequestTicketMemo) SetMemo(v string) *FinishTicketRequestTicketMemo {
	s.Memo = &v
	return s
}

func (s *FinishTicketRequestTicketMemo) SetAttachments(v []*FinishTicketRequestTicketMemoAttachments) *FinishTicketRequestTicketMemo {
	s.Attachments = v
	return s
}

type FinishTicketRequestTicketMemoAttachments struct {
	// 文件名
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s FinishTicketRequestTicketMemoAttachments) String() string {
	return tea.Prettify(s)
}

func (s FinishTicketRequestTicketMemoAttachments) GoString() string {
	return s.String()
}

func (s *FinishTicketRequestTicketMemoAttachments) SetFileName(v string) *FinishTicketRequestTicketMemoAttachments {
	s.FileName = &v
	return s
}

func (s *FinishTicketRequestTicketMemoAttachments) SetKey(v string) *FinishTicketRequestTicketMemoAttachments {
	s.Key = &v
	return s
}

type FinishTicketRequestNotify struct {
	// 企业工作通知接收人（钉钉UnionId）
	WorkNoticeReceiverUnionIds []*string `json:"workNoticeReceiverUnionIds,omitempty" xml:"workNoticeReceiverUnionIds,omitempty" type:"Repeated"`
	// 群中通知接收人（钉钉UnionId）
	GroupNoticeReceiverUnionIds []*string `json:"groupNoticeReceiverUnionIds,omitempty" xml:"groupNoticeReceiverUnionIds,omitempty" type:"Repeated"`
	// 是否向群内推送一个全员可见工单通知卡片
	NoticeAllGroupMember *bool `json:"noticeAllGroupMember,omitempty" xml:"noticeAllGroupMember,omitempty"`
}

func (s FinishTicketRequestNotify) String() string {
	return tea.Prettify(s)
}

func (s FinishTicketRequestNotify) GoString() string {
	return s.String()
}

func (s *FinishTicketRequestNotify) SetWorkNoticeReceiverUnionIds(v []*string) *FinishTicketRequestNotify {
	s.WorkNoticeReceiverUnionIds = v
	return s
}

func (s *FinishTicketRequestNotify) SetGroupNoticeReceiverUnionIds(v []*string) *FinishTicketRequestNotify {
	s.GroupNoticeReceiverUnionIds = v
	return s
}

func (s *FinishTicketRequestNotify) SetNoticeAllGroupMember(v bool) *FinishTicketRequestNotify {
	s.NoticeAllGroupMember = &v
	return s
}

type FinishTicketResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s FinishTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s FinishTicketResponse) GoString() string {
	return s.String()
}

func (s *FinishTicketResponse) SetHeaders(v map[string]*string) *FinishTicketResponse {
	s.Headers = v
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

type UpdateTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTicketHeaders) SetCommonHeaders(v map[string]*string) *UpdateTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTicketHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateTicketHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateTicketRequest struct {
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	// 团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 工单处理人unionId
	ProcessorUnionId *string `json:"processorUnionId,omitempty" xml:"processorUnionId,omitempty"`
	// 工单开放id
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	// 自定义字段值JSON格式
	CustomFields *string `json:"customFields,omitempty" xml:"customFields,omitempty"`
	// 备注
	TicketMemo *UpdateTicketRequestTicketMemo `json:"ticketMemo,omitempty" xml:"ticketMemo,omitempty" type:"Struct"`
}

func (s UpdateTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketRequest) GoString() string {
	return s.String()
}

func (s *UpdateTicketRequest) SetDingIsvOrgId(v int64) *UpdateTicketRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *UpdateTicketRequest) SetDingOrgId(v int64) *UpdateTicketRequest {
	s.DingOrgId = &v
	return s
}

func (s *UpdateTicketRequest) SetDingSuiteKey(v string) *UpdateTicketRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *UpdateTicketRequest) SetDingTokenGrantType(v int64) *UpdateTicketRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *UpdateTicketRequest) SetOpenTeamId(v string) *UpdateTicketRequest {
	s.OpenTeamId = &v
	return s
}

func (s *UpdateTicketRequest) SetProcessorUnionId(v string) *UpdateTicketRequest {
	s.ProcessorUnionId = &v
	return s
}

func (s *UpdateTicketRequest) SetOpenTicketId(v string) *UpdateTicketRequest {
	s.OpenTicketId = &v
	return s
}

func (s *UpdateTicketRequest) SetCustomFields(v string) *UpdateTicketRequest {
	s.CustomFields = &v
	return s
}

func (s *UpdateTicketRequest) SetTicketMemo(v *UpdateTicketRequestTicketMemo) *UpdateTicketRequest {
	s.TicketMemo = v
	return s
}

type UpdateTicketRequestTicketMemo struct {
	// 备注文字
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 备注相关的附件
	Attachments []*UpdateTicketRequestTicketMemoAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
}

func (s UpdateTicketRequestTicketMemo) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketRequestTicketMemo) GoString() string {
	return s.String()
}

func (s *UpdateTicketRequestTicketMemo) SetMemo(v string) *UpdateTicketRequestTicketMemo {
	s.Memo = &v
	return s
}

func (s *UpdateTicketRequestTicketMemo) SetAttachments(v []*UpdateTicketRequestTicketMemoAttachments) *UpdateTicketRequestTicketMemo {
	s.Attachments = v
	return s
}

type UpdateTicketRequestTicketMemoAttachments struct {
	// 文件名
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s UpdateTicketRequestTicketMemoAttachments) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketRequestTicketMemoAttachments) GoString() string {
	return s.String()
}

func (s *UpdateTicketRequestTicketMemoAttachments) SetFileName(v string) *UpdateTicketRequestTicketMemoAttachments {
	s.FileName = &v
	return s
}

func (s *UpdateTicketRequestTicketMemoAttachments) SetKey(v string) *UpdateTicketRequestTicketMemoAttachments {
	s.Key = &v
	return s
}

type UpdateTicketResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpdateTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketResponse) GoString() string {
	return s.String()
}

func (s *UpdateTicketResponse) SetHeaders(v map[string]*string) *UpdateTicketResponse {
	s.Headers = v
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
	MemberStaffIds []*string `json:"memberStaffIds,omitempty" xml:"memberStaffIds,omitempty" type:"Repeated"`
	// 群标签
	GroupTagNames      []*string `json:"groupTagNames,omitempty" xml:"groupTagNames,omitempty" type:"Repeated"`
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

func (s *CreateGroupRequest) SetGroupTagNames(v []*string) *CreateGroupRequest {
	s.GroupTagNames = v
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

type TransferTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TransferTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s TransferTicketHeaders) GoString() string {
	return s.String()
}

func (s *TransferTicketHeaders) SetCommonHeaders(v map[string]*string) *TransferTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TransferTicketHeaders) SetXAcsDingtalkAccessToken(v string) *TransferTicketHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TransferTicketRequest struct {
	// 工单处理人
	ProcessorUnionId *string `json:"processorUnionId,omitempty" xml:"processorUnionId,omitempty"`
	// 工单开放ID
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	// 被转单人UnionId列表
	ProcessorUnionIds []*string `json:"processorUnionIds,omitempty" xml:"processorUnionIds,omitempty" type:"Repeated"`
	// 工单备注
	TicketMemo *TransferTicketRequestTicketMemo `json:"ticketMemo,omitempty" xml:"ticketMemo,omitempty" type:"Struct"`
	Notify     *TransferTicketRequestNotify     `json:"notify,omitempty" xml:"notify,omitempty" type:"Struct"`
	// 开放团队ID
	OpenTeamId         *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
}

func (s TransferTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferTicketRequest) GoString() string {
	return s.String()
}

func (s *TransferTicketRequest) SetProcessorUnionId(v string) *TransferTicketRequest {
	s.ProcessorUnionId = &v
	return s
}

func (s *TransferTicketRequest) SetOpenTicketId(v string) *TransferTicketRequest {
	s.OpenTicketId = &v
	return s
}

func (s *TransferTicketRequest) SetProcessorUnionIds(v []*string) *TransferTicketRequest {
	s.ProcessorUnionIds = v
	return s
}

func (s *TransferTicketRequest) SetTicketMemo(v *TransferTicketRequestTicketMemo) *TransferTicketRequest {
	s.TicketMemo = v
	return s
}

func (s *TransferTicketRequest) SetNotify(v *TransferTicketRequestNotify) *TransferTicketRequest {
	s.Notify = v
	return s
}

func (s *TransferTicketRequest) SetOpenTeamId(v string) *TransferTicketRequest {
	s.OpenTeamId = &v
	return s
}

func (s *TransferTicketRequest) SetDingIsvOrgId(v int64) *TransferTicketRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *TransferTicketRequest) SetDingOrgId(v int64) *TransferTicketRequest {
	s.DingOrgId = &v
	return s
}

func (s *TransferTicketRequest) SetDingSuiteKey(v string) *TransferTicketRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *TransferTicketRequest) SetDingTokenGrantType(v int64) *TransferTicketRequest {
	s.DingTokenGrantType = &v
	return s
}

type TransferTicketRequestTicketMemo struct {
	// 文字备注
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 备注相关的附件
	Attachments []*TransferTicketRequestTicketMemoAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
}

func (s TransferTicketRequestTicketMemo) String() string {
	return tea.Prettify(s)
}

func (s TransferTicketRequestTicketMemo) GoString() string {
	return s.String()
}

func (s *TransferTicketRequestTicketMemo) SetMemo(v string) *TransferTicketRequestTicketMemo {
	s.Memo = &v
	return s
}

func (s *TransferTicketRequestTicketMemo) SetAttachments(v []*TransferTicketRequestTicketMemoAttachments) *TransferTicketRequestTicketMemo {
	s.Attachments = v
	return s
}

type TransferTicketRequestTicketMemoAttachments struct {
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	Key      *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s TransferTicketRequestTicketMemoAttachments) String() string {
	return tea.Prettify(s)
}

func (s TransferTicketRequestTicketMemoAttachments) GoString() string {
	return s.String()
}

func (s *TransferTicketRequestTicketMemoAttachments) SetFileName(v string) *TransferTicketRequestTicketMemoAttachments {
	s.FileName = &v
	return s
}

func (s *TransferTicketRequestTicketMemoAttachments) SetKey(v string) *TransferTicketRequestTicketMemoAttachments {
	s.Key = &v
	return s
}

type TransferTicketRequestNotify struct {
	// 企业工作通知接收人（钉钉UnionId）
	WorkNoticeReceiverUnionIds []*string `json:"workNoticeReceiverUnionIds,omitempty" xml:"workNoticeReceiverUnionIds,omitempty" type:"Repeated"`
	// 群中通知接收人（钉钉UnionId）
	GroupNoticeReceiverUnionIds []*string `json:"groupNoticeReceiverUnionIds,omitempty" xml:"groupNoticeReceiverUnionIds,omitempty" type:"Repeated"`
	// 是否向群内推送一个全员可见工单通知卡片
	NoticeAllGroupMember *bool `json:"noticeAllGroupMember,omitempty" xml:"noticeAllGroupMember,omitempty"`
}

func (s TransferTicketRequestNotify) String() string {
	return tea.Prettify(s)
}

func (s TransferTicketRequestNotify) GoString() string {
	return s.String()
}

func (s *TransferTicketRequestNotify) SetWorkNoticeReceiverUnionIds(v []*string) *TransferTicketRequestNotify {
	s.WorkNoticeReceiverUnionIds = v
	return s
}

func (s *TransferTicketRequestNotify) SetGroupNoticeReceiverUnionIds(v []*string) *TransferTicketRequestNotify {
	s.GroupNoticeReceiverUnionIds = v
	return s
}

func (s *TransferTicketRequestNotify) SetNoticeAllGroupMember(v bool) *TransferTicketRequestNotify {
	s.NoticeAllGroupMember = &v
	return s
}

type TransferTicketResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s TransferTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferTicketResponse) GoString() string {
	return s.String()
}

func (s *TransferTicketResponse) SetHeaders(v map[string]*string) *TransferTicketResponse {
	s.Headers = v
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

type BatchGetGroupSetConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchGetGroupSetConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchGetGroupSetConfigHeaders) GoString() string {
	return s.String()
}

func (s *BatchGetGroupSetConfigHeaders) SetCommonHeaders(v map[string]*string) *BatchGetGroupSetConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchGetGroupSetConfigHeaders) SetXAcsDingtalkAccessToken(v string) *BatchGetGroupSetConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchGetGroupSetConfigRequest struct {
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	// 开放团队id
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 开放群组id
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// 配置项key列表
	ConfigKeys []*string `json:"configKeys,omitempty" xml:"configKeys,omitempty" type:"Repeated"`
}

func (s BatchGetGroupSetConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetGroupSetConfigRequest) GoString() string {
	return s.String()
}

func (s *BatchGetGroupSetConfigRequest) SetDingIsvOrgId(v int64) *BatchGetGroupSetConfigRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *BatchGetGroupSetConfigRequest) SetDingOrgId(v int64) *BatchGetGroupSetConfigRequest {
	s.DingOrgId = &v
	return s
}

func (s *BatchGetGroupSetConfigRequest) SetDingSuiteKey(v string) *BatchGetGroupSetConfigRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *BatchGetGroupSetConfigRequest) SetDingTokenGrantType(v int64) *BatchGetGroupSetConfigRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *BatchGetGroupSetConfigRequest) SetOpenTeamId(v string) *BatchGetGroupSetConfigRequest {
	s.OpenTeamId = &v
	return s
}

func (s *BatchGetGroupSetConfigRequest) SetOpenGroupSetId(v string) *BatchGetGroupSetConfigRequest {
	s.OpenGroupSetId = &v
	return s
}

func (s *BatchGetGroupSetConfigRequest) SetConfigKeys(v []*string) *BatchGetGroupSetConfigRequest {
	s.ConfigKeys = v
	return s
}

type BatchGetGroupSetConfigResponseBody struct {
	// 群粗配置列表
	GroupSetConfigs []*BatchGetGroupSetConfigResponseBodyGroupSetConfigs `json:"groupSetConfigs,omitempty" xml:"groupSetConfigs,omitempty" type:"Repeated"`
}

func (s BatchGetGroupSetConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchGetGroupSetConfigResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetGroupSetConfigResponseBody) SetGroupSetConfigs(v []*BatchGetGroupSetConfigResponseBodyGroupSetConfigs) *BatchGetGroupSetConfigResponseBody {
	s.GroupSetConfigs = v
	return s
}

type BatchGetGroupSetConfigResponseBodyGroupSetConfigs struct {
	// 配置项key
	ConfigKey *string `json:"configKey,omitempty" xml:"configKey,omitempty"`
	// 配置项值
	ConfigValue *string `json:"configValue,omitempty" xml:"configValue,omitempty"`
}

func (s BatchGetGroupSetConfigResponseBodyGroupSetConfigs) String() string {
	return tea.Prettify(s)
}

func (s BatchGetGroupSetConfigResponseBodyGroupSetConfigs) GoString() string {
	return s.String()
}

func (s *BatchGetGroupSetConfigResponseBodyGroupSetConfigs) SetConfigKey(v string) *BatchGetGroupSetConfigResponseBodyGroupSetConfigs {
	s.ConfigKey = &v
	return s
}

func (s *BatchGetGroupSetConfigResponseBodyGroupSetConfigs) SetConfigValue(v string) *BatchGetGroupSetConfigResponseBodyGroupSetConfigs {
	s.ConfigValue = &v
	return s
}

type BatchGetGroupSetConfigResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchGetGroupSetConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchGetGroupSetConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchGetGroupSetConfigResponse) GoString() string {
	return s.String()
}

func (s *BatchGetGroupSetConfigResponse) SetHeaders(v map[string]*string) *BatchGetGroupSetConfigResponse {
	s.Headers = v
	return s
}

func (s *BatchGetGroupSetConfigResponse) SetBody(v *BatchGetGroupSetConfigResponseBody) *BatchGetGroupSetConfigResponse {
	s.Body = v
	return s
}

type ListTicketOperateRecordHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListTicketOperateRecordHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListTicketOperateRecordHeaders) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordHeaders) SetCommonHeaders(v map[string]*string) *ListTicketOperateRecordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListTicketOperateRecordHeaders) SetXAcsDingtalkAccessToken(v string) *ListTicketOperateRecordHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListTicketOperateRecordRequest struct {
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 工单开放ID
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
}

func (s ListTicketOperateRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTicketOperateRecordRequest) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordRequest) SetOpenTeamId(v string) *ListTicketOperateRecordRequest {
	s.OpenTeamId = &v
	return s
}

func (s *ListTicketOperateRecordRequest) SetOpenTicketId(v string) *ListTicketOperateRecordRequest {
	s.OpenTicketId = &v
	return s
}

type ListTicketOperateRecordResponseBody struct {
	// Id of the request
	Records []*ListTicketOperateRecordResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
}

func (s ListTicketOperateRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTicketOperateRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordResponseBody) SetRecords(v []*ListTicketOperateRecordResponseBodyRecords) *ListTicketOperateRecordResponseBody {
	s.Records = v
	return s
}

type ListTicketOperateRecordResponseBodyRecords struct {
	// 工单开放ID
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	// 操作时间
	OperateTime *string `json:"operateTime,omitempty" xml:"operateTime,omitempty"`
	// 动作
	Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
	// 动作展示名
	OperationDisplayName *string                                               `json:"operationDisplayName,omitempty" xml:"operationDisplayName,omitempty"`
	Operator             *ListTicketOperateRecordResponseBodyRecordsOperator   `json:"operator,omitempty" xml:"operator,omitempty" type:"Struct"`
	TicketMemo           *ListTicketOperateRecordResponseBodyRecordsTicketMemo `json:"ticketMemo,omitempty" xml:"ticketMemo,omitempty" type:"Struct"`
	OperateData          *string                                               `json:"operateData,omitempty" xml:"operateData,omitempty"`
}

func (s ListTicketOperateRecordResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ListTicketOperateRecordResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordResponseBodyRecords) SetOpenTicketId(v string) *ListTicketOperateRecordResponseBodyRecords {
	s.OpenTicketId = &v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecords) SetOperateTime(v string) *ListTicketOperateRecordResponseBodyRecords {
	s.OperateTime = &v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecords) SetOperation(v string) *ListTicketOperateRecordResponseBodyRecords {
	s.Operation = &v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecords) SetOperationDisplayName(v string) *ListTicketOperateRecordResponseBodyRecords {
	s.OperationDisplayName = &v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecords) SetOperator(v *ListTicketOperateRecordResponseBodyRecordsOperator) *ListTicketOperateRecordResponseBodyRecords {
	s.Operator = v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecords) SetTicketMemo(v *ListTicketOperateRecordResponseBodyRecordsTicketMemo) *ListTicketOperateRecordResponseBodyRecords {
	s.TicketMemo = v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecords) SetOperateData(v string) *ListTicketOperateRecordResponseBodyRecords {
	s.OperateData = &v
	return s
}

type ListTicketOperateRecordResponseBodyRecordsOperator struct {
	UnionId  *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
}

func (s ListTicketOperateRecordResponseBodyRecordsOperator) String() string {
	return tea.Prettify(s)
}

func (s ListTicketOperateRecordResponseBodyRecordsOperator) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordResponseBodyRecordsOperator) SetUnionId(v string) *ListTicketOperateRecordResponseBodyRecordsOperator {
	s.UnionId = &v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecordsOperator) SetNickName(v string) *ListTicketOperateRecordResponseBodyRecordsOperator {
	s.NickName = &v
	return s
}

type ListTicketOperateRecordResponseBodyRecordsTicketMemo struct {
	Memo        *string                                                            `json:"memo,omitempty" xml:"memo,omitempty"`
	Attachments []*ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
}

func (s ListTicketOperateRecordResponseBodyRecordsTicketMemo) String() string {
	return tea.Prettify(s)
}

func (s ListTicketOperateRecordResponseBodyRecordsTicketMemo) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordResponseBodyRecordsTicketMemo) SetMemo(v string) *ListTicketOperateRecordResponseBodyRecordsTicketMemo {
	s.Memo = &v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecordsTicketMemo) SetAttachments(v []*ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments) *ListTicketOperateRecordResponseBodyRecordsTicketMemo {
	s.Attachments = v
	return s
}

type ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments struct {
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	Key      *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments) String() string {
	return tea.Prettify(s)
}

func (s ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments) SetFileName(v string) *ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments {
	s.FileName = &v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments) SetKey(v string) *ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments {
	s.Key = &v
	return s
}

type ListTicketOperateRecordResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTicketOperateRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTicketOperateRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTicketOperateRecordResponse) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordResponse) SetHeaders(v map[string]*string) *ListTicketOperateRecordResponse {
	s.Headers = v
	return s
}

func (s *ListTicketOperateRecordResponse) SetBody(v *ListTicketOperateRecordResponseBody) *ListTicketOperateRecordResponse {
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
	// 接收者昵称
	ReceiverName *string `json:"receiverName,omitempty" xml:"receiverName,omitempty"`
	// 接收者dingtalkId
	ReceiverDingTalkId *string `json:"receiverDingTalkId,omitempty" xml:"receiverDingTalkId,omitempty"`
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

func (s *QueryServiceGroupMessageReadStatusResponseBodyRecords) SetReceiverName(v string) *QueryServiceGroupMessageReadStatusResponseBodyRecords {
	s.ReceiverName = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusResponseBodyRecords) SetReceiverDingTalkId(v string) *QueryServiceGroupMessageReadStatusResponseBodyRecords {
	s.ReceiverDingTalkId = &v
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
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 开放群ID
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 业务关联ID，和开放群ID二选一传
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
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

func (s *QueryGroupRequest) SetOpenTeamId(v string) *QueryGroupRequest {
	s.OpenTeamId = &v
	return s
}

func (s *QueryGroupRequest) SetOpenConversationId(v string) *QueryGroupRequest {
	s.OpenConversationId = &v
	return s
}

func (s *QueryGroupRequest) SetBizId(v string) *QueryGroupRequest {
	s.BizId = &v
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
	// 群bizId
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
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

func (s *QueryGroupResponseBody) SetBizId(v string) *QueryGroupResponseBody {
	s.BizId = &v
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

type CloseHumanSessionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CloseHumanSessionHeaders) String() string {
	return tea.Prettify(s)
}

func (s CloseHumanSessionHeaders) GoString() string {
	return s.String()
}

func (s *CloseHumanSessionHeaders) SetCommonHeaders(v map[string]*string) *CloseHumanSessionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CloseHumanSessionHeaders) SetXAcsDingtalkAccessToken(v string) *CloseHumanSessionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CloseHumanSessionRequest struct {
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	// 开放团队id
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 开放会话id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 访客unionid
	VisitorUnionId *int64 `json:"visitorUnionId,omitempty" xml:"visitorUnionId,omitempty"`
}

func (s CloseHumanSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseHumanSessionRequest) GoString() string {
	return s.String()
}

func (s *CloseHumanSessionRequest) SetDingIsvOrgId(v int64) *CloseHumanSessionRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *CloseHumanSessionRequest) SetDingOrgId(v int64) *CloseHumanSessionRequest {
	s.DingOrgId = &v
	return s
}

func (s *CloseHumanSessionRequest) SetDingSuiteKey(v string) *CloseHumanSessionRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *CloseHumanSessionRequest) SetDingTokenGrantType(v int64) *CloseHumanSessionRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *CloseHumanSessionRequest) SetOpenTeamId(v string) *CloseHumanSessionRequest {
	s.OpenTeamId = &v
	return s
}

func (s *CloseHumanSessionRequest) SetOpenConversationId(v string) *CloseHumanSessionRequest {
	s.OpenConversationId = &v
	return s
}

func (s *CloseHumanSessionRequest) SetVisitorUnionId(v int64) *CloseHumanSessionRequest {
	s.VisitorUnionId = &v
	return s
}

type CloseHumanSessionResponseBody struct {
	SessionId *int64 `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s CloseHumanSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseHumanSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CloseHumanSessionResponseBody) SetSessionId(v int64) *CloseHumanSessionResponseBody {
	s.SessionId = &v
	return s
}

type CloseHumanSessionResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CloseHumanSessionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloseHumanSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseHumanSessionResponse) GoString() string {
	return s.String()
}

func (s *CloseHumanSessionResponse) SetHeaders(v map[string]*string) *CloseHumanSessionResponse {
	s.Headers = v
	return s
}

func (s *CloseHumanSessionResponse) SetBody(v *CloseHumanSessionResponseBody) *CloseHumanSessionResponse {
	s.Body = v
	return s
}

type UrgeTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UrgeTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s UrgeTicketHeaders) GoString() string {
	return s.String()
}

func (s *UrgeTicketHeaders) SetCommonHeaders(v map[string]*string) *UrgeTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UrgeTicketHeaders) SetXAcsDingtalkAccessToken(v string) *UrgeTicketHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UrgeTicketRequest struct {
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	// 工单催单操作人UnionId
	OperatorUnionId *string `json:"operatorUnionId,omitempty" xml:"operatorUnionId,omitempty"`
	// 工单开放id
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	// 备注
	TicketMemo *UrgeTicketRequestTicketMemo `json:"ticketMemo,omitempty" xml:"ticketMemo,omitempty" type:"Struct"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s UrgeTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s UrgeTicketRequest) GoString() string {
	return s.String()
}

func (s *UrgeTicketRequest) SetDingIsvOrgId(v int64) *UrgeTicketRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *UrgeTicketRequest) SetDingOrgId(v int64) *UrgeTicketRequest {
	s.DingOrgId = &v
	return s
}

func (s *UrgeTicketRequest) SetDingSuiteKey(v string) *UrgeTicketRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *UrgeTicketRequest) SetDingTokenGrantType(v int64) *UrgeTicketRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *UrgeTicketRequest) SetOperatorUnionId(v string) *UrgeTicketRequest {
	s.OperatorUnionId = &v
	return s
}

func (s *UrgeTicketRequest) SetOpenTicketId(v string) *UrgeTicketRequest {
	s.OpenTicketId = &v
	return s
}

func (s *UrgeTicketRequest) SetTicketMemo(v *UrgeTicketRequestTicketMemo) *UrgeTicketRequest {
	s.TicketMemo = v
	return s
}

func (s *UrgeTicketRequest) SetOpenTeamId(v string) *UrgeTicketRequest {
	s.OpenTeamId = &v
	return s
}

type UrgeTicketRequestTicketMemo struct {
	// 备注文字
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 备注相关的附件
	Attachments []*UrgeTicketRequestTicketMemoAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
}

func (s UrgeTicketRequestTicketMemo) String() string {
	return tea.Prettify(s)
}

func (s UrgeTicketRequestTicketMemo) GoString() string {
	return s.String()
}

func (s *UrgeTicketRequestTicketMemo) SetMemo(v string) *UrgeTicketRequestTicketMemo {
	s.Memo = &v
	return s
}

func (s *UrgeTicketRequestTicketMemo) SetAttachments(v []*UrgeTicketRequestTicketMemoAttachments) *UrgeTicketRequestTicketMemo {
	s.Attachments = v
	return s
}

type UrgeTicketRequestTicketMemoAttachments struct {
	// 文件名
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s UrgeTicketRequestTicketMemoAttachments) String() string {
	return tea.Prettify(s)
}

func (s UrgeTicketRequestTicketMemoAttachments) GoString() string {
	return s.String()
}

func (s *UrgeTicketRequestTicketMemoAttachments) SetFileName(v string) *UrgeTicketRequestTicketMemoAttachments {
	s.FileName = &v
	return s
}

func (s *UrgeTicketRequestTicketMemoAttachments) SetKey(v string) *UrgeTicketRequestTicketMemoAttachments {
	s.Key = &v
	return s
}

type UrgeTicketResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UrgeTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s UrgeTicketResponse) GoString() string {
	return s.String()
}

func (s *UrgeTicketResponse) SetHeaders(v map[string]*string) *UrgeTicketResponse {
	s.Headers = v
	return s
}

type GetTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTicketHeaders) GoString() string {
	return s.String()
}

func (s *GetTicketHeaders) SetCommonHeaders(v map[string]*string) *GetTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTicketHeaders) SetXAcsDingtalkAccessToken(v string) *GetTicketHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTicketRequest struct {
	// eKWh3GBwsKEiE
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// hNiPO2OVktNMiE
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
}

func (s GetTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTicketRequest) GoString() string {
	return s.String()
}

func (s *GetTicketRequest) SetOpenTeamId(v string) *GetTicketRequest {
	s.OpenTeamId = &v
	return s
}

func (s *GetTicketRequest) SetOpenTicketId(v string) *GetTicketRequest {
	s.OpenTicketId = &v
	return s
}

type GetTicketResponseBody struct {
	// Id of the request
	OpenTicketId       *string                         `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	CreateTime         *string                         `json:"createTime,omitempty" xml:"createTime,omitempty"`
	UpdateTime         *string                         `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	OpenConversationId *string                         `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	Creator            *GetTicketResponseBodyCreator   `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	Processor          *GetTicketResponseBodyProcessor `json:"processor,omitempty" xml:"processor,omitempty" type:"Struct"`
	Takers             []*GetTicketResponseBodyTakers  `json:"takers,omitempty" xml:"takers,omitempty" type:"Repeated"`
	Stage              *string                         `json:"stage,omitempty" xml:"stage,omitempty"`
	Title              *string                         `json:"title,omitempty" xml:"title,omitempty"`
	CustomFields       *string                         `json:"customFields,omitempty" xml:"customFields,omitempty"`
	Scene              *string                         `json:"scene,omitempty" xml:"scene,omitempty"`
	SceneContext       *string                         `json:"sceneContext,omitempty" xml:"sceneContext,omitempty"`
	Template           *GetTicketResponseBodyTemplate  `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
}

func (s GetTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBody) SetOpenTicketId(v string) *GetTicketResponseBody {
	s.OpenTicketId = &v
	return s
}

func (s *GetTicketResponseBody) SetCreateTime(v string) *GetTicketResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetTicketResponseBody) SetUpdateTime(v string) *GetTicketResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetTicketResponseBody) SetOpenConversationId(v string) *GetTicketResponseBody {
	s.OpenConversationId = &v
	return s
}

func (s *GetTicketResponseBody) SetCreator(v *GetTicketResponseBodyCreator) *GetTicketResponseBody {
	s.Creator = v
	return s
}

func (s *GetTicketResponseBody) SetProcessor(v *GetTicketResponseBodyProcessor) *GetTicketResponseBody {
	s.Processor = v
	return s
}

func (s *GetTicketResponseBody) SetTakers(v []*GetTicketResponseBodyTakers) *GetTicketResponseBody {
	s.Takers = v
	return s
}

func (s *GetTicketResponseBody) SetStage(v string) *GetTicketResponseBody {
	s.Stage = &v
	return s
}

func (s *GetTicketResponseBody) SetTitle(v string) *GetTicketResponseBody {
	s.Title = &v
	return s
}

func (s *GetTicketResponseBody) SetCustomFields(v string) *GetTicketResponseBody {
	s.CustomFields = &v
	return s
}

func (s *GetTicketResponseBody) SetScene(v string) *GetTicketResponseBody {
	s.Scene = &v
	return s
}

func (s *GetTicketResponseBody) SetSceneContext(v string) *GetTicketResponseBody {
	s.SceneContext = &v
	return s
}

func (s *GetTicketResponseBody) SetTemplate(v *GetTicketResponseBodyTemplate) *GetTicketResponseBody {
	s.Template = v
	return s
}

type GetTicketResponseBodyCreator struct {
	UnionId  *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
}

func (s GetTicketResponseBodyCreator) String() string {
	return tea.Prettify(s)
}

func (s GetTicketResponseBodyCreator) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBodyCreator) SetUnionId(v string) *GetTicketResponseBodyCreator {
	s.UnionId = &v
	return s
}

func (s *GetTicketResponseBodyCreator) SetNickName(v string) *GetTicketResponseBodyCreator {
	s.NickName = &v
	return s
}

type GetTicketResponseBodyProcessor struct {
	UnionId  *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
}

func (s GetTicketResponseBodyProcessor) String() string {
	return tea.Prettify(s)
}

func (s GetTicketResponseBodyProcessor) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBodyProcessor) SetUnionId(v string) *GetTicketResponseBodyProcessor {
	s.UnionId = &v
	return s
}

func (s *GetTicketResponseBodyProcessor) SetNickName(v string) *GetTicketResponseBodyProcessor {
	s.NickName = &v
	return s
}

type GetTicketResponseBodyTakers struct {
	UnionId  *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
}

func (s GetTicketResponseBodyTakers) String() string {
	return tea.Prettify(s)
}

func (s GetTicketResponseBodyTakers) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBodyTakers) SetUnionId(v string) *GetTicketResponseBodyTakers {
	s.UnionId = &v
	return s
}

func (s *GetTicketResponseBodyTakers) SetNickName(v string) *GetTicketResponseBodyTakers {
	s.NickName = &v
	return s
}

type GetTicketResponseBodyTemplate struct {
	// 工单模版ID
	OpenTemplateId *string `json:"openTemplateId,omitempty" xml:"openTemplateId,omitempty"`
	// 工单模版业务ID
	OpenTemplateBizId *string `json:"openTemplateBizId,omitempty" xml:"openTemplateBizId,omitempty"`
	// 工单模版名称
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s GetTicketResponseBodyTemplate) String() string {
	return tea.Prettify(s)
}

func (s GetTicketResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBodyTemplate) SetOpenTemplateId(v string) *GetTicketResponseBodyTemplate {
	s.OpenTemplateId = &v
	return s
}

func (s *GetTicketResponseBodyTemplate) SetOpenTemplateBizId(v string) *GetTicketResponseBodyTemplate {
	s.OpenTemplateBizId = &v
	return s
}

func (s *GetTicketResponseBodyTemplate) SetTemplateName(v string) *GetTicketResponseBodyTemplate {
	s.TemplateName = &v
	return s
}

type GetTicketResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTicketResponse) GoString() string {
	return s.String()
}

func (s *GetTicketResponse) SetHeaders(v map[string]*string) *GetTicketResponse {
	s.Headers = v
	return s
}

func (s *GetTicketResponse) SetBody(v *GetTicketResponseBody) *GetTicketResponse {
	s.Body = v
	return s
}

type GetOssTempUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOssTempUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOssTempUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetOssTempUrlHeaders) SetCommonHeaders(v map[string]*string) *GetOssTempUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOssTempUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GetOssTempUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOssTempUrlRequest struct {
	// 团队开放ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// oss文件key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 文件名
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 访问模式 AUTO(自动，例如在浏览器中如果是图片，PDF等可以在线直接查看，不能在线看时自动下载)、DOWNLOAD（直接下载）
	FetchMode *string `json:"fetchMode,omitempty" xml:"fetchMode,omitempty"`
}

func (s GetOssTempUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOssTempUrlRequest) GoString() string {
	return s.String()
}

func (s *GetOssTempUrlRequest) SetOpenTeamId(v string) *GetOssTempUrlRequest {
	s.OpenTeamId = &v
	return s
}

func (s *GetOssTempUrlRequest) SetKey(v string) *GetOssTempUrlRequest {
	s.Key = &v
	return s
}

func (s *GetOssTempUrlRequest) SetFileName(v string) *GetOssTempUrlRequest {
	s.FileName = &v
	return s
}

func (s *GetOssTempUrlRequest) SetFetchMode(v string) *GetOssTempUrlRequest {
	s.FetchMode = &v
	return s
}

type GetOssTempUrlResponseBody struct {
	// Id of the request
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetOssTempUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOssTempUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssTempUrlResponseBody) SetUrl(v string) *GetOssTempUrlResponseBody {
	s.Url = &v
	return s
}

type GetOssTempUrlResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOssTempUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOssTempUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOssTempUrlResponse) GoString() string {
	return s.String()
}

func (s *GetOssTempUrlResponse) SetHeaders(v map[string]*string) *GetOssTempUrlResponse {
	s.Headers = v
	return s
}

func (s *GetOssTempUrlResponse) SetBody(v *GetOssTempUrlResponseBody) *GetOssTempUrlResponse {
	s.Body = v
	return s
}

type TakeTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TakeTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s TakeTicketHeaders) GoString() string {
	return s.String()
}

func (s *TakeTicketHeaders) SetCommonHeaders(v map[string]*string) *TakeTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TakeTicketHeaders) SetXAcsDingtalkAccessToken(v string) *TakeTicketHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TakeTicketRequest struct {
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	OpenTeamId         *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	TakerUnionId       *string `json:"takerUnionId,omitempty" xml:"takerUnionId,omitempty"`
	OpenTicketId       *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
}

func (s TakeTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s TakeTicketRequest) GoString() string {
	return s.String()
}

func (s *TakeTicketRequest) SetDingIsvOrgId(v int64) *TakeTicketRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *TakeTicketRequest) SetDingOrgId(v int64) *TakeTicketRequest {
	s.DingOrgId = &v
	return s
}

func (s *TakeTicketRequest) SetDingSuiteKey(v string) *TakeTicketRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *TakeTicketRequest) SetDingTokenGrantType(v int64) *TakeTicketRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *TakeTicketRequest) SetOpenTeamId(v string) *TakeTicketRequest {
	s.OpenTeamId = &v
	return s
}

func (s *TakeTicketRequest) SetTakerUnionId(v string) *TakeTicketRequest {
	s.TakerUnionId = &v
	return s
}

func (s *TakeTicketRequest) SetOpenTicketId(v string) *TakeTicketRequest {
	s.OpenTicketId = &v
	return s
}

type TakeTicketResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s TakeTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s TakeTicketResponse) GoString() string {
	return s.String()
}

func (s *TakeTicketResponse) SetHeaders(v map[string]*string) *TakeTicketResponse {
	s.Headers = v
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

type GetStoragePolicyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetStoragePolicyHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetStoragePolicyHeaders) GoString() string {
	return s.String()
}

func (s *GetStoragePolicyHeaders) SetCommonHeaders(v map[string]*string) *GetStoragePolicyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetStoragePolicyHeaders) SetXAcsDingtalkAccessToken(v string) *GetStoragePolicyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetStoragePolicyRequest struct {
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	// 团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 业务类型
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 文件大小，单位字节
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// 文件名称
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s GetStoragePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStoragePolicyRequest) GoString() string {
	return s.String()
}

func (s *GetStoragePolicyRequest) SetDingIsvOrgId(v int64) *GetStoragePolicyRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *GetStoragePolicyRequest) SetDingOrgId(v int64) *GetStoragePolicyRequest {
	s.DingOrgId = &v
	return s
}

func (s *GetStoragePolicyRequest) SetDingSuiteKey(v string) *GetStoragePolicyRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *GetStoragePolicyRequest) SetDingTokenGrantType(v int64) *GetStoragePolicyRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *GetStoragePolicyRequest) SetOpenTeamId(v string) *GetStoragePolicyRequest {
	s.OpenTeamId = &v
	return s
}

func (s *GetStoragePolicyRequest) SetBizType(v string) *GetStoragePolicyRequest {
	s.BizType = &v
	return s
}

func (s *GetStoragePolicyRequest) SetFileSize(v int64) *GetStoragePolicyRequest {
	s.FileSize = &v
	return s
}

func (s *GetStoragePolicyRequest) SetFileName(v string) *GetStoragePolicyRequest {
	s.FileName = &v
	return s
}

type GetStoragePolicyResponseBody struct {
	// Id of the request
	Key         *string `json:"key,omitempty" xml:"key,omitempty"`
	Policy      *string `json:"policy,omitempty" xml:"policy,omitempty"`
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	Endpoint    *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Signature   *string `json:"signature,omitempty" xml:"signature,omitempty"`
}

func (s GetStoragePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStoragePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetStoragePolicyResponseBody) SetKey(v string) *GetStoragePolicyResponseBody {
	s.Key = &v
	return s
}

func (s *GetStoragePolicyResponseBody) SetPolicy(v string) *GetStoragePolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *GetStoragePolicyResponseBody) SetAccessKeyId(v string) *GetStoragePolicyResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *GetStoragePolicyResponseBody) SetEndpoint(v string) *GetStoragePolicyResponseBody {
	s.Endpoint = &v
	return s
}

func (s *GetStoragePolicyResponseBody) SetSignature(v string) *GetStoragePolicyResponseBody {
	s.Signature = &v
	return s
}

type GetStoragePolicyResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetStoragePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStoragePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStoragePolicyResponse) GoString() string {
	return s.String()
}

func (s *GetStoragePolicyResponse) SetHeaders(v map[string]*string) *GetStoragePolicyResponse {
	s.Headers = v
	return s
}

func (s *GetStoragePolicyResponse) SetBody(v *GetStoragePolicyResponseBody) *GetStoragePolicyResponse {
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

type AddTicketMemoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddTicketMemoHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddTicketMemoHeaders) GoString() string {
	return s.String()
}

func (s *AddTicketMemoHeaders) SetCommonHeaders(v map[string]*string) *AddTicketMemoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddTicketMemoHeaders) SetXAcsDingtalkAccessToken(v string) *AddTicketMemoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddTicketMemoRequest struct {
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 当前工单处理人
	ProcessorUnionId *string `json:"processorUnionId,omitempty" xml:"processorUnionId,omitempty"`
	// 工单开放ID
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	// 备注
	TicketMemo *AddTicketMemoRequestTicketMemo `json:"ticketMemo,omitempty" xml:"ticketMemo,omitempty" type:"Struct"`
}

func (s AddTicketMemoRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTicketMemoRequest) GoString() string {
	return s.String()
}

func (s *AddTicketMemoRequest) SetDingIsvOrgId(v int64) *AddTicketMemoRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *AddTicketMemoRequest) SetDingOrgId(v int64) *AddTicketMemoRequest {
	s.DingOrgId = &v
	return s
}

func (s *AddTicketMemoRequest) SetDingSuiteKey(v string) *AddTicketMemoRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *AddTicketMemoRequest) SetDingTokenGrantType(v int64) *AddTicketMemoRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *AddTicketMemoRequest) SetOpenTeamId(v string) *AddTicketMemoRequest {
	s.OpenTeamId = &v
	return s
}

func (s *AddTicketMemoRequest) SetProcessorUnionId(v string) *AddTicketMemoRequest {
	s.ProcessorUnionId = &v
	return s
}

func (s *AddTicketMemoRequest) SetOpenTicketId(v string) *AddTicketMemoRequest {
	s.OpenTicketId = &v
	return s
}

func (s *AddTicketMemoRequest) SetTicketMemo(v *AddTicketMemoRequestTicketMemo) *AddTicketMemoRequest {
	s.TicketMemo = v
	return s
}

type AddTicketMemoRequestTicketMemo struct {
	// 文字备注
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 备注相关的附件
	Attachments []*AddTicketMemoRequestTicketMemoAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
}

func (s AddTicketMemoRequestTicketMemo) String() string {
	return tea.Prettify(s)
}

func (s AddTicketMemoRequestTicketMemo) GoString() string {
	return s.String()
}

func (s *AddTicketMemoRequestTicketMemo) SetMemo(v string) *AddTicketMemoRequestTicketMemo {
	s.Memo = &v
	return s
}

func (s *AddTicketMemoRequestTicketMemo) SetAttachments(v []*AddTicketMemoRequestTicketMemoAttachments) *AddTicketMemoRequestTicketMemo {
	s.Attachments = v
	return s
}

type AddTicketMemoRequestTicketMemoAttachments struct {
	// 文件名
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s AddTicketMemoRequestTicketMemoAttachments) String() string {
	return tea.Prettify(s)
}

func (s AddTicketMemoRequestTicketMemoAttachments) GoString() string {
	return s.String()
}

func (s *AddTicketMemoRequestTicketMemoAttachments) SetFileName(v string) *AddTicketMemoRequestTicketMemoAttachments {
	s.FileName = &v
	return s
}

func (s *AddTicketMemoRequestTicketMemoAttachments) SetKey(v string) *AddTicketMemoRequestTicketMemoAttachments {
	s.Key = &v
	return s
}

type AddTicketMemoResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s AddTicketMemoResponse) String() string {
	return tea.Prettify(s)
}

func (s AddTicketMemoResponse) GoString() string {
	return s.String()
}

func (s *AddTicketMemoResponse) SetHeaders(v map[string]*string) *AddTicketMemoResponse {
	s.Headers = v
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

	if !tea.BoolValue(util.IsUnset(request.CreatorUnionId)) {
		body["creatorUnionId"] = request.CreatorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessorUnionIds)) {
		body["processorUnionIds"] = request.ProcessorUnionIds
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SceneContext))) {
		body["sceneContext"] = request.SceneContext
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTemplateBizId)) {
		body["openTemplateBizId"] = request.OpenTemplateBizId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.CustomFields)) {
		body["customFields"] = request.CustomFields
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Notify))) {
		body["notify"] = request.Notify
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
	_body, _err := client.DoROARequest(tea.String("CreateTicket"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tickets"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssignTicket(request *AssignTicketRequest) (_result *AssignTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AssignTicketHeaders{}
	_result = &AssignTicketResponse{}
	_body, _err := client.AssignTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssignTicketWithOptions(request *AssignTicketRequest, headers *AssignTicketHeaders, runtime *util.RuntimeOptions) (_result *AssignTicketResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.OperatorUnionId)) {
		body["operatorUnionId"] = request.OperatorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTicketId)) {
		body["openTicketId"] = request.OpenTicketId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessorUnionIds)) {
		body["processorUnionIds"] = request.ProcessorUnionIds
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TicketMemo))) {
		body["ticketMemo"] = request.TicketMemo
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Notify))) {
		body["notify"] = request.Notify
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
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
	_result = &AssignTicketResponse{}
	_body, _err := client.DoROARequest(tea.String("AssignTicket"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tickets/assign"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FinishTicket(request *FinishTicketRequest) (_result *FinishTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &FinishTicketHeaders{}
	_result = &FinishTicketResponse{}
	_body, _err := client.FinishTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FinishTicketWithOptions(request *FinishTicketRequest, headers *FinishTicketHeaders, runtime *util.RuntimeOptions) (_result *FinishTicketResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ProcessorUnionId)) {
		body["processorUnionId"] = request.ProcessorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTicketId)) {
		body["openTicketId"] = request.OpenTicketId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TicketMemo))) {
		body["ticketMemo"] = request.TicketMemo
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Notify))) {
		body["notify"] = request.Notify
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
	_result = &FinishTicketResponse{}
	_body, _err := client.DoROARequest(tea.String("FinishTicket"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tickets/finish"), tea.String("none"), req, runtime)
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

func (client *Client) UpdateTicket(request *UpdateTicketRequest) (_result *UpdateTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTicketHeaders{}
	_result = &UpdateTicketResponse{}
	_body, _err := client.UpdateTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTicketWithOptions(request *UpdateTicketRequest, headers *UpdateTicketHeaders, runtime *util.RuntimeOptions) (_result *UpdateTicketResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ProcessorUnionId)) {
		body["processorUnionId"] = request.ProcessorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTicketId)) {
		body["openTicketId"] = request.OpenTicketId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomFields)) {
		body["customFields"] = request.CustomFields
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TicketMemo))) {
		body["ticketMemo"] = request.TicketMemo
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
	_result = &UpdateTicketResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateTicket"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tickets"), tea.String("none"), req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.GroupTagNames)) {
		body["groupTagNames"] = request.GroupTagNames
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

func (client *Client) TransferTicket(request *TransferTicketRequest) (_result *TransferTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TransferTicketHeaders{}
	_result = &TransferTicketResponse{}
	_body, _err := client.TransferTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferTicketWithOptions(request *TransferTicketRequest, headers *TransferTicketHeaders, runtime *util.RuntimeOptions) (_result *TransferTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProcessorUnionId)) {
		body["processorUnionId"] = request.ProcessorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTicketId)) {
		body["openTicketId"] = request.OpenTicketId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessorUnionIds)) {
		body["processorUnionIds"] = request.ProcessorUnionIds
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TicketMemo))) {
		body["ticketMemo"] = request.TicketMemo
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Notify))) {
		body["notify"] = request.Notify
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
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
	_result = &TransferTicketResponse{}
	_body, _err := client.DoROARequest(tea.String("TransferTicket"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tickets/transfer"), tea.String("none"), req, runtime)
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

func (client *Client) BatchGetGroupSetConfig(request *BatchGetGroupSetConfigRequest) (_result *BatchGetGroupSetConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchGetGroupSetConfigHeaders{}
	_result = &BatchGetGroupSetConfigResponse{}
	_body, _err := client.BatchGetGroupSetConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchGetGroupSetConfigWithOptions(request *BatchGetGroupSetConfigRequest, headers *BatchGetGroupSetConfigHeaders, runtime *util.RuntimeOptions) (_result *BatchGetGroupSetConfigResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.OpenGroupSetId)) {
		body["openGroupSetId"] = request.OpenGroupSetId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigKeys)) {
		body["configKeys"] = request.ConfigKeys
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
	_result = &BatchGetGroupSetConfigResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchGetGroupSetConfig"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/groupSetConfigs/batchQuery"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTicketOperateRecord(request *ListTicketOperateRecordRequest) (_result *ListTicketOperateRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListTicketOperateRecordHeaders{}
	_result = &ListTicketOperateRecordResponse{}
	_body, _err := client.ListTicketOperateRecordWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTicketOperateRecordWithOptions(request *ListTicketOperateRecordRequest, headers *ListTicketOperateRecordHeaders, runtime *util.RuntimeOptions) (_result *ListTicketOperateRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		query["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTicketId)) {
		query["openTicketId"] = request.OpenTicketId
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
	_result = &ListTicketOperateRecordResponse{}
	_body, _err := client.DoROARequest(tea.String("ListTicketOperateRecord"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tickets/operateRecords"), tea.String("json"), req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
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

func (client *Client) CloseHumanSession(request *CloseHumanSessionRequest) (_result *CloseHumanSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CloseHumanSessionHeaders{}
	_result = &CloseHumanSessionResponse{}
	_body, _err := client.CloseHumanSessionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloseHumanSessionWithOptions(request *CloseHumanSessionRequest, headers *CloseHumanSessionHeaders, runtime *util.RuntimeOptions) (_result *CloseHumanSessionResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.VisitorUnionId)) {
		body["visitorUnionId"] = request.VisitorUnionId
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
	_result = &CloseHumanSessionResponse{}
	_body, _err := client.DoROARequest(tea.String("CloseHumanSession"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/humanSessions/close"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UrgeTicket(request *UrgeTicketRequest) (_result *UrgeTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UrgeTicketHeaders{}
	_result = &UrgeTicketResponse{}
	_body, _err := client.UrgeTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UrgeTicketWithOptions(request *UrgeTicketRequest, headers *UrgeTicketHeaders, runtime *util.RuntimeOptions) (_result *UrgeTicketResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.OperatorUnionId)) {
		body["operatorUnionId"] = request.OperatorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTicketId)) {
		body["openTicketId"] = request.OpenTicketId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TicketMemo))) {
		body["ticketMemo"] = request.TicketMemo
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
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
	_result = &UrgeTicketResponse{}
	_body, _err := client.DoROARequest(tea.String("UrgeTicket"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tickets/urge"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTicket(request *GetTicketRequest) (_result *GetTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTicketHeaders{}
	_result = &GetTicketResponse{}
	_body, _err := client.GetTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTicketWithOptions(request *GetTicketRequest, headers *GetTicketHeaders, runtime *util.RuntimeOptions) (_result *GetTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		query["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTicketId)) {
		query["openTicketId"] = request.OpenTicketId
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
	_result = &GetTicketResponse{}
	_body, _err := client.DoROARequest(tea.String("GetTicket"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tickets"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOssTempUrl(request *GetOssTempUrlRequest) (_result *GetOssTempUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOssTempUrlHeaders{}
	_result = &GetOssTempUrlResponse{}
	_body, _err := client.GetOssTempUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOssTempUrlWithOptions(request *GetOssTempUrlRequest, headers *GetOssTempUrlHeaders, runtime *util.RuntimeOptions) (_result *GetOssTempUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		query["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FetchMode)) {
		query["fetchMode"] = request.FetchMode
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
	_result = &GetOssTempUrlResponse{}
	_body, _err := client.DoROARequest(tea.String("GetOssTempUrl"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/serviceGroup/ossServices/tempUrls"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TakeTicket(request *TakeTicketRequest) (_result *TakeTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TakeTicketHeaders{}
	_result = &TakeTicketResponse{}
	_body, _err := client.TakeTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TakeTicketWithOptions(request *TakeTicketRequest, headers *TakeTicketHeaders, runtime *util.RuntimeOptions) (_result *TakeTicketResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.TakerUnionId)) {
		body["takerUnionId"] = request.TakerUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTicketId)) {
		body["openTicketId"] = request.OpenTicketId
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
	_result = &TakeTicketResponse{}
	_body, _err := client.DoROARequest(tea.String("TakeTicket"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tickets/apply"), tea.String("none"), req, runtime)
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

func (client *Client) GetStoragePolicy(request *GetStoragePolicyRequest) (_result *GetStoragePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetStoragePolicyHeaders{}
	_result = &GetStoragePolicyResponse{}
	_body, _err := client.GetStoragePolicyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStoragePolicyWithOptions(request *GetStoragePolicyRequest, headers *GetStoragePolicyHeaders, runtime *util.RuntimeOptions) (_result *GetStoragePolicyResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.FileSize)) {
		body["fileSize"] = request.FileSize
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["fileName"] = request.FileName
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
	_result = &GetStoragePolicyResponse{}
	_body, _err := client.DoROARequest(tea.String("GetStoragePolicy"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/ossServices/policies"), tea.String("json"), req, runtime)
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

func (client *Client) AddTicketMemo(request *AddTicketMemoRequest) (_result *AddTicketMemoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddTicketMemoHeaders{}
	_result = &AddTicketMemoResponse{}
	_body, _err := client.AddTicketMemoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddTicketMemoWithOptions(request *AddTicketMemoRequest, headers *AddTicketMemoHeaders, runtime *util.RuntimeOptions) (_result *AddTicketMemoResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ProcessorUnionId)) {
		body["processorUnionId"] = request.ProcessorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTicketId)) {
		body["openTicketId"] = request.OpenTicketId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TicketMemo))) {
		body["ticketMemo"] = request.TicketMemo
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
	_result = &AddTicketMemoResponse{}
	_body, _err := client.DoROARequest(tea.String("AddTicketMemo"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tickets/memos"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
