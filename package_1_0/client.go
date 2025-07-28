// This file is auto-generated, don't edit it. Thanks.
package package_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CloseHPackageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CloseHPackageHeaders) String() string {
	return tea.Prettify(s)
}

func (s CloseHPackageHeaders) GoString() string {
	return s.String()
}

func (s *CloseHPackageHeaders) SetCommonHeaders(v map[string]*string) *CloseHPackageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CloseHPackageHeaders) SetXAcsDingtalkAccessToken(v string) *CloseHPackageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CloseHPackageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5000000000000000
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
}

func (s CloseHPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseHPackageRequest) GoString() string {
	return s.String()
}

func (s *CloseHPackageRequest) SetMiniAppId(v string) *CloseHPackageRequest {
	s.MiniAppId = &v
	return s
}

type CloseHPackageResponseBody struct {
	Result interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CloseHPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseHPackageResponseBody) GoString() string {
	return s.String()
}

func (s *CloseHPackageResponseBody) SetResult(v interface{}) *CloseHPackageResponseBody {
	s.Result = v
	return s
}

type CloseHPackageResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseHPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseHPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseHPackageResponse) GoString() string {
	return s.String()
}

func (s *CloseHPackageResponse) SetHeaders(v map[string]*string) *CloseHPackageResponse {
	s.Headers = v
	return s
}

func (s *CloseHPackageResponse) SetStatusCode(v int32) *CloseHPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseHPackageResponse) SetBody(v *CloseHPackageResponseBody) *CloseHPackageResponse {
	s.Body = v
	return s
}

type GetUploadTokenHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUploadTokenHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUploadTokenHeaders) GoString() string {
	return s.String()
}

func (s *GetUploadTokenHeaders) SetCommonHeaders(v map[string]*string) *GetUploadTokenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUploadTokenHeaders) SetXAcsDingtalkAccessToken(v string) *GetUploadTokenHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUploadTokenRequest struct {
	// example:
	//
	// 5000000002761167
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
}

func (s GetUploadTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUploadTokenRequest) GoString() string {
	return s.String()
}

func (s *GetUploadTokenRequest) SetMiniAppId(v string) *GetUploadTokenRequest {
	s.MiniAppId = &v
	return s
}

type GetUploadTokenResponseBody struct {
	// example:
	//
	// STS.NUPjgnMhCVWvo1HSxfftf
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// example:
	//
	// ASviryNDy9tTuS5KiYMA6fCYf81vHg4KdoX7CVHz4CSx
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	// example:
	//
	// dingtalk-bucket
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// example:
	//
	// oss-cn-shanghai.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// example:
	//
	// 2022-09-21T09:32:16Z
	Expiration *string `json:"expiration,omitempty" xml:"expiration,omitempty"`
	// example:
	//
	// 5000000002761167/1663751835956
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// oss-cn-shanghai
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// CAIS0QJ1q6Ft5B2yfSjIr5blId3aoLdi4ZWdbRf5t3gzavt...
	StsToken *string `json:"stsToken,omitempty" xml:"stsToken,omitempty"`
}

func (s GetUploadTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUploadTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadTokenResponseBody) SetAccessKeyId(v string) *GetUploadTokenResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *GetUploadTokenResponseBody) SetAccessKeySecret(v string) *GetUploadTokenResponseBody {
	s.AccessKeySecret = &v
	return s
}

func (s *GetUploadTokenResponseBody) SetBucket(v string) *GetUploadTokenResponseBody {
	s.Bucket = &v
	return s
}

func (s *GetUploadTokenResponseBody) SetEndpoint(v string) *GetUploadTokenResponseBody {
	s.Endpoint = &v
	return s
}

func (s *GetUploadTokenResponseBody) SetExpiration(v string) *GetUploadTokenResponseBody {
	s.Expiration = &v
	return s
}

func (s *GetUploadTokenResponseBody) SetName(v string) *GetUploadTokenResponseBody {
	s.Name = &v
	return s
}

func (s *GetUploadTokenResponseBody) SetRegion(v string) *GetUploadTokenResponseBody {
	s.Region = &v
	return s
}

func (s *GetUploadTokenResponseBody) SetStsToken(v string) *GetUploadTokenResponseBody {
	s.StsToken = &v
	return s
}

type GetUploadTokenResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUploadTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUploadTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUploadTokenResponse) GoString() string {
	return s.String()
}

func (s *GetUploadTokenResponse) SetHeaders(v map[string]*string) *GetUploadTokenResponse {
	s.Headers = v
	return s
}

func (s *GetUploadTokenResponse) SetStatusCode(v int32) *GetUploadTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUploadTokenResponse) SetBody(v *GetUploadTokenResponseBody) *GetUploadTokenResponse {
	s.Body = v
	return s
}

type HPackageListGetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HPackageListGetHeaders) String() string {
	return tea.Prettify(s)
}

func (s HPackageListGetHeaders) GoString() string {
	return s.String()
}

func (s *HPackageListGetHeaders) SetCommonHeaders(v map[string]*string) *HPackageListGetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HPackageListGetHeaders) SetXAcsDingtalkAccessToken(v string) *HPackageListGetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HPackageListGetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5000000000000000
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	// example:
	//
	// 0
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s HPackageListGetRequest) String() string {
	return tea.Prettify(s)
}

func (s HPackageListGetRequest) GoString() string {
	return s.String()
}

func (s *HPackageListGetRequest) SetMiniAppId(v string) *HPackageListGetRequest {
	s.MiniAppId = &v
	return s
}

func (s *HPackageListGetRequest) SetPageNumber(v int64) *HPackageListGetRequest {
	s.PageNumber = &v
	return s
}

func (s *HPackageListGetRequest) SetPageSize(v int64) *HPackageListGetRequest {
	s.PageSize = &v
	return s
}

type HPackageListGetResponseBody struct {
	List       []*HPackageListGetResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	TotalCount *int64                             `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s HPackageListGetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HPackageListGetResponseBody) GoString() string {
	return s.String()
}

func (s *HPackageListGetResponseBody) SetList(v []*HPackageListGetResponseBodyList) *HPackageListGetResponseBody {
	s.List = v
	return s
}

func (s *HPackageListGetResponseBody) SetTotalCount(v int64) *HPackageListGetResponseBody {
	s.TotalCount = &v
	return s
}

type HPackageListGetResponseBodyList struct {
	// example:
	//
	// 1
	Avaliable *int64  `json:"avaliable,omitempty" xml:"avaliable,omitempty"`
	Creator   *string `json:"creator,omitempty" xml:"creator,omitempty"`
	Finished  *bool   `json:"finished,omitempty" xml:"finished,omitempty"`
	// example:
	//
	// 1669261911344
	OperationTime *int64 `json:"operationTime,omitempty" xml:"operationTime,omitempty"`
	// example:
	//
	// 100
	PackageSize *int64 `json:"packageSize,omitempty" xml:"packageSize,omitempty"`
	// example:
	//
	// 1
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 00000000Azksf
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 0.0.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s HPackageListGetResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s HPackageListGetResponseBodyList) GoString() string {
	return s.String()
}

func (s *HPackageListGetResponseBodyList) SetAvaliable(v int64) *HPackageListGetResponseBodyList {
	s.Avaliable = &v
	return s
}

func (s *HPackageListGetResponseBodyList) SetCreator(v string) *HPackageListGetResponseBodyList {
	s.Creator = &v
	return s
}

func (s *HPackageListGetResponseBodyList) SetFinished(v bool) *HPackageListGetResponseBodyList {
	s.Finished = &v
	return s
}

func (s *HPackageListGetResponseBodyList) SetOperationTime(v int64) *HPackageListGetResponseBodyList {
	s.OperationTime = &v
	return s
}

func (s *HPackageListGetResponseBodyList) SetPackageSize(v int64) *HPackageListGetResponseBodyList {
	s.PackageSize = &v
	return s
}

func (s *HPackageListGetResponseBodyList) SetStatus(v string) *HPackageListGetResponseBodyList {
	s.Status = &v
	return s
}

func (s *HPackageListGetResponseBodyList) SetTaskId(v string) *HPackageListGetResponseBodyList {
	s.TaskId = &v
	return s
}

func (s *HPackageListGetResponseBodyList) SetVersion(v string) *HPackageListGetResponseBodyList {
	s.Version = &v
	return s
}

type HPackageListGetResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HPackageListGetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HPackageListGetResponse) String() string {
	return tea.Prettify(s)
}

func (s HPackageListGetResponse) GoString() string {
	return s.String()
}

func (s *HPackageListGetResponse) SetHeaders(v map[string]*string) *HPackageListGetResponse {
	s.Headers = v
	return s
}

func (s *HPackageListGetResponse) SetStatusCode(v int32) *HPackageListGetResponse {
	s.StatusCode = &v
	return s
}

func (s *HPackageListGetResponse) SetBody(v *HPackageListGetResponseBody) *HPackageListGetResponse {
	s.Body = v
	return s
}

type HPublishPackageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HPublishPackageHeaders) String() string {
	return tea.Prettify(s)
}

func (s HPublishPackageHeaders) GoString() string {
	return s.String()
}

func (s *HPublishPackageHeaders) SetCommonHeaders(v map[string]*string) *HPublishPackageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HPublishPackageHeaders) SetXAcsDingtalkAccessToken(v string) *HPublishPackageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HPublishPackageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5000000002761167
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.0.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s HPublishPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s HPublishPackageRequest) GoString() string {
	return s.String()
}

func (s *HPublishPackageRequest) SetMiniAppId(v string) *HPublishPackageRequest {
	s.MiniAppId = &v
	return s
}

func (s *HPublishPackageRequest) SetVersion(v string) *HPublishPackageRequest {
	s.Version = &v
	return s
}

type HPublishPackageResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HPublishPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HPublishPackageResponseBody) GoString() string {
	return s.String()
}

func (s *HPublishPackageResponseBody) SetSuccess(v bool) *HPublishPackageResponseBody {
	s.Success = &v
	return s
}

type HPublishPackageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HPublishPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HPublishPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s HPublishPackageResponse) GoString() string {
	return s.String()
}

func (s *HPublishPackageResponse) SetHeaders(v map[string]*string) *HPublishPackageResponse {
	s.Headers = v
	return s
}

func (s *HPublishPackageResponse) SetStatusCode(v int32) *HPublishPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *HPublishPackageResponse) SetBody(v *HPublishPackageResponseBody) *HPublishPackageResponse {
	s.Body = v
	return s
}

type HUploadPackageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HUploadPackageHeaders) String() string {
	return tea.Prettify(s)
}

func (s HUploadPackageHeaders) GoString() string {
	return s.String()
}

func (s *HUploadPackageHeaders) SetCommonHeaders(v map[string]*string) *HUploadPackageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HUploadPackageHeaders) SetXAcsDingtalkAccessToken(v string) *HUploadPackageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HUploadPackageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5000000002793153
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	OssObjectKey *string `json:"ossObjectKey,omitempty" xml:"ossObjectKey,omitempty"`
}

func (s HUploadPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s HUploadPackageRequest) GoString() string {
	return s.String()
}

func (s *HUploadPackageRequest) SetMiniAppId(v string) *HUploadPackageRequest {
	s.MiniAppId = &v
	return s
}

func (s *HUploadPackageRequest) SetOssObjectKey(v string) *HUploadPackageRequest {
	s.OssObjectKey = &v
	return s
}

type HUploadPackageResponseBody struct {
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s HUploadPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HUploadPackageResponseBody) GoString() string {
	return s.String()
}

func (s *HUploadPackageResponseBody) SetTaskId(v string) *HUploadPackageResponseBody {
	s.TaskId = &v
	return s
}

type HUploadPackageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HUploadPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HUploadPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s HUploadPackageResponse) GoString() string {
	return s.String()
}

func (s *HUploadPackageResponse) SetHeaders(v map[string]*string) *HUploadPackageResponse {
	s.Headers = v
	return s
}

func (s *HUploadPackageResponse) SetStatusCode(v int32) *HUploadPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *HUploadPackageResponse) SetBody(v *HUploadPackageResponseBody) *HUploadPackageResponse {
	s.Body = v
	return s
}

type HUploadPackageStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HUploadPackageStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s HUploadPackageStatusHeaders) GoString() string {
	return s.String()
}

func (s *HUploadPackageStatusHeaders) SetCommonHeaders(v map[string]*string) *HUploadPackageStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HUploadPackageStatusHeaders) SetXAcsDingtalkAccessToken(v string) *HUploadPackageStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HUploadPackageStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5000000002761167
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s HUploadPackageStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s HUploadPackageStatusRequest) GoString() string {
	return s.String()
}

func (s *HUploadPackageStatusRequest) SetMiniAppId(v string) *HUploadPackageStatusRequest {
	s.MiniAppId = &v
	return s
}

func (s *HUploadPackageStatusRequest) SetTaskId(v string) *HUploadPackageStatusRequest {
	s.TaskId = &v
	return s
}

type HUploadPackageStatusResponseBody struct {
	// example:
	//
	// 1663743241146
	BuildTime *int64 `json:"buildTime,omitempty" xml:"buildTime,omitempty"`
	// This parameter is required.
	Finished *bool `json:"finished,omitempty" xml:"finished,omitempty"`
	// example:
	//
	// 100
	PackageSize *int64 `json:"packageSize,omitempty" xml:"packageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1663748308644pjpF
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.0.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s HUploadPackageStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HUploadPackageStatusResponseBody) GoString() string {
	return s.String()
}

func (s *HUploadPackageStatusResponseBody) SetBuildTime(v int64) *HUploadPackageStatusResponseBody {
	s.BuildTime = &v
	return s
}

func (s *HUploadPackageStatusResponseBody) SetFinished(v bool) *HUploadPackageStatusResponseBody {
	s.Finished = &v
	return s
}

func (s *HUploadPackageStatusResponseBody) SetPackageSize(v int64) *HUploadPackageStatusResponseBody {
	s.PackageSize = &v
	return s
}

func (s *HUploadPackageStatusResponseBody) SetStatus(v string) *HUploadPackageStatusResponseBody {
	s.Status = &v
	return s
}

func (s *HUploadPackageStatusResponseBody) SetTaskId(v string) *HUploadPackageStatusResponseBody {
	s.TaskId = &v
	return s
}

func (s *HUploadPackageStatusResponseBody) SetVersion(v string) *HUploadPackageStatusResponseBody {
	s.Version = &v
	return s
}

type HUploadPackageStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HUploadPackageStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HUploadPackageStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s HUploadPackageStatusResponse) GoString() string {
	return s.String()
}

func (s *HUploadPackageStatusResponse) SetHeaders(v map[string]*string) *HUploadPackageStatusResponse {
	s.Headers = v
	return s
}

func (s *HUploadPackageStatusResponse) SetStatusCode(v int32) *HUploadPackageStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *HUploadPackageStatusResponse) SetBody(v *HUploadPackageStatusResponseBody) *HUploadPackageStatusResponse {
	s.Body = v
	return s
}

type OpenMicroAppPackageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s OpenMicroAppPackageHeaders) String() string {
	return tea.Prettify(s)
}

func (s OpenMicroAppPackageHeaders) GoString() string {
	return s.String()
}

func (s *OpenMicroAppPackageHeaders) SetCommonHeaders(v map[string]*string) *OpenMicroAppPackageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OpenMicroAppPackageHeaders) SetXAcsDingtalkAccessToken(v string) *OpenMicroAppPackageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type OpenMicroAppPackageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
}

func (s OpenMicroAppPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenMicroAppPackageRequest) GoString() string {
	return s.String()
}

func (s *OpenMicroAppPackageRequest) SetAgentId(v int64) *OpenMicroAppPackageRequest {
	s.AgentId = &v
	return s
}

type OpenMicroAppPackageResponseBody struct {
	// example:
	//
	// 5000000000000000
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
}

func (s OpenMicroAppPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenMicroAppPackageResponseBody) GoString() string {
	return s.String()
}

func (s *OpenMicroAppPackageResponseBody) SetMiniAppId(v string) *OpenMicroAppPackageResponseBody {
	s.MiniAppId = &v
	return s
}

type OpenMicroAppPackageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenMicroAppPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenMicroAppPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenMicroAppPackageResponse) GoString() string {
	return s.String()
}

func (s *OpenMicroAppPackageResponse) SetHeaders(v map[string]*string) *OpenMicroAppPackageResponse {
	s.Headers = v
	return s
}

func (s *OpenMicroAppPackageResponse) SetStatusCode(v int32) *OpenMicroAppPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenMicroAppPackageResponse) SetBody(v *OpenMicroAppPackageResponseBody) *OpenMicroAppPackageResponse {
	s.Body = v
	return s
}

type ReleaseGrayDeployHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ReleaseGrayDeployHeaders) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayDeployHeaders) GoString() string {
	return s.String()
}

func (s *ReleaseGrayDeployHeaders) SetCommonHeaders(v map[string]*string) *ReleaseGrayDeployHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReleaseGrayDeployHeaders) SetXAcsDingtalkAccessToken(v string) *ReleaseGrayDeployHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ReleaseGrayDeployRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5000000000000000
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.0.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ReleaseGrayDeployRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayDeployRequest) GoString() string {
	return s.String()
}

func (s *ReleaseGrayDeployRequest) SetMiniAppId(v string) *ReleaseGrayDeployRequest {
	s.MiniAppId = &v
	return s
}

func (s *ReleaseGrayDeployRequest) SetVersion(v string) *ReleaseGrayDeployRequest {
	s.Version = &v
	return s
}

type ReleaseGrayDeployResponseBody struct {
	Result interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ReleaseGrayDeployResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayDeployResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseGrayDeployResponseBody) SetResult(v interface{}) *ReleaseGrayDeployResponseBody {
	s.Result = v
	return s
}

type ReleaseGrayDeployResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseGrayDeployResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseGrayDeployResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayDeployResponse) GoString() string {
	return s.String()
}

func (s *ReleaseGrayDeployResponse) SetHeaders(v map[string]*string) *ReleaseGrayDeployResponse {
	s.Headers = v
	return s
}

func (s *ReleaseGrayDeployResponse) SetStatusCode(v int32) *ReleaseGrayDeployResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseGrayDeployResponse) SetBody(v *ReleaseGrayDeployResponseBody) *ReleaseGrayDeployResponse {
	s.Body = v
	return s
}

type ReleaseGrayExitHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ReleaseGrayExitHeaders) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayExitHeaders) GoString() string {
	return s.String()
}

func (s *ReleaseGrayExitHeaders) SetCommonHeaders(v map[string]*string) *ReleaseGrayExitHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReleaseGrayExitHeaders) SetXAcsDingtalkAccessToken(v string) *ReleaseGrayExitHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ReleaseGrayExitRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5000000000000000
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.0.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ReleaseGrayExitRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayExitRequest) GoString() string {
	return s.String()
}

func (s *ReleaseGrayExitRequest) SetMiniAppId(v string) *ReleaseGrayExitRequest {
	s.MiniAppId = &v
	return s
}

func (s *ReleaseGrayExitRequest) SetVersion(v string) *ReleaseGrayExitRequest {
	s.Version = &v
	return s
}

type ReleaseGrayExitResponseBody struct {
	Reuslt interface{} `json:"reuslt,omitempty" xml:"reuslt,omitempty"`
}

func (s ReleaseGrayExitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayExitResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseGrayExitResponseBody) SetReuslt(v interface{}) *ReleaseGrayExitResponseBody {
	s.Reuslt = v
	return s
}

type ReleaseGrayExitResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseGrayExitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseGrayExitResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayExitResponse) GoString() string {
	return s.String()
}

func (s *ReleaseGrayExitResponse) SetHeaders(v map[string]*string) *ReleaseGrayExitResponse {
	s.Headers = v
	return s
}

func (s *ReleaseGrayExitResponse) SetStatusCode(v int32) *ReleaseGrayExitResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseGrayExitResponse) SetBody(v *ReleaseGrayExitResponseBody) *ReleaseGrayExitResponse {
	s.Body = v
	return s
}

type ReleaseGrayOrgGetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ReleaseGrayOrgGetHeaders) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayOrgGetHeaders) GoString() string {
	return s.String()
}

func (s *ReleaseGrayOrgGetHeaders) SetCommonHeaders(v map[string]*string) *ReleaseGrayOrgGetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReleaseGrayOrgGetHeaders) SetXAcsDingtalkAccessToken(v string) *ReleaseGrayOrgGetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ReleaseGrayOrgGetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5000000000000000
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.0.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ReleaseGrayOrgGetRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayOrgGetRequest) GoString() string {
	return s.String()
}

func (s *ReleaseGrayOrgGetRequest) SetMiniAppId(v string) *ReleaseGrayOrgGetRequest {
	s.MiniAppId = &v
	return s
}

func (s *ReleaseGrayOrgGetRequest) SetVersion(v string) *ReleaseGrayOrgGetRequest {
	s.Version = &v
	return s
}

type ReleaseGrayOrgGetResponseBody struct {
	Value []*string `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s ReleaseGrayOrgGetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayOrgGetResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseGrayOrgGetResponseBody) SetValue(v []*string) *ReleaseGrayOrgGetResponseBody {
	s.Value = v
	return s
}

type ReleaseGrayOrgGetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseGrayOrgGetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseGrayOrgGetResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayOrgGetResponse) GoString() string {
	return s.String()
}

func (s *ReleaseGrayOrgGetResponse) SetHeaders(v map[string]*string) *ReleaseGrayOrgGetResponse {
	s.Headers = v
	return s
}

func (s *ReleaseGrayOrgGetResponse) SetStatusCode(v int32) *ReleaseGrayOrgGetResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseGrayOrgGetResponse) SetBody(v *ReleaseGrayOrgGetResponseBody) *ReleaseGrayOrgGetResponse {
	s.Body = v
	return s
}

type ReleaseGrayOrgSetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ReleaseGrayOrgSetHeaders) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayOrgSetHeaders) GoString() string {
	return s.String()
}

func (s *ReleaseGrayOrgSetHeaders) SetCommonHeaders(v map[string]*string) *ReleaseGrayOrgSetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReleaseGrayOrgSetHeaders) SetXAcsDingtalkAccessToken(v string) *ReleaseGrayOrgSetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ReleaseGrayOrgSetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5000000000000000
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	// This parameter is required.
	Value []*string `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 0.0.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ReleaseGrayOrgSetRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayOrgSetRequest) GoString() string {
	return s.String()
}

func (s *ReleaseGrayOrgSetRequest) SetMiniAppId(v string) *ReleaseGrayOrgSetRequest {
	s.MiniAppId = &v
	return s
}

func (s *ReleaseGrayOrgSetRequest) SetValue(v []*string) *ReleaseGrayOrgSetRequest {
	s.Value = v
	return s
}

func (s *ReleaseGrayOrgSetRequest) SetVersion(v string) *ReleaseGrayOrgSetRequest {
	s.Version = &v
	return s
}

type ReleaseGrayOrgSetResponseBody struct {
	Result interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ReleaseGrayOrgSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayOrgSetResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseGrayOrgSetResponseBody) SetResult(v interface{}) *ReleaseGrayOrgSetResponseBody {
	s.Result = v
	return s
}

type ReleaseGrayOrgSetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseGrayOrgSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseGrayOrgSetResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayOrgSetResponse) GoString() string {
	return s.String()
}

func (s *ReleaseGrayOrgSetResponse) SetHeaders(v map[string]*string) *ReleaseGrayOrgSetResponse {
	s.Headers = v
	return s
}

func (s *ReleaseGrayOrgSetResponse) SetStatusCode(v int32) *ReleaseGrayOrgSetResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseGrayOrgSetResponse) SetBody(v *ReleaseGrayOrgSetResponseBody) *ReleaseGrayOrgSetResponse {
	s.Body = v
	return s
}

type ReleaseGrayPercentGetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ReleaseGrayPercentGetHeaders) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayPercentGetHeaders) GoString() string {
	return s.String()
}

func (s *ReleaseGrayPercentGetHeaders) SetCommonHeaders(v map[string]*string) *ReleaseGrayPercentGetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReleaseGrayPercentGetHeaders) SetXAcsDingtalkAccessToken(v string) *ReleaseGrayPercentGetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ReleaseGrayPercentGetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5000000000000000
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.0.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ReleaseGrayPercentGetRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayPercentGetRequest) GoString() string {
	return s.String()
}

func (s *ReleaseGrayPercentGetRequest) SetMiniAppId(v string) *ReleaseGrayPercentGetRequest {
	s.MiniAppId = &v
	return s
}

func (s *ReleaseGrayPercentGetRequest) SetVersion(v string) *ReleaseGrayPercentGetRequest {
	s.Version = &v
	return s
}

type ReleaseGrayPercentGetResponseBody struct {
	Value *float32 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ReleaseGrayPercentGetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayPercentGetResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseGrayPercentGetResponseBody) SetValue(v float32) *ReleaseGrayPercentGetResponseBody {
	s.Value = &v
	return s
}

type ReleaseGrayPercentGetResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseGrayPercentGetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseGrayPercentGetResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayPercentGetResponse) GoString() string {
	return s.String()
}

func (s *ReleaseGrayPercentGetResponse) SetHeaders(v map[string]*string) *ReleaseGrayPercentGetResponse {
	s.Headers = v
	return s
}

func (s *ReleaseGrayPercentGetResponse) SetStatusCode(v int32) *ReleaseGrayPercentGetResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseGrayPercentGetResponse) SetBody(v *ReleaseGrayPercentGetResponseBody) *ReleaseGrayPercentGetResponse {
	s.Body = v
	return s
}

type ReleaseGrayPercentSetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ReleaseGrayPercentSetHeaders) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayPercentSetHeaders) GoString() string {
	return s.String()
}

func (s *ReleaseGrayPercentSetHeaders) SetCommonHeaders(v map[string]*string) *ReleaseGrayPercentSetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReleaseGrayPercentSetHeaders) SetXAcsDingtalkAccessToken(v string) *ReleaseGrayPercentSetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ReleaseGrayPercentSetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5000000000000000
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	Value *float64 `json:"value,omitempty" xml:"value,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.0.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ReleaseGrayPercentSetRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayPercentSetRequest) GoString() string {
	return s.String()
}

func (s *ReleaseGrayPercentSetRequest) SetMiniAppId(v string) *ReleaseGrayPercentSetRequest {
	s.MiniAppId = &v
	return s
}

func (s *ReleaseGrayPercentSetRequest) SetValue(v float64) *ReleaseGrayPercentSetRequest {
	s.Value = &v
	return s
}

func (s *ReleaseGrayPercentSetRequest) SetVersion(v string) *ReleaseGrayPercentSetRequest {
	s.Version = &v
	return s
}

type ReleaseGrayPercentSetResponseBody struct {
	Result interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ReleaseGrayPercentSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayPercentSetResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseGrayPercentSetResponseBody) SetResult(v interface{}) *ReleaseGrayPercentSetResponseBody {
	s.Result = v
	return s
}

type ReleaseGrayPercentSetResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseGrayPercentSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseGrayPercentSetResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayPercentSetResponse) GoString() string {
	return s.String()
}

func (s *ReleaseGrayPercentSetResponse) SetHeaders(v map[string]*string) *ReleaseGrayPercentSetResponse {
	s.Headers = v
	return s
}

func (s *ReleaseGrayPercentSetResponse) SetStatusCode(v int32) *ReleaseGrayPercentSetResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseGrayPercentSetResponse) SetBody(v *ReleaseGrayPercentSetResponseBody) *ReleaseGrayPercentSetResponse {
	s.Body = v
	return s
}

type ReleaseGrayUserIdGetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ReleaseGrayUserIdGetHeaders) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayUserIdGetHeaders) GoString() string {
	return s.String()
}

func (s *ReleaseGrayUserIdGetHeaders) SetCommonHeaders(v map[string]*string) *ReleaseGrayUserIdGetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReleaseGrayUserIdGetHeaders) SetXAcsDingtalkAccessToken(v string) *ReleaseGrayUserIdGetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ReleaseGrayUserIdGetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5000000000000000
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.0.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ReleaseGrayUserIdGetRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayUserIdGetRequest) GoString() string {
	return s.String()
}

func (s *ReleaseGrayUserIdGetRequest) SetMiniAppId(v string) *ReleaseGrayUserIdGetRequest {
	s.MiniAppId = &v
	return s
}

func (s *ReleaseGrayUserIdGetRequest) SetVersion(v string) *ReleaseGrayUserIdGetRequest {
	s.Version = &v
	return s
}

type ReleaseGrayUserIdGetResponseBody struct {
	Value []*string `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s ReleaseGrayUserIdGetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayUserIdGetResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseGrayUserIdGetResponseBody) SetValue(v []*string) *ReleaseGrayUserIdGetResponseBody {
	s.Value = v
	return s
}

type ReleaseGrayUserIdGetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseGrayUserIdGetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseGrayUserIdGetResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseGrayUserIdGetResponse) GoString() string {
	return s.String()
}

func (s *ReleaseGrayUserIdGetResponse) SetHeaders(v map[string]*string) *ReleaseGrayUserIdGetResponse {
	s.Headers = v
	return s
}

func (s *ReleaseGrayUserIdGetResponse) SetStatusCode(v int32) *ReleaseGrayUserIdGetResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseGrayUserIdGetResponse) SetBody(v *ReleaseGrayUserIdGetResponseBody) *ReleaseGrayUserIdGetResponse {
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
// 关闭企业自建应用H5离线包
//
// @param request - CloseHPackageRequest
//
// @param headers - CloseHPackageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseHPackageResponse
func (client *Client) CloseHPackageWithOptions(request *CloseHPackageRequest, headers *CloseHPackageHeaders, runtime *util.RuntimeOptions) (_result *CloseHPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		body["miniAppId"] = request.MiniAppId
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
		Action:      tea.String("CloseHPackage"),
		Version:     tea.String("package_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/package/h5/microApps/close"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseHPackageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关闭企业自建应用H5离线包
//
// @param request - CloseHPackageRequest
//
// @return CloseHPackageResponse
func (client *Client) CloseHPackage(request *CloseHPackageRequest) (_result *CloseHPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CloseHPackageHeaders{}
	_result = &CloseHPackageResponse{}
	_body, _err := client.CloseHPackageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取离线包上传凭证
//
// @param request - GetUploadTokenRequest
//
// @param headers - GetUploadTokenHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUploadTokenResponse
func (client *Client) GetUploadTokenWithOptions(request *GetUploadTokenRequest, headers *GetUploadTokenHeaders, runtime *util.RuntimeOptions) (_result *GetUploadTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		query["miniAppId"] = request.MiniAppId
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
		Action:      tea.String("GetUploadToken"),
		Version:     tea.String("package_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/package/uploadTokens"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUploadTokenResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取离线包上传凭证
//
// @param request - GetUploadTokenRequest
//
// @return GetUploadTokenResponse
func (client *Client) GetUploadToken(request *GetUploadTokenRequest) (_result *GetUploadTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUploadTokenHeaders{}
	_result = &GetUploadTokenResponse{}
	_body, _err := client.GetUploadTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取H5离线包版本列表
//
// @param request - HPackageListGetRequest
//
// @param headers - HPackageListGetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HPackageListGetResponse
func (client *Client) HPackageListGetWithOptions(request *HPackageListGetRequest, headers *HPackageListGetHeaders, runtime *util.RuntimeOptions) (_result *HPackageListGetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		query["miniAppId"] = request.MiniAppId
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("HPackageListGet"),
		Version:     tea.String("package_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/package/h5/versions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HPackageListGetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取H5离线包版本列表
//
// @param request - HPackageListGetRequest
//
// @return HPackageListGetResponse
func (client *Client) HPackageListGet(request *HPackageListGetRequest) (_result *HPackageListGetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HPackageListGetHeaders{}
	_result = &HPackageListGetResponse{}
	_body, _err := client.HPackageListGetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发布离线包
//
// @param request - HPublishPackageRequest
//
// @param headers - HPublishPackageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HPublishPackageResponse
func (client *Client) HPublishPackageWithOptions(request *HPublishPackageRequest, headers *HPublishPackageHeaders, runtime *util.RuntimeOptions) (_result *HPublishPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		body["miniAppId"] = request.MiniAppId
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["version"] = request.Version
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
		Action:      tea.String("HPublishPackage"),
		Version:     tea.String("package_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/package/h5/publish"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HPublishPackageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发布离线包
//
// @param request - HPublishPackageRequest
//
// @return HPublishPackageResponse
func (client *Client) HPublishPackage(request *HPublishPackageRequest) (_result *HPublishPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HPublishPackageHeaders{}
	_result = &HPublishPackageResponse{}
	_body, _err := client.HPublishPackageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上传H5离线包
//
// @param request - HUploadPackageRequest
//
// @param headers - HUploadPackageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HUploadPackageResponse
func (client *Client) HUploadPackageWithOptions(request *HUploadPackageRequest, headers *HUploadPackageHeaders, runtime *util.RuntimeOptions) (_result *HUploadPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		body["miniAppId"] = request.MiniAppId
	}

	if !tea.BoolValue(util.IsUnset(request.OssObjectKey)) {
		body["ossObjectKey"] = request.OssObjectKey
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
		Action:      tea.String("HUploadPackage"),
		Version:     tea.String("package_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/package/h5/asyncUpload"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HUploadPackageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传H5离线包
//
// @param request - HUploadPackageRequest
//
// @return HUploadPackageResponse
func (client *Client) HUploadPackage(request *HUploadPackageRequest) (_result *HUploadPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HUploadPackageHeaders{}
	_result = &HUploadPackageResponse{}
	_body, _err := client.HUploadPackageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上传H5离线包进度
//
// @param request - HUploadPackageStatusRequest
//
// @param headers - HUploadPackageStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HUploadPackageStatusResponse
func (client *Client) HUploadPackageStatusWithOptions(request *HUploadPackageStatusRequest, headers *HUploadPackageStatusHeaders, runtime *util.RuntimeOptions) (_result *HUploadPackageStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		query["miniAppId"] = request.MiniAppId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
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
		Action:      tea.String("HUploadPackageStatus"),
		Version:     tea.String("package_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/package/h5/uploadStatus"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HUploadPackageStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传H5离线包进度
//
// @param request - HUploadPackageStatusRequest
//
// @return HUploadPackageStatusResponse
func (client *Client) HUploadPackageStatus(request *HUploadPackageStatusRequest) (_result *HUploadPackageStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HUploadPackageStatusHeaders{}
	_result = &HUploadPackageStatusResponse{}
	_body, _err := client.HUploadPackageStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开启企业自建应用H5离线包
//
// @param request - OpenMicroAppPackageRequest
//
// @param headers - OpenMicroAppPackageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenMicroAppPackageResponse
func (client *Client) OpenMicroAppPackageWithOptions(request *OpenMicroAppPackageRequest, headers *OpenMicroAppPackageHeaders, runtime *util.RuntimeOptions) (_result *OpenMicroAppPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		body["agentId"] = request.AgentId
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
		Action:      tea.String("OpenMicroAppPackage"),
		Version:     tea.String("package_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/package/h5/microApps/open"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenMicroAppPackageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启企业自建应用H5离线包
//
// @param request - OpenMicroAppPackageRequest
//
// @return OpenMicroAppPackageResponse
func (client *Client) OpenMicroAppPackage(request *OpenMicroAppPackageRequest) (_result *OpenMicroAppPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OpenMicroAppPackageHeaders{}
	_result = &OpenMicroAppPackageResponse{}
	_body, _err := client.OpenMicroAppPackageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发布离线包到灰度
//
// @param request - ReleaseGrayDeployRequest
//
// @param headers - ReleaseGrayDeployHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseGrayDeployResponse
func (client *Client) ReleaseGrayDeployWithOptions(request *ReleaseGrayDeployRequest, headers *ReleaseGrayDeployHeaders, runtime *util.RuntimeOptions) (_result *ReleaseGrayDeployResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		body["miniAppId"] = request.MiniAppId
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["version"] = request.Version
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
		Action:      tea.String("ReleaseGrayDeploy"),
		Version:     tea.String("package_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/package/greys/deploy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseGrayDeployResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发布离线包到灰度
//
// @param request - ReleaseGrayDeployRequest
//
// @return ReleaseGrayDeployResponse
func (client *Client) ReleaseGrayDeploy(request *ReleaseGrayDeployRequest) (_result *ReleaseGrayDeployResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ReleaseGrayDeployHeaders{}
	_result = &ReleaseGrayDeployResponse{}
	_body, _err := client.ReleaseGrayDeployWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 退出灰度
//
// @param request - ReleaseGrayExitRequest
//
// @param headers - ReleaseGrayExitHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseGrayExitResponse
func (client *Client) ReleaseGrayExitWithOptions(request *ReleaseGrayExitRequest, headers *ReleaseGrayExitHeaders, runtime *util.RuntimeOptions) (_result *ReleaseGrayExitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		body["miniAppId"] = request.MiniAppId
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["version"] = request.Version
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
		Action:      tea.String("ReleaseGrayExit"),
		Version:     tea.String("package_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/package/greys/exit"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseGrayExitResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 退出灰度
//
// @param request - ReleaseGrayExitRequest
//
// @return ReleaseGrayExitResponse
func (client *Client) ReleaseGrayExit(request *ReleaseGrayExitRequest) (_result *ReleaseGrayExitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ReleaseGrayExitHeaders{}
	_result = &ReleaseGrayExitResponse{}
	_body, _err := client.ReleaseGrayExitWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业灰度白名单
//
// @param request - ReleaseGrayOrgGetRequest
//
// @param headers - ReleaseGrayOrgGetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseGrayOrgGetResponse
func (client *Client) ReleaseGrayOrgGetWithOptions(request *ReleaseGrayOrgGetRequest, headers *ReleaseGrayOrgGetHeaders, runtime *util.RuntimeOptions) (_result *ReleaseGrayOrgGetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		query["miniAppId"] = request.MiniAppId
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		query["version"] = request.Version
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
		Action:      tea.String("ReleaseGrayOrgGet"),
		Version:     tea.String("package_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/package/greys/organizations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseGrayOrgGetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业灰度白名单
//
// @param request - ReleaseGrayOrgGetRequest
//
// @return ReleaseGrayOrgGetResponse
func (client *Client) ReleaseGrayOrgGet(request *ReleaseGrayOrgGetRequest) (_result *ReleaseGrayOrgGetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ReleaseGrayOrgGetHeaders{}
	_result = &ReleaseGrayOrgGetResponse{}
	_body, _err := client.ReleaseGrayOrgGetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置企业灰度白名单
//
// @param request - ReleaseGrayOrgSetRequest
//
// @param headers - ReleaseGrayOrgSetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseGrayOrgSetResponse
func (client *Client) ReleaseGrayOrgSetWithOptions(request *ReleaseGrayOrgSetRequest, headers *ReleaseGrayOrgSetHeaders, runtime *util.RuntimeOptions) (_result *ReleaseGrayOrgSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		body["miniAppId"] = request.MiniAppId
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["value"] = request.Value
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["version"] = request.Version
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
		Action:      tea.String("ReleaseGrayOrgSet"),
		Version:     tea.String("package_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/package/greys/organizations/release"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseGrayOrgSetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置企业灰度白名单
//
// @param request - ReleaseGrayOrgSetRequest
//
// @return ReleaseGrayOrgSetResponse
func (client *Client) ReleaseGrayOrgSet(request *ReleaseGrayOrgSetRequest) (_result *ReleaseGrayOrgSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ReleaseGrayOrgSetHeaders{}
	_result = &ReleaseGrayOrgSetResponse{}
	_body, _err := client.ReleaseGrayOrgSetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取灰度离线包百分比值
//
// @param request - ReleaseGrayPercentGetRequest
//
// @param headers - ReleaseGrayPercentGetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseGrayPercentGetResponse
func (client *Client) ReleaseGrayPercentGetWithOptions(request *ReleaseGrayPercentGetRequest, headers *ReleaseGrayPercentGetHeaders, runtime *util.RuntimeOptions) (_result *ReleaseGrayPercentGetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		query["miniAppId"] = request.MiniAppId
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		query["version"] = request.Version
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
		Action:      tea.String("ReleaseGrayPercentGet"),
		Version:     tea.String("package_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/package/greys/users/percents"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseGrayPercentGetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取灰度离线包百分比值
//
// @param request - ReleaseGrayPercentGetRequest
//
// @return ReleaseGrayPercentGetResponse
func (client *Client) ReleaseGrayPercentGet(request *ReleaseGrayPercentGetRequest) (_result *ReleaseGrayPercentGetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ReleaseGrayPercentGetHeaders{}
	_result = &ReleaseGrayPercentGetResponse{}
	_body, _err := client.ReleaseGrayPercentGetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置灰度离线包百分比值
//
// @param request - ReleaseGrayPercentSetRequest
//
// @param headers - ReleaseGrayPercentSetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseGrayPercentSetResponse
func (client *Client) ReleaseGrayPercentSetWithOptions(request *ReleaseGrayPercentSetRequest, headers *ReleaseGrayPercentSetHeaders, runtime *util.RuntimeOptions) (_result *ReleaseGrayPercentSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		body["miniAppId"] = request.MiniAppId
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["value"] = request.Value
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["version"] = request.Version
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
		Action:      tea.String("ReleaseGrayPercentSet"),
		Version:     tea.String("package_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/package/greys/users/percents/release"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseGrayPercentSetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置灰度离线包百分比值
//
// @param request - ReleaseGrayPercentSetRequest
//
// @return ReleaseGrayPercentSetResponse
func (client *Client) ReleaseGrayPercentSet(request *ReleaseGrayPercentSetRequest) (_result *ReleaseGrayPercentSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ReleaseGrayPercentSetHeaders{}
	_result = &ReleaseGrayPercentSetResponse{}
	_body, _err := client.ReleaseGrayPercentSetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户灰度名单
//
// @param request - ReleaseGrayUserIdGetRequest
//
// @param headers - ReleaseGrayUserIdGetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseGrayUserIdGetResponse
func (client *Client) ReleaseGrayUserIdGetWithOptions(request *ReleaseGrayUserIdGetRequest, headers *ReleaseGrayUserIdGetHeaders, runtime *util.RuntimeOptions) (_result *ReleaseGrayUserIdGetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		query["miniAppId"] = request.MiniAppId
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		query["version"] = request.Version
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
		Action:      tea.String("ReleaseGrayUserIdGet"),
		Version:     tea.String("package_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/package/greys/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseGrayUserIdGetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户灰度名单
//
// @param request - ReleaseGrayUserIdGetRequest
//
// @return ReleaseGrayUserIdGetResponse
func (client *Client) ReleaseGrayUserIdGet(request *ReleaseGrayUserIdGetRequest) (_result *ReleaseGrayUserIdGetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ReleaseGrayUserIdGetHeaders{}
	_result = &ReleaseGrayUserIdGetResponse{}
	_body, _err := client.ReleaseGrayUserIdGetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
