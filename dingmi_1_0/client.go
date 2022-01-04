// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package dingmi_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddRobotInstanceToGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddRobotInstanceToGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddRobotInstanceToGroupHeaders) GoString() string {
	return s.String()
}

func (s *AddRobotInstanceToGroupHeaders) SetCommonHeaders(v map[string]*string) *AddRobotInstanceToGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddRobotInstanceToGroupHeaders) SetXAcsDingtalkAccessToken(v string) *AddRobotInstanceToGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddRobotInstanceToGroupRequest struct {
	// 机器人id
	ChatbotId *string `json:"chatbotId,omitempty" xml:"chatbotId,omitempty"`
	// 对话id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s AddRobotInstanceToGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRobotInstanceToGroupRequest) GoString() string {
	return s.String()
}

func (s *AddRobotInstanceToGroupRequest) SetChatbotId(v string) *AddRobotInstanceToGroupRequest {
	s.ChatbotId = &v
	return s
}

func (s *AddRobotInstanceToGroupRequest) SetOpenConversationId(v string) *AddRobotInstanceToGroupRequest {
	s.OpenConversationId = &v
	return s
}

type AddRobotInstanceToGroupResponseBody struct {
	// 是否成功拉入群
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s AddRobotInstanceToGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRobotInstanceToGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddRobotInstanceToGroupResponseBody) SetResult(v bool) *AddRobotInstanceToGroupResponseBody {
	s.Result = &v
	return s
}

type AddRobotInstanceToGroupResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddRobotInstanceToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddRobotInstanceToGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddRobotInstanceToGroupResponse) GoString() string {
	return s.String()
}

func (s *AddRobotInstanceToGroupResponse) SetHeaders(v map[string]*string) *AddRobotInstanceToGroupResponse {
	s.Headers = v
	return s
}

func (s *AddRobotInstanceToGroupResponse) SetBody(v *AddRobotInstanceToGroupResponseBody) *AddRobotInstanceToGroupResponse {
	s.Body = v
	return s
}

type AskRobotHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AskRobotHeaders) String() string {
	return tea.Prettify(s)
}

func (s AskRobotHeaders) GoString() string {
	return s.String()
}

func (s *AskRobotHeaders) SetCommonHeaders(v map[string]*string) *AskRobotHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AskRobotHeaders) SetXAcsDingtalkAccessToken(v string) *AskRobotHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AskRobotRequest struct {
	DingUserId *string `json:"dingUserId,omitempty" xml:"dingUserId,omitempty"`
	// 问题
	Question *string `json:"question,omitempty" xml:"question,omitempty"`
	// 机器人id
	RobotAppKey *string `json:"robotAppKey,omitempty" xml:"robotAppKey,omitempty"`
	// sessionId(非必传)
	SessionUuid *string `json:"sessionUuid,omitempty" xml:"sessionUuid,omitempty"`
}

func (s AskRobotRequest) String() string {
	return tea.Prettify(s)
}

func (s AskRobotRequest) GoString() string {
	return s.String()
}

func (s *AskRobotRequest) SetDingUserId(v string) *AskRobotRequest {
	s.DingUserId = &v
	return s
}

func (s *AskRobotRequest) SetQuestion(v string) *AskRobotRequest {
	s.Question = &v
	return s
}

func (s *AskRobotRequest) SetRobotAppKey(v string) *AskRobotRequest {
	s.RobotAppKey = &v
	return s
}

func (s *AskRobotRequest) SetSessionUuid(v string) *AskRobotRequest {
	s.SessionUuid = &v
	return s
}

type AskRobotResponseBody struct {
	// 答案的json string
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s AskRobotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AskRobotResponseBody) GoString() string {
	return s.String()
}

func (s *AskRobotResponseBody) SetResult(v string) *AskRobotResponseBody {
	s.Result = &v
	return s
}

type AskRobotResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AskRobotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AskRobotResponse) String() string {
	return tea.Prettify(s)
}

func (s AskRobotResponse) GoString() string {
	return s.String()
}

func (s *AskRobotResponse) SetHeaders(v map[string]*string) *AskRobotResponse {
	s.Headers = v
	return s
}

func (s *AskRobotResponse) SetBody(v *AskRobotResponseBody) *AskRobotResponse {
	s.Body = v
	return s
}

type GetDingMeBaseDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDingMeBaseDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDingMeBaseDataHeaders) GoString() string {
	return s.String()
}

func (s *GetDingMeBaseDataHeaders) SetCommonHeaders(v map[string]*string) *GetDingMeBaseDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDingMeBaseDataHeaders) SetXAcsDingtalkAccessToken(v string) *GetDingMeBaseDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDingMeBaseDataRequest struct {
	// 机器人ID
	AppKey *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	// 是否按天分组
	ByDay *bool `json:"byDay,omitempty" xml:"byDay,omitempty"`
	// 结束时间
	EndDay *string `json:"endDay,omitempty" xml:"endDay,omitempty"`
	// 开始时间
	StartDay *string `json:"startDay,omitempty" xml:"startDay,omitempty"`
}

func (s GetDingMeBaseDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDingMeBaseDataRequest) GoString() string {
	return s.String()
}

func (s *GetDingMeBaseDataRequest) SetAppKey(v string) *GetDingMeBaseDataRequest {
	s.AppKey = &v
	return s
}

func (s *GetDingMeBaseDataRequest) SetByDay(v bool) *GetDingMeBaseDataRequest {
	s.ByDay = &v
	return s
}

func (s *GetDingMeBaseDataRequest) SetEndDay(v string) *GetDingMeBaseDataRequest {
	s.EndDay = &v
	return s
}

func (s *GetDingMeBaseDataRequest) SetStartDay(v string) *GetDingMeBaseDataRequest {
	s.StartDay = &v
	return s
}

type GetDingMeBaseDataResponseBody struct {
	// 是否缓存
	FromCache *bool `json:"fromCache,omitempty" xml:"fromCache,omitempty"`
	// 结果集
	Rawset []map[string]*string `json:"rawset,omitempty" xml:"rawset,omitempty" type:"Repeated"`
	// 运行时间
	Runtime *int64 `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// 字段解释
	Tips map[string]interface{} `json:"tips,omitempty" xml:"tips,omitempty"`
}

func (s GetDingMeBaseDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDingMeBaseDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetDingMeBaseDataResponseBody) SetFromCache(v bool) *GetDingMeBaseDataResponseBody {
	s.FromCache = &v
	return s
}

func (s *GetDingMeBaseDataResponseBody) SetRawset(v []map[string]*string) *GetDingMeBaseDataResponseBody {
	s.Rawset = v
	return s
}

func (s *GetDingMeBaseDataResponseBody) SetRuntime(v int64) *GetDingMeBaseDataResponseBody {
	s.Runtime = &v
	return s
}

func (s *GetDingMeBaseDataResponseBody) SetTips(v map[string]interface{}) *GetDingMeBaseDataResponseBody {
	s.Tips = v
	return s
}

type GetDingMeBaseDataResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDingMeBaseDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDingMeBaseDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDingMeBaseDataResponse) GoString() string {
	return s.String()
}

func (s *GetDingMeBaseDataResponse) SetHeaders(v map[string]*string) *GetDingMeBaseDataResponse {
	s.Headers = v
	return s
}

func (s *GetDingMeBaseDataResponse) SetBody(v *GetDingMeBaseDataResponseBody) *GetDingMeBaseDataResponse {
	s.Body = v
	return s
}

type GetIntelligentRobotInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetIntelligentRobotInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetIntelligentRobotInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetIntelligentRobotInfoHeaders) SetCommonHeaders(v map[string]*string) *GetIntelligentRobotInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetIntelligentRobotInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetIntelligentRobotInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetIntelligentRobotInfoRequest struct {
	// 机器人业务标识
	RobotAppKey *string `json:"robotAppKey,omitempty" xml:"robotAppKey,omitempty"`
}

func (s GetIntelligentRobotInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIntelligentRobotInfoRequest) GoString() string {
	return s.String()
}

func (s *GetIntelligentRobotInfoRequest) SetRobotAppKey(v string) *GetIntelligentRobotInfoRequest {
	s.RobotAppKey = &v
	return s
}

type GetIntelligentRobotInfoResponseBody struct {
	// 机器人id
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetIntelligentRobotInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIntelligentRobotInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetIntelligentRobotInfoResponseBody) SetResult(v string) *GetIntelligentRobotInfoResponseBody {
	s.Result = &v
	return s
}

type GetIntelligentRobotInfoResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetIntelligentRobotInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIntelligentRobotInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIntelligentRobotInfoResponse) GoString() string {
	return s.String()
}

func (s *GetIntelligentRobotInfoResponse) SetHeaders(v map[string]*string) *GetIntelligentRobotInfoResponse {
	s.Headers = v
	return s
}

func (s *GetIntelligentRobotInfoResponse) SetBody(v *GetIntelligentRobotInfoResponseBody) *GetIntelligentRobotInfoResponse {
	s.Body = v
	return s
}

type GetOfficialAccountRobotInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOfficialAccountRobotInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountRobotInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountRobotInfoHeaders) SetCommonHeaders(v map[string]*string) *GetOfficialAccountRobotInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOfficialAccountRobotInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetOfficialAccountRobotInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOfficialAccountRobotInfoRequest struct {
	// 机器人类型参数
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetOfficialAccountRobotInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountRobotInfoRequest) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountRobotInfoRequest) SetType(v string) *GetOfficialAccountRobotInfoRequest {
	s.Type = &v
	return s
}

type GetOfficialAccountRobotInfoResponseBody struct {
	// 机器人id
	AppId *int64 `json:"appId,omitempty" xml:"appId,omitempty"`
	// 机器人简介
	Brief *string `json:"brief,omitempty" xml:"brief,omitempty"`
	// 机器人描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 机器人icon
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 机器人名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 机器人预览图
	PreviewMediaUrl *string `json:"previewMediaUrl,omitempty" xml:"previewMediaUrl,omitempty"`
}

func (s GetOfficialAccountRobotInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountRobotInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountRobotInfoResponseBody) SetAppId(v int64) *GetOfficialAccountRobotInfoResponseBody {
	s.AppId = &v
	return s
}

func (s *GetOfficialAccountRobotInfoResponseBody) SetBrief(v string) *GetOfficialAccountRobotInfoResponseBody {
	s.Brief = &v
	return s
}

func (s *GetOfficialAccountRobotInfoResponseBody) SetDescription(v string) *GetOfficialAccountRobotInfoResponseBody {
	s.Description = &v
	return s
}

func (s *GetOfficialAccountRobotInfoResponseBody) SetIcon(v string) *GetOfficialAccountRobotInfoResponseBody {
	s.Icon = &v
	return s
}

func (s *GetOfficialAccountRobotInfoResponseBody) SetName(v string) *GetOfficialAccountRobotInfoResponseBody {
	s.Name = &v
	return s
}

func (s *GetOfficialAccountRobotInfoResponseBody) SetPreviewMediaUrl(v string) *GetOfficialAccountRobotInfoResponseBody {
	s.PreviewMediaUrl = &v
	return s
}

type GetOfficialAccountRobotInfoResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOfficialAccountRobotInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOfficialAccountRobotInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountRobotInfoResponse) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountRobotInfoResponse) SetHeaders(v map[string]*string) *GetOfficialAccountRobotInfoResponse {
	s.Headers = v
	return s
}

func (s *GetOfficialAccountRobotInfoResponse) SetBody(v *GetOfficialAccountRobotInfoResponseBody) *GetOfficialAccountRobotInfoResponse {
	s.Body = v
	return s
}

type GetWebChannelUserTokenHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetWebChannelUserTokenHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetWebChannelUserTokenHeaders) GoString() string {
	return s.String()
}

func (s *GetWebChannelUserTokenHeaders) SetCommonHeaders(v map[string]*string) *GetWebChannelUserTokenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetWebChannelUserTokenHeaders) SetXAcsDingtalkAccessToken(v string) *GetWebChannelUserTokenHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetWebChannelUserTokenRequest struct {
	// 登录用户在业务账号体系内的用户id
	ForeignId *string `json:"foreignId,omitempty" xml:"foreignId,omitempty"`
	// 登录用户在业务账号体系内的昵称
	Nick *string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 调用方在小蜜客服平台申请的业务账号体系的id
	Source *int64 `json:"source,omitempty" xml:"source,omitempty"`
}

func (s GetWebChannelUserTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWebChannelUserTokenRequest) GoString() string {
	return s.String()
}

func (s *GetWebChannelUserTokenRequest) SetForeignId(v string) *GetWebChannelUserTokenRequest {
	s.ForeignId = &v
	return s
}

func (s *GetWebChannelUserTokenRequest) SetNick(v string) *GetWebChannelUserTokenRequest {
	s.Nick = &v
	return s
}

func (s *GetWebChannelUserTokenRequest) SetSource(v int64) *GetWebChannelUserTokenRequest {
	s.Source = &v
	return s
}

type GetWebChannelUserTokenResponseBody struct {
	// 返回结果
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetWebChannelUserTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWebChannelUserTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetWebChannelUserTokenResponseBody) SetResult(v string) *GetWebChannelUserTokenResponseBody {
	s.Result = &v
	return s
}

type GetWebChannelUserTokenResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWebChannelUserTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWebChannelUserTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWebChannelUserTokenResponse) GoString() string {
	return s.String()
}

func (s *GetWebChannelUserTokenResponse) SetHeaders(v map[string]*string) *GetWebChannelUserTokenResponse {
	s.Headers = v
	return s
}

func (s *GetWebChannelUserTokenResponse) SetBody(v *GetWebChannelUserTokenResponseBody) *GetWebChannelUserTokenResponse {
	s.Body = v
	return s
}

type PushCustomerGroupMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PushCustomerGroupMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s PushCustomerGroupMessageHeaders) GoString() string {
	return s.String()
}

func (s *PushCustomerGroupMessageHeaders) SetCommonHeaders(v map[string]*string) *PushCustomerGroupMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PushCustomerGroupMessageHeaders) SetXAcsDingtalkAccessToken(v string) *PushCustomerGroupMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PushCustomerGroupMessageRequest struct {
	// 客户群会话id
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// 消息类型
	MsgKey *string `json:"msgKey,omitempty" xml:"msgKey,omitempty"`
	// 消息模板替换参数
	MsgParam *string `json:"msgParam,omitempty" xml:"msgParam,omitempty"`
}

func (s PushCustomerGroupMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s PushCustomerGroupMessageRequest) GoString() string {
	return s.String()
}

func (s *PushCustomerGroupMessageRequest) SetConversationId(v string) *PushCustomerGroupMessageRequest {
	s.ConversationId = &v
	return s
}

func (s *PushCustomerGroupMessageRequest) SetMsgKey(v string) *PushCustomerGroupMessageRequest {
	s.MsgKey = &v
	return s
}

func (s *PushCustomerGroupMessageRequest) SetMsgParam(v string) *PushCustomerGroupMessageRequest {
	s.MsgParam = &v
	return s
}

type PushCustomerGroupMessageResponseBody struct {
	// 推送queryKey
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s PushCustomerGroupMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushCustomerGroupMessageResponseBody) GoString() string {
	return s.String()
}

func (s *PushCustomerGroupMessageResponseBody) SetResult(v string) *PushCustomerGroupMessageResponseBody {
	s.Result = &v
	return s
}

type PushCustomerGroupMessageResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PushCustomerGroupMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushCustomerGroupMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s PushCustomerGroupMessageResponse) GoString() string {
	return s.String()
}

func (s *PushCustomerGroupMessageResponse) SetHeaders(v map[string]*string) *PushCustomerGroupMessageResponse {
	s.Headers = v
	return s
}

func (s *PushCustomerGroupMessageResponse) SetBody(v *PushCustomerGroupMessageResponseBody) *PushCustomerGroupMessageResponse {
	s.Body = v
	return s
}

type PushIntelligentRobotGroupMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PushIntelligentRobotGroupMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s PushIntelligentRobotGroupMessageHeaders) GoString() string {
	return s.String()
}

func (s *PushIntelligentRobotGroupMessageHeaders) SetCommonHeaders(v map[string]*string) *PushIntelligentRobotGroupMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PushIntelligentRobotGroupMessageHeaders) SetXAcsDingtalkAccessToken(v string) *PushIntelligentRobotGroupMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PushIntelligentRobotGroupMessageRequest struct {
	// 机器人id
	ChatbotId *string `json:"chatbotId,omitempty" xml:"chatbotId,omitempty"`
	// 消息类型
	MsgKey *string `json:"msgKey,omitempty" xml:"msgKey,omitempty"`
	// 消息模板替换参数
	MsgParam *string `json:"msgParam,omitempty" xml:"msgParam,omitempty"`
	// 群对话id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s PushIntelligentRobotGroupMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s PushIntelligentRobotGroupMessageRequest) GoString() string {
	return s.String()
}

func (s *PushIntelligentRobotGroupMessageRequest) SetChatbotId(v string) *PushIntelligentRobotGroupMessageRequest {
	s.ChatbotId = &v
	return s
}

func (s *PushIntelligentRobotGroupMessageRequest) SetMsgKey(v string) *PushIntelligentRobotGroupMessageRequest {
	s.MsgKey = &v
	return s
}

func (s *PushIntelligentRobotGroupMessageRequest) SetMsgParam(v string) *PushIntelligentRobotGroupMessageRequest {
	s.MsgParam = &v
	return s
}

func (s *PushIntelligentRobotGroupMessageRequest) SetOpenConversationId(v string) *PushIntelligentRobotGroupMessageRequest {
	s.OpenConversationId = &v
	return s
}

type PushIntelligentRobotGroupMessageResponseBody struct {
	// 推送queryKey
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s PushIntelligentRobotGroupMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushIntelligentRobotGroupMessageResponseBody) GoString() string {
	return s.String()
}

func (s *PushIntelligentRobotGroupMessageResponseBody) SetResult(v string) *PushIntelligentRobotGroupMessageResponseBody {
	s.Result = &v
	return s
}

type PushIntelligentRobotGroupMessageResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PushIntelligentRobotGroupMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushIntelligentRobotGroupMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s PushIntelligentRobotGroupMessageResponse) GoString() string {
	return s.String()
}

func (s *PushIntelligentRobotGroupMessageResponse) SetHeaders(v map[string]*string) *PushIntelligentRobotGroupMessageResponse {
	s.Headers = v
	return s
}

func (s *PushIntelligentRobotGroupMessageResponse) SetBody(v *PushIntelligentRobotGroupMessageResponseBody) *PushIntelligentRobotGroupMessageResponse {
	s.Body = v
	return s
}

type PushIntelligentRobotMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PushIntelligentRobotMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s PushIntelligentRobotMessageHeaders) GoString() string {
	return s.String()
}

func (s *PushIntelligentRobotMessageHeaders) SetCommonHeaders(v map[string]*string) *PushIntelligentRobotMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PushIntelligentRobotMessageHeaders) SetXAcsDingtalkAccessToken(v string) *PushIntelligentRobotMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PushIntelligentRobotMessageRequest struct {
	// 机器人id
	ChatbotId *string `json:"chatbotId,omitempty" xml:"chatbotId,omitempty"`
	// 消息类型
	MsgKey *string `json:"msgKey,omitempty" xml:"msgKey,omitempty"`
	// 消息模板替换参数
	MsgParam *string `json:"msgParam,omitempty" xml:"msgParam,omitempty"`
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PushIntelligentRobotMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s PushIntelligentRobotMessageRequest) GoString() string {
	return s.String()
}

func (s *PushIntelligentRobotMessageRequest) SetChatbotId(v string) *PushIntelligentRobotMessageRequest {
	s.ChatbotId = &v
	return s
}

func (s *PushIntelligentRobotMessageRequest) SetMsgKey(v string) *PushIntelligentRobotMessageRequest {
	s.MsgKey = &v
	return s
}

func (s *PushIntelligentRobotMessageRequest) SetMsgParam(v string) *PushIntelligentRobotMessageRequest {
	s.MsgParam = &v
	return s
}

func (s *PushIntelligentRobotMessageRequest) SetUserId(v string) *PushIntelligentRobotMessageRequest {
	s.UserId = &v
	return s
}

type PushIntelligentRobotMessageResponseBody struct {
	// 推送queryKey
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s PushIntelligentRobotMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushIntelligentRobotMessageResponseBody) GoString() string {
	return s.String()
}

func (s *PushIntelligentRobotMessageResponseBody) SetResult(v string) *PushIntelligentRobotMessageResponseBody {
	s.Result = &v
	return s
}

type PushIntelligentRobotMessageResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PushIntelligentRobotMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushIntelligentRobotMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s PushIntelligentRobotMessageResponse) GoString() string {
	return s.String()
}

func (s *PushIntelligentRobotMessageResponse) SetHeaders(v map[string]*string) *PushIntelligentRobotMessageResponse {
	s.Headers = v
	return s
}

func (s *PushIntelligentRobotMessageResponse) SetBody(v *PushIntelligentRobotMessageResponseBody) *PushIntelligentRobotMessageResponse {
	s.Body = v
	return s
}

type PushOfficialAccountMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PushOfficialAccountMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s PushOfficialAccountMessageHeaders) GoString() string {
	return s.String()
}

func (s *PushOfficialAccountMessageHeaders) SetCommonHeaders(v map[string]*string) *PushOfficialAccountMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PushOfficialAccountMessageHeaders) SetXAcsDingtalkAccessToken(v string) *PushOfficialAccountMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PushOfficialAccountMessageRequest struct {
	// 消息类型
	MsgKey *string `json:"msgKey,omitempty" xml:"msgKey,omitempty"`
	// 消息模板替换参数
	MsgParam *string `json:"msgParam,omitempty" xml:"msgParam,omitempty"`
	// 用户id(在服务窗对应虚拟企业中的用户id)
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PushOfficialAccountMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s PushOfficialAccountMessageRequest) GoString() string {
	return s.String()
}

func (s *PushOfficialAccountMessageRequest) SetMsgKey(v string) *PushOfficialAccountMessageRequest {
	s.MsgKey = &v
	return s
}

func (s *PushOfficialAccountMessageRequest) SetMsgParam(v string) *PushOfficialAccountMessageRequest {
	s.MsgParam = &v
	return s
}

func (s *PushOfficialAccountMessageRequest) SetUserId(v string) *PushOfficialAccountMessageRequest {
	s.UserId = &v
	return s
}

type PushOfficialAccountMessageResponseBody struct {
	// 推送queryKey
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s PushOfficialAccountMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushOfficialAccountMessageResponseBody) GoString() string {
	return s.String()
}

func (s *PushOfficialAccountMessageResponseBody) SetResult(v string) *PushOfficialAccountMessageResponseBody {
	s.Result = &v
	return s
}

type PushOfficialAccountMessageResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PushOfficialAccountMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushOfficialAccountMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s PushOfficialAccountMessageResponse) GoString() string {
	return s.String()
}

func (s *PushOfficialAccountMessageResponse) SetHeaders(v map[string]*string) *PushOfficialAccountMessageResponse {
	s.Headers = v
	return s
}

func (s *PushOfficialAccountMessageResponse) SetBody(v *PushOfficialAccountMessageResponseBody) *PushOfficialAccountMessageResponse {
	s.Body = v
	return s
}

type PushRobotMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PushRobotMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s PushRobotMessageHeaders) GoString() string {
	return s.String()
}

func (s *PushRobotMessageHeaders) SetCommonHeaders(v map[string]*string) *PushRobotMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PushRobotMessageHeaders) SetXAcsDingtalkAccessToken(v string) *PushRobotMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PushRobotMessageRequest struct {
	// 机器人id
	ChatbotId *string `json:"chatbotId,omitempty" xml:"chatbotId,omitempty"`
	// 消息类型
	MsgKey *string `json:"msgKey,omitempty" xml:"msgKey,omitempty"`
	// 消息模板替换参数
	MsgParam *string `json:"msgParam,omitempty" xml:"msgParam,omitempty"`
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PushRobotMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s PushRobotMessageRequest) GoString() string {
	return s.String()
}

func (s *PushRobotMessageRequest) SetChatbotId(v string) *PushRobotMessageRequest {
	s.ChatbotId = &v
	return s
}

func (s *PushRobotMessageRequest) SetMsgKey(v string) *PushRobotMessageRequest {
	s.MsgKey = &v
	return s
}

func (s *PushRobotMessageRequest) SetMsgParam(v string) *PushRobotMessageRequest {
	s.MsgParam = &v
	return s
}

func (s *PushRobotMessageRequest) SetUserId(v string) *PushRobotMessageRequest {
	s.UserId = &v
	return s
}

type PushRobotMessageResponseBody struct {
	// 推送queryKey
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s PushRobotMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushRobotMessageResponseBody) GoString() string {
	return s.String()
}

func (s *PushRobotMessageResponseBody) SetResult(v string) *PushRobotMessageResponseBody {
	s.Result = &v
	return s
}

type PushRobotMessageResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PushRobotMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushRobotMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s PushRobotMessageResponse) GoString() string {
	return s.String()
}

func (s *PushRobotMessageResponse) SetHeaders(v map[string]*string) *PushRobotMessageResponse {
	s.Headers = v
	return s
}

func (s *PushRobotMessageResponse) SetBody(v *PushRobotMessageResponseBody) *PushRobotMessageResponse {
	s.Body = v
	return s
}

type ReplyRobotHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ReplyRobotHeaders) String() string {
	return tea.Prettify(s)
}

func (s ReplyRobotHeaders) GoString() string {
	return s.String()
}

func (s *ReplyRobotHeaders) SetCommonHeaders(v map[string]*string) *ReplyRobotHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReplyRobotHeaders) SetXAcsDingtalkAccessToken(v string) *ReplyRobotHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ReplyRobotRequest struct {
	// 回复消息内容的json string
	ProxyMessageStr *string `json:"proxyMessageStr,omitempty" xml:"proxyMessageStr,omitempty"`
}

func (s ReplyRobotRequest) String() string {
	return tea.Prettify(s)
}

func (s ReplyRobotRequest) GoString() string {
	return s.String()
}

func (s *ReplyRobotRequest) SetProxyMessageStr(v string) *ReplyRobotRequest {
	s.ProxyMessageStr = &v
	return s
}

type ReplyRobotResponseBody struct {
	// 回复是否成功结果
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ReplyRobotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReplyRobotResponseBody) GoString() string {
	return s.String()
}

func (s *ReplyRobotResponseBody) SetResult(v bool) *ReplyRobotResponseBody {
	s.Result = &v
	return s
}

type ReplyRobotResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReplyRobotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReplyRobotResponse) String() string {
	return tea.Prettify(s)
}

func (s ReplyRobotResponse) GoString() string {
	return s.String()
}

func (s *ReplyRobotResponse) SetHeaders(v map[string]*string) *ReplyRobotResponse {
	s.Headers = v
	return s
}

func (s *ReplyRobotResponse) SetBody(v *ReplyRobotResponseBody) *ReplyRobotResponse {
	s.Body = v
	return s
}

type UpdateOfficialAccountRobotInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateOfficialAccountRobotInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateOfficialAccountRobotInfoHeaders) GoString() string {
	return s.String()
}

func (s *UpdateOfficialAccountRobotInfoHeaders) SetCommonHeaders(v map[string]*string) *UpdateOfficialAccountRobotInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateOfficialAccountRobotInfoHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateOfficialAccountRobotInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateOfficialAccountRobotInfoRequest struct {
	// 机器人头像
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// 机器人简介
	Brief *string `json:"brief,omitempty" xml:"brief,omitempty"`
	// 机器人描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 机器人名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 机器人预览图
	PreviewMediaUrl *string `json:"previewMediaUrl,omitempty" xml:"previewMediaUrl,omitempty"`
	// 机器人类型参数
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateOfficialAccountRobotInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOfficialAccountRobotInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateOfficialAccountRobotInfoRequest) SetAvatar(v string) *UpdateOfficialAccountRobotInfoRequest {
	s.Avatar = &v
	return s
}

func (s *UpdateOfficialAccountRobotInfoRequest) SetBrief(v string) *UpdateOfficialAccountRobotInfoRequest {
	s.Brief = &v
	return s
}

func (s *UpdateOfficialAccountRobotInfoRequest) SetDescription(v string) *UpdateOfficialAccountRobotInfoRequest {
	s.Description = &v
	return s
}

func (s *UpdateOfficialAccountRobotInfoRequest) SetName(v string) *UpdateOfficialAccountRobotInfoRequest {
	s.Name = &v
	return s
}

func (s *UpdateOfficialAccountRobotInfoRequest) SetPreviewMediaUrl(v string) *UpdateOfficialAccountRobotInfoRequest {
	s.PreviewMediaUrl = &v
	return s
}

func (s *UpdateOfficialAccountRobotInfoRequest) SetType(v string) *UpdateOfficialAccountRobotInfoRequest {
	s.Type = &v
	return s
}

type UpdateOfficialAccountRobotInfoResponseBody struct {
	// 结果
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateOfficialAccountRobotInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOfficialAccountRobotInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOfficialAccountRobotInfoResponseBody) SetResult(v string) *UpdateOfficialAccountRobotInfoResponseBody {
	s.Result = &v
	return s
}

type UpdateOfficialAccountRobotInfoResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateOfficialAccountRobotInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateOfficialAccountRobotInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOfficialAccountRobotInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateOfficialAccountRobotInfoResponse) SetHeaders(v map[string]*string) *UpdateOfficialAccountRobotInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateOfficialAccountRobotInfoResponse) SetBody(v *UpdateOfficialAccountRobotInfoResponseBody) *UpdateOfficialAccountRobotInfoResponse {
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

func (client *Client) AddRobotInstanceToGroup(request *AddRobotInstanceToGroupRequest) (_result *AddRobotInstanceToGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddRobotInstanceToGroupHeaders{}
	_result = &AddRobotInstanceToGroupResponse{}
	_body, _err := client.AddRobotInstanceToGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddRobotInstanceToGroupWithOptions(request *AddRobotInstanceToGroupRequest, headers *AddRobotInstanceToGroupHeaders, runtime *util.RuntimeOptions) (_result *AddRobotInstanceToGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChatbotId)) {
		body["chatbotId"] = request.ChatbotId
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
	_result = &AddRobotInstanceToGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("AddRobotInstanceToGroup"), tea.String("dingmi_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/dingmi/intelligentRobots/groups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AskRobot(request *AskRobotRequest) (_result *AskRobotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AskRobotHeaders{}
	_result = &AskRobotResponse{}
	_body, _err := client.AskRobotWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AskRobotWithOptions(request *AskRobotRequest, headers *AskRobotHeaders, runtime *util.RuntimeOptions) (_result *AskRobotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingUserId)) {
		body["dingUserId"] = request.DingUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Question)) {
		body["question"] = request.Question
	}

	if !tea.BoolValue(util.IsUnset(request.RobotAppKey)) {
		body["robotAppKey"] = request.RobotAppKey
	}

	if !tea.BoolValue(util.IsUnset(request.SessionUuid)) {
		body["sessionUuid"] = request.SessionUuid
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
	_result = &AskRobotResponse{}
	_body, _err := client.DoROARequest(tea.String("AskRobot"), tea.String("dingmi_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/dingmi/robots/ask"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDingMeBaseData(request *GetDingMeBaseDataRequest) (_result *GetDingMeBaseDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDingMeBaseDataHeaders{}
	_result = &GetDingMeBaseDataResponse{}
	_body, _err := client.GetDingMeBaseDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDingMeBaseDataWithOptions(request *GetDingMeBaseDataRequest, headers *GetDingMeBaseDataHeaders, runtime *util.RuntimeOptions) (_result *GetDingMeBaseDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["appKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.ByDay)) {
		query["byDay"] = request.ByDay
	}

	if !tea.BoolValue(util.IsUnset(request.EndDay)) {
		query["endDay"] = request.EndDay
	}

	if !tea.BoolValue(util.IsUnset(request.StartDay)) {
		query["startDay"] = request.StartDay
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
	_result = &GetDingMeBaseDataResponse{}
	_body, _err := client.DoROARequest(tea.String("GetDingMeBaseData"), tea.String("dingmi_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/dingmi/robots/data"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIntelligentRobotInfo(request *GetIntelligentRobotInfoRequest) (_result *GetIntelligentRobotInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetIntelligentRobotInfoHeaders{}
	_result = &GetIntelligentRobotInfoResponse{}
	_body, _err := client.GetIntelligentRobotInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIntelligentRobotInfoWithOptions(request *GetIntelligentRobotInfoRequest, headers *GetIntelligentRobotInfoHeaders, runtime *util.RuntimeOptions) (_result *GetIntelligentRobotInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RobotAppKey)) {
		query["robotAppKey"] = request.RobotAppKey
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
	_result = &GetIntelligentRobotInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetIntelligentRobotInfo"), tea.String("dingmi_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/dingmi/intelligentRobots/infos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOfficialAccountRobotInfo(request *GetOfficialAccountRobotInfoRequest) (_result *GetOfficialAccountRobotInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOfficialAccountRobotInfoHeaders{}
	_result = &GetOfficialAccountRobotInfoResponse{}
	_body, _err := client.GetOfficialAccountRobotInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOfficialAccountRobotInfoWithOptions(request *GetOfficialAccountRobotInfoRequest, headers *GetOfficialAccountRobotInfoHeaders, runtime *util.RuntimeOptions) (_result *GetOfficialAccountRobotInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
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
	_result = &GetOfficialAccountRobotInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetOfficialAccountRobotInfo"), tea.String("dingmi_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/dingmi/officialAccounts/robots"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWebChannelUserToken(request *GetWebChannelUserTokenRequest) (_result *GetWebChannelUserTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetWebChannelUserTokenHeaders{}
	_result = &GetWebChannelUserTokenResponse{}
	_body, _err := client.GetWebChannelUserTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWebChannelUserTokenWithOptions(request *GetWebChannelUserTokenRequest, headers *GetWebChannelUserTokenHeaders, runtime *util.RuntimeOptions) (_result *GetWebChannelUserTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ForeignId)) {
		body["foreignId"] = request.ForeignId
	}

	if !tea.BoolValue(util.IsUnset(request.Nick)) {
		body["nick"] = request.Nick
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
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
	_result = &GetWebChannelUserTokenResponse{}
	_body, _err := client.DoROARequest(tea.String("GetWebChannelUserToken"), tea.String("dingmi_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/dingmi/webChannels/userTokens"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushCustomerGroupMessage(request *PushCustomerGroupMessageRequest) (_result *PushCustomerGroupMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PushCustomerGroupMessageHeaders{}
	_result = &PushCustomerGroupMessageResponse{}
	_body, _err := client.PushCustomerGroupMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushCustomerGroupMessageWithOptions(request *PushCustomerGroupMessageRequest, headers *PushCustomerGroupMessageHeaders, runtime *util.RuntimeOptions) (_result *PushCustomerGroupMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		body["conversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.MsgKey)) {
		body["msgKey"] = request.MsgKey
	}

	if !tea.BoolValue(util.IsUnset(request.MsgParam)) {
		body["msgParam"] = request.MsgParam
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
	_result = &PushCustomerGroupMessageResponse{}
	_body, _err := client.DoROARequest(tea.String("PushCustomerGroupMessage"), tea.String("dingmi_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/dingmi/officialAccounts/robots/groupMessages/send"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushIntelligentRobotGroupMessage(request *PushIntelligentRobotGroupMessageRequest) (_result *PushIntelligentRobotGroupMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PushIntelligentRobotGroupMessageHeaders{}
	_result = &PushIntelligentRobotGroupMessageResponse{}
	_body, _err := client.PushIntelligentRobotGroupMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushIntelligentRobotGroupMessageWithOptions(request *PushIntelligentRobotGroupMessageRequest, headers *PushIntelligentRobotGroupMessageHeaders, runtime *util.RuntimeOptions) (_result *PushIntelligentRobotGroupMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChatbotId)) {
		body["chatbotId"] = request.ChatbotId
	}

	if !tea.BoolValue(util.IsUnset(request.MsgKey)) {
		body["msgKey"] = request.MsgKey
	}

	if !tea.BoolValue(util.IsUnset(request.MsgParam)) {
		body["msgParam"] = request.MsgParam
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
	_result = &PushIntelligentRobotGroupMessageResponse{}
	_body, _err := client.DoROARequest(tea.String("PushIntelligentRobotGroupMessage"), tea.String("dingmi_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/dingmi/intelligentRobots/groupMessages/send"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushIntelligentRobotMessage(request *PushIntelligentRobotMessageRequest) (_result *PushIntelligentRobotMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PushIntelligentRobotMessageHeaders{}
	_result = &PushIntelligentRobotMessageResponse{}
	_body, _err := client.PushIntelligentRobotMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushIntelligentRobotMessageWithOptions(request *PushIntelligentRobotMessageRequest, headers *PushIntelligentRobotMessageHeaders, runtime *util.RuntimeOptions) (_result *PushIntelligentRobotMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChatbotId)) {
		body["chatbotId"] = request.ChatbotId
	}

	if !tea.BoolValue(util.IsUnset(request.MsgKey)) {
		body["msgKey"] = request.MsgKey
	}

	if !tea.BoolValue(util.IsUnset(request.MsgParam)) {
		body["msgParam"] = request.MsgParam
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
	_result = &PushIntelligentRobotMessageResponse{}
	_body, _err := client.DoROARequest(tea.String("PushIntelligentRobotMessage"), tea.String("dingmi_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/dingmi/intelligentRobots/oToMessages/send"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushOfficialAccountMessage(request *PushOfficialAccountMessageRequest) (_result *PushOfficialAccountMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PushOfficialAccountMessageHeaders{}
	_result = &PushOfficialAccountMessageResponse{}
	_body, _err := client.PushOfficialAccountMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushOfficialAccountMessageWithOptions(request *PushOfficialAccountMessageRequest, headers *PushOfficialAccountMessageHeaders, runtime *util.RuntimeOptions) (_result *PushOfficialAccountMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MsgKey)) {
		body["msgKey"] = request.MsgKey
	}

	if !tea.BoolValue(util.IsUnset(request.MsgParam)) {
		body["msgParam"] = request.MsgParam
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
	_result = &PushOfficialAccountMessageResponse{}
	_body, _err := client.DoROARequest(tea.String("PushOfficialAccountMessage"), tea.String("dingmi_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/dingmi/officialAccounts/robots/oToMessages/send"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushRobotMessage(request *PushRobotMessageRequest) (_result *PushRobotMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PushRobotMessageHeaders{}
	_result = &PushRobotMessageResponse{}
	_body, _err := client.PushRobotMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushRobotMessageWithOptions(request *PushRobotMessageRequest, headers *PushRobotMessageHeaders, runtime *util.RuntimeOptions) (_result *PushRobotMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChatbotId)) {
		body["chatbotId"] = request.ChatbotId
	}

	if !tea.BoolValue(util.IsUnset(request.MsgKey)) {
		body["msgKey"] = request.MsgKey
	}

	if !tea.BoolValue(util.IsUnset(request.MsgParam)) {
		body["msgParam"] = request.MsgParam
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
	_result = &PushRobotMessageResponse{}
	_body, _err := client.DoROARequest(tea.String("PushRobotMessage"), tea.String("dingmi_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/dingmi/robots/oToMessages/send"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReplyRobot(request *ReplyRobotRequest) (_result *ReplyRobotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ReplyRobotHeaders{}
	_result = &ReplyRobotResponse{}
	_body, _err := client.ReplyRobotWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReplyRobotWithOptions(request *ReplyRobotRequest, headers *ReplyRobotHeaders, runtime *util.RuntimeOptions) (_result *ReplyRobotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProxyMessageStr)) {
		body["proxyMessageStr"] = request.ProxyMessageStr
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
	_result = &ReplyRobotResponse{}
	_body, _err := client.DoROARequest(tea.String("ReplyRobot"), tea.String("dingmi_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/dingmi/robots/reply"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateOfficialAccountRobotInfo(request *UpdateOfficialAccountRobotInfoRequest) (_result *UpdateOfficialAccountRobotInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateOfficialAccountRobotInfoHeaders{}
	_result = &UpdateOfficialAccountRobotInfoResponse{}
	_body, _err := client.UpdateOfficialAccountRobotInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateOfficialAccountRobotInfoWithOptions(request *UpdateOfficialAccountRobotInfoRequest, headers *UpdateOfficialAccountRobotInfoHeaders, runtime *util.RuntimeOptions) (_result *UpdateOfficialAccountRobotInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Avatar)) {
		body["avatar"] = request.Avatar
	}

	if !tea.BoolValue(util.IsUnset(request.Brief)) {
		body["brief"] = request.Brief
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PreviewMediaUrl)) {
		body["previewMediaUrl"] = request.PreviewMediaUrl
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
	_result = &UpdateOfficialAccountRobotInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateOfficialAccountRobotInfo"), tea.String("dingmi_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/dingmi/officialAccounts/robots"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
