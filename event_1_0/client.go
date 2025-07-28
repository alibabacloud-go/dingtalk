// This file is auto-generated, don't edit it. Thanks.
package event_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetCallBackFaileResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCallBackFaileResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCallBackFaileResultHeaders) GoString() string {
	return s.String()
}

func (s *GetCallBackFaileResultHeaders) SetCommonHeaders(v map[string]*string) *GetCallBackFaileResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCallBackFaileResultHeaders) SetXAcsDingtalkAccessToken(v string) *GetCallBackFaileResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCallBackFaileResultRequest struct {
	// example:
	//
	// 1606126433000
	BeginTime *int64 `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// example:
	//
	// 1606126493000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
}

func (s GetCallBackFaileResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCallBackFaileResultRequest) GoString() string {
	return s.String()
}

func (s *GetCallBackFaileResultRequest) SetBeginTime(v int64) *GetCallBackFaileResultRequest {
	s.BeginTime = &v
	return s
}

func (s *GetCallBackFaileResultRequest) SetEndTime(v int64) *GetCallBackFaileResultRequest {
	s.EndTime = &v
	return s
}

type GetCallBackFaileResultResponseBody struct {
	FailedList []*GetCallBackFaileResultResponseBodyFailedList `json:"failedList,omitempty" xml:"failedList,omitempty" type:"Repeated"`
	HasMore    *bool                                           `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
}

func (s GetCallBackFaileResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCallBackFaileResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetCallBackFaileResultResponseBody) SetFailedList(v []*GetCallBackFaileResultResponseBodyFailedList) *GetCallBackFaileResultResponseBody {
	s.FailedList = v
	return s
}

func (s *GetCallBackFaileResultResponseBody) SetHasMore(v bool) *GetCallBackFaileResultResponseBody {
	s.HasMore = &v
	return s
}

type GetCallBackFaileResultResponseBodyFailedList struct {
	// example:
	//
	// {\"CalendarEventUpdateTime\":1668735924619,\"CorpId\":\"ding9**cd16741\",\"ChangeType\":\"updated\",\"EventType\":\"calendar_event_change\",\"CalendarId\":\"NzE3MjU0NEB1c2V***5jb218MTQwMDE2\",\"EventTime\":1668735924640,\"LegacyCalendarEventId\":\"1C1BB56076***8A338\",\"BizId\":\"1668**4640\",\"CalendarEventId\":\"RVNUZllHK**elEydz09\",\"operator\":{\"type\":\"user\"},\"UnionIdList\":[\"QQa**mYiE\"]}
	CallBackData *string `json:"callBackData,omitempty" xml:"callBackData,omitempty"`
	// example:
	//
	// calendar_event_change
	CallBackTag *string `json:"callBackTag,omitempty" xml:"callBackTag,omitempty"`
	// example:
	//
	// ding9f50b15b*****41
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// 166***39184
	EventTime *int64 `json:"eventTime,omitempty" xml:"eventTime,omitempty"`
}

func (s GetCallBackFaileResultResponseBodyFailedList) String() string {
	return tea.Prettify(s)
}

func (s GetCallBackFaileResultResponseBodyFailedList) GoString() string {
	return s.String()
}

func (s *GetCallBackFaileResultResponseBodyFailedList) SetCallBackData(v string) *GetCallBackFaileResultResponseBodyFailedList {
	s.CallBackData = &v
	return s
}

func (s *GetCallBackFaileResultResponseBodyFailedList) SetCallBackTag(v string) *GetCallBackFaileResultResponseBodyFailedList {
	s.CallBackTag = &v
	return s
}

func (s *GetCallBackFaileResultResponseBodyFailedList) SetCorpId(v string) *GetCallBackFaileResultResponseBodyFailedList {
	s.CorpId = &v
	return s
}

func (s *GetCallBackFaileResultResponseBodyFailedList) SetEventTime(v int64) *GetCallBackFaileResultResponseBodyFailedList {
	s.EventTime = &v
	return s
}

type GetCallBackFaileResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCallBackFaileResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCallBackFaileResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCallBackFaileResultResponse) GoString() string {
	return s.String()
}

func (s *GetCallBackFaileResultResponse) SetHeaders(v map[string]*string) *GetCallBackFaileResultResponse {
	s.Headers = v
	return s
}

func (s *GetCallBackFaileResultResponse) SetStatusCode(v int32) *GetCallBackFaileResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCallBackFaileResultResponse) SetBody(v *GetCallBackFaileResultResponseBody) *GetCallBackFaileResultResponse {
	s.Body = v
	return s
}

type InstallAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InstallAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s InstallAppHeaders) GoString() string {
	return s.String()
}

func (s *InstallAppHeaders) SetCommonHeaders(v map[string]*string) *InstallAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InstallAppHeaders) SetXAcsDingtalkAccessToken(v string) *InstallAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InstallAppRequest struct {
	AppId      *int64  `json:"appId,omitempty" xml:"appId,omitempty"`
	CorpId     *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DingUserId *string `json:"dingUserId,omitempty" xml:"dingUserId,omitempty"`
	SuiteId    *int64  `json:"suiteId,omitempty" xml:"suiteId,omitempty"`
}

func (s InstallAppRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallAppRequest) GoString() string {
	return s.String()
}

func (s *InstallAppRequest) SetAppId(v int64) *InstallAppRequest {
	s.AppId = &v
	return s
}

func (s *InstallAppRequest) SetCorpId(v string) *InstallAppRequest {
	s.CorpId = &v
	return s
}

func (s *InstallAppRequest) SetDingUserId(v string) *InstallAppRequest {
	s.DingUserId = &v
	return s
}

func (s *InstallAppRequest) SetSuiteId(v int64) *InstallAppRequest {
	s.SuiteId = &v
	return s
}

type InstallAppResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s InstallAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallAppResponseBody) GoString() string {
	return s.String()
}

func (s *InstallAppResponseBody) SetResult(v bool) *InstallAppResponseBody {
	s.Result = &v
	return s
}

type InstallAppResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallAppResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallAppResponse) GoString() string {
	return s.String()
}

func (s *InstallAppResponse) SetHeaders(v map[string]*string) *InstallAppResponse {
	s.Headers = v
	return s
}

func (s *InstallAppResponse) SetStatusCode(v int32) *InstallAppResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallAppResponse) SetBody(v *InstallAppResponseBody) *InstallAppResponse {
	s.Body = v
	return s
}

type InstallCoolAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InstallCoolAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s InstallCoolAppHeaders) GoString() string {
	return s.String()
}

func (s *InstallCoolAppHeaders) SetCommonHeaders(v map[string]*string) *InstallCoolAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InstallCoolAppHeaders) SetXAcsDingtalkAccessToken(v string) *InstallCoolAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InstallCoolAppRequest struct {
	// This parameter is required.
	AppId *int64 `json:"appId,omitempty" xml:"appId,omitempty"`
	// This parameter is required.
	CoolAppCode *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	// This parameter is required.
	CorpId  *string                `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Feature map[string]interface{} `json:"feature,omitempty" xml:"feature,omitempty"`
	// This parameter is required.
	InstallUid *string `json:"installUid,omitempty" xml:"installUid,omitempty"`
	// This parameter is required.
	OpenConversationId *string                `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	Options            map[string]interface{} `json:"options,omitempty" xml:"options,omitempty"`
	// This parameter is required.
	SuiteId *string `json:"suiteId,omitempty" xml:"suiteId,omitempty"`
}

func (s InstallCoolAppRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallCoolAppRequest) GoString() string {
	return s.String()
}

func (s *InstallCoolAppRequest) SetAppId(v int64) *InstallCoolAppRequest {
	s.AppId = &v
	return s
}

func (s *InstallCoolAppRequest) SetCoolAppCode(v string) *InstallCoolAppRequest {
	s.CoolAppCode = &v
	return s
}

func (s *InstallCoolAppRequest) SetCorpId(v string) *InstallCoolAppRequest {
	s.CorpId = &v
	return s
}

func (s *InstallCoolAppRequest) SetFeature(v map[string]interface{}) *InstallCoolAppRequest {
	s.Feature = v
	return s
}

func (s *InstallCoolAppRequest) SetInstallUid(v string) *InstallCoolAppRequest {
	s.InstallUid = &v
	return s
}

func (s *InstallCoolAppRequest) SetOpenConversationId(v string) *InstallCoolAppRequest {
	s.OpenConversationId = &v
	return s
}

func (s *InstallCoolAppRequest) SetOptions(v map[string]interface{}) *InstallCoolAppRequest {
	s.Options = v
	return s
}

func (s *InstallCoolAppRequest) SetSuiteId(v string) *InstallCoolAppRequest {
	s.SuiteId = &v
	return s
}

type InstallCoolAppShrinkRequest struct {
	// This parameter is required.
	AppId *int64 `json:"appId,omitempty" xml:"appId,omitempty"`
	// This parameter is required.
	CoolAppCode *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	// This parameter is required.
	CorpId        *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	FeatureShrink *string `json:"feature,omitempty" xml:"feature,omitempty"`
	// This parameter is required.
	InstallUid *string `json:"installUid,omitempty" xml:"installUid,omitempty"`
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OptionsShrink      *string `json:"options,omitempty" xml:"options,omitempty"`
	// This parameter is required.
	SuiteId *string `json:"suiteId,omitempty" xml:"suiteId,omitempty"`
}

func (s InstallCoolAppShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallCoolAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *InstallCoolAppShrinkRequest) SetAppId(v int64) *InstallCoolAppShrinkRequest {
	s.AppId = &v
	return s
}

func (s *InstallCoolAppShrinkRequest) SetCoolAppCode(v string) *InstallCoolAppShrinkRequest {
	s.CoolAppCode = &v
	return s
}

func (s *InstallCoolAppShrinkRequest) SetCorpId(v string) *InstallCoolAppShrinkRequest {
	s.CorpId = &v
	return s
}

func (s *InstallCoolAppShrinkRequest) SetFeatureShrink(v string) *InstallCoolAppShrinkRequest {
	s.FeatureShrink = &v
	return s
}

func (s *InstallCoolAppShrinkRequest) SetInstallUid(v string) *InstallCoolAppShrinkRequest {
	s.InstallUid = &v
	return s
}

func (s *InstallCoolAppShrinkRequest) SetOpenConversationId(v string) *InstallCoolAppShrinkRequest {
	s.OpenConversationId = &v
	return s
}

func (s *InstallCoolAppShrinkRequest) SetOptionsShrink(v string) *InstallCoolAppShrinkRequest {
	s.OptionsShrink = &v
	return s
}

func (s *InstallCoolAppShrinkRequest) SetSuiteId(v string) *InstallCoolAppShrinkRequest {
	s.SuiteId = &v
	return s
}

type InstallCoolAppResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s InstallCoolAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallCoolAppResponseBody) GoString() string {
	return s.String()
}

func (s *InstallCoolAppResponseBody) SetResult(v string) *InstallCoolAppResponseBody {
	s.Result = &v
	return s
}

type InstallCoolAppResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallCoolAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallCoolAppResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallCoolAppResponse) GoString() string {
	return s.String()
}

func (s *InstallCoolAppResponse) SetHeaders(v map[string]*string) *InstallCoolAppResponse {
	s.Headers = v
	return s
}

func (s *InstallCoolAppResponse) SetStatusCode(v int32) *InstallCoolAppResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallCoolAppResponse) SetBody(v *InstallCoolAppResponseBody) *InstallCoolAppResponse {
	s.Body = v
	return s
}

type RePushSuiteTicketRequest struct {
	SuiteId     *int64  `json:"suiteId,omitempty" xml:"suiteId,omitempty"`
	SuiteSecret *string `json:"suiteSecret,omitempty" xml:"suiteSecret,omitempty"`
}

func (s RePushSuiteTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s RePushSuiteTicketRequest) GoString() string {
	return s.String()
}

func (s *RePushSuiteTicketRequest) SetSuiteId(v int64) *RePushSuiteTicketRequest {
	s.SuiteId = &v
	return s
}

func (s *RePushSuiteTicketRequest) SetSuiteSecret(v string) *RePushSuiteTicketRequest {
	s.SuiteSecret = &v
	return s
}

type RePushSuiteTicketResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RePushSuiteTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RePushSuiteTicketResponseBody) GoString() string {
	return s.String()
}

func (s *RePushSuiteTicketResponseBody) SetResult(v string) *RePushSuiteTicketResponseBody {
	s.Result = &v
	return s
}

type RePushSuiteTicketResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RePushSuiteTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RePushSuiteTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s RePushSuiteTicketResponse) GoString() string {
	return s.String()
}

func (s *RePushSuiteTicketResponse) SetHeaders(v map[string]*string) *RePushSuiteTicketResponse {
	s.Headers = v
	return s
}

func (s *RePushSuiteTicketResponse) SetStatusCode(v int32) *RePushSuiteTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *RePushSuiteTicketResponse) SetBody(v *RePushSuiteTicketResponseBody) *RePushSuiteTicketResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

// Summary:
//
// 调用本获取推送失败的变更事件。
//
// @param request - GetCallBackFaileResultRequest
//
// @param headers - GetCallBackFaileResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCallBackFaileResultResponse
func (client *Client) GetCallBackFaileResultWithOptions(request *GetCallBackFaileResultRequest, headers *GetCallBackFaileResultHeaders, runtime *util.RuntimeOptions) (_result *GetCallBackFaileResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["beginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
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
		Action:      tea.String("GetCallBackFaileResult"),
		Version:     tea.String("event_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/event/callbacks/failedResults"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCallBackFaileResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用本获取推送失败的变更事件。
//
// @param request - GetCallBackFaileResultRequest
//
// @return GetCallBackFaileResultResponse
func (client *Client) GetCallBackFaileResult(request *GetCallBackFaileResultRequest) (_result *GetCallBackFaileResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCallBackFaileResultHeaders{}
	_result = &GetCallBackFaileResultResponse{}
	_body, _err := client.GetCallBackFaileResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 安装一方应用
//
// @param request - InstallAppRequest
//
// @param headers - InstallAppHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallAppResponse
func (client *Client) InstallAppWithOptions(request *InstallAppRequest, headers *InstallAppHeaders, runtime *util.RuntimeOptions) (_result *InstallAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DingUserId)) {
		query["dingUserId"] = request.DingUserId
	}

	if !tea.BoolValue(util.IsUnset(request.SuiteId)) {
		query["suiteId"] = request.SuiteId
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
		Action:      tea.String("InstallApp"),
		Version:     tea.String("event_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/event/elm/apps/install"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallAppResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 安装一方应用
//
// @param request - InstallAppRequest
//
// @return InstallAppResponse
func (client *Client) InstallApp(request *InstallAppRequest) (_result *InstallAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InstallAppHeaders{}
	_result = &InstallAppResponse{}
	_body, _err := client.InstallAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 安装酷应用
//
// @param tmpReq - InstallCoolAppRequest
//
// @param headers - InstallCoolAppHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallCoolAppResponse
func (client *Client) InstallCoolAppWithOptions(tmpReq *InstallCoolAppRequest, headers *InstallCoolAppHeaders, runtime *util.RuntimeOptions) (_result *InstallCoolAppResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &InstallCoolAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Feature)) {
		request.FeatureShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Feature, tea.String("feature"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Options)) {
		request.OptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Options, tea.String("options"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CoolAppCode)) {
		query["coolAppCode"] = request.CoolAppCode
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureShrink)) {
		query["feature"] = request.FeatureShrink
	}

	if !tea.BoolValue(util.IsUnset(request.InstallUid)) {
		query["installUid"] = request.InstallUid
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		query["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OptionsShrink)) {
		query["options"] = request.OptionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SuiteId)) {
		query["suiteId"] = request.SuiteId
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
		Action:      tea.String("InstallCoolApp"),
		Version:     tea.String("event_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/event/elm/coolApps/install"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallCoolAppResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 安装酷应用
//
// @param request - InstallCoolAppRequest
//
// @return InstallCoolAppResponse
func (client *Client) InstallCoolApp(request *InstallCoolAppRequest) (_result *InstallCoolAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InstallCoolAppHeaders{}
	_result = &InstallCoolAppResponse{}
	_body, _err := client.InstallCoolAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重新获取suiteTicket
//
// @param request - RePushSuiteTicketRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RePushSuiteTicketResponse
func (client *Client) RePushSuiteTicketWithOptions(request *RePushSuiteTicketRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RePushSuiteTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SuiteId)) {
		query["suiteId"] = request.SuiteId
	}

	if !tea.BoolValue(util.IsUnset(request.SuiteSecret)) {
		query["suiteSecret"] = request.SuiteSecret
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RePushSuiteTicket"),
		Version:     tea.String("event_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/event/suiteTicket/rePush"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RePushSuiteTicketResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重新获取suiteTicket
//
// @param request - RePushSuiteTicketRequest
//
// @return RePushSuiteTicketResponse
func (client *Client) RePushSuiteTicket(request *RePushSuiteTicketRequest) (_result *RePushSuiteTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RePushSuiteTicketResponse{}
	_body, _err := client.RePushSuiteTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
