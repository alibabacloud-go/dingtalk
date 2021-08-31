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
