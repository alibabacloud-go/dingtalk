// This file is auto-generated, don't edit it. Thanks.
package aiot_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type PropertiesValue struct {
	Value     interface{} `json:"value,omitempty" xml:"value,omitempty"`
	Timestamp *string     `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s PropertiesValue) String() string {
	return tea.Prettify(s)
}

func (s PropertiesValue) GoString() string {
	return s.String()
}

func (s *PropertiesValue) SetValue(v interface{}) *PropertiesValue {
	s.Value = v
	return s
}

func (s *PropertiesValue) SetTimestamp(v string) *PropertiesValue {
	s.Timestamp = &v
	return s
}

type CheckDeviceUpdateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CheckDeviceUpdateHeaders) String() string {
	return tea.Prettify(s)
}

func (s CheckDeviceUpdateHeaders) GoString() string {
	return s.String()
}

func (s *CheckDeviceUpdateHeaders) SetCommonHeaders(v map[string]*string) *CheckDeviceUpdateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CheckDeviceUpdateHeaders) SetXAcsDingtalkAccessToken(v string) *CheckDeviceUpdateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CheckDeviceUpdateRequest struct {
	Body []*CheckDeviceUpdateRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s CheckDeviceUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckDeviceUpdateRequest) GoString() string {
	return s.String()
}

func (s *CheckDeviceUpdateRequest) SetBody(v []*CheckDeviceUpdateRequestBody) *CheckDeviceUpdateRequest {
	s.Body = v
	return s
}

type CheckDeviceUpdateRequestBody struct {
	CurrentVersion *string `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	ModuleName     *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
}

func (s CheckDeviceUpdateRequestBody) String() string {
	return tea.Prettify(s)
}

func (s CheckDeviceUpdateRequestBody) GoString() string {
	return s.String()
}

func (s *CheckDeviceUpdateRequestBody) SetCurrentVersion(v string) *CheckDeviceUpdateRequestBody {
	s.CurrentVersion = &v
	return s
}

func (s *CheckDeviceUpdateRequestBody) SetModuleName(v string) *CheckDeviceUpdateRequestBody {
	s.ModuleName = &v
	return s
}

type CheckDeviceUpdateResponseBody struct {
	Modules []*CheckDeviceUpdateResponseBodyModules `json:"modules,omitempty" xml:"modules,omitempty" type:"Repeated"`
}

func (s CheckDeviceUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckDeviceUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDeviceUpdateResponseBody) SetModules(v []*CheckDeviceUpdateResponseBodyModules) *CheckDeviceUpdateResponseBody {
	s.Modules = v
	return s
}

type CheckDeviceUpdateResponseBodyModules struct {
	Checksum          *string `json:"checksum,omitempty" xml:"checksum,omitempty"`
	ChecksumAlgorithm *string `json:"checksumAlgorithm,omitempty" xml:"checksumAlgorithm,omitempty"`
	CriticalNext      *string `json:"criticalNext,omitempty" xml:"criticalNext,omitempty"`
	CurrentVersion    *string `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	FileUrl           *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	Latest            *string `json:"latest,omitempty" xml:"latest,omitempty"`
	ModuleName        *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	NoticeEn          *string `json:"noticeEn,omitempty" xml:"noticeEn,omitempty"`
	NoticeZh          *string `json:"noticeZh,omitempty" xml:"noticeZh,omitempty"`
	UpgradeMode       *string `json:"upgradeMode,omitempty" xml:"upgradeMode,omitempty"`
}

func (s CheckDeviceUpdateResponseBodyModules) String() string {
	return tea.Prettify(s)
}

func (s CheckDeviceUpdateResponseBodyModules) GoString() string {
	return s.String()
}

func (s *CheckDeviceUpdateResponseBodyModules) SetChecksum(v string) *CheckDeviceUpdateResponseBodyModules {
	s.Checksum = &v
	return s
}

func (s *CheckDeviceUpdateResponseBodyModules) SetChecksumAlgorithm(v string) *CheckDeviceUpdateResponseBodyModules {
	s.ChecksumAlgorithm = &v
	return s
}

func (s *CheckDeviceUpdateResponseBodyModules) SetCriticalNext(v string) *CheckDeviceUpdateResponseBodyModules {
	s.CriticalNext = &v
	return s
}

func (s *CheckDeviceUpdateResponseBodyModules) SetCurrentVersion(v string) *CheckDeviceUpdateResponseBodyModules {
	s.CurrentVersion = &v
	return s
}

func (s *CheckDeviceUpdateResponseBodyModules) SetFileUrl(v string) *CheckDeviceUpdateResponseBodyModules {
	s.FileUrl = &v
	return s
}

func (s *CheckDeviceUpdateResponseBodyModules) SetLatest(v string) *CheckDeviceUpdateResponseBodyModules {
	s.Latest = &v
	return s
}

func (s *CheckDeviceUpdateResponseBodyModules) SetModuleName(v string) *CheckDeviceUpdateResponseBodyModules {
	s.ModuleName = &v
	return s
}

func (s *CheckDeviceUpdateResponseBodyModules) SetNoticeEn(v string) *CheckDeviceUpdateResponseBodyModules {
	s.NoticeEn = &v
	return s
}

func (s *CheckDeviceUpdateResponseBodyModules) SetNoticeZh(v string) *CheckDeviceUpdateResponseBodyModules {
	s.NoticeZh = &v
	return s
}

func (s *CheckDeviceUpdateResponseBodyModules) SetUpgradeMode(v string) *CheckDeviceUpdateResponseBodyModules {
	s.UpgradeMode = &v
	return s
}

type CheckDeviceUpdateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDeviceUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDeviceUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckDeviceUpdateResponse) GoString() string {
	return s.String()
}

func (s *CheckDeviceUpdateResponse) SetHeaders(v map[string]*string) *CheckDeviceUpdateResponse {
	s.Headers = v
	return s
}

func (s *CheckDeviceUpdateResponse) SetStatusCode(v int32) *CheckDeviceUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDeviceUpdateResponse) SetBody(v *CheckDeviceUpdateResponseBody) *CheckDeviceUpdateResponse {
	s.Body = v
	return s
}

type ConfirmFirmwareUpgradeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ConfirmFirmwareUpgradeHeaders) String() string {
	return tea.Prettify(s)
}

func (s ConfirmFirmwareUpgradeHeaders) GoString() string {
	return s.String()
}

func (s *ConfirmFirmwareUpgradeHeaders) SetCommonHeaders(v map[string]*string) *ConfirmFirmwareUpgradeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ConfirmFirmwareUpgradeHeaders) SetXAcsDingtalkAccessToken(v string) *ConfirmFirmwareUpgradeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ConfirmFirmwareUpgradeRequest struct {
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
}

func (s ConfirmFirmwareUpgradeRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfirmFirmwareUpgradeRequest) GoString() string {
	return s.String()
}

func (s *ConfirmFirmwareUpgradeRequest) SetModuleName(v string) *ConfirmFirmwareUpgradeRequest {
	s.ModuleName = &v
	return s
}

type ConfirmFirmwareUpgradeResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ConfirmFirmwareUpgradeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfirmFirmwareUpgradeResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmFirmwareUpgradeResponseBody) SetResult(v bool) *ConfirmFirmwareUpgradeResponseBody {
	s.Result = &v
	return s
}

type ConfirmFirmwareUpgradeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfirmFirmwareUpgradeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmFirmwareUpgradeResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfirmFirmwareUpgradeResponse) GoString() string {
	return s.String()
}

func (s *ConfirmFirmwareUpgradeResponse) SetHeaders(v map[string]*string) *ConfirmFirmwareUpgradeResponse {
	s.Headers = v
	return s
}

func (s *ConfirmFirmwareUpgradeResponse) SetStatusCode(v int32) *ConfirmFirmwareUpgradeResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmFirmwareUpgradeResponse) SetBody(v *ConfirmFirmwareUpgradeResponseBody) *ConfirmFirmwareUpgradeResponse {
	s.Body = v
	return s
}

type GetDeviceDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDeviceDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetDeviceDetailHeaders) SetCommonHeaders(v map[string]*string) *GetDeviceDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeviceDetailHeaders) SetXAcsDingtalkAccessToken(v string) *GetDeviceDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDeviceDetailResponseBody struct {
	ActivatedAt     *string `json:"activatedAt,omitempty" xml:"activatedAt,omitempty"`
	Connectivity    *string `json:"connectivity,omitempty" xml:"connectivity,omitempty"`
	LastOfflineTime *string `json:"lastOfflineTime,omitempty" xml:"lastOfflineTime,omitempty"`
	LastOnlineTime  *string `json:"lastOnlineTime,omitempty" xml:"lastOnlineTime,omitempty"`
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetDeviceDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceDetailResponseBody) SetActivatedAt(v string) *GetDeviceDetailResponseBody {
	s.ActivatedAt = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetConnectivity(v string) *GetDeviceDetailResponseBody {
	s.Connectivity = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetLastOfflineTime(v string) *GetDeviceDetailResponseBody {
	s.LastOfflineTime = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetLastOnlineTime(v string) *GetDeviceDetailResponseBody {
	s.LastOnlineTime = &v
	return s
}

func (s *GetDeviceDetailResponseBody) SetStatus(v string) *GetDeviceDetailResponseBody {
	s.Status = &v
	return s
}

type GetDeviceDetailResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceDetailResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceDetailResponse) SetHeaders(v map[string]*string) *GetDeviceDetailResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceDetailResponse) SetStatusCode(v int32) *GetDeviceDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceDetailResponse) SetBody(v *GetDeviceDetailResponseBody) *GetDeviceDetailResponse {
	s.Body = v
	return s
}

type GetDevicePropertiesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDevicePropertiesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDevicePropertiesHeaders) GoString() string {
	return s.String()
}

func (s *GetDevicePropertiesHeaders) SetCommonHeaders(v map[string]*string) *GetDevicePropertiesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDevicePropertiesHeaders) SetXAcsDingtalkAccessToken(v string) *GetDevicePropertiesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDevicePropertiesRequest struct {
	Body []*string `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s GetDevicePropertiesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDevicePropertiesRequest) GoString() string {
	return s.String()
}

func (s *GetDevicePropertiesRequest) SetBody(v []*string) *GetDevicePropertiesRequest {
	s.Body = v
	return s
}

type GetDevicePropertiesResponseBody struct {
	DeviceName *string                     `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	ProductKey *string                     `json:"productKey,omitempty" xml:"productKey,omitempty"`
	Properties map[string]*PropertiesValue `json:"properties,omitempty" xml:"properties,omitempty"`
	SnapshotAt *string                     `json:"snapshotAt,omitempty" xml:"snapshotAt,omitempty"`
}

func (s GetDevicePropertiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDevicePropertiesResponseBody) GoString() string {
	return s.String()
}

func (s *GetDevicePropertiesResponseBody) SetDeviceName(v string) *GetDevicePropertiesResponseBody {
	s.DeviceName = &v
	return s
}

func (s *GetDevicePropertiesResponseBody) SetProductKey(v string) *GetDevicePropertiesResponseBody {
	s.ProductKey = &v
	return s
}

func (s *GetDevicePropertiesResponseBody) SetProperties(v map[string]*PropertiesValue) *GetDevicePropertiesResponseBody {
	s.Properties = v
	return s
}

func (s *GetDevicePropertiesResponseBody) SetSnapshotAt(v string) *GetDevicePropertiesResponseBody {
	s.SnapshotAt = &v
	return s
}

type GetDevicePropertiesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDevicePropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDevicePropertiesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDevicePropertiesResponse) GoString() string {
	return s.String()
}

func (s *GetDevicePropertiesResponse) SetHeaders(v map[string]*string) *GetDevicePropertiesResponse {
	s.Headers = v
	return s
}

func (s *GetDevicePropertiesResponse) SetStatusCode(v int32) *GetDevicePropertiesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDevicePropertiesResponse) SetBody(v *GetDevicePropertiesResponseBody) *GetDevicePropertiesResponse {
	s.Body = v
	return s
}

type GetServiceInvocationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetServiceInvocationHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInvocationHeaders) GoString() string {
	return s.String()
}

func (s *GetServiceInvocationHeaders) SetCommonHeaders(v map[string]*string) *GetServiceInvocationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetServiceInvocationHeaders) SetXAcsDingtalkAccessToken(v string) *GetServiceInvocationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetServiceInvocationResponseBody struct {
	CompletedAt      *string                `json:"completedAt,omitempty" xml:"completedAt,omitempty"`
	CreatedAt        *string                `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	DeviceName       *string                `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	ErrorCode        *string                `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg         *string                `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Identifier       *string                `json:"identifier,omitempty" xml:"identifier,omitempty"`
	InvocationId     *string                `json:"invocationId,omitempty" xml:"invocationId,omitempty"`
	OutputData       map[string]interface{} `json:"outputData,omitempty" xml:"outputData,omitempty"`
	ProcessingTimeMs *int64                 `json:"processingTimeMs,omitempty" xml:"processingTimeMs,omitempty"`
	ProductKey       *string                `json:"productKey,omitempty" xml:"productKey,omitempty"`
	Status           *string                `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetServiceInvocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInvocationResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceInvocationResponseBody) SetCompletedAt(v string) *GetServiceInvocationResponseBody {
	s.CompletedAt = &v
	return s
}

func (s *GetServiceInvocationResponseBody) SetCreatedAt(v string) *GetServiceInvocationResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *GetServiceInvocationResponseBody) SetDeviceName(v string) *GetServiceInvocationResponseBody {
	s.DeviceName = &v
	return s
}

func (s *GetServiceInvocationResponseBody) SetErrorCode(v string) *GetServiceInvocationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetServiceInvocationResponseBody) SetErrorMsg(v string) *GetServiceInvocationResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetServiceInvocationResponseBody) SetIdentifier(v string) *GetServiceInvocationResponseBody {
	s.Identifier = &v
	return s
}

func (s *GetServiceInvocationResponseBody) SetInvocationId(v string) *GetServiceInvocationResponseBody {
	s.InvocationId = &v
	return s
}

func (s *GetServiceInvocationResponseBody) SetOutputData(v map[string]interface{}) *GetServiceInvocationResponseBody {
	s.OutputData = v
	return s
}

func (s *GetServiceInvocationResponseBody) SetProcessingTimeMs(v int64) *GetServiceInvocationResponseBody {
	s.ProcessingTimeMs = &v
	return s
}

func (s *GetServiceInvocationResponseBody) SetProductKey(v string) *GetServiceInvocationResponseBody {
	s.ProductKey = &v
	return s
}

func (s *GetServiceInvocationResponseBody) SetStatus(v string) *GetServiceInvocationResponseBody {
	s.Status = &v
	return s
}

type GetServiceInvocationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceInvocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceInvocationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInvocationResponse) GoString() string {
	return s.String()
}

func (s *GetServiceInvocationResponse) SetHeaders(v map[string]*string) *GetServiceInvocationResponse {
	s.Headers = v
	return s
}

func (s *GetServiceInvocationResponse) SetStatusCode(v int32) *GetServiceInvocationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceInvocationResponse) SetBody(v *GetServiceInvocationResponseBody) *GetServiceInvocationResponse {
	s.Body = v
	return s
}

type InvokeDeviceServiceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InvokeDeviceServiceHeaders) String() string {
	return tea.Prettify(s)
}

func (s InvokeDeviceServiceHeaders) GoString() string {
	return s.String()
}

func (s *InvokeDeviceServiceHeaders) SetCommonHeaders(v map[string]*string) *InvokeDeviceServiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvokeDeviceServiceHeaders) SetXAcsDingtalkAccessToken(v string) *InvokeDeviceServiceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InvokeDeviceServiceRequest struct {
	Args           map[string]interface{} `json:"args,omitempty" xml:"args,omitempty"`
	TimeoutSeconds *int64                 `json:"timeoutSeconds,omitempty" xml:"timeoutSeconds,omitempty"`
}

func (s InvokeDeviceServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokeDeviceServiceRequest) GoString() string {
	return s.String()
}

func (s *InvokeDeviceServiceRequest) SetArgs(v map[string]interface{}) *InvokeDeviceServiceRequest {
	s.Args = v
	return s
}

func (s *InvokeDeviceServiceRequest) SetTimeoutSeconds(v int64) *InvokeDeviceServiceRequest {
	s.TimeoutSeconds = &v
	return s
}

type InvokeDeviceServiceResponseBody struct {
	ErrorCode    *string                `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg     *string                `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	InvocationId *string                `json:"invocationId,omitempty" xml:"invocationId,omitempty"`
	OutputData   map[string]interface{} `json:"outputData,omitempty" xml:"outputData,omitempty"`
	Status       *string                `json:"status,omitempty" xml:"status,omitempty"`
}

func (s InvokeDeviceServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvokeDeviceServiceResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeDeviceServiceResponseBody) SetErrorCode(v string) *InvokeDeviceServiceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *InvokeDeviceServiceResponseBody) SetErrorMsg(v string) *InvokeDeviceServiceResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *InvokeDeviceServiceResponseBody) SetInvocationId(v string) *InvokeDeviceServiceResponseBody {
	s.InvocationId = &v
	return s
}

func (s *InvokeDeviceServiceResponseBody) SetOutputData(v map[string]interface{}) *InvokeDeviceServiceResponseBody {
	s.OutputData = v
	return s
}

func (s *InvokeDeviceServiceResponseBody) SetStatus(v string) *InvokeDeviceServiceResponseBody {
	s.Status = &v
	return s
}

type InvokeDeviceServiceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvokeDeviceServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokeDeviceServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s InvokeDeviceServiceResponse) GoString() string {
	return s.String()
}

func (s *InvokeDeviceServiceResponse) SetHeaders(v map[string]*string) *InvokeDeviceServiceResponse {
	s.Headers = v
	return s
}

func (s *InvokeDeviceServiceResponse) SetStatusCode(v int32) *InvokeDeviceServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeDeviceServiceResponse) SetBody(v *InvokeDeviceServiceResponseBody) *InvokeDeviceServiceResponse {
	s.Body = v
	return s
}

type SetDevicePropertiesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetDevicePropertiesHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetDevicePropertiesHeaders) GoString() string {
	return s.String()
}

func (s *SetDevicePropertiesHeaders) SetCommonHeaders(v map[string]*string) *SetDevicePropertiesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetDevicePropertiesHeaders) SetXAcsDingtalkAccessToken(v string) *SetDevicePropertiesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetDevicePropertiesRequest struct {
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDevicePropertiesRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDevicePropertiesRequest) GoString() string {
	return s.String()
}

func (s *SetDevicePropertiesRequest) SetBody(v map[string]interface{}) *SetDevicePropertiesRequest {
	s.Body = v
	return s
}

type SetDevicePropertiesResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SetDevicePropertiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDevicePropertiesResponseBody) GoString() string {
	return s.String()
}

func (s *SetDevicePropertiesResponseBody) SetResult(v bool) *SetDevicePropertiesResponseBody {
	s.Result = &v
	return s
}

type SetDevicePropertiesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDevicePropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDevicePropertiesResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDevicePropertiesResponse) GoString() string {
	return s.String()
}

func (s *SetDevicePropertiesResponse) SetHeaders(v map[string]*string) *SetDevicePropertiesResponse {
	s.Headers = v
	return s
}

func (s *SetDevicePropertiesResponse) SetStatusCode(v int32) *SetDevicePropertiesResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDevicePropertiesResponse) SetBody(v *SetDevicePropertiesResponseBody) *SetDevicePropertiesResponse {
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
// 检查指定设备的固件升级
//
// @param request - CheckDeviceUpdateRequest
//
// @param headers - CheckDeviceUpdateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDeviceUpdateResponse
func (client *Client) CheckDeviceUpdateWithOptions(productKey *string, deviceName *string, request *CheckDeviceUpdateRequest, headers *CheckDeviceUpdateHeaders, runtime *util.RuntimeOptions) (_result *CheckDeviceUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
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
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckDeviceUpdate"),
		Version:     tea.String("aiot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiot/products/" + tea.StringValue(productKey) + "/devices/" + tea.StringValue(deviceName) + "/firmware/checkUpdate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckDeviceUpdateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查指定设备的固件升级
//
// @param request - CheckDeviceUpdateRequest
//
// @return CheckDeviceUpdateResponse
func (client *Client) CheckDeviceUpdate(productKey *string, deviceName *string, request *CheckDeviceUpdateRequest) (_result *CheckDeviceUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CheckDeviceUpdateHeaders{}
	_result = &CheckDeviceUpdateResponse{}
	_body, _err := client.CheckDeviceUpdateWithOptions(productKey, deviceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 确认执行设备固件升级
//
// @param request - ConfirmFirmwareUpgradeRequest
//
// @param headers - ConfirmFirmwareUpgradeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfirmFirmwareUpgradeResponse
func (client *Client) ConfirmFirmwareUpgradeWithOptions(productKey *string, deviceName *string, request *ConfirmFirmwareUpgradeRequest, headers *ConfirmFirmwareUpgradeHeaders, runtime *util.RuntimeOptions) (_result *ConfirmFirmwareUpgradeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModuleName)) {
		query["moduleName"] = request.ModuleName
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
		Action:      tea.String("ConfirmFirmwareUpgrade"),
		Version:     tea.String("aiot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiot/products/" + tea.StringValue(productKey) + "/devices/" + tea.StringValue(deviceName) + "/firmware/confirmUpgrade"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfirmFirmwareUpgradeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 确认执行设备固件升级
//
// @param request - ConfirmFirmwareUpgradeRequest
//
// @return ConfirmFirmwareUpgradeResponse
func (client *Client) ConfirmFirmwareUpgrade(productKey *string, deviceName *string, request *ConfirmFirmwareUpgradeRequest) (_result *ConfirmFirmwareUpgradeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ConfirmFirmwareUpgradeHeaders{}
	_result = &ConfirmFirmwareUpgradeResponse{}
	_body, _err := client.ConfirmFirmwareUpgradeWithOptions(productKey, deviceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询指定设备的详情
//
// @param headers - GetDeviceDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceDetailResponse
func (client *Client) GetDeviceDetailWithOptions(productKey *string, deviceName *string, headers *GetDeviceDetailHeaders, runtime *util.RuntimeOptions) (_result *GetDeviceDetailResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceDetail"),
		Version:     tea.String("aiot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiot/products/" + tea.StringValue(productKey) + "/devices/" + tea.StringValue(deviceName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定设备的详情
//
// @return GetDeviceDetailResponse
func (client *Client) GetDeviceDetail(productKey *string, deviceName *string) (_result *GetDeviceDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDeviceDetailHeaders{}
	_result = &GetDeviceDetailResponse{}
	_body, _err := client.GetDeviceDetailWithOptions(productKey, deviceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 读取指定设备的属性快照
//
// @param request - GetDevicePropertiesRequest
//
// @param headers - GetDevicePropertiesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDevicePropertiesResponse
func (client *Client) GetDevicePropertiesWithOptions(productKey *string, deviceName *string, request *GetDevicePropertiesRequest, headers *GetDevicePropertiesHeaders, runtime *util.RuntimeOptions) (_result *GetDevicePropertiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
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
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceProperties"),
		Version:     tea.String("aiot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiot/products/" + tea.StringValue(productKey) + "/devices/" + tea.StringValue(deviceName) + "/properties/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDevicePropertiesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 读取指定设备的属性快照
//
// @param request - GetDevicePropertiesRequest
//
// @return GetDevicePropertiesResponse
func (client *Client) GetDeviceProperties(productKey *string, deviceName *string, request *GetDevicePropertiesRequest) (_result *GetDevicePropertiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDevicePropertiesHeaders{}
	_result = &GetDevicePropertiesResponse{}
	_body, _err := client.GetDevicePropertiesWithOptions(productKey, deviceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询设备服务调用记录
//
// @param headers - GetServiceInvocationHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceInvocationResponse
func (client *Client) GetServiceInvocationWithOptions(invocationId *string, headers *GetServiceInvocationHeaders, runtime *util.RuntimeOptions) (_result *GetServiceInvocationResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceInvocation"),
		Version:     tea.String("aiot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiot/serviceInvocations/" + tea.StringValue(invocationId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceInvocationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询设备服务调用记录
//
// @return GetServiceInvocationResponse
func (client *Client) GetServiceInvocation(invocationId *string) (_result *GetServiceInvocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetServiceInvocationHeaders{}
	_result = &GetServiceInvocationResponse{}
	_body, _err := client.GetServiceInvocationWithOptions(invocationId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用指定设备的物模型服务
//
// @param request - InvokeDeviceServiceRequest
//
// @param headers - InvokeDeviceServiceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeDeviceServiceResponse
func (client *Client) InvokeDeviceServiceWithOptions(productKey *string, deviceName *string, serviceIdentifier *string, request *InvokeDeviceServiceRequest, headers *InvokeDeviceServiceHeaders, runtime *util.RuntimeOptions) (_result *InvokeDeviceServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Args)) {
		body["args"] = request.Args
	}

	if !tea.BoolValue(util.IsUnset(request.TimeoutSeconds)) {
		body["timeoutSeconds"] = request.TimeoutSeconds
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
		Action:      tea.String("InvokeDeviceService"),
		Version:     tea.String("aiot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiot/products/" + tea.StringValue(productKey) + "/devices/" + tea.StringValue(deviceName) + "/services/" + tea.StringValue(serviceIdentifier) + "/invoke"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &InvokeDeviceServiceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用指定设备的物模型服务
//
// @param request - InvokeDeviceServiceRequest
//
// @return InvokeDeviceServiceResponse
func (client *Client) InvokeDeviceService(productKey *string, deviceName *string, serviceIdentifier *string, request *InvokeDeviceServiceRequest) (_result *InvokeDeviceServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InvokeDeviceServiceHeaders{}
	_result = &InvokeDeviceServiceResponse{}
	_body, _err := client.InvokeDeviceServiceWithOptions(productKey, deviceName, serviceIdentifier, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置指定设备的属性
//
// @param request - SetDevicePropertiesRequest
//
// @param headers - SetDevicePropertiesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDevicePropertiesResponse
func (client *Client) SetDevicePropertiesWithOptions(productKey *string, deviceName *string, request *SetDevicePropertiesRequest, headers *SetDevicePropertiesHeaders, runtime *util.RuntimeOptions) (_result *SetDevicePropertiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["body"] = request.Body
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
		Action:      tea.String("SetDeviceProperties"),
		Version:     tea.String("aiot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiot/products/" + tea.StringValue(productKey) + "/devices/" + tea.StringValue(deviceName) + "/properties"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDevicePropertiesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置指定设备的属性
//
// @param request - SetDevicePropertiesRequest
//
// @return SetDevicePropertiesResponse
func (client *Client) SetDeviceProperties(productKey *string, deviceName *string, request *SetDevicePropertiesRequest) (_result *SetDevicePropertiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetDevicePropertiesHeaders{}
	_result = &SetDevicePropertiesResponse{}
	_body, _err := client.SetDevicePropertiesWithOptions(productKey, deviceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
