// This file is auto-generated, don't edit it. Thanks.
package minutes_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BatchGetMinutesDetailsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchGetMinutesDetailsHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchGetMinutesDetailsHeaders) GoString() string {
	return s.String()
}

func (s *BatchGetMinutesDetailsHeaders) SetCommonHeaders(v map[string]*string) *BatchGetMinutesDetailsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchGetMinutesDetailsHeaders) SetXAcsDingtalkAccessToken(v string) *BatchGetMinutesDetailsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchGetMinutesDetailsRequest struct {
	// This parameter is required.
	TaskUuids []*string `json:"taskUuids,omitempty" xml:"taskUuids,omitempty" type:"Repeated"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s BatchGetMinutesDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetMinutesDetailsRequest) GoString() string {
	return s.String()
}

func (s *BatchGetMinutesDetailsRequest) SetTaskUuids(v []*string) *BatchGetMinutesDetailsRequest {
	s.TaskUuids = v
	return s
}

func (s *BatchGetMinutesDetailsRequest) SetUnionId(v string) *BatchGetMinutesDetailsRequest {
	s.UnionId = &v
	return s
}

type BatchGetMinutesDetailsResponseBody struct {
	MinutesDetails []*BatchGetMinutesDetailsResponseBodyMinutesDetails `json:"minutesDetails,omitempty" xml:"minutesDetails,omitempty" type:"Repeated"`
}

func (s BatchGetMinutesDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchGetMinutesDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetMinutesDetailsResponseBody) SetMinutesDetails(v []*BatchGetMinutesDetailsResponseBodyMinutesDetails) *BatchGetMinutesDetailsResponseBody {
	s.MinutesDetails = v
	return s
}

type BatchGetMinutesDetailsResponseBodyMinutesDetails struct {
	BizType        *int32  `json:"bizType,omitempty" xml:"bizType,omitempty"`
	CreatorNick    *string `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	CreatorUnionId *string `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	DurationMicros *int64  `json:"durationMicros,omitempty" xml:"durationMicros,omitempty"`
	IsDeleted      *int32  `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	Size           *int64  `json:"size,omitempty" xml:"size,omitempty"`
	StartTime      *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	TaskUuid       *string `json:"taskUuid,omitempty" xml:"taskUuid,omitempty"`
	Title          *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s BatchGetMinutesDetailsResponseBodyMinutesDetails) String() string {
	return tea.Prettify(s)
}

func (s BatchGetMinutesDetailsResponseBodyMinutesDetails) GoString() string {
	return s.String()
}

func (s *BatchGetMinutesDetailsResponseBodyMinutesDetails) SetBizType(v int32) *BatchGetMinutesDetailsResponseBodyMinutesDetails {
	s.BizType = &v
	return s
}

func (s *BatchGetMinutesDetailsResponseBodyMinutesDetails) SetCreatorNick(v string) *BatchGetMinutesDetailsResponseBodyMinutesDetails {
	s.CreatorNick = &v
	return s
}

func (s *BatchGetMinutesDetailsResponseBodyMinutesDetails) SetCreatorUnionId(v string) *BatchGetMinutesDetailsResponseBodyMinutesDetails {
	s.CreatorUnionId = &v
	return s
}

func (s *BatchGetMinutesDetailsResponseBodyMinutesDetails) SetDurationMicros(v int64) *BatchGetMinutesDetailsResponseBodyMinutesDetails {
	s.DurationMicros = &v
	return s
}

func (s *BatchGetMinutesDetailsResponseBodyMinutesDetails) SetIsDeleted(v int32) *BatchGetMinutesDetailsResponseBodyMinutesDetails {
	s.IsDeleted = &v
	return s
}

func (s *BatchGetMinutesDetailsResponseBodyMinutesDetails) SetSize(v int64) *BatchGetMinutesDetailsResponseBodyMinutesDetails {
	s.Size = &v
	return s
}

func (s *BatchGetMinutesDetailsResponseBodyMinutesDetails) SetStartTime(v int64) *BatchGetMinutesDetailsResponseBodyMinutesDetails {
	s.StartTime = &v
	return s
}

func (s *BatchGetMinutesDetailsResponseBodyMinutesDetails) SetStatus(v int32) *BatchGetMinutesDetailsResponseBodyMinutesDetails {
	s.Status = &v
	return s
}

func (s *BatchGetMinutesDetailsResponseBodyMinutesDetails) SetTaskUuid(v string) *BatchGetMinutesDetailsResponseBodyMinutesDetails {
	s.TaskUuid = &v
	return s
}

func (s *BatchGetMinutesDetailsResponseBodyMinutesDetails) SetTitle(v string) *BatchGetMinutesDetailsResponseBodyMinutesDetails {
	s.Title = &v
	return s
}

type BatchGetMinutesDetailsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetMinutesDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetMinutesDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchGetMinutesDetailsResponse) GoString() string {
	return s.String()
}

func (s *BatchGetMinutesDetailsResponse) SetHeaders(v map[string]*string) *BatchGetMinutesDetailsResponse {
	s.Headers = v
	return s
}

func (s *BatchGetMinutesDetailsResponse) SetStatusCode(v int32) *BatchGetMinutesDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetMinutesDetailsResponse) SetBody(v *BatchGetMinutesDetailsResponseBody) *BatchGetMinutesDetailsResponse {
	s.Body = v
	return s
}

type DeleteMinutesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteMinutesHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteMinutesHeaders) GoString() string {
	return s.String()
}

func (s *DeleteMinutesHeaders) SetCommonHeaders(v map[string]*string) *DeleteMinutesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteMinutesHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteMinutesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteMinutesRequest struct {
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DeleteMinutesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMinutesRequest) GoString() string {
	return s.String()
}

func (s *DeleteMinutesRequest) SetUnionId(v string) *DeleteMinutesRequest {
	s.UnionId = &v
	return s
}

type DeleteMinutesResponseBody struct {
	TaskUuid *string `json:"taskUuid,omitempty" xml:"taskUuid,omitempty"`
}

func (s DeleteMinutesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMinutesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMinutesResponseBody) SetTaskUuid(v string) *DeleteMinutesResponseBody {
	s.TaskUuid = &v
	return s
}

type DeleteMinutesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMinutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMinutesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMinutesResponse) GoString() string {
	return s.String()
}

func (s *DeleteMinutesResponse) SetHeaders(v map[string]*string) *DeleteMinutesResponse {
	s.Headers = v
	return s
}

func (s *DeleteMinutesResponse) SetStatusCode(v int32) *DeleteMinutesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMinutesResponse) SetBody(v *DeleteMinutesResponseBody) *DeleteMinutesResponse {
	s.Body = v
	return s
}

type QueryCreateMinutesListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCreateMinutesListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCreateMinutesListHeaders) GoString() string {
	return s.String()
}

func (s *QueryCreateMinutesListHeaders) SetCommonHeaders(v map[string]*string) *QueryCreateMinutesListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCreateMinutesListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCreateMinutesListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCreateMinutesListRequest struct {
	// This parameter is required.
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryCreateMinutesListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCreateMinutesListRequest) GoString() string {
	return s.String()
}

func (s *QueryCreateMinutesListRequest) SetMaxResults(v int32) *QueryCreateMinutesListRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryCreateMinutesListRequest) SetNextToken(v string) *QueryCreateMinutesListRequest {
	s.NextToken = &v
	return s
}

func (s *QueryCreateMinutesListRequest) SetUnionId(v string) *QueryCreateMinutesListRequest {
	s.UnionId = &v
	return s
}

type QueryCreateMinutesListResponseBody struct {
	HasNext        *bool                                               `json:"hasNext,omitempty" xml:"hasNext,omitempty"`
	MinutesDetails []*QueryCreateMinutesListResponseBodyMinutesDetails `json:"minutesDetails,omitempty" xml:"minutesDetails,omitempty" type:"Repeated"`
	NextToken      *string                                             `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryCreateMinutesListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCreateMinutesListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCreateMinutesListResponseBody) SetHasNext(v bool) *QueryCreateMinutesListResponseBody {
	s.HasNext = &v
	return s
}

func (s *QueryCreateMinutesListResponseBody) SetMinutesDetails(v []*QueryCreateMinutesListResponseBodyMinutesDetails) *QueryCreateMinutesListResponseBody {
	s.MinutesDetails = v
	return s
}

func (s *QueryCreateMinutesListResponseBody) SetNextToken(v string) *QueryCreateMinutesListResponseBody {
	s.NextToken = &v
	return s
}

type QueryCreateMinutesListResponseBodyMinutesDetails struct {
	BizType        *int32  `json:"bizType,omitempty" xml:"bizType,omitempty"`
	CreatorNick    *string `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	CreatorUnionId *string `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	DurationMicros *int64  `json:"durationMicros,omitempty" xml:"durationMicros,omitempty"`
	IsDeleted      *int32  `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	Size           *int64  `json:"size,omitempty" xml:"size,omitempty"`
	StartTime      *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	TaskUuid       *string `json:"taskUuid,omitempty" xml:"taskUuid,omitempty"`
	Title          *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryCreateMinutesListResponseBodyMinutesDetails) String() string {
	return tea.Prettify(s)
}

func (s QueryCreateMinutesListResponseBodyMinutesDetails) GoString() string {
	return s.String()
}

func (s *QueryCreateMinutesListResponseBodyMinutesDetails) SetBizType(v int32) *QueryCreateMinutesListResponseBodyMinutesDetails {
	s.BizType = &v
	return s
}

func (s *QueryCreateMinutesListResponseBodyMinutesDetails) SetCreatorNick(v string) *QueryCreateMinutesListResponseBodyMinutesDetails {
	s.CreatorNick = &v
	return s
}

func (s *QueryCreateMinutesListResponseBodyMinutesDetails) SetCreatorUnionId(v string) *QueryCreateMinutesListResponseBodyMinutesDetails {
	s.CreatorUnionId = &v
	return s
}

func (s *QueryCreateMinutesListResponseBodyMinutesDetails) SetDurationMicros(v int64) *QueryCreateMinutesListResponseBodyMinutesDetails {
	s.DurationMicros = &v
	return s
}

func (s *QueryCreateMinutesListResponseBodyMinutesDetails) SetIsDeleted(v int32) *QueryCreateMinutesListResponseBodyMinutesDetails {
	s.IsDeleted = &v
	return s
}

func (s *QueryCreateMinutesListResponseBodyMinutesDetails) SetSize(v int64) *QueryCreateMinutesListResponseBodyMinutesDetails {
	s.Size = &v
	return s
}

func (s *QueryCreateMinutesListResponseBodyMinutesDetails) SetStartTime(v int64) *QueryCreateMinutesListResponseBodyMinutesDetails {
	s.StartTime = &v
	return s
}

func (s *QueryCreateMinutesListResponseBodyMinutesDetails) SetStatus(v int32) *QueryCreateMinutesListResponseBodyMinutesDetails {
	s.Status = &v
	return s
}

func (s *QueryCreateMinutesListResponseBodyMinutesDetails) SetTaskUuid(v string) *QueryCreateMinutesListResponseBodyMinutesDetails {
	s.TaskUuid = &v
	return s
}

func (s *QueryCreateMinutesListResponseBodyMinutesDetails) SetTitle(v string) *QueryCreateMinutesListResponseBodyMinutesDetails {
	s.Title = &v
	return s
}

type QueryCreateMinutesListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCreateMinutesListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCreateMinutesListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCreateMinutesListResponse) GoString() string {
	return s.String()
}

func (s *QueryCreateMinutesListResponse) SetHeaders(v map[string]*string) *QueryCreateMinutesListResponse {
	s.Headers = v
	return s
}

func (s *QueryCreateMinutesListResponse) SetStatusCode(v int32) *QueryCreateMinutesListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCreateMinutesListResponse) SetBody(v *QueryCreateMinutesListResponseBody) *QueryCreateMinutesListResponse {
	s.Body = v
	return s
}

type QueryMinutesPlayInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMinutesPlayInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesPlayInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryMinutesPlayInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryMinutesPlayInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMinutesPlayInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMinutesPlayInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMinutesPlayInfoRequest struct {
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryMinutesPlayInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesPlayInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryMinutesPlayInfoRequest) SetUnionId(v string) *QueryMinutesPlayInfoRequest {
	s.UnionId = &v
	return s
}

type QueryMinutesPlayInfoResponseBody struct {
	PlayInfo *QueryMinutesPlayInfoResponseBodyPlayInfo `json:"playInfo,omitempty" xml:"playInfo,omitempty" type:"Struct"`
}

func (s QueryMinutesPlayInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesPlayInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMinutesPlayInfoResponseBody) SetPlayInfo(v *QueryMinutesPlayInfoResponseBodyPlayInfo) *QueryMinutesPlayInfoResponseBody {
	s.PlayInfo = v
	return s
}

type QueryMinutesPlayInfoResponseBodyPlayInfo struct {
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	Duration    *string `json:"duration,omitempty" xml:"duration,omitempty"`
	MediaType   *string `json:"mediaType,omitempty" xml:"mediaType,omitempty"`
	PlayUrl     *string `json:"playUrl,omitempty" xml:"playUrl,omitempty"`
	Size        *string `json:"size,omitempty" xml:"size,omitempty"`
	Status      *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryMinutesPlayInfoResponseBodyPlayInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesPlayInfoResponseBodyPlayInfo) GoString() string {
	return s.String()
}

func (s *QueryMinutesPlayInfoResponseBodyPlayInfo) SetDownloadUrl(v string) *QueryMinutesPlayInfoResponseBodyPlayInfo {
	s.DownloadUrl = &v
	return s
}

func (s *QueryMinutesPlayInfoResponseBodyPlayInfo) SetDuration(v string) *QueryMinutesPlayInfoResponseBodyPlayInfo {
	s.Duration = &v
	return s
}

func (s *QueryMinutesPlayInfoResponseBodyPlayInfo) SetMediaType(v string) *QueryMinutesPlayInfoResponseBodyPlayInfo {
	s.MediaType = &v
	return s
}

func (s *QueryMinutesPlayInfoResponseBodyPlayInfo) SetPlayUrl(v string) *QueryMinutesPlayInfoResponseBodyPlayInfo {
	s.PlayUrl = &v
	return s
}

func (s *QueryMinutesPlayInfoResponseBodyPlayInfo) SetSize(v string) *QueryMinutesPlayInfoResponseBodyPlayInfo {
	s.Size = &v
	return s
}

func (s *QueryMinutesPlayInfoResponseBodyPlayInfo) SetStatus(v string) *QueryMinutesPlayInfoResponseBodyPlayInfo {
	s.Status = &v
	return s
}

type QueryMinutesPlayInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMinutesPlayInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMinutesPlayInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesPlayInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryMinutesPlayInfoResponse) SetHeaders(v map[string]*string) *QueryMinutesPlayInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryMinutesPlayInfoResponse) SetStatusCode(v int32) *QueryMinutesPlayInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMinutesPlayInfoResponse) SetBody(v *QueryMinutesPlayInfoResponseBody) *QueryMinutesPlayInfoResponse {
	s.Body = v
	return s
}

type QueryMinutesShareListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMinutesShareListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesShareListHeaders) GoString() string {
	return s.String()
}

func (s *QueryMinutesShareListHeaders) SetCommonHeaders(v map[string]*string) *QueryMinutesShareListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMinutesShareListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMinutesShareListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMinutesShareListRequest struct {
	// This parameter is required.
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Scene      *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryMinutesShareListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesShareListRequest) GoString() string {
	return s.String()
}

func (s *QueryMinutesShareListRequest) SetMaxResults(v int32) *QueryMinutesShareListRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryMinutesShareListRequest) SetNextToken(v string) *QueryMinutesShareListRequest {
	s.NextToken = &v
	return s
}

func (s *QueryMinutesShareListRequest) SetScene(v string) *QueryMinutesShareListRequest {
	s.Scene = &v
	return s
}

func (s *QueryMinutesShareListRequest) SetUnionId(v string) *QueryMinutesShareListRequest {
	s.UnionId = &v
	return s
}

type QueryMinutesShareListResponseBody struct {
	HasNext        *bool                                              `json:"hasNext,omitempty" xml:"hasNext,omitempty"`
	MinutesDetails []*QueryMinutesShareListResponseBodyMinutesDetails `json:"minutesDetails,omitempty" xml:"minutesDetails,omitempty" type:"Repeated"`
	NextToken      *string                                            `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryMinutesShareListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesShareListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMinutesShareListResponseBody) SetHasNext(v bool) *QueryMinutesShareListResponseBody {
	s.HasNext = &v
	return s
}

func (s *QueryMinutesShareListResponseBody) SetMinutesDetails(v []*QueryMinutesShareListResponseBodyMinutesDetails) *QueryMinutesShareListResponseBody {
	s.MinutesDetails = v
	return s
}

func (s *QueryMinutesShareListResponseBody) SetNextToken(v string) *QueryMinutesShareListResponseBody {
	s.NextToken = &v
	return s
}

type QueryMinutesShareListResponseBodyMinutesDetails struct {
	BizType        *int32  `json:"bizType,omitempty" xml:"bizType,omitempty"`
	CreatorNick    *string `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	CreatorUnionId *string `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	DurationMicros *int64  `json:"durationMicros,omitempty" xml:"durationMicros,omitempty"`
	IsDeleted      *int32  `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	Size           *int64  `json:"size,omitempty" xml:"size,omitempty"`
	StartTime      *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	TaskUuid       *string `json:"taskUuid,omitempty" xml:"taskUuid,omitempty"`
	Title          *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryMinutesShareListResponseBodyMinutesDetails) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesShareListResponseBodyMinutesDetails) GoString() string {
	return s.String()
}

func (s *QueryMinutesShareListResponseBodyMinutesDetails) SetBizType(v int32) *QueryMinutesShareListResponseBodyMinutesDetails {
	s.BizType = &v
	return s
}

func (s *QueryMinutesShareListResponseBodyMinutesDetails) SetCreatorNick(v string) *QueryMinutesShareListResponseBodyMinutesDetails {
	s.CreatorNick = &v
	return s
}

func (s *QueryMinutesShareListResponseBodyMinutesDetails) SetCreatorUnionId(v string) *QueryMinutesShareListResponseBodyMinutesDetails {
	s.CreatorUnionId = &v
	return s
}

func (s *QueryMinutesShareListResponseBodyMinutesDetails) SetDurationMicros(v int64) *QueryMinutesShareListResponseBodyMinutesDetails {
	s.DurationMicros = &v
	return s
}

func (s *QueryMinutesShareListResponseBodyMinutesDetails) SetIsDeleted(v int32) *QueryMinutesShareListResponseBodyMinutesDetails {
	s.IsDeleted = &v
	return s
}

func (s *QueryMinutesShareListResponseBodyMinutesDetails) SetSize(v int64) *QueryMinutesShareListResponseBodyMinutesDetails {
	s.Size = &v
	return s
}

func (s *QueryMinutesShareListResponseBodyMinutesDetails) SetStartTime(v int64) *QueryMinutesShareListResponseBodyMinutesDetails {
	s.StartTime = &v
	return s
}

func (s *QueryMinutesShareListResponseBodyMinutesDetails) SetStatus(v int32) *QueryMinutesShareListResponseBodyMinutesDetails {
	s.Status = &v
	return s
}

func (s *QueryMinutesShareListResponseBodyMinutesDetails) SetTaskUuid(v string) *QueryMinutesShareListResponseBodyMinutesDetails {
	s.TaskUuid = &v
	return s
}

func (s *QueryMinutesShareListResponseBodyMinutesDetails) SetTitle(v string) *QueryMinutesShareListResponseBodyMinutesDetails {
	s.Title = &v
	return s
}

type QueryMinutesShareListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMinutesShareListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMinutesShareListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesShareListResponse) GoString() string {
	return s.String()
}

func (s *QueryMinutesShareListResponse) SetHeaders(v map[string]*string) *QueryMinutesShareListResponse {
	s.Headers = v
	return s
}

func (s *QueryMinutesShareListResponse) SetStatusCode(v int32) *QueryMinutesShareListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMinutesShareListResponse) SetBody(v *QueryMinutesShareListResponseBody) *QueryMinutesShareListResponse {
	s.Body = v
	return s
}

type QueryMinutesStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMinutesStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesStatusHeaders) GoString() string {
	return s.String()
}

func (s *QueryMinutesStatusHeaders) SetCommonHeaders(v map[string]*string) *QueryMinutesStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMinutesStatusHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMinutesStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMinutesStatusRequest struct {
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryMinutesStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryMinutesStatusRequest) SetUnionId(v string) *QueryMinutesStatusRequest {
	s.UnionId = &v
	return s
}

type QueryMinutesStatusResponseBody struct {
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryMinutesStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMinutesStatusResponseBody) SetStatus(v int32) *QueryMinutesStatusResponseBody {
	s.Status = &v
	return s
}

type QueryMinutesStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMinutesStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMinutesStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryMinutesStatusResponse) SetHeaders(v map[string]*string) *QueryMinutesStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryMinutesStatusResponse) SetStatusCode(v int32) *QueryMinutesStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMinutesStatusResponse) SetBody(v *QueryMinutesStatusResponseBody) *QueryMinutesStatusResponse {
	s.Body = v
	return s
}

type QueryMinutesTextHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMinutesTextHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesTextHeaders) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextHeaders) SetCommonHeaders(v map[string]*string) *QueryMinutesTextHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMinutesTextHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMinutesTextHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMinutesTextRequest struct {
	Direction  *int32  `json:"direction,omitempty" xml:"direction,omitempty"`
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryMinutesTextRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesTextRequest) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextRequest) SetDirection(v int32) *QueryMinutesTextRequest {
	s.Direction = &v
	return s
}

func (s *QueryMinutesTextRequest) SetMaxResults(v int32) *QueryMinutesTextRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryMinutesTextRequest) SetNextToken(v string) *QueryMinutesTextRequest {
	s.NextToken = &v
	return s
}

func (s *QueryMinutesTextRequest) SetUnionId(v string) *QueryMinutesTextRequest {
	s.UnionId = &v
	return s
}

type QueryMinutesTextResponseBody struct {
	HasNext       *bool                                        `json:"hasNext,omitempty" xml:"hasNext,omitempty"`
	NextToken     *string                                      `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ParagraphList []*QueryMinutesTextResponseBodyParagraphList `json:"paragraphList,omitempty" xml:"paragraphList,omitempty" type:"Repeated"`
}

func (s QueryMinutesTextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesTextResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextResponseBody) SetHasNext(v bool) *QueryMinutesTextResponseBody {
	s.HasNext = &v
	return s
}

func (s *QueryMinutesTextResponseBody) SetNextToken(v string) *QueryMinutesTextResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryMinutesTextResponseBody) SetParagraphList(v []*QueryMinutesTextResponseBodyParagraphList) *QueryMinutesTextResponseBody {
	s.ParagraphList = v
	return s
}

type QueryMinutesTextResponseBodyParagraphList struct {
	EndTime      *int64                                                   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	NickName     *string                                                  `json:"nickName,omitempty" xml:"nickName,omitempty"`
	Paragraph    *string                                                  `json:"paragraph,omitempty" xml:"paragraph,omitempty"`
	ParagraphId  *int64                                                   `json:"paragraphId,omitempty" xml:"paragraphId,omitempty"`
	RecordId     *int64                                                   `json:"recordId,omitempty" xml:"recordId,omitempty"`
	SentenceList []*QueryMinutesTextResponseBodyParagraphListSentenceList `json:"sentenceList,omitempty" xml:"sentenceList,omitempty" type:"Repeated"`
	StartTime    *int64                                                   `json:"startTime,omitempty" xml:"startTime,omitempty"`
	UnionId      *string                                                  `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryMinutesTextResponseBodyParagraphList) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesTextResponseBodyParagraphList) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetEndTime(v int64) *QueryMinutesTextResponseBodyParagraphList {
	s.EndTime = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetNickName(v string) *QueryMinutesTextResponseBodyParagraphList {
	s.NickName = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetParagraph(v string) *QueryMinutesTextResponseBodyParagraphList {
	s.Paragraph = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetParagraphId(v int64) *QueryMinutesTextResponseBodyParagraphList {
	s.ParagraphId = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetRecordId(v int64) *QueryMinutesTextResponseBodyParagraphList {
	s.RecordId = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetSentenceList(v []*QueryMinutesTextResponseBodyParagraphListSentenceList) *QueryMinutesTextResponseBodyParagraphList {
	s.SentenceList = v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetStartTime(v int64) *QueryMinutesTextResponseBodyParagraphList {
	s.StartTime = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetUnionId(v string) *QueryMinutesTextResponseBodyParagraphList {
	s.UnionId = &v
	return s
}

type QueryMinutesTextResponseBodyParagraphListSentenceList struct {
	EndTime   *int64                                                           `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Sentence  *string                                                          `json:"sentence,omitempty" xml:"sentence,omitempty"`
	StartTime *int64                                                           `json:"startTime,omitempty" xml:"startTime,omitempty"`
	WordList  []*QueryMinutesTextResponseBodyParagraphListSentenceListWordList `json:"wordList,omitempty" xml:"wordList,omitempty" type:"Repeated"`
}

func (s QueryMinutesTextResponseBodyParagraphListSentenceList) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesTextResponseBodyParagraphListSentenceList) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceList) SetEndTime(v int64) *QueryMinutesTextResponseBodyParagraphListSentenceList {
	s.EndTime = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceList) SetSentence(v string) *QueryMinutesTextResponseBodyParagraphListSentenceList {
	s.Sentence = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceList) SetStartTime(v int64) *QueryMinutesTextResponseBodyParagraphListSentenceList {
	s.StartTime = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceList) SetWordList(v []*QueryMinutesTextResponseBodyParagraphListSentenceListWordList) *QueryMinutesTextResponseBodyParagraphListSentenceList {
	s.WordList = v
	return s
}

type QueryMinutesTextResponseBodyParagraphListSentenceListWordList struct {
	EndTime   *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Word      *string `json:"word,omitempty" xml:"word,omitempty"`
	WordId    *string `json:"wordId,omitempty" xml:"wordId,omitempty"`
}

func (s QueryMinutesTextResponseBodyParagraphListSentenceListWordList) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesTextResponseBodyParagraphListSentenceListWordList) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceListWordList) SetEndTime(v int64) *QueryMinutesTextResponseBodyParagraphListSentenceListWordList {
	s.EndTime = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceListWordList) SetStartTime(v int64) *QueryMinutesTextResponseBodyParagraphListSentenceListWordList {
	s.StartTime = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceListWordList) SetWord(v string) *QueryMinutesTextResponseBodyParagraphListSentenceListWordList {
	s.Word = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceListWordList) SetWordId(v string) *QueryMinutesTextResponseBodyParagraphListSentenceListWordList {
	s.WordId = &v
	return s
}

type QueryMinutesTextResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMinutesTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMinutesTextResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesTextResponse) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextResponse) SetHeaders(v map[string]*string) *QueryMinutesTextResponse {
	s.Headers = v
	return s
}

func (s *QueryMinutesTextResponse) SetStatusCode(v int32) *QueryMinutesTextResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMinutesTextResponse) SetBody(v *QueryMinutesTextResponseBody) *QueryMinutesTextResponse {
	s.Body = v
	return s
}

type QueryUploadVideoPlayInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUploadVideoPlayInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUploadVideoPlayInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryUploadVideoPlayInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryUploadVideoPlayInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUploadVideoPlayInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUploadVideoPlayInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUploadVideoPlayInfoRequest struct {
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryUploadVideoPlayInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUploadVideoPlayInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryUploadVideoPlayInfoRequest) SetUnionId(v string) *QueryUploadVideoPlayInfoRequest {
	s.UnionId = &v
	return s
}

type QueryUploadVideoPlayInfoResponseBody struct {
	PlayInfo *QueryUploadVideoPlayInfoResponseBodyPlayInfo `json:"playInfo,omitempty" xml:"playInfo,omitempty" type:"Struct"`
}

func (s QueryUploadVideoPlayInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUploadVideoPlayInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUploadVideoPlayInfoResponseBody) SetPlayInfo(v *QueryUploadVideoPlayInfoResponseBodyPlayInfo) *QueryUploadVideoPlayInfoResponseBody {
	s.PlayInfo = v
	return s
}

type QueryUploadVideoPlayInfoResponseBodyPlayInfo struct {
	Duration *int64  `json:"duration,omitempty" xml:"duration,omitempty"`
	PlayUrl  *string `json:"playUrl,omitempty" xml:"playUrl,omitempty"`
	Size     *int64  `json:"size,omitempty" xml:"size,omitempty"`
	Status   *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryUploadVideoPlayInfoResponseBodyPlayInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryUploadVideoPlayInfoResponseBodyPlayInfo) GoString() string {
	return s.String()
}

func (s *QueryUploadVideoPlayInfoResponseBodyPlayInfo) SetDuration(v int64) *QueryUploadVideoPlayInfoResponseBodyPlayInfo {
	s.Duration = &v
	return s
}

func (s *QueryUploadVideoPlayInfoResponseBodyPlayInfo) SetPlayUrl(v string) *QueryUploadVideoPlayInfoResponseBodyPlayInfo {
	s.PlayUrl = &v
	return s
}

func (s *QueryUploadVideoPlayInfoResponseBodyPlayInfo) SetSize(v int64) *QueryUploadVideoPlayInfoResponseBodyPlayInfo {
	s.Size = &v
	return s
}

func (s *QueryUploadVideoPlayInfoResponseBodyPlayInfo) SetStatus(v string) *QueryUploadVideoPlayInfoResponseBodyPlayInfo {
	s.Status = &v
	return s
}

type QueryUploadVideoPlayInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUploadVideoPlayInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUploadVideoPlayInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUploadVideoPlayInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryUploadVideoPlayInfoResponse) SetHeaders(v map[string]*string) *QueryUploadVideoPlayInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryUploadVideoPlayInfoResponse) SetStatusCode(v int32) *QueryUploadVideoPlayInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUploadVideoPlayInfoResponse) SetBody(v *QueryUploadVideoPlayInfoResponseBody) *QueryUploadVideoPlayInfoResponse {
	s.Body = v
	return s
}

type UpdateMinutesTitleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateMinutesTitleHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateMinutesTitleHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMinutesTitleHeaders) SetCommonHeaders(v map[string]*string) *UpdateMinutesTitleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMinutesTitleHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateMinutesTitleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateMinutesTitleRequest struct {
	// This parameter is required.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s UpdateMinutesTitleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMinutesTitleRequest) GoString() string {
	return s.String()
}

func (s *UpdateMinutesTitleRequest) SetTitle(v string) *UpdateMinutesTitleRequest {
	s.Title = &v
	return s
}

func (s *UpdateMinutesTitleRequest) SetUnionId(v string) *UpdateMinutesTitleRequest {
	s.UnionId = &v
	return s
}

type UpdateMinutesTitleResponseBody struct {
	TaskUuid *string `json:"taskUuid,omitempty" xml:"taskUuid,omitempty"`
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s UpdateMinutesTitleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMinutesTitleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMinutesTitleResponseBody) SetTaskUuid(v string) *UpdateMinutesTitleResponseBody {
	s.TaskUuid = &v
	return s
}

func (s *UpdateMinutesTitleResponseBody) SetTitle(v string) *UpdateMinutesTitleResponseBody {
	s.Title = &v
	return s
}

type UpdateMinutesTitleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMinutesTitleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMinutesTitleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMinutesTitleResponse) GoString() string {
	return s.String()
}

func (s *UpdateMinutesTitleResponse) SetHeaders(v map[string]*string) *UpdateMinutesTitleResponse {
	s.Headers = v
	return s
}

func (s *UpdateMinutesTitleResponse) SetStatusCode(v int32) *UpdateMinutesTitleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMinutesTitleResponse) SetBody(v *UpdateMinutesTitleResponseBody) *UpdateMinutesTitleResponse {
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
// 批量获取闪记详情
//
// @param request - BatchGetMinutesDetailsRequest
//
// @param headers - BatchGetMinutesDetailsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetMinutesDetailsResponse
func (client *Client) BatchGetMinutesDetailsWithOptions(request *BatchGetMinutesDetailsRequest, headers *BatchGetMinutesDetailsHeaders, runtime *util.RuntimeOptions) (_result *BatchGetMinutesDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskUuids)) {
		body["taskUuids"] = request.TaskUuids
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
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchGetMinutesDetails"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/details/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchGetMinutesDetailsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量获取闪记详情
//
// @param request - BatchGetMinutesDetailsRequest
//
// @return BatchGetMinutesDetailsResponse
func (client *Client) BatchGetMinutesDetails(request *BatchGetMinutesDetailsRequest) (_result *BatchGetMinutesDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchGetMinutesDetailsHeaders{}
	_result = &BatchGetMinutesDetailsResponse{}
	_body, _err := client.BatchGetMinutesDetailsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除闪记
//
// @param request - DeleteMinutesRequest
//
// @param headers - DeleteMinutesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMinutesResponse
func (client *Client) DeleteMinutesWithOptions(taskUuid *string, request *DeleteMinutesRequest, headers *DeleteMinutesHeaders, runtime *util.RuntimeOptions) (_result *DeleteMinutesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
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
		Action:      tea.String("DeleteMinutes"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/tasks/" + tea.StringValue(taskUuid)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMinutesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除闪记
//
// @param request - DeleteMinutesRequest
//
// @return DeleteMinutesResponse
func (client *Client) DeleteMinutes(taskUuid *string, request *DeleteMinutesRequest) (_result *DeleteMinutesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteMinutesHeaders{}
	_result = &DeleteMinutesResponse{}
	_body, _err := client.DeleteMinutesWithOptions(taskUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询自己创建的闪记列表
//
// @param request - QueryCreateMinutesListRequest
//
// @param headers - QueryCreateMinutesListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCreateMinutesListResponse
func (client *Client) QueryCreateMinutesListWithOptions(request *QueryCreateMinutesListRequest, headers *QueryCreateMinutesListHeaders, runtime *util.RuntimeOptions) (_result *QueryCreateMinutesListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
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
		Action:      tea.String("QueryCreateMinutesList"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/createLists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCreateMinutesListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询自己创建的闪记列表
//
// @param request - QueryCreateMinutesListRequest
//
// @return QueryCreateMinutesListResponse
func (client *Client) QueryCreateMinutesList(request *QueryCreateMinutesListRequest) (_result *QueryCreateMinutesListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCreateMinutesListHeaders{}
	_result = &QueryCreateMinutesListResponse{}
	_body, _err := client.QueryCreateMinutesListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询闪记媒体播放信息
//
// @param request - QueryMinutesPlayInfoRequest
//
// @param headers - QueryMinutesPlayInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMinutesPlayInfoResponse
func (client *Client) QueryMinutesPlayInfoWithOptions(taskUuid *string, request *QueryMinutesPlayInfoRequest, headers *QueryMinutesPlayInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryMinutesPlayInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
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
		Action:      tea.String("QueryMinutesPlayInfo"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/tasks/" + tea.StringValue(taskUuid) + "/playInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMinutesPlayInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询闪记媒体播放信息
//
// @param request - QueryMinutesPlayInfoRequest
//
// @return QueryMinutesPlayInfoResponse
func (client *Client) QueryMinutesPlayInfo(taskUuid *string, request *QueryMinutesPlayInfoRequest) (_result *QueryMinutesPlayInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMinutesPlayInfoHeaders{}
	_result = &QueryMinutesPlayInfoResponse{}
	_body, _err := client.QueryMinutesPlayInfoWithOptions(taskUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取与我共享闪记列表
//
// @param request - QueryMinutesShareListRequest
//
// @param headers - QueryMinutesShareListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMinutesShareListResponse
func (client *Client) QueryMinutesShareListWithOptions(request *QueryMinutesShareListRequest, headers *QueryMinutesShareListHeaders, runtime *util.RuntimeOptions) (_result *QueryMinutesShareListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		query["scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
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
		Action:      tea.String("QueryMinutesShareList"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/shareLists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMinutesShareListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取与我共享闪记列表
//
// @param request - QueryMinutesShareListRequest
//
// @return QueryMinutesShareListResponse
func (client *Client) QueryMinutesShareList(request *QueryMinutesShareListRequest) (_result *QueryMinutesShareListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMinutesShareListHeaders{}
	_result = &QueryMinutesShareListResponse{}
	_body, _err := client.QueryMinutesShareListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询闪记状态
//
// @param request - QueryMinutesStatusRequest
//
// @param headers - QueryMinutesStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMinutesStatusResponse
func (client *Client) QueryMinutesStatusWithOptions(taskUuid *string, request *QueryMinutesStatusRequest, headers *QueryMinutesStatusHeaders, runtime *util.RuntimeOptions) (_result *QueryMinutesStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
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
		Action:      tea.String("QueryMinutesStatus"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/" + tea.StringValue(taskUuid) + "/taskStatus"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMinutesStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询闪记状态
//
// @param request - QueryMinutesStatusRequest
//
// @return QueryMinutesStatusResponse
func (client *Client) QueryMinutesStatus(taskUuid *string, request *QueryMinutesStatusRequest) (_result *QueryMinutesStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMinutesStatusHeaders{}
	_result = &QueryMinutesStatusResponse{}
	_body, _err := client.QueryMinutesStatusWithOptions(taskUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询闪记转写文本内容
//
// @param request - QueryMinutesTextRequest
//
// @param headers - QueryMinutesTextHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMinutesTextResponse
func (client *Client) QueryMinutesTextWithOptions(taskUuid *string, request *QueryMinutesTextRequest, headers *QueryMinutesTextHeaders, runtime *util.RuntimeOptions) (_result *QueryMinutesTextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
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
		Action:      tea.String("QueryMinutesText"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/" + tea.StringValue(taskUuid) + "/textInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMinutesTextResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询闪记转写文本内容
//
// @param request - QueryMinutesTextRequest
//
// @return QueryMinutesTextResponse
func (client *Client) QueryMinutesText(taskUuid *string, request *QueryMinutesTextRequest) (_result *QueryMinutesTextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMinutesTextHeaders{}
	_result = &QueryMinutesTextResponse{}
	_body, _err := client.QueryMinutesTextWithOptions(taskUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询上传视频播放信息
//
// @param request - QueryUploadVideoPlayInfoRequest
//
// @param headers - QueryUploadVideoPlayInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUploadVideoPlayInfoResponse
func (client *Client) QueryUploadVideoPlayInfoWithOptions(videoId *string, request *QueryUploadVideoPlayInfoRequest, headers *QueryUploadVideoPlayInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryUploadVideoPlayInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
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
		Action:      tea.String("QueryUploadVideoPlayInfo"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/uploadVideos/" + tea.StringValue(videoId) + "/playInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUploadVideoPlayInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询上传视频播放信息
//
// @param request - QueryUploadVideoPlayInfoRequest
//
// @return QueryUploadVideoPlayInfoResponse
func (client *Client) QueryUploadVideoPlayInfo(videoId *string, request *QueryUploadVideoPlayInfoRequest) (_result *QueryUploadVideoPlayInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUploadVideoPlayInfoHeaders{}
	_result = &QueryUploadVideoPlayInfoResponse{}
	_body, _err := client.QueryUploadVideoPlayInfoWithOptions(videoId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新闪记标题
//
// @param request - UpdateMinutesTitleRequest
//
// @param headers - UpdateMinutesTitleHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMinutesTitleResponse
func (client *Client) UpdateMinutesTitleWithOptions(taskUuid *string, request *UpdateMinutesTitleRequest, headers *UpdateMinutesTitleHeaders, runtime *util.RuntimeOptions) (_result *UpdateMinutesTitleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
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
		Action:      tea.String("UpdateMinutesTitle"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/tasks/" + tea.StringValue(taskUuid) + "/titles"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMinutesTitleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新闪记标题
//
// @param request - UpdateMinutesTitleRequest
//
// @return UpdateMinutesTitleResponse
func (client *Client) UpdateMinutesTitle(taskUuid *string, request *UpdateMinutesTitleRequest) (_result *UpdateMinutesTitleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateMinutesTitleHeaders{}
	_result = &UpdateMinutesTitleResponse{}
	_body, _err := client.UpdateMinutesTitleWithOptions(taskUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
