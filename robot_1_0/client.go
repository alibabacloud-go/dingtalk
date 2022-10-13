// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package robot_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BatchOTOQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchOTOQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchOTOQueryHeaders) GoString() string {
	return s.String()
}

func (s *BatchOTOQueryHeaders) SetCommonHeaders(v map[string]*string) *BatchOTOQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchOTOQueryHeaders) SetXAcsDingtalkAccessToken(v string) *BatchOTOQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchOTOQueryRequest struct {
	// 消息已读查询标志
	ProcessQueryKey *string `json:"processQueryKey,omitempty" xml:"processQueryKey,omitempty"`
	// 机器人robotCode
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s BatchOTOQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchOTOQueryRequest) GoString() string {
	return s.String()
}

func (s *BatchOTOQueryRequest) SetProcessQueryKey(v string) *BatchOTOQueryRequest {
	s.ProcessQueryKey = &v
	return s
}

func (s *BatchOTOQueryRequest) SetRobotCode(v string) *BatchOTOQueryRequest {
	s.RobotCode = &v
	return s
}

type BatchOTOQueryResponseBody struct {
	// 消息已读情况
	MessageReadInfoList []*BatchOTOQueryResponseBodyMessageReadInfoList `json:"messageReadInfoList,omitempty" xml:"messageReadInfoList,omitempty" type:"Repeated"`
	// 消息发送状态
	SendStatus *string `json:"sendStatus,omitempty" xml:"sendStatus,omitempty"`
}

func (s BatchOTOQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchOTOQueryResponseBody) GoString() string {
	return s.String()
}

func (s *BatchOTOQueryResponseBody) SetMessageReadInfoList(v []*BatchOTOQueryResponseBodyMessageReadInfoList) *BatchOTOQueryResponseBody {
	s.MessageReadInfoList = v
	return s
}

func (s *BatchOTOQueryResponseBody) SetSendStatus(v string) *BatchOTOQueryResponseBody {
	s.SendStatus = &v
	return s
}

type BatchOTOQueryResponseBodyMessageReadInfoList struct {
	// 姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 已读状态
	ReadStatus *string `json:"readStatus,omitempty" xml:"readStatus,omitempty"`
	// 已读时间
	ReadTimestamp *int64 `json:"readTimestamp,omitempty" xml:"readTimestamp,omitempty"`
	// 工号
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchOTOQueryResponseBodyMessageReadInfoList) String() string {
	return tea.Prettify(s)
}

func (s BatchOTOQueryResponseBodyMessageReadInfoList) GoString() string {
	return s.String()
}

func (s *BatchOTOQueryResponseBodyMessageReadInfoList) SetName(v string) *BatchOTOQueryResponseBodyMessageReadInfoList {
	s.Name = &v
	return s
}

func (s *BatchOTOQueryResponseBodyMessageReadInfoList) SetReadStatus(v string) *BatchOTOQueryResponseBodyMessageReadInfoList {
	s.ReadStatus = &v
	return s
}

func (s *BatchOTOQueryResponseBodyMessageReadInfoList) SetReadTimestamp(v int64) *BatchOTOQueryResponseBodyMessageReadInfoList {
	s.ReadTimestamp = &v
	return s
}

func (s *BatchOTOQueryResponseBodyMessageReadInfoList) SetUserId(v string) *BatchOTOQueryResponseBodyMessageReadInfoList {
	s.UserId = &v
	return s
}

type BatchOTOQueryResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchOTOQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchOTOQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchOTOQueryResponse) GoString() string {
	return s.String()
}

func (s *BatchOTOQueryResponse) SetHeaders(v map[string]*string) *BatchOTOQueryResponse {
	s.Headers = v
	return s
}

func (s *BatchOTOQueryResponse) SetBody(v *BatchOTOQueryResponseBody) *BatchOTOQueryResponse {
	s.Body = v
	return s
}

type BatchRecallGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchRecallGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchRecallGroupHeaders) GoString() string {
	return s.String()
}

func (s *BatchRecallGroupHeaders) SetCommonHeaders(v map[string]*string) *BatchRecallGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchRecallGroupHeaders) SetXAcsDingtalkAccessToken(v string) *BatchRecallGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchRecallGroupRequest struct {
	// 机器人的robotCode
	ChatbotId *string `json:"chatbotId,omitempty" xml:"chatbotId,omitempty"`
	// 开放的群id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 消息id
	ProcessQueryKeys []*string `json:"processQueryKeys,omitempty" xml:"processQueryKeys,omitempty" type:"Repeated"`
}

func (s BatchRecallGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchRecallGroupRequest) GoString() string {
	return s.String()
}

func (s *BatchRecallGroupRequest) SetChatbotId(v string) *BatchRecallGroupRequest {
	s.ChatbotId = &v
	return s
}

func (s *BatchRecallGroupRequest) SetOpenConversationId(v string) *BatchRecallGroupRequest {
	s.OpenConversationId = &v
	return s
}

func (s *BatchRecallGroupRequest) SetProcessQueryKeys(v []*string) *BatchRecallGroupRequest {
	s.ProcessQueryKeys = v
	return s
}

type BatchRecallGroupResponseBody struct {
	// 撤回失败的消息id及原因
	FailedResult map[string]*string `json:"failedResult,omitempty" xml:"failedResult,omitempty"`
	// 撤回成功的消息id
	SuccessResult []*string `json:"successResult,omitempty" xml:"successResult,omitempty" type:"Repeated"`
}

func (s BatchRecallGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchRecallGroupResponseBody) GoString() string {
	return s.String()
}

func (s *BatchRecallGroupResponseBody) SetFailedResult(v map[string]*string) *BatchRecallGroupResponseBody {
	s.FailedResult = v
	return s
}

func (s *BatchRecallGroupResponseBody) SetSuccessResult(v []*string) *BatchRecallGroupResponseBody {
	s.SuccessResult = v
	return s
}

type BatchRecallGroupResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchRecallGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchRecallGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchRecallGroupResponse) GoString() string {
	return s.String()
}

func (s *BatchRecallGroupResponse) SetHeaders(v map[string]*string) *BatchRecallGroupResponse {
	s.Headers = v
	return s
}

func (s *BatchRecallGroupResponse) SetBody(v *BatchRecallGroupResponseBody) *BatchRecallGroupResponse {
	s.Body = v
	return s
}

type BatchRecallOTOHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchRecallOTOHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchRecallOTOHeaders) GoString() string {
	return s.String()
}

func (s *BatchRecallOTOHeaders) SetCommonHeaders(v map[string]*string) *BatchRecallOTOHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchRecallOTOHeaders) SetXAcsDingtalkAccessToken(v string) *BatchRecallOTOHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchRecallOTORequest struct {
	// 消息id
	ProcessQueryKeys []*string `json:"processQueryKeys,omitempty" xml:"processQueryKeys,omitempty" type:"Repeated"`
	// 机器人的robotCode
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s BatchRecallOTORequest) String() string {
	return tea.Prettify(s)
}

func (s BatchRecallOTORequest) GoString() string {
	return s.String()
}

func (s *BatchRecallOTORequest) SetProcessQueryKeys(v []*string) *BatchRecallOTORequest {
	s.ProcessQueryKeys = v
	return s
}

func (s *BatchRecallOTORequest) SetRobotCode(v string) *BatchRecallOTORequest {
	s.RobotCode = &v
	return s
}

type BatchRecallOTOResponseBody struct {
	// 撤回失败的消息id及对应的失败原因
	FailedResult map[string]*string `json:"failedResult,omitempty" xml:"failedResult,omitempty"`
	// 撤回成功的消息id
	SuccessResult []*string `json:"successResult,omitempty" xml:"successResult,omitempty" type:"Repeated"`
}

func (s BatchRecallOTOResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchRecallOTOResponseBody) GoString() string {
	return s.String()
}

func (s *BatchRecallOTOResponseBody) SetFailedResult(v map[string]*string) *BatchRecallOTOResponseBody {
	s.FailedResult = v
	return s
}

func (s *BatchRecallOTOResponseBody) SetSuccessResult(v []*string) *BatchRecallOTOResponseBody {
	s.SuccessResult = v
	return s
}

type BatchRecallOTOResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchRecallOTOResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchRecallOTOResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchRecallOTOResponse) GoString() string {
	return s.String()
}

func (s *BatchRecallOTOResponse) SetHeaders(v map[string]*string) *BatchRecallOTOResponse {
	s.Headers = v
	return s
}

func (s *BatchRecallOTOResponse) SetBody(v *BatchRecallOTOResponseBody) *BatchRecallOTOResponse {
	s.Body = v
	return s
}

type BatchSendOTOHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchSendOTOHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOTOHeaders) GoString() string {
	return s.String()
}

func (s *BatchSendOTOHeaders) SetCommonHeaders(v map[string]*string) *BatchSendOTOHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchSendOTOHeaders) SetXAcsDingtalkAccessToken(v string) *BatchSendOTOHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchSendOTORequest struct {
	// 消息的msgKey
	MsgKey *string `json:"msgKey,omitempty" xml:"msgKey,omitempty"`
	// 消息体
	MsgParam *string `json:"msgParam,omitempty" xml:"msgParam,omitempty"`
	// 机器人的robotCode
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	// 被推送会话人员的userId列表
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s BatchSendOTORequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOTORequest) GoString() string {
	return s.String()
}

func (s *BatchSendOTORequest) SetMsgKey(v string) *BatchSendOTORequest {
	s.MsgKey = &v
	return s
}

func (s *BatchSendOTORequest) SetMsgParam(v string) *BatchSendOTORequest {
	s.MsgParam = &v
	return s
}

func (s *BatchSendOTORequest) SetRobotCode(v string) *BatchSendOTORequest {
	s.RobotCode = &v
	return s
}

func (s *BatchSendOTORequest) SetUserIds(v []*string) *BatchSendOTORequest {
	s.UserIds = v
	return s
}

type BatchSendOTOResponseBody struct {
	// 推送频繁，被限流的用户userId列表
	FlowControlledStaffIdList []*string `json:"flowControlledStaffIdList,omitempty" xml:"flowControlledStaffIdList,omitempty" type:"Repeated"`
	// 无效的用户userId列表
	InvalidStaffIdList []*string `json:"invalidStaffIdList,omitempty" xml:"invalidStaffIdList,omitempty" type:"Repeated"`
	// 消息id
	ProcessQueryKey *string `json:"processQueryKey,omitempty" xml:"processQueryKey,omitempty"`
}

func (s BatchSendOTOResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOTOResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSendOTOResponseBody) SetFlowControlledStaffIdList(v []*string) *BatchSendOTOResponseBody {
	s.FlowControlledStaffIdList = v
	return s
}

func (s *BatchSendOTOResponseBody) SetInvalidStaffIdList(v []*string) *BatchSendOTOResponseBody {
	s.InvalidStaffIdList = v
	return s
}

func (s *BatchSendOTOResponseBody) SetProcessQueryKey(v string) *BatchSendOTOResponseBody {
	s.ProcessQueryKey = &v
	return s
}

type BatchSendOTOResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchSendOTOResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchSendOTOResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOTOResponse) GoString() string {
	return s.String()
}

func (s *BatchSendOTOResponse) SetHeaders(v map[string]*string) *BatchSendOTOResponse {
	s.Headers = v
	return s
}

func (s *BatchSendOTOResponse) SetBody(v *BatchSendOTOResponseBody) *BatchSendOTOResponse {
	s.Body = v
	return s
}

type ClearRobotPluginHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ClearRobotPluginHeaders) String() string {
	return tea.Prettify(s)
}

func (s ClearRobotPluginHeaders) GoString() string {
	return s.String()
}

func (s *ClearRobotPluginHeaders) SetCommonHeaders(v map[string]*string) *ClearRobotPluginHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ClearRobotPluginHeaders) SetXAcsDingtalkAccessToken(v string) *ClearRobotPluginHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ClearRobotPluginRequest struct {
	// 钉钉开放平台后台机器人的robotCode
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s ClearRobotPluginRequest) String() string {
	return tea.Prettify(s)
}

func (s ClearRobotPluginRequest) GoString() string {
	return s.String()
}

func (s *ClearRobotPluginRequest) SetRobotCode(v string) *ClearRobotPluginRequest {
	s.RobotCode = &v
	return s
}

type ClearRobotPluginResponseBody struct {
	// 是否成功清除机器人快捷入口
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ClearRobotPluginResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClearRobotPluginResponseBody) GoString() string {
	return s.String()
}

func (s *ClearRobotPluginResponseBody) SetResult(v bool) *ClearRobotPluginResponseBody {
	s.Result = &v
	return s
}

type ClearRobotPluginResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ClearRobotPluginResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ClearRobotPluginResponse) String() string {
	return tea.Prettify(s)
}

func (s ClearRobotPluginResponse) GoString() string {
	return s.String()
}

func (s *ClearRobotPluginResponse) SetHeaders(v map[string]*string) *ClearRobotPluginResponse {
	s.Headers = v
	return s
}

func (s *ClearRobotPluginResponse) SetBody(v *ClearRobotPluginResponseBody) *ClearRobotPluginResponse {
	s.Body = v
	return s
}

type ManageSingleChatRobotStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ManageSingleChatRobotStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s ManageSingleChatRobotStatusHeaders) GoString() string {
	return s.String()
}

func (s *ManageSingleChatRobotStatusHeaders) SetCommonHeaders(v map[string]*string) *ManageSingleChatRobotStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ManageSingleChatRobotStatusHeaders) SetXAcsDingtalkAccessToken(v string) *ManageSingleChatRobotStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ManageSingleChatRobotStatusRequest struct {
	// 钉钉开放平台后台机器人的robotCode
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	// 机器人的可用状态，enable-启用、disable-停用
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ManageSingleChatRobotStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ManageSingleChatRobotStatusRequest) GoString() string {
	return s.String()
}

func (s *ManageSingleChatRobotStatusRequest) SetRobotCode(v string) *ManageSingleChatRobotStatusRequest {
	s.RobotCode = &v
	return s
}

func (s *ManageSingleChatRobotStatusRequest) SetStatus(v string) *ManageSingleChatRobotStatusRequest {
	s.Status = &v
	return s
}

type ManageSingleChatRobotStatusResponseBody struct {
	// 是否成功更新机器人状态
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ManageSingleChatRobotStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ManageSingleChatRobotStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ManageSingleChatRobotStatusResponseBody) SetResult(v bool) *ManageSingleChatRobotStatusResponseBody {
	s.Result = &v
	return s
}

type ManageSingleChatRobotStatusResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ManageSingleChatRobotStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ManageSingleChatRobotStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ManageSingleChatRobotStatusResponse) GoString() string {
	return s.String()
}

func (s *ManageSingleChatRobotStatusResponse) SetHeaders(v map[string]*string) *ManageSingleChatRobotStatusResponse {
	s.Headers = v
	return s
}

func (s *ManageSingleChatRobotStatusResponse) SetBody(v *ManageSingleChatRobotStatusResponseBody) *ManageSingleChatRobotStatusResponse {
	s.Body = v
	return s
}

type OrgGroupQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s OrgGroupQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s OrgGroupQueryHeaders) GoString() string {
	return s.String()
}

func (s *OrgGroupQueryHeaders) SetCommonHeaders(v map[string]*string) *OrgGroupQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OrgGroupQueryHeaders) SetXAcsDingtalkAccessToken(v string) *OrgGroupQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type OrgGroupQueryRequest struct {
	// 分页查询每页的数量
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 一次查询后返回的加密的分页凭证，首次查询不填
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 开放的群id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 发送消息返回的加密消息id
	ProcessQueryKey *string `json:"processQueryKey,omitempty" xml:"processQueryKey,omitempty"`
	// 企业机器人的robotcode
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	// 群内机器人的webhook中的Token
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s OrgGroupQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s OrgGroupQueryRequest) GoString() string {
	return s.String()
}

func (s *OrgGroupQueryRequest) SetMaxResults(v int64) *OrgGroupQueryRequest {
	s.MaxResults = &v
	return s
}

func (s *OrgGroupQueryRequest) SetNextToken(v string) *OrgGroupQueryRequest {
	s.NextToken = &v
	return s
}

func (s *OrgGroupQueryRequest) SetOpenConversationId(v string) *OrgGroupQueryRequest {
	s.OpenConversationId = &v
	return s
}

func (s *OrgGroupQueryRequest) SetProcessQueryKey(v string) *OrgGroupQueryRequest {
	s.ProcessQueryKey = &v
	return s
}

func (s *OrgGroupQueryRequest) SetRobotCode(v string) *OrgGroupQueryRequest {
	s.RobotCode = &v
	return s
}

func (s *OrgGroupQueryRequest) SetToken(v string) *OrgGroupQueryRequest {
	s.Token = &v
	return s
}

type OrgGroupQueryResponseBody struct {
	// 分页查询是否还有人员可查询消息已读状态
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 下次分页查询的加密凭证
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 消息已读人的userId列表
	ReadUserIds []*string `json:"readUserIds,omitempty" xml:"readUserIds,omitempty" type:"Repeated"`
	// 消息发送状态
	SendStatus *string `json:"sendStatus,omitempty" xml:"sendStatus,omitempty"`
}

func (s OrgGroupQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OrgGroupQueryResponseBody) GoString() string {
	return s.String()
}

func (s *OrgGroupQueryResponseBody) SetHasMore(v bool) *OrgGroupQueryResponseBody {
	s.HasMore = &v
	return s
}

func (s *OrgGroupQueryResponseBody) SetNextToken(v string) *OrgGroupQueryResponseBody {
	s.NextToken = &v
	return s
}

func (s *OrgGroupQueryResponseBody) SetReadUserIds(v []*string) *OrgGroupQueryResponseBody {
	s.ReadUserIds = v
	return s
}

func (s *OrgGroupQueryResponseBody) SetSendStatus(v string) *OrgGroupQueryResponseBody {
	s.SendStatus = &v
	return s
}

type OrgGroupQueryResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OrgGroupQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OrgGroupQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s OrgGroupQueryResponse) GoString() string {
	return s.String()
}

func (s *OrgGroupQueryResponse) SetHeaders(v map[string]*string) *OrgGroupQueryResponse {
	s.Headers = v
	return s
}

func (s *OrgGroupQueryResponse) SetBody(v *OrgGroupQueryResponseBody) *OrgGroupQueryResponse {
	s.Body = v
	return s
}

type OrgGroupRecallHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s OrgGroupRecallHeaders) String() string {
	return tea.Prettify(s)
}

func (s OrgGroupRecallHeaders) GoString() string {
	return s.String()
}

func (s *OrgGroupRecallHeaders) SetCommonHeaders(v map[string]*string) *OrgGroupRecallHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OrgGroupRecallHeaders) SetXAcsDingtalkAccessToken(v string) *OrgGroupRecallHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type OrgGroupRecallRequest struct {
	// 开放的群id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 消息id
	ProcessQueryKeys []*string `json:"processQueryKeys,omitempty" xml:"processQueryKeys,omitempty" type:"Repeated"`
	// 机器人的robotCode
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s OrgGroupRecallRequest) String() string {
	return tea.Prettify(s)
}

func (s OrgGroupRecallRequest) GoString() string {
	return s.String()
}

func (s *OrgGroupRecallRequest) SetOpenConversationId(v string) *OrgGroupRecallRequest {
	s.OpenConversationId = &v
	return s
}

func (s *OrgGroupRecallRequest) SetProcessQueryKeys(v []*string) *OrgGroupRecallRequest {
	s.ProcessQueryKeys = v
	return s
}

func (s *OrgGroupRecallRequest) SetRobotCode(v string) *OrgGroupRecallRequest {
	s.RobotCode = &v
	return s
}

type OrgGroupRecallResponseBody struct {
	// 撤回失败的消息id及原因
	FailedResult map[string]*string `json:"failedResult,omitempty" xml:"failedResult,omitempty"`
	// 撤回成功的消息id
	SuccessResult []*string `json:"successResult,omitempty" xml:"successResult,omitempty" type:"Repeated"`
}

func (s OrgGroupRecallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OrgGroupRecallResponseBody) GoString() string {
	return s.String()
}

func (s *OrgGroupRecallResponseBody) SetFailedResult(v map[string]*string) *OrgGroupRecallResponseBody {
	s.FailedResult = v
	return s
}

func (s *OrgGroupRecallResponseBody) SetSuccessResult(v []*string) *OrgGroupRecallResponseBody {
	s.SuccessResult = v
	return s
}

type OrgGroupRecallResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OrgGroupRecallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OrgGroupRecallResponse) String() string {
	return tea.Prettify(s)
}

func (s OrgGroupRecallResponse) GoString() string {
	return s.String()
}

func (s *OrgGroupRecallResponse) SetHeaders(v map[string]*string) *OrgGroupRecallResponse {
	s.Headers = v
	return s
}

func (s *OrgGroupRecallResponse) SetBody(v *OrgGroupRecallResponseBody) *OrgGroupRecallResponse {
	s.Body = v
	return s
}

type OrgGroupSendHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s OrgGroupSendHeaders) String() string {
	return tea.Prettify(s)
}

func (s OrgGroupSendHeaders) GoString() string {
	return s.String()
}

func (s *OrgGroupSendHeaders) SetCommonHeaders(v map[string]*string) *OrgGroupSendHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OrgGroupSendHeaders) SetXAcsDingtalkAccessToken(v string) *OrgGroupSendHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type OrgGroupSendRequest struct {
	// 酷应用的code
	CoolAppCode *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	// 消息类型的key
	MsgKey *string `json:"msgKey,omitempty" xml:"msgKey,omitempty"`
	// 消息体
	MsgParam *string `json:"msgParam,omitempty" xml:"msgParam,omitempty"`
	// 开放的群id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 群内机器人的code
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	// 群内机器人的webhook中的Token
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s OrgGroupSendRequest) String() string {
	return tea.Prettify(s)
}

func (s OrgGroupSendRequest) GoString() string {
	return s.String()
}

func (s *OrgGroupSendRequest) SetCoolAppCode(v string) *OrgGroupSendRequest {
	s.CoolAppCode = &v
	return s
}

func (s *OrgGroupSendRequest) SetMsgKey(v string) *OrgGroupSendRequest {
	s.MsgKey = &v
	return s
}

func (s *OrgGroupSendRequest) SetMsgParam(v string) *OrgGroupSendRequest {
	s.MsgParam = &v
	return s
}

func (s *OrgGroupSendRequest) SetOpenConversationId(v string) *OrgGroupSendRequest {
	s.OpenConversationId = &v
	return s
}

func (s *OrgGroupSendRequest) SetRobotCode(v string) *OrgGroupSendRequest {
	s.RobotCode = &v
	return s
}

func (s *OrgGroupSendRequest) SetToken(v string) *OrgGroupSendRequest {
	s.Token = &v
	return s
}

type OrgGroupSendResponseBody struct {
	// 加密消息id
	ProcessQueryKey *string `json:"processQueryKey,omitempty" xml:"processQueryKey,omitempty"`
}

func (s OrgGroupSendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OrgGroupSendResponseBody) GoString() string {
	return s.String()
}

func (s *OrgGroupSendResponseBody) SetProcessQueryKey(v string) *OrgGroupSendResponseBody {
	s.ProcessQueryKey = &v
	return s
}

type OrgGroupSendResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OrgGroupSendResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OrgGroupSendResponse) String() string {
	return tea.Prettify(s)
}

func (s OrgGroupSendResponse) GoString() string {
	return s.String()
}

func (s *OrgGroupSendResponse) SetHeaders(v map[string]*string) *OrgGroupSendResponse {
	s.Headers = v
	return s
}

func (s *OrgGroupSendResponse) SetBody(v *OrgGroupSendResponseBody) *OrgGroupSendResponse {
	s.Body = v
	return s
}

type QueryRobotPluginHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryRobotPluginHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotPluginHeaders) GoString() string {
	return s.String()
}

func (s *QueryRobotPluginHeaders) SetCommonHeaders(v map[string]*string) *QueryRobotPluginHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryRobotPluginHeaders) SetXAcsDingtalkAccessToken(v string) *QueryRobotPluginHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryRobotPluginRequest struct {
	// 钉钉开放平台后台机器人的robotCode
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s QueryRobotPluginRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotPluginRequest) GoString() string {
	return s.String()
}

func (s *QueryRobotPluginRequest) SetRobotCode(v string) *QueryRobotPluginRequest {
	s.RobotCode = &v
	return s
}

type QueryRobotPluginResponseBody struct {
	PluginInfoList []*QueryRobotPluginResponseBodyPluginInfoList `json:"pluginInfoList,omitempty" xml:"pluginInfoList,omitempty" type:"Repeated"`
}

func (s QueryRobotPluginResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotPluginResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRobotPluginResponseBody) SetPluginInfoList(v []*QueryRobotPluginResponseBodyPluginInfoList) *QueryRobotPluginResponseBody {
	s.PluginInfoList = v
	return s
}

type QueryRobotPluginResponseBodyPluginInfoList struct {
	// 快捷入口的图标id
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 手机端快捷入口跳转链接
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	// 快捷入口的名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// pc端会话快捷入口跳转链接
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
}

func (s QueryRobotPluginResponseBodyPluginInfoList) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotPluginResponseBodyPluginInfoList) GoString() string {
	return s.String()
}

func (s *QueryRobotPluginResponseBodyPluginInfoList) SetIcon(v string) *QueryRobotPluginResponseBodyPluginInfoList {
	s.Icon = &v
	return s
}

func (s *QueryRobotPluginResponseBodyPluginInfoList) SetMobileUrl(v string) *QueryRobotPluginResponseBodyPluginInfoList {
	s.MobileUrl = &v
	return s
}

func (s *QueryRobotPluginResponseBodyPluginInfoList) SetName(v string) *QueryRobotPluginResponseBodyPluginInfoList {
	s.Name = &v
	return s
}

func (s *QueryRobotPluginResponseBodyPluginInfoList) SetPcUrl(v string) *QueryRobotPluginResponseBodyPluginInfoList {
	s.PcUrl = &v
	return s
}

type QueryRobotPluginResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRobotPluginResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRobotPluginResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotPluginResponse) GoString() string {
	return s.String()
}

func (s *QueryRobotPluginResponse) SetHeaders(v map[string]*string) *QueryRobotPluginResponse {
	s.Headers = v
	return s
}

func (s *QueryRobotPluginResponse) SetBody(v *QueryRobotPluginResponseBody) *QueryRobotPluginResponse {
	s.Body = v
	return s
}

type SendRobotDingMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendRobotDingMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendRobotDingMessageHeaders) GoString() string {
	return s.String()
}

func (s *SendRobotDingMessageHeaders) SetCommonHeaders(v map[string]*string) *SendRobotDingMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendRobotDingMessageHeaders) SetXAcsDingtalkAccessToken(v string) *SendRobotDingMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendRobotDingMessageRequest struct {
	// 模版对应的参数
	ContentParams map[string]*string `json:"contentParams,omitempty" xml:"contentParams,omitempty"`
	// 颁发的模版id，可通过宜搭申请：https://yida.alibaba-inc.com/alibaba/web/APP_NSUGAGIQUMI4ESRA7O7D/inst/homepage/#/FORM-WO866371VGXSECXX4M0NC9KSGAT92VSA3TZSK9B
	DingTemplateId *string `json:"dingTemplateId,omitempty" xml:"dingTemplateId,omitempty"`
	// 群聊的对外开放Id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 接受人的userId列表
	ReceiverUserIdList []*string `json:"receiverUserIdList,omitempty" xml:"receiverUserIdList,omitempty" type:"Repeated"`
	// 机器人的Id
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s SendRobotDingMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendRobotDingMessageRequest) GoString() string {
	return s.String()
}

func (s *SendRobotDingMessageRequest) SetContentParams(v map[string]*string) *SendRobotDingMessageRequest {
	s.ContentParams = v
	return s
}

func (s *SendRobotDingMessageRequest) SetDingTemplateId(v string) *SendRobotDingMessageRequest {
	s.DingTemplateId = &v
	return s
}

func (s *SendRobotDingMessageRequest) SetOpenConversationId(v string) *SendRobotDingMessageRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SendRobotDingMessageRequest) SetReceiverUserIdList(v []*string) *SendRobotDingMessageRequest {
	s.ReceiverUserIdList = v
	return s
}

func (s *SendRobotDingMessageRequest) SetRobotCode(v string) *SendRobotDingMessageRequest {
	s.RobotCode = &v
	return s
}

type SendRobotDingMessageResponseBody struct {
	// 返回的ding消息id
	DingSendResultId *string `json:"dingSendResultId,omitempty" xml:"dingSendResultId,omitempty"`
}

func (s SendRobotDingMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendRobotDingMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendRobotDingMessageResponseBody) SetDingSendResultId(v string) *SendRobotDingMessageResponseBody {
	s.DingSendResultId = &v
	return s
}

type SendRobotDingMessageResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendRobotDingMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendRobotDingMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendRobotDingMessageResponse) GoString() string {
	return s.String()
}

func (s *SendRobotDingMessageResponse) SetHeaders(v map[string]*string) *SendRobotDingMessageResponse {
	s.Headers = v
	return s
}

func (s *SendRobotDingMessageResponse) SetBody(v *SendRobotDingMessageResponseBody) *SendRobotDingMessageResponse {
	s.Body = v
	return s
}

type SetRobotPluginHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetRobotPluginHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetRobotPluginHeaders) GoString() string {
	return s.String()
}

func (s *SetRobotPluginHeaders) SetCommonHeaders(v map[string]*string) *SetRobotPluginHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetRobotPluginHeaders) SetXAcsDingtalkAccessToken(v string) *SetRobotPluginHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetRobotPluginRequest struct {
	PluginInfoList []*SetRobotPluginRequestPluginInfoList `json:"pluginInfoList,omitempty" xml:"pluginInfoList,omitempty" type:"Repeated"`
	// 钉钉开放平台后台机器人的robotCode
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s SetRobotPluginRequest) String() string {
	return tea.Prettify(s)
}

func (s SetRobotPluginRequest) GoString() string {
	return s.String()
}

func (s *SetRobotPluginRequest) SetPluginInfoList(v []*SetRobotPluginRequestPluginInfoList) *SetRobotPluginRequest {
	s.PluginInfoList = v
	return s
}

func (s *SetRobotPluginRequest) SetRobotCode(v string) *SetRobotPluginRequest {
	s.RobotCode = &v
	return s
}

type SetRobotPluginRequestPluginInfoList struct {
	// 快捷入口的图标id
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 手机端快捷入口跳转链接
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	// 快捷入口的名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// pc端会话快捷入口跳转链接
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
}

func (s SetRobotPluginRequestPluginInfoList) String() string {
	return tea.Prettify(s)
}

func (s SetRobotPluginRequestPluginInfoList) GoString() string {
	return s.String()
}

func (s *SetRobotPluginRequestPluginInfoList) SetIcon(v string) *SetRobotPluginRequestPluginInfoList {
	s.Icon = &v
	return s
}

func (s *SetRobotPluginRequestPluginInfoList) SetMobileUrl(v string) *SetRobotPluginRequestPluginInfoList {
	s.MobileUrl = &v
	return s
}

func (s *SetRobotPluginRequestPluginInfoList) SetName(v string) *SetRobotPluginRequestPluginInfoList {
	s.Name = &v
	return s
}

func (s *SetRobotPluginRequestPluginInfoList) SetPcUrl(v string) *SetRobotPluginRequestPluginInfoList {
	s.PcUrl = &v
	return s
}

type SetRobotPluginResponseBody struct {
	// 是否成功设置机器人插件
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SetRobotPluginResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetRobotPluginResponseBody) GoString() string {
	return s.String()
}

func (s *SetRobotPluginResponseBody) SetResult(v bool) *SetRobotPluginResponseBody {
	s.Result = &v
	return s
}

type SetRobotPluginResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetRobotPluginResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetRobotPluginResponse) String() string {
	return tea.Prettify(s)
}

func (s SetRobotPluginResponse) GoString() string {
	return s.String()
}

func (s *SetRobotPluginResponse) SetHeaders(v map[string]*string) *SetRobotPluginResponse {
	s.Headers = v
	return s
}

func (s *SetRobotPluginResponse) SetBody(v *SetRobotPluginResponseBody) *SetRobotPluginResponse {
	s.Body = v
	return s
}

type UpdateInstalledRobotHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateInstalledRobotHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstalledRobotHeaders) GoString() string {
	return s.String()
}

func (s *UpdateInstalledRobotHeaders) SetCommonHeaders(v map[string]*string) *UpdateInstalledRobotHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateInstalledRobotHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateInstalledRobotHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateInstalledRobotRequest struct {
	// 机器人的简要描述。
	Brief *string `json:"brief,omitempty" xml:"brief,omitempty"`
	// 机器人的详细描述。
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 机器人图标的mediaId。
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 机器人的名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 机器人的robotCode。
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	// 更新名字或头像时是否更新群里已添加机器人的名字或头像。
	// 0-不更新群里机器人名字或头像
	// 1-更新群里机器人名字或头像
	UpdateType *int32 `json:"updateType,omitempty" xml:"updateType,omitempty"`
}

func (s UpdateInstalledRobotRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstalledRobotRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstalledRobotRequest) SetBrief(v string) *UpdateInstalledRobotRequest {
	s.Brief = &v
	return s
}

func (s *UpdateInstalledRobotRequest) SetDescription(v string) *UpdateInstalledRobotRequest {
	s.Description = &v
	return s
}

func (s *UpdateInstalledRobotRequest) SetIcon(v string) *UpdateInstalledRobotRequest {
	s.Icon = &v
	return s
}

func (s *UpdateInstalledRobotRequest) SetName(v string) *UpdateInstalledRobotRequest {
	s.Name = &v
	return s
}

func (s *UpdateInstalledRobotRequest) SetRobotCode(v string) *UpdateInstalledRobotRequest {
	s.RobotCode = &v
	return s
}

func (s *UpdateInstalledRobotRequest) SetUpdateType(v int32) *UpdateInstalledRobotRequest {
	s.UpdateType = &v
	return s
}

type UpdateInstalledRobotResponseBody struct {
	// 本次更新操作是否成功。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateInstalledRobotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstalledRobotResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstalledRobotResponseBody) SetSuccess(v bool) *UpdateInstalledRobotResponseBody {
	s.Success = &v
	return s
}

type UpdateInstalledRobotResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInstalledRobotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInstalledRobotResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstalledRobotResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstalledRobotResponse) SetHeaders(v map[string]*string) *UpdateInstalledRobotResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstalledRobotResponse) SetBody(v *UpdateInstalledRobotResponseBody) *UpdateInstalledRobotResponse {
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

func (client *Client) BatchOTOQuery(request *BatchOTOQueryRequest) (_result *BatchOTOQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchOTOQueryHeaders{}
	_result = &BatchOTOQueryResponse{}
	_body, _err := client.BatchOTOQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchOTOQueryWithOptions(request *BatchOTOQueryRequest, headers *BatchOTOQueryHeaders, runtime *util.RuntimeOptions) (_result *BatchOTOQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProcessQueryKey)) {
		query["processQueryKey"] = request.ProcessQueryKey
	}

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		query["robotCode"] = request.RobotCode
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
	_result = &BatchOTOQueryResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchOTOQuery"), tea.String("robot_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/robot/oToMessages/readStatus"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchRecallGroup(request *BatchRecallGroupRequest) (_result *BatchRecallGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchRecallGroupHeaders{}
	_result = &BatchRecallGroupResponse{}
	_body, _err := client.BatchRecallGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchRecallGroupWithOptions(request *BatchRecallGroupRequest, headers *BatchRecallGroupHeaders, runtime *util.RuntimeOptions) (_result *BatchRecallGroupResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ProcessQueryKeys)) {
		body["processQueryKeys"] = request.ProcessQueryKeys
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
	_result = &BatchRecallGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchRecallGroup"), tea.String("robot_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/robot/groupMessages/batchRecall"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchRecallOTO(request *BatchRecallOTORequest) (_result *BatchRecallOTOResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchRecallOTOHeaders{}
	_result = &BatchRecallOTOResponse{}
	_body, _err := client.BatchRecallOTOWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchRecallOTOWithOptions(request *BatchRecallOTORequest, headers *BatchRecallOTOHeaders, runtime *util.RuntimeOptions) (_result *BatchRecallOTOResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProcessQueryKeys)) {
		body["processQueryKeys"] = request.ProcessQueryKeys
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
	_result = &BatchRecallOTOResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchRecallOTO"), tea.String("robot_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/robot/otoMessages/batchRecall"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchSendOTO(request *BatchSendOTORequest) (_result *BatchSendOTOResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchSendOTOHeaders{}
	_result = &BatchSendOTOResponse{}
	_body, _err := client.BatchSendOTOWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchSendOTOWithOptions(request *BatchSendOTORequest, headers *BatchSendOTOHeaders, runtime *util.RuntimeOptions) (_result *BatchSendOTOResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
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
	_result = &BatchSendOTOResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchSendOTO"), tea.String("robot_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/robot/oToMessages/batchSend"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClearRobotPlugin(request *ClearRobotPluginRequest) (_result *ClearRobotPluginResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ClearRobotPluginHeaders{}
	_result = &ClearRobotPluginResponse{}
	_body, _err := client.ClearRobotPluginWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClearRobotPluginWithOptions(request *ClearRobotPluginRequest, headers *ClearRobotPluginHeaders, runtime *util.RuntimeOptions) (_result *ClearRobotPluginResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
	_result = &ClearRobotPluginResponse{}
	_body, _err := client.DoROARequest(tea.String("ClearRobotPlugin"), tea.String("robot_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/robot/plugins/clear"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ManageSingleChatRobotStatus(request *ManageSingleChatRobotStatusRequest) (_result *ManageSingleChatRobotStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ManageSingleChatRobotStatusHeaders{}
	_result = &ManageSingleChatRobotStatusResponse{}
	_body, _err := client.ManageSingleChatRobotStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ManageSingleChatRobotStatusWithOptions(request *ManageSingleChatRobotStatusRequest, headers *ManageSingleChatRobotStatusHeaders, runtime *util.RuntimeOptions) (_result *ManageSingleChatRobotStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
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
	_result = &ManageSingleChatRobotStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("ManageSingleChatRobotStatus"), tea.String("robot_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/robot/statuses/manage"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OrgGroupQuery(request *OrgGroupQueryRequest) (_result *OrgGroupQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OrgGroupQueryHeaders{}
	_result = &OrgGroupQueryResponse{}
	_body, _err := client.OrgGroupQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OrgGroupQueryWithOptions(request *OrgGroupQueryRequest, headers *OrgGroupQueryHeaders, runtime *util.RuntimeOptions) (_result *OrgGroupQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessQueryKey)) {
		body["processQueryKey"] = request.ProcessQueryKey
	}

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
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
	_result = &OrgGroupQueryResponse{}
	_body, _err := client.DoROARequest(tea.String("OrgGroupQuery"), tea.String("robot_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/robot/groupMessages/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OrgGroupRecall(request *OrgGroupRecallRequest) (_result *OrgGroupRecallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OrgGroupRecallHeaders{}
	_result = &OrgGroupRecallResponse{}
	_body, _err := client.OrgGroupRecallWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OrgGroupRecallWithOptions(request *OrgGroupRecallRequest, headers *OrgGroupRecallHeaders, runtime *util.RuntimeOptions) (_result *OrgGroupRecallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessQueryKeys)) {
		body["processQueryKeys"] = request.ProcessQueryKeys
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
	_result = &OrgGroupRecallResponse{}
	_body, _err := client.DoROARequest(tea.String("OrgGroupRecall"), tea.String("robot_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/robot/groupMessages/recall"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OrgGroupSend(request *OrgGroupSendRequest) (_result *OrgGroupSendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OrgGroupSendHeaders{}
	_result = &OrgGroupSendResponse{}
	_body, _err := client.OrgGroupSendWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OrgGroupSendWithOptions(request *OrgGroupSendRequest, headers *OrgGroupSendHeaders, runtime *util.RuntimeOptions) (_result *OrgGroupSendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CoolAppCode)) {
		body["coolAppCode"] = request.CoolAppCode
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

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
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
	_result = &OrgGroupSendResponse{}
	_body, _err := client.DoROARequest(tea.String("OrgGroupSend"), tea.String("robot_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/robot/groupMessages/send"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRobotPlugin(request *QueryRobotPluginRequest) (_result *QueryRobotPluginResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryRobotPluginHeaders{}
	_result = &QueryRobotPluginResponse{}
	_body, _err := client.QueryRobotPluginWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRobotPluginWithOptions(request *QueryRobotPluginRequest, headers *QueryRobotPluginHeaders, runtime *util.RuntimeOptions) (_result *QueryRobotPluginResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
	_result = &QueryRobotPluginResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryRobotPlugin"), tea.String("robot_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/robot/plugins/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendRobotDingMessage(request *SendRobotDingMessageRequest) (_result *SendRobotDingMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendRobotDingMessageHeaders{}
	_result = &SendRobotDingMessageResponse{}
	_body, _err := client.SendRobotDingMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendRobotDingMessageWithOptions(request *SendRobotDingMessageRequest, headers *SendRobotDingMessageHeaders, runtime *util.RuntimeOptions) (_result *SendRobotDingMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContentParams)) {
		body["contentParams"] = request.ContentParams
	}

	if !tea.BoolValue(util.IsUnset(request.DingTemplateId)) {
		body["dingTemplateId"] = request.DingTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
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
	_result = &SendRobotDingMessageResponse{}
	_body, _err := client.DoROARequest(tea.String("SendRobotDingMessage"), tea.String("robot_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/robot/dingMessages/send"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetRobotPlugin(request *SetRobotPluginRequest) (_result *SetRobotPluginResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetRobotPluginHeaders{}
	_result = &SetRobotPluginResponse{}
	_body, _err := client.SetRobotPluginWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetRobotPluginWithOptions(request *SetRobotPluginRequest, headers *SetRobotPluginHeaders, runtime *util.RuntimeOptions) (_result *SetRobotPluginResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PluginInfoList)) {
		body["pluginInfoList"] = request.PluginInfoList
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
	_result = &SetRobotPluginResponse{}
	_body, _err := client.DoROARequest(tea.String("SetRobotPlugin"), tea.String("robot_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/robot/plugins/set"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInstalledRobot(request *UpdateInstalledRobotRequest) (_result *UpdateInstalledRobotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateInstalledRobotHeaders{}
	_result = &UpdateInstalledRobotResponse{}
	_body, _err := client.UpdateInstalledRobotWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInstalledRobotWithOptions(request *UpdateInstalledRobotRequest, headers *UpdateInstalledRobotHeaders, runtime *util.RuntimeOptions) (_result *UpdateInstalledRobotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Brief)) {
		body["brief"] = request.Brief
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateType)) {
		body["updateType"] = request.UpdateType
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
	_result = &UpdateInstalledRobotResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateInstalledRobot"), tea.String("robot_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/robot/managements/infos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
