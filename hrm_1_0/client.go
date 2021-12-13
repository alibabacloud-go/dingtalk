// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package hrm_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ECertQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ECertQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s ECertQueryHeaders) GoString() string {
	return s.String()
}

func (s *ECertQueryHeaders) SetCommonHeaders(v map[string]*string) *ECertQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ECertQueryHeaders) SetXAcsDingtalkAccessToken(v string) *ECertQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ECertQueryRequest struct {
	// 用户ID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ECertQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s ECertQueryRequest) GoString() string {
	return s.String()
}

func (s *ECertQueryRequest) SetUserId(v string) *ECertQueryRequest {
	s.UserId = &v
	return s
}

type ECertQueryResponseBody struct {
	// 身份证姓名
	RealName *string `json:"realName,omitempty" xml:"realName,omitempty"`
	// 身份证号码
	CertNO *string `json:"certNO,omitempty" xml:"certNO,omitempty"`
	// 主部门ID
	MainDeptId *int64 `json:"mainDeptId,omitempty" xml:"mainDeptId,omitempty"`
	// 主部门
	MainDeptName *string `json:"mainDeptName,omitempty" xml:"mainDeptName,omitempty"`
	// 职务ID
	EmployJobId *string `json:"employJobId,omitempty" xml:"employJobId,omitempty"`
	// 职务名称
	EmployJobIdLabel *string `json:"employJobIdLabel,omitempty" xml:"employJobIdLabel,omitempty"`
	// 职位ID
	EmployPositionId *string `json:"employPositionId,omitempty" xml:"employPositionId,omitempty"`
	// 职位名称
	EmployPositionIdLabel *string `json:"employPositionIdLabel,omitempty" xml:"employPositionIdLabel,omitempty"`
	// 职级ID
	EmployPositionRankId *string `json:"employPositionRankId,omitempty" xml:"employPositionRankId,omitempty"`
	// 职级名称
	EmployPositionRankIdLabel *string `json:"employPositionRankIdLabel,omitempty" xml:"employPositionRankIdLabel,omitempty"`
	// 入职日期
	HiredDate *string `json:"hiredDate,omitempty" xml:"hiredDate,omitempty"`
	// 离职日期
	LastWorkDay *string `json:"lastWorkDay,omitempty" xml:"lastWorkDay,omitempty"`
	// 主动离职原因
	TerminationReasonVoluntary []*string `json:"terminationReasonVoluntary,omitempty" xml:"terminationReasonVoluntary,omitempty" type:"Repeated"`
	// 被动离职原因
	TerminationReasonPassive []*string `json:"terminationReasonPassive,omitempty" xml:"terminationReasonPassive,omitempty" type:"Repeated"`
	// 姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ECertQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ECertQueryResponseBody) GoString() string {
	return s.String()
}

func (s *ECertQueryResponseBody) SetRealName(v string) *ECertQueryResponseBody {
	s.RealName = &v
	return s
}

func (s *ECertQueryResponseBody) SetCertNO(v string) *ECertQueryResponseBody {
	s.CertNO = &v
	return s
}

func (s *ECertQueryResponseBody) SetMainDeptId(v int64) *ECertQueryResponseBody {
	s.MainDeptId = &v
	return s
}

func (s *ECertQueryResponseBody) SetMainDeptName(v string) *ECertQueryResponseBody {
	s.MainDeptName = &v
	return s
}

func (s *ECertQueryResponseBody) SetEmployJobId(v string) *ECertQueryResponseBody {
	s.EmployJobId = &v
	return s
}

func (s *ECertQueryResponseBody) SetEmployJobIdLabel(v string) *ECertQueryResponseBody {
	s.EmployJobIdLabel = &v
	return s
}

func (s *ECertQueryResponseBody) SetEmployPositionId(v string) *ECertQueryResponseBody {
	s.EmployPositionId = &v
	return s
}

func (s *ECertQueryResponseBody) SetEmployPositionIdLabel(v string) *ECertQueryResponseBody {
	s.EmployPositionIdLabel = &v
	return s
}

func (s *ECertQueryResponseBody) SetEmployPositionRankId(v string) *ECertQueryResponseBody {
	s.EmployPositionRankId = &v
	return s
}

func (s *ECertQueryResponseBody) SetEmployPositionRankIdLabel(v string) *ECertQueryResponseBody {
	s.EmployPositionRankIdLabel = &v
	return s
}

func (s *ECertQueryResponseBody) SetHiredDate(v string) *ECertQueryResponseBody {
	s.HiredDate = &v
	return s
}

func (s *ECertQueryResponseBody) SetLastWorkDay(v string) *ECertQueryResponseBody {
	s.LastWorkDay = &v
	return s
}

func (s *ECertQueryResponseBody) SetTerminationReasonVoluntary(v []*string) *ECertQueryResponseBody {
	s.TerminationReasonVoluntary = v
	return s
}

func (s *ECertQueryResponseBody) SetTerminationReasonPassive(v []*string) *ECertQueryResponseBody {
	s.TerminationReasonPassive = v
	return s
}

func (s *ECertQueryResponseBody) SetName(v string) *ECertQueryResponseBody {
	s.Name = &v
	return s
}

type ECertQueryResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ECertQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ECertQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s ECertQueryResponse) GoString() string {
	return s.String()
}

func (s *ECertQueryResponse) SetHeaders(v map[string]*string) *ECertQueryResponse {
	s.Headers = v
	return s
}

func (s *ECertQueryResponse) SetBody(v *ECertQueryResponseBody) *ECertQueryResponse {
	s.Body = v
	return s
}

type QueryJobRanksHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryJobRanksHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryJobRanksHeaders) GoString() string {
	return s.String()
}

func (s *QueryJobRanksHeaders) SetCommonHeaders(v map[string]*string) *QueryJobRanksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryJobRanksHeaders) SetXAcsDingtalkAccessToken(v string) *QueryJobRanksHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryJobRanksRequest struct {
	// 职级序列
	RankCategoryId *string `json:"rankCategoryId,omitempty" xml:"rankCategoryId,omitempty"`
	// 职级编码
	RankCode *string `json:"rankCode,omitempty" xml:"rankCode,omitempty"`
	// 职级名称
	RankName *string `json:"rankName,omitempty" xml:"rankName,omitempty"`
	// 标记当前开始读取的位置
	NextToken *int32 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 本次读取的最大数据记录数量
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
}

func (s QueryJobRanksRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryJobRanksRequest) GoString() string {
	return s.String()
}

func (s *QueryJobRanksRequest) SetRankCategoryId(v string) *QueryJobRanksRequest {
	s.RankCategoryId = &v
	return s
}

func (s *QueryJobRanksRequest) SetRankCode(v string) *QueryJobRanksRequest {
	s.RankCode = &v
	return s
}

func (s *QueryJobRanksRequest) SetRankName(v string) *QueryJobRanksRequest {
	s.RankName = &v
	return s
}

func (s *QueryJobRanksRequest) SetNextToken(v int32) *QueryJobRanksRequest {
	s.NextToken = &v
	return s
}

func (s *QueryJobRanksRequest) SetMaxResults(v int32) *QueryJobRanksRequest {
	s.MaxResults = &v
	return s
}

type QueryJobRanksResponseBody struct {
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 是否有更多数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 职级列表
	List []*QueryJobRanksResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s QueryJobRanksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryJobRanksResponseBody) GoString() string {
	return s.String()
}

func (s *QueryJobRanksResponseBody) SetNextToken(v int64) *QueryJobRanksResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryJobRanksResponseBody) SetHasMore(v bool) *QueryJobRanksResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryJobRanksResponseBody) SetList(v []*QueryJobRanksResponseBodyList) *QueryJobRanksResponseBody {
	s.List = v
	return s
}

type QueryJobRanksResponseBodyList struct {
	// 职级ID
	RankId *string `json:"rankId,omitempty" xml:"rankId,omitempty"`
	// 职级序列ID
	RankCategoryId *string `json:"rankCategoryId,omitempty" xml:"rankCategoryId,omitempty"`
	// 职级编码
	RankCode *string `json:"rankCode,omitempty" xml:"rankCode,omitempty"`
	// 职级名称
	RankName *string `json:"rankName,omitempty" xml:"rankName,omitempty"`
	// 最小等级
	MinJobGrade *int32 `json:"minJobGrade,omitempty" xml:"minJobGrade,omitempty"`
	// 最大等级
	MaxJobGrade *int32 `json:"maxJobGrade,omitempty" xml:"maxJobGrade,omitempty"`
	// 职级描述
	RankDescription *string `json:"rankDescription,omitempty" xml:"rankDescription,omitempty"`
}

func (s QueryJobRanksResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryJobRanksResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryJobRanksResponseBodyList) SetRankId(v string) *QueryJobRanksResponseBodyList {
	s.RankId = &v
	return s
}

func (s *QueryJobRanksResponseBodyList) SetRankCategoryId(v string) *QueryJobRanksResponseBodyList {
	s.RankCategoryId = &v
	return s
}

func (s *QueryJobRanksResponseBodyList) SetRankCode(v string) *QueryJobRanksResponseBodyList {
	s.RankCode = &v
	return s
}

func (s *QueryJobRanksResponseBodyList) SetRankName(v string) *QueryJobRanksResponseBodyList {
	s.RankName = &v
	return s
}

func (s *QueryJobRanksResponseBodyList) SetMinJobGrade(v int32) *QueryJobRanksResponseBodyList {
	s.MinJobGrade = &v
	return s
}

func (s *QueryJobRanksResponseBodyList) SetMaxJobGrade(v int32) *QueryJobRanksResponseBodyList {
	s.MaxJobGrade = &v
	return s
}

func (s *QueryJobRanksResponseBodyList) SetRankDescription(v string) *QueryJobRanksResponseBodyList {
	s.RankDescription = &v
	return s
}

type QueryJobRanksResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryJobRanksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryJobRanksResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryJobRanksResponse) GoString() string {
	return s.String()
}

func (s *QueryJobRanksResponse) SetHeaders(v map[string]*string) *QueryJobRanksResponse {
	s.Headers = v
	return s
}

func (s *QueryJobRanksResponse) SetBody(v *QueryJobRanksResponseBody) *QueryJobRanksResponse {
	s.Body = v
	return s
}

type QueryJobsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryJobsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryJobsHeaders) GoString() string {
	return s.String()
}

func (s *QueryJobsHeaders) SetCommonHeaders(v map[string]*string) *QueryJobsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryJobsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryJobsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryJobsRequest struct {
	// 职务名称
	JobName *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
	// 偏移量
	NextToken *int32 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 最大值
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
}

func (s QueryJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryJobsRequest) GoString() string {
	return s.String()
}

func (s *QueryJobsRequest) SetJobName(v string) *QueryJobsRequest {
	s.JobName = &v
	return s
}

func (s *QueryJobsRequest) SetNextToken(v int32) *QueryJobsRequest {
	s.NextToken = &v
	return s
}

func (s *QueryJobsRequest) SetMaxResults(v int32) *QueryJobsRequest {
	s.MaxResults = &v
	return s
}

type QueryJobsResponseBody struct {
	// 下次获取数据的起始游标
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 是否有更多数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 职务列表
	List []*QueryJobsResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s QueryJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryJobsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryJobsResponseBody) SetNextToken(v int64) *QueryJobsResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryJobsResponseBody) SetHasMore(v bool) *QueryJobsResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryJobsResponseBody) SetList(v []*QueryJobsResponseBodyList) *QueryJobsResponseBody {
	s.List = v
	return s
}

type QueryJobsResponseBodyList struct {
	// 职务ID
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// 职务名称
	JobName *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
	// 职务描述
	JobDescription *string `json:"jobDescription,omitempty" xml:"jobDescription,omitempty"`
}

func (s QueryJobsResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryJobsResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryJobsResponseBodyList) SetJobId(v string) *QueryJobsResponseBodyList {
	s.JobId = &v
	return s
}

func (s *QueryJobsResponseBodyList) SetJobName(v string) *QueryJobsResponseBodyList {
	s.JobName = &v
	return s
}

func (s *QueryJobsResponseBodyList) SetJobDescription(v string) *QueryJobsResponseBodyList {
	s.JobDescription = &v
	return s
}

type QueryJobsResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryJobsResponse) GoString() string {
	return s.String()
}

func (s *QueryJobsResponse) SetHeaders(v map[string]*string) *QueryJobsResponse {
	s.Headers = v
	return s
}

func (s *QueryJobsResponse) SetBody(v *QueryJobsResponseBody) *QueryJobsResponse {
	s.Body = v
	return s
}

type QueryCustomEntryProcessesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCustomEntryProcessesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomEntryProcessesHeaders) GoString() string {
	return s.String()
}

func (s *QueryCustomEntryProcessesHeaders) SetCommonHeaders(v map[string]*string) *QueryCustomEntryProcessesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCustomEntryProcessesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCustomEntryProcessesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCustomEntryProcessesRequest struct {
	// 操作人id
	OperateUserId *string `json:"operateUserId,omitempty" xml:"operateUserId,omitempty"`
	// 偏移量
	NextToken *int32 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 最大值
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
}

func (s QueryCustomEntryProcessesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomEntryProcessesRequest) GoString() string {
	return s.String()
}

func (s *QueryCustomEntryProcessesRequest) SetOperateUserId(v string) *QueryCustomEntryProcessesRequest {
	s.OperateUserId = &v
	return s
}

func (s *QueryCustomEntryProcessesRequest) SetNextToken(v int32) *QueryCustomEntryProcessesRequest {
	s.NextToken = &v
	return s
}

func (s *QueryCustomEntryProcessesRequest) SetMaxResults(v int32) *QueryCustomEntryProcessesRequest {
	s.MaxResults = &v
	return s
}

type QueryCustomEntryProcessesResponseBody struct {
	// 下次获取数据的起始游标
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 是否有更多
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 表单信息列表
	List []*QueryCustomEntryProcessesResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s QueryCustomEntryProcessesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomEntryProcessesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCustomEntryProcessesResponseBody) SetNextToken(v int64) *QueryCustomEntryProcessesResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryCustomEntryProcessesResponseBody) SetHasMore(v bool) *QueryCustomEntryProcessesResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryCustomEntryProcessesResponseBody) SetList(v []*QueryCustomEntryProcessesResponseBodyList) *QueryCustomEntryProcessesResponseBody {
	s.List = v
	return s
}

type QueryCustomEntryProcessesResponseBodyList struct {
	FormId   *string `json:"formId,omitempty" xml:"formId,omitempty"`
	FormName *string `json:"formName,omitempty" xml:"formName,omitempty"`
	FormDesc *string `json:"formDesc,omitempty" xml:"formDesc,omitempty"`
	ShortUrl *string `json:"shortUrl,omitempty" xml:"shortUrl,omitempty"`
}

func (s QueryCustomEntryProcessesResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomEntryProcessesResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryCustomEntryProcessesResponseBodyList) SetFormId(v string) *QueryCustomEntryProcessesResponseBodyList {
	s.FormId = &v
	return s
}

func (s *QueryCustomEntryProcessesResponseBodyList) SetFormName(v string) *QueryCustomEntryProcessesResponseBodyList {
	s.FormName = &v
	return s
}

func (s *QueryCustomEntryProcessesResponseBodyList) SetFormDesc(v string) *QueryCustomEntryProcessesResponseBodyList {
	s.FormDesc = &v
	return s
}

func (s *QueryCustomEntryProcessesResponseBodyList) SetShortUrl(v string) *QueryCustomEntryProcessesResponseBodyList {
	s.ShortUrl = &v
	return s
}

type QueryCustomEntryProcessesResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCustomEntryProcessesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCustomEntryProcessesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomEntryProcessesResponse) GoString() string {
	return s.String()
}

func (s *QueryCustomEntryProcessesResponse) SetHeaders(v map[string]*string) *QueryCustomEntryProcessesResponse {
	s.Headers = v
	return s
}

func (s *QueryCustomEntryProcessesResponse) SetBody(v *QueryCustomEntryProcessesResponseBody) *QueryCustomEntryProcessesResponse {
	s.Body = v
	return s
}

type QueryPositionsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryPositionsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryPositionsHeaders) GoString() string {
	return s.String()
}

func (s *QueryPositionsHeaders) SetCommonHeaders(v map[string]*string) *QueryPositionsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryPositionsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryPositionsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryPositionsRequest struct {
	// 职位名称
	PositionName *string `json:"positionName,omitempty" xml:"positionName,omitempty"`
	// 职位类别列表
	InCategoryIds []*string `json:"inCategoryIds,omitempty" xml:"inCategoryIds,omitempty" type:"Repeated"`
	// 职位id列表
	InPositionIds []*string `json:"inPositionIds,omitempty" xml:"inPositionIds,omitempty" type:"Repeated"`
	// 偏移量
	NextToken *int32 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 一次查询获取记录数
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
}

func (s QueryPositionsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPositionsRequest) GoString() string {
	return s.String()
}

func (s *QueryPositionsRequest) SetPositionName(v string) *QueryPositionsRequest {
	s.PositionName = &v
	return s
}

func (s *QueryPositionsRequest) SetInCategoryIds(v []*string) *QueryPositionsRequest {
	s.InCategoryIds = v
	return s
}

func (s *QueryPositionsRequest) SetInPositionIds(v []*string) *QueryPositionsRequest {
	s.InPositionIds = v
	return s
}

func (s *QueryPositionsRequest) SetNextToken(v int32) *QueryPositionsRequest {
	s.NextToken = &v
	return s
}

func (s *QueryPositionsRequest) SetMaxResults(v int32) *QueryPositionsRequest {
	s.MaxResults = &v
	return s
}

type QueryPositionsResponseBody struct {
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 是否有更多数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 职位列表
	List []*QueryPositionsResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s QueryPositionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPositionsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPositionsResponseBody) SetNextToken(v int64) *QueryPositionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryPositionsResponseBody) SetHasMore(v bool) *QueryPositionsResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryPositionsResponseBody) SetList(v []*QueryPositionsResponseBodyList) *QueryPositionsResponseBody {
	s.List = v
	return s
}

type QueryPositionsResponseBodyList struct {
	// 职位ID
	PositionId *string `json:"positionId,omitempty" xml:"positionId,omitempty"`
	// 职位名称
	PositionName *string `json:"positionName,omitempty" xml:"positionName,omitempty"`
	// 职位类别ID
	PositionCategoryId *string `json:"positionCategoryId,omitempty" xml:"positionCategoryId,omitempty"`
	// 所属职务ID
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// 职位描述
	PositionDes *string `json:"positionDes,omitempty" xml:"positionDes,omitempty"`
	// 职位对应职级列表
	RankIdList []*string `json:"rankIdList,omitempty" xml:"rankIdList,omitempty" type:"Repeated"`
	// 职位状态-0，启用；1，停用
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryPositionsResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryPositionsResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryPositionsResponseBodyList) SetPositionId(v string) *QueryPositionsResponseBodyList {
	s.PositionId = &v
	return s
}

func (s *QueryPositionsResponseBodyList) SetPositionName(v string) *QueryPositionsResponseBodyList {
	s.PositionName = &v
	return s
}

func (s *QueryPositionsResponseBodyList) SetPositionCategoryId(v string) *QueryPositionsResponseBodyList {
	s.PositionCategoryId = &v
	return s
}

func (s *QueryPositionsResponseBodyList) SetJobId(v string) *QueryPositionsResponseBodyList {
	s.JobId = &v
	return s
}

func (s *QueryPositionsResponseBodyList) SetPositionDes(v string) *QueryPositionsResponseBodyList {
	s.PositionDes = &v
	return s
}

func (s *QueryPositionsResponseBodyList) SetRankIdList(v []*string) *QueryPositionsResponseBodyList {
	s.RankIdList = v
	return s
}

func (s *QueryPositionsResponseBodyList) SetStatus(v int32) *QueryPositionsResponseBodyList {
	s.Status = &v
	return s
}

type QueryPositionsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryPositionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPositionsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPositionsResponse) GoString() string {
	return s.String()
}

func (s *QueryPositionsResponse) SetHeaders(v map[string]*string) *QueryPositionsResponse {
	s.Headers = v
	return s
}

func (s *QueryPositionsResponse) SetBody(v *QueryPositionsResponseBody) *QueryPositionsResponse {
	s.Body = v
	return s
}

type MasterDataQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MasterDataQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryHeaders) GoString() string {
	return s.String()
}

func (s *MasterDataQueryHeaders) SetCommonHeaders(v map[string]*string) *MasterDataQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MasterDataQueryHeaders) SetXAcsDingtalkAccessToken(v string) *MasterDataQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MasterDataQueryRequest struct {
	// 领域code 由钉钉分配
	ScopeCode *string `json:"scopeCode,omitempty" xml:"scopeCode,omitempty"`
	// 实体code
	ViewEntityCode *string `json:"viewEntityCode,omitempty" xml:"viewEntityCode,omitempty"`
	// 数据生产方的租户id，由钉钉分配
	TenantId *int64 `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 数据唯一键
	BizUK *string `json:"bizUK,omitempty" xml:"bizUK,omitempty"`
	// 关联id列表，一般为userId
	RelationIds []*string `json:"relationIds,omitempty" xml:"relationIds,omitempty" type:"Repeated"`
	// 当前操作人userId
	OptUserId *string `json:"optUserId,omitempty" xml:"optUserId,omitempty"`
	// 分页查询的游标
	NextToken *int32 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 分页查询每页数据条数
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 其他查询条件
	QueryParams []*MasterDataQueryRequestQueryParams `json:"queryParams,omitempty" xml:"queryParams,omitempty" type:"Repeated"`
}

func (s MasterDataQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryRequest) GoString() string {
	return s.String()
}

func (s *MasterDataQueryRequest) SetScopeCode(v string) *MasterDataQueryRequest {
	s.ScopeCode = &v
	return s
}

func (s *MasterDataQueryRequest) SetViewEntityCode(v string) *MasterDataQueryRequest {
	s.ViewEntityCode = &v
	return s
}

func (s *MasterDataQueryRequest) SetTenantId(v int64) *MasterDataQueryRequest {
	s.TenantId = &v
	return s
}

func (s *MasterDataQueryRequest) SetBizUK(v string) *MasterDataQueryRequest {
	s.BizUK = &v
	return s
}

func (s *MasterDataQueryRequest) SetRelationIds(v []*string) *MasterDataQueryRequest {
	s.RelationIds = v
	return s
}

func (s *MasterDataQueryRequest) SetOptUserId(v string) *MasterDataQueryRequest {
	s.OptUserId = &v
	return s
}

func (s *MasterDataQueryRequest) SetNextToken(v int32) *MasterDataQueryRequest {
	s.NextToken = &v
	return s
}

func (s *MasterDataQueryRequest) SetMaxResults(v int32) *MasterDataQueryRequest {
	s.MaxResults = &v
	return s
}

func (s *MasterDataQueryRequest) SetQueryParams(v []*MasterDataQueryRequestQueryParams) *MasterDataQueryRequest {
	s.QueryParams = v
	return s
}

type MasterDataQueryRequestQueryParams struct {
	// 需要筛选的字段
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	// 筛选条件连接类型
	JoinType *string `json:"joinType,omitempty" xml:"joinType,omitempty"`
	// 筛选条件
	ConditionList []*MasterDataQueryRequestQueryParamsConditionList `json:"conditionList,omitempty" xml:"conditionList,omitempty" type:"Repeated"`
}

func (s MasterDataQueryRequestQueryParams) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryRequestQueryParams) GoString() string {
	return s.String()
}

func (s *MasterDataQueryRequestQueryParams) SetFieldCode(v string) *MasterDataQueryRequestQueryParams {
	s.FieldCode = &v
	return s
}

func (s *MasterDataQueryRequestQueryParams) SetJoinType(v string) *MasterDataQueryRequestQueryParams {
	s.JoinType = &v
	return s
}

func (s *MasterDataQueryRequestQueryParams) SetConditionList(v []*MasterDataQueryRequestQueryParamsConditionList) *MasterDataQueryRequestQueryParams {
	s.ConditionList = v
	return s
}

type MasterDataQueryRequestQueryParamsConditionList struct {
	// 字段关系符
	Operate *string `json:"operate,omitempty" xml:"operate,omitempty"`
	// 操作值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s MasterDataQueryRequestQueryParamsConditionList) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryRequestQueryParamsConditionList) GoString() string {
	return s.String()
}

func (s *MasterDataQueryRequestQueryParamsConditionList) SetOperate(v string) *MasterDataQueryRequestQueryParamsConditionList {
	s.Operate = &v
	return s
}

func (s *MasterDataQueryRequestQueryParamsConditionList) SetValue(v string) *MasterDataQueryRequestQueryParamsConditionList {
	s.Value = &v
	return s
}

type MasterDataQueryResponseBody struct {
	// 总条目数
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 是否还有更多
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 分页游标
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 结果
	Result []*MasterDataQueryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s MasterDataQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryResponseBody) GoString() string {
	return s.String()
}

func (s *MasterDataQueryResponseBody) SetTotal(v int64) *MasterDataQueryResponseBody {
	s.Total = &v
	return s
}

func (s *MasterDataQueryResponseBody) SetHasMore(v bool) *MasterDataQueryResponseBody {
	s.HasMore = &v
	return s
}

func (s *MasterDataQueryResponseBody) SetNextToken(v int64) *MasterDataQueryResponseBody {
	s.NextToken = &v
	return s
}

func (s *MasterDataQueryResponseBody) SetSuccess(v bool) *MasterDataQueryResponseBody {
	s.Success = &v
	return s
}

func (s *MasterDataQueryResponseBody) SetResult(v []*MasterDataQueryResponseBodyResult) *MasterDataQueryResponseBody {
	s.Result = v
	return s
}

type MasterDataQueryResponseBodyResult struct {
	// 唯一id
	OuterId *string `json:"outerId,omitempty" xml:"outerId,omitempty"`
	// 领域
	ScopeCode *string `json:"scopeCode,omitempty" xml:"scopeCode,omitempty"`
	// 编码
	ViewEntityCode *string `json:"viewEntityCode,omitempty" xml:"viewEntityCode,omitempty"`
	// 字段列表
	ViewEntityFieldVOList []*MasterDataQueryResponseBodyResultViewEntityFieldVOList `json:"viewEntityFieldVOList,omitempty" xml:"viewEntityFieldVOList,omitempty" type:"Repeated"`
	// 关联id列表，一般为userId
	RelationId *string `json:"relationId,omitempty" xml:"relationId,omitempty"`
}

func (s MasterDataQueryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *MasterDataQueryResponseBodyResult) SetOuterId(v string) *MasterDataQueryResponseBodyResult {
	s.OuterId = &v
	return s
}

func (s *MasterDataQueryResponseBodyResult) SetScopeCode(v string) *MasterDataQueryResponseBodyResult {
	s.ScopeCode = &v
	return s
}

func (s *MasterDataQueryResponseBodyResult) SetViewEntityCode(v string) *MasterDataQueryResponseBodyResult {
	s.ViewEntityCode = &v
	return s
}

func (s *MasterDataQueryResponseBodyResult) SetViewEntityFieldVOList(v []*MasterDataQueryResponseBodyResultViewEntityFieldVOList) *MasterDataQueryResponseBodyResult {
	s.ViewEntityFieldVOList = v
	return s
}

func (s *MasterDataQueryResponseBodyResult) SetRelationId(v string) *MasterDataQueryResponseBodyResult {
	s.RelationId = &v
	return s
}

type MasterDataQueryResponseBodyResultViewEntityFieldVOList struct {
	// 字段code
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	// 字段值
	FieldDataVO *MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO `json:"fieldDataVO,omitempty" xml:"fieldDataVO,omitempty" type:"Struct"`
	// 字段名称
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// 字段类型
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
}

func (s MasterDataQueryResponseBodyResultViewEntityFieldVOList) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryResponseBodyResultViewEntityFieldVOList) GoString() string {
	return s.String()
}

func (s *MasterDataQueryResponseBodyResultViewEntityFieldVOList) SetFieldCode(v string) *MasterDataQueryResponseBodyResultViewEntityFieldVOList {
	s.FieldCode = &v
	return s
}

func (s *MasterDataQueryResponseBodyResultViewEntityFieldVOList) SetFieldDataVO(v *MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO) *MasterDataQueryResponseBodyResultViewEntityFieldVOList {
	s.FieldDataVO = v
	return s
}

func (s *MasterDataQueryResponseBodyResultViewEntityFieldVOList) SetFieldName(v string) *MasterDataQueryResponseBodyResultViewEntityFieldVOList {
	s.FieldName = &v
	return s
}

func (s *MasterDataQueryResponseBodyResultViewEntityFieldVOList) SetFieldType(v string) *MasterDataQueryResponseBodyResultViewEntityFieldVOList {
	s.FieldType = &v
	return s
}

type MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO struct {
	// 字段值的key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 字段值的文本
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO) GoString() string {
	return s.String()
}

func (s *MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO) SetKey(v string) *MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO {
	s.Key = &v
	return s
}

func (s *MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO) SetValue(v string) *MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO {
	s.Value = &v
	return s
}

type MasterDataQueryResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MasterDataQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MasterDataQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryResponse) GoString() string {
	return s.String()
}

func (s *MasterDataQueryResponse) SetHeaders(v map[string]*string) *MasterDataQueryResponse {
	s.Headers = v
	return s
}

func (s *MasterDataQueryResponse) SetBody(v *MasterDataQueryResponseBody) *MasterDataQueryResponse {
	s.Body = v
	return s
}

type AddHrmPreentryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddHrmPreentryHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryHeaders) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryHeaders) SetCommonHeaders(v map[string]*string) *AddHrmPreentryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddHrmPreentryHeaders) SetXAcsDingtalkAccessToken(v string) *AddHrmPreentryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddHrmPreentryRequest struct {
	PreEntryTime *int64                         `json:"preEntryTime,omitempty" xml:"preEntryTime,omitempty"`
	Name         *string                        `json:"name,omitempty" xml:"name,omitempty"`
	Mobile       *string                        `json:"mobile,omitempty" xml:"mobile,omitempty"`
	AgentId      *int64                         `json:"agentId,omitempty" xml:"agentId,omitempty"`
	Groups       []*AddHrmPreentryRequestGroups `json:"groups,omitempty" xml:"groups,omitempty" type:"Repeated"`
}

func (s AddHrmPreentryRequest) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryRequest) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryRequest) SetPreEntryTime(v int64) *AddHrmPreentryRequest {
	s.PreEntryTime = &v
	return s
}

func (s *AddHrmPreentryRequest) SetName(v string) *AddHrmPreentryRequest {
	s.Name = &v
	return s
}

func (s *AddHrmPreentryRequest) SetMobile(v string) *AddHrmPreentryRequest {
	s.Mobile = &v
	return s
}

func (s *AddHrmPreentryRequest) SetAgentId(v int64) *AddHrmPreentryRequest {
	s.AgentId = &v
	return s
}

func (s *AddHrmPreentryRequest) SetGroups(v []*AddHrmPreentryRequestGroups) *AddHrmPreentryRequest {
	s.Groups = v
	return s
}

type AddHrmPreentryRequestGroups struct {
	GroupId  *string                                `json:"groupId,omitempty" xml:"groupId,omitempty"`
	Sections []*AddHrmPreentryRequestGroupsSections `json:"sections,omitempty" xml:"sections,omitempty" type:"Repeated"`
}

func (s AddHrmPreentryRequestGroups) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryRequestGroups) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryRequestGroups) SetGroupId(v string) *AddHrmPreentryRequestGroups {
	s.GroupId = &v
	return s
}

func (s *AddHrmPreentryRequestGroups) SetSections(v []*AddHrmPreentryRequestGroupsSections) *AddHrmPreentryRequestGroups {
	s.Sections = v
	return s
}

type AddHrmPreentryRequestGroupsSections struct {
	OldIndex       *int32                                               `json:"oldIndex,omitempty" xml:"oldIndex,omitempty"`
	EmpFieldVOList []*AddHrmPreentryRequestGroupsSectionsEmpFieldVOList `json:"empFieldVOList,omitempty" xml:"empFieldVOList,omitempty" type:"Repeated"`
}

func (s AddHrmPreentryRequestGroupsSections) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryRequestGroupsSections) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryRequestGroupsSections) SetOldIndex(v int32) *AddHrmPreentryRequestGroupsSections {
	s.OldIndex = &v
	return s
}

func (s *AddHrmPreentryRequestGroupsSections) SetEmpFieldVOList(v []*AddHrmPreentryRequestGroupsSectionsEmpFieldVOList) *AddHrmPreentryRequestGroupsSections {
	s.EmpFieldVOList = v
	return s
}

type AddHrmPreentryRequestGroupsSectionsEmpFieldVOList struct {
	Value     *string `json:"value,omitempty" xml:"value,omitempty"`
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
}

func (s AddHrmPreentryRequestGroupsSectionsEmpFieldVOList) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryRequestGroupsSectionsEmpFieldVOList) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryRequestGroupsSectionsEmpFieldVOList) SetValue(v string) *AddHrmPreentryRequestGroupsSectionsEmpFieldVOList {
	s.Value = &v
	return s
}

func (s *AddHrmPreentryRequestGroupsSectionsEmpFieldVOList) SetFieldCode(v string) *AddHrmPreentryRequestGroupsSectionsEmpFieldVOList {
	s.FieldCode = &v
	return s
}

type AddHrmPreentryResponseBody struct {
	// result
	TmpUserId *string `json:"tmpUserId,omitempty" xml:"tmpUserId,omitempty"`
}

func (s AddHrmPreentryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryResponseBody) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryResponseBody) SetTmpUserId(v string) *AddHrmPreentryResponseBody {
	s.TmpUserId = &v
	return s
}

type AddHrmPreentryResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddHrmPreentryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddHrmPreentryResponse) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryResponse) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryResponse) SetHeaders(v map[string]*string) *AddHrmPreentryResponse {
	s.Headers = v
	return s
}

func (s *AddHrmPreentryResponse) SetBody(v *AddHrmPreentryResponseBody) *AddHrmPreentryResponse {
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

func (client *Client) ECertQuery(request *ECertQueryRequest) (_result *ECertQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ECertQueryHeaders{}
	_result = &ECertQueryResponse{}
	_body, _err := client.ECertQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ECertQueryWithOptions(request *ECertQueryRequest, headers *ECertQueryHeaders, runtime *util.RuntimeOptions) (_result *ECertQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
	_result = &ECertQueryResponse{}
	_body, _err := client.DoROARequest(tea.String("ECertQuery"), tea.String("hrm_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/hrm/eCerts"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryJobRanks(request *QueryJobRanksRequest) (_result *QueryJobRanksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryJobRanksHeaders{}
	_result = &QueryJobRanksResponse{}
	_body, _err := client.QueryJobRanksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryJobRanksWithOptions(request *QueryJobRanksRequest, headers *QueryJobRanksHeaders, runtime *util.RuntimeOptions) (_result *QueryJobRanksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RankCategoryId)) {
		query["rankCategoryId"] = request.RankCategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.RankCode)) {
		query["rankCode"] = request.RankCode
	}

	if !tea.BoolValue(util.IsUnset(request.RankName)) {
		query["rankName"] = request.RankName
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
	_result = &QueryJobRanksResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryJobRanks"), tea.String("hrm_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/hrm/jobRanks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryJobs(request *QueryJobsRequest) (_result *QueryJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryJobsHeaders{}
	_result = &QueryJobsResponse{}
	_body, _err := client.QueryJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryJobsWithOptions(request *QueryJobsRequest, headers *QueryJobsHeaders, runtime *util.RuntimeOptions) (_result *QueryJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobName)) {
		query["jobName"] = request.JobName
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
	_result = &QueryJobsResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryJobs"), tea.String("hrm_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/hrm/jobs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCustomEntryProcesses(request *QueryCustomEntryProcessesRequest) (_result *QueryCustomEntryProcessesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCustomEntryProcessesHeaders{}
	_result = &QueryCustomEntryProcessesResponse{}
	_body, _err := client.QueryCustomEntryProcessesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCustomEntryProcessesWithOptions(request *QueryCustomEntryProcessesRequest, headers *QueryCustomEntryProcessesHeaders, runtime *util.RuntimeOptions) (_result *QueryCustomEntryProcessesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperateUserId)) {
		query["operateUserId"] = request.OperateUserId
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
	_result = &QueryCustomEntryProcessesResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryCustomEntryProcesses"), tea.String("hrm_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/hrm/customEntryProcesses"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPositions(request *QueryPositionsRequest) (_result *QueryPositionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryPositionsHeaders{}
	_result = &QueryPositionsResponse{}
	_body, _err := client.QueryPositionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPositionsWithOptions(request *QueryPositionsRequest, headers *QueryPositionsHeaders, runtime *util.RuntimeOptions) (_result *QueryPositionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PositionName)) {
		body["positionName"] = request.PositionName
	}

	if !tea.BoolValue(util.IsUnset(request.InCategoryIds)) {
		body["inCategoryIds"] = request.InCategoryIds
	}

	if !tea.BoolValue(util.IsUnset(request.InPositionIds)) {
		body["inPositionIds"] = request.InPositionIds
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &QueryPositionsResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryPositions"), tea.String("hrm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/hrm/positions/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MasterDataQuery(request *MasterDataQueryRequest) (_result *MasterDataQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MasterDataQueryHeaders{}
	_result = &MasterDataQueryResponse{}
	_body, _err := client.MasterDataQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MasterDataQueryWithOptions(request *MasterDataQueryRequest, headers *MasterDataQueryHeaders, runtime *util.RuntimeOptions) (_result *MasterDataQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ScopeCode)) {
		body["scopeCode"] = request.ScopeCode
	}

	if !tea.BoolValue(util.IsUnset(request.ViewEntityCode)) {
		body["viewEntityCode"] = request.ViewEntityCode
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["tenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.BizUK)) {
		body["bizUK"] = request.BizUK
	}

	if !tea.BoolValue(util.IsUnset(request.RelationIds)) {
		body["relationIds"] = request.RelationIds
	}

	if !tea.BoolValue(util.IsUnset(request.OptUserId)) {
		body["optUserId"] = request.OptUserId
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.QueryParams)) {
		body["queryParams"] = request.QueryParams
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
	_result = &MasterDataQueryResponse{}
	_body, _err := client.DoROARequest(tea.String("MasterDataQuery"), tea.String("hrm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/hrm/masters/datas/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddHrmPreentry(request *AddHrmPreentryRequest) (_result *AddHrmPreentryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddHrmPreentryHeaders{}
	_result = &AddHrmPreentryResponse{}
	_body, _err := client.AddHrmPreentryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddHrmPreentryWithOptions(request *AddHrmPreentryRequest, headers *AddHrmPreentryHeaders, runtime *util.RuntimeOptions) (_result *AddHrmPreentryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PreEntryTime)) {
		body["preEntryTime"] = request.PreEntryTime
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		body["mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		body["agentId"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.Groups)) {
		body["groups"] = request.Groups
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
	_result = &AddHrmPreentryResponse{}
	_body, _err := client.DoROARequest(tea.String("AddHrmPreentry"), tea.String("hrm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/hrm/preentries"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
