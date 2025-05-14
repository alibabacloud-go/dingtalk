// This file is auto-generated, don't edit it. Thanks.
package yun_shu_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type SaveOpenExternalLogHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveOpenExternalLogHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveOpenExternalLogHeaders) GoString() string {
	return s.String()
}

func (s *SaveOpenExternalLogHeaders) SetCommonHeaders(v map[string]*string) *SaveOpenExternalLogHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveOpenExternalLogHeaders) SetXAcsDingtalkAccessToken(v string) *SaveOpenExternalLogHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveOpenExternalLogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dingf8d907412a586
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yunshu
	LogSource *string `json:"logSource,omitempty" xml:"logSource,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// terminalInfo
	LogType *string `json:"logType,omitempty" xml:"logType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"date":"2023-05-10","macAddress":"34-2E-B7-AF-EA-JF","devSn":"68D1F0-B76A-5CC9-BCFC-BD7548BA","staffId":"05166141678164"}]
	OpenExt *string `json:"openExt,omitempty" xml:"openExt,omitempty"`
}

func (s SaveOpenExternalLogRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveOpenExternalLogRequest) GoString() string {
	return s.String()
}

func (s *SaveOpenExternalLogRequest) SetCorpId(v string) *SaveOpenExternalLogRequest {
	s.CorpId = &v
	return s
}

func (s *SaveOpenExternalLogRequest) SetLogSource(v string) *SaveOpenExternalLogRequest {
	s.LogSource = &v
	return s
}

func (s *SaveOpenExternalLogRequest) SetLogType(v string) *SaveOpenExternalLogRequest {
	s.LogType = &v
	return s
}

func (s *SaveOpenExternalLogRequest) SetOpenExt(v string) *SaveOpenExternalLogRequest {
	s.OpenExt = &v
	return s
}

type SaveOpenExternalLogResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SaveOpenExternalLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveOpenExternalLogResponseBody) GoString() string {
	return s.String()
}

func (s *SaveOpenExternalLogResponseBody) SetSuccess(v bool) *SaveOpenExternalLogResponseBody {
	s.Success = &v
	return s
}

type SaveOpenExternalLogResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveOpenExternalLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveOpenExternalLogResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveOpenExternalLogResponse) GoString() string {
	return s.String()
}

func (s *SaveOpenExternalLogResponse) SetHeaders(v map[string]*string) *SaveOpenExternalLogResponse {
	s.Headers = v
	return s
}

func (s *SaveOpenExternalLogResponse) SetStatusCode(v int32) *SaveOpenExternalLogResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveOpenExternalLogResponse) SetBody(v *SaveOpenExternalLogResponseBody) *SaveOpenExternalLogResponse {
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
// 生态日志数据互通
//
// @param request - SaveOpenExternalLogRequest
//
// @param headers - SaveOpenExternalLogHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveOpenExternalLogResponse
func (client *Client) SaveOpenExternalLogWithOptions(request *SaveOpenExternalLogRequest, headers *SaveOpenExternalLogHeaders, runtime *util.RuntimeOptions) (_result *SaveOpenExternalLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.LogSource)) {
		body["logSource"] = request.LogSource
	}

	if !tea.BoolValue(util.IsUnset(request.LogType)) {
		body["logType"] = request.LogType
	}

	if !tea.BoolValue(util.IsUnset(request.OpenExt)) {
		body["openExt"] = request.OpenExt
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
		Action:      tea.String("SaveOpenExternalLog"),
		Version:     tea.String("yunShu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/yunShu/externalLogs/save"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveOpenExternalLogResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生态日志数据互通
//
// @param request - SaveOpenExternalLogRequest
//
// @return SaveOpenExternalLogResponse
func (client *Client) SaveOpenExternalLog(request *SaveOpenExternalLogRequest) (_result *SaveOpenExternalLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveOpenExternalLogHeaders{}
	_result = &SaveOpenExternalLogResponse{}
	_body, _err := client.SaveOpenExternalLogWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
