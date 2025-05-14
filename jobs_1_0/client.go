// This file is auto-generated, don't edit it. Thanks.
package jobs_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateResumeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateResumeHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateResumeHeaders) GoString() string {
	return s.String()
}

func (s *CreateResumeHeaders) SetCommonHeaders(v map[string]*string) *CreateResumeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateResumeHeaders) SetXAcsDingtalkAccessToken(v string) *CreateResumeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateResumeRequest struct {
	BizCode      *string                          `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	Ext          *string                          `json:"ext,omitempty" xml:"ext,omitempty"`
	ResumeDataVO *CreateResumeRequestResumeDataVO `json:"resumeDataVO,omitempty" xml:"resumeDataVO,omitempty" type:"Struct"`
	Scene        *string                          `json:"scene,omitempty" xml:"scene,omitempty"`
	Types        []*string                        `json:"types,omitempty" xml:"types,omitempty" type:"Repeated"`
	// if can be null:
	// false
	UserIdentify *string `json:"userIdentify,omitempty" xml:"userIdentify,omitempty"`
}

func (s CreateResumeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResumeRequest) GoString() string {
	return s.String()
}

func (s *CreateResumeRequest) SetBizCode(v string) *CreateResumeRequest {
	s.BizCode = &v
	return s
}

func (s *CreateResumeRequest) SetExt(v string) *CreateResumeRequest {
	s.Ext = &v
	return s
}

func (s *CreateResumeRequest) SetResumeDataVO(v *CreateResumeRequestResumeDataVO) *CreateResumeRequest {
	s.ResumeDataVO = v
	return s
}

func (s *CreateResumeRequest) SetScene(v string) *CreateResumeRequest {
	s.Scene = &v
	return s
}

func (s *CreateResumeRequest) SetTypes(v []*string) *CreateResumeRequest {
	s.Types = v
	return s
}

func (s *CreateResumeRequest) SetUserIdentify(v string) *CreateResumeRequest {
	s.UserIdentify = &v
	return s
}

type CreateResumeRequestResumeDataVO struct {
	BaseInfo           *CreateResumeRequestResumeDataVOBaseInfo             `json:"baseInfo,omitempty" xml:"baseInfo,omitempty" type:"Struct"`
	Certificates       []*CreateResumeRequestResumeDataVOCertificates       `json:"certificates,omitempty" xml:"certificates,omitempty" type:"Repeated"`
	JobExpects         []*CreateResumeRequestResumeDataVOJobExpects         `json:"jobExpects,omitempty" xml:"jobExpects,omitempty" type:"Repeated"`
	PersonalHonors     []*CreateResumeRequestResumeDataVOPersonalHonors     `json:"personalHonors,omitempty" xml:"personalHonors,omitempty" type:"Repeated"`
	ProjectExperiences []*CreateResumeRequestResumeDataVOProjectExperiences `json:"projectExperiences,omitempty" xml:"projectExperiences,omitempty" type:"Repeated"`
	Tags               []*CreateResumeRequestResumeDataVOTags               `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	WorkExperiences    []*CreateResumeRequestResumeDataVOWorkExperiences    `json:"workExperiences,omitempty" xml:"workExperiences,omitempty" type:"Repeated"`
}

func (s CreateResumeRequestResumeDataVO) String() string {
	return tea.Prettify(s)
}

func (s CreateResumeRequestResumeDataVO) GoString() string {
	return s.String()
}

func (s *CreateResumeRequestResumeDataVO) SetBaseInfo(v *CreateResumeRequestResumeDataVOBaseInfo) *CreateResumeRequestResumeDataVO {
	s.BaseInfo = v
	return s
}

func (s *CreateResumeRequestResumeDataVO) SetCertificates(v []*CreateResumeRequestResumeDataVOCertificates) *CreateResumeRequestResumeDataVO {
	s.Certificates = v
	return s
}

func (s *CreateResumeRequestResumeDataVO) SetJobExpects(v []*CreateResumeRequestResumeDataVOJobExpects) *CreateResumeRequestResumeDataVO {
	s.JobExpects = v
	return s
}

func (s *CreateResumeRequestResumeDataVO) SetPersonalHonors(v []*CreateResumeRequestResumeDataVOPersonalHonors) *CreateResumeRequestResumeDataVO {
	s.PersonalHonors = v
	return s
}

func (s *CreateResumeRequestResumeDataVO) SetProjectExperiences(v []*CreateResumeRequestResumeDataVOProjectExperiences) *CreateResumeRequestResumeDataVO {
	s.ProjectExperiences = v
	return s
}

func (s *CreateResumeRequestResumeDataVO) SetTags(v []*CreateResumeRequestResumeDataVOTags) *CreateResumeRequestResumeDataVO {
	s.Tags = v
	return s
}

func (s *CreateResumeRequestResumeDataVO) SetWorkExperiences(v []*CreateResumeRequestResumeDataVOWorkExperiences) *CreateResumeRequestResumeDataVO {
	s.WorkExperiences = v
	return s
}

type CreateResumeRequestResumeDataVOBaseInfo struct {
	Age                 *int64    `json:"age,omitempty" xml:"age,omitempty"`
	Avatar              *string   `json:"avatar,omitempty" xml:"avatar,omitempty"`
	BeginWorkTime       *string   `json:"beginWorkTime,omitempty" xml:"beginWorkTime,omitempty"`
	Birthday            *string   `json:"birthday,omitempty" xml:"birthday,omitempty"`
	CandidateBackground *int32    `json:"candidateBackground,omitempty" xml:"candidateBackground,omitempty"`
	DingTalk            *string   `json:"dingTalk,omitempty" xml:"dingTalk,omitempty"`
	Email               *string   `json:"email,omitempty" xml:"email,omitempty"`
	EnglishName         *string   `json:"englishName,omitempty" xml:"englishName,omitempty"`
	Ethnicity           *string   `json:"ethnicity,omitempty" xml:"ethnicity,omitempty"`
	GaduateTime         *string   `json:"gaduateTime,omitempty" xml:"gaduateTime,omitempty"`
	HighestAcademic     *string   `json:"highestAcademic,omitempty" xml:"highestAcademic,omitempty"`
	HighestEducation    *int32    `json:"highestEducation,omitempty" xml:"highestEducation,omitempty"`
	Identify            *string   `json:"identify,omitempty" xml:"identify,omitempty"`
	Industry            *string   `json:"industry,omitempty" xml:"industry,omitempty"`
	IndustryCode        *string   `json:"industryCode,omitempty" xml:"industryCode,omitempty"`
	JobTitle            *string   `json:"jobTitle,omitempty" xml:"jobTitle,omitempty"`
	LastSchoolName      *string   `json:"lastSchoolName,omitempty" xml:"lastSchoolName,omitempty"`
	Married             *int64    `json:"married,omitempty" xml:"married,omitempty"`
	MbtiType            *int32    `json:"mbtiType,omitempty" xml:"mbtiType,omitempty"`
	Name                *string   `json:"name,omitempty" xml:"name,omitempty"`
	Nationality         *string   `json:"nationality,omitempty" xml:"nationality,omitempty"`
	NativePlace         *string   `json:"nativePlace,omitempty" xml:"nativePlace,omitempty"`
	NativePlaceCode     *string   `json:"nativePlaceCode,omitempty" xml:"nativePlaceCode,omitempty"`
	NowLocation         *string   `json:"nowLocation,omitempty" xml:"nowLocation,omitempty"`
	NowLocationCode     *string   `json:"nowLocationCode,omitempty" xml:"nowLocationCode,omitempty"`
	ParentIndustry      *string   `json:"parentIndustry,omitempty" xml:"parentIndustry,omitempty"`
	ParentIndustryCode  *string   `json:"parentIndustryCode,omitempty" xml:"parentIndustryCode,omitempty"`
	PersonalHonor       *string   `json:"personalHonor,omitempty" xml:"personalHonor,omitempty"`
	PersonalUrls        []*string `json:"personalUrls,omitempty" xml:"personalUrls,omitempty" type:"Repeated"`
	PhoneNum            *string   `json:"phoneNum,omitempty" xml:"phoneNum,omitempty"`
	PoliticalStatus     *int32    `json:"politicalStatus,omitempty" xml:"politicalStatus,omitempty"`
	Qq                  *string   `json:"qq,omitempty" xml:"qq,omitempty"`
	RealAvatar          *int32    `json:"realAvatar,omitempty" xml:"realAvatar,omitempty"`
	SelfEvaluation      *string   `json:"selfEvaluation,omitempty" xml:"selfEvaluation,omitempty"`
	Sex                 *int32    `json:"sex,omitempty" xml:"sex,omitempty"`
	SkillSummary        *string   `json:"skillSummary,omitempty" xml:"skillSummary,omitempty"`
	StateCode           *string   `json:"stateCode,omitempty" xml:"stateCode,omitempty"`
	Status              *string   `json:"status,omitempty" xml:"status,omitempty"`
	VirtualPhoneNum     *string   `json:"virtualPhoneNum,omitempty" xml:"virtualPhoneNum,omitempty"`
	WeChat              *string   `json:"weChat,omitempty" xml:"weChat,omitempty"`
	Weibo               *string   `json:"weibo,omitempty" xml:"weibo,omitempty"`
	WorkingYears        *int32    `json:"workingYears,omitempty" xml:"workingYears,omitempty"`
}

func (s CreateResumeRequestResumeDataVOBaseInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateResumeRequestResumeDataVOBaseInfo) GoString() string {
	return s.String()
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetAge(v int64) *CreateResumeRequestResumeDataVOBaseInfo {
	s.Age = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetAvatar(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.Avatar = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetBeginWorkTime(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.BeginWorkTime = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetBirthday(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.Birthday = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetCandidateBackground(v int32) *CreateResumeRequestResumeDataVOBaseInfo {
	s.CandidateBackground = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetDingTalk(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.DingTalk = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetEmail(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.Email = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetEnglishName(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.EnglishName = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetEthnicity(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.Ethnicity = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetGaduateTime(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.GaduateTime = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetHighestAcademic(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.HighestAcademic = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetHighestEducation(v int32) *CreateResumeRequestResumeDataVOBaseInfo {
	s.HighestEducation = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetIdentify(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.Identify = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetIndustry(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.Industry = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetIndustryCode(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.IndustryCode = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetJobTitle(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.JobTitle = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetLastSchoolName(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.LastSchoolName = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetMarried(v int64) *CreateResumeRequestResumeDataVOBaseInfo {
	s.Married = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetMbtiType(v int32) *CreateResumeRequestResumeDataVOBaseInfo {
	s.MbtiType = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetName(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.Name = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetNationality(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.Nationality = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetNativePlace(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.NativePlace = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetNativePlaceCode(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.NativePlaceCode = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetNowLocation(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.NowLocation = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetNowLocationCode(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.NowLocationCode = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetParentIndustry(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.ParentIndustry = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetParentIndustryCode(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.ParentIndustryCode = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetPersonalHonor(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.PersonalHonor = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetPersonalUrls(v []*string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.PersonalUrls = v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetPhoneNum(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.PhoneNum = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetPoliticalStatus(v int32) *CreateResumeRequestResumeDataVOBaseInfo {
	s.PoliticalStatus = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetQq(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.Qq = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetRealAvatar(v int32) *CreateResumeRequestResumeDataVOBaseInfo {
	s.RealAvatar = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetSelfEvaluation(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.SelfEvaluation = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetSex(v int32) *CreateResumeRequestResumeDataVOBaseInfo {
	s.Sex = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetSkillSummary(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.SkillSummary = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetStateCode(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.StateCode = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetStatus(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.Status = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetVirtualPhoneNum(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.VirtualPhoneNum = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetWeChat(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.WeChat = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetWeibo(v string) *CreateResumeRequestResumeDataVOBaseInfo {
	s.Weibo = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOBaseInfo) SetWorkingYears(v int32) *CreateResumeRequestResumeDataVOBaseInfo {
	s.WorkingYears = &v
	return s
}

type CreateResumeRequestResumeDataVOCertificates struct {
	CertificateId   *string `json:"certificateId,omitempty" xml:"certificateId,omitempty"`
	CertificateName *string `json:"certificateName,omitempty" xml:"certificateName,omitempty"`
	CrantTime       *string `json:"crantTime,omitempty" xml:"crantTime,omitempty"`
}

func (s CreateResumeRequestResumeDataVOCertificates) String() string {
	return tea.Prettify(s)
}

func (s CreateResumeRequestResumeDataVOCertificates) GoString() string {
	return s.String()
}

func (s *CreateResumeRequestResumeDataVOCertificates) SetCertificateId(v string) *CreateResumeRequestResumeDataVOCertificates {
	s.CertificateId = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOCertificates) SetCertificateName(v string) *CreateResumeRequestResumeDataVOCertificates {
	s.CertificateName = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOCertificates) SetCrantTime(v string) *CreateResumeRequestResumeDataVOCertificates {
	s.CrantTime = &v
	return s
}

type CreateResumeRequestResumeDataVOJobExpects struct {
	CityList         []*CreateResumeRequestResumeDataVOJobExpectsCityList      `json:"cityList,omitempty" xml:"cityList,omitempty" type:"Repeated"`
	GmtCreate        *int64                                                    `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified      *int64                                                    `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	IndustryList     []*CreateResumeRequestResumeDataVOJobExpectsIndustryList  `json:"industryList,omitempty" xml:"industryList,omitempty" type:"Repeated"`
	JobList          []*CreateResumeRequestResumeDataVOJobExpectsJobList       `json:"jobList,omitempty" xml:"jobList,omitempty" type:"Repeated"`
	JobNature        *string                                                   `json:"jobNature,omitempty" xml:"jobNature,omitempty"`
	MaxSalary        *string                                                   `json:"maxSalary,omitempty" xml:"maxSalary,omitempty"`
	MinSalary        *string                                                   `json:"minSalary,omitempty" xml:"minSalary,omitempty"`
	OtherCityList    []*CreateResumeRequestResumeDataVOJobExpectsOtherCityList `json:"otherCityList,omitempty" xml:"otherCityList,omitempty" type:"Repeated"`
	SalaryDesc       *string                                                   `json:"salaryDesc,omitempty" xml:"salaryDesc,omitempty"`
	SalarySettleType *string                                                   `json:"salarySettleType,omitempty" xml:"salarySettleType,omitempty"`
	SalaryType       *string                                                   `json:"salaryType,omitempty" xml:"salaryType,omitempty"`
	SalaryYear       *string                                                   `json:"salaryYear,omitempty" xml:"salaryYear,omitempty"`
}

func (s CreateResumeRequestResumeDataVOJobExpects) String() string {
	return tea.Prettify(s)
}

func (s CreateResumeRequestResumeDataVOJobExpects) GoString() string {
	return s.String()
}

func (s *CreateResumeRequestResumeDataVOJobExpects) SetCityList(v []*CreateResumeRequestResumeDataVOJobExpectsCityList) *CreateResumeRequestResumeDataVOJobExpects {
	s.CityList = v
	return s
}

func (s *CreateResumeRequestResumeDataVOJobExpects) SetGmtCreate(v int64) *CreateResumeRequestResumeDataVOJobExpects {
	s.GmtCreate = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOJobExpects) SetGmtModified(v int64) *CreateResumeRequestResumeDataVOJobExpects {
	s.GmtModified = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOJobExpects) SetIndustryList(v []*CreateResumeRequestResumeDataVOJobExpectsIndustryList) *CreateResumeRequestResumeDataVOJobExpects {
	s.IndustryList = v
	return s
}

func (s *CreateResumeRequestResumeDataVOJobExpects) SetJobList(v []*CreateResumeRequestResumeDataVOJobExpectsJobList) *CreateResumeRequestResumeDataVOJobExpects {
	s.JobList = v
	return s
}

func (s *CreateResumeRequestResumeDataVOJobExpects) SetJobNature(v string) *CreateResumeRequestResumeDataVOJobExpects {
	s.JobNature = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOJobExpects) SetMaxSalary(v string) *CreateResumeRequestResumeDataVOJobExpects {
	s.MaxSalary = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOJobExpects) SetMinSalary(v string) *CreateResumeRequestResumeDataVOJobExpects {
	s.MinSalary = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOJobExpects) SetOtherCityList(v []*CreateResumeRequestResumeDataVOJobExpectsOtherCityList) *CreateResumeRequestResumeDataVOJobExpects {
	s.OtherCityList = v
	return s
}

func (s *CreateResumeRequestResumeDataVOJobExpects) SetSalaryDesc(v string) *CreateResumeRequestResumeDataVOJobExpects {
	s.SalaryDesc = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOJobExpects) SetSalarySettleType(v string) *CreateResumeRequestResumeDataVOJobExpects {
	s.SalarySettleType = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOJobExpects) SetSalaryType(v string) *CreateResumeRequestResumeDataVOJobExpects {
	s.SalaryType = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOJobExpects) SetSalaryYear(v string) *CreateResumeRequestResumeDataVOJobExpects {
	s.SalaryYear = &v
	return s
}

type CreateResumeRequestResumeDataVOJobExpectsCityList struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateResumeRequestResumeDataVOJobExpectsCityList) String() string {
	return tea.Prettify(s)
}

func (s CreateResumeRequestResumeDataVOJobExpectsCityList) GoString() string {
	return s.String()
}

func (s *CreateResumeRequestResumeDataVOJobExpectsCityList) SetCode(v string) *CreateResumeRequestResumeDataVOJobExpectsCityList {
	s.Code = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOJobExpectsCityList) SetName(v string) *CreateResumeRequestResumeDataVOJobExpectsCityList {
	s.Name = &v
	return s
}

type CreateResumeRequestResumeDataVOJobExpectsIndustryList struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateResumeRequestResumeDataVOJobExpectsIndustryList) String() string {
	return tea.Prettify(s)
}

func (s CreateResumeRequestResumeDataVOJobExpectsIndustryList) GoString() string {
	return s.String()
}

func (s *CreateResumeRequestResumeDataVOJobExpectsIndustryList) SetCode(v string) *CreateResumeRequestResumeDataVOJobExpectsIndustryList {
	s.Code = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOJobExpectsIndustryList) SetName(v string) *CreateResumeRequestResumeDataVOJobExpectsIndustryList {
	s.Name = &v
	return s
}

type CreateResumeRequestResumeDataVOJobExpectsJobList struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateResumeRequestResumeDataVOJobExpectsJobList) String() string {
	return tea.Prettify(s)
}

func (s CreateResumeRequestResumeDataVOJobExpectsJobList) GoString() string {
	return s.String()
}

func (s *CreateResumeRequestResumeDataVOJobExpectsJobList) SetCode(v string) *CreateResumeRequestResumeDataVOJobExpectsJobList {
	s.Code = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOJobExpectsJobList) SetName(v string) *CreateResumeRequestResumeDataVOJobExpectsJobList {
	s.Name = &v
	return s
}

type CreateResumeRequestResumeDataVOJobExpectsOtherCityList struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateResumeRequestResumeDataVOJobExpectsOtherCityList) String() string {
	return tea.Prettify(s)
}

func (s CreateResumeRequestResumeDataVOJobExpectsOtherCityList) GoString() string {
	return s.String()
}

func (s *CreateResumeRequestResumeDataVOJobExpectsOtherCityList) SetCode(v string) *CreateResumeRequestResumeDataVOJobExpectsOtherCityList {
	s.Code = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOJobExpectsOtherCityList) SetName(v string) *CreateResumeRequestResumeDataVOJobExpectsOtherCityList {
	s.Name = &v
	return s
}

type CreateResumeRequestResumeDataVOPersonalHonors struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	GrantTime   *string `json:"grantTime,omitempty" xml:"grantTime,omitempty"`
}

func (s CreateResumeRequestResumeDataVOPersonalHonors) String() string {
	return tea.Prettify(s)
}

func (s CreateResumeRequestResumeDataVOPersonalHonors) GoString() string {
	return s.String()
}

func (s *CreateResumeRequestResumeDataVOPersonalHonors) SetDescription(v string) *CreateResumeRequestResumeDataVOPersonalHonors {
	s.Description = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOPersonalHonors) SetGrantTime(v string) *CreateResumeRequestResumeDataVOPersonalHonors {
	s.GrantTime = &v
	return s
}

type CreateResumeRequestResumeDataVOProjectExperiences struct {
	Achievement    *string `json:"achievement,omitempty" xml:"achievement,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	EndDate        *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	ProjectUrl     *string `json:"projectUrl,omitempty" xml:"projectUrl,omitempty"`
	Responsibility *string `json:"responsibility,omitempty" xml:"responsibility,omitempty"`
	StartDate      *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	Technology     *string `json:"technology,omitempty" xml:"technology,omitempty"`
	Title          *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateResumeRequestResumeDataVOProjectExperiences) String() string {
	return tea.Prettify(s)
}

func (s CreateResumeRequestResumeDataVOProjectExperiences) GoString() string {
	return s.String()
}

func (s *CreateResumeRequestResumeDataVOProjectExperiences) SetAchievement(v string) *CreateResumeRequestResumeDataVOProjectExperiences {
	s.Achievement = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOProjectExperiences) SetDescription(v string) *CreateResumeRequestResumeDataVOProjectExperiences {
	s.Description = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOProjectExperiences) SetEndDate(v string) *CreateResumeRequestResumeDataVOProjectExperiences {
	s.EndDate = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOProjectExperiences) SetName(v string) *CreateResumeRequestResumeDataVOProjectExperiences {
	s.Name = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOProjectExperiences) SetProjectUrl(v string) *CreateResumeRequestResumeDataVOProjectExperiences {
	s.ProjectUrl = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOProjectExperiences) SetResponsibility(v string) *CreateResumeRequestResumeDataVOProjectExperiences {
	s.Responsibility = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOProjectExperiences) SetStartDate(v string) *CreateResumeRequestResumeDataVOProjectExperiences {
	s.StartDate = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOProjectExperiences) SetTechnology(v string) *CreateResumeRequestResumeDataVOProjectExperiences {
	s.Technology = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOProjectExperiences) SetTitle(v string) *CreateResumeRequestResumeDataVOProjectExperiences {
	s.Title = &v
	return s
}

type CreateResumeRequestResumeDataVOTags struct {
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s CreateResumeRequestResumeDataVOTags) String() string {
	return tea.Prettify(s)
}

func (s CreateResumeRequestResumeDataVOTags) GoString() string {
	return s.String()
}

func (s *CreateResumeRequestResumeDataVOTags) SetTag(v string) *CreateResumeRequestResumeDataVOTags {
	s.Tag = &v
	return s
}

type CreateResumeRequestResumeDataVOWorkExperiences struct {
	Achievement          *string                                                      `json:"achievement,omitempty" xml:"achievement,omitempty"`
	CompanyCode          *string                                                      `json:"companyCode,omitempty" xml:"companyCode,omitempty"`
	CompanyName          *string                                                      `json:"companyName,omitempty" xml:"companyName,omitempty"`
	Description          *string                                                      `json:"description,omitempty" xml:"description,omitempty"`
	EndDate              *string                                                      `json:"endDate,omitempty" xml:"endDate,omitempty"`
	Industry             *string                                                      `json:"industry,omitempty" xml:"industry,omitempty"`
	IndustryCode         *string                                                      `json:"industryCode,omitempty" xml:"industryCode,omitempty"`
	Internship           *bool                                                        `json:"internship,omitempty" xml:"internship,omitempty"`
	JobCode              *string                                                      `json:"jobCode,omitempty" xml:"jobCode,omitempty"`
	JobNature            *string                                                      `json:"jobNature,omitempty" xml:"jobNature,omitempty"`
	JobTitle             *string                                                      `json:"jobTitle,omitempty" xml:"jobTitle,omitempty"`
	Leader               *string                                                      `json:"leader,omitempty" xml:"leader,omitempty"`
	Location             *string                                                      `json:"location,omitempty" xml:"location,omitempty"`
	LocationCode         *string                                                      `json:"locationCode,omitempty" xml:"locationCode,omitempty"`
	ParentIndustry       *string                                                      `json:"parentIndustry,omitempty" xml:"parentIndustry,omitempty"`
	ParentIndustryCode   *string                                                      `json:"parentIndustryCode,omitempty" xml:"parentIndustryCode,omitempty"`
	ReasonOfLeaving      *string                                                      `json:"reasonOfLeaving,omitempty" xml:"reasonOfLeaving,omitempty"`
	Responsibility       *string                                                      `json:"responsibility,omitempty" xml:"responsibility,omitempty"`
	ResumePrivacy        *CreateResumeRequestResumeDataVOWorkExperiencesResumePrivacy `json:"resumePrivacy,omitempty" xml:"resumePrivacy,omitempty" type:"Struct"`
	Salary               *string                                                      `json:"salary,omitempty" xml:"salary,omitempty"`
	SelectedSkillOptions []*string                                                    `json:"selectedSkillOptions,omitempty" xml:"selectedSkillOptions,omitempty" type:"Repeated"`
	StartDate            *string                                                      `json:"startDate,omitempty" xml:"startDate,omitempty"`
	UnderlingNumber      *string                                                      `json:"underlingNumber,omitempty" xml:"underlingNumber,omitempty"`
}

func (s CreateResumeRequestResumeDataVOWorkExperiences) String() string {
	return tea.Prettify(s)
}

func (s CreateResumeRequestResumeDataVOWorkExperiences) GoString() string {
	return s.String()
}

func (s *CreateResumeRequestResumeDataVOWorkExperiences) SetAchievement(v string) *CreateResumeRequestResumeDataVOWorkExperiences {
	s.Achievement = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOWorkExperiences) SetCompanyCode(v string) *CreateResumeRequestResumeDataVOWorkExperiences {
	s.CompanyCode = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOWorkExperiences) SetCompanyName(v string) *CreateResumeRequestResumeDataVOWorkExperiences {
	s.CompanyName = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOWorkExperiences) SetDescription(v string) *CreateResumeRequestResumeDataVOWorkExperiences {
	s.Description = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOWorkExperiences) SetEndDate(v string) *CreateResumeRequestResumeDataVOWorkExperiences {
	s.EndDate = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOWorkExperiences) SetIndustry(v string) *CreateResumeRequestResumeDataVOWorkExperiences {
	s.Industry = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOWorkExperiences) SetIndustryCode(v string) *CreateResumeRequestResumeDataVOWorkExperiences {
	s.IndustryCode = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOWorkExperiences) SetInternship(v bool) *CreateResumeRequestResumeDataVOWorkExperiences {
	s.Internship = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOWorkExperiences) SetJobCode(v string) *CreateResumeRequestResumeDataVOWorkExperiences {
	s.JobCode = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOWorkExperiences) SetJobNature(v string) *CreateResumeRequestResumeDataVOWorkExperiences {
	s.JobNature = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOWorkExperiences) SetJobTitle(v string) *CreateResumeRequestResumeDataVOWorkExperiences {
	s.JobTitle = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOWorkExperiences) SetLeader(v string) *CreateResumeRequestResumeDataVOWorkExperiences {
	s.Leader = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOWorkExperiences) SetLocation(v string) *CreateResumeRequestResumeDataVOWorkExperiences {
	s.Location = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOWorkExperiences) SetLocationCode(v string) *CreateResumeRequestResumeDataVOWorkExperiences {
	s.LocationCode = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOWorkExperiences) SetParentIndustry(v string) *CreateResumeRequestResumeDataVOWorkExperiences {
	s.ParentIndustry = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOWorkExperiences) SetParentIndustryCode(v string) *CreateResumeRequestResumeDataVOWorkExperiences {
	s.ParentIndustryCode = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOWorkExperiences) SetReasonOfLeaving(v string) *CreateResumeRequestResumeDataVOWorkExperiences {
	s.ReasonOfLeaving = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOWorkExperiences) SetResponsibility(v string) *CreateResumeRequestResumeDataVOWorkExperiences {
	s.Responsibility = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOWorkExperiences) SetResumePrivacy(v *CreateResumeRequestResumeDataVOWorkExperiencesResumePrivacy) *CreateResumeRequestResumeDataVOWorkExperiences {
	s.ResumePrivacy = v
	return s
}

func (s *CreateResumeRequestResumeDataVOWorkExperiences) SetSalary(v string) *CreateResumeRequestResumeDataVOWorkExperiences {
	s.Salary = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOWorkExperiences) SetSelectedSkillOptions(v []*string) *CreateResumeRequestResumeDataVOWorkExperiences {
	s.SelectedSkillOptions = v
	return s
}

func (s *CreateResumeRequestResumeDataVOWorkExperiences) SetStartDate(v string) *CreateResumeRequestResumeDataVOWorkExperiences {
	s.StartDate = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOWorkExperiences) SetUnderlingNumber(v string) *CreateResumeRequestResumeDataVOWorkExperiences {
	s.UnderlingNumber = &v
	return s
}

type CreateResumeRequestResumeDataVOWorkExperiencesResumePrivacy struct {
	ShieldedCompany        *bool `json:"shieldedCompany,omitempty" xml:"shieldedCompany,omitempty"`
	ShieldedRelatedCompany *bool `json:"shieldedRelatedCompany,omitempty" xml:"shieldedRelatedCompany,omitempty"`
}

func (s CreateResumeRequestResumeDataVOWorkExperiencesResumePrivacy) String() string {
	return tea.Prettify(s)
}

func (s CreateResumeRequestResumeDataVOWorkExperiencesResumePrivacy) GoString() string {
	return s.String()
}

func (s *CreateResumeRequestResumeDataVOWorkExperiencesResumePrivacy) SetShieldedCompany(v bool) *CreateResumeRequestResumeDataVOWorkExperiencesResumePrivacy {
	s.ShieldedCompany = &v
	return s
}

func (s *CreateResumeRequestResumeDataVOWorkExperiencesResumePrivacy) SetShieldedRelatedCompany(v bool) *CreateResumeRequestResumeDataVOWorkExperiencesResumePrivacy {
	s.ShieldedRelatedCompany = &v
	return s
}

type CreateResumeResponseBody struct {
	Result  *CreateResumeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateResumeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResumeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResumeResponseBody) SetResult(v *CreateResumeResponseBodyResult) *CreateResumeResponseBody {
	s.Result = v
	return s
}

func (s *CreateResumeResponseBody) SetSuccess(v bool) *CreateResumeResponseBody {
	s.Success = &v
	return s
}

type CreateResumeResponseBodyResult struct {
	ResumeId *string `json:"resumeId,omitempty" xml:"resumeId,omitempty"`
}

func (s CreateResumeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateResumeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateResumeResponseBodyResult) SetResumeId(v string) *CreateResumeResponseBodyResult {
	s.ResumeId = &v
	return s
}

type CreateResumeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResumeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResumeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResumeResponse) GoString() string {
	return s.String()
}

func (s *CreateResumeResponse) SetHeaders(v map[string]*string) *CreateResumeResponse {
	s.Headers = v
	return s
}

func (s *CreateResumeResponse) SetStatusCode(v int32) *CreateResumeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResumeResponse) SetBody(v *CreateResumeResponseBody) *CreateResumeResponse {
	s.Body = v
	return s
}

type PostResumeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PostResumeHeaders) String() string {
	return tea.Prettify(s)
}

func (s PostResumeHeaders) GoString() string {
	return s.String()
}

func (s *PostResumeHeaders) SetCommonHeaders(v map[string]*string) *PostResumeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PostResumeHeaders) SetXAcsDingtalkAccessToken(v string) *PostResumeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PostResumeRequest struct {
	JobId        *int64  `json:"jobId,omitempty" xml:"jobId,omitempty"`
	UserIdentify *string `json:"userIdentify,omitempty" xml:"userIdentify,omitempty"`
}

func (s PostResumeRequest) String() string {
	return tea.Prettify(s)
}

func (s PostResumeRequest) GoString() string {
	return s.String()
}

func (s *PostResumeRequest) SetJobId(v int64) *PostResumeRequest {
	s.JobId = &v
	return s
}

func (s *PostResumeRequest) SetUserIdentify(v string) *PostResumeRequest {
	s.UserIdentify = &v
	return s
}

type PostResumeResponseBody struct {
	JobId        *int64  `json:"jobId,omitempty" xml:"jobId,omitempty"`
	UserIdentify *string `json:"userIdentify,omitempty" xml:"userIdentify,omitempty"`
}

func (s PostResumeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PostResumeResponseBody) GoString() string {
	return s.String()
}

func (s *PostResumeResponseBody) SetJobId(v int64) *PostResumeResponseBody {
	s.JobId = &v
	return s
}

func (s *PostResumeResponseBody) SetUserIdentify(v string) *PostResumeResponseBody {
	s.UserIdentify = &v
	return s
}

type PostResumeResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostResumeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostResumeResponse) String() string {
	return tea.Prettify(s)
}

func (s PostResumeResponse) GoString() string {
	return s.String()
}

func (s *PostResumeResponse) SetHeaders(v map[string]*string) *PostResumeResponse {
	s.Headers = v
	return s
}

func (s *PostResumeResponse) SetStatusCode(v int32) *PostResumeResponse {
	s.StatusCode = &v
	return s
}

func (s *PostResumeResponse) SetBody(v *PostResumeResponseBody) *PostResumeResponse {
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
// 创建简历
//
// @param request - CreateResumeRequest
//
// @param headers - CreateResumeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResumeResponse
func (client *Client) CreateResumeWithOptions(request *CreateResumeRequest, headers *CreateResumeHeaders, runtime *util.RuntimeOptions) (_result *CreateResumeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		body["bizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.ResumeDataVO)) {
		body["resumeDataVO"] = request.ResumeDataVO
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.Types)) {
		body["types"] = request.Types
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdentify)) {
		body["userIdentify"] = request.UserIdentify
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
		Action:      tea.String("CreateResume"),
		Version:     tea.String("jobs_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/jobs/resumes"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateResumeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建简历
//
// @param request - CreateResumeRequest
//
// @return CreateResumeResponse
func (client *Client) CreateResume(request *CreateResumeRequest) (_result *CreateResumeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateResumeHeaders{}
	_result = &CreateResumeResponse{}
	_body, _err := client.CreateResumeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 投递简历
//
// @param request - PostResumeRequest
//
// @param headers - PostResumeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PostResumeResponse
func (client *Client) PostResumeWithOptions(request *PostResumeRequest, headers *PostResumeHeaders, runtime *util.RuntimeOptions) (_result *PostResumeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["jobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdentify)) {
		body["userIdentify"] = request.UserIdentify
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
		Action:      tea.String("PostResume"),
		Version:     tea.String("jobs_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/jobs/resumes/post"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PostResumeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 投递简历
//
// @param request - PostResumeRequest
//
// @return PostResumeResponse
func (client *Client) PostResume(request *PostResumeRequest) (_result *PostResumeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PostResumeHeaders{}
	_result = &PostResumeResponse{}
	_body, _err := client.PostResumeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
