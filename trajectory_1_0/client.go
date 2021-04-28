// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package trajectory_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
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
	// 是否需要返回位置信息
	NeedPositionInfo *bool `json:"needPositionInfo,omitempty" xml:"needPositionInfo,omitempty"`
	// 标记当前开始读取的位置，置空表示从头开始
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 本次读取的最大数据记录数量
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
}

func (s QueryAppActiveUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAppActiveUsersRequest) GoString() string {
	return s.String()
}

func (s *QueryAppActiveUsersRequest) SetNeedPositionInfo(v bool) *QueryAppActiveUsersRequest {
	s.NeedPositionInfo = &v
	return s
}

func (s *QueryAppActiveUsersRequest) SetNextToken(v int64) *QueryAppActiveUsersRequest {
	s.NextToken = &v
	return s
}

func (s *QueryAppActiveUsersRequest) SetMaxResults(v int64) *QueryAppActiveUsersRequest {
	s.MaxResults = &v
	return s
}

type QueryAppActiveUsersResponseBody struct {
	// 是否存在更多数据需要获取
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 总数
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 数据集合
	List []*QueryAppActiveUsersResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// 下一次获取开始位置
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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

func (s *QueryAppActiveUsersResponseBody) SetTotalCount(v int64) *QueryAppActiveUsersResponseBody {
	s.TotalCount = &v
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

type QueryAppActiveUsersResponseBodyList struct {
	// 轨迹采集开启时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 应用轨迹ID
	AppTraceId *string `json:"appTraceId,omitempty" xml:"appTraceId,omitempty"`
	// 经度
	Longitude *float32 `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 纬度
	Latitude *float32 `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 该位置采集时间
	ReportTime *int64 `json:"reportTime,omitempty" xml:"reportTime,omitempty"`
	// 员工Id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryAppActiveUsersResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryAppActiveUsersResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryAppActiveUsersResponseBodyList) SetStartTime(v int64) *QueryAppActiveUsersResponseBodyList {
	s.StartTime = &v
	return s
}

func (s *QueryAppActiveUsersResponseBodyList) SetAppTraceId(v string) *QueryAppActiveUsersResponseBodyList {
	s.AppTraceId = &v
	return s
}

func (s *QueryAppActiveUsersResponseBodyList) SetLongitude(v float32) *QueryAppActiveUsersResponseBodyList {
	s.Longitude = &v
	return s
}

func (s *QueryAppActiveUsersResponseBodyList) SetLatitude(v float32) *QueryAppActiveUsersResponseBodyList {
	s.Latitude = &v
	return s
}

func (s *QueryAppActiveUsersResponseBodyList) SetReportTime(v int64) *QueryAppActiveUsersResponseBodyList {
	s.ReportTime = &v
	return s
}

func (s *QueryAppActiveUsersResponseBodyList) SetUserId(v string) *QueryAppActiveUsersResponseBodyList {
	s.UserId = &v
	return s
}

type QueryAppActiveUsersResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAppActiveUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 员工用户ID列表
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
	// isvOrgId
	DingIsvOrgId *int64 `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	// tokenGrantType
	DingTokenGrantType *int64 `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	// appKey
	DingClientId *string `json:"dingClientId,omitempty" xml:"dingClientId,omitempty"`
	// orgId
	DingOrgId *int64 `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	// oauthAppId
	DingOauthAppId *int64 `json:"dingOauthAppId,omitempty" xml:"dingOauthAppId,omitempty"`
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

func (s *QueryCollectingTraceTaskRequest) SetDingIsvOrgId(v int64) *QueryCollectingTraceTaskRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *QueryCollectingTraceTaskRequest) SetDingTokenGrantType(v int64) *QueryCollectingTraceTaskRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *QueryCollectingTraceTaskRequest) SetDingClientId(v string) *QueryCollectingTraceTaskRequest {
	s.DingClientId = &v
	return s
}

func (s *QueryCollectingTraceTaskRequest) SetDingOrgId(v int64) *QueryCollectingTraceTaskRequest {
	s.DingOrgId = &v
	return s
}

func (s *QueryCollectingTraceTaskRequest) SetDingOauthAppId(v int64) *QueryCollectingTraceTaskRequest {
	s.DingOauthAppId = &v
	return s
}

type QueryCollectingTraceTaskResponseBody struct {
	// result
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
	// 应用轨迹ID
	AppTraceId *string `json:"appTraceId,omitempty" xml:"appTraceId,omitempty"`
	// 组织下员工Id
	UserId           *string `json:"userId,omitempty" xml:"userId,omitempty"`
	GeoReportStatus  *int64  `json:"geoReportStatus,omitempty" xml:"geoReportStatus,omitempty"`
	GeoReportPeriod  *int64  `json:"geoReportPeriod,omitempty" xml:"geoReportPeriod,omitempty"`
	GeoCollectPeriod *int64  `json:"geoCollectPeriod,omitempty" xml:"geoCollectPeriod,omitempty"`
	ReportStartTime  *int64  `json:"reportStartTime,omitempty" xml:"reportStartTime,omitempty"`
	ReportEndTime    *int64  `json:"reportEndTime,omitempty" xml:"reportEndTime,omitempty"`
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

func (s *QueryCollectingTraceTaskResponseBodyList) SetUserId(v string) *QueryCollectingTraceTaskResponseBodyList {
	s.UserId = &v
	return s
}

func (s *QueryCollectingTraceTaskResponseBodyList) SetGeoReportStatus(v int64) *QueryCollectingTraceTaskResponseBodyList {
	s.GeoReportStatus = &v
	return s
}

func (s *QueryCollectingTraceTaskResponseBodyList) SetGeoReportPeriod(v int64) *QueryCollectingTraceTaskResponseBodyList {
	s.GeoReportPeriod = &v
	return s
}

func (s *QueryCollectingTraceTaskResponseBodyList) SetGeoCollectPeriod(v int64) *QueryCollectingTraceTaskResponseBodyList {
	s.GeoCollectPeriod = &v
	return s
}

func (s *QueryCollectingTraceTaskResponseBodyList) SetReportStartTime(v int64) *QueryCollectingTraceTaskResponseBodyList {
	s.ReportStartTime = &v
	return s
}

func (s *QueryCollectingTraceTaskResponseBodyList) SetReportEndTime(v int64) *QueryCollectingTraceTaskResponseBodyList {
	s.ReportEndTime = &v
	return s
}

type QueryCollectingTraceTaskResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCollectingTraceTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// traceId
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
	// 起始位置
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 查询数量
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 终止时间
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 员工ID
	StaffId *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
}

func (s QueryPageTraceDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPageTraceDataRequest) GoString() string {
	return s.String()
}

func (s *QueryPageTraceDataRequest) SetTraceId(v string) *QueryPageTraceDataRequest {
	s.TraceId = &v
	return s
}

func (s *QueryPageTraceDataRequest) SetNextToken(v int64) *QueryPageTraceDataRequest {
	s.NextToken = &v
	return s
}

func (s *QueryPageTraceDataRequest) SetMaxResults(v int64) *QueryPageTraceDataRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryPageTraceDataRequest) SetStartTime(v int64) *QueryPageTraceDataRequest {
	s.StartTime = &v
	return s
}

func (s *QueryPageTraceDataRequest) SetEndTime(v int64) *QueryPageTraceDataRequest {
	s.EndTime = &v
	return s
}

func (s *QueryPageTraceDataRequest) SetStaffId(v string) *QueryPageTraceDataRequest {
	s.StaffId = &v
	return s
}

type QueryPageTraceDataResponseBody struct {
	// 轨迹点列表
	List []*QueryPageTraceDataResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// 是否结束
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 下一个开始位置
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryPageTraceDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPageTraceDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPageTraceDataResponseBody) SetList(v []*QueryPageTraceDataResponseBodyList) *QueryPageTraceDataResponseBody {
	s.List = v
	return s
}

func (s *QueryPageTraceDataResponseBody) SetHasMore(v bool) *QueryPageTraceDataResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryPageTraceDataResponseBody) SetNextToken(v int64) *QueryPageTraceDataResponseBody {
	s.NextToken = &v
	return s
}

type QueryPageTraceDataResponseBodyList struct {
	// 经纬度
	Coordinates *QueryPageTraceDataResponseBodyListCoordinates `json:"coordinates,omitempty" xml:"coordinates,omitempty" type:"Struct"`
	// 定位时间
	GmtLocation *int64 `json:"gmtLocation,omitempty" xml:"gmtLocation,omitempty"`
	// 上报时间
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
	// 经度
	Longitude *float32 `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 纬度
	Latitude *float32 `json:"latitude,omitempty" xml:"latitude,omitempty"`
}

func (s QueryPageTraceDataResponseBodyListCoordinates) String() string {
	return tea.Prettify(s)
}

func (s QueryPageTraceDataResponseBodyListCoordinates) GoString() string {
	return s.String()
}

func (s *QueryPageTraceDataResponseBodyListCoordinates) SetLongitude(v float32) *QueryPageTraceDataResponseBodyListCoordinates {
	s.Longitude = &v
	return s
}

func (s *QueryPageTraceDataResponseBodyListCoordinates) SetLatitude(v float32) *QueryPageTraceDataResponseBodyListCoordinates {
	s.Latitude = &v
	return s
}

type QueryPageTraceDataResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryPageTraceDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

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

func (client *Client) QueryAppActiveUsersWithOptions(request *QueryAppActiveUsersRequest, headers *QueryAppActiveUsersHeaders, runtime *util.RuntimeOptions) (_result *QueryAppActiveUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedPositionInfo)) {
		query["needPositionInfo"] = request.NeedPositionInfo
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
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
	_result = &QueryAppActiveUsersResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAppActiveUsers"), tea.String("trajectory_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/trajectory/activeUsers"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) QueryCollectingTraceTaskWithOptions(request *QueryCollectingTraceTaskRequest, headers *QueryCollectingTraceTaskHeaders, runtime *util.RuntimeOptions) (_result *QueryCollectingTraceTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.DingClientId)) {
		body["dingClientId"] = request.DingClientId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOauthAppId)) {
		body["dingOauthAppId"] = request.DingOauthAppId
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
	_result = &QueryCollectingTraceTaskResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryCollectingTraceTask"), tea.String("trajectory_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/trajectory/currentTasks/queryByUserIds"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) QueryPageTraceDataWithOptions(request *QueryPageTraceDataRequest, headers *QueryPageTraceDataHeaders, runtime *util.RuntimeOptions) (_result *QueryPageTraceDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TraceId)) {
		query["traceId"] = request.TraceId
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StaffId)) {
		query["staffId"] = request.StaffId
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
	_result = &QueryPageTraceDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryPageTraceData"), tea.String("trajectory_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/trajectory/data"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
