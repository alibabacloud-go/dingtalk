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

type AutoOpenDingTalkConnectHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AutoOpenDingTalkConnectHeaders) String() string {
	return tea.Prettify(s)
}

func (s AutoOpenDingTalkConnectHeaders) GoString() string {
	return s.String()
}

func (s *AutoOpenDingTalkConnectHeaders) SetCommonHeaders(v map[string]*string) *AutoOpenDingTalkConnectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AutoOpenDingTalkConnectHeaders) SetXAcsDingtalkAccessToken(v string) *AutoOpenDingTalkConnectHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AutoOpenDingTalkConnectResponseBody struct {
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s AutoOpenDingTalkConnectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AutoOpenDingTalkConnectResponseBody) GoString() string {
	return s.String()
}

func (s *AutoOpenDingTalkConnectResponseBody) SetMessage(v string) *AutoOpenDingTalkConnectResponseBody {
	s.Message = &v
	return s
}

type AutoOpenDingTalkConnectResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AutoOpenDingTalkConnectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AutoOpenDingTalkConnectResponse) String() string {
	return tea.Prettify(s)
}

func (s AutoOpenDingTalkConnectResponse) GoString() string {
	return s.String()
}

func (s *AutoOpenDingTalkConnectResponse) SetHeaders(v map[string]*string) *AutoOpenDingTalkConnectResponse {
	s.Headers = v
	return s
}

func (s *AutoOpenDingTalkConnectResponse) SetBody(v *AutoOpenDingTalkConnectResponseBody) *AutoOpenDingTalkConnectResponse {
	s.Body = v
	return s
}

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

type ChatSubAdminUpdateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ChatSubAdminUpdateHeaders) String() string {
	return tea.Prettify(s)
}

func (s ChatSubAdminUpdateHeaders) GoString() string {
	return s.String()
}

func (s *ChatSubAdminUpdateHeaders) SetCommonHeaders(v map[string]*string) *ChatSubAdminUpdateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ChatSubAdminUpdateHeaders) SetXAcsDingtalkAccessToken(v string) *ChatSubAdminUpdateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ChatSubAdminUpdateRequest struct {
	// 开放群ID
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 设置2添加为管理员，设置3删除该管理员
	Role *int32 `json:"role,omitempty" xml:"role,omitempty"`
	// 企业员工工号列表
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s ChatSubAdminUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s ChatSubAdminUpdateRequest) GoString() string {
	return s.String()
}

func (s *ChatSubAdminUpdateRequest) SetOpenConversationId(v string) *ChatSubAdminUpdateRequest {
	s.OpenConversationId = &v
	return s
}

func (s *ChatSubAdminUpdateRequest) SetRole(v int32) *ChatSubAdminUpdateRequest {
	s.Role = &v
	return s
}

func (s *ChatSubAdminUpdateRequest) SetUserIds(v []*string) *ChatSubAdminUpdateRequest {
	s.UserIds = v
	return s
}

type ChatSubAdminUpdateResponseBody struct {
	// result
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ChatSubAdminUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChatSubAdminUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *ChatSubAdminUpdateResponseBody) SetSuccess(v string) *ChatSubAdminUpdateResponseBody {
	s.Success = &v
	return s
}

type ChatSubAdminUpdateResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ChatSubAdminUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChatSubAdminUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s ChatSubAdminUpdateResponse) GoString() string {
	return s.String()
}

func (s *ChatSubAdminUpdateResponse) SetHeaders(v map[string]*string) *ChatSubAdminUpdateResponse {
	s.Headers = v
	return s
}

func (s *ChatSubAdminUpdateResponse) SetBody(v *ChatSubAdminUpdateResponseBody) *ChatSubAdminUpdateResponse {
	s.Body = v
	return s
}

type CreateCoupleGroupConversationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateCoupleGroupConversationHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateCoupleGroupConversationHeaders) GoString() string {
	return s.String()
}

func (s *CreateCoupleGroupConversationHeaders) SetCommonHeaders(v map[string]*string) *CreateCoupleGroupConversationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateCoupleGroupConversationHeaders) SetXAcsDingtalkAccessToken(v string) *CreateCoupleGroupConversationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateCoupleGroupConversationRequest struct {
	// 钉外人员业务Id
	AppUserId *string `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
	// 群头像
	GroupAvatar *string `json:"groupAvatar,omitempty" xml:"groupAvatar,omitempty"`
	// 群名称
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// 钉外群主业务Id
	GroupOwnerId *string `json:"groupOwnerId,omitempty" xml:"groupOwnerId,omitempty"`
	// 群模板
	GroupTemplateId *string `json:"groupTemplateId,omitempty" xml:"groupTemplateId,omitempty"`
	// 操作者的钉钉Id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s CreateCoupleGroupConversationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCoupleGroupConversationRequest) GoString() string {
	return s.String()
}

func (s *CreateCoupleGroupConversationRequest) SetAppUserId(v string) *CreateCoupleGroupConversationRequest {
	s.AppUserId = &v
	return s
}

func (s *CreateCoupleGroupConversationRequest) SetGroupAvatar(v string) *CreateCoupleGroupConversationRequest {
	s.GroupAvatar = &v
	return s
}

func (s *CreateCoupleGroupConversationRequest) SetGroupName(v string) *CreateCoupleGroupConversationRequest {
	s.GroupName = &v
	return s
}

func (s *CreateCoupleGroupConversationRequest) SetGroupOwnerId(v string) *CreateCoupleGroupConversationRequest {
	s.GroupOwnerId = &v
	return s
}

func (s *CreateCoupleGroupConversationRequest) SetGroupTemplateId(v string) *CreateCoupleGroupConversationRequest {
	s.GroupTemplateId = &v
	return s
}

func (s *CreateCoupleGroupConversationRequest) SetOperatorId(v string) *CreateCoupleGroupConversationRequest {
	s.OperatorId = &v
	return s
}

type CreateCoupleGroupConversationResponseBody struct {
	// 群会话Id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s CreateCoupleGroupConversationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCoupleGroupConversationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCoupleGroupConversationResponseBody) SetOpenConversationId(v string) *CreateCoupleGroupConversationResponseBody {
	s.OpenConversationId = &v
	return s
}

type CreateCoupleGroupConversationResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCoupleGroupConversationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCoupleGroupConversationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCoupleGroupConversationResponse) GoString() string {
	return s.String()
}

func (s *CreateCoupleGroupConversationResponse) SetHeaders(v map[string]*string) *CreateCoupleGroupConversationResponse {
	s.Headers = v
	return s
}

func (s *CreateCoupleGroupConversationResponse) SetBody(v *CreateCoupleGroupConversationResponseBody) *CreateCoupleGroupConversationResponse {
	s.Body = v
	return s
}

type CreateGroupConversationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateGroupConversationHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupConversationHeaders) GoString() string {
	return s.String()
}

func (s *CreateGroupConversationHeaders) SetCommonHeaders(v map[string]*string) *CreateGroupConversationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateGroupConversationHeaders) SetXAcsDingtalkAccessToken(v string) *CreateGroupConversationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateGroupConversationRequest struct {
	// C端客户成员列表
	AppUserIds []*string `json:"appUserIds,omitempty" xml:"appUserIds,omitempty" type:"Repeated"`
	// 群头像
	GroupAvatar *string `json:"groupAvatar,omitempty" xml:"groupAvatar,omitempty"`
	// 群名称
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// 群主钉钉Id
	GroupOwnerId *string `json:"groupOwnerId,omitempty" xml:"groupOwnerId,omitempty"`
	// 群模板
	GroupTemplateId *string `json:"groupTemplateId,omitempty" xml:"groupTemplateId,omitempty"`
	// 操作者的钉钉Id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// B端客服成员列表
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s CreateGroupConversationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupConversationRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupConversationRequest) SetAppUserIds(v []*string) *CreateGroupConversationRequest {
	s.AppUserIds = v
	return s
}

func (s *CreateGroupConversationRequest) SetGroupAvatar(v string) *CreateGroupConversationRequest {
	s.GroupAvatar = &v
	return s
}

func (s *CreateGroupConversationRequest) SetGroupName(v string) *CreateGroupConversationRequest {
	s.GroupName = &v
	return s
}

func (s *CreateGroupConversationRequest) SetGroupOwnerId(v string) *CreateGroupConversationRequest {
	s.GroupOwnerId = &v
	return s
}

func (s *CreateGroupConversationRequest) SetGroupTemplateId(v string) *CreateGroupConversationRequest {
	s.GroupTemplateId = &v
	return s
}

func (s *CreateGroupConversationRequest) SetOperatorId(v string) *CreateGroupConversationRequest {
	s.OperatorId = &v
	return s
}

func (s *CreateGroupConversationRequest) SetUserIds(v []*string) *CreateGroupConversationRequest {
	s.UserIds = v
	return s
}

type CreateGroupConversationResponseBody struct {
	// 添加成功的C端客户列表
	AppUserIds []*string `json:"appUserIds,omitempty" xml:"appUserIds,omitempty" type:"Repeated"`
	// 群会话Id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 添加成功的B端客服列表
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s CreateGroupConversationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupConversationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupConversationResponseBody) SetAppUserIds(v []*string) *CreateGroupConversationResponseBody {
	s.AppUserIds = v
	return s
}

func (s *CreateGroupConversationResponseBody) SetOpenConversationId(v string) *CreateGroupConversationResponseBody {
	s.OpenConversationId = &v
	return s
}

func (s *CreateGroupConversationResponseBody) SetUserIds(v []*string) *CreateGroupConversationResponseBody {
	s.UserIds = v
	return s
}

type CreateGroupConversationResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGroupConversationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGroupConversationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupConversationResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupConversationResponse) SetHeaders(v map[string]*string) *CreateGroupConversationResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupConversationResponse) SetBody(v *CreateGroupConversationResponseBody) *CreateGroupConversationResponse {
	s.Body = v
	return s
}

type CreateInterconnectionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateInterconnectionHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateInterconnectionHeaders) GoString() string {
	return s.String()
}

func (s *CreateInterconnectionHeaders) SetCommonHeaders(v map[string]*string) *CreateInterconnectionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateInterconnectionHeaders) SetXAcsDingtalkAccessToken(v string) *CreateInterconnectionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateInterconnectionRequest struct {
	// bc关系列表
	Interconnections []*CreateInterconnectionRequestInterconnections `json:"interconnections,omitempty" xml:"interconnections,omitempty" type:"Repeated"`
	// 参数签名
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
}

func (s CreateInterconnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInterconnectionRequest) GoString() string {
	return s.String()
}

func (s *CreateInterconnectionRequest) SetInterconnections(v []*CreateInterconnectionRequestInterconnections) *CreateInterconnectionRequest {
	s.Interconnections = v
	return s
}

func (s *CreateInterconnectionRequest) SetSignature(v string) *CreateInterconnectionRequest {
	s.Signature = &v
	return s
}

type CreateInterconnectionRequestInterconnections struct {
	// 客户头像链接
	AppUserAvatar *string `json:"appUserAvatar,omitempty" xml:"appUserAvatar,omitempty"`
	// 客户头像类型，取值：
	// 1：http
	AppUserAvatarMediaType *int32 `json:"appUserAvatarMediaType,omitempty" xml:"appUserAvatarMediaType,omitempty"`
	// 客户动态
	AppUserDynamics *string `json:"appUserDynamics,omitempty" xml:"appUserDynamics,omitempty"`
	// 客户业务系统唯一标识
	AppUserId *string `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
	// 客户手机号
	AppUserMobile *string `json:"appUserMobile,omitempty" xml:"appUserMobile,omitempty"`
	// 客户名称
	AppUserName *string `json:"appUserName,omitempty" xml:"appUserName,omitempty"`
	// 客户渠道code
	ChannelCode *string `json:"channelCode,omitempty" xml:"channelCode,omitempty"`
	// 客服钉钉Id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateInterconnectionRequestInterconnections) String() string {
	return tea.Prettify(s)
}

func (s CreateInterconnectionRequestInterconnections) GoString() string {
	return s.String()
}

func (s *CreateInterconnectionRequestInterconnections) SetAppUserAvatar(v string) *CreateInterconnectionRequestInterconnections {
	s.AppUserAvatar = &v
	return s
}

func (s *CreateInterconnectionRequestInterconnections) SetAppUserAvatarMediaType(v int32) *CreateInterconnectionRequestInterconnections {
	s.AppUserAvatarMediaType = &v
	return s
}

func (s *CreateInterconnectionRequestInterconnections) SetAppUserDynamics(v string) *CreateInterconnectionRequestInterconnections {
	s.AppUserDynamics = &v
	return s
}

func (s *CreateInterconnectionRequestInterconnections) SetAppUserId(v string) *CreateInterconnectionRequestInterconnections {
	s.AppUserId = &v
	return s
}

func (s *CreateInterconnectionRequestInterconnections) SetAppUserMobile(v string) *CreateInterconnectionRequestInterconnections {
	s.AppUserMobile = &v
	return s
}

func (s *CreateInterconnectionRequestInterconnections) SetAppUserName(v string) *CreateInterconnectionRequestInterconnections {
	s.AppUserName = &v
	return s
}

func (s *CreateInterconnectionRequestInterconnections) SetChannelCode(v string) *CreateInterconnectionRequestInterconnections {
	s.ChannelCode = &v
	return s
}

func (s *CreateInterconnectionRequestInterconnections) SetUserId(v string) *CreateInterconnectionRequestInterconnections {
	s.UserId = &v
	return s
}

type CreateInterconnectionResponseBody struct {
	// 失败的bc关系列表
	Results []*CreateInterconnectionResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s CreateInterconnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInterconnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInterconnectionResponseBody) SetResults(v []*CreateInterconnectionResponseBodyResults) *CreateInterconnectionResponseBody {
	s.Results = v
	return s
}

type CreateInterconnectionResponseBodyResults struct {
	// 客户业务身份唯一标识
	AppUserId *string `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
	// 客服钉钉Id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateInterconnectionResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s CreateInterconnectionResponseBodyResults) GoString() string {
	return s.String()
}

func (s *CreateInterconnectionResponseBodyResults) SetAppUserId(v string) *CreateInterconnectionResponseBodyResults {
	s.AppUserId = &v
	return s
}

func (s *CreateInterconnectionResponseBodyResults) SetUserId(v string) *CreateInterconnectionResponseBodyResults {
	s.UserId = &v
	return s
}

type CreateInterconnectionResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInterconnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInterconnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInterconnectionResponse) GoString() string {
	return s.String()
}

func (s *CreateInterconnectionResponse) SetHeaders(v map[string]*string) *CreateInterconnectionResponse {
	s.Headers = v
	return s
}

func (s *CreateInterconnectionResponse) SetBody(v *CreateInterconnectionResponseBody) *CreateInterconnectionResponse {
	s.Body = v
	return s
}

type CreateStoreGroupConversationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateStoreGroupConversationHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateStoreGroupConversationHeaders) GoString() string {
	return s.String()
}

func (s *CreateStoreGroupConversationHeaders) SetCommonHeaders(v map[string]*string) *CreateStoreGroupConversationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateStoreGroupConversationHeaders) SetXAcsDingtalkAccessToken(v string) *CreateStoreGroupConversationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateStoreGroupConversationRequest struct {
	// 钉外用户在业务系统内的唯一标识
	AppUserId *string `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
	// 外部业务唯一标识（店铺唯一标识）
	BusinessUniqueKey *string `json:"businessUniqueKey,omitempty" xml:"businessUniqueKey,omitempty"`
	// 群头像
	GroupAvatar *string `json:"groupAvatar,omitempty" xml:"groupAvatar,omitempty"`
	// 群名称
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// 群模板
	GroupTemplateId *string `json:"groupTemplateId,omitempty" xml:"groupTemplateId,omitempty"`
	// 操作者在业务系统内的唯一标识
	OperatorId *string   `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	UserIds    []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s CreateStoreGroupConversationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStoreGroupConversationRequest) GoString() string {
	return s.String()
}

func (s *CreateStoreGroupConversationRequest) SetAppUserId(v string) *CreateStoreGroupConversationRequest {
	s.AppUserId = &v
	return s
}

func (s *CreateStoreGroupConversationRequest) SetBusinessUniqueKey(v string) *CreateStoreGroupConversationRequest {
	s.BusinessUniqueKey = &v
	return s
}

func (s *CreateStoreGroupConversationRequest) SetGroupAvatar(v string) *CreateStoreGroupConversationRequest {
	s.GroupAvatar = &v
	return s
}

func (s *CreateStoreGroupConversationRequest) SetGroupName(v string) *CreateStoreGroupConversationRequest {
	s.GroupName = &v
	return s
}

func (s *CreateStoreGroupConversationRequest) SetGroupTemplateId(v string) *CreateStoreGroupConversationRequest {
	s.GroupTemplateId = &v
	return s
}

func (s *CreateStoreGroupConversationRequest) SetOperatorId(v string) *CreateStoreGroupConversationRequest {
	s.OperatorId = &v
	return s
}

func (s *CreateStoreGroupConversationRequest) SetUserIds(v []*string) *CreateStoreGroupConversationRequest {
	s.UserIds = v
	return s
}

type CreateStoreGroupConversationResponseBody struct {
	// 群会话Id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s CreateStoreGroupConversationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStoreGroupConversationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStoreGroupConversationResponseBody) SetOpenConversationId(v string) *CreateStoreGroupConversationResponseBody {
	s.OpenConversationId = &v
	return s
}

type CreateStoreGroupConversationResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateStoreGroupConversationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateStoreGroupConversationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStoreGroupConversationResponse) GoString() string {
	return s.String()
}

func (s *CreateStoreGroupConversationResponse) SetHeaders(v map[string]*string) *CreateStoreGroupConversationResponse {
	s.Headers = v
	return s
}

func (s *CreateStoreGroupConversationResponse) SetBody(v *CreateStoreGroupConversationResponseBody) *CreateStoreGroupConversationResponse {
	s.Body = v
	return s
}

type GetConversationUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetConversationUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetConversationUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetConversationUrlHeaders) SetCommonHeaders(v map[string]*string) *GetConversationUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetConversationUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GetConversationUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetConversationUrlRequest struct {
	// C端用户在业务账号体系内的用户userid，长度限制为1~64个字符。
	AppUserId *string `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
	// C端客户渠道code。
	ChannelCode *string `json:"channelCode,omitempty" xml:"channelCode,omitempty"`
	// 群会话Id。
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// C端客户标识。
	SourceCode *string `json:"sourceCode,omitempty" xml:"sourceCode,omitempty"`
	// B端用户的钉钉userId。
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetConversationUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConversationUrlRequest) GoString() string {
	return s.String()
}

func (s *GetConversationUrlRequest) SetAppUserId(v string) *GetConversationUrlRequest {
	s.AppUserId = &v
	return s
}

func (s *GetConversationUrlRequest) SetChannelCode(v string) *GetConversationUrlRequest {
	s.ChannelCode = &v
	return s
}

func (s *GetConversationUrlRequest) SetOpenConversationId(v string) *GetConversationUrlRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GetConversationUrlRequest) SetSourceCode(v string) *GetConversationUrlRequest {
	s.SourceCode = &v
	return s
}

func (s *GetConversationUrlRequest) SetUserId(v string) *GetConversationUrlRequest {
	s.UserId = &v
	return s
}

type GetConversationUrlResponseBody struct {
	// 会话url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetConversationUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConversationUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetConversationUrlResponseBody) SetUrl(v string) *GetConversationUrlResponseBody {
	s.Url = &v
	return s
}

type GetConversationUrlResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetConversationUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConversationUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConversationUrlResponse) GoString() string {
	return s.String()
}

func (s *GetConversationUrlResponse) SetHeaders(v map[string]*string) *GetConversationUrlResponse {
	s.Headers = v
	return s
}

func (s *GetConversationUrlResponse) SetBody(v *GetConversationUrlResponseBody) *GetConversationUrlResponse {
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

type GroupBanWordsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GroupBanWordsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GroupBanWordsHeaders) GoString() string {
	return s.String()
}

func (s *GroupBanWordsHeaders) SetCommonHeaders(v map[string]*string) *GroupBanWordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GroupBanWordsHeaders) SetXAcsDingtalkAccessToken(v string) *GroupBanWordsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GroupBanWordsRequest struct {
	// 禁言模式
	BanWordsMode *int32 `json:"banWordsMode,omitempty" xml:"banWordsMode,omitempty"`
	// 开放群id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 扩展参数
	Options map[string]interface{} `json:"options,omitempty" xml:"options,omitempty"`
}

func (s GroupBanWordsRequest) String() string {
	return tea.Prettify(s)
}

func (s GroupBanWordsRequest) GoString() string {
	return s.String()
}

func (s *GroupBanWordsRequest) SetBanWordsMode(v int32) *GroupBanWordsRequest {
	s.BanWordsMode = &v
	return s
}

func (s *GroupBanWordsRequest) SetOpenConversationId(v string) *GroupBanWordsRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GroupBanWordsRequest) SetOptions(v map[string]interface{}) *GroupBanWordsRequest {
	s.Options = v
	return s
}

type GroupBanWordsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s GroupBanWordsResponse) String() string {
	return tea.Prettify(s)
}

func (s GroupBanWordsResponse) GoString() string {
	return s.String()
}

func (s *GroupBanWordsResponse) SetHeaders(v map[string]*string) *GroupBanWordsResponse {
	s.Headers = v
	return s
}

type GroupCapacityInquiryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GroupCapacityInquiryHeaders) String() string {
	return tea.Prettify(s)
}

func (s GroupCapacityInquiryHeaders) GoString() string {
	return s.String()
}

func (s *GroupCapacityInquiryHeaders) SetCommonHeaders(v map[string]*string) *GroupCapacityInquiryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GroupCapacityInquiryHeaders) SetXAcsDingtalkAccessToken(v string) *GroupCapacityInquiryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GroupCapacityInquiryRequest struct {
	// 有效生命周期
	EffectiveDuration *string `json:"effectiveDuration,omitempty" xml:"effectiveDuration,omitempty"`
	// 开放的群id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 当前操作人工号
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 扩展参数
	Options map[string]interface{} `json:"options,omitempty" xml:"options,omitempty"`
	// 目标容量
	TargetCapacity *int32 `json:"targetCapacity,omitempty" xml:"targetCapacity,omitempty"`
}

func (s GroupCapacityInquiryRequest) String() string {
	return tea.Prettify(s)
}

func (s GroupCapacityInquiryRequest) GoString() string {
	return s.String()
}

func (s *GroupCapacityInquiryRequest) SetEffectiveDuration(v string) *GroupCapacityInquiryRequest {
	s.EffectiveDuration = &v
	return s
}

func (s *GroupCapacityInquiryRequest) SetOpenConversationId(v string) *GroupCapacityInquiryRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GroupCapacityInquiryRequest) SetOperator(v string) *GroupCapacityInquiryRequest {
	s.Operator = &v
	return s
}

func (s *GroupCapacityInquiryRequest) SetOptions(v map[string]interface{}) *GroupCapacityInquiryRequest {
	s.Options = v
	return s
}

func (s *GroupCapacityInquiryRequest) SetTargetCapacity(v int32) *GroupCapacityInquiryRequest {
	s.TargetCapacity = &v
	return s
}

type GroupCapacityInquiryResponseBody struct {
	// 实际价格
	ActualPrice *int64 `json:"actualPrice,omitempty" xml:"actualPrice,omitempty"`
	// 群创建时间
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 当前容量
	CurrentCapacity *int32 `json:"currentCapacity,omitempty" xml:"currentCapacity,omitempty"`
	// 当前容量生效至何时
	CurrentEffectUntil *int64 `json:"currentEffectUntil,omitempty" xml:"currentEffectUntil,omitempty"`
	// 折扣
	Discount *int32 `json:"discount,omitempty" xml:"discount,omitempty"`
	// 扩展信息
	ExtInfo map[string]interface{} `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// 群主userId
	GroupOwner *string `json:"groupOwner,omitempty" xml:"groupOwner,omitempty"`
	// 群标题
	GroupTitle *string `json:"groupTitle,omitempty" xml:"groupTitle,omitempty"`
	// 标价
	MarkedPrice *int64 `json:"markedPrice,omitempty" xml:"markedPrice,omitempty"`
	// 群人数
	MemberCount *int32 `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	// 开放的群id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 当前操作人工号
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 目标容量
	TargetCapacity *int32 `json:"targetCapacity,omitempty" xml:"targetCapacity,omitempty"`
	// 目标容量生效至何时
	TargetEffectUntil *int64 `json:"targetEffectUntil,omitempty" xml:"targetEffectUntil,omitempty"`
	// 校验令牌
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s GroupCapacityInquiryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GroupCapacityInquiryResponseBody) GoString() string {
	return s.String()
}

func (s *GroupCapacityInquiryResponseBody) SetActualPrice(v int64) *GroupCapacityInquiryResponseBody {
	s.ActualPrice = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetCreatedAt(v int64) *GroupCapacityInquiryResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetCurrentCapacity(v int32) *GroupCapacityInquiryResponseBody {
	s.CurrentCapacity = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetCurrentEffectUntil(v int64) *GroupCapacityInquiryResponseBody {
	s.CurrentEffectUntil = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetDiscount(v int32) *GroupCapacityInquiryResponseBody {
	s.Discount = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetExtInfo(v map[string]interface{}) *GroupCapacityInquiryResponseBody {
	s.ExtInfo = v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetGroupOwner(v string) *GroupCapacityInquiryResponseBody {
	s.GroupOwner = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetGroupTitle(v string) *GroupCapacityInquiryResponseBody {
	s.GroupTitle = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetMarkedPrice(v int64) *GroupCapacityInquiryResponseBody {
	s.MarkedPrice = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetMemberCount(v int32) *GroupCapacityInquiryResponseBody {
	s.MemberCount = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetOpenConversationId(v string) *GroupCapacityInquiryResponseBody {
	s.OpenConversationId = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetOperator(v string) *GroupCapacityInquiryResponseBody {
	s.Operator = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetTargetCapacity(v int32) *GroupCapacityInquiryResponseBody {
	s.TargetCapacity = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetTargetEffectUntil(v int64) *GroupCapacityInquiryResponseBody {
	s.TargetEffectUntil = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetToken(v string) *GroupCapacityInquiryResponseBody {
	s.Token = &v
	return s
}

type GroupCapacityInquiryResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GroupCapacityInquiryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GroupCapacityInquiryResponse) String() string {
	return tea.Prettify(s)
}

func (s GroupCapacityInquiryResponse) GoString() string {
	return s.String()
}

func (s *GroupCapacityInquiryResponse) SetHeaders(v map[string]*string) *GroupCapacityInquiryResponse {
	s.Headers = v
	return s
}

func (s *GroupCapacityInquiryResponse) SetBody(v *GroupCapacityInquiryResponseBody) *GroupCapacityInquiryResponse {
	s.Body = v
	return s
}

type GroupCapacityOrderConfirmHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GroupCapacityOrderConfirmHeaders) String() string {
	return tea.Prettify(s)
}

func (s GroupCapacityOrderConfirmHeaders) GoString() string {
	return s.String()
}

func (s *GroupCapacityOrderConfirmHeaders) SetCommonHeaders(v map[string]*string) *GroupCapacityOrderConfirmHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GroupCapacityOrderConfirmHeaders) SetXAcsDingtalkAccessToken(v string) *GroupCapacityOrderConfirmHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GroupCapacityOrderConfirmRequest struct {
	// 操作人工号
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 订单号
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s GroupCapacityOrderConfirmRequest) String() string {
	return tea.Prettify(s)
}

func (s GroupCapacityOrderConfirmRequest) GoString() string {
	return s.String()
}

func (s *GroupCapacityOrderConfirmRequest) SetOperator(v string) *GroupCapacityOrderConfirmRequest {
	s.Operator = &v
	return s
}

func (s *GroupCapacityOrderConfirmRequest) SetOrderId(v string) *GroupCapacityOrderConfirmRequest {
	s.OrderId = &v
	return s
}

type GroupCapacityOrderConfirmResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s GroupCapacityOrderConfirmResponse) String() string {
	return tea.Prettify(s)
}

func (s GroupCapacityOrderConfirmResponse) GoString() string {
	return s.String()
}

func (s *GroupCapacityOrderConfirmResponse) SetHeaders(v map[string]*string) *GroupCapacityOrderConfirmResponse {
	s.Headers = v
	return s
}

type GroupCapacityOrderPlaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GroupCapacityOrderPlaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s GroupCapacityOrderPlaceHeaders) GoString() string {
	return s.String()
}

func (s *GroupCapacityOrderPlaceHeaders) SetCommonHeaders(v map[string]*string) *GroupCapacityOrderPlaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GroupCapacityOrderPlaceHeaders) SetXAcsDingtalkAccessToken(v string) *GroupCapacityOrderPlaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GroupCapacityOrderPlaceRequest struct {
	// 实际价格
	ActualPrice *int64 `json:"actualPrice,omitempty" xml:"actualPrice,omitempty"`
	// 当前容量
	CurrentCapacity *int32 `json:"currentCapacity,omitempty" xml:"currentCapacity,omitempty"`
	// 当前操当前容量生效至何时
	CurrentEffectUntil *int64 `json:"currentEffectUntil,omitempty" xml:"currentEffectUntil,omitempty"`
	// 折扣
	Discount *int32 `json:"discount,omitempty" xml:"discount,omitempty"`
	// 扩展信息
	ExtInfo map[string]interface{} `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// 标价
	MarkedPrice *int64 `json:"markedPrice,omitempty" xml:"markedPrice,omitempty"`
	// 开放的群id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 当前操作人工号
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 目标容量
	TargetCapacity *int32 `json:"targetCapacity,omitempty" xml:"targetCapacity,omitempty"`
	// 目标容量生效至何时
	TargetEffectUntil *int64 `json:"targetEffectUntil,omitempty" xml:"targetEffectUntil,omitempty"`
	// 校验令牌
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s GroupCapacityOrderPlaceRequest) String() string {
	return tea.Prettify(s)
}

func (s GroupCapacityOrderPlaceRequest) GoString() string {
	return s.String()
}

func (s *GroupCapacityOrderPlaceRequest) SetActualPrice(v int64) *GroupCapacityOrderPlaceRequest {
	s.ActualPrice = &v
	return s
}

func (s *GroupCapacityOrderPlaceRequest) SetCurrentCapacity(v int32) *GroupCapacityOrderPlaceRequest {
	s.CurrentCapacity = &v
	return s
}

func (s *GroupCapacityOrderPlaceRequest) SetCurrentEffectUntil(v int64) *GroupCapacityOrderPlaceRequest {
	s.CurrentEffectUntil = &v
	return s
}

func (s *GroupCapacityOrderPlaceRequest) SetDiscount(v int32) *GroupCapacityOrderPlaceRequest {
	s.Discount = &v
	return s
}

func (s *GroupCapacityOrderPlaceRequest) SetExtInfo(v map[string]interface{}) *GroupCapacityOrderPlaceRequest {
	s.ExtInfo = v
	return s
}

func (s *GroupCapacityOrderPlaceRequest) SetMarkedPrice(v int64) *GroupCapacityOrderPlaceRequest {
	s.MarkedPrice = &v
	return s
}

func (s *GroupCapacityOrderPlaceRequest) SetOpenConversationId(v string) *GroupCapacityOrderPlaceRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GroupCapacityOrderPlaceRequest) SetOperator(v string) *GroupCapacityOrderPlaceRequest {
	s.Operator = &v
	return s
}

func (s *GroupCapacityOrderPlaceRequest) SetTargetCapacity(v int32) *GroupCapacityOrderPlaceRequest {
	s.TargetCapacity = &v
	return s
}

func (s *GroupCapacityOrderPlaceRequest) SetTargetEffectUntil(v int64) *GroupCapacityOrderPlaceRequest {
	s.TargetEffectUntil = &v
	return s
}

func (s *GroupCapacityOrderPlaceRequest) SetToken(v string) *GroupCapacityOrderPlaceRequest {
	s.Token = &v
	return s
}

type GroupCapacityOrderPlaceResponseBody struct {
	// 实际价格
	ActualPrice *int64 `json:"actualPrice,omitempty" xml:"actualPrice,omitempty"`
	// 当前容量
	CurrentCapacity *int32 `json:"currentCapacity,omitempty" xml:"currentCapacity,omitempty"`
	// 当前容量生效至何时
	CurrentEffectUntil *int64 `json:"currentEffectUntil,omitempty" xml:"currentEffectUntil,omitempty"`
	// 折扣
	Discount *int32 `json:"discount,omitempty" xml:"discount,omitempty"`
	// 扩展信息
	ExtInfo map[string]*string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// 标价
	MarkedPrice *int64 `json:"markedPrice,omitempty" xml:"markedPrice,omitempty"`
	// 群标识
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 当前操作人工号
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 订单号
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// 目标容量
	TargetCapacity *int32 `json:"targetCapacity,omitempty" xml:"targetCapacity,omitempty"`
	// 目标容量生效至何时
	TargetEffectUntil *int64 `json:"targetEffectUntil,omitempty" xml:"targetEffectUntil,omitempty"`
	// 校验令牌
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s GroupCapacityOrderPlaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GroupCapacityOrderPlaceResponseBody) GoString() string {
	return s.String()
}

func (s *GroupCapacityOrderPlaceResponseBody) SetActualPrice(v int64) *GroupCapacityOrderPlaceResponseBody {
	s.ActualPrice = &v
	return s
}

func (s *GroupCapacityOrderPlaceResponseBody) SetCurrentCapacity(v int32) *GroupCapacityOrderPlaceResponseBody {
	s.CurrentCapacity = &v
	return s
}

func (s *GroupCapacityOrderPlaceResponseBody) SetCurrentEffectUntil(v int64) *GroupCapacityOrderPlaceResponseBody {
	s.CurrentEffectUntil = &v
	return s
}

func (s *GroupCapacityOrderPlaceResponseBody) SetDiscount(v int32) *GroupCapacityOrderPlaceResponseBody {
	s.Discount = &v
	return s
}

func (s *GroupCapacityOrderPlaceResponseBody) SetExtInfo(v map[string]*string) *GroupCapacityOrderPlaceResponseBody {
	s.ExtInfo = v
	return s
}

func (s *GroupCapacityOrderPlaceResponseBody) SetMarkedPrice(v int64) *GroupCapacityOrderPlaceResponseBody {
	s.MarkedPrice = &v
	return s
}

func (s *GroupCapacityOrderPlaceResponseBody) SetOpenConversationId(v string) *GroupCapacityOrderPlaceResponseBody {
	s.OpenConversationId = &v
	return s
}

func (s *GroupCapacityOrderPlaceResponseBody) SetOperator(v string) *GroupCapacityOrderPlaceResponseBody {
	s.Operator = &v
	return s
}

func (s *GroupCapacityOrderPlaceResponseBody) SetOrderId(v string) *GroupCapacityOrderPlaceResponseBody {
	s.OrderId = &v
	return s
}

func (s *GroupCapacityOrderPlaceResponseBody) SetTargetCapacity(v int32) *GroupCapacityOrderPlaceResponseBody {
	s.TargetCapacity = &v
	return s
}

func (s *GroupCapacityOrderPlaceResponseBody) SetTargetEffectUntil(v int64) *GroupCapacityOrderPlaceResponseBody {
	s.TargetEffectUntil = &v
	return s
}

func (s *GroupCapacityOrderPlaceResponseBody) SetToken(v string) *GroupCapacityOrderPlaceResponseBody {
	s.Token = &v
	return s
}

type GroupCapacityOrderPlaceResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GroupCapacityOrderPlaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GroupCapacityOrderPlaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GroupCapacityOrderPlaceResponse) GoString() string {
	return s.String()
}

func (s *GroupCapacityOrderPlaceResponse) SetHeaders(v map[string]*string) *GroupCapacityOrderPlaceResponse {
	s.Headers = v
	return s
}

func (s *GroupCapacityOrderPlaceResponse) SetBody(v *GroupCapacityOrderPlaceResponseBody) *GroupCapacityOrderPlaceResponse {
	s.Body = v
	return s
}

type GroupManageQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GroupManageQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s GroupManageQueryHeaders) GoString() string {
	return s.String()
}

func (s *GroupManageQueryHeaders) SetCommonHeaders(v map[string]*string) *GroupManageQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GroupManageQueryHeaders) SetXAcsDingtalkAccessToken(v string) *GroupManageQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GroupManageQueryRequest struct {
	// 建群时间不早于（辅助性条件，结合非排他条件使用）
	CreatedAfter *int64 `json:"createdAfter,omitempty" xml:"createdAfter,omitempty"`
	// 群号
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// 群成员样例工号列表
	GroupMemberSamples []*string `json:"groupMemberSamples,omitempty" xml:"groupMemberSamples,omitempty" type:"Repeated"`
	// 群主工号
	GroupOwner *string `json:"groupOwner,omitempty" xml:"groupOwner,omitempty"`
	// 群标题关键词列表
	GroupTitleKeywords []*string `json:"groupTitleKeywords,omitempty" xml:"groupTitleKeywords,omitempty" type:"Repeated"`
	// 群链接
	GroupUrl *string `json:"groupUrl,omitempty" xml:"groupUrl,omitempty"`
	// 分页拉取的页大小, 最大不可超过1000
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 群成员数下限（辅助性条件，结合非排他条件使用）
	MembersOver *int32 `json:"membersOver,omitempty" xml:"membersOver,omitempty"`
	// 分页拉取游标, 若不指定，则从头开始拉
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 开放群id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s GroupManageQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s GroupManageQueryRequest) GoString() string {
	return s.String()
}

func (s *GroupManageQueryRequest) SetCreatedAfter(v int64) *GroupManageQueryRequest {
	s.CreatedAfter = &v
	return s
}

func (s *GroupManageQueryRequest) SetGroupId(v string) *GroupManageQueryRequest {
	s.GroupId = &v
	return s
}

func (s *GroupManageQueryRequest) SetGroupMemberSamples(v []*string) *GroupManageQueryRequest {
	s.GroupMemberSamples = v
	return s
}

func (s *GroupManageQueryRequest) SetGroupOwner(v string) *GroupManageQueryRequest {
	s.GroupOwner = &v
	return s
}

func (s *GroupManageQueryRequest) SetGroupTitleKeywords(v []*string) *GroupManageQueryRequest {
	s.GroupTitleKeywords = v
	return s
}

func (s *GroupManageQueryRequest) SetGroupUrl(v string) *GroupManageQueryRequest {
	s.GroupUrl = &v
	return s
}

func (s *GroupManageQueryRequest) SetMaxResults(v int32) *GroupManageQueryRequest {
	s.MaxResults = &v
	return s
}

func (s *GroupManageQueryRequest) SetMembersOver(v int32) *GroupManageQueryRequest {
	s.MembersOver = &v
	return s
}

func (s *GroupManageQueryRequest) SetNextToken(v string) *GroupManageQueryRequest {
	s.NextToken = &v
	return s
}

func (s *GroupManageQueryRequest) SetOpenConversationId(v string) *GroupManageQueryRequest {
	s.OpenConversationId = &v
	return s
}

type GroupManageQueryResponseBody struct {
	// 群信息列表
	GroupInfoList []*GroupManageQueryResponseBodyGroupInfoList `json:"groupInfoList,omitempty" xml:"groupInfoList,omitempty" type:"Repeated"`
	// 分页拉取时, 是否还有更多
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 分页拉取游标, 请求下一页时回传
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GroupManageQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GroupManageQueryResponseBody) GoString() string {
	return s.String()
}

func (s *GroupManageQueryResponseBody) SetGroupInfoList(v []*GroupManageQueryResponseBodyGroupInfoList) *GroupManageQueryResponseBody {
	s.GroupInfoList = v
	return s
}

func (s *GroupManageQueryResponseBody) SetHasMore(v bool) *GroupManageQueryResponseBody {
	s.HasMore = &v
	return s
}

func (s *GroupManageQueryResponseBody) SetNextToken(v string) *GroupManageQueryResponseBody {
	s.NextToken = &v
	return s
}

type GroupManageQueryResponseBodyGroupInfoList struct {
	// 禁言模式
	BanWordsMode *int32 `json:"banWordsMode,omitempty" xml:"banWordsMode,omitempty"`
	// 群容量
	Capacity *int32 `json:"capacity,omitempty" xml:"capacity,omitempty"`
	// 群创建时间
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 扩展信息
	ExtInfo        map[string]interface{} `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	GroupAdminList []*string              `json:"groupAdminList,omitempty" xml:"groupAdminList,omitempty" type:"Repeated"`
	// 群主userid
	GroupOwner *string `json:"groupOwner,omitempty" xml:"groupOwner,omitempty"`
	// 群标题
	GroupTitle *string `json:"groupTitle,omitempty" xml:"groupTitle,omitempty"`
	// 当前群人数
	MemberCount *int32 `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	// 开放的群id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 群类型
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GroupManageQueryResponseBodyGroupInfoList) String() string {
	return tea.Prettify(s)
}

func (s GroupManageQueryResponseBodyGroupInfoList) GoString() string {
	return s.String()
}

func (s *GroupManageQueryResponseBodyGroupInfoList) SetBanWordsMode(v int32) *GroupManageQueryResponseBodyGroupInfoList {
	s.BanWordsMode = &v
	return s
}

func (s *GroupManageQueryResponseBodyGroupInfoList) SetCapacity(v int32) *GroupManageQueryResponseBodyGroupInfoList {
	s.Capacity = &v
	return s
}

func (s *GroupManageQueryResponseBodyGroupInfoList) SetCreatedAt(v int64) *GroupManageQueryResponseBodyGroupInfoList {
	s.CreatedAt = &v
	return s
}

func (s *GroupManageQueryResponseBodyGroupInfoList) SetExtInfo(v map[string]interface{}) *GroupManageQueryResponseBodyGroupInfoList {
	s.ExtInfo = v
	return s
}

func (s *GroupManageQueryResponseBodyGroupInfoList) SetGroupAdminList(v []*string) *GroupManageQueryResponseBodyGroupInfoList {
	s.GroupAdminList = v
	return s
}

func (s *GroupManageQueryResponseBodyGroupInfoList) SetGroupOwner(v string) *GroupManageQueryResponseBodyGroupInfoList {
	s.GroupOwner = &v
	return s
}

func (s *GroupManageQueryResponseBodyGroupInfoList) SetGroupTitle(v string) *GroupManageQueryResponseBodyGroupInfoList {
	s.GroupTitle = &v
	return s
}

func (s *GroupManageQueryResponseBodyGroupInfoList) SetMemberCount(v int32) *GroupManageQueryResponseBodyGroupInfoList {
	s.MemberCount = &v
	return s
}

func (s *GroupManageQueryResponseBodyGroupInfoList) SetOpenConversationId(v string) *GroupManageQueryResponseBodyGroupInfoList {
	s.OpenConversationId = &v
	return s
}

func (s *GroupManageQueryResponseBodyGroupInfoList) SetType(v string) *GroupManageQueryResponseBodyGroupInfoList {
	s.Type = &v
	return s
}

type GroupManageQueryResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GroupManageQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GroupManageQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s GroupManageQueryResponse) GoString() string {
	return s.String()
}

func (s *GroupManageQueryResponse) SetHeaders(v map[string]*string) *GroupManageQueryResponse {
	s.Headers = v
	return s
}

func (s *GroupManageQueryResponse) SetBody(v *GroupManageQueryResponseBody) *GroupManageQueryResponse {
	s.Body = v
	return s
}

type GroupManageReduceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GroupManageReduceHeaders) String() string {
	return tea.Prettify(s)
}

func (s GroupManageReduceHeaders) GoString() string {
	return s.String()
}

func (s *GroupManageReduceHeaders) SetCommonHeaders(v map[string]*string) *GroupManageReduceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GroupManageReduceHeaders) SetXAcsDingtalkAccessToken(v string) *GroupManageReduceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GroupManageReduceRequest struct {
	// 群容量限定值
	CapacityLimit *int32 `json:"capacityLimit,omitempty" xml:"capacityLimit,omitempty"`
	// 开放群id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 扩展参数
	Options map[string]interface{} `json:"options,omitempty" xml:"options,omitempty"`
}

func (s GroupManageReduceRequest) String() string {
	return tea.Prettify(s)
}

func (s GroupManageReduceRequest) GoString() string {
	return s.String()
}

func (s *GroupManageReduceRequest) SetCapacityLimit(v int32) *GroupManageReduceRequest {
	s.CapacityLimit = &v
	return s
}

func (s *GroupManageReduceRequest) SetOpenConversationId(v string) *GroupManageReduceRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GroupManageReduceRequest) SetOptions(v map[string]interface{}) *GroupManageReduceRequest {
	s.Options = v
	return s
}

type GroupManageReduceResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s GroupManageReduceResponse) String() string {
	return tea.Prettify(s)
}

func (s GroupManageReduceResponse) GoString() string {
	return s.String()
}

func (s *GroupManageReduceResponse) SetHeaders(v map[string]*string) *GroupManageReduceResponse {
	s.Headers = v
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
	// 是否开启卡片纯拉模式
	PullStrategy *bool `json:"pullStrategy,omitempty" xml:"pullStrategy,omitempty"`
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

func (s *InteractiveCardCreateInstanceRequest) SetPullStrategy(v bool) *InteractiveCardCreateInstanceRequest {
	s.PullStrategy = &v
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

type QueryGroupInfoByMemberAuthHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGroupInfoByMemberAuthHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoByMemberAuthHeaders) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoByMemberAuthHeaders) SetCommonHeaders(v map[string]*string) *QueryGroupInfoByMemberAuthHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGroupInfoByMemberAuthHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGroupInfoByMemberAuthHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGroupInfoByMemberAuthRequest struct {
	// 酷应用编码
	CoolAppCode *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	// 群的openConversationId
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s QueryGroupInfoByMemberAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoByMemberAuthRequest) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoByMemberAuthRequest) SetCoolAppCode(v string) *QueryGroupInfoByMemberAuthRequest {
	s.CoolAppCode = &v
	return s
}

func (s *QueryGroupInfoByMemberAuthRequest) SetOpenConversationId(v string) *QueryGroupInfoByMemberAuthRequest {
	s.OpenConversationId = &v
	return s
}

type QueryGroupInfoByMemberAuthResponseBody struct {
	// 群内总人数
	MemberCount *int32 `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
}

func (s QueryGroupInfoByMemberAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoByMemberAuthResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoByMemberAuthResponseBody) SetMemberCount(v int32) *QueryGroupInfoByMemberAuthResponseBody {
	s.MemberCount = &v
	return s
}

type QueryGroupInfoByMemberAuthResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryGroupInfoByMemberAuthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryGroupInfoByMemberAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoByMemberAuthResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoByMemberAuthResponse) SetHeaders(v map[string]*string) *QueryGroupInfoByMemberAuthResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupInfoByMemberAuthResponse) SetBody(v *QueryGroupInfoByMemberAuthResponseBody) *QueryGroupInfoByMemberAuthResponse {
	s.Body = v
	return s
}

type QueryGroupMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGroupMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberHeaders) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberHeaders) SetCommonHeaders(v map[string]*string) *QueryGroupMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGroupMemberHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGroupMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGroupMemberRequest struct {
	// 群会话Id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s QueryGroupMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberRequest) SetOpenConversationId(v string) *QueryGroupMemberRequest {
	s.OpenConversationId = &v
	return s
}

type QueryGroupMemberResponseBody struct {
	// 群成员列表
	GroupMembers []*QueryGroupMemberResponseBodyGroupMembers `json:"groupMembers,omitempty" xml:"groupMembers,omitempty" type:"Repeated"`
	// 群会话Id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s QueryGroupMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberResponseBody) SetGroupMembers(v []*QueryGroupMemberResponseBodyGroupMembers) *QueryGroupMemberResponseBody {
	s.GroupMembers = v
	return s
}

func (s *QueryGroupMemberResponseBody) SetOpenConversationId(v string) *QueryGroupMemberResponseBody {
	s.OpenConversationId = &v
	return s
}

type QueryGroupMemberResponseBodyGroupMembers struct {
	// 群成员头像
	GroupMemberAvatar *string `json:"groupMemberAvatar,omitempty" xml:"groupMemberAvatar,omitempty"`
	// 群成员动态信息
	GroupMemberDynamics *string `json:"groupMemberDynamics,omitempty" xml:"groupMemberDynamics,omitempty"`
	// 群成员Id
	GroupMemberId *string `json:"groupMemberId,omitempty" xml:"groupMemberId,omitempty"`
	// 群成员名称
	GroupMemberName *string `json:"groupMemberName,omitempty" xml:"groupMemberName,omitempty"`
	// 群成员类型
	GroupMemberType *int32 `json:"groupMemberType,omitempty" xml:"groupMemberType,omitempty"`
}

func (s QueryGroupMemberResponseBodyGroupMembers) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberResponseBodyGroupMembers) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberResponseBodyGroupMembers) SetGroupMemberAvatar(v string) *QueryGroupMemberResponseBodyGroupMembers {
	s.GroupMemberAvatar = &v
	return s
}

func (s *QueryGroupMemberResponseBodyGroupMembers) SetGroupMemberDynamics(v string) *QueryGroupMemberResponseBodyGroupMembers {
	s.GroupMemberDynamics = &v
	return s
}

func (s *QueryGroupMemberResponseBodyGroupMembers) SetGroupMemberId(v string) *QueryGroupMemberResponseBodyGroupMembers {
	s.GroupMemberId = &v
	return s
}

func (s *QueryGroupMemberResponseBodyGroupMembers) SetGroupMemberName(v string) *QueryGroupMemberResponseBodyGroupMembers {
	s.GroupMemberName = &v
	return s
}

func (s *QueryGroupMemberResponseBodyGroupMembers) SetGroupMemberType(v int32) *QueryGroupMemberResponseBodyGroupMembers {
	s.GroupMemberType = &v
	return s
}

type QueryGroupMemberResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryGroupMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberResponse) SetHeaders(v map[string]*string) *QueryGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupMemberResponse) SetBody(v *QueryGroupMemberResponseBody) *QueryGroupMemberResponse {
	s.Body = v
	return s
}

type QueryGroupMemberByMemberAuthHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGroupMemberByMemberAuthHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberByMemberAuthHeaders) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberByMemberAuthHeaders) SetCommonHeaders(v map[string]*string) *QueryGroupMemberByMemberAuthHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGroupMemberByMemberAuthHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGroupMemberByMemberAuthHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGroupMemberByMemberAuthRequest struct {
	// 酷应用编码
	CoolAppCode *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	// 群的openConversationId
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s QueryGroupMemberByMemberAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberByMemberAuthRequest) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberByMemberAuthRequest) SetCoolAppCode(v string) *QueryGroupMemberByMemberAuthRequest {
	s.CoolAppCode = &v
	return s
}

func (s *QueryGroupMemberByMemberAuthRequest) SetOpenConversationId(v string) *QueryGroupMemberByMemberAuthRequest {
	s.OpenConversationId = &v
	return s
}

type QueryGroupMemberByMemberAuthResponseBody struct {
	// 群成员列表
	GroupMemberList []*QueryGroupMemberByMemberAuthResponseBodyGroupMemberList `json:"groupMemberList,omitempty" xml:"groupMemberList,omitempty" type:"Repeated"`
}

func (s QueryGroupMemberByMemberAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberByMemberAuthResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberByMemberAuthResponseBody) SetGroupMemberList(v []*QueryGroupMemberByMemberAuthResponseBodyGroupMemberList) *QueryGroupMemberByMemberAuthResponseBody {
	s.GroupMemberList = v
	return s
}

type QueryGroupMemberByMemberAuthResponseBodyGroupMemberList struct {
	// 群内昵称
	//
	GroupNickName *string `json:"groupNickName,omitempty" xml:"groupNickName,omitempty"`
	// 企业内成员姓名
	OrgName *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	// 头像url
	ProfilePhotoUrl *string `json:"profilePhotoUrl,omitempty" xml:"profilePhotoUrl,omitempty"`
	// 员工id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryGroupMemberByMemberAuthResponseBodyGroupMemberList) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberByMemberAuthResponseBodyGroupMemberList) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberByMemberAuthResponseBodyGroupMemberList) SetGroupNickName(v string) *QueryGroupMemberByMemberAuthResponseBodyGroupMemberList {
	s.GroupNickName = &v
	return s
}

func (s *QueryGroupMemberByMemberAuthResponseBodyGroupMemberList) SetOrgName(v string) *QueryGroupMemberByMemberAuthResponseBodyGroupMemberList {
	s.OrgName = &v
	return s
}

func (s *QueryGroupMemberByMemberAuthResponseBodyGroupMemberList) SetProfilePhotoUrl(v string) *QueryGroupMemberByMemberAuthResponseBodyGroupMemberList {
	s.ProfilePhotoUrl = &v
	return s
}

func (s *QueryGroupMemberByMemberAuthResponseBodyGroupMemberList) SetUserId(v string) *QueryGroupMemberByMemberAuthResponseBodyGroupMemberList {
	s.UserId = &v
	return s
}

type QueryGroupMemberByMemberAuthResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryGroupMemberByMemberAuthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryGroupMemberByMemberAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberByMemberAuthResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberByMemberAuthResponse) SetHeaders(v map[string]*string) *QueryGroupMemberByMemberAuthResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupMemberByMemberAuthResponse) SetBody(v *QueryGroupMemberByMemberAuthResponseBody) *QueryGroupMemberByMemberAuthResponse {
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

type QuerySingleGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QuerySingleGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleGroupHeaders) GoString() string {
	return s.String()
}

func (s *QuerySingleGroupHeaders) SetCommonHeaders(v map[string]*string) *QuerySingleGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QuerySingleGroupHeaders) SetXAcsDingtalkAccessToken(v string) *QuerySingleGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QuerySingleGroupRequest struct {
	// 群成员列表
	GroupMembers []*QuerySingleGroupRequestGroupMembers `json:"groupMembers,omitempty" xml:"groupMembers,omitempty" type:"Repeated"`
	// 群模版Id
	GroupTemplateId *string `json:"groupTemplateId,omitempty" xml:"groupTemplateId,omitempty"`
}

func (s QuerySingleGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleGroupRequest) GoString() string {
	return s.String()
}

func (s *QuerySingleGroupRequest) SetGroupMembers(v []*QuerySingleGroupRequestGroupMembers) *QuerySingleGroupRequest {
	s.GroupMembers = v
	return s
}

func (s *QuerySingleGroupRequest) SetGroupTemplateId(v string) *QuerySingleGroupRequest {
	s.GroupTemplateId = &v
	return s
}

type QuerySingleGroupRequestGroupMembers struct {
	// 客户appUserId
	AppUserId *string `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
	// 客服钉钉Id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QuerySingleGroupRequestGroupMembers) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleGroupRequestGroupMembers) GoString() string {
	return s.String()
}

func (s *QuerySingleGroupRequestGroupMembers) SetAppUserId(v string) *QuerySingleGroupRequestGroupMembers {
	s.AppUserId = &v
	return s
}

func (s *QuerySingleGroupRequestGroupMembers) SetUserId(v string) *QuerySingleGroupRequestGroupMembers {
	s.UserId = &v
	return s
}

type QuerySingleGroupResponseBody struct {
	// 群会话列表
	OpenConversations []*QuerySingleGroupResponseBodyOpenConversations `json:"openConversations,omitempty" xml:"openConversations,omitempty" type:"Repeated"`
}

func (s QuerySingleGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySingleGroupResponseBody) SetOpenConversations(v []*QuerySingleGroupResponseBodyOpenConversations) *QuerySingleGroupResponseBody {
	s.OpenConversations = v
	return s
}

type QuerySingleGroupResponseBodyOpenConversations struct {
	// 客户appUserId
	AppUserId *string `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
	// 群会话Id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 客服钉钉Id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QuerySingleGroupResponseBodyOpenConversations) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleGroupResponseBodyOpenConversations) GoString() string {
	return s.String()
}

func (s *QuerySingleGroupResponseBodyOpenConversations) SetAppUserId(v string) *QuerySingleGroupResponseBodyOpenConversations {
	s.AppUserId = &v
	return s
}

func (s *QuerySingleGroupResponseBodyOpenConversations) SetOpenConversationId(v string) *QuerySingleGroupResponseBodyOpenConversations {
	s.OpenConversationId = &v
	return s
}

func (s *QuerySingleGroupResponseBodyOpenConversations) SetUserId(v string) *QuerySingleGroupResponseBodyOpenConversations {
	s.UserId = &v
	return s
}

type QuerySingleGroupResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySingleGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySingleGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleGroupResponse) GoString() string {
	return s.String()
}

func (s *QuerySingleGroupResponse) SetHeaders(v map[string]*string) *QuerySingleGroupResponse {
	s.Headers = v
	return s
}

func (s *QuerySingleGroupResponse) SetBody(v *QuerySingleGroupResponseBody) *QuerySingleGroupResponse {
	s.Body = v
	return s
}

type QueryUnReadMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUnReadMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUnReadMessageHeaders) GoString() string {
	return s.String()
}

func (s *QueryUnReadMessageHeaders) SetCommonHeaders(v map[string]*string) *QueryUnReadMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUnReadMessageHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUnReadMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUnReadMessageRequest struct {
	// C端客户appUserId
	AppUserId *string `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
	// 群会话Id列表
	OpenConversationIds []*string `json:"openConversationIds,omitempty" xml:"openConversationIds,omitempty" type:"Repeated"`
}

func (s QueryUnReadMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUnReadMessageRequest) GoString() string {
	return s.String()
}

func (s *QueryUnReadMessageRequest) SetAppUserId(v string) *QueryUnReadMessageRequest {
	s.AppUserId = &v
	return s
}

func (s *QueryUnReadMessageRequest) SetOpenConversationIds(v []*string) *QueryUnReadMessageRequest {
	s.OpenConversationIds = v
	return s
}

type QueryUnReadMessageResponseBody struct {
	// 未读消息数
	UnReadCount *int64 `json:"unReadCount,omitempty" xml:"unReadCount,omitempty"`
	// 未读消息列表
	UnReadItems []*QueryUnReadMessageResponseBodyUnReadItems `json:"unReadItems,omitempty" xml:"unReadItems,omitempty" type:"Repeated"`
}

func (s QueryUnReadMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUnReadMessageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUnReadMessageResponseBody) SetUnReadCount(v int64) *QueryUnReadMessageResponseBody {
	s.UnReadCount = &v
	return s
}

func (s *QueryUnReadMessageResponseBody) SetUnReadItems(v []*QueryUnReadMessageResponseBodyUnReadItems) *QueryUnReadMessageResponseBody {
	s.UnReadItems = v
	return s
}

type QueryUnReadMessageResponseBodyUnReadItems struct {
	// 群会话Id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 未读消息数
	UnReadCount *int64 `json:"unReadCount,omitempty" xml:"unReadCount,omitempty"`
}

func (s QueryUnReadMessageResponseBodyUnReadItems) String() string {
	return tea.Prettify(s)
}

func (s QueryUnReadMessageResponseBodyUnReadItems) GoString() string {
	return s.String()
}

func (s *QueryUnReadMessageResponseBodyUnReadItems) SetOpenConversationId(v string) *QueryUnReadMessageResponseBodyUnReadItems {
	s.OpenConversationId = &v
	return s
}

func (s *QueryUnReadMessageResponseBodyUnReadItems) SetUnReadCount(v int64) *QueryUnReadMessageResponseBodyUnReadItems {
	s.UnReadCount = &v
	return s
}

type QueryUnReadMessageResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUnReadMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUnReadMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUnReadMessageResponse) GoString() string {
	return s.String()
}

func (s *QueryUnReadMessageResponse) SetHeaders(v map[string]*string) *QueryUnReadMessageResponse {
	s.Headers = v
	return s
}

func (s *QueryUnReadMessageResponse) SetBody(v *QueryUnReadMessageResponseBody) *QueryUnReadMessageResponse {
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
	// 是否开启卡片纯拉模式
	PullStrategy *bool `json:"pullStrategy,omitempty" xml:"pullStrategy,omitempty"`
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

func (s *SendInteractiveCardRequest) SetPullStrategy(v bool) *SendInteractiveCardRequest {
	s.PullStrategy = &v
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
	// 可交互卡片回调的url【可空：不填写无需回调】
	CallbackUrl *string `json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
	// 唯一标识一张卡片的外部ID（卡片幂等ID，可用于更新或重复发送同一卡片到多个群会话）【备注：同一个outTrackId重复创建，卡片数据不覆盖更新】
	CardBizId *string `json:"cardBizId,omitempty" xml:"cardBizId,omitempty"`
	// 卡片模板-文本内容参数（卡片json结构体）
	CardData *string `json:"cardData,omitempty" xml:"cardData,omitempty"`
	// 卡片搭建平台模板ID
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	// 【openConversationId & singleChatReceiver 二选一必填】接收卡片的加密群ID，特指多人群会话（非单聊）
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 是否开启卡片纯拉模式
	PullStrategy *bool `json:"pullStrategy,omitempty" xml:"pullStrategy,omitempty"`
	// 机器人代码，群模板机器人网页有机器人ID；企业内部机器人为机器人appKey，企业三方机器人有robotCode
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	// 互动卡片发送选项
	SendOptions *SendRobotInteractiveCardRequestSendOptions `json:"sendOptions,omitempty" xml:"sendOptions,omitempty" type:"Struct"`
	// 【openConversationId & singleChatReceiver 二选一必填】单聊会话接受者json串
	SingleChatReceiver *string `json:"singleChatReceiver,omitempty" xml:"singleChatReceiver,omitempty"`
	// 卡片模板-userId差异用户参数（json结构体）
	UnionIdPrivateDataMap *string `json:"unionIdPrivateDataMap,omitempty" xml:"unionIdPrivateDataMap,omitempty"`
	// 卡片模板-userId差异用户参数（json结构体）
	UserIdPrivateDataMap *string `json:"userIdPrivateDataMap,omitempty" xml:"userIdPrivateDataMap,omitempty"`
}

func (s SendRobotInteractiveCardRequest) String() string {
	return tea.Prettify(s)
}

func (s SendRobotInteractiveCardRequest) GoString() string {
	return s.String()
}

func (s *SendRobotInteractiveCardRequest) SetCallbackUrl(v string) *SendRobotInteractiveCardRequest {
	s.CallbackUrl = &v
	return s
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

func (s *SendRobotInteractiveCardRequest) SetPullStrategy(v bool) *SendRobotInteractiveCardRequest {
	s.PullStrategy = &v
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

func (s *SendRobotInteractiveCardRequest) SetUnionIdPrivateDataMap(v string) *SendRobotInteractiveCardRequest {
	s.UnionIdPrivateDataMap = &v
	return s
}

func (s *SendRobotInteractiveCardRequest) SetUserIdPrivateDataMap(v string) *SendRobotInteractiveCardRequest {
	s.UserIdPrivateDataMap = &v
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
	// 发送的会话类型：单聊-0, 群聊-1
	ConversationType *int32 `json:"conversationType,omitempty" xml:"conversationType,omitempty"`
	// 酷应用编码
	CoolAppCode *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	// 接收卡片的群的openConversationId
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 唯一标识一张卡片的外部ID（卡片幂等ID，可用于更新或重复发送同一卡片到多个群会话）
	OutTrackId *string `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	// 接收人的员工号列表
	ReceiverUserIdList []*string `json:"receiverUserIdList,omitempty" xml:"receiverUserIdList,omitempty" type:"Repeated"`
	// 机器人编码
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s TopboxCloseRequest) String() string {
	return tea.Prettify(s)
}

func (s TopboxCloseRequest) GoString() string {
	return s.String()
}

func (s *TopboxCloseRequest) SetConversationType(v int32) *TopboxCloseRequest {
	s.ConversationType = &v
	return s
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

func (s *TopboxCloseRequest) SetReceiverUserIdList(v []*string) *TopboxCloseRequest {
	s.ReceiverUserIdList = v
	return s
}

func (s *TopboxCloseRequest) SetRobotCode(v string) *TopboxCloseRequest {
	s.RobotCode = &v
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
	// 发送的会话类型：单聊-0, 群聊-1
	ConversationType *int32 `json:"conversationType,omitempty" xml:"conversationType,omitempty"`
	// 酷应用编码
	CoolAppCode *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	// 吊顶的过期时间（绝对时间）
	ExpiredTime *int64 `json:"expiredTime,omitempty" xml:"expiredTime,omitempty"`
	// 接收卡片的群的openConversationId
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 唯一标识一张卡片的外部ID（卡片幂等ID，可用于更新或重复发送同一卡片到多个群会话）
	OutTrackId *string `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	// 期望吊顶的端（多个"|"隔开，如："ios|win|"）
	Platforms *string `json:"platforms,omitempty" xml:"platforms,omitempty"`
	// 接收人的员工号列表
	ReceiverUserIdList []*string `json:"receiverUserIdList,omitempty" xml:"receiverUserIdList,omitempty" type:"Repeated"`
	// 机器人编码
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s TopboxOpenRequest) String() string {
	return tea.Prettify(s)
}

func (s TopboxOpenRequest) GoString() string {
	return s.String()
}

func (s *TopboxOpenRequest) SetConversationType(v int32) *TopboxOpenRequest {
	s.ConversationType = &v
	return s
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

func (s *TopboxOpenRequest) SetReceiverUserIdList(v []*string) *TopboxOpenRequest {
	s.ReceiverUserIdList = v
	return s
}

func (s *TopboxOpenRequest) SetRobotCode(v string) *TopboxOpenRequest {
	s.RobotCode = &v
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

type UpdateGroupAvatarHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateGroupAvatarHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupAvatarHeaders) GoString() string {
	return s.String()
}

func (s *UpdateGroupAvatarHeaders) SetCommonHeaders(v map[string]*string) *UpdateGroupAvatarHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateGroupAvatarHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateGroupAvatarHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateGroupAvatarRequest struct {
	// 群头像地址
	GroupAvatar *string `json:"groupAvatar,omitempty" xml:"groupAvatar,omitempty"`
	// 群会话id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s UpdateGroupAvatarRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupAvatarRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupAvatarRequest) SetGroupAvatar(v string) *UpdateGroupAvatarRequest {
	s.GroupAvatar = &v
	return s
}

func (s *UpdateGroupAvatarRequest) SetOpenConversationId(v string) *UpdateGroupAvatarRequest {
	s.OpenConversationId = &v
	return s
}

type UpdateGroupAvatarResponseBody struct {
	// 新头像
	NewGroupAvatar *string `json:"newGroupAvatar,omitempty" xml:"newGroupAvatar,omitempty"`
}

func (s UpdateGroupAvatarResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupAvatarResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupAvatarResponseBody) SetNewGroupAvatar(v string) *UpdateGroupAvatarResponseBody {
	s.NewGroupAvatar = &v
	return s
}

type UpdateGroupAvatarResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateGroupAvatarResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGroupAvatarResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupAvatarResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupAvatarResponse) SetHeaders(v map[string]*string) *UpdateGroupAvatarResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupAvatarResponse) SetBody(v *UpdateGroupAvatarResponseBody) *UpdateGroupAvatarResponse {
	s.Body = v
	return s
}

type UpdateGroupNameHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
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

func (s *UpdateGroupNameHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateGroupNameHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateGroupNameRequest struct {
	// 群名称
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// 群会话id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s UpdateGroupNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupNameRequest) SetGroupName(v string) *UpdateGroupNameRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateGroupNameRequest) SetOpenConversationId(v string) *UpdateGroupNameRequest {
	s.OpenConversationId = &v
	return s
}

type UpdateGroupNameResponseBody struct {
	// Id of the request
	NewGroupName *string `json:"newGroupName,omitempty" xml:"newGroupName,omitempty"`
}

func (s UpdateGroupNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupNameResponseBody) SetNewGroupName(v string) *UpdateGroupNameResponseBody {
	s.NewGroupName = &v
	return s
}

type UpdateGroupNameResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateGroupNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateGroupNameResponse) SetBody(v *UpdateGroupNameResponseBody) *UpdateGroupNameResponse {
	s.Body = v
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

type UpdateRobotInteractiveCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateRobotInteractiveCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateRobotInteractiveCardHeaders) GoString() string {
	return s.String()
}

func (s *UpdateRobotInteractiveCardHeaders) SetCommonHeaders(v map[string]*string) *UpdateRobotInteractiveCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateRobotInteractiveCardHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateRobotInteractiveCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateRobotInteractiveCardRequest struct {
	// 唯一标识一张卡片的外部ID（卡片幂等ID，可用于更新或重复发送同一卡片到多个群会话）【备注：同一个outTrackId重复创建，卡片数据不覆盖更新】
	CardBizId *string `json:"cardBizId,omitempty" xml:"cardBizId,omitempty"`
	// 卡片模板-文本内容参数（卡片json结构体）
	CardData *string `json:"cardData,omitempty" xml:"cardData,omitempty"`
	// 卡片模板-userId差异用户参数（json结构体）
	UnionIdPrivateDataMap *string `json:"unionIdPrivateDataMap,omitempty" xml:"unionIdPrivateDataMap,omitempty"`
	// 互动卡片更新选项
	UpdateOptions *UpdateRobotInteractiveCardRequestUpdateOptions `json:"updateOptions,omitempty" xml:"updateOptions,omitempty" type:"Struct"`
	// 卡片模板-userId差异用户参数（json结构体）
	UserIdPrivateDataMap *string `json:"userIdPrivateDataMap,omitempty" xml:"userIdPrivateDataMap,omitempty"`
}

func (s UpdateRobotInteractiveCardRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRobotInteractiveCardRequest) GoString() string {
	return s.String()
}

func (s *UpdateRobotInteractiveCardRequest) SetCardBizId(v string) *UpdateRobotInteractiveCardRequest {
	s.CardBizId = &v
	return s
}

func (s *UpdateRobotInteractiveCardRequest) SetCardData(v string) *UpdateRobotInteractiveCardRequest {
	s.CardData = &v
	return s
}

func (s *UpdateRobotInteractiveCardRequest) SetUnionIdPrivateDataMap(v string) *UpdateRobotInteractiveCardRequest {
	s.UnionIdPrivateDataMap = &v
	return s
}

func (s *UpdateRobotInteractiveCardRequest) SetUpdateOptions(v *UpdateRobotInteractiveCardRequestUpdateOptions) *UpdateRobotInteractiveCardRequest {
	s.UpdateOptions = v
	return s
}

func (s *UpdateRobotInteractiveCardRequest) SetUserIdPrivateDataMap(v string) *UpdateRobotInteractiveCardRequest {
	s.UserIdPrivateDataMap = &v
	return s
}

type UpdateRobotInteractiveCardRequestUpdateOptions struct {
	// 按key更新数据(默认全量更新)
	UpdateCardDataByKey *bool `json:"updateCardDataByKey,omitempty" xml:"updateCardDataByKey,omitempty"`
	// 按key更新用户数据(默认全量更新)
	UpdatePrivateDataByKey *bool `json:"updatePrivateDataByKey,omitempty" xml:"updatePrivateDataByKey,omitempty"`
}

func (s UpdateRobotInteractiveCardRequestUpdateOptions) String() string {
	return tea.Prettify(s)
}

func (s UpdateRobotInteractiveCardRequestUpdateOptions) GoString() string {
	return s.String()
}

func (s *UpdateRobotInteractiveCardRequestUpdateOptions) SetUpdateCardDataByKey(v bool) *UpdateRobotInteractiveCardRequestUpdateOptions {
	s.UpdateCardDataByKey = &v
	return s
}

func (s *UpdateRobotInteractiveCardRequestUpdateOptions) SetUpdatePrivateDataByKey(v bool) *UpdateRobotInteractiveCardRequestUpdateOptions {
	s.UpdatePrivateDataByKey = &v
	return s
}

type UpdateRobotInteractiveCardResponseBody struct {
	// 用于业务方后续查看已读列表的查询key
	ProcessQueryKey *string `json:"processQueryKey,omitempty" xml:"processQueryKey,omitempty"`
}

func (s UpdateRobotInteractiveCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRobotInteractiveCardResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRobotInteractiveCardResponseBody) SetProcessQueryKey(v string) *UpdateRobotInteractiveCardResponseBody {
	s.ProcessQueryKey = &v
	return s
}

type UpdateRobotInteractiveCardResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateRobotInteractiveCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRobotInteractiveCardResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRobotInteractiveCardResponse) GoString() string {
	return s.String()
}

func (s *UpdateRobotInteractiveCardResponse) SetHeaders(v map[string]*string) *UpdateRobotInteractiveCardResponse {
	s.Headers = v
	return s
}

func (s *UpdateRobotInteractiveCardResponse) SetBody(v *UpdateRobotInteractiveCardResponseBody) *UpdateRobotInteractiveCardResponse {
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

type AddGroupMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddGroupMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMemberHeaders) GoString() string {
	return s.String()
}

func (s *AddGroupMemberHeaders) SetCommonHeaders(v map[string]*string) *AddGroupMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddGroupMemberHeaders) SetXAcsDingtalkAccessToken(v string) *AddGroupMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddGroupMemberRequest struct {
	// C端客户成员列表
	AppUserIds []*string `json:"appUserIds,omitempty" xml:"appUserIds,omitempty" type:"Repeated"`
	// 群会话Id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 操作者钉钉Id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// B端客服成员列表
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s AddGroupMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *AddGroupMemberRequest) SetAppUserIds(v []*string) *AddGroupMemberRequest {
	s.AppUserIds = v
	return s
}

func (s *AddGroupMemberRequest) SetOpenConversationId(v string) *AddGroupMemberRequest {
	s.OpenConversationId = &v
	return s
}

func (s *AddGroupMemberRequest) SetOperatorId(v string) *AddGroupMemberRequest {
	s.OperatorId = &v
	return s
}

func (s *AddGroupMemberRequest) SetUserIds(v []*string) *AddGroupMemberRequest {
	s.UserIds = v
	return s
}

type AddGroupMemberResponseBody struct {
	// 拉取成功的C端客户列表
	AppUserIds []*string `json:"appUserIds,omitempty" xml:"appUserIds,omitempty" type:"Repeated"`
	// 拉取成功的B端客服列表
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s AddGroupMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddGroupMemberResponseBody) SetAppUserIds(v []*string) *AddGroupMemberResponseBody {
	s.AppUserIds = v
	return s
}

func (s *AddGroupMemberResponseBody) SetUserIds(v []*string) *AddGroupMemberResponseBody {
	s.UserIds = v
	return s
}

type AddGroupMemberResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddGroupMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *AddGroupMemberResponse) SetHeaders(v map[string]*string) *AddGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *AddGroupMemberResponse) SetBody(v *AddGroupMemberResponseBody) *AddGroupMemberResponse {
	s.Body = v
	return s
}

type RemoveGroupMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RemoveGroupMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupMemberHeaders) GoString() string {
	return s.String()
}

func (s *RemoveGroupMemberHeaders) SetCommonHeaders(v map[string]*string) *RemoveGroupMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveGroupMemberHeaders) SetXAcsDingtalkAccessToken(v string) *RemoveGroupMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RemoveGroupMemberRequest struct {
	// C端客户成员列表
	AppUserIds []*string `json:"appUserIds,omitempty" xml:"appUserIds,omitempty" type:"Repeated"`
	// 群会话Id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 操作者钉钉Id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// B端客服成员列表
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s RemoveGroupMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *RemoveGroupMemberRequest) SetAppUserIds(v []*string) *RemoveGroupMemberRequest {
	s.AppUserIds = v
	return s
}

func (s *RemoveGroupMemberRequest) SetOpenConversationId(v string) *RemoveGroupMemberRequest {
	s.OpenConversationId = &v
	return s
}

func (s *RemoveGroupMemberRequest) SetOperatorId(v string) *RemoveGroupMemberRequest {
	s.OperatorId = &v
	return s
}

func (s *RemoveGroupMemberRequest) SetUserIds(v []*string) *RemoveGroupMemberRequest {
	s.UserIds = v
	return s
}

type RemoveGroupMemberResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s RemoveGroupMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *RemoveGroupMemberResponse) SetHeaders(v map[string]*string) *RemoveGroupMemberResponse {
	s.Headers = v
	return s
}

type SendDingMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendDingMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendDingMessageHeaders) GoString() string {
	return s.String()
}

func (s *SendDingMessageHeaders) SetCommonHeaders(v map[string]*string) *SendDingMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendDingMessageHeaders) SetXAcsDingtalkAccessToken(v string) *SendDingMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendDingMessageRequest struct {
	// 客服oauth2.0授权码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 消息内容
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 消息类型
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// 群会话Id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// C端客户appUserId
	ReceiverId *string `json:"receiverId,omitempty" xml:"receiverId,omitempty"`
	// B端客服钉钉Id
	SenderId *string `json:"senderId,omitempty" xml:"senderId,omitempty"`
}

func (s SendDingMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendDingMessageRequest) GoString() string {
	return s.String()
}

func (s *SendDingMessageRequest) SetCode(v string) *SendDingMessageRequest {
	s.Code = &v
	return s
}

func (s *SendDingMessageRequest) SetMessage(v string) *SendDingMessageRequest {
	s.Message = &v
	return s
}

func (s *SendDingMessageRequest) SetMessageType(v string) *SendDingMessageRequest {
	s.MessageType = &v
	return s
}

func (s *SendDingMessageRequest) SetOpenConversationId(v string) *SendDingMessageRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SendDingMessageRequest) SetReceiverId(v string) *SendDingMessageRequest {
	s.ReceiverId = &v
	return s
}

func (s *SendDingMessageRequest) SetSenderId(v string) *SendDingMessageRequest {
	s.SenderId = &v
	return s
}

type SendDingMessageResponseBody struct {
	// 发送消息请求Id
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SendDingMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendDingMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendDingMessageResponseBody) SetRequestId(v string) *SendDingMessageResponseBody {
	s.RequestId = &v
	return s
}

type SendDingMessageResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendDingMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendDingMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendDingMessageResponse) GoString() string {
	return s.String()
}

func (s *SendDingMessageResponse) SetHeaders(v map[string]*string) *SendDingMessageResponse {
	s.Headers = v
	return s
}

func (s *SendDingMessageResponse) SetBody(v *SendDingMessageResponseBody) *SendDingMessageResponse {
	s.Body = v
	return s
}

type SendMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
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

func (s *SendMessageHeaders) SetXAcsDingtalkAccessToken(v string) *SendMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendMessageRequest struct {
	// 消息内容
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 消息类型
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// 群会话Id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// B端客服钉钉Id
	ReceiverId *string `json:"receiverId,omitempty" xml:"receiverId,omitempty"`
	// C端客户appUserId
	SenderId *string `json:"senderId,omitempty" xml:"senderId,omitempty"`
}

func (s SendMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequest) SetMessage(v string) *SendMessageRequest {
	s.Message = &v
	return s
}

func (s *SendMessageRequest) SetMessageType(v string) *SendMessageRequest {
	s.MessageType = &v
	return s
}

func (s *SendMessageRequest) SetOpenConversationId(v string) *SendMessageRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SendMessageRequest) SetReceiverId(v string) *SendMessageRequest {
	s.ReceiverId = &v
	return s
}

func (s *SendMessageRequest) SetSenderId(v string) *SendMessageRequest {
	s.SenderId = &v
	return s
}

type SendMessageResponseBody struct {
	// 发送消息请求Id
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SendMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageResponseBody) SetRequestId(v string) *SendMessageResponseBody {
	s.RequestId = &v
	return s
}

type SendMessageResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SendMessageResponse) SetBody(v *SendMessageResponseBody) *SendMessageResponse {
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

func (client *Client) AutoOpenDingTalkConnect() (_result *AutoOpenDingTalkConnectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AutoOpenDingTalkConnectHeaders{}
	_result = &AutoOpenDingTalkConnectResponse{}
	_body, _err := client.AutoOpenDingTalkConnectWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AutoOpenDingTalkConnectWithOptions(headers *AutoOpenDingTalkConnectHeaders, runtime *util.RuntimeOptions) (_result *AutoOpenDingTalkConnectResponse, _err error) {
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
	_result = &AutoOpenDingTalkConnectResponse{}
	_body, _err := client.DoROARequest(tea.String("AutoOpenDingTalkConnect"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/interconnections/apps/open"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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

func (client *Client) ChatSubAdminUpdate(request *ChatSubAdminUpdateRequest) (_result *ChatSubAdminUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ChatSubAdminUpdateHeaders{}
	_result = &ChatSubAdminUpdateResponse{}
	_body, _err := client.ChatSubAdminUpdateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChatSubAdminUpdateWithOptions(request *ChatSubAdminUpdateRequest, headers *ChatSubAdminUpdateHeaders, runtime *util.RuntimeOptions) (_result *ChatSubAdminUpdateResponse, _err error) {
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
	_result = &ChatSubAdminUpdateResponse{}
	_body, _err := client.DoROARequest(tea.String("ChatSubAdminUpdate"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/subAdministrators"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCoupleGroupConversation(request *CreateCoupleGroupConversationRequest) (_result *CreateCoupleGroupConversationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateCoupleGroupConversationHeaders{}
	_result = &CreateCoupleGroupConversationResponse{}
	_body, _err := client.CreateCoupleGroupConversationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCoupleGroupConversationWithOptions(request *CreateCoupleGroupConversationRequest, headers *CreateCoupleGroupConversationHeaders, runtime *util.RuntimeOptions) (_result *CreateCoupleGroupConversationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUserId)) {
		body["appUserId"] = request.AppUserId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupAvatar)) {
		body["groupAvatar"] = request.GroupAvatar
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupOwnerId)) {
		body["groupOwnerId"] = request.GroupOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupTemplateId)) {
		body["groupTemplateId"] = request.GroupTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
	_result = &CreateCoupleGroupConversationResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateCoupleGroupConversation"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/interconnections/coupleGroups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGroupConversation(request *CreateGroupConversationRequest) (_result *CreateGroupConversationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateGroupConversationHeaders{}
	_result = &CreateGroupConversationResponse{}
	_body, _err := client.CreateGroupConversationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGroupConversationWithOptions(request *CreateGroupConversationRequest, headers *CreateGroupConversationHeaders, runtime *util.RuntimeOptions) (_result *CreateGroupConversationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUserIds)) {
		body["appUserIds"] = request.AppUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.GroupAvatar)) {
		body["groupAvatar"] = request.GroupAvatar
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupOwnerId)) {
		body["groupOwnerId"] = request.GroupOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupTemplateId)) {
		body["groupTemplateId"] = request.GroupTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
	_result = &CreateGroupConversationResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateGroupConversation"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/interconnections/groups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInterconnection(request *CreateInterconnectionRequest) (_result *CreateInterconnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateInterconnectionHeaders{}
	_result = &CreateInterconnectionResponse{}
	_body, _err := client.CreateInterconnectionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInterconnectionWithOptions(request *CreateInterconnectionRequest, headers *CreateInterconnectionHeaders, runtime *util.RuntimeOptions) (_result *CreateInterconnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Interconnections)) {
		body["interconnections"] = request.Interconnections
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		body["signature"] = request.Signature
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
	_result = &CreateInterconnectionResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateInterconnection"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/interconnections"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateStoreGroupConversation(request *CreateStoreGroupConversationRequest) (_result *CreateStoreGroupConversationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateStoreGroupConversationHeaders{}
	_result = &CreateStoreGroupConversationResponse{}
	_body, _err := client.CreateStoreGroupConversationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateStoreGroupConversationWithOptions(request *CreateStoreGroupConversationRequest, headers *CreateStoreGroupConversationHeaders, runtime *util.RuntimeOptions) (_result *CreateStoreGroupConversationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUserId)) {
		body["appUserId"] = request.AppUserId
	}

	if !tea.BoolValue(util.IsUnset(request.BusinessUniqueKey)) {
		body["businessUniqueKey"] = request.BusinessUniqueKey
	}

	if !tea.BoolValue(util.IsUnset(request.GroupAvatar)) {
		body["groupAvatar"] = request.GroupAvatar
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupTemplateId)) {
		body["groupTemplateId"] = request.GroupTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
	_result = &CreateStoreGroupConversationResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateStoreGroupConversation"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/interconnections/storeGroups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConversationUrl(request *GetConversationUrlRequest) (_result *GetConversationUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetConversationUrlHeaders{}
	_result = &GetConversationUrlResponse{}
	_body, _err := client.GetConversationUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConversationUrlWithOptions(request *GetConversationUrlRequest, headers *GetConversationUrlHeaders, runtime *util.RuntimeOptions) (_result *GetConversationUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUserId)) {
		body["appUserId"] = request.AppUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelCode)) {
		body["channelCode"] = request.ChannelCode
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCode)) {
		body["sourceCode"] = request.SourceCode
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
	_result = &GetConversationUrlResponse{}
	_body, _err := client.DoROARequest(tea.String("GetConversationUrl"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/conversations/urls"), tea.String("json"), req, runtime)
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

func (client *Client) GroupBanWords(request *GroupBanWordsRequest) (_result *GroupBanWordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GroupBanWordsHeaders{}
	_result = &GroupBanWordsResponse{}
	_body, _err := client.GroupBanWordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GroupBanWordsWithOptions(request *GroupBanWordsRequest, headers *GroupBanWordsHeaders, runtime *util.RuntimeOptions) (_result *GroupBanWordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BanWordsMode)) {
		body["banWordsMode"] = request.BanWordsMode
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		body["options"] = request.Options
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
	_result = &GroupBanWordsResponse{}
	_body, _err := client.DoROARequest(tea.String("GroupBanWords"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/groups/words/ban"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GroupCapacityInquiry(request *GroupCapacityInquiryRequest) (_result *GroupCapacityInquiryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GroupCapacityInquiryHeaders{}
	_result = &GroupCapacityInquiryResponse{}
	_body, _err := client.GroupCapacityInquiryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GroupCapacityInquiryWithOptions(request *GroupCapacityInquiryRequest, headers *GroupCapacityInquiryHeaders, runtime *util.RuntimeOptions) (_result *GroupCapacityInquiryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EffectiveDuration)) {
		body["effectiveDuration"] = request.EffectiveDuration
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		body["options"] = request.Options
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCapacity)) {
		body["targetCapacity"] = request.TargetCapacity
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
	_result = &GroupCapacityInquiryResponse{}
	_body, _err := client.DoROARequest(tea.String("GroupCapacityInquiry"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/groups/capacities/inquiries/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GroupCapacityOrderConfirm(request *GroupCapacityOrderConfirmRequest) (_result *GroupCapacityOrderConfirmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GroupCapacityOrderConfirmHeaders{}
	_result = &GroupCapacityOrderConfirmResponse{}
	_body, _err := client.GroupCapacityOrderConfirmWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GroupCapacityOrderConfirmWithOptions(request *GroupCapacityOrderConfirmRequest, headers *GroupCapacityOrderConfirmHeaders, runtime *util.RuntimeOptions) (_result *GroupCapacityOrderConfirmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		body["orderId"] = request.OrderId
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
	_result = &GroupCapacityOrderConfirmResponse{}
	_body, _err := client.DoROARequest(tea.String("GroupCapacityOrderConfirm"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/groups/capacities/orders/confirm"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GroupCapacityOrderPlace(request *GroupCapacityOrderPlaceRequest) (_result *GroupCapacityOrderPlaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GroupCapacityOrderPlaceHeaders{}
	_result = &GroupCapacityOrderPlaceResponse{}
	_body, _err := client.GroupCapacityOrderPlaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GroupCapacityOrderPlaceWithOptions(request *GroupCapacityOrderPlaceRequest, headers *GroupCapacityOrderPlaceHeaders, runtime *util.RuntimeOptions) (_result *GroupCapacityOrderPlaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActualPrice)) {
		body["actualPrice"] = request.ActualPrice
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentCapacity)) {
		body["currentCapacity"] = request.CurrentCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentEffectUntil)) {
		body["currentEffectUntil"] = request.CurrentEffectUntil
	}

	if !tea.BoolValue(util.IsUnset(request.Discount)) {
		body["discount"] = request.Discount
	}

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["extInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.MarkedPrice)) {
		body["markedPrice"] = request.MarkedPrice
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCapacity)) {
		body["targetCapacity"] = request.TargetCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.TargetEffectUntil)) {
		body["targetEffectUntil"] = request.TargetEffectUntil
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		body["token"] = request.Token
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
	_result = &GroupCapacityOrderPlaceResponse{}
	_body, _err := client.DoROARequest(tea.String("GroupCapacityOrderPlace"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/groups/capacities/orders/place"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GroupManageQuery(request *GroupManageQueryRequest) (_result *GroupManageQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GroupManageQueryHeaders{}
	_result = &GroupManageQueryResponse{}
	_body, _err := client.GroupManageQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GroupManageQueryWithOptions(request *GroupManageQueryRequest, headers *GroupManageQueryHeaders, runtime *util.RuntimeOptions) (_result *GroupManageQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreatedAfter)) {
		body["createdAfter"] = request.CreatedAfter
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["groupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupMemberSamples)) {
		body["groupMemberSamples"] = request.GroupMemberSamples
	}

	if !tea.BoolValue(util.IsUnset(request.GroupOwner)) {
		body["groupOwner"] = request.GroupOwner
	}

	if !tea.BoolValue(util.IsUnset(request.GroupTitleKeywords)) {
		body["groupTitleKeywords"] = request.GroupTitleKeywords
	}

	if !tea.BoolValue(util.IsUnset(request.GroupUrl)) {
		body["groupUrl"] = request.GroupUrl
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.MembersOver)) {
		body["membersOver"] = request.MembersOver
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
	_result = &GroupManageQueryResponse{}
	_body, _err := client.DoROARequest(tea.String("GroupManageQuery"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/groups/managements/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GroupManageReduce(request *GroupManageReduceRequest) (_result *GroupManageReduceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GroupManageReduceHeaders{}
	_result = &GroupManageReduceResponse{}
	_body, _err := client.GroupManageReduceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GroupManageReduceWithOptions(request *GroupManageReduceRequest, headers *GroupManageReduceHeaders, runtime *util.RuntimeOptions) (_result *GroupManageReduceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CapacityLimit)) {
		body["capacityLimit"] = request.CapacityLimit
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		body["options"] = request.Options
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
	_result = &GroupManageReduceResponse{}
	_body, _err := client.DoROARequest(tea.String("GroupManageReduce"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/groups/capacities/reduce"), tea.String("none"), req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.PullStrategy)) {
		body["pullStrategy"] = request.PullStrategy
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

func (client *Client) QueryGroupInfoByMemberAuth(request *QueryGroupInfoByMemberAuthRequest) (_result *QueryGroupInfoByMemberAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGroupInfoByMemberAuthHeaders{}
	_result = &QueryGroupInfoByMemberAuthResponse{}
	_body, _err := client.QueryGroupInfoByMemberAuthWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryGroupInfoByMemberAuthWithOptions(request *QueryGroupInfoByMemberAuthRequest, headers *QueryGroupInfoByMemberAuthHeaders, runtime *util.RuntimeOptions) (_result *QueryGroupInfoByMemberAuthResponse, _err error) {
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
	_result = &QueryGroupInfoByMemberAuthResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryGroupInfoByMemberAuth"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/memberAuthorizations/groups/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryGroupMember(request *QueryGroupMemberRequest) (_result *QueryGroupMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGroupMemberHeaders{}
	_result = &QueryGroupMemberResponse{}
	_body, _err := client.QueryGroupMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryGroupMemberWithOptions(request *QueryGroupMemberRequest, headers *QueryGroupMemberHeaders, runtime *util.RuntimeOptions) (_result *QueryGroupMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		query["openConversationId"] = request.OpenConversationId
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
	_result = &QueryGroupMemberResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryGroupMember"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/im/interconnections/conversations/members"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryGroupMemberByMemberAuth(request *QueryGroupMemberByMemberAuthRequest) (_result *QueryGroupMemberByMemberAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGroupMemberByMemberAuthHeaders{}
	_result = &QueryGroupMemberByMemberAuthResponse{}
	_body, _err := client.QueryGroupMemberByMemberAuthWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryGroupMemberByMemberAuthWithOptions(request *QueryGroupMemberByMemberAuthRequest, headers *QueryGroupMemberByMemberAuthHeaders, runtime *util.RuntimeOptions) (_result *QueryGroupMemberByMemberAuthResponse, _err error) {
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
	_result = &QueryGroupMemberByMemberAuthResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryGroupMemberByMemberAuth"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/memberAuthorizations/groups/members/query"), tea.String("json"), req, runtime)
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

func (client *Client) QuerySingleGroup(request *QuerySingleGroupRequest) (_result *QuerySingleGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QuerySingleGroupHeaders{}
	_result = &QuerySingleGroupResponse{}
	_body, _err := client.QuerySingleGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySingleGroupWithOptions(request *QuerySingleGroupRequest, headers *QuerySingleGroupHeaders, runtime *util.RuntimeOptions) (_result *QuerySingleGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupMembers)) {
		body["groupMembers"] = request.GroupMembers
	}

	if !tea.BoolValue(util.IsUnset(request.GroupTemplateId)) {
		body["groupTemplateId"] = request.GroupTemplateId
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
	_result = &QuerySingleGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("QuerySingleGroup"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/interconnections/doubleGroups/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUnReadMessage(request *QueryUnReadMessageRequest) (_result *QueryUnReadMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUnReadMessageHeaders{}
	_result = &QueryUnReadMessageResponse{}
	_body, _err := client.QueryUnReadMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUnReadMessageWithOptions(request *QueryUnReadMessageRequest, headers *QueryUnReadMessageHeaders, runtime *util.RuntimeOptions) (_result *QueryUnReadMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUserId)) {
		body["appUserId"] = request.AppUserId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationIds)) {
		body["openConversationIds"] = request.OpenConversationIds
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
	_result = &QueryUnReadMessageResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryUnReadMessage"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/interconnections/unReadMsgs/query"), tea.String("json"), req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.PullStrategy)) {
		body["pullStrategy"] = request.PullStrategy
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
	if !tea.BoolValue(util.IsUnset(request.CallbackUrl)) {
		body["callbackUrl"] = request.CallbackUrl
	}

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

	if !tea.BoolValue(util.IsUnset(request.PullStrategy)) {
		body["pullStrategy"] = request.PullStrategy
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

	if !tea.BoolValue(util.IsUnset(request.UnionIdPrivateDataMap)) {
		body["unionIdPrivateDataMap"] = request.UnionIdPrivateDataMap
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdPrivateDataMap)) {
		body["userIdPrivateDataMap"] = request.UserIdPrivateDataMap
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
	if !tea.BoolValue(util.IsUnset(request.ConversationType)) {
		body["conversationType"] = request.ConversationType
	}

	if !tea.BoolValue(util.IsUnset(request.CoolAppCode)) {
		body["coolAppCode"] = request.CoolAppCode
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OutTrackId)) {
		body["outTrackId"] = request.OutTrackId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUserIdList)) {
		body["receiverUserIdList"] = request.ReceiverUserIdList
	}

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
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
	if !tea.BoolValue(util.IsUnset(request.ConversationType)) {
		body["conversationType"] = request.ConversationType
	}

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

	if !tea.BoolValue(util.IsUnset(request.ReceiverUserIdList)) {
		body["receiverUserIdList"] = request.ReceiverUserIdList
	}

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
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

func (client *Client) UpdateGroupAvatar(request *UpdateGroupAvatarRequest) (_result *UpdateGroupAvatarResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateGroupAvatarHeaders{}
	_result = &UpdateGroupAvatarResponse{}
	_body, _err := client.UpdateGroupAvatarWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGroupAvatarWithOptions(request *UpdateGroupAvatarRequest, headers *UpdateGroupAvatarHeaders, runtime *util.RuntimeOptions) (_result *UpdateGroupAvatarResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupAvatar)) {
		body["groupAvatar"] = request.GroupAvatar
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
	_result = &UpdateGroupAvatarResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateGroupAvatar"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/im/interconnections/groups/avatars"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) UpdateGroupNameWithOptions(request *UpdateGroupNameRequest, headers *UpdateGroupNameHeaders, runtime *util.RuntimeOptions) (_result *UpdateGroupNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
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
	_result = &UpdateGroupNameResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateGroupName"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/im/interconnections/groups/names"), tea.String("json"), req, runtime)
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

func (client *Client) UpdateRobotInteractiveCard(request *UpdateRobotInteractiveCardRequest) (_result *UpdateRobotInteractiveCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateRobotInteractiveCardHeaders{}
	_result = &UpdateRobotInteractiveCardResponse{}
	_body, _err := client.UpdateRobotInteractiveCardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRobotInteractiveCardWithOptions(request *UpdateRobotInteractiveCardRequest, headers *UpdateRobotInteractiveCardHeaders, runtime *util.RuntimeOptions) (_result *UpdateRobotInteractiveCardResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.UnionIdPrivateDataMap)) {
		body["unionIdPrivateDataMap"] = request.UnionIdPrivateDataMap
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.UpdateOptions))) {
		body["updateOptions"] = request.UpdateOptions
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdPrivateDataMap)) {
		body["userIdPrivateDataMap"] = request.UserIdPrivateDataMap
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
	_result = &UpdateRobotInteractiveCardResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateRobotInteractiveCard"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/im/robots/interactiveCards"), tea.String("json"), req, runtime)
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

func (client *Client) AddGroupMember(request *AddGroupMemberRequest) (_result *AddGroupMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddGroupMemberHeaders{}
	_result = &AddGroupMemberResponse{}
	_body, _err := client.AddGroupMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddGroupMemberWithOptions(request *AddGroupMemberRequest, headers *AddGroupMemberHeaders, runtime *util.RuntimeOptions) (_result *AddGroupMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUserIds)) {
		body["appUserIds"] = request.AppUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
	_result = &AddGroupMemberResponse{}
	_body, _err := client.DoROARequest(tea.String("addGroupMember"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/interconnections/groups/members"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveGroupMember(request *RemoveGroupMemberRequest) (_result *RemoveGroupMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveGroupMemberHeaders{}
	_result = &RemoveGroupMemberResponse{}
	_body, _err := client.RemoveGroupMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveGroupMemberWithOptions(request *RemoveGroupMemberRequest, headers *RemoveGroupMemberHeaders, runtime *util.RuntimeOptions) (_result *RemoveGroupMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUserIds)) {
		body["appUserIds"] = request.AppUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
	_result = &RemoveGroupMemberResponse{}
	_body, _err := client.DoROARequest(tea.String("removeGroupMember"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/interconnections/groups/members/remove"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendDingMessage(request *SendDingMessageRequest) (_result *SendDingMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendDingMessageHeaders{}
	_result = &SendDingMessageResponse{}
	_body, _err := client.SendDingMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendDingMessageWithOptions(request *SendDingMessageRequest, headers *SendDingMessageHeaders, runtime *util.RuntimeOptions) (_result *SendDingMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		body["message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		body["messageType"] = request.MessageType
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverId)) {
		body["receiverId"] = request.ReceiverId
	}

	if !tea.BoolValue(util.IsUnset(request.SenderId)) {
		body["senderId"] = request.SenderId
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
	_result = &SendDingMessageResponse{}
	_body, _err := client.DoROARequest(tea.String("sendDingMessage"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/interconnections/dingMessages/send"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) SendMessageWithOptions(request *SendMessageRequest, headers *SendMessageHeaders, runtime *util.RuntimeOptions) (_result *SendMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Message)) {
		body["message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		body["messageType"] = request.MessageType
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverId)) {
		body["receiverId"] = request.ReceiverId
	}

	if !tea.BoolValue(util.IsUnset(request.SenderId)) {
		body["senderId"] = request.SenderId
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
	_result = &SendMessageResponse{}
	_body, _err := client.DoROARequest(tea.String("sendMessage"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/interconnections/messages/send"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
