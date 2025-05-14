// This file is auto-generated, don't edit it. Thanks.
package vip_member_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type QueryVipMemberInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryVipMemberInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryVipMemberInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryVipMemberInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryVipMemberInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryVipMemberInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryVipMemberInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryVipMemberInfoRequest struct {
	ChannelCode *string `json:"channelCode,omitempty" xml:"channelCode,omitempty"`
}

func (s QueryVipMemberInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryVipMemberInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryVipMemberInfoRequest) SetChannelCode(v string) *QueryVipMemberInfoRequest {
	s.ChannelCode = &v
	return s
}

type QueryVipMemberInfoResponseBody struct {
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	IsVip      *bool   `json:"isVip,omitempty" xml:"isVip,omitempty"`
}

func (s QueryVipMemberInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryVipMemberInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVipMemberInfoResponseBody) SetExpireTime(v string) *QueryVipMemberInfoResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *QueryVipMemberInfoResponseBody) SetIsVip(v bool) *QueryVipMemberInfoResponseBody {
	s.IsVip = &v
	return s
}

type QueryVipMemberInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryVipMemberInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryVipMemberInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryVipMemberInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryVipMemberInfoResponse) SetHeaders(v map[string]*string) *QueryVipMemberInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryVipMemberInfoResponse) SetStatusCode(v int32) *QueryVipMemberInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVipMemberInfoResponse) SetBody(v *QueryVipMemberInfoResponseBody) *QueryVipMemberInfoResponse {
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
// 查询365会员信息
//
// @param request - QueryVipMemberInfoRequest
//
// @param headers - QueryVipMemberInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryVipMemberInfoResponse
func (client *Client) QueryVipMemberInfoWithOptions(request *QueryVipMemberInfoRequest, headers *QueryVipMemberInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryVipMemberInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelCode)) {
		body["channelCode"] = request.ChannelCode
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
		Action:      tea.String("QueryVipMemberInfo"),
		Version:     tea.String("vipMember_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/vipMember/users/memberInfos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryVipMemberInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询365会员信息
//
// @param request - QueryVipMemberInfoRequest
//
// @return QueryVipMemberInfoResponse
func (client *Client) QueryVipMemberInfo(request *QueryVipMemberInfoRequest) (_result *QueryVipMemberInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryVipMemberInfoHeaders{}
	_result = &QueryVipMemberInfoResponse{}
	_body, _err := client.QueryVipMemberInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
