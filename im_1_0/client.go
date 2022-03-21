// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package im_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BatchQueryGroupMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchQueryGroupMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryGroupMemberHeaders) GoString() string {
	return s.String()
}

func (s *BatchQueryGroupMemberHeaders) SetCommonHeaders(v map[string]*string) *BatchQueryGroupMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchQueryGroupMemberHeaders) SetXAcsDingtalkAccessToken(v string) *BatchQueryGroupMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchQueryGroupMemberRequest struct {
	// 酷应用编码
	CoolAppCode *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	// 本次读取的最大数据记录数量（该入参传入值小于钉钉阈值时返回全部）
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 标记当前开始读取的位置，置空表示从头开始
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 开放群ID
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s BatchQueryGroupMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryGroupMemberRequest) SetCoolAppCode(v string) *BatchQueryGroupMemberRequest {
	s.CoolAppCode = &v
	return s
}

func (s *BatchQueryGroupMemberRequest) SetMaxResults(v int64) *BatchQueryGroupMemberRequest {
	s.MaxResults = &v
	return s
}

func (s *BatchQueryGroupMemberRequest) SetNextToken(v string) *BatchQueryGroupMemberRequest {
	s.NextToken = &v
	return s
}

func (s *BatchQueryGroupMemberRequest) SetOpenConversationId(v string) *BatchQueryGroupMemberRequest {
	s.OpenConversationId = &v
	return s
}

type BatchQueryGroupMemberResponseBody struct {
	// 是否还有更多数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 群成员员工号
	MemberUserIds []*string `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty" type:"Repeated"`
	// 下一次请求的游标
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// result
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchQueryGroupMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQueryGroupMemberResponseBody) SetHasMore(v bool) *BatchQueryGroupMemberResponseBody {
	s.HasMore = &v
	return s
}

func (s *BatchQueryGroupMemberResponseBody) SetMemberUserIds(v []*string) *BatchQueryGroupMemberResponseBody {
	s.MemberUserIds = v
	return s
}

func (s *BatchQueryGroupMemberResponseBody) SetNextToken(v string) *BatchQueryGroupMemberResponseBody {
	s.NextToken = &v
	return s
}

func (s *BatchQueryGroupMemberResponseBody) SetSuccess(v bool) *BatchQueryGroupMemberResponseBody {
	s.Success = &v
	return s
}

type BatchQueryGroupMemberResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchQueryGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchQueryGroupMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *BatchQueryGroupMemberResponse) SetHeaders(v map[string]*string) *BatchQueryGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *BatchQueryGroupMemberResponse) SetBody(v *BatchQueryGroupMemberResponseBody) *BatchQueryGroupMemberResponse {
	s.Body = v
	return s
}

type CardTemplateBuildActionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CardTemplateBuildActionHeaders) String() string {
	return tea.Prettify(s)
}

func (s CardTemplateBuildActionHeaders) GoString() string {
	return s.String()
}

func (s *CardTemplateBuildActionHeaders) SetCommonHeaders(v map[string]*string) *CardTemplateBuildActionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CardTemplateBuildActionHeaders) SetXAcsDingtalkAccessToken(v string) *CardTemplateBuildActionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CardTemplateBuildActionRequest struct {
	// 模板构建的action：含create、save、deploy
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 模板构建的dto对象
	CardTemplateJson *string `json:"cardTemplateJson,omitempty" xml:"cardTemplateJson,omitempty"`
}

func (s CardTemplateBuildActionRequest) String() string {
	return tea.Prettify(s)
}

func (s CardTemplateBuildActionRequest) GoString() string {
	return s.String()
}

func (s *CardTemplateBuildActionRequest) SetAction(v string) *CardTemplateBuildActionRequest {
	s.Action = &v
	return s
}

func (s *CardTemplateBuildActionRequest) SetCardTemplateJson(v string) *CardTemplateBuildActionRequest {
	s.CardTemplateJson = &v
	return s
}

type CardTemplateBuildActionResponseBody struct {
	// 模板构建的dto对象
	CardTemplateJson *string `json:"cardTemplateJson,omitempty" xml:"cardTemplateJson,omitempty"`
}

func (s CardTemplateBuildActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CardTemplateBuildActionResponseBody) GoString() string {
	return s.String()
}

func (s *CardTemplateBuildActionResponseBody) SetCardTemplateJson(v string) *CardTemplateBuildActionResponseBody {
	s.CardTemplateJson = &v
	return s
}

type CardTemplateBuildActionResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CardTemplateBuildActionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CardTemplateBuildActionResponse) String() string {
	return tea.Prettify(s)
}

func (s CardTemplateBuildActionResponse) GoString() string {
	return s.String()
}

func (s *CardTemplateBuildActionResponse) SetHeaders(v map[string]*string) *CardTemplateBuildActionResponse {
	s.Headers = v
	return s
}

func (s *CardTemplateBuildActionResponse) SetBody(v *CardTemplateBuildActionResponseBody) *CardTemplateBuildActionResponse {
	s.Body = v
	return s
}

type ChatIdToOpenConversationIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ChatIdToOpenConversationIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s ChatIdToOpenConversationIdHeaders) GoString() string {
	return s.String()
}

func (s *ChatIdToOpenConversationIdHeaders) SetCommonHeaders(v map[string]*string) *ChatIdToOpenConversationIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ChatIdToOpenConversationIdHeaders) SetXAcsDingtalkAccessToken(v string) *ChatIdToOpenConversationIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ChatIdToOpenConversationIdResponseBody struct {
	// openConversationId
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s ChatIdToOpenConversationIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChatIdToOpenConversationIdResponseBody) GoString() string {
	return s.String()
}

func (s *ChatIdToOpenConversationIdResponseBody) SetOpenConversationId(v string) *ChatIdToOpenConversationIdResponseBody {
	s.OpenConversationId = &v
	return s
}

type ChatIdToOpenConversationIdResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ChatIdToOpenConversationIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChatIdToOpenConversationIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ChatIdToOpenConversationIdResponse) GoString() string {
	return s.String()
}

func (s *ChatIdToOpenConversationIdResponse) SetHeaders(v map[string]*string) *ChatIdToOpenConversationIdResponse {
	s.Headers = v
	return s
}

func (s *ChatIdToOpenConversationIdResponse) SetBody(v *ChatIdToOpenConversationIdResponseBody) *ChatIdToOpenConversationIdResponse {
	s.Body = v
	return s
}

type GetInterconnectionUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetInterconnectionUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetInterconnectionUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetInterconnectionUrlHeaders) SetCommonHeaders(v map[string]*string) *GetInterconnectionUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInterconnectionUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GetInterconnectionUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetInterconnectionUrlRequest struct {
	// appUserAvatar
	AppUserAvatar *string `json:"appUserAvatar,omitempty" xml:"appUserAvatar,omitempty"`
	// appUserAvatarType
	AppUserAvatarType *int32 `json:"appUserAvatarType,omitempty" xml:"appUserAvatarType,omitempty"`
	// appUserId
	AppUserId *string `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
	// appUserMobileNumber
	AppUserMobileNumber *string `json:"appUserMobileNumber,omitempty" xml:"appUserMobileNumber,omitempty"`
	// appUserName
	AppUserName *string `json:"appUserName,omitempty" xml:"appUserName,omitempty"`
	// msgPageType
	MsgPageType *int32 `json:"msgPageType,omitempty" xml:"msgPageType,omitempty"`
	// qrCode
	QrCode *string `json:"qrCode,omitempty" xml:"qrCode,omitempty"`
	// signature
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	// sourceCode
	SourceCode *string `json:"sourceCode,omitempty" xml:"sourceCode,omitempty"`
	// sourceType
	SourceType *int32 `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetInterconnectionUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInterconnectionUrlRequest) GoString() string {
	return s.String()
}

func (s *GetInterconnectionUrlRequest) SetAppUserAvatar(v string) *GetInterconnectionUrlRequest {
	s.AppUserAvatar = &v
	return s
}

func (s *GetInterconnectionUrlRequest) SetAppUserAvatarType(v int32) *GetInterconnectionUrlRequest {
	s.AppUserAvatarType = &v
	return s
}

func (s *GetInterconnectionUrlRequest) SetAppUserId(v string) *GetInterconnectionUrlRequest {
	s.AppUserId = &v
	return s
}

func (s *GetInterconnectionUrlRequest) SetAppUserMobileNumber(v string) *GetInterconnectionUrlRequest {
	s.AppUserMobileNumber = &v
	return s
}

func (s *GetInterconnectionUrlRequest) SetAppUserName(v string) *GetInterconnectionUrlRequest {
	s.AppUserName = &v
	return s
}

func (s *GetInterconnectionUrlRequest) SetMsgPageType(v int32) *GetInterconnectionUrlRequest {
	s.MsgPageType = &v
	return s
}

func (s *GetInterconnectionUrlRequest) SetQrCode(v string) *GetInterconnectionUrlRequest {
	s.QrCode = &v
	return s
}

func (s *GetInterconnectionUrlRequest) SetSignature(v string) *GetInterconnectionUrlRequest {
	s.Signature = &v
	return s
}

func (s *GetInterconnectionUrlRequest) SetSourceCode(v string) *GetInterconnectionUrlRequest {
	s.SourceCode = &v
	return s
}

func (s *GetInterconnectionUrlRequest) SetSourceType(v int32) *GetInterconnectionUrlRequest {
	s.SourceType = &v
	return s
}

func (s *GetInterconnectionUrlRequest) SetUserId(v string) *GetInterconnectionUrlRequest {
	s.UserId = &v
	return s
}

type GetInterconnectionUrlResponseBody struct {
	// 会话url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetInterconnectionUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInterconnectionUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetInterconnectionUrlResponseBody) SetUrl(v string) *GetInterconnectionUrlResponseBody {
	s.Url = &v
	return s
}

type GetInterconnectionUrlResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInterconnectionUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInterconnectionUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInterconnectionUrlResponse) GoString() string {
	return s.String()
}

func (s *GetInterconnectionUrlResponse) SetHeaders(v map[string]*string) *GetInterconnectionUrlResponse {
	s.Headers = v
	return s
}

func (s *GetInterconnectionUrlResponse) SetBody(v *GetInterconnectionUrlResponseBody) *GetInterconnectionUrlResponse {
	s.Body = v
	return s
}

type GetSceneGroupInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSceneGroupInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSceneGroupInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetSceneGroupInfoHeaders) SetCommonHeaders(v map[string]*string) *GetSceneGroupInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSceneGroupInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetSceneGroupInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSceneGroupInfoRequest struct {
	// 酷应用编码
	CoolAppCode *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	// 开放群ID
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s GetSceneGroupInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSceneGroupInfoRequest) GoString() string {
	return s.String()
}

func (s *GetSceneGroupInfoRequest) SetCoolAppCode(v string) *GetSceneGroupInfoRequest {
	s.CoolAppCode = &v
	return s
}

func (s *GetSceneGroupInfoRequest) SetOpenConversationId(v string) *GetSceneGroupInfoRequest {
	s.OpenConversationId = &v
	return s
}

type GetSceneGroupInfoResponseBody struct {
	// 群url
	GroupUrl *string `json:"groupUrl,omitempty" xml:"groupUrl,omitempty"`
	// 群头像mediaId
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 开放群id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 群主员工id
	OwnerUserId *string `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	// result
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 场景群模板ID
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// 群名称
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetSceneGroupInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSceneGroupInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetSceneGroupInfoResponseBody) SetGroupUrl(v string) *GetSceneGroupInfoResponseBody {
	s.GroupUrl = &v
	return s
}

func (s *GetSceneGroupInfoResponseBody) SetIcon(v string) *GetSceneGroupInfoResponseBody {
	s.Icon = &v
	return s
}

func (s *GetSceneGroupInfoResponseBody) SetOpenConversationId(v string) *GetSceneGroupInfoResponseBody {
	s.OpenConversationId = &v
	return s
}

func (s *GetSceneGroupInfoResponseBody) SetOwnerUserId(v string) *GetSceneGroupInfoResponseBody {
	s.OwnerUserId = &v
	return s
}

func (s *GetSceneGroupInfoResponseBody) SetSuccess(v bool) *GetSceneGroupInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetSceneGroupInfoResponseBody) SetTemplateId(v string) *GetSceneGroupInfoResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetSceneGroupInfoResponseBody) SetTitle(v string) *GetSceneGroupInfoResponseBody {
	s.Title = &v
	return s
}

type GetSceneGroupInfoResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSceneGroupInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSceneGroupInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSceneGroupInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSceneGroupInfoResponse) SetHeaders(v map[string]*string) *GetSceneGroupInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSceneGroupInfoResponse) SetBody(v *GetSceneGroupInfoResponseBody) *GetSceneGroupInfoResponse {
	s.Body = v
	return s
}

type GetSceneGroupMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSceneGroupMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSceneGroupMembersHeaders) GoString() string {
	return s.String()
}

func (s *GetSceneGroupMembersHeaders) SetCommonHeaders(v map[string]*string) *GetSceneGroupMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSceneGroupMembersHeaders) SetXAcsDingtalkAccessToken(v string) *GetSceneGroupMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSceneGroupMembersRequest struct {
	// 酷应用编码
	CoolAppCode *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	// 分页游标，首页传0
	Cursor *string `json:"cursor,omitempty" xml:"cursor,omitempty"`
	// 开放群ID
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 本次查询返回数量上限（该入参传入值小于钉钉阈值时不启用）
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s GetSceneGroupMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSceneGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *GetSceneGroupMembersRequest) SetCoolAppCode(v string) *GetSceneGroupMembersRequest {
	s.CoolAppCode = &v
	return s
}

func (s *GetSceneGroupMembersRequest) SetCursor(v string) *GetSceneGroupMembersRequest {
	s.Cursor = &v
	return s
}

func (s *GetSceneGroupMembersRequest) SetOpenConversationId(v string) *GetSceneGroupMembersRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GetSceneGroupMembersRequest) SetSize(v int64) *GetSceneGroupMembersRequest {
	s.Size = &v
	return s
}

type GetSceneGroupMembersResponseBody struct {
	// 是否还有更多数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 群成员员工号
	MemberUserIds []*string `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty" type:"Repeated"`
	// 下一次请求的游标
	NextCursor *string `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	// result
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetSceneGroupMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSceneGroupMembersResponseBody) GoString() string {
	return s.String()
}

func (s *GetSceneGroupMembersResponseBody) SetHasMore(v bool) *GetSceneGroupMembersResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetSceneGroupMembersResponseBody) SetMemberUserIds(v []*string) *GetSceneGroupMembersResponseBody {
	s.MemberUserIds = v
	return s
}

func (s *GetSceneGroupMembersResponseBody) SetNextCursor(v string) *GetSceneGroupMembersResponseBody {
	s.NextCursor = &v
	return s
}

func (s *GetSceneGroupMembersResponseBody) SetSuccess(v bool) *GetSceneGroupMembersResponseBody {
	s.Success = &v
	return s
}

type GetSceneGroupMembersResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSceneGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSceneGroupMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSceneGroupMembersResponse) GoString() string {
	return s.String()
}

func (s *GetSceneGroupMembersResponse) SetHeaders(v map[string]*string) *GetSceneGroupMembersResponse {
	s.Headers = v
	return s
}

func (s *GetSceneGroupMembersResponse) SetBody(v *GetSceneGroupMembersResponseBody) *GetSceneGroupMembersResponse {
	s.Body = v
	return s
}

type InteractiveCardCreateInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InteractiveCardCreateInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s InteractiveCardCreateInstanceHeaders) GoString() string {
	return s.String()
}

func (s *InteractiveCardCreateInstanceHeaders) SetCommonHeaders(v map[string]*string) *InteractiveCardCreateInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InteractiveCardCreateInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *InteractiveCardCreateInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InteractiveCardCreateInstanceRequest struct {
	// 可控制卡片回调时的路由Key，用于指定特定的callbackUrl【可空：不填写默认用企业的回调地址】
	CallbackRouteKey *string                                       `json:"callbackRouteKey,omitempty" xml:"callbackRouteKey,omitempty"`
	CardData         *InteractiveCardCreateInstanceRequestCardData `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	// 卡片模板ID
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	// 【robotCode & chatBotId二选一必填】机器人ID（企业机器人）
	ChatBotId *string `json:"chatBotId,omitempty" xml:"chatBotId,omitempty"`
	// 发送的会话类型：单聊-0, 群聊-1（单聊时：openConversationId不用填写；receiverUserIdList填写有且一个员工号）
	ConversationType *int32 `json:"conversationType,omitempty" xml:"conversationType,omitempty"`
	// 接收卡片的群的openConversationId
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 唯一标识一张卡片的外部ID（卡片幂等ID，可用于更新或重复发送同一卡片到多个群会话）
	OutTrackId *string `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	// 指定用户可见的按钮列表（key：用户userId；value：用户数据）
	PrivateData map[string]*PrivateDataValue `json:"privateData,omitempty" xml:"privateData,omitempty"`
	// 接收人userId列表
	ReceiverUserIdList []*string `json:"receiverUserIdList,omitempty" xml:"receiverUserIdList,omitempty" type:"Repeated"`
	// 【robotCode & chatBotId二选一必填】机器人编码（群模板机器人）
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	// 用户ID类型：1：staffId模式【默认】；2：unionId模式；对应receiverUserIdList、privateData字段关于用户id的值填写方式
	UserIdType *int32 `json:"userIdType,omitempty" xml:"userIdType,omitempty"`
}

func (s InteractiveCardCreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s InteractiveCardCreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *InteractiveCardCreateInstanceRequest) SetCallbackRouteKey(v string) *InteractiveCardCreateInstanceRequest {
	s.CallbackRouteKey = &v
	return s
}

func (s *InteractiveCardCreateInstanceRequest) SetCardData(v *InteractiveCardCreateInstanceRequestCardData) *InteractiveCardCreateInstanceRequest {
	s.CardData = v
	return s
}

func (s *InteractiveCardCreateInstanceRequest) SetCardTemplateId(v string) *InteractiveCardCreateInstanceRequest {
	s.CardTemplateId = &v
	return s
}

func (s *InteractiveCardCreateInstanceRequest) SetChatBotId(v string) *InteractiveCardCreateInstanceRequest {
	s.ChatBotId = &v
	return s
}

func (s *InteractiveCardCreateInstanceRequest) SetConversationType(v int32) *InteractiveCardCreateInstanceRequest {
	s.ConversationType = &v
	return s
}

func (s *InteractiveCardCreateInstanceRequest) SetOpenConversationId(v string) *InteractiveCardCreateInstanceRequest {
	s.OpenConversationId = &v
	return s
}

func (s *InteractiveCardCreateInstanceRequest) SetOutTrackId(v string) *InteractiveCardCreateInstanceRequest {
	s.OutTrackId = &v
	return s
}

func (s *InteractiveCardCreateInstanceRequest) SetPrivateData(v map[string]*PrivateDataValue) *InteractiveCardCreateInstanceRequest {
	s.PrivateData = v
	return s
}

func (s *InteractiveCardCreateInstanceRequest) SetReceiverUserIdList(v []*string) *InteractiveCardCreateInstanceRequest {
	s.ReceiverUserIdList = v
	return s
}

func (s *InteractiveCardCreateInstanceRequest) SetRobotCode(v string) *InteractiveCardCreateInstanceRequest {
	s.RobotCode = &v
	return s
}

func (s *InteractiveCardCreateInstanceRequest) SetUserIdType(v int32) *InteractiveCardCreateInstanceRequest {
	s.UserIdType = &v
	return s
}

type InteractiveCardCreateInstanceRequestCardData struct {
	// 卡片模板内容替换参数-多媒体类型
	CardMediaIdParamMap map[string]*string `json:"cardMediaIdParamMap,omitempty" xml:"cardMediaIdParamMap,omitempty"`
	// 卡片模板内容替换参数-普通文本类型
	CardParamMap map[string]*string `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
}

func (s InteractiveCardCreateInstanceRequestCardData) String() string {
	return tea.Prettify(s)
}

func (s InteractiveCardCreateInstanceRequestCardData) GoString() string {
	return s.String()
}

func (s *InteractiveCardCreateInstanceRequestCardData) SetCardMediaIdParamMap(v map[string]*string) *InteractiveCardCreateInstanceRequestCardData {
	s.CardMediaIdParamMap = v
	return s
}

func (s *InteractiveCardCreateInstanceRequestCardData) SetCardParamMap(v map[string]*string) *InteractiveCardCreateInstanceRequestCardData {
	s.CardParamMap = v
	return s
}

type InteractiveCardCreateInstanceResponseBody struct {
	// 用于业务方后续查看已读列表的查询key
	ProcessQueryKey *string `json:"processQueryKey,omitempty" xml:"processQueryKey,omitempty"`
}

func (s InteractiveCardCreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InteractiveCardCreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *InteractiveCardCreateInstanceResponseBody) SetProcessQueryKey(v string) *InteractiveCardCreateInstanceResponseBody {
	s.ProcessQueryKey = &v
	return s
}

type InteractiveCardCreateInstanceResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InteractiveCardCreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InteractiveCardCreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s InteractiveCardCreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *InteractiveCardCreateInstanceResponse) SetHeaders(v map[string]*string) *InteractiveCardCreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *InteractiveCardCreateInstanceResponse) SetBody(v *InteractiveCardCreateInstanceResponseBody) *InteractiveCardCreateInstanceResponse {
	s.Body = v
	return s
}

type QueryGroupMuteStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGroupMuteStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMuteStatusHeaders) GoString() string {
	return s.String()
}

func (s *QueryGroupMuteStatusHeaders) SetCommonHeaders(v map[string]*string) *QueryGroupMuteStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGroupMuteStatusHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGroupMuteStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGroupMuteStatusRequest struct {
	// 开放的会话ID
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 群成员的员工工号
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryGroupMuteStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMuteStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryGroupMuteStatusRequest) SetOpenConversationId(v string) *QueryGroupMuteStatusRequest {
	s.OpenConversationId = &v
	return s
}

func (s *QueryGroupMuteStatusRequest) SetUserId(v string) *QueryGroupMuteStatusRequest {
	s.UserId = &v
	return s
}

type QueryGroupMuteStatusResponseBody struct {
	// 群禁言状态
	GroupMuteMode  *bool                                           `json:"groupMuteMode,omitempty" xml:"groupMuteMode,omitempty"`
	UserMuteResult *QueryGroupMuteStatusResponseBodyUserMuteResult `json:"userMuteResult,omitempty" xml:"userMuteResult,omitempty" type:"Struct"`
}

func (s QueryGroupMuteStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMuteStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupMuteStatusResponseBody) SetGroupMuteMode(v bool) *QueryGroupMuteStatusResponseBody {
	s.GroupMuteMode = &v
	return s
}

func (s *QueryGroupMuteStatusResponseBody) SetUserMuteResult(v *QueryGroupMuteStatusResponseBodyUserMuteResult) *QueryGroupMuteStatusResponseBody {
	s.UserMuteResult = v
	return s
}

type QueryGroupMuteStatusResponseBodyUserMuteResult struct {
	// 禁言结束时间
	MuteEndTime *int64 `json:"muteEndTime,omitempty" xml:"muteEndTime,omitempty"`
	// 禁言开始时间
	MuteStartTime *int64 `json:"muteStartTime,omitempty" xml:"muteStartTime,omitempty"`
	// 成员禁言状态
	UserMuteMode *bool `json:"userMuteMode,omitempty" xml:"userMuteMode,omitempty"`
}

func (s QueryGroupMuteStatusResponseBodyUserMuteResult) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMuteStatusResponseBodyUserMuteResult) GoString() string {
	return s.String()
}

func (s *QueryGroupMuteStatusResponseBodyUserMuteResult) SetMuteEndTime(v int64) *QueryGroupMuteStatusResponseBodyUserMuteResult {
	s.MuteEndTime = &v
	return s
}

func (s *QueryGroupMuteStatusResponseBodyUserMuteResult) SetMuteStartTime(v int64) *QueryGroupMuteStatusResponseBodyUserMuteResult {
	s.MuteStartTime = &v
	return s
}

func (s *QueryGroupMuteStatusResponseBodyUserMuteResult) SetUserMuteMode(v bool) *QueryGroupMuteStatusResponseBodyUserMuteResult {
	s.UserMuteMode = &v
	return s
}

type QueryGroupMuteStatusResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryGroupMuteStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryGroupMuteStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMuteStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupMuteStatusResponse) SetHeaders(v map[string]*string) *QueryGroupMuteStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupMuteStatusResponse) SetBody(v *QueryGroupMuteStatusResponseBody) *QueryGroupMuteStatusResponse {
	s.Body = v
	return s
}

type QueryMembersOfGroupRoleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMembersOfGroupRoleHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMembersOfGroupRoleHeaders) GoString() string {
	return s.String()
}

func (s *QueryMembersOfGroupRoleHeaders) SetCommonHeaders(v map[string]*string) *QueryMembersOfGroupRoleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMembersOfGroupRoleHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMembersOfGroupRoleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMembersOfGroupRoleRequest struct {
	// 开放群ID
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 开放群角色id
	OpenRoleId *string `json:"openRoleId,omitempty" xml:"openRoleId,omitempty"`
	// 时间戳
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s QueryMembersOfGroupRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMembersOfGroupRoleRequest) GoString() string {
	return s.String()
}

func (s *QueryMembersOfGroupRoleRequest) SetOpenConversationId(v string) *QueryMembersOfGroupRoleRequest {
	s.OpenConversationId = &v
	return s
}

func (s *QueryMembersOfGroupRoleRequest) SetOpenRoleId(v string) *QueryMembersOfGroupRoleRequest {
	s.OpenRoleId = &v
	return s
}

func (s *QueryMembersOfGroupRoleRequest) SetTimestamp(v int64) *QueryMembersOfGroupRoleRequest {
	s.Timestamp = &v
	return s
}

type QueryMembersOfGroupRoleResponseBody struct {
	// userIds
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s QueryMembersOfGroupRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMembersOfGroupRoleResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMembersOfGroupRoleResponseBody) SetUserIds(v []*string) *QueryMembersOfGroupRoleResponseBody {
	s.UserIds = v
	return s
}

type QueryMembersOfGroupRoleResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMembersOfGroupRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMembersOfGroupRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMembersOfGroupRoleResponse) GoString() string {
	return s.String()
}

func (s *QueryMembersOfGroupRoleResponse) SetHeaders(v map[string]*string) *QueryMembersOfGroupRoleResponse {
	s.Headers = v
	return s
}

func (s *QueryMembersOfGroupRoleResponse) SetBody(v *QueryMembersOfGroupRoleResponseBody) *QueryMembersOfGroupRoleResponse {
	s.Body = v
	return s
}

type SendInteractiveCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendInteractiveCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendInteractiveCardHeaders) GoString() string {
	return s.String()
}

func (s *SendInteractiveCardHeaders) SetCommonHeaders(v map[string]*string) *SendInteractiveCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendInteractiveCardHeaders) SetXAcsDingtalkAccessToken(v string) *SendInteractiveCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendInteractiveCardRequest struct {
	// 消息@人，{123456:"钉三多"}，key：根据userIdType来设置，【特殊设置：如果key、value都为"@ALL"则判断at所有人】
	AtOpenIds map[string]*string `json:"atOpenIds,omitempty" xml:"atOpenIds,omitempty"`
	// 可控制卡片回调时的路由Key，用于指定特定的callbackUrl【可空：不填写默认用企业的回调地址】
	CallbackRouteKey *string `json:"callbackRouteKey,omitempty" xml:"callbackRouteKey,omitempty"`
	// 卡片公共主体部分数据
	CardData *SendInteractiveCardRequestCardData `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	// 卡片属性
	CardOptions *SendInteractiveCardRequestCardOptions `json:"cardOptions,omitempty" xml:"cardOptions,omitempty" type:"Struct"`
	// 卡片模板ID
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	// 【robotCode & chatBotId二选一必填】机器人ID（企业机器人）
	ChatBotId *string `json:"chatBotId,omitempty" xml:"chatBotId,omitempty"`
	// 发送的会话类型：单聊-0, 群聊-1（单聊时：openConversationId不用填写；receiverUserIdList填写有且一个员工号）
	ConversationType *int32 `json:"conversationType,omitempty" xml:"conversationType,omitempty"`
	// 接收卡片的群的openConversationId
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 唯一标识一张卡片的外部ID（卡片幂等ID，可用于更新或重复发送同一卡片到多个群会话）
	OutTrackId *string `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	// 卡片用户私有差异部分数据（如卡片不同人显示不同按钮；key：用户userId；value：用户数据变量）
	PrivateData map[string]*PrivateDataValue `json:"privateData,omitempty" xml:"privateData,omitempty"`
	// 互动卡片消息需要群会话部分人可见时的接收人列表，不填写默认群会话所有人可见
	ReceiverUserIdList []*string `json:"receiverUserIdList,omitempty" xml:"receiverUserIdList,omitempty" type:"Repeated"`
	// 【robotCode & chatBotId二选一必填】机器人编码（群模板机器人）
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	// 用户ID类型：1：userId模式【默认】；2：unionId模式；对应receiverUserIdList、privateData字段关于用户id的值填写方式
	UserIdType *int32 `json:"userIdType,omitempty" xml:"userIdType,omitempty"`
}

func (s SendInteractiveCardRequest) String() string {
	return tea.Prettify(s)
}

func (s SendInteractiveCardRequest) GoString() string {
	return s.String()
}

func (s *SendInteractiveCardRequest) SetAtOpenIds(v map[string]*string) *SendInteractiveCardRequest {
	s.AtOpenIds = v
	return s
}

func (s *SendInteractiveCardRequest) SetCallbackRouteKey(v string) *SendInteractiveCardRequest {
	s.CallbackRouteKey = &v
	return s
}

func (s *SendInteractiveCardRequest) SetCardData(v *SendInteractiveCardRequestCardData) *SendInteractiveCardRequest {
	s.CardData = v
	return s
}

func (s *SendInteractiveCardRequest) SetCardOptions(v *SendInteractiveCardRequestCardOptions) *SendInteractiveCardRequest {
	s.CardOptions = v
	return s
}

func (s *SendInteractiveCardRequest) SetCardTemplateId(v string) *SendInteractiveCardRequest {
	s.CardTemplateId = &v
	return s
}

func (s *SendInteractiveCardRequest) SetChatBotId(v string) *SendInteractiveCardRequest {
	s.ChatBotId = &v
	return s
}

func (s *SendInteractiveCardRequest) SetConversationType(v int32) *SendInteractiveCardRequest {
	s.ConversationType = &v
	return s
}

func (s *SendInteractiveCardRequest) SetOpenConversationId(v string) *SendInteractiveCardRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SendInteractiveCardRequest) SetOutTrackId(v string) *SendInteractiveCardRequest {
	s.OutTrackId = &v
	return s
}

func (s *SendInteractiveCardRequest) SetPrivateData(v map[string]*PrivateDataValue) *SendInteractiveCardRequest {
	s.PrivateData = v
	return s
}

func (s *SendInteractiveCardRequest) SetReceiverUserIdList(v []*string) *SendInteractiveCardRequest {
	s.ReceiverUserIdList = v
	return s
}

func (s *SendInteractiveCardRequest) SetRobotCode(v string) *SendInteractiveCardRequest {
	s.RobotCode = &v
	return s
}

func (s *SendInteractiveCardRequest) SetUserIdType(v int32) *SendInteractiveCardRequest {
	s.UserIdType = &v
	return s
}

type SendInteractiveCardRequestCardData struct {
	// 卡片模板内容替换参数-多媒体类型
	CardMediaIdParamMap map[string]*string `json:"cardMediaIdParamMap,omitempty" xml:"cardMediaIdParamMap,omitempty"`
	// 卡片模板内容替换参数-普通文本类型
	CardParamMap map[string]*string `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
}

func (s SendInteractiveCardRequestCardData) String() string {
	return tea.Prettify(s)
}

func (s SendInteractiveCardRequestCardData) GoString() string {
	return s.String()
}

func (s *SendInteractiveCardRequestCardData) SetCardMediaIdParamMap(v map[string]*string) *SendInteractiveCardRequestCardData {
	s.CardMediaIdParamMap = v
	return s
}

func (s *SendInteractiveCardRequestCardData) SetCardParamMap(v map[string]*string) *SendInteractiveCardRequestCardData {
	s.CardParamMap = v
	return s
}

type SendInteractiveCardRequestCardOptions struct {
	// 是否支持转发
	SupportForward *bool `json:"supportForward,omitempty" xml:"supportForward,omitempty"`
}

func (s SendInteractiveCardRequestCardOptions) String() string {
	return tea.Prettify(s)
}

func (s SendInteractiveCardRequestCardOptions) GoString() string {
	return s.String()
}

func (s *SendInteractiveCardRequestCardOptions) SetSupportForward(v bool) *SendInteractiveCardRequestCardOptions {
	s.SupportForward = &v
	return s
}

type SendInteractiveCardResponseBody struct {
	// 创建卡片结果
	Result *SendInteractiveCardResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SendInteractiveCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendInteractiveCardResponseBody) GoString() string {
	return s.String()
}

func (s *SendInteractiveCardResponseBody) SetResult(v *SendInteractiveCardResponseBodyResult) *SendInteractiveCardResponseBody {
	s.Result = v
	return s
}

func (s *SendInteractiveCardResponseBody) SetSuccess(v bool) *SendInteractiveCardResponseBody {
	s.Success = &v
	return s
}

type SendInteractiveCardResponseBodyResult struct {
	// 用于业务方后续查看已读列表的查询key
	ProcessQueryKey *string `json:"processQueryKey,omitempty" xml:"processQueryKey,omitempty"`
}

func (s SendInteractiveCardResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SendInteractiveCardResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SendInteractiveCardResponseBodyResult) SetProcessQueryKey(v string) *SendInteractiveCardResponseBodyResult {
	s.ProcessQueryKey = &v
	return s
}

type SendInteractiveCardResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendInteractiveCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendInteractiveCardResponse) String() string {
	return tea.Prettify(s)
}

func (s SendInteractiveCardResponse) GoString() string {
	return s.String()
}

func (s *SendInteractiveCardResponse) SetHeaders(v map[string]*string) *SendInteractiveCardResponse {
	s.Headers = v
	return s
}

func (s *SendInteractiveCardResponse) SetBody(v *SendInteractiveCardResponseBody) *SendInteractiveCardResponse {
	s.Body = v
	return s
}

type SendRobotInteractiveCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendRobotInteractiveCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendRobotInteractiveCardHeaders) GoString() string {
	return s.String()
}

func (s *SendRobotInteractiveCardHeaders) SetCommonHeaders(v map[string]*string) *SendRobotInteractiveCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendRobotInteractiveCardHeaders) SetXAcsDingtalkAccessToken(v string) *SendRobotInteractiveCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendRobotInteractiveCardRequest struct {
	// 唯一标识一张卡片的外部ID（卡片幂等ID，可用于更新或重复发送同一卡片到多个群会话）【备注：同一个outTrackId重复创建，卡片数据不覆盖更新】
	CardBizId *string `json:"cardBizId,omitempty" xml:"cardBizId,omitempty"`
	// 卡片模板-文本内容参数（卡片json结构体）
	CardData *string `json:"cardData,omitempty" xml:"cardData,omitempty"`
	// 卡片搭建平台模板ID
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	// 【openConversationId & singleChatReceiver 二选一必填】接收卡片的加密群ID，特指多人群会话（非单聊）
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 机器人代码，群模板机器人网页有机器人ID；企业内部机器人为机器人appKey，企业三方机器人有robotCode
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	// 互动卡片发送选项
	SendOptions *SendRobotInteractiveCardRequestSendOptions `json:"sendOptions,omitempty" xml:"sendOptions,omitempty" type:"Struct"`
	// 【openConversationId & singleChatReceiver 二选一必填】单聊会话接受者json串
	SingleChatReceiver *string `json:"singleChatReceiver,omitempty" xml:"singleChatReceiver,omitempty"`
}

func (s SendRobotInteractiveCardRequest) String() string {
	return tea.Prettify(s)
}

func (s SendRobotInteractiveCardRequest) GoString() string {
	return s.String()
}

func (s *SendRobotInteractiveCardRequest) SetCardBizId(v string) *SendRobotInteractiveCardRequest {
	s.CardBizId = &v
	return s
}

func (s *SendRobotInteractiveCardRequest) SetCardData(v string) *SendRobotInteractiveCardRequest {
	s.CardData = &v
	return s
}

func (s *SendRobotInteractiveCardRequest) SetCardTemplateId(v string) *SendRobotInteractiveCardRequest {
	s.CardTemplateId = &v
	return s
}

func (s *SendRobotInteractiveCardRequest) SetOpenConversationId(v string) *SendRobotInteractiveCardRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SendRobotInteractiveCardRequest) SetRobotCode(v string) *SendRobotInteractiveCardRequest {
	s.RobotCode = &v
	return s
}

func (s *SendRobotInteractiveCardRequest) SetSendOptions(v *SendRobotInteractiveCardRequestSendOptions) *SendRobotInteractiveCardRequest {
	s.SendOptions = v
	return s
}

func (s *SendRobotInteractiveCardRequest) SetSingleChatReceiver(v string) *SendRobotInteractiveCardRequest {
	s.SingleChatReceiver = &v
	return s
}

type SendRobotInteractiveCardRequestSendOptions struct {
	// 是否@所有人
	AtAll *bool `json:"atAll,omitempty" xml:"atAll,omitempty"`
	// 消息@人，JSON格式：[{"nickName":"张三","userId":"userId0001"},{"nickName":"李四","unionId":"unionId001"}]
	AtUserListJson *string `json:"atUserListJson,omitempty" xml:"atUserListJson,omitempty"`
	// 卡片特殊属性json串
	CardPropertyJson *string `json:"cardPropertyJson,omitempty" xml:"cardPropertyJson,omitempty"`
	// 消息仅部分人可见的接收人列表【可空：为空则群所有人可见】，JSON格式：[{"userId":"userId0001"},{"unionId":"unionId001"}]
	ReceiverListJson *string `json:"receiverListJson,omitempty" xml:"receiverListJson,omitempty"`
}

func (s SendRobotInteractiveCardRequestSendOptions) String() string {
	return tea.Prettify(s)
}

func (s SendRobotInteractiveCardRequestSendOptions) GoString() string {
	return s.String()
}

func (s *SendRobotInteractiveCardRequestSendOptions) SetAtAll(v bool) *SendRobotInteractiveCardRequestSendOptions {
	s.AtAll = &v
	return s
}

func (s *SendRobotInteractiveCardRequestSendOptions) SetAtUserListJson(v string) *SendRobotInteractiveCardRequestSendOptions {
	s.AtUserListJson = &v
	return s
}

func (s *SendRobotInteractiveCardRequestSendOptions) SetCardPropertyJson(v string) *SendRobotInteractiveCardRequestSendOptions {
	s.CardPropertyJson = &v
	return s
}

func (s *SendRobotInteractiveCardRequestSendOptions) SetReceiverListJson(v string) *SendRobotInteractiveCardRequestSendOptions {
	s.ReceiverListJson = &v
	return s
}

type SendRobotInteractiveCardResponseBody struct {
	// 用于业务方后续查看已读列表的查询key
	ProcessQueryKey *string `json:"processQueryKey,omitempty" xml:"processQueryKey,omitempty"`
}

func (s SendRobotInteractiveCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendRobotInteractiveCardResponseBody) GoString() string {
	return s.String()
}

func (s *SendRobotInteractiveCardResponseBody) SetProcessQueryKey(v string) *SendRobotInteractiveCardResponseBody {
	s.ProcessQueryKey = &v
	return s
}

type SendRobotInteractiveCardResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendRobotInteractiveCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendRobotInteractiveCardResponse) String() string {
	return tea.Prettify(s)
}

func (s SendRobotInteractiveCardResponse) GoString() string {
	return s.String()
}

func (s *SendRobotInteractiveCardResponse) SetHeaders(v map[string]*string) *SendRobotInteractiveCardResponse {
	s.Headers = v
	return s
}

func (s *SendRobotInteractiveCardResponse) SetBody(v *SendRobotInteractiveCardResponseBody) *SendRobotInteractiveCardResponse {
	s.Body = v
	return s
}

type SendTemplateInteractiveCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendTemplateInteractiveCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendTemplateInteractiveCardHeaders) GoString() string {
	return s.String()
}

func (s *SendTemplateInteractiveCardHeaders) SetCommonHeaders(v map[string]*string) *SendTemplateInteractiveCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendTemplateInteractiveCardHeaders) SetXAcsDingtalkAccessToken(v string) *SendTemplateInteractiveCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendTemplateInteractiveCardRequest struct {
	// 可控制卡片回调的url【可空：不填写无需回调】
	CallbackUrl *string `json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
	// 卡片模板-文本内容参数（卡片json结构体）
	CardData *string `json:"cardData,omitempty" xml:"cardData,omitempty"`
	// 卡片内容模板ID，响应模板目前有：TuWenCard01、TuWenCard02、TuWenCard03、TuWenCard04 4种
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	// 【openConversationId & singleChatReceiver 二选一必填】接收卡片的加密群ID，特指多人群会话（非单聊）
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 唯一标识一张卡片的外部ID（卡片幂等ID，可用于更新或重复发送同一卡片到多个群会话）【备注：同一个outTrackId重复创建，卡片数据不覆盖更新】
	OutTrackId *string `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	// 机器人代码，群模板机器人网页有机器人ID；企业内部机器人为机器人appKey，企业三方机器人有robotCode
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	// 互动卡片发送选项
	SendOptions *SendTemplateInteractiveCardRequestSendOptions `json:"sendOptions,omitempty" xml:"sendOptions,omitempty" type:"Struct"`
	// 【openConversationId & singleChatReceiver 二选一必填】单聊会话接受者json串
	SingleChatReceiver *string `json:"singleChatReceiver,omitempty" xml:"singleChatReceiver,omitempty"`
}

func (s SendTemplateInteractiveCardRequest) String() string {
	return tea.Prettify(s)
}

func (s SendTemplateInteractiveCardRequest) GoString() string {
	return s.String()
}

func (s *SendTemplateInteractiveCardRequest) SetCallbackUrl(v string) *SendTemplateInteractiveCardRequest {
	s.CallbackUrl = &v
	return s
}

func (s *SendTemplateInteractiveCardRequest) SetCardData(v string) *SendTemplateInteractiveCardRequest {
	s.CardData = &v
	return s
}

func (s *SendTemplateInteractiveCardRequest) SetCardTemplateId(v string) *SendTemplateInteractiveCardRequest {
	s.CardTemplateId = &v
	return s
}

func (s *SendTemplateInteractiveCardRequest) SetOpenConversationId(v string) *SendTemplateInteractiveCardRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SendTemplateInteractiveCardRequest) SetOutTrackId(v string) *SendTemplateInteractiveCardRequest {
	s.OutTrackId = &v
	return s
}

func (s *SendTemplateInteractiveCardRequest) SetRobotCode(v string) *SendTemplateInteractiveCardRequest {
	s.RobotCode = &v
	return s
}

func (s *SendTemplateInteractiveCardRequest) SetSendOptions(v *SendTemplateInteractiveCardRequestSendOptions) *SendTemplateInteractiveCardRequest {
	s.SendOptions = v
	return s
}

func (s *SendTemplateInteractiveCardRequest) SetSingleChatReceiver(v string) *SendTemplateInteractiveCardRequest {
	s.SingleChatReceiver = &v
	return s
}

type SendTemplateInteractiveCardRequestSendOptions struct {
	// 是否@所有人
	AtAll *bool `json:"atAll,omitempty" xml:"atAll,omitempty"`
	// 消息@人，JSON格式：[{"nickName":"张三","userId":"userId0001"},{"nickName":"李四","unionId":"unionId001"}]
	AtUserListJson *string `json:"atUserListJson,omitempty" xml:"atUserListJson,omitempty"`
	// 卡片特殊属性json串
	CardPropertyJson *string `json:"cardPropertyJson,omitempty" xml:"cardPropertyJson,omitempty"`
	// 消息仅部分人可见的接收人列表【可空：为空则群所有人可见】，JSON格式：[{"userId":"userId0001"},{"unionId":"unionId001"}]
	ReceiverListJson *string `json:"receiverListJson,omitempty" xml:"receiverListJson,omitempty"`
}

func (s SendTemplateInteractiveCardRequestSendOptions) String() string {
	return tea.Prettify(s)
}

func (s SendTemplateInteractiveCardRequestSendOptions) GoString() string {
	return s.String()
}

func (s *SendTemplateInteractiveCardRequestSendOptions) SetAtAll(v bool) *SendTemplateInteractiveCardRequestSendOptions {
	s.AtAll = &v
	return s
}

func (s *SendTemplateInteractiveCardRequestSendOptions) SetAtUserListJson(v string) *SendTemplateInteractiveCardRequestSendOptions {
	s.AtUserListJson = &v
	return s
}

func (s *SendTemplateInteractiveCardRequestSendOptions) SetCardPropertyJson(v string) *SendTemplateInteractiveCardRequestSendOptions {
	s.CardPropertyJson = &v
	return s
}

func (s *SendTemplateInteractiveCardRequestSendOptions) SetReceiverListJson(v string) *SendTemplateInteractiveCardRequestSendOptions {
	s.ReceiverListJson = &v
	return s
}

type SendTemplateInteractiveCardResponseBody struct {
	// 用于业务方后续查看已读列表的查询key
	ProcessQueryKey *string `json:"processQueryKey,omitempty" xml:"processQueryKey,omitempty"`
}

func (s SendTemplateInteractiveCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendTemplateInteractiveCardResponseBody) GoString() string {
	return s.String()
}

func (s *SendTemplateInteractiveCardResponseBody) SetProcessQueryKey(v string) *SendTemplateInteractiveCardResponseBody {
	s.ProcessQueryKey = &v
	return s
}

type SendTemplateInteractiveCardResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendTemplateInteractiveCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendTemplateInteractiveCardResponse) String() string {
	return tea.Prettify(s)
}

func (s SendTemplateInteractiveCardResponse) GoString() string {
	return s.String()
}

func (s *SendTemplateInteractiveCardResponse) SetHeaders(v map[string]*string) *SendTemplateInteractiveCardResponse {
	s.Headers = v
	return s
}

func (s *SendTemplateInteractiveCardResponse) SetBody(v *SendTemplateInteractiveCardResponseBody) *SendTemplateInteractiveCardResponse {
	s.Body = v
	return s
}

type TopboxCloseHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TopboxCloseHeaders) String() string {
	return tea.Prettify(s)
}

func (s TopboxCloseHeaders) GoString() string {
	return s.String()
}

func (s *TopboxCloseHeaders) SetCommonHeaders(v map[string]*string) *TopboxCloseHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TopboxCloseHeaders) SetXAcsDingtalkAccessToken(v string) *TopboxCloseHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TopboxCloseRequest struct {
	// 酷应用编码
	CoolAppCode *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	// 接收卡片的群的openConversationId
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 唯一标识一张卡片的外部ID（卡片幂等ID，可用于更新或重复发送同一卡片到多个群会话）
	OutTrackId *string `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
}

func (s TopboxCloseRequest) String() string {
	return tea.Prettify(s)
}

func (s TopboxCloseRequest) GoString() string {
	return s.String()
}

func (s *TopboxCloseRequest) SetCoolAppCode(v string) *TopboxCloseRequest {
	s.CoolAppCode = &v
	return s
}

func (s *TopboxCloseRequest) SetOpenConversationId(v string) *TopboxCloseRequest {
	s.OpenConversationId = &v
	return s
}

func (s *TopboxCloseRequest) SetOutTrackId(v string) *TopboxCloseRequest {
	s.OutTrackId = &v
	return s
}

type TopboxCloseResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s TopboxCloseResponse) String() string {
	return tea.Prettify(s)
}

func (s TopboxCloseResponse) GoString() string {
	return s.String()
}

func (s *TopboxCloseResponse) SetHeaders(v map[string]*string) *TopboxCloseResponse {
	s.Headers = v
	return s
}

type TopboxOpenHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TopboxOpenHeaders) String() string {
	return tea.Prettify(s)
}

func (s TopboxOpenHeaders) GoString() string {
	return s.String()
}

func (s *TopboxOpenHeaders) SetCommonHeaders(v map[string]*string) *TopboxOpenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TopboxOpenHeaders) SetXAcsDingtalkAccessToken(v string) *TopboxOpenHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TopboxOpenRequest struct {
	// 酷应用编码
	CoolAppCode *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	// 吊顶的过期时间（绝对时间）
	ExpiredTime *int64 `json:"expiredTime,omitempty" xml:"expiredTime,omitempty"`
	// 接收卡片的群的openConversationId
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 唯一标识一张卡片的外部ID（卡片幂等ID，可用于更新或重复发送同一卡片到多个群会话）
	OutTrackId *string `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	// 期望吊顶的端（多个'|'隔开，如："ios|win|"）
	Platforms *string `json:"platforms,omitempty" xml:"platforms,omitempty"`
}

func (s TopboxOpenRequest) String() string {
	return tea.Prettify(s)
}

func (s TopboxOpenRequest) GoString() string {
	return s.String()
}

func (s *TopboxOpenRequest) SetCoolAppCode(v string) *TopboxOpenRequest {
	s.CoolAppCode = &v
	return s
}

func (s *TopboxOpenRequest) SetExpiredTime(v int64) *TopboxOpenRequest {
	s.ExpiredTime = &v
	return s
}

func (s *TopboxOpenRequest) SetOpenConversationId(v string) *TopboxOpenRequest {
	s.OpenConversationId = &v
	return s
}

func (s *TopboxOpenRequest) SetOutTrackId(v string) *TopboxOpenRequest {
	s.OutTrackId = &v
	return s
}

func (s *TopboxOpenRequest) SetPlatforms(v string) *TopboxOpenRequest {
	s.Platforms = &v
	return s
}

type TopboxOpenResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s TopboxOpenResponse) String() string {
	return tea.Prettify(s)
}

func (s TopboxOpenResponse) GoString() string {
	return s.String()
}

func (s *TopboxOpenResponse) SetHeaders(v map[string]*string) *TopboxOpenResponse {
	s.Headers = v
	return s
}

type UpdateGroupPermissionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateGroupPermissionHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupPermissionHeaders) GoString() string {
	return s.String()
}

func (s *UpdateGroupPermissionHeaders) SetCommonHeaders(v map[string]*string) *UpdateGroupPermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateGroupPermissionHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateGroupPermissionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateGroupPermissionRequest struct {
	// 开放群ID
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 群权限项
	PermissionGroup *string `json:"permissionGroup,omitempty" xml:"permissionGroup,omitempty"`
	// 状态,0-关闭，1-开启
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateGroupPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupPermissionRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupPermissionRequest) SetOpenConversationId(v string) *UpdateGroupPermissionRequest {
	s.OpenConversationId = &v
	return s
}

func (s *UpdateGroupPermissionRequest) SetPermissionGroup(v string) *UpdateGroupPermissionRequest {
	s.PermissionGroup = &v
	return s
}

func (s *UpdateGroupPermissionRequest) SetStatus(v string) *UpdateGroupPermissionRequest {
	s.Status = &v
	return s
}

type UpdateGroupPermissionResponseBody struct {
	// result
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateGroupPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupPermissionResponseBody) SetSuccess(v bool) *UpdateGroupPermissionResponseBody {
	s.Success = &v
	return s
}

type UpdateGroupPermissionResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateGroupPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGroupPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupPermissionResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupPermissionResponse) SetHeaders(v map[string]*string) *UpdateGroupPermissionResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupPermissionResponse) SetBody(v *UpdateGroupPermissionResponseBody) *UpdateGroupPermissionResponse {
	s.Body = v
	return s
}

type UpdateGroupSubAdminHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateGroupSubAdminHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupSubAdminHeaders) GoString() string {
	return s.String()
}

func (s *UpdateGroupSubAdminHeaders) SetCommonHeaders(v map[string]*string) *UpdateGroupSubAdminHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateGroupSubAdminHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateGroupSubAdminHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateGroupSubAdminRequest struct {
	// 开放群ID
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 2-群管理员 3-普通群成员
	Role *int64 `json:"role,omitempty" xml:"role,omitempty"`
	// 用户ID清单
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s UpdateGroupSubAdminRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupSubAdminRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupSubAdminRequest) SetOpenConversationId(v string) *UpdateGroupSubAdminRequest {
	s.OpenConversationId = &v
	return s
}

func (s *UpdateGroupSubAdminRequest) SetRole(v int64) *UpdateGroupSubAdminRequest {
	s.Role = &v
	return s
}

func (s *UpdateGroupSubAdminRequest) SetUserIds(v []*string) *UpdateGroupSubAdminRequest {
	s.UserIds = v
	return s
}

type UpdateGroupSubAdminResponseBody struct {
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateGroupSubAdminResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupSubAdminResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupSubAdminResponseBody) SetSuccess(v bool) *UpdateGroupSubAdminResponseBody {
	s.Success = &v
	return s
}

type UpdateGroupSubAdminResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateGroupSubAdminResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGroupSubAdminResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupSubAdminResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupSubAdminResponse) SetHeaders(v map[string]*string) *UpdateGroupSubAdminResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupSubAdminResponse) SetBody(v *UpdateGroupSubAdminResponseBody) *UpdateGroupSubAdminResponse {
	s.Body = v
	return s
}

type UpdateInteractiveCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateInteractiveCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateInteractiveCardHeaders) GoString() string {
	return s.String()
}

func (s *UpdateInteractiveCardHeaders) SetCommonHeaders(v map[string]*string) *UpdateInteractiveCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateInteractiveCardHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateInteractiveCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateInteractiveCardRequest struct {
	// 卡片公共主体部分数据
	CardData *UpdateInteractiveCardRequestCardData `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	// 发送可交互卡片的一些功能选项
	CardOptions *UpdateInteractiveCardRequestCardOptions `json:"cardOptions,omitempty" xml:"cardOptions,omitempty" type:"Struct"`
	// 唯一标识一张卡片的外部ID
	OutTrackId *string `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	// 卡片用户私有差异部分数据（如卡片不同人显示不同按钮；key：用户userId；value：用户数据变量）
	PrivateData map[string]*PrivateDataValue `json:"privateData,omitempty" xml:"privateData,omitempty"`
	// 用户ID类型：1：userId模式【默认】；2：unionId模式；对应receiverUserIdList、privateData字段关于用户id的值填写方式
	UserIdType *int32 `json:"userIdType,omitempty" xml:"userIdType,omitempty"`
}

func (s UpdateInteractiveCardRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInteractiveCardRequest) GoString() string {
	return s.String()
}

func (s *UpdateInteractiveCardRequest) SetCardData(v *UpdateInteractiveCardRequestCardData) *UpdateInteractiveCardRequest {
	s.CardData = v
	return s
}

func (s *UpdateInteractiveCardRequest) SetCardOptions(v *UpdateInteractiveCardRequestCardOptions) *UpdateInteractiveCardRequest {
	s.CardOptions = v
	return s
}

func (s *UpdateInteractiveCardRequest) SetOutTrackId(v string) *UpdateInteractiveCardRequest {
	s.OutTrackId = &v
	return s
}

func (s *UpdateInteractiveCardRequest) SetPrivateData(v map[string]*PrivateDataValue) *UpdateInteractiveCardRequest {
	s.PrivateData = v
	return s
}

func (s *UpdateInteractiveCardRequest) SetUserIdType(v int32) *UpdateInteractiveCardRequest {
	s.UserIdType = &v
	return s
}

type UpdateInteractiveCardRequestCardData struct {
	// 卡片模板内容替换参数-多媒体类型
	CardMediaIdParamMap map[string]*string `json:"cardMediaIdParamMap,omitempty" xml:"cardMediaIdParamMap,omitempty"`
	// 卡片模板内容替换参数-普通文本类型
	CardParamMap map[string]*string `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
}

func (s UpdateInteractiveCardRequestCardData) String() string {
	return tea.Prettify(s)
}

func (s UpdateInteractiveCardRequestCardData) GoString() string {
	return s.String()
}

func (s *UpdateInteractiveCardRequestCardData) SetCardMediaIdParamMap(v map[string]*string) *UpdateInteractiveCardRequestCardData {
	s.CardMediaIdParamMap = v
	return s
}

func (s *UpdateInteractiveCardRequestCardData) SetCardParamMap(v map[string]*string) *UpdateInteractiveCardRequestCardData {
	s.CardParamMap = v
	return s
}

type UpdateInteractiveCardRequestCardOptions struct {
	// 按key更新cardData数据(不填默认覆盖更新)
	UpdateCardDataByKey *bool `json:"updateCardDataByKey,omitempty" xml:"updateCardDataByKey,omitempty"`
	// 按key更新privateData用户数据(不填默认覆盖更新)
	UpdatePrivateDataByKey *bool `json:"updatePrivateDataByKey,omitempty" xml:"updatePrivateDataByKey,omitempty"`
}

func (s UpdateInteractiveCardRequestCardOptions) String() string {
	return tea.Prettify(s)
}

func (s UpdateInteractiveCardRequestCardOptions) GoString() string {
	return s.String()
}

func (s *UpdateInteractiveCardRequestCardOptions) SetUpdateCardDataByKey(v bool) *UpdateInteractiveCardRequestCardOptions {
	s.UpdateCardDataByKey = &v
	return s
}

func (s *UpdateInteractiveCardRequestCardOptions) SetUpdatePrivateDataByKey(v bool) *UpdateInteractiveCardRequestCardOptions {
	s.UpdatePrivateDataByKey = &v
	return s
}

type UpdateInteractiveCardResponseBody struct {
	// result
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateInteractiveCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInteractiveCardResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInteractiveCardResponseBody) SetSuccess(v string) *UpdateInteractiveCardResponseBody {
	s.Success = &v
	return s
}

type UpdateInteractiveCardResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInteractiveCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInteractiveCardResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInteractiveCardResponse) GoString() string {
	return s.String()
}

func (s *UpdateInteractiveCardResponse) SetHeaders(v map[string]*string) *UpdateInteractiveCardResponse {
	s.Headers = v
	return s
}

func (s *UpdateInteractiveCardResponse) SetBody(v *UpdateInteractiveCardResponseBody) *UpdateInteractiveCardResponse {
	s.Body = v
	return s
}

type UpdateMemberBanWordsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateMemberBanWordsHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemberBanWordsHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMemberBanWordsHeaders) SetCommonHeaders(v map[string]*string) *UpdateMemberBanWordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMemberBanWordsHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateMemberBanWordsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateMemberBanWordsRequest struct {
	// 禁言持续时长（单位：毫秒）
	MuteDuration *int64 `json:"muteDuration,omitempty" xml:"muteDuration,omitempty"`
	// 禁言状态(0表示取消禁言，1表示禁言)
	MuteStatus *int32 `json:"muteStatus,omitempty" xml:"muteStatus,omitempty"`
	// 开放群id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 需要禁言或取消禁言的群成员列表
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s UpdateMemberBanWordsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemberBanWordsRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemberBanWordsRequest) SetMuteDuration(v int64) *UpdateMemberBanWordsRequest {
	s.MuteDuration = &v
	return s
}

func (s *UpdateMemberBanWordsRequest) SetMuteStatus(v int32) *UpdateMemberBanWordsRequest {
	s.MuteStatus = &v
	return s
}

func (s *UpdateMemberBanWordsRequest) SetOpenConversationId(v string) *UpdateMemberBanWordsRequest {
	s.OpenConversationId = &v
	return s
}

func (s *UpdateMemberBanWordsRequest) SetUserIdList(v []*string) *UpdateMemberBanWordsRequest {
	s.UserIdList = v
	return s
}

type UpdateMemberBanWordsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpdateMemberBanWordsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemberBanWordsResponse) GoString() string {
	return s.String()
}

func (s *UpdateMemberBanWordsResponse) SetHeaders(v map[string]*string) *UpdateMemberBanWordsResponse {
	s.Headers = v
	return s
}

type UpdateMemberGroupNickHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateMemberGroupNickHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemberGroupNickHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMemberGroupNickHeaders) SetCommonHeaders(v map[string]*string) *UpdateMemberGroupNickHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMemberGroupNickHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateMemberGroupNickHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateMemberGroupNickRequest struct {
	// 群昵称
	GroupNick *string `json:"groupNick,omitempty" xml:"groupNick,omitempty"`
	// 开放群ID
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 用户ID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateMemberGroupNickRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemberGroupNickRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemberGroupNickRequest) SetGroupNick(v string) *UpdateMemberGroupNickRequest {
	s.GroupNick = &v
	return s
}

func (s *UpdateMemberGroupNickRequest) SetOpenConversationId(v string) *UpdateMemberGroupNickRequest {
	s.OpenConversationId = &v
	return s
}

func (s *UpdateMemberGroupNickRequest) SetUserId(v string) *UpdateMemberGroupNickRequest {
	s.UserId = &v
	return s
}

type UpdateMemberGroupNickResponseBody struct {
	// result
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateMemberGroupNickResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemberGroupNickResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMemberGroupNickResponseBody) SetSuccess(v bool) *UpdateMemberGroupNickResponseBody {
	s.Success = &v
	return s
}

type UpdateMemberGroupNickResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateMemberGroupNickResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMemberGroupNickResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemberGroupNickResponse) GoString() string {
	return s.String()
}

func (s *UpdateMemberGroupNickResponse) SetHeaders(v map[string]*string) *UpdateMemberGroupNickResponse {
	s.Headers = v
	return s
}

func (s *UpdateMemberGroupNickResponse) SetBody(v *UpdateMemberGroupNickResponseBody) *UpdateMemberGroupNickResponse {
	s.Body = v
	return s
}

type UpdateTheGroupRolesOfGroupMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateTheGroupRolesOfGroupMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTheGroupRolesOfGroupMemberHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTheGroupRolesOfGroupMemberHeaders) SetCommonHeaders(v map[string]*string) *UpdateTheGroupRolesOfGroupMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTheGroupRolesOfGroupMemberHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateTheGroupRolesOfGroupMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateTheGroupRolesOfGroupMemberRequest struct {
	// 开放群ID
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 群角色列表
	OpenRoleIds []*string `json:"openRoleIds,omitempty" xml:"openRoleIds,omitempty" type:"Repeated"`
	// 用户ID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateTheGroupRolesOfGroupMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTheGroupRolesOfGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdateTheGroupRolesOfGroupMemberRequest) SetOpenConversationId(v string) *UpdateTheGroupRolesOfGroupMemberRequest {
	s.OpenConversationId = &v
	return s
}

func (s *UpdateTheGroupRolesOfGroupMemberRequest) SetOpenRoleIds(v []*string) *UpdateTheGroupRolesOfGroupMemberRequest {
	s.OpenRoleIds = v
	return s
}

func (s *UpdateTheGroupRolesOfGroupMemberRequest) SetUserId(v string) *UpdateTheGroupRolesOfGroupMemberRequest {
	s.UserId = &v
	return s
}

type UpdateTheGroupRolesOfGroupMemberResponseBody struct {
	// result
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateTheGroupRolesOfGroupMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTheGroupRolesOfGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTheGroupRolesOfGroupMemberResponseBody) SetSuccess(v bool) *UpdateTheGroupRolesOfGroupMemberResponseBody {
	s.Success = &v
	return s
}

type UpdateTheGroupRolesOfGroupMemberResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTheGroupRolesOfGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTheGroupRolesOfGroupMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTheGroupRolesOfGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *UpdateTheGroupRolesOfGroupMemberResponse) SetHeaders(v map[string]*string) *UpdateTheGroupRolesOfGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *UpdateTheGroupRolesOfGroupMemberResponse) SetBody(v *UpdateTheGroupRolesOfGroupMemberResponseBody) *UpdateTheGroupRolesOfGroupMemberResponse {
	s.Body = v
	return s
}

type PrivateDataValue struct {
	// 卡片模板内容替换参数-普通文本类型
	CardParamMap map[string]*string `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
	// 卡片模板内容替换参数-多媒体类型
	CardMediaIdParamMap map[string]*string `json:"cardMediaIdParamMap,omitempty" xml:"cardMediaIdParamMap,omitempty"`
}

func (s PrivateDataValue) String() string {
	return tea.Prettify(s)
}

func (s PrivateDataValue) GoString() string {
	return s.String()
}

func (s *PrivateDataValue) SetCardParamMap(v map[string]*string) *PrivateDataValue {
	s.CardParamMap = v
	return s
}

func (s *PrivateDataValue) SetCardMediaIdParamMap(v map[string]*string) *PrivateDataValue {
	s.CardMediaIdParamMap = v
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

func (client *Client) BatchQueryGroupMember(request *BatchQueryGroupMemberRequest) (_result *BatchQueryGroupMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchQueryGroupMemberHeaders{}
	_result = &BatchQueryGroupMemberResponse{}
	_body, _err := client.BatchQueryGroupMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchQueryGroupMemberWithOptions(request *BatchQueryGroupMemberRequest, headers *BatchQueryGroupMemberHeaders, runtime *util.RuntimeOptions) (_result *BatchQueryGroupMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CoolAppCode)) {
		body["coolAppCode"] = request.CoolAppCode
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
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
	_result = &BatchQueryGroupMemberResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchQueryGroupMember"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/sceneGroups/members/batchQuery"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CardTemplateBuildAction(request *CardTemplateBuildActionRequest) (_result *CardTemplateBuildActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CardTemplateBuildActionHeaders{}
	_result = &CardTemplateBuildActionResponse{}
	_body, _err := client.CardTemplateBuildActionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CardTemplateBuildActionWithOptions(request *CardTemplateBuildActionRequest, headers *CardTemplateBuildActionHeaders, runtime *util.RuntimeOptions) (_result *CardTemplateBuildActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		body["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.CardTemplateJson)) {
		body["cardTemplateJson"] = request.CardTemplateJson
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
	_result = &CardTemplateBuildActionResponse{}
	_body, _err := client.DoROARequest(tea.String("CardTemplateBuildAction"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/interactiveCards/templates/buildAction"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChatIdToOpenConversationId(chatId *string) (_result *ChatIdToOpenConversationIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ChatIdToOpenConversationIdHeaders{}
	_result = &ChatIdToOpenConversationIdResponse{}
	_body, _err := client.ChatIdToOpenConversationIdWithOptions(chatId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChatIdToOpenConversationIdWithOptions(chatId *string, headers *ChatIdToOpenConversationIdHeaders, runtime *util.RuntimeOptions) (_result *ChatIdToOpenConversationIdResponse, _err error) {
	chatId = openapiutil.GetEncodeParam(chatId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &ChatIdToOpenConversationIdResponse{}
	_body, _err := client.DoROARequest(tea.String("ChatIdToOpenConversationId"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/chat/"+tea.StringValue(chatId)+"/convertToOpenConversationId"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInterconnectionUrl(request *GetInterconnectionUrlRequest) (_result *GetInterconnectionUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetInterconnectionUrlHeaders{}
	_result = &GetInterconnectionUrlResponse{}
	_body, _err := client.GetInterconnectionUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInterconnectionUrlWithOptions(request *GetInterconnectionUrlRequest, headers *GetInterconnectionUrlHeaders, runtime *util.RuntimeOptions) (_result *GetInterconnectionUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUserAvatar)) {
		body["appUserAvatar"] = request.AppUserAvatar
	}

	if !tea.BoolValue(util.IsUnset(request.AppUserAvatarType)) {
		body["appUserAvatarType"] = request.AppUserAvatarType
	}

	if !tea.BoolValue(util.IsUnset(request.AppUserId)) {
		body["appUserId"] = request.AppUserId
	}

	if !tea.BoolValue(util.IsUnset(request.AppUserMobileNumber)) {
		body["appUserMobileNumber"] = request.AppUserMobileNumber
	}

	if !tea.BoolValue(util.IsUnset(request.AppUserName)) {
		body["appUserName"] = request.AppUserName
	}

	if !tea.BoolValue(util.IsUnset(request.MsgPageType)) {
		body["msgPageType"] = request.MsgPageType
	}

	if !tea.BoolValue(util.IsUnset(request.QrCode)) {
		body["qrCode"] = request.QrCode
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		body["signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCode)) {
		body["sourceCode"] = request.SourceCode
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		body["sourceType"] = request.SourceType
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
	_result = &GetInterconnectionUrlResponse{}
	_body, _err := client.DoROARequest(tea.String("GetInterconnectionUrl"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/interconnections/sessions/urls"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSceneGroupInfo(request *GetSceneGroupInfoRequest) (_result *GetSceneGroupInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSceneGroupInfoHeaders{}
	_result = &GetSceneGroupInfoResponse{}
	_body, _err := client.GetSceneGroupInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSceneGroupInfoWithOptions(request *GetSceneGroupInfoRequest, headers *GetSceneGroupInfoHeaders, runtime *util.RuntimeOptions) (_result *GetSceneGroupInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CoolAppCode)) {
		body["coolAppCode"] = request.CoolAppCode
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
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
	_result = &GetSceneGroupInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSceneGroupInfo"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/sceneGroups/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSceneGroupMembers(request *GetSceneGroupMembersRequest) (_result *GetSceneGroupMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSceneGroupMembersHeaders{}
	_result = &GetSceneGroupMembersResponse{}
	_body, _err := client.GetSceneGroupMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSceneGroupMembersWithOptions(request *GetSceneGroupMembersRequest, headers *GetSceneGroupMembersHeaders, runtime *util.RuntimeOptions) (_result *GetSceneGroupMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CoolAppCode)) {
		body["coolAppCode"] = request.CoolAppCode
	}

	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		body["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["size"] = request.Size
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
	_result = &GetSceneGroupMembersResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSceneGroupMembers"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/sceneGroups/members/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InteractiveCardCreateInstance(request *InteractiveCardCreateInstanceRequest) (_result *InteractiveCardCreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InteractiveCardCreateInstanceHeaders{}
	_result = &InteractiveCardCreateInstanceResponse{}
	_body, _err := client.InteractiveCardCreateInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InteractiveCardCreateInstanceWithOptions(request *InteractiveCardCreateInstanceRequest, headers *InteractiveCardCreateInstanceHeaders, runtime *util.RuntimeOptions) (_result *InteractiveCardCreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallbackRouteKey)) {
		body["callbackRouteKey"] = request.CallbackRouteKey
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CardData))) {
		body["cardData"] = request.CardData
	}

	if !tea.BoolValue(util.IsUnset(request.CardTemplateId)) {
		body["cardTemplateId"] = request.CardTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.ChatBotId)) {
		body["chatBotId"] = request.ChatBotId
	}

	if !tea.BoolValue(util.IsUnset(request.ConversationType)) {
		body["conversationType"] = request.ConversationType
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OutTrackId)) {
		body["outTrackId"] = request.OutTrackId
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateData)) {
		body["privateData"] = request.PrivateData
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUserIdList)) {
		body["receiverUserIdList"] = request.ReceiverUserIdList
	}

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdType)) {
		body["userIdType"] = request.UserIdType
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
	_result = &InteractiveCardCreateInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("InteractiveCardCreateInstance"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/interactiveCards/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryGroupMuteStatus(request *QueryGroupMuteStatusRequest) (_result *QueryGroupMuteStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGroupMuteStatusHeaders{}
	_result = &QueryGroupMuteStatusResponse{}
	_body, _err := client.QueryGroupMuteStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryGroupMuteStatusWithOptions(request *QueryGroupMuteStatusRequest, headers *QueryGroupMuteStatusHeaders, runtime *util.RuntimeOptions) (_result *QueryGroupMuteStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		query["openConversationId"] = request.OpenConversationId
	}

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
	_result = &QueryGroupMuteStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryGroupMuteStatus"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/im/sceneGroups/muteSettings"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMembersOfGroupRole(request *QueryMembersOfGroupRoleRequest) (_result *QueryMembersOfGroupRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMembersOfGroupRoleHeaders{}
	_result = &QueryMembersOfGroupRoleResponse{}
	_body, _err := client.QueryMembersOfGroupRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMembersOfGroupRoleWithOptions(request *QueryMembersOfGroupRoleRequest, headers *QueryMembersOfGroupRoleHeaders, runtime *util.RuntimeOptions) (_result *QueryMembersOfGroupRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenRoleId)) {
		body["openRoleId"] = request.OpenRoleId
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamp)) {
		body["timestamp"] = request.Timestamp
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
	_result = &QueryMembersOfGroupRoleResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryMembersOfGroupRole"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/sceneGroups/roles/members/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendInteractiveCard(request *SendInteractiveCardRequest) (_result *SendInteractiveCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendInteractiveCardHeaders{}
	_result = &SendInteractiveCardResponse{}
	_body, _err := client.SendInteractiveCardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendInteractiveCardWithOptions(request *SendInteractiveCardRequest, headers *SendInteractiveCardHeaders, runtime *util.RuntimeOptions) (_result *SendInteractiveCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AtOpenIds)) {
		body["atOpenIds"] = request.AtOpenIds
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackRouteKey)) {
		body["callbackRouteKey"] = request.CallbackRouteKey
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CardData))) {
		body["cardData"] = request.CardData
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CardOptions))) {
		body["cardOptions"] = request.CardOptions
	}

	if !tea.BoolValue(util.IsUnset(request.CardTemplateId)) {
		body["cardTemplateId"] = request.CardTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.ChatBotId)) {
		body["chatBotId"] = request.ChatBotId
	}

	if !tea.BoolValue(util.IsUnset(request.ConversationType)) {
		body["conversationType"] = request.ConversationType
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OutTrackId)) {
		body["outTrackId"] = request.OutTrackId
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateData)) {
		body["privateData"] = request.PrivateData
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUserIdList)) {
		body["receiverUserIdList"] = request.ReceiverUserIdList
	}

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdType)) {
		body["userIdType"] = request.UserIdType
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
	_result = &SendInteractiveCardResponse{}
	_body, _err := client.DoROARequest(tea.String("SendInteractiveCard"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/interactiveCards/send"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendRobotInteractiveCard(request *SendRobotInteractiveCardRequest) (_result *SendRobotInteractiveCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendRobotInteractiveCardHeaders{}
	_result = &SendRobotInteractiveCardResponse{}
	_body, _err := client.SendRobotInteractiveCardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendRobotInteractiveCardWithOptions(request *SendRobotInteractiveCardRequest, headers *SendRobotInteractiveCardHeaders, runtime *util.RuntimeOptions) (_result *SendRobotInteractiveCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CardBizId)) {
		body["cardBizId"] = request.CardBizId
	}

	if !tea.BoolValue(util.IsUnset(request.CardData)) {
		body["cardData"] = request.CardData
	}

	if !tea.BoolValue(util.IsUnset(request.CardTemplateId)) {
		body["cardTemplateId"] = request.CardTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SendOptions))) {
		body["sendOptions"] = request.SendOptions
	}

	if !tea.BoolValue(util.IsUnset(request.SingleChatReceiver)) {
		body["singleChatReceiver"] = request.SingleChatReceiver
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
	_result = &SendRobotInteractiveCardResponse{}
	_body, _err := client.DoROARequest(tea.String("SendRobotInteractiveCard"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/v1.0/robot/interactiveCards/send"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendTemplateInteractiveCard(request *SendTemplateInteractiveCardRequest) (_result *SendTemplateInteractiveCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendTemplateInteractiveCardHeaders{}
	_result = &SendTemplateInteractiveCardResponse{}
	_body, _err := client.SendTemplateInteractiveCardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendTemplateInteractiveCardWithOptions(request *SendTemplateInteractiveCardRequest, headers *SendTemplateInteractiveCardHeaders, runtime *util.RuntimeOptions) (_result *SendTemplateInteractiveCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallbackUrl)) {
		body["callbackUrl"] = request.CallbackUrl
	}

	if !tea.BoolValue(util.IsUnset(request.CardData)) {
		body["cardData"] = request.CardData
	}

	if !tea.BoolValue(util.IsUnset(request.CardTemplateId)) {
		body["cardTemplateId"] = request.CardTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OutTrackId)) {
		body["outTrackId"] = request.OutTrackId
	}

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SendOptions))) {
		body["sendOptions"] = request.SendOptions
	}

	if !tea.BoolValue(util.IsUnset(request.SingleChatReceiver)) {
		body["singleChatReceiver"] = request.SingleChatReceiver
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
	_result = &SendTemplateInteractiveCardResponse{}
	_body, _err := client.DoROARequest(tea.String("SendTemplateInteractiveCard"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/interactiveCards/templates/send"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TopboxClose(request *TopboxCloseRequest) (_result *TopboxCloseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TopboxCloseHeaders{}
	_result = &TopboxCloseResponse{}
	_body, _err := client.TopboxCloseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TopboxCloseWithOptions(request *TopboxCloseRequest, headers *TopboxCloseHeaders, runtime *util.RuntimeOptions) (_result *TopboxCloseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CoolAppCode)) {
		body["coolAppCode"] = request.CoolAppCode
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OutTrackId)) {
		body["outTrackId"] = request.OutTrackId
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
	_result = &TopboxCloseResponse{}
	_body, _err := client.DoROARequest(tea.String("TopboxClose"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/topBoxes/close"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TopboxOpen(request *TopboxOpenRequest) (_result *TopboxOpenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TopboxOpenHeaders{}
	_result = &TopboxOpenResponse{}
	_body, _err := client.TopboxOpenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TopboxOpenWithOptions(request *TopboxOpenRequest, headers *TopboxOpenHeaders, runtime *util.RuntimeOptions) (_result *TopboxOpenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CoolAppCode)) {
		body["coolAppCode"] = request.CoolAppCode
	}

	if !tea.BoolValue(util.IsUnset(request.ExpiredTime)) {
		body["expiredTime"] = request.ExpiredTime
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OutTrackId)) {
		body["outTrackId"] = request.OutTrackId
	}

	if !tea.BoolValue(util.IsUnset(request.Platforms)) {
		body["platforms"] = request.Platforms
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
	_result = &TopboxOpenResponse{}
	_body, _err := client.DoROARequest(tea.String("TopboxOpen"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/topBoxes/open"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGroupPermission(request *UpdateGroupPermissionRequest) (_result *UpdateGroupPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateGroupPermissionHeaders{}
	_result = &UpdateGroupPermissionResponse{}
	_body, _err := client.UpdateGroupPermissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGroupPermissionWithOptions(request *UpdateGroupPermissionRequest, headers *UpdateGroupPermissionHeaders, runtime *util.RuntimeOptions) (_result *UpdateGroupPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.PermissionGroup)) {
		body["permissionGroup"] = request.PermissionGroup
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
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
	_result = &UpdateGroupPermissionResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateGroupPermission"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/im/sceneGroups/permissions"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGroupSubAdmin(request *UpdateGroupSubAdminRequest) (_result *UpdateGroupSubAdminResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateGroupSubAdminHeaders{}
	_result = &UpdateGroupSubAdminResponse{}
	_body, _err := client.UpdateGroupSubAdminWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGroupSubAdminWithOptions(request *UpdateGroupSubAdminRequest, headers *UpdateGroupSubAdminHeaders, runtime *util.RuntimeOptions) (_result *UpdateGroupSubAdminResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		body["role"] = request.Role
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
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
	_result = &UpdateGroupSubAdminResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateGroupSubAdmin"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/im/sceneGroups/subAdmins"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInteractiveCard(request *UpdateInteractiveCardRequest) (_result *UpdateInteractiveCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateInteractiveCardHeaders{}
	_result = &UpdateInteractiveCardResponse{}
	_body, _err := client.UpdateInteractiveCardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInteractiveCardWithOptions(request *UpdateInteractiveCardRequest, headers *UpdateInteractiveCardHeaders, runtime *util.RuntimeOptions) (_result *UpdateInteractiveCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CardData))) {
		body["cardData"] = request.CardData
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CardOptions))) {
		body["cardOptions"] = request.CardOptions
	}

	if !tea.BoolValue(util.IsUnset(request.OutTrackId)) {
		body["outTrackId"] = request.OutTrackId
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateData)) {
		body["privateData"] = request.PrivateData
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdType)) {
		body["userIdType"] = request.UserIdType
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
	_result = &UpdateInteractiveCardResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateInteractiveCard"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/im/interactiveCards"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMemberBanWords(request *UpdateMemberBanWordsRequest) (_result *UpdateMemberBanWordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateMemberBanWordsHeaders{}
	_result = &UpdateMemberBanWordsResponse{}
	_body, _err := client.UpdateMemberBanWordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMemberBanWordsWithOptions(request *UpdateMemberBanWordsRequest, headers *UpdateMemberBanWordsHeaders, runtime *util.RuntimeOptions) (_result *UpdateMemberBanWordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MuteDuration)) {
		body["muteDuration"] = request.MuteDuration
	}

	if !tea.BoolValue(util.IsUnset(request.MuteStatus)) {
		body["muteStatus"] = request.MuteStatus
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdList)) {
		body["userIdList"] = request.UserIdList
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
	_result = &UpdateMemberBanWordsResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateMemberBanWords"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/sceneGroups/muteMembers/set"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMemberGroupNick(request *UpdateMemberGroupNickRequest) (_result *UpdateMemberGroupNickResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateMemberGroupNickHeaders{}
	_result = &UpdateMemberGroupNickResponse{}
	_body, _err := client.UpdateMemberGroupNickWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMemberGroupNickWithOptions(request *UpdateMemberGroupNickRequest, headers *UpdateMemberGroupNickHeaders, runtime *util.RuntimeOptions) (_result *UpdateMemberGroupNickResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupNick)) {
		body["groupNick"] = request.GroupNick
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
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
	_result = &UpdateMemberGroupNickResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateMemberGroupNick"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/im/sceneGroups/members/groupNicks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTheGroupRolesOfGroupMember(request *UpdateTheGroupRolesOfGroupMemberRequest) (_result *UpdateTheGroupRolesOfGroupMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTheGroupRolesOfGroupMemberHeaders{}
	_result = &UpdateTheGroupRolesOfGroupMemberResponse{}
	_body, _err := client.UpdateTheGroupRolesOfGroupMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTheGroupRolesOfGroupMemberWithOptions(request *UpdateTheGroupRolesOfGroupMemberRequest, headers *UpdateTheGroupRolesOfGroupMemberHeaders, runtime *util.RuntimeOptions) (_result *UpdateTheGroupRolesOfGroupMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenRoleIds)) {
		body["openRoleIds"] = request.OpenRoleIds
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
	_result = &UpdateTheGroupRolesOfGroupMemberResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateTheGroupRolesOfGroupMember"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/im/sceneGroups/members/groupRoles"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
