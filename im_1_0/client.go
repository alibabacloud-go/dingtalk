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
	Status             *string `json:"status,omitempty" xml:"status,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOauthAppId     *int64  `json:"dingOauthAppId,omitempty" xml:"dingOauthAppId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
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

func (s *UpdateGroupPermissionRequest) SetDingTokenGrantType(v int64) *UpdateGroupPermissionRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *UpdateGroupPermissionRequest) SetDingOrgId(v int64) *UpdateGroupPermissionRequest {
	s.DingOrgId = &v
	return s
}

func (s *UpdateGroupPermissionRequest) SetDingIsvOrgId(v int64) *UpdateGroupPermissionRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *UpdateGroupPermissionRequest) SetDingOauthAppId(v int64) *UpdateGroupPermissionRequest {
	s.DingOauthAppId = &v
	return s
}

func (s *UpdateGroupPermissionRequest) SetDingSuiteKey(v string) *UpdateGroupPermissionRequest {
	s.DingSuiteKey = &v
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
	// 用户ID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 群角色列表
	OpenRoleIds        []*string `json:"openRoleIds,omitempty" xml:"openRoleIds,omitempty" type:"Repeated"`
	DingTokenGrantType *int64    `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	DingOrgId          *int64    `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingIsvOrgId       *int64    `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingSuiteKey       *string   `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingOauthAppId     *int64    `json:"dingOauthAppId,omitempty" xml:"dingOauthAppId,omitempty"`
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

func (s *UpdateTheGroupRolesOfGroupMemberRequest) SetUserId(v string) *UpdateTheGroupRolesOfGroupMemberRequest {
	s.UserId = &v
	return s
}

func (s *UpdateTheGroupRolesOfGroupMemberRequest) SetOpenRoleIds(v []*string) *UpdateTheGroupRolesOfGroupMemberRequest {
	s.OpenRoleIds = v
	return s
}

func (s *UpdateTheGroupRolesOfGroupMemberRequest) SetDingTokenGrantType(v int64) *UpdateTheGroupRolesOfGroupMemberRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *UpdateTheGroupRolesOfGroupMemberRequest) SetDingOrgId(v int64) *UpdateTheGroupRolesOfGroupMemberRequest {
	s.DingOrgId = &v
	return s
}

func (s *UpdateTheGroupRolesOfGroupMemberRequest) SetDingIsvOrgId(v int64) *UpdateTheGroupRolesOfGroupMemberRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *UpdateTheGroupRolesOfGroupMemberRequest) SetDingSuiteKey(v string) *UpdateTheGroupRolesOfGroupMemberRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *UpdateTheGroupRolesOfGroupMemberRequest) SetDingOauthAppId(v int64) *UpdateTheGroupRolesOfGroupMemberRequest {
	s.DingOauthAppId = &v
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
	DingIsvOrgId *int64 `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	// 卡片模板ID
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	// 接收卡片的群的openConversationId
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 接收人userId列表
	ReceiverUserIdList []*string `json:"receiverUserIdList,omitempty" xml:"receiverUserIdList,omitempty" type:"Repeated"`
	DingTokenGrantType *int64    `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	// 唯一标识一张卡片的外部ID（卡片幂等ID，可用于更新或重复发送同一卡片到多个群会话）
	OutTrackId   *string `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	DingSuiteKey *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	// 【robotCode & chatBotId二选一必填】机器人编码（群模板机器人）
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	DingOrgId *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	// 发送的会话类型：单聊-0, 群聊-1（单聊时：openConversationId不用填写；receiverUserIdList填写有且一个员工号）
	ConversationType *int32 `json:"conversationType,omitempty" xml:"conversationType,omitempty"`
	// 可控制卡片回调时的路由Key，用于指定特定的callbackUrl【可空：不填写默认用企业的回调地址】
	CallbackRouteKey *string                             `json:"callbackRouteKey,omitempty" xml:"callbackRouteKey,omitempty"`
	CardData         *SendInteractiveCardRequestCardData `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	// 指定用户可见的按钮列表（key：用户userId；value：用户数据）
	PrivateData    map[string]*PrivateDataValue `json:"privateData,omitempty" xml:"privateData,omitempty"`
	DingOauthAppId *int64                       `json:"dingOauthAppId,omitempty" xml:"dingOauthAppId,omitempty"`
	// 【robotCode & chatBotId二选一必填】机器人ID（企业机器人）
	ChatBotId *string `json:"chatBotId,omitempty" xml:"chatBotId,omitempty"`
	// 用户ID类型：1：staffId模式【默认】；2：unionId模式；对应receiverUserIdList、privateData字段关于用户id的值填写方式
	UserIdType *int32 `json:"userIdType,omitempty" xml:"userIdType,omitempty"`
	// 消息@人，{123456:"钉三多"}，key：根据userIdType来设置，【特殊设置：如果key、value都为"@ALL"则判断at所有人】
	AtOpenIds map[string]*string `json:"atOpenIds,omitempty" xml:"atOpenIds,omitempty"`
}

func (s SendInteractiveCardRequest) String() string {
	return tea.Prettify(s)
}

func (s SendInteractiveCardRequest) GoString() string {
	return s.String()
}

func (s *SendInteractiveCardRequest) SetDingIsvOrgId(v int64) *SendInteractiveCardRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *SendInteractiveCardRequest) SetCardTemplateId(v string) *SendInteractiveCardRequest {
	s.CardTemplateId = &v
	return s
}

func (s *SendInteractiveCardRequest) SetOpenConversationId(v string) *SendInteractiveCardRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SendInteractiveCardRequest) SetReceiverUserIdList(v []*string) *SendInteractiveCardRequest {
	s.ReceiverUserIdList = v
	return s
}

func (s *SendInteractiveCardRequest) SetDingTokenGrantType(v int64) *SendInteractiveCardRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *SendInteractiveCardRequest) SetOutTrackId(v string) *SendInteractiveCardRequest {
	s.OutTrackId = &v
	return s
}

func (s *SendInteractiveCardRequest) SetDingSuiteKey(v string) *SendInteractiveCardRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *SendInteractiveCardRequest) SetRobotCode(v string) *SendInteractiveCardRequest {
	s.RobotCode = &v
	return s
}

func (s *SendInteractiveCardRequest) SetDingOrgId(v int64) *SendInteractiveCardRequest {
	s.DingOrgId = &v
	return s
}

func (s *SendInteractiveCardRequest) SetConversationType(v int32) *SendInteractiveCardRequest {
	s.ConversationType = &v
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

func (s *SendInteractiveCardRequest) SetPrivateData(v map[string]*PrivateDataValue) *SendInteractiveCardRequest {
	s.PrivateData = v
	return s
}

func (s *SendInteractiveCardRequest) SetDingOauthAppId(v int64) *SendInteractiveCardRequest {
	s.DingOauthAppId = &v
	return s
}

func (s *SendInteractiveCardRequest) SetChatBotId(v string) *SendInteractiveCardRequest {
	s.ChatBotId = &v
	return s
}

func (s *SendInteractiveCardRequest) SetUserIdType(v int32) *SendInteractiveCardRequest {
	s.UserIdType = &v
	return s
}

func (s *SendInteractiveCardRequest) SetAtOpenIds(v map[string]*string) *SendInteractiveCardRequest {
	s.AtOpenIds = v
	return s
}

type SendInteractiveCardRequestCardData struct {
	// 卡片模板内容替换参数-普通文本类型
	CardParamMap map[string]*string `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
	// 卡片模板内容替换参数-多媒体类型
	CardMediaIdParamMap map[string]*string `json:"cardMediaIdParamMap,omitempty" xml:"cardMediaIdParamMap,omitempty"`
}

func (s SendInteractiveCardRequestCardData) String() string {
	return tea.Prettify(s)
}

func (s SendInteractiveCardRequestCardData) GoString() string {
	return s.String()
}

func (s *SendInteractiveCardRequestCardData) SetCardParamMap(v map[string]*string) *SendInteractiveCardRequestCardData {
	s.CardParamMap = v
	return s
}

func (s *SendInteractiveCardRequestCardData) SetCardMediaIdParamMap(v map[string]*string) *SendInteractiveCardRequestCardData {
	s.CardMediaIdParamMap = v
	return s
}

type SendInteractiveCardResponseBody struct {
	// 结果
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SendInteractiveCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendInteractiveCardResponseBody) GoString() string {
	return s.String()
}

func (s *SendInteractiveCardResponseBody) SetSuccess(v bool) *SendInteractiveCardResponseBody {
	s.Success = &v
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
	// 唯一标识一张卡片的外部ID
	OutTrackId         *string                               `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	CardData           *UpdateInteractiveCardRequestCardData `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	PrivateData        map[string]*PrivateDataValue          `json:"privateData,omitempty" xml:"privateData,omitempty"`
	DingTokenGrantType *int64                                `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	DingOrgId          *int64                                `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingIsvOrgId       *int64                                `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingSuiteKey       *string                               `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingOauthAppId     *int64                                `json:"dingOauthAppId,omitempty" xml:"dingOauthAppId,omitempty"`
	// 用户ID类型：1：staffId模式【默认】；2：unionId模式；对应receiverUserIdList、privateData字段关于用户id的值填写方式
	UserIdType *int32 `json:"userIdType,omitempty" xml:"userIdType,omitempty"`
}

func (s UpdateInteractiveCardRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInteractiveCardRequest) GoString() string {
	return s.String()
}

func (s *UpdateInteractiveCardRequest) SetOutTrackId(v string) *UpdateInteractiveCardRequest {
	s.OutTrackId = &v
	return s
}

func (s *UpdateInteractiveCardRequest) SetCardData(v *UpdateInteractiveCardRequestCardData) *UpdateInteractiveCardRequest {
	s.CardData = v
	return s
}

func (s *UpdateInteractiveCardRequest) SetPrivateData(v map[string]*PrivateDataValue) *UpdateInteractiveCardRequest {
	s.PrivateData = v
	return s
}

func (s *UpdateInteractiveCardRequest) SetDingTokenGrantType(v int64) *UpdateInteractiveCardRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *UpdateInteractiveCardRequest) SetDingOrgId(v int64) *UpdateInteractiveCardRequest {
	s.DingOrgId = &v
	return s
}

func (s *UpdateInteractiveCardRequest) SetDingIsvOrgId(v int64) *UpdateInteractiveCardRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *UpdateInteractiveCardRequest) SetDingSuiteKey(v string) *UpdateInteractiveCardRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *UpdateInteractiveCardRequest) SetDingOauthAppId(v int64) *UpdateInteractiveCardRequest {
	s.DingOauthAppId = &v
	return s
}

func (s *UpdateInteractiveCardRequest) SetUserIdType(v int32) *UpdateInteractiveCardRequest {
	s.UserIdType = &v
	return s
}

type UpdateInteractiveCardRequestCardData struct {
	CardParamMap        map[string]*string `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
	CardMediaIdParamMap map[string]*string `json:"cardMediaIdParamMap,omitempty" xml:"cardMediaIdParamMap,omitempty"`
}

func (s UpdateInteractiveCardRequestCardData) String() string {
	return tea.Prettify(s)
}

func (s UpdateInteractiveCardRequestCardData) GoString() string {
	return s.String()
}

func (s *UpdateInteractiveCardRequestCardData) SetCardParamMap(v map[string]*string) *UpdateInteractiveCardRequestCardData {
	s.CardParamMap = v
	return s
}

func (s *UpdateInteractiveCardRequestCardData) SetCardMediaIdParamMap(v map[string]*string) *UpdateInteractiveCardRequestCardData {
	s.CardMediaIdParamMap = v
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
	Timestamp          *int64  `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingOauthAppId     *int64  `json:"dingOauthAppId,omitempty" xml:"dingOauthAppId,omitempty"`
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

func (s *QueryMembersOfGroupRoleRequest) SetDingTokenGrantType(v int64) *QueryMembersOfGroupRoleRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *QueryMembersOfGroupRoleRequest) SetDingOrgId(v int64) *QueryMembersOfGroupRoleRequest {
	s.DingOrgId = &v
	return s
}

func (s *QueryMembersOfGroupRoleRequest) SetDingIsvOrgId(v int64) *QueryMembersOfGroupRoleRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *QueryMembersOfGroupRoleRequest) SetDingSuiteKey(v string) *QueryMembersOfGroupRoleRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *QueryMembersOfGroupRoleRequest) SetDingOauthAppId(v int64) *QueryMembersOfGroupRoleRequest {
	s.DingOauthAppId = &v
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

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOauthAppId)) {
		body["dingOauthAppId"] = request.DingOauthAppId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
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
	_result = &UpdateGroupPermissionResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateGroupPermission"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/im/sceneGroups/permissions"), tea.String("json"), req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenRoleIds)) {
		body["openRoleIds"] = request.OpenRoleIds
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.DingOauthAppId)) {
		body["dingOauthAppId"] = request.DingOauthAppId
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
	_result = &UpdateTheGroupRolesOfGroupMemberResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateTheGroupRolesOfGroupMember"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/im/sceneGroups/members/groupRoles"), tea.String("json"), req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.CardTemplateId)) {
		body["cardTemplateId"] = request.CardTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUserIdList)) {
		body["receiverUserIdList"] = request.ReceiverUserIdList
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.OutTrackId)) {
		body["outTrackId"] = request.OutTrackId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.ConversationType)) {
		body["conversationType"] = request.ConversationType
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackRouteKey)) {
		body["callbackRouteKey"] = request.CallbackRouteKey
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CardData))) {
		body["cardData"] = request.CardData
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateData)) {
		body["privateData"] = request.PrivateData
	}

	if !tea.BoolValue(util.IsUnset(request.DingOauthAppId)) {
		body["dingOauthAppId"] = request.DingOauthAppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChatBotId)) {
		body["chatBotId"] = request.ChatBotId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdType)) {
		body["userIdType"] = request.UserIdType
	}

	if !tea.BoolValue(util.IsUnset(request.AtOpenIds)) {
		body["atOpenIds"] = request.AtOpenIds
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
	_result = &SendInteractiveCardResponse{}
	_body, _err := client.DoROARequest(tea.String("SendInteractiveCard"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/interactiveCards/send"), tea.String("json"), req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.OutTrackId)) {
		body["outTrackId"] = request.OutTrackId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CardData))) {
		body["cardData"] = request.CardData
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateData)) {
		body["privateData"] = request.PrivateData
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.DingOauthAppId)) {
		body["dingOauthAppId"] = request.DingOauthAppId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdType)) {
		body["userIdType"] = request.UserIdType
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
	_result = &UpdateInteractiveCardResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateInteractiveCard"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/im/interactiveCards"), tea.String("json"), req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.DingOauthAppId)) {
		body["dingOauthAppId"] = request.DingOauthAppId
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
	_result = &QueryMembersOfGroupRoleResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryMembersOfGroupRole"), tea.String("im_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/im/sceneGroups/roles/members/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
