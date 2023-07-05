// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package blackboard_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type QueryBlackboardReadUnReadHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryBlackboardReadUnReadHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardReadUnReadHeaders) GoString() string {
	return s.String()
}

func (s *QueryBlackboardReadUnReadHeaders) SetCommonHeaders(v map[string]*string) *QueryBlackboardReadUnReadHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryBlackboardReadUnReadHeaders) SetXAcsDingtalkAccessToken(v string) *QueryBlackboardReadUnReadHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryBlackboardReadUnReadRequest struct {
	BlackboardId    *string `json:"blackboardId,omitempty" xml:"blackboardId,omitempty"`
	MaxResults      *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken       *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OperationUserId *string `json:"operationUserId,omitempty" xml:"operationUserId,omitempty"`
}

func (s QueryBlackboardReadUnReadRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardReadUnReadRequest) GoString() string {
	return s.String()
}

func (s *QueryBlackboardReadUnReadRequest) SetBlackboardId(v string) *QueryBlackboardReadUnReadRequest {
	s.BlackboardId = &v
	return s
}

func (s *QueryBlackboardReadUnReadRequest) SetMaxResults(v int32) *QueryBlackboardReadUnReadRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryBlackboardReadUnReadRequest) SetNextToken(v string) *QueryBlackboardReadUnReadRequest {
	s.NextToken = &v
	return s
}

func (s *QueryBlackboardReadUnReadRequest) SetOperationUserId(v string) *QueryBlackboardReadUnReadRequest {
	s.OperationUserId = &v
	return s
}

type QueryBlackboardReadUnReadResponseBody struct {
	NextToken *string                                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Users     []*QueryBlackboardReadUnReadResponseBodyUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s QueryBlackboardReadUnReadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardReadUnReadResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBlackboardReadUnReadResponseBody) SetNextToken(v string) *QueryBlackboardReadUnReadResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryBlackboardReadUnReadResponseBody) SetUsers(v []*QueryBlackboardReadUnReadResponseBodyUsers) *QueryBlackboardReadUnReadResponseBody {
	s.Users = v
	return s
}

type QueryBlackboardReadUnReadResponseBodyUsers struct {
	Read          *string `json:"read,omitempty" xml:"read,omitempty"`
	ReadTimestamp *int64  `json:"readTimestamp,omitempty" xml:"readTimestamp,omitempty"`
	UserId        *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryBlackboardReadUnReadResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardReadUnReadResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *QueryBlackboardReadUnReadResponseBodyUsers) SetRead(v string) *QueryBlackboardReadUnReadResponseBodyUsers {
	s.Read = &v
	return s
}

func (s *QueryBlackboardReadUnReadResponseBodyUsers) SetReadTimestamp(v int64) *QueryBlackboardReadUnReadResponseBodyUsers {
	s.ReadTimestamp = &v
	return s
}

func (s *QueryBlackboardReadUnReadResponseBodyUsers) SetUserId(v string) *QueryBlackboardReadUnReadResponseBodyUsers {
	s.UserId = &v
	return s
}

type QueryBlackboardReadUnReadResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryBlackboardReadUnReadResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBlackboardReadUnReadResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardReadUnReadResponse) GoString() string {
	return s.String()
}

func (s *QueryBlackboardReadUnReadResponse) SetHeaders(v map[string]*string) *QueryBlackboardReadUnReadResponse {
	s.Headers = v
	return s
}

func (s *QueryBlackboardReadUnReadResponse) SetStatusCode(v int32) *QueryBlackboardReadUnReadResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBlackboardReadUnReadResponse) SetBody(v *QueryBlackboardReadUnReadResponseBody) *QueryBlackboardReadUnReadResponse {
	s.Body = v
	return s
}

type QueryBlackboardSpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryBlackboardSpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardSpaceHeaders) GoString() string {
	return s.String()
}

func (s *QueryBlackboardSpaceHeaders) SetCommonHeaders(v map[string]*string) *QueryBlackboardSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryBlackboardSpaceHeaders) SetXAcsDingtalkAccessToken(v string) *QueryBlackboardSpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryBlackboardSpaceRequest struct {
	OperationUserId *string `json:"operationUserId,omitempty" xml:"operationUserId,omitempty"`
}

func (s QueryBlackboardSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardSpaceRequest) GoString() string {
	return s.String()
}

func (s *QueryBlackboardSpaceRequest) SetOperationUserId(v string) *QueryBlackboardSpaceRequest {
	s.OperationUserId = &v
	return s
}

type QueryBlackboardSpaceResponseBody struct {
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s QueryBlackboardSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBlackboardSpaceResponseBody) SetSpaceId(v string) *QueryBlackboardSpaceResponseBody {
	s.SpaceId = &v
	return s
}

type QueryBlackboardSpaceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryBlackboardSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBlackboardSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardSpaceResponse) GoString() string {
	return s.String()
}

func (s *QueryBlackboardSpaceResponse) SetHeaders(v map[string]*string) *QueryBlackboardSpaceResponse {
	s.Headers = v
	return s
}

func (s *QueryBlackboardSpaceResponse) SetStatusCode(v int32) *QueryBlackboardSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBlackboardSpaceResponse) SetBody(v *QueryBlackboardSpaceResponseBody) *QueryBlackboardSpaceResponse {
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

func (client *Client) QueryBlackboardReadUnReadWithOptions(request *QueryBlackboardReadUnReadRequest, headers *QueryBlackboardReadUnReadHeaders, runtime *util.RuntimeOptions) (_result *QueryBlackboardReadUnReadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BlackboardId)) {
		query["blackboardId"] = request.BlackboardId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OperationUserId)) {
		query["operationUserId"] = request.OperationUserId
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
		Action:      tea.String("QueryBlackboardReadUnRead"),
		Version:     tea.String("blackboard_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/blackboard/readers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBlackboardReadUnReadResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBlackboardReadUnRead(request *QueryBlackboardReadUnReadRequest) (_result *QueryBlackboardReadUnReadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryBlackboardReadUnReadHeaders{}
	_result = &QueryBlackboardReadUnReadResponse{}
	_body, _err := client.QueryBlackboardReadUnReadWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBlackboardSpaceWithOptions(request *QueryBlackboardSpaceRequest, headers *QueryBlackboardSpaceHeaders, runtime *util.RuntimeOptions) (_result *QueryBlackboardSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperationUserId)) {
		query["operationUserId"] = request.OperationUserId
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
		Action:      tea.String("QueryBlackboardSpace"),
		Version:     tea.String("blackboard_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/blackboard/spaces"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBlackboardSpaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBlackboardSpace(request *QueryBlackboardSpaceRequest) (_result *QueryBlackboardSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryBlackboardSpaceHeaders{}
	_result = &QueryBlackboardSpaceResponse{}
	_body, _err := client.QueryBlackboardSpaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
