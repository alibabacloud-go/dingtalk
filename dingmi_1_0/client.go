// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package dingmi_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetDingMeBaseDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDingMeBaseDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDingMeBaseDataHeaders) GoString() string {
	return s.String()
}

func (s *GetDingMeBaseDataHeaders) SetCommonHeaders(v map[string]*string) *GetDingMeBaseDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDingMeBaseDataHeaders) SetXAcsDingtalkAccessToken(v string) *GetDingMeBaseDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDingMeBaseDataRequest struct {
	// 机器人ID
	AppKey *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	// 开始时间
	StartDay *string `json:"startDay,omitempty" xml:"startDay,omitempty"`
	// 结束时间
	EndDay *string `json:"endDay,omitempty" xml:"endDay,omitempty"`
	// 是否按天分组
	ByDay *bool `json:"byDay,omitempty" xml:"byDay,omitempty"`
}

func (s GetDingMeBaseDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDingMeBaseDataRequest) GoString() string {
	return s.String()
}

func (s *GetDingMeBaseDataRequest) SetAppKey(v string) *GetDingMeBaseDataRequest {
	s.AppKey = &v
	return s
}

func (s *GetDingMeBaseDataRequest) SetStartDay(v string) *GetDingMeBaseDataRequest {
	s.StartDay = &v
	return s
}

func (s *GetDingMeBaseDataRequest) SetEndDay(v string) *GetDingMeBaseDataRequest {
	s.EndDay = &v
	return s
}

func (s *GetDingMeBaseDataRequest) SetByDay(v bool) *GetDingMeBaseDataRequest {
	s.ByDay = &v
	return s
}

type GetDingMeBaseDataResponseBody struct {
	// 是否缓存
	FromCache *bool `json:"fromCache,omitempty" xml:"fromCache,omitempty"`
	// 运行时间
	Runtime *int64 `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// 结果集
	Rawset []map[string]*string `json:"rawset,omitempty" xml:"rawset,omitempty" type:"Repeated"`
	// 字段解释
	Tips map[string]interface{} `json:"tips,omitempty" xml:"tips,omitempty"`
}

func (s GetDingMeBaseDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDingMeBaseDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetDingMeBaseDataResponseBody) SetFromCache(v bool) *GetDingMeBaseDataResponseBody {
	s.FromCache = &v
	return s
}

func (s *GetDingMeBaseDataResponseBody) SetRuntime(v int64) *GetDingMeBaseDataResponseBody {
	s.Runtime = &v
	return s
}

func (s *GetDingMeBaseDataResponseBody) SetRawset(v []map[string]*string) *GetDingMeBaseDataResponseBody {
	s.Rawset = v
	return s
}

func (s *GetDingMeBaseDataResponseBody) SetTips(v map[string]interface{}) *GetDingMeBaseDataResponseBody {
	s.Tips = v
	return s
}

type GetDingMeBaseDataResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDingMeBaseDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDingMeBaseDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDingMeBaseDataResponse) GoString() string {
	return s.String()
}

func (s *GetDingMeBaseDataResponse) SetHeaders(v map[string]*string) *GetDingMeBaseDataResponse {
	s.Headers = v
	return s
}

func (s *GetDingMeBaseDataResponse) SetBody(v *GetDingMeBaseDataResponseBody) *GetDingMeBaseDataResponse {
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

func (client *Client) GetDingMeBaseData(request *GetDingMeBaseDataRequest) (_result *GetDingMeBaseDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDingMeBaseDataHeaders{}
	_result = &GetDingMeBaseDataResponse{}
	_body, _err := client.GetDingMeBaseDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDingMeBaseDataWithOptions(request *GetDingMeBaseDataRequest, headers *GetDingMeBaseDataHeaders, runtime *util.RuntimeOptions) (_result *GetDingMeBaseDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["appKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.StartDay)) {
		query["startDay"] = request.StartDay
	}

	if !tea.BoolValue(util.IsUnset(request.EndDay)) {
		query["endDay"] = request.EndDay
	}

	if !tea.BoolValue(util.IsUnset(request.ByDay)) {
		query["byDay"] = request.ByDay
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
	_result = &GetDingMeBaseDataResponse{}
	_body, _err := client.DoROARequest(tea.String("GetDingMeBaseData"), tea.String("dingmi_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/dingmi/robots/data"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
