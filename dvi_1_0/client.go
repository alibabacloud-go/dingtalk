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
	BizKey       *string                                   `json:"bizKey,omitempty" xml:"bizKey,omitempty"`
	ErrorCode    *string                                   `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                   `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	ResultInfo   *QueryAsrTaskResponseBodyResultResultInfo `json:"resultInfo,omitempty" xml:"resultInfo,omitempty" type:"Struct"`
	TaskId       *string                                   `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TaskStatus   *string                                   `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
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

func (s *QueryAsrTaskResponseBodyResult) SetErrorCode(v string) *QueryAsrTaskResponseBodyResult {
	s.ErrorCode = &v
	return s
}

func (s *QueryAsrTaskResponseBodyResult) SetErrorMessage(v string) *QueryAsrTaskResponseBodyResult {
	s.ErrorMessage = &v
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
