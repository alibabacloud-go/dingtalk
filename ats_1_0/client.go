// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package ats_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetJobAuthHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetJobAuthHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetJobAuthHeaders) GoString() string {
	return s.String()
}

func (s *GetJobAuthHeaders) SetCommonHeaders(v map[string]*string) *GetJobAuthHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetJobAuthHeaders) SetXAcsDingtalkAccessToken(v string) *GetJobAuthHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetJobAuthRequest struct {
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s GetJobAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobAuthRequest) GoString() string {
	return s.String()
}

func (s *GetJobAuthRequest) SetOpUserId(v string) *GetJobAuthRequest {
	s.OpUserId = &v
	return s
}

type GetJobAuthResponseBody struct {
	// 职位ID
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// 职位负责人
	JobOwners []*GetJobAuthResponseBodyJobOwners `json:"jobOwners,omitempty" xml:"jobOwners,omitempty" type:"Repeated"`
}

func (s GetJobAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobAuthResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobAuthResponseBody) SetJobId(v string) *GetJobAuthResponseBody {
	s.JobId = &v
	return s
}

func (s *GetJobAuthResponseBody) SetJobOwners(v []*GetJobAuthResponseBodyJobOwners) *GetJobAuthResponseBody {
	s.JobOwners = v
	return s
}

type GetJobAuthResponseBodyJobOwners struct {
	// 员工标识
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 员工姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetJobAuthResponseBodyJobOwners) String() string {
	return tea.Prettify(s)
}

func (s GetJobAuthResponseBodyJobOwners) GoString() string {
	return s.String()
}

func (s *GetJobAuthResponseBodyJobOwners) SetUserId(v string) *GetJobAuthResponseBodyJobOwners {
	s.UserId = &v
	return s
}

func (s *GetJobAuthResponseBodyJobOwners) SetName(v string) *GetJobAuthResponseBodyJobOwners {
	s.Name = &v
	return s
}

type GetJobAuthResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetJobAuthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobAuthResponse) GoString() string {
	return s.String()
}

func (s *GetJobAuthResponse) SetHeaders(v map[string]*string) *GetJobAuthResponse {
	s.Headers = v
	return s
}

func (s *GetJobAuthResponse) SetBody(v *GetJobAuthResponseBody) *GetJobAuthResponse {
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

func (client *Client) GetJobAuth(jobId *string, request *GetJobAuthRequest) (_result *GetJobAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetJobAuthHeaders{}
	_result = &GetJobAuthResponse{}
	_body, _err := client.GetJobAuthWithOptions(jobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobAuthWithOptions(jobId *string, request *GetJobAuthRequest, headers *GetJobAuthHeaders, runtime *util.RuntimeOptions) (_result *GetJobAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
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
		Query:   openapiutil.Query(query),
	}
	_result = &GetJobAuthResponse{}
	_body, _err := client.DoROARequest(tea.String("GetJobAuth"), tea.String("ats_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/ats/auths/jobs/"+tea.StringValue(jobId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
