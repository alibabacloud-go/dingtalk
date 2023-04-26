// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package robot_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
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
	ProcessQueryKey *string `json:"processQueryKey,omitempty" xml:"processQueryKey,omitempty"`
	RobotCode       *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
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
	MessageReadInfoList []*BatchOTOQueryResponseBodyMessageReadInfoList `json:"messageReadInfoList,omitempty" xml:"messageReadInfoList,omitempty" type:"Repeated"`
	SendStatus          *string                                         `json:"sendStatus,omitempty" xml:"sendStatus,omitempty"`
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
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	ReadStatus    *string `json:"readStatus,omitempty" xml:"readStatus,omitempty"`
	ReadTimestamp *int64  `json:"readTimestamp,omitempty" xml:"readTimestamp,omitempty"`
	UserId        *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchOTOQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *BatchOTOQueryResponse) SetStatusCode(v int32) *BatchOTOQueryResponse {
	s.StatusCode = &v
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
	ChatbotId          *string   `json:"chatbotId,omitempty" xml:"chatbotId,omitempty"`
	OpenConversationId *string   `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	ProcessQueryKeys   []*string `json:"processQueryKeys,omitempty" xml:"processQueryKeys,omitempty" type:"Repeated"`
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
	FailedResult  map[string]*string `json:"failedResult,omitempty" xml:"failedResult,omitempty"`
	SuccessResult []*string          `json:"successResult,omitempty" xml:"successResult,omitempty" type:"Repeated"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchRecallGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *BatchRecallGroupResponse) SetStatusCode(v int32) *BatchRecallGroupResponse {
	s.StatusCode = &v
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
	ProcessQueryKeys []*string `json:"processQueryKeys,omitempty" xml:"processQueryKeys,omitempty" type:"Repeated"`
	RobotCode        *string   `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
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
	FailedResult  map[string]*string `json:"failedResult,omitempty" xml:"failedResult,omitempty"`
	SuccessResult []*string          `json:"successResult,omitempty" xml:"successResult,omitempty" type:"Repeated"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchRecallOTOResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *BatchRecallOTOResponse) SetStatusCode(v int32) *BatchRecallOTOResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchRecallOTOResponse) SetBody(v *BatchRecallOTOResponseBody) *BatchRecallOTOResponse {
	s.Body = v
	return s
}

type BatchRecallPrivateChatHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchRecallPrivateChatHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchRecallPrivateChatHeaders) GoString() string {
	return s.String()
}

func (s *BatchRecallPrivateChatHeaders) SetCommonHeaders(v map[string]*string) *BatchRecallPrivateChatHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchRecallPrivateChatHeaders) SetXAcsDingtalkAccessToken(v string) *BatchRecallPrivateChatHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchRecallPrivateChatRequest struct {
	OpenConversationId *string   `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	ProcessQueryKeys   []*string `json:"processQueryKeys,omitempty" xml:"processQueryKeys,omitempty" type:"Repeated"`
	RobotCode          *string   `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s BatchRecallPrivateChatRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchRecallPrivateChatRequest) GoString() string {
	return s.String()
}

func (s *BatchRecallPrivateChatRequest) SetOpenConversationId(v string) *BatchRecallPrivateChatRequest {
	s.OpenConversationId = &v
	return s
}

func (s *BatchRecallPrivateChatRequest) SetProcessQueryKeys(v []*string) *BatchRecallPrivateChatRequest {
	s.ProcessQueryKeys = v
	return s
}

func (s *BatchRecallPrivateChatRequest) SetRobotCode(v string) *BatchRecallPrivateChatRequest {
	s.RobotCode = &v
	return s
}

type BatchRecallPrivateChatResponseBody struct {
	FailedResult  map[string]*string `json:"failedResult,omitempty" xml:"failedResult,omitempty"`
	SuccessResult []*string          `json:"successResult,omitempty" xml:"successResult,omitempty" type:"Repeated"`
}

func (s BatchRecallPrivateChatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchRecallPrivateChatResponseBody) GoString() string {
	return s.String()
}

func (s *BatchRecallPrivateChatResponseBody) SetFailedResult(v map[string]*string) *BatchRecallPrivateChatResponseBody {
	s.FailedResult = v
	return s
}

func (s *BatchRecallPrivateChatResponseBody) SetSuccessResult(v []*string) *BatchRecallPrivateChatResponseBody {
	s.SuccessResult = v
	return s
}

type BatchRecallPrivateChatResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchRecallPrivateChatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchRecallPrivateChatResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchRecallPrivateChatResponse) GoString() string {
	return s.String()
}

func (s *BatchRecallPrivateChatResponse) SetHeaders(v map[string]*string) *BatchRecallPrivateChatResponse {
	s.Headers = v
	return s
}

func (s *BatchRecallPrivateChatResponse) SetStatusCode(v int32) *BatchRecallPrivateChatResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchRecallPrivateChatResponse) SetBody(v *BatchRecallPrivateChatResponseBody) *BatchRecallPrivateChatResponse {
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
	MsgKey    *string   `json:"msgKey,omitempty" xml:"msgKey,omitempty"`
	MsgParam  *string   `json:"msgParam,omitempty" xml:"msgParam,omitempty"`
	RobotCode *string   `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	UserIds   []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
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
	FlowControlledStaffIdList []*string `json:"flowControlledStaffIdList,omitempty" xml:"flowControlledStaffIdList,omitempty" type:"Repeated"`
	InvalidStaffIdList        []*string `json:"invalidStaffIdList,omitempty" xml:"invalidStaffIdList,omitempty" type:"Repeated"`
	ProcessQueryKey           *string   `json:"processQueryKey,omitempty" xml:"processQueryKey,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchSendOTOResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *BatchSendOTOResponse) SetStatusCode(v int32) *BatchSendOTOResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ClearRobotPluginResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ClearRobotPluginResponse) SetStatusCode(v int32) *ClearRobotPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearRobotPluginResponse) SetBody(v *ClearRobotPluginResponseBody) *ClearRobotPluginResponse {
	s.Body = v
	return s
}

type GetBotListInGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetBotListInGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetBotListInGroupHeaders) GoString() string {
	return s.String()
}

func (s *GetBotListInGroupHeaders) SetCommonHeaders(v map[string]*string) *GetBotListInGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetBotListInGroupHeaders) SetXAcsDingtalkAccessToken(v string) *GetBotListInGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetBotListInGroupRequest struct {
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s GetBotListInGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBotListInGroupRequest) GoString() string {
	return s.String()
}

func (s *GetBotListInGroupRequest) SetOpenConversationId(v string) *GetBotListInGroupRequest {
	s.OpenConversationId = &v
	return s
}

type GetBotListInGroupResponseBody struct {
	ChatbotInstanceVOList []*GetBotListInGroupResponseBodyChatbotInstanceVOList `json:"chatbotInstanceVOList,omitempty" xml:"chatbotInstanceVOList,omitempty" type:"Repeated"`
}

func (s GetBotListInGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBotListInGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetBotListInGroupResponseBody) SetChatbotInstanceVOList(v []*GetBotListInGroupResponseBodyChatbotInstanceVOList) *GetBotListInGroupResponseBody {
	s.ChatbotInstanceVOList = v
	return s
}

type GetBotListInGroupResponseBodyChatbotInstanceVOList struct {
	DownloadIconURL *string `json:"downloadIconURL,omitempty" xml:"downloadIconURL,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	OpenRobotType   *int32  `json:"openRobotType,omitempty" xml:"openRobotType,omitempty"`
	RobotCode       *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s GetBotListInGroupResponseBodyChatbotInstanceVOList) String() string {
	return tea.Prettify(s)
}

func (s GetBotListInGroupResponseBodyChatbotInstanceVOList) GoString() string {
	return s.String()
}

func (s *GetBotListInGroupResponseBodyChatbotInstanceVOList) SetDownloadIconURL(v string) *GetBotListInGroupResponseBodyChatbotInstanceVOList {
	s.DownloadIconURL = &v
	return s
}

func (s *GetBotListInGroupResponseBodyChatbotInstanceVOList) SetName(v string) *GetBotListInGroupResponseBodyChatbotInstanceVOList {
	s.Name = &v
	return s
}

func (s *GetBotListInGroupResponseBodyChatbotInstanceVOList) SetOpenRobotType(v int32) *GetBotListInGroupResponseBodyChatbotInstanceVOList {
	s.OpenRobotType = &v
	return s
}

func (s *GetBotListInGroupResponseBodyChatbotInstanceVOList) SetRobotCode(v string) *GetBotListInGroupResponseBodyChatbotInstanceVOList {
	s.RobotCode = &v
	return s
}

type GetBotListInGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBotListInGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBotListInGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBotListInGroupResponse) GoString() string {
	return s.String()
}

func (s *GetBotListInGroupResponse) SetHeaders(v map[string]*string) *GetBotListInGroupResponse {
	s.Headers = v
	return s
}

func (s *GetBotListInGroupResponse) SetStatusCode(v int32) *GetBotListInGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBotListInGroupResponse) SetBody(v *GetBotListInGroupResponseBody) *GetBotListInGroupResponse {
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
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ManageSingleChatRobotStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ManageSingleChatRobotStatusResponse) SetStatusCode(v int32) *ManageSingleChatRobotStatusResponse {
	s.StatusCode = &v
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
	MaxResults         *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken          *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	ProcessQueryKey    *string `json:"processQueryKey,omitempty" xml:"processQueryKey,omitempty"`
	RobotCode          *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	Token              *string `json:"token,omitempty" xml:"token,omitempty"`
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
	HasMore     *bool     `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken   *string   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ReadUserIds []*string `json:"readUserIds,omitempty" xml:"readUserIds,omitempty" type:"Repeated"`
	SendStatus  *string   `json:"sendStatus,omitempty" xml:"sendStatus,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OrgGroupQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OrgGroupQueryResponse) SetStatusCode(v int32) *OrgGroupQueryResponse {
	s.StatusCode = &v
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
	OpenConversationId *string   `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	ProcessQueryKeys   []*string `json:"processQueryKeys,omitempty" xml:"processQueryKeys,omitempty" type:"Repeated"`
	RobotCode          *string   `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
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
	FailedResult  map[string]*string `json:"failedResult,omitempty" xml:"failedResult,omitempty"`
	SuccessResult []*string          `json:"successResult,omitempty" xml:"successResult,omitempty" type:"Repeated"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OrgGroupRecallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OrgGroupRecallResponse) SetStatusCode(v int32) *OrgGroupRecallResponse {
	s.StatusCode = &v
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
	CoolAppCode        *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	MsgKey             *string `json:"msgKey,omitempty" xml:"msgKey,omitempty"`
	MsgParam           *string `json:"msgParam,omitempty" xml:"msgParam,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	RobotCode          *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	Token              *string `json:"token,omitempty" xml:"token,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OrgGroupSendResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OrgGroupSendResponse) SetStatusCode(v int32) *OrgGroupSendResponse {
	s.StatusCode = &v
	return s
}

func (s *OrgGroupSendResponse) SetBody(v *OrgGroupSendResponseBody) *OrgGroupSendResponse {
	s.Body = v
	return s
}

type PrivateChatQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PrivateChatQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s PrivateChatQueryHeaders) GoString() string {
	return s.String()
}

func (s *PrivateChatQueryHeaders) SetCommonHeaders(v map[string]*string) *PrivateChatQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PrivateChatQueryHeaders) SetXAcsDingtalkAccessToken(v string) *PrivateChatQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PrivateChatQueryRequest struct {
	MaxResults         *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken          *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	ProcessQueryKey    *string `json:"processQueryKey,omitempty" xml:"processQueryKey,omitempty"`
	RobotCode          *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s PrivateChatQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s PrivateChatQueryRequest) GoString() string {
	return s.String()
}

func (s *PrivateChatQueryRequest) SetMaxResults(v int64) *PrivateChatQueryRequest {
	s.MaxResults = &v
	return s
}

func (s *PrivateChatQueryRequest) SetNextToken(v string) *PrivateChatQueryRequest {
	s.NextToken = &v
	return s
}

func (s *PrivateChatQueryRequest) SetOpenConversationId(v string) *PrivateChatQueryRequest {
	s.OpenConversationId = &v
	return s
}

func (s *PrivateChatQueryRequest) SetProcessQueryKey(v string) *PrivateChatQueryRequest {
	s.ProcessQueryKey = &v
	return s
}

func (s *PrivateChatQueryRequest) SetRobotCode(v string) *PrivateChatQueryRequest {
	s.RobotCode = &v
	return s
}

type PrivateChatQueryResponseBody struct {
	HasMore     *bool     `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken   *string   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ReadUserIds []*string `json:"readUserIds,omitempty" xml:"readUserIds,omitempty" type:"Repeated"`
	SendStatus  *string   `json:"sendStatus,omitempty" xml:"sendStatus,omitempty"`
}

func (s PrivateChatQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PrivateChatQueryResponseBody) GoString() string {
	return s.String()
}

func (s *PrivateChatQueryResponseBody) SetHasMore(v bool) *PrivateChatQueryResponseBody {
	s.HasMore = &v
	return s
}

func (s *PrivateChatQueryResponseBody) SetNextToken(v string) *PrivateChatQueryResponseBody {
	s.NextToken = &v
	return s
}

func (s *PrivateChatQueryResponseBody) SetReadUserIds(v []*string) *PrivateChatQueryResponseBody {
	s.ReadUserIds = v
	return s
}

func (s *PrivateChatQueryResponseBody) SetSendStatus(v string) *PrivateChatQueryResponseBody {
	s.SendStatus = &v
	return s
}

type PrivateChatQueryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PrivateChatQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PrivateChatQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s PrivateChatQueryResponse) GoString() string {
	return s.String()
}

func (s *PrivateChatQueryResponse) SetHeaders(v map[string]*string) *PrivateChatQueryResponse {
	s.Headers = v
	return s
}

func (s *PrivateChatQueryResponse) SetStatusCode(v int32) *PrivateChatQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *PrivateChatQueryResponse) SetBody(v *PrivateChatQueryResponseBody) *PrivateChatQueryResponse {
	s.Body = v
	return s
}

type PrivateChatSendHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PrivateChatSendHeaders) String() string {
	return tea.Prettify(s)
}

func (s PrivateChatSendHeaders) GoString() string {
	return s.String()
}

func (s *PrivateChatSendHeaders) SetCommonHeaders(v map[string]*string) *PrivateChatSendHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PrivateChatSendHeaders) SetXAcsDingtalkAccessToken(v string) *PrivateChatSendHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PrivateChatSendRequest struct {
	CoolAppCode        *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	MsgKey             *string `json:"msgKey,omitempty" xml:"msgKey,omitempty"`
	MsgParam           *string `json:"msgParam,omitempty" xml:"msgParam,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	RobotCode          *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s PrivateChatSendRequest) String() string {
	return tea.Prettify(s)
}

func (s PrivateChatSendRequest) GoString() string {
	return s.String()
}

func (s *PrivateChatSendRequest) SetCoolAppCode(v string) *PrivateChatSendRequest {
	s.CoolAppCode = &v
	return s
}

func (s *PrivateChatSendRequest) SetMsgKey(v string) *PrivateChatSendRequest {
	s.MsgKey = &v
	return s
}

func (s *PrivateChatSendRequest) SetMsgParam(v string) *PrivateChatSendRequest {
	s.MsgParam = &v
	return s
}

func (s *PrivateChatSendRequest) SetOpenConversationId(v string) *PrivateChatSendRequest {
	s.OpenConversationId = &v
	return s
}

func (s *PrivateChatSendRequest) SetRobotCode(v string) *PrivateChatSendRequest {
	s.RobotCode = &v
	return s
}

type PrivateChatSendResponseBody struct {
	ProcessQueryKey *string `json:"processQueryKey,omitempty" xml:"processQueryKey,omitempty"`
}

func (s PrivateChatSendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PrivateChatSendResponseBody) GoString() string {
	return s.String()
}

func (s *PrivateChatSendResponseBody) SetProcessQueryKey(v string) *PrivateChatSendResponseBody {
	s.ProcessQueryKey = &v
	return s
}

type PrivateChatSendResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PrivateChatSendResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PrivateChatSendResponse) String() string {
	return tea.Prettify(s)
}

func (s PrivateChatSendResponse) GoString() string {
	return s.String()
}

func (s *PrivateChatSendResponse) SetHeaders(v map[string]*string) *PrivateChatSendResponse {
	s.Headers = v
	return s
}

func (s *PrivateChatSendResponse) SetStatusCode(v int32) *PrivateChatSendResponse {
	s.StatusCode = &v
	return s
}

func (s *PrivateChatSendResponse) SetBody(v *PrivateChatSendResponseBody) *PrivateChatSendResponse {
	s.Body = v
	return s
}

type QueryBotInstanceInGroupInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryBotInstanceInGroupInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryBotInstanceInGroupInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryBotInstanceInGroupInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryBotInstanceInGroupInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryBotInstanceInGroupInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryBotInstanceInGroupInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryBotInstanceInGroupInfoRequest struct {
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RobotCode  *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s QueryBotInstanceInGroupInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBotInstanceInGroupInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryBotInstanceInGroupInfoRequest) SetPageNumber(v int32) *QueryBotInstanceInGroupInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryBotInstanceInGroupInfoRequest) SetPageSize(v int32) *QueryBotInstanceInGroupInfoRequest {
	s.PageSize = &v
	return s
}

func (s *QueryBotInstanceInGroupInfoRequest) SetRobotCode(v string) *QueryBotInstanceInGroupInfoRequest {
	s.RobotCode = &v
	return s
}

type QueryBotInstanceInGroupInfoResponseBody struct {
	HasMore             *bool     `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	OpenConversationIds []*string `json:"openConversationIds,omitempty" xml:"openConversationIds,omitempty" type:"Repeated"`
}

func (s QueryBotInstanceInGroupInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBotInstanceInGroupInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBotInstanceInGroupInfoResponseBody) SetHasMore(v bool) *QueryBotInstanceInGroupInfoResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryBotInstanceInGroupInfoResponseBody) SetOpenConversationIds(v []*string) *QueryBotInstanceInGroupInfoResponseBody {
	s.OpenConversationIds = v
	return s
}

type QueryBotInstanceInGroupInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryBotInstanceInGroupInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBotInstanceInGroupInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBotInstanceInGroupInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryBotInstanceInGroupInfoResponse) SetHeaders(v map[string]*string) *QueryBotInstanceInGroupInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryBotInstanceInGroupInfoResponse) SetStatusCode(v int32) *QueryBotInstanceInGroupInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBotInstanceInGroupInfoResponse) SetBody(v *QueryBotInstanceInGroupInfoResponseBody) *QueryBotInstanceInGroupInfoResponse {
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
	Icon      *string `json:"icon,omitempty" xml:"icon,omitempty"`
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	PcUrl     *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryRobotPluginResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryRobotPluginResponse) SetStatusCode(v int32) *QueryRobotPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRobotPluginResponse) SetBody(v *QueryRobotPluginResponseBody) *QueryRobotPluginResponse {
	s.Body = v
	return s
}

type RobotMessageFileDownloadHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RobotMessageFileDownloadHeaders) String() string {
	return tea.Prettify(s)
}

func (s RobotMessageFileDownloadHeaders) GoString() string {
	return s.String()
}

func (s *RobotMessageFileDownloadHeaders) SetCommonHeaders(v map[string]*string) *RobotMessageFileDownloadHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RobotMessageFileDownloadHeaders) SetXAcsDingtalkAccessToken(v string) *RobotMessageFileDownloadHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RobotMessageFileDownloadRequest struct {
	DownloadCode *string `json:"downloadCode,omitempty" xml:"downloadCode,omitempty"`
	RobotCode    *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s RobotMessageFileDownloadRequest) String() string {
	return tea.Prettify(s)
}

func (s RobotMessageFileDownloadRequest) GoString() string {
	return s.String()
}

func (s *RobotMessageFileDownloadRequest) SetDownloadCode(v string) *RobotMessageFileDownloadRequest {
	s.DownloadCode = &v
	return s
}

func (s *RobotMessageFileDownloadRequest) SetRobotCode(v string) *RobotMessageFileDownloadRequest {
	s.RobotCode = &v
	return s
}

type RobotMessageFileDownloadResponseBody struct {
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
}

func (s RobotMessageFileDownloadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RobotMessageFileDownloadResponseBody) GoString() string {
	return s.String()
}

func (s *RobotMessageFileDownloadResponseBody) SetDownloadUrl(v string) *RobotMessageFileDownloadResponseBody {
	s.DownloadUrl = &v
	return s
}

type RobotMessageFileDownloadResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RobotMessageFileDownloadResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RobotMessageFileDownloadResponse) String() string {
	return tea.Prettify(s)
}

func (s RobotMessageFileDownloadResponse) GoString() string {
	return s.String()
}

func (s *RobotMessageFileDownloadResponse) SetHeaders(v map[string]*string) *RobotMessageFileDownloadResponse {
	s.Headers = v
	return s
}

func (s *RobotMessageFileDownloadResponse) SetStatusCode(v int32) *RobotMessageFileDownloadResponse {
	s.StatusCode = &v
	return s
}

func (s *RobotMessageFileDownloadResponse) SetBody(v *RobotMessageFileDownloadResponseBody) *RobotMessageFileDownloadResponse {
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
	ContentParams      map[string]*string `json:"contentParams,omitempty" xml:"contentParams,omitempty"`
	DingTemplateId     *string            `json:"dingTemplateId,omitempty" xml:"dingTemplateId,omitempty"`
	OpenConversationId *string            `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	ReceiverUserIdList []*string          `json:"receiverUserIdList,omitempty" xml:"receiverUserIdList,omitempty" type:"Repeated"`
	RobotCode          *string            `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendRobotDingMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SendRobotDingMessageResponse) SetStatusCode(v int32) *SendRobotDingMessageResponse {
	s.StatusCode = &v
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
	RobotCode      *string                                `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
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
	Icon      *string `json:"icon,omitempty" xml:"icon,omitempty"`
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	PcUrl     *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetRobotPluginResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SetRobotPluginResponse) SetStatusCode(v int32) *SetRobotPluginResponse {
	s.StatusCode = &v
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
	Brief       *string `json:"brief,omitempty" xml:"brief,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Icon        *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	RobotCode   *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	UpdateType  *int32  `json:"updateType,omitempty" xml:"updateType,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateInstalledRobotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateInstalledRobotResponse) SetStatusCode(v int32) *UpdateInstalledRobotResponse {
	s.StatusCode = &v
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
	interfaceSPI, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = interfaceSPI
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
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
	params := &openapi.Params{
		Action:      tea.String("BatchOTOQuery"),
		Version:     tea.String("robot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/robot/oToMessages/readStatus"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchOTOQueryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
	params := &openapi.Params{
		Action:      tea.String("BatchRecallGroup"),
		Version:     tea.String("robot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/robot/groupMessages/batchRecall"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchRecallGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("BatchRecallOTO"),
		Version:     tea.String("robot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/robot/otoMessages/batchRecall"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchRecallOTOResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) BatchRecallPrivateChatWithOptions(request *BatchRecallPrivateChatRequest, headers *BatchRecallPrivateChatHeaders, runtime *util.RuntimeOptions) (_result *BatchRecallPrivateChatResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("BatchRecallPrivateChat"),
		Version:     tea.String("robot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/robot/privateChatMessages/batchRecall"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchRecallPrivateChatResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchRecallPrivateChat(request *BatchRecallPrivateChatRequest) (_result *BatchRecallPrivateChatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchRecallPrivateChatHeaders{}
	_result = &BatchRecallPrivateChatResponse{}
	_body, _err := client.BatchRecallPrivateChatWithOptions(request, headers, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("BatchSendOTO"),
		Version:     tea.String("robot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/robot/oToMessages/batchSend"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchSendOTOResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("ClearRobotPlugin"),
		Version:     tea.String("robot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/robot/plugins/clear"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ClearRobotPluginResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) GetBotListInGroupWithOptions(request *GetBotListInGroupRequest, headers *GetBotListInGroupHeaders, runtime *util.RuntimeOptions) (_result *GetBotListInGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
	params := &openapi.Params{
		Action:      tea.String("GetBotListInGroup"),
		Version:     tea.String("robot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/robot/groups/robots/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBotListInGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBotListInGroup(request *GetBotListInGroupRequest) (_result *GetBotListInGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetBotListInGroupHeaders{}
	_result = &GetBotListInGroupResponse{}
	_body, _err := client.GetBotListInGroupWithOptions(request, headers, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("ManageSingleChatRobotStatus"),
		Version:     tea.String("robot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/robot/statuses/manage"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ManageSingleChatRobotStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("OrgGroupQuery"),
		Version:     tea.String("robot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/robot/groupMessages/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &OrgGroupQueryResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("OrgGroupRecall"),
		Version:     tea.String("robot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/robot/groupMessages/recall"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &OrgGroupRecallResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("OrgGroupSend"),
		Version:     tea.String("robot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/robot/groupMessages/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &OrgGroupSendResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) PrivateChatQueryWithOptions(request *PrivateChatQueryRequest, headers *PrivateChatQueryHeaders, runtime *util.RuntimeOptions) (_result *PrivateChatQueryResponse, _err error) {
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
		Action:      tea.String("PrivateChatQuery"),
		Version:     tea.String("robot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/robot/privateChatMessages/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PrivateChatQueryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PrivateChatQuery(request *PrivateChatQueryRequest) (_result *PrivateChatQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PrivateChatQueryHeaders{}
	_result = &PrivateChatQueryResponse{}
	_body, _err := client.PrivateChatQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PrivateChatSendWithOptions(request *PrivateChatSendRequest, headers *PrivateChatSendHeaders, runtime *util.RuntimeOptions) (_result *PrivateChatSendResponse, _err error) {
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
		Action:      tea.String("PrivateChatSend"),
		Version:     tea.String("robot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/robot/privateChatMessages/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PrivateChatSendResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PrivateChatSend(request *PrivateChatSendRequest) (_result *PrivateChatSendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PrivateChatSendHeaders{}
	_result = &PrivateChatSendResponse{}
	_body, _err := client.PrivateChatSendWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBotInstanceInGroupInfoWithOptions(request *QueryBotInstanceInGroupInfoRequest, headers *QueryBotInstanceInGroupInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryBotInstanceInGroupInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
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
	params := &openapi.Params{
		Action:      tea.String("QueryBotInstanceInGroupInfo"),
		Version:     tea.String("robot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/robot/groups/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBotInstanceInGroupInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBotInstanceInGroupInfo(request *QueryBotInstanceInGroupInfoRequest) (_result *QueryBotInstanceInGroupInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryBotInstanceInGroupInfoHeaders{}
	_result = &QueryBotInstanceInGroupInfoResponse{}
	_body, _err := client.QueryBotInstanceInGroupInfoWithOptions(request, headers, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("QueryRobotPlugin"),
		Version:     tea.String("robot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/robot/plugins/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRobotPluginResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) RobotMessageFileDownloadWithOptions(request *RobotMessageFileDownloadRequest, headers *RobotMessageFileDownloadHeaders, runtime *util.RuntimeOptions) (_result *RobotMessageFileDownloadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DownloadCode)) {
		body["downloadCode"] = request.DownloadCode
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
	params := &openapi.Params{
		Action:      tea.String("RobotMessageFileDownload"),
		Version:     tea.String("robot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/robot/messageFiles/download"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RobotMessageFileDownloadResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RobotMessageFileDownload(request *RobotMessageFileDownloadRequest) (_result *RobotMessageFileDownloadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RobotMessageFileDownloadHeaders{}
	_result = &RobotMessageFileDownloadResponse{}
	_body, _err := client.RobotMessageFileDownloadWithOptions(request, headers, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("SendRobotDingMessage"),
		Version:     tea.String("robot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/robot/dingMessages/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SendRobotDingMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("SetRobotPlugin"),
		Version:     tea.String("robot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/robot/plugins/set"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SetRobotPluginResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("UpdateInstalledRobot"),
		Version:     tea.String("robot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/robot/managements/infos"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstalledRobotResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
