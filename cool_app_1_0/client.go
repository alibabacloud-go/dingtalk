// This file is auto-generated, don't edit it. Thanks.
package cool_app_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type InstallCoolAppOrderToGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InstallCoolAppOrderToGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s InstallCoolAppOrderToGroupHeaders) GoString() string {
	return s.String()
}

func (s *InstallCoolAppOrderToGroupHeaders) SetCommonHeaders(v map[string]*string) *InstallCoolAppOrderToGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InstallCoolAppOrderToGroupHeaders) SetXAcsDingtalkAccessToken(v string) *InstallCoolAppOrderToGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InstallCoolAppOrderToGroupRequest struct {
	// example:
	//
	// cidxxx
	ConversationId     *string  `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	SortedPluginIdList []*int64 `json:"sortedPluginIdList,omitempty" xml:"sortedPluginIdList,omitempty" type:"Repeated"`
	// example:
	//
	// template-id-xxx
	TemplateId           *string  `json:"templateId,omitempty" xml:"templateId,omitempty"`
	UnsortedPluginIdList []*int64 `json:"unsortedPluginIdList,omitempty" xml:"unsortedPluginIdList,omitempty" type:"Repeated"`
}

func (s InstallCoolAppOrderToGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallCoolAppOrderToGroupRequest) GoString() string {
	return s.String()
}

func (s *InstallCoolAppOrderToGroupRequest) SetConversationId(v string) *InstallCoolAppOrderToGroupRequest {
	s.ConversationId = &v
	return s
}

func (s *InstallCoolAppOrderToGroupRequest) SetSortedPluginIdList(v []*int64) *InstallCoolAppOrderToGroupRequest {
	s.SortedPluginIdList = v
	return s
}

func (s *InstallCoolAppOrderToGroupRequest) SetTemplateId(v string) *InstallCoolAppOrderToGroupRequest {
	s.TemplateId = &v
	return s
}

func (s *InstallCoolAppOrderToGroupRequest) SetUnsortedPluginIdList(v []*int64) *InstallCoolAppOrderToGroupRequest {
	s.UnsortedPluginIdList = v
	return s
}

type InstallCoolAppOrderToGroupResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s InstallCoolAppOrderToGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallCoolAppOrderToGroupResponseBody) GoString() string {
	return s.String()
}

func (s *InstallCoolAppOrderToGroupResponseBody) SetResult(v string) *InstallCoolAppOrderToGroupResponseBody {
	s.Result = &v
	return s
}

func (s *InstallCoolAppOrderToGroupResponseBody) SetSuccess(v bool) *InstallCoolAppOrderToGroupResponseBody {
	s.Success = &v
	return s
}

type InstallCoolAppOrderToGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallCoolAppOrderToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallCoolAppOrderToGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallCoolAppOrderToGroupResponse) GoString() string {
	return s.String()
}

func (s *InstallCoolAppOrderToGroupResponse) SetHeaders(v map[string]*string) *InstallCoolAppOrderToGroupResponse {
	s.Headers = v
	return s
}

func (s *InstallCoolAppOrderToGroupResponse) SetStatusCode(v int32) *InstallCoolAppOrderToGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallCoolAppOrderToGroupResponse) SetBody(v *InstallCoolAppOrderToGroupResponseBody) *InstallCoolAppOrderToGroupResponse {
	s.Body = v
	return s
}

type InstallCoolAppToGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InstallCoolAppToGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s InstallCoolAppToGroupHeaders) GoString() string {
	return s.String()
}

func (s *InstallCoolAppToGroupHeaders) SetCommonHeaders(v map[string]*string) *InstallCoolAppToGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InstallCoolAppToGroupHeaders) SetXAcsDingtalkAccessToken(v string) *InstallCoolAppToGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InstallCoolAppToGroupRequest struct {
	// example:
	//
	// cidxxxx
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// example:
	//
	// CoolApp-xxx
	OperateCoolAppCode *string `json:"operateCoolAppCode,omitempty" xml:"operateCoolAppCode,omitempty"`
	// example:
	//
	// staffid12
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// example:
	//
	// template-id-xxx
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s InstallCoolAppToGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallCoolAppToGroupRequest) GoString() string {
	return s.String()
}

func (s *InstallCoolAppToGroupRequest) SetConversationId(v string) *InstallCoolAppToGroupRequest {
	s.ConversationId = &v
	return s
}

func (s *InstallCoolAppToGroupRequest) SetOperateCoolAppCode(v string) *InstallCoolAppToGroupRequest {
	s.OperateCoolAppCode = &v
	return s
}

func (s *InstallCoolAppToGroupRequest) SetOperatorId(v string) *InstallCoolAppToGroupRequest {
	s.OperatorId = &v
	return s
}

func (s *InstallCoolAppToGroupRequest) SetTemplateId(v string) *InstallCoolAppToGroupRequest {
	s.TemplateId = &v
	return s
}

type InstallCoolAppToGroupResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s InstallCoolAppToGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallCoolAppToGroupResponseBody) GoString() string {
	return s.String()
}

func (s *InstallCoolAppToGroupResponseBody) SetResult(v string) *InstallCoolAppToGroupResponseBody {
	s.Result = &v
	return s
}

func (s *InstallCoolAppToGroupResponseBody) SetSuccess(v bool) *InstallCoolAppToGroupResponseBody {
	s.Success = &v
	return s
}

type InstallCoolAppToGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallCoolAppToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallCoolAppToGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallCoolAppToGroupResponse) GoString() string {
	return s.String()
}

func (s *InstallCoolAppToGroupResponse) SetHeaders(v map[string]*string) *InstallCoolAppToGroupResponse {
	s.Headers = v
	return s
}

func (s *InstallCoolAppToGroupResponse) SetStatusCode(v int32) *InstallCoolAppToGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallCoolAppToGroupResponse) SetBody(v *InstallCoolAppToGroupResponseBody) *InstallCoolAppToGroupResponse {
	s.Body = v
	return s
}

type QueryCoolAppShortcutOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCoolAppShortcutOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCoolAppShortcutOrderHeaders) GoString() string {
	return s.String()
}

func (s *QueryCoolAppShortcutOrderHeaders) SetCommonHeaders(v map[string]*string) *QueryCoolAppShortcutOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCoolAppShortcutOrderHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCoolAppShortcutOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCoolAppShortcutOrderRequest struct {
	// example:
	//
	// cidxxx
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// example:
	//
	// staff1
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// example:
	//
	// templateId1
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s QueryCoolAppShortcutOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCoolAppShortcutOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryCoolAppShortcutOrderRequest) SetConversationId(v string) *QueryCoolAppShortcutOrderRequest {
	s.ConversationId = &v
	return s
}

func (s *QueryCoolAppShortcutOrderRequest) SetOperatorId(v string) *QueryCoolAppShortcutOrderRequest {
	s.OperatorId = &v
	return s
}

func (s *QueryCoolAppShortcutOrderRequest) SetTemplateId(v string) *QueryCoolAppShortcutOrderRequest {
	s.TemplateId = &v
	return s
}

type QueryCoolAppShortcutOrderResponseBody struct {
	Result  *QueryCoolAppShortcutOrderResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                        `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryCoolAppShortcutOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCoolAppShortcutOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCoolAppShortcutOrderResponseBody) SetResult(v *QueryCoolAppShortcutOrderResponseBodyResult) *QueryCoolAppShortcutOrderResponseBody {
	s.Result = v
	return s
}

func (s *QueryCoolAppShortcutOrderResponseBody) SetSuccess(v bool) *QueryCoolAppShortcutOrderResponseBody {
	s.Success = &v
	return s
}

type QueryCoolAppShortcutOrderResponseBodyResult struct {
	ForbiddenPluginList []*QueryCoolAppShortcutOrderResponseBodyResultForbiddenPluginList `json:"forbiddenPluginList,omitempty" xml:"forbiddenPluginList,omitempty" type:"Repeated"`
	MyPluginList        []*QueryCoolAppShortcutOrderResponseBodyResultMyPluginList        `json:"myPluginList,omitempty" xml:"myPluginList,omitempty" type:"Repeated"`
	OtherPluginList     []*QueryCoolAppShortcutOrderResponseBodyResultOtherPluginList     `json:"otherPluginList,omitempty" xml:"otherPluginList,omitempty" type:"Repeated"`
}

func (s QueryCoolAppShortcutOrderResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryCoolAppShortcutOrderResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryCoolAppShortcutOrderResponseBodyResult) SetForbiddenPluginList(v []*QueryCoolAppShortcutOrderResponseBodyResultForbiddenPluginList) *QueryCoolAppShortcutOrderResponseBodyResult {
	s.ForbiddenPluginList = v
	return s
}

func (s *QueryCoolAppShortcutOrderResponseBodyResult) SetMyPluginList(v []*QueryCoolAppShortcutOrderResponseBodyResultMyPluginList) *QueryCoolAppShortcutOrderResponseBodyResult {
	s.MyPluginList = v
	return s
}

func (s *QueryCoolAppShortcutOrderResponseBodyResult) SetOtherPluginList(v []*QueryCoolAppShortcutOrderResponseBodyResultOtherPluginList) *QueryCoolAppShortcutOrderResponseBodyResult {
	s.OtherPluginList = v
	return s
}

type QueryCoolAppShortcutOrderResponseBodyResultForbiddenPluginList struct {
	AppCode  *string `json:"appCode,omitempty" xml:"appCode,omitempty"`
	Desc     *string `json:"desc,omitempty" xml:"desc,omitempty"`
	PluginId *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	Source   *string `json:"source,omitempty" xml:"source,omitempty"`
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryCoolAppShortcutOrderResponseBodyResultForbiddenPluginList) String() string {
	return tea.Prettify(s)
}

func (s QueryCoolAppShortcutOrderResponseBodyResultForbiddenPluginList) GoString() string {
	return s.String()
}

func (s *QueryCoolAppShortcutOrderResponseBodyResultForbiddenPluginList) SetAppCode(v string) *QueryCoolAppShortcutOrderResponseBodyResultForbiddenPluginList {
	s.AppCode = &v
	return s
}

func (s *QueryCoolAppShortcutOrderResponseBodyResultForbiddenPluginList) SetDesc(v string) *QueryCoolAppShortcutOrderResponseBodyResultForbiddenPluginList {
	s.Desc = &v
	return s
}

func (s *QueryCoolAppShortcutOrderResponseBodyResultForbiddenPluginList) SetPluginId(v string) *QueryCoolAppShortcutOrderResponseBodyResultForbiddenPluginList {
	s.PluginId = &v
	return s
}

func (s *QueryCoolAppShortcutOrderResponseBodyResultForbiddenPluginList) SetSource(v string) *QueryCoolAppShortcutOrderResponseBodyResultForbiddenPluginList {
	s.Source = &v
	return s
}

func (s *QueryCoolAppShortcutOrderResponseBodyResultForbiddenPluginList) SetTitle(v string) *QueryCoolAppShortcutOrderResponseBodyResultForbiddenPluginList {
	s.Title = &v
	return s
}

type QueryCoolAppShortcutOrderResponseBodyResultMyPluginList struct {
	AppCode  *string `json:"appCode,omitempty" xml:"appCode,omitempty"`
	Desc     *string `json:"desc,omitempty" xml:"desc,omitempty"`
	PluginId *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	Source   *string `json:"source,omitempty" xml:"source,omitempty"`
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryCoolAppShortcutOrderResponseBodyResultMyPluginList) String() string {
	return tea.Prettify(s)
}

func (s QueryCoolAppShortcutOrderResponseBodyResultMyPluginList) GoString() string {
	return s.String()
}

func (s *QueryCoolAppShortcutOrderResponseBodyResultMyPluginList) SetAppCode(v string) *QueryCoolAppShortcutOrderResponseBodyResultMyPluginList {
	s.AppCode = &v
	return s
}

func (s *QueryCoolAppShortcutOrderResponseBodyResultMyPluginList) SetDesc(v string) *QueryCoolAppShortcutOrderResponseBodyResultMyPluginList {
	s.Desc = &v
	return s
}

func (s *QueryCoolAppShortcutOrderResponseBodyResultMyPluginList) SetPluginId(v string) *QueryCoolAppShortcutOrderResponseBodyResultMyPluginList {
	s.PluginId = &v
	return s
}

func (s *QueryCoolAppShortcutOrderResponseBodyResultMyPluginList) SetSource(v string) *QueryCoolAppShortcutOrderResponseBodyResultMyPluginList {
	s.Source = &v
	return s
}

func (s *QueryCoolAppShortcutOrderResponseBodyResultMyPluginList) SetTitle(v string) *QueryCoolAppShortcutOrderResponseBodyResultMyPluginList {
	s.Title = &v
	return s
}

type QueryCoolAppShortcutOrderResponseBodyResultOtherPluginList struct {
	AppCode  *string `json:"appCode,omitempty" xml:"appCode,omitempty"`
	Desc     *string `json:"desc,omitempty" xml:"desc,omitempty"`
	PluginId *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	Source   *string `json:"source,omitempty" xml:"source,omitempty"`
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryCoolAppShortcutOrderResponseBodyResultOtherPluginList) String() string {
	return tea.Prettify(s)
}

func (s QueryCoolAppShortcutOrderResponseBodyResultOtherPluginList) GoString() string {
	return s.String()
}

func (s *QueryCoolAppShortcutOrderResponseBodyResultOtherPluginList) SetAppCode(v string) *QueryCoolAppShortcutOrderResponseBodyResultOtherPluginList {
	s.AppCode = &v
	return s
}

func (s *QueryCoolAppShortcutOrderResponseBodyResultOtherPluginList) SetDesc(v string) *QueryCoolAppShortcutOrderResponseBodyResultOtherPluginList {
	s.Desc = &v
	return s
}

func (s *QueryCoolAppShortcutOrderResponseBodyResultOtherPluginList) SetPluginId(v string) *QueryCoolAppShortcutOrderResponseBodyResultOtherPluginList {
	s.PluginId = &v
	return s
}

func (s *QueryCoolAppShortcutOrderResponseBodyResultOtherPluginList) SetSource(v string) *QueryCoolAppShortcutOrderResponseBodyResultOtherPluginList {
	s.Source = &v
	return s
}

func (s *QueryCoolAppShortcutOrderResponseBodyResultOtherPluginList) SetTitle(v string) *QueryCoolAppShortcutOrderResponseBodyResultOtherPluginList {
	s.Title = &v
	return s
}

type QueryCoolAppShortcutOrderResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCoolAppShortcutOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCoolAppShortcutOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCoolAppShortcutOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryCoolAppShortcutOrderResponse) SetHeaders(v map[string]*string) *QueryCoolAppShortcutOrderResponse {
	s.Headers = v
	return s
}

func (s *QueryCoolAppShortcutOrderResponse) SetStatusCode(v int32) *QueryCoolAppShortcutOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCoolAppShortcutOrderResponse) SetBody(v *QueryCoolAppShortcutOrderResponseBody) *QueryCoolAppShortcutOrderResponse {
	s.Body = v
	return s
}

type UninstallCoolAppFromGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UninstallCoolAppFromGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s UninstallCoolAppFromGroupHeaders) GoString() string {
	return s.String()
}

func (s *UninstallCoolAppFromGroupHeaders) SetCommonHeaders(v map[string]*string) *UninstallCoolAppFromGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UninstallCoolAppFromGroupHeaders) SetXAcsDingtalkAccessToken(v string) *UninstallCoolAppFromGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UninstallCoolAppFromGroupRequest struct {
	// example:
	//
	// cidxxx
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// example:
	//
	// CoolApp-xxx
	OperateCoolAppCode *string `json:"operateCoolAppCode,omitempty" xml:"operateCoolAppCode,omitempty"`
	// example:
	//
	// staffid111
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// example:
	//
	// template-id-xxx
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s UninstallCoolAppFromGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UninstallCoolAppFromGroupRequest) GoString() string {
	return s.String()
}

func (s *UninstallCoolAppFromGroupRequest) SetConversationId(v string) *UninstallCoolAppFromGroupRequest {
	s.ConversationId = &v
	return s
}

func (s *UninstallCoolAppFromGroupRequest) SetOperateCoolAppCode(v string) *UninstallCoolAppFromGroupRequest {
	s.OperateCoolAppCode = &v
	return s
}

func (s *UninstallCoolAppFromGroupRequest) SetOperatorId(v string) *UninstallCoolAppFromGroupRequest {
	s.OperatorId = &v
	return s
}

func (s *UninstallCoolAppFromGroupRequest) SetTemplateId(v string) *UninstallCoolAppFromGroupRequest {
	s.TemplateId = &v
	return s
}

type UninstallCoolAppFromGroupResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UninstallCoolAppFromGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UninstallCoolAppFromGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallCoolAppFromGroupResponseBody) SetResult(v string) *UninstallCoolAppFromGroupResponseBody {
	s.Result = &v
	return s
}

func (s *UninstallCoolAppFromGroupResponseBody) SetSuccess(v bool) *UninstallCoolAppFromGroupResponseBody {
	s.Success = &v
	return s
}

type UninstallCoolAppFromGroupResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallCoolAppFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallCoolAppFromGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UninstallCoolAppFromGroupResponse) GoString() string {
	return s.String()
}

func (s *UninstallCoolAppFromGroupResponse) SetHeaders(v map[string]*string) *UninstallCoolAppFromGroupResponse {
	s.Headers = v
	return s
}

func (s *UninstallCoolAppFromGroupResponse) SetStatusCode(v int32) *UninstallCoolAppFromGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallCoolAppFromGroupResponse) SetBody(v *UninstallCoolAppFromGroupResponseBody) *UninstallCoolAppFromGroupResponse {
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
// 群酷应用排序
//
// @param request - InstallCoolAppOrderToGroupRequest
//
// @param headers - InstallCoolAppOrderToGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallCoolAppOrderToGroupResponse
func (client *Client) InstallCoolAppOrderToGroupWithOptions(request *InstallCoolAppOrderToGroupRequest, headers *InstallCoolAppOrderToGroupHeaders, runtime *util.RuntimeOptions) (_result *InstallCoolAppOrderToGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		body["conversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.SortedPluginIdList)) {
		body["sortedPluginIdList"] = request.SortedPluginIdList
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.UnsortedPluginIdList)) {
		body["unsortedPluginIdList"] = request.UnsortedPluginIdList
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
		Action:      tea.String("InstallCoolAppOrderToGroup"),
		Version:     tea.String("coolApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/coolApp/shortcuts/plugins/order"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallCoolAppOrderToGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 群酷应用排序
//
// @param request - InstallCoolAppOrderToGroupRequest
//
// @return InstallCoolAppOrderToGroupResponse
func (client *Client) InstallCoolAppOrderToGroup(request *InstallCoolAppOrderToGroupRequest) (_result *InstallCoolAppOrderToGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InstallCoolAppOrderToGroupHeaders{}
	_result = &InstallCoolAppOrderToGroupResponse{}
	_body, _err := client.InstallCoolAppOrderToGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 安装酷应用到群
//
// @param request - InstallCoolAppToGroupRequest
//
// @param headers - InstallCoolAppToGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallCoolAppToGroupResponse
func (client *Client) InstallCoolAppToGroupWithOptions(request *InstallCoolAppToGroupRequest, headers *InstallCoolAppToGroupHeaders, runtime *util.RuntimeOptions) (_result *InstallCoolAppToGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		body["conversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OperateCoolAppCode)) {
		body["operateCoolAppCode"] = request.OperateCoolAppCode
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
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
		Action:      tea.String("InstallCoolAppToGroup"),
		Version:     tea.String("coolApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/coolApp/shortcuts/plugins/install"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallCoolAppToGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 安装酷应用到群
//
// @param request - InstallCoolAppToGroupRequest
//
// @return InstallCoolAppToGroupResponse
func (client *Client) InstallCoolAppToGroup(request *InstallCoolAppToGroupRequest) (_result *InstallCoolAppToGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InstallCoolAppToGroupHeaders{}
	_result = &InstallCoolAppToGroupResponse{}
	_body, _err := client.InstallCoolAppToGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询群插件栏
//
// @param request - QueryCoolAppShortcutOrderRequest
//
// @param headers - QueryCoolAppShortcutOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCoolAppShortcutOrderResponse
func (client *Client) QueryCoolAppShortcutOrderWithOptions(request *QueryCoolAppShortcutOrderRequest, headers *QueryCoolAppShortcutOrderHeaders, runtime *util.RuntimeOptions) (_result *QueryCoolAppShortcutOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		body["conversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
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
		Action:      tea.String("QueryCoolAppShortcutOrder"),
		Version:     tea.String("coolApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/coolApp/shortcuts/plugins/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCoolAppShortcutOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询群插件栏
//
// @param request - QueryCoolAppShortcutOrderRequest
//
// @return QueryCoolAppShortcutOrderResponse
func (client *Client) QueryCoolAppShortcutOrder(request *QueryCoolAppShortcutOrderRequest) (_result *QueryCoolAppShortcutOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCoolAppShortcutOrderHeaders{}
	_result = &QueryCoolAppShortcutOrderResponse{}
	_body, _err := client.QueryCoolAppShortcutOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 从群内卸载酷应用
//
// @param request - UninstallCoolAppFromGroupRequest
//
// @param headers - UninstallCoolAppFromGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallCoolAppFromGroupResponse
func (client *Client) UninstallCoolAppFromGroupWithOptions(request *UninstallCoolAppFromGroupRequest, headers *UninstallCoolAppFromGroupHeaders, runtime *util.RuntimeOptions) (_result *UninstallCoolAppFromGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		body["conversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OperateCoolAppCode)) {
		body["operateCoolAppCode"] = request.OperateCoolAppCode
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
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
		Action:      tea.String("UninstallCoolAppFromGroup"),
		Version:     tea.String("coolApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/coolApp/shortcuts/plugins/uninstall"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UninstallCoolAppFromGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 从群内卸载酷应用
//
// @param request - UninstallCoolAppFromGroupRequest
//
// @return UninstallCoolAppFromGroupResponse
func (client *Client) UninstallCoolAppFromGroup(request *UninstallCoolAppFromGroupRequest) (_result *UninstallCoolAppFromGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UninstallCoolAppFromGroupHeaders{}
	_result = &UninstallCoolAppFromGroupResponse{}
	_body, _err := client.UninstallCoolAppFromGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
