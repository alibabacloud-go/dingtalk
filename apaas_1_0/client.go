// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package apaas_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BatchCreateTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchCreateTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateTemplateHeaders) GoString() string {
	return s.String()
}

func (s *BatchCreateTemplateHeaders) SetCommonHeaders(v map[string]*string) *BatchCreateTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchCreateTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *BatchCreateTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchCreateTemplateRequest struct {
	TemplateList []*BatchCreateTemplateRequestTemplateList `json:"templateList,omitempty" xml:"templateList,omitempty" type:"Repeated"`
}

func (s BatchCreateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateTemplateRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateTemplateRequest) SetTemplateList(v []*BatchCreateTemplateRequestTemplateList) *BatchCreateTemplateRequest {
	s.TemplateList = v
	return s
}

type BatchCreateTemplateRequestTemplateList struct {
	// adaptEnv
	AdaptEnv []*string `json:"adaptEnv,omitempty" xml:"adaptEnv,omitempty" type:"Repeated"`
	// appDesc
	AppDesc *string `json:"appDesc,omitempty" xml:"appDesc,omitempty"`
	// appIcon
	AppIcon *string `json:"appIcon,omitempty" xml:"appIcon,omitempty"`
	// caseVideoList
	CaseVideoList []*string `json:"caseVideoList,omitempty" xml:"caseVideoList,omitempty" type:"Repeated"`
	// category
	CategoryCode *string `json:"categoryCode,omitempty" xml:"categoryCode,omitempty"`
	// coverImgList
	CoverImgList []*string `json:"coverImgList,omitempty" xml:"coverImgList,omitempty" type:"Repeated"`
	// expUrl
	ExpUrl *string `json:"expUrl,omitempty" xml:"expUrl,omitempty"`
	// industryLabelList
	IndustryLabelList []*string `json:"industryLabelList,omitempty" xml:"industryLabelList,omitempty" type:"Repeated"`
	// installTimes
	InstallTimes *int64 `json:"installTimes,omitempty" xml:"installTimes,omitempty"`
	// mobilePreviewMediaList
	MobilePreviewMediaList []*string `json:"mobilePreviewMediaList,omitempty" xml:"mobilePreviewMediaList,omitempty" type:"Repeated"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// previewMediaList
	PreviewMediaList []*string `json:"previewMediaList,omitempty" xml:"previewMediaList,omitempty" type:"Repeated"`
	// providerName
	ProviderName *string `json:"providerName,omitempty" xml:"providerName,omitempty"`
	// roleLabelList
	RoleLabelList []*string `json:"roleLabelList,omitempty" xml:"roleLabelList,omitempty" type:"Repeated"`
	// simpleDesc
	SimpleDesc *string `json:"simpleDesc,omitempty" xml:"simpleDesc,omitempty"`
	// templateKey
	TemplateKey *string `json:"templateKey,omitempty" xml:"templateKey,omitempty"`
	// useCasesMediaList
	UseCasesMediaList []*string `json:"useCasesMediaList,omitempty" xml:"useCasesMediaList,omitempty" type:"Repeated"`
}

func (s BatchCreateTemplateRequestTemplateList) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateTemplateRequestTemplateList) GoString() string {
	return s.String()
}

func (s *BatchCreateTemplateRequestTemplateList) SetAdaptEnv(v []*string) *BatchCreateTemplateRequestTemplateList {
	s.AdaptEnv = v
	return s
}

func (s *BatchCreateTemplateRequestTemplateList) SetAppDesc(v string) *BatchCreateTemplateRequestTemplateList {
	s.AppDesc = &v
	return s
}

func (s *BatchCreateTemplateRequestTemplateList) SetAppIcon(v string) *BatchCreateTemplateRequestTemplateList {
	s.AppIcon = &v
	return s
}

func (s *BatchCreateTemplateRequestTemplateList) SetCaseVideoList(v []*string) *BatchCreateTemplateRequestTemplateList {
	s.CaseVideoList = v
	return s
}

func (s *BatchCreateTemplateRequestTemplateList) SetCategoryCode(v string) *BatchCreateTemplateRequestTemplateList {
	s.CategoryCode = &v
	return s
}

func (s *BatchCreateTemplateRequestTemplateList) SetCoverImgList(v []*string) *BatchCreateTemplateRequestTemplateList {
	s.CoverImgList = v
	return s
}

func (s *BatchCreateTemplateRequestTemplateList) SetExpUrl(v string) *BatchCreateTemplateRequestTemplateList {
	s.ExpUrl = &v
	return s
}

func (s *BatchCreateTemplateRequestTemplateList) SetIndustryLabelList(v []*string) *BatchCreateTemplateRequestTemplateList {
	s.IndustryLabelList = v
	return s
}

func (s *BatchCreateTemplateRequestTemplateList) SetInstallTimes(v int64) *BatchCreateTemplateRequestTemplateList {
	s.InstallTimes = &v
	return s
}

func (s *BatchCreateTemplateRequestTemplateList) SetMobilePreviewMediaList(v []*string) *BatchCreateTemplateRequestTemplateList {
	s.MobilePreviewMediaList = v
	return s
}

func (s *BatchCreateTemplateRequestTemplateList) SetName(v string) *BatchCreateTemplateRequestTemplateList {
	s.Name = &v
	return s
}

func (s *BatchCreateTemplateRequestTemplateList) SetPreviewMediaList(v []*string) *BatchCreateTemplateRequestTemplateList {
	s.PreviewMediaList = v
	return s
}

func (s *BatchCreateTemplateRequestTemplateList) SetProviderName(v string) *BatchCreateTemplateRequestTemplateList {
	s.ProviderName = &v
	return s
}

func (s *BatchCreateTemplateRequestTemplateList) SetRoleLabelList(v []*string) *BatchCreateTemplateRequestTemplateList {
	s.RoleLabelList = v
	return s
}

func (s *BatchCreateTemplateRequestTemplateList) SetSimpleDesc(v string) *BatchCreateTemplateRequestTemplateList {
	s.SimpleDesc = &v
	return s
}

func (s *BatchCreateTemplateRequestTemplateList) SetTemplateKey(v string) *BatchCreateTemplateRequestTemplateList {
	s.TemplateKey = &v
	return s
}

func (s *BatchCreateTemplateRequestTemplateList) SetUseCasesMediaList(v []*string) *BatchCreateTemplateRequestTemplateList {
	s.UseCasesMediaList = v
	return s
}

type BatchCreateTemplateResponseBody struct {
	CreateResultList []*BatchCreateTemplateResponseBodyCreateResultList `json:"createResultList,omitempty" xml:"createResultList,omitempty" type:"Repeated"`
}

func (s BatchCreateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateTemplateResponseBody) SetCreateResultList(v []*BatchCreateTemplateResponseBodyCreateResultList) *BatchCreateTemplateResponseBody {
	s.CreateResultList = v
	return s
}

type BatchCreateTemplateResponseBodyCreateResultList struct {
	TemplateKey *string `json:"templateKey,omitempty" xml:"templateKey,omitempty"`
	Value       *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s BatchCreateTemplateResponseBodyCreateResultList) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateTemplateResponseBodyCreateResultList) GoString() string {
	return s.String()
}

func (s *BatchCreateTemplateResponseBodyCreateResultList) SetTemplateKey(v string) *BatchCreateTemplateResponseBodyCreateResultList {
	s.TemplateKey = &v
	return s
}

func (s *BatchCreateTemplateResponseBodyCreateResultList) SetValue(v string) *BatchCreateTemplateResponseBodyCreateResultList {
	s.Value = &v
	return s
}

type BatchCreateTemplateResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchCreateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchCreateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateTemplateResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateTemplateResponse) SetHeaders(v map[string]*string) *BatchCreateTemplateResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateTemplateResponse) SetBody(v *BatchCreateTemplateResponseBody) *BatchCreateTemplateResponse {
	s.Body = v
	return s
}

type BatchQueryByTemplateKeyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchQueryByTemplateKeyHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryByTemplateKeyHeaders) GoString() string {
	return s.String()
}

func (s *BatchQueryByTemplateKeyHeaders) SetCommonHeaders(v map[string]*string) *BatchQueryByTemplateKeyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchQueryByTemplateKeyHeaders) SetXAcsDingtalkAccessToken(v string) *BatchQueryByTemplateKeyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchQueryByTemplateKeyRequest struct {
	TemplateKeys []*string `json:"templateKeys,omitempty" xml:"templateKeys,omitempty" type:"Repeated"`
}

func (s BatchQueryByTemplateKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryByTemplateKeyRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryByTemplateKeyRequest) SetTemplateKeys(v []*string) *BatchQueryByTemplateKeyRequest {
	s.TemplateKeys = v
	return s
}

type BatchQueryByTemplateKeyResponseBody struct {
	TemplateList []*BatchQueryByTemplateKeyResponseBodyTemplateList `json:"templateList,omitempty" xml:"templateList,omitempty" type:"Repeated"`
}

func (s BatchQueryByTemplateKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryByTemplateKeyResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQueryByTemplateKeyResponseBody) SetTemplateList(v []*BatchQueryByTemplateKeyResponseBodyTemplateList) *BatchQueryByTemplateKeyResponseBody {
	s.TemplateList = v
	return s
}

type BatchQueryByTemplateKeyResponseBodyTemplateList struct {
	AdaptEnv               []*string `json:"adaptEnv,omitempty" xml:"adaptEnv,omitempty" type:"Repeated"`
	AppDesc                *string   `json:"appDesc,omitempty" xml:"appDesc,omitempty"`
	AppIcon                *string   `json:"appIcon,omitempty" xml:"appIcon,omitempty"`
	CaseVideoList          []*string `json:"caseVideoList,omitempty" xml:"caseVideoList,omitempty" type:"Repeated"`
	Category               *string   `json:"category,omitempty" xml:"category,omitempty"`
	CoverImgList           []*string `json:"coverImgList,omitempty" xml:"coverImgList,omitempty" type:"Repeated"`
	ExpUrl                 *string   `json:"expUrl,omitempty" xml:"expUrl,omitempty"`
	IndustryLabelList      []*string `json:"industryLabelList,omitempty" xml:"industryLabelList,omitempty" type:"Repeated"`
	InstallTimes           *float32  `json:"installTimes,omitempty" xml:"installTimes,omitempty"`
	MobilePreviewMediaList []*string `json:"mobilePreviewMediaList,omitempty" xml:"mobilePreviewMediaList,omitempty" type:"Repeated"`
	Name                   *string   `json:"name,omitempty" xml:"name,omitempty"`
	PreviewMediaList       []*string `json:"previewMediaList,omitempty" xml:"previewMediaList,omitempty" type:"Repeated"`
	ProviderName           *string   `json:"providerName,omitempty" xml:"providerName,omitempty"`
	RoleLabelList          []*string `json:"roleLabelList,omitempty" xml:"roleLabelList,omitempty" type:"Repeated"`
	SimpleDesc             *string   `json:"simpleDesc,omitempty" xml:"simpleDesc,omitempty"`
	TemplateKey            *string   `json:"templateKey,omitempty" xml:"templateKey,omitempty"`
	UseCasesMediaList      []*string `json:"useCasesMediaList,omitempty" xml:"useCasesMediaList,omitempty" type:"Repeated"`
}

func (s BatchQueryByTemplateKeyResponseBodyTemplateList) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryByTemplateKeyResponseBodyTemplateList) GoString() string {
	return s.String()
}

func (s *BatchQueryByTemplateKeyResponseBodyTemplateList) SetAdaptEnv(v []*string) *BatchQueryByTemplateKeyResponseBodyTemplateList {
	s.AdaptEnv = v
	return s
}

func (s *BatchQueryByTemplateKeyResponseBodyTemplateList) SetAppDesc(v string) *BatchQueryByTemplateKeyResponseBodyTemplateList {
	s.AppDesc = &v
	return s
}

func (s *BatchQueryByTemplateKeyResponseBodyTemplateList) SetAppIcon(v string) *BatchQueryByTemplateKeyResponseBodyTemplateList {
	s.AppIcon = &v
	return s
}

func (s *BatchQueryByTemplateKeyResponseBodyTemplateList) SetCaseVideoList(v []*string) *BatchQueryByTemplateKeyResponseBodyTemplateList {
	s.CaseVideoList = v
	return s
}

func (s *BatchQueryByTemplateKeyResponseBodyTemplateList) SetCategory(v string) *BatchQueryByTemplateKeyResponseBodyTemplateList {
	s.Category = &v
	return s
}

func (s *BatchQueryByTemplateKeyResponseBodyTemplateList) SetCoverImgList(v []*string) *BatchQueryByTemplateKeyResponseBodyTemplateList {
	s.CoverImgList = v
	return s
}

func (s *BatchQueryByTemplateKeyResponseBodyTemplateList) SetExpUrl(v string) *BatchQueryByTemplateKeyResponseBodyTemplateList {
	s.ExpUrl = &v
	return s
}

func (s *BatchQueryByTemplateKeyResponseBodyTemplateList) SetIndustryLabelList(v []*string) *BatchQueryByTemplateKeyResponseBodyTemplateList {
	s.IndustryLabelList = v
	return s
}

func (s *BatchQueryByTemplateKeyResponseBodyTemplateList) SetInstallTimes(v float32) *BatchQueryByTemplateKeyResponseBodyTemplateList {
	s.InstallTimes = &v
	return s
}

func (s *BatchQueryByTemplateKeyResponseBodyTemplateList) SetMobilePreviewMediaList(v []*string) *BatchQueryByTemplateKeyResponseBodyTemplateList {
	s.MobilePreviewMediaList = v
	return s
}

func (s *BatchQueryByTemplateKeyResponseBodyTemplateList) SetName(v string) *BatchQueryByTemplateKeyResponseBodyTemplateList {
	s.Name = &v
	return s
}

func (s *BatchQueryByTemplateKeyResponseBodyTemplateList) SetPreviewMediaList(v []*string) *BatchQueryByTemplateKeyResponseBodyTemplateList {
	s.PreviewMediaList = v
	return s
}

func (s *BatchQueryByTemplateKeyResponseBodyTemplateList) SetProviderName(v string) *BatchQueryByTemplateKeyResponseBodyTemplateList {
	s.ProviderName = &v
	return s
}

func (s *BatchQueryByTemplateKeyResponseBodyTemplateList) SetRoleLabelList(v []*string) *BatchQueryByTemplateKeyResponseBodyTemplateList {
	s.RoleLabelList = v
	return s
}

func (s *BatchQueryByTemplateKeyResponseBodyTemplateList) SetSimpleDesc(v string) *BatchQueryByTemplateKeyResponseBodyTemplateList {
	s.SimpleDesc = &v
	return s
}

func (s *BatchQueryByTemplateKeyResponseBodyTemplateList) SetTemplateKey(v string) *BatchQueryByTemplateKeyResponseBodyTemplateList {
	s.TemplateKey = &v
	return s
}

func (s *BatchQueryByTemplateKeyResponseBodyTemplateList) SetUseCasesMediaList(v []*string) *BatchQueryByTemplateKeyResponseBodyTemplateList {
	s.UseCasesMediaList = v
	return s
}

type BatchQueryByTemplateKeyResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchQueryByTemplateKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchQueryByTemplateKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryByTemplateKeyResponse) GoString() string {
	return s.String()
}

func (s *BatchQueryByTemplateKeyResponse) SetHeaders(v map[string]*string) *BatchQueryByTemplateKeyResponse {
	s.Headers = v
	return s
}

func (s *BatchQueryByTemplateKeyResponse) SetBody(v *BatchQueryByTemplateKeyResponseBody) *BatchQueryByTemplateKeyResponse {
	s.Body = v
	return s
}

type BatchUpdateTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchUpdateTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateTemplateHeaders) GoString() string {
	return s.String()
}

func (s *BatchUpdateTemplateHeaders) SetCommonHeaders(v map[string]*string) *BatchUpdateTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchUpdateTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *BatchUpdateTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchUpdateTemplateRequest struct {
	TemplateList []*BatchUpdateTemplateRequestTemplateList `json:"templateList,omitempty" xml:"templateList,omitempty" type:"Repeated"`
}

func (s BatchUpdateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateTemplateRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateTemplateRequest) SetTemplateList(v []*BatchUpdateTemplateRequestTemplateList) *BatchUpdateTemplateRequest {
	s.TemplateList = v
	return s
}

type BatchUpdateTemplateRequestTemplateList struct {
	// adaptEnv
	AdaptEnv []*string `json:"adaptEnv,omitempty" xml:"adaptEnv,omitempty" type:"Repeated"`
	// appDesc
	AppDesc *string `json:"appDesc,omitempty" xml:"appDesc,omitempty"`
	// appIcon
	AppIcon *string `json:"appIcon,omitempty" xml:"appIcon,omitempty"`
	// caseVideoList
	CaseVideoList []*string `json:"caseVideoList,omitempty" xml:"caseVideoList,omitempty" type:"Repeated"`
	// category
	CategoryCode *string `json:"categoryCode,omitempty" xml:"categoryCode,omitempty"`
	// coverImgList
	CoverImgList []*string `json:"coverImgList,omitempty" xml:"coverImgList,omitempty" type:"Repeated"`
	// expUrl
	ExpUrl *string `json:"expUrl,omitempty" xml:"expUrl,omitempty"`
	// industryLabelList
	IndustryLabelList []*string `json:"industryLabelList,omitempty" xml:"industryLabelList,omitempty" type:"Repeated"`
	// mobilePreviewMediaList
	MobilePreviewMediaList []*string `json:"mobilePreviewMediaList,omitempty" xml:"mobilePreviewMediaList,omitempty" type:"Repeated"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// previewMediaList
	PreviewMediaList []*string `json:"previewMediaList,omitempty" xml:"previewMediaList,omitempty" type:"Repeated"`
	// providerName
	ProviderName *string `json:"providerName,omitempty" xml:"providerName,omitempty"`
	// roleLabelList
	RoleLabelList []*string `json:"roleLabelList,omitempty" xml:"roleLabelList,omitempty" type:"Repeated"`
	// simpleDesc
	SimpleDesc *string `json:"simpleDesc,omitempty" xml:"simpleDesc,omitempty"`
	// templateKey
	TemplateKey *string `json:"templateKey,omitempty" xml:"templateKey,omitempty"`
	// useCasesMediaList
	UseCasesMediaList []*string `json:"useCasesMediaList,omitempty" xml:"useCasesMediaList,omitempty" type:"Repeated"`
}

func (s BatchUpdateTemplateRequestTemplateList) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateTemplateRequestTemplateList) GoString() string {
	return s.String()
}

func (s *BatchUpdateTemplateRequestTemplateList) SetAdaptEnv(v []*string) *BatchUpdateTemplateRequestTemplateList {
	s.AdaptEnv = v
	return s
}

func (s *BatchUpdateTemplateRequestTemplateList) SetAppDesc(v string) *BatchUpdateTemplateRequestTemplateList {
	s.AppDesc = &v
	return s
}

func (s *BatchUpdateTemplateRequestTemplateList) SetAppIcon(v string) *BatchUpdateTemplateRequestTemplateList {
	s.AppIcon = &v
	return s
}

func (s *BatchUpdateTemplateRequestTemplateList) SetCaseVideoList(v []*string) *BatchUpdateTemplateRequestTemplateList {
	s.CaseVideoList = v
	return s
}

func (s *BatchUpdateTemplateRequestTemplateList) SetCategoryCode(v string) *BatchUpdateTemplateRequestTemplateList {
	s.CategoryCode = &v
	return s
}

func (s *BatchUpdateTemplateRequestTemplateList) SetCoverImgList(v []*string) *BatchUpdateTemplateRequestTemplateList {
	s.CoverImgList = v
	return s
}

func (s *BatchUpdateTemplateRequestTemplateList) SetExpUrl(v string) *BatchUpdateTemplateRequestTemplateList {
	s.ExpUrl = &v
	return s
}

func (s *BatchUpdateTemplateRequestTemplateList) SetIndustryLabelList(v []*string) *BatchUpdateTemplateRequestTemplateList {
	s.IndustryLabelList = v
	return s
}

func (s *BatchUpdateTemplateRequestTemplateList) SetMobilePreviewMediaList(v []*string) *BatchUpdateTemplateRequestTemplateList {
	s.MobilePreviewMediaList = v
	return s
}

func (s *BatchUpdateTemplateRequestTemplateList) SetName(v string) *BatchUpdateTemplateRequestTemplateList {
	s.Name = &v
	return s
}

func (s *BatchUpdateTemplateRequestTemplateList) SetPreviewMediaList(v []*string) *BatchUpdateTemplateRequestTemplateList {
	s.PreviewMediaList = v
	return s
}

func (s *BatchUpdateTemplateRequestTemplateList) SetProviderName(v string) *BatchUpdateTemplateRequestTemplateList {
	s.ProviderName = &v
	return s
}

func (s *BatchUpdateTemplateRequestTemplateList) SetRoleLabelList(v []*string) *BatchUpdateTemplateRequestTemplateList {
	s.RoleLabelList = v
	return s
}

func (s *BatchUpdateTemplateRequestTemplateList) SetSimpleDesc(v string) *BatchUpdateTemplateRequestTemplateList {
	s.SimpleDesc = &v
	return s
}

func (s *BatchUpdateTemplateRequestTemplateList) SetTemplateKey(v string) *BatchUpdateTemplateRequestTemplateList {
	s.TemplateKey = &v
	return s
}

func (s *BatchUpdateTemplateRequestTemplateList) SetUseCasesMediaList(v []*string) *BatchUpdateTemplateRequestTemplateList {
	s.UseCasesMediaList = v
	return s
}

type BatchUpdateTemplateResponseBody struct {
	UpdateResultList []*BatchUpdateTemplateResponseBodyUpdateResultList `json:"updateResultList,omitempty" xml:"updateResultList,omitempty" type:"Repeated"`
}

func (s BatchUpdateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateTemplateResponseBody) SetUpdateResultList(v []*BatchUpdateTemplateResponseBodyUpdateResultList) *BatchUpdateTemplateResponseBody {
	s.UpdateResultList = v
	return s
}

type BatchUpdateTemplateResponseBodyUpdateResultList struct {
	TemplateKey *string `json:"templateKey,omitempty" xml:"templateKey,omitempty"`
	Value       *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s BatchUpdateTemplateResponseBodyUpdateResultList) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateTemplateResponseBodyUpdateResultList) GoString() string {
	return s.String()
}

func (s *BatchUpdateTemplateResponseBodyUpdateResultList) SetTemplateKey(v string) *BatchUpdateTemplateResponseBodyUpdateResultList {
	s.TemplateKey = &v
	return s
}

func (s *BatchUpdateTemplateResponseBodyUpdateResultList) SetValue(v string) *BatchUpdateTemplateResponseBodyUpdateResultList {
	s.Value = &v
	return s
}

type BatchUpdateTemplateResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchUpdateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchUpdateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateTemplateResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateTemplateResponse) SetHeaders(v map[string]*string) *BatchUpdateTemplateResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateTemplateResponse) SetBody(v *BatchUpdateTemplateResponseBody) *BatchUpdateTemplateResponse {
	s.Body = v
	return s
}

type QueryIndustryTagListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryIndustryTagListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryIndustryTagListHeaders) GoString() string {
	return s.String()
}

func (s *QueryIndustryTagListHeaders) SetCommonHeaders(v map[string]*string) *QueryIndustryTagListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryIndustryTagListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryIndustryTagListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryIndustryTagListResponseBody struct {
	IndustryList []*string `json:"industryList,omitempty" xml:"industryList,omitempty" type:"Repeated"`
}

func (s QueryIndustryTagListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryIndustryTagListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryIndustryTagListResponseBody) SetIndustryList(v []*string) *QueryIndustryTagListResponseBody {
	s.IndustryList = v
	return s
}

type QueryIndustryTagListResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryIndustryTagListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryIndustryTagListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryIndustryTagListResponse) GoString() string {
	return s.String()
}

func (s *QueryIndustryTagListResponse) SetHeaders(v map[string]*string) *QueryIndustryTagListResponse {
	s.Headers = v
	return s
}

func (s *QueryIndustryTagListResponse) SetBody(v *QueryIndustryTagListResponseBody) *QueryIndustryTagListResponse {
	s.Body = v
	return s
}

type QueryRoleTagListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryRoleTagListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryRoleTagListHeaders) GoString() string {
	return s.String()
}

func (s *QueryRoleTagListHeaders) SetCommonHeaders(v map[string]*string) *QueryRoleTagListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryRoleTagListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryRoleTagListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryRoleTagListResponseBody struct {
	RoleList []*string `json:"roleList,omitempty" xml:"roleList,omitempty" type:"Repeated"`
}

func (s QueryRoleTagListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRoleTagListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRoleTagListResponseBody) SetRoleList(v []*string) *QueryRoleTagListResponseBody {
	s.RoleList = v
	return s
}

type QueryRoleTagListResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRoleTagListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRoleTagListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRoleTagListResponse) GoString() string {
	return s.String()
}

func (s *QueryRoleTagListResponse) SetHeaders(v map[string]*string) *QueryRoleTagListResponse {
	s.Headers = v
	return s
}

func (s *QueryRoleTagListResponse) SetBody(v *QueryRoleTagListResponseBody) *QueryRoleTagListResponse {
	s.Body = v
	return s
}

type QueryTemplateCategorysHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryTemplateCategorysHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateCategorysHeaders) GoString() string {
	return s.String()
}

func (s *QueryTemplateCategorysHeaders) SetCommonHeaders(v map[string]*string) *QueryTemplateCategorysHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryTemplateCategorysHeaders) SetXAcsDingtalkAccessToken(v string) *QueryTemplateCategorysHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryTemplateCategorysResponseBody struct {
	CategoryList []*QueryTemplateCategorysResponseBodyCategoryList `json:"categoryList,omitempty" xml:"categoryList,omitempty" type:"Repeated"`
	// 总数
	Total *string `json:"total,omitempty" xml:"total,omitempty"`
}

func (s QueryTemplateCategorysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateCategorysResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTemplateCategorysResponseBody) SetCategoryList(v []*QueryTemplateCategorysResponseBodyCategoryList) *QueryTemplateCategorysResponseBody {
	s.CategoryList = v
	return s
}

func (s *QueryTemplateCategorysResponseBody) SetTotal(v string) *QueryTemplateCategorysResponseBody {
	s.Total = &v
	return s
}

type QueryTemplateCategorysResponseBodyCategoryList struct {
	// 分类编码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 分类名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryTemplateCategorysResponseBodyCategoryList) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateCategorysResponseBodyCategoryList) GoString() string {
	return s.String()
}

func (s *QueryTemplateCategorysResponseBodyCategoryList) SetCode(v string) *QueryTemplateCategorysResponseBodyCategoryList {
	s.Code = &v
	return s
}

func (s *QueryTemplateCategorysResponseBodyCategoryList) SetName(v string) *QueryTemplateCategorysResponseBodyCategoryList {
	s.Name = &v
	return s
}

type QueryTemplateCategorysResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTemplateCategorysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTemplateCategorysResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateCategorysResponse) GoString() string {
	return s.String()
}

func (s *QueryTemplateCategorysResponse) SetHeaders(v map[string]*string) *QueryTemplateCategorysResponse {
	s.Headers = v
	return s
}

func (s *QueryTemplateCategorysResponse) SetBody(v *QueryTemplateCategorysResponseBody) *QueryTemplateCategorysResponse {
	s.Body = v
	return s
}

type RecallAuditTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RecallAuditTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s RecallAuditTemplateHeaders) GoString() string {
	return s.String()
}

func (s *RecallAuditTemplateHeaders) SetCommonHeaders(v map[string]*string) *RecallAuditTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RecallAuditTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *RecallAuditTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RecallAuditTemplateRequest struct {
	TemplateKeys []*string `json:"templateKeys,omitempty" xml:"templateKeys,omitempty" type:"Repeated"`
}

func (s RecallAuditTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s RecallAuditTemplateRequest) GoString() string {
	return s.String()
}

func (s *RecallAuditTemplateRequest) SetTemplateKeys(v []*string) *RecallAuditTemplateRequest {
	s.TemplateKeys = v
	return s
}

type RecallAuditTemplateResponseBody struct {
	RecallResult []*RecallAuditTemplateResponseBodyRecallResult `json:"recallResult,omitempty" xml:"recallResult,omitempty" type:"Repeated"`
}

func (s RecallAuditTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecallAuditTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *RecallAuditTemplateResponseBody) SetRecallResult(v []*RecallAuditTemplateResponseBodyRecallResult) *RecallAuditTemplateResponseBody {
	s.RecallResult = v
	return s
}

type RecallAuditTemplateResponseBodyRecallResult struct {
	TemplateKey *string `json:"templateKey,omitempty" xml:"templateKey,omitempty"`
	Value       *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s RecallAuditTemplateResponseBodyRecallResult) String() string {
	return tea.Prettify(s)
}

func (s RecallAuditTemplateResponseBodyRecallResult) GoString() string {
	return s.String()
}

func (s *RecallAuditTemplateResponseBodyRecallResult) SetTemplateKey(v string) *RecallAuditTemplateResponseBodyRecallResult {
	s.TemplateKey = &v
	return s
}

func (s *RecallAuditTemplateResponseBodyRecallResult) SetValue(v string) *RecallAuditTemplateResponseBodyRecallResult {
	s.Value = &v
	return s
}

type RecallAuditTemplateResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecallAuditTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecallAuditTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s RecallAuditTemplateResponse) GoString() string {
	return s.String()
}

func (s *RecallAuditTemplateResponse) SetHeaders(v map[string]*string) *RecallAuditTemplateResponse {
	s.Headers = v
	return s
}

func (s *RecallAuditTemplateResponse) SetBody(v *RecallAuditTemplateResponseBody) *RecallAuditTemplateResponse {
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

func (client *Client) BatchCreateTemplate(request *BatchCreateTemplateRequest) (_result *BatchCreateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchCreateTemplateHeaders{}
	_result = &BatchCreateTemplateResponse{}
	_body, _err := client.BatchCreateTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchCreateTemplateWithOptions(request *BatchCreateTemplateRequest, headers *BatchCreateTemplateHeaders, runtime *util.RuntimeOptions) (_result *BatchCreateTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateList)) {
		body["templateList"] = request.TemplateList
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
	_result = &BatchCreateTemplateResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchCreateTemplate"), tea.String("apaas_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/apaas/templates"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchQueryByTemplateKey(request *BatchQueryByTemplateKeyRequest) (_result *BatchQueryByTemplateKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchQueryByTemplateKeyHeaders{}
	_result = &BatchQueryByTemplateKeyResponse{}
	_body, _err := client.BatchQueryByTemplateKeyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchQueryByTemplateKeyWithOptions(request *BatchQueryByTemplateKeyRequest, headers *BatchQueryByTemplateKeyHeaders, runtime *util.RuntimeOptions) (_result *BatchQueryByTemplateKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateKeys)) {
		body["templateKeys"] = request.TemplateKeys
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
	_result = &BatchQueryByTemplateKeyResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchQueryByTemplateKey"), tea.String("apaas_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/apaas/templates/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchUpdateTemplate(request *BatchUpdateTemplateRequest) (_result *BatchUpdateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchUpdateTemplateHeaders{}
	_result = &BatchUpdateTemplateResponse{}
	_body, _err := client.BatchUpdateTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchUpdateTemplateWithOptions(request *BatchUpdateTemplateRequest, headers *BatchUpdateTemplateHeaders, runtime *util.RuntimeOptions) (_result *BatchUpdateTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateList)) {
		body["templateList"] = request.TemplateList
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
	_result = &BatchUpdateTemplateResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchUpdateTemplate"), tea.String("apaas_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/apaas/templates"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryIndustryTagList() (_result *QueryIndustryTagListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryIndustryTagListHeaders{}
	_result = &QueryIndustryTagListResponse{}
	_body, _err := client.QueryIndustryTagListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryIndustryTagListWithOptions(headers *QueryIndustryTagListHeaders, runtime *util.RuntimeOptions) (_result *QueryIndustryTagListResponse, _err error) {
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
	_result = &QueryIndustryTagListResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryIndustryTagList"), tea.String("apaas_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/apaas/templates/industries"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRoleTagList() (_result *QueryRoleTagListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryRoleTagListHeaders{}
	_result = &QueryRoleTagListResponse{}
	_body, _err := client.QueryRoleTagListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRoleTagListWithOptions(headers *QueryRoleTagListHeaders, runtime *util.RuntimeOptions) (_result *QueryRoleTagListResponse, _err error) {
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
	_result = &QueryRoleTagListResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryRoleTagList"), tea.String("apaas_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/apaas/templates/roles"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTemplateCategorys() (_result *QueryTemplateCategorysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryTemplateCategorysHeaders{}
	_result = &QueryTemplateCategorysResponse{}
	_body, _err := client.QueryTemplateCategorysWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTemplateCategorysWithOptions(headers *QueryTemplateCategorysHeaders, runtime *util.RuntimeOptions) (_result *QueryTemplateCategorysResponse, _err error) {
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
	_result = &QueryTemplateCategorysResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryTemplateCategorys"), tea.String("apaas_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/apaas/templates/categories"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecallAuditTemplate(request *RecallAuditTemplateRequest) (_result *RecallAuditTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RecallAuditTemplateHeaders{}
	_result = &RecallAuditTemplateResponse{}
	_body, _err := client.RecallAuditTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecallAuditTemplateWithOptions(request *RecallAuditTemplateRequest, headers *RecallAuditTemplateHeaders, runtime *util.RuntimeOptions) (_result *RecallAuditTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateKeys)) {
		body["templateKeys"] = request.TemplateKeys
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
	_result = &RecallAuditTemplateResponse{}
	_body, _err := client.DoROARequest(tea.String("RecallAuditTemplate"), tea.String("apaas_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/apaas/templates/audits/recall"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
