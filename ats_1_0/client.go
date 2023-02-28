// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package ats_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddApplicationRegFormTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddApplicationRegFormTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddApplicationRegFormTemplateHeaders) GoString() string {
	return s.String()
}

func (s *AddApplicationRegFormTemplateHeaders) SetCommonHeaders(v map[string]*string) *AddApplicationRegFormTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddApplicationRegFormTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *AddApplicationRegFormTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddApplicationRegFormTemplateRequest struct {
	// 业务标识
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	// 模板内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 模板名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 外部唯一标识
	OuterId *string `json:"outerId,omitempty" xml:"outerId,omitempty"`
	// 操作人员工标识
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s AddApplicationRegFormTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s AddApplicationRegFormTemplateRequest) GoString() string {
	return s.String()
}

func (s *AddApplicationRegFormTemplateRequest) SetBizCode(v string) *AddApplicationRegFormTemplateRequest {
	s.BizCode = &v
	return s
}

func (s *AddApplicationRegFormTemplateRequest) SetContent(v string) *AddApplicationRegFormTemplateRequest {
	s.Content = &v
	return s
}

func (s *AddApplicationRegFormTemplateRequest) SetName(v string) *AddApplicationRegFormTemplateRequest {
	s.Name = &v
	return s
}

func (s *AddApplicationRegFormTemplateRequest) SetOuterId(v string) *AddApplicationRegFormTemplateRequest {
	s.OuterId = &v
	return s
}

func (s *AddApplicationRegFormTemplateRequest) SetOpUserId(v string) *AddApplicationRegFormTemplateRequest {
	s.OpUserId = &v
	return s
}

type AddApplicationRegFormTemplateResponseBody struct {
	// 模板标识
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// 模板版本
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s AddApplicationRegFormTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddApplicationRegFormTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *AddApplicationRegFormTemplateResponseBody) SetTemplateId(v string) *AddApplicationRegFormTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *AddApplicationRegFormTemplateResponseBody) SetVersion(v int32) *AddApplicationRegFormTemplateResponseBody {
	s.Version = &v
	return s
}

type AddApplicationRegFormTemplateResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddApplicationRegFormTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddApplicationRegFormTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s AddApplicationRegFormTemplateResponse) GoString() string {
	return s.String()
}

func (s *AddApplicationRegFormTemplateResponse) SetHeaders(v map[string]*string) *AddApplicationRegFormTemplateResponse {
	s.Headers = v
	return s
}

func (s *AddApplicationRegFormTemplateResponse) SetBody(v *AddApplicationRegFormTemplateResponseBody) *AddApplicationRegFormTemplateResponse {
	s.Body = v
	return s
}

type AddFileHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddFileHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddFileHeaders) GoString() string {
	return s.String()
}

func (s *AddFileHeaders) SetCommonHeaders(v map[string]*string) *AddFileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddFileHeaders) SetXAcsDingtalkAccessToken(v string) *AddFileHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddFileRequest struct {
	// 业务标识
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	// 文件名称
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件mediaId
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	// 操作人员工标识，为空时默认以企业管理员身份进行操作
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s AddFileRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFileRequest) GoString() string {
	return s.String()
}

func (s *AddFileRequest) SetBizCode(v string) *AddFileRequest {
	s.BizCode = &v
	return s
}

func (s *AddFileRequest) SetFileName(v string) *AddFileRequest {
	s.FileName = &v
	return s
}

func (s *AddFileRequest) SetMediaId(v string) *AddFileRequest {
	s.MediaId = &v
	return s
}

func (s *AddFileRequest) SetOpUserId(v string) *AddFileRequest {
	s.OpUserId = &v
	return s
}

type AddFileResponseBody struct {
	// 文件标识
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// 文件名
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 空间标识
	SpaceId *int64 `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s AddFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFileResponseBody) GoString() string {
	return s.String()
}

func (s *AddFileResponseBody) SetFileId(v string) *AddFileResponseBody {
	s.FileId = &v
	return s
}

func (s *AddFileResponseBody) SetFileName(v string) *AddFileResponseBody {
	s.FileName = &v
	return s
}

func (s *AddFileResponseBody) SetSpaceId(v int64) *AddFileResponseBody {
	s.SpaceId = &v
	return s
}

type AddFileResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddFileResponse) String() string {
	return tea.Prettify(s)
}

func (s AddFileResponse) GoString() string {
	return s.String()
}

func (s *AddFileResponse) SetHeaders(v map[string]*string) *AddFileResponse {
	s.Headers = v
	return s
}

func (s *AddFileResponse) SetBody(v *AddFileResponseBody) *AddFileResponse {
	s.Body = v
	return s
}

type AddUserAccountHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddUserAccountHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddUserAccountHeaders) GoString() string {
	return s.String()
}

func (s *AddUserAccountHeaders) SetCommonHeaders(v map[string]*string) *AddUserAccountHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddUserAccountHeaders) SetXAcsDingtalkAccessToken(v string) *AddUserAccountHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddUserAccountRequest struct {
	// 业务标识
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	// 渠道账号名
	ChannelAccountName *string `json:"channelAccountName,omitempty" xml:"channelAccountName,omitempty"`
	// 渠道用户标识
	ChannelUserIdentify *string `json:"channelUserIdentify,omitempty" xml:"channelUserIdentify,omitempty"`
	// 手机号
	PhoneNumber *string `json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	// 企业标识
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 用户标识
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AddUserAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserAccountRequest) GoString() string {
	return s.String()
}

func (s *AddUserAccountRequest) SetBizCode(v string) *AddUserAccountRequest {
	s.BizCode = &v
	return s
}

func (s *AddUserAccountRequest) SetChannelAccountName(v string) *AddUserAccountRequest {
	s.ChannelAccountName = &v
	return s
}

func (s *AddUserAccountRequest) SetChannelUserIdentify(v string) *AddUserAccountRequest {
	s.ChannelUserIdentify = &v
	return s
}

func (s *AddUserAccountRequest) SetPhoneNumber(v string) *AddUserAccountRequest {
	s.PhoneNumber = &v
	return s
}

func (s *AddUserAccountRequest) SetCorpId(v string) *AddUserAccountRequest {
	s.CorpId = &v
	return s
}

func (s *AddUserAccountRequest) SetUserId(v string) *AddUserAccountRequest {
	s.UserId = &v
	return s
}

type AddUserAccountResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddUserAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUserAccountResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserAccountResponseBody) SetSuccess(v bool) *AddUserAccountResponseBody {
	s.Success = &v
	return s
}

type AddUserAccountResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddUserAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddUserAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUserAccountResponse) GoString() string {
	return s.String()
}

func (s *AddUserAccountResponse) SetHeaders(v map[string]*string) *AddUserAccountResponse {
	s.Headers = v
	return s
}

func (s *AddUserAccountResponse) SetBody(v *AddUserAccountResponseBody) *AddUserAccountResponse {
	s.Body = v
	return s
}

type CollectRecruitJobDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CollectRecruitJobDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s CollectRecruitJobDetailHeaders) GoString() string {
	return s.String()
}

func (s *CollectRecruitJobDetailHeaders) SetCommonHeaders(v map[string]*string) *CollectRecruitJobDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CollectRecruitJobDetailHeaders) SetXAcsDingtalkAccessToken(v string) *CollectRecruitJobDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CollectRecruitJobDetailRequest struct {
	// 业务标识，目前固定为ddats
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	// 业务标识，目前固定为ddats
	Channel *string                                `json:"channel,omitempty" xml:"channel,omitempty"`
	JonInfo *CollectRecruitJobDetailRequestJonInfo `json:"jonInfo,omitempty" xml:"jonInfo,omitempty" type:"Struct"`
	// 渠道侧外部企业唯一ID
	OutCorpId *string `json:"outCorpId,omitempty" xml:"outCorpId,omitempty"`
	// 企业名称
	OutCorpName *string `json:"outCorpName,omitempty" xml:"outCorpName,omitempty"`
	// 招聘人信息
	RecruitUserInfo *CollectRecruitJobDetailRequestRecruitUserInfo `json:"recruitUserInfo,omitempty" xml:"recruitUserInfo,omitempty" type:"Struct"`
	// 来源
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 数据源更新时间
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s CollectRecruitJobDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s CollectRecruitJobDetailRequest) GoString() string {
	return s.String()
}

func (s *CollectRecruitJobDetailRequest) SetBizCode(v string) *CollectRecruitJobDetailRequest {
	s.BizCode = &v
	return s
}

func (s *CollectRecruitJobDetailRequest) SetChannel(v string) *CollectRecruitJobDetailRequest {
	s.Channel = &v
	return s
}

func (s *CollectRecruitJobDetailRequest) SetJonInfo(v *CollectRecruitJobDetailRequestJonInfo) *CollectRecruitJobDetailRequest {
	s.JonInfo = v
	return s
}

func (s *CollectRecruitJobDetailRequest) SetOutCorpId(v string) *CollectRecruitJobDetailRequest {
	s.OutCorpId = &v
	return s
}

func (s *CollectRecruitJobDetailRequest) SetOutCorpName(v string) *CollectRecruitJobDetailRequest {
	s.OutCorpName = &v
	return s
}

func (s *CollectRecruitJobDetailRequest) SetRecruitUserInfo(v *CollectRecruitJobDetailRequestRecruitUserInfo) *CollectRecruitJobDetailRequest {
	s.RecruitUserInfo = v
	return s
}

func (s *CollectRecruitJobDetailRequest) SetSource(v string) *CollectRecruitJobDetailRequest {
	s.Source = &v
	return s
}

func (s *CollectRecruitJobDetailRequest) SetUpdateTime(v int64) *CollectRecruitJobDetailRequest {
	s.UpdateTime = &v
	return s
}

type CollectRecruitJobDetailRequestJonInfo struct {
	// 地址信息
	Address *CollectRecruitJobDetailRequestJonInfoAddress `json:"address,omitempty" xml:"address,omitempty" type:"Struct"`
	// 职位分类编码
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 职位描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	ExtInfo     *string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// 全职信息
	FullTimeInfo *CollectRecruitJobDetailRequestJonInfoFullTimeInfo `json:"fullTimeInfo,omitempty" xml:"fullTimeInfo,omitempty" type:"Struct"`
	// 招聘人数
	HeadCount *string `json:"headCount,omitempty" xml:"headCount,omitempty"`
	// 职位性质
	JobNature *string `json:"jobNature,omitempty" xml:"jobNature,omitempty"`
	// 职位标签，字符串列表
	JobTags []*string `json:"jobTags,omitempty" xml:"jobTags,omitempty" type:"Repeated"`
	// 最高薪资
	MaxSalary *string `json:"maxSalary,omitempty" xml:"maxSalary,omitempty"`
	// 最低薪资
	MinSalary *string `json:"minSalary,omitempty" xml:"minSalary,omitempty"`
	// 职位名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 渠道职位ID
	OutJobId *string `json:"outJobId,omitempty" xml:"outJobId,omitempty"`
	// 兼职信息
	PartTimeInfo *CollectRecruitJobDetailRequestJonInfoPartTimeInfo `json:"partTimeInfo,omitempty" xml:"partTimeInfo,omitempty" type:"Struct"`
	// 学历要求
	RequiredEdu *string `json:"requiredEdu,omitempty" xml:"requiredEdu,omitempty"`
}

func (s CollectRecruitJobDetailRequestJonInfo) String() string {
	return tea.Prettify(s)
}

func (s CollectRecruitJobDetailRequestJonInfo) GoString() string {
	return s.String()
}

func (s *CollectRecruitJobDetailRequestJonInfo) SetAddress(v *CollectRecruitJobDetailRequestJonInfoAddress) *CollectRecruitJobDetailRequestJonInfo {
	s.Address = v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfo) SetCategory(v string) *CollectRecruitJobDetailRequestJonInfo {
	s.Category = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfo) SetDescription(v string) *CollectRecruitJobDetailRequestJonInfo {
	s.Description = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfo) SetExtInfo(v string) *CollectRecruitJobDetailRequestJonInfo {
	s.ExtInfo = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfo) SetFullTimeInfo(v *CollectRecruitJobDetailRequestJonInfoFullTimeInfo) *CollectRecruitJobDetailRequestJonInfo {
	s.FullTimeInfo = v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfo) SetHeadCount(v string) *CollectRecruitJobDetailRequestJonInfo {
	s.HeadCount = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfo) SetJobNature(v string) *CollectRecruitJobDetailRequestJonInfo {
	s.JobNature = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfo) SetJobTags(v []*string) *CollectRecruitJobDetailRequestJonInfo {
	s.JobTags = v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfo) SetMaxSalary(v string) *CollectRecruitJobDetailRequestJonInfo {
	s.MaxSalary = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfo) SetMinSalary(v string) *CollectRecruitJobDetailRequestJonInfo {
	s.MinSalary = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfo) SetName(v string) *CollectRecruitJobDetailRequestJonInfo {
	s.Name = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfo) SetOutJobId(v string) *CollectRecruitJobDetailRequestJonInfo {
	s.OutJobId = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfo) SetPartTimeInfo(v *CollectRecruitJobDetailRequestJonInfoPartTimeInfo) *CollectRecruitJobDetailRequestJonInfo {
	s.PartTimeInfo = v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfo) SetRequiredEdu(v string) *CollectRecruitJobDetailRequestJonInfo {
	s.RequiredEdu = &v
	return s
}

type CollectRecruitJobDetailRequestJonInfoAddress struct {
	// 城市编码
	CityCode *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	// 位置详情描述
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	// 区县编码
	DistrictCode *string `json:"districtCode,omitempty" xml:"districtCode,omitempty"`
	// 经度（高德地图选点）
	Latitude *string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 纬度（高德地图选点）
	Longitude *string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 位置名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 省份编码
	ProvinceCode *string `json:"provinceCode,omitempty" xml:"provinceCode,omitempty"`
}

func (s CollectRecruitJobDetailRequestJonInfoAddress) String() string {
	return tea.Prettify(s)
}

func (s CollectRecruitJobDetailRequestJonInfoAddress) GoString() string {
	return s.String()
}

func (s *CollectRecruitJobDetailRequestJonInfoAddress) SetCityCode(v string) *CollectRecruitJobDetailRequestJonInfoAddress {
	s.CityCode = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfoAddress) SetDetail(v string) *CollectRecruitJobDetailRequestJonInfoAddress {
	s.Detail = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfoAddress) SetDistrictCode(v string) *CollectRecruitJobDetailRequestJonInfoAddress {
	s.DistrictCode = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfoAddress) SetLatitude(v string) *CollectRecruitJobDetailRequestJonInfoAddress {
	s.Latitude = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfoAddress) SetLongitude(v string) *CollectRecruitJobDetailRequestJonInfoAddress {
	s.Longitude = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfoAddress) SetName(v string) *CollectRecruitJobDetailRequestJonInfoAddress {
	s.Name = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfoAddress) SetProvinceCode(v string) *CollectRecruitJobDetailRequestJonInfoAddress {
	s.ProvinceCode = &v
	return s
}

type CollectRecruitJobDetailRequestJonInfoFullTimeInfo struct {
	// 工作经验要求最高年限
	MaxJobExperience *string `json:"maxJobExperience,omitempty" xml:"maxJobExperience,omitempty"`
	// 工作经验要求最低年限
	MinJobExperience *string `json:"minJobExperience,omitempty" xml:"minJobExperience,omitempty"`
	// 薪资发放月数
	SalaryMonth *string `json:"salaryMonth,omitempty" xml:"salaryMonth,omitempty"`
}

func (s CollectRecruitJobDetailRequestJonInfoFullTimeInfo) String() string {
	return tea.Prettify(s)
}

func (s CollectRecruitJobDetailRequestJonInfoFullTimeInfo) GoString() string {
	return s.String()
}

func (s *CollectRecruitJobDetailRequestJonInfoFullTimeInfo) SetMaxJobExperience(v string) *CollectRecruitJobDetailRequestJonInfoFullTimeInfo {
	s.MaxJobExperience = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfoFullTimeInfo) SetMinJobExperience(v string) *CollectRecruitJobDetailRequestJonInfoFullTimeInfo {
	s.MinJobExperience = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfoFullTimeInfo) SetSalaryMonth(v string) *CollectRecruitJobDetailRequestJonInfoFullTimeInfo {
	s.SalaryMonth = &v
	return s
}

type CollectRecruitJobDetailRequestJonInfoPartTimeInfo struct {
	// 联系电话
	ContactNumber *string `json:"contactNumber,omitempty" xml:"contactNumber,omitempty"`
	// 薪资发放周期
	SalaryPeriod *string `json:"salaryPeriod,omitempty" xml:"salaryPeriod,omitempty"`
	// 薪资结算类型
	SettleType *string `json:"settleType,omitempty" xml:"settleType,omitempty"`
	// 是否指定工作日期
	SpecifyWorkDate *string `json:"specifyWorkDate,omitempty" xml:"specifyWorkDate,omitempty"`
	// 是否指定工作时间
	SpecifyWorkTime *string `json:"specifyWorkTime,omitempty" xml:"specifyWorkTime,omitempty"`
	// 工作开始时间
	WorkBeginTimeMin *string `json:"workBeginTimeMin,omitempty" xml:"workBeginTimeMin,omitempty"`
	// 工作日期类型
	WorkDateType *string `json:"workDateType,omitempty" xml:"workDateType,omitempty"`
	// 工作结束日期
	WorkEndDate *string `json:"workEndDate,omitempty" xml:"workEndDate,omitempty"`
	// 工作结束时间
	WorkEndTimeMin *string `json:"workEndTimeMin,omitempty" xml:"workEndTimeMin,omitempty"`
	// 工作开始日期
	WorkStartDate *string `json:"workStartDate,omitempty" xml:"workStartDate,omitempty"`
}

func (s CollectRecruitJobDetailRequestJonInfoPartTimeInfo) String() string {
	return tea.Prettify(s)
}

func (s CollectRecruitJobDetailRequestJonInfoPartTimeInfo) GoString() string {
	return s.String()
}

func (s *CollectRecruitJobDetailRequestJonInfoPartTimeInfo) SetContactNumber(v string) *CollectRecruitJobDetailRequestJonInfoPartTimeInfo {
	s.ContactNumber = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfoPartTimeInfo) SetSalaryPeriod(v string) *CollectRecruitJobDetailRequestJonInfoPartTimeInfo {
	s.SalaryPeriod = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfoPartTimeInfo) SetSettleType(v string) *CollectRecruitJobDetailRequestJonInfoPartTimeInfo {
	s.SettleType = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfoPartTimeInfo) SetSpecifyWorkDate(v string) *CollectRecruitJobDetailRequestJonInfoPartTimeInfo {
	s.SpecifyWorkDate = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfoPartTimeInfo) SetSpecifyWorkTime(v string) *CollectRecruitJobDetailRequestJonInfoPartTimeInfo {
	s.SpecifyWorkTime = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfoPartTimeInfo) SetWorkBeginTimeMin(v string) *CollectRecruitJobDetailRequestJonInfoPartTimeInfo {
	s.WorkBeginTimeMin = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfoPartTimeInfo) SetWorkDateType(v string) *CollectRecruitJobDetailRequestJonInfoPartTimeInfo {
	s.WorkDateType = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfoPartTimeInfo) SetWorkEndDate(v string) *CollectRecruitJobDetailRequestJonInfoPartTimeInfo {
	s.WorkEndDate = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfoPartTimeInfo) SetWorkEndTimeMin(v string) *CollectRecruitJobDetailRequestJonInfoPartTimeInfo {
	s.WorkEndTimeMin = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJonInfoPartTimeInfo) SetWorkStartDate(v string) *CollectRecruitJobDetailRequestJonInfoPartTimeInfo {
	s.WorkStartDate = &v
	return s
}

type CollectRecruitJobDetailRequestRecruitUserInfo struct {
	// 额外信息
	ExtInfo *string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// 招聘员工唯一ID
	OutUserId *string `json:"outUserId,omitempty" xml:"outUserId,omitempty"`
	// 招聘员工手机号码
	UserMobile *string `json:"userMobile,omitempty" xml:"userMobile,omitempty"`
	// 招聘员工姓名
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s CollectRecruitJobDetailRequestRecruitUserInfo) String() string {
	return tea.Prettify(s)
}

func (s CollectRecruitJobDetailRequestRecruitUserInfo) GoString() string {
	return s.String()
}

func (s *CollectRecruitJobDetailRequestRecruitUserInfo) SetExtInfo(v string) *CollectRecruitJobDetailRequestRecruitUserInfo {
	s.ExtInfo = &v
	return s
}

func (s *CollectRecruitJobDetailRequestRecruitUserInfo) SetOutUserId(v string) *CollectRecruitJobDetailRequestRecruitUserInfo {
	s.OutUserId = &v
	return s
}

func (s *CollectRecruitJobDetailRequestRecruitUserInfo) SetUserMobile(v string) *CollectRecruitJobDetailRequestRecruitUserInfo {
	s.UserMobile = &v
	return s
}

func (s *CollectRecruitJobDetailRequestRecruitUserInfo) SetUserName(v string) *CollectRecruitJobDetailRequestRecruitUserInfo {
	s.UserName = &v
	return s
}

type CollectRecruitJobDetailResponseBody struct {
	// 是否成功
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CollectRecruitJobDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollectRecruitJobDetailResponseBody) GoString() string {
	return s.String()
}

func (s *CollectRecruitJobDetailResponseBody) SetResult(v bool) *CollectRecruitJobDetailResponseBody {
	s.Result = &v
	return s
}

type CollectRecruitJobDetailResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CollectRecruitJobDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CollectRecruitJobDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s CollectRecruitJobDetailResponse) GoString() string {
	return s.String()
}

func (s *CollectRecruitJobDetailResponse) SetHeaders(v map[string]*string) *CollectRecruitJobDetailResponse {
	s.Headers = v
	return s
}

func (s *CollectRecruitJobDetailResponse) SetBody(v *CollectRecruitJobDetailResponseBody) *CollectRecruitJobDetailResponse {
	s.Body = v
	return s
}

type CollectResumeDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CollectResumeDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s CollectResumeDetailHeaders) GoString() string {
	return s.String()
}

func (s *CollectResumeDetailHeaders) SetCommonHeaders(v map[string]*string) *CollectResumeDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CollectResumeDetailHeaders) SetXAcsDingtalkAccessToken(v string) *CollectResumeDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CollectResumeDetailRequest struct {
	// 业务标识，目前固定为ddats
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	// 渠道侧简历标识
	ChannelOuterId *string `json:"channelOuterId,omitempty" xml:"channelOuterId,omitempty"`
	// 渠道侧候选人标识。
	ChannelTalentId *string `json:"channelTalentId,omitempty" xml:"channelTalentId,omitempty"`
	// 简历投递职位标识
	DeliverJobId *string `json:"deliverJobId,omitempty" xml:"deliverJobId,omitempty"`
	OptUserId    *string `json:"optUserId,omitempty" xml:"optUserId,omitempty"`
	// 简历详情信息
	ResumeData *CollectResumeDetailRequestResumeData `json:"resumeData,omitempty" xml:"resumeData,omitempty" type:"Struct"`
}

func (s CollectResumeDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s CollectResumeDetailRequest) GoString() string {
	return s.String()
}

func (s *CollectResumeDetailRequest) SetBizCode(v string) *CollectResumeDetailRequest {
	s.BizCode = &v
	return s
}

func (s *CollectResumeDetailRequest) SetChannelOuterId(v string) *CollectResumeDetailRequest {
	s.ChannelOuterId = &v
	return s
}

func (s *CollectResumeDetailRequest) SetChannelTalentId(v string) *CollectResumeDetailRequest {
	s.ChannelTalentId = &v
	return s
}

func (s *CollectResumeDetailRequest) SetDeliverJobId(v string) *CollectResumeDetailRequest {
	s.DeliverJobId = &v
	return s
}

func (s *CollectResumeDetailRequest) SetOptUserId(v string) *CollectResumeDetailRequest {
	s.OptUserId = &v
	return s
}

func (s *CollectResumeDetailRequest) SetResumeData(v *CollectResumeDetailRequestResumeData) *CollectResumeDetailRequest {
	s.ResumeData = v
	return s
}

type CollectResumeDetailRequestResumeData struct {
	// 简历基础信息
	BaseInfo *CollectResumeDetailRequestResumeDataBaseInfo `json:"baseInfo,omitempty" xml:"baseInfo,omitempty" type:"Struct"`
	// 证书信息
	Certificates []*CollectResumeDetailRequestResumeDataCertificates `json:"certificates,omitempty" xml:"certificates,omitempty" type:"Repeated"`
	// 教育经历
	EducationExperiences []*CollectResumeDetailRequestResumeDataEducationExperiences `json:"educationExperiences,omitempty" xml:"educationExperiences,omitempty" type:"Repeated"`
	// 期望职位信息
	JobExpect *CollectResumeDetailRequestResumeDataJobExpect `json:"jobExpect,omitempty" xml:"jobExpect,omitempty" type:"Struct"`
	// 语言能力
	LanguageSkill []*CollectResumeDetailRequestResumeDataLanguageSkill `json:"languageSkill,omitempty" xml:"languageSkill,omitempty" type:"Repeated"`
	// 培训经历
	TrainingExperiences []*CollectResumeDetailRequestResumeDataTrainingExperiences `json:"trainingExperiences,omitempty" xml:"trainingExperiences,omitempty" type:"Repeated"`
	// 工作经历
	WorkExperiences []*CollectResumeDetailRequestResumeDataWorkExperiences `json:"workExperiences,omitempty" xml:"workExperiences,omitempty" type:"Repeated"`
}

func (s CollectResumeDetailRequestResumeData) String() string {
	return tea.Prettify(s)
}

func (s CollectResumeDetailRequestResumeData) GoString() string {
	return s.String()
}

func (s *CollectResumeDetailRequestResumeData) SetBaseInfo(v *CollectResumeDetailRequestResumeDataBaseInfo) *CollectResumeDetailRequestResumeData {
	s.BaseInfo = v
	return s
}

func (s *CollectResumeDetailRequestResumeData) SetCertificates(v []*CollectResumeDetailRequestResumeDataCertificates) *CollectResumeDetailRequestResumeData {
	s.Certificates = v
	return s
}

func (s *CollectResumeDetailRequestResumeData) SetEducationExperiences(v []*CollectResumeDetailRequestResumeDataEducationExperiences) *CollectResumeDetailRequestResumeData {
	s.EducationExperiences = v
	return s
}

func (s *CollectResumeDetailRequestResumeData) SetJobExpect(v *CollectResumeDetailRequestResumeDataJobExpect) *CollectResumeDetailRequestResumeData {
	s.JobExpect = v
	return s
}

func (s *CollectResumeDetailRequestResumeData) SetLanguageSkill(v []*CollectResumeDetailRequestResumeDataLanguageSkill) *CollectResumeDetailRequestResumeData {
	s.LanguageSkill = v
	return s
}

func (s *CollectResumeDetailRequestResumeData) SetTrainingExperiences(v []*CollectResumeDetailRequestResumeDataTrainingExperiences) *CollectResumeDetailRequestResumeData {
	s.TrainingExperiences = v
	return s
}

func (s *CollectResumeDetailRequestResumeData) SetWorkExperiences(v []*CollectResumeDetailRequestResumeDataWorkExperiences) *CollectResumeDetailRequestResumeData {
	s.WorkExperiences = v
	return s
}

type CollectResumeDetailRequestResumeDataBaseInfo struct {
	// 年龄
	Age *int32 `json:"age,omitempty" xml:"age,omitempty"`
	// 头像cdn地址，http链接
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// 初次工作时间
	BeginWorkTime *string `json:"beginWorkTime,omitempty" xml:"beginWorkTime,omitempty"`
	// 生日
	Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// 邮箱地址
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// 英文名称
	EnglishName *string `json:"englishName,omitempty" xml:"englishName,omitempty"`
	// 毕业时间
	GraduateTime *string `json:"graduateTime,omitempty" xml:"graduateTime,omitempty"`
	// 最高学历
	HighestEducation *int32 `json:"highestEducation,omitempty" xml:"highestEducation,omitempty"`
	// 当前工作职位名称
	JobTitle *string `json:"jobTitle,omitempty" xml:"jobTitle,omitempty"`
	// 最高学历毕业院校名称
	LastSchoolName *string `json:"lastSchoolName,omitempty" xml:"lastSchoolName,omitempty"`
	// 婚姻状况
	Married *int32 `json:"married,omitempty" xml:"married,omitempty"`
	// 名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 籍贯地址
	NativePlace *string `json:"nativePlace,omitempty" xml:"nativePlace,omitempty"`
	// 现居住地址
	NowLocation *string `json:"nowLocation,omitempty" xml:"nowLocation,omitempty"`
	// 个人荣誉
	PersonalHonor *string `json:"personalHonor,omitempty" xml:"personalHonor,omitempty"`
	// 手机号
	PhoneNum *string `json:"phoneNum,omitempty" xml:"phoneNum,omitempty"`
	// 政治面貌
	PoliticalStatus *int32 `json:"politicalStatus,omitempty" xml:"politicalStatus,omitempty"`
	// 自我评价
	SelfEvaluation *string `json:"selfEvaluation,omitempty" xml:"selfEvaluation,omitempty"`
	// 性别
	Sex *int32 `json:"sex,omitempty" xml:"sex,omitempty"`
	// 虚拟手机号
	VirtualPhoneNum *string `json:"virtualPhoneNum,omitempty" xml:"virtualPhoneNum,omitempty"`
	// 工作年限
	WorkingYears *int32 `json:"workingYears,omitempty" xml:"workingYears,omitempty"`
}

func (s CollectResumeDetailRequestResumeDataBaseInfo) String() string {
	return tea.Prettify(s)
}

func (s CollectResumeDetailRequestResumeDataBaseInfo) GoString() string {
	return s.String()
}

func (s *CollectResumeDetailRequestResumeDataBaseInfo) SetAge(v int32) *CollectResumeDetailRequestResumeDataBaseInfo {
	s.Age = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataBaseInfo) SetAvatar(v string) *CollectResumeDetailRequestResumeDataBaseInfo {
	s.Avatar = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataBaseInfo) SetBeginWorkTime(v string) *CollectResumeDetailRequestResumeDataBaseInfo {
	s.BeginWorkTime = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataBaseInfo) SetBirthday(v string) *CollectResumeDetailRequestResumeDataBaseInfo {
	s.Birthday = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataBaseInfo) SetEmail(v string) *CollectResumeDetailRequestResumeDataBaseInfo {
	s.Email = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataBaseInfo) SetEnglishName(v string) *CollectResumeDetailRequestResumeDataBaseInfo {
	s.EnglishName = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataBaseInfo) SetGraduateTime(v string) *CollectResumeDetailRequestResumeDataBaseInfo {
	s.GraduateTime = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataBaseInfo) SetHighestEducation(v int32) *CollectResumeDetailRequestResumeDataBaseInfo {
	s.HighestEducation = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataBaseInfo) SetJobTitle(v string) *CollectResumeDetailRequestResumeDataBaseInfo {
	s.JobTitle = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataBaseInfo) SetLastSchoolName(v string) *CollectResumeDetailRequestResumeDataBaseInfo {
	s.LastSchoolName = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataBaseInfo) SetMarried(v int32) *CollectResumeDetailRequestResumeDataBaseInfo {
	s.Married = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataBaseInfo) SetName(v string) *CollectResumeDetailRequestResumeDataBaseInfo {
	s.Name = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataBaseInfo) SetNativePlace(v string) *CollectResumeDetailRequestResumeDataBaseInfo {
	s.NativePlace = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataBaseInfo) SetNowLocation(v string) *CollectResumeDetailRequestResumeDataBaseInfo {
	s.NowLocation = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataBaseInfo) SetPersonalHonor(v string) *CollectResumeDetailRequestResumeDataBaseInfo {
	s.PersonalHonor = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataBaseInfo) SetPhoneNum(v string) *CollectResumeDetailRequestResumeDataBaseInfo {
	s.PhoneNum = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataBaseInfo) SetPoliticalStatus(v int32) *CollectResumeDetailRequestResumeDataBaseInfo {
	s.PoliticalStatus = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataBaseInfo) SetSelfEvaluation(v string) *CollectResumeDetailRequestResumeDataBaseInfo {
	s.SelfEvaluation = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataBaseInfo) SetSex(v int32) *CollectResumeDetailRequestResumeDataBaseInfo {
	s.Sex = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataBaseInfo) SetVirtualPhoneNum(v string) *CollectResumeDetailRequestResumeDataBaseInfo {
	s.VirtualPhoneNum = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataBaseInfo) SetWorkingYears(v int32) *CollectResumeDetailRequestResumeDataBaseInfo {
	s.WorkingYears = &v
	return s
}

type CollectResumeDetailRequestResumeDataCertificates struct {
	// 证书名称
	CertificateName *string `json:"certificateName,omitempty" xml:"certificateName,omitempty"`
	// 证书授予时间
	GrantTime *string `json:"grantTime,omitempty" xml:"grantTime,omitempty"`
}

func (s CollectResumeDetailRequestResumeDataCertificates) String() string {
	return tea.Prettify(s)
}

func (s CollectResumeDetailRequestResumeDataCertificates) GoString() string {
	return s.String()
}

func (s *CollectResumeDetailRequestResumeDataCertificates) SetCertificateName(v string) *CollectResumeDetailRequestResumeDataCertificates {
	s.CertificateName = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataCertificates) SetGrantTime(v string) *CollectResumeDetailRequestResumeDataCertificates {
	s.GrantTime = &v
	return s
}

type CollectResumeDetailRequestResumeDataEducationExperiences struct {
	// 学历
	Degree *int32 `json:"degree,omitempty" xml:"degree,omitempty"`
	// 院系
	Department *string `json:"department,omitempty" xml:"department,omitempty"`
	// 详细描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 结束时间
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// 专业
	Major *string `json:"major,omitempty" xml:"major,omitempty"`
	// 学校名称
	SchoolName *string `json:"schoolName,omitempty" xml:"schoolName,omitempty"`
	// 开始时间
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
}

func (s CollectResumeDetailRequestResumeDataEducationExperiences) String() string {
	return tea.Prettify(s)
}

func (s CollectResumeDetailRequestResumeDataEducationExperiences) GoString() string {
	return s.String()
}

func (s *CollectResumeDetailRequestResumeDataEducationExperiences) SetDegree(v int32) *CollectResumeDetailRequestResumeDataEducationExperiences {
	s.Degree = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataEducationExperiences) SetDepartment(v string) *CollectResumeDetailRequestResumeDataEducationExperiences {
	s.Department = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataEducationExperiences) SetDescription(v string) *CollectResumeDetailRequestResumeDataEducationExperiences {
	s.Description = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataEducationExperiences) SetEndDate(v string) *CollectResumeDetailRequestResumeDataEducationExperiences {
	s.EndDate = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataEducationExperiences) SetMajor(v string) *CollectResumeDetailRequestResumeDataEducationExperiences {
	s.Major = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataEducationExperiences) SetSchoolName(v string) *CollectResumeDetailRequestResumeDataEducationExperiences {
	s.SchoolName = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataEducationExperiences) SetStartDate(v string) *CollectResumeDetailRequestResumeDataEducationExperiences {
	s.StartDate = &v
	return s
}

type CollectResumeDetailRequestResumeDataJobExpect struct {
	// 期望职位名称
	JobName *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
	// 期望工作地
	Locations []*string `json:"locations,omitempty" xml:"locations,omitempty" type:"Repeated"`
	// 最高期望工资
	MaxSalary *string `json:"maxSalary,omitempty" xml:"maxSalary,omitempty"`
	// 最低期望工资
	MinSalary *string `json:"minSalary,omitempty" xml:"minSalary,omitempty"`
	// 期望入职时间
	OnboardTime *string `json:"onboardTime,omitempty" xml:"onboardTime,omitempty"`
}

func (s CollectResumeDetailRequestResumeDataJobExpect) String() string {
	return tea.Prettify(s)
}

func (s CollectResumeDetailRequestResumeDataJobExpect) GoString() string {
	return s.String()
}

func (s *CollectResumeDetailRequestResumeDataJobExpect) SetJobName(v string) *CollectResumeDetailRequestResumeDataJobExpect {
	s.JobName = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataJobExpect) SetLocations(v []*string) *CollectResumeDetailRequestResumeDataJobExpect {
	s.Locations = v
	return s
}

func (s *CollectResumeDetailRequestResumeDataJobExpect) SetMaxSalary(v string) *CollectResumeDetailRequestResumeDataJobExpect {
	s.MaxSalary = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataJobExpect) SetMinSalary(v string) *CollectResumeDetailRequestResumeDataJobExpect {
	s.MinSalary = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataJobExpect) SetOnboardTime(v string) *CollectResumeDetailRequestResumeDataJobExpect {
	s.OnboardTime = &v
	return s
}

type CollectResumeDetailRequestResumeDataLanguageSkill struct {
	// 证书名称
	CertificateName *string `json:"certificateName,omitempty" xml:"certificateName,omitempty"`
	// 语言名称
	LanguageName *string `json:"languageName,omitempty" xml:"languageName,omitempty"`
}

func (s CollectResumeDetailRequestResumeDataLanguageSkill) String() string {
	return tea.Prettify(s)
}

func (s CollectResumeDetailRequestResumeDataLanguageSkill) GoString() string {
	return s.String()
}

func (s *CollectResumeDetailRequestResumeDataLanguageSkill) SetCertificateName(v string) *CollectResumeDetailRequestResumeDataLanguageSkill {
	s.CertificateName = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataLanguageSkill) SetLanguageName(v string) *CollectResumeDetailRequestResumeDataLanguageSkill {
	s.LanguageName = &v
	return s
}

type CollectResumeDetailRequestResumeDataTrainingExperiences struct {
	// 详细内容描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 结束时间
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// 培训机构名称
	InstitutionName *string `json:"institutionName,omitempty" xml:"institutionName,omitempty"`
	// 培训地点
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// 培训名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 开始时间
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
}

func (s CollectResumeDetailRequestResumeDataTrainingExperiences) String() string {
	return tea.Prettify(s)
}

func (s CollectResumeDetailRequestResumeDataTrainingExperiences) GoString() string {
	return s.String()
}

func (s *CollectResumeDetailRequestResumeDataTrainingExperiences) SetDescription(v string) *CollectResumeDetailRequestResumeDataTrainingExperiences {
	s.Description = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataTrainingExperiences) SetEndDate(v string) *CollectResumeDetailRequestResumeDataTrainingExperiences {
	s.EndDate = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataTrainingExperiences) SetInstitutionName(v string) *CollectResumeDetailRequestResumeDataTrainingExperiences {
	s.InstitutionName = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataTrainingExperiences) SetLocation(v string) *CollectResumeDetailRequestResumeDataTrainingExperiences {
	s.Location = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataTrainingExperiences) SetName(v string) *CollectResumeDetailRequestResumeDataTrainingExperiences {
	s.Name = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataTrainingExperiences) SetStartDate(v string) *CollectResumeDetailRequestResumeDataTrainingExperiences {
	s.StartDate = &v
	return s
}

type CollectResumeDetailRequestResumeDataWorkExperiences struct {
	// 公司名称
	CompanyName *string `json:"companyName,omitempty" xml:"companyName,omitempty"`
	// 部门
	Department *string `json:"department,omitempty" xml:"department,omitempty"`
	// 工作详情描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	EndDate     *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// 职位名称
	JobTitle *string `json:"jobTitle,omitempty" xml:"jobTitle,omitempty"`
	// 工作地点
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// 工作职责
	Responsibility *string `json:"responsibility,omitempty" xml:"responsibility,omitempty"`
	StartDate      *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
}

func (s CollectResumeDetailRequestResumeDataWorkExperiences) String() string {
	return tea.Prettify(s)
}

func (s CollectResumeDetailRequestResumeDataWorkExperiences) GoString() string {
	return s.String()
}

func (s *CollectResumeDetailRequestResumeDataWorkExperiences) SetCompanyName(v string) *CollectResumeDetailRequestResumeDataWorkExperiences {
	s.CompanyName = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataWorkExperiences) SetDepartment(v string) *CollectResumeDetailRequestResumeDataWorkExperiences {
	s.Department = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataWorkExperiences) SetDescription(v string) *CollectResumeDetailRequestResumeDataWorkExperiences {
	s.Description = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataWorkExperiences) SetEndDate(v string) *CollectResumeDetailRequestResumeDataWorkExperiences {
	s.EndDate = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataWorkExperiences) SetJobTitle(v string) *CollectResumeDetailRequestResumeDataWorkExperiences {
	s.JobTitle = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataWorkExperiences) SetLocation(v string) *CollectResumeDetailRequestResumeDataWorkExperiences {
	s.Location = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataWorkExperiences) SetResponsibility(v string) *CollectResumeDetailRequestResumeDataWorkExperiences {
	s.Responsibility = &v
	return s
}

func (s *CollectResumeDetailRequestResumeDataWorkExperiences) SetStartDate(v string) *CollectResumeDetailRequestResumeDataWorkExperiences {
	s.StartDate = &v
	return s
}

type CollectResumeDetailResponseBody struct {
	// 简历标识
	ResumeId *string `json:"resumeId,omitempty" xml:"resumeId,omitempty"`
}

func (s CollectResumeDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollectResumeDetailResponseBody) GoString() string {
	return s.String()
}

func (s *CollectResumeDetailResponseBody) SetResumeId(v string) *CollectResumeDetailResponseBody {
	s.ResumeId = &v
	return s
}

type CollectResumeDetailResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CollectResumeDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CollectResumeDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s CollectResumeDetailResponse) GoString() string {
	return s.String()
}

func (s *CollectResumeDetailResponse) SetHeaders(v map[string]*string) *CollectResumeDetailResponse {
	s.Headers = v
	return s
}

func (s *CollectResumeDetailResponse) SetBody(v *CollectResumeDetailResponseBody) *CollectResumeDetailResponse {
	s.Body = v
	return s
}

type ConfirmRightsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ConfirmRightsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ConfirmRightsHeaders) GoString() string {
	return s.String()
}

func (s *ConfirmRightsHeaders) SetCommonHeaders(v map[string]*string) *ConfirmRightsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ConfirmRightsHeaders) SetXAcsDingtalkAccessToken(v string) *ConfirmRightsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ConfirmRightsRequest struct {
	// 业务标识
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
}

func (s ConfirmRightsRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfirmRightsRequest) GoString() string {
	return s.String()
}

func (s *ConfirmRightsRequest) SetBizCode(v string) *ConfirmRightsRequest {
	s.BizCode = &v
	return s
}

type ConfirmRightsResponseBody struct {
	// 结果
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ConfirmRightsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfirmRightsResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmRightsResponseBody) SetResult(v bool) *ConfirmRightsResponseBody {
	s.Result = &v
	return s
}

type ConfirmRightsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfirmRightsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfirmRightsResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfirmRightsResponse) GoString() string {
	return s.String()
}

func (s *ConfirmRightsResponse) SetHeaders(v map[string]*string) *ConfirmRightsResponse {
	s.Headers = v
	return s
}

func (s *ConfirmRightsResponse) SetBody(v *ConfirmRightsResponseBody) *ConfirmRightsResponse {
	s.Body = v
	return s
}

type FinishBeginnerTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s FinishBeginnerTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s FinishBeginnerTaskHeaders) GoString() string {
	return s.String()
}

func (s *FinishBeginnerTaskHeaders) SetCommonHeaders(v map[string]*string) *FinishBeginnerTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FinishBeginnerTaskHeaders) SetXAcsDingtalkAccessToken(v string) *FinishBeginnerTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type FinishBeginnerTaskRequest struct {
	// 任务范围
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// 员工标识
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s FinishBeginnerTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s FinishBeginnerTaskRequest) GoString() string {
	return s.String()
}

func (s *FinishBeginnerTaskRequest) SetScope(v string) *FinishBeginnerTaskRequest {
	s.Scope = &v
	return s
}

func (s *FinishBeginnerTaskRequest) SetUserId(v string) *FinishBeginnerTaskRequest {
	s.UserId = &v
	return s
}

type FinishBeginnerTaskResponseBody struct {
	// 是否成功
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s FinishBeginnerTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FinishBeginnerTaskResponseBody) GoString() string {
	return s.String()
}

func (s *FinishBeginnerTaskResponseBody) SetResult(v bool) *FinishBeginnerTaskResponseBody {
	s.Result = &v
	return s
}

type FinishBeginnerTaskResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FinishBeginnerTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FinishBeginnerTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s FinishBeginnerTaskResponse) GoString() string {
	return s.String()
}

func (s *FinishBeginnerTaskResponse) SetHeaders(v map[string]*string) *FinishBeginnerTaskResponse {
	s.Headers = v
	return s
}

func (s *FinishBeginnerTaskResponse) SetBody(v *FinishBeginnerTaskResponseBody) *FinishBeginnerTaskResponse {
	s.Body = v
	return s
}

type GetApplicationRegFormByFlowIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetApplicationRegFormByFlowIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationRegFormByFlowIdHeaders) GoString() string {
	return s.String()
}

func (s *GetApplicationRegFormByFlowIdHeaders) SetCommonHeaders(v map[string]*string) *GetApplicationRegFormByFlowIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetApplicationRegFormByFlowIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetApplicationRegFormByFlowIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetApplicationRegFormByFlowIdRequest struct {
	// 业务标识
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
}

func (s GetApplicationRegFormByFlowIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationRegFormByFlowIdRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationRegFormByFlowIdRequest) SetBizCode(v string) *GetApplicationRegFormByFlowIdRequest {
	s.BizCode = &v
	return s
}

type GetApplicationRegFormByFlowIdResponseBody struct {
	// 候选人标识
	CandidateId *string `json:"candidateId,omitempty" xml:"candidateId,omitempty"`
	// 邀填人员工标识
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// 招聘流程标识
	FlowId *string `json:"flowId,omitempty" xml:"flowId,omitempty"`
	// 表单标识
	FormId *string `json:"formId,omitempty" xml:"formId,omitempty"`
	// 创建时间（邀填时间，单位：毫秒）
	GmtCreateMillis *int64 `json:"gmtCreateMillis,omitempty" xml:"gmtCreateMillis,omitempty"`
	// 更新时间（填写时间，单位：毫秒），仅当表单状态为已填写时有效
	GmtModifiedMillis *int64 `json:"gmtModifiedMillis,omitempty" xml:"gmtModifiedMillis,omitempty"`
	// 职位标识
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// 表单状态（0表示未填写，1表示已填写）
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// 模板标识
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// 模板版本
	TemplateVersion *int32 `json:"templateVersion,omitempty" xml:"templateVersion,omitempty"`
}

func (s GetApplicationRegFormByFlowIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationRegFormByFlowIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationRegFormByFlowIdResponseBody) SetCandidateId(v string) *GetApplicationRegFormByFlowIdResponseBody {
	s.CandidateId = &v
	return s
}

func (s *GetApplicationRegFormByFlowIdResponseBody) SetCreatorUserId(v string) *GetApplicationRegFormByFlowIdResponseBody {
	s.CreatorUserId = &v
	return s
}

func (s *GetApplicationRegFormByFlowIdResponseBody) SetFlowId(v string) *GetApplicationRegFormByFlowIdResponseBody {
	s.FlowId = &v
	return s
}

func (s *GetApplicationRegFormByFlowIdResponseBody) SetFormId(v string) *GetApplicationRegFormByFlowIdResponseBody {
	s.FormId = &v
	return s
}

func (s *GetApplicationRegFormByFlowIdResponseBody) SetGmtCreateMillis(v int64) *GetApplicationRegFormByFlowIdResponseBody {
	s.GmtCreateMillis = &v
	return s
}

func (s *GetApplicationRegFormByFlowIdResponseBody) SetGmtModifiedMillis(v int64) *GetApplicationRegFormByFlowIdResponseBody {
	s.GmtModifiedMillis = &v
	return s
}

func (s *GetApplicationRegFormByFlowIdResponseBody) SetJobId(v string) *GetApplicationRegFormByFlowIdResponseBody {
	s.JobId = &v
	return s
}

func (s *GetApplicationRegFormByFlowIdResponseBody) SetStatus(v int32) *GetApplicationRegFormByFlowIdResponseBody {
	s.Status = &v
	return s
}

func (s *GetApplicationRegFormByFlowIdResponseBody) SetTemplateId(v string) *GetApplicationRegFormByFlowIdResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetApplicationRegFormByFlowIdResponseBody) SetTemplateVersion(v int32) *GetApplicationRegFormByFlowIdResponseBody {
	s.TemplateVersion = &v
	return s
}

type GetApplicationRegFormByFlowIdResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetApplicationRegFormByFlowIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetApplicationRegFormByFlowIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationRegFormByFlowIdResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationRegFormByFlowIdResponse) SetHeaders(v map[string]*string) *GetApplicationRegFormByFlowIdResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationRegFormByFlowIdResponse) SetBody(v *GetApplicationRegFormByFlowIdResponseBody) *GetApplicationRegFormByFlowIdResponse {
	s.Body = v
	return s
}

type GetCandidateByPhoneNumberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCandidateByPhoneNumberHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCandidateByPhoneNumberHeaders) GoString() string {
	return s.String()
}

func (s *GetCandidateByPhoneNumberHeaders) SetCommonHeaders(v map[string]*string) *GetCandidateByPhoneNumberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCandidateByPhoneNumberHeaders) SetXAcsDingtalkAccessToken(v string) *GetCandidateByPhoneNumberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCandidateByPhoneNumberRequest struct {
	// 业务标识
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	// 候选人手机号
	PhoneNumber *string `json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
}

func (s GetCandidateByPhoneNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCandidateByPhoneNumberRequest) GoString() string {
	return s.String()
}

func (s *GetCandidateByPhoneNumberRequest) SetBizCode(v string) *GetCandidateByPhoneNumberRequest {
	s.BizCode = &v
	return s
}

func (s *GetCandidateByPhoneNumberRequest) SetPhoneNumber(v string) *GetCandidateByPhoneNumberRequest {
	s.PhoneNumber = &v
	return s
}

type GetCandidateByPhoneNumberResponseBody struct {
	// 候选人标识
	CandidateId *string `json:"candidateId,omitempty" xml:"candidateId,omitempty"`
	// 候选人姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetCandidateByPhoneNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCandidateByPhoneNumberResponseBody) GoString() string {
	return s.String()
}

func (s *GetCandidateByPhoneNumberResponseBody) SetCandidateId(v string) *GetCandidateByPhoneNumberResponseBody {
	s.CandidateId = &v
	return s
}

func (s *GetCandidateByPhoneNumberResponseBody) SetName(v string) *GetCandidateByPhoneNumberResponseBody {
	s.Name = &v
	return s
}

type GetCandidateByPhoneNumberResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCandidateByPhoneNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCandidateByPhoneNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCandidateByPhoneNumberResponse) GoString() string {
	return s.String()
}

func (s *GetCandidateByPhoneNumberResponse) SetHeaders(v map[string]*string) *GetCandidateByPhoneNumberResponse {
	s.Headers = v
	return s
}

func (s *GetCandidateByPhoneNumberResponse) SetBody(v *GetCandidateByPhoneNumberResponseBody) *GetCandidateByPhoneNumberResponse {
	s.Body = v
	return s
}

type GetFileUploadInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFileUploadInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFileUploadInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoHeaders) SetCommonHeaders(v map[string]*string) *GetFileUploadInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFileUploadInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetFileUploadInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFileUploadInfoRequest struct {
	// 业务标识
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	// 文件名称
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件大小（单位：字节）
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// 文件MD5摘要
	Md5 *string `json:"md5,omitempty" xml:"md5,omitempty"`
	// 操作人员工标识，为空时默认以企业管理员身份进行操作
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s GetFileUploadInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFileUploadInfoRequest) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoRequest) SetBizCode(v string) *GetFileUploadInfoRequest {
	s.BizCode = &v
	return s
}

func (s *GetFileUploadInfoRequest) SetFileName(v string) *GetFileUploadInfoRequest {
	s.FileName = &v
	return s
}

func (s *GetFileUploadInfoRequest) SetFileSize(v int64) *GetFileUploadInfoRequest {
	s.FileSize = &v
	return s
}

func (s *GetFileUploadInfoRequest) SetMd5(v string) *GetFileUploadInfoRequest {
	s.Md5 = &v
	return s
}

func (s *GetFileUploadInfoRequest) SetOpUserId(v string) *GetFileUploadInfoRequest {
	s.OpUserId = &v
	return s
}

type GetFileUploadInfoResponseBody struct {
	// OSS上传所需信息：accessKeyId
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// OSS上传所需信息：accessKeySecret
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	// OSS上传所需信息：accessToken
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// accessToken有效期截止时间（单位：毫秒），需要在此时间之前完成文件上传
	AccessTokenExpirationMillis *int64 `json:"accessTokenExpirationMillis,omitempty" xml:"accessTokenExpirationMillis,omitempty"`
	// OSS上传所需信息：bucket
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// OSS上传所需信息：endPoint
	EndPoint *string `json:"endPoint,omitempty" xml:"endPoint,omitempty"`
	// 文件mediaId
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
}

func (s GetFileUploadInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFileUploadInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoResponseBody) SetAccessKeyId(v string) *GetFileUploadInfoResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *GetFileUploadInfoResponseBody) SetAccessKeySecret(v string) *GetFileUploadInfoResponseBody {
	s.AccessKeySecret = &v
	return s
}

func (s *GetFileUploadInfoResponseBody) SetAccessToken(v string) *GetFileUploadInfoResponseBody {
	s.AccessToken = &v
	return s
}

func (s *GetFileUploadInfoResponseBody) SetAccessTokenExpirationMillis(v int64) *GetFileUploadInfoResponseBody {
	s.AccessTokenExpirationMillis = &v
	return s
}

func (s *GetFileUploadInfoResponseBody) SetBucket(v string) *GetFileUploadInfoResponseBody {
	s.Bucket = &v
	return s
}

func (s *GetFileUploadInfoResponseBody) SetEndPoint(v string) *GetFileUploadInfoResponseBody {
	s.EndPoint = &v
	return s
}

func (s *GetFileUploadInfoResponseBody) SetMediaId(v string) *GetFileUploadInfoResponseBody {
	s.MediaId = &v
	return s
}

type GetFileUploadInfoResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFileUploadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFileUploadInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFileUploadInfoResponse) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoResponse) SetHeaders(v map[string]*string) *GetFileUploadInfoResponse {
	s.Headers = v
	return s
}

func (s *GetFileUploadInfoResponse) SetBody(v *GetFileUploadInfoResponseBody) *GetFileUploadInfoResponse {
	s.Body = v
	return s
}

type GetFlowIdByRelationEntityIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFlowIdByRelationEntityIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFlowIdByRelationEntityIdHeaders) GoString() string {
	return s.String()
}

func (s *GetFlowIdByRelationEntityIdHeaders) SetCommonHeaders(v map[string]*string) *GetFlowIdByRelationEntityIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFlowIdByRelationEntityIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetFlowIdByRelationEntityIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFlowIdByRelationEntityIdRequest struct {
	// 业务标识
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	// 招聘流程关联实体
	RelationEntity *string `json:"relationEntity,omitempty" xml:"relationEntity,omitempty"`
	// 招聘流程关联实体标识
	RelationEntityId *string `json:"relationEntityId,omitempty" xml:"relationEntityId,omitempty"`
}

func (s GetFlowIdByRelationEntityIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFlowIdByRelationEntityIdRequest) GoString() string {
	return s.String()
}

func (s *GetFlowIdByRelationEntityIdRequest) SetBizCode(v string) *GetFlowIdByRelationEntityIdRequest {
	s.BizCode = &v
	return s
}

func (s *GetFlowIdByRelationEntityIdRequest) SetRelationEntity(v string) *GetFlowIdByRelationEntityIdRequest {
	s.RelationEntity = &v
	return s
}

func (s *GetFlowIdByRelationEntityIdRequest) SetRelationEntityId(v string) *GetFlowIdByRelationEntityIdRequest {
	s.RelationEntityId = &v
	return s
}

type GetFlowIdByRelationEntityIdResponseBody struct {
	// 招聘流程标识
	FlowId *string `json:"flowId,omitempty" xml:"flowId,omitempty"`
}

func (s GetFlowIdByRelationEntityIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFlowIdByRelationEntityIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetFlowIdByRelationEntityIdResponseBody) SetFlowId(v string) *GetFlowIdByRelationEntityIdResponseBody {
	s.FlowId = &v
	return s
}

type GetFlowIdByRelationEntityIdResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFlowIdByRelationEntityIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFlowIdByRelationEntityIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFlowIdByRelationEntityIdResponse) GoString() string {
	return s.String()
}

func (s *GetFlowIdByRelationEntityIdResponse) SetHeaders(v map[string]*string) *GetFlowIdByRelationEntityIdResponse {
	s.Headers = v
	return s
}

func (s *GetFlowIdByRelationEntityIdResponse) SetBody(v *GetFlowIdByRelationEntityIdResponseBody) *GetFlowIdByRelationEntityIdResponse {
	s.Body = v
	return s
}

type GetJobAuthHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetJobAuthHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetJobAuthHeaders) GoString() string {
	return s.String()
}

func (s *GetJobAuthHeaders) SetCommonHeaders(v map[string]*string) *GetJobAuthHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetJobAuthHeaders) SetXAcsDingtalkAccessToken(v string) *GetJobAuthHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetJobAuthRequest struct {
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s GetJobAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobAuthRequest) GoString() string {
	return s.String()
}

func (s *GetJobAuthRequest) SetOpUserId(v string) *GetJobAuthRequest {
	s.OpUserId = &v
	return s
}

type GetJobAuthResponseBody struct {
	// 职位ID
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// 职位负责人
	JobOwners []*GetJobAuthResponseBodyJobOwners `json:"jobOwners,omitempty" xml:"jobOwners,omitempty" type:"Repeated"`
}

func (s GetJobAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobAuthResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobAuthResponseBody) SetJobId(v string) *GetJobAuthResponseBody {
	s.JobId = &v
	return s
}

func (s *GetJobAuthResponseBody) SetJobOwners(v []*GetJobAuthResponseBodyJobOwners) *GetJobAuthResponseBody {
	s.JobOwners = v
	return s
}

type GetJobAuthResponseBodyJobOwners struct {
	// 员工姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 员工标识
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetJobAuthResponseBodyJobOwners) String() string {
	return tea.Prettify(s)
}

func (s GetJobAuthResponseBodyJobOwners) GoString() string {
	return s.String()
}

func (s *GetJobAuthResponseBodyJobOwners) SetName(v string) *GetJobAuthResponseBodyJobOwners {
	s.Name = &v
	return s
}

func (s *GetJobAuthResponseBodyJobOwners) SetUserId(v string) *GetJobAuthResponseBodyJobOwners {
	s.UserId = &v
	return s
}

type GetJobAuthResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetJobAuthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobAuthResponse) GoString() string {
	return s.String()
}

func (s *GetJobAuthResponse) SetHeaders(v map[string]*string) *GetJobAuthResponse {
	s.Headers = v
	return s
}

func (s *GetJobAuthResponse) SetBody(v *GetJobAuthResponseBody) *GetJobAuthResponse {
	s.Body = v
	return s
}

type QueryInterviewsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryInterviewsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryInterviewsHeaders) GoString() string {
	return s.String()
}

func (s *QueryInterviewsHeaders) SetCommonHeaders(v map[string]*string) *QueryInterviewsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryInterviewsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryInterviewsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryInterviewsRequest struct {
	// 业务标识
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	// 候选人标识
	CandidateId *string `json:"candidateId,omitempty" xml:"candidateId,omitempty"`
	// 面试开始时间的查询起始时间（单位：毫秒）
	StartTimeBeginMillis *int64 `json:"startTimeBeginMillis,omitempty" xml:"startTimeBeginMillis,omitempty"`
	// 面试开始时间的查询结束时间（单位：毫秒）
	StartTimeEndMillis *int64 `json:"startTimeEndMillis,omitempty" xml:"startTimeEndMillis,omitempty"`
	// 分页游标，首次调用传空
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 分页大小
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s QueryInterviewsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInterviewsRequest) GoString() string {
	return s.String()
}

func (s *QueryInterviewsRequest) SetBizCode(v string) *QueryInterviewsRequest {
	s.BizCode = &v
	return s
}

func (s *QueryInterviewsRequest) SetCandidateId(v string) *QueryInterviewsRequest {
	s.CandidateId = &v
	return s
}

func (s *QueryInterviewsRequest) SetStartTimeBeginMillis(v int64) *QueryInterviewsRequest {
	s.StartTimeBeginMillis = &v
	return s
}

func (s *QueryInterviewsRequest) SetStartTimeEndMillis(v int64) *QueryInterviewsRequest {
	s.StartTimeEndMillis = &v
	return s
}

func (s *QueryInterviewsRequest) SetNextToken(v string) *QueryInterviewsRequest {
	s.NextToken = &v
	return s
}

func (s *QueryInterviewsRequest) SetSize(v int64) *QueryInterviewsRequest {
	s.Size = &v
	return s
}

type QueryInterviewsResponseBody struct {
	// 是否有更多数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 数据列表
	List []*QueryInterviewsResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// 下次查询的分页游标
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 总数量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryInterviewsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInterviewsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInterviewsResponseBody) SetHasMore(v bool) *QueryInterviewsResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryInterviewsResponseBody) SetList(v []*QueryInterviewsResponseBodyList) *QueryInterviewsResponseBody {
	s.List = v
	return s
}

func (s *QueryInterviewsResponseBody) SetNextToken(v string) *QueryInterviewsResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryInterviewsResponseBody) SetTotalCount(v int64) *QueryInterviewsResponseBody {
	s.TotalCount = &v
	return s
}

type QueryInterviewsResponseBodyList struct {
	// 面试是否已取消
	Cancelled *bool `json:"cancelled,omitempty" xml:"cancelled,omitempty"`
	// 面试创建人员工标识
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// 面试结束时间（单位：毫秒）
	EndTimeMillis *int64 `json:"endTimeMillis,omitempty" xml:"endTimeMillis,omitempty"`
	// 面试标识
	InterviewId *string `json:"interviewId,omitempty" xml:"interviewId,omitempty"`
	// 面试官列表
	Interviewers []*QueryInterviewsResponseBodyListInterviewers `json:"interviewers,omitempty" xml:"interviewers,omitempty" type:"Repeated"`
	// 职位标识
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// 面试开始时间（单位：毫秒）
	StartTimeMillis *int64 `json:"startTimeMillis,omitempty" xml:"startTimeMillis,omitempty"`
}

func (s QueryInterviewsResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryInterviewsResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryInterviewsResponseBodyList) SetCancelled(v bool) *QueryInterviewsResponseBodyList {
	s.Cancelled = &v
	return s
}

func (s *QueryInterviewsResponseBodyList) SetCreatorUserId(v string) *QueryInterviewsResponseBodyList {
	s.CreatorUserId = &v
	return s
}

func (s *QueryInterviewsResponseBodyList) SetEndTimeMillis(v int64) *QueryInterviewsResponseBodyList {
	s.EndTimeMillis = &v
	return s
}

func (s *QueryInterviewsResponseBodyList) SetInterviewId(v string) *QueryInterviewsResponseBodyList {
	s.InterviewId = &v
	return s
}

func (s *QueryInterviewsResponseBodyList) SetInterviewers(v []*QueryInterviewsResponseBodyListInterviewers) *QueryInterviewsResponseBodyList {
	s.Interviewers = v
	return s
}

func (s *QueryInterviewsResponseBodyList) SetJobId(v string) *QueryInterviewsResponseBodyList {
	s.JobId = &v
	return s
}

func (s *QueryInterviewsResponseBodyList) SetStartTimeMillis(v int64) *QueryInterviewsResponseBodyList {
	s.StartTimeMillis = &v
	return s
}

type QueryInterviewsResponseBodyListInterviewers struct {
	// 面试官员工标识
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryInterviewsResponseBodyListInterviewers) String() string {
	return tea.Prettify(s)
}

func (s QueryInterviewsResponseBodyListInterviewers) GoString() string {
	return s.String()
}

func (s *QueryInterviewsResponseBodyListInterviewers) SetUserId(v string) *QueryInterviewsResponseBodyListInterviewers {
	s.UserId = &v
	return s
}

type QueryInterviewsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryInterviewsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryInterviewsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryInterviewsResponse) GoString() string {
	return s.String()
}

func (s *QueryInterviewsResponse) SetHeaders(v map[string]*string) *QueryInterviewsResponse {
	s.Headers = v
	return s
}

func (s *QueryInterviewsResponse) SetBody(v *QueryInterviewsResponseBody) *QueryInterviewsResponse {
	s.Body = v
	return s
}

type ReportMessageStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ReportMessageStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s ReportMessageStatusHeaders) GoString() string {
	return s.String()
}

func (s *ReportMessageStatusHeaders) SetCommonHeaders(v map[string]*string) *ReportMessageStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReportMessageStatusHeaders) SetXAcsDingtalkAccessToken(v string) *ReportMessageStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ReportMessageStatusRequest struct {
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	// 渠道标识。
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 错误码。
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息描述。
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// 消息ID。
	MessageId      *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	ReceiverUserId *string `json:"receiverUserId,omitempty" xml:"receiverUserId,omitempty"`
	SenderUserId   *string `json:"senderUserId,omitempty" xml:"senderUserId,omitempty"`
}

func (s ReportMessageStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportMessageStatusRequest) GoString() string {
	return s.String()
}

func (s *ReportMessageStatusRequest) SetBizCode(v string) *ReportMessageStatusRequest {
	s.BizCode = &v
	return s
}

func (s *ReportMessageStatusRequest) SetChannel(v string) *ReportMessageStatusRequest {
	s.Channel = &v
	return s
}

func (s *ReportMessageStatusRequest) SetErrorCode(v string) *ReportMessageStatusRequest {
	s.ErrorCode = &v
	return s
}

func (s *ReportMessageStatusRequest) SetErrorMsg(v string) *ReportMessageStatusRequest {
	s.ErrorMsg = &v
	return s
}

func (s *ReportMessageStatusRequest) SetMessageId(v string) *ReportMessageStatusRequest {
	s.MessageId = &v
	return s
}

func (s *ReportMessageStatusRequest) SetReceiverUserId(v string) *ReportMessageStatusRequest {
	s.ReceiverUserId = &v
	return s
}

func (s *ReportMessageStatusRequest) SetSenderUserId(v string) *ReportMessageStatusRequest {
	s.SenderUserId = &v
	return s
}

type ReportMessageStatusResponseBody struct {
	// Id of the request
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ReportMessageStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportMessageStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ReportMessageStatusResponseBody) SetResult(v string) *ReportMessageStatusResponseBody {
	s.Result = &v
	return s
}

type ReportMessageStatusResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReportMessageStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReportMessageStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportMessageStatusResponse) GoString() string {
	return s.String()
}

func (s *ReportMessageStatusResponse) SetHeaders(v map[string]*string) *ReportMessageStatusResponse {
	s.Headers = v
	return s
}

func (s *ReportMessageStatusResponse) SetBody(v *ReportMessageStatusResponseBody) *ReportMessageStatusResponse {
	s.Body = v
	return s
}

type SyncChannelMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SyncChannelMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncChannelMessageHeaders) GoString() string {
	return s.String()
}

func (s *SyncChannelMessageHeaders) SetCommonHeaders(v map[string]*string) *SyncChannelMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncChannelMessageHeaders) SetXAcsDingtalkAccessToken(v string) *SyncChannelMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SyncChannelMessageRequest struct {
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	// 渠道标识。
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 消息内容。
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 消息创建时间。
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 消息接收者ID。
	ReceiverUserId *string `json:"receiverUserId,omitempty" xml:"receiverUserId,omitempty"`
	// 消息发送者用户ID。
	SenderUserId *string `json:"senderUserId,omitempty" xml:"senderUserId,omitempty"`
	// 消息UUID，业务方产生用于去重。
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s SyncChannelMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncChannelMessageRequest) GoString() string {
	return s.String()
}

func (s *SyncChannelMessageRequest) SetBizCode(v string) *SyncChannelMessageRequest {
	s.BizCode = &v
	return s
}

func (s *SyncChannelMessageRequest) SetChannel(v string) *SyncChannelMessageRequest {
	s.Channel = &v
	return s
}

func (s *SyncChannelMessageRequest) SetContent(v string) *SyncChannelMessageRequest {
	s.Content = &v
	return s
}

func (s *SyncChannelMessageRequest) SetCreateTime(v int64) *SyncChannelMessageRequest {
	s.CreateTime = &v
	return s
}

func (s *SyncChannelMessageRequest) SetReceiverUserId(v string) *SyncChannelMessageRequest {
	s.ReceiverUserId = &v
	return s
}

func (s *SyncChannelMessageRequest) SetSenderUserId(v string) *SyncChannelMessageRequest {
	s.SenderUserId = &v
	return s
}

func (s *SyncChannelMessageRequest) SetUuid(v string) *SyncChannelMessageRequest {
	s.Uuid = &v
	return s
}

type SyncChannelMessageResponseBody struct {
	// Id of the request
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SyncChannelMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncChannelMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SyncChannelMessageResponseBody) SetResult(v string) *SyncChannelMessageResponseBody {
	s.Result = &v
	return s
}

type SyncChannelMessageResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SyncChannelMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncChannelMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncChannelMessageResponse) GoString() string {
	return s.String()
}

func (s *SyncChannelMessageResponse) SetHeaders(v map[string]*string) *SyncChannelMessageResponse {
	s.Headers = v
	return s
}

func (s *SyncChannelMessageResponse) SetBody(v *SyncChannelMessageResponseBody) *SyncChannelMessageResponse {
	s.Body = v
	return s
}

type UpdateApplicationRegFormHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateApplicationRegFormHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationRegFormHeaders) GoString() string {
	return s.String()
}

func (s *UpdateApplicationRegFormHeaders) SetCommonHeaders(v map[string]*string) *UpdateApplicationRegFormHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateApplicationRegFormHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateApplicationRegFormHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateApplicationRegFormRequest struct {
	// 业务标识
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	// 应聘登记表的表单内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 钉盘文件信息
	DingPanFile *UpdateApplicationRegFormRequestDingPanFile `json:"dingPanFile,omitempty" xml:"dingPanFile,omitempty" type:"Struct"`
}

func (s UpdateApplicationRegFormRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationRegFormRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationRegFormRequest) SetBizCode(v string) *UpdateApplicationRegFormRequest {
	s.BizCode = &v
	return s
}

func (s *UpdateApplicationRegFormRequest) SetContent(v string) *UpdateApplicationRegFormRequest {
	s.Content = &v
	return s
}

func (s *UpdateApplicationRegFormRequest) SetDingPanFile(v *UpdateApplicationRegFormRequestDingPanFile) *UpdateApplicationRegFormRequest {
	s.DingPanFile = v
	return s
}

type UpdateApplicationRegFormRequestDingPanFile struct {
	// 钉盘文件标识
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// 文件名
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件大小（单位：字节）
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// 文件类型（支持：pdf，doc，docx，ppt，pptx，jpg等）
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// 钉盘空间标识
	SpaceId *int64 `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s UpdateApplicationRegFormRequestDingPanFile) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationRegFormRequestDingPanFile) GoString() string {
	return s.String()
}

func (s *UpdateApplicationRegFormRequestDingPanFile) SetFileId(v string) *UpdateApplicationRegFormRequestDingPanFile {
	s.FileId = &v
	return s
}

func (s *UpdateApplicationRegFormRequestDingPanFile) SetFileName(v string) *UpdateApplicationRegFormRequestDingPanFile {
	s.FileName = &v
	return s
}

func (s *UpdateApplicationRegFormRequestDingPanFile) SetFileSize(v int64) *UpdateApplicationRegFormRequestDingPanFile {
	s.FileSize = &v
	return s
}

func (s *UpdateApplicationRegFormRequestDingPanFile) SetFileType(v string) *UpdateApplicationRegFormRequestDingPanFile {
	s.FileType = &v
	return s
}

func (s *UpdateApplicationRegFormRequestDingPanFile) SetSpaceId(v int64) *UpdateApplicationRegFormRequestDingPanFile {
	s.SpaceId = &v
	return s
}

type UpdateApplicationRegFormResponseBody struct {
	// 邀填人员工标识
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// 表单标识
	FormId *string `json:"formId,omitempty" xml:"formId,omitempty"`
	// 创建时间（邀填时间，单位：毫秒）
	GmtCreateMillis *int64 `json:"gmtCreateMillis,omitempty" xml:"gmtCreateMillis,omitempty"`
	// 更新时间（填写时间，单位：毫秒），仅当表单状态为已填写时有效
	GmtModifiedMillis *int64 `json:"gmtModifiedMillis,omitempty" xml:"gmtModifiedMillis,omitempty"`
	// 表单状态（0表示未填写，1表示已填写）
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// 模板标识
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// 模板版本
	TemplateVersion *int32 `json:"templateVersion,omitempty" xml:"templateVersion,omitempty"`
}

func (s UpdateApplicationRegFormResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationRegFormResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationRegFormResponseBody) SetCreatorUserId(v string) *UpdateApplicationRegFormResponseBody {
	s.CreatorUserId = &v
	return s
}

func (s *UpdateApplicationRegFormResponseBody) SetFormId(v string) *UpdateApplicationRegFormResponseBody {
	s.FormId = &v
	return s
}

func (s *UpdateApplicationRegFormResponseBody) SetGmtCreateMillis(v int64) *UpdateApplicationRegFormResponseBody {
	s.GmtCreateMillis = &v
	return s
}

func (s *UpdateApplicationRegFormResponseBody) SetGmtModifiedMillis(v int64) *UpdateApplicationRegFormResponseBody {
	s.GmtModifiedMillis = &v
	return s
}

func (s *UpdateApplicationRegFormResponseBody) SetStatus(v int32) *UpdateApplicationRegFormResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateApplicationRegFormResponseBody) SetTemplateId(v string) *UpdateApplicationRegFormResponseBody {
	s.TemplateId = &v
	return s
}

func (s *UpdateApplicationRegFormResponseBody) SetTemplateVersion(v int32) *UpdateApplicationRegFormResponseBody {
	s.TemplateVersion = &v
	return s
}

type UpdateApplicationRegFormResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateApplicationRegFormResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateApplicationRegFormResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationRegFormResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationRegFormResponse) SetHeaders(v map[string]*string) *UpdateApplicationRegFormResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationRegFormResponse) SetBody(v *UpdateApplicationRegFormResponseBody) *UpdateApplicationRegFormResponse {
	s.Body = v
	return s
}

type UpdateInterviewSignInInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateInterviewSignInInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateInterviewSignInInfoHeaders) GoString() string {
	return s.String()
}

func (s *UpdateInterviewSignInInfoHeaders) SetCommonHeaders(v map[string]*string) *UpdateInterviewSignInInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateInterviewSignInInfoHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateInterviewSignInInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateInterviewSignInInfoRequest struct {
	// 业务标识
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	// 面试签到时间（单位：毫秒）
	SignInTimeMillis *int64 `json:"signInTimeMillis,omitempty" xml:"signInTimeMillis,omitempty"`
}

func (s UpdateInterviewSignInInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInterviewSignInInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateInterviewSignInInfoRequest) SetBizCode(v string) *UpdateInterviewSignInInfoRequest {
	s.BizCode = &v
	return s
}

func (s *UpdateInterviewSignInInfoRequest) SetSignInTimeMillis(v int64) *UpdateInterviewSignInInfoRequest {
	s.SignInTimeMillis = &v
	return s
}

type UpdateInterviewSignInInfoResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpdateInterviewSignInInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInterviewSignInInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateInterviewSignInInfoResponse) SetHeaders(v map[string]*string) *UpdateInterviewSignInInfoResponse {
	s.Headers = v
	return s
}

type UpdateJobDeliverHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateJobDeliverHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobDeliverHeaders) GoString() string {
	return s.String()
}

func (s *UpdateJobDeliverHeaders) SetCommonHeaders(v map[string]*string) *UpdateJobDeliverHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateJobDeliverHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateJobDeliverHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateJobDeliverRequest struct {
	// 招聘业务标识，目前固定ddats
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	// 渠道侧职位唯一标识
	ChannelOuterId *string `json:"channelOuterId,omitempty" xml:"channelOuterId,omitempty"`
	// 职位投递人userId
	DeliverUserId *string `json:"deliverUserId,omitempty" xml:"deliverUserId,omitempty"`
	// 渠道侧错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 渠道侧错误信息
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// 操作时间
	OpTime *int64 `json:"opTime,omitempty" xml:"opTime,omitempty"`
	// 操作人userId
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 职位投递状态
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// 钉钉侧职位唯一标识
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
}

func (s UpdateJobDeliverRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobDeliverRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobDeliverRequest) SetBizCode(v string) *UpdateJobDeliverRequest {
	s.BizCode = &v
	return s
}

func (s *UpdateJobDeliverRequest) SetChannelOuterId(v string) *UpdateJobDeliverRequest {
	s.ChannelOuterId = &v
	return s
}

func (s *UpdateJobDeliverRequest) SetDeliverUserId(v string) *UpdateJobDeliverRequest {
	s.DeliverUserId = &v
	return s
}

func (s *UpdateJobDeliverRequest) SetErrorCode(v string) *UpdateJobDeliverRequest {
	s.ErrorCode = &v
	return s
}

func (s *UpdateJobDeliverRequest) SetErrorMsg(v string) *UpdateJobDeliverRequest {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateJobDeliverRequest) SetOpTime(v int64) *UpdateJobDeliverRequest {
	s.OpTime = &v
	return s
}

func (s *UpdateJobDeliverRequest) SetOpUserId(v string) *UpdateJobDeliverRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateJobDeliverRequest) SetStatus(v int32) *UpdateJobDeliverRequest {
	s.Status = &v
	return s
}

func (s *UpdateJobDeliverRequest) SetJobId(v string) *UpdateJobDeliverRequest {
	s.JobId = &v
	return s
}

type UpdateJobDeliverResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateJobDeliverResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobDeliverResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateJobDeliverResponseBody) SetSuccess(v bool) *UpdateJobDeliverResponseBody {
	s.Success = &v
	return s
}

type UpdateJobDeliverResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateJobDeliverResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateJobDeliverResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobDeliverResponse) GoString() string {
	return s.String()
}

func (s *UpdateJobDeliverResponse) SetHeaders(v map[string]*string) *UpdateJobDeliverResponse {
	s.Headers = v
	return s
}

func (s *UpdateJobDeliverResponse) SetBody(v *UpdateJobDeliverResponseBody) *UpdateJobDeliverResponse {
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

func (client *Client) AddApplicationRegFormTemplate(request *AddApplicationRegFormTemplateRequest) (_result *AddApplicationRegFormTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddApplicationRegFormTemplateHeaders{}
	_result = &AddApplicationRegFormTemplateResponse{}
	_body, _err := client.AddApplicationRegFormTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddApplicationRegFormTemplateWithOptions(request *AddApplicationRegFormTemplateRequest, headers *AddApplicationRegFormTemplateHeaders, runtime *util.RuntimeOptions) (_result *AddApplicationRegFormTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["bizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OuterId)) {
		body["outerId"] = request.OuterId
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &AddApplicationRegFormTemplateResponse{}
	_body, _err := client.DoROARequest(tea.String("AddApplicationRegFormTemplate"), tea.String("ats_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/ats/flows/applicationRegForms/templates"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddFile(request *AddFileRequest) (_result *AddFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddFileHeaders{}
	_result = &AddFileResponse{}
	_body, _err := client.AddFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddFileWithOptions(request *AddFileRequest, headers *AddFileHeaders, runtime *util.RuntimeOptions) (_result *AddFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["bizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.MediaId)) {
		body["mediaId"] = request.MediaId
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &AddFileResponse{}
	_body, _err := client.DoROARequest(tea.String("AddFile"), tea.String("ats_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/ats/files"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddUserAccount(request *AddUserAccountRequest) (_result *AddUserAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddUserAccountHeaders{}
	_result = &AddUserAccountResponse{}
	_body, _err := client.AddUserAccountWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddUserAccountWithOptions(request *AddUserAccountRequest, headers *AddUserAccountHeaders, runtime *util.RuntimeOptions) (_result *AddUserAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["bizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelAccountName)) {
		body["channelAccountName"] = request.ChannelAccountName
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelUserIdentify)) {
		body["channelUserIdentify"] = request.ChannelUserIdentify
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		body["phoneNumber"] = request.PhoneNumber
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &AddUserAccountResponse{}
	_body, _err := client.DoROARequest(tea.String("AddUserAccount"), tea.String("ats_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/ats/channels/users/accounts"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CollectRecruitJobDetail(request *CollectRecruitJobDetailRequest) (_result *CollectRecruitJobDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CollectRecruitJobDetailHeaders{}
	_result = &CollectRecruitJobDetailResponse{}
	_body, _err := client.CollectRecruitJobDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CollectRecruitJobDetailWithOptions(request *CollectRecruitJobDetailRequest, headers *CollectRecruitJobDetailHeaders, runtime *util.RuntimeOptions) (_result *CollectRecruitJobDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["bizCode"] = request.BizCode
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		body["channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.JonInfo)) {
		body["jonInfo"] = request.JonInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OutCorpId)) {
		body["outCorpId"] = request.OutCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.OutCorpName)) {
		body["outCorpName"] = request.OutCorpName
	}

	if !tea.BoolValue(util.IsUnset(request.RecruitUserInfo)) {
		body["recruitUserInfo"] = request.RecruitUserInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateTime)) {
		body["updateTime"] = request.UpdateTime
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CollectRecruitJobDetailResponse{}
	_body, _err := client.DoROARequest(tea.String("CollectRecruitJobDetail"), tea.String("ats_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/ats/channels/jobs/import"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CollectResumeDetail(request *CollectResumeDetailRequest) (_result *CollectResumeDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CollectResumeDetailHeaders{}
	_result = &CollectResumeDetailResponse{}
	_body, _err := client.CollectResumeDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CollectResumeDetailWithOptions(request *CollectResumeDetailRequest, headers *CollectResumeDetailHeaders, runtime *util.RuntimeOptions) (_result *CollectResumeDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["bizCode"] = request.BizCode
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelOuterId)) {
		body["channelOuterId"] = request.ChannelOuterId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelTalentId)) {
		body["channelTalentId"] = request.ChannelTalentId
	}

	if !tea.BoolValue(util.IsUnset(request.DeliverJobId)) {
		body["deliverJobId"] = request.DeliverJobId
	}

	if !tea.BoolValue(util.IsUnset(request.OptUserId)) {
		body["optUserId"] = request.OptUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ResumeData)) {
		body["resumeData"] = request.ResumeData
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CollectResumeDetailResponse{}
	_body, _err := client.DoROARequest(tea.String("CollectResumeDetail"), tea.String("ats_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/ats/resumes/details"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfirmRights(rightsCode *string, request *ConfirmRightsRequest) (_result *ConfirmRightsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ConfirmRightsHeaders{}
	_result = &ConfirmRightsResponse{}
	_body, _err := client.ConfirmRightsWithOptions(rightsCode, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfirmRightsWithOptions(rightsCode *string, request *ConfirmRightsRequest, headers *ConfirmRightsHeaders, runtime *util.RuntimeOptions) (_result *ConfirmRightsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	rightsCode = openapiutil.GetEncodeParam(rightsCode)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["bizCode"] = request.BizCode
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
	_result = &ConfirmRightsResponse{}
	_body, _err := client.DoROARequest(tea.String("ConfirmRights"), tea.String("ats_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/ats/rights/"+tea.StringValue(rightsCode)+"/confirm"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FinishBeginnerTask(taskCode *string, request *FinishBeginnerTaskRequest) (_result *FinishBeginnerTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &FinishBeginnerTaskHeaders{}
	_result = &FinishBeginnerTaskResponse{}
	_body, _err := client.FinishBeginnerTaskWithOptions(taskCode, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FinishBeginnerTaskWithOptions(taskCode *string, request *FinishBeginnerTaskRequest, headers *FinishBeginnerTaskHeaders, runtime *util.RuntimeOptions) (_result *FinishBeginnerTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	taskCode = openapiutil.GetEncodeParam(taskCode)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["scope"] = request.Scope
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
	_result = &FinishBeginnerTaskResponse{}
	_body, _err := client.DoROARequest(tea.String("FinishBeginnerTask"), tea.String("ats_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/ats/beginnerTasks/"+tea.StringValue(taskCode)+"/finish"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetApplicationRegFormByFlowId(flowId *string, request *GetApplicationRegFormByFlowIdRequest) (_result *GetApplicationRegFormByFlowIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetApplicationRegFormByFlowIdHeaders{}
	_result = &GetApplicationRegFormByFlowIdResponse{}
	_body, _err := client.GetApplicationRegFormByFlowIdWithOptions(flowId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetApplicationRegFormByFlowIdWithOptions(flowId *string, request *GetApplicationRegFormByFlowIdRequest, headers *GetApplicationRegFormByFlowIdHeaders, runtime *util.RuntimeOptions) (_result *GetApplicationRegFormByFlowIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	flowId = openapiutil.GetEncodeParam(flowId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["bizCode"] = request.BizCode
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
	_result = &GetApplicationRegFormByFlowIdResponse{}
	_body, _err := client.DoROARequest(tea.String("GetApplicationRegFormByFlowId"), tea.String("ats_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/ats/flows/"+tea.StringValue(flowId)+"/applicationRegForms"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCandidateByPhoneNumber(request *GetCandidateByPhoneNumberRequest) (_result *GetCandidateByPhoneNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCandidateByPhoneNumberHeaders{}
	_result = &GetCandidateByPhoneNumberResponse{}
	_body, _err := client.GetCandidateByPhoneNumberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCandidateByPhoneNumberWithOptions(request *GetCandidateByPhoneNumberRequest, headers *GetCandidateByPhoneNumberHeaders, runtime *util.RuntimeOptions) (_result *GetCandidateByPhoneNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["bizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["phoneNumber"] = request.PhoneNumber
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
	_result = &GetCandidateByPhoneNumberResponse{}
	_body, _err := client.DoROARequest(tea.String("GetCandidateByPhoneNumber"), tea.String("ats_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/ats/candidates"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFileUploadInfo(request *GetFileUploadInfoRequest) (_result *GetFileUploadInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFileUploadInfoHeaders{}
	_result = &GetFileUploadInfoResponse{}
	_body, _err := client.GetFileUploadInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFileUploadInfoWithOptions(request *GetFileUploadInfoRequest, headers *GetFileUploadInfoHeaders, runtime *util.RuntimeOptions) (_result *GetFileUploadInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["bizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileSize)) {
		query["fileSize"] = request.FileSize
	}

	if !tea.BoolValue(util.IsUnset(request.Md5)) {
		query["md5"] = request.Md5
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
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
	_result = &GetFileUploadInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetFileUploadInfo"), tea.String("ats_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/ats/files/uploadInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFlowIdByRelationEntityId(request *GetFlowIdByRelationEntityIdRequest) (_result *GetFlowIdByRelationEntityIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFlowIdByRelationEntityIdHeaders{}
	_result = &GetFlowIdByRelationEntityIdResponse{}
	_body, _err := client.GetFlowIdByRelationEntityIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFlowIdByRelationEntityIdWithOptions(request *GetFlowIdByRelationEntityIdRequest, headers *GetFlowIdByRelationEntityIdHeaders, runtime *util.RuntimeOptions) (_result *GetFlowIdByRelationEntityIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["bizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.RelationEntity)) {
		query["relationEntity"] = request.RelationEntity
	}

	if !tea.BoolValue(util.IsUnset(request.RelationEntityId)) {
		query["relationEntityId"] = request.RelationEntityId
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
	_result = &GetFlowIdByRelationEntityIdResponse{}
	_body, _err := client.DoROARequest(tea.String("GetFlowIdByRelationEntityId"), tea.String("ats_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/ats/flows/ids"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobAuth(jobId *string, request *GetJobAuthRequest) (_result *GetJobAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetJobAuthHeaders{}
	_result = &GetJobAuthResponse{}
	_body, _err := client.GetJobAuthWithOptions(jobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobAuthWithOptions(jobId *string, request *GetJobAuthRequest, headers *GetJobAuthHeaders, runtime *util.RuntimeOptions) (_result *GetJobAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	jobId = openapiutil.GetEncodeParam(jobId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
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
	_result = &GetJobAuthResponse{}
	_body, _err := client.DoROARequest(tea.String("GetJobAuth"), tea.String("ats_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/ats/auths/jobs/"+tea.StringValue(jobId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryInterviews(request *QueryInterviewsRequest) (_result *QueryInterviewsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryInterviewsHeaders{}
	_result = &QueryInterviewsResponse{}
	_body, _err := client.QueryInterviewsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryInterviewsWithOptions(request *QueryInterviewsRequest, headers *QueryInterviewsHeaders, runtime *util.RuntimeOptions) (_result *QueryInterviewsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["bizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CandidateId)) {
		body["candidateId"] = request.CandidateId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimeBeginMillis)) {
		body["startTimeBeginMillis"] = request.StartTimeBeginMillis
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimeEndMillis)) {
		body["startTimeEndMillis"] = request.StartTimeEndMillis
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &QueryInterviewsResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryInterviews"), tea.String("ats_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/ats/interviews/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportMessageStatus(request *ReportMessageStatusRequest) (_result *ReportMessageStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ReportMessageStatusHeaders{}
	_result = &ReportMessageStatusResponse{}
	_body, _err := client.ReportMessageStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportMessageStatusWithOptions(request *ReportMessageStatusRequest, headers *ReportMessageStatusHeaders, runtime *util.RuntimeOptions) (_result *ReportMessageStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["bizCode"] = request.BizCode
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		body["channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.ErrorCode)) {
		body["errorCode"] = request.ErrorCode
	}

	if !tea.BoolValue(util.IsUnset(request.ErrorMsg)) {
		body["errorMsg"] = request.ErrorMsg
	}

	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		body["messageId"] = request.MessageId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUserId)) {
		body["receiverUserId"] = request.ReceiverUserId
	}

	if !tea.BoolValue(util.IsUnset(request.SenderUserId)) {
		body["senderUserId"] = request.SenderUserId
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &ReportMessageStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("ReportMessageStatus"), tea.String("ats_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/ats/channels/messages/statuses/report"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncChannelMessage(request *SyncChannelMessageRequest) (_result *SyncChannelMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncChannelMessageHeaders{}
	_result = &SyncChannelMessageResponse{}
	_body, _err := client.SyncChannelMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncChannelMessageWithOptions(request *SyncChannelMessageRequest, headers *SyncChannelMessageHeaders, runtime *util.RuntimeOptions) (_result *SyncChannelMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["bizCode"] = request.BizCode
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		body["channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTime)) {
		body["createTime"] = request.CreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUserId)) {
		body["receiverUserId"] = request.ReceiverUserId
	}

	if !tea.BoolValue(util.IsUnset(request.SenderUserId)) {
		body["senderUserId"] = request.SenderUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &SyncChannelMessageResponse{}
	_body, _err := client.DoROARequest(tea.String("SyncChannelMessage"), tea.String("ats_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/ats/channels/messages/sync"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateApplicationRegForm(flowId *string, request *UpdateApplicationRegFormRequest) (_result *UpdateApplicationRegFormResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateApplicationRegFormHeaders{}
	_result = &UpdateApplicationRegFormResponse{}
	_body, _err := client.UpdateApplicationRegFormWithOptions(flowId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateApplicationRegFormWithOptions(flowId *string, request *UpdateApplicationRegFormRequest, headers *UpdateApplicationRegFormHeaders, runtime *util.RuntimeOptions) (_result *UpdateApplicationRegFormResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	flowId = openapiutil.GetEncodeParam(flowId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["bizCode"] = request.BizCode
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DingPanFile)) {
		body["dingPanFile"] = request.DingPanFile
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateApplicationRegFormResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateApplicationRegForm"), tea.String("ats_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/ats/flows/"+tea.StringValue(flowId)+"/applicationRegForms"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInterviewSignInInfo(interviewId *string, request *UpdateInterviewSignInInfoRequest) (_result *UpdateInterviewSignInInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateInterviewSignInInfoHeaders{}
	_result = &UpdateInterviewSignInInfoResponse{}
	_body, _err := client.UpdateInterviewSignInInfoWithOptions(interviewId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInterviewSignInInfoWithOptions(interviewId *string, request *UpdateInterviewSignInInfoRequest, headers *UpdateInterviewSignInInfoHeaders, runtime *util.RuntimeOptions) (_result *UpdateInterviewSignInInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	interviewId = openapiutil.GetEncodeParam(interviewId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["bizCode"] = request.BizCode
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SignInTimeMillis)) {
		body["signInTimeMillis"] = request.SignInTimeMillis
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateInterviewSignInInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateInterviewSignInInfo"), tea.String("ats_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/ats/interviews/"+tea.StringValue(interviewId)+"/signInInfos"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateJobDeliver(request *UpdateJobDeliverRequest) (_result *UpdateJobDeliverResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateJobDeliverHeaders{}
	_result = &UpdateJobDeliverResponse{}
	_body, _err := client.UpdateJobDeliverWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateJobDeliverWithOptions(request *UpdateJobDeliverRequest, headers *UpdateJobDeliverHeaders, runtime *util.RuntimeOptions) (_result *UpdateJobDeliverResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["bizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["jobId"] = request.JobId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelOuterId)) {
		body["channelOuterId"] = request.ChannelOuterId
	}

	if !tea.BoolValue(util.IsUnset(request.DeliverUserId)) {
		body["deliverUserId"] = request.DeliverUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ErrorCode)) {
		body["errorCode"] = request.ErrorCode
	}

	if !tea.BoolValue(util.IsUnset(request.ErrorMsg)) {
		body["errorMsg"] = request.ErrorMsg
	}

	if !tea.BoolValue(util.IsUnset(request.OpTime)) {
		body["opTime"] = request.OpTime
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateJobDeliverResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateJobDeliver"), tea.String("ats_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/ats/jobs/deliveryStatus"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
