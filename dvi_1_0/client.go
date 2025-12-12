// This file is auto-generated, don't edit it. Thanks.
package dvi_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetAudioFileDownloadInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAudioFileDownloadInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAudioFileDownloadInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetAudioFileDownloadInfoHeaders) SetCommonHeaders(v map[string]*string) *GetAudioFileDownloadInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAudioFileDownloadInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetAudioFileDownloadInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAudioFileDownloadInfoRequest struct {
	// This parameter is required.
	DeviceType *string `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// da5ad92d-79dc-4599-8f92-ba8209c68569
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
}

func (s GetAudioFileDownloadInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAudioFileDownloadInfoRequest) GoString() string {
	return s.String()
}

func (s *GetAudioFileDownloadInfoRequest) SetDeviceType(v string) *GetAudioFileDownloadInfoRequest {
	s.DeviceType = &v
	return s
}

func (s *GetAudioFileDownloadInfoRequest) SetFileId(v string) *GetAudioFileDownloadInfoRequest {
	s.FileId = &v
	return s
}

type GetAudioFileDownloadInfoResponseBody struct {
	Result *GetAudioFileDownloadInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetAudioFileDownloadInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAudioFileDownloadInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAudioFileDownloadInfoResponseBody) SetResult(v *GetAudioFileDownloadInfoResponseBodyResult) *GetAudioFileDownloadInfoResponseBody {
	s.Result = v
	return s
}

type GetAudioFileDownloadInfoResponseBodyResult struct {
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetAudioFileDownloadInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetAudioFileDownloadInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAudioFileDownloadInfoResponseBodyResult) SetUrl(v string) *GetAudioFileDownloadInfoResponseBodyResult {
	s.Url = &v
	return s
}

type GetAudioFileDownloadInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAudioFileDownloadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAudioFileDownloadInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAudioFileDownloadInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAudioFileDownloadInfoResponse) SetHeaders(v map[string]*string) *GetAudioFileDownloadInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAudioFileDownloadInfoResponse) SetStatusCode(v int32) *GetAudioFileDownloadInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAudioFileDownloadInfoResponse) SetBody(v *GetAudioFileDownloadInfoResponseBody) *GetAudioFileDownloadInfoResponse {
	s.Body = v
	return s
}

type GetAudioFileInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAudioFileInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAudioFileInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetAudioFileInfoHeaders) SetCommonHeaders(v map[string]*string) *GetAudioFileInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAudioFileInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetAudioFileInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAudioFileInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// A1
	DeviceType *string `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
	// This parameter is required.
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
}

func (s GetAudioFileInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAudioFileInfoRequest) GoString() string {
	return s.String()
}

func (s *GetAudioFileInfoRequest) SetDeviceType(v string) *GetAudioFileInfoRequest {
	s.DeviceType = &v
	return s
}

func (s *GetAudioFileInfoRequest) SetFileId(v string) *GetAudioFileInfoRequest {
	s.FileId = &v
	return s
}

type GetAudioFileInfoResponseBody struct {
	Result *GetAudioFileInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetAudioFileInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAudioFileInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAudioFileInfoResponseBody) SetResult(v *GetAudioFileInfoResponseBodyResult) *GetAudioFileInfoResponseBody {
	s.Result = v
	return s
}

type GetAudioFileInfoResponseBodyResult struct {
	CreateTime    *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	Duration      *int64  `json:"duration,omitempty" xml:"duration,omitempty"`
	FileId        *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	FileName      *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileSize      *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
}

func (s GetAudioFileInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetAudioFileInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAudioFileInfoResponseBodyResult) SetCreateTime(v int64) *GetAudioFileInfoResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetAudioFileInfoResponseBodyResult) SetCreatorUserId(v string) *GetAudioFileInfoResponseBodyResult {
	s.CreatorUserId = &v
	return s
}

func (s *GetAudioFileInfoResponseBodyResult) SetDuration(v int64) *GetAudioFileInfoResponseBodyResult {
	s.Duration = &v
	return s
}

func (s *GetAudioFileInfoResponseBodyResult) SetFileId(v string) *GetAudioFileInfoResponseBodyResult {
	s.FileId = &v
	return s
}

func (s *GetAudioFileInfoResponseBodyResult) SetFileName(v string) *GetAudioFileInfoResponseBodyResult {
	s.FileName = &v
	return s
}

func (s *GetAudioFileInfoResponseBodyResult) SetFileSize(v int64) *GetAudioFileInfoResponseBodyResult {
	s.FileSize = &v
	return s
}

type GetAudioFileInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAudioFileInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAudioFileInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAudioFileInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAudioFileInfoResponse) SetHeaders(v map[string]*string) *GetAudioFileInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAudioFileInfoResponse) SetStatusCode(v int32) *GetAudioFileInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAudioFileInfoResponse) SetBody(v *GetAudioFileInfoResponseBody) *GetAudioFileInfoResponse {
	s.Body = v
	return s
}

type GetCustomerInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCustomerInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetCustomerInfoHeaders) SetCommonHeaders(v map[string]*string) *GetCustomerInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCustomerInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetCustomerInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCustomerInfoRequest struct {
	// This parameter is required.
	CustomerId *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
}

func (s GetCustomerInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerInfoRequest) GoString() string {
	return s.String()
}

func (s *GetCustomerInfoRequest) SetCustomerId(v string) *GetCustomerInfoRequest {
	s.CustomerId = &v
	return s
}

type GetCustomerInfoResponseBody struct {
	Result *GetCustomerInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetCustomerInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomerInfoResponseBody) SetResult(v *GetCustomerInfoResponseBodyResult) *GetCustomerInfoResponseBody {
	s.Result = v
	return s
}

type GetCustomerInfoResponseBodyResult struct {
	CreateAt    *string `json:"createAt,omitempty" xml:"createAt,omitempty"`
	Id          *string `json:"id,omitempty" xml:"id,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	OwnerUserId *string `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	TeamCode    *string `json:"teamCode,omitempty" xml:"teamCode,omitempty"`
}

func (s GetCustomerInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetCustomerInfoResponseBodyResult) SetCreateAt(v string) *GetCustomerInfoResponseBodyResult {
	s.CreateAt = &v
	return s
}

func (s *GetCustomerInfoResponseBodyResult) SetId(v string) *GetCustomerInfoResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetCustomerInfoResponseBodyResult) SetName(v string) *GetCustomerInfoResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetCustomerInfoResponseBodyResult) SetOwnerUserId(v string) *GetCustomerInfoResponseBodyResult {
	s.OwnerUserId = &v
	return s
}

func (s *GetCustomerInfoResponseBodyResult) SetTeamCode(v string) *GetCustomerInfoResponseBodyResult {
	s.TeamCode = &v
	return s
}

type GetCustomerInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomerInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomerInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerInfoResponse) GoString() string {
	return s.String()
}

func (s *GetCustomerInfoResponse) SetHeaders(v map[string]*string) *GetCustomerInfoResponse {
	s.Headers = v
	return s
}

func (s *GetCustomerInfoResponse) SetStatusCode(v int32) *GetCustomerInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomerInfoResponse) SetBody(v *GetCustomerInfoResponseBody) *GetCustomerInfoResponse {
	s.Body = v
	return s
}

type GetServiceChapterSummaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetServiceChapterSummaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetServiceChapterSummaryHeaders) GoString() string {
	return s.String()
}

func (s *GetServiceChapterSummaryHeaders) SetCommonHeaders(v map[string]*string) *GetServiceChapterSummaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetServiceChapterSummaryHeaders) SetXAcsDingtalkAccessToken(v string) *GetServiceChapterSummaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetServiceChapterSummaryRequest struct {
	// This parameter is required.
	RecordId *string `json:"recordId,omitempty" xml:"recordId,omitempty"`
}

func (s GetServiceChapterSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceChapterSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetServiceChapterSummaryRequest) SetRecordId(v string) *GetServiceChapterSummaryRequest {
	s.RecordId = &v
	return s
}

type GetServiceChapterSummaryResponseBody struct {
	NextToken  *string                                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Result     []*GetServiceChapterSummaryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int32                                        `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetServiceChapterSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceChapterSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceChapterSummaryResponseBody) SetNextToken(v string) *GetServiceChapterSummaryResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetServiceChapterSummaryResponseBody) SetResult(v []*GetServiceChapterSummaryResponseBodyResult) *GetServiceChapterSummaryResponseBody {
	s.Result = v
	return s
}

func (s *GetServiceChapterSummaryResponseBody) SetTotalCount(v int32) *GetServiceChapterSummaryResponseBody {
	s.TotalCount = &v
	return s
}

type GetServiceChapterSummaryResponseBodyResult struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetServiceChapterSummaryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetServiceChapterSummaryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetServiceChapterSummaryResponseBodyResult) SetContent(v string) *GetServiceChapterSummaryResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetServiceChapterSummaryResponseBodyResult) SetName(v string) *GetServiceChapterSummaryResponseBodyResult {
	s.Name = &v
	return s
}

type GetServiceChapterSummaryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceChapterSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceChapterSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceChapterSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetServiceChapterSummaryResponse) SetHeaders(v map[string]*string) *GetServiceChapterSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetServiceChapterSummaryResponse) SetStatusCode(v int32) *GetServiceChapterSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceChapterSummaryResponse) SetBody(v *GetServiceChapterSummaryResponseBody) *GetServiceChapterSummaryResponse {
	s.Body = v
	return s
}

type GetServiceChatSummaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetServiceChatSummaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetServiceChatSummaryHeaders) GoString() string {
	return s.String()
}

func (s *GetServiceChatSummaryHeaders) SetCommonHeaders(v map[string]*string) *GetServiceChatSummaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetServiceChatSummaryHeaders) SetXAcsDingtalkAccessToken(v string) *GetServiceChatSummaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetServiceChatSummaryRequest struct {
	// This parameter is required.
	RecordId *string `json:"recordId,omitempty" xml:"recordId,omitempty"`
}

func (s GetServiceChatSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceChatSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetServiceChatSummaryRequest) SetRecordId(v string) *GetServiceChatSummaryRequest {
	s.RecordId = &v
	return s
}

type GetServiceChatSummaryResponseBody struct {
	Result *GetServiceChatSummaryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetServiceChatSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceChatSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceChatSummaryResponseBody) SetResult(v *GetServiceChatSummaryResponseBodyResult) *GetServiceChatSummaryResponseBody {
	s.Result = v
	return s
}

type GetServiceChatSummaryResponseBodyResult struct {
	Basic   []*GetServiceChatSummaryResponseBodyResultBasic   `json:"basic,omitempty" xml:"basic,omitempty" type:"Repeated"`
	Product []*GetServiceChatSummaryResponseBodyResultProduct `json:"product,omitempty" xml:"product,omitempty" type:"Repeated"`
}

func (s GetServiceChatSummaryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetServiceChatSummaryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetServiceChatSummaryResponseBodyResult) SetBasic(v []*GetServiceChatSummaryResponseBodyResultBasic) *GetServiceChatSummaryResponseBodyResult {
	s.Basic = v
	return s
}

func (s *GetServiceChatSummaryResponseBodyResult) SetProduct(v []*GetServiceChatSummaryResponseBodyResultProduct) *GetServiceChatSummaryResponseBodyResult {
	s.Product = v
	return s
}

type GetServiceChatSummaryResponseBodyResultBasic struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetServiceChatSummaryResponseBodyResultBasic) String() string {
	return tea.Prettify(s)
}

func (s GetServiceChatSummaryResponseBodyResultBasic) GoString() string {
	return s.String()
}

func (s *GetServiceChatSummaryResponseBodyResultBasic) SetContent(v string) *GetServiceChatSummaryResponseBodyResultBasic {
	s.Content = &v
	return s
}

func (s *GetServiceChatSummaryResponseBodyResultBasic) SetName(v string) *GetServiceChatSummaryResponseBodyResultBasic {
	s.Name = &v
	return s
}

type GetServiceChatSummaryResponseBodyResultProduct struct {
	ItemList []*GetServiceChatSummaryResponseBodyResultProductItemList `json:"itemList,omitempty" xml:"itemList,omitempty" type:"Repeated"`
	Product  *string                                                   `json:"product,omitempty" xml:"product,omitempty"`
}

func (s GetServiceChatSummaryResponseBodyResultProduct) String() string {
	return tea.Prettify(s)
}

func (s GetServiceChatSummaryResponseBodyResultProduct) GoString() string {
	return s.String()
}

func (s *GetServiceChatSummaryResponseBodyResultProduct) SetItemList(v []*GetServiceChatSummaryResponseBodyResultProductItemList) *GetServiceChatSummaryResponseBodyResultProduct {
	s.ItemList = v
	return s
}

func (s *GetServiceChatSummaryResponseBodyResultProduct) SetProduct(v string) *GetServiceChatSummaryResponseBodyResultProduct {
	s.Product = &v
	return s
}

type GetServiceChatSummaryResponseBodyResultProductItemList struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetServiceChatSummaryResponseBodyResultProductItemList) String() string {
	return tea.Prettify(s)
}

func (s GetServiceChatSummaryResponseBodyResultProductItemList) GoString() string {
	return s.String()
}

func (s *GetServiceChatSummaryResponseBodyResultProductItemList) SetContent(v string) *GetServiceChatSummaryResponseBodyResultProductItemList {
	s.Content = &v
	return s
}

func (s *GetServiceChatSummaryResponseBodyResultProductItemList) SetName(v string) *GetServiceChatSummaryResponseBodyResultProductItemList {
	s.Name = &v
	return s
}

type GetServiceChatSummaryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceChatSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceChatSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceChatSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetServiceChatSummaryResponse) SetHeaders(v map[string]*string) *GetServiceChatSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetServiceChatSummaryResponse) SetStatusCode(v int32) *GetServiceChatSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceChatSummaryResponse) SetBody(v *GetServiceChatSummaryResponseBody) *GetServiceChatSummaryResponse {
	s.Body = v
	return s
}

type GetServiceQualityInspectionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetServiceQualityInspectionHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetServiceQualityInspectionHeaders) GoString() string {
	return s.String()
}

func (s *GetServiceQualityInspectionHeaders) SetCommonHeaders(v map[string]*string) *GetServiceQualityInspectionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetServiceQualityInspectionHeaders) SetXAcsDingtalkAccessToken(v string) *GetServiceQualityInspectionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetServiceQualityInspectionRequest struct {
	// This parameter is required.
	RecordId *string `json:"recordId,omitempty" xml:"recordId,omitempty"`
}

func (s GetServiceQualityInspectionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceQualityInspectionRequest) GoString() string {
	return s.String()
}

func (s *GetServiceQualityInspectionRequest) SetRecordId(v string) *GetServiceQualityInspectionRequest {
	s.RecordId = &v
	return s
}

type GetServiceQualityInspectionResponseBody struct {
	Result *GetServiceQualityInspectionResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetServiceQualityInspectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceQualityInspectionResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceQualityInspectionResponseBody) SetResult(v *GetServiceQualityInspectionResponseBodyResult) *GetServiceQualityInspectionResponseBody {
	s.Result = v
	return s
}

type GetServiceQualityInspectionResponseBodyResult struct {
	GroupList []*GetServiceQualityInspectionResponseBodyResultGroupList `json:"groupList,omitempty" xml:"groupList,omitempty" type:"Repeated"`
	Score     *int32                                                    `json:"score,omitempty" xml:"score,omitempty"`
	Summary   *string                                                   `json:"summary,omitempty" xml:"summary,omitempty"`
}

func (s GetServiceQualityInspectionResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetServiceQualityInspectionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetServiceQualityInspectionResponseBodyResult) SetGroupList(v []*GetServiceQualityInspectionResponseBodyResultGroupList) *GetServiceQualityInspectionResponseBodyResult {
	s.GroupList = v
	return s
}

func (s *GetServiceQualityInspectionResponseBodyResult) SetScore(v int32) *GetServiceQualityInspectionResponseBodyResult {
	s.Score = &v
	return s
}

func (s *GetServiceQualityInspectionResponseBodyResult) SetSummary(v string) *GetServiceQualityInspectionResponseBodyResult {
	s.Summary = &v
	return s
}

type GetServiceQualityInspectionResponseBodyResultGroupList struct {
	ItemList []*GetServiceQualityInspectionResponseBodyResultGroupListItemList `json:"itemList,omitempty" xml:"itemList,omitempty" type:"Repeated"`
	Name     *string                                                           `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetServiceQualityInspectionResponseBodyResultGroupList) String() string {
	return tea.Prettify(s)
}

func (s GetServiceQualityInspectionResponseBodyResultGroupList) GoString() string {
	return s.String()
}

func (s *GetServiceQualityInspectionResponseBodyResultGroupList) SetItemList(v []*GetServiceQualityInspectionResponseBodyResultGroupListItemList) *GetServiceQualityInspectionResponseBodyResultGroupList {
	s.ItemList = v
	return s
}

func (s *GetServiceQualityInspectionResponseBodyResultGroupList) SetName(v string) *GetServiceQualityInspectionResponseBodyResultGroupList {
	s.Name = &v
	return s
}

type GetServiceQualityInspectionResponseBodyResultGroupListItemList struct {
	FlowName *string `json:"flowName,omitempty" xml:"flowName,omitempty"`
	IsHit    *string `json:"isHit,omitempty" xml:"isHit,omitempty"`
	Reason   *string `json:"reason,omitempty" xml:"reason,omitempty"`
	Score    *int32  `json:"score,omitempty" xml:"score,omitempty"`
	Script   *string `json:"script,omitempty" xml:"script,omitempty"`
}

func (s GetServiceQualityInspectionResponseBodyResultGroupListItemList) String() string {
	return tea.Prettify(s)
}

func (s GetServiceQualityInspectionResponseBodyResultGroupListItemList) GoString() string {
	return s.String()
}

func (s *GetServiceQualityInspectionResponseBodyResultGroupListItemList) SetFlowName(v string) *GetServiceQualityInspectionResponseBodyResultGroupListItemList {
	s.FlowName = &v
	return s
}

func (s *GetServiceQualityInspectionResponseBodyResultGroupListItemList) SetIsHit(v string) *GetServiceQualityInspectionResponseBodyResultGroupListItemList {
	s.IsHit = &v
	return s
}

func (s *GetServiceQualityInspectionResponseBodyResultGroupListItemList) SetReason(v string) *GetServiceQualityInspectionResponseBodyResultGroupListItemList {
	s.Reason = &v
	return s
}

func (s *GetServiceQualityInspectionResponseBodyResultGroupListItemList) SetScore(v int32) *GetServiceQualityInspectionResponseBodyResultGroupListItemList {
	s.Score = &v
	return s
}

func (s *GetServiceQualityInspectionResponseBodyResultGroupListItemList) SetScript(v string) *GetServiceQualityInspectionResponseBodyResultGroupListItemList {
	s.Script = &v
	return s
}

type GetServiceQualityInspectionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceQualityInspectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceQualityInspectionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceQualityInspectionResponse) GoString() string {
	return s.String()
}

func (s *GetServiceQualityInspectionResponse) SetHeaders(v map[string]*string) *GetServiceQualityInspectionResponse {
	s.Headers = v
	return s
}

func (s *GetServiceQualityInspectionResponse) SetStatusCode(v int32) *GetServiceQualityInspectionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceQualityInspectionResponse) SetBody(v *GetServiceQualityInspectionResponseBody) *GetServiceQualityInspectionResponse {
	s.Body = v
	return s
}

type GetServiceRecordTranscriptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetServiceRecordTranscriptHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetServiceRecordTranscriptHeaders) GoString() string {
	return s.String()
}

func (s *GetServiceRecordTranscriptHeaders) SetCommonHeaders(v map[string]*string) *GetServiceRecordTranscriptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetServiceRecordTranscriptHeaders) SetXAcsDingtalkAccessToken(v string) *GetServiceRecordTranscriptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetServiceRecordTranscriptRequest struct {
	// This parameter is required.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s GetServiceRecordTranscriptRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceRecordTranscriptRequest) GoString() string {
	return s.String()
}

func (s *GetServiceRecordTranscriptRequest) SetId(v string) *GetServiceRecordTranscriptRequest {
	s.Id = &v
	return s
}

type GetServiceRecordTranscriptResponseBody struct {
	Result *GetServiceRecordTranscriptResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetServiceRecordTranscriptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceRecordTranscriptResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceRecordTranscriptResponseBody) SetResult(v *GetServiceRecordTranscriptResponseBodyResult) *GetServiceRecordTranscriptResponseBody {
	s.Result = v
	return s
}

type GetServiceRecordTranscriptResponseBodyResult struct {
	AudionText *GetServiceRecordTranscriptResponseBodyResultAudionText `json:"audionText,omitempty" xml:"audionText,omitempty" type:"Struct"`
	Speaker    *GetServiceRecordTranscriptResponseBodyResultSpeaker    `json:"speaker,omitempty" xml:"speaker,omitempty" type:"Struct"`
}

func (s GetServiceRecordTranscriptResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetServiceRecordTranscriptResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetServiceRecordTranscriptResponseBodyResult) SetAudionText(v *GetServiceRecordTranscriptResponseBodyResultAudionText) *GetServiceRecordTranscriptResponseBodyResult {
	s.AudionText = v
	return s
}

func (s *GetServiceRecordTranscriptResponseBodyResult) SetSpeaker(v *GetServiceRecordTranscriptResponseBodyResultSpeaker) *GetServiceRecordTranscriptResponseBodyResult {
	s.Speaker = v
	return s
}

type GetServiceRecordTranscriptResponseBodyResultAudionText struct {
	DataList []*GetServiceRecordTranscriptResponseBodyResultAudionTextDataList `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	Status   *string                                                           `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetServiceRecordTranscriptResponseBodyResultAudionText) String() string {
	return tea.Prettify(s)
}

func (s GetServiceRecordTranscriptResponseBodyResultAudionText) GoString() string {
	return s.String()
}

func (s *GetServiceRecordTranscriptResponseBodyResultAudionText) SetDataList(v []*GetServiceRecordTranscriptResponseBodyResultAudionTextDataList) *GetServiceRecordTranscriptResponseBodyResultAudionText {
	s.DataList = v
	return s
}

func (s *GetServiceRecordTranscriptResponseBodyResultAudionText) SetStatus(v string) *GetServiceRecordTranscriptResponseBodyResultAudionText {
	s.Status = &v
	return s
}

type GetServiceRecordTranscriptResponseBodyResultAudionTextDataList struct {
	Channel   *string `json:"channel,omitempty" xml:"channel,omitempty"`
	EndTime   *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Text      *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s GetServiceRecordTranscriptResponseBodyResultAudionTextDataList) String() string {
	return tea.Prettify(s)
}

func (s GetServiceRecordTranscriptResponseBodyResultAudionTextDataList) GoString() string {
	return s.String()
}

func (s *GetServiceRecordTranscriptResponseBodyResultAudionTextDataList) SetChannel(v string) *GetServiceRecordTranscriptResponseBodyResultAudionTextDataList {
	s.Channel = &v
	return s
}

func (s *GetServiceRecordTranscriptResponseBodyResultAudionTextDataList) SetEndTime(v string) *GetServiceRecordTranscriptResponseBodyResultAudionTextDataList {
	s.EndTime = &v
	return s
}

func (s *GetServiceRecordTranscriptResponseBodyResultAudionTextDataList) SetStartTime(v string) *GetServiceRecordTranscriptResponseBodyResultAudionTextDataList {
	s.StartTime = &v
	return s
}

func (s *GetServiceRecordTranscriptResponseBodyResultAudionTextDataList) SetText(v string) *GetServiceRecordTranscriptResponseBodyResultAudionTextDataList {
	s.Text = &v
	return s
}

type GetServiceRecordTranscriptResponseBodyResultSpeaker struct {
	DataList []*GetServiceRecordTranscriptResponseBodyResultSpeakerDataList `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	Status   *string                                                        `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetServiceRecordTranscriptResponseBodyResultSpeaker) String() string {
	return tea.Prettify(s)
}

func (s GetServiceRecordTranscriptResponseBodyResultSpeaker) GoString() string {
	return s.String()
}

func (s *GetServiceRecordTranscriptResponseBodyResultSpeaker) SetDataList(v []*GetServiceRecordTranscriptResponseBodyResultSpeakerDataList) *GetServiceRecordTranscriptResponseBodyResultSpeaker {
	s.DataList = v
	return s
}

func (s *GetServiceRecordTranscriptResponseBodyResultSpeaker) SetStatus(v string) *GetServiceRecordTranscriptResponseBodyResultSpeaker {
	s.Status = &v
	return s
}

type GetServiceRecordTranscriptResponseBodyResultSpeakerDataList struct {
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	Role    *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s GetServiceRecordTranscriptResponseBodyResultSpeakerDataList) String() string {
	return tea.Prettify(s)
}

func (s GetServiceRecordTranscriptResponseBodyResultSpeakerDataList) GoString() string {
	return s.String()
}

func (s *GetServiceRecordTranscriptResponseBodyResultSpeakerDataList) SetChannel(v string) *GetServiceRecordTranscriptResponseBodyResultSpeakerDataList {
	s.Channel = &v
	return s
}

func (s *GetServiceRecordTranscriptResponseBodyResultSpeakerDataList) SetRole(v string) *GetServiceRecordTranscriptResponseBodyResultSpeakerDataList {
	s.Role = &v
	return s
}

type GetServiceRecordTranscriptResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceRecordTranscriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceRecordTranscriptResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceRecordTranscriptResponse) GoString() string {
	return s.String()
}

func (s *GetServiceRecordTranscriptResponse) SetHeaders(v map[string]*string) *GetServiceRecordTranscriptResponse {
	s.Headers = v
	return s
}

func (s *GetServiceRecordTranscriptResponse) SetStatusCode(v int32) *GetServiceRecordTranscriptResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceRecordTranscriptResponse) SetBody(v *GetServiceRecordTranscriptResponseBody) *GetServiceRecordTranscriptResponse {
	s.Body = v
	return s
}

type GetTranscriptSummaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTranscriptSummaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTranscriptSummaryHeaders) GoString() string {
	return s.String()
}

func (s *GetTranscriptSummaryHeaders) SetCommonHeaders(v map[string]*string) *GetTranscriptSummaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTranscriptSummaryHeaders) SetXAcsDingtalkAccessToken(v string) *GetTranscriptSummaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTranscriptSummaryRequest struct {
	// This parameter is required.
	DeviceType *string `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
	// This parameter is required.
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
}

func (s GetTranscriptSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTranscriptSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetTranscriptSummaryRequest) SetDeviceType(v string) *GetTranscriptSummaryRequest {
	s.DeviceType = &v
	return s
}

func (s *GetTranscriptSummaryRequest) SetFileId(v string) *GetTranscriptSummaryRequest {
	s.FileId = &v
	return s
}

type GetTranscriptSummaryResponseBody struct {
	Result *GetTranscriptSummaryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetTranscriptSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTranscriptSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetTranscriptSummaryResponseBody) SetResult(v *GetTranscriptSummaryResponseBodyResult) *GetTranscriptSummaryResponseBody {
	s.Result = v
	return s
}

type GetTranscriptSummaryResponseBodyResult struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s GetTranscriptSummaryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetTranscriptSummaryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetTranscriptSummaryResponseBodyResult) SetContent(v string) *GetTranscriptSummaryResponseBodyResult {
	s.Content = &v
	return s
}

type GetTranscriptSummaryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTranscriptSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTranscriptSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTranscriptSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetTranscriptSummaryResponse) SetHeaders(v map[string]*string) *GetTranscriptSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetTranscriptSummaryResponse) SetStatusCode(v int32) *GetTranscriptSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTranscriptSummaryResponse) SetBody(v *GetTranscriptSummaryResponseBody) *GetTranscriptSummaryResponse {
	s.Body = v
	return s
}

type ListCustomerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListCustomerHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListCustomerHeaders) GoString() string {
	return s.String()
}

func (s *ListCustomerHeaders) SetCommonHeaders(v map[string]*string) *ListCustomerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListCustomerHeaders) SetXAcsDingtalkAccessToken(v string) *ListCustomerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListCustomerRequest struct {
	EndTime     *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	MaxResults  *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken   *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OwnerUserId *string `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	StartTime   *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// This parameter is required.
	TeamCode *string `json:"teamCode,omitempty" xml:"teamCode,omitempty"`
}

func (s ListCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCustomerRequest) GoString() string {
	return s.String()
}

func (s *ListCustomerRequest) SetEndTime(v int64) *ListCustomerRequest {
	s.EndTime = &v
	return s
}

func (s *ListCustomerRequest) SetMaxResults(v int32) *ListCustomerRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCustomerRequest) SetNextToken(v string) *ListCustomerRequest {
	s.NextToken = &v
	return s
}

func (s *ListCustomerRequest) SetOwnerUserId(v string) *ListCustomerRequest {
	s.OwnerUserId = &v
	return s
}

func (s *ListCustomerRequest) SetStartTime(v int64) *ListCustomerRequest {
	s.StartTime = &v
	return s
}

func (s *ListCustomerRequest) SetTeamCode(v string) *ListCustomerRequest {
	s.TeamCode = &v
	return s
}

type ListCustomerResponseBody struct {
	NextToken  *string                           `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Result     []*ListCustomerResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int32                            `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomerResponseBody) SetNextToken(v string) *ListCustomerResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCustomerResponseBody) SetResult(v []*ListCustomerResponseBodyResult) *ListCustomerResponseBody {
	s.Result = v
	return s
}

func (s *ListCustomerResponseBody) SetTotalCount(v int32) *ListCustomerResponseBody {
	s.TotalCount = &v
	return s
}

type ListCustomerResponseBodyResult struct {
	CreateAt    *string `json:"createAt,omitempty" xml:"createAt,omitempty"`
	Id          *string `json:"id,omitempty" xml:"id,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	OwnerUserId *string `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	TeamCode    *string `json:"teamCode,omitempty" xml:"teamCode,omitempty"`
}

func (s ListCustomerResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListCustomerResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCustomerResponseBodyResult) SetCreateAt(v string) *ListCustomerResponseBodyResult {
	s.CreateAt = &v
	return s
}

func (s *ListCustomerResponseBodyResult) SetId(v string) *ListCustomerResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListCustomerResponseBodyResult) SetName(v string) *ListCustomerResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListCustomerResponseBodyResult) SetOwnerUserId(v string) *ListCustomerResponseBodyResult {
	s.OwnerUserId = &v
	return s
}

func (s *ListCustomerResponseBodyResult) SetTeamCode(v string) *ListCustomerResponseBodyResult {
	s.TeamCode = &v
	return s
}

type ListCustomerResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCustomerResponse) GoString() string {
	return s.String()
}

func (s *ListCustomerResponse) SetHeaders(v map[string]*string) *ListCustomerResponse {
	s.Headers = v
	return s
}

func (s *ListCustomerResponse) SetStatusCode(v int32) *ListCustomerResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomerResponse) SetBody(v *ListCustomerResponseBody) *ListCustomerResponse {
	s.Body = v
	return s
}

type ListServiceRecordHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListServiceRecordHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListServiceRecordHeaders) GoString() string {
	return s.String()
}

func (s *ListServiceRecordHeaders) SetCommonHeaders(v map[string]*string) *ListServiceRecordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListServiceRecordHeaders) SetXAcsDingtalkAccessToken(v string) *ListServiceRecordHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListServiceRecordRequest struct {
	EndTime    *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	StartTime  *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	TeamCode   *string `json:"teamCode,omitempty" xml:"teamCode,omitempty"`
	UserId     *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListServiceRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceRecordRequest) GoString() string {
	return s.String()
}

func (s *ListServiceRecordRequest) SetEndTime(v int64) *ListServiceRecordRequest {
	s.EndTime = &v
	return s
}

func (s *ListServiceRecordRequest) SetMaxResults(v int32) *ListServiceRecordRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceRecordRequest) SetNextToken(v string) *ListServiceRecordRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceRecordRequest) SetStartTime(v int64) *ListServiceRecordRequest {
	s.StartTime = &v
	return s
}

func (s *ListServiceRecordRequest) SetTeamCode(v string) *ListServiceRecordRequest {
	s.TeamCode = &v
	return s
}

func (s *ListServiceRecordRequest) SetUserId(v string) *ListServiceRecordRequest {
	s.UserId = &v
	return s
}

type ListServiceRecordResponseBody struct {
	NextToken  *string                                `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Result     []*ListServiceRecordResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int32                                 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListServiceRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceRecordResponseBody) SetNextToken(v string) *ListServiceRecordResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceRecordResponseBody) SetResult(v []*ListServiceRecordResponseBodyResult) *ListServiceRecordResponseBody {
	s.Result = v
	return s
}

func (s *ListServiceRecordResponseBody) SetTotalCount(v int32) *ListServiceRecordResponseBody {
	s.TotalCount = &v
	return s
}

type ListServiceRecordResponseBodyResult struct {
	CustomerId     *string                                  `json:"customerId,omitempty" xml:"customerId,omitempty"`
	DeviceSn       *string                                  `json:"deviceSn,omitempty" xml:"deviceSn,omitempty"`
	Duration       *string                                  `json:"duration,omitempty" xml:"duration,omitempty"`
	EndTimestamp   *int64                                   `json:"endTimestamp,omitempty" xml:"endTimestamp,omitempty"`
	RecordId       *string                                  `json:"recordId,omitempty" xml:"recordId,omitempty"`
	StartTimestamp *int64                                   `json:"startTimestamp,omitempty" xml:"startTimestamp,omitempty"`
	Team           *ListServiceRecordResponseBodyResultTeam `json:"team,omitempty" xml:"team,omitempty" type:"Struct"`
	User           *ListServiceRecordResponseBodyResultUser `json:"user,omitempty" xml:"user,omitempty" type:"Struct"`
}

func (s ListServiceRecordResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListServiceRecordResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListServiceRecordResponseBodyResult) SetCustomerId(v string) *ListServiceRecordResponseBodyResult {
	s.CustomerId = &v
	return s
}

func (s *ListServiceRecordResponseBodyResult) SetDeviceSn(v string) *ListServiceRecordResponseBodyResult {
	s.DeviceSn = &v
	return s
}

func (s *ListServiceRecordResponseBodyResult) SetDuration(v string) *ListServiceRecordResponseBodyResult {
	s.Duration = &v
	return s
}

func (s *ListServiceRecordResponseBodyResult) SetEndTimestamp(v int64) *ListServiceRecordResponseBodyResult {
	s.EndTimestamp = &v
	return s
}

func (s *ListServiceRecordResponseBodyResult) SetRecordId(v string) *ListServiceRecordResponseBodyResult {
	s.RecordId = &v
	return s
}

func (s *ListServiceRecordResponseBodyResult) SetStartTimestamp(v int64) *ListServiceRecordResponseBodyResult {
	s.StartTimestamp = &v
	return s
}

func (s *ListServiceRecordResponseBodyResult) SetTeam(v *ListServiceRecordResponseBodyResultTeam) *ListServiceRecordResponseBodyResult {
	s.Team = v
	return s
}

func (s *ListServiceRecordResponseBodyResult) SetUser(v *ListServiceRecordResponseBodyResultUser) *ListServiceRecordResponseBodyResult {
	s.User = v
	return s
}

type ListServiceRecordResponseBodyResultTeam struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListServiceRecordResponseBodyResultTeam) String() string {
	return tea.Prettify(s)
}

func (s ListServiceRecordResponseBodyResultTeam) GoString() string {
	return s.String()
}

func (s *ListServiceRecordResponseBodyResultTeam) SetCode(v string) *ListServiceRecordResponseBodyResultTeam {
	s.Code = &v
	return s
}

func (s *ListServiceRecordResponseBodyResultTeam) SetName(v string) *ListServiceRecordResponseBodyResultTeam {
	s.Name = &v
	return s
}

type ListServiceRecordResponseBodyResultUser struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListServiceRecordResponseBodyResultUser) String() string {
	return tea.Prettify(s)
}

func (s ListServiceRecordResponseBodyResultUser) GoString() string {
	return s.String()
}

func (s *ListServiceRecordResponseBodyResultUser) SetName(v string) *ListServiceRecordResponseBodyResultUser {
	s.Name = &v
	return s
}

func (s *ListServiceRecordResponseBodyResultUser) SetUserId(v string) *ListServiceRecordResponseBodyResultUser {
	s.UserId = &v
	return s
}

type ListServiceRecordResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceRecordResponse) GoString() string {
	return s.String()
}

func (s *ListServiceRecordResponse) SetHeaders(v map[string]*string) *ListServiceRecordResponse {
	s.Headers = v
	return s
}

func (s *ListServiceRecordResponse) SetStatusCode(v int32) *ListServiceRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceRecordResponse) SetBody(v *ListServiceRecordResponseBody) *ListServiceRecordResponse {
	s.Body = v
	return s
}

type ListServiceTodoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListServiceTodoHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListServiceTodoHeaders) GoString() string {
	return s.String()
}

func (s *ListServiceTodoHeaders) SetCommonHeaders(v map[string]*string) *ListServiceTodoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListServiceTodoHeaders) SetXAcsDingtalkAccessToken(v string) *ListServiceTodoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListServiceTodoRequest struct {
	// This parameter is required.
	RecordId *string `json:"recordId,omitempty" xml:"recordId,omitempty"`
}

func (s ListServiceTodoRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceTodoRequest) GoString() string {
	return s.String()
}

func (s *ListServiceTodoRequest) SetRecordId(v string) *ListServiceTodoRequest {
	s.RecordId = &v
	return s
}

type ListServiceTodoResponseBody struct {
	Result []*ListServiceTodoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListServiceTodoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceTodoResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceTodoResponseBody) SetResult(v []*ListServiceTodoResponseBodyResult) *ListServiceTodoResponseBody {
	s.Result = v
	return s
}

type ListServiceTodoResponseBodyResult struct {
	Creator        *string                                       `json:"creator,omitempty" xml:"creator,omitempty"`
	DingTodoId     *string                                       `json:"dingTodoId,omitempty" xml:"dingTodoId,omitempty"`
	Executors      []*ListServiceTodoResponseBodyResultExecutors `json:"executors,omitempty" xml:"executors,omitempty" type:"Repeated"`
	Finished       *bool                                         `json:"finished,omitempty" xml:"finished,omitempty"`
	PlanFinishDate *int64                                        `json:"planFinishDate,omitempty" xml:"planFinishDate,omitempty"`
	TodoContent    *string                                       `json:"todoContent,omitempty" xml:"todoContent,omitempty"`
	Uuid           *string                                       `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s ListServiceTodoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListServiceTodoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListServiceTodoResponseBodyResult) SetCreator(v string) *ListServiceTodoResponseBodyResult {
	s.Creator = &v
	return s
}

func (s *ListServiceTodoResponseBodyResult) SetDingTodoId(v string) *ListServiceTodoResponseBodyResult {
	s.DingTodoId = &v
	return s
}

func (s *ListServiceTodoResponseBodyResult) SetExecutors(v []*ListServiceTodoResponseBodyResultExecutors) *ListServiceTodoResponseBodyResult {
	s.Executors = v
	return s
}

func (s *ListServiceTodoResponseBodyResult) SetFinished(v bool) *ListServiceTodoResponseBodyResult {
	s.Finished = &v
	return s
}

func (s *ListServiceTodoResponseBodyResult) SetPlanFinishDate(v int64) *ListServiceTodoResponseBodyResult {
	s.PlanFinishDate = &v
	return s
}

func (s *ListServiceTodoResponseBodyResult) SetTodoContent(v string) *ListServiceTodoResponseBodyResult {
	s.TodoContent = &v
	return s
}

func (s *ListServiceTodoResponseBodyResult) SetUuid(v string) *ListServiceTodoResponseBodyResult {
	s.Uuid = &v
	return s
}

type ListServiceTodoResponseBodyResultExecutors struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListServiceTodoResponseBodyResultExecutors) String() string {
	return tea.Prettify(s)
}

func (s ListServiceTodoResponseBodyResultExecutors) GoString() string {
	return s.String()
}

func (s *ListServiceTodoResponseBodyResultExecutors) SetName(v string) *ListServiceTodoResponseBodyResultExecutors {
	s.Name = &v
	return s
}

func (s *ListServiceTodoResponseBodyResultExecutors) SetUserId(v string) *ListServiceTodoResponseBodyResultExecutors {
	s.UserId = &v
	return s
}

type ListServiceTodoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceTodoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceTodoResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceTodoResponse) GoString() string {
	return s.String()
}

func (s *ListServiceTodoResponse) SetHeaders(v map[string]*string) *ListServiceTodoResponse {
	s.Headers = v
	return s
}

func (s *ListServiceTodoResponse) SetStatusCode(v int32) *ListServiceTodoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceTodoResponse) SetBody(v *ListServiceTodoResponseBody) *ListServiceTodoResponse {
	s.Body = v
	return s
}

type ListTeamHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListTeamHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListTeamHeaders) GoString() string {
	return s.String()
}

func (s *ListTeamHeaders) SetCommonHeaders(v map[string]*string) *ListTeamHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListTeamHeaders) SetXAcsDingtalkAccessToken(v string) *ListTeamHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListTeamRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListTeamRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTeamRequest) GoString() string {
	return s.String()
}

func (s *ListTeamRequest) SetMaxResults(v int32) *ListTeamRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTeamRequest) SetNextToken(v string) *ListTeamRequest {
	s.NextToken = &v
	return s
}

type ListTeamResponseBody struct {
	NextToken  *string                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Result     []*ListTeamResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int32                        `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListTeamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTeamResponseBody) GoString() string {
	return s.String()
}

func (s *ListTeamResponseBody) SetNextToken(v string) *ListTeamResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTeamResponseBody) SetResult(v []*ListTeamResponseBodyResult) *ListTeamResponseBody {
	s.Result = v
	return s
}

func (s *ListTeamResponseBody) SetTotalCount(v int32) *ListTeamResponseBody {
	s.TotalCount = &v
	return s
}

type ListTeamResponseBodyResult struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListTeamResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListTeamResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListTeamResponseBodyResult) SetCode(v string) *ListTeamResponseBodyResult {
	s.Code = &v
	return s
}

func (s *ListTeamResponseBodyResult) SetName(v string) *ListTeamResponseBodyResult {
	s.Name = &v
	return s
}

type ListTeamResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTeamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTeamResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTeamResponse) GoString() string {
	return s.String()
}

func (s *ListTeamResponse) SetHeaders(v map[string]*string) *ListTeamResponse {
	s.Headers = v
	return s
}

func (s *ListTeamResponse) SetStatusCode(v int32) *ListTeamResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTeamResponse) SetBody(v *ListTeamResponseBody) *ListTeamResponse {
	s.Body = v
	return s
}

type QueryAsrTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAsrTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAsrTaskHeaders) GoString() string {
	return s.String()
}

func (s *QueryAsrTaskHeaders) SetCommonHeaders(v map[string]*string) *QueryAsrTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAsrTaskHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAsrTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAsrTaskRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	TaskId  *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryAsrTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAsrTaskRequest) GoString() string {
	return s.String()
}

func (s *QueryAsrTaskRequest) SetMaxResults(v int32) *QueryAsrTaskRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryAsrTaskRequest) SetNextToken(v string) *QueryAsrTaskRequest {
	s.NextToken = &v
	return s
}

func (s *QueryAsrTaskRequest) SetTaskId(v string) *QueryAsrTaskRequest {
	s.TaskId = &v
	return s
}

func (s *QueryAsrTaskRequest) SetUnionId(v string) *QueryAsrTaskRequest {
	s.UnionId = &v
	return s
}

type QueryAsrTaskResponseBody struct {
	ErrorCode *string                         `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                         `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result    *QueryAsrTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success   *bool                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryAsrTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAsrTaskResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAsrTaskResponseBody) SetErrorCode(v string) *QueryAsrTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryAsrTaskResponseBody) SetErrorMsg(v string) *QueryAsrTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryAsrTaskResponseBody) SetResult(v *QueryAsrTaskResponseBodyResult) *QueryAsrTaskResponseBody {
	s.Result = v
	return s
}

func (s *QueryAsrTaskResponseBody) SetSuccess(v bool) *QueryAsrTaskResponseBody {
	s.Success = &v
	return s
}

type QueryAsrTaskResponseBodyResult struct {
	BizKey     *string                                   `json:"bizKey,omitempty" xml:"bizKey,omitempty"`
	NextToken  *string                                   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ResultInfo *QueryAsrTaskResponseBodyResultResultInfo `json:"resultInfo,omitempty" xml:"resultInfo,omitempty" type:"Struct"`
	TaskId     *string                                   `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TaskStatus *string                                   `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s QueryAsrTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryAsrTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryAsrTaskResponseBodyResult) SetBizKey(v string) *QueryAsrTaskResponseBodyResult {
	s.BizKey = &v
	return s
}

func (s *QueryAsrTaskResponseBodyResult) SetNextToken(v string) *QueryAsrTaskResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *QueryAsrTaskResponseBodyResult) SetResultInfo(v *QueryAsrTaskResponseBodyResultResultInfo) *QueryAsrTaskResponseBodyResult {
	s.ResultInfo = v
	return s
}

func (s *QueryAsrTaskResponseBodyResult) SetTaskId(v string) *QueryAsrTaskResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *QueryAsrTaskResponseBodyResult) SetTaskStatus(v string) *QueryAsrTaskResponseBodyResult {
	s.TaskStatus = &v
	return s
}

type QueryAsrTaskResponseBodyResultResultInfo struct {
	ParagraphList []*QueryAsrTaskResponseBodyResultResultInfoParagraphList `json:"paragraphList,omitempty" xml:"paragraphList,omitempty" type:"Repeated"`
}

func (s QueryAsrTaskResponseBodyResultResultInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryAsrTaskResponseBodyResultResultInfo) GoString() string {
	return s.String()
}

func (s *QueryAsrTaskResponseBodyResultResultInfo) SetParagraphList(v []*QueryAsrTaskResponseBodyResultResultInfoParagraphList) *QueryAsrTaskResponseBodyResultResultInfo {
	s.ParagraphList = v
	return s
}

type QueryAsrTaskResponseBodyResultResultInfoParagraphList struct {
	EndTime      *int64                                                               `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Paragraph    *string                                                              `json:"paragraph,omitempty" xml:"paragraph,omitempty"`
	SentenceList []*QueryAsrTaskResponseBodyResultResultInfoParagraphListSentenceList `json:"sentenceList,omitempty" xml:"sentenceList,omitempty" type:"Repeated"`
	SpeakerId    *string                                                              `json:"speakerId,omitempty" xml:"speakerId,omitempty"`
	StartTime    *int64                                                               `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s QueryAsrTaskResponseBodyResultResultInfoParagraphList) String() string {
	return tea.Prettify(s)
}

func (s QueryAsrTaskResponseBodyResultResultInfoParagraphList) GoString() string {
	return s.String()
}

func (s *QueryAsrTaskResponseBodyResultResultInfoParagraphList) SetEndTime(v int64) *QueryAsrTaskResponseBodyResultResultInfoParagraphList {
	s.EndTime = &v
	return s
}

func (s *QueryAsrTaskResponseBodyResultResultInfoParagraphList) SetParagraph(v string) *QueryAsrTaskResponseBodyResultResultInfoParagraphList {
	s.Paragraph = &v
	return s
}

func (s *QueryAsrTaskResponseBodyResultResultInfoParagraphList) SetSentenceList(v []*QueryAsrTaskResponseBodyResultResultInfoParagraphListSentenceList) *QueryAsrTaskResponseBodyResultResultInfoParagraphList {
	s.SentenceList = v
	return s
}

func (s *QueryAsrTaskResponseBodyResultResultInfoParagraphList) SetSpeakerId(v string) *QueryAsrTaskResponseBodyResultResultInfoParagraphList {
	s.SpeakerId = &v
	return s
}

func (s *QueryAsrTaskResponseBodyResultResultInfoParagraphList) SetStartTime(v int64) *QueryAsrTaskResponseBodyResultResultInfoParagraphList {
	s.StartTime = &v
	return s
}

type QueryAsrTaskResponseBodyResultResultInfoParagraphListSentenceList struct {
	EndTime   *int64                                                                       `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Sentence  *string                                                                      `json:"sentence,omitempty" xml:"sentence,omitempty"`
	StartTime *int64                                                                       `json:"startTime,omitempty" xml:"startTime,omitempty"`
	WordList  []*QueryAsrTaskResponseBodyResultResultInfoParagraphListSentenceListWordList `json:"wordList,omitempty" xml:"wordList,omitempty" type:"Repeated"`
}

func (s QueryAsrTaskResponseBodyResultResultInfoParagraphListSentenceList) String() string {
	return tea.Prettify(s)
}

func (s QueryAsrTaskResponseBodyResultResultInfoParagraphListSentenceList) GoString() string {
	return s.String()
}

func (s *QueryAsrTaskResponseBodyResultResultInfoParagraphListSentenceList) SetEndTime(v int64) *QueryAsrTaskResponseBodyResultResultInfoParagraphListSentenceList {
	s.EndTime = &v
	return s
}

func (s *QueryAsrTaskResponseBodyResultResultInfoParagraphListSentenceList) SetSentence(v string) *QueryAsrTaskResponseBodyResultResultInfoParagraphListSentenceList {
	s.Sentence = &v
	return s
}

func (s *QueryAsrTaskResponseBodyResultResultInfoParagraphListSentenceList) SetStartTime(v int64) *QueryAsrTaskResponseBodyResultResultInfoParagraphListSentenceList {
	s.StartTime = &v
	return s
}

func (s *QueryAsrTaskResponseBodyResultResultInfoParagraphListSentenceList) SetWordList(v []*QueryAsrTaskResponseBodyResultResultInfoParagraphListSentenceListWordList) *QueryAsrTaskResponseBodyResultResultInfoParagraphListSentenceList {
	s.WordList = v
	return s
}

type QueryAsrTaskResponseBodyResultResultInfoParagraphListSentenceListWordList struct {
	EndTime   *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Text      *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s QueryAsrTaskResponseBodyResultResultInfoParagraphListSentenceListWordList) String() string {
	return tea.Prettify(s)
}

func (s QueryAsrTaskResponseBodyResultResultInfoParagraphListSentenceListWordList) GoString() string {
	return s.String()
}

func (s *QueryAsrTaskResponseBodyResultResultInfoParagraphListSentenceListWordList) SetEndTime(v int64) *QueryAsrTaskResponseBodyResultResultInfoParagraphListSentenceListWordList {
	s.EndTime = &v
	return s
}

func (s *QueryAsrTaskResponseBodyResultResultInfoParagraphListSentenceListWordList) SetStartTime(v int64) *QueryAsrTaskResponseBodyResultResultInfoParagraphListSentenceListWordList {
	s.StartTime = &v
	return s
}

func (s *QueryAsrTaskResponseBodyResultResultInfoParagraphListSentenceListWordList) SetText(v string) *QueryAsrTaskResponseBodyResultResultInfoParagraphListSentenceListWordList {
	s.Text = &v
	return s
}

type QueryAsrTaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAsrTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAsrTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAsrTaskResponse) GoString() string {
	return s.String()
}

func (s *QueryAsrTaskResponse) SetHeaders(v map[string]*string) *QueryAsrTaskResponse {
	s.Headers = v
	return s
}

func (s *QueryAsrTaskResponse) SetStatusCode(v int32) *QueryAsrTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAsrTaskResponse) SetBody(v *QueryAsrTaskResponseBody) *QueryAsrTaskResponse {
	s.Body = v
	return s
}

type QueryAudioFileHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAudioFileHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAudioFileHeaders) GoString() string {
	return s.String()
}

func (s *QueryAudioFileHeaders) SetCommonHeaders(v map[string]*string) *QueryAudioFileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAudioFileHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAudioFileHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAudioFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// A1
	DeviceType   *string `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
	EndTimestamp *int64  `json:"endTimestamp,omitempty" xml:"endTimestamp,omitempty"`
	// example:
	//
	// 5
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	Sn             *string `json:"sn,omitempty" xml:"sn,omitempty"`
	StartTimestamp *int64  `json:"startTimestamp,omitempty" xml:"startTimestamp,omitempty"`
}

func (s QueryAudioFileRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAudioFileRequest) GoString() string {
	return s.String()
}

func (s *QueryAudioFileRequest) SetDeviceType(v string) *QueryAudioFileRequest {
	s.DeviceType = &v
	return s
}

func (s *QueryAudioFileRequest) SetEndTimestamp(v int64) *QueryAudioFileRequest {
	s.EndTimestamp = &v
	return s
}

func (s *QueryAudioFileRequest) SetMaxResults(v int32) *QueryAudioFileRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryAudioFileRequest) SetNextToken(v string) *QueryAudioFileRequest {
	s.NextToken = &v
	return s
}

func (s *QueryAudioFileRequest) SetSn(v string) *QueryAudioFileRequest {
	s.Sn = &v
	return s
}

func (s *QueryAudioFileRequest) SetStartTimestamp(v int64) *QueryAudioFileRequest {
	s.StartTimestamp = &v
	return s
}

type QueryAudioFileResponseBody struct {
	NextToken  *string                             `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Result     []*QueryAudioFileResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int64                              `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryAudioFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAudioFileResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAudioFileResponseBody) SetNextToken(v string) *QueryAudioFileResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryAudioFileResponseBody) SetResult(v []*QueryAudioFileResponseBodyResult) *QueryAudioFileResponseBody {
	s.Result = v
	return s
}

func (s *QueryAudioFileResponseBody) SetTotalCount(v int64) *QueryAudioFileResponseBody {
	s.TotalCount = &v
	return s
}

type QueryAudioFileResponseBodyResult struct {
	CreateTime    *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	Duration      *int64  `json:"duration,omitempty" xml:"duration,omitempty"`
	FileId        *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	FileName      *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileSize      *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
}

func (s QueryAudioFileResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryAudioFileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryAudioFileResponseBodyResult) SetCreateTime(v int64) *QueryAudioFileResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *QueryAudioFileResponseBodyResult) SetCreatorUserId(v string) *QueryAudioFileResponseBodyResult {
	s.CreatorUserId = &v
	return s
}

func (s *QueryAudioFileResponseBodyResult) SetDuration(v int64) *QueryAudioFileResponseBodyResult {
	s.Duration = &v
	return s
}

func (s *QueryAudioFileResponseBodyResult) SetFileId(v string) *QueryAudioFileResponseBodyResult {
	s.FileId = &v
	return s
}

func (s *QueryAudioFileResponseBodyResult) SetFileName(v string) *QueryAudioFileResponseBodyResult {
	s.FileName = &v
	return s
}

func (s *QueryAudioFileResponseBodyResult) SetFileSize(v int64) *QueryAudioFileResponseBodyResult {
	s.FileSize = &v
	return s
}

type QueryAudioFileResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAudioFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAudioFileResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAudioFileResponse) GoString() string {
	return s.String()
}

func (s *QueryAudioFileResponse) SetHeaders(v map[string]*string) *QueryAudioFileResponse {
	s.Headers = v
	return s
}

func (s *QueryAudioFileResponse) SetStatusCode(v int32) *QueryAudioFileResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAudioFileResponse) SetBody(v *QueryAudioFileResponseBody) *QueryAudioFileResponse {
	s.Body = v
	return s
}

type QueryDeviceDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDeviceDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceDetailHeaders) GoString() string {
	return s.String()
}

func (s *QueryDeviceDetailHeaders) SetCommonHeaders(v map[string]*string) *QueryDeviceDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDeviceDetailHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDeviceDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDeviceDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// A1
	DeviceType *string `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
	// This parameter is required.
	SnList []*string `json:"snList,omitempty" xml:"snList,omitempty" type:"Repeated"`
}

func (s QueryDeviceDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceDetailRequest) SetDeviceType(v string) *QueryDeviceDetailRequest {
	s.DeviceType = &v
	return s
}

func (s *QueryDeviceDetailRequest) SetSnList(v []*string) *QueryDeviceDetailRequest {
	s.SnList = v
	return s
}

type QueryDeviceDetailResponseBody struct {
	Result []*QueryDeviceDetailResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryDeviceDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceDetailResponseBody) SetResult(v []*QueryDeviceDetailResponseBodyResult) *QueryDeviceDetailResponseBody {
	s.Result = v
	return s
}

type QueryDeviceDetailResponseBodyResult struct {
	BindTimestamp *int64  `json:"bindTimestamp,omitempty" xml:"bindTimestamp,omitempty"`
	DeviceName    *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	Sn            *string `json:"sn,omitempty" xml:"sn,omitempty"`
	UserId        *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryDeviceDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryDeviceDetailResponseBodyResult) SetBindTimestamp(v int64) *QueryDeviceDetailResponseBodyResult {
	s.BindTimestamp = &v
	return s
}

func (s *QueryDeviceDetailResponseBodyResult) SetDeviceName(v string) *QueryDeviceDetailResponseBodyResult {
	s.DeviceName = &v
	return s
}

func (s *QueryDeviceDetailResponseBodyResult) SetSn(v string) *QueryDeviceDetailResponseBodyResult {
	s.Sn = &v
	return s
}

func (s *QueryDeviceDetailResponseBodyResult) SetUserId(v string) *QueryDeviceDetailResponseBodyResult {
	s.UserId = &v
	return s
}

type QueryDeviceDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDeviceDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDeviceDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceDetailResponse) SetHeaders(v map[string]*string) *QueryDeviceDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceDetailResponse) SetStatusCode(v int32) *QueryDeviceDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceDetailResponse) SetBody(v *QueryDeviceDetailResponseBody) *QueryDeviceDetailResponse {
	s.Body = v
	return s
}

type QueryDeviceStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDeviceStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatusHeaders) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusHeaders) SetCommonHeaders(v map[string]*string) *QueryDeviceStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDeviceStatusHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDeviceStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDeviceStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// A1
	DeviceType *string `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
	// This parameter is required.
	SnList []*string `json:"snList,omitempty" xml:"snList,omitempty" type:"Repeated"`
}

func (s QueryDeviceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusRequest) SetDeviceType(v string) *QueryDeviceStatusRequest {
	s.DeviceType = &v
	return s
}

func (s *QueryDeviceStatusRequest) SetSnList(v []*string) *QueryDeviceStatusRequest {
	s.SnList = v
	return s
}

type QueryDeviceStatusResponseBody struct {
	Result []*QueryDeviceStatusResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryDeviceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusResponseBody) SetResult(v []*QueryDeviceStatusResponseBodyResult) *QueryDeviceStatusResponseBody {
	s.Result = v
	return s
}

type QueryDeviceStatusResponseBodyResult struct {
	Battery            *QueryDeviceStatusResponseBodyResultBattery            `json:"battery,omitempty" xml:"battery,omitempty" type:"Struct"`
	Firmware           *QueryDeviceStatusResponseBodyResultFirmware           `json:"firmware,omitempty" xml:"firmware,omitempty" type:"Struct"`
	RecordingStartTime *QueryDeviceStatusResponseBodyResultRecordingStartTime `json:"recordingStartTime,omitempty" xml:"recordingStartTime,omitempty" type:"Struct"`
	Sn                 *string                                                `json:"sn,omitempty" xml:"sn,omitempty"`
	Status             *QueryDeviceStatusResponseBodyResultStatus             `json:"status,omitempty" xml:"status,omitempty" type:"Struct"`
}

func (s QueryDeviceStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusResponseBodyResult) SetBattery(v *QueryDeviceStatusResponseBodyResultBattery) *QueryDeviceStatusResponseBodyResult {
	s.Battery = v
	return s
}

func (s *QueryDeviceStatusResponseBodyResult) SetFirmware(v *QueryDeviceStatusResponseBodyResultFirmware) *QueryDeviceStatusResponseBodyResult {
	s.Firmware = v
	return s
}

func (s *QueryDeviceStatusResponseBodyResult) SetRecordingStartTime(v *QueryDeviceStatusResponseBodyResultRecordingStartTime) *QueryDeviceStatusResponseBodyResult {
	s.RecordingStartTime = v
	return s
}

func (s *QueryDeviceStatusResponseBodyResult) SetSn(v string) *QueryDeviceStatusResponseBodyResult {
	s.Sn = &v
	return s
}

func (s *QueryDeviceStatusResponseBodyResult) SetStatus(v *QueryDeviceStatusResponseBodyResultStatus) *QueryDeviceStatusResponseBodyResult {
	s.Status = v
	return s
}

type QueryDeviceStatusResponseBodyResultBattery struct {
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Value     *int32 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s QueryDeviceStatusResponseBodyResultBattery) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatusResponseBodyResultBattery) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusResponseBodyResultBattery) SetTimestamp(v int64) *QueryDeviceStatusResponseBodyResultBattery {
	s.Timestamp = &v
	return s
}

func (s *QueryDeviceStatusResponseBodyResultBattery) SetValue(v int32) *QueryDeviceStatusResponseBodyResultBattery {
	s.Value = &v
	return s
}

type QueryDeviceStatusResponseBodyResultFirmware struct {
	Timestamp *int64  `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Value     *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s QueryDeviceStatusResponseBodyResultFirmware) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatusResponseBodyResultFirmware) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusResponseBodyResultFirmware) SetTimestamp(v int64) *QueryDeviceStatusResponseBodyResultFirmware {
	s.Timestamp = &v
	return s
}

func (s *QueryDeviceStatusResponseBodyResultFirmware) SetValue(v string) *QueryDeviceStatusResponseBodyResultFirmware {
	s.Value = &v
	return s
}

type QueryDeviceStatusResponseBodyResultRecordingStartTime struct {
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Value     *int64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s QueryDeviceStatusResponseBodyResultRecordingStartTime) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatusResponseBodyResultRecordingStartTime) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusResponseBodyResultRecordingStartTime) SetTimestamp(v int64) *QueryDeviceStatusResponseBodyResultRecordingStartTime {
	s.Timestamp = &v
	return s
}

func (s *QueryDeviceStatusResponseBodyResultRecordingStartTime) SetValue(v int64) *QueryDeviceStatusResponseBodyResultRecordingStartTime {
	s.Value = &v
	return s
}

type QueryDeviceStatusResponseBodyResultStatus struct {
	Timestamp *int64  `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Value     *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s QueryDeviceStatusResponseBodyResultStatus) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatusResponseBodyResultStatus) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusResponseBodyResultStatus) SetTimestamp(v int64) *QueryDeviceStatusResponseBodyResultStatus {
	s.Timestamp = &v
	return s
}

func (s *QueryDeviceStatusResponseBodyResultStatus) SetValue(v string) *QueryDeviceStatusResponseBodyResultStatus {
	s.Value = &v
	return s
}

type QueryDeviceStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDeviceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDeviceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusResponse) SetHeaders(v map[string]*string) *QueryDeviceStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceStatusResponse) SetStatusCode(v int32) *QueryDeviceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceStatusResponse) SetBody(v *QueryDeviceStatusResponseBody) *QueryDeviceStatusResponse {
	s.Body = v
	return s
}

type SubmitAsrTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SubmitAsrTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s SubmitAsrTaskHeaders) GoString() string {
	return s.String()
}

func (s *SubmitAsrTaskHeaders) SetCommonHeaders(v map[string]*string) *SubmitAsrTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SubmitAsrTaskHeaders) SetXAcsDingtalkAccessToken(v string) *SubmitAsrTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SubmitAsrTaskRequest struct {
	BizKey *string `json:"bizKey,omitempty" xml:"bizKey,omitempty"`
	// This parameter is required.
	DentryId *string   `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	Phrases  []*string `json:"phrases,omitempty" xml:"phrases,omitempty" type:"Repeated"`
	// This parameter is required.
	SourceLanguage *string `json:"sourceLanguage,omitempty" xml:"sourceLanguage,omitempty"`
	// This parameter is required.
	SpaceId       *string                            `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	Transcription *SubmitAsrTaskRequestTranscription `json:"transcription,omitempty" xml:"transcription,omitempty" type:"Struct"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s SubmitAsrTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitAsrTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitAsrTaskRequest) SetBizKey(v string) *SubmitAsrTaskRequest {
	s.BizKey = &v
	return s
}

func (s *SubmitAsrTaskRequest) SetDentryId(v string) *SubmitAsrTaskRequest {
	s.DentryId = &v
	return s
}

func (s *SubmitAsrTaskRequest) SetPhrases(v []*string) *SubmitAsrTaskRequest {
	s.Phrases = v
	return s
}

func (s *SubmitAsrTaskRequest) SetSourceLanguage(v string) *SubmitAsrTaskRequest {
	s.SourceLanguage = &v
	return s
}

func (s *SubmitAsrTaskRequest) SetSpaceId(v string) *SubmitAsrTaskRequest {
	s.SpaceId = &v
	return s
}

func (s *SubmitAsrTaskRequest) SetTranscription(v *SubmitAsrTaskRequestTranscription) *SubmitAsrTaskRequest {
	s.Transcription = v
	return s
}

func (s *SubmitAsrTaskRequest) SetUnionId(v string) *SubmitAsrTaskRequest {
	s.UnionId = &v
	return s
}

type SubmitAsrTaskRequestTranscription struct {
	Diarization        *SubmitAsrTaskRequestTranscriptionDiarization `json:"diarization,omitempty" xml:"diarization,omitempty" type:"Struct"`
	DiarizationEnabled *bool                                         `json:"diarizationEnabled,omitempty" xml:"diarizationEnabled,omitempty"`
}

func (s SubmitAsrTaskRequestTranscription) String() string {
	return tea.Prettify(s)
}

func (s SubmitAsrTaskRequestTranscription) GoString() string {
	return s.String()
}

func (s *SubmitAsrTaskRequestTranscription) SetDiarization(v *SubmitAsrTaskRequestTranscriptionDiarization) *SubmitAsrTaskRequestTranscription {
	s.Diarization = v
	return s
}

func (s *SubmitAsrTaskRequestTranscription) SetDiarizationEnabled(v bool) *SubmitAsrTaskRequestTranscription {
	s.DiarizationEnabled = &v
	return s
}

type SubmitAsrTaskRequestTranscriptionDiarization struct {
	SpeakerCount *int64 `json:"speakerCount,omitempty" xml:"speakerCount,omitempty"`
}

func (s SubmitAsrTaskRequestTranscriptionDiarization) String() string {
	return tea.Prettify(s)
}

func (s SubmitAsrTaskRequestTranscriptionDiarization) GoString() string {
	return s.String()
}

func (s *SubmitAsrTaskRequestTranscriptionDiarization) SetSpeakerCount(v int64) *SubmitAsrTaskRequestTranscriptionDiarization {
	s.SpeakerCount = &v
	return s
}

type SubmitAsrTaskResponseBody struct {
	ErrorCode *string                          `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                          `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result    *SubmitAsrTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success   *string                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SubmitAsrTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitAsrTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAsrTaskResponseBody) SetErrorCode(v string) *SubmitAsrTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitAsrTaskResponseBody) SetErrorMsg(v string) *SubmitAsrTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SubmitAsrTaskResponseBody) SetResult(v *SubmitAsrTaskResponseBodyResult) *SubmitAsrTaskResponseBody {
	s.Result = v
	return s
}

func (s *SubmitAsrTaskResponseBody) SetSuccess(v string) *SubmitAsrTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitAsrTaskResponseBodyResult struct {
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SubmitAsrTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SubmitAsrTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SubmitAsrTaskResponseBodyResult) SetTaskId(v string) *SubmitAsrTaskResponseBodyResult {
	s.TaskId = &v
	return s
}

type SubmitAsrTaskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAsrTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAsrTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitAsrTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitAsrTaskResponse) SetHeaders(v map[string]*string) *SubmitAsrTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitAsrTaskResponse) SetStatusCode(v int32) *SubmitAsrTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAsrTaskResponse) SetBody(v *SubmitAsrTaskResponseBody) *SubmitAsrTaskResponse {
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
// 获取音频文件下载地址
//
// @param request - GetAudioFileDownloadInfoRequest
//
// @param headers - GetAudioFileDownloadInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAudioFileDownloadInfoResponse
func (client *Client) GetAudioFileDownloadInfoWithOptions(request *GetAudioFileDownloadInfoRequest, headers *GetAudioFileDownloadInfoHeaders, runtime *util.RuntimeOptions) (_result *GetAudioFileDownloadInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceType)) {
		body["deviceType"] = request.DeviceType
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["fileId"] = request.FileId
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
		Action:      tea.String("GetAudioFileDownloadInfo"),
		Version:     tea.String("dvi_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/dvi/device/audio/download"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAudioFileDownloadInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取音频文件下载地址
//
// @param request - GetAudioFileDownloadInfoRequest
//
// @return GetAudioFileDownloadInfoResponse
func (client *Client) GetAudioFileDownloadInfo(request *GetAudioFileDownloadInfoRequest) (_result *GetAudioFileDownloadInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAudioFileDownloadInfoHeaders{}
	_result = &GetAudioFileDownloadInfoResponse{}
	_body, _err := client.GetAudioFileDownloadInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取音频文件信息
//
// @param request - GetAudioFileInfoRequest
//
// @param headers - GetAudioFileInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAudioFileInfoResponse
func (client *Client) GetAudioFileInfoWithOptions(request *GetAudioFileInfoRequest, headers *GetAudioFileInfoHeaders, runtime *util.RuntimeOptions) (_result *GetAudioFileInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceType)) {
		body["deviceType"] = request.DeviceType
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["fileId"] = request.FileId
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
		Action:      tea.String("GetAudioFileInfo"),
		Version:     tea.String("dvi_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/dvi/device/audio/get"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAudioFileInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取音频文件信息
//
// @param request - GetAudioFileInfoRequest
//
// @return GetAudioFileInfoResponse
func (client *Client) GetAudioFileInfo(request *GetAudioFileInfoRequest) (_result *GetAudioFileInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAudioFileInfoHeaders{}
	_result = &GetAudioFileInfoResponse{}
	_body, _err := client.GetAudioFileInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询客户数据
//
// @param request - GetCustomerInfoRequest
//
// @param headers - GetCustomerInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomerInfoResponse
func (client *Client) GetCustomerInfoWithOptions(request *GetCustomerInfoRequest, headers *GetCustomerInfoHeaders, runtime *util.RuntimeOptions) (_result *GetCustomerInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomerId)) {
		query["customerId"] = request.CustomerId
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
		Action:      tea.String("GetCustomerInfo"),
		Version:     tea.String("dvi_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/dvi/customer"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCustomerInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询客户数据
//
// @param request - GetCustomerInfoRequest
//
// @return GetCustomerInfoResponse
func (client *Client) GetCustomerInfo(request *GetCustomerInfoRequest) (_result *GetCustomerInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCustomerInfoHeaders{}
	_result = &GetCustomerInfoResponse{}
	_body, _err := client.GetCustomerInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取服务章节摘要
//
// @param request - GetServiceChapterSummaryRequest
//
// @param headers - GetServiceChapterSummaryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceChapterSummaryResponse
func (client *Client) GetServiceChapterSummaryWithOptions(request *GetServiceChapterSummaryRequest, headers *GetServiceChapterSummaryHeaders, runtime *util.RuntimeOptions) (_result *GetServiceChapterSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RecordId)) {
		query["recordId"] = request.RecordId
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
		Action:      tea.String("GetServiceChapterSummary"),
		Version:     tea.String("dvi_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/dvi/service/chapters/summary"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceChapterSummaryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取服务章节摘要
//
// @param request - GetServiceChapterSummaryRequest
//
// @return GetServiceChapterSummaryResponse
func (client *Client) GetServiceChapterSummary(request *GetServiceChapterSummaryRequest) (_result *GetServiceChapterSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetServiceChapterSummaryHeaders{}
	_result = &GetServiceChapterSummaryResponse{}
	_body, _err := client.GetServiceChapterSummaryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取服务会话总结
//
// @param request - GetServiceChatSummaryRequest
//
// @param headers - GetServiceChatSummaryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceChatSummaryResponse
func (client *Client) GetServiceChatSummaryWithOptions(request *GetServiceChatSummaryRequest, headers *GetServiceChatSummaryHeaders, runtime *util.RuntimeOptions) (_result *GetServiceChatSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RecordId)) {
		query["recordId"] = request.RecordId
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
		Action:      tea.String("GetServiceChatSummary"),
		Version:     tea.String("dvi_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/dvi/service/chats/summary"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceChatSummaryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取服务会话总结
//
// @param request - GetServiceChatSummaryRequest
//
// @return GetServiceChatSummaryResponse
func (client *Client) GetServiceChatSummary(request *GetServiceChatSummaryRequest) (_result *GetServiceChatSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetServiceChatSummaryHeaders{}
	_result = &GetServiceChatSummaryResponse{}
	_body, _err := client.GetServiceChatSummaryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询服务质检信息
//
// @param request - GetServiceQualityInspectionRequest
//
// @param headers - GetServiceQualityInspectionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceQualityInspectionResponse
func (client *Client) GetServiceQualityInspectionWithOptions(request *GetServiceQualityInspectionRequest, headers *GetServiceQualityInspectionHeaders, runtime *util.RuntimeOptions) (_result *GetServiceQualityInspectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RecordId)) {
		query["recordId"] = request.RecordId
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
		Action:      tea.String("GetServiceQualityInspection"),
		Version:     tea.String("dvi_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/dvi/service/quality-inspections"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceQualityInspectionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务质检信息
//
// @param request - GetServiceQualityInspectionRequest
//
// @return GetServiceQualityInspectionResponse
func (client *Client) GetServiceQualityInspection(request *GetServiceQualityInspectionRequest) (_result *GetServiceQualityInspectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetServiceQualityInspectionHeaders{}
	_result = &GetServiceQualityInspectionResponse{}
	_body, _err := client.GetServiceQualityInspectionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取服务记录音频转写信息
//
// @param request - GetServiceRecordTranscriptRequest
//
// @param headers - GetServiceRecordTranscriptHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceRecordTranscriptResponse
func (client *Client) GetServiceRecordTranscriptWithOptions(request *GetServiceRecordTranscriptRequest, headers *GetServiceRecordTranscriptHeaders, runtime *util.RuntimeOptions) (_result *GetServiceRecordTranscriptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["id"] = request.Id
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
		Action:      tea.String("GetServiceRecordTranscript"),
		Version:     tea.String("dvi_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/dvi/service/transcript"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceRecordTranscriptResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取服务记录音频转写信息
//
// @param request - GetServiceRecordTranscriptRequest
//
// @return GetServiceRecordTranscriptResponse
func (client *Client) GetServiceRecordTranscript(request *GetServiceRecordTranscriptRequest) (_result *GetServiceRecordTranscriptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetServiceRecordTranscriptHeaders{}
	_result = &GetServiceRecordTranscriptResponse{}
	_body, _err := client.GetServiceRecordTranscriptWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文件转写的概要信息
//
// @param request - GetTranscriptSummaryRequest
//
// @param headers - GetTranscriptSummaryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTranscriptSummaryResponse
func (client *Client) GetTranscriptSummaryWithOptions(request *GetTranscriptSummaryRequest, headers *GetTranscriptSummaryHeaders, runtime *util.RuntimeOptions) (_result *GetTranscriptSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceType)) {
		query["deviceType"] = request.DeviceType
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		query["fileId"] = request.FileId
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
		Action:      tea.String("GetTranscriptSummary"),
		Version:     tea.String("dvi_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/dvi/transcripts/summary"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTranscriptSummaryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件转写的概要信息
//
// @param request - GetTranscriptSummaryRequest
//
// @return GetTranscriptSummaryResponse
func (client *Client) GetTranscriptSummary(request *GetTranscriptSummaryRequest) (_result *GetTranscriptSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTranscriptSummaryHeaders{}
	_result = &GetTranscriptSummaryResponse{}
	_body, _err := client.GetTranscriptSummaryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询客户列表
//
// @param request - ListCustomerRequest
//
// @param headers - ListCustomerHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomerResponse
func (client *Client) ListCustomerWithOptions(request *ListCustomerRequest, headers *ListCustomerHeaders, runtime *util.RuntimeOptions) (_result *ListCustomerResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.OwnerUserId)) {
		query["ownerUserId"] = request.OwnerUserId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TeamCode)) {
		query["teamCode"] = request.TeamCode
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
		Action:      tea.String("ListCustomer"),
		Version:     tea.String("dvi_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/dvi/customers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCustomerResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询客户列表
//
// @param request - ListCustomerRequest
//
// @return ListCustomerResponse
func (client *Client) ListCustomer(request *ListCustomerRequest) (_result *ListCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListCustomerHeaders{}
	_result = &ListCustomerResponse{}
	_body, _err := client.ListCustomerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询服务记录
//
// @param request - ListServiceRecordRequest
//
// @param headers - ListServiceRecordHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceRecordResponse
func (client *Client) ListServiceRecordWithOptions(request *ListServiceRecordRequest, headers *ListServiceRecordHeaders, runtime *util.RuntimeOptions) (_result *ListServiceRecordResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TeamCode)) {
		query["teamCode"] = request.TeamCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
		Action:      tea.String("ListServiceRecord"),
		Version:     tea.String("dvi_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/dvi/service-records"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceRecordResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询服务记录
//
// @param request - ListServiceRecordRequest
//
// @return ListServiceRecordResponse
func (client *Client) ListServiceRecord(request *ListServiceRecordRequest) (_result *ListServiceRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListServiceRecordHeaders{}
	_result = &ListServiceRecordResponse{}
	_body, _err := client.ListServiceRecordWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询服务记录待办列表
//
// @param request - ListServiceTodoRequest
//
// @param headers - ListServiceTodoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceTodoResponse
func (client *Client) ListServiceTodoWithOptions(request *ListServiceTodoRequest, headers *ListServiceTodoHeaders, runtime *util.RuntimeOptions) (_result *ListServiceTodoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RecordId)) {
		query["recordId"] = request.RecordId
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
		Action:      tea.String("ListServiceTodo"),
		Version:     tea.String("dvi_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/dvi/service-todos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceTodoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务记录待办列表
//
// @param request - ListServiceTodoRequest
//
// @return ListServiceTodoResponse
func (client *Client) ListServiceTodo(request *ListServiceTodoRequest) (_result *ListServiceTodoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListServiceTodoHeaders{}
	_result = &ListServiceTodoResponse{}
	_body, _err := client.ListServiceTodoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询团队列表
//
// @param request - ListTeamRequest
//
// @param headers - ListTeamHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTeamResponse
func (client *Client) ListTeamWithOptions(request *ListTeamRequest, headers *ListTeamHeaders, runtime *util.RuntimeOptions) (_result *ListTeamResponse, _err error) {
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
		Action:      tea.String("ListTeam"),
		Version:     tea.String("dvi_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/dvi/teams"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTeamResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询团队列表
//
// @param request - ListTeamRequest
//
// @return ListTeamResponse
func (client *Client) ListTeam(request *ListTeamRequest) (_result *ListTeamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListTeamHeaders{}
	_result = &ListTeamResponse{}
	_body, _err := client.ListTeamWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询asr结果
//
// @param request - QueryAsrTaskRequest
//
// @param headers - QueryAsrTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAsrTaskResponse
func (client *Client) QueryAsrTaskWithOptions(request *QueryAsrTaskRequest, headers *QueryAsrTaskHeaders, runtime *util.RuntimeOptions) (_result *QueryAsrTaskResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
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
		Action:      tea.String("QueryAsrTask"),
		Version:     tea.String("dvi_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/dvi/asr/query"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAsrTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询asr结果
//
// @param request - QueryAsrTaskRequest
//
// @return QueryAsrTaskResponse
func (client *Client) QueryAsrTask(request *QueryAsrTaskRequest) (_result *QueryAsrTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAsrTaskHeaders{}
	_result = &QueryAsrTaskResponse{}
	_body, _err := client.QueryAsrTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询指定设备的音频文件列表
//
// @param request - QueryAudioFileRequest
//
// @param headers - QueryAudioFileHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAudioFileResponse
func (client *Client) QueryAudioFileWithOptions(request *QueryAudioFileRequest, headers *QueryAudioFileHeaders, runtime *util.RuntimeOptions) (_result *QueryAudioFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceType)) {
		body["deviceType"] = request.DeviceType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		body["endTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		body["sn"] = request.Sn
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		body["startTimestamp"] = request.StartTimestamp
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
		Action:      tea.String("QueryAudioFile"),
		Version:     tea.String("dvi_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/dvi/device/audio/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAudioFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询指定设备的音频文件列表
//
// @param request - QueryAudioFileRequest
//
// @return QueryAudioFileResponse
func (client *Client) QueryAudioFile(request *QueryAudioFileRequest) (_result *QueryAudioFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAudioFileHeaders{}
	_result = &QueryAudioFileResponse{}
	_body, _err := client.QueryAudioFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取设备列表
//
// @param request - QueryDeviceDetailRequest
//
// @param headers - QueryDeviceDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDeviceDetailResponse
func (client *Client) QueryDeviceDetailWithOptions(request *QueryDeviceDetailRequest, headers *QueryDeviceDetailHeaders, runtime *util.RuntimeOptions) (_result *QueryDeviceDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceType)) {
		body["deviceType"] = request.DeviceType
	}

	if !tea.BoolValue(util.IsUnset(request.SnList)) {
		body["snList"] = request.SnList
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
		Action:      tea.String("QueryDeviceDetail"),
		Version:     tea.String("dvi_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/dvi/device/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDeviceDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取设备列表
//
// @param request - QueryDeviceDetailRequest
//
// @return QueryDeviceDetailResponse
func (client *Client) QueryDeviceDetail(request *QueryDeviceDetailRequest) (_result *QueryDeviceDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDeviceDetailHeaders{}
	_result = &QueryDeviceDetailResponse{}
	_body, _err := client.QueryDeviceDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询设备属性
//
// @param request - QueryDeviceStatusRequest
//
// @param headers - QueryDeviceStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDeviceStatusResponse
func (client *Client) QueryDeviceStatusWithOptions(request *QueryDeviceStatusRequest, headers *QueryDeviceStatusHeaders, runtime *util.RuntimeOptions) (_result *QueryDeviceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceType)) {
		body["deviceType"] = request.DeviceType
	}

	if !tea.BoolValue(util.IsUnset(request.SnList)) {
		body["snList"] = request.SnList
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
		Action:      tea.String("QueryDeviceStatus"),
		Version:     tea.String("dvi_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/dvi/device/status"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDeviceStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询设备属性
//
// @param request - QueryDeviceStatusRequest
//
// @return QueryDeviceStatusResponse
func (client *Client) QueryDeviceStatus(request *QueryDeviceStatusRequest) (_result *QueryDeviceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDeviceStatusHeaders{}
	_result = &QueryDeviceStatusResponse{}
	_body, _err := client.QueryDeviceStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// asr离线任务
//
// @param request - SubmitAsrTaskRequest
//
// @param headers - SubmitAsrTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAsrTaskResponse
func (client *Client) SubmitAsrTaskWithOptions(request *SubmitAsrTaskRequest, headers *SubmitAsrTaskHeaders, runtime *util.RuntimeOptions) (_result *SubmitAsrTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizKey)) {
		body["bizKey"] = request.BizKey
	}

	if !tea.BoolValue(util.IsUnset(request.DentryId)) {
		body["dentryId"] = request.DentryId
	}

	if !tea.BoolValue(util.IsUnset(request.Phrases)) {
		body["phrases"] = request.Phrases
	}

	if !tea.BoolValue(util.IsUnset(request.SourceLanguage)) {
		body["sourceLanguage"] = request.SourceLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["spaceId"] = request.SpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.Transcription)) {
		body["transcription"] = request.Transcription
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("SubmitAsrTask"),
		Version:     tea.String("dvi_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/dvi/asr/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitAsrTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// asr离线任务
//
// @param request - SubmitAsrTaskRequest
//
// @return SubmitAsrTaskResponse
func (client *Client) SubmitAsrTask(request *SubmitAsrTaskRequest) (_result *SubmitAsrTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SubmitAsrTaskHeaders{}
	_result = &SubmitAsrTaskResponse{}
	_body, _err := client.SubmitAsrTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
