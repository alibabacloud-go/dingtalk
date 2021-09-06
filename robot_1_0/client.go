// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package robot_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	// 机器人的robotCode
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	// 被推送会话人员的userId列表
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
	// 消息的msgKey
	MsgKey *string `json:"msgKey,omitempty" xml:"msgKey,omitempty"`
	// 消息体
	MsgParam *string `json:"msgParam,omitempty" xml:"msgParam,omitempty"`
}

func (s BatchSendOTORequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOTORequest) GoString() string {
	return s.String()
}

func (s *BatchSendOTORequest) SetRobotCode(v string) *BatchSendOTORequest {
	s.RobotCode = &v
	return s
}

func (s *BatchSendOTORequest) SetUserIds(v []*string) *BatchSendOTORequest {
	s.UserIds = v
	return s
}

func (s *BatchSendOTORequest) SetMsgKey(v string) *BatchSendOTORequest {
	s.MsgKey = &v
	return s
}

func (s *BatchSendOTORequest) SetMsgParam(v string) *BatchSendOTORequest {
	s.MsgParam = &v
	return s
}

type BatchSendOTOResponseBody struct {
	// 消息id
	ProcessQueryKey *string `json:"processQueryKey,omitempty" xml:"processQueryKey,omitempty"`
	// 无效的用户userId列表
	InvalidStaffIdList []*string `json:"invalidStaffIdList,omitempty" xml:"invalidStaffIdList,omitempty" type:"Repeated"`
	// 推送频繁，被限流的用户userId列表
	FlowControlledStaffIdList []*string `json:"flowControlledStaffIdList,omitempty" xml:"flowControlledStaffIdList,omitempty" type:"Repeated"`
}

func (s BatchSendOTOResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOTOResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSendOTOResponseBody) SetProcessQueryKey(v string) *BatchSendOTOResponseBody {
	s.ProcessQueryKey = &v
	return s
}

func (s *BatchSendOTOResponseBody) SetInvalidStaffIdList(v []*string) *BatchSendOTOResponseBody {
	s.InvalidStaffIdList = v
	return s
}

func (s *BatchSendOTOResponseBody) SetFlowControlledStaffIdList(v []*string) *BatchSendOTOResponseBody {
	s.FlowControlledStaffIdList = v
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
	// 机器人robotCode
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	// 消息已读查询标志
	ProcessQueryKey *string `json:"processQueryKey,omitempty" xml:"processQueryKey,omitempty"`
}

func (s BatchOTOQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchOTOQueryRequest) GoString() string {
	return s.String()
}

func (s *BatchOTOQueryRequest) SetRobotCode(v string) *BatchOTOQueryRequest {
	s.RobotCode = &v
	return s
}

func (s *BatchOTOQueryRequest) SetProcessQueryKey(v string) *BatchOTOQueryRequest {
	s.ProcessQueryKey = &v
	return s
}

type BatchOTOQueryResponseBody struct {
	// 消息发送状态
	SendStatus *string `json:"sendStatus,omitempty" xml:"sendStatus,omitempty"`
	// 消息已读情况
	MessageReadInfoList []*BatchOTOQueryResponseBodyMessageReadInfoList `json:"messageReadInfoList,omitempty" xml:"messageReadInfoList,omitempty" type:"Repeated"`
}

func (s BatchOTOQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchOTOQueryResponseBody) GoString() string {
	return s.String()
}

func (s *BatchOTOQueryResponseBody) SetSendStatus(v string) *BatchOTOQueryResponseBody {
	s.SendStatus = &v
	return s
}

func (s *BatchOTOQueryResponseBody) SetMessageReadInfoList(v []*BatchOTOQueryResponseBodyMessageReadInfoList) *BatchOTOQueryResponseBody {
	s.MessageReadInfoList = v
	return s
}

type BatchOTOQueryResponseBodyMessageReadInfoList struct {
	// 姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 工号
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 已读状态
	ReadStatus *string `json:"readStatus,omitempty" xml:"readStatus,omitempty"`
	// 已读时间
	ReadTimestamp *int64 `json:"readTimestamp,omitempty" xml:"readTimestamp,omitempty"`
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

func (s *BatchOTOQueryResponseBodyMessageReadInfoList) SetUserId(v string) *BatchOTOQueryResponseBodyMessageReadInfoList {
	s.UserId = &v
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
	// 机器人的Id
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	// 群聊的对外开放Id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 接受人的userId列表
	ReceiverUserIdList []*string `json:"receiverUserIdList,omitempty" xml:"receiverUserIdList,omitempty" type:"Repeated"`
	// 颁发的模版id，可通过宜搭申请：https://yida.alibaba-inc.com/alibaba/web/APP_NSUGAGIQUMI4ESRA7O7D/inst/homepage/#/FORM-WO866371VGXSECXX4M0NC9KSGAT92VSA3TZSK9B
	DingTemplateId *string `json:"dingTemplateId,omitempty" xml:"dingTemplateId,omitempty"`
	// 模版对应的参数
	ContentParams map[string]*string `json:"contentParams,omitempty" xml:"contentParams,omitempty"`
}

func (s SendRobotDingMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendRobotDingMessageRequest) GoString() string {
	return s.String()
}

func (s *SendRobotDingMessageRequest) SetRobotCode(v string) *SendRobotDingMessageRequest {
	s.RobotCode = &v
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

func (s *SendRobotDingMessageRequest) SetDingTemplateId(v string) *SendRobotDingMessageRequest {
	s.DingTemplateId = &v
	return s
}

func (s *SendRobotDingMessageRequest) SetContentParams(v map[string]*string) *SendRobotDingMessageRequest {
	s.ContentParams = v
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
	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
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
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
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
	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		query["robotCode"] = request.RobotCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessQueryKey)) {
		query["processQueryKey"] = request.ProcessQueryKey
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
	_result = &BatchOTOQueryResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchOTOQuery"), tea.String("robot_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/robot/oToMessages/readStatus"), tea.String("json"), req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUserIdList)) {
		body["receiverUserIdList"] = request.ReceiverUserIdList
	}

	if !tea.BoolValue(util.IsUnset(request.DingTemplateId)) {
		body["dingTemplateId"] = request.DingTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.ContentParams)) {
		body["contentParams"] = request.ContentParams
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
	_result = &SendRobotDingMessageResponse{}
	_body, _err := client.DoROARequest(tea.String("SendRobotDingMessage"), tea.String("robot_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/robot/dingMessages/send"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
