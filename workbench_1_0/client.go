// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package workbench_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type QueryShortcutScopesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryShortcutScopesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryShortcutScopesHeaders) GoString() string {
	return s.String()
}

func (s *QueryShortcutScopesHeaders) SetCommonHeaders(v map[string]*string) *QueryShortcutScopesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryShortcutScopesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryShortcutScopesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryShortcutScopesResponseBody struct {
	// errorMsg
	UserVisibleScopes []*string `json:"userVisibleScopes,omitempty" xml:"userVisibleScopes,omitempty" type:"Repeated"`
	DeptVisibleScopes []*string `json:"deptVisibleScopes,omitempty" xml:"deptVisibleScopes,omitempty" type:"Repeated"`
}

func (s QueryShortcutScopesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryShortcutScopesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryShortcutScopesResponseBody) SetUserVisibleScopes(v []*string) *QueryShortcutScopesResponseBody {
	s.UserVisibleScopes = v
	return s
}

func (s *QueryShortcutScopesResponseBody) SetDeptVisibleScopes(v []*string) *QueryShortcutScopesResponseBody {
	s.DeptVisibleScopes = v
	return s
}

type QueryShortcutScopesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryShortcutScopesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryShortcutScopesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryShortcutScopesResponse) GoString() string {
	return s.String()
}

func (s *QueryShortcutScopesResponse) SetHeaders(v map[string]*string) *QueryShortcutScopesResponse {
	s.Headers = v
	return s
}

func (s *QueryShortcutScopesResponse) SetBody(v *QueryShortcutScopesResponseBody) *QueryShortcutScopesResponse {
	s.Body = v
	return s
}

type QueryComponentScopesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryComponentScopesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentScopesHeaders) GoString() string {
	return s.String()
}

func (s *QueryComponentScopesHeaders) SetCommonHeaders(v map[string]*string) *QueryComponentScopesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryComponentScopesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryComponentScopesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryComponentScopesResponseBody struct {
	// scopes
	UserVisibleScopes []*string `json:"userVisibleScopes,omitempty" xml:"userVisibleScopes,omitempty" type:"Repeated"`
	DeptVisibleScopes []*string `json:"deptVisibleScopes,omitempty" xml:"deptVisibleScopes,omitempty" type:"Repeated"`
}

func (s QueryComponentScopesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentScopesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryComponentScopesResponseBody) SetUserVisibleScopes(v []*string) *QueryComponentScopesResponseBody {
	s.UserVisibleScopes = v
	return s
}

func (s *QueryComponentScopesResponseBody) SetDeptVisibleScopes(v []*string) *QueryComponentScopesResponseBody {
	s.DeptVisibleScopes = v
	return s
}

type QueryComponentScopesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryComponentScopesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryComponentScopesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentScopesResponse) GoString() string {
	return s.String()
}

func (s *QueryComponentScopesResponse) SetHeaders(v map[string]*string) *QueryComponentScopesResponse {
	s.Headers = v
	return s
}

func (s *QueryComponentScopesResponse) SetBody(v *QueryComponentScopesResponseBody) *QueryComponentScopesResponse {
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

func (client *Client) QueryShortcutScopes(shortcutKey *string) (_result *QueryShortcutScopesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryShortcutScopesHeaders{}
	_result = &QueryShortcutScopesResponse{}
	_body, _err := client.QueryShortcutScopesWithOptions(shortcutKey, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryShortcutScopesWithOptions(shortcutKey *string, headers *QueryShortcutScopesHeaders, runtime *util.RuntimeOptions) (_result *QueryShortcutScopesResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &QueryShortcutScopesResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryShortcutScopes"), tea.String("workbench_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/workbench/shortcuts/"+tea.StringValue(shortcutKey)+"/scopes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryComponentScopes(componentId *string) (_result *QueryComponentScopesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryComponentScopesHeaders{}
	_result = &QueryComponentScopesResponse{}
	_body, _err := client.QueryComponentScopesWithOptions(componentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryComponentScopesWithOptions(componentId *string, headers *QueryComponentScopesHeaders, runtime *util.RuntimeOptions) (_result *QueryComponentScopesResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &QueryComponentScopesResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryComponentScopes"), tea.String("workbench_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/workbench/components/"+tea.StringValue(componentId)+"/scopes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
