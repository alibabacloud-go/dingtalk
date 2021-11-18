// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package edu_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type QueryStatisticsDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryStatisticsDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryStatisticsDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryStatisticsDataHeaders) SetCommonHeaders(v map[string]*string) *QueryStatisticsDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryStatisticsDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryStatisticsDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryStatisticsDataRequest struct {
	// startTime
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// endTime
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// opUserId
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 课程节次列表
	SectionIndexList []*int64 `json:"sectionIndexList,omitempty" xml:"sectionIndexList,omitempty" type:"Repeated"`
	// 老师UserIds
	TeacherUserIds []*string `json:"teacherUserIds,omitempty" xml:"teacherUserIds,omitempty" type:"Repeated"`
}

func (s QueryStatisticsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryStatisticsDataRequest) GoString() string {
	return s.String()
}

func (s *QueryStatisticsDataRequest) SetStartTime(v int64) *QueryStatisticsDataRequest {
	s.StartTime = &v
	return s
}

func (s *QueryStatisticsDataRequest) SetEndTime(v int64) *QueryStatisticsDataRequest {
	s.EndTime = &v
	return s
}

func (s *QueryStatisticsDataRequest) SetOpUserId(v string) *QueryStatisticsDataRequest {
	s.OpUserId = &v
	return s
}

func (s *QueryStatisticsDataRequest) SetSectionIndexList(v []*int64) *QueryStatisticsDataRequest {
	s.SectionIndexList = v
	return s
}

func (s *QueryStatisticsDataRequest) SetTeacherUserIds(v []*string) *QueryStatisticsDataRequest {
	s.TeacherUserIds = v
	return s
}

type QueryStatisticsDataResponseBody struct {
	// result
	Result []*QueryStatisticsDataResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryStatisticsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryStatisticsDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryStatisticsDataResponseBody) SetResult(v []*QueryStatisticsDataResponseBodyResult) *QueryStatisticsDataResponseBody {
	s.Result = v
	return s
}

type QueryStatisticsDataResponseBodyResult struct {
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户名称
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// 班级id
	ClassId *int64 `json:"classId,omitempty" xml:"classId,omitempty"`
	// 学科名称
	SubjectName *int64 `json:"subjectName,omitempty" xml:"subjectName,omitempty"`
	// 学科code
	SubjectCode *string `json:"subjectCode,omitempty" xml:"subjectCode,omitempty"`
	// 总学时
	CourseHours *float32 `json:"courseHours,omitempty" xml:"courseHours,omitempty"`
	// 总课程数
	CourseCount *int64 `json:"courseCount,omitempty" xml:"courseCount,omitempty"`
}

func (s QueryStatisticsDataResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryStatisticsDataResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryStatisticsDataResponseBodyResult) SetUserId(v string) *QueryStatisticsDataResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *QueryStatisticsDataResponseBodyResult) SetUserName(v string) *QueryStatisticsDataResponseBodyResult {
	s.UserName = &v
	return s
}

func (s *QueryStatisticsDataResponseBodyResult) SetClassId(v int64) *QueryStatisticsDataResponseBodyResult {
	s.ClassId = &v
	return s
}

func (s *QueryStatisticsDataResponseBodyResult) SetSubjectName(v int64) *QueryStatisticsDataResponseBodyResult {
	s.SubjectName = &v
	return s
}

func (s *QueryStatisticsDataResponseBodyResult) SetSubjectCode(v string) *QueryStatisticsDataResponseBodyResult {
	s.SubjectCode = &v
	return s
}

func (s *QueryStatisticsDataResponseBodyResult) SetCourseHours(v float32) *QueryStatisticsDataResponseBodyResult {
	s.CourseHours = &v
	return s
}

func (s *QueryStatisticsDataResponseBodyResult) SetCourseCount(v int64) *QueryStatisticsDataResponseBodyResult {
	s.CourseCount = &v
	return s
}

type QueryStatisticsDataResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryStatisticsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryStatisticsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryStatisticsDataResponse) GoString() string {
	return s.String()
}

func (s *QueryStatisticsDataResponse) SetHeaders(v map[string]*string) *QueryStatisticsDataResponse {
	s.Headers = v
	return s
}

func (s *QueryStatisticsDataResponse) SetBody(v *QueryStatisticsDataResponseBody) *QueryStatisticsDataResponse {
	s.Body = v
	return s
}

type QueryClassScheduleConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryClassScheduleConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleConfigHeaders) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleConfigHeaders) SetCommonHeaders(v map[string]*string) *QueryClassScheduleConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryClassScheduleConfigHeaders) SetXAcsDingtalkAccessToken(v string) *QueryClassScheduleConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryClassScheduleConfigRequest struct {
	// 课程id列表
	ClassIds []*int64 `json:"classIds,omitempty" xml:"classIds,omitempty" type:"Repeated"`
	// 操作者的UserID
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s QueryClassScheduleConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleConfigRequest) SetClassIds(v []*int64) *QueryClassScheduleConfigRequest {
	s.ClassIds = v
	return s
}

func (s *QueryClassScheduleConfigRequest) SetOpUserId(v string) *QueryClassScheduleConfigRequest {
	s.OpUserId = &v
	return s
}

type QueryClassScheduleConfigShrinkRequest struct {
	// 课程id列表
	ClassIdsShrink *string `json:"classIds,omitempty" xml:"classIds,omitempty"`
	// 操作者的UserID
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s QueryClassScheduleConfigShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleConfigShrinkRequest) SetClassIdsShrink(v string) *QueryClassScheduleConfigShrinkRequest {
	s.ClassIdsShrink = &v
	return s
}

func (s *QueryClassScheduleConfigShrinkRequest) SetOpUserId(v string) *QueryClassScheduleConfigShrinkRequest {
	s.OpUserId = &v
	return s
}

type QueryClassScheduleConfigResponseBody struct {
	// 查询结果
	Result []*QueryClassScheduleConfigResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryClassScheduleConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleConfigResponseBody) SetResult(v []*QueryClassScheduleConfigResponseBodyResult) *QueryClassScheduleConfigResponseBody {
	s.Result = v
	return s
}

type QueryClassScheduleConfigResponseBodyResult struct {
	// 班级的Id.
	ClassId *int64 `json:"classId,omitempty" xml:"classId,omitempty"`
	// 开始时间
	Start *QueryClassScheduleConfigResponseBodyResultStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
	End   *QueryClassScheduleConfigResponseBodyResultEnd   `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
	// 节次模型。
	SectionModels []*QueryClassScheduleConfigResponseBodyResultSectionModels `json:"sectionModels,omitempty" xml:"sectionModels,omitempty" type:"Repeated"`
}

func (s QueryClassScheduleConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleConfigResponseBodyResult) SetClassId(v int64) *QueryClassScheduleConfigResponseBodyResult {
	s.ClassId = &v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResult) SetStart(v *QueryClassScheduleConfigResponseBodyResultStart) *QueryClassScheduleConfigResponseBodyResult {
	s.Start = v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResult) SetEnd(v *QueryClassScheduleConfigResponseBodyResultEnd) *QueryClassScheduleConfigResponseBodyResult {
	s.End = v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResult) SetSectionModels(v []*QueryClassScheduleConfigResponseBodyResultSectionModels) *QueryClassScheduleConfigResponseBodyResult {
	s.SectionModels = v
	return s
}

type QueryClassScheduleConfigResponseBodyResultStart struct {
	// 年份
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
	// 月份
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// 一个月中的第几天
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
}

func (s QueryClassScheduleConfigResponseBodyResultStart) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleConfigResponseBodyResultStart) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleConfigResponseBodyResultStart) SetYear(v int32) *QueryClassScheduleConfigResponseBodyResultStart {
	s.Year = &v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResultStart) SetMonth(v int32) *QueryClassScheduleConfigResponseBodyResultStart {
	s.Month = &v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResultStart) SetDayOfMonth(v int32) *QueryClassScheduleConfigResponseBodyResultStart {
	s.DayOfMonth = &v
	return s
}

type QueryClassScheduleConfigResponseBodyResultEnd struct {
	// 年份
	Year *int64 `json:"year,omitempty" xml:"year,omitempty"`
	// 月份
	Month *int64 `json:"month,omitempty" xml:"month,omitempty"`
	// 一个月中第几天
	DayOfMonth *int64 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
}

func (s QueryClassScheduleConfigResponseBodyResultEnd) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleConfigResponseBodyResultEnd) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleConfigResponseBodyResultEnd) SetYear(v int64) *QueryClassScheduleConfigResponseBodyResultEnd {
	s.Year = &v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResultEnd) SetMonth(v int64) *QueryClassScheduleConfigResponseBodyResultEnd {
	s.Month = &v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResultEnd) SetDayOfMonth(v int64) *QueryClassScheduleConfigResponseBodyResultEnd {
	s.DayOfMonth = &v
	return s
}

type QueryClassScheduleConfigResponseBodyResultSectionModels struct {
	// 节次类型：COURSE/REST
	SectionType *string `json:"sectionType,omitempty" xml:"sectionType,omitempty"`
	// 开始时间（时分）
	Start *QueryClassScheduleConfigResponseBodyResultSectionModelsStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
	// 节次设置
	SectionIndex *int32 `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty"`
	// 结束时间
	End *QueryClassScheduleConfigResponseBodyResultSectionModelsEnd `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
	// 节次名称
	SectionName *string `json:"sectionName,omitempty" xml:"sectionName,omitempty"`
}

func (s QueryClassScheduleConfigResponseBodyResultSectionModels) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleConfigResponseBodyResultSectionModels) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleConfigResponseBodyResultSectionModels) SetSectionType(v string) *QueryClassScheduleConfigResponseBodyResultSectionModels {
	s.SectionType = &v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResultSectionModels) SetStart(v *QueryClassScheduleConfigResponseBodyResultSectionModelsStart) *QueryClassScheduleConfigResponseBodyResultSectionModels {
	s.Start = v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResultSectionModels) SetSectionIndex(v int32) *QueryClassScheduleConfigResponseBodyResultSectionModels {
	s.SectionIndex = &v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResultSectionModels) SetEnd(v *QueryClassScheduleConfigResponseBodyResultSectionModelsEnd) *QueryClassScheduleConfigResponseBodyResultSectionModels {
	s.End = v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResultSectionModels) SetSectionName(v string) *QueryClassScheduleConfigResponseBodyResultSectionModels {
	s.SectionName = &v
	return s
}

type QueryClassScheduleConfigResponseBodyResultSectionModelsStart struct {
	// 分钟
	Min *int32 `json:"min,omitempty" xml:"min,omitempty"`
	// 小时
	Hour *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
}

func (s QueryClassScheduleConfigResponseBodyResultSectionModelsStart) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleConfigResponseBodyResultSectionModelsStart) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleConfigResponseBodyResultSectionModelsStart) SetMin(v int32) *QueryClassScheduleConfigResponseBodyResultSectionModelsStart {
	s.Min = &v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResultSectionModelsStart) SetHour(v int32) *QueryClassScheduleConfigResponseBodyResultSectionModelsStart {
	s.Hour = &v
	return s
}

type QueryClassScheduleConfigResponseBodyResultSectionModelsEnd struct {
	// 分钟
	Min *int32 `json:"min,omitempty" xml:"min,omitempty"`
	// 小时
	Hour *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
}

func (s QueryClassScheduleConfigResponseBodyResultSectionModelsEnd) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleConfigResponseBodyResultSectionModelsEnd) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleConfigResponseBodyResultSectionModelsEnd) SetMin(v int32) *QueryClassScheduleConfigResponseBodyResultSectionModelsEnd {
	s.Min = &v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResultSectionModelsEnd) SetHour(v int32) *QueryClassScheduleConfigResponseBodyResultSectionModelsEnd {
	s.Hour = &v
	return s
}

type QueryClassScheduleConfigResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryClassScheduleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryClassScheduleConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleConfigResponse) SetHeaders(v map[string]*string) *QueryClassScheduleConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryClassScheduleConfigResponse) SetBody(v *QueryClassScheduleConfigResponseBody) *QueryClassScheduleConfigResponse {
	s.Body = v
	return s
}

type GetDefaultChildHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDefaultChildHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultChildHeaders) GoString() string {
	return s.String()
}

func (s *GetDefaultChildHeaders) SetCommonHeaders(v map[string]*string) *GetDefaultChildHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDefaultChildHeaders) SetXAcsDingtalkAccessToken(v string) *GetDefaultChildHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDefaultChildResponseBody struct {
	Name         *string                                    `json:"name,omitempty" xml:"name,omitempty"`
	Nick         *string                                    `json:"nick,omitempty" xml:"nick,omitempty"`
	Avatar       *string                                    `json:"avatar,omitempty" xml:"avatar,omitempty"`
	UnionId      *string                                    `json:"unionId,omitempty" xml:"unionId,omitempty"`
	Period       *string                                    `json:"period,omitempty" xml:"period,omitempty"`
	Grade        *int32                                     `json:"grade,omitempty" xml:"grade,omitempty"`
	BindStudents []*GetDefaultChildResponseBodyBindStudents `json:"bindStudents,omitempty" xml:"bindStudents,omitempty" type:"Repeated"`
}

func (s GetDefaultChildResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultChildResponseBody) GoString() string {
	return s.String()
}

func (s *GetDefaultChildResponseBody) SetName(v string) *GetDefaultChildResponseBody {
	s.Name = &v
	return s
}

func (s *GetDefaultChildResponseBody) SetNick(v string) *GetDefaultChildResponseBody {
	s.Nick = &v
	return s
}

func (s *GetDefaultChildResponseBody) SetAvatar(v string) *GetDefaultChildResponseBody {
	s.Avatar = &v
	return s
}

func (s *GetDefaultChildResponseBody) SetUnionId(v string) *GetDefaultChildResponseBody {
	s.UnionId = &v
	return s
}

func (s *GetDefaultChildResponseBody) SetPeriod(v string) *GetDefaultChildResponseBody {
	s.Period = &v
	return s
}

func (s *GetDefaultChildResponseBody) SetGrade(v int32) *GetDefaultChildResponseBody {
	s.Grade = &v
	return s
}

func (s *GetDefaultChildResponseBody) SetBindStudents(v []*GetDefaultChildResponseBodyBindStudents) *GetDefaultChildResponseBody {
	s.BindStudents = v
	return s
}

type GetDefaultChildResponseBodyBindStudents struct {
	CorpId  *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	StaffId *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
	ClassId *int64  `json:"classId,omitempty" xml:"classId,omitempty"`
	Period  *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s GetDefaultChildResponseBodyBindStudents) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultChildResponseBodyBindStudents) GoString() string {
	return s.String()
}

func (s *GetDefaultChildResponseBodyBindStudents) SetCorpId(v string) *GetDefaultChildResponseBodyBindStudents {
	s.CorpId = &v
	return s
}

func (s *GetDefaultChildResponseBodyBindStudents) SetStaffId(v string) *GetDefaultChildResponseBodyBindStudents {
	s.StaffId = &v
	return s
}

func (s *GetDefaultChildResponseBodyBindStudents) SetClassId(v int64) *GetDefaultChildResponseBodyBindStudents {
	s.ClassId = &v
	return s
}

func (s *GetDefaultChildResponseBodyBindStudents) SetPeriod(v string) *GetDefaultChildResponseBodyBindStudents {
	s.Period = &v
	return s
}

type GetDefaultChildResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDefaultChildResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDefaultChildResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultChildResponse) GoString() string {
	return s.String()
}

func (s *GetDefaultChildResponse) SetHeaders(v map[string]*string) *GetDefaultChildResponse {
	s.Headers = v
	return s
}

func (s *GetDefaultChildResponse) SetBody(v *GetDefaultChildResponseBody) *GetDefaultChildResponse {
	s.Body = v
	return s
}

type GetOpenCoursesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOpenCoursesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOpenCoursesHeaders) GoString() string {
	return s.String()
}

func (s *GetOpenCoursesHeaders) SetCommonHeaders(v map[string]*string) *GetOpenCoursesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOpenCoursesHeaders) SetXAcsDingtalkAccessToken(v string) *GetOpenCoursesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOpenCoursesRequest struct {
	// 分页起始, 起始值为0
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页大小
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetOpenCoursesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOpenCoursesRequest) GoString() string {
	return s.String()
}

func (s *GetOpenCoursesRequest) SetPageNumber(v int64) *GetOpenCoursesRequest {
	s.PageNumber = &v
	return s
}

func (s *GetOpenCoursesRequest) SetPageSize(v int64) *GetOpenCoursesRequest {
	s.PageSize = &v
	return s
}

type GetOpenCoursesResponseBody struct {
	// 总记录数
	TotalCount *int64                                  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	CourseList []*GetOpenCoursesResponseBodyCourseList `json:"courseList,omitempty" xml:"courseList,omitempty" type:"Repeated"`
}

func (s GetOpenCoursesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOpenCoursesResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpenCoursesResponseBody) SetTotalCount(v int64) *GetOpenCoursesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetOpenCoursesResponseBody) SetCourseList(v []*GetOpenCoursesResponseBodyCourseList) *GetOpenCoursesResponseBody {
	s.CourseList = v
	return s
}

type GetOpenCoursesResponseBodyCourseList struct {
	// 课程id
	CourseId *string `json:"courseId,omitempty" xml:"courseId,omitempty"`
	// 课程标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 课程类型: 0-直播 2-视频内容
	FeedType *int64 `json:"feedType,omitempty" xml:"feedType,omitempty"`
	// 老师名称
	TeacherName *string `json:"teacherName,omitempty" xml:"teacherName,omitempty"`
	// 封面图片地址
	CoverUrl *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	// 课程开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 课程观看地址
	JumpUrl *string `json:"jumpUrl,omitempty" xml:"jumpUrl,omitempty"`
	// 老师的userId
	TeacherId *string `json:"teacherId,omitempty" xml:"teacherId,omitempty"`
}

func (s GetOpenCoursesResponseBodyCourseList) String() string {
	return tea.Prettify(s)
}

func (s GetOpenCoursesResponseBodyCourseList) GoString() string {
	return s.String()
}

func (s *GetOpenCoursesResponseBodyCourseList) SetCourseId(v string) *GetOpenCoursesResponseBodyCourseList {
	s.CourseId = &v
	return s
}

func (s *GetOpenCoursesResponseBodyCourseList) SetTitle(v string) *GetOpenCoursesResponseBodyCourseList {
	s.Title = &v
	return s
}

func (s *GetOpenCoursesResponseBodyCourseList) SetFeedType(v int64) *GetOpenCoursesResponseBodyCourseList {
	s.FeedType = &v
	return s
}

func (s *GetOpenCoursesResponseBodyCourseList) SetTeacherName(v string) *GetOpenCoursesResponseBodyCourseList {
	s.TeacherName = &v
	return s
}

func (s *GetOpenCoursesResponseBodyCourseList) SetCoverUrl(v string) *GetOpenCoursesResponseBodyCourseList {
	s.CoverUrl = &v
	return s
}

func (s *GetOpenCoursesResponseBodyCourseList) SetStartTime(v int64) *GetOpenCoursesResponseBodyCourseList {
	s.StartTime = &v
	return s
}

func (s *GetOpenCoursesResponseBodyCourseList) SetJumpUrl(v string) *GetOpenCoursesResponseBodyCourseList {
	s.JumpUrl = &v
	return s
}

func (s *GetOpenCoursesResponseBodyCourseList) SetTeacherId(v string) *GetOpenCoursesResponseBodyCourseList {
	s.TeacherId = &v
	return s
}

type GetOpenCoursesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOpenCoursesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOpenCoursesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOpenCoursesResponse) GoString() string {
	return s.String()
}

func (s *GetOpenCoursesResponse) SetHeaders(v map[string]*string) *GetOpenCoursesResponse {
	s.Headers = v
	return s
}

func (s *GetOpenCoursesResponse) SetBody(v *GetOpenCoursesResponseBody) *GetOpenCoursesResponse {
	s.Body = v
	return s
}

type BatchCreateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchCreateHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateHeaders) GoString() string {
	return s.String()
}

func (s *BatchCreateHeaders) SetCommonHeaders(v map[string]*string) *BatchCreateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchCreateHeaders) SetXAcsDingtalkAccessToken(v string) *BatchCreateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchCreateRequest struct {
	// 卡片业务类型，打卡写死：industry_center
	CardBizCode *string `json:"cardBizCode,omitempty" xml:"cardBizCode,omitempty"`
	// 卡片详细数据
	Data *BatchCreateRequestData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 卡片幂等唯一键
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// isv业务类型
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// 老师用户id
	Userid *string `json:"userid,omitempty" xml:"userid,omitempty"`
	// 老师corpId
	DingCorpId *string `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	// 小程序版本号
	JsVersion *int32 `json:"jsVersion,omitempty" xml:"jsVersion,omitempty"`
}

func (s BatchCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateRequest) SetCardBizCode(v string) *BatchCreateRequest {
	s.CardBizCode = &v
	return s
}

func (s *BatchCreateRequest) SetData(v *BatchCreateRequestData) *BatchCreateRequest {
	s.Data = v
	return s
}

func (s *BatchCreateRequest) SetIdentifier(v string) *BatchCreateRequest {
	s.Identifier = &v
	return s
}

func (s *BatchCreateRequest) SetSourceType(v string) *BatchCreateRequest {
	s.SourceType = &v
	return s
}

func (s *BatchCreateRequest) SetUserid(v string) *BatchCreateRequest {
	s.Userid = &v
	return s
}

func (s *BatchCreateRequest) SetDingCorpId(v string) *BatchCreateRequest {
	s.DingCorpId = &v
	return s
}

func (s *BatchCreateRequest) SetJsVersion(v int32) *BatchCreateRequest {
	s.JsVersion = &v
	return s
}

type BatchCreateRequestData struct {
	// 是否可以补卡
	CanReissueCard *bool `json:"canReissueCard,omitempty" xml:"canReissueCard,omitempty"`
	// 打卡周期,单位为天
	CardCycle *int32 `json:"cardCycle,omitempty" xml:"cardCycle,omitempty"`
	// 打卡的频次设置："cardFrequency":[             1,//周天             2,//周一             3,//周二             4,//周三             5,//周四             6,//周五             7//周六         ]
	CardFrequency         []*int32                                       `json:"cardFrequency,omitempty" xml:"cardFrequency,omitempty" type:"Repeated"`
	CardRuleItemParamList []*BatchCreateRequestDataCardRuleItemParamList `json:"cardRuleItemParamList,omitempty" xml:"cardRuleItemParamList,omitempty" type:"Repeated"`
	// 班级列表
	ClassIds []*string `json:"classIds,omitempty" xml:"classIds,omitempty" type:"Repeated"`
	// 班级名称列表
	ClassNames []*string `json:"classNames,omitempty" xml:"classNames,omitempty" type:"Repeated"`
	// 打卡的内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 卡片生效时间
	EffectDate *int64 `json:"effectDate,omitempty" xml:"effectDate,omitempty"`
	// 上传相册，图片，录音，盯盘的信息
	Medias *string `json:"medias,omitempty" xml:"medias,omitempty"`
	// 计量开启
	NeedMetering             *string                                           `json:"needMetering,omitempty" xml:"needMetering,omitempty"`
	OrgClassStudentGroupList []*BatchCreateRequestDataOrgClassStudentGroupList `json:"orgClassStudentGroupList,omitempty" xml:"orgClassStudentGroupList,omitempty" type:"Repeated"`
	// 提醒时间（小时）
	RemindHour *int32 `json:"remindHour,omitempty" xml:"remindHour,omitempty"`
	// 提醒时间（分钟）
	RemindMinute *int32 `json:"remindMinute,omitempty" xml:"remindMinute,omitempty"`
	// 默认：student_guardian
	TargetRole *string `json:"targetRole,omitempty" xml:"targetRole,omitempty"`
	// 打卡模板id
	TemplateId *int64 `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// 卡片标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 计量单位
	UnitOfMeasurement *string `json:"unitOfMeasurement,omitempty" xml:"unitOfMeasurement,omitempty"`
}

func (s BatchCreateRequestData) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateRequestData) GoString() string {
	return s.String()
}

func (s *BatchCreateRequestData) SetCanReissueCard(v bool) *BatchCreateRequestData {
	s.CanReissueCard = &v
	return s
}

func (s *BatchCreateRequestData) SetCardCycle(v int32) *BatchCreateRequestData {
	s.CardCycle = &v
	return s
}

func (s *BatchCreateRequestData) SetCardFrequency(v []*int32) *BatchCreateRequestData {
	s.CardFrequency = v
	return s
}

func (s *BatchCreateRequestData) SetCardRuleItemParamList(v []*BatchCreateRequestDataCardRuleItemParamList) *BatchCreateRequestData {
	s.CardRuleItemParamList = v
	return s
}

func (s *BatchCreateRequestData) SetClassIds(v []*string) *BatchCreateRequestData {
	s.ClassIds = v
	return s
}

func (s *BatchCreateRequestData) SetClassNames(v []*string) *BatchCreateRequestData {
	s.ClassNames = v
	return s
}

func (s *BatchCreateRequestData) SetContent(v string) *BatchCreateRequestData {
	s.Content = &v
	return s
}

func (s *BatchCreateRequestData) SetEffectDate(v int64) *BatchCreateRequestData {
	s.EffectDate = &v
	return s
}

func (s *BatchCreateRequestData) SetMedias(v string) *BatchCreateRequestData {
	s.Medias = &v
	return s
}

func (s *BatchCreateRequestData) SetNeedMetering(v string) *BatchCreateRequestData {
	s.NeedMetering = &v
	return s
}

func (s *BatchCreateRequestData) SetOrgClassStudentGroupList(v []*BatchCreateRequestDataOrgClassStudentGroupList) *BatchCreateRequestData {
	s.OrgClassStudentGroupList = v
	return s
}

func (s *BatchCreateRequestData) SetRemindHour(v int32) *BatchCreateRequestData {
	s.RemindHour = &v
	return s
}

func (s *BatchCreateRequestData) SetRemindMinute(v int32) *BatchCreateRequestData {
	s.RemindMinute = &v
	return s
}

func (s *BatchCreateRequestData) SetTargetRole(v string) *BatchCreateRequestData {
	s.TargetRole = &v
	return s
}

func (s *BatchCreateRequestData) SetTemplateId(v int64) *BatchCreateRequestData {
	s.TemplateId = &v
	return s
}

func (s *BatchCreateRequestData) SetTitle(v string) *BatchCreateRequestData {
	s.Title = &v
	return s
}

func (s *BatchCreateRequestData) SetUnitOfMeasurement(v string) *BatchCreateRequestData {
	s.UnitOfMeasurement = &v
	return s
}

type BatchCreateRequestDataCardRuleItemParamList struct {
	// 卡片taskCode
	CardTaskCode *string `json:"cardTaskCode,omitempty" xml:"cardTaskCode,omitempty"`
	// 关联的外部Id
	RelationId *string `json:"relationId,omitempty" xml:"relationId,omitempty"`
	// 扩展属性，存放配音难度、每日配音视频的url等
	CardRuleAttr *string `json:"cardRuleAttr,omitempty" xml:"cardRuleAttr,omitempty"`
	// 每日配音数
	DailyDubbing *int32 `json:"dailyDubbing,omitempty" xml:"dailyDubbing,omitempty"`
	// 关联内容标题（会在打卡详页页展示）
	RelationTitle *string `json:"relationTitle,omitempty" xml:"relationTitle,omitempty"`
	// relationUrl（点击打卡内容后跳转的链接）（点击卡片内容后跳转的链接）
	RelationUrl *string `json:"relationUrl,omitempty" xml:"relationUrl,omitempty"`
}

func (s BatchCreateRequestDataCardRuleItemParamList) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateRequestDataCardRuleItemParamList) GoString() string {
	return s.String()
}

func (s *BatchCreateRequestDataCardRuleItemParamList) SetCardTaskCode(v string) *BatchCreateRequestDataCardRuleItemParamList {
	s.CardTaskCode = &v
	return s
}

func (s *BatchCreateRequestDataCardRuleItemParamList) SetRelationId(v string) *BatchCreateRequestDataCardRuleItemParamList {
	s.RelationId = &v
	return s
}

func (s *BatchCreateRequestDataCardRuleItemParamList) SetCardRuleAttr(v string) *BatchCreateRequestDataCardRuleItemParamList {
	s.CardRuleAttr = &v
	return s
}

func (s *BatchCreateRequestDataCardRuleItemParamList) SetDailyDubbing(v int32) *BatchCreateRequestDataCardRuleItemParamList {
	s.DailyDubbing = &v
	return s
}

func (s *BatchCreateRequestDataCardRuleItemParamList) SetRelationTitle(v string) *BatchCreateRequestDataCardRuleItemParamList {
	s.RelationTitle = &v
	return s
}

func (s *BatchCreateRequestDataCardRuleItemParamList) SetRelationUrl(v string) *BatchCreateRequestDataCardRuleItemParamList {
	s.RelationUrl = &v
	return s
}

type BatchCreateRequestDataOrgClassStudentGroupList struct {
	// 组织id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 班级列表
	ClassList []*BatchCreateRequestDataOrgClassStudentGroupListClassList `json:"classList,omitempty" xml:"classList,omitempty" type:"Repeated"`
}

func (s BatchCreateRequestDataOrgClassStudentGroupList) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateRequestDataOrgClassStudentGroupList) GoString() string {
	return s.String()
}

func (s *BatchCreateRequestDataOrgClassStudentGroupList) SetCorpId(v string) *BatchCreateRequestDataOrgClassStudentGroupList {
	s.CorpId = &v
	return s
}

func (s *BatchCreateRequestDataOrgClassStudentGroupList) SetClassList(v []*BatchCreateRequestDataOrgClassStudentGroupListClassList) *BatchCreateRequestDataOrgClassStudentGroupList {
	s.ClassList = v
	return s
}

type BatchCreateRequestDataOrgClassStudentGroupListClassList struct {
	// 班级id
	ClassId *int64 `json:"classId,omitempty" xml:"classId,omitempty"`
	// 班级名称
	ClassName *string `json:"className,omitempty" xml:"className,omitempty"`
	// 班级学生
	Students []*BatchCreateRequestDataOrgClassStudentGroupListClassListStudents `json:"students,omitempty" xml:"students,omitempty" type:"Repeated"`
}

func (s BatchCreateRequestDataOrgClassStudentGroupListClassList) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateRequestDataOrgClassStudentGroupListClassList) GoString() string {
	return s.String()
}

func (s *BatchCreateRequestDataOrgClassStudentGroupListClassList) SetClassId(v int64) *BatchCreateRequestDataOrgClassStudentGroupListClassList {
	s.ClassId = &v
	return s
}

func (s *BatchCreateRequestDataOrgClassStudentGroupListClassList) SetClassName(v string) *BatchCreateRequestDataOrgClassStudentGroupListClassList {
	s.ClassName = &v
	return s
}

func (s *BatchCreateRequestDataOrgClassStudentGroupListClassList) SetStudents(v []*BatchCreateRequestDataOrgClassStudentGroupListClassListStudents) *BatchCreateRequestDataOrgClassStudentGroupListClassList {
	s.Students = v
	return s
}

type BatchCreateRequestDataOrgClassStudentGroupListClassListStudents struct {
	// 学生名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 学生id
	StaffId *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
}

func (s BatchCreateRequestDataOrgClassStudentGroupListClassListStudents) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateRequestDataOrgClassStudentGroupListClassListStudents) GoString() string {
	return s.String()
}

func (s *BatchCreateRequestDataOrgClassStudentGroupListClassListStudents) SetName(v string) *BatchCreateRequestDataOrgClassStudentGroupListClassListStudents {
	s.Name = &v
	return s
}

func (s *BatchCreateRequestDataOrgClassStudentGroupListClassListStudents) SetStaffId(v string) *BatchCreateRequestDataOrgClassStudentGroupListClassListStudents {
	s.StaffId = &v
	return s
}

type BatchCreateResponseBody struct {
	// result
	Result *BatchCreateResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s BatchCreateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateResponseBody) SetResult(v *BatchCreateResponseBodyResult) *BatchCreateResponseBody {
	s.Result = v
	return s
}

type BatchCreateResponseBodyResult struct {
	CorpIdCardIdMap map[string]*string `json:"corpIdCardIdMap,omitempty" xml:"corpIdCardIdMap,omitempty"`
}

func (s BatchCreateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *BatchCreateResponseBodyResult) SetCorpIdCardIdMap(v map[string]*string) *BatchCreateResponseBodyResult {
	s.CorpIdCardIdMap = v
	return s
}

type BatchCreateResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchCreateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateResponse) SetHeaders(v map[string]*string) *BatchCreateResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateResponse) SetBody(v *BatchCreateResponseBody) *BatchCreateResponse {
	s.Body = v
	return s
}

type InitCoursesOfClassHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InitCoursesOfClassHeaders) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassHeaders) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassHeaders) SetCommonHeaders(v map[string]*string) *InitCoursesOfClassHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InitCoursesOfClassHeaders) SetXAcsDingtalkAccessToken(v string) *InitCoursesOfClassHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InitCoursesOfClassRequest struct {
	// 课程设置。
	Courses []*InitCoursesOfClassRequestCourses `json:"courses,omitempty" xml:"courses,omitempty" type:"Repeated"`
	// 节次设置
	SectionConfig *InitCoursesOfClassRequestSectionConfig `json:"sectionConfig,omitempty" xml:"sectionConfig,omitempty" type:"Struct"`
	// 操作人的userid。
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s InitCoursesOfClassRequest) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassRequest) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassRequest) SetCourses(v []*InitCoursesOfClassRequestCourses) *InitCoursesOfClassRequest {
	s.Courses = v
	return s
}

func (s *InitCoursesOfClassRequest) SetSectionConfig(v *InitCoursesOfClassRequestSectionConfig) *InitCoursesOfClassRequest {
	s.SectionConfig = v
	return s
}

func (s *InitCoursesOfClassRequest) SetOpUserId(v string) *InitCoursesOfClassRequest {
	s.OpUserId = &v
	return s
}

type InitCoursesOfClassRequestCourses struct {
	// 老师的staffId。
	TeacherStaffIds []*string `json:"teacherStaffIds,omitempty" xml:"teacherStaffIds,omitempty" type:"Repeated"`
	// 课程名称。
	CourseName *string `json:"courseName,omitempty" xml:"courseName,omitempty"`
	// 上课时间。
	DateModel *InitCoursesOfClassRequestCoursesDateModel `json:"dateModel,omitempty" xml:"dateModel,omitempty" type:"Struct"`
	// 课程节次。
	SectionModel *InitCoursesOfClassRequestCoursesSectionModel `json:"sectionModel,omitempty" xml:"sectionModel,omitempty" type:"Struct"`
	// 创建者名称。
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// 上课地点
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
}

func (s InitCoursesOfClassRequestCourses) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassRequestCourses) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassRequestCourses) SetTeacherStaffIds(v []*string) *InitCoursesOfClassRequestCourses {
	s.TeacherStaffIds = v
	return s
}

func (s *InitCoursesOfClassRequestCourses) SetCourseName(v string) *InitCoursesOfClassRequestCourses {
	s.CourseName = &v
	return s
}

func (s *InitCoursesOfClassRequestCourses) SetDateModel(v *InitCoursesOfClassRequestCoursesDateModel) *InitCoursesOfClassRequestCourses {
	s.DateModel = v
	return s
}

func (s *InitCoursesOfClassRequestCourses) SetSectionModel(v *InitCoursesOfClassRequestCoursesSectionModel) *InitCoursesOfClassRequestCourses {
	s.SectionModel = v
	return s
}

func (s *InitCoursesOfClassRequestCourses) SetCreatorName(v string) *InitCoursesOfClassRequestCourses {
	s.CreatorName = &v
	return s
}

func (s *InitCoursesOfClassRequestCourses) SetLocation(v string) *InitCoursesOfClassRequestCourses {
	s.Location = &v
	return s
}

type InitCoursesOfClassRequestCoursesDateModel struct {
	// 月份。
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// 年份。
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
	// 每个月的第几天。
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
}

func (s InitCoursesOfClassRequestCoursesDateModel) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassRequestCoursesDateModel) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassRequestCoursesDateModel) SetMonth(v int32) *InitCoursesOfClassRequestCoursesDateModel {
	s.Month = &v
	return s
}

func (s *InitCoursesOfClassRequestCoursesDateModel) SetYear(v int32) *InitCoursesOfClassRequestCoursesDateModel {
	s.Year = &v
	return s
}

func (s *InitCoursesOfClassRequestCoursesDateModel) SetDayOfMonth(v int32) *InitCoursesOfClassRequestCoursesDateModel {
	s.DayOfMonth = &v
	return s
}

type InitCoursesOfClassRequestCoursesSectionModel struct {
	// 节次序列号。
	SectionIndex *int32 `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty"`
	// 节次名称。
	SectionName *string `json:"sectionName,omitempty" xml:"sectionName,omitempty"`
}

func (s InitCoursesOfClassRequestCoursesSectionModel) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassRequestCoursesSectionModel) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassRequestCoursesSectionModel) SetSectionIndex(v int32) *InitCoursesOfClassRequestCoursesSectionModel {
	s.SectionIndex = &v
	return s
}

func (s *InitCoursesOfClassRequestCoursesSectionModel) SetSectionName(v string) *InitCoursesOfClassRequestCoursesSectionModel {
	s.SectionName = &v
	return s
}

type InitCoursesOfClassRequestSectionConfig struct {
	// 节次模型
	SectionModels []*InitCoursesOfClassRequestSectionConfigSectionModels `json:"sectionModels,omitempty" xml:"sectionModels,omitempty" type:"Repeated"`
	// 课程表开始时间（精确到日）
	Start *InitCoursesOfClassRequestSectionConfigStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
	// 课程表结束开始时间（精确到日）
	End *InitCoursesOfClassRequestSectionConfigEnd `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
}

func (s InitCoursesOfClassRequestSectionConfig) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassRequestSectionConfig) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassRequestSectionConfig) SetSectionModels(v []*InitCoursesOfClassRequestSectionConfigSectionModels) *InitCoursesOfClassRequestSectionConfig {
	s.SectionModels = v
	return s
}

func (s *InitCoursesOfClassRequestSectionConfig) SetStart(v *InitCoursesOfClassRequestSectionConfigStart) *InitCoursesOfClassRequestSectionConfig {
	s.Start = v
	return s
}

func (s *InitCoursesOfClassRequestSectionConfig) SetEnd(v *InitCoursesOfClassRequestSectionConfigEnd) *InitCoursesOfClassRequestSectionConfig {
	s.End = v
	return s
}

type InitCoursesOfClassRequestSectionConfigSectionModels struct {
	// 节次类型枚举：COURSE/REST
	SectionType *string `json:"sectionType,omitempty" xml:"sectionType,omitempty"`
	// 开始时间
	Start *InitCoursesOfClassRequestSectionConfigSectionModelsStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
	// 第几节。
	SectionIndex *int32 `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty"`
	// 结束时间
	End *InitCoursesOfClassRequestSectionConfigSectionModelsEnd `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
}

func (s InitCoursesOfClassRequestSectionConfigSectionModels) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassRequestSectionConfigSectionModels) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassRequestSectionConfigSectionModels) SetSectionType(v string) *InitCoursesOfClassRequestSectionConfigSectionModels {
	s.SectionType = &v
	return s
}

func (s *InitCoursesOfClassRequestSectionConfigSectionModels) SetStart(v *InitCoursesOfClassRequestSectionConfigSectionModelsStart) *InitCoursesOfClassRequestSectionConfigSectionModels {
	s.Start = v
	return s
}

func (s *InitCoursesOfClassRequestSectionConfigSectionModels) SetSectionIndex(v int32) *InitCoursesOfClassRequestSectionConfigSectionModels {
	s.SectionIndex = &v
	return s
}

func (s *InitCoursesOfClassRequestSectionConfigSectionModels) SetEnd(v *InitCoursesOfClassRequestSectionConfigSectionModelsEnd) *InitCoursesOfClassRequestSectionConfigSectionModels {
	s.End = v
	return s
}

type InitCoursesOfClassRequestSectionConfigSectionModelsStart struct {
	// 分钟
	Min *int32 `json:"min,omitempty" xml:"min,omitempty"`
	// 小时
	Hour *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
}

func (s InitCoursesOfClassRequestSectionConfigSectionModelsStart) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassRequestSectionConfigSectionModelsStart) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassRequestSectionConfigSectionModelsStart) SetMin(v int32) *InitCoursesOfClassRequestSectionConfigSectionModelsStart {
	s.Min = &v
	return s
}

func (s *InitCoursesOfClassRequestSectionConfigSectionModelsStart) SetHour(v int32) *InitCoursesOfClassRequestSectionConfigSectionModelsStart {
	s.Hour = &v
	return s
}

type InitCoursesOfClassRequestSectionConfigSectionModelsEnd struct {
	// 分钟
	Min *int32 `json:"min,omitempty" xml:"min,omitempty"`
	// 小时
	Hour *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
}

func (s InitCoursesOfClassRequestSectionConfigSectionModelsEnd) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassRequestSectionConfigSectionModelsEnd) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassRequestSectionConfigSectionModelsEnd) SetMin(v int32) *InitCoursesOfClassRequestSectionConfigSectionModelsEnd {
	s.Min = &v
	return s
}

func (s *InitCoursesOfClassRequestSectionConfigSectionModelsEnd) SetHour(v int32) *InitCoursesOfClassRequestSectionConfigSectionModelsEnd {
	s.Hour = &v
	return s
}

type InitCoursesOfClassRequestSectionConfigStart struct {
	// 月份。
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// 年份。
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
	// 每个月的第几天。
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
}

func (s InitCoursesOfClassRequestSectionConfigStart) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassRequestSectionConfigStart) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassRequestSectionConfigStart) SetMonth(v int32) *InitCoursesOfClassRequestSectionConfigStart {
	s.Month = &v
	return s
}

func (s *InitCoursesOfClassRequestSectionConfigStart) SetYear(v int32) *InitCoursesOfClassRequestSectionConfigStart {
	s.Year = &v
	return s
}

func (s *InitCoursesOfClassRequestSectionConfigStart) SetDayOfMonth(v int32) *InitCoursesOfClassRequestSectionConfigStart {
	s.DayOfMonth = &v
	return s
}

type InitCoursesOfClassRequestSectionConfigEnd struct {
	// 月份。
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// 年份。
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
	// 每个月的第几天。
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
}

func (s InitCoursesOfClassRequestSectionConfigEnd) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassRequestSectionConfigEnd) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassRequestSectionConfigEnd) SetMonth(v int32) *InitCoursesOfClassRequestSectionConfigEnd {
	s.Month = &v
	return s
}

func (s *InitCoursesOfClassRequestSectionConfigEnd) SetYear(v int32) *InitCoursesOfClassRequestSectionConfigEnd {
	s.Year = &v
	return s
}

func (s *InitCoursesOfClassRequestSectionConfigEnd) SetDayOfMonth(v int32) *InitCoursesOfClassRequestSectionConfigEnd {
	s.DayOfMonth = &v
	return s
}

type InitCoursesOfClassResponseBody struct {
	// 初始化是否成功。
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s InitCoursesOfClassResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassResponseBody) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassResponseBody) SetResult(v bool) *InitCoursesOfClassResponseBody {
	s.Result = &v
	return s
}

type InitCoursesOfClassResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InitCoursesOfClassResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitCoursesOfClassResponse) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassResponse) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassResponse) SetHeaders(v map[string]*string) *InitCoursesOfClassResponse {
	s.Headers = v
	return s
}

func (s *InitCoursesOfClassResponse) SetBody(v *InitCoursesOfClassResponseBody) *InitCoursesOfClassResponse {
	s.Body = v
	return s
}

type CreateInviteUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateInviteUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateInviteUrlHeaders) GoString() string {
	return s.String()
}

func (s *CreateInviteUrlHeaders) SetCommonHeaders(v map[string]*string) *CreateInviteUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateInviteUrlHeaders) SetXAcsDingtalkAccessToken(v string) *CreateInviteUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateInviteUrlRequest struct {
	DingCorpId         *string `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int32  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	DingOauthAppId     *int64  `json:"dingOauthAppId,omitempty" xml:"dingOauthAppId,omitempty"`
	TargetCorpId       *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
	AuthCode           *string `json:"authCode,omitempty" xml:"authCode,omitempty"`
	TargetOperator     *string `json:"targetOperator,omitempty" xml:"targetOperator,omitempty"`
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
}

func (s CreateInviteUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInviteUrlRequest) GoString() string {
	return s.String()
}

func (s *CreateInviteUrlRequest) SetDingCorpId(v string) *CreateInviteUrlRequest {
	s.DingCorpId = &v
	return s
}

func (s *CreateInviteUrlRequest) SetDingSuiteKey(v string) *CreateInviteUrlRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *CreateInviteUrlRequest) SetDingTokenGrantType(v int32) *CreateInviteUrlRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *CreateInviteUrlRequest) SetDingOauthAppId(v int64) *CreateInviteUrlRequest {
	s.DingOauthAppId = &v
	return s
}

func (s *CreateInviteUrlRequest) SetTargetCorpId(v string) *CreateInviteUrlRequest {
	s.TargetCorpId = &v
	return s
}

func (s *CreateInviteUrlRequest) SetAuthCode(v string) *CreateInviteUrlRequest {
	s.AuthCode = &v
	return s
}

func (s *CreateInviteUrlRequest) SetTargetOperator(v string) *CreateInviteUrlRequest {
	s.TargetOperator = &v
	return s
}

func (s *CreateInviteUrlRequest) SetDingIsvOrgId(v int64) *CreateInviteUrlRequest {
	s.DingIsvOrgId = &v
	return s
}

type CreateInviteUrlResponseBody struct {
	Success *bool                              `json:"success,omitempty" xml:"success,omitempty"`
	Result  *CreateInviteUrlResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateInviteUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInviteUrlResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInviteUrlResponseBody) SetSuccess(v bool) *CreateInviteUrlResponseBody {
	s.Success = &v
	return s
}

func (s *CreateInviteUrlResponseBody) SetResult(v *CreateInviteUrlResponseBodyResult) *CreateInviteUrlResponseBody {
	s.Result = v
	return s
}

type CreateInviteUrlResponseBodyResult struct {
	InviteUrl *string `json:"inviteUrl,omitempty" xml:"inviteUrl,omitempty"`
}

func (s CreateInviteUrlResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateInviteUrlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateInviteUrlResponseBodyResult) SetInviteUrl(v string) *CreateInviteUrlResponseBodyResult {
	s.InviteUrl = &v
	return s
}

type CreateInviteUrlResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInviteUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInviteUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInviteUrlResponse) GoString() string {
	return s.String()
}

func (s *CreateInviteUrlResponse) SetHeaders(v map[string]*string) *CreateInviteUrlResponse {
	s.Headers = v
	return s
}

func (s *CreateInviteUrlResponse) SetBody(v *CreateInviteUrlResponseBody) *CreateInviteUrlResponse {
	s.Body = v
	return s
}

type PollingConfirmStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PollingConfirmStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s PollingConfirmStatusHeaders) GoString() string {
	return s.String()
}

func (s *PollingConfirmStatusHeaders) SetCommonHeaders(v map[string]*string) *PollingConfirmStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PollingConfirmStatusHeaders) SetXAcsDingtalkAccessToken(v string) *PollingConfirmStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PollingConfirmStatusRequest struct {
	// ext
	Ext *string `json:"ext,omitempty" xml:"ext,omitempty"`
	// isvCode
	IsvCode *string `json:"isvCode,omitempty" xml:"isvCode,omitempty"`
	// courseCode
	CourseCode *string `json:"courseCode,omitempty" xml:"courseCode,omitempty"`
	// opUserId
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s PollingConfirmStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s PollingConfirmStatusRequest) GoString() string {
	return s.String()
}

func (s *PollingConfirmStatusRequest) SetExt(v string) *PollingConfirmStatusRequest {
	s.Ext = &v
	return s
}

func (s *PollingConfirmStatusRequest) SetIsvCode(v string) *PollingConfirmStatusRequest {
	s.IsvCode = &v
	return s
}

func (s *PollingConfirmStatusRequest) SetCourseCode(v string) *PollingConfirmStatusRequest {
	s.CourseCode = &v
	return s
}

func (s *PollingConfirmStatusRequest) SetOpUserId(v string) *PollingConfirmStatusRequest {
	s.OpUserId = &v
	return s
}

type PollingConfirmStatusResponseBody struct {
	UniversityPollingCourseStatusResponse *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponse `json:"universityPollingCourseStatusResponse,omitempty" xml:"universityPollingCourseStatusResponse,omitempty" type:"Struct"`
}

func (s PollingConfirmStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PollingConfirmStatusResponseBody) GoString() string {
	return s.String()
}

func (s *PollingConfirmStatusResponseBody) SetUniversityPollingCourseStatusResponse(v *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponse) *PollingConfirmStatusResponseBody {
	s.UniversityPollingCourseStatusResponse = v
	return s
}

type PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponse struct {
	// 确认状态
	ConfirmStatus *bool `json:"confirmStatus,omitempty" xml:"confirmStatus,omitempty"`
	// 课程编码
	CourseCode       *string                                                                                  `json:"courseCode,omitempty" xml:"courseCode,omitempty"`
	LivePlayInfoList []*PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList `json:"livePlayInfoList,omitempty" xml:"livePlayInfoList,omitempty" type:"Repeated"`
}

func (s PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponse) GoString() string {
	return s.String()
}

func (s *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponse) SetConfirmStatus(v bool) *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponse {
	s.ConfirmStatus = &v
	return s
}

func (s *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponse) SetCourseCode(v string) *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponse {
	s.CourseCode = &v
	return s
}

func (s *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponse) SetLivePlayInfoList(v []*PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList) *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponse {
	s.LivePlayInfoList = v
	return s
}

type PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList struct {
	// 推流地址
	LiveInputUrl *string `json:"liveInputUrl,omitempty" xml:"liveInputUrl,omitempty"`
	// 直播拉流地址
	LiveOutputUrl *string `json:"liveOutputUrl,omitempty" xml:"liveOutputUrl,omitempty"`
	// 视频流类型
	LiveType *int64 `json:"liveType,omitempty" xml:"liveType,omitempty"`
	// 视频回放地址
	ReplayUrl *string `json:"replayUrl,omitempty" xml:"replayUrl,omitempty"`
}

func (s PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList) String() string {
	return tea.Prettify(s)
}

func (s PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList) GoString() string {
	return s.String()
}

func (s *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList) SetLiveInputUrl(v string) *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList {
	s.LiveInputUrl = &v
	return s
}

func (s *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList) SetLiveOutputUrl(v string) *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList {
	s.LiveOutputUrl = &v
	return s
}

func (s *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList) SetLiveType(v int64) *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList {
	s.LiveType = &v
	return s
}

func (s *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList) SetReplayUrl(v string) *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList {
	s.ReplayUrl = &v
	return s
}

type PollingConfirmStatusResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PollingConfirmStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PollingConfirmStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s PollingConfirmStatusResponse) GoString() string {
	return s.String()
}

func (s *PollingConfirmStatusResponse) SetHeaders(v map[string]*string) *PollingConfirmStatusResponse {
	s.Headers = v
	return s
}

func (s *PollingConfirmStatusResponse) SetBody(v *PollingConfirmStatusResponseBody) *PollingConfirmStatusResponse {
	s.Body = v
	return s
}

type DeleteTeacherHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteTeacherHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteTeacherHeaders) GoString() string {
	return s.String()
}

func (s *DeleteTeacherHeaders) SetCommonHeaders(v map[string]*string) *DeleteTeacherHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteTeacherHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteTeacherHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteTeacherRequest struct {
	// 是否班主任；1:班主任；0:非班主任
	Adviser *int32 `json:"adviser,omitempty" xml:"adviser,omitempty"`
	// 钉钉企业管理员员工ID
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
}

func (s DeleteTeacherRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTeacherRequest) GoString() string {
	return s.String()
}

func (s *DeleteTeacherRequest) SetAdviser(v int32) *DeleteTeacherRequest {
	s.Adviser = &v
	return s
}

func (s *DeleteTeacherRequest) SetOperator(v string) *DeleteTeacherRequest {
	s.Operator = &v
	return s
}

type DeleteTeacherResponseBody struct {
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteTeacherResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTeacherResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTeacherResponseBody) SetSuccess(v bool) *DeleteTeacherResponseBody {
	s.Success = &v
	return s
}

type DeleteTeacherResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTeacherResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTeacherResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTeacherResponse) GoString() string {
	return s.String()
}

func (s *DeleteTeacherResponse) SetHeaders(v map[string]*string) *DeleteTeacherResponse {
	s.Headers = v
	return s
}

func (s *DeleteTeacherResponse) SetBody(v *DeleteTeacherResponseBody) *DeleteTeacherResponse {
	s.Body = v
	return s
}

type DeleteDeviceOrgHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteDeviceOrgHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceOrgHeaders) GoString() string {
	return s.String()
}

func (s *DeleteDeviceOrgHeaders) SetCommonHeaders(v map[string]*string) *DeleteDeviceOrgHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteDeviceOrgHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteDeviceOrgHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteDeviceOrgRequest struct {
	DeviceCode *string `json:"deviceCode,omitempty" xml:"deviceCode,omitempty"`
	AuthCode   *string `json:"authCode,omitempty" xml:"authCode,omitempty"`
}

func (s DeleteDeviceOrgRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceOrgRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeviceOrgRequest) SetDeviceCode(v string) *DeleteDeviceOrgRequest {
	s.DeviceCode = &v
	return s
}

func (s *DeleteDeviceOrgRequest) SetAuthCode(v string) *DeleteDeviceOrgRequest {
	s.AuthCode = &v
	return s
}

type DeleteDeviceOrgResponseBody struct {
	// Id of the request
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteDeviceOrgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceOrgResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeviceOrgResponseBody) SetSuccess(v bool) *DeleteDeviceOrgResponseBody {
	s.Success = &v
	return s
}

type DeleteDeviceOrgResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDeviceOrgResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDeviceOrgResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceOrgResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeviceOrgResponse) SetHeaders(v map[string]*string) *DeleteDeviceOrgResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeviceOrgResponse) SetBody(v *DeleteDeviceOrgResponseBody) *DeleteDeviceOrgResponse {
	s.Body = v
	return s
}

type DeleteRemoteClassCourseHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteRemoteClassCourseHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteRemoteClassCourseHeaders) GoString() string {
	return s.String()
}

func (s *DeleteRemoteClassCourseHeaders) SetCommonHeaders(v map[string]*string) *DeleteRemoteClassCourseHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteRemoteClassCourseHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteRemoteClassCourseHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteRemoteClassCourseRequest struct {
	// 免登码
	AuthCode *string `json:"authCode,omitempty" xml:"authCode,omitempty"`
}

func (s DeleteRemoteClassCourseRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRemoteClassCourseRequest) GoString() string {
	return s.String()
}

func (s *DeleteRemoteClassCourseRequest) SetAuthCode(v string) *DeleteRemoteClassCourseRequest {
	s.AuthCode = &v
	return s
}

type DeleteRemoteClassCourseResponseBody struct {
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteRemoteClassCourseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRemoteClassCourseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRemoteClassCourseResponseBody) SetSuccess(v bool) *DeleteRemoteClassCourseResponseBody {
	s.Success = &v
	return s
}

type DeleteRemoteClassCourseResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRemoteClassCourseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRemoteClassCourseResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRemoteClassCourseResponse) GoString() string {
	return s.String()
}

func (s *DeleteRemoteClassCourseResponse) SetHeaders(v map[string]*string) *DeleteRemoteClassCourseResponse {
	s.Headers = v
	return s
}

func (s *DeleteRemoteClassCourseResponse) SetBody(v *DeleteRemoteClassCourseResponseBody) *DeleteRemoteClassCourseResponse {
	s.Body = v
	return s
}

type GetOpenCourseDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOpenCourseDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOpenCourseDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetOpenCourseDetailHeaders) SetCommonHeaders(v map[string]*string) *GetOpenCourseDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOpenCourseDetailHeaders) SetXAcsDingtalkAccessToken(v string) *GetOpenCourseDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOpenCourseDetailResponseBody struct {
	// 课程id
	CourseId *string `json:"courseId,omitempty" xml:"courseId,omitempty"`
	// 课程标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 课程类型: 0-直播 2-视频内容
	CourseType *int64 `json:"courseType,omitempty" xml:"courseType,omitempty"`
	// 老师名称
	TeacherName *string `json:"teacherName,omitempty" xml:"teacherName,omitempty"`
	// 封面图片地址
	CoverUrl *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	// 课程开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 老师的userId
	TeacherId *string `json:"teacherId,omitempty" xml:"teacherId,omitempty"`
	// 课程介绍
	Introduction *string `json:"introduction,omitempty" xml:"introduction,omitempty"`
	// 发布详情model
	PushModel *GetOpenCourseDetailResponseBodyPushModel `json:"pushModel,omitempty" xml:"pushModel,omitempty" type:"Struct"`
}

func (s GetOpenCourseDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOpenCourseDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpenCourseDetailResponseBody) SetCourseId(v string) *GetOpenCourseDetailResponseBody {
	s.CourseId = &v
	return s
}

func (s *GetOpenCourseDetailResponseBody) SetTitle(v string) *GetOpenCourseDetailResponseBody {
	s.Title = &v
	return s
}

func (s *GetOpenCourseDetailResponseBody) SetCourseType(v int64) *GetOpenCourseDetailResponseBody {
	s.CourseType = &v
	return s
}

func (s *GetOpenCourseDetailResponseBody) SetTeacherName(v string) *GetOpenCourseDetailResponseBody {
	s.TeacherName = &v
	return s
}

func (s *GetOpenCourseDetailResponseBody) SetCoverUrl(v string) *GetOpenCourseDetailResponseBody {
	s.CoverUrl = &v
	return s
}

func (s *GetOpenCourseDetailResponseBody) SetStartTime(v int64) *GetOpenCourseDetailResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetOpenCourseDetailResponseBody) SetTeacherId(v string) *GetOpenCourseDetailResponseBody {
	s.TeacherId = &v
	return s
}

func (s *GetOpenCourseDetailResponseBody) SetIntroduction(v string) *GetOpenCourseDetailResponseBody {
	s.Introduction = &v
	return s
}

func (s *GetOpenCourseDetailResponseBody) SetPushModel(v *GetOpenCourseDetailResponseBodyPushModel) *GetOpenCourseDetailResponseBody {
	s.PushModel = v
	return s
}

type GetOpenCourseDetailResponseBodyPushModel struct {
	// 参与学校的名称列表
	PushOrgNameList []*string `json:"pushOrgNameList,omitempty" xml:"pushOrgNameList,omitempty" type:"Repeated"`
	// 参与角色的名称列表
	PushRoleNameList []*string `json:"pushRoleNameList,omitempty" xml:"pushRoleNameList,omitempty" type:"Repeated"`
}

func (s GetOpenCourseDetailResponseBodyPushModel) String() string {
	return tea.Prettify(s)
}

func (s GetOpenCourseDetailResponseBodyPushModel) GoString() string {
	return s.String()
}

func (s *GetOpenCourseDetailResponseBodyPushModel) SetPushOrgNameList(v []*string) *GetOpenCourseDetailResponseBodyPushModel {
	s.PushOrgNameList = v
	return s
}

func (s *GetOpenCourseDetailResponseBodyPushModel) SetPushRoleNameList(v []*string) *GetOpenCourseDetailResponseBodyPushModel {
	s.PushRoleNameList = v
	return s
}

type GetOpenCourseDetailResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOpenCourseDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOpenCourseDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOpenCourseDetailResponse) GoString() string {
	return s.String()
}

func (s *GetOpenCourseDetailResponse) SetHeaders(v map[string]*string) *GetOpenCourseDetailResponse {
	s.Headers = v
	return s
}

func (s *GetOpenCourseDetailResponse) SetBody(v *GetOpenCourseDetailResponseBody) *GetOpenCourseDetailResponse {
	s.Body = v
	return s
}

type DeleteStudentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteStudentHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteStudentHeaders) GoString() string {
	return s.String()
}

func (s *DeleteStudentHeaders) SetCommonHeaders(v map[string]*string) *DeleteStudentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteStudentHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteStudentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteStudentRequest struct {
	// 钉钉管理员员工ID
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
}

func (s DeleteStudentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteStudentRequest) GoString() string {
	return s.String()
}

func (s *DeleteStudentRequest) SetOperator(v string) *DeleteStudentRequest {
	s.Operator = &v
	return s
}

type DeleteStudentResponseBody struct {
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteStudentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteStudentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStudentResponseBody) SetSuccess(v bool) *DeleteStudentResponseBody {
	s.Success = &v
	return s
}

type DeleteStudentResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteStudentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteStudentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteStudentResponse) GoString() string {
	return s.String()
}

func (s *DeleteStudentResponse) SetHeaders(v map[string]*string) *DeleteStudentResponse {
	s.Headers = v
	return s
}

func (s *DeleteStudentResponse) SetBody(v *DeleteStudentResponseBody) *DeleteStudentResponse {
	s.Body = v
	return s
}

type UpdateCoursesOfClassHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateCoursesOfClassHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoursesOfClassHeaders) GoString() string {
	return s.String()
}

func (s *UpdateCoursesOfClassHeaders) SetCommonHeaders(v map[string]*string) *UpdateCoursesOfClassHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateCoursesOfClassHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateCoursesOfClassHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateCoursesOfClassRequest struct {
	Courses []*UpdateCoursesOfClassRequestCourses `json:"courses,omitempty" xml:"courses,omitempty" type:"Repeated"`
	// 节次设置
	SectionConfig *UpdateCoursesOfClassRequestSectionConfig `json:"sectionConfig,omitempty" xml:"sectionConfig,omitempty" type:"Struct"`
	// 操作者id
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s UpdateCoursesOfClassRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoursesOfClassRequest) GoString() string {
	return s.String()
}

func (s *UpdateCoursesOfClassRequest) SetCourses(v []*UpdateCoursesOfClassRequestCourses) *UpdateCoursesOfClassRequest {
	s.Courses = v
	return s
}

func (s *UpdateCoursesOfClassRequest) SetSectionConfig(v *UpdateCoursesOfClassRequestSectionConfig) *UpdateCoursesOfClassRequest {
	s.SectionConfig = v
	return s
}

func (s *UpdateCoursesOfClassRequest) SetOpUserId(v string) *UpdateCoursesOfClassRequest {
	s.OpUserId = &v
	return s
}

type UpdateCoursesOfClassRequestCourses struct {
	// 老师Staffid
	TeacherStaffIds []*string `json:"teacherStaffIds,omitempty" xml:"teacherStaffIds,omitempty" type:"Repeated"`
	// 课程名称
	CourseName *string `json:"courseName,omitempty" xml:"courseName,omitempty"`
	// 上课日期
	DateModel *UpdateCoursesOfClassRequestCoursesDateModel `json:"dateModel,omitempty" xml:"dateModel,omitempty" type:"Struct"`
	// 节次模型
	SectionModel *UpdateCoursesOfClassRequestCoursesSectionModel `json:"sectionModel,omitempty" xml:"sectionModel,omitempty" type:"Struct"`
	// 创建者名字
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// 上课地点
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// 删除标记：要删除为ture
	DeleteTag *bool `json:"deleteTag,omitempty" xml:"deleteTag,omitempty"`
	// 课程code：删除/更新必填
	CourseCode *string `json:"courseCode,omitempty" xml:"courseCode,omitempty"`
	// 课组code
	CourseGroupCode *string `json:"courseGroupCode,omitempty" xml:"courseGroupCode,omitempty"`
}

func (s UpdateCoursesOfClassRequestCourses) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoursesOfClassRequestCourses) GoString() string {
	return s.String()
}

func (s *UpdateCoursesOfClassRequestCourses) SetTeacherStaffIds(v []*string) *UpdateCoursesOfClassRequestCourses {
	s.TeacherStaffIds = v
	return s
}

func (s *UpdateCoursesOfClassRequestCourses) SetCourseName(v string) *UpdateCoursesOfClassRequestCourses {
	s.CourseName = &v
	return s
}

func (s *UpdateCoursesOfClassRequestCourses) SetDateModel(v *UpdateCoursesOfClassRequestCoursesDateModel) *UpdateCoursesOfClassRequestCourses {
	s.DateModel = v
	return s
}

func (s *UpdateCoursesOfClassRequestCourses) SetSectionModel(v *UpdateCoursesOfClassRequestCoursesSectionModel) *UpdateCoursesOfClassRequestCourses {
	s.SectionModel = v
	return s
}

func (s *UpdateCoursesOfClassRequestCourses) SetCreatorName(v string) *UpdateCoursesOfClassRequestCourses {
	s.CreatorName = &v
	return s
}

func (s *UpdateCoursesOfClassRequestCourses) SetLocation(v string) *UpdateCoursesOfClassRequestCourses {
	s.Location = &v
	return s
}

func (s *UpdateCoursesOfClassRequestCourses) SetDeleteTag(v bool) *UpdateCoursesOfClassRequestCourses {
	s.DeleteTag = &v
	return s
}

func (s *UpdateCoursesOfClassRequestCourses) SetCourseCode(v string) *UpdateCoursesOfClassRequestCourses {
	s.CourseCode = &v
	return s
}

func (s *UpdateCoursesOfClassRequestCourses) SetCourseGroupCode(v string) *UpdateCoursesOfClassRequestCourses {
	s.CourseGroupCode = &v
	return s
}

type UpdateCoursesOfClassRequestCoursesDateModel struct {
	// month
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// year
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
	// dayOfMonth
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
}

func (s UpdateCoursesOfClassRequestCoursesDateModel) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoursesOfClassRequestCoursesDateModel) GoString() string {
	return s.String()
}

func (s *UpdateCoursesOfClassRequestCoursesDateModel) SetMonth(v int32) *UpdateCoursesOfClassRequestCoursesDateModel {
	s.Month = &v
	return s
}

func (s *UpdateCoursesOfClassRequestCoursesDateModel) SetYear(v int32) *UpdateCoursesOfClassRequestCoursesDateModel {
	s.Year = &v
	return s
}

func (s *UpdateCoursesOfClassRequestCoursesDateModel) SetDayOfMonth(v int32) *UpdateCoursesOfClassRequestCoursesDateModel {
	s.DayOfMonth = &v
	return s
}

type UpdateCoursesOfClassRequestCoursesSectionModel struct {
	// sectionType
	SectionType *string `json:"sectionType,omitempty" xml:"sectionType,omitempty"`
	// 节次index
	SectionIndex *int32 `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty"`
	// 节次名称
	SectionName *string `json:"sectionName,omitempty" xml:"sectionName,omitempty"`
}

func (s UpdateCoursesOfClassRequestCoursesSectionModel) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoursesOfClassRequestCoursesSectionModel) GoString() string {
	return s.String()
}

func (s *UpdateCoursesOfClassRequestCoursesSectionModel) SetSectionType(v string) *UpdateCoursesOfClassRequestCoursesSectionModel {
	s.SectionType = &v
	return s
}

func (s *UpdateCoursesOfClassRequestCoursesSectionModel) SetSectionIndex(v int32) *UpdateCoursesOfClassRequestCoursesSectionModel {
	s.SectionIndex = &v
	return s
}

func (s *UpdateCoursesOfClassRequestCoursesSectionModel) SetSectionName(v string) *UpdateCoursesOfClassRequestCoursesSectionModel {
	s.SectionName = &v
	return s
}

type UpdateCoursesOfClassRequestSectionConfig struct {
	// 节次模型
	SectionModels []*UpdateCoursesOfClassRequestSectionConfigSectionModels `json:"sectionModels,omitempty" xml:"sectionModels,omitempty" type:"Repeated"`
}

func (s UpdateCoursesOfClassRequestSectionConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoursesOfClassRequestSectionConfig) GoString() string {
	return s.String()
}

func (s *UpdateCoursesOfClassRequestSectionConfig) SetSectionModels(v []*UpdateCoursesOfClassRequestSectionConfigSectionModels) *UpdateCoursesOfClassRequestSectionConfig {
	s.SectionModels = v
	return s
}

type UpdateCoursesOfClassRequestSectionConfigSectionModels struct {
	// 节次类型枚举：COURSE/REST
	SectionType *string `json:"sectionType,omitempty" xml:"sectionType,omitempty"`
	// 开始时间
	Start *UpdateCoursesOfClassRequestSectionConfigSectionModelsStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
	// 第几节。
	SectionIndex *int32 `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty"`
	// 结束时间
	End *UpdateCoursesOfClassRequestSectionConfigSectionModelsEnd `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
}

func (s UpdateCoursesOfClassRequestSectionConfigSectionModels) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoursesOfClassRequestSectionConfigSectionModels) GoString() string {
	return s.String()
}

func (s *UpdateCoursesOfClassRequestSectionConfigSectionModels) SetSectionType(v string) *UpdateCoursesOfClassRequestSectionConfigSectionModels {
	s.SectionType = &v
	return s
}

func (s *UpdateCoursesOfClassRequestSectionConfigSectionModels) SetStart(v *UpdateCoursesOfClassRequestSectionConfigSectionModelsStart) *UpdateCoursesOfClassRequestSectionConfigSectionModels {
	s.Start = v
	return s
}

func (s *UpdateCoursesOfClassRequestSectionConfigSectionModels) SetSectionIndex(v int32) *UpdateCoursesOfClassRequestSectionConfigSectionModels {
	s.SectionIndex = &v
	return s
}

func (s *UpdateCoursesOfClassRequestSectionConfigSectionModels) SetEnd(v *UpdateCoursesOfClassRequestSectionConfigSectionModelsEnd) *UpdateCoursesOfClassRequestSectionConfigSectionModels {
	s.End = v
	return s
}

type UpdateCoursesOfClassRequestSectionConfigSectionModelsStart struct {
	// 分钟
	Min *int32 `json:"min,omitempty" xml:"min,omitempty"`
	// 小时
	Hour *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
}

func (s UpdateCoursesOfClassRequestSectionConfigSectionModelsStart) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoursesOfClassRequestSectionConfigSectionModelsStart) GoString() string {
	return s.String()
}

func (s *UpdateCoursesOfClassRequestSectionConfigSectionModelsStart) SetMin(v int32) *UpdateCoursesOfClassRequestSectionConfigSectionModelsStart {
	s.Min = &v
	return s
}

func (s *UpdateCoursesOfClassRequestSectionConfigSectionModelsStart) SetHour(v int32) *UpdateCoursesOfClassRequestSectionConfigSectionModelsStart {
	s.Hour = &v
	return s
}

type UpdateCoursesOfClassRequestSectionConfigSectionModelsEnd struct {
	// 分钟
	Min *int32 `json:"min,omitempty" xml:"min,omitempty"`
	// 小时
	Hour *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
}

func (s UpdateCoursesOfClassRequestSectionConfigSectionModelsEnd) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoursesOfClassRequestSectionConfigSectionModelsEnd) GoString() string {
	return s.String()
}

func (s *UpdateCoursesOfClassRequestSectionConfigSectionModelsEnd) SetMin(v int32) *UpdateCoursesOfClassRequestSectionConfigSectionModelsEnd {
	s.Min = &v
	return s
}

func (s *UpdateCoursesOfClassRequestSectionConfigSectionModelsEnd) SetHour(v int32) *UpdateCoursesOfClassRequestSectionConfigSectionModelsEnd {
	s.Hour = &v
	return s
}

type UpdateCoursesOfClassResponseBody struct {
	// 结果
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateCoursesOfClassResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoursesOfClassResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCoursesOfClassResponseBody) SetResult(v bool) *UpdateCoursesOfClassResponseBody {
	s.Result = &v
	return s
}

type UpdateCoursesOfClassResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCoursesOfClassResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCoursesOfClassResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoursesOfClassResponse) GoString() string {
	return s.String()
}

func (s *UpdateCoursesOfClassResponse) SetHeaders(v map[string]*string) *UpdateCoursesOfClassResponse {
	s.Headers = v
	return s
}

func (s *UpdateCoursesOfClassResponse) SetBody(v *UpdateCoursesOfClassResponseBody) *UpdateCoursesOfClassResponse {
	s.Body = v
	return s
}

type CreateSectionConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateSectionConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateSectionConfigHeaders) GoString() string {
	return s.String()
}

func (s *CreateSectionConfigHeaders) SetCommonHeaders(v map[string]*string) *CreateSectionConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSectionConfigHeaders) SetXAcsDingtalkAccessToken(v string) *CreateSectionConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateSectionConfigRequest struct {
	// 扩展参数
	Ext *string `json:"ext,omitempty" xml:"ext,omitempty"`
	// 课表模板信息
	SectionConfigs []*CreateSectionConfigRequestSectionConfigs `json:"sectionConfigs,omitempty" xml:"sectionConfigs,omitempty" type:"Repeated"`
	// 操作人的userid。
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s CreateSectionConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSectionConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateSectionConfigRequest) SetExt(v string) *CreateSectionConfigRequest {
	s.Ext = &v
	return s
}

func (s *CreateSectionConfigRequest) SetSectionConfigs(v []*CreateSectionConfigRequestSectionConfigs) *CreateSectionConfigRequest {
	s.SectionConfigs = v
	return s
}

func (s *CreateSectionConfigRequest) SetOpUserId(v string) *CreateSectionConfigRequest {
	s.OpUserId = &v
	return s
}

type CreateSectionConfigRequestSectionConfigs struct {
	// 学期
	Semester *int32 `json:"semester,omitempty" xml:"semester,omitempty"`
	// 开始时间（精确到日）
	SectionStartDate *CreateSectionConfigRequestSectionConfigsSectionStartDate `json:"sectionStartDate,omitempty" xml:"sectionStartDate,omitempty" type:"Struct"`
	// 学年
	SchoolYear *string `json:"schoolYear,omitempty" xml:"schoolYear,omitempty"`
	// 课表名称
	ScheduleName *string `json:"scheduleName,omitempty" xml:"scheduleName,omitempty"`
	// 节次模型
	SectionModels []*CreateSectionConfigRequestSectionConfigsSectionModels `json:"sectionModels,omitempty" xml:"sectionModels,omitempty" type:"Repeated"`
	// 结束时间
	SectionEndDate *CreateSectionConfigRequestSectionConfigsSectionEndDate `json:"sectionEndDate,omitempty" xml:"sectionEndDate,omitempty" type:"Struct"`
	// 学期开始时间
	SemesterStartDate *CreateSectionConfigRequestSectionConfigsSemesterStartDate `json:"semesterStartDate,omitempty" xml:"semesterStartDate,omitempty" type:"Struct"`
	// 学期结束时间
	SemesterEndDate *CreateSectionConfigRequestSectionConfigsSemesterEndDate `json:"semesterEndDate,omitempty" xml:"semesterEndDate,omitempty" type:"Struct"`
}

func (s CreateSectionConfigRequestSectionConfigs) String() string {
	return tea.Prettify(s)
}

func (s CreateSectionConfigRequestSectionConfigs) GoString() string {
	return s.String()
}

func (s *CreateSectionConfigRequestSectionConfigs) SetSemester(v int32) *CreateSectionConfigRequestSectionConfigs {
	s.Semester = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigs) SetSectionStartDate(v *CreateSectionConfigRequestSectionConfigsSectionStartDate) *CreateSectionConfigRequestSectionConfigs {
	s.SectionStartDate = v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigs) SetSchoolYear(v string) *CreateSectionConfigRequestSectionConfigs {
	s.SchoolYear = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigs) SetScheduleName(v string) *CreateSectionConfigRequestSectionConfigs {
	s.ScheduleName = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigs) SetSectionModels(v []*CreateSectionConfigRequestSectionConfigsSectionModels) *CreateSectionConfigRequestSectionConfigs {
	s.SectionModels = v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigs) SetSectionEndDate(v *CreateSectionConfigRequestSectionConfigsSectionEndDate) *CreateSectionConfigRequestSectionConfigs {
	s.SectionEndDate = v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigs) SetSemesterStartDate(v *CreateSectionConfigRequestSectionConfigsSemesterStartDate) *CreateSectionConfigRequestSectionConfigs {
	s.SemesterStartDate = v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigs) SetSemesterEndDate(v *CreateSectionConfigRequestSectionConfigsSemesterEndDate) *CreateSectionConfigRequestSectionConfigs {
	s.SemesterEndDate = v
	return s
}

type CreateSectionConfigRequestSectionConfigsSectionStartDate struct {
	// 月份。
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// 年份。
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
	// 每个月的第几天。
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
}

func (s CreateSectionConfigRequestSectionConfigsSectionStartDate) String() string {
	return tea.Prettify(s)
}

func (s CreateSectionConfigRequestSectionConfigsSectionStartDate) GoString() string {
	return s.String()
}

func (s *CreateSectionConfigRequestSectionConfigsSectionStartDate) SetMonth(v int32) *CreateSectionConfigRequestSectionConfigsSectionStartDate {
	s.Month = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSectionStartDate) SetYear(v int32) *CreateSectionConfigRequestSectionConfigsSectionStartDate {
	s.Year = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSectionStartDate) SetDayOfMonth(v int32) *CreateSectionConfigRequestSectionConfigsSectionStartDate {
	s.DayOfMonth = &v
	return s
}

type CreateSectionConfigRequestSectionConfigsSectionModels struct {
	// 节次类型枚举：COURSE/REST
	SectionType *string `json:"sectionType,omitempty" xml:"sectionType,omitempty"`
	// 开始时间
	SectionStartTime *CreateSectionConfigRequestSectionConfigsSectionModelsSectionStartTime `json:"sectionStartTime,omitempty" xml:"sectionStartTime,omitempty" type:"Struct"`
	// 第几节。
	SectionIndex *int32 `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty"`
	// 结束时间
	SectionEndTime *CreateSectionConfigRequestSectionConfigsSectionModelsSectionEndTime `json:"sectionEndTime,omitempty" xml:"sectionEndTime,omitempty" type:"Struct"`
	// 节次名称
	SectionName *string `json:"sectionName,omitempty" xml:"sectionName,omitempty"`
}

func (s CreateSectionConfigRequestSectionConfigsSectionModels) String() string {
	return tea.Prettify(s)
}

func (s CreateSectionConfigRequestSectionConfigsSectionModels) GoString() string {
	return s.String()
}

func (s *CreateSectionConfigRequestSectionConfigsSectionModels) SetSectionType(v string) *CreateSectionConfigRequestSectionConfigsSectionModels {
	s.SectionType = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSectionModels) SetSectionStartTime(v *CreateSectionConfigRequestSectionConfigsSectionModelsSectionStartTime) *CreateSectionConfigRequestSectionConfigsSectionModels {
	s.SectionStartTime = v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSectionModels) SetSectionIndex(v int32) *CreateSectionConfigRequestSectionConfigsSectionModels {
	s.SectionIndex = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSectionModels) SetSectionEndTime(v *CreateSectionConfigRequestSectionConfigsSectionModelsSectionEndTime) *CreateSectionConfigRequestSectionConfigsSectionModels {
	s.SectionEndTime = v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSectionModels) SetSectionName(v string) *CreateSectionConfigRequestSectionConfigsSectionModels {
	s.SectionName = &v
	return s
}

type CreateSectionConfigRequestSectionConfigsSectionModelsSectionStartTime struct {
	// 分
	Min *int32 `json:"min,omitempty" xml:"min,omitempty"`
	// 小时
	Hour *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
}

func (s CreateSectionConfigRequestSectionConfigsSectionModelsSectionStartTime) String() string {
	return tea.Prettify(s)
}

func (s CreateSectionConfigRequestSectionConfigsSectionModelsSectionStartTime) GoString() string {
	return s.String()
}

func (s *CreateSectionConfigRequestSectionConfigsSectionModelsSectionStartTime) SetMin(v int32) *CreateSectionConfigRequestSectionConfigsSectionModelsSectionStartTime {
	s.Min = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSectionModelsSectionStartTime) SetHour(v int32) *CreateSectionConfigRequestSectionConfigsSectionModelsSectionStartTime {
	s.Hour = &v
	return s
}

type CreateSectionConfigRequestSectionConfigsSectionModelsSectionEndTime struct {
	// 分
	Min *int32 `json:"min,omitempty" xml:"min,omitempty"`
	// 小时
	Hour *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
}

func (s CreateSectionConfigRequestSectionConfigsSectionModelsSectionEndTime) String() string {
	return tea.Prettify(s)
}

func (s CreateSectionConfigRequestSectionConfigsSectionModelsSectionEndTime) GoString() string {
	return s.String()
}

func (s *CreateSectionConfigRequestSectionConfigsSectionModelsSectionEndTime) SetMin(v int32) *CreateSectionConfigRequestSectionConfigsSectionModelsSectionEndTime {
	s.Min = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSectionModelsSectionEndTime) SetHour(v int32) *CreateSectionConfigRequestSectionConfigsSectionModelsSectionEndTime {
	s.Hour = &v
	return s
}

type CreateSectionConfigRequestSectionConfigsSectionEndDate struct {
	// 月
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// 年
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
	// 日
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
}

func (s CreateSectionConfigRequestSectionConfigsSectionEndDate) String() string {
	return tea.Prettify(s)
}

func (s CreateSectionConfigRequestSectionConfigsSectionEndDate) GoString() string {
	return s.String()
}

func (s *CreateSectionConfigRequestSectionConfigsSectionEndDate) SetMonth(v int32) *CreateSectionConfigRequestSectionConfigsSectionEndDate {
	s.Month = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSectionEndDate) SetYear(v int32) *CreateSectionConfigRequestSectionConfigsSectionEndDate {
	s.Year = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSectionEndDate) SetDayOfMonth(v int32) *CreateSectionConfigRequestSectionConfigsSectionEndDate {
	s.DayOfMonth = &v
	return s
}

type CreateSectionConfigRequestSectionConfigsSemesterStartDate struct {
	// 月
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// 年
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
	// 日
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
}

func (s CreateSectionConfigRequestSectionConfigsSemesterStartDate) String() string {
	return tea.Prettify(s)
}

func (s CreateSectionConfigRequestSectionConfigsSemesterStartDate) GoString() string {
	return s.String()
}

func (s *CreateSectionConfigRequestSectionConfigsSemesterStartDate) SetMonth(v int32) *CreateSectionConfigRequestSectionConfigsSemesterStartDate {
	s.Month = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSemesterStartDate) SetYear(v int32) *CreateSectionConfigRequestSectionConfigsSemesterStartDate {
	s.Year = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSemesterStartDate) SetDayOfMonth(v int32) *CreateSectionConfigRequestSectionConfigsSemesterStartDate {
	s.DayOfMonth = &v
	return s
}

type CreateSectionConfigRequestSectionConfigsSemesterEndDate struct {
	// 月
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// 年
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
	// 每月第几天
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
}

func (s CreateSectionConfigRequestSectionConfigsSemesterEndDate) String() string {
	return tea.Prettify(s)
}

func (s CreateSectionConfigRequestSectionConfigsSemesterEndDate) GoString() string {
	return s.String()
}

func (s *CreateSectionConfigRequestSectionConfigsSemesterEndDate) SetMonth(v int32) *CreateSectionConfigRequestSectionConfigsSemesterEndDate {
	s.Month = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSemesterEndDate) SetYear(v int32) *CreateSectionConfigRequestSectionConfigsSemesterEndDate {
	s.Year = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSemesterEndDate) SetDayOfMonth(v int32) *CreateSectionConfigRequestSectionConfigsSemesterEndDate {
	s.DayOfMonth = &v
	return s
}

type CreateSectionConfigResponseBody struct {
	// 初始化是否成功。
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateSectionConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSectionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSectionConfigResponseBody) SetResult(v bool) *CreateSectionConfigResponseBody {
	s.Result = &v
	return s
}

type CreateSectionConfigResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSectionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSectionConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSectionConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateSectionConfigResponse) SetHeaders(v map[string]*string) *CreateSectionConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateSectionConfigResponse) SetBody(v *CreateSectionConfigResponseBody) *CreateSectionConfigResponse {
	s.Body = v
	return s
}

type QuerySubjectTeachersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QuerySubjectTeachersHeaders) String() string {
	return tea.Prettify(s)
}

func (s QuerySubjectTeachersHeaders) GoString() string {
	return s.String()
}

func (s *QuerySubjectTeachersHeaders) SetCommonHeaders(v map[string]*string) *QuerySubjectTeachersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QuerySubjectTeachersHeaders) SetXAcsDingtalkAccessToken(v string) *QuerySubjectTeachersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QuerySubjectTeachersRequest struct {
	// 班级ids
	ClassIds []*int64 `json:"classIds,omitempty" xml:"classIds,omitempty" type:"Repeated"`
	// 操作人id
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 学科code
	SubjectCode *string `json:"subjectCode,omitempty" xml:"subjectCode,omitempty"`
}

func (s QuerySubjectTeachersRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySubjectTeachersRequest) GoString() string {
	return s.String()
}

func (s *QuerySubjectTeachersRequest) SetClassIds(v []*int64) *QuerySubjectTeachersRequest {
	s.ClassIds = v
	return s
}

func (s *QuerySubjectTeachersRequest) SetOpUserId(v string) *QuerySubjectTeachersRequest {
	s.OpUserId = &v
	return s
}

func (s *QuerySubjectTeachersRequest) SetSubjectCode(v string) *QuerySubjectTeachersRequest {
	s.SubjectCode = &v
	return s
}

type QuerySubjectTeachersResponseBody struct {
	// 结果
	Result []*QuerySubjectTeachersResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QuerySubjectTeachersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySubjectTeachersResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySubjectTeachersResponseBody) SetResult(v []*QuerySubjectTeachersResponseBodyResult) *QuerySubjectTeachersResponseBody {
	s.Result = v
	return s
}

type QuerySubjectTeachersResponseBodyResult struct {
	// 老师名称
	TeacherName *string `json:"teacherName,omitempty" xml:"teacherName,omitempty"`
	// 学科名称
	SubjectName *string `json:"subjectName,omitempty" xml:"subjectName,omitempty"`
	// 学科code
	SubjectCode *string `json:"subjectCode,omitempty" xml:"subjectCode,omitempty"`
	// 学段code
	PeriodCode *string `json:"periodCode,omitempty" xml:"periodCode,omitempty"`
	// 学校corpId
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 老师Userid
	TeacherUserId *string `json:"teacherUserId,omitempty" xml:"teacherUserId,omitempty"`
	// 班级id
	ClassId *int64 `json:"classId,omitempty" xml:"classId,omitempty"`
}

func (s QuerySubjectTeachersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QuerySubjectTeachersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QuerySubjectTeachersResponseBodyResult) SetTeacherName(v string) *QuerySubjectTeachersResponseBodyResult {
	s.TeacherName = &v
	return s
}

func (s *QuerySubjectTeachersResponseBodyResult) SetSubjectName(v string) *QuerySubjectTeachersResponseBodyResult {
	s.SubjectName = &v
	return s
}

func (s *QuerySubjectTeachersResponseBodyResult) SetSubjectCode(v string) *QuerySubjectTeachersResponseBodyResult {
	s.SubjectCode = &v
	return s
}

func (s *QuerySubjectTeachersResponseBodyResult) SetPeriodCode(v string) *QuerySubjectTeachersResponseBodyResult {
	s.PeriodCode = &v
	return s
}

func (s *QuerySubjectTeachersResponseBodyResult) SetCorpId(v string) *QuerySubjectTeachersResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *QuerySubjectTeachersResponseBodyResult) SetTeacherUserId(v string) *QuerySubjectTeachersResponseBodyResult {
	s.TeacherUserId = &v
	return s
}

func (s *QuerySubjectTeachersResponseBodyResult) SetClassId(v int64) *QuerySubjectTeachersResponseBodyResult {
	s.ClassId = &v
	return s
}

type QuerySubjectTeachersResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySubjectTeachersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySubjectTeachersResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySubjectTeachersResponse) GoString() string {
	return s.String()
}

func (s *QuerySubjectTeachersResponse) SetHeaders(v map[string]*string) *QuerySubjectTeachersResponse {
	s.Headers = v
	return s
}

func (s *QuerySubjectTeachersResponse) SetBody(v *QuerySubjectTeachersResponseBody) *QuerySubjectTeachersResponse {
	s.Body = v
	return s
}

type QueryRemoteClassCourseHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryRemoteClassCourseHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryRemoteClassCourseHeaders) GoString() string {
	return s.String()
}

func (s *QueryRemoteClassCourseHeaders) SetCommonHeaders(v map[string]*string) *QueryRemoteClassCourseHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryRemoteClassCourseHeaders) SetXAcsDingtalkAccessToken(v string) *QueryRemoteClassCourseHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryRemoteClassCourseRequest struct {
	// 操作者用户ID
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 结束时间
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
}

func (s QueryRemoteClassCourseRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRemoteClassCourseRequest) GoString() string {
	return s.String()
}

func (s *QueryRemoteClassCourseRequest) SetOperator(v string) *QueryRemoteClassCourseRequest {
	s.Operator = &v
	return s
}

func (s *QueryRemoteClassCourseRequest) SetStartTime(v int64) *QueryRemoteClassCourseRequest {
	s.StartTime = &v
	return s
}

func (s *QueryRemoteClassCourseRequest) SetEndTime(v int64) *QueryRemoteClassCourseRequest {
	s.EndTime = &v
	return s
}

type QueryRemoteClassCourseResponseBody struct {
	// 是否成功
	Success *bool                                       `json:"success,omitempty" xml:"success,omitempty"`
	Result  []*QueryRemoteClassCourseResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryRemoteClassCourseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRemoteClassCourseResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRemoteClassCourseResponseBody) SetSuccess(v bool) *QueryRemoteClassCourseResponseBody {
	s.Success = &v
	return s
}

func (s *QueryRemoteClassCourseResponseBody) SetResult(v []*QueryRemoteClassCourseResponseBodyResult) *QueryRemoteClassCourseResponseBody {
	s.Result = v
	return s
}

type QueryRemoteClassCourseResponseBodyResult struct {
	// 课程code
	CourseCode *string `json:"courseCode,omitempty" xml:"courseCode,omitempty"`
	// 课程名称
	CourseName *string `json:"courseName,omitempty" xml:"courseName,omitempty"`
	// 课程状态：0: 未开始；1: 已开始；2: 已结束
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// 开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 结束时间
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 课程是否可以编辑或删除
	CanEdit *bool `json:"canEdit,omitempty" xml:"canEdit,omitempty"`
	// 授课设备
	TeachingParticipant *QueryRemoteClassCourseResponseBodyResultTeachingParticipant `json:"teachingParticipant,omitempty" xml:"teachingParticipant,omitempty" type:"Struct"`
	// 听课设备列表
	AttendParticipants []*QueryRemoteClassCourseResponseBodyResultAttendParticipants `json:"attendParticipants,omitempty" xml:"attendParticipants,omitempty" type:"Repeated"`
}

func (s QueryRemoteClassCourseResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryRemoteClassCourseResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryRemoteClassCourseResponseBodyResult) SetCourseCode(v string) *QueryRemoteClassCourseResponseBodyResult {
	s.CourseCode = &v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResult) SetCourseName(v string) *QueryRemoteClassCourseResponseBodyResult {
	s.CourseName = &v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResult) SetStatus(v int32) *QueryRemoteClassCourseResponseBodyResult {
	s.Status = &v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResult) SetStartTime(v int64) *QueryRemoteClassCourseResponseBodyResult {
	s.StartTime = &v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResult) SetEndTime(v int64) *QueryRemoteClassCourseResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResult) SetCanEdit(v bool) *QueryRemoteClassCourseResponseBodyResult {
	s.CanEdit = &v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResult) SetTeachingParticipant(v *QueryRemoteClassCourseResponseBodyResultTeachingParticipant) *QueryRemoteClassCourseResponseBodyResult {
	s.TeachingParticipant = v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResult) SetAttendParticipants(v []*QueryRemoteClassCourseResponseBodyResultAttendParticipants) *QueryRemoteClassCourseResponseBodyResult {
	s.AttendParticipants = v
	return s
}

type QueryRemoteClassCourseResponseBodyResultTeachingParticipant struct {
	// 参与方名称
	ParticipantName *string `json:"participantName,omitempty" xml:"participantName,omitempty"`
	// 参与方ID
	ParticipantId *string `json:"participantId,omitempty" xml:"participantId,omitempty"`
	// 组织ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 组织名称
	OrgName *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
}

func (s QueryRemoteClassCourseResponseBodyResultTeachingParticipant) String() string {
	return tea.Prettify(s)
}

func (s QueryRemoteClassCourseResponseBodyResultTeachingParticipant) GoString() string {
	return s.String()
}

func (s *QueryRemoteClassCourseResponseBodyResultTeachingParticipant) SetParticipantName(v string) *QueryRemoteClassCourseResponseBodyResultTeachingParticipant {
	s.ParticipantName = &v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResultTeachingParticipant) SetParticipantId(v string) *QueryRemoteClassCourseResponseBodyResultTeachingParticipant {
	s.ParticipantId = &v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResultTeachingParticipant) SetCorpId(v string) *QueryRemoteClassCourseResponseBodyResultTeachingParticipant {
	s.CorpId = &v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResultTeachingParticipant) SetOrgName(v string) *QueryRemoteClassCourseResponseBodyResultTeachingParticipant {
	s.OrgName = &v
	return s
}

type QueryRemoteClassCourseResponseBodyResultAttendParticipants struct {
	// 参与方名称
	ParticipantName *string `json:"participantName,omitempty" xml:"participantName,omitempty"`
	// 参与方ID
	ParticipantId *string `json:"participantId,omitempty" xml:"participantId,omitempty"`
	// 组织ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 组织名称
	OrgName *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
}

func (s QueryRemoteClassCourseResponseBodyResultAttendParticipants) String() string {
	return tea.Prettify(s)
}

func (s QueryRemoteClassCourseResponseBodyResultAttendParticipants) GoString() string {
	return s.String()
}

func (s *QueryRemoteClassCourseResponseBodyResultAttendParticipants) SetParticipantName(v string) *QueryRemoteClassCourseResponseBodyResultAttendParticipants {
	s.ParticipantName = &v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResultAttendParticipants) SetParticipantId(v string) *QueryRemoteClassCourseResponseBodyResultAttendParticipants {
	s.ParticipantId = &v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResultAttendParticipants) SetCorpId(v string) *QueryRemoteClassCourseResponseBodyResultAttendParticipants {
	s.CorpId = &v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResultAttendParticipants) SetOrgName(v string) *QueryRemoteClassCourseResponseBodyResultAttendParticipants {
	s.OrgName = &v
	return s
}

type QueryRemoteClassCourseResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRemoteClassCourseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRemoteClassCourseResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRemoteClassCourseResponse) GoString() string {
	return s.String()
}

func (s *QueryRemoteClassCourseResponse) SetHeaders(v map[string]*string) *QueryRemoteClassCourseResponse {
	s.Headers = v
	return s
}

func (s *QueryRemoteClassCourseResponse) SetBody(v *QueryRemoteClassCourseResponseBody) *QueryRemoteClassCourseResponse {
	s.Body = v
	return s
}

type CreatePhysicalClassroomHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreatePhysicalClassroomHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreatePhysicalClassroomHeaders) GoString() string {
	return s.String()
}

func (s *CreatePhysicalClassroomHeaders) SetCommonHeaders(v map[string]*string) *CreatePhysicalClassroomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreatePhysicalClassroomHeaders) SetXAcsDingtalkAccessToken(v string) *CreatePhysicalClassroomHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreatePhysicalClassroomRequest struct {
	// opUserId
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 教室楼层
	ClassroomFloor *string `json:"classroomFloor,omitempty" xml:"classroomFloor,omitempty"`
	// 扩展信息
	Ext *string `json:"ext,omitempty" xml:"ext,omitempty"`
	// 教室教学楼
	ClassroomBuilding *string `json:"classroomBuilding,omitempty" xml:"classroomBuilding,omitempty"`
	// 是否支持直播
	DirectBroadcast *string `json:"directBroadcast,omitempty" xml:"directBroadcast,omitempty"`
	// 教室名称
	ClassroomName *string `json:"classroomName,omitempty" xml:"classroomName,omitempty"`
	// 教室校区
	ClassroomCampus *string `json:"classroomCampus,omitempty" xml:"classroomCampus,omitempty"`
	// 教室房间号
	ClassroomNumber *string `json:"classroomNumber,omitempty" xml:"classroomNumber,omitempty"`
}

func (s CreatePhysicalClassroomRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePhysicalClassroomRequest) GoString() string {
	return s.String()
}

func (s *CreatePhysicalClassroomRequest) SetOpUserId(v string) *CreatePhysicalClassroomRequest {
	s.OpUserId = &v
	return s
}

func (s *CreatePhysicalClassroomRequest) SetClassroomFloor(v string) *CreatePhysicalClassroomRequest {
	s.ClassroomFloor = &v
	return s
}

func (s *CreatePhysicalClassroomRequest) SetExt(v string) *CreatePhysicalClassroomRequest {
	s.Ext = &v
	return s
}

func (s *CreatePhysicalClassroomRequest) SetClassroomBuilding(v string) *CreatePhysicalClassroomRequest {
	s.ClassroomBuilding = &v
	return s
}

func (s *CreatePhysicalClassroomRequest) SetDirectBroadcast(v string) *CreatePhysicalClassroomRequest {
	s.DirectBroadcast = &v
	return s
}

func (s *CreatePhysicalClassroomRequest) SetClassroomName(v string) *CreatePhysicalClassroomRequest {
	s.ClassroomName = &v
	return s
}

func (s *CreatePhysicalClassroomRequest) SetClassroomCampus(v string) *CreatePhysicalClassroomRequest {
	s.ClassroomCampus = &v
	return s
}

func (s *CreatePhysicalClassroomRequest) SetClassroomNumber(v string) *CreatePhysicalClassroomRequest {
	s.ClassroomNumber = &v
	return s
}

type CreatePhysicalClassroomResponseBody struct {
	// 教室id
	ClassroomId *int64 `json:"classroomId,omitempty" xml:"classroomId,omitempty"`
}

func (s CreatePhysicalClassroomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePhysicalClassroomResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePhysicalClassroomResponseBody) SetClassroomId(v int64) *CreatePhysicalClassroomResponseBody {
	s.ClassroomId = &v
	return s
}

type CreatePhysicalClassroomResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatePhysicalClassroomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePhysicalClassroomResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePhysicalClassroomResponse) GoString() string {
	return s.String()
}

func (s *CreatePhysicalClassroomResponse) SetHeaders(v map[string]*string) *CreatePhysicalClassroomResponse {
	s.Headers = v
	return s
}

func (s *CreatePhysicalClassroomResponse) SetBody(v *CreatePhysicalClassroomResponseBody) *CreatePhysicalClassroomResponse {
	s.Body = v
	return s
}

type QueryClassScheduleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryClassScheduleHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleHeaders) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleHeaders) SetCommonHeaders(v map[string]*string) *QueryClassScheduleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryClassScheduleHeaders) SetXAcsDingtalkAccessToken(v string) *QueryClassScheduleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryClassScheduleRequest struct {
	// 订阅者类型：  DEPARTMENT：班级订阅 USER：老师订阅
	SubscriberType *string `json:"subscriberType,omitempty" xml:"subscriberType,omitempty"`
	// 开始时间（unix时间戳）
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 结束时间（unix时间戳）
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 操作者UserId
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 订阅者的Id。
	SubscriberIds []*string `json:"subscriberIds,omitempty" xml:"subscriberIds,omitempty" type:"Repeated"`
	// 查询课程的节次。
	SectionIndexList []*int64 `json:"sectionIndexList,omitempty" xml:"sectionIndexList,omitempty" type:"Repeated"`
}

func (s QueryClassScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleRequest) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleRequest) SetSubscriberType(v string) *QueryClassScheduleRequest {
	s.SubscriberType = &v
	return s
}

func (s *QueryClassScheduleRequest) SetStartTime(v int64) *QueryClassScheduleRequest {
	s.StartTime = &v
	return s
}

func (s *QueryClassScheduleRequest) SetEndTime(v int64) *QueryClassScheduleRequest {
	s.EndTime = &v
	return s
}

func (s *QueryClassScheduleRequest) SetOpUserId(v string) *QueryClassScheduleRequest {
	s.OpUserId = &v
	return s
}

func (s *QueryClassScheduleRequest) SetSubscriberIds(v []*string) *QueryClassScheduleRequest {
	s.SubscriberIds = v
	return s
}

func (s *QueryClassScheduleRequest) SetSectionIndexList(v []*int64) *QueryClassScheduleRequest {
	s.SectionIndexList = v
	return s
}

type QueryClassScheduleResponseBody struct {
	// 课表时间节次配置。
	Config *QueryClassScheduleResponseBodyConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// 课程列表
	CourseDTOS []*QueryClassScheduleResponseBodyCourseDTOS `json:"courseDTOS,omitempty" xml:"courseDTOS,omitempty" type:"Repeated"`
}

func (s QueryClassScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleResponseBody) SetConfig(v *QueryClassScheduleResponseBodyConfig) *QueryClassScheduleResponseBody {
	s.Config = v
	return s
}

func (s *QueryClassScheduleResponseBody) SetCourseDTOS(v []*QueryClassScheduleResponseBodyCourseDTOS) *QueryClassScheduleResponseBody {
	s.CourseDTOS = v
	return s
}

type QueryClassScheduleResponseBodyConfig struct {
	// 开始时间（到日）。
	Start *QueryClassScheduleResponseBodyConfigStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
	// 节次模型。
	SectionModels []*QueryClassScheduleResponseBodyConfigSectionModels `json:"sectionModels,omitempty" xml:"sectionModels,omitempty" type:"Repeated"`
	// 开始时间（到日）。
	End *QueryClassScheduleResponseBodyConfigEnd `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
}

func (s QueryClassScheduleResponseBodyConfig) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleResponseBodyConfig) SetStart(v *QueryClassScheduleResponseBodyConfigStart) *QueryClassScheduleResponseBodyConfig {
	s.Start = v
	return s
}

func (s *QueryClassScheduleResponseBodyConfig) SetSectionModels(v []*QueryClassScheduleResponseBodyConfigSectionModels) *QueryClassScheduleResponseBodyConfig {
	s.SectionModels = v
	return s
}

func (s *QueryClassScheduleResponseBodyConfig) SetEnd(v *QueryClassScheduleResponseBodyConfigEnd) *QueryClassScheduleResponseBodyConfig {
	s.End = v
	return s
}

type QueryClassScheduleResponseBodyConfigStart struct {
	// 年份。
	Year *int64 `json:"year,omitempty" xml:"year,omitempty"`
	// 月份。
	Month *int64 `json:"month,omitempty" xml:"month,omitempty"`
	// 一个月中第几天
	DayOfMonth *int64 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
}

func (s QueryClassScheduleResponseBodyConfigStart) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleResponseBodyConfigStart) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleResponseBodyConfigStart) SetYear(v int64) *QueryClassScheduleResponseBodyConfigStart {
	s.Year = &v
	return s
}

func (s *QueryClassScheduleResponseBodyConfigStart) SetMonth(v int64) *QueryClassScheduleResponseBodyConfigStart {
	s.Month = &v
	return s
}

func (s *QueryClassScheduleResponseBodyConfigStart) SetDayOfMonth(v int64) *QueryClassScheduleResponseBodyConfigStart {
	s.DayOfMonth = &v
	return s
}

type QueryClassScheduleResponseBodyConfigSectionModels struct {
	// 节次名称。
	SectionName *string `json:"sectionName,omitempty" xml:"sectionName,omitempty"`
	// 节次类型：  COURSE：上课节次 REST：休息节次
	SectionType *string `json:"sectionType,omitempty" xml:"sectionType,omitempty"`
	// 节次序列。
	SectionIndex *int64 `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty"`
	// 开始时间（时分级别）。
	Start *QueryClassScheduleResponseBodyConfigSectionModelsStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
	// 结束时间（时分级别）
	End *QueryClassScheduleResponseBodyConfigSectionModelsEnd `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
}

func (s QueryClassScheduleResponseBodyConfigSectionModels) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleResponseBodyConfigSectionModels) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleResponseBodyConfigSectionModels) SetSectionName(v string) *QueryClassScheduleResponseBodyConfigSectionModels {
	s.SectionName = &v
	return s
}

func (s *QueryClassScheduleResponseBodyConfigSectionModels) SetSectionType(v string) *QueryClassScheduleResponseBodyConfigSectionModels {
	s.SectionType = &v
	return s
}

func (s *QueryClassScheduleResponseBodyConfigSectionModels) SetSectionIndex(v int64) *QueryClassScheduleResponseBodyConfigSectionModels {
	s.SectionIndex = &v
	return s
}

func (s *QueryClassScheduleResponseBodyConfigSectionModels) SetStart(v *QueryClassScheduleResponseBodyConfigSectionModelsStart) *QueryClassScheduleResponseBodyConfigSectionModels {
	s.Start = v
	return s
}

func (s *QueryClassScheduleResponseBodyConfigSectionModels) SetEnd(v *QueryClassScheduleResponseBodyConfigSectionModelsEnd) *QueryClassScheduleResponseBodyConfigSectionModels {
	s.End = v
	return s
}

type QueryClassScheduleResponseBodyConfigSectionModelsStart struct {
	// 小时。
	Hour *int64 `json:"hour,omitempty" xml:"hour,omitempty"`
	// 分钟
	Min *int64 `json:"min,omitempty" xml:"min,omitempty"`
}

func (s QueryClassScheduleResponseBodyConfigSectionModelsStart) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleResponseBodyConfigSectionModelsStart) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleResponseBodyConfigSectionModelsStart) SetHour(v int64) *QueryClassScheduleResponseBodyConfigSectionModelsStart {
	s.Hour = &v
	return s
}

func (s *QueryClassScheduleResponseBodyConfigSectionModelsStart) SetMin(v int64) *QueryClassScheduleResponseBodyConfigSectionModelsStart {
	s.Min = &v
	return s
}

type QueryClassScheduleResponseBodyConfigSectionModelsEnd struct {
	// 小时。
	Hour *int64 `json:"hour,omitempty" xml:"hour,omitempty"`
	// 分钟。
	Min *int64 `json:"min,omitempty" xml:"min,omitempty"`
}

func (s QueryClassScheduleResponseBodyConfigSectionModelsEnd) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleResponseBodyConfigSectionModelsEnd) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleResponseBodyConfigSectionModelsEnd) SetHour(v int64) *QueryClassScheduleResponseBodyConfigSectionModelsEnd {
	s.Hour = &v
	return s
}

func (s *QueryClassScheduleResponseBodyConfigSectionModelsEnd) SetMin(v int64) *QueryClassScheduleResponseBodyConfigSectionModelsEnd {
	s.Min = &v
	return s
}

type QueryClassScheduleResponseBodyConfigEnd struct {
	// 年份。
	Year *int64 `json:"year,omitempty" xml:"year,omitempty"`
	// 月份。
	Month *int64 `json:"month,omitempty" xml:"month,omitempty"`
	// 一个月中第几天
	DayOfMonth *int64 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
}

func (s QueryClassScheduleResponseBodyConfigEnd) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleResponseBodyConfigEnd) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleResponseBodyConfigEnd) SetYear(v int64) *QueryClassScheduleResponseBodyConfigEnd {
	s.Year = &v
	return s
}

func (s *QueryClassScheduleResponseBodyConfigEnd) SetMonth(v int64) *QueryClassScheduleResponseBodyConfigEnd {
	s.Month = &v
	return s
}

func (s *QueryClassScheduleResponseBodyConfigEnd) SetDayOfMonth(v int64) *QueryClassScheduleResponseBodyConfigEnd {
	s.DayOfMonth = &v
	return s
}

type QueryClassScheduleResponseBodyCourseDTOS struct {
	// 课程编码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 课程名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 课程介绍
	Introduce *string `json:"introduce,omitempty" xml:"introduce,omitempty"`
	// 课程封面地址
	CoverUrl *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	// 开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 结束时间
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 创建者组织id
	CreatorCorpId *string `json:"creatorCorpId,omitempty" xml:"creatorCorpId,omitempty"`
	// 创建者UserId
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// 创建者UserName
	CreatorUserName *string `json:"creatorUserName,omitempty" xml:"creatorUserName,omitempty"`
	// 老师CorpId
	TeacherCorpId *string `json:"teacherCorpId,omitempty" xml:"teacherCorpId,omitempty"`
	// 老师UserId
	TeacherUserId *string `json:"teacherUserId,omitempty" xml:"teacherUserId,omitempty"`
	// 老师UserName
	TeacherUserName *string `json:"teacherUserName,omitempty" xml:"teacherUserName,omitempty"`
	// 学科编码
	SubjectCode *string `json:"subjectCode,omitempty" xml:"subjectCode,omitempty"`
	// 课程组编码
	CourseGroupCode *string `json:"courseGroupCode,omitempty" xml:"courseGroupCode,omitempty"`
	// 课程状态
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 课堂列表
	Classrooms []*QueryClassScheduleResponseBodyCourseDTOSClassrooms `json:"classrooms,omitempty" xml:"classrooms,omitempty" type:"Repeated"`
	// 课程参与人列表
	EduUserModels []*QueryClassScheduleResponseBodyCourseDTOSEduUserModels `json:"eduUserModels,omitempty" xml:"eduUserModels,omitempty" type:"Repeated"`
	// 课程名称
	SectionName *string `json:"sectionName,omitempty" xml:"sectionName,omitempty"`
	// 课程所在节次序列号
	SectionIndex *int64 `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty"`
	// 课程所在班级id
	ClassId *int64 `json:"classId,omitempty" xml:"classId,omitempty"`
	// 课程扩展信息
	ExtInfo *string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
}

func (s QueryClassScheduleResponseBodyCourseDTOS) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleResponseBodyCourseDTOS) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetCode(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.Code = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetName(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.Name = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetIntroduce(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.Introduce = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetCoverUrl(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.CoverUrl = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetStartTime(v int64) *QueryClassScheduleResponseBodyCourseDTOS {
	s.StartTime = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetEndTime(v int64) *QueryClassScheduleResponseBodyCourseDTOS {
	s.EndTime = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetCreatorCorpId(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.CreatorCorpId = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetCreatorUserId(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.CreatorUserId = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetCreatorUserName(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.CreatorUserName = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetTeacherCorpId(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.TeacherCorpId = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetTeacherUserId(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.TeacherUserId = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetTeacherUserName(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.TeacherUserName = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetSubjectCode(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.SubjectCode = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetCourseGroupCode(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.CourseGroupCode = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetStatus(v int64) *QueryClassScheduleResponseBodyCourseDTOS {
	s.Status = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetClassrooms(v []*QueryClassScheduleResponseBodyCourseDTOSClassrooms) *QueryClassScheduleResponseBodyCourseDTOS {
	s.Classrooms = v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetEduUserModels(v []*QueryClassScheduleResponseBodyCourseDTOSEduUserModels) *QueryClassScheduleResponseBodyCourseDTOS {
	s.EduUserModels = v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetSectionName(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.SectionName = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetSectionIndex(v int64) *QueryClassScheduleResponseBodyCourseDTOS {
	s.SectionIndex = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetClassId(v int64) *QueryClassScheduleResponseBodyCourseDTOS {
	s.ClassId = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetExtInfo(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.ExtInfo = &v
	return s
}

type QueryClassScheduleResponseBodyCourseDTOSClassrooms struct {
	// 课堂唯一标识
	TargetId *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
	// 交互信息
	InteractInfo *string `json:"interactInfo,omitempty" xml:"interactInfo,omitempty"`
}

func (s QueryClassScheduleResponseBodyCourseDTOSClassrooms) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleResponseBodyCourseDTOSClassrooms) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleResponseBodyCourseDTOSClassrooms) SetTargetId(v string) *QueryClassScheduleResponseBodyCourseDTOSClassrooms {
	s.TargetId = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOSClassrooms) SetInteractInfo(v string) *QueryClassScheduleResponseBodyCourseDTOSClassrooms {
	s.InteractInfo = &v
	return s
}

type QueryClassScheduleResponseBodyCourseDTOSEduUserModels struct {
	// 用户userid
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户uid
	Uid  *int64  `json:"uid,omitempty" xml:"uid,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryClassScheduleResponseBodyCourseDTOSEduUserModels) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleResponseBodyCourseDTOSEduUserModels) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleResponseBodyCourseDTOSEduUserModels) SetUserId(v string) *QueryClassScheduleResponseBodyCourseDTOSEduUserModels {
	s.UserId = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOSEduUserModels) SetUid(v int64) *QueryClassScheduleResponseBodyCourseDTOSEduUserModels {
	s.Uid = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOSEduUserModels) SetName(v string) *QueryClassScheduleResponseBodyCourseDTOSEduUserModels {
	s.Name = &v
	return s
}

type QueryClassScheduleResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryClassScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryClassScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleResponse) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleResponse) SetHeaders(v map[string]*string) *QueryClassScheduleResponse {
	s.Headers = v
	return s
}

func (s *QueryClassScheduleResponse) SetBody(v *QueryClassScheduleResponseBody) *QueryClassScheduleResponse {
	s.Body = v
	return s
}

type CreateRemoteClassCourseHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateRemoteClassCourseHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateRemoteClassCourseHeaders) GoString() string {
	return s.String()
}

func (s *CreateRemoteClassCourseHeaders) SetCommonHeaders(v map[string]*string) *CreateRemoteClassCourseHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateRemoteClassCourseHeaders) SetXAcsDingtalkAccessToken(v string) *CreateRemoteClassCourseHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateRemoteClassCourseRequest struct {
	// 课程名称
	CourseName *string `json:"courseName,omitempty" xml:"courseName,omitempty"`
	// 开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 结束时间
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 授课设备
	TeachingParticipant *CreateRemoteClassCourseRequestTeachingParticipant `json:"teachingParticipant,omitempty" xml:"teachingParticipant,omitempty" type:"Struct"`
	// 听课设备列表
	AttendParticipants []*CreateRemoteClassCourseRequestAttendParticipants `json:"attendParticipants,omitempty" xml:"attendParticipants,omitempty" type:"Repeated"`
	// 免登码
	AuthCode           *string `json:"authCode,omitempty" xml:"authCode,omitempty"`
	DingClientId       *string `json:"dingClientId,omitempty" xml:"dingClientId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int32  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	DingOauthAppId     *int64  `json:"dingOauthAppId,omitempty" xml:"dingOauthAppId,omitempty"`
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingCorpId         *string `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
}

func (s CreateRemoteClassCourseRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRemoteClassCourseRequest) GoString() string {
	return s.String()
}

func (s *CreateRemoteClassCourseRequest) SetCourseName(v string) *CreateRemoteClassCourseRequest {
	s.CourseName = &v
	return s
}

func (s *CreateRemoteClassCourseRequest) SetStartTime(v int64) *CreateRemoteClassCourseRequest {
	s.StartTime = &v
	return s
}

func (s *CreateRemoteClassCourseRequest) SetEndTime(v int64) *CreateRemoteClassCourseRequest {
	s.EndTime = &v
	return s
}

func (s *CreateRemoteClassCourseRequest) SetTeachingParticipant(v *CreateRemoteClassCourseRequestTeachingParticipant) *CreateRemoteClassCourseRequest {
	s.TeachingParticipant = v
	return s
}

func (s *CreateRemoteClassCourseRequest) SetAttendParticipants(v []*CreateRemoteClassCourseRequestAttendParticipants) *CreateRemoteClassCourseRequest {
	s.AttendParticipants = v
	return s
}

func (s *CreateRemoteClassCourseRequest) SetAuthCode(v string) *CreateRemoteClassCourseRequest {
	s.AuthCode = &v
	return s
}

func (s *CreateRemoteClassCourseRequest) SetDingClientId(v string) *CreateRemoteClassCourseRequest {
	s.DingClientId = &v
	return s
}

func (s *CreateRemoteClassCourseRequest) SetDingSuiteKey(v string) *CreateRemoteClassCourseRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *CreateRemoteClassCourseRequest) SetDingTokenGrantType(v int32) *CreateRemoteClassCourseRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *CreateRemoteClassCourseRequest) SetDingOauthAppId(v int64) *CreateRemoteClassCourseRequest {
	s.DingOauthAppId = &v
	return s
}

func (s *CreateRemoteClassCourseRequest) SetDingIsvOrgId(v int64) *CreateRemoteClassCourseRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *CreateRemoteClassCourseRequest) SetDingCorpId(v string) *CreateRemoteClassCourseRequest {
	s.DingCorpId = &v
	return s
}

type CreateRemoteClassCourseRequestTeachingParticipant struct {
	// 参与方ID
	ParticipantId *string `json:"participantId,omitempty" xml:"participantId,omitempty"`
	// 组织ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s CreateRemoteClassCourseRequestTeachingParticipant) String() string {
	return tea.Prettify(s)
}

func (s CreateRemoteClassCourseRequestTeachingParticipant) GoString() string {
	return s.String()
}

func (s *CreateRemoteClassCourseRequestTeachingParticipant) SetParticipantId(v string) *CreateRemoteClassCourseRequestTeachingParticipant {
	s.ParticipantId = &v
	return s
}

func (s *CreateRemoteClassCourseRequestTeachingParticipant) SetCorpId(v string) *CreateRemoteClassCourseRequestTeachingParticipant {
	s.CorpId = &v
	return s
}

type CreateRemoteClassCourseRequestAttendParticipants struct {
	// 参与方ID
	ParticipantId *string `json:"participantId,omitempty" xml:"participantId,omitempty"`
	// 组织ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s CreateRemoteClassCourseRequestAttendParticipants) String() string {
	return tea.Prettify(s)
}

func (s CreateRemoteClassCourseRequestAttendParticipants) GoString() string {
	return s.String()
}

func (s *CreateRemoteClassCourseRequestAttendParticipants) SetParticipantId(v string) *CreateRemoteClassCourseRequestAttendParticipants {
	s.ParticipantId = &v
	return s
}

func (s *CreateRemoteClassCourseRequestAttendParticipants) SetCorpId(v string) *CreateRemoteClassCourseRequestAttendParticipants {
	s.CorpId = &v
	return s
}

type CreateRemoteClassCourseResponseBody struct {
	Success *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
	Result  *CreateRemoteClassCourseResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateRemoteClassCourseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRemoteClassCourseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRemoteClassCourseResponseBody) SetSuccess(v bool) *CreateRemoteClassCourseResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRemoteClassCourseResponseBody) SetResult(v *CreateRemoteClassCourseResponseBodyResult) *CreateRemoteClassCourseResponseBody {
	s.Result = v
	return s
}

type CreateRemoteClassCourseResponseBodyResult struct {
	// 课程码
	CourseCode *string `json:"courseCode,omitempty" xml:"courseCode,omitempty"`
}

func (s CreateRemoteClassCourseResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateRemoteClassCourseResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateRemoteClassCourseResponseBodyResult) SetCourseCode(v string) *CreateRemoteClassCourseResponseBodyResult {
	s.CourseCode = &v
	return s
}

type CreateRemoteClassCourseResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRemoteClassCourseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRemoteClassCourseResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRemoteClassCourseResponse) GoString() string {
	return s.String()
}

func (s *CreateRemoteClassCourseResponse) SetHeaders(v map[string]*string) *CreateRemoteClassCourseResponse {
	s.Headers = v
	return s
}

func (s *CreateRemoteClassCourseResponse) SetBody(v *CreateRemoteClassCourseResponseBody) *CreateRemoteClassCourseResponse {
	s.Body = v
	return s
}

type CreateCustomClassHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateCustomClassHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomClassHeaders) GoString() string {
	return s.String()
}

func (s *CreateCustomClassHeaders) SetCommonHeaders(v map[string]*string) *CreateCustomClassHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateCustomClassHeaders) SetXAcsDingtalkAccessToken(v string) *CreateCustomClassHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateCustomClassRequest struct {
	// 班级信息
	CustomClass *CreateCustomClassRequestCustomClass `json:"customClass,omitempty" xml:"customClass,omitempty" type:"Struct"`
	// 上级部门ID
	SuperId *int64 `json:"superId,omitempty" xml:"superId,omitempty"`
	// 钉钉企业管理员工ID
	Operator           *string `json:"operator,omitempty" xml:"operator,omitempty"`
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingCorpId         *string `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	DingOauthAppId     *int64  `json:"dingOauthAppId,omitempty" xml:"dingOauthAppId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int32  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
}

func (s CreateCustomClassRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomClassRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomClassRequest) SetCustomClass(v *CreateCustomClassRequestCustomClass) *CreateCustomClassRequest {
	s.CustomClass = v
	return s
}

func (s *CreateCustomClassRequest) SetSuperId(v int64) *CreateCustomClassRequest {
	s.SuperId = &v
	return s
}

func (s *CreateCustomClassRequest) SetOperator(v string) *CreateCustomClassRequest {
	s.Operator = &v
	return s
}

func (s *CreateCustomClassRequest) SetDingIsvOrgId(v int64) *CreateCustomClassRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *CreateCustomClassRequest) SetDingCorpId(v string) *CreateCustomClassRequest {
	s.DingCorpId = &v
	return s
}

func (s *CreateCustomClassRequest) SetDingOauthAppId(v int64) *CreateCustomClassRequest {
	s.DingOauthAppId = &v
	return s
}

func (s *CreateCustomClassRequest) SetDingSuiteKey(v string) *CreateCustomClassRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *CreateCustomClassRequest) SetDingTokenGrantType(v int32) *CreateCustomClassRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *CreateCustomClassRequest) SetDingOrgId(v int64) *CreateCustomClassRequest {
	s.DingOrgId = &v
	return s
}

type CreateCustomClassRequestCustomClass struct {
	// 班级名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateCustomClassRequestCustomClass) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomClassRequestCustomClass) GoString() string {
	return s.String()
}

func (s *CreateCustomClassRequestCustomClass) SetName(v string) *CreateCustomClassRequestCustomClass {
	s.Name = &v
	return s
}

type CreateCustomClassResponseBody struct {
	// result
	Result *CreateCustomClassResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateCustomClassResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomClassResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomClassResponseBody) SetResult(v *CreateCustomClassResponseBodyResult) *CreateCustomClassResponseBody {
	s.Result = v
	return s
}

func (s *CreateCustomClassResponseBody) SetSuccess(v bool) *CreateCustomClassResponseBody {
	s.Success = &v
	return s
}

type CreateCustomClassResponseBodyResult struct {
	// 班级ID
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s CreateCustomClassResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomClassResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateCustomClassResponseBodyResult) SetDeptId(v int64) *CreateCustomClassResponseBodyResult {
	s.DeptId = &v
	return s
}

type CreateCustomClassResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCustomClassResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCustomClassResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomClassResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomClassResponse) SetHeaders(v map[string]*string) *CreateCustomClassResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomClassResponse) SetBody(v *CreateCustomClassResponseBody) *CreateCustomClassResponse {
	s.Body = v
	return s
}

type UpdatePhysicalClassroomHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdatePhysicalClassroomHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhysicalClassroomHeaders) GoString() string {
	return s.String()
}

func (s *UpdatePhysicalClassroomHeaders) SetCommonHeaders(v map[string]*string) *UpdatePhysicalClassroomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdatePhysicalClassroomHeaders) SetXAcsDingtalkAccessToken(v string) *UpdatePhysicalClassroomHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdatePhysicalClassroomRequest struct {
	// 操作人id
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 教室楼层
	ClassroomFloor *string `json:"classroomFloor,omitempty" xml:"classroomFloor,omitempty"`
	// 扩展信息
	Ext *string `json:"ext,omitempty" xml:"ext,omitempty"`
	// 教室教学楼
	ClassroomBuilding *string `json:"classroomBuilding,omitempty" xml:"classroomBuilding,omitempty"`
	// 是否支持直播
	DirectBroadcast *string `json:"directBroadcast,omitempty" xml:"directBroadcast,omitempty"`
	// 教室名称
	ClassroomName *string `json:"classroomName,omitempty" xml:"classroomName,omitempty"`
	// 教室校区
	ClassroomCampus *string `json:"classroomCampus,omitempty" xml:"classroomCampus,omitempty"`
	// 教室房间号
	ClassroomNumber *string `json:"classroomNumber,omitempty" xml:"classroomNumber,omitempty"`
	// 教室id
	ClassroomId *int64 `json:"classroomId,omitempty" xml:"classroomId,omitempty"`
}

func (s UpdatePhysicalClassroomRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhysicalClassroomRequest) GoString() string {
	return s.String()
}

func (s *UpdatePhysicalClassroomRequest) SetOpUserId(v string) *UpdatePhysicalClassroomRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdatePhysicalClassroomRequest) SetClassroomFloor(v string) *UpdatePhysicalClassroomRequest {
	s.ClassroomFloor = &v
	return s
}

func (s *UpdatePhysicalClassroomRequest) SetExt(v string) *UpdatePhysicalClassroomRequest {
	s.Ext = &v
	return s
}

func (s *UpdatePhysicalClassroomRequest) SetClassroomBuilding(v string) *UpdatePhysicalClassroomRequest {
	s.ClassroomBuilding = &v
	return s
}

func (s *UpdatePhysicalClassroomRequest) SetDirectBroadcast(v string) *UpdatePhysicalClassroomRequest {
	s.DirectBroadcast = &v
	return s
}

func (s *UpdatePhysicalClassroomRequest) SetClassroomName(v string) *UpdatePhysicalClassroomRequest {
	s.ClassroomName = &v
	return s
}

func (s *UpdatePhysicalClassroomRequest) SetClassroomCampus(v string) *UpdatePhysicalClassroomRequest {
	s.ClassroomCampus = &v
	return s
}

func (s *UpdatePhysicalClassroomRequest) SetClassroomNumber(v string) *UpdatePhysicalClassroomRequest {
	s.ClassroomNumber = &v
	return s
}

func (s *UpdatePhysicalClassroomRequest) SetClassroomId(v int64) *UpdatePhysicalClassroomRequest {
	s.ClassroomId = &v
	return s
}

type UpdatePhysicalClassroomResponseBody struct {
	// 是否成功
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdatePhysicalClassroomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhysicalClassroomResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePhysicalClassroomResponseBody) SetResult(v bool) *UpdatePhysicalClassroomResponseBody {
	s.Result = &v
	return s
}

type UpdatePhysicalClassroomResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdatePhysicalClassroomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePhysicalClassroomResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhysicalClassroomResponse) GoString() string {
	return s.String()
}

func (s *UpdatePhysicalClassroomResponse) SetHeaders(v map[string]*string) *UpdatePhysicalClassroomResponse {
	s.Headers = v
	return s
}

func (s *UpdatePhysicalClassroomResponse) SetBody(v *UpdatePhysicalClassroomResponseBody) *UpdatePhysicalClassroomResponse {
	s.Body = v
	return s
}

type MoveStudentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MoveStudentHeaders) String() string {
	return tea.Prettify(s)
}

func (s MoveStudentHeaders) GoString() string {
	return s.String()
}

func (s *MoveStudentHeaders) SetCommonHeaders(v map[string]*string) *MoveStudentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MoveStudentHeaders) SetXAcsDingtalkAccessToken(v string) *MoveStudentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MoveStudentRequest struct {
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingCorpId         *string `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	DingOauthAppId     *int64  `json:"dingOauthAppId,omitempty" xml:"dingOauthAppId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int32  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	// 管理员id
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 学生id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 愿班级id
	OriginClassId *int64 `json:"originClassId,omitempty" xml:"originClassId,omitempty"`
	// 目标班级id
	TargetClassId *int64 `json:"targetClassId,omitempty" xml:"targetClassId,omitempty"`
}

func (s MoveStudentRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveStudentRequest) GoString() string {
	return s.String()
}

func (s *MoveStudentRequest) SetDingIsvOrgId(v int64) *MoveStudentRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *MoveStudentRequest) SetDingCorpId(v string) *MoveStudentRequest {
	s.DingCorpId = &v
	return s
}

func (s *MoveStudentRequest) SetDingOauthAppId(v int64) *MoveStudentRequest {
	s.DingOauthAppId = &v
	return s
}

func (s *MoveStudentRequest) SetDingSuiteKey(v string) *MoveStudentRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *MoveStudentRequest) SetDingTokenGrantType(v int32) *MoveStudentRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *MoveStudentRequest) SetDingOrgId(v int64) *MoveStudentRequest {
	s.DingOrgId = &v
	return s
}

func (s *MoveStudentRequest) SetOperator(v string) *MoveStudentRequest {
	s.Operator = &v
	return s
}

func (s *MoveStudentRequest) SetUserId(v string) *MoveStudentRequest {
	s.UserId = &v
	return s
}

func (s *MoveStudentRequest) SetOriginClassId(v int64) *MoveStudentRequest {
	s.OriginClassId = &v
	return s
}

func (s *MoveStudentRequest) SetTargetClassId(v int64) *MoveStudentRequest {
	s.TargetClassId = &v
	return s
}

type MoveStudentResponseBody struct {
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s MoveStudentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveStudentResponseBody) GoString() string {
	return s.String()
}

func (s *MoveStudentResponseBody) SetSuccess(v bool) *MoveStudentResponseBody {
	s.Success = &v
	return s
}

type MoveStudentResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MoveStudentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MoveStudentResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveStudentResponse) GoString() string {
	return s.String()
}

func (s *MoveStudentResponse) SetHeaders(v map[string]*string) *MoveStudentResponse {
	s.Headers = v
	return s
}

func (s *MoveStudentResponse) SetBody(v *MoveStudentResponseBody) *MoveStudentResponse {
	s.Body = v
	return s
}

type CreateUniversityCourseGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateUniversityCourseGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityCourseGroupHeaders) GoString() string {
	return s.String()
}

func (s *CreateUniversityCourseGroupHeaders) SetCommonHeaders(v map[string]*string) *CreateUniversityCourseGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateUniversityCourseGroupHeaders) SetXAcsDingtalkAccessToken(v string) *CreateUniversityCourseGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateUniversityCourseGroupRequest struct {
	// 操作人
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 合作方课程组code
	IsvCourseGroupCode *string `json:"isvCourseGroupCode,omitempty" xml:"isvCourseGroupCode,omitempty"`
	// 扩展参数
	Ext *string `json:"ext,omitempty" xml:"ext,omitempty"`
	// 课程组介绍
	CourseGroupIntroduce *string `json:"courseGroupIntroduce,omitempty" xml:"courseGroupIntroduce,omitempty"`
	// 学期
	Semester *int32 `json:"semester,omitempty" xml:"semester,omitempty"`
	// 学科
	SubjectName *string `json:"subjectName,omitempty" xml:"subjectName,omitempty"`
	// 课程组名称
	CourseGroupName *string `json:"courseGroupName,omitempty" xml:"courseGroupName,omitempty"`
	// 学年
	SchoolYear *string `json:"schoolYear,omitempty" xml:"schoolYear,omitempty"`
	// 学段code
	PeriodCode *string `json:"periodCode,omitempty" xml:"periodCode,omitempty"`
	// 教师信息
	TeacherInfos []*CreateUniversityCourseGroupRequestTeacherInfos `json:"teacherInfos,omitempty" xml:"teacherInfos,omitempty" type:"Repeated"`
	// 课程详细
	CourserGroupItemModels []*CreateUniversityCourseGroupRequestCourserGroupItemModels `json:"courserGroupItemModels,omitempty" xml:"courserGroupItemModels,omitempty" type:"Repeated"`
}

func (s CreateUniversityCourseGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityCourseGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateUniversityCourseGroupRequest) SetOpUserId(v string) *CreateUniversityCourseGroupRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateUniversityCourseGroupRequest) SetIsvCourseGroupCode(v string) *CreateUniversityCourseGroupRequest {
	s.IsvCourseGroupCode = &v
	return s
}

func (s *CreateUniversityCourseGroupRequest) SetExt(v string) *CreateUniversityCourseGroupRequest {
	s.Ext = &v
	return s
}

func (s *CreateUniversityCourseGroupRequest) SetCourseGroupIntroduce(v string) *CreateUniversityCourseGroupRequest {
	s.CourseGroupIntroduce = &v
	return s
}

func (s *CreateUniversityCourseGroupRequest) SetSemester(v int32) *CreateUniversityCourseGroupRequest {
	s.Semester = &v
	return s
}

func (s *CreateUniversityCourseGroupRequest) SetSubjectName(v string) *CreateUniversityCourseGroupRequest {
	s.SubjectName = &v
	return s
}

func (s *CreateUniversityCourseGroupRequest) SetCourseGroupName(v string) *CreateUniversityCourseGroupRequest {
	s.CourseGroupName = &v
	return s
}

func (s *CreateUniversityCourseGroupRequest) SetSchoolYear(v string) *CreateUniversityCourseGroupRequest {
	s.SchoolYear = &v
	return s
}

func (s *CreateUniversityCourseGroupRequest) SetPeriodCode(v string) *CreateUniversityCourseGroupRequest {
	s.PeriodCode = &v
	return s
}

func (s *CreateUniversityCourseGroupRequest) SetTeacherInfos(v []*CreateUniversityCourseGroupRequestTeacherInfos) *CreateUniversityCourseGroupRequest {
	s.TeacherInfos = v
	return s
}

func (s *CreateUniversityCourseGroupRequest) SetCourserGroupItemModels(v []*CreateUniversityCourseGroupRequestCourserGroupItemModels) *CreateUniversityCourseGroupRequest {
	s.CourserGroupItemModels = v
	return s
}

type CreateUniversityCourseGroupRequestTeacherInfos struct {
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 角色类型
	ParticipantRole *string `json:"participantRole,omitempty" xml:"participantRole,omitempty"`
}

func (s CreateUniversityCourseGroupRequestTeacherInfos) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityCourseGroupRequestTeacherInfos) GoString() string {
	return s.String()
}

func (s *CreateUniversityCourseGroupRequestTeacherInfos) SetUserId(v string) *CreateUniversityCourseGroupRequestTeacherInfos {
	s.UserId = &v
	return s
}

func (s *CreateUniversityCourseGroupRequestTeacherInfos) SetParticipantRole(v string) *CreateUniversityCourseGroupRequestTeacherInfos {
	s.ParticipantRole = &v
	return s
}

type CreateUniversityCourseGroupRequestCourserGroupItemModels struct {
	// 一周的第几天
	DayOfWeek *int32 `json:"dayOfWeek,omitempty" xml:"dayOfWeek,omitempty"`
	// 上课周期
	ClassPeriodType *int32 `json:"classPeriodType,omitempty" xml:"classPeriodType,omitempty"`
	// 课程组详细开始时间
	CourserGroupItemStartDate *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate `json:"courserGroupItemStartDate,omitempty" xml:"courserGroupItemStartDate,omitempty" type:"Struct"`
	// 课节
	SectionIndex []*int32 `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty" type:"Repeated"`
	// 课程组详细结束时间
	CourserGroupItemEndDate *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate `json:"courserGroupItemEndDate,omitempty" xml:"courserGroupItemEndDate,omitempty" type:"Struct"`
	// 课程类型
	CourseType *int32 `json:"courseType,omitempty" xml:"courseType,omitempty"`
	// 教室id
	ClassroomId *int64 `json:"classroomId,omitempty" xml:"classroomId,omitempty"`
}

func (s CreateUniversityCourseGroupRequestCourserGroupItemModels) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityCourseGroupRequestCourserGroupItemModels) GoString() string {
	return s.String()
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModels) SetDayOfWeek(v int32) *CreateUniversityCourseGroupRequestCourserGroupItemModels {
	s.DayOfWeek = &v
	return s
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModels) SetClassPeriodType(v int32) *CreateUniversityCourseGroupRequestCourserGroupItemModels {
	s.ClassPeriodType = &v
	return s
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModels) SetCourserGroupItemStartDate(v *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate) *CreateUniversityCourseGroupRequestCourserGroupItemModels {
	s.CourserGroupItemStartDate = v
	return s
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModels) SetSectionIndex(v []*int32) *CreateUniversityCourseGroupRequestCourserGroupItemModels {
	s.SectionIndex = v
	return s
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModels) SetCourserGroupItemEndDate(v *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate) *CreateUniversityCourseGroupRequestCourserGroupItemModels {
	s.CourserGroupItemEndDate = v
	return s
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModels) SetCourseType(v int32) *CreateUniversityCourseGroupRequestCourserGroupItemModels {
	s.CourseType = &v
	return s
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModels) SetClassroomId(v int64) *CreateUniversityCourseGroupRequestCourserGroupItemModels {
	s.ClassroomId = &v
	return s
}

type CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate struct {
	// 月
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// 年
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
	// 一月的第几天
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
}

func (s CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate) GoString() string {
	return s.String()
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate) SetMonth(v int32) *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate {
	s.Month = &v
	return s
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate) SetYear(v int32) *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate {
	s.Year = &v
	return s
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate) SetDayOfMonth(v int32) *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate {
	s.DayOfMonth = &v
	return s
}

type CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate struct {
	// 月
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// 年
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
	// 一月的第几天
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
}

func (s CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate) GoString() string {
	return s.String()
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate) SetMonth(v int32) *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate {
	s.Month = &v
	return s
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate) SetYear(v int32) *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate {
	s.Year = &v
	return s
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate) SetDayOfMonth(v int32) *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate {
	s.DayOfMonth = &v
	return s
}

type CreateUniversityCourseGroupResponseBody struct {
	// 课程组信息
	CourseGroupInfo *CreateUniversityCourseGroupResponseBodyCourseGroupInfo `json:"courseGroupInfo,omitempty" xml:"courseGroupInfo,omitempty" type:"Struct"`
}

func (s CreateUniversityCourseGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityCourseGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUniversityCourseGroupResponseBody) SetCourseGroupInfo(v *CreateUniversityCourseGroupResponseBodyCourseGroupInfo) *CreateUniversityCourseGroupResponseBody {
	s.CourseGroupInfo = v
	return s
}

type CreateUniversityCourseGroupResponseBodyCourseGroupInfo struct {
	// 课程组编码
	CourseGroupCode *string `json:"courseGroupCode,omitempty" xml:"courseGroupCode,omitempty"`
}

func (s CreateUniversityCourseGroupResponseBodyCourseGroupInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityCourseGroupResponseBodyCourseGroupInfo) GoString() string {
	return s.String()
}

func (s *CreateUniversityCourseGroupResponseBodyCourseGroupInfo) SetCourseGroupCode(v string) *CreateUniversityCourseGroupResponseBodyCourseGroupInfo {
	s.CourseGroupCode = &v
	return s
}

type CreateUniversityCourseGroupResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUniversityCourseGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUniversityCourseGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityCourseGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateUniversityCourseGroupResponse) SetHeaders(v map[string]*string) *CreateUniversityCourseGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateUniversityCourseGroupResponse) SetBody(v *CreateUniversityCourseGroupResponseBody) *CreateUniversityCourseGroupResponse {
	s.Body = v
	return s
}

type GetRemoteClassCourseHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetRemoteClassCourseHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetRemoteClassCourseHeaders) GoString() string {
	return s.String()
}

func (s *GetRemoteClassCourseHeaders) SetCommonHeaders(v map[string]*string) *GetRemoteClassCourseHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRemoteClassCourseHeaders) SetXAcsDingtalkAccessToken(v string) *GetRemoteClassCourseHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetRemoteClassCourseRequest struct {
	// 操作者用户ID
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
}

func (s GetRemoteClassCourseRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRemoteClassCourseRequest) GoString() string {
	return s.String()
}

func (s *GetRemoteClassCourseRequest) SetOperator(v string) *GetRemoteClassCourseRequest {
	s.Operator = &v
	return s
}

type GetRemoteClassCourseResponseBody struct {
	// 是否成功
	Success *bool                                   `json:"success,omitempty" xml:"success,omitempty"`
	Result  *GetRemoteClassCourseResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetRemoteClassCourseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRemoteClassCourseResponseBody) GoString() string {
	return s.String()
}

func (s *GetRemoteClassCourseResponseBody) SetSuccess(v bool) *GetRemoteClassCourseResponseBody {
	s.Success = &v
	return s
}

func (s *GetRemoteClassCourseResponseBody) SetResult(v *GetRemoteClassCourseResponseBodyResult) *GetRemoteClassCourseResponseBody {
	s.Result = v
	return s
}

type GetRemoteClassCourseResponseBodyResult struct {
	// 课程code
	CourseCode *string `json:"courseCode,omitempty" xml:"courseCode,omitempty"`
	// 课程名称
	CourseName *string `json:"courseName,omitempty" xml:"courseName,omitempty"`
	// 开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 结束时间
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 课程状态：0: 未开始；1: 已开始；2: 已结束
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// 课堂当前状态：0: 未进行；1: 进行中
	RoomStatus *int32 `json:"roomStatus,omitempty" xml:"roomStatus,omitempty"`
	// 课程是否可以编辑或删除
	CanEdit *bool `json:"canEdit,omitempty" xml:"canEdit,omitempty"`
	// 授课设备
	TeachingParticipant *GetRemoteClassCourseResponseBodyResultTeachingParticipant `json:"teachingParticipant,omitempty" xml:"teachingParticipant,omitempty" type:"Struct"`
	// 听课设备列表
	AttendParticipants []*GetRemoteClassCourseResponseBodyResultAttendParticipants `json:"attendParticipants,omitempty" xml:"attendParticipants,omitempty" type:"Repeated"`
}

func (s GetRemoteClassCourseResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetRemoteClassCourseResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRemoteClassCourseResponseBodyResult) SetCourseCode(v string) *GetRemoteClassCourseResponseBodyResult {
	s.CourseCode = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResult) SetCourseName(v string) *GetRemoteClassCourseResponseBodyResult {
	s.CourseName = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResult) SetStartTime(v int64) *GetRemoteClassCourseResponseBodyResult {
	s.StartTime = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResult) SetEndTime(v int64) *GetRemoteClassCourseResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResult) SetStatus(v int32) *GetRemoteClassCourseResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResult) SetRoomStatus(v int32) *GetRemoteClassCourseResponseBodyResult {
	s.RoomStatus = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResult) SetCanEdit(v bool) *GetRemoteClassCourseResponseBodyResult {
	s.CanEdit = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResult) SetTeachingParticipant(v *GetRemoteClassCourseResponseBodyResultTeachingParticipant) *GetRemoteClassCourseResponseBodyResult {
	s.TeachingParticipant = v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResult) SetAttendParticipants(v []*GetRemoteClassCourseResponseBodyResultAttendParticipants) *GetRemoteClassCourseResponseBodyResult {
	s.AttendParticipants = v
	return s
}

type GetRemoteClassCourseResponseBodyResultTeachingParticipant struct {
	// 参与方ID
	ParticipantId *string `json:"participantId,omitempty" xml:"participantId,omitempty"`
	// 参与方名称
	ParticipantName *string `json:"participantName,omitempty" xml:"participantName,omitempty"`
	// 组织ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 组织名称
	OrgName *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
}

func (s GetRemoteClassCourseResponseBodyResultTeachingParticipant) String() string {
	return tea.Prettify(s)
}

func (s GetRemoteClassCourseResponseBodyResultTeachingParticipant) GoString() string {
	return s.String()
}

func (s *GetRemoteClassCourseResponseBodyResultTeachingParticipant) SetParticipantId(v string) *GetRemoteClassCourseResponseBodyResultTeachingParticipant {
	s.ParticipantId = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResultTeachingParticipant) SetParticipantName(v string) *GetRemoteClassCourseResponseBodyResultTeachingParticipant {
	s.ParticipantName = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResultTeachingParticipant) SetCorpId(v string) *GetRemoteClassCourseResponseBodyResultTeachingParticipant {
	s.CorpId = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResultTeachingParticipant) SetOrgName(v string) *GetRemoteClassCourseResponseBodyResultTeachingParticipant {
	s.OrgName = &v
	return s
}

type GetRemoteClassCourseResponseBodyResultAttendParticipants struct {
	// 参与方ID
	ParticipantId *string `json:"participantId,omitempty" xml:"participantId,omitempty"`
	// 参与方名称
	ParticipantName *string `json:"participantName,omitempty" xml:"participantName,omitempty"`
	// 组织ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 组织名称
	OrgName *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
}

func (s GetRemoteClassCourseResponseBodyResultAttendParticipants) String() string {
	return tea.Prettify(s)
}

func (s GetRemoteClassCourseResponseBodyResultAttendParticipants) GoString() string {
	return s.String()
}

func (s *GetRemoteClassCourseResponseBodyResultAttendParticipants) SetParticipantId(v string) *GetRemoteClassCourseResponseBodyResultAttendParticipants {
	s.ParticipantId = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResultAttendParticipants) SetParticipantName(v string) *GetRemoteClassCourseResponseBodyResultAttendParticipants {
	s.ParticipantName = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResultAttendParticipants) SetCorpId(v string) *GetRemoteClassCourseResponseBodyResultAttendParticipants {
	s.CorpId = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResultAttendParticipants) SetOrgName(v string) *GetRemoteClassCourseResponseBodyResultAttendParticipants {
	s.OrgName = &v
	return s
}

type GetRemoteClassCourseResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRemoteClassCourseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRemoteClassCourseResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRemoteClassCourseResponse) GoString() string {
	return s.String()
}

func (s *GetRemoteClassCourseResponse) SetHeaders(v map[string]*string) *GetRemoteClassCourseResponse {
	s.Headers = v
	return s
}

func (s *GetRemoteClassCourseResponse) SetBody(v *GetRemoteClassCourseResponseBody) *GetRemoteClassCourseResponse {
	s.Body = v
	return s
}

type QueryUniversityCourseGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUniversityCourseGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUniversityCourseGroupHeaders) GoString() string {
	return s.String()
}

func (s *QueryUniversityCourseGroupHeaders) SetCommonHeaders(v map[string]*string) *QueryUniversityCourseGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUniversityCourseGroupHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUniversityCourseGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUniversityCourseGroupRequest struct {
	// 操作人
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 课程编码
	CourseGroupCode *string `json:"courseGroupCode,omitempty" xml:"courseGroupCode,omitempty"`
}

func (s QueryUniversityCourseGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUniversityCourseGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryUniversityCourseGroupRequest) SetOpUserId(v string) *QueryUniversityCourseGroupRequest {
	s.OpUserId = &v
	return s
}

func (s *QueryUniversityCourseGroupRequest) SetCourseGroupCode(v string) *QueryUniversityCourseGroupRequest {
	s.CourseGroupCode = &v
	return s
}

type QueryUniversityCourseGroupResponseBody struct {
	// 课程组信息
	UniversityCourseGroupInfo *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo `json:"universityCourseGroupInfo,omitempty" xml:"universityCourseGroupInfo,omitempty" type:"Struct"`
}

func (s QueryUniversityCourseGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUniversityCourseGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUniversityCourseGroupResponseBody) SetUniversityCourseGroupInfo(v *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo) *QueryUniversityCourseGroupResponseBody {
	s.UniversityCourseGroupInfo = v
	return s
}

type QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo struct {
	// 合作方课程组code
	IsvCourseGroupCode *string `json:"isvCourseGroupCode,omitempty" xml:"isvCourseGroupCode,omitempty"`
	// 课程组编码
	CourseGroupCode *string `json:"courseGroupCode,omitempty" xml:"courseGroupCode,omitempty"`
	// 课程组名称
	CourseGroupName *string `json:"courseGroupName,omitempty" xml:"courseGroupName,omitempty"`
	// 课程组介绍
	CourseGroupIntroduce *string `json:"courseGroupIntroduce,omitempty" xml:"courseGroupIntroduce,omitempty"`
	// 学年
	SchoolYear *string `json:"schoolYear,omitempty" xml:"schoolYear,omitempty"`
	// 学期
	Semester *int32 `json:"semester,omitempty" xml:"semester,omitempty"`
	// 学段编码
	PeriodCode *string `json:"periodCode,omitempty" xml:"periodCode,omitempty"`
	// 学科名称
	SubjectName *string `json:"subjectName,omitempty" xml:"subjectName,omitempty"`
	// 课程组详细
	CourserGroupItemModels []*QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels `json:"courserGroupItemModels,omitempty" xml:"courserGroupItemModels,omitempty" type:"Repeated"`
}

func (s QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo) GoString() string {
	return s.String()
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo) SetIsvCourseGroupCode(v string) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo {
	s.IsvCourseGroupCode = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo) SetCourseGroupCode(v string) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo {
	s.CourseGroupCode = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo) SetCourseGroupName(v string) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo {
	s.CourseGroupName = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo) SetCourseGroupIntroduce(v string) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo {
	s.CourseGroupIntroduce = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo) SetSchoolYear(v string) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo {
	s.SchoolYear = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo) SetSemester(v int32) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo {
	s.Semester = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo) SetPeriodCode(v string) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo {
	s.PeriodCode = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo) SetSubjectName(v string) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo {
	s.SubjectName = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo) SetCourserGroupItemModels(v []*QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo {
	s.CourserGroupItemModels = v
	return s
}

type QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels struct {
	// 教室主键
	ClassroomId *int64 `json:"classroomId,omitempty" xml:"classroomId,omitempty"`
	// 上课周期
	ClassPeriodType *int32 `json:"classPeriodType,omitempty" xml:"classPeriodType,omitempty"`
	// 一周的第几天
	DayOfWeek *int32 `json:"dayOfWeek,omitempty" xml:"dayOfWeek,omitempty"`
	// 课节
	SectionIndex []*int32 `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty" type:"Repeated"`
	// 开始时间
	CourserGroupItemStartDate *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemStartDate `json:"courserGroupItemStartDate,omitempty" xml:"courserGroupItemStartDate,omitempty" type:"Struct"`
	// 结束时间
	CourserGroupItemEndDate *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemEndDate `json:"courserGroupItemEndDate,omitempty" xml:"courserGroupItemEndDate,omitempty" type:"Struct"`
	// 课程类型
	CourseType *int32 `json:"courseType,omitempty" xml:"courseType,omitempty"`
}

func (s QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels) String() string {
	return tea.Prettify(s)
}

func (s QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels) GoString() string {
	return s.String()
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels) SetClassroomId(v int64) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels {
	s.ClassroomId = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels) SetClassPeriodType(v int32) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels {
	s.ClassPeriodType = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels) SetDayOfWeek(v int32) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels {
	s.DayOfWeek = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels) SetSectionIndex(v []*int32) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels {
	s.SectionIndex = v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels) SetCourserGroupItemStartDate(v *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemStartDate) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels {
	s.CourserGroupItemStartDate = v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels) SetCourserGroupItemEndDate(v *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemEndDate) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels {
	s.CourserGroupItemEndDate = v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels) SetCourseType(v int32) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels {
	s.CourseType = &v
	return s
}

type QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemStartDate struct {
	// 年
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
	// 月
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// 日
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
}

func (s QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemStartDate) String() string {
	return tea.Prettify(s)
}

func (s QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemStartDate) GoString() string {
	return s.String()
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemStartDate) SetYear(v int32) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemStartDate {
	s.Year = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemStartDate) SetMonth(v int32) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemStartDate {
	s.Month = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemStartDate) SetDayOfMonth(v int32) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemStartDate {
	s.DayOfMonth = &v
	return s
}

type QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemEndDate struct {
	// 年
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
	// 月
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// 日
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
}

func (s QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemEndDate) String() string {
	return tea.Prettify(s)
}

func (s QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemEndDate) GoString() string {
	return s.String()
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemEndDate) SetYear(v int32) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemEndDate {
	s.Year = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemEndDate) SetMonth(v int32) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemEndDate {
	s.Month = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemEndDate) SetDayOfMonth(v int32) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemEndDate {
	s.DayOfMonth = &v
	return s
}

type QueryUniversityCourseGroupResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUniversityCourseGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUniversityCourseGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUniversityCourseGroupResponse) GoString() string {
	return s.String()
}

func (s *QueryUniversityCourseGroupResponse) SetHeaders(v map[string]*string) *QueryUniversityCourseGroupResponse {
	s.Headers = v
	return s
}

func (s *QueryUniversityCourseGroupResponse) SetBody(v *QueryUniversityCourseGroupResponseBody) *QueryUniversityCourseGroupResponse {
	s.Body = v
	return s
}

type QueryAllSubjectsFromClassScheduleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllSubjectsFromClassScheduleHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllSubjectsFromClassScheduleHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllSubjectsFromClassScheduleHeaders) SetCommonHeaders(v map[string]*string) *QueryAllSubjectsFromClassScheduleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllSubjectsFromClassScheduleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllSubjectsFromClassScheduleRequest struct {
	// 班级ID。
	ClassIds []*int64 `json:"classIds,omitempty" xml:"classIds,omitempty" type:"Repeated"`
	// 操作者的userid。
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 学段编码：KINDERGARTEN：小学 PRIMARY_SCHOOL：小学 MODDLE_SCHOOL： 初中 HIGH_SCHOOL： 高中 UNIVERSITY：大学 NOT_SCHOOL：无学段
	PeriodCode *string `json:"periodCode,omitempty" xml:"periodCode,omitempty"`
}

func (s QueryAllSubjectsFromClassScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllSubjectsFromClassScheduleRequest) GoString() string {
	return s.String()
}

func (s *QueryAllSubjectsFromClassScheduleRequest) SetClassIds(v []*int64) *QueryAllSubjectsFromClassScheduleRequest {
	s.ClassIds = v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleRequest) SetOpUserId(v string) *QueryAllSubjectsFromClassScheduleRequest {
	s.OpUserId = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleRequest) SetPeriodCode(v string) *QueryAllSubjectsFromClassScheduleRequest {
	s.PeriodCode = &v
	return s
}

type QueryAllSubjectsFromClassScheduleShrinkRequest struct {
	// 班级ID。
	ClassIdsShrink *string `json:"classIds,omitempty" xml:"classIds,omitempty"`
	// 操作者的userid。
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 学段编码：KINDERGARTEN：小学 PRIMARY_SCHOOL：小学 MODDLE_SCHOOL： 初中 HIGH_SCHOOL： 高中 UNIVERSITY：大学 NOT_SCHOOL：无学段
	PeriodCode *string `json:"periodCode,omitempty" xml:"periodCode,omitempty"`
}

func (s QueryAllSubjectsFromClassScheduleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllSubjectsFromClassScheduleShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryAllSubjectsFromClassScheduleShrinkRequest) SetClassIdsShrink(v string) *QueryAllSubjectsFromClassScheduleShrinkRequest {
	s.ClassIdsShrink = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleShrinkRequest) SetOpUserId(v string) *QueryAllSubjectsFromClassScheduleShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleShrinkRequest) SetPeriodCode(v string) *QueryAllSubjectsFromClassScheduleShrinkRequest {
	s.PeriodCode = &v
	return s
}

type QueryAllSubjectsFromClassScheduleResponseBody struct {
	// 查询结果。
	Result []*QueryAllSubjectsFromClassScheduleResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryAllSubjectsFromClassScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllSubjectsFromClassScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllSubjectsFromClassScheduleResponseBody) SetResult(v []*QueryAllSubjectsFromClassScheduleResponseBodyResult) *QueryAllSubjectsFromClassScheduleResponseBody {
	s.Result = v
	return s
}

type QueryAllSubjectsFromClassScheduleResponseBodyResult struct {
	// 学科名称。
	SubjectName *string `json:"subjectName,omitempty" xml:"subjectName,omitempty"`
	// 学科code。
	SubjectCode *string `json:"subjectCode,omitempty" xml:"subjectCode,omitempty"`
	// 学段编码：  KINDERGARTEN：小学 PRIMARY_SCHOOL：小学 MODDLE_SCHOOL： 初中 HIGH_SCHOOL： 高中 UNIVERSITY：大学 NOT_SCHOOL：无学段
	PeriodCode *string `json:"periodCode,omitempty" xml:"periodCode,omitempty"`
	// creatorOrgId
	CreatorOrgId *int64 `json:"creatorOrgId,omitempty" xml:"creatorOrgId,omitempty"`
	// 拓展信息
	Ext *QueryAllSubjectsFromClassScheduleResponseBodyResultExt `json:"ext,omitempty" xml:"ext,omitempty" type:"Struct"`
}

func (s QueryAllSubjectsFromClassScheduleResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryAllSubjectsFromClassScheduleResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResult) SetSubjectName(v string) *QueryAllSubjectsFromClassScheduleResponseBodyResult {
	s.SubjectName = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResult) SetSubjectCode(v string) *QueryAllSubjectsFromClassScheduleResponseBodyResult {
	s.SubjectCode = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResult) SetPeriodCode(v string) *QueryAllSubjectsFromClassScheduleResponseBodyResult {
	s.PeriodCode = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResult) SetCreatorOrgId(v int64) *QueryAllSubjectsFromClassScheduleResponseBodyResult {
	s.CreatorOrgId = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResult) SetExt(v *QueryAllSubjectsFromClassScheduleResponseBodyResultExt) *QueryAllSubjectsFromClassScheduleResponseBodyResult {
	s.Ext = v
	return s
}

type QueryAllSubjectsFromClassScheduleResponseBodyResultExt struct {
	// 学科字体颜色
	FontColor *string `json:"fontColor,omitempty" xml:"fontColor,omitempty"`
	// 学科背景颜色
	BackgroundColor *string `json:"backgroundColor,omitempty" xml:"backgroundColor,omitempty"`
	// 班级id。
	ClassId *int64 `json:"classId,omitempty" xml:"classId,omitempty"`
	// 老师列表
	TeacherList []*QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList `json:"teacherList,omitempty" xml:"teacherList,omitempty" type:"Repeated"`
}

func (s QueryAllSubjectsFromClassScheduleResponseBodyResultExt) String() string {
	return tea.Prettify(s)
}

func (s QueryAllSubjectsFromClassScheduleResponseBodyResultExt) GoString() string {
	return s.String()
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResultExt) SetFontColor(v string) *QueryAllSubjectsFromClassScheduleResponseBodyResultExt {
	s.FontColor = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResultExt) SetBackgroundColor(v string) *QueryAllSubjectsFromClassScheduleResponseBodyResultExt {
	s.BackgroundColor = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResultExt) SetClassId(v int64) *QueryAllSubjectsFromClassScheduleResponseBodyResultExt {
	s.ClassId = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResultExt) SetTeacherList(v []*QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList) *QueryAllSubjectsFromClassScheduleResponseBodyResultExt {
	s.TeacherList = v
	return s
}

type QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList struct {
	// 老师的userid。
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 老师名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 老师的头像url
	Avator *string `json:"avator,omitempty" xml:"avator,omitempty"`
	// 老师的uid。
	Uid *int64 `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList) String() string {
	return tea.Prettify(s)
}

func (s QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList) GoString() string {
	return s.String()
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList) SetUserId(v string) *QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList {
	s.UserId = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList) SetName(v string) *QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList {
	s.Name = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList) SetAvator(v string) *QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList {
	s.Avator = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList) SetUid(v int64) *QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList {
	s.Uid = &v
	return s
}

type QueryAllSubjectsFromClassScheduleResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAllSubjectsFromClassScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllSubjectsFromClassScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllSubjectsFromClassScheduleResponse) GoString() string {
	return s.String()
}

func (s *QueryAllSubjectsFromClassScheduleResponse) SetHeaders(v map[string]*string) *QueryAllSubjectsFromClassScheduleResponse {
	s.Headers = v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleResponse) SetBody(v *QueryAllSubjectsFromClassScheduleResponseBody) *QueryAllSubjectsFromClassScheduleResponse {
	s.Body = v
	return s
}

type CourseSchedulingComplimentNoticeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CourseSchedulingComplimentNoticeHeaders) String() string {
	return tea.Prettify(s)
}

func (s CourseSchedulingComplimentNoticeHeaders) GoString() string {
	return s.String()
}

func (s *CourseSchedulingComplimentNoticeHeaders) SetCommonHeaders(v map[string]*string) *CourseSchedulingComplimentNoticeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CourseSchedulingComplimentNoticeHeaders) SetXAcsDingtalkAccessToken(v string) *CourseSchedulingComplimentNoticeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CourseSchedulingComplimentNoticeRequest struct {
	// opUserId
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s CourseSchedulingComplimentNoticeRequest) String() string {
	return tea.Prettify(s)
}

func (s CourseSchedulingComplimentNoticeRequest) GoString() string {
	return s.String()
}

func (s *CourseSchedulingComplimentNoticeRequest) SetOpUserId(v string) *CourseSchedulingComplimentNoticeRequest {
	s.OpUserId = &v
	return s
}

type CourseSchedulingComplimentNoticeResponseBody struct {
	// 通知课程导入完成是否成功。
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CourseSchedulingComplimentNoticeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CourseSchedulingComplimentNoticeResponseBody) GoString() string {
	return s.String()
}

func (s *CourseSchedulingComplimentNoticeResponseBody) SetResult(v bool) *CourseSchedulingComplimentNoticeResponseBody {
	s.Result = &v
	return s
}

type CourseSchedulingComplimentNoticeResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CourseSchedulingComplimentNoticeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CourseSchedulingComplimentNoticeResponse) String() string {
	return tea.Prettify(s)
}

func (s CourseSchedulingComplimentNoticeResponse) GoString() string {
	return s.String()
}

func (s *CourseSchedulingComplimentNoticeResponse) SetHeaders(v map[string]*string) *CourseSchedulingComplimentNoticeResponse {
	s.Headers = v
	return s
}

func (s *CourseSchedulingComplimentNoticeResponse) SetBody(v *CourseSchedulingComplimentNoticeResponseBody) *CourseSchedulingComplimentNoticeResponse {
	s.Body = v
	return s
}

type DeleteUniversityCourseGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteUniversityCourseGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteUniversityCourseGroupHeaders) GoString() string {
	return s.String()
}

func (s *DeleteUniversityCourseGroupHeaders) SetCommonHeaders(v map[string]*string) *DeleteUniversityCourseGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteUniversityCourseGroupHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteUniversityCourseGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteUniversityCourseGroupRequest struct {
	// 操作人
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 课程组编码
	CourseGroupCode *string `json:"courseGroupCode,omitempty" xml:"courseGroupCode,omitempty"`
}

func (s DeleteUniversityCourseGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUniversityCourseGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteUniversityCourseGroupRequest) SetOpUserId(v string) *DeleteUniversityCourseGroupRequest {
	s.OpUserId = &v
	return s
}

func (s *DeleteUniversityCourseGroupRequest) SetCourseGroupCode(v string) *DeleteUniversityCourseGroupRequest {
	s.CourseGroupCode = &v
	return s
}

type DeleteUniversityCourseGroupResponseBody struct {
	// 操作结果
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteUniversityCourseGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUniversityCourseGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUniversityCourseGroupResponseBody) SetResult(v bool) *DeleteUniversityCourseGroupResponseBody {
	s.Result = &v
	return s
}

type DeleteUniversityCourseGroupResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUniversityCourseGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUniversityCourseGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUniversityCourseGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteUniversityCourseGroupResponse) SetHeaders(v map[string]*string) *DeleteUniversityCourseGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteUniversityCourseGroupResponse) SetBody(v *DeleteUniversityCourseGroupResponseBody) *DeleteUniversityCourseGroupResponse {
	s.Body = v
	return s
}

type DeleteDeptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteDeptHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeptHeaders) GoString() string {
	return s.String()
}

func (s *DeleteDeptHeaders) SetCommonHeaders(v map[string]*string) *DeleteDeptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteDeptHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteDeptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteDeptRequest struct {
	// 钉钉企业管理员员工ID
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
}

func (s DeleteDeptRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeptRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeptRequest) SetOperator(v string) *DeleteDeptRequest {
	s.Operator = &v
	return s
}

type DeleteDeptResponseBody struct {
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteDeptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeptResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeptResponseBody) SetSuccess(v bool) *DeleteDeptResponseBody {
	s.Success = &v
	return s
}

type DeleteDeptResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDeptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDeptResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeptResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeptResponse) SetHeaders(v map[string]*string) *DeleteDeptResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeptResponse) SetBody(v *DeleteDeptResponseBody) *DeleteDeptResponse {
	s.Body = v
	return s
}

type DeleteGuardianHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteGuardianHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteGuardianHeaders) GoString() string {
	return s.String()
}

func (s *DeleteGuardianHeaders) SetCommonHeaders(v map[string]*string) *DeleteGuardianHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteGuardianHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteGuardianHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteGuardianRequest struct {
	// 学生ID
	StuId *string `json:"stuId,omitempty" xml:"stuId,omitempty"`
	// 钉钉企业管理员员工ID
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
}

func (s DeleteGuardianRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGuardianRequest) GoString() string {
	return s.String()
}

func (s *DeleteGuardianRequest) SetStuId(v string) *DeleteGuardianRequest {
	s.StuId = &v
	return s
}

func (s *DeleteGuardianRequest) SetOperator(v string) *DeleteGuardianRequest {
	s.Operator = &v
	return s
}

type DeleteGuardianResponseBody struct {
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteGuardianResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGuardianResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGuardianResponseBody) SetSuccess(v bool) *DeleteGuardianResponseBody {
	s.Success = &v
	return s
}

type DeleteGuardianResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteGuardianResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGuardianResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGuardianResponse) GoString() string {
	return s.String()
}

func (s *DeleteGuardianResponse) SetHeaders(v map[string]*string) *DeleteGuardianResponse {
	s.Headers = v
	return s
}

func (s *DeleteGuardianResponse) SetBody(v *DeleteGuardianResponseBody) *DeleteGuardianResponse {
	s.Body = v
	return s
}

type InsertSectionConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InsertSectionConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s InsertSectionConfigHeaders) GoString() string {
	return s.String()
}

func (s *InsertSectionConfigHeaders) SetCommonHeaders(v map[string]*string) *InsertSectionConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsertSectionConfigHeaders) SetXAcsDingtalkAccessToken(v string) *InsertSectionConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InsertSectionConfigRequest struct {
	// 节次模型
	SectionModels []*InsertSectionConfigRequestSectionModels `json:"sectionModels,omitempty" xml:"sectionModels,omitempty" type:"Repeated"`
	// 开始日期
	Start *InsertSectionConfigRequestStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
	// 结束日期
	End *InsertSectionConfigRequestEnd `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
	// 课程表名称
	ScheduleName *string `json:"scheduleName,omitempty" xml:"scheduleName,omitempty"`
	// 操作人的userid。
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s InsertSectionConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertSectionConfigRequest) GoString() string {
	return s.String()
}

func (s *InsertSectionConfigRequest) SetSectionModels(v []*InsertSectionConfigRequestSectionModels) *InsertSectionConfigRequest {
	s.SectionModels = v
	return s
}

func (s *InsertSectionConfigRequest) SetStart(v *InsertSectionConfigRequestStart) *InsertSectionConfigRequest {
	s.Start = v
	return s
}

func (s *InsertSectionConfigRequest) SetEnd(v *InsertSectionConfigRequestEnd) *InsertSectionConfigRequest {
	s.End = v
	return s
}

func (s *InsertSectionConfigRequest) SetScheduleName(v string) *InsertSectionConfigRequest {
	s.ScheduleName = &v
	return s
}

func (s *InsertSectionConfigRequest) SetOpUserId(v string) *InsertSectionConfigRequest {
	s.OpUserId = &v
	return s
}

type InsertSectionConfigRequestSectionModels struct {
	// 节次类型
	SectionType *string `json:"sectionType,omitempty" xml:"sectionType,omitempty"`
	// 开始时间
	Start *InsertSectionConfigRequestSectionModelsStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
	// 节次序号
	SectionIndex *int32 `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty"`
	// 结束时间
	End *InsertSectionConfigRequestSectionModelsEnd `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
	// 节次名称
	SectionName *string `json:"sectionName,omitempty" xml:"sectionName,omitempty"`
}

func (s InsertSectionConfigRequestSectionModels) String() string {
	return tea.Prettify(s)
}

func (s InsertSectionConfigRequestSectionModels) GoString() string {
	return s.String()
}

func (s *InsertSectionConfigRequestSectionModels) SetSectionType(v string) *InsertSectionConfigRequestSectionModels {
	s.SectionType = &v
	return s
}

func (s *InsertSectionConfigRequestSectionModels) SetStart(v *InsertSectionConfigRequestSectionModelsStart) *InsertSectionConfigRequestSectionModels {
	s.Start = v
	return s
}

func (s *InsertSectionConfigRequestSectionModels) SetSectionIndex(v int32) *InsertSectionConfigRequestSectionModels {
	s.SectionIndex = &v
	return s
}

func (s *InsertSectionConfigRequestSectionModels) SetEnd(v *InsertSectionConfigRequestSectionModelsEnd) *InsertSectionConfigRequestSectionModels {
	s.End = v
	return s
}

func (s *InsertSectionConfigRequestSectionModels) SetSectionName(v string) *InsertSectionConfigRequestSectionModels {
	s.SectionName = &v
	return s
}

type InsertSectionConfigRequestSectionModelsStart struct {
	// 分钟
	Min *int32 `json:"min,omitempty" xml:"min,omitempty"`
	// 小时
	Hour *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
}

func (s InsertSectionConfigRequestSectionModelsStart) String() string {
	return tea.Prettify(s)
}

func (s InsertSectionConfigRequestSectionModelsStart) GoString() string {
	return s.String()
}

func (s *InsertSectionConfigRequestSectionModelsStart) SetMin(v int32) *InsertSectionConfigRequestSectionModelsStart {
	s.Min = &v
	return s
}

func (s *InsertSectionConfigRequestSectionModelsStart) SetHour(v int32) *InsertSectionConfigRequestSectionModelsStart {
	s.Hour = &v
	return s
}

type InsertSectionConfigRequestSectionModelsEnd struct {
	// 分钟
	Min *int32 `json:"min,omitempty" xml:"min,omitempty"`
	// 小时
	Hour *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
}

func (s InsertSectionConfigRequestSectionModelsEnd) String() string {
	return tea.Prettify(s)
}

func (s InsertSectionConfigRequestSectionModelsEnd) GoString() string {
	return s.String()
}

func (s *InsertSectionConfigRequestSectionModelsEnd) SetMin(v int32) *InsertSectionConfigRequestSectionModelsEnd {
	s.Min = &v
	return s
}

func (s *InsertSectionConfigRequestSectionModelsEnd) SetHour(v int32) *InsertSectionConfigRequestSectionModelsEnd {
	s.Hour = &v
	return s
}

type InsertSectionConfigRequestStart struct {
	// 月份
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// 年份
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
	// 一月中的第几天
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
}

func (s InsertSectionConfigRequestStart) String() string {
	return tea.Prettify(s)
}

func (s InsertSectionConfigRequestStart) GoString() string {
	return s.String()
}

func (s *InsertSectionConfigRequestStart) SetMonth(v int32) *InsertSectionConfigRequestStart {
	s.Month = &v
	return s
}

func (s *InsertSectionConfigRequestStart) SetYear(v int32) *InsertSectionConfigRequestStart {
	s.Year = &v
	return s
}

func (s *InsertSectionConfigRequestStart) SetDayOfMonth(v int32) *InsertSectionConfigRequestStart {
	s.DayOfMonth = &v
	return s
}

type InsertSectionConfigRequestEnd struct {
	// 月份
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// 年份
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
	// 一月中的第几天
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
}

func (s InsertSectionConfigRequestEnd) String() string {
	return tea.Prettify(s)
}

func (s InsertSectionConfigRequestEnd) GoString() string {
	return s.String()
}

func (s *InsertSectionConfigRequestEnd) SetMonth(v int32) *InsertSectionConfigRequestEnd {
	s.Month = &v
	return s
}

func (s *InsertSectionConfigRequestEnd) SetYear(v int32) *InsertSectionConfigRequestEnd {
	s.Year = &v
	return s
}

func (s *InsertSectionConfigRequestEnd) SetDayOfMonth(v int32) *InsertSectionConfigRequestEnd {
	s.DayOfMonth = &v
	return s
}

type InsertSectionConfigResponseBody struct {
	// 初始化是否成功。
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s InsertSectionConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertSectionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *InsertSectionConfigResponseBody) SetResult(v bool) *InsertSectionConfigResponseBody {
	s.Result = &v
	return s
}

type InsertSectionConfigResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertSectionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertSectionConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertSectionConfigResponse) GoString() string {
	return s.String()
}

func (s *InsertSectionConfigResponse) SetHeaders(v map[string]*string) *InsertSectionConfigResponse {
	s.Headers = v
	return s
}

func (s *InsertSectionConfigResponse) SetBody(v *InsertSectionConfigResponseBody) *InsertSectionConfigResponse {
	s.Body = v
	return s
}

type QueryDeviceListByCorpIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDeviceListByCorpIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceListByCorpIdHeaders) GoString() string {
	return s.String()
}

func (s *QueryDeviceListByCorpIdHeaders) SetCommonHeaders(v map[string]*string) *QueryDeviceListByCorpIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDeviceListByCorpIdHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDeviceListByCorpIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDeviceListByCorpIdRequest struct {
	Operator   *string `json:"operator,omitempty" xml:"operator,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryDeviceListByCorpIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceListByCorpIdRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceListByCorpIdRequest) SetOperator(v string) *QueryDeviceListByCorpIdRequest {
	s.Operator = &v
	return s
}

func (s *QueryDeviceListByCorpIdRequest) SetPageNumber(v int32) *QueryDeviceListByCorpIdRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryDeviceListByCorpIdRequest) SetPageSize(v int32) *QueryDeviceListByCorpIdRequest {
	s.PageSize = &v
	return s
}

type QueryDeviceListByCorpIdResponseBody struct {
	// Id of the request
	Success *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
	Result  *QueryDeviceListByCorpIdResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryDeviceListByCorpIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceListByCorpIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceListByCorpIdResponseBody) SetSuccess(v bool) *QueryDeviceListByCorpIdResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDeviceListByCorpIdResponseBody) SetResult(v *QueryDeviceListByCorpIdResponseBodyResult) *QueryDeviceListByCorpIdResponseBody {
	s.Result = v
	return s
}

type QueryDeviceListByCorpIdResponseBodyResult struct {
	HasMore *bool                                            `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List    []*QueryDeviceListByCorpIdResponseBodyResultList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s QueryDeviceListByCorpIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceListByCorpIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryDeviceListByCorpIdResponseBodyResult) SetHasMore(v bool) *QueryDeviceListByCorpIdResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *QueryDeviceListByCorpIdResponseBodyResult) SetList(v []*QueryDeviceListByCorpIdResponseBodyResultList) *QueryDeviceListByCorpIdResponseBodyResult {
	s.List = v
	return s
}

type QueryDeviceListByCorpIdResponseBodyResultList struct {
	AppStatus  *int32  `json:"appStatus,omitempty" xml:"appStatus,omitempty"`
	DeviceCode *string `json:"deviceCode,omitempty" xml:"deviceCode,omitempty"`
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
}

func (s QueryDeviceListByCorpIdResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceListByCorpIdResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *QueryDeviceListByCorpIdResponseBodyResultList) SetAppStatus(v int32) *QueryDeviceListByCorpIdResponseBodyResultList {
	s.AppStatus = &v
	return s
}

func (s *QueryDeviceListByCorpIdResponseBodyResultList) SetDeviceCode(v string) *QueryDeviceListByCorpIdResponseBodyResultList {
	s.DeviceCode = &v
	return s
}

func (s *QueryDeviceListByCorpIdResponseBodyResultList) SetDeviceName(v string) *QueryDeviceListByCorpIdResponseBodyResultList {
	s.DeviceName = &v
	return s
}

type QueryDeviceListByCorpIdResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDeviceListByCorpIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDeviceListByCorpIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceListByCorpIdResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceListByCorpIdResponse) SetHeaders(v map[string]*string) *QueryDeviceListByCorpIdResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceListByCorpIdResponse) SetBody(v *QueryDeviceListByCorpIdResponseBody) *QueryDeviceListByCorpIdResponse {
	s.Body = v
	return s
}

type UpdateUniversityCourseGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateUniversityCourseGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateUniversityCourseGroupHeaders) GoString() string {
	return s.String()
}

func (s *UpdateUniversityCourseGroupHeaders) SetCommonHeaders(v map[string]*string) *UpdateUniversityCourseGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateUniversityCourseGroupHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateUniversityCourseGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateUniversityCourseGroupRequest struct {
	// opUserId
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 扩展信息
	Ext *string `json:"ext,omitempty" xml:"ext,omitempty"`
	// 课程组编码
	CourseGroupCode *string `json:"courseGroupCode,omitempty" xml:"courseGroupCode,omitempty"`
	// 课程组介绍
	CourseGroupIntroduce *string `json:"courseGroupIntroduce,omitempty" xml:"courseGroupIntroduce,omitempty"`
	// 课程组名称
	CourseGroupName *string `json:"courseGroupName,omitempty" xml:"courseGroupName,omitempty"`
	// 课程组详细
	CourserGroupItemModels []*UpdateUniversityCourseGroupRequestCourserGroupItemModels `json:"courserGroupItemModels,omitempty" xml:"courserGroupItemModels,omitempty" type:"Repeated"`
}

func (s UpdateUniversityCourseGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUniversityCourseGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateUniversityCourseGroupRequest) SetOpUserId(v string) *UpdateUniversityCourseGroupRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateUniversityCourseGroupRequest) SetExt(v string) *UpdateUniversityCourseGroupRequest {
	s.Ext = &v
	return s
}

func (s *UpdateUniversityCourseGroupRequest) SetCourseGroupCode(v string) *UpdateUniversityCourseGroupRequest {
	s.CourseGroupCode = &v
	return s
}

func (s *UpdateUniversityCourseGroupRequest) SetCourseGroupIntroduce(v string) *UpdateUniversityCourseGroupRequest {
	s.CourseGroupIntroduce = &v
	return s
}

func (s *UpdateUniversityCourseGroupRequest) SetCourseGroupName(v string) *UpdateUniversityCourseGroupRequest {
	s.CourseGroupName = &v
	return s
}

func (s *UpdateUniversityCourseGroupRequest) SetCourserGroupItemModels(v []*UpdateUniversityCourseGroupRequestCourserGroupItemModels) *UpdateUniversityCourseGroupRequest {
	s.CourserGroupItemModels = v
	return s
}

type UpdateUniversityCourseGroupRequestCourserGroupItemModels struct {
	// 一周的第几天
	DayOfWeek *int32 `json:"dayOfWeek,omitempty" xml:"dayOfWeek,omitempty"`
	// 上课周期
	ClassPeriodType *int32 `json:"classPeriodType,omitempty" xml:"classPeriodType,omitempty"`
	// 开始时间
	CourserGroupItemStartDate *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate `json:"courserGroupItemStartDate,omitempty" xml:"courserGroupItemStartDate,omitempty" type:"Struct"`
	// 课节
	SectionIndex []*int32 `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty" type:"Repeated"`
	// 结束时间
	CourserGroupItemEndDate *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate `json:"courserGroupItemEndDate,omitempty" xml:"courserGroupItemEndDate,omitempty" type:"Struct"`
	// 课程类型
	CourseType *int32 `json:"courseType,omitempty" xml:"courseType,omitempty"`
	// classroomId
	ClassroomId *int64 `json:"classroomId,omitempty" xml:"classroomId,omitempty"`
}

func (s UpdateUniversityCourseGroupRequestCourserGroupItemModels) String() string {
	return tea.Prettify(s)
}

func (s UpdateUniversityCourseGroupRequestCourserGroupItemModels) GoString() string {
	return s.String()
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModels) SetDayOfWeek(v int32) *UpdateUniversityCourseGroupRequestCourserGroupItemModels {
	s.DayOfWeek = &v
	return s
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModels) SetClassPeriodType(v int32) *UpdateUniversityCourseGroupRequestCourserGroupItemModels {
	s.ClassPeriodType = &v
	return s
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModels) SetCourserGroupItemStartDate(v *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate) *UpdateUniversityCourseGroupRequestCourserGroupItemModels {
	s.CourserGroupItemStartDate = v
	return s
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModels) SetSectionIndex(v []*int32) *UpdateUniversityCourseGroupRequestCourserGroupItemModels {
	s.SectionIndex = v
	return s
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModels) SetCourserGroupItemEndDate(v *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate) *UpdateUniversityCourseGroupRequestCourserGroupItemModels {
	s.CourserGroupItemEndDate = v
	return s
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModels) SetCourseType(v int32) *UpdateUniversityCourseGroupRequestCourserGroupItemModels {
	s.CourseType = &v
	return s
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModels) SetClassroomId(v int64) *UpdateUniversityCourseGroupRequestCourserGroupItemModels {
	s.ClassroomId = &v
	return s
}

type UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate struct {
	// 月
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// 年
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
	// 一月的第几天
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
}

func (s UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate) String() string {
	return tea.Prettify(s)
}

func (s UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate) GoString() string {
	return s.String()
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate) SetMonth(v int32) *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate {
	s.Month = &v
	return s
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate) SetYear(v int32) *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate {
	s.Year = &v
	return s
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate) SetDayOfMonth(v int32) *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate {
	s.DayOfMonth = &v
	return s
}

type UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate struct {
	// 月
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// 年
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
	// 一月的第几天
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
}

func (s UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate) String() string {
	return tea.Prettify(s)
}

func (s UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate) GoString() string {
	return s.String()
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate) SetMonth(v int32) *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate {
	s.Month = &v
	return s
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate) SetYear(v int32) *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate {
	s.Year = &v
	return s
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate) SetDayOfMonth(v int32) *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate {
	s.DayOfMonth = &v
	return s
}

type UpdateUniversityCourseGroupResponseBody struct {
	// 是否成功
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateUniversityCourseGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUniversityCourseGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUniversityCourseGroupResponseBody) SetResult(v bool) *UpdateUniversityCourseGroupResponseBody {
	s.Result = &v
	return s
}

type UpdateUniversityCourseGroupResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateUniversityCourseGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateUniversityCourseGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUniversityCourseGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateUniversityCourseGroupResponse) SetHeaders(v map[string]*string) *UpdateUniversityCourseGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateUniversityCourseGroupResponse) SetBody(v *UpdateUniversityCourseGroupResponseBody) *UpdateUniversityCourseGroupResponse {
	s.Body = v
	return s
}

type BatchOrgCreateHWHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchOrgCreateHWHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchOrgCreateHWHeaders) GoString() string {
	return s.String()
}

func (s *BatchOrgCreateHWHeaders) SetCommonHeaders(v map[string]*string) *BatchOrgCreateHWHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchOrgCreateHWHeaders) SetXAcsDingtalkAccessToken(v string) *BatchOrgCreateHWHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchOrgCreateHWRequest struct {
	// 作业视频
	HwMedia *string `json:"hwMedia,omitempty" xml:"hwMedia,omitempty"`
	// 作业内容
	HwContent *string `json:"hwContent,omitempty" xml:"hwContent,omitempty"`
	// 作业标题
	HwTitle *string `json:"hwTitle,omitempty" xml:"hwTitle,omitempty"`
	// 作业课程名称
	CourseName *string `json:"courseName,omitempty" xml:"courseName,omitempty"`
	// 作业图片
	HwPhoto *string `json:"hwPhoto,omitempty" xml:"hwPhoto,omitempty"`
	// 作业音视频
	HwVideo *string `json:"hwVideo,omitempty" xml:"hwVideo,omitempty"`
	// 老师名称
	TeacherName *string `json:"teacherName,omitempty" xml:"teacherName,omitempty"`
	// 老师userid
	TeacherUserId *string `json:"teacherUserId,omitempty" xml:"teacherUserId,omitempty"`
	// 幂等ID字段
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// 扩展属性
	Attributes *string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 发送对象
	TargetRole *string `json:"targetRole,omitempty" xml:"targetRole,omitempty"`
	// 业务编码
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	// 状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 定时调度
	ScheduledRelease *string `json:"scheduledRelease,omitempty" xml:"scheduledRelease,omitempty"`
	// 定时调度时间
	ScheduledTime *string `json:"scheduledTime,omitempty" xml:"scheduledTime,omitempty"`
	// 截止时间开启
	HwDeadlineOpen *string `json:"hwDeadlineOpen,omitempty" xml:"hwDeadlineOpen,omitempty"`
	// 截止时间
	HwDeadline *int64 `json:"hwDeadline,omitempty" xml:"hwDeadline,omitempty"`
	// 作业类型
	HwType *string `json:"hwType,omitempty" xml:"hwType,omitempty"`
	// 选择跨组织班级
	OpenSelectItemList []*BatchOrgCreateHWRequestOpenSelectItemList `json:"openSelectItemList,omitempty" xml:"openSelectItemList,omitempty" type:"Repeated"`
	// 组织ID
	DingOrgId *int64 `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
}

func (s BatchOrgCreateHWRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchOrgCreateHWRequest) GoString() string {
	return s.String()
}

func (s *BatchOrgCreateHWRequest) SetHwMedia(v string) *BatchOrgCreateHWRequest {
	s.HwMedia = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetHwContent(v string) *BatchOrgCreateHWRequest {
	s.HwContent = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetHwTitle(v string) *BatchOrgCreateHWRequest {
	s.HwTitle = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetCourseName(v string) *BatchOrgCreateHWRequest {
	s.CourseName = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetHwPhoto(v string) *BatchOrgCreateHWRequest {
	s.HwPhoto = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetHwVideo(v string) *BatchOrgCreateHWRequest {
	s.HwVideo = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetTeacherName(v string) *BatchOrgCreateHWRequest {
	s.TeacherName = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetTeacherUserId(v string) *BatchOrgCreateHWRequest {
	s.TeacherUserId = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetIdentifier(v string) *BatchOrgCreateHWRequest {
	s.Identifier = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetAttributes(v string) *BatchOrgCreateHWRequest {
	s.Attributes = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetTargetRole(v string) *BatchOrgCreateHWRequest {
	s.TargetRole = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetBizCode(v string) *BatchOrgCreateHWRequest {
	s.BizCode = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetStatus(v string) *BatchOrgCreateHWRequest {
	s.Status = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetScheduledRelease(v string) *BatchOrgCreateHWRequest {
	s.ScheduledRelease = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetScheduledTime(v string) *BatchOrgCreateHWRequest {
	s.ScheduledTime = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetHwDeadlineOpen(v string) *BatchOrgCreateHWRequest {
	s.HwDeadlineOpen = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetHwDeadline(v int64) *BatchOrgCreateHWRequest {
	s.HwDeadline = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetHwType(v string) *BatchOrgCreateHWRequest {
	s.HwType = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetOpenSelectItemList(v []*BatchOrgCreateHWRequestOpenSelectItemList) *BatchOrgCreateHWRequest {
	s.OpenSelectItemList = v
	return s
}

func (s *BatchOrgCreateHWRequest) SetDingOrgId(v int64) *BatchOrgCreateHWRequest {
	s.DingOrgId = &v
	return s
}

type BatchOrgCreateHWRequestOpenSelectItemList struct {
	// 组织corpId
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 选择内容
	SelectedClassesDesc *string `json:"selectedClassesDesc,omitempty" xml:"selectedClassesDesc,omitempty"`
	// 班级列表
	ClassList []*BatchOrgCreateHWRequestOpenSelectItemListClassList `json:"classList,omitempty" xml:"classList,omitempty" type:"Repeated"`
}

func (s BatchOrgCreateHWRequestOpenSelectItemList) String() string {
	return tea.Prettify(s)
}

func (s BatchOrgCreateHWRequestOpenSelectItemList) GoString() string {
	return s.String()
}

func (s *BatchOrgCreateHWRequestOpenSelectItemList) SetCorpId(v string) *BatchOrgCreateHWRequestOpenSelectItemList {
	s.CorpId = &v
	return s
}

func (s *BatchOrgCreateHWRequestOpenSelectItemList) SetSelectedClassesDesc(v string) *BatchOrgCreateHWRequestOpenSelectItemList {
	s.SelectedClassesDesc = &v
	return s
}

func (s *BatchOrgCreateHWRequestOpenSelectItemList) SetClassList(v []*BatchOrgCreateHWRequestOpenSelectItemListClassList) *BatchOrgCreateHWRequestOpenSelectItemList {
	s.ClassList = v
	return s
}

type BatchOrgCreateHWRequestOpenSelectItemListClassList struct {
	// 班级id
	ClassId *string `json:"classId,omitempty" xml:"classId,omitempty"`
	// 班级名称
	ClassName *string `json:"className,omitempty" xml:"className,omitempty"`
	// 是否全选
	All *bool `json:"all,omitempty" xml:"all,omitempty"`
	// 学生列表
	Students []*BatchOrgCreateHWRequestOpenSelectItemListClassListStudents `json:"students,omitempty" xml:"students,omitempty" type:"Repeated"`
}

func (s BatchOrgCreateHWRequestOpenSelectItemListClassList) String() string {
	return tea.Prettify(s)
}

func (s BatchOrgCreateHWRequestOpenSelectItemListClassList) GoString() string {
	return s.String()
}

func (s *BatchOrgCreateHWRequestOpenSelectItemListClassList) SetClassId(v string) *BatchOrgCreateHWRequestOpenSelectItemListClassList {
	s.ClassId = &v
	return s
}

func (s *BatchOrgCreateHWRequestOpenSelectItemListClassList) SetClassName(v string) *BatchOrgCreateHWRequestOpenSelectItemListClassList {
	s.ClassName = &v
	return s
}

func (s *BatchOrgCreateHWRequestOpenSelectItemListClassList) SetAll(v bool) *BatchOrgCreateHWRequestOpenSelectItemListClassList {
	s.All = &v
	return s
}

func (s *BatchOrgCreateHWRequestOpenSelectItemListClassList) SetStudents(v []*BatchOrgCreateHWRequestOpenSelectItemListClassListStudents) *BatchOrgCreateHWRequestOpenSelectItemListClassList {
	s.Students = v
	return s
}

type BatchOrgCreateHWRequestOpenSelectItemListClassListStudents struct {
	// 学生姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 学生userid
	StaffId *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
	// 学生头像
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
}

func (s BatchOrgCreateHWRequestOpenSelectItemListClassListStudents) String() string {
	return tea.Prettify(s)
}

func (s BatchOrgCreateHWRequestOpenSelectItemListClassListStudents) GoString() string {
	return s.String()
}

func (s *BatchOrgCreateHWRequestOpenSelectItemListClassListStudents) SetName(v string) *BatchOrgCreateHWRequestOpenSelectItemListClassListStudents {
	s.Name = &v
	return s
}

func (s *BatchOrgCreateHWRequestOpenSelectItemListClassListStudents) SetStaffId(v string) *BatchOrgCreateHWRequestOpenSelectItemListClassListStudents {
	s.StaffId = &v
	return s
}

func (s *BatchOrgCreateHWRequestOpenSelectItemListClassListStudents) SetAvatar(v string) *BatchOrgCreateHWRequestOpenSelectItemListClassListStudents {
	s.Avatar = &v
	return s
}

type BatchOrgCreateHWResponseBody struct {
	Result *BatchOrgCreateHWResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s BatchOrgCreateHWResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchOrgCreateHWResponseBody) GoString() string {
	return s.String()
}

func (s *BatchOrgCreateHWResponseBody) SetResult(v *BatchOrgCreateHWResponseBodyResult) *BatchOrgCreateHWResponseBody {
	s.Result = v
	return s
}

type BatchOrgCreateHWResponseBodyResult struct {
	PublishList []*BatchOrgCreateHWResponseBodyResultPublishList `json:"publishList,omitempty" xml:"publishList,omitempty" type:"Repeated"`
}

func (s BatchOrgCreateHWResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s BatchOrgCreateHWResponseBodyResult) GoString() string {
	return s.String()
}

func (s *BatchOrgCreateHWResponseBodyResult) SetPublishList(v []*BatchOrgCreateHWResponseBodyResultPublishList) *BatchOrgCreateHWResponseBodyResult {
	s.PublishList = v
	return s
}

type BatchOrgCreateHWResponseBodyResultPublishList struct {
	Corpid *string `json:"corpid,omitempty" xml:"corpid,omitempty"`
	Hwid   *int64  `json:"hwid,omitempty" xml:"hwid,omitempty"`
}

func (s BatchOrgCreateHWResponseBodyResultPublishList) String() string {
	return tea.Prettify(s)
}

func (s BatchOrgCreateHWResponseBodyResultPublishList) GoString() string {
	return s.String()
}

func (s *BatchOrgCreateHWResponseBodyResultPublishList) SetCorpid(v string) *BatchOrgCreateHWResponseBodyResultPublishList {
	s.Corpid = &v
	return s
}

func (s *BatchOrgCreateHWResponseBodyResultPublishList) SetHwid(v int64) *BatchOrgCreateHWResponseBodyResultPublishList {
	s.Hwid = &v
	return s
}

type BatchOrgCreateHWResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchOrgCreateHWResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchOrgCreateHWResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchOrgCreateHWResponse) GoString() string {
	return s.String()
}

func (s *BatchOrgCreateHWResponse) SetHeaders(v map[string]*string) *BatchOrgCreateHWResponse {
	s.Headers = v
	return s
}

func (s *BatchOrgCreateHWResponse) SetBody(v *BatchOrgCreateHWResponseBody) *BatchOrgCreateHWResponse {
	s.Body = v
	return s
}

type CreateCustomDeptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateCustomDeptHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomDeptHeaders) GoString() string {
	return s.String()
}

func (s *CreateCustomDeptHeaders) SetCommonHeaders(v map[string]*string) *CreateCustomDeptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateCustomDeptHeaders) SetXAcsDingtalkAccessToken(v string) *CreateCustomDeptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateCustomDeptRequest struct {
	CustomDept *CreateCustomDeptRequestCustomDept `json:"customDept,omitempty" xml:"customDept,omitempty" type:"Struct"`
	// 上级部门ID（type为custom_campus时，必须为-7）
	SuperId *int64 `json:"superId,omitempty" xml:"superId,omitempty"`
	// 钉钉管理员员工ID
	Operator           *string `json:"operator,omitempty" xml:"operator,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingTokenGrantType *int32  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingOauthAppId     *int64  `json:"dingOauthAppId,omitempty" xml:"dingOauthAppId,omitempty"`
	DingCorpId         *string `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
}

func (s CreateCustomDeptRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomDeptRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomDeptRequest) SetCustomDept(v *CreateCustomDeptRequestCustomDept) *CreateCustomDeptRequest {
	s.CustomDept = v
	return s
}

func (s *CreateCustomDeptRequest) SetSuperId(v int64) *CreateCustomDeptRequest {
	s.SuperId = &v
	return s
}

func (s *CreateCustomDeptRequest) SetOperator(v string) *CreateCustomDeptRequest {
	s.Operator = &v
	return s
}

func (s *CreateCustomDeptRequest) SetDingOrgId(v int64) *CreateCustomDeptRequest {
	s.DingOrgId = &v
	return s
}

func (s *CreateCustomDeptRequest) SetDingTokenGrantType(v int32) *CreateCustomDeptRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *CreateCustomDeptRequest) SetDingSuiteKey(v string) *CreateCustomDeptRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *CreateCustomDeptRequest) SetDingOauthAppId(v int64) *CreateCustomDeptRequest {
	s.DingOauthAppId = &v
	return s
}

func (s *CreateCustomDeptRequest) SetDingCorpId(v string) *CreateCustomDeptRequest {
	s.DingCorpId = &v
	return s
}

func (s *CreateCustomDeptRequest) SetDingIsvOrgId(v int64) *CreateCustomDeptRequest {
	s.DingIsvOrgId = &v
	return s
}

type CreateCustomDeptRequestCustomDept struct {
	// 部门类型：custom_campus: 自定义校区；custom_dept: 自定义部门
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 自定义校区或部门名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateCustomDeptRequestCustomDept) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomDeptRequestCustomDept) GoString() string {
	return s.String()
}

func (s *CreateCustomDeptRequestCustomDept) SetType(v string) *CreateCustomDeptRequestCustomDept {
	s.Type = &v
	return s
}

func (s *CreateCustomDeptRequestCustomDept) SetName(v string) *CreateCustomDeptRequestCustomDept {
	s.Name = &v
	return s
}

type CreateCustomDeptResponseBody struct {
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// result
	Result *CreateCustomDeptResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateCustomDeptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomDeptResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomDeptResponseBody) SetSuccess(v bool) *CreateCustomDeptResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCustomDeptResponseBody) SetResult(v *CreateCustomDeptResponseBodyResult) *CreateCustomDeptResponseBody {
	s.Result = v
	return s
}

type CreateCustomDeptResponseBodyResult struct {
	// 部门ID
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s CreateCustomDeptResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomDeptResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateCustomDeptResponseBodyResult) SetDeptId(v int64) *CreateCustomDeptResponseBodyResult {
	s.DeptId = &v
	return s
}

type CreateCustomDeptResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCustomDeptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCustomDeptResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomDeptResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomDeptResponse) SetHeaders(v map[string]*string) *CreateCustomDeptResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomDeptResponse) SetBody(v *CreateCustomDeptResponseBody) *CreateCustomDeptResponse {
	s.Body = v
	return s
}

type StartCoursePrepareHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s StartCoursePrepareHeaders) String() string {
	return tea.Prettify(s)
}

func (s StartCoursePrepareHeaders) GoString() string {
	return s.String()
}

func (s *StartCoursePrepareHeaders) SetCommonHeaders(v map[string]*string) *StartCoursePrepareHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartCoursePrepareHeaders) SetXAcsDingtalkAccessToken(v string) *StartCoursePrepareHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type StartCoursePrepareRequest struct {
	// 拓展信息
	Ext *string `json:"ext,omitempty" xml:"ext,omitempty"`
	// 课程组编号
	CourseGroupCode *string `json:"courseGroupCode,omitempty" xml:"courseGroupCode,omitempty"`
	// 课节信息
	SectionIndex []*int32 `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty" type:"Repeated"`
	// 设备id
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// 封面url
	LiveCoverImage *string `json:"liveCoverImage,omitempty" xml:"liveCoverImage,omitempty"`
	// 上课日期
	CourseDate *string `json:"courseDate,omitempty" xml:"courseDate,omitempty"`
	// isv编号
	IsvCode *string `json:"isvCode,omitempty" xml:"isvCode,omitempty"`
	// 操作人
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s StartCoursePrepareRequest) String() string {
	return tea.Prettify(s)
}

func (s StartCoursePrepareRequest) GoString() string {
	return s.String()
}

func (s *StartCoursePrepareRequest) SetExt(v string) *StartCoursePrepareRequest {
	s.Ext = &v
	return s
}

func (s *StartCoursePrepareRequest) SetCourseGroupCode(v string) *StartCoursePrepareRequest {
	s.CourseGroupCode = &v
	return s
}

func (s *StartCoursePrepareRequest) SetSectionIndex(v []*int32) *StartCoursePrepareRequest {
	s.SectionIndex = v
	return s
}

func (s *StartCoursePrepareRequest) SetDeviceId(v string) *StartCoursePrepareRequest {
	s.DeviceId = &v
	return s
}

func (s *StartCoursePrepareRequest) SetLiveCoverImage(v string) *StartCoursePrepareRequest {
	s.LiveCoverImage = &v
	return s
}

func (s *StartCoursePrepareRequest) SetCourseDate(v string) *StartCoursePrepareRequest {
	s.CourseDate = &v
	return s
}

func (s *StartCoursePrepareRequest) SetIsvCode(v string) *StartCoursePrepareRequest {
	s.IsvCode = &v
	return s
}

func (s *StartCoursePrepareRequest) SetOpUserId(v string) *StartCoursePrepareRequest {
	s.OpUserId = &v
	return s
}

type StartCoursePrepareResponseBody struct {
	UniversityCourseCommonResponse *StartCoursePrepareResponseBodyUniversityCourseCommonResponse `json:"universityCourseCommonResponse,omitempty" xml:"universityCourseCommonResponse,omitempty" type:"Struct"`
}

func (s StartCoursePrepareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartCoursePrepareResponseBody) GoString() string {
	return s.String()
}

func (s *StartCoursePrepareResponseBody) SetUniversityCourseCommonResponse(v *StartCoursePrepareResponseBodyUniversityCourseCommonResponse) *StartCoursePrepareResponseBody {
	s.UniversityCourseCommonResponse = v
	return s
}

type StartCoursePrepareResponseBodyUniversityCourseCommonResponse struct {
	// 调用是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 课程编码
	CourseCode *string `json:"courseCode,omitempty" xml:"courseCode,omitempty"`
}

func (s StartCoursePrepareResponseBodyUniversityCourseCommonResponse) String() string {
	return tea.Prettify(s)
}

func (s StartCoursePrepareResponseBodyUniversityCourseCommonResponse) GoString() string {
	return s.String()
}

func (s *StartCoursePrepareResponseBodyUniversityCourseCommonResponse) SetSuccess(v bool) *StartCoursePrepareResponseBodyUniversityCourseCommonResponse {
	s.Success = &v
	return s
}

func (s *StartCoursePrepareResponseBodyUniversityCourseCommonResponse) SetCourseCode(v string) *StartCoursePrepareResponseBodyUniversityCourseCommonResponse {
	s.CourseCode = &v
	return s
}

type StartCoursePrepareResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartCoursePrepareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartCoursePrepareResponse) String() string {
	return tea.Prettify(s)
}

func (s StartCoursePrepareResponse) GoString() string {
	return s.String()
}

func (s *StartCoursePrepareResponse) SetHeaders(v map[string]*string) *StartCoursePrepareResponse {
	s.Headers = v
	return s
}

func (s *StartCoursePrepareResponse) SetBody(v *StartCoursePrepareResponseBody) *StartCoursePrepareResponse {
	s.Body = v
	return s
}

type QueryClassScheduleByTimeSchoolHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryClassScheduleByTimeSchoolHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleByTimeSchoolHeaders) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleByTimeSchoolHeaders) SetCommonHeaders(v map[string]*string) *QueryClassScheduleByTimeSchoolHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryClassScheduleByTimeSchoolHeaders) SetXAcsDingtalkAccessToken(v string) *QueryClassScheduleByTimeSchoolHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryClassScheduleByTimeSchoolRequest struct {
	// 开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 1621676000000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 1621566000000
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s QueryClassScheduleByTimeSchoolRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleByTimeSchoolRequest) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleByTimeSchoolRequest) SetStartTime(v int64) *QueryClassScheduleByTimeSchoolRequest {
	s.StartTime = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolRequest) SetEndTime(v int64) *QueryClassScheduleByTimeSchoolRequest {
	s.EndTime = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolRequest) SetOpUserId(v string) *QueryClassScheduleByTimeSchoolRequest {
	s.OpUserId = &v
	return s
}

type QueryClassScheduleByTimeSchoolResponseBody struct {
	Result []*QueryClassScheduleByTimeSchoolResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryClassScheduleByTimeSchoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleByTimeSchoolResponseBody) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleByTimeSchoolResponseBody) SetResult(v []*QueryClassScheduleByTimeSchoolResponseBodyResult) *QueryClassScheduleByTimeSchoolResponseBody {
	s.Result = v
	return s
}

type QueryClassScheduleByTimeSchoolResponseBodyResult struct {
	// 课程编码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 课程名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 课程介绍
	Introduce *string `json:"introduce,omitempty" xml:"introduce,omitempty"`
	// 课程封面地址
	CoverUrl *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	// 开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 结束时间
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 创建者组织id
	CreatorCorpId *string `json:"creatorCorpId,omitempty" xml:"creatorCorpId,omitempty"`
	// 创建者UserId
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// 创建者UserName
	CreatorUserName *string `json:"creatorUserName,omitempty" xml:"creatorUserName,omitempty"`
	// 老师CorpId
	TeacherCorpId *string `json:"teacherCorpId,omitempty" xml:"teacherCorpId,omitempty"`
	// 老师UserId
	TeacherUserId *string `json:"teacherUserId,omitempty" xml:"teacherUserId,omitempty"`
	// 老师UserName
	TeacherUserName *string `json:"teacherUserName,omitempty" xml:"teacherUserName,omitempty"`
	// 业务唯一键
	BizKey *string `json:"bizKey,omitempty" xml:"bizKey,omitempty"`
	// 学科编码
	SubjectCode *string `json:"subjectCode,omitempty" xml:"subjectCode,omitempty"`
	// 课程组编码
	CourseGroupCode *string `json:"courseGroupCode,omitempty" xml:"courseGroupCode,omitempty"`
	// 课程状态
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 课堂列表
	Classrooms []*QueryClassScheduleByTimeSchoolResponseBodyResultClassrooms `json:"classrooms,omitempty" xml:"classrooms,omitempty" type:"Repeated"`
	// 课程参与人列表
	EduUserModels []*QueryClassScheduleByTimeSchoolResponseBodyResultEduUserModels `json:"eduUserModels,omitempty" xml:"eduUserModels,omitempty" type:"Repeated"`
	// 课程编码
	SectionName *string `json:"sectionName,omitempty" xml:"sectionName,omitempty"`
	// 课程所在节次序列号
	SectionIndex *int64 `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty"`
	// 课程所在班级id
	ClassId *int64 `json:"classId,omitempty" xml:"classId,omitempty"`
	// 课程扩展信息
	ExtInfo *string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
}

func (s QueryClassScheduleByTimeSchoolResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleByTimeSchoolResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetCode(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.Code = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetName(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.Name = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetIntroduce(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.Introduce = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetCoverUrl(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.CoverUrl = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetStartTime(v int64) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.StartTime = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetEndTime(v int64) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetCreatorCorpId(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.CreatorCorpId = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetCreatorUserId(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.CreatorUserId = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetCreatorUserName(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.CreatorUserName = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetTeacherCorpId(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.TeacherCorpId = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetTeacherUserId(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.TeacherUserId = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetTeacherUserName(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.TeacherUserName = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetBizKey(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.BizKey = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetSubjectCode(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.SubjectCode = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetCourseGroupCode(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.CourseGroupCode = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetStatus(v int64) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.Status = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetClassrooms(v []*QueryClassScheduleByTimeSchoolResponseBodyResultClassrooms) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.Classrooms = v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetEduUserModels(v []*QueryClassScheduleByTimeSchoolResponseBodyResultEduUserModels) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.EduUserModels = v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetSectionName(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.SectionName = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetSectionIndex(v int64) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.SectionIndex = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetClassId(v int64) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.ClassId = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetExtInfo(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.ExtInfo = &v
	return s
}

type QueryClassScheduleByTimeSchoolResponseBodyResultClassrooms struct {
	// 课堂唯一标识
	TargetId *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
	// 交互信息
	InteractInfo *string `json:"interactInfo,omitempty" xml:"interactInfo,omitempty"`
}

func (s QueryClassScheduleByTimeSchoolResponseBodyResultClassrooms) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleByTimeSchoolResponseBodyResultClassrooms) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResultClassrooms) SetTargetId(v string) *QueryClassScheduleByTimeSchoolResponseBodyResultClassrooms {
	s.TargetId = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResultClassrooms) SetInteractInfo(v string) *QueryClassScheduleByTimeSchoolResponseBodyResultClassrooms {
	s.InteractInfo = &v
	return s
}

type QueryClassScheduleByTimeSchoolResponseBodyResultEduUserModels struct {
	// 用户userid
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户uid
	Uid  *int64  `json:"uid,omitempty" xml:"uid,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryClassScheduleByTimeSchoolResponseBodyResultEduUserModels) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleByTimeSchoolResponseBodyResultEduUserModels) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResultEduUserModels) SetUserId(v string) *QueryClassScheduleByTimeSchoolResponseBodyResultEduUserModels {
	s.UserId = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResultEduUserModels) SetUid(v int64) *QueryClassScheduleByTimeSchoolResponseBodyResultEduUserModels {
	s.Uid = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResultEduUserModels) SetName(v string) *QueryClassScheduleByTimeSchoolResponseBodyResultEduUserModels {
	s.Name = &v
	return s
}

type QueryClassScheduleByTimeSchoolResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryClassScheduleByTimeSchoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryClassScheduleByTimeSchoolResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleByTimeSchoolResponse) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleByTimeSchoolResponse) SetHeaders(v map[string]*string) *QueryClassScheduleByTimeSchoolResponse {
	s.Headers = v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponse) SetBody(v *QueryClassScheduleByTimeSchoolResponseBody) *QueryClassScheduleByTimeSchoolResponse {
	s.Body = v
	return s
}

type QueryTeachSubjectsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryTeachSubjectsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryTeachSubjectsHeaders) GoString() string {
	return s.String()
}

func (s *QueryTeachSubjectsHeaders) SetCommonHeaders(v map[string]*string) *QueryTeachSubjectsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryTeachSubjectsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryTeachSubjectsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryTeachSubjectsRequest struct {
	// 班级ids
	ClassIds []*int64 `json:"classIds,omitempty" xml:"classIds,omitempty" type:"Repeated"`
	// 操作者UserId
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s QueryTeachSubjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTeachSubjectsRequest) GoString() string {
	return s.String()
}

func (s *QueryTeachSubjectsRequest) SetClassIds(v []*int64) *QueryTeachSubjectsRequest {
	s.ClassIds = v
	return s
}

func (s *QueryTeachSubjectsRequest) SetOpUserId(v string) *QueryTeachSubjectsRequest {
	s.OpUserId = &v
	return s
}

type QueryTeachSubjectsResponseBody struct {
	// 结果
	Result []*QueryTeachSubjectsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryTeachSubjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTeachSubjectsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTeachSubjectsResponseBody) SetResult(v []*QueryTeachSubjectsResponseBodyResult) *QueryTeachSubjectsResponseBody {
	s.Result = v
	return s
}

type QueryTeachSubjectsResponseBodyResult struct {
	// 老师名称
	TeacherName *string `json:"teacherName,omitempty" xml:"teacherName,omitempty"`
	// 学科名称
	SubjectName *string `json:"subjectName,omitempty" xml:"subjectName,omitempty"`
	// 学科code
	SubjectCode *string `json:"subjectCode,omitempty" xml:"subjectCode,omitempty"`
	// 学段code
	PeriodCode *string `json:"periodCode,omitempty" xml:"periodCode,omitempty"`
	// 学校corpId
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 老师Userid
	TeacherUserId *string `json:"teacherUserId,omitempty" xml:"teacherUserId,omitempty"`
	// 班级id
	ClassId *int64 `json:"classId,omitempty" xml:"classId,omitempty"`
}

func (s QueryTeachSubjectsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryTeachSubjectsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryTeachSubjectsResponseBodyResult) SetTeacherName(v string) *QueryTeachSubjectsResponseBodyResult {
	s.TeacherName = &v
	return s
}

func (s *QueryTeachSubjectsResponseBodyResult) SetSubjectName(v string) *QueryTeachSubjectsResponseBodyResult {
	s.SubjectName = &v
	return s
}

func (s *QueryTeachSubjectsResponseBodyResult) SetSubjectCode(v string) *QueryTeachSubjectsResponseBodyResult {
	s.SubjectCode = &v
	return s
}

func (s *QueryTeachSubjectsResponseBodyResult) SetPeriodCode(v string) *QueryTeachSubjectsResponseBodyResult {
	s.PeriodCode = &v
	return s
}

func (s *QueryTeachSubjectsResponseBodyResult) SetCorpId(v string) *QueryTeachSubjectsResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *QueryTeachSubjectsResponseBodyResult) SetTeacherUserId(v string) *QueryTeachSubjectsResponseBodyResult {
	s.TeacherUserId = &v
	return s
}

func (s *QueryTeachSubjectsResponseBodyResult) SetClassId(v int64) *QueryTeachSubjectsResponseBodyResult {
	s.ClassId = &v
	return s
}

type QueryTeachSubjectsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTeachSubjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTeachSubjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTeachSubjectsResponse) GoString() string {
	return s.String()
}

func (s *QueryTeachSubjectsResponse) SetHeaders(v map[string]*string) *QueryTeachSubjectsResponse {
	s.Headers = v
	return s
}

func (s *QueryTeachSubjectsResponse) SetBody(v *QueryTeachSubjectsResponseBody) *QueryTeachSubjectsResponse {
	s.Body = v
	return s
}

type GetShareRolesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetShareRolesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetShareRolesHeaders) GoString() string {
	return s.String()
}

func (s *GetShareRolesHeaders) SetCommonHeaders(v map[string]*string) *GetShareRolesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetShareRolesHeaders) SetXAcsDingtalkAccessToken(v string) *GetShareRolesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetShareRolesResponseBody struct {
	Result []*GetShareRolesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetShareRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetShareRolesResponseBody) GoString() string {
	return s.String()
}

func (s *GetShareRolesResponseBody) SetResult(v []*GetShareRolesResponseBodyResult) *GetShareRolesResponseBody {
	s.Result = v
	return s
}

type GetShareRolesResponseBodyResult struct {
	// 角色code
	ShareRoleCode *string `json:"shareRoleCode,omitempty" xml:"shareRoleCode,omitempty"`
	// 角色名称
	ShareRoleName *string `json:"shareRoleName,omitempty" xml:"shareRoleName,omitempty"`
}

func (s GetShareRolesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetShareRolesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetShareRolesResponseBodyResult) SetShareRoleCode(v string) *GetShareRolesResponseBodyResult {
	s.ShareRoleCode = &v
	return s
}

func (s *GetShareRolesResponseBodyResult) SetShareRoleName(v string) *GetShareRolesResponseBodyResult {
	s.ShareRoleName = &v
	return s
}

type GetShareRolesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetShareRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetShareRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetShareRolesResponse) GoString() string {
	return s.String()
}

func (s *GetShareRolesResponse) SetHeaders(v map[string]*string) *GetShareRolesResponse {
	s.Headers = v
	return s
}

func (s *GetShareRolesResponse) SetBody(v *GetShareRolesResponseBody) *GetShareRolesResponse {
	s.Body = v
	return s
}

type SubscribeUniversityCourseGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SubscribeUniversityCourseGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s SubscribeUniversityCourseGroupHeaders) GoString() string {
	return s.String()
}

func (s *SubscribeUniversityCourseGroupHeaders) SetCommonHeaders(v map[string]*string) *SubscribeUniversityCourseGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SubscribeUniversityCourseGroupHeaders) SetXAcsDingtalkAccessToken(v string) *SubscribeUniversityCourseGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SubscribeUniversityCourseGroupRequest struct {
	// 课程组编号
	CourseGroupCode *string `json:"courseGroupCode,omitempty" xml:"courseGroupCode,omitempty"`
	// 学生用户Id
	StudentUserIds []*string `json:"studentUserIds,omitempty" xml:"studentUserIds,omitempty" type:"Repeated"`
	// 操作人id
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s SubscribeUniversityCourseGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s SubscribeUniversityCourseGroupRequest) GoString() string {
	return s.String()
}

func (s *SubscribeUniversityCourseGroupRequest) SetCourseGroupCode(v string) *SubscribeUniversityCourseGroupRequest {
	s.CourseGroupCode = &v
	return s
}

func (s *SubscribeUniversityCourseGroupRequest) SetStudentUserIds(v []*string) *SubscribeUniversityCourseGroupRequest {
	s.StudentUserIds = v
	return s
}

func (s *SubscribeUniversityCourseGroupRequest) SetOpUserId(v string) *SubscribeUniversityCourseGroupRequest {
	s.OpUserId = &v
	return s
}

type SubscribeUniversityCourseGroupResponseBody struct {
	// 订阅结果
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SubscribeUniversityCourseGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubscribeUniversityCourseGroupResponseBody) GoString() string {
	return s.String()
}

func (s *SubscribeUniversityCourseGroupResponseBody) SetResult(v bool) *SubscribeUniversityCourseGroupResponseBody {
	s.Result = &v
	return s
}

type SubscribeUniversityCourseGroupResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubscribeUniversityCourseGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubscribeUniversityCourseGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s SubscribeUniversityCourseGroupResponse) GoString() string {
	return s.String()
}

func (s *SubscribeUniversityCourseGroupResponse) SetHeaders(v map[string]*string) *SubscribeUniversityCourseGroupResponse {
	s.Headers = v
	return s
}

func (s *SubscribeUniversityCourseGroupResponse) SetBody(v *SubscribeUniversityCourseGroupResponseBody) *SubscribeUniversityCourseGroupResponse {
	s.Body = v
	return s
}

type DeleteOrgRelationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteOrgRelationHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrgRelationHeaders) GoString() string {
	return s.String()
}

func (s *DeleteOrgRelationHeaders) SetCommonHeaders(v map[string]*string) *DeleteOrgRelationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteOrgRelationHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteOrgRelationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteOrgRelationRequest struct {
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
	AuthCode     *string `json:"authCode,omitempty" xml:"authCode,omitempty"`
}

func (s DeleteOrgRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrgRelationRequest) GoString() string {
	return s.String()
}

func (s *DeleteOrgRelationRequest) SetTargetCorpId(v string) *DeleteOrgRelationRequest {
	s.TargetCorpId = &v
	return s
}

func (s *DeleteOrgRelationRequest) SetAuthCode(v string) *DeleteOrgRelationRequest {
	s.AuthCode = &v
	return s
}

type DeleteOrgRelationResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteOrgRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrgRelationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOrgRelationResponseBody) SetSuccess(v bool) *DeleteOrgRelationResponseBody {
	s.Success = &v
	return s
}

type DeleteOrgRelationResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteOrgRelationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteOrgRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrgRelationResponse) GoString() string {
	return s.String()
}

func (s *DeleteOrgRelationResponse) SetHeaders(v map[string]*string) *DeleteOrgRelationResponse {
	s.Headers = v
	return s
}

func (s *DeleteOrgRelationResponse) SetBody(v *DeleteOrgRelationResponseBody) *DeleteOrgRelationResponse {
	s.Body = v
	return s
}

type StartCourseRequest struct {
	// opUserId
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 拓展字段
	Ext *string `json:"ext,omitempty" xml:"ext,omitempty"`
	// 课程编码
	CourseCode *string `json:"courseCode,omitempty" xml:"courseCode,omitempty"`
	// livePlayInfoList
	LivePlayInfoList []*StartCourseRequestLivePlayInfoList `json:"livePlayInfoList,omitempty" xml:"livePlayInfoList,omitempty" type:"Repeated"`
	// isvCode
	IsvCode *string `json:"isvCode,omitempty" xml:"isvCode,omitempty"`
}

func (s StartCourseRequest) String() string {
	return tea.Prettify(s)
}

func (s StartCourseRequest) GoString() string {
	return s.String()
}

func (s *StartCourseRequest) SetOpUserId(v string) *StartCourseRequest {
	s.OpUserId = &v
	return s
}

func (s *StartCourseRequest) SetExt(v string) *StartCourseRequest {
	s.Ext = &v
	return s
}

func (s *StartCourseRequest) SetCourseCode(v string) *StartCourseRequest {
	s.CourseCode = &v
	return s
}

func (s *StartCourseRequest) SetLivePlayInfoList(v []*StartCourseRequestLivePlayInfoList) *StartCourseRequest {
	s.LivePlayInfoList = v
	return s
}

func (s *StartCourseRequest) SetIsvCode(v string) *StartCourseRequest {
	s.IsvCode = &v
	return s
}

type StartCourseRequestLivePlayInfoList struct {
	// 直播拉流地址
	LiveOutputUrl *string `json:"liveOutputUrl,omitempty" xml:"liveOutputUrl,omitempty"`
	// 直播流类型
	LiveType *int32 `json:"liveType,omitempty" xml:"liveType,omitempty"`
	// 直播推流地址
	LiveInputUrl *string `json:"liveInputUrl,omitempty" xml:"liveInputUrl,omitempty"`
	// 视频回放地址
	ReplayUrl *string `json:"replayUrl,omitempty" xml:"replayUrl,omitempty"`
}

func (s StartCourseRequestLivePlayInfoList) String() string {
	return tea.Prettify(s)
}

func (s StartCourseRequestLivePlayInfoList) GoString() string {
	return s.String()
}

func (s *StartCourseRequestLivePlayInfoList) SetLiveOutputUrl(v string) *StartCourseRequestLivePlayInfoList {
	s.LiveOutputUrl = &v
	return s
}

func (s *StartCourseRequestLivePlayInfoList) SetLiveType(v int32) *StartCourseRequestLivePlayInfoList {
	s.LiveType = &v
	return s
}

func (s *StartCourseRequestLivePlayInfoList) SetLiveInputUrl(v string) *StartCourseRequestLivePlayInfoList {
	s.LiveInputUrl = &v
	return s
}

func (s *StartCourseRequestLivePlayInfoList) SetReplayUrl(v string) *StartCourseRequestLivePlayInfoList {
	s.ReplayUrl = &v
	return s
}

type StartCourseResponseBody struct {
	UniversityCourseCommonResponse *StartCourseResponseBodyUniversityCourseCommonResponse `json:"universityCourseCommonResponse,omitempty" xml:"universityCourseCommonResponse,omitempty" type:"Struct"`
}

func (s StartCourseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartCourseResponseBody) GoString() string {
	return s.String()
}

func (s *StartCourseResponseBody) SetUniversityCourseCommonResponse(v *StartCourseResponseBodyUniversityCourseCommonResponse) *StartCourseResponseBody {
	s.UniversityCourseCommonResponse = v
	return s
}

type StartCourseResponseBodyUniversityCourseCommonResponse struct {
	// 调用是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 课程编码
	CourseCode *string `json:"courseCode,omitempty" xml:"courseCode,omitempty"`
}

func (s StartCourseResponseBodyUniversityCourseCommonResponse) String() string {
	return tea.Prettify(s)
}

func (s StartCourseResponseBodyUniversityCourseCommonResponse) GoString() string {
	return s.String()
}

func (s *StartCourseResponseBodyUniversityCourseCommonResponse) SetSuccess(v bool) *StartCourseResponseBodyUniversityCourseCommonResponse {
	s.Success = &v
	return s
}

func (s *StartCourseResponseBodyUniversityCourseCommonResponse) SetCourseCode(v string) *StartCourseResponseBodyUniversityCourseCommonResponse {
	s.CourseCode = &v
	return s
}

type StartCourseResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartCourseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartCourseResponse) String() string {
	return tea.Prettify(s)
}

func (s StartCourseResponse) GoString() string {
	return s.String()
}

func (s *StartCourseResponse) SetHeaders(v map[string]*string) *StartCourseResponse {
	s.Headers = v
	return s
}

func (s *StartCourseResponse) SetBody(v *StartCourseResponseBody) *StartCourseResponse {
	s.Body = v
	return s
}

type QueryPhysicalClassroomHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryPhysicalClassroomHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryPhysicalClassroomHeaders) GoString() string {
	return s.String()
}

func (s *QueryPhysicalClassroomHeaders) SetCommonHeaders(v map[string]*string) *QueryPhysicalClassroomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryPhysicalClassroomHeaders) SetXAcsDingtalkAccessToken(v string) *QueryPhysicalClassroomHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryPhysicalClassroomRequest struct {
	// 操作人id
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 教室id
	ClassroomId *int64 `json:"classroomId,omitempty" xml:"classroomId,omitempty"`
}

func (s QueryPhysicalClassroomRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPhysicalClassroomRequest) GoString() string {
	return s.String()
}

func (s *QueryPhysicalClassroomRequest) SetOpUserId(v string) *QueryPhysicalClassroomRequest {
	s.OpUserId = &v
	return s
}

func (s *QueryPhysicalClassroomRequest) SetClassroomId(v int64) *QueryPhysicalClassroomRequest {
	s.ClassroomId = &v
	return s
}

type QueryPhysicalClassroomResponseBody struct {
	// 请求是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回结果
	Result *QueryPhysicalClassroomResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryPhysicalClassroomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPhysicalClassroomResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPhysicalClassroomResponseBody) SetSuccess(v bool) *QueryPhysicalClassroomResponseBody {
	s.Success = &v
	return s
}

func (s *QueryPhysicalClassroomResponseBody) SetResult(v *QueryPhysicalClassroomResponseBodyResult) *QueryPhysicalClassroomResponseBody {
	s.Result = v
	return s
}

type QueryPhysicalClassroomResponseBodyResult struct {
	// 教室ID
	ClassroomId *int64 `json:"classroomId,omitempty" xml:"classroomId,omitempty"`
	// 教室名称
	ClassroomName *string `json:"classroomName,omitempty" xml:"classroomName,omitempty"`
	// 教室校区
	ClassroomCampus *string `json:"classroomCampus,omitempty" xml:"classroomCampus,omitempty"`
	// 教室教学楼
	ClassroomBuilding *string `json:"classroomBuilding,omitempty" xml:"classroomBuilding,omitempty"`
	// 教室楼层
	ClassroomFloor *string `json:"classroomFloor,omitempty" xml:"classroomFloor,omitempty"`
	// 教室房间号
	ClassroomNumber *string `json:"classroomNumber,omitempty" xml:"classroomNumber,omitempty"`
	// 是否支持直播
	DirectBroadcast *string `json:"directBroadcast,omitempty" xml:"directBroadcast,omitempty"`
}

func (s QueryPhysicalClassroomResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryPhysicalClassroomResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryPhysicalClassroomResponseBodyResult) SetClassroomId(v int64) *QueryPhysicalClassroomResponseBodyResult {
	s.ClassroomId = &v
	return s
}

func (s *QueryPhysicalClassroomResponseBodyResult) SetClassroomName(v string) *QueryPhysicalClassroomResponseBodyResult {
	s.ClassroomName = &v
	return s
}

func (s *QueryPhysicalClassroomResponseBodyResult) SetClassroomCampus(v string) *QueryPhysicalClassroomResponseBodyResult {
	s.ClassroomCampus = &v
	return s
}

func (s *QueryPhysicalClassroomResponseBodyResult) SetClassroomBuilding(v string) *QueryPhysicalClassroomResponseBodyResult {
	s.ClassroomBuilding = &v
	return s
}

func (s *QueryPhysicalClassroomResponseBodyResult) SetClassroomFloor(v string) *QueryPhysicalClassroomResponseBodyResult {
	s.ClassroomFloor = &v
	return s
}

func (s *QueryPhysicalClassroomResponseBodyResult) SetClassroomNumber(v string) *QueryPhysicalClassroomResponseBodyResult {
	s.ClassroomNumber = &v
	return s
}

func (s *QueryPhysicalClassroomResponseBodyResult) SetDirectBroadcast(v string) *QueryPhysicalClassroomResponseBodyResult {
	s.DirectBroadcast = &v
	return s
}

type QueryPhysicalClassroomResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryPhysicalClassroomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPhysicalClassroomResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPhysicalClassroomResponse) GoString() string {
	return s.String()
}

func (s *QueryPhysicalClassroomResponse) SetHeaders(v map[string]*string) *QueryPhysicalClassroomResponse {
	s.Headers = v
	return s
}

func (s *QueryPhysicalClassroomResponse) SetBody(v *QueryPhysicalClassroomResponseBody) *QueryPhysicalClassroomResponse {
	s.Body = v
	return s
}

type UpdateRemoteClassCourseHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateRemoteClassCourseHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateRemoteClassCourseHeaders) GoString() string {
	return s.String()
}

func (s *UpdateRemoteClassCourseHeaders) SetCommonHeaders(v map[string]*string) *UpdateRemoteClassCourseHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateRemoteClassCourseHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateRemoteClassCourseHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateRemoteClassCourseRequest struct {
	// 课程名称
	CourseName *string `json:"courseName,omitempty" xml:"courseName,omitempty"`
	// 课程开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 课程结束时间
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 授课设备
	TeachingParticipant *UpdateRemoteClassCourseRequestTeachingParticipant `json:"teachingParticipant,omitempty" xml:"teachingParticipant,omitempty" type:"Struct"`
	// 听课设备
	AttendParticipants []*UpdateRemoteClassCourseRequestAttendParticipants `json:"attendParticipants,omitempty" xml:"attendParticipants,omitempty" type:"Repeated"`
	// 课程码
	CourseCode         *string `json:"courseCode,omitempty" xml:"courseCode,omitempty"`
	DingClientId       *string `json:"dingClientId,omitempty" xml:"dingClientId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int32  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	DingOauthAppId     *int64  `json:"dingOauthAppId,omitempty" xml:"dingOauthAppId,omitempty"`
	DingCorpId         *string `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	AuthCode           *string `json:"authCode,omitempty" xml:"authCode,omitempty"`
}

func (s UpdateRemoteClassCourseRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRemoteClassCourseRequest) GoString() string {
	return s.String()
}

func (s *UpdateRemoteClassCourseRequest) SetCourseName(v string) *UpdateRemoteClassCourseRequest {
	s.CourseName = &v
	return s
}

func (s *UpdateRemoteClassCourseRequest) SetStartTime(v int64) *UpdateRemoteClassCourseRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateRemoteClassCourseRequest) SetEndTime(v int64) *UpdateRemoteClassCourseRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateRemoteClassCourseRequest) SetTeachingParticipant(v *UpdateRemoteClassCourseRequestTeachingParticipant) *UpdateRemoteClassCourseRequest {
	s.TeachingParticipant = v
	return s
}

func (s *UpdateRemoteClassCourseRequest) SetAttendParticipants(v []*UpdateRemoteClassCourseRequestAttendParticipants) *UpdateRemoteClassCourseRequest {
	s.AttendParticipants = v
	return s
}

func (s *UpdateRemoteClassCourseRequest) SetCourseCode(v string) *UpdateRemoteClassCourseRequest {
	s.CourseCode = &v
	return s
}

func (s *UpdateRemoteClassCourseRequest) SetDingClientId(v string) *UpdateRemoteClassCourseRequest {
	s.DingClientId = &v
	return s
}

func (s *UpdateRemoteClassCourseRequest) SetDingSuiteKey(v string) *UpdateRemoteClassCourseRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *UpdateRemoteClassCourseRequest) SetDingTokenGrantType(v int32) *UpdateRemoteClassCourseRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *UpdateRemoteClassCourseRequest) SetDingOauthAppId(v int64) *UpdateRemoteClassCourseRequest {
	s.DingOauthAppId = &v
	return s
}

func (s *UpdateRemoteClassCourseRequest) SetDingCorpId(v string) *UpdateRemoteClassCourseRequest {
	s.DingCorpId = &v
	return s
}

func (s *UpdateRemoteClassCourseRequest) SetDingIsvOrgId(v int64) *UpdateRemoteClassCourseRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *UpdateRemoteClassCourseRequest) SetAuthCode(v string) *UpdateRemoteClassCourseRequest {
	s.AuthCode = &v
	return s
}

type UpdateRemoteClassCourseRequestTeachingParticipant struct {
	// 参与方ID
	ParticipantId *string `json:"participantId,omitempty" xml:"participantId,omitempty"`
	// 组织ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s UpdateRemoteClassCourseRequestTeachingParticipant) String() string {
	return tea.Prettify(s)
}

func (s UpdateRemoteClassCourseRequestTeachingParticipant) GoString() string {
	return s.String()
}

func (s *UpdateRemoteClassCourseRequestTeachingParticipant) SetParticipantId(v string) *UpdateRemoteClassCourseRequestTeachingParticipant {
	s.ParticipantId = &v
	return s
}

func (s *UpdateRemoteClassCourseRequestTeachingParticipant) SetCorpId(v string) *UpdateRemoteClassCourseRequestTeachingParticipant {
	s.CorpId = &v
	return s
}

type UpdateRemoteClassCourseRequestAttendParticipants struct {
	// 参与方ID
	ParticipantId *string `json:"participantId,omitempty" xml:"participantId,omitempty"`
	// 组织ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s UpdateRemoteClassCourseRequestAttendParticipants) String() string {
	return tea.Prettify(s)
}

func (s UpdateRemoteClassCourseRequestAttendParticipants) GoString() string {
	return s.String()
}

func (s *UpdateRemoteClassCourseRequestAttendParticipants) SetParticipantId(v string) *UpdateRemoteClassCourseRequestAttendParticipants {
	s.ParticipantId = &v
	return s
}

func (s *UpdateRemoteClassCourseRequestAttendParticipants) SetCorpId(v string) *UpdateRemoteClassCourseRequestAttendParticipants {
	s.CorpId = &v
	return s
}

type UpdateRemoteClassCourseResponseBody struct {
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// result
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateRemoteClassCourseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRemoteClassCourseResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRemoteClassCourseResponseBody) SetSuccess(v bool) *UpdateRemoteClassCourseResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateRemoteClassCourseResponseBody) SetResult(v string) *UpdateRemoteClassCourseResponseBody {
	s.Result = &v
	return s
}

type UpdateRemoteClassCourseResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateRemoteClassCourseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRemoteClassCourseResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRemoteClassCourseResponse) GoString() string {
	return s.String()
}

func (s *UpdateRemoteClassCourseResponse) SetHeaders(v map[string]*string) *UpdateRemoteClassCourseResponse {
	s.Headers = v
	return s
}

func (s *UpdateRemoteClassCourseResponse) SetBody(v *UpdateRemoteClassCourseResponseBody) *UpdateRemoteClassCourseResponse {
	s.Body = v
	return s
}

type UnsubscribeUniversityCourseGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UnsubscribeUniversityCourseGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeUniversityCourseGroupHeaders) GoString() string {
	return s.String()
}

func (s *UnsubscribeUniversityCourseGroupHeaders) SetCommonHeaders(v map[string]*string) *UnsubscribeUniversityCourseGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UnsubscribeUniversityCourseGroupHeaders) SetXAcsDingtalkAccessToken(v string) *UnsubscribeUniversityCourseGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UnsubscribeUniversityCourseGroupRequest struct {
	// 课程组
	CourseGroupCode *string `json:"courseGroupCode,omitempty" xml:"courseGroupCode,omitempty"`
	// 学生用户id
	StudentUserIds []*string `json:"studentUserIds,omitempty" xml:"studentUserIds,omitempty" type:"Repeated"`
	// opUserId
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s UnsubscribeUniversityCourseGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeUniversityCourseGroupRequest) GoString() string {
	return s.String()
}

func (s *UnsubscribeUniversityCourseGroupRequest) SetCourseGroupCode(v string) *UnsubscribeUniversityCourseGroupRequest {
	s.CourseGroupCode = &v
	return s
}

func (s *UnsubscribeUniversityCourseGroupRequest) SetStudentUserIds(v []*string) *UnsubscribeUniversityCourseGroupRequest {
	s.StudentUserIds = v
	return s
}

func (s *UnsubscribeUniversityCourseGroupRequest) SetOpUserId(v string) *UnsubscribeUniversityCourseGroupRequest {
	s.OpUserId = &v
	return s
}

type UnsubscribeUniversityCourseGroupResponseBody struct {
	// 取消订阅结果
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UnsubscribeUniversityCourseGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeUniversityCourseGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UnsubscribeUniversityCourseGroupResponseBody) SetResult(v bool) *UnsubscribeUniversityCourseGroupResponseBody {
	s.Result = &v
	return s
}

type UnsubscribeUniversityCourseGroupResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnsubscribeUniversityCourseGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnsubscribeUniversityCourseGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeUniversityCourseGroupResponse) GoString() string {
	return s.String()
}

func (s *UnsubscribeUniversityCourseGroupResponse) SetHeaders(v map[string]*string) *UnsubscribeUniversityCourseGroupResponse {
	s.Headers = v
	return s
}

func (s *UnsubscribeUniversityCourseGroupResponse) SetBody(v *UnsubscribeUniversityCourseGroupResponseBody) *UnsubscribeUniversityCourseGroupResponse {
	s.Body = v
	return s
}

type QueryOrgSecretKeyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOrgSecretKeyHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgSecretKeyHeaders) GoString() string {
	return s.String()
}

func (s *QueryOrgSecretKeyHeaders) SetCommonHeaders(v map[string]*string) *QueryOrgSecretKeyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOrgSecretKeyHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOrgSecretKeyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOrgSecretKeyRequest struct {
	// 合作方编码
	IsvCode *string `json:"isvCode,omitempty" xml:"isvCode,omitempty"`
	// 操作人
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s QueryOrgSecretKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgSecretKeyRequest) GoString() string {
	return s.String()
}

func (s *QueryOrgSecretKeyRequest) SetIsvCode(v string) *QueryOrgSecretKeyRequest {
	s.IsvCode = &v
	return s
}

func (s *QueryOrgSecretKeyRequest) SetOpUserId(v string) *QueryOrgSecretKeyRequest {
	s.OpUserId = &v
	return s
}

type QueryOrgSecretKeyResponseBody struct {
	// 秘钥信息
	UniversitySecretKeyInfo *QueryOrgSecretKeyResponseBodyUniversitySecretKeyInfo `json:"universitySecretKeyInfo,omitempty" xml:"universitySecretKeyInfo,omitempty" type:"Struct"`
}

func (s QueryOrgSecretKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgSecretKeyResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrgSecretKeyResponseBody) SetUniversitySecretKeyInfo(v *QueryOrgSecretKeyResponseBodyUniversitySecretKeyInfo) *QueryOrgSecretKeyResponseBody {
	s.UniversitySecretKeyInfo = v
	return s
}

type QueryOrgSecretKeyResponseBodyUniversitySecretKeyInfo struct {
	// 秘钥key
	SecretKey *string `json:"secretKey,omitempty" xml:"secretKey,omitempty"`
}

func (s QueryOrgSecretKeyResponseBodyUniversitySecretKeyInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgSecretKeyResponseBodyUniversitySecretKeyInfo) GoString() string {
	return s.String()
}

func (s *QueryOrgSecretKeyResponseBodyUniversitySecretKeyInfo) SetSecretKey(v string) *QueryOrgSecretKeyResponseBodyUniversitySecretKeyInfo {
	s.SecretKey = &v
	return s
}

type QueryOrgSecretKeyResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryOrgSecretKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOrgSecretKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgSecretKeyResponse) GoString() string {
	return s.String()
}

func (s *QueryOrgSecretKeyResponse) SetHeaders(v map[string]*string) *QueryOrgSecretKeyResponse {
	s.Headers = v
	return s
}

func (s *QueryOrgSecretKeyResponse) SetBody(v *QueryOrgSecretKeyResponseBody) *QueryOrgSecretKeyResponse {
	s.Body = v
	return s
}

type SearchTeachersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchTeachersHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchTeachersHeaders) GoString() string {
	return s.String()
}

func (s *SearchTeachersHeaders) SetCommonHeaders(v map[string]*string) *SearchTeachersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchTeachersHeaders) SetXAcsDingtalkAccessToken(v string) *SearchTeachersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchTeachersRequest struct {
	NameKeyword *string `json:"nameKeyword,omitempty" xml:"nameKeyword,omitempty"`
}

func (s SearchTeachersRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTeachersRequest) GoString() string {
	return s.String()
}

func (s *SearchTeachersRequest) SetNameKeyword(v string) *SearchTeachersRequest {
	s.NameKeyword = &v
	return s
}

type SearchTeachersResponseBody struct {
	Users []*SearchTeachersResponseBodyUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s SearchTeachersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTeachersResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTeachersResponseBody) SetUsers(v []*SearchTeachersResponseBodyUsers) *SearchTeachersResponseBody {
	s.Users = v
	return s
}

type SearchTeachersResponseBodyUsers struct {
	// 用户ID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 所在其中一个班级ID
	ClassId *int64 `json:"classId,omitempty" xml:"classId,omitempty"`
	// 所在其中一个班级名称
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
}

func (s SearchTeachersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s SearchTeachersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *SearchTeachersResponseBodyUsers) SetUserId(v string) *SearchTeachersResponseBodyUsers {
	s.UserId = &v
	return s
}

func (s *SearchTeachersResponseBodyUsers) SetName(v string) *SearchTeachersResponseBodyUsers {
	s.Name = &v
	return s
}

func (s *SearchTeachersResponseBodyUsers) SetClassId(v int64) *SearchTeachersResponseBodyUsers {
	s.ClassId = &v
	return s
}

func (s *SearchTeachersResponseBodyUsers) SetDeptName(v string) *SearchTeachersResponseBodyUsers {
	s.DeptName = &v
	return s
}

type SearchTeachersResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchTeachersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchTeachersResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTeachersResponse) GoString() string {
	return s.String()
}

func (s *SearchTeachersResponse) SetHeaders(v map[string]*string) *SearchTeachersResponse {
	s.Headers = v
	return s
}

func (s *SearchTeachersResponse) SetBody(v *SearchTeachersResponseBody) *SearchTeachersResponse {
	s.Body = v
	return s
}

type DeletePhysicalClassroomHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeletePhysicalClassroomHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeletePhysicalClassroomHeaders) GoString() string {
	return s.String()
}

func (s *DeletePhysicalClassroomHeaders) SetCommonHeaders(v map[string]*string) *DeletePhysicalClassroomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeletePhysicalClassroomHeaders) SetXAcsDingtalkAccessToken(v string) *DeletePhysicalClassroomHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeletePhysicalClassroomRequest struct {
	// 操作人id
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 教室主键
	ClassroomId *int64 `json:"classroomId,omitempty" xml:"classroomId,omitempty"`
}

func (s DeletePhysicalClassroomRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePhysicalClassroomRequest) GoString() string {
	return s.String()
}

func (s *DeletePhysicalClassroomRequest) SetOpUserId(v string) *DeletePhysicalClassroomRequest {
	s.OpUserId = &v
	return s
}

func (s *DeletePhysicalClassroomRequest) SetClassroomId(v int64) *DeletePhysicalClassroomRequest {
	s.ClassroomId = &v
	return s
}

type DeletePhysicalClassroomResponseBody struct {
	// 是否成功
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeletePhysicalClassroomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePhysicalClassroomResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePhysicalClassroomResponseBody) SetResult(v bool) *DeletePhysicalClassroomResponseBody {
	s.Result = &v
	return s
}

type DeletePhysicalClassroomResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeletePhysicalClassroomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePhysicalClassroomResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePhysicalClassroomResponse) GoString() string {
	return s.String()
}

func (s *DeletePhysicalClassroomResponse) SetHeaders(v map[string]*string) *DeletePhysicalClassroomResponse {
	s.Headers = v
	return s
}

func (s *DeletePhysicalClassroomResponse) SetBody(v *DeletePhysicalClassroomResponseBody) *DeletePhysicalClassroomResponse {
	s.Body = v
	return s
}

type QueryOrgRelationListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOrgRelationListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgRelationListHeaders) GoString() string {
	return s.String()
}

func (s *QueryOrgRelationListHeaders) SetCommonHeaders(v map[string]*string) *QueryOrgRelationListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOrgRelationListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOrgRelationListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOrgRelationListRequest struct {
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
}

func (s QueryOrgRelationListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgRelationListRequest) GoString() string {
	return s.String()
}

func (s *QueryOrgRelationListRequest) SetOperator(v string) *QueryOrgRelationListRequest {
	s.Operator = &v
	return s
}

type QueryOrgRelationListResponseBody struct {
	// Id of the request
	Success *bool                                     `json:"success,omitempty" xml:"success,omitempty"`
	Result  []*QueryOrgRelationListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryOrgRelationListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgRelationListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrgRelationListResponseBody) SetSuccess(v bool) *QueryOrgRelationListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryOrgRelationListResponseBody) SetResult(v []*QueryOrgRelationListResponseBodyResult) *QueryOrgRelationListResponseBody {
	s.Result = v
	return s
}

type QueryOrgRelationListResponseBodyResult struct {
	CorpId      *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	DeviceCount *int32  `json:"deviceCount,omitempty" xml:"deviceCount,omitempty"`
}

func (s QueryOrgRelationListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgRelationListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryOrgRelationListResponseBodyResult) SetCorpId(v string) *QueryOrgRelationListResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *QueryOrgRelationListResponseBodyResult) SetName(v string) *QueryOrgRelationListResponseBodyResult {
	s.Name = &v
	return s
}

func (s *QueryOrgRelationListResponseBodyResult) SetDeviceCount(v int32) *QueryOrgRelationListResponseBodyResult {
	s.DeviceCount = &v
	return s
}

type QueryOrgRelationListResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryOrgRelationListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOrgRelationListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgRelationListResponse) GoString() string {
	return s.String()
}

func (s *QueryOrgRelationListResponse) SetHeaders(v map[string]*string) *QueryOrgRelationListResponse {
	s.Headers = v
	return s
}

func (s *QueryOrgRelationListResponse) SetBody(v *QueryOrgRelationListResponseBody) *QueryOrgRelationListResponse {
	s.Body = v
	return s
}

type QueryOrgTypeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOrgTypeHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTypeHeaders) GoString() string {
	return s.String()
}

func (s *QueryOrgTypeHeaders) SetCommonHeaders(v map[string]*string) *QueryOrgTypeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOrgTypeHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOrgTypeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOrgTypeResponseBody struct {
	// 组织类型
	OrgType *int64 `json:"orgType,omitempty" xml:"orgType,omitempty"`
}

func (s QueryOrgTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTypeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrgTypeResponseBody) SetOrgType(v int64) *QueryOrgTypeResponseBody {
	s.OrgType = &v
	return s
}

type QueryOrgTypeResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryOrgTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOrgTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTypeResponse) GoString() string {
	return s.String()
}

func (s *QueryOrgTypeResponse) SetHeaders(v map[string]*string) *QueryOrgTypeResponse {
	s.Headers = v
	return s
}

func (s *QueryOrgTypeResponse) SetBody(v *QueryOrgTypeResponseBody) *QueryOrgTypeResponse {
	s.Body = v
	return s
}

type EndCourseHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EndCourseHeaders) String() string {
	return tea.Prettify(s)
}

func (s EndCourseHeaders) GoString() string {
	return s.String()
}

func (s *EndCourseHeaders) SetCommonHeaders(v map[string]*string) *EndCourseHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EndCourseHeaders) SetXAcsDingtalkAccessToken(v string) *EndCourseHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EndCourseRequest struct {
	// 用户id
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 拓展字段
	Ext *string `json:"ext,omitempty" xml:"ext,omitempty"`
	// 课程编码
	CourseCode *string `json:"courseCode,omitempty" xml:"courseCode,omitempty"`
	// 直播流信息
	LivePlayInfoList []*EndCourseRequestLivePlayInfoList `json:"livePlayInfoList,omitempty" xml:"livePlayInfoList,omitempty" type:"Repeated"`
	// isv编码
	IsvCode *string `json:"isvCode,omitempty" xml:"isvCode,omitempty"`
}

func (s EndCourseRequest) String() string {
	return tea.Prettify(s)
}

func (s EndCourseRequest) GoString() string {
	return s.String()
}

func (s *EndCourseRequest) SetOpUserId(v string) *EndCourseRequest {
	s.OpUserId = &v
	return s
}

func (s *EndCourseRequest) SetExt(v string) *EndCourseRequest {
	s.Ext = &v
	return s
}

func (s *EndCourseRequest) SetCourseCode(v string) *EndCourseRequest {
	s.CourseCode = &v
	return s
}

func (s *EndCourseRequest) SetLivePlayInfoList(v []*EndCourseRequestLivePlayInfoList) *EndCourseRequest {
	s.LivePlayInfoList = v
	return s
}

func (s *EndCourseRequest) SetIsvCode(v string) *EndCourseRequest {
	s.IsvCode = &v
	return s
}

type EndCourseRequestLivePlayInfoList struct {
	// 直播拉流地址
	LiveOutputUrl *string `json:"liveOutputUrl,omitempty" xml:"liveOutputUrl,omitempty"`
	// 直播流类型
	LiveType *int32 `json:"liveType,omitempty" xml:"liveType,omitempty"`
	// 直播推流地址
	LiveInputUrl *string `json:"liveInputUrl,omitempty" xml:"liveInputUrl,omitempty"`
	// 回放视频地址
	ReplayUrl *string `json:"replayUrl,omitempty" xml:"replayUrl,omitempty"`
}

func (s EndCourseRequestLivePlayInfoList) String() string {
	return tea.Prettify(s)
}

func (s EndCourseRequestLivePlayInfoList) GoString() string {
	return s.String()
}

func (s *EndCourseRequestLivePlayInfoList) SetLiveOutputUrl(v string) *EndCourseRequestLivePlayInfoList {
	s.LiveOutputUrl = &v
	return s
}

func (s *EndCourseRequestLivePlayInfoList) SetLiveType(v int32) *EndCourseRequestLivePlayInfoList {
	s.LiveType = &v
	return s
}

func (s *EndCourseRequestLivePlayInfoList) SetLiveInputUrl(v string) *EndCourseRequestLivePlayInfoList {
	s.LiveInputUrl = &v
	return s
}

func (s *EndCourseRequestLivePlayInfoList) SetReplayUrl(v string) *EndCourseRequestLivePlayInfoList {
	s.ReplayUrl = &v
	return s
}

type EndCourseResponseBody struct {
	UniversityCourseCommonResponse *EndCourseResponseBodyUniversityCourseCommonResponse `json:"universityCourseCommonResponse,omitempty" xml:"universityCourseCommonResponse,omitempty" type:"Struct"`
}

func (s EndCourseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EndCourseResponseBody) GoString() string {
	return s.String()
}

func (s *EndCourseResponseBody) SetUniversityCourseCommonResponse(v *EndCourseResponseBodyUniversityCourseCommonResponse) *EndCourseResponseBody {
	s.UniversityCourseCommonResponse = v
	return s
}

type EndCourseResponseBodyUniversityCourseCommonResponse struct {
	// 调用是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 课程编码
	CourseCode *string `json:"courseCode,omitempty" xml:"courseCode,omitempty"`
}

func (s EndCourseResponseBodyUniversityCourseCommonResponse) String() string {
	return tea.Prettify(s)
}

func (s EndCourseResponseBodyUniversityCourseCommonResponse) GoString() string {
	return s.String()
}

func (s *EndCourseResponseBodyUniversityCourseCommonResponse) SetSuccess(v bool) *EndCourseResponseBodyUniversityCourseCommonResponse {
	s.Success = &v
	return s
}

func (s *EndCourseResponseBodyUniversityCourseCommonResponse) SetCourseCode(v string) *EndCourseResponseBodyUniversityCourseCommonResponse {
	s.CourseCode = &v
	return s
}

type EndCourseResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EndCourseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EndCourseResponse) String() string {
	return tea.Prettify(s)
}

func (s EndCourseResponse) GoString() string {
	return s.String()
}

func (s *EndCourseResponse) SetHeaders(v map[string]*string) *EndCourseResponse {
	s.Headers = v
	return s
}

func (s *EndCourseResponse) SetBody(v *EndCourseResponseBody) *EndCourseResponse {
	s.Body = v
	return s
}

type GetShareRoleMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetShareRoleMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetShareRoleMembersHeaders) GoString() string {
	return s.String()
}

func (s *GetShareRoleMembersHeaders) SetCommonHeaders(v map[string]*string) *GetShareRoleMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetShareRoleMembersHeaders) SetXAcsDingtalkAccessToken(v string) *GetShareRoleMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetShareRoleMembersResponseBody struct {
	Result []*GetShareRoleMembersResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetShareRoleMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetShareRoleMembersResponseBody) GoString() string {
	return s.String()
}

func (s *GetShareRoleMembersResponseBody) SetResult(v []*GetShareRoleMembersResponseBodyResult) *GetShareRoleMembersResponseBody {
	s.Result = v
	return s
}

type GetShareRoleMembersResponseBodyResult struct {
	// 分支组织corpId
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 角色成员在主干组织中的userId列表
	MemberUserIdListInTrunkOrg []*string `json:"memberUserIdListInTrunkOrg,omitempty" xml:"memberUserIdListInTrunkOrg,omitempty" type:"Repeated"`
}

func (s GetShareRoleMembersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetShareRoleMembersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetShareRoleMembersResponseBodyResult) SetCorpId(v string) *GetShareRoleMembersResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *GetShareRoleMembersResponseBodyResult) SetMemberUserIdListInTrunkOrg(v []*string) *GetShareRoleMembersResponseBodyResult {
	s.MemberUserIdListInTrunkOrg = v
	return s
}

type GetShareRoleMembersResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetShareRoleMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetShareRoleMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s GetShareRoleMembersResponse) GoString() string {
	return s.String()
}

func (s *GetShareRoleMembersResponse) SetHeaders(v map[string]*string) *GetShareRoleMembersResponse {
	s.Headers = v
	return s
}

func (s *GetShareRoleMembersResponse) SetBody(v *GetShareRoleMembersResponseBody) *GetShareRoleMembersResponse {
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

func (client *Client) QueryStatisticsData(request *QueryStatisticsDataRequest) (_result *QueryStatisticsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryStatisticsDataHeaders{}
	_result = &QueryStatisticsDataResponse{}
	_body, _err := client.QueryStatisticsDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryStatisticsDataWithOptions(request *QueryStatisticsDataRequest, headers *QueryStatisticsDataHeaders, runtime *util.RuntimeOptions) (_result *QueryStatisticsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SectionIndexList)) {
		body["sectionIndexList"] = request.SectionIndexList
	}

	if !tea.BoolValue(util.IsUnset(request.TeacherUserIds)) {
		body["teacherUserIds"] = request.TeacherUserIds
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
	_result = &QueryStatisticsDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryStatisticsData"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/edu/classes/schedules/statisticData/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryClassScheduleConfig(request *QueryClassScheduleConfigRequest) (_result *QueryClassScheduleConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryClassScheduleConfigHeaders{}
	_result = &QueryClassScheduleConfigResponse{}
	_body, _err := client.QueryClassScheduleConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryClassScheduleConfigWithOptions(tmpReq *QueryClassScheduleConfigRequest, headers *QueryClassScheduleConfigHeaders, runtime *util.RuntimeOptions) (_result *QueryClassScheduleConfigResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryClassScheduleConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ClassIds)) {
		request.ClassIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClassIds, tea.String("classIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassIdsShrink)) {
		query["classIds"] = request.ClassIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
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
	_result = &QueryClassScheduleConfigResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryClassScheduleConfig"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/edu/schedules/configs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDefaultChild() (_result *GetDefaultChildResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDefaultChildHeaders{}
	_result = &GetDefaultChildResponse{}
	_body, _err := client.GetDefaultChildWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDefaultChildWithOptions(headers *GetDefaultChildHeaders, runtime *util.RuntimeOptions) (_result *GetDefaultChildResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &GetDefaultChildResponse{}
	_body, _err := client.DoROARequest(tea.String("GetDefaultChild"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/edu/defaultChildren"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOpenCourses(request *GetOpenCoursesRequest) (_result *GetOpenCoursesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOpenCoursesHeaders{}
	_result = &GetOpenCoursesResponse{}
	_body, _err := client.GetOpenCoursesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOpenCoursesWithOptions(request *GetOpenCoursesRequest, headers *GetOpenCoursesHeaders, runtime *util.RuntimeOptions) (_result *GetOpenCoursesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
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
	_result = &GetOpenCoursesResponse{}
	_body, _err := client.DoROARequest(tea.String("GetOpenCourses"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/edu/openCourses"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchCreate(request *BatchCreateRequest) (_result *BatchCreateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchCreateHeaders{}
	_result = &BatchCreateResponse{}
	_body, _err := client.BatchCreateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchCreateWithOptions(request *BatchCreateRequest, headers *BatchCreateHeaders, runtime *util.RuntimeOptions) (_result *BatchCreateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CardBizCode)) {
		body["cardBizCode"] = request.CardBizCode
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Data))) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.Identifier)) {
		body["identifier"] = request.Identifier
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		body["sourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Userid)) {
		body["userid"] = request.Userid
	}

	if !tea.BoolValue(util.IsUnset(request.DingCorpId)) {
		body["dingCorpId"] = request.DingCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.JsVersion)) {
		body["jsVersion"] = request.JsVersion
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
	_result = &BatchCreateResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchCreate"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/edu/cards"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitCoursesOfClass(classId *string, request *InitCoursesOfClassRequest) (_result *InitCoursesOfClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InitCoursesOfClassHeaders{}
	_result = &InitCoursesOfClassResponse{}
	_body, _err := client.InitCoursesOfClassWithOptions(classId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitCoursesOfClassWithOptions(classId *string, request *InitCoursesOfClassRequest, headers *InitCoursesOfClassHeaders, runtime *util.RuntimeOptions) (_result *InitCoursesOfClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Courses)) {
		body["courses"] = request.Courses
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SectionConfig))) {
		body["sectionConfig"] = request.SectionConfig
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
	_result = &InitCoursesOfClassResponse{}
	_body, _err := client.DoROARequest(tea.String("InitCoursesOfClass"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/edu/classes/"+tea.StringValue(classId)+"/courses/init"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInviteUrl(request *CreateInviteUrlRequest) (_result *CreateInviteUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateInviteUrlHeaders{}
	_result = &CreateInviteUrlResponse{}
	_body, _err := client.CreateInviteUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInviteUrlWithOptions(request *CreateInviteUrlRequest, headers *CreateInviteUrlHeaders, runtime *util.RuntimeOptions) (_result *CreateInviteUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingCorpId)) {
		body["dingCorpId"] = request.DingCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.DingOauthAppId)) {
		body["dingOauthAppId"] = request.DingOauthAppId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		body["targetCorpId"] = request.TargetCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		body["authCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.TargetOperator)) {
		body["targetOperator"] = request.TargetOperator
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
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
	_result = &CreateInviteUrlResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateInviteUrl"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/edu/remoteClasses/orgRelations/inviteUrls"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PollingConfirmStatus(request *PollingConfirmStatusRequest) (_result *PollingConfirmStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PollingConfirmStatusHeaders{}
	_result = &PollingConfirmStatusResponse{}
	_body, _err := client.PollingConfirmStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PollingConfirmStatusWithOptions(request *PollingConfirmStatusRequest, headers *PollingConfirmStatusHeaders, runtime *util.RuntimeOptions) (_result *PollingConfirmStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		query["ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.IsvCode)) {
		query["isvCode"] = request.IsvCode
	}

	if !tea.BoolValue(util.IsUnset(request.CourseCode)) {
		query["courseCode"] = request.CourseCode
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
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
	_result = &PollingConfirmStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("PollingConfirmStatus"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/edu/universities/courses/pollingConfirmStatus"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTeacher(classId *string, userId *string, request *DeleteTeacherRequest) (_result *DeleteTeacherResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteTeacherHeaders{}
	_result = &DeleteTeacherResponse{}
	_body, _err := client.DeleteTeacherWithOptions(classId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTeacherWithOptions(classId *string, userId *string, request *DeleteTeacherRequest, headers *DeleteTeacherHeaders, runtime *util.RuntimeOptions) (_result *DeleteTeacherResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Adviser)) {
		query["adviser"] = request.Adviser
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		query["operator"] = request.Operator
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
	_result = &DeleteTeacherResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteTeacher"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/edu/classes/"+tea.StringValue(classId)+"/teachers/"+tea.StringValue(userId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDeviceOrg(request *DeleteDeviceOrgRequest) (_result *DeleteDeviceOrgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteDeviceOrgHeaders{}
	_result = &DeleteDeviceOrgResponse{}
	_body, _err := client.DeleteDeviceOrgWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDeviceOrgWithOptions(request *DeleteDeviceOrgRequest, headers *DeleteDeviceOrgHeaders, runtime *util.RuntimeOptions) (_result *DeleteDeviceOrgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceCode)) {
		query["deviceCode"] = request.DeviceCode
	}

	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["authCode"] = request.AuthCode
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
	_result = &DeleteDeviceOrgResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteDeviceOrg"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/edu/remoteClasses/deviceOrgs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRemoteClassCourse(courseCode *string, request *DeleteRemoteClassCourseRequest) (_result *DeleteRemoteClassCourseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteRemoteClassCourseHeaders{}
	_result = &DeleteRemoteClassCourseResponse{}
	_body, _err := client.DeleteRemoteClassCourseWithOptions(courseCode, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRemoteClassCourseWithOptions(courseCode *string, request *DeleteRemoteClassCourseRequest, headers *DeleteRemoteClassCourseHeaders, runtime *util.RuntimeOptions) (_result *DeleteRemoteClassCourseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["authCode"] = request.AuthCode
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
	_result = &DeleteRemoteClassCourseResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteRemoteClassCourse"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/edu/remoteClasses/courses/"+tea.StringValue(courseCode)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOpenCourseDetail(courseId *string) (_result *GetOpenCourseDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOpenCourseDetailHeaders{}
	_result = &GetOpenCourseDetailResponse{}
	_body, _err := client.GetOpenCourseDetailWithOptions(courseId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOpenCourseDetailWithOptions(courseId *string, headers *GetOpenCourseDetailHeaders, runtime *util.RuntimeOptions) (_result *GetOpenCourseDetailResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &GetOpenCourseDetailResponse{}
	_body, _err := client.DoROARequest(tea.String("GetOpenCourseDetail"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/edu/openCourse/"+tea.StringValue(courseId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteStudent(classId *string, userId *string, request *DeleteStudentRequest) (_result *DeleteStudentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteStudentHeaders{}
	_result = &DeleteStudentResponse{}
	_body, _err := client.DeleteStudentWithOptions(classId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteStudentWithOptions(classId *string, userId *string, request *DeleteStudentRequest, headers *DeleteStudentHeaders, runtime *util.RuntimeOptions) (_result *DeleteStudentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		query["operator"] = request.Operator
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
	_result = &DeleteStudentResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteStudent"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/edu/classes/"+tea.StringValue(classId)+"/students/"+tea.StringValue(userId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCoursesOfClass(classId *string, request *UpdateCoursesOfClassRequest) (_result *UpdateCoursesOfClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateCoursesOfClassHeaders{}
	_result = &UpdateCoursesOfClassResponse{}
	_body, _err := client.UpdateCoursesOfClassWithOptions(classId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCoursesOfClassWithOptions(classId *string, request *UpdateCoursesOfClassRequest, headers *UpdateCoursesOfClassHeaders, runtime *util.RuntimeOptions) (_result *UpdateCoursesOfClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Courses)) {
		body["courses"] = request.Courses
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SectionConfig))) {
		body["sectionConfig"] = request.SectionConfig
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
	_result = &UpdateCoursesOfClassResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateCoursesOfClass"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/edu/classes/"+tea.StringValue(classId)+"/courses/schedules"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSectionConfig(request *CreateSectionConfigRequest) (_result *CreateSectionConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateSectionConfigHeaders{}
	_result = &CreateSectionConfigResponse{}
	_body, _err := client.CreateSectionConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSectionConfigWithOptions(request *CreateSectionConfigRequest, headers *CreateSectionConfigHeaders, runtime *util.RuntimeOptions) (_result *CreateSectionConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.SectionConfigs)) {
		body["sectionConfigs"] = request.SectionConfigs
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
	_result = &CreateSectionConfigResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateSectionConfig"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/edu/universities/sectionConfigs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySubjectTeachers(request *QuerySubjectTeachersRequest) (_result *QuerySubjectTeachersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QuerySubjectTeachersHeaders{}
	_result = &QuerySubjectTeachersResponse{}
	_body, _err := client.QuerySubjectTeachersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySubjectTeachersWithOptions(request *QuerySubjectTeachersRequest, headers *QuerySubjectTeachersHeaders, runtime *util.RuntimeOptions) (_result *QuerySubjectTeachersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassIds)) {
		query["classIds"] = request.ClassIds
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.SubjectCode)) {
		query["subjectCode"] = request.SubjectCode
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
	_result = &QuerySubjectTeachersResponse{}
	_body, _err := client.DoROARequest(tea.String("QuerySubjectTeachers"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/edu/subjects/teachers"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRemoteClassCourse(request *QueryRemoteClassCourseRequest) (_result *QueryRemoteClassCourseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryRemoteClassCourseHeaders{}
	_result = &QueryRemoteClassCourseResponse{}
	_body, _err := client.QueryRemoteClassCourseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRemoteClassCourseWithOptions(request *QueryRemoteClassCourseRequest, headers *QueryRemoteClassCourseHeaders, runtime *util.RuntimeOptions) (_result *QueryRemoteClassCourseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		query["operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
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
	_result = &QueryRemoteClassCourseResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryRemoteClassCourse"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/edu/remoteClasses/courses"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePhysicalClassroom(request *CreatePhysicalClassroomRequest) (_result *CreatePhysicalClassroomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreatePhysicalClassroomHeaders{}
	_result = &CreatePhysicalClassroomResponse{}
	_body, _err := client.CreatePhysicalClassroomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePhysicalClassroomWithOptions(request *CreatePhysicalClassroomRequest, headers *CreatePhysicalClassroomHeaders, runtime *util.RuntimeOptions) (_result *CreatePhysicalClassroomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassroomFloor)) {
		body["classroomFloor"] = request.ClassroomFloor
	}

	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.ClassroomBuilding)) {
		body["classroomBuilding"] = request.ClassroomBuilding
	}

	if !tea.BoolValue(util.IsUnset(request.DirectBroadcast)) {
		body["directBroadcast"] = request.DirectBroadcast
	}

	if !tea.BoolValue(util.IsUnset(request.ClassroomName)) {
		body["classroomName"] = request.ClassroomName
	}

	if !tea.BoolValue(util.IsUnset(request.ClassroomCampus)) {
		body["classroomCampus"] = request.ClassroomCampus
	}

	if !tea.BoolValue(util.IsUnset(request.ClassroomNumber)) {
		body["classroomNumber"] = request.ClassroomNumber
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
	_result = &CreatePhysicalClassroomResponse{}
	_body, _err := client.DoROARequest(tea.String("CreatePhysicalClassroom"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/edu/physicalClassrooms"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryClassSchedule(request *QueryClassScheduleRequest) (_result *QueryClassScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryClassScheduleHeaders{}
	_result = &QueryClassScheduleResponse{}
	_body, _err := client.QueryClassScheduleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryClassScheduleWithOptions(request *QueryClassScheduleRequest, headers *QueryClassScheduleHeaders, runtime *util.RuntimeOptions) (_result *QueryClassScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubscriberType)) {
		query["subscriberType"] = request.SubscriberType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubscriberIds)) {
		body["subscriberIds"] = request.SubscriberIds
	}

	if !tea.BoolValue(util.IsUnset(request.SectionIndexList)) {
		body["sectionIndexList"] = request.SectionIndexList
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
	_result = &QueryClassScheduleResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryClassSchedule"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/edu/classes/schedules/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRemoteClassCourse(request *CreateRemoteClassCourseRequest) (_result *CreateRemoteClassCourseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateRemoteClassCourseHeaders{}
	_result = &CreateRemoteClassCourseResponse{}
	_body, _err := client.CreateRemoteClassCourseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRemoteClassCourseWithOptions(request *CreateRemoteClassCourseRequest, headers *CreateRemoteClassCourseHeaders, runtime *util.RuntimeOptions) (_result *CreateRemoteClassCourseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CourseName)) {
		body["courseName"] = request.CourseName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TeachingParticipant))) {
		body["teachingParticipant"] = request.TeachingParticipant
	}

	if !tea.BoolValue(util.IsUnset(request.AttendParticipants)) {
		body["attendParticipants"] = request.AttendParticipants
	}

	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		body["authCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.DingClientId)) {
		body["dingClientId"] = request.DingClientId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.DingOauthAppId)) {
		body["dingOauthAppId"] = request.DingOauthAppId
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingCorpId)) {
		body["dingCorpId"] = request.DingCorpId
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
	_result = &CreateRemoteClassCourseResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateRemoteClassCourse"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/edu/remoteClasses/courses"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCustomClass(request *CreateCustomClassRequest) (_result *CreateCustomClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateCustomClassHeaders{}
	_result = &CreateCustomClassResponse{}
	_body, _err := client.CreateCustomClassWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCustomClassWithOptions(request *CreateCustomClassRequest, headers *CreateCustomClassHeaders, runtime *util.RuntimeOptions) (_result *CreateCustomClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CustomClass))) {
		body["customClass"] = request.CustomClass
	}

	if !tea.BoolValue(util.IsUnset(request.SuperId)) {
		body["superId"] = request.SuperId
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingCorpId)) {
		body["dingCorpId"] = request.DingCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOauthAppId)) {
		body["dingOauthAppId"] = request.DingOauthAppId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
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
	_result = &CreateCustomClassResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateCustomClass"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/edu/customClasses"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePhysicalClassroom(request *UpdatePhysicalClassroomRequest) (_result *UpdatePhysicalClassroomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdatePhysicalClassroomHeaders{}
	_result = &UpdatePhysicalClassroomResponse{}
	_body, _err := client.UpdatePhysicalClassroomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePhysicalClassroomWithOptions(request *UpdatePhysicalClassroomRequest, headers *UpdatePhysicalClassroomHeaders, runtime *util.RuntimeOptions) (_result *UpdatePhysicalClassroomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassroomFloor)) {
		body["classroomFloor"] = request.ClassroomFloor
	}

	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.ClassroomBuilding)) {
		body["classroomBuilding"] = request.ClassroomBuilding
	}

	if !tea.BoolValue(util.IsUnset(request.DirectBroadcast)) {
		body["directBroadcast"] = request.DirectBroadcast
	}

	if !tea.BoolValue(util.IsUnset(request.ClassroomName)) {
		body["classroomName"] = request.ClassroomName
	}

	if !tea.BoolValue(util.IsUnset(request.ClassroomCampus)) {
		body["classroomCampus"] = request.ClassroomCampus
	}

	if !tea.BoolValue(util.IsUnset(request.ClassroomNumber)) {
		body["classroomNumber"] = request.ClassroomNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ClassroomId)) {
		body["classroomId"] = request.ClassroomId
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
	_result = &UpdatePhysicalClassroomResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdatePhysicalClassroom"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/edu/physicalClassrooms"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveStudent(request *MoveStudentRequest) (_result *MoveStudentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MoveStudentHeaders{}
	_result = &MoveStudentResponse{}
	_body, _err := client.MoveStudentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveStudentWithOptions(request *MoveStudentRequest, headers *MoveStudentHeaders, runtime *util.RuntimeOptions) (_result *MoveStudentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingCorpId)) {
		body["dingCorpId"] = request.DingCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOauthAppId)) {
		body["dingOauthAppId"] = request.DingOauthAppId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.OriginClassId)) {
		body["originClassId"] = request.OriginClassId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetClassId)) {
		body["targetClassId"] = request.TargetClassId
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
	_result = &MoveStudentResponse{}
	_body, _err := client.DoROARequest(tea.String("MoveStudent"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/edu/students/move"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUniversityCourseGroup(request *CreateUniversityCourseGroupRequest) (_result *CreateUniversityCourseGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateUniversityCourseGroupHeaders{}
	_result = &CreateUniversityCourseGroupResponse{}
	_body, _err := client.CreateUniversityCourseGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUniversityCourseGroupWithOptions(request *CreateUniversityCourseGroupRequest, headers *CreateUniversityCourseGroupHeaders, runtime *util.RuntimeOptions) (_result *CreateUniversityCourseGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsvCourseGroupCode)) {
		body["isvCourseGroupCode"] = request.IsvCourseGroupCode
	}

	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.CourseGroupIntroduce)) {
		body["courseGroupIntroduce"] = request.CourseGroupIntroduce
	}

	if !tea.BoolValue(util.IsUnset(request.Semester)) {
		body["semester"] = request.Semester
	}

	if !tea.BoolValue(util.IsUnset(request.SubjectName)) {
		body["subjectName"] = request.SubjectName
	}

	if !tea.BoolValue(util.IsUnset(request.CourseGroupName)) {
		body["courseGroupName"] = request.CourseGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.SchoolYear)) {
		body["schoolYear"] = request.SchoolYear
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodCode)) {
		body["periodCode"] = request.PeriodCode
	}

	if !tea.BoolValue(util.IsUnset(request.TeacherInfos)) {
		body["teacherInfos"] = request.TeacherInfos
	}

	if !tea.BoolValue(util.IsUnset(request.CourserGroupItemModels)) {
		body["courserGroupItemModels"] = request.CourserGroupItemModels
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
	_result = &CreateUniversityCourseGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateUniversityCourseGroup"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/edu/universities/courseGroups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRemoteClassCourse(courseCode *string, request *GetRemoteClassCourseRequest) (_result *GetRemoteClassCourseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetRemoteClassCourseHeaders{}
	_result = &GetRemoteClassCourseResponse{}
	_body, _err := client.GetRemoteClassCourseWithOptions(courseCode, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRemoteClassCourseWithOptions(courseCode *string, request *GetRemoteClassCourseRequest, headers *GetRemoteClassCourseHeaders, runtime *util.RuntimeOptions) (_result *GetRemoteClassCourseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		query["operator"] = request.Operator
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
	_result = &GetRemoteClassCourseResponse{}
	_body, _err := client.DoROARequest(tea.String("GetRemoteClassCourse"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/edu/remoteClasses/courses/"+tea.StringValue(courseCode)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUniversityCourseGroup(request *QueryUniversityCourseGroupRequest) (_result *QueryUniversityCourseGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUniversityCourseGroupHeaders{}
	_result = &QueryUniversityCourseGroupResponse{}
	_body, _err := client.QueryUniversityCourseGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUniversityCourseGroupWithOptions(request *QueryUniversityCourseGroupRequest, headers *QueryUniversityCourseGroupHeaders, runtime *util.RuntimeOptions) (_result *QueryUniversityCourseGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.CourseGroupCode)) {
		query["courseGroupCode"] = request.CourseGroupCode
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
	_result = &QueryUniversityCourseGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryUniversityCourseGroup"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/edu/universities/courseGroups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllSubjectsFromClassSchedule(request *QueryAllSubjectsFromClassScheduleRequest) (_result *QueryAllSubjectsFromClassScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllSubjectsFromClassScheduleHeaders{}
	_result = &QueryAllSubjectsFromClassScheduleResponse{}
	_body, _err := client.QueryAllSubjectsFromClassScheduleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllSubjectsFromClassScheduleWithOptions(tmpReq *QueryAllSubjectsFromClassScheduleRequest, headers *QueryAllSubjectsFromClassScheduleHeaders, runtime *util.RuntimeOptions) (_result *QueryAllSubjectsFromClassScheduleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryAllSubjectsFromClassScheduleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ClassIds)) {
		request.ClassIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClassIds, tea.String("classIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassIdsShrink)) {
		query["classIds"] = request.ClassIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodCode)) {
		query["periodCode"] = request.PeriodCode
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
	_result = &QueryAllSubjectsFromClassScheduleResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAllSubjectsFromClassSchedule"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/edu/subjects/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CourseSchedulingComplimentNotice(request *CourseSchedulingComplimentNoticeRequest) (_result *CourseSchedulingComplimentNoticeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CourseSchedulingComplimentNoticeHeaders{}
	_result = &CourseSchedulingComplimentNoticeResponse{}
	_body, _err := client.CourseSchedulingComplimentNoticeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CourseSchedulingComplimentNoticeWithOptions(request *CourseSchedulingComplimentNoticeRequest, headers *CourseSchedulingComplimentNoticeHeaders, runtime *util.RuntimeOptions) (_result *CourseSchedulingComplimentNoticeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
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
	_result = &CourseSchedulingComplimentNoticeResponse{}
	_body, _err := client.DoROARequest(tea.String("CourseSchedulingComplimentNotice"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/edu/schedules/finishNotify"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUniversityCourseGroup(request *DeleteUniversityCourseGroupRequest) (_result *DeleteUniversityCourseGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteUniversityCourseGroupHeaders{}
	_result = &DeleteUniversityCourseGroupResponse{}
	_body, _err := client.DeleteUniversityCourseGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUniversityCourseGroupWithOptions(request *DeleteUniversityCourseGroupRequest, headers *DeleteUniversityCourseGroupHeaders, runtime *util.RuntimeOptions) (_result *DeleteUniversityCourseGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.CourseGroupCode)) {
		query["courseGroupCode"] = request.CourseGroupCode
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
	_result = &DeleteUniversityCourseGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteUniversityCourseGroup"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/edu/universities/courseGroups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDept(deptId *string, request *DeleteDeptRequest) (_result *DeleteDeptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteDeptHeaders{}
	_result = &DeleteDeptResponse{}
	_body, _err := client.DeleteDeptWithOptions(deptId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDeptWithOptions(deptId *string, request *DeleteDeptRequest, headers *DeleteDeptHeaders, runtime *util.RuntimeOptions) (_result *DeleteDeptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		query["operator"] = request.Operator
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
	_result = &DeleteDeptResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteDept"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/edu/depts/"+tea.StringValue(deptId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGuardian(classId *string, userId *string, request *DeleteGuardianRequest) (_result *DeleteGuardianResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteGuardianHeaders{}
	_result = &DeleteGuardianResponse{}
	_body, _err := client.DeleteGuardianWithOptions(classId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGuardianWithOptions(classId *string, userId *string, request *DeleteGuardianRequest, headers *DeleteGuardianHeaders, runtime *util.RuntimeOptions) (_result *DeleteGuardianResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StuId)) {
		query["stuId"] = request.StuId
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		query["operator"] = request.Operator
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
	_result = &DeleteGuardianResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteGuardian"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/edu/classes/"+tea.StringValue(classId)+"/guardians/"+tea.StringValue(userId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertSectionConfig(request *InsertSectionConfigRequest) (_result *InsertSectionConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InsertSectionConfigHeaders{}
	_result = &InsertSectionConfigResponse{}
	_body, _err := client.InsertSectionConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertSectionConfigWithOptions(request *InsertSectionConfigRequest, headers *InsertSectionConfigHeaders, runtime *util.RuntimeOptions) (_result *InsertSectionConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SectionModels)) {
		body["sectionModels"] = request.SectionModels
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Start))) {
		body["start"] = request.Start
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.End))) {
		body["end"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleName)) {
		body["scheduleName"] = request.ScheduleName
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
	_result = &InsertSectionConfigResponse{}
	_body, _err := client.DoROARequest(tea.String("InsertSectionConfig"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/edu/schedules/configs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDeviceListByCorpId(request *QueryDeviceListByCorpIdRequest) (_result *QueryDeviceListByCorpIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDeviceListByCorpIdHeaders{}
	_result = &QueryDeviceListByCorpIdResponse{}
	_body, _err := client.QueryDeviceListByCorpIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDeviceListByCorpIdWithOptions(request *QueryDeviceListByCorpIdRequest, headers *QueryDeviceListByCorpIdHeaders, runtime *util.RuntimeOptions) (_result *QueryDeviceListByCorpIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		query["operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
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
	_result = &QueryDeviceListByCorpIdResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryDeviceListByCorpId"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/edu/remoteClasses/devices"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUniversityCourseGroup(request *UpdateUniversityCourseGroupRequest) (_result *UpdateUniversityCourseGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateUniversityCourseGroupHeaders{}
	_result = &UpdateUniversityCourseGroupResponse{}
	_body, _err := client.UpdateUniversityCourseGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUniversityCourseGroupWithOptions(request *UpdateUniversityCourseGroupRequest, headers *UpdateUniversityCourseGroupHeaders, runtime *util.RuntimeOptions) (_result *UpdateUniversityCourseGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.CourseGroupCode)) {
		body["courseGroupCode"] = request.CourseGroupCode
	}

	if !tea.BoolValue(util.IsUnset(request.CourseGroupIntroduce)) {
		body["courseGroupIntroduce"] = request.CourseGroupIntroduce
	}

	if !tea.BoolValue(util.IsUnset(request.CourseGroupName)) {
		body["courseGroupName"] = request.CourseGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.CourserGroupItemModels)) {
		body["courserGroupItemModels"] = request.CourserGroupItemModels
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
	_result = &UpdateUniversityCourseGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateUniversityCourseGroup"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/edu/universities/courseGroups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchOrgCreateHW(request *BatchOrgCreateHWRequest) (_result *BatchOrgCreateHWResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchOrgCreateHWHeaders{}
	_result = &BatchOrgCreateHWResponse{}
	_body, _err := client.BatchOrgCreateHWWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchOrgCreateHWWithOptions(request *BatchOrgCreateHWRequest, headers *BatchOrgCreateHWHeaders, runtime *util.RuntimeOptions) (_result *BatchOrgCreateHWResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HwMedia)) {
		body["hwMedia"] = request.HwMedia
	}

	if !tea.BoolValue(util.IsUnset(request.HwContent)) {
		body["hwContent"] = request.HwContent
	}

	if !tea.BoolValue(util.IsUnset(request.HwTitle)) {
		body["hwTitle"] = request.HwTitle
	}

	if !tea.BoolValue(util.IsUnset(request.CourseName)) {
		body["courseName"] = request.CourseName
	}

	if !tea.BoolValue(util.IsUnset(request.HwPhoto)) {
		body["hwPhoto"] = request.HwPhoto
	}

	if !tea.BoolValue(util.IsUnset(request.HwVideo)) {
		body["hwVideo"] = request.HwVideo
	}

	if !tea.BoolValue(util.IsUnset(request.TeacherName)) {
		body["teacherName"] = request.TeacherName
	}

	if !tea.BoolValue(util.IsUnset(request.TeacherUserId)) {
		body["teacherUserId"] = request.TeacherUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Identifier)) {
		body["identifier"] = request.Identifier
	}

	if !tea.BoolValue(util.IsUnset(request.Attributes)) {
		body["attributes"] = request.Attributes
	}

	if !tea.BoolValue(util.IsUnset(request.TargetRole)) {
		body["targetRole"] = request.TargetRole
	}

	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		body["bizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduledRelease)) {
		body["scheduledRelease"] = request.ScheduledRelease
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduledTime)) {
		body["scheduledTime"] = request.ScheduledTime
	}

	if !tea.BoolValue(util.IsUnset(request.HwDeadlineOpen)) {
		body["hwDeadlineOpen"] = request.HwDeadlineOpen
	}

	if !tea.BoolValue(util.IsUnset(request.HwDeadline)) {
		body["hwDeadline"] = request.HwDeadline
	}

	if !tea.BoolValue(util.IsUnset(request.HwType)) {
		body["hwType"] = request.HwType
	}

	if !tea.BoolValue(util.IsUnset(request.OpenSelectItemList)) {
		body["openSelectItemList"] = request.OpenSelectItemList
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
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
	_result = &BatchOrgCreateHWResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchOrgCreateHW"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/edu/homeworks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCustomDept(request *CreateCustomDeptRequest) (_result *CreateCustomDeptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateCustomDeptHeaders{}
	_result = &CreateCustomDeptResponse{}
	_body, _err := client.CreateCustomDeptWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCustomDeptWithOptions(request *CreateCustomDeptRequest, headers *CreateCustomDeptHeaders, runtime *util.RuntimeOptions) (_result *CreateCustomDeptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CustomDept))) {
		body["customDept"] = request.CustomDept
	}

	if !tea.BoolValue(util.IsUnset(request.SuperId)) {
		body["superId"] = request.SuperId
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.DingOauthAppId)) {
		body["dingOauthAppId"] = request.DingOauthAppId
	}

	if !tea.BoolValue(util.IsUnset(request.DingCorpId)) {
		body["dingCorpId"] = request.DingCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
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
	_result = &CreateCustomDeptResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateCustomDept"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/edu/customDepts"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartCoursePrepare(request *StartCoursePrepareRequest) (_result *StartCoursePrepareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StartCoursePrepareHeaders{}
	_result = &StartCoursePrepareResponse{}
	_body, _err := client.StartCoursePrepareWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartCoursePrepareWithOptions(request *StartCoursePrepareRequest, headers *StartCoursePrepareHeaders, runtime *util.RuntimeOptions) (_result *StartCoursePrepareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.CourseGroupCode)) {
		body["courseGroupCode"] = request.CourseGroupCode
	}

	if !tea.BoolValue(util.IsUnset(request.SectionIndex)) {
		body["sectionIndex"] = request.SectionIndex
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		body["deviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.LiveCoverImage)) {
		body["liveCoverImage"] = request.LiveCoverImage
	}

	if !tea.BoolValue(util.IsUnset(request.CourseDate)) {
		body["courseDate"] = request.CourseDate
	}

	if !tea.BoolValue(util.IsUnset(request.IsvCode)) {
		body["isvCode"] = request.IsvCode
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
	_result = &StartCoursePrepareResponse{}
	_body, _err := client.DoROARequest(tea.String("StartCoursePrepare"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/edu/universities/courses/prepare"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryClassScheduleByTimeSchool(request *QueryClassScheduleByTimeSchoolRequest) (_result *QueryClassScheduleByTimeSchoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryClassScheduleByTimeSchoolHeaders{}
	_result = &QueryClassScheduleByTimeSchoolResponse{}
	_body, _err := client.QueryClassScheduleByTimeSchoolWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryClassScheduleByTimeSchoolWithOptions(request *QueryClassScheduleByTimeSchoolRequest, headers *QueryClassScheduleByTimeSchoolHeaders, runtime *util.RuntimeOptions) (_result *QueryClassScheduleByTimeSchoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
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
	_result = &QueryClassScheduleByTimeSchoolResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryClassScheduleByTimeSchool"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/edu/schools/classes/courses "), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTeachSubjects(request *QueryTeachSubjectsRequest) (_result *QueryTeachSubjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryTeachSubjectsHeaders{}
	_result = &QueryTeachSubjectsResponse{}
	_body, _err := client.QueryTeachSubjectsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTeachSubjectsWithOptions(request *QueryTeachSubjectsRequest, headers *QueryTeachSubjectsHeaders, runtime *util.RuntimeOptions) (_result *QueryTeachSubjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassIds)) {
		query["classIds"] = request.ClassIds
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
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
	_result = &QueryTeachSubjectsResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryTeachSubjects"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/edu/teachers/subjects"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetShareRoles() (_result *GetShareRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetShareRolesHeaders{}
	_result = &GetShareRolesResponse{}
	_body, _err := client.GetShareRolesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetShareRolesWithOptions(headers *GetShareRolesHeaders, runtime *util.RuntimeOptions) (_result *GetShareRolesResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &GetShareRolesResponse{}
	_body, _err := client.DoROARequest(tea.String("GetShareRoles"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/edu/shareRoles"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubscribeUniversityCourseGroup(request *SubscribeUniversityCourseGroupRequest) (_result *SubscribeUniversityCourseGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SubscribeUniversityCourseGroupHeaders{}
	_result = &SubscribeUniversityCourseGroupResponse{}
	_body, _err := client.SubscribeUniversityCourseGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubscribeUniversityCourseGroupWithOptions(request *SubscribeUniversityCourseGroupRequest, headers *SubscribeUniversityCourseGroupHeaders, runtime *util.RuntimeOptions) (_result *SubscribeUniversityCourseGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CourseGroupCode)) {
		body["courseGroupCode"] = request.CourseGroupCode
	}

	if !tea.BoolValue(util.IsUnset(request.StudentUserIds)) {
		body["studentUserIds"] = request.StudentUserIds
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
	_result = &SubscribeUniversityCourseGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("SubscribeUniversityCourseGroup"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/edu/universities/courseGroups/subscribe"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteOrgRelation(request *DeleteOrgRelationRequest) (_result *DeleteOrgRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteOrgRelationHeaders{}
	_result = &DeleteOrgRelationResponse{}
	_body, _err := client.DeleteOrgRelationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteOrgRelationWithOptions(request *DeleteOrgRelationRequest, headers *DeleteOrgRelationHeaders, runtime *util.RuntimeOptions) (_result *DeleteOrgRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		query["targetCorpId"] = request.TargetCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["authCode"] = request.AuthCode
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
	_result = &DeleteOrgRelationResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteOrgRelation"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/edu/remoteClasses/orgRelations"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartCourse(request *StartCourseRequest) (_result *StartCourseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartCourseResponse{}
	_body, _err := client.StartCourseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartCourseWithOptions(request *StartCourseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartCourseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.CourseCode)) {
		body["courseCode"] = request.CourseCode
	}

	if !tea.BoolValue(util.IsUnset(request.LivePlayInfoList)) {
		body["livePlayInfoList"] = request.LivePlayInfoList
	}

	if !tea.BoolValue(util.IsUnset(request.IsvCode)) {
		body["isvCode"] = request.IsvCode
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &StartCourseResponse{}
	_body, _err := client.DoROARequest(tea.String("StartCourse"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/edu/universities/courses/start"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPhysicalClassroom(request *QueryPhysicalClassroomRequest) (_result *QueryPhysicalClassroomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryPhysicalClassroomHeaders{}
	_result = &QueryPhysicalClassroomResponse{}
	_body, _err := client.QueryPhysicalClassroomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPhysicalClassroomWithOptions(request *QueryPhysicalClassroomRequest, headers *QueryPhysicalClassroomHeaders, runtime *util.RuntimeOptions) (_result *QueryPhysicalClassroomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ClassroomId)) {
		query["classroomId"] = request.ClassroomId
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
	_result = &QueryPhysicalClassroomResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryPhysicalClassroom"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/edu/physicalClassrooms"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRemoteClassCourse(request *UpdateRemoteClassCourseRequest) (_result *UpdateRemoteClassCourseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateRemoteClassCourseHeaders{}
	_result = &UpdateRemoteClassCourseResponse{}
	_body, _err := client.UpdateRemoteClassCourseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRemoteClassCourseWithOptions(request *UpdateRemoteClassCourseRequest, headers *UpdateRemoteClassCourseHeaders, runtime *util.RuntimeOptions) (_result *UpdateRemoteClassCourseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CourseName)) {
		body["courseName"] = request.CourseName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TeachingParticipant))) {
		body["teachingParticipant"] = request.TeachingParticipant
	}

	if !tea.BoolValue(util.IsUnset(request.AttendParticipants)) {
		body["attendParticipants"] = request.AttendParticipants
	}

	if !tea.BoolValue(util.IsUnset(request.CourseCode)) {
		body["courseCode"] = request.CourseCode
	}

	if !tea.BoolValue(util.IsUnset(request.DingClientId)) {
		body["dingClientId"] = request.DingClientId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.DingOauthAppId)) {
		body["dingOauthAppId"] = request.DingOauthAppId
	}

	if !tea.BoolValue(util.IsUnset(request.DingCorpId)) {
		body["dingCorpId"] = request.DingCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		body["authCode"] = request.AuthCode
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
	_result = &UpdateRemoteClassCourseResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateRemoteClassCourse"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/edu/remoteClasses/courses"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnsubscribeUniversityCourseGroup(request *UnsubscribeUniversityCourseGroupRequest) (_result *UnsubscribeUniversityCourseGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UnsubscribeUniversityCourseGroupHeaders{}
	_result = &UnsubscribeUniversityCourseGroupResponse{}
	_body, _err := client.UnsubscribeUniversityCourseGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnsubscribeUniversityCourseGroupWithOptions(request *UnsubscribeUniversityCourseGroupRequest, headers *UnsubscribeUniversityCourseGroupHeaders, runtime *util.RuntimeOptions) (_result *UnsubscribeUniversityCourseGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CourseGroupCode)) {
		body["courseGroupCode"] = request.CourseGroupCode
	}

	if !tea.BoolValue(util.IsUnset(request.StudentUserIds)) {
		body["studentUserIds"] = request.StudentUserIds
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
	_result = &UnsubscribeUniversityCourseGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("UnsubscribeUniversityCourseGroup"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/edu/universities/courseGroups/unsubscribe"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrgSecretKey(request *QueryOrgSecretKeyRequest) (_result *QueryOrgSecretKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOrgSecretKeyHeaders{}
	_result = &QueryOrgSecretKeyResponse{}
	_body, _err := client.QueryOrgSecretKeyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrgSecretKeyWithOptions(request *QueryOrgSecretKeyRequest, headers *QueryOrgSecretKeyHeaders, runtime *util.RuntimeOptions) (_result *QueryOrgSecretKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsvCode)) {
		query["isvCode"] = request.IsvCode
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
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
	_result = &QueryOrgSecretKeyResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryOrgSecretKey"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/edu/universities/secretKeys"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchTeachers(request *SearchTeachersRequest) (_result *SearchTeachersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchTeachersHeaders{}
	_result = &SearchTeachersResponse{}
	_body, _err := client.SearchTeachersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchTeachersWithOptions(request *SearchTeachersRequest, headers *SearchTeachersHeaders, runtime *util.RuntimeOptions) (_result *SearchTeachersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NameKeyword)) {
		query["nameKeyword"] = request.NameKeyword
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
	_result = &SearchTeachersResponse{}
	_body, _err := client.DoROARequest(tea.String("SearchTeachers"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/edu/teachers/search"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePhysicalClassroom(request *DeletePhysicalClassroomRequest) (_result *DeletePhysicalClassroomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeletePhysicalClassroomHeaders{}
	_result = &DeletePhysicalClassroomResponse{}
	_body, _err := client.DeletePhysicalClassroomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePhysicalClassroomWithOptions(request *DeletePhysicalClassroomRequest, headers *DeletePhysicalClassroomHeaders, runtime *util.RuntimeOptions) (_result *DeletePhysicalClassroomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ClassroomId)) {
		query["classroomId"] = request.ClassroomId
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
	_result = &DeletePhysicalClassroomResponse{}
	_body, _err := client.DoROARequest(tea.String("DeletePhysicalClassroom"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/edu/physicalClassrooms"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrgRelationList(request *QueryOrgRelationListRequest) (_result *QueryOrgRelationListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOrgRelationListHeaders{}
	_result = &QueryOrgRelationListResponse{}
	_body, _err := client.QueryOrgRelationListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrgRelationListWithOptions(request *QueryOrgRelationListRequest, headers *QueryOrgRelationListHeaders, runtime *util.RuntimeOptions) (_result *QueryOrgRelationListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		query["operator"] = request.Operator
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
	_result = &QueryOrgRelationListResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryOrgRelationList"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/edu/remoteClasses/orgRelations"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrgType() (_result *QueryOrgTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOrgTypeHeaders{}
	_result = &QueryOrgTypeResponse{}
	_body, _err := client.QueryOrgTypeWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrgTypeWithOptions(headers *QueryOrgTypeHeaders, runtime *util.RuntimeOptions) (_result *QueryOrgTypeResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &QueryOrgTypeResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryOrgType"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/edu/orgTypes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EndCourse(request *EndCourseRequest) (_result *EndCourseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EndCourseHeaders{}
	_result = &EndCourseResponse{}
	_body, _err := client.EndCourseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EndCourseWithOptions(request *EndCourseRequest, headers *EndCourseHeaders, runtime *util.RuntimeOptions) (_result *EndCourseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.CourseCode)) {
		body["courseCode"] = request.CourseCode
	}

	if !tea.BoolValue(util.IsUnset(request.LivePlayInfoList)) {
		body["livePlayInfoList"] = request.LivePlayInfoList
	}

	if !tea.BoolValue(util.IsUnset(request.IsvCode)) {
		body["isvCode"] = request.IsvCode
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
	_result = &EndCourseResponse{}
	_body, _err := client.DoROARequest(tea.String("EndCourse"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/edu/universities/courses/end"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetShareRoleMembers(shareRoleCode *string) (_result *GetShareRoleMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetShareRoleMembersHeaders{}
	_result = &GetShareRoleMembersResponse{}
	_body, _err := client.GetShareRoleMembersWithOptions(shareRoleCode, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetShareRoleMembersWithOptions(shareRoleCode *string, headers *GetShareRoleMembersHeaders, runtime *util.RuntimeOptions) (_result *GetShareRoleMembersResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &GetShareRoleMembersResponse{}
	_body, _err := client.DoROARequest(tea.String("GetShareRoleMembers"), tea.String("edu_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/edu/shareRoles/"+tea.StringValue(shareRoleCode)+"/members"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
