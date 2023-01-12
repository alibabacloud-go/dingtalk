// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package chengfeng_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetEmployeeInfoByWorkNoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetEmployeeInfoByWorkNoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetEmployeeInfoByWorkNoHeaders) GoString() string {
	return s.String()
}

func (s *GetEmployeeInfoByWorkNoHeaders) SetCommonHeaders(v map[string]*string) *GetEmployeeInfoByWorkNoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetEmployeeInfoByWorkNoHeaders) SetXAcsDingtalkAccessToken(v string) *GetEmployeeInfoByWorkNoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetEmployeeInfoByWorkNoRequest struct {
	// 工号
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s GetEmployeeInfoByWorkNoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEmployeeInfoByWorkNoRequest) GoString() string {
	return s.String()
}

func (s *GetEmployeeInfoByWorkNoRequest) SetWorkNo(v string) *GetEmployeeInfoByWorkNoRequest {
	s.WorkNo = &v
	return s
}

type GetEmployeeInfoByWorkNoResponseBody struct {
	// 请求返回数据对象
	Content *GetEmployeeInfoByWorkNoResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	// 接口请求成功标识,成功为true,失败为false
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetEmployeeInfoByWorkNoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEmployeeInfoByWorkNoResponseBody) GoString() string {
	return s.String()
}

func (s *GetEmployeeInfoByWorkNoResponseBody) SetContent(v *GetEmployeeInfoByWorkNoResponseBodyContent) *GetEmployeeInfoByWorkNoResponseBody {
	s.Content = v
	return s
}

func (s *GetEmployeeInfoByWorkNoResponseBody) SetSuccess(v bool) *GetEmployeeInfoByWorkNoResponseBody {
	s.Success = &v
	return s
}

type GetEmployeeInfoByWorkNoResponseBodyContent struct {
	// 员工姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 员工工号
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s GetEmployeeInfoByWorkNoResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetEmployeeInfoByWorkNoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetEmployeeInfoByWorkNoResponseBodyContent) SetName(v string) *GetEmployeeInfoByWorkNoResponseBodyContent {
	s.Name = &v
	return s
}

func (s *GetEmployeeInfoByWorkNoResponseBodyContent) SetWorkNo(v string) *GetEmployeeInfoByWorkNoResponseBodyContent {
	s.WorkNo = &v
	return s
}

type GetEmployeeInfoByWorkNoResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEmployeeInfoByWorkNoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEmployeeInfoByWorkNoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEmployeeInfoByWorkNoResponse) GoString() string {
	return s.String()
}

func (s *GetEmployeeInfoByWorkNoResponse) SetHeaders(v map[string]*string) *GetEmployeeInfoByWorkNoResponse {
	s.Headers = v
	return s
}

func (s *GetEmployeeInfoByWorkNoResponse) SetBody(v *GetEmployeeInfoByWorkNoResponseBody) *GetEmployeeInfoByWorkNoResponse {
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

func (client *Client) GetEmployeeInfoByWorkNo(request *GetEmployeeInfoByWorkNoRequest) (_result *GetEmployeeInfoByWorkNoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetEmployeeInfoByWorkNoHeaders{}
	_result = &GetEmployeeInfoByWorkNoResponse{}
	_body, _err := client.GetEmployeeInfoByWorkNoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEmployeeInfoByWorkNoWithOptions(request *GetEmployeeInfoByWorkNoRequest, headers *GetEmployeeInfoByWorkNoHeaders, runtime *util.RuntimeOptions) (_result *GetEmployeeInfoByWorkNoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkNo)) {
		query["workNo"] = request.WorkNo
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
	_result = &GetEmployeeInfoByWorkNoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetEmployeeInfoByWorkNo"), tea.String("chengfeng_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/chengfeng/workNumbers/employees"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
