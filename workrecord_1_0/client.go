// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package workrecord_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CountWorkRecordHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CountWorkRecordHeaders) String() string {
	return tea.Prettify(s)
}

func (s CountWorkRecordHeaders) GoString() string {
	return s.String()
}

func (s *CountWorkRecordHeaders) SetCommonHeaders(v map[string]*string) *CountWorkRecordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CountWorkRecordHeaders) SetXAcsDingtalkAccessToken(v string) *CountWorkRecordHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CountWorkRecordRequest struct {
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CountWorkRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s CountWorkRecordRequest) GoString() string {
	return s.String()
}

func (s *CountWorkRecordRequest) SetUserId(v string) *CountWorkRecordRequest {
	s.UserId = &v
	return s
}

type CountWorkRecordResponseBody struct {
	UndoCount *int64 `json:"undoCount,omitempty" xml:"undoCount,omitempty"`
}

func (s CountWorkRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountWorkRecordResponseBody) GoString() string {
	return s.String()
}

func (s *CountWorkRecordResponseBody) SetUndoCount(v int64) *CountWorkRecordResponseBody {
	s.UndoCount = &v
	return s
}

type CountWorkRecordResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CountWorkRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CountWorkRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s CountWorkRecordResponse) GoString() string {
	return s.String()
}

func (s *CountWorkRecordResponse) SetHeaders(v map[string]*string) *CountWorkRecordResponse {
	s.Headers = v
	return s
}

func (s *CountWorkRecordResponse) SetStatusCode(v int32) *CountWorkRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *CountWorkRecordResponse) SetBody(v *CountWorkRecordResponseBody) *CountWorkRecordResponse {
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

func (client *Client) CountWorkRecordWithOptions(request *CountWorkRecordRequest, headers *CountWorkRecordHeaders, runtime *util.RuntimeOptions) (_result *CountWorkRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
		Action:      tea.String("CountWorkRecord"),
		Version:     tea.String("workrecord_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workrecord/counts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CountWorkRecordResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CountWorkRecord(request *CountWorkRecordRequest) (_result *CountWorkRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CountWorkRecordHeaders{}
	_result = &CountWorkRecordResponse{}
	_body, _err := client.CountWorkRecordWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
