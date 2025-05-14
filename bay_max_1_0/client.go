// This file is auto-generated, don't edit it. Thanks.
package bay_max_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type QueryBaymaxSkillLogHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryBaymaxSkillLogHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryBaymaxSkillLogHeaders) GoString() string {
	return s.String()
}

func (s *QueryBaymaxSkillLogHeaders) SetCommonHeaders(v map[string]*string) *QueryBaymaxSkillLogHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryBaymaxSkillLogHeaders) SetXAcsDingtalkAccessToken(v string) *QueryBaymaxSkillLogHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryBaymaxSkillLogRequest struct {
	From *int32 `json:"from,omitempty" xml:"from,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 14da****2760
	SkillExecuteId *string `json:"skillExecuteId,omitempty" xml:"skillExecuteId,omitempty"`
	To             *int32  `json:"to,omitempty" xml:"to,omitempty"`
}

func (s QueryBaymaxSkillLogRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBaymaxSkillLogRequest) GoString() string {
	return s.String()
}

func (s *QueryBaymaxSkillLogRequest) SetFrom(v int32) *QueryBaymaxSkillLogRequest {
	s.From = &v
	return s
}

func (s *QueryBaymaxSkillLogRequest) SetSkillExecuteId(v string) *QueryBaymaxSkillLogRequest {
	s.SkillExecuteId = &v
	return s
}

func (s *QueryBaymaxSkillLogRequest) SetTo(v int32) *QueryBaymaxSkillLogRequest {
	s.To = &v
	return s
}

type QueryBaymaxSkillLogResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 14da****2760
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s QueryBaymaxSkillLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBaymaxSkillLogResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBaymaxSkillLogResponseBody) SetResult(v string) *QueryBaymaxSkillLogResponseBody {
	s.Result = &v
	return s
}

type QueryBaymaxSkillLogResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBaymaxSkillLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBaymaxSkillLogResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBaymaxSkillLogResponse) GoString() string {
	return s.String()
}

func (s *QueryBaymaxSkillLogResponse) SetHeaders(v map[string]*string) *QueryBaymaxSkillLogResponse {
	s.Headers = v
	return s
}

func (s *QueryBaymaxSkillLogResponse) SetStatusCode(v int32) *QueryBaymaxSkillLogResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBaymaxSkillLogResponse) SetBody(v *QueryBaymaxSkillLogResponseBody) *QueryBaymaxSkillLogResponse {
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
// # Baymax技能执行日志
//
// @param request - QueryBaymaxSkillLogRequest
//
// @param headers - QueryBaymaxSkillLogHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBaymaxSkillLogResponse
func (client *Client) QueryBaymaxSkillLogWithOptions(request *QueryBaymaxSkillLogRequest, headers *QueryBaymaxSkillLogHeaders, runtime *util.RuntimeOptions) (_result *QueryBaymaxSkillLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.From)) {
		query["from"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.SkillExecuteId)) {
		query["skillExecuteId"] = request.SkillExecuteId
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		query["to"] = request.To
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
		Action:      tea.String("QueryBaymaxSkillLog"),
		Version:     tea.String("bayMax_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/bayMax/skills/logs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBaymaxSkillLogResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Baymax技能执行日志
//
// @param request - QueryBaymaxSkillLogRequest
//
// @return QueryBaymaxSkillLogResponse
func (client *Client) QueryBaymaxSkillLog(request *QueryBaymaxSkillLogRequest) (_result *QueryBaymaxSkillLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryBaymaxSkillLogHeaders{}
	_result = &QueryBaymaxSkillLogResponse{}
	_body, _err := client.QueryBaymaxSkillLogWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
