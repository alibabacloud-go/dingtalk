// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package contact_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddContactHideBySceneSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddContactHideBySceneSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddContactHideBySceneSettingHeaders) GoString() string {
	return s.String()
}

func (s *AddContactHideBySceneSettingHeaders) SetCommonHeaders(v map[string]*string) *AddContactHideBySceneSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddContactHideBySceneSettingHeaders) SetXAcsDingtalkAccessToken(v string) *AddContactHideBySceneSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddContactHideBySceneSettingRequest struct {
	Description         *string                                                 `json:"description,omitempty" xml:"description,omitempty"`
	ExcludeDeptIds      []*int64                                                `json:"excludeDeptIds,omitempty" xml:"excludeDeptIds,omitempty" type:"Repeated"`
	ExcludeTagIds       []*int64                                                `json:"excludeTagIds,omitempty" xml:"excludeTagIds,omitempty" type:"Repeated"`
	ExcludeUserIds      []*string                                               `json:"excludeUserIds,omitempty" xml:"excludeUserIds,omitempty" type:"Repeated"`
	Name                *string                                                 `json:"name,omitempty" xml:"name,omitempty"`
	NodeListSceneConfig *AddContactHideBySceneSettingRequestNodeListSceneConfig `json:"nodeListSceneConfig,omitempty" xml:"nodeListSceneConfig,omitempty" type:"Struct"`
	ObjectDeptIds       []*int64                                                `json:"objectDeptIds,omitempty" xml:"objectDeptIds,omitempty" type:"Repeated"`
	ObjectTagIds        []*int64                                                `json:"objectTagIds,omitempty" xml:"objectTagIds,omitempty" type:"Repeated"`
	ObjectUserIds       []*string                                               `json:"objectUserIds,omitempty" xml:"objectUserIds,omitempty" type:"Repeated"`
	ProfileSceneConfig  *AddContactHideBySceneSettingRequestProfileSceneConfig  `json:"profileSceneConfig,omitempty" xml:"profileSceneConfig,omitempty" type:"Struct"`
	SearchSceneConfig   *AddContactHideBySceneSettingRequestSearchSceneConfig   `json:"searchSceneConfig,omitempty" xml:"searchSceneConfig,omitempty" type:"Struct"`
}

func (s AddContactHideBySceneSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s AddContactHideBySceneSettingRequest) GoString() string {
	return s.String()
}

func (s *AddContactHideBySceneSettingRequest) SetDescription(v string) *AddContactHideBySceneSettingRequest {
	s.Description = &v
	return s
}

func (s *AddContactHideBySceneSettingRequest) SetExcludeDeptIds(v []*int64) *AddContactHideBySceneSettingRequest {
	s.ExcludeDeptIds = v
	return s
}

func (s *AddContactHideBySceneSettingRequest) SetExcludeTagIds(v []*int64) *AddContactHideBySceneSettingRequest {
	s.ExcludeTagIds = v
	return s
}

func (s *AddContactHideBySceneSettingRequest) SetExcludeUserIds(v []*string) *AddContactHideBySceneSettingRequest {
	s.ExcludeUserIds = v
	return s
}

func (s *AddContactHideBySceneSettingRequest) SetName(v string) *AddContactHideBySceneSettingRequest {
	s.Name = &v
	return s
}

func (s *AddContactHideBySceneSettingRequest) SetNodeListSceneConfig(v *AddContactHideBySceneSettingRequestNodeListSceneConfig) *AddContactHideBySceneSettingRequest {
	s.NodeListSceneConfig = v
	return s
}

func (s *AddContactHideBySceneSettingRequest) SetObjectDeptIds(v []*int64) *AddContactHideBySceneSettingRequest {
	s.ObjectDeptIds = v
	return s
}

func (s *AddContactHideBySceneSettingRequest) SetObjectTagIds(v []*int64) *AddContactHideBySceneSettingRequest {
	s.ObjectTagIds = v
	return s
}

func (s *AddContactHideBySceneSettingRequest) SetObjectUserIds(v []*string) *AddContactHideBySceneSettingRequest {
	s.ObjectUserIds = v
	return s
}

func (s *AddContactHideBySceneSettingRequest) SetProfileSceneConfig(v *AddContactHideBySceneSettingRequestProfileSceneConfig) *AddContactHideBySceneSettingRequest {
	s.ProfileSceneConfig = v
	return s
}

func (s *AddContactHideBySceneSettingRequest) SetSearchSceneConfig(v *AddContactHideBySceneSettingRequestSearchSceneConfig) *AddContactHideBySceneSettingRequest {
	s.SearchSceneConfig = v
	return s
}

type AddContactHideBySceneSettingRequestNodeListSceneConfig struct {
	Active               *bool `json:"active,omitempty" xml:"active,omitempty"`
	DeptObjectIncludeEmp *bool `json:"deptObjectIncludeEmp,omitempty" xml:"deptObjectIncludeEmp,omitempty"`
}

func (s AddContactHideBySceneSettingRequestNodeListSceneConfig) String() string {
	return tea.Prettify(s)
}

func (s AddContactHideBySceneSettingRequestNodeListSceneConfig) GoString() string {
	return s.String()
}

func (s *AddContactHideBySceneSettingRequestNodeListSceneConfig) SetActive(v bool) *AddContactHideBySceneSettingRequestNodeListSceneConfig {
	s.Active = &v
	return s
}

func (s *AddContactHideBySceneSettingRequestNodeListSceneConfig) SetDeptObjectIncludeEmp(v bool) *AddContactHideBySceneSettingRequestNodeListSceneConfig {
	s.DeptObjectIncludeEmp = &v
	return s
}

type AddContactHideBySceneSettingRequestProfileSceneConfig struct {
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
}

func (s AddContactHideBySceneSettingRequestProfileSceneConfig) String() string {
	return tea.Prettify(s)
}

func (s AddContactHideBySceneSettingRequestProfileSceneConfig) GoString() string {
	return s.String()
}

func (s *AddContactHideBySceneSettingRequestProfileSceneConfig) SetActive(v bool) *AddContactHideBySceneSettingRequestProfileSceneConfig {
	s.Active = &v
	return s
}

type AddContactHideBySceneSettingRequestSearchSceneConfig struct {
	Active               *bool `json:"active,omitempty" xml:"active,omitempty"`
	DeptObjectIncludeEmp *bool `json:"deptObjectIncludeEmp,omitempty" xml:"deptObjectIncludeEmp,omitempty"`
}

func (s AddContactHideBySceneSettingRequestSearchSceneConfig) String() string {
	return tea.Prettify(s)
}

func (s AddContactHideBySceneSettingRequestSearchSceneConfig) GoString() string {
	return s.String()
}

func (s *AddContactHideBySceneSettingRequestSearchSceneConfig) SetActive(v bool) *AddContactHideBySceneSettingRequestSearchSceneConfig {
	s.Active = &v
	return s
}

func (s *AddContactHideBySceneSettingRequestSearchSceneConfig) SetDeptObjectIncludeEmp(v bool) *AddContactHideBySceneSettingRequestSearchSceneConfig {
	s.DeptObjectIncludeEmp = &v
	return s
}

type AddContactHideBySceneSettingResponseBody struct {
	SettingId *int64 `json:"settingId,omitempty" xml:"settingId,omitempty"`
}

func (s AddContactHideBySceneSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddContactHideBySceneSettingResponseBody) GoString() string {
	return s.String()
}

func (s *AddContactHideBySceneSettingResponseBody) SetSettingId(v int64) *AddContactHideBySceneSettingResponseBody {
	s.SettingId = &v
	return s
}

type AddContactHideBySceneSettingResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddContactHideBySceneSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddContactHideBySceneSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s AddContactHideBySceneSettingResponse) GoString() string {
	return s.String()
}

func (s *AddContactHideBySceneSettingResponse) SetHeaders(v map[string]*string) *AddContactHideBySceneSettingResponse {
	s.Headers = v
	return s
}

func (s *AddContactHideBySceneSettingResponse) SetStatusCode(v int32) *AddContactHideBySceneSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *AddContactHideBySceneSettingResponse) SetBody(v *AddContactHideBySceneSettingResponseBody) *AddContactHideBySceneSettingResponse {
	s.Body = v
	return s
}

type AddEmpAttributeHideBySceneSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddEmpAttributeHideBySceneSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddEmpAttributeHideBySceneSettingHeaders) GoString() string {
	return s.String()
}

func (s *AddEmpAttributeHideBySceneSettingHeaders) SetCommonHeaders(v map[string]*string) *AddEmpAttributeHideBySceneSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddEmpAttributeHideBySceneSettingHeaders) SetXAcsDingtalkAccessToken(v string) *AddEmpAttributeHideBySceneSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddEmpAttributeHideBySceneSettingRequest struct {
	ChatSubtitleConfig *AddEmpAttributeHideBySceneSettingRequestChatSubtitleConfig `json:"chatSubtitleConfig,omitempty" xml:"chatSubtitleConfig,omitempty" type:"Struct"`
	Description        *string                                                     `json:"description,omitempty" xml:"description,omitempty"`
	ExcludeDeptIds     []*int64                                                    `json:"excludeDeptIds,omitempty" xml:"excludeDeptIds,omitempty" type:"Repeated"`
	ExcludeTagIds      []*int64                                                    `json:"excludeTagIds,omitempty" xml:"excludeTagIds,omitempty" type:"Repeated"`
	ExcludeUserIds     []*string                                                   `json:"excludeUserIds,omitempty" xml:"excludeUserIds,omitempty" type:"Repeated"`
	HideFields         []*string                                                   `json:"hideFields,omitempty" xml:"hideFields,omitempty" type:"Repeated"`
	Name               *string                                                     `json:"name,omitempty" xml:"name,omitempty"`
	ObjectDeptIds      []*int64                                                    `json:"objectDeptIds,omitempty" xml:"objectDeptIds,omitempty" type:"Repeated"`
	ObjectTagIds       []*int64                                                    `json:"objectTagIds,omitempty" xml:"objectTagIds,omitempty" type:"Repeated"`
	ObjectUserIds      []*string                                                   `json:"objectUserIds,omitempty" xml:"objectUserIds,omitempty" type:"Repeated"`
	ProfileSceneConfig *AddEmpAttributeHideBySceneSettingRequestProfileSceneConfig `json:"profileSceneConfig,omitempty" xml:"profileSceneConfig,omitempty" type:"Struct"`
	SearchSceneConfig  *AddEmpAttributeHideBySceneSettingRequestSearchSceneConfig  `json:"searchSceneConfig,omitempty" xml:"searchSceneConfig,omitempty" type:"Struct"`
}

func (s AddEmpAttributeHideBySceneSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s AddEmpAttributeHideBySceneSettingRequest) GoString() string {
	return s.String()
}

func (s *AddEmpAttributeHideBySceneSettingRequest) SetChatSubtitleConfig(v *AddEmpAttributeHideBySceneSettingRequestChatSubtitleConfig) *AddEmpAttributeHideBySceneSettingRequest {
	s.ChatSubtitleConfig = v
	return s
}

func (s *AddEmpAttributeHideBySceneSettingRequest) SetDescription(v string) *AddEmpAttributeHideBySceneSettingRequest {
	s.Description = &v
	return s
}

func (s *AddEmpAttributeHideBySceneSettingRequest) SetExcludeDeptIds(v []*int64) *AddEmpAttributeHideBySceneSettingRequest {
	s.ExcludeDeptIds = v
	return s
}

func (s *AddEmpAttributeHideBySceneSettingRequest) SetExcludeTagIds(v []*int64) *AddEmpAttributeHideBySceneSettingRequest {
	s.ExcludeTagIds = v
	return s
}

func (s *AddEmpAttributeHideBySceneSettingRequest) SetExcludeUserIds(v []*string) *AddEmpAttributeHideBySceneSettingRequest {
	s.ExcludeUserIds = v
	return s
}

func (s *AddEmpAttributeHideBySceneSettingRequest) SetHideFields(v []*string) *AddEmpAttributeHideBySceneSettingRequest {
	s.HideFields = v
	return s
}

func (s *AddEmpAttributeHideBySceneSettingRequest) SetName(v string) *AddEmpAttributeHideBySceneSettingRequest {
	s.Name = &v
	return s
}

func (s *AddEmpAttributeHideBySceneSettingRequest) SetObjectDeptIds(v []*int64) *AddEmpAttributeHideBySceneSettingRequest {
	s.ObjectDeptIds = v
	return s
}

func (s *AddEmpAttributeHideBySceneSettingRequest) SetObjectTagIds(v []*int64) *AddEmpAttributeHideBySceneSettingRequest {
	s.ObjectTagIds = v
	return s
}

func (s *AddEmpAttributeHideBySceneSettingRequest) SetObjectUserIds(v []*string) *AddEmpAttributeHideBySceneSettingRequest {
	s.ObjectUserIds = v
	return s
}

func (s *AddEmpAttributeHideBySceneSettingRequest) SetProfileSceneConfig(v *AddEmpAttributeHideBySceneSettingRequestProfileSceneConfig) *AddEmpAttributeHideBySceneSettingRequest {
	s.ProfileSceneConfig = v
	return s
}

func (s *AddEmpAttributeHideBySceneSettingRequest) SetSearchSceneConfig(v *AddEmpAttributeHideBySceneSettingRequestSearchSceneConfig) *AddEmpAttributeHideBySceneSettingRequest {
	s.SearchSceneConfig = v
	return s
}

type AddEmpAttributeHideBySceneSettingRequestChatSubtitleConfig struct {
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
}

func (s AddEmpAttributeHideBySceneSettingRequestChatSubtitleConfig) String() string {
	return tea.Prettify(s)
}

func (s AddEmpAttributeHideBySceneSettingRequestChatSubtitleConfig) GoString() string {
	return s.String()
}

func (s *AddEmpAttributeHideBySceneSettingRequestChatSubtitleConfig) SetActive(v bool) *AddEmpAttributeHideBySceneSettingRequestChatSubtitleConfig {
	s.Active = &v
	return s
}

type AddEmpAttributeHideBySceneSettingRequestProfileSceneConfig struct {
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
}

func (s AddEmpAttributeHideBySceneSettingRequestProfileSceneConfig) String() string {
	return tea.Prettify(s)
}

func (s AddEmpAttributeHideBySceneSettingRequestProfileSceneConfig) GoString() string {
	return s.String()
}

func (s *AddEmpAttributeHideBySceneSettingRequestProfileSceneConfig) SetActive(v bool) *AddEmpAttributeHideBySceneSettingRequestProfileSceneConfig {
	s.Active = &v
	return s
}

type AddEmpAttributeHideBySceneSettingRequestSearchSceneConfig struct {
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
}

func (s AddEmpAttributeHideBySceneSettingRequestSearchSceneConfig) String() string {
	return tea.Prettify(s)
}

func (s AddEmpAttributeHideBySceneSettingRequestSearchSceneConfig) GoString() string {
	return s.String()
}

func (s *AddEmpAttributeHideBySceneSettingRequestSearchSceneConfig) SetActive(v bool) *AddEmpAttributeHideBySceneSettingRequestSearchSceneConfig {
	s.Active = &v
	return s
}

type AddEmpAttributeHideBySceneSettingResponseBody struct {
	SettingId *int64 `json:"settingId,omitempty" xml:"settingId,omitempty"`
}

func (s AddEmpAttributeHideBySceneSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddEmpAttributeHideBySceneSettingResponseBody) GoString() string {
	return s.String()
}

func (s *AddEmpAttributeHideBySceneSettingResponseBody) SetSettingId(v int64) *AddEmpAttributeHideBySceneSettingResponseBody {
	s.SettingId = &v
	return s
}

type AddEmpAttributeHideBySceneSettingResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddEmpAttributeHideBySceneSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddEmpAttributeHideBySceneSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s AddEmpAttributeHideBySceneSettingResponse) GoString() string {
	return s.String()
}

func (s *AddEmpAttributeHideBySceneSettingResponse) SetHeaders(v map[string]*string) *AddEmpAttributeHideBySceneSettingResponse {
	s.Headers = v
	return s
}

func (s *AddEmpAttributeHideBySceneSettingResponse) SetStatusCode(v int32) *AddEmpAttributeHideBySceneSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *AddEmpAttributeHideBySceneSettingResponse) SetBody(v *AddEmpAttributeHideBySceneSettingResponseBody) *AddEmpAttributeHideBySceneSettingResponse {
	s.Body = v
	return s
}

type AnnualCertificationAuditHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AnnualCertificationAuditHeaders) String() string {
	return tea.Prettify(s)
}

func (s AnnualCertificationAuditHeaders) GoString() string {
	return s.String()
}

func (s *AnnualCertificationAuditHeaders) SetCommonHeaders(v map[string]*string) *AnnualCertificationAuditHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AnnualCertificationAuditHeaders) SetXAcsDingtalkAccessToken(v string) *AnnualCertificationAuditHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AnnualCertificationAuditRequest struct {
	ApplicantMobile   *string `json:"applicantMobile,omitempty" xml:"applicantMobile,omitempty"`
	ApplicantName     *string `json:"applicantName,omitempty" xml:"applicantName,omitempty"`
	ApplicationLetter *string `json:"applicationLetter,omitempty" xml:"applicationLetter,omitempty"`
	AuthStatus        *int32  `json:"authStatus,omitempty" xml:"authStatus,omitempty"`
	CertificateType   *int32  `json:"certificateType,omitempty" xml:"certificateType,omitempty"`
	CorpName          *string `json:"corpName,omitempty" xml:"corpName,omitempty"`
	DepositaryBank    *string `json:"depositaryBank,omitempty" xml:"depositaryBank,omitempty"`
	Extension         *string `json:"extension,omitempty" xml:"extension,omitempty"`
	LegalPerson       *string `json:"legalPerson,omitempty" xml:"legalPerson,omitempty"`
	LicenseNumber     *string `json:"licenseNumber,omitempty" xml:"licenseNumber,omitempty"`
	LicenseUrl        *string `json:"licenseUrl,omitempty" xml:"licenseUrl,omitempty"`
	OrderId           *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	PublicAccount     *string `json:"publicAccount,omitempty" xml:"publicAccount,omitempty"`
	ReasonCode        *string `json:"reasonCode,omitempty" xml:"reasonCode,omitempty"`
	ReasonMsg         *string `json:"reasonMsg,omitempty" xml:"reasonMsg,omitempty"`
	Tag               *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s AnnualCertificationAuditRequest) String() string {
	return tea.Prettify(s)
}

func (s AnnualCertificationAuditRequest) GoString() string {
	return s.String()
}

func (s *AnnualCertificationAuditRequest) SetApplicantMobile(v string) *AnnualCertificationAuditRequest {
	s.ApplicantMobile = &v
	return s
}

func (s *AnnualCertificationAuditRequest) SetApplicantName(v string) *AnnualCertificationAuditRequest {
	s.ApplicantName = &v
	return s
}

func (s *AnnualCertificationAuditRequest) SetApplicationLetter(v string) *AnnualCertificationAuditRequest {
	s.ApplicationLetter = &v
	return s
}

func (s *AnnualCertificationAuditRequest) SetAuthStatus(v int32) *AnnualCertificationAuditRequest {
	s.AuthStatus = &v
	return s
}

func (s *AnnualCertificationAuditRequest) SetCertificateType(v int32) *AnnualCertificationAuditRequest {
	s.CertificateType = &v
	return s
}

func (s *AnnualCertificationAuditRequest) SetCorpName(v string) *AnnualCertificationAuditRequest {
	s.CorpName = &v
	return s
}

func (s *AnnualCertificationAuditRequest) SetDepositaryBank(v string) *AnnualCertificationAuditRequest {
	s.DepositaryBank = &v
	return s
}

func (s *AnnualCertificationAuditRequest) SetExtension(v string) *AnnualCertificationAuditRequest {
	s.Extension = &v
	return s
}

func (s *AnnualCertificationAuditRequest) SetLegalPerson(v string) *AnnualCertificationAuditRequest {
	s.LegalPerson = &v
	return s
}

func (s *AnnualCertificationAuditRequest) SetLicenseNumber(v string) *AnnualCertificationAuditRequest {
	s.LicenseNumber = &v
	return s
}

func (s *AnnualCertificationAuditRequest) SetLicenseUrl(v string) *AnnualCertificationAuditRequest {
	s.LicenseUrl = &v
	return s
}

func (s *AnnualCertificationAuditRequest) SetOrderId(v string) *AnnualCertificationAuditRequest {
	s.OrderId = &v
	return s
}

func (s *AnnualCertificationAuditRequest) SetPublicAccount(v string) *AnnualCertificationAuditRequest {
	s.PublicAccount = &v
	return s
}

func (s *AnnualCertificationAuditRequest) SetReasonCode(v string) *AnnualCertificationAuditRequest {
	s.ReasonCode = &v
	return s
}

func (s *AnnualCertificationAuditRequest) SetReasonMsg(v string) *AnnualCertificationAuditRequest {
	s.ReasonMsg = &v
	return s
}

func (s *AnnualCertificationAuditRequest) SetTag(v string) *AnnualCertificationAuditRequest {
	s.Tag = &v
	return s
}

type AnnualCertificationAuditResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s AnnualCertificationAuditResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AnnualCertificationAuditResponseBody) GoString() string {
	return s.String()
}

func (s *AnnualCertificationAuditResponseBody) SetResult(v bool) *AnnualCertificationAuditResponseBody {
	s.Result = &v
	return s
}

type AnnualCertificationAuditResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AnnualCertificationAuditResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AnnualCertificationAuditResponse) String() string {
	return tea.Prettify(s)
}

func (s AnnualCertificationAuditResponse) GoString() string {
	return s.String()
}

func (s *AnnualCertificationAuditResponse) SetHeaders(v map[string]*string) *AnnualCertificationAuditResponse {
	s.Headers = v
	return s
}

func (s *AnnualCertificationAuditResponse) SetStatusCode(v int32) *AnnualCertificationAuditResponse {
	s.StatusCode = &v
	return s
}

func (s *AnnualCertificationAuditResponse) SetBody(v *AnnualCertificationAuditResponseBody) *AnnualCertificationAuditResponse {
	s.Body = v
	return s
}

type BatchApproveUnionApplyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchApproveUnionApplyHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchApproveUnionApplyHeaders) GoString() string {
	return s.String()
}

func (s *BatchApproveUnionApplyHeaders) SetCommonHeaders(v map[string]*string) *BatchApproveUnionApplyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchApproveUnionApplyHeaders) SetXAcsDingtalkAccessToken(v string) *BatchApproveUnionApplyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchApproveUnionApplyRequest struct {
	Body []*BatchApproveUnionApplyRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s BatchApproveUnionApplyRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchApproveUnionApplyRequest) GoString() string {
	return s.String()
}

func (s *BatchApproveUnionApplyRequest) SetBody(v []*BatchApproveUnionApplyRequestBody) *BatchApproveUnionApplyRequest {
	s.Body = v
	return s
}

type BatchApproveUnionApplyRequestBody struct {
	BranchCorpId  *string `json:"branchCorpId,omitempty" xml:"branchCorpId,omitempty"`
	LinkDeptId    *int64  `json:"linkDeptId,omitempty" xml:"linkDeptId,omitempty"`
	UnionRootName *string `json:"unionRootName,omitempty" xml:"unionRootName,omitempty"`
}

func (s BatchApproveUnionApplyRequestBody) String() string {
	return tea.Prettify(s)
}

func (s BatchApproveUnionApplyRequestBody) GoString() string {
	return s.String()
}

func (s *BatchApproveUnionApplyRequestBody) SetBranchCorpId(v string) *BatchApproveUnionApplyRequestBody {
	s.BranchCorpId = &v
	return s
}

func (s *BatchApproveUnionApplyRequestBody) SetLinkDeptId(v int64) *BatchApproveUnionApplyRequestBody {
	s.LinkDeptId = &v
	return s
}

func (s *BatchApproveUnionApplyRequestBody) SetUnionRootName(v string) *BatchApproveUnionApplyRequestBody {
	s.UnionRootName = &v
	return s
}

type BatchApproveUnionApplyResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s BatchApproveUnionApplyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchApproveUnionApplyResponseBody) GoString() string {
	return s.String()
}

func (s *BatchApproveUnionApplyResponseBody) SetResult(v bool) *BatchApproveUnionApplyResponseBody {
	s.Result = &v
	return s
}

type BatchApproveUnionApplyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchApproveUnionApplyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchApproveUnionApplyResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchApproveUnionApplyResponse) GoString() string {
	return s.String()
}

func (s *BatchApproveUnionApplyResponse) SetHeaders(v map[string]*string) *BatchApproveUnionApplyResponse {
	s.Headers = v
	return s
}

func (s *BatchApproveUnionApplyResponse) SetStatusCode(v int32) *BatchApproveUnionApplyResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchApproveUnionApplyResponse) SetBody(v *BatchApproveUnionApplyResponseBody) *BatchApproveUnionApplyResponse {
	s.Body = v
	return s
}

type ChangeMainAdminHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ChangeMainAdminHeaders) String() string {
	return tea.Prettify(s)
}

func (s ChangeMainAdminHeaders) GoString() string {
	return s.String()
}

func (s *ChangeMainAdminHeaders) SetCommonHeaders(v map[string]*string) *ChangeMainAdminHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ChangeMainAdminHeaders) SetXAcsDingtalkAccessToken(v string) *ChangeMainAdminHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ChangeMainAdminRequest struct {
	EffectCorpId *string `json:"effectCorpId,omitempty" xml:"effectCorpId,omitempty"`
	SourceUserId *string `json:"sourceUserId,omitempty" xml:"sourceUserId,omitempty"`
	TargetUserId *string `json:"targetUserId,omitempty" xml:"targetUserId,omitempty"`
}

func (s ChangeMainAdminRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeMainAdminRequest) GoString() string {
	return s.String()
}

func (s *ChangeMainAdminRequest) SetEffectCorpId(v string) *ChangeMainAdminRequest {
	s.EffectCorpId = &v
	return s
}

func (s *ChangeMainAdminRequest) SetSourceUserId(v string) *ChangeMainAdminRequest {
	s.SourceUserId = &v
	return s
}

func (s *ChangeMainAdminRequest) SetTargetUserId(v string) *ChangeMainAdminRequest {
	s.TargetUserId = &v
	return s
}

type ChangeMainAdminResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s ChangeMainAdminResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeMainAdminResponse) GoString() string {
	return s.String()
}

func (s *ChangeMainAdminResponse) SetHeaders(v map[string]*string) *ChangeMainAdminResponse {
	s.Headers = v
	return s
}

func (s *ChangeMainAdminResponse) SetStatusCode(v int32) *ChangeMainAdminResponse {
	s.StatusCode = &v
	return s
}

type CreateCooperateOrgHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateCooperateOrgHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateCooperateOrgHeaders) GoString() string {
	return s.String()
}

func (s *CreateCooperateOrgHeaders) SetCommonHeaders(v map[string]*string) *CreateCooperateOrgHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateCooperateOrgHeaders) SetXAcsDingtalkAccessToken(v string) *CreateCooperateOrgHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateCooperateOrgRequest struct {
	IndustryCode *int64  `json:"industryCode,omitempty" xml:"industryCode,omitempty"`
	LogoMediaId  *string `json:"logoMediaId,omitempty" xml:"logoMediaId,omitempty"`
	OrgName      *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
}

func (s CreateCooperateOrgRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCooperateOrgRequest) GoString() string {
	return s.String()
}

func (s *CreateCooperateOrgRequest) SetIndustryCode(v int64) *CreateCooperateOrgRequest {
	s.IndustryCode = &v
	return s
}

func (s *CreateCooperateOrgRequest) SetLogoMediaId(v string) *CreateCooperateOrgRequest {
	s.LogoMediaId = &v
	return s
}

func (s *CreateCooperateOrgRequest) SetOrgName(v string) *CreateCooperateOrgRequest {
	s.OrgName = &v
	return s
}

type CreateCooperateOrgResponseBody struct {
	CooperateCorpId *string `json:"cooperateCorpId,omitempty" xml:"cooperateCorpId,omitempty"`
}

func (s CreateCooperateOrgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCooperateOrgResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCooperateOrgResponseBody) SetCooperateCorpId(v string) *CreateCooperateOrgResponseBody {
	s.CooperateCorpId = &v
	return s
}

type CreateCooperateOrgResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCooperateOrgResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCooperateOrgResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCooperateOrgResponse) GoString() string {
	return s.String()
}

func (s *CreateCooperateOrgResponse) SetHeaders(v map[string]*string) *CreateCooperateOrgResponse {
	s.Headers = v
	return s
}

func (s *CreateCooperateOrgResponse) SetStatusCode(v int32) *CreateCooperateOrgResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCooperateOrgResponse) SetBody(v *CreateCooperateOrgResponseBody) *CreateCooperateOrgResponse {
	s.Body = v
	return s
}

type CreateManagementGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateManagementGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateManagementGroupHeaders) GoString() string {
	return s.String()
}

func (s *CreateManagementGroupHeaders) SetCommonHeaders(v map[string]*string) *CreateManagementGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateManagementGroupHeaders) SetXAcsDingtalkAccessToken(v string) *CreateManagementGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateManagementGroupRequest struct {
	GroupName   *string                                `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Members     []*CreateManagementGroupRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	ResourceIds []*string                              `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	Scope       *CreateManagementGroupRequestScope     `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
}

func (s CreateManagementGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateManagementGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateManagementGroupRequest) SetGroupName(v string) *CreateManagementGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateManagementGroupRequest) SetMembers(v []*CreateManagementGroupRequestMembers) *CreateManagementGroupRequest {
	s.Members = v
	return s
}

func (s *CreateManagementGroupRequest) SetResourceIds(v []*string) *CreateManagementGroupRequest {
	s.ResourceIds = v
	return s
}

func (s *CreateManagementGroupRequest) SetScope(v *CreateManagementGroupRequestScope) *CreateManagementGroupRequest {
	s.Scope = v
	return s
}

type CreateManagementGroupRequestMembers struct {
	MemberId   *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
}

func (s CreateManagementGroupRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s CreateManagementGroupRequestMembers) GoString() string {
	return s.String()
}

func (s *CreateManagementGroupRequestMembers) SetMemberId(v string) *CreateManagementGroupRequestMembers {
	s.MemberId = &v
	return s
}

func (s *CreateManagementGroupRequestMembers) SetMemberType(v string) *CreateManagementGroupRequestMembers {
	s.MemberType = &v
	return s
}

type CreateManagementGroupRequestScope struct {
	DeptIds   []*int64 `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	ScopeType *int32   `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
}

func (s CreateManagementGroupRequestScope) String() string {
	return tea.Prettify(s)
}

func (s CreateManagementGroupRequestScope) GoString() string {
	return s.String()
}

func (s *CreateManagementGroupRequestScope) SetDeptIds(v []*int64) *CreateManagementGroupRequestScope {
	s.DeptIds = v
	return s
}

func (s *CreateManagementGroupRequestScope) SetScopeType(v int32) *CreateManagementGroupRequestScope {
	s.ScopeType = &v
	return s
}

type CreateManagementGroupResponseBody struct {
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
}

func (s CreateManagementGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateManagementGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateManagementGroupResponseBody) SetGroupId(v string) *CreateManagementGroupResponseBody {
	s.GroupId = &v
	return s
}

type CreateManagementGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateManagementGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateManagementGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateManagementGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateManagementGroupResponse) SetHeaders(v map[string]*string) *CreateManagementGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateManagementGroupResponse) SetStatusCode(v int32) *CreateManagementGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateManagementGroupResponse) SetBody(v *CreateManagementGroupResponseBody) *CreateManagementGroupResponse {
	s.Body = v
	return s
}

type CreateSecondaryManagementGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateSecondaryManagementGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateSecondaryManagementGroupHeaders) GoString() string {
	return s.String()
}

func (s *CreateSecondaryManagementGroupHeaders) SetCommonHeaders(v map[string]*string) *CreateSecondaryManagementGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSecondaryManagementGroupHeaders) SetXAcsDingtalkAccessToken(v string) *CreateSecondaryManagementGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateSecondaryManagementGroupRequest struct {
	GroupName   *string                                         `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Members     []*CreateSecondaryManagementGroupRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	ResourceIds []*string                                       `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	Scope       *CreateSecondaryManagementGroupRequestScope     `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
	UserId      *string                                         `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateSecondaryManagementGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSecondaryManagementGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateSecondaryManagementGroupRequest) SetGroupName(v string) *CreateSecondaryManagementGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateSecondaryManagementGroupRequest) SetMembers(v []*CreateSecondaryManagementGroupRequestMembers) *CreateSecondaryManagementGroupRequest {
	s.Members = v
	return s
}

func (s *CreateSecondaryManagementGroupRequest) SetResourceIds(v []*string) *CreateSecondaryManagementGroupRequest {
	s.ResourceIds = v
	return s
}

func (s *CreateSecondaryManagementGroupRequest) SetScope(v *CreateSecondaryManagementGroupRequestScope) *CreateSecondaryManagementGroupRequest {
	s.Scope = v
	return s
}

func (s *CreateSecondaryManagementGroupRequest) SetUserId(v string) *CreateSecondaryManagementGroupRequest {
	s.UserId = &v
	return s
}

type CreateSecondaryManagementGroupRequestMembers struct {
	MemberId   *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
}

func (s CreateSecondaryManagementGroupRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s CreateSecondaryManagementGroupRequestMembers) GoString() string {
	return s.String()
}

func (s *CreateSecondaryManagementGroupRequestMembers) SetMemberId(v string) *CreateSecondaryManagementGroupRequestMembers {
	s.MemberId = &v
	return s
}

func (s *CreateSecondaryManagementGroupRequestMembers) SetMemberType(v string) *CreateSecondaryManagementGroupRequestMembers {
	s.MemberType = &v
	return s
}

type CreateSecondaryManagementGroupRequestScope struct {
	DeptIds   []*int64 `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	ScopeType *int32   `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
}

func (s CreateSecondaryManagementGroupRequestScope) String() string {
	return tea.Prettify(s)
}

func (s CreateSecondaryManagementGroupRequestScope) GoString() string {
	return s.String()
}

func (s *CreateSecondaryManagementGroupRequestScope) SetDeptIds(v []*int64) *CreateSecondaryManagementGroupRequestScope {
	s.DeptIds = v
	return s
}

func (s *CreateSecondaryManagementGroupRequestScope) SetScopeType(v int32) *CreateSecondaryManagementGroupRequestScope {
	s.ScopeType = &v
	return s
}

type CreateSecondaryManagementGroupResponseBody struct {
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
}

func (s CreateSecondaryManagementGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSecondaryManagementGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecondaryManagementGroupResponseBody) SetGroupId(v string) *CreateSecondaryManagementGroupResponseBody {
	s.GroupId = &v
	return s
}

type CreateSecondaryManagementGroupResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSecondaryManagementGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSecondaryManagementGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSecondaryManagementGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateSecondaryManagementGroupResponse) SetHeaders(v map[string]*string) *CreateSecondaryManagementGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateSecondaryManagementGroupResponse) SetStatusCode(v int32) *CreateSecondaryManagementGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSecondaryManagementGroupResponse) SetBody(v *CreateSecondaryManagementGroupResponseBody) *CreateSecondaryManagementGroupResponse {
	s.Body = v
	return s
}

type DeleteContactHideBySceneSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteContactHideBySceneSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactHideBySceneSettingHeaders) GoString() string {
	return s.String()
}

func (s *DeleteContactHideBySceneSettingHeaders) SetCommonHeaders(v map[string]*string) *DeleteContactHideBySceneSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteContactHideBySceneSettingHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteContactHideBySceneSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteContactHideBySceneSettingResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteContactHideBySceneSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactHideBySceneSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContactHideBySceneSettingResponseBody) SetSuccess(v bool) *DeleteContactHideBySceneSettingResponseBody {
	s.Success = &v
	return s
}

type DeleteContactHideBySceneSettingResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteContactHideBySceneSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteContactHideBySceneSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactHideBySceneSettingResponse) GoString() string {
	return s.String()
}

func (s *DeleteContactHideBySceneSettingResponse) SetHeaders(v map[string]*string) *DeleteContactHideBySceneSettingResponse {
	s.Headers = v
	return s
}

func (s *DeleteContactHideBySceneSettingResponse) SetStatusCode(v int32) *DeleteContactHideBySceneSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContactHideBySceneSettingResponse) SetBody(v *DeleteContactHideBySceneSettingResponseBody) *DeleteContactHideBySceneSettingResponse {
	s.Body = v
	return s
}

type DeleteContactHideSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteContactHideSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactHideSettingHeaders) GoString() string {
	return s.String()
}

func (s *DeleteContactHideSettingHeaders) SetCommonHeaders(v map[string]*string) *DeleteContactHideSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteContactHideSettingHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteContactHideSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteContactHideSettingResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteContactHideSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactHideSettingResponse) GoString() string {
	return s.String()
}

func (s *DeleteContactHideSettingResponse) SetHeaders(v map[string]*string) *DeleteContactHideSettingResponse {
	s.Headers = v
	return s
}

func (s *DeleteContactHideSettingResponse) SetStatusCode(v int32) *DeleteContactHideSettingResponse {
	s.StatusCode = &v
	return s
}

type DeleteContactRestrictSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteContactRestrictSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactRestrictSettingHeaders) GoString() string {
	return s.String()
}

func (s *DeleteContactRestrictSettingHeaders) SetCommonHeaders(v map[string]*string) *DeleteContactRestrictSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteContactRestrictSettingHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteContactRestrictSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteContactRestrictSettingResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteContactRestrictSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactRestrictSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContactRestrictSettingResponseBody) SetResult(v bool) *DeleteContactRestrictSettingResponseBody {
	s.Result = &v
	return s
}

type DeleteContactRestrictSettingResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteContactRestrictSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteContactRestrictSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactRestrictSettingResponse) GoString() string {
	return s.String()
}

func (s *DeleteContactRestrictSettingResponse) SetHeaders(v map[string]*string) *DeleteContactRestrictSettingResponse {
	s.Headers = v
	return s
}

func (s *DeleteContactRestrictSettingResponse) SetStatusCode(v int32) *DeleteContactRestrictSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContactRestrictSettingResponse) SetBody(v *DeleteContactRestrictSettingResponseBody) *DeleteContactRestrictSettingResponse {
	s.Body = v
	return s
}

type DeleteEmpAttributeHideBySceneSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteEmpAttributeHideBySceneSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteEmpAttributeHideBySceneSettingHeaders) GoString() string {
	return s.String()
}

func (s *DeleteEmpAttributeHideBySceneSettingHeaders) SetCommonHeaders(v map[string]*string) *DeleteEmpAttributeHideBySceneSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteEmpAttributeHideBySceneSettingHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteEmpAttributeHideBySceneSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteEmpAttributeHideBySceneSettingResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteEmpAttributeHideBySceneSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEmpAttributeHideBySceneSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEmpAttributeHideBySceneSettingResponseBody) SetSuccess(v bool) *DeleteEmpAttributeHideBySceneSettingResponseBody {
	s.Success = &v
	return s
}

type DeleteEmpAttributeHideBySceneSettingResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteEmpAttributeHideBySceneSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEmpAttributeHideBySceneSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEmpAttributeHideBySceneSettingResponse) GoString() string {
	return s.String()
}

func (s *DeleteEmpAttributeHideBySceneSettingResponse) SetHeaders(v map[string]*string) *DeleteEmpAttributeHideBySceneSettingResponse {
	s.Headers = v
	return s
}

func (s *DeleteEmpAttributeHideBySceneSettingResponse) SetStatusCode(v int32) *DeleteEmpAttributeHideBySceneSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEmpAttributeHideBySceneSettingResponse) SetBody(v *DeleteEmpAttributeHideBySceneSettingResponseBody) *DeleteEmpAttributeHideBySceneSettingResponse {
	s.Body = v
	return s
}

type DeleteEmpAttributeVisibilityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteEmpAttributeVisibilityHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteEmpAttributeVisibilityHeaders) GoString() string {
	return s.String()
}

func (s *DeleteEmpAttributeVisibilityHeaders) SetCommonHeaders(v map[string]*string) *DeleteEmpAttributeVisibilityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteEmpAttributeVisibilityHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteEmpAttributeVisibilityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteEmpAttributeVisibilityResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteEmpAttributeVisibilityResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEmpAttributeVisibilityResponse) GoString() string {
	return s.String()
}

func (s *DeleteEmpAttributeVisibilityResponse) SetHeaders(v map[string]*string) *DeleteEmpAttributeVisibilityResponse {
	s.Headers = v
	return s
}

func (s *DeleteEmpAttributeVisibilityResponse) SetStatusCode(v int32) *DeleteEmpAttributeVisibilityResponse {
	s.StatusCode = &v
	return s
}

type DeleteManagementGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteManagementGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteManagementGroupHeaders) GoString() string {
	return s.String()
}

func (s *DeleteManagementGroupHeaders) SetCommonHeaders(v map[string]*string) *DeleteManagementGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteManagementGroupHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteManagementGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteManagementGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteManagementGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteManagementGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteManagementGroupResponse) SetHeaders(v map[string]*string) *DeleteManagementGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteManagementGroupResponse) SetStatusCode(v int32) *DeleteManagementGroupResponse {
	s.StatusCode = &v
	return s
}

type GetApplyInviteInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetApplyInviteInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetApplyInviteInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetApplyInviteInfoHeaders) SetCommonHeaders(v map[string]*string) *GetApplyInviteInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetApplyInviteInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetApplyInviteInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetApplyInviteInfoRequest struct {
	DeptId        *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	InviterUserId *string `json:"inviterUserId,omitempty" xml:"inviterUserId,omitempty"`
}

func (s GetApplyInviteInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApplyInviteInfoRequest) GoString() string {
	return s.String()
}

func (s *GetApplyInviteInfoRequest) SetDeptId(v int64) *GetApplyInviteInfoRequest {
	s.DeptId = &v
	return s
}

func (s *GetApplyInviteInfoRequest) SetInviterUserId(v string) *GetApplyInviteInfoRequest {
	s.InviterUserId = &v
	return s
}

type GetApplyInviteInfoResponseBody struct {
	AuditType          *int64  `json:"auditType,omitempty" xml:"auditType,omitempty"`
	EmpApplyJoinDept   *bool   `json:"empApplyJoinDept,omitempty" xml:"empApplyJoinDept,omitempty"`
	InviteSwitch       *bool   `json:"inviteSwitch,omitempty" xml:"inviteSwitch,omitempty"`
	InviteUrl          *string `json:"inviteUrl,omitempty" xml:"inviteUrl,omitempty"`
	LinkInvite         *bool   `json:"linkInvite,omitempty" xml:"linkInvite,omitempty"`
	OrgApplyCodeInvite *bool   `json:"orgApplyCodeInvite,omitempty" xml:"orgApplyCodeInvite,omitempty"`
	SearchNameInvite   *bool   `json:"searchNameInvite,omitempty" xml:"searchNameInvite,omitempty"`
}

func (s GetApplyInviteInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApplyInviteInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplyInviteInfoResponseBody) SetAuditType(v int64) *GetApplyInviteInfoResponseBody {
	s.AuditType = &v
	return s
}

func (s *GetApplyInviteInfoResponseBody) SetEmpApplyJoinDept(v bool) *GetApplyInviteInfoResponseBody {
	s.EmpApplyJoinDept = &v
	return s
}

func (s *GetApplyInviteInfoResponseBody) SetInviteSwitch(v bool) *GetApplyInviteInfoResponseBody {
	s.InviteSwitch = &v
	return s
}

func (s *GetApplyInviteInfoResponseBody) SetInviteUrl(v string) *GetApplyInviteInfoResponseBody {
	s.InviteUrl = &v
	return s
}

func (s *GetApplyInviteInfoResponseBody) SetLinkInvite(v bool) *GetApplyInviteInfoResponseBody {
	s.LinkInvite = &v
	return s
}

func (s *GetApplyInviteInfoResponseBody) SetOrgApplyCodeInvite(v bool) *GetApplyInviteInfoResponseBody {
	s.OrgApplyCodeInvite = &v
	return s
}

func (s *GetApplyInviteInfoResponseBody) SetSearchNameInvite(v bool) *GetApplyInviteInfoResponseBody {
	s.SearchNameInvite = &v
	return s
}

type GetApplyInviteInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetApplyInviteInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetApplyInviteInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApplyInviteInfoResponse) GoString() string {
	return s.String()
}

func (s *GetApplyInviteInfoResponse) SetHeaders(v map[string]*string) *GetApplyInviteInfoResponse {
	s.Headers = v
	return s
}

func (s *GetApplyInviteInfoResponse) SetStatusCode(v int32) *GetApplyInviteInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplyInviteInfoResponse) SetBody(v *GetApplyInviteInfoResponseBody) *GetApplyInviteInfoResponse {
	s.Body = v
	return s
}

type GetBranchAuthDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetBranchAuthDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetBranchAuthDataHeaders) GoString() string {
	return s.String()
}

func (s *GetBranchAuthDataHeaders) SetCommonHeaders(v map[string]*string) *GetBranchAuthDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetBranchAuthDataHeaders) SetXAcsDingtalkAccessToken(v string) *GetBranchAuthDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetBranchAuthDataRequest struct {
	Body         map[string]*string `json:"body,omitempty" xml:"body,omitempty"`
	BranchCorpId *string            `json:"branchCorpId,omitempty" xml:"branchCorpId,omitempty"`
	Code         *string            `json:"code,omitempty" xml:"code,omitempty"`
}

func (s GetBranchAuthDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBranchAuthDataRequest) GoString() string {
	return s.String()
}

func (s *GetBranchAuthDataRequest) SetBody(v map[string]*string) *GetBranchAuthDataRequest {
	s.Body = v
	return s
}

func (s *GetBranchAuthDataRequest) SetBranchCorpId(v string) *GetBranchAuthDataRequest {
	s.BranchCorpId = &v
	return s
}

func (s *GetBranchAuthDataRequest) SetCode(v string) *GetBranchAuthDataRequest {
	s.Code = &v
	return s
}

type GetBranchAuthDataResponseBody struct {
	Result []*GetBranchAuthDataResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetBranchAuthDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBranchAuthDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetBranchAuthDataResponseBody) SetResult(v []*GetBranchAuthDataResponseBodyResult) *GetBranchAuthDataResponseBody {
	s.Result = v
	return s
}

type GetBranchAuthDataResponseBodyResult struct {
	FieldCode  *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	FieldName  *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
}

func (s GetBranchAuthDataResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetBranchAuthDataResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetBranchAuthDataResponseBodyResult) SetFieldCode(v string) *GetBranchAuthDataResponseBodyResult {
	s.FieldCode = &v
	return s
}

func (s *GetBranchAuthDataResponseBodyResult) SetFieldName(v string) *GetBranchAuthDataResponseBodyResult {
	s.FieldName = &v
	return s
}

func (s *GetBranchAuthDataResponseBodyResult) SetFieldValue(v string) *GetBranchAuthDataResponseBodyResult {
	s.FieldValue = &v
	return s
}

type GetBranchAuthDataResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBranchAuthDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBranchAuthDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBranchAuthDataResponse) GoString() string {
	return s.String()
}

func (s *GetBranchAuthDataResponse) SetHeaders(v map[string]*string) *GetBranchAuthDataResponse {
	s.Headers = v
	return s
}

func (s *GetBranchAuthDataResponse) SetStatusCode(v int32) *GetBranchAuthDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBranchAuthDataResponse) SetBody(v *GetBranchAuthDataResponseBody) *GetBranchAuthDataResponse {
	s.Body = v
	return s
}

type GetCardInUserHolderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCardInUserHolderHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCardInUserHolderHeaders) GoString() string {
	return s.String()
}

func (s *GetCardInUserHolderHeaders) SetCommonHeaders(v map[string]*string) *GetCardInUserHolderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCardInUserHolderHeaders) SetXAcsDingtalkAccessToken(v string) *GetCardInUserHolderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCardInUserHolderResponseBody struct {
	AvatarUrl          *string                `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	CardAcceptStatus   *int32                 `json:"cardAcceptStatus,omitempty" xml:"cardAcceptStatus,omitempty"`
	CardAcceptTimeLong *int64                 `json:"cardAcceptTimeLong,omitempty" xml:"cardAcceptTimeLong,omitempty"`
	CardId             *string                `json:"cardId,omitempty" xml:"cardId,omitempty"`
	CardSource         *int32                 `json:"cardSource,omitempty" xml:"cardSource,omitempty"`
	Extension          map[string]interface{} `json:"extension,omitempty" xml:"extension,omitempty"`
	IndustryName       *string                `json:"industryName,omitempty" xml:"industryName,omitempty"`
	Introduce          *string                `json:"introduce,omitempty" xml:"introduce,omitempty"`
	Name               *string                `json:"name,omitempty" xml:"name,omitempty"`
	OrgName            *string                `json:"orgName,omitempty" xml:"orgName,omitempty"`
	TemplateId         *string                `json:"templateId,omitempty" xml:"templateId,omitempty"`
	Title              *string                `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetCardInUserHolderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCardInUserHolderResponseBody) GoString() string {
	return s.String()
}

func (s *GetCardInUserHolderResponseBody) SetAvatarUrl(v string) *GetCardInUserHolderResponseBody {
	s.AvatarUrl = &v
	return s
}

func (s *GetCardInUserHolderResponseBody) SetCardAcceptStatus(v int32) *GetCardInUserHolderResponseBody {
	s.CardAcceptStatus = &v
	return s
}

func (s *GetCardInUserHolderResponseBody) SetCardAcceptTimeLong(v int64) *GetCardInUserHolderResponseBody {
	s.CardAcceptTimeLong = &v
	return s
}

func (s *GetCardInUserHolderResponseBody) SetCardId(v string) *GetCardInUserHolderResponseBody {
	s.CardId = &v
	return s
}

func (s *GetCardInUserHolderResponseBody) SetCardSource(v int32) *GetCardInUserHolderResponseBody {
	s.CardSource = &v
	return s
}

func (s *GetCardInUserHolderResponseBody) SetExtension(v map[string]interface{}) *GetCardInUserHolderResponseBody {
	s.Extension = v
	return s
}

func (s *GetCardInUserHolderResponseBody) SetIndustryName(v string) *GetCardInUserHolderResponseBody {
	s.IndustryName = &v
	return s
}

func (s *GetCardInUserHolderResponseBody) SetIntroduce(v string) *GetCardInUserHolderResponseBody {
	s.Introduce = &v
	return s
}

func (s *GetCardInUserHolderResponseBody) SetName(v string) *GetCardInUserHolderResponseBody {
	s.Name = &v
	return s
}

func (s *GetCardInUserHolderResponseBody) SetOrgName(v string) *GetCardInUserHolderResponseBody {
	s.OrgName = &v
	return s
}

func (s *GetCardInUserHolderResponseBody) SetTemplateId(v string) *GetCardInUserHolderResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetCardInUserHolderResponseBody) SetTitle(v string) *GetCardInUserHolderResponseBody {
	s.Title = &v
	return s
}

type GetCardInUserHolderResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCardInUserHolderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCardInUserHolderResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCardInUserHolderResponse) GoString() string {
	return s.String()
}

func (s *GetCardInUserHolderResponse) SetHeaders(v map[string]*string) *GetCardInUserHolderResponse {
	s.Headers = v
	return s
}

func (s *GetCardInUserHolderResponse) SetStatusCode(v int32) *GetCardInUserHolderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCardInUserHolderResponse) SetBody(v *GetCardInUserHolderResponseBody) *GetCardInUserHolderResponse {
	s.Body = v
	return s
}

type GetCardInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCardInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCardInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetCardInfoHeaders) SetCommonHeaders(v map[string]*string) *GetCardInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCardInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetCardInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCardInfoResponseBody struct {
	AdminRole    *int64                            `json:"adminRole,omitempty" xml:"adminRole,omitempty"`
	AvatarUrl    *string                           `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	CardId       *string                           `json:"cardId,omitempty" xml:"cardId,omitempty"`
	Extension    *GetCardInfoResponseBodyExtension `json:"extension,omitempty" xml:"extension,omitempty" type:"Struct"`
	IndustryName *string                           `json:"industryName,omitempty" xml:"industryName,omitempty"`
	Introduce    map[string]interface{}            `json:"introduce,omitempty" xml:"introduce,omitempty"`
	Name         *string                           `json:"name,omitempty" xml:"name,omitempty"`
	OrgName      *string                           `json:"orgName,omitempty" xml:"orgName,omitempty"`
	Settings     map[string]interface{}            `json:"settings,omitempty" xml:"settings,omitempty"`
	TemplateId   *string                           `json:"templateId,omitempty" xml:"templateId,omitempty"`
	Title        *string                           `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetCardInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCardInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetCardInfoResponseBody) SetAdminRole(v int64) *GetCardInfoResponseBody {
	s.AdminRole = &v
	return s
}

func (s *GetCardInfoResponseBody) SetAvatarUrl(v string) *GetCardInfoResponseBody {
	s.AvatarUrl = &v
	return s
}

func (s *GetCardInfoResponseBody) SetCardId(v string) *GetCardInfoResponseBody {
	s.CardId = &v
	return s
}

func (s *GetCardInfoResponseBody) SetExtension(v *GetCardInfoResponseBodyExtension) *GetCardInfoResponseBody {
	s.Extension = v
	return s
}

func (s *GetCardInfoResponseBody) SetIndustryName(v string) *GetCardInfoResponseBody {
	s.IndustryName = &v
	return s
}

func (s *GetCardInfoResponseBody) SetIntroduce(v map[string]interface{}) *GetCardInfoResponseBody {
	s.Introduce = v
	return s
}

func (s *GetCardInfoResponseBody) SetName(v string) *GetCardInfoResponseBody {
	s.Name = &v
	return s
}

func (s *GetCardInfoResponseBody) SetOrgName(v string) *GetCardInfoResponseBody {
	s.OrgName = &v
	return s
}

func (s *GetCardInfoResponseBody) SetSettings(v map[string]interface{}) *GetCardInfoResponseBody {
	s.Settings = v
	return s
}

func (s *GetCardInfoResponseBody) SetTemplateId(v string) *GetCardInfoResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetCardInfoResponseBody) SetTitle(v string) *GetCardInfoResponseBody {
	s.Title = &v
	return s
}

type GetCardInfoResponseBodyExtension struct {
	CardContactInfo *GetCardInfoResponseBodyExtensionCardContactInfo `json:"cardContactInfo,omitempty" xml:"cardContactInfo,omitempty" type:"Struct"`
	CorpId          *string                                          `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Department      *string                                          `json:"department,omitempty" xml:"department,omitempty"`
	OrgAuthed       *bool                                            `json:"orgAuthed,omitempty" xml:"orgAuthed,omitempty"`
	OrgLogo         *string                                          `json:"orgLogo,omitempty" xml:"orgLogo,omitempty"`
	OriginCardUrl   *string                                          `json:"originCardUrl,omitempty" xml:"originCardUrl,omitempty"`
	ShareContent    *string                                          `json:"shareContent,omitempty" xml:"shareContent,omitempty"`
	ThumbnailUrl    *string                                          `json:"thumbnailUrl,omitempty" xml:"thumbnailUrl,omitempty"`
	VideoFileName   *string                                          `json:"videoFileName,omitempty" xml:"videoFileName,omitempty"`
	VideoTitle      *string                                          `json:"videoTitle,omitempty" xml:"videoTitle,omitempty"`
	VideoUrl        *string                                          `json:"videoUrl,omitempty" xml:"videoUrl,omitempty"`
}

func (s GetCardInfoResponseBodyExtension) String() string {
	return tea.Prettify(s)
}

func (s GetCardInfoResponseBodyExtension) GoString() string {
	return s.String()
}

func (s *GetCardInfoResponseBodyExtension) SetCardContactInfo(v *GetCardInfoResponseBodyExtensionCardContactInfo) *GetCardInfoResponseBodyExtension {
	s.CardContactInfo = v
	return s
}

func (s *GetCardInfoResponseBodyExtension) SetCorpId(v string) *GetCardInfoResponseBodyExtension {
	s.CorpId = &v
	return s
}

func (s *GetCardInfoResponseBodyExtension) SetDepartment(v string) *GetCardInfoResponseBodyExtension {
	s.Department = &v
	return s
}

func (s *GetCardInfoResponseBodyExtension) SetOrgAuthed(v bool) *GetCardInfoResponseBodyExtension {
	s.OrgAuthed = &v
	return s
}

func (s *GetCardInfoResponseBodyExtension) SetOrgLogo(v string) *GetCardInfoResponseBodyExtension {
	s.OrgLogo = &v
	return s
}

func (s *GetCardInfoResponseBodyExtension) SetOriginCardUrl(v string) *GetCardInfoResponseBodyExtension {
	s.OriginCardUrl = &v
	return s
}

func (s *GetCardInfoResponseBodyExtension) SetShareContent(v string) *GetCardInfoResponseBodyExtension {
	s.ShareContent = &v
	return s
}

func (s *GetCardInfoResponseBodyExtension) SetThumbnailUrl(v string) *GetCardInfoResponseBodyExtension {
	s.ThumbnailUrl = &v
	return s
}

func (s *GetCardInfoResponseBodyExtension) SetVideoFileName(v string) *GetCardInfoResponseBodyExtension {
	s.VideoFileName = &v
	return s
}

func (s *GetCardInfoResponseBodyExtension) SetVideoTitle(v string) *GetCardInfoResponseBodyExtension {
	s.VideoTitle = &v
	return s
}

func (s *GetCardInfoResponseBodyExtension) SetVideoUrl(v string) *GetCardInfoResponseBodyExtension {
	s.VideoUrl = &v
	return s
}

type GetCardInfoResponseBodyExtensionCardContactInfo struct {
	Address   []*GetCardInfoResponseBodyExtensionCardContactInfoAddress   `json:"address,omitempty" xml:"address,omitempty" type:"Repeated"`
	Email     []*GetCardInfoResponseBodyExtensionCardContactInfoEmail     `json:"email,omitempty" xml:"email,omitempty" type:"Repeated"`
	Link      []*GetCardInfoResponseBodyExtensionCardContactInfoLink      `json:"link,omitempty" xml:"link,omitempty" type:"Repeated"`
	Telephone []*GetCardInfoResponseBodyExtensionCardContactInfoTelephone `json:"telephone,omitempty" xml:"telephone,omitempty" type:"Repeated"`
	WorkPhone []*GetCardInfoResponseBodyExtensionCardContactInfoWorkPhone `json:"workPhone,omitempty" xml:"workPhone,omitempty" type:"Repeated"`
}

func (s GetCardInfoResponseBodyExtensionCardContactInfo) String() string {
	return tea.Prettify(s)
}

func (s GetCardInfoResponseBodyExtensionCardContactInfo) GoString() string {
	return s.String()
}

func (s *GetCardInfoResponseBodyExtensionCardContactInfo) SetAddress(v []*GetCardInfoResponseBodyExtensionCardContactInfoAddress) *GetCardInfoResponseBodyExtensionCardContactInfo {
	s.Address = v
	return s
}

func (s *GetCardInfoResponseBodyExtensionCardContactInfo) SetEmail(v []*GetCardInfoResponseBodyExtensionCardContactInfoEmail) *GetCardInfoResponseBodyExtensionCardContactInfo {
	s.Email = v
	return s
}

func (s *GetCardInfoResponseBodyExtensionCardContactInfo) SetLink(v []*GetCardInfoResponseBodyExtensionCardContactInfoLink) *GetCardInfoResponseBodyExtensionCardContactInfo {
	s.Link = v
	return s
}

func (s *GetCardInfoResponseBodyExtensionCardContactInfo) SetTelephone(v []*GetCardInfoResponseBodyExtensionCardContactInfoTelephone) *GetCardInfoResponseBodyExtensionCardContactInfo {
	s.Telephone = v
	return s
}

func (s *GetCardInfoResponseBodyExtensionCardContactInfo) SetWorkPhone(v []*GetCardInfoResponseBodyExtensionCardContactInfoWorkPhone) *GetCardInfoResponseBodyExtensionCardContactInfo {
	s.WorkPhone = v
	return s
}

type GetCardInfoResponseBodyExtensionCardContactInfoAddress struct {
	Area   *GetCardInfoResponseBodyExtensionCardContactInfoAddressArea `json:"area,omitempty" xml:"area,omitempty" type:"Struct"`
	Detail *string                                                     `json:"detail,omitempty" xml:"detail,omitempty"`
}

func (s GetCardInfoResponseBodyExtensionCardContactInfoAddress) String() string {
	return tea.Prettify(s)
}

func (s GetCardInfoResponseBodyExtensionCardContactInfoAddress) GoString() string {
	return s.String()
}

func (s *GetCardInfoResponseBodyExtensionCardContactInfoAddress) SetArea(v *GetCardInfoResponseBodyExtensionCardContactInfoAddressArea) *GetCardInfoResponseBodyExtensionCardContactInfoAddress {
	s.Area = v
	return s
}

func (s *GetCardInfoResponseBodyExtensionCardContactInfoAddress) SetDetail(v string) *GetCardInfoResponseBodyExtensionCardContactInfoAddress {
	s.Detail = &v
	return s
}

type GetCardInfoResponseBodyExtensionCardContactInfoAddressArea struct {
	Region         *string `json:"region,omitempty" xml:"region,omitempty"`
	RegionFullName *string `json:"regionFullName,omitempty" xml:"regionFullName,omitempty"`
}

func (s GetCardInfoResponseBodyExtensionCardContactInfoAddressArea) String() string {
	return tea.Prettify(s)
}

func (s GetCardInfoResponseBodyExtensionCardContactInfoAddressArea) GoString() string {
	return s.String()
}

func (s *GetCardInfoResponseBodyExtensionCardContactInfoAddressArea) SetRegion(v string) *GetCardInfoResponseBodyExtensionCardContactInfoAddressArea {
	s.Region = &v
	return s
}

func (s *GetCardInfoResponseBodyExtensionCardContactInfoAddressArea) SetRegionFullName(v string) *GetCardInfoResponseBodyExtensionCardContactInfoAddressArea {
	s.RegionFullName = &v
	return s
}

type GetCardInfoResponseBodyExtensionCardContactInfoEmail struct {
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetCardInfoResponseBodyExtensionCardContactInfoEmail) String() string {
	return tea.Prettify(s)
}

func (s GetCardInfoResponseBodyExtensionCardContactInfoEmail) GoString() string {
	return s.String()
}

func (s *GetCardInfoResponseBodyExtensionCardContactInfoEmail) SetLabel(v string) *GetCardInfoResponseBodyExtensionCardContactInfoEmail {
	s.Label = &v
	return s
}

func (s *GetCardInfoResponseBodyExtensionCardContactInfoEmail) SetValue(v string) *GetCardInfoResponseBodyExtensionCardContactInfoEmail {
	s.Value = &v
	return s
}

type GetCardInfoResponseBodyExtensionCardContactInfoLink struct {
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetCardInfoResponseBodyExtensionCardContactInfoLink) String() string {
	return tea.Prettify(s)
}

func (s GetCardInfoResponseBodyExtensionCardContactInfoLink) GoString() string {
	return s.String()
}

func (s *GetCardInfoResponseBodyExtensionCardContactInfoLink) SetLabel(v string) *GetCardInfoResponseBodyExtensionCardContactInfoLink {
	s.Label = &v
	return s
}

func (s *GetCardInfoResponseBodyExtensionCardContactInfoLink) SetValue(v string) *GetCardInfoResponseBodyExtensionCardContactInfoLink {
	s.Value = &v
	return s
}

type GetCardInfoResponseBodyExtensionCardContactInfoTelephone struct {
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetCardInfoResponseBodyExtensionCardContactInfoTelephone) String() string {
	return tea.Prettify(s)
}

func (s GetCardInfoResponseBodyExtensionCardContactInfoTelephone) GoString() string {
	return s.String()
}

func (s *GetCardInfoResponseBodyExtensionCardContactInfoTelephone) SetLabel(v string) *GetCardInfoResponseBodyExtensionCardContactInfoTelephone {
	s.Label = &v
	return s
}

func (s *GetCardInfoResponseBodyExtensionCardContactInfoTelephone) SetValue(v string) *GetCardInfoResponseBodyExtensionCardContactInfoTelephone {
	s.Value = &v
	return s
}

type GetCardInfoResponseBodyExtensionCardContactInfoWorkPhone struct {
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetCardInfoResponseBodyExtensionCardContactInfoWorkPhone) String() string {
	return tea.Prettify(s)
}

func (s GetCardInfoResponseBodyExtensionCardContactInfoWorkPhone) GoString() string {
	return s.String()
}

func (s *GetCardInfoResponseBodyExtensionCardContactInfoWorkPhone) SetLabel(v string) *GetCardInfoResponseBodyExtensionCardContactInfoWorkPhone {
	s.Label = &v
	return s
}

func (s *GetCardInfoResponseBodyExtensionCardContactInfoWorkPhone) SetValue(v string) *GetCardInfoResponseBodyExtensionCardContactInfoWorkPhone {
	s.Value = &v
	return s
}

type GetCardInfoResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCardInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCardInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCardInfoResponse) GoString() string {
	return s.String()
}

func (s *GetCardInfoResponse) SetHeaders(v map[string]*string) *GetCardInfoResponse {
	s.Headers = v
	return s
}

func (s *GetCardInfoResponse) SetStatusCode(v int32) *GetCardInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCardInfoResponse) SetBody(v *GetCardInfoResponseBody) *GetCardInfoResponse {
	s.Body = v
	return s
}

type GetContactHideBySceneSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetContactHideBySceneSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetContactHideBySceneSettingHeaders) GoString() string {
	return s.String()
}

func (s *GetContactHideBySceneSettingHeaders) SetCommonHeaders(v map[string]*string) *GetContactHideBySceneSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetContactHideBySceneSettingHeaders) SetXAcsDingtalkAccessToken(v string) *GetContactHideBySceneSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetContactHideBySceneSettingResponseBody struct {
	Description         *string                                                      `json:"description,omitempty" xml:"description,omitempty"`
	ExcludeDeptIds      []*int64                                                     `json:"excludeDeptIds,omitempty" xml:"excludeDeptIds,omitempty" type:"Repeated"`
	ExcludeTagIds       []*int64                                                     `json:"excludeTagIds,omitempty" xml:"excludeTagIds,omitempty" type:"Repeated"`
	ExcludeUserIds      []*string                                                    `json:"excludeUserIds,omitempty" xml:"excludeUserIds,omitempty" type:"Repeated"`
	Id                  *int64                                                       `json:"id,omitempty" xml:"id,omitempty"`
	Name                *string                                                      `json:"name,omitempty" xml:"name,omitempty"`
	NodeListSceneConfig *GetContactHideBySceneSettingResponseBodyNodeListSceneConfig `json:"nodeListSceneConfig,omitempty" xml:"nodeListSceneConfig,omitempty" type:"Struct"`
	ObjectDeptIds       []*int64                                                     `json:"objectDeptIds,omitempty" xml:"objectDeptIds,omitempty" type:"Repeated"`
	ObjectTagIds        []*int64                                                     `json:"objectTagIds,omitempty" xml:"objectTagIds,omitempty" type:"Repeated"`
	ObjectUserIds       []*string                                                    `json:"objectUserIds,omitempty" xml:"objectUserIds,omitempty" type:"Repeated"`
	ProfileSceneConfig  *GetContactHideBySceneSettingResponseBodyProfileSceneConfig  `json:"profileSceneConfig,omitempty" xml:"profileSceneConfig,omitempty" type:"Struct"`
	SearchSceneConfig   *GetContactHideBySceneSettingResponseBodySearchSceneConfig   `json:"searchSceneConfig,omitempty" xml:"searchSceneConfig,omitempty" type:"Struct"`
}

func (s GetContactHideBySceneSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetContactHideBySceneSettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetContactHideBySceneSettingResponseBody) SetDescription(v string) *GetContactHideBySceneSettingResponseBody {
	s.Description = &v
	return s
}

func (s *GetContactHideBySceneSettingResponseBody) SetExcludeDeptIds(v []*int64) *GetContactHideBySceneSettingResponseBody {
	s.ExcludeDeptIds = v
	return s
}

func (s *GetContactHideBySceneSettingResponseBody) SetExcludeTagIds(v []*int64) *GetContactHideBySceneSettingResponseBody {
	s.ExcludeTagIds = v
	return s
}

func (s *GetContactHideBySceneSettingResponseBody) SetExcludeUserIds(v []*string) *GetContactHideBySceneSettingResponseBody {
	s.ExcludeUserIds = v
	return s
}

func (s *GetContactHideBySceneSettingResponseBody) SetId(v int64) *GetContactHideBySceneSettingResponseBody {
	s.Id = &v
	return s
}

func (s *GetContactHideBySceneSettingResponseBody) SetName(v string) *GetContactHideBySceneSettingResponseBody {
	s.Name = &v
	return s
}

func (s *GetContactHideBySceneSettingResponseBody) SetNodeListSceneConfig(v *GetContactHideBySceneSettingResponseBodyNodeListSceneConfig) *GetContactHideBySceneSettingResponseBody {
	s.NodeListSceneConfig = v
	return s
}

func (s *GetContactHideBySceneSettingResponseBody) SetObjectDeptIds(v []*int64) *GetContactHideBySceneSettingResponseBody {
	s.ObjectDeptIds = v
	return s
}

func (s *GetContactHideBySceneSettingResponseBody) SetObjectTagIds(v []*int64) *GetContactHideBySceneSettingResponseBody {
	s.ObjectTagIds = v
	return s
}

func (s *GetContactHideBySceneSettingResponseBody) SetObjectUserIds(v []*string) *GetContactHideBySceneSettingResponseBody {
	s.ObjectUserIds = v
	return s
}

func (s *GetContactHideBySceneSettingResponseBody) SetProfileSceneConfig(v *GetContactHideBySceneSettingResponseBodyProfileSceneConfig) *GetContactHideBySceneSettingResponseBody {
	s.ProfileSceneConfig = v
	return s
}

func (s *GetContactHideBySceneSettingResponseBody) SetSearchSceneConfig(v *GetContactHideBySceneSettingResponseBodySearchSceneConfig) *GetContactHideBySceneSettingResponseBody {
	s.SearchSceneConfig = v
	return s
}

type GetContactHideBySceneSettingResponseBodyNodeListSceneConfig struct {
	Active               *bool `json:"active,omitempty" xml:"active,omitempty"`
	DeptObjectIncludeEmp *bool `json:"deptObjectIncludeEmp,omitempty" xml:"deptObjectIncludeEmp,omitempty"`
}

func (s GetContactHideBySceneSettingResponseBodyNodeListSceneConfig) String() string {
	return tea.Prettify(s)
}

func (s GetContactHideBySceneSettingResponseBodyNodeListSceneConfig) GoString() string {
	return s.String()
}

func (s *GetContactHideBySceneSettingResponseBodyNodeListSceneConfig) SetActive(v bool) *GetContactHideBySceneSettingResponseBodyNodeListSceneConfig {
	s.Active = &v
	return s
}

func (s *GetContactHideBySceneSettingResponseBodyNodeListSceneConfig) SetDeptObjectIncludeEmp(v bool) *GetContactHideBySceneSettingResponseBodyNodeListSceneConfig {
	s.DeptObjectIncludeEmp = &v
	return s
}

type GetContactHideBySceneSettingResponseBodyProfileSceneConfig struct {
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
}

func (s GetContactHideBySceneSettingResponseBodyProfileSceneConfig) String() string {
	return tea.Prettify(s)
}

func (s GetContactHideBySceneSettingResponseBodyProfileSceneConfig) GoString() string {
	return s.String()
}

func (s *GetContactHideBySceneSettingResponseBodyProfileSceneConfig) SetActive(v bool) *GetContactHideBySceneSettingResponseBodyProfileSceneConfig {
	s.Active = &v
	return s
}

type GetContactHideBySceneSettingResponseBodySearchSceneConfig struct {
	Active               *bool `json:"active,omitempty" xml:"active,omitempty"`
	DeptObjectIncludeEmp *bool `json:"deptObjectIncludeEmp,omitempty" xml:"deptObjectIncludeEmp,omitempty"`
}

func (s GetContactHideBySceneSettingResponseBodySearchSceneConfig) String() string {
	return tea.Prettify(s)
}

func (s GetContactHideBySceneSettingResponseBodySearchSceneConfig) GoString() string {
	return s.String()
}

func (s *GetContactHideBySceneSettingResponseBodySearchSceneConfig) SetActive(v bool) *GetContactHideBySceneSettingResponseBodySearchSceneConfig {
	s.Active = &v
	return s
}

func (s *GetContactHideBySceneSettingResponseBodySearchSceneConfig) SetDeptObjectIncludeEmp(v bool) *GetContactHideBySceneSettingResponseBodySearchSceneConfig {
	s.DeptObjectIncludeEmp = &v
	return s
}

type GetContactHideBySceneSettingResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetContactHideBySceneSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetContactHideBySceneSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetContactHideBySceneSettingResponse) GoString() string {
	return s.String()
}

func (s *GetContactHideBySceneSettingResponse) SetHeaders(v map[string]*string) *GetContactHideBySceneSettingResponse {
	s.Headers = v
	return s
}

func (s *GetContactHideBySceneSettingResponse) SetStatusCode(v int32) *GetContactHideBySceneSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContactHideBySceneSettingResponse) SetBody(v *GetContactHideBySceneSettingResponseBody) *GetContactHideBySceneSettingResponse {
	s.Body = v
	return s
}

type GetCooperateOrgInviteInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCooperateOrgInviteInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCooperateOrgInviteInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetCooperateOrgInviteInfoHeaders) SetCommonHeaders(v map[string]*string) *GetCooperateOrgInviteInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCooperateOrgInviteInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetCooperateOrgInviteInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCooperateOrgInviteInfoResponseBody struct {
	InviteUrl *string `json:"inviteUrl,omitempty" xml:"inviteUrl,omitempty"`
}

func (s GetCooperateOrgInviteInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCooperateOrgInviteInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetCooperateOrgInviteInfoResponseBody) SetInviteUrl(v string) *GetCooperateOrgInviteInfoResponseBody {
	s.InviteUrl = &v
	return s
}

type GetCooperateOrgInviteInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCooperateOrgInviteInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCooperateOrgInviteInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCooperateOrgInviteInfoResponse) GoString() string {
	return s.String()
}

func (s *GetCooperateOrgInviteInfoResponse) SetHeaders(v map[string]*string) *GetCooperateOrgInviteInfoResponse {
	s.Headers = v
	return s
}

func (s *GetCooperateOrgInviteInfoResponse) SetStatusCode(v int32) *GetCooperateOrgInviteInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCooperateOrgInviteInfoResponse) SetBody(v *GetCooperateOrgInviteInfoResponseBody) *GetCooperateOrgInviteInfoResponse {
	s.Body = v
	return s
}

type GetCorpCardStyleListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCorpCardStyleListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCorpCardStyleListHeaders) GoString() string {
	return s.String()
}

func (s *GetCorpCardStyleListHeaders) SetCommonHeaders(v map[string]*string) *GetCorpCardStyleListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCorpCardStyleListHeaders) SetXAcsDingtalkAccessToken(v string) *GetCorpCardStyleListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCorpCardStyleListResponseBody struct {
	Result []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetCorpCardStyleListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCorpCardStyleListResponseBody) GoString() string {
	return s.String()
}

func (s *GetCorpCardStyleListResponseBody) SetResult(v []map[string]interface{}) *GetCorpCardStyleListResponseBody {
	s.Result = v
	return s
}

type GetCorpCardStyleListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCorpCardStyleListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCorpCardStyleListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCorpCardStyleListResponse) GoString() string {
	return s.String()
}

func (s *GetCorpCardStyleListResponse) SetHeaders(v map[string]*string) *GetCorpCardStyleListResponse {
	s.Headers = v
	return s
}

func (s *GetCorpCardStyleListResponse) SetStatusCode(v int32) *GetCorpCardStyleListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCorpCardStyleListResponse) SetBody(v *GetCorpCardStyleListResponseBody) *GetCorpCardStyleListResponse {
	s.Body = v
	return s
}

type GetDingIdByMigrationDingIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDingIdByMigrationDingIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDingIdByMigrationDingIdHeaders) GoString() string {
	return s.String()
}

func (s *GetDingIdByMigrationDingIdHeaders) SetCommonHeaders(v map[string]*string) *GetDingIdByMigrationDingIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDingIdByMigrationDingIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetDingIdByMigrationDingIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDingIdByMigrationDingIdRequest struct {
	MigrationDingId *string `json:"migrationDingId,omitempty" xml:"migrationDingId,omitempty"`
}

func (s GetDingIdByMigrationDingIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDingIdByMigrationDingIdRequest) GoString() string {
	return s.String()
}

func (s *GetDingIdByMigrationDingIdRequest) SetMigrationDingId(v string) *GetDingIdByMigrationDingIdRequest {
	s.MigrationDingId = &v
	return s
}

type GetDingIdByMigrationDingIdResponseBody struct {
	DingId *string `json:"dingId,omitempty" xml:"dingId,omitempty"`
}

func (s GetDingIdByMigrationDingIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDingIdByMigrationDingIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetDingIdByMigrationDingIdResponseBody) SetDingId(v string) *GetDingIdByMigrationDingIdResponseBody {
	s.DingId = &v
	return s
}

type GetDingIdByMigrationDingIdResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDingIdByMigrationDingIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDingIdByMigrationDingIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDingIdByMigrationDingIdResponse) GoString() string {
	return s.String()
}

func (s *GetDingIdByMigrationDingIdResponse) SetHeaders(v map[string]*string) *GetDingIdByMigrationDingIdResponse {
	s.Headers = v
	return s
}

func (s *GetDingIdByMigrationDingIdResponse) SetStatusCode(v int32) *GetDingIdByMigrationDingIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDingIdByMigrationDingIdResponse) SetBody(v *GetDingIdByMigrationDingIdResponseBody) *GetDingIdByMigrationDingIdResponse {
	s.Body = v
	return s
}

type GetEmpAttributeHideBySceneSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetEmpAttributeHideBySceneSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetEmpAttributeHideBySceneSettingHeaders) GoString() string {
	return s.String()
}

func (s *GetEmpAttributeHideBySceneSettingHeaders) SetCommonHeaders(v map[string]*string) *GetEmpAttributeHideBySceneSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetEmpAttributeHideBySceneSettingHeaders) SetXAcsDingtalkAccessToken(v string) *GetEmpAttributeHideBySceneSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetEmpAttributeHideBySceneSettingResponseBody struct {
	ChatSubtitleConfig *GetEmpAttributeHideBySceneSettingResponseBodyChatSubtitleConfig `json:"chatSubtitleConfig,omitempty" xml:"chatSubtitleConfig,omitempty" type:"Struct"`
	Description        *string                                                          `json:"description,omitempty" xml:"description,omitempty"`
	ExcludeDeptIds     []*int64                                                         `json:"excludeDeptIds,omitempty" xml:"excludeDeptIds,omitempty" type:"Repeated"`
	ExcludeTagIds      []*int64                                                         `json:"excludeTagIds,omitempty" xml:"excludeTagIds,omitempty" type:"Repeated"`
	ExcludeUserIds     []*string                                                        `json:"excludeUserIds,omitempty" xml:"excludeUserIds,omitempty" type:"Repeated"`
	HideFields         []*string                                                        `json:"hideFields,omitempty" xml:"hideFields,omitempty" type:"Repeated"`
	Id                 *int64                                                           `json:"id,omitempty" xml:"id,omitempty"`
	Name               *string                                                          `json:"name,omitempty" xml:"name,omitempty"`
	ObjectDeptIds      []*int64                                                         `json:"objectDeptIds,omitempty" xml:"objectDeptIds,omitempty" type:"Repeated"`
	ObjectTagIds       []*int64                                                         `json:"objectTagIds,omitempty" xml:"objectTagIds,omitempty" type:"Repeated"`
	ObjectUserIds      []*string                                                        `json:"objectUserIds,omitempty" xml:"objectUserIds,omitempty" type:"Repeated"`
	ProfileSceneConfig *GetEmpAttributeHideBySceneSettingResponseBodyProfileSceneConfig `json:"profileSceneConfig,omitempty" xml:"profileSceneConfig,omitempty" type:"Struct"`
	SearchSceneConfig  *GetEmpAttributeHideBySceneSettingResponseBodySearchSceneConfig  `json:"searchSceneConfig,omitempty" xml:"searchSceneConfig,omitempty" type:"Struct"`
}

func (s GetEmpAttributeHideBySceneSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEmpAttributeHideBySceneSettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetEmpAttributeHideBySceneSettingResponseBody) SetChatSubtitleConfig(v *GetEmpAttributeHideBySceneSettingResponseBodyChatSubtitleConfig) *GetEmpAttributeHideBySceneSettingResponseBody {
	s.ChatSubtitleConfig = v
	return s
}

func (s *GetEmpAttributeHideBySceneSettingResponseBody) SetDescription(v string) *GetEmpAttributeHideBySceneSettingResponseBody {
	s.Description = &v
	return s
}

func (s *GetEmpAttributeHideBySceneSettingResponseBody) SetExcludeDeptIds(v []*int64) *GetEmpAttributeHideBySceneSettingResponseBody {
	s.ExcludeDeptIds = v
	return s
}

func (s *GetEmpAttributeHideBySceneSettingResponseBody) SetExcludeTagIds(v []*int64) *GetEmpAttributeHideBySceneSettingResponseBody {
	s.ExcludeTagIds = v
	return s
}

func (s *GetEmpAttributeHideBySceneSettingResponseBody) SetExcludeUserIds(v []*string) *GetEmpAttributeHideBySceneSettingResponseBody {
	s.ExcludeUserIds = v
	return s
}

func (s *GetEmpAttributeHideBySceneSettingResponseBody) SetHideFields(v []*string) *GetEmpAttributeHideBySceneSettingResponseBody {
	s.HideFields = v
	return s
}

func (s *GetEmpAttributeHideBySceneSettingResponseBody) SetId(v int64) *GetEmpAttributeHideBySceneSettingResponseBody {
	s.Id = &v
	return s
}

func (s *GetEmpAttributeHideBySceneSettingResponseBody) SetName(v string) *GetEmpAttributeHideBySceneSettingResponseBody {
	s.Name = &v
	return s
}

func (s *GetEmpAttributeHideBySceneSettingResponseBody) SetObjectDeptIds(v []*int64) *GetEmpAttributeHideBySceneSettingResponseBody {
	s.ObjectDeptIds = v
	return s
}

func (s *GetEmpAttributeHideBySceneSettingResponseBody) SetObjectTagIds(v []*int64) *GetEmpAttributeHideBySceneSettingResponseBody {
	s.ObjectTagIds = v
	return s
}

func (s *GetEmpAttributeHideBySceneSettingResponseBody) SetObjectUserIds(v []*string) *GetEmpAttributeHideBySceneSettingResponseBody {
	s.ObjectUserIds = v
	return s
}

func (s *GetEmpAttributeHideBySceneSettingResponseBody) SetProfileSceneConfig(v *GetEmpAttributeHideBySceneSettingResponseBodyProfileSceneConfig) *GetEmpAttributeHideBySceneSettingResponseBody {
	s.ProfileSceneConfig = v
	return s
}

func (s *GetEmpAttributeHideBySceneSettingResponseBody) SetSearchSceneConfig(v *GetEmpAttributeHideBySceneSettingResponseBodySearchSceneConfig) *GetEmpAttributeHideBySceneSettingResponseBody {
	s.SearchSceneConfig = v
	return s
}

type GetEmpAttributeHideBySceneSettingResponseBodyChatSubtitleConfig struct {
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
}

func (s GetEmpAttributeHideBySceneSettingResponseBodyChatSubtitleConfig) String() string {
	return tea.Prettify(s)
}

func (s GetEmpAttributeHideBySceneSettingResponseBodyChatSubtitleConfig) GoString() string {
	return s.String()
}

func (s *GetEmpAttributeHideBySceneSettingResponseBodyChatSubtitleConfig) SetActive(v bool) *GetEmpAttributeHideBySceneSettingResponseBodyChatSubtitleConfig {
	s.Active = &v
	return s
}

type GetEmpAttributeHideBySceneSettingResponseBodyProfileSceneConfig struct {
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
}

func (s GetEmpAttributeHideBySceneSettingResponseBodyProfileSceneConfig) String() string {
	return tea.Prettify(s)
}

func (s GetEmpAttributeHideBySceneSettingResponseBodyProfileSceneConfig) GoString() string {
	return s.String()
}

func (s *GetEmpAttributeHideBySceneSettingResponseBodyProfileSceneConfig) SetActive(v bool) *GetEmpAttributeHideBySceneSettingResponseBodyProfileSceneConfig {
	s.Active = &v
	return s
}

type GetEmpAttributeHideBySceneSettingResponseBodySearchSceneConfig struct {
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
}

func (s GetEmpAttributeHideBySceneSettingResponseBodySearchSceneConfig) String() string {
	return tea.Prettify(s)
}

func (s GetEmpAttributeHideBySceneSettingResponseBodySearchSceneConfig) GoString() string {
	return s.String()
}

func (s *GetEmpAttributeHideBySceneSettingResponseBodySearchSceneConfig) SetActive(v bool) *GetEmpAttributeHideBySceneSettingResponseBodySearchSceneConfig {
	s.Active = &v
	return s
}

type GetEmpAttributeHideBySceneSettingResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetEmpAttributeHideBySceneSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEmpAttributeHideBySceneSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEmpAttributeHideBySceneSettingResponse) GoString() string {
	return s.String()
}

func (s *GetEmpAttributeHideBySceneSettingResponse) SetHeaders(v map[string]*string) *GetEmpAttributeHideBySceneSettingResponse {
	s.Headers = v
	return s
}

func (s *GetEmpAttributeHideBySceneSettingResponse) SetStatusCode(v int32) *GetEmpAttributeHideBySceneSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEmpAttributeHideBySceneSettingResponse) SetBody(v *GetEmpAttributeHideBySceneSettingResponseBody) *GetEmpAttributeHideBySceneSettingResponse {
	s.Body = v
	return s
}

type GetLatestDingIndexHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetLatestDingIndexHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetLatestDingIndexHeaders) GoString() string {
	return s.String()
}

func (s *GetLatestDingIndexHeaders) SetCommonHeaders(v map[string]*string) *GetLatestDingIndexHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetLatestDingIndexHeaders) SetXAcsDingtalkAccessToken(v string) *GetLatestDingIndexHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetLatestDingIndexResponseBody struct {
	IdxCarbon     *float32 `json:"idxCarbon,omitempty" xml:"idxCarbon,omitempty"`
	IdxEfficiency *float32 `json:"idxEfficiency,omitempty" xml:"idxEfficiency,omitempty"`
	IdxMonthlyAvg *float32 `json:"idxMonthlyAvg,omitempty" xml:"idxMonthlyAvg,omitempty"`
	IdxTotal      *float32 `json:"idxTotal,omitempty" xml:"idxTotal,omitempty"`
	StatDate      *string  `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s GetLatestDingIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLatestDingIndexResponseBody) GoString() string {
	return s.String()
}

func (s *GetLatestDingIndexResponseBody) SetIdxCarbon(v float32) *GetLatestDingIndexResponseBody {
	s.IdxCarbon = &v
	return s
}

func (s *GetLatestDingIndexResponseBody) SetIdxEfficiency(v float32) *GetLatestDingIndexResponseBody {
	s.IdxEfficiency = &v
	return s
}

func (s *GetLatestDingIndexResponseBody) SetIdxMonthlyAvg(v float32) *GetLatestDingIndexResponseBody {
	s.IdxMonthlyAvg = &v
	return s
}

func (s *GetLatestDingIndexResponseBody) SetIdxTotal(v float32) *GetLatestDingIndexResponseBody {
	s.IdxTotal = &v
	return s
}

func (s *GetLatestDingIndexResponseBody) SetStatDate(v string) *GetLatestDingIndexResponseBody {
	s.StatDate = &v
	return s
}

type GetLatestDingIndexResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLatestDingIndexResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLatestDingIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLatestDingIndexResponse) GoString() string {
	return s.String()
}

func (s *GetLatestDingIndexResponse) SetHeaders(v map[string]*string) *GetLatestDingIndexResponse {
	s.Headers = v
	return s
}

func (s *GetLatestDingIndexResponse) SetStatusCode(v int32) *GetLatestDingIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLatestDingIndexResponse) SetBody(v *GetLatestDingIndexResponseBody) *GetLatestDingIndexResponse {
	s.Body = v
	return s
}

type GetMigrationDingIdByDingIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetMigrationDingIdByDingIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMigrationDingIdByDingIdHeaders) GoString() string {
	return s.String()
}

func (s *GetMigrationDingIdByDingIdHeaders) SetCommonHeaders(v map[string]*string) *GetMigrationDingIdByDingIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMigrationDingIdByDingIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetMigrationDingIdByDingIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetMigrationDingIdByDingIdRequest struct {
	DingId *string `json:"dingId,omitempty" xml:"dingId,omitempty"`
}

func (s GetMigrationDingIdByDingIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMigrationDingIdByDingIdRequest) GoString() string {
	return s.String()
}

func (s *GetMigrationDingIdByDingIdRequest) SetDingId(v string) *GetMigrationDingIdByDingIdRequest {
	s.DingId = &v
	return s
}

type GetMigrationDingIdByDingIdResponseBody struct {
	MigrationDingIdList map[string]interface{} `json:"migrationDingIdList,omitempty" xml:"migrationDingIdList,omitempty"`
}

func (s GetMigrationDingIdByDingIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMigrationDingIdByDingIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetMigrationDingIdByDingIdResponseBody) SetMigrationDingIdList(v map[string]interface{}) *GetMigrationDingIdByDingIdResponseBody {
	s.MigrationDingIdList = v
	return s
}

type GetMigrationDingIdByDingIdResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMigrationDingIdByDingIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMigrationDingIdByDingIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMigrationDingIdByDingIdResponse) GoString() string {
	return s.String()
}

func (s *GetMigrationDingIdByDingIdResponse) SetHeaders(v map[string]*string) *GetMigrationDingIdByDingIdResponse {
	s.Headers = v
	return s
}

func (s *GetMigrationDingIdByDingIdResponse) SetStatusCode(v int32) *GetMigrationDingIdByDingIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMigrationDingIdByDingIdResponse) SetBody(v *GetMigrationDingIdByDingIdResponseBody) *GetMigrationDingIdByDingIdResponse {
	s.Body = v
	return s
}

type GetMigrationUnionIdByUnionIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetMigrationUnionIdByUnionIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMigrationUnionIdByUnionIdHeaders) GoString() string {
	return s.String()
}

func (s *GetMigrationUnionIdByUnionIdHeaders) SetCommonHeaders(v map[string]*string) *GetMigrationUnionIdByUnionIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMigrationUnionIdByUnionIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetMigrationUnionIdByUnionIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetMigrationUnionIdByUnionIdRequest struct {
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetMigrationUnionIdByUnionIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMigrationUnionIdByUnionIdRequest) GoString() string {
	return s.String()
}

func (s *GetMigrationUnionIdByUnionIdRequest) SetUnionId(v string) *GetMigrationUnionIdByUnionIdRequest {
	s.UnionId = &v
	return s
}

type GetMigrationUnionIdByUnionIdResponseBody struct {
	MigrationUnionIdList map[string]interface{} `json:"migrationUnionIdList,omitempty" xml:"migrationUnionIdList,omitempty"`
}

func (s GetMigrationUnionIdByUnionIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMigrationUnionIdByUnionIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetMigrationUnionIdByUnionIdResponseBody) SetMigrationUnionIdList(v map[string]interface{}) *GetMigrationUnionIdByUnionIdResponseBody {
	s.MigrationUnionIdList = v
	return s
}

type GetMigrationUnionIdByUnionIdResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMigrationUnionIdByUnionIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMigrationUnionIdByUnionIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMigrationUnionIdByUnionIdResponse) GoString() string {
	return s.String()
}

func (s *GetMigrationUnionIdByUnionIdResponse) SetHeaders(v map[string]*string) *GetMigrationUnionIdByUnionIdResponse {
	s.Headers = v
	return s
}

func (s *GetMigrationUnionIdByUnionIdResponse) SetStatusCode(v int32) *GetMigrationUnionIdByUnionIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMigrationUnionIdByUnionIdResponse) SetBody(v *GetMigrationUnionIdByUnionIdResponseBody) *GetMigrationUnionIdByUnionIdResponse {
	s.Body = v
	return s
}

type GetOrgAuthInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOrgAuthInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOrgAuthInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetOrgAuthInfoHeaders) SetCommonHeaders(v map[string]*string) *GetOrgAuthInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOrgAuthInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetOrgAuthInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOrgAuthInfoRequest struct {
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
}

func (s GetOrgAuthInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOrgAuthInfoRequest) GoString() string {
	return s.String()
}

func (s *GetOrgAuthInfoRequest) SetTargetCorpId(v string) *GetOrgAuthInfoRequest {
	s.TargetCorpId = &v
	return s
}

type GetOrgAuthInfoResponseBody struct {
	AuthLevel           *int64  `json:"authLevel,omitempty" xml:"authLevel,omitempty"`
	LegalPerson         *string `json:"legalPerson,omitempty" xml:"legalPerson,omitempty"`
	LicenseOrgName      *string `json:"licenseOrgName,omitempty" xml:"licenseOrgName,omitempty"`
	LicenseUrl          *string `json:"licenseUrl,omitempty" xml:"licenseUrl,omitempty"`
	OrgName             *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	OrganizationCode    *string `json:"organizationCode,omitempty" xml:"organizationCode,omitempty"`
	RegistrationNum     *string `json:"registrationNum,omitempty" xml:"registrationNum,omitempty"`
	UnifiedSocialCredit *string `json:"unifiedSocialCredit,omitempty" xml:"unifiedSocialCredit,omitempty"`
}

func (s GetOrgAuthInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrgAuthInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrgAuthInfoResponseBody) SetAuthLevel(v int64) *GetOrgAuthInfoResponseBody {
	s.AuthLevel = &v
	return s
}

func (s *GetOrgAuthInfoResponseBody) SetLegalPerson(v string) *GetOrgAuthInfoResponseBody {
	s.LegalPerson = &v
	return s
}

func (s *GetOrgAuthInfoResponseBody) SetLicenseOrgName(v string) *GetOrgAuthInfoResponseBody {
	s.LicenseOrgName = &v
	return s
}

func (s *GetOrgAuthInfoResponseBody) SetLicenseUrl(v string) *GetOrgAuthInfoResponseBody {
	s.LicenseUrl = &v
	return s
}

func (s *GetOrgAuthInfoResponseBody) SetOrgName(v string) *GetOrgAuthInfoResponseBody {
	s.OrgName = &v
	return s
}

func (s *GetOrgAuthInfoResponseBody) SetOrganizationCode(v string) *GetOrgAuthInfoResponseBody {
	s.OrganizationCode = &v
	return s
}

func (s *GetOrgAuthInfoResponseBody) SetRegistrationNum(v string) *GetOrgAuthInfoResponseBody {
	s.RegistrationNum = &v
	return s
}

func (s *GetOrgAuthInfoResponseBody) SetUnifiedSocialCredit(v string) *GetOrgAuthInfoResponseBody {
	s.UnifiedSocialCredit = &v
	return s
}

type GetOrgAuthInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOrgAuthInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOrgAuthInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrgAuthInfoResponse) GoString() string {
	return s.String()
}

func (s *GetOrgAuthInfoResponse) SetHeaders(v map[string]*string) *GetOrgAuthInfoResponse {
	s.Headers = v
	return s
}

func (s *GetOrgAuthInfoResponse) SetStatusCode(v int32) *GetOrgAuthInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrgAuthInfoResponse) SetBody(v *GetOrgAuthInfoResponseBody) *GetOrgAuthInfoResponse {
	s.Body = v
	return s
}

type GetTranslateFileJobResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTranslateFileJobResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTranslateFileJobResultHeaders) GoString() string {
	return s.String()
}

func (s *GetTranslateFileJobResultHeaders) SetCommonHeaders(v map[string]*string) *GetTranslateFileJobResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTranslateFileJobResultHeaders) SetXAcsDingtalkAccessToken(v string) *GetTranslateFileJobResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTranslateFileJobResultRequest struct {
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
}

func (s GetTranslateFileJobResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTranslateFileJobResultRequest) GoString() string {
	return s.String()
}

func (s *GetTranslateFileJobResultRequest) SetJobId(v string) *GetTranslateFileJobResultRequest {
	s.JobId = &v
	return s
}

type GetTranslateFileJobResultResponseBody struct {
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	Url    *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetTranslateFileJobResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTranslateFileJobResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetTranslateFileJobResultResponseBody) SetStatus(v string) *GetTranslateFileJobResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetTranslateFileJobResultResponseBody) SetUrl(v string) *GetTranslateFileJobResultResponseBody {
	s.Url = &v
	return s
}

type GetTranslateFileJobResultResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTranslateFileJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTranslateFileJobResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTranslateFileJobResultResponse) GoString() string {
	return s.String()
}

func (s *GetTranslateFileJobResultResponse) SetHeaders(v map[string]*string) *GetTranslateFileJobResultResponse {
	s.Headers = v
	return s
}

func (s *GetTranslateFileJobResultResponse) SetStatusCode(v int32) *GetTranslateFileJobResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTranslateFileJobResultResponse) SetBody(v *GetTranslateFileJobResultResponseBody) *GetTranslateFileJobResultResponse {
	s.Body = v
	return s
}

type GetUnionIdByMigrationUnionIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUnionIdByMigrationUnionIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUnionIdByMigrationUnionIdHeaders) GoString() string {
	return s.String()
}

func (s *GetUnionIdByMigrationUnionIdHeaders) SetCommonHeaders(v map[string]*string) *GetUnionIdByMigrationUnionIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUnionIdByMigrationUnionIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetUnionIdByMigrationUnionIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUnionIdByMigrationUnionIdRequest struct {
	MigrationUnionId *string `json:"migrationUnionId,omitempty" xml:"migrationUnionId,omitempty"`
}

func (s GetUnionIdByMigrationUnionIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUnionIdByMigrationUnionIdRequest) GoString() string {
	return s.String()
}

func (s *GetUnionIdByMigrationUnionIdRequest) SetMigrationUnionId(v string) *GetUnionIdByMigrationUnionIdRequest {
	s.MigrationUnionId = &v
	return s
}

type GetUnionIdByMigrationUnionIdResponseBody struct {
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetUnionIdByMigrationUnionIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUnionIdByMigrationUnionIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetUnionIdByMigrationUnionIdResponseBody) SetUnionId(v string) *GetUnionIdByMigrationUnionIdResponseBody {
	s.UnionId = &v
	return s
}

type GetUnionIdByMigrationUnionIdResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUnionIdByMigrationUnionIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUnionIdByMigrationUnionIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUnionIdByMigrationUnionIdResponse) GoString() string {
	return s.String()
}

func (s *GetUnionIdByMigrationUnionIdResponse) SetHeaders(v map[string]*string) *GetUnionIdByMigrationUnionIdResponse {
	s.Headers = v
	return s
}

func (s *GetUnionIdByMigrationUnionIdResponse) SetStatusCode(v int32) *GetUnionIdByMigrationUnionIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUnionIdByMigrationUnionIdResponse) SetBody(v *GetUnionIdByMigrationUnionIdResponseBody) *GetUnionIdByMigrationUnionIdResponse {
	s.Body = v
	return s
}

type GetUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserHeaders) GoString() string {
	return s.String()
}

func (s *GetUserHeaders) SetCommonHeaders(v map[string]*string) *GetUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserResponseBody struct {
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	Email     *string `json:"email,omitempty" xml:"email,omitempty"`
	Mobile    *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	Nick      *string `json:"nick,omitempty" xml:"nick,omitempty"`
	OpenId    *string `json:"openId,omitempty" xml:"openId,omitempty"`
	StateCode *string `json:"stateCode,omitempty" xml:"stateCode,omitempty"`
	UnionId   *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetAvatarUrl(v string) *GetUserResponseBody {
	s.AvatarUrl = &v
	return s
}

func (s *GetUserResponseBody) SetEmail(v string) *GetUserResponseBody {
	s.Email = &v
	return s
}

func (s *GetUserResponseBody) SetMobile(v string) *GetUserResponseBody {
	s.Mobile = &v
	return s
}

func (s *GetUserResponseBody) SetNick(v string) *GetUserResponseBody {
	s.Nick = &v
	return s
}

func (s *GetUserResponseBody) SetOpenId(v string) *GetUserResponseBody {
	s.OpenId = &v
	return s
}

func (s *GetUserResponseBody) SetStateCode(v string) *GetUserResponseBody {
	s.StateCode = &v
	return s
}

func (s *GetUserResponseBody) SetUnionId(v string) *GetUserResponseBody {
	s.UnionId = &v
	return s
}

type GetUserResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetStatusCode(v int32) *GetUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type GetUserCardHolderListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserCardHolderListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserCardHolderListHeaders) GoString() string {
	return s.String()
}

func (s *GetUserCardHolderListHeaders) SetCommonHeaders(v map[string]*string) *GetUserCardHolderListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserCardHolderListHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserCardHolderListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserCardHolderListRequest struct {
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetUserCardHolderListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserCardHolderListRequest) GoString() string {
	return s.String()
}

func (s *GetUserCardHolderListRequest) SetMaxResults(v int32) *GetUserCardHolderListRequest {
	s.MaxResults = &v
	return s
}

func (s *GetUserCardHolderListRequest) SetNextToken(v int64) *GetUserCardHolderListRequest {
	s.NextToken = &v
	return s
}

type GetUserCardHolderListResponseBody struct {
	HasMore    *bool                                    `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*GetUserCardHolderListResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextToken  *int64                                   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	TotalCount *int32                                   `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetUserCardHolderListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserCardHolderListResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserCardHolderListResponseBody) SetHasMore(v bool) *GetUserCardHolderListResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetUserCardHolderListResponseBody) SetList(v []*GetUserCardHolderListResponseBodyList) *GetUserCardHolderListResponseBody {
	s.List = v
	return s
}

func (s *GetUserCardHolderListResponseBody) SetNextToken(v int64) *GetUserCardHolderListResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetUserCardHolderListResponseBody) SetTotalCount(v int32) *GetUserCardHolderListResponseBody {
	s.TotalCount = &v
	return s
}

type GetUserCardHolderListResponseBodyList struct {
	AvatarUrl          *string                `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	CardAcceptStatus   *int32                 `json:"cardAcceptStatus,omitempty" xml:"cardAcceptStatus,omitempty"`
	CardAcceptTimeLong *int64                 `json:"cardAcceptTimeLong,omitempty" xml:"cardAcceptTimeLong,omitempty"`
	CardId             *string                `json:"cardId,omitempty" xml:"cardId,omitempty"`
	CardSource         *int32                 `json:"cardSource,omitempty" xml:"cardSource,omitempty"`
	Extension          map[string]interface{} `json:"extension,omitempty" xml:"extension,omitempty"`
	IndustryName       *string                `json:"industryName,omitempty" xml:"industryName,omitempty"`
	Introduce          *string                `json:"introduce,omitempty" xml:"introduce,omitempty"`
	Name               *string                `json:"name,omitempty" xml:"name,omitempty"`
	OrgName            *string                `json:"orgName,omitempty" xml:"orgName,omitempty"`
	TemplateId         *string                `json:"templateId,omitempty" xml:"templateId,omitempty"`
	Title              *string                `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetUserCardHolderListResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s GetUserCardHolderListResponseBodyList) GoString() string {
	return s.String()
}

func (s *GetUserCardHolderListResponseBodyList) SetAvatarUrl(v string) *GetUserCardHolderListResponseBodyList {
	s.AvatarUrl = &v
	return s
}

func (s *GetUserCardHolderListResponseBodyList) SetCardAcceptStatus(v int32) *GetUserCardHolderListResponseBodyList {
	s.CardAcceptStatus = &v
	return s
}

func (s *GetUserCardHolderListResponseBodyList) SetCardAcceptTimeLong(v int64) *GetUserCardHolderListResponseBodyList {
	s.CardAcceptTimeLong = &v
	return s
}

func (s *GetUserCardHolderListResponseBodyList) SetCardId(v string) *GetUserCardHolderListResponseBodyList {
	s.CardId = &v
	return s
}

func (s *GetUserCardHolderListResponseBodyList) SetCardSource(v int32) *GetUserCardHolderListResponseBodyList {
	s.CardSource = &v
	return s
}

func (s *GetUserCardHolderListResponseBodyList) SetExtension(v map[string]interface{}) *GetUserCardHolderListResponseBodyList {
	s.Extension = v
	return s
}

func (s *GetUserCardHolderListResponseBodyList) SetIndustryName(v string) *GetUserCardHolderListResponseBodyList {
	s.IndustryName = &v
	return s
}

func (s *GetUserCardHolderListResponseBodyList) SetIntroduce(v string) *GetUserCardHolderListResponseBodyList {
	s.Introduce = &v
	return s
}

func (s *GetUserCardHolderListResponseBodyList) SetName(v string) *GetUserCardHolderListResponseBodyList {
	s.Name = &v
	return s
}

func (s *GetUserCardHolderListResponseBodyList) SetOrgName(v string) *GetUserCardHolderListResponseBodyList {
	s.OrgName = &v
	return s
}

func (s *GetUserCardHolderListResponseBodyList) SetTemplateId(v string) *GetUserCardHolderListResponseBodyList {
	s.TemplateId = &v
	return s
}

func (s *GetUserCardHolderListResponseBodyList) SetTitle(v string) *GetUserCardHolderListResponseBodyList {
	s.Title = &v
	return s
}

type GetUserCardHolderListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserCardHolderListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserCardHolderListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserCardHolderListResponse) GoString() string {
	return s.String()
}

func (s *GetUserCardHolderListResponse) SetHeaders(v map[string]*string) *GetUserCardHolderListResponse {
	s.Headers = v
	return s
}

func (s *GetUserCardHolderListResponse) SetStatusCode(v int32) *GetUserCardHolderListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserCardHolderListResponse) SetBody(v *GetUserCardHolderListResponseBody) *GetUserCardHolderListResponse {
	s.Body = v
	return s
}

type IsFriendHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IsFriendHeaders) String() string {
	return tea.Prettify(s)
}

func (s IsFriendHeaders) GoString() string {
	return s.String()
}

func (s *IsFriendHeaders) SetCommonHeaders(v map[string]*string) *IsFriendHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IsFriendHeaders) SetXAcsDingtalkAccessToken(v string) *IsFriendHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IsFriendRequest struct {
	MobileNo *string `json:"mobileNo,omitempty" xml:"mobileNo,omitempty"`
	UserId   *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s IsFriendRequest) String() string {
	return tea.Prettify(s)
}

func (s IsFriendRequest) GoString() string {
	return s.String()
}

func (s *IsFriendRequest) SetMobileNo(v string) *IsFriendRequest {
	s.MobileNo = &v
	return s
}

func (s *IsFriendRequest) SetUserId(v string) *IsFriendRequest {
	s.UserId = &v
	return s
}

type IsFriendResponseBody struct {
	IsFriend *bool `json:"isFriend,omitempty" xml:"isFriend,omitempty"`
}

func (s IsFriendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IsFriendResponseBody) GoString() string {
	return s.String()
}

func (s *IsFriendResponseBody) SetIsFriend(v bool) *IsFriendResponseBody {
	s.IsFriend = &v
	return s
}

type IsFriendResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IsFriendResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IsFriendResponse) String() string {
	return tea.Prettify(s)
}

func (s IsFriendResponse) GoString() string {
	return s.String()
}

func (s *IsFriendResponse) SetHeaders(v map[string]*string) *IsFriendResponse {
	s.Headers = v
	return s
}

func (s *IsFriendResponse) SetStatusCode(v int32) *IsFriendResponse {
	s.StatusCode = &v
	return s
}

func (s *IsFriendResponse) SetBody(v *IsFriendResponseBody) *IsFriendResponse {
	s.Body = v
	return s
}

type IsvCardEventPushHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IsvCardEventPushHeaders) String() string {
	return tea.Prettify(s)
}

func (s IsvCardEventPushHeaders) GoString() string {
	return s.String()
}

func (s *IsvCardEventPushHeaders) SetCommonHeaders(v map[string]*string) *IsvCardEventPushHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IsvCardEventPushHeaders) SetXAcsDingtalkAccessToken(v string) *IsvCardEventPushHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IsvCardEventPushRequest struct {
	EventParams map[string]interface{} `json:"eventParams,omitempty" xml:"eventParams,omitempty"`
	EventType   *string                `json:"eventType,omitempty" xml:"eventType,omitempty"`
	IsvCardId   *string                `json:"isvCardId,omitempty" xml:"isvCardId,omitempty"`
	IsvToken    *string                `json:"isvToken,omitempty" xml:"isvToken,omitempty"`
	IsvUid      *string                `json:"isvUid,omitempty" xml:"isvUid,omitempty"`
}

func (s IsvCardEventPushRequest) String() string {
	return tea.Prettify(s)
}

func (s IsvCardEventPushRequest) GoString() string {
	return s.String()
}

func (s *IsvCardEventPushRequest) SetEventParams(v map[string]interface{}) *IsvCardEventPushRequest {
	s.EventParams = v
	return s
}

func (s *IsvCardEventPushRequest) SetEventType(v string) *IsvCardEventPushRequest {
	s.EventType = &v
	return s
}

func (s *IsvCardEventPushRequest) SetIsvCardId(v string) *IsvCardEventPushRequest {
	s.IsvCardId = &v
	return s
}

func (s *IsvCardEventPushRequest) SetIsvToken(v string) *IsvCardEventPushRequest {
	s.IsvToken = &v
	return s
}

func (s *IsvCardEventPushRequest) SetIsvUid(v string) *IsvCardEventPushRequest {
	s.IsvUid = &v
	return s
}

type IsvCardEventPushResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s IsvCardEventPushResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IsvCardEventPushResponseBody) GoString() string {
	return s.String()
}

func (s *IsvCardEventPushResponseBody) SetSuccess(v bool) *IsvCardEventPushResponseBody {
	s.Success = &v
	return s
}

type IsvCardEventPushResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IsvCardEventPushResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IsvCardEventPushResponse) String() string {
	return tea.Prettify(s)
}

func (s IsvCardEventPushResponse) GoString() string {
	return s.String()
}

func (s *IsvCardEventPushResponse) SetHeaders(v map[string]*string) *IsvCardEventPushResponse {
	s.Headers = v
	return s
}

func (s *IsvCardEventPushResponse) SetStatusCode(v int32) *IsvCardEventPushResponse {
	s.StatusCode = &v
	return s
}

func (s *IsvCardEventPushResponse) SetBody(v *IsvCardEventPushResponseBody) *IsvCardEventPushResponse {
	s.Body = v
	return s
}

type ListBasicRoleInPageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListBasicRoleInPageHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListBasicRoleInPageHeaders) GoString() string {
	return s.String()
}

func (s *ListBasicRoleInPageHeaders) SetCommonHeaders(v map[string]*string) *ListBasicRoleInPageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListBasicRoleInPageHeaders) SetXAcsDingtalkAccessToken(v string) *ListBasicRoleInPageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListBasicRoleInPageRequest struct {
	AgentId    *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int64  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListBasicRoleInPageRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBasicRoleInPageRequest) GoString() string {
	return s.String()
}

func (s *ListBasicRoleInPageRequest) SetAgentId(v string) *ListBasicRoleInPageRequest {
	s.AgentId = &v
	return s
}

func (s *ListBasicRoleInPageRequest) SetMaxResults(v int32) *ListBasicRoleInPageRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBasicRoleInPageRequest) SetNextToken(v int64) *ListBasicRoleInPageRequest {
	s.NextToken = &v
	return s
}

type ListBasicRoleInPageResponseBody struct {
	HasMore   *bool                                  `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List      []*ListBasicRoleInPageResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextToken *int64                                 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListBasicRoleInPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBasicRoleInPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListBasicRoleInPageResponseBody) SetHasMore(v bool) *ListBasicRoleInPageResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListBasicRoleInPageResponseBody) SetList(v []*ListBasicRoleInPageResponseBodyList) *ListBasicRoleInPageResponseBody {
	s.List = v
	return s
}

func (s *ListBasicRoleInPageResponseBody) SetNextToken(v int64) *ListBasicRoleInPageResponseBody {
	s.NextToken = &v
	return s
}

type ListBasicRoleInPageResponseBodyList struct {
	OpenAction    *ListBasicRoleInPageResponseBodyListOpenAction    `json:"openAction,omitempty" xml:"openAction,omitempty" type:"Struct"`
	OpenMembers   []*ListBasicRoleInPageResponseBodyListOpenMembers `json:"openMembers,omitempty" xml:"openMembers,omitempty" type:"Repeated"`
	OpenResources []*string                                         `json:"openResources,omitempty" xml:"openResources,omitempty" type:"Repeated"`
	OpenRoleId    *string                                           `json:"openRoleId,omitempty" xml:"openRoleId,omitempty"`
	OpenRoleName  *string                                           `json:"openRoleName,omitempty" xml:"openRoleName,omitempty"`
}

func (s ListBasicRoleInPageResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListBasicRoleInPageResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListBasicRoleInPageResponseBodyList) SetOpenAction(v *ListBasicRoleInPageResponseBodyListOpenAction) *ListBasicRoleInPageResponseBodyList {
	s.OpenAction = v
	return s
}

func (s *ListBasicRoleInPageResponseBodyList) SetOpenMembers(v []*ListBasicRoleInPageResponseBodyListOpenMembers) *ListBasicRoleInPageResponseBodyList {
	s.OpenMembers = v
	return s
}

func (s *ListBasicRoleInPageResponseBodyList) SetOpenResources(v []*string) *ListBasicRoleInPageResponseBodyList {
	s.OpenResources = v
	return s
}

func (s *ListBasicRoleInPageResponseBodyList) SetOpenRoleId(v string) *ListBasicRoleInPageResponseBodyList {
	s.OpenRoleId = &v
	return s
}

func (s *ListBasicRoleInPageResponseBodyList) SetOpenRoleName(v string) *ListBasicRoleInPageResponseBodyList {
	s.OpenRoleName = &v
	return s
}

type ListBasicRoleInPageResponseBodyListOpenAction struct {
	ActionIds     []*string                                                   `json:"actionIds,omitempty" xml:"actionIds,omitempty" type:"Repeated"`
	OpenCondition *ListBasicRoleInPageResponseBodyListOpenActionOpenCondition `json:"openCondition,omitempty" xml:"openCondition,omitempty" type:"Struct"`
}

func (s ListBasicRoleInPageResponseBodyListOpenAction) String() string {
	return tea.Prettify(s)
}

func (s ListBasicRoleInPageResponseBodyListOpenAction) GoString() string {
	return s.String()
}

func (s *ListBasicRoleInPageResponseBodyListOpenAction) SetActionIds(v []*string) *ListBasicRoleInPageResponseBodyListOpenAction {
	s.ActionIds = v
	return s
}

func (s *ListBasicRoleInPageResponseBodyListOpenAction) SetOpenCondition(v *ListBasicRoleInPageResponseBodyListOpenActionOpenCondition) *ListBasicRoleInPageResponseBodyListOpenAction {
	s.OpenCondition = v
	return s
}

type ListBasicRoleInPageResponseBodyListOpenActionOpenCondition struct {
	OpenContactScope *ListBasicRoleInPageResponseBodyListOpenActionOpenConditionOpenContactScope `json:"openContactScope,omitempty" xml:"openContactScope,omitempty" type:"Struct"`
}

func (s ListBasicRoleInPageResponseBodyListOpenActionOpenCondition) String() string {
	return tea.Prettify(s)
}

func (s ListBasicRoleInPageResponseBodyListOpenActionOpenCondition) GoString() string {
	return s.String()
}

func (s *ListBasicRoleInPageResponseBodyListOpenActionOpenCondition) SetOpenContactScope(v *ListBasicRoleInPageResponseBodyListOpenActionOpenConditionOpenContactScope) *ListBasicRoleInPageResponseBodyListOpenActionOpenCondition {
	s.OpenContactScope = v
	return s
}

type ListBasicRoleInPageResponseBodyListOpenActionOpenConditionOpenContactScope struct {
	DeptIds                []*int64  `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	IncludeMemberDepts     *bool     `json:"includeMemberDepts,omitempty" xml:"includeMemberDepts,omitempty"`
	IncludeSelfManageDepts *bool     `json:"includeSelfManageDepts,omitempty" xml:"includeSelfManageDepts,omitempty"`
	UserIds                []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s ListBasicRoleInPageResponseBodyListOpenActionOpenConditionOpenContactScope) String() string {
	return tea.Prettify(s)
}

func (s ListBasicRoleInPageResponseBodyListOpenActionOpenConditionOpenContactScope) GoString() string {
	return s.String()
}

func (s *ListBasicRoleInPageResponseBodyListOpenActionOpenConditionOpenContactScope) SetDeptIds(v []*int64) *ListBasicRoleInPageResponseBodyListOpenActionOpenConditionOpenContactScope {
	s.DeptIds = v
	return s
}

func (s *ListBasicRoleInPageResponseBodyListOpenActionOpenConditionOpenContactScope) SetIncludeMemberDepts(v bool) *ListBasicRoleInPageResponseBodyListOpenActionOpenConditionOpenContactScope {
	s.IncludeMemberDepts = &v
	return s
}

func (s *ListBasicRoleInPageResponseBodyListOpenActionOpenConditionOpenContactScope) SetIncludeSelfManageDepts(v bool) *ListBasicRoleInPageResponseBodyListOpenActionOpenConditionOpenContactScope {
	s.IncludeSelfManageDepts = &v
	return s
}

func (s *ListBasicRoleInPageResponseBodyListOpenActionOpenConditionOpenContactScope) SetUserIds(v []*string) *ListBasicRoleInPageResponseBodyListOpenActionOpenConditionOpenContactScope {
	s.UserIds = v
	return s
}

type ListBasicRoleInPageResponseBodyListOpenMembers struct {
	BelongCorpId  *string `json:"belongCorpId,omitempty" xml:"belongCorpId,omitempty"`
	MemberId      *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	MemberType    *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	OperateUserId *string `json:"operateUserId,omitempty" xml:"operateUserId,omitempty"`
}

func (s ListBasicRoleInPageResponseBodyListOpenMembers) String() string {
	return tea.Prettify(s)
}

func (s ListBasicRoleInPageResponseBodyListOpenMembers) GoString() string {
	return s.String()
}

func (s *ListBasicRoleInPageResponseBodyListOpenMembers) SetBelongCorpId(v string) *ListBasicRoleInPageResponseBodyListOpenMembers {
	s.BelongCorpId = &v
	return s
}

func (s *ListBasicRoleInPageResponseBodyListOpenMembers) SetMemberId(v string) *ListBasicRoleInPageResponseBodyListOpenMembers {
	s.MemberId = &v
	return s
}

func (s *ListBasicRoleInPageResponseBodyListOpenMembers) SetMemberType(v string) *ListBasicRoleInPageResponseBodyListOpenMembers {
	s.MemberType = &v
	return s
}

func (s *ListBasicRoleInPageResponseBodyListOpenMembers) SetOperateUserId(v string) *ListBasicRoleInPageResponseBodyListOpenMembers {
	s.OperateUserId = &v
	return s
}

type ListBasicRoleInPageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListBasicRoleInPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBasicRoleInPageResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBasicRoleInPageResponse) GoString() string {
	return s.String()
}

func (s *ListBasicRoleInPageResponse) SetHeaders(v map[string]*string) *ListBasicRoleInPageResponse {
	s.Headers = v
	return s
}

func (s *ListBasicRoleInPageResponse) SetStatusCode(v int32) *ListBasicRoleInPageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBasicRoleInPageResponse) SetBody(v *ListBasicRoleInPageResponseBody) *ListBasicRoleInPageResponse {
	s.Body = v
	return s
}

type ListContactHideSettingsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListContactHideSettingsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListContactHideSettingsHeaders) GoString() string {
	return s.String()
}

func (s *ListContactHideSettingsHeaders) SetCommonHeaders(v map[string]*string) *ListContactHideSettingsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListContactHideSettingsHeaders) SetXAcsDingtalkAccessToken(v string) *ListContactHideSettingsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListContactHideSettingsRequest struct {
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListContactHideSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListContactHideSettingsRequest) GoString() string {
	return s.String()
}

func (s *ListContactHideSettingsRequest) SetMaxResults(v int32) *ListContactHideSettingsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListContactHideSettingsRequest) SetNextToken(v int64) *ListContactHideSettingsRequest {
	s.NextToken = &v
	return s
}

type ListContactHideSettingsResponseBody struct {
	HasMore   *bool                                      `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List      []*ListContactHideSettingsResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextToken *int64                                     `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListContactHideSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListContactHideSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListContactHideSettingsResponseBody) SetHasMore(v bool) *ListContactHideSettingsResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListContactHideSettingsResponseBody) SetList(v []*ListContactHideSettingsResponseBodyList) *ListContactHideSettingsResponseBody {
	s.List = v
	return s
}

func (s *ListContactHideSettingsResponseBody) SetNextToken(v int64) *ListContactHideSettingsResponseBody {
	s.NextToken = &v
	return s
}

type ListContactHideSettingsResponseBodyList struct {
	Active          *bool     `json:"active,omitempty" xml:"active,omitempty"`
	Description     *string   `json:"description,omitempty" xml:"description,omitempty"`
	ExcludeDeptIds  []*int64  `json:"excludeDeptIds,omitempty" xml:"excludeDeptIds,omitempty" type:"Repeated"`
	ExcludeStaffIds []*string `json:"excludeStaffIds,omitempty" xml:"excludeStaffIds,omitempty" type:"Repeated"`
	ExcludeTagIds   []*int64  `json:"excludeTagIds,omitempty" xml:"excludeTagIds,omitempty" type:"Repeated"`
	Id              *int64    `json:"id,omitempty" xml:"id,omitempty"`
	Name            *string   `json:"name,omitempty" xml:"name,omitempty"`
	ObjectDeptIds   []*int64  `json:"objectDeptIds,omitempty" xml:"objectDeptIds,omitempty" type:"Repeated"`
	ObjectStaffIds  []*string `json:"objectStaffIds,omitempty" xml:"objectStaffIds,omitempty" type:"Repeated"`
	ObjectTagIds    []*int64  `json:"objectTagIds,omitempty" xml:"objectTagIds,omitempty" type:"Repeated"`
}

func (s ListContactHideSettingsResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListContactHideSettingsResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListContactHideSettingsResponseBodyList) SetActive(v bool) *ListContactHideSettingsResponseBodyList {
	s.Active = &v
	return s
}

func (s *ListContactHideSettingsResponseBodyList) SetDescription(v string) *ListContactHideSettingsResponseBodyList {
	s.Description = &v
	return s
}

func (s *ListContactHideSettingsResponseBodyList) SetExcludeDeptIds(v []*int64) *ListContactHideSettingsResponseBodyList {
	s.ExcludeDeptIds = v
	return s
}

func (s *ListContactHideSettingsResponseBodyList) SetExcludeStaffIds(v []*string) *ListContactHideSettingsResponseBodyList {
	s.ExcludeStaffIds = v
	return s
}

func (s *ListContactHideSettingsResponseBodyList) SetExcludeTagIds(v []*int64) *ListContactHideSettingsResponseBodyList {
	s.ExcludeTagIds = v
	return s
}

func (s *ListContactHideSettingsResponseBodyList) SetId(v int64) *ListContactHideSettingsResponseBodyList {
	s.Id = &v
	return s
}

func (s *ListContactHideSettingsResponseBodyList) SetName(v string) *ListContactHideSettingsResponseBodyList {
	s.Name = &v
	return s
}

func (s *ListContactHideSettingsResponseBodyList) SetObjectDeptIds(v []*int64) *ListContactHideSettingsResponseBodyList {
	s.ObjectDeptIds = v
	return s
}

func (s *ListContactHideSettingsResponseBodyList) SetObjectStaffIds(v []*string) *ListContactHideSettingsResponseBodyList {
	s.ObjectStaffIds = v
	return s
}

func (s *ListContactHideSettingsResponseBodyList) SetObjectTagIds(v []*int64) *ListContactHideSettingsResponseBodyList {
	s.ObjectTagIds = v
	return s
}

type ListContactHideSettingsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListContactHideSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListContactHideSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListContactHideSettingsResponse) GoString() string {
	return s.String()
}

func (s *ListContactHideSettingsResponse) SetHeaders(v map[string]*string) *ListContactHideSettingsResponse {
	s.Headers = v
	return s
}

func (s *ListContactHideSettingsResponse) SetStatusCode(v int32) *ListContactHideSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContactHideSettingsResponse) SetBody(v *ListContactHideSettingsResponseBody) *ListContactHideSettingsResponse {
	s.Body = v
	return s
}

type ListContactRestrictSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListContactRestrictSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListContactRestrictSettingHeaders) GoString() string {
	return s.String()
}

func (s *ListContactRestrictSettingHeaders) SetCommonHeaders(v map[string]*string) *ListContactRestrictSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListContactRestrictSettingHeaders) SetXAcsDingtalkAccessToken(v string) *ListContactRestrictSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListContactRestrictSettingRequest struct {
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListContactRestrictSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s ListContactRestrictSettingRequest) GoString() string {
	return s.String()
}

func (s *ListContactRestrictSettingRequest) SetMaxResults(v int32) *ListContactRestrictSettingRequest {
	s.MaxResults = &v
	return s
}

func (s *ListContactRestrictSettingRequest) SetNextToken(v int64) *ListContactRestrictSettingRequest {
	s.NextToken = &v
	return s
}

type ListContactRestrictSettingResponseBody struct {
	HasMore   *bool                                         `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List      []*ListContactRestrictSettingResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextToken *int64                                        `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListContactRestrictSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListContactRestrictSettingResponseBody) GoString() string {
	return s.String()
}

func (s *ListContactRestrictSettingResponseBody) SetHasMore(v bool) *ListContactRestrictSettingResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListContactRestrictSettingResponseBody) SetList(v []*ListContactRestrictSettingResponseBodyList) *ListContactRestrictSettingResponseBody {
	s.List = v
	return s
}

func (s *ListContactRestrictSettingResponseBody) SetNextToken(v int64) *ListContactRestrictSettingResponseBody {
	s.NextToken = &v
	return s
}

type ListContactRestrictSettingResponseBodyList struct {
	Active                *bool     `json:"active,omitempty" xml:"active,omitempty"`
	Description           *string   `json:"description,omitempty" xml:"description,omitempty"`
	ExcludeDeptIds        []*int64  `json:"excludeDeptIds,omitempty" xml:"excludeDeptIds,omitempty" type:"Repeated"`
	ExcludeTagIds         []*int64  `json:"excludeTagIds,omitempty" xml:"excludeTagIds,omitempty" type:"Repeated"`
	ExcludeUserIds        []*string `json:"excludeUserIds,omitempty" xml:"excludeUserIds,omitempty" type:"Repeated"`
	Id                    *int64    `json:"id,omitempty" xml:"id,omitempty"`
	Name                  *string   `json:"name,omitempty" xml:"name,omitempty"`
	RestrictInSearch      *bool     `json:"restrictInSearch,omitempty" xml:"restrictInSearch,omitempty"`
	RestrictInUserProfile *bool     `json:"restrictInUserProfile,omitempty" xml:"restrictInUserProfile,omitempty"`
	SubjectDeptIds        []*int64  `json:"subjectDeptIds,omitempty" xml:"subjectDeptIds,omitempty" type:"Repeated"`
	SubjectTagIds         []*int64  `json:"subjectTagIds,omitempty" xml:"subjectTagIds,omitempty" type:"Repeated"`
	SubjectUserIds        []*string `json:"subjectUserIds,omitempty" xml:"subjectUserIds,omitempty" type:"Repeated"`
	Type                  *string   `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListContactRestrictSettingResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListContactRestrictSettingResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListContactRestrictSettingResponseBodyList) SetActive(v bool) *ListContactRestrictSettingResponseBodyList {
	s.Active = &v
	return s
}

func (s *ListContactRestrictSettingResponseBodyList) SetDescription(v string) *ListContactRestrictSettingResponseBodyList {
	s.Description = &v
	return s
}

func (s *ListContactRestrictSettingResponseBodyList) SetExcludeDeptIds(v []*int64) *ListContactRestrictSettingResponseBodyList {
	s.ExcludeDeptIds = v
	return s
}

func (s *ListContactRestrictSettingResponseBodyList) SetExcludeTagIds(v []*int64) *ListContactRestrictSettingResponseBodyList {
	s.ExcludeTagIds = v
	return s
}

func (s *ListContactRestrictSettingResponseBodyList) SetExcludeUserIds(v []*string) *ListContactRestrictSettingResponseBodyList {
	s.ExcludeUserIds = v
	return s
}

func (s *ListContactRestrictSettingResponseBodyList) SetId(v int64) *ListContactRestrictSettingResponseBodyList {
	s.Id = &v
	return s
}

func (s *ListContactRestrictSettingResponseBodyList) SetName(v string) *ListContactRestrictSettingResponseBodyList {
	s.Name = &v
	return s
}

func (s *ListContactRestrictSettingResponseBodyList) SetRestrictInSearch(v bool) *ListContactRestrictSettingResponseBodyList {
	s.RestrictInSearch = &v
	return s
}

func (s *ListContactRestrictSettingResponseBodyList) SetRestrictInUserProfile(v bool) *ListContactRestrictSettingResponseBodyList {
	s.RestrictInUserProfile = &v
	return s
}

func (s *ListContactRestrictSettingResponseBodyList) SetSubjectDeptIds(v []*int64) *ListContactRestrictSettingResponseBodyList {
	s.SubjectDeptIds = v
	return s
}

func (s *ListContactRestrictSettingResponseBodyList) SetSubjectTagIds(v []*int64) *ListContactRestrictSettingResponseBodyList {
	s.SubjectTagIds = v
	return s
}

func (s *ListContactRestrictSettingResponseBodyList) SetSubjectUserIds(v []*string) *ListContactRestrictSettingResponseBodyList {
	s.SubjectUserIds = v
	return s
}

func (s *ListContactRestrictSettingResponseBodyList) SetType(v string) *ListContactRestrictSettingResponseBodyList {
	s.Type = &v
	return s
}

type ListContactRestrictSettingResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListContactRestrictSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListContactRestrictSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s ListContactRestrictSettingResponse) GoString() string {
	return s.String()
}

func (s *ListContactRestrictSettingResponse) SetHeaders(v map[string]*string) *ListContactRestrictSettingResponse {
	s.Headers = v
	return s
}

func (s *ListContactRestrictSettingResponse) SetStatusCode(v int32) *ListContactRestrictSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContactRestrictSettingResponse) SetBody(v *ListContactRestrictSettingResponseBody) *ListContactRestrictSettingResponse {
	s.Body = v
	return s
}

type ListEmpAttributeVisibilityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListEmpAttributeVisibilityHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListEmpAttributeVisibilityHeaders) GoString() string {
	return s.String()
}

func (s *ListEmpAttributeVisibilityHeaders) SetCommonHeaders(v map[string]*string) *ListEmpAttributeVisibilityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListEmpAttributeVisibilityHeaders) SetXAcsDingtalkAccessToken(v string) *ListEmpAttributeVisibilityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListEmpAttributeVisibilityRequest struct {
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListEmpAttributeVisibilityRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEmpAttributeVisibilityRequest) GoString() string {
	return s.String()
}

func (s *ListEmpAttributeVisibilityRequest) SetMaxResults(v int32) *ListEmpAttributeVisibilityRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEmpAttributeVisibilityRequest) SetNextToken(v int64) *ListEmpAttributeVisibilityRequest {
	s.NextToken = &v
	return s
}

type ListEmpAttributeVisibilityResponseBody struct {
	HasMore    *bool                                         `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*ListEmpAttributeVisibilityResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextCursor *int64                                        `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
}

func (s ListEmpAttributeVisibilityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEmpAttributeVisibilityResponseBody) GoString() string {
	return s.String()
}

func (s *ListEmpAttributeVisibilityResponseBody) SetHasMore(v bool) *ListEmpAttributeVisibilityResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBody) SetList(v []*ListEmpAttributeVisibilityResponseBodyList) *ListEmpAttributeVisibilityResponseBody {
	s.List = v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBody) SetNextCursor(v int64) *ListEmpAttributeVisibilityResponseBody {
	s.NextCursor = &v
	return s
}

type ListEmpAttributeVisibilityResponseBodyList struct {
	Active          *bool     `json:"active,omitempty" xml:"active,omitempty"`
	Description     *string   `json:"description,omitempty" xml:"description,omitempty"`
	ExcludeDeptIds  []*int64  `json:"excludeDeptIds,omitempty" xml:"excludeDeptIds,omitempty" type:"Repeated"`
	ExcludeStaffIds []*string `json:"excludeStaffIds,omitempty" xml:"excludeStaffIds,omitempty" type:"Repeated"`
	ExcludeTagIds   []*int64  `json:"excludeTagIds,omitempty" xml:"excludeTagIds,omitempty" type:"Repeated"`
	GmtCreate       *string   `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified     *string   `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	HideFields      []*string `json:"hideFields,omitempty" xml:"hideFields,omitempty" type:"Repeated"`
	Id              *int64    `json:"id,omitempty" xml:"id,omitempty"`
	Name            *string   `json:"name,omitempty" xml:"name,omitempty"`
	ObjectDeptIds   []*int64  `json:"objectDeptIds,omitempty" xml:"objectDeptIds,omitempty" type:"Repeated"`
	ObjectStaffIds  []*string `json:"objectStaffIds,omitempty" xml:"objectStaffIds,omitempty" type:"Repeated"`
	ObjectTagIds    []*int64  `json:"objectTagIds,omitempty" xml:"objectTagIds,omitempty" type:"Repeated"`
}

func (s ListEmpAttributeVisibilityResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListEmpAttributeVisibilityResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetActive(v bool) *ListEmpAttributeVisibilityResponseBodyList {
	s.Active = &v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetDescription(v string) *ListEmpAttributeVisibilityResponseBodyList {
	s.Description = &v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetExcludeDeptIds(v []*int64) *ListEmpAttributeVisibilityResponseBodyList {
	s.ExcludeDeptIds = v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetExcludeStaffIds(v []*string) *ListEmpAttributeVisibilityResponseBodyList {
	s.ExcludeStaffIds = v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetExcludeTagIds(v []*int64) *ListEmpAttributeVisibilityResponseBodyList {
	s.ExcludeTagIds = v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetGmtCreate(v string) *ListEmpAttributeVisibilityResponseBodyList {
	s.GmtCreate = &v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetGmtModified(v string) *ListEmpAttributeVisibilityResponseBodyList {
	s.GmtModified = &v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetHideFields(v []*string) *ListEmpAttributeVisibilityResponseBodyList {
	s.HideFields = v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetId(v int64) *ListEmpAttributeVisibilityResponseBodyList {
	s.Id = &v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetName(v string) *ListEmpAttributeVisibilityResponseBodyList {
	s.Name = &v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetObjectDeptIds(v []*int64) *ListEmpAttributeVisibilityResponseBodyList {
	s.ObjectDeptIds = v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetObjectStaffIds(v []*string) *ListEmpAttributeVisibilityResponseBodyList {
	s.ObjectStaffIds = v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetObjectTagIds(v []*int64) *ListEmpAttributeVisibilityResponseBodyList {
	s.ObjectTagIds = v
	return s
}

type ListEmpAttributeVisibilityResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEmpAttributeVisibilityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEmpAttributeVisibilityResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEmpAttributeVisibilityResponse) GoString() string {
	return s.String()
}

func (s *ListEmpAttributeVisibilityResponse) SetHeaders(v map[string]*string) *ListEmpAttributeVisibilityResponse {
	s.Headers = v
	return s
}

func (s *ListEmpAttributeVisibilityResponse) SetStatusCode(v int32) *ListEmpAttributeVisibilityResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEmpAttributeVisibilityResponse) SetBody(v *ListEmpAttributeVisibilityResponseBody) *ListEmpAttributeVisibilityResponse {
	s.Body = v
	return s
}

type ListEmpLeaveRecordsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListEmpLeaveRecordsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListEmpLeaveRecordsHeaders) GoString() string {
	return s.String()
}

func (s *ListEmpLeaveRecordsHeaders) SetCommonHeaders(v map[string]*string) *ListEmpLeaveRecordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListEmpLeaveRecordsHeaders) SetXAcsDingtalkAccessToken(v string) *ListEmpLeaveRecordsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListEmpLeaveRecordsRequest struct {
	EndTime    *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	StartTime  *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListEmpLeaveRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEmpLeaveRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListEmpLeaveRecordsRequest) SetEndTime(v string) *ListEmpLeaveRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *ListEmpLeaveRecordsRequest) SetMaxResults(v int32) *ListEmpLeaveRecordsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEmpLeaveRecordsRequest) SetNextToken(v string) *ListEmpLeaveRecordsRequest {
	s.NextToken = &v
	return s
}

func (s *ListEmpLeaveRecordsRequest) SetStartTime(v string) *ListEmpLeaveRecordsRequest {
	s.StartTime = &v
	return s
}

type ListEmpLeaveRecordsResponseBody struct {
	NextToken *string                                   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Records   []*ListEmpLeaveRecordsResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
}

func (s ListEmpLeaveRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEmpLeaveRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEmpLeaveRecordsResponseBody) SetNextToken(v string) *ListEmpLeaveRecordsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListEmpLeaveRecordsResponseBody) SetRecords(v []*ListEmpLeaveRecordsResponseBodyRecords) *ListEmpLeaveRecordsResponseBody {
	s.Records = v
	return s
}

type ListEmpLeaveRecordsResponseBodyRecords struct {
	LeaveReason *string `json:"leaveReason,omitempty" xml:"leaveReason,omitempty"`
	LeaveTime   *string `json:"leaveTime,omitempty" xml:"leaveTime,omitempty"`
	Mobile      *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	StateCode   *string `json:"stateCode,omitempty" xml:"stateCode,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListEmpLeaveRecordsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ListEmpLeaveRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListEmpLeaveRecordsResponseBodyRecords) SetLeaveReason(v string) *ListEmpLeaveRecordsResponseBodyRecords {
	s.LeaveReason = &v
	return s
}

func (s *ListEmpLeaveRecordsResponseBodyRecords) SetLeaveTime(v string) *ListEmpLeaveRecordsResponseBodyRecords {
	s.LeaveTime = &v
	return s
}

func (s *ListEmpLeaveRecordsResponseBodyRecords) SetMobile(v string) *ListEmpLeaveRecordsResponseBodyRecords {
	s.Mobile = &v
	return s
}

func (s *ListEmpLeaveRecordsResponseBodyRecords) SetName(v string) *ListEmpLeaveRecordsResponseBodyRecords {
	s.Name = &v
	return s
}

func (s *ListEmpLeaveRecordsResponseBodyRecords) SetStateCode(v string) *ListEmpLeaveRecordsResponseBodyRecords {
	s.StateCode = &v
	return s
}

func (s *ListEmpLeaveRecordsResponseBodyRecords) SetUserId(v string) *ListEmpLeaveRecordsResponseBodyRecords {
	s.UserId = &v
	return s
}

type ListEmpLeaveRecordsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEmpLeaveRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEmpLeaveRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEmpLeaveRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListEmpLeaveRecordsResponse) SetHeaders(v map[string]*string) *ListEmpLeaveRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListEmpLeaveRecordsResponse) SetStatusCode(v int32) *ListEmpLeaveRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEmpLeaveRecordsResponse) SetBody(v *ListEmpLeaveRecordsResponseBody) *ListEmpLeaveRecordsResponse {
	s.Body = v
	return s
}

type ListManagementGroupsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListManagementGroupsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListManagementGroupsHeaders) GoString() string {
	return s.String()
}

func (s *ListManagementGroupsHeaders) SetCommonHeaders(v map[string]*string) *ListManagementGroupsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListManagementGroupsHeaders) SetXAcsDingtalkAccessToken(v string) *ListManagementGroupsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListManagementGroupsRequest struct {
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListManagementGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListManagementGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListManagementGroupsRequest) SetMaxResults(v int32) *ListManagementGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListManagementGroupsRequest) SetNextToken(v int64) *ListManagementGroupsRequest {
	s.NextToken = &v
	return s
}

type ListManagementGroupsResponseBody struct {
	Groups    []*ListManagementGroupsResponseBodyGroups `json:"groups,omitempty" xml:"groups,omitempty" type:"Repeated"`
	HasMore   *bool                                     `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken *int64                                    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListManagementGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListManagementGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListManagementGroupsResponseBody) SetGroups(v []*ListManagementGroupsResponseBodyGroups) *ListManagementGroupsResponseBody {
	s.Groups = v
	return s
}

func (s *ListManagementGroupsResponseBody) SetHasMore(v bool) *ListManagementGroupsResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListManagementGroupsResponseBody) SetNextToken(v int64) *ListManagementGroupsResponseBody {
	s.NextToken = &v
	return s
}

type ListManagementGroupsResponseBodyGroups struct {
	GroupId     *string                                          `json:"groupId,omitempty" xml:"groupId,omitempty"`
	GroupName   *string                                          `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Members     []*ListManagementGroupsResponseBodyGroupsMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	ResourceIds []*string                                        `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	Scope       *ListManagementGroupsResponseBodyGroupsScope     `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
}

func (s ListManagementGroupsResponseBodyGroups) String() string {
	return tea.Prettify(s)
}

func (s ListManagementGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListManagementGroupsResponseBodyGroups) SetGroupId(v string) *ListManagementGroupsResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *ListManagementGroupsResponseBodyGroups) SetGroupName(v string) *ListManagementGroupsResponseBodyGroups {
	s.GroupName = &v
	return s
}

func (s *ListManagementGroupsResponseBodyGroups) SetMembers(v []*ListManagementGroupsResponseBodyGroupsMembers) *ListManagementGroupsResponseBodyGroups {
	s.Members = v
	return s
}

func (s *ListManagementGroupsResponseBodyGroups) SetResourceIds(v []*string) *ListManagementGroupsResponseBodyGroups {
	s.ResourceIds = v
	return s
}

func (s *ListManagementGroupsResponseBodyGroups) SetScope(v *ListManagementGroupsResponseBodyGroupsScope) *ListManagementGroupsResponseBodyGroups {
	s.Scope = v
	return s
}

type ListManagementGroupsResponseBodyGroupsMembers struct {
	MemberId   *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
}

func (s ListManagementGroupsResponseBodyGroupsMembers) String() string {
	return tea.Prettify(s)
}

func (s ListManagementGroupsResponseBodyGroupsMembers) GoString() string {
	return s.String()
}

func (s *ListManagementGroupsResponseBodyGroupsMembers) SetMemberId(v string) *ListManagementGroupsResponseBodyGroupsMembers {
	s.MemberId = &v
	return s
}

func (s *ListManagementGroupsResponseBodyGroupsMembers) SetMemberType(v string) *ListManagementGroupsResponseBodyGroupsMembers {
	s.MemberType = &v
	return s
}

type ListManagementGroupsResponseBodyGroupsScope struct {
	DeptIds   []*int64 `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	ScopeType *int32   `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
}

func (s ListManagementGroupsResponseBodyGroupsScope) String() string {
	return tea.Prettify(s)
}

func (s ListManagementGroupsResponseBodyGroupsScope) GoString() string {
	return s.String()
}

func (s *ListManagementGroupsResponseBodyGroupsScope) SetDeptIds(v []*int64) *ListManagementGroupsResponseBodyGroupsScope {
	s.DeptIds = v
	return s
}

func (s *ListManagementGroupsResponseBodyGroupsScope) SetScopeType(v int32) *ListManagementGroupsResponseBodyGroupsScope {
	s.ScopeType = &v
	return s
}

type ListManagementGroupsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListManagementGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListManagementGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListManagementGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListManagementGroupsResponse) SetHeaders(v map[string]*string) *ListManagementGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListManagementGroupsResponse) SetStatusCode(v int32) *ListManagementGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListManagementGroupsResponse) SetBody(v *ListManagementGroupsResponseBody) *ListManagementGroupsResponse {
	s.Body = v
	return s
}

type ListOwnedOrgByStaffIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListOwnedOrgByStaffIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListOwnedOrgByStaffIdHeaders) GoString() string {
	return s.String()
}

func (s *ListOwnedOrgByStaffIdHeaders) SetCommonHeaders(v map[string]*string) *ListOwnedOrgByStaffIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListOwnedOrgByStaffIdHeaders) SetXAcsDingtalkAccessToken(v string) *ListOwnedOrgByStaffIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListOwnedOrgByStaffIdRequest struct {
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListOwnedOrgByStaffIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOwnedOrgByStaffIdRequest) GoString() string {
	return s.String()
}

func (s *ListOwnedOrgByStaffIdRequest) SetUserId(v string) *ListOwnedOrgByStaffIdRequest {
	s.UserId = &v
	return s
}

type ListOwnedOrgByStaffIdResponseBody struct {
	OrgList []*ListOwnedOrgByStaffIdResponseBodyOrgList `json:"orgList,omitempty" xml:"orgList,omitempty" type:"Repeated"`
}

func (s ListOwnedOrgByStaffIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOwnedOrgByStaffIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListOwnedOrgByStaffIdResponseBody) SetOrgList(v []*ListOwnedOrgByStaffIdResponseBodyOrgList) *ListOwnedOrgByStaffIdResponseBody {
	s.OrgList = v
	return s
}

type ListOwnedOrgByStaffIdResponseBodyOrgList struct {
	CorpId   *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	CorpName *string `json:"corpName,omitempty" xml:"corpName,omitempty"`
}

func (s ListOwnedOrgByStaffIdResponseBodyOrgList) String() string {
	return tea.Prettify(s)
}

func (s ListOwnedOrgByStaffIdResponseBodyOrgList) GoString() string {
	return s.String()
}

func (s *ListOwnedOrgByStaffIdResponseBodyOrgList) SetCorpId(v string) *ListOwnedOrgByStaffIdResponseBodyOrgList {
	s.CorpId = &v
	return s
}

func (s *ListOwnedOrgByStaffIdResponseBodyOrgList) SetCorpName(v string) *ListOwnedOrgByStaffIdResponseBodyOrgList {
	s.CorpName = &v
	return s
}

type ListOwnedOrgByStaffIdResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListOwnedOrgByStaffIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOwnedOrgByStaffIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOwnedOrgByStaffIdResponse) GoString() string {
	return s.String()
}

func (s *ListOwnedOrgByStaffIdResponse) SetHeaders(v map[string]*string) *ListOwnedOrgByStaffIdResponse {
	s.Headers = v
	return s
}

func (s *ListOwnedOrgByStaffIdResponse) SetStatusCode(v int32) *ListOwnedOrgByStaffIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOwnedOrgByStaffIdResponse) SetBody(v *ListOwnedOrgByStaffIdResponseBody) *ListOwnedOrgByStaffIdResponse {
	s.Body = v
	return s
}

type ListSeniorSettingsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListSeniorSettingsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSeniorSettingsHeaders) GoString() string {
	return s.String()
}

func (s *ListSeniorSettingsHeaders) SetCommonHeaders(v map[string]*string) *ListSeniorSettingsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSeniorSettingsHeaders) SetXAcsDingtalkAccessToken(v string) *ListSeniorSettingsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListSeniorSettingsRequest struct {
	SeniorStaffId *string `json:"seniorStaffId,omitempty" xml:"seniorStaffId,omitempty"`
}

func (s ListSeniorSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSeniorSettingsRequest) GoString() string {
	return s.String()
}

func (s *ListSeniorSettingsRequest) SetSeniorStaffId(v string) *ListSeniorSettingsRequest {
	s.SeniorStaffId = &v
	return s
}

type ListSeniorSettingsResponseBody struct {
	ProtectScenes   []*string                                        `json:"protectScenes,omitempty" xml:"protectScenes,omitempty" type:"Repeated"`
	SeniorStaffId   *string                                          `json:"seniorStaffId,omitempty" xml:"seniorStaffId,omitempty"`
	SeniorWhiteList []*ListSeniorSettingsResponseBodySeniorWhiteList `json:"seniorWhiteList,omitempty" xml:"seniorWhiteList,omitempty" type:"Repeated"`
}

func (s ListSeniorSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSeniorSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSeniorSettingsResponseBody) SetProtectScenes(v []*string) *ListSeniorSettingsResponseBody {
	s.ProtectScenes = v
	return s
}

func (s *ListSeniorSettingsResponseBody) SetSeniorStaffId(v string) *ListSeniorSettingsResponseBody {
	s.SeniorStaffId = &v
	return s
}

func (s *ListSeniorSettingsResponseBody) SetSeniorWhiteList(v []*ListSeniorSettingsResponseBodySeniorWhiteList) *ListSeniorSettingsResponseBody {
	s.SeniorWhiteList = v
	return s
}

type ListSeniorSettingsResponseBodySeniorWhiteList struct {
	Id   *string `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Type *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListSeniorSettingsResponseBodySeniorWhiteList) String() string {
	return tea.Prettify(s)
}

func (s ListSeniorSettingsResponseBodySeniorWhiteList) GoString() string {
	return s.String()
}

func (s *ListSeniorSettingsResponseBodySeniorWhiteList) SetId(v string) *ListSeniorSettingsResponseBodySeniorWhiteList {
	s.Id = &v
	return s
}

func (s *ListSeniorSettingsResponseBodySeniorWhiteList) SetName(v string) *ListSeniorSettingsResponseBodySeniorWhiteList {
	s.Name = &v
	return s
}

func (s *ListSeniorSettingsResponseBodySeniorWhiteList) SetType(v int32) *ListSeniorSettingsResponseBodySeniorWhiteList {
	s.Type = &v
	return s
}

type ListSeniorSettingsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSeniorSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSeniorSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSeniorSettingsResponse) GoString() string {
	return s.String()
}

func (s *ListSeniorSettingsResponse) SetHeaders(v map[string]*string) *ListSeniorSettingsResponse {
	s.Headers = v
	return s
}

func (s *ListSeniorSettingsResponse) SetStatusCode(v int32) *ListSeniorSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSeniorSettingsResponse) SetBody(v *ListSeniorSettingsResponseBody) *ListSeniorSettingsResponse {
	s.Body = v
	return s
}

type MultiOrgPermissionGrantHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MultiOrgPermissionGrantHeaders) String() string {
	return tea.Prettify(s)
}

func (s MultiOrgPermissionGrantHeaders) GoString() string {
	return s.String()
}

func (s *MultiOrgPermissionGrantHeaders) SetCommonHeaders(v map[string]*string) *MultiOrgPermissionGrantHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MultiOrgPermissionGrantHeaders) SetXAcsDingtalkAccessToken(v string) *MultiOrgPermissionGrantHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MultiOrgPermissionGrantRequest struct {
	GrantDeptIdList []*int64 `json:"grantDeptIdList,omitempty" xml:"grantDeptIdList,omitempty" type:"Repeated"`
	JoinCorpId      *string  `json:"joinCorpId,omitempty" xml:"joinCorpId,omitempty"`
}

func (s MultiOrgPermissionGrantRequest) String() string {
	return tea.Prettify(s)
}

func (s MultiOrgPermissionGrantRequest) GoString() string {
	return s.String()
}

func (s *MultiOrgPermissionGrantRequest) SetGrantDeptIdList(v []*int64) *MultiOrgPermissionGrantRequest {
	s.GrantDeptIdList = v
	return s
}

func (s *MultiOrgPermissionGrantRequest) SetJoinCorpId(v string) *MultiOrgPermissionGrantRequest {
	s.JoinCorpId = &v
	return s
}

type MultiOrgPermissionGrantResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s MultiOrgPermissionGrantResponse) String() string {
	return tea.Prettify(s)
}

func (s MultiOrgPermissionGrantResponse) GoString() string {
	return s.String()
}

func (s *MultiOrgPermissionGrantResponse) SetHeaders(v map[string]*string) *MultiOrgPermissionGrantResponse {
	s.Headers = v
	return s
}

func (s *MultiOrgPermissionGrantResponse) SetStatusCode(v int32) *MultiOrgPermissionGrantResponse {
	s.StatusCode = &v
	return s
}

type QueryCardVisitorStatisticDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCardVisitorStatisticDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCardVisitorStatisticDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryCardVisitorStatisticDataHeaders) SetCommonHeaders(v map[string]*string) *QueryCardVisitorStatisticDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCardVisitorStatisticDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCardVisitorStatisticDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCardVisitorStatisticDataRequest struct {
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryCardVisitorStatisticDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCardVisitorStatisticDataRequest) GoString() string {
	return s.String()
}

func (s *QueryCardVisitorStatisticDataRequest) SetUnionId(v string) *QueryCardVisitorStatisticDataRequest {
	s.UnionId = &v
	return s
}

type QueryCardVisitorStatisticDataResponseBody struct {
	CardSendCnt            *int64 `json:"cardSendCnt,omitempty" xml:"cardSendCnt,omitempty"`
	TodayVisitAddCnt       *int64 `json:"todayVisitAddCnt,omitempty" xml:"todayVisitAddCnt,omitempty"`
	TodayVisitCnt          *int64 `json:"todayVisitCnt,omitempty" xml:"todayVisitCnt,omitempty"`
	TotalVisitAddCnt       *int64 `json:"totalVisitAddCnt,omitempty" xml:"totalVisitAddCnt,omitempty"`
	TotalVisitCnt          *int64 `json:"totalVisitCnt,omitempty" xml:"totalVisitCnt,omitempty"`
	WechatTodayVisitAddCnt *int64 `json:"wechatTodayVisitAddCnt,omitempty" xml:"wechatTodayVisitAddCnt,omitempty"`
	WechatTodayVisitCnt    *int64 `json:"wechatTodayVisitCnt,omitempty" xml:"wechatTodayVisitCnt,omitempty"`
	WechatTotalVisitAddCnt *int64 `json:"wechatTotalVisitAddCnt,omitempty" xml:"wechatTotalVisitAddCnt,omitempty"`
	WechatTotalVisitCnt    *int64 `json:"wechatTotalVisitCnt,omitempty" xml:"wechatTotalVisitCnt,omitempty"`
}

func (s QueryCardVisitorStatisticDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCardVisitorStatisticDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCardVisitorStatisticDataResponseBody) SetCardSendCnt(v int64) *QueryCardVisitorStatisticDataResponseBody {
	s.CardSendCnt = &v
	return s
}

func (s *QueryCardVisitorStatisticDataResponseBody) SetTodayVisitAddCnt(v int64) *QueryCardVisitorStatisticDataResponseBody {
	s.TodayVisitAddCnt = &v
	return s
}

func (s *QueryCardVisitorStatisticDataResponseBody) SetTodayVisitCnt(v int64) *QueryCardVisitorStatisticDataResponseBody {
	s.TodayVisitCnt = &v
	return s
}

func (s *QueryCardVisitorStatisticDataResponseBody) SetTotalVisitAddCnt(v int64) *QueryCardVisitorStatisticDataResponseBody {
	s.TotalVisitAddCnt = &v
	return s
}

func (s *QueryCardVisitorStatisticDataResponseBody) SetTotalVisitCnt(v int64) *QueryCardVisitorStatisticDataResponseBody {
	s.TotalVisitCnt = &v
	return s
}

func (s *QueryCardVisitorStatisticDataResponseBody) SetWechatTodayVisitAddCnt(v int64) *QueryCardVisitorStatisticDataResponseBody {
	s.WechatTodayVisitAddCnt = &v
	return s
}

func (s *QueryCardVisitorStatisticDataResponseBody) SetWechatTodayVisitCnt(v int64) *QueryCardVisitorStatisticDataResponseBody {
	s.WechatTodayVisitCnt = &v
	return s
}

func (s *QueryCardVisitorStatisticDataResponseBody) SetWechatTotalVisitAddCnt(v int64) *QueryCardVisitorStatisticDataResponseBody {
	s.WechatTotalVisitAddCnt = &v
	return s
}

func (s *QueryCardVisitorStatisticDataResponseBody) SetWechatTotalVisitCnt(v int64) *QueryCardVisitorStatisticDataResponseBody {
	s.WechatTotalVisitCnt = &v
	return s
}

type QueryCardVisitorStatisticDataResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCardVisitorStatisticDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCardVisitorStatisticDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCardVisitorStatisticDataResponse) GoString() string {
	return s.String()
}

func (s *QueryCardVisitorStatisticDataResponse) SetHeaders(v map[string]*string) *QueryCardVisitorStatisticDataResponse {
	s.Headers = v
	return s
}

func (s *QueryCardVisitorStatisticDataResponse) SetStatusCode(v int32) *QueryCardVisitorStatisticDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCardVisitorStatisticDataResponse) SetBody(v *QueryCardVisitorStatisticDataResponseBody) *QueryCardVisitorStatisticDataResponse {
	s.Body = v
	return s
}

type QueryCorpStatisticDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCorpStatisticDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCorpStatisticDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryCorpStatisticDataHeaders) SetCommonHeaders(v map[string]*string) *QueryCorpStatisticDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCorpStatisticDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCorpStatisticDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCorpStatisticDataRequest struct {
	EndTime     *string   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime   *string   `json:"startTime,omitempty" xml:"startTime,omitempty"`
	TemplateIds []*string `json:"templateIds,omitempty" xml:"templateIds,omitempty" type:"Repeated"`
	UnionId     *string   `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryCorpStatisticDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCorpStatisticDataRequest) GoString() string {
	return s.String()
}

func (s *QueryCorpStatisticDataRequest) SetEndTime(v string) *QueryCorpStatisticDataRequest {
	s.EndTime = &v
	return s
}

func (s *QueryCorpStatisticDataRequest) SetStartTime(v string) *QueryCorpStatisticDataRequest {
	s.StartTime = &v
	return s
}

func (s *QueryCorpStatisticDataRequest) SetTemplateIds(v []*string) *QueryCorpStatisticDataRequest {
	s.TemplateIds = v
	return s
}

func (s *QueryCorpStatisticDataRequest) SetUnionId(v string) *QueryCorpStatisticDataRequest {
	s.UnionId = &v
	return s
}

type QueryCorpStatisticDataResponseBody struct {
	Result []*QueryCorpStatisticDataResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryCorpStatisticDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCorpStatisticDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCorpStatisticDataResponseBody) SetResult(v []*QueryCorpStatisticDataResponseBodyResult) *QueryCorpStatisticDataResponseBody {
	s.Result = v
	return s
}

type QueryCorpStatisticDataResponseBodyResult struct {
	CardBeReceivedTotalCnt *int64  `json:"cardBeReceivedTotalCnt,omitempty" xml:"cardBeReceivedTotalCnt,omitempty"`
	CardReceiveTotalCnt    *int64  `json:"cardReceiveTotalCnt,omitempty" xml:"cardReceiveTotalCnt,omitempty"`
	CardTotalBeVisitedCnt  *int64  `json:"cardTotalBeVisitedCnt,omitempty" xml:"cardTotalBeVisitedCnt,omitempty"`
	DataDate               *string `json:"dataDate,omitempty" xml:"dataDate,omitempty"`
	DingTotalShareCnt      *int64  `json:"dingTotalShareCnt,omitempty" xml:"dingTotalShareCnt,omitempty"`
	TotalSendCnt           *int64  `json:"totalSendCnt,omitempty" xml:"totalSendCnt,omitempty"`
	WechatTotalShareCnt    *int64  `json:"wechatTotalShareCnt,omitempty" xml:"wechatTotalShareCnt,omitempty"`
}

func (s QueryCorpStatisticDataResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryCorpStatisticDataResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryCorpStatisticDataResponseBodyResult) SetCardBeReceivedTotalCnt(v int64) *QueryCorpStatisticDataResponseBodyResult {
	s.CardBeReceivedTotalCnt = &v
	return s
}

func (s *QueryCorpStatisticDataResponseBodyResult) SetCardReceiveTotalCnt(v int64) *QueryCorpStatisticDataResponseBodyResult {
	s.CardReceiveTotalCnt = &v
	return s
}

func (s *QueryCorpStatisticDataResponseBodyResult) SetCardTotalBeVisitedCnt(v int64) *QueryCorpStatisticDataResponseBodyResult {
	s.CardTotalBeVisitedCnt = &v
	return s
}

func (s *QueryCorpStatisticDataResponseBodyResult) SetDataDate(v string) *QueryCorpStatisticDataResponseBodyResult {
	s.DataDate = &v
	return s
}

func (s *QueryCorpStatisticDataResponseBodyResult) SetDingTotalShareCnt(v int64) *QueryCorpStatisticDataResponseBodyResult {
	s.DingTotalShareCnt = &v
	return s
}

func (s *QueryCorpStatisticDataResponseBodyResult) SetTotalSendCnt(v int64) *QueryCorpStatisticDataResponseBodyResult {
	s.TotalSendCnt = &v
	return s
}

func (s *QueryCorpStatisticDataResponseBodyResult) SetWechatTotalShareCnt(v int64) *QueryCorpStatisticDataResponseBodyResult {
	s.WechatTotalShareCnt = &v
	return s
}

type QueryCorpStatisticDataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCorpStatisticDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCorpStatisticDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCorpStatisticDataResponse) GoString() string {
	return s.String()
}

func (s *QueryCorpStatisticDataResponse) SetHeaders(v map[string]*string) *QueryCorpStatisticDataResponse {
	s.Headers = v
	return s
}

func (s *QueryCorpStatisticDataResponse) SetStatusCode(v int32) *QueryCorpStatisticDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCorpStatisticDataResponse) SetBody(v *QueryCorpStatisticDataResponseBody) *QueryCorpStatisticDataResponse {
	s.Body = v
	return s
}

type QueryCorpUserStatisticHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCorpUserStatisticHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCorpUserStatisticHeaders) GoString() string {
	return s.String()
}

func (s *QueryCorpUserStatisticHeaders) SetCommonHeaders(v map[string]*string) *QueryCorpUserStatisticHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCorpUserStatisticHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCorpUserStatisticHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCorpUserStatisticRequest struct {
	EndTime     *string   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	MaxResults  *int64    `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken   *int64    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	StartTime   *string   `json:"startTime,omitempty" xml:"startTime,omitempty"`
	TemplateIds []*string `json:"templateIds,omitempty" xml:"templateIds,omitempty" type:"Repeated"`
	UnionId     *string   `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryCorpUserStatisticRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCorpUserStatisticRequest) GoString() string {
	return s.String()
}

func (s *QueryCorpUserStatisticRequest) SetEndTime(v string) *QueryCorpUserStatisticRequest {
	s.EndTime = &v
	return s
}

func (s *QueryCorpUserStatisticRequest) SetMaxResults(v int64) *QueryCorpUserStatisticRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryCorpUserStatisticRequest) SetNextToken(v int64) *QueryCorpUserStatisticRequest {
	s.NextToken = &v
	return s
}

func (s *QueryCorpUserStatisticRequest) SetStartTime(v string) *QueryCorpUserStatisticRequest {
	s.StartTime = &v
	return s
}

func (s *QueryCorpUserStatisticRequest) SetTemplateIds(v []*string) *QueryCorpUserStatisticRequest {
	s.TemplateIds = v
	return s
}

func (s *QueryCorpUserStatisticRequest) SetUnionId(v string) *QueryCorpUserStatisticRequest {
	s.UnionId = &v
	return s
}

type QueryCorpUserStatisticResponseBody struct {
	HasMore    *bool                                     `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*QueryCorpUserStatisticResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextToken  *int64                                    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	TotalCount *int64                                    `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryCorpUserStatisticResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCorpUserStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCorpUserStatisticResponseBody) SetHasMore(v bool) *QueryCorpUserStatisticResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryCorpUserStatisticResponseBody) SetList(v []*QueryCorpUserStatisticResponseBodyList) *QueryCorpUserStatisticResponseBody {
	s.List = v
	return s
}

func (s *QueryCorpUserStatisticResponseBody) SetNextToken(v int64) *QueryCorpUserStatisticResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryCorpUserStatisticResponseBody) SetTotalCount(v int64) *QueryCorpUserStatisticResponseBody {
	s.TotalCount = &v
	return s
}

type QueryCorpUserStatisticResponseBodyList struct {
	AvatarUrl  *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	ReceiveCnt *int64  `json:"receiveCnt,omitempty" xml:"receiveCnt,omitempty"`
	SendCnt    *int64  `json:"sendCnt,omitempty" xml:"sendCnt,omitempty"`
	UnionId    *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryCorpUserStatisticResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryCorpUserStatisticResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryCorpUserStatisticResponseBodyList) SetAvatarUrl(v string) *QueryCorpUserStatisticResponseBodyList {
	s.AvatarUrl = &v
	return s
}

func (s *QueryCorpUserStatisticResponseBodyList) SetName(v string) *QueryCorpUserStatisticResponseBodyList {
	s.Name = &v
	return s
}

func (s *QueryCorpUserStatisticResponseBodyList) SetReceiveCnt(v int64) *QueryCorpUserStatisticResponseBodyList {
	s.ReceiveCnt = &v
	return s
}

func (s *QueryCorpUserStatisticResponseBodyList) SetSendCnt(v int64) *QueryCorpUserStatisticResponseBodyList {
	s.SendCnt = &v
	return s
}

func (s *QueryCorpUserStatisticResponseBodyList) SetUnionId(v string) *QueryCorpUserStatisticResponseBodyList {
	s.UnionId = &v
	return s
}

type QueryCorpUserStatisticResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCorpUserStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCorpUserStatisticResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCorpUserStatisticResponse) GoString() string {
	return s.String()
}

func (s *QueryCorpUserStatisticResponse) SetHeaders(v map[string]*string) *QueryCorpUserStatisticResponse {
	s.Headers = v
	return s
}

func (s *QueryCorpUserStatisticResponse) SetStatusCode(v int32) *QueryCorpUserStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCorpUserStatisticResponse) SetBody(v *QueryCorpUserStatisticResponseBody) *QueryCorpUserStatisticResponse {
	s.Body = v
	return s
}

type QueryResourceManagementMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryResourceManagementMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryResourceManagementMembersHeaders) GoString() string {
	return s.String()
}

func (s *QueryResourceManagementMembersHeaders) SetCommonHeaders(v map[string]*string) *QueryResourceManagementMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryResourceManagementMembersHeaders) SetXAcsDingtalkAccessToken(v string) *QueryResourceManagementMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryResourceManagementMembersResponseBody struct {
	Members []*QueryResourceManagementMembersResponseBodyMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
}

func (s QueryResourceManagementMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryResourceManagementMembersResponseBody) GoString() string {
	return s.String()
}

func (s *QueryResourceManagementMembersResponseBody) SetMembers(v []*QueryResourceManagementMembersResponseBodyMembers) *QueryResourceManagementMembersResponseBody {
	s.Members = v
	return s
}

type QueryResourceManagementMembersResponseBodyMembers struct {
	MemberId   *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
}

func (s QueryResourceManagementMembersResponseBodyMembers) String() string {
	return tea.Prettify(s)
}

func (s QueryResourceManagementMembersResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *QueryResourceManagementMembersResponseBodyMembers) SetMemberId(v string) *QueryResourceManagementMembersResponseBodyMembers {
	s.MemberId = &v
	return s
}

func (s *QueryResourceManagementMembersResponseBodyMembers) SetMemberType(v string) *QueryResourceManagementMembersResponseBodyMembers {
	s.MemberType = &v
	return s
}

type QueryResourceManagementMembersResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryResourceManagementMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryResourceManagementMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryResourceManagementMembersResponse) GoString() string {
	return s.String()
}

func (s *QueryResourceManagementMembersResponse) SetHeaders(v map[string]*string) *QueryResourceManagementMembersResponse {
	s.Headers = v
	return s
}

func (s *QueryResourceManagementMembersResponse) SetStatusCode(v int32) *QueryResourceManagementMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryResourceManagementMembersResponse) SetBody(v *QueryResourceManagementMembersResponseBody) *QueryResourceManagementMembersResponse {
	s.Body = v
	return s
}

type QueryStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryStatusHeaders) GoString() string {
	return s.String()
}

func (s *QueryStatusHeaders) SetCommonHeaders(v map[string]*string) *QueryStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryStatusHeaders) SetXAcsDingtalkAccessToken(v string) *QueryStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryStatusRequest struct {
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryStatusRequest) SetUserId(v string) *QueryStatusRequest {
	s.UserId = &v
	return s
}

type QueryStatusResponseBody struct {
	Disable *bool `json:"disable,omitempty" xml:"disable,omitempty"`
}

func (s QueryStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryStatusResponseBody) SetDisable(v bool) *QueryStatusResponseBody {
	s.Disable = &v
	return s
}

type QueryStatusResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryStatusResponse) SetHeaders(v map[string]*string) *QueryStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryStatusResponse) SetStatusCode(v int32) *QueryStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryStatusResponse) SetBody(v *QueryStatusResponseBody) *QueryStatusResponse {
	s.Body = v
	return s
}

type QueryUserManagementResourcesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserManagementResourcesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserManagementResourcesHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserManagementResourcesHeaders) SetCommonHeaders(v map[string]*string) *QueryUserManagementResourcesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserManagementResourcesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserManagementResourcesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserManagementResourcesResponseBody struct {
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
}

func (s QueryUserManagementResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserManagementResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserManagementResourcesResponseBody) SetResourceIds(v []*string) *QueryUserManagementResourcesResponseBody {
	s.ResourceIds = v
	return s
}

type QueryUserManagementResourcesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUserManagementResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserManagementResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserManagementResourcesResponse) GoString() string {
	return s.String()
}

func (s *QueryUserManagementResourcesResponse) SetHeaders(v map[string]*string) *QueryUserManagementResourcesResponse {
	s.Headers = v
	return s
}

func (s *QueryUserManagementResourcesResponse) SetStatusCode(v int32) *QueryUserManagementResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserManagementResourcesResponse) SetBody(v *QueryUserManagementResourcesResponseBody) *QueryUserManagementResourcesResponse {
	s.Body = v
	return s
}

type QueryVerifyResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryVerifyResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryVerifyResultHeaders) GoString() string {
	return s.String()
}

func (s *QueryVerifyResultHeaders) SetCommonHeaders(v map[string]*string) *QueryVerifyResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryVerifyResultHeaders) SetXAcsDingtalkAccessToken(v string) *QueryVerifyResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryVerifyResultRequest struct {
	VerifyId *string `json:"verifyId,omitempty" xml:"verifyId,omitempty"`
}

func (s QueryVerifyResultRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryVerifyResultRequest) GoString() string {
	return s.String()
}

func (s *QueryVerifyResultRequest) SetVerifyId(v string) *QueryVerifyResultRequest {
	s.VerifyId = &v
	return s
}

type QueryVerifyResultResponseBody struct {
	CorpId          *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	FactorCode      *string `json:"factorCode,omitempty" xml:"factorCode,omitempty"`
	FactorDesc      *string `json:"factorDesc,omitempty" xml:"factorDesc,omitempty"`
	ResultCode      *string `json:"resultCode,omitempty" xml:"resultCode,omitempty"`
	ResultDesc      *string `json:"resultDesc,omitempty" xml:"resultDesc,omitempty"`
	State           *string `json:"state,omitempty" xml:"state,omitempty"`
	UserId          *string `json:"userId,omitempty" xml:"userId,omitempty"`
	VerifyTimestamp *int64  `json:"verifyTimestamp,omitempty" xml:"verifyTimestamp,omitempty"`
}

func (s QueryVerifyResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryVerifyResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVerifyResultResponseBody) SetCorpId(v string) *QueryVerifyResultResponseBody {
	s.CorpId = &v
	return s
}

func (s *QueryVerifyResultResponseBody) SetFactorCode(v string) *QueryVerifyResultResponseBody {
	s.FactorCode = &v
	return s
}

func (s *QueryVerifyResultResponseBody) SetFactorDesc(v string) *QueryVerifyResultResponseBody {
	s.FactorDesc = &v
	return s
}

func (s *QueryVerifyResultResponseBody) SetResultCode(v string) *QueryVerifyResultResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryVerifyResultResponseBody) SetResultDesc(v string) *QueryVerifyResultResponseBody {
	s.ResultDesc = &v
	return s
}

func (s *QueryVerifyResultResponseBody) SetState(v string) *QueryVerifyResultResponseBody {
	s.State = &v
	return s
}

func (s *QueryVerifyResultResponseBody) SetUserId(v string) *QueryVerifyResultResponseBody {
	s.UserId = &v
	return s
}

func (s *QueryVerifyResultResponseBody) SetVerifyTimestamp(v int64) *QueryVerifyResultResponseBody {
	s.VerifyTimestamp = &v
	return s
}

type QueryVerifyResultResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryVerifyResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryVerifyResultResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryVerifyResultResponse) GoString() string {
	return s.String()
}

func (s *QueryVerifyResultResponse) SetHeaders(v map[string]*string) *QueryVerifyResultResponse {
	s.Headers = v
	return s
}

func (s *QueryVerifyResultResponse) SetStatusCode(v int32) *QueryVerifyResultResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVerifyResultResponse) SetBody(v *QueryVerifyResultResponseBody) *QueryVerifyResultResponse {
	s.Body = v
	return s
}

type SearchDepartmentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchDepartmentHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchDepartmentHeaders) GoString() string {
	return s.String()
}

func (s *SearchDepartmentHeaders) SetCommonHeaders(v map[string]*string) *SearchDepartmentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchDepartmentHeaders) SetXAcsDingtalkAccessToken(v string) *SearchDepartmentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchDepartmentRequest struct {
	Offset    *int32  `json:"offset,omitempty" xml:"offset,omitempty"`
	QueryWord *string `json:"queryWord,omitempty" xml:"queryWord,omitempty"`
	Size      *int32  `json:"size,omitempty" xml:"size,omitempty"`
}

func (s SearchDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchDepartmentRequest) GoString() string {
	return s.String()
}

func (s *SearchDepartmentRequest) SetOffset(v int32) *SearchDepartmentRequest {
	s.Offset = &v
	return s
}

func (s *SearchDepartmentRequest) SetQueryWord(v string) *SearchDepartmentRequest {
	s.QueryWord = &v
	return s
}

func (s *SearchDepartmentRequest) SetSize(v int32) *SearchDepartmentRequest {
	s.Size = &v
	return s
}

type SearchDepartmentResponseBody struct {
	HasMore    *bool    `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*int64 `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	TotalCount *int64   `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s SearchDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *SearchDepartmentResponseBody) SetHasMore(v bool) *SearchDepartmentResponseBody {
	s.HasMore = &v
	return s
}

func (s *SearchDepartmentResponseBody) SetList(v []*int64) *SearchDepartmentResponseBody {
	s.List = v
	return s
}

func (s *SearchDepartmentResponseBody) SetTotalCount(v int64) *SearchDepartmentResponseBody {
	s.TotalCount = &v
	return s
}

type SearchDepartmentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchDepartmentResponse) GoString() string {
	return s.String()
}

func (s *SearchDepartmentResponse) SetHeaders(v map[string]*string) *SearchDepartmentResponse {
	s.Headers = v
	return s
}

func (s *SearchDepartmentResponse) SetStatusCode(v int32) *SearchDepartmentResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchDepartmentResponse) SetBody(v *SearchDepartmentResponseBody) *SearchDepartmentResponse {
	s.Body = v
	return s
}

type SearchUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchUserHeaders) GoString() string {
	return s.String()
}

func (s *SearchUserHeaders) SetCommonHeaders(v map[string]*string) *SearchUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchUserHeaders) SetXAcsDingtalkAccessToken(v string) *SearchUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchUserRequest struct {
	FullMatchField *int32  `json:"fullMatchField,omitempty" xml:"fullMatchField,omitempty"`
	Offset         *int32  `json:"offset,omitempty" xml:"offset,omitempty"`
	QueryWord      *string `json:"queryWord,omitempty" xml:"queryWord,omitempty"`
	Size           *int32  `json:"size,omitempty" xml:"size,omitempty"`
}

func (s SearchUserRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchUserRequest) GoString() string {
	return s.String()
}

func (s *SearchUserRequest) SetFullMatchField(v int32) *SearchUserRequest {
	s.FullMatchField = &v
	return s
}

func (s *SearchUserRequest) SetOffset(v int32) *SearchUserRequest {
	s.Offset = &v
	return s
}

func (s *SearchUserRequest) SetQueryWord(v string) *SearchUserRequest {
	s.QueryWord = &v
	return s
}

func (s *SearchUserRequest) SetSize(v int32) *SearchUserRequest {
	s.Size = &v
	return s
}

type SearchUserResponseBody struct {
	HasMore    *bool     `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*string `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	TotalCount *int64    `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s SearchUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchUserResponseBody) GoString() string {
	return s.String()
}

func (s *SearchUserResponseBody) SetHasMore(v bool) *SearchUserResponseBody {
	s.HasMore = &v
	return s
}

func (s *SearchUserResponseBody) SetList(v []*string) *SearchUserResponseBody {
	s.List = v
	return s
}

func (s *SearchUserResponseBody) SetTotalCount(v int64) *SearchUserResponseBody {
	s.TotalCount = &v
	return s
}

type SearchUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchUserResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchUserResponse) GoString() string {
	return s.String()
}

func (s *SearchUserResponse) SetHeaders(v map[string]*string) *SearchUserResponse {
	s.Headers = v
	return s
}

func (s *SearchUserResponse) SetStatusCode(v int32) *SearchUserResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchUserResponse) SetBody(v *SearchUserResponseBody) *SearchUserResponse {
	s.Body = v
	return s
}

type SeparateBranchOrgHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SeparateBranchOrgHeaders) String() string {
	return tea.Prettify(s)
}

func (s SeparateBranchOrgHeaders) GoString() string {
	return s.String()
}

func (s *SeparateBranchOrgHeaders) SetCommonHeaders(v map[string]*string) *SeparateBranchOrgHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SeparateBranchOrgHeaders) SetXAcsDingtalkAccessToken(v string) *SeparateBranchOrgHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SeparateBranchOrgRequest struct {
	AttachDeptId *int64 `json:"attachDeptId,omitempty" xml:"attachDeptId,omitempty"`
}

func (s SeparateBranchOrgRequest) String() string {
	return tea.Prettify(s)
}

func (s SeparateBranchOrgRequest) GoString() string {
	return s.String()
}

func (s *SeparateBranchOrgRequest) SetAttachDeptId(v int64) *SeparateBranchOrgRequest {
	s.AttachDeptId = &v
	return s
}

type SeparateBranchOrgResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SeparateBranchOrgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SeparateBranchOrgResponseBody) GoString() string {
	return s.String()
}

func (s *SeparateBranchOrgResponseBody) SetResult(v bool) *SeparateBranchOrgResponseBody {
	s.Result = &v
	return s
}

type SeparateBranchOrgResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SeparateBranchOrgResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SeparateBranchOrgResponse) String() string {
	return tea.Prettify(s)
}

func (s SeparateBranchOrgResponse) GoString() string {
	return s.String()
}

func (s *SeparateBranchOrgResponse) SetHeaders(v map[string]*string) *SeparateBranchOrgResponse {
	s.Headers = v
	return s
}

func (s *SeparateBranchOrgResponse) SetStatusCode(v int32) *SeparateBranchOrgResponse {
	s.StatusCode = &v
	return s
}

func (s *SeparateBranchOrgResponse) SetBody(v *SeparateBranchOrgResponseBody) *SeparateBranchOrgResponse {
	s.Body = v
	return s
}

type SetDisableHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetDisableHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetDisableHeaders) GoString() string {
	return s.String()
}

func (s *SetDisableHeaders) SetCommonHeaders(v map[string]*string) *SetDisableHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetDisableHeaders) SetXAcsDingtalkAccessToken(v string) *SetDisableHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetDisableRequest struct {
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SetDisableRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDisableRequest) GoString() string {
	return s.String()
}

func (s *SetDisableRequest) SetReason(v string) *SetDisableRequest {
	s.Reason = &v
	return s
}

func (s *SetDisableRequest) SetUserId(v string) *SetDisableRequest {
	s.UserId = &v
	return s
}

type SetDisableResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SetDisableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDisableResponseBody) GoString() string {
	return s.String()
}

func (s *SetDisableResponseBody) SetResult(v bool) *SetDisableResponseBody {
	s.Result = &v
	return s
}

type SetDisableResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetDisableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDisableResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDisableResponse) GoString() string {
	return s.String()
}

func (s *SetDisableResponse) SetHeaders(v map[string]*string) *SetDisableResponse {
	s.Headers = v
	return s
}

func (s *SetDisableResponse) SetStatusCode(v int32) *SetDisableResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDisableResponse) SetBody(v *SetDisableResponseBody) *SetDisableResponse {
	s.Body = v
	return s
}

type SetEnableHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetEnableHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetEnableHeaders) GoString() string {
	return s.String()
}

func (s *SetEnableHeaders) SetCommonHeaders(v map[string]*string) *SetEnableHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetEnableHeaders) SetXAcsDingtalkAccessToken(v string) *SetEnableHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetEnableRequest struct {
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SetEnableRequest) String() string {
	return tea.Prettify(s)
}

func (s SetEnableRequest) GoString() string {
	return s.String()
}

func (s *SetEnableRequest) SetUserId(v string) *SetEnableRequest {
	s.UserId = &v
	return s
}

type SetEnableResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SetEnableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetEnableResponseBody) GoString() string {
	return s.String()
}

func (s *SetEnableResponseBody) SetResult(v bool) *SetEnableResponseBody {
	s.Result = &v
	return s
}

type SetEnableResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetEnableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetEnableResponse) String() string {
	return tea.Prettify(s)
}

func (s SetEnableResponse) GoString() string {
	return s.String()
}

func (s *SetEnableResponse) SetHeaders(v map[string]*string) *SetEnableResponse {
	s.Headers = v
	return s
}

func (s *SetEnableResponse) SetStatusCode(v int32) *SetEnableResponse {
	s.StatusCode = &v
	return s
}

func (s *SetEnableResponse) SetBody(v *SetEnableResponseBody) *SetEnableResponse {
	s.Body = v
	return s
}

type SignOutHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SignOutHeaders) String() string {
	return tea.Prettify(s)
}

func (s SignOutHeaders) GoString() string {
	return s.String()
}

func (s *SignOutHeaders) SetCommonHeaders(v map[string]*string) *SignOutHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SignOutHeaders) SetXAcsDingtalkAccessToken(v string) *SignOutHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SignOutRequest struct {
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SignOutRequest) String() string {
	return tea.Prettify(s)
}

func (s SignOutRequest) GoString() string {
	return s.String()
}

func (s *SignOutRequest) SetReason(v string) *SignOutRequest {
	s.Reason = &v
	return s
}

func (s *SignOutRequest) SetUserId(v string) *SignOutRequest {
	s.UserId = &v
	return s
}

type SignOutResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SignOutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SignOutResponseBody) GoString() string {
	return s.String()
}

func (s *SignOutResponseBody) SetResult(v bool) *SignOutResponseBody {
	s.Result = &v
	return s
}

type SignOutResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SignOutResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SignOutResponse) String() string {
	return tea.Prettify(s)
}

func (s SignOutResponse) GoString() string {
	return s.String()
}

func (s *SignOutResponse) SetHeaders(v map[string]*string) *SignOutResponse {
	s.Headers = v
	return s
}

func (s *SignOutResponse) SetStatusCode(v int32) *SignOutResponse {
	s.StatusCode = &v
	return s
}

func (s *SignOutResponse) SetBody(v *SignOutResponseBody) *SignOutResponse {
	s.Body = v
	return s
}

type SortUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SortUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s SortUserHeaders) GoString() string {
	return s.String()
}

func (s *SortUserHeaders) SetCommonHeaders(v map[string]*string) *SortUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SortUserHeaders) SetXAcsDingtalkAccessToken(v string) *SortUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SortUserRequest struct {
	SortType   *int32    `json:"sortType,omitempty" xml:"sortType,omitempty"`
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s SortUserRequest) String() string {
	return tea.Prettify(s)
}

func (s SortUserRequest) GoString() string {
	return s.String()
}

func (s *SortUserRequest) SetSortType(v int32) *SortUserRequest {
	s.SortType = &v
	return s
}

func (s *SortUserRequest) SetUserIdList(v []*string) *SortUserRequest {
	s.UserIdList = v
	return s
}

type SortUserResponseBody struct {
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s SortUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SortUserResponseBody) GoString() string {
	return s.String()
}

func (s *SortUserResponseBody) SetUserIdList(v []*string) *SortUserResponseBody {
	s.UserIdList = v
	return s
}

type SortUserResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SortUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SortUserResponse) String() string {
	return tea.Prettify(s)
}

func (s SortUserResponse) GoString() string {
	return s.String()
}

func (s *SortUserResponse) SetHeaders(v map[string]*string) *SortUserResponse {
	s.Headers = v
	return s
}

func (s *SortUserResponse) SetStatusCode(v int32) *SortUserResponse {
	s.StatusCode = &v
	return s
}

func (s *SortUserResponse) SetBody(v *SortUserResponseBody) *SortUserResponse {
	s.Body = v
	return s
}

type TransformToExclusiveAccountHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TransformToExclusiveAccountHeaders) String() string {
	return tea.Prettify(s)
}

func (s TransformToExclusiveAccountHeaders) GoString() string {
	return s.String()
}

func (s *TransformToExclusiveAccountHeaders) SetCommonHeaders(v map[string]*string) *TransformToExclusiveAccountHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TransformToExclusiveAccountHeaders) SetXAcsDingtalkAccessToken(v string) *TransformToExclusiveAccountHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TransformToExclusiveAccountRequest struct {
	IdpDingTalk   *bool   `json:"idpDingTalk,omitempty" xml:"idpDingTalk,omitempty"`
	InitPassword  *string `json:"initPassword,omitempty" xml:"initPassword,omitempty"`
	LoginId       *string `json:"loginId,omitempty" xml:"loginId,omitempty"`
	TransformType *string `json:"transformType,omitempty" xml:"transformType,omitempty"`
	UserId        *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s TransformToExclusiveAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s TransformToExclusiveAccountRequest) GoString() string {
	return s.String()
}

func (s *TransformToExclusiveAccountRequest) SetIdpDingTalk(v bool) *TransformToExclusiveAccountRequest {
	s.IdpDingTalk = &v
	return s
}

func (s *TransformToExclusiveAccountRequest) SetInitPassword(v string) *TransformToExclusiveAccountRequest {
	s.InitPassword = &v
	return s
}

func (s *TransformToExclusiveAccountRequest) SetLoginId(v string) *TransformToExclusiveAccountRequest {
	s.LoginId = &v
	return s
}

func (s *TransformToExclusiveAccountRequest) SetTransformType(v string) *TransformToExclusiveAccountRequest {
	s.TransformType = &v
	return s
}

func (s *TransformToExclusiveAccountRequest) SetUserId(v string) *TransformToExclusiveAccountRequest {
	s.UserId = &v
	return s
}

type TransformToExclusiveAccountResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s TransformToExclusiveAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s TransformToExclusiveAccountResponse) GoString() string {
	return s.String()
}

func (s *TransformToExclusiveAccountResponse) SetHeaders(v map[string]*string) *TransformToExclusiveAccountResponse {
	s.Headers = v
	return s
}

func (s *TransformToExclusiveAccountResponse) SetStatusCode(v int32) *TransformToExclusiveAccountResponse {
	s.StatusCode = &v
	return s
}

type TranslateFileHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TranslateFileHeaders) String() string {
	return tea.Prettify(s)
}

func (s TranslateFileHeaders) GoString() string {
	return s.String()
}

func (s *TranslateFileHeaders) SetCommonHeaders(v map[string]*string) *TranslateFileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TranslateFileHeaders) SetXAcsDingtalkAccessToken(v string) *TranslateFileHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TranslateFileRequest struct {
	Medias         map[string]*string `json:"medias,omitempty" xml:"medias,omitempty"`
	OutputFileName *string            `json:"outputFileName,omitempty" xml:"outputFileName,omitempty"`
	UnionId        *string            `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s TranslateFileRequest) String() string {
	return tea.Prettify(s)
}

func (s TranslateFileRequest) GoString() string {
	return s.String()
}

func (s *TranslateFileRequest) SetMedias(v map[string]*string) *TranslateFileRequest {
	s.Medias = v
	return s
}

func (s *TranslateFileRequest) SetOutputFileName(v string) *TranslateFileRequest {
	s.OutputFileName = &v
	return s
}

func (s *TranslateFileRequest) SetUnionId(v string) *TranslateFileRequest {
	s.UnionId = &v
	return s
}

type TranslateFileResponseBody struct {
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
}

func (s TranslateFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TranslateFileResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateFileResponseBody) SetJobId(v string) *TranslateFileResponseBody {
	s.JobId = &v
	return s
}

type TranslateFileResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TranslateFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TranslateFileResponse) String() string {
	return tea.Prettify(s)
}

func (s TranslateFileResponse) GoString() string {
	return s.String()
}

func (s *TranslateFileResponse) SetHeaders(v map[string]*string) *TranslateFileResponse {
	s.Headers = v
	return s
}

func (s *TranslateFileResponse) SetStatusCode(v int32) *TranslateFileResponse {
	s.StatusCode = &v
	return s
}

func (s *TranslateFileResponse) SetBody(v *TranslateFileResponseBody) *TranslateFileResponse {
	s.Body = v
	return s
}

type UniqueQueryUserCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UniqueQueryUserCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s UniqueQueryUserCardHeaders) GoString() string {
	return s.String()
}

func (s *UniqueQueryUserCardHeaders) SetCommonHeaders(v map[string]*string) *UniqueQueryUserCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UniqueQueryUserCardHeaders) SetXAcsDingtalkAccessToken(v string) *UniqueQueryUserCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UniqueQueryUserCardRequest struct {
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	UnionId    *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s UniqueQueryUserCardRequest) String() string {
	return tea.Prettify(s)
}

func (s UniqueQueryUserCardRequest) GoString() string {
	return s.String()
}

func (s *UniqueQueryUserCardRequest) SetTemplateId(v string) *UniqueQueryUserCardRequest {
	s.TemplateId = &v
	return s
}

func (s *UniqueQueryUserCardRequest) SetUnionId(v string) *UniqueQueryUserCardRequest {
	s.UnionId = &v
	return s
}

type UniqueQueryUserCardResponseBody struct {
	AvatarUrl    *string                `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	CardId       *string                `json:"cardId,omitempty" xml:"cardId,omitempty"`
	Extension    map[string]interface{} `json:"extension,omitempty" xml:"extension,omitempty"`
	IndustryName *string                `json:"industryName,omitempty" xml:"industryName,omitempty"`
	Introduce    *string                `json:"introduce,omitempty" xml:"introduce,omitempty"`
	Name         *string                `json:"name,omitempty" xml:"name,omitempty"`
	OrgName      *string                `json:"orgName,omitempty" xml:"orgName,omitempty"`
	Settings     map[string]interface{} `json:"settings,omitempty" xml:"settings,omitempty"`
	TemplateId   *string                `json:"templateId,omitempty" xml:"templateId,omitempty"`
	Title        *string                `json:"title,omitempty" xml:"title,omitempty"`
}

func (s UniqueQueryUserCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UniqueQueryUserCardResponseBody) GoString() string {
	return s.String()
}

func (s *UniqueQueryUserCardResponseBody) SetAvatarUrl(v string) *UniqueQueryUserCardResponseBody {
	s.AvatarUrl = &v
	return s
}

func (s *UniqueQueryUserCardResponseBody) SetCardId(v string) *UniqueQueryUserCardResponseBody {
	s.CardId = &v
	return s
}

func (s *UniqueQueryUserCardResponseBody) SetExtension(v map[string]interface{}) *UniqueQueryUserCardResponseBody {
	s.Extension = v
	return s
}

func (s *UniqueQueryUserCardResponseBody) SetIndustryName(v string) *UniqueQueryUserCardResponseBody {
	s.IndustryName = &v
	return s
}

func (s *UniqueQueryUserCardResponseBody) SetIntroduce(v string) *UniqueQueryUserCardResponseBody {
	s.Introduce = &v
	return s
}

func (s *UniqueQueryUserCardResponseBody) SetName(v string) *UniqueQueryUserCardResponseBody {
	s.Name = &v
	return s
}

func (s *UniqueQueryUserCardResponseBody) SetOrgName(v string) *UniqueQueryUserCardResponseBody {
	s.OrgName = &v
	return s
}

func (s *UniqueQueryUserCardResponseBody) SetSettings(v map[string]interface{}) *UniqueQueryUserCardResponseBody {
	s.Settings = v
	return s
}

func (s *UniqueQueryUserCardResponseBody) SetTemplateId(v string) *UniqueQueryUserCardResponseBody {
	s.TemplateId = &v
	return s
}

func (s *UniqueQueryUserCardResponseBody) SetTitle(v string) *UniqueQueryUserCardResponseBody {
	s.Title = &v
	return s
}

type UniqueQueryUserCardResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UniqueQueryUserCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UniqueQueryUserCardResponse) String() string {
	return tea.Prettify(s)
}

func (s UniqueQueryUserCardResponse) GoString() string {
	return s.String()
}

func (s *UniqueQueryUserCardResponse) SetHeaders(v map[string]*string) *UniqueQueryUserCardResponse {
	s.Headers = v
	return s
}

func (s *UniqueQueryUserCardResponse) SetStatusCode(v int32) *UniqueQueryUserCardResponse {
	s.StatusCode = &v
	return s
}

func (s *UniqueQueryUserCardResponse) SetBody(v *UniqueQueryUserCardResponseBody) *UniqueQueryUserCardResponse {
	s.Body = v
	return s
}

type UpdateBranchAttributesInCooperateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateBranchAttributesInCooperateHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateBranchAttributesInCooperateHeaders) GoString() string {
	return s.String()
}

func (s *UpdateBranchAttributesInCooperateHeaders) SetCommonHeaders(v map[string]*string) *UpdateBranchAttributesInCooperateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateBranchAttributesInCooperateHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateBranchAttributesInCooperateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateBranchAttributesInCooperateRequest struct {
	Body []*UpdateBranchAttributesInCooperateRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s UpdateBranchAttributesInCooperateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBranchAttributesInCooperateRequest) GoString() string {
	return s.String()
}

func (s *UpdateBranchAttributesInCooperateRequest) SetBody(v []*UpdateBranchAttributesInCooperateRequestBody) *UpdateBranchAttributesInCooperateRequest {
	s.Body = v
	return s
}

type UpdateBranchAttributesInCooperateRequestBody struct {
	BranchCorpId  *string `json:"branchCorpId,omitempty" xml:"branchCorpId,omitempty"`
	LinkDeptId    *int64  `json:"linkDeptId,omitempty" xml:"linkDeptId,omitempty"`
	UnionRootName *string `json:"unionRootName,omitempty" xml:"unionRootName,omitempty"`
}

func (s UpdateBranchAttributesInCooperateRequestBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBranchAttributesInCooperateRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateBranchAttributesInCooperateRequestBody) SetBranchCorpId(v string) *UpdateBranchAttributesInCooperateRequestBody {
	s.BranchCorpId = &v
	return s
}

func (s *UpdateBranchAttributesInCooperateRequestBody) SetLinkDeptId(v int64) *UpdateBranchAttributesInCooperateRequestBody {
	s.LinkDeptId = &v
	return s
}

func (s *UpdateBranchAttributesInCooperateRequestBody) SetUnionRootName(v string) *UpdateBranchAttributesInCooperateRequestBody {
	s.UnionRootName = &v
	return s
}

type UpdateBranchAttributesInCooperateResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateBranchAttributesInCooperateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBranchAttributesInCooperateResponse) GoString() string {
	return s.String()
}

func (s *UpdateBranchAttributesInCooperateResponse) SetHeaders(v map[string]*string) *UpdateBranchAttributesInCooperateResponse {
	s.Headers = v
	return s
}

func (s *UpdateBranchAttributesInCooperateResponse) SetStatusCode(v int32) *UpdateBranchAttributesInCooperateResponse {
	s.StatusCode = &v
	return s
}

type UpdateBranchVisibleSettingInCooperateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateBranchVisibleSettingInCooperateHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateBranchVisibleSettingInCooperateHeaders) GoString() string {
	return s.String()
}

func (s *UpdateBranchVisibleSettingInCooperateHeaders) SetCommonHeaders(v map[string]*string) *UpdateBranchVisibleSettingInCooperateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateBranchVisibleSettingInCooperateHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateBranchVisibleSettingInCooperateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateBranchVisibleSettingInCooperateRequest struct {
	Body []*UpdateBranchVisibleSettingInCooperateRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s UpdateBranchVisibleSettingInCooperateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBranchVisibleSettingInCooperateRequest) GoString() string {
	return s.String()
}

func (s *UpdateBranchVisibleSettingInCooperateRequest) SetBody(v []*UpdateBranchVisibleSettingInCooperateRequestBody) *UpdateBranchVisibleSettingInCooperateRequest {
	s.Body = v
	return s
}

type UpdateBranchVisibleSettingInCooperateRequestBody struct {
	BranchCorpId         *string   `json:"branchCorpId,omitempty" xml:"branchCorpId,omitempty"`
	Open                 *bool     `json:"open,omitempty" xml:"open,omitempty"`
	Type                 *int64    `json:"type,omitempty" xml:"type,omitempty"`
	VisibleBranchCorpIds []*string `json:"visibleBranchCorpIds,omitempty" xml:"visibleBranchCorpIds,omitempty" type:"Repeated"`
	VisibleDeptIds       []*int64  `json:"visibleDeptIds,omitempty" xml:"visibleDeptIds,omitempty" type:"Repeated"`
}

func (s UpdateBranchVisibleSettingInCooperateRequestBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBranchVisibleSettingInCooperateRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateBranchVisibleSettingInCooperateRequestBody) SetBranchCorpId(v string) *UpdateBranchVisibleSettingInCooperateRequestBody {
	s.BranchCorpId = &v
	return s
}

func (s *UpdateBranchVisibleSettingInCooperateRequestBody) SetOpen(v bool) *UpdateBranchVisibleSettingInCooperateRequestBody {
	s.Open = &v
	return s
}

func (s *UpdateBranchVisibleSettingInCooperateRequestBody) SetType(v int64) *UpdateBranchVisibleSettingInCooperateRequestBody {
	s.Type = &v
	return s
}

func (s *UpdateBranchVisibleSettingInCooperateRequestBody) SetVisibleBranchCorpIds(v []*string) *UpdateBranchVisibleSettingInCooperateRequestBody {
	s.VisibleBranchCorpIds = v
	return s
}

func (s *UpdateBranchVisibleSettingInCooperateRequestBody) SetVisibleDeptIds(v []*int64) *UpdateBranchVisibleSettingInCooperateRequestBody {
	s.VisibleDeptIds = v
	return s
}

type UpdateBranchVisibleSettingInCooperateResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateBranchVisibleSettingInCooperateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBranchVisibleSettingInCooperateResponse) GoString() string {
	return s.String()
}

func (s *UpdateBranchVisibleSettingInCooperateResponse) SetHeaders(v map[string]*string) *UpdateBranchVisibleSettingInCooperateResponse {
	s.Headers = v
	return s
}

func (s *UpdateBranchVisibleSettingInCooperateResponse) SetStatusCode(v int32) *UpdateBranchVisibleSettingInCooperateResponse {
	s.StatusCode = &v
	return s
}

type UpdateContactHideBySceneSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateContactHideBySceneSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactHideBySceneSettingHeaders) GoString() string {
	return s.String()
}

func (s *UpdateContactHideBySceneSettingHeaders) SetCommonHeaders(v map[string]*string) *UpdateContactHideBySceneSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateContactHideBySceneSettingHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateContactHideBySceneSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateContactHideBySceneSettingRequest struct {
	Description         *string                                                    `json:"description,omitempty" xml:"description,omitempty"`
	ExcludeDeptIds      []*int64                                                   `json:"excludeDeptIds,omitempty" xml:"excludeDeptIds,omitempty" type:"Repeated"`
	ExcludeTagIds       []*int64                                                   `json:"excludeTagIds,omitempty" xml:"excludeTagIds,omitempty" type:"Repeated"`
	ExcludeUserIds      []*string                                                  `json:"excludeUserIds,omitempty" xml:"excludeUserIds,omitempty" type:"Repeated"`
	Name                *string                                                    `json:"name,omitempty" xml:"name,omitempty"`
	NodeListSceneConfig *UpdateContactHideBySceneSettingRequestNodeListSceneConfig `json:"nodeListSceneConfig,omitempty" xml:"nodeListSceneConfig,omitempty" type:"Struct"`
	ObjectDeptIds       []*int64                                                   `json:"objectDeptIds,omitempty" xml:"objectDeptIds,omitempty" type:"Repeated"`
	ObjectTagIds        []*int64                                                   `json:"objectTagIds,omitempty" xml:"objectTagIds,omitempty" type:"Repeated"`
	ObjectUserIds       []*string                                                  `json:"objectUserIds,omitempty" xml:"objectUserIds,omitempty" type:"Repeated"`
	ProfileSceneConfig  *UpdateContactHideBySceneSettingRequestProfileSceneConfig  `json:"profileSceneConfig,omitempty" xml:"profileSceneConfig,omitempty" type:"Struct"`
	SearchSceneConfig   *UpdateContactHideBySceneSettingRequestSearchSceneConfig   `json:"searchSceneConfig,omitempty" xml:"searchSceneConfig,omitempty" type:"Struct"`
}

func (s UpdateContactHideBySceneSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactHideBySceneSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateContactHideBySceneSettingRequest) SetDescription(v string) *UpdateContactHideBySceneSettingRequest {
	s.Description = &v
	return s
}

func (s *UpdateContactHideBySceneSettingRequest) SetExcludeDeptIds(v []*int64) *UpdateContactHideBySceneSettingRequest {
	s.ExcludeDeptIds = v
	return s
}

func (s *UpdateContactHideBySceneSettingRequest) SetExcludeTagIds(v []*int64) *UpdateContactHideBySceneSettingRequest {
	s.ExcludeTagIds = v
	return s
}

func (s *UpdateContactHideBySceneSettingRequest) SetExcludeUserIds(v []*string) *UpdateContactHideBySceneSettingRequest {
	s.ExcludeUserIds = v
	return s
}

func (s *UpdateContactHideBySceneSettingRequest) SetName(v string) *UpdateContactHideBySceneSettingRequest {
	s.Name = &v
	return s
}

func (s *UpdateContactHideBySceneSettingRequest) SetNodeListSceneConfig(v *UpdateContactHideBySceneSettingRequestNodeListSceneConfig) *UpdateContactHideBySceneSettingRequest {
	s.NodeListSceneConfig = v
	return s
}

func (s *UpdateContactHideBySceneSettingRequest) SetObjectDeptIds(v []*int64) *UpdateContactHideBySceneSettingRequest {
	s.ObjectDeptIds = v
	return s
}

func (s *UpdateContactHideBySceneSettingRequest) SetObjectTagIds(v []*int64) *UpdateContactHideBySceneSettingRequest {
	s.ObjectTagIds = v
	return s
}

func (s *UpdateContactHideBySceneSettingRequest) SetObjectUserIds(v []*string) *UpdateContactHideBySceneSettingRequest {
	s.ObjectUserIds = v
	return s
}

func (s *UpdateContactHideBySceneSettingRequest) SetProfileSceneConfig(v *UpdateContactHideBySceneSettingRequestProfileSceneConfig) *UpdateContactHideBySceneSettingRequest {
	s.ProfileSceneConfig = v
	return s
}

func (s *UpdateContactHideBySceneSettingRequest) SetSearchSceneConfig(v *UpdateContactHideBySceneSettingRequestSearchSceneConfig) *UpdateContactHideBySceneSettingRequest {
	s.SearchSceneConfig = v
	return s
}

type UpdateContactHideBySceneSettingRequestNodeListSceneConfig struct {
	Active               *bool `json:"active,omitempty" xml:"active,omitempty"`
	DeptObjectIncludeEmp *bool `json:"deptObjectIncludeEmp,omitempty" xml:"deptObjectIncludeEmp,omitempty"`
}

func (s UpdateContactHideBySceneSettingRequestNodeListSceneConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactHideBySceneSettingRequestNodeListSceneConfig) GoString() string {
	return s.String()
}

func (s *UpdateContactHideBySceneSettingRequestNodeListSceneConfig) SetActive(v bool) *UpdateContactHideBySceneSettingRequestNodeListSceneConfig {
	s.Active = &v
	return s
}

func (s *UpdateContactHideBySceneSettingRequestNodeListSceneConfig) SetDeptObjectIncludeEmp(v bool) *UpdateContactHideBySceneSettingRequestNodeListSceneConfig {
	s.DeptObjectIncludeEmp = &v
	return s
}

type UpdateContactHideBySceneSettingRequestProfileSceneConfig struct {
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
}

func (s UpdateContactHideBySceneSettingRequestProfileSceneConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactHideBySceneSettingRequestProfileSceneConfig) GoString() string {
	return s.String()
}

func (s *UpdateContactHideBySceneSettingRequestProfileSceneConfig) SetActive(v bool) *UpdateContactHideBySceneSettingRequestProfileSceneConfig {
	s.Active = &v
	return s
}

type UpdateContactHideBySceneSettingRequestSearchSceneConfig struct {
	Active               *bool `json:"active,omitempty" xml:"active,omitempty"`
	DeptObjectIncludeEmp *bool `json:"deptObjectIncludeEmp,omitempty" xml:"deptObjectIncludeEmp,omitempty"`
}

func (s UpdateContactHideBySceneSettingRequestSearchSceneConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactHideBySceneSettingRequestSearchSceneConfig) GoString() string {
	return s.String()
}

func (s *UpdateContactHideBySceneSettingRequestSearchSceneConfig) SetActive(v bool) *UpdateContactHideBySceneSettingRequestSearchSceneConfig {
	s.Active = &v
	return s
}

func (s *UpdateContactHideBySceneSettingRequestSearchSceneConfig) SetDeptObjectIncludeEmp(v bool) *UpdateContactHideBySceneSettingRequestSearchSceneConfig {
	s.DeptObjectIncludeEmp = &v
	return s
}

type UpdateContactHideBySceneSettingResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateContactHideBySceneSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactHideBySceneSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContactHideBySceneSettingResponseBody) SetSuccess(v bool) *UpdateContactHideBySceneSettingResponseBody {
	s.Success = &v
	return s
}

type UpdateContactHideBySceneSettingResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateContactHideBySceneSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateContactHideBySceneSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactHideBySceneSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateContactHideBySceneSettingResponse) SetHeaders(v map[string]*string) *UpdateContactHideBySceneSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateContactHideBySceneSettingResponse) SetStatusCode(v int32) *UpdateContactHideBySceneSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContactHideBySceneSettingResponse) SetBody(v *UpdateContactHideBySceneSettingResponseBody) *UpdateContactHideBySceneSettingResponse {
	s.Body = v
	return s
}

type UpdateContactHideSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateContactHideSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactHideSettingHeaders) GoString() string {
	return s.String()
}

func (s *UpdateContactHideSettingHeaders) SetCommonHeaders(v map[string]*string) *UpdateContactHideSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateContactHideSettingHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateContactHideSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateContactHideSettingRequest struct {
	Active            *bool     `json:"active,omitempty" xml:"active,omitempty"`
	Description       *string   `json:"description,omitempty" xml:"description,omitempty"`
	ExcludeDeptIds    []*int64  `json:"excludeDeptIds,omitempty" xml:"excludeDeptIds,omitempty" type:"Repeated"`
	ExcludeStaffIds   []*string `json:"excludeStaffIds,omitempty" xml:"excludeStaffIds,omitempty" type:"Repeated"`
	ExcludeTagIds     []*int64  `json:"excludeTagIds,omitempty" xml:"excludeTagIds,omitempty" type:"Repeated"`
	HideInSearch      *bool     `json:"hideInSearch,omitempty" xml:"hideInSearch,omitempty"`
	HideInUserProfile *bool     `json:"hideInUserProfile,omitempty" xml:"hideInUserProfile,omitempty"`
	Id                *int64    `json:"id,omitempty" xml:"id,omitempty"`
	Name              *string   `json:"name,omitempty" xml:"name,omitempty"`
	ObjectDeptIds     []*int64  `json:"objectDeptIds,omitempty" xml:"objectDeptIds,omitempty" type:"Repeated"`
	ObjectStaffIds    []*string `json:"objectStaffIds,omitempty" xml:"objectStaffIds,omitempty" type:"Repeated"`
	ObjectTagIds      []*int64  `json:"objectTagIds,omitempty" xml:"objectTagIds,omitempty" type:"Repeated"`
}

func (s UpdateContactHideSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactHideSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateContactHideSettingRequest) SetActive(v bool) *UpdateContactHideSettingRequest {
	s.Active = &v
	return s
}

func (s *UpdateContactHideSettingRequest) SetDescription(v string) *UpdateContactHideSettingRequest {
	s.Description = &v
	return s
}

func (s *UpdateContactHideSettingRequest) SetExcludeDeptIds(v []*int64) *UpdateContactHideSettingRequest {
	s.ExcludeDeptIds = v
	return s
}

func (s *UpdateContactHideSettingRequest) SetExcludeStaffIds(v []*string) *UpdateContactHideSettingRequest {
	s.ExcludeStaffIds = v
	return s
}

func (s *UpdateContactHideSettingRequest) SetExcludeTagIds(v []*int64) *UpdateContactHideSettingRequest {
	s.ExcludeTagIds = v
	return s
}

func (s *UpdateContactHideSettingRequest) SetHideInSearch(v bool) *UpdateContactHideSettingRequest {
	s.HideInSearch = &v
	return s
}

func (s *UpdateContactHideSettingRequest) SetHideInUserProfile(v bool) *UpdateContactHideSettingRequest {
	s.HideInUserProfile = &v
	return s
}

func (s *UpdateContactHideSettingRequest) SetId(v int64) *UpdateContactHideSettingRequest {
	s.Id = &v
	return s
}

func (s *UpdateContactHideSettingRequest) SetName(v string) *UpdateContactHideSettingRequest {
	s.Name = &v
	return s
}

func (s *UpdateContactHideSettingRequest) SetObjectDeptIds(v []*int64) *UpdateContactHideSettingRequest {
	s.ObjectDeptIds = v
	return s
}

func (s *UpdateContactHideSettingRequest) SetObjectStaffIds(v []*string) *UpdateContactHideSettingRequest {
	s.ObjectStaffIds = v
	return s
}

func (s *UpdateContactHideSettingRequest) SetObjectTagIds(v []*int64) *UpdateContactHideSettingRequest {
	s.ObjectTagIds = v
	return s
}

type UpdateContactHideSettingResponseBody struct {
	Result *int64 `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateContactHideSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactHideSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContactHideSettingResponseBody) SetResult(v int64) *UpdateContactHideSettingResponseBody {
	s.Result = &v
	return s
}

type UpdateContactHideSettingResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateContactHideSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateContactHideSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactHideSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateContactHideSettingResponse) SetHeaders(v map[string]*string) *UpdateContactHideSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateContactHideSettingResponse) SetStatusCode(v int32) *UpdateContactHideSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContactHideSettingResponse) SetBody(v *UpdateContactHideSettingResponseBody) *UpdateContactHideSettingResponse {
	s.Body = v
	return s
}

type UpdateContactRestrictSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateContactRestrictSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactRestrictSettingHeaders) GoString() string {
	return s.String()
}

func (s *UpdateContactRestrictSettingHeaders) SetCommonHeaders(v map[string]*string) *UpdateContactRestrictSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateContactRestrictSettingHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateContactRestrictSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateContactRestrictSettingRequest struct {
	Active                *bool     `json:"active,omitempty" xml:"active,omitempty"`
	Description           *string   `json:"description,omitempty" xml:"description,omitempty"`
	ExcludeDeptIds        []*int64  `json:"excludeDeptIds,omitempty" xml:"excludeDeptIds,omitempty" type:"Repeated"`
	ExcludeTagIds         []*int64  `json:"excludeTagIds,omitempty" xml:"excludeTagIds,omitempty" type:"Repeated"`
	ExcludeUserIds        []*string `json:"excludeUserIds,omitempty" xml:"excludeUserIds,omitempty" type:"Repeated"`
	Id                    *int64    `json:"id,omitempty" xml:"id,omitempty"`
	Name                  *string   `json:"name,omitempty" xml:"name,omitempty"`
	RestrictInSearch      *bool     `json:"restrictInSearch,omitempty" xml:"restrictInSearch,omitempty"`
	RestrictInUserProfile *bool     `json:"restrictInUserProfile,omitempty" xml:"restrictInUserProfile,omitempty"`
	SubjectDeptIds        []*int64  `json:"subjectDeptIds,omitempty" xml:"subjectDeptIds,omitempty" type:"Repeated"`
	SubjectTagIds         []*int64  `json:"subjectTagIds,omitempty" xml:"subjectTagIds,omitempty" type:"Repeated"`
	SubjectUserIds        []*string `json:"subjectUserIds,omitempty" xml:"subjectUserIds,omitempty" type:"Repeated"`
	Type                  *string   `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateContactRestrictSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactRestrictSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateContactRestrictSettingRequest) SetActive(v bool) *UpdateContactRestrictSettingRequest {
	s.Active = &v
	return s
}

func (s *UpdateContactRestrictSettingRequest) SetDescription(v string) *UpdateContactRestrictSettingRequest {
	s.Description = &v
	return s
}

func (s *UpdateContactRestrictSettingRequest) SetExcludeDeptIds(v []*int64) *UpdateContactRestrictSettingRequest {
	s.ExcludeDeptIds = v
	return s
}

func (s *UpdateContactRestrictSettingRequest) SetExcludeTagIds(v []*int64) *UpdateContactRestrictSettingRequest {
	s.ExcludeTagIds = v
	return s
}

func (s *UpdateContactRestrictSettingRequest) SetExcludeUserIds(v []*string) *UpdateContactRestrictSettingRequest {
	s.ExcludeUserIds = v
	return s
}

func (s *UpdateContactRestrictSettingRequest) SetId(v int64) *UpdateContactRestrictSettingRequest {
	s.Id = &v
	return s
}

func (s *UpdateContactRestrictSettingRequest) SetName(v string) *UpdateContactRestrictSettingRequest {
	s.Name = &v
	return s
}

func (s *UpdateContactRestrictSettingRequest) SetRestrictInSearch(v bool) *UpdateContactRestrictSettingRequest {
	s.RestrictInSearch = &v
	return s
}

func (s *UpdateContactRestrictSettingRequest) SetRestrictInUserProfile(v bool) *UpdateContactRestrictSettingRequest {
	s.RestrictInUserProfile = &v
	return s
}

func (s *UpdateContactRestrictSettingRequest) SetSubjectDeptIds(v []*int64) *UpdateContactRestrictSettingRequest {
	s.SubjectDeptIds = v
	return s
}

func (s *UpdateContactRestrictSettingRequest) SetSubjectTagIds(v []*int64) *UpdateContactRestrictSettingRequest {
	s.SubjectTagIds = v
	return s
}

func (s *UpdateContactRestrictSettingRequest) SetSubjectUserIds(v []*string) *UpdateContactRestrictSettingRequest {
	s.SubjectUserIds = v
	return s
}

func (s *UpdateContactRestrictSettingRequest) SetType(v string) *UpdateContactRestrictSettingRequest {
	s.Type = &v
	return s
}

type UpdateContactRestrictSettingResponseBody struct {
	Result *int64 `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateContactRestrictSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactRestrictSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContactRestrictSettingResponseBody) SetResult(v int64) *UpdateContactRestrictSettingResponseBody {
	s.Result = &v
	return s
}

type UpdateContactRestrictSettingResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateContactRestrictSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateContactRestrictSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactRestrictSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateContactRestrictSettingResponse) SetHeaders(v map[string]*string) *UpdateContactRestrictSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateContactRestrictSettingResponse) SetStatusCode(v int32) *UpdateContactRestrictSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContactRestrictSettingResponse) SetBody(v *UpdateContactRestrictSettingResponseBody) *UpdateContactRestrictSettingResponse {
	s.Body = v
	return s
}

type UpdateDeptSettngTailFirstHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateDeptSettngTailFirstHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeptSettngTailFirstHeaders) GoString() string {
	return s.String()
}

func (s *UpdateDeptSettngTailFirstHeaders) SetCommonHeaders(v map[string]*string) *UpdateDeptSettngTailFirstHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateDeptSettngTailFirstHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateDeptSettngTailFirstHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateDeptSettngTailFirstRequest struct {
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s UpdateDeptSettngTailFirstRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeptSettngTailFirstRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeptSettngTailFirstRequest) SetEnable(v bool) *UpdateDeptSettngTailFirstRequest {
	s.Enable = &v
	return s
}

type UpdateDeptSettngTailFirstResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateDeptSettngTailFirstResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeptSettngTailFirstResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeptSettngTailFirstResponse) SetHeaders(v map[string]*string) *UpdateDeptSettngTailFirstResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeptSettngTailFirstResponse) SetStatusCode(v int32) *UpdateDeptSettngTailFirstResponse {
	s.StatusCode = &v
	return s
}

type UpdateEmpAttrbuteVisibilitySettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateEmpAttrbuteVisibilitySettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateEmpAttrbuteVisibilitySettingHeaders) GoString() string {
	return s.String()
}

func (s *UpdateEmpAttrbuteVisibilitySettingHeaders) SetCommonHeaders(v map[string]*string) *UpdateEmpAttrbuteVisibilitySettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateEmpAttrbuteVisibilitySettingHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateEmpAttrbuteVisibilitySettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateEmpAttrbuteVisibilitySettingRequest struct {
	Active          *bool     `json:"active,omitempty" xml:"active,omitempty"`
	Description     *string   `json:"description,omitempty" xml:"description,omitempty"`
	ExcludeDeptIds  []*int64  `json:"excludeDeptIds,omitempty" xml:"excludeDeptIds,omitempty" type:"Repeated"`
	ExcludeStaffIds []*string `json:"excludeStaffIds,omitempty" xml:"excludeStaffIds,omitempty" type:"Repeated"`
	ExcludeTagIds   []*int64  `json:"excludeTagIds,omitempty" xml:"excludeTagIds,omitempty" type:"Repeated"`
	HideFields      []*string `json:"hideFields,omitempty" xml:"hideFields,omitempty" type:"Repeated"`
	Id              *int64    `json:"id,omitempty" xml:"id,omitempty"`
	Name            *string   `json:"name,omitempty" xml:"name,omitempty"`
	ObjectDeptIds   []*int64  `json:"objectDeptIds,omitempty" xml:"objectDeptIds,omitempty" type:"Repeated"`
	ObjectStaffIds  []*string `json:"objectStaffIds,omitempty" xml:"objectStaffIds,omitempty" type:"Repeated"`
	ObjectTagIds    []*int64  `json:"objectTagIds,omitempty" xml:"objectTagIds,omitempty" type:"Repeated"`
}

func (s UpdateEmpAttrbuteVisibilitySettingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEmpAttrbuteVisibilitySettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateEmpAttrbuteVisibilitySettingRequest) SetActive(v bool) *UpdateEmpAttrbuteVisibilitySettingRequest {
	s.Active = &v
	return s
}

func (s *UpdateEmpAttrbuteVisibilitySettingRequest) SetDescription(v string) *UpdateEmpAttrbuteVisibilitySettingRequest {
	s.Description = &v
	return s
}

func (s *UpdateEmpAttrbuteVisibilitySettingRequest) SetExcludeDeptIds(v []*int64) *UpdateEmpAttrbuteVisibilitySettingRequest {
	s.ExcludeDeptIds = v
	return s
}

func (s *UpdateEmpAttrbuteVisibilitySettingRequest) SetExcludeStaffIds(v []*string) *UpdateEmpAttrbuteVisibilitySettingRequest {
	s.ExcludeStaffIds = v
	return s
}

func (s *UpdateEmpAttrbuteVisibilitySettingRequest) SetExcludeTagIds(v []*int64) *UpdateEmpAttrbuteVisibilitySettingRequest {
	s.ExcludeTagIds = v
	return s
}

func (s *UpdateEmpAttrbuteVisibilitySettingRequest) SetHideFields(v []*string) *UpdateEmpAttrbuteVisibilitySettingRequest {
	s.HideFields = v
	return s
}

func (s *UpdateEmpAttrbuteVisibilitySettingRequest) SetId(v int64) *UpdateEmpAttrbuteVisibilitySettingRequest {
	s.Id = &v
	return s
}

func (s *UpdateEmpAttrbuteVisibilitySettingRequest) SetName(v string) *UpdateEmpAttrbuteVisibilitySettingRequest {
	s.Name = &v
	return s
}

func (s *UpdateEmpAttrbuteVisibilitySettingRequest) SetObjectDeptIds(v []*int64) *UpdateEmpAttrbuteVisibilitySettingRequest {
	s.ObjectDeptIds = v
	return s
}

func (s *UpdateEmpAttrbuteVisibilitySettingRequest) SetObjectStaffIds(v []*string) *UpdateEmpAttrbuteVisibilitySettingRequest {
	s.ObjectStaffIds = v
	return s
}

func (s *UpdateEmpAttrbuteVisibilitySettingRequest) SetObjectTagIds(v []*int64) *UpdateEmpAttrbuteVisibilitySettingRequest {
	s.ObjectTagIds = v
	return s
}

type UpdateEmpAttrbuteVisibilitySettingResponseBody struct {
	Result *int64 `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateEmpAttrbuteVisibilitySettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEmpAttrbuteVisibilitySettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEmpAttrbuteVisibilitySettingResponseBody) SetResult(v int64) *UpdateEmpAttrbuteVisibilitySettingResponseBody {
	s.Result = &v
	return s
}

type UpdateEmpAttrbuteVisibilitySettingResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateEmpAttrbuteVisibilitySettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateEmpAttrbuteVisibilitySettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEmpAttrbuteVisibilitySettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateEmpAttrbuteVisibilitySettingResponse) SetHeaders(v map[string]*string) *UpdateEmpAttrbuteVisibilitySettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateEmpAttrbuteVisibilitySettingResponse) SetStatusCode(v int32) *UpdateEmpAttrbuteVisibilitySettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEmpAttrbuteVisibilitySettingResponse) SetBody(v *UpdateEmpAttrbuteVisibilitySettingResponseBody) *UpdateEmpAttrbuteVisibilitySettingResponse {
	s.Body = v
	return s
}

type UpdateEmpAttributeHideBySceneSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateEmpAttributeHideBySceneSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateEmpAttributeHideBySceneSettingHeaders) GoString() string {
	return s.String()
}

func (s *UpdateEmpAttributeHideBySceneSettingHeaders) SetCommonHeaders(v map[string]*string) *UpdateEmpAttributeHideBySceneSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateEmpAttributeHideBySceneSettingHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateEmpAttributeHideBySceneSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateEmpAttributeHideBySceneSettingRequest struct {
	ChatSubtitleConfig *UpdateEmpAttributeHideBySceneSettingRequestChatSubtitleConfig `json:"chatSubtitleConfig,omitempty" xml:"chatSubtitleConfig,omitempty" type:"Struct"`
	Description        *string                                                        `json:"description,omitempty" xml:"description,omitempty"`
	ExcludeDeptIds     []*int64                                                       `json:"excludeDeptIds,omitempty" xml:"excludeDeptIds,omitempty" type:"Repeated"`
	ExcludeTagIds      []*int64                                                       `json:"excludeTagIds,omitempty" xml:"excludeTagIds,omitempty" type:"Repeated"`
	ExcludeUserIds     []*string                                                      `json:"excludeUserIds,omitempty" xml:"excludeUserIds,omitempty" type:"Repeated"`
	HideFields         []*string                                                      `json:"hideFields,omitempty" xml:"hideFields,omitempty" type:"Repeated"`
	Name               *string                                                        `json:"name,omitempty" xml:"name,omitempty"`
	ObjectDeptIds      []*int64                                                       `json:"objectDeptIds,omitempty" xml:"objectDeptIds,omitempty" type:"Repeated"`
	ObjectTagIds       []*int64                                                       `json:"objectTagIds,omitempty" xml:"objectTagIds,omitempty" type:"Repeated"`
	ObjectUserIds      []*string                                                      `json:"objectUserIds,omitempty" xml:"objectUserIds,omitempty" type:"Repeated"`
	ProfileSceneConfig *UpdateEmpAttributeHideBySceneSettingRequestProfileSceneConfig `json:"profileSceneConfig,omitempty" xml:"profileSceneConfig,omitempty" type:"Struct"`
	SearchSceneConfig  *UpdateEmpAttributeHideBySceneSettingRequestSearchSceneConfig  `json:"searchSceneConfig,omitempty" xml:"searchSceneConfig,omitempty" type:"Struct"`
}

func (s UpdateEmpAttributeHideBySceneSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEmpAttributeHideBySceneSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateEmpAttributeHideBySceneSettingRequest) SetChatSubtitleConfig(v *UpdateEmpAttributeHideBySceneSettingRequestChatSubtitleConfig) *UpdateEmpAttributeHideBySceneSettingRequest {
	s.ChatSubtitleConfig = v
	return s
}

func (s *UpdateEmpAttributeHideBySceneSettingRequest) SetDescription(v string) *UpdateEmpAttributeHideBySceneSettingRequest {
	s.Description = &v
	return s
}

func (s *UpdateEmpAttributeHideBySceneSettingRequest) SetExcludeDeptIds(v []*int64) *UpdateEmpAttributeHideBySceneSettingRequest {
	s.ExcludeDeptIds = v
	return s
}

func (s *UpdateEmpAttributeHideBySceneSettingRequest) SetExcludeTagIds(v []*int64) *UpdateEmpAttributeHideBySceneSettingRequest {
	s.ExcludeTagIds = v
	return s
}

func (s *UpdateEmpAttributeHideBySceneSettingRequest) SetExcludeUserIds(v []*string) *UpdateEmpAttributeHideBySceneSettingRequest {
	s.ExcludeUserIds = v
	return s
}

func (s *UpdateEmpAttributeHideBySceneSettingRequest) SetHideFields(v []*string) *UpdateEmpAttributeHideBySceneSettingRequest {
	s.HideFields = v
	return s
}

func (s *UpdateEmpAttributeHideBySceneSettingRequest) SetName(v string) *UpdateEmpAttributeHideBySceneSettingRequest {
	s.Name = &v
	return s
}

func (s *UpdateEmpAttributeHideBySceneSettingRequest) SetObjectDeptIds(v []*int64) *UpdateEmpAttributeHideBySceneSettingRequest {
	s.ObjectDeptIds = v
	return s
}

func (s *UpdateEmpAttributeHideBySceneSettingRequest) SetObjectTagIds(v []*int64) *UpdateEmpAttributeHideBySceneSettingRequest {
	s.ObjectTagIds = v
	return s
}

func (s *UpdateEmpAttributeHideBySceneSettingRequest) SetObjectUserIds(v []*string) *UpdateEmpAttributeHideBySceneSettingRequest {
	s.ObjectUserIds = v
	return s
}

func (s *UpdateEmpAttributeHideBySceneSettingRequest) SetProfileSceneConfig(v *UpdateEmpAttributeHideBySceneSettingRequestProfileSceneConfig) *UpdateEmpAttributeHideBySceneSettingRequest {
	s.ProfileSceneConfig = v
	return s
}

func (s *UpdateEmpAttributeHideBySceneSettingRequest) SetSearchSceneConfig(v *UpdateEmpAttributeHideBySceneSettingRequestSearchSceneConfig) *UpdateEmpAttributeHideBySceneSettingRequest {
	s.SearchSceneConfig = v
	return s
}

type UpdateEmpAttributeHideBySceneSettingRequestChatSubtitleConfig struct {
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
}

func (s UpdateEmpAttributeHideBySceneSettingRequestChatSubtitleConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateEmpAttributeHideBySceneSettingRequestChatSubtitleConfig) GoString() string {
	return s.String()
}

func (s *UpdateEmpAttributeHideBySceneSettingRequestChatSubtitleConfig) SetActive(v bool) *UpdateEmpAttributeHideBySceneSettingRequestChatSubtitleConfig {
	s.Active = &v
	return s
}

type UpdateEmpAttributeHideBySceneSettingRequestProfileSceneConfig struct {
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
}

func (s UpdateEmpAttributeHideBySceneSettingRequestProfileSceneConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateEmpAttributeHideBySceneSettingRequestProfileSceneConfig) GoString() string {
	return s.String()
}

func (s *UpdateEmpAttributeHideBySceneSettingRequestProfileSceneConfig) SetActive(v bool) *UpdateEmpAttributeHideBySceneSettingRequestProfileSceneConfig {
	s.Active = &v
	return s
}

type UpdateEmpAttributeHideBySceneSettingRequestSearchSceneConfig struct {
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
}

func (s UpdateEmpAttributeHideBySceneSettingRequestSearchSceneConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateEmpAttributeHideBySceneSettingRequestSearchSceneConfig) GoString() string {
	return s.String()
}

func (s *UpdateEmpAttributeHideBySceneSettingRequestSearchSceneConfig) SetActive(v bool) *UpdateEmpAttributeHideBySceneSettingRequestSearchSceneConfig {
	s.Active = &v
	return s
}

type UpdateEmpAttributeHideBySceneSettingResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateEmpAttributeHideBySceneSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEmpAttributeHideBySceneSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEmpAttributeHideBySceneSettingResponseBody) SetSuccess(v bool) *UpdateEmpAttributeHideBySceneSettingResponseBody {
	s.Success = &v
	return s
}

type UpdateEmpAttributeHideBySceneSettingResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateEmpAttributeHideBySceneSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateEmpAttributeHideBySceneSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEmpAttributeHideBySceneSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateEmpAttributeHideBySceneSettingResponse) SetHeaders(v map[string]*string) *UpdateEmpAttributeHideBySceneSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateEmpAttributeHideBySceneSettingResponse) SetStatusCode(v int32) *UpdateEmpAttributeHideBySceneSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEmpAttributeHideBySceneSettingResponse) SetBody(v *UpdateEmpAttributeHideBySceneSettingResponseBody) *UpdateEmpAttributeHideBySceneSettingResponse {
	s.Body = v
	return s
}

type UpdateManagementGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateManagementGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateManagementGroupHeaders) GoString() string {
	return s.String()
}

func (s *UpdateManagementGroupHeaders) SetCommonHeaders(v map[string]*string) *UpdateManagementGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateManagementGroupHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateManagementGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateManagementGroupRequest struct {
	GroupName   *string                                `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Members     []*UpdateManagementGroupRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	ResourceIds []*string                              `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	Scope       *UpdateManagementGroupRequestScope     `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
}

func (s UpdateManagementGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateManagementGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateManagementGroupRequest) SetGroupName(v string) *UpdateManagementGroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateManagementGroupRequest) SetMembers(v []*UpdateManagementGroupRequestMembers) *UpdateManagementGroupRequest {
	s.Members = v
	return s
}

func (s *UpdateManagementGroupRequest) SetResourceIds(v []*string) *UpdateManagementGroupRequest {
	s.ResourceIds = v
	return s
}

func (s *UpdateManagementGroupRequest) SetScope(v *UpdateManagementGroupRequestScope) *UpdateManagementGroupRequest {
	s.Scope = v
	return s
}

type UpdateManagementGroupRequestMembers struct {
	MemberId   *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
}

func (s UpdateManagementGroupRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s UpdateManagementGroupRequestMembers) GoString() string {
	return s.String()
}

func (s *UpdateManagementGroupRequestMembers) SetMemberId(v string) *UpdateManagementGroupRequestMembers {
	s.MemberId = &v
	return s
}

func (s *UpdateManagementGroupRequestMembers) SetMemberType(v string) *UpdateManagementGroupRequestMembers {
	s.MemberType = &v
	return s
}

type UpdateManagementGroupRequestScope struct {
	DeptIds   []*int64 `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	ScopeType *int32   `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
}

func (s UpdateManagementGroupRequestScope) String() string {
	return tea.Prettify(s)
}

func (s UpdateManagementGroupRequestScope) GoString() string {
	return s.String()
}

func (s *UpdateManagementGroupRequestScope) SetDeptIds(v []*int64) *UpdateManagementGroupRequestScope {
	s.DeptIds = v
	return s
}

func (s *UpdateManagementGroupRequestScope) SetScopeType(v int32) *UpdateManagementGroupRequestScope {
	s.ScopeType = &v
	return s
}

type UpdateManagementGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateManagementGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateManagementGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateManagementGroupResponse) SetHeaders(v map[string]*string) *UpdateManagementGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateManagementGroupResponse) SetStatusCode(v int32) *UpdateManagementGroupResponse {
	s.StatusCode = &v
	return s
}

type UpdateSeniorSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateSeniorSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateSeniorSettingHeaders) GoString() string {
	return s.String()
}

func (s *UpdateSeniorSettingHeaders) SetCommonHeaders(v map[string]*string) *UpdateSeniorSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateSeniorSettingHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateSeniorSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateSeniorSettingRequest struct {
	Open           *bool     `json:"open,omitempty" xml:"open,omitempty"`
	PermitDeptIds  []*int64  `json:"permitDeptIds,omitempty" xml:"permitDeptIds,omitempty" type:"Repeated"`
	PermitStaffIds []*string `json:"permitStaffIds,omitempty" xml:"permitStaffIds,omitempty" type:"Repeated"`
	PermitTagIds   []*int64  `json:"permitTagIds,omitempty" xml:"permitTagIds,omitempty" type:"Repeated"`
	ProtectScenes  []*string `json:"protectScenes,omitempty" xml:"protectScenes,omitempty" type:"Repeated"`
	SeniorStaffId  *string   `json:"seniorStaffId,omitempty" xml:"seniorStaffId,omitempty"`
}

func (s UpdateSeniorSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSeniorSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateSeniorSettingRequest) SetOpen(v bool) *UpdateSeniorSettingRequest {
	s.Open = &v
	return s
}

func (s *UpdateSeniorSettingRequest) SetPermitDeptIds(v []*int64) *UpdateSeniorSettingRequest {
	s.PermitDeptIds = v
	return s
}

func (s *UpdateSeniorSettingRequest) SetPermitStaffIds(v []*string) *UpdateSeniorSettingRequest {
	s.PermitStaffIds = v
	return s
}

func (s *UpdateSeniorSettingRequest) SetPermitTagIds(v []*int64) *UpdateSeniorSettingRequest {
	s.PermitTagIds = v
	return s
}

func (s *UpdateSeniorSettingRequest) SetProtectScenes(v []*string) *UpdateSeniorSettingRequest {
	s.ProtectScenes = v
	return s
}

func (s *UpdateSeniorSettingRequest) SetSeniorStaffId(v string) *UpdateSeniorSettingRequest {
	s.SeniorStaffId = &v
	return s
}

type UpdateSeniorSettingResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateSeniorSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSeniorSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateSeniorSettingResponse) SetHeaders(v map[string]*string) *UpdateSeniorSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateSeniorSettingResponse) SetStatusCode(v int32) *UpdateSeniorSettingResponse {
	s.StatusCode = &v
	return s
}

type UpdateUserOwnnessHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateUserOwnnessHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserOwnnessHeaders) GoString() string {
	return s.String()
}

func (s *UpdateUserOwnnessHeaders) SetCommonHeaders(v map[string]*string) *UpdateUserOwnnessHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateUserOwnnessHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateUserOwnnessHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateUserOwnnessRequest struct {
	DeletedFlag *int32 `json:"deletedFlag,omitempty" xml:"deletedFlag,omitempty"`
	EndTime     *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Id          *int64 `json:"id,omitempty" xml:"id,omitempty"`
	OwnenssType *int32 `json:"ownenssType,omitempty" xml:"ownenssType,omitempty"`
	StartTime   *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s UpdateUserOwnnessRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserOwnnessRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserOwnnessRequest) SetDeletedFlag(v int32) *UpdateUserOwnnessRequest {
	s.DeletedFlag = &v
	return s
}

func (s *UpdateUserOwnnessRequest) SetEndTime(v int64) *UpdateUserOwnnessRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateUserOwnnessRequest) SetId(v int64) *UpdateUserOwnnessRequest {
	s.Id = &v
	return s
}

func (s *UpdateUserOwnnessRequest) SetOwnenssType(v int32) *UpdateUserOwnnessRequest {
	s.OwnenssType = &v
	return s
}

func (s *UpdateUserOwnnessRequest) SetStartTime(v int64) *UpdateUserOwnnessRequest {
	s.StartTime = &v
	return s
}

type UpdateUserOwnnessResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateUserOwnnessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserOwnnessResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserOwnnessResponseBody) SetResult(v string) *UpdateUserOwnnessResponseBody {
	s.Result = &v
	return s
}

type UpdateUserOwnnessResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateUserOwnnessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateUserOwnnessResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserOwnnessResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserOwnnessResponse) SetHeaders(v map[string]*string) *UpdateUserOwnnessResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserOwnnessResponse) SetStatusCode(v int32) *UpdateUserOwnnessResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserOwnnessResponse) SetBody(v *UpdateUserOwnnessResponseBody) *UpdateUserOwnnessResponse {
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

func (client *Client) AddContactHideBySceneSettingWithOptions(request *AddContactHideBySceneSettingRequest, headers *AddContactHideBySceneSettingHeaders, runtime *util.RuntimeOptions) (_result *AddContactHideBySceneSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeDeptIds)) {
		body["excludeDeptIds"] = request.ExcludeDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeTagIds)) {
		body["excludeTagIds"] = request.ExcludeTagIds
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeUserIds)) {
		body["excludeUserIds"] = request.ExcludeUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NodeListSceneConfig)) {
		body["nodeListSceneConfig"] = request.NodeListSceneConfig
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectDeptIds)) {
		body["objectDeptIds"] = request.ObjectDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectTagIds)) {
		body["objectTagIds"] = request.ObjectTagIds
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectUserIds)) {
		body["objectUserIds"] = request.ObjectUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProfileSceneConfig)) {
		body["profileSceneConfig"] = request.ProfileSceneConfig
	}

	if !tea.BoolValue(util.IsUnset(request.SearchSceneConfig)) {
		body["searchSceneConfig"] = request.SearchSceneConfig
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
		Action:      tea.String("AddContactHideBySceneSetting"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/organizations/hides/settings"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddContactHideBySceneSettingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddContactHideBySceneSetting(request *AddContactHideBySceneSettingRequest) (_result *AddContactHideBySceneSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddContactHideBySceneSettingHeaders{}
	_result = &AddContactHideBySceneSettingResponse{}
	_body, _err := client.AddContactHideBySceneSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddEmpAttributeHideBySceneSettingWithOptions(request *AddEmpAttributeHideBySceneSettingRequest, headers *AddEmpAttributeHideBySceneSettingHeaders, runtime *util.RuntimeOptions) (_result *AddEmpAttributeHideBySceneSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChatSubtitleConfig)) {
		body["chatSubtitleConfig"] = request.ChatSubtitleConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeDeptIds)) {
		body["excludeDeptIds"] = request.ExcludeDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeTagIds)) {
		body["excludeTagIds"] = request.ExcludeTagIds
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeUserIds)) {
		body["excludeUserIds"] = request.ExcludeUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.HideFields)) {
		body["hideFields"] = request.HideFields
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectDeptIds)) {
		body["objectDeptIds"] = request.ObjectDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectTagIds)) {
		body["objectTagIds"] = request.ObjectTagIds
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectUserIds)) {
		body["objectUserIds"] = request.ObjectUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProfileSceneConfig)) {
		body["profileSceneConfig"] = request.ProfileSceneConfig
	}

	if !tea.BoolValue(util.IsUnset(request.SearchSceneConfig)) {
		body["searchSceneConfig"] = request.SearchSceneConfig
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
		Action:      tea.String("AddEmpAttributeHideBySceneSetting"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/empAttributes/hides/settings"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddEmpAttributeHideBySceneSettingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddEmpAttributeHideBySceneSetting(request *AddEmpAttributeHideBySceneSettingRequest) (_result *AddEmpAttributeHideBySceneSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddEmpAttributeHideBySceneSettingHeaders{}
	_result = &AddEmpAttributeHideBySceneSettingResponse{}
	_body, _err := client.AddEmpAttributeHideBySceneSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AnnualCertificationAuditWithOptions(request *AnnualCertificationAuditRequest, headers *AnnualCertificationAuditHeaders, runtime *util.RuntimeOptions) (_result *AnnualCertificationAuditResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicantMobile)) {
		body["applicantMobile"] = request.ApplicantMobile
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicantName)) {
		body["applicantName"] = request.ApplicantName
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationLetter)) {
		body["applicationLetter"] = request.ApplicationLetter
	}

	if !tea.BoolValue(util.IsUnset(request.AuthStatus)) {
		body["authStatus"] = request.AuthStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CertificateType)) {
		body["certificateType"] = request.CertificateType
	}

	if !tea.BoolValue(util.IsUnset(request.CorpName)) {
		body["corpName"] = request.CorpName
	}

	if !tea.BoolValue(util.IsUnset(request.DepositaryBank)) {
		body["depositaryBank"] = request.DepositaryBank
	}

	if !tea.BoolValue(util.IsUnset(request.Extension)) {
		body["extension"] = request.Extension
	}

	if !tea.BoolValue(util.IsUnset(request.LegalPerson)) {
		body["legalPerson"] = request.LegalPerson
	}

	if !tea.BoolValue(util.IsUnset(request.LicenseNumber)) {
		body["licenseNumber"] = request.LicenseNumber
	}

	if !tea.BoolValue(util.IsUnset(request.LicenseUrl)) {
		body["licenseUrl"] = request.LicenseUrl
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		body["orderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.PublicAccount)) {
		body["publicAccount"] = request.PublicAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ReasonCode)) {
		body["reasonCode"] = request.ReasonCode
	}

	if !tea.BoolValue(util.IsUnset(request.ReasonMsg)) {
		body["reasonMsg"] = request.ReasonMsg
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["tag"] = request.Tag
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
		Action:      tea.String("AnnualCertificationAudit"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/organizations/authorities/audit"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AnnualCertificationAuditResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AnnualCertificationAudit(request *AnnualCertificationAuditRequest) (_result *AnnualCertificationAuditResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AnnualCertificationAuditHeaders{}
	_result = &AnnualCertificationAuditResponse{}
	_body, _err := client.AnnualCertificationAuditWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchApproveUnionApplyWithOptions(request *BatchApproveUnionApplyRequest, headers *BatchApproveUnionApplyHeaders, runtime *util.RuntimeOptions) (_result *BatchApproveUnionApplyResponse, _err error) {
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
		Action:      tea.String("BatchApproveUnionApply"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/cooperateCorps/unionApplications/approve"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchApproveUnionApplyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchApproveUnionApply(request *BatchApproveUnionApplyRequest) (_result *BatchApproveUnionApplyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchApproveUnionApplyHeaders{}
	_result = &BatchApproveUnionApplyResponse{}
	_body, _err := client.BatchApproveUnionApplyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeMainAdminWithOptions(request *ChangeMainAdminRequest, headers *ChangeMainAdminHeaders, runtime *util.RuntimeOptions) (_result *ChangeMainAdminResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EffectCorpId)) {
		body["effectCorpId"] = request.EffectCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceUserId)) {
		body["sourceUserId"] = request.SourceUserId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetUserId)) {
		body["targetUserId"] = request.TargetUserId
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
		Action:      tea.String("ChangeMainAdmin"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/orgAccounts/mainAdministrators/change"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("none"),
	}
	_result = &ChangeMainAdminResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeMainAdmin(request *ChangeMainAdminRequest) (_result *ChangeMainAdminResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ChangeMainAdminHeaders{}
	_result = &ChangeMainAdminResponse{}
	_body, _err := client.ChangeMainAdminWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCooperateOrgWithOptions(request *CreateCooperateOrgRequest, headers *CreateCooperateOrgHeaders, runtime *util.RuntimeOptions) (_result *CreateCooperateOrgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IndustryCode)) {
		body["industryCode"] = request.IndustryCode
	}

	if !tea.BoolValue(util.IsUnset(request.LogoMediaId)) {
		body["logoMediaId"] = request.LogoMediaId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		body["orgName"] = request.OrgName
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
		Action:      tea.String("CreateCooperateOrg"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/cooperateCorps"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCooperateOrgResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCooperateOrg(request *CreateCooperateOrgRequest) (_result *CreateCooperateOrgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateCooperateOrgHeaders{}
	_result = &CreateCooperateOrgResponse{}
	_body, _err := client.CreateCooperateOrgWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateManagementGroupWithOptions(request *CreateManagementGroupRequest, headers *CreateManagementGroupHeaders, runtime *util.RuntimeOptions) (_result *CreateManagementGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		body["resourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["scope"] = request.Scope
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
		Action:      tea.String("CreateManagementGroup"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/managementGroups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateManagementGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateManagementGroup(request *CreateManagementGroupRequest) (_result *CreateManagementGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateManagementGroupHeaders{}
	_result = &CreateManagementGroupResponse{}
	_body, _err := client.CreateManagementGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSecondaryManagementGroupWithOptions(request *CreateSecondaryManagementGroupRequest, headers *CreateSecondaryManagementGroupHeaders, runtime *util.RuntimeOptions) (_result *CreateSecondaryManagementGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		body["resourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["scope"] = request.Scope
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
		Action:      tea.String("CreateSecondaryManagementGroup"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/secondaryAdministrators/managementGroups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSecondaryManagementGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSecondaryManagementGroup(request *CreateSecondaryManagementGroupRequest) (_result *CreateSecondaryManagementGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateSecondaryManagementGroupHeaders{}
	_result = &CreateSecondaryManagementGroupResponse{}
	_body, _err := client.CreateSecondaryManagementGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteContactHideBySceneSettingWithOptions(settingId *string, headers *DeleteContactHideBySceneSettingHeaders, runtime *util.RuntimeOptions) (_result *DeleteContactHideBySceneSettingResponse, _err error) {
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
		Action:      tea.String("DeleteContactHideBySceneSetting"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/organizations/hides/settings/" + tea.StringValue(settingId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteContactHideBySceneSettingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteContactHideBySceneSetting(settingId *string) (_result *DeleteContactHideBySceneSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteContactHideBySceneSettingHeaders{}
	_result = &DeleteContactHideBySceneSettingResponse{}
	_body, _err := client.DeleteContactHideBySceneSettingWithOptions(settingId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteContactHideSettingWithOptions(settingId *string, headers *DeleteContactHideSettingHeaders, runtime *util.RuntimeOptions) (_result *DeleteContactHideSettingResponse, _err error) {
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
		Action:      tea.String("DeleteContactHideSetting"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/contactHideSettings/" + tea.StringValue(settingId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteContactHideSettingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteContactHideSetting(settingId *string) (_result *DeleteContactHideSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteContactHideSettingHeaders{}
	_result = &DeleteContactHideSettingResponse{}
	_body, _err := client.DeleteContactHideSettingWithOptions(settingId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteContactRestrictSettingWithOptions(settingId *string, headers *DeleteContactRestrictSettingHeaders, runtime *util.RuntimeOptions) (_result *DeleteContactRestrictSettingResponse, _err error) {
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
		Action:      tea.String("DeleteContactRestrictSetting"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/restrictions/settings/" + tea.StringValue(settingId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteContactRestrictSettingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteContactRestrictSetting(settingId *string) (_result *DeleteContactRestrictSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteContactRestrictSettingHeaders{}
	_result = &DeleteContactRestrictSettingResponse{}
	_body, _err := client.DeleteContactRestrictSettingWithOptions(settingId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEmpAttributeHideBySceneSettingWithOptions(settingId *string, headers *DeleteEmpAttributeHideBySceneSettingHeaders, runtime *util.RuntimeOptions) (_result *DeleteEmpAttributeHideBySceneSettingResponse, _err error) {
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
		Action:      tea.String("DeleteEmpAttributeHideBySceneSetting"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/empAttributes/hides/settings/" + tea.StringValue(settingId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEmpAttributeHideBySceneSettingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEmpAttributeHideBySceneSetting(settingId *string) (_result *DeleteEmpAttributeHideBySceneSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteEmpAttributeHideBySceneSettingHeaders{}
	_result = &DeleteEmpAttributeHideBySceneSettingResponse{}
	_body, _err := client.DeleteEmpAttributeHideBySceneSettingWithOptions(settingId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEmpAttributeVisibilityWithOptions(settingId *string, headers *DeleteEmpAttributeVisibilityHeaders, runtime *util.RuntimeOptions) (_result *DeleteEmpAttributeVisibilityResponse, _err error) {
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
		Action:      tea.String("DeleteEmpAttributeVisibility"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/staffAttributes/visibilitySettings/" + tea.StringValue(settingId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEmpAttributeVisibilityResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEmpAttributeVisibility(settingId *string) (_result *DeleteEmpAttributeVisibilityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteEmpAttributeVisibilityHeaders{}
	_result = &DeleteEmpAttributeVisibilityResponse{}
	_body, _err := client.DeleteEmpAttributeVisibilityWithOptions(settingId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteManagementGroupWithOptions(groupId *string, headers *DeleteManagementGroupHeaders, runtime *util.RuntimeOptions) (_result *DeleteManagementGroupResponse, _err error) {
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
		Action:      tea.String("DeleteManagementGroup"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/managementGroups/" + tea.StringValue(groupId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteManagementGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteManagementGroup(groupId *string) (_result *DeleteManagementGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteManagementGroupHeaders{}
	_result = &DeleteManagementGroupResponse{}
	_body, _err := client.DeleteManagementGroupWithOptions(groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetApplyInviteInfoWithOptions(request *GetApplyInviteInfoRequest, headers *GetApplyInviteInfoHeaders, runtime *util.RuntimeOptions) (_result *GetApplyInviteInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.InviterUserId)) {
		query["inviterUserId"] = request.InviterUserId
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
		Action:      tea.String("GetApplyInviteInfo"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/invites/infos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetApplyInviteInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetApplyInviteInfo(request *GetApplyInviteInfoRequest) (_result *GetApplyInviteInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetApplyInviteInfoHeaders{}
	_result = &GetApplyInviteInfoResponse{}
	_body, _err := client.GetApplyInviteInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBranchAuthDataWithOptions(request *GetBranchAuthDataRequest, headers *GetBranchAuthDataHeaders, runtime *util.RuntimeOptions) (_result *GetBranchAuthDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BranchCorpId)) {
		query["branchCorpId"] = request.BranchCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["code"] = request.Code
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
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBranchAuthData"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/branchAuthDatas/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBranchAuthDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBranchAuthData(request *GetBranchAuthDataRequest) (_result *GetBranchAuthDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetBranchAuthDataHeaders{}
	_result = &GetBranchAuthDataResponse{}
	_body, _err := client.GetBranchAuthDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCardInUserHolderWithOptions(cardId *string, headers *GetCardInUserHolderHeaders, runtime *util.RuntimeOptions) (_result *GetCardInUserHolderResponse, _err error) {
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
		Action:      tea.String("GetCardInUserHolder"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/cards/holders/infos/" + tea.StringValue(cardId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCardInUserHolderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCardInUserHolder(cardId *string) (_result *GetCardInUserHolderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCardInUserHolderHeaders{}
	_result = &GetCardInUserHolderResponse{}
	_body, _err := client.GetCardInUserHolderWithOptions(cardId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCardInfoWithOptions(cardId *string, headers *GetCardInfoHeaders, runtime *util.RuntimeOptions) (_result *GetCardInfoResponse, _err error) {
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
		Action:      tea.String("GetCardInfo"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/cards/infos/" + tea.StringValue(cardId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCardInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCardInfo(cardId *string) (_result *GetCardInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCardInfoHeaders{}
	_result = &GetCardInfoResponse{}
	_body, _err := client.GetCardInfoWithOptions(cardId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetContactHideBySceneSettingWithOptions(settingId *string, headers *GetContactHideBySceneSettingHeaders, runtime *util.RuntimeOptions) (_result *GetContactHideBySceneSettingResponse, _err error) {
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
		Action:      tea.String("GetContactHideBySceneSetting"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/organizations/hides/settings/" + tea.StringValue(settingId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetContactHideBySceneSettingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetContactHideBySceneSetting(settingId *string) (_result *GetContactHideBySceneSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetContactHideBySceneSettingHeaders{}
	_result = &GetContactHideBySceneSettingResponse{}
	_body, _err := client.GetContactHideBySceneSettingWithOptions(settingId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCooperateOrgInviteInfoWithOptions(cooperateCorpId *string, headers *GetCooperateOrgInviteInfoHeaders, runtime *util.RuntimeOptions) (_result *GetCooperateOrgInviteInfoResponse, _err error) {
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
		Action:      tea.String("GetCooperateOrgInviteInfo"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/cooperateCorps/" + tea.StringValue(cooperateCorpId) + "/inviteInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCooperateOrgInviteInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCooperateOrgInviteInfo(cooperateCorpId *string) (_result *GetCooperateOrgInviteInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCooperateOrgInviteInfoHeaders{}
	_result = &GetCooperateOrgInviteInfoResponse{}
	_body, _err := client.GetCooperateOrgInviteInfoWithOptions(cooperateCorpId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCorpCardStyleListWithOptions(headers *GetCorpCardStyleListHeaders, runtime *util.RuntimeOptions) (_result *GetCorpCardStyleListResponse, _err error) {
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
		Action:      tea.String("GetCorpCardStyleList"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/cards/styles/lists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCorpCardStyleListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCorpCardStyleList() (_result *GetCorpCardStyleListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCorpCardStyleListHeaders{}
	_result = &GetCorpCardStyleListResponse{}
	_body, _err := client.GetCorpCardStyleListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDingIdByMigrationDingIdWithOptions(request *GetDingIdByMigrationDingIdRequest, headers *GetDingIdByMigrationDingIdHeaders, runtime *util.RuntimeOptions) (_result *GetDingIdByMigrationDingIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MigrationDingId)) {
		query["migrationDingId"] = request.MigrationDingId
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
		Action:      tea.String("GetDingIdByMigrationDingId"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/orgAccount/getDingIdByMigrationDingIds"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDingIdByMigrationDingIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDingIdByMigrationDingId(request *GetDingIdByMigrationDingIdRequest) (_result *GetDingIdByMigrationDingIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDingIdByMigrationDingIdHeaders{}
	_result = &GetDingIdByMigrationDingIdResponse{}
	_body, _err := client.GetDingIdByMigrationDingIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEmpAttributeHideBySceneSettingWithOptions(settingId *string, headers *GetEmpAttributeHideBySceneSettingHeaders, runtime *util.RuntimeOptions) (_result *GetEmpAttributeHideBySceneSettingResponse, _err error) {
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
		Action:      tea.String("GetEmpAttributeHideBySceneSetting"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/empAttributes/hides/settings/" + tea.StringValue(settingId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEmpAttributeHideBySceneSettingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEmpAttributeHideBySceneSetting(settingId *string) (_result *GetEmpAttributeHideBySceneSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetEmpAttributeHideBySceneSettingHeaders{}
	_result = &GetEmpAttributeHideBySceneSettingResponse{}
	_body, _err := client.GetEmpAttributeHideBySceneSettingWithOptions(settingId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLatestDingIndexWithOptions(headers *GetLatestDingIndexHeaders, runtime *util.RuntimeOptions) (_result *GetLatestDingIndexResponse, _err error) {
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
		Action:      tea.String("GetLatestDingIndex"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/dingIndexs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLatestDingIndexResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLatestDingIndex() (_result *GetLatestDingIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetLatestDingIndexHeaders{}
	_result = &GetLatestDingIndexResponse{}
	_body, _err := client.GetLatestDingIndexWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMigrationDingIdByDingIdWithOptions(request *GetMigrationDingIdByDingIdRequest, headers *GetMigrationDingIdByDingIdHeaders, runtime *util.RuntimeOptions) (_result *GetMigrationDingIdByDingIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingId)) {
		query["dingId"] = request.DingId
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
		Action:      tea.String("GetMigrationDingIdByDingId"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/orgAccount/getMigrationDingIdByDingIds"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMigrationDingIdByDingIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMigrationDingIdByDingId(request *GetMigrationDingIdByDingIdRequest) (_result *GetMigrationDingIdByDingIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMigrationDingIdByDingIdHeaders{}
	_result = &GetMigrationDingIdByDingIdResponse{}
	_body, _err := client.GetMigrationDingIdByDingIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMigrationUnionIdByUnionIdWithOptions(request *GetMigrationUnionIdByUnionIdRequest, headers *GetMigrationUnionIdByUnionIdHeaders, runtime *util.RuntimeOptions) (_result *GetMigrationUnionIdByUnionIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
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
		Action:      tea.String("GetMigrationUnionIdByUnionId"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/orgAccount/getMigrationUnionIdByUnionIds"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMigrationUnionIdByUnionIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMigrationUnionIdByUnionId(request *GetMigrationUnionIdByUnionIdRequest) (_result *GetMigrationUnionIdByUnionIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMigrationUnionIdByUnionIdHeaders{}
	_result = &GetMigrationUnionIdByUnionIdResponse{}
	_body, _err := client.GetMigrationUnionIdByUnionIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOrgAuthInfoWithOptions(request *GetOrgAuthInfoRequest, headers *GetOrgAuthInfoHeaders, runtime *util.RuntimeOptions) (_result *GetOrgAuthInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		query["targetCorpId"] = request.TargetCorpId
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
		Action:      tea.String("GetOrgAuthInfo"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/organizations/authInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOrgAuthInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOrgAuthInfo(request *GetOrgAuthInfoRequest) (_result *GetOrgAuthInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOrgAuthInfoHeaders{}
	_result = &GetOrgAuthInfoResponse{}
	_body, _err := client.GetOrgAuthInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTranslateFileJobResultWithOptions(request *GetTranslateFileJobResultRequest, headers *GetTranslateFileJobResultHeaders, runtime *util.RuntimeOptions) (_result *GetTranslateFileJobResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["jobId"] = request.JobId
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
		Action:      tea.String("GetTranslateFileJobResult"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/files/translateResults"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTranslateFileJobResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTranslateFileJobResult(request *GetTranslateFileJobResultRequest) (_result *GetTranslateFileJobResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTranslateFileJobResultHeaders{}
	_result = &GetTranslateFileJobResultResponse{}
	_body, _err := client.GetTranslateFileJobResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUnionIdByMigrationUnionIdWithOptions(request *GetUnionIdByMigrationUnionIdRequest, headers *GetUnionIdByMigrationUnionIdHeaders, runtime *util.RuntimeOptions) (_result *GetUnionIdByMigrationUnionIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MigrationUnionId)) {
		query["migrationUnionId"] = request.MigrationUnionId
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
		Action:      tea.String("GetUnionIdByMigrationUnionId"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/orgAccount/getUnionIdByMigrationUnionIds"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUnionIdByMigrationUnionIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUnionIdByMigrationUnionId(request *GetUnionIdByMigrationUnionIdRequest) (_result *GetUnionIdByMigrationUnionIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUnionIdByMigrationUnionIdHeaders{}
	_result = &GetUnionIdByMigrationUnionIdResponse{}
	_body, _err := client.GetUnionIdByMigrationUnionIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserWithOptions(unionId *string, headers *GetUserHeaders, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
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
		Action:      tea.String("GetUser"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/users/" + tea.StringValue(unionId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUser(unionId *string) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserHeaders{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(unionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserCardHolderListWithOptions(request *GetUserCardHolderListRequest, headers *GetUserCardHolderListHeaders, runtime *util.RuntimeOptions) (_result *GetUserCardHolderListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserCardHolderList"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/cards/holders/lists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserCardHolderListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserCardHolderList(request *GetUserCardHolderListRequest) (_result *GetUserCardHolderListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserCardHolderListHeaders{}
	_result = &GetUserCardHolderListResponse{}
	_body, _err := client.GetUserCardHolderListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IsFriendWithOptions(request *IsFriendRequest, headers *IsFriendHeaders, runtime *util.RuntimeOptions) (_result *IsFriendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MobileNo)) {
		body["mobileNo"] = request.MobileNo
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
		Action:      tea.String("IsFriend"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/relationships/friends/judge"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &IsFriendResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IsFriend(request *IsFriendRequest) (_result *IsFriendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IsFriendHeaders{}
	_result = &IsFriendResponse{}
	_body, _err := client.IsFriendWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IsvCardEventPushWithOptions(request *IsvCardEventPushRequest, headers *IsvCardEventPushHeaders, runtime *util.RuntimeOptions) (_result *IsvCardEventPushResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsvCardId)) {
		query["isvCardId"] = request.IsvCardId
	}

	if !tea.BoolValue(util.IsUnset(request.IsvToken)) {
		query["isvToken"] = request.IsvToken
	}

	if !tea.BoolValue(util.IsUnset(request.IsvUid)) {
		query["isvUid"] = request.IsvUid
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventParams)) {
		body["eventParams"] = request.EventParams
	}

	if !tea.BoolValue(util.IsUnset(request.EventType)) {
		body["eventType"] = request.EventType
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
		Action:      tea.String("IsvCardEventPush"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/cards/events/push"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &IsvCardEventPushResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IsvCardEventPush(request *IsvCardEventPushRequest) (_result *IsvCardEventPushResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IsvCardEventPushHeaders{}
	_result = &IsvCardEventPushResponse{}
	_body, _err := client.IsvCardEventPushWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBasicRoleInPageWithOptions(request *ListBasicRoleInPageRequest, headers *ListBasicRoleInPageHeaders, runtime *util.RuntimeOptions) (_result *ListBasicRoleInPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		query["agentId"] = request.AgentId
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
	}
	params := &openapi.Params{
		Action:      tea.String("ListBasicRoleInPage"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/rbac/administrativeGroups/baseInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBasicRoleInPageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBasicRoleInPage(request *ListBasicRoleInPageRequest) (_result *ListBasicRoleInPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListBasicRoleInPageHeaders{}
	_result = &ListBasicRoleInPageResponse{}
	_body, _err := client.ListBasicRoleInPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListContactHideSettingsWithOptions(request *ListContactHideSettingsRequest, headers *ListContactHideSettingsHeaders, runtime *util.RuntimeOptions) (_result *ListContactHideSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
	}
	params := &openapi.Params{
		Action:      tea.String("ListContactHideSettings"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/contactHideSettings"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListContactHideSettingsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListContactHideSettings(request *ListContactHideSettingsRequest) (_result *ListContactHideSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListContactHideSettingsHeaders{}
	_result = &ListContactHideSettingsResponse{}
	_body, _err := client.ListContactHideSettingsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListContactRestrictSettingWithOptions(request *ListContactRestrictSettingRequest, headers *ListContactRestrictSettingHeaders, runtime *util.RuntimeOptions) (_result *ListContactRestrictSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
	}
	params := &openapi.Params{
		Action:      tea.String("ListContactRestrictSetting"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/restrictions/settings"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListContactRestrictSettingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListContactRestrictSetting(request *ListContactRestrictSettingRequest) (_result *ListContactRestrictSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListContactRestrictSettingHeaders{}
	_result = &ListContactRestrictSettingResponse{}
	_body, _err := client.ListContactRestrictSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEmpAttributeVisibilityWithOptions(request *ListEmpAttributeVisibilityRequest, headers *ListEmpAttributeVisibilityHeaders, runtime *util.RuntimeOptions) (_result *ListEmpAttributeVisibilityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
	}
	params := &openapi.Params{
		Action:      tea.String("ListEmpAttributeVisibility"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/staffAttributes/visibilitySettings"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEmpAttributeVisibilityResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEmpAttributeVisibility(request *ListEmpAttributeVisibilityRequest) (_result *ListEmpAttributeVisibilityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListEmpAttributeVisibilityHeaders{}
	_result = &ListEmpAttributeVisibilityResponse{}
	_body, _err := client.ListEmpAttributeVisibilityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEmpLeaveRecordsWithOptions(request *ListEmpLeaveRecordsRequest, headers *ListEmpLeaveRecordsHeaders, runtime *util.RuntimeOptions) (_result *ListEmpLeaveRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
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
		Action:      tea.String("ListEmpLeaveRecords"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/empLeaveRecords"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEmpLeaveRecordsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEmpLeaveRecords(request *ListEmpLeaveRecordsRequest) (_result *ListEmpLeaveRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListEmpLeaveRecordsHeaders{}
	_result = &ListEmpLeaveRecordsResponse{}
	_body, _err := client.ListEmpLeaveRecordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListManagementGroupsWithOptions(request *ListManagementGroupsRequest, headers *ListManagementGroupsHeaders, runtime *util.RuntimeOptions) (_result *ListManagementGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
	}
	params := &openapi.Params{
		Action:      tea.String("ListManagementGroups"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/managementGroups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListManagementGroupsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListManagementGroups(request *ListManagementGroupsRequest) (_result *ListManagementGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListManagementGroupsHeaders{}
	_result = &ListManagementGroupsResponse{}
	_body, _err := client.ListManagementGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOwnedOrgByStaffIdWithOptions(request *ListOwnedOrgByStaffIdRequest, headers *ListOwnedOrgByStaffIdHeaders, runtime *util.RuntimeOptions) (_result *ListOwnedOrgByStaffIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      tea.String("ListOwnedOrgByStaffId"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/orgAccounts/ownedOrganizations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOwnedOrgByStaffIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOwnedOrgByStaffId(request *ListOwnedOrgByStaffIdRequest) (_result *ListOwnedOrgByStaffIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListOwnedOrgByStaffIdHeaders{}
	_result = &ListOwnedOrgByStaffIdResponse{}
	_body, _err := client.ListOwnedOrgByStaffIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSeniorSettingsWithOptions(request *ListSeniorSettingsRequest, headers *ListSeniorSettingsHeaders, runtime *util.RuntimeOptions) (_result *ListSeniorSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SeniorStaffId)) {
		query["seniorStaffId"] = request.SeniorStaffId
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
		Action:      tea.String("ListSeniorSettings"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/seniorSettings"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSeniorSettingsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSeniorSettings(request *ListSeniorSettingsRequest) (_result *ListSeniorSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSeniorSettingsHeaders{}
	_result = &ListSeniorSettingsResponse{}
	_body, _err := client.ListSeniorSettingsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MultiOrgPermissionGrantWithOptions(request *MultiOrgPermissionGrantRequest, headers *MultiOrgPermissionGrantHeaders, runtime *util.RuntimeOptions) (_result *MultiOrgPermissionGrantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GrantDeptIdList)) {
		body["grantDeptIdList"] = request.GrantDeptIdList
	}

	if !tea.BoolValue(util.IsUnset(request.JoinCorpId)) {
		body["joinCorpId"] = request.JoinCorpId
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
		Action:      tea.String("MultiOrgPermissionGrant"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/orgAccounts/multiOrgPermissions/auth"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("none"),
	}
	_result = &MultiOrgPermissionGrantResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MultiOrgPermissionGrant(request *MultiOrgPermissionGrantRequest) (_result *MultiOrgPermissionGrantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MultiOrgPermissionGrantHeaders{}
	_result = &MultiOrgPermissionGrantResponse{}
	_body, _err := client.MultiOrgPermissionGrantWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCardVisitorStatisticDataWithOptions(request *QueryCardVisitorStatisticDataRequest, headers *QueryCardVisitorStatisticDataHeaders, runtime *util.RuntimeOptions) (_result *QueryCardVisitorStatisticDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
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
		Action:      tea.String("QueryCardVisitorStatisticData"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/cards/visitors/statistics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCardVisitorStatisticDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCardVisitorStatisticData(request *QueryCardVisitorStatisticDataRequest) (_result *QueryCardVisitorStatisticDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCardVisitorStatisticDataHeaders{}
	_result = &QueryCardVisitorStatisticDataResponse{}
	_body, _err := client.QueryCardVisitorStatisticDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCorpStatisticDataWithOptions(request *QueryCorpStatisticDataRequest, headers *QueryCorpStatisticDataHeaders, runtime *util.RuntimeOptions) (_result *QueryCorpStatisticDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateIds)) {
		body["templateIds"] = request.TemplateIds
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
		Action:      tea.String("QueryCorpStatisticData"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/cards/templates/statistics/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCorpStatisticDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCorpStatisticData(request *QueryCorpStatisticDataRequest) (_result *QueryCorpStatisticDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCorpStatisticDataHeaders{}
	_result = &QueryCorpStatisticDataResponse{}
	_body, _err := client.QueryCorpStatisticDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCorpUserStatisticWithOptions(request *QueryCorpUserStatisticRequest, headers *QueryCorpUserStatisticHeaders, runtime *util.RuntimeOptions) (_result *QueryCorpUserStatisticResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateIds)) {
		body["templateIds"] = request.TemplateIds
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
		Action:      tea.String("QueryCorpUserStatistic"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/cards/users/statistics/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCorpUserStatisticResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCorpUserStatistic(request *QueryCorpUserStatisticRequest) (_result *QueryCorpUserStatisticResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCorpUserStatisticHeaders{}
	_result = &QueryCorpUserStatisticResponse{}
	_body, _err := client.QueryCorpUserStatisticWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryResourceManagementMembersWithOptions(resourceId *string, headers *QueryResourceManagementMembersHeaders, runtime *util.RuntimeOptions) (_result *QueryResourceManagementMembersResponse, _err error) {
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
		Action:      tea.String("QueryResourceManagementMembers"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/resources/" + tea.StringValue(resourceId) + "/managementMembers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryResourceManagementMembersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryResourceManagementMembers(resourceId *string) (_result *QueryResourceManagementMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryResourceManagementMembersHeaders{}
	_result = &QueryResourceManagementMembersResponse{}
	_body, _err := client.QueryResourceManagementMembersWithOptions(resourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryStatusWithOptions(request *QueryStatusRequest, headers *QueryStatusHeaders, runtime *util.RuntimeOptions) (_result *QueryStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      tea.String("QueryStatus"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/orgAccounts/status"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryStatus(request *QueryStatusRequest) (_result *QueryStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryStatusHeaders{}
	_result = &QueryStatusResponse{}
	_body, _err := client.QueryStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserManagementResourcesWithOptions(userId *string, headers *QueryUserManagementResourcesHeaders, runtime *util.RuntimeOptions) (_result *QueryUserManagementResourcesResponse, _err error) {
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
		Action:      tea.String("QueryUserManagementResources"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/users/" + tea.StringValue(userId) + "/managemementResources"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserManagementResourcesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserManagementResources(userId *string) (_result *QueryUserManagementResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserManagementResourcesHeaders{}
	_result = &QueryUserManagementResourcesResponse{}
	_body, _err := client.QueryUserManagementResourcesWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryVerifyResultWithOptions(request *QueryVerifyResultRequest, headers *QueryVerifyResultHeaders, runtime *util.RuntimeOptions) (_result *QueryVerifyResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VerifyId)) {
		query["verifyId"] = request.VerifyId
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
		Action:      tea.String("QueryVerifyResult"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/verifyIdentitys/verifyResults"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryVerifyResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryVerifyResult(request *QueryVerifyResultRequest) (_result *QueryVerifyResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryVerifyResultHeaders{}
	_result = &QueryVerifyResultResponse{}
	_body, _err := client.QueryVerifyResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchDepartmentWithOptions(request *SearchDepartmentRequest, headers *SearchDepartmentHeaders, runtime *util.RuntimeOptions) (_result *SearchDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		body["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.QueryWord)) {
		body["queryWord"] = request.QueryWord
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["size"] = request.Size
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
		Action:      tea.String("SearchDepartment"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/departments/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchDepartmentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchDepartment(request *SearchDepartmentRequest) (_result *SearchDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchDepartmentHeaders{}
	_result = &SearchDepartmentResponse{}
	_body, _err := client.SearchDepartmentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchUserWithOptions(request *SearchUserRequest, headers *SearchUserHeaders, runtime *util.RuntimeOptions) (_result *SearchUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FullMatchField)) {
		body["fullMatchField"] = request.FullMatchField
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		body["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.QueryWord)) {
		body["queryWord"] = request.QueryWord
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["size"] = request.Size
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
		Action:      tea.String("SearchUser"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/users/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchUser(request *SearchUserRequest) (_result *SearchUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchUserHeaders{}
	_result = &SearchUserResponse{}
	_body, _err := client.SearchUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SeparateBranchOrgWithOptions(request *SeparateBranchOrgRequest, headers *SeparateBranchOrgHeaders, runtime *util.RuntimeOptions) (_result *SeparateBranchOrgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttachDeptId)) {
		body["attachDeptId"] = request.AttachDeptId
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
		Action:      tea.String("SeparateBranchOrg"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/cooperateCorps/separate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SeparateBranchOrgResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SeparateBranchOrg(request *SeparateBranchOrgRequest) (_result *SeparateBranchOrgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SeparateBranchOrgHeaders{}
	_result = &SeparateBranchOrgResponse{}
	_body, _err := client.SeparateBranchOrgWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDisableWithOptions(request *SetDisableRequest, headers *SetDisableHeaders, runtime *util.RuntimeOptions) (_result *SetDisableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Reason)) {
		body["reason"] = request.Reason
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
		Action:      tea.String("SetDisable"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/orgAccounts/disable"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDisableResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDisable(request *SetDisableRequest) (_result *SetDisableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetDisableHeaders{}
	_result = &SetDisableResponse{}
	_body, _err := client.SetDisableWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetEnableWithOptions(request *SetEnableRequest, headers *SetEnableHeaders, runtime *util.RuntimeOptions) (_result *SetEnableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
		Action:      tea.String("SetEnable"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/orgAccounts/enable"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SetEnableResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetEnable(request *SetEnableRequest) (_result *SetEnableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetEnableHeaders{}
	_result = &SetEnableResponse{}
	_body, _err := client.SetEnableWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SignOutWithOptions(request *SignOutRequest, headers *SignOutHeaders, runtime *util.RuntimeOptions) (_result *SignOutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Reason)) {
		body["reason"] = request.Reason
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
		Action:      tea.String("SignOut"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/orgAccounts/signOut"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SignOutResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SignOut(request *SignOutRequest) (_result *SignOutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SignOutHeaders{}
	_result = &SignOutResponse{}
	_body, _err := client.SignOutWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SortUserWithOptions(request *SortUserRequest, headers *SortUserHeaders, runtime *util.RuntimeOptions) (_result *SortUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SortType)) {
		body["sortType"] = request.SortType
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdList)) {
		body["userIdList"] = request.UserIdList
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
		Action:      tea.String("SortUser"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/users/sort"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SortUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SortUser(request *SortUserRequest) (_result *SortUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SortUserHeaders{}
	_result = &SortUserResponse{}
	_body, _err := client.SortUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransformToExclusiveAccountWithOptions(request *TransformToExclusiveAccountRequest, headers *TransformToExclusiveAccountHeaders, runtime *util.RuntimeOptions) (_result *TransformToExclusiveAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdpDingTalk)) {
		body["idpDingTalk"] = request.IdpDingTalk
	}

	if !tea.BoolValue(util.IsUnset(request.InitPassword)) {
		body["initPassword"] = request.InitPassword
	}

	if !tea.BoolValue(util.IsUnset(request.LoginId)) {
		body["loginId"] = request.LoginId
	}

	if !tea.BoolValue(util.IsUnset(request.TransformType)) {
		body["transformType"] = request.TransformType
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
		Action:      tea.String("TransformToExclusiveAccount"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/orgAccount/transformToExclusiveAccounts"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TransformToExclusiveAccountResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransformToExclusiveAccount(request *TransformToExclusiveAccountRequest) (_result *TransformToExclusiveAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TransformToExclusiveAccountHeaders{}
	_result = &TransformToExclusiveAccountResponse{}
	_body, _err := client.TransformToExclusiveAccountWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TranslateFileWithOptions(request *TranslateFileRequest, headers *TranslateFileHeaders, runtime *util.RuntimeOptions) (_result *TranslateFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Medias)) {
		body["medias"] = request.Medias
	}

	if !tea.BoolValue(util.IsUnset(request.OutputFileName)) {
		body["outputFileName"] = request.OutputFileName
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
		Action:      tea.String("TranslateFile"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/files/translate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TranslateFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TranslateFile(request *TranslateFileRequest) (_result *TranslateFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TranslateFileHeaders{}
	_result = &TranslateFileResponse{}
	_body, _err := client.TranslateFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UniqueQueryUserCardWithOptions(request *UniqueQueryUserCardRequest, headers *UniqueQueryUserCardHeaders, runtime *util.RuntimeOptions) (_result *UniqueQueryUserCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["templateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
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
		Action:      tea.String("UniqueQueryUserCard"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/uniques/cards"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UniqueQueryUserCardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UniqueQueryUserCard(request *UniqueQueryUserCardRequest) (_result *UniqueQueryUserCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UniqueQueryUserCardHeaders{}
	_result = &UniqueQueryUserCardResponse{}
	_body, _err := client.UniqueQueryUserCardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateBranchAttributesInCooperateWithOptions(request *UpdateBranchAttributesInCooperateRequest, headers *UpdateBranchAttributesInCooperateHeaders, runtime *util.RuntimeOptions) (_result *UpdateBranchAttributesInCooperateResponse, _err error) {
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
		Action:      tea.String("UpdateBranchAttributesInCooperate"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/cooperateCorps/branchAttributes"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateBranchAttributesInCooperateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateBranchAttributesInCooperate(request *UpdateBranchAttributesInCooperateRequest) (_result *UpdateBranchAttributesInCooperateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateBranchAttributesInCooperateHeaders{}
	_result = &UpdateBranchAttributesInCooperateResponse{}
	_body, _err := client.UpdateBranchAttributesInCooperateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateBranchVisibleSettingInCooperateWithOptions(request *UpdateBranchVisibleSettingInCooperateRequest, headers *UpdateBranchVisibleSettingInCooperateHeaders, runtime *util.RuntimeOptions) (_result *UpdateBranchVisibleSettingInCooperateResponse, _err error) {
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
		Action:      tea.String("UpdateBranchVisibleSettingInCooperate"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/cooperateCorps/branchVisibleSettings"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateBranchVisibleSettingInCooperateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateBranchVisibleSettingInCooperate(request *UpdateBranchVisibleSettingInCooperateRequest) (_result *UpdateBranchVisibleSettingInCooperateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateBranchVisibleSettingInCooperateHeaders{}
	_result = &UpdateBranchVisibleSettingInCooperateResponse{}
	_body, _err := client.UpdateBranchVisibleSettingInCooperateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateContactHideBySceneSettingWithOptions(settingId *string, request *UpdateContactHideBySceneSettingRequest, headers *UpdateContactHideBySceneSettingHeaders, runtime *util.RuntimeOptions) (_result *UpdateContactHideBySceneSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeDeptIds)) {
		body["excludeDeptIds"] = request.ExcludeDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeTagIds)) {
		body["excludeTagIds"] = request.ExcludeTagIds
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeUserIds)) {
		body["excludeUserIds"] = request.ExcludeUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NodeListSceneConfig)) {
		body["nodeListSceneConfig"] = request.NodeListSceneConfig
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectDeptIds)) {
		body["objectDeptIds"] = request.ObjectDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectTagIds)) {
		body["objectTagIds"] = request.ObjectTagIds
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectUserIds)) {
		body["objectUserIds"] = request.ObjectUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProfileSceneConfig)) {
		body["profileSceneConfig"] = request.ProfileSceneConfig
	}

	if !tea.BoolValue(util.IsUnset(request.SearchSceneConfig)) {
		body["searchSceneConfig"] = request.SearchSceneConfig
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
		Action:      tea.String("UpdateContactHideBySceneSetting"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/organizations/hides/settings/" + tea.StringValue(settingId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateContactHideBySceneSettingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateContactHideBySceneSetting(settingId *string, request *UpdateContactHideBySceneSettingRequest) (_result *UpdateContactHideBySceneSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateContactHideBySceneSettingHeaders{}
	_result = &UpdateContactHideBySceneSettingResponse{}
	_body, _err := client.UpdateContactHideBySceneSettingWithOptions(settingId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateContactHideSettingWithOptions(request *UpdateContactHideSettingRequest, headers *UpdateContactHideSettingHeaders, runtime *util.RuntimeOptions) (_result *UpdateContactHideSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Active)) {
		body["active"] = request.Active
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeDeptIds)) {
		body["excludeDeptIds"] = request.ExcludeDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeStaffIds)) {
		body["excludeStaffIds"] = request.ExcludeStaffIds
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeTagIds)) {
		body["excludeTagIds"] = request.ExcludeTagIds
	}

	if !tea.BoolValue(util.IsUnset(request.HideInSearch)) {
		body["hideInSearch"] = request.HideInSearch
	}

	if !tea.BoolValue(util.IsUnset(request.HideInUserProfile)) {
		body["hideInUserProfile"] = request.HideInUserProfile
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectDeptIds)) {
		body["objectDeptIds"] = request.ObjectDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectStaffIds)) {
		body["objectStaffIds"] = request.ObjectStaffIds
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectTagIds)) {
		body["objectTagIds"] = request.ObjectTagIds
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
		Action:      tea.String("UpdateContactHideSetting"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/contactHideSettings"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateContactHideSettingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateContactHideSetting(request *UpdateContactHideSettingRequest) (_result *UpdateContactHideSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateContactHideSettingHeaders{}
	_result = &UpdateContactHideSettingResponse{}
	_body, _err := client.UpdateContactHideSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateContactRestrictSettingWithOptions(request *UpdateContactRestrictSettingRequest, headers *UpdateContactRestrictSettingHeaders, runtime *util.RuntimeOptions) (_result *UpdateContactRestrictSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Active)) {
		body["active"] = request.Active
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeDeptIds)) {
		body["excludeDeptIds"] = request.ExcludeDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeTagIds)) {
		body["excludeTagIds"] = request.ExcludeTagIds
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeUserIds)) {
		body["excludeUserIds"] = request.ExcludeUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RestrictInSearch)) {
		body["restrictInSearch"] = request.RestrictInSearch
	}

	if !tea.BoolValue(util.IsUnset(request.RestrictInUserProfile)) {
		body["restrictInUserProfile"] = request.RestrictInUserProfile
	}

	if !tea.BoolValue(util.IsUnset(request.SubjectDeptIds)) {
		body["subjectDeptIds"] = request.SubjectDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.SubjectTagIds)) {
		body["subjectTagIds"] = request.SubjectTagIds
	}

	if !tea.BoolValue(util.IsUnset(request.SubjectUserIds)) {
		body["subjectUserIds"] = request.SubjectUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
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
		Action:      tea.String("UpdateContactRestrictSetting"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/restrictions/settings"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateContactRestrictSettingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateContactRestrictSetting(request *UpdateContactRestrictSettingRequest) (_result *UpdateContactRestrictSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateContactRestrictSettingHeaders{}
	_result = &UpdateContactRestrictSettingResponse{}
	_body, _err := client.UpdateContactRestrictSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDeptSettngTailFirstWithOptions(request *UpdateDeptSettngTailFirstRequest, headers *UpdateDeptSettngTailFirstHeaders, runtime *util.RuntimeOptions) (_result *UpdateDeptSettngTailFirstResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		body["enable"] = request.Enable
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
		Action:      tea.String("UpdateDeptSettngTailFirst"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/depts/settings/priorities"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDeptSettngTailFirstResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDeptSettngTailFirst(request *UpdateDeptSettngTailFirstRequest) (_result *UpdateDeptSettngTailFirstResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateDeptSettngTailFirstHeaders{}
	_result = &UpdateDeptSettngTailFirstResponse{}
	_body, _err := client.UpdateDeptSettngTailFirstWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEmpAttrbuteVisibilitySettingWithOptions(request *UpdateEmpAttrbuteVisibilitySettingRequest, headers *UpdateEmpAttrbuteVisibilitySettingHeaders, runtime *util.RuntimeOptions) (_result *UpdateEmpAttrbuteVisibilitySettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Active)) {
		body["active"] = request.Active
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeDeptIds)) {
		body["excludeDeptIds"] = request.ExcludeDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeStaffIds)) {
		body["excludeStaffIds"] = request.ExcludeStaffIds
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeTagIds)) {
		body["excludeTagIds"] = request.ExcludeTagIds
	}

	if !tea.BoolValue(util.IsUnset(request.HideFields)) {
		body["hideFields"] = request.HideFields
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectDeptIds)) {
		body["objectDeptIds"] = request.ObjectDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectStaffIds)) {
		body["objectStaffIds"] = request.ObjectStaffIds
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectTagIds)) {
		body["objectTagIds"] = request.ObjectTagIds
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
		Action:      tea.String("UpdateEmpAttrbuteVisibilitySetting"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/staffAttributes/visibilitySettings"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEmpAttrbuteVisibilitySettingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEmpAttrbuteVisibilitySetting(request *UpdateEmpAttrbuteVisibilitySettingRequest) (_result *UpdateEmpAttrbuteVisibilitySettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateEmpAttrbuteVisibilitySettingHeaders{}
	_result = &UpdateEmpAttrbuteVisibilitySettingResponse{}
	_body, _err := client.UpdateEmpAttrbuteVisibilitySettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEmpAttributeHideBySceneSettingWithOptions(settingId *string, request *UpdateEmpAttributeHideBySceneSettingRequest, headers *UpdateEmpAttributeHideBySceneSettingHeaders, runtime *util.RuntimeOptions) (_result *UpdateEmpAttributeHideBySceneSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChatSubtitleConfig)) {
		body["chatSubtitleConfig"] = request.ChatSubtitleConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeDeptIds)) {
		body["excludeDeptIds"] = request.ExcludeDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeTagIds)) {
		body["excludeTagIds"] = request.ExcludeTagIds
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeUserIds)) {
		body["excludeUserIds"] = request.ExcludeUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.HideFields)) {
		body["hideFields"] = request.HideFields
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectDeptIds)) {
		body["objectDeptIds"] = request.ObjectDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectTagIds)) {
		body["objectTagIds"] = request.ObjectTagIds
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectUserIds)) {
		body["objectUserIds"] = request.ObjectUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProfileSceneConfig)) {
		body["profileSceneConfig"] = request.ProfileSceneConfig
	}

	if !tea.BoolValue(util.IsUnset(request.SearchSceneConfig)) {
		body["searchSceneConfig"] = request.SearchSceneConfig
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
		Action:      tea.String("UpdateEmpAttributeHideBySceneSetting"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/empAttributes/hides/settings/" + tea.StringValue(settingId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEmpAttributeHideBySceneSettingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEmpAttributeHideBySceneSetting(settingId *string, request *UpdateEmpAttributeHideBySceneSettingRequest) (_result *UpdateEmpAttributeHideBySceneSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateEmpAttributeHideBySceneSettingHeaders{}
	_result = &UpdateEmpAttributeHideBySceneSettingResponse{}
	_body, _err := client.UpdateEmpAttributeHideBySceneSettingWithOptions(settingId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateManagementGroupWithOptions(groupId *string, request *UpdateManagementGroupRequest, headers *UpdateManagementGroupHeaders, runtime *util.RuntimeOptions) (_result *UpdateManagementGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		body["resourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["scope"] = request.Scope
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
		Action:      tea.String("UpdateManagementGroup"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/managementGroups/" + tea.StringValue(groupId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateManagementGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateManagementGroup(groupId *string, request *UpdateManagementGroupRequest) (_result *UpdateManagementGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateManagementGroupHeaders{}
	_result = &UpdateManagementGroupResponse{}
	_body, _err := client.UpdateManagementGroupWithOptions(groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSeniorSettingWithOptions(request *UpdateSeniorSettingRequest, headers *UpdateSeniorSettingHeaders, runtime *util.RuntimeOptions) (_result *UpdateSeniorSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Open)) {
		body["open"] = request.Open
	}

	if !tea.BoolValue(util.IsUnset(request.PermitDeptIds)) {
		body["permitDeptIds"] = request.PermitDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.PermitStaffIds)) {
		body["permitStaffIds"] = request.PermitStaffIds
	}

	if !tea.BoolValue(util.IsUnset(request.PermitTagIds)) {
		body["permitTagIds"] = request.PermitTagIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProtectScenes)) {
		body["protectScenes"] = request.ProtectScenes
	}

	if !tea.BoolValue(util.IsUnset(request.SeniorStaffId)) {
		body["seniorStaffId"] = request.SeniorStaffId
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
		Action:      tea.String("UpdateSeniorSetting"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/seniorSettings"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateSeniorSettingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSeniorSetting(request *UpdateSeniorSettingRequest) (_result *UpdateSeniorSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateSeniorSettingHeaders{}
	_result = &UpdateSeniorSettingResponse{}
	_body, _err := client.UpdateSeniorSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUserOwnnessWithOptions(userId *string, request *UpdateUserOwnnessRequest, headers *UpdateUserOwnnessHeaders, runtime *util.RuntimeOptions) (_result *UpdateUserOwnnessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeletedFlag)) {
		body["deletedFlag"] = request.DeletedFlag
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnenssType)) {
		body["ownenssType"] = request.OwnenssType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
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
		Action:      tea.String("UpdateUserOwnness"),
		Version:     tea.String("contact_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/contact/user/" + tea.StringValue(userId) + "/ownness"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserOwnnessResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUserOwnness(userId *string, request *UpdateUserOwnnessRequest) (_result *UpdateUserOwnnessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateUserOwnnessHeaders{}
	_result = &UpdateUserOwnnessResponse{}
	_body, _err := client.UpdateUserOwnnessWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
