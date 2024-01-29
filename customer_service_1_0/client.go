// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package customer_service_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketHeaders) GoString() string {
	return s.String()
}

func (s *CreateTicketHeaders) SetCommonHeaders(v map[string]*string) *CreateTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTicketHeaders) SetXAcsDingtalkAccessToken(v string) *CreateTicketHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateTicketRequest struct {
	ForeignId      *string                          `json:"foreignId,omitempty" xml:"foreignId,omitempty"`
	ForeignName    *string                          `json:"foreignName,omitempty" xml:"foreignName,omitempty"`
	OpenInstanceId *string                          `json:"openInstanceId,omitempty" xml:"openInstanceId,omitempty"`
	ProductionType *int32                           `json:"productionType,omitempty" xml:"productionType,omitempty"`
	Properties     []*CreateTicketRequestProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Repeated"`
	SourceId       *string                          `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	TemplateId     *string                          `json:"templateId,omitempty" xml:"templateId,omitempty"`
	Title          *string                          `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketRequest) GoString() string {
	return s.String()
}

func (s *CreateTicketRequest) SetForeignId(v string) *CreateTicketRequest {
	s.ForeignId = &v
	return s
}

func (s *CreateTicketRequest) SetForeignName(v string) *CreateTicketRequest {
	s.ForeignName = &v
	return s
}

func (s *CreateTicketRequest) SetOpenInstanceId(v string) *CreateTicketRequest {
	s.OpenInstanceId = &v
	return s
}

func (s *CreateTicketRequest) SetProductionType(v int32) *CreateTicketRequest {
	s.ProductionType = &v
	return s
}

func (s *CreateTicketRequest) SetProperties(v []*CreateTicketRequestProperties) *CreateTicketRequest {
	s.Properties = v
	return s
}

func (s *CreateTicketRequest) SetSourceId(v string) *CreateTicketRequest {
	s.SourceId = &v
	return s
}

func (s *CreateTicketRequest) SetTemplateId(v string) *CreateTicketRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateTicketRequest) SetTitle(v string) *CreateTicketRequest {
	s.Title = &v
	return s
}

type CreateTicketRequestProperties struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateTicketRequestProperties) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketRequestProperties) GoString() string {
	return s.String()
}

func (s *CreateTicketRequestProperties) SetName(v string) *CreateTicketRequestProperties {
	s.Name = &v
	return s
}

func (s *CreateTicketRequestProperties) SetValue(v string) *CreateTicketRequestProperties {
	s.Value = &v
	return s
}

type CreateTicketResponseBody struct {
	TicketId *string `json:"ticketId,omitempty" xml:"ticketId,omitempty"`
}

func (s CreateTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTicketResponseBody) SetTicketId(v string) *CreateTicketResponseBody {
	s.TicketId = &v
	return s
}

type CreateTicketResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketResponse) GoString() string {
	return s.String()
}

func (s *CreateTicketResponse) SetHeaders(v map[string]*string) *CreateTicketResponse {
	s.Headers = v
	return s
}

func (s *CreateTicketResponse) SetStatusCode(v int32) *CreateTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTicketResponse) SetBody(v *CreateTicketResponseBody) *CreateTicketResponse {
	s.Body = v
	return s
}

type ExecuteActivityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ExecuteActivityHeaders) String() string {
	return tea.Prettify(s)
}

func (s ExecuteActivityHeaders) GoString() string {
	return s.String()
}

func (s *ExecuteActivityHeaders) SetCommonHeaders(v map[string]*string) *ExecuteActivityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ExecuteActivityHeaders) SetXAcsDingtalkAccessToken(v string) *ExecuteActivityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ExecuteActivityRequest struct {
	ActivityCode   *string                             `json:"activityCode,omitempty" xml:"activityCode,omitempty"`
	ForeignId      *string                             `json:"foreignId,omitempty" xml:"foreignId,omitempty"`
	ForeignName    *string                             `json:"foreignName,omitempty" xml:"foreignName,omitempty"`
	OpenInstanceId *string                             `json:"openInstanceId,omitempty" xml:"openInstanceId,omitempty"`
	ProductionType *int32                              `json:"productionType,omitempty" xml:"productionType,omitempty"`
	Properties     []*ExecuteActivityRequestProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Repeated"`
	SourceId       *string                             `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
}

func (s ExecuteActivityRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteActivityRequest) GoString() string {
	return s.String()
}

func (s *ExecuteActivityRequest) SetActivityCode(v string) *ExecuteActivityRequest {
	s.ActivityCode = &v
	return s
}

func (s *ExecuteActivityRequest) SetForeignId(v string) *ExecuteActivityRequest {
	s.ForeignId = &v
	return s
}

func (s *ExecuteActivityRequest) SetForeignName(v string) *ExecuteActivityRequest {
	s.ForeignName = &v
	return s
}

func (s *ExecuteActivityRequest) SetOpenInstanceId(v string) *ExecuteActivityRequest {
	s.OpenInstanceId = &v
	return s
}

func (s *ExecuteActivityRequest) SetProductionType(v int32) *ExecuteActivityRequest {
	s.ProductionType = &v
	return s
}

func (s *ExecuteActivityRequest) SetProperties(v []*ExecuteActivityRequestProperties) *ExecuteActivityRequest {
	s.Properties = v
	return s
}

func (s *ExecuteActivityRequest) SetSourceId(v string) *ExecuteActivityRequest {
	s.SourceId = &v
	return s
}

type ExecuteActivityRequestProperties struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ExecuteActivityRequestProperties) String() string {
	return tea.Prettify(s)
}

func (s ExecuteActivityRequestProperties) GoString() string {
	return s.String()
}

func (s *ExecuteActivityRequestProperties) SetName(v string) *ExecuteActivityRequestProperties {
	s.Name = &v
	return s
}

func (s *ExecuteActivityRequestProperties) SetValue(v string) *ExecuteActivityRequestProperties {
	s.Value = &v
	return s
}

type ExecuteActivityResponseBody struct {
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ExecuteActivityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteActivityResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteActivityResponseBody) SetTaskId(v string) *ExecuteActivityResponseBody {
	s.TaskId = &v
	return s
}

type ExecuteActivityResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteActivityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteActivityResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteActivityResponse) GoString() string {
	return s.String()
}

func (s *ExecuteActivityResponse) SetHeaders(v map[string]*string) *ExecuteActivityResponse {
	s.Headers = v
	return s
}

func (s *ExecuteActivityResponse) SetStatusCode(v int32) *ExecuteActivityResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteActivityResponse) SetBody(v *ExecuteActivityResponseBody) *ExecuteActivityResponse {
	s.Body = v
	return s
}

type GetUserSourceListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserSourceListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserSourceListHeaders) GoString() string {
	return s.String()
}

func (s *GetUserSourceListHeaders) SetCommonHeaders(v map[string]*string) *GetUserSourceListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserSourceListHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserSourceListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserSourceListRequest struct {
	CorpId         *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	OpenInstanceId *string `json:"openInstanceId,omitempty" xml:"openInstanceId,omitempty"`
	OrgId          *int64  `json:"orgId,omitempty" xml:"orgId,omitempty"`
	OrgName        *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	ProductionType *int32  `json:"productionType,omitempty" xml:"productionType,omitempty"`
}

func (s GetUserSourceListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserSourceListRequest) GoString() string {
	return s.String()
}

func (s *GetUserSourceListRequest) SetCorpId(v string) *GetUserSourceListRequest {
	s.CorpId = &v
	return s
}

func (s *GetUserSourceListRequest) SetDescription(v string) *GetUserSourceListRequest {
	s.Description = &v
	return s
}

func (s *GetUserSourceListRequest) SetOpenInstanceId(v string) *GetUserSourceListRequest {
	s.OpenInstanceId = &v
	return s
}

func (s *GetUserSourceListRequest) SetOrgId(v int64) *GetUserSourceListRequest {
	s.OrgId = &v
	return s
}

func (s *GetUserSourceListRequest) SetOrgName(v string) *GetUserSourceListRequest {
	s.OrgName = &v
	return s
}

func (s *GetUserSourceListRequest) SetProductionType(v int32) *GetUserSourceListRequest {
	s.ProductionType = &v
	return s
}

type GetUserSourceListResponseBody struct {
	Result []*GetUserSourceListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetUserSourceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserSourceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserSourceListResponseBody) SetResult(v []*GetUserSourceListResponseBodyResult) *GetUserSourceListResponseBody {
	s.Result = v
	return s
}

type GetUserSourceListResponseBodyResult struct {
	Config      *string `json:"config,omitempty" xml:"config,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Id          *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Status      *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Vendor      *string `json:"vendor,omitempty" xml:"vendor,omitempty"`
}

func (s GetUserSourceListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetUserSourceListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetUserSourceListResponseBodyResult) SetConfig(v string) *GetUserSourceListResponseBodyResult {
	s.Config = &v
	return s
}

func (s *GetUserSourceListResponseBodyResult) SetDescription(v string) *GetUserSourceListResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetUserSourceListResponseBodyResult) SetId(v int64) *GetUserSourceListResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetUserSourceListResponseBodyResult) SetName(v string) *GetUserSourceListResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetUserSourceListResponseBodyResult) SetStatus(v int32) *GetUserSourceListResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetUserSourceListResponseBodyResult) SetVendor(v string) *GetUserSourceListResponseBodyResult {
	s.Vendor = &v
	return s
}

type GetUserSourceListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserSourceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserSourceListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserSourceListResponse) GoString() string {
	return s.String()
}

func (s *GetUserSourceListResponse) SetHeaders(v map[string]*string) *GetUserSourceListResponse {
	s.Headers = v
	return s
}

func (s *GetUserSourceListResponse) SetStatusCode(v int32) *GetUserSourceListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserSourceListResponse) SetBody(v *GetUserSourceListResponseBody) *GetUserSourceListResponse {
	s.Body = v
	return s
}

type PageListActionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PageListActionHeaders) String() string {
	return tea.Prettify(s)
}

func (s PageListActionHeaders) GoString() string {
	return s.String()
}

func (s *PageListActionHeaders) SetCommonHeaders(v map[string]*string) *PageListActionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PageListActionHeaders) SetXAcsDingtalkAccessToken(v string) *PageListActionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PageListActionRequest struct {
	MaxResults     *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken      *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OpenInstanceId *string `json:"openInstanceId,omitempty" xml:"openInstanceId,omitempty"`
	ProductionType *int64  `json:"productionType,omitempty" xml:"productionType,omitempty"`
}

func (s PageListActionRequest) String() string {
	return tea.Prettify(s)
}

func (s PageListActionRequest) GoString() string {
	return s.String()
}

func (s *PageListActionRequest) SetMaxResults(v int64) *PageListActionRequest {
	s.MaxResults = &v
	return s
}

func (s *PageListActionRequest) SetNextToken(v string) *PageListActionRequest {
	s.NextToken = &v
	return s
}

func (s *PageListActionRequest) SetOpenInstanceId(v string) *PageListActionRequest {
	s.OpenInstanceId = &v
	return s
}

func (s *PageListActionRequest) SetProductionType(v int64) *PageListActionRequest {
	s.ProductionType = &v
	return s
}

type PageListActionResponseBody struct {
	List       []*PageListActionResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextCursor *int64                            `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	Total      *int64                            `json:"total,omitempty" xml:"total,omitempty"`
}

func (s PageListActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageListActionResponseBody) GoString() string {
	return s.String()
}

func (s *PageListActionResponseBody) SetList(v []*PageListActionResponseBodyList) *PageListActionResponseBody {
	s.List = v
	return s
}

func (s *PageListActionResponseBody) SetNextCursor(v int64) *PageListActionResponseBody {
	s.NextCursor = &v
	return s
}

func (s *PageListActionResponseBody) SetTotal(v int64) *PageListActionResponseBody {
	s.Total = &v
	return s
}

type PageListActionResponseBodyList struct {
	ActionCode    *string                                        `json:"actionCode,omitempty" xml:"actionCode,omitempty"`
	ActionContent []*PageListActionResponseBodyListActionContent `json:"actionContent,omitempty" xml:"actionContent,omitempty" type:"Repeated"`
	Operator      *string                                        `json:"operator,omitempty" xml:"operator,omitempty"`
	OperatorId    *string                                        `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	OperatorRole  *string                                        `json:"operatorRole,omitempty" xml:"operatorRole,omitempty"`
}

func (s PageListActionResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s PageListActionResponseBodyList) GoString() string {
	return s.String()
}

func (s *PageListActionResponseBodyList) SetActionCode(v string) *PageListActionResponseBodyList {
	s.ActionCode = &v
	return s
}

func (s *PageListActionResponseBodyList) SetActionContent(v []*PageListActionResponseBodyListActionContent) *PageListActionResponseBodyList {
	s.ActionContent = v
	return s
}

func (s *PageListActionResponseBodyList) SetOperator(v string) *PageListActionResponseBodyList {
	s.Operator = &v
	return s
}

func (s *PageListActionResponseBodyList) SetOperatorId(v string) *PageListActionResponseBodyList {
	s.OperatorId = &v
	return s
}

func (s *PageListActionResponseBodyList) SetOperatorRole(v string) *PageListActionResponseBodyList {
	s.OperatorRole = &v
	return s
}

type PageListActionResponseBodyListActionContent struct {
	DisplayName  *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	DisplayValue *string `json:"displayValue,omitempty" xml:"displayValue,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	Value        *string `json:"value,omitempty" xml:"value,omitempty"`
	ValueType    *string `json:"valueType,omitempty" xml:"valueType,omitempty"`
}

func (s PageListActionResponseBodyListActionContent) String() string {
	return tea.Prettify(s)
}

func (s PageListActionResponseBodyListActionContent) GoString() string {
	return s.String()
}

func (s *PageListActionResponseBodyListActionContent) SetDisplayName(v string) *PageListActionResponseBodyListActionContent {
	s.DisplayName = &v
	return s
}

func (s *PageListActionResponseBodyListActionContent) SetDisplayValue(v string) *PageListActionResponseBodyListActionContent {
	s.DisplayValue = &v
	return s
}

func (s *PageListActionResponseBodyListActionContent) SetName(v string) *PageListActionResponseBodyListActionContent {
	s.Name = &v
	return s
}

func (s *PageListActionResponseBodyListActionContent) SetValue(v string) *PageListActionResponseBodyListActionContent {
	s.Value = &v
	return s
}

func (s *PageListActionResponseBodyListActionContent) SetValueType(v string) *PageListActionResponseBodyListActionContent {
	s.ValueType = &v
	return s
}

type PageListActionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PageListActionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PageListActionResponse) String() string {
	return tea.Prettify(s)
}

func (s PageListActionResponse) GoString() string {
	return s.String()
}

func (s *PageListActionResponse) SetHeaders(v map[string]*string) *PageListActionResponse {
	s.Headers = v
	return s
}

func (s *PageListActionResponse) SetStatusCode(v int32) *PageListActionResponse {
	s.StatusCode = &v
	return s
}

func (s *PageListActionResponse) SetBody(v *PageListActionResponseBody) *PageListActionResponse {
	s.Body = v
	return s
}

type PageListRobotHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PageListRobotHeaders) String() string {
	return tea.Prettify(s)
}

func (s PageListRobotHeaders) GoString() string {
	return s.String()
}

func (s *PageListRobotHeaders) SetCommonHeaders(v map[string]*string) *PageListRobotHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PageListRobotHeaders) SetXAcsDingtalkAccessToken(v string) *PageListRobotHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PageListRobotRequest struct {
	CorpId         *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	MaxResults     *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken      *int64  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OpenInstanceId *string `json:"openInstanceId,omitempty" xml:"openInstanceId,omitempty"`
	ProductionType *int32  `json:"productionType,omitempty" xml:"productionType,omitempty"`
}

func (s PageListRobotRequest) String() string {
	return tea.Prettify(s)
}

func (s PageListRobotRequest) GoString() string {
	return s.String()
}

func (s *PageListRobotRequest) SetCorpId(v string) *PageListRobotRequest {
	s.CorpId = &v
	return s
}

func (s *PageListRobotRequest) SetMaxResults(v int32) *PageListRobotRequest {
	s.MaxResults = &v
	return s
}

func (s *PageListRobotRequest) SetNextToken(v int64) *PageListRobotRequest {
	s.NextToken = &v
	return s
}

func (s *PageListRobotRequest) SetOpenInstanceId(v string) *PageListRobotRequest {
	s.OpenInstanceId = &v
	return s
}

func (s *PageListRobotRequest) SetProductionType(v int32) *PageListRobotRequest {
	s.ProductionType = &v
	return s
}

type PageListRobotResponseBody struct {
	HasMore    *bool                            `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*PageListRobotResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextCursor *int64                           `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	Total      *int64                           `json:"total,omitempty" xml:"total,omitempty"`
}

func (s PageListRobotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageListRobotResponseBody) GoString() string {
	return s.String()
}

func (s *PageListRobotResponseBody) SetHasMore(v bool) *PageListRobotResponseBody {
	s.HasMore = &v
	return s
}

func (s *PageListRobotResponseBody) SetList(v []*PageListRobotResponseBodyList) *PageListRobotResponseBody {
	s.List = v
	return s
}

func (s *PageListRobotResponseBody) SetNextCursor(v int64) *PageListRobotResponseBody {
	s.NextCursor = &v
	return s
}

func (s *PageListRobotResponseBody) SetTotal(v int64) *PageListRobotResponseBody {
	s.Total = &v
	return s
}

type PageListRobotResponseBodyList struct {
	AccountId *int64  `json:"accountId,omitempty" xml:"accountId,omitempty"`
	AppKey    *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	Id        *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Status    *int32  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PageListRobotResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s PageListRobotResponseBodyList) GoString() string {
	return s.String()
}

func (s *PageListRobotResponseBodyList) SetAccountId(v int64) *PageListRobotResponseBodyList {
	s.AccountId = &v
	return s
}

func (s *PageListRobotResponseBodyList) SetAppKey(v string) *PageListRobotResponseBodyList {
	s.AppKey = &v
	return s
}

func (s *PageListRobotResponseBodyList) SetId(v int64) *PageListRobotResponseBodyList {
	s.Id = &v
	return s
}

func (s *PageListRobotResponseBodyList) SetName(v string) *PageListRobotResponseBodyList {
	s.Name = &v
	return s
}

func (s *PageListRobotResponseBodyList) SetStatus(v int32) *PageListRobotResponseBodyList {
	s.Status = &v
	return s
}

type PageListRobotResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PageListRobotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PageListRobotResponse) String() string {
	return tea.Prettify(s)
}

func (s PageListRobotResponse) GoString() string {
	return s.String()
}

func (s *PageListRobotResponse) SetHeaders(v map[string]*string) *PageListRobotResponse {
	s.Headers = v
	return s
}

func (s *PageListRobotResponse) SetStatusCode(v int32) *PageListRobotResponse {
	s.StatusCode = &v
	return s
}

func (s *PageListRobotResponse) SetBody(v *PageListRobotResponseBody) *PageListRobotResponse {
	s.Body = v
	return s
}

type PageListTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PageListTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s PageListTicketHeaders) GoString() string {
	return s.String()
}

func (s *PageListTicketHeaders) SetCommonHeaders(v map[string]*string) *PageListTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PageListTicketHeaders) SetXAcsDingtalkAccessToken(v string) *PageListTicketHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PageListTicketRequest struct {
	EndTime        *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ForeignId      *string `json:"foreignId,omitempty" xml:"foreignId,omitempty"`
	MaxResults     *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken      *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OpenInstanceId *string `json:"openInstanceId,omitempty" xml:"openInstanceId,omitempty"`
	ProductionType *int32  `json:"productionType,omitempty" xml:"productionType,omitempty"`
	SourceId       *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	StartTime      *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	TemplateId     *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	TicketId       *string `json:"ticketId,omitempty" xml:"ticketId,omitempty"`
	TicketStatus   *string `json:"ticketStatus,omitempty" xml:"ticketStatus,omitempty"`
}

func (s PageListTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s PageListTicketRequest) GoString() string {
	return s.String()
}

func (s *PageListTicketRequest) SetEndTime(v int64) *PageListTicketRequest {
	s.EndTime = &v
	return s
}

func (s *PageListTicketRequest) SetForeignId(v string) *PageListTicketRequest {
	s.ForeignId = &v
	return s
}

func (s *PageListTicketRequest) SetMaxResults(v int32) *PageListTicketRequest {
	s.MaxResults = &v
	return s
}

func (s *PageListTicketRequest) SetNextToken(v string) *PageListTicketRequest {
	s.NextToken = &v
	return s
}

func (s *PageListTicketRequest) SetOpenInstanceId(v string) *PageListTicketRequest {
	s.OpenInstanceId = &v
	return s
}

func (s *PageListTicketRequest) SetProductionType(v int32) *PageListTicketRequest {
	s.ProductionType = &v
	return s
}

func (s *PageListTicketRequest) SetSourceId(v string) *PageListTicketRequest {
	s.SourceId = &v
	return s
}

func (s *PageListTicketRequest) SetStartTime(v int64) *PageListTicketRequest {
	s.StartTime = &v
	return s
}

func (s *PageListTicketRequest) SetTemplateId(v string) *PageListTicketRequest {
	s.TemplateId = &v
	return s
}

func (s *PageListTicketRequest) SetTicketId(v string) *PageListTicketRequest {
	s.TicketId = &v
	return s
}

func (s *PageListTicketRequest) SetTicketStatus(v string) *PageListTicketRequest {
	s.TicketStatus = &v
	return s
}

type PageListTicketResponseBody struct {
	List       []*PageListTicketResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextCursor *int64                            `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	Total      *int64                            `json:"total,omitempty" xml:"total,omitempty"`
}

func (s PageListTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageListTicketResponseBody) GoString() string {
	return s.String()
}

func (s *PageListTicketResponseBody) SetList(v []*PageListTicketResponseBodyList) *PageListTicketResponseBody {
	s.List = v
	return s
}

func (s *PageListTicketResponseBody) SetNextCursor(v int64) *PageListTicketResponseBody {
	s.NextCursor = &v
	return s
}

func (s *PageListTicketResponseBody) SetTotal(v int64) *PageListTicketResponseBody {
	s.Total = &v
	return s
}

type PageListTicketResponseBodyList struct {
	BizDataMap     map[string]interface{} `json:"bizDataMap,omitempty" xml:"bizDataMap,omitempty"`
	ForeignId      *string                `json:"foreignId,omitempty" xml:"foreignId,omitempty"`
	ForeignName    *string                `json:"foreignName,omitempty" xml:"foreignName,omitempty"`
	GmtCreate      *string                `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string                `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	OpenInstanceId *string                `json:"openInstanceId,omitempty" xml:"openInstanceId,omitempty"`
	ProductionType *int32                 `json:"productionType,omitempty" xml:"productionType,omitempty"`
	SourceId       *string                `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	TemplateId     *string                `json:"templateId,omitempty" xml:"templateId,omitempty"`
	TicketId       *string                `json:"ticketId,omitempty" xml:"ticketId,omitempty"`
	TicketStatus   *string                `json:"ticketStatus,omitempty" xml:"ticketStatus,omitempty"`
	Title          *string                `json:"title,omitempty" xml:"title,omitempty"`
}

func (s PageListTicketResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s PageListTicketResponseBodyList) GoString() string {
	return s.String()
}

func (s *PageListTicketResponseBodyList) SetBizDataMap(v map[string]interface{}) *PageListTicketResponseBodyList {
	s.BizDataMap = v
	return s
}

func (s *PageListTicketResponseBodyList) SetForeignId(v string) *PageListTicketResponseBodyList {
	s.ForeignId = &v
	return s
}

func (s *PageListTicketResponseBodyList) SetForeignName(v string) *PageListTicketResponseBodyList {
	s.ForeignName = &v
	return s
}

func (s *PageListTicketResponseBodyList) SetGmtCreate(v string) *PageListTicketResponseBodyList {
	s.GmtCreate = &v
	return s
}

func (s *PageListTicketResponseBodyList) SetGmtModified(v string) *PageListTicketResponseBodyList {
	s.GmtModified = &v
	return s
}

func (s *PageListTicketResponseBodyList) SetOpenInstanceId(v string) *PageListTicketResponseBodyList {
	s.OpenInstanceId = &v
	return s
}

func (s *PageListTicketResponseBodyList) SetProductionType(v int32) *PageListTicketResponseBodyList {
	s.ProductionType = &v
	return s
}

func (s *PageListTicketResponseBodyList) SetSourceId(v string) *PageListTicketResponseBodyList {
	s.SourceId = &v
	return s
}

func (s *PageListTicketResponseBodyList) SetTemplateId(v string) *PageListTicketResponseBodyList {
	s.TemplateId = &v
	return s
}

func (s *PageListTicketResponseBodyList) SetTicketId(v string) *PageListTicketResponseBodyList {
	s.TicketId = &v
	return s
}

func (s *PageListTicketResponseBodyList) SetTicketStatus(v string) *PageListTicketResponseBodyList {
	s.TicketStatus = &v
	return s
}

func (s *PageListTicketResponseBodyList) SetTitle(v string) *PageListTicketResponseBodyList {
	s.Title = &v
	return s
}

type PageListTicketResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PageListTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PageListTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s PageListTicketResponse) GoString() string {
	return s.String()
}

func (s *PageListTicketResponse) SetHeaders(v map[string]*string) *PageListTicketResponse {
	s.Headers = v
	return s
}

func (s *PageListTicketResponse) SetStatusCode(v int32) *PageListTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *PageListTicketResponse) SetBody(v *PageListTicketResponseBody) *PageListTicketResponse {
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

func (client *Client) CreateTicketWithOptions(request *CreateTicketRequest, headers *CreateTicketHeaders, runtime *util.RuntimeOptions) (_result *CreateTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ForeignId)) {
		body["foreignId"] = request.ForeignId
	}

	if !tea.BoolValue(util.IsUnset(request.ForeignName)) {
		body["foreignName"] = request.ForeignName
	}

	if !tea.BoolValue(util.IsUnset(request.OpenInstanceId)) {
		body["openInstanceId"] = request.OpenInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductionType)) {
		body["productionType"] = request.ProductionType
	}

	if !tea.BoolValue(util.IsUnset(request.Properties)) {
		body["properties"] = request.Properties
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		body["sourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
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
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTicket"),
		Version:     tea.String("customerService_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/customerService/tickets"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTicketResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTicket(request *CreateTicketRequest) (_result *CreateTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTicketHeaders{}
	_result = &CreateTicketResponse{}
	_body, _err := client.CreateTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteActivityWithOptions(ticketId *string, request *ExecuteActivityRequest, headers *ExecuteActivityHeaders, runtime *util.RuntimeOptions) (_result *ExecuteActivityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivityCode)) {
		body["activityCode"] = request.ActivityCode
	}

	if !tea.BoolValue(util.IsUnset(request.ForeignId)) {
		body["foreignId"] = request.ForeignId
	}

	if !tea.BoolValue(util.IsUnset(request.ForeignName)) {
		body["foreignName"] = request.ForeignName
	}

	if !tea.BoolValue(util.IsUnset(request.OpenInstanceId)) {
		body["openInstanceId"] = request.OpenInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductionType)) {
		body["productionType"] = request.ProductionType
	}

	if !tea.BoolValue(util.IsUnset(request.Properties)) {
		body["properties"] = request.Properties
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		body["sourceId"] = request.SourceId
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
		Action:      tea.String("ExecuteActivity"),
		Version:     tea.String("customerService_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/customerService/tickets/" + tea.StringValue(ticketId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteActivityResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteActivity(ticketId *string, request *ExecuteActivityRequest) (_result *ExecuteActivityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ExecuteActivityHeaders{}
	_result = &ExecuteActivityResponse{}
	_body, _err := client.ExecuteActivityWithOptions(ticketId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserSourceListWithOptions(request *GetUserSourceListRequest, headers *GetUserSourceListHeaders, runtime *util.RuntimeOptions) (_result *GetUserSourceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.OpenInstanceId)) {
		query["openInstanceId"] = request.OpenInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		query["orgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		query["orgName"] = request.OrgName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductionType)) {
		query["productionType"] = request.ProductionType
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
		Action:      tea.String("GetUserSourceList"),
		Version:     tea.String("customerService_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/customerService/customers/sources"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserSourceListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserSourceList(request *GetUserSourceListRequest) (_result *GetUserSourceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserSourceListHeaders{}
	_result = &GetUserSourceListResponse{}
	_body, _err := client.GetUserSourceListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PageListActionWithOptions(ticketId *string, request *PageListActionRequest, headers *PageListActionHeaders, runtime *util.RuntimeOptions) (_result *PageListActionResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.OpenInstanceId)) {
		query["openInstanceId"] = request.OpenInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductionType)) {
		query["productionType"] = request.ProductionType
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
		Action:      tea.String("PageListAction"),
		Version:     tea.String("customerService_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/customerService/tickets/" + tea.StringValue(ticketId) + "/actions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PageListActionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PageListAction(ticketId *string, request *PageListActionRequest) (_result *PageListActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PageListActionHeaders{}
	_result = &PageListActionResponse{}
	_body, _err := client.PageListActionWithOptions(ticketId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PageListRobotWithOptions(request *PageListRobotRequest, headers *PageListRobotHeaders, runtime *util.RuntimeOptions) (_result *PageListRobotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OpenInstanceId)) {
		query["openInstanceId"] = request.OpenInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductionType)) {
		query["productionType"] = request.ProductionType
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
		Action:      tea.String("PageListRobot"),
		Version:     tea.String("customerService_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/customerService/robots"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PageListRobotResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PageListRobot(request *PageListRobotRequest) (_result *PageListRobotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PageListRobotHeaders{}
	_result = &PageListRobotResponse{}
	_body, _err := client.PageListRobotWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PageListTicketWithOptions(request *PageListTicketRequest, headers *PageListTicketHeaders, runtime *util.RuntimeOptions) (_result *PageListTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ForeignId)) {
		query["foreignId"] = request.ForeignId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OpenInstanceId)) {
		query["openInstanceId"] = request.OpenInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductionType)) {
		query["productionType"] = request.ProductionType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		query["sourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["templateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TicketId)) {
		query["ticketId"] = request.TicketId
	}

	if !tea.BoolValue(util.IsUnset(request.TicketStatus)) {
		query["ticketStatus"] = request.TicketStatus
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
		Action:      tea.String("PageListTicket"),
		Version:     tea.String("customerService_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/customerService/tickets"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PageListTicketResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PageListTicket(request *PageListTicketRequest) (_result *PageListTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PageListTicketHeaders{}
	_result = &PageListTicketResponse{}
	_body, _err := client.PageListTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
