// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package im_2_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type UnionIdPrivateDataMapValue struct {
	CardParamMap map[string]*string `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
}

func (s UnionIdPrivateDataMapValue) String() string {
	return tea.Prettify(s)
}

func (s UnionIdPrivateDataMapValue) GoString() string {
	return s.String()
}

func (s *UnionIdPrivateDataMapValue) SetCardParamMap(v map[string]*string) *UnionIdPrivateDataMapValue {
	s.CardParamMap = v
	return s
}

type UserIdPrivateDataMapValue struct {
	CardParamMap map[string]*string `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
}

func (s UserIdPrivateDataMapValue) String() string {
	return tea.Prettify(s)
}

func (s UserIdPrivateDataMapValue) GoString() string {
	return s.String()
}

func (s *UserIdPrivateDataMapValue) SetCardParamMap(v map[string]*string) *UserIdPrivateDataMapValue {
	s.CardParamMap = v
	return s
}

type CloseTopboxHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CloseTopboxHeaders) String() string {
	return tea.Prettify(s)
}

func (s CloseTopboxHeaders) GoString() string {
	return s.String()
}

func (s *CloseTopboxHeaders) SetCommonHeaders(v map[string]*string) *CloseTopboxHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CloseTopboxHeaders) SetXAcsDingtalkAccessToken(v string) *CloseTopboxHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CloseTopboxRequest struct {
	ConversationType   *int32  `json:"conversationType,omitempty" xml:"conversationType,omitempty"`
	CoolAppCode        *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	GroupTemplateId    *string `json:"groupTemplateId,omitempty" xml:"groupTemplateId,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OutTrackId         *string `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	RobotCode          *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	UnoinId            *string `json:"unoinId,omitempty" xml:"unoinId,omitempty"`
	UserId             *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CloseTopboxRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseTopboxRequest) GoString() string {
	return s.String()
}

func (s *CloseTopboxRequest) SetConversationType(v int32) *CloseTopboxRequest {
	s.ConversationType = &v
	return s
}

func (s *CloseTopboxRequest) SetCoolAppCode(v string) *CloseTopboxRequest {
	s.CoolAppCode = &v
	return s
}

func (s *CloseTopboxRequest) SetGroupTemplateId(v string) *CloseTopboxRequest {
	s.GroupTemplateId = &v
	return s
}

func (s *CloseTopboxRequest) SetOpenConversationId(v string) *CloseTopboxRequest {
	s.OpenConversationId = &v
	return s
}

func (s *CloseTopboxRequest) SetOutTrackId(v string) *CloseTopboxRequest {
	s.OutTrackId = &v
	return s
}

func (s *CloseTopboxRequest) SetRobotCode(v string) *CloseTopboxRequest {
	s.RobotCode = &v
	return s
}

func (s *CloseTopboxRequest) SetUnoinId(v string) *CloseTopboxRequest {
	s.UnoinId = &v
	return s
}

func (s *CloseTopboxRequest) SetUserId(v string) *CloseTopboxRequest {
	s.UserId = &v
	return s
}

type CloseTopboxResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CloseTopboxResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseTopboxResponseBody) GoString() string {
	return s.String()
}

func (s *CloseTopboxResponseBody) SetSuccess(v bool) *CloseTopboxResponseBody {
	s.Success = &v
	return s
}

type CloseTopboxResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CloseTopboxResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloseTopboxResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseTopboxResponse) GoString() string {
	return s.String()
}

func (s *CloseTopboxResponse) SetHeaders(v map[string]*string) *CloseTopboxResponse {
	s.Headers = v
	return s
}

func (s *CloseTopboxResponse) SetStatusCode(v int32) *CloseTopboxResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseTopboxResponse) SetBody(v *CloseTopboxResponseBody) *CloseTopboxResponse {
	s.Body = v
	return s
}

type CreateTopboxHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateTopboxHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTopboxHeaders) GoString() string {
	return s.String()
}

func (s *CreateTopboxHeaders) SetCommonHeaders(v map[string]*string) *CreateTopboxHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTopboxHeaders) SetXAcsDingtalkAccessToken(v string) *CreateTopboxHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateTopboxRequest struct {
	CallbackRouteKey      *string                                `json:"callbackRouteKey,omitempty" xml:"callbackRouteKey,omitempty"`
	CardData              *CreateTopboxRequestCardData           `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	CardSettings          *CreateTopboxRequestCardSettings       `json:"cardSettings,omitempty" xml:"cardSettings,omitempty" type:"Struct"`
	CardTemplateId        *string                                `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	ConversationType      *int32                                 `json:"conversationType,omitempty" xml:"conversationType,omitempty"`
	CoolAppCode           *string                                `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	ExpiredTime           *int64                                 `json:"expiredTime,omitempty" xml:"expiredTime,omitempty"`
	GroupTemplateId       *string                                `json:"groupTemplateId,omitempty" xml:"groupTemplateId,omitempty"`
	OpenConversationId    *string                                `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OutTrackId            *string                                `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	Platforms             *string                                `json:"platforms,omitempty" xml:"platforms,omitempty"`
	ReceiverUnionIdList   []*string                              `json:"receiverUnionIdList,omitempty" xml:"receiverUnionIdList,omitempty" type:"Repeated"`
	ReceiverUserIdList    []*string                              `json:"receiverUserIdList,omitempty" xml:"receiverUserIdList,omitempty" type:"Repeated"`
	RobotCode             *string                                `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	UnionIdPrivateDataMap map[string]*UnionIdPrivateDataMapValue `json:"unionIdPrivateDataMap,omitempty" xml:"unionIdPrivateDataMap,omitempty"`
	UnoinId               *string                                `json:"unoinId,omitempty" xml:"unoinId,omitempty"`
	UserId                *string                                `json:"userId,omitempty" xml:"userId,omitempty"`
	UserIdPrivateDataMap  map[string]*UserIdPrivateDataMapValue  `json:"userIdPrivateDataMap,omitempty" xml:"userIdPrivateDataMap,omitempty"`
}

func (s CreateTopboxRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTopboxRequest) GoString() string {
	return s.String()
}

func (s *CreateTopboxRequest) SetCallbackRouteKey(v string) *CreateTopboxRequest {
	s.CallbackRouteKey = &v
	return s
}

func (s *CreateTopboxRequest) SetCardData(v *CreateTopboxRequestCardData) *CreateTopboxRequest {
	s.CardData = v
	return s
}

func (s *CreateTopboxRequest) SetCardSettings(v *CreateTopboxRequestCardSettings) *CreateTopboxRequest {
	s.CardSettings = v
	return s
}

func (s *CreateTopboxRequest) SetCardTemplateId(v string) *CreateTopboxRequest {
	s.CardTemplateId = &v
	return s
}

func (s *CreateTopboxRequest) SetConversationType(v int32) *CreateTopboxRequest {
	s.ConversationType = &v
	return s
}

func (s *CreateTopboxRequest) SetCoolAppCode(v string) *CreateTopboxRequest {
	s.CoolAppCode = &v
	return s
}

func (s *CreateTopboxRequest) SetExpiredTime(v int64) *CreateTopboxRequest {
	s.ExpiredTime = &v
	return s
}

func (s *CreateTopboxRequest) SetGroupTemplateId(v string) *CreateTopboxRequest {
	s.GroupTemplateId = &v
	return s
}

func (s *CreateTopboxRequest) SetOpenConversationId(v string) *CreateTopboxRequest {
	s.OpenConversationId = &v
	return s
}

func (s *CreateTopboxRequest) SetOutTrackId(v string) *CreateTopboxRequest {
	s.OutTrackId = &v
	return s
}

func (s *CreateTopboxRequest) SetPlatforms(v string) *CreateTopboxRequest {
	s.Platforms = &v
	return s
}

func (s *CreateTopboxRequest) SetReceiverUnionIdList(v []*string) *CreateTopboxRequest {
	s.ReceiverUnionIdList = v
	return s
}

func (s *CreateTopboxRequest) SetReceiverUserIdList(v []*string) *CreateTopboxRequest {
	s.ReceiverUserIdList = v
	return s
}

func (s *CreateTopboxRequest) SetRobotCode(v string) *CreateTopboxRequest {
	s.RobotCode = &v
	return s
}

func (s *CreateTopboxRequest) SetUnionIdPrivateDataMap(v map[string]*UnionIdPrivateDataMapValue) *CreateTopboxRequest {
	s.UnionIdPrivateDataMap = v
	return s
}

func (s *CreateTopboxRequest) SetUnoinId(v string) *CreateTopboxRequest {
	s.UnoinId = &v
	return s
}

func (s *CreateTopboxRequest) SetUserId(v string) *CreateTopboxRequest {
	s.UserId = &v
	return s
}

func (s *CreateTopboxRequest) SetUserIdPrivateDataMap(v map[string]*UserIdPrivateDataMapValue) *CreateTopboxRequest {
	s.UserIdPrivateDataMap = v
	return s
}

type CreateTopboxRequestCardData struct {
	CardParamMap map[string]*string `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
}

func (s CreateTopboxRequestCardData) String() string {
	return tea.Prettify(s)
}

func (s CreateTopboxRequestCardData) GoString() string {
	return s.String()
}

func (s *CreateTopboxRequestCardData) SetCardParamMap(v map[string]*string) *CreateTopboxRequestCardData {
	s.CardParamMap = v
	return s
}

type CreateTopboxRequestCardSettings struct {
	PullStrategy *bool `json:"pullStrategy,omitempty" xml:"pullStrategy,omitempty"`
}

func (s CreateTopboxRequestCardSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateTopboxRequestCardSettings) GoString() string {
	return s.String()
}

func (s *CreateTopboxRequestCardSettings) SetPullStrategy(v bool) *CreateTopboxRequestCardSettings {
	s.PullStrategy = &v
	return s
}

type CreateTopboxResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateTopboxResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTopboxResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTopboxResponseBody) SetSuccess(v bool) *CreateTopboxResponseBody {
	s.Success = &v
	return s
}

type CreateTopboxResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTopboxResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTopboxResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTopboxResponse) GoString() string {
	return s.String()
}

func (s *CreateTopboxResponse) SetHeaders(v map[string]*string) *CreateTopboxResponse {
	s.Headers = v
	return s
}

func (s *CreateTopboxResponse) SetStatusCode(v int32) *CreateTopboxResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTopboxResponse) SetBody(v *CreateTopboxResponseBody) *CreateTopboxResponse {
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

func (client *Client) CloseTopboxWithOptions(request *CloseTopboxRequest, headers *CloseTopboxHeaders, runtime *util.RuntimeOptions) (_result *CloseTopboxResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.GroupTemplateId)) {
		body["groupTemplateId"] = request.GroupTemplateId
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

	if !tea.BoolValue(util.IsUnset(request.UnoinId)) {
		body["unoinId"] = request.UnoinId
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
		Action:      tea.String("CloseTopbox"),
		Version:     tea.String("im_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/im/topBoxes/close"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseTopboxResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloseTopbox(request *CloseTopboxRequest) (_result *CloseTopboxResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CloseTopboxHeaders{}
	_result = &CloseTopboxResponse{}
	_body, _err := client.CloseTopboxWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTopboxWithOptions(request *CreateTopboxRequest, headers *CreateTopboxHeaders, runtime *util.RuntimeOptions) (_result *CreateTopboxResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallbackRouteKey)) {
		body["callbackRouteKey"] = request.CallbackRouteKey
	}

	if !tea.BoolValue(util.IsUnset(request.CardData)) {
		body["cardData"] = request.CardData
	}

	if !tea.BoolValue(util.IsUnset(request.CardSettings)) {
		body["cardSettings"] = request.CardSettings
	}

	if !tea.BoolValue(util.IsUnset(request.CardTemplateId)) {
		body["cardTemplateId"] = request.CardTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.ConversationType)) {
		body["conversationType"] = request.ConversationType
	}

	if !tea.BoolValue(util.IsUnset(request.CoolAppCode)) {
		body["coolAppCode"] = request.CoolAppCode
	}

	if !tea.BoolValue(util.IsUnset(request.ExpiredTime)) {
		body["expiredTime"] = request.ExpiredTime
	}

	if !tea.BoolValue(util.IsUnset(request.GroupTemplateId)) {
		body["groupTemplateId"] = request.GroupTemplateId
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

	if !tea.BoolValue(util.IsUnset(request.ReceiverUnionIdList)) {
		body["receiverUnionIdList"] = request.ReceiverUnionIdList
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUserIdList)) {
		body["receiverUserIdList"] = request.ReceiverUserIdList
	}

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
	}

	if !tea.BoolValue(util.IsUnset(request.UnionIdPrivateDataMap)) {
		body["unionIdPrivateDataMap"] = request.UnionIdPrivateDataMap
	}

	if !tea.BoolValue(util.IsUnset(request.UnoinId)) {
		body["unoinId"] = request.UnoinId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
	params := &openapi.Params{
		Action:      tea.String("CreateTopbox"),
		Version:     tea.String("im_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/im/topBoxes"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTopboxResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTopbox(request *CreateTopboxRequest) (_result *CreateTopboxResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTopboxHeaders{}
	_result = &CreateTopboxResponse{}
	_body, _err := client.CreateTopboxWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
