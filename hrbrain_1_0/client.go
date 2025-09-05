// This file is auto-generated, don't edit it. Thanks.
package hrbrain_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type HrbrainDeleteAwardRecordsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainDeleteAwardRecordsHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteAwardRecordsHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteAwardRecordsHeaders) SetCommonHeaders(v map[string]*string) *HrbrainDeleteAwardRecordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainDeleteAwardRecordsHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainDeleteAwardRecordsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainDeleteAwardRecordsRequest struct {
	Params []*HrbrainDeleteAwardRecordsRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Repeated"`
}

func (s HrbrainDeleteAwardRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteAwardRecordsRequest) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteAwardRecordsRequest) SetParams(v []*HrbrainDeleteAwardRecordsRequestParams) *HrbrainDeleteAwardRecordsRequest {
	s.Params = v
	return s
}

type HrbrainDeleteAwardRecordsRequestParams struct {
	AwardDate *string `json:"awardDate,omitempty" xml:"awardDate,omitempty"`
	AwardName *string `json:"awardName,omitempty" xml:"awardName,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainDeleteAwardRecordsRequestParams) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteAwardRecordsRequestParams) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteAwardRecordsRequestParams) SetAwardDate(v string) *HrbrainDeleteAwardRecordsRequestParams {
	s.AwardDate = &v
	return s
}

func (s *HrbrainDeleteAwardRecordsRequestParams) SetAwardName(v string) *HrbrainDeleteAwardRecordsRequestParams {
	s.AwardName = &v
	return s
}

func (s *HrbrainDeleteAwardRecordsRequestParams) SetWorkNo(v string) *HrbrainDeleteAwardRecordsRequestParams {
	s.WorkNo = &v
	return s
}

type HrbrainDeleteAwardRecordsResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainDeleteAwardRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteAwardRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteAwardRecordsResponseBody) SetRequestId(v string) *HrbrainDeleteAwardRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainDeleteAwardRecordsResponseBody) SetResult(v bool) *HrbrainDeleteAwardRecordsResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainDeleteAwardRecordsResponseBody) SetSuccess(v bool) *HrbrainDeleteAwardRecordsResponseBody {
	s.Success = &v
	return s
}

type HrbrainDeleteAwardRecordsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainDeleteAwardRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainDeleteAwardRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteAwardRecordsResponse) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteAwardRecordsResponse) SetHeaders(v map[string]*string) *HrbrainDeleteAwardRecordsResponse {
	s.Headers = v
	return s
}

func (s *HrbrainDeleteAwardRecordsResponse) SetStatusCode(v int32) *HrbrainDeleteAwardRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainDeleteAwardRecordsResponse) SetBody(v *HrbrainDeleteAwardRecordsResponseBody) *HrbrainDeleteAwardRecordsResponse {
	s.Body = v
	return s
}

type HrbrainDeleteCustomHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainDeleteCustomHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteCustomHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteCustomHeaders) SetCommonHeaders(v map[string]*string) *HrbrainDeleteCustomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainDeleteCustomHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainDeleteCustomHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainDeleteCustomRequest struct {
	// This parameter is required.
	ModelCode *string `json:"modelCode,omitempty" xml:"modelCode,omitempty"`
	// This parameter is required.
	Params []map[string]interface{} `json:"params,omitempty" xml:"params,omitempty" type:"Repeated"`
}

func (s HrbrainDeleteCustomRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteCustomRequest) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteCustomRequest) SetModelCode(v string) *HrbrainDeleteCustomRequest {
	s.ModelCode = &v
	return s
}

func (s *HrbrainDeleteCustomRequest) SetParams(v []map[string]interface{}) *HrbrainDeleteCustomRequest {
	s.Params = v
	return s
}

type HrbrainDeleteCustomResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainDeleteCustomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteCustomResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteCustomResponseBody) SetRequestId(v string) *HrbrainDeleteCustomResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainDeleteCustomResponseBody) SetResult(v bool) *HrbrainDeleteCustomResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainDeleteCustomResponseBody) SetSuccess(v bool) *HrbrainDeleteCustomResponseBody {
	s.Success = &v
	return s
}

type HrbrainDeleteCustomResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainDeleteCustomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainDeleteCustomResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteCustomResponse) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteCustomResponse) SetHeaders(v map[string]*string) *HrbrainDeleteCustomResponse {
	s.Headers = v
	return s
}

func (s *HrbrainDeleteCustomResponse) SetStatusCode(v int32) *HrbrainDeleteCustomResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainDeleteCustomResponse) SetBody(v *HrbrainDeleteCustomResponseBody) *HrbrainDeleteCustomResponse {
	s.Body = v
	return s
}

type HrbrainDeleteDeptInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainDeleteDeptInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteDeptInfoHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteDeptInfoHeaders) SetCommonHeaders(v map[string]*string) *HrbrainDeleteDeptInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainDeleteDeptInfoHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainDeleteDeptInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainDeleteDeptInfoRequest struct {
	Params []*HrbrainDeleteDeptInfoRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Repeated"`
}

func (s HrbrainDeleteDeptInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteDeptInfoRequest) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteDeptInfoRequest) SetParams(v []*HrbrainDeleteDeptInfoRequestParams) *HrbrainDeleteDeptInfoRequest {
	s.Params = v
	return s
}

type HrbrainDeleteDeptInfoRequestParams struct {
	// This parameter is required.
	DeptNo *string `json:"deptNo,omitempty" xml:"deptNo,omitempty"`
}

func (s HrbrainDeleteDeptInfoRequestParams) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteDeptInfoRequestParams) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteDeptInfoRequestParams) SetDeptNo(v string) *HrbrainDeleteDeptInfoRequestParams {
	s.DeptNo = &v
	return s
}

type HrbrainDeleteDeptInfoResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainDeleteDeptInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteDeptInfoResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteDeptInfoResponseBody) SetRequestId(v string) *HrbrainDeleteDeptInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainDeleteDeptInfoResponseBody) SetResult(v bool) *HrbrainDeleteDeptInfoResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainDeleteDeptInfoResponseBody) SetSuccess(v bool) *HrbrainDeleteDeptInfoResponseBody {
	s.Success = &v
	return s
}

type HrbrainDeleteDeptInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainDeleteDeptInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainDeleteDeptInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteDeptInfoResponse) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteDeptInfoResponse) SetHeaders(v map[string]*string) *HrbrainDeleteDeptInfoResponse {
	s.Headers = v
	return s
}

func (s *HrbrainDeleteDeptInfoResponse) SetStatusCode(v int32) *HrbrainDeleteDeptInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainDeleteDeptInfoResponse) SetBody(v *HrbrainDeleteDeptInfoResponseBody) *HrbrainDeleteDeptInfoResponse {
	s.Body = v
	return s
}

type HrbrainDeleteDimissionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainDeleteDimissionHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteDimissionHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteDimissionHeaders) SetCommonHeaders(v map[string]*string) *HrbrainDeleteDimissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainDeleteDimissionHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainDeleteDimissionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainDeleteDimissionRequest struct {
	Params []*HrbrainDeleteDimissionRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Repeated"`
}

func (s HrbrainDeleteDimissionRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteDimissionRequest) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteDimissionRequest) SetParams(v []*HrbrainDeleteDimissionRequestParams) *HrbrainDeleteDimissionRequest {
	s.Params = v
	return s
}

type HrbrainDeleteDimissionRequestParams struct {
	DimissionDate *string `json:"dimissionDate,omitempty" xml:"dimissionDate,omitempty"`
	WorkNo        *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainDeleteDimissionRequestParams) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteDimissionRequestParams) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteDimissionRequestParams) SetDimissionDate(v string) *HrbrainDeleteDimissionRequestParams {
	s.DimissionDate = &v
	return s
}

func (s *HrbrainDeleteDimissionRequestParams) SetWorkNo(v string) *HrbrainDeleteDimissionRequestParams {
	s.WorkNo = &v
	return s
}

type HrbrainDeleteDimissionResponseBody struct {
	// example:
	//
	// 480021443f9f37fcbf464c4a6b85d289
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainDeleteDimissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteDimissionResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteDimissionResponseBody) SetRequestId(v string) *HrbrainDeleteDimissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainDeleteDimissionResponseBody) SetResult(v bool) *HrbrainDeleteDimissionResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainDeleteDimissionResponseBody) SetSuccess(v bool) *HrbrainDeleteDimissionResponseBody {
	s.Success = &v
	return s
}

type HrbrainDeleteDimissionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainDeleteDimissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainDeleteDimissionResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteDimissionResponse) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteDimissionResponse) SetHeaders(v map[string]*string) *HrbrainDeleteDimissionResponse {
	s.Headers = v
	return s
}

func (s *HrbrainDeleteDimissionResponse) SetStatusCode(v int32) *HrbrainDeleteDimissionResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainDeleteDimissionResponse) SetBody(v *HrbrainDeleteDimissionResponseBody) *HrbrainDeleteDimissionResponse {
	s.Body = v
	return s
}

type HrbrainDeleteEduExpHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainDeleteEduExpHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteEduExpHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteEduExpHeaders) SetCommonHeaders(v map[string]*string) *HrbrainDeleteEduExpHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainDeleteEduExpHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainDeleteEduExpHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainDeleteEduExpRequest struct {
	Params []*HrbrainDeleteEduExpRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Repeated"`
}

func (s HrbrainDeleteEduExpRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteEduExpRequest) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteEduExpRequest) SetParams(v []*HrbrainDeleteEduExpRequestParams) *HrbrainDeleteEduExpRequest {
	s.Params = v
	return s
}

type HrbrainDeleteEduExpRequestParams struct {
	EduName    *string `json:"eduName,omitempty" xml:"eduName,omitempty"`
	EndDate    *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	SchoolName *string `json:"schoolName,omitempty" xml:"schoolName,omitempty"`
	StartDate  *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainDeleteEduExpRequestParams) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteEduExpRequestParams) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteEduExpRequestParams) SetEduName(v string) *HrbrainDeleteEduExpRequestParams {
	s.EduName = &v
	return s
}

func (s *HrbrainDeleteEduExpRequestParams) SetEndDate(v string) *HrbrainDeleteEduExpRequestParams {
	s.EndDate = &v
	return s
}

func (s *HrbrainDeleteEduExpRequestParams) SetSchoolName(v string) *HrbrainDeleteEduExpRequestParams {
	s.SchoolName = &v
	return s
}

func (s *HrbrainDeleteEduExpRequestParams) SetStartDate(v string) *HrbrainDeleteEduExpRequestParams {
	s.StartDate = &v
	return s
}

func (s *HrbrainDeleteEduExpRequestParams) SetWorkNo(v string) *HrbrainDeleteEduExpRequestParams {
	s.WorkNo = &v
	return s
}

type HrbrainDeleteEduExpResponseBody struct {
	// example:
	//
	// 480021443f9f37fcbf464c4a6b85d289
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainDeleteEduExpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteEduExpResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteEduExpResponseBody) SetRequestId(v string) *HrbrainDeleteEduExpResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainDeleteEduExpResponseBody) SetResult(v bool) *HrbrainDeleteEduExpResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainDeleteEduExpResponseBody) SetSuccess(v bool) *HrbrainDeleteEduExpResponseBody {
	s.Success = &v
	return s
}

type HrbrainDeleteEduExpResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainDeleteEduExpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainDeleteEduExpResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteEduExpResponse) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteEduExpResponse) SetHeaders(v map[string]*string) *HrbrainDeleteEduExpResponse {
	s.Headers = v
	return s
}

func (s *HrbrainDeleteEduExpResponse) SetStatusCode(v int32) *HrbrainDeleteEduExpResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainDeleteEduExpResponse) SetBody(v *HrbrainDeleteEduExpResponseBody) *HrbrainDeleteEduExpResponse {
	s.Body = v
	return s
}

type HrbrainDeleteEmpInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainDeleteEmpInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteEmpInfoHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteEmpInfoHeaders) SetCommonHeaders(v map[string]*string) *HrbrainDeleteEmpInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainDeleteEmpInfoHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainDeleteEmpInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainDeleteEmpInfoRequest struct {
	Params []*HrbrainDeleteEmpInfoRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Repeated"`
}

func (s HrbrainDeleteEmpInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteEmpInfoRequest) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteEmpInfoRequest) SetParams(v []*HrbrainDeleteEmpInfoRequestParams) *HrbrainDeleteEmpInfoRequest {
	s.Params = v
	return s
}

type HrbrainDeleteEmpInfoRequestParams struct {
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainDeleteEmpInfoRequestParams) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteEmpInfoRequestParams) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteEmpInfoRequestParams) SetWorkNo(v string) *HrbrainDeleteEmpInfoRequestParams {
	s.WorkNo = &v
	return s
}

type HrbrainDeleteEmpInfoResponseBody struct {
	// example:
	//
	// 480021443f9f37fcbf464c4a6b85d289
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainDeleteEmpInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteEmpInfoResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteEmpInfoResponseBody) SetRequestId(v string) *HrbrainDeleteEmpInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainDeleteEmpInfoResponseBody) SetResult(v bool) *HrbrainDeleteEmpInfoResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainDeleteEmpInfoResponseBody) SetSuccess(v bool) *HrbrainDeleteEmpInfoResponseBody {
	s.Success = &v
	return s
}

type HrbrainDeleteEmpInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainDeleteEmpInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainDeleteEmpInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteEmpInfoResponse) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteEmpInfoResponse) SetHeaders(v map[string]*string) *HrbrainDeleteEmpInfoResponse {
	s.Headers = v
	return s
}

func (s *HrbrainDeleteEmpInfoResponse) SetStatusCode(v int32) *HrbrainDeleteEmpInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainDeleteEmpInfoResponse) SetBody(v *HrbrainDeleteEmpInfoResponseBody) *HrbrainDeleteEmpInfoResponse {
	s.Body = v
	return s
}

type HrbrainDeleteLabelIndustryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainDeleteLabelIndustryHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteLabelIndustryHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteLabelIndustryHeaders) SetCommonHeaders(v map[string]*string) *HrbrainDeleteLabelIndustryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainDeleteLabelIndustryHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainDeleteLabelIndustryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainDeleteLabelIndustryRequest struct {
	Params []*HrbrainDeleteLabelIndustryRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Repeated"`
}

func (s HrbrainDeleteLabelIndustryRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteLabelIndustryRequest) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteLabelIndustryRequest) SetParams(v []*HrbrainDeleteLabelIndustryRequestParams) *HrbrainDeleteLabelIndustryRequest {
	s.Params = v
	return s
}

type HrbrainDeleteLabelIndustryRequestParams struct {
	Label map[string]interface{} `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainDeleteLabelIndustryRequestParams) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteLabelIndustryRequestParams) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteLabelIndustryRequestParams) SetLabel(v map[string]interface{}) *HrbrainDeleteLabelIndustryRequestParams {
	s.Label = v
	return s
}

func (s *HrbrainDeleteLabelIndustryRequestParams) SetWorkNo(v string) *HrbrainDeleteLabelIndustryRequestParams {
	s.WorkNo = &v
	return s
}

type HrbrainDeleteLabelIndustryResponseBody struct {
	// example:
	//
	// 480021443f9f37fcbf464c4a6b85d299
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainDeleteLabelIndustryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteLabelIndustryResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteLabelIndustryResponseBody) SetRequestId(v string) *HrbrainDeleteLabelIndustryResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainDeleteLabelIndustryResponseBody) SetResult(v bool) *HrbrainDeleteLabelIndustryResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainDeleteLabelIndustryResponseBody) SetSuccess(v bool) *HrbrainDeleteLabelIndustryResponseBody {
	s.Success = &v
	return s
}

type HrbrainDeleteLabelIndustryResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainDeleteLabelIndustryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainDeleteLabelIndustryResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteLabelIndustryResponse) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteLabelIndustryResponse) SetHeaders(v map[string]*string) *HrbrainDeleteLabelIndustryResponse {
	s.Headers = v
	return s
}

func (s *HrbrainDeleteLabelIndustryResponse) SetStatusCode(v int32) *HrbrainDeleteLabelIndustryResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainDeleteLabelIndustryResponse) SetBody(v *HrbrainDeleteLabelIndustryResponseBody) *HrbrainDeleteLabelIndustryResponse {
	s.Body = v
	return s
}

type HrbrainDeleteLabelInventoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainDeleteLabelInventoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteLabelInventoryHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteLabelInventoryHeaders) SetCommonHeaders(v map[string]*string) *HrbrainDeleteLabelInventoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainDeleteLabelInventoryHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainDeleteLabelInventoryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainDeleteLabelInventoryRequest struct {
	Params []*HrbrainDeleteLabelInventoryRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Repeated"`
}

func (s HrbrainDeleteLabelInventoryRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteLabelInventoryRequest) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteLabelInventoryRequest) SetParams(v []*HrbrainDeleteLabelInventoryRequestParams) *HrbrainDeleteLabelInventoryRequest {
	s.Params = v
	return s
}

type HrbrainDeleteLabelInventoryRequestParams struct {
	Label  map[string]interface{} `json:"label,omitempty" xml:"label,omitempty"`
	Period *string                `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainDeleteLabelInventoryRequestParams) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteLabelInventoryRequestParams) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteLabelInventoryRequestParams) SetLabel(v map[string]interface{}) *HrbrainDeleteLabelInventoryRequestParams {
	s.Label = v
	return s
}

func (s *HrbrainDeleteLabelInventoryRequestParams) SetPeriod(v string) *HrbrainDeleteLabelInventoryRequestParams {
	s.Period = &v
	return s
}

func (s *HrbrainDeleteLabelInventoryRequestParams) SetWorkNo(v string) *HrbrainDeleteLabelInventoryRequestParams {
	s.WorkNo = &v
	return s
}

type HrbrainDeleteLabelInventoryResponseBody struct {
	// example:
	//
	// 480021443f9f37fcbf464c4a6b85d299
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainDeleteLabelInventoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteLabelInventoryResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteLabelInventoryResponseBody) SetRequestId(v string) *HrbrainDeleteLabelInventoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainDeleteLabelInventoryResponseBody) SetResult(v bool) *HrbrainDeleteLabelInventoryResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainDeleteLabelInventoryResponseBody) SetSuccess(v bool) *HrbrainDeleteLabelInventoryResponseBody {
	s.Success = &v
	return s
}

type HrbrainDeleteLabelInventoryResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainDeleteLabelInventoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainDeleteLabelInventoryResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteLabelInventoryResponse) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteLabelInventoryResponse) SetHeaders(v map[string]*string) *HrbrainDeleteLabelInventoryResponse {
	s.Headers = v
	return s
}

func (s *HrbrainDeleteLabelInventoryResponse) SetStatusCode(v int32) *HrbrainDeleteLabelInventoryResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainDeleteLabelInventoryResponse) SetBody(v *HrbrainDeleteLabelInventoryResponseBody) *HrbrainDeleteLabelInventoryResponse {
	s.Body = v
	return s
}

type HrbrainDeleteLabelProfSkillHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainDeleteLabelProfSkillHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteLabelProfSkillHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteLabelProfSkillHeaders) SetCommonHeaders(v map[string]*string) *HrbrainDeleteLabelProfSkillHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainDeleteLabelProfSkillHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainDeleteLabelProfSkillHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainDeleteLabelProfSkillRequest struct {
	Params []*HrbrainDeleteLabelProfSkillRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Repeated"`
}

func (s HrbrainDeleteLabelProfSkillRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteLabelProfSkillRequest) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteLabelProfSkillRequest) SetParams(v []*HrbrainDeleteLabelProfSkillRequestParams) *HrbrainDeleteLabelProfSkillRequest {
	s.Params = v
	return s
}

type HrbrainDeleteLabelProfSkillRequestParams struct {
	Label map[string]interface{} `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainDeleteLabelProfSkillRequestParams) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteLabelProfSkillRequestParams) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteLabelProfSkillRequestParams) SetLabel(v map[string]interface{}) *HrbrainDeleteLabelProfSkillRequestParams {
	s.Label = v
	return s
}

func (s *HrbrainDeleteLabelProfSkillRequestParams) SetWorkNo(v string) *HrbrainDeleteLabelProfSkillRequestParams {
	s.WorkNo = &v
	return s
}

type HrbrainDeleteLabelProfSkillResponseBody struct {
	// example:
	//
	// 480021443f9f37fcbf464c4a6b85d289
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainDeleteLabelProfSkillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteLabelProfSkillResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteLabelProfSkillResponseBody) SetRequestId(v string) *HrbrainDeleteLabelProfSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainDeleteLabelProfSkillResponseBody) SetResult(v bool) *HrbrainDeleteLabelProfSkillResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainDeleteLabelProfSkillResponseBody) SetSuccess(v bool) *HrbrainDeleteLabelProfSkillResponseBody {
	s.Success = &v
	return s
}

type HrbrainDeleteLabelProfSkillResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainDeleteLabelProfSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainDeleteLabelProfSkillResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteLabelProfSkillResponse) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteLabelProfSkillResponse) SetHeaders(v map[string]*string) *HrbrainDeleteLabelProfSkillResponse {
	s.Headers = v
	return s
}

func (s *HrbrainDeleteLabelProfSkillResponse) SetStatusCode(v int32) *HrbrainDeleteLabelProfSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainDeleteLabelProfSkillResponse) SetBody(v *HrbrainDeleteLabelProfSkillResponseBody) *HrbrainDeleteLabelProfSkillResponse {
	s.Body = v
	return s
}

type HrbrainDeletePerfEvalHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainDeletePerfEvalHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeletePerfEvalHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainDeletePerfEvalHeaders) SetCommonHeaders(v map[string]*string) *HrbrainDeletePerfEvalHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainDeletePerfEvalHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainDeletePerfEvalHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainDeletePerfEvalRequest struct {
	Params []*HrbrainDeletePerfEvalRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Repeated"`
}

func (s HrbrainDeletePerfEvalRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeletePerfEvalRequest) GoString() string {
	return s.String()
}

func (s *HrbrainDeletePerfEvalRequest) SetParams(v []*HrbrainDeletePerfEvalRequestParams) *HrbrainDeletePerfEvalRequest {
	s.Params = v
	return s
}

type HrbrainDeletePerfEvalRequestParams struct {
	PerfPlanName *string `json:"perfPlanName,omitempty" xml:"perfPlanName,omitempty"`
	Period       *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainDeletePerfEvalRequestParams) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeletePerfEvalRequestParams) GoString() string {
	return s.String()
}

func (s *HrbrainDeletePerfEvalRequestParams) SetPerfPlanName(v string) *HrbrainDeletePerfEvalRequestParams {
	s.PerfPlanName = &v
	return s
}

func (s *HrbrainDeletePerfEvalRequestParams) SetPeriod(v string) *HrbrainDeletePerfEvalRequestParams {
	s.Period = &v
	return s
}

func (s *HrbrainDeletePerfEvalRequestParams) SetWorkNo(v string) *HrbrainDeletePerfEvalRequestParams {
	s.WorkNo = &v
	return s
}

type HrbrainDeletePerfEvalResponseBody struct {
	// example:
	//
	// 480021443f9f37fcbf464c4a6b85d289
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainDeletePerfEvalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeletePerfEvalResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainDeletePerfEvalResponseBody) SetRequestId(v string) *HrbrainDeletePerfEvalResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainDeletePerfEvalResponseBody) SetResult(v bool) *HrbrainDeletePerfEvalResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainDeletePerfEvalResponseBody) SetSuccess(v bool) *HrbrainDeletePerfEvalResponseBody {
	s.Success = &v
	return s
}

type HrbrainDeletePerfEvalResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainDeletePerfEvalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainDeletePerfEvalResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeletePerfEvalResponse) GoString() string {
	return s.String()
}

func (s *HrbrainDeletePerfEvalResponse) SetHeaders(v map[string]*string) *HrbrainDeletePerfEvalResponse {
	s.Headers = v
	return s
}

func (s *HrbrainDeletePerfEvalResponse) SetStatusCode(v int32) *HrbrainDeletePerfEvalResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainDeletePerfEvalResponse) SetBody(v *HrbrainDeletePerfEvalResponseBody) *HrbrainDeletePerfEvalResponse {
	s.Body = v
	return s
}

type HrbrainDeletePromRecordsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainDeletePromRecordsHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeletePromRecordsHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainDeletePromRecordsHeaders) SetCommonHeaders(v map[string]*string) *HrbrainDeletePromRecordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainDeletePromRecordsHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainDeletePromRecordsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainDeletePromRecordsRequest struct {
	Params []*HrbrainDeletePromRecordsRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Repeated"`
}

func (s HrbrainDeletePromRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeletePromRecordsRequest) GoString() string {
	return s.String()
}

func (s *HrbrainDeletePromRecordsRequest) SetParams(v []*HrbrainDeletePromRecordsRequestParams) *HrbrainDeletePromRecordsRequest {
	s.Params = v
	return s
}

type HrbrainDeletePromRecordsRequestParams struct {
	AwardDate *string `json:"awardDate,omitempty" xml:"awardDate,omitempty"`
	AwardName *string `json:"awardName,omitempty" xml:"awardName,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainDeletePromRecordsRequestParams) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeletePromRecordsRequestParams) GoString() string {
	return s.String()
}

func (s *HrbrainDeletePromRecordsRequestParams) SetAwardDate(v string) *HrbrainDeletePromRecordsRequestParams {
	s.AwardDate = &v
	return s
}

func (s *HrbrainDeletePromRecordsRequestParams) SetAwardName(v string) *HrbrainDeletePromRecordsRequestParams {
	s.AwardName = &v
	return s
}

func (s *HrbrainDeletePromRecordsRequestParams) SetWorkNo(v string) *HrbrainDeletePromRecordsRequestParams {
	s.WorkNo = &v
	return s
}

type HrbrainDeletePromRecordsResponseBody struct {
	// example:
	//
	// 480021443f9f37fcbf464c4a6b85d299
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainDeletePromRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeletePromRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainDeletePromRecordsResponseBody) SetRequestId(v string) *HrbrainDeletePromRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainDeletePromRecordsResponseBody) SetResult(v bool) *HrbrainDeletePromRecordsResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainDeletePromRecordsResponseBody) SetSuccess(v bool) *HrbrainDeletePromRecordsResponseBody {
	s.Success = &v
	return s
}

type HrbrainDeletePromRecordsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainDeletePromRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainDeletePromRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeletePromRecordsResponse) GoString() string {
	return s.String()
}

func (s *HrbrainDeletePromRecordsResponse) SetHeaders(v map[string]*string) *HrbrainDeletePromRecordsResponse {
	s.Headers = v
	return s
}

func (s *HrbrainDeletePromRecordsResponse) SetStatusCode(v int32) *HrbrainDeletePromRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainDeletePromRecordsResponse) SetBody(v *HrbrainDeletePromRecordsResponseBody) *HrbrainDeletePromRecordsResponse {
	s.Body = v
	return s
}

type HrbrainDeletePunDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainDeletePunDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeletePunDetailHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainDeletePunDetailHeaders) SetCommonHeaders(v map[string]*string) *HrbrainDeletePunDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainDeletePunDetailHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainDeletePunDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainDeletePunDetailRequest struct {
	Params []*HrbrainDeletePunDetailRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Repeated"`
}

func (s HrbrainDeletePunDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeletePunDetailRequest) GoString() string {
	return s.String()
}

func (s *HrbrainDeletePunDetailRequest) SetParams(v []*HrbrainDeletePunDetailRequestParams) *HrbrainDeletePunDetailRequest {
	s.Params = v
	return s
}

type HrbrainDeletePunDetailRequestParams struct {
	EffectiveDate *string `json:"effectiveDate,omitempty" xml:"effectiveDate,omitempty"`
	PunName       *string `json:"punName,omitempty" xml:"punName,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainDeletePunDetailRequestParams) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeletePunDetailRequestParams) GoString() string {
	return s.String()
}

func (s *HrbrainDeletePunDetailRequestParams) SetEffectiveDate(v string) *HrbrainDeletePunDetailRequestParams {
	s.EffectiveDate = &v
	return s
}

func (s *HrbrainDeletePunDetailRequestParams) SetPunName(v string) *HrbrainDeletePunDetailRequestParams {
	s.PunName = &v
	return s
}

func (s *HrbrainDeletePunDetailRequestParams) SetWorkNo(v string) *HrbrainDeletePunDetailRequestParams {
	s.WorkNo = &v
	return s
}

type HrbrainDeletePunDetailResponseBody struct {
	// example:
	//
	// 480021443f9f37fcbf464c4a6b85d289
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainDeletePunDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeletePunDetailResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainDeletePunDetailResponseBody) SetRequestId(v string) *HrbrainDeletePunDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainDeletePunDetailResponseBody) SetResult(v bool) *HrbrainDeletePunDetailResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainDeletePunDetailResponseBody) SetSuccess(v bool) *HrbrainDeletePunDetailResponseBody {
	s.Success = &v
	return s
}

type HrbrainDeletePunDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainDeletePunDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainDeletePunDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeletePunDetailResponse) GoString() string {
	return s.String()
}

func (s *HrbrainDeletePunDetailResponse) SetHeaders(v map[string]*string) *HrbrainDeletePunDetailResponse {
	s.Headers = v
	return s
}

func (s *HrbrainDeletePunDetailResponse) SetStatusCode(v int32) *HrbrainDeletePunDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainDeletePunDetailResponse) SetBody(v *HrbrainDeletePunDetailResponseBody) *HrbrainDeletePunDetailResponse {
	s.Body = v
	return s
}

type HrbrainDeleteRegistHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainDeleteRegistHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteRegistHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteRegistHeaders) SetCommonHeaders(v map[string]*string) *HrbrainDeleteRegistHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainDeleteRegistHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainDeleteRegistHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainDeleteRegistRequest struct {
	Params []*HrbrainDeleteRegistRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Repeated"`
}

func (s HrbrainDeleteRegistRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteRegistRequest) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteRegistRequest) SetParams(v []*HrbrainDeleteRegistRequestParams) *HrbrainDeleteRegistRequest {
	s.Params = v
	return s
}

type HrbrainDeleteRegistRequestParams struct {
	RegistDate *string `json:"registDate,omitempty" xml:"registDate,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainDeleteRegistRequestParams) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteRegistRequestParams) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteRegistRequestParams) SetRegistDate(v string) *HrbrainDeleteRegistRequestParams {
	s.RegistDate = &v
	return s
}

func (s *HrbrainDeleteRegistRequestParams) SetWorkNo(v string) *HrbrainDeleteRegistRequestParams {
	s.WorkNo = &v
	return s
}

type HrbrainDeleteRegistResponseBody struct {
	// example:
	//
	// 480021443f9f37fcbf464c4a6b85d289
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainDeleteRegistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteRegistResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteRegistResponseBody) SetRequestId(v string) *HrbrainDeleteRegistResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainDeleteRegistResponseBody) SetResult(v bool) *HrbrainDeleteRegistResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainDeleteRegistResponseBody) SetSuccess(v bool) *HrbrainDeleteRegistResponseBody {
	s.Success = &v
	return s
}

type HrbrainDeleteRegistResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainDeleteRegistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainDeleteRegistResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteRegistResponse) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteRegistResponse) SetHeaders(v map[string]*string) *HrbrainDeleteRegistResponse {
	s.Headers = v
	return s
}

func (s *HrbrainDeleteRegistResponse) SetStatusCode(v int32) *HrbrainDeleteRegistResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainDeleteRegistResponse) SetBody(v *HrbrainDeleteRegistResponseBody) *HrbrainDeleteRegistResponse {
	s.Body = v
	return s
}

type HrbrainDeleteRegularHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainDeleteRegularHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteRegularHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteRegularHeaders) SetCommonHeaders(v map[string]*string) *HrbrainDeleteRegularHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainDeleteRegularHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainDeleteRegularHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainDeleteRegularRequest struct {
	Params []*HrbrainDeleteRegularRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Repeated"`
}

func (s HrbrainDeleteRegularRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteRegularRequest) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteRegularRequest) SetParams(v []*HrbrainDeleteRegularRequestParams) *HrbrainDeleteRegularRequest {
	s.Params = v
	return s
}

type HrbrainDeleteRegularRequestParams struct {
	// This parameter is required.
	RegularDate *string `json:"regularDate,omitempty" xml:"regularDate,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainDeleteRegularRequestParams) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteRegularRequestParams) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteRegularRequestParams) SetRegularDate(v string) *HrbrainDeleteRegularRequestParams {
	s.RegularDate = &v
	return s
}

func (s *HrbrainDeleteRegularRequestParams) SetWorkNo(v string) *HrbrainDeleteRegularRequestParams {
	s.WorkNo = &v
	return s
}

type HrbrainDeleteRegularResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainDeleteRegularResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteRegularResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteRegularResponseBody) SetRequestId(v string) *HrbrainDeleteRegularResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainDeleteRegularResponseBody) SetResult(v bool) *HrbrainDeleteRegularResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainDeleteRegularResponseBody) SetSuccess(v bool) *HrbrainDeleteRegularResponseBody {
	s.Success = &v
	return s
}

type HrbrainDeleteRegularResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainDeleteRegularResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainDeleteRegularResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteRegularResponse) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteRegularResponse) SetHeaders(v map[string]*string) *HrbrainDeleteRegularResponse {
	s.Headers = v
	return s
}

func (s *HrbrainDeleteRegularResponse) SetStatusCode(v int32) *HrbrainDeleteRegularResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainDeleteRegularResponse) SetBody(v *HrbrainDeleteRegularResponseBody) *HrbrainDeleteRegularResponse {
	s.Body = v
	return s
}

type HrbrainDeleteTrainingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainDeleteTrainingHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteTrainingHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteTrainingHeaders) SetCommonHeaders(v map[string]*string) *HrbrainDeleteTrainingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainDeleteTrainingHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainDeleteTrainingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainDeleteTrainingRequest struct {
	Params []*HrbrainDeleteTrainingRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Repeated"`
}

func (s HrbrainDeleteTrainingRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteTrainingRequest) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteTrainingRequest) SetParams(v []*HrbrainDeleteTrainingRequestParams) *HrbrainDeleteTrainingRequest {
	s.Params = v
	return s
}

type HrbrainDeleteTrainingRequestParams struct {
	TrainEndDate   *string `json:"trainEndDate,omitempty" xml:"trainEndDate,omitempty"`
	TrainName      *string `json:"trainName,omitempty" xml:"trainName,omitempty"`
	TrainStartDate *string `json:"trainStartDate,omitempty" xml:"trainStartDate,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainDeleteTrainingRequestParams) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteTrainingRequestParams) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteTrainingRequestParams) SetTrainEndDate(v string) *HrbrainDeleteTrainingRequestParams {
	s.TrainEndDate = &v
	return s
}

func (s *HrbrainDeleteTrainingRequestParams) SetTrainName(v string) *HrbrainDeleteTrainingRequestParams {
	s.TrainName = &v
	return s
}

func (s *HrbrainDeleteTrainingRequestParams) SetTrainStartDate(v string) *HrbrainDeleteTrainingRequestParams {
	s.TrainStartDate = &v
	return s
}

func (s *HrbrainDeleteTrainingRequestParams) SetWorkNo(v string) *HrbrainDeleteTrainingRequestParams {
	s.WorkNo = &v
	return s
}

type HrbrainDeleteTrainingResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainDeleteTrainingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteTrainingResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteTrainingResponseBody) SetRequestId(v string) *HrbrainDeleteTrainingResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainDeleteTrainingResponseBody) SetResult(v bool) *HrbrainDeleteTrainingResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainDeleteTrainingResponseBody) SetSuccess(v bool) *HrbrainDeleteTrainingResponseBody {
	s.Success = &v
	return s
}

type HrbrainDeleteTrainingResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainDeleteTrainingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainDeleteTrainingResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteTrainingResponse) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteTrainingResponse) SetHeaders(v map[string]*string) *HrbrainDeleteTrainingResponse {
	s.Headers = v
	return s
}

func (s *HrbrainDeleteTrainingResponse) SetStatusCode(v int32) *HrbrainDeleteTrainingResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainDeleteTrainingResponse) SetBody(v *HrbrainDeleteTrainingResponseBody) *HrbrainDeleteTrainingResponse {
	s.Body = v
	return s
}

type HrbrainDeleteTransferEvalHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainDeleteTransferEvalHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteTransferEvalHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteTransferEvalHeaders) SetCommonHeaders(v map[string]*string) *HrbrainDeleteTransferEvalHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainDeleteTransferEvalHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainDeleteTransferEvalHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainDeleteTransferEvalRequest struct {
	Params []*HrbrainDeleteTransferEvalRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Repeated"`
}

func (s HrbrainDeleteTransferEvalRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteTransferEvalRequest) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteTransferEvalRequest) SetParams(v []*HrbrainDeleteTransferEvalRequestParams) *HrbrainDeleteTransferEvalRequest {
	s.Params = v
	return s
}

type HrbrainDeleteTransferEvalRequestParams struct {
	TransferDate *string `json:"transferDate,omitempty" xml:"transferDate,omitempty"`
	TransferType *string `json:"transferType,omitempty" xml:"transferType,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainDeleteTransferEvalRequestParams) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteTransferEvalRequestParams) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteTransferEvalRequestParams) SetTransferDate(v string) *HrbrainDeleteTransferEvalRequestParams {
	s.TransferDate = &v
	return s
}

func (s *HrbrainDeleteTransferEvalRequestParams) SetTransferType(v string) *HrbrainDeleteTransferEvalRequestParams {
	s.TransferType = &v
	return s
}

func (s *HrbrainDeleteTransferEvalRequestParams) SetWorkNo(v string) *HrbrainDeleteTransferEvalRequestParams {
	s.WorkNo = &v
	return s
}

type HrbrainDeleteTransferEvalResponseBody struct {
	// example:
	//
	// 480021443f9f37fcbf464c4a6b85d289
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainDeleteTransferEvalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteTransferEvalResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteTransferEvalResponseBody) SetRequestId(v string) *HrbrainDeleteTransferEvalResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainDeleteTransferEvalResponseBody) SetResult(v bool) *HrbrainDeleteTransferEvalResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainDeleteTransferEvalResponseBody) SetSuccess(v bool) *HrbrainDeleteTransferEvalResponseBody {
	s.Success = &v
	return s
}

type HrbrainDeleteTransferEvalResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainDeleteTransferEvalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainDeleteTransferEvalResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteTransferEvalResponse) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteTransferEvalResponse) SetHeaders(v map[string]*string) *HrbrainDeleteTransferEvalResponse {
	s.Headers = v
	return s
}

func (s *HrbrainDeleteTransferEvalResponse) SetStatusCode(v int32) *HrbrainDeleteTransferEvalResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainDeleteTransferEvalResponse) SetBody(v *HrbrainDeleteTransferEvalResponseBody) *HrbrainDeleteTransferEvalResponse {
	s.Body = v
	return s
}

type HrbrainDeleteWorkExpHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainDeleteWorkExpHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteWorkExpHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteWorkExpHeaders) SetCommonHeaders(v map[string]*string) *HrbrainDeleteWorkExpHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainDeleteWorkExpHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainDeleteWorkExpHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainDeleteWorkExpRequest struct {
	Params []*HrbrainDeleteWorkExpRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Repeated"`
}

func (s HrbrainDeleteWorkExpRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteWorkExpRequest) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteWorkExpRequest) SetParams(v []*HrbrainDeleteWorkExpRequestParams) *HrbrainDeleteWorkExpRequest {
	s.Params = v
	return s
}

type HrbrainDeleteWorkExpRequestParams struct {
	CompanyName *string `json:"companyName,omitempty" xml:"companyName,omitempty"`
	EndDate     *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	StartDate   *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainDeleteWorkExpRequestParams) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteWorkExpRequestParams) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteWorkExpRequestParams) SetCompanyName(v string) *HrbrainDeleteWorkExpRequestParams {
	s.CompanyName = &v
	return s
}

func (s *HrbrainDeleteWorkExpRequestParams) SetEndDate(v string) *HrbrainDeleteWorkExpRequestParams {
	s.EndDate = &v
	return s
}

func (s *HrbrainDeleteWorkExpRequestParams) SetStartDate(v string) *HrbrainDeleteWorkExpRequestParams {
	s.StartDate = &v
	return s
}

func (s *HrbrainDeleteWorkExpRequestParams) SetWorkNo(v string) *HrbrainDeleteWorkExpRequestParams {
	s.WorkNo = &v
	return s
}

type HrbrainDeleteWorkExpResponseBody struct {
	// example:
	//
	// 480021443f9f37fcbf464c4a6b85d289
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainDeleteWorkExpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteWorkExpResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteWorkExpResponseBody) SetRequestId(v string) *HrbrainDeleteWorkExpResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainDeleteWorkExpResponseBody) SetResult(v bool) *HrbrainDeleteWorkExpResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainDeleteWorkExpResponseBody) SetSuccess(v bool) *HrbrainDeleteWorkExpResponseBody {
	s.Success = &v
	return s
}

type HrbrainDeleteWorkExpResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainDeleteWorkExpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainDeleteWorkExpResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeleteWorkExpResponse) GoString() string {
	return s.String()
}

func (s *HrbrainDeleteWorkExpResponse) SetHeaders(v map[string]*string) *HrbrainDeleteWorkExpResponse {
	s.Headers = v
	return s
}

func (s *HrbrainDeleteWorkExpResponse) SetStatusCode(v int32) *HrbrainDeleteWorkExpResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainDeleteWorkExpResponse) SetBody(v *HrbrainDeleteWorkExpResponseBody) *HrbrainDeleteWorkExpResponse {
	s.Body = v
	return s
}

type HrbrainDeletetLabelBaseHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainDeletetLabelBaseHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeletetLabelBaseHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainDeletetLabelBaseHeaders) SetCommonHeaders(v map[string]*string) *HrbrainDeletetLabelBaseHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainDeletetLabelBaseHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainDeletetLabelBaseHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainDeletetLabelBaseRequest struct {
	Params []*HrbrainDeletetLabelBaseRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Repeated"`
}

func (s HrbrainDeletetLabelBaseRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeletetLabelBaseRequest) GoString() string {
	return s.String()
}

func (s *HrbrainDeletetLabelBaseRequest) SetParams(v []*HrbrainDeletetLabelBaseRequestParams) *HrbrainDeletetLabelBaseRequest {
	s.Params = v
	return s
}

type HrbrainDeletetLabelBaseRequestParams struct {
	Label map[string]interface{} `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainDeletetLabelBaseRequestParams) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeletetLabelBaseRequestParams) GoString() string {
	return s.String()
}

func (s *HrbrainDeletetLabelBaseRequestParams) SetLabel(v map[string]interface{}) *HrbrainDeletetLabelBaseRequestParams {
	s.Label = v
	return s
}

func (s *HrbrainDeletetLabelBaseRequestParams) SetWorkNo(v string) *HrbrainDeletetLabelBaseRequestParams {
	s.WorkNo = &v
	return s
}

type HrbrainDeletetLabelBaseResponseBody struct {
	// example:
	//
	// 480021443f9f37fcbf464c4a6b85d289
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainDeletetLabelBaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeletetLabelBaseResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainDeletetLabelBaseResponseBody) SetRequestId(v string) *HrbrainDeletetLabelBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainDeletetLabelBaseResponseBody) SetResult(v bool) *HrbrainDeletetLabelBaseResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainDeletetLabelBaseResponseBody) SetSuccess(v bool) *HrbrainDeletetLabelBaseResponseBody {
	s.Success = &v
	return s
}

type HrbrainDeletetLabelBaseResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainDeletetLabelBaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainDeletetLabelBaseResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainDeletetLabelBaseResponse) GoString() string {
	return s.String()
}

func (s *HrbrainDeletetLabelBaseResponse) SetHeaders(v map[string]*string) *HrbrainDeletetLabelBaseResponse {
	s.Headers = v
	return s
}

func (s *HrbrainDeletetLabelBaseResponse) SetStatusCode(v int32) *HrbrainDeletetLabelBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainDeletetLabelBaseResponse) SetBody(v *HrbrainDeletetLabelBaseResponseBody) *HrbrainDeletetLabelBaseResponse {
	s.Body = v
	return s
}

type HrbrainEmpPoolQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainEmpPoolQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainEmpPoolQueryHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainEmpPoolQueryHeaders) SetCommonHeaders(v map[string]*string) *HrbrainEmpPoolQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainEmpPoolQueryHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainEmpPoolQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainEmpPoolQueryRequest struct {
	Keyword    *string   `json:"keyword,omitempty" xml:"keyword,omitempty"`
	Labels     []*string `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	MaxResults *int32    `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int32    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s HrbrainEmpPoolQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainEmpPoolQueryRequest) GoString() string {
	return s.String()
}

func (s *HrbrainEmpPoolQueryRequest) SetKeyword(v string) *HrbrainEmpPoolQueryRequest {
	s.Keyword = &v
	return s
}

func (s *HrbrainEmpPoolQueryRequest) SetLabels(v []*string) *HrbrainEmpPoolQueryRequest {
	s.Labels = v
	return s
}

func (s *HrbrainEmpPoolQueryRequest) SetMaxResults(v int32) *HrbrainEmpPoolQueryRequest {
	s.MaxResults = &v
	return s
}

func (s *HrbrainEmpPoolQueryRequest) SetNextToken(v int32) *HrbrainEmpPoolQueryRequest {
	s.NextToken = &v
	return s
}

type HrbrainEmpPoolQueryResponseBody struct {
	Content   *HrbrainEmpPoolQueryResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	RequestId *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool                                   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool                                   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainEmpPoolQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainEmpPoolQueryResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainEmpPoolQueryResponseBody) SetContent(v *HrbrainEmpPoolQueryResponseBodyContent) *HrbrainEmpPoolQueryResponseBody {
	s.Content = v
	return s
}

func (s *HrbrainEmpPoolQueryResponseBody) SetRequestId(v string) *HrbrainEmpPoolQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainEmpPoolQueryResponseBody) SetResult(v bool) *HrbrainEmpPoolQueryResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainEmpPoolQueryResponseBody) SetSuccess(v bool) *HrbrainEmpPoolQueryResponseBody {
	s.Success = &v
	return s
}

type HrbrainEmpPoolQueryResponseBodyContent struct {
	MaxResults *int32                                             `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int32                                             `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	PoolInfos  []*HrbrainEmpPoolQueryResponseBodyContentPoolInfos `json:"poolInfos,omitempty" xml:"poolInfos,omitempty" type:"Repeated"`
	TotalCount *int32                                             `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s HrbrainEmpPoolQueryResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s HrbrainEmpPoolQueryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *HrbrainEmpPoolQueryResponseBodyContent) SetMaxResults(v int32) *HrbrainEmpPoolQueryResponseBodyContent {
	s.MaxResults = &v
	return s
}

func (s *HrbrainEmpPoolQueryResponseBodyContent) SetNextToken(v int32) *HrbrainEmpPoolQueryResponseBodyContent {
	s.NextToken = &v
	return s
}

func (s *HrbrainEmpPoolQueryResponseBodyContent) SetPoolInfos(v []*HrbrainEmpPoolQueryResponseBodyContentPoolInfos) *HrbrainEmpPoolQueryResponseBodyContent {
	s.PoolInfos = v
	return s
}

func (s *HrbrainEmpPoolQueryResponseBodyContent) SetTotalCount(v int32) *HrbrainEmpPoolQueryResponseBodyContent {
	s.TotalCount = &v
	return s
}

type HrbrainEmpPoolQueryResponseBodyContentPoolInfos struct {
	PoolCode *string                                                    `json:"poolCode,omitempty" xml:"poolCode,omitempty"`
	PoolDesc *string                                                    `json:"poolDesc,omitempty" xml:"poolDesc,omitempty"`
	PoolName *string                                                    `json:"poolName,omitempty" xml:"poolName,omitempty"`
	PoolTags []*HrbrainEmpPoolQueryResponseBodyContentPoolInfosPoolTags `json:"poolTags,omitempty" xml:"poolTags,omitempty" type:"Repeated"`
}

func (s HrbrainEmpPoolQueryResponseBodyContentPoolInfos) String() string {
	return tea.Prettify(s)
}

func (s HrbrainEmpPoolQueryResponseBodyContentPoolInfos) GoString() string {
	return s.String()
}

func (s *HrbrainEmpPoolQueryResponseBodyContentPoolInfos) SetPoolCode(v string) *HrbrainEmpPoolQueryResponseBodyContentPoolInfos {
	s.PoolCode = &v
	return s
}

func (s *HrbrainEmpPoolQueryResponseBodyContentPoolInfos) SetPoolDesc(v string) *HrbrainEmpPoolQueryResponseBodyContentPoolInfos {
	s.PoolDesc = &v
	return s
}

func (s *HrbrainEmpPoolQueryResponseBodyContentPoolInfos) SetPoolName(v string) *HrbrainEmpPoolQueryResponseBodyContentPoolInfos {
	s.PoolName = &v
	return s
}

func (s *HrbrainEmpPoolQueryResponseBodyContentPoolInfos) SetPoolTags(v []*HrbrainEmpPoolQueryResponseBodyContentPoolInfosPoolTags) *HrbrainEmpPoolQueryResponseBodyContentPoolInfos {
	s.PoolTags = v
	return s
}

type HrbrainEmpPoolQueryResponseBodyContentPoolInfosPoolTags struct {
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s HrbrainEmpPoolQueryResponseBodyContentPoolInfosPoolTags) String() string {
	return tea.Prettify(s)
}

func (s HrbrainEmpPoolQueryResponseBodyContentPoolInfosPoolTags) GoString() string {
	return s.String()
}

func (s *HrbrainEmpPoolQueryResponseBodyContentPoolInfosPoolTags) SetLabel(v string) *HrbrainEmpPoolQueryResponseBodyContentPoolInfosPoolTags {
	s.Label = &v
	return s
}

func (s *HrbrainEmpPoolQueryResponseBodyContentPoolInfosPoolTags) SetValue(v string) *HrbrainEmpPoolQueryResponseBodyContentPoolInfosPoolTags {
	s.Value = &v
	return s
}

type HrbrainEmpPoolQueryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainEmpPoolQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainEmpPoolQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainEmpPoolQueryResponse) GoString() string {
	return s.String()
}

func (s *HrbrainEmpPoolQueryResponse) SetHeaders(v map[string]*string) *HrbrainEmpPoolQueryResponse {
	s.Headers = v
	return s
}

func (s *HrbrainEmpPoolQueryResponse) SetStatusCode(v int32) *HrbrainEmpPoolQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainEmpPoolQueryResponse) SetBody(v *HrbrainEmpPoolQueryResponseBody) *HrbrainEmpPoolQueryResponse {
	s.Body = v
	return s
}

type HrbrainEmpPoolUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainEmpPoolUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainEmpPoolUserHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainEmpPoolUserHeaders) SetCommonHeaders(v map[string]*string) *HrbrainEmpPoolUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainEmpPoolUserHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainEmpPoolUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainEmpPoolUserRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int32  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	PoolCode   *string `json:"poolCode,omitempty" xml:"poolCode,omitempty"`
}

func (s HrbrainEmpPoolUserRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainEmpPoolUserRequest) GoString() string {
	return s.String()
}

func (s *HrbrainEmpPoolUserRequest) SetMaxResults(v int32) *HrbrainEmpPoolUserRequest {
	s.MaxResults = &v
	return s
}

func (s *HrbrainEmpPoolUserRequest) SetNextToken(v int32) *HrbrainEmpPoolUserRequest {
	s.NextToken = &v
	return s
}

func (s *HrbrainEmpPoolUserRequest) SetPoolCode(v string) *HrbrainEmpPoolUserRequest {
	s.PoolCode = &v
	return s
}

type HrbrainEmpPoolUserResponseBody struct {
	Content   *HrbrainEmpPoolUserResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool                                  `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainEmpPoolUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainEmpPoolUserResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainEmpPoolUserResponseBody) SetContent(v *HrbrainEmpPoolUserResponseBodyContent) *HrbrainEmpPoolUserResponseBody {
	s.Content = v
	return s
}

func (s *HrbrainEmpPoolUserResponseBody) SetRequestId(v string) *HrbrainEmpPoolUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainEmpPoolUserResponseBody) SetResult(v bool) *HrbrainEmpPoolUserResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainEmpPoolUserResponseBody) SetSuccess(v bool) *HrbrainEmpPoolUserResponseBody {
	s.Success = &v
	return s
}

type HrbrainEmpPoolUserResponseBodyContent struct {
	EmpVos     []*HrbrainEmpPoolUserResponseBodyContentEmpVos `json:"empVos,omitempty" xml:"empVos,omitempty" type:"Repeated"`
	MaxResults *int32                                         `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int32                                         `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	TotalCount *int32                                         `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s HrbrainEmpPoolUserResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s HrbrainEmpPoolUserResponseBodyContent) GoString() string {
	return s.String()
}

func (s *HrbrainEmpPoolUserResponseBodyContent) SetEmpVos(v []*HrbrainEmpPoolUserResponseBodyContentEmpVos) *HrbrainEmpPoolUserResponseBodyContent {
	s.EmpVos = v
	return s
}

func (s *HrbrainEmpPoolUserResponseBodyContent) SetMaxResults(v int32) *HrbrainEmpPoolUserResponseBodyContent {
	s.MaxResults = &v
	return s
}

func (s *HrbrainEmpPoolUserResponseBodyContent) SetNextToken(v int32) *HrbrainEmpPoolUserResponseBodyContent {
	s.NextToken = &v
	return s
}

func (s *HrbrainEmpPoolUserResponseBodyContent) SetTotalCount(v int32) *HrbrainEmpPoolUserResponseBodyContent {
	s.TotalCount = &v
	return s
}

type HrbrainEmpPoolUserResponseBodyContentEmpVos struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s HrbrainEmpPoolUserResponseBodyContentEmpVos) String() string {
	return tea.Prettify(s)
}

func (s HrbrainEmpPoolUserResponseBodyContentEmpVos) GoString() string {
	return s.String()
}

func (s *HrbrainEmpPoolUserResponseBodyContentEmpVos) SetName(v string) *HrbrainEmpPoolUserResponseBodyContentEmpVos {
	s.Name = &v
	return s
}

func (s *HrbrainEmpPoolUserResponseBodyContentEmpVos) SetUserId(v string) *HrbrainEmpPoolUserResponseBodyContentEmpVos {
	s.UserId = &v
	return s
}

type HrbrainEmpPoolUserResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainEmpPoolUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainEmpPoolUserResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainEmpPoolUserResponse) GoString() string {
	return s.String()
}

func (s *HrbrainEmpPoolUserResponse) SetHeaders(v map[string]*string) *HrbrainEmpPoolUserResponse {
	s.Headers = v
	return s
}

func (s *HrbrainEmpPoolUserResponse) SetStatusCode(v int32) *HrbrainEmpPoolUserResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainEmpPoolUserResponse) SetBody(v *HrbrainEmpPoolUserResponseBody) *HrbrainEmpPoolUserResponse {
	s.Body = v
	return s
}

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
	// This parameter is required.
	Body []*HrbrainImportAwardDetailRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// This parameter is required.
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
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
	// This parameter is required.
	AwardDate *string `json:"awardDate,omitempty" xml:"awardDate,omitempty"`
	// This parameter is required.
	AwardName  *string                `json:"awardName,omitempty" xml:"awardName,omitempty"`
	AwardOrg   *string                `json:"awardOrg,omitempty" xml:"awardOrg,omitempty"`
	AwardType  *string                `json:"awardType,omitempty" xml:"awardType,omitempty"`
	Comment    *string                `json:"comment,omitempty" xml:"comment,omitempty"`
	ExtendInfo map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
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

type HrbrainImportCustomHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainImportCustomHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportCustomHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainImportCustomHeaders) SetCommonHeaders(v map[string]*string) *HrbrainImportCustomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainImportCustomHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainImportCustomHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainImportCustomRequest struct {
	// This parameter is required.
	Body []map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// This parameter is required.
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// This parameter is required.
	ModelCode *string `json:"modelCode,omitempty" xml:"modelCode,omitempty"`
}

func (s HrbrainImportCustomRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportCustomRequest) GoString() string {
	return s.String()
}

func (s *HrbrainImportCustomRequest) SetBody(v []map[string]interface{}) *HrbrainImportCustomRequest {
	s.Body = v
	return s
}

func (s *HrbrainImportCustomRequest) SetCorpId(v string) *HrbrainImportCustomRequest {
	s.CorpId = &v
	return s
}

func (s *HrbrainImportCustomRequest) SetModelCode(v string) *HrbrainImportCustomRequest {
	s.ModelCode = &v
	return s
}

type HrbrainImportCustomResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainImportCustomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportCustomResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportCustomResponseBody) SetRequestId(v string) *HrbrainImportCustomResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainImportCustomResponseBody) SetResult(v bool) *HrbrainImportCustomResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainImportCustomResponseBody) SetSuccess(v bool) *HrbrainImportCustomResponseBody {
	s.Success = &v
	return s
}

type HrbrainImportCustomResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainImportCustomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainImportCustomResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportCustomResponse) GoString() string {
	return s.String()
}

func (s *HrbrainImportCustomResponse) SetHeaders(v map[string]*string) *HrbrainImportCustomResponse {
	s.Headers = v
	return s
}

func (s *HrbrainImportCustomResponse) SetStatusCode(v int32) *HrbrainImportCustomResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainImportCustomResponse) SetBody(v *HrbrainImportCustomResponseBody) *HrbrainImportCustomResponse {
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
	// This parameter is required.
	Body []*HrbrainImportDeptInfoRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// This parameter is required.
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
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
	CreateDate *string `json:"createDate,omitempty" xml:"createDate,omitempty"`
	// This parameter is required.
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// This parameter is required.
	DeptNo        *string                `json:"deptNo,omitempty" xml:"deptNo,omitempty"`
	EffectiveDate *string                `json:"effectiveDate,omitempty" xml:"effectiveDate,omitempty"`
	ExtendInfo    map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	IsEffective   *string                `json:"isEffective,omitempty" xml:"isEffective,omitempty"`
	SuperDeptName *string                `json:"superDeptName,omitempty" xml:"superDeptName,omitempty"`
	// This parameter is required.
	SuperDeptNo *string `json:"superDeptNo,omitempty" xml:"superDeptNo,omitempty"`
	SuperEmpId  *string `json:"superEmpId,omitempty" xml:"superEmpId,omitempty"`
	SuperName   *string `json:"superName,omitempty" xml:"superName,omitempty"`
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
	// This parameter is required.
	Body []*HrbrainImportDimissionRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// This parameter is required.
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
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
	// This parameter is required.
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// This parameter is required.
	DeptNo *string `json:"deptNo,omitempty" xml:"deptNo,omitempty"`
	// This parameter is required.
	DimissionDate *string `json:"dimissionDate,omitempty" xml:"dimissionDate,omitempty"`
	// This parameter is required.
	DimissionReaasonDesc *string `json:"dimissionReaasonDesc,omitempty" xml:"dimissionReaasonDesc,omitempty"`
	// This parameter is required.
	DimissionReason *string `json:"dimissionReason,omitempty" xml:"dimissionReason,omitempty"`
	// This parameter is required.
	EmpType     *string                `json:"empType,omitempty" xml:"empType,omitempty"`
	ExtendInfo  map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	JobCodeName *string                `json:"jobCodeName,omitempty" xml:"jobCodeName,omitempty"`
	JobLevel    *string                `json:"jobLevel,omitempty" xml:"jobLevel,omitempty"`
	// This parameter is required.
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	PostName *string `json:"postName,omitempty" xml:"postName,omitempty"`
	// This parameter is required.
	SuperName   *string `json:"superName,omitempty" xml:"superName,omitempty"`
	WorkLocAddr *string `json:"workLocAddr,omitempty" xml:"workLocAddr,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
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
	// This parameter is required.
	Body []*HrbrainImportEduExpRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// This parameter is required.
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
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
	// This parameter is required.
	EduName *string `json:"eduName,omitempty" xml:"eduName,omitempty"`
	// This parameter is required.
	EndDate    *string                `json:"endDate,omitempty" xml:"endDate,omitempty"`
	ExtendInfo map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	Major      *string                `json:"major,omitempty" xml:"major,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	SchoolName *string `json:"schoolName,omitempty" xml:"schoolName,omitempty"`
	// This parameter is required.
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
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
	// This parameter is required.
	Body []*HrbrainImportEmpInfoRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// This parameter is required.
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
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
	// This parameter is required.
	Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// This parameter is required.
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// This parameter is required.
	DeptNo        *string `json:"deptNo,omitempty" xml:"deptNo,omitempty"`
	DimissionDate *string `json:"dimissionDate,omitempty" xml:"dimissionDate,omitempty"`
	// This parameter is required.
	EmpSource *string `json:"empSource,omitempty" xml:"empSource,omitempty"`
	// This parameter is required.
	EmpStatus *string `json:"empStatus,omitempty" xml:"empStatus,omitempty"`
	// This parameter is required.
	EmpType    *string                `json:"empType,omitempty" xml:"empType,omitempty"`
	ExtendInfo map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	// This parameter is required.
	Gender         *string `json:"gender,omitempty" xml:"gender,omitempty"`
	HighestDegree  *string `json:"highestDegree,omitempty" xml:"highestDegree,omitempty"`
	HighestEduName *string `json:"highestEduName,omitempty" xml:"highestEduName,omitempty"`
	IsDimission    *string `json:"isDimission,omitempty" xml:"isDimission,omitempty"`
	// This parameter is required.
	JobCodeName    *string `json:"jobCodeName,omitempty" xml:"jobCodeName,omitempty"`
	JobLevel       *string `json:"jobLevel,omitempty" xml:"jobLevel,omitempty"`
	LastSchoolName *string `json:"lastSchoolName,omitempty" xml:"lastSchoolName,omitempty"`
	Marriage       *string `json:"marriage,omitempty" xml:"marriage,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	Nation *string `json:"nation,omitempty" xml:"nation,omitempty"`
	// This parameter is required.
	NationCtry *string `json:"nationCtry,omitempty" xml:"nationCtry,omitempty"`
	// This parameter is required.
	PoliticalStatus *string `json:"politicalStatus,omitempty" xml:"politicalStatus,omitempty"`
	// This parameter is required.
	PostName    *string `json:"postName,omitempty" xml:"postName,omitempty"`
	RegistDate  *string `json:"registDate,omitempty" xml:"registDate,omitempty"`
	RegularDate *string `json:"regularDate,omitempty" xml:"regularDate,omitempty"`
	SuperEmpId  *string `json:"superEmpId,omitempty" xml:"superEmpId,omitempty"`
	SuperName   *string `json:"superName,omitempty" xml:"superName,omitempty"`
	WorkEmail   *string `json:"workEmail,omitempty" xml:"workEmail,omitempty"`
	WorkLocAddr *string `json:"workLocAddr,omitempty" xml:"workLocAddr,omitempty"`
	WorkLocCity *string `json:"workLocCity,omitempty" xml:"workLocCity,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
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
	// This parameter is required.
	Body []*HrbrainImportLabelBaseRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// This parameter is required.
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
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
	// This parameter is required.
	ExtendInfo map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
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
	// This parameter is required.
	Body []*HrbrainImportLabelCustomRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// This parameter is required.
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
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
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
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
	// This parameter is required.
	Body []*HrbrainImportLabelIndustryRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// This parameter is required.
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
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
	// This parameter is required.
	Level1 *string `json:"level1,omitempty" xml:"level1,omitempty"`
	// This parameter is required.
	Level2 *string `json:"level2,omitempty" xml:"level2,omitempty"`
	// This parameter is required.
	Level3 *string `json:"level3,omitempty" xml:"level3,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
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
	// This parameter is required.
	Body []*HrbrainImportLabelInventoryRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// This parameter is required.
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
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
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
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
	// This parameter is required.
	Body []*HrbrainImportLabelProfSkillRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// This parameter is required.
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
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
	// This parameter is required.
	Level1 *string `json:"level1,omitempty" xml:"level1,omitempty"`
	// This parameter is required.
	Level2 *string `json:"level2,omitempty" xml:"level2,omitempty"`
	// This parameter is required.
	Level3 *string `json:"level3,omitempty" xml:"level3,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
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
	// This parameter is required.
	Body []*HrbrainImportPerfEvalRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// This parameter is required.
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
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
	Comment    *string                `json:"comment,omitempty" xml:"comment,omitempty"`
	ExtendInfo map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	// This parameter is required.
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	PerfCate *string `json:"perfCate,omitempty" xml:"perfCate,omitempty"`
	// This parameter is required.
	PerfPlanName *string `json:"perfPlanName,omitempty" xml:"perfPlanName,omitempty"`
	PerfScore    *string `json:"perfScore,omitempty" xml:"perfScore,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	PeriodEndDate *string `json:"periodEndDate,omitempty" xml:"periodEndDate,omitempty"`
	// This parameter is required.
	PeriodStartDate *string `json:"periodStartDate,omitempty" xml:"periodStartDate,omitempty"`
	// This parameter is required.
	Score *string `json:"score,omitempty" xml:"score,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
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
	// This parameter is required.
	Body []*HrbrainImportPromEvalRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// This parameter is required.
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
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
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// This parameter is required.
	EffectiveDate *string                `json:"effectiveDate,omitempty" xml:"effectiveDate,omitempty"`
	ExtendInfo    map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	NewJobLevel *string `json:"newJobLevel,omitempty" xml:"newJobLevel,omitempty"`
	// This parameter is required.
	Period          *string `json:"period,omitempty" xml:"period,omitempty"`
	PeriodEndDate   *string `json:"periodEndDate,omitempty" xml:"periodEndDate,omitempty"`
	PeriodStartDate *string `json:"periodStartDate,omitempty" xml:"periodStartDate,omitempty"`
	// This parameter is required.
	PreJobLevel *string `json:"preJobLevel,omitempty" xml:"preJobLevel,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
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
	// This parameter is required.
	Body []*HrbrainImportPunDetailRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// This parameter is required.
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
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
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// This parameter is required.
	EffectiveDate *string                `json:"effectiveDate,omitempty" xml:"effectiveDate,omitempty"`
	ExtendInfo    map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	PunName *string `json:"punName,omitempty" xml:"punName,omitempty"`
	PunOrg  *string `json:"punOrg,omitempty" xml:"punOrg,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
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
	// This parameter is required.
	Body []*HrbrainImportRegistRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// This parameter is required.
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
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
	// This parameter is required.
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// This parameter is required.
	DeptNo *string `json:"deptNo,omitempty" xml:"deptNo,omitempty"`
	// This parameter is required.
	EmpSource *string `json:"empSource,omitempty" xml:"empSource,omitempty"`
	// This parameter is required.
	EmpType     *string                `json:"empType,omitempty" xml:"empType,omitempty"`
	ExtendInfo  map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	JobCodeName *string                `json:"jobCodeName,omitempty" xml:"jobCodeName,omitempty"`
	JobLevel    *string                `json:"jobLevel,omitempty" xml:"jobLevel,omitempty"`
	// This parameter is required.
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	PostName *string `json:"postName,omitempty" xml:"postName,omitempty"`
	// This parameter is required.
	RegistDate  *string `json:"registDate,omitempty" xml:"registDate,omitempty"`
	SuperName   *string `json:"superName,omitempty" xml:"superName,omitempty"`
	WorkLocAddr *string `json:"workLocAddr,omitempty" xml:"workLocAddr,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
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

type HrbrainImportRegularHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainImportRegularHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportRegularHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainImportRegularHeaders) SetCommonHeaders(v map[string]*string) *HrbrainImportRegularHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainImportRegularHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainImportRegularHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainImportRegularRequest struct {
	Body []*HrbrainImportRegularRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// This parameter is required.
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s HrbrainImportRegularRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportRegularRequest) GoString() string {
	return s.String()
}

func (s *HrbrainImportRegularRequest) SetBody(v []*HrbrainImportRegularRequestBody) *HrbrainImportRegularRequest {
	s.Body = v
	return s
}

func (s *HrbrainImportRegularRequest) SetCorpId(v string) *HrbrainImportRegularRequest {
	s.CorpId = &v
	return s
}

type HrbrainImportRegularRequestBody struct {
	DeptName        *string                `json:"deptName,omitempty" xml:"deptName,omitempty"`
	DeptNo          *string                `json:"deptNo,omitempty" xml:"deptNo,omitempty"`
	ExtendInfo      map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	JobCodeName     *string                `json:"jobCodeName,omitempty" xml:"jobCodeName,omitempty"`
	JobLevel        *string                `json:"jobLevel,omitempty" xml:"jobLevel,omitempty"`
	Name            *string                `json:"name,omitempty" xml:"name,omitempty"`
	PlanRegularDate *string                `json:"planRegularDate,omitempty" xml:"planRegularDate,omitempty"`
	PostName        *string                `json:"postName,omitempty" xml:"postName,omitempty"`
	// This parameter is required.
	RegularDate *string `json:"regularDate,omitempty" xml:"regularDate,omitempty"`
	SuperEmpId  *string `json:"superEmpId,omitempty" xml:"superEmpId,omitempty"`
	SuperName   *string `json:"superName,omitempty" xml:"superName,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainImportRegularRequestBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportRegularRequestBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportRegularRequestBody) SetDeptName(v string) *HrbrainImportRegularRequestBody {
	s.DeptName = &v
	return s
}

func (s *HrbrainImportRegularRequestBody) SetDeptNo(v string) *HrbrainImportRegularRequestBody {
	s.DeptNo = &v
	return s
}

func (s *HrbrainImportRegularRequestBody) SetExtendInfo(v map[string]interface{}) *HrbrainImportRegularRequestBody {
	s.ExtendInfo = v
	return s
}

func (s *HrbrainImportRegularRequestBody) SetJobCodeName(v string) *HrbrainImportRegularRequestBody {
	s.JobCodeName = &v
	return s
}

func (s *HrbrainImportRegularRequestBody) SetJobLevel(v string) *HrbrainImportRegularRequestBody {
	s.JobLevel = &v
	return s
}

func (s *HrbrainImportRegularRequestBody) SetName(v string) *HrbrainImportRegularRequestBody {
	s.Name = &v
	return s
}

func (s *HrbrainImportRegularRequestBody) SetPlanRegularDate(v string) *HrbrainImportRegularRequestBody {
	s.PlanRegularDate = &v
	return s
}

func (s *HrbrainImportRegularRequestBody) SetPostName(v string) *HrbrainImportRegularRequestBody {
	s.PostName = &v
	return s
}

func (s *HrbrainImportRegularRequestBody) SetRegularDate(v string) *HrbrainImportRegularRequestBody {
	s.RegularDate = &v
	return s
}

func (s *HrbrainImportRegularRequestBody) SetSuperEmpId(v string) *HrbrainImportRegularRequestBody {
	s.SuperEmpId = &v
	return s
}

func (s *HrbrainImportRegularRequestBody) SetSuperName(v string) *HrbrainImportRegularRequestBody {
	s.SuperName = &v
	return s
}

func (s *HrbrainImportRegularRequestBody) SetWorkNo(v string) *HrbrainImportRegularRequestBody {
	s.WorkNo = &v
	return s
}

type HrbrainImportRegularResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainImportRegularResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportRegularResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportRegularResponseBody) SetRequestId(v string) *HrbrainImportRegularResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainImportRegularResponseBody) SetResult(v bool) *HrbrainImportRegularResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainImportRegularResponseBody) SetSuccess(v bool) *HrbrainImportRegularResponseBody {
	s.Success = &v
	return s
}

type HrbrainImportRegularResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainImportRegularResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainImportRegularResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportRegularResponse) GoString() string {
	return s.String()
}

func (s *HrbrainImportRegularResponse) SetHeaders(v map[string]*string) *HrbrainImportRegularResponse {
	s.Headers = v
	return s
}

func (s *HrbrainImportRegularResponse) SetStatusCode(v int32) *HrbrainImportRegularResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainImportRegularResponse) SetBody(v *HrbrainImportRegularResponseBody) *HrbrainImportRegularResponse {
	s.Body = v
	return s
}

type HrbrainImportTrainingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainImportTrainingHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportTrainingHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainImportTrainingHeaders) SetCommonHeaders(v map[string]*string) *HrbrainImportTrainingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainImportTrainingHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainImportTrainingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainImportTrainingRequest struct {
	Body []*HrbrainImportTrainingRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// This parameter is required.
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s HrbrainImportTrainingRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportTrainingRequest) GoString() string {
	return s.String()
}

func (s *HrbrainImportTrainingRequest) SetBody(v []*HrbrainImportTrainingRequestBody) *HrbrainImportTrainingRequest {
	s.Body = v
	return s
}

func (s *HrbrainImportTrainingRequest) SetCorpId(v string) *HrbrainImportTrainingRequest {
	s.CorpId = &v
	return s
}

type HrbrainImportTrainingRequestBody struct {
	CertifCnt   *string                `json:"certifCnt,omitempty" xml:"certifCnt,omitempty"`
	CreditScore *string                `json:"creditScore,omitempty" xml:"creditScore,omitempty"`
	DeptName    *string                `json:"deptName,omitempty" xml:"deptName,omitempty"`
	DeptNo      *string                `json:"deptNo,omitempty" xml:"deptNo,omitempty"`
	ExtendInfo  map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	JobCodeName *string                `json:"jobCodeName,omitempty" xml:"jobCodeName,omitempty"`
	JobLevel    *string                `json:"jobLevel,omitempty" xml:"jobLevel,omitempty"`
	Name        *string                `json:"name,omitempty" xml:"name,omitempty"`
	PostName    *string                `json:"postName,omitempty" xml:"postName,omitempty"`
	// This parameter is required.
	TrainEndDate *string `json:"trainEndDate,omitempty" xml:"trainEndDate,omitempty"`
	// This parameter is required.
	TrainName *string `json:"trainName,omitempty" xml:"trainName,omitempty"`
	// This parameter is required.
	TrainStartDate *string `json:"trainStartDate,omitempty" xml:"trainStartDate,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainImportTrainingRequestBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportTrainingRequestBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportTrainingRequestBody) SetCertifCnt(v string) *HrbrainImportTrainingRequestBody {
	s.CertifCnt = &v
	return s
}

func (s *HrbrainImportTrainingRequestBody) SetCreditScore(v string) *HrbrainImportTrainingRequestBody {
	s.CreditScore = &v
	return s
}

func (s *HrbrainImportTrainingRequestBody) SetDeptName(v string) *HrbrainImportTrainingRequestBody {
	s.DeptName = &v
	return s
}

func (s *HrbrainImportTrainingRequestBody) SetDeptNo(v string) *HrbrainImportTrainingRequestBody {
	s.DeptNo = &v
	return s
}

func (s *HrbrainImportTrainingRequestBody) SetExtendInfo(v map[string]interface{}) *HrbrainImportTrainingRequestBody {
	s.ExtendInfo = v
	return s
}

func (s *HrbrainImportTrainingRequestBody) SetJobCodeName(v string) *HrbrainImportTrainingRequestBody {
	s.JobCodeName = &v
	return s
}

func (s *HrbrainImportTrainingRequestBody) SetJobLevel(v string) *HrbrainImportTrainingRequestBody {
	s.JobLevel = &v
	return s
}

func (s *HrbrainImportTrainingRequestBody) SetName(v string) *HrbrainImportTrainingRequestBody {
	s.Name = &v
	return s
}

func (s *HrbrainImportTrainingRequestBody) SetPostName(v string) *HrbrainImportTrainingRequestBody {
	s.PostName = &v
	return s
}

func (s *HrbrainImportTrainingRequestBody) SetTrainEndDate(v string) *HrbrainImportTrainingRequestBody {
	s.TrainEndDate = &v
	return s
}

func (s *HrbrainImportTrainingRequestBody) SetTrainName(v string) *HrbrainImportTrainingRequestBody {
	s.TrainName = &v
	return s
}

func (s *HrbrainImportTrainingRequestBody) SetTrainStartDate(v string) *HrbrainImportTrainingRequestBody {
	s.TrainStartDate = &v
	return s
}

func (s *HrbrainImportTrainingRequestBody) SetWorkNo(v string) *HrbrainImportTrainingRequestBody {
	s.WorkNo = &v
	return s
}

type HrbrainImportTrainingResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainImportTrainingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportTrainingResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainImportTrainingResponseBody) SetRequestId(v string) *HrbrainImportTrainingResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainImportTrainingResponseBody) SetResult(v bool) *HrbrainImportTrainingResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainImportTrainingResponseBody) SetSuccess(v bool) *HrbrainImportTrainingResponseBody {
	s.Success = &v
	return s
}

type HrbrainImportTrainingResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainImportTrainingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainImportTrainingResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainImportTrainingResponse) GoString() string {
	return s.String()
}

func (s *HrbrainImportTrainingResponse) SetHeaders(v map[string]*string) *HrbrainImportTrainingResponse {
	s.Headers = v
	return s
}

func (s *HrbrainImportTrainingResponse) SetStatusCode(v int32) *HrbrainImportTrainingResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainImportTrainingResponse) SetBody(v *HrbrainImportTrainingResponseBody) *HrbrainImportTrainingResponse {
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
	// This parameter is required.
	Body []*HrbrainImportTransferEvalRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// This parameter is required.
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
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
	// This parameter is required.
	CurrInfo   map[string]interface{} `json:"currInfo,omitempty" xml:"currInfo,omitempty"`
	ExtendInfo map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	PreInfo map[string]interface{} `json:"preInfo,omitempty" xml:"preInfo,omitempty"`
	// This parameter is required.
	TransferDate *string `json:"transferDate,omitempty" xml:"transferDate,omitempty"`
	// This parameter is required.
	TransferReason *string `json:"transferReason,omitempty" xml:"transferReason,omitempty"`
	// This parameter is required.
	TransferType *string `json:"transferType,omitempty" xml:"transferType,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
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
	// This parameter is required.
	Body []*HrbrainImportWorkExpRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// This parameter is required.
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
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
	// This parameter is required.
	CompanyName *string `json:"companyName,omitempty" xml:"companyName,omitempty"`
	// This parameter is required.
	EndDate    *string                `json:"endDate,omitempty" xml:"endDate,omitempty"`
	ExtendInfo map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	JobDesc    *string                `json:"jobDesc,omitempty" xml:"jobDesc,omitempty"`
	// This parameter is required.
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	PostName *string `json:"postName,omitempty" xml:"postName,omitempty"`
	// This parameter is required.
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// This parameter is required.
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
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

type HrbrainTalentProfileAttachmentQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainTalentProfileAttachmentQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainTalentProfileAttachmentQueryHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainTalentProfileAttachmentQueryHeaders) SetCommonHeaders(v map[string]*string) *HrbrainTalentProfileAttachmentQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainTalentProfileAttachmentQueryHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainTalentProfileAttachmentQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainTalentProfileAttachmentQueryRequest struct {
	Body       []*string `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	DingCorpId *string   `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
}

func (s HrbrainTalentProfileAttachmentQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainTalentProfileAttachmentQueryRequest) GoString() string {
	return s.String()
}

func (s *HrbrainTalentProfileAttachmentQueryRequest) SetBody(v []*string) *HrbrainTalentProfileAttachmentQueryRequest {
	s.Body = v
	return s
}

func (s *HrbrainTalentProfileAttachmentQueryRequest) SetDingCorpId(v string) *HrbrainTalentProfileAttachmentQueryRequest {
	s.DingCorpId = &v
	return s
}

type HrbrainTalentProfileAttachmentQueryResponseBody struct {
	Content   *HrbrainTalentProfileAttachmentQueryResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	RequestId *string                                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool                                                   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool                                                   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainTalentProfileAttachmentQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainTalentProfileAttachmentQueryResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainTalentProfileAttachmentQueryResponseBody) SetContent(v *HrbrainTalentProfileAttachmentQueryResponseBodyContent) *HrbrainTalentProfileAttachmentQueryResponseBody {
	s.Content = v
	return s
}

func (s *HrbrainTalentProfileAttachmentQueryResponseBody) SetRequestId(v string) *HrbrainTalentProfileAttachmentQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainTalentProfileAttachmentQueryResponseBody) SetResult(v bool) *HrbrainTalentProfileAttachmentQueryResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainTalentProfileAttachmentQueryResponseBody) SetSuccess(v bool) *HrbrainTalentProfileAttachmentQueryResponseBody {
	s.Success = &v
	return s
}

type HrbrainTalentProfileAttachmentQueryResponseBodyContent struct {
	StaffAttachmentInfoList []*HrbrainTalentProfileAttachmentQueryResponseBodyContentStaffAttachmentInfoList `json:"staffAttachmentInfoList,omitempty" xml:"staffAttachmentInfoList,omitempty" type:"Repeated"`
}

func (s HrbrainTalentProfileAttachmentQueryResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s HrbrainTalentProfileAttachmentQueryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *HrbrainTalentProfileAttachmentQueryResponseBodyContent) SetStaffAttachmentInfoList(v []*HrbrainTalentProfileAttachmentQueryResponseBodyContentStaffAttachmentInfoList) *HrbrainTalentProfileAttachmentQueryResponseBodyContent {
	s.StaffAttachmentInfoList = v
	return s
}

type HrbrainTalentProfileAttachmentQueryResponseBodyContentStaffAttachmentInfoList struct {
	AttachmentInfoList []*HrbrainTalentProfileAttachmentQueryResponseBodyContentStaffAttachmentInfoListAttachmentInfoList `json:"attachmentInfoList,omitempty" xml:"attachmentInfoList,omitempty" type:"Repeated"`
	WorkNo             *string                                                                                            `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s HrbrainTalentProfileAttachmentQueryResponseBodyContentStaffAttachmentInfoList) String() string {
	return tea.Prettify(s)
}

func (s HrbrainTalentProfileAttachmentQueryResponseBodyContentStaffAttachmentInfoList) GoString() string {
	return s.String()
}

func (s *HrbrainTalentProfileAttachmentQueryResponseBodyContentStaffAttachmentInfoList) SetAttachmentInfoList(v []*HrbrainTalentProfileAttachmentQueryResponseBodyContentStaffAttachmentInfoListAttachmentInfoList) *HrbrainTalentProfileAttachmentQueryResponseBodyContentStaffAttachmentInfoList {
	s.AttachmentInfoList = v
	return s
}

func (s *HrbrainTalentProfileAttachmentQueryResponseBodyContentStaffAttachmentInfoList) SetWorkNo(v string) *HrbrainTalentProfileAttachmentQueryResponseBodyContentStaffAttachmentInfoList {
	s.WorkNo = &v
	return s
}

type HrbrainTalentProfileAttachmentQueryResponseBodyContentStaffAttachmentInfoListAttachmentInfoList struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Url  *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s HrbrainTalentProfileAttachmentQueryResponseBodyContentStaffAttachmentInfoListAttachmentInfoList) String() string {
	return tea.Prettify(s)
}

func (s HrbrainTalentProfileAttachmentQueryResponseBodyContentStaffAttachmentInfoListAttachmentInfoList) GoString() string {
	return s.String()
}

func (s *HrbrainTalentProfileAttachmentQueryResponseBodyContentStaffAttachmentInfoListAttachmentInfoList) SetName(v string) *HrbrainTalentProfileAttachmentQueryResponseBodyContentStaffAttachmentInfoListAttachmentInfoList {
	s.Name = &v
	return s
}

func (s *HrbrainTalentProfileAttachmentQueryResponseBodyContentStaffAttachmentInfoListAttachmentInfoList) SetUrl(v string) *HrbrainTalentProfileAttachmentQueryResponseBodyContentStaffAttachmentInfoListAttachmentInfoList {
	s.Url = &v
	return s
}

type HrbrainTalentProfileAttachmentQueryResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainTalentProfileAttachmentQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainTalentProfileAttachmentQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainTalentProfileAttachmentQueryResponse) GoString() string {
	return s.String()
}

func (s *HrbrainTalentProfileAttachmentQueryResponse) SetHeaders(v map[string]*string) *HrbrainTalentProfileAttachmentQueryResponse {
	s.Headers = v
	return s
}

func (s *HrbrainTalentProfileAttachmentQueryResponse) SetStatusCode(v int32) *HrbrainTalentProfileAttachmentQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainTalentProfileAttachmentQueryResponse) SetBody(v *HrbrainTalentProfileAttachmentQueryResponseBody) *HrbrainTalentProfileAttachmentQueryResponse {
	s.Body = v
	return s
}

type HrbrainTalentProfileBasicQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrbrainTalentProfileBasicQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrbrainTalentProfileBasicQueryHeaders) GoString() string {
	return s.String()
}

func (s *HrbrainTalentProfileBasicQueryHeaders) SetCommonHeaders(v map[string]*string) *HrbrainTalentProfileBasicQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrbrainTalentProfileBasicQueryHeaders) SetXAcsDingtalkAccessToken(v string) *HrbrainTalentProfileBasicQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrbrainTalentProfileBasicQueryRequest struct {
	Body       []*string `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	DingCorpId *string   `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
}

func (s HrbrainTalentProfileBasicQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s HrbrainTalentProfileBasicQueryRequest) GoString() string {
	return s.String()
}

func (s *HrbrainTalentProfileBasicQueryRequest) SetBody(v []*string) *HrbrainTalentProfileBasicQueryRequest {
	s.Body = v
	return s
}

func (s *HrbrainTalentProfileBasicQueryRequest) SetDingCorpId(v string) *HrbrainTalentProfileBasicQueryRequest {
	s.DingCorpId = &v
	return s
}

type HrbrainTalentProfileBasicQueryResponseBody struct {
	Content   *HrbrainTalentProfileBasicQueryResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	RequestId *string                                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool                                              `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool                                              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HrbrainTalentProfileBasicQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrbrainTalentProfileBasicQueryResponseBody) GoString() string {
	return s.String()
}

func (s *HrbrainTalentProfileBasicQueryResponseBody) SetContent(v *HrbrainTalentProfileBasicQueryResponseBodyContent) *HrbrainTalentProfileBasicQueryResponseBody {
	s.Content = v
	return s
}

func (s *HrbrainTalentProfileBasicQueryResponseBody) SetRequestId(v string) *HrbrainTalentProfileBasicQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *HrbrainTalentProfileBasicQueryResponseBody) SetResult(v bool) *HrbrainTalentProfileBasicQueryResponseBody {
	s.Result = &v
	return s
}

func (s *HrbrainTalentProfileBasicQueryResponseBody) SetSuccess(v bool) *HrbrainTalentProfileBasicQueryResponseBody {
	s.Success = &v
	return s
}

type HrbrainTalentProfileBasicQueryResponseBodyContent struct {
	ProfileBaseInfoList []*HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList `json:"profileBaseInfoList,omitempty" xml:"profileBaseInfoList,omitempty" type:"Repeated"`
}

func (s HrbrainTalentProfileBasicQueryResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s HrbrainTalentProfileBasicQueryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *HrbrainTalentProfileBasicQueryResponseBodyContent) SetProfileBaseInfoList(v []*HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList) *HrbrainTalentProfileBasicQueryResponseBodyContent {
	s.ProfileBaseInfoList = v
	return s
}

type HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList struct {
	Age            *string `json:"age,omitempty" xml:"age,omitempty"`
	Birthday       *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	DeptName       *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	DeptNo         *string `json:"deptNo,omitempty" xml:"deptNo,omitempty"`
	Gender         *string `json:"gender,omitempty" xml:"gender,omitempty"`
	JobLevel       *string `json:"jobLevel,omitempty" xml:"jobLevel,omitempty"`
	Jobcode        *string `json:"jobcode,omitempty" xml:"jobcode,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	Position       *string `json:"position,omitempty" xml:"position,omitempty"`
	SeniorityYears *string `json:"seniorityYears,omitempty" xml:"seniorityYears,omitempty"`
	SuperName      *string `json:"superName,omitempty" xml:"superName,omitempty"`
	SuperWorkNo    *string `json:"superWorkNo,omitempty" xml:"superWorkNo,omitempty"`
	WorkNo         *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
	WorkPlace      *string `json:"workPlace,omitempty" xml:"workPlace,omitempty"`
}

func (s HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList) String() string {
	return tea.Prettify(s)
}

func (s HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList) GoString() string {
	return s.String()
}

func (s *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList) SetAge(v string) *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList {
	s.Age = &v
	return s
}

func (s *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList) SetBirthday(v string) *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList {
	s.Birthday = &v
	return s
}

func (s *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList) SetDeptName(v string) *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList {
	s.DeptName = &v
	return s
}

func (s *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList) SetDeptNo(v string) *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList {
	s.DeptNo = &v
	return s
}

func (s *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList) SetGender(v string) *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList {
	s.Gender = &v
	return s
}

func (s *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList) SetJobLevel(v string) *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList {
	s.JobLevel = &v
	return s
}

func (s *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList) SetJobcode(v string) *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList {
	s.Jobcode = &v
	return s
}

func (s *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList) SetName(v string) *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList {
	s.Name = &v
	return s
}

func (s *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList) SetPosition(v string) *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList {
	s.Position = &v
	return s
}

func (s *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList) SetSeniorityYears(v string) *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList {
	s.SeniorityYears = &v
	return s
}

func (s *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList) SetSuperName(v string) *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList {
	s.SuperName = &v
	return s
}

func (s *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList) SetSuperWorkNo(v string) *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList {
	s.SuperWorkNo = &v
	return s
}

func (s *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList) SetWorkNo(v string) *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList {
	s.WorkNo = &v
	return s
}

func (s *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList) SetWorkPlace(v string) *HrbrainTalentProfileBasicQueryResponseBodyContentProfileBaseInfoList {
	s.WorkPlace = &v
	return s
}

type HrbrainTalentProfileBasicQueryResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrbrainTalentProfileBasicQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrbrainTalentProfileBasicQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s HrbrainTalentProfileBasicQueryResponse) GoString() string {
	return s.String()
}

func (s *HrbrainTalentProfileBasicQueryResponse) SetHeaders(v map[string]*string) *HrbrainTalentProfileBasicQueryResponse {
	s.Headers = v
	return s
}

func (s *HrbrainTalentProfileBasicQueryResponse) SetStatusCode(v int32) *HrbrainTalentProfileBasicQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *HrbrainTalentProfileBasicQueryResponse) SetBody(v *HrbrainTalentProfileBasicQueryResponseBody) *HrbrainTalentProfileBasicQueryResponse {
	s.Body = v
	return s
}

type StaffLabelRecordsQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s StaffLabelRecordsQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s StaffLabelRecordsQueryHeaders) GoString() string {
	return s.String()
}

func (s *StaffLabelRecordsQueryHeaders) SetCommonHeaders(v map[string]*string) *StaffLabelRecordsQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StaffLabelRecordsQueryHeaders) SetXAcsDingtalkAccessToken(v string) *StaffLabelRecordsQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type StaffLabelRecordsQueryRequest struct {
	Body []*StaffLabelRecordsQueryRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// example:
	//
	// 0140180438261064274667
	DingCorpId *string `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s StaffLabelRecordsQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s StaffLabelRecordsQueryRequest) GoString() string {
	return s.String()
}

func (s *StaffLabelRecordsQueryRequest) SetBody(v []*StaffLabelRecordsQueryRequestBody) *StaffLabelRecordsQueryRequest {
	s.Body = v
	return s
}

func (s *StaffLabelRecordsQueryRequest) SetDingCorpId(v string) *StaffLabelRecordsQueryRequest {
	s.DingCorpId = &v
	return s
}

func (s *StaffLabelRecordsQueryRequest) SetMaxResults(v int64) *StaffLabelRecordsQueryRequest {
	s.MaxResults = &v
	return s
}

func (s *StaffLabelRecordsQueryRequest) SetNextToken(v string) *StaffLabelRecordsQueryRequest {
	s.NextToken = &v
	return s
}

type StaffLabelRecordsQueryRequestBody struct {
	Labels []*StaffLabelRecordsQueryRequestBodyLabels `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// example:
	//
	// 0140180438261064274667
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s StaffLabelRecordsQueryRequestBody) String() string {
	return tea.Prettify(s)
}

func (s StaffLabelRecordsQueryRequestBody) GoString() string {
	return s.String()
}

func (s *StaffLabelRecordsQueryRequestBody) SetLabels(v []*StaffLabelRecordsQueryRequestBodyLabels) *StaffLabelRecordsQueryRequestBody {
	s.Labels = v
	return s
}

func (s *StaffLabelRecordsQueryRequestBody) SetUserId(v string) *StaffLabelRecordsQueryRequestBody {
	s.UserId = &v
	return s
}

type StaffLabelRecordsQueryRequestBodyLabels struct {
	// example:
	//
	// long_termism_score
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// values
	TypeCode *string `json:"typeCode,omitempty" xml:"typeCode,omitempty"`
}

func (s StaffLabelRecordsQueryRequestBodyLabels) String() string {
	return tea.Prettify(s)
}

func (s StaffLabelRecordsQueryRequestBodyLabels) GoString() string {
	return s.String()
}

func (s *StaffLabelRecordsQueryRequestBodyLabels) SetCode(v string) *StaffLabelRecordsQueryRequestBodyLabels {
	s.Code = &v
	return s
}

func (s *StaffLabelRecordsQueryRequestBodyLabels) SetTypeCode(v string) *StaffLabelRecordsQueryRequestBodyLabels {
	s.TypeCode = &v
	return s
}

type StaffLabelRecordsQueryResponseBody struct {
	Content *StaffLabelRecordsQueryResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	// example:
	//
	// 0140180438261064274667
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StaffLabelRecordsQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StaffLabelRecordsQueryResponseBody) GoString() string {
	return s.String()
}

func (s *StaffLabelRecordsQueryResponseBody) SetContent(v *StaffLabelRecordsQueryResponseBodyContent) *StaffLabelRecordsQueryResponseBody {
	s.Content = v
	return s
}

func (s *StaffLabelRecordsQueryResponseBody) SetRequestId(v string) *StaffLabelRecordsQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *StaffLabelRecordsQueryResponseBody) SetResult(v bool) *StaffLabelRecordsQueryResponseBody {
	s.Result = &v
	return s
}

func (s *StaffLabelRecordsQueryResponseBody) SetSuccess(v bool) *StaffLabelRecordsQueryResponseBody {
	s.Success = &v
	return s
}

type StaffLabelRecordsQueryResponseBodyContent struct {
	Data []*StaffLabelRecordsQueryResponseBodyContentData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 100
	TotalCountt *int64 `json:"totalCountt,omitempty" xml:"totalCountt,omitempty"`
}

func (s StaffLabelRecordsQueryResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s StaffLabelRecordsQueryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *StaffLabelRecordsQueryResponseBodyContent) SetData(v []*StaffLabelRecordsQueryResponseBodyContentData) *StaffLabelRecordsQueryResponseBodyContent {
	s.Data = v
	return s
}

func (s *StaffLabelRecordsQueryResponseBodyContent) SetMaxResults(v int64) *StaffLabelRecordsQueryResponseBodyContent {
	s.MaxResults = &v
	return s
}

func (s *StaffLabelRecordsQueryResponseBodyContent) SetNextToken(v string) *StaffLabelRecordsQueryResponseBodyContent {
	s.NextToken = &v
	return s
}

func (s *StaffLabelRecordsQueryResponseBodyContent) SetTotalCountt(v int64) *StaffLabelRecordsQueryResponseBodyContent {
	s.TotalCountt = &v
	return s
}

type StaffLabelRecordsQueryResponseBodyContentData struct {
	Labels []*StaffLabelRecordsQueryResponseBodyContentDataLabels `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// example:
	//
	// 0140180438261064274667
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s StaffLabelRecordsQueryResponseBodyContentData) String() string {
	return tea.Prettify(s)
}

func (s StaffLabelRecordsQueryResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *StaffLabelRecordsQueryResponseBodyContentData) SetLabels(v []*StaffLabelRecordsQueryResponseBodyContentDataLabels) *StaffLabelRecordsQueryResponseBodyContentData {
	s.Labels = v
	return s
}

func (s *StaffLabelRecordsQueryResponseBodyContentData) SetUserId(v string) *StaffLabelRecordsQueryResponseBodyContentData {
	s.UserId = &v
	return s
}

type StaffLabelRecordsQueryResponseBodyContentDataLabels struct {
	// example:
	//
	// long_termism_score
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// values.long_termism_score
	Guid *string `json:"guid,omitempty" xml:"guid,omitempty"`
	// example:
	//
	// 持续业绩
	Name    *string                                                       `json:"name,omitempty" xml:"name,omitempty"`
	Options []*StaffLabelRecordsQueryResponseBodyContentDataLabelsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// example:
	//
	// values
	TypeCode *string `json:"typeCode,omitempty" xml:"typeCode,omitempty"`
	// example:
	//
	// 价值
	TypeName *string `json:"typeName,omitempty" xml:"typeName,omitempty"`
	// example:
	//
	// 5（总是）
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s StaffLabelRecordsQueryResponseBodyContentDataLabels) String() string {
	return tea.Prettify(s)
}

func (s StaffLabelRecordsQueryResponseBodyContentDataLabels) GoString() string {
	return s.String()
}

func (s *StaffLabelRecordsQueryResponseBodyContentDataLabels) SetCode(v string) *StaffLabelRecordsQueryResponseBodyContentDataLabels {
	s.Code = &v
	return s
}

func (s *StaffLabelRecordsQueryResponseBodyContentDataLabels) SetGuid(v string) *StaffLabelRecordsQueryResponseBodyContentDataLabels {
	s.Guid = &v
	return s
}

func (s *StaffLabelRecordsQueryResponseBodyContentDataLabels) SetName(v string) *StaffLabelRecordsQueryResponseBodyContentDataLabels {
	s.Name = &v
	return s
}

func (s *StaffLabelRecordsQueryResponseBodyContentDataLabels) SetOptions(v []*StaffLabelRecordsQueryResponseBodyContentDataLabelsOptions) *StaffLabelRecordsQueryResponseBodyContentDataLabels {
	s.Options = v
	return s
}

func (s *StaffLabelRecordsQueryResponseBodyContentDataLabels) SetTypeCode(v string) *StaffLabelRecordsQueryResponseBodyContentDataLabels {
	s.TypeCode = &v
	return s
}

func (s *StaffLabelRecordsQueryResponseBodyContentDataLabels) SetTypeName(v string) *StaffLabelRecordsQueryResponseBodyContentDataLabels {
	s.TypeName = &v
	return s
}

func (s *StaffLabelRecordsQueryResponseBodyContentDataLabels) SetValue(v string) *StaffLabelRecordsQueryResponseBodyContentDataLabels {
	s.Value = &v
	return s
}

type StaffLabelRecordsQueryResponseBodyContentDataLabelsOptions struct {
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	Tip   *string `json:"tip,omitempty" xml:"tip,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s StaffLabelRecordsQueryResponseBodyContentDataLabelsOptions) String() string {
	return tea.Prettify(s)
}

func (s StaffLabelRecordsQueryResponseBodyContentDataLabelsOptions) GoString() string {
	return s.String()
}

func (s *StaffLabelRecordsQueryResponseBodyContentDataLabelsOptions) SetLabel(v string) *StaffLabelRecordsQueryResponseBodyContentDataLabelsOptions {
	s.Label = &v
	return s
}

func (s *StaffLabelRecordsQueryResponseBodyContentDataLabelsOptions) SetTip(v string) *StaffLabelRecordsQueryResponseBodyContentDataLabelsOptions {
	s.Tip = &v
	return s
}

func (s *StaffLabelRecordsQueryResponseBodyContentDataLabelsOptions) SetValue(v string) *StaffLabelRecordsQueryResponseBodyContentDataLabelsOptions {
	s.Value = &v
	return s
}

type StaffLabelRecordsQueryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StaffLabelRecordsQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StaffLabelRecordsQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s StaffLabelRecordsQueryResponse) GoString() string {
	return s.String()
}

func (s *StaffLabelRecordsQueryResponse) SetHeaders(v map[string]*string) *StaffLabelRecordsQueryResponse {
	s.Headers = v
	return s
}

func (s *StaffLabelRecordsQueryResponse) SetStatusCode(v int32) *StaffLabelRecordsQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *StaffLabelRecordsQueryResponse) SetBody(v *StaffLabelRecordsQueryResponseBody) *StaffLabelRecordsQueryResponse {
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
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	DataId *string `json:"dataId,omitempty" xml:"dataId,omitempty"`
	// This parameter is required.
	EtlTime *string `json:"etlTime,omitempty" xml:"etlTime,omitempty"`
	// This parameter is required.
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// This parameter is required.
	SchemaId *string `json:"schemaId,omitempty" xml:"schemaId,omitempty"`
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
// 删除奖励记录
//
// @param request - HrbrainDeleteAwardRecordsRequest
//
// @param headers - HrbrainDeleteAwardRecordsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainDeleteAwardRecordsResponse
func (client *Client) HrbrainDeleteAwardRecordsWithOptions(request *HrbrainDeleteAwardRecordsRequest, headers *HrbrainDeleteAwardRecordsHeaders, runtime *util.RuntimeOptions) (_result *HrbrainDeleteAwardRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
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
		Action:      tea.String("HrbrainDeleteAwardRecords"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/awardRecords/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainDeleteAwardRecordsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除奖励记录
//
// @param request - HrbrainDeleteAwardRecordsRequest
//
// @return HrbrainDeleteAwardRecordsResponse
func (client *Client) HrbrainDeleteAwardRecords(request *HrbrainDeleteAwardRecordsRequest) (_result *HrbrainDeleteAwardRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainDeleteAwardRecordsHeaders{}
	_result = &HrbrainDeleteAwardRecordsResponse{}
	_body, _err := client.HrbrainDeleteAwardRecordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除自定义模型记录
//
// @param request - HrbrainDeleteCustomRequest
//
// @param headers - HrbrainDeleteCustomHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainDeleteCustomResponse
func (client *Client) HrbrainDeleteCustomWithOptions(request *HrbrainDeleteCustomRequest, headers *HrbrainDeleteCustomHeaders, runtime *util.RuntimeOptions) (_result *HrbrainDeleteCustomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModelCode)) {
		body["modelCode"] = request.ModelCode
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
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
		Action:      tea.String("HrbrainDeleteCustom"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/customModels/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainDeleteCustomResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除自定义模型记录
//
// @param request - HrbrainDeleteCustomRequest
//
// @return HrbrainDeleteCustomResponse
func (client *Client) HrbrainDeleteCustom(request *HrbrainDeleteCustomRequest) (_result *HrbrainDeleteCustomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainDeleteCustomHeaders{}
	_result = &HrbrainDeleteCustomResponse{}
	_body, _err := client.HrbrainDeleteCustomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除组织架构
//
// @param request - HrbrainDeleteDeptInfoRequest
//
// @param headers - HrbrainDeleteDeptInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainDeleteDeptInfoResponse
func (client *Client) HrbrainDeleteDeptInfoWithOptions(request *HrbrainDeleteDeptInfoRequest, headers *HrbrainDeleteDeptInfoHeaders, runtime *util.RuntimeOptions) (_result *HrbrainDeleteDeptInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
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
		Action:      tea.String("HrbrainDeleteDeptInfo"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/deptInfos/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainDeleteDeptInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除组织架构
//
// @param request - HrbrainDeleteDeptInfoRequest
//
// @return HrbrainDeleteDeptInfoResponse
func (client *Client) HrbrainDeleteDeptInfo(request *HrbrainDeleteDeptInfoRequest) (_result *HrbrainDeleteDeptInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainDeleteDeptInfoHeaders{}
	_result = &HrbrainDeleteDeptInfoResponse{}
	_body, _err := client.HrbrainDeleteDeptInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除离职记录
//
// @param request - HrbrainDeleteDimissionRequest
//
// @param headers - HrbrainDeleteDimissionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainDeleteDimissionResponse
func (client *Client) HrbrainDeleteDimissionWithOptions(request *HrbrainDeleteDimissionRequest, headers *HrbrainDeleteDimissionHeaders, runtime *util.RuntimeOptions) (_result *HrbrainDeleteDimissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
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
		Action:      tea.String("HrbrainDeleteDimission"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/dimissionInfos/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainDeleteDimissionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除离职记录
//
// @param request - HrbrainDeleteDimissionRequest
//
// @return HrbrainDeleteDimissionResponse
func (client *Client) HrbrainDeleteDimission(request *HrbrainDeleteDimissionRequest) (_result *HrbrainDeleteDimissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainDeleteDimissionHeaders{}
	_result = &HrbrainDeleteDimissionResponse{}
	_body, _err := client.HrbrainDeleteDimissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除教育经历
//
// @param request - HrbrainDeleteEduExpRequest
//
// @param headers - HrbrainDeleteEduExpHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainDeleteEduExpResponse
func (client *Client) HrbrainDeleteEduExpWithOptions(request *HrbrainDeleteEduExpRequest, headers *HrbrainDeleteEduExpHeaders, runtime *util.RuntimeOptions) (_result *HrbrainDeleteEduExpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
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
		Action:      tea.String("HrbrainDeleteEduExp"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/eduExperiences/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainDeleteEduExpResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除教育经历
//
// @param request - HrbrainDeleteEduExpRequest
//
// @return HrbrainDeleteEduExpResponse
func (client *Client) HrbrainDeleteEduExp(request *HrbrainDeleteEduExpRequest) (_result *HrbrainDeleteEduExpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainDeleteEduExpHeaders{}
	_result = &HrbrainDeleteEduExpResponse{}
	_body, _err := client.HrbrainDeleteEduExpWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除人员信息
//
// @param request - HrbrainDeleteEmpInfoRequest
//
// @param headers - HrbrainDeleteEmpInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainDeleteEmpInfoResponse
func (client *Client) HrbrainDeleteEmpInfoWithOptions(request *HrbrainDeleteEmpInfoRequest, headers *HrbrainDeleteEmpInfoHeaders, runtime *util.RuntimeOptions) (_result *HrbrainDeleteEmpInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
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
		Action:      tea.String("HrbrainDeleteEmpInfo"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/empInfos/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainDeleteEmpInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除人员信息
//
// @param request - HrbrainDeleteEmpInfoRequest
//
// @return HrbrainDeleteEmpInfoResponse
func (client *Client) HrbrainDeleteEmpInfo(request *HrbrainDeleteEmpInfoRequest) (_result *HrbrainDeleteEmpInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainDeleteEmpInfoHeaders{}
	_result = &HrbrainDeleteEmpInfoResponse{}
	_body, _err := client.HrbrainDeleteEmpInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除领域经验
//
// @param request - HrbrainDeleteLabelIndustryRequest
//
// @param headers - HrbrainDeleteLabelIndustryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainDeleteLabelIndustryResponse
func (client *Client) HrbrainDeleteLabelIndustryWithOptions(request *HrbrainDeleteLabelIndustryRequest, headers *HrbrainDeleteLabelIndustryHeaders, runtime *util.RuntimeOptions) (_result *HrbrainDeleteLabelIndustryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
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
		Action:      tea.String("HrbrainDeleteLabelIndustry"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/industries/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainDeleteLabelIndustryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除领域经验
//
// @param request - HrbrainDeleteLabelIndustryRequest
//
// @return HrbrainDeleteLabelIndustryResponse
func (client *Client) HrbrainDeleteLabelIndustry(request *HrbrainDeleteLabelIndustryRequest) (_result *HrbrainDeleteLabelIndustryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainDeleteLabelIndustryHeaders{}
	_result = &HrbrainDeleteLabelIndustryResponse{}
	_body, _err := client.HrbrainDeleteLabelIndustryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除盘点数据
//
// @param request - HrbrainDeleteLabelInventoryRequest
//
// @param headers - HrbrainDeleteLabelInventoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainDeleteLabelInventoryResponse
func (client *Client) HrbrainDeleteLabelInventoryWithOptions(request *HrbrainDeleteLabelInventoryRequest, headers *HrbrainDeleteLabelInventoryHeaders, runtime *util.RuntimeOptions) (_result *HrbrainDeleteLabelInventoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
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
		Action:      tea.String("HrbrainDeleteLabelInventory"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/inventories/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainDeleteLabelInventoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除盘点数据
//
// @param request - HrbrainDeleteLabelInventoryRequest
//
// @return HrbrainDeleteLabelInventoryResponse
func (client *Client) HrbrainDeleteLabelInventory(request *HrbrainDeleteLabelInventoryRequest) (_result *HrbrainDeleteLabelInventoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainDeleteLabelInventoryHeaders{}
	_result = &HrbrainDeleteLabelInventoryResponse{}
	_body, _err := client.HrbrainDeleteLabelInventoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除专业技能
//
// @param request - HrbrainDeleteLabelProfSkillRequest
//
// @param headers - HrbrainDeleteLabelProfSkillHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainDeleteLabelProfSkillResponse
func (client *Client) HrbrainDeleteLabelProfSkillWithOptions(request *HrbrainDeleteLabelProfSkillRequest, headers *HrbrainDeleteLabelProfSkillHeaders, runtime *util.RuntimeOptions) (_result *HrbrainDeleteLabelProfSkillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
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
		Action:      tea.String("HrbrainDeleteLabelProfSkill"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/profSkills/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainDeleteLabelProfSkillResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除专业技能
//
// @param request - HrbrainDeleteLabelProfSkillRequest
//
// @return HrbrainDeleteLabelProfSkillResponse
func (client *Client) HrbrainDeleteLabelProfSkill(request *HrbrainDeleteLabelProfSkillRequest) (_result *HrbrainDeleteLabelProfSkillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainDeleteLabelProfSkillHeaders{}
	_result = &HrbrainDeleteLabelProfSkillResponse{}
	_body, _err := client.HrbrainDeleteLabelProfSkillWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除绩效记录
//
// @param request - HrbrainDeletePerfEvalRequest
//
// @param headers - HrbrainDeletePerfEvalHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainDeletePerfEvalResponse
func (client *Client) HrbrainDeletePerfEvalWithOptions(request *HrbrainDeletePerfEvalRequest, headers *HrbrainDeletePerfEvalHeaders, runtime *util.RuntimeOptions) (_result *HrbrainDeletePerfEvalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
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
		Action:      tea.String("HrbrainDeletePerfEval"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/perfRecords/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainDeletePerfEvalResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除绩效记录
//
// @param request - HrbrainDeletePerfEvalRequest
//
// @return HrbrainDeletePerfEvalResponse
func (client *Client) HrbrainDeletePerfEval(request *HrbrainDeletePerfEvalRequest) (_result *HrbrainDeletePerfEvalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainDeletePerfEvalHeaders{}
	_result = &HrbrainDeletePerfEvalResponse{}
	_body, _err := client.HrbrainDeletePerfEvalWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 数据集成晋升记录删除
//
// @param request - HrbrainDeletePromRecordsRequest
//
// @param headers - HrbrainDeletePromRecordsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainDeletePromRecordsResponse
func (client *Client) HrbrainDeletePromRecordsWithOptions(request *HrbrainDeletePromRecordsRequest, headers *HrbrainDeletePromRecordsHeaders, runtime *util.RuntimeOptions) (_result *HrbrainDeletePromRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
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
		Action:      tea.String("HrbrainDeletePromRecords"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/promEvals/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainDeletePromRecordsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 数据集成晋升记录删除
//
// @param request - HrbrainDeletePromRecordsRequest
//
// @return HrbrainDeletePromRecordsResponse
func (client *Client) HrbrainDeletePromRecords(request *HrbrainDeletePromRecordsRequest) (_result *HrbrainDeletePromRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainDeletePromRecordsHeaders{}
	_result = &HrbrainDeletePromRecordsResponse{}
	_body, _err := client.HrbrainDeletePromRecordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除处分记录
//
// @param request - HrbrainDeletePunDetailRequest
//
// @param headers - HrbrainDeletePunDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainDeletePunDetailResponse
func (client *Client) HrbrainDeletePunDetailWithOptions(request *HrbrainDeletePunDetailRequest, headers *HrbrainDeletePunDetailHeaders, runtime *util.RuntimeOptions) (_result *HrbrainDeletePunDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
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
		Action:      tea.String("HrbrainDeletePunDetail"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/punDetails/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainDeletePunDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除处分记录
//
// @param request - HrbrainDeletePunDetailRequest
//
// @return HrbrainDeletePunDetailResponse
func (client *Client) HrbrainDeletePunDetail(request *HrbrainDeletePunDetailRequest) (_result *HrbrainDeletePunDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainDeletePunDetailHeaders{}
	_result = &HrbrainDeletePunDetailResponse{}
	_body, _err := client.HrbrainDeletePunDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除入职记录
//
// @param request - HrbrainDeleteRegistRequest
//
// @param headers - HrbrainDeleteRegistHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainDeleteRegistResponse
func (client *Client) HrbrainDeleteRegistWithOptions(request *HrbrainDeleteRegistRequest, headers *HrbrainDeleteRegistHeaders, runtime *util.RuntimeOptions) (_result *HrbrainDeleteRegistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
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
		Action:      tea.String("HrbrainDeleteRegist"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/registerInfos/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainDeleteRegistResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除入职记录
//
// @param request - HrbrainDeleteRegistRequest
//
// @return HrbrainDeleteRegistResponse
func (client *Client) HrbrainDeleteRegist(request *HrbrainDeleteRegistRequest) (_result *HrbrainDeleteRegistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainDeleteRegistHeaders{}
	_result = &HrbrainDeleteRegistResponse{}
	_body, _err := client.HrbrainDeleteRegistWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除转正记录
//
// @param request - HrbrainDeleteRegularRequest
//
// @param headers - HrbrainDeleteRegularHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainDeleteRegularResponse
func (client *Client) HrbrainDeleteRegularWithOptions(request *HrbrainDeleteRegularRequest, headers *HrbrainDeleteRegularHeaders, runtime *util.RuntimeOptions) (_result *HrbrainDeleteRegularResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
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
		Action:      tea.String("HrbrainDeleteRegular"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/regulars/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainDeleteRegularResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除转正记录
//
// @param request - HrbrainDeleteRegularRequest
//
// @return HrbrainDeleteRegularResponse
func (client *Client) HrbrainDeleteRegular(request *HrbrainDeleteRegularRequest) (_result *HrbrainDeleteRegularResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainDeleteRegularHeaders{}
	_result = &HrbrainDeleteRegularResponse{}
	_body, _err := client.HrbrainDeleteRegularWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除培训学习记录
//
// @param request - HrbrainDeleteTrainingRequest
//
// @param headers - HrbrainDeleteTrainingHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainDeleteTrainingResponse
func (client *Client) HrbrainDeleteTrainingWithOptions(request *HrbrainDeleteTrainingRequest, headers *HrbrainDeleteTrainingHeaders, runtime *util.RuntimeOptions) (_result *HrbrainDeleteTrainingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
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
		Action:      tea.String("HrbrainDeleteTraining"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/trainings/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainDeleteTrainingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除培训学习记录
//
// @param request - HrbrainDeleteTrainingRequest
//
// @return HrbrainDeleteTrainingResponse
func (client *Client) HrbrainDeleteTraining(request *HrbrainDeleteTrainingRequest) (_result *HrbrainDeleteTrainingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainDeleteTrainingHeaders{}
	_result = &HrbrainDeleteTrainingResponse{}
	_body, _err := client.HrbrainDeleteTrainingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除调岗记录
//
// @param request - HrbrainDeleteTransferEvalRequest
//
// @param headers - HrbrainDeleteTransferEvalHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainDeleteTransferEvalResponse
func (client *Client) HrbrainDeleteTransferEvalWithOptions(request *HrbrainDeleteTransferEvalRequest, headers *HrbrainDeleteTransferEvalHeaders, runtime *util.RuntimeOptions) (_result *HrbrainDeleteTransferEvalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
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
		Action:      tea.String("HrbrainDeleteTransferEval"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/changeRecords/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainDeleteTransferEvalResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除调岗记录
//
// @param request - HrbrainDeleteTransferEvalRequest
//
// @return HrbrainDeleteTransferEvalResponse
func (client *Client) HrbrainDeleteTransferEval(request *HrbrainDeleteTransferEvalRequest) (_result *HrbrainDeleteTransferEvalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainDeleteTransferEvalHeaders{}
	_result = &HrbrainDeleteTransferEvalResponse{}
	_body, _err := client.HrbrainDeleteTransferEvalWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除工作经历
//
// @param request - HrbrainDeleteWorkExpRequest
//
// @param headers - HrbrainDeleteWorkExpHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainDeleteWorkExpResponse
func (client *Client) HrbrainDeleteWorkExpWithOptions(request *HrbrainDeleteWorkExpRequest, headers *HrbrainDeleteWorkExpHeaders, runtime *util.RuntimeOptions) (_result *HrbrainDeleteWorkExpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
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
		Action:      tea.String("HrbrainDeleteWorkExp"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/workExperiences/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainDeleteWorkExpResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除工作经历
//
// @param request - HrbrainDeleteWorkExpRequest
//
// @return HrbrainDeleteWorkExpResponse
func (client *Client) HrbrainDeleteWorkExp(request *HrbrainDeleteWorkExpRequest) (_result *HrbrainDeleteWorkExpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainDeleteWorkExpHeaders{}
	_result = &HrbrainDeleteWorkExpResponse{}
	_body, _err := client.HrbrainDeleteWorkExpWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除标签数据
//
// @param request - HrbrainDeletetLabelBaseRequest
//
// @param headers - HrbrainDeletetLabelBaseHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainDeletetLabelBaseResponse
func (client *Client) HrbrainDeletetLabelBaseWithOptions(request *HrbrainDeletetLabelBaseRequest, headers *HrbrainDeletetLabelBaseHeaders, runtime *util.RuntimeOptions) (_result *HrbrainDeletetLabelBaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
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
		Action:      tea.String("HrbrainDeletetLabelBase"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/baseLabels/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainDeletetLabelBaseResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除标签数据
//
// @param request - HrbrainDeletetLabelBaseRequest
//
// @return HrbrainDeletetLabelBaseResponse
func (client *Client) HrbrainDeletetLabelBase(request *HrbrainDeletetLabelBaseRequest) (_result *HrbrainDeletetLabelBaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainDeletetLabelBaseHeaders{}
	_result = &HrbrainDeletetLabelBaseResponse{}
	_body, _err := client.HrbrainDeletetLabelBaseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 人才池信息查询
//
// @param request - HrbrainEmpPoolQueryRequest
//
// @param headers - HrbrainEmpPoolQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainEmpPoolQueryResponse
func (client *Client) HrbrainEmpPoolQueryWithOptions(request *HrbrainEmpPoolQueryRequest, headers *HrbrainEmpPoolQueryHeaders, runtime *util.RuntimeOptions) (_result *HrbrainEmpPoolQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		body["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
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
		Action:      tea.String("HrbrainEmpPoolQuery"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/empPools/infos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainEmpPoolQueryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人才池信息查询
//
// @param request - HrbrainEmpPoolQueryRequest
//
// @return HrbrainEmpPoolQueryResponse
func (client *Client) HrbrainEmpPoolQuery(request *HrbrainEmpPoolQueryRequest) (_result *HrbrainEmpPoolQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainEmpPoolQueryHeaders{}
	_result = &HrbrainEmpPoolQueryResponse{}
	_body, _err := client.HrbrainEmpPoolQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 人才池人员查询
//
// @param request - HrbrainEmpPoolUserRequest
//
// @param headers - HrbrainEmpPoolUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainEmpPoolUserResponse
func (client *Client) HrbrainEmpPoolUserWithOptions(request *HrbrainEmpPoolUserRequest, headers *HrbrainEmpPoolUserHeaders, runtime *util.RuntimeOptions) (_result *HrbrainEmpPoolUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PoolCode)) {
		body["poolCode"] = request.PoolCode
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
		Action:      tea.String("HrbrainEmpPoolUser"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/empPools/users/lists/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainEmpPoolUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人才池人员查询
//
// @param request - HrbrainEmpPoolUserRequest
//
// @return HrbrainEmpPoolUserResponse
func (client *Client) HrbrainEmpPoolUser(request *HrbrainEmpPoolUserRequest) (_result *HrbrainEmpPoolUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainEmpPoolUserHeaders{}
	_result = &HrbrainEmpPoolUserResponse{}
	_body, _err := client.HrbrainEmpPoolUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 集成奖励记录
//
// @param request - HrbrainImportAwardDetailRequest
//
// @param headers - HrbrainImportAwardDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainImportAwardDetailResponse
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

// Summary:
//
// 集成奖励记录
//
// @param request - HrbrainImportAwardDetailRequest
//
// @return HrbrainImportAwardDetailResponse
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

// Summary:
//
// 集成自定义模型记录
//
// @param request - HrbrainImportCustomRequest
//
// @param headers - HrbrainImportCustomHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainImportCustomResponse
func (client *Client) HrbrainImportCustomWithOptions(request *HrbrainImportCustomRequest, headers *HrbrainImportCustomHeaders, runtime *util.RuntimeOptions) (_result *HrbrainImportCustomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelCode)) {
		query["modelCode"] = request.ModelCode
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
		Action:      tea.String("HrbrainImportCustom"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/customModels/import"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainImportCustomResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 集成自定义模型记录
//
// @param request - HrbrainImportCustomRequest
//
// @return HrbrainImportCustomResponse
func (client *Client) HrbrainImportCustom(request *HrbrainImportCustomRequest) (_result *HrbrainImportCustomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainImportCustomHeaders{}
	_result = &HrbrainImportCustomResponse{}
	_body, _err := client.HrbrainImportCustomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 集成组织架构
//
// @param request - HrbrainImportDeptInfoRequest
//
// @param headers - HrbrainImportDeptInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainImportDeptInfoResponse
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

// Summary:
//
// 集成组织架构
//
// @param request - HrbrainImportDeptInfoRequest
//
// @return HrbrainImportDeptInfoResponse
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

// Summary:
//
// 集成离职信息
//
// @param request - HrbrainImportDimissionRequest
//
// @param headers - HrbrainImportDimissionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainImportDimissionResponse
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

// Summary:
//
// 集成离职信息
//
// @param request - HrbrainImportDimissionRequest
//
// @return HrbrainImportDimissionResponse
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

// Summary:
//
// 集成教育经历
//
// @param request - HrbrainImportEduExpRequest
//
// @param headers - HrbrainImportEduExpHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainImportEduExpResponse
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

// Summary:
//
// 集成教育经历
//
// @param request - HrbrainImportEduExpRequest
//
// @return HrbrainImportEduExpResponse
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

// Summary:
//
// 集成人员信息
//
// @param request - HrbrainImportEmpInfoRequest
//
// @param headers - HrbrainImportEmpInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainImportEmpInfoResponse
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

// Summary:
//
// 集成人员信息
//
// @param request - HrbrainImportEmpInfoRequest
//
// @return HrbrainImportEmpInfoResponse
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

// Summary:
//
// 集成基础标签
//
// @param request - HrbrainImportLabelBaseRequest
//
// @param headers - HrbrainImportLabelBaseHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainImportLabelBaseResponse
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

// Summary:
//
// 集成基础标签
//
// @param request - HrbrainImportLabelBaseRequest
//
// @return HrbrainImportLabelBaseResponse
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

// Summary:
//
// 集成自定义标签
//
// @param request - HrbrainImportLabelCustomRequest
//
// @param headers - HrbrainImportLabelCustomHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainImportLabelCustomResponse
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

// Summary:
//
// 集成自定义标签
//
// @param request - HrbrainImportLabelCustomRequest
//
// @return HrbrainImportLabelCustomResponse
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

// Summary:
//
// 集成领域经验
//
// @param request - HrbrainImportLabelIndustryRequest
//
// @param headers - HrbrainImportLabelIndustryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainImportLabelIndustryResponse
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

// Summary:
//
// 集成领域经验
//
// @param request - HrbrainImportLabelIndustryRequest
//
// @return HrbrainImportLabelIndustryResponse
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

// Summary:
//
// 集成盘点数据
//
// @param request - HrbrainImportLabelInventoryRequest
//
// @param headers - HrbrainImportLabelInventoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainImportLabelInventoryResponse
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

// Summary:
//
// 集成盘点数据
//
// @param request - HrbrainImportLabelInventoryRequest
//
// @return HrbrainImportLabelInventoryResponse
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

// Summary:
//
// 集成专业技能
//
// @param request - HrbrainImportLabelProfSkillRequest
//
// @param headers - HrbrainImportLabelProfSkillHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainImportLabelProfSkillResponse
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

// Summary:
//
// 集成专业技能
//
// @param request - HrbrainImportLabelProfSkillRequest
//
// @return HrbrainImportLabelProfSkillResponse
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

// Summary:
//
// 集成绩效记录
//
// @param request - HrbrainImportPerfEvalRequest
//
// @param headers - HrbrainImportPerfEvalHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainImportPerfEvalResponse
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

// Summary:
//
// 集成绩效记录
//
// @param request - HrbrainImportPerfEvalRequest
//
// @return HrbrainImportPerfEvalResponse
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

// Summary:
//
// 集成晋升记录
//
// @param request - HrbrainImportPromEvalRequest
//
// @param headers - HrbrainImportPromEvalHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainImportPromEvalResponse
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

// Summary:
//
// 集成晋升记录
//
// @param request - HrbrainImportPromEvalRequest
//
// @return HrbrainImportPromEvalResponse
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

// Summary:
//
// 集成处分记录
//
// @param request - HrbrainImportPunDetailRequest
//
// @param headers - HrbrainImportPunDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainImportPunDetailResponse
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

// Summary:
//
// 集成处分记录
//
// @param request - HrbrainImportPunDetailRequest
//
// @return HrbrainImportPunDetailResponse
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

// Summary:
//
// 集成入职信息
//
// @param request - HrbrainImportRegistRequest
//
// @param headers - HrbrainImportRegistHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainImportRegistResponse
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

// Summary:
//
// 集成入职信息
//
// @param request - HrbrainImportRegistRequest
//
// @return HrbrainImportRegistResponse
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

// Summary:
//
// 集成转正记录
//
// @param request - HrbrainImportRegularRequest
//
// @param headers - HrbrainImportRegularHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainImportRegularResponse
func (client *Client) HrbrainImportRegularWithOptions(request *HrbrainImportRegularRequest, headers *HrbrainImportRegularHeaders, runtime *util.RuntimeOptions) (_result *HrbrainImportRegularResponse, _err error) {
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
		Action:      tea.String("HrbrainImportRegular"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/regulars/import"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainImportRegularResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 集成转正记录
//
// @param request - HrbrainImportRegularRequest
//
// @return HrbrainImportRegularResponse
func (client *Client) HrbrainImportRegular(request *HrbrainImportRegularRequest) (_result *HrbrainImportRegularResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainImportRegularHeaders{}
	_result = &HrbrainImportRegularResponse{}
	_body, _err := client.HrbrainImportRegularWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 集成培训学习记录
//
// @param request - HrbrainImportTrainingRequest
//
// @param headers - HrbrainImportTrainingHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainImportTrainingResponse
func (client *Client) HrbrainImportTrainingWithOptions(request *HrbrainImportTrainingRequest, headers *HrbrainImportTrainingHeaders, runtime *util.RuntimeOptions) (_result *HrbrainImportTrainingResponse, _err error) {
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
		Action:      tea.String("HrbrainImportTraining"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/trainings/import"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainImportTrainingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 集成培训学习记录
//
// @param request - HrbrainImportTrainingRequest
//
// @return HrbrainImportTrainingResponse
func (client *Client) HrbrainImportTraining(request *HrbrainImportTrainingRequest) (_result *HrbrainImportTrainingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainImportTrainingHeaders{}
	_result = &HrbrainImportTrainingResponse{}
	_body, _err := client.HrbrainImportTrainingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 集成异动记录
//
// @param request - HrbrainImportTransferEvalRequest
//
// @param headers - HrbrainImportTransferEvalHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainImportTransferEvalResponse
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

// Summary:
//
// 集成异动记录
//
// @param request - HrbrainImportTransferEvalRequest
//
// @return HrbrainImportTransferEvalResponse
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

// Summary:
//
// 集成工作经历
//
// @param request - HrbrainImportWorkExpRequest
//
// @param headers - HrbrainImportWorkExpHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainImportWorkExpResponse
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

// Summary:
//
// 集成工作经历
//
// @param request - HrbrainImportWorkExpRequest
//
// @return HrbrainImportWorkExpResponse
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

// Summary:
//
// 查询人才档案附件照片
//
// @param request - HrbrainTalentProfileAttachmentQueryRequest
//
// @param headers - HrbrainTalentProfileAttachmentQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainTalentProfileAttachmentQueryResponse
func (client *Client) HrbrainTalentProfileAttachmentQueryWithOptions(request *HrbrainTalentProfileAttachmentQueryRequest, headers *HrbrainTalentProfileAttachmentQueryHeaders, runtime *util.RuntimeOptions) (_result *HrbrainTalentProfileAttachmentQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingCorpId)) {
		query["dingCorpId"] = request.DingCorpId
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
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("HrbrainTalentProfileAttachmentQuery"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/profiles/attachmentPhotos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainTalentProfileAttachmentQueryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询人才档案附件照片
//
// @param request - HrbrainTalentProfileAttachmentQueryRequest
//
// @return HrbrainTalentProfileAttachmentQueryResponse
func (client *Client) HrbrainTalentProfileAttachmentQuery(request *HrbrainTalentProfileAttachmentQueryRequest) (_result *HrbrainTalentProfileAttachmentQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainTalentProfileAttachmentQueryHeaders{}
	_result = &HrbrainTalentProfileAttachmentQueryResponse{}
	_body, _err := client.HrbrainTalentProfileAttachmentQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询人才档案基础数据
//
// @param request - HrbrainTalentProfileBasicQueryRequest
//
// @param headers - HrbrainTalentProfileBasicQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrbrainTalentProfileBasicQueryResponse
func (client *Client) HrbrainTalentProfileBasicQueryWithOptions(request *HrbrainTalentProfileBasicQueryRequest, headers *HrbrainTalentProfileBasicQueryHeaders, runtime *util.RuntimeOptions) (_result *HrbrainTalentProfileBasicQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingCorpId)) {
		query["dingCorpId"] = request.DingCorpId
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
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("HrbrainTalentProfileBasicQuery"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/profiles/basicData/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrbrainTalentProfileBasicQueryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询人才档案基础数据
//
// @param request - HrbrainTalentProfileBasicQueryRequest
//
// @return HrbrainTalentProfileBasicQueryResponse
func (client *Client) HrbrainTalentProfileBasicQuery(request *HrbrainTalentProfileBasicQueryRequest) (_result *HrbrainTalentProfileBasicQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrbrainTalentProfileBasicQueryHeaders{}
	_result = &HrbrainTalentProfileBasicQueryResponse{}
	_body, _err := client.HrbrainTalentProfileBasicQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 人员标签查询
//
// @param request - StaffLabelRecordsQueryRequest
//
// @param headers - StaffLabelRecordsQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StaffLabelRecordsQueryResponse
func (client *Client) StaffLabelRecordsQueryWithOptions(request *StaffLabelRecordsQueryRequest, headers *StaffLabelRecordsQueryHeaders, runtime *util.RuntimeOptions) (_result *StaffLabelRecordsQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingCorpId)) {
		query["dingCorpId"] = request.DingCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
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
		Action:      tea.String("StaffLabelRecordsQuery"),
		Version:     tea.String("hrbrain_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrbrain/datas/labelRecords/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &StaffLabelRecordsQueryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人员标签查询
//
// @param request - StaffLabelRecordsQueryRequest
//
// @return StaffLabelRecordsQueryResponse
func (client *Client) StaffLabelRecordsQuery(request *StaffLabelRecordsQueryRequest) (_result *StaffLabelRecordsQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StaffLabelRecordsQueryHeaders{}
	_result = &StaffLabelRecordsQueryResponse{}
	_body, _err := client.StaffLabelRecordsQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 同步统计基础数据
//
// @param request - SyncDataRequest
//
// @param headers - SyncDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncDataResponse
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

// Summary:
//
// 同步统计基础数据
//
// @param request - SyncDataRequest
//
// @return SyncDataResponse
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
