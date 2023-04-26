// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package watt_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CheckInCrowdsByMobileRequest struct {
	CrowdIds []byte  `json:"crowdIds,omitempty" xml:"crowdIds,omitempty"`
	Mobile   *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
}

func (s CheckInCrowdsByMobileRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckInCrowdsByMobileRequest) GoString() string {
	return s.String()
}

func (s *CheckInCrowdsByMobileRequest) SetCrowdIds(v []byte) *CheckInCrowdsByMobileRequest {
	s.CrowdIds = v
	return s
}

func (s *CheckInCrowdsByMobileRequest) SetMobile(v string) *CheckInCrowdsByMobileRequest {
	s.Mobile = &v
	return s
}

type CheckInCrowdsByMobileResponseBody struct {
	Data    *bool  `json:"data,omitempty" xml:"data,omitempty"`
	Success *bool  `json:"success,omitempty" xml:"success,omitempty"`
	Total   *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s CheckInCrowdsByMobileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckInCrowdsByMobileResponseBody) GoString() string {
	return s.String()
}

func (s *CheckInCrowdsByMobileResponseBody) SetData(v bool) *CheckInCrowdsByMobileResponseBody {
	s.Data = &v
	return s
}

func (s *CheckInCrowdsByMobileResponseBody) SetSuccess(v bool) *CheckInCrowdsByMobileResponseBody {
	s.Success = &v
	return s
}

func (s *CheckInCrowdsByMobileResponseBody) SetTotal(v int32) *CheckInCrowdsByMobileResponseBody {
	s.Total = &v
	return s
}

type CheckInCrowdsByMobileResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckInCrowdsByMobileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckInCrowdsByMobileResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckInCrowdsByMobileResponse) GoString() string {
	return s.String()
}

func (s *CheckInCrowdsByMobileResponse) SetHeaders(v map[string]*string) *CheckInCrowdsByMobileResponse {
	s.Headers = v
	return s
}

func (s *CheckInCrowdsByMobileResponse) SetStatusCode(v int32) *CheckInCrowdsByMobileResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckInCrowdsByMobileResponse) SetBody(v *CheckInCrowdsByMobileResponseBody) *CheckInCrowdsByMobileResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

func (client *Client) CheckInCrowdsByMobileWithOptions(request *CheckInCrowdsByMobileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CheckInCrowdsByMobileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CrowdIds)) {
		query["crowdIds"] = request.CrowdIds
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["mobile"] = request.Mobile
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckInCrowdsByMobile"),
		Version:     tea.String("watt_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/watt/crowdIdentifications/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckInCrowdsByMobileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckInCrowdsByMobile(request *CheckInCrowdsByMobileRequest) (_result *CheckInCrowdsByMobileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckInCrowdsByMobileResponse{}
	_body, _err := client.CheckInCrowdsByMobileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
