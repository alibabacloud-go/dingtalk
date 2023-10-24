// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package cool_ops_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type UpdateIsvOppStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateIsvOppStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateIsvOppStatusHeaders) GoString() string {
	return s.String()
}

func (s *UpdateIsvOppStatusHeaders) SetCommonHeaders(v map[string]*string) *UpdateIsvOppStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateIsvOppStatusHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateIsvOppStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateIsvOppStatusRequest struct {
	IsvOpportunityStatusList []*UpdateIsvOppStatusRequestIsvOpportunityStatusList `json:"isvOpportunityStatusList,omitempty" xml:"isvOpportunityStatusList,omitempty" type:"Repeated"`
}

func (s UpdateIsvOppStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIsvOppStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateIsvOppStatusRequest) SetIsvOpportunityStatusList(v []*UpdateIsvOppStatusRequestIsvOpportunityStatusList) *UpdateIsvOppStatusRequest {
	s.IsvOpportunityStatusList = v
	return s
}

type UpdateIsvOppStatusRequestIsvOpportunityStatusList struct {
	IsvCorpId         *string `json:"isvCorpId,omitempty" xml:"isvCorpId,omitempty"`
	MicroAppId        *string `json:"microAppId,omitempty" xml:"microAppId,omitempty"`
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	Note              *string `json:"note,omitempty" xml:"note,omitempty"`
	OperName          *string `json:"operName,omitempty" xml:"operName,omitempty"`
	OperTime          *string `json:"operTime,omitempty" xml:"operTime,omitempty"`
	OperUserId        *string `json:"operUserId,omitempty" xml:"operUserId,omitempty"`
	OppSourceCorpId   *string `json:"oppSourceCorpId,omitempty" xml:"oppSourceCorpId,omitempty"`
	OpportunityStatus *string `json:"opportunityStatus,omitempty" xml:"opportunityStatus,omitempty"`
	UserId            *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateIsvOppStatusRequestIsvOpportunityStatusList) String() string {
	return tea.Prettify(s)
}

func (s UpdateIsvOppStatusRequestIsvOpportunityStatusList) GoString() string {
	return s.String()
}

func (s *UpdateIsvOppStatusRequestIsvOpportunityStatusList) SetIsvCorpId(v string) *UpdateIsvOppStatusRequestIsvOpportunityStatusList {
	s.IsvCorpId = &v
	return s
}

func (s *UpdateIsvOppStatusRequestIsvOpportunityStatusList) SetMicroAppId(v string) *UpdateIsvOppStatusRequestIsvOpportunityStatusList {
	s.MicroAppId = &v
	return s
}

func (s *UpdateIsvOppStatusRequestIsvOpportunityStatusList) SetName(v string) *UpdateIsvOppStatusRequestIsvOpportunityStatusList {
	s.Name = &v
	return s
}

func (s *UpdateIsvOppStatusRequestIsvOpportunityStatusList) SetNote(v string) *UpdateIsvOppStatusRequestIsvOpportunityStatusList {
	s.Note = &v
	return s
}

func (s *UpdateIsvOppStatusRequestIsvOpportunityStatusList) SetOperName(v string) *UpdateIsvOppStatusRequestIsvOpportunityStatusList {
	s.OperName = &v
	return s
}

func (s *UpdateIsvOppStatusRequestIsvOpportunityStatusList) SetOperTime(v string) *UpdateIsvOppStatusRequestIsvOpportunityStatusList {
	s.OperTime = &v
	return s
}

func (s *UpdateIsvOppStatusRequestIsvOpportunityStatusList) SetOperUserId(v string) *UpdateIsvOppStatusRequestIsvOpportunityStatusList {
	s.OperUserId = &v
	return s
}

func (s *UpdateIsvOppStatusRequestIsvOpportunityStatusList) SetOppSourceCorpId(v string) *UpdateIsvOppStatusRequestIsvOpportunityStatusList {
	s.OppSourceCorpId = &v
	return s
}

func (s *UpdateIsvOppStatusRequestIsvOpportunityStatusList) SetOpportunityStatus(v string) *UpdateIsvOppStatusRequestIsvOpportunityStatusList {
	s.OpportunityStatus = &v
	return s
}

func (s *UpdateIsvOppStatusRequestIsvOpportunityStatusList) SetUserId(v string) *UpdateIsvOppStatusRequestIsvOpportunityStatusList {
	s.UserId = &v
	return s
}

type UpdateIsvOppStatusResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateIsvOppStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIsvOppStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIsvOppStatusResponseBody) SetResult(v bool) *UpdateIsvOppStatusResponseBody {
	s.Result = &v
	return s
}

type UpdateIsvOppStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateIsvOppStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateIsvOppStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIsvOppStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateIsvOppStatusResponse) SetHeaders(v map[string]*string) *UpdateIsvOppStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateIsvOppStatusResponse) SetStatusCode(v int32) *UpdateIsvOppStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIsvOppStatusResponse) SetBody(v *UpdateIsvOppStatusResponseBody) *UpdateIsvOppStatusResponse {
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

func (client *Client) UpdateIsvOppStatusWithOptions(request *UpdateIsvOppStatusRequest, headers *UpdateIsvOppStatusHeaders, runtime *util.RuntimeOptions) (_result *UpdateIsvOppStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsvOpportunityStatusList)) {
		body["isvOpportunityStatusList"] = request.IsvOpportunityStatusList
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
		Action:      tea.String("UpdateIsvOppStatus"),
		Version:     tea.String("coolOps_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/coolOps/isvOpportunities/statuses"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIsvOppStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateIsvOppStatus(request *UpdateIsvOppStatusRequest) (_result *UpdateIsvOppStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateIsvOppStatusHeaders{}
	_result = &UpdateIsvOppStatusResponse{}
	_body, _err := client.UpdateIsvOppStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
