// This file is auto-generated, don't edit it. Thanks.
package tre_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketHeaders) GoString() string {
	return s.String()
}

func (s *CreateTicketHeaders) SetCommonHeaders(v map[string]*string) *CreateTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTicketHeaders) SetXAcsDingtalkAccessToken(v string) *CreateTicketHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateTicketRequest struct {
	// This parameter is required.
	Description *string                         `json:"description,omitempty" xml:"description,omitempty"`
	Priority    *string                         `json:"priority,omitempty" xml:"priority,omitempty"`
	Reporters   []*CreateTicketRequestReporters `json:"reporters,omitempty" xml:"reporters,omitempty" type:"Repeated"`
}

func (s CreateTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketRequest) GoString() string {
	return s.String()
}

func (s *CreateTicketRequest) SetDescription(v string) *CreateTicketRequest {
	s.Description = &v
	return s
}

func (s *CreateTicketRequest) SetPriority(v string) *CreateTicketRequest {
	s.Priority = &v
	return s
}

func (s *CreateTicketRequest) SetReporters(v []*CreateTicketRequestReporters) *CreateTicketRequest {
	s.Reporters = v
	return s
}

type CreateTicketRequestReporters struct {
	Carrier     *string   `json:"carrier,omitempty" xml:"carrier,omitempty"`
	Phone       *string   `json:"phone,omitempty" xml:"phone,omitempty"`
	Region      *string   `json:"region,omitempty" xml:"region,omitempty"`
	Role        *string   `json:"role,omitempty" xml:"role,omitempty"`
	Screenshots []*string `json:"screenshots,omitempty" xml:"screenshots,omitempty" type:"Repeated"`
	Timestamp   *int64    `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Uid         *string   `json:"uid,omitempty" xml:"uid,omitempty"`
	Version     *string   `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CreateTicketRequestReporters) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketRequestReporters) GoString() string {
	return s.String()
}

func (s *CreateTicketRequestReporters) SetCarrier(v string) *CreateTicketRequestReporters {
	s.Carrier = &v
	return s
}

func (s *CreateTicketRequestReporters) SetPhone(v string) *CreateTicketRequestReporters {
	s.Phone = &v
	return s
}

func (s *CreateTicketRequestReporters) SetRegion(v string) *CreateTicketRequestReporters {
	s.Region = &v
	return s
}

func (s *CreateTicketRequestReporters) SetRole(v string) *CreateTicketRequestReporters {
	s.Role = &v
	return s
}

func (s *CreateTicketRequestReporters) SetScreenshots(v []*string) *CreateTicketRequestReporters {
	s.Screenshots = v
	return s
}

func (s *CreateTicketRequestReporters) SetTimestamp(v int64) *CreateTicketRequestReporters {
	s.Timestamp = &v
	return s
}

func (s *CreateTicketRequestReporters) SetUid(v string) *CreateTicketRequestReporters {
	s.Uid = &v
	return s
}

func (s *CreateTicketRequestReporters) SetVersion(v string) *CreateTicketRequestReporters {
	s.Version = &v
	return s
}

type CreateTicketShrinkRequest struct {
	// This parameter is required.
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	Priority        *string `json:"priority,omitempty" xml:"priority,omitempty"`
	ReportersShrink *string `json:"reporters,omitempty" xml:"reporters,omitempty"`
}

func (s CreateTicketShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTicketShrinkRequest) SetDescription(v string) *CreateTicketShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateTicketShrinkRequest) SetPriority(v string) *CreateTicketShrinkRequest {
	s.Priority = &v
	return s
}

func (s *CreateTicketShrinkRequest) SetReportersShrink(v string) *CreateTicketShrinkRequest {
	s.ReportersShrink = &v
	return s
}

type CreateTicketResponseBody struct {
	Data      *CreateTicketResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTicketResponseBody) SetData(v *CreateTicketResponseBodyData) *CreateTicketResponseBody {
	s.Data = v
	return s
}

func (s *CreateTicketResponseBody) SetRequestId(v string) *CreateTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTicketResponseBody) SetSuccess(v bool) *CreateTicketResponseBody {
	s.Success = &v
	return s
}

type CreateTicketResponseBodyData struct {
	TicketId *string `json:"ticketId,omitempty" xml:"ticketId,omitempty"`
}

func (s CreateTicketResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTicketResponseBodyData) SetTicketId(v string) *CreateTicketResponseBodyData {
	s.TicketId = &v
	return s
}

type CreateTicketResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketResponse) GoString() string {
	return s.String()
}

func (s *CreateTicketResponse) SetHeaders(v map[string]*string) *CreateTicketResponse {
	s.Headers = v
	return s
}

func (s *CreateTicketResponse) SetStatusCode(v int32) *CreateTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTicketResponse) SetBody(v *CreateTicketResponseBody) *CreateTicketResponse {
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
// # TRE拉铃服务
//
// @param tmpReq - CreateTicketRequest
//
// @param headers - CreateTicketHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTicketResponse
func (client *Client) CreateTicketWithOptions(tmpReq *CreateTicketRequest, headers *CreateTicketHeaders, runtime *util.RuntimeOptions) (_result *CreateTicketResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateTicketShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Reporters)) {
		request.ReportersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Reporters, tea.String("reporters"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.ReportersShrink)) {
		body["reporters"] = request.ReportersShrink
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
		Action:      tea.String("CreateTicket"),
		Version:     tea.String("tre_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/tre/ticket"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTicketResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # TRE拉铃服务
//
// @param request - CreateTicketRequest
//
// @return CreateTicketResponse
func (client *Client) CreateTicket(request *CreateTicketRequest) (_result *CreateTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTicketHeaders{}
	_result = &CreateTicketResponse{}
	_body, _err := client.CreateTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
