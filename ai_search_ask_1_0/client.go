// This file is auto-generated, don't edit it. Thanks.
package ai_search_ask_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type RetrievalExtendParamsValue struct {
	StartTime        *int64                                        `json:"startTime,omitempty" xml:"startTime,omitempty"`
	EndTime          *int64                                        `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Keywords         []*string                                     `json:"keywords,omitempty" xml:"keywords,omitempty" type:"Repeated"`
	SourceUserParams []*RetrievalExtendParamsValueSourceUserParams `json:"sourceUserParams,omitempty" xml:"sourceUserParams,omitempty" type:"Repeated"`
	TargetUserParams []*RetrievalExtendParamsValueTargetUserParams `json:"targetUserParams,omitempty" xml:"targetUserParams,omitempty" type:"Repeated"`
}

func (s RetrievalExtendParamsValue) String() string {
	return tea.Prettify(s)
}

func (s RetrievalExtendParamsValue) GoString() string {
	return s.String()
}

func (s *RetrievalExtendParamsValue) SetStartTime(v int64) *RetrievalExtendParamsValue {
	s.StartTime = &v
	return s
}

func (s *RetrievalExtendParamsValue) SetEndTime(v int64) *RetrievalExtendParamsValue {
	s.EndTime = &v
	return s
}

func (s *RetrievalExtendParamsValue) SetKeywords(v []*string) *RetrievalExtendParamsValue {
	s.Keywords = v
	return s
}

func (s *RetrievalExtendParamsValue) SetSourceUserParams(v []*RetrievalExtendParamsValueSourceUserParams) *RetrievalExtendParamsValue {
	s.SourceUserParams = v
	return s
}

func (s *RetrievalExtendParamsValue) SetTargetUserParams(v []*RetrievalExtendParamsValueTargetUserParams) *RetrievalExtendParamsValue {
	s.TargetUserParams = v
	return s
}

type RetrievalExtendParamsValueSourceUserParams struct {
	Nick    *string `json:"nick,omitempty" xml:"nick,omitempty"`
	StaffId *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
}

func (s RetrievalExtendParamsValueSourceUserParams) String() string {
	return tea.Prettify(s)
}

func (s RetrievalExtendParamsValueSourceUserParams) GoString() string {
	return s.String()
}

func (s *RetrievalExtendParamsValueSourceUserParams) SetNick(v string) *RetrievalExtendParamsValueSourceUserParams {
	s.Nick = &v
	return s
}

func (s *RetrievalExtendParamsValueSourceUserParams) SetStaffId(v string) *RetrievalExtendParamsValueSourceUserParams {
	s.StaffId = &v
	return s
}

type RetrievalExtendParamsValueTargetUserParams struct {
	Nick    *string `json:"nick,omitempty" xml:"nick,omitempty"`
	StaffId *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
}

func (s RetrievalExtendParamsValueTargetUserParams) String() string {
	return tea.Prettify(s)
}

func (s RetrievalExtendParamsValueTargetUserParams) GoString() string {
	return s.String()
}

func (s *RetrievalExtendParamsValueTargetUserParams) SetNick(v string) *RetrievalExtendParamsValueTargetUserParams {
	s.Nick = &v
	return s
}

func (s *RetrievalExtendParamsValueTargetUserParams) SetStaffId(v string) *RetrievalExtendParamsValueTargetUserParams {
	s.StaffId = &v
	return s
}

type FetchLoginStatusDevicesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s FetchLoginStatusDevicesHeaders) String() string {
	return tea.Prettify(s)
}

func (s FetchLoginStatusDevicesHeaders) GoString() string {
	return s.String()
}

func (s *FetchLoginStatusDevicesHeaders) SetCommonHeaders(v map[string]*string) *FetchLoginStatusDevicesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FetchLoginStatusDevicesHeaders) SetXAcsDingtalkAccessToken(v string) *FetchLoginStatusDevicesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type FetchLoginStatusDevicesRequest struct {
	// This parameter is required.
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// This parameter is required.
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s FetchLoginStatusDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s FetchLoginStatusDevicesRequest) GoString() string {
	return s.String()
}

func (s *FetchLoginStatusDevicesRequest) SetDomain(v string) *FetchLoginStatusDevicesRequest {
	s.Domain = &v
	return s
}

func (s *FetchLoginStatusDevicesRequest) SetUserIdList(v []*string) *FetchLoginStatusDevicesRequest {
	s.UserIdList = v
	return s
}

type FetchLoginStatusDevicesResponseBody struct {
	ErrorCode *string                                      `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                                      `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result    []*FetchLoginStatusDevicesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success   *bool                                        `json:"success,omitempty" xml:"success,omitempty"`
}

func (s FetchLoginStatusDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FetchLoginStatusDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *FetchLoginStatusDevicesResponseBody) SetErrorCode(v string) *FetchLoginStatusDevicesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *FetchLoginStatusDevicesResponseBody) SetErrorMsg(v string) *FetchLoginStatusDevicesResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *FetchLoginStatusDevicesResponseBody) SetResult(v []*FetchLoginStatusDevicesResponseBodyResult) *FetchLoginStatusDevicesResponseBody {
	s.Result = v
	return s
}

func (s *FetchLoginStatusDevicesResponseBody) SetSuccess(v bool) *FetchLoginStatusDevicesResponseBody {
	s.Success = &v
	return s
}

type FetchLoginStatusDevicesResponseBodyResult struct {
	DeviceList []*FetchLoginStatusDevicesResponseBodyResultDeviceList `json:"deviceList,omitempty" xml:"deviceList,omitempty" type:"Repeated"`
	UserId     *string                                                `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s FetchLoginStatusDevicesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s FetchLoginStatusDevicesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *FetchLoginStatusDevicesResponseBodyResult) SetDeviceList(v []*FetchLoginStatusDevicesResponseBodyResultDeviceList) *FetchLoginStatusDevicesResponseBodyResult {
	s.DeviceList = v
	return s
}

func (s *FetchLoginStatusDevicesResponseBodyResult) SetUserId(v string) *FetchLoginStatusDevicesResponseBodyResult {
	s.UserId = &v
	return s
}

type FetchLoginStatusDevicesResponseBodyResultDeviceList struct {
	ActionTimestamp *int64  `json:"actionTimestamp,omitempty" xml:"actionTimestamp,omitempty"`
	ClientType      *string `json:"clientType,omitempty" xml:"clientType,omitempty"`
	IsLoggedIn      *bool   `json:"isLoggedIn,omitempty" xml:"isLoggedIn,omitempty"`
	OpenDeviceId    *string `json:"openDeviceId,omitempty" xml:"openDeviceId,omitempty"`
}

func (s FetchLoginStatusDevicesResponseBodyResultDeviceList) String() string {
	return tea.Prettify(s)
}

func (s FetchLoginStatusDevicesResponseBodyResultDeviceList) GoString() string {
	return s.String()
}

func (s *FetchLoginStatusDevicesResponseBodyResultDeviceList) SetActionTimestamp(v int64) *FetchLoginStatusDevicesResponseBodyResultDeviceList {
	s.ActionTimestamp = &v
	return s
}

func (s *FetchLoginStatusDevicesResponseBodyResultDeviceList) SetClientType(v string) *FetchLoginStatusDevicesResponseBodyResultDeviceList {
	s.ClientType = &v
	return s
}

func (s *FetchLoginStatusDevicesResponseBodyResultDeviceList) SetIsLoggedIn(v bool) *FetchLoginStatusDevicesResponseBodyResultDeviceList {
	s.IsLoggedIn = &v
	return s
}

func (s *FetchLoginStatusDevicesResponseBodyResultDeviceList) SetOpenDeviceId(v string) *FetchLoginStatusDevicesResponseBodyResultDeviceList {
	s.OpenDeviceId = &v
	return s
}

type FetchLoginStatusDevicesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FetchLoginStatusDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FetchLoginStatusDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s FetchLoginStatusDevicesResponse) GoString() string {
	return s.String()
}

func (s *FetchLoginStatusDevicesResponse) SetHeaders(v map[string]*string) *FetchLoginStatusDevicesResponse {
	s.Headers = v
	return s
}

func (s *FetchLoginStatusDevicesResponse) SetStatusCode(v int32) *FetchLoginStatusDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *FetchLoginStatusDevicesResponse) SetBody(v *FetchLoginStatusDevicesResponseBody) *FetchLoginStatusDevicesResponse {
	s.Body = v
	return s
}

type GetAiTaskResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAiTaskResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAiTaskResultHeaders) GoString() string {
	return s.String()
}

func (s *GetAiTaskResultHeaders) SetCommonHeaders(v map[string]*string) *GetAiTaskResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAiTaskResultHeaders) SetXAcsDingtalkAccessToken(v string) *GetAiTaskResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAiTaskResultRequest struct {
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetAiTaskResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAiTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetAiTaskResultRequest) SetTaskId(v string) *GetAiTaskResultRequest {
	s.TaskId = &v
	return s
}

type GetAiTaskResultResponseBody struct {
	Content     *string `json:"content,omitempty" xml:"content,omitempty"`
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	ErrorCode   *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg    *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Success     *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetAiTaskResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAiTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAiTaskResultResponseBody) SetContent(v string) *GetAiTaskResultResponseBody {
	s.Content = &v
	return s
}

func (s *GetAiTaskResultResponseBody) SetContentType(v string) *GetAiTaskResultResponseBody {
	s.ContentType = &v
	return s
}

func (s *GetAiTaskResultResponseBody) SetErrorCode(v string) *GetAiTaskResultResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAiTaskResultResponseBody) SetErrorMsg(v string) *GetAiTaskResultResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetAiTaskResultResponseBody) SetSuccess(v bool) *GetAiTaskResultResponseBody {
	s.Success = &v
	return s
}

type GetAiTaskResultResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAiTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAiTaskResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAiTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetAiTaskResultResponse) SetHeaders(v map[string]*string) *GetAiTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetAiTaskResultResponse) SetStatusCode(v int32) *GetAiTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAiTaskResultResponse) SetBody(v *GetAiTaskResultResponseBody) *GetAiTaskResultResponse {
	s.Body = v
	return s
}

type RetrieveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RetrieveHeaders) String() string {
	return tea.Prettify(s)
}

func (s RetrieveHeaders) GoString() string {
	return s.String()
}

func (s *RetrieveHeaders) SetCommonHeaders(v map[string]*string) *RetrieveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RetrieveHeaders) SetXAcsDingtalkAccessToken(v string) *RetrieveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RetrieveRequest struct {
	// This parameter is required.
	Answerer           *string                            `json:"answerer,omitempty" xml:"answerer,omitempty"`
	DragRequestContext *RetrieveRequestDragRequestContext `json:"dragRequestContext,omitempty" xml:"dragRequestContext,omitempty" type:"Struct"`
	KeywordList        []*string                          `json:"keywordList,omitempty" xml:"keywordList,omitempty" type:"Repeated"`
	// This parameter is required.
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// This parameter is required.
	Question *string `json:"question,omitempty" xml:"question,omitempty"`
	// This parameter is required.
	RetrievalExtendParams  map[string]*RetrievalExtendParamsValue `json:"retrievalExtendParams,omitempty" xml:"retrievalExtendParams,omitempty"`
	RetrieveScoreThreshold *float64                               `json:"retrieveScoreThreshold,omitempty" xml:"retrieveScoreThreshold,omitempty"`
	// This parameter is required.
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// This parameter is required.
	Tenant     *string `json:"tenant,omitempty" xml:"tenant,omitempty"`
	TokenLimit *int32  `json:"tokenLimit,omitempty" xml:"tokenLimit,omitempty"`
}

func (s RetrieveRequest) String() string {
	return tea.Prettify(s)
}

func (s RetrieveRequest) GoString() string {
	return s.String()
}

func (s *RetrieveRequest) SetAnswerer(v string) *RetrieveRequest {
	s.Answerer = &v
	return s
}

func (s *RetrieveRequest) SetDragRequestContext(v *RetrieveRequestDragRequestContext) *RetrieveRequest {
	s.DragRequestContext = v
	return s
}

func (s *RetrieveRequest) SetKeywordList(v []*string) *RetrieveRequest {
	s.KeywordList = v
	return s
}

func (s *RetrieveRequest) SetLimit(v int32) *RetrieveRequest {
	s.Limit = &v
	return s
}

func (s *RetrieveRequest) SetQuestion(v string) *RetrieveRequest {
	s.Question = &v
	return s
}

func (s *RetrieveRequest) SetRetrievalExtendParams(v map[string]*RetrievalExtendParamsValue) *RetrieveRequest {
	s.RetrievalExtendParams = v
	return s
}

func (s *RetrieveRequest) SetRetrieveScoreThreshold(v float64) *RetrieveRequest {
	s.RetrieveScoreThreshold = &v
	return s
}

func (s *RetrieveRequest) SetScene(v string) *RetrieveRequest {
	s.Scene = &v
	return s
}

func (s *RetrieveRequest) SetTenant(v string) *RetrieveRequest {
	s.Tenant = &v
	return s
}

func (s *RetrieveRequest) SetTokenLimit(v int32) *RetrieveRequest {
	s.TokenLimit = &v
	return s
}

type RetrieveRequestDragRequestContext struct {
	EvaluationContext *RetrieveRequestDragRequestContextEvaluationContext `json:"evaluationContext,omitempty" xml:"evaluationContext,omitempty" type:"Struct"`
}

func (s RetrieveRequestDragRequestContext) String() string {
	return tea.Prettify(s)
}

func (s RetrieveRequestDragRequestContext) GoString() string {
	return s.String()
}

func (s *RetrieveRequestDragRequestContext) SetEvaluationContext(v *RetrieveRequestDragRequestContextEvaluationContext) *RetrieveRequestDragRequestContext {
	s.EvaluationContext = v
	return s
}

type RetrieveRequestDragRequestContextEvaluationContext struct {
	SourceDentryId *string `json:"sourceDentryId,omitempty" xml:"sourceDentryId,omitempty"`
}

func (s RetrieveRequestDragRequestContextEvaluationContext) String() string {
	return tea.Prettify(s)
}

func (s RetrieveRequestDragRequestContextEvaluationContext) GoString() string {
	return s.String()
}

func (s *RetrieveRequestDragRequestContextEvaluationContext) SetSourceDentryId(v string) *RetrieveRequestDragRequestContextEvaluationContext {
	s.SourceDentryId = &v
	return s
}

type RetrieveResponseBody struct {
	ErrorCode *string                       `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                       `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result    []*RetrieveResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success   *bool                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RetrieveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBody) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBody) SetErrorCode(v string) *RetrieveResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RetrieveResponseBody) SetErrorMsg(v string) *RetrieveResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RetrieveResponseBody) SetResult(v []*RetrieveResponseBodyResult) *RetrieveResponseBody {
	s.Result = v
	return s
}

func (s *RetrieveResponseBody) SetSuccess(v bool) *RetrieveResponseBody {
	s.Success = &v
	return s
}

type RetrieveResponseBodyResult struct {
	Calendars      []*RetrieveResponseBodyResultCalendars      `json:"calendars,omitempty" xml:"calendars,omitempty" type:"Repeated"`
	DingHelperDocs []*RetrieveResponseBodyResultDingHelperDocs `json:"dingHelperDocs,omitempty" xml:"dingHelperDocs,omitempty" type:"Repeated"`
	Docs           []*RetrieveResponseBodyResultDocs           `json:"docs,omitempty" xml:"docs,omitempty" type:"Repeated"`
	Faqs           []*RetrieveResponseBodyResultFaqs           `json:"faqs,omitempty" xml:"faqs,omitempty" type:"Repeated"`
	FinalResults   []*RetrieveResponseBodyResultFinalResults   `json:"finalResults,omitempty" xml:"finalResults,omitempty" type:"Repeated"`
	Minutes        []*RetrieveResponseBodyResultMinutes        `json:"minutes,omitempty" xml:"minutes,omitempty" type:"Repeated"`
	Reports        []*RetrieveResponseBodyResultReports        `json:"reports,omitempty" xml:"reports,omitempty" type:"Repeated"`
	Wikis          []*RetrieveResponseBodyResultWikis          `json:"wikis,omitempty" xml:"wikis,omitempty" type:"Repeated"`
}

func (s RetrieveResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyResult) SetCalendars(v []*RetrieveResponseBodyResultCalendars) *RetrieveResponseBodyResult {
	s.Calendars = v
	return s
}

func (s *RetrieveResponseBodyResult) SetDingHelperDocs(v []*RetrieveResponseBodyResultDingHelperDocs) *RetrieveResponseBodyResult {
	s.DingHelperDocs = v
	return s
}

func (s *RetrieveResponseBodyResult) SetDocs(v []*RetrieveResponseBodyResultDocs) *RetrieveResponseBodyResult {
	s.Docs = v
	return s
}

func (s *RetrieveResponseBodyResult) SetFaqs(v []*RetrieveResponseBodyResultFaqs) *RetrieveResponseBodyResult {
	s.Faqs = v
	return s
}

func (s *RetrieveResponseBodyResult) SetFinalResults(v []*RetrieveResponseBodyResultFinalResults) *RetrieveResponseBodyResult {
	s.FinalResults = v
	return s
}

func (s *RetrieveResponseBodyResult) SetMinutes(v []*RetrieveResponseBodyResultMinutes) *RetrieveResponseBodyResult {
	s.Minutes = v
	return s
}

func (s *RetrieveResponseBodyResult) SetReports(v []*RetrieveResponseBodyResultReports) *RetrieveResponseBodyResult {
	s.Reports = v
	return s
}

func (s *RetrieveResponseBodyResult) SetWikis(v []*RetrieveResponseBodyResultWikis) *RetrieveResponseBodyResult {
	s.Wikis = v
	return s
}

type RetrieveResponseBodyResultCalendars struct {
	CreatorStaffId      *int64                                        `json:"creatorStaffId,omitempty" xml:"creatorStaffId,omitempty"`
	EndTime             *int64                                        `json:"endTime,omitempty" xml:"endTime,omitempty"`
	GmtCreate           *int64                                        `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified         *int64                                        `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Icon                *string                                       `json:"icon,omitempty" xml:"icon,omitempty"`
	MatchIntimacy       *bool                                         `json:"matchIntimacy,omitempty" xml:"matchIntimacy,omitempty"`
	Material            *string                                       `json:"material,omitempty" xml:"material,omitempty"`
	MeetingRoom         *string                                       `json:"meetingRoom,omitempty" xml:"meetingRoom,omitempty"`
	ParticipantStaffIds []*string                                     `json:"participantStaffIds,omitempty" xml:"participantStaffIds,omitempty" type:"Repeated"`
	RawComment          *string                                       `json:"rawComment,omitempty" xml:"rawComment,omitempty"`
	RefType             *string                                       `json:"refType,omitempty" xml:"refType,omitempty"`
	Score               *float64                                      `json:"score,omitempty" xml:"score,omitempty"`
	ScoreItem           *RetrieveResponseBodyResultCalendarsScoreItem `json:"scoreItem,omitempty" xml:"scoreItem,omitempty" type:"Struct"`
	StartTime           *int64                                        `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Timestamp           *int64                                        `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Title               *string                                       `json:"title,omitempty" xml:"title,omitempty"`
	Type                *string                                       `json:"type,omitempty" xml:"type,omitempty"`
	Url                 *string                                       `json:"url,omitempty" xml:"url,omitempty"`
}

func (s RetrieveResponseBodyResultCalendars) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBodyResultCalendars) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyResultCalendars) SetCreatorStaffId(v int64) *RetrieveResponseBodyResultCalendars {
	s.CreatorStaffId = &v
	return s
}

func (s *RetrieveResponseBodyResultCalendars) SetEndTime(v int64) *RetrieveResponseBodyResultCalendars {
	s.EndTime = &v
	return s
}

func (s *RetrieveResponseBodyResultCalendars) SetGmtCreate(v int64) *RetrieveResponseBodyResultCalendars {
	s.GmtCreate = &v
	return s
}

func (s *RetrieveResponseBodyResultCalendars) SetGmtModified(v int64) *RetrieveResponseBodyResultCalendars {
	s.GmtModified = &v
	return s
}

func (s *RetrieveResponseBodyResultCalendars) SetIcon(v string) *RetrieveResponseBodyResultCalendars {
	s.Icon = &v
	return s
}

func (s *RetrieveResponseBodyResultCalendars) SetMatchIntimacy(v bool) *RetrieveResponseBodyResultCalendars {
	s.MatchIntimacy = &v
	return s
}

func (s *RetrieveResponseBodyResultCalendars) SetMaterial(v string) *RetrieveResponseBodyResultCalendars {
	s.Material = &v
	return s
}

func (s *RetrieveResponseBodyResultCalendars) SetMeetingRoom(v string) *RetrieveResponseBodyResultCalendars {
	s.MeetingRoom = &v
	return s
}

func (s *RetrieveResponseBodyResultCalendars) SetParticipantStaffIds(v []*string) *RetrieveResponseBodyResultCalendars {
	s.ParticipantStaffIds = v
	return s
}

func (s *RetrieveResponseBodyResultCalendars) SetRawComment(v string) *RetrieveResponseBodyResultCalendars {
	s.RawComment = &v
	return s
}

func (s *RetrieveResponseBodyResultCalendars) SetRefType(v string) *RetrieveResponseBodyResultCalendars {
	s.RefType = &v
	return s
}

func (s *RetrieveResponseBodyResultCalendars) SetScore(v float64) *RetrieveResponseBodyResultCalendars {
	s.Score = &v
	return s
}

func (s *RetrieveResponseBodyResultCalendars) SetScoreItem(v *RetrieveResponseBodyResultCalendarsScoreItem) *RetrieveResponseBodyResultCalendars {
	s.ScoreItem = v
	return s
}

func (s *RetrieveResponseBodyResultCalendars) SetStartTime(v int64) *RetrieveResponseBodyResultCalendars {
	s.StartTime = &v
	return s
}

func (s *RetrieveResponseBodyResultCalendars) SetTimestamp(v int64) *RetrieveResponseBodyResultCalendars {
	s.Timestamp = &v
	return s
}

func (s *RetrieveResponseBodyResultCalendars) SetTitle(v string) *RetrieveResponseBodyResultCalendars {
	s.Title = &v
	return s
}

func (s *RetrieveResponseBodyResultCalendars) SetType(v string) *RetrieveResponseBodyResultCalendars {
	s.Type = &v
	return s
}

func (s *RetrieveResponseBodyResultCalendars) SetUrl(v string) *RetrieveResponseBodyResultCalendars {
	s.Url = &v
	return s
}

type RetrieveResponseBodyResultCalendarsScoreItem struct {
	FinallyScore     *float64            `json:"finallyScore,omitempty" xml:"finallyScore,omitempty"`
	ScoreMap         map[string]*float64 `json:"scoreMap,omitempty" xml:"scoreMap,omitempty"`
	ScoreThreshold   *float64            `json:"scoreThreshold,omitempty" xml:"scoreThreshold,omitempty"`
	SelectedBranch   []*string           `json:"selectedBranch,omitempty" xml:"selectedBranch,omitempty" type:"Repeated"`
	SelectedCategory *string             `json:"selectedCategory,omitempty" xml:"selectedCategory,omitempty"`
}

func (s RetrieveResponseBodyResultCalendarsScoreItem) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBodyResultCalendarsScoreItem) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyResultCalendarsScoreItem) SetFinallyScore(v float64) *RetrieveResponseBodyResultCalendarsScoreItem {
	s.FinallyScore = &v
	return s
}

func (s *RetrieveResponseBodyResultCalendarsScoreItem) SetScoreMap(v map[string]*float64) *RetrieveResponseBodyResultCalendarsScoreItem {
	s.ScoreMap = v
	return s
}

func (s *RetrieveResponseBodyResultCalendarsScoreItem) SetScoreThreshold(v float64) *RetrieveResponseBodyResultCalendarsScoreItem {
	s.ScoreThreshold = &v
	return s
}

func (s *RetrieveResponseBodyResultCalendarsScoreItem) SetSelectedBranch(v []*string) *RetrieveResponseBodyResultCalendarsScoreItem {
	s.SelectedBranch = v
	return s
}

func (s *RetrieveResponseBodyResultCalendarsScoreItem) SetSelectedCategory(v string) *RetrieveResponseBodyResultCalendarsScoreItem {
	s.SelectedCategory = &v
	return s
}

type RetrieveResponseBodyResultDingHelperDocs struct {
	AbleDashboardModel *RetrieveResponseBodyResultDingHelperDocsAbleDashboardModel `json:"ableDashboardModel,omitempty" xml:"ableDashboardModel,omitempty" type:"Struct"`
	ChunkId            *int32                                                      `json:"chunkId,omitempty" xml:"chunkId,omitempty"`
	ChunkIds           []*int32                                                    `json:"chunkIds,omitempty" xml:"chunkIds,omitempty" type:"Repeated"`
	ChunkModel         *string                                                     `json:"chunkModel,omitempty" xml:"chunkModel,omitempty"`
	Creator            *string                                                     `json:"creator,omitempty" xml:"creator,omitempty"`
	DentryUuid         *string                                                     `json:"dentryUuid,omitempty" xml:"dentryUuid,omitempty"`
	Extension          *string                                                     `json:"extension,omitempty" xml:"extension,omitempty"`
	GmtCreate          *int64                                                      `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified        *int64                                                      `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	HasExtendChunk     *bool                                                       `json:"hasExtendChunk,omitempty" xml:"hasExtendChunk,omitempty"`
	Icon               *string                                                     `json:"icon,omitempty" xml:"icon,omitempty"`
	MatchIntimacy      *bool                                                       `json:"matchIntimacy,omitempty" xml:"matchIntimacy,omitempty"`
	Material           *string                                                     `json:"material,omitempty" xml:"material,omitempty"`
	RefType            *string                                                     `json:"refType,omitempty" xml:"refType,omitempty"`
	Score              *float64                                                    `json:"score,omitempty" xml:"score,omitempty"`
	ScoreItem          *RetrieveResponseBodyResultDingHelperDocsScoreItem          `json:"scoreItem,omitempty" xml:"scoreItem,omitempty" type:"Struct"`
	Small2bigText      *string                                                     `json:"small2bigText,omitempty" xml:"small2bigText,omitempty"`
	Text               *string                                                     `json:"text,omitempty" xml:"text,omitempty"`
	Timestamp          *int64                                                      `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Title              *string                                                     `json:"title,omitempty" xml:"title,omitempty"`
	Type               *string                                                     `json:"type,omitempty" xml:"type,omitempty"`
	UploadSource       *string                                                     `json:"uploadSource,omitempty" xml:"uploadSource,omitempty"`
	Url                *string                                                     `json:"url,omitempty" xml:"url,omitempty"`
}

func (s RetrieveResponseBodyResultDingHelperDocs) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBodyResultDingHelperDocs) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyResultDingHelperDocs) SetAbleDashboardModel(v *RetrieveResponseBodyResultDingHelperDocsAbleDashboardModel) *RetrieveResponseBodyResultDingHelperDocs {
	s.AbleDashboardModel = v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocs) SetChunkId(v int32) *RetrieveResponseBodyResultDingHelperDocs {
	s.ChunkId = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocs) SetChunkIds(v []*int32) *RetrieveResponseBodyResultDingHelperDocs {
	s.ChunkIds = v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocs) SetChunkModel(v string) *RetrieveResponseBodyResultDingHelperDocs {
	s.ChunkModel = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocs) SetCreator(v string) *RetrieveResponseBodyResultDingHelperDocs {
	s.Creator = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocs) SetDentryUuid(v string) *RetrieveResponseBodyResultDingHelperDocs {
	s.DentryUuid = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocs) SetExtension(v string) *RetrieveResponseBodyResultDingHelperDocs {
	s.Extension = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocs) SetGmtCreate(v int64) *RetrieveResponseBodyResultDingHelperDocs {
	s.GmtCreate = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocs) SetGmtModified(v int64) *RetrieveResponseBodyResultDingHelperDocs {
	s.GmtModified = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocs) SetHasExtendChunk(v bool) *RetrieveResponseBodyResultDingHelperDocs {
	s.HasExtendChunk = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocs) SetIcon(v string) *RetrieveResponseBodyResultDingHelperDocs {
	s.Icon = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocs) SetMatchIntimacy(v bool) *RetrieveResponseBodyResultDingHelperDocs {
	s.MatchIntimacy = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocs) SetMaterial(v string) *RetrieveResponseBodyResultDingHelperDocs {
	s.Material = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocs) SetRefType(v string) *RetrieveResponseBodyResultDingHelperDocs {
	s.RefType = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocs) SetScore(v float64) *RetrieveResponseBodyResultDingHelperDocs {
	s.Score = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocs) SetScoreItem(v *RetrieveResponseBodyResultDingHelperDocsScoreItem) *RetrieveResponseBodyResultDingHelperDocs {
	s.ScoreItem = v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocs) SetSmall2bigText(v string) *RetrieveResponseBodyResultDingHelperDocs {
	s.Small2bigText = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocs) SetText(v string) *RetrieveResponseBodyResultDingHelperDocs {
	s.Text = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocs) SetTimestamp(v int64) *RetrieveResponseBodyResultDingHelperDocs {
	s.Timestamp = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocs) SetTitle(v string) *RetrieveResponseBodyResultDingHelperDocs {
	s.Title = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocs) SetType(v string) *RetrieveResponseBodyResultDingHelperDocs {
	s.Type = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocs) SetUploadSource(v string) *RetrieveResponseBodyResultDingHelperDocs {
	s.UploadSource = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocs) SetUrl(v string) *RetrieveResponseBodyResultDingHelperDocs {
	s.Url = &v
	return s
}

type RetrieveResponseBodyResultDingHelperDocsAbleDashboardModel struct {
	ChartName     *string `json:"chartName,omitempty" xml:"chartName,omitempty"`
	ChartType     *string `json:"chartType,omitempty" xml:"chartType,omitempty"`
	DashboardName *string `json:"dashboardName,omitempty" xml:"dashboardName,omitempty"`
	Data          *string `json:"data,omitempty" xml:"data,omitempty"`
	SheetName     *string `json:"sheetName,omitempty" xml:"sheetName,omitempty"`
	Type          *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RetrieveResponseBodyResultDingHelperDocsAbleDashboardModel) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBodyResultDingHelperDocsAbleDashboardModel) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyResultDingHelperDocsAbleDashboardModel) SetChartName(v string) *RetrieveResponseBodyResultDingHelperDocsAbleDashboardModel {
	s.ChartName = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocsAbleDashboardModel) SetChartType(v string) *RetrieveResponseBodyResultDingHelperDocsAbleDashboardModel {
	s.ChartType = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocsAbleDashboardModel) SetDashboardName(v string) *RetrieveResponseBodyResultDingHelperDocsAbleDashboardModel {
	s.DashboardName = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocsAbleDashboardModel) SetData(v string) *RetrieveResponseBodyResultDingHelperDocsAbleDashboardModel {
	s.Data = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocsAbleDashboardModel) SetSheetName(v string) *RetrieveResponseBodyResultDingHelperDocsAbleDashboardModel {
	s.SheetName = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocsAbleDashboardModel) SetType(v string) *RetrieveResponseBodyResultDingHelperDocsAbleDashboardModel {
	s.Type = &v
	return s
}

type RetrieveResponseBodyResultDingHelperDocsScoreItem struct {
	FinallyScore     *float64            `json:"finallyScore,omitempty" xml:"finallyScore,omitempty"`
	ScoreMap         map[string]*float64 `json:"scoreMap,omitempty" xml:"scoreMap,omitempty"`
	ScoreThreshold   *float64            `json:"scoreThreshold,omitempty" xml:"scoreThreshold,omitempty"`
	SelectedBranch   []*string           `json:"selectedBranch,omitempty" xml:"selectedBranch,omitempty" type:"Repeated"`
	SelectedCategory *string             `json:"selectedCategory,omitempty" xml:"selectedCategory,omitempty"`
}

func (s RetrieveResponseBodyResultDingHelperDocsScoreItem) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBodyResultDingHelperDocsScoreItem) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyResultDingHelperDocsScoreItem) SetFinallyScore(v float64) *RetrieveResponseBodyResultDingHelperDocsScoreItem {
	s.FinallyScore = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocsScoreItem) SetScoreMap(v map[string]*float64) *RetrieveResponseBodyResultDingHelperDocsScoreItem {
	s.ScoreMap = v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocsScoreItem) SetScoreThreshold(v float64) *RetrieveResponseBodyResultDingHelperDocsScoreItem {
	s.ScoreThreshold = &v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocsScoreItem) SetSelectedBranch(v []*string) *RetrieveResponseBodyResultDingHelperDocsScoreItem {
	s.SelectedBranch = v
	return s
}

func (s *RetrieveResponseBodyResultDingHelperDocsScoreItem) SetSelectedCategory(v string) *RetrieveResponseBodyResultDingHelperDocsScoreItem {
	s.SelectedCategory = &v
	return s
}

type RetrieveResponseBodyResultDocs struct {
	AbleDashboardModel *RetrieveResponseBodyResultDocsAbleDashboardModel `json:"ableDashboardModel,omitempty" xml:"ableDashboardModel,omitempty" type:"Struct"`
	ChunkId            *int32                                            `json:"chunkId,omitempty" xml:"chunkId,omitempty"`
	ChunkIds           []*int32                                          `json:"chunkIds,omitempty" xml:"chunkIds,omitempty" type:"Repeated"`
	ChunkModel         *string                                           `json:"chunkModel,omitempty" xml:"chunkModel,omitempty"`
	Creator            *string                                           `json:"creator,omitempty" xml:"creator,omitempty"`
	DentryUuid         *string                                           `json:"dentryUuid,omitempty" xml:"dentryUuid,omitempty"`
	Extension          *string                                           `json:"extension,omitempty" xml:"extension,omitempty"`
	GmtCreate          *int64                                            `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified        *int64                                            `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	HasExtendChunk     *bool                                             `json:"hasExtendChunk,omitempty" xml:"hasExtendChunk,omitempty"`
	Icon               *string                                           `json:"icon,omitempty" xml:"icon,omitempty"`
	MatchIntimacy      *bool                                             `json:"matchIntimacy,omitempty" xml:"matchIntimacy,omitempty"`
	Material           *string                                           `json:"material,omitempty" xml:"material,omitempty"`
	RefType            *string                                           `json:"refType,omitempty" xml:"refType,omitempty"`
	Score              *float64                                          `json:"score,omitempty" xml:"score,omitempty"`
	ScoreItem          *RetrieveResponseBodyResultDocsScoreItem          `json:"scoreItem,omitempty" xml:"scoreItem,omitempty" type:"Struct"`
	Small2bigText      *string                                           `json:"small2bigText,omitempty" xml:"small2bigText,omitempty"`
	Text               *string                                           `json:"text,omitempty" xml:"text,omitempty"`
	Timestamp          *int64                                            `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Title              *string                                           `json:"title,omitempty" xml:"title,omitempty"`
	Type               *string                                           `json:"type,omitempty" xml:"type,omitempty"`
	UploadSource       *string                                           `json:"uploadSource,omitempty" xml:"uploadSource,omitempty"`
	Url                *string                                           `json:"url,omitempty" xml:"url,omitempty"`
}

func (s RetrieveResponseBodyResultDocs) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBodyResultDocs) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyResultDocs) SetAbleDashboardModel(v *RetrieveResponseBodyResultDocsAbleDashboardModel) *RetrieveResponseBodyResultDocs {
	s.AbleDashboardModel = v
	return s
}

func (s *RetrieveResponseBodyResultDocs) SetChunkId(v int32) *RetrieveResponseBodyResultDocs {
	s.ChunkId = &v
	return s
}

func (s *RetrieveResponseBodyResultDocs) SetChunkIds(v []*int32) *RetrieveResponseBodyResultDocs {
	s.ChunkIds = v
	return s
}

func (s *RetrieveResponseBodyResultDocs) SetChunkModel(v string) *RetrieveResponseBodyResultDocs {
	s.ChunkModel = &v
	return s
}

func (s *RetrieveResponseBodyResultDocs) SetCreator(v string) *RetrieveResponseBodyResultDocs {
	s.Creator = &v
	return s
}

func (s *RetrieveResponseBodyResultDocs) SetDentryUuid(v string) *RetrieveResponseBodyResultDocs {
	s.DentryUuid = &v
	return s
}

func (s *RetrieveResponseBodyResultDocs) SetExtension(v string) *RetrieveResponseBodyResultDocs {
	s.Extension = &v
	return s
}

func (s *RetrieveResponseBodyResultDocs) SetGmtCreate(v int64) *RetrieveResponseBodyResultDocs {
	s.GmtCreate = &v
	return s
}

func (s *RetrieveResponseBodyResultDocs) SetGmtModified(v int64) *RetrieveResponseBodyResultDocs {
	s.GmtModified = &v
	return s
}

func (s *RetrieveResponseBodyResultDocs) SetHasExtendChunk(v bool) *RetrieveResponseBodyResultDocs {
	s.HasExtendChunk = &v
	return s
}

func (s *RetrieveResponseBodyResultDocs) SetIcon(v string) *RetrieveResponseBodyResultDocs {
	s.Icon = &v
	return s
}

func (s *RetrieveResponseBodyResultDocs) SetMatchIntimacy(v bool) *RetrieveResponseBodyResultDocs {
	s.MatchIntimacy = &v
	return s
}

func (s *RetrieveResponseBodyResultDocs) SetMaterial(v string) *RetrieveResponseBodyResultDocs {
	s.Material = &v
	return s
}

func (s *RetrieveResponseBodyResultDocs) SetRefType(v string) *RetrieveResponseBodyResultDocs {
	s.RefType = &v
	return s
}

func (s *RetrieveResponseBodyResultDocs) SetScore(v float64) *RetrieveResponseBodyResultDocs {
	s.Score = &v
	return s
}

func (s *RetrieveResponseBodyResultDocs) SetScoreItem(v *RetrieveResponseBodyResultDocsScoreItem) *RetrieveResponseBodyResultDocs {
	s.ScoreItem = v
	return s
}

func (s *RetrieveResponseBodyResultDocs) SetSmall2bigText(v string) *RetrieveResponseBodyResultDocs {
	s.Small2bigText = &v
	return s
}

func (s *RetrieveResponseBodyResultDocs) SetText(v string) *RetrieveResponseBodyResultDocs {
	s.Text = &v
	return s
}

func (s *RetrieveResponseBodyResultDocs) SetTimestamp(v int64) *RetrieveResponseBodyResultDocs {
	s.Timestamp = &v
	return s
}

func (s *RetrieveResponseBodyResultDocs) SetTitle(v string) *RetrieveResponseBodyResultDocs {
	s.Title = &v
	return s
}

func (s *RetrieveResponseBodyResultDocs) SetType(v string) *RetrieveResponseBodyResultDocs {
	s.Type = &v
	return s
}

func (s *RetrieveResponseBodyResultDocs) SetUploadSource(v string) *RetrieveResponseBodyResultDocs {
	s.UploadSource = &v
	return s
}

func (s *RetrieveResponseBodyResultDocs) SetUrl(v string) *RetrieveResponseBodyResultDocs {
	s.Url = &v
	return s
}

type RetrieveResponseBodyResultDocsAbleDashboardModel struct {
	ChartName     *string `json:"chartName,omitempty" xml:"chartName,omitempty"`
	ChartType     *string `json:"chartType,omitempty" xml:"chartType,omitempty"`
	DashboardName *string `json:"dashboardName,omitempty" xml:"dashboardName,omitempty"`
	Data          *string `json:"data,omitempty" xml:"data,omitempty"`
	SheetName     *string `json:"sheetName,omitempty" xml:"sheetName,omitempty"`
	Type          *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RetrieveResponseBodyResultDocsAbleDashboardModel) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBodyResultDocsAbleDashboardModel) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyResultDocsAbleDashboardModel) SetChartName(v string) *RetrieveResponseBodyResultDocsAbleDashboardModel {
	s.ChartName = &v
	return s
}

func (s *RetrieveResponseBodyResultDocsAbleDashboardModel) SetChartType(v string) *RetrieveResponseBodyResultDocsAbleDashboardModel {
	s.ChartType = &v
	return s
}

func (s *RetrieveResponseBodyResultDocsAbleDashboardModel) SetDashboardName(v string) *RetrieveResponseBodyResultDocsAbleDashboardModel {
	s.DashboardName = &v
	return s
}

func (s *RetrieveResponseBodyResultDocsAbleDashboardModel) SetData(v string) *RetrieveResponseBodyResultDocsAbleDashboardModel {
	s.Data = &v
	return s
}

func (s *RetrieveResponseBodyResultDocsAbleDashboardModel) SetSheetName(v string) *RetrieveResponseBodyResultDocsAbleDashboardModel {
	s.SheetName = &v
	return s
}

func (s *RetrieveResponseBodyResultDocsAbleDashboardModel) SetType(v string) *RetrieveResponseBodyResultDocsAbleDashboardModel {
	s.Type = &v
	return s
}

type RetrieveResponseBodyResultDocsScoreItem struct {
	FinallyScore     *float64            `json:"finallyScore,omitempty" xml:"finallyScore,omitempty"`
	ScoreMap         map[string]*float64 `json:"scoreMap,omitempty" xml:"scoreMap,omitempty"`
	ScoreThreshold   *float64            `json:"scoreThreshold,omitempty" xml:"scoreThreshold,omitempty"`
	SelectedBranch   []*string           `json:"selectedBranch,omitempty" xml:"selectedBranch,omitempty" type:"Repeated"`
	SelectedCategory *string             `json:"selectedCategory,omitempty" xml:"selectedCategory,omitempty"`
}

func (s RetrieveResponseBodyResultDocsScoreItem) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBodyResultDocsScoreItem) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyResultDocsScoreItem) SetFinallyScore(v float64) *RetrieveResponseBodyResultDocsScoreItem {
	s.FinallyScore = &v
	return s
}

func (s *RetrieveResponseBodyResultDocsScoreItem) SetScoreMap(v map[string]*float64) *RetrieveResponseBodyResultDocsScoreItem {
	s.ScoreMap = v
	return s
}

func (s *RetrieveResponseBodyResultDocsScoreItem) SetScoreThreshold(v float64) *RetrieveResponseBodyResultDocsScoreItem {
	s.ScoreThreshold = &v
	return s
}

func (s *RetrieveResponseBodyResultDocsScoreItem) SetSelectedBranch(v []*string) *RetrieveResponseBodyResultDocsScoreItem {
	s.SelectedBranch = v
	return s
}

func (s *RetrieveResponseBodyResultDocsScoreItem) SetSelectedCategory(v string) *RetrieveResponseBodyResultDocsScoreItem {
	s.SelectedCategory = &v
	return s
}

type RetrieveResponseBodyResultFaqs struct {
	Icon          *string                                  `json:"icon,omitempty" xml:"icon,omitempty"`
	MatchIntimacy *bool                                    `json:"matchIntimacy,omitempty" xml:"matchIntimacy,omitempty"`
	Material      *string                                  `json:"material,omitempty" xml:"material,omitempty"`
	RefType       *string                                  `json:"refType,omitempty" xml:"refType,omitempty"`
	Score         *float64                                 `json:"score,omitempty" xml:"score,omitempty"`
	ScoreItem     *RetrieveResponseBodyResultFaqsScoreItem `json:"scoreItem,omitempty" xml:"scoreItem,omitempty" type:"Struct"`
	Timestamp     *int64                                   `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Title         *string                                  `json:"title,omitempty" xml:"title,omitempty"`
	Type          *string                                  `json:"type,omitempty" xml:"type,omitempty"`
	Url           *string                                  `json:"url,omitempty" xml:"url,omitempty"`
}

func (s RetrieveResponseBodyResultFaqs) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBodyResultFaqs) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyResultFaqs) SetIcon(v string) *RetrieveResponseBodyResultFaqs {
	s.Icon = &v
	return s
}

func (s *RetrieveResponseBodyResultFaqs) SetMatchIntimacy(v bool) *RetrieveResponseBodyResultFaqs {
	s.MatchIntimacy = &v
	return s
}

func (s *RetrieveResponseBodyResultFaqs) SetMaterial(v string) *RetrieveResponseBodyResultFaqs {
	s.Material = &v
	return s
}

func (s *RetrieveResponseBodyResultFaqs) SetRefType(v string) *RetrieveResponseBodyResultFaqs {
	s.RefType = &v
	return s
}

func (s *RetrieveResponseBodyResultFaqs) SetScore(v float64) *RetrieveResponseBodyResultFaqs {
	s.Score = &v
	return s
}

func (s *RetrieveResponseBodyResultFaqs) SetScoreItem(v *RetrieveResponseBodyResultFaqsScoreItem) *RetrieveResponseBodyResultFaqs {
	s.ScoreItem = v
	return s
}

func (s *RetrieveResponseBodyResultFaqs) SetTimestamp(v int64) *RetrieveResponseBodyResultFaqs {
	s.Timestamp = &v
	return s
}

func (s *RetrieveResponseBodyResultFaqs) SetTitle(v string) *RetrieveResponseBodyResultFaqs {
	s.Title = &v
	return s
}

func (s *RetrieveResponseBodyResultFaqs) SetType(v string) *RetrieveResponseBodyResultFaqs {
	s.Type = &v
	return s
}

func (s *RetrieveResponseBodyResultFaqs) SetUrl(v string) *RetrieveResponseBodyResultFaqs {
	s.Url = &v
	return s
}

type RetrieveResponseBodyResultFaqsScoreItem struct {
	FinallyScore     *float64            `json:"finallyScore,omitempty" xml:"finallyScore,omitempty"`
	ScoreMap         map[string]*float64 `json:"scoreMap,omitempty" xml:"scoreMap,omitempty"`
	ScoreThreshold   *float64            `json:"scoreThreshold,omitempty" xml:"scoreThreshold,omitempty"`
	SelectedBranch   []*string           `json:"selectedBranch,omitempty" xml:"selectedBranch,omitempty" type:"Repeated"`
	SelectedCategory *string             `json:"selectedCategory,omitempty" xml:"selectedCategory,omitempty"`
}

func (s RetrieveResponseBodyResultFaqsScoreItem) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBodyResultFaqsScoreItem) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyResultFaqsScoreItem) SetFinallyScore(v float64) *RetrieveResponseBodyResultFaqsScoreItem {
	s.FinallyScore = &v
	return s
}

func (s *RetrieveResponseBodyResultFaqsScoreItem) SetScoreMap(v map[string]*float64) *RetrieveResponseBodyResultFaqsScoreItem {
	s.ScoreMap = v
	return s
}

func (s *RetrieveResponseBodyResultFaqsScoreItem) SetScoreThreshold(v float64) *RetrieveResponseBodyResultFaqsScoreItem {
	s.ScoreThreshold = &v
	return s
}

func (s *RetrieveResponseBodyResultFaqsScoreItem) SetSelectedBranch(v []*string) *RetrieveResponseBodyResultFaqsScoreItem {
	s.SelectedBranch = v
	return s
}

func (s *RetrieveResponseBodyResultFaqsScoreItem) SetSelectedCategory(v string) *RetrieveResponseBodyResultFaqsScoreItem {
	s.SelectedCategory = &v
	return s
}

type RetrieveResponseBodyResultFinalResults struct {
	Icon          *string                                          `json:"icon,omitempty" xml:"icon,omitempty"`
	MatchIntimacy *bool                                            `json:"matchIntimacy,omitempty" xml:"matchIntimacy,omitempty"`
	Material      *string                                          `json:"material,omitempty" xml:"material,omitempty"`
	RefType       *string                                          `json:"refType,omitempty" xml:"refType,omitempty"`
	Score         *float64                                         `json:"score,omitempty" xml:"score,omitempty"`
	ScoreItem     *RetrieveResponseBodyResultFinalResultsScoreItem `json:"scoreItem,omitempty" xml:"scoreItem,omitempty" type:"Struct"`
	Timestamp     *int64                                           `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Title         *string                                          `json:"title,omitempty" xml:"title,omitempty"`
	Type          *string                                          `json:"type,omitempty" xml:"type,omitempty"`
	Url           *string                                          `json:"url,omitempty" xml:"url,omitempty"`
}

func (s RetrieveResponseBodyResultFinalResults) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBodyResultFinalResults) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyResultFinalResults) SetIcon(v string) *RetrieveResponseBodyResultFinalResults {
	s.Icon = &v
	return s
}

func (s *RetrieveResponseBodyResultFinalResults) SetMatchIntimacy(v bool) *RetrieveResponseBodyResultFinalResults {
	s.MatchIntimacy = &v
	return s
}

func (s *RetrieveResponseBodyResultFinalResults) SetMaterial(v string) *RetrieveResponseBodyResultFinalResults {
	s.Material = &v
	return s
}

func (s *RetrieveResponseBodyResultFinalResults) SetRefType(v string) *RetrieveResponseBodyResultFinalResults {
	s.RefType = &v
	return s
}

func (s *RetrieveResponseBodyResultFinalResults) SetScore(v float64) *RetrieveResponseBodyResultFinalResults {
	s.Score = &v
	return s
}

func (s *RetrieveResponseBodyResultFinalResults) SetScoreItem(v *RetrieveResponseBodyResultFinalResultsScoreItem) *RetrieveResponseBodyResultFinalResults {
	s.ScoreItem = v
	return s
}

func (s *RetrieveResponseBodyResultFinalResults) SetTimestamp(v int64) *RetrieveResponseBodyResultFinalResults {
	s.Timestamp = &v
	return s
}

func (s *RetrieveResponseBodyResultFinalResults) SetTitle(v string) *RetrieveResponseBodyResultFinalResults {
	s.Title = &v
	return s
}

func (s *RetrieveResponseBodyResultFinalResults) SetType(v string) *RetrieveResponseBodyResultFinalResults {
	s.Type = &v
	return s
}

func (s *RetrieveResponseBodyResultFinalResults) SetUrl(v string) *RetrieveResponseBodyResultFinalResults {
	s.Url = &v
	return s
}

type RetrieveResponseBodyResultFinalResultsScoreItem struct {
	FinallyScore     *float64            `json:"finallyScore,omitempty" xml:"finallyScore,omitempty"`
	ScoreMap         map[string]*float64 `json:"scoreMap,omitempty" xml:"scoreMap,omitempty"`
	ScoreThreshold   *float64            `json:"scoreThreshold,omitempty" xml:"scoreThreshold,omitempty"`
	SelectedBranch   []*string           `json:"selectedBranch,omitempty" xml:"selectedBranch,omitempty" type:"Repeated"`
	SelectedCategory *string             `json:"selectedCategory,omitempty" xml:"selectedCategory,omitempty"`
}

func (s RetrieveResponseBodyResultFinalResultsScoreItem) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBodyResultFinalResultsScoreItem) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyResultFinalResultsScoreItem) SetFinallyScore(v float64) *RetrieveResponseBodyResultFinalResultsScoreItem {
	s.FinallyScore = &v
	return s
}

func (s *RetrieveResponseBodyResultFinalResultsScoreItem) SetScoreMap(v map[string]*float64) *RetrieveResponseBodyResultFinalResultsScoreItem {
	s.ScoreMap = v
	return s
}

func (s *RetrieveResponseBodyResultFinalResultsScoreItem) SetScoreThreshold(v float64) *RetrieveResponseBodyResultFinalResultsScoreItem {
	s.ScoreThreshold = &v
	return s
}

func (s *RetrieveResponseBodyResultFinalResultsScoreItem) SetSelectedBranch(v []*string) *RetrieveResponseBodyResultFinalResultsScoreItem {
	s.SelectedBranch = v
	return s
}

func (s *RetrieveResponseBodyResultFinalResultsScoreItem) SetSelectedCategory(v string) *RetrieveResponseBodyResultFinalResultsScoreItem {
	s.SelectedCategory = &v
	return s
}

type RetrieveResponseBodyResultMinutes struct {
	CorpId          *int64                                      `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Creator         *string                                     `json:"creator,omitempty" xml:"creator,omitempty"`
	Extension       *string                                     `json:"extension,omitempty" xml:"extension,omitempty"`
	FullTextSummary *string                                     `json:"fullTextSummary,omitempty" xml:"fullTextSummary,omitempty"`
	GmtModified     *int64                                      `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Icon            *string                                     `json:"icon,omitempty" xml:"icon,omitempty"`
	MatchIntimacy   *bool                                       `json:"matchIntimacy,omitempty" xml:"matchIntimacy,omitempty"`
	Material        *string                                     `json:"material,omitempty" xml:"material,omitempty"`
	OriginText      *string                                     `json:"originText,omitempty" xml:"originText,omitempty"`
	RefType         *string                                     `json:"refType,omitempty" xml:"refType,omitempty"`
	Score           *float64                                    `json:"score,omitempty" xml:"score,omitempty"`
	ScoreItem       *RetrieveResponseBodyResultMinutesScoreItem `json:"scoreItem,omitempty" xml:"scoreItem,omitempty" type:"Struct"`
	Timestamp       *int64                                      `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Title           *string                                     `json:"title,omitempty" xml:"title,omitempty"`
	Type            *string                                     `json:"type,omitempty" xml:"type,omitempty"`
	Url             *string                                     `json:"url,omitempty" xml:"url,omitempty"`
}

func (s RetrieveResponseBodyResultMinutes) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBodyResultMinutes) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyResultMinutes) SetCorpId(v int64) *RetrieveResponseBodyResultMinutes {
	s.CorpId = &v
	return s
}

func (s *RetrieveResponseBodyResultMinutes) SetCreator(v string) *RetrieveResponseBodyResultMinutes {
	s.Creator = &v
	return s
}

func (s *RetrieveResponseBodyResultMinutes) SetExtension(v string) *RetrieveResponseBodyResultMinutes {
	s.Extension = &v
	return s
}

func (s *RetrieveResponseBodyResultMinutes) SetFullTextSummary(v string) *RetrieveResponseBodyResultMinutes {
	s.FullTextSummary = &v
	return s
}

func (s *RetrieveResponseBodyResultMinutes) SetGmtModified(v int64) *RetrieveResponseBodyResultMinutes {
	s.GmtModified = &v
	return s
}

func (s *RetrieveResponseBodyResultMinutes) SetIcon(v string) *RetrieveResponseBodyResultMinutes {
	s.Icon = &v
	return s
}

func (s *RetrieveResponseBodyResultMinutes) SetMatchIntimacy(v bool) *RetrieveResponseBodyResultMinutes {
	s.MatchIntimacy = &v
	return s
}

func (s *RetrieveResponseBodyResultMinutes) SetMaterial(v string) *RetrieveResponseBodyResultMinutes {
	s.Material = &v
	return s
}

func (s *RetrieveResponseBodyResultMinutes) SetOriginText(v string) *RetrieveResponseBodyResultMinutes {
	s.OriginText = &v
	return s
}

func (s *RetrieveResponseBodyResultMinutes) SetRefType(v string) *RetrieveResponseBodyResultMinutes {
	s.RefType = &v
	return s
}

func (s *RetrieveResponseBodyResultMinutes) SetScore(v float64) *RetrieveResponseBodyResultMinutes {
	s.Score = &v
	return s
}

func (s *RetrieveResponseBodyResultMinutes) SetScoreItem(v *RetrieveResponseBodyResultMinutesScoreItem) *RetrieveResponseBodyResultMinutes {
	s.ScoreItem = v
	return s
}

func (s *RetrieveResponseBodyResultMinutes) SetTimestamp(v int64) *RetrieveResponseBodyResultMinutes {
	s.Timestamp = &v
	return s
}

func (s *RetrieveResponseBodyResultMinutes) SetTitle(v string) *RetrieveResponseBodyResultMinutes {
	s.Title = &v
	return s
}

func (s *RetrieveResponseBodyResultMinutes) SetType(v string) *RetrieveResponseBodyResultMinutes {
	s.Type = &v
	return s
}

func (s *RetrieveResponseBodyResultMinutes) SetUrl(v string) *RetrieveResponseBodyResultMinutes {
	s.Url = &v
	return s
}

type RetrieveResponseBodyResultMinutesScoreItem struct {
	FinallyScore     *float64            `json:"finallyScore,omitempty" xml:"finallyScore,omitempty"`
	ScoreMap         map[string]*float64 `json:"scoreMap,omitempty" xml:"scoreMap,omitempty"`
	ScoreThreshold   *float64            `json:"scoreThreshold,omitempty" xml:"scoreThreshold,omitempty"`
	SelectedBranch   []*string           `json:"selectedBranch,omitempty" xml:"selectedBranch,omitempty" type:"Repeated"`
	SelectedCategory *string             `json:"selectedCategory,omitempty" xml:"selectedCategory,omitempty"`
}

func (s RetrieveResponseBodyResultMinutesScoreItem) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBodyResultMinutesScoreItem) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyResultMinutesScoreItem) SetFinallyScore(v float64) *RetrieveResponseBodyResultMinutesScoreItem {
	s.FinallyScore = &v
	return s
}

func (s *RetrieveResponseBodyResultMinutesScoreItem) SetScoreMap(v map[string]*float64) *RetrieveResponseBodyResultMinutesScoreItem {
	s.ScoreMap = v
	return s
}

func (s *RetrieveResponseBodyResultMinutesScoreItem) SetScoreThreshold(v float64) *RetrieveResponseBodyResultMinutesScoreItem {
	s.ScoreThreshold = &v
	return s
}

func (s *RetrieveResponseBodyResultMinutesScoreItem) SetSelectedBranch(v []*string) *RetrieveResponseBodyResultMinutesScoreItem {
	s.SelectedBranch = v
	return s
}

func (s *RetrieveResponseBodyResultMinutesScoreItem) SetSelectedCategory(v string) *RetrieveResponseBodyResultMinutesScoreItem {
	s.SelectedCategory = &v
	return s
}

type RetrieveResponseBodyResultReports struct {
	Content       *string                                     `json:"content,omitempty" xml:"content,omitempty"`
	CorpId        *int64                                      `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Creator       *string                                     `json:"creator,omitempty" xml:"creator,omitempty"`
	GmtCreate     *int64                                      `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified   *int64                                      `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Icon          *string                                     `json:"icon,omitempty" xml:"icon,omitempty"`
	MatchIntimacy *bool                                       `json:"matchIntimacy,omitempty" xml:"matchIntimacy,omitempty"`
	Material      *string                                     `json:"material,omitempty" xml:"material,omitempty"`
	RefType       *string                                     `json:"refType,omitempty" xml:"refType,omitempty"`
	Score         *float64                                    `json:"score,omitempty" xml:"score,omitempty"`
	ScoreItem     *RetrieveResponseBodyResultReportsScoreItem `json:"scoreItem,omitempty" xml:"scoreItem,omitempty" type:"Struct"`
	Timestamp     *int64                                      `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Title         *string                                     `json:"title,omitempty" xml:"title,omitempty"`
	Type          *string                                     `json:"type,omitempty" xml:"type,omitempty"`
	Url           *string                                     `json:"url,omitempty" xml:"url,omitempty"`
}

func (s RetrieveResponseBodyResultReports) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBodyResultReports) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyResultReports) SetContent(v string) *RetrieveResponseBodyResultReports {
	s.Content = &v
	return s
}

func (s *RetrieveResponseBodyResultReports) SetCorpId(v int64) *RetrieveResponseBodyResultReports {
	s.CorpId = &v
	return s
}

func (s *RetrieveResponseBodyResultReports) SetCreator(v string) *RetrieveResponseBodyResultReports {
	s.Creator = &v
	return s
}

func (s *RetrieveResponseBodyResultReports) SetGmtCreate(v int64) *RetrieveResponseBodyResultReports {
	s.GmtCreate = &v
	return s
}

func (s *RetrieveResponseBodyResultReports) SetGmtModified(v int64) *RetrieveResponseBodyResultReports {
	s.GmtModified = &v
	return s
}

func (s *RetrieveResponseBodyResultReports) SetIcon(v string) *RetrieveResponseBodyResultReports {
	s.Icon = &v
	return s
}

func (s *RetrieveResponseBodyResultReports) SetMatchIntimacy(v bool) *RetrieveResponseBodyResultReports {
	s.MatchIntimacy = &v
	return s
}

func (s *RetrieveResponseBodyResultReports) SetMaterial(v string) *RetrieveResponseBodyResultReports {
	s.Material = &v
	return s
}

func (s *RetrieveResponseBodyResultReports) SetRefType(v string) *RetrieveResponseBodyResultReports {
	s.RefType = &v
	return s
}

func (s *RetrieveResponseBodyResultReports) SetScore(v float64) *RetrieveResponseBodyResultReports {
	s.Score = &v
	return s
}

func (s *RetrieveResponseBodyResultReports) SetScoreItem(v *RetrieveResponseBodyResultReportsScoreItem) *RetrieveResponseBodyResultReports {
	s.ScoreItem = v
	return s
}

func (s *RetrieveResponseBodyResultReports) SetTimestamp(v int64) *RetrieveResponseBodyResultReports {
	s.Timestamp = &v
	return s
}

func (s *RetrieveResponseBodyResultReports) SetTitle(v string) *RetrieveResponseBodyResultReports {
	s.Title = &v
	return s
}

func (s *RetrieveResponseBodyResultReports) SetType(v string) *RetrieveResponseBodyResultReports {
	s.Type = &v
	return s
}

func (s *RetrieveResponseBodyResultReports) SetUrl(v string) *RetrieveResponseBodyResultReports {
	s.Url = &v
	return s
}

type RetrieveResponseBodyResultReportsScoreItem struct {
	FinallyScore     *float64            `json:"finallyScore,omitempty" xml:"finallyScore,omitempty"`
	ScoreMap         map[string]*float64 `json:"scoreMap,omitempty" xml:"scoreMap,omitempty"`
	ScoreThreshold   *float64            `json:"scoreThreshold,omitempty" xml:"scoreThreshold,omitempty"`
	SelectedBranch   []*string           `json:"selectedBranch,omitempty" xml:"selectedBranch,omitempty" type:"Repeated"`
	SelectedCategory *string             `json:"selectedCategory,omitempty" xml:"selectedCategory,omitempty"`
}

func (s RetrieveResponseBodyResultReportsScoreItem) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBodyResultReportsScoreItem) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyResultReportsScoreItem) SetFinallyScore(v float64) *RetrieveResponseBodyResultReportsScoreItem {
	s.FinallyScore = &v
	return s
}

func (s *RetrieveResponseBodyResultReportsScoreItem) SetScoreMap(v map[string]*float64) *RetrieveResponseBodyResultReportsScoreItem {
	s.ScoreMap = v
	return s
}

func (s *RetrieveResponseBodyResultReportsScoreItem) SetScoreThreshold(v float64) *RetrieveResponseBodyResultReportsScoreItem {
	s.ScoreThreshold = &v
	return s
}

func (s *RetrieveResponseBodyResultReportsScoreItem) SetSelectedBranch(v []*string) *RetrieveResponseBodyResultReportsScoreItem {
	s.SelectedBranch = v
	return s
}

func (s *RetrieveResponseBodyResultReportsScoreItem) SetSelectedCategory(v string) *RetrieveResponseBodyResultReportsScoreItem {
	s.SelectedCategory = &v
	return s
}

type RetrieveResponseBodyResultWikis struct {
	CorpId        *int64                                    `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Icon          *string                                   `json:"icon,omitempty" xml:"icon,omitempty"`
	MatchIntimacy *bool                                     `json:"matchIntimacy,omitempty" xml:"matchIntimacy,omitempty"`
	Material      *string                                   `json:"material,omitempty" xml:"material,omitempty"`
	RefType       *string                                   `json:"refType,omitempty" xml:"refType,omitempty"`
	Score         *float64                                  `json:"score,omitempty" xml:"score,omitempty"`
	ScoreItem     *RetrieveResponseBodyResultWikisScoreItem `json:"scoreItem,omitempty" xml:"scoreItem,omitempty" type:"Struct"`
	Timestamp     *int64                                    `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Title         *string                                   `json:"title,omitempty" xml:"title,omitempty"`
	Type          *string                                   `json:"type,omitempty" xml:"type,omitempty"`
	Url           *string                                   `json:"url,omitempty" xml:"url,omitempty"`
}

func (s RetrieveResponseBodyResultWikis) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBodyResultWikis) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyResultWikis) SetCorpId(v int64) *RetrieveResponseBodyResultWikis {
	s.CorpId = &v
	return s
}

func (s *RetrieveResponseBodyResultWikis) SetIcon(v string) *RetrieveResponseBodyResultWikis {
	s.Icon = &v
	return s
}

func (s *RetrieveResponseBodyResultWikis) SetMatchIntimacy(v bool) *RetrieveResponseBodyResultWikis {
	s.MatchIntimacy = &v
	return s
}

func (s *RetrieveResponseBodyResultWikis) SetMaterial(v string) *RetrieveResponseBodyResultWikis {
	s.Material = &v
	return s
}

func (s *RetrieveResponseBodyResultWikis) SetRefType(v string) *RetrieveResponseBodyResultWikis {
	s.RefType = &v
	return s
}

func (s *RetrieveResponseBodyResultWikis) SetScore(v float64) *RetrieveResponseBodyResultWikis {
	s.Score = &v
	return s
}

func (s *RetrieveResponseBodyResultWikis) SetScoreItem(v *RetrieveResponseBodyResultWikisScoreItem) *RetrieveResponseBodyResultWikis {
	s.ScoreItem = v
	return s
}

func (s *RetrieveResponseBodyResultWikis) SetTimestamp(v int64) *RetrieveResponseBodyResultWikis {
	s.Timestamp = &v
	return s
}

func (s *RetrieveResponseBodyResultWikis) SetTitle(v string) *RetrieveResponseBodyResultWikis {
	s.Title = &v
	return s
}

func (s *RetrieveResponseBodyResultWikis) SetType(v string) *RetrieveResponseBodyResultWikis {
	s.Type = &v
	return s
}

func (s *RetrieveResponseBodyResultWikis) SetUrl(v string) *RetrieveResponseBodyResultWikis {
	s.Url = &v
	return s
}

type RetrieveResponseBodyResultWikisScoreItem struct {
	FinallyScore     *float64            `json:"finallyScore,omitempty" xml:"finallyScore,omitempty"`
	ScoreMap         map[string]*float64 `json:"scoreMap,omitempty" xml:"scoreMap,omitempty"`
	ScoreThreshold   *float64            `json:"scoreThreshold,omitempty" xml:"scoreThreshold,omitempty"`
	SelectedBranch   []*string           `json:"selectedBranch,omitempty" xml:"selectedBranch,omitempty" type:"Repeated"`
	SelectedCategory *string             `json:"selectedCategory,omitempty" xml:"selectedCategory,omitempty"`
}

func (s RetrieveResponseBodyResultWikisScoreItem) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBodyResultWikisScoreItem) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyResultWikisScoreItem) SetFinallyScore(v float64) *RetrieveResponseBodyResultWikisScoreItem {
	s.FinallyScore = &v
	return s
}

func (s *RetrieveResponseBodyResultWikisScoreItem) SetScoreMap(v map[string]*float64) *RetrieveResponseBodyResultWikisScoreItem {
	s.ScoreMap = v
	return s
}

func (s *RetrieveResponseBodyResultWikisScoreItem) SetScoreThreshold(v float64) *RetrieveResponseBodyResultWikisScoreItem {
	s.ScoreThreshold = &v
	return s
}

func (s *RetrieveResponseBodyResultWikisScoreItem) SetSelectedBranch(v []*string) *RetrieveResponseBodyResultWikisScoreItem {
	s.SelectedBranch = v
	return s
}

func (s *RetrieveResponseBodyResultWikisScoreItem) SetSelectedCategory(v string) *RetrieveResponseBodyResultWikisScoreItem {
	s.SelectedCategory = &v
	return s
}

type RetrieveResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetrieveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetrieveResponse) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponse) GoString() string {
	return s.String()
}

func (s *RetrieveResponse) SetHeaders(v map[string]*string) *RetrieveResponse {
	s.Headers = v
	return s
}

func (s *RetrieveResponse) SetStatusCode(v int32) *RetrieveResponse {
	s.StatusCode = &v
	return s
}

func (s *RetrieveResponse) SetBody(v *RetrieveResponseBody) *RetrieveResponse {
	s.Body = v
	return s
}

type SubmitAiTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SubmitAiTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s SubmitAiTaskHeaders) GoString() string {
	return s.String()
}

func (s *SubmitAiTaskHeaders) SetCommonHeaders(v map[string]*string) *SubmitAiTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SubmitAiTaskHeaders) SetXAcsDingtalkAccessToken(v string) *SubmitAiTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SubmitAiTaskRequest struct {
	Messages []*SubmitAiTaskRequestMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
}

func (s SubmitAiTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitAiTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitAiTaskRequest) SetMessages(v []*SubmitAiTaskRequestMessages) *SubmitAiTaskRequest {
	s.Messages = v
	return s
}

type SubmitAiTaskRequestMessages struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s SubmitAiTaskRequestMessages) String() string {
	return tea.Prettify(s)
}

func (s SubmitAiTaskRequestMessages) GoString() string {
	return s.String()
}

func (s *SubmitAiTaskRequestMessages) SetContent(v string) *SubmitAiTaskRequestMessages {
	s.Content = &v
	return s
}

type SubmitAiTaskShrinkRequest struct {
	MessagesShrink *string `json:"messages,omitempty" xml:"messages,omitempty"`
}

func (s SubmitAiTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitAiTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitAiTaskShrinkRequest) SetMessagesShrink(v string) *SubmitAiTaskShrinkRequest {
	s.MessagesShrink = &v
	return s
}

type SubmitAiTaskResponseBody struct {
	ErrorCode *string                         `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                         `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result    *SubmitAiTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success   *bool                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SubmitAiTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitAiTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAiTaskResponseBody) SetErrorCode(v string) *SubmitAiTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitAiTaskResponseBody) SetErrorMsg(v string) *SubmitAiTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SubmitAiTaskResponseBody) SetResult(v *SubmitAiTaskResponseBodyResult) *SubmitAiTaskResponseBody {
	s.Result = v
	return s
}

func (s *SubmitAiTaskResponseBody) SetSuccess(v bool) *SubmitAiTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitAiTaskResponseBodyResult struct {
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SubmitAiTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SubmitAiTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SubmitAiTaskResponseBodyResult) SetTaskId(v string) *SubmitAiTaskResponseBodyResult {
	s.TaskId = &v
	return s
}

type SubmitAiTaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAiTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAiTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitAiTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitAiTaskResponse) SetHeaders(v map[string]*string) *SubmitAiTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitAiTaskResponse) SetStatusCode(v int32) *SubmitAiTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAiTaskResponse) SetBody(v *SubmitAiTaskResponseBody) *SubmitAiTaskResponse {
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
// 查询指定用户的全部已登录设备及其状态
//
// @param request - FetchLoginStatusDevicesRequest
//
// @param headers - FetchLoginStatusDevicesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FetchLoginStatusDevicesResponse
func (client *Client) FetchLoginStatusDevicesWithOptions(request *FetchLoginStatusDevicesRequest, headers *FetchLoginStatusDevicesHeaders, runtime *util.RuntimeOptions) (_result *FetchLoginStatusDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		body["domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdList)) {
		body["userIdList"] = request.UserIdList
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
		Action:      tea.String("FetchLoginStatusDevices"),
		Version:     tea.String("aiSearchAsk_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiSearchAsk/fetchLoginStatusDevices"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &FetchLoginStatusDevicesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定用户的全部已登录设备及其状态
//
// @param request - FetchLoginStatusDevicesRequest
//
// @return FetchLoginStatusDevicesResponse
func (client *Client) FetchLoginStatusDevices(request *FetchLoginStatusDevicesRequest) (_result *FetchLoginStatusDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &FetchLoginStatusDevicesHeaders{}
	_result = &FetchLoginStatusDevicesResponse{}
	_body, _err := client.FetchLoginStatusDevicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获得ai任务结果
//
// @param request - GetAiTaskResultRequest
//
// @param headers - GetAiTaskResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAiTaskResultResponse
func (client *Client) GetAiTaskResultWithOptions(request *GetAiTaskResultRequest, headers *GetAiTaskResultHeaders, runtime *util.RuntimeOptions) (_result *GetAiTaskResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
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
		Action:      tea.String("GetAiTaskResult"),
		Version:     tea.String("aiSearchAsk_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiSearchAsk/getAiTaskResult"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAiTaskResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得ai任务结果
//
// @param request - GetAiTaskResultRequest
//
// @return GetAiTaskResultResponse
func (client *Client) GetAiTaskResult(request *GetAiTaskResultRequest) (_result *GetAiTaskResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAiTaskResultHeaders{}
	_result = &GetAiTaskResultResponse{}
	_body, _err := client.GetAiTaskResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 多数据源召回接口
//
// @param request - RetrieveRequest
//
// @param headers - RetrieveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetrieveResponse
func (client *Client) RetrieveWithOptions(request *RetrieveRequest, headers *RetrieveHeaders, runtime *util.RuntimeOptions) (_result *RetrieveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Answerer)) {
		body["answerer"] = request.Answerer
	}

	if !tea.BoolValue(util.IsUnset(request.DragRequestContext)) {
		body["dragRequestContext"] = request.DragRequestContext
	}

	if !tea.BoolValue(util.IsUnset(request.KeywordList)) {
		body["keywordList"] = request.KeywordList
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Question)) {
		body["question"] = request.Question
	}

	if !tea.BoolValue(util.IsUnset(request.RetrievalExtendParams)) {
		body["retrievalExtendParams"] = request.RetrievalExtendParams
	}

	if !tea.BoolValue(util.IsUnset(request.RetrieveScoreThreshold)) {
		body["retrieveScoreThreshold"] = request.RetrieveScoreThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.Tenant)) {
		body["tenant"] = request.Tenant
	}

	if !tea.BoolValue(util.IsUnset(request.TokenLimit)) {
		body["tokenLimit"] = request.TokenLimit
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
		Action:      tea.String("Retrieve"),
		Version:     tea.String("aiSearchAsk_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiSearchAsk/retrieve"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RetrieveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 多数据源召回接口
//
// @param request - RetrieveRequest
//
// @return RetrieveResponse
func (client *Client) Retrieve(request *RetrieveRequest) (_result *RetrieveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RetrieveHeaders{}
	_result = &RetrieveResponse{}
	_body, _err := client.RetrieveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交ai任务
//
// @param tmpReq - SubmitAiTaskRequest
//
// @param headers - SubmitAiTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAiTaskResponse
func (client *Client) SubmitAiTaskWithOptions(tmpReq *SubmitAiTaskRequest, headers *SubmitAiTaskHeaders, runtime *util.RuntimeOptions) (_result *SubmitAiTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SubmitAiTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Messages)) {
		request.MessagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Messages, tea.String("messages"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MessagesShrink)) {
		query["messages"] = request.MessagesShrink
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
		Action:      tea.String("SubmitAiTask"),
		Version:     tea.String("aiSearchAsk_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiSearchAsk/submitAiTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitAiTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交ai任务
//
// @param request - SubmitAiTaskRequest
//
// @return SubmitAiTaskResponse
func (client *Client) SubmitAiTask(request *SubmitAiTaskRequest) (_result *SubmitAiTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SubmitAiTaskHeaders{}
	_result = &SubmitAiTaskResponse{}
	_body, _err := client.SubmitAiTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
