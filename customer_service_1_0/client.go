// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package customer_service_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
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
	// 会员来源
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 第三方会员ID
	ForeignId *string `json:"foreignId,omitempty" xml:"foreignId,omitempty"`
	// 第三方会员名称
	ForeignName *string `json:"foreignName,omitempty" xml:"foreignName,omitempty"`
	// 实例ID
	OpenInstanceId *string `json:"openInstanceId,omitempty" xml:"openInstanceId,omitempty"`
	// 智能客服产品
	ProductionType *int32 `json:"productionType,omitempty" xml:"productionType,omitempty"`
	// 工单模板ID
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// 工单标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 工单表单
	Properties []*CreateTicketRequestProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Repeated"`
}

func (s CreateTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketRequest) GoString() string {
	return s.String()
}

func (s *CreateTicketRequest) SetSourceId(v string) *CreateTicketRequest {
	s.SourceId = &v
	return s
}

func (s *CreateTicketRequest) SetForeignId(v string) *CreateTicketRequest {
	s.ForeignId = &v
	return s
}

func (s *CreateTicketRequest) SetForeignName(v string) *CreateTicketRequest {
	s.ForeignName = &v
	return s
}

func (s *CreateTicketRequest) SetOpenInstanceId(v string) *CreateTicketRequest {
	s.OpenInstanceId = &v
	return s
}

func (s *CreateTicketRequest) SetProductionType(v int32) *CreateTicketRequest {
	s.ProductionType = &v
	return s
}

func (s *CreateTicketRequest) SetTemplateId(v string) *CreateTicketRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateTicketRequest) SetTitle(v string) *CreateTicketRequest {
	s.Title = &v
	return s
}

func (s *CreateTicketRequest) SetProperties(v []*CreateTicketRequestProperties) *CreateTicketRequest {
	s.Properties = v
	return s
}

type CreateTicketRequestProperties struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateTicketRequestProperties) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketRequestProperties) GoString() string {
	return s.String()
}

func (s *CreateTicketRequestProperties) SetName(v string) *CreateTicketRequestProperties {
	s.Name = &v
	return s
}

func (s *CreateTicketRequestProperties) SetValue(v string) *CreateTicketRequestProperties {
	s.Value = &v
	return s
}

type CreateTicketResponseBody struct {
	// 新创建工单ID
	TicketId *string `json:"ticketId,omitempty" xml:"ticketId,omitempty"`
}

func (s CreateTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTicketResponseBody) SetTicketId(v string) *CreateTicketResponseBody {
	s.TicketId = &v
	return s
}

type CreateTicketResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

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

func (client *Client) CreateTicketWithOptions(request *CreateTicketRequest, headers *CreateTicketHeaders, runtime *util.RuntimeOptions) (_result *CreateTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		body["sourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ForeignId)) {
		body["foreignId"] = request.ForeignId
	}

	if !tea.BoolValue(util.IsUnset(request.ForeignName)) {
		body["foreignName"] = request.ForeignName
	}

	if !tea.BoolValue(util.IsUnset(request.OpenInstanceId)) {
		body["openInstanceId"] = request.OpenInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductionType)) {
		body["productionType"] = request.ProductionType
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Properties)) {
		body["properties"] = request.Properties
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateTicketResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateTicket"), tea.String("customerService_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/customerService/tickets"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
