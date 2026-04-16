// This file is auto-generated, don't edit it. Thanks.
package minutes_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AdminSearchMinutesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AdminSearchMinutesHeaders) String() string {
	return tea.Prettify(s)
}

func (s AdminSearchMinutesHeaders) GoString() string {
	return s.String()
}

func (s *AdminSearchMinutesHeaders) SetCommonHeaders(v map[string]*string) *AdminSearchMinutesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AdminSearchMinutesHeaders) SetXAcsDingtalkAccessToken(v string) *AdminSearchMinutesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AdminSearchMinutesRequest struct {
	CreatorUnionIds []*string `json:"creatorUnionIds,omitempty" xml:"creatorUnionIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1700100000000
	EndTime   *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 会议纪要
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// example:
	//
	// fulltext
	SearchType *string `json:"searchType,omitempty" xml:"searchType,omitempty"`
	// example:
	//
	// 1700000000000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lJcRnm39OsU4jlFVmRGXXXXX
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s AdminSearchMinutesRequest) String() string {
	return tea.Prettify(s)
}

func (s AdminSearchMinutesRequest) GoString() string {
	return s.String()
}

func (s *AdminSearchMinutesRequest) SetCreatorUnionIds(v []*string) *AdminSearchMinutesRequest {
	s.CreatorUnionIds = v
	return s
}

func (s *AdminSearchMinutesRequest) SetEndTime(v int64) *AdminSearchMinutesRequest {
	s.EndTime = &v
	return s
}

func (s *AdminSearchMinutesRequest) SetNextToken(v string) *AdminSearchMinutesRequest {
	s.NextToken = &v
	return s
}

func (s *AdminSearchMinutesRequest) SetPageSize(v int32) *AdminSearchMinutesRequest {
	s.PageSize = &v
	return s
}

func (s *AdminSearchMinutesRequest) SetQuery(v string) *AdminSearchMinutesRequest {
	s.Query = &v
	return s
}

func (s *AdminSearchMinutesRequest) SetSearchType(v string) *AdminSearchMinutesRequest {
	s.SearchType = &v
	return s
}

func (s *AdminSearchMinutesRequest) SetStartTime(v int64) *AdminSearchMinutesRequest {
	s.StartTime = &v
	return s
}

func (s *AdminSearchMinutesRequest) SetUnionId(v string) *AdminSearchMinutesRequest {
	s.UnionId = &v
	return s
}

type AdminSearchMinutesResponseBody struct {
	HasMore     *bool                                        `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	MinutesList []*AdminSearchMinutesResponseBodyMinutesList `json:"minutesList,omitempty" xml:"minutesList,omitempty" type:"Repeated"`
	NextToken   *string                                      `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s AdminSearchMinutesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AdminSearchMinutesResponseBody) GoString() string {
	return s.String()
}

func (s *AdminSearchMinutesResponseBody) SetHasMore(v bool) *AdminSearchMinutesResponseBody {
	s.HasMore = &v
	return s
}

func (s *AdminSearchMinutesResponseBody) SetMinutesList(v []*AdminSearchMinutesResponseBodyMinutesList) *AdminSearchMinutesResponseBody {
	s.MinutesList = v
	return s
}

func (s *AdminSearchMinutesResponseBody) SetNextToken(v string) *AdminSearchMinutesResponseBody {
	s.NextToken = &v
	return s
}

type AdminSearchMinutesResponseBodyMinutesList struct {
	BizType         *int32  `json:"bizType,omitempty" xml:"bizType,omitempty"`
	CreatorNick     *string `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	CreatorUnionId  *string `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	Duration        *int64  `json:"duration,omitempty" xml:"duration,omitempty"`
	FullTextSummary *string `json:"fullTextSummary,omitempty" xml:"fullTextSummary,omitempty"`
	OriginalText    *string `json:"originalText,omitempty" xml:"originalText,omitempty"`
	StartTime       *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status          *int32  `json:"status,omitempty" xml:"status,omitempty"`
	TaskUuid        *string `json:"taskUuid,omitempty" xml:"taskUuid,omitempty"`
	Title           *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s AdminSearchMinutesResponseBodyMinutesList) String() string {
	return tea.Prettify(s)
}

func (s AdminSearchMinutesResponseBodyMinutesList) GoString() string {
	return s.String()
}

func (s *AdminSearchMinutesResponseBodyMinutesList) SetBizType(v int32) *AdminSearchMinutesResponseBodyMinutesList {
	s.BizType = &v
	return s
}

func (s *AdminSearchMinutesResponseBodyMinutesList) SetCreatorNick(v string) *AdminSearchMinutesResponseBodyMinutesList {
	s.CreatorNick = &v
	return s
}

func (s *AdminSearchMinutesResponseBodyMinutesList) SetCreatorUnionId(v string) *AdminSearchMinutesResponseBodyMinutesList {
	s.CreatorUnionId = &v
	return s
}

func (s *AdminSearchMinutesResponseBodyMinutesList) SetDuration(v int64) *AdminSearchMinutesResponseBodyMinutesList {
	s.Duration = &v
	return s
}

func (s *AdminSearchMinutesResponseBodyMinutesList) SetFullTextSummary(v string) *AdminSearchMinutesResponseBodyMinutesList {
	s.FullTextSummary = &v
	return s
}

func (s *AdminSearchMinutesResponseBodyMinutesList) SetOriginalText(v string) *AdminSearchMinutesResponseBodyMinutesList {
	s.OriginalText = &v
	return s
}

func (s *AdminSearchMinutesResponseBodyMinutesList) SetStartTime(v int64) *AdminSearchMinutesResponseBodyMinutesList {
	s.StartTime = &v
	return s
}

func (s *AdminSearchMinutesResponseBodyMinutesList) SetStatus(v int32) *AdminSearchMinutesResponseBodyMinutesList {
	s.Status = &v
	return s
}

func (s *AdminSearchMinutesResponseBodyMinutesList) SetTaskUuid(v string) *AdminSearchMinutesResponseBodyMinutesList {
	s.TaskUuid = &v
	return s
}

func (s *AdminSearchMinutesResponseBodyMinutesList) SetTitle(v string) *AdminSearchMinutesResponseBodyMinutesList {
	s.Title = &v
	return s
}

type AdminSearchMinutesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AdminSearchMinutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AdminSearchMinutesResponse) String() string {
	return tea.Prettify(s)
}

func (s AdminSearchMinutesResponse) GoString() string {
	return s.String()
}

func (s *AdminSearchMinutesResponse) SetHeaders(v map[string]*string) *AdminSearchMinutesResponse {
	s.Headers = v
	return s
}

func (s *AdminSearchMinutesResponse) SetStatusCode(v int32) *AdminSearchMinutesResponse {
	s.StatusCode = &v
	return s
}

func (s *AdminSearchMinutesResponse) SetBody(v *AdminSearchMinutesResponseBody) *AdminSearchMinutesResponse {
	s.Body = v
	return s
}

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

type BatchGetVoicePrintIdentifyConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchGetVoicePrintIdentifyConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchGetVoicePrintIdentifyConfigHeaders) GoString() string {
	return s.String()
}

func (s *BatchGetVoicePrintIdentifyConfigHeaders) SetCommonHeaders(v map[string]*string) *BatchGetVoicePrintIdentifyConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchGetVoicePrintIdentifyConfigHeaders) SetXAcsDingtalkAccessToken(v string) *BatchGetVoicePrintIdentifyConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchGetVoicePrintIdentifyConfigRequest struct {
	// This parameter is required.
	Items []*BatchGetVoicePrintIdentifyConfigRequestItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s BatchGetVoicePrintIdentifyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetVoicePrintIdentifyConfigRequest) GoString() string {
	return s.String()
}

func (s *BatchGetVoicePrintIdentifyConfigRequest) SetItems(v []*BatchGetVoicePrintIdentifyConfigRequestItems) *BatchGetVoicePrintIdentifyConfigRequest {
	s.Items = v
	return s
}

type BatchGetVoicePrintIdentifyConfigRequestItems struct {
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s BatchGetVoicePrintIdentifyConfigRequestItems) String() string {
	return tea.Prettify(s)
}

func (s BatchGetVoicePrintIdentifyConfigRequestItems) GoString() string {
	return s.String()
}

func (s *BatchGetVoicePrintIdentifyConfigRequestItems) SetLang(v string) *BatchGetVoicePrintIdentifyConfigRequestItems {
	s.Lang = &v
	return s
}

func (s *BatchGetVoicePrintIdentifyConfigRequestItems) SetUnionId(v string) *BatchGetVoicePrintIdentifyConfigRequestItems {
	s.UnionId = &v
	return s
}

type BatchGetVoicePrintIdentifyConfigResponseBody struct {
	ConfigItems []*BatchGetVoicePrintIdentifyConfigResponseBodyConfigItems `json:"configItems,omitempty" xml:"configItems,omitempty" type:"Repeated"`
}

func (s BatchGetVoicePrintIdentifyConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchGetVoicePrintIdentifyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetVoicePrintIdentifyConfigResponseBody) SetConfigItems(v []*BatchGetVoicePrintIdentifyConfigResponseBodyConfigItems) *BatchGetVoicePrintIdentifyConfigResponseBody {
	s.ConfigItems = v
	return s
}

type BatchGetVoicePrintIdentifyConfigResponseBodyConfigItems struct {
	AllowConfigVoicePrint *bool   `json:"allowConfigVoicePrint,omitempty" xml:"allowConfigVoicePrint,omitempty"`
	EnableVoicePrint      *bool   `json:"enableVoicePrint,omitempty" xml:"enableVoicePrint,omitempty"`
	HasVoicePrintRecord   *bool   `json:"hasVoicePrintRecord,omitempty" xml:"hasVoicePrintRecord,omitempty"`
	UnionId               *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s BatchGetVoicePrintIdentifyConfigResponseBodyConfigItems) String() string {
	return tea.Prettify(s)
}

func (s BatchGetVoicePrintIdentifyConfigResponseBodyConfigItems) GoString() string {
	return s.String()
}

func (s *BatchGetVoicePrintIdentifyConfigResponseBodyConfigItems) SetAllowConfigVoicePrint(v bool) *BatchGetVoicePrintIdentifyConfigResponseBodyConfigItems {
	s.AllowConfigVoicePrint = &v
	return s
}

func (s *BatchGetVoicePrintIdentifyConfigResponseBodyConfigItems) SetEnableVoicePrint(v bool) *BatchGetVoicePrintIdentifyConfigResponseBodyConfigItems {
	s.EnableVoicePrint = &v
	return s
}

func (s *BatchGetVoicePrintIdentifyConfigResponseBodyConfigItems) SetHasVoicePrintRecord(v bool) *BatchGetVoicePrintIdentifyConfigResponseBodyConfigItems {
	s.HasVoicePrintRecord = &v
	return s
}

func (s *BatchGetVoicePrintIdentifyConfigResponseBodyConfigItems) SetUnionId(v string) *BatchGetVoicePrintIdentifyConfigResponseBodyConfigItems {
	s.UnionId = &v
	return s
}

type BatchGetVoicePrintIdentifyConfigResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetVoicePrintIdentifyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetVoicePrintIdentifyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchGetVoicePrintIdentifyConfigResponse) GoString() string {
	return s.String()
}

func (s *BatchGetVoicePrintIdentifyConfigResponse) SetHeaders(v map[string]*string) *BatchGetVoicePrintIdentifyConfigResponse {
	s.Headers = v
	return s
}

func (s *BatchGetVoicePrintIdentifyConfigResponse) SetStatusCode(v int32) *BatchGetVoicePrintIdentifyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetVoicePrintIdentifyConfigResponse) SetBody(v *BatchGetVoicePrintIdentifyConfigResponseBody) *BatchGetVoicePrintIdentifyConfigResponse {
	s.Body = v
	return s
}

type BatchToggleVoicePrintSwitchStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchToggleVoicePrintSwitchStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchToggleVoicePrintSwitchStatusHeaders) GoString() string {
	return s.String()
}

func (s *BatchToggleVoicePrintSwitchStatusHeaders) SetCommonHeaders(v map[string]*string) *BatchToggleVoicePrintSwitchStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchToggleVoicePrintSwitchStatusHeaders) SetXAcsDingtalkAccessToken(v string) *BatchToggleVoicePrintSwitchStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchToggleVoicePrintSwitchStatusRequest struct {
	// This parameter is required.
	Items []*BatchToggleVoicePrintSwitchStatusRequestItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s BatchToggleVoicePrintSwitchStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchToggleVoicePrintSwitchStatusRequest) GoString() string {
	return s.String()
}

func (s *BatchToggleVoicePrintSwitchStatusRequest) SetItems(v []*BatchToggleVoicePrintSwitchStatusRequestItems) *BatchToggleVoicePrintSwitchStatusRequest {
	s.Items = v
	return s
}

type BatchToggleVoicePrintSwitchStatusRequestItems struct {
	// This parameter is required.
	Open                  *bool `json:"open,omitempty" xml:"open,omitempty"`
	ShouldClearVoicePrint *bool `json:"shouldClearVoicePrint,omitempty" xml:"shouldClearVoicePrint,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s BatchToggleVoicePrintSwitchStatusRequestItems) String() string {
	return tea.Prettify(s)
}

func (s BatchToggleVoicePrintSwitchStatusRequestItems) GoString() string {
	return s.String()
}

func (s *BatchToggleVoicePrintSwitchStatusRequestItems) SetOpen(v bool) *BatchToggleVoicePrintSwitchStatusRequestItems {
	s.Open = &v
	return s
}

func (s *BatchToggleVoicePrintSwitchStatusRequestItems) SetShouldClearVoicePrint(v bool) *BatchToggleVoicePrintSwitchStatusRequestItems {
	s.ShouldClearVoicePrint = &v
	return s
}

func (s *BatchToggleVoicePrintSwitchStatusRequestItems) SetUnionId(v string) *BatchToggleVoicePrintSwitchStatusRequestItems {
	s.UnionId = &v
	return s
}

type BatchToggleVoicePrintSwitchStatusResponseBody struct {
	ResultItems []*BatchToggleVoicePrintSwitchStatusResponseBodyResultItems `json:"resultItems,omitempty" xml:"resultItems,omitempty" type:"Repeated"`
}

func (s BatchToggleVoicePrintSwitchStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchToggleVoicePrintSwitchStatusResponseBody) GoString() string {
	return s.String()
}

func (s *BatchToggleVoicePrintSwitchStatusResponseBody) SetResultItems(v []*BatchToggleVoicePrintSwitchStatusResponseBodyResultItems) *BatchToggleVoicePrintSwitchStatusResponseBody {
	s.ResultItems = v
	return s
}

type BatchToggleVoicePrintSwitchStatusResponseBodyResultItems struct {
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s BatchToggleVoicePrintSwitchStatusResponseBodyResultItems) String() string {
	return tea.Prettify(s)
}

func (s BatchToggleVoicePrintSwitchStatusResponseBodyResultItems) GoString() string {
	return s.String()
}

func (s *BatchToggleVoicePrintSwitchStatusResponseBodyResultItems) SetMessage(v string) *BatchToggleVoicePrintSwitchStatusResponseBodyResultItems {
	s.Message = &v
	return s
}

func (s *BatchToggleVoicePrintSwitchStatusResponseBodyResultItems) SetSuccess(v bool) *BatchToggleVoicePrintSwitchStatusResponseBodyResultItems {
	s.Success = &v
	return s
}

func (s *BatchToggleVoicePrintSwitchStatusResponseBodyResultItems) SetUnionId(v string) *BatchToggleVoicePrintSwitchStatusResponseBodyResultItems {
	s.UnionId = &v
	return s
}

type BatchToggleVoicePrintSwitchStatusResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchToggleVoicePrintSwitchStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchToggleVoicePrintSwitchStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchToggleVoicePrintSwitchStatusResponse) GoString() string {
	return s.String()
}

func (s *BatchToggleVoicePrintSwitchStatusResponse) SetHeaders(v map[string]*string) *BatchToggleVoicePrintSwitchStatusResponse {
	s.Headers = v
	return s
}

func (s *BatchToggleVoicePrintSwitchStatusResponse) SetStatusCode(v int32) *BatchToggleVoicePrintSwitchStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchToggleVoicePrintSwitchStatusResponse) SetBody(v *BatchToggleVoicePrintSwitchStatusResponseBody) *BatchToggleVoicePrintSwitchStatusResponse {
	s.Body = v
	return s
}

type CreateMinutesByUploadFileHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateMinutesByUploadFileHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateMinutesByUploadFileHeaders) GoString() string {
	return s.String()
}

func (s *CreateMinutesByUploadFileHeaders) SetCommonHeaders(v map[string]*string) *CreateMinutesByUploadFileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateMinutesByUploadFileHeaders) SetXAcsDingtalkAccessToken(v string) *CreateMinutesByUploadFileHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateMinutesByUploadFileRequest struct {
	// This parameter is required.
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lJcRnm39OsU4jlFVmRGXXXXX
	CreatorId    *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	CustomPrompt *string `json:"customPrompt,omitempty" xml:"customPrompt,omitempty"`
	// example:
	//
	// true
	EnablePushCard *bool `json:"enablePushCard,omitempty" xml:"enablePushCard,omitempty"`
	// example:
	//
	// false
	HiddenMinutes *bool `json:"hiddenMinutes,omitempty" xml:"hiddenMinutes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://media.source/audiotominutes.ogg
	MediaFileUrl *string `json:"mediaFileUrl,omitempty" xml:"mediaFileUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// audio
	MediaType *string `json:"mediaType,omitempty" xml:"mediaType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11-20 录音
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lJcRnm39OsU4jlFVmRGXXXXX
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s CreateMinutesByUploadFileRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMinutesByUploadFileRequest) GoString() string {
	return s.String()
}

func (s *CreateMinutesByUploadFileRequest) SetBizId(v string) *CreateMinutesByUploadFileRequest {
	s.BizId = &v
	return s
}

func (s *CreateMinutesByUploadFileRequest) SetCreatorId(v string) *CreateMinutesByUploadFileRequest {
	s.CreatorId = &v
	return s
}

func (s *CreateMinutesByUploadFileRequest) SetCustomPrompt(v string) *CreateMinutesByUploadFileRequest {
	s.CustomPrompt = &v
	return s
}

func (s *CreateMinutesByUploadFileRequest) SetEnablePushCard(v bool) *CreateMinutesByUploadFileRequest {
	s.EnablePushCard = &v
	return s
}

func (s *CreateMinutesByUploadFileRequest) SetHiddenMinutes(v bool) *CreateMinutesByUploadFileRequest {
	s.HiddenMinutes = &v
	return s
}

func (s *CreateMinutesByUploadFileRequest) SetMediaFileUrl(v string) *CreateMinutesByUploadFileRequest {
	s.MediaFileUrl = &v
	return s
}

func (s *CreateMinutesByUploadFileRequest) SetMediaType(v string) *CreateMinutesByUploadFileRequest {
	s.MediaType = &v
	return s
}

func (s *CreateMinutesByUploadFileRequest) SetTitle(v string) *CreateMinutesByUploadFileRequest {
	s.Title = &v
	return s
}

func (s *CreateMinutesByUploadFileRequest) SetUnionId(v string) *CreateMinutesByUploadFileRequest {
	s.UnionId = &v
	return s
}

type CreateMinutesByUploadFileResponseBody struct {
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// example:
	//
	// 7632756964313430aaaaaaaaaaaaaaaaaaaaaaaaaa7363731333633305f35
	TaskUuid *string `json:"taskUuid,omitempty" xml:"taskUuid,omitempty"`
}

func (s CreateMinutesByUploadFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMinutesByUploadFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMinutesByUploadFileResponseBody) SetBizId(v string) *CreateMinutesByUploadFileResponseBody {
	s.BizId = &v
	return s
}

func (s *CreateMinutesByUploadFileResponseBody) SetTaskUuid(v string) *CreateMinutesByUploadFileResponseBody {
	s.TaskUuid = &v
	return s
}

type CreateMinutesByUploadFileResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMinutesByUploadFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMinutesByUploadFileResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMinutesByUploadFileResponse) GoString() string {
	return s.String()
}

func (s *CreateMinutesByUploadFileResponse) SetHeaders(v map[string]*string) *CreateMinutesByUploadFileResponse {
	s.Headers = v
	return s
}

func (s *CreateMinutesByUploadFileResponse) SetStatusCode(v int32) *CreateMinutesByUploadFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMinutesByUploadFileResponse) SetBody(v *CreateMinutesByUploadFileResponseBody) *CreateMinutesByUploadFileResponse {
	s.Body = v
	return s
}

type CreateSmartDeviceAiSummaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateSmartDeviceAiSummaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateSmartDeviceAiSummaryHeaders) GoString() string {
	return s.String()
}

func (s *CreateSmartDeviceAiSummaryHeaders) SetCommonHeaders(v map[string]*string) *CreateSmartDeviceAiSummaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSmartDeviceAiSummaryHeaders) SetXAcsDingtalkAccessToken(v string) *CreateSmartDeviceAiSummaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateSmartDeviceAiSummaryRequest struct {
	AgentId    *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	IsvContext *string `json:"isvContext,omitempty" xml:"isvContext,omitempty"`
	OpenFileId *string `json:"openFileId,omitempty" xml:"openFileId,omitempty"`
}

func (s CreateSmartDeviceAiSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSmartDeviceAiSummaryRequest) GoString() string {
	return s.String()
}

func (s *CreateSmartDeviceAiSummaryRequest) SetAgentId(v string) *CreateSmartDeviceAiSummaryRequest {
	s.AgentId = &v
	return s
}

func (s *CreateSmartDeviceAiSummaryRequest) SetInstanceId(v string) *CreateSmartDeviceAiSummaryRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSmartDeviceAiSummaryRequest) SetIsvContext(v string) *CreateSmartDeviceAiSummaryRequest {
	s.IsvContext = &v
	return s
}

func (s *CreateSmartDeviceAiSummaryRequest) SetOpenFileId(v string) *CreateSmartDeviceAiSummaryRequest {
	s.OpenFileId = &v
	return s
}

type CreateSmartDeviceAiSummaryResponseBody struct {
	Async *bool `json:"async,omitempty" xml:"async,omitempty"`
}

func (s CreateSmartDeviceAiSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSmartDeviceAiSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSmartDeviceAiSummaryResponseBody) SetAsync(v bool) *CreateSmartDeviceAiSummaryResponseBody {
	s.Async = &v
	return s
}

type CreateSmartDeviceAiSummaryResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSmartDeviceAiSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSmartDeviceAiSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSmartDeviceAiSummaryResponse) GoString() string {
	return s.String()
}

func (s *CreateSmartDeviceAiSummaryResponse) SetHeaders(v map[string]*string) *CreateSmartDeviceAiSummaryResponse {
	s.Headers = v
	return s
}

func (s *CreateSmartDeviceAiSummaryResponse) SetStatusCode(v int32) *CreateSmartDeviceAiSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSmartDeviceAiSummaryResponse) SetBody(v *CreateSmartDeviceAiSummaryResponseBody) *CreateSmartDeviceAiSummaryResponse {
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

type ExportMinutesTaskResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ExportMinutesTaskResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s ExportMinutesTaskResultHeaders) GoString() string {
	return s.String()
}

func (s *ExportMinutesTaskResultHeaders) SetCommonHeaders(v map[string]*string) *ExportMinutesTaskResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ExportMinutesTaskResultHeaders) SetXAcsDingtalkAccessToken(v string) *ExportMinutesTaskResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ExportMinutesTaskResultRequest struct {
	ExpireTime           *int64                                              `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	SummaryExportSetting *ExportMinutesTaskResultRequestSummaryExportSetting `json:"summaryExportSetting,omitempty" xml:"summaryExportSetting,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// text
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 763xxxxxx325f32
	TaskUuid *string `json:"taskUuid,omitempty" xml:"taskUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// D5xxxxxxxxxxxxxxEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s ExportMinutesTaskResultRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportMinutesTaskResultRequest) GoString() string {
	return s.String()
}

func (s *ExportMinutesTaskResultRequest) SetExpireTime(v int64) *ExportMinutesTaskResultRequest {
	s.ExpireTime = &v
	return s
}

func (s *ExportMinutesTaskResultRequest) SetSummaryExportSetting(v *ExportMinutesTaskResultRequestSummaryExportSetting) *ExportMinutesTaskResultRequest {
	s.SummaryExportSetting = v
	return s
}

func (s *ExportMinutesTaskResultRequest) SetTaskType(v string) *ExportMinutesTaskResultRequest {
	s.TaskType = &v
	return s
}

func (s *ExportMinutesTaskResultRequest) SetTaskUuid(v string) *ExportMinutesTaskResultRequest {
	s.TaskUuid = &v
	return s
}

func (s *ExportMinutesTaskResultRequest) SetUnionId(v string) *ExportMinutesTaskResultRequest {
	s.UnionId = &v
	return s
}

type ExportMinutesTaskResultRequestSummaryExportSetting struct {
	EnableBilingual *bool `json:"enableBilingual,omitempty" xml:"enableBilingual,omitempty"`
	// example:
	//
	// zh
	TargetLang *string `json:"targetLang,omitempty" xml:"targetLang,omitempty"`
}

func (s ExportMinutesTaskResultRequestSummaryExportSetting) String() string {
	return tea.Prettify(s)
}

func (s ExportMinutesTaskResultRequestSummaryExportSetting) GoString() string {
	return s.String()
}

func (s *ExportMinutesTaskResultRequestSummaryExportSetting) SetEnableBilingual(v bool) *ExportMinutesTaskResultRequestSummaryExportSetting {
	s.EnableBilingual = &v
	return s
}

func (s *ExportMinutesTaskResultRequestSummaryExportSetting) SetTargetLang(v string) *ExportMinutesTaskResultRequestSummaryExportSetting {
	s.TargetLang = &v
	return s
}

type ExportMinutesTaskResultResponseBody struct {
	MinutesDocUrl *string `json:"minutesDocUrl,omitempty" xml:"minutesDocUrl,omitempty"`
}

func (s ExportMinutesTaskResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportMinutesTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *ExportMinutesTaskResultResponseBody) SetMinutesDocUrl(v string) *ExportMinutesTaskResultResponseBody {
	s.MinutesDocUrl = &v
	return s
}

type ExportMinutesTaskResultResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportMinutesTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportMinutesTaskResultResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportMinutesTaskResultResponse) GoString() string {
	return s.String()
}

func (s *ExportMinutesTaskResultResponse) SetHeaders(v map[string]*string) *ExportMinutesTaskResultResponse {
	s.Headers = v
	return s
}

func (s *ExportMinutesTaskResultResponse) SetStatusCode(v int32) *ExportMinutesTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportMinutesTaskResultResponse) SetBody(v *ExportMinutesTaskResultResponseBody) *ExportMinutesTaskResultResponse {
	s.Body = v
	return s
}

type GenerateSummaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GenerateSummaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s GenerateSummaryHeaders) GoString() string {
	return s.String()
}

func (s *GenerateSummaryHeaders) SetCommonHeaders(v map[string]*string) *GenerateSummaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GenerateSummaryHeaders) SetXAcsDingtalkAccessToken(v string) *GenerateSummaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GenerateSummaryRequest struct {
	// example:
	//
	// 1
	DiyTemplateVersion *string `json:"diyTemplateVersion,omitempty" xml:"diyTemplateVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// meeting-assistant
	SummaryTemplateId *string `json:"summaryTemplateId,omitempty" xml:"summaryTemplateId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	SummaryTemplateType *string `json:"summaryTemplateType,omitempty" xml:"summaryTemplateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lJcRnm39OsU4jlFVmRGXXXXX
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GenerateSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateSummaryRequest) GoString() string {
	return s.String()
}

func (s *GenerateSummaryRequest) SetDiyTemplateVersion(v string) *GenerateSummaryRequest {
	s.DiyTemplateVersion = &v
	return s
}

func (s *GenerateSummaryRequest) SetSummaryTemplateId(v string) *GenerateSummaryRequest {
	s.SummaryTemplateId = &v
	return s
}

func (s *GenerateSummaryRequest) SetSummaryTemplateType(v string) *GenerateSummaryRequest {
	s.SummaryTemplateType = &v
	return s
}

func (s *GenerateSummaryRequest) SetUnionId(v string) *GenerateSummaryRequest {
	s.UnionId = &v
	return s
}

type GenerateSummaryResponseBody struct {
	SummaryText *string `json:"summaryText,omitempty" xml:"summaryText,omitempty"`
	TaskUuid    *string `json:"taskUuid,omitempty" xml:"taskUuid,omitempty"`
}

func (s GenerateSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateSummaryResponseBody) SetSummaryText(v string) *GenerateSummaryResponseBody {
	s.SummaryText = &v
	return s
}

func (s *GenerateSummaryResponseBody) SetTaskUuid(v string) *GenerateSummaryResponseBody {
	s.TaskUuid = &v
	return s
}

type GenerateSummaryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateSummaryResponse) GoString() string {
	return s.String()
}

func (s *GenerateSummaryResponse) SetHeaders(v map[string]*string) *GenerateSummaryResponse {
	s.Headers = v
	return s
}

func (s *GenerateSummaryResponse) SetStatusCode(v int32) *GenerateSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateSummaryResponse) SetBody(v *GenerateSummaryResponseBody) *GenerateSummaryResponse {
	s.Body = v
	return s
}

type OpenQueryMinutesSummaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s OpenQueryMinutesSummaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s OpenQueryMinutesSummaryHeaders) GoString() string {
	return s.String()
}

func (s *OpenQueryMinutesSummaryHeaders) SetCommonHeaders(v map[string]*string) *OpenQueryMinutesSummaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OpenQueryMinutesSummaryHeaders) SetXAcsDingtalkAccessToken(v string) *OpenQueryMinutesSummaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type OpenQueryMinutesSummaryRequest struct {
	// This parameter is required.
	TaskUuid *string `json:"taskUuid,omitempty" xml:"taskUuid,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s OpenQueryMinutesSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenQueryMinutesSummaryRequest) GoString() string {
	return s.String()
}

func (s *OpenQueryMinutesSummaryRequest) SetTaskUuid(v string) *OpenQueryMinutesSummaryRequest {
	s.TaskUuid = &v
	return s
}

func (s *OpenQueryMinutesSummaryRequest) SetUnionId(v string) *OpenQueryMinutesSummaryRequest {
	s.UnionId = &v
	return s
}

type OpenQueryMinutesSummaryResponseBody struct {
	FullSummary *string `json:"fullSummary,omitempty" xml:"fullSummary,omitempty"`
}

func (s OpenQueryMinutesSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenQueryMinutesSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *OpenQueryMinutesSummaryResponseBody) SetFullSummary(v string) *OpenQueryMinutesSummaryResponseBody {
	s.FullSummary = &v
	return s
}

type OpenQueryMinutesSummaryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenQueryMinutesSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenQueryMinutesSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenQueryMinutesSummaryResponse) GoString() string {
	return s.String()
}

func (s *OpenQueryMinutesSummaryResponse) SetHeaders(v map[string]*string) *OpenQueryMinutesSummaryResponse {
	s.Headers = v
	return s
}

func (s *OpenQueryMinutesSummaryResponse) SetStatusCode(v int32) *OpenQueryMinutesSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenQueryMinutesSummaryResponse) SetBody(v *OpenQueryMinutesSummaryResponseBody) *OpenQueryMinutesSummaryResponse {
	s.Body = v
	return s
}

type QueryBizMinutesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryBizMinutesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryBizMinutesHeaders) GoString() string {
	return s.String()
}

func (s *QueryBizMinutesHeaders) SetCommonHeaders(v map[string]*string) *QueryBizMinutesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryBizMinutesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryBizMinutesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryBizMinutesRequest struct {
	// This parameter is required.
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// This parameter is required.
	BizType *int32 `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryBizMinutesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBizMinutesRequest) GoString() string {
	return s.String()
}

func (s *QueryBizMinutesRequest) SetBizId(v string) *QueryBizMinutesRequest {
	s.BizId = &v
	return s
}

func (s *QueryBizMinutesRequest) SetBizType(v int32) *QueryBizMinutesRequest {
	s.BizType = &v
	return s
}

func (s *QueryBizMinutesRequest) SetUnionId(v string) *QueryBizMinutesRequest {
	s.UnionId = &v
	return s
}

type QueryBizMinutesResponseBody struct {
	MinutesDetails []*QueryBizMinutesResponseBodyMinutesDetails `json:"minutesDetails,omitempty" xml:"minutesDetails,omitempty" type:"Repeated"`
}

func (s QueryBizMinutesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBizMinutesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBizMinutesResponseBody) SetMinutesDetails(v []*QueryBizMinutesResponseBodyMinutesDetails) *QueryBizMinutesResponseBody {
	s.MinutesDetails = v
	return s
}

type QueryBizMinutesResponseBodyMinutesDetails struct {
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

func (s QueryBizMinutesResponseBodyMinutesDetails) String() string {
	return tea.Prettify(s)
}

func (s QueryBizMinutesResponseBodyMinutesDetails) GoString() string {
	return s.String()
}

func (s *QueryBizMinutesResponseBodyMinutesDetails) SetBizType(v int32) *QueryBizMinutesResponseBodyMinutesDetails {
	s.BizType = &v
	return s
}

func (s *QueryBizMinutesResponseBodyMinutesDetails) SetCreatorNick(v string) *QueryBizMinutesResponseBodyMinutesDetails {
	s.CreatorNick = &v
	return s
}

func (s *QueryBizMinutesResponseBodyMinutesDetails) SetCreatorUnionId(v string) *QueryBizMinutesResponseBodyMinutesDetails {
	s.CreatorUnionId = &v
	return s
}

func (s *QueryBizMinutesResponseBodyMinutesDetails) SetDurationMicros(v int64) *QueryBizMinutesResponseBodyMinutesDetails {
	s.DurationMicros = &v
	return s
}

func (s *QueryBizMinutesResponseBodyMinutesDetails) SetIsDeleted(v int32) *QueryBizMinutesResponseBodyMinutesDetails {
	s.IsDeleted = &v
	return s
}

func (s *QueryBizMinutesResponseBodyMinutesDetails) SetSize(v int64) *QueryBizMinutesResponseBodyMinutesDetails {
	s.Size = &v
	return s
}

func (s *QueryBizMinutesResponseBodyMinutesDetails) SetStartTime(v int64) *QueryBizMinutesResponseBodyMinutesDetails {
	s.StartTime = &v
	return s
}

func (s *QueryBizMinutesResponseBodyMinutesDetails) SetStatus(v int32) *QueryBizMinutesResponseBodyMinutesDetails {
	s.Status = &v
	return s
}

func (s *QueryBizMinutesResponseBodyMinutesDetails) SetTaskUuid(v string) *QueryBizMinutesResponseBodyMinutesDetails {
	s.TaskUuid = &v
	return s
}

func (s *QueryBizMinutesResponseBodyMinutesDetails) SetTitle(v string) *QueryBizMinutesResponseBodyMinutesDetails {
	s.Title = &v
	return s
}

type QueryBizMinutesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBizMinutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBizMinutesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBizMinutesResponse) GoString() string {
	return s.String()
}

func (s *QueryBizMinutesResponse) SetHeaders(v map[string]*string) *QueryBizMinutesResponse {
	s.Headers = v
	return s
}

func (s *QueryBizMinutesResponse) SetStatusCode(v int32) *QueryBizMinutesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBizMinutesResponse) SetBody(v *QueryBizMinutesResponseBody) *QueryBizMinutesResponse {
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

type QueryMinutesBasicInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMinutesBasicInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesBasicInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryMinutesBasicInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryMinutesBasicInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMinutesBasicInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMinutesBasicInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMinutesBasicInfoRequest struct {
	// This parameter is required.
	TaskUuid *string `json:"taskUuid,omitempty" xml:"taskUuid,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryMinutesBasicInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryMinutesBasicInfoRequest) SetTaskUuid(v string) *QueryMinutesBasicInfoRequest {
	s.TaskUuid = &v
	return s
}

func (s *QueryMinutesBasicInfoRequest) SetUnionId(v string) *QueryMinutesBasicInfoRequest {
	s.UnionId = &v
	return s
}

type QueryMinutesBasicInfoResponseBody struct {
	Duration    *int64  `json:"duration,omitempty" xml:"duration,omitempty"`
	EndTime     *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime   *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	TaskCreator *string `json:"taskCreator,omitempty" xml:"taskCreator,omitempty"`
	TaskUuid    *string `json:"taskUuid,omitempty" xml:"taskUuid,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
	Url         *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s QueryMinutesBasicInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMinutesBasicInfoResponseBody) SetDuration(v int64) *QueryMinutesBasicInfoResponseBody {
	s.Duration = &v
	return s
}

func (s *QueryMinutesBasicInfoResponseBody) SetEndTime(v int64) *QueryMinutesBasicInfoResponseBody {
	s.EndTime = &v
	return s
}

func (s *QueryMinutesBasicInfoResponseBody) SetStartTime(v int64) *QueryMinutesBasicInfoResponseBody {
	s.StartTime = &v
	return s
}

func (s *QueryMinutesBasicInfoResponseBody) SetTaskCreator(v string) *QueryMinutesBasicInfoResponseBody {
	s.TaskCreator = &v
	return s
}

func (s *QueryMinutesBasicInfoResponseBody) SetTaskUuid(v string) *QueryMinutesBasicInfoResponseBody {
	s.TaskUuid = &v
	return s
}

func (s *QueryMinutesBasicInfoResponseBody) SetTitle(v string) *QueryMinutesBasicInfoResponseBody {
	s.Title = &v
	return s
}

func (s *QueryMinutesBasicInfoResponseBody) SetUrl(v string) *QueryMinutesBasicInfoResponseBody {
	s.Url = &v
	return s
}

type QueryMinutesBasicInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMinutesBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMinutesBasicInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryMinutesBasicInfoResponse) SetHeaders(v map[string]*string) *QueryMinutesBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryMinutesBasicInfoResponse) SetStatusCode(v int32) *QueryMinutesBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMinutesBasicInfoResponse) SetBody(v *QueryMinutesBasicInfoResponseBody) *QueryMinutesBasicInfoResponse {
	s.Body = v
	return s
}

type QueryMinutesChaptersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMinutesChaptersHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesChaptersHeaders) GoString() string {
	return s.String()
}

func (s *QueryMinutesChaptersHeaders) SetCommonHeaders(v map[string]*string) *QueryMinutesChaptersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMinutesChaptersHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMinutesChaptersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMinutesChaptersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// D5xxxxxxxxxgiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryMinutesChaptersRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesChaptersRequest) GoString() string {
	return s.String()
}

func (s *QueryMinutesChaptersRequest) SetUnionId(v string) *QueryMinutesChaptersRequest {
	s.UnionId = &v
	return s
}

type QueryMinutesChaptersResponseBody struct {
	ChapterList []*QueryMinutesChaptersResponseBodyChapterList `json:"chapterList,omitempty" xml:"chapterList,omitempty" type:"Repeated"`
	Status      *int32                                         `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryMinutesChaptersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesChaptersResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMinutesChaptersResponseBody) SetChapterList(v []*QueryMinutesChaptersResponseBodyChapterList) *QueryMinutesChaptersResponseBody {
	s.ChapterList = v
	return s
}

func (s *QueryMinutesChaptersResponseBody) SetStatus(v int32) *QueryMinutesChaptersResponseBody {
	s.Status = &v
	return s
}

type QueryMinutesChaptersResponseBodyChapterList struct {
	Content   *string `json:"content,omitempty" xml:"content,omitempty"`
	EndTime   *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	Uuid      *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s QueryMinutesChaptersResponseBodyChapterList) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesChaptersResponseBodyChapterList) GoString() string {
	return s.String()
}

func (s *QueryMinutesChaptersResponseBodyChapterList) SetContent(v string) *QueryMinutesChaptersResponseBodyChapterList {
	s.Content = &v
	return s
}

func (s *QueryMinutesChaptersResponseBodyChapterList) SetEndTime(v int64) *QueryMinutesChaptersResponseBodyChapterList {
	s.EndTime = &v
	return s
}

func (s *QueryMinutesChaptersResponseBodyChapterList) SetStartTime(v int64) *QueryMinutesChaptersResponseBodyChapterList {
	s.StartTime = &v
	return s
}

func (s *QueryMinutesChaptersResponseBodyChapterList) SetTitle(v string) *QueryMinutesChaptersResponseBodyChapterList {
	s.Title = &v
	return s
}

func (s *QueryMinutesChaptersResponseBodyChapterList) SetUuid(v string) *QueryMinutesChaptersResponseBodyChapterList {
	s.Uuid = &v
	return s
}

type QueryMinutesChaptersResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMinutesChaptersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMinutesChaptersResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesChaptersResponse) GoString() string {
	return s.String()
}

func (s *QueryMinutesChaptersResponse) SetHeaders(v map[string]*string) *QueryMinutesChaptersResponse {
	s.Headers = v
	return s
}

func (s *QueryMinutesChaptersResponse) SetStatusCode(v int32) *QueryMinutesChaptersResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMinutesChaptersResponse) SetBody(v *QueryMinutesChaptersResponseBody) *QueryMinutesChaptersResponse {
	s.Body = v
	return s
}

type QueryMinutesKeywordHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMinutesKeywordHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesKeywordHeaders) GoString() string {
	return s.String()
}

func (s *QueryMinutesKeywordHeaders) SetCommonHeaders(v map[string]*string) *QueryMinutesKeywordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMinutesKeywordHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMinutesKeywordHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMinutesKeywordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// D5xxxxxxxxxgiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryMinutesKeywordRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesKeywordRequest) GoString() string {
	return s.String()
}

func (s *QueryMinutesKeywordRequest) SetUnionId(v string) *QueryMinutesKeywordRequest {
	s.UnionId = &v
	return s
}

type QueryMinutesKeywordResponseBody struct {
	Keywords []*string `json:"keywords,omitempty" xml:"keywords,omitempty" type:"Repeated"`
}

func (s QueryMinutesKeywordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesKeywordResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMinutesKeywordResponseBody) SetKeywords(v []*string) *QueryMinutesKeywordResponseBody {
	s.Keywords = v
	return s
}

type QueryMinutesKeywordResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMinutesKeywordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMinutesKeywordResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesKeywordResponse) GoString() string {
	return s.String()
}

func (s *QueryMinutesKeywordResponse) SetHeaders(v map[string]*string) *QueryMinutesKeywordResponse {
	s.Headers = v
	return s
}

func (s *QueryMinutesKeywordResponse) SetStatusCode(v int32) *QueryMinutesKeywordResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMinutesKeywordResponse) SetBody(v *QueryMinutesKeywordResponseBody) *QueryMinutesKeywordResponse {
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
	EndTime        *int64                                                   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	NickName       *string                                                  `json:"nickName,omitempty" xml:"nickName,omitempty"`
	Paragraph      *string                                                  `json:"paragraph,omitempty" xml:"paragraph,omitempty"`
	ParagraphId    *int64                                                   `json:"paragraphId,omitempty" xml:"paragraphId,omitempty"`
	RecordId       *int64                                                   `json:"recordId,omitempty" xml:"recordId,omitempty"`
	SentenceList   []*QueryMinutesTextResponseBodyParagraphListSentenceList `json:"sentenceList,omitempty" xml:"sentenceList,omitempty" type:"Repeated"`
	SpeakerDisplay *QueryMinutesTextResponseBodyParagraphListSpeakerDisplay `json:"speakerDisplay,omitempty" xml:"speakerDisplay,omitempty" type:"Struct"`
	StartTime      *int64                                                   `json:"startTime,omitempty" xml:"startTime,omitempty"`
	SubSpeakerId   *string                                                  `json:"subSpeakerId,omitempty" xml:"subSpeakerId,omitempty"`
	UnionId        *string                                                  `json:"unionId,omitempty" xml:"unionId,omitempty"`
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

func (s *QueryMinutesTextResponseBodyParagraphList) SetSpeakerDisplay(v *QueryMinutesTextResponseBodyParagraphListSpeakerDisplay) *QueryMinutesTextResponseBodyParagraphList {
	s.SpeakerDisplay = v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetStartTime(v int64) *QueryMinutesTextResponseBodyParagraphList {
	s.StartTime = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetSubSpeakerId(v string) *QueryMinutesTextResponseBodyParagraphList {
	s.SubSpeakerId = &v
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

type QueryMinutesTextResponseBodyParagraphListSpeakerDisplay struct {
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	NickName  *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
}

func (s QueryMinutesTextResponseBodyParagraphListSpeakerDisplay) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesTextResponseBodyParagraphListSpeakerDisplay) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextResponseBodyParagraphListSpeakerDisplay) SetAvatarUrl(v string) *QueryMinutesTextResponseBodyParagraphListSpeakerDisplay {
	s.AvatarUrl = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphListSpeakerDisplay) SetNickName(v string) *QueryMinutesTextResponseBodyParagraphListSpeakerDisplay {
	s.NickName = &v
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

type QueryMinutesTodoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMinutesTodoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesTodoHeaders) GoString() string {
	return s.String()
}

func (s *QueryMinutesTodoHeaders) SetCommonHeaders(v map[string]*string) *QueryMinutesTodoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMinutesTodoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMinutesTodoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMinutesTodoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// D5xxxxxxxxxgiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryMinutesTodoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesTodoRequest) GoString() string {
	return s.String()
}

func (s *QueryMinutesTodoRequest) SetUnionId(v string) *QueryMinutesTodoRequest {
	s.UnionId = &v
	return s
}

type QueryMinutesTodoResponseBody struct {
	Actions          []*string                                       `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	DingtalkTodoList []*QueryMinutesTodoResponseBodyDingtalkTodoList `json:"dingtalkTodoList,omitempty" xml:"dingtalkTodoList,omitempty" type:"Repeated"`
}

func (s QueryMinutesTodoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesTodoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMinutesTodoResponseBody) SetActions(v []*string) *QueryMinutesTodoResponseBody {
	s.Actions = v
	return s
}

func (s *QueryMinutesTodoResponseBody) SetDingtalkTodoList(v []*QueryMinutesTodoResponseBodyDingtalkTodoList) *QueryMinutesTodoResponseBody {
	s.DingtalkTodoList = v
	return s
}

type QueryMinutesTodoResponseBodyDingtalkTodoList struct {
	CreatedTime    *int64                                                      `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	CreatorUnionId *string                                                     `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	Deadline       *string                                                     `json:"deadline,omitempty" xml:"deadline,omitempty"`
	DingtalkTodoId *string                                                     `json:"dingtalkTodoId,omitempty" xml:"dingtalkTodoId,omitempty"`
	ExecutorList   []*QueryMinutesTodoResponseBodyDingtalkTodoListExecutorList `json:"executorList,omitempty" xml:"executorList,omitempty" type:"Repeated"`
	IsDone         *bool                                                       `json:"isDone,omitempty" xml:"isDone,omitempty"`
	MinutesTodoId  *string                                                     `json:"minutesTodoId,omitempty" xml:"minutesTodoId,omitempty"`
	Title          *string                                                     `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryMinutesTodoResponseBodyDingtalkTodoList) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesTodoResponseBodyDingtalkTodoList) GoString() string {
	return s.String()
}

func (s *QueryMinutesTodoResponseBodyDingtalkTodoList) SetCreatedTime(v int64) *QueryMinutesTodoResponseBodyDingtalkTodoList {
	s.CreatedTime = &v
	return s
}

func (s *QueryMinutesTodoResponseBodyDingtalkTodoList) SetCreatorUnionId(v string) *QueryMinutesTodoResponseBodyDingtalkTodoList {
	s.CreatorUnionId = &v
	return s
}

func (s *QueryMinutesTodoResponseBodyDingtalkTodoList) SetDeadline(v string) *QueryMinutesTodoResponseBodyDingtalkTodoList {
	s.Deadline = &v
	return s
}

func (s *QueryMinutesTodoResponseBodyDingtalkTodoList) SetDingtalkTodoId(v string) *QueryMinutesTodoResponseBodyDingtalkTodoList {
	s.DingtalkTodoId = &v
	return s
}

func (s *QueryMinutesTodoResponseBodyDingtalkTodoList) SetExecutorList(v []*QueryMinutesTodoResponseBodyDingtalkTodoListExecutorList) *QueryMinutesTodoResponseBodyDingtalkTodoList {
	s.ExecutorList = v
	return s
}

func (s *QueryMinutesTodoResponseBodyDingtalkTodoList) SetIsDone(v bool) *QueryMinutesTodoResponseBodyDingtalkTodoList {
	s.IsDone = &v
	return s
}

func (s *QueryMinutesTodoResponseBodyDingtalkTodoList) SetMinutesTodoId(v string) *QueryMinutesTodoResponseBodyDingtalkTodoList {
	s.MinutesTodoId = &v
	return s
}

func (s *QueryMinutesTodoResponseBodyDingtalkTodoList) SetTitle(v string) *QueryMinutesTodoResponseBodyDingtalkTodoList {
	s.Title = &v
	return s
}

type QueryMinutesTodoResponseBodyDingtalkTodoListExecutorList struct {
	Avatar  *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	Nick    *string `json:"nick,omitempty" xml:"nick,omitempty"`
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryMinutesTodoResponseBodyDingtalkTodoListExecutorList) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesTodoResponseBodyDingtalkTodoListExecutorList) GoString() string {
	return s.String()
}

func (s *QueryMinutesTodoResponseBodyDingtalkTodoListExecutorList) SetAvatar(v string) *QueryMinutesTodoResponseBodyDingtalkTodoListExecutorList {
	s.Avatar = &v
	return s
}

func (s *QueryMinutesTodoResponseBodyDingtalkTodoListExecutorList) SetNick(v string) *QueryMinutesTodoResponseBodyDingtalkTodoListExecutorList {
	s.Nick = &v
	return s
}

func (s *QueryMinutesTodoResponseBodyDingtalkTodoListExecutorList) SetUnionId(v string) *QueryMinutesTodoResponseBodyDingtalkTodoListExecutorList {
	s.UnionId = &v
	return s
}

type QueryMinutesTodoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMinutesTodoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMinutesTodoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesTodoResponse) GoString() string {
	return s.String()
}

func (s *QueryMinutesTodoResponse) SetHeaders(v map[string]*string) *QueryMinutesTodoResponse {
	s.Headers = v
	return s
}

func (s *QueryMinutesTodoResponse) SetStatusCode(v int32) *QueryMinutesTodoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMinutesTodoResponse) SetBody(v *QueryMinutesTodoResponseBody) *QueryMinutesTodoResponse {
	s.Body = v
	return s
}

type QueryOrgDiyTemplatesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOrgDiyTemplatesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgDiyTemplatesHeaders) GoString() string {
	return s.String()
}

func (s *QueryOrgDiyTemplatesHeaders) SetCommonHeaders(v map[string]*string) *QueryOrgDiyTemplatesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOrgDiyTemplatesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOrgDiyTemplatesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOrgDiyTemplatesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 0
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lJcRnm39OsU4jlFVmRGXXXXX
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryOrgDiyTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgDiyTemplatesRequest) GoString() string {
	return s.String()
}

func (s *QueryOrgDiyTemplatesRequest) SetMaxResults(v int32) *QueryOrgDiyTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryOrgDiyTemplatesRequest) SetNextToken(v string) *QueryOrgDiyTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *QueryOrgDiyTemplatesRequest) SetUnionId(v string) *QueryOrgDiyTemplatesRequest {
	s.UnionId = &v
	return s
}

type QueryOrgDiyTemplatesResponseBody struct {
	DiyTemplates []*QueryOrgDiyTemplatesResponseBodyDiyTemplates `json:"diyTemplates,omitempty" xml:"diyTemplates,omitempty" type:"Repeated"`
	HasNext      *bool                                           `json:"hasNext,omitempty" xml:"hasNext,omitempty"`
	NextToken    *string                                         `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Total        *int64                                          `json:"total,omitempty" xml:"total,omitempty"`
}

func (s QueryOrgDiyTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgDiyTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrgDiyTemplatesResponseBody) SetDiyTemplates(v []*QueryOrgDiyTemplatesResponseBodyDiyTemplates) *QueryOrgDiyTemplatesResponseBody {
	s.DiyTemplates = v
	return s
}

func (s *QueryOrgDiyTemplatesResponseBody) SetHasNext(v bool) *QueryOrgDiyTemplatesResponseBody {
	s.HasNext = &v
	return s
}

func (s *QueryOrgDiyTemplatesResponseBody) SetNextToken(v string) *QueryOrgDiyTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryOrgDiyTemplatesResponseBody) SetTotal(v int64) *QueryOrgDiyTemplatesResponseBody {
	s.Total = &v
	return s
}

type QueryOrgDiyTemplatesResponseBodyDiyTemplates struct {
	AcceptTimes              *int64  `json:"acceptTimes,omitempty" xml:"acceptTimes,omitempty"`
	CreateTime               *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorNick              *string `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	CreatorUnionId           *string `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	DiyTemplateBriefIntro    *string `json:"diyTemplateBriefIntro,omitempty" xml:"diyTemplateBriefIntro,omitempty"`
	DiyTemplateIconUrl       *string `json:"diyTemplateIconUrl,omitempty" xml:"diyTemplateIconUrl,omitempty"`
	DiyTemplateKey           *string `json:"diyTemplateKey,omitempty" xml:"diyTemplateKey,omitempty"`
	DiyTemplateLatestVersion *int64  `json:"diyTemplateLatestVersion,omitempty" xml:"diyTemplateLatestVersion,omitempty"`
	DiyTemplateSource        *string `json:"diyTemplateSource,omitempty" xml:"diyTemplateSource,omitempty"`
	DiyTemplateTitle         *string `json:"diyTemplateTitle,omitempty" xml:"diyTemplateTitle,omitempty"`
	Status                   *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryOrgDiyTemplatesResponseBodyDiyTemplates) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgDiyTemplatesResponseBodyDiyTemplates) GoString() string {
	return s.String()
}

func (s *QueryOrgDiyTemplatesResponseBodyDiyTemplates) SetAcceptTimes(v int64) *QueryOrgDiyTemplatesResponseBodyDiyTemplates {
	s.AcceptTimes = &v
	return s
}

func (s *QueryOrgDiyTemplatesResponseBodyDiyTemplates) SetCreateTime(v int64) *QueryOrgDiyTemplatesResponseBodyDiyTemplates {
	s.CreateTime = &v
	return s
}

func (s *QueryOrgDiyTemplatesResponseBodyDiyTemplates) SetCreatorNick(v string) *QueryOrgDiyTemplatesResponseBodyDiyTemplates {
	s.CreatorNick = &v
	return s
}

func (s *QueryOrgDiyTemplatesResponseBodyDiyTemplates) SetCreatorUnionId(v string) *QueryOrgDiyTemplatesResponseBodyDiyTemplates {
	s.CreatorUnionId = &v
	return s
}

func (s *QueryOrgDiyTemplatesResponseBodyDiyTemplates) SetDiyTemplateBriefIntro(v string) *QueryOrgDiyTemplatesResponseBodyDiyTemplates {
	s.DiyTemplateBriefIntro = &v
	return s
}

func (s *QueryOrgDiyTemplatesResponseBodyDiyTemplates) SetDiyTemplateIconUrl(v string) *QueryOrgDiyTemplatesResponseBodyDiyTemplates {
	s.DiyTemplateIconUrl = &v
	return s
}

func (s *QueryOrgDiyTemplatesResponseBodyDiyTemplates) SetDiyTemplateKey(v string) *QueryOrgDiyTemplatesResponseBodyDiyTemplates {
	s.DiyTemplateKey = &v
	return s
}

func (s *QueryOrgDiyTemplatesResponseBodyDiyTemplates) SetDiyTemplateLatestVersion(v int64) *QueryOrgDiyTemplatesResponseBodyDiyTemplates {
	s.DiyTemplateLatestVersion = &v
	return s
}

func (s *QueryOrgDiyTemplatesResponseBodyDiyTemplates) SetDiyTemplateSource(v string) *QueryOrgDiyTemplatesResponseBodyDiyTemplates {
	s.DiyTemplateSource = &v
	return s
}

func (s *QueryOrgDiyTemplatesResponseBodyDiyTemplates) SetDiyTemplateTitle(v string) *QueryOrgDiyTemplatesResponseBodyDiyTemplates {
	s.DiyTemplateTitle = &v
	return s
}

func (s *QueryOrgDiyTemplatesResponseBodyDiyTemplates) SetStatus(v string) *QueryOrgDiyTemplatesResponseBodyDiyTemplates {
	s.Status = &v
	return s
}

type QueryOrgDiyTemplatesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOrgDiyTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOrgDiyTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgDiyTemplatesResponse) GoString() string {
	return s.String()
}

func (s *QueryOrgDiyTemplatesResponse) SetHeaders(v map[string]*string) *QueryOrgDiyTemplatesResponse {
	s.Headers = v
	return s
}

func (s *QueryOrgDiyTemplatesResponse) SetStatusCode(v int32) *QueryOrgDiyTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrgDiyTemplatesResponse) SetBody(v *QueryOrgDiyTemplatesResponseBody) *QueryOrgDiyTemplatesResponse {
	s.Body = v
	return s
}

type QueryScheduleConfMinutesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryScheduleConfMinutesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryScheduleConfMinutesHeaders) GoString() string {
	return s.String()
}

func (s *QueryScheduleConfMinutesHeaders) SetCommonHeaders(v map[string]*string) *QueryScheduleConfMinutesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryScheduleConfMinutesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryScheduleConfMinutesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryScheduleConfMinutesRequest struct {
	// This parameter is required.
	EventId *string `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryScheduleConfMinutesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryScheduleConfMinutesRequest) GoString() string {
	return s.String()
}

func (s *QueryScheduleConfMinutesRequest) SetEventId(v string) *QueryScheduleConfMinutesRequest {
	s.EventId = &v
	return s
}

func (s *QueryScheduleConfMinutesRequest) SetUnionId(v string) *QueryScheduleConfMinutesRequest {
	s.UnionId = &v
	return s
}

type QueryScheduleConfMinutesResponseBody struct {
	MinutesDetails []*QueryScheduleConfMinutesResponseBodyMinutesDetails `json:"minutesDetails,omitempty" xml:"minutesDetails,omitempty" type:"Repeated"`
}

func (s QueryScheduleConfMinutesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryScheduleConfMinutesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryScheduleConfMinutesResponseBody) SetMinutesDetails(v []*QueryScheduleConfMinutesResponseBodyMinutesDetails) *QueryScheduleConfMinutesResponseBody {
	s.MinutesDetails = v
	return s
}

type QueryScheduleConfMinutesResponseBodyMinutesDetails struct {
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

func (s QueryScheduleConfMinutesResponseBodyMinutesDetails) String() string {
	return tea.Prettify(s)
}

func (s QueryScheduleConfMinutesResponseBodyMinutesDetails) GoString() string {
	return s.String()
}

func (s *QueryScheduleConfMinutesResponseBodyMinutesDetails) SetBizType(v int32) *QueryScheduleConfMinutesResponseBodyMinutesDetails {
	s.BizType = &v
	return s
}

func (s *QueryScheduleConfMinutesResponseBodyMinutesDetails) SetCreatorNick(v string) *QueryScheduleConfMinutesResponseBodyMinutesDetails {
	s.CreatorNick = &v
	return s
}

func (s *QueryScheduleConfMinutesResponseBodyMinutesDetails) SetCreatorUnionId(v string) *QueryScheduleConfMinutesResponseBodyMinutesDetails {
	s.CreatorUnionId = &v
	return s
}

func (s *QueryScheduleConfMinutesResponseBodyMinutesDetails) SetDurationMicros(v int64) *QueryScheduleConfMinutesResponseBodyMinutesDetails {
	s.DurationMicros = &v
	return s
}

func (s *QueryScheduleConfMinutesResponseBodyMinutesDetails) SetIsDeleted(v int32) *QueryScheduleConfMinutesResponseBodyMinutesDetails {
	s.IsDeleted = &v
	return s
}

func (s *QueryScheduleConfMinutesResponseBodyMinutesDetails) SetSize(v int64) *QueryScheduleConfMinutesResponseBodyMinutesDetails {
	s.Size = &v
	return s
}

func (s *QueryScheduleConfMinutesResponseBodyMinutesDetails) SetStartTime(v int64) *QueryScheduleConfMinutesResponseBodyMinutesDetails {
	s.StartTime = &v
	return s
}

func (s *QueryScheduleConfMinutesResponseBodyMinutesDetails) SetStatus(v int32) *QueryScheduleConfMinutesResponseBodyMinutesDetails {
	s.Status = &v
	return s
}

func (s *QueryScheduleConfMinutesResponseBodyMinutesDetails) SetTaskUuid(v string) *QueryScheduleConfMinutesResponseBodyMinutesDetails {
	s.TaskUuid = &v
	return s
}

func (s *QueryScheduleConfMinutesResponseBodyMinutesDetails) SetTitle(v string) *QueryScheduleConfMinutesResponseBodyMinutesDetails {
	s.Title = &v
	return s
}

type QueryScheduleConfMinutesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryScheduleConfMinutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryScheduleConfMinutesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryScheduleConfMinutesResponse) GoString() string {
	return s.String()
}

func (s *QueryScheduleConfMinutesResponse) SetHeaders(v map[string]*string) *QueryScheduleConfMinutesResponse {
	s.Headers = v
	return s
}

func (s *QueryScheduleConfMinutesResponse) SetStatusCode(v int32) *QueryScheduleConfMinutesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryScheduleConfMinutesResponse) SetBody(v *QueryScheduleConfMinutesResponseBody) *QueryScheduleConfMinutesResponse {
	s.Body = v
	return s
}

type QuerySmartDeviceAiSummaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QuerySmartDeviceAiSummaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s QuerySmartDeviceAiSummaryHeaders) GoString() string {
	return s.String()
}

func (s *QuerySmartDeviceAiSummaryHeaders) SetCommonHeaders(v map[string]*string) *QuerySmartDeviceAiSummaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QuerySmartDeviceAiSummaryHeaders) SetXAcsDingtalkAccessToken(v string) *QuerySmartDeviceAiSummaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QuerySmartDeviceAiSummaryRequest struct {
	AgentId    *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	OpenFileId *string `json:"openFileId,omitempty" xml:"openFileId,omitempty"`
}

func (s QuerySmartDeviceAiSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySmartDeviceAiSummaryRequest) GoString() string {
	return s.String()
}

func (s *QuerySmartDeviceAiSummaryRequest) SetAgentId(v string) *QuerySmartDeviceAiSummaryRequest {
	s.AgentId = &v
	return s
}

func (s *QuerySmartDeviceAiSummaryRequest) SetOpenFileId(v string) *QuerySmartDeviceAiSummaryRequest {
	s.OpenFileId = &v
	return s
}

type QuerySmartDeviceAiSummaryResponseBody struct {
	AiSummaryList []*QuerySmartDeviceAiSummaryResponseBodyAiSummaryList `json:"aiSummaryList,omitempty" xml:"aiSummaryList,omitempty" type:"Repeated"`
	State         *int32                                                `json:"state,omitempty" xml:"state,omitempty"`
}

func (s QuerySmartDeviceAiSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySmartDeviceAiSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySmartDeviceAiSummaryResponseBody) SetAiSummaryList(v []*QuerySmartDeviceAiSummaryResponseBodyAiSummaryList) *QuerySmartDeviceAiSummaryResponseBody {
	s.AiSummaryList = v
	return s
}

func (s *QuerySmartDeviceAiSummaryResponseBody) SetState(v int32) *QuerySmartDeviceAiSummaryResponseBody {
	s.State = &v
	return s
}

type QuerySmartDeviceAiSummaryResponseBodyAiSummaryList struct {
	AgentId              *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	AiSceneRuleAvatarUrl *string `json:"aiSceneRuleAvatarUrl,omitempty" xml:"aiSceneRuleAvatarUrl,omitempty"`
	CreatorUnionId       *string `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	InstanceId           *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	OpenFileId           *string `json:"openFileId,omitempty" xml:"openFileId,omitempty"`
	Order                *int32  `json:"order,omitempty" xml:"order,omitempty"`
	Summary              *string `json:"summary,omitempty" xml:"summary,omitempty"`
	TemplateId           *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	Title                *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QuerySmartDeviceAiSummaryResponseBodyAiSummaryList) String() string {
	return tea.Prettify(s)
}

func (s QuerySmartDeviceAiSummaryResponseBodyAiSummaryList) GoString() string {
	return s.String()
}

func (s *QuerySmartDeviceAiSummaryResponseBodyAiSummaryList) SetAgentId(v string) *QuerySmartDeviceAiSummaryResponseBodyAiSummaryList {
	s.AgentId = &v
	return s
}

func (s *QuerySmartDeviceAiSummaryResponseBodyAiSummaryList) SetAiSceneRuleAvatarUrl(v string) *QuerySmartDeviceAiSummaryResponseBodyAiSummaryList {
	s.AiSceneRuleAvatarUrl = &v
	return s
}

func (s *QuerySmartDeviceAiSummaryResponseBodyAiSummaryList) SetCreatorUnionId(v string) *QuerySmartDeviceAiSummaryResponseBodyAiSummaryList {
	s.CreatorUnionId = &v
	return s
}

func (s *QuerySmartDeviceAiSummaryResponseBodyAiSummaryList) SetInstanceId(v string) *QuerySmartDeviceAiSummaryResponseBodyAiSummaryList {
	s.InstanceId = &v
	return s
}

func (s *QuerySmartDeviceAiSummaryResponseBodyAiSummaryList) SetOpenFileId(v string) *QuerySmartDeviceAiSummaryResponseBodyAiSummaryList {
	s.OpenFileId = &v
	return s
}

func (s *QuerySmartDeviceAiSummaryResponseBodyAiSummaryList) SetOrder(v int32) *QuerySmartDeviceAiSummaryResponseBodyAiSummaryList {
	s.Order = &v
	return s
}

func (s *QuerySmartDeviceAiSummaryResponseBodyAiSummaryList) SetSummary(v string) *QuerySmartDeviceAiSummaryResponseBodyAiSummaryList {
	s.Summary = &v
	return s
}

func (s *QuerySmartDeviceAiSummaryResponseBodyAiSummaryList) SetTemplateId(v string) *QuerySmartDeviceAiSummaryResponseBodyAiSummaryList {
	s.TemplateId = &v
	return s
}

func (s *QuerySmartDeviceAiSummaryResponseBodyAiSummaryList) SetTitle(v string) *QuerySmartDeviceAiSummaryResponseBodyAiSummaryList {
	s.Title = &v
	return s
}

type QuerySmartDeviceAiSummaryResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySmartDeviceAiSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySmartDeviceAiSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySmartDeviceAiSummaryResponse) GoString() string {
	return s.String()
}

func (s *QuerySmartDeviceAiSummaryResponse) SetHeaders(v map[string]*string) *QuerySmartDeviceAiSummaryResponse {
	s.Headers = v
	return s
}

func (s *QuerySmartDeviceAiSummaryResponse) SetStatusCode(v int32) *QuerySmartDeviceAiSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySmartDeviceAiSummaryResponse) SetBody(v *QuerySmartDeviceAiSummaryResponseBody) *QuerySmartDeviceAiSummaryResponse {
	s.Body = v
	return s
}

type QuerySummaryWithTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QuerySummaryWithTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s QuerySummaryWithTemplateHeaders) GoString() string {
	return s.String()
}

func (s *QuerySummaryWithTemplateHeaders) SetCommonHeaders(v map[string]*string) *QuerySummaryWithTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QuerySummaryWithTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *QuerySummaryWithTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QuerySummaryWithTemplateRequest struct {
	// example:
	//
	// 1
	DiyTemplateVersion *string `json:"diyTemplateVersion,omitempty" xml:"diyTemplateVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// meeting-assistant
	SummaryTemplateId *string `json:"summaryTemplateId,omitempty" xml:"summaryTemplateId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	SummaryTemplateType *string `json:"summaryTemplateType,omitempty" xml:"summaryTemplateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lJcRnm39OsU4jlFVmRGXXXXX
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QuerySummaryWithTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySummaryWithTemplateRequest) GoString() string {
	return s.String()
}

func (s *QuerySummaryWithTemplateRequest) SetDiyTemplateVersion(v string) *QuerySummaryWithTemplateRequest {
	s.DiyTemplateVersion = &v
	return s
}

func (s *QuerySummaryWithTemplateRequest) SetSummaryTemplateId(v string) *QuerySummaryWithTemplateRequest {
	s.SummaryTemplateId = &v
	return s
}

func (s *QuerySummaryWithTemplateRequest) SetSummaryTemplateType(v string) *QuerySummaryWithTemplateRequest {
	s.SummaryTemplateType = &v
	return s
}

func (s *QuerySummaryWithTemplateRequest) SetUnionId(v string) *QuerySummaryWithTemplateRequest {
	s.UnionId = &v
	return s
}

type QuerySummaryWithTemplateResponseBody struct {
	SummaryText *string `json:"summaryText,omitempty" xml:"summaryText,omitempty"`
}

func (s QuerySummaryWithTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySummaryWithTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySummaryWithTemplateResponseBody) SetSummaryText(v string) *QuerySummaryWithTemplateResponseBody {
	s.SummaryText = &v
	return s
}

type QuerySummaryWithTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySummaryWithTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySummaryWithTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySummaryWithTemplateResponse) GoString() string {
	return s.String()
}

func (s *QuerySummaryWithTemplateResponse) SetHeaders(v map[string]*string) *QuerySummaryWithTemplateResponse {
	s.Headers = v
	return s
}

func (s *QuerySummaryWithTemplateResponse) SetStatusCode(v int32) *QuerySummaryWithTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySummaryWithTemplateResponse) SetBody(v *QuerySummaryWithTemplateResponseBody) *QuerySummaryWithTemplateResponse {
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

type QueryUserAvailableDiyTemplatesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserAvailableDiyTemplatesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserAvailableDiyTemplatesHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserAvailableDiyTemplatesHeaders) SetCommonHeaders(v map[string]*string) *QueryUserAvailableDiyTemplatesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserAvailableDiyTemplatesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserAvailableDiyTemplatesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserAvailableDiyTemplatesRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 0
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lJcRnm39OsU4jlFVmRGXXXXX
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryUserAvailableDiyTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserAvailableDiyTemplatesRequest) GoString() string {
	return s.String()
}

func (s *QueryUserAvailableDiyTemplatesRequest) SetMaxResults(v int32) *QueryUserAvailableDiyTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryUserAvailableDiyTemplatesRequest) SetNextToken(v string) *QueryUserAvailableDiyTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *QueryUserAvailableDiyTemplatesRequest) SetUnionId(v string) *QueryUserAvailableDiyTemplatesRequest {
	s.UnionId = &v
	return s
}

type QueryUserAvailableDiyTemplatesResponseBody struct {
	HasNext          *bool                                                         `json:"hasNext,omitempty" xml:"hasNext,omitempty"`
	NextToken        *string                                                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OrgDiyTemplates  []*QueryUserAvailableDiyTemplatesResponseBodyOrgDiyTemplates  `json:"orgDiyTemplates,omitempty" xml:"orgDiyTemplates,omitempty" type:"Repeated"`
	UserDiyTemplates []*QueryUserAvailableDiyTemplatesResponseBodyUserDiyTemplates `json:"userDiyTemplates,omitempty" xml:"userDiyTemplates,omitempty" type:"Repeated"`
}

func (s QueryUserAvailableDiyTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserAvailableDiyTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserAvailableDiyTemplatesResponseBody) SetHasNext(v bool) *QueryUserAvailableDiyTemplatesResponseBody {
	s.HasNext = &v
	return s
}

func (s *QueryUserAvailableDiyTemplatesResponseBody) SetNextToken(v string) *QueryUserAvailableDiyTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryUserAvailableDiyTemplatesResponseBody) SetOrgDiyTemplates(v []*QueryUserAvailableDiyTemplatesResponseBodyOrgDiyTemplates) *QueryUserAvailableDiyTemplatesResponseBody {
	s.OrgDiyTemplates = v
	return s
}

func (s *QueryUserAvailableDiyTemplatesResponseBody) SetUserDiyTemplates(v []*QueryUserAvailableDiyTemplatesResponseBodyUserDiyTemplates) *QueryUserAvailableDiyTemplatesResponseBody {
	s.UserDiyTemplates = v
	return s
}

type QueryUserAvailableDiyTemplatesResponseBodyOrgDiyTemplates struct {
	CreateTime               *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DiyTemplateBriefIntro    *string `json:"diyTemplateBriefIntro,omitempty" xml:"diyTemplateBriefIntro,omitempty"`
	DiyTemplateIconUrl       *string `json:"diyTemplateIconUrl,omitempty" xml:"diyTemplateIconUrl,omitempty"`
	DiyTemplateKey           *string `json:"diyTemplateKey,omitempty" xml:"diyTemplateKey,omitempty"`
	DiyTemplateLatestVersion *int64  `json:"diyTemplateLatestVersion,omitempty" xml:"diyTemplateLatestVersion,omitempty"`
	DiyTemplateSource        *string `json:"diyTemplateSource,omitempty" xml:"diyTemplateSource,omitempty"`
	DiyTemplateTitle         *string `json:"diyTemplateTitle,omitempty" xml:"diyTemplateTitle,omitempty"`
	ModifiedTime             *int64  `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
}

func (s QueryUserAvailableDiyTemplatesResponseBodyOrgDiyTemplates) String() string {
	return tea.Prettify(s)
}

func (s QueryUserAvailableDiyTemplatesResponseBodyOrgDiyTemplates) GoString() string {
	return s.String()
}

func (s *QueryUserAvailableDiyTemplatesResponseBodyOrgDiyTemplates) SetCreateTime(v int64) *QueryUserAvailableDiyTemplatesResponseBodyOrgDiyTemplates {
	s.CreateTime = &v
	return s
}

func (s *QueryUserAvailableDiyTemplatesResponseBodyOrgDiyTemplates) SetDiyTemplateBriefIntro(v string) *QueryUserAvailableDiyTemplatesResponseBodyOrgDiyTemplates {
	s.DiyTemplateBriefIntro = &v
	return s
}

func (s *QueryUserAvailableDiyTemplatesResponseBodyOrgDiyTemplates) SetDiyTemplateIconUrl(v string) *QueryUserAvailableDiyTemplatesResponseBodyOrgDiyTemplates {
	s.DiyTemplateIconUrl = &v
	return s
}

func (s *QueryUserAvailableDiyTemplatesResponseBodyOrgDiyTemplates) SetDiyTemplateKey(v string) *QueryUserAvailableDiyTemplatesResponseBodyOrgDiyTemplates {
	s.DiyTemplateKey = &v
	return s
}

func (s *QueryUserAvailableDiyTemplatesResponseBodyOrgDiyTemplates) SetDiyTemplateLatestVersion(v int64) *QueryUserAvailableDiyTemplatesResponseBodyOrgDiyTemplates {
	s.DiyTemplateLatestVersion = &v
	return s
}

func (s *QueryUserAvailableDiyTemplatesResponseBodyOrgDiyTemplates) SetDiyTemplateSource(v string) *QueryUserAvailableDiyTemplatesResponseBodyOrgDiyTemplates {
	s.DiyTemplateSource = &v
	return s
}

func (s *QueryUserAvailableDiyTemplatesResponseBodyOrgDiyTemplates) SetDiyTemplateTitle(v string) *QueryUserAvailableDiyTemplatesResponseBodyOrgDiyTemplates {
	s.DiyTemplateTitle = &v
	return s
}

func (s *QueryUserAvailableDiyTemplatesResponseBodyOrgDiyTemplates) SetModifiedTime(v int64) *QueryUserAvailableDiyTemplatesResponseBodyOrgDiyTemplates {
	s.ModifiedTime = &v
	return s
}

type QueryUserAvailableDiyTemplatesResponseBodyUserDiyTemplates struct {
	CreateTime               *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DiyTemplateBriefIntro    *string `json:"diyTemplateBriefIntro,omitempty" xml:"diyTemplateBriefIntro,omitempty"`
	DiyTemplateIconUrl       *string `json:"diyTemplateIconUrl,omitempty" xml:"diyTemplateIconUrl,omitempty"`
	DiyTemplateKey           *string `json:"diyTemplateKey,omitempty" xml:"diyTemplateKey,omitempty"`
	DiyTemplateLatestVersion *int64  `json:"diyTemplateLatestVersion,omitempty" xml:"diyTemplateLatestVersion,omitempty"`
	DiyTemplateSource        *string `json:"diyTemplateSource,omitempty" xml:"diyTemplateSource,omitempty"`
	DiyTemplateTitle         *string `json:"diyTemplateTitle,omitempty" xml:"diyTemplateTitle,omitempty"`
	ModifiedTime             *int64  `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
}

func (s QueryUserAvailableDiyTemplatesResponseBodyUserDiyTemplates) String() string {
	return tea.Prettify(s)
}

func (s QueryUserAvailableDiyTemplatesResponseBodyUserDiyTemplates) GoString() string {
	return s.String()
}

func (s *QueryUserAvailableDiyTemplatesResponseBodyUserDiyTemplates) SetCreateTime(v int64) *QueryUserAvailableDiyTemplatesResponseBodyUserDiyTemplates {
	s.CreateTime = &v
	return s
}

func (s *QueryUserAvailableDiyTemplatesResponseBodyUserDiyTemplates) SetDiyTemplateBriefIntro(v string) *QueryUserAvailableDiyTemplatesResponseBodyUserDiyTemplates {
	s.DiyTemplateBriefIntro = &v
	return s
}

func (s *QueryUserAvailableDiyTemplatesResponseBodyUserDiyTemplates) SetDiyTemplateIconUrl(v string) *QueryUserAvailableDiyTemplatesResponseBodyUserDiyTemplates {
	s.DiyTemplateIconUrl = &v
	return s
}

func (s *QueryUserAvailableDiyTemplatesResponseBodyUserDiyTemplates) SetDiyTemplateKey(v string) *QueryUserAvailableDiyTemplatesResponseBodyUserDiyTemplates {
	s.DiyTemplateKey = &v
	return s
}

func (s *QueryUserAvailableDiyTemplatesResponseBodyUserDiyTemplates) SetDiyTemplateLatestVersion(v int64) *QueryUserAvailableDiyTemplatesResponseBodyUserDiyTemplates {
	s.DiyTemplateLatestVersion = &v
	return s
}

func (s *QueryUserAvailableDiyTemplatesResponseBodyUserDiyTemplates) SetDiyTemplateSource(v string) *QueryUserAvailableDiyTemplatesResponseBodyUserDiyTemplates {
	s.DiyTemplateSource = &v
	return s
}

func (s *QueryUserAvailableDiyTemplatesResponseBodyUserDiyTemplates) SetDiyTemplateTitle(v string) *QueryUserAvailableDiyTemplatesResponseBodyUserDiyTemplates {
	s.DiyTemplateTitle = &v
	return s
}

func (s *QueryUserAvailableDiyTemplatesResponseBodyUserDiyTemplates) SetModifiedTime(v int64) *QueryUserAvailableDiyTemplatesResponseBodyUserDiyTemplates {
	s.ModifiedTime = &v
	return s
}

type QueryUserAvailableDiyTemplatesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserAvailableDiyTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserAvailableDiyTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserAvailableDiyTemplatesResponse) GoString() string {
	return s.String()
}

func (s *QueryUserAvailableDiyTemplatesResponse) SetHeaders(v map[string]*string) *QueryUserAvailableDiyTemplatesResponse {
	s.Headers = v
	return s
}

func (s *QueryUserAvailableDiyTemplatesResponse) SetStatusCode(v int32) *QueryUserAvailableDiyTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserAvailableDiyTemplatesResponse) SetBody(v *QueryUserAvailableDiyTemplatesResponseBody) *QueryUserAvailableDiyTemplatesResponse {
	s.Body = v
	return s
}

type QueryUserMinutesPermissionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserMinutesPermissionHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserMinutesPermissionHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserMinutesPermissionHeaders) SetCommonHeaders(v map[string]*string) *QueryUserMinutesPermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserMinutesPermissionHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserMinutesPermissionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserMinutesPermissionResponseBody struct {
	HasPermission *bool `json:"hasPermission,omitempty" xml:"hasPermission,omitempty"`
	// 角色类型：manager-管理员, owner-所有者, editor-可编辑, read_download-可查看/下载, read-仅查看, none-无权限
	RoleType *string `json:"roleType,omitempty" xml:"roleType,omitempty"`
}

func (s QueryUserMinutesPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserMinutesPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserMinutesPermissionResponseBody) SetHasPermission(v bool) *QueryUserMinutesPermissionResponseBody {
	s.HasPermission = &v
	return s
}

func (s *QueryUserMinutesPermissionResponseBody) SetRoleType(v string) *QueryUserMinutesPermissionResponseBody {
	s.RoleType = &v
	return s
}

type QueryUserMinutesPermissionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserMinutesPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserMinutesPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserMinutesPermissionResponse) GoString() string {
	return s.String()
}

func (s *QueryUserMinutesPermissionResponse) SetHeaders(v map[string]*string) *QueryUserMinutesPermissionResponse {
	s.Headers = v
	return s
}

func (s *QueryUserMinutesPermissionResponse) SetStatusCode(v int32) *QueryUserMinutesPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserMinutesPermissionResponse) SetBody(v *QueryUserMinutesPermissionResponseBody) *QueryUserMinutesPermissionResponse {
	s.Body = v
	return s
}

type RegenerateChaptersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RegenerateChaptersHeaders) String() string {
	return tea.Prettify(s)
}

func (s RegenerateChaptersHeaders) GoString() string {
	return s.String()
}

func (s *RegenerateChaptersHeaders) SetCommonHeaders(v map[string]*string) *RegenerateChaptersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RegenerateChaptersHeaders) SetXAcsDingtalkAccessToken(v string) *RegenerateChaptersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RegenerateChaptersRequest struct {
	CustomPrompt *string `json:"customPrompt,omitempty" xml:"customPrompt,omitempty"`
	// This parameter is required.
	TaskUuid *string `json:"taskUuid,omitempty" xml:"taskUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lJcRnm39OsU4jlFVmRGXXXXX
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s RegenerateChaptersRequest) String() string {
	return tea.Prettify(s)
}

func (s RegenerateChaptersRequest) GoString() string {
	return s.String()
}

func (s *RegenerateChaptersRequest) SetCustomPrompt(v string) *RegenerateChaptersRequest {
	s.CustomPrompt = &v
	return s
}

func (s *RegenerateChaptersRequest) SetTaskUuid(v string) *RegenerateChaptersRequest {
	s.TaskUuid = &v
	return s
}

func (s *RegenerateChaptersRequest) SetUnionId(v string) *RegenerateChaptersRequest {
	s.UnionId = &v
	return s
}

type RegenerateChaptersResponseBody struct {
	TaskUuid *string `json:"taskUuid,omitempty" xml:"taskUuid,omitempty"`
}

func (s RegenerateChaptersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegenerateChaptersResponseBody) GoString() string {
	return s.String()
}

func (s *RegenerateChaptersResponseBody) SetTaskUuid(v string) *RegenerateChaptersResponseBody {
	s.TaskUuid = &v
	return s
}

type RegenerateChaptersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegenerateChaptersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegenerateChaptersResponse) String() string {
	return tea.Prettify(s)
}

func (s RegenerateChaptersResponse) GoString() string {
	return s.String()
}

func (s *RegenerateChaptersResponse) SetHeaders(v map[string]*string) *RegenerateChaptersResponse {
	s.Headers = v
	return s
}

func (s *RegenerateChaptersResponse) SetStatusCode(v int32) *RegenerateChaptersResponse {
	s.StatusCode = &v
	return s
}

func (s *RegenerateChaptersResponse) SetBody(v *RegenerateChaptersResponseBody) *RegenerateChaptersResponse {
	s.Body = v
	return s
}

type SetDetailPageCustomTabHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetDetailPageCustomTabHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetDetailPageCustomTabHeaders) GoString() string {
	return s.String()
}

func (s *SetDetailPageCustomTabHeaders) SetCommonHeaders(v map[string]*string) *SetDetailPageCustomTabHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetDetailPageCustomTabHeaders) SetXAcsDingtalkAccessToken(v string) *SetDetailPageCustomTabHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetDetailPageCustomTabRequest struct {
	// This parameter is required.
	CustomTabList []*SetDetailPageCustomTabRequestCustomTabList `json:"customTabList,omitempty" xml:"customTabList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// lJcRnm39OsU4jlFVmRGXXXXX
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s SetDetailPageCustomTabRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDetailPageCustomTabRequest) GoString() string {
	return s.String()
}

func (s *SetDetailPageCustomTabRequest) SetCustomTabList(v []*SetDetailPageCustomTabRequestCustomTabList) *SetDetailPageCustomTabRequest {
	s.CustomTabList = v
	return s
}

func (s *SetDetailPageCustomTabRequest) SetUnionId(v string) *SetDetailPageCustomTabRequest {
	s.UnionId = &v
	return s
}

type SetDetailPageCustomTabRequestCustomTabList struct {
	// This parameter is required.
	//
	// example:
	//
	// analyze
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn_ZH
	DefaultLocale *string `json:"defaultLocale,omitempty" xml:"defaultLocale,omitempty"`
	// This parameter is required.
	NameI18nMap map[string]interface{} `json:"nameI18nMap,omitempty" xml:"nameI18nMap,omitempty"`
	// example:
	//
	// https://example.com/pc/tab
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/tab
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SetDetailPageCustomTabRequestCustomTabList) String() string {
	return tea.Prettify(s)
}

func (s SetDetailPageCustomTabRequestCustomTabList) GoString() string {
	return s.String()
}

func (s *SetDetailPageCustomTabRequestCustomTabList) SetBizType(v string) *SetDetailPageCustomTabRequestCustomTabList {
	s.BizType = &v
	return s
}

func (s *SetDetailPageCustomTabRequestCustomTabList) SetDefaultLocale(v string) *SetDetailPageCustomTabRequestCustomTabList {
	s.DefaultLocale = &v
	return s
}

func (s *SetDetailPageCustomTabRequestCustomTabList) SetNameI18nMap(v map[string]interface{}) *SetDetailPageCustomTabRequestCustomTabList {
	s.NameI18nMap = v
	return s
}

func (s *SetDetailPageCustomTabRequestCustomTabList) SetPcUrl(v string) *SetDetailPageCustomTabRequestCustomTabList {
	s.PcUrl = &v
	return s
}

func (s *SetDetailPageCustomTabRequestCustomTabList) SetUrl(v string) *SetDetailPageCustomTabRequestCustomTabList {
	s.Url = &v
	return s
}

type SetDetailPageCustomTabResponseBody struct {
	TaskUuid *string `json:"taskUuid,omitempty" xml:"taskUuid,omitempty"`
}

func (s SetDetailPageCustomTabResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDetailPageCustomTabResponseBody) GoString() string {
	return s.String()
}

func (s *SetDetailPageCustomTabResponseBody) SetTaskUuid(v string) *SetDetailPageCustomTabResponseBody {
	s.TaskUuid = &v
	return s
}

type SetDetailPageCustomTabResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDetailPageCustomTabResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDetailPageCustomTabResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDetailPageCustomTabResponse) GoString() string {
	return s.String()
}

func (s *SetDetailPageCustomTabResponse) SetHeaders(v map[string]*string) *SetDetailPageCustomTabResponse {
	s.Headers = v
	return s
}

func (s *SetDetailPageCustomTabResponse) SetStatusCode(v int32) *SetDetailPageCustomTabResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDetailPageCustomTabResponse) SetBody(v *SetDetailPageCustomTabResponseBody) *SetDetailPageCustomTabResponse {
	s.Body = v
	return s
}

type SetInProgressCustomTabsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetInProgressCustomTabsHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetInProgressCustomTabsHeaders) GoString() string {
	return s.String()
}

func (s *SetInProgressCustomTabsHeaders) SetCommonHeaders(v map[string]*string) *SetInProgressCustomTabsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetInProgressCustomTabsHeaders) SetXAcsDingtalkAccessToken(v string) *SetInProgressCustomTabsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetInProgressCustomTabsRequest struct {
	// This parameter is required.
	CustomTabList []*SetInProgressCustomTabsRequestCustomTabList `json:"customTabList,omitempty" xml:"customTabList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// lJcRnm39OsU4jlFVmRGXXXXX
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s SetInProgressCustomTabsRequest) String() string {
	return tea.Prettify(s)
}

func (s SetInProgressCustomTabsRequest) GoString() string {
	return s.String()
}

func (s *SetInProgressCustomTabsRequest) SetCustomTabList(v []*SetInProgressCustomTabsRequestCustomTabList) *SetInProgressCustomTabsRequest {
	s.CustomTabList = v
	return s
}

func (s *SetInProgressCustomTabsRequest) SetUnionId(v string) *SetInProgressCustomTabsRequest {
	s.UnionId = &v
	return s
}

type SetInProgressCustomTabsRequestCustomTabList struct {
	// This parameter is required.
	//
	// example:
	//
	// data_analysis
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh_CN
	DefaultLocale *string `json:"defaultLocale,omitempty" xml:"defaultLocale,omitempty"`
	// This parameter is required.
	NameI18nMap map[string]interface{} `json:"nameI18nMap,omitempty" xml:"nameI18nMap,omitempty"`
	// example:
	//
	// https://example.com/app/minutes/analysis_pc
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/app/minutes/analysis_h5
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SetInProgressCustomTabsRequestCustomTabList) String() string {
	return tea.Prettify(s)
}

func (s SetInProgressCustomTabsRequestCustomTabList) GoString() string {
	return s.String()
}

func (s *SetInProgressCustomTabsRequestCustomTabList) SetBizType(v string) *SetInProgressCustomTabsRequestCustomTabList {
	s.BizType = &v
	return s
}

func (s *SetInProgressCustomTabsRequestCustomTabList) SetDefaultLocale(v string) *SetInProgressCustomTabsRequestCustomTabList {
	s.DefaultLocale = &v
	return s
}

func (s *SetInProgressCustomTabsRequestCustomTabList) SetNameI18nMap(v map[string]interface{}) *SetInProgressCustomTabsRequestCustomTabList {
	s.NameI18nMap = v
	return s
}

func (s *SetInProgressCustomTabsRequestCustomTabList) SetPcUrl(v string) *SetInProgressCustomTabsRequestCustomTabList {
	s.PcUrl = &v
	return s
}

func (s *SetInProgressCustomTabsRequestCustomTabList) SetUrl(v string) *SetInProgressCustomTabsRequestCustomTabList {
	s.Url = &v
	return s
}

type SetInProgressCustomTabsResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SetInProgressCustomTabsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetInProgressCustomTabsResponseBody) GoString() string {
	return s.String()
}

func (s *SetInProgressCustomTabsResponseBody) SetSuccess(v bool) *SetInProgressCustomTabsResponseBody {
	s.Success = &v
	return s
}

type SetInProgressCustomTabsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetInProgressCustomTabsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetInProgressCustomTabsResponse) String() string {
	return tea.Prettify(s)
}

func (s SetInProgressCustomTabsResponse) GoString() string {
	return s.String()
}

func (s *SetInProgressCustomTabsResponse) SetHeaders(v map[string]*string) *SetInProgressCustomTabsResponse {
	s.Headers = v
	return s
}

func (s *SetInProgressCustomTabsResponse) SetStatusCode(v int32) *SetInProgressCustomTabsResponse {
	s.StatusCode = &v
	return s
}

func (s *SetInProgressCustomTabsResponse) SetBody(v *SetInProgressCustomTabsResponseBody) *SetInProgressCustomTabsResponse {
	s.Body = v
	return s
}

type SetMinutesTodosVisibleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetMinutesTodosVisibleHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetMinutesTodosVisibleHeaders) GoString() string {
	return s.String()
}

func (s *SetMinutesTodosVisibleHeaders) SetCommonHeaders(v map[string]*string) *SetMinutesTodosVisibleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetMinutesTodosVisibleHeaders) SetXAcsDingtalkAccessToken(v string) *SetMinutesTodosVisibleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetMinutesTodosVisibleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// false
	TodosVisible *bool `json:"todosVisible,omitempty" xml:"todosVisible,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lJcRnm39OsU4jlFVmRGXXXXX
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s SetMinutesTodosVisibleRequest) String() string {
	return tea.Prettify(s)
}

func (s SetMinutesTodosVisibleRequest) GoString() string {
	return s.String()
}

func (s *SetMinutesTodosVisibleRequest) SetTodosVisible(v bool) *SetMinutesTodosVisibleRequest {
	s.TodosVisible = &v
	return s
}

func (s *SetMinutesTodosVisibleRequest) SetUnionId(v string) *SetMinutesTodosVisibleRequest {
	s.UnionId = &v
	return s
}

type SetMinutesTodosVisibleResponseBody struct {
	TaskUuid     *string `json:"taskUuid,omitempty" xml:"taskUuid,omitempty"`
	TodosVisible *bool   `json:"todosVisible,omitempty" xml:"todosVisible,omitempty"`
}

func (s SetMinutesTodosVisibleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetMinutesTodosVisibleResponseBody) GoString() string {
	return s.String()
}

func (s *SetMinutesTodosVisibleResponseBody) SetTaskUuid(v string) *SetMinutesTodosVisibleResponseBody {
	s.TaskUuid = &v
	return s
}

func (s *SetMinutesTodosVisibleResponseBody) SetTodosVisible(v bool) *SetMinutesTodosVisibleResponseBody {
	s.TodosVisible = &v
	return s
}

type SetMinutesTodosVisibleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetMinutesTodosVisibleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetMinutesTodosVisibleResponse) String() string {
	return tea.Prettify(s)
}

func (s SetMinutesTodosVisibleResponse) GoString() string {
	return s.String()
}

func (s *SetMinutesTodosVisibleResponse) SetHeaders(v map[string]*string) *SetMinutesTodosVisibleResponse {
	s.Headers = v
	return s
}

func (s *SetMinutesTodosVisibleResponse) SetStatusCode(v int32) *SetMinutesTodosVisibleResponse {
	s.StatusCode = &v
	return s
}

func (s *SetMinutesTodosVisibleResponse) SetBody(v *SetMinutesTodosVisibleResponseBody) *SetMinutesTodosVisibleResponse {
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

type UpdatePermissionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdatePermissionHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdatePermissionHeaders) GoString() string {
	return s.String()
}

func (s *UpdatePermissionHeaders) SetCommonHeaders(v map[string]*string) *UpdatePermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdatePermissionHeaders) SetXAcsDingtalkAccessToken(v string) *UpdatePermissionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdatePermissionRequest struct {
	MemberInfoList []*UpdatePermissionRequestMemberInfoList `json:"memberInfoList,omitempty" xml:"memberInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	OpType *int32 `json:"opType,omitempty" xml:"opType,omitempty"`
	// example:
	//
	// 1000
	RoleCode           *string   `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	RoleSubResourceIds []*string `json:"roleSubResourceIds,omitempty" xml:"roleSubResourceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	ShareScope *int32 `json:"shareScope,omitempty" xml:"shareScope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lJcRnm39OsU4jlFVmRGXXXXX
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s UpdatePermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePermissionRequest) GoString() string {
	return s.String()
}

func (s *UpdatePermissionRequest) SetMemberInfoList(v []*UpdatePermissionRequestMemberInfoList) *UpdatePermissionRequest {
	s.MemberInfoList = v
	return s
}

func (s *UpdatePermissionRequest) SetOpType(v int32) *UpdatePermissionRequest {
	s.OpType = &v
	return s
}

func (s *UpdatePermissionRequest) SetRoleCode(v string) *UpdatePermissionRequest {
	s.RoleCode = &v
	return s
}

func (s *UpdatePermissionRequest) SetRoleSubResourceIds(v []*string) *UpdatePermissionRequest {
	s.RoleSubResourceIds = v
	return s
}

func (s *UpdatePermissionRequest) SetShareScope(v int32) *UpdatePermissionRequest {
	s.ShareScope = &v
	return s
}

func (s *UpdatePermissionRequest) SetUnionId(v string) *UpdatePermissionRequest {
	s.UnionId = &v
	return s
}

type UpdatePermissionRequestMemberInfoList struct {
	// example:
	//
	// 2
	MemberType *int32 `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// example:
	//
	// lJcRnm39OsU4jlFVmRGXXXXX
	MemberUnionId *string `json:"memberUnionId,omitempty" xml:"memberUnionId,omitempty"`
	// example:
	//
	// 2
	PolicyId *int64 `json:"policyId,omitempty" xml:"policyId,omitempty"`
}

func (s UpdatePermissionRequestMemberInfoList) String() string {
	return tea.Prettify(s)
}

func (s UpdatePermissionRequestMemberInfoList) GoString() string {
	return s.String()
}

func (s *UpdatePermissionRequestMemberInfoList) SetMemberType(v int32) *UpdatePermissionRequestMemberInfoList {
	s.MemberType = &v
	return s
}

func (s *UpdatePermissionRequestMemberInfoList) SetMemberUnionId(v string) *UpdatePermissionRequestMemberInfoList {
	s.MemberUnionId = &v
	return s
}

func (s *UpdatePermissionRequestMemberInfoList) SetPolicyId(v int64) *UpdatePermissionRequestMemberInfoList {
	s.PolicyId = &v
	return s
}

type UpdatePermissionResponseBody struct {
	FailMemberInfoList []*UpdatePermissionResponseBodyFailMemberInfoList `json:"failMemberInfoList,omitempty" xml:"failMemberInfoList,omitempty" type:"Repeated"`
}

func (s UpdatePermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePermissionResponseBody) SetFailMemberInfoList(v []*UpdatePermissionResponseBodyFailMemberInfoList) *UpdatePermissionResponseBody {
	s.FailMemberInfoList = v
	return s
}

type UpdatePermissionResponseBodyFailMemberInfoList struct {
	// example:
	//
	// 2
	MemberType *int32 `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// example:
	//
	// lJcRnm39OsU4jlFVmRGXXXXX
	MemberUnionId *string `json:"memberUnionId,omitempty" xml:"memberUnionId,omitempty"`
	// example:
	//
	// 2
	PolicyId *int64 `json:"policyId,omitempty" xml:"policyId,omitempty"`
}

func (s UpdatePermissionResponseBodyFailMemberInfoList) String() string {
	return tea.Prettify(s)
}

func (s UpdatePermissionResponseBodyFailMemberInfoList) GoString() string {
	return s.String()
}

func (s *UpdatePermissionResponseBodyFailMemberInfoList) SetMemberType(v int32) *UpdatePermissionResponseBodyFailMemberInfoList {
	s.MemberType = &v
	return s
}

func (s *UpdatePermissionResponseBodyFailMemberInfoList) SetMemberUnionId(v string) *UpdatePermissionResponseBodyFailMemberInfoList {
	s.MemberUnionId = &v
	return s
}

func (s *UpdatePermissionResponseBodyFailMemberInfoList) SetPolicyId(v int64) *UpdatePermissionResponseBodyFailMemberInfoList {
	s.PolicyId = &v
	return s
}

type UpdatePermissionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePermissionResponse) GoString() string {
	return s.String()
}

func (s *UpdatePermissionResponse) SetHeaders(v map[string]*string) *UpdatePermissionResponse {
	s.Headers = v
	return s
}

func (s *UpdatePermissionResponse) SetStatusCode(v int32) *UpdatePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePermissionResponse) SetBody(v *UpdatePermissionResponseBody) *UpdatePermissionResponse {
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
// 搜索企业内听记
//
// @param request - AdminSearchMinutesRequest
//
// @param headers - AdminSearchMinutesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AdminSearchMinutesResponse
func (client *Client) AdminSearchMinutesWithOptions(request *AdminSearchMinutesRequest, headers *AdminSearchMinutesHeaders, runtime *util.RuntimeOptions) (_result *AdminSearchMinutesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreatorUnionIds)) {
		body["creatorUnionIds"] = request.CreatorUnionIds
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.SearchType)) {
		body["searchType"] = request.SearchType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
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
		Action:      tea.String("AdminSearchMinutes"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/adminSearch"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AdminSearchMinutesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索企业内听记
//
// @param request - AdminSearchMinutesRequest
//
// @return AdminSearchMinutesResponse
func (client *Client) AdminSearchMinutes(request *AdminSearchMinutesRequest) (_result *AdminSearchMinutesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AdminSearchMinutesHeaders{}
	_result = &AdminSearchMinutesResponse{}
	_body, _err := client.AdminSearchMinutesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
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
// 批量查询声纹配置信息
//
// @param request - BatchGetVoicePrintIdentifyConfigRequest
//
// @param headers - BatchGetVoicePrintIdentifyConfigHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetVoicePrintIdentifyConfigResponse
func (client *Client) BatchGetVoicePrintIdentifyConfigWithOptions(request *BatchGetVoicePrintIdentifyConfigRequest, headers *BatchGetVoicePrintIdentifyConfigHeaders, runtime *util.RuntimeOptions) (_result *BatchGetVoicePrintIdentifyConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Items)) {
		body["items"] = request.Items
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
		Action:      tea.String("BatchGetVoicePrintIdentifyConfig"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/voicePrint/configs/batchGet"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchGetVoicePrintIdentifyConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询声纹配置信息
//
// @param request - BatchGetVoicePrintIdentifyConfigRequest
//
// @return BatchGetVoicePrintIdentifyConfigResponse
func (client *Client) BatchGetVoicePrintIdentifyConfig(request *BatchGetVoicePrintIdentifyConfigRequest) (_result *BatchGetVoicePrintIdentifyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchGetVoicePrintIdentifyConfigHeaders{}
	_result = &BatchGetVoicePrintIdentifyConfigResponse{}
	_body, _err := client.BatchGetVoicePrintIdentifyConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量设置声纹开关状态
//
// @param request - BatchToggleVoicePrintSwitchStatusRequest
//
// @param headers - BatchToggleVoicePrintSwitchStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchToggleVoicePrintSwitchStatusResponse
func (client *Client) BatchToggleVoicePrintSwitchStatusWithOptions(request *BatchToggleVoicePrintSwitchStatusRequest, headers *BatchToggleVoicePrintSwitchStatusHeaders, runtime *util.RuntimeOptions) (_result *BatchToggleVoicePrintSwitchStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Items)) {
		body["items"] = request.Items
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
		Action:      tea.String("BatchToggleVoicePrintSwitchStatus"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/voicePrint/switches/batchToggle"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchToggleVoicePrintSwitchStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量设置声纹开关状态
//
// @param request - BatchToggleVoicePrintSwitchStatusRequest
//
// @return BatchToggleVoicePrintSwitchStatusResponse
func (client *Client) BatchToggleVoicePrintSwitchStatus(request *BatchToggleVoicePrintSwitchStatusRequest) (_result *BatchToggleVoicePrintSwitchStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchToggleVoicePrintSwitchStatusHeaders{}
	_result = &BatchToggleVoicePrintSwitchStatusResponse{}
	_body, _err := client.BatchToggleVoicePrintSwitchStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 从上传的媒体文件创建闪记
//
// @param request - CreateMinutesByUploadFileRequest
//
// @param headers - CreateMinutesByUploadFileHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMinutesByUploadFileResponse
func (client *Client) CreateMinutesByUploadFileWithOptions(request *CreateMinutesByUploadFileRequest, headers *CreateMinutesByUploadFileHeaders, runtime *util.RuntimeOptions) (_result *CreateMinutesByUploadFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorId)) {
		body["creatorId"] = request.CreatorId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomPrompt)) {
		body["customPrompt"] = request.CustomPrompt
	}

	if !tea.BoolValue(util.IsUnset(request.EnablePushCard)) {
		body["enablePushCard"] = request.EnablePushCard
	}

	if !tea.BoolValue(util.IsUnset(request.HiddenMinutes)) {
		body["hiddenMinutes"] = request.HiddenMinutes
	}

	if !tea.BoolValue(util.IsUnset(request.MediaFileUrl)) {
		body["mediaFileUrl"] = request.MediaFileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.MediaType)) {
		body["mediaType"] = request.MediaType
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
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
		Action:      tea.String("CreateMinutesByUploadFile"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/createMinutesByUploadFile"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMinutesByUploadFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 从上传的媒体文件创建闪记
//
// @param request - CreateMinutesByUploadFileRequest
//
// @return CreateMinutesByUploadFileResponse
func (client *Client) CreateMinutesByUploadFile(request *CreateMinutesByUploadFileRequest) (_result *CreateMinutesByUploadFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateMinutesByUploadFileHeaders{}
	_result = &CreateMinutesByUploadFileResponse{}
	_body, _err := client.CreateMinutesByUploadFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建DingTalkA1小助理分析
//
// @param request - CreateSmartDeviceAiSummaryRequest
//
// @param headers - CreateSmartDeviceAiSummaryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSmartDeviceAiSummaryResponse
func (client *Client) CreateSmartDeviceAiSummaryWithOptions(request *CreateSmartDeviceAiSummaryRequest, headers *CreateSmartDeviceAiSummaryHeaders, runtime *util.RuntimeOptions) (_result *CreateSmartDeviceAiSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		body["agentId"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsvContext)) {
		body["isvContext"] = request.IsvContext
	}

	if !tea.BoolValue(util.IsUnset(request.OpenFileId)) {
		body["openFileId"] = request.OpenFileId
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
		Action:      tea.String("CreateSmartDeviceAiSummary"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/smartdevice/aisummary/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSmartDeviceAiSummaryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建DingTalkA1小助理分析
//
// @param request - CreateSmartDeviceAiSummaryRequest
//
// @return CreateSmartDeviceAiSummaryResponse
func (client *Client) CreateSmartDeviceAiSummary(request *CreateSmartDeviceAiSummaryRequest) (_result *CreateSmartDeviceAiSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateSmartDeviceAiSummaryHeaders{}
	_result = &CreateSmartDeviceAiSummaryResponse{}
	_body, _err := client.CreateSmartDeviceAiSummaryWithOptions(request, headers, runtime)
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
// 导出闪记任务结果
//
// @param request - ExportMinutesTaskResultRequest
//
// @param headers - ExportMinutesTaskResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportMinutesTaskResultResponse
func (client *Client) ExportMinutesTaskResultWithOptions(request *ExportMinutesTaskResultRequest, headers *ExportMinutesTaskResultHeaders, runtime *util.RuntimeOptions) (_result *ExportMinutesTaskResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		body["expireTime"] = request.ExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.SummaryExportSetting)) {
		body["summaryExportSetting"] = request.SummaryExportSetting
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		body["taskType"] = request.TaskType
	}

	if !tea.BoolValue(util.IsUnset(request.TaskUuid)) {
		body["taskUuid"] = request.TaskUuid
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
		Action:      tea.String("ExportMinutesTaskResult"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/minutesTask/export"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportMinutesTaskResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出闪记任务结果
//
// @param request - ExportMinutesTaskResultRequest
//
// @return ExportMinutesTaskResultResponse
func (client *Client) ExportMinutesTaskResult(request *ExportMinutesTaskResultRequest) (_result *ExportMinutesTaskResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ExportMinutesTaskResultHeaders{}
	_result = &ExportMinutesTaskResultResponse{}
	_body, _err := client.ExportMinutesTaskResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生成摘要
//
// @param request - GenerateSummaryRequest
//
// @param headers - GenerateSummaryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateSummaryResponse
func (client *Client) GenerateSummaryWithOptions(taskUuid *string, request *GenerateSummaryRequest, headers *GenerateSummaryHeaders, runtime *util.RuntimeOptions) (_result *GenerateSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DiyTemplateVersion)) {
		body["diyTemplateVersion"] = request.DiyTemplateVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SummaryTemplateId)) {
		body["summaryTemplateId"] = request.SummaryTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.SummaryTemplateType)) {
		body["summaryTemplateType"] = request.SummaryTemplateType
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
		Action:      tea.String("GenerateSummary"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/tasks/" + tea.StringValue(taskUuid) + "/summary"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateSummaryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成摘要
//
// @param request - GenerateSummaryRequest
//
// @return GenerateSummaryResponse
func (client *Client) GenerateSummary(taskUuid *string, request *GenerateSummaryRequest) (_result *GenerateSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GenerateSummaryHeaders{}
	_result = &GenerateSummaryResponse{}
	_body, _err := client.GenerateSummaryWithOptions(taskUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询闪记摘要
//
// @param request - OpenQueryMinutesSummaryRequest
//
// @param headers - OpenQueryMinutesSummaryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenQueryMinutesSummaryResponse
func (client *Client) OpenQueryMinutesSummaryWithOptions(request *OpenQueryMinutesSummaryRequest, headers *OpenQueryMinutesSummaryHeaders, runtime *util.RuntimeOptions) (_result *OpenQueryMinutesSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskUuid)) {
		query["taskUuid"] = request.TaskUuid
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
		Action:      tea.String("OpenQueryMinutesSummary"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/queryMinutesSummary"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenQueryMinutesSummaryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询闪记摘要
//
// @param request - OpenQueryMinutesSummaryRequest
//
// @return OpenQueryMinutesSummaryResponse
func (client *Client) OpenQueryMinutesSummary(request *OpenQueryMinutesSummaryRequest) (_result *OpenQueryMinutesSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OpenQueryMinutesSummaryHeaders{}
	_result = &OpenQueryMinutesSummaryResponse{}
	_body, _err := client.OpenQueryMinutesSummaryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询闪会、文档等来源的闪记列表
//
// @param request - QueryBizMinutesRequest
//
// @param headers - QueryBizMinutesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBizMinutesResponse
func (client *Client) QueryBizMinutesWithOptions(request *QueryBizMinutesRequest, headers *QueryBizMinutesHeaders, runtime *util.RuntimeOptions) (_result *QueryBizMinutesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["bizType"] = request.BizType
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
		Action:      tea.String("QueryBizMinutes"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/queryBizMinutes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBizMinutesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询闪会、文档等来源的闪记列表
//
// @param request - QueryBizMinutesRequest
//
// @return QueryBizMinutesResponse
func (client *Client) QueryBizMinutes(request *QueryBizMinutesRequest) (_result *QueryBizMinutesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryBizMinutesHeaders{}
	_result = &QueryBizMinutesResponse{}
	_body, _err := client.QueryBizMinutesWithOptions(request, headers, runtime)
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
// 查询闪记基本信息
//
// @param request - QueryMinutesBasicInfoRequest
//
// @param headers - QueryMinutesBasicInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMinutesBasicInfoResponse
func (client *Client) QueryMinutesBasicInfoWithOptions(request *QueryMinutesBasicInfoRequest, headers *QueryMinutesBasicInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryMinutesBasicInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskUuid)) {
		query["taskUuid"] = request.TaskUuid
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
		Action:      tea.String("QueryMinutesBasicInfo"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/queryMinutesBasicInfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMinutesBasicInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询闪记基本信息
//
// @param request - QueryMinutesBasicInfoRequest
//
// @return QueryMinutesBasicInfoResponse
func (client *Client) QueryMinutesBasicInfo(request *QueryMinutesBasicInfoRequest) (_result *QueryMinutesBasicInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMinutesBasicInfoHeaders{}
	_result = &QueryMinutesBasicInfoResponse{}
	_body, _err := client.QueryMinutesBasicInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询听记智能章节列表
//
// @param request - QueryMinutesChaptersRequest
//
// @param headers - QueryMinutesChaptersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMinutesChaptersResponse
func (client *Client) QueryMinutesChaptersWithOptions(taskUuid *string, request *QueryMinutesChaptersRequest, headers *QueryMinutesChaptersHeaders, runtime *util.RuntimeOptions) (_result *QueryMinutesChaptersResponse, _err error) {
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
		Action:      tea.String("QueryMinutesChapters"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/" + tea.StringValue(taskUuid) + "/chapters"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMinutesChaptersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询听记智能章节列表
//
// @param request - QueryMinutesChaptersRequest
//
// @return QueryMinutesChaptersResponse
func (client *Client) QueryMinutesChapters(taskUuid *string, request *QueryMinutesChaptersRequest) (_result *QueryMinutesChaptersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMinutesChaptersHeaders{}
	_result = &QueryMinutesChaptersResponse{}
	_body, _err := client.QueryMinutesChaptersWithOptions(taskUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询闪记关键字
//
// @param request - QueryMinutesKeywordRequest
//
// @param headers - QueryMinutesKeywordHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMinutesKeywordResponse
func (client *Client) QueryMinutesKeywordWithOptions(taskUuid *string, request *QueryMinutesKeywordRequest, headers *QueryMinutesKeywordHeaders, runtime *util.RuntimeOptions) (_result *QueryMinutesKeywordResponse, _err error) {
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
		Action:      tea.String("QueryMinutesKeyword"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/" + tea.StringValue(taskUuid) + "/keywords"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMinutesKeywordResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询闪记关键字
//
// @param request - QueryMinutesKeywordRequest
//
// @return QueryMinutesKeywordResponse
func (client *Client) QueryMinutesKeyword(taskUuid *string, request *QueryMinutesKeywordRequest) (_result *QueryMinutesKeywordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMinutesKeywordHeaders{}
	_result = &QueryMinutesKeywordResponse{}
	_body, _err := client.QueryMinutesKeywordWithOptions(taskUuid, request, headers, runtime)
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
// 查询闪记待办
//
// @param request - QueryMinutesTodoRequest
//
// @param headers - QueryMinutesTodoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMinutesTodoResponse
func (client *Client) QueryMinutesTodoWithOptions(taskUuid *string, request *QueryMinutesTodoRequest, headers *QueryMinutesTodoHeaders, runtime *util.RuntimeOptions) (_result *QueryMinutesTodoResponse, _err error) {
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
		Action:      tea.String("QueryMinutesTodo"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/" + tea.StringValue(taskUuid) + "/todos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMinutesTodoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询闪记待办
//
// @param request - QueryMinutesTodoRequest
//
// @return QueryMinutesTodoResponse
func (client *Client) QueryMinutesTodo(taskUuid *string, request *QueryMinutesTodoRequest) (_result *QueryMinutesTodoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMinutesTodoHeaders{}
	_result = &QueryMinutesTodoResponse{}
	_body, _err := client.QueryMinutesTodoWithOptions(taskUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询企业所有自定义纪要模板列表
//
// @param request - QueryOrgDiyTemplatesRequest
//
// @param headers - QueryOrgDiyTemplatesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOrgDiyTemplatesResponse
func (client *Client) QueryOrgDiyTemplatesWithOptions(request *QueryOrgDiyTemplatesRequest, headers *QueryOrgDiyTemplatesHeaders, runtime *util.RuntimeOptions) (_result *QueryOrgDiyTemplatesResponse, _err error) {
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
		Action:      tea.String("QueryOrgDiyTemplates"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/diyTemplates/orgDeclared"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOrgDiyTemplatesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询企业所有自定义纪要模板列表
//
// @param request - QueryOrgDiyTemplatesRequest
//
// @return QueryOrgDiyTemplatesResponse
func (client *Client) QueryOrgDiyTemplates(request *QueryOrgDiyTemplatesRequest) (_result *QueryOrgDiyTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOrgDiyTemplatesHeaders{}
	_result = &QueryOrgDiyTemplatesResponse{}
	_body, _err := client.QueryOrgDiyTemplatesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询预约会议闪记列表
//
// @param request - QueryScheduleConfMinutesRequest
//
// @param headers - QueryScheduleConfMinutesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryScheduleConfMinutesResponse
func (client *Client) QueryScheduleConfMinutesWithOptions(scheduleConferenceId *string, request *QueryScheduleConfMinutesRequest, headers *QueryScheduleConfMinutesHeaders, runtime *util.RuntimeOptions) (_result *QueryScheduleConfMinutesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventId)) {
		query["eventId"] = request.EventId
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
		Action:      tea.String("QueryScheduleConfMinutes"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/scheduleConference/" + tea.StringValue(scheduleConferenceId) + "/minutes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryScheduleConfMinutesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询预约会议闪记列表
//
// @param request - QueryScheduleConfMinutesRequest
//
// @return QueryScheduleConfMinutesResponse
func (client *Client) QueryScheduleConfMinutes(scheduleConferenceId *string, request *QueryScheduleConfMinutesRequest) (_result *QueryScheduleConfMinutesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryScheduleConfMinutesHeaders{}
	_result = &QueryScheduleConfMinutesResponse{}
	_body, _err := client.QueryScheduleConfMinutesWithOptions(scheduleConferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询DingTalkA1小助理分析
//
// @param request - QuerySmartDeviceAiSummaryRequest
//
// @param headers - QuerySmartDeviceAiSummaryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmartDeviceAiSummaryResponse
func (client *Client) QuerySmartDeviceAiSummaryWithOptions(request *QuerySmartDeviceAiSummaryRequest, headers *QuerySmartDeviceAiSummaryHeaders, runtime *util.RuntimeOptions) (_result *QuerySmartDeviceAiSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		body["agentId"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenFileId)) {
		body["openFileId"] = request.OpenFileId
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
		Action:      tea.String("QuerySmartDeviceAiSummary"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/smartdevice/aisummary"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySmartDeviceAiSummaryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询DingTalkA1小助理分析
//
// @param request - QuerySmartDeviceAiSummaryRequest
//
// @return QuerySmartDeviceAiSummaryResponse
func (client *Client) QuerySmartDeviceAiSummary(request *QuerySmartDeviceAiSummaryRequest) (_result *QuerySmartDeviceAiSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QuerySmartDeviceAiSummaryHeaders{}
	_result = &QuerySmartDeviceAiSummaryResponse{}
	_body, _err := client.QuerySmartDeviceAiSummaryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据模板id查询摘要
//
// @param request - QuerySummaryWithTemplateRequest
//
// @param headers - QuerySummaryWithTemplateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySummaryWithTemplateResponse
func (client *Client) QuerySummaryWithTemplateWithOptions(taskUuid *string, request *QuerySummaryWithTemplateRequest, headers *QuerySummaryWithTemplateHeaders, runtime *util.RuntimeOptions) (_result *QuerySummaryWithTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DiyTemplateVersion)) {
		query["diyTemplateVersion"] = request.DiyTemplateVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SummaryTemplateId)) {
		query["summaryTemplateId"] = request.SummaryTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.SummaryTemplateType)) {
		query["summaryTemplateType"] = request.SummaryTemplateType
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
		Action:      tea.String("QuerySummaryWithTemplate"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/tasks/" + tea.StringValue(taskUuid) + "/summary/template"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySummaryWithTemplateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据模板id查询摘要
//
// @param request - QuerySummaryWithTemplateRequest
//
// @return QuerySummaryWithTemplateResponse
func (client *Client) QuerySummaryWithTemplate(taskUuid *string, request *QuerySummaryWithTemplateRequest) (_result *QuerySummaryWithTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QuerySummaryWithTemplateHeaders{}
	_result = &QuerySummaryWithTemplateResponse{}
	_body, _err := client.QuerySummaryWithTemplateWithOptions(taskUuid, request, headers, runtime)
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
// 查询用户可见的企业自定义纪要模版列表
//
// @param request - QueryUserAvailableDiyTemplatesRequest
//
// @param headers - QueryUserAvailableDiyTemplatesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserAvailableDiyTemplatesResponse
func (client *Client) QueryUserAvailableDiyTemplatesWithOptions(request *QueryUserAvailableDiyTemplatesRequest, headers *QueryUserAvailableDiyTemplatesHeaders, runtime *util.RuntimeOptions) (_result *QueryUserAvailableDiyTemplatesResponse, _err error) {
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
		Action:      tea.String("QueryUserAvailableDiyTemplates"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/diyTemplates/userAvailable"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserAvailableDiyTemplatesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户可见的企业自定义纪要模版列表
//
// @param request - QueryUserAvailableDiyTemplatesRequest
//
// @return QueryUserAvailableDiyTemplatesResponse
func (client *Client) QueryUserAvailableDiyTemplates(request *QueryUserAvailableDiyTemplatesRequest) (_result *QueryUserAvailableDiyTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserAvailableDiyTemplatesHeaders{}
	_result = &QueryUserAvailableDiyTemplatesResponse{}
	_body, _err := client.QueryUserAvailableDiyTemplatesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询指定用户对某篇听记的权限
//
// @param headers - QueryUserMinutesPermissionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserMinutesPermissionResponse
func (client *Client) QueryUserMinutesPermissionWithOptions(taskUuid *string, unionId *string, headers *QueryUserMinutesPermissionHeaders, runtime *util.RuntimeOptions) (_result *QueryUserMinutesPermissionResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("QueryUserMinutesPermission"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/tasks/" + tea.StringValue(taskUuid) + "/permissions/" + tea.StringValue(unionId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserMinutesPermissionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定用户对某篇听记的权限
//
// @return QueryUserMinutesPermissionResponse
func (client *Client) QueryUserMinutesPermission(taskUuid *string, unionId *string) (_result *QueryUserMinutesPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserMinutesPermissionHeaders{}
	_result = &QueryUserMinutesPermissionResponse{}
	_body, _err := client.QueryUserMinutesPermissionWithOptions(taskUuid, unionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重新生成听记智能章节
//
// @param request - RegenerateChaptersRequest
//
// @param headers - RegenerateChaptersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegenerateChaptersResponse
func (client *Client) RegenerateChaptersWithOptions(request *RegenerateChaptersRequest, headers *RegenerateChaptersHeaders, runtime *util.RuntimeOptions) (_result *RegenerateChaptersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskUuid)) {
		query["taskUuid"] = request.TaskUuid
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomPrompt)) {
		body["customPrompt"] = request.CustomPrompt
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
		Action:      tea.String("RegenerateChapters"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/chapters/regenerate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RegenerateChaptersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重新生成听记智能章节
//
// @param request - RegenerateChaptersRequest
//
// @return RegenerateChaptersResponse
func (client *Client) RegenerateChapters(request *RegenerateChaptersRequest) (_result *RegenerateChaptersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RegenerateChaptersHeaders{}
	_result = &RegenerateChaptersResponse{}
	_body, _err := client.RegenerateChaptersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 自定义听记详情页tab
//
// @param request - SetDetailPageCustomTabRequest
//
// @param headers - SetDetailPageCustomTabHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDetailPageCustomTabResponse
func (client *Client) SetDetailPageCustomTabWithOptions(taskUuid *string, request *SetDetailPageCustomTabRequest, headers *SetDetailPageCustomTabHeaders, runtime *util.RuntimeOptions) (_result *SetDetailPageCustomTabResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomTabList)) {
		body["customTabList"] = request.CustomTabList
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
		Action:      tea.String("SetDetailPageCustomTab"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/tasks/" + tea.StringValue(taskUuid) + "/customTabs"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDetailPageCustomTabResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 自定义听记详情页tab
//
// @param request - SetDetailPageCustomTabRequest
//
// @return SetDetailPageCustomTabResponse
func (client *Client) SetDetailPageCustomTab(taskUuid *string, request *SetDetailPageCustomTabRequest) (_result *SetDetailPageCustomTabResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetDetailPageCustomTabHeaders{}
	_result = &SetDetailPageCustomTabResponse{}
	_body, _err := client.SetDetailPageCustomTabWithOptions(taskUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 配置应用在听记录制页的自定义Tab
//
// @param request - SetInProgressCustomTabsRequest
//
// @param headers - SetInProgressCustomTabsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetInProgressCustomTabsResponse
func (client *Client) SetInProgressCustomTabsWithOptions(request *SetInProgressCustomTabsRequest, headers *SetInProgressCustomTabsHeaders, runtime *util.RuntimeOptions) (_result *SetInProgressCustomTabsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomTabList)) {
		body["customTabList"] = request.CustomTabList
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
		Action:      tea.String("SetInProgressCustomTabs"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/apps/settings/inProgressTabs"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SetInProgressCustomTabsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 配置应用在听记录制页的自定义Tab
//
// @param request - SetInProgressCustomTabsRequest
//
// @return SetInProgressCustomTabsResponse
func (client *Client) SetInProgressCustomTabs(request *SetInProgressCustomTabsRequest) (_result *SetInProgressCustomTabsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetInProgressCustomTabsHeaders{}
	_result = &SetInProgressCustomTabsResponse{}
	_body, _err := client.SetInProgressCustomTabsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置AI纪要待办模块可见性
//
// @param request - SetMinutesTodosVisibleRequest
//
// @param headers - SetMinutesTodosVisibleHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetMinutesTodosVisibleResponse
func (client *Client) SetMinutesTodosVisibleWithOptions(taskUuid *string, request *SetMinutesTodosVisibleRequest, headers *SetMinutesTodosVisibleHeaders, runtime *util.RuntimeOptions) (_result *SetMinutesTodosVisibleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TodosVisible)) {
		body["todosVisible"] = request.TodosVisible
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
		Action:      tea.String("SetMinutesTodosVisible"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/flashMinutes/tasks/" + tea.StringValue(taskUuid) + "/todosVisible"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SetMinutesTodosVisibleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置AI纪要待办模块可见性
//
// @param request - SetMinutesTodosVisibleRequest
//
// @return SetMinutesTodosVisibleResponse
func (client *Client) SetMinutesTodosVisible(taskUuid *string, request *SetMinutesTodosVisibleRequest) (_result *SetMinutesTodosVisibleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetMinutesTodosVisibleHeaders{}
	_result = &SetMinutesTodosVisibleResponse{}
	_body, _err := client.SetMinutesTodosVisibleWithOptions(taskUuid, request, headers, runtime)
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

// Summary:
//
// 更新闪记权限
//
// @param request - UpdatePermissionRequest
//
// @param headers - UpdatePermissionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePermissionResponse
func (client *Client) UpdatePermissionWithOptions(taskUuid *string, request *UpdatePermissionRequest, headers *UpdatePermissionHeaders, runtime *util.RuntimeOptions) (_result *UpdatePermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MemberInfoList)) {
		body["memberInfoList"] = request.MemberInfoList
	}

	if !tea.BoolValue(util.IsUnset(request.OpType)) {
		body["opType"] = request.OpType
	}

	if !tea.BoolValue(util.IsUnset(request.RoleCode)) {
		body["roleCode"] = request.RoleCode
	}

	if !tea.BoolValue(util.IsUnset(request.RoleSubResourceIds)) {
		body["roleSubResourceIds"] = request.RoleSubResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ShareScope)) {
		body["shareScope"] = request.ShareScope
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
		Action:      tea.String("UpdatePermission"),
		Version:     tea.String("minutes_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/minutes/" + tea.StringValue(taskUuid) + "/permission"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePermissionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新闪记权限
//
// @param request - UpdatePermissionRequest
//
// @return UpdatePermissionResponse
func (client *Client) UpdatePermission(taskUuid *string, request *UpdatePermissionRequest) (_result *UpdatePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdatePermissionHeaders{}
	_result = &UpdatePermissionResponse{}
	_body, _err := client.UpdatePermissionWithOptions(taskUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
