// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package ai_interaction_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type FinishHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s FinishHeaders) String() string {
	return tea.Prettify(s)
}

func (s FinishHeaders) GoString() string {
	return s.String()
}

func (s *FinishHeaders) SetCommonHeaders(v map[string]*string) *FinishHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FinishHeaders) SetXAcsDingtalkAccessToken(v string) *FinishHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type FinishRequest struct {
	ConversationToken *string `json:"conversationToken,omitempty" xml:"conversationToken,omitempty"`
}

func (s FinishRequest) String() string {
	return tea.Prettify(s)
}

func (s FinishRequest) GoString() string {
	return s.String()
}

func (s *FinishRequest) SetConversationToken(v string) *FinishRequest {
	s.ConversationToken = &v
	return s
}

type FinishResponseBody struct {
	Result *FinishResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s FinishResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FinishResponseBody) GoString() string {
	return s.String()
}

func (s *FinishResponseBody) SetResult(v *FinishResponseBodyResult) *FinishResponseBody {
	s.Result = v
	return s
}

type FinishResponseBodyResult struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s FinishResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s FinishResponseBodyResult) GoString() string {
	return s.String()
}

func (s *FinishResponseBodyResult) SetSuccess(v bool) *FinishResponseBodyResult {
	s.Success = &v
	return s
}

type FinishResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FinishResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FinishResponse) String() string {
	return tea.Prettify(s)
}

func (s FinishResponse) GoString() string {
	return s.String()
}

func (s *FinishResponse) SetHeaders(v map[string]*string) *FinishResponse {
	s.Headers = v
	return s
}

func (s *FinishResponse) SetStatusCode(v int32) *FinishResponse {
	s.StatusCode = &v
	return s
}

func (s *FinishResponse) SetBody(v *FinishResponseBody) *FinishResponse {
	s.Body = v
	return s
}

type PrepareHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PrepareHeaders) String() string {
	return tea.Prettify(s)
}

func (s PrepareHeaders) GoString() string {
	return s.String()
}

func (s *PrepareHeaders) SetCommonHeaders(v map[string]*string) *PrepareHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PrepareHeaders) SetXAcsDingtalkAccessToken(v string) *PrepareHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PrepareRequest struct {
	Content            *string `json:"content,omitempty" xml:"content,omitempty"`
	ContentType        *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	UnionId            *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s PrepareRequest) String() string {
	return tea.Prettify(s)
}

func (s PrepareRequest) GoString() string {
	return s.String()
}

func (s *PrepareRequest) SetContent(v string) *PrepareRequest {
	s.Content = &v
	return s
}

func (s *PrepareRequest) SetContentType(v string) *PrepareRequest {
	s.ContentType = &v
	return s
}

func (s *PrepareRequest) SetOpenConversationId(v string) *PrepareRequest {
	s.OpenConversationId = &v
	return s
}

func (s *PrepareRequest) SetUnionId(v string) *PrepareRequest {
	s.UnionId = &v
	return s
}

type PrepareResponseBody struct {
	Result *PrepareResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s PrepareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PrepareResponseBody) GoString() string {
	return s.String()
}

func (s *PrepareResponseBody) SetResult(v *PrepareResponseBodyResult) *PrepareResponseBody {
	s.Result = v
	return s
}

type PrepareResponseBodyResult struct {
	ConversationToken *string `json:"conversationToken,omitempty" xml:"conversationToken,omitempty"`
}

func (s PrepareResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PrepareResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PrepareResponseBodyResult) SetConversationToken(v string) *PrepareResponseBodyResult {
	s.ConversationToken = &v
	return s
}

type PrepareResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PrepareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PrepareResponse) String() string {
	return tea.Prettify(s)
}

func (s PrepareResponse) GoString() string {
	return s.String()
}

func (s *PrepareResponse) SetHeaders(v map[string]*string) *PrepareResponse {
	s.Headers = v
	return s
}

func (s *PrepareResponse) SetStatusCode(v int32) *PrepareResponse {
	s.StatusCode = &v
	return s
}

func (s *PrepareResponse) SetBody(v *PrepareResponseBody) *PrepareResponse {
	s.Body = v
	return s
}

type ReplyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ReplyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ReplyHeaders) GoString() string {
	return s.String()
}

func (s *ReplyHeaders) SetCommonHeaders(v map[string]*string) *ReplyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReplyHeaders) SetXAcsDingtalkAccessToken(v string) *ReplyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ReplyRequest struct {
	Content           *string `json:"content,omitempty" xml:"content,omitempty"`
	ContentType       *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	ConversationToken *string `json:"conversationToken,omitempty" xml:"conversationToken,omitempty"`
}

func (s ReplyRequest) String() string {
	return tea.Prettify(s)
}

func (s ReplyRequest) GoString() string {
	return s.String()
}

func (s *ReplyRequest) SetContent(v string) *ReplyRequest {
	s.Content = &v
	return s
}

func (s *ReplyRequest) SetContentType(v string) *ReplyRequest {
	s.ContentType = &v
	return s
}

func (s *ReplyRequest) SetConversationToken(v string) *ReplyRequest {
	s.ConversationToken = &v
	return s
}

type ReplyResponseBody struct {
	Result *ReplyResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ReplyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReplyResponseBody) GoString() string {
	return s.String()
}

func (s *ReplyResponseBody) SetResult(v *ReplyResponseBodyResult) *ReplyResponseBody {
	s.Result = v
	return s
}

type ReplyResponseBodyResult struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ReplyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ReplyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ReplyResponseBodyResult) SetSuccess(v bool) *ReplyResponseBodyResult {
	s.Success = &v
	return s
}

type ReplyResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReplyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReplyResponse) String() string {
	return tea.Prettify(s)
}

func (s ReplyResponse) GoString() string {
	return s.String()
}

func (s *ReplyResponse) SetHeaders(v map[string]*string) *ReplyResponse {
	s.Headers = v
	return s
}

func (s *ReplyResponse) SetStatusCode(v int32) *ReplyResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplyResponse) SetBody(v *ReplyResponseBody) *ReplyResponse {
	s.Body = v
	return s
}

type SendHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendHeaders) GoString() string {
	return s.String()
}

func (s *SendHeaders) SetCommonHeaders(v map[string]*string) *SendHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendHeaders) SetXAcsDingtalkAccessToken(v string) *SendHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendRequest struct {
	Content            *string `json:"content,omitempty" xml:"content,omitempty"`
	ContentType        *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	UnionId            *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s SendRequest) String() string {
	return tea.Prettify(s)
}

func (s SendRequest) GoString() string {
	return s.String()
}

func (s *SendRequest) SetContent(v string) *SendRequest {
	s.Content = &v
	return s
}

func (s *SendRequest) SetContentType(v string) *SendRequest {
	s.ContentType = &v
	return s
}

func (s *SendRequest) SetOpenConversationId(v string) *SendRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SendRequest) SetUnionId(v string) *SendRequest {
	s.UnionId = &v
	return s
}

type SendResponseBody struct {
	Result *SendResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s SendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendResponseBody) GoString() string {
	return s.String()
}

func (s *SendResponseBody) SetResult(v *SendResponseBodyResult) *SendResponseBody {
	s.Result = v
	return s
}

type SendResponseBodyResult struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SendResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SendResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SendResponseBodyResult) SetSuccess(v bool) *SendResponseBodyResult {
	s.Success = &v
	return s
}

type SendResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendResponseBody  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendResponse) String() string {
	return tea.Prettify(s)
}

func (s SendResponse) GoString() string {
	return s.String()
}

func (s *SendResponse) SetHeaders(v map[string]*string) *SendResponse {
	s.Headers = v
	return s
}

func (s *SendResponse) SetStatusCode(v int32) *SendResponse {
	s.StatusCode = &v
	return s
}

func (s *SendResponse) SetBody(v *SendResponseBody) *SendResponse {
	s.Body = v
	return s
}

type UpdateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateHeaders) GoString() string {
	return s.String()
}

func (s *UpdateHeaders) SetCommonHeaders(v map[string]*string) *UpdateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateRequest struct {
	Content           *string `json:"content,omitempty" xml:"content,omitempty"`
	ContentType       *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	ConversationToken *string `json:"conversationToken,omitempty" xml:"conversationToken,omitempty"`
}

func (s UpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRequest) GoString() string {
	return s.String()
}

func (s *UpdateRequest) SetContent(v string) *UpdateRequest {
	s.Content = &v
	return s
}

func (s *UpdateRequest) SetContentType(v string) *UpdateRequest {
	s.ContentType = &v
	return s
}

func (s *UpdateRequest) SetConversationToken(v string) *UpdateRequest {
	s.ConversationToken = &v
	return s
}

type UpdateResponseBody struct {
	Result *UpdateResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResponseBody) SetResult(v *UpdateResponseBodyResult) *UpdateResponseBody {
	s.Result = v
	return s
}

type UpdateResponseBodyResult struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateResponseBodyResult) SetSuccess(v bool) *UpdateResponseBodyResult {
	s.Success = &v
	return s
}

type UpdateResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResponse) GoString() string {
	return s.String()
}

func (s *UpdateResponse) SetHeaders(v map[string]*string) *UpdateResponse {
	s.Headers = v
	return s
}

func (s *UpdateResponse) SetStatusCode(v int32) *UpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResponse) SetBody(v *UpdateResponseBody) *UpdateResponse {
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

func (client *Client) FinishWithOptions(request *FinishRequest, headers *FinishHeaders, runtime *util.RuntimeOptions) (_result *FinishResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationToken)) {
		body["conversationToken"] = request.ConversationToken
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
		Action:      tea.String("Finish"),
		Version:     tea.String("aiInteraction_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiInteraction/finish"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &FinishResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Finish(request *FinishRequest) (_result *FinishResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &FinishHeaders{}
	_result = &FinishResponse{}
	_body, _err := client.FinishWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PrepareWithOptions(request *PrepareRequest, headers *PrepareHeaders, runtime *util.RuntimeOptions) (_result *PrepareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ContentType)) {
		body["contentType"] = request.ContentType
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("Prepare"),
		Version:     tea.String("aiInteraction_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiInteraction/prepare"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PrepareResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Prepare(request *PrepareRequest) (_result *PrepareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PrepareHeaders{}
	_result = &PrepareResponse{}
	_body, _err := client.PrepareWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReplyWithOptions(request *ReplyRequest, headers *ReplyHeaders, runtime *util.RuntimeOptions) (_result *ReplyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ContentType)) {
		body["contentType"] = request.ContentType
	}

	if !tea.BoolValue(util.IsUnset(request.ConversationToken)) {
		body["conversationToken"] = request.ConversationToken
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
		Action:      tea.String("Reply"),
		Version:     tea.String("aiInteraction_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiInteraction/reply"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ReplyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Reply(request *ReplyRequest) (_result *ReplyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ReplyHeaders{}
	_result = &ReplyResponse{}
	_body, _err := client.ReplyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendWithOptions(request *SendRequest, headers *SendHeaders, runtime *util.RuntimeOptions) (_result *SendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ContentType)) {
		body["contentType"] = request.ContentType
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("Send"),
		Version:     tea.String("aiInteraction_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiInteraction/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Send(request *SendRequest) (_result *SendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendHeaders{}
	_result = &SendResponse{}
	_body, _err := client.SendWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWithOptions(request *UpdateRequest, headers *UpdateHeaders, runtime *util.RuntimeOptions) (_result *UpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ContentType)) {
		body["contentType"] = request.ContentType
	}

	if !tea.BoolValue(util.IsUnset(request.ConversationToken)) {
		body["conversationToken"] = request.ConversationToken
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
		Action:      tea.String("Update"),
		Version:     tea.String("aiInteraction_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiInteraction/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Update(request *UpdateRequest) (_result *UpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateHeaders{}
	_result = &UpdateResponse{}
	_body, _err := client.UpdateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
