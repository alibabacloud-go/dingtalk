// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package event_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetCallBackFaileResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCallBackFaileResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCallBackFaileResultHeaders) GoString() string {
	return s.String()
}

func (s *GetCallBackFaileResultHeaders) SetCommonHeaders(v map[string]*string) *GetCallBackFaileResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCallBackFaileResultHeaders) SetXAcsDingtalkAccessToken(v string) *GetCallBackFaileResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCallBackFaileResultRequest struct {
	BeginTime *int64 `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	EndTime   *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
}

func (s GetCallBackFaileResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCallBackFaileResultRequest) GoString() string {
	return s.String()
}

func (s *GetCallBackFaileResultRequest) SetBeginTime(v int64) *GetCallBackFaileResultRequest {
	s.BeginTime = &v
	return s
}

func (s *GetCallBackFaileResultRequest) SetEndTime(v int64) *GetCallBackFaileResultRequest {
	s.EndTime = &v
	return s
}

type GetCallBackFaileResultResponseBody struct {
	FailedList []*GetCallBackFaileResultResponseBodyFailedList `json:"failedList,omitempty" xml:"failedList,omitempty" type:"Repeated"`
	HasMore    *bool                                           `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
}

func (s GetCallBackFaileResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCallBackFaileResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetCallBackFaileResultResponseBody) SetFailedList(v []*GetCallBackFaileResultResponseBodyFailedList) *GetCallBackFaileResultResponseBody {
	s.FailedList = v
	return s
}

func (s *GetCallBackFaileResultResponseBody) SetHasMore(v bool) *GetCallBackFaileResultResponseBody {
	s.HasMore = &v
	return s
}

type GetCallBackFaileResultResponseBodyFailedList struct {
	CallBackData *string `json:"callBackData,omitempty" xml:"callBackData,omitempty"`
	CallBackTag  *string `json:"callBackTag,omitempty" xml:"callBackTag,omitempty"`
	CorpId       *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	EventTime    *int64  `json:"eventTime,omitempty" xml:"eventTime,omitempty"`
}

func (s GetCallBackFaileResultResponseBodyFailedList) String() string {
	return tea.Prettify(s)
}

func (s GetCallBackFaileResultResponseBodyFailedList) GoString() string {
	return s.String()
}

func (s *GetCallBackFaileResultResponseBodyFailedList) SetCallBackData(v string) *GetCallBackFaileResultResponseBodyFailedList {
	s.CallBackData = &v
	return s
}

func (s *GetCallBackFaileResultResponseBodyFailedList) SetCallBackTag(v string) *GetCallBackFaileResultResponseBodyFailedList {
	s.CallBackTag = &v
	return s
}

func (s *GetCallBackFaileResultResponseBodyFailedList) SetCorpId(v string) *GetCallBackFaileResultResponseBodyFailedList {
	s.CorpId = &v
	return s
}

func (s *GetCallBackFaileResultResponseBodyFailedList) SetEventTime(v int64) *GetCallBackFaileResultResponseBodyFailedList {
	s.EventTime = &v
	return s
}

type GetCallBackFaileResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCallBackFaileResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCallBackFaileResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCallBackFaileResultResponse) GoString() string {
	return s.String()
}

func (s *GetCallBackFaileResultResponse) SetHeaders(v map[string]*string) *GetCallBackFaileResultResponse {
	s.Headers = v
	return s
}

func (s *GetCallBackFaileResultResponse) SetStatusCode(v int32) *GetCallBackFaileResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCallBackFaileResultResponse) SetBody(v *GetCallBackFaileResultResponseBody) *GetCallBackFaileResultResponse {
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

func (client *Client) GetCallBackFaileResultWithOptions(request *GetCallBackFaileResultRequest, headers *GetCallBackFaileResultHeaders, runtime *util.RuntimeOptions) (_result *GetCallBackFaileResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["beginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
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
		Action:      tea.String("GetCallBackFaileResult"),
		Version:     tea.String("event_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/event/callbacks/failedResults"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCallBackFaileResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCallBackFaileResult(request *GetCallBackFaileResultRequest) (_result *GetCallBackFaileResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCallBackFaileResultHeaders{}
	_result = &GetCallBackFaileResultResponse{}
	_body, _err := client.GetCallBackFaileResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
