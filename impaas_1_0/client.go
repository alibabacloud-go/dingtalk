// This file is auto-generated, don't edit it. Thanks.
package impaas_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddGroupMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	OperationSource         *string            `json:"operationSource,omitempty" xml:"operationSource,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddGroupMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMembersHeaders) GoString() string {
	return s.String()
}

func (s *AddGroupMembersHeaders) SetCommonHeaders(v map[string]*string) *AddGroupMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddGroupMembersHeaders) SetOperationSource(v string) *AddGroupMembersHeaders {
	s.OperationSource = &v
	return s
}

func (s *AddGroupMembersHeaders) SetXAcsDingtalkAccessToken(v string) *AddGroupMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddGroupMembersRequest struct {
	// This parameter is required.
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// true
	Members     []*AddGroupMembersRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	OperatorUid *string                          `json:"operatorUid,omitempty" xml:"operatorUid,omitempty"`
}

func (s AddGroupMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *AddGroupMembersRequest) SetConversationId(v string) *AddGroupMembersRequest {
	s.ConversationId = &v
	return s
}

func (s *AddGroupMembersRequest) SetMembers(v []*AddGroupMembersRequestMembers) *AddGroupMembersRequest {
	s.Members = v
	return s
}

func (s *AddGroupMembersRequest) SetOperatorUid(v string) *AddGroupMembersRequest {
	s.OperatorUid = &v
	return s
}

type AddGroupMembersRequestMembers struct {
	Nick *string `json:"nick,omitempty" xml:"nick,omitempty"`
	// This parameter is required.
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s AddGroupMembersRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *AddGroupMembersRequestMembers) SetNick(v string) *AddGroupMembersRequestMembers {
	s.Nick = &v
	return s
}

func (s *AddGroupMembersRequestMembers) SetUid(v string) *AddGroupMembersRequestMembers {
	s.Uid = &v
	return s
}

type AddGroupMembersResponseBody struct {
	// This parameter is required.
	MemberUids []*string `json:"memberUids,omitempty" xml:"memberUids,omitempty" type:"Repeated"`
}

func (s AddGroupMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMembersResponseBody) GoString() string {
	return s.String()
}

func (s *AddGroupMembersResponseBody) SetMemberUids(v []*string) *AddGroupMembersResponseBody {
	s.MemberUids = v
	return s
}

type AddGroupMembersResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddGroupMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMembersResponse) GoString() string {
	return s.String()
}

func (s *AddGroupMembersResponse) SetHeaders(v map[string]*string) *AddGroupMembersResponse {
	s.Headers = v
	return s
}

func (s *AddGroupMembersResponse) SetStatusCode(v int32) *AddGroupMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddGroupMembersResponse) SetBody(v *AddGroupMembersResponseBody) *AddGroupMembersResponse {
	s.Body = v
	return s
}

type AddProfileHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddProfileHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddProfileHeaders) GoString() string {
	return s.String()
}

func (s *AddProfileHeaders) SetCommonHeaders(v map[string]*string) *AddProfileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddProfileHeaders) SetXAcsDingtalkAccessToken(v string) *AddProfileHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddProfileRequest struct {
	// This parameter is required.
	AppUid        *string `json:"appUid,omitempty" xml:"appUid,omitempty"`
	AvatarMediaId *string `json:"avatarMediaId,omitempty" xml:"avatarMediaId,omitempty"`
	// This parameter is required.
	MobileNumber *string `json:"mobileNumber,omitempty" xml:"mobileNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// usertest
	Nick *string `json:"nick,omitempty" xml:"nick,omitempty"`
}

func (s AddProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s AddProfileRequest) GoString() string {
	return s.String()
}

func (s *AddProfileRequest) SetAppUid(v string) *AddProfileRequest {
	s.AppUid = &v
	return s
}

func (s *AddProfileRequest) SetAvatarMediaId(v string) *AddProfileRequest {
	s.AvatarMediaId = &v
	return s
}

func (s *AddProfileRequest) SetMobileNumber(v string) *AddProfileRequest {
	s.MobileNumber = &v
	return s
}

func (s *AddProfileRequest) SetNick(v string) *AddProfileRequest {
	s.Nick = &v
	return s
}

type AddProfileResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s AddProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s AddProfileResponse) GoString() string {
	return s.String()
}

func (s *AddProfileResponse) SetHeaders(v map[string]*string) *AddProfileResponse {
	s.Headers = v
	return s
}

func (s *AddProfileResponse) SetStatusCode(v int32) *AddProfileResponse {
	s.StatusCode = &v
	return s
}

type BatchSendHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchSendHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchSendHeaders) GoString() string {
	return s.String()
}

func (s *BatchSendHeaders) SetCommonHeaders(v map[string]*string) *BatchSendHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchSendHeaders) SetXAcsDingtalkAccessToken(v string) *BatchSendHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchSendRequest struct {
	// if can be null:
	// true
	AppUids []*string `json:"appUids,omitempty" xml:"appUids,omitempty" type:"Repeated"`
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// if can be null:
	// true
	ConversationIds []*string `json:"conversationIds,omitempty" xml:"conversationIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 234#dingxxx@dingding
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchSendRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSendRequest) GoString() string {
	return s.String()
}

func (s *BatchSendRequest) SetAppUids(v []*string) *BatchSendRequest {
	s.AppUids = v
	return s
}

func (s *BatchSendRequest) SetContent(v string) *BatchSendRequest {
	s.Content = &v
	return s
}

func (s *BatchSendRequest) SetConversationIds(v []*string) *BatchSendRequest {
	s.ConversationIds = v
	return s
}

func (s *BatchSendRequest) SetUserId(v string) *BatchSendRequest {
	s.UserId = &v
	return s
}

type BatchSendResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 23233434
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s BatchSendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSendResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSendResponseBody) SetTaskId(v string) *BatchSendResponseBody {
	s.TaskId = &v
	return s
}

type BatchSendResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchSendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchSendResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchSendResponse) GoString() string {
	return s.String()
}

func (s *BatchSendResponse) SetHeaders(v map[string]*string) *BatchSendResponse {
	s.Headers = v
	return s
}

func (s *BatchSendResponse) SetStatusCode(v int32) *BatchSendResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSendResponse) SetBody(v *BatchSendResponseBody) *BatchSendResponse {
	s.Body = v
	return s
}

type CreateGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	OperationSource         *string            `json:"operationSource,omitempty" xml:"operationSource,omitempty"`
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

func (s *CreateGroupHeaders) SetOperationSource(v string) *CreateGroupHeaders {
	s.OperationSource = &v
	return s
}

func (s *CreateGroupHeaders) SetXAcsDingtalkAccessToken(v string) *CreateGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateGroupRequest struct {
	// This parameter is required.
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// This parameter is required.
	CreatorUid  *string `json:"creatorUid,omitempty" xml:"creatorUid,omitempty"`
	IconMediaId *string `json:"iconMediaId,omitempty" xml:"iconMediaId,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	Properties map[string]*string `json:"properties,omitempty" xml:"properties,omitempty"`
	// This parameter is required.
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s CreateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) SetChannel(v string) *CreateGroupRequest {
	s.Channel = &v
	return s
}

func (s *CreateGroupRequest) SetCreatorUid(v string) *CreateGroupRequest {
	s.CreatorUid = &v
	return s
}

func (s *CreateGroupRequest) SetIconMediaId(v string) *CreateGroupRequest {
	s.IconMediaId = &v
	return s
}

func (s *CreateGroupRequest) SetName(v string) *CreateGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateGroupRequest) SetProperties(v map[string]*string) *CreateGroupRequest {
	s.Properties = v
	return s
}

func (s *CreateGroupRequest) SetUuid(v string) *CreateGroupRequest {
	s.Uuid = &v
	return s
}

type CreateGroupResponseBody struct {
	// This parameter is required.
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// This parameter is required.
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// This parameter is required.
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
}

func (s CreateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBody) SetChatId(v string) *CreateGroupResponseBody {
	s.ChatId = &v
	return s
}

func (s *CreateGroupResponseBody) SetConversationId(v string) *CreateGroupResponseBody {
	s.ConversationId = &v
	return s
}

func (s *CreateGroupResponseBody) SetCreateTime(v int64) *CreateGroupResponseBody {
	s.CreateTime = &v
	return s
}

type CreateGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateGroupResponse) SetStatusCode(v int32) *CreateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupResponse) SetBody(v *CreateGroupResponseBody) *CreateGroupResponse {
	s.Body = v
	return s
}

type CreateTrustGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	OperationSource         *string            `json:"operationSource,omitempty" xml:"operationSource,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateTrustGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTrustGroupHeaders) GoString() string {
	return s.String()
}

func (s *CreateTrustGroupHeaders) SetCommonHeaders(v map[string]*string) *CreateTrustGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTrustGroupHeaders) SetOperationSource(v string) *CreateTrustGroupHeaders {
	s.OperationSource = &v
	return s
}

func (s *CreateTrustGroupHeaders) SetXAcsDingtalkAccessToken(v string) *CreateTrustGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateTrustGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// channel_abcd
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// example:
	//
	// @lALOKACADDA
	IconMediaId *string                           `json:"iconMediaId,omitempty" xml:"iconMediaId,omitempty"`
	Members     []*CreateTrustGroupRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 测试群名称XXX
	Name       *string            `json:"name,omitempty" xml:"name,omitempty"`
	Properties map[string]*string `json:"properties,omitempty" xml:"properties,omitempty"`
	// example:
	//
	// 你有新的会话
	SystemMsg *string `json:"systemMsg,omitempty" xml:"systemMsg,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1657099913071
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s CreateTrustGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTrustGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateTrustGroupRequest) SetChannel(v string) *CreateTrustGroupRequest {
	s.Channel = &v
	return s
}

func (s *CreateTrustGroupRequest) SetIconMediaId(v string) *CreateTrustGroupRequest {
	s.IconMediaId = &v
	return s
}

func (s *CreateTrustGroupRequest) SetMembers(v []*CreateTrustGroupRequestMembers) *CreateTrustGroupRequest {
	s.Members = v
	return s
}

func (s *CreateTrustGroupRequest) SetName(v string) *CreateTrustGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateTrustGroupRequest) SetProperties(v map[string]*string) *CreateTrustGroupRequest {
	s.Properties = v
	return s
}

func (s *CreateTrustGroupRequest) SetSystemMsg(v string) *CreateTrustGroupRequest {
	s.SystemMsg = &v
	return s
}

func (s *CreateTrustGroupRequest) SetUuid(v string) *CreateTrustGroupRequest {
	s.Uuid = &v
	return s
}

type CreateTrustGroupRequestMembers struct {
	Nick *string `json:"nick,omitempty" xml:"nick,omitempty"`
	Uid  *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s CreateTrustGroupRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s CreateTrustGroupRequestMembers) GoString() string {
	return s.String()
}

func (s *CreateTrustGroupRequestMembers) SetNick(v string) *CreateTrustGroupRequestMembers {
	s.Nick = &v
	return s
}

func (s *CreateTrustGroupRequestMembers) SetUid(v string) *CreateTrustGroupRequestMembers {
	s.Uid = &v
	return s
}

type CreateTrustGroupResponseBody struct {
	// This parameter is required.
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// This parameter is required.
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s CreateTrustGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTrustGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrustGroupResponseBody) SetChatId(v string) *CreateTrustGroupResponseBody {
	s.ChatId = &v
	return s
}

func (s *CreateTrustGroupResponseBody) SetCreateTime(v int64) *CreateTrustGroupResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateTrustGroupResponseBody) SetOpenConversationId(v string) *CreateTrustGroupResponseBody {
	s.OpenConversationId = &v
	return s
}

type CreateTrustGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTrustGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTrustGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTrustGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateTrustGroupResponse) SetHeaders(v map[string]*string) *CreateTrustGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateTrustGroupResponse) SetStatusCode(v int32) *CreateTrustGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrustGroupResponse) SetBody(v *CreateTrustGroupResponseBody) *CreateTrustGroupResponse {
	s.Body = v
	return s
}

type DismissGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	OperationSource         *string            `json:"operationSource,omitempty" xml:"operationSource,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DismissGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s DismissGroupHeaders) GoString() string {
	return s.String()
}

func (s *DismissGroupHeaders) SetCommonHeaders(v map[string]*string) *DismissGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DismissGroupHeaders) SetOperationSource(v string) *DismissGroupHeaders {
	s.OperationSource = &v
	return s
}

func (s *DismissGroupHeaders) SetXAcsDingtalkAccessToken(v string) *DismissGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DismissGroupRequest struct {
	// This parameter is required.
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	OperatorUid    *string `json:"operatorUid,omitempty" xml:"operatorUid,omitempty"`
}

func (s DismissGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DismissGroupRequest) GoString() string {
	return s.String()
}

func (s *DismissGroupRequest) SetConversationId(v string) *DismissGroupRequest {
	s.ConversationId = &v
	return s
}

func (s *DismissGroupRequest) SetOperatorUid(v string) *DismissGroupRequest {
	s.OperatorUid = &v
	return s
}

type DismissGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DismissGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DismissGroupResponse) GoString() string {
	return s.String()
}

func (s *DismissGroupResponse) SetHeaders(v map[string]*string) *DismissGroupResponse {
	s.Headers = v
	return s
}

func (s *DismissGroupResponse) SetStatusCode(v int32) *DismissGroupResponse {
	s.StatusCode = &v
	return s
}

type GetConversationIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetConversationIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetConversationIdHeaders) GoString() string {
	return s.String()
}

func (s *GetConversationIdHeaders) SetCommonHeaders(v map[string]*string) *GetConversationIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetConversationIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetConversationIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetConversationIdRequest struct {
	// This parameter is required.
	AppUid *string `json:"appUid,omitempty" xml:"appUid,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetConversationIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConversationIdRequest) GoString() string {
	return s.String()
}

func (s *GetConversationIdRequest) SetAppUid(v string) *GetConversationIdRequest {
	s.AppUid = &v
	return s
}

func (s *GetConversationIdRequest) SetUserId(v string) *GetConversationIdRequest {
	s.UserId = &v
	return s
}

type GetConversationIdResponseBody struct {
	// This parameter is required.
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
}

func (s GetConversationIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConversationIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetConversationIdResponseBody) SetConversationId(v string) *GetConversationIdResponseBody {
	s.ConversationId = &v
	return s
}

type GetConversationIdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConversationIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConversationIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConversationIdResponse) GoString() string {
	return s.String()
}

func (s *GetConversationIdResponse) SetHeaders(v map[string]*string) *GetConversationIdResponse {
	s.Headers = v
	return s
}

func (s *GetConversationIdResponse) SetStatusCode(v int32) *GetConversationIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConversationIdResponse) SetBody(v *GetConversationIdResponseBody) *GetConversationIdResponse {
	s.Body = v
	return s
}

type GetMediaUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetMediaUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMediaUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetMediaUrlHeaders) SetCommonHeaders(v map[string]*string) *GetMediaUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMediaUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GetMediaUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetMediaUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// @wesfsdfsfwe
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	// example:
	//
	// 86399
	UrlExpireTime *int32 `json:"urlExpireTime,omitempty" xml:"urlExpireTime,omitempty"`
}

func (s GetMediaUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMediaUrlRequest) GoString() string {
	return s.String()
}

func (s *GetMediaUrlRequest) SetMediaId(v string) *GetMediaUrlRequest {
	s.MediaId = &v
	return s
}

func (s *GetMediaUrlRequest) SetUrlExpireTime(v int32) *GetMediaUrlRequest {
	s.UrlExpireTime = &v
	return s
}

type GetMediaUrlResponseBody struct {
	// This parameter is required.
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetMediaUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMediaUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaUrlResponseBody) SetUrl(v string) *GetMediaUrlResponseBody {
	s.Url = &v
	return s
}

type GetMediaUrlResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMediaUrlResponse) GoString() string {
	return s.String()
}

func (s *GetMediaUrlResponse) SetHeaders(v map[string]*string) *GetMediaUrlResponse {
	s.Headers = v
	return s
}

func (s *GetMediaUrlResponse) SetStatusCode(v int32) *GetMediaUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaUrlResponse) SetBody(v *GetMediaUrlResponseBody) *GetMediaUrlResponse {
	s.Body = v
	return s
}

type GetMediaUrlsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetMediaUrlsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMediaUrlsHeaders) GoString() string {
	return s.String()
}

func (s *GetMediaUrlsHeaders) SetCommonHeaders(v map[string]*string) *GetMediaUrlsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMediaUrlsHeaders) SetXAcsDingtalkAccessToken(v string) *GetMediaUrlsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetMediaUrlsRequest struct {
	// This parameter is required.
	MediaIds []*string `json:"mediaIds,omitempty" xml:"mediaIds,omitempty" type:"Repeated"`
	// example:
	//
	// 86399
	UrlExpireTime *int32 `json:"urlExpireTime,omitempty" xml:"urlExpireTime,omitempty"`
}

func (s GetMediaUrlsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMediaUrlsRequest) GoString() string {
	return s.String()
}

func (s *GetMediaUrlsRequest) SetMediaIds(v []*string) *GetMediaUrlsRequest {
	s.MediaIds = v
	return s
}

func (s *GetMediaUrlsRequest) SetUrlExpireTime(v int32) *GetMediaUrlsRequest {
	s.UrlExpireTime = &v
	return s
}

type GetMediaUrlsResponseBody struct {
	Urls map[string]interface{} `json:"urls,omitempty" xml:"urls,omitempty"`
}

func (s GetMediaUrlsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMediaUrlsResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaUrlsResponseBody) SetUrls(v map[string]interface{}) *GetMediaUrlsResponseBody {
	s.Urls = v
	return s
}

type GetMediaUrlsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaUrlsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaUrlsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMediaUrlsResponse) GoString() string {
	return s.String()
}

func (s *GetMediaUrlsResponse) SetHeaders(v map[string]*string) *GetMediaUrlsResponse {
	s.Headers = v
	return s
}

func (s *GetMediaUrlsResponse) SetStatusCode(v int32) *GetMediaUrlsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaUrlsResponse) SetBody(v *GetMediaUrlsResponseBody) *GetMediaUrlsResponse {
	s.Body = v
	return s
}

type GetSpaceFileUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSpaceFileUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceFileUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetSpaceFileUrlHeaders) SetCommonHeaders(v map[string]*string) *GetSpaceFileUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSpaceFileUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GetSpaceFileUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSpaceFileUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123#123@dingding
	SenderUid *string `json:"senderUid,omitempty" xml:"senderUid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s GetSpaceFileUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceFileUrlRequest) GoString() string {
	return s.String()
}

func (s *GetSpaceFileUrlRequest) SetFileId(v string) *GetSpaceFileUrlRequest {
	s.FileId = &v
	return s
}

func (s *GetSpaceFileUrlRequest) SetSenderUid(v string) *GetSpaceFileUrlRequest {
	s.SenderUid = &v
	return s
}

func (s *GetSpaceFileUrlRequest) SetSpaceId(v string) *GetSpaceFileUrlRequest {
	s.SpaceId = &v
	return s
}

type GetSpaceFileUrlResponseBody struct {
	// This parameter is required.
	Headers map[string]interface{} `json:"headers,omitempty" xml:"headers,omitempty"`
	// This parameter is required.
	InternalResourceUrl *string `json:"internalResourceUrl,omitempty" xml:"internalResourceUrl,omitempty"`
	// This parameter is required.
	ResourceUrl *string `json:"resourceUrl,omitempty" xml:"resourceUrl,omitempty"`
}

func (s GetSpaceFileUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceFileUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetSpaceFileUrlResponseBody) SetHeaders(v map[string]interface{}) *GetSpaceFileUrlResponseBody {
	s.Headers = v
	return s
}

func (s *GetSpaceFileUrlResponseBody) SetInternalResourceUrl(v string) *GetSpaceFileUrlResponseBody {
	s.InternalResourceUrl = &v
	return s
}

func (s *GetSpaceFileUrlResponseBody) SetResourceUrl(v string) *GetSpaceFileUrlResponseBody {
	s.ResourceUrl = &v
	return s
}

type GetSpaceFileUrlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSpaceFileUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSpaceFileUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceFileUrlResponse) GoString() string {
	return s.String()
}

func (s *GetSpaceFileUrlResponse) SetHeaders(v map[string]*string) *GetSpaceFileUrlResponse {
	s.Headers = v
	return s
}

func (s *GetSpaceFileUrlResponse) SetStatusCode(v int32) *GetSpaceFileUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSpaceFileUrlResponse) SetBody(v *GetSpaceFileUrlResponseBody) *GetSpaceFileUrlResponse {
	s.Body = v
	return s
}

type ListGroupStaffMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListGroupStaffMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListGroupStaffMembersHeaders) GoString() string {
	return s.String()
}

func (s *ListGroupStaffMembersHeaders) SetCommonHeaders(v map[string]*string) *ListGroupStaffMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListGroupStaffMembersHeaders) SetXAcsDingtalkAccessToken(v string) *ListGroupStaffMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListGroupStaffMembersRequest struct {
	// This parameter is required.
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
}

func (s ListGroupStaffMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupStaffMembersRequest) GoString() string {
	return s.String()
}

func (s *ListGroupStaffMembersRequest) SetConversationId(v string) *ListGroupStaffMembersRequest {
	s.ConversationId = &v
	return s
}

type ListGroupStaffMembersResponseBody struct {
	// This parameter is required.
	Members []*ListGroupStaffMembersResponseBodyMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
}

func (s ListGroupStaffMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupStaffMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupStaffMembersResponseBody) SetMembers(v []*ListGroupStaffMembersResponseBodyMembers) *ListGroupStaffMembersResponseBody {
	s.Members = v
	return s
}

type ListGroupStaffMembersResponseBodyMembers struct {
	Nick *string `json:"nick,omitempty" xml:"nick,omitempty"`
	// This parameter is required.
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ListGroupStaffMembersResponseBodyMembers) String() string {
	return tea.Prettify(s)
}

func (s ListGroupStaffMembersResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *ListGroupStaffMembersResponseBodyMembers) SetNick(v string) *ListGroupStaffMembersResponseBodyMembers {
	s.Nick = &v
	return s
}

func (s *ListGroupStaffMembersResponseBodyMembers) SetUid(v string) *ListGroupStaffMembersResponseBodyMembers {
	s.Uid = &v
	return s
}

type ListGroupStaffMembersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupStaffMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupStaffMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupStaffMembersResponse) GoString() string {
	return s.String()
}

func (s *ListGroupStaffMembersResponse) SetHeaders(v map[string]*string) *ListGroupStaffMembersResponse {
	s.Headers = v
	return s
}

func (s *ListGroupStaffMembersResponse) SetStatusCode(v int32) *ListGroupStaffMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupStaffMembersResponse) SetBody(v *ListGroupStaffMembersResponseBody) *ListGroupStaffMembersResponse {
	s.Body = v
	return s
}

type QueryBatchSendResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryBatchSendResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchSendResultHeaders) GoString() string {
	return s.String()
}

func (s *QueryBatchSendResultHeaders) SetCommonHeaders(v map[string]*string) *QueryBatchSendResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryBatchSendResultHeaders) SetXAcsDingtalkAccessToken(v string) *QueryBatchSendResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryBatchSendResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 234#dingxxx@dingding
	SenderUserId *string `json:"senderUserId,omitempty" xml:"senderUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567890
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s QueryBatchSendResultRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchSendResultRequest) GoString() string {
	return s.String()
}

func (s *QueryBatchSendResultRequest) SetSenderUserId(v string) *QueryBatchSendResultRequest {
	s.SenderUserId = &v
	return s
}

func (s *QueryBatchSendResultRequest) SetTaskId(v string) *QueryBatchSendResultRequest {
	s.TaskId = &v
	return s
}

type QueryBatchSendResultResponseBody struct {
	// This parameter is required.
	Results []*QueryBatchSendResultResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryBatchSendResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchSendResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBatchSendResultResponseBody) SetResults(v []*QueryBatchSendResultResponseBodyResults) *QueryBatchSendResultResponseBody {
	s.Results = v
	return s
}

func (s *QueryBatchSendResultResponseBody) SetStatus(v int32) *QueryBatchSendResultResponseBody {
	s.Status = &v
	return s
}

type QueryBatchSendResultResponseBodyResults struct {
	// example:
	//
	// 123@channel
	AppUid *string `json:"appUid,omitempty" xml:"appUid,omitempty"`
	// example:
	//
	// cid1234567890==
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// example:
	//
	// 0
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// success
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// msg1234567890==
	MsgId *string `json:"msgId,omitempty" xml:"msgId,omitempty"`
}

func (s QueryBatchSendResultResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchSendResultResponseBodyResults) GoString() string {
	return s.String()
}

func (s *QueryBatchSendResultResponseBodyResults) SetAppUid(v string) *QueryBatchSendResultResponseBodyResults {
	s.AppUid = &v
	return s
}

func (s *QueryBatchSendResultResponseBodyResults) SetConversationId(v string) *QueryBatchSendResultResponseBodyResults {
	s.ConversationId = &v
	return s
}

func (s *QueryBatchSendResultResponseBodyResults) SetErrorCode(v string) *QueryBatchSendResultResponseBodyResults {
	s.ErrorCode = &v
	return s
}

func (s *QueryBatchSendResultResponseBodyResults) SetErrorMessage(v string) *QueryBatchSendResultResponseBodyResults {
	s.ErrorMessage = &v
	return s
}

func (s *QueryBatchSendResultResponseBodyResults) SetMsgId(v string) *QueryBatchSendResultResponseBodyResults {
	s.MsgId = &v
	return s
}

type QueryBatchSendResultResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBatchSendResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBatchSendResultResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchSendResultResponse) GoString() string {
	return s.String()
}

func (s *QueryBatchSendResultResponse) SetHeaders(v map[string]*string) *QueryBatchSendResultResponse {
	s.Headers = v
	return s
}

func (s *QueryBatchSendResultResponse) SetStatusCode(v int32) *QueryBatchSendResultResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBatchSendResultResponse) SetBody(v *QueryBatchSendResultResponseBody) *QueryBatchSendResultResponse {
	s.Body = v
	return s
}

type ReadMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	OperationSource         *string            `json:"operationSource,omitempty" xml:"operationSource,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ReadMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageHeaders) GoString() string {
	return s.String()
}

func (s *ReadMessageHeaders) SetCommonHeaders(v map[string]*string) *ReadMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReadMessageHeaders) SetOperationSource(v string) *ReadMessageHeaders {
	s.OperationSource = &v
	return s
}

func (s *ReadMessageHeaders) SetXAcsDingtalkAccessToken(v string) *ReadMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ReadMessageRequest struct {
	// This parameter is required.
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// This parameter is required.
	OperatorUid *string `json:"operatorUid,omitempty" xml:"operatorUid,omitempty"`
}

func (s ReadMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageRequest) GoString() string {
	return s.String()
}

func (s *ReadMessageRequest) SetMessageId(v string) *ReadMessageRequest {
	s.MessageId = &v
	return s
}

func (s *ReadMessageRequest) SetOperatorUid(v string) *ReadMessageRequest {
	s.OperatorUid = &v
	return s
}

type ReadMessageResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s ReadMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageResponse) GoString() string {
	return s.String()
}

func (s *ReadMessageResponse) SetHeaders(v map[string]*string) *ReadMessageResponse {
	s.Headers = v
	return s
}

func (s *ReadMessageResponse) SetStatusCode(v int32) *ReadMessageResponse {
	s.StatusCode = &v
	return s
}

type RecallMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	OperationSource         *string            `json:"operationSource,omitempty" xml:"operationSource,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RecallMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s RecallMessageHeaders) GoString() string {
	return s.String()
}

func (s *RecallMessageHeaders) SetCommonHeaders(v map[string]*string) *RecallMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RecallMessageHeaders) SetOperationSource(v string) *RecallMessageHeaders {
	s.OperationSource = &v
	return s
}

func (s *RecallMessageHeaders) SetXAcsDingtalkAccessToken(v string) *RecallMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RecallMessageRequest struct {
	// This parameter is required.
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// This parameter is required.
	OperatorUid *string `json:"operatorUid,omitempty" xml:"operatorUid,omitempty"`
	Type        *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RecallMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s RecallMessageRequest) GoString() string {
	return s.String()
}

func (s *RecallMessageRequest) SetMessageId(v string) *RecallMessageRequest {
	s.MessageId = &v
	return s
}

func (s *RecallMessageRequest) SetOperatorUid(v string) *RecallMessageRequest {
	s.OperatorUid = &v
	return s
}

func (s *RecallMessageRequest) SetType(v int32) *RecallMessageRequest {
	s.Type = &v
	return s
}

type RecallMessageResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s RecallMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s RecallMessageResponse) GoString() string {
	return s.String()
}

func (s *RecallMessageResponse) SetHeaders(v map[string]*string) *RecallMessageResponse {
	s.Headers = v
	return s
}

func (s *RecallMessageResponse) SetStatusCode(v int32) *RecallMessageResponse {
	s.StatusCode = &v
	return s
}

type RemoveGroupMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	OperationSource         *string            `json:"operationSource,omitempty" xml:"operationSource,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RemoveGroupMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupMembersHeaders) GoString() string {
	return s.String()
}

func (s *RemoveGroupMembersHeaders) SetCommonHeaders(v map[string]*string) *RemoveGroupMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveGroupMembersHeaders) SetOperationSource(v string) *RemoveGroupMembersHeaders {
	s.OperationSource = &v
	return s
}

func (s *RemoveGroupMembersHeaders) SetXAcsDingtalkAccessToken(v string) *RemoveGroupMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RemoveGroupMembersRequest struct {
	// This parameter is required.
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// true
	MemberUids  []*string `json:"memberUids,omitempty" xml:"memberUids,omitempty" type:"Repeated"`
	OperatorUid *string   `json:"operatorUid,omitempty" xml:"operatorUid,omitempty"`
}

func (s RemoveGroupMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *RemoveGroupMembersRequest) SetConversationId(v string) *RemoveGroupMembersRequest {
	s.ConversationId = &v
	return s
}

func (s *RemoveGroupMembersRequest) SetMemberUids(v []*string) *RemoveGroupMembersRequest {
	s.MemberUids = v
	return s
}

func (s *RemoveGroupMembersRequest) SetOperatorUid(v string) *RemoveGroupMembersRequest {
	s.OperatorUid = &v
	return s
}

type RemoveGroupMembersResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s RemoveGroupMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupMembersResponse) GoString() string {
	return s.String()
}

func (s *RemoveGroupMembersResponse) SetHeaders(v map[string]*string) *RemoveGroupMembersResponse {
	s.Headers = v
	return s
}

func (s *RemoveGroupMembersResponse) SetStatusCode(v int32) *RemoveGroupMembersResponse {
	s.StatusCode = &v
	return s
}

type SendMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	OperationSource         *string            `json:"operationSource,omitempty" xml:"operationSource,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendMessageHeaders) GoString() string {
	return s.String()
}

func (s *SendMessageHeaders) SetCommonHeaders(v map[string]*string) *SendMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendMessageHeaders) SetOperationSource(v string) *SendMessageHeaders {
	s.OperationSource = &v
	return s
}

func (s *SendMessageHeaders) SetXAcsDingtalkAccessToken(v string) *SendMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendMessageRequest struct {
	// This parameter is required.
	Content        *string `json:"content,omitempty" xml:"content,omitempty"`
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	CreateTime     *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	ReceiverUid    *string `json:"receiverUid,omitempty" xml:"receiverUid,omitempty"`
	// This parameter is required.
	SenderUid *string `json:"senderUid,omitempty" xml:"senderUid,omitempty"`
	// This parameter is required.
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s SendMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequest) SetContent(v string) *SendMessageRequest {
	s.Content = &v
	return s
}

func (s *SendMessageRequest) SetConversationId(v string) *SendMessageRequest {
	s.ConversationId = &v
	return s
}

func (s *SendMessageRequest) SetCreateTime(v int64) *SendMessageRequest {
	s.CreateTime = &v
	return s
}

func (s *SendMessageRequest) SetReceiverUid(v string) *SendMessageRequest {
	s.ReceiverUid = &v
	return s
}

func (s *SendMessageRequest) SetSenderUid(v string) *SendMessageRequest {
	s.SenderUid = &v
	return s
}

func (s *SendMessageRequest) SetUuid(v string) *SendMessageRequest {
	s.Uuid = &v
	return s
}

type SendMessageResponseBody struct {
	// This parameter is required.
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// This parameter is required.
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// This parameter is required.
	MsgId *string `json:"msgId,omitempty" xml:"msgId,omitempty"`
}

func (s SendMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageResponseBody) SetCreateTime(v int64) *SendMessageResponseBody {
	s.CreateTime = &v
	return s
}

func (s *SendMessageResponseBody) SetMessageId(v string) *SendMessageResponseBody {
	s.MessageId = &v
	return s
}

func (s *SendMessageResponseBody) SetMsgId(v string) *SendMessageResponseBody {
	s.MsgId = &v
	return s
}

type SendMessageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponse) GoString() string {
	return s.String()
}

func (s *SendMessageResponse) SetHeaders(v map[string]*string) *SendMessageResponse {
	s.Headers = v
	return s
}

func (s *SendMessageResponse) SetStatusCode(v int32) *SendMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMessageResponse) SetBody(v *SendMessageResponseBody) *SendMessageResponse {
	s.Body = v
	return s
}

type SendRobotMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendRobotMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendRobotMessageHeaders) GoString() string {
	return s.String()
}

func (s *SendRobotMessageHeaders) SetCommonHeaders(v map[string]*string) *SendRobotMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendRobotMessageHeaders) SetXAcsDingtalkAccessToken(v string) *SendRobotMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendRobotMessageRequest struct {
	// example:
	//
	// false
	AtAll      *bool     `json:"atAll,omitempty" xml:"atAll,omitempty"`
	AtAppUids  []*string `json:"atAppUids,omitempty" xml:"atAppUids,omitempty" type:"Repeated"`
	AtMobiles  []*string `json:"atMobiles,omitempty" xml:"atMobiles,omitempty" type:"Repeated"`
	AtUnionIds []*string `json:"atUnionIds,omitempty" xml:"atUnionIds,omitempty" type:"Repeated"`
	AtUsers    []*string `json:"atUsers,omitempty" xml:"atUsers,omitempty" type:"Repeated"`
	// example:
	//
	// 123
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// example:
	//
	// {\"pic1\":\"@123\",\"pic2\":\"@456\"}
	MsgMediaIdParamMap map[string]interface{} `json:"msgMediaIdParamMap,omitempty" xml:"msgMediaIdParamMap,omitempty"`
	// example:
	//
	// {\"text1\":\"hello\",\"text2\":\"world\"}
	MsgParamMap map[string]interface{} `json:"msgParamMap,omitempty" xml:"msgParamMap,omitempty"`
	// example:
	//
	// 123
	MsgTemplateId    *string   `json:"msgTemplateId,omitempty" xml:"msgTemplateId,omitempty"`
	ReceiverAppUids  []*string `json:"receiverAppUids,omitempty" xml:"receiverAppUids,omitempty" type:"Repeated"`
	ReceiverMobiles  []*string `json:"receiverMobiles,omitempty" xml:"receiverMobiles,omitempty" type:"Repeated"`
	ReceiverUnionIds []*string `json:"receiverUnionIds,omitempty" xml:"receiverUnionIds,omitempty" type:"Repeated"`
	ReceiverUserIds  []*string `json:"receiverUserIds,omitempty" xml:"receiverUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// 123
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	// example:
	//
	// 123
	TargetOpenConversationId *string `json:"targetOpenConversationId,omitempty" xml:"targetOpenConversationId,omitempty"`
}

func (s SendRobotMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendRobotMessageRequest) GoString() string {
	return s.String()
}

func (s *SendRobotMessageRequest) SetAtAll(v bool) *SendRobotMessageRequest {
	s.AtAll = &v
	return s
}

func (s *SendRobotMessageRequest) SetAtAppUids(v []*string) *SendRobotMessageRequest {
	s.AtAppUids = v
	return s
}

func (s *SendRobotMessageRequest) SetAtMobiles(v []*string) *SendRobotMessageRequest {
	s.AtMobiles = v
	return s
}

func (s *SendRobotMessageRequest) SetAtUnionIds(v []*string) *SendRobotMessageRequest {
	s.AtUnionIds = v
	return s
}

func (s *SendRobotMessageRequest) SetAtUsers(v []*string) *SendRobotMessageRequest {
	s.AtUsers = v
	return s
}

func (s *SendRobotMessageRequest) SetChannel(v string) *SendRobotMessageRequest {
	s.Channel = &v
	return s
}

func (s *SendRobotMessageRequest) SetMsgMediaIdParamMap(v map[string]interface{}) *SendRobotMessageRequest {
	s.MsgMediaIdParamMap = v
	return s
}

func (s *SendRobotMessageRequest) SetMsgParamMap(v map[string]interface{}) *SendRobotMessageRequest {
	s.MsgParamMap = v
	return s
}

func (s *SendRobotMessageRequest) SetMsgTemplateId(v string) *SendRobotMessageRequest {
	s.MsgTemplateId = &v
	return s
}

func (s *SendRobotMessageRequest) SetReceiverAppUids(v []*string) *SendRobotMessageRequest {
	s.ReceiverAppUids = v
	return s
}

func (s *SendRobotMessageRequest) SetReceiverMobiles(v []*string) *SendRobotMessageRequest {
	s.ReceiverMobiles = v
	return s
}

func (s *SendRobotMessageRequest) SetReceiverUnionIds(v []*string) *SendRobotMessageRequest {
	s.ReceiverUnionIds = v
	return s
}

func (s *SendRobotMessageRequest) SetReceiverUserIds(v []*string) *SendRobotMessageRequest {
	s.ReceiverUserIds = v
	return s
}

func (s *SendRobotMessageRequest) SetRobotCode(v string) *SendRobotMessageRequest {
	s.RobotCode = &v
	return s
}

func (s *SendRobotMessageRequest) SetTargetOpenConversationId(v string) *SendRobotMessageRequest {
	s.TargetOpenConversationId = &v
	return s
}

type SendRobotMessageResponseBody struct {
	OpenMsgId *string `json:"openMsgId,omitempty" xml:"openMsgId,omitempty"`
}

func (s SendRobotMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendRobotMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendRobotMessageResponseBody) SetOpenMsgId(v string) *SendRobotMessageResponseBody {
	s.OpenMsgId = &v
	return s
}

type SendRobotMessageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendRobotMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendRobotMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendRobotMessageResponse) GoString() string {
	return s.String()
}

func (s *SendRobotMessageResponse) SetHeaders(v map[string]*string) *SendRobotMessageResponse {
	s.Headers = v
	return s
}

func (s *SendRobotMessageResponse) SetStatusCode(v int32) *SendRobotMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendRobotMessageResponse) SetBody(v *SendRobotMessageResponseBody) *SendRobotMessageResponse {
	s.Body = v
	return s
}

type UpdateGroupNameHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	OperationSource         *string            `json:"operationSource,omitempty" xml:"operationSource,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateGroupNameHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupNameHeaders) GoString() string {
	return s.String()
}

func (s *UpdateGroupNameHeaders) SetCommonHeaders(v map[string]*string) *UpdateGroupNameHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateGroupNameHeaders) SetOperationSource(v string) *UpdateGroupNameHeaders {
	s.OperationSource = &v
	return s
}

func (s *UpdateGroupNameHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateGroupNameHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateGroupNameRequest struct {
	// This parameter is required.
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// This parameter is required.
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	OperatorUid *string `json:"operatorUid,omitempty" xml:"operatorUid,omitempty"`
}

func (s UpdateGroupNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupNameRequest) SetConversationId(v string) *UpdateGroupNameRequest {
	s.ConversationId = &v
	return s
}

func (s *UpdateGroupNameRequest) SetName(v string) *UpdateGroupNameRequest {
	s.Name = &v
	return s
}

func (s *UpdateGroupNameRequest) SetOperatorUid(v string) *UpdateGroupNameRequest {
	s.OperatorUid = &v
	return s
}

type UpdateGroupNameResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateGroupNameResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupNameResponse) SetHeaders(v map[string]*string) *UpdateGroupNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupNameResponse) SetStatusCode(v int32) *UpdateGroupNameResponse {
	s.StatusCode = &v
	return s
}

type UpdateGroupOwnerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateGroupOwnerHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupOwnerHeaders) GoString() string {
	return s.String()
}

func (s *UpdateGroupOwnerHeaders) SetCommonHeaders(v map[string]*string) *UpdateGroupOwnerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateGroupOwnerHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateGroupOwnerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateGroupOwnerRequest struct {
	// This parameter is required.
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	OperatorUid    *string `json:"operatorUid,omitempty" xml:"operatorUid,omitempty"`
	// This parameter is required.
	OwnerUid *string `json:"ownerUid,omitempty" xml:"ownerUid,omitempty"`
}

func (s UpdateGroupOwnerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupOwnerRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupOwnerRequest) SetConversationId(v string) *UpdateGroupOwnerRequest {
	s.ConversationId = &v
	return s
}

func (s *UpdateGroupOwnerRequest) SetOperatorUid(v string) *UpdateGroupOwnerRequest {
	s.OperatorUid = &v
	return s
}

func (s *UpdateGroupOwnerRequest) SetOwnerUid(v string) *UpdateGroupOwnerRequest {
	s.OwnerUid = &v
	return s
}

type UpdateGroupOwnerResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateGroupOwnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupOwnerResponseBody) SetResult(v bool) *UpdateGroupOwnerResponseBody {
	s.Result = &v
	return s
}

type UpdateGroupOwnerResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGroupOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGroupOwnerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupOwnerResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupOwnerResponse) SetHeaders(v map[string]*string) *UpdateGroupOwnerResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupOwnerResponse) SetStatusCode(v int32) *UpdateGroupOwnerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupOwnerResponse) SetBody(v *UpdateGroupOwnerResponseBody) *UpdateGroupOwnerResponse {
	s.Body = v
	return s
}

type UploadFileHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UploadFileHeaders) String() string {
	return tea.Prettify(s)
}

func (s UploadFileHeaders) GoString() string {
	return s.String()
}

func (s *UploadFileHeaders) SetCommonHeaders(v map[string]*string) *UploadFileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UploadFileHeaders) SetXAcsDingtalkAccessToken(v string) *UploadFileHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UploadFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 111.png
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// image
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// This parameter is required.
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111@dingdingkelian
	SenderUid *string `json:"senderUid,omitempty" xml:"senderUid,omitempty"`
}

func (s UploadFileRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadFileRequest) GoString() string {
	return s.String()
}

func (s *UploadFileRequest) SetFileName(v string) *UploadFileRequest {
	s.FileName = &v
	return s
}

func (s *UploadFileRequest) SetFileType(v string) *UploadFileRequest {
	s.FileType = &v
	return s
}

func (s *UploadFileRequest) SetFileUrl(v string) *UploadFileRequest {
	s.FileUrl = &v
	return s
}

func (s *UploadFileRequest) SetSenderUid(v string) *UploadFileRequest {
	s.SenderUid = &v
	return s
}

type UploadFileResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// $csvsvsdvsvdscs
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
}

func (s UploadFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadFileResponseBody) GoString() string {
	return s.String()
}

func (s *UploadFileResponseBody) SetMediaId(v string) *UploadFileResponseBody {
	s.MediaId = &v
	return s
}

type UploadFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadFileResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadFileResponse) GoString() string {
	return s.String()
}

func (s *UploadFileResponse) SetHeaders(v map[string]*string) *UploadFileResponse {
	s.Headers = v
	return s
}

func (s *UploadFileResponse) SetStatusCode(v int32) *UploadFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadFileResponse) SetBody(v *UploadFileResponseBody) *UploadFileResponse {
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
// 添加群成员
//
// @param request - AddGroupMembersRequest
//
// @param headers - AddGroupMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGroupMembersResponse
func (client *Client) AddGroupMembersWithOptions(request *AddGroupMembersRequest, headers *AddGroupMembersHeaders, runtime *util.RuntimeOptions) (_result *AddGroupMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		body["conversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUid)) {
		body["operatorUid"] = request.OperatorUid
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.OperationSource)) {
		realHeaders["operationSource"] = util.ToJSONString(headers.OperationSource)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddGroupMembers"),
		Version:     tea.String("impaas_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/impaas/interconnections/groups/members/batchAdd"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddGroupMembersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加群成员
//
// @param request - AddGroupMembersRequest
//
// @return AddGroupMembersResponse
func (client *Client) AddGroupMembers(request *AddGroupMembersRequest) (_result *AddGroupMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddGroupMembersHeaders{}
	_result = &AddGroupMembersResponse{}
	_body, _err := client.AddGroupMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 外部用户导入profile
//
// @param request - AddProfileRequest
//
// @param headers - AddProfileHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddProfileResponse
func (client *Client) AddProfileWithOptions(request *AddProfileRequest, headers *AddProfileHeaders, runtime *util.RuntimeOptions) (_result *AddProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUid)) {
		body["appUid"] = request.AppUid
	}

	if !tea.BoolValue(util.IsUnset(request.AvatarMediaId)) {
		body["avatarMediaId"] = request.AvatarMediaId
	}

	if !tea.BoolValue(util.IsUnset(request.MobileNumber)) {
		body["mobileNumber"] = request.MobileNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Nick)) {
		body["nick"] = request.Nick
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
		Action:      tea.String("AddProfile"),
		Version:     tea.String("impaas_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/impaas/interconnections/users/profiles"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("none"),
	}
	_result = &AddProfileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 外部用户导入profile
//
// @param request - AddProfileRequest
//
// @return AddProfileResponse
func (client *Client) AddProfile(request *AddProfileRequest) (_result *AddProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddProfileHeaders{}
	_result = &AddProfileResponse{}
	_body, _err := client.AddProfileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 消息批量接口
//
// @param request - BatchSendRequest
//
// @param headers - BatchSendHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSendResponse
func (client *Client) BatchSendWithOptions(request *BatchSendRequest, headers *BatchSendHeaders, runtime *util.RuntimeOptions) (_result *BatchSendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUids)) {
		body["appUids"] = request.AppUids
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ConversationIds)) {
		body["conversationIds"] = request.ConversationIds
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
	params := &openapi.Params{
		Action:      tea.String("BatchSend"),
		Version:     tea.String("impaas_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/impaas/interconnections/messages/batchSend"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchSendResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 消息批量接口
//
// @param request - BatchSendRequest
//
// @return BatchSendResponse
func (client *Client) BatchSend(request *BatchSendRequest) (_result *BatchSendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchSendHeaders{}
	_result = &BatchSendResponse{}
	_body, _err := client.BatchSendWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建群
//
// @param request - CreateGroupRequest
//
// @param headers - CreateGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupResponse
func (client *Client) CreateGroupWithOptions(request *CreateGroupRequest, headers *CreateGroupHeaders, runtime *util.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		body["channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorUid)) {
		body["creatorUid"] = request.CreatorUid
	}

	if !tea.BoolValue(util.IsUnset(request.IconMediaId)) {
		body["iconMediaId"] = request.IconMediaId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Properties)) {
		body["properties"] = request.Properties
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.OperationSource)) {
		realHeaders["operationSource"] = util.ToJSONString(headers.OperationSource)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGroup"),
		Version:     tea.String("impaas_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/impaas/interconnections/groups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建群
//
// @param request - CreateGroupRequest
//
// @return CreateGroupResponse
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

// Summary:
//
// 创建托管账号为群主的群
//
// @param request - CreateTrustGroupRequest
//
// @param headers - CreateTrustGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrustGroupResponse
func (client *Client) CreateTrustGroupWithOptions(request *CreateTrustGroupRequest, headers *CreateTrustGroupHeaders, runtime *util.RuntimeOptions) (_result *CreateTrustGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		body["channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.IconMediaId)) {
		body["iconMediaId"] = request.IconMediaId
	}

	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Properties)) {
		body["properties"] = request.Properties
	}

	if !tea.BoolValue(util.IsUnset(request.SystemMsg)) {
		body["systemMsg"] = request.SystemMsg
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.OperationSource)) {
		realHeaders["operationSource"] = util.ToJSONString(headers.OperationSource)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTrustGroup"),
		Version:     tea.String("impaas_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/impaas/interconnections/groups/trusts"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTrustGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建托管账号为群主的群
//
// @param request - CreateTrustGroupRequest
//
// @return CreateTrustGroupResponse
func (client *Client) CreateTrustGroup(request *CreateTrustGroupRequest) (_result *CreateTrustGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTrustGroupHeaders{}
	_result = &CreateTrustGroupResponse{}
	_body, _err := client.CreateTrustGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解散群
//
// @param request - DismissGroupRequest
//
// @param headers - DismissGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DismissGroupResponse
func (client *Client) DismissGroupWithOptions(request *DismissGroupRequest, headers *DismissGroupHeaders, runtime *util.RuntimeOptions) (_result *DismissGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		body["conversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUid)) {
		body["operatorUid"] = request.OperatorUid
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.OperationSource)) {
		realHeaders["operationSource"] = util.ToJSONString(headers.OperationSource)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DismissGroup"),
		Version:     tea.String("impaas_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/impaas/interconnections/groups/dismiss"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DismissGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解散群
//
// @param request - DismissGroupRequest
//
// @return DismissGroupResponse
func (client *Client) DismissGroup(request *DismissGroupRequest) (_result *DismissGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DismissGroupHeaders{}
	_result = &DismissGroupResponse{}
	_body, _err := client.DismissGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生成单聊会话Id
//
// @param request - GetConversationIdRequest
//
// @param headers - GetConversationIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConversationIdResponse
func (client *Client) GetConversationIdWithOptions(request *GetConversationIdRequest, headers *GetConversationIdHeaders, runtime *util.RuntimeOptions) (_result *GetConversationIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUid)) {
		body["appUid"] = request.AppUid
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
	params := &openapi.Params{
		Action:      tea.String("GetConversationId"),
		Version:     tea.String("impaas_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/impaas/interconnections/conversations"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConversationIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成单聊会话Id
//
// @param request - GetConversationIdRequest
//
// @return GetConversationIdResponse
func (client *Client) GetConversationId(request *GetConversationIdRequest) (_result *GetConversationIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetConversationIdHeaders{}
	_result = &GetConversationIdResponse{}
	_body, _err := client.GetConversationIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 多媒体文件下载
//
// @param request - GetMediaUrlRequest
//
// @param headers - GetMediaUrlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaUrlResponse
func (client *Client) GetMediaUrlWithOptions(request *GetMediaUrlRequest, headers *GetMediaUrlHeaders, runtime *util.RuntimeOptions) (_result *GetMediaUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MediaId)) {
		body["mediaId"] = request.MediaId
	}

	if !tea.BoolValue(util.IsUnset(request.UrlExpireTime)) {
		body["urlExpireTime"] = request.UrlExpireTime
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
		Action:      tea.String("GetMediaUrl"),
		Version:     tea.String("impaas_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/impaas/interconnections/medium/urls"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMediaUrlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 多媒体文件下载
//
// @param request - GetMediaUrlRequest
//
// @return GetMediaUrlResponse
func (client *Client) GetMediaUrl(request *GetMediaUrlRequest) (_result *GetMediaUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMediaUrlHeaders{}
	_result = &GetMediaUrlResponse{}
	_body, _err := client.GetMediaUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 多媒体文件批量下载
//
// @param request - GetMediaUrlsRequest
//
// @param headers - GetMediaUrlsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaUrlsResponse
func (client *Client) GetMediaUrlsWithOptions(request *GetMediaUrlsRequest, headers *GetMediaUrlsHeaders, runtime *util.RuntimeOptions) (_result *GetMediaUrlsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MediaIds)) {
		body["mediaIds"] = request.MediaIds
	}

	if !tea.BoolValue(util.IsUnset(request.UrlExpireTime)) {
		body["urlExpireTime"] = request.UrlExpireTime
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
		Action:      tea.String("GetMediaUrls"),
		Version:     tea.String("impaas_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/impaas/interconnections/mediaUrls/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMediaUrlsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 多媒体文件批量下载
//
// @param request - GetMediaUrlsRequest
//
// @return GetMediaUrlsResponse
func (client *Client) GetMediaUrls(request *GetMediaUrlsRequest) (_result *GetMediaUrlsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMediaUrlsHeaders{}
	_result = &GetMediaUrlsResponse{}
	_body, _err := client.GetMediaUrlsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取钉盘文件下载信息
//
// @param request - GetSpaceFileUrlRequest
//
// @param headers - GetSpaceFileUrlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSpaceFileUrlResponse
func (client *Client) GetSpaceFileUrlWithOptions(request *GetSpaceFileUrlRequest, headers *GetSpaceFileUrlHeaders, runtime *util.RuntimeOptions) (_result *GetSpaceFileUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		query["fileId"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.SenderUid)) {
		query["senderUid"] = request.SenderUid
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		query["spaceId"] = request.SpaceId
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
		Action:      tea.String("GetSpaceFileUrl"),
		Version:     tea.String("impaas_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/impaas/interconnections/spaces/files/urls"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSpaceFileUrlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取钉盘文件下载信息
//
// @param request - GetSpaceFileUrlRequest
//
// @return GetSpaceFileUrlResponse
func (client *Client) GetSpaceFileUrl(request *GetSpaceFileUrlRequest) (_result *GetSpaceFileUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSpaceFileUrlHeaders{}
	_result = &GetSpaceFileUrlResponse{}
	_body, _err := client.GetSpaceFileUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业员工类型的群成员
//
// @param request - ListGroupStaffMembersRequest
//
// @param headers - ListGroupStaffMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupStaffMembersResponse
func (client *Client) ListGroupStaffMembersWithOptions(request *ListGroupStaffMembersRequest, headers *ListGroupStaffMembersHeaders, runtime *util.RuntimeOptions) (_result *ListGroupStaffMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		body["conversationId"] = request.ConversationId
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
		Action:      tea.String("ListGroupStaffMembers"),
		Version:     tea.String("impaas_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/impaas/interconnections/groups/staffMemers/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGroupStaffMembersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业员工类型的群成员
//
// @param request - ListGroupStaffMembersRequest
//
// @return ListGroupStaffMembersResponse
func (client *Client) ListGroupStaffMembers(request *ListGroupStaffMembersRequest) (_result *ListGroupStaffMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListGroupStaffMembersHeaders{}
	_result = &ListGroupStaffMembersResponse{}
	_body, _err := client.ListGroupStaffMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询batchSend结果
//
// @param request - QueryBatchSendResultRequest
//
// @param headers - QueryBatchSendResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBatchSendResultResponse
func (client *Client) QueryBatchSendResultWithOptions(request *QueryBatchSendResultRequest, headers *QueryBatchSendResultHeaders, runtime *util.RuntimeOptions) (_result *QueryBatchSendResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SenderUserId)) {
		query["senderUserId"] = request.SenderUserId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
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
		Action:      tea.String("QueryBatchSendResult"),
		Version:     tea.String("impaas_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/impaas/interconnections/messages/batchSendResults"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBatchSendResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询batchSend结果
//
// @param request - QueryBatchSendResultRequest
//
// @return QueryBatchSendResultResponse
func (client *Client) QueryBatchSendResult(request *QueryBatchSendResultRequest) (_result *QueryBatchSendResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryBatchSendResultHeaders{}
	_result = &QueryBatchSendResultResponse{}
	_body, _err := client.QueryBatchSendResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 消息已读
//
// @param request - ReadMessageRequest
//
// @param headers - ReadMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadMessageResponse
func (client *Client) ReadMessageWithOptions(request *ReadMessageRequest, headers *ReadMessageHeaders, runtime *util.RuntimeOptions) (_result *ReadMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		body["messageId"] = request.MessageId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUid)) {
		body["operatorUid"] = request.OperatorUid
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.OperationSource)) {
		realHeaders["operationSource"] = util.ToJSONString(headers.OperationSource)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReadMessage"),
		Version:     tea.String("impaas_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/impaas/interconnections/messages/read"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("none"),
	}
	_result = &ReadMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 消息已读
//
// @param request - ReadMessageRequest
//
// @return ReadMessageResponse
func (client *Client) ReadMessage(request *ReadMessageRequest) (_result *ReadMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ReadMessageHeaders{}
	_result = &ReadMessageResponse{}
	_body, _err := client.ReadMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 消息撤回
//
// @param request - RecallMessageRequest
//
// @param headers - RecallMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecallMessageResponse
func (client *Client) RecallMessageWithOptions(request *RecallMessageRequest, headers *RecallMessageHeaders, runtime *util.RuntimeOptions) (_result *RecallMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		body["messageId"] = request.MessageId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUid)) {
		body["operatorUid"] = request.OperatorUid
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.OperationSource)) {
		realHeaders["operationSource"] = util.ToJSONString(headers.OperationSource)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecallMessage"),
		Version:     tea.String("impaas_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/impaas/interconnections/messages/recall"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("none"),
	}
	_result = &RecallMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 消息撤回
//
// @param request - RecallMessageRequest
//
// @return RecallMessageResponse
func (client *Client) RecallMessage(request *RecallMessageRequest) (_result *RecallMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RecallMessageHeaders{}
	_result = &RecallMessageResponse{}
	_body, _err := client.RecallMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移除群成员
//
// @param request - RemoveGroupMembersRequest
//
// @param headers - RemoveGroupMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveGroupMembersResponse
func (client *Client) RemoveGroupMembersWithOptions(request *RemoveGroupMembersRequest, headers *RemoveGroupMembersHeaders, runtime *util.RuntimeOptions) (_result *RemoveGroupMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		body["conversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUids)) {
		body["memberUids"] = request.MemberUids
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUid)) {
		body["operatorUid"] = request.OperatorUid
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.OperationSource)) {
		realHeaders["operationSource"] = util.ToJSONString(headers.OperationSource)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveGroupMembers"),
		Version:     tea.String("impaas_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/impaas/interconnections/groups/members/batchRemove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveGroupMembersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移除群成员
//
// @param request - RemoveGroupMembersRequest
//
// @return RemoveGroupMembersResponse
func (client *Client) RemoveGroupMembers(request *RemoveGroupMembersRequest) (_result *RemoveGroupMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveGroupMembersHeaders{}
	_result = &RemoveGroupMembersResponse{}
	_body, _err := client.RemoveGroupMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 消息发送
//
// @param request - SendMessageRequest
//
// @param headers - SendMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendMessageResponse
func (client *Client) SendMessageWithOptions(request *SendMessageRequest, headers *SendMessageHeaders, runtime *util.RuntimeOptions) (_result *SendMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		body["conversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTime)) {
		body["createTime"] = request.CreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUid)) {
		body["receiverUid"] = request.ReceiverUid
	}

	if !tea.BoolValue(util.IsUnset(request.SenderUid)) {
		body["senderUid"] = request.SenderUid
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.OperationSource)) {
		realHeaders["operationSource"] = util.ToJSONString(headers.OperationSource)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendMessage"),
		Version:     tea.String("impaas_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/impaas/interconnections/messages/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 消息发送
//
// @param request - SendMessageRequest
//
// @return SendMessageResponse
func (client *Client) SendMessage(request *SendMessageRequest) (_result *SendMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendMessageHeaders{}
	_result = &SendMessageResponse{}
	_body, _err := client.SendMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过群模板机器人发送消息
//
// @param request - SendRobotMessageRequest
//
// @param headers - SendRobotMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendRobotMessageResponse
func (client *Client) SendRobotMessageWithOptions(request *SendRobotMessageRequest, headers *SendRobotMessageHeaders, runtime *util.RuntimeOptions) (_result *SendRobotMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AtAll)) {
		body["atAll"] = request.AtAll
	}

	if !tea.BoolValue(util.IsUnset(request.AtAppUids)) {
		body["atAppUids"] = request.AtAppUids
	}

	if !tea.BoolValue(util.IsUnset(request.AtMobiles)) {
		body["atMobiles"] = request.AtMobiles
	}

	if !tea.BoolValue(util.IsUnset(request.AtUnionIds)) {
		body["atUnionIds"] = request.AtUnionIds
	}

	if !tea.BoolValue(util.IsUnset(request.AtUsers)) {
		body["atUsers"] = request.AtUsers
	}

	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		body["channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.MsgMediaIdParamMap)) {
		body["msgMediaIdParamMap"] = request.MsgMediaIdParamMap
	}

	if !tea.BoolValue(util.IsUnset(request.MsgParamMap)) {
		body["msgParamMap"] = request.MsgParamMap
	}

	if !tea.BoolValue(util.IsUnset(request.MsgTemplateId)) {
		body["msgTemplateId"] = request.MsgTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverAppUids)) {
		body["receiverAppUids"] = request.ReceiverAppUids
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverMobiles)) {
		body["receiverMobiles"] = request.ReceiverMobiles
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUnionIds)) {
		body["receiverUnionIds"] = request.ReceiverUnionIds
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUserIds)) {
		body["receiverUserIds"] = request.ReceiverUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
	}

	if !tea.BoolValue(util.IsUnset(request.TargetOpenConversationId)) {
		body["targetOpenConversationId"] = request.TargetOpenConversationId
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
		Action:      tea.String("SendRobotMessage"),
		Version:     tea.String("impaas_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/impaas/interconnections/robots/messages/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendRobotMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过群模板机器人发送消息
//
// @param request - SendRobotMessageRequest
//
// @return SendRobotMessageResponse
func (client *Client) SendRobotMessage(request *SendRobotMessageRequest) (_result *SendRobotMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendRobotMessageHeaders{}
	_result = &SendRobotMessageResponse{}
	_body, _err := client.SendRobotMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改群名称
//
// @param request - UpdateGroupNameRequest
//
// @param headers - UpdateGroupNameHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupNameResponse
func (client *Client) UpdateGroupNameWithOptions(request *UpdateGroupNameRequest, headers *UpdateGroupNameHeaders, runtime *util.RuntimeOptions) (_result *UpdateGroupNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		body["conversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUid)) {
		body["operatorUid"] = request.OperatorUid
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.OperationSource)) {
		realHeaders["operationSource"] = util.ToJSONString(headers.OperationSource)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGroupName"),
		Version:     tea.String("impaas_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/impaas/interconnections/groups/names"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGroupNameResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改群名称
//
// @param request - UpdateGroupNameRequest
//
// @return UpdateGroupNameResponse
func (client *Client) UpdateGroupName(request *UpdateGroupNameRequest) (_result *UpdateGroupNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateGroupNameHeaders{}
	_result = &UpdateGroupNameResponse{}
	_body, _err := client.UpdateGroupNameWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 转让群主
//
// @param request - UpdateGroupOwnerRequest
//
// @param headers - UpdateGroupOwnerHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupOwnerResponse
func (client *Client) UpdateGroupOwnerWithOptions(request *UpdateGroupOwnerRequest, headers *UpdateGroupOwnerHeaders, runtime *util.RuntimeOptions) (_result *UpdateGroupOwnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		body["conversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUid)) {
		body["operatorUid"] = request.OperatorUid
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUid)) {
		body["ownerUid"] = request.OwnerUid
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
		Action:      tea.String("UpdateGroupOwner"),
		Version:     tea.String("impaas_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/impaas/interconnections/groups/owners"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGroupOwnerResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 转让群主
//
// @param request - UpdateGroupOwnerRequest
//
// @return UpdateGroupOwnerResponse
func (client *Client) UpdateGroupOwner(request *UpdateGroupOwnerRequest) (_result *UpdateGroupOwnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateGroupOwnerHeaders{}
	_result = &UpdateGroupOwnerResponse{}
	_body, _err := client.UpdateGroupOwnerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 互联互通上传文件
//
// @param request - UploadFileRequest
//
// @param headers - UploadFileHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadFileResponse
func (client *Client) UploadFileWithOptions(request *UploadFileRequest, headers *UploadFileHeaders, runtime *util.RuntimeOptions) (_result *UploadFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		body["fileType"] = request.FileType
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		body["fileUrl"] = request.FileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SenderUid)) {
		body["senderUid"] = request.SenderUid
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
		Action:      tea.String("UploadFile"),
		Version:     tea.String("impaas_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/impaas/interconnections/files/upload"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 互联互通上传文件
//
// @param request - UploadFileRequest
//
// @return UploadFileResponse
func (client *Client) UploadFile(request *UploadFileRequest) (_result *UploadFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UploadFileHeaders{}
	_result = &UploadFileResponse{}
	_body, _err := client.UploadFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
