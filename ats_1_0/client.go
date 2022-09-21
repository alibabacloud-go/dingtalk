// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package ats_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
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

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.DingPanFile))) {
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
