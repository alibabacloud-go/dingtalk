// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package workbench_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	DeptVisibleScopes []*int64  `json:"deptVisibleScopes,omitempty" xml:"deptVisibleScopes,omitempty" type:"Repeated"`
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

func (s *QueryComponentScopesResponseBody) SetDeptVisibleScopes(v []*int64) *QueryComponentScopesResponseBody {
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

type UpdateDingPortalPageScopeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateDingPortalPageScopeHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateDingPortalPageScopeHeaders) GoString() string {
	return s.String()
}

func (s *UpdateDingPortalPageScopeHeaders) SetCommonHeaders(v map[string]*string) *UpdateDingPortalPageScopeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateDingPortalPageScopeHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateDingPortalPageScopeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateDingPortalPageScopeRequest struct {
	// 可见用户列表
	Userids []*string `json:"userids,omitempty" xml:"userids,omitempty" type:"Repeated"`
	// 可见部门列表
	DeptIds []*int64 `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	// 可见角色列表
	RoleIds []*int64 `json:"roleIds,omitempty" xml:"roleIds,omitempty" type:"Repeated"`
	// 是否全员可见
	AllVisible *bool `json:"allVisible,omitempty" xml:"allVisible,omitempty"`
}

func (s UpdateDingPortalPageScopeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDingPortalPageScopeRequest) GoString() string {
	return s.String()
}

func (s *UpdateDingPortalPageScopeRequest) SetUserids(v []*string) *UpdateDingPortalPageScopeRequest {
	s.Userids = v
	return s
}

func (s *UpdateDingPortalPageScopeRequest) SetDeptIds(v []*int64) *UpdateDingPortalPageScopeRequest {
	s.DeptIds = v
	return s
}

func (s *UpdateDingPortalPageScopeRequest) SetRoleIds(v []*int64) *UpdateDingPortalPageScopeRequest {
	s.RoleIds = v
	return s
}

func (s *UpdateDingPortalPageScopeRequest) SetAllVisible(v bool) *UpdateDingPortalPageScopeRequest {
	s.AllVisible = &v
	return s
}

type UpdateDingPortalPageScopeResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpdateDingPortalPageScopeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDingPortalPageScopeResponse) GoString() string {
	return s.String()
}

func (s *UpdateDingPortalPageScopeResponse) SetHeaders(v map[string]*string) *UpdateDingPortalPageScopeResponse {
	s.Headers = v
	return s
}

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
	DeptVisibleScopes []*int64  `json:"deptVisibleScopes,omitempty" xml:"deptVisibleScopes,omitempty" type:"Repeated"`
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

func (s *QueryShortcutScopesResponseBody) SetDeptVisibleScopes(v []*int64) *QueryShortcutScopesResponseBody {
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

type GetDingPortalDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDingPortalDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDingPortalDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetDingPortalDetailHeaders) SetCommonHeaders(v map[string]*string) *GetDingPortalDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDingPortalDetailHeaders) SetXAcsDingtalkAccessToken(v string) *GetDingPortalDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDingPortalDetailResponseBody struct {
	// 工作台ID
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// 工作台名称
	DingPortalName *string `json:"dingPortalName,omitempty" xml:"dingPortalName,omitempty"`
	// 工作台页面信息
	Pages []*GetDingPortalDetailResponseBodyPages `json:"pages,omitempty" xml:"pages,omitempty" type:"Repeated"`
}

func (s GetDingPortalDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDingPortalDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetDingPortalDetailResponseBody) SetAppUuid(v string) *GetDingPortalDetailResponseBody {
	s.AppUuid = &v
	return s
}

func (s *GetDingPortalDetailResponseBody) SetDingPortalName(v string) *GetDingPortalDetailResponseBody {
	s.DingPortalName = &v
	return s
}

func (s *GetDingPortalDetailResponseBody) SetPages(v []*GetDingPortalDetailResponseBodyPages) *GetDingPortalDetailResponseBody {
	s.Pages = v
	return s
}

type GetDingPortalDetailResponseBodyPages struct {
	// 页面ID
	PageUuid *string `json:"pageUuid,omitempty" xml:"pageUuid,omitempty"`
	// 页面名称
	PageName *string `json:"pageName,omitempty" xml:"pageName,omitempty"`
	// 可见员工 ID 列表
	Userids []*string `json:"userids,omitempty" xml:"userids,omitempty" type:"Repeated"`
	// 可见部门 ID 铺
	DeptIds []*int64 `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	// 可见角色列表
	RoleIds []*int64 `json:"roleIds,omitempty" xml:"roleIds,omitempty" type:"Repeated"`
	// 是否全公司可见
	AllVisible *bool `json:"allVisible,omitempty" xml:"allVisible,omitempty"`
}

func (s GetDingPortalDetailResponseBodyPages) String() string {
	return tea.Prettify(s)
}

func (s GetDingPortalDetailResponseBodyPages) GoString() string {
	return s.String()
}

func (s *GetDingPortalDetailResponseBodyPages) SetPageUuid(v string) *GetDingPortalDetailResponseBodyPages {
	s.PageUuid = &v
	return s
}

func (s *GetDingPortalDetailResponseBodyPages) SetPageName(v string) *GetDingPortalDetailResponseBodyPages {
	s.PageName = &v
	return s
}

func (s *GetDingPortalDetailResponseBodyPages) SetUserids(v []*string) *GetDingPortalDetailResponseBodyPages {
	s.Userids = v
	return s
}

func (s *GetDingPortalDetailResponseBodyPages) SetDeptIds(v []*int64) *GetDingPortalDetailResponseBodyPages {
	s.DeptIds = v
	return s
}

func (s *GetDingPortalDetailResponseBodyPages) SetRoleIds(v []*int64) *GetDingPortalDetailResponseBodyPages {
	s.RoleIds = v
	return s
}

func (s *GetDingPortalDetailResponseBodyPages) SetAllVisible(v bool) *GetDingPortalDetailResponseBodyPages {
	s.AllVisible = &v
	return s
}

type GetDingPortalDetailResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDingPortalDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDingPortalDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDingPortalDetailResponse) GoString() string {
	return s.String()
}

func (s *GetDingPortalDetailResponse) SetHeaders(v map[string]*string) *GetDingPortalDetailResponse {
	s.Headers = v
	return s
}

func (s *GetDingPortalDetailResponse) SetBody(v *GetDingPortalDetailResponseBody) *GetDingPortalDetailResponse {
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

func (client *Client) UpdateDingPortalPageScope(pageUuid *string, appUuid *string, request *UpdateDingPortalPageScopeRequest) (_result *UpdateDingPortalPageScopeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateDingPortalPageScopeHeaders{}
	_result = &UpdateDingPortalPageScopeResponse{}
	_body, _err := client.UpdateDingPortalPageScopeWithOptions(pageUuid, appUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDingPortalPageScopeWithOptions(pageUuid *string, appUuid *string, request *UpdateDingPortalPageScopeRequest, headers *UpdateDingPortalPageScopeHeaders, runtime *util.RuntimeOptions) (_result *UpdateDingPortalPageScopeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Userids)) {
		body["userids"] = request.Userids
	}

	if !tea.BoolValue(util.IsUnset(request.DeptIds)) {
		body["deptIds"] = request.DeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.RoleIds)) {
		body["roleIds"] = request.RoleIds
	}

	if !tea.BoolValue(util.IsUnset(request.AllVisible)) {
		body["allVisible"] = request.AllVisible
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
	_result = &UpdateDingPortalPageScopeResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateDingPortalPageScope"), tea.String("workbench_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/workbench/dingPortals/"+tea.StringValue(appUuid)+"/pageScopes/"+tea.StringValue(pageUuid)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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

func (client *Client) GetDingPortalDetail(appUuid *string) (_result *GetDingPortalDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDingPortalDetailHeaders{}
	_result = &GetDingPortalDetailResponse{}
	_body, _err := client.GetDingPortalDetailWithOptions(appUuid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDingPortalDetailWithOptions(appUuid *string, headers *GetDingPortalDetailHeaders, runtime *util.RuntimeOptions) (_result *GetDingPortalDetailResponse, _err error) {
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
	_result = &GetDingPortalDetailResponse{}
	_body, _err := client.DoROARequest(tea.String("GetDingPortalDetail"), tea.String("workbench_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/workbench/dingPortals/"+tea.StringValue(appUuid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
