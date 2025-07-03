// This file is auto-generated, don't edit it. Thanks.
package salary_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetSalaryCalculationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSalaryCalculationHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSalaryCalculationHeaders) GoString() string {
	return s.String()
}

func (s *GetSalaryCalculationHeaders) SetCommonHeaders(v map[string]*string) *GetSalaryCalculationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSalaryCalculationHeaders) SetXAcsDingtalkAccessToken(v string) *GetSalaryCalculationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSalaryCalculationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2025-06
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SalaryGroupxxx
	SalaryGroupId *string `json:"salaryGroupId,omitempty" xml:"salaryGroupId,omitempty"`
}

func (s GetSalaryCalculationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSalaryCalculationRequest) GoString() string {
	return s.String()
}

func (s *GetSalaryCalculationRequest) SetDate(v string) *GetSalaryCalculationRequest {
	s.Date = &v
	return s
}

func (s *GetSalaryCalculationRequest) SetSalaryGroupId(v string) *GetSalaryCalculationRequest {
	s.SalaryGroupId = &v
	return s
}

type GetSalaryCalculationResponseBody struct {
	Result *GetSalaryCalculationResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetSalaryCalculationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSalaryCalculationResponseBody) GoString() string {
	return s.String()
}

func (s *GetSalaryCalculationResponseBody) SetResult(v *GetSalaryCalculationResponseBodyResult) *GetSalaryCalculationResponseBody {
	s.Result = v
	return s
}

type GetSalaryCalculationResponseBodyResult struct {
	CalStatus *bool `json:"calStatus,omitempty" xml:"calStatus,omitempty"`
	// example:
	//
	// 2025-06-30
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// example:
	//
	// 2025-06-01
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// example:
	//
	// NONE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetSalaryCalculationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetSalaryCalculationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSalaryCalculationResponseBodyResult) SetCalStatus(v bool) *GetSalaryCalculationResponseBodyResult {
	s.CalStatus = &v
	return s
}

func (s *GetSalaryCalculationResponseBodyResult) SetEndDate(v string) *GetSalaryCalculationResponseBodyResult {
	s.EndDate = &v
	return s
}

func (s *GetSalaryCalculationResponseBodyResult) SetStartDate(v string) *GetSalaryCalculationResponseBodyResult {
	s.StartDate = &v
	return s
}

func (s *GetSalaryCalculationResponseBodyResult) SetStatus(v string) *GetSalaryCalculationResponseBodyResult {
	s.Status = &v
	return s
}

type GetSalaryCalculationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSalaryCalculationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSalaryCalculationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSalaryCalculationResponse) GoString() string {
	return s.String()
}

func (s *GetSalaryCalculationResponse) SetHeaders(v map[string]*string) *GetSalaryCalculationResponse {
	s.Headers = v
	return s
}

func (s *GetSalaryCalculationResponse) SetStatusCode(v int32) *GetSalaryCalculationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSalaryCalculationResponse) SetBody(v *GetSalaryCalculationResponseBody) *GetSalaryCalculationResponse {
	s.Body = v
	return s
}

type GetSalaryGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSalaryGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSalaryGroupHeaders) GoString() string {
	return s.String()
}

func (s *GetSalaryGroupHeaders) SetCommonHeaders(v map[string]*string) *GetSalaryGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSalaryGroupHeaders) SetXAcsDingtalkAccessToken(v string) *GetSalaryGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSalaryGroupResponseBody struct {
	Result []*GetSalaryGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetSalaryGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSalaryGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetSalaryGroupResponseBody) SetResult(v []*GetSalaryGroupResponseBodyResult) *GetSalaryGroupResponseBody {
	s.Result = v
	return s
}

type GetSalaryGroupResponseBodyResult struct {
	EnableFlag *bool `json:"enableFlag,omitempty" xml:"enableFlag,omitempty"`
	// example:
	//
	// SalaryItemGroupIdxxx
	SalaryGroupId *string `json:"salaryGroupId,omitempty" xml:"salaryGroupId,omitempty"`
	// example:
	//
	// 123
	SalaryGroupName *string `json:"salaryGroupName,omitempty" xml:"salaryGroupName,omitempty"`
}

func (s GetSalaryGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetSalaryGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSalaryGroupResponseBodyResult) SetEnableFlag(v bool) *GetSalaryGroupResponseBodyResult {
	s.EnableFlag = &v
	return s
}

func (s *GetSalaryGroupResponseBodyResult) SetSalaryGroupId(v string) *GetSalaryGroupResponseBodyResult {
	s.SalaryGroupId = &v
	return s
}

func (s *GetSalaryGroupResponseBodyResult) SetSalaryGroupName(v string) *GetSalaryGroupResponseBodyResult {
	s.SalaryGroupName = &v
	return s
}

type GetSalaryGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSalaryGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSalaryGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSalaryGroupResponse) GoString() string {
	return s.String()
}

func (s *GetSalaryGroupResponse) SetHeaders(v map[string]*string) *GetSalaryGroupResponse {
	s.Headers = v
	return s
}

func (s *GetSalaryGroupResponse) SetStatusCode(v int32) *GetSalaryGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSalaryGroupResponse) SetBody(v *GetSalaryGroupResponseBody) *GetSalaryGroupResponse {
	s.Body = v
	return s
}

type GetSalaryItemHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSalaryItemHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSalaryItemHeaders) GoString() string {
	return s.String()
}

func (s *GetSalaryItemHeaders) SetCommonHeaders(v map[string]*string) *GetSalaryItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSalaryItemHeaders) SetXAcsDingtalkAccessToken(v string) *GetSalaryItemHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSalaryItemRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// SalaryItemGroupxxx
	SalaryItemGroupId *string `json:"salaryItemGroupId,omitempty" xml:"salaryItemGroupId,omitempty"`
}

func (s GetSalaryItemRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSalaryItemRequest) GoString() string {
	return s.String()
}

func (s *GetSalaryItemRequest) SetSalaryItemGroupId(v string) *GetSalaryItemRequest {
	s.SalaryItemGroupId = &v
	return s
}

type GetSalaryItemResponseBody struct {
	Result []*GetSalaryItemResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetSalaryItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSalaryItemResponseBody) GoString() string {
	return s.String()
}

func (s *GetSalaryItemResponseBody) SetResult(v []*GetSalaryItemResponseBodyResult) *GetSalaryItemResponseBody {
	s.Result = v
	return s
}

type GetSalaryItemResponseBodyResult struct {
	// example:
	//
	// SalaryItemIdxxx
	SalaryItemId *string `json:"salaryItemId,omitempty" xml:"salaryItemId,omitempty"`
	// example:
	//
	// 绩效xx
	SalaryItemName *string `json:"salaryItemName,omitempty" xml:"salaryItemName,omitempty"`
}

func (s GetSalaryItemResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetSalaryItemResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSalaryItemResponseBodyResult) SetSalaryItemId(v string) *GetSalaryItemResponseBodyResult {
	s.SalaryItemId = &v
	return s
}

func (s *GetSalaryItemResponseBodyResult) SetSalaryItemName(v string) *GetSalaryItemResponseBodyResult {
	s.SalaryItemName = &v
	return s
}

type GetSalaryItemResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSalaryItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSalaryItemResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSalaryItemResponse) GoString() string {
	return s.String()
}

func (s *GetSalaryItemResponse) SetHeaders(v map[string]*string) *GetSalaryItemResponse {
	s.Headers = v
	return s
}

func (s *GetSalaryItemResponse) SetStatusCode(v int32) *GetSalaryItemResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSalaryItemResponse) SetBody(v *GetSalaryItemResponseBody) *GetSalaryItemResponse {
	s.Body = v
	return s
}

type GetSalaryItemGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSalaryItemGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSalaryItemGroupHeaders) GoString() string {
	return s.String()
}

func (s *GetSalaryItemGroupHeaders) SetCommonHeaders(v map[string]*string) *GetSalaryItemGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSalaryItemGroupHeaders) SetXAcsDingtalkAccessToken(v string) *GetSalaryItemGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSalaryItemGroupResponseBody struct {
	Result []*GetSalaryItemGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetSalaryItemGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSalaryItemGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetSalaryItemGroupResponseBody) SetResult(v []*GetSalaryItemGroupResponseBodyResult) *GetSalaryItemGroupResponseBody {
	s.Result = v
	return s
}

type GetSalaryItemGroupResponseBodyResult struct {
	// example:
	//
	// SalaryItemGroupIdxxx
	SalaryItemGroupId *string `json:"salaryItemGroupId,omitempty" xml:"salaryItemGroupId,omitempty"`
	// example:
	//
	// 浮动薪资xx
	SalaryItemGroupName *string `json:"salaryItemGroupName,omitempty" xml:"salaryItemGroupName,omitempty"`
}

func (s GetSalaryItemGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetSalaryItemGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSalaryItemGroupResponseBodyResult) SetSalaryItemGroupId(v string) *GetSalaryItemGroupResponseBodyResult {
	s.SalaryItemGroupId = &v
	return s
}

func (s *GetSalaryItemGroupResponseBodyResult) SetSalaryItemGroupName(v string) *GetSalaryItemGroupResponseBodyResult {
	s.SalaryItemGroupName = &v
	return s
}

type GetSalaryItemGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSalaryItemGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSalaryItemGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSalaryItemGroupResponse) GoString() string {
	return s.String()
}

func (s *GetSalaryItemGroupResponse) SetHeaders(v map[string]*string) *GetSalaryItemGroupResponse {
	s.Headers = v
	return s
}

func (s *GetSalaryItemGroupResponse) SetStatusCode(v int32) *GetSalaryItemGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSalaryItemGroupResponse) SetBody(v *GetSalaryItemGroupResponseBody) *GetSalaryItemGroupResponse {
	s.Body = v
	return s
}

type ListSalaryCalculationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListSalaryCalculationHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSalaryCalculationHeaders) GoString() string {
	return s.String()
}

func (s *ListSalaryCalculationHeaders) SetCommonHeaders(v map[string]*string) *ListSalaryCalculationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSalaryCalculationHeaders) SetXAcsDingtalkAccessToken(v string) *ListSalaryCalculationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListSalaryCalculationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2025-06
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
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
	// SalaryGroupXXX
	SalaryGroupId *string `json:"salaryGroupId,omitempty" xml:"salaryGroupId,omitempty"`
}

func (s ListSalaryCalculationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSalaryCalculationRequest) GoString() string {
	return s.String()
}

func (s *ListSalaryCalculationRequest) SetDate(v string) *ListSalaryCalculationRequest {
	s.Date = &v
	return s
}

func (s *ListSalaryCalculationRequest) SetPageIndex(v int32) *ListSalaryCalculationRequest {
	s.PageIndex = &v
	return s
}

func (s *ListSalaryCalculationRequest) SetPageSize(v int32) *ListSalaryCalculationRequest {
	s.PageSize = &v
	return s
}

func (s *ListSalaryCalculationRequest) SetSalaryGroupId(v string) *ListSalaryCalculationRequest {
	s.SalaryGroupId = &v
	return s
}

type ListSalaryCalculationResponseBody struct {
	Result *ListSalaryCalculationResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ListSalaryCalculationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSalaryCalculationResponseBody) GoString() string {
	return s.String()
}

func (s *ListSalaryCalculationResponseBody) SetResult(v *ListSalaryCalculationResponseBodyResult) *ListSalaryCalculationResponseBody {
	s.Result = v
	return s
}

type ListSalaryCalculationResponseBodyResult struct {
	Data    []*ListSalaryCalculationResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	HasMore *bool                                          `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
}

func (s ListSalaryCalculationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListSalaryCalculationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSalaryCalculationResponseBodyResult) SetData(v []*ListSalaryCalculationResponseBodyResultData) *ListSalaryCalculationResponseBodyResult {
	s.Data = v
	return s
}

func (s *ListSalaryCalculationResponseBodyResult) SetHasMore(v bool) *ListSalaryCalculationResponseBodyResult {
	s.HasMore = &v
	return s
}

type ListSalaryCalculationResponseBodyResultData struct {
	DataList []*ListSalaryCalculationResponseBodyResultDataDataList `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	UserId   *string                                                `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListSalaryCalculationResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s ListSalaryCalculationResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *ListSalaryCalculationResponseBodyResultData) SetDataList(v []*ListSalaryCalculationResponseBodyResultDataDataList) *ListSalaryCalculationResponseBodyResultData {
	s.DataList = v
	return s
}

func (s *ListSalaryCalculationResponseBodyResultData) SetUserId(v string) *ListSalaryCalculationResponseBodyResultData {
	s.UserId = &v
	return s
}

type ListSalaryCalculationResponseBodyResultDataDataList struct {
	SalaryItemId    *string `json:"salaryItemId,omitempty" xml:"salaryItemId,omitempty"`
	SalaryItemName  *string `json:"salaryItemName,omitempty" xml:"salaryItemName,omitempty"`
	SalaryItemValue *string `json:"salaryItemValue,omitempty" xml:"salaryItemValue,omitempty"`
}

func (s ListSalaryCalculationResponseBodyResultDataDataList) String() string {
	return tea.Prettify(s)
}

func (s ListSalaryCalculationResponseBodyResultDataDataList) GoString() string {
	return s.String()
}

func (s *ListSalaryCalculationResponseBodyResultDataDataList) SetSalaryItemId(v string) *ListSalaryCalculationResponseBodyResultDataDataList {
	s.SalaryItemId = &v
	return s
}

func (s *ListSalaryCalculationResponseBodyResultDataDataList) SetSalaryItemName(v string) *ListSalaryCalculationResponseBodyResultDataDataList {
	s.SalaryItemName = &v
	return s
}

func (s *ListSalaryCalculationResponseBodyResultDataDataList) SetSalaryItemValue(v string) *ListSalaryCalculationResponseBodyResultDataDataList {
	s.SalaryItemValue = &v
	return s
}

type ListSalaryCalculationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSalaryCalculationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSalaryCalculationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSalaryCalculationResponse) GoString() string {
	return s.String()
}

func (s *ListSalaryCalculationResponse) SetHeaders(v map[string]*string) *ListSalaryCalculationResponse {
	s.Headers = v
	return s
}

func (s *ListSalaryCalculationResponse) SetStatusCode(v int32) *ListSalaryCalculationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSalaryCalculationResponse) SetBody(v *ListSalaryCalculationResponseBody) *ListSalaryCalculationResponse {
	s.Body = v
	return s
}

type WriteSalaryCalculationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s WriteSalaryCalculationHeaders) String() string {
	return tea.Prettify(s)
}

func (s WriteSalaryCalculationHeaders) GoString() string {
	return s.String()
}

func (s *WriteSalaryCalculationHeaders) SetCommonHeaders(v map[string]*string) *WriteSalaryCalculationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *WriteSalaryCalculationHeaders) SetXAcsDingtalkAccessToken(v string) *WriteSalaryCalculationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type WriteSalaryCalculationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2025-06
	Date  *string                               `json:"date,omitempty" xml:"date,omitempty"`
	Items []*WriteSalaryCalculationRequestItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s WriteSalaryCalculationRequest) String() string {
	return tea.Prettify(s)
}

func (s WriteSalaryCalculationRequest) GoString() string {
	return s.String()
}

func (s *WriteSalaryCalculationRequest) SetDate(v string) *WriteSalaryCalculationRequest {
	s.Date = &v
	return s
}

func (s *WriteSalaryCalculationRequest) SetItems(v []*WriteSalaryCalculationRequestItems) *WriteSalaryCalculationRequest {
	s.Items = v
	return s
}

type WriteSalaryCalculationRequestItems struct {
	Contents []*WriteSalaryCalculationRequestItemsContents `json:"contents,omitempty" xml:"contents,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// user001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s WriteSalaryCalculationRequestItems) String() string {
	return tea.Prettify(s)
}

func (s WriteSalaryCalculationRequestItems) GoString() string {
	return s.String()
}

func (s *WriteSalaryCalculationRequestItems) SetContents(v []*WriteSalaryCalculationRequestItemsContents) *WriteSalaryCalculationRequestItems {
	s.Contents = v
	return s
}

func (s *WriteSalaryCalculationRequestItems) SetUserId(v string) *WriteSalaryCalculationRequestItems {
	s.UserId = &v
	return s
}

type WriteSalaryCalculationRequestItemsContents struct {
	// This parameter is required.
	//
	// example:
	//
	// SalaryItemXXX
	SalaryItemId *string `json:"salaryItemId,omitempty" xml:"salaryItemId,omitempty"`
	// example:
	//
	// 26
	SalaryItemValue *string `json:"salaryItemValue,omitempty" xml:"salaryItemValue,omitempty"`
}

func (s WriteSalaryCalculationRequestItemsContents) String() string {
	return tea.Prettify(s)
}

func (s WriteSalaryCalculationRequestItemsContents) GoString() string {
	return s.String()
}

func (s *WriteSalaryCalculationRequestItemsContents) SetSalaryItemId(v string) *WriteSalaryCalculationRequestItemsContents {
	s.SalaryItemId = &v
	return s
}

func (s *WriteSalaryCalculationRequestItemsContents) SetSalaryItemValue(v string) *WriteSalaryCalculationRequestItemsContents {
	s.SalaryItemValue = &v
	return s
}

type WriteSalaryCalculationResponseBody struct {
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s WriteSalaryCalculationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WriteSalaryCalculationResponseBody) GoString() string {
	return s.String()
}

func (s *WriteSalaryCalculationResponseBody) SetResult(v map[string]interface{}) *WriteSalaryCalculationResponseBody {
	s.Result = v
	return s
}

type WriteSalaryCalculationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WriteSalaryCalculationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WriteSalaryCalculationResponse) String() string {
	return tea.Prettify(s)
}

func (s WriteSalaryCalculationResponse) GoString() string {
	return s.String()
}

func (s *WriteSalaryCalculationResponse) SetHeaders(v map[string]*string) *WriteSalaryCalculationResponse {
	s.Headers = v
	return s
}

func (s *WriteSalaryCalculationResponse) SetStatusCode(v int32) *WriteSalaryCalculationResponse {
	s.StatusCode = &v
	return s
}

func (s *WriteSalaryCalculationResponse) SetBody(v *WriteSalaryCalculationResponseBody) *WriteSalaryCalculationResponse {
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
// 小微薪酬获取薪资记录
//
// @param request - GetSalaryCalculationRequest
//
// @param headers - GetSalaryCalculationHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSalaryCalculationResponse
func (client *Client) GetSalaryCalculationWithOptions(request *GetSalaryCalculationRequest, headers *GetSalaryCalculationHeaders, runtime *util.RuntimeOptions) (_result *GetSalaryCalculationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Date)) {
		query["date"] = request.Date
	}

	if !tea.BoolValue(util.IsUnset(request.SalaryGroupId)) {
		query["salaryGroupId"] = request.SalaryGroupId
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
		Action:      tea.String("GetSalaryCalculation"),
		Version:     tea.String("salary_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/salary/calculation"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSalaryCalculationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 小微薪酬获取薪资记录
//
// @param request - GetSalaryCalculationRequest
//
// @return GetSalaryCalculationResponse
func (client *Client) GetSalaryCalculation(request *GetSalaryCalculationRequest) (_result *GetSalaryCalculationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSalaryCalculationHeaders{}
	_result = &GetSalaryCalculationResponse{}
	_body, _err := client.GetSalaryCalculationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 小微薪酬获取薪资组
//
// @param headers - GetSalaryGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSalaryGroupResponse
func (client *Client) GetSalaryGroupWithOptions(headers *GetSalaryGroupHeaders, runtime *util.RuntimeOptions) (_result *GetSalaryGroupResponse, _err error) {
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
		Action:      tea.String("GetSalaryGroup"),
		Version:     tea.String("salary_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/salary/group"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSalaryGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 小微薪酬获取薪资组
//
// @return GetSalaryGroupResponse
func (client *Client) GetSalaryGroup() (_result *GetSalaryGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSalaryGroupHeaders{}
	_result = &GetSalaryGroupResponse{}
	_body, _err := client.GetSalaryGroupWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 小微薪酬获取薪资项目
//
// @param request - GetSalaryItemRequest
//
// @param headers - GetSalaryItemHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSalaryItemResponse
func (client *Client) GetSalaryItemWithOptions(request *GetSalaryItemRequest, headers *GetSalaryItemHeaders, runtime *util.RuntimeOptions) (_result *GetSalaryItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SalaryItemGroupId)) {
		query["salaryItemGroupId"] = request.SalaryItemGroupId
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
		Action:      tea.String("GetSalaryItem"),
		Version:     tea.String("salary_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/salary/item"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSalaryItemResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 小微薪酬获取薪资项目
//
// @param request - GetSalaryItemRequest
//
// @return GetSalaryItemResponse
func (client *Client) GetSalaryItem(request *GetSalaryItemRequest) (_result *GetSalaryItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSalaryItemHeaders{}
	_result = &GetSalaryItemResponse{}
	_body, _err := client.GetSalaryItemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 小微薪酬获取薪资项目大类
//
// @param headers - GetSalaryItemGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSalaryItemGroupResponse
func (client *Client) GetSalaryItemGroupWithOptions(headers *GetSalaryItemGroupHeaders, runtime *util.RuntimeOptions) (_result *GetSalaryItemGroupResponse, _err error) {
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
		Action:      tea.String("GetSalaryItemGroup"),
		Version:     tea.String("salary_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/salary/itemGroup"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSalaryItemGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 小微薪酬获取薪资项目大类
//
// @return GetSalaryItemGroupResponse
func (client *Client) GetSalaryItemGroup() (_result *GetSalaryItemGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSalaryItemGroupHeaders{}
	_result = &GetSalaryItemGroupResponse{}
	_body, _err := client.GetSalaryItemGroupWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 小微薪酬获取薪资记录数据
//
// @param request - ListSalaryCalculationRequest
//
// @param headers - ListSalaryCalculationHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSalaryCalculationResponse
func (client *Client) ListSalaryCalculationWithOptions(request *ListSalaryCalculationRequest, headers *ListSalaryCalculationHeaders, runtime *util.RuntimeOptions) (_result *ListSalaryCalculationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Date)) {
		body["date"] = request.Date
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		body["pageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SalaryGroupId)) {
		body["salaryGroupId"] = request.SalaryGroupId
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
		Action:      tea.String("ListSalaryCalculation"),
		Version:     tea.String("salary_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/salary/calculation"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSalaryCalculationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 小微薪酬获取薪资记录数据
//
// @param request - ListSalaryCalculationRequest
//
// @return ListSalaryCalculationResponse
func (client *Client) ListSalaryCalculation(request *ListSalaryCalculationRequest) (_result *ListSalaryCalculationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSalaryCalculationHeaders{}
	_result = &ListSalaryCalculationResponse{}
	_body, _err := client.ListSalaryCalculationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 小微薪酬获取薪资记录写入
//
// @param request - WriteSalaryCalculationRequest
//
// @param headers - WriteSalaryCalculationHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WriteSalaryCalculationResponse
func (client *Client) WriteSalaryCalculationWithOptions(request *WriteSalaryCalculationRequest, headers *WriteSalaryCalculationHeaders, runtime *util.RuntimeOptions) (_result *WriteSalaryCalculationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Date)) {
		body["date"] = request.Date
	}

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
		Action:      tea.String("WriteSalaryCalculation"),
		Version:     tea.String("salary_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/salary/calculation/write"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &WriteSalaryCalculationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 小微薪酬获取薪资记录写入
//
// @param request - WriteSalaryCalculationRequest
//
// @return WriteSalaryCalculationResponse
func (client *Client) WriteSalaryCalculation(request *WriteSalaryCalculationRequest) (_result *WriteSalaryCalculationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &WriteSalaryCalculationHeaders{}
	_result = &WriteSalaryCalculationResponse{}
	_body, _err := client.WriteSalaryCalculationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
