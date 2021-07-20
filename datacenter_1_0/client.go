// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package datacenter_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type QueryDigitalDistrictOrgInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDigitalDistrictOrgInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDigitalDistrictOrgInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryDigitalDistrictOrgInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryDigitalDistrictOrgInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDigitalDistrictOrgInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDigitalDistrictOrgInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDigitalDistrictOrgInfoRequest struct {
	KpiGroupId *string   `json:"kpiGroupId,omitempty" xml:"kpiGroupId,omitempty"`
	StatDates  []*string `json:"statDates,omitempty" xml:"statDates,omitempty" type:"Repeated"`
	OrgIds     []*string `json:"orgIds,omitempty" xml:"orgIds,omitempty" type:"Repeated"`
}

func (s QueryDigitalDistrictOrgInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDigitalDistrictOrgInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryDigitalDistrictOrgInfoRequest) SetKpiGroupId(v string) *QueryDigitalDistrictOrgInfoRequest {
	s.KpiGroupId = &v
	return s
}

func (s *QueryDigitalDistrictOrgInfoRequest) SetStatDates(v []*string) *QueryDigitalDistrictOrgInfoRequest {
	s.StatDates = v
	return s
}

func (s *QueryDigitalDistrictOrgInfoRequest) SetOrgIds(v []*string) *QueryDigitalDistrictOrgInfoRequest {
	s.OrgIds = v
	return s
}

type QueryDigitalDistrictOrgInfoResponseBody struct {
	// arguments
	Arguments []*string `json:"arguments,omitempty" xml:"arguments,omitempty" type:"Repeated"`
	// result
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s QueryDigitalDistrictOrgInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDigitalDistrictOrgInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDigitalDistrictOrgInfoResponseBody) SetArguments(v []*string) *QueryDigitalDistrictOrgInfoResponseBody {
	s.Arguments = v
	return s
}

func (s *QueryDigitalDistrictOrgInfoResponseBody) SetResult(v string) *QueryDigitalDistrictOrgInfoResponseBody {
	s.Result = &v
	return s
}

type QueryDigitalDistrictOrgInfoResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDigitalDistrictOrgInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDigitalDistrictOrgInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDigitalDistrictOrgInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryDigitalDistrictOrgInfoResponse) SetHeaders(v map[string]*string) *QueryDigitalDistrictOrgInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryDigitalDistrictOrgInfoResponse) SetBody(v *QueryDigitalDistrictOrgInfoResponseBody) *QueryDigitalDistrictOrgInfoResponse {
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

func (client *Client) QueryDigitalDistrictOrgInfo(request *QueryDigitalDistrictOrgInfoRequest) (_result *QueryDigitalDistrictOrgInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDigitalDistrictOrgInfoHeaders{}
	_result = &QueryDigitalDistrictOrgInfoResponse{}
	_body, _err := client.QueryDigitalDistrictOrgInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDigitalDistrictOrgInfoWithOptions(request *QueryDigitalDistrictOrgInfoRequest, headers *QueryDigitalDistrictOrgInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryDigitalDistrictOrgInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KpiGroupId)) {
		body["kpiGroupId"] = request.KpiGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StatDates)) {
		body["statDates"] = request.StatDates
	}

	if !tea.BoolValue(util.IsUnset(request.OrgIds)) {
		body["orgIds"] = request.OrgIds
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
	_result = &QueryDigitalDistrictOrgInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryDigitalDistrictOrgInfo"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/datacenter/digitalCounty/orgInfos/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
