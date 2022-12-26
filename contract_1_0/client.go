// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package contract_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type SendContractCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendContractCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendContractCardHeaders) GoString() string {
	return s.String()
}

func (s *SendContractCardHeaders) SetCommonHeaders(v map[string]*string) *SendContractCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendContractCardHeaders) SetXAcsDingtalkAccessToken(v string) *SendContractCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendContractCardRequest struct {
	// 卡片类型
	CardType *string `json:"cardType,omitempty" xml:"cardType,omitempty"`
	// 合同信息
	ContractInfo *SendContractCardRequestContractInfo `json:"contractInfo,omitempty" xml:"contractInfo,omitempty" type:"Struct"`
	// 合同的企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 额外信息
	Extension map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 审批实例id
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// 接收群id
	ReceiveGroups []*string `json:"receiveGroups,omitempty" xml:"receiveGroups,omitempty" type:"Repeated"`
	// 接收者
	Receivers []*SendContractCardRequestReceivers `json:"receivers,omitempty" xml:"receivers,omitempty" type:"Repeated"`
	// 发送者
	Sender *SendContractCardRequestSender `json:"sender,omitempty" xml:"sender,omitempty" type:"Struct"`
	// 是否同步单聊
	SyncSingleChat *bool `json:"syncSingleChat,omitempty" xml:"syncSingleChat,omitempty"`
}

func (s SendContractCardRequest) String() string {
	return tea.Prettify(s)
}

func (s SendContractCardRequest) GoString() string {
	return s.String()
}

func (s *SendContractCardRequest) SetCardType(v string) *SendContractCardRequest {
	s.CardType = &v
	return s
}

func (s *SendContractCardRequest) SetContractInfo(v *SendContractCardRequestContractInfo) *SendContractCardRequest {
	s.ContractInfo = v
	return s
}

func (s *SendContractCardRequest) SetCorpId(v string) *SendContractCardRequest {
	s.CorpId = &v
	return s
}

func (s *SendContractCardRequest) SetExtension(v map[string]*string) *SendContractCardRequest {
	s.Extension = v
	return s
}

func (s *SendContractCardRequest) SetProcessInstanceId(v string) *SendContractCardRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *SendContractCardRequest) SetReceiveGroups(v []*string) *SendContractCardRequest {
	s.ReceiveGroups = v
	return s
}

func (s *SendContractCardRequest) SetReceivers(v []*SendContractCardRequestReceivers) *SendContractCardRequest {
	s.Receivers = v
	return s
}

func (s *SendContractCardRequest) SetSender(v *SendContractCardRequestSender) *SendContractCardRequest {
	s.Sender = v
	return s
}

func (s *SendContractCardRequest) SetSyncSingleChat(v bool) *SendContractCardRequest {
	s.SyncSingleChat = &v
	return s
}

type SendContractCardRequestContractInfo struct {
	// 合同编号
	ContractCode *string `json:"contractCode,omitempty" xml:"contractCode,omitempty"`
	// 合同名称
	ContractName *string `json:"contractName,omitempty" xml:"contractName,omitempty"`
	// 签署时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 签署人名称
	SignUserName *string `json:"signUserName,omitempty" xml:"signUserName,omitempty"`
}

func (s SendContractCardRequestContractInfo) String() string {
	return tea.Prettify(s)
}

func (s SendContractCardRequestContractInfo) GoString() string {
	return s.String()
}

func (s *SendContractCardRequestContractInfo) SetContractCode(v string) *SendContractCardRequestContractInfo {
	s.ContractCode = &v
	return s
}

func (s *SendContractCardRequestContractInfo) SetContractName(v string) *SendContractCardRequestContractInfo {
	s.ContractName = &v
	return s
}

func (s *SendContractCardRequestContractInfo) SetCreateTime(v int64) *SendContractCardRequestContractInfo {
	s.CreateTime = &v
	return s
}

func (s *SendContractCardRequestContractInfo) SetSignUserName(v string) *SendContractCardRequestContractInfo {
	s.SignUserName = &v
	return s
}

type SendContractCardRequestReceivers struct {
	// 接收者所在组织
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户类型
	UserType *string `json:"userType,omitempty" xml:"userType,omitempty"`
}

func (s SendContractCardRequestReceivers) String() string {
	return tea.Prettify(s)
}

func (s SendContractCardRequestReceivers) GoString() string {
	return s.String()
}

func (s *SendContractCardRequestReceivers) SetCorpId(v string) *SendContractCardRequestReceivers {
	s.CorpId = &v
	return s
}

func (s *SendContractCardRequestReceivers) SetUserId(v string) *SendContractCardRequestReceivers {
	s.UserId = &v
	return s
}

func (s *SendContractCardRequestReceivers) SetUserType(v string) *SendContractCardRequestReceivers {
	s.UserType = &v
	return s
}

type SendContractCardRequestSender struct {
	// 发起人所在组织
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户类型
	UserType *string `json:"userType,omitempty" xml:"userType,omitempty"`
}

func (s SendContractCardRequestSender) String() string {
	return tea.Prettify(s)
}

func (s SendContractCardRequestSender) GoString() string {
	return s.String()
}

func (s *SendContractCardRequestSender) SetCorpId(v string) *SendContractCardRequestSender {
	s.CorpId = &v
	return s
}

func (s *SendContractCardRequestSender) SetUserId(v string) *SendContractCardRequestSender {
	s.UserId = &v
	return s
}

func (s *SendContractCardRequestSender) SetUserType(v string) *SendContractCardRequestSender {
	s.UserType = &v
	return s
}

type SendContractCardResponseBody struct {
	// 成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SendContractCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendContractCardResponseBody) GoString() string {
	return s.String()
}

func (s *SendContractCardResponseBody) SetSuccess(v bool) *SendContractCardResponseBody {
	s.Success = &v
	return s
}

type SendContractCardResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendContractCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendContractCardResponse) String() string {
	return tea.Prettify(s)
}

func (s SendContractCardResponse) GoString() string {
	return s.String()
}

func (s *SendContractCardResponse) SetHeaders(v map[string]*string) *SendContractCardResponse {
	s.Headers = v
	return s
}

func (s *SendContractCardResponse) SetBody(v *SendContractCardResponseBody) *SendContractCardResponse {
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

func (client *Client) SendContractCard(request *SendContractCardRequest) (_result *SendContractCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendContractCardHeaders{}
	_result = &SendContractCardResponse{}
	_body, _err := client.SendContractCardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendContractCardWithOptions(request *SendContractCardRequest, headers *SendContractCardHeaders, runtime *util.RuntimeOptions) (_result *SendContractCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CardType)) {
		body["cardType"] = request.CardType
	}

	if !tea.BoolValue(util.IsUnset(request.ContractInfo)) {
		body["contractInfo"] = request.ContractInfo
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Extension)) {
		body["extension"] = request.Extension
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		body["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiveGroups)) {
		body["receiveGroups"] = request.ReceiveGroups
	}

	if !tea.BoolValue(util.IsUnset(request.Receivers)) {
		body["receivers"] = request.Receivers
	}

	if !tea.BoolValue(util.IsUnset(request.Sender)) {
		body["sender"] = request.Sender
	}

	if !tea.BoolValue(util.IsUnset(request.SyncSingleChat)) {
		body["syncSingleChat"] = request.SyncSingleChat
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
	_result = &SendContractCardResponse{}
	_body, _err := client.DoROARequest(tea.String("SendContractCard"), tea.String("contract_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/contract/cards/send"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
