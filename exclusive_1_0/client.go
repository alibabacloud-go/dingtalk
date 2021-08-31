// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package exclusive_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetConferenceDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetConferenceDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetConferenceDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetConferenceDetailHeaders) SetCommonHeaders(v map[string]*string) *GetConferenceDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetConferenceDetailHeaders) SetXAcsDingtalkAccessToken(v string) *GetConferenceDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetConferenceDetailResponseBody struct {
	// 会议ID
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// 会议标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 开始时间
	ConfStartTime *float32 `json:"confStartTime,omitempty" xml:"confStartTime,omitempty"`
	// 持续时间
	Duration *float32 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 会议人数
	TotalNum *int64 `json:"totalNum,omitempty" xml:"totalNum,omitempty"`
	// 出席会议人数
	AttendeeNum *int64 `json:"attendeeNum,omitempty" xml:"attendeeNum,omitempty"`
	// 出席率
	AttendeePercentage *string `json:"attendeePercentage,omitempty" xml:"attendeePercentage,omitempty"`
	// 发起人uid
	CallerId *string `json:"callerId,omitempty" xml:"callerId,omitempty"`
	// 发起人昵称
	CallerName *string `json:"callerName,omitempty" xml:"callerName,omitempty"`
	// 参会人员列表
	MemberList []*GetConferenceDetailResponseBodyMemberList `json:"memberList,omitempty" xml:"memberList,omitempty" type:"Repeated"`
}

func (s GetConferenceDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConferenceDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetConferenceDetailResponseBody) SetConferenceId(v string) *GetConferenceDetailResponseBody {
	s.ConferenceId = &v
	return s
}

func (s *GetConferenceDetailResponseBody) SetTitle(v string) *GetConferenceDetailResponseBody {
	s.Title = &v
	return s
}

func (s *GetConferenceDetailResponseBody) SetConfStartTime(v float32) *GetConferenceDetailResponseBody {
	s.ConfStartTime = &v
	return s
}

func (s *GetConferenceDetailResponseBody) SetDuration(v float32) *GetConferenceDetailResponseBody {
	s.Duration = &v
	return s
}

func (s *GetConferenceDetailResponseBody) SetTotalNum(v int64) *GetConferenceDetailResponseBody {
	s.TotalNum = &v
	return s
}

func (s *GetConferenceDetailResponseBody) SetAttendeeNum(v int64) *GetConferenceDetailResponseBody {
	s.AttendeeNum = &v
	return s
}

func (s *GetConferenceDetailResponseBody) SetAttendeePercentage(v string) *GetConferenceDetailResponseBody {
	s.AttendeePercentage = &v
	return s
}

func (s *GetConferenceDetailResponseBody) SetCallerId(v string) *GetConferenceDetailResponseBody {
	s.CallerId = &v
	return s
}

func (s *GetConferenceDetailResponseBody) SetCallerName(v string) *GetConferenceDetailResponseBody {
	s.CallerName = &v
	return s
}

func (s *GetConferenceDetailResponseBody) SetMemberList(v []*GetConferenceDetailResponseBodyMemberList) *GetConferenceDetailResponseBody {
	s.MemberList = v
	return s
}

type GetConferenceDetailResponseBodyMemberList struct {
	// 用户uid
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// 用户昵称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 参会时长
	AttendDuration *float32 `json:"attendDuration,omitempty" xml:"attendDuration,omitempty"`
	// 员工id
	StaffId *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
}

func (s GetConferenceDetailResponseBodyMemberList) String() string {
	return tea.Prettify(s)
}

func (s GetConferenceDetailResponseBodyMemberList) GoString() string {
	return s.String()
}

func (s *GetConferenceDetailResponseBodyMemberList) SetUnionId(v string) *GetConferenceDetailResponseBodyMemberList {
	s.UnionId = &v
	return s
}

func (s *GetConferenceDetailResponseBodyMemberList) SetName(v string) *GetConferenceDetailResponseBodyMemberList {
	s.Name = &v
	return s
}

func (s *GetConferenceDetailResponseBodyMemberList) SetAttendDuration(v float32) *GetConferenceDetailResponseBodyMemberList {
	s.AttendDuration = &v
	return s
}

func (s *GetConferenceDetailResponseBodyMemberList) SetStaffId(v string) *GetConferenceDetailResponseBodyMemberList {
	s.StaffId = &v
	return s
}

type GetConferenceDetailResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetConferenceDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConferenceDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConferenceDetailResponse) GoString() string {
	return s.String()
}

func (s *GetConferenceDetailResponse) SetHeaders(v map[string]*string) *GetConferenceDetailResponse {
	s.Headers = v
	return s
}

func (s *GetConferenceDetailResponse) SetBody(v *GetConferenceDetailResponseBody) *GetConferenceDetailResponse {
	s.Body = v
	return s
}

type GetOaOperatorLogListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOaOperatorLogListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOaOperatorLogListHeaders) GoString() string {
	return s.String()
}

func (s *GetOaOperatorLogListHeaders) SetCommonHeaders(v map[string]*string) *GetOaOperatorLogListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOaOperatorLogListHeaders) SetXAcsDingtalkAccessToken(v string) *GetOaOperatorLogListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOaOperatorLogListRequest struct {
	// 操作员userId
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 起始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 结束时间
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 分页起始页
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页大小
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 操作分类（一级目录）
	CategoryList []*string `json:"categoryList,omitempty" xml:"categoryList,omitempty" type:"Repeated"`
}

func (s GetOaOperatorLogListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOaOperatorLogListRequest) GoString() string {
	return s.String()
}

func (s *GetOaOperatorLogListRequest) SetOpUserId(v string) *GetOaOperatorLogListRequest {
	s.OpUserId = &v
	return s
}

func (s *GetOaOperatorLogListRequest) SetStartTime(v int64) *GetOaOperatorLogListRequest {
	s.StartTime = &v
	return s
}

func (s *GetOaOperatorLogListRequest) SetEndTime(v int64) *GetOaOperatorLogListRequest {
	s.EndTime = &v
	return s
}

func (s *GetOaOperatorLogListRequest) SetPageNumber(v int64) *GetOaOperatorLogListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetOaOperatorLogListRequest) SetPageSize(v int32) *GetOaOperatorLogListRequest {
	s.PageSize = &v
	return s
}

func (s *GetOaOperatorLogListRequest) SetCategoryList(v []*string) *GetOaOperatorLogListRequest {
	s.CategoryList = v
	return s
}

type GetOaOperatorLogListResponseBody struct {
	Data []*GetOaOperatorLogListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 当前获取记录数
	ItemCount *int64 `json:"itemCount,omitempty" xml:"itemCount,omitempty"`
}

func (s GetOaOperatorLogListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOaOperatorLogListResponseBody) GoString() string {
	return s.String()
}

func (s *GetOaOperatorLogListResponseBody) SetData(v []*GetOaOperatorLogListResponseBodyData) *GetOaOperatorLogListResponseBody {
	s.Data = v
	return s
}

func (s *GetOaOperatorLogListResponseBody) SetItemCount(v int64) *GetOaOperatorLogListResponseBody {
	s.ItemCount = &v
	return s
}

type GetOaOperatorLogListResponseBodyData struct {
	// 操作员userId
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 操作员名字
	OpName *string `json:"opName,omitempty" xml:"opName,omitempty"`
	// 操作时间
	OpTime *int64 `json:"opTime,omitempty" xml:"opTime,omitempty"`
	// 操作分类（一级）
	Category1Name *string `json:"category1Name,omitempty" xml:"category1Name,omitempty"`
	// 操作分类（二级）
	Category2Name *string `json:"category2Name,omitempty" xml:"category2Name,omitempty"`
	// 操作详情
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s GetOaOperatorLogListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOaOperatorLogListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOaOperatorLogListResponseBodyData) SetOpUserId(v string) *GetOaOperatorLogListResponseBodyData {
	s.OpUserId = &v
	return s
}

func (s *GetOaOperatorLogListResponseBodyData) SetOpName(v string) *GetOaOperatorLogListResponseBodyData {
	s.OpName = &v
	return s
}

func (s *GetOaOperatorLogListResponseBodyData) SetOpTime(v int64) *GetOaOperatorLogListResponseBodyData {
	s.OpTime = &v
	return s
}

func (s *GetOaOperatorLogListResponseBodyData) SetCategory1Name(v string) *GetOaOperatorLogListResponseBodyData {
	s.Category1Name = &v
	return s
}

func (s *GetOaOperatorLogListResponseBodyData) SetCategory2Name(v string) *GetOaOperatorLogListResponseBodyData {
	s.Category2Name = &v
	return s
}

func (s *GetOaOperatorLogListResponseBodyData) SetContent(v string) *GetOaOperatorLogListResponseBodyData {
	s.Content = &v
	return s
}

type GetOaOperatorLogListResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOaOperatorLogListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOaOperatorLogListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOaOperatorLogListResponse) GoString() string {
	return s.String()
}

func (s *GetOaOperatorLogListResponse) SetHeaders(v map[string]*string) *GetOaOperatorLogListResponse {
	s.Headers = v
	return s
}

func (s *GetOaOperatorLogListResponse) SetBody(v *GetOaOperatorLogListResponseBody) *GetOaOperatorLogListResponse {
	s.Body = v
	return s
}

type DeleteCommentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteCommentHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommentHeaders) GoString() string {
	return s.String()
}

func (s *DeleteCommentHeaders) SetCommonHeaders(v map[string]*string) *DeleteCommentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteCommentHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteCommentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteCommentResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *bool              `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommentResponse) GoString() string {
	return s.String()
}

func (s *DeleteCommentResponse) SetHeaders(v map[string]*string) *DeleteCommentResponse {
	s.Headers = v
	return s
}

func (s *DeleteCommentResponse) SetBody(v bool) *DeleteCommentResponse {
	s.Body = &v
	return s
}

type GetAllLabelableDeptsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAllLabelableDeptsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAllLabelableDeptsHeaders) GoString() string {
	return s.String()
}

func (s *GetAllLabelableDeptsHeaders) SetCommonHeaders(v map[string]*string) *GetAllLabelableDeptsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAllLabelableDeptsHeaders) SetXAcsDingtalkAccessToken(v string) *GetAllLabelableDeptsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAllLabelableDeptsResponseBody struct {
	// 伙伴钉可打标部门列表
	Data []*GetAllLabelableDeptsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s GetAllLabelableDeptsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAllLabelableDeptsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllLabelableDeptsResponseBody) SetData(v []*GetAllLabelableDeptsResponseBodyData) *GetAllLabelableDeptsResponseBody {
	s.Data = v
	return s
}

type GetAllLabelableDeptsResponseBodyData struct {
	// 部门id
	DeptId *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 父部门id
	SuperDeptId *string `json:"superDeptId,omitempty" xml:"superDeptId,omitempty"`
	// 部门名称
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// 部门人数
	MemberCount *int64 `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	// 部门伙伴编码
	PartnerNum *string `json:"partnerNum,omitempty" xml:"partnerNum,omitempty"`
	// 部门一级伙伴类型
	PartnerLabelVOLevel1 *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel1 `json:"partnerLabelVOLevel1,omitempty" xml:"partnerLabelVOLevel1,omitempty" type:"Struct"`
	// 部门二级伙伴类型
	PartnerLabelVOLevel2 *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel2 `json:"partnerLabelVOLevel2,omitempty" xml:"partnerLabelVOLevel2,omitempty" type:"Struct"`
	// 部门三级伙伴类型
	PartnerLabelVOLevel3 *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel3 `json:"partnerLabelVOLevel3,omitempty" xml:"partnerLabelVOLevel3,omitempty" type:"Struct"`
	// 部门四级伙伴类型
	PartnerLabelVOLevel4 *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel4 `json:"partnerLabelVOLevel4,omitempty" xml:"partnerLabelVOLevel4,omitempty" type:"Struct"`
	// 部门五级伙伴类型
	PartnerLabelVOLevel5 *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel5 `json:"partnerLabelVOLevel5,omitempty" xml:"partnerLabelVOLevel5,omitempty" type:"Struct"`
}

func (s GetAllLabelableDeptsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAllLabelableDeptsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAllLabelableDeptsResponseBodyData) SetDeptId(v string) *GetAllLabelableDeptsResponseBodyData {
	s.DeptId = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyData) SetSuperDeptId(v string) *GetAllLabelableDeptsResponseBodyData {
	s.SuperDeptId = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyData) SetDeptName(v string) *GetAllLabelableDeptsResponseBodyData {
	s.DeptName = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyData) SetMemberCount(v int64) *GetAllLabelableDeptsResponseBodyData {
	s.MemberCount = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyData) SetPartnerNum(v string) *GetAllLabelableDeptsResponseBodyData {
	s.PartnerNum = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyData) SetPartnerLabelVOLevel1(v *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel1) *GetAllLabelableDeptsResponseBodyData {
	s.PartnerLabelVOLevel1 = v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyData) SetPartnerLabelVOLevel2(v *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel2) *GetAllLabelableDeptsResponseBodyData {
	s.PartnerLabelVOLevel2 = v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyData) SetPartnerLabelVOLevel3(v *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel3) *GetAllLabelableDeptsResponseBodyData {
	s.PartnerLabelVOLevel3 = v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyData) SetPartnerLabelVOLevel4(v *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel4) *GetAllLabelableDeptsResponseBodyData {
	s.PartnerLabelVOLevel4 = v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyData) SetPartnerLabelVOLevel5(v *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel5) *GetAllLabelableDeptsResponseBodyData {
	s.PartnerLabelVOLevel5 = v
	return s
}

type GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel1 struct {
	// 伙伴类型id
	LabelId *int64 `json:"labelId,omitempty" xml:"labelId,omitempty"`
	// 伙伴类型
	LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty"`
	// 伙伴类型层级
	LevelNum *int64 `json:"levelNum,omitempty" xml:"levelNum,omitempty"`
}

func (s GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel1) String() string {
	return tea.Prettify(s)
}

func (s GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel1) GoString() string {
	return s.String()
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel1) SetLabelId(v int64) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel1 {
	s.LabelId = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel1) SetLabelName(v string) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel1 {
	s.LabelName = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel1) SetLevelNum(v int64) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel1 {
	s.LevelNum = &v
	return s
}

type GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel2 struct {
	// 伙伴类型id
	LabelId *int64 `json:"labelId,omitempty" xml:"labelId,omitempty"`
	// 伙伴类型
	LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty"`
	// 伙伴类型层级
	LevelNum *int64 `json:"levelNum,omitempty" xml:"levelNum,omitempty"`
}

func (s GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel2) String() string {
	return tea.Prettify(s)
}

func (s GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel2) GoString() string {
	return s.String()
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel2) SetLabelId(v int64) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel2 {
	s.LabelId = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel2) SetLabelName(v string) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel2 {
	s.LabelName = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel2) SetLevelNum(v int64) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel2 {
	s.LevelNum = &v
	return s
}

type GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel3 struct {
	// 伙伴类型id
	LabelId *int64 `json:"labelId,omitempty" xml:"labelId,omitempty"`
	// 伙伴类型
	LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty"`
	// 伙伴类型层级
	LevelNum *int64 `json:"levelNum,omitempty" xml:"levelNum,omitempty"`
}

func (s GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel3) String() string {
	return tea.Prettify(s)
}

func (s GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel3) GoString() string {
	return s.String()
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel3) SetLabelId(v int64) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel3 {
	s.LabelId = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel3) SetLabelName(v string) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel3 {
	s.LabelName = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel3) SetLevelNum(v int64) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel3 {
	s.LevelNum = &v
	return s
}

type GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel4 struct {
	// 伙伴类型id
	LabelId *int64 `json:"labelId,omitempty" xml:"labelId,omitempty"`
	// 伙伴类型
	LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty"`
	// 伙伴类型层级
	LevelNum *int64 `json:"levelNum,omitempty" xml:"levelNum,omitempty"`
}

func (s GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel4) String() string {
	return tea.Prettify(s)
}

func (s GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel4) GoString() string {
	return s.String()
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel4) SetLabelId(v int64) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel4 {
	s.LabelId = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel4) SetLabelName(v string) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel4 {
	s.LabelName = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel4) SetLevelNum(v int64) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel4 {
	s.LevelNum = &v
	return s
}

type GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel5 struct {
	// 伙伴类型id
	LabelId *int64 `json:"labelId,omitempty" xml:"labelId,omitempty"`
	// 伙伴类型
	LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty"`
	// 伙伴类型层级
	LevelNum *int64 `json:"levelNum,omitempty" xml:"levelNum,omitempty"`
}

func (s GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel5) String() string {
	return tea.Prettify(s)
}

func (s GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel5) GoString() string {
	return s.String()
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel5) SetLabelId(v int64) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel5 {
	s.LabelId = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel5) SetLabelName(v string) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel5 {
	s.LabelName = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel5) SetLevelNum(v int64) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel5 {
	s.LevelNum = &v
	return s
}

type GetAllLabelableDeptsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAllLabelableDeptsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAllLabelableDeptsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllLabelableDeptsResponse) GoString() string {
	return s.String()
}

func (s *GetAllLabelableDeptsResponse) SetHeaders(v map[string]*string) *GetAllLabelableDeptsResponse {
	s.Headers = v
	return s
}

func (s *GetAllLabelableDeptsResponse) SetBody(v *GetAllLabelableDeptsResponseBody) *GetAllLabelableDeptsResponse {
	s.Body = v
	return s
}

type GetTrustDeviceListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTrustDeviceListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTrustDeviceListHeaders) GoString() string {
	return s.String()
}

func (s *GetTrustDeviceListHeaders) SetCommonHeaders(v map[string]*string) *GetTrustDeviceListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTrustDeviceListHeaders) SetXAcsDingtalkAccessToken(v string) *GetTrustDeviceListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTrustDeviceListRequest struct {
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s GetTrustDeviceListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTrustDeviceListRequest) GoString() string {
	return s.String()
}

func (s *GetTrustDeviceListRequest) SetUserIds(v []*string) *GetTrustDeviceListRequest {
	s.UserIds = v
	return s
}

type GetTrustDeviceListResponseBody struct {
	Data []*GetTrustDeviceListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s GetTrustDeviceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrustDeviceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrustDeviceListResponseBody) SetData(v []*GetTrustDeviceListResponseBodyData) *GetTrustDeviceListResponseBody {
	s.Data = v
	return s
}

type GetTrustDeviceListResponseBodyData struct {
	// 员工Id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 平台类型
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// mac地址
	MacAddress *string `json:"macAddress,omitempty" xml:"macAddress,omitempty"`
	// 设备状态
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// 创建时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
}

func (s GetTrustDeviceListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTrustDeviceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTrustDeviceListResponseBodyData) SetUserId(v string) *GetTrustDeviceListResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetTrustDeviceListResponseBodyData) SetPlatform(v string) *GetTrustDeviceListResponseBodyData {
	s.Platform = &v
	return s
}

func (s *GetTrustDeviceListResponseBodyData) SetMacAddress(v string) *GetTrustDeviceListResponseBodyData {
	s.MacAddress = &v
	return s
}

func (s *GetTrustDeviceListResponseBodyData) SetStatus(v int32) *GetTrustDeviceListResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTrustDeviceListResponseBodyData) SetCreateTime(v int64) *GetTrustDeviceListResponseBodyData {
	s.CreateTime = &v
	return s
}

type GetTrustDeviceListResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTrustDeviceListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTrustDeviceListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrustDeviceListResponse) GoString() string {
	return s.String()
}

func (s *GetTrustDeviceListResponse) SetHeaders(v map[string]*string) *GetTrustDeviceListResponse {
	s.Headers = v
	return s
}

func (s *GetTrustDeviceListResponse) SetBody(v *GetTrustDeviceListResponseBody) *GetTrustDeviceListResponse {
	s.Body = v
	return s
}

type SearchOrgInnerGroupInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchOrgInnerGroupInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchOrgInnerGroupInfoHeaders) GoString() string {
	return s.String()
}

func (s *SearchOrgInnerGroupInfoHeaders) SetCommonHeaders(v map[string]*string) *SearchOrgInnerGroupInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchOrgInnerGroupInfoHeaders) SetXAcsDingtalkAccessToken(v string) *SearchOrgInnerGroupInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchOrgInnerGroupInfoRequest struct {
	// groupMembersCntEnd
	GroupMembersCountEnd *int32 `json:"groupMembersCountEnd,omitempty" xml:"groupMembersCountEnd,omitempty"`
	// syncToDingpan
	SyncToDingpan *int32 `json:"syncToDingpan,omitempty" xml:"syncToDingpan,omitempty"`
	// groupOwner
	GroupOwner *string `json:"groupOwner,omitempty" xml:"groupOwner,omitempty"`
	// createTimeEnd
	CreateTimeEnd *int64 `json:"createTimeEnd,omitempty" xml:"createTimeEnd,omitempty"`
	// pageSize
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// createTimeStart
	CreateTimeStart *int64 `json:"createTimeStart,omitempty" xml:"createTimeStart,omitempty"`
	// uuid
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// groupMembersCntStart
	GroupMembersCountStart *int32 `json:"groupMembersCountStart,omitempty" xml:"groupMembersCountStart,omitempty"`
	// lastActiveTimeEnd
	LastActiveTimeEnd *int64 `json:"lastActiveTimeEnd,omitempty" xml:"lastActiveTimeEnd,omitempty"`
	// operatorUserId
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// groupName
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// pageStart
	PageStart *int32 `json:"pageStart,omitempty" xml:"pageStart,omitempty"`
	// lastActiveTimeStart
	LastActiveTimeStart *int64 `json:"lastActiveTimeStart,omitempty" xml:"lastActiveTimeStart,omitempty"`
}

func (s SearchOrgInnerGroupInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchOrgInnerGroupInfoRequest) GoString() string {
	return s.String()
}

func (s *SearchOrgInnerGroupInfoRequest) SetGroupMembersCountEnd(v int32) *SearchOrgInnerGroupInfoRequest {
	s.GroupMembersCountEnd = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetSyncToDingpan(v int32) *SearchOrgInnerGroupInfoRequest {
	s.SyncToDingpan = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetGroupOwner(v string) *SearchOrgInnerGroupInfoRequest {
	s.GroupOwner = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetCreateTimeEnd(v int64) *SearchOrgInnerGroupInfoRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetPageSize(v int32) *SearchOrgInnerGroupInfoRequest {
	s.PageSize = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetCreateTimeStart(v int64) *SearchOrgInnerGroupInfoRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetUuid(v string) *SearchOrgInnerGroupInfoRequest {
	s.Uuid = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetGroupMembersCountStart(v int32) *SearchOrgInnerGroupInfoRequest {
	s.GroupMembersCountStart = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetLastActiveTimeEnd(v int64) *SearchOrgInnerGroupInfoRequest {
	s.LastActiveTimeEnd = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetOperatorUserId(v string) *SearchOrgInnerGroupInfoRequest {
	s.OperatorUserId = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetGroupName(v string) *SearchOrgInnerGroupInfoRequest {
	s.GroupName = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetPageStart(v int32) *SearchOrgInnerGroupInfoRequest {
	s.PageStart = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetLastActiveTimeStart(v int64) *SearchOrgInnerGroupInfoRequest {
	s.LastActiveTimeStart = &v
	return s
}

type SearchOrgInnerGroupInfoResponseBody struct {
	TotalCount *int64                                      `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	ItemCount  *int32                                      `json:"itemCount,omitempty" xml:"itemCount,omitempty"`
	Items      []*SearchOrgInnerGroupInfoResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s SearchOrgInnerGroupInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchOrgInnerGroupInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SearchOrgInnerGroupInfoResponseBody) SetTotalCount(v int64) *SearchOrgInnerGroupInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBody) SetItemCount(v int32) *SearchOrgInnerGroupInfoResponseBody {
	s.ItemCount = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBody) SetItems(v []*SearchOrgInnerGroupInfoResponseBodyItems) *SearchOrgInnerGroupInfoResponseBody {
	s.Items = v
	return s
}

type SearchOrgInnerGroupInfoResponseBodyItems struct {
	OpenConversationId      *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	GroupOwner              *string `json:"groupOwner,omitempty" xml:"groupOwner,omitempty"`
	GroupName               *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	GroupAdminsCount        *int32  `json:"groupAdminsCount,omitempty" xml:"groupAdminsCount,omitempty"`
	GroupMembersCount       *int32  `json:"groupMembersCount,omitempty" xml:"groupMembersCount,omitempty"`
	GroupCreateTime         *int64  `json:"groupCreateTime,omitempty" xml:"groupCreateTime,omitempty"`
	GroupLastActiveTime     *int64  `json:"groupLastActiveTime,omitempty" xml:"groupLastActiveTime,omitempty"`
	GroupLastActiveTimeShow *string `json:"groupLastActiveTimeShow,omitempty" xml:"groupLastActiveTimeShow,omitempty"`
	SyncToDingpan           *int32  `json:"syncToDingpan,omitempty" xml:"syncToDingpan,omitempty"`
	UsedQuota               *int64  `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
	GroupOwnerUserId        *string `json:"groupOwnerUserId,omitempty" xml:"groupOwnerUserId,omitempty"`
	Status                  *int32  `json:"status,omitempty" xml:"status,omitempty"`
	TemplateId              *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	TemplateName            *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s SearchOrgInnerGroupInfoResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s SearchOrgInnerGroupInfoResponseBodyItems) GoString() string {
	return s.String()
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetOpenConversationId(v string) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.OpenConversationId = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetGroupOwner(v string) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.GroupOwner = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetGroupName(v string) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.GroupName = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetGroupAdminsCount(v int32) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.GroupAdminsCount = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetGroupMembersCount(v int32) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.GroupMembersCount = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetGroupCreateTime(v int64) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.GroupCreateTime = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetGroupLastActiveTime(v int64) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.GroupLastActiveTime = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetGroupLastActiveTimeShow(v string) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.GroupLastActiveTimeShow = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetSyncToDingpan(v int32) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.SyncToDingpan = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetUsedQuota(v int64) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.UsedQuota = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetGroupOwnerUserId(v string) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.GroupOwnerUserId = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetStatus(v int32) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.Status = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetTemplateId(v string) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.TemplateId = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetTemplateName(v string) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.TemplateName = &v
	return s
}

type SearchOrgInnerGroupInfoResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchOrgInnerGroupInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchOrgInnerGroupInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchOrgInnerGroupInfoResponse) GoString() string {
	return s.String()
}

func (s *SearchOrgInnerGroupInfoResponse) SetHeaders(v map[string]*string) *SearchOrgInnerGroupInfoResponse {
	s.Headers = v
	return s
}

func (s *SearchOrgInnerGroupInfoResponse) SetBody(v *SearchOrgInnerGroupInfoResponseBody) *SearchOrgInnerGroupInfoResponse {
	s.Body = v
	return s
}

type CreateTrustedDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateTrustedDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTrustedDeviceHeaders) GoString() string {
	return s.String()
}

func (s *CreateTrustedDeviceHeaders) SetCommonHeaders(v map[string]*string) *CreateTrustedDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTrustedDeviceHeaders) SetXAcsDingtalkAccessToken(v string) *CreateTrustedDeviceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateTrustedDeviceRequest struct {
	// 员工userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 平台类型
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// mac地址
	MacAddress *string `json:"macAddress,omitempty" xml:"macAddress,omitempty"`
	// 设备状态
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateTrustedDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTrustedDeviceRequest) GoString() string {
	return s.String()
}

func (s *CreateTrustedDeviceRequest) SetUserId(v string) *CreateTrustedDeviceRequest {
	s.UserId = &v
	return s
}

func (s *CreateTrustedDeviceRequest) SetPlatform(v string) *CreateTrustedDeviceRequest {
	s.Platform = &v
	return s
}

func (s *CreateTrustedDeviceRequest) SetMacAddress(v string) *CreateTrustedDeviceRequest {
	s.MacAddress = &v
	return s
}

func (s *CreateTrustedDeviceRequest) SetStatus(v int32) *CreateTrustedDeviceRequest {
	s.Status = &v
	return s
}

type CreateTrustedDeviceResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateTrustedDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTrustedDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrustedDeviceResponseBody) SetSuccess(v bool) *CreateTrustedDeviceResponseBody {
	s.Success = &v
	return s
}

type CreateTrustedDeviceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTrustedDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTrustedDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTrustedDeviceResponse) GoString() string {
	return s.String()
}

func (s *CreateTrustedDeviceResponse) SetHeaders(v map[string]*string) *CreateTrustedDeviceResponse {
	s.Headers = v
	return s
}

func (s *CreateTrustedDeviceResponse) SetBody(v *CreateTrustedDeviceResponseBody) *CreateTrustedDeviceResponse {
	s.Body = v
	return s
}

type GetGroupActiveInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetGroupActiveInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetGroupActiveInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetGroupActiveInfoHeaders) SetCommonHeaders(v map[string]*string) *GetGroupActiveInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetGroupActiveInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetGroupActiveInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetGroupActiveInfoRequest struct {
	// 统计日期
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
	// 钉钉群组id
	DingGroupId *string `json:"dingGroupId,omitempty" xml:"dingGroupId,omitempty"`
	// 分页起始页
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页大小
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetGroupActiveInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGroupActiveInfoRequest) GoString() string {
	return s.String()
}

func (s *GetGroupActiveInfoRequest) SetStatDate(v string) *GetGroupActiveInfoRequest {
	s.StatDate = &v
	return s
}

func (s *GetGroupActiveInfoRequest) SetDingGroupId(v string) *GetGroupActiveInfoRequest {
	s.DingGroupId = &v
	return s
}

func (s *GetGroupActiveInfoRequest) SetPageNumber(v int64) *GetGroupActiveInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *GetGroupActiveInfoRequest) SetPageSize(v int64) *GetGroupActiveInfoRequest {
	s.PageSize = &v
	return s
}

type GetGroupActiveInfoResponseBody struct {
	Data       []*GetGroupActiveInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	TotalCount *int64                                `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetGroupActiveInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGroupActiveInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupActiveInfoResponseBody) SetData(v []*GetGroupActiveInfoResponseBodyData) *GetGroupActiveInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetGroupActiveInfoResponseBody) SetTotalCount(v int64) *GetGroupActiveInfoResponseBody {
	s.TotalCount = &v
	return s
}

type GetGroupActiveInfoResponseBodyData struct {
	// 统计时间
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
	// 群组id
	DingGroupId *string `json:"dingGroupId,omitempty" xml:"dingGroupId,omitempty"`
	// 群组创建时间
	GroupCreateTime *string `json:"groupCreateTime,omitempty" xml:"groupCreateTime,omitempty"`
	// 群组创建用户id
	GroupCreateUserId *string `json:"groupCreateUserId,omitempty" xml:"groupCreateUserId,omitempty"`
	// 群组创建用户姓名
	GroupCreateUserName *string `json:"groupCreateUserName,omitempty" xml:"groupCreateUserName,omitempty"`
	// 群名称
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// 群类型：1-全员群，2-部门群，3-（其他）内部群，4-场景群
	GroupType *int64 `json:"groupType,omitempty" xml:"groupType,omitempty"`
	// 最近1天群人数
	GroupUserCnt1d *int32 `json:"groupUserCnt1d,omitempty" xml:"groupUserCnt1d,omitempty"`
	// 最近1天发消息人数
	SendMessageUserCnt1d *int64 `json:"sendMessageUserCnt1d,omitempty" xml:"sendMessageUserCnt1d,omitempty"`
	// 最近1天发消息次数
	SendMessageCnt1d *int64 `json:"sendMessageCnt1d,omitempty" xml:"sendMessageCnt1d,omitempty"`
	// 最近1天打开群人数
	OpenConvUv1d *int32 `json:"openConvUv1d,omitempty" xml:"openConvUv1d,omitempty"`
}

func (s GetGroupActiveInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetGroupActiveInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGroupActiveInfoResponseBodyData) SetStatDate(v string) *GetGroupActiveInfoResponseBodyData {
	s.StatDate = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetDingGroupId(v string) *GetGroupActiveInfoResponseBodyData {
	s.DingGroupId = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetGroupCreateTime(v string) *GetGroupActiveInfoResponseBodyData {
	s.GroupCreateTime = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetGroupCreateUserId(v string) *GetGroupActiveInfoResponseBodyData {
	s.GroupCreateUserId = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetGroupCreateUserName(v string) *GetGroupActiveInfoResponseBodyData {
	s.GroupCreateUserName = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetGroupName(v string) *GetGroupActiveInfoResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetGroupType(v int64) *GetGroupActiveInfoResponseBodyData {
	s.GroupType = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetGroupUserCnt1d(v int32) *GetGroupActiveInfoResponseBodyData {
	s.GroupUserCnt1d = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetSendMessageUserCnt1d(v int64) *GetGroupActiveInfoResponseBodyData {
	s.SendMessageUserCnt1d = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetSendMessageCnt1d(v int64) *GetGroupActiveInfoResponseBodyData {
	s.SendMessageCnt1d = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetOpenConvUv1d(v int32) *GetGroupActiveInfoResponseBodyData {
	s.OpenConvUv1d = &v
	return s
}

type GetGroupActiveInfoResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetGroupActiveInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGroupActiveInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGroupActiveInfoResponse) GoString() string {
	return s.String()
}

func (s *GetGroupActiveInfoResponse) SetHeaders(v map[string]*string) *GetGroupActiveInfoResponse {
	s.Headers = v
	return s
}

func (s *GetGroupActiveInfoResponse) SetBody(v *GetGroupActiveInfoResponseBody) *GetGroupActiveInfoResponse {
	s.Body = v
	return s
}

type GetCommentListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCommentListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCommentListHeaders) GoString() string {
	return s.String()
}

func (s *GetCommentListHeaders) SetCommonHeaders(v map[string]*string) *GetCommentListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCommentListHeaders) SetXAcsDingtalkAccessToken(v string) *GetCommentListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCommentListRequest struct {
	// 分页起始页
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页大小
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetCommentListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCommentListRequest) GoString() string {
	return s.String()
}

func (s *GetCommentListRequest) SetPageNumber(v int64) *GetCommentListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetCommentListRequest) SetPageSize(v int64) *GetCommentListRequest {
	s.PageSize = &v
	return s
}

type GetCommentListResponseBody struct {
	Data       []*GetCommentListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	TotalCount *int64                            `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetCommentListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCommentListResponseBody) GoString() string {
	return s.String()
}

func (s *GetCommentListResponseBody) SetData(v []*GetCommentListResponseBodyData) *GetCommentListResponseBody {
	s.Data = v
	return s
}

func (s *GetCommentListResponseBody) SetTotalCount(v int64) *GetCommentListResponseBody {
	s.TotalCount = &v
	return s
}

type GetCommentListResponseBodyData struct {
	// 评论者姓名
	CommentUserName *string `json:"commentUserName,omitempty" xml:"commentUserName,omitempty"`
	// 评论内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 评论时间
	CommentTime *float32 `json:"commentTime,omitempty" xml:"commentTime,omitempty"`
	// 评论ID
	CommentId *string `json:"commentId,omitempty" xml:"commentId,omitempty"`
}

func (s GetCommentListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCommentListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCommentListResponseBodyData) SetCommentUserName(v string) *GetCommentListResponseBodyData {
	s.CommentUserName = &v
	return s
}

func (s *GetCommentListResponseBodyData) SetContent(v string) *GetCommentListResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetCommentListResponseBodyData) SetCommentTime(v float32) *GetCommentListResponseBodyData {
	s.CommentTime = &v
	return s
}

func (s *GetCommentListResponseBodyData) SetCommentId(v string) *GetCommentListResponseBodyData {
	s.CommentId = &v
	return s
}

type GetCommentListResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCommentListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCommentListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCommentListResponse) GoString() string {
	return s.String()
}

func (s *GetCommentListResponse) SetHeaders(v map[string]*string) *GetCommentListResponse {
	s.Headers = v
	return s
}

func (s *GetCommentListResponse) SetBody(v *GetCommentListResponseBody) *GetCommentListResponse {
	s.Body = v
	return s
}

type GetPartnerTypeByParentIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPartnerTypeByParentIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPartnerTypeByParentIdHeaders) GoString() string {
	return s.String()
}

func (s *GetPartnerTypeByParentIdHeaders) SetCommonHeaders(v map[string]*string) *GetPartnerTypeByParentIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPartnerTypeByParentIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetPartnerTypeByParentIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPartnerTypeByParentIdResponseBody struct {
	// 子标签列表
	Data []*GetPartnerTypeByParentIdResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s GetPartnerTypeByParentIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPartnerTypeByParentIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetPartnerTypeByParentIdResponseBody) SetData(v []*GetPartnerTypeByParentIdResponseBodyData) *GetPartnerTypeByParentIdResponseBody {
	s.Data = v
	return s
}

type GetPartnerTypeByParentIdResponseBodyData struct {
	// 自标签id
	TypeId *float32 `json:"typeId,omitempty" xml:"typeId,omitempty"`
	// 子标签名
	TypeName *string `json:"typeName,omitempty" xml:"typeName,omitempty"`
}

func (s GetPartnerTypeByParentIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPartnerTypeByParentIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPartnerTypeByParentIdResponseBodyData) SetTypeId(v float32) *GetPartnerTypeByParentIdResponseBodyData {
	s.TypeId = &v
	return s
}

func (s *GetPartnerTypeByParentIdResponseBodyData) SetTypeName(v string) *GetPartnerTypeByParentIdResponseBodyData {
	s.TypeName = &v
	return s
}

type GetPartnerTypeByParentIdResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPartnerTypeByParentIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPartnerTypeByParentIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPartnerTypeByParentIdResponse) GoString() string {
	return s.String()
}

func (s *GetPartnerTypeByParentIdResponse) SetHeaders(v map[string]*string) *GetPartnerTypeByParentIdResponse {
	s.Headers = v
	return s
}

func (s *GetPartnerTypeByParentIdResponse) SetBody(v *GetPartnerTypeByParentIdResponseBody) *GetPartnerTypeByParentIdResponse {
	s.Body = v
	return s
}

type SetDeptPartnerTypeAndNumHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetDeptPartnerTypeAndNumHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetDeptPartnerTypeAndNumHeaders) GoString() string {
	return s.String()
}

func (s *SetDeptPartnerTypeAndNumHeaders) SetCommonHeaders(v map[string]*string) *SetDeptPartnerTypeAndNumHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetDeptPartnerTypeAndNumHeaders) SetXAcsDingtalkAccessToken(v string) *SetDeptPartnerTypeAndNumHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetDeptPartnerTypeAndNumRequest struct {
	// 部门id
	DeptId *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 伙伴编码
	PartnerNum *string `json:"partnerNum,omitempty" xml:"partnerNum,omitempty"`
	// 伙伴类型id列表
	LabelIds []*string `json:"labelIds,omitempty" xml:"labelIds,omitempty" type:"Repeated"`
}

func (s SetDeptPartnerTypeAndNumRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDeptPartnerTypeAndNumRequest) GoString() string {
	return s.String()
}

func (s *SetDeptPartnerTypeAndNumRequest) SetDeptId(v string) *SetDeptPartnerTypeAndNumRequest {
	s.DeptId = &v
	return s
}

func (s *SetDeptPartnerTypeAndNumRequest) SetPartnerNum(v string) *SetDeptPartnerTypeAndNumRequest {
	s.PartnerNum = &v
	return s
}

func (s *SetDeptPartnerTypeAndNumRequest) SetLabelIds(v []*string) *SetDeptPartnerTypeAndNumRequest {
	s.LabelIds = v
	return s
}

type SetDeptPartnerTypeAndNumResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s SetDeptPartnerTypeAndNumResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDeptPartnerTypeAndNumResponse) GoString() string {
	return s.String()
}

func (s *SetDeptPartnerTypeAndNumResponse) SetHeaders(v map[string]*string) *SetDeptPartnerTypeAndNumResponse {
	s.Headers = v
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

func (client *Client) GetConferenceDetail(conferenceId *string) (_result *GetConferenceDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetConferenceDetailHeaders{}
	_result = &GetConferenceDetailResponse{}
	_body, _err := client.GetConferenceDetailWithOptions(conferenceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConferenceDetailWithOptions(conferenceId *string, headers *GetConferenceDetailHeaders, runtime *util.RuntimeOptions) (_result *GetConferenceDetailResponse, _err error) {
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
	_result = &GetConferenceDetailResponse{}
	_body, _err := client.DoROARequest(tea.String("GetConferenceDetail"), tea.String("exclusive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/exclusive/data/conferences/"+tea.StringValue(conferenceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOaOperatorLogList(request *GetOaOperatorLogListRequest) (_result *GetOaOperatorLogListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOaOperatorLogListHeaders{}
	_result = &GetOaOperatorLogListResponse{}
	_body, _err := client.GetOaOperatorLogListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOaOperatorLogListWithOptions(request *GetOaOperatorLogListRequest, headers *GetOaOperatorLogListHeaders, runtime *util.RuntimeOptions) (_result *GetOaOperatorLogListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryList)) {
		body["categoryList"] = request.CategoryList
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
	_result = &GetOaOperatorLogListResponse{}
	_body, _err := client.DoROARequest(tea.String("GetOaOperatorLogList"), tea.String("exclusive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/exclusive/oaOperatorLogs/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteComment(publisherId *string, commentId *string) (_result *DeleteCommentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteCommentHeaders{}
	_result = &DeleteCommentResponse{}
	_body, _err := client.DeleteCommentWithOptions(publisherId, commentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCommentWithOptions(publisherId *string, commentId *string, headers *DeleteCommentHeaders, runtime *util.RuntimeOptions) (_result *DeleteCommentResponse, _err error) {
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
	_result = &DeleteCommentResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteComment"), tea.String("exclusive_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/exclusive/publishers/"+tea.StringValue(publisherId)+"/comments/"+tea.StringValue(commentId)), tea.String("boolean"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAllLabelableDepts() (_result *GetAllLabelableDeptsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAllLabelableDeptsHeaders{}
	_result = &GetAllLabelableDeptsResponse{}
	_body, _err := client.GetAllLabelableDeptsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAllLabelableDeptsWithOptions(headers *GetAllLabelableDeptsHeaders, runtime *util.RuntimeOptions) (_result *GetAllLabelableDeptsResponse, _err error) {
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
	_result = &GetAllLabelableDeptsResponse{}
	_body, _err := client.DoROARequest(tea.String("GetAllLabelableDepts"), tea.String("exclusive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/exclusive/partnerDepartments"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTrustDeviceList(request *GetTrustDeviceListRequest) (_result *GetTrustDeviceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTrustDeviceListHeaders{}
	_result = &GetTrustDeviceListResponse{}
	_body, _err := client.GetTrustDeviceListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTrustDeviceListWithOptions(request *GetTrustDeviceListRequest, headers *GetTrustDeviceListHeaders, runtime *util.RuntimeOptions) (_result *GetTrustDeviceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
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
	_result = &GetTrustDeviceListResponse{}
	_body, _err := client.DoROARequest(tea.String("GetTrustDeviceList"), tea.String("exclusive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/exclusive/trustedDevices/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchOrgInnerGroupInfo(request *SearchOrgInnerGroupInfoRequest) (_result *SearchOrgInnerGroupInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchOrgInnerGroupInfoHeaders{}
	_result = &SearchOrgInnerGroupInfoResponse{}
	_body, _err := client.SearchOrgInnerGroupInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchOrgInnerGroupInfoWithOptions(request *SearchOrgInnerGroupInfoRequest, headers *SearchOrgInnerGroupInfoHeaders, runtime *util.RuntimeOptions) (_result *SearchOrgInnerGroupInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupMembersCountEnd)) {
		query["groupMembersCountEnd"] = request.GroupMembersCountEnd
	}

	if !tea.BoolValue(util.IsUnset(request.SyncToDingpan)) {
		query["syncToDingpan"] = request.SyncToDingpan
	}

	if !tea.BoolValue(util.IsUnset(request.GroupOwner)) {
		query["groupOwner"] = request.GroupOwner
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeEnd)) {
		query["createTimeEnd"] = request.CreateTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeStart)) {
		query["createTimeStart"] = request.CreateTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["uuid"] = request.Uuid
	}

	if !tea.BoolValue(util.IsUnset(request.GroupMembersCountStart)) {
		query["groupMembersCountStart"] = request.GroupMembersCountStart
	}

	if !tea.BoolValue(util.IsUnset(request.LastActiveTimeEnd)) {
		query["lastActiveTimeEnd"] = request.LastActiveTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		query["operatorUserId"] = request.OperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.PageStart)) {
		query["pageStart"] = request.PageStart
	}

	if !tea.BoolValue(util.IsUnset(request.LastActiveTimeStart)) {
		query["lastActiveTimeStart"] = request.LastActiveTimeStart
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
	_result = &SearchOrgInnerGroupInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("SearchOrgInnerGroupInfo"), tea.String("exclusive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/exclusive/securities/orgGroupInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTrustedDevice(request *CreateTrustedDeviceRequest) (_result *CreateTrustedDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTrustedDeviceHeaders{}
	_result = &CreateTrustedDeviceResponse{}
	_body, _err := client.CreateTrustedDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTrustedDeviceWithOptions(request *CreateTrustedDeviceRequest, headers *CreateTrustedDeviceHeaders, runtime *util.RuntimeOptions) (_result *CreateTrustedDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		body["platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.MacAddress)) {
		body["macAddress"] = request.MacAddress
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
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
	_result = &CreateTrustedDeviceResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateTrustedDevice"), tea.String("exclusive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/exclusive/trustedDevices"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGroupActiveInfo(request *GetGroupActiveInfoRequest) (_result *GetGroupActiveInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetGroupActiveInfoHeaders{}
	_result = &GetGroupActiveInfoResponse{}
	_body, _err := client.GetGroupActiveInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGroupActiveInfoWithOptions(request *GetGroupActiveInfoRequest, headers *GetGroupActiveInfoHeaders, runtime *util.RuntimeOptions) (_result *GetGroupActiveInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	if !tea.BoolValue(util.IsUnset(request.DingGroupId)) {
		query["dingGroupId"] = request.DingGroupId
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
	_result = &GetGroupActiveInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetGroupActiveInfo"), tea.String("exclusive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/exclusive/data/activeGroups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCommentList(publisherId *string, request *GetCommentListRequest) (_result *GetCommentListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCommentListHeaders{}
	_result = &GetCommentListResponse{}
	_body, _err := client.GetCommentListWithOptions(publisherId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCommentListWithOptions(publisherId *string, request *GetCommentListRequest, headers *GetCommentListHeaders, runtime *util.RuntimeOptions) (_result *GetCommentListResponse, _err error) {
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
	_result = &GetCommentListResponse{}
	_body, _err := client.DoROARequest(tea.String("GetCommentList"), tea.String("exclusive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/exclusive/publishers/"+tea.StringValue(publisherId)+"/comments/list"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPartnerTypeByParentId(parentId *string) (_result *GetPartnerTypeByParentIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPartnerTypeByParentIdHeaders{}
	_result = &GetPartnerTypeByParentIdResponse{}
	_body, _err := client.GetPartnerTypeByParentIdWithOptions(parentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPartnerTypeByParentIdWithOptions(parentId *string, headers *GetPartnerTypeByParentIdHeaders, runtime *util.RuntimeOptions) (_result *GetPartnerTypeByParentIdResponse, _err error) {
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
	_result = &GetPartnerTypeByParentIdResponse{}
	_body, _err := client.DoROARequest(tea.String("GetPartnerTypeByParentId"), tea.String("exclusive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/exclusive/partnerLabels/"+tea.StringValue(parentId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDeptPartnerTypeAndNum(request *SetDeptPartnerTypeAndNumRequest) (_result *SetDeptPartnerTypeAndNumResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetDeptPartnerTypeAndNumHeaders{}
	_result = &SetDeptPartnerTypeAndNumResponse{}
	_body, _err := client.SetDeptPartnerTypeAndNumWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDeptPartnerTypeAndNumWithOptions(request *SetDeptPartnerTypeAndNumRequest, headers *SetDeptPartnerTypeAndNumHeaders, runtime *util.RuntimeOptions) (_result *SetDeptPartnerTypeAndNumResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.PartnerNum)) {
		body["partnerNum"] = request.PartnerNum
	}

	if !tea.BoolValue(util.IsUnset(request.LabelIds)) {
		body["labelIds"] = request.LabelIds
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
	_result = &SetDeptPartnerTypeAndNumResponse{}
	_body, _err := client.DoROARequest(tea.String("SetDeptPartnerTypeAndNum"), tea.String("exclusive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/exclusive/partnerDepartments"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
