// This file is auto-generated, don't edit it. Thanks.
package ding_one_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
