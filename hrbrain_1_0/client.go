// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package hrbrain_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type HrbrainImportAwardDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainImportAwardDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportAwardDetailHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainImportAwardDetailHeaders) SetCommonHeaders(v map[string]*string) *HrbrainImportAwardDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainImportAwardDetailHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainImportAwardDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainImportAwardDetailRequest struct {
	Body   []*HrbrainImportAwardDetailRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	CorpId *string                                `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s HrbrainImportAwardDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportAwardDetailRequest) GoString() string {
	return s.String()
}

func (s *HrbrainImportAwardDetailRequest) SetBody(v []*HrbrainImportAwardDetailRequestBody) *HrbrainImportAwardDetailRequest {
	s.Body = v
	return s
}

func (s *HrbrainImportAwardDetailRequest) SetCorpId(v string) *HrbrainImportAwardDetailRequest {
	s.CorpId = &v
	return s
}

type HrbrainImportAwardDetailRequestBody struct {
	AwardDate  *string                `json:"awardDate,omitempty" xml:"awardDate,omitempty"`
	AwardName  *string                `json:"awardName,omitempty" xml:"awardName,omitempty"`
	AwardOrg   *string                `json:"awardOrg,omitempty" xml:"awardOrg,omitempty"`
	AwardType  *string                `json:"awardType,omitempty" xml:"awardType,omitempty"`
	Comment    *string                `json:"comment,omitempty" xml:"comment,omitempty"`
	ExtendInfo map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	Name       *string                `json:"name,omitempty" xml:"name,omitempty"`
	WorkNo     *string                `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainImportAwardDetailRequestBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportAwardDetailRequestBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportAwardDetailRequestBody) SetAwardDate(v string) *HrbrainImportAwardDetailRequestBody {
	s.AwardDate = &v
	return s
}

func (s *HrbrainImportAwardDetailRequestBody) SetAwardName(v string) *HrbrainImportAwardDetailRequestBody {
	s.AwardName = &v
	return s
}

func (s *HrbrainImportAwardDetailRequestBody) SetAwardOrg(v string) *HrbrainImportAwardDetailRequestBody {
	s.AwardOrg = &v
	return s
}

func (s *HrbrainImportAwardDetailRequestBody) SetAwardType(v string) *HrbrainImportAwardDetailRequestBody {
	s.AwardType = &v
	return s
}

func (s *HrbrainImportAwardDetailRequestBody) SetComment(v string) *HrbrainImportAwardDetailRequestBody {
	s.Comment = &v
	return s
}

func (s *HrbrainImportAwardDetailRequestBody) SetExtendInfo(v map[string]interface{}) *HrbrainImportAwardDetailRequestBody {
	s.ExtendInfo = v
	return s
}

func (s *HrbrainImportAwardDetailRequestBody) SetName(v string) *HrbrainImportAwardDetailRequestBody {
	s.Name = &v
	return s
}

func (s *HrbrainImportAwardDetailRequestBody) SetWorkNo(v string) *HrbrainImportAwardDetailRequestBody {
	s.WorkNo = &v
	return s
}

type HrbrainImportAwardDetailResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainImportAwardDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportAwardDetailResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportAwardDetailResponseBody) SetRequestId(v string) *HrbrainImportAwardDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainImportAwardDetailResponseBody) SetResult(v bool) *HrbrainImportAwardDetailResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainImportAwardDetailResponseBody) SetSuccess(v bool) *HrbrainImportAwardDetailResponseBody {
	s.Success = &v
	return s
}

type HrbrainImportAwardDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainImportAwardDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainImportAwardDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportAwardDetailResponse) GoString() string {
	return s.String()
}

func (s *HrbrainImportAwardDetailResponse) SetHeaders(v map[string]*string) *HrbrainImportAwardDetailResponse {
	s.Headers = v
	return s
}

func (s *HrbrainImportAwardDetailResponse) SetStatusCode(v int32) *HrbrainImportAwardDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainImportAwardDetailResponse) SetBody(v *HrbrainImportAwardDetailResponseBody) *HrbrainImportAwardDetailResponse {
	s.Body = v
	return s
}

type HrbrainImportDeptInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainImportDeptInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportDeptInfoHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainImportDeptInfoHeaders) SetCommonHeaders(v map[string]*string) *HrbrainImportDeptInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainImportDeptInfoHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainImportDeptInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainImportDeptInfoRequest struct {
	Body   []*HrbrainImportDeptInfoRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	CorpId *string                             `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s HrbrainImportDeptInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportDeptInfoRequest) GoString() string {
	return s.String()
}

func (s *HrbrainImportDeptInfoRequest) SetBody(v []*HrbrainImportDeptInfoRequestBody) *HrbrainImportDeptInfoRequest {
	s.Body = v
	return s
}

func (s *HrbrainImportDeptInfoRequest) SetCorpId(v string) *HrbrainImportDeptInfoRequest {
	s.CorpId = &v
	return s
}

type HrbrainImportDeptInfoRequestBody struct {
	CreateDate    *string                `json:"createDate,omitempty" xml:"createDate,omitempty"`
	DeptName      *string                `json:"deptName,omitempty" xml:"deptName,omitempty"`
	DeptNo        *string                `json:"deptNo,omitempty" xml:"deptNo,omitempty"`
	EffectiveDate *string                `json:"effectiveDate,omitempty" xml:"effectiveDate,omitempty"`
	ExtendInfo    map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	IsEffective   *string                `json:"isEffective,omitempty" xml:"isEffective,omitempty"`
	SuperDeptName *string                `json:"superDeptName,omitempty" xml:"superDeptName,omitempty"`
	SuperDeptNo   *string                `json:"superDeptNo,omitempty" xml:"superDeptNo,omitempty"`
	SuperEmpId    *string                `json:"superEmpId,omitempty" xml:"superEmpId,omitempty"`
	SuperName     *string                `json:"superName,omitempty" xml:"superName,omitempty"`
}

func (s HrbrainImportDeptInfoRequestBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportDeptInfoRequestBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportDeptInfoRequestBody) SetCreateDate(v string) *HrbrainImportDeptInfoRequestBody {
	s.CreateDate = &v
	return s
}

func (s *HrbrainImportDeptInfoRequestBody) SetDeptName(v string) *HrbrainImportDeptInfoRequestBody {
	s.DeptName = &v
	return s
}

func (s *HrbrainImportDeptInfoRequestBody) SetDeptNo(v string) *HrbrainImportDeptInfoRequestBody {
	s.DeptNo = &v
	return s
}

func (s *HrbrainImportDeptInfoRequestBody) SetEffectiveDate(v string) *HrbrainImportDeptInfoRequestBody {
	s.EffectiveDate = &v
	return s
}

func (s *HrbrainImportDeptInfoRequestBody) SetExtendInfo(v map[string]interface{}) *HrbrainImportDeptInfoRequestBody {
	s.ExtendInfo = v
	return s
}

func (s *HrbrainImportDeptInfoRequestBody) SetIsEffective(v string) *HrbrainImportDeptInfoRequestBody {
	s.IsEffective = &v
	return s
}

func (s *HrbrainImportDeptInfoRequestBody) SetSuperDeptName(v string) *HrbrainImportDeptInfoRequestBody {
	s.SuperDeptName = &v
	return s
}

func (s *HrbrainImportDeptInfoRequestBody) SetSuperDeptNo(v string) *HrbrainImportDeptInfoRequestBody {
	s.SuperDeptNo = &v
	return s
}

func (s *HrbrainImportDeptInfoRequestBody) SetSuperEmpId(v string) *HrbrainImportDeptInfoRequestBody {
	s.SuperEmpId = &v
	return s
}

func (s *HrbrainImportDeptInfoRequestBody) SetSuperName(v string) *HrbrainImportDeptInfoRequestBody {
	s.SuperName = &v
	return s
}

type HrbrainImportDeptInfoResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainImportDeptInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportDeptInfoResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportDeptInfoResponseBody) SetRequestId(v string) *HrbrainImportDeptInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainImportDeptInfoResponseBody) SetResult(v bool) *HrbrainImportDeptInfoResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainImportDeptInfoResponseBody) SetSuccess(v bool) *HrbrainImportDeptInfoResponseBody {
	s.Success = &v
	return s
}

type HrbrainImportDeptInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainImportDeptInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainImportDeptInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportDeptInfoResponse) GoString() string {
	return s.String()
}

func (s *HrbrainImportDeptInfoResponse) SetHeaders(v map[string]*string) *HrbrainImportDeptInfoResponse {
	s.Headers = v
	return s
}

func (s *HrbrainImportDeptInfoResponse) SetStatusCode(v int32) *HrbrainImportDeptInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainImportDeptInfoResponse) SetBody(v *HrbrainImportDeptInfoResponseBody) *HrbrainImportDeptInfoResponse {
	s.Body = v
	return s
}

type HrbrainImportDimissionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainImportDimissionHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportDimissionHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainImportDimissionHeaders) SetCommonHeaders(v map[string]*string) *HrbrainImportDimissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainImportDimissionHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainImportDimissionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainImportDimissionRequest struct {
	Body   []*HrbrainImportDimissionRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	CorpId *string                              `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s HrbrainImportDimissionRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportDimissionRequest) GoString() string {
	return s.String()
}

func (s *HrbrainImportDimissionRequest) SetBody(v []*HrbrainImportDimissionRequestBody) *HrbrainImportDimissionRequest {
	s.Body = v
	return s
}

func (s *HrbrainImportDimissionRequest) SetCorpId(v string) *HrbrainImportDimissionRequest {
	s.CorpId = &v
	return s
}

type HrbrainImportDimissionRequestBody struct {
	DeptName             *string                `json:"deptName,omitempty" xml:"deptName,omitempty"`
	DeptNo               *string                `json:"deptNo,omitempty" xml:"deptNo,omitempty"`
	DimissionDate        *string                `json:"dimissionDate,omitempty" xml:"dimissionDate,omitempty"`
	DimissionReaasonDesc *string                `json:"dimissionReaasonDesc,omitempty" xml:"dimissionReaasonDesc,omitempty"`
	DimissionReason      *string                `json:"dimissionReason,omitempty" xml:"dimissionReason,omitempty"`
	EmpType              *string                `json:"empType,omitempty" xml:"empType,omitempty"`
	ExtendInfo           map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	JobCodeName          *string                `json:"jobCodeName,omitempty" xml:"jobCodeName,omitempty"`
	JobLevel             *string                `json:"jobLevel,omitempty" xml:"jobLevel,omitempty"`
	Name                 *string                `json:"name,omitempty" xml:"name,omitempty"`
	PostName             *string                `json:"postName,omitempty" xml:"postName,omitempty"`
	SuperName            *string                `json:"superName,omitempty" xml:"superName,omitempty"`
	WorkLocAddr          *string                `json:"workLocAddr,omitempty" xml:"workLocAddr,omitempty"`
	WorkNo               *string                `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainImportDimissionRequestBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportDimissionRequestBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportDimissionRequestBody) SetDeptName(v string) *HrbrainImportDimissionRequestBody {
	s.DeptName = &v
	return s
}

func (s *HrbrainImportDimissionRequestBody) SetDeptNo(v string) *HrbrainImportDimissionRequestBody {
	s.DeptNo = &v
	return s
}

func (s *HrbrainImportDimissionRequestBody) SetDimissionDate(v string) *HrbrainImportDimissionRequestBody {
	s.DimissionDate = &v
	return s
}

func (s *HrbrainImportDimissionRequestBody) SetDimissionReaasonDesc(v string) *HrbrainImportDimissionRequestBody {
	s.DimissionReaasonDesc = &v
	return s
}

func (s *HrbrainImportDimissionRequestBody) SetDimissionReason(v string) *HrbrainImportDimissionRequestBody {
	s.DimissionReason = &v
	return s
}

func (s *HrbrainImportDimissionRequestBody) SetEmpType(v string) *HrbrainImportDimissionRequestBody {
	s.EmpType = &v
	return s
}

func (s *HrbrainImportDimissionRequestBody) SetExtendInfo(v map[string]interface{}) *HrbrainImportDimissionRequestBody {
	s.ExtendInfo = v
	return s
}

func (s *HrbrainImportDimissionRequestBody) SetJobCodeName(v string) *HrbrainImportDimissionRequestBody {
	s.JobCodeName = &v
	return s
}

func (s *HrbrainImportDimissionRequestBody) SetJobLevel(v string) *HrbrainImportDimissionRequestBody {
	s.JobLevel = &v
	return s
}

func (s *HrbrainImportDimissionRequestBody) SetName(v string) *HrbrainImportDimissionRequestBody {
	s.Name = &v
	return s
}

func (s *HrbrainImportDimissionRequestBody) SetPostName(v string) *HrbrainImportDimissionRequestBody {
	s.PostName = &v
	return s
}

func (s *HrbrainImportDimissionRequestBody) SetSuperName(v string) *HrbrainImportDimissionRequestBody {
	s.SuperName = &v
	return s
}

func (s *HrbrainImportDimissionRequestBody) SetWorkLocAddr(v string) *HrbrainImportDimissionRequestBody {
	s.WorkLocAddr = &v
	return s
}

func (s *HrbrainImportDimissionRequestBody) SetWorkNo(v string) *HrbrainImportDimissionRequestBody {
	s.WorkNo = &v
	return s
}

type HrbrainImportDimissionResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainImportDimissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportDimissionResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportDimissionResponseBody) SetRequestId(v string) *HrbrainImportDimissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainImportDimissionResponseBody) SetResult(v bool) *HrbrainImportDimissionResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainImportDimissionResponseBody) SetSuccess(v bool) *HrbrainImportDimissionResponseBody {
	s.Success = &v
	return s
}

type HrbrainImportDimissionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainImportDimissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainImportDimissionResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportDimissionResponse) GoString() string {
	return s.String()
}

func (s *HrbrainImportDimissionResponse) SetHeaders(v map[string]*string) *HrbrainImportDimissionResponse {
	s.Headers = v
	return s
}

func (s *HrbrainImportDimissionResponse) SetStatusCode(v int32) *HrbrainImportDimissionResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainImportDimissionResponse) SetBody(v *HrbrainImportDimissionResponseBody) *HrbrainImportDimissionResponse {
	s.Body = v
	return s
}

type HrbrainImportEduExpHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainImportEduExpHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportEduExpHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainImportEduExpHeaders) SetCommonHeaders(v map[string]*string) *HrbrainImportEduExpHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainImportEduExpHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainImportEduExpHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainImportEduExpRequest struct {
	Body   []*HrbrainImportEduExpRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	CorpId *string                           `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s HrbrainImportEduExpRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportEduExpRequest) GoString() string {
	return s.String()
}

func (s *HrbrainImportEduExpRequest) SetBody(v []*HrbrainImportEduExpRequestBody) *HrbrainImportEduExpRequest {
	s.Body = v
	return s
}

func (s *HrbrainImportEduExpRequest) SetCorpId(v string) *HrbrainImportEduExpRequest {
	s.CorpId = &v
	return s
}

type HrbrainImportEduExpRequestBody struct {
	EduName    *string                `json:"eduName,omitempty" xml:"eduName,omitempty"`
	EndDate    *string                `json:"endDate,omitempty" xml:"endDate,omitempty"`
	ExtendInfo map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	Major      *string                `json:"major,omitempty" xml:"major,omitempty"`
	Name       *string                `json:"name,omitempty" xml:"name,omitempty"`
	SchoolName *string                `json:"schoolName,omitempty" xml:"schoolName,omitempty"`
	StartDate  *string                `json:"startDate,omitempty" xml:"startDate,omitempty"`
	WorkNo     *string                `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainImportEduExpRequestBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportEduExpRequestBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportEduExpRequestBody) SetEduName(v string) *HrbrainImportEduExpRequestBody {
	s.EduName = &v
	return s
}

func (s *HrbrainImportEduExpRequestBody) SetEndDate(v string) *HrbrainImportEduExpRequestBody {
	s.EndDate = &v
	return s
}

func (s *HrbrainImportEduExpRequestBody) SetExtendInfo(v map[string]interface{}) *HrbrainImportEduExpRequestBody {
	s.ExtendInfo = v
	return s
}

func (s *HrbrainImportEduExpRequestBody) SetMajor(v string) *HrbrainImportEduExpRequestBody {
	s.Major = &v
	return s
}

func (s *HrbrainImportEduExpRequestBody) SetName(v string) *HrbrainImportEduExpRequestBody {
	s.Name = &v
	return s
}

func (s *HrbrainImportEduExpRequestBody) SetSchoolName(v string) *HrbrainImportEduExpRequestBody {
	s.SchoolName = &v
	return s
}

func (s *HrbrainImportEduExpRequestBody) SetStartDate(v string) *HrbrainImportEduExpRequestBody {
	s.StartDate = &v
	return s
}

func (s *HrbrainImportEduExpRequestBody) SetWorkNo(v string) *HrbrainImportEduExpRequestBody {
	s.WorkNo = &v
	return s
}

type HrbrainImportEduExpResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainImportEduExpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportEduExpResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportEduExpResponseBody) SetRequestId(v string) *HrbrainImportEduExpResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainImportEduExpResponseBody) SetResult(v bool) *HrbrainImportEduExpResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainImportEduExpResponseBody) SetSuccess(v bool) *HrbrainImportEduExpResponseBody {
	s.Success = &v
	return s
}

type HrbrainImportEduExpResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainImportEduExpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainImportEduExpResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportEduExpResponse) GoString() string {
	return s.String()
}

func (s *HrbrainImportEduExpResponse) SetHeaders(v map[string]*string) *HrbrainImportEduExpResponse {
	s.Headers = v
	return s
}

func (s *HrbrainImportEduExpResponse) SetStatusCode(v int32) *HrbrainImportEduExpResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainImportEduExpResponse) SetBody(v *HrbrainImportEduExpResponseBody) *HrbrainImportEduExpResponse {
	s.Body = v
	return s
}

type HrbrainImportEmpInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainImportEmpInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportEmpInfoHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainImportEmpInfoHeaders) SetCommonHeaders(v map[string]*string) *HrbrainImportEmpInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainImportEmpInfoHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainImportEmpInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainImportEmpInfoRequest struct {
	Body   []*HrbrainImportEmpInfoRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	CorpId *string                            `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s HrbrainImportEmpInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportEmpInfoRequest) GoString() string {
	return s.String()
}

func (s *HrbrainImportEmpInfoRequest) SetBody(v []*HrbrainImportEmpInfoRequestBody) *HrbrainImportEmpInfoRequest {
	s.Body = v
	return s
}

func (s *HrbrainImportEmpInfoRequest) SetCorpId(v string) *HrbrainImportEmpInfoRequest {
	s.CorpId = &v
	return s
}

type HrbrainImportEmpInfoRequestBody struct {
	Birthday        *string                `json:"birthday,omitempty" xml:"birthday,omitempty"`
	DeptName        *string                `json:"deptName,omitempty" xml:"deptName,omitempty"`
	DeptNo          *string                `json:"deptNo,omitempty" xml:"deptNo,omitempty"`
	DimissionDate   *string                `json:"dimissionDate,omitempty" xml:"dimissionDate,omitempty"`
	EmpSource       *string                `json:"empSource,omitempty" xml:"empSource,omitempty"`
	EmpStatus       *string                `json:"empStatus,omitempty" xml:"empStatus,omitempty"`
	EmpType         *string                `json:"empType,omitempty" xml:"empType,omitempty"`
	ExtendInfo      map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	Gender          *string                `json:"gender,omitempty" xml:"gender,omitempty"`
	HighestDegree   *string                `json:"highestDegree,omitempty" xml:"highestDegree,omitempty"`
	HighestEduName  *string                `json:"highestEduName,omitempty" xml:"highestEduName,omitempty"`
	IsDimission     *string                `json:"isDimission,omitempty" xml:"isDimission,omitempty"`
	JobCodeName     *string                `json:"jobCodeName,omitempty" xml:"jobCodeName,omitempty"`
	JobLevel        *string                `json:"jobLevel,omitempty" xml:"jobLevel,omitempty"`
	LastSchoolName  *string                `json:"lastSchoolName,omitempty" xml:"lastSchoolName,omitempty"`
	Marriage        *string                `json:"marriage,omitempty" xml:"marriage,omitempty"`
	Name            *string                `json:"name,omitempty" xml:"name,omitempty"`
	Nation          *string                `json:"nation,omitempty" xml:"nation,omitempty"`
	NationCtry      *string                `json:"nationCtry,omitempty" xml:"nationCtry,omitempty"`
	PoliticalStatus *string                `json:"politicalStatus,omitempty" xml:"politicalStatus,omitempty"`
	PostName        *string                `json:"postName,omitempty" xml:"postName,omitempty"`
	RegistDate      *string                `json:"registDate,omitempty" xml:"registDate,omitempty"`
	RegularDate     *string                `json:"regularDate,omitempty" xml:"regularDate,omitempty"`
	SuperEmpId      *string                `json:"superEmpId,omitempty" xml:"superEmpId,omitempty"`
	SuperName       *string                `json:"superName,omitempty" xml:"superName,omitempty"`
	WorkEmail       *string                `json:"workEmail,omitempty" xml:"workEmail,omitempty"`
	WorkLocAddr     *string                `json:"workLocAddr,omitempty" xml:"workLocAddr,omitempty"`
	WorkLocCity     *string                `json:"workLocCity,omitempty" xml:"workLocCity,omitempty"`
	WorkNo          *string                `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainImportEmpInfoRequestBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportEmpInfoRequestBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportEmpInfoRequestBody) SetBirthday(v string) *HrbrainImportEmpInfoRequestBody {
	s.Birthday = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetDeptName(v string) *HrbrainImportEmpInfoRequestBody {
	s.DeptName = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetDeptNo(v string) *HrbrainImportEmpInfoRequestBody {
	s.DeptNo = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetDimissionDate(v string) *HrbrainImportEmpInfoRequestBody {
	s.DimissionDate = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetEmpSource(v string) *HrbrainImportEmpInfoRequestBody {
	s.EmpSource = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetEmpStatus(v string) *HrbrainImportEmpInfoRequestBody {
	s.EmpStatus = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetEmpType(v string) *HrbrainImportEmpInfoRequestBody {
	s.EmpType = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetExtendInfo(v map[string]interface{}) *HrbrainImportEmpInfoRequestBody {
	s.ExtendInfo = v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetGender(v string) *HrbrainImportEmpInfoRequestBody {
	s.Gender = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetHighestDegree(v string) *HrbrainImportEmpInfoRequestBody {
	s.HighestDegree = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetHighestEduName(v string) *HrbrainImportEmpInfoRequestBody {
	s.HighestEduName = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetIsDimission(v string) *HrbrainImportEmpInfoRequestBody {
	s.IsDimission = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetJobCodeName(v string) *HrbrainImportEmpInfoRequestBody {
	s.JobCodeName = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetJobLevel(v string) *HrbrainImportEmpInfoRequestBody {
	s.JobLevel = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetLastSchoolName(v string) *HrbrainImportEmpInfoRequestBody {
	s.LastSchoolName = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetMarriage(v string) *HrbrainImportEmpInfoRequestBody {
	s.Marriage = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetName(v string) *HrbrainImportEmpInfoRequestBody {
	s.Name = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetNation(v string) *HrbrainImportEmpInfoRequestBody {
	s.Nation = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetNationCtry(v string) *HrbrainImportEmpInfoRequestBody {
	s.NationCtry = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetPoliticalStatus(v string) *HrbrainImportEmpInfoRequestBody {
	s.PoliticalStatus = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetPostName(v string) *HrbrainImportEmpInfoRequestBody {
	s.PostName = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetRegistDate(v string) *HrbrainImportEmpInfoRequestBody {
	s.RegistDate = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetRegularDate(v string) *HrbrainImportEmpInfoRequestBody {
	s.RegularDate = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetSuperEmpId(v string) *HrbrainImportEmpInfoRequestBody {
	s.SuperEmpId = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetSuperName(v string) *HrbrainImportEmpInfoRequestBody {
	s.SuperName = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetWorkEmail(v string) *HrbrainImportEmpInfoRequestBody {
	s.WorkEmail = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetWorkLocAddr(v string) *HrbrainImportEmpInfoRequestBody {
	s.WorkLocAddr = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetWorkLocCity(v string) *HrbrainImportEmpInfoRequestBody {
	s.WorkLocCity = &v
	return s
}

func (s *HrbrainImportEmpInfoRequestBody) SetWorkNo(v string) *HrbrainImportEmpInfoRequestBody {
	s.WorkNo = &v
	return s
}

type HrbrainImportEmpInfoResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainImportEmpInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportEmpInfoResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportEmpInfoResponseBody) SetRequestId(v string) *HrbrainImportEmpInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainImportEmpInfoResponseBody) SetResult(v bool) *HrbrainImportEmpInfoResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainImportEmpInfoResponseBody) SetSuccess(v bool) *HrbrainImportEmpInfoResponseBody {
	s.Success = &v
	return s
}

type HrbrainImportEmpInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainImportEmpInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainImportEmpInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportEmpInfoResponse) GoString() string {
	return s.String()
}

func (s *HrbrainImportEmpInfoResponse) SetHeaders(v map[string]*string) *HrbrainImportEmpInfoResponse {
	s.Headers = v
	return s
}

func (s *HrbrainImportEmpInfoResponse) SetStatusCode(v int32) *HrbrainImportEmpInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainImportEmpInfoResponse) SetBody(v *HrbrainImportEmpInfoResponseBody) *HrbrainImportEmpInfoResponse {
	s.Body = v
	return s
}

type HrbrainImportLabelBaseHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainImportLabelBaseHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportLabelBaseHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainImportLabelBaseHeaders) SetCommonHeaders(v map[string]*string) *HrbrainImportLabelBaseHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainImportLabelBaseHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainImportLabelBaseHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainImportLabelBaseRequest struct {
	Body   []*HrbrainImportLabelBaseRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	CorpId *string                              `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s HrbrainImportLabelBaseRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportLabelBaseRequest) GoString() string {
	return s.String()
}

func (s *HrbrainImportLabelBaseRequest) SetBody(v []*HrbrainImportLabelBaseRequestBody) *HrbrainImportLabelBaseRequest {
	s.Body = v
	return s
}

func (s *HrbrainImportLabelBaseRequest) SetCorpId(v string) *HrbrainImportLabelBaseRequest {
	s.CorpId = &v
	return s
}

type HrbrainImportLabelBaseRequestBody struct {
	ExtendInfo map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	Name       *string                `json:"name,omitempty" xml:"name,omitempty"`
	WorkNo     *string                `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainImportLabelBaseRequestBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportLabelBaseRequestBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportLabelBaseRequestBody) SetExtendInfo(v map[string]interface{}) *HrbrainImportLabelBaseRequestBody {
	s.ExtendInfo = v
	return s
}

func (s *HrbrainImportLabelBaseRequestBody) SetName(v string) *HrbrainImportLabelBaseRequestBody {
	s.Name = &v
	return s
}

func (s *HrbrainImportLabelBaseRequestBody) SetWorkNo(v string) *HrbrainImportLabelBaseRequestBody {
	s.WorkNo = &v
	return s
}

type HrbrainImportLabelBaseResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainImportLabelBaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportLabelBaseResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportLabelBaseResponseBody) SetRequestId(v string) *HrbrainImportLabelBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainImportLabelBaseResponseBody) SetResult(v bool) *HrbrainImportLabelBaseResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainImportLabelBaseResponseBody) SetSuccess(v bool) *HrbrainImportLabelBaseResponseBody {
	s.Success = &v
	return s
}

type HrbrainImportLabelBaseResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainImportLabelBaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainImportLabelBaseResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportLabelBaseResponse) GoString() string {
	return s.String()
}

func (s *HrbrainImportLabelBaseResponse) SetHeaders(v map[string]*string) *HrbrainImportLabelBaseResponse {
	s.Headers = v
	return s
}

func (s *HrbrainImportLabelBaseResponse) SetStatusCode(v int32) *HrbrainImportLabelBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainImportLabelBaseResponse) SetBody(v *HrbrainImportLabelBaseResponseBody) *HrbrainImportLabelBaseResponse {
	s.Body = v
	return s
}

type HrbrainImportLabelCustomHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainImportLabelCustomHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportLabelCustomHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainImportLabelCustomHeaders) SetCommonHeaders(v map[string]*string) *HrbrainImportLabelCustomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainImportLabelCustomHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainImportLabelCustomHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainImportLabelCustomRequest struct {
	Body   []*HrbrainImportLabelCustomRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	CorpId *string                                `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s HrbrainImportLabelCustomRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportLabelCustomRequest) GoString() string {
	return s.String()
}

func (s *HrbrainImportLabelCustomRequest) SetBody(v []*HrbrainImportLabelCustomRequestBody) *HrbrainImportLabelCustomRequest {
	s.Body = v
	return s
}

func (s *HrbrainImportLabelCustomRequest) SetCorpId(v string) *HrbrainImportLabelCustomRequest {
	s.CorpId = &v
	return s
}

type HrbrainImportLabelCustomRequestBody struct {
	ExtendInfo map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	Name       *string                `json:"name,omitempty" xml:"name,omitempty"`
	Tag        *string                `json:"tag,omitempty" xml:"tag,omitempty"`
	WorkNo     *string                `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainImportLabelCustomRequestBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportLabelCustomRequestBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportLabelCustomRequestBody) SetExtendInfo(v map[string]interface{}) *HrbrainImportLabelCustomRequestBody {
	s.ExtendInfo = v
	return s
}

func (s *HrbrainImportLabelCustomRequestBody) SetName(v string) *HrbrainImportLabelCustomRequestBody {
	s.Name = &v
	return s
}

func (s *HrbrainImportLabelCustomRequestBody) SetTag(v string) *HrbrainImportLabelCustomRequestBody {
	s.Tag = &v
	return s
}

func (s *HrbrainImportLabelCustomRequestBody) SetWorkNo(v string) *HrbrainImportLabelCustomRequestBody {
	s.WorkNo = &v
	return s
}

type HrbrainImportLabelCustomResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainImportLabelCustomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportLabelCustomResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportLabelCustomResponseBody) SetRequestId(v string) *HrbrainImportLabelCustomResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainImportLabelCustomResponseBody) SetResult(v bool) *HrbrainImportLabelCustomResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainImportLabelCustomResponseBody) SetSuccess(v bool) *HrbrainImportLabelCustomResponseBody {
	s.Success = &v
	return s
}

type HrbrainImportLabelCustomResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainImportLabelCustomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainImportLabelCustomResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportLabelCustomResponse) GoString() string {
	return s.String()
}

func (s *HrbrainImportLabelCustomResponse) SetHeaders(v map[string]*string) *HrbrainImportLabelCustomResponse {
	s.Headers = v
	return s
}

func (s *HrbrainImportLabelCustomResponse) SetStatusCode(v int32) *HrbrainImportLabelCustomResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainImportLabelCustomResponse) SetBody(v *HrbrainImportLabelCustomResponseBody) *HrbrainImportLabelCustomResponse {
	s.Body = v
	return s
}

type HrbrainImportLabelIndustryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainImportLabelIndustryHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportLabelIndustryHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainImportLabelIndustryHeaders) SetCommonHeaders(v map[string]*string) *HrbrainImportLabelIndustryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainImportLabelIndustryHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainImportLabelIndustryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainImportLabelIndustryRequest struct {
	Body   []*HrbrainImportLabelIndustryRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	CorpId *string                                  `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s HrbrainImportLabelIndustryRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportLabelIndustryRequest) GoString() string {
	return s.String()
}

func (s *HrbrainImportLabelIndustryRequest) SetBody(v []*HrbrainImportLabelIndustryRequestBody) *HrbrainImportLabelIndustryRequest {
	s.Body = v
	return s
}

func (s *HrbrainImportLabelIndustryRequest) SetCorpId(v string) *HrbrainImportLabelIndustryRequest {
	s.CorpId = &v
	return s
}

type HrbrainImportLabelIndustryRequestBody struct {
	ExtendInfo map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	Level1     *string                `json:"level1,omitempty" xml:"level1,omitempty"`
	Level2     *string                `json:"level2,omitempty" xml:"level2,omitempty"`
	Level3     *string                `json:"level3,omitempty" xml:"level3,omitempty"`
	Name       *string                `json:"name,omitempty" xml:"name,omitempty"`
	WorkNo     *string                `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainImportLabelIndustryRequestBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportLabelIndustryRequestBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportLabelIndustryRequestBody) SetExtendInfo(v map[string]interface{}) *HrbrainImportLabelIndustryRequestBody {
	s.ExtendInfo = v
	return s
}

func (s *HrbrainImportLabelIndustryRequestBody) SetLevel1(v string) *HrbrainImportLabelIndustryRequestBody {
	s.Level1 = &v
	return s
}

func (s *HrbrainImportLabelIndustryRequestBody) SetLevel2(v string) *HrbrainImportLabelIndustryRequestBody {
	s.Level2 = &v
	return s
}

func (s *HrbrainImportLabelIndustryRequestBody) SetLevel3(v string) *HrbrainImportLabelIndustryRequestBody {
	s.Level3 = &v
	return s
}

func (s *HrbrainImportLabelIndustryRequestBody) SetName(v string) *HrbrainImportLabelIndustryRequestBody {
	s.Name = &v
	return s
}

func (s *HrbrainImportLabelIndustryRequestBody) SetWorkNo(v string) *HrbrainImportLabelIndustryRequestBody {
	s.WorkNo = &v
	return s
}

type HrbrainImportLabelIndustryResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainImportLabelIndustryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportLabelIndustryResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportLabelIndustryResponseBody) SetRequestId(v string) *HrbrainImportLabelIndustryResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainImportLabelIndustryResponseBody) SetResult(v bool) *HrbrainImportLabelIndustryResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainImportLabelIndustryResponseBody) SetSuccess(v bool) *HrbrainImportLabelIndustryResponseBody {
	s.Success = &v
	return s
}

type HrbrainImportLabelIndustryResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainImportLabelIndustryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainImportLabelIndustryResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportLabelIndustryResponse) GoString() string {
	return s.String()
}

func (s *HrbrainImportLabelIndustryResponse) SetHeaders(v map[string]*string) *HrbrainImportLabelIndustryResponse {
	s.Headers = v
	return s
}

func (s *HrbrainImportLabelIndustryResponse) SetStatusCode(v int32) *HrbrainImportLabelIndustryResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainImportLabelIndustryResponse) SetBody(v *HrbrainImportLabelIndustryResponseBody) *HrbrainImportLabelIndustryResponse {
	s.Body = v
	return s
}

type HrbrainImportLabelInventoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainImportLabelInventoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportLabelInventoryHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainImportLabelInventoryHeaders) SetCommonHeaders(v map[string]*string) *HrbrainImportLabelInventoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainImportLabelInventoryHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainImportLabelInventoryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainImportLabelInventoryRequest struct {
	Body   []*HrbrainImportLabelInventoryRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	CorpId *string                                   `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s HrbrainImportLabelInventoryRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportLabelInventoryRequest) GoString() string {
	return s.String()
}

func (s *HrbrainImportLabelInventoryRequest) SetBody(v []*HrbrainImportLabelInventoryRequestBody) *HrbrainImportLabelInventoryRequest {
	s.Body = v
	return s
}

func (s *HrbrainImportLabelInventoryRequest) SetCorpId(v string) *HrbrainImportLabelInventoryRequest {
	s.CorpId = &v
	return s
}

type HrbrainImportLabelInventoryRequestBody struct {
	ExtendInfo map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	Name       *string                `json:"name,omitempty" xml:"name,omitempty"`
	Period     *string                `json:"period,omitempty" xml:"period,omitempty"`
	WorkNo     *string                `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainImportLabelInventoryRequestBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportLabelInventoryRequestBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportLabelInventoryRequestBody) SetExtendInfo(v map[string]interface{}) *HrbrainImportLabelInventoryRequestBody {
	s.ExtendInfo = v
	return s
}

func (s *HrbrainImportLabelInventoryRequestBody) SetName(v string) *HrbrainImportLabelInventoryRequestBody {
	s.Name = &v
	return s
}

func (s *HrbrainImportLabelInventoryRequestBody) SetPeriod(v string) *HrbrainImportLabelInventoryRequestBody {
	s.Period = &v
	return s
}

func (s *HrbrainImportLabelInventoryRequestBody) SetWorkNo(v string) *HrbrainImportLabelInventoryRequestBody {
	s.WorkNo = &v
	return s
}

type HrbrainImportLabelInventoryResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainImportLabelInventoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportLabelInventoryResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportLabelInventoryResponseBody) SetRequestId(v string) *HrbrainImportLabelInventoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainImportLabelInventoryResponseBody) SetResult(v bool) *HrbrainImportLabelInventoryResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainImportLabelInventoryResponseBody) SetSuccess(v bool) *HrbrainImportLabelInventoryResponseBody {
	s.Success = &v
	return s
}

type HrbrainImportLabelInventoryResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainImportLabelInventoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainImportLabelInventoryResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportLabelInventoryResponse) GoString() string {
	return s.String()
}

func (s *HrbrainImportLabelInventoryResponse) SetHeaders(v map[string]*string) *HrbrainImportLabelInventoryResponse {
	s.Headers = v
	return s
}

func (s *HrbrainImportLabelInventoryResponse) SetStatusCode(v int32) *HrbrainImportLabelInventoryResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainImportLabelInventoryResponse) SetBody(v *HrbrainImportLabelInventoryResponseBody) *HrbrainImportLabelInventoryResponse {
	s.Body = v
	return s
}

type HrbrainImportLabelProfSkillHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainImportLabelProfSkillHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportLabelProfSkillHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainImportLabelProfSkillHeaders) SetCommonHeaders(v map[string]*string) *HrbrainImportLabelProfSkillHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainImportLabelProfSkillHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainImportLabelProfSkillHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainImportLabelProfSkillRequest struct {
	Body   []*HrbrainImportLabelProfSkillRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	CorpId *string                                   `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s HrbrainImportLabelProfSkillRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportLabelProfSkillRequest) GoString() string {
	return s.String()
}

func (s *HrbrainImportLabelProfSkillRequest) SetBody(v []*HrbrainImportLabelProfSkillRequestBody) *HrbrainImportLabelProfSkillRequest {
	s.Body = v
	return s
}

func (s *HrbrainImportLabelProfSkillRequest) SetCorpId(v string) *HrbrainImportLabelProfSkillRequest {
	s.CorpId = &v
	return s
}

type HrbrainImportLabelProfSkillRequestBody struct {
	ExtendInfo map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	Level1     *string                `json:"level1,omitempty" xml:"level1,omitempty"`
	Level2     *string                `json:"level2,omitempty" xml:"level2,omitempty"`
	Level3     *string                `json:"level3,omitempty" xml:"level3,omitempty"`
	Name       *string                `json:"name,omitempty" xml:"name,omitempty"`
	WorkNo     *string                `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainImportLabelProfSkillRequestBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportLabelProfSkillRequestBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportLabelProfSkillRequestBody) SetExtendInfo(v map[string]interface{}) *HrbrainImportLabelProfSkillRequestBody {
	s.ExtendInfo = v
	return s
}

func (s *HrbrainImportLabelProfSkillRequestBody) SetLevel1(v string) *HrbrainImportLabelProfSkillRequestBody {
	s.Level1 = &v
	return s
}

func (s *HrbrainImportLabelProfSkillRequestBody) SetLevel2(v string) *HrbrainImportLabelProfSkillRequestBody {
	s.Level2 = &v
	return s
}

func (s *HrbrainImportLabelProfSkillRequestBody) SetLevel3(v string) *HrbrainImportLabelProfSkillRequestBody {
	s.Level3 = &v
	return s
}

func (s *HrbrainImportLabelProfSkillRequestBody) SetName(v string) *HrbrainImportLabelProfSkillRequestBody {
	s.Name = &v
	return s
}

func (s *HrbrainImportLabelProfSkillRequestBody) SetWorkNo(v string) *HrbrainImportLabelProfSkillRequestBody {
	s.WorkNo = &v
	return s
}

type HrbrainImportLabelProfSkillResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainImportLabelProfSkillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportLabelProfSkillResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportLabelProfSkillResponseBody) SetRequestId(v string) *HrbrainImportLabelProfSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainImportLabelProfSkillResponseBody) SetResult(v bool) *HrbrainImportLabelProfSkillResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainImportLabelProfSkillResponseBody) SetSuccess(v bool) *HrbrainImportLabelProfSkillResponseBody {
	s.Success = &v
	return s
}

type HrbrainImportLabelProfSkillResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainImportLabelProfSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainImportLabelProfSkillResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportLabelProfSkillResponse) GoString() string {
	return s.String()
}

func (s *HrbrainImportLabelProfSkillResponse) SetHeaders(v map[string]*string) *HrbrainImportLabelProfSkillResponse {
	s.Headers = v
	return s
}

func (s *HrbrainImportLabelProfSkillResponse) SetStatusCode(v int32) *HrbrainImportLabelProfSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainImportLabelProfSkillResponse) SetBody(v *HrbrainImportLabelProfSkillResponseBody) *HrbrainImportLabelProfSkillResponse {
	s.Body = v
	return s
}

type HrbrainImportPerfEvalHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainImportPerfEvalHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportPerfEvalHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainImportPerfEvalHeaders) SetCommonHeaders(v map[string]*string) *HrbrainImportPerfEvalHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainImportPerfEvalHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainImportPerfEvalHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainImportPerfEvalRequest struct {
	Body   []*HrbrainImportPerfEvalRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	CorpId *string                             `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s HrbrainImportPerfEvalRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportPerfEvalRequest) GoString() string {
	return s.String()
}

func (s *HrbrainImportPerfEvalRequest) SetBody(v []*HrbrainImportPerfEvalRequestBody) *HrbrainImportPerfEvalRequest {
	s.Body = v
	return s
}

func (s *HrbrainImportPerfEvalRequest) SetCorpId(v string) *HrbrainImportPerfEvalRequest {
	s.CorpId = &v
	return s
}

type HrbrainImportPerfEvalRequestBody struct {
	Comment         *string                `json:"comment,omitempty" xml:"comment,omitempty"`
	ExtendInfo      map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	Name            *string                `json:"name,omitempty" xml:"name,omitempty"`
	PerfCate        *string                `json:"perfCate,omitempty" xml:"perfCate,omitempty"`
	PerfPlanName    *string                `json:"perfPlanName,omitempty" xml:"perfPlanName,omitempty"`
	PerfScore       *string                `json:"perfScore,omitempty" xml:"perfScore,omitempty"`
	Period          *string                `json:"period,omitempty" xml:"period,omitempty"`
	PeriodEndDate   *string                `json:"periodEndDate,omitempty" xml:"periodEndDate,omitempty"`
	PeriodStartDate *string                `json:"periodStartDate,omitempty" xml:"periodStartDate,omitempty"`
	Score           *string                `json:"score,omitempty" xml:"score,omitempty"`
	WorkNo          *string                `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainImportPerfEvalRequestBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportPerfEvalRequestBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportPerfEvalRequestBody) SetComment(v string) *HrbrainImportPerfEvalRequestBody {
	s.Comment = &v
	return s
}

func (s *HrbrainImportPerfEvalRequestBody) SetExtendInfo(v map[string]interface{}) *HrbrainImportPerfEvalRequestBody {
	s.ExtendInfo = v
	return s
}

func (s *HrbrainImportPerfEvalRequestBody) SetName(v string) *HrbrainImportPerfEvalRequestBody {
	s.Name = &v
	return s
}

func (s *HrbrainImportPerfEvalRequestBody) SetPerfCate(v string) *HrbrainImportPerfEvalRequestBody {
	s.PerfCate = &v
	return s
}

func (s *HrbrainImportPerfEvalRequestBody) SetPerfPlanName(v string) *HrbrainImportPerfEvalRequestBody {
	s.PerfPlanName = &v
	return s
}

func (s *HrbrainImportPerfEvalRequestBody) SetPerfScore(v string) *HrbrainImportPerfEvalRequestBody {
	s.PerfScore = &v
	return s
}

func (s *HrbrainImportPerfEvalRequestBody) SetPeriod(v string) *HrbrainImportPerfEvalRequestBody {
	s.Period = &v
	return s
}

func (s *HrbrainImportPerfEvalRequestBody) SetPeriodEndDate(v string) *HrbrainImportPerfEvalRequestBody {
	s.PeriodEndDate = &v
	return s
}

func (s *HrbrainImportPerfEvalRequestBody) SetPeriodStartDate(v string) *HrbrainImportPerfEvalRequestBody {
	s.PeriodStartDate = &v
	return s
}

func (s *HrbrainImportPerfEvalRequestBody) SetScore(v string) *HrbrainImportPerfEvalRequestBody {
	s.Score = &v
	return s
}

func (s *HrbrainImportPerfEvalRequestBody) SetWorkNo(v string) *HrbrainImportPerfEvalRequestBody {
	s.WorkNo = &v
	return s
}

type HrbrainImportPerfEvalResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainImportPerfEvalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportPerfEvalResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportPerfEvalResponseBody) SetRequestId(v string) *HrbrainImportPerfEvalResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainImportPerfEvalResponseBody) SetResult(v bool) *HrbrainImportPerfEvalResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainImportPerfEvalResponseBody) SetSuccess(v bool) *HrbrainImportPerfEvalResponseBody {
	s.Success = &v
	return s
}

type HrbrainImportPerfEvalResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainImportPerfEvalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainImportPerfEvalResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportPerfEvalResponse) GoString() string {
	return s.String()
}

func (s *HrbrainImportPerfEvalResponse) SetHeaders(v map[string]*string) *HrbrainImportPerfEvalResponse {
	s.Headers = v
	return s
}

func (s *HrbrainImportPerfEvalResponse) SetStatusCode(v int32) *HrbrainImportPerfEvalResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainImportPerfEvalResponse) SetBody(v *HrbrainImportPerfEvalResponseBody) *HrbrainImportPerfEvalResponse {
	s.Body = v
	return s
}

type HrbrainImportPromEvalHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainImportPromEvalHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportPromEvalHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainImportPromEvalHeaders) SetCommonHeaders(v map[string]*string) *HrbrainImportPromEvalHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainImportPromEvalHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainImportPromEvalHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainImportPromEvalRequest struct {
	Body   []*HrbrainImportPromEvalRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	CorpId *string                             `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s HrbrainImportPromEvalRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportPromEvalRequest) GoString() string {
	return s.String()
}

func (s *HrbrainImportPromEvalRequest) SetBody(v []*HrbrainImportPromEvalRequestBody) *HrbrainImportPromEvalRequest {
	s.Body = v
	return s
}

func (s *HrbrainImportPromEvalRequest) SetCorpId(v string) *HrbrainImportPromEvalRequest {
	s.CorpId = &v
	return s
}

type HrbrainImportPromEvalRequestBody struct {
	Comment         *string                `json:"comment,omitempty" xml:"comment,omitempty"`
	EffectiveDate   *string                `json:"effectiveDate,omitempty" xml:"effectiveDate,omitempty"`
	ExtendInfo      map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	Name            *string                `json:"name,omitempty" xml:"name,omitempty"`
	NewJobLevel     *string                `json:"newJobLevel,omitempty" xml:"newJobLevel,omitempty"`
	Period          *string                `json:"period,omitempty" xml:"period,omitempty"`
	PeriodEndDate   *string                `json:"periodEndDate,omitempty" xml:"periodEndDate,omitempty"`
	PeriodStartDate *string                `json:"periodStartDate,omitempty" xml:"periodStartDate,omitempty"`
	PreJobLevel     *string                `json:"preJobLevel,omitempty" xml:"preJobLevel,omitempty"`
	WorkNo          *string                `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainImportPromEvalRequestBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportPromEvalRequestBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportPromEvalRequestBody) SetComment(v string) *HrbrainImportPromEvalRequestBody {
	s.Comment = &v
	return s
}

func (s *HrbrainImportPromEvalRequestBody) SetEffectiveDate(v string) *HrbrainImportPromEvalRequestBody {
	s.EffectiveDate = &v
	return s
}

func (s *HrbrainImportPromEvalRequestBody) SetExtendInfo(v map[string]interface{}) *HrbrainImportPromEvalRequestBody {
	s.ExtendInfo = v
	return s
}

func (s *HrbrainImportPromEvalRequestBody) SetName(v string) *HrbrainImportPromEvalRequestBody {
	s.Name = &v
	return s
}

func (s *HrbrainImportPromEvalRequestBody) SetNewJobLevel(v string) *HrbrainImportPromEvalRequestBody {
	s.NewJobLevel = &v
	return s
}

func (s *HrbrainImportPromEvalRequestBody) SetPeriod(v string) *HrbrainImportPromEvalRequestBody {
	s.Period = &v
	return s
}

func (s *HrbrainImportPromEvalRequestBody) SetPeriodEndDate(v string) *HrbrainImportPromEvalRequestBody {
	s.PeriodEndDate = &v
	return s
}

func (s *HrbrainImportPromEvalRequestBody) SetPeriodStartDate(v string) *HrbrainImportPromEvalRequestBody {
	s.PeriodStartDate = &v
	return s
}

func (s *HrbrainImportPromEvalRequestBody) SetPreJobLevel(v string) *HrbrainImportPromEvalRequestBody {
	s.PreJobLevel = &v
	return s
}

func (s *HrbrainImportPromEvalRequestBody) SetWorkNo(v string) *HrbrainImportPromEvalRequestBody {
	s.WorkNo = &v
	return s
}

type HrbrainImportPromEvalResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainImportPromEvalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportPromEvalResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportPromEvalResponseBody) SetRequestId(v string) *HrbrainImportPromEvalResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainImportPromEvalResponseBody) SetResult(v bool) *HrbrainImportPromEvalResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainImportPromEvalResponseBody) SetSuccess(v bool) *HrbrainImportPromEvalResponseBody {
	s.Success = &v
	return s
}

type HrbrainImportPromEvalResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainImportPromEvalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainImportPromEvalResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportPromEvalResponse) GoString() string {
	return s.String()
}

func (s *HrbrainImportPromEvalResponse) SetHeaders(v map[string]*string) *HrbrainImportPromEvalResponse {
	s.Headers = v
	return s
}

func (s *HrbrainImportPromEvalResponse) SetStatusCode(v int32) *HrbrainImportPromEvalResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainImportPromEvalResponse) SetBody(v *HrbrainImportPromEvalResponseBody) *HrbrainImportPromEvalResponse {
	s.Body = v
	return s
}

type HrbrainImportPunDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainImportPunDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportPunDetailHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainImportPunDetailHeaders) SetCommonHeaders(v map[string]*string) *HrbrainImportPunDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainImportPunDetailHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainImportPunDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainImportPunDetailRequest struct {
	Body   []*HrbrainImportPunDetailRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	CorpId *string                              `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s HrbrainImportPunDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportPunDetailRequest) GoString() string {
	return s.String()
}

func (s *HrbrainImportPunDetailRequest) SetBody(v []*HrbrainImportPunDetailRequestBody) *HrbrainImportPunDetailRequest {
	s.Body = v
	return s
}

func (s *HrbrainImportPunDetailRequest) SetCorpId(v string) *HrbrainImportPunDetailRequest {
	s.CorpId = &v
	return s
}

type HrbrainImportPunDetailRequestBody struct {
	Comment       *string                `json:"comment,omitempty" xml:"comment,omitempty"`
	EffectiveDate *string                `json:"effectiveDate,omitempty" xml:"effectiveDate,omitempty"`
	ExtendInfo    map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	Name          *string                `json:"name,omitempty" xml:"name,omitempty"`
	PunName       *string                `json:"punName,omitempty" xml:"punName,omitempty"`
	PunOrg        *string                `json:"punOrg,omitempty" xml:"punOrg,omitempty"`
	WorkNo        *string                `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainImportPunDetailRequestBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportPunDetailRequestBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportPunDetailRequestBody) SetComment(v string) *HrbrainImportPunDetailRequestBody {
	s.Comment = &v
	return s
}

func (s *HrbrainImportPunDetailRequestBody) SetEffectiveDate(v string) *HrbrainImportPunDetailRequestBody {
	s.EffectiveDate = &v
	return s
}

func (s *HrbrainImportPunDetailRequestBody) SetExtendInfo(v map[string]interface{}) *HrbrainImportPunDetailRequestBody {
	s.ExtendInfo = v
	return s
}

func (s *HrbrainImportPunDetailRequestBody) SetName(v string) *HrbrainImportPunDetailRequestBody {
	s.Name = &v
	return s
}

func (s *HrbrainImportPunDetailRequestBody) SetPunName(v string) *HrbrainImportPunDetailRequestBody {
	s.PunName = &v
	return s
}

func (s *HrbrainImportPunDetailRequestBody) SetPunOrg(v string) *HrbrainImportPunDetailRequestBody {
	s.PunOrg = &v
	return s
}

func (s *HrbrainImportPunDetailRequestBody) SetWorkNo(v string) *HrbrainImportPunDetailRequestBody {
	s.WorkNo = &v
	return s
}

type HrbrainImportPunDetailResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainImportPunDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportPunDetailResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportPunDetailResponseBody) SetRequestId(v string) *HrbrainImportPunDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainImportPunDetailResponseBody) SetResult(v bool) *HrbrainImportPunDetailResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainImportPunDetailResponseBody) SetSuccess(v bool) *HrbrainImportPunDetailResponseBody {
	s.Success = &v
	return s
}

type HrbrainImportPunDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainImportPunDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainImportPunDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportPunDetailResponse) GoString() string {
	return s.String()
}

func (s *HrbrainImportPunDetailResponse) SetHeaders(v map[string]*string) *HrbrainImportPunDetailResponse {
	s.Headers = v
	return s
}

func (s *HrbrainImportPunDetailResponse) SetStatusCode(v int32) *HrbrainImportPunDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainImportPunDetailResponse) SetBody(v *HrbrainImportPunDetailResponseBody) *HrbrainImportPunDetailResponse {
	s.Body = v
	return s
}

type HrbrainImportRegistHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainImportRegistHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportRegistHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainImportRegistHeaders) SetCommonHeaders(v map[string]*string) *HrbrainImportRegistHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainImportRegistHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainImportRegistHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainImportRegistRequest struct {
	Body   []*HrbrainImportRegistRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	CorpId *string                           `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s HrbrainImportRegistRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportRegistRequest) GoString() string {
	return s.String()
}

func (s *HrbrainImportRegistRequest) SetBody(v []*HrbrainImportRegistRequestBody) *HrbrainImportRegistRequest {
	s.Body = v
	return s
}

func (s *HrbrainImportRegistRequest) SetCorpId(v string) *HrbrainImportRegistRequest {
	s.CorpId = &v
	return s
}

type HrbrainImportRegistRequestBody struct {
	DeptName    *string                `json:"deptName,omitempty" xml:"deptName,omitempty"`
	DeptNo      *string                `json:"deptNo,omitempty" xml:"deptNo,omitempty"`
	EmpSource   *string                `json:"empSource,omitempty" xml:"empSource,omitempty"`
	EmpType     *string                `json:"empType,omitempty" xml:"empType,omitempty"`
	ExtendInfo  map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	JobCodeName *string                `json:"jobCodeName,omitempty" xml:"jobCodeName,omitempty"`
	JobLevel    *string                `json:"jobLevel,omitempty" xml:"jobLevel,omitempty"`
	Name        *string                `json:"name,omitempty" xml:"name,omitempty"`
	PostName    *string                `json:"postName,omitempty" xml:"postName,omitempty"`
	RegistDate  *string                `json:"registDate,omitempty" xml:"registDate,omitempty"`
	SuperName   *string                `json:"superName,omitempty" xml:"superName,omitempty"`
	WorkLocAddr *string                `json:"workLocAddr,omitempty" xml:"workLocAddr,omitempty"`
	WorkNo      *string                `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainImportRegistRequestBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportRegistRequestBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportRegistRequestBody) SetDeptName(v string) *HrbrainImportRegistRequestBody {
	s.DeptName = &v
	return s
}

func (s *HrbrainImportRegistRequestBody) SetDeptNo(v string) *HrbrainImportRegistRequestBody {
	s.DeptNo = &v
	return s
}

func (s *HrbrainImportRegistRequestBody) SetEmpSource(v string) *HrbrainImportRegistRequestBody {
	s.EmpSource = &v
	return s
}

func (s *HrbrainImportRegistRequestBody) SetEmpType(v string) *HrbrainImportRegistRequestBody {
	s.EmpType = &v
	return s
}

func (s *HrbrainImportRegistRequestBody) SetExtendInfo(v map[string]interface{}) *HrbrainImportRegistRequestBody {
	s.ExtendInfo = v
	return s
}

func (s *HrbrainImportRegistRequestBody) SetJobCodeName(v string) *HrbrainImportRegistRequestBody {
	s.JobCodeName = &v
	return s
}

func (s *HrbrainImportRegistRequestBody) SetJobLevel(v string) *HrbrainImportRegistRequestBody {
	s.JobLevel = &v
	return s
}

func (s *HrbrainImportRegistRequestBody) SetName(v string) *HrbrainImportRegistRequestBody {
	s.Name = &v
	return s
}

func (s *HrbrainImportRegistRequestBody) SetPostName(v string) *HrbrainImportRegistRequestBody {
	s.PostName = &v
	return s
}

func (s *HrbrainImportRegistRequestBody) SetRegistDate(v string) *HrbrainImportRegistRequestBody {
	s.RegistDate = &v
	return s
}

func (s *HrbrainImportRegistRequestBody) SetSuperName(v string) *HrbrainImportRegistRequestBody {
	s.SuperName = &v
	return s
}

func (s *HrbrainImportRegistRequestBody) SetWorkLocAddr(v string) *HrbrainImportRegistRequestBody {
	s.WorkLocAddr = &v
	return s
}

func (s *HrbrainImportRegistRequestBody) SetWorkNo(v string) *HrbrainImportRegistRequestBody {
	s.WorkNo = &v
	return s
}

type HrbrainImportRegistResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainImportRegistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportRegistResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportRegistResponseBody) SetRequestId(v string) *HrbrainImportRegistResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainImportRegistResponseBody) SetResult(v bool) *HrbrainImportRegistResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainImportRegistResponseBody) SetSuccess(v bool) *HrbrainImportRegistResponseBody {
	s.Success = &v
	return s
}

type HrbrainImportRegistResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainImportRegistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainImportRegistResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportRegistResponse) GoString() string {
	return s.String()
}

func (s *HrbrainImportRegistResponse) SetHeaders(v map[string]*string) *HrbrainImportRegistResponse {
	s.Headers = v
	return s
}

func (s *HrbrainImportRegistResponse) SetStatusCode(v int32) *HrbrainImportRegistResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainImportRegistResponse) SetBody(v *HrbrainImportRegistResponseBody) *HrbrainImportRegistResponse {
	s.Body = v
	return s
}

type HrbrainImportTransferEvalHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainImportTransferEvalHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportTransferEvalHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainImportTransferEvalHeaders) SetCommonHeaders(v map[string]*string) *HrbrainImportTransferEvalHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainImportTransferEvalHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainImportTransferEvalHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainImportTransferEvalRequest struct {
	Body   []*HrbrainImportTransferEvalRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	CorpId *string                                 `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s HrbrainImportTransferEvalRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportTransferEvalRequest) GoString() string {
	return s.String()
}

func (s *HrbrainImportTransferEvalRequest) SetBody(v []*HrbrainImportTransferEvalRequestBody) *HrbrainImportTransferEvalRequest {
	s.Body = v
	return s
}

func (s *HrbrainImportTransferEvalRequest) SetCorpId(v string) *HrbrainImportTransferEvalRequest {
	s.CorpId = &v
	return s
}

type HrbrainImportTransferEvalRequestBody struct {
	CurrInfo       map[string]interface{} `json:"currInfo,omitempty" xml:"currInfo,omitempty"`
	ExtendInfo     map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	Name           *string                `json:"name,omitempty" xml:"name,omitempty"`
	PreInfo        map[string]interface{} `json:"preInfo,omitempty" xml:"preInfo,omitempty"`
	TransferDate   *string                `json:"transferDate,omitempty" xml:"transferDate,omitempty"`
	TransferReason *string                `json:"transferReason,omitempty" xml:"transferReason,omitempty"`
	TransferType   *string                `json:"transferType,omitempty" xml:"transferType,omitempty"`
	WorkNo         *string                `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainImportTransferEvalRequestBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportTransferEvalRequestBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportTransferEvalRequestBody) SetCurrInfo(v map[string]interface{}) *HrbrainImportTransferEvalRequestBody {
	s.CurrInfo = v
	return s
}

func (s *HrbrainImportTransferEvalRequestBody) SetExtendInfo(v map[string]interface{}) *HrbrainImportTransferEvalRequestBody {
	s.ExtendInfo = v
	return s
}

func (s *HrbrainImportTransferEvalRequestBody) SetName(v string) *HrbrainImportTransferEvalRequestBody {
	s.Name = &v
	return s
}

func (s *HrbrainImportTransferEvalRequestBody) SetPreInfo(v map[string]interface{}) *HrbrainImportTransferEvalRequestBody {
	s.PreInfo = v
	return s
}

func (s *HrbrainImportTransferEvalRequestBody) SetTransferDate(v string) *HrbrainImportTransferEvalRequestBody {
	s.TransferDate = &v
	return s
}

func (s *HrbrainImportTransferEvalRequestBody) SetTransferReason(v string) *HrbrainImportTransferEvalRequestBody {
	s.TransferReason = &v
	return s
}

func (s *HrbrainImportTransferEvalRequestBody) SetTransferType(v string) *HrbrainImportTransferEvalRequestBody {
	s.TransferType = &v
	return s
}

func (s *HrbrainImportTransferEvalRequestBody) SetWorkNo(v string) *HrbrainImportTransferEvalRequestBody {
	s.WorkNo = &v
	return s
}

type HrbrainImportTransferEvalResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainImportTransferEvalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportTransferEvalResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportTransferEvalResponseBody) SetRequestId(v string) *HrbrainImportTransferEvalResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainImportTransferEvalResponseBody) SetResult(v bool) *HrbrainImportTransferEvalResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainImportTransferEvalResponseBody) SetSuccess(v bool) *HrbrainImportTransferEvalResponseBody {
	s.Success = &v
	return s
}

type HrbrainImportTransferEvalResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainImportTransferEvalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainImportTransferEvalResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportTransferEvalResponse) GoString() string {
	return s.String()
}

func (s *HrbrainImportTransferEvalResponse) SetHeaders(v map[string]*string) *HrbrainImportTransferEvalResponse {
	s.Headers = v
	return s
}

func (s *HrbrainImportTransferEvalResponse) SetStatusCode(v int32) *HrbrainImportTransferEvalResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainImportTransferEvalResponse) SetBody(v *HrbrainImportTransferEvalResponseBody) *HrbrainImportTransferEvalResponse {
	s.Body = v
	return s
}

type HrbrainImportWorkExpHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainImportWorkExpHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportWorkExpHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainImportWorkExpHeaders) SetCommonHeaders(v map[string]*string) *HrbrainImportWorkExpHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainImportWorkExpHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainImportWorkExpHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainImportWorkExpRequest struct {
	Body   []*HrbrainImportWorkExpRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	CorpId *string                            `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s HrbrainImportWorkExpRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportWorkExpRequest) GoString() string {
	return s.String()
}

func (s *HrbrainImportWorkExpRequest) SetBody(v []*HrbrainImportWorkExpRequestBody) *HrbrainImportWorkExpRequest {
	s.Body = v
	return s
}

func (s *HrbrainImportWorkExpRequest) SetCorpId(v string) *HrbrainImportWorkExpRequest {
	s.CorpId = &v
	return s
}

type HrbrainImportWorkExpRequestBody struct {
	CompanyName *string                `json:"companyName,omitempty" xml:"companyName,omitempty"`
	EndDate     *string                `json:"endDate,omitempty" xml:"endDate,omitempty"`
	ExtendInfo  map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	JobDesc     *string                `json:"jobDesc,omitempty" xml:"jobDesc,omitempty"`
	Name        *string                `json:"name,omitempty" xml:"name,omitempty"`
	PostName    *string                `json:"postName,omitempty" xml:"postName,omitempty"`
	StartDate   *string                `json:"startDate,omitempty" xml:"startDate,omitempty"`
	WorkNo      *string                `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainImportWorkExpRequestBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportWorkExpRequestBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportWorkExpRequestBody) SetCompanyName(v string) *HrbrainImportWorkExpRequestBody {
	s.CompanyName = &v
	return s
}

func (s *HrbrainImportWorkExpRequestBody) SetEndDate(v string) *HrbrainImportWorkExpRequestBody {
	s.EndDate = &v
	return s
}

func (s *HrbrainImportWorkExpRequestBody) SetExtendInfo(v map[string]interface{}) *HrbrainImportWorkExpRequestBody {
	s.ExtendInfo = v
	return s
}

func (s *HrbrainImportWorkExpRequestBody) SetJobDesc(v string) *HrbrainImportWorkExpRequestBody {
	s.JobDesc = &v
	return s
}

func (s *HrbrainImportWorkExpRequestBody) SetName(v string) *HrbrainImportWorkExpRequestBody {
	s.Name = &v
	return s
}

func (s *HrbrainImportWorkExpRequestBody) SetPostName(v string) *HrbrainImportWorkExpRequestBody {
	s.PostName = &v
	return s
}

func (s *HrbrainImportWorkExpRequestBody) SetStartDate(v string) *HrbrainImportWorkExpRequestBody {
	s.StartDate = &v
	return s
}

func (s *HrbrainImportWorkExpRequestBody) SetWorkNo(v string) *HrbrainImportWorkExpRequestBody {
	s.WorkNo = &v
	return s
}

type HrbrainImportWorkExpResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainImportWorkExpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportWorkExpResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportWorkExpResponseBody) SetRequestId(v string) *HrbrainImportWorkExpResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainImportWorkExpResponseBody) SetResult(v bool) *HrbrainImportWorkExpResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainImportWorkExpResponseBody) SetSuccess(v bool) *HrbrainImportWorkExpResponseBody {
	s.Success = &v
	return s
}

type HrbrainImportWorkExpResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainImportWorkExpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainImportWorkExpResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportWorkExpResponse) GoString() string {
	return s.String()
}

func (s *HrbrainImportWorkExpResponse) SetHeaders(v map[string]*string) *HrbrainImportWorkExpResponse {
	s.Headers = v
	return s
}

func (s *HrbrainImportWorkExpResponse) SetStatusCode(v int32) *HrbrainImportWorkExpResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainImportWorkExpResponse) SetBody(v *HrbrainImportWorkExpResponseBody) *HrbrainImportWorkExpResponse {
	s.Body = v
	return s
}

type SyncDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SyncDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncDataHeaders) GoString() string {
	return s.String()
}

func (s *SyncDataHeaders) SetCommonHeaders(v map[string]*string) *SyncDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncDataHeaders) SetXAcsDingtalkAccessToken(v string) *SyncDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SyncDataRequest struct {
	Content   *string `json:"content,omitempty" xml:"content,omitempty"`
	DataId    *string `json:"dataId,omitempty" xml:"dataId,omitempty"`
	EtlTime   *string `json:"etlTime,omitempty" xml:"etlTime,omitempty"`
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	SchemaId  *string `json:"schemaId,omitempty" xml:"schemaId,omitempty"`
}

func (s SyncDataRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncDataRequest) GoString() string {
	return s.String()
}

func (s *SyncDataRequest) SetContent(v string) *SyncDataRequest {
	s.Content = &v
	return s
}

func (s *SyncDataRequest) SetDataId(v string) *SyncDataRequest {
	s.DataId = &v
	return s
}

func (s *SyncDataRequest) SetEtlTime(v string) *SyncDataRequest {
	s.EtlTime = &v
	return s
}

func (s *SyncDataRequest) SetProjectId(v string) *SyncDataRequest {
	s.ProjectId = &v
	return s
}

func (s *SyncDataRequest) SetSchemaId(v string) *SyncDataRequest {
	s.SchemaId = &v
	return s
}

type SyncDataResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SyncDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncDataResponseBody) GoString() string {
	return s.String()
}

func (s *SyncDataResponseBody) SetSuccess(v bool) *SyncDataResponseBody {
	s.Success = &v
	return s
}

type SyncDataResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncDataResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncDataResponse) GoString() string {
	return s.String()
}

func (s *SyncDataResponse) SetHeaders(v map[string]*string) *SyncDataResponse {
	s.Headers = v
	return s
}

func (s *SyncDataResponse) SetStatusCode(v int32) *SyncDataResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncDataResponse) SetBody(v *SyncDataResponseBody) *SyncDataResponse {
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

func (client *Client) HrbrainImportAwardDetailWithOptions(request *HrbrainImportAwardDetailRequest, headers *HrbrainImportAwardDetailHeaders, runtime *util.RuntimeOptions) (_result *HrbrainImportAwardDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
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
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("HrbrainImportAwardDetail"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/awardDetails/import"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainImportAwardDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HrbrainImportAwardDetail(request *HrbrainImportAwardDetailRequest) (_result *HrbrainImportAwardDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainImportAwardDetailHeaders{}
	_result = &HrbrainImportAwardDetailResponse{}
	_body, _err := client.HrbrainImportAwardDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HrbrainImportDeptInfoWithOptions(request *HrbrainImportDeptInfoRequest, headers *HrbrainImportDeptInfoHeaders, runtime *util.RuntimeOptions) (_result *HrbrainImportDeptInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
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
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("HrbrainImportDeptInfo"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/deptInfos/import"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainImportDeptInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HrbrainImportDeptInfo(request *HrbrainImportDeptInfoRequest) (_result *HrbrainImportDeptInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainImportDeptInfoHeaders{}
	_result = &HrbrainImportDeptInfoResponse{}
	_body, _err := client.HrbrainImportDeptInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HrbrainImportDimissionWithOptions(request *HrbrainImportDimissionRequest, headers *HrbrainImportDimissionHeaders, runtime *util.RuntimeOptions) (_result *HrbrainImportDimissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
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
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("HrbrainImportDimission"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/dimissionInfos/import"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainImportDimissionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HrbrainImportDimission(request *HrbrainImportDimissionRequest) (_result *HrbrainImportDimissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainImportDimissionHeaders{}
	_result = &HrbrainImportDimissionResponse{}
	_body, _err := client.HrbrainImportDimissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HrbrainImportEduExpWithOptions(request *HrbrainImportEduExpRequest, headers *HrbrainImportEduExpHeaders, runtime *util.RuntimeOptions) (_result *HrbrainImportEduExpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
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
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("HrbrainImportEduExp"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/eduExperiences/import"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainImportEduExpResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HrbrainImportEduExp(request *HrbrainImportEduExpRequest) (_result *HrbrainImportEduExpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainImportEduExpHeaders{}
	_result = &HrbrainImportEduExpResponse{}
	_body, _err := client.HrbrainImportEduExpWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HrbrainImportEmpInfoWithOptions(request *HrbrainImportEmpInfoRequest, headers *HrbrainImportEmpInfoHeaders, runtime *util.RuntimeOptions) (_result *HrbrainImportEmpInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
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
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("HrbrainImportEmpInfo"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/empInfos/import"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainImportEmpInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HrbrainImportEmpInfo(request *HrbrainImportEmpInfoRequest) (_result *HrbrainImportEmpInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainImportEmpInfoHeaders{}
	_result = &HrbrainImportEmpInfoResponse{}
	_body, _err := client.HrbrainImportEmpInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HrbrainImportLabelBaseWithOptions(request *HrbrainImportLabelBaseRequest, headers *HrbrainImportLabelBaseHeaders, runtime *util.RuntimeOptions) (_result *HrbrainImportLabelBaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
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
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("HrbrainImportLabelBase"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/baseLabels/import"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainImportLabelBaseResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HrbrainImportLabelBase(request *HrbrainImportLabelBaseRequest) (_result *HrbrainImportLabelBaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainImportLabelBaseHeaders{}
	_result = &HrbrainImportLabelBaseResponse{}
	_body, _err := client.HrbrainImportLabelBaseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HrbrainImportLabelCustomWithOptions(request *HrbrainImportLabelCustomRequest, headers *HrbrainImportLabelCustomHeaders, runtime *util.RuntimeOptions) (_result *HrbrainImportLabelCustomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
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
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("HrbrainImportLabelCustom"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/customLabels/import"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainImportLabelCustomResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HrbrainImportLabelCustom(request *HrbrainImportLabelCustomRequest) (_result *HrbrainImportLabelCustomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainImportLabelCustomHeaders{}
	_result = &HrbrainImportLabelCustomResponse{}
	_body, _err := client.HrbrainImportLabelCustomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HrbrainImportLabelIndustryWithOptions(request *HrbrainImportLabelIndustryRequest, headers *HrbrainImportLabelIndustryHeaders, runtime *util.RuntimeOptions) (_result *HrbrainImportLabelIndustryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
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
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("HrbrainImportLabelIndustry"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/industries/import"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainImportLabelIndustryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HrbrainImportLabelIndustry(request *HrbrainImportLabelIndustryRequest) (_result *HrbrainImportLabelIndustryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainImportLabelIndustryHeaders{}
	_result = &HrbrainImportLabelIndustryResponse{}
	_body, _err := client.HrbrainImportLabelIndustryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HrbrainImportLabelInventoryWithOptions(request *HrbrainImportLabelInventoryRequest, headers *HrbrainImportLabelInventoryHeaders, runtime *util.RuntimeOptions) (_result *HrbrainImportLabelInventoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
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
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("HrbrainImportLabelInventory"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/inventories/import"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainImportLabelInventoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HrbrainImportLabelInventory(request *HrbrainImportLabelInventoryRequest) (_result *HrbrainImportLabelInventoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainImportLabelInventoryHeaders{}
	_result = &HrbrainImportLabelInventoryResponse{}
	_body, _err := client.HrbrainImportLabelInventoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HrbrainImportLabelProfSkillWithOptions(request *HrbrainImportLabelProfSkillRequest, headers *HrbrainImportLabelProfSkillHeaders, runtime *util.RuntimeOptions) (_result *HrbrainImportLabelProfSkillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
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
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("HrbrainImportLabelProfSkill"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/profSkills/import"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainImportLabelProfSkillResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HrbrainImportLabelProfSkill(request *HrbrainImportLabelProfSkillRequest) (_result *HrbrainImportLabelProfSkillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainImportLabelProfSkillHeaders{}
	_result = &HrbrainImportLabelProfSkillResponse{}
	_body, _err := client.HrbrainImportLabelProfSkillWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HrbrainImportPerfEvalWithOptions(request *HrbrainImportPerfEvalRequest, headers *HrbrainImportPerfEvalHeaders, runtime *util.RuntimeOptions) (_result *HrbrainImportPerfEvalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
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
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("HrbrainImportPerfEval"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/perfRecords/import"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainImportPerfEvalResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HrbrainImportPerfEval(request *HrbrainImportPerfEvalRequest) (_result *HrbrainImportPerfEvalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainImportPerfEvalHeaders{}
	_result = &HrbrainImportPerfEvalResponse{}
	_body, _err := client.HrbrainImportPerfEvalWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HrbrainImportPromEvalWithOptions(request *HrbrainImportPromEvalRequest, headers *HrbrainImportPromEvalHeaders, runtime *util.RuntimeOptions) (_result *HrbrainImportPromEvalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
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
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("HrbrainImportPromEval"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/promRecords/import"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainImportPromEvalResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HrbrainImportPromEval(request *HrbrainImportPromEvalRequest) (_result *HrbrainImportPromEvalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainImportPromEvalHeaders{}
	_result = &HrbrainImportPromEvalResponse{}
	_body, _err := client.HrbrainImportPromEvalWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HrbrainImportPunDetailWithOptions(request *HrbrainImportPunDetailRequest, headers *HrbrainImportPunDetailHeaders, runtime *util.RuntimeOptions) (_result *HrbrainImportPunDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
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
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("HrbrainImportPunDetail"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/punDetails/import"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainImportPunDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HrbrainImportPunDetail(request *HrbrainImportPunDetailRequest) (_result *HrbrainImportPunDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainImportPunDetailHeaders{}
	_result = &HrbrainImportPunDetailResponse{}
	_body, _err := client.HrbrainImportPunDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HrbrainImportRegistWithOptions(request *HrbrainImportRegistRequest, headers *HrbrainImportRegistHeaders, runtime *util.RuntimeOptions) (_result *HrbrainImportRegistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
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
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("HrbrainImportRegist"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/registerInfos/import"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainImportRegistResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HrbrainImportRegist(request *HrbrainImportRegistRequest) (_result *HrbrainImportRegistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainImportRegistHeaders{}
	_result = &HrbrainImportRegistResponse{}
	_body, _err := client.HrbrainImportRegistWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HrbrainImportTransferEvalWithOptions(request *HrbrainImportTransferEvalRequest, headers *HrbrainImportTransferEvalHeaders, runtime *util.RuntimeOptions) (_result *HrbrainImportTransferEvalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
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
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("HrbrainImportTransferEval"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/changeRecords/import"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainImportTransferEvalResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HrbrainImportTransferEval(request *HrbrainImportTransferEvalRequest) (_result *HrbrainImportTransferEvalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainImportTransferEvalHeaders{}
	_result = &HrbrainImportTransferEvalResponse{}
	_body, _err := client.HrbrainImportTransferEvalWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HrbrainImportWorkExpWithOptions(request *HrbrainImportWorkExpRequest, headers *HrbrainImportWorkExpHeaders, runtime *util.RuntimeOptions) (_result *HrbrainImportWorkExpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
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
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("HrbrainImportWorkExp"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/workExperiences/import"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainImportWorkExpResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HrbrainImportWorkExp(request *HrbrainImportWorkExpRequest) (_result *HrbrainImportWorkExpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainImportWorkExpHeaders{}
	_result = &HrbrainImportWorkExpResponse{}
	_body, _err := client.HrbrainImportWorkExpWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncDataWithOptions(request *SyncDataRequest, headers *SyncDataHeaders, runtime *util.RuntimeOptions) (_result *SyncDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DataId)) {
		body["dataId"] = request.DataId
	}

	if !tea.BoolValue(util.IsUnset(request.EtlTime)) {
		body["etlTime"] = request.EtlTime
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["projectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaId)) {
		body["schemaId"] = request.SchemaId
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
		Action:      tea.String("SyncData"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncData(request *SyncDataRequest) (_result *SyncDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncDataHeaders{}
	_result = &SyncDataResponse{}
	_body, _err := client.SyncDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
