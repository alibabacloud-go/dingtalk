// This file is auto-generated, don't edit it. Thanks.
package ding_one_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DeleteFeedHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteFeedHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteFeedHeaders) GoString() string {
	return s.String()
}

func (s *DeleteFeedHeaders) SetCommonHeaders(v map[string]*string) *DeleteFeedHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteFeedHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteFeedHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteFeedRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dd-one-work-eSo869uR9VhWe2sqTw3dDF
	FeedId *string `json:"feedId,omitempty" xml:"feedId,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// ntThCP2X44FWclaliPIXPUQiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DeleteFeedRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFeedRequest) GoString() string {
	return s.String()
}

func (s *DeleteFeedRequest) SetFeedId(v string) *DeleteFeedRequest {
	s.FeedId = &v
	return s
}

func (s *DeleteFeedRequest) SetUnionId(v string) *DeleteFeedRequest {
	s.UnionId = &v
	return s
}

type DeleteFeedResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteFeedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFeedResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFeedResponseBody) SetRequestId(v string) *DeleteFeedResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFeedResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFeedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFeedResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFeedResponse) GoString() string {
	return s.String()
}

func (s *DeleteFeedResponse) SetHeaders(v map[string]*string) *DeleteFeedResponse {
	s.Headers = v
	return s
}

func (s *DeleteFeedResponse) SetStatusCode(v int32) *DeleteFeedResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFeedResponse) SetBody(v *DeleteFeedResponseBody) *DeleteFeedResponse {
	s.Body = v
	return s
}

type DeliverOneFeedHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeliverOneFeedHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeliverOneFeedHeaders) GoString() string {
	return s.String()
}

func (s *DeliverOneFeedHeaders) SetCommonHeaders(v map[string]*string) *DeliverOneFeedHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeliverOneFeedHeaders) SetXAcsDingtalkAccessToken(v string) *DeliverOneFeedHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeliverOneFeedRequest struct {
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	CoverMediaId *string `json:"coverMediaId,omitempty" xml:"coverMediaId,omitempty"`
	ExpireTime   *int64  `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// This parameter is required.
	Summary *string   `json:"summary,omitempty" xml:"summary,omitempty"`
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
	// This parameter is required.
	VideoMediaId *string `json:"videoMediaId,omitempty" xml:"videoMediaId,omitempty"`
	// This parameter is required.
	VideoTags map[string]*string `json:"videoTags,omitempty" xml:"videoTags,omitempty"`
}

func (s DeliverOneFeedRequest) String() string {
	return tea.Prettify(s)
}

func (s DeliverOneFeedRequest) GoString() string {
	return s.String()
}

func (s *DeliverOneFeedRequest) SetContent(v string) *DeliverOneFeedRequest {
	s.Content = &v
	return s
}

func (s *DeliverOneFeedRequest) SetCoverMediaId(v string) *DeliverOneFeedRequest {
	s.CoverMediaId = &v
	return s
}

func (s *DeliverOneFeedRequest) SetExpireTime(v int64) *DeliverOneFeedRequest {
	s.ExpireTime = &v
	return s
}

func (s *DeliverOneFeedRequest) SetSummary(v string) *DeliverOneFeedRequest {
	s.Summary = &v
	return s
}

func (s *DeliverOneFeedRequest) SetUserIds(v []*string) *DeliverOneFeedRequest {
	s.UserIds = v
	return s
}

func (s *DeliverOneFeedRequest) SetVideoMediaId(v string) *DeliverOneFeedRequest {
	s.VideoMediaId = &v
	return s
}

func (s *DeliverOneFeedRequest) SetVideoTags(v map[string]*string) *DeliverOneFeedRequest {
	s.VideoTags = v
	return s
}

type DeliverOneFeedResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeliverOneFeedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeliverOneFeedResponseBody) GoString() string {
	return s.String()
}

func (s *DeliverOneFeedResponseBody) SetResult(v string) *DeliverOneFeedResponseBody {
	s.Result = &v
	return s
}

type DeliverOneFeedResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeliverOneFeedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeliverOneFeedResponse) String() string {
	return tea.Prettify(s)
}

func (s DeliverOneFeedResponse) GoString() string {
	return s.String()
}

func (s *DeliverOneFeedResponse) SetHeaders(v map[string]*string) *DeliverOneFeedResponse {
	s.Headers = v
	return s
}

func (s *DeliverOneFeedResponse) SetStatusCode(v int32) *DeliverOneFeedResponse {
	s.StatusCode = &v
	return s
}

func (s *DeliverOneFeedResponse) SetBody(v *DeliverOneFeedResponseBody) *DeliverOneFeedResponse {
	s.Body = v
	return s
}

type PushFeedHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PushFeedHeaders) String() string {
	return tea.Prettify(s)
}

func (s PushFeedHeaders) GoString() string {
	return s.String()
}

func (s *PushFeedHeaders) SetCommonHeaders(v map[string]*string) *PushFeedHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PushFeedHeaders) SetXAcsDingtalkAccessToken(v string) *PushFeedHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PushFeedRequest struct {
	// This parameter is required.
	Content *PushFeedRequestContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 1772177962000
	ExpireTime *int64 `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// This parameter is required.
	FeedFeature map[string]interface{} `json:"feedFeature,omitempty" xml:"feedFeature,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// iwElAqN6aXADBgQABQAGsO9WlNbxvoXaCN
	IdempotentKey *string `json:"idempotentKey,omitempty" xml:"idempotentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ntThCP2X44Fw73IXPUQiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s PushFeedRequest) String() string {
	return tea.Prettify(s)
}

func (s PushFeedRequest) GoString() string {
	return s.String()
}

func (s *PushFeedRequest) SetContent(v *PushFeedRequestContent) *PushFeedRequest {
	s.Content = v
	return s
}

func (s *PushFeedRequest) SetExpireTime(v int64) *PushFeedRequest {
	s.ExpireTime = &v
	return s
}

func (s *PushFeedRequest) SetFeedFeature(v map[string]interface{}) *PushFeedRequest {
	s.FeedFeature = v
	return s
}

func (s *PushFeedRequest) SetIdempotentKey(v string) *PushFeedRequest {
	s.IdempotentKey = &v
	return s
}

func (s *PushFeedRequest) SetUnionId(v string) *PushFeedRequest {
	s.UnionId = &v
	return s
}

type PushFeedRequestContent struct {
	// This parameter is required.
	//
	// example:
	//
	// dsTemplate
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// This parameter is required.
	DslTemplateContent *PushFeedRequestContentDslTemplateContent `json:"dslTemplateContent,omitempty" xml:"dslTemplateContent,omitempty" type:"Struct"`
}

func (s PushFeedRequestContent) String() string {
	return tea.Prettify(s)
}

func (s PushFeedRequestContent) GoString() string {
	return s.String()
}

func (s *PushFeedRequestContent) SetContentType(v string) *PushFeedRequestContent {
	s.ContentType = &v
	return s
}

func (s *PushFeedRequestContent) SetDslTemplateContent(v *PushFeedRequestContentDslTemplateContent) *PushFeedRequestContent {
	s.DslTemplateContent = v
	return s
}

type PushFeedRequestContentDslTemplateContent struct {
	// This parameter is required.
	//
	// example:
	//
	// https://xxxxx.xxx.com/v1.0/test.html
	ApplyUrl *string `json:"applyUrl,omitempty" xml:"applyUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"date":"2025-11-06"}
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushFeedRequestContentDslTemplateContent) String() string {
	return tea.Prettify(s)
}

func (s PushFeedRequestContentDslTemplateContent) GoString() string {
	return s.String()
}

func (s *PushFeedRequestContentDslTemplateContent) SetApplyUrl(v string) *PushFeedRequestContentDslTemplateContent {
	s.ApplyUrl = &v
	return s
}

func (s *PushFeedRequestContentDslTemplateContent) SetBody(v string) *PushFeedRequestContentDslTemplateContent {
	s.Body = &v
	return s
}

type PushFeedResponseBody struct {
	RequestId *string                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *PushFeedResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s PushFeedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushFeedResponseBody) GoString() string {
	return s.String()
}

func (s *PushFeedResponseBody) SetRequestId(v string) *PushFeedResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushFeedResponseBody) SetResult(v *PushFeedResponseBodyResult) *PushFeedResponseBody {
	s.Result = v
	return s
}

type PushFeedResponseBodyResult struct {
	FeedId *string `json:"feedId,omitempty" xml:"feedId,omitempty"`
}

func (s PushFeedResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PushFeedResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PushFeedResponseBodyResult) SetFeedId(v string) *PushFeedResponseBodyResult {
	s.FeedId = &v
	return s
}

type PushFeedResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushFeedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushFeedResponse) String() string {
	return tea.Prettify(s)
}

func (s PushFeedResponse) GoString() string {
	return s.String()
}

func (s *PushFeedResponse) SetHeaders(v map[string]*string) *PushFeedResponse {
	s.Headers = v
	return s
}

func (s *PushFeedResponse) SetStatusCode(v int32) *PushFeedResponse {
	s.StatusCode = &v
	return s
}

func (s *PushFeedResponse) SetBody(v *PushFeedResponseBody) *PushFeedResponse {
	s.Body = v
	return s
}

type UpdateFeedContentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateFeedContentHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateFeedContentHeaders) GoString() string {
	return s.String()
}

func (s *UpdateFeedContentHeaders) SetCommonHeaders(v map[string]*string) *UpdateFeedContentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateFeedContentHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateFeedContentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateFeedContentRequest struct {
	// This parameter is required.
	Content *UpdateFeedContentRequestContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// dd-one-work-eSo869uR9Vhse2sqTw3dDF
	FeedId *string `json:"feedId,omitempty" xml:"feedId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ntThCP2X44FlskVliPIXPUQiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s UpdateFeedContentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFeedContentRequest) GoString() string {
	return s.String()
}

func (s *UpdateFeedContentRequest) SetContent(v *UpdateFeedContentRequestContent) *UpdateFeedContentRequest {
	s.Content = v
	return s
}

func (s *UpdateFeedContentRequest) SetFeedId(v string) *UpdateFeedContentRequest {
	s.FeedId = &v
	return s
}

func (s *UpdateFeedContentRequest) SetUnionId(v string) *UpdateFeedContentRequest {
	s.UnionId = &v
	return s
}

type UpdateFeedContentRequestContent struct {
	// This parameter is required.
	//
	// example:
	//
	// dslTemplate
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// This parameter is required.
	DslTemplateContent *UpdateFeedContentRequestContentDslTemplateContent `json:"dslTemplateContent,omitempty" xml:"dslTemplateContent,omitempty" type:"Struct"`
}

func (s UpdateFeedContentRequestContent) String() string {
	return tea.Prettify(s)
}

func (s UpdateFeedContentRequestContent) GoString() string {
	return s.String()
}

func (s *UpdateFeedContentRequestContent) SetContentType(v string) *UpdateFeedContentRequestContent {
	s.ContentType = &v
	return s
}

func (s *UpdateFeedContentRequestContent) SetDslTemplateContent(v *UpdateFeedContentRequestContentDslTemplateContent) *UpdateFeedContentRequestContent {
	s.DslTemplateContent = v
	return s
}

type UpdateFeedContentRequestContentDslTemplateContent struct {
	// This parameter is required.
	//
	// example:
	//
	// https://xxxxx.xxx.com/v1.0/test.html
	ApplyUrl *string `json:"applyUrl,omitempty" xml:"applyUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"date":"2025-11-06"}
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFeedContentRequestContentDslTemplateContent) String() string {
	return tea.Prettify(s)
}

func (s UpdateFeedContentRequestContentDslTemplateContent) GoString() string {
	return s.String()
}

func (s *UpdateFeedContentRequestContentDslTemplateContent) SetApplyUrl(v string) *UpdateFeedContentRequestContentDslTemplateContent {
	s.ApplyUrl = &v
	return s
}

func (s *UpdateFeedContentRequestContentDslTemplateContent) SetBody(v string) *UpdateFeedContentRequestContentDslTemplateContent {
	s.Body = &v
	return s
}

type UpdateFeedContentResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateFeedContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFeedContentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFeedContentResponseBody) SetRequestId(v string) *UpdateFeedContentResponseBody {
	s.RequestId = &v
	return s
}

type UpdateFeedContentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFeedContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFeedContentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFeedContentResponse) GoString() string {
	return s.String()
}

func (s *UpdateFeedContentResponse) SetHeaders(v map[string]*string) *UpdateFeedContentResponse {
	s.Headers = v
	return s
}

func (s *UpdateFeedContentResponse) SetStatusCode(v int32) *UpdateFeedContentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFeedContentResponse) SetBody(v *UpdateFeedContentResponseBody) *UpdateFeedContentResponse {
	s.Body = v
	return s
}

type UpdateFeedExpireTimeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateFeedExpireTimeHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateFeedExpireTimeHeaders) GoString() string {
	return s.String()
}

func (s *UpdateFeedExpireTimeHeaders) SetCommonHeaders(v map[string]*string) *UpdateFeedExpireTimeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateFeedExpireTimeHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateFeedExpireTimeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateFeedExpireTimeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1772177962000
	ExpireTime *int64 `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dd-one-work-eSo869uR9Vhse2sqTw3dDF
	FeedId *string `json:"feedId,omitempty" xml:"feedId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ntThCP2X44FlskVliPIXPUQiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s UpdateFeedExpireTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFeedExpireTimeRequest) GoString() string {
	return s.String()
}

func (s *UpdateFeedExpireTimeRequest) SetExpireTime(v int64) *UpdateFeedExpireTimeRequest {
	s.ExpireTime = &v
	return s
}

func (s *UpdateFeedExpireTimeRequest) SetFeedId(v string) *UpdateFeedExpireTimeRequest {
	s.FeedId = &v
	return s
}

func (s *UpdateFeedExpireTimeRequest) SetUnionId(v string) *UpdateFeedExpireTimeRequest {
	s.UnionId = &v
	return s
}

type UpdateFeedExpireTimeResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateFeedExpireTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFeedExpireTimeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFeedExpireTimeResponseBody) SetRequestId(v string) *UpdateFeedExpireTimeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateFeedExpireTimeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFeedExpireTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFeedExpireTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFeedExpireTimeResponse) GoString() string {
	return s.String()
}

func (s *UpdateFeedExpireTimeResponse) SetHeaders(v map[string]*string) *UpdateFeedExpireTimeResponse {
	s.Headers = v
	return s
}

func (s *UpdateFeedExpireTimeResponse) SetStatusCode(v int32) *UpdateFeedExpireTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFeedExpireTimeResponse) SetBody(v *UpdateFeedExpireTimeResponseBody) *UpdateFeedExpireTimeResponse {
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
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

// Summary:
//
// 删除Feed
//
// @param request - DeleteFeedRequest
//
// @param headers - DeleteFeedHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFeedResponse
func (client *Client) DeleteFeedWithOptions(request *DeleteFeedRequest, headers *DeleteFeedHeaders, runtime *util.RuntimeOptions) (_result *DeleteFeedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FeedId)) {
		body["feedId"] = request.FeedId
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
		Action:      tea.String("DeleteFeed"),
		Version:     tea.String("dingOne_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/dingOne/frame/feeds/deleteFeed"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFeedResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Feed
//
// @param request - DeleteFeedRequest
//
// @return DeleteFeedResponse
func (client *Client) DeleteFeed(request *DeleteFeedRequest) (_result *DeleteFeedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteFeedHeaders{}
	_result = &DeleteFeedResponse{}
	_body, _err := client.DeleteFeedWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 投放钉钉one中feed流资讯内容
//
// @param request - DeliverOneFeedRequest
//
// @param headers - DeliverOneFeedHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeliverOneFeedResponse
func (client *Client) DeliverOneFeedWithOptions(request *DeliverOneFeedRequest, headers *DeliverOneFeedHeaders, runtime *util.RuntimeOptions) (_result *DeliverOneFeedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.CoverMediaId)) {
		body["coverMediaId"] = request.CoverMediaId
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		body["expireTime"] = request.ExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.Summary)) {
		body["summary"] = request.Summary
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
	}

	if !tea.BoolValue(util.IsUnset(request.VideoMediaId)) {
		body["videoMediaId"] = request.VideoMediaId
	}

	if !tea.BoolValue(util.IsUnset(request.VideoTags)) {
		body["videoTags"] = request.VideoTags
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
		Action:      tea.String("DeliverOneFeed"),
		Version:     tea.String("dingOne_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/dingOne/feed/deliver"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeliverOneFeedResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 投放钉钉one中feed流资讯内容
//
// @param request - DeliverOneFeedRequest
//
// @return DeliverOneFeedResponse
func (client *Client) DeliverOneFeed(request *DeliverOneFeedRequest) (_result *DeliverOneFeedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeliverOneFeedHeaders{}
	_result = &DeliverOneFeedResponse{}
	_body, _err := client.DeliverOneFeedWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Feed推送
//
// @param request - PushFeedRequest
//
// @param headers - PushFeedHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushFeedResponse
func (client *Client) PushFeedWithOptions(request *PushFeedRequest, headers *PushFeedHeaders, runtime *util.RuntimeOptions) (_result *PushFeedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		body["expireTime"] = request.ExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.FeedFeature)) {
		body["feedFeature"] = request.FeedFeature
	}

	if !tea.BoolValue(util.IsUnset(request.IdempotentKey)) {
		body["idempotentKey"] = request.IdempotentKey
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
		Action:      tea.String("PushFeed"),
		Version:     tea.String("dingOne_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/dingOne/frame/feeds/pushFeed"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PushFeedResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Feed推送
//
// @param request - PushFeedRequest
//
// @return PushFeedResponse
func (client *Client) PushFeed(request *PushFeedRequest) (_result *PushFeedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PushFeedHeaders{}
	_result = &PushFeedResponse{}
	_body, _err := client.PushFeedWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新feed内容
//
// @param request - UpdateFeedContentRequest
//
// @param headers - UpdateFeedContentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFeedContentResponse
func (client *Client) UpdateFeedContentWithOptions(request *UpdateFeedContentRequest, headers *UpdateFeedContentHeaders, runtime *util.RuntimeOptions) (_result *UpdateFeedContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.FeedId)) {
		body["feedId"] = request.FeedId
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
		Action:      tea.String("UpdateFeedContent"),
		Version:     tea.String("dingOne_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/dingOne/frame/feeds/updateFeedContent"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFeedContentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新feed内容
//
// @param request - UpdateFeedContentRequest
//
// @return UpdateFeedContentResponse
func (client *Client) UpdateFeedContent(request *UpdateFeedContentRequest) (_result *UpdateFeedContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateFeedContentHeaders{}
	_result = &UpdateFeedContentResponse{}
	_body, _err := client.UpdateFeedContentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新feed过期时间
//
// @param request - UpdateFeedExpireTimeRequest
//
// @param headers - UpdateFeedExpireTimeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFeedExpireTimeResponse
func (client *Client) UpdateFeedExpireTimeWithOptions(request *UpdateFeedExpireTimeRequest, headers *UpdateFeedExpireTimeHeaders, runtime *util.RuntimeOptions) (_result *UpdateFeedExpireTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		body["expireTime"] = request.ExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.FeedId)) {
		body["feedId"] = request.FeedId
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
		Action:      tea.String("UpdateFeedExpireTime"),
		Version:     tea.String("dingOne_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/dingOne/frame/feeds/updateFeedExpireTime"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFeedExpireTimeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新feed过期时间
//
// @param request - UpdateFeedExpireTimeRequest
//
// @return UpdateFeedExpireTimeResponse
func (client *Client) UpdateFeedExpireTime(request *UpdateFeedExpireTimeRequest) (_result *UpdateFeedExpireTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateFeedExpireTimeHeaders{}
	_result = &UpdateFeedExpireTimeResponse{}
	_body, _err := client.UpdateFeedExpireTimeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
