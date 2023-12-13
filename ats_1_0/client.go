// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package ats_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
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
	BizCode  *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	Content  *string `json:"content,omitempty" xml:"content,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	OuterId  *string `json:"outerId,omitempty" xml:"outerId,omitempty"`
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
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	Version    *int32  `json:"version,omitempty" xml:"version,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddApplicationRegFormTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddApplicationRegFormTemplateResponse) SetStatusCode(v int32) *AddApplicationRegFormTemplateResponse {
	s.StatusCode = &v
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
	BizCode  *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	MediaId  *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
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
	FileId   *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	SpaceId  *int64  `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
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
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddFileResponse) SetStatusCode(v int32) *AddFileResponse {
	s.StatusCode = &v
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
	BizCode             *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	ChannelAccountName  *string `json:"channelAccountName,omitempty" xml:"channelAccountName,omitempty"`
	ChannelUserIdentify *string `json:"channelUserIdentify,omitempty" xml:"channelUserIdentify,omitempty"`
	PhoneNumber         *string `json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	CorpId              *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	UserId              *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddUserAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddUserAccountResponse) SetStatusCode(v int32) *AddUserAccountResponse {
	s.StatusCode = &v
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
	BizCode         *string                                        `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	Channel         *string                                        `json:"channel,omitempty" xml:"channel,omitempty"`
	JobInfo         *CollectRecruitJobDetailRequestJobInfo         `json:"jobInfo,omitempty" xml:"jobInfo,omitempty" type:"Struct"`
	OutCorpId       *string                                        `json:"outCorpId,omitempty" xml:"outCorpId,omitempty"`
	OutCorpName     *string                                        `json:"outCorpName,omitempty" xml:"outCorpName,omitempty"`
	RecruitUserInfo *CollectRecruitJobDetailRequestRecruitUserInfo `json:"recruitUserInfo,omitempty" xml:"recruitUserInfo,omitempty" type:"Struct"`
	Source          *string                                        `json:"source,omitempty" xml:"source,omitempty"`
	UpdateTime      *int64                                         `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
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

func (s *CollectRecruitJobDetailRequest) SetJobInfo(v *CollectRecruitJobDetailRequestJobInfo) *CollectRecruitJobDetailRequest {
	s.JobInfo = v
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

type CollectRecruitJobDetailRequestJobInfo struct {
	Address      *CollectRecruitJobDetailRequestJobInfoAddress      `json:"address,omitempty" xml:"address,omitempty" type:"Struct"`
	Category     *string                                            `json:"category,omitempty" xml:"category,omitempty"`
	Description  *string                                            `json:"description,omitempty" xml:"description,omitempty"`
	ExtInfo      *string                                            `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	FullTimeInfo *CollectRecruitJobDetailRequestJobInfoFullTimeInfo `json:"fullTimeInfo,omitempty" xml:"fullTimeInfo,omitempty" type:"Struct"`
	HeadCount    *string                                            `json:"headCount,omitempty" xml:"headCount,omitempty"`
	JobNature    *string                                            `json:"jobNature,omitempty" xml:"jobNature,omitempty"`
	JobTags      []*string                                          `json:"jobTags,omitempty" xml:"jobTags,omitempty" type:"Repeated"`
	MaxSalary    *string                                            `json:"maxSalary,omitempty" xml:"maxSalary,omitempty"`
	MinSalary    *string                                            `json:"minSalary,omitempty" xml:"minSalary,omitempty"`
	Name         *string                                            `json:"name,omitempty" xml:"name,omitempty"`
	OutJobId     *string                                            `json:"outJobId,omitempty" xml:"outJobId,omitempty"`
	PartTimeInfo *CollectRecruitJobDetailRequestJobInfoPartTimeInfo `json:"partTimeInfo,omitempty" xml:"partTimeInfo,omitempty" type:"Struct"`
	RequiredEdu  *string                                            `json:"requiredEdu,omitempty" xml:"requiredEdu,omitempty"`
}

func (s CollectRecruitJobDetailRequestJobInfo) String() string {
	return tea.Prettify(s)
}

func (s CollectRecruitJobDetailRequestJobInfo) GoString() string {
	return s.String()
}

func (s *CollectRecruitJobDetailRequestJobInfo) SetAddress(v *CollectRecruitJobDetailRequestJobInfoAddress) *CollectRecruitJobDetailRequestJobInfo {
	s.Address = v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfo) SetCategory(v string) *CollectRecruitJobDetailRequestJobInfo {
	s.Category = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfo) SetDescription(v string) *CollectRecruitJobDetailRequestJobInfo {
	s.Description = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfo) SetExtInfo(v string) *CollectRecruitJobDetailRequestJobInfo {
	s.ExtInfo = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfo) SetFullTimeInfo(v *CollectRecruitJobDetailRequestJobInfoFullTimeInfo) *CollectRecruitJobDetailRequestJobInfo {
	s.FullTimeInfo = v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfo) SetHeadCount(v string) *CollectRecruitJobDetailRequestJobInfo {
	s.HeadCount = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfo) SetJobNature(v string) *CollectRecruitJobDetailRequestJobInfo {
	s.JobNature = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfo) SetJobTags(v []*string) *CollectRecruitJobDetailRequestJobInfo {
	s.JobTags = v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfo) SetMaxSalary(v string) *CollectRecruitJobDetailRequestJobInfo {
	s.MaxSalary = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfo) SetMinSalary(v string) *CollectRecruitJobDetailRequestJobInfo {
	s.MinSalary = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfo) SetName(v string) *CollectRecruitJobDetailRequestJobInfo {
	s.Name = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfo) SetOutJobId(v string) *CollectRecruitJobDetailRequestJobInfo {
	s.OutJobId = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfo) SetPartTimeInfo(v *CollectRecruitJobDetailRequestJobInfoPartTimeInfo) *CollectRecruitJobDetailRequestJobInfo {
	s.PartTimeInfo = v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfo) SetRequiredEdu(v string) *CollectRecruitJobDetailRequestJobInfo {
	s.RequiredEdu = &v
	return s
}

type CollectRecruitJobDetailRequestJobInfoAddress struct {
	CityCode     *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	Detail       *string `json:"detail,omitempty" xml:"detail,omitempty"`
	DistrictCode *string `json:"districtCode,omitempty" xml:"districtCode,omitempty"`
	Latitude     *string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	Longitude    *string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	ProvinceCode *string `json:"provinceCode,omitempty" xml:"provinceCode,omitempty"`
}

func (s CollectRecruitJobDetailRequestJobInfoAddress) String() string {
	return tea.Prettify(s)
}

func (s CollectRecruitJobDetailRequestJobInfoAddress) GoString() string {
	return s.String()
}

func (s *CollectRecruitJobDetailRequestJobInfoAddress) SetCityCode(v string) *CollectRecruitJobDetailRequestJobInfoAddress {
	s.CityCode = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfoAddress) SetDetail(v string) *CollectRecruitJobDetailRequestJobInfoAddress {
	s.Detail = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfoAddress) SetDistrictCode(v string) *CollectRecruitJobDetailRequestJobInfoAddress {
	s.DistrictCode = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfoAddress) SetLatitude(v string) *CollectRecruitJobDetailRequestJobInfoAddress {
	s.Latitude = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfoAddress) SetLongitude(v string) *CollectRecruitJobDetailRequestJobInfoAddress {
	s.Longitude = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfoAddress) SetName(v string) *CollectRecruitJobDetailRequestJobInfoAddress {
	s.Name = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfoAddress) SetProvinceCode(v string) *CollectRecruitJobDetailRequestJobInfoAddress {
	s.ProvinceCode = &v
	return s
}

type CollectRecruitJobDetailRequestJobInfoFullTimeInfo struct {
	MaxJobExperience *string `json:"maxJobExperience,omitempty" xml:"maxJobExperience,omitempty"`
	MinJobExperience *string `json:"minJobExperience,omitempty" xml:"minJobExperience,omitempty"`
	SalaryMonth      *string `json:"salaryMonth,omitempty" xml:"salaryMonth,omitempty"`
}

func (s CollectRecruitJobDetailRequestJobInfoFullTimeInfo) String() string {
	return tea.Prettify(s)
}

func (s CollectRecruitJobDetailRequestJobInfoFullTimeInfo) GoString() string {
	return s.String()
}

func (s *CollectRecruitJobDetailRequestJobInfoFullTimeInfo) SetMaxJobExperience(v string) *CollectRecruitJobDetailRequestJobInfoFullTimeInfo {
	s.MaxJobExperience = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfoFullTimeInfo) SetMinJobExperience(v string) *CollectRecruitJobDetailRequestJobInfoFullTimeInfo {
	s.MinJobExperience = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfoFullTimeInfo) SetSalaryMonth(v string) *CollectRecruitJobDetailRequestJobInfoFullTimeInfo {
	s.SalaryMonth = &v
	return s
}

type CollectRecruitJobDetailRequestJobInfoPartTimeInfo struct {
	ContactNumber    *string `json:"contactNumber,omitempty" xml:"contactNumber,omitempty"`
	SalaryPeriod     *string `json:"salaryPeriod,omitempty" xml:"salaryPeriod,omitempty"`
	SettleType       *string `json:"settleType,omitempty" xml:"settleType,omitempty"`
	SpecifyWorkDate  *string `json:"specifyWorkDate,omitempty" xml:"specifyWorkDate,omitempty"`
	SpecifyWorkTime  *string `json:"specifyWorkTime,omitempty" xml:"specifyWorkTime,omitempty"`
	WorkBeginTimeMin *string `json:"workBeginTimeMin,omitempty" xml:"workBeginTimeMin,omitempty"`
	WorkDateType     *string `json:"workDateType,omitempty" xml:"workDateType,omitempty"`
	WorkEndDate      *string `json:"workEndDate,omitempty" xml:"workEndDate,omitempty"`
	WorkEndTimeMin   *string `json:"workEndTimeMin,omitempty" xml:"workEndTimeMin,omitempty"`
	WorkStartDate    *string `json:"workStartDate,omitempty" xml:"workStartDate,omitempty"`
}

func (s CollectRecruitJobDetailRequestJobInfoPartTimeInfo) String() string {
	return tea.Prettify(s)
}

func (s CollectRecruitJobDetailRequestJobInfoPartTimeInfo) GoString() string {
	return s.String()
}

func (s *CollectRecruitJobDetailRequestJobInfoPartTimeInfo) SetContactNumber(v string) *CollectRecruitJobDetailRequestJobInfoPartTimeInfo {
	s.ContactNumber = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfoPartTimeInfo) SetSalaryPeriod(v string) *CollectRecruitJobDetailRequestJobInfoPartTimeInfo {
	s.SalaryPeriod = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfoPartTimeInfo) SetSettleType(v string) *CollectRecruitJobDetailRequestJobInfoPartTimeInfo {
	s.SettleType = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfoPartTimeInfo) SetSpecifyWorkDate(v string) *CollectRecruitJobDetailRequestJobInfoPartTimeInfo {
	s.SpecifyWorkDate = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfoPartTimeInfo) SetSpecifyWorkTime(v string) *CollectRecruitJobDetailRequestJobInfoPartTimeInfo {
	s.SpecifyWorkTime = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfoPartTimeInfo) SetWorkBeginTimeMin(v string) *CollectRecruitJobDetailRequestJobInfoPartTimeInfo {
	s.WorkBeginTimeMin = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfoPartTimeInfo) SetWorkDateType(v string) *CollectRecruitJobDetailRequestJobInfoPartTimeInfo {
	s.WorkDateType = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfoPartTimeInfo) SetWorkEndDate(v string) *CollectRecruitJobDetailRequestJobInfoPartTimeInfo {
	s.WorkEndDate = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfoPartTimeInfo) SetWorkEndTimeMin(v string) *CollectRecruitJobDetailRequestJobInfoPartTimeInfo {
	s.WorkEndTimeMin = &v
	return s
}

func (s *CollectRecruitJobDetailRequestJobInfoPartTimeInfo) SetWorkStartDate(v string) *CollectRecruitJobDetailRequestJobInfoPartTimeInfo {
	s.WorkStartDate = &v
	return s
}

type CollectRecruitJobDetailRequestRecruitUserInfo struct {
	ExtInfo    *string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	OutUserId  *string `json:"outUserId,omitempty" xml:"outUserId,omitempty"`
	UserMobile *string `json:"userMobile,omitempty" xml:"userMobile,omitempty"`
	UserName   *string `json:"userName,omitempty" xml:"userName,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CollectRecruitJobDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CollectRecruitJobDetailResponse) SetStatusCode(v int32) *CollectRecruitJobDetailResponse {
	s.StatusCode = &v
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
	BizCode          *string                               `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	ChannelCode      *string                               `json:"channelCode,omitempty" xml:"channelCode,omitempty"`
	ChannelOuterId   *string                               `json:"channelOuterId,omitempty" xml:"channelOuterId,omitempty"`
	ChannelTalentId  *string                               `json:"channelTalentId,omitempty" xml:"channelTalentId,omitempty"`
	DeliverJobId     *string                               `json:"deliverJobId,omitempty" xml:"deliverJobId,omitempty"`
	OptUserId        *string                               `json:"optUserId,omitempty" xml:"optUserId,omitempty"`
	ResumeChannelUrl *string                               `json:"resumeChannelUrl,omitempty" xml:"resumeChannelUrl,omitempty"`
	ResumeData       *CollectResumeDetailRequestResumeData `json:"resumeData,omitempty" xml:"resumeData,omitempty" type:"Struct"`
	ResumeFile       *CollectResumeDetailRequestResumeFile `json:"resumeFile,omitempty" xml:"resumeFile,omitempty" type:"Struct"`
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

func (s *CollectResumeDetailRequest) SetChannelCode(v string) *CollectResumeDetailRequest {
	s.ChannelCode = &v
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

func (s *CollectResumeDetailRequest) SetResumeChannelUrl(v string) *CollectResumeDetailRequest {
	s.ResumeChannelUrl = &v
	return s
}

func (s *CollectResumeDetailRequest) SetResumeData(v *CollectResumeDetailRequestResumeData) *CollectResumeDetailRequest {
	s.ResumeData = v
	return s
}

func (s *CollectResumeDetailRequest) SetResumeFile(v *CollectResumeDetailRequestResumeFile) *CollectResumeDetailRequest {
	s.ResumeFile = v
	return s
}

type CollectResumeDetailRequestResumeData struct {
	BaseInfo             *CollectResumeDetailRequestResumeDataBaseInfo               `json:"baseInfo,omitempty" xml:"baseInfo,omitempty" type:"Struct"`
	Certificates         []*CollectResumeDetailRequestResumeDataCertificates         `json:"certificates,omitempty" xml:"certificates,omitempty" type:"Repeated"`
	EducationExperiences []*CollectResumeDetailRequestResumeDataEducationExperiences `json:"educationExperiences,omitempty" xml:"educationExperiences,omitempty" type:"Repeated"`
	JobExpect            *CollectResumeDetailRequestResumeDataJobExpect              `json:"jobExpect,omitempty" xml:"jobExpect,omitempty" type:"Struct"`
	LanguageSkill        []*CollectResumeDetailRequestResumeDataLanguageSkill        `json:"languageSkill,omitempty" xml:"languageSkill,omitempty" type:"Repeated"`
	TrainingExperiences  []*CollectResumeDetailRequestResumeDataTrainingExperiences  `json:"trainingExperiences,omitempty" xml:"trainingExperiences,omitempty" type:"Repeated"`
	WorkExperiences      []*CollectResumeDetailRequestResumeDataWorkExperiences      `json:"workExperiences,omitempty" xml:"workExperiences,omitempty" type:"Repeated"`
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
	Age              *int32  `json:"age,omitempty" xml:"age,omitempty"`
	Avatar           *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	BeginWorkTime    *string `json:"beginWorkTime,omitempty" xml:"beginWorkTime,omitempty"`
	Birthday         *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	Email            *string `json:"email,omitempty" xml:"email,omitempty"`
	EnglishName      *string `json:"englishName,omitempty" xml:"englishName,omitempty"`
	GraduateTime     *string `json:"graduateTime,omitempty" xml:"graduateTime,omitempty"`
	HighestEducation *int32  `json:"highestEducation,omitempty" xml:"highestEducation,omitempty"`
	JobTitle         *string `json:"jobTitle,omitempty" xml:"jobTitle,omitempty"`
	LastSchoolName   *string `json:"lastSchoolName,omitempty" xml:"lastSchoolName,omitempty"`
	Married          *int32  `json:"married,omitempty" xml:"married,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	NativePlace      *string `json:"nativePlace,omitempty" xml:"nativePlace,omitempty"`
	NowLocation      *string `json:"nowLocation,omitempty" xml:"nowLocation,omitempty"`
	PersonalHonor    *string `json:"personalHonor,omitempty" xml:"personalHonor,omitempty"`
	PhoneNum         *string `json:"phoneNum,omitempty" xml:"phoneNum,omitempty"`
	PoliticalStatus  *int32  `json:"politicalStatus,omitempty" xml:"politicalStatus,omitempty"`
	SelfEvaluation   *string `json:"selfEvaluation,omitempty" xml:"selfEvaluation,omitempty"`
	Sex              *int32  `json:"sex,omitempty" xml:"sex,omitempty"`
	VirtualPhoneNum  *string `json:"virtualPhoneNum,omitempty" xml:"virtualPhoneNum,omitempty"`
	WorkingYears     *int32  `json:"workingYears,omitempty" xml:"workingYears,omitempty"`
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
	CertificateName *string `json:"certificateName,omitempty" xml:"certificateName,omitempty"`
	GrantTime       *string `json:"grantTime,omitempty" xml:"grantTime,omitempty"`
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
	Degree      *int32  `json:"degree,omitempty" xml:"degree,omitempty"`
	Department  *string `json:"department,omitempty" xml:"department,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	EndDate     *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	Major       *string `json:"major,omitempty" xml:"major,omitempty"`
	SchoolName  *string `json:"schoolName,omitempty" xml:"schoolName,omitempty"`
	StartDate   *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
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
	JobName     *string   `json:"jobName,omitempty" xml:"jobName,omitempty"`
	Locations   []*string `json:"locations,omitempty" xml:"locations,omitempty" type:"Repeated"`
	MaxSalary   *string   `json:"maxSalary,omitempty" xml:"maxSalary,omitempty"`
	MinSalary   *string   `json:"minSalary,omitempty" xml:"minSalary,omitempty"`
	OnboardTime *string   `json:"onboardTime,omitempty" xml:"onboardTime,omitempty"`
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
	CertificateName *string `json:"certificateName,omitempty" xml:"certificateName,omitempty"`
	LanguageName    *string `json:"languageName,omitempty" xml:"languageName,omitempty"`
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
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	EndDate         *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	InstitutionName *string `json:"institutionName,omitempty" xml:"institutionName,omitempty"`
	Location        *string `json:"location,omitempty" xml:"location,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	StartDate       *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
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
	CompanyName    *string `json:"companyName,omitempty" xml:"companyName,omitempty"`
	Department     *string `json:"department,omitempty" xml:"department,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	EndDate        *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	JobTitle       *string `json:"jobTitle,omitempty" xml:"jobTitle,omitempty"`
	Location       *string `json:"location,omitempty" xml:"location,omitempty"`
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

type CollectResumeDetailRequestResumeFile struct {
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	FileName    *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType    *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
}

func (s CollectResumeDetailRequestResumeFile) String() string {
	return tea.Prettify(s)
}

func (s CollectResumeDetailRequestResumeFile) GoString() string {
	return s.String()
}

func (s *CollectResumeDetailRequestResumeFile) SetDownloadUrl(v string) *CollectResumeDetailRequestResumeFile {
	s.DownloadUrl = &v
	return s
}

func (s *CollectResumeDetailRequestResumeFile) SetFileName(v string) *CollectResumeDetailRequestResumeFile {
	s.FileName = &v
	return s
}

func (s *CollectResumeDetailRequestResumeFile) SetFileType(v string) *CollectResumeDetailRequestResumeFile {
	s.FileType = &v
	return s
}

type CollectResumeDetailResponseBody struct {
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CollectResumeDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CollectResumeDetailResponse) SetStatusCode(v int32) *CollectResumeDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *CollectResumeDetailResponse) SetBody(v *CollectResumeDetailResponseBody) *CollectResumeDetailResponse {
	s.Body = v
	return s
}

type CollectResumeMailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CollectResumeMailHeaders) String() string {
	return tea.Prettify(s)
}

func (s CollectResumeMailHeaders) GoString() string {
	return s.String()
}

func (s *CollectResumeMailHeaders) SetCommonHeaders(v map[string]*string) *CollectResumeMailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CollectResumeMailHeaders) SetXAcsDingtalkAccessToken(v string) *CollectResumeMailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CollectResumeMailRequest struct {
	BizCode            *string                             `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	ChannelCode        *string                             `json:"channelCode,omitempty" xml:"channelCode,omitempty"`
	DeliverJobId       *string                             `json:"deliverJobId,omitempty" xml:"deliverJobId,omitempty"`
	FromMailAddress    *string                             `json:"fromMailAddress,omitempty" xml:"fromMailAddress,omitempty"`
	HistoryMailImport  *bool                               `json:"historyMailImport,omitempty" xml:"historyMailImport,omitempty"`
	MailId             *string                             `json:"mailId,omitempty" xml:"mailId,omitempty"`
	MailTitle          *string                             `json:"mailTitle,omitempty" xml:"mailTitle,omitempty"`
	OptUserId          *string                             `json:"optUserId,omitempty" xml:"optUserId,omitempty"`
	ReceiveMailAddress *string                             `json:"receiveMailAddress,omitempty" xml:"receiveMailAddress,omitempty"`
	ReceiveMailType    *int32                              `json:"receiveMailType,omitempty" xml:"receiveMailType,omitempty"`
	ReceivedTime       *int64                              `json:"receivedTime,omitempty" xml:"receivedTime,omitempty"`
	ResumeChannelUrl   *string                             `json:"resumeChannelUrl,omitempty" xml:"resumeChannelUrl,omitempty"`
	ResumeFile         *CollectResumeMailRequestResumeFile `json:"resumeFile,omitempty" xml:"resumeFile,omitempty" type:"Struct"`
}

func (s CollectResumeMailRequest) String() string {
	return tea.Prettify(s)
}

func (s CollectResumeMailRequest) GoString() string {
	return s.String()
}

func (s *CollectResumeMailRequest) SetBizCode(v string) *CollectResumeMailRequest {
	s.BizCode = &v
	return s
}

func (s *CollectResumeMailRequest) SetChannelCode(v string) *CollectResumeMailRequest {
	s.ChannelCode = &v
	return s
}

func (s *CollectResumeMailRequest) SetDeliverJobId(v string) *CollectResumeMailRequest {
	s.DeliverJobId = &v
	return s
}

func (s *CollectResumeMailRequest) SetFromMailAddress(v string) *CollectResumeMailRequest {
	s.FromMailAddress = &v
	return s
}

func (s *CollectResumeMailRequest) SetHistoryMailImport(v bool) *CollectResumeMailRequest {
	s.HistoryMailImport = &v
	return s
}

func (s *CollectResumeMailRequest) SetMailId(v string) *CollectResumeMailRequest {
	s.MailId = &v
	return s
}

func (s *CollectResumeMailRequest) SetMailTitle(v string) *CollectResumeMailRequest {
	s.MailTitle = &v
	return s
}

func (s *CollectResumeMailRequest) SetOptUserId(v string) *CollectResumeMailRequest {
	s.OptUserId = &v
	return s
}

func (s *CollectResumeMailRequest) SetReceiveMailAddress(v string) *CollectResumeMailRequest {
	s.ReceiveMailAddress = &v
	return s
}

func (s *CollectResumeMailRequest) SetReceiveMailType(v int32) *CollectResumeMailRequest {
	s.ReceiveMailType = &v
	return s
}

func (s *CollectResumeMailRequest) SetReceivedTime(v int64) *CollectResumeMailRequest {
	s.ReceivedTime = &v
	return s
}

func (s *CollectResumeMailRequest) SetResumeChannelUrl(v string) *CollectResumeMailRequest {
	s.ResumeChannelUrl = &v
	return s
}

func (s *CollectResumeMailRequest) SetResumeFile(v *CollectResumeMailRequestResumeFile) *CollectResumeMailRequest {
	s.ResumeFile = v
	return s
}

type CollectResumeMailRequestResumeFile struct {
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	FileName    *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType    *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
}

func (s CollectResumeMailRequestResumeFile) String() string {
	return tea.Prettify(s)
}

func (s CollectResumeMailRequestResumeFile) GoString() string {
	return s.String()
}

func (s *CollectResumeMailRequestResumeFile) SetDownloadUrl(v string) *CollectResumeMailRequestResumeFile {
	s.DownloadUrl = &v
	return s
}

func (s *CollectResumeMailRequestResumeFile) SetFileName(v string) *CollectResumeMailRequestResumeFile {
	s.FileName = &v
	return s
}

func (s *CollectResumeMailRequestResumeFile) SetFileType(v string) *CollectResumeMailRequestResumeFile {
	s.FileType = &v
	return s
}

type CollectResumeMailResponseBody struct {
	ResumeId *string `json:"resumeId,omitempty" xml:"resumeId,omitempty"`
}

func (s CollectResumeMailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollectResumeMailResponseBody) GoString() string {
	return s.String()
}

func (s *CollectResumeMailResponseBody) SetResumeId(v string) *CollectResumeMailResponseBody {
	s.ResumeId = &v
	return s
}

type CollectResumeMailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CollectResumeMailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CollectResumeMailResponse) String() string {
	return tea.Prettify(s)
}

func (s CollectResumeMailResponse) GoString() string {
	return s.String()
}

func (s *CollectResumeMailResponse) SetHeaders(v map[string]*string) *CollectResumeMailResponse {
	s.Headers = v
	return s
}

func (s *CollectResumeMailResponse) SetStatusCode(v int32) *CollectResumeMailResponse {
	s.StatusCode = &v
	return s
}

func (s *CollectResumeMailResponse) SetBody(v *CollectResumeMailResponseBody) *CollectResumeMailResponse {
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ConfirmRightsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ConfirmRightsResponse) SetStatusCode(v int32) *ConfirmRightsResponse {
	s.StatusCode = &v
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
	Scope  *string `json:"scope,omitempty" xml:"scope,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FinishBeginnerTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *FinishBeginnerTaskResponse) SetStatusCode(v int32) *FinishBeginnerTaskResponse {
	s.StatusCode = &v
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
	CandidateId       *string `json:"candidateId,omitempty" xml:"candidateId,omitempty"`
	CreatorUserId     *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	FlowId            *string `json:"flowId,omitempty" xml:"flowId,omitempty"`
	FormId            *string `json:"formId,omitempty" xml:"formId,omitempty"`
	GmtCreateMillis   *int64  `json:"gmtCreateMillis,omitempty" xml:"gmtCreateMillis,omitempty"`
	GmtModifiedMillis *int64  `json:"gmtModifiedMillis,omitempty" xml:"gmtModifiedMillis,omitempty"`
	JobId             *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	Status            *int32  `json:"status,omitempty" xml:"status,omitempty"`
	TemplateId        *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	TemplateVersion   *int32  `json:"templateVersion,omitempty" xml:"templateVersion,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetApplicationRegFormByFlowIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetApplicationRegFormByFlowIdResponse) SetStatusCode(v int32) *GetApplicationRegFormByFlowIdResponse {
	s.StatusCode = &v
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
	BizCode     *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
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
	CandidateId *string `json:"candidateId,omitempty" xml:"candidateId,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCandidateByPhoneNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetCandidateByPhoneNumberResponse) SetStatusCode(v int32) *GetCandidateByPhoneNumberResponse {
	s.StatusCode = &v
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
	BizCode  *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileSize *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	Md5      *string `json:"md5,omitempty" xml:"md5,omitempty"`
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
	AccessKeyId                 *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	AccessKeySecret             *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	AccessToken                 *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	AccessTokenExpirationMillis *int64  `json:"accessTokenExpirationMillis,omitempty" xml:"accessTokenExpirationMillis,omitempty"`
	Bucket                      *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	EndPoint                    *string `json:"endPoint,omitempty" xml:"endPoint,omitempty"`
	MediaId                     *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFileUploadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetFileUploadInfoResponse) SetStatusCode(v int32) *GetFileUploadInfoResponse {
	s.StatusCode = &v
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
	BizCode          *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	RelationEntity   *string `json:"relationEntity,omitempty" xml:"relationEntity,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFlowIdByRelationEntityIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetFlowIdByRelationEntityIdResponse) SetStatusCode(v int32) *GetFlowIdByRelationEntityIdResponse {
	s.StatusCode = &v
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
	JobId     *string                            `json:"jobId,omitempty" xml:"jobId,omitempty"`
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
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetJobAuthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetJobAuthResponse) SetStatusCode(v int32) *GetJobAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobAuthResponse) SetBody(v *GetJobAuthResponseBody) *GetJobAuthResponse {
	s.Body = v
	return s
}

type QueryCandidatesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCandidatesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCandidatesHeaders) GoString() string {
	return s.String()
}

func (s *QueryCandidatesHeaders) SetCommonHeaders(v map[string]*string) *QueryCandidatesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCandidatesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCandidatesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCandidatesRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	StatId     *string `json:"statId,omitempty" xml:"statId,omitempty"`
	OpUserId   *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s QueryCandidatesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCandidatesRequest) GoString() string {
	return s.String()
}

func (s *QueryCandidatesRequest) SetMaxResults(v int32) *QueryCandidatesRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryCandidatesRequest) SetNextToken(v string) *QueryCandidatesRequest {
	s.NextToken = &v
	return s
}

func (s *QueryCandidatesRequest) SetStatId(v string) *QueryCandidatesRequest {
	s.StatId = &v
	return s
}

func (s *QueryCandidatesRequest) SetOpUserId(v string) *QueryCandidatesRequest {
	s.OpUserId = &v
	return s
}

type QueryCandidatesResponseBody struct {
	HasMore    *bool                              `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*QueryCandidatesResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextToken  *int64                             `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	TotalCount *int64                             `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryCandidatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCandidatesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCandidatesResponseBody) SetHasMore(v bool) *QueryCandidatesResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryCandidatesResponseBody) SetList(v []*QueryCandidatesResponseBodyList) *QueryCandidatesResponseBody {
	s.List = v
	return s
}

func (s *QueryCandidatesResponseBody) SetNextToken(v int64) *QueryCandidatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryCandidatesResponseBody) SetTotalCount(v int64) *QueryCandidatesResponseBody {
	s.TotalCount = &v
	return s
}

type QueryCandidatesResponseBodyList struct {
	CandidateId *string `json:"candidateId,omitempty" xml:"candidateId,omitempty"`
	CorpId      *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	EntryDate   *int32  `json:"entryDate,omitempty" xml:"entryDate,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryCandidatesResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryCandidatesResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryCandidatesResponseBodyList) SetCandidateId(v string) *QueryCandidatesResponseBodyList {
	s.CandidateId = &v
	return s
}

func (s *QueryCandidatesResponseBodyList) SetCorpId(v string) *QueryCandidatesResponseBodyList {
	s.CorpId = &v
	return s
}

func (s *QueryCandidatesResponseBodyList) SetEntryDate(v int32) *QueryCandidatesResponseBodyList {
	s.EntryDate = &v
	return s
}

func (s *QueryCandidatesResponseBodyList) SetUserId(v string) *QueryCandidatesResponseBodyList {
	s.UserId = &v
	return s
}

type QueryCandidatesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCandidatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCandidatesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCandidatesResponse) GoString() string {
	return s.String()
}

func (s *QueryCandidatesResponse) SetHeaders(v map[string]*string) *QueryCandidatesResponse {
	s.Headers = v
	return s
}

func (s *QueryCandidatesResponse) SetStatusCode(v int32) *QueryCandidatesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCandidatesResponse) SetBody(v *QueryCandidatesResponseBody) *QueryCandidatesResponse {
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
	BizCode              *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	CandidateId          *string `json:"candidateId,omitempty" xml:"candidateId,omitempty"`
	StartTimeBeginMillis *int64  `json:"startTimeBeginMillis,omitempty" xml:"startTimeBeginMillis,omitempty"`
	StartTimeEndMillis   *int64  `json:"startTimeEndMillis,omitempty" xml:"startTimeEndMillis,omitempty"`
	NextToken            *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Size                 *int64  `json:"size,omitempty" xml:"size,omitempty"`
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
	HasMore    *bool                              `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*QueryInterviewsResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextToken  *string                            `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	TotalCount *int64                             `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	Cancelled       *bool                                          `json:"cancelled,omitempty" xml:"cancelled,omitempty"`
	CreatorUserId   *string                                        `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	EndTimeMillis   *int64                                         `json:"endTimeMillis,omitempty" xml:"endTimeMillis,omitempty"`
	InterviewId     *string                                        `json:"interviewId,omitempty" xml:"interviewId,omitempty"`
	Interviewers    []*QueryInterviewsResponseBodyListInterviewers `json:"interviewers,omitempty" xml:"interviewers,omitempty" type:"Repeated"`
	JobId           *string                                        `json:"jobId,omitempty" xml:"jobId,omitempty"`
	StartTimeMillis *int64                                         `json:"startTimeMillis,omitempty" xml:"startTimeMillis,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryInterviewsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryInterviewsResponse) SetStatusCode(v int32) *QueryInterviewsResponse {
	s.StatusCode = &v
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
	BizCode        *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	Channel        *string `json:"channel,omitempty" xml:"channel,omitempty"`
	ErrorCode      *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg       *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReportMessageStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ReportMessageStatusResponse) SetStatusCode(v int32) *ReportMessageStatusResponse {
	s.StatusCode = &v
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
	BizCode        *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	Channel        *string `json:"channel,omitempty" xml:"channel,omitempty"`
	Content        *string `json:"content,omitempty" xml:"content,omitempty"`
	CreateTime     *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	ReceiverUserId *string `json:"receiverUserId,omitempty" xml:"receiverUserId,omitempty"`
	SenderUserId   *string `json:"senderUserId,omitempty" xml:"senderUserId,omitempty"`
	Uuid           *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SyncChannelMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SyncChannelMessageResponse) SetStatusCode(v int32) *SyncChannelMessageResponse {
	s.StatusCode = &v
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
	BizCode     *string                                     `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	Content     *string                                     `json:"content,omitempty" xml:"content,omitempty"`
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
	FileId   *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileSize *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	SpaceId  *int64  `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
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
	CreatorUserId     *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	FormId            *string `json:"formId,omitempty" xml:"formId,omitempty"`
	GmtCreateMillis   *int64  `json:"gmtCreateMillis,omitempty" xml:"gmtCreateMillis,omitempty"`
	GmtModifiedMillis *int64  `json:"gmtModifiedMillis,omitempty" xml:"gmtModifiedMillis,omitempty"`
	Status            *int32  `json:"status,omitempty" xml:"status,omitempty"`
	TemplateId        *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	TemplateVersion   *int32  `json:"templateVersion,omitempty" xml:"templateVersion,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateApplicationRegFormResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateApplicationRegFormResponse) SetStatusCode(v int32) *UpdateApplicationRegFormResponse {
	s.StatusCode = &v
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
	BizCode          *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	SignInTimeMillis *int64  `json:"signInTimeMillis,omitempty" xml:"signInTimeMillis,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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

func (s *UpdateInterviewSignInInfoResponse) SetStatusCode(v int32) *UpdateInterviewSignInInfoResponse {
	s.StatusCode = &v
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
	BizCode        *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	ChannelOuterId *string `json:"channelOuterId,omitempty" xml:"channelOuterId,omitempty"`
	DeliverUserId  *string `json:"deliverUserId,omitempty" xml:"deliverUserId,omitempty"`
	ErrorCode      *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg       *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	OpTime         *int64  `json:"opTime,omitempty" xml:"opTime,omitempty"`
	OpUserId       *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	JobId          *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateJobDeliverResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateJobDeliverResponse) SetStatusCode(v int32) *UpdateJobDeliverResponse {
	s.StatusCode = &v
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
	interfaceSPI, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = interfaceSPI
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
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
	params := &openapi.Params{
		Action:      tea.String("AddApplicationRegFormTemplate"),
		Version:     tea.String("ats_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/ats/flows/applicationRegForms/templates"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddApplicationRegFormTemplateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
	params := &openapi.Params{
		Action:      tea.String("AddFile"),
		Version:     tea.String("ats_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/ats/files"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("AddUserAccount"),
		Version:     tea.String("ats_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/ats/channels/users/accounts"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddUserAccountResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.JobInfo)) {
		body["jobInfo"] = request.JobInfo
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
	params := &openapi.Params{
		Action:      tea.String("CollectRecruitJobDetail"),
		Version:     tea.String("ats_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/ats/channels/jobs/import"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CollectRecruitJobDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.ChannelCode)) {
		body["channelCode"] = request.ChannelCode
	}

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

	if !tea.BoolValue(util.IsUnset(request.ResumeChannelUrl)) {
		body["resumeChannelUrl"] = request.ResumeChannelUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ResumeData)) {
		body["resumeData"] = request.ResumeData
	}

	if !tea.BoolValue(util.IsUnset(request.ResumeFile)) {
		body["resumeFile"] = request.ResumeFile
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
	params := &openapi.Params{
		Action:      tea.String("CollectResumeDetail"),
		Version:     tea.String("ats_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/ats/resumes/details"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CollectResumeDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) CollectResumeMailWithOptions(request *CollectResumeMailRequest, headers *CollectResumeMailHeaders, runtime *util.RuntimeOptions) (_result *CollectResumeMailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["bizCode"] = request.BizCode
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelCode)) {
		body["channelCode"] = request.ChannelCode
	}

	if !tea.BoolValue(util.IsUnset(request.DeliverJobId)) {
		body["deliverJobId"] = request.DeliverJobId
	}

	if !tea.BoolValue(util.IsUnset(request.FromMailAddress)) {
		body["fromMailAddress"] = request.FromMailAddress
	}

	if !tea.BoolValue(util.IsUnset(request.HistoryMailImport)) {
		body["historyMailImport"] = request.HistoryMailImport
	}

	if !tea.BoolValue(util.IsUnset(request.MailId)) {
		body["mailId"] = request.MailId
	}

	if !tea.BoolValue(util.IsUnset(request.MailTitle)) {
		body["mailTitle"] = request.MailTitle
	}

	if !tea.BoolValue(util.IsUnset(request.OptUserId)) {
		body["optUserId"] = request.OptUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiveMailAddress)) {
		body["receiveMailAddress"] = request.ReceiveMailAddress
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiveMailType)) {
		body["receiveMailType"] = request.ReceiveMailType
	}

	if !tea.BoolValue(util.IsUnset(request.ReceivedTime)) {
		body["receivedTime"] = request.ReceivedTime
	}

	if !tea.BoolValue(util.IsUnset(request.ResumeChannelUrl)) {
		body["resumeChannelUrl"] = request.ResumeChannelUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ResumeFile)) {
		body["resumeFile"] = request.ResumeFile
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
	params := &openapi.Params{
		Action:      tea.String("CollectResumeMail"),
		Version:     tea.String("ats_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/ats/resumes/mails"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CollectResumeMailResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CollectResumeMail(request *CollectResumeMailRequest) (_result *CollectResumeMailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CollectResumeMailHeaders{}
	_result = &CollectResumeMailResponse{}
	_body, _err := client.CollectResumeMailWithOptions(request, headers, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("ConfirmRights"),
		Version:     tea.String("ats_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/ats/rights/" + tea.StringValue(rightsCode) + "/confirm"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfirmRightsResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) FinishBeginnerTaskWithOptions(taskCode *string, request *FinishBeginnerTaskRequest, headers *FinishBeginnerTaskHeaders, runtime *util.RuntimeOptions) (_result *FinishBeginnerTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
	params := &openapi.Params{
		Action:      tea.String("FinishBeginnerTask"),
		Version:     tea.String("ats_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/ats/beginnerTasks/" + tea.StringValue(taskCode) + "/finish"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FinishBeginnerTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) GetApplicationRegFormByFlowIdWithOptions(flowId *string, request *GetApplicationRegFormByFlowIdRequest, headers *GetApplicationRegFormByFlowIdHeaders, runtime *util.RuntimeOptions) (_result *GetApplicationRegFormByFlowIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
	params := &openapi.Params{
		Action:      tea.String("GetApplicationRegFormByFlowId"),
		Version:     tea.String("ats_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/ats/flows/" + tea.StringValue(flowId) + "/applicationRegForms"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetApplicationRegFormByFlowIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetCandidateByPhoneNumber"),
		Version:     tea.String("ats_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/ats/candidates"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCandidateByPhoneNumberResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetFileUploadInfo"),
		Version:     tea.String("ats_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/ats/files/uploadInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFileUploadInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetFlowIdByRelationEntityId"),
		Version:     tea.String("ats_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/ats/flows/ids"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFlowIdByRelationEntityIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) GetJobAuthWithOptions(jobId *string, request *GetJobAuthRequest, headers *GetJobAuthHeaders, runtime *util.RuntimeOptions) (_result *GetJobAuthResponse, _err error) {
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobAuth"),
		Version:     tea.String("ats_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/ats/auths/jobs/" + tea.StringValue(jobId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobAuthResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) QueryCandidatesWithOptions(request *QueryCandidatesRequest, headers *QueryCandidatesHeaders, runtime *util.RuntimeOptions) (_result *QueryCandidatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.StatId)) {
		body["statId"] = request.StatId
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
	params := &openapi.Params{
		Action:      tea.String("QueryCandidates"),
		Version:     tea.String("ats_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/ats/candidates/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCandidatesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCandidates(request *QueryCandidatesRequest) (_result *QueryCandidatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCandidatesHeaders{}
	_result = &QueryCandidatesResponse{}
	_body, _err := client.QueryCandidatesWithOptions(request, headers, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("QueryInterviews"),
		Version:     tea.String("ats_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/ats/interviews/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryInterviewsResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("ReportMessageStatus"),
		Version:     tea.String("ats_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/ats/channels/messages/statuses/report"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportMessageStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("SyncChannelMessage"),
		Version:     tea.String("ats_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/ats/channels/messages/sync"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncChannelMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) UpdateApplicationRegFormWithOptions(flowId *string, request *UpdateApplicationRegFormRequest, headers *UpdateApplicationRegFormHeaders, runtime *util.RuntimeOptions) (_result *UpdateApplicationRegFormResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
	params := &openapi.Params{
		Action:      tea.String("UpdateApplicationRegForm"),
		Version:     tea.String("ats_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/ats/flows/" + tea.StringValue(flowId) + "/applicationRegForms"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateApplicationRegFormResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) UpdateInterviewSignInInfoWithOptions(interviewId *string, request *UpdateInterviewSignInInfoRequest, headers *UpdateInterviewSignInInfoHeaders, runtime *util.RuntimeOptions) (_result *UpdateInterviewSignInInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
	params := &openapi.Params{
		Action:      tea.String("UpdateInterviewSignInInfo"),
		Version:     tea.String("ats_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/ats/interviews/" + tea.StringValue(interviewId) + "/signInInfos"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInterviewSignInInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("UpdateJobDeliver"),
		Version:     tea.String("ats_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/ats/jobs/deliveryStatus"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateJobDeliverResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
