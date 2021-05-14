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
