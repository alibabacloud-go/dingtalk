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
	JsVersion *int64 `json:"jsVersion,omitempty" xml:"jsVersion,omitempty"`
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

func (s *BatchCreateRequest) SetJsVersion(v int64) *BatchCreateRequest {
	s.JsVersion = &v
	return s
}

type BatchCreateRequestData struct {
	// 是否可以补卡
	CanReissueCard *bool `json:"canReissueCard,omitempty" xml:"canReissueCard,omitempty"`
	// 打卡周期,单位为天
	CardCycle *int64 `json:"cardCycle,omitempty" xml:"cardCycle,omitempty"`
	// 打卡的频次设置："cardFrequency":[             1,//周天             2,//周一             3,//周二             4,//周三             5,//周四             6,//周五             7//周六         ]
	CardFrequency         []*int64                                       `json:"cardFrequency,omitempty" xml:"cardFrequency,omitempty" type:"Repeated"`
	CardRuleItemParamList []*BatchCreateRequestDataCardRuleItemParamList `json:"cardRuleItemParamList,omitempty" xml:"cardRuleItemParamList,omitempty" type:"Repeated"`
	// 班级列表
	ClassIds []*string `json:"classIds,omitempty" xml:"classIds,omitempty" type:"Repeated"`
	// 班级名称列表
	ClassNames []*string `json:"classNames,omitempty" xml:"classNames,omitempty" type:"Repeated"`
	// 打卡的内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 卡片生效时间
	EffectDate *float32 `json:"effectDate,omitempty" xml:"effectDate,omitempty"`
	// 上传相册，图片，录音，盯盘的信息
	Medias *string `json:"medias,omitempty" xml:"medias,omitempty"`
	// 计量开启
	NeedMetering             *bool                                             `json:"needMetering,omitempty" xml:"needMetering,omitempty"`
	OrgClassStudentGroupList []*BatchCreateRequestDataOrgClassStudentGroupList `json:"orgClassStudentGroupList,omitempty" xml:"orgClassStudentGroupList,omitempty" type:"Repeated"`
	// 提醒时间（小时）
	RemindHour *int64 `json:"remindHour,omitempty" xml:"remindHour,omitempty"`
	// 提醒时间（分钟）
	RemindMinute *int64 `json:"remindMinute,omitempty" xml:"remindMinute,omitempty"`
	// 默认：student_guardian
	TargetRole *string `json:"targetRole,omitempty" xml:"targetRole,omitempty"`
	// 打卡模板id
	TemplateId *float32 `json:"templateId,omitempty" xml:"templateId,omitempty"`
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

func (s *BatchCreateRequestData) SetCardCycle(v int64) *BatchCreateRequestData {
	s.CardCycle = &v
	return s
}

func (s *BatchCreateRequestData) SetCardFrequency(v []*int64) *BatchCreateRequestData {
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

func (s *BatchCreateRequestData) SetEffectDate(v float32) *BatchCreateRequestData {
	s.EffectDate = &v
	return s
}

func (s *BatchCreateRequestData) SetMedias(v string) *BatchCreateRequestData {
	s.Medias = &v
	return s
}

func (s *BatchCreateRequestData) SetNeedMetering(v bool) *BatchCreateRequestData {
	s.NeedMetering = &v
	return s
}

func (s *BatchCreateRequestData) SetOrgClassStudentGroupList(v []*BatchCreateRequestDataOrgClassStudentGroupList) *BatchCreateRequestData {
	s.OrgClassStudentGroupList = v
	return s
}

func (s *BatchCreateRequestData) SetRemindHour(v int64) *BatchCreateRequestData {
	s.RemindHour = &v
	return s
}

func (s *BatchCreateRequestData) SetRemindMinute(v int64) *BatchCreateRequestData {
	s.RemindMinute = &v
	return s
}

func (s *BatchCreateRequestData) SetTargetRole(v string) *BatchCreateRequestData {
	s.TargetRole = &v
	return s
}

func (s *BatchCreateRequestData) SetTemplateId(v float32) *BatchCreateRequestData {
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
	DailyDubbing *int64 `json:"dailyDubbing,omitempty" xml:"dailyDubbing,omitempty"`
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

func (s *BatchCreateRequestDataCardRuleItemParamList) SetDailyDubbing(v int64) *BatchCreateRequestDataCardRuleItemParamList {
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
	ClassId *float32 `json:"classId,omitempty" xml:"classId,omitempty"`
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

func (s *BatchCreateRequestDataOrgClassStudentGroupListClassList) SetClassId(v float32) *BatchCreateRequestDataOrgClassStudentGroupListClassList {
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
	StuName *string `json:"stuName,omitempty" xml:"stuName,omitempty"`
	// 学生id
	StuId *string `json:"stuId,omitempty" xml:"stuId,omitempty"`
}

func (s BatchCreateRequestDataOrgClassStudentGroupListClassListStudents) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateRequestDataOrgClassStudentGroupListClassListStudents) GoString() string {
	return s.String()
}

func (s *BatchCreateRequestDataOrgClassStudentGroupListClassListStudents) SetStuName(v string) *BatchCreateRequestDataOrgClassStudentGroupListClassListStudents {
	s.StuName = &v
	return s
}

func (s *BatchCreateRequestDataOrgClassStudentGroupListClassListStudents) SetStuId(v string) *BatchCreateRequestDataOrgClassStudentGroupListClassListStudents {
	s.StuId = &v
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
	HwMedia            *string                                      `json:"hwMedia,omitempty" xml:"hwMedia,omitempty"`
	HwContent          *string                                      `json:"hwContent,omitempty" xml:"hwContent,omitempty"`
	HwTitle            *string                                      `json:"hwTitle,omitempty" xml:"hwTitle,omitempty"`
	CourseName         *string                                      `json:"courseName,omitempty" xml:"courseName,omitempty"`
	HwPhoto            *string                                      `json:"hwPhoto,omitempty" xml:"hwPhoto,omitempty"`
	HwVideo            *string                                      `json:"hwVideo,omitempty" xml:"hwVideo,omitempty"`
	TeacherName        *string                                      `json:"teacherName,omitempty" xml:"teacherName,omitempty"`
	TeacherUserId      *string                                      `json:"teacherUserId,omitempty" xml:"teacherUserId,omitempty"`
	Identifier         *string                                      `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Attributes         *string                                      `json:"attributes,omitempty" xml:"attributes,omitempty"`
	TargetRole         *string                                      `json:"targetRole,omitempty" xml:"targetRole,omitempty"`
	BizCode            *string                                      `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	Status             *string                                      `json:"status,omitempty" xml:"status,omitempty"`
	ScheduledRelease   *string                                      `json:"scheduledRelease,omitempty" xml:"scheduledRelease,omitempty"`
	ScheduledTime      *string                                      `json:"scheduledTime,omitempty" xml:"scheduledTime,omitempty"`
	HwDeadlineOpen     *string                                      `json:"hwDeadlineOpen,omitempty" xml:"hwDeadlineOpen,omitempty"`
	HwDeadline         *float32                                     `json:"hwDeadline,omitempty" xml:"hwDeadline,omitempty"`
	HwType             *string                                      `json:"hwType,omitempty" xml:"hwType,omitempty"`
	OpenSelectItemList []*BatchOrgCreateHWRequestOpenSelectItemList `json:"openSelectItemList,omitempty" xml:"openSelectItemList,omitempty" type:"Repeated"`
	DingOrgId          *int64                                       `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
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

func (s *BatchOrgCreateHWRequest) SetHwDeadline(v float32) *BatchOrgCreateHWRequest {
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
	CorpId              *string                                               `json:"corpId,omitempty" xml:"corpId,omitempty"`
	SelectedClassesDesc *string                                               `json:"selectedClassesDesc,omitempty" xml:"selectedClassesDesc,omitempty"`
	ClassList           []*BatchOrgCreateHWRequestOpenSelectItemListClassList `json:"classList,omitempty" xml:"classList,omitempty" type:"Repeated"`
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
	ClassId   *string                                                       `json:"classId,omitempty" xml:"classId,omitempty"`
	ClassName *string                                                       `json:"className,omitempty" xml:"className,omitempty"`
	All       *bool                                                         `json:"all,omitempty" xml:"all,omitempty"`
	Students  []*BatchOrgCreateHWRequestOpenSelectItemListClassListStudents `json:"students,omitempty" xml:"students,omitempty" type:"Repeated"`
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
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	StaffId *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
	Avatar  *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
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
	Corpid *string  `json:"corpid,omitempty" xml:"corpid,omitempty"`
	Hwid   *float32 `json:"hwid,omitempty" xml:"hwid,omitempty"`
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

func (s *BatchOrgCreateHWResponseBodyResultPublishList) SetHwid(v float32) *BatchOrgCreateHWResponseBodyResultPublishList {
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
