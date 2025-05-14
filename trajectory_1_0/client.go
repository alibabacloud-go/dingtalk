// This file is auto-generated, don't edit it. Thanks.
package trajectory_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type QueryAppActiveUsersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAppActiveUsersHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAppActiveUsersHeaders) GoString() string {
	return s.String()
}

func (s *QueryAppActiveUsersHeaders) SetCommonHeaders(v map[string]*string) *QueryAppActiveUsersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAppActiveUsersHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAppActiveUsersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAppActiveUsersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 100
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	NeedPositionInfo *bool `json:"needPositionInfo,omitempty" xml:"needPositionInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryAppActiveUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAppActiveUsersRequest) GoString() string {
	return s.String()
}

func (s *QueryAppActiveUsersRequest) SetMaxResults(v int64) *QueryAppActiveUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryAppActiveUsersRequest) SetNeedPositionInfo(v bool) *QueryAppActiveUsersRequest {
	s.NeedPositionInfo = &v
	return s
}

func (s *QueryAppActiveUsersRequest) SetNextToken(v int64) *QueryAppActiveUsersRequest {
	s.NextToken = &v
	return s
}

type QueryAppActiveUsersResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// This parameter is required.
	List []*QueryAppActiveUsersResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// This parameter is required.
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 23153
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryAppActiveUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAppActiveUsersResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAppActiveUsersResponseBody) SetHasMore(v bool) *QueryAppActiveUsersResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryAppActiveUsersResponseBody) SetList(v []*QueryAppActiveUsersResponseBodyList) *QueryAppActiveUsersResponseBody {
	s.List = v
	return s
}

func (s *QueryAppActiveUsersResponseBody) SetNextToken(v int64) *QueryAppActiveUsersResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryAppActiveUsersResponseBody) SetTotalCount(v int64) *QueryAppActiveUsersResponseBody {
	s.TotalCount = &v
	return s
}

type QueryAppActiveUsersResponseBodyList struct {
	// This parameter is required.
	//
	// example:
	//
	// kxm9dhfs01jd98cuv
	AppTraceId *string `json:"appTraceId,omitempty" xml:"appTraceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123.123
	Latitude *float32 `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123.123
	Longitude *float32 `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1619341954123
	ReportTime *int64 `json:"reportTime,omitempty" xml:"reportTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1619341754123
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// I0912384771
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryAppActiveUsersResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryAppActiveUsersResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryAppActiveUsersResponseBodyList) SetAppTraceId(v string) *QueryAppActiveUsersResponseBodyList {
	s.AppTraceId = &v
	return s
}

func (s *QueryAppActiveUsersResponseBodyList) SetLatitude(v float32) *QueryAppActiveUsersResponseBodyList {
	s.Latitude = &v
	return s
}

func (s *QueryAppActiveUsersResponseBodyList) SetLongitude(v float32) *QueryAppActiveUsersResponseBodyList {
	s.Longitude = &v
	return s
}

func (s *QueryAppActiveUsersResponseBodyList) SetReportTime(v int64) *QueryAppActiveUsersResponseBodyList {
	s.ReportTime = &v
	return s
}

func (s *QueryAppActiveUsersResponseBodyList) SetStartTime(v int64) *QueryAppActiveUsersResponseBodyList {
	s.StartTime = &v
	return s
}

func (s *QueryAppActiveUsersResponseBodyList) SetUserId(v string) *QueryAppActiveUsersResponseBodyList {
	s.UserId = &v
	return s
}

type QueryAppActiveUsersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAppActiveUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAppActiveUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAppActiveUsersResponse) GoString() string {
	return s.String()
}

func (s *QueryAppActiveUsersResponse) SetHeaders(v map[string]*string) *QueryAppActiveUsersResponse {
	s.Headers = v
	return s
}

func (s *QueryAppActiveUsersResponse) SetStatusCode(v int32) *QueryAppActiveUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAppActiveUsersResponse) SetBody(v *QueryAppActiveUsersResponseBody) *QueryAppActiveUsersResponse {
	s.Body = v
	return s
}

type QueryCollectingTraceTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCollectingTraceTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCollectingTraceTaskHeaders) GoString() string {
	return s.String()
}

func (s *QueryCollectingTraceTaskHeaders) SetCommonHeaders(v map[string]*string) *QueryCollectingTraceTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCollectingTraceTaskHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCollectingTraceTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCollectingTraceTaskRequest struct {
	// This parameter is required.
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s QueryCollectingTraceTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCollectingTraceTaskRequest) GoString() string {
	return s.String()
}

func (s *QueryCollectingTraceTaskRequest) SetUserIds(v []*string) *QueryCollectingTraceTaskRequest {
	s.UserIds = v
	return s
}

type QueryCollectingTraceTaskResponseBody struct {
	// This parameter is required.
	List []*QueryCollectingTraceTaskResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s QueryCollectingTraceTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCollectingTraceTaskResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCollectingTraceTaskResponseBody) SetList(v []*QueryCollectingTraceTaskResponseBodyList) *QueryCollectingTraceTaskResponseBody {
	s.List = v
	return s
}

type QueryCollectingTraceTaskResponseBodyList struct {
	// This parameter is required.
	//
	// example:
	//
	// ffsfsdf
	AppTraceId *string `json:"appTraceId,omitempty" xml:"appTraceId,omitempty"`
	// This parameter is required.
	GeoCollectPeriod *int64 `json:"geoCollectPeriod,omitempty" xml:"geoCollectPeriod,omitempty"`
	// This parameter is required.
	GeoReportPeriod *int64 `json:"geoReportPeriod,omitempty" xml:"geoReportPeriod,omitempty"`
	// This parameter is required.
	GeoReportStatus *int64 `json:"geoReportStatus,omitempty" xml:"geoReportStatus,omitempty"`
	// This parameter is required.
	ReportEndTime *int64 `json:"reportEndTime,omitempty" xml:"reportEndTime,omitempty"`
	// This parameter is required.
	ReportStartTime *int64 `json:"reportStartTime,omitempty" xml:"reportStartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// I01231231ads1
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryCollectingTraceTaskResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryCollectingTraceTaskResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryCollectingTraceTaskResponseBodyList) SetAppTraceId(v string) *QueryCollectingTraceTaskResponseBodyList {
	s.AppTraceId = &v
	return s
}

func (s *QueryCollectingTraceTaskResponseBodyList) SetGeoCollectPeriod(v int64) *QueryCollectingTraceTaskResponseBodyList {
	s.GeoCollectPeriod = &v
	return s
}

func (s *QueryCollectingTraceTaskResponseBodyList) SetGeoReportPeriod(v int64) *QueryCollectingTraceTaskResponseBodyList {
	s.GeoReportPeriod = &v
	return s
}

func (s *QueryCollectingTraceTaskResponseBodyList) SetGeoReportStatus(v int64) *QueryCollectingTraceTaskResponseBodyList {
	s.GeoReportStatus = &v
	return s
}

func (s *QueryCollectingTraceTaskResponseBodyList) SetReportEndTime(v int64) *QueryCollectingTraceTaskResponseBodyList {
	s.ReportEndTime = &v
	return s
}

func (s *QueryCollectingTraceTaskResponseBodyList) SetReportStartTime(v int64) *QueryCollectingTraceTaskResponseBodyList {
	s.ReportStartTime = &v
	return s
}

func (s *QueryCollectingTraceTaskResponseBodyList) SetUserId(v string) *QueryCollectingTraceTaskResponseBodyList {
	s.UserId = &v
	return s
}

type QueryCollectingTraceTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCollectingTraceTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCollectingTraceTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCollectingTraceTaskResponse) GoString() string {
	return s.String()
}

func (s *QueryCollectingTraceTaskResponse) SetHeaders(v map[string]*string) *QueryCollectingTraceTaskResponse {
	s.Headers = v
	return s
}

func (s *QueryCollectingTraceTaskResponse) SetStatusCode(v int32) *QueryCollectingTraceTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCollectingTraceTaskResponse) SetBody(v *QueryCollectingTraceTaskResponseBody) *QueryCollectingTraceTaskResponse {
	s.Body = v
	return s
}

type QueryPageTraceDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryPageTraceDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryPageTraceDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryPageTraceDataHeaders) SetCommonHeaders(v map[string]*string) *QueryPageTraceDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryPageTraceDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryPageTraceDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryPageTraceDataRequest struct {
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 0
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	StaffId   *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
	StartTime *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s QueryPageTraceDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPageTraceDataRequest) GoString() string {
	return s.String()
}

func (s *QueryPageTraceDataRequest) SetEndTime(v int64) *QueryPageTraceDataRequest {
	s.EndTime = &v
	return s
}

func (s *QueryPageTraceDataRequest) SetMaxResults(v int64) *QueryPageTraceDataRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryPageTraceDataRequest) SetNextToken(v int64) *QueryPageTraceDataRequest {
	s.NextToken = &v
	return s
}

func (s *QueryPageTraceDataRequest) SetStaffId(v string) *QueryPageTraceDataRequest {
	s.StaffId = &v
	return s
}

func (s *QueryPageTraceDataRequest) SetStartTime(v int64) *QueryPageTraceDataRequest {
	s.StartTime = &v
	return s
}

func (s *QueryPageTraceDataRequest) SetTraceId(v string) *QueryPageTraceDataRequest {
	s.TraceId = &v
	return s
}

type QueryPageTraceDataResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// This parameter is required.
	List []*QueryPageTraceDataResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryPageTraceDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPageTraceDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPageTraceDataResponseBody) SetHasMore(v bool) *QueryPageTraceDataResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryPageTraceDataResponseBody) SetList(v []*QueryPageTraceDataResponseBodyList) *QueryPageTraceDataResponseBody {
	s.List = v
	return s
}

func (s *QueryPageTraceDataResponseBody) SetNextToken(v int64) *QueryPageTraceDataResponseBody {
	s.NextToken = &v
	return s
}

type QueryPageTraceDataResponseBodyList struct {
	// This parameter is required.
	Coordinates *QueryPageTraceDataResponseBodyListCoordinates `json:"coordinates,omitempty" xml:"coordinates,omitempty" type:"Struct"`
	// This parameter is required.
	GmtLocation *int64 `json:"gmtLocation,omitempty" xml:"gmtLocation,omitempty"`
	// This parameter is required.
	GmtUpload *int64 `json:"gmtUpload,omitempty" xml:"gmtUpload,omitempty"`
}

func (s QueryPageTraceDataResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryPageTraceDataResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryPageTraceDataResponseBodyList) SetCoordinates(v *QueryPageTraceDataResponseBodyListCoordinates) *QueryPageTraceDataResponseBodyList {
	s.Coordinates = v
	return s
}

func (s *QueryPageTraceDataResponseBodyList) SetGmtLocation(v int64) *QueryPageTraceDataResponseBodyList {
	s.GmtLocation = &v
	return s
}

func (s *QueryPageTraceDataResponseBodyList) SetGmtUpload(v int64) *QueryPageTraceDataResponseBodyList {
	s.GmtUpload = &v
	return s
}

type QueryPageTraceDataResponseBodyListCoordinates struct {
	// This parameter is required.
	Latitude *float32 `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// This parameter is required.
	Longitude *float32 `json:"longitude,omitempty" xml:"longitude,omitempty"`
}

func (s QueryPageTraceDataResponseBodyListCoordinates) String() string {
	return tea.Prettify(s)
}

func (s QueryPageTraceDataResponseBodyListCoordinates) GoString() string {
	return s.String()
}

func (s *QueryPageTraceDataResponseBodyListCoordinates) SetLatitude(v float32) *QueryPageTraceDataResponseBodyListCoordinates {
	s.Latitude = &v
	return s
}

func (s *QueryPageTraceDataResponseBodyListCoordinates) SetLongitude(v float32) *QueryPageTraceDataResponseBodyListCoordinates {
	s.Longitude = &v
	return s
}

type QueryPageTraceDataResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPageTraceDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPageTraceDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPageTraceDataResponse) GoString() string {
	return s.String()
}

func (s *QueryPageTraceDataResponse) SetHeaders(v map[string]*string) *QueryPageTraceDataResponse {
	s.Headers = v
	return s
}

func (s *QueryPageTraceDataResponse) SetStatusCode(v int32) *QueryPageTraceDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPageTraceDataResponse) SetBody(v *QueryPageTraceDataResponseBody) *QueryPageTraceDataResponse {
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
// 查询APP当前开启轨迹采集的用户
//
// @param request - QueryAppActiveUsersRequest
//
// @param headers - QueryAppActiveUsersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAppActiveUsersResponse
func (client *Client) QueryAppActiveUsersWithOptions(request *QueryAppActiveUsersRequest, headers *QueryAppActiveUsersHeaders, runtime *util.RuntimeOptions) (_result *QueryAppActiveUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NeedPositionInfo)) {
		query["needPositionInfo"] = request.NeedPositionInfo
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
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
		Action:      tea.String("QueryAppActiveUsers"),
		Version:     tea.String("trajectory_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/trajectory/activeUsers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAppActiveUsersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询APP当前开启轨迹采集的用户
//
// @param request - QueryAppActiveUsersRequest
//
// @return QueryAppActiveUsersResponse
func (client *Client) QueryAppActiveUsers(request *QueryAppActiveUsersRequest) (_result *QueryAppActiveUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAppActiveUsersHeaders{}
	_result = &QueryAppActiveUsersResponse{}
	_body, _err := client.QueryAppActiveUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询应用采集中的轨迹任务
//
// @param request - QueryCollectingTraceTaskRequest
//
// @param headers - QueryCollectingTraceTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCollectingTraceTaskResponse
func (client *Client) QueryCollectingTraceTaskWithOptions(request *QueryCollectingTraceTaskRequest, headers *QueryCollectingTraceTaskHeaders, runtime *util.RuntimeOptions) (_result *QueryCollectingTraceTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
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
		Action:      tea.String("QueryCollectingTraceTask"),
		Version:     tea.String("trajectory_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/trajectory/currentTasks/queryByUserIds"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCollectingTraceTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询应用采集中的轨迹任务
//
// @param request - QueryCollectingTraceTaskRequest
//
// @return QueryCollectingTraceTaskResponse
func (client *Client) QueryCollectingTraceTask(request *QueryCollectingTraceTaskRequest) (_result *QueryCollectingTraceTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCollectingTraceTaskHeaders{}
	_result = &QueryCollectingTraceTaskResponse{}
	_body, _err := client.QueryCollectingTraceTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询轨迹数据
//
// @param request - QueryPageTraceDataRequest
//
// @param headers - QueryPageTraceDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPageTraceDataResponse
func (client *Client) QueryPageTraceDataWithOptions(request *QueryPageTraceDataRequest, headers *QueryPageTraceDataHeaders, runtime *util.RuntimeOptions) (_result *QueryPageTraceDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.StaffId)) {
		query["staffId"] = request.StaffId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TraceId)) {
		query["traceId"] = request.TraceId
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
		Action:      tea.String("QueryPageTraceData"),
		Version:     tea.String("trajectory_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/trajectory/data"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPageTraceDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询轨迹数据
//
// @param request - QueryPageTraceDataRequest
//
// @return QueryPageTraceDataResponse
func (client *Client) QueryPageTraceData(request *QueryPageTraceDataRequest) (_result *QueryPageTraceDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryPageTraceDataHeaders{}
	_result = &QueryPageTraceDataResponse{}
	_body, _err := client.QueryPageTraceDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
