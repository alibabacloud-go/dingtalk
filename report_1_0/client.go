// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package report_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type UserMapValue struct {
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	DeptId *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s UserMapValue) String() string {
	return tea.Prettify(s)
}

func (s UserMapValue) GoString() string {
	return s.String()
}

func (s *UserMapValue) SetUserId(v string) *UserMapValue {
	s.UserId = &v
	return s
}

func (s *UserMapValue) SetName(v string) *UserMapValue {
	s.Name = &v
	return s
}

func (s *UserMapValue) SetDeptId(v string) *UserMapValue {
	s.DeptId = &v
	return s
}

type CreateTemplatesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateTemplatesHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplatesHeaders) GoString() string {
	return s.String()
}

func (s *CreateTemplatesHeaders) SetCommonHeaders(v map[string]*string) *CreateTemplatesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTemplatesHeaders) SetXAcsDingtalkAccessToken(v string) *CreateTemplatesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateTemplatesRequest struct {
	AllowAddReceivers           *bool                           `json:"allowAddReceivers,omitempty" xml:"allowAddReceivers,omitempty"`
	AllowEdit                   *bool                           `json:"allowEdit,omitempty" xml:"allowEdit,omitempty"`
	AllowGetLocation            *bool                           `json:"allowGetLocation,omitempty" xml:"allowGetLocation,omitempty"`
	AuthDeptIds                 []*string                       `json:"authDeptIds,omitempty" xml:"authDeptIds,omitempty" type:"Repeated"`
	AuthUserIds                 []*string                       `json:"authUserIds,omitempty" xml:"authUserIds,omitempty" type:"Repeated"`
	Creator                     *string                         `json:"creator,omitempty" xml:"creator,omitempty"`
	DefaultReceivedCids         []*string                       `json:"defaultReceivedCids,omitempty" xml:"defaultReceivedCids,omitempty" type:"Repeated"`
	DefaultReceivedMasterLevels []*string                       `json:"defaultReceivedMasterLevels,omitempty" xml:"defaultReceivedMasterLevels,omitempty" type:"Repeated"`
	DefaultReceivers            []*string                       `json:"defaultReceivers,omitempty" xml:"defaultReceivers,omitempty" type:"Repeated"`
	Fields                      []*CreateTemplatesRequestFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	Logo                        *string                         `json:"logo,omitempty" xml:"logo,omitempty"`
	MaxWordCount                *int32                          `json:"maxWordCount,omitempty" xml:"maxWordCount,omitempty"`
	MinWordCount                *int32                          `json:"minWordCount,omitempty" xml:"minWordCount,omitempty"`
	Name                        *string                         `json:"name,omitempty" xml:"name,omitempty"`
	TemplateManagers            []*string                       `json:"templateManagers,omitempty" xml:"templateManagers,omitempty" type:"Repeated"`
}

func (s CreateTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplatesRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplatesRequest) SetAllowAddReceivers(v bool) *CreateTemplatesRequest {
	s.AllowAddReceivers = &v
	return s
}

func (s *CreateTemplatesRequest) SetAllowEdit(v bool) *CreateTemplatesRequest {
	s.AllowEdit = &v
	return s
}

func (s *CreateTemplatesRequest) SetAllowGetLocation(v bool) *CreateTemplatesRequest {
	s.AllowGetLocation = &v
	return s
}

func (s *CreateTemplatesRequest) SetAuthDeptIds(v []*string) *CreateTemplatesRequest {
	s.AuthDeptIds = v
	return s
}

func (s *CreateTemplatesRequest) SetAuthUserIds(v []*string) *CreateTemplatesRequest {
	s.AuthUserIds = v
	return s
}

func (s *CreateTemplatesRequest) SetCreator(v string) *CreateTemplatesRequest {
	s.Creator = &v
	return s
}

func (s *CreateTemplatesRequest) SetDefaultReceivedCids(v []*string) *CreateTemplatesRequest {
	s.DefaultReceivedCids = v
	return s
}

func (s *CreateTemplatesRequest) SetDefaultReceivedMasterLevels(v []*string) *CreateTemplatesRequest {
	s.DefaultReceivedMasterLevels = v
	return s
}

func (s *CreateTemplatesRequest) SetDefaultReceivers(v []*string) *CreateTemplatesRequest {
	s.DefaultReceivers = v
	return s
}

func (s *CreateTemplatesRequest) SetFields(v []*CreateTemplatesRequestFields) *CreateTemplatesRequest {
	s.Fields = v
	return s
}

func (s *CreateTemplatesRequest) SetLogo(v string) *CreateTemplatesRequest {
	s.Logo = &v
	return s
}

func (s *CreateTemplatesRequest) SetMaxWordCount(v int32) *CreateTemplatesRequest {
	s.MaxWordCount = &v
	return s
}

func (s *CreateTemplatesRequest) SetMinWordCount(v int32) *CreateTemplatesRequest {
	s.MinWordCount = &v
	return s
}

func (s *CreateTemplatesRequest) SetName(v string) *CreateTemplatesRequest {
	s.Name = &v
	return s
}

func (s *CreateTemplatesRequest) SetTemplateManagers(v []*string) *CreateTemplatesRequest {
	s.TemplateManagers = v
	return s
}

type CreateTemplatesRequestFields struct {
	DataType  *int32                                 `json:"dataType,omitempty" xml:"dataType,omitempty"`
	DataValue *CreateTemplatesRequestFieldsDataValue `json:"dataValue,omitempty" xml:"dataValue,omitempty" type:"Struct"`
	FieldName *string                                `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	Need      *bool                                  `json:"need,omitempty" xml:"need,omitempty"`
	Order     *int32                                 `json:"order,omitempty" xml:"order,omitempty"`
	Sort      *int32                                 `json:"sort,omitempty" xml:"sort,omitempty"`
}

func (s CreateTemplatesRequestFields) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplatesRequestFields) GoString() string {
	return s.String()
}

func (s *CreateTemplatesRequestFields) SetDataType(v int32) *CreateTemplatesRequestFields {
	s.DataType = &v
	return s
}

func (s *CreateTemplatesRequestFields) SetDataValue(v *CreateTemplatesRequestFieldsDataValue) *CreateTemplatesRequestFields {
	s.DataValue = v
	return s
}

func (s *CreateTemplatesRequestFields) SetFieldName(v string) *CreateTemplatesRequestFields {
	s.FieldName = &v
	return s
}

func (s *CreateTemplatesRequestFields) SetNeed(v bool) *CreateTemplatesRequestFields {
	s.Need = &v
	return s
}

func (s *CreateTemplatesRequestFields) SetOrder(v int32) *CreateTemplatesRequestFields {
	s.Order = &v
	return s
}

func (s *CreateTemplatesRequestFields) SetSort(v int32) *CreateTemplatesRequestFields {
	s.Sort = &v
	return s
}

type CreateTemplatesRequestFieldsDataValue struct {
	OpenInfo    *CreateTemplatesRequestFieldsDataValueOpenInfo `json:"openInfo,omitempty" xml:"openInfo,omitempty" type:"Struct"`
	Options     []*string                                      `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	Placeholder *string                                        `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
}

func (s CreateTemplatesRequestFieldsDataValue) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplatesRequestFieldsDataValue) GoString() string {
	return s.String()
}

func (s *CreateTemplatesRequestFieldsDataValue) SetOpenInfo(v *CreateTemplatesRequestFieldsDataValueOpenInfo) *CreateTemplatesRequestFieldsDataValue {
	s.OpenInfo = v
	return s
}

func (s *CreateTemplatesRequestFieldsDataValue) SetOptions(v []*string) *CreateTemplatesRequestFieldsDataValue {
	s.Options = v
	return s
}

func (s *CreateTemplatesRequestFieldsDataValue) SetPlaceholder(v string) *CreateTemplatesRequestFieldsDataValue {
	s.Placeholder = &v
	return s
}

type CreateTemplatesRequestFieldsDataValueOpenInfo struct {
	Attribute map[string]*string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	OpenId    *string            `json:"openId,omitempty" xml:"openId,omitempty"`
}

func (s CreateTemplatesRequestFieldsDataValueOpenInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplatesRequestFieldsDataValueOpenInfo) GoString() string {
	return s.String()
}

func (s *CreateTemplatesRequestFieldsDataValueOpenInfo) SetAttribute(v map[string]*string) *CreateTemplatesRequestFieldsDataValueOpenInfo {
	s.Attribute = v
	return s
}

func (s *CreateTemplatesRequestFieldsDataValueOpenInfo) SetOpenId(v string) *CreateTemplatesRequestFieldsDataValueOpenInfo {
	s.OpenId = &v
	return s
}

type CreateTemplatesResponseBody struct {
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s CreateTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplatesResponseBody) SetTemplateId(v string) *CreateTemplatesResponseBody {
	s.TemplateId = &v
	return s
}

type CreateTemplatesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplatesResponse) GoString() string {
	return s.String()
}

func (s *CreateTemplatesResponse) SetHeaders(v map[string]*string) *CreateTemplatesResponse {
	s.Headers = v
	return s
}

func (s *CreateTemplatesResponse) SetStatusCode(v int32) *CreateTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTemplatesResponse) SetBody(v *CreateTemplatesResponseBody) *CreateTemplatesResponse {
	s.Body = v
	return s
}

type GetSendAndReceiveReportListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSendAndReceiveReportListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSendAndReceiveReportListHeaders) GoString() string {
	return s.String()
}

func (s *GetSendAndReceiveReportListHeaders) SetCommonHeaders(v map[string]*string) *GetSendAndReceiveReportListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSendAndReceiveReportListHeaders) SetXAcsDingtalkAccessToken(v string) *GetSendAndReceiveReportListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSendAndReceiveReportListRequest struct {
	EndTime         *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	MaxResults      *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken       *int64  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OperationUserId *string `json:"operationUserId,omitempty" xml:"operationUserId,omitempty"`
	StartTime       *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetSendAndReceiveReportListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSendAndReceiveReportListRequest) GoString() string {
	return s.String()
}

func (s *GetSendAndReceiveReportListRequest) SetEndTime(v int64) *GetSendAndReceiveReportListRequest {
	s.EndTime = &v
	return s
}

func (s *GetSendAndReceiveReportListRequest) SetMaxResults(v int32) *GetSendAndReceiveReportListRequest {
	s.MaxResults = &v
	return s
}

func (s *GetSendAndReceiveReportListRequest) SetNextToken(v int64) *GetSendAndReceiveReportListRequest {
	s.NextToken = &v
	return s
}

func (s *GetSendAndReceiveReportListRequest) SetOperationUserId(v string) *GetSendAndReceiveReportListRequest {
	s.OperationUserId = &v
	return s
}

func (s *GetSendAndReceiveReportListRequest) SetStartTime(v int64) *GetSendAndReceiveReportListRequest {
	s.StartTime = &v
	return s
}

type GetSendAndReceiveReportListResponseBody struct {
	DataList   []*GetSendAndReceiveReportListResponseBodyDataList `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	HasMore    *bool                                              `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	MaxResults *int32                                             `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int64                                             `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetSendAndReceiveReportListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSendAndReceiveReportListResponseBody) GoString() string {
	return s.String()
}

func (s *GetSendAndReceiveReportListResponseBody) SetDataList(v []*GetSendAndReceiveReportListResponseBodyDataList) *GetSendAndReceiveReportListResponseBody {
	s.DataList = v
	return s
}

func (s *GetSendAndReceiveReportListResponseBody) SetHasMore(v bool) *GetSendAndReceiveReportListResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetSendAndReceiveReportListResponseBody) SetMaxResults(v int32) *GetSendAndReceiveReportListResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetSendAndReceiveReportListResponseBody) SetNextToken(v int64) *GetSendAndReceiveReportListResponseBody {
	s.NextToken = &v
	return s
}

type GetSendAndReceiveReportListResponseBodyDataList struct {
	CreateTime   *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId    *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	CreatorName  *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	ModifiedTime *int64  `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ReportId     *string `json:"reportId,omitempty" xml:"reportId,omitempty"`
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s GetSendAndReceiveReportListResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetSendAndReceiveReportListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetSendAndReceiveReportListResponseBodyDataList) SetCreateTime(v int64) *GetSendAndReceiveReportListResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *GetSendAndReceiveReportListResponseBodyDataList) SetCreatorId(v string) *GetSendAndReceiveReportListResponseBodyDataList {
	s.CreatorId = &v
	return s
}

func (s *GetSendAndReceiveReportListResponseBodyDataList) SetCreatorName(v string) *GetSendAndReceiveReportListResponseBodyDataList {
	s.CreatorName = &v
	return s
}

func (s *GetSendAndReceiveReportListResponseBodyDataList) SetModifiedTime(v int64) *GetSendAndReceiveReportListResponseBodyDataList {
	s.ModifiedTime = &v
	return s
}

func (s *GetSendAndReceiveReportListResponseBodyDataList) SetReportId(v string) *GetSendAndReceiveReportListResponseBodyDataList {
	s.ReportId = &v
	return s
}

func (s *GetSendAndReceiveReportListResponseBodyDataList) SetTemplateName(v string) *GetSendAndReceiveReportListResponseBodyDataList {
	s.TemplateName = &v
	return s
}

type GetSendAndReceiveReportListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSendAndReceiveReportListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSendAndReceiveReportListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSendAndReceiveReportListResponse) GoString() string {
	return s.String()
}

func (s *GetSendAndReceiveReportListResponse) SetHeaders(v map[string]*string) *GetSendAndReceiveReportListResponse {
	s.Headers = v
	return s
}

func (s *GetSendAndReceiveReportListResponse) SetStatusCode(v int32) *GetSendAndReceiveReportListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSendAndReceiveReportListResponse) SetBody(v *GetSendAndReceiveReportListResponseBody) *GetSendAndReceiveReportListResponse {
	s.Body = v
	return s
}

type GetSubmitStatisticsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSubmitStatisticsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSubmitStatisticsHeaders) GoString() string {
	return s.String()
}

func (s *GetSubmitStatisticsHeaders) SetCommonHeaders(v map[string]*string) *GetSubmitStatisticsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSubmitStatisticsHeaders) SetXAcsDingtalkAccessToken(v string) *GetSubmitStatisticsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSubmitStatisticsRequest struct {
	EndTime         *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	OperationUserId *string `json:"operationUserId,omitempty" xml:"operationUserId,omitempty"`
	RemindId        *int64  `json:"remindId,omitempty" xml:"remindId,omitempty"`
	StartTime       *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	TemplateId      *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s GetSubmitStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSubmitStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetSubmitStatisticsRequest) SetEndTime(v int64) *GetSubmitStatisticsRequest {
	s.EndTime = &v
	return s
}

func (s *GetSubmitStatisticsRequest) SetOperationUserId(v string) *GetSubmitStatisticsRequest {
	s.OperationUserId = &v
	return s
}

func (s *GetSubmitStatisticsRequest) SetRemindId(v int64) *GetSubmitStatisticsRequest {
	s.RemindId = &v
	return s
}

func (s *GetSubmitStatisticsRequest) SetStartTime(v int64) *GetSubmitStatisticsRequest {
	s.StartTime = &v
	return s
}

func (s *GetSubmitStatisticsRequest) SetTemplateId(v string) *GetSubmitStatisticsRequest {
	s.TemplateId = &v
	return s
}

type GetSubmitStatisticsResponseBody struct {
	ShouldRemindTimes *int32                            `json:"shouldRemindTimes,omitempty" xml:"shouldRemindTimes,omitempty"`
	TemplateName      *string                           `json:"templateName,omitempty" xml:"templateName,omitempty"`
	UserDeptMap       map[string]*string                `json:"userDeptMap,omitempty" xml:"userDeptMap,omitempty"`
	UserIdCountMap    map[string]*int64                 `json:"userIdCountMap,omitempty" xml:"userIdCountMap,omitempty"`
	UserIdStatusMap   map[string]map[string]interface{} `json:"userIdStatusMap,omitempty" xml:"userIdStatusMap,omitempty"`
	UserIds           []*string                         `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
	UserMap           map[string]*UserMapValue          `json:"userMap,omitempty" xml:"userMap,omitempty"`
}

func (s GetSubmitStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSubmitStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubmitStatisticsResponseBody) SetShouldRemindTimes(v int32) *GetSubmitStatisticsResponseBody {
	s.ShouldRemindTimes = &v
	return s
}

func (s *GetSubmitStatisticsResponseBody) SetTemplateName(v string) *GetSubmitStatisticsResponseBody {
	s.TemplateName = &v
	return s
}

func (s *GetSubmitStatisticsResponseBody) SetUserDeptMap(v map[string]*string) *GetSubmitStatisticsResponseBody {
	s.UserDeptMap = v
	return s
}

func (s *GetSubmitStatisticsResponseBody) SetUserIdCountMap(v map[string]*int64) *GetSubmitStatisticsResponseBody {
	s.UserIdCountMap = v
	return s
}

func (s *GetSubmitStatisticsResponseBody) SetUserIdStatusMap(v map[string]map[string]interface{}) *GetSubmitStatisticsResponseBody {
	s.UserIdStatusMap = v
	return s
}

func (s *GetSubmitStatisticsResponseBody) SetUserIds(v []*string) *GetSubmitStatisticsResponseBody {
	s.UserIds = v
	return s
}

func (s *GetSubmitStatisticsResponseBody) SetUserMap(v map[string]*UserMapValue) *GetSubmitStatisticsResponseBody {
	s.UserMap = v
	return s
}

type GetSubmitStatisticsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSubmitStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSubmitStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSubmitStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetSubmitStatisticsResponse) SetHeaders(v map[string]*string) *GetSubmitStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetSubmitStatisticsResponse) SetStatusCode(v int32) *GetSubmitStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubmitStatisticsResponse) SetBody(v *GetSubmitStatisticsResponseBody) *GetSubmitStatisticsResponse {
	s.Body = v
	return s
}

type QueryRemindResultsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryRemindResultsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryRemindResultsHeaders) GoString() string {
	return s.String()
}

func (s *QueryRemindResultsHeaders) SetCommonHeaders(v map[string]*string) *QueryRemindResultsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryRemindResultsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryRemindResultsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryRemindResultsRequest struct {
	MaxResults      *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken       *int64  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OperationUserId *string `json:"operationUserId,omitempty" xml:"operationUserId,omitempty"`
	TemplateId      *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s QueryRemindResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRemindResultsRequest) GoString() string {
	return s.String()
}

func (s *QueryRemindResultsRequest) SetMaxResults(v int32) *QueryRemindResultsRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryRemindResultsRequest) SetNextToken(v int64) *QueryRemindResultsRequest {
	s.NextToken = &v
	return s
}

func (s *QueryRemindResultsRequest) SetOperationUserId(v string) *QueryRemindResultsRequest {
	s.OperationUserId = &v
	return s
}

func (s *QueryRemindResultsRequest) SetTemplateId(v string) *QueryRemindResultsRequest {
	s.TemplateId = &v
	return s
}

type QueryRemindResultsResponseBody struct {
	DataList  []*QueryRemindResultsResponseBodyDataList `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	HasMore   *bool                                     `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken *int64                                    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryRemindResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRemindResultsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRemindResultsResponseBody) SetDataList(v []*QueryRemindResultsResponseBodyDataList) *QueryRemindResultsResponseBody {
	s.DataList = v
	return s
}

func (s *QueryRemindResultsResponseBody) SetHasMore(v bool) *QueryRemindResultsResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryRemindResultsResponseBody) SetNextToken(v int64) *QueryRemindResultsResponseBody {
	s.NextToken = &v
	return s
}

type QueryRemindResultsResponseBodyDataList struct {
	CreatorId     *string                                           `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	EndDateTime   []*string                                         `json:"endDateTime,omitempty" xml:"endDateTime,omitempty" type:"Repeated"`
	ModifierId    *string                                           `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	PeriodType    *int32                                            `json:"periodType,omitempty" xml:"periodType,omitempty"`
	RemindId      *int64                                            `json:"remindId,omitempty" xml:"remindId,omitempty"`
	StartDateTime []*string                                         `json:"startDateTime,omitempty" xml:"startDateTime,omitempty" type:"Repeated"`
	TemplateId    *string                                           `json:"templateId,omitempty" xml:"templateId,omitempty"`
	ToGroups      []*QueryRemindResultsResponseBodyDataListToGroups `json:"toGroups,omitempty" xml:"toGroups,omitempty" type:"Repeated"`
}

func (s QueryRemindResultsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryRemindResultsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryRemindResultsResponseBodyDataList) SetCreatorId(v string) *QueryRemindResultsResponseBodyDataList {
	s.CreatorId = &v
	return s
}

func (s *QueryRemindResultsResponseBodyDataList) SetEndDateTime(v []*string) *QueryRemindResultsResponseBodyDataList {
	s.EndDateTime = v
	return s
}

func (s *QueryRemindResultsResponseBodyDataList) SetModifierId(v string) *QueryRemindResultsResponseBodyDataList {
	s.ModifierId = &v
	return s
}

func (s *QueryRemindResultsResponseBodyDataList) SetPeriodType(v int32) *QueryRemindResultsResponseBodyDataList {
	s.PeriodType = &v
	return s
}

func (s *QueryRemindResultsResponseBodyDataList) SetRemindId(v int64) *QueryRemindResultsResponseBodyDataList {
	s.RemindId = &v
	return s
}

func (s *QueryRemindResultsResponseBodyDataList) SetStartDateTime(v []*string) *QueryRemindResultsResponseBodyDataList {
	s.StartDateTime = v
	return s
}

func (s *QueryRemindResultsResponseBodyDataList) SetTemplateId(v string) *QueryRemindResultsResponseBodyDataList {
	s.TemplateId = &v
	return s
}

func (s *QueryRemindResultsResponseBodyDataList) SetToGroups(v []*QueryRemindResultsResponseBodyDataListToGroups) *QueryRemindResultsResponseBodyDataList {
	s.ToGroups = v
	return s
}

type QueryRemindResultsResponseBodyDataListToGroups struct {
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryRemindResultsResponseBodyDataListToGroups) String() string {
	return tea.Prettify(s)
}

func (s QueryRemindResultsResponseBodyDataListToGroups) GoString() string {
	return s.String()
}

func (s *QueryRemindResultsResponseBodyDataListToGroups) SetTitle(v string) *QueryRemindResultsResponseBodyDataListToGroups {
	s.Title = &v
	return s
}

type QueryRemindResultsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRemindResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRemindResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRemindResultsResponse) GoString() string {
	return s.String()
}

func (s *QueryRemindResultsResponse) SetHeaders(v map[string]*string) *QueryRemindResultsResponse {
	s.Headers = v
	return s
}

func (s *QueryRemindResultsResponse) SetStatusCode(v int32) *QueryRemindResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRemindResultsResponse) SetBody(v *QueryRemindResultsResponseBody) *QueryRemindResultsResponse {
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
	interfaceSPI, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = interfaceSPI
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

func (client *Client) CreateTemplatesWithOptions(request *CreateTemplatesRequest, headers *CreateTemplatesHeaders, runtime *util.RuntimeOptions) (_result *CreateTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowAddReceivers)) {
		body["allowAddReceivers"] = request.AllowAddReceivers
	}

	if !tea.BoolValue(util.IsUnset(request.AllowEdit)) {
		body["allowEdit"] = request.AllowEdit
	}

	if !tea.BoolValue(util.IsUnset(request.AllowGetLocation)) {
		body["allowGetLocation"] = request.AllowGetLocation
	}

	if !tea.BoolValue(util.IsUnset(request.AuthDeptIds)) {
		body["authDeptIds"] = request.AuthDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.AuthUserIds)) {
		body["authUserIds"] = request.AuthUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.Creator)) {
		body["creator"] = request.Creator
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultReceivedCids)) {
		body["defaultReceivedCids"] = request.DefaultReceivedCids
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultReceivedMasterLevels)) {
		body["defaultReceivedMasterLevels"] = request.DefaultReceivedMasterLevels
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultReceivers)) {
		body["defaultReceivers"] = request.DefaultReceivers
	}

	if !tea.BoolValue(util.IsUnset(request.Fields)) {
		body["fields"] = request.Fields
	}

	if !tea.BoolValue(util.IsUnset(request.Logo)) {
		body["logo"] = request.Logo
	}

	if !tea.BoolValue(util.IsUnset(request.MaxWordCount)) {
		body["maxWordCount"] = request.MaxWordCount
	}

	if !tea.BoolValue(util.IsUnset(request.MinWordCount)) {
		body["minWordCount"] = request.MinWordCount
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateManagers)) {
		body["templateManagers"] = request.TemplateManagers
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
		Action:      tea.String("CreateTemplates"),
		Version:     tea.String("report_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/report/templates"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTemplatesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTemplates(request *CreateTemplatesRequest) (_result *CreateTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTemplatesHeaders{}
	_result = &CreateTemplatesResponse{}
	_body, _err := client.CreateTemplatesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSendAndReceiveReportListWithOptions(request *GetSendAndReceiveReportListRequest, headers *GetSendAndReceiveReportListHeaders, runtime *util.RuntimeOptions) (_result *GetSendAndReceiveReportListResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.OperationUserId)) {
		query["operationUserId"] = request.OperationUserId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
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
		Action:      tea.String("GetSendAndReceiveReportList"),
		Version:     tea.String("report_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/report/users/sendAndReceiveLists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSendAndReceiveReportListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSendAndReceiveReportList(request *GetSendAndReceiveReportListRequest) (_result *GetSendAndReceiveReportListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSendAndReceiveReportListHeaders{}
	_result = &GetSendAndReceiveReportListResponse{}
	_body, _err := client.GetSendAndReceiveReportListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSubmitStatisticsWithOptions(request *GetSubmitStatisticsRequest, headers *GetSubmitStatisticsHeaders, runtime *util.RuntimeOptions) (_result *GetSubmitStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OperationUserId)) {
		query["operationUserId"] = request.OperationUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RemindId)) {
		query["remindId"] = request.RemindId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["templateId"] = request.TemplateId
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
		Action:      tea.String("GetSubmitStatistics"),
		Version:     tea.String("report_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/report/submitStatisticalResults"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSubmitStatisticsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSubmitStatistics(request *GetSubmitStatisticsRequest) (_result *GetSubmitStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSubmitStatisticsHeaders{}
	_result = &GetSubmitStatisticsResponse{}
	_body, _err := client.GetSubmitStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRemindResultsWithOptions(request *QueryRemindResultsRequest, headers *QueryRemindResultsHeaders, runtime *util.RuntimeOptions) (_result *QueryRemindResultsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.OperationUserId)) {
		query["operationUserId"] = request.OperationUserId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["templateId"] = request.TemplateId
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
		Action:      tea.String("QueryRemindResults"),
		Version:     tea.String("report_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/report/statisticalRules/results"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRemindResultsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRemindResults(request *QueryRemindResultsRequest) (_result *QueryRemindResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryRemindResultsHeaders{}
	_result = &QueryRemindResultsResponse{}
	_body, _err := client.QueryRemindResultsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
