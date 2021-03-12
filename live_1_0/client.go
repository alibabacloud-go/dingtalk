// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package live_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type StartCloudFeedHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s StartCloudFeedHeaders) String() string {
	return tea.Prettify(s)
}

func (s StartCloudFeedHeaders) GoString() string {
	return s.String()
}

func (s *StartCloudFeedHeaders) SetCommonHeaders(v map[string]*string) *StartCloudFeedHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartCloudFeedHeaders) SetXAcsDingtalkAccessToken(v string) *StartCloudFeedHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type StartCloudFeedRequest struct {
	// 操作者的组织内id（staffId）
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s StartCloudFeedRequest) String() string {
	return tea.Prettify(s)
}

func (s StartCloudFeedRequest) GoString() string {
	return s.String()
}

func (s *StartCloudFeedRequest) SetUserId(v string) *StartCloudFeedRequest {
	s.UserId = &v
	return s
}

type StartCloudFeedResponseBody struct {
	// 状态更改是否成功
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s StartCloudFeedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartCloudFeedResponseBody) GoString() string {
	return s.String()
}

func (s *StartCloudFeedResponseBody) SetResult(v bool) *StartCloudFeedResponseBody {
	s.Result = &v
	return s
}

type StartCloudFeedResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartCloudFeedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartCloudFeedResponse) String() string {
	return tea.Prettify(s)
}

func (s StartCloudFeedResponse) GoString() string {
	return s.String()
}

func (s *StartCloudFeedResponse) SetHeaders(v map[string]*string) *StartCloudFeedResponse {
	s.Headers = v
	return s
}

func (s *StartCloudFeedResponse) SetBody(v *StartCloudFeedResponseBody) *StartCloudFeedResponse {
	s.Body = v
	return s
}

type StopCloudFeedHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s StopCloudFeedHeaders) String() string {
	return tea.Prettify(s)
}

func (s StopCloudFeedHeaders) GoString() string {
	return s.String()
}

func (s *StopCloudFeedHeaders) SetCommonHeaders(v map[string]*string) *StopCloudFeedHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StopCloudFeedHeaders) SetXAcsDingtalkAccessToken(v string) *StopCloudFeedHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type StopCloudFeedRequest struct {
	// 操作者的组织内id（staffId）
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s StopCloudFeedRequest) String() string {
	return tea.Prettify(s)
}

func (s StopCloudFeedRequest) GoString() string {
	return s.String()
}

func (s *StopCloudFeedRequest) SetUserId(v string) *StopCloudFeedRequest {
	s.UserId = &v
	return s
}

type StopCloudFeedResponseBody struct {
	// 状态更改是否成功
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s StopCloudFeedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopCloudFeedResponseBody) GoString() string {
	return s.String()
}

func (s *StopCloudFeedResponseBody) SetResult(v bool) *StopCloudFeedResponseBody {
	s.Result = &v
	return s
}

type StopCloudFeedResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopCloudFeedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopCloudFeedResponse) String() string {
	return tea.Prettify(s)
}

func (s StopCloudFeedResponse) GoString() string {
	return s.String()
}

func (s *StopCloudFeedResponse) SetHeaders(v map[string]*string) *StopCloudFeedResponse {
	s.Headers = v
	return s
}

func (s *StopCloudFeedResponse) SetBody(v *StopCloudFeedResponseBody) *StopCloudFeedResponse {
	s.Body = v
	return s
}

type CreateCloudFeedHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateCloudFeedHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudFeedHeaders) GoString() string {
	return s.String()
}

func (s *CreateCloudFeedHeaders) SetCommonHeaders(v map[string]*string) *CreateCloudFeedHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateCloudFeedHeaders) SetXAcsDingtalkAccessToken(v string) *CreateCloudFeedHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateCloudFeedRequest struct {
	// 课程标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 课程简介
	Intro *string `json:"intro,omitempty" xml:"intro,omitempty"`
	// 创建课程的主播id（staffId）
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 预计开始的时间戳(未来的时间点)
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 课程封面Url
	CoverUrl *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	// 云导播课程资源的url
	VideoUrl *string `json:"videoUrl,omitempty" xml:"videoUrl,omitempty"`
}

func (s CreateCloudFeedRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudFeedRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudFeedRequest) SetTitle(v string) *CreateCloudFeedRequest {
	s.Title = &v
	return s
}

func (s *CreateCloudFeedRequest) SetIntro(v string) *CreateCloudFeedRequest {
	s.Intro = &v
	return s
}

func (s *CreateCloudFeedRequest) SetUserId(v string) *CreateCloudFeedRequest {
	s.UserId = &v
	return s
}

func (s *CreateCloudFeedRequest) SetStartTime(v int64) *CreateCloudFeedRequest {
	s.StartTime = &v
	return s
}

func (s *CreateCloudFeedRequest) SetCoverUrl(v string) *CreateCloudFeedRequest {
	s.CoverUrl = &v
	return s
}

func (s *CreateCloudFeedRequest) SetVideoUrl(v string) *CreateCloudFeedRequest {
	s.VideoUrl = &v
	return s
}

type CreateCloudFeedResponseBody struct {
	// 8c0ed3c3-e125-4a9d-aa40-18ad999398d4
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateCloudFeedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudFeedResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudFeedResponseBody) SetResult(v string) *CreateCloudFeedResponseBody {
	s.Result = &v
	return s
}

type CreateCloudFeedResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCloudFeedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCloudFeedResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudFeedResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudFeedResponse) SetHeaders(v map[string]*string) *CreateCloudFeedResponse {
	s.Headers = v
	return s
}

func (s *CreateCloudFeedResponse) SetBody(v *CreateCloudFeedResponseBody) *CreateCloudFeedResponse {
	s.Body = v
	return s
}

type AddShareCidListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddShareCidListHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddShareCidListHeaders) GoString() string {
	return s.String()
}

func (s *AddShareCidListHeaders) SetCommonHeaders(v map[string]*string) *AddShareCidListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddShareCidListHeaders) SetXAcsDingtalkAccessToken(v string) *AddShareCidListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddShareCidListRequest struct {
	// 操作的的组织内id(staffId)
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 添加的联播群列表
	GroupIds []*string `json:"groupIds,omitempty" xml:"groupIds,omitempty" type:"Repeated"`
	// 传入的群id类型（1 chatId / 2 openConversationId ）
	GroupIdType *int64 `json:"groupIdType,omitempty" xml:"groupIdType,omitempty"`
}

func (s AddShareCidListRequest) String() string {
	return tea.Prettify(s)
}

func (s AddShareCidListRequest) GoString() string {
	return s.String()
}

func (s *AddShareCidListRequest) SetUserId(v string) *AddShareCidListRequest {
	s.UserId = &v
	return s
}

func (s *AddShareCidListRequest) SetGroupIds(v []*string) *AddShareCidListRequest {
	s.GroupIds = v
	return s
}

func (s *AddShareCidListRequest) SetGroupIdType(v int64) *AddShareCidListRequest {
	s.GroupIdType = &v
	return s
}

type AddShareCidListResponseBody struct {
	// 是否联播成功
	HasShareSuccess *bool `json:"hasShareSuccess,omitempty" xml:"hasShareSuccess,omitempty"`
	// 本次请求成功联播的群列表
	ShareSuccessGroupList []*string `json:"shareSuccessGroupList,omitempty" xml:"shareSuccessGroupList,omitempty" type:"Repeated"`
}

func (s AddShareCidListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddShareCidListResponseBody) GoString() string {
	return s.String()
}

func (s *AddShareCidListResponseBody) SetHasShareSuccess(v bool) *AddShareCidListResponseBody {
	s.HasShareSuccess = &v
	return s
}

func (s *AddShareCidListResponseBody) SetShareSuccessGroupList(v []*string) *AddShareCidListResponseBody {
	s.ShareSuccessGroupList = v
	return s
}

type AddShareCidListResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddShareCidListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddShareCidListResponse) String() string {
	return tea.Prettify(s)
}

func (s AddShareCidListResponse) GoString() string {
	return s.String()
}

func (s *AddShareCidListResponse) SetHeaders(v map[string]*string) *AddShareCidListResponse {
	s.Headers = v
	return s
}

func (s *AddShareCidListResponse) SetBody(v *AddShareCidListResponseBody) *AddShareCidListResponse {
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

func (client *Client) StartCloudFeed(feedId *string, request *StartCloudFeedRequest) (_result *StartCloudFeedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StartCloudFeedHeaders{}
	_result = &StartCloudFeedResponse{}
	_body, _err := client.StartCloudFeedWithOptions(feedId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartCloudFeedWithOptions(feedId *string, request *StartCloudFeedRequest, headers *StartCloudFeedHeaders, runtime *util.RuntimeOptions) (_result *StartCloudFeedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
	_result = &StartCloudFeedResponse{}
	_body, _err := client.DoROARequest(tea.String("StartCloudFeed"), tea.String("live_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/live/cloudFeeds/"+tea.StringValue(feedId)+"/start"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopCloudFeed(feedId *string, request *StopCloudFeedRequest) (_result *StopCloudFeedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StopCloudFeedHeaders{}
	_result = &StopCloudFeedResponse{}
	_body, _err := client.StopCloudFeedWithOptions(feedId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopCloudFeedWithOptions(feedId *string, request *StopCloudFeedRequest, headers *StopCloudFeedHeaders, runtime *util.RuntimeOptions) (_result *StopCloudFeedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
	_result = &StopCloudFeedResponse{}
	_body, _err := client.DoROARequest(tea.String("StopCloudFeed"), tea.String("live_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/live/cloudFeeds/"+tea.StringValue(feedId)+"/stop"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCloudFeed(request *CreateCloudFeedRequest) (_result *CreateCloudFeedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateCloudFeedHeaders{}
	_result = &CreateCloudFeedResponse{}
	_body, _err := client.CreateCloudFeedWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCloudFeedWithOptions(request *CreateCloudFeedRequest, headers *CreateCloudFeedHeaders, runtime *util.RuntimeOptions) (_result *CreateCloudFeedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Intro)) {
		body["intro"] = request.Intro
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.CoverUrl)) {
		body["coverUrl"] = request.CoverUrl
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUrl)) {
		body["videoUrl"] = request.VideoUrl
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
	_result = &CreateCloudFeedResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateCloudFeed"), tea.String("live_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/live/cloudFeeds"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddShareCidList(feedId *string, request *AddShareCidListRequest) (_result *AddShareCidListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddShareCidListHeaders{}
	_result = &AddShareCidListResponse{}
	_body, _err := client.AddShareCidListWithOptions(feedId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddShareCidListWithOptions(feedId *string, request *AddShareCidListRequest, headers *AddShareCidListHeaders, runtime *util.RuntimeOptions) (_result *AddShareCidListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupIds)) {
		body["groupIds"] = request.GroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.GroupIdType)) {
		body["groupIdType"] = request.GroupIdType
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
	_result = &AddShareCidListResponse{}
	_body, _err := client.DoROARequest(tea.String("AddShareCidList"), tea.String("live_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/live/cloudFeeds/"+tea.StringValue(feedId)+"/share"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
