// This file is auto-generated, don't edit it. Thanks.
package contract_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CancelReviewOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CancelReviewOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s CancelReviewOrderHeaders) GoString() string {
	return s.String()
}

func (s *CancelReviewOrderHeaders) SetCommonHeaders(v map[string]*string) *CancelReviewOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CancelReviewOrderHeaders) SetXAcsDingtalkAccessToken(v string) *CancelReviewOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CancelReviewOrderRequest struct {
	EndFiles  []*CancelReviewOrderRequestEndFiles `json:"endFiles,omitempty" xml:"endFiles,omitempty" type:"Repeated"`
	Extension *string                             `json:"extension,omitempty" xml:"extension,omitempty"`
	OrderId   *string                             `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s CancelReviewOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelReviewOrderRequest) GoString() string {
	return s.String()
}

func (s *CancelReviewOrderRequest) SetEndFiles(v []*CancelReviewOrderRequestEndFiles) *CancelReviewOrderRequest {
	s.EndFiles = v
	return s
}

func (s *CancelReviewOrderRequest) SetExtension(v string) *CancelReviewOrderRequest {
	s.Extension = &v
	return s
}

func (s *CancelReviewOrderRequest) SetOrderId(v string) *CancelReviewOrderRequest {
	s.OrderId = &v
	return s
}

type CancelReviewOrderRequestEndFiles struct {
	FileName    *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileSize    *string `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	FileType    *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	FileVersion *int32  `json:"fileVersion,omitempty" xml:"fileVersion,omitempty"`
	Url         *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CancelReviewOrderRequestEndFiles) String() string {
	return tea.Prettify(s)
}

func (s CancelReviewOrderRequestEndFiles) GoString() string {
	return s.String()
}

func (s *CancelReviewOrderRequestEndFiles) SetFileName(v string) *CancelReviewOrderRequestEndFiles {
	s.FileName = &v
	return s
}

func (s *CancelReviewOrderRequestEndFiles) SetFileSize(v string) *CancelReviewOrderRequestEndFiles {
	s.FileSize = &v
	return s
}

func (s *CancelReviewOrderRequestEndFiles) SetFileType(v string) *CancelReviewOrderRequestEndFiles {
	s.FileType = &v
	return s
}

func (s *CancelReviewOrderRequestEndFiles) SetFileVersion(v int32) *CancelReviewOrderRequestEndFiles {
	s.FileVersion = &v
	return s
}

func (s *CancelReviewOrderRequestEndFiles) SetUrl(v string) *CancelReviewOrderRequestEndFiles {
	s.Url = &v
	return s
}

type CancelReviewOrderResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CancelReviewOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelReviewOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CancelReviewOrderResponseBody) SetResult(v string) *CancelReviewOrderResponseBody {
	s.Result = &v
	return s
}

func (s *CancelReviewOrderResponseBody) SetSuccess(v bool) *CancelReviewOrderResponseBody {
	s.Success = &v
	return s
}

type CancelReviewOrderResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelReviewOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelReviewOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelReviewOrderResponse) GoString() string {
	return s.String()
}

func (s *CancelReviewOrderResponse) SetHeaders(v map[string]*string) *CancelReviewOrderResponse {
	s.Headers = v
	return s
}

func (s *CancelReviewOrderResponse) SetStatusCode(v int32) *CancelReviewOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelReviewOrderResponse) SetBody(v *CancelReviewOrderResponseBody) *CancelReviewOrderResponse {
	s.Body = v
	return s
}

type CheckEsignFileHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CheckEsignFileHeaders) String() string {
	return tea.Prettify(s)
}

func (s CheckEsignFileHeaders) GoString() string {
	return s.String()
}

func (s *CheckEsignFileHeaders) SetCommonHeaders(v map[string]*string) *CheckEsignFileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CheckEsignFileHeaders) SetXAcsDingtalkAccessToken(v string) *CheckEsignFileHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CheckEsignFileRequest struct {
	// This parameter is required.
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// This parameter is required.
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// This parameter is required.
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CheckEsignFileRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckEsignFileRequest) GoString() string {
	return s.String()
}

func (s *CheckEsignFileRequest) SetCorpId(v string) *CheckEsignFileRequest {
	s.CorpId = &v
	return s
}

func (s *CheckEsignFileRequest) SetFileId(v string) *CheckEsignFileRequest {
	s.FileId = &v
	return s
}

func (s *CheckEsignFileRequest) SetSpaceId(v string) *CheckEsignFileRequest {
	s.SpaceId = &v
	return s
}

func (s *CheckEsignFileRequest) SetUserId(v string) *CheckEsignFileRequest {
	s.UserId = &v
	return s
}

type CheckEsignFileResponseBody struct {
	Result  *bool `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CheckEsignFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckEsignFileResponseBody) GoString() string {
	return s.String()
}

func (s *CheckEsignFileResponseBody) SetResult(v bool) *CheckEsignFileResponseBody {
	s.Result = &v
	return s
}

func (s *CheckEsignFileResponseBody) SetSuccess(v bool) *CheckEsignFileResponseBody {
	s.Success = &v
	return s
}

type CheckEsignFileResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckEsignFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckEsignFileResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckEsignFileResponse) GoString() string {
	return s.String()
}

func (s *CheckEsignFileResponse) SetHeaders(v map[string]*string) *CheckEsignFileResponse {
	s.Headers = v
	return s
}

func (s *CheckEsignFileResponse) SetStatusCode(v int32) *CheckEsignFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckEsignFileResponse) SetBody(v *CheckEsignFileResponseBody) *CheckEsignFileResponse {
	s.Body = v
	return s
}

type ContractBenefitConsumeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ContractBenefitConsumeHeaders) String() string {
	return tea.Prettify(s)
}

func (s ContractBenefitConsumeHeaders) GoString() string {
	return s.String()
}

func (s *ContractBenefitConsumeHeaders) SetCommonHeaders(v map[string]*string) *ContractBenefitConsumeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ContractBenefitConsumeHeaders) SetXAcsDingtalkAccessToken(v string) *ContractBenefitConsumeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ContractBenefitConsumeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// esign
	BenefitPoint *string `json:"benefitPoint,omitempty" xml:"benefitPoint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sjdaujii129w9qej2nqas
	BizRequestId *string `json:"bizRequestId,omitempty" xml:"bizRequestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	ConsumeQuota *int64 `json:"consumeQuota,omitempty" xml:"consumeQuota,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ding1231ndu29923312
	CorpId    *string            `json:"corpId,omitempty" xml:"corpId,omitempty"`
	ExtParams map[string]*string `json:"extParams,omitempty" xml:"extParams,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ding12939912nduaejjwe
	IsvCorpId *string `json:"isvCorpId,omitempty" xml:"isvCorpId,omitempty"`
	// example:
	//
	// djauihjauiwnkndjknkjanaae
	OptUnionId *string `json:"optUnionId,omitempty" xml:"optUnionId,omitempty"`
}

func (s ContractBenefitConsumeRequest) String() string {
	return tea.Prettify(s)
}

func (s ContractBenefitConsumeRequest) GoString() string {
	return s.String()
}

func (s *ContractBenefitConsumeRequest) SetBenefitPoint(v string) *ContractBenefitConsumeRequest {
	s.BenefitPoint = &v
	return s
}

func (s *ContractBenefitConsumeRequest) SetBizRequestId(v string) *ContractBenefitConsumeRequest {
	s.BizRequestId = &v
	return s
}

func (s *ContractBenefitConsumeRequest) SetConsumeQuota(v int64) *ContractBenefitConsumeRequest {
	s.ConsumeQuota = &v
	return s
}

func (s *ContractBenefitConsumeRequest) SetCorpId(v string) *ContractBenefitConsumeRequest {
	s.CorpId = &v
	return s
}

func (s *ContractBenefitConsumeRequest) SetExtParams(v map[string]*string) *ContractBenefitConsumeRequest {
	s.ExtParams = v
	return s
}

func (s *ContractBenefitConsumeRequest) SetIsvCorpId(v string) *ContractBenefitConsumeRequest {
	s.IsvCorpId = &v
	return s
}

func (s *ContractBenefitConsumeRequest) SetOptUnionId(v string) *ContractBenefitConsumeRequest {
	s.OptUnionId = &v
	return s
}

type ContractBenefitConsumeResponseBody struct {
	Result  *ContractBenefitConsumeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ContractBenefitConsumeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ContractBenefitConsumeResponseBody) GoString() string {
	return s.String()
}

func (s *ContractBenefitConsumeResponseBody) SetResult(v *ContractBenefitConsumeResponseBodyResult) *ContractBenefitConsumeResponseBody {
	s.Result = v
	return s
}

func (s *ContractBenefitConsumeResponseBody) SetSuccess(v bool) *ContractBenefitConsumeResponseBody {
	s.Success = &v
	return s
}

type ContractBenefitConsumeResponseBodyResult struct {
	ConsumeResult *bool `json:"consumeResult,omitempty" xml:"consumeResult,omitempty"`
}

func (s ContractBenefitConsumeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ContractBenefitConsumeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ContractBenefitConsumeResponseBodyResult) SetConsumeResult(v bool) *ContractBenefitConsumeResponseBodyResult {
	s.ConsumeResult = &v
	return s
}

type ContractBenefitConsumeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ContractBenefitConsumeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContractBenefitConsumeResponse) String() string {
	return tea.Prettify(s)
}

func (s ContractBenefitConsumeResponse) GoString() string {
	return s.String()
}

func (s *ContractBenefitConsumeResponse) SetHeaders(v map[string]*string) *ContractBenefitConsumeResponse {
	s.Headers = v
	return s
}

func (s *ContractBenefitConsumeResponse) SetStatusCode(v int32) *ContractBenefitConsumeResponse {
	s.StatusCode = &v
	return s
}

func (s *ContractBenefitConsumeResponse) SetBody(v *ContractBenefitConsumeResponseBody) *ContractBenefitConsumeResponse {
	s.Body = v
	return s
}

type CreateContractAppsCompareTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateContractAppsCompareTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateContractAppsCompareTaskHeaders) GoString() string {
	return s.String()
}

func (s *CreateContractAppsCompareTaskHeaders) SetCommonHeaders(v map[string]*string) *CreateContractAppsCompareTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateContractAppsCompareTaskHeaders) SetXAcsDingtalkAccessToken(v string) *CreateContractAppsCompareTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateContractAppsCompareTaskRequest struct {
	ComparativeFile            *CreateContractAppsCompareTaskRequestComparativeFile `json:"comparativeFile,omitempty" xml:"comparativeFile,omitempty" type:"Struct"`
	ComparativeFileDownloadUrl *string                                              `json:"comparativeFileDownloadUrl,omitempty" xml:"comparativeFileDownloadUrl,omitempty"`
	// This parameter is required.
	ComparativeFileName *string `json:"comparativeFileName,omitempty" xml:"comparativeFileName,omitempty"`
	// This parameter is required.
	FileSource *string `json:"fileSource,omitempty" xml:"fileSource,omitempty"`
	// This parameter is required.
	RequestId               *string                                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	StandardFile            *CreateContractAppsCompareTaskRequestStandardFile `json:"standardFile,omitempty" xml:"standardFile,omitempty" type:"Struct"`
	StandardFileDownloadUrl *string                                           `json:"standardFileDownloadUrl,omitempty" xml:"standardFileDownloadUrl,omitempty"`
	// This parameter is required.
	StandardFileName *string `json:"standardFileName,omitempty" xml:"standardFileName,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s CreateContractAppsCompareTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateContractAppsCompareTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateContractAppsCompareTaskRequest) SetComparativeFile(v *CreateContractAppsCompareTaskRequestComparativeFile) *CreateContractAppsCompareTaskRequest {
	s.ComparativeFile = v
	return s
}

func (s *CreateContractAppsCompareTaskRequest) SetComparativeFileDownloadUrl(v string) *CreateContractAppsCompareTaskRequest {
	s.ComparativeFileDownloadUrl = &v
	return s
}

func (s *CreateContractAppsCompareTaskRequest) SetComparativeFileName(v string) *CreateContractAppsCompareTaskRequest {
	s.ComparativeFileName = &v
	return s
}

func (s *CreateContractAppsCompareTaskRequest) SetFileSource(v string) *CreateContractAppsCompareTaskRequest {
	s.FileSource = &v
	return s
}

func (s *CreateContractAppsCompareTaskRequest) SetRequestId(v string) *CreateContractAppsCompareTaskRequest {
	s.RequestId = &v
	return s
}

func (s *CreateContractAppsCompareTaskRequest) SetStandardFile(v *CreateContractAppsCompareTaskRequestStandardFile) *CreateContractAppsCompareTaskRequest {
	s.StandardFile = v
	return s
}

func (s *CreateContractAppsCompareTaskRequest) SetStandardFileDownloadUrl(v string) *CreateContractAppsCompareTaskRequest {
	s.StandardFileDownloadUrl = &v
	return s
}

func (s *CreateContractAppsCompareTaskRequest) SetStandardFileName(v string) *CreateContractAppsCompareTaskRequest {
	s.StandardFileName = &v
	return s
}

func (s *CreateContractAppsCompareTaskRequest) SetUnionId(v string) *CreateContractAppsCompareTaskRequest {
	s.UnionId = &v
	return s
}

type CreateContractAppsCompareTaskRequestComparativeFile struct {
	FileId   *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileSize *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	SpaceId  *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s CreateContractAppsCompareTaskRequestComparativeFile) String() string {
	return tea.Prettify(s)
}

func (s CreateContractAppsCompareTaskRequestComparativeFile) GoString() string {
	return s.String()
}

func (s *CreateContractAppsCompareTaskRequestComparativeFile) SetFileId(v string) *CreateContractAppsCompareTaskRequestComparativeFile {
	s.FileId = &v
	return s
}

func (s *CreateContractAppsCompareTaskRequestComparativeFile) SetFileName(v string) *CreateContractAppsCompareTaskRequestComparativeFile {
	s.FileName = &v
	return s
}

func (s *CreateContractAppsCompareTaskRequestComparativeFile) SetFileSize(v int64) *CreateContractAppsCompareTaskRequestComparativeFile {
	s.FileSize = &v
	return s
}

func (s *CreateContractAppsCompareTaskRequestComparativeFile) SetFileType(v string) *CreateContractAppsCompareTaskRequestComparativeFile {
	s.FileType = &v
	return s
}

func (s *CreateContractAppsCompareTaskRequestComparativeFile) SetSpaceId(v string) *CreateContractAppsCompareTaskRequestComparativeFile {
	s.SpaceId = &v
	return s
}

type CreateContractAppsCompareTaskRequestStandardFile struct {
	FileId   *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileSize *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	SpaceId  *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s CreateContractAppsCompareTaskRequestStandardFile) String() string {
	return tea.Prettify(s)
}

func (s CreateContractAppsCompareTaskRequestStandardFile) GoString() string {
	return s.String()
}

func (s *CreateContractAppsCompareTaskRequestStandardFile) SetFileId(v string) *CreateContractAppsCompareTaskRequestStandardFile {
	s.FileId = &v
	return s
}

func (s *CreateContractAppsCompareTaskRequestStandardFile) SetFileName(v string) *CreateContractAppsCompareTaskRequestStandardFile {
	s.FileName = &v
	return s
}

func (s *CreateContractAppsCompareTaskRequestStandardFile) SetFileSize(v int64) *CreateContractAppsCompareTaskRequestStandardFile {
	s.FileSize = &v
	return s
}

func (s *CreateContractAppsCompareTaskRequestStandardFile) SetFileType(v string) *CreateContractAppsCompareTaskRequestStandardFile {
	s.FileType = &v
	return s
}

func (s *CreateContractAppsCompareTaskRequestStandardFile) SetSpaceId(v string) *CreateContractAppsCompareTaskRequestStandardFile {
	s.SpaceId = &v
	return s
}

type CreateContractAppsCompareTaskResponseBody struct {
	Result  *CreateContractAppsCompareTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                            `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateContractAppsCompareTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateContractAppsCompareTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContractAppsCompareTaskResponseBody) SetResult(v *CreateContractAppsCompareTaskResponseBodyResult) *CreateContractAppsCompareTaskResponseBody {
	s.Result = v
	return s
}

func (s *CreateContractAppsCompareTaskResponseBody) SetSuccess(v bool) *CreateContractAppsCompareTaskResponseBody {
	s.Success = &v
	return s
}

type CreateContractAppsCompareTaskResponseBodyResult struct {
	Data      *CreateContractAppsCompareTaskResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateContractAppsCompareTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateContractAppsCompareTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateContractAppsCompareTaskResponseBodyResult) SetData(v *CreateContractAppsCompareTaskResponseBodyResultData) *CreateContractAppsCompareTaskResponseBodyResult {
	s.Data = v
	return s
}

func (s *CreateContractAppsCompareTaskResponseBodyResult) SetRequestId(v string) *CreateContractAppsCompareTaskResponseBodyResult {
	s.RequestId = &v
	return s
}

type CreateContractAppsCompareTaskResponseBodyResultData struct {
	CompareTaskId *string `json:"compareTaskId,omitempty" xml:"compareTaskId,omitempty"`
}

func (s CreateContractAppsCompareTaskResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s CreateContractAppsCompareTaskResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *CreateContractAppsCompareTaskResponseBodyResultData) SetCompareTaskId(v string) *CreateContractAppsCompareTaskResponseBodyResultData {
	s.CompareTaskId = &v
	return s
}

type CreateContractAppsCompareTaskResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateContractAppsCompareTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateContractAppsCompareTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateContractAppsCompareTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateContractAppsCompareTaskResponse) SetHeaders(v map[string]*string) *CreateContractAppsCompareTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateContractAppsCompareTaskResponse) SetStatusCode(v int32) *CreateContractAppsCompareTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateContractAppsCompareTaskResponse) SetBody(v *CreateContractAppsCompareTaskResponseBody) *CreateContractAppsCompareTaskResponse {
	s.Body = v
	return s
}

type CreateContractAppsExtractTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateContractAppsExtractTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateContractAppsExtractTaskHeaders) GoString() string {
	return s.String()
}

func (s *CreateContractAppsExtractTaskHeaders) SetCommonHeaders(v map[string]*string) *CreateContractAppsExtractTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateContractAppsExtractTaskHeaders) SetXAcsDingtalkAccessToken(v string) *CreateContractAppsExtractTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateContractAppsExtractTaskRequest struct {
	ContractFile            *CreateContractAppsExtractTaskRequestContractFile `json:"contractFile,omitempty" xml:"contractFile,omitempty" type:"Struct"`
	ContractFileDownloadUrl *string                                           `json:"contractFileDownloadUrl,omitempty" xml:"contractFileDownloadUrl,omitempty"`
	// This parameter is required.
	ContractFileName *string `json:"contractFileName,omitempty" xml:"contractFileName,omitempty"`
	// This parameter is required.
	ExtractKeys []*string `json:"extractKeys,omitempty" xml:"extractKeys,omitempty" type:"Repeated"`
	// This parameter is required.
	FileSource *string `json:"fileSource,omitempty" xml:"fileSource,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s CreateContractAppsExtractTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateContractAppsExtractTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateContractAppsExtractTaskRequest) SetContractFile(v *CreateContractAppsExtractTaskRequestContractFile) *CreateContractAppsExtractTaskRequest {
	s.ContractFile = v
	return s
}

func (s *CreateContractAppsExtractTaskRequest) SetContractFileDownloadUrl(v string) *CreateContractAppsExtractTaskRequest {
	s.ContractFileDownloadUrl = &v
	return s
}

func (s *CreateContractAppsExtractTaskRequest) SetContractFileName(v string) *CreateContractAppsExtractTaskRequest {
	s.ContractFileName = &v
	return s
}

func (s *CreateContractAppsExtractTaskRequest) SetExtractKeys(v []*string) *CreateContractAppsExtractTaskRequest {
	s.ExtractKeys = v
	return s
}

func (s *CreateContractAppsExtractTaskRequest) SetFileSource(v string) *CreateContractAppsExtractTaskRequest {
	s.FileSource = &v
	return s
}

func (s *CreateContractAppsExtractTaskRequest) SetRequestId(v string) *CreateContractAppsExtractTaskRequest {
	s.RequestId = &v
	return s
}

func (s *CreateContractAppsExtractTaskRequest) SetUnionId(v string) *CreateContractAppsExtractTaskRequest {
	s.UnionId = &v
	return s
}

type CreateContractAppsExtractTaskRequestContractFile struct {
	FileId   *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileSize *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	SpaceId  *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s CreateContractAppsExtractTaskRequestContractFile) String() string {
	return tea.Prettify(s)
}

func (s CreateContractAppsExtractTaskRequestContractFile) GoString() string {
	return s.String()
}

func (s *CreateContractAppsExtractTaskRequestContractFile) SetFileId(v string) *CreateContractAppsExtractTaskRequestContractFile {
	s.FileId = &v
	return s
}

func (s *CreateContractAppsExtractTaskRequestContractFile) SetFileName(v string) *CreateContractAppsExtractTaskRequestContractFile {
	s.FileName = &v
	return s
}

func (s *CreateContractAppsExtractTaskRequestContractFile) SetFileSize(v int64) *CreateContractAppsExtractTaskRequestContractFile {
	s.FileSize = &v
	return s
}

func (s *CreateContractAppsExtractTaskRequestContractFile) SetFileType(v string) *CreateContractAppsExtractTaskRequestContractFile {
	s.FileType = &v
	return s
}

func (s *CreateContractAppsExtractTaskRequestContractFile) SetSpaceId(v string) *CreateContractAppsExtractTaskRequestContractFile {
	s.SpaceId = &v
	return s
}

type CreateContractAppsExtractTaskResponseBody struct {
	Result  *CreateContractAppsExtractTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                            `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateContractAppsExtractTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateContractAppsExtractTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContractAppsExtractTaskResponseBody) SetResult(v *CreateContractAppsExtractTaskResponseBodyResult) *CreateContractAppsExtractTaskResponseBody {
	s.Result = v
	return s
}

func (s *CreateContractAppsExtractTaskResponseBody) SetSuccess(v bool) *CreateContractAppsExtractTaskResponseBody {
	s.Success = &v
	return s
}

type CreateContractAppsExtractTaskResponseBodyResult struct {
	Data      *CreateContractAppsExtractTaskResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateContractAppsExtractTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateContractAppsExtractTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateContractAppsExtractTaskResponseBodyResult) SetData(v *CreateContractAppsExtractTaskResponseBodyResultData) *CreateContractAppsExtractTaskResponseBodyResult {
	s.Data = v
	return s
}

func (s *CreateContractAppsExtractTaskResponseBodyResult) SetRequestId(v string) *CreateContractAppsExtractTaskResponseBodyResult {
	s.RequestId = &v
	return s
}

type CreateContractAppsExtractTaskResponseBodyResultData struct {
	ExtractTaskId *string `json:"extractTaskId,omitempty" xml:"extractTaskId,omitempty"`
}

func (s CreateContractAppsExtractTaskResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s CreateContractAppsExtractTaskResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *CreateContractAppsExtractTaskResponseBodyResultData) SetExtractTaskId(v string) *CreateContractAppsExtractTaskResponseBodyResultData {
	s.ExtractTaskId = &v
	return s
}

type CreateContractAppsExtractTaskResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateContractAppsExtractTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateContractAppsExtractTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateContractAppsExtractTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateContractAppsExtractTaskResponse) SetHeaders(v map[string]*string) *CreateContractAppsExtractTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateContractAppsExtractTaskResponse) SetStatusCode(v int32) *CreateContractAppsExtractTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateContractAppsExtractTaskResponse) SetBody(v *CreateContractAppsExtractTaskResponseBody) *CreateContractAppsExtractTaskResponse {
	s.Body = v
	return s
}

type CreateContractAppsReviewTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateContractAppsReviewTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateContractAppsReviewTaskHeaders) GoString() string {
	return s.String()
}

func (s *CreateContractAppsReviewTaskHeaders) SetCommonHeaders(v map[string]*string) *CreateContractAppsReviewTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateContractAppsReviewTaskHeaders) SetXAcsDingtalkAccessToken(v string) *CreateContractAppsReviewTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateContractAppsReviewTaskRequest struct {
	ContractFile            *CreateContractAppsReviewTaskRequestContractFile `json:"contractFile,omitempty" xml:"contractFile,omitempty" type:"Struct"`
	ContractFileDownloadUrl *string                                          `json:"contractFileDownloadUrl,omitempty" xml:"contractFileDownloadUrl,omitempty"`
	// This parameter is required.
	ContractFileName *string `json:"contractFileName,omitempty" xml:"contractFileName,omitempty"`
	// This parameter is required.
	FileSource *string `json:"fileSource,omitempty" xml:"fileSource,omitempty"`
	// This parameter is required.
	RequestId         *string                                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ReviewCustomRules []*CreateContractAppsReviewTaskRequestReviewCustomRules `json:"reviewCustomRules,omitempty" xml:"reviewCustomRules,omitempty" type:"Repeated"`
	RuleType          *string                                                 `json:"ruleType,omitempty" xml:"ruleType,omitempty"`
	Standpoint        *string                                                 `json:"standpoint,omitempty" xml:"standpoint,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s CreateContractAppsReviewTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateContractAppsReviewTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateContractAppsReviewTaskRequest) SetContractFile(v *CreateContractAppsReviewTaskRequestContractFile) *CreateContractAppsReviewTaskRequest {
	s.ContractFile = v
	return s
}

func (s *CreateContractAppsReviewTaskRequest) SetContractFileDownloadUrl(v string) *CreateContractAppsReviewTaskRequest {
	s.ContractFileDownloadUrl = &v
	return s
}

func (s *CreateContractAppsReviewTaskRequest) SetContractFileName(v string) *CreateContractAppsReviewTaskRequest {
	s.ContractFileName = &v
	return s
}

func (s *CreateContractAppsReviewTaskRequest) SetFileSource(v string) *CreateContractAppsReviewTaskRequest {
	s.FileSource = &v
	return s
}

func (s *CreateContractAppsReviewTaskRequest) SetRequestId(v string) *CreateContractAppsReviewTaskRequest {
	s.RequestId = &v
	return s
}

func (s *CreateContractAppsReviewTaskRequest) SetReviewCustomRules(v []*CreateContractAppsReviewTaskRequestReviewCustomRules) *CreateContractAppsReviewTaskRequest {
	s.ReviewCustomRules = v
	return s
}

func (s *CreateContractAppsReviewTaskRequest) SetRuleType(v string) *CreateContractAppsReviewTaskRequest {
	s.RuleType = &v
	return s
}

func (s *CreateContractAppsReviewTaskRequest) SetStandpoint(v string) *CreateContractAppsReviewTaskRequest {
	s.Standpoint = &v
	return s
}

func (s *CreateContractAppsReviewTaskRequest) SetUnionId(v string) *CreateContractAppsReviewTaskRequest {
	s.UnionId = &v
	return s
}

type CreateContractAppsReviewTaskRequestContractFile struct {
	FileId   *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileSize *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	SpaceId  *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s CreateContractAppsReviewTaskRequestContractFile) String() string {
	return tea.Prettify(s)
}

func (s CreateContractAppsReviewTaskRequestContractFile) GoString() string {
	return s.String()
}

func (s *CreateContractAppsReviewTaskRequestContractFile) SetFileId(v string) *CreateContractAppsReviewTaskRequestContractFile {
	s.FileId = &v
	return s
}

func (s *CreateContractAppsReviewTaskRequestContractFile) SetFileName(v string) *CreateContractAppsReviewTaskRequestContractFile {
	s.FileName = &v
	return s
}

func (s *CreateContractAppsReviewTaskRequestContractFile) SetFileSize(v int64) *CreateContractAppsReviewTaskRequestContractFile {
	s.FileSize = &v
	return s
}

func (s *CreateContractAppsReviewTaskRequestContractFile) SetFileType(v string) *CreateContractAppsReviewTaskRequestContractFile {
	s.FileType = &v
	return s
}

func (s *CreateContractAppsReviewTaskRequestContractFile) SetSpaceId(v string) *CreateContractAppsReviewTaskRequestContractFile {
	s.SpaceId = &v
	return s
}

type CreateContractAppsReviewTaskRequestReviewCustomRules struct {
	RiskLevel    *string `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
	RuleDesc     *string `json:"ruleDesc,omitempty" xml:"ruleDesc,omitempty"`
	RuleSequence *string `json:"ruleSequence,omitempty" xml:"ruleSequence,omitempty"`
	RuleTag      *string `json:"ruleTag,omitempty" xml:"ruleTag,omitempty"`
	RuleTitle    *string `json:"ruleTitle,omitempty" xml:"ruleTitle,omitempty"`
}

func (s CreateContractAppsReviewTaskRequestReviewCustomRules) String() string {
	return tea.Prettify(s)
}

func (s CreateContractAppsReviewTaskRequestReviewCustomRules) GoString() string {
	return s.String()
}

func (s *CreateContractAppsReviewTaskRequestReviewCustomRules) SetRiskLevel(v string) *CreateContractAppsReviewTaskRequestReviewCustomRules {
	s.RiskLevel = &v
	return s
}

func (s *CreateContractAppsReviewTaskRequestReviewCustomRules) SetRuleDesc(v string) *CreateContractAppsReviewTaskRequestReviewCustomRules {
	s.RuleDesc = &v
	return s
}

func (s *CreateContractAppsReviewTaskRequestReviewCustomRules) SetRuleSequence(v string) *CreateContractAppsReviewTaskRequestReviewCustomRules {
	s.RuleSequence = &v
	return s
}

func (s *CreateContractAppsReviewTaskRequestReviewCustomRules) SetRuleTag(v string) *CreateContractAppsReviewTaskRequestReviewCustomRules {
	s.RuleTag = &v
	return s
}

func (s *CreateContractAppsReviewTaskRequestReviewCustomRules) SetRuleTitle(v string) *CreateContractAppsReviewTaskRequestReviewCustomRules {
	s.RuleTitle = &v
	return s
}

type CreateContractAppsReviewTaskResponseBody struct {
	Result  *CreateContractAppsReviewTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateContractAppsReviewTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateContractAppsReviewTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContractAppsReviewTaskResponseBody) SetResult(v *CreateContractAppsReviewTaskResponseBodyResult) *CreateContractAppsReviewTaskResponseBody {
	s.Result = v
	return s
}

func (s *CreateContractAppsReviewTaskResponseBody) SetSuccess(v bool) *CreateContractAppsReviewTaskResponseBody {
	s.Success = &v
	return s
}

type CreateContractAppsReviewTaskResponseBodyResult struct {
	Data      *CreateContractAppsReviewTaskResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateContractAppsReviewTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateContractAppsReviewTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateContractAppsReviewTaskResponseBodyResult) SetData(v *CreateContractAppsReviewTaskResponseBodyResultData) *CreateContractAppsReviewTaskResponseBodyResult {
	s.Data = v
	return s
}

func (s *CreateContractAppsReviewTaskResponseBodyResult) SetRequestId(v string) *CreateContractAppsReviewTaskResponseBodyResult {
	s.RequestId = &v
	return s
}

type CreateContractAppsReviewTaskResponseBodyResultData struct {
	ReviewTaskId *string `json:"reviewTaskId,omitempty" xml:"reviewTaskId,omitempty"`
}

func (s CreateContractAppsReviewTaskResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s CreateContractAppsReviewTaskResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *CreateContractAppsReviewTaskResponseBodyResultData) SetReviewTaskId(v string) *CreateContractAppsReviewTaskResponseBodyResultData {
	s.ReviewTaskId = &v
	return s
}

type CreateContractAppsReviewTaskResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateContractAppsReviewTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateContractAppsReviewTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateContractAppsReviewTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateContractAppsReviewTaskResponse) SetHeaders(v map[string]*string) *CreateContractAppsReviewTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateContractAppsReviewTaskResponse) SetStatusCode(v int32) *CreateContractAppsReviewTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateContractAppsReviewTaskResponse) SetBody(v *CreateContractAppsReviewTaskResponseBody) *CreateContractAppsReviewTaskResponse {
	s.Body = v
	return s
}

type CreateContractCompareTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateContractCompareTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateContractCompareTaskHeaders) GoString() string {
	return s.String()
}

func (s *CreateContractCompareTaskHeaders) SetCommonHeaders(v map[string]*string) *CreateContractCompareTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateContractCompareTaskHeaders) SetXAcsDingtalkAccessToken(v string) *CreateContractCompareTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateContractCompareTaskRequest struct {
	ComparativeFile            *CreateContractCompareTaskRequestComparativeFile `json:"comparativeFile,omitempty" xml:"comparativeFile,omitempty" type:"Struct"`
	ComparativeFileDownloadUrl *string                                          `json:"comparativeFileDownloadUrl,omitempty" xml:"comparativeFileDownloadUrl,omitempty"`
	// This parameter is required.
	ComparativeFileName *string `json:"comparativeFileName,omitempty" xml:"comparativeFileName,omitempty"`
	// This parameter is required.
	FileSource *string `json:"fileSource,omitempty" xml:"fileSource,omitempty"`
	// This parameter is required.
	RequestId               *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	StandardFile            *CreateContractCompareTaskRequestStandardFile `json:"standardFile,omitempty" xml:"standardFile,omitempty" type:"Struct"`
	StandardFileDownloadUrl *string                                       `json:"standardFileDownloadUrl,omitempty" xml:"standardFileDownloadUrl,omitempty"`
	// This parameter is required.
	StandardFileName *string `json:"standardFileName,omitempty" xml:"standardFileName,omitempty"`
}

func (s CreateContractCompareTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateContractCompareTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateContractCompareTaskRequest) SetComparativeFile(v *CreateContractCompareTaskRequestComparativeFile) *CreateContractCompareTaskRequest {
	s.ComparativeFile = v
	return s
}

func (s *CreateContractCompareTaskRequest) SetComparativeFileDownloadUrl(v string) *CreateContractCompareTaskRequest {
	s.ComparativeFileDownloadUrl = &v
	return s
}

func (s *CreateContractCompareTaskRequest) SetComparativeFileName(v string) *CreateContractCompareTaskRequest {
	s.ComparativeFileName = &v
	return s
}

func (s *CreateContractCompareTaskRequest) SetFileSource(v string) *CreateContractCompareTaskRequest {
	s.FileSource = &v
	return s
}

func (s *CreateContractCompareTaskRequest) SetRequestId(v string) *CreateContractCompareTaskRequest {
	s.RequestId = &v
	return s
}

func (s *CreateContractCompareTaskRequest) SetStandardFile(v *CreateContractCompareTaskRequestStandardFile) *CreateContractCompareTaskRequest {
	s.StandardFile = v
	return s
}

func (s *CreateContractCompareTaskRequest) SetStandardFileDownloadUrl(v string) *CreateContractCompareTaskRequest {
	s.StandardFileDownloadUrl = &v
	return s
}

func (s *CreateContractCompareTaskRequest) SetStandardFileName(v string) *CreateContractCompareTaskRequest {
	s.StandardFileName = &v
	return s
}

type CreateContractCompareTaskRequestComparativeFile struct {
	FileId   *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileSize *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	SpaceId  *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s CreateContractCompareTaskRequestComparativeFile) String() string {
	return tea.Prettify(s)
}

func (s CreateContractCompareTaskRequestComparativeFile) GoString() string {
	return s.String()
}

func (s *CreateContractCompareTaskRequestComparativeFile) SetFileId(v string) *CreateContractCompareTaskRequestComparativeFile {
	s.FileId = &v
	return s
}

func (s *CreateContractCompareTaskRequestComparativeFile) SetFileName(v string) *CreateContractCompareTaskRequestComparativeFile {
	s.FileName = &v
	return s
}

func (s *CreateContractCompareTaskRequestComparativeFile) SetFileSize(v int64) *CreateContractCompareTaskRequestComparativeFile {
	s.FileSize = &v
	return s
}

func (s *CreateContractCompareTaskRequestComparativeFile) SetFileType(v string) *CreateContractCompareTaskRequestComparativeFile {
	s.FileType = &v
	return s
}

func (s *CreateContractCompareTaskRequestComparativeFile) SetSpaceId(v string) *CreateContractCompareTaskRequestComparativeFile {
	s.SpaceId = &v
	return s
}

type CreateContractCompareTaskRequestStandardFile struct {
	FileId   *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileSize *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	SpaceId  *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s CreateContractCompareTaskRequestStandardFile) String() string {
	return tea.Prettify(s)
}

func (s CreateContractCompareTaskRequestStandardFile) GoString() string {
	return s.String()
}

func (s *CreateContractCompareTaskRequestStandardFile) SetFileId(v string) *CreateContractCompareTaskRequestStandardFile {
	s.FileId = &v
	return s
}

func (s *CreateContractCompareTaskRequestStandardFile) SetFileName(v string) *CreateContractCompareTaskRequestStandardFile {
	s.FileName = &v
	return s
}

func (s *CreateContractCompareTaskRequestStandardFile) SetFileSize(v int64) *CreateContractCompareTaskRequestStandardFile {
	s.FileSize = &v
	return s
}

func (s *CreateContractCompareTaskRequestStandardFile) SetFileType(v string) *CreateContractCompareTaskRequestStandardFile {
	s.FileType = &v
	return s
}

func (s *CreateContractCompareTaskRequestStandardFile) SetSpaceId(v string) *CreateContractCompareTaskRequestStandardFile {
	s.SpaceId = &v
	return s
}

type CreateContractCompareTaskResponseBody struct {
	Result  *CreateContractCompareTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                        `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateContractCompareTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateContractCompareTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContractCompareTaskResponseBody) SetResult(v *CreateContractCompareTaskResponseBodyResult) *CreateContractCompareTaskResponseBody {
	s.Result = v
	return s
}

func (s *CreateContractCompareTaskResponseBody) SetSuccess(v bool) *CreateContractCompareTaskResponseBody {
	s.Success = &v
	return s
}

type CreateContractCompareTaskResponseBodyResult struct {
	Data      *CreateContractCompareTaskResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateContractCompareTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateContractCompareTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateContractCompareTaskResponseBodyResult) SetData(v *CreateContractCompareTaskResponseBodyResultData) *CreateContractCompareTaskResponseBodyResult {
	s.Data = v
	return s
}

func (s *CreateContractCompareTaskResponseBodyResult) SetRequestId(v string) *CreateContractCompareTaskResponseBodyResult {
	s.RequestId = &v
	return s
}

type CreateContractCompareTaskResponseBodyResultData struct {
	CompareTaskId *string `json:"compareTaskId,omitempty" xml:"compareTaskId,omitempty"`
}

func (s CreateContractCompareTaskResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s CreateContractCompareTaskResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *CreateContractCompareTaskResponseBodyResultData) SetCompareTaskId(v string) *CreateContractCompareTaskResponseBodyResultData {
	s.CompareTaskId = &v
	return s
}

type CreateContractCompareTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateContractCompareTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateContractCompareTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateContractCompareTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateContractCompareTaskResponse) SetHeaders(v map[string]*string) *CreateContractCompareTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateContractCompareTaskResponse) SetStatusCode(v int32) *CreateContractCompareTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateContractCompareTaskResponse) SetBody(v *CreateContractCompareTaskResponseBody) *CreateContractCompareTaskResponse {
	s.Body = v
	return s
}

type CreateContractExtractTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateContractExtractTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateContractExtractTaskHeaders) GoString() string {
	return s.String()
}

func (s *CreateContractExtractTaskHeaders) SetCommonHeaders(v map[string]*string) *CreateContractExtractTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateContractExtractTaskHeaders) SetXAcsDingtalkAccessToken(v string) *CreateContractExtractTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateContractExtractTaskRequest struct {
	// if can be null:
	// false
	ContractFile            *CreateContractExtractTaskRequestContractFile `json:"contractFile,omitempty" xml:"contractFile,omitempty" type:"Struct"`
	ContractFileDownloadUrl *string                                       `json:"contractFileDownloadUrl,omitempty" xml:"contractFileDownloadUrl,omitempty"`
	// This parameter is required.
	ContractFileName *string `json:"contractFileName,omitempty" xml:"contractFileName,omitempty"`
	// This parameter is required.
	ExtractKeys []*string `json:"extractKeys,omitempty" xml:"extractKeys,omitempty" type:"Repeated"`
	// This parameter is required.
	FileSource *string `json:"fileSource,omitempty" xml:"fileSource,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateContractExtractTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateContractExtractTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateContractExtractTaskRequest) SetContractFile(v *CreateContractExtractTaskRequestContractFile) *CreateContractExtractTaskRequest {
	s.ContractFile = v
	return s
}

func (s *CreateContractExtractTaskRequest) SetContractFileDownloadUrl(v string) *CreateContractExtractTaskRequest {
	s.ContractFileDownloadUrl = &v
	return s
}

func (s *CreateContractExtractTaskRequest) SetContractFileName(v string) *CreateContractExtractTaskRequest {
	s.ContractFileName = &v
	return s
}

func (s *CreateContractExtractTaskRequest) SetExtractKeys(v []*string) *CreateContractExtractTaskRequest {
	s.ExtractKeys = v
	return s
}

func (s *CreateContractExtractTaskRequest) SetFileSource(v string) *CreateContractExtractTaskRequest {
	s.FileSource = &v
	return s
}

func (s *CreateContractExtractTaskRequest) SetRequestId(v string) *CreateContractExtractTaskRequest {
	s.RequestId = &v
	return s
}

type CreateContractExtractTaskRequestContractFile struct {
	FileId   *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileSize *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	SpaceId  *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s CreateContractExtractTaskRequestContractFile) String() string {
	return tea.Prettify(s)
}

func (s CreateContractExtractTaskRequestContractFile) GoString() string {
	return s.String()
}

func (s *CreateContractExtractTaskRequestContractFile) SetFileId(v string) *CreateContractExtractTaskRequestContractFile {
	s.FileId = &v
	return s
}

func (s *CreateContractExtractTaskRequestContractFile) SetFileName(v string) *CreateContractExtractTaskRequestContractFile {
	s.FileName = &v
	return s
}

func (s *CreateContractExtractTaskRequestContractFile) SetFileSize(v int64) *CreateContractExtractTaskRequestContractFile {
	s.FileSize = &v
	return s
}

func (s *CreateContractExtractTaskRequestContractFile) SetFileType(v string) *CreateContractExtractTaskRequestContractFile {
	s.FileType = &v
	return s
}

func (s *CreateContractExtractTaskRequestContractFile) SetSpaceId(v string) *CreateContractExtractTaskRequestContractFile {
	s.SpaceId = &v
	return s
}

type CreateContractExtractTaskResponseBody struct {
	Result  *CreateContractExtractTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                        `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateContractExtractTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateContractExtractTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContractExtractTaskResponseBody) SetResult(v *CreateContractExtractTaskResponseBodyResult) *CreateContractExtractTaskResponseBody {
	s.Result = v
	return s
}

func (s *CreateContractExtractTaskResponseBody) SetSuccess(v bool) *CreateContractExtractTaskResponseBody {
	s.Success = &v
	return s
}

type CreateContractExtractTaskResponseBodyResult struct {
	Data      *CreateContractExtractTaskResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateContractExtractTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateContractExtractTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateContractExtractTaskResponseBodyResult) SetData(v *CreateContractExtractTaskResponseBodyResultData) *CreateContractExtractTaskResponseBodyResult {
	s.Data = v
	return s
}

func (s *CreateContractExtractTaskResponseBodyResult) SetRequestId(v string) *CreateContractExtractTaskResponseBodyResult {
	s.RequestId = &v
	return s
}

type CreateContractExtractTaskResponseBodyResultData struct {
	ExtractTaskId *string `json:"extractTaskId,omitempty" xml:"extractTaskId,omitempty"`
}

func (s CreateContractExtractTaskResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s CreateContractExtractTaskResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *CreateContractExtractTaskResponseBodyResultData) SetExtractTaskId(v string) *CreateContractExtractTaskResponseBodyResultData {
	s.ExtractTaskId = &v
	return s
}

type CreateContractExtractTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateContractExtractTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateContractExtractTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateContractExtractTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateContractExtractTaskResponse) SetHeaders(v map[string]*string) *CreateContractExtractTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateContractExtractTaskResponse) SetStatusCode(v int32) *CreateContractExtractTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateContractExtractTaskResponse) SetBody(v *CreateContractExtractTaskResponseBody) *CreateContractExtractTaskResponse {
	s.Body = v
	return s
}

type CreateContractReviewTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateContractReviewTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateContractReviewTaskHeaders) GoString() string {
	return s.String()
}

func (s *CreateContractReviewTaskHeaders) SetCommonHeaders(v map[string]*string) *CreateContractReviewTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateContractReviewTaskHeaders) SetXAcsDingtalkAccessToken(v string) *CreateContractReviewTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateContractReviewTaskRequest struct {
	// if can be null:
	// false
	ContractFile            *CreateContractReviewTaskRequestContractFile `json:"contractFile,omitempty" xml:"contractFile,omitempty" type:"Struct"`
	ContractFileDownloadUrl *string                                      `json:"contractFileDownloadUrl,omitempty" xml:"contractFileDownloadUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 采购合同.doc
	ContractFileName *string `json:"contractFileName,omitempty" xml:"contractFileName,omitempty"`
	// This parameter is required.
	FileSource *string `json:"fileSource,omitempty" xml:"fileSource,omitempty"`
	// This parameter is required.
	RequestId         *string                                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ReviewCustomRules []*CreateContractReviewTaskRequestReviewCustomRules `json:"reviewCustomRules,omitempty" xml:"reviewCustomRules,omitempty" type:"Repeated"`
	// example:
	//
	// model
	RuleType *string `json:"ruleType,omitempty" xml:"ruleType,omitempty"`
	// example:
	//
	// 0
	Standpoint *string `json:"standpoint,omitempty" xml:"standpoint,omitempty"`
}

func (s CreateContractReviewTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateContractReviewTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateContractReviewTaskRequest) SetContractFile(v *CreateContractReviewTaskRequestContractFile) *CreateContractReviewTaskRequest {
	s.ContractFile = v
	return s
}

func (s *CreateContractReviewTaskRequest) SetContractFileDownloadUrl(v string) *CreateContractReviewTaskRequest {
	s.ContractFileDownloadUrl = &v
	return s
}

func (s *CreateContractReviewTaskRequest) SetContractFileName(v string) *CreateContractReviewTaskRequest {
	s.ContractFileName = &v
	return s
}

func (s *CreateContractReviewTaskRequest) SetFileSource(v string) *CreateContractReviewTaskRequest {
	s.FileSource = &v
	return s
}

func (s *CreateContractReviewTaskRequest) SetRequestId(v string) *CreateContractReviewTaskRequest {
	s.RequestId = &v
	return s
}

func (s *CreateContractReviewTaskRequest) SetReviewCustomRules(v []*CreateContractReviewTaskRequestReviewCustomRules) *CreateContractReviewTaskRequest {
	s.ReviewCustomRules = v
	return s
}

func (s *CreateContractReviewTaskRequest) SetRuleType(v string) *CreateContractReviewTaskRequest {
	s.RuleType = &v
	return s
}

func (s *CreateContractReviewTaskRequest) SetStandpoint(v string) *CreateContractReviewTaskRequest {
	s.Standpoint = &v
	return s
}

type CreateContractReviewTaskRequestContractFile struct {
	FileId   *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileSize *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	SpaceId  *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s CreateContractReviewTaskRequestContractFile) String() string {
	return tea.Prettify(s)
}

func (s CreateContractReviewTaskRequestContractFile) GoString() string {
	return s.String()
}

func (s *CreateContractReviewTaskRequestContractFile) SetFileId(v string) *CreateContractReviewTaskRequestContractFile {
	s.FileId = &v
	return s
}

func (s *CreateContractReviewTaskRequestContractFile) SetFileName(v string) *CreateContractReviewTaskRequestContractFile {
	s.FileName = &v
	return s
}

func (s *CreateContractReviewTaskRequestContractFile) SetFileSize(v int64) *CreateContractReviewTaskRequestContractFile {
	s.FileSize = &v
	return s
}

func (s *CreateContractReviewTaskRequestContractFile) SetFileType(v string) *CreateContractReviewTaskRequestContractFile {
	s.FileType = &v
	return s
}

func (s *CreateContractReviewTaskRequestContractFile) SetSpaceId(v string) *CreateContractReviewTaskRequestContractFile {
	s.SpaceId = &v
	return s
}

type CreateContractReviewTaskRequestReviewCustomRules struct {
	// example:
	//
	// high
	RiskLevel *string `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
	// example:
	//
	// 审查合同金额大小，不得低于1000元
	RuleDesc *string `json:"ruleDesc,omitempty" xml:"ruleDesc,omitempty"`
	// example:
	//
	// 1.1
	RuleSequence *string `json:"ruleSequence,omitempty" xml:"ruleSequence,omitempty"`
	// example:
	//
	// 金额相关规则
	RuleTag *string `json:"ruleTag,omitempty" xml:"ruleTag,omitempty"`
	// example:
	//
	// 合同金额校验
	RuleTitle *string `json:"ruleTitle,omitempty" xml:"ruleTitle,omitempty"`
}

func (s CreateContractReviewTaskRequestReviewCustomRules) String() string {
	return tea.Prettify(s)
}

func (s CreateContractReviewTaskRequestReviewCustomRules) GoString() string {
	return s.String()
}

func (s *CreateContractReviewTaskRequestReviewCustomRules) SetRiskLevel(v string) *CreateContractReviewTaskRequestReviewCustomRules {
	s.RiskLevel = &v
	return s
}

func (s *CreateContractReviewTaskRequestReviewCustomRules) SetRuleDesc(v string) *CreateContractReviewTaskRequestReviewCustomRules {
	s.RuleDesc = &v
	return s
}

func (s *CreateContractReviewTaskRequestReviewCustomRules) SetRuleSequence(v string) *CreateContractReviewTaskRequestReviewCustomRules {
	s.RuleSequence = &v
	return s
}

func (s *CreateContractReviewTaskRequestReviewCustomRules) SetRuleTag(v string) *CreateContractReviewTaskRequestReviewCustomRules {
	s.RuleTag = &v
	return s
}

func (s *CreateContractReviewTaskRequestReviewCustomRules) SetRuleTitle(v string) *CreateContractReviewTaskRequestReviewCustomRules {
	s.RuleTitle = &v
	return s
}

type CreateContractReviewTaskResponseBody struct {
	Result  *CreateContractReviewTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                       `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateContractReviewTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateContractReviewTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContractReviewTaskResponseBody) SetResult(v *CreateContractReviewTaskResponseBodyResult) *CreateContractReviewTaskResponseBody {
	s.Result = v
	return s
}

func (s *CreateContractReviewTaskResponseBody) SetSuccess(v bool) *CreateContractReviewTaskResponseBody {
	s.Success = &v
	return s
}

type CreateContractReviewTaskResponseBodyResult struct {
	Data      *CreateContractReviewTaskResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateContractReviewTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateContractReviewTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateContractReviewTaskResponseBodyResult) SetData(v *CreateContractReviewTaskResponseBodyResultData) *CreateContractReviewTaskResponseBodyResult {
	s.Data = v
	return s
}

func (s *CreateContractReviewTaskResponseBodyResult) SetRequestId(v string) *CreateContractReviewTaskResponseBodyResult {
	s.RequestId = &v
	return s
}

type CreateContractReviewTaskResponseBodyResultData struct {
	ReviewTaskId *string `json:"reviewTaskId,omitempty" xml:"reviewTaskId,omitempty"`
}

func (s CreateContractReviewTaskResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s CreateContractReviewTaskResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *CreateContractReviewTaskResponseBodyResultData) SetReviewTaskId(v string) *CreateContractReviewTaskResponseBodyResultData {
	s.ReviewTaskId = &v
	return s
}

type CreateContractReviewTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateContractReviewTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateContractReviewTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateContractReviewTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateContractReviewTaskResponse) SetHeaders(v map[string]*string) *CreateContractReviewTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateContractReviewTaskResponse) SetStatusCode(v int32) *CreateContractReviewTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateContractReviewTaskResponse) SetBody(v *CreateContractReviewTaskResponseBody) *CreateContractReviewTaskResponse {
	s.Body = v
	return s
}

type EsignQueryApprovalInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EsignQueryApprovalInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryApprovalInfoHeaders) GoString() string {
	return s.String()
}

func (s *EsignQueryApprovalInfoHeaders) SetCommonHeaders(v map[string]*string) *EsignQueryApprovalInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EsignQueryApprovalInfoHeaders) SetXAcsDingtalkAccessToken(v string) *EsignQueryApprovalInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EsignQueryApprovalInfoRequest struct {
	// example:
	//
	// dingd0c871e2dfc941a34ac5d6980864d335
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// 5556ae0359c64c4b9491c0c3c339341f
	EsignFlowId *string `json:"esignFlowId,omitempty" xml:"esignFlowId,omitempty"`
	// example:
	//
	// PbnhW6rVXRg8u6T4NiiOwwQiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s EsignQueryApprovalInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryApprovalInfoRequest) GoString() string {
	return s.String()
}

func (s *EsignQueryApprovalInfoRequest) SetCorpId(v string) *EsignQueryApprovalInfoRequest {
	s.CorpId = &v
	return s
}

func (s *EsignQueryApprovalInfoRequest) SetEsignFlowId(v string) *EsignQueryApprovalInfoRequest {
	s.EsignFlowId = &v
	return s
}

func (s *EsignQueryApprovalInfoRequest) SetUnionId(v string) *EsignQueryApprovalInfoRequest {
	s.UnionId = &v
	return s
}

type EsignQueryApprovalInfoResponseBody struct {
	Result  *EsignQueryApprovalInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s EsignQueryApprovalInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryApprovalInfoResponseBody) GoString() string {
	return s.String()
}

func (s *EsignQueryApprovalInfoResponseBody) SetResult(v *EsignQueryApprovalInfoResponseBodyResult) *EsignQueryApprovalInfoResponseBody {
	s.Result = v
	return s
}

func (s *EsignQueryApprovalInfoResponseBody) SetSuccess(v bool) *EsignQueryApprovalInfoResponseBody {
	s.Success = &v
	return s
}

type EsignQueryApprovalInfoResponseBodyResult struct {
	// example:
	//
	// 202311081619000455501
	BpmsProcessBusinessId *string `json:"bpmsProcessBusinessId,omitempty" xml:"bpmsProcessBusinessId,omitempty"`
	// example:
	//
	// O6wNhB4wTMajvNW8Dc_Rjg09301699431585
	BpmsProcessInstanceId *string `json:"bpmsProcessInstanceId,omitempty" xml:"bpmsProcessInstanceId,omitempty"`
	// example:
	//
	// https://aflow.dingtalk.com/dingtalk/pc/query/pchomepage.htm?corpid=ding6c3948a9e37c439c35c2f4657eb6378f&swfrom=https://n.dingtalk.com/dingding/h5-contract/contractPC/index.html#/approval?procInstId=O6wNhB4wTMajvNW8Dc_Rjg09301699431585
	BpmsProcessInstanceUrl *string `json:"bpmsProcessInstanceUrl,omitempty" xml:"bpmsProcessInstanceUrl,omitempty"`
}

func (s EsignQueryApprovalInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryApprovalInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *EsignQueryApprovalInfoResponseBodyResult) SetBpmsProcessBusinessId(v string) *EsignQueryApprovalInfoResponseBodyResult {
	s.BpmsProcessBusinessId = &v
	return s
}

func (s *EsignQueryApprovalInfoResponseBodyResult) SetBpmsProcessInstanceId(v string) *EsignQueryApprovalInfoResponseBodyResult {
	s.BpmsProcessInstanceId = &v
	return s
}

func (s *EsignQueryApprovalInfoResponseBodyResult) SetBpmsProcessInstanceUrl(v string) *EsignQueryApprovalInfoResponseBodyResult {
	s.BpmsProcessInstanceUrl = &v
	return s
}

type EsignQueryApprovalInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EsignQueryApprovalInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EsignQueryApprovalInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryApprovalInfoResponse) GoString() string {
	return s.String()
}

func (s *EsignQueryApprovalInfoResponse) SetHeaders(v map[string]*string) *EsignQueryApprovalInfoResponse {
	s.Headers = v
	return s
}

func (s *EsignQueryApprovalInfoResponse) SetStatusCode(v int32) *EsignQueryApprovalInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *EsignQueryApprovalInfoResponse) SetBody(v *EsignQueryApprovalInfoResponseBody) *EsignQueryApprovalInfoResponse {
	s.Body = v
	return s
}

type EsignQueryGrantInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EsignQueryGrantInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryGrantInfoHeaders) GoString() string {
	return s.String()
}

func (s *EsignQueryGrantInfoHeaders) SetCommonHeaders(v map[string]*string) *EsignQueryGrantInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EsignQueryGrantInfoHeaders) SetXAcsDingtalkAccessToken(v string) *EsignQueryGrantInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EsignQueryGrantInfoRequest struct {
	// example:
	//
	// dingd0c871e2dfc941a34ac5d6980864d335
	CorpId    *string            `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Extension map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
	// example:
	//
	// PbnhW6rVXRg8u6T4NiiOwwQiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s EsignQueryGrantInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryGrantInfoRequest) GoString() string {
	return s.String()
}

func (s *EsignQueryGrantInfoRequest) SetCorpId(v string) *EsignQueryGrantInfoRequest {
	s.CorpId = &v
	return s
}

func (s *EsignQueryGrantInfoRequest) SetExtension(v map[string]*string) *EsignQueryGrantInfoRequest {
	s.Extension = v
	return s
}

func (s *EsignQueryGrantInfoRequest) SetUnionId(v string) *EsignQueryGrantInfoRequest {
	s.UnionId = &v
	return s
}

type EsignQueryGrantInfoResponseBody struct {
	Result  *EsignQueryGrantInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s EsignQueryGrantInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryGrantInfoResponseBody) GoString() string {
	return s.String()
}

func (s *EsignQueryGrantInfoResponseBody) SetResult(v *EsignQueryGrantInfoResponseBodyResult) *EsignQueryGrantInfoResponseBody {
	s.Result = v
	return s
}

func (s *EsignQueryGrantInfoResponseBody) SetSuccess(v bool) *EsignQueryGrantInfoResponseBody {
	s.Success = &v
	return s
}

type EsignQueryGrantInfoResponseBodyResult struct {
	// example:
	//
	// 何一兵
	LegalPerson *string `json:"legalPerson,omitempty" xml:"legalPerson,omitempty"`
	// example:
	//
	// 18667021101
	MobilePhoneNumber *string `json:"mobilePhoneNumber,omitempty" xml:"mobilePhoneNumber,omitempty"`
	// example:
	//
	// 杭州天谷信息科技有限公司
	OrgName *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	// example:
	//
	// 86
	StateCode *string `json:"stateCode,omitempty" xml:"stateCode,omitempty"`
	// example:
	//
	// 913301087458306077
	UnifiedSocialCredit *string `json:"unifiedSocialCredit,omitempty" xml:"unifiedSocialCredit,omitempty"`
	// example:
	//
	// 某人名
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s EsignQueryGrantInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryGrantInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *EsignQueryGrantInfoResponseBodyResult) SetLegalPerson(v string) *EsignQueryGrantInfoResponseBodyResult {
	s.LegalPerson = &v
	return s
}

func (s *EsignQueryGrantInfoResponseBodyResult) SetMobilePhoneNumber(v string) *EsignQueryGrantInfoResponseBodyResult {
	s.MobilePhoneNumber = &v
	return s
}

func (s *EsignQueryGrantInfoResponseBodyResult) SetOrgName(v string) *EsignQueryGrantInfoResponseBodyResult {
	s.OrgName = &v
	return s
}

func (s *EsignQueryGrantInfoResponseBodyResult) SetStateCode(v string) *EsignQueryGrantInfoResponseBodyResult {
	s.StateCode = &v
	return s
}

func (s *EsignQueryGrantInfoResponseBodyResult) SetUnifiedSocialCredit(v string) *EsignQueryGrantInfoResponseBodyResult {
	s.UnifiedSocialCredit = &v
	return s
}

func (s *EsignQueryGrantInfoResponseBodyResult) SetUserName(v string) *EsignQueryGrantInfoResponseBodyResult {
	s.UserName = &v
	return s
}

type EsignQueryGrantInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EsignQueryGrantInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EsignQueryGrantInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryGrantInfoResponse) GoString() string {
	return s.String()
}

func (s *EsignQueryGrantInfoResponse) SetHeaders(v map[string]*string) *EsignQueryGrantInfoResponse {
	s.Headers = v
	return s
}

func (s *EsignQueryGrantInfoResponse) SetStatusCode(v int32) *EsignQueryGrantInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *EsignQueryGrantInfoResponse) SetBody(v *EsignQueryGrantInfoResponseBody) *EsignQueryGrantInfoResponse {
	s.Body = v
	return s
}

type EsignQueryIdentityByTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EsignQueryIdentityByTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryIdentityByTicketHeaders) GoString() string {
	return s.String()
}

func (s *EsignQueryIdentityByTicketHeaders) SetCommonHeaders(v map[string]*string) *EsignQueryIdentityByTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EsignQueryIdentityByTicketHeaders) SetXAcsDingtalkAccessToken(v string) *EsignQueryIdentityByTicketHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EsignQueryIdentityByTicketRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dingd0c871e2dfc941a34ac5d6980864d335
	CorpId    *string            `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Extension map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// feb4b8e5-d6d9-4d22-a6b8-c8e26823a73a
	Ticket *string `json:"ticket,omitempty" xml:"ticket,omitempty"`
}

func (s EsignQueryIdentityByTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryIdentityByTicketRequest) GoString() string {
	return s.String()
}

func (s *EsignQueryIdentityByTicketRequest) SetCorpId(v string) *EsignQueryIdentityByTicketRequest {
	s.CorpId = &v
	return s
}

func (s *EsignQueryIdentityByTicketRequest) SetExtension(v map[string]*string) *EsignQueryIdentityByTicketRequest {
	s.Extension = v
	return s
}

func (s *EsignQueryIdentityByTicketRequest) SetTicket(v string) *EsignQueryIdentityByTicketRequest {
	s.Ticket = &v
	return s
}

type EsignQueryIdentityByTicketResponseBody struct {
	Result  *EsignQueryIdentityByTicketResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s EsignQueryIdentityByTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryIdentityByTicketResponseBody) GoString() string {
	return s.String()
}

func (s *EsignQueryIdentityByTicketResponseBody) SetResult(v *EsignQueryIdentityByTicketResponseBodyResult) *EsignQueryIdentityByTicketResponseBody {
	s.Result = v
	return s
}

func (s *EsignQueryIdentityByTicketResponseBody) SetSuccess(v bool) *EsignQueryIdentityByTicketResponseBody {
	s.Success = &v
	return s
}

type EsignQueryIdentityByTicketResponseBodyResult struct {
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s EsignQueryIdentityByTicketResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryIdentityByTicketResponseBodyResult) GoString() string {
	return s.String()
}

func (s *EsignQueryIdentityByTicketResponseBodyResult) SetUnionId(v string) *EsignQueryIdentityByTicketResponseBodyResult {
	s.UnionId = &v
	return s
}

type EsignQueryIdentityByTicketResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EsignQueryIdentityByTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EsignQueryIdentityByTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s EsignQueryIdentityByTicketResponse) GoString() string {
	return s.String()
}

func (s *EsignQueryIdentityByTicketResponse) SetHeaders(v map[string]*string) *EsignQueryIdentityByTicketResponse {
	s.Headers = v
	return s
}

func (s *EsignQueryIdentityByTicketResponse) SetStatusCode(v int32) *EsignQueryIdentityByTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *EsignQueryIdentityByTicketResponse) SetBody(v *EsignQueryIdentityByTicketResponseBody) *EsignQueryIdentityByTicketResponse {
	s.Body = v
	return s
}

type EsignSyncEventHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EsignSyncEventHeaders) String() string {
	return tea.Prettify(s)
}

func (s EsignSyncEventHeaders) GoString() string {
	return s.String()
}

func (s *EsignSyncEventHeaders) SetCommonHeaders(v map[string]*string) *EsignSyncEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EsignSyncEventHeaders) SetXAcsDingtalkAccessToken(v string) *EsignSyncEventHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EsignSyncEventRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// openEsign
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// example:
	//
	// dingd0c871e2dfc941a34ac5d6980864d335
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// {"name": "名字"}
	EsignData *string            `json:"esignData,omitempty" xml:"esignData,omitempty"`
	Extension map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
	// example:
	//
	// PbnhW6rVXRg8u6T4NiiOwwQiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s EsignSyncEventRequest) String() string {
	return tea.Prettify(s)
}

func (s EsignSyncEventRequest) GoString() string {
	return s.String()
}

func (s *EsignSyncEventRequest) SetAction(v string) *EsignSyncEventRequest {
	s.Action = &v
	return s
}

func (s *EsignSyncEventRequest) SetCorpId(v string) *EsignSyncEventRequest {
	s.CorpId = &v
	return s
}

func (s *EsignSyncEventRequest) SetEsignData(v string) *EsignSyncEventRequest {
	s.EsignData = &v
	return s
}

func (s *EsignSyncEventRequest) SetExtension(v map[string]*string) *EsignSyncEventRequest {
	s.Extension = v
	return s
}

func (s *EsignSyncEventRequest) SetUnionId(v string) *EsignSyncEventRequest {
	s.UnionId = &v
	return s
}

type EsignSyncEventResponseBody struct {
	Result  *EsignSyncEventResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s EsignSyncEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EsignSyncEventResponseBody) GoString() string {
	return s.String()
}

func (s *EsignSyncEventResponseBody) SetResult(v *EsignSyncEventResponseBodyResult) *EsignSyncEventResponseBody {
	s.Result = v
	return s
}

func (s *EsignSyncEventResponseBody) SetSuccess(v bool) *EsignSyncEventResponseBody {
	s.Success = &v
	return s
}

type EsignSyncEventResponseBodyResult struct {
	// example:
	//
	// 外部服务异常
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s EsignSyncEventResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s EsignSyncEventResponseBodyResult) GoString() string {
	return s.String()
}

func (s *EsignSyncEventResponseBodyResult) SetMessage(v string) *EsignSyncEventResponseBodyResult {
	s.Message = &v
	return s
}

type EsignSyncEventResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EsignSyncEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EsignSyncEventResponse) String() string {
	return tea.Prettify(s)
}

func (s EsignSyncEventResponse) GoString() string {
	return s.String()
}

func (s *EsignSyncEventResponse) SetHeaders(v map[string]*string) *EsignSyncEventResponse {
	s.Headers = v
	return s
}

func (s *EsignSyncEventResponse) SetStatusCode(v int32) *EsignSyncEventResponse {
	s.StatusCode = &v
	return s
}

func (s *EsignSyncEventResponse) SetBody(v *EsignSyncEventResponseBody) *EsignSyncEventResponse {
	s.Body = v
	return s
}

type EsignUserVerifyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EsignUserVerifyHeaders) String() string {
	return tea.Prettify(s)
}

func (s EsignUserVerifyHeaders) GoString() string {
	return s.String()
}

func (s *EsignUserVerifyHeaders) SetCommonHeaders(v map[string]*string) *EsignUserVerifyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EsignUserVerifyHeaders) SetXAcsDingtalkAccessToken(v string) *EsignUserVerifyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EsignUserVerifyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dingbfc7ac3ddee0c1acffe93478753d9884
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1Wgkx59vuKA8u6T4NiiOwwQiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s EsignUserVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s EsignUserVerifyRequest) GoString() string {
	return s.String()
}

func (s *EsignUserVerifyRequest) SetCorpId(v string) *EsignUserVerifyRequest {
	s.CorpId = &v
	return s
}

func (s *EsignUserVerifyRequest) SetUnionId(v string) *EsignUserVerifyRequest {
	s.UnionId = &v
	return s
}

type EsignUserVerifyResponseBody struct {
	Result *EsignUserVerifyResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// This parameter is required.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s EsignUserVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EsignUserVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *EsignUserVerifyResponseBody) SetResult(v *EsignUserVerifyResponseBodyResult) *EsignUserVerifyResponseBody {
	s.Result = v
	return s
}

func (s *EsignUserVerifyResponseBody) SetSuccess(v bool) *EsignUserVerifyResponseBody {
	s.Success = &v
	return s
}

type EsignUserVerifyResponseBodyResult struct {
	CanAccess *bool `json:"canAccess,omitempty" xml:"canAccess,omitempty"`
}

func (s EsignUserVerifyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s EsignUserVerifyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *EsignUserVerifyResponseBodyResult) SetCanAccess(v bool) *EsignUserVerifyResponseBodyResult {
	s.CanAccess = &v
	return s
}

type EsignUserVerifyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EsignUserVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EsignUserVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s EsignUserVerifyResponse) GoString() string {
	return s.String()
}

func (s *EsignUserVerifyResponse) SetHeaders(v map[string]*string) *EsignUserVerifyResponse {
	s.Headers = v
	return s
}

func (s *EsignUserVerifyResponse) SetStatusCode(v int32) *EsignUserVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *EsignUserVerifyResponse) SetBody(v *EsignUserVerifyResponseBody) *EsignUserVerifyResponse {
	s.Body = v
	return s
}

type FinishReviewOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s FinishReviewOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s FinishReviewOrderHeaders) GoString() string {
	return s.String()
}

func (s *FinishReviewOrderHeaders) SetCommonHeaders(v map[string]*string) *FinishReviewOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FinishReviewOrderHeaders) SetXAcsDingtalkAccessToken(v string) *FinishReviewOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type FinishReviewOrderRequest struct {
	EndFiles []*FinishReviewOrderRequestEndFiles `json:"endFiles,omitempty" xml:"endFiles,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// example:
	//
	// 12345678
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s FinishReviewOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s FinishReviewOrderRequest) GoString() string {
	return s.String()
}

func (s *FinishReviewOrderRequest) SetEndFiles(v []*FinishReviewOrderRequestEndFiles) *FinishReviewOrderRequest {
	s.EndFiles = v
	return s
}

func (s *FinishReviewOrderRequest) SetExtension(v string) *FinishReviewOrderRequest {
	s.Extension = &v
	return s
}

func (s *FinishReviewOrderRequest) SetOrderId(v string) *FinishReviewOrderRequest {
	s.OrderId = &v
	return s
}

type FinishReviewOrderRequestEndFiles struct {
	// example:
	//
	// 合同文件
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// 12
	FileSize *string `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// example:
	//
	// word
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// 0
	FileVersion *int32 `json:"fileVersion,omitempty" xml:"fileVersion,omitempty"`
	// example:
	//
	// http://oss.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s FinishReviewOrderRequestEndFiles) String() string {
	return tea.Prettify(s)
}

func (s FinishReviewOrderRequestEndFiles) GoString() string {
	return s.String()
}

func (s *FinishReviewOrderRequestEndFiles) SetFileName(v string) *FinishReviewOrderRequestEndFiles {
	s.FileName = &v
	return s
}

func (s *FinishReviewOrderRequestEndFiles) SetFileSize(v string) *FinishReviewOrderRequestEndFiles {
	s.FileSize = &v
	return s
}

func (s *FinishReviewOrderRequestEndFiles) SetFileType(v string) *FinishReviewOrderRequestEndFiles {
	s.FileType = &v
	return s
}

func (s *FinishReviewOrderRequestEndFiles) SetFileVersion(v int32) *FinishReviewOrderRequestEndFiles {
	s.FileVersion = &v
	return s
}

func (s *FinishReviewOrderRequestEndFiles) SetUrl(v string) *FinishReviewOrderRequestEndFiles {
	s.Url = &v
	return s
}

type FinishReviewOrderResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s FinishReviewOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FinishReviewOrderResponseBody) GoString() string {
	return s.String()
}

func (s *FinishReviewOrderResponseBody) SetSuccess(v bool) *FinishReviewOrderResponseBody {
	s.Success = &v
	return s
}

type FinishReviewOrderResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FinishReviewOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FinishReviewOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s FinishReviewOrderResponse) GoString() string {
	return s.String()
}

func (s *FinishReviewOrderResponse) SetHeaders(v map[string]*string) *FinishReviewOrderResponse {
	s.Headers = v
	return s
}

func (s *FinishReviewOrderResponse) SetStatusCode(v int32) *FinishReviewOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *FinishReviewOrderResponse) SetBody(v *FinishReviewOrderResponseBody) *FinishReviewOrderResponse {
	s.Body = v
	return s
}

type QueryAdvancedContractVersionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAdvancedContractVersionHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAdvancedContractVersionHeaders) GoString() string {
	return s.String()
}

func (s *QueryAdvancedContractVersionHeaders) SetCommonHeaders(v map[string]*string) *QueryAdvancedContractVersionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAdvancedContractVersionHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAdvancedContractVersionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAdvancedContractVersionRequest struct {
	// example:
	//
	// dingff048f704a8b6d8e4ac5d6980864d335
	CorpId    *string            `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Extension map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
}

func (s QueryAdvancedContractVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAdvancedContractVersionRequest) GoString() string {
	return s.String()
}

func (s *QueryAdvancedContractVersionRequest) SetCorpId(v string) *QueryAdvancedContractVersionRequest {
	s.CorpId = &v
	return s
}

func (s *QueryAdvancedContractVersionRequest) SetExtension(v map[string]*string) *QueryAdvancedContractVersionRequest {
	s.Extension = v
	return s
}

type QueryAdvancedContractVersionResponseBody struct {
	Result  *QueryAdvancedContractVersionResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryAdvancedContractVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAdvancedContractVersionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAdvancedContractVersionResponseBody) SetResult(v *QueryAdvancedContractVersionResponseBodyResult) *QueryAdvancedContractVersionResponseBody {
	s.Result = v
	return s
}

func (s *QueryAdvancedContractVersionResponseBody) SetSuccess(v bool) *QueryAdvancedContractVersionResponseBody {
	s.Success = &v
	return s
}

type QueryAdvancedContractVersionResponseBodyResult struct {
	EnableEsignAttachmentSign *bool              `json:"enableEsignAttachmentSign,omitempty" xml:"enableEsignAttachmentSign,omitempty"`
	Extension                 map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
	// example:
	//
	// advanced
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s QueryAdvancedContractVersionResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryAdvancedContractVersionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryAdvancedContractVersionResponseBodyResult) SetEnableEsignAttachmentSign(v bool) *QueryAdvancedContractVersionResponseBodyResult {
	s.EnableEsignAttachmentSign = &v
	return s
}

func (s *QueryAdvancedContractVersionResponseBodyResult) SetExtension(v map[string]*string) *QueryAdvancedContractVersionResponseBodyResult {
	s.Extension = v
	return s
}

func (s *QueryAdvancedContractVersionResponseBodyResult) SetVersion(v string) *QueryAdvancedContractVersionResponseBodyResult {
	s.Version = &v
	return s
}

type QueryAdvancedContractVersionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAdvancedContractVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAdvancedContractVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAdvancedContractVersionResponse) GoString() string {
	return s.String()
}

func (s *QueryAdvancedContractVersionResponse) SetHeaders(v map[string]*string) *QueryAdvancedContractVersionResponse {
	s.Headers = v
	return s
}

func (s *QueryAdvancedContractVersionResponse) SetStatusCode(v int32) *QueryAdvancedContractVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAdvancedContractVersionResponse) SetBody(v *QueryAdvancedContractVersionResponseBody) *QueryAdvancedContractVersionResponse {
	s.Body = v
	return s
}

type QueryContractAppsCompareResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryContractAppsCompareResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsCompareResultHeaders) GoString() string {
	return s.String()
}

func (s *QueryContractAppsCompareResultHeaders) SetCommonHeaders(v map[string]*string) *QueryContractAppsCompareResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryContractAppsCompareResultHeaders) SetXAcsDingtalkAccessToken(v string) *QueryContractAppsCompareResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryContractAppsCompareResultRequest struct {
	// This parameter is required.
	CompareTaskId *string `json:"compareTaskId,omitempty" xml:"compareTaskId,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryContractAppsCompareResultRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsCompareResultRequest) GoString() string {
	return s.String()
}

func (s *QueryContractAppsCompareResultRequest) SetCompareTaskId(v string) *QueryContractAppsCompareResultRequest {
	s.CompareTaskId = &v
	return s
}

func (s *QueryContractAppsCompareResultRequest) SetRequestId(v string) *QueryContractAppsCompareResultRequest {
	s.RequestId = &v
	return s
}

func (s *QueryContractAppsCompareResultRequest) SetUnionId(v string) *QueryContractAppsCompareResultRequest {
	s.UnionId = &v
	return s
}

type QueryContractAppsCompareResultResponseBody struct {
	Result  *QueryContractAppsCompareResultResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryContractAppsCompareResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsCompareResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryContractAppsCompareResultResponseBody) SetResult(v *QueryContractAppsCompareResultResponseBodyResult) *QueryContractAppsCompareResultResponseBody {
	s.Result = v
	return s
}

func (s *QueryContractAppsCompareResultResponseBody) SetSuccess(v bool) *QueryContractAppsCompareResultResponseBody {
	s.Success = &v
	return s
}

type QueryContractAppsCompareResultResponseBodyResult struct {
	Data      *QueryContractAppsCompareResultResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryContractAppsCompareResultResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsCompareResultResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryContractAppsCompareResultResponseBodyResult) SetData(v *QueryContractAppsCompareResultResponseBodyResultData) *QueryContractAppsCompareResultResponseBodyResult {
	s.Data = v
	return s
}

func (s *QueryContractAppsCompareResultResponseBodyResult) SetRequestId(v string) *QueryContractAppsCompareResultResponseBodyResult {
	s.RequestId = &v
	return s
}

type QueryContractAppsCompareResultResponseBodyResultData struct {
	CompareDetail    *QueryContractAppsCompareResultResponseBodyResultDataCompareDetail `json:"compareDetail,omitempty" xml:"compareDetail,omitempty" type:"Struct"`
	CompareDetailUrl *string                                                            `json:"compareDetailUrl,omitempty" xml:"compareDetailUrl,omitempty"`
	CompareStatus    *string                                                            `json:"compareStatus,omitempty" xml:"compareStatus,omitempty"`
}

func (s QueryContractAppsCompareResultResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsCompareResultResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryContractAppsCompareResultResponseBodyResultData) SetCompareDetail(v *QueryContractAppsCompareResultResponseBodyResultDataCompareDetail) *QueryContractAppsCompareResultResponseBodyResultData {
	s.CompareDetail = v
	return s
}

func (s *QueryContractAppsCompareResultResponseBodyResultData) SetCompareDetailUrl(v string) *QueryContractAppsCompareResultResponseBodyResultData {
	s.CompareDetailUrl = &v
	return s
}

func (s *QueryContractAppsCompareResultResponseBodyResultData) SetCompareStatus(v string) *QueryContractAppsCompareResultResponseBodyResultData {
	s.CompareStatus = &v
	return s
}

type QueryContractAppsCompareResultResponseBodyResultDataCompareDetail struct {
	Details         []*QueryContractAppsCompareResultResponseBodyResultDataCompareDetailDetails       `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
	DifferenceCount *QueryContractAppsCompareResultResponseBodyResultDataCompareDetailDifferenceCount `json:"differenceCount,omitempty" xml:"differenceCount,omitempty" type:"Struct"`
}

func (s QueryContractAppsCompareResultResponseBodyResultDataCompareDetail) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsCompareResultResponseBodyResultDataCompareDetail) GoString() string {
	return s.String()
}

func (s *QueryContractAppsCompareResultResponseBodyResultDataCompareDetail) SetDetails(v []*QueryContractAppsCompareResultResponseBodyResultDataCompareDetailDetails) *QueryContractAppsCompareResultResponseBodyResultDataCompareDetail {
	s.Details = v
	return s
}

func (s *QueryContractAppsCompareResultResponseBodyResultDataCompareDetail) SetDifferenceCount(v *QueryContractAppsCompareResultResponseBodyResultDataCompareDetailDifferenceCount) *QueryContractAppsCompareResultResponseBodyResultDataCompareDetail {
	s.DifferenceCount = v
	return s
}

type QueryContractAppsCompareResultResponseBodyResultDataCompareDetailDetails struct {
	CompareWords  *string `json:"compareWords,omitempty" xml:"compareWords,omitempty"`
	OriginalWords *string `json:"originalWords,omitempty" xml:"originalWords,omitempty"`
	Type          *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryContractAppsCompareResultResponseBodyResultDataCompareDetailDetails) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsCompareResultResponseBodyResultDataCompareDetailDetails) GoString() string {
	return s.String()
}

func (s *QueryContractAppsCompareResultResponseBodyResultDataCompareDetailDetails) SetCompareWords(v string) *QueryContractAppsCompareResultResponseBodyResultDataCompareDetailDetails {
	s.CompareWords = &v
	return s
}

func (s *QueryContractAppsCompareResultResponseBodyResultDataCompareDetailDetails) SetOriginalWords(v string) *QueryContractAppsCompareResultResponseBodyResultDataCompareDetailDetails {
	s.OriginalWords = &v
	return s
}

func (s *QueryContractAppsCompareResultResponseBodyResultDataCompareDetailDetails) SetType(v int32) *QueryContractAppsCompareResultResponseBodyResultDataCompareDetailDetails {
	s.Type = &v
	return s
}

type QueryContractAppsCompareResultResponseBodyResultDataCompareDetailDifferenceCount struct {
	Add     *int32 `json:"add,omitempty" xml:"add,omitempty"`
	Delete  *int32 `json:"delete,omitempty" xml:"delete,omitempty"`
	Replace *int32 `json:"replace,omitempty" xml:"replace,omitempty"`
	Total   *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s QueryContractAppsCompareResultResponseBodyResultDataCompareDetailDifferenceCount) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsCompareResultResponseBodyResultDataCompareDetailDifferenceCount) GoString() string {
	return s.String()
}

func (s *QueryContractAppsCompareResultResponseBodyResultDataCompareDetailDifferenceCount) SetAdd(v int32) *QueryContractAppsCompareResultResponseBodyResultDataCompareDetailDifferenceCount {
	s.Add = &v
	return s
}

func (s *QueryContractAppsCompareResultResponseBodyResultDataCompareDetailDifferenceCount) SetDelete(v int32) *QueryContractAppsCompareResultResponseBodyResultDataCompareDetailDifferenceCount {
	s.Delete = &v
	return s
}

func (s *QueryContractAppsCompareResultResponseBodyResultDataCompareDetailDifferenceCount) SetReplace(v int32) *QueryContractAppsCompareResultResponseBodyResultDataCompareDetailDifferenceCount {
	s.Replace = &v
	return s
}

func (s *QueryContractAppsCompareResultResponseBodyResultDataCompareDetailDifferenceCount) SetTotal(v int32) *QueryContractAppsCompareResultResponseBodyResultDataCompareDetailDifferenceCount {
	s.Total = &v
	return s
}

type QueryContractAppsCompareResultResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryContractAppsCompareResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryContractAppsCompareResultResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsCompareResultResponse) GoString() string {
	return s.String()
}

func (s *QueryContractAppsCompareResultResponse) SetHeaders(v map[string]*string) *QueryContractAppsCompareResultResponse {
	s.Headers = v
	return s
}

func (s *QueryContractAppsCompareResultResponse) SetStatusCode(v int32) *QueryContractAppsCompareResultResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryContractAppsCompareResultResponse) SetBody(v *QueryContractAppsCompareResultResponseBody) *QueryContractAppsCompareResultResponse {
	s.Body = v
	return s
}

type QueryContractAppsExtractResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryContractAppsExtractResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsExtractResultHeaders) GoString() string {
	return s.String()
}

func (s *QueryContractAppsExtractResultHeaders) SetCommonHeaders(v map[string]*string) *QueryContractAppsExtractResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryContractAppsExtractResultHeaders) SetXAcsDingtalkAccessToken(v string) *QueryContractAppsExtractResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryContractAppsExtractResultRequest struct {
	// This parameter is required.
	ExtractTaskId *string `json:"extractTaskId,omitempty" xml:"extractTaskId,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryContractAppsExtractResultRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsExtractResultRequest) GoString() string {
	return s.String()
}

func (s *QueryContractAppsExtractResultRequest) SetExtractTaskId(v string) *QueryContractAppsExtractResultRequest {
	s.ExtractTaskId = &v
	return s
}

func (s *QueryContractAppsExtractResultRequest) SetRequestId(v string) *QueryContractAppsExtractResultRequest {
	s.RequestId = &v
	return s
}

func (s *QueryContractAppsExtractResultRequest) SetUnionId(v string) *QueryContractAppsExtractResultRequest {
	s.UnionId = &v
	return s
}

type QueryContractAppsExtractResultResponseBody struct {
	Result  *QueryContractAppsExtractResultResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryContractAppsExtractResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsExtractResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryContractAppsExtractResultResponseBody) SetResult(v *QueryContractAppsExtractResultResponseBodyResult) *QueryContractAppsExtractResultResponseBody {
	s.Result = v
	return s
}

func (s *QueryContractAppsExtractResultResponseBody) SetSuccess(v bool) *QueryContractAppsExtractResultResponseBody {
	s.Success = &v
	return s
}

type QueryContractAppsExtractResultResponseBodyResult struct {
	Data      *QueryContractAppsExtractResultResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryContractAppsExtractResultResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsExtractResultResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryContractAppsExtractResultResponseBodyResult) SetData(v *QueryContractAppsExtractResultResponseBodyResultData) *QueryContractAppsExtractResultResponseBodyResult {
	s.Data = v
	return s
}

func (s *QueryContractAppsExtractResultResponseBodyResult) SetRequestId(v string) *QueryContractAppsExtractResultResponseBodyResult {
	s.RequestId = &v
	return s
}

type QueryContractAppsExtractResultResponseBodyResultData struct {
	ExtractEntities []*QueryContractAppsExtractResultResponseBodyResultDataExtractEntities `json:"extractEntities,omitempty" xml:"extractEntities,omitempty" type:"Repeated"`
	ExtractStatus   *string                                                                `json:"extractStatus,omitempty" xml:"extractStatus,omitempty"`
}

func (s QueryContractAppsExtractResultResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsExtractResultResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryContractAppsExtractResultResponseBodyResultData) SetExtractEntities(v []*QueryContractAppsExtractResultResponseBodyResultDataExtractEntities) *QueryContractAppsExtractResultResponseBodyResultData {
	s.ExtractEntities = v
	return s
}

func (s *QueryContractAppsExtractResultResponseBodyResultData) SetExtractStatus(v string) *QueryContractAppsExtractResultResponseBodyResultData {
	s.ExtractStatus = &v
	return s
}

type QueryContractAppsExtractResultResponseBodyResultDataExtractEntities struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s QueryContractAppsExtractResultResponseBodyResultDataExtractEntities) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsExtractResultResponseBodyResultDataExtractEntities) GoString() string {
	return s.String()
}

func (s *QueryContractAppsExtractResultResponseBodyResultDataExtractEntities) SetKey(v string) *QueryContractAppsExtractResultResponseBodyResultDataExtractEntities {
	s.Key = &v
	return s
}

func (s *QueryContractAppsExtractResultResponseBodyResultDataExtractEntities) SetValue(v string) *QueryContractAppsExtractResultResponseBodyResultDataExtractEntities {
	s.Value = &v
	return s
}

type QueryContractAppsExtractResultResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryContractAppsExtractResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryContractAppsExtractResultResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsExtractResultResponse) GoString() string {
	return s.String()
}

func (s *QueryContractAppsExtractResultResponse) SetHeaders(v map[string]*string) *QueryContractAppsExtractResultResponse {
	s.Headers = v
	return s
}

func (s *QueryContractAppsExtractResultResponse) SetStatusCode(v int32) *QueryContractAppsExtractResultResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryContractAppsExtractResultResponse) SetBody(v *QueryContractAppsExtractResultResponseBody) *QueryContractAppsExtractResultResponse {
	s.Body = v
	return s
}

type QueryContractAppsReviewResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryContractAppsReviewResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsReviewResultHeaders) GoString() string {
	return s.String()
}

func (s *QueryContractAppsReviewResultHeaders) SetCommonHeaders(v map[string]*string) *QueryContractAppsReviewResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryContractAppsReviewResultHeaders) SetXAcsDingtalkAccessToken(v string) *QueryContractAppsReviewResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryContractAppsReviewResultRequest struct {
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	ReviewTaskId *string `json:"reviewTaskId,omitempty" xml:"reviewTaskId,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryContractAppsReviewResultRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsReviewResultRequest) GoString() string {
	return s.String()
}

func (s *QueryContractAppsReviewResultRequest) SetRequestId(v string) *QueryContractAppsReviewResultRequest {
	s.RequestId = &v
	return s
}

func (s *QueryContractAppsReviewResultRequest) SetReviewTaskId(v string) *QueryContractAppsReviewResultRequest {
	s.ReviewTaskId = &v
	return s
}

func (s *QueryContractAppsReviewResultRequest) SetUnionId(v string) *QueryContractAppsReviewResultRequest {
	s.UnionId = &v
	return s
}

type QueryContractAppsReviewResultResponseBody struct {
	Result  *QueryContractAppsReviewResultResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                            `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryContractAppsReviewResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsReviewResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryContractAppsReviewResultResponseBody) SetResult(v *QueryContractAppsReviewResultResponseBodyResult) *QueryContractAppsReviewResultResponseBody {
	s.Result = v
	return s
}

func (s *QueryContractAppsReviewResultResponseBody) SetSuccess(v bool) *QueryContractAppsReviewResultResponseBody {
	s.Success = &v
	return s
}

type QueryContractAppsReviewResultResponseBodyResult struct {
	Data      *QueryContractAppsReviewResultResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryContractAppsReviewResultResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsReviewResultResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryContractAppsReviewResultResponseBodyResult) SetData(v *QueryContractAppsReviewResultResponseBodyResultData) *QueryContractAppsReviewResultResponseBodyResult {
	s.Data = v
	return s
}

func (s *QueryContractAppsReviewResultResponseBodyResult) SetRequestId(v string) *QueryContractAppsReviewResultResponseBodyResult {
	s.RequestId = &v
	return s
}

type QueryContractAppsReviewResultResponseBodyResultData struct {
	ReviewRiskDetail   []*QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetail `json:"reviewRiskDetail,omitempty" xml:"reviewRiskDetail,omitempty" type:"Repeated"`
	ReviewRiskOverview *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskOverview `json:"reviewRiskOverview,omitempty" xml:"reviewRiskOverview,omitempty" type:"Struct"`
	ReviewStatus       *QueryContractAppsReviewResultResponseBodyResultDataReviewStatus       `json:"reviewStatus,omitempty" xml:"reviewStatus,omitempty" type:"Struct"`
}

func (s QueryContractAppsReviewResultResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsReviewResultResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryContractAppsReviewResultResponseBodyResultData) SetReviewRiskDetail(v []*QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetail) *QueryContractAppsReviewResultResponseBodyResultData {
	s.ReviewRiskDetail = v
	return s
}

func (s *QueryContractAppsReviewResultResponseBodyResultData) SetReviewRiskOverview(v *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskOverview) *QueryContractAppsReviewResultResponseBodyResultData {
	s.ReviewRiskOverview = v
	return s
}

func (s *QueryContractAppsReviewResultResponseBodyResultData) SetReviewStatus(v *QueryContractAppsReviewResultResponseBodyResultDataReviewStatus) *QueryContractAppsReviewResultResponseBodyResultData {
	s.ReviewStatus = v
	return s
}

type QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetail struct {
	ExamineBrief  *string                                                                        `json:"examineBrief,omitempty" xml:"examineBrief,omitempty"`
	ExamineResult *string                                                                        `json:"examineResult,omitempty" xml:"examineResult,omitempty"`
	ExamineStatus *string                                                                        `json:"examineStatus,omitempty" xml:"examineStatus,omitempty"`
	RiskLevel     *string                                                                        `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
	RuleSequence  *string                                                                        `json:"ruleSequence,omitempty" xml:"ruleSequence,omitempty"`
	RuleTag       *string                                                                        `json:"ruleTag,omitempty" xml:"ruleTag,omitempty"`
	RuleTitle     *string                                                                        `json:"ruleTitle,omitempty" xml:"ruleTitle,omitempty"`
	SubRisks      []*QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetailSubRisks `json:"subRisks,omitempty" xml:"subRisks,omitempty" type:"Repeated"`
}

func (s QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetail) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetail) GoString() string {
	return s.String()
}

func (s *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetail) SetExamineBrief(v string) *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetail {
	s.ExamineBrief = &v
	return s
}

func (s *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetail) SetExamineResult(v string) *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetail {
	s.ExamineResult = &v
	return s
}

func (s *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetail) SetExamineStatus(v string) *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetail {
	s.ExamineStatus = &v
	return s
}

func (s *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetail) SetRiskLevel(v string) *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetail {
	s.RiskLevel = &v
	return s
}

func (s *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetail) SetRuleSequence(v string) *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetail {
	s.RuleSequence = &v
	return s
}

func (s *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetail) SetRuleTag(v string) *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetail {
	s.RuleTag = &v
	return s
}

func (s *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetail) SetRuleTitle(v string) *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetail {
	s.RuleTitle = &v
	return s
}

func (s *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetail) SetSubRisks(v []*QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetailSubRisks) *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetail {
	s.SubRisks = v
	return s
}

type QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetailSubRisks struct {
	OriginalContent *string `json:"originalContent,omitempty" xml:"originalContent,omitempty"`
	ResultContent   *string `json:"resultContent,omitempty" xml:"resultContent,omitempty"`
	ResultType      *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
	RiskBrief       *string `json:"riskBrief,omitempty" xml:"riskBrief,omitempty"`
	RiskClause      *string `json:"riskClause,omitempty" xml:"riskClause,omitempty"`
	RiskExplain     *string `json:"riskExplain,omitempty" xml:"riskExplain,omitempty"`
}

func (s QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetailSubRisks) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetailSubRisks) GoString() string {
	return s.String()
}

func (s *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetailSubRisks) SetOriginalContent(v string) *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetailSubRisks {
	s.OriginalContent = &v
	return s
}

func (s *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetailSubRisks) SetResultContent(v string) *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetailSubRisks {
	s.ResultContent = &v
	return s
}

func (s *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetailSubRisks) SetResultType(v string) *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetailSubRisks {
	s.ResultType = &v
	return s
}

func (s *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetailSubRisks) SetRiskBrief(v string) *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetailSubRisks {
	s.RiskBrief = &v
	return s
}

func (s *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetailSubRisks) SetRiskClause(v string) *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetailSubRisks {
	s.RiskClause = &v
	return s
}

func (s *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetailSubRisks) SetRiskExplain(v string) *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskDetailSubRisks {
	s.RiskExplain = &v
	return s
}

type QueryContractAppsReviewResultResponseBodyResultDataReviewRiskOverview struct {
	HasRisk    *bool  `json:"hasRisk,omitempty" xml:"hasRisk,omitempty"`
	HighRisk   *int32 `json:"highRisk,omitempty" xml:"highRisk,omitempty"`
	LowRisk    *int32 `json:"lowRisk,omitempty" xml:"lowRisk,omitempty"`
	MediumRisk *int32 `json:"mediumRisk,omitempty" xml:"mediumRisk,omitempty"`
}

func (s QueryContractAppsReviewResultResponseBodyResultDataReviewRiskOverview) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsReviewResultResponseBodyResultDataReviewRiskOverview) GoString() string {
	return s.String()
}

func (s *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskOverview) SetHasRisk(v bool) *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskOverview {
	s.HasRisk = &v
	return s
}

func (s *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskOverview) SetHighRisk(v int32) *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskOverview {
	s.HighRisk = &v
	return s
}

func (s *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskOverview) SetLowRisk(v int32) *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskOverview {
	s.LowRisk = &v
	return s
}

func (s *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskOverview) SetMediumRisk(v int32) *QueryContractAppsReviewResultResponseBodyResultDataReviewRiskOverview {
	s.MediumRisk = &v
	return s
}

type QueryContractAppsReviewResultResponseBodyResultDataReviewStatus struct {
	Overview *string `json:"overview,omitempty" xml:"overview,omitempty"`
	Result   *string `json:"result,omitempty" xml:"result,omitempty"`
	Rule     *string `json:"rule,omitempty" xml:"rule,omitempty"`
	Stage    *string `json:"stage,omitempty" xml:"stage,omitempty"`
}

func (s QueryContractAppsReviewResultResponseBodyResultDataReviewStatus) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsReviewResultResponseBodyResultDataReviewStatus) GoString() string {
	return s.String()
}

func (s *QueryContractAppsReviewResultResponseBodyResultDataReviewStatus) SetOverview(v string) *QueryContractAppsReviewResultResponseBodyResultDataReviewStatus {
	s.Overview = &v
	return s
}

func (s *QueryContractAppsReviewResultResponseBodyResultDataReviewStatus) SetResult(v string) *QueryContractAppsReviewResultResponseBodyResultDataReviewStatus {
	s.Result = &v
	return s
}

func (s *QueryContractAppsReviewResultResponseBodyResultDataReviewStatus) SetRule(v string) *QueryContractAppsReviewResultResponseBodyResultDataReviewStatus {
	s.Rule = &v
	return s
}

func (s *QueryContractAppsReviewResultResponseBodyResultDataReviewStatus) SetStage(v string) *QueryContractAppsReviewResultResponseBodyResultDataReviewStatus {
	s.Stage = &v
	return s
}

type QueryContractAppsReviewResultResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryContractAppsReviewResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryContractAppsReviewResultResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryContractAppsReviewResultResponse) GoString() string {
	return s.String()
}

func (s *QueryContractAppsReviewResultResponse) SetHeaders(v map[string]*string) *QueryContractAppsReviewResultResponse {
	s.Headers = v
	return s
}

func (s *QueryContractAppsReviewResultResponse) SetStatusCode(v int32) *QueryContractAppsReviewResultResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryContractAppsReviewResultResponse) SetBody(v *QueryContractAppsReviewResultResponseBody) *QueryContractAppsReviewResultResponse {
	s.Body = v
	return s
}

type QueryContractCompareResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryContractCompareResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryContractCompareResultHeaders) GoString() string {
	return s.String()
}

func (s *QueryContractCompareResultHeaders) SetCommonHeaders(v map[string]*string) *QueryContractCompareResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryContractCompareResultHeaders) SetXAcsDingtalkAccessToken(v string) *QueryContractCompareResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryContractCompareResultRequest struct {
	// This parameter is required.
	CompareTaskId *string `json:"compareTaskId,omitempty" xml:"compareTaskId,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryContractCompareResultRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryContractCompareResultRequest) GoString() string {
	return s.String()
}

func (s *QueryContractCompareResultRequest) SetCompareTaskId(v string) *QueryContractCompareResultRequest {
	s.CompareTaskId = &v
	return s
}

func (s *QueryContractCompareResultRequest) SetRequestId(v string) *QueryContractCompareResultRequest {
	s.RequestId = &v
	return s
}

type QueryContractCompareResultResponseBody struct {
	Result  *QueryContractCompareResultResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryContractCompareResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryContractCompareResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryContractCompareResultResponseBody) SetResult(v *QueryContractCompareResultResponseBodyResult) *QueryContractCompareResultResponseBody {
	s.Result = v
	return s
}

func (s *QueryContractCompareResultResponseBody) SetSuccess(v bool) *QueryContractCompareResultResponseBody {
	s.Success = &v
	return s
}

type QueryContractCompareResultResponseBodyResult struct {
	Data      *QueryContractCompareResultResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryContractCompareResultResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryContractCompareResultResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryContractCompareResultResponseBodyResult) SetData(v *QueryContractCompareResultResponseBodyResultData) *QueryContractCompareResultResponseBodyResult {
	s.Data = v
	return s
}

func (s *QueryContractCompareResultResponseBodyResult) SetRequestId(v string) *QueryContractCompareResultResponseBodyResult {
	s.RequestId = &v
	return s
}

type QueryContractCompareResultResponseBodyResultData struct {
	CompareDetail    *QueryContractCompareResultResponseBodyResultDataCompareDetail `json:"compareDetail,omitempty" xml:"compareDetail,omitempty" type:"Struct"`
	CompareDetailUrl *string                                                        `json:"compareDetailUrl,omitempty" xml:"compareDetailUrl,omitempty"`
	CompareStatus    *string                                                        `json:"compareStatus,omitempty" xml:"compareStatus,omitempty"`
}

func (s QueryContractCompareResultResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s QueryContractCompareResultResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryContractCompareResultResponseBodyResultData) SetCompareDetail(v *QueryContractCompareResultResponseBodyResultDataCompareDetail) *QueryContractCompareResultResponseBodyResultData {
	s.CompareDetail = v
	return s
}

func (s *QueryContractCompareResultResponseBodyResultData) SetCompareDetailUrl(v string) *QueryContractCompareResultResponseBodyResultData {
	s.CompareDetailUrl = &v
	return s
}

func (s *QueryContractCompareResultResponseBodyResultData) SetCompareStatus(v string) *QueryContractCompareResultResponseBodyResultData {
	s.CompareStatus = &v
	return s
}

type QueryContractCompareResultResponseBodyResultDataCompareDetail struct {
	Details         []*QueryContractCompareResultResponseBodyResultDataCompareDetailDetails       `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
	DifferenceCount *QueryContractCompareResultResponseBodyResultDataCompareDetailDifferenceCount `json:"differenceCount,omitempty" xml:"differenceCount,omitempty" type:"Struct"`
}

func (s QueryContractCompareResultResponseBodyResultDataCompareDetail) String() string {
	return tea.Prettify(s)
}

func (s QueryContractCompareResultResponseBodyResultDataCompareDetail) GoString() string {
	return s.String()
}

func (s *QueryContractCompareResultResponseBodyResultDataCompareDetail) SetDetails(v []*QueryContractCompareResultResponseBodyResultDataCompareDetailDetails) *QueryContractCompareResultResponseBodyResultDataCompareDetail {
	s.Details = v
	return s
}

func (s *QueryContractCompareResultResponseBodyResultDataCompareDetail) SetDifferenceCount(v *QueryContractCompareResultResponseBodyResultDataCompareDetailDifferenceCount) *QueryContractCompareResultResponseBodyResultDataCompareDetail {
	s.DifferenceCount = v
	return s
}

type QueryContractCompareResultResponseBodyResultDataCompareDetailDetails struct {
	CompareWords  *string `json:"compareWords,omitempty" xml:"compareWords,omitempty"`
	OriginalWords *string `json:"originalWords,omitempty" xml:"originalWords,omitempty"`
	Type          *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryContractCompareResultResponseBodyResultDataCompareDetailDetails) String() string {
	return tea.Prettify(s)
}

func (s QueryContractCompareResultResponseBodyResultDataCompareDetailDetails) GoString() string {
	return s.String()
}

func (s *QueryContractCompareResultResponseBodyResultDataCompareDetailDetails) SetCompareWords(v string) *QueryContractCompareResultResponseBodyResultDataCompareDetailDetails {
	s.CompareWords = &v
	return s
}

func (s *QueryContractCompareResultResponseBodyResultDataCompareDetailDetails) SetOriginalWords(v string) *QueryContractCompareResultResponseBodyResultDataCompareDetailDetails {
	s.OriginalWords = &v
	return s
}

func (s *QueryContractCompareResultResponseBodyResultDataCompareDetailDetails) SetType(v int32) *QueryContractCompareResultResponseBodyResultDataCompareDetailDetails {
	s.Type = &v
	return s
}

type QueryContractCompareResultResponseBodyResultDataCompareDetailDifferenceCount struct {
	Add     *int32 `json:"add,omitempty" xml:"add,omitempty"`
	Delete  *int32 `json:"delete,omitempty" xml:"delete,omitempty"`
	Replace *int32 `json:"replace,omitempty" xml:"replace,omitempty"`
	Total   *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s QueryContractCompareResultResponseBodyResultDataCompareDetailDifferenceCount) String() string {
	return tea.Prettify(s)
}

func (s QueryContractCompareResultResponseBodyResultDataCompareDetailDifferenceCount) GoString() string {
	return s.String()
}

func (s *QueryContractCompareResultResponseBodyResultDataCompareDetailDifferenceCount) SetAdd(v int32) *QueryContractCompareResultResponseBodyResultDataCompareDetailDifferenceCount {
	s.Add = &v
	return s
}

func (s *QueryContractCompareResultResponseBodyResultDataCompareDetailDifferenceCount) SetDelete(v int32) *QueryContractCompareResultResponseBodyResultDataCompareDetailDifferenceCount {
	s.Delete = &v
	return s
}

func (s *QueryContractCompareResultResponseBodyResultDataCompareDetailDifferenceCount) SetReplace(v int32) *QueryContractCompareResultResponseBodyResultDataCompareDetailDifferenceCount {
	s.Replace = &v
	return s
}

func (s *QueryContractCompareResultResponseBodyResultDataCompareDetailDifferenceCount) SetTotal(v int32) *QueryContractCompareResultResponseBodyResultDataCompareDetailDifferenceCount {
	s.Total = &v
	return s
}

type QueryContractCompareResultResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryContractCompareResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryContractCompareResultResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryContractCompareResultResponse) GoString() string {
	return s.String()
}

func (s *QueryContractCompareResultResponse) SetHeaders(v map[string]*string) *QueryContractCompareResultResponse {
	s.Headers = v
	return s
}

func (s *QueryContractCompareResultResponse) SetStatusCode(v int32) *QueryContractCompareResultResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryContractCompareResultResponse) SetBody(v *QueryContractCompareResultResponseBody) *QueryContractCompareResultResponse {
	s.Body = v
	return s
}

type QueryContractExtractResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryContractExtractResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryContractExtractResultHeaders) GoString() string {
	return s.String()
}

func (s *QueryContractExtractResultHeaders) SetCommonHeaders(v map[string]*string) *QueryContractExtractResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryContractExtractResultHeaders) SetXAcsDingtalkAccessToken(v string) *QueryContractExtractResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryContractExtractResultRequest struct {
	// This parameter is required.
	ExtractTaskId *string `json:"extractTaskId,omitempty" xml:"extractTaskId,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryContractExtractResultRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryContractExtractResultRequest) GoString() string {
	return s.String()
}

func (s *QueryContractExtractResultRequest) SetExtractTaskId(v string) *QueryContractExtractResultRequest {
	s.ExtractTaskId = &v
	return s
}

func (s *QueryContractExtractResultRequest) SetRequestId(v string) *QueryContractExtractResultRequest {
	s.RequestId = &v
	return s
}

type QueryContractExtractResultResponseBody struct {
	Result  *QueryContractExtractResultResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryContractExtractResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryContractExtractResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryContractExtractResultResponseBody) SetResult(v *QueryContractExtractResultResponseBodyResult) *QueryContractExtractResultResponseBody {
	s.Result = v
	return s
}

func (s *QueryContractExtractResultResponseBody) SetSuccess(v bool) *QueryContractExtractResultResponseBody {
	s.Success = &v
	return s
}

type QueryContractExtractResultResponseBodyResult struct {
	Data      *QueryContractExtractResultResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryContractExtractResultResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryContractExtractResultResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryContractExtractResultResponseBodyResult) SetData(v *QueryContractExtractResultResponseBodyResultData) *QueryContractExtractResultResponseBodyResult {
	s.Data = v
	return s
}

func (s *QueryContractExtractResultResponseBodyResult) SetRequestId(v string) *QueryContractExtractResultResponseBodyResult {
	s.RequestId = &v
	return s
}

type QueryContractExtractResultResponseBodyResultData struct {
	ExtractEntities []*QueryContractExtractResultResponseBodyResultDataExtractEntities `json:"extractEntities,omitempty" xml:"extractEntities,omitempty" type:"Repeated"`
	ExtractStatus   *string                                                            `json:"extractStatus,omitempty" xml:"extractStatus,omitempty"`
}

func (s QueryContractExtractResultResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s QueryContractExtractResultResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryContractExtractResultResponseBodyResultData) SetExtractEntities(v []*QueryContractExtractResultResponseBodyResultDataExtractEntities) *QueryContractExtractResultResponseBodyResultData {
	s.ExtractEntities = v
	return s
}

func (s *QueryContractExtractResultResponseBodyResultData) SetExtractStatus(v string) *QueryContractExtractResultResponseBodyResultData {
	s.ExtractStatus = &v
	return s
}

type QueryContractExtractResultResponseBodyResultDataExtractEntities struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s QueryContractExtractResultResponseBodyResultDataExtractEntities) String() string {
	return tea.Prettify(s)
}

func (s QueryContractExtractResultResponseBodyResultDataExtractEntities) GoString() string {
	return s.String()
}

func (s *QueryContractExtractResultResponseBodyResultDataExtractEntities) SetKey(v string) *QueryContractExtractResultResponseBodyResultDataExtractEntities {
	s.Key = &v
	return s
}

func (s *QueryContractExtractResultResponseBodyResultDataExtractEntities) SetValue(v string) *QueryContractExtractResultResponseBodyResultDataExtractEntities {
	s.Value = &v
	return s
}

type QueryContractExtractResultResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryContractExtractResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryContractExtractResultResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryContractExtractResultResponse) GoString() string {
	return s.String()
}

func (s *QueryContractExtractResultResponse) SetHeaders(v map[string]*string) *QueryContractExtractResultResponse {
	s.Headers = v
	return s
}

func (s *QueryContractExtractResultResponse) SetStatusCode(v int32) *QueryContractExtractResultResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryContractExtractResultResponse) SetBody(v *QueryContractExtractResultResponseBody) *QueryContractExtractResultResponse {
	s.Body = v
	return s
}

type QueryContractReviewResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryContractReviewResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryContractReviewResultHeaders) GoString() string {
	return s.String()
}

func (s *QueryContractReviewResultHeaders) SetCommonHeaders(v map[string]*string) *QueryContractReviewResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryContractReviewResultHeaders) SetXAcsDingtalkAccessToken(v string) *QueryContractReviewResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryContractReviewResultRequest struct {
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	ReviewTaskId *string `json:"reviewTaskId,omitempty" xml:"reviewTaskId,omitempty"`
}

func (s QueryContractReviewResultRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryContractReviewResultRequest) GoString() string {
	return s.String()
}

func (s *QueryContractReviewResultRequest) SetRequestId(v string) *QueryContractReviewResultRequest {
	s.RequestId = &v
	return s
}

func (s *QueryContractReviewResultRequest) SetReviewTaskId(v string) *QueryContractReviewResultRequest {
	s.ReviewTaskId = &v
	return s
}

type QueryContractReviewResultResponseBody struct {
	Result  *QueryContractReviewResultResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                        `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryContractReviewResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryContractReviewResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryContractReviewResultResponseBody) SetResult(v *QueryContractReviewResultResponseBodyResult) *QueryContractReviewResultResponseBody {
	s.Result = v
	return s
}

func (s *QueryContractReviewResultResponseBody) SetSuccess(v bool) *QueryContractReviewResultResponseBody {
	s.Success = &v
	return s
}

type QueryContractReviewResultResponseBodyResult struct {
	Data      *QueryContractReviewResultResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryContractReviewResultResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryContractReviewResultResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryContractReviewResultResponseBodyResult) SetData(v *QueryContractReviewResultResponseBodyResultData) *QueryContractReviewResultResponseBodyResult {
	s.Data = v
	return s
}

func (s *QueryContractReviewResultResponseBodyResult) SetRequestId(v string) *QueryContractReviewResultResponseBodyResult {
	s.RequestId = &v
	return s
}

type QueryContractReviewResultResponseBodyResultData struct {
	ReviewRiskDetail   []*QueryContractReviewResultResponseBodyResultDataReviewRiskDetail `json:"reviewRiskDetail,omitempty" xml:"reviewRiskDetail,omitempty" type:"Repeated"`
	ReviewRiskOverview *QueryContractReviewResultResponseBodyResultDataReviewRiskOverview `json:"reviewRiskOverview,omitempty" xml:"reviewRiskOverview,omitempty" type:"Struct"`
	ReviewStatus       *QueryContractReviewResultResponseBodyResultDataReviewStatus       `json:"reviewStatus,omitempty" xml:"reviewStatus,omitempty" type:"Struct"`
}

func (s QueryContractReviewResultResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s QueryContractReviewResultResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryContractReviewResultResponseBodyResultData) SetReviewRiskDetail(v []*QueryContractReviewResultResponseBodyResultDataReviewRiskDetail) *QueryContractReviewResultResponseBodyResultData {
	s.ReviewRiskDetail = v
	return s
}

func (s *QueryContractReviewResultResponseBodyResultData) SetReviewRiskOverview(v *QueryContractReviewResultResponseBodyResultDataReviewRiskOverview) *QueryContractReviewResultResponseBodyResultData {
	s.ReviewRiskOverview = v
	return s
}

func (s *QueryContractReviewResultResponseBodyResultData) SetReviewStatus(v *QueryContractReviewResultResponseBodyResultDataReviewStatus) *QueryContractReviewResultResponseBodyResultData {
	s.ReviewStatus = v
	return s
}

type QueryContractReviewResultResponseBodyResultDataReviewRiskDetail struct {
	ExamineBrief  *string                                                                    `json:"examineBrief,omitempty" xml:"examineBrief,omitempty"`
	ExamineResult *string                                                                    `json:"examineResult,omitempty" xml:"examineResult,omitempty"`
	ExamineStatus *string                                                                    `json:"examineStatus,omitempty" xml:"examineStatus,omitempty"`
	RiskLevel     *string                                                                    `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
	RuleSequence  *string                                                                    `json:"ruleSequence,omitempty" xml:"ruleSequence,omitempty"`
	RuleTag       *string                                                                    `json:"ruleTag,omitempty" xml:"ruleTag,omitempty"`
	RuleTitle     *string                                                                    `json:"ruleTitle,omitempty" xml:"ruleTitle,omitempty"`
	SubRisks      []*QueryContractReviewResultResponseBodyResultDataReviewRiskDetailSubRisks `json:"subRisks,omitempty" xml:"subRisks,omitempty" type:"Repeated"`
}

func (s QueryContractReviewResultResponseBodyResultDataReviewRiskDetail) String() string {
	return tea.Prettify(s)
}

func (s QueryContractReviewResultResponseBodyResultDataReviewRiskDetail) GoString() string {
	return s.String()
}

func (s *QueryContractReviewResultResponseBodyResultDataReviewRiskDetail) SetExamineBrief(v string) *QueryContractReviewResultResponseBodyResultDataReviewRiskDetail {
	s.ExamineBrief = &v
	return s
}

func (s *QueryContractReviewResultResponseBodyResultDataReviewRiskDetail) SetExamineResult(v string) *QueryContractReviewResultResponseBodyResultDataReviewRiskDetail {
	s.ExamineResult = &v
	return s
}

func (s *QueryContractReviewResultResponseBodyResultDataReviewRiskDetail) SetExamineStatus(v string) *QueryContractReviewResultResponseBodyResultDataReviewRiskDetail {
	s.ExamineStatus = &v
	return s
}

func (s *QueryContractReviewResultResponseBodyResultDataReviewRiskDetail) SetRiskLevel(v string) *QueryContractReviewResultResponseBodyResultDataReviewRiskDetail {
	s.RiskLevel = &v
	return s
}

func (s *QueryContractReviewResultResponseBodyResultDataReviewRiskDetail) SetRuleSequence(v string) *QueryContractReviewResultResponseBodyResultDataReviewRiskDetail {
	s.RuleSequence = &v
	return s
}

func (s *QueryContractReviewResultResponseBodyResultDataReviewRiskDetail) SetRuleTag(v string) *QueryContractReviewResultResponseBodyResultDataReviewRiskDetail {
	s.RuleTag = &v
	return s
}

func (s *QueryContractReviewResultResponseBodyResultDataReviewRiskDetail) SetRuleTitle(v string) *QueryContractReviewResultResponseBodyResultDataReviewRiskDetail {
	s.RuleTitle = &v
	return s
}

func (s *QueryContractReviewResultResponseBodyResultDataReviewRiskDetail) SetSubRisks(v []*QueryContractReviewResultResponseBodyResultDataReviewRiskDetailSubRisks) *QueryContractReviewResultResponseBodyResultDataReviewRiskDetail {
	s.SubRisks = v
	return s
}

type QueryContractReviewResultResponseBodyResultDataReviewRiskDetailSubRisks struct {
	OriginalContent *string `json:"originalContent,omitempty" xml:"originalContent,omitempty"`
	ResultContent   *string `json:"resultContent,omitempty" xml:"resultContent,omitempty"`
	ResultType      *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
	RiskBrief       *string `json:"riskBrief,omitempty" xml:"riskBrief,omitempty"`
	RiskClause      *string `json:"riskClause,omitempty" xml:"riskClause,omitempty"`
	RiskExplain     *string `json:"riskExplain,omitempty" xml:"riskExplain,omitempty"`
}

func (s QueryContractReviewResultResponseBodyResultDataReviewRiskDetailSubRisks) String() string {
	return tea.Prettify(s)
}

func (s QueryContractReviewResultResponseBodyResultDataReviewRiskDetailSubRisks) GoString() string {
	return s.String()
}

func (s *QueryContractReviewResultResponseBodyResultDataReviewRiskDetailSubRisks) SetOriginalContent(v string) *QueryContractReviewResultResponseBodyResultDataReviewRiskDetailSubRisks {
	s.OriginalContent = &v
	return s
}

func (s *QueryContractReviewResultResponseBodyResultDataReviewRiskDetailSubRisks) SetResultContent(v string) *QueryContractReviewResultResponseBodyResultDataReviewRiskDetailSubRisks {
	s.ResultContent = &v
	return s
}

func (s *QueryContractReviewResultResponseBodyResultDataReviewRiskDetailSubRisks) SetResultType(v string) *QueryContractReviewResultResponseBodyResultDataReviewRiskDetailSubRisks {
	s.ResultType = &v
	return s
}

func (s *QueryContractReviewResultResponseBodyResultDataReviewRiskDetailSubRisks) SetRiskBrief(v string) *QueryContractReviewResultResponseBodyResultDataReviewRiskDetailSubRisks {
	s.RiskBrief = &v
	return s
}

func (s *QueryContractReviewResultResponseBodyResultDataReviewRiskDetailSubRisks) SetRiskClause(v string) *QueryContractReviewResultResponseBodyResultDataReviewRiskDetailSubRisks {
	s.RiskClause = &v
	return s
}

func (s *QueryContractReviewResultResponseBodyResultDataReviewRiskDetailSubRisks) SetRiskExplain(v string) *QueryContractReviewResultResponseBodyResultDataReviewRiskDetailSubRisks {
	s.RiskExplain = &v
	return s
}

type QueryContractReviewResultResponseBodyResultDataReviewRiskOverview struct {
	HasRisk    *bool  `json:"hasRisk,omitempty" xml:"hasRisk,omitempty"`
	HighRisk   *int32 `json:"highRisk,omitempty" xml:"highRisk,omitempty"`
	LowRisk    *int32 `json:"lowRisk,omitempty" xml:"lowRisk,omitempty"`
	MediumRisk *int32 `json:"mediumRisk,omitempty" xml:"mediumRisk,omitempty"`
}

func (s QueryContractReviewResultResponseBodyResultDataReviewRiskOverview) String() string {
	return tea.Prettify(s)
}

func (s QueryContractReviewResultResponseBodyResultDataReviewRiskOverview) GoString() string {
	return s.String()
}

func (s *QueryContractReviewResultResponseBodyResultDataReviewRiskOverview) SetHasRisk(v bool) *QueryContractReviewResultResponseBodyResultDataReviewRiskOverview {
	s.HasRisk = &v
	return s
}

func (s *QueryContractReviewResultResponseBodyResultDataReviewRiskOverview) SetHighRisk(v int32) *QueryContractReviewResultResponseBodyResultDataReviewRiskOverview {
	s.HighRisk = &v
	return s
}

func (s *QueryContractReviewResultResponseBodyResultDataReviewRiskOverview) SetLowRisk(v int32) *QueryContractReviewResultResponseBodyResultDataReviewRiskOverview {
	s.LowRisk = &v
	return s
}

func (s *QueryContractReviewResultResponseBodyResultDataReviewRiskOverview) SetMediumRisk(v int32) *QueryContractReviewResultResponseBodyResultDataReviewRiskOverview {
	s.MediumRisk = &v
	return s
}

type QueryContractReviewResultResponseBodyResultDataReviewStatus struct {
	Overview *string `json:"overview,omitempty" xml:"overview,omitempty"`
	Result   *string `json:"result,omitempty" xml:"result,omitempty"`
	Rule     *string `json:"rule,omitempty" xml:"rule,omitempty"`
	Stage    *string `json:"stage,omitempty" xml:"stage,omitempty"`
}

func (s QueryContractReviewResultResponseBodyResultDataReviewStatus) String() string {
	return tea.Prettify(s)
}

func (s QueryContractReviewResultResponseBodyResultDataReviewStatus) GoString() string {
	return s.String()
}

func (s *QueryContractReviewResultResponseBodyResultDataReviewStatus) SetOverview(v string) *QueryContractReviewResultResponseBodyResultDataReviewStatus {
	s.Overview = &v
	return s
}

func (s *QueryContractReviewResultResponseBodyResultDataReviewStatus) SetResult(v string) *QueryContractReviewResultResponseBodyResultDataReviewStatus {
	s.Result = &v
	return s
}

func (s *QueryContractReviewResultResponseBodyResultDataReviewStatus) SetRule(v string) *QueryContractReviewResultResponseBodyResultDataReviewStatus {
	s.Rule = &v
	return s
}

func (s *QueryContractReviewResultResponseBodyResultDataReviewStatus) SetStage(v string) *QueryContractReviewResultResponseBodyResultDataReviewStatus {
	s.Stage = &v
	return s
}

type QueryContractReviewResultResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryContractReviewResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryContractReviewResultResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryContractReviewResultResponse) GoString() string {
	return s.String()
}

func (s *QueryContractReviewResultResponse) SetHeaders(v map[string]*string) *QueryContractReviewResultResponse {
	s.Headers = v
	return s
}

func (s *QueryContractReviewResultResponse) SetStatusCode(v int32) *QueryContractReviewResultResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryContractReviewResultResponse) SetBody(v *QueryContractReviewResultResponseBody) *QueryContractReviewResultResponse {
	s.Body = v
	return s
}

type SendContractCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendContractCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendContractCardHeaders) GoString() string {
	return s.String()
}

func (s *SendContractCardHeaders) SetCommonHeaders(v map[string]*string) *SendContractCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendContractCardHeaders) SetXAcsDingtalkAccessToken(v string) *SendContractCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendContractCardRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// signing
	CardType     *string                              `json:"cardType,omitempty" xml:"cardType,omitempty"`
	ContractInfo *SendContractCardRequestContractInfo `json:"contractInfo,omitempty" xml:"contractInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// ding5f62ac8a3c24952ebc961a6cb783455b
	CorpId    *string            `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Extension map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
	// example:
	//
	// PROC_Xxxxxxxx
	ProcessInstanceId *string   `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	ReceiveGroups     []*string `json:"receiveGroups,omitempty" xml:"receiveGroups,omitempty" type:"Repeated"`
	// This parameter is required.
	Receivers []*SendContractCardRequestReceivers `json:"receivers,omitempty" xml:"receivers,omitempty" type:"Repeated"`
	// This parameter is required.
	Sender *SendContractCardRequestSender `json:"sender,omitempty" xml:"sender,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// true
	SyncSingleChat *bool `json:"syncSingleChat,omitempty" xml:"syncSingleChat,omitempty"`
}

func (s SendContractCardRequest) String() string {
	return tea.Prettify(s)
}

func (s SendContractCardRequest) GoString() string {
	return s.String()
}

func (s *SendContractCardRequest) SetCardType(v string) *SendContractCardRequest {
	s.CardType = &v
	return s
}

func (s *SendContractCardRequest) SetContractInfo(v *SendContractCardRequestContractInfo) *SendContractCardRequest {
	s.ContractInfo = v
	return s
}

func (s *SendContractCardRequest) SetCorpId(v string) *SendContractCardRequest {
	s.CorpId = &v
	return s
}

func (s *SendContractCardRequest) SetExtension(v map[string]*string) *SendContractCardRequest {
	s.Extension = v
	return s
}

func (s *SendContractCardRequest) SetProcessInstanceId(v string) *SendContractCardRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *SendContractCardRequest) SetReceiveGroups(v []*string) *SendContractCardRequest {
	s.ReceiveGroups = v
	return s
}

func (s *SendContractCardRequest) SetReceivers(v []*SendContractCardRequestReceivers) *SendContractCardRequest {
	s.Receivers = v
	return s
}

func (s *SendContractCardRequest) SetSender(v *SendContractCardRequestSender) *SendContractCardRequest {
	s.Sender = v
	return s
}

func (s *SendContractCardRequest) SetSyncSingleChat(v bool) *SendContractCardRequest {
	s.SyncSingleChat = &v
	return s
}

type SendContractCardRequestContractInfo struct {
	// example:
	//
	// HT_xxxxxxx
	ContractCode *string `json:"contractCode,omitempty" xml:"contractCode,omitempty"`
	// example:
	//
	// 合同
	ContractName *string `json:"contractName,omitempty" xml:"contractName,omitempty"`
	// example:
	//
	// 1242153453
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 张三
	SignUserName *string `json:"signUserName,omitempty" xml:"signUserName,omitempty"`
}

func (s SendContractCardRequestContractInfo) String() string {
	return tea.Prettify(s)
}

func (s SendContractCardRequestContractInfo) GoString() string {
	return s.String()
}

func (s *SendContractCardRequestContractInfo) SetContractCode(v string) *SendContractCardRequestContractInfo {
	s.ContractCode = &v
	return s
}

func (s *SendContractCardRequestContractInfo) SetContractName(v string) *SendContractCardRequestContractInfo {
	s.ContractName = &v
	return s
}

func (s *SendContractCardRequestContractInfo) SetCreateTime(v int64) *SendContractCardRequestContractInfo {
	s.CreateTime = &v
	return s
}

func (s *SendContractCardRequestContractInfo) SetSignUserName(v string) *SendContractCardRequestContractInfo {
	s.SignUserName = &v
	return s
}

type SendContractCardRequestReceivers struct {
	// example:
	//
	// ding5f62ac8a3c24952ebc961a6cb783455b
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// 1622265907855672
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// 可以为空
	UserType *string `json:"userType,omitempty" xml:"userType,omitempty"`
}

func (s SendContractCardRequestReceivers) String() string {
	return tea.Prettify(s)
}

func (s SendContractCardRequestReceivers) GoString() string {
	return s.String()
}

func (s *SendContractCardRequestReceivers) SetCorpId(v string) *SendContractCardRequestReceivers {
	s.CorpId = &v
	return s
}

func (s *SendContractCardRequestReceivers) SetUserId(v string) *SendContractCardRequestReceivers {
	s.UserId = &v
	return s
}

func (s *SendContractCardRequestReceivers) SetUserType(v string) *SendContractCardRequestReceivers {
	s.UserType = &v
	return s
}

type SendContractCardRequestSender struct {
	// example:
	//
	// ding5f62ac8a3c24952ebc961a6cb783455b
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// 1622265907855672
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// 可以为空
	UserType *string `json:"userType,omitempty" xml:"userType,omitempty"`
}

func (s SendContractCardRequestSender) String() string {
	return tea.Prettify(s)
}

func (s SendContractCardRequestSender) GoString() string {
	return s.String()
}

func (s *SendContractCardRequestSender) SetCorpId(v string) *SendContractCardRequestSender {
	s.CorpId = &v
	return s
}

func (s *SendContractCardRequestSender) SetUserId(v string) *SendContractCardRequestSender {
	s.UserId = &v
	return s
}

func (s *SendContractCardRequestSender) SetUserType(v string) *SendContractCardRequestSender {
	s.UserType = &v
	return s
}

type SendContractCardResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SendContractCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendContractCardResponseBody) GoString() string {
	return s.String()
}

func (s *SendContractCardResponseBody) SetSuccess(v bool) *SendContractCardResponseBody {
	s.Success = &v
	return s
}

type SendContractCardResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendContractCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendContractCardResponse) String() string {
	return tea.Prettify(s)
}

func (s SendContractCardResponse) GoString() string {
	return s.String()
}

func (s *SendContractCardResponse) SetHeaders(v map[string]*string) *SendContractCardResponse {
	s.Headers = v
	return s
}

func (s *SendContractCardResponse) SetStatusCode(v int32) *SendContractCardResponse {
	s.StatusCode = &v
	return s
}

func (s *SendContractCardResponse) SetBody(v *SendContractCardResponseBody) *SendContractCardResponse {
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
// 取消审查工单接口
//
// @param request - CancelReviewOrderRequest
//
// @param headers - CancelReviewOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelReviewOrderResponse
func (client *Client) CancelReviewOrderWithOptions(request *CancelReviewOrderRequest, headers *CancelReviewOrderHeaders, runtime *util.RuntimeOptions) (_result *CancelReviewOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndFiles)) {
		body["endFiles"] = request.EndFiles
	}

	if !tea.BoolValue(util.IsUnset(request.Extension)) {
		body["extension"] = request.Extension
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		body["orderId"] = request.OrderId
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
		Action:      tea.String("CancelReviewOrder"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/reviews/cancel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelReviewOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消审查工单接口
//
// @param request - CancelReviewOrderRequest
//
// @return CancelReviewOrderResponse
func (client *Client) CancelReviewOrder(request *CancelReviewOrderRequest) (_result *CancelReviewOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CancelReviewOrderHeaders{}
	_result = &CancelReviewOrderResponse{}
	_body, _err := client.CancelReviewOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询esign文件是否正常
//
// @param request - CheckEsignFileRequest
//
// @param headers - CheckEsignFileHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckEsignFileResponse
func (client *Client) CheckEsignFileWithOptions(request *CheckEsignFileRequest, headers *CheckEsignFileHeaders, runtime *util.RuntimeOptions) (_result *CheckEsignFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["fileId"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["spaceId"] = request.SpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
		Action:      tea.String("CheckEsignFile"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/esignFiles/check"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckEsignFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询esign文件是否正常
//
// @param request - CheckEsignFileRequest
//
// @return CheckEsignFileResponse
func (client *Client) CheckEsignFile(request *CheckEsignFileRequest) (_result *CheckEsignFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CheckEsignFileHeaders{}
	_result = &CheckEsignFileResponse{}
	_body, _err := client.CheckEsignFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 合同权益核销
//
// @param request - ContractBenefitConsumeRequest
//
// @param headers - ContractBenefitConsumeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContractBenefitConsumeResponse
func (client *Client) ContractBenefitConsumeWithOptions(request *ContractBenefitConsumeRequest, headers *ContractBenefitConsumeHeaders, runtime *util.RuntimeOptions) (_result *ContractBenefitConsumeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BenefitPoint)) {
		body["benefitPoint"] = request.BenefitPoint
	}

	if !tea.BoolValue(util.IsUnset(request.BizRequestId)) {
		body["bizRequestId"] = request.BizRequestId
	}

	if !tea.BoolValue(util.IsUnset(request.ConsumeQuota)) {
		body["consumeQuota"] = request.ConsumeQuota
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtParams)) {
		body["extParams"] = request.ExtParams
	}

	if !tea.BoolValue(util.IsUnset(request.IsvCorpId)) {
		body["isvCorpId"] = request.IsvCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.OptUnionId)) {
		body["optUnionId"] = request.OptUnionId
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
		Action:      tea.String("ContractBenefitConsume"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/benefits/consume"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ContractBenefitConsumeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合同权益核销
//
// @param request - ContractBenefitConsumeRequest
//
// @return ContractBenefitConsumeResponse
func (client *Client) ContractBenefitConsume(request *ContractBenefitConsumeRequest) (_result *ContractBenefitConsumeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ContractBenefitConsumeHeaders{}
	_result = &ContractBenefitConsumeResponse{}
	_body, _err := client.ContractBenefitConsumeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建合同比对任务
//
// @param request - CreateContractAppsCompareTaskRequest
//
// @param headers - CreateContractAppsCompareTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateContractAppsCompareTaskResponse
func (client *Client) CreateContractAppsCompareTaskWithOptions(request *CreateContractAppsCompareTaskRequest, headers *CreateContractAppsCompareTaskHeaders, runtime *util.RuntimeOptions) (_result *CreateContractAppsCompareTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ComparativeFile)) {
		body["comparativeFile"] = request.ComparativeFile
	}

	if !tea.BoolValue(util.IsUnset(request.ComparativeFileDownloadUrl)) {
		body["comparativeFileDownloadUrl"] = request.ComparativeFileDownloadUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ComparativeFileName)) {
		body["comparativeFileName"] = request.ComparativeFileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileSource)) {
		body["fileSource"] = request.FileSource
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.StandardFile)) {
		body["standardFile"] = request.StandardFile
	}

	if !tea.BoolValue(util.IsUnset(request.StandardFileDownloadUrl)) {
		body["standardFileDownloadUrl"] = request.StandardFileDownloadUrl
	}

	if !tea.BoolValue(util.IsUnset(request.StandardFileName)) {
		body["standardFileName"] = request.StandardFileName
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("CreateContractAppsCompareTask"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/apps/comparisonTasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateContractAppsCompareTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建合同比对任务
//
// @param request - CreateContractAppsCompareTaskRequest
//
// @return CreateContractAppsCompareTaskResponse
func (client *Client) CreateContractAppsCompareTask(request *CreateContractAppsCompareTaskRequest) (_result *CreateContractAppsCompareTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateContractAppsCompareTaskHeaders{}
	_result = &CreateContractAppsCompareTaskResponse{}
	_body, _err := client.CreateContractAppsCompareTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建合同提取任务
//
// @param request - CreateContractAppsExtractTaskRequest
//
// @param headers - CreateContractAppsExtractTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateContractAppsExtractTaskResponse
func (client *Client) CreateContractAppsExtractTaskWithOptions(request *CreateContractAppsExtractTaskRequest, headers *CreateContractAppsExtractTaskHeaders, runtime *util.RuntimeOptions) (_result *CreateContractAppsExtractTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContractFile)) {
		body["contractFile"] = request.ContractFile
	}

	if !tea.BoolValue(util.IsUnset(request.ContractFileDownloadUrl)) {
		body["contractFileDownloadUrl"] = request.ContractFileDownloadUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ContractFileName)) {
		body["contractFileName"] = request.ContractFileName
	}

	if !tea.BoolValue(util.IsUnset(request.ExtractKeys)) {
		body["extractKeys"] = request.ExtractKeys
	}

	if !tea.BoolValue(util.IsUnset(request.FileSource)) {
		body["fileSource"] = request.FileSource
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("CreateContractAppsExtractTask"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/apps/extractTasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateContractAppsExtractTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建合同提取任务
//
// @param request - CreateContractAppsExtractTaskRequest
//
// @return CreateContractAppsExtractTaskResponse
func (client *Client) CreateContractAppsExtractTask(request *CreateContractAppsExtractTaskRequest) (_result *CreateContractAppsExtractTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateContractAppsExtractTaskHeaders{}
	_result = &CreateContractAppsExtractTaskResponse{}
	_body, _err := client.CreateContractAppsExtractTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建合同审查任务
//
// @param request - CreateContractAppsReviewTaskRequest
//
// @param headers - CreateContractAppsReviewTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateContractAppsReviewTaskResponse
func (client *Client) CreateContractAppsReviewTaskWithOptions(request *CreateContractAppsReviewTaskRequest, headers *CreateContractAppsReviewTaskHeaders, runtime *util.RuntimeOptions) (_result *CreateContractAppsReviewTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContractFile)) {
		body["contractFile"] = request.ContractFile
	}

	if !tea.BoolValue(util.IsUnset(request.ContractFileDownloadUrl)) {
		body["contractFileDownloadUrl"] = request.ContractFileDownloadUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ContractFileName)) {
		body["contractFileName"] = request.ContractFileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileSource)) {
		body["fileSource"] = request.FileSource
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.ReviewCustomRules)) {
		body["reviewCustomRules"] = request.ReviewCustomRules
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		body["ruleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.Standpoint)) {
		body["standpoint"] = request.Standpoint
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("CreateContractAppsReviewTask"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/apps/reviewTasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateContractAppsReviewTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建合同审查任务
//
// @param request - CreateContractAppsReviewTaskRequest
//
// @return CreateContractAppsReviewTaskResponse
func (client *Client) CreateContractAppsReviewTask(request *CreateContractAppsReviewTaskRequest) (_result *CreateContractAppsReviewTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateContractAppsReviewTaskHeaders{}
	_result = &CreateContractAppsReviewTaskResponse{}
	_body, _err := client.CreateContractAppsReviewTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建合同比对任务
//
// @param request - CreateContractCompareTaskRequest
//
// @param headers - CreateContractCompareTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateContractCompareTaskResponse
func (client *Client) CreateContractCompareTaskWithOptions(request *CreateContractCompareTaskRequest, headers *CreateContractCompareTaskHeaders, runtime *util.RuntimeOptions) (_result *CreateContractCompareTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ComparativeFile)) {
		body["comparativeFile"] = request.ComparativeFile
	}

	if !tea.BoolValue(util.IsUnset(request.ComparativeFileDownloadUrl)) {
		body["comparativeFileDownloadUrl"] = request.ComparativeFileDownloadUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ComparativeFileName)) {
		body["comparativeFileName"] = request.ComparativeFileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileSource)) {
		body["fileSource"] = request.FileSource
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.StandardFile)) {
		body["standardFile"] = request.StandardFile
	}

	if !tea.BoolValue(util.IsUnset(request.StandardFileDownloadUrl)) {
		body["standardFileDownloadUrl"] = request.StandardFileDownloadUrl
	}

	if !tea.BoolValue(util.IsUnset(request.StandardFileName)) {
		body["standardFileName"] = request.StandardFileName
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
		Action:      tea.String("CreateContractCompareTask"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/comparisonTasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateContractCompareTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建合同比对任务
//
// @param request - CreateContractCompareTaskRequest
//
// @return CreateContractCompareTaskResponse
func (client *Client) CreateContractCompareTask(request *CreateContractCompareTaskRequest) (_result *CreateContractCompareTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateContractCompareTaskHeaders{}
	_result = &CreateContractCompareTaskResponse{}
	_body, _err := client.CreateContractCompareTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建合同提取任务
//
// @param request - CreateContractExtractTaskRequest
//
// @param headers - CreateContractExtractTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateContractExtractTaskResponse
func (client *Client) CreateContractExtractTaskWithOptions(request *CreateContractExtractTaskRequest, headers *CreateContractExtractTaskHeaders, runtime *util.RuntimeOptions) (_result *CreateContractExtractTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContractFile)) {
		body["contractFile"] = request.ContractFile
	}

	if !tea.BoolValue(util.IsUnset(request.ContractFileDownloadUrl)) {
		body["contractFileDownloadUrl"] = request.ContractFileDownloadUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ContractFileName)) {
		body["contractFileName"] = request.ContractFileName
	}

	if !tea.BoolValue(util.IsUnset(request.ExtractKeys)) {
		body["extractKeys"] = request.ExtractKeys
	}

	if !tea.BoolValue(util.IsUnset(request.FileSource)) {
		body["fileSource"] = request.FileSource
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
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
		Action:      tea.String("CreateContractExtractTask"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/extractTasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateContractExtractTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建合同提取任务
//
// @param request - CreateContractExtractTaskRequest
//
// @return CreateContractExtractTaskResponse
func (client *Client) CreateContractExtractTask(request *CreateContractExtractTaskRequest) (_result *CreateContractExtractTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateContractExtractTaskHeaders{}
	_result = &CreateContractExtractTaskResponse{}
	_body, _err := client.CreateContractExtractTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建合同审查任务
//
// @param request - CreateContractReviewTaskRequest
//
// @param headers - CreateContractReviewTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateContractReviewTaskResponse
func (client *Client) CreateContractReviewTaskWithOptions(request *CreateContractReviewTaskRequest, headers *CreateContractReviewTaskHeaders, runtime *util.RuntimeOptions) (_result *CreateContractReviewTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContractFile)) {
		body["contractFile"] = request.ContractFile
	}

	if !tea.BoolValue(util.IsUnset(request.ContractFileDownloadUrl)) {
		body["contractFileDownloadUrl"] = request.ContractFileDownloadUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ContractFileName)) {
		body["contractFileName"] = request.ContractFileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileSource)) {
		body["fileSource"] = request.FileSource
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.ReviewCustomRules)) {
		body["reviewCustomRules"] = request.ReviewCustomRules
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		body["ruleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.Standpoint)) {
		body["standpoint"] = request.Standpoint
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
		Action:      tea.String("CreateContractReviewTask"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/reviewTasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateContractReviewTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建合同审查任务
//
// @param request - CreateContractReviewTaskRequest
//
// @return CreateContractReviewTaskResponse
func (client *Client) CreateContractReviewTask(request *CreateContractReviewTaskRequest) (_result *CreateContractReviewTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateContractReviewTaskHeaders{}
	_result = &CreateContractReviewTaskResponse{}
	_body, _err := client.CreateContractReviewTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 天谷侧查询审批单
//
// @param request - EsignQueryApprovalInfoRequest
//
// @param headers - EsignQueryApprovalInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EsignQueryApprovalInfoResponse
func (client *Client) EsignQueryApprovalInfoWithOptions(request *EsignQueryApprovalInfoRequest, headers *EsignQueryApprovalInfoHeaders, runtime *util.RuntimeOptions) (_result *EsignQueryApprovalInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.EsignFlowId)) {
		body["esignFlowId"] = request.EsignFlowId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("EsignQueryApprovalInfo"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/esign/approvalInfos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &EsignQueryApprovalInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 天谷侧查询审批单
//
// @param request - EsignQueryApprovalInfoRequest
//
// @return EsignQueryApprovalInfoResponse
func (client *Client) EsignQueryApprovalInfo(request *EsignQueryApprovalInfoRequest) (_result *EsignQueryApprovalInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EsignQueryApprovalInfoHeaders{}
	_result = &EsignQueryApprovalInfoResponse{}
	_body, _err := client.EsignQueryApprovalInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 天谷侧查询授权信息接口
//
// @param request - EsignQueryGrantInfoRequest
//
// @param headers - EsignQueryGrantInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EsignQueryGrantInfoResponse
func (client *Client) EsignQueryGrantInfoWithOptions(request *EsignQueryGrantInfoRequest, headers *EsignQueryGrantInfoHeaders, runtime *util.RuntimeOptions) (_result *EsignQueryGrantInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Extension)) {
		body["extension"] = request.Extension
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("EsignQueryGrantInfo"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/esign/anthInfos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &EsignQueryGrantInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 天谷侧查询授权信息接口
//
// @param request - EsignQueryGrantInfoRequest
//
// @return EsignQueryGrantInfoResponse
func (client *Client) EsignQueryGrantInfo(request *EsignQueryGrantInfoRequest) (_result *EsignQueryGrantInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EsignQueryGrantInfoHeaders{}
	_result = &EsignQueryGrantInfoResponse{}
	_body, _err := client.EsignQueryGrantInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 天谷侧查询获取免登信息
//
// @param request - EsignQueryIdentityByTicketRequest
//
// @param headers - EsignQueryIdentityByTicketHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EsignQueryIdentityByTicketResponse
func (client *Client) EsignQueryIdentityByTicketWithOptions(request *EsignQueryIdentityByTicketRequest, headers *EsignQueryIdentityByTicketHeaders, runtime *util.RuntimeOptions) (_result *EsignQueryIdentityByTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Extension)) {
		body["extension"] = request.Extension
	}

	if !tea.BoolValue(util.IsUnset(request.Ticket)) {
		body["ticket"] = request.Ticket
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
		Action:      tea.String("EsignQueryIdentityByTicket"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/esign/tickets/identities/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &EsignQueryIdentityByTicketResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 天谷侧查询获取免登信息
//
// @param request - EsignQueryIdentityByTicketRequest
//
// @return EsignQueryIdentityByTicketResponse
func (client *Client) EsignQueryIdentityByTicket(request *EsignQueryIdentityByTicketRequest) (_result *EsignQueryIdentityByTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EsignQueryIdentityByTicketHeaders{}
	_result = &EsignQueryIdentityByTicketResponse{}
	_body, _err := client.EsignQueryIdentityByTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// e签宝电子签事件同步回传接口
//
// @param request - EsignSyncEventRequest
//
// @param headers - EsignSyncEventHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EsignSyncEventResponse
func (client *Client) EsignSyncEventWithOptions(request *EsignSyncEventRequest, headers *EsignSyncEventHeaders, runtime *util.RuntimeOptions) (_result *EsignSyncEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		body["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.EsignData)) {
		body["esignData"] = request.EsignData
	}

	if !tea.BoolValue(util.IsUnset(request.Extension)) {
		body["extension"] = request.Extension
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("EsignSyncEvent"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/esign/events/sync"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &EsignSyncEventResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// e签宝电子签事件同步回传接口
//
// @param request - EsignSyncEventRequest
//
// @return EsignSyncEventResponse
func (client *Client) EsignSyncEvent(request *EsignSyncEventRequest) (_result *EsignSyncEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EsignSyncEventHeaders{}
	_result = &EsignSyncEventResponse{}
	_body, _err := client.EsignSyncEventWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 校验钉钉用户能否访问e签宝页面接口
//
// @param request - EsignUserVerifyRequest
//
// @param headers - EsignUserVerifyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EsignUserVerifyResponse
func (client *Client) EsignUserVerifyWithOptions(request *EsignUserVerifyRequest, headers *EsignUserVerifyHeaders, runtime *util.RuntimeOptions) (_result *EsignUserVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("EsignUserVerify"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/esign/user/verify"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &EsignUserVerifyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验钉钉用户能否访问e签宝页面接口
//
// @param request - EsignUserVerifyRequest
//
// @return EsignUserVerifyResponse
func (client *Client) EsignUserVerify(request *EsignUserVerifyRequest) (_result *EsignUserVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EsignUserVerifyHeaders{}
	_result = &EsignUserVerifyResponse{}
	_body, _err := client.EsignUserVerifyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 完成工单审查接口
//
// @param request - FinishReviewOrderRequest
//
// @param headers - FinishReviewOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FinishReviewOrderResponse
func (client *Client) FinishReviewOrderWithOptions(request *FinishReviewOrderRequest, headers *FinishReviewOrderHeaders, runtime *util.RuntimeOptions) (_result *FinishReviewOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndFiles)) {
		body["endFiles"] = request.EndFiles
	}

	if !tea.BoolValue(util.IsUnset(request.Extension)) {
		body["extension"] = request.Extension
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		body["orderId"] = request.OrderId
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
		Action:      tea.String("FinishReviewOrder"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/reviews/finish"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &FinishReviewOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 完成工单审查接口
//
// @param request - FinishReviewOrderRequest
//
// @return FinishReviewOrderResponse
func (client *Client) FinishReviewOrder(request *FinishReviewOrderRequest) (_result *FinishReviewOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &FinishReviewOrderHeaders{}
	_result = &FinishReviewOrderResponse{}
	_body, _err := client.FinishReviewOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// e签宝查询智能合同版本接口
//
// @param request - QueryAdvancedContractVersionRequest
//
// @param headers - QueryAdvancedContractVersionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAdvancedContractVersionResponse
func (client *Client) QueryAdvancedContractVersionWithOptions(request *QueryAdvancedContractVersionRequest, headers *QueryAdvancedContractVersionHeaders, runtime *util.RuntimeOptions) (_result *QueryAdvancedContractVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Extension)) {
		body["extension"] = request.Extension
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
		Action:      tea.String("QueryAdvancedContractVersion"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/versions/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAdvancedContractVersionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// e签宝查询智能合同版本接口
//
// @param request - QueryAdvancedContractVersionRequest
//
// @return QueryAdvancedContractVersionResponse
func (client *Client) QueryAdvancedContractVersion(request *QueryAdvancedContractVersionRequest) (_result *QueryAdvancedContractVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAdvancedContractVersionHeaders{}
	_result = &QueryAdvancedContractVersionResponse{}
	_body, _err := client.QueryAdvancedContractVersionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询合同比对结果
//
// @param request - QueryContractAppsCompareResultRequest
//
// @param headers - QueryContractAppsCompareResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryContractAppsCompareResultResponse
func (client *Client) QueryContractAppsCompareResultWithOptions(request *QueryContractAppsCompareResultRequest, headers *QueryContractAppsCompareResultHeaders, runtime *util.RuntimeOptions) (_result *QueryContractAppsCompareResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompareTaskId)) {
		body["compareTaskId"] = request.CompareTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("QueryContractAppsCompareResult"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/apps/comparisonResults/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryContractAppsCompareResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询合同比对结果
//
// @param request - QueryContractAppsCompareResultRequest
//
// @return QueryContractAppsCompareResultResponse
func (client *Client) QueryContractAppsCompareResult(request *QueryContractAppsCompareResultRequest) (_result *QueryContractAppsCompareResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryContractAppsCompareResultHeaders{}
	_result = &QueryContractAppsCompareResultResponse{}
	_body, _err := client.QueryContractAppsCompareResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询合同提取结果
//
// @param request - QueryContractAppsExtractResultRequest
//
// @param headers - QueryContractAppsExtractResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryContractAppsExtractResultResponse
func (client *Client) QueryContractAppsExtractResultWithOptions(request *QueryContractAppsExtractResultRequest, headers *QueryContractAppsExtractResultHeaders, runtime *util.RuntimeOptions) (_result *QueryContractAppsExtractResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtractTaskId)) {
		body["extractTaskId"] = request.ExtractTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("QueryContractAppsExtractResult"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/apps/extractResults/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryContractAppsExtractResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询合同提取结果
//
// @param request - QueryContractAppsExtractResultRequest
//
// @return QueryContractAppsExtractResultResponse
func (client *Client) QueryContractAppsExtractResult(request *QueryContractAppsExtractResultRequest) (_result *QueryContractAppsExtractResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryContractAppsExtractResultHeaders{}
	_result = &QueryContractAppsExtractResultResponse{}
	_body, _err := client.QueryContractAppsExtractResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询合同审查结果
//
// @param request - QueryContractAppsReviewResultRequest
//
// @param headers - QueryContractAppsReviewResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryContractAppsReviewResultResponse
func (client *Client) QueryContractAppsReviewResultWithOptions(request *QueryContractAppsReviewResultRequest, headers *QueryContractAppsReviewResultHeaders, runtime *util.RuntimeOptions) (_result *QueryContractAppsReviewResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.ReviewTaskId)) {
		body["reviewTaskId"] = request.ReviewTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("QueryContractAppsReviewResult"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/apps/reviewResults/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryContractAppsReviewResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询合同审查结果
//
// @param request - QueryContractAppsReviewResultRequest
//
// @return QueryContractAppsReviewResultResponse
func (client *Client) QueryContractAppsReviewResult(request *QueryContractAppsReviewResultRequest) (_result *QueryContractAppsReviewResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryContractAppsReviewResultHeaders{}
	_result = &QueryContractAppsReviewResultResponse{}
	_body, _err := client.QueryContractAppsReviewResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询合同比对结果
//
// @param request - QueryContractCompareResultRequest
//
// @param headers - QueryContractCompareResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryContractCompareResultResponse
func (client *Client) QueryContractCompareResultWithOptions(request *QueryContractCompareResultRequest, headers *QueryContractCompareResultHeaders, runtime *util.RuntimeOptions) (_result *QueryContractCompareResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompareTaskId)) {
		body["compareTaskId"] = request.CompareTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
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
		Action:      tea.String("QueryContractCompareResult"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/comparisonResults/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryContractCompareResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询合同比对结果
//
// @param request - QueryContractCompareResultRequest
//
// @return QueryContractCompareResultResponse
func (client *Client) QueryContractCompareResult(request *QueryContractCompareResultRequest) (_result *QueryContractCompareResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryContractCompareResultHeaders{}
	_result = &QueryContractCompareResultResponse{}
	_body, _err := client.QueryContractCompareResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询合同提取结果
//
// @param request - QueryContractExtractResultRequest
//
// @param headers - QueryContractExtractResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryContractExtractResultResponse
func (client *Client) QueryContractExtractResultWithOptions(request *QueryContractExtractResultRequest, headers *QueryContractExtractResultHeaders, runtime *util.RuntimeOptions) (_result *QueryContractExtractResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtractTaskId)) {
		body["extractTaskId"] = request.ExtractTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
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
		Action:      tea.String("QueryContractExtractResult"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/extractResults/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryContractExtractResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询合同提取结果
//
// @param request - QueryContractExtractResultRequest
//
// @return QueryContractExtractResultResponse
func (client *Client) QueryContractExtractResult(request *QueryContractExtractResultRequest) (_result *QueryContractExtractResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryContractExtractResultHeaders{}
	_result = &QueryContractExtractResultResponse{}
	_body, _err := client.QueryContractExtractResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询合同审查结果
//
// @param request - QueryContractReviewResultRequest
//
// @param headers - QueryContractReviewResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryContractReviewResultResponse
func (client *Client) QueryContractReviewResultWithOptions(request *QueryContractReviewResultRequest, headers *QueryContractReviewResultHeaders, runtime *util.RuntimeOptions) (_result *QueryContractReviewResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.ReviewTaskId)) {
		body["reviewTaskId"] = request.ReviewTaskId
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
		Action:      tea.String("QueryContractReviewResult"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/reviewResults/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryContractReviewResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询合同审查结果
//
// @param request - QueryContractReviewResultRequest
//
// @return QueryContractReviewResultResponse
func (client *Client) QueryContractReviewResult(request *QueryContractReviewResultRequest) (_result *QueryContractReviewResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryContractReviewResultHeaders{}
	_result = &QueryContractReviewResultResponse{}
	_body, _err := client.QueryContractReviewResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发送合同相关卡片
//
// @param request - SendContractCardRequest
//
// @param headers - SendContractCardHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendContractCardResponse
func (client *Client) SendContractCardWithOptions(request *SendContractCardRequest, headers *SendContractCardHeaders, runtime *util.RuntimeOptions) (_result *SendContractCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CardType)) {
		body["cardType"] = request.CardType
	}

	if !tea.BoolValue(util.IsUnset(request.ContractInfo)) {
		body["contractInfo"] = request.ContractInfo
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Extension)) {
		body["extension"] = request.Extension
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		body["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiveGroups)) {
		body["receiveGroups"] = request.ReceiveGroups
	}

	if !tea.BoolValue(util.IsUnset(request.Receivers)) {
		body["receivers"] = request.Receivers
	}

	if !tea.BoolValue(util.IsUnset(request.Sender)) {
		body["sender"] = request.Sender
	}

	if !tea.BoolValue(util.IsUnset(request.SyncSingleChat)) {
		body["syncSingleChat"] = request.SyncSingleChat
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
		Action:      tea.String("SendContractCard"),
		Version:     tea.String("contract_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contract/cards/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendContractCardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送合同相关卡片
//
// @param request - SendContractCardRequest
//
// @return SendContractCardResponse
func (client *Client) SendContractCard(request *SendContractCardRequest) (_result *SendContractCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendContractCardHeaders{}
	_result = &SendContractCardResponse{}
	_body, _err := client.SendContractCardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
