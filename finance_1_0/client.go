// This file is auto-generated, don't edit it. Thanks.
package finance_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ApplyBatchPayHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ApplyBatchPayHeaders) String() string {
	return tea.Prettify(s)
}

func (s ApplyBatchPayHeaders) GoString() string {
	return s.String()
}

func (s *ApplyBatchPayHeaders) SetCommonHeaders(v map[string]*string) *ApplyBatchPayHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ApplyBatchPayHeaders) SetXAcsDingtalkAccessToken(v string) *ApplyBatchPayHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ApplyBatchPayRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2021070712440326300185114
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20210909153300000002734753314700
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	// example:
	//
	// map
	PassBackParams map[string]interface{} `json:"passBackParams,omitempty" xml:"passBackParams,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PC
	PayTerminal *string `json:"payTerminal,omitempty" xml:"payTerminal,omitempty"`
	// example:
	//
	// http://xx
	ReturnUrl *string `json:"returnUrl,omitempty" xml:"returnUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8754214873
	StaffId *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10.00
	TransAmount *string `json:"transAmount,omitempty" xml:"transAmount,omitempty"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	TransExpireTime *string `json:"transExpireTime,omitempty" xml:"transExpireTime,omitempty"`
}

func (s ApplyBatchPayRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyBatchPayRequest) GoString() string {
	return s.String()
}

func (s *ApplyBatchPayRequest) SetAccountId(v string) *ApplyBatchPayRequest {
	s.AccountId = &v
	return s
}

func (s *ApplyBatchPayRequest) SetOrderNo(v string) *ApplyBatchPayRequest {
	s.OrderNo = &v
	return s
}

func (s *ApplyBatchPayRequest) SetPassBackParams(v map[string]interface{}) *ApplyBatchPayRequest {
	s.PassBackParams = v
	return s
}

func (s *ApplyBatchPayRequest) SetPayTerminal(v string) *ApplyBatchPayRequest {
	s.PayTerminal = &v
	return s
}

func (s *ApplyBatchPayRequest) SetReturnUrl(v string) *ApplyBatchPayRequest {
	s.ReturnUrl = &v
	return s
}

func (s *ApplyBatchPayRequest) SetStaffId(v string) *ApplyBatchPayRequest {
	s.StaffId = &v
	return s
}

func (s *ApplyBatchPayRequest) SetTransAmount(v string) *ApplyBatchPayRequest {
	s.TransAmount = &v
	return s
}

func (s *ApplyBatchPayRequest) SetTransExpireTime(v string) *ApplyBatchPayRequest {
	s.TransExpireTime = &v
	return s
}

type ApplyBatchPayResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 20210909153300000002734753314700
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// payData
	PayData *string `json:"payData,omitempty" xml:"payData,omitempty"`
}

func (s ApplyBatchPayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyBatchPayResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyBatchPayResponseBody) SetOrderNo(v string) *ApplyBatchPayResponseBody {
	s.OrderNo = &v
	return s
}

func (s *ApplyBatchPayResponseBody) SetPayData(v string) *ApplyBatchPayResponseBody {
	s.PayData = &v
	return s
}

type ApplyBatchPayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyBatchPayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyBatchPayResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyBatchPayResponse) GoString() string {
	return s.String()
}

func (s *ApplyBatchPayResponse) SetHeaders(v map[string]*string) *ApplyBatchPayResponse {
	s.Headers = v
	return s
}

func (s *ApplyBatchPayResponse) SetStatusCode(v int32) *ApplyBatchPayResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyBatchPayResponse) SetBody(v *ApplyBatchPayResponseBody) *ApplyBatchPayResponse {
	s.Body = v
	return s
}

type CloseLoanEntranceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CloseLoanEntranceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CloseLoanEntranceHeaders) GoString() string {
	return s.String()
}

func (s *CloseLoanEntranceHeaders) SetCommonHeaders(v map[string]*string) *CloseLoanEntranceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CloseLoanEntranceHeaders) SetXAcsDingtalkAccessToken(v string) *CloseLoanEntranceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CloseLoanEntranceRequest struct {
	// This parameter is required.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 1a23qdfa
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CloseLoanEntranceRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseLoanEntranceRequest) GoString() string {
	return s.String()
}

func (s *CloseLoanEntranceRequest) SetRequestId(v string) *CloseLoanEntranceRequest {
	s.RequestId = &v
	return s
}

type CloseLoanEntranceResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 1a23qdfa
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// Y
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CloseLoanEntranceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseLoanEntranceResponseBody) GoString() string {
	return s.String()
}

func (s *CloseLoanEntranceResponseBody) SetRequestId(v string) *CloseLoanEntranceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseLoanEntranceResponseBody) SetResult(v string) *CloseLoanEntranceResponseBody {
	s.Result = &v
	return s
}

type CloseLoanEntranceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseLoanEntranceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseLoanEntranceResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseLoanEntranceResponse) GoString() string {
	return s.String()
}

func (s *CloseLoanEntranceResponse) SetHeaders(v map[string]*string) *CloseLoanEntranceResponse {
	s.Headers = v
	return s
}

func (s *CloseLoanEntranceResponse) SetStatusCode(v int32) *CloseLoanEntranceResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseLoanEntranceResponse) SetBody(v *CloseLoanEntranceResponseBody) *CloseLoanEntranceResponse {
	s.Body = v
	return s
}

type ConsultCreateSubInstitutionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ConsultCreateSubInstitutionHeaders) String() string {
	return tea.Prettify(s)
}

func (s ConsultCreateSubInstitutionHeaders) GoString() string {
	return s.String()
}

func (s *ConsultCreateSubInstitutionHeaders) SetCommonHeaders(v map[string]*string) *ConsultCreateSubInstitutionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ConsultCreateSubInstitutionHeaders) SetXAcsDingtalkAccessToken(v string) *ConsultCreateSubInstitutionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ConsultCreateSubInstitutionRequest struct {
	// example:
	//
	// asdf@163.com
	BindingAlipayLogonId *string `json:"bindingAlipayLogonId,omitempty" xml:"bindingAlipayLogonId,omitempty"`
	// This parameter is required.
	ContactInfo *ConsultCreateSubInstitutionRequestContactInfo `json:"contactInfo,omitempty" xml:"contactInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 202111090001
	InstId *string `json:"instId,omitempty" xml:"instId,omitempty"`
	// This parameter is required.
	LegalPersonCertInfo *ConsultCreateSubInstitutionRequestLegalPersonCertInfo `json:"legalPersonCertInfo,omitempty" xml:"legalPersonCertInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 2021000001
	OutTradeNo *string `json:"outTradeNo,omitempty" xml:"outTradeNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALIPAY
	PayChannel         *string                                                 `json:"payChannel,omitempty" xml:"payChannel,omitempty"`
	QualificationInfos []*ConsultCreateSubInstitutionRequestQualificationInfos `json:"qualificationInfos,omitempty" xml:"qualificationInfos,omitempty" type:"Repeated"`
	// This parameter is required.
	Services []*string `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
	// This parameter is required.
	SettleInfo *ConsultCreateSubInstitutionRequestSettleInfo `json:"settleInfo,omitempty" xml:"settleInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// S001
	Solution           *string                                               `json:"solution,omitempty" xml:"solution,omitempty"`
	SubInstAddressInfo *ConsultCreateSubInstitutionRequestSubInstAddressInfo `json:"subInstAddressInfo,omitempty" xml:"subInstAddressInfo,omitempty" type:"Struct"`
	SubInstAuthInfo    *ConsultCreateSubInstitutionRequestSubInstAuthInfo    `json:"subInstAuthInfo,omitempty" xml:"subInstAuthInfo,omitempty" type:"Struct"`
	// This parameter is required.
	SubInstBasicInfo *ConsultCreateSubInstitutionRequestSubInstBasicInfo `json:"subInstBasicInfo,omitempty" xml:"subInstBasicInfo,omitempty" type:"Struct"`
	// This parameter is required.
	SubInstCertifyInfo *ConsultCreateSubInstitutionRequestSubInstCertifyInfo `json:"subInstCertifyInfo,omitempty" xml:"subInstCertifyInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 1001
	SubInstId          *string                                               `json:"subInstId,omitempty" xml:"subInstId,omitempty"`
	SubInstInvoiceInfo *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo `json:"subInstInvoiceInfo,omitempty" xml:"subInstInvoiceInfo,omitempty" type:"Struct"`
	SubInstShopInfo    *ConsultCreateSubInstitutionRequestSubInstShopInfo    `json:"subInstShopInfo,omitempty" xml:"subInstShopInfo,omitempty" type:"Struct"`
}

func (s ConsultCreateSubInstitutionRequest) String() string {
	return tea.Prettify(s)
}

func (s ConsultCreateSubInstitutionRequest) GoString() string {
	return s.String()
}

func (s *ConsultCreateSubInstitutionRequest) SetBindingAlipayLogonId(v string) *ConsultCreateSubInstitutionRequest {
	s.BindingAlipayLogonId = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequest) SetContactInfo(v *ConsultCreateSubInstitutionRequestContactInfo) *ConsultCreateSubInstitutionRequest {
	s.ContactInfo = v
	return s
}

func (s *ConsultCreateSubInstitutionRequest) SetInstId(v string) *ConsultCreateSubInstitutionRequest {
	s.InstId = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequest) SetLegalPersonCertInfo(v *ConsultCreateSubInstitutionRequestLegalPersonCertInfo) *ConsultCreateSubInstitutionRequest {
	s.LegalPersonCertInfo = v
	return s
}

func (s *ConsultCreateSubInstitutionRequest) SetOutTradeNo(v string) *ConsultCreateSubInstitutionRequest {
	s.OutTradeNo = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequest) SetPayChannel(v string) *ConsultCreateSubInstitutionRequest {
	s.PayChannel = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequest) SetQualificationInfos(v []*ConsultCreateSubInstitutionRequestQualificationInfos) *ConsultCreateSubInstitutionRequest {
	s.QualificationInfos = v
	return s
}

func (s *ConsultCreateSubInstitutionRequest) SetServices(v []*string) *ConsultCreateSubInstitutionRequest {
	s.Services = v
	return s
}

func (s *ConsultCreateSubInstitutionRequest) SetSettleInfo(v *ConsultCreateSubInstitutionRequestSettleInfo) *ConsultCreateSubInstitutionRequest {
	s.SettleInfo = v
	return s
}

func (s *ConsultCreateSubInstitutionRequest) SetSolution(v string) *ConsultCreateSubInstitutionRequest {
	s.Solution = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequest) SetSubInstAddressInfo(v *ConsultCreateSubInstitutionRequestSubInstAddressInfo) *ConsultCreateSubInstitutionRequest {
	s.SubInstAddressInfo = v
	return s
}

func (s *ConsultCreateSubInstitutionRequest) SetSubInstAuthInfo(v *ConsultCreateSubInstitutionRequestSubInstAuthInfo) *ConsultCreateSubInstitutionRequest {
	s.SubInstAuthInfo = v
	return s
}

func (s *ConsultCreateSubInstitutionRequest) SetSubInstBasicInfo(v *ConsultCreateSubInstitutionRequestSubInstBasicInfo) *ConsultCreateSubInstitutionRequest {
	s.SubInstBasicInfo = v
	return s
}

func (s *ConsultCreateSubInstitutionRequest) SetSubInstCertifyInfo(v *ConsultCreateSubInstitutionRequestSubInstCertifyInfo) *ConsultCreateSubInstitutionRequest {
	s.SubInstCertifyInfo = v
	return s
}

func (s *ConsultCreateSubInstitutionRequest) SetSubInstId(v string) *ConsultCreateSubInstitutionRequest {
	s.SubInstId = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequest) SetSubInstInvoiceInfo(v *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo) *ConsultCreateSubInstitutionRequest {
	s.SubInstInvoiceInfo = v
	return s
}

func (s *ConsultCreateSubInstitutionRequest) SetSubInstShopInfo(v *ConsultCreateSubInstitutionRequestSubInstShopInfo) *ConsultCreateSubInstitutionRequest {
	s.SubInstShopInfo = v
	return s
}

type ConsultCreateSubInstitutionRequestContactInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 李某某
	ContactName *string `json:"contactName,omitempty" xml:"contactName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13900000000
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
}

func (s ConsultCreateSubInstitutionRequestContactInfo) String() string {
	return tea.Prettify(s)
}

func (s ConsultCreateSubInstitutionRequestContactInfo) GoString() string {
	return s.String()
}

func (s *ConsultCreateSubInstitutionRequestContactInfo) SetContactName(v string) *ConsultCreateSubInstitutionRequestContactInfo {
	s.ContactName = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestContactInfo) SetMobile(v string) *ConsultCreateSubInstitutionRequestContactInfo {
	s.Mobile = &v
	return s
}

type ConsultCreateSubInstitutionRequestLegalPersonCertInfo struct {
	// example:
	//
	// ossUrl
	CertBackImage *string `json:"certBackImage,omitempty" xml:"certBackImage,omitempty"`
	// example:
	//
	// ossUrl
	CertFrontImage *string `json:"certFrontImage,omitempty" xml:"certFrontImage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 李某某
	CertName *string `json:"certName,omitempty" xml:"certName,omitempty"`
	// example:
	//
	// 100
	CertType *string `json:"certType,omitempty" xml:"certType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 330104200010109999
	IdCardNo *string `json:"idCardNo,omitempty" xml:"idCardNo,omitempty"`
}

func (s ConsultCreateSubInstitutionRequestLegalPersonCertInfo) String() string {
	return tea.Prettify(s)
}

func (s ConsultCreateSubInstitutionRequestLegalPersonCertInfo) GoString() string {
	return s.String()
}

func (s *ConsultCreateSubInstitutionRequestLegalPersonCertInfo) SetCertBackImage(v string) *ConsultCreateSubInstitutionRequestLegalPersonCertInfo {
	s.CertBackImage = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestLegalPersonCertInfo) SetCertFrontImage(v string) *ConsultCreateSubInstitutionRequestLegalPersonCertInfo {
	s.CertFrontImage = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestLegalPersonCertInfo) SetCertName(v string) *ConsultCreateSubInstitutionRequestLegalPersonCertInfo {
	s.CertName = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestLegalPersonCertInfo) SetCertType(v string) *ConsultCreateSubInstitutionRequestLegalPersonCertInfo {
	s.CertType = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestLegalPersonCertInfo) SetIdCardNo(v string) *ConsultCreateSubInstitutionRequestLegalPersonCertInfo {
	s.IdCardNo = &v
	return s
}

type ConsultCreateSubInstitutionRequestQualificationInfos struct {
	// This parameter is required.
	//
	// example:
	//
	// ossUrl
	QualificationImage *string `json:"qualificationImage,omitempty" xml:"qualificationImage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 301
	QualificationType *string `json:"qualificationType,omitempty" xml:"qualificationType,omitempty"`
}

func (s ConsultCreateSubInstitutionRequestQualificationInfos) String() string {
	return tea.Prettify(s)
}

func (s ConsultCreateSubInstitutionRequestQualificationInfos) GoString() string {
	return s.String()
}

func (s *ConsultCreateSubInstitutionRequestQualificationInfos) SetQualificationImage(v string) *ConsultCreateSubInstitutionRequestQualificationInfos {
	s.QualificationImage = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestQualificationInfos) SetQualificationType(v string) *ConsultCreateSubInstitutionRequestQualificationInfos {
	s.QualificationType = &v
	return s
}

type ConsultCreateSubInstitutionRequestSettleInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 622202120200000000
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// 李某某
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	// example:
	//
	// DEBIT_CARD
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	// example:
	//
	// 城东支行
	BankBranchName *string `json:"bankBranchName,omitempty" xml:"bankBranchName,omitempty"`
	// example:
	//
	// 杭州市
	BankCity *string `json:"bankCity,omitempty" xml:"bankCity,omitempty"`
	// example:
	//
	// 313791000023
	BankCode *string `json:"bankCode,omitempty" xml:"bankCode,omitempty"`
	// example:
	//
	// 工商银行
	BankName *string `json:"bankName,omitempty" xml:"bankName,omitempty"`
	// example:
	//
	// 浙江省
	BankProvince *string `json:"bankProvince,omitempty" xml:"bankProvince,omitempty"`
	// example:
	//
	// ICBC
	BankShortNameCode *string `json:"bankShortNameCode,omitempty" xml:"bankShortNameCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALIPAY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// TO_PRI
	UsageType *string `json:"usageType,omitempty" xml:"usageType,omitempty"`
}

func (s ConsultCreateSubInstitutionRequestSettleInfo) String() string {
	return tea.Prettify(s)
}

func (s ConsultCreateSubInstitutionRequestSettleInfo) GoString() string {
	return s.String()
}

func (s *ConsultCreateSubInstitutionRequestSettleInfo) SetAccountId(v string) *ConsultCreateSubInstitutionRequestSettleInfo {
	s.AccountId = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSettleInfo) SetAccountName(v string) *ConsultCreateSubInstitutionRequestSettleInfo {
	s.AccountName = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSettleInfo) SetAccountType(v string) *ConsultCreateSubInstitutionRequestSettleInfo {
	s.AccountType = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSettleInfo) SetBankBranchName(v string) *ConsultCreateSubInstitutionRequestSettleInfo {
	s.BankBranchName = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSettleInfo) SetBankCity(v string) *ConsultCreateSubInstitutionRequestSettleInfo {
	s.BankCity = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSettleInfo) SetBankCode(v string) *ConsultCreateSubInstitutionRequestSettleInfo {
	s.BankCode = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSettleInfo) SetBankName(v string) *ConsultCreateSubInstitutionRequestSettleInfo {
	s.BankName = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSettleInfo) SetBankProvince(v string) *ConsultCreateSubInstitutionRequestSettleInfo {
	s.BankProvince = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSettleInfo) SetBankShortNameCode(v string) *ConsultCreateSubInstitutionRequestSettleInfo {
	s.BankShortNameCode = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSettleInfo) SetType(v string) *ConsultCreateSubInstitutionRequestSettleInfo {
	s.Type = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSettleInfo) SetUsageType(v string) *ConsultCreateSubInstitutionRequestSettleInfo {
	s.UsageType = &v
	return s
}

type ConsultCreateSubInstitutionRequestSubInstAddressInfo struct {
	// example:
	//
	// 未来park
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// 330100
	CityCode *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	// example:
	//
	// 330104
	DistrictCode *string `json:"districtCode,omitempty" xml:"districtCode,omitempty"`
	// example:
	//
	// 330000
	ProvinceCode *string `json:"provinceCode,omitempty" xml:"provinceCode,omitempty"`
}

func (s ConsultCreateSubInstitutionRequestSubInstAddressInfo) String() string {
	return tea.Prettify(s)
}

func (s ConsultCreateSubInstitutionRequestSubInstAddressInfo) GoString() string {
	return s.String()
}

func (s *ConsultCreateSubInstitutionRequestSubInstAddressInfo) SetAddress(v string) *ConsultCreateSubInstitutionRequestSubInstAddressInfo {
	s.Address = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSubInstAddressInfo) SetCityCode(v string) *ConsultCreateSubInstitutionRequestSubInstAddressInfo {
	s.CityCode = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSubInstAddressInfo) SetDistrictCode(v string) *ConsultCreateSubInstitutionRequestSubInstAddressInfo {
	s.DistrictCode = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSubInstAddressInfo) SetProvinceCode(v string) *ConsultCreateSubInstitutionRequestSubInstAddressInfo {
	s.ProvinceCode = &v
	return s
}

type ConsultCreateSubInstitutionRequestSubInstAuthInfo struct {
	// example:
	//
	// ossUrl
	AuthorizationLetterUrl *string `json:"authorizationLetterUrl,omitempty" xml:"authorizationLetterUrl,omitempty"`
}

func (s ConsultCreateSubInstitutionRequestSubInstAuthInfo) String() string {
	return tea.Prettify(s)
}

func (s ConsultCreateSubInstitutionRequestSubInstAuthInfo) GoString() string {
	return s.String()
}

func (s *ConsultCreateSubInstitutionRequestSubInstAuthInfo) SetAuthorizationLetterUrl(v string) *ConsultCreateSubInstitutionRequestSubInstAuthInfo {
	s.AuthorizationLetterUrl = &v
	return s
}

type ConsultCreateSubInstitutionRequestSubInstBasicInfo struct {
	// example:
	//
	// 一食堂
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5812
	Mcc *string `json:"mcc,omitempty" xml:"mcc,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 食堂
	SubInstName *string `json:"subInstName,omitempty" xml:"subInstName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 01
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ConsultCreateSubInstitutionRequestSubInstBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s ConsultCreateSubInstitutionRequestSubInstBasicInfo) GoString() string {
	return s.String()
}

func (s *ConsultCreateSubInstitutionRequestSubInstBasicInfo) SetAliasName(v string) *ConsultCreateSubInstitutionRequestSubInstBasicInfo {
	s.AliasName = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSubInstBasicInfo) SetMcc(v string) *ConsultCreateSubInstitutionRequestSubInstBasicInfo {
	s.Mcc = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSubInstBasicInfo) SetSubInstName(v string) *ConsultCreateSubInstitutionRequestSubInstBasicInfo {
	s.SubInstName = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSubInstBasicInfo) SetType(v string) *ConsultCreateSubInstitutionRequestSubInstBasicInfo {
	s.Type = &v
	return s
}

type ConsultCreateSubInstitutionRequestSubInstCertifyInfo struct {
	// example:
	//
	// ossUrl
	CertImage *string `json:"certImage,omitempty" xml:"certImage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 331081198611111111
	CertNo *string `json:"certNo,omitempty" xml:"certNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 201
	CertType *string `json:"certType,omitempty" xml:"certType,omitempty"`
}

func (s ConsultCreateSubInstitutionRequestSubInstCertifyInfo) String() string {
	return tea.Prettify(s)
}

func (s ConsultCreateSubInstitutionRequestSubInstCertifyInfo) GoString() string {
	return s.String()
}

func (s *ConsultCreateSubInstitutionRequestSubInstCertifyInfo) SetCertImage(v string) *ConsultCreateSubInstitutionRequestSubInstCertifyInfo {
	s.CertImage = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSubInstCertifyInfo) SetCertNo(v string) *ConsultCreateSubInstitutionRequestSubInstCertifyInfo {
	s.CertNo = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSubInstCertifyInfo) SetCertType(v string) *ConsultCreateSubInstitutionRequestSubInstCertifyInfo {
	s.CertType = &v
	return s
}

type ConsultCreateSubInstitutionRequestSubInstInvoiceInfo struct {
	// example:
	//
	// true
	AcceptElectronic *bool `json:"acceptElectronic,omitempty" xml:"acceptElectronic,omitempty"`
	// example:
	//
	// 浙江省杭州市西湖区西溪路蚂蚁金服
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// false
	AutoInvoice *bool `json:"autoInvoice,omitempty" xml:"autoInvoice,omitempty"`
	// example:
	//
	// 1234567812345678123
	BankAccount *string `json:"bankAccount,omitempty" xml:"bankAccount,omitempty"`
	// example:
	//
	// 中国银行
	BankName    *string                                                          `json:"bankName,omitempty" xml:"bankName,omitempty"`
	MailAddress *ConsultCreateSubInstitutionRequestSubInstInvoiceInfoMailAddress `json:"mailAddress,omitempty" xml:"mailAddress,omitempty" type:"Struct"`
	// example:
	//
	// 张三
	MailName *string `json:"mailName,omitempty" xml:"mailName,omitempty"`
	// example:
	//
	// 057162288888
	MailPhone *string `json:"mailPhone,omitempty" xml:"mailPhone,omitempty"`
	// example:
	//
	// 51010482542598631219
	TaxNo *string `json:"taxNo,omitempty" xml:"taxNo,omitempty"`
	// example:
	//
	// 01
	TaxPayerQualification *string `json:"taxPayerQualification,omitempty" xml:"taxPayerQualification,omitempty"`
	// example:
	//
	// 19981011
	TaxPayerValidDate *string `json:"taxPayerValidDate,omitempty" xml:"taxPayerValidDate,omitempty"`
	// example:
	//
	// 057162288888
	Telephone *string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	// example:
	//
	// **有限公司
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ConsultCreateSubInstitutionRequestSubInstInvoiceInfo) String() string {
	return tea.Prettify(s)
}

func (s ConsultCreateSubInstitutionRequestSubInstInvoiceInfo) GoString() string {
	return s.String()
}

func (s *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo) SetAcceptElectronic(v bool) *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo {
	s.AcceptElectronic = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo) SetAddress(v string) *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo {
	s.Address = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo) SetAutoInvoice(v bool) *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo {
	s.AutoInvoice = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo) SetBankAccount(v string) *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo {
	s.BankAccount = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo) SetBankName(v string) *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo {
	s.BankName = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo) SetMailAddress(v *ConsultCreateSubInstitutionRequestSubInstInvoiceInfoMailAddress) *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo {
	s.MailAddress = v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo) SetMailName(v string) *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo {
	s.MailName = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo) SetMailPhone(v string) *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo {
	s.MailPhone = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo) SetTaxNo(v string) *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo {
	s.TaxNo = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo) SetTaxPayerQualification(v string) *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo {
	s.TaxPayerQualification = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo) SetTaxPayerValidDate(v string) *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo {
	s.TaxPayerValidDate = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo) SetTelephone(v string) *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo {
	s.Telephone = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo) SetTitle(v string) *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo {
	s.Title = &v
	return s
}

type ConsultCreateSubInstitutionRequestSubInstInvoiceInfoMailAddress struct {
	// example:
	//
	// 未来park
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// 330100
	CityCode *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	// example:
	//
	// 330104
	DistrictCode *string `json:"districtCode,omitempty" xml:"districtCode,omitempty"`
	// example:
	//
	// 330000
	ProvinceCode *string `json:"provinceCode,omitempty" xml:"provinceCode,omitempty"`
}

func (s ConsultCreateSubInstitutionRequestSubInstInvoiceInfoMailAddress) String() string {
	return tea.Prettify(s)
}

func (s ConsultCreateSubInstitutionRequestSubInstInvoiceInfoMailAddress) GoString() string {
	return s.String()
}

func (s *ConsultCreateSubInstitutionRequestSubInstInvoiceInfoMailAddress) SetAddress(v string) *ConsultCreateSubInstitutionRequestSubInstInvoiceInfoMailAddress {
	s.Address = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSubInstInvoiceInfoMailAddress) SetCityCode(v string) *ConsultCreateSubInstitutionRequestSubInstInvoiceInfoMailAddress {
	s.CityCode = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSubInstInvoiceInfoMailAddress) SetDistrictCode(v string) *ConsultCreateSubInstitutionRequestSubInstInvoiceInfoMailAddress {
	s.DistrictCode = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSubInstInvoiceInfoMailAddress) SetProvinceCode(v string) *ConsultCreateSubInstitutionRequestSubInstInvoiceInfoMailAddress {
	s.ProvinceCode = &v
	return s
}

type ConsultCreateSubInstitutionRequestSubInstShopInfo struct {
	InDoorImages  []*string `json:"inDoorImages,omitempty" xml:"inDoorImages,omitempty" type:"Repeated"`
	OutDoorImages []*string `json:"outDoorImages,omitempty" xml:"outDoorImages,omitempty" type:"Repeated"`
}

func (s ConsultCreateSubInstitutionRequestSubInstShopInfo) String() string {
	return tea.Prettify(s)
}

func (s ConsultCreateSubInstitutionRequestSubInstShopInfo) GoString() string {
	return s.String()
}

func (s *ConsultCreateSubInstitutionRequestSubInstShopInfo) SetInDoorImages(v []*string) *ConsultCreateSubInstitutionRequestSubInstShopInfo {
	s.InDoorImages = v
	return s
}

func (s *ConsultCreateSubInstitutionRequestSubInstShopInfo) SetOutDoorImages(v []*string) *ConsultCreateSubInstitutionRequestSubInstShopInfo {
	s.OutDoorImages = v
	return s
}

type ConsultCreateSubInstitutionResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 202110110000001
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s ConsultCreateSubInstitutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConsultCreateSubInstitutionResponseBody) GoString() string {
	return s.String()
}

func (s *ConsultCreateSubInstitutionResponseBody) SetOrderId(v string) *ConsultCreateSubInstitutionResponseBody {
	s.OrderId = &v
	return s
}

type ConsultCreateSubInstitutionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConsultCreateSubInstitutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConsultCreateSubInstitutionResponse) String() string {
	return tea.Prettify(s)
}

func (s ConsultCreateSubInstitutionResponse) GoString() string {
	return s.String()
}

func (s *ConsultCreateSubInstitutionResponse) SetHeaders(v map[string]*string) *ConsultCreateSubInstitutionResponse {
	s.Headers = v
	return s
}

func (s *ConsultCreateSubInstitutionResponse) SetStatusCode(v int32) *ConsultCreateSubInstitutionResponse {
	s.StatusCode = &v
	return s
}

func (s *ConsultCreateSubInstitutionResponse) SetBody(v *ConsultCreateSubInstitutionResponseBody) *ConsultCreateSubInstitutionResponse {
	s.Body = v
	return s
}

type CreatWithholdingOrderAndPayHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreatWithholdingOrderAndPayHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreatWithholdingOrderAndPayHeaders) GoString() string {
	return s.String()
}

func (s *CreatWithholdingOrderAndPayHeaders) SetCommonHeaders(v map[string]*string) *CreatWithholdingOrderAndPayHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreatWithholdingOrderAndPayHeaders) SetXAcsDingtalkAccessToken(v string) *CreatWithholdingOrderAndPayHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreatWithholdingOrderAndPayRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10.01
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 202111090001
	InstId                        *string                                                            `json:"instId,omitempty" xml:"instId,omitempty"`
	OtherPayChannelDetailInfoList []*CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoList `json:"otherPayChannelDetailInfoList,omitempty" xml:"otherPayChannelDetailInfoList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 2021113000001
	OutTradeNo *string `json:"outTradeNo,omitempty" xml:"outTradeNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALIPAY
	PayChannel *string `json:"payChannel,omitempty" xml:"payChannel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2120493284
	PayerUserId *string `json:"payerUserId,omitempty" xml:"payerUserId,omitempty"`
	// example:
	//
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1001
	SubInstId *string `json:"subInstId,omitempty" xml:"subInstId,omitempty"`
	// example:
	//
	// 15m
	TimeOutExpress *string `json:"timeOutExpress,omitempty" xml:"timeOutExpress,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 餐费
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreatWithholdingOrderAndPayRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatWithholdingOrderAndPayRequest) GoString() string {
	return s.String()
}

func (s *CreatWithholdingOrderAndPayRequest) SetAmount(v string) *CreatWithholdingOrderAndPayRequest {
	s.Amount = &v
	return s
}

func (s *CreatWithholdingOrderAndPayRequest) SetInstId(v string) *CreatWithholdingOrderAndPayRequest {
	s.InstId = &v
	return s
}

func (s *CreatWithholdingOrderAndPayRequest) SetOtherPayChannelDetailInfoList(v []*CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoList) *CreatWithholdingOrderAndPayRequest {
	s.OtherPayChannelDetailInfoList = v
	return s
}

func (s *CreatWithholdingOrderAndPayRequest) SetOutTradeNo(v string) *CreatWithholdingOrderAndPayRequest {
	s.OutTradeNo = &v
	return s
}

func (s *CreatWithholdingOrderAndPayRequest) SetPayChannel(v string) *CreatWithholdingOrderAndPayRequest {
	s.PayChannel = &v
	return s
}

func (s *CreatWithholdingOrderAndPayRequest) SetPayerUserId(v string) *CreatWithholdingOrderAndPayRequest {
	s.PayerUserId = &v
	return s
}

func (s *CreatWithholdingOrderAndPayRequest) SetRemark(v string) *CreatWithholdingOrderAndPayRequest {
	s.Remark = &v
	return s
}

func (s *CreatWithholdingOrderAndPayRequest) SetSubInstId(v string) *CreatWithholdingOrderAndPayRequest {
	s.SubInstId = &v
	return s
}

func (s *CreatWithholdingOrderAndPayRequest) SetTimeOutExpress(v string) *CreatWithholdingOrderAndPayRequest {
	s.TimeOutExpress = &v
	return s
}

func (s *CreatWithholdingOrderAndPayRequest) SetTitle(v string) *CreatWithholdingOrderAndPayRequest {
	s.Title = &v
	return s
}

type CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoList struct {
	// This parameter is required.
	//
	// example:
	//
	// 5.00
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// This parameter is required.
	FundToolDetailInfoList []*CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoListFundToolDetailInfoList `json:"fundToolDetailInfoList,omitempty" xml:"fundToolDetailInfoList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 支付宝
	PayChannelName *string `json:"payChannelName,omitempty" xml:"payChannelName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021110100001
	PayChannelOrderNo *string `json:"payChannelOrderNo,omitempty" xml:"payChannelOrderNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALIPAY
	PayChannelType *string `json:"payChannelType,omitempty" xml:"payChannelType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4.00
	PromotionAmount *string `json:"promotionAmount,omitempty" xml:"promotionAmount,omitempty"`
}

func (s CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoList) String() string {
	return tea.Prettify(s)
}

func (s CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoList) GoString() string {
	return s.String()
}

func (s *CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoList) SetAmount(v string) *CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoList {
	s.Amount = &v
	return s
}

func (s *CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoList) SetFundToolDetailInfoList(v []*CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoListFundToolDetailInfoList) *CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoList {
	s.FundToolDetailInfoList = v
	return s
}

func (s *CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoList) SetPayChannelName(v string) *CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoList {
	s.PayChannelName = &v
	return s
}

func (s *CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoList) SetPayChannelOrderNo(v string) *CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoList {
	s.PayChannelOrderNo = &v
	return s
}

func (s *CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoList) SetPayChannelType(v string) *CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoList {
	s.PayChannelType = &v
	return s
}

func (s *CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoList) SetPromotionAmount(v string) *CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoList {
	s.PromotionAmount = &v
	return s
}

type CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoListFundToolDetailInfoList struct {
	// This parameter is required.
	//
	// example:
	//
	// 5.00
	Amount  *string `json:"amount,omitempty" xml:"amount,omitempty"`
	ExtInfo *string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 余额
	FundToolName *string `json:"fundToolName,omitempty" xml:"fundToolName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-15 10:10:10
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-15 10:10:11
	GmtFinish *string `json:"gmtFinish,omitempty" xml:"gmtFinish,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	PromotionFundTool *bool `json:"promotionFundTool,omitempty" xml:"promotionFundTool,omitempty"`
}

func (s CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoListFundToolDetailInfoList) String() string {
	return tea.Prettify(s)
}

func (s CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoListFundToolDetailInfoList) GoString() string {
	return s.String()
}

func (s *CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoListFundToolDetailInfoList) SetAmount(v string) *CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoListFundToolDetailInfoList {
	s.Amount = &v
	return s
}

func (s *CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoListFundToolDetailInfoList) SetExtInfo(v string) *CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoListFundToolDetailInfoList {
	s.ExtInfo = &v
	return s
}

func (s *CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoListFundToolDetailInfoList) SetFundToolName(v string) *CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoListFundToolDetailInfoList {
	s.FundToolName = &v
	return s
}

func (s *CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoListFundToolDetailInfoList) SetGmtCreate(v string) *CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoListFundToolDetailInfoList {
	s.GmtCreate = &v
	return s
}

func (s *CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoListFundToolDetailInfoList) SetGmtFinish(v string) *CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoListFundToolDetailInfoList {
	s.GmtFinish = &v
	return s
}

func (s *CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoListFundToolDetailInfoList) SetPromotionFundTool(v bool) *CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoListFundToolDetailInfoList {
	s.PromotionFundTool = &v
	return s
}

type CreatWithholdingOrderAndPayResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 10.01
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// example:
	//
	// 2021-11-15 10:10:10
	GmtPay *string `json:"gmtPay,omitempty" xml:"gmtPay,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 202111010001
	InstId *string `json:"instId,omitempty" xml:"instId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 202121241343151
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 202111020001
	OutTradeNo *string `json:"outTradeNo,omitempty" xml:"outTradeNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALIPAY
	PayChannel *string `json:"payChannel,omitempty" xml:"payChannel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13****09
	PayChannelAccountNo *string `json:"payChannelAccountNo,omitempty" xml:"payChannelAccountNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123124
	PayerStaffId *string `json:"payerStaffId,omitempty" xml:"payerStaffId,omitempty"`
	// example:
	//
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1001
	SubInstId *string `json:"subInstId,omitempty" xml:"subInstId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 餐费
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreatWithholdingOrderAndPayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatWithholdingOrderAndPayResponseBody) GoString() string {
	return s.String()
}

func (s *CreatWithholdingOrderAndPayResponseBody) SetAmount(v string) *CreatWithholdingOrderAndPayResponseBody {
	s.Amount = &v
	return s
}

func (s *CreatWithholdingOrderAndPayResponseBody) SetGmtPay(v string) *CreatWithholdingOrderAndPayResponseBody {
	s.GmtPay = &v
	return s
}

func (s *CreatWithholdingOrderAndPayResponseBody) SetInstId(v string) *CreatWithholdingOrderAndPayResponseBody {
	s.InstId = &v
	return s
}

func (s *CreatWithholdingOrderAndPayResponseBody) SetOrderNo(v string) *CreatWithholdingOrderAndPayResponseBody {
	s.OrderNo = &v
	return s
}

func (s *CreatWithholdingOrderAndPayResponseBody) SetOutTradeNo(v string) *CreatWithholdingOrderAndPayResponseBody {
	s.OutTradeNo = &v
	return s
}

func (s *CreatWithholdingOrderAndPayResponseBody) SetPayChannel(v string) *CreatWithholdingOrderAndPayResponseBody {
	s.PayChannel = &v
	return s
}

func (s *CreatWithholdingOrderAndPayResponseBody) SetPayChannelAccountNo(v string) *CreatWithholdingOrderAndPayResponseBody {
	s.PayChannelAccountNo = &v
	return s
}

func (s *CreatWithholdingOrderAndPayResponseBody) SetPayerStaffId(v string) *CreatWithholdingOrderAndPayResponseBody {
	s.PayerStaffId = &v
	return s
}

func (s *CreatWithholdingOrderAndPayResponseBody) SetRemark(v string) *CreatWithholdingOrderAndPayResponseBody {
	s.Remark = &v
	return s
}

func (s *CreatWithholdingOrderAndPayResponseBody) SetStatus(v string) *CreatWithholdingOrderAndPayResponseBody {
	s.Status = &v
	return s
}

func (s *CreatWithholdingOrderAndPayResponseBody) SetSubInstId(v string) *CreatWithholdingOrderAndPayResponseBody {
	s.SubInstId = &v
	return s
}

func (s *CreatWithholdingOrderAndPayResponseBody) SetTitle(v string) *CreatWithholdingOrderAndPayResponseBody {
	s.Title = &v
	return s
}

type CreatWithholdingOrderAndPayResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatWithholdingOrderAndPayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatWithholdingOrderAndPayResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatWithholdingOrderAndPayResponse) GoString() string {
	return s.String()
}

func (s *CreatWithholdingOrderAndPayResponse) SetHeaders(v map[string]*string) *CreatWithholdingOrderAndPayResponse {
	s.Headers = v
	return s
}

func (s *CreatWithholdingOrderAndPayResponse) SetStatusCode(v int32) *CreatWithholdingOrderAndPayResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatWithholdingOrderAndPayResponse) SetBody(v *CreatWithholdingOrderAndPayResponseBody) *CreatWithholdingOrderAndPayResponse {
	s.Body = v
	return s
}

type CreateAcquireRefundOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateAcquireRefundOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateAcquireRefundOrderHeaders) GoString() string {
	return s.String()
}

func (s *CreateAcquireRefundOrderHeaders) SetCommonHeaders(v map[string]*string) *CreateAcquireRefundOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateAcquireRefundOrderHeaders) SetXAcsDingtalkAccessToken(v string) *CreateAcquireRefundOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateAcquireRefundOrderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 202111090001
	InstId *string `json:"instId,omitempty" xml:"instId,omitempty"`
	// example:
	//
	// 2120493284
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021113000001
	OriginOutTradeNo              *string                                                         `json:"originOutTradeNo,omitempty" xml:"originOutTradeNo,omitempty"`
	OtherPayChannelDetailInfoList []*CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoList `json:"otherPayChannelDetailInfoList,omitempty" xml:"otherPayChannelDetailInfoList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// r2021113000001
	OutRefundNo *string `json:"outRefundNo,omitempty" xml:"outRefundNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10.01
	RefundAmount *string `json:"refundAmount,omitempty" xml:"refundAmount,omitempty"`
	// example:
	//
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1001
	SubInstId *string `json:"subInstId,omitempty" xml:"subInstId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 餐费
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateAcquireRefundOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAcquireRefundOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateAcquireRefundOrderRequest) SetInstId(v string) *CreateAcquireRefundOrderRequest {
	s.InstId = &v
	return s
}

func (s *CreateAcquireRefundOrderRequest) SetOperatorUserId(v string) *CreateAcquireRefundOrderRequest {
	s.OperatorUserId = &v
	return s
}

func (s *CreateAcquireRefundOrderRequest) SetOriginOutTradeNo(v string) *CreateAcquireRefundOrderRequest {
	s.OriginOutTradeNo = &v
	return s
}

func (s *CreateAcquireRefundOrderRequest) SetOtherPayChannelDetailInfoList(v []*CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoList) *CreateAcquireRefundOrderRequest {
	s.OtherPayChannelDetailInfoList = v
	return s
}

func (s *CreateAcquireRefundOrderRequest) SetOutRefundNo(v string) *CreateAcquireRefundOrderRequest {
	s.OutRefundNo = &v
	return s
}

func (s *CreateAcquireRefundOrderRequest) SetRefundAmount(v string) *CreateAcquireRefundOrderRequest {
	s.RefundAmount = &v
	return s
}

func (s *CreateAcquireRefundOrderRequest) SetRemark(v string) *CreateAcquireRefundOrderRequest {
	s.Remark = &v
	return s
}

func (s *CreateAcquireRefundOrderRequest) SetSubInstId(v string) *CreateAcquireRefundOrderRequest {
	s.SubInstId = &v
	return s
}

func (s *CreateAcquireRefundOrderRequest) SetTitle(v string) *CreateAcquireRefundOrderRequest {
	s.Title = &v
	return s
}

type CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoList struct {
	// This parameter is required.
	//
	// example:
	//
	// 5.00
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// This parameter is required.
	FundToolDetailInfoList []*CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoListFundToolDetailInfoList `json:"fundToolDetailInfoList,omitempty" xml:"fundToolDetailInfoList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 支付宝
	PayChannelName *string `json:"payChannelName,omitempty" xml:"payChannelName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021110100001
	PayChannelOrderNo *string `json:"payChannelOrderNo,omitempty" xml:"payChannelOrderNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALIPAY
	PayChannelType *string `json:"payChannelType,omitempty" xml:"payChannelType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4.00
	PromotionAmount *string `json:"promotionAmount,omitempty" xml:"promotionAmount,omitempty"`
}

func (s CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoList) String() string {
	return tea.Prettify(s)
}

func (s CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoList) GoString() string {
	return s.String()
}

func (s *CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoList) SetAmount(v string) *CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoList {
	s.Amount = &v
	return s
}

func (s *CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoList) SetFundToolDetailInfoList(v []*CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoListFundToolDetailInfoList) *CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoList {
	s.FundToolDetailInfoList = v
	return s
}

func (s *CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoList) SetPayChannelName(v string) *CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoList {
	s.PayChannelName = &v
	return s
}

func (s *CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoList) SetPayChannelOrderNo(v string) *CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoList {
	s.PayChannelOrderNo = &v
	return s
}

func (s *CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoList) SetPayChannelType(v string) *CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoList {
	s.PayChannelType = &v
	return s
}

func (s *CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoList) SetPromotionAmount(v string) *CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoList {
	s.PromotionAmount = &v
	return s
}

type CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoListFundToolDetailInfoList struct {
	// This parameter is required.
	//
	// example:
	//
	// 5.00
	Amount  *string `json:"amount,omitempty" xml:"amount,omitempty"`
	ExtInfo *string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 余额
	FundToolName *string `json:"fundToolName,omitempty" xml:"fundToolName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-15 10:10:10
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-15 10:10:11
	GmtFinish *string `json:"gmtFinish,omitempty" xml:"gmtFinish,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	PromotionFundTool *bool `json:"promotionFundTool,omitempty" xml:"promotionFundTool,omitempty"`
}

func (s CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoListFundToolDetailInfoList) String() string {
	return tea.Prettify(s)
}

func (s CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoListFundToolDetailInfoList) GoString() string {
	return s.String()
}

func (s *CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoListFundToolDetailInfoList) SetAmount(v string) *CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoListFundToolDetailInfoList {
	s.Amount = &v
	return s
}

func (s *CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoListFundToolDetailInfoList) SetExtInfo(v string) *CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoListFundToolDetailInfoList {
	s.ExtInfo = &v
	return s
}

func (s *CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoListFundToolDetailInfoList) SetFundToolName(v string) *CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoListFundToolDetailInfoList {
	s.FundToolName = &v
	return s
}

func (s *CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoListFundToolDetailInfoList) SetGmtCreate(v string) *CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoListFundToolDetailInfoList {
	s.GmtCreate = &v
	return s
}

func (s *CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoListFundToolDetailInfoList) SetGmtFinish(v string) *CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoListFundToolDetailInfoList {
	s.GmtFinish = &v
	return s
}

func (s *CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoListFundToolDetailInfoList) SetPromotionFundTool(v bool) *CreateAcquireRefundOrderRequestOtherPayChannelDetailInfoListFundToolDetailInfoList {
	s.PromotionFundTool = &v
	return s
}

type CreateAcquireRefundOrderResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// r2021113000001
	OutRefundNo *string `json:"outRefundNo,omitempty" xml:"outRefundNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 202111110000111
	RefundOrderNo *string `json:"refundOrderNo,omitempty" xml:"refundOrderNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateAcquireRefundOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAcquireRefundOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAcquireRefundOrderResponseBody) SetOutRefundNo(v string) *CreateAcquireRefundOrderResponseBody {
	s.OutRefundNo = &v
	return s
}

func (s *CreateAcquireRefundOrderResponseBody) SetRefundOrderNo(v string) *CreateAcquireRefundOrderResponseBody {
	s.RefundOrderNo = &v
	return s
}

func (s *CreateAcquireRefundOrderResponseBody) SetStatus(v string) *CreateAcquireRefundOrderResponseBody {
	s.Status = &v
	return s
}

type CreateAcquireRefundOrderResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAcquireRefundOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAcquireRefundOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAcquireRefundOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateAcquireRefundOrderResponse) SetHeaders(v map[string]*string) *CreateAcquireRefundOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateAcquireRefundOrderResponse) SetStatusCode(v int32) *CreateAcquireRefundOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAcquireRefundOrderResponse) SetBody(v *CreateAcquireRefundOrderResponseBody) *CreateAcquireRefundOrderResponse {
	s.Body = v
	return s
}

type CreateBatchTradeOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateBatchTradeOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchTradeOrderHeaders) GoString() string {
	return s.String()
}

func (s *CreateBatchTradeOrderHeaders) SetCommonHeaders(v map[string]*string) *CreateBatchTradeOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateBatchTradeOrderHeaders) SetXAcsDingtalkAccessToken(v string) *CreateBatchTradeOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateBatchTradeOrderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2021070712440326300185114
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13****9
	AccountNo *string `json:"accountNo,omitempty" xml:"accountNo,omitempty"`
	// example:
	//
	// 备注
	BatchRemark *string `json:"batchRemark,omitempty" xml:"batchRemark,omitempty"`
	// This parameter is required.
	BatchTradeDetails []*CreateBatchTradeOrderRequestBatchTradeDetails `json:"batchTradeDetails,omitempty" xml:"batchTradeDetails,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 20210901001
	OutBatchNo *string `json:"outBatchNo,omitempty" xml:"outBatchNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8476212471
	StaffId *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1.00
	TotalAmount *string `json:"totalAmount,omitempty" xml:"totalAmount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 工资
	TradeTitle *string `json:"tradeTitle,omitempty" xml:"tradeTitle,omitempty"`
}

func (s CreateBatchTradeOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchTradeOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateBatchTradeOrderRequest) SetAccountId(v string) *CreateBatchTradeOrderRequest {
	s.AccountId = &v
	return s
}

func (s *CreateBatchTradeOrderRequest) SetAccountNo(v string) *CreateBatchTradeOrderRequest {
	s.AccountNo = &v
	return s
}

func (s *CreateBatchTradeOrderRequest) SetBatchRemark(v string) *CreateBatchTradeOrderRequest {
	s.BatchRemark = &v
	return s
}

func (s *CreateBatchTradeOrderRequest) SetBatchTradeDetails(v []*CreateBatchTradeOrderRequestBatchTradeDetails) *CreateBatchTradeOrderRequest {
	s.BatchTradeDetails = v
	return s
}

func (s *CreateBatchTradeOrderRequest) SetOutBatchNo(v string) *CreateBatchTradeOrderRequest {
	s.OutBatchNo = &v
	return s
}

func (s *CreateBatchTradeOrderRequest) SetStaffId(v string) *CreateBatchTradeOrderRequest {
	s.StaffId = &v
	return s
}

func (s *CreateBatchTradeOrderRequest) SetTotalAmount(v string) *CreateBatchTradeOrderRequest {
	s.TotalAmount = &v
	return s
}

func (s *CreateBatchTradeOrderRequest) SetTotalCount(v int64) *CreateBatchTradeOrderRequest {
	s.TotalCount = &v
	return s
}

func (s *CreateBatchTradeOrderRequest) SetTradeTitle(v string) *CreateBatchTradeOrderRequest {
	s.TradeTitle = &v
	return s
}

type CreateBatchTradeOrderRequestBatchTradeDetails struct {
	// This parameter is required.
	//
	// example:
	//
	// 1.00
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// example:
	//
	// 工资
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试
	PayeeAccountName *string `json:"payeeAccountName,omitempty" xml:"payeeAccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13000000000
	PayeeAccountNo *string `json:"payeeAccountNo,omitempty" xml:"payeeAccountNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALIPAY
	PayeeAccountType *string `json:"payeeAccountType,omitempty" xml:"payeeAccountType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SerialNo *int64 `json:"serialNo,omitempty" xml:"serialNo,omitempty"`
}

func (s CreateBatchTradeOrderRequestBatchTradeDetails) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchTradeOrderRequestBatchTradeDetails) GoString() string {
	return s.String()
}

func (s *CreateBatchTradeOrderRequestBatchTradeDetails) SetAmount(v string) *CreateBatchTradeOrderRequestBatchTradeDetails {
	s.Amount = &v
	return s
}

func (s *CreateBatchTradeOrderRequestBatchTradeDetails) SetMemo(v string) *CreateBatchTradeOrderRequestBatchTradeDetails {
	s.Memo = &v
	return s
}

func (s *CreateBatchTradeOrderRequestBatchTradeDetails) SetPayeeAccountName(v string) *CreateBatchTradeOrderRequestBatchTradeDetails {
	s.PayeeAccountName = &v
	return s
}

func (s *CreateBatchTradeOrderRequestBatchTradeDetails) SetPayeeAccountNo(v string) *CreateBatchTradeOrderRequestBatchTradeDetails {
	s.PayeeAccountNo = &v
	return s
}

func (s *CreateBatchTradeOrderRequestBatchTradeDetails) SetPayeeAccountType(v string) *CreateBatchTradeOrderRequestBatchTradeDetails {
	s.PayeeAccountType = &v
	return s
}

func (s *CreateBatchTradeOrderRequestBatchTradeDetails) SetSerialNo(v int64) *CreateBatchTradeOrderRequestBatchTradeDetails {
	s.SerialNo = &v
	return s
}

type CreateBatchTradeOrderResponseBody struct {
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	// This parameter is required.
	OrderStatus *string `json:"orderStatus,omitempty" xml:"orderStatus,omitempty"`
	// This parameter is required.
	OutBatchNo *string `json:"outBatchNo,omitempty" xml:"outBatchNo,omitempty"`
}

func (s CreateBatchTradeOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchTradeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBatchTradeOrderResponseBody) SetOrderNo(v string) *CreateBatchTradeOrderResponseBody {
	s.OrderNo = &v
	return s
}

func (s *CreateBatchTradeOrderResponseBody) SetOrderStatus(v string) *CreateBatchTradeOrderResponseBody {
	s.OrderStatus = &v
	return s
}

func (s *CreateBatchTradeOrderResponseBody) SetOutBatchNo(v string) *CreateBatchTradeOrderResponseBody {
	s.OutBatchNo = &v
	return s
}

type CreateBatchTradeOrderResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBatchTradeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBatchTradeOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchTradeOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateBatchTradeOrderResponse) SetHeaders(v map[string]*string) *CreateBatchTradeOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateBatchTradeOrderResponse) SetStatusCode(v int32) *CreateBatchTradeOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBatchTradeOrderResponse) SetBody(v *CreateBatchTradeOrderResponseBody) *CreateBatchTradeOrderResponse {
	s.Body = v
	return s
}

type CreateSubInstitutionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateSubInstitutionHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateSubInstitutionHeaders) GoString() string {
	return s.String()
}

func (s *CreateSubInstitutionHeaders) SetCommonHeaders(v map[string]*string) *CreateSubInstitutionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSubInstitutionHeaders) SetXAcsDingtalkAccessToken(v string) *CreateSubInstitutionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateSubInstitutionRequest struct {
	// example:
	//
	// asdf@163.com
	BindingAlipayLogonId *string `json:"bindingAlipayLogonId,omitempty" xml:"bindingAlipayLogonId,omitempty"`
	// This parameter is required.
	ContactInfo *CreateSubInstitutionRequestContactInfo `json:"contactInfo,omitempty" xml:"contactInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 202111090001
	InstId *string `json:"instId,omitempty" xml:"instId,omitempty"`
	// This parameter is required.
	LegalPersonCertInfo *CreateSubInstitutionRequestLegalPersonCertInfo `json:"legalPersonCertInfo,omitempty" xml:"legalPersonCertInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 2021000001
	OutTradeNo *string `json:"outTradeNo,omitempty" xml:"outTradeNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALIPAY
	PayChannel         *string                                          `json:"payChannel,omitempty" xml:"payChannel,omitempty"`
	QualificationInfos []*CreateSubInstitutionRequestQualificationInfos `json:"qualificationInfos,omitempty" xml:"qualificationInfos,omitempty" type:"Repeated"`
	// This parameter is required.
	Services []*string `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
	// This parameter is required.
	SettleInfo *CreateSubInstitutionRequestSettleInfo `json:"settleInfo,omitempty" xml:"settleInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// S001
	Solution           *string                                        `json:"solution,omitempty" xml:"solution,omitempty"`
	SubInstAddressInfo *CreateSubInstitutionRequestSubInstAddressInfo `json:"subInstAddressInfo,omitempty" xml:"subInstAddressInfo,omitempty" type:"Struct"`
	SubInstAuthInfo    *CreateSubInstitutionRequestSubInstAuthInfo    `json:"subInstAuthInfo,omitempty" xml:"subInstAuthInfo,omitempty" type:"Struct"`
	// This parameter is required.
	SubInstBasicInfo *CreateSubInstitutionRequestSubInstBasicInfo `json:"subInstBasicInfo,omitempty" xml:"subInstBasicInfo,omitempty" type:"Struct"`
	// This parameter is required.
	SubInstCertifyInfo *CreateSubInstitutionRequestSubInstCertifyInfo `json:"subInstCertifyInfo,omitempty" xml:"subInstCertifyInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 1001
	SubInstId          *string                                        `json:"subInstId,omitempty" xml:"subInstId,omitempty"`
	SubInstInvoiceInfo *CreateSubInstitutionRequestSubInstInvoiceInfo `json:"subInstInvoiceInfo,omitempty" xml:"subInstInvoiceInfo,omitempty" type:"Struct"`
	SubInstShopInfo    *CreateSubInstitutionRequestSubInstShopInfo    `json:"subInstShopInfo,omitempty" xml:"subInstShopInfo,omitempty" type:"Struct"`
}

func (s CreateSubInstitutionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSubInstitutionRequest) GoString() string {
	return s.String()
}

func (s *CreateSubInstitutionRequest) SetBindingAlipayLogonId(v string) *CreateSubInstitutionRequest {
	s.BindingAlipayLogonId = &v
	return s
}

func (s *CreateSubInstitutionRequest) SetContactInfo(v *CreateSubInstitutionRequestContactInfo) *CreateSubInstitutionRequest {
	s.ContactInfo = v
	return s
}

func (s *CreateSubInstitutionRequest) SetInstId(v string) *CreateSubInstitutionRequest {
	s.InstId = &v
	return s
}

func (s *CreateSubInstitutionRequest) SetLegalPersonCertInfo(v *CreateSubInstitutionRequestLegalPersonCertInfo) *CreateSubInstitutionRequest {
	s.LegalPersonCertInfo = v
	return s
}

func (s *CreateSubInstitutionRequest) SetOutTradeNo(v string) *CreateSubInstitutionRequest {
	s.OutTradeNo = &v
	return s
}

func (s *CreateSubInstitutionRequest) SetPayChannel(v string) *CreateSubInstitutionRequest {
	s.PayChannel = &v
	return s
}

func (s *CreateSubInstitutionRequest) SetQualificationInfos(v []*CreateSubInstitutionRequestQualificationInfos) *CreateSubInstitutionRequest {
	s.QualificationInfos = v
	return s
}

func (s *CreateSubInstitutionRequest) SetServices(v []*string) *CreateSubInstitutionRequest {
	s.Services = v
	return s
}

func (s *CreateSubInstitutionRequest) SetSettleInfo(v *CreateSubInstitutionRequestSettleInfo) *CreateSubInstitutionRequest {
	s.SettleInfo = v
	return s
}

func (s *CreateSubInstitutionRequest) SetSolution(v string) *CreateSubInstitutionRequest {
	s.Solution = &v
	return s
}

func (s *CreateSubInstitutionRequest) SetSubInstAddressInfo(v *CreateSubInstitutionRequestSubInstAddressInfo) *CreateSubInstitutionRequest {
	s.SubInstAddressInfo = v
	return s
}

func (s *CreateSubInstitutionRequest) SetSubInstAuthInfo(v *CreateSubInstitutionRequestSubInstAuthInfo) *CreateSubInstitutionRequest {
	s.SubInstAuthInfo = v
	return s
}

func (s *CreateSubInstitutionRequest) SetSubInstBasicInfo(v *CreateSubInstitutionRequestSubInstBasicInfo) *CreateSubInstitutionRequest {
	s.SubInstBasicInfo = v
	return s
}

func (s *CreateSubInstitutionRequest) SetSubInstCertifyInfo(v *CreateSubInstitutionRequestSubInstCertifyInfo) *CreateSubInstitutionRequest {
	s.SubInstCertifyInfo = v
	return s
}

func (s *CreateSubInstitutionRequest) SetSubInstId(v string) *CreateSubInstitutionRequest {
	s.SubInstId = &v
	return s
}

func (s *CreateSubInstitutionRequest) SetSubInstInvoiceInfo(v *CreateSubInstitutionRequestSubInstInvoiceInfo) *CreateSubInstitutionRequest {
	s.SubInstInvoiceInfo = v
	return s
}

func (s *CreateSubInstitutionRequest) SetSubInstShopInfo(v *CreateSubInstitutionRequestSubInstShopInfo) *CreateSubInstitutionRequest {
	s.SubInstShopInfo = v
	return s
}

type CreateSubInstitutionRequestContactInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 李某某
	ContactName *string `json:"contactName,omitempty" xml:"contactName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13900000000
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
}

func (s CreateSubInstitutionRequestContactInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateSubInstitutionRequestContactInfo) GoString() string {
	return s.String()
}

func (s *CreateSubInstitutionRequestContactInfo) SetContactName(v string) *CreateSubInstitutionRequestContactInfo {
	s.ContactName = &v
	return s
}

func (s *CreateSubInstitutionRequestContactInfo) SetMobile(v string) *CreateSubInstitutionRequestContactInfo {
	s.Mobile = &v
	return s
}

type CreateSubInstitutionRequestLegalPersonCertInfo struct {
	// example:
	//
	// ossUrl
	CertBackImage *string `json:"certBackImage,omitempty" xml:"certBackImage,omitempty"`
	// example:
	//
	// ossUrl
	CertFrontImage *string `json:"certFrontImage,omitempty" xml:"certFrontImage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 李某某
	CertName *string `json:"certName,omitempty" xml:"certName,omitempty"`
	// example:
	//
	// 100
	CertType *string `json:"certType,omitempty" xml:"certType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 330104200010109999
	IdCardNo *string `json:"idCardNo,omitempty" xml:"idCardNo,omitempty"`
}

func (s CreateSubInstitutionRequestLegalPersonCertInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateSubInstitutionRequestLegalPersonCertInfo) GoString() string {
	return s.String()
}

func (s *CreateSubInstitutionRequestLegalPersonCertInfo) SetCertBackImage(v string) *CreateSubInstitutionRequestLegalPersonCertInfo {
	s.CertBackImage = &v
	return s
}

func (s *CreateSubInstitutionRequestLegalPersonCertInfo) SetCertFrontImage(v string) *CreateSubInstitutionRequestLegalPersonCertInfo {
	s.CertFrontImage = &v
	return s
}

func (s *CreateSubInstitutionRequestLegalPersonCertInfo) SetCertName(v string) *CreateSubInstitutionRequestLegalPersonCertInfo {
	s.CertName = &v
	return s
}

func (s *CreateSubInstitutionRequestLegalPersonCertInfo) SetCertType(v string) *CreateSubInstitutionRequestLegalPersonCertInfo {
	s.CertType = &v
	return s
}

func (s *CreateSubInstitutionRequestLegalPersonCertInfo) SetIdCardNo(v string) *CreateSubInstitutionRequestLegalPersonCertInfo {
	s.IdCardNo = &v
	return s
}

type CreateSubInstitutionRequestQualificationInfos struct {
	// This parameter is required.
	//
	// example:
	//
	// ossUrl
	QualificationImage *string `json:"qualificationImage,omitempty" xml:"qualificationImage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 301
	QualificationType *string `json:"qualificationType,omitempty" xml:"qualificationType,omitempty"`
}

func (s CreateSubInstitutionRequestQualificationInfos) String() string {
	return tea.Prettify(s)
}

func (s CreateSubInstitutionRequestQualificationInfos) GoString() string {
	return s.String()
}

func (s *CreateSubInstitutionRequestQualificationInfos) SetQualificationImage(v string) *CreateSubInstitutionRequestQualificationInfos {
	s.QualificationImage = &v
	return s
}

func (s *CreateSubInstitutionRequestQualificationInfos) SetQualificationType(v string) *CreateSubInstitutionRequestQualificationInfos {
	s.QualificationType = &v
	return s
}

type CreateSubInstitutionRequestSettleInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 622202120200000000
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// 李某某
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	// example:
	//
	// DEBIT_CARD
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	// example:
	//
	// 城东支行
	BankBranchName *string `json:"bankBranchName,omitempty" xml:"bankBranchName,omitempty"`
	// example:
	//
	// 杭州市
	BankCity *string `json:"bankCity,omitempty" xml:"bankCity,omitempty"`
	// example:
	//
	// 313791000023
	BankCode *string `json:"bankCode,omitempty" xml:"bankCode,omitempty"`
	// example:
	//
	// 工商银行
	BankName *string `json:"bankName,omitempty" xml:"bankName,omitempty"`
	// example:
	//
	// 浙江省
	BankProvince *string `json:"bankProvince,omitempty" xml:"bankProvince,omitempty"`
	// example:
	//
	// ICBC
	BankShortNameCode *string `json:"bankShortNameCode,omitempty" xml:"bankShortNameCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALIPAY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// TO_PRI
	UsageType *string `json:"usageType,omitempty" xml:"usageType,omitempty"`
}

func (s CreateSubInstitutionRequestSettleInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateSubInstitutionRequestSettleInfo) GoString() string {
	return s.String()
}

func (s *CreateSubInstitutionRequestSettleInfo) SetAccountId(v string) *CreateSubInstitutionRequestSettleInfo {
	s.AccountId = &v
	return s
}

func (s *CreateSubInstitutionRequestSettleInfo) SetAccountName(v string) *CreateSubInstitutionRequestSettleInfo {
	s.AccountName = &v
	return s
}

func (s *CreateSubInstitutionRequestSettleInfo) SetAccountType(v string) *CreateSubInstitutionRequestSettleInfo {
	s.AccountType = &v
	return s
}

func (s *CreateSubInstitutionRequestSettleInfo) SetBankBranchName(v string) *CreateSubInstitutionRequestSettleInfo {
	s.BankBranchName = &v
	return s
}

func (s *CreateSubInstitutionRequestSettleInfo) SetBankCity(v string) *CreateSubInstitutionRequestSettleInfo {
	s.BankCity = &v
	return s
}

func (s *CreateSubInstitutionRequestSettleInfo) SetBankCode(v string) *CreateSubInstitutionRequestSettleInfo {
	s.BankCode = &v
	return s
}

func (s *CreateSubInstitutionRequestSettleInfo) SetBankName(v string) *CreateSubInstitutionRequestSettleInfo {
	s.BankName = &v
	return s
}

func (s *CreateSubInstitutionRequestSettleInfo) SetBankProvince(v string) *CreateSubInstitutionRequestSettleInfo {
	s.BankProvince = &v
	return s
}

func (s *CreateSubInstitutionRequestSettleInfo) SetBankShortNameCode(v string) *CreateSubInstitutionRequestSettleInfo {
	s.BankShortNameCode = &v
	return s
}

func (s *CreateSubInstitutionRequestSettleInfo) SetType(v string) *CreateSubInstitutionRequestSettleInfo {
	s.Type = &v
	return s
}

func (s *CreateSubInstitutionRequestSettleInfo) SetUsageType(v string) *CreateSubInstitutionRequestSettleInfo {
	s.UsageType = &v
	return s
}

type CreateSubInstitutionRequestSubInstAddressInfo struct {
	// example:
	//
	// 未来park
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// 330100
	CityCode *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	// example:
	//
	// 330104
	DistrictCode *string `json:"districtCode,omitempty" xml:"districtCode,omitempty"`
	// example:
	//
	// 330000
	ProvinceCode *string `json:"provinceCode,omitempty" xml:"provinceCode,omitempty"`
}

func (s CreateSubInstitutionRequestSubInstAddressInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateSubInstitutionRequestSubInstAddressInfo) GoString() string {
	return s.String()
}

func (s *CreateSubInstitutionRequestSubInstAddressInfo) SetAddress(v string) *CreateSubInstitutionRequestSubInstAddressInfo {
	s.Address = &v
	return s
}

func (s *CreateSubInstitutionRequestSubInstAddressInfo) SetCityCode(v string) *CreateSubInstitutionRequestSubInstAddressInfo {
	s.CityCode = &v
	return s
}

func (s *CreateSubInstitutionRequestSubInstAddressInfo) SetDistrictCode(v string) *CreateSubInstitutionRequestSubInstAddressInfo {
	s.DistrictCode = &v
	return s
}

func (s *CreateSubInstitutionRequestSubInstAddressInfo) SetProvinceCode(v string) *CreateSubInstitutionRequestSubInstAddressInfo {
	s.ProvinceCode = &v
	return s
}

type CreateSubInstitutionRequestSubInstAuthInfo struct {
	// example:
	//
	// ossUrl
	AuthorizationLetterUrl *string `json:"authorizationLetterUrl,omitempty" xml:"authorizationLetterUrl,omitempty"`
}

func (s CreateSubInstitutionRequestSubInstAuthInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateSubInstitutionRequestSubInstAuthInfo) GoString() string {
	return s.String()
}

func (s *CreateSubInstitutionRequestSubInstAuthInfo) SetAuthorizationLetterUrl(v string) *CreateSubInstitutionRequestSubInstAuthInfo {
	s.AuthorizationLetterUrl = &v
	return s
}

type CreateSubInstitutionRequestSubInstBasicInfo struct {
	// example:
	//
	// 一食堂
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5812
	Mcc *string `json:"mcc,omitempty" xml:"mcc,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 食堂
	SubInstName *string `json:"subInstName,omitempty" xml:"subInstName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 01
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateSubInstitutionRequestSubInstBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateSubInstitutionRequestSubInstBasicInfo) GoString() string {
	return s.String()
}

func (s *CreateSubInstitutionRequestSubInstBasicInfo) SetAliasName(v string) *CreateSubInstitutionRequestSubInstBasicInfo {
	s.AliasName = &v
	return s
}

func (s *CreateSubInstitutionRequestSubInstBasicInfo) SetMcc(v string) *CreateSubInstitutionRequestSubInstBasicInfo {
	s.Mcc = &v
	return s
}

func (s *CreateSubInstitutionRequestSubInstBasicInfo) SetSubInstName(v string) *CreateSubInstitutionRequestSubInstBasicInfo {
	s.SubInstName = &v
	return s
}

func (s *CreateSubInstitutionRequestSubInstBasicInfo) SetType(v string) *CreateSubInstitutionRequestSubInstBasicInfo {
	s.Type = &v
	return s
}

type CreateSubInstitutionRequestSubInstCertifyInfo struct {
	// example:
	//
	// ossUrl
	CertImage *string `json:"certImage,omitempty" xml:"certImage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 331081198611111111
	CertNo *string `json:"certNo,omitempty" xml:"certNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 201
	CertType *string `json:"certType,omitempty" xml:"certType,omitempty"`
}

func (s CreateSubInstitutionRequestSubInstCertifyInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateSubInstitutionRequestSubInstCertifyInfo) GoString() string {
	return s.String()
}

func (s *CreateSubInstitutionRequestSubInstCertifyInfo) SetCertImage(v string) *CreateSubInstitutionRequestSubInstCertifyInfo {
	s.CertImage = &v
	return s
}

func (s *CreateSubInstitutionRequestSubInstCertifyInfo) SetCertNo(v string) *CreateSubInstitutionRequestSubInstCertifyInfo {
	s.CertNo = &v
	return s
}

func (s *CreateSubInstitutionRequestSubInstCertifyInfo) SetCertType(v string) *CreateSubInstitutionRequestSubInstCertifyInfo {
	s.CertType = &v
	return s
}

type CreateSubInstitutionRequestSubInstInvoiceInfo struct {
	// example:
	//
	// true
	AcceptElectronic *bool `json:"acceptElectronic,omitempty" xml:"acceptElectronic,omitempty"`
	// example:
	//
	// 浙江省杭州市西湖区西溪路蚂蚁金服
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// false
	AutoInvoice *bool `json:"autoInvoice,omitempty" xml:"autoInvoice,omitempty"`
	// example:
	//
	// 1234567812345678123
	BankAccount *string `json:"bankAccount,omitempty" xml:"bankAccount,omitempty"`
	// example:
	//
	// 中国银行
	BankName    *string                                                   `json:"bankName,omitempty" xml:"bankName,omitempty"`
	MailAddress *CreateSubInstitutionRequestSubInstInvoiceInfoMailAddress `json:"mailAddress,omitempty" xml:"mailAddress,omitempty" type:"Struct"`
	// example:
	//
	// 张三
	MailName *string `json:"mailName,omitempty" xml:"mailName,omitempty"`
	// example:
	//
	// 057162288888
	MailPhone *string `json:"mailPhone,omitempty" xml:"mailPhone,omitempty"`
	// example:
	//
	// 51010482542598631219
	TaxNo *string `json:"taxNo,omitempty" xml:"taxNo,omitempty"`
	// example:
	//
	// 01
	TaxPayerQualification *string `json:"taxPayerQualification,omitempty" xml:"taxPayerQualification,omitempty"`
	// example:
	//
	// 19981011
	TaxPayerValidDate *string `json:"taxPayerValidDate,omitempty" xml:"taxPayerValidDate,omitempty"`
	// example:
	//
	// 057162288888
	Telephone *string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	// example:
	//
	// **有限公司
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateSubInstitutionRequestSubInstInvoiceInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateSubInstitutionRequestSubInstInvoiceInfo) GoString() string {
	return s.String()
}

func (s *CreateSubInstitutionRequestSubInstInvoiceInfo) SetAcceptElectronic(v bool) *CreateSubInstitutionRequestSubInstInvoiceInfo {
	s.AcceptElectronic = &v
	return s
}

func (s *CreateSubInstitutionRequestSubInstInvoiceInfo) SetAddress(v string) *CreateSubInstitutionRequestSubInstInvoiceInfo {
	s.Address = &v
	return s
}

func (s *CreateSubInstitutionRequestSubInstInvoiceInfo) SetAutoInvoice(v bool) *CreateSubInstitutionRequestSubInstInvoiceInfo {
	s.AutoInvoice = &v
	return s
}

func (s *CreateSubInstitutionRequestSubInstInvoiceInfo) SetBankAccount(v string) *CreateSubInstitutionRequestSubInstInvoiceInfo {
	s.BankAccount = &v
	return s
}

func (s *CreateSubInstitutionRequestSubInstInvoiceInfo) SetBankName(v string) *CreateSubInstitutionRequestSubInstInvoiceInfo {
	s.BankName = &v
	return s
}

func (s *CreateSubInstitutionRequestSubInstInvoiceInfo) SetMailAddress(v *CreateSubInstitutionRequestSubInstInvoiceInfoMailAddress) *CreateSubInstitutionRequestSubInstInvoiceInfo {
	s.MailAddress = v
	return s
}

func (s *CreateSubInstitutionRequestSubInstInvoiceInfo) SetMailName(v string) *CreateSubInstitutionRequestSubInstInvoiceInfo {
	s.MailName = &v
	return s
}

func (s *CreateSubInstitutionRequestSubInstInvoiceInfo) SetMailPhone(v string) *CreateSubInstitutionRequestSubInstInvoiceInfo {
	s.MailPhone = &v
	return s
}

func (s *CreateSubInstitutionRequestSubInstInvoiceInfo) SetTaxNo(v string) *CreateSubInstitutionRequestSubInstInvoiceInfo {
	s.TaxNo = &v
	return s
}

func (s *CreateSubInstitutionRequestSubInstInvoiceInfo) SetTaxPayerQualification(v string) *CreateSubInstitutionRequestSubInstInvoiceInfo {
	s.TaxPayerQualification = &v
	return s
}

func (s *CreateSubInstitutionRequestSubInstInvoiceInfo) SetTaxPayerValidDate(v string) *CreateSubInstitutionRequestSubInstInvoiceInfo {
	s.TaxPayerValidDate = &v
	return s
}

func (s *CreateSubInstitutionRequestSubInstInvoiceInfo) SetTelephone(v string) *CreateSubInstitutionRequestSubInstInvoiceInfo {
	s.Telephone = &v
	return s
}

func (s *CreateSubInstitutionRequestSubInstInvoiceInfo) SetTitle(v string) *CreateSubInstitutionRequestSubInstInvoiceInfo {
	s.Title = &v
	return s
}

type CreateSubInstitutionRequestSubInstInvoiceInfoMailAddress struct {
	// example:
	//
	// 未来park
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// 330100
	CityCode *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	// example:
	//
	// 330104
	DistrictCode *string `json:"districtCode,omitempty" xml:"districtCode,omitempty"`
	// example:
	//
	// 330000
	ProvinceCode *string `json:"provinceCode,omitempty" xml:"provinceCode,omitempty"`
}

func (s CreateSubInstitutionRequestSubInstInvoiceInfoMailAddress) String() string {
	return tea.Prettify(s)
}

func (s CreateSubInstitutionRequestSubInstInvoiceInfoMailAddress) GoString() string {
	return s.String()
}

func (s *CreateSubInstitutionRequestSubInstInvoiceInfoMailAddress) SetAddress(v string) *CreateSubInstitutionRequestSubInstInvoiceInfoMailAddress {
	s.Address = &v
	return s
}

func (s *CreateSubInstitutionRequestSubInstInvoiceInfoMailAddress) SetCityCode(v string) *CreateSubInstitutionRequestSubInstInvoiceInfoMailAddress {
	s.CityCode = &v
	return s
}

func (s *CreateSubInstitutionRequestSubInstInvoiceInfoMailAddress) SetDistrictCode(v string) *CreateSubInstitutionRequestSubInstInvoiceInfoMailAddress {
	s.DistrictCode = &v
	return s
}

func (s *CreateSubInstitutionRequestSubInstInvoiceInfoMailAddress) SetProvinceCode(v string) *CreateSubInstitutionRequestSubInstInvoiceInfoMailAddress {
	s.ProvinceCode = &v
	return s
}

type CreateSubInstitutionRequestSubInstShopInfo struct {
	InDoorImages  []*string `json:"inDoorImages,omitempty" xml:"inDoorImages,omitempty" type:"Repeated"`
	OutDoorImages []*string `json:"outDoorImages,omitempty" xml:"outDoorImages,omitempty" type:"Repeated"`
}

func (s CreateSubInstitutionRequestSubInstShopInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateSubInstitutionRequestSubInstShopInfo) GoString() string {
	return s.String()
}

func (s *CreateSubInstitutionRequestSubInstShopInfo) SetInDoorImages(v []*string) *CreateSubInstitutionRequestSubInstShopInfo {
	s.InDoorImages = v
	return s
}

func (s *CreateSubInstitutionRequestSubInstShopInfo) SetOutDoorImages(v []*string) *CreateSubInstitutionRequestSubInstShopInfo {
	s.OutDoorImages = v
	return s
}

type CreateSubInstitutionResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 202110110000001
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s CreateSubInstitutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSubInstitutionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubInstitutionResponseBody) SetOrderId(v string) *CreateSubInstitutionResponseBody {
	s.OrderId = &v
	return s
}

type CreateSubInstitutionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSubInstitutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSubInstitutionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSubInstitutionResponse) GoString() string {
	return s.String()
}

func (s *CreateSubInstitutionResponse) SetHeaders(v map[string]*string) *CreateSubInstitutionResponse {
	s.Headers = v
	return s
}

func (s *CreateSubInstitutionResponse) SetStatusCode(v int32) *CreateSubInstitutionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSubInstitutionResponse) SetBody(v *CreateSubInstitutionResponseBody) *CreateSubInstitutionResponse {
	s.Body = v
	return s
}

type CreateUserCodeInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateUserCodeInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateUserCodeInstanceHeaders) GoString() string {
	return s.String()
}

func (s *CreateUserCodeInstanceHeaders) SetCommonHeaders(v map[string]*string) *CreateUserCodeInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateUserCodeInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *CreateUserCodeInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateUserCodeInstanceRequest struct {
	// This parameter is required.
	AvailableTimes []*CreateUserCodeInstanceRequestAvailableTimes `json:"availableTimes,omitempty" xml:"availableTimes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// TEST
	CodeIdentity *string `json:"codeIdentity,omitempty" xml:"codeIdentity,omitempty"`
	CodeValue    *string `json:"codeValue,omitempty" xml:"codeValue,omitempty"`
	// example:
	//
	// DING_STATIC
	CodeValueType *string `json:"codeValueType,omitempty" xml:"codeValueType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// corpid1234
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// This parameter is required.
	ExtInfo map[string]interface{} `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	GmtExpired *string `json:"gmtExpired,omitempty" xml:"gmtExpired,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 202102021212
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// INTERNAL_STAFF
	UserCorpRelationType *string `json:"userCorpRelationType,omitempty" xml:"userCorpRelationType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 86-xxxxxx
	UserIdentity *string `json:"userIdentity,omitempty" xml:"userIdentity,omitempty"`
}

func (s CreateUserCodeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserCodeInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateUserCodeInstanceRequest) SetAvailableTimes(v []*CreateUserCodeInstanceRequestAvailableTimes) *CreateUserCodeInstanceRequest {
	s.AvailableTimes = v
	return s
}

func (s *CreateUserCodeInstanceRequest) SetCodeIdentity(v string) *CreateUserCodeInstanceRequest {
	s.CodeIdentity = &v
	return s
}

func (s *CreateUserCodeInstanceRequest) SetCodeValue(v string) *CreateUserCodeInstanceRequest {
	s.CodeValue = &v
	return s
}

func (s *CreateUserCodeInstanceRequest) SetCodeValueType(v string) *CreateUserCodeInstanceRequest {
	s.CodeValueType = &v
	return s
}

func (s *CreateUserCodeInstanceRequest) SetCorpId(v string) *CreateUserCodeInstanceRequest {
	s.CorpId = &v
	return s
}

func (s *CreateUserCodeInstanceRequest) SetExtInfo(v map[string]interface{}) *CreateUserCodeInstanceRequest {
	s.ExtInfo = v
	return s
}

func (s *CreateUserCodeInstanceRequest) SetGmtExpired(v string) *CreateUserCodeInstanceRequest {
	s.GmtExpired = &v
	return s
}

func (s *CreateUserCodeInstanceRequest) SetRequestId(v string) *CreateUserCodeInstanceRequest {
	s.RequestId = &v
	return s
}

func (s *CreateUserCodeInstanceRequest) SetStatus(v string) *CreateUserCodeInstanceRequest {
	s.Status = &v
	return s
}

func (s *CreateUserCodeInstanceRequest) SetUserCorpRelationType(v string) *CreateUserCodeInstanceRequest {
	s.UserCorpRelationType = &v
	return s
}

func (s *CreateUserCodeInstanceRequest) SetUserIdentity(v string) *CreateUserCodeInstanceRequest {
	s.UserIdentity = &v
	return s
}

type CreateUserCodeInstanceRequestAvailableTimes struct {
	// This parameter is required.
	//
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	GmtEnd *string `json:"gmtEnd,omitempty" xml:"gmtEnd,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	GmtStart *string `json:"gmtStart,omitempty" xml:"gmtStart,omitempty"`
}

func (s CreateUserCodeInstanceRequestAvailableTimes) String() string {
	return tea.Prettify(s)
}

func (s CreateUserCodeInstanceRequestAvailableTimes) GoString() string {
	return s.String()
}

func (s *CreateUserCodeInstanceRequestAvailableTimes) SetGmtEnd(v string) *CreateUserCodeInstanceRequestAvailableTimes {
	s.GmtEnd = &v
	return s
}

func (s *CreateUserCodeInstanceRequestAvailableTimes) SetGmtStart(v string) *CreateUserCodeInstanceRequestAvailableTimes {
	s.GmtStart = &v
	return s
}

type CreateUserCodeInstanceResponseBody struct {
	CodeDetailUrl *string `json:"codeDetailUrl,omitempty" xml:"codeDetailUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// codexxxxxx
	CodeId *string `json:"codeId,omitempty" xml:"codeId,omitempty"`
}

func (s CreateUserCodeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserCodeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserCodeInstanceResponseBody) SetCodeDetailUrl(v string) *CreateUserCodeInstanceResponseBody {
	s.CodeDetailUrl = &v
	return s
}

func (s *CreateUserCodeInstanceResponseBody) SetCodeId(v string) *CreateUserCodeInstanceResponseBody {
	s.CodeId = &v
	return s
}

type CreateUserCodeInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserCodeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserCodeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserCodeInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateUserCodeInstanceResponse) SetHeaders(v map[string]*string) *CreateUserCodeInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateUserCodeInstanceResponse) SetStatusCode(v int32) *CreateUserCodeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserCodeInstanceResponse) SetBody(v *CreateUserCodeInstanceResponseBody) *CreateUserCodeInstanceResponse {
	s.Body = v
	return s
}

type DecodePayCodeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DecodePayCodeHeaders) String() string {
	return tea.Prettify(s)
}

func (s DecodePayCodeHeaders) GoString() string {
	return s.String()
}

func (s *DecodePayCodeHeaders) SetCommonHeaders(v map[string]*string) *DecodePayCodeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DecodePayCodeHeaders) SetXAcsDingtalkAccessToken(v string) *DecodePayCodeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DecodePayCodeRequest struct {
	PayCode   *string `json:"payCode,omitempty" xml:"payCode,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DecodePayCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s DecodePayCodeRequest) GoString() string {
	return s.String()
}

func (s *DecodePayCodeRequest) SetPayCode(v string) *DecodePayCodeRequest {
	s.PayCode = &v
	return s
}

func (s *DecodePayCodeRequest) SetRequestId(v string) *DecodePayCodeRequest {
	s.RequestId = &v
	return s
}

type DecodePayCodeResponseBody struct {
	// example:
	//
	// 2512345678
	AlipayCode *string `json:"alipayCode,omitempty" xml:"alipayCode,omitempty"`
	// example:
	//
	// codeIdxxxxx
	CodeId *string `json:"codeId,omitempty" xml:"codeId,omitempty"`
	// example:
	//
	// DT_VISITOR
	CodeIdentity *string `json:"codeIdentity,omitempty" xml:"codeIdentity,omitempty"`
	// example:
	//
	// PURE_IDENTIFY_CODE
	CodeType *string `json:"codeType,omitempty" xml:"codeType,omitempty"`
	// example:
	//
	// ding1234
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// {"authRules":{}}
	ExtInfo *string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// example:
	//
	// xxxxx
	OutBizId *string `json:"outBizId,omitempty" xml:"outBizId,omitempty"`
	// example:
	//
	// INTERNAL_STAFF
	UserCorpRelationType *string `json:"userCorpRelationType,omitempty" xml:"userCorpRelationType,omitempty"`
	// example:
	//
	// staffId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// true
	UserInCorp *bool `json:"userInCorp,omitempty" xml:"userInCorp,omitempty"`
}

func (s DecodePayCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DecodePayCodeResponseBody) GoString() string {
	return s.String()
}

func (s *DecodePayCodeResponseBody) SetAlipayCode(v string) *DecodePayCodeResponseBody {
	s.AlipayCode = &v
	return s
}

func (s *DecodePayCodeResponseBody) SetCodeId(v string) *DecodePayCodeResponseBody {
	s.CodeId = &v
	return s
}

func (s *DecodePayCodeResponseBody) SetCodeIdentity(v string) *DecodePayCodeResponseBody {
	s.CodeIdentity = &v
	return s
}

func (s *DecodePayCodeResponseBody) SetCodeType(v string) *DecodePayCodeResponseBody {
	s.CodeType = &v
	return s
}

func (s *DecodePayCodeResponseBody) SetCorpId(v string) *DecodePayCodeResponseBody {
	s.CorpId = &v
	return s
}

func (s *DecodePayCodeResponseBody) SetExtInfo(v string) *DecodePayCodeResponseBody {
	s.ExtInfo = &v
	return s
}

func (s *DecodePayCodeResponseBody) SetOutBizId(v string) *DecodePayCodeResponseBody {
	s.OutBizId = &v
	return s
}

func (s *DecodePayCodeResponseBody) SetUserCorpRelationType(v string) *DecodePayCodeResponseBody {
	s.UserCorpRelationType = &v
	return s
}

func (s *DecodePayCodeResponseBody) SetUserId(v string) *DecodePayCodeResponseBody {
	s.UserId = &v
	return s
}

func (s *DecodePayCodeResponseBody) SetUserInCorp(v bool) *DecodePayCodeResponseBody {
	s.UserInCorp = &v
	return s
}

type DecodePayCodeResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DecodePayCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DecodePayCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DecodePayCodeResponse) GoString() string {
	return s.String()
}

func (s *DecodePayCodeResponse) SetHeaders(v map[string]*string) *DecodePayCodeResponse {
	s.Headers = v
	return s
}

func (s *DecodePayCodeResponse) SetStatusCode(v int32) *DecodePayCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DecodePayCodeResponse) SetBody(v *DecodePayCodeResponseBody) *DecodePayCodeResponse {
	s.Body = v
	return s
}

type FinanceLoanNotifyRegisterHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s FinanceLoanNotifyRegisterHeaders) String() string {
	return tea.Prettify(s)
}

func (s FinanceLoanNotifyRegisterHeaders) GoString() string {
	return s.String()
}

func (s *FinanceLoanNotifyRegisterHeaders) SetCommonHeaders(v map[string]*string) *FinanceLoanNotifyRegisterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FinanceLoanNotifyRegisterHeaders) SetXAcsDingtalkAccessToken(v string) *FinanceLoanNotifyRegisterHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type FinanceLoanNotifyRegisterRequest struct {
	// example:
	//
	// 2024-06-18 14:53:33
	CompleteTime *string `json:"completeTime,omitempty" xml:"completeTime,omitempty"`
	// example:
	//
	// {}
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// example:
	//
	// 330725189509101234
	IdCardNo *string `json:"idCardNo,omitempty" xml:"idCardNo,omitempty"`
	// example:
	//
	// 中原消费金融
	OpenChannelName *string `json:"openChannelName,omitempty" xml:"openChannelName,omitempty"`
	// example:
	//
	// XFD201909210001
	OpenProductCode *string `json:"openProductCode,omitempty" xml:"openProductCode,omitempty"`
	// example:
	//
	// 员工贷
	OpenProductName *string `json:"openProductName,omitempty" xml:"openProductName,omitempty"`
	// example:
	//
	// ZYXJ_XFD
	OpenProductType *string `json:"openProductType,omitempty" xml:"openProductType,omitempty"`
	// example:
	//
	// 0
	ProcessingStatus *string `json:"processingStatus,omitempty" xml:"processingStatus,omitempty"`
	// example:
	//
	// ZRSB2020
	RefuseCode *string `json:"refuseCode,omitempty" xml:"refuseCode,omitempty"`
	// example:
	//
	// 进件准入失败
	RefuseReason *string `json:"refuseReason,omitempty" xml:"refuseReason,omitempty"`
	// example:
	//
	// 2024061814654041710801
	RegisterNo *string `json:"registerNo,omitempty" xml:"registerNo,omitempty"`
	// example:
	//
	// 0
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 2024-06-18 14:53:33
	SubmitTime *string `json:"submitTime,omitempty" xml:"submitTime,omitempty"`
	// example:
	//
	// 18092149430
	UserMobile *string `json:"userMobile,omitempty" xml:"userMobile,omitempty"`
}

func (s FinanceLoanNotifyRegisterRequest) String() string {
	return tea.Prettify(s)
}

func (s FinanceLoanNotifyRegisterRequest) GoString() string {
	return s.String()
}

func (s *FinanceLoanNotifyRegisterRequest) SetCompleteTime(v string) *FinanceLoanNotifyRegisterRequest {
	s.CompleteTime = &v
	return s
}

func (s *FinanceLoanNotifyRegisterRequest) SetExtension(v string) *FinanceLoanNotifyRegisterRequest {
	s.Extension = &v
	return s
}

func (s *FinanceLoanNotifyRegisterRequest) SetIdCardNo(v string) *FinanceLoanNotifyRegisterRequest {
	s.IdCardNo = &v
	return s
}

func (s *FinanceLoanNotifyRegisterRequest) SetOpenChannelName(v string) *FinanceLoanNotifyRegisterRequest {
	s.OpenChannelName = &v
	return s
}

func (s *FinanceLoanNotifyRegisterRequest) SetOpenProductCode(v string) *FinanceLoanNotifyRegisterRequest {
	s.OpenProductCode = &v
	return s
}

func (s *FinanceLoanNotifyRegisterRequest) SetOpenProductName(v string) *FinanceLoanNotifyRegisterRequest {
	s.OpenProductName = &v
	return s
}

func (s *FinanceLoanNotifyRegisterRequest) SetOpenProductType(v string) *FinanceLoanNotifyRegisterRequest {
	s.OpenProductType = &v
	return s
}

func (s *FinanceLoanNotifyRegisterRequest) SetProcessingStatus(v string) *FinanceLoanNotifyRegisterRequest {
	s.ProcessingStatus = &v
	return s
}

func (s *FinanceLoanNotifyRegisterRequest) SetRefuseCode(v string) *FinanceLoanNotifyRegisterRequest {
	s.RefuseCode = &v
	return s
}

func (s *FinanceLoanNotifyRegisterRequest) SetRefuseReason(v string) *FinanceLoanNotifyRegisterRequest {
	s.RefuseReason = &v
	return s
}

func (s *FinanceLoanNotifyRegisterRequest) SetRegisterNo(v string) *FinanceLoanNotifyRegisterRequest {
	s.RegisterNo = &v
	return s
}

func (s *FinanceLoanNotifyRegisterRequest) SetStatus(v string) *FinanceLoanNotifyRegisterRequest {
	s.Status = &v
	return s
}

func (s *FinanceLoanNotifyRegisterRequest) SetSubmitTime(v string) *FinanceLoanNotifyRegisterRequest {
	s.SubmitTime = &v
	return s
}

func (s *FinanceLoanNotifyRegisterRequest) SetUserMobile(v string) *FinanceLoanNotifyRegisterRequest {
	s.UserMobile = &v
	return s
}

type FinanceLoanNotifyRegisterResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s FinanceLoanNotifyRegisterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FinanceLoanNotifyRegisterResponseBody) GoString() string {
	return s.String()
}

func (s *FinanceLoanNotifyRegisterResponseBody) SetRequestId(v string) *FinanceLoanNotifyRegisterResponseBody {
	s.RequestId = &v
	return s
}

func (s *FinanceLoanNotifyRegisterResponseBody) SetResult(v string) *FinanceLoanNotifyRegisterResponseBody {
	s.Result = &v
	return s
}

type FinanceLoanNotifyRegisterResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FinanceLoanNotifyRegisterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FinanceLoanNotifyRegisterResponse) String() string {
	return tea.Prettify(s)
}

func (s FinanceLoanNotifyRegisterResponse) GoString() string {
	return s.String()
}

func (s *FinanceLoanNotifyRegisterResponse) SetHeaders(v map[string]*string) *FinanceLoanNotifyRegisterResponse {
	s.Headers = v
	return s
}

func (s *FinanceLoanNotifyRegisterResponse) SetStatusCode(v int32) *FinanceLoanNotifyRegisterResponse {
	s.StatusCode = &v
	return s
}

func (s *FinanceLoanNotifyRegisterResponse) SetBody(v *FinanceLoanNotifyRegisterResponseBody) *FinanceLoanNotifyRegisterResponse {
	s.Body = v
	return s
}

type ModifySubInstitutionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ModifySubInstitutionHeaders) String() string {
	return tea.Prettify(s)
}

func (s ModifySubInstitutionHeaders) GoString() string {
	return s.String()
}

func (s *ModifySubInstitutionHeaders) SetCommonHeaders(v map[string]*string) *ModifySubInstitutionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ModifySubInstitutionHeaders) SetXAcsDingtalkAccessToken(v string) *ModifySubInstitutionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ModifySubInstitutionRequest struct {
	// example:
	//
	// asdf@163.com
	BindingAlipayLogonId *string `json:"bindingAlipayLogonId,omitempty" xml:"bindingAlipayLogonId,omitempty"`
	// This parameter is required.
	ContactInfo *ModifySubInstitutionRequestContactInfo `json:"contactInfo,omitempty" xml:"contactInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 202111090001
	InstId *string `json:"instId,omitempty" xml:"instId,omitempty"`
	// This parameter is required.
	LegalPersonCertInfo *ModifySubInstitutionRequestLegalPersonCertInfo `json:"legalPersonCertInfo,omitempty" xml:"legalPersonCertInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 2021000001
	OutTradeNo *string `json:"outTradeNo,omitempty" xml:"outTradeNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALIPAY
	PayChannel         *string                                          `json:"payChannel,omitempty" xml:"payChannel,omitempty"`
	QualificationInfos []*ModifySubInstitutionRequestQualificationInfos `json:"qualificationInfos,omitempty" xml:"qualificationInfos,omitempty" type:"Repeated"`
	// This parameter is required.
	Services []*string `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
	// This parameter is required.
	SettleInfo         *ModifySubInstitutionRequestSettleInfo         `json:"settleInfo,omitempty" xml:"settleInfo,omitempty" type:"Struct"`
	SubInstAddressInfo *ModifySubInstitutionRequestSubInstAddressInfo `json:"subInstAddressInfo,omitempty" xml:"subInstAddressInfo,omitempty" type:"Struct"`
	SubInstAuthInfo    *ModifySubInstitutionRequestSubInstAuthInfo    `json:"subInstAuthInfo,omitempty" xml:"subInstAuthInfo,omitempty" type:"Struct"`
	// This parameter is required.
	SubInstBasicInfo *ModifySubInstitutionRequestSubInstBasicInfo `json:"subInstBasicInfo,omitempty" xml:"subInstBasicInfo,omitempty" type:"Struct"`
	// This parameter is required.
	SubInstCertifyInfo *ModifySubInstitutionRequestSubInstCertifyInfo `json:"subInstCertifyInfo,omitempty" xml:"subInstCertifyInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 1001
	SubInstId          *string                                        `json:"subInstId,omitempty" xml:"subInstId,omitempty"`
	SubInstInvoiceInfo *ModifySubInstitutionRequestSubInstInvoiceInfo `json:"subInstInvoiceInfo,omitempty" xml:"subInstInvoiceInfo,omitempty" type:"Struct"`
	SubInstShopInfo    *ModifySubInstitutionRequestSubInstShopInfo    `json:"subInstShopInfo,omitempty" xml:"subInstShopInfo,omitempty" type:"Struct"`
}

func (s ModifySubInstitutionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySubInstitutionRequest) GoString() string {
	return s.String()
}

func (s *ModifySubInstitutionRequest) SetBindingAlipayLogonId(v string) *ModifySubInstitutionRequest {
	s.BindingAlipayLogonId = &v
	return s
}

func (s *ModifySubInstitutionRequest) SetContactInfo(v *ModifySubInstitutionRequestContactInfo) *ModifySubInstitutionRequest {
	s.ContactInfo = v
	return s
}

func (s *ModifySubInstitutionRequest) SetInstId(v string) *ModifySubInstitutionRequest {
	s.InstId = &v
	return s
}

func (s *ModifySubInstitutionRequest) SetLegalPersonCertInfo(v *ModifySubInstitutionRequestLegalPersonCertInfo) *ModifySubInstitutionRequest {
	s.LegalPersonCertInfo = v
	return s
}

func (s *ModifySubInstitutionRequest) SetOutTradeNo(v string) *ModifySubInstitutionRequest {
	s.OutTradeNo = &v
	return s
}

func (s *ModifySubInstitutionRequest) SetPayChannel(v string) *ModifySubInstitutionRequest {
	s.PayChannel = &v
	return s
}

func (s *ModifySubInstitutionRequest) SetQualificationInfos(v []*ModifySubInstitutionRequestQualificationInfos) *ModifySubInstitutionRequest {
	s.QualificationInfos = v
	return s
}

func (s *ModifySubInstitutionRequest) SetServices(v []*string) *ModifySubInstitutionRequest {
	s.Services = v
	return s
}

func (s *ModifySubInstitutionRequest) SetSettleInfo(v *ModifySubInstitutionRequestSettleInfo) *ModifySubInstitutionRequest {
	s.SettleInfo = v
	return s
}

func (s *ModifySubInstitutionRequest) SetSubInstAddressInfo(v *ModifySubInstitutionRequestSubInstAddressInfo) *ModifySubInstitutionRequest {
	s.SubInstAddressInfo = v
	return s
}

func (s *ModifySubInstitutionRequest) SetSubInstAuthInfo(v *ModifySubInstitutionRequestSubInstAuthInfo) *ModifySubInstitutionRequest {
	s.SubInstAuthInfo = v
	return s
}

func (s *ModifySubInstitutionRequest) SetSubInstBasicInfo(v *ModifySubInstitutionRequestSubInstBasicInfo) *ModifySubInstitutionRequest {
	s.SubInstBasicInfo = v
	return s
}

func (s *ModifySubInstitutionRequest) SetSubInstCertifyInfo(v *ModifySubInstitutionRequestSubInstCertifyInfo) *ModifySubInstitutionRequest {
	s.SubInstCertifyInfo = v
	return s
}

func (s *ModifySubInstitutionRequest) SetSubInstId(v string) *ModifySubInstitutionRequest {
	s.SubInstId = &v
	return s
}

func (s *ModifySubInstitutionRequest) SetSubInstInvoiceInfo(v *ModifySubInstitutionRequestSubInstInvoiceInfo) *ModifySubInstitutionRequest {
	s.SubInstInvoiceInfo = v
	return s
}

func (s *ModifySubInstitutionRequest) SetSubInstShopInfo(v *ModifySubInstitutionRequestSubInstShopInfo) *ModifySubInstitutionRequest {
	s.SubInstShopInfo = v
	return s
}

type ModifySubInstitutionRequestContactInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 李某某
	ContactName *string `json:"contactName,omitempty" xml:"contactName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13900000000
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
}

func (s ModifySubInstitutionRequestContactInfo) String() string {
	return tea.Prettify(s)
}

func (s ModifySubInstitutionRequestContactInfo) GoString() string {
	return s.String()
}

func (s *ModifySubInstitutionRequestContactInfo) SetContactName(v string) *ModifySubInstitutionRequestContactInfo {
	s.ContactName = &v
	return s
}

func (s *ModifySubInstitutionRequestContactInfo) SetMobile(v string) *ModifySubInstitutionRequestContactInfo {
	s.Mobile = &v
	return s
}

type ModifySubInstitutionRequestLegalPersonCertInfo struct {
	// example:
	//
	// ossUrl
	CertBackImage *string `json:"certBackImage,omitempty" xml:"certBackImage,omitempty"`
	// example:
	//
	// ossUrl
	CertFrontImage *string `json:"certFrontImage,omitempty" xml:"certFrontImage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 李某某
	CertName *string `json:"certName,omitempty" xml:"certName,omitempty"`
	// example:
	//
	// 100
	CertType *string `json:"certType,omitempty" xml:"certType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 330104200010109999
	IdCardNo *string `json:"idCardNo,omitempty" xml:"idCardNo,omitempty"`
}

func (s ModifySubInstitutionRequestLegalPersonCertInfo) String() string {
	return tea.Prettify(s)
}

func (s ModifySubInstitutionRequestLegalPersonCertInfo) GoString() string {
	return s.String()
}

func (s *ModifySubInstitutionRequestLegalPersonCertInfo) SetCertBackImage(v string) *ModifySubInstitutionRequestLegalPersonCertInfo {
	s.CertBackImage = &v
	return s
}

func (s *ModifySubInstitutionRequestLegalPersonCertInfo) SetCertFrontImage(v string) *ModifySubInstitutionRequestLegalPersonCertInfo {
	s.CertFrontImage = &v
	return s
}

func (s *ModifySubInstitutionRequestLegalPersonCertInfo) SetCertName(v string) *ModifySubInstitutionRequestLegalPersonCertInfo {
	s.CertName = &v
	return s
}

func (s *ModifySubInstitutionRequestLegalPersonCertInfo) SetCertType(v string) *ModifySubInstitutionRequestLegalPersonCertInfo {
	s.CertType = &v
	return s
}

func (s *ModifySubInstitutionRequestLegalPersonCertInfo) SetIdCardNo(v string) *ModifySubInstitutionRequestLegalPersonCertInfo {
	s.IdCardNo = &v
	return s
}

type ModifySubInstitutionRequestQualificationInfos struct {
	// This parameter is required.
	//
	// example:
	//
	// ossUrl
	QualificationImage *string `json:"qualificationImage,omitempty" xml:"qualificationImage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 301
	QualificationType *string `json:"qualificationType,omitempty" xml:"qualificationType,omitempty"`
}

func (s ModifySubInstitutionRequestQualificationInfos) String() string {
	return tea.Prettify(s)
}

func (s ModifySubInstitutionRequestQualificationInfos) GoString() string {
	return s.String()
}

func (s *ModifySubInstitutionRequestQualificationInfos) SetQualificationImage(v string) *ModifySubInstitutionRequestQualificationInfos {
	s.QualificationImage = &v
	return s
}

func (s *ModifySubInstitutionRequestQualificationInfos) SetQualificationType(v string) *ModifySubInstitutionRequestQualificationInfos {
	s.QualificationType = &v
	return s
}

type ModifySubInstitutionRequestSettleInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 622202120200000000
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// 李某某
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	// example:
	//
	// DEBIT_CARD
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	// example:
	//
	// 城东支行
	BankBranchName *string `json:"bankBranchName,omitempty" xml:"bankBranchName,omitempty"`
	// example:
	//
	// 杭州市
	BankCity *string `json:"bankCity,omitempty" xml:"bankCity,omitempty"`
	// example:
	//
	// 313791000023
	BankCode *string `json:"bankCode,omitempty" xml:"bankCode,omitempty"`
	// example:
	//
	// 工商银行
	BankName *string `json:"bankName,omitempty" xml:"bankName,omitempty"`
	// example:
	//
	// 浙江省
	BankProvince *string `json:"bankProvince,omitempty" xml:"bankProvince,omitempty"`
	// example:
	//
	// ICBC
	BankShortNameCode *string `json:"bankShortNameCode,omitempty" xml:"bankShortNameCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALIPAY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// TO_PRI
	UsageType *string `json:"usageType,omitempty" xml:"usageType,omitempty"`
}

func (s ModifySubInstitutionRequestSettleInfo) String() string {
	return tea.Prettify(s)
}

func (s ModifySubInstitutionRequestSettleInfo) GoString() string {
	return s.String()
}

func (s *ModifySubInstitutionRequestSettleInfo) SetAccountId(v string) *ModifySubInstitutionRequestSettleInfo {
	s.AccountId = &v
	return s
}

func (s *ModifySubInstitutionRequestSettleInfo) SetAccountName(v string) *ModifySubInstitutionRequestSettleInfo {
	s.AccountName = &v
	return s
}

func (s *ModifySubInstitutionRequestSettleInfo) SetAccountType(v string) *ModifySubInstitutionRequestSettleInfo {
	s.AccountType = &v
	return s
}

func (s *ModifySubInstitutionRequestSettleInfo) SetBankBranchName(v string) *ModifySubInstitutionRequestSettleInfo {
	s.BankBranchName = &v
	return s
}

func (s *ModifySubInstitutionRequestSettleInfo) SetBankCity(v string) *ModifySubInstitutionRequestSettleInfo {
	s.BankCity = &v
	return s
}

func (s *ModifySubInstitutionRequestSettleInfo) SetBankCode(v string) *ModifySubInstitutionRequestSettleInfo {
	s.BankCode = &v
	return s
}

func (s *ModifySubInstitutionRequestSettleInfo) SetBankName(v string) *ModifySubInstitutionRequestSettleInfo {
	s.BankName = &v
	return s
}

func (s *ModifySubInstitutionRequestSettleInfo) SetBankProvince(v string) *ModifySubInstitutionRequestSettleInfo {
	s.BankProvince = &v
	return s
}

func (s *ModifySubInstitutionRequestSettleInfo) SetBankShortNameCode(v string) *ModifySubInstitutionRequestSettleInfo {
	s.BankShortNameCode = &v
	return s
}

func (s *ModifySubInstitutionRequestSettleInfo) SetType(v string) *ModifySubInstitutionRequestSettleInfo {
	s.Type = &v
	return s
}

func (s *ModifySubInstitutionRequestSettleInfo) SetUsageType(v string) *ModifySubInstitutionRequestSettleInfo {
	s.UsageType = &v
	return s
}

type ModifySubInstitutionRequestSubInstAddressInfo struct {
	// example:
	//
	// 未来park
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// 330100
	CityCode *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	// example:
	//
	// 330104
	DistrictCode *string `json:"districtCode,omitempty" xml:"districtCode,omitempty"`
	// example:
	//
	// 330000
	ProvinceCode *string `json:"provinceCode,omitempty" xml:"provinceCode,omitempty"`
}

func (s ModifySubInstitutionRequestSubInstAddressInfo) String() string {
	return tea.Prettify(s)
}

func (s ModifySubInstitutionRequestSubInstAddressInfo) GoString() string {
	return s.String()
}

func (s *ModifySubInstitutionRequestSubInstAddressInfo) SetAddress(v string) *ModifySubInstitutionRequestSubInstAddressInfo {
	s.Address = &v
	return s
}

func (s *ModifySubInstitutionRequestSubInstAddressInfo) SetCityCode(v string) *ModifySubInstitutionRequestSubInstAddressInfo {
	s.CityCode = &v
	return s
}

func (s *ModifySubInstitutionRequestSubInstAddressInfo) SetDistrictCode(v string) *ModifySubInstitutionRequestSubInstAddressInfo {
	s.DistrictCode = &v
	return s
}

func (s *ModifySubInstitutionRequestSubInstAddressInfo) SetProvinceCode(v string) *ModifySubInstitutionRequestSubInstAddressInfo {
	s.ProvinceCode = &v
	return s
}

type ModifySubInstitutionRequestSubInstAuthInfo struct {
	// example:
	//
	// ossUrl
	AuthorizationLetterUrl *string `json:"authorizationLetterUrl,omitempty" xml:"authorizationLetterUrl,omitempty"`
}

func (s ModifySubInstitutionRequestSubInstAuthInfo) String() string {
	return tea.Prettify(s)
}

func (s ModifySubInstitutionRequestSubInstAuthInfo) GoString() string {
	return s.String()
}

func (s *ModifySubInstitutionRequestSubInstAuthInfo) SetAuthorizationLetterUrl(v string) *ModifySubInstitutionRequestSubInstAuthInfo {
	s.AuthorizationLetterUrl = &v
	return s
}

type ModifySubInstitutionRequestSubInstBasicInfo struct {
	// example:
	//
	// 一食堂
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5812
	Mcc *string `json:"mcc,omitempty" xml:"mcc,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 食堂
	SubInstName *string `json:"subInstName,omitempty" xml:"subInstName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 01
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModifySubInstitutionRequestSubInstBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s ModifySubInstitutionRequestSubInstBasicInfo) GoString() string {
	return s.String()
}

func (s *ModifySubInstitutionRequestSubInstBasicInfo) SetAliasName(v string) *ModifySubInstitutionRequestSubInstBasicInfo {
	s.AliasName = &v
	return s
}

func (s *ModifySubInstitutionRequestSubInstBasicInfo) SetMcc(v string) *ModifySubInstitutionRequestSubInstBasicInfo {
	s.Mcc = &v
	return s
}

func (s *ModifySubInstitutionRequestSubInstBasicInfo) SetSubInstName(v string) *ModifySubInstitutionRequestSubInstBasicInfo {
	s.SubInstName = &v
	return s
}

func (s *ModifySubInstitutionRequestSubInstBasicInfo) SetType(v string) *ModifySubInstitutionRequestSubInstBasicInfo {
	s.Type = &v
	return s
}

type ModifySubInstitutionRequestSubInstCertifyInfo struct {
	// example:
	//
	// ossUrl
	CertImage *string `json:"certImage,omitempty" xml:"certImage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 331081198611111111
	CertNo *string `json:"certNo,omitempty" xml:"certNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 201
	CertType *string `json:"certType,omitempty" xml:"certType,omitempty"`
}

func (s ModifySubInstitutionRequestSubInstCertifyInfo) String() string {
	return tea.Prettify(s)
}

func (s ModifySubInstitutionRequestSubInstCertifyInfo) GoString() string {
	return s.String()
}

func (s *ModifySubInstitutionRequestSubInstCertifyInfo) SetCertImage(v string) *ModifySubInstitutionRequestSubInstCertifyInfo {
	s.CertImage = &v
	return s
}

func (s *ModifySubInstitutionRequestSubInstCertifyInfo) SetCertNo(v string) *ModifySubInstitutionRequestSubInstCertifyInfo {
	s.CertNo = &v
	return s
}

func (s *ModifySubInstitutionRequestSubInstCertifyInfo) SetCertType(v string) *ModifySubInstitutionRequestSubInstCertifyInfo {
	s.CertType = &v
	return s
}

type ModifySubInstitutionRequestSubInstInvoiceInfo struct {
	// example:
	//
	// true
	AcceptElectronic *bool `json:"acceptElectronic,omitempty" xml:"acceptElectronic,omitempty"`
	// example:
	//
	// 浙江省杭州市西湖区西溪路蚂蚁金服
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// false
	AutoInvoice *bool `json:"autoInvoice,omitempty" xml:"autoInvoice,omitempty"`
	// example:
	//
	// 1234567812345678123
	BankAccount *string `json:"bankAccount,omitempty" xml:"bankAccount,omitempty"`
	// example:
	//
	// 中国银行
	BankName    *string                                                   `json:"bankName,omitempty" xml:"bankName,omitempty"`
	MailAddress *ModifySubInstitutionRequestSubInstInvoiceInfoMailAddress `json:"mailAddress,omitempty" xml:"mailAddress,omitempty" type:"Struct"`
	// example:
	//
	// 张三
	MailName *string `json:"mailName,omitempty" xml:"mailName,omitempty"`
	// example:
	//
	// 057162288888
	MailPhone *string `json:"mailPhone,omitempty" xml:"mailPhone,omitempty"`
	// example:
	//
	// 51010482542598631219
	TaxNo *string `json:"taxNo,omitempty" xml:"taxNo,omitempty"`
	// example:
	//
	// 01
	TaxPayerQualification *string `json:"taxPayerQualification,omitempty" xml:"taxPayerQualification,omitempty"`
	// example:
	//
	// 19981011
	TaxPayerValidDate *string `json:"taxPayerValidDate,omitempty" xml:"taxPayerValidDate,omitempty"`
	// example:
	//
	// 057162288888
	Telephone *string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	// example:
	//
	// **有限公司
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ModifySubInstitutionRequestSubInstInvoiceInfo) String() string {
	return tea.Prettify(s)
}

func (s ModifySubInstitutionRequestSubInstInvoiceInfo) GoString() string {
	return s.String()
}

func (s *ModifySubInstitutionRequestSubInstInvoiceInfo) SetAcceptElectronic(v bool) *ModifySubInstitutionRequestSubInstInvoiceInfo {
	s.AcceptElectronic = &v
	return s
}

func (s *ModifySubInstitutionRequestSubInstInvoiceInfo) SetAddress(v string) *ModifySubInstitutionRequestSubInstInvoiceInfo {
	s.Address = &v
	return s
}

func (s *ModifySubInstitutionRequestSubInstInvoiceInfo) SetAutoInvoice(v bool) *ModifySubInstitutionRequestSubInstInvoiceInfo {
	s.AutoInvoice = &v
	return s
}

func (s *ModifySubInstitutionRequestSubInstInvoiceInfo) SetBankAccount(v string) *ModifySubInstitutionRequestSubInstInvoiceInfo {
	s.BankAccount = &v
	return s
}

func (s *ModifySubInstitutionRequestSubInstInvoiceInfo) SetBankName(v string) *ModifySubInstitutionRequestSubInstInvoiceInfo {
	s.BankName = &v
	return s
}

func (s *ModifySubInstitutionRequestSubInstInvoiceInfo) SetMailAddress(v *ModifySubInstitutionRequestSubInstInvoiceInfoMailAddress) *ModifySubInstitutionRequestSubInstInvoiceInfo {
	s.MailAddress = v
	return s
}

func (s *ModifySubInstitutionRequestSubInstInvoiceInfo) SetMailName(v string) *ModifySubInstitutionRequestSubInstInvoiceInfo {
	s.MailName = &v
	return s
}

func (s *ModifySubInstitutionRequestSubInstInvoiceInfo) SetMailPhone(v string) *ModifySubInstitutionRequestSubInstInvoiceInfo {
	s.MailPhone = &v
	return s
}

func (s *ModifySubInstitutionRequestSubInstInvoiceInfo) SetTaxNo(v string) *ModifySubInstitutionRequestSubInstInvoiceInfo {
	s.TaxNo = &v
	return s
}

func (s *ModifySubInstitutionRequestSubInstInvoiceInfo) SetTaxPayerQualification(v string) *ModifySubInstitutionRequestSubInstInvoiceInfo {
	s.TaxPayerQualification = &v
	return s
}

func (s *ModifySubInstitutionRequestSubInstInvoiceInfo) SetTaxPayerValidDate(v string) *ModifySubInstitutionRequestSubInstInvoiceInfo {
	s.TaxPayerValidDate = &v
	return s
}

func (s *ModifySubInstitutionRequestSubInstInvoiceInfo) SetTelephone(v string) *ModifySubInstitutionRequestSubInstInvoiceInfo {
	s.Telephone = &v
	return s
}

func (s *ModifySubInstitutionRequestSubInstInvoiceInfo) SetTitle(v string) *ModifySubInstitutionRequestSubInstInvoiceInfo {
	s.Title = &v
	return s
}

type ModifySubInstitutionRequestSubInstInvoiceInfoMailAddress struct {
	// example:
	//
	// 未来park
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// 330100
	CityCode *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	// example:
	//
	// 330104
	DistrictCode *string `json:"districtCode,omitempty" xml:"districtCode,omitempty"`
	// example:
	//
	// 330000
	ProvinceCode *string `json:"provinceCode,omitempty" xml:"provinceCode,omitempty"`
}

func (s ModifySubInstitutionRequestSubInstInvoiceInfoMailAddress) String() string {
	return tea.Prettify(s)
}

func (s ModifySubInstitutionRequestSubInstInvoiceInfoMailAddress) GoString() string {
	return s.String()
}

func (s *ModifySubInstitutionRequestSubInstInvoiceInfoMailAddress) SetAddress(v string) *ModifySubInstitutionRequestSubInstInvoiceInfoMailAddress {
	s.Address = &v
	return s
}

func (s *ModifySubInstitutionRequestSubInstInvoiceInfoMailAddress) SetCityCode(v string) *ModifySubInstitutionRequestSubInstInvoiceInfoMailAddress {
	s.CityCode = &v
	return s
}

func (s *ModifySubInstitutionRequestSubInstInvoiceInfoMailAddress) SetDistrictCode(v string) *ModifySubInstitutionRequestSubInstInvoiceInfoMailAddress {
	s.DistrictCode = &v
	return s
}

func (s *ModifySubInstitutionRequestSubInstInvoiceInfoMailAddress) SetProvinceCode(v string) *ModifySubInstitutionRequestSubInstInvoiceInfoMailAddress {
	s.ProvinceCode = &v
	return s
}

type ModifySubInstitutionRequestSubInstShopInfo struct {
	InDoorImages  []*string `json:"inDoorImages,omitempty" xml:"inDoorImages,omitempty" type:"Repeated"`
	OutDoorImages []*string `json:"outDoorImages,omitempty" xml:"outDoorImages,omitempty" type:"Repeated"`
}

func (s ModifySubInstitutionRequestSubInstShopInfo) String() string {
	return tea.Prettify(s)
}

func (s ModifySubInstitutionRequestSubInstShopInfo) GoString() string {
	return s.String()
}

func (s *ModifySubInstitutionRequestSubInstShopInfo) SetInDoorImages(v []*string) *ModifySubInstitutionRequestSubInstShopInfo {
	s.InDoorImages = v
	return s
}

func (s *ModifySubInstitutionRequestSubInstShopInfo) SetOutDoorImages(v []*string) *ModifySubInstitutionRequestSubInstShopInfo {
	s.OutDoorImages = v
	return s
}

type ModifySubInstitutionResponseBody struct {
	// example:
	//
	// 202110110000001
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s ModifySubInstitutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySubInstitutionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySubInstitutionResponseBody) SetOrderId(v string) *ModifySubInstitutionResponseBody {
	s.OrderId = &v
	return s
}

type ModifySubInstitutionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySubInstitutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySubInstitutionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySubInstitutionResponse) GoString() string {
	return s.String()
}

func (s *ModifySubInstitutionResponse) SetHeaders(v map[string]*string) *ModifySubInstitutionResponse {
	s.Headers = v
	return s
}

func (s *ModifySubInstitutionResponse) SetStatusCode(v int32) *ModifySubInstitutionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySubInstitutionResponse) SetBody(v *ModifySubInstitutionResponseBody) *ModifySubInstitutionResponse {
	s.Body = v
	return s
}

type NotifyPayCodePayResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s NotifyPayCodePayResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayCodePayResultHeaders) GoString() string {
	return s.String()
}

func (s *NotifyPayCodePayResultHeaders) SetCommonHeaders(v map[string]*string) *NotifyPayCodePayResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *NotifyPayCodePayResultHeaders) SetXAcsDingtalkAccessToken(v string) *NotifyPayCodePayResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type NotifyPayCodePayResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234.56
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1.00, 没有传0.00
	ChargeAmount *string `json:"chargeAmount,omitempty" xml:"chargeAmount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ding1234
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// { "akey": "avalue“}
	ExtInfo *string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-01-01 11:11:11
	GmtTradeCreate *string `json:"gmtTradeCreate,omitempty" xml:"gmtTradeCreate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-01-01 11:11:11
	GmtTradeFinish *string `json:"gmtTradeFinish,omitempty" xml:"gmtTradeFinish,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// XX公司食堂
	MerchantName *string `json:"merchantName,omitempty" xml:"merchantName,omitempty"`
	// This parameter is required.
	PayChannelDetailList []*NotifyPayCodePayResultRequestPayChannelDetailList `json:"payChannelDetailList,omitempty" xml:"payChannelDetailList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 261234567890
	PayCode *string `json:"payCode,omitempty" xml:"payCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1.23，没有传0.00
	PromotionAmount *string `json:"promotionAmount,omitempty" xml:"promotionAmount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 晚餐100.0元
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// BALANCE_NOT_ENOUGH
	TradeErrorCode *string `json:"tradeErrorCode,omitempty" xml:"tradeErrorCode,omitempty"`
	// example:
	//
	// 余额不足，请充值
	TradeErrorMsg *string `json:"tradeErrorMsg,omitempty" xml:"tradeErrorMsg,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 202101012345678
	TradeNo *string `json:"tradeNo,omitempty" xml:"tradeNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SUCCESS/FAIL
	TradeStatus *string `json:"tradeStatus,omitempty" xml:"tradeStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// userId1234
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s NotifyPayCodePayResultRequest) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayCodePayResultRequest) GoString() string {
	return s.String()
}

func (s *NotifyPayCodePayResultRequest) SetAmount(v string) *NotifyPayCodePayResultRequest {
	s.Amount = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetChargeAmount(v string) *NotifyPayCodePayResultRequest {
	s.ChargeAmount = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetCorpId(v string) *NotifyPayCodePayResultRequest {
	s.CorpId = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetExtInfo(v string) *NotifyPayCodePayResultRequest {
	s.ExtInfo = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetGmtTradeCreate(v string) *NotifyPayCodePayResultRequest {
	s.GmtTradeCreate = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetGmtTradeFinish(v string) *NotifyPayCodePayResultRequest {
	s.GmtTradeFinish = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetMerchantName(v string) *NotifyPayCodePayResultRequest {
	s.MerchantName = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetPayChannelDetailList(v []*NotifyPayCodePayResultRequestPayChannelDetailList) *NotifyPayCodePayResultRequest {
	s.PayChannelDetailList = v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetPayCode(v string) *NotifyPayCodePayResultRequest {
	s.PayCode = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetPromotionAmount(v string) *NotifyPayCodePayResultRequest {
	s.PromotionAmount = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetRemark(v string) *NotifyPayCodePayResultRequest {
	s.Remark = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetTitle(v string) *NotifyPayCodePayResultRequest {
	s.Title = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetTradeErrorCode(v string) *NotifyPayCodePayResultRequest {
	s.TradeErrorCode = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetTradeErrorMsg(v string) *NotifyPayCodePayResultRequest {
	s.TradeErrorMsg = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetTradeNo(v string) *NotifyPayCodePayResultRequest {
	s.TradeNo = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetTradeStatus(v string) *NotifyPayCodePayResultRequest {
	s.TradeStatus = &v
	return s
}

func (s *NotifyPayCodePayResultRequest) SetUserId(v string) *NotifyPayCodePayResultRequest {
	s.UserId = &v
	return s
}

type NotifyPayCodePayResultRequestPayChannelDetailList struct {
	// This parameter is required.
	//
	// example:
	//
	// 1.23
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// This parameter is required.
	FundToolDetailList []*NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList `json:"fundToolDetailList,omitempty" xml:"fundToolDetailList,omitempty" type:"Repeated"`
	// example:
	//
	// 2021-01-01 11:11:11
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2021-01-01 11:11:11
	GmtFinish *string `json:"gmtFinish,omitempty" xml:"gmtFinish,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 卡余额
	PayChannelName *string `json:"payChannelName,omitempty" xml:"payChannelName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20211234
	PayChannelOrderNo *string `json:"payChannelOrderNo,omitempty" xml:"payChannelOrderNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALIPAY|BALANCE
	PayChannelType *string `json:"payChannelType,omitempty" xml:"payChannelType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.00
	PromotionAmount *string `json:"promotionAmount,omitempty" xml:"promotionAmount,omitempty"`
}

func (s NotifyPayCodePayResultRequestPayChannelDetailList) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayCodePayResultRequestPayChannelDetailList) GoString() string {
	return s.String()
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailList) SetAmount(v string) *NotifyPayCodePayResultRequestPayChannelDetailList {
	s.Amount = &v
	return s
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailList) SetFundToolDetailList(v []*NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList) *NotifyPayCodePayResultRequestPayChannelDetailList {
	s.FundToolDetailList = v
	return s
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailList) SetGmtCreate(v string) *NotifyPayCodePayResultRequestPayChannelDetailList {
	s.GmtCreate = &v
	return s
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailList) SetGmtFinish(v string) *NotifyPayCodePayResultRequestPayChannelDetailList {
	s.GmtFinish = &v
	return s
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailList) SetPayChannelName(v string) *NotifyPayCodePayResultRequestPayChannelDetailList {
	s.PayChannelName = &v
	return s
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailList) SetPayChannelOrderNo(v string) *NotifyPayCodePayResultRequestPayChannelDetailList {
	s.PayChannelOrderNo = &v
	return s
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailList) SetPayChannelType(v string) *NotifyPayCodePayResultRequestPayChannelDetailList {
	s.PayChannelType = &v
	return s
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailList) SetPromotionAmount(v string) *NotifyPayCodePayResultRequestPayChannelDetailList {
	s.PromotionAmount = &v
	return s
}

type NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList struct {
	// This parameter is required.
	//
	// example:
	//
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// example:
	//
	// {"key":"value"}
	ExtInfo *string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 优惠券
	FundToolName *string `json:"fundToolName,omitempty" xml:"fundToolName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-01-01
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-01-01 11:11:11
	GmtFinish *string `json:"gmtFinish,omitempty" xml:"gmtFinish,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	PromotionFundTool *bool `json:"promotionFundTool,omitempty" xml:"promotionFundTool,omitempty"`
}

func (s NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList) GoString() string {
	return s.String()
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList) SetAmount(v string) *NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList {
	s.Amount = &v
	return s
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList) SetExtInfo(v string) *NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList {
	s.ExtInfo = &v
	return s
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList) SetFundToolName(v string) *NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList {
	s.FundToolName = &v
	return s
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList) SetGmtCreate(v string) *NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList {
	s.GmtCreate = &v
	return s
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList) SetGmtFinish(v string) *NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList {
	s.GmtFinish = &v
	return s
}

func (s *NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList) SetPromotionFundTool(v bool) *NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList {
	s.PromotionFundTool = &v
	return s
}

type NotifyPayCodePayResultResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// SUCCESS
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s NotifyPayCodePayResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayCodePayResultResponseBody) GoString() string {
	return s.String()
}

func (s *NotifyPayCodePayResultResponseBody) SetResult(v string) *NotifyPayCodePayResultResponseBody {
	s.Result = &v
	return s
}

type NotifyPayCodePayResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NotifyPayCodePayResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NotifyPayCodePayResultResponse) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayCodePayResultResponse) GoString() string {
	return s.String()
}

func (s *NotifyPayCodePayResultResponse) SetHeaders(v map[string]*string) *NotifyPayCodePayResultResponse {
	s.Headers = v
	return s
}

func (s *NotifyPayCodePayResultResponse) SetStatusCode(v int32) *NotifyPayCodePayResultResponse {
	s.StatusCode = &v
	return s
}

func (s *NotifyPayCodePayResultResponse) SetBody(v *NotifyPayCodePayResultResponseBody) *NotifyPayCodePayResultResponse {
	s.Body = v
	return s
}

type NotifyPayCodeRefundResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s NotifyPayCodeRefundResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayCodeRefundResultHeaders) GoString() string {
	return s.String()
}

func (s *NotifyPayCodeRefundResultHeaders) SetCommonHeaders(v map[string]*string) *NotifyPayCodeRefundResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *NotifyPayCodeRefundResultHeaders) SetXAcsDingtalkAccessToken(v string) *NotifyPayCodeRefundResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type NotifyPayCodeRefundResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ding1234
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-11 11:11:11
	GmtRefund *string `json:"gmtRefund,omitempty" xml:"gmtRefund,omitempty"`
	// This parameter is required.
	PayChannelDetailList []*NotifyPayCodeRefundResultRequestPayChannelDetailList `json:"payChannelDetailList,omitempty" xml:"payChannelDetailList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// payCode
	PayCode *string `json:"payCode,omitempty" xml:"payCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1.00
	RefundAmount *string `json:"refundAmount,omitempty" xml:"refundAmount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// refundOrderNo
	RefundOrderNo *string `json:"refundOrderNo,omitempty" xml:"refundOrderNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.00
	RefundPromotionAmount *string `json:"refundPromotionAmount,omitempty" xml:"refundPromotionAmount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 晚餐退款
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// tradeNo
	TradeNo *string `json:"tradeNo,omitempty" xml:"tradeNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s NotifyPayCodeRefundResultRequest) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayCodeRefundResultRequest) GoString() string {
	return s.String()
}

func (s *NotifyPayCodeRefundResultRequest) SetCorpId(v string) *NotifyPayCodeRefundResultRequest {
	s.CorpId = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequest) SetGmtRefund(v string) *NotifyPayCodeRefundResultRequest {
	s.GmtRefund = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequest) SetPayChannelDetailList(v []*NotifyPayCodeRefundResultRequestPayChannelDetailList) *NotifyPayCodeRefundResultRequest {
	s.PayChannelDetailList = v
	return s
}

func (s *NotifyPayCodeRefundResultRequest) SetPayCode(v string) *NotifyPayCodeRefundResultRequest {
	s.PayCode = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequest) SetRefundAmount(v string) *NotifyPayCodeRefundResultRequest {
	s.RefundAmount = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequest) SetRefundOrderNo(v string) *NotifyPayCodeRefundResultRequest {
	s.RefundOrderNo = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequest) SetRefundPromotionAmount(v string) *NotifyPayCodeRefundResultRequest {
	s.RefundPromotionAmount = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequest) SetRemark(v string) *NotifyPayCodeRefundResultRequest {
	s.Remark = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequest) SetTradeNo(v string) *NotifyPayCodeRefundResultRequest {
	s.TradeNo = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequest) SetUserId(v string) *NotifyPayCodeRefundResultRequest {
	s.UserId = &v
	return s
}

type NotifyPayCodeRefundResultRequestPayChannelDetailList struct {
	// This parameter is required.
	//
	// example:
	//
	// 1.00
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// This parameter is required.
	FundToolDetailList []*NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList `json:"fundToolDetailList,omitempty" xml:"fundToolDetailList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// ALIPAY
	PayChannelName *string `json:"payChannelName,omitempty" xml:"payChannelName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20210531123456
	PayChannelOrderNo *string `json:"payChannelOrderNo,omitempty" xml:"payChannelOrderNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021053112345678
	PayChannelRefundOrderNo *string `json:"payChannelRefundOrderNo,omitempty" xml:"payChannelRefundOrderNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALIPAY
	PayChannelType *string `json:"payChannelType,omitempty" xml:"payChannelType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.00
	PromotionAmount *string `json:"promotionAmount,omitempty" xml:"promotionAmount,omitempty"`
}

func (s NotifyPayCodeRefundResultRequestPayChannelDetailList) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayCodeRefundResultRequestPayChannelDetailList) GoString() string {
	return s.String()
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailList) SetAmount(v string) *NotifyPayCodeRefundResultRequestPayChannelDetailList {
	s.Amount = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailList) SetFundToolDetailList(v []*NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList) *NotifyPayCodeRefundResultRequestPayChannelDetailList {
	s.FundToolDetailList = v
	return s
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailList) SetPayChannelName(v string) *NotifyPayCodeRefundResultRequestPayChannelDetailList {
	s.PayChannelName = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailList) SetPayChannelOrderNo(v string) *NotifyPayCodeRefundResultRequestPayChannelDetailList {
	s.PayChannelOrderNo = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailList) SetPayChannelRefundOrderNo(v string) *NotifyPayCodeRefundResultRequestPayChannelDetailList {
	s.PayChannelRefundOrderNo = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailList) SetPayChannelType(v string) *NotifyPayCodeRefundResultRequestPayChannelDetailList {
	s.PayChannelType = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailList) SetPromotionAmount(v string) *NotifyPayCodeRefundResultRequestPayChannelDetailList {
	s.PromotionAmount = &v
	return s
}

type NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList struct {
	// This parameter is required.
	//
	// example:
	//
	// 1.00
	Amount  *string `json:"amount,omitempty" xml:"amount,omitempty"`
	ExtInfo *string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 余额
	FundToolName *string `json:"fundToolName,omitempty" xml:"fundToolName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-05-31 11:11:11
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-05-31 11:11:11
	GmtFinish *string `json:"gmtFinish,omitempty" xml:"gmtFinish,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	PromotionFundTool *bool `json:"promotionFundTool,omitempty" xml:"promotionFundTool,omitempty"`
}

func (s NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList) GoString() string {
	return s.String()
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList) SetAmount(v string) *NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList {
	s.Amount = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList) SetExtInfo(v string) *NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList {
	s.ExtInfo = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList) SetFundToolName(v string) *NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList {
	s.FundToolName = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList) SetGmtCreate(v string) *NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList {
	s.GmtCreate = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList) SetGmtFinish(v string) *NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList {
	s.GmtFinish = &v
	return s
}

func (s *NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList) SetPromotionFundTool(v bool) *NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList {
	s.PromotionFundTool = &v
	return s
}

type NotifyPayCodeRefundResultResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// SUCCESS
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s NotifyPayCodeRefundResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayCodeRefundResultResponseBody) GoString() string {
	return s.String()
}

func (s *NotifyPayCodeRefundResultResponseBody) SetResult(v string) *NotifyPayCodeRefundResultResponseBody {
	s.Result = &v
	return s
}

type NotifyPayCodeRefundResultResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NotifyPayCodeRefundResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NotifyPayCodeRefundResultResponse) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayCodeRefundResultResponse) GoString() string {
	return s.String()
}

func (s *NotifyPayCodeRefundResultResponse) SetHeaders(v map[string]*string) *NotifyPayCodeRefundResultResponse {
	s.Headers = v
	return s
}

func (s *NotifyPayCodeRefundResultResponse) SetStatusCode(v int32) *NotifyPayCodeRefundResultResponse {
	s.StatusCode = &v
	return s
}

func (s *NotifyPayCodeRefundResultResponse) SetBody(v *NotifyPayCodeRefundResultResponseBody) *NotifyPayCodeRefundResultResponse {
	s.Body = v
	return s
}

type NotifyVerifyResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s NotifyVerifyResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s NotifyVerifyResultHeaders) GoString() string {
	return s.String()
}

func (s *NotifyVerifyResultHeaders) SetCommonHeaders(v map[string]*string) *NotifyVerifyResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *NotifyVerifyResultHeaders) SetXAcsDingtalkAccessToken(v string) *NotifyVerifyResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type NotifyVerifyResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// corpxxxx
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 261234567890
	PayCode *string `json:"payCode,omitempty" xml:"payCode,omitempty"`
	Remark  *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// INTERNAL_STAFF
	UserCorpRelationType *string `json:"userCorpRelationType,omitempty" xml:"userCorpRelationType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxxxx
	UserIdentity *string `json:"userIdentity,omitempty" xml:"userIdentity,omitempty"`
	// example:
	//
	// 门禁验证
	VerifyEvent *string `json:"verifyEvent,omitempty" xml:"verifyEvent,omitempty"`
	// example:
	//
	// 1号食堂
	VerifyLocation *string `json:"verifyLocation,omitempty" xml:"verifyLocation,omitempty"`
	// example:
	//
	// 202112120003232
	VerifyNo *string `json:"verifyNo,omitempty" xml:"verifyNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 是否通过
	VerifyResult *bool `json:"verifyResult,omitempty" xml:"verifyResult,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-01-01 12:12:12
	VerifyTime *string `json:"verifyTime,omitempty" xml:"verifyTime,omitempty"`
}

func (s NotifyVerifyResultRequest) String() string {
	return tea.Prettify(s)
}

func (s NotifyVerifyResultRequest) GoString() string {
	return s.String()
}

func (s *NotifyVerifyResultRequest) SetCorpId(v string) *NotifyVerifyResultRequest {
	s.CorpId = &v
	return s
}

func (s *NotifyVerifyResultRequest) SetPayCode(v string) *NotifyVerifyResultRequest {
	s.PayCode = &v
	return s
}

func (s *NotifyVerifyResultRequest) SetRemark(v string) *NotifyVerifyResultRequest {
	s.Remark = &v
	return s
}

func (s *NotifyVerifyResultRequest) SetUserCorpRelationType(v string) *NotifyVerifyResultRequest {
	s.UserCorpRelationType = &v
	return s
}

func (s *NotifyVerifyResultRequest) SetUserIdentity(v string) *NotifyVerifyResultRequest {
	s.UserIdentity = &v
	return s
}

func (s *NotifyVerifyResultRequest) SetVerifyEvent(v string) *NotifyVerifyResultRequest {
	s.VerifyEvent = &v
	return s
}

func (s *NotifyVerifyResultRequest) SetVerifyLocation(v string) *NotifyVerifyResultRequest {
	s.VerifyLocation = &v
	return s
}

func (s *NotifyVerifyResultRequest) SetVerifyNo(v string) *NotifyVerifyResultRequest {
	s.VerifyNo = &v
	return s
}

func (s *NotifyVerifyResultRequest) SetVerifyResult(v bool) *NotifyVerifyResultRequest {
	s.VerifyResult = &v
	return s
}

func (s *NotifyVerifyResultRequest) SetVerifyTime(v string) *NotifyVerifyResultRequest {
	s.VerifyTime = &v
	return s
}

type NotifyVerifyResultResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// SUCCESS
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s NotifyVerifyResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NotifyVerifyResultResponseBody) GoString() string {
	return s.String()
}

func (s *NotifyVerifyResultResponseBody) SetResult(v string) *NotifyVerifyResultResponseBody {
	s.Result = &v
	return s
}

type NotifyVerifyResultResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NotifyVerifyResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NotifyVerifyResultResponse) String() string {
	return tea.Prettify(s)
}

func (s NotifyVerifyResultResponse) GoString() string {
	return s.String()
}

func (s *NotifyVerifyResultResponse) SetHeaders(v map[string]*string) *NotifyVerifyResultResponse {
	s.Headers = v
	return s
}

func (s *NotifyVerifyResultResponse) SetStatusCode(v int32) *NotifyVerifyResultResponse {
	s.StatusCode = &v
	return s
}

func (s *NotifyVerifyResultResponse) SetBody(v *NotifyVerifyResultResponseBody) *NotifyVerifyResultResponse {
	s.Body = v
	return s
}

type PreCreateGroupBillOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PreCreateGroupBillOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s PreCreateGroupBillOrderHeaders) GoString() string {
	return s.String()
}

func (s *PreCreateGroupBillOrderHeaders) SetCommonHeaders(v map[string]*string) *PreCreateGroupBillOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PreCreateGroupBillOrderHeaders) SetXAcsDingtalkAccessToken(v string) *PreCreateGroupBillOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PreCreateGroupBillOrderRequest struct {
	// This parameter is required.
	BillItemList []*PreCreateGroupBillOrderRequestBillItemList `json:"billItemList,omitempty" xml:"billItemList,omitempty" type:"Repeated"`
	ExtParams    map[string]*string                            `json:"extParams,omitempty" xml:"extParams,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	HeadCount *int64 `json:"headCount,omitempty" xml:"headCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	IsAverageAmount *int32 `json:"isAverageAmount,omitempty" xml:"isAverageAmount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dhqhadsnkj2qweqsw2
	MerchantId *string `json:"merchantId,omitempty" xml:"merchantId,omitempty"`
	// example:
	//
	// opemcesdjuwqw2uwnedj==
	OpenCid *string `json:"openCid,omitempty" xml:"openCid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20230918291921929193911
	OutBizNo *string `json:"outBizNo,omitempty" xml:"outBizNo,omitempty"`
	// example:
	//
	// ding32fff839a3e0105d
	PayeeCorpId *string `json:"payeeCorpId,omitempty" xml:"payeeCorpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ECEjwiiwenwnw2q2sdd
	PayeeUnionId *string `json:"payeeUnionId,omitempty" xml:"payeeUnionId,omitempty"`
	// example:
	//
	// 饿了么拼单-测试
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10.24
	TotalAmount *string `json:"totalAmount,omitempty" xml:"totalAmount,omitempty"`
}

func (s PreCreateGroupBillOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s PreCreateGroupBillOrderRequest) GoString() string {
	return s.String()
}

func (s *PreCreateGroupBillOrderRequest) SetBillItemList(v []*PreCreateGroupBillOrderRequestBillItemList) *PreCreateGroupBillOrderRequest {
	s.BillItemList = v
	return s
}

func (s *PreCreateGroupBillOrderRequest) SetExtParams(v map[string]*string) *PreCreateGroupBillOrderRequest {
	s.ExtParams = v
	return s
}

func (s *PreCreateGroupBillOrderRequest) SetHeadCount(v int64) *PreCreateGroupBillOrderRequest {
	s.HeadCount = &v
	return s
}

func (s *PreCreateGroupBillOrderRequest) SetIsAverageAmount(v int32) *PreCreateGroupBillOrderRequest {
	s.IsAverageAmount = &v
	return s
}

func (s *PreCreateGroupBillOrderRequest) SetMerchantId(v string) *PreCreateGroupBillOrderRequest {
	s.MerchantId = &v
	return s
}

func (s *PreCreateGroupBillOrderRequest) SetOpenCid(v string) *PreCreateGroupBillOrderRequest {
	s.OpenCid = &v
	return s
}

func (s *PreCreateGroupBillOrderRequest) SetOutBizNo(v string) *PreCreateGroupBillOrderRequest {
	s.OutBizNo = &v
	return s
}

func (s *PreCreateGroupBillOrderRequest) SetPayeeCorpId(v string) *PreCreateGroupBillOrderRequest {
	s.PayeeCorpId = &v
	return s
}

func (s *PreCreateGroupBillOrderRequest) SetPayeeUnionId(v string) *PreCreateGroupBillOrderRequest {
	s.PayeeUnionId = &v
	return s
}

func (s *PreCreateGroupBillOrderRequest) SetRemark(v string) *PreCreateGroupBillOrderRequest {
	s.Remark = &v
	return s
}

func (s *PreCreateGroupBillOrderRequest) SetTotalAmount(v string) *PreCreateGroupBillOrderRequest {
	s.TotalAmount = &v
	return s
}

type PreCreateGroupBillOrderRequestBillItemList struct {
	// This parameter is required.
	//
	// example:
	//
	// 5.12
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cshadbikahdksnajhada
	PayerUnionId *string `json:"payerUnionId,omitempty" xml:"payerUnionId,omitempty"`
}

func (s PreCreateGroupBillOrderRequestBillItemList) String() string {
	return tea.Prettify(s)
}

func (s PreCreateGroupBillOrderRequestBillItemList) GoString() string {
	return s.String()
}

func (s *PreCreateGroupBillOrderRequestBillItemList) SetAmount(v string) *PreCreateGroupBillOrderRequestBillItemList {
	s.Amount = &v
	return s
}

func (s *PreCreateGroupBillOrderRequestBillItemList) SetPayerUnionId(v string) *PreCreateGroupBillOrderRequestBillItemList {
	s.PayerUnionId = &v
	return s
}

type PreCreateGroupBillOrderResponseBody struct {
	Result  *PreCreateGroupBillOrderResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PreCreateGroupBillOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PreCreateGroupBillOrderResponseBody) GoString() string {
	return s.String()
}

func (s *PreCreateGroupBillOrderResponseBody) SetResult(v *PreCreateGroupBillOrderResponseBodyResult) *PreCreateGroupBillOrderResponseBody {
	s.Result = v
	return s
}

func (s *PreCreateGroupBillOrderResponseBody) SetSuccess(v bool) *PreCreateGroupBillOrderResponseBody {
	s.Success = &v
	return s
}

type PreCreateGroupBillOrderResponseBodyResult struct {
	// example:
	//
	// 2023100914312930910100002107362525
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
}

func (s PreCreateGroupBillOrderResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PreCreateGroupBillOrderResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PreCreateGroupBillOrderResponseBodyResult) SetOrderNo(v string) *PreCreateGroupBillOrderResponseBodyResult {
	s.OrderNo = &v
	return s
}

type PreCreateGroupBillOrderResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreCreateGroupBillOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreCreateGroupBillOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s PreCreateGroupBillOrderResponse) GoString() string {
	return s.String()
}

func (s *PreCreateGroupBillOrderResponse) SetHeaders(v map[string]*string) *PreCreateGroupBillOrderResponse {
	s.Headers = v
	return s
}

func (s *PreCreateGroupBillOrderResponse) SetStatusCode(v int32) *PreCreateGroupBillOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *PreCreateGroupBillOrderResponse) SetBody(v *PreCreateGroupBillOrderResponseBody) *PreCreateGroupBillOrderResponse {
	s.Body = v
	return s
}

type QueryAcquireRefundOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAcquireRefundOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAcquireRefundOrderHeaders) GoString() string {
	return s.String()
}

func (s *QueryAcquireRefundOrderHeaders) SetCommonHeaders(v map[string]*string) *QueryAcquireRefundOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAcquireRefundOrderHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAcquireRefundOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAcquireRefundOrderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 202100001
	OutRefundNo *string `json:"outRefundNo,omitempty" xml:"outRefundNo,omitempty"`
}

func (s QueryAcquireRefundOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAcquireRefundOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryAcquireRefundOrderRequest) SetOutRefundNo(v string) *QueryAcquireRefundOrderRequest {
	s.OutRefundNo = &v
	return s
}

type QueryAcquireRefundOrderResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 10.01
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-15 10:10:09
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2021-11-15 10:10:10
	GmtRefund *string `json:"gmtRefund,omitempty" xml:"gmtRefund,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 202111010001
	InstId *string `json:"instId,omitempty" xml:"instId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 202121241343151
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 202111020001
	OriginOutTradeNo *string `json:"originOutTradeNo,omitempty" xml:"originOutTradeNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// r202111020001
	OutRefundNo *string `json:"outRefundNo,omitempty" xml:"outRefundNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALIPAY
	PayChannel *string `json:"payChannel,omitempty" xml:"payChannel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13****09
	PayChannelAccountNo *string `json:"payChannelAccountNo,omitempty" xml:"payChannelAccountNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123124
	PayerUserId *string `json:"payerUserId,omitempty" xml:"payerUserId,omitempty"`
	// example:
	//
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1001
	SubInstId *string `json:"subInstId,omitempty" xml:"subInstId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 餐费
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryAcquireRefundOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAcquireRefundOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAcquireRefundOrderResponseBody) SetAmount(v string) *QueryAcquireRefundOrderResponseBody {
	s.Amount = &v
	return s
}

func (s *QueryAcquireRefundOrderResponseBody) SetGmtCreate(v string) *QueryAcquireRefundOrderResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *QueryAcquireRefundOrderResponseBody) SetGmtRefund(v string) *QueryAcquireRefundOrderResponseBody {
	s.GmtRefund = &v
	return s
}

func (s *QueryAcquireRefundOrderResponseBody) SetInstId(v string) *QueryAcquireRefundOrderResponseBody {
	s.InstId = &v
	return s
}

func (s *QueryAcquireRefundOrderResponseBody) SetOrderNo(v string) *QueryAcquireRefundOrderResponseBody {
	s.OrderNo = &v
	return s
}

func (s *QueryAcquireRefundOrderResponseBody) SetOriginOutTradeNo(v string) *QueryAcquireRefundOrderResponseBody {
	s.OriginOutTradeNo = &v
	return s
}

func (s *QueryAcquireRefundOrderResponseBody) SetOutRefundNo(v string) *QueryAcquireRefundOrderResponseBody {
	s.OutRefundNo = &v
	return s
}

func (s *QueryAcquireRefundOrderResponseBody) SetPayChannel(v string) *QueryAcquireRefundOrderResponseBody {
	s.PayChannel = &v
	return s
}

func (s *QueryAcquireRefundOrderResponseBody) SetPayChannelAccountNo(v string) *QueryAcquireRefundOrderResponseBody {
	s.PayChannelAccountNo = &v
	return s
}

func (s *QueryAcquireRefundOrderResponseBody) SetPayerUserId(v string) *QueryAcquireRefundOrderResponseBody {
	s.PayerUserId = &v
	return s
}

func (s *QueryAcquireRefundOrderResponseBody) SetRemark(v string) *QueryAcquireRefundOrderResponseBody {
	s.Remark = &v
	return s
}

func (s *QueryAcquireRefundOrderResponseBody) SetStatus(v string) *QueryAcquireRefundOrderResponseBody {
	s.Status = &v
	return s
}

func (s *QueryAcquireRefundOrderResponseBody) SetSubInstId(v string) *QueryAcquireRefundOrderResponseBody {
	s.SubInstId = &v
	return s
}

func (s *QueryAcquireRefundOrderResponseBody) SetTitle(v string) *QueryAcquireRefundOrderResponseBody {
	s.Title = &v
	return s
}

type QueryAcquireRefundOrderResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAcquireRefundOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAcquireRefundOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAcquireRefundOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryAcquireRefundOrderResponse) SetHeaders(v map[string]*string) *QueryAcquireRefundOrderResponse {
	s.Headers = v
	return s
}

func (s *QueryAcquireRefundOrderResponse) SetStatusCode(v int32) *QueryAcquireRefundOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAcquireRefundOrderResponse) SetBody(v *QueryAcquireRefundOrderResponseBody) *QueryAcquireRefundOrderResponse {
	s.Body = v
	return s
}

type QueryBatchTradeDetailListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryBatchTradeDetailListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchTradeDetailListHeaders) GoString() string {
	return s.String()
}

func (s *QueryBatchTradeDetailListHeaders) SetCommonHeaders(v map[string]*string) *QueryBatchTradeDetailListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryBatchTradeDetailListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryBatchTradeDetailListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryBatchTradeDetailListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20210901001
	OutBatchNo *string `json:"outBatchNo,omitempty" xml:"outBatchNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryBatchTradeDetailListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchTradeDetailListRequest) GoString() string {
	return s.String()
}

func (s *QueryBatchTradeDetailListRequest) SetOutBatchNo(v string) *QueryBatchTradeDetailListRequest {
	s.OutBatchNo = &v
	return s
}

func (s *QueryBatchTradeDetailListRequest) SetPageNumber(v int32) *QueryBatchTradeDetailListRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryBatchTradeDetailListRequest) SetPageSize(v int32) *QueryBatchTradeDetailListRequest {
	s.PageSize = &v
	return s
}

type QueryBatchTradeDetailListResponseBody struct {
	// This parameter is required.
	BatchTradeDetailList []*QueryBatchTradeDetailListResponseBodyBatchTradeDetailList `json:"batchTradeDetailList,omitempty" xml:"batchTradeDetailList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TotalPageNumber *int32 `json:"totalPageNumber,omitempty" xml:"totalPageNumber,omitempty"`
}

func (s QueryBatchTradeDetailListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchTradeDetailListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBatchTradeDetailListResponseBody) SetBatchTradeDetailList(v []*QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) *QueryBatchTradeDetailListResponseBody {
	s.BatchTradeDetailList = v
	return s
}

func (s *QueryBatchTradeDetailListResponseBody) SetPageNumber(v int32) *QueryBatchTradeDetailListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryBatchTradeDetailListResponseBody) SetPageSize(v int32) *QueryBatchTradeDetailListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryBatchTradeDetailListResponseBody) SetTotal(v int32) *QueryBatchTradeDetailListResponseBody {
	s.Total = &v
	return s
}

func (s *QueryBatchTradeDetailListResponseBody) SetTotalPageNumber(v int32) *QueryBatchTradeDetailListResponseBody {
	s.TotalPageNumber = &v
	return s
}

type QueryBatchTradeDetailListResponseBodyBatchTradeDetailList struct {
	// This parameter is required.
	//
	// example:
	//
	// 1.00
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20210909153300000002734746770740
	DetailNo *string `json:"detailNo,omitempty" xml:"detailNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-09-29 14:46:48
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-09-29 16:05:00
	GmtFinish *string `json:"gmtFinish,omitempty" xml:"gmtFinish,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 收款人
	PayeeAccountName *string `json:"payeeAccountName,omitempty" xml:"payeeAccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13900000000
	PayeeAccountNo *string `json:"payeeAccountNo,omitempty" xml:"payeeAccountNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALIPAY
	PayeeAccountType *string `json:"payeeAccountType,omitempty" xml:"payeeAccountType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SerialNo *int64 `json:"serialNo,omitempty" xml:"serialNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) GoString() string {
	return s.String()
}

func (s *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) SetAmount(v string) *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList {
	s.Amount = &v
	return s
}

func (s *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) SetDetailNo(v string) *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList {
	s.DetailNo = &v
	return s
}

func (s *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) SetGmtCreate(v string) *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList {
	s.GmtCreate = &v
	return s
}

func (s *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) SetGmtFinish(v string) *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList {
	s.GmtFinish = &v
	return s
}

func (s *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) SetMemo(v string) *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList {
	s.Memo = &v
	return s
}

func (s *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) SetPayeeAccountName(v string) *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList {
	s.PayeeAccountName = &v
	return s
}

func (s *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) SetPayeeAccountNo(v string) *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList {
	s.PayeeAccountNo = &v
	return s
}

func (s *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) SetPayeeAccountType(v string) *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList {
	s.PayeeAccountType = &v
	return s
}

func (s *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) SetSerialNo(v int64) *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList {
	s.SerialNo = &v
	return s
}

func (s *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList) SetStatus(v string) *QueryBatchTradeDetailListResponseBodyBatchTradeDetailList {
	s.Status = &v
	return s
}

type QueryBatchTradeDetailListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBatchTradeDetailListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBatchTradeDetailListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchTradeDetailListResponse) GoString() string {
	return s.String()
}

func (s *QueryBatchTradeDetailListResponse) SetHeaders(v map[string]*string) *QueryBatchTradeDetailListResponse {
	s.Headers = v
	return s
}

func (s *QueryBatchTradeDetailListResponse) SetStatusCode(v int32) *QueryBatchTradeDetailListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBatchTradeDetailListResponse) SetBody(v *QueryBatchTradeDetailListResponseBody) *QueryBatchTradeDetailListResponse {
	s.Body = v
	return s
}

type QueryBatchTradeOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryBatchTradeOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchTradeOrderHeaders) GoString() string {
	return s.String()
}

func (s *QueryBatchTradeOrderHeaders) SetCommonHeaders(v map[string]*string) *QueryBatchTradeOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryBatchTradeOrderHeaders) SetXAcsDingtalkAccessToken(v string) *QueryBatchTradeOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryBatchTradeOrderRequest struct {
	// This parameter is required.
	OutBatchNos []*string `json:"outBatchNos,omitempty" xml:"outBatchNos,omitempty" type:"Repeated"`
}

func (s QueryBatchTradeOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchTradeOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryBatchTradeOrderRequest) SetOutBatchNos(v []*string) *QueryBatchTradeOrderRequest {
	s.OutBatchNos = v
	return s
}

type QueryBatchTradeOrderResponseBody struct {
	// This parameter is required.
	BatchTradeOrderVOs []*QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs `json:"batchTradeOrderVOs,omitempty" xml:"batchTradeOrderVOs,omitempty" type:"Repeated"`
}

func (s QueryBatchTradeOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchTradeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBatchTradeOrderResponseBody) SetBatchTradeOrderVOs(v []*QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) *QueryBatchTradeOrderResponseBody {
	s.BatchTradeOrderVOs = v
	return s
}

type QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs struct {
	// This parameter is required.
	//
	// example:
	//
	// 2021090102102122200002121
	AlipayTransId *string `json:"alipayTransId,omitempty" xml:"alipayTransId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	FailAmount *string `json:"failAmount,omitempty" xml:"failAmount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	FailCount *int64 `json:"failCount,omitempty" xml:"failCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 收款人不存在
	FailReason *string `json:"failReason,omitempty" xml:"failReason,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-09-01 12:00:00
	GmtFinish *string `json:"gmtFinish,omitempty" xml:"gmtFinish,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-09-01 11:00:00
	GmtSubmit *string `json:"gmtSubmit,omitempty" xml:"gmtSubmit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20210901001
	OutBatchNo *string `json:"outBatchNo,omitempty" xml:"outBatchNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 213654465
	PayerStaffId *string `json:"payerStaffId,omitempty" xml:"payerStaffId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1.00
	PaymentAmount *string `json:"paymentAmount,omitempty" xml:"paymentAmount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CNY
	PaymentCurrency *string `json:"paymentCurrency,omitempty" xml:"paymentCurrency,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1.00
	SuccessAmount *string `json:"successAmount,omitempty" xml:"successAmount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SuccessCount *int64 `json:"successCount,omitempty" xml:"successCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1.00
	TotalAmount *string `json:"totalAmount,omitempty" xml:"totalAmount,omitempty"`
}

func (s QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) GoString() string {
	return s.String()
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetAlipayTransId(v string) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.AlipayTransId = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetFailAmount(v string) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.FailAmount = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetFailCount(v int64) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.FailCount = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetFailReason(v string) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.FailReason = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetGmtFinish(v string) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.GmtFinish = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetGmtSubmit(v string) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.GmtSubmit = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetOutBatchNo(v string) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.OutBatchNo = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetPayerStaffId(v string) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.PayerStaffId = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetPaymentAmount(v string) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.PaymentAmount = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetPaymentCurrency(v string) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.PaymentCurrency = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetStatus(v string) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.Status = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetSuccessAmount(v string) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.SuccessAmount = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetSuccessCount(v int64) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.SuccessCount = &v
	return s
}

func (s *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs) SetTotalAmount(v string) *QueryBatchTradeOrderResponseBodyBatchTradeOrderVOs {
	s.TotalAmount = &v
	return s
}

type QueryBatchTradeOrderResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBatchTradeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBatchTradeOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchTradeOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryBatchTradeOrderResponse) SetHeaders(v map[string]*string) *QueryBatchTradeOrderResponse {
	s.Headers = v
	return s
}

func (s *QueryBatchTradeOrderResponse) SetStatusCode(v int32) *QueryBatchTradeOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBatchTradeOrderResponse) SetBody(v *QueryBatchTradeOrderResponseBody) *QueryBatchTradeOrderResponse {
	s.Body = v
	return s
}

type QueryPayAccountListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryPayAccountListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryPayAccountListHeaders) GoString() string {
	return s.String()
}

func (s *QueryPayAccountListHeaders) SetCommonHeaders(v map[string]*string) *QueryPayAccountListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryPayAccountListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryPayAccountListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryPayAccountListResponseBody struct {
	// This parameter is required.
	PayAccountVOList []*QueryPayAccountListResponseBodyPayAccountVOList `json:"payAccountVOList,omitempty" xml:"payAccountVOList,omitempty" type:"Repeated"`
}

func (s QueryPayAccountListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPayAccountListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPayAccountListResponseBody) SetPayAccountVOList(v []*QueryPayAccountListResponseBodyPayAccountVOList) *QueryPayAccountListResponseBody {
	s.PayAccountVOList = v
	return s
}

type QueryPayAccountListResponseBodyPayAccountVOList struct {
	// This parameter is required.
	//
	// example:
	//
	// B
	AccountClass *string `json:"accountClass,omitempty" xml:"accountClass,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20210912001
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 139****1
	AccountNo *string `json:"accountNo,omitempty" xml:"accountNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 备注
	AccountRemark *string `json:"accountRemark,omitempty" xml:"accountRemark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALIPAY
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
}

func (s QueryPayAccountListResponseBodyPayAccountVOList) String() string {
	return tea.Prettify(s)
}

func (s QueryPayAccountListResponseBodyPayAccountVOList) GoString() string {
	return s.String()
}

func (s *QueryPayAccountListResponseBodyPayAccountVOList) SetAccountClass(v string) *QueryPayAccountListResponseBodyPayAccountVOList {
	s.AccountClass = &v
	return s
}

func (s *QueryPayAccountListResponseBodyPayAccountVOList) SetAccountId(v string) *QueryPayAccountListResponseBodyPayAccountVOList {
	s.AccountId = &v
	return s
}

func (s *QueryPayAccountListResponseBodyPayAccountVOList) SetAccountName(v string) *QueryPayAccountListResponseBodyPayAccountVOList {
	s.AccountName = &v
	return s
}

func (s *QueryPayAccountListResponseBodyPayAccountVOList) SetAccountNo(v string) *QueryPayAccountListResponseBodyPayAccountVOList {
	s.AccountNo = &v
	return s
}

func (s *QueryPayAccountListResponseBodyPayAccountVOList) SetAccountRemark(v string) *QueryPayAccountListResponseBodyPayAccountVOList {
	s.AccountRemark = &v
	return s
}

func (s *QueryPayAccountListResponseBodyPayAccountVOList) SetAccountType(v string) *QueryPayAccountListResponseBodyPayAccountVOList {
	s.AccountType = &v
	return s
}

type QueryPayAccountListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPayAccountListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPayAccountListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPayAccountListResponse) GoString() string {
	return s.String()
}

func (s *QueryPayAccountListResponse) SetHeaders(v map[string]*string) *QueryPayAccountListResponse {
	s.Headers = v
	return s
}

func (s *QueryPayAccountListResponse) SetStatusCode(v int32) *QueryPayAccountListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPayAccountListResponse) SetBody(v *QueryPayAccountListResponseBody) *QueryPayAccountListResponse {
	s.Body = v
	return s
}

type QueryRegisterOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryRegisterOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryRegisterOrderHeaders) GoString() string {
	return s.String()
}

func (s *QueryRegisterOrderHeaders) SetCommonHeaders(v map[string]*string) *QueryRegisterOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryRegisterOrderHeaders) SetXAcsDingtalkAccessToken(v string) *QueryRegisterOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryRegisterOrderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 202111090001
	InstId *string `json:"instId,omitempty" xml:"instId,omitempty"`
	// example:
	//
	// 20211222000000001
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// example:
	//
	// 202112220001
	OutTradeNo *string `json:"outTradeNo,omitempty" xml:"outTradeNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	SubInstId *string `json:"subInstId,omitempty" xml:"subInstId,omitempty"`
}

func (s QueryRegisterOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRegisterOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryRegisterOrderRequest) SetInstId(v string) *QueryRegisterOrderRequest {
	s.InstId = &v
	return s
}

func (s *QueryRegisterOrderRequest) SetOrderId(v string) *QueryRegisterOrderRequest {
	s.OrderId = &v
	return s
}

func (s *QueryRegisterOrderRequest) SetOutTradeNo(v string) *QueryRegisterOrderRequest {
	s.OutTradeNo = &v
	return s
}

func (s *QueryRegisterOrderRequest) SetSubInstId(v string) *QueryRegisterOrderRequest {
	s.SubInstId = &v
	return s
}

type QueryRegisterOrderResponseBody struct {
	FailReason  *string `json:"failReason,omitempty" xml:"failReason,omitempty"`
	GmtAudit    *string `json:"gmtAudit,omitempty" xml:"gmtAudit,omitempty"`
	InstId      *string `json:"instId,omitempty" xml:"instId,omitempty"`
	OrderId     *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	OutTradeNo  *string `json:"outTradeNo,omitempty" xml:"outTradeNo,omitempty"`
	Status      *string `json:"status,omitempty" xml:"status,omitempty"`
	SubInstId   *string `json:"subInstId,omitempty" xml:"subInstId,omitempty"`
	SubInstName *string `json:"subInstName,omitempty" xml:"subInstName,omitempty"`
}

func (s QueryRegisterOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRegisterOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRegisterOrderResponseBody) SetFailReason(v string) *QueryRegisterOrderResponseBody {
	s.FailReason = &v
	return s
}

func (s *QueryRegisterOrderResponseBody) SetGmtAudit(v string) *QueryRegisterOrderResponseBody {
	s.GmtAudit = &v
	return s
}

func (s *QueryRegisterOrderResponseBody) SetInstId(v string) *QueryRegisterOrderResponseBody {
	s.InstId = &v
	return s
}

func (s *QueryRegisterOrderResponseBody) SetOrderId(v string) *QueryRegisterOrderResponseBody {
	s.OrderId = &v
	return s
}

func (s *QueryRegisterOrderResponseBody) SetOutTradeNo(v string) *QueryRegisterOrderResponseBody {
	s.OutTradeNo = &v
	return s
}

func (s *QueryRegisterOrderResponseBody) SetStatus(v string) *QueryRegisterOrderResponseBody {
	s.Status = &v
	return s
}

func (s *QueryRegisterOrderResponseBody) SetSubInstId(v string) *QueryRegisterOrderResponseBody {
	s.SubInstId = &v
	return s
}

func (s *QueryRegisterOrderResponseBody) SetSubInstName(v string) *QueryRegisterOrderResponseBody {
	s.SubInstName = &v
	return s
}

type QueryRegisterOrderResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRegisterOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRegisterOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRegisterOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryRegisterOrderResponse) SetHeaders(v map[string]*string) *QueryRegisterOrderResponse {
	s.Headers = v
	return s
}

func (s *QueryRegisterOrderResponse) SetStatusCode(v int32) *QueryRegisterOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRegisterOrderResponse) SetBody(v *QueryRegisterOrderResponseBody) *QueryRegisterOrderResponse {
	s.Body = v
	return s
}

type QueryUserAgreementHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserAgreementHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserAgreementHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserAgreementHeaders) SetCommonHeaders(v map[string]*string) *QueryUserAgreementHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserAgreementHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserAgreementHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserAgreementRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// TRADE
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// WITHHOLDING
	BizScene *string `json:"bizScene,omitempty" xml:"bizScene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 202111090001
	InstId *string `json:"instId,omitempty" xml:"instId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1001
	SubInstId *string `json:"subInstId,omitempty" xml:"subInstId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2120493284
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryUserAgreementRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserAgreementRequest) GoString() string {
	return s.String()
}

func (s *QueryUserAgreementRequest) SetBizCode(v string) *QueryUserAgreementRequest {
	s.BizCode = &v
	return s
}

func (s *QueryUserAgreementRequest) SetBizScene(v string) *QueryUserAgreementRequest {
	s.BizScene = &v
	return s
}

func (s *QueryUserAgreementRequest) SetInstId(v string) *QueryUserAgreementRequest {
	s.InstId = &v
	return s
}

func (s *QueryUserAgreementRequest) SetSubInstId(v string) *QueryUserAgreementRequest {
	s.SubInstId = &v
	return s
}

func (s *QueryUserAgreementRequest) SetUserId(v string) *QueryUserAgreementRequest {
	s.UserId = &v
	return s
}

type QueryUserAgreementResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 201220123212312
	AgreementNo *string `json:"agreementNo,omitempty" xml:"agreementNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ding123123234234
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-11 10:10:10
	GmtExpire *string `json:"gmtExpire,omitempty" xml:"gmtExpire,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-11 10:10:10
	GmtSign *string `json:"gmtSign,omitempty" xml:"gmtSign,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-11 10:10:10
	GmtValid *string `json:"gmtValid,omitempty" xml:"gmtValid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021000001
	InstId *string `json:"instId,omitempty" xml:"instId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALIPAY
	PayChannel *string `json:"payChannel,omitempty" xml:"payChannel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// *川
	PayChannelAccountName *string `json:"payChannelAccountName,omitempty" xml:"payChannelAccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 139*****0
	PayChannelAccountNo *string `json:"payChannelAccountNo,omitempty" xml:"payChannelAccountNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 101
	SubInstId *string `json:"subInstId,omitempty" xml:"subInstId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 54646545
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryUserAgreementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserAgreementResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserAgreementResponseBody) SetAgreementNo(v string) *QueryUserAgreementResponseBody {
	s.AgreementNo = &v
	return s
}

func (s *QueryUserAgreementResponseBody) SetCorpId(v string) *QueryUserAgreementResponseBody {
	s.CorpId = &v
	return s
}

func (s *QueryUserAgreementResponseBody) SetGmtExpire(v string) *QueryUserAgreementResponseBody {
	s.GmtExpire = &v
	return s
}

func (s *QueryUserAgreementResponseBody) SetGmtSign(v string) *QueryUserAgreementResponseBody {
	s.GmtSign = &v
	return s
}

func (s *QueryUserAgreementResponseBody) SetGmtValid(v string) *QueryUserAgreementResponseBody {
	s.GmtValid = &v
	return s
}

func (s *QueryUserAgreementResponseBody) SetInstId(v string) *QueryUserAgreementResponseBody {
	s.InstId = &v
	return s
}

func (s *QueryUserAgreementResponseBody) SetPayChannel(v string) *QueryUserAgreementResponseBody {
	s.PayChannel = &v
	return s
}

func (s *QueryUserAgreementResponseBody) SetPayChannelAccountName(v string) *QueryUserAgreementResponseBody {
	s.PayChannelAccountName = &v
	return s
}

func (s *QueryUserAgreementResponseBody) SetPayChannelAccountNo(v string) *QueryUserAgreementResponseBody {
	s.PayChannelAccountNo = &v
	return s
}

func (s *QueryUserAgreementResponseBody) SetStatus(v string) *QueryUserAgreementResponseBody {
	s.Status = &v
	return s
}

func (s *QueryUserAgreementResponseBody) SetSubInstId(v string) *QueryUserAgreementResponseBody {
	s.SubInstId = &v
	return s
}

func (s *QueryUserAgreementResponseBody) SetUserId(v string) *QueryUserAgreementResponseBody {
	s.UserId = &v
	return s
}

type QueryUserAgreementResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserAgreementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserAgreementResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserAgreementResponse) GoString() string {
	return s.String()
}

func (s *QueryUserAgreementResponse) SetHeaders(v map[string]*string) *QueryUserAgreementResponse {
	s.Headers = v
	return s
}

func (s *QueryUserAgreementResponse) SetStatusCode(v int32) *QueryUserAgreementResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserAgreementResponse) SetBody(v *QueryUserAgreementResponseBody) *QueryUserAgreementResponse {
	s.Body = v
	return s
}

type QueryUserAlipayAccountHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserAlipayAccountHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserAlipayAccountHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserAlipayAccountHeaders) SetCommonHeaders(v map[string]*string) *QueryUserAlipayAccountHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserAlipayAccountHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserAlipayAccountHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserAlipayAccountResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 2088894773487
	AlipayUid *string `json:"alipayUid,omitempty" xml:"alipayUid,omitempty"`
}

func (s QueryUserAlipayAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserAlipayAccountResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserAlipayAccountResponseBody) SetAlipayUid(v string) *QueryUserAlipayAccountResponseBody {
	s.AlipayUid = &v
	return s
}

type QueryUserAlipayAccountResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserAlipayAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserAlipayAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserAlipayAccountResponse) GoString() string {
	return s.String()
}

func (s *QueryUserAlipayAccountResponse) SetHeaders(v map[string]*string) *QueryUserAlipayAccountResponse {
	s.Headers = v
	return s
}

func (s *QueryUserAlipayAccountResponse) SetStatusCode(v int32) *QueryUserAlipayAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserAlipayAccountResponse) SetBody(v *QueryUserAlipayAccountResponseBody) *QueryUserAlipayAccountResponse {
	s.Body = v
	return s
}

type QueryWithholdingOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryWithholdingOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryWithholdingOrderHeaders) GoString() string {
	return s.String()
}

func (s *QueryWithholdingOrderHeaders) SetCommonHeaders(v map[string]*string) *QueryWithholdingOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryWithholdingOrderHeaders) SetXAcsDingtalkAccessToken(v string) *QueryWithholdingOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryWithholdingOrderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 202100001
	OutTradeNo *string `json:"outTradeNo,omitempty" xml:"outTradeNo,omitempty"`
}

func (s QueryWithholdingOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryWithholdingOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryWithholdingOrderRequest) SetOutTradeNo(v string) *QueryWithholdingOrderRequest {
	s.OutTradeNo = &v
	return s
}

type QueryWithholdingOrderResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 10.01
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-15 10:10:09
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2021-11-15 10:10:10
	GmtPay *string `json:"gmtPay,omitempty" xml:"gmtPay,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 202111010001
	InstId *string `json:"instId,omitempty" xml:"instId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 202121241343151
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 202111020001
	OutTradeNo *string `json:"outTradeNo,omitempty" xml:"outTradeNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALIPAY
	PayChannel *string `json:"payChannel,omitempty" xml:"payChannel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13****09
	PayChannelAccountNo *string `json:"payChannelAccountNo,omitempty" xml:"payChannelAccountNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123124
	PayerUserId *string `json:"payerUserId,omitempty" xml:"payerUserId,omitempty"`
	// example:
	//
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1001
	SubInstId *string `json:"subInstId,omitempty" xml:"subInstId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 餐费
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryWithholdingOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryWithholdingOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QueryWithholdingOrderResponseBody) SetAmount(v string) *QueryWithholdingOrderResponseBody {
	s.Amount = &v
	return s
}

func (s *QueryWithholdingOrderResponseBody) SetGmtCreate(v string) *QueryWithholdingOrderResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *QueryWithholdingOrderResponseBody) SetGmtPay(v string) *QueryWithholdingOrderResponseBody {
	s.GmtPay = &v
	return s
}

func (s *QueryWithholdingOrderResponseBody) SetInstId(v string) *QueryWithholdingOrderResponseBody {
	s.InstId = &v
	return s
}

func (s *QueryWithholdingOrderResponseBody) SetOrderNo(v string) *QueryWithholdingOrderResponseBody {
	s.OrderNo = &v
	return s
}

func (s *QueryWithholdingOrderResponseBody) SetOutTradeNo(v string) *QueryWithholdingOrderResponseBody {
	s.OutTradeNo = &v
	return s
}

func (s *QueryWithholdingOrderResponseBody) SetPayChannel(v string) *QueryWithholdingOrderResponseBody {
	s.PayChannel = &v
	return s
}

func (s *QueryWithholdingOrderResponseBody) SetPayChannelAccountNo(v string) *QueryWithholdingOrderResponseBody {
	s.PayChannelAccountNo = &v
	return s
}

func (s *QueryWithholdingOrderResponseBody) SetPayerUserId(v string) *QueryWithholdingOrderResponseBody {
	s.PayerUserId = &v
	return s
}

func (s *QueryWithholdingOrderResponseBody) SetRemark(v string) *QueryWithholdingOrderResponseBody {
	s.Remark = &v
	return s
}

func (s *QueryWithholdingOrderResponseBody) SetStatus(v string) *QueryWithholdingOrderResponseBody {
	s.Status = &v
	return s
}

func (s *QueryWithholdingOrderResponseBody) SetSubInstId(v string) *QueryWithholdingOrderResponseBody {
	s.SubInstId = &v
	return s
}

func (s *QueryWithholdingOrderResponseBody) SetTitle(v string) *QueryWithholdingOrderResponseBody {
	s.Title = &v
	return s
}

type QueryWithholdingOrderResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryWithholdingOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryWithholdingOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryWithholdingOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryWithholdingOrderResponse) SetHeaders(v map[string]*string) *QueryWithholdingOrderResponse {
	s.Headers = v
	return s
}

func (s *QueryWithholdingOrderResponse) SetStatusCode(v int32) *QueryWithholdingOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryWithholdingOrderResponse) SetBody(v *QueryWithholdingOrderResponseBody) *QueryWithholdingOrderResponse {
	s.Body = v
	return s
}

type SaveCorpPayCodeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveCorpPayCodeHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveCorpPayCodeHeaders) GoString() string {
	return s.String()
}

func (s *SaveCorpPayCodeHeaders) SetCommonHeaders(v map[string]*string) *SaveCorpPayCodeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveCorpPayCodeHeaders) SetXAcsDingtalkAccessToken(v string) *SaveCorpPayCodeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveCorpPayCodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// H_FISH_CANTEEN
	CodeIdentity *string `json:"codeIdentity,omitempty" xml:"codeIdentity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dingxxxx
	CorpId  *string            `json:"corpId,omitempty" xml:"corpId,omitempty"`
	ExtInfo map[string]*string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SaveCorpPayCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveCorpPayCodeRequest) GoString() string {
	return s.String()
}

func (s *SaveCorpPayCodeRequest) SetCodeIdentity(v string) *SaveCorpPayCodeRequest {
	s.CodeIdentity = &v
	return s
}

func (s *SaveCorpPayCodeRequest) SetCorpId(v string) *SaveCorpPayCodeRequest {
	s.CorpId = &v
	return s
}

func (s *SaveCorpPayCodeRequest) SetExtInfo(v map[string]*string) *SaveCorpPayCodeRequest {
	s.ExtInfo = v
	return s
}

func (s *SaveCorpPayCodeRequest) SetStatus(v string) *SaveCorpPayCodeRequest {
	s.Status = &v
	return s
}

type SaveCorpPayCodeResponseBody struct {
	CodeIdentity *string            `json:"codeIdentity,omitempty" xml:"codeIdentity,omitempty"`
	CorpId       *string            `json:"corpId,omitempty" xml:"corpId,omitempty"`
	ExtInfo      map[string]*string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	Status       *string            `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SaveCorpPayCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveCorpPayCodeResponseBody) GoString() string {
	return s.String()
}

func (s *SaveCorpPayCodeResponseBody) SetCodeIdentity(v string) *SaveCorpPayCodeResponseBody {
	s.CodeIdentity = &v
	return s
}

func (s *SaveCorpPayCodeResponseBody) SetCorpId(v string) *SaveCorpPayCodeResponseBody {
	s.CorpId = &v
	return s
}

func (s *SaveCorpPayCodeResponseBody) SetExtInfo(v map[string]*string) *SaveCorpPayCodeResponseBody {
	s.ExtInfo = v
	return s
}

func (s *SaveCorpPayCodeResponseBody) SetStatus(v string) *SaveCorpPayCodeResponseBody {
	s.Status = &v
	return s
}

type SaveCorpPayCodeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveCorpPayCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveCorpPayCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveCorpPayCodeResponse) GoString() string {
	return s.String()
}

func (s *SaveCorpPayCodeResponse) SetHeaders(v map[string]*string) *SaveCorpPayCodeResponse {
	s.Headers = v
	return s
}

func (s *SaveCorpPayCodeResponse) SetStatusCode(v int32) *SaveCorpPayCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveCorpPayCodeResponse) SetBody(v *SaveCorpPayCodeResponseBody) *SaveCorpPayCodeResponse {
	s.Body = v
	return s
}

type UnsignUserAgreementHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UnsignUserAgreementHeaders) String() string {
	return tea.Prettify(s)
}

func (s UnsignUserAgreementHeaders) GoString() string {
	return s.String()
}

func (s *UnsignUserAgreementHeaders) SetCommonHeaders(v map[string]*string) *UnsignUserAgreementHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UnsignUserAgreementHeaders) SetXAcsDingtalkAccessToken(v string) *UnsignUserAgreementHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UnsignUserAgreementRequest struct {
	// example:
	//
	// 23021_12342134
	AgreementNo *string `json:"agreementNo,omitempty" xml:"agreementNo,omitempty"`
	// example:
	//
	// TRADE
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	// example:
	//
	// WITHHOLDING
	BizScene *string `json:"bizScene,omitempty" xml:"bizScene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 202111090001
	InstId *string `json:"instId,omitempty" xml:"instId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1001
	SubInstId *string `json:"subInstId,omitempty" xml:"subInstId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2120493284
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UnsignUserAgreementRequest) String() string {
	return tea.Prettify(s)
}

func (s UnsignUserAgreementRequest) GoString() string {
	return s.String()
}

func (s *UnsignUserAgreementRequest) SetAgreementNo(v string) *UnsignUserAgreementRequest {
	s.AgreementNo = &v
	return s
}

func (s *UnsignUserAgreementRequest) SetBizCode(v string) *UnsignUserAgreementRequest {
	s.BizCode = &v
	return s
}

func (s *UnsignUserAgreementRequest) SetBizScene(v string) *UnsignUserAgreementRequest {
	s.BizScene = &v
	return s
}

func (s *UnsignUserAgreementRequest) SetInstId(v string) *UnsignUserAgreementRequest {
	s.InstId = &v
	return s
}

func (s *UnsignUserAgreementRequest) SetSubInstId(v string) *UnsignUserAgreementRequest {
	s.SubInstId = &v
	return s
}

func (s *UnsignUserAgreementRequest) SetUserId(v string) *UnsignUserAgreementRequest {
	s.UserId = &v
	return s
}

type UnsignUserAgreementResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UnsignUserAgreementResponse) String() string {
	return tea.Prettify(s)
}

func (s UnsignUserAgreementResponse) GoString() string {
	return s.String()
}

func (s *UnsignUserAgreementResponse) SetHeaders(v map[string]*string) *UnsignUserAgreementResponse {
	s.Headers = v
	return s
}

func (s *UnsignUserAgreementResponse) SetStatusCode(v int32) *UnsignUserAgreementResponse {
	s.StatusCode = &v
	return s
}

type UpateUserCodeInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpateUserCodeInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpateUserCodeInstanceHeaders) GoString() string {
	return s.String()
}

func (s *UpateUserCodeInstanceHeaders) SetCommonHeaders(v map[string]*string) *UpateUserCodeInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpateUserCodeInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *UpateUserCodeInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpateUserCodeInstanceRequest struct {
	// This parameter is required.
	AvailableTimes []*UpateUserCodeInstanceRequestAvailableTimes `json:"availableTimes,omitempty" xml:"availableTimes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// ccodexxxxx
	CodeId *string `json:"codeId,omitempty" xml:"codeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TEST
	CodeIdentity *string `json:"codeIdentity,omitempty" xml:"codeIdentity,omitempty"`
	CodeValue    *string `json:"codeValue,omitempty" xml:"codeValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// corpIdxxxx
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// This parameter is required.
	ExtInfo map[string]interface{} `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// This parameter is required.
	GmtExpired *string `json:"gmtExpired,omitempty" xml:"gmtExpired,omitempty"`
	// example:
	//
	// OPEN
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// INTERNAL_STAFF
	UserCorpRelationType *string `json:"userCorpRelationType,omitempty" xml:"userCorpRelationType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 86-xxxxxx
	UserIdentity *string `json:"userIdentity,omitempty" xml:"userIdentity,omitempty"`
}

func (s UpateUserCodeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpateUserCodeInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpateUserCodeInstanceRequest) SetAvailableTimes(v []*UpateUserCodeInstanceRequestAvailableTimes) *UpateUserCodeInstanceRequest {
	s.AvailableTimes = v
	return s
}

func (s *UpateUserCodeInstanceRequest) SetCodeId(v string) *UpateUserCodeInstanceRequest {
	s.CodeId = &v
	return s
}

func (s *UpateUserCodeInstanceRequest) SetCodeIdentity(v string) *UpateUserCodeInstanceRequest {
	s.CodeIdentity = &v
	return s
}

func (s *UpateUserCodeInstanceRequest) SetCodeValue(v string) *UpateUserCodeInstanceRequest {
	s.CodeValue = &v
	return s
}

func (s *UpateUserCodeInstanceRequest) SetCorpId(v string) *UpateUserCodeInstanceRequest {
	s.CorpId = &v
	return s
}

func (s *UpateUserCodeInstanceRequest) SetExtInfo(v map[string]interface{}) *UpateUserCodeInstanceRequest {
	s.ExtInfo = v
	return s
}

func (s *UpateUserCodeInstanceRequest) SetGmtExpired(v string) *UpateUserCodeInstanceRequest {
	s.GmtExpired = &v
	return s
}

func (s *UpateUserCodeInstanceRequest) SetStatus(v string) *UpateUserCodeInstanceRequest {
	s.Status = &v
	return s
}

func (s *UpateUserCodeInstanceRequest) SetUserCorpRelationType(v string) *UpateUserCodeInstanceRequest {
	s.UserCorpRelationType = &v
	return s
}

func (s *UpateUserCodeInstanceRequest) SetUserIdentity(v string) *UpateUserCodeInstanceRequest {
	s.UserIdentity = &v
	return s
}

type UpateUserCodeInstanceRequestAvailableTimes struct {
	// This parameter is required.
	//
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	GmtEnd *string `json:"gmtEnd,omitempty" xml:"gmtEnd,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	GmtStart *string `json:"gmtStart,omitempty" xml:"gmtStart,omitempty"`
}

func (s UpateUserCodeInstanceRequestAvailableTimes) String() string {
	return tea.Prettify(s)
}

func (s UpateUserCodeInstanceRequestAvailableTimes) GoString() string {
	return s.String()
}

func (s *UpateUserCodeInstanceRequestAvailableTimes) SetGmtEnd(v string) *UpateUserCodeInstanceRequestAvailableTimes {
	s.GmtEnd = &v
	return s
}

func (s *UpateUserCodeInstanceRequestAvailableTimes) SetGmtStart(v string) *UpateUserCodeInstanceRequestAvailableTimes {
	s.GmtStart = &v
	return s
}

type UpateUserCodeInstanceResponseBody struct {
	// example:
	//
	// codexxxxxx
	CodeId *string `json:"codeId,omitempty" xml:"codeId,omitempty"`
}

func (s UpateUserCodeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpateUserCodeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpateUserCodeInstanceResponseBody) SetCodeId(v string) *UpateUserCodeInstanceResponseBody {
	s.CodeId = &v
	return s
}

type UpateUserCodeInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpateUserCodeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpateUserCodeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpateUserCodeInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpateUserCodeInstanceResponse) SetHeaders(v map[string]*string) *UpateUserCodeInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpateUserCodeInstanceResponse) SetStatusCode(v int32) *UpateUserCodeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpateUserCodeInstanceResponse) SetBody(v *UpateUserCodeInstanceResponseBody) *UpateUserCodeInstanceResponse {
	s.Body = v
	return s
}

type UpdateInvoiceVerifyStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateInvoiceVerifyStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceVerifyStatusHeaders) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceVerifyStatusHeaders) SetCommonHeaders(v map[string]*string) *UpdateInvoiceVerifyStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateInvoiceVerifyStatusHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateInvoiceVerifyStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateInvoiceVerifyStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// lpKgSTzGSy
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// example:
	//
	// 1
	CheckingResult *int32 `json:"checkingResult,omitempty" xml:"checkingResult,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CheckingStatus *int32 `json:"checkingStatus,omitempty" xml:"checkingStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ding673cxxxxxxxxxxxx85
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// {"restCheckTimes":10,"noticeFlag":1}
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// example:
	//
	// 034012100111
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 61235725
	InvoiceNo *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
	// example:
	//
	// 1
	InvoiceStatus *int32 `json:"invoiceStatus,omitempty" xml:"invoiceStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1299999
	InvoiceVerifyId *string `json:"invoiceVerifyId,omitempty" xml:"invoiceVerifyId,omitempty"`
	// This parameter is required.
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BPq7qiSIH8PJHlB9kPuii1NQiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s UpdateInvoiceVerifyStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceVerifyStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceVerifyStatusRequest) SetBizId(v string) *UpdateInvoiceVerifyStatusRequest {
	s.BizId = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequest) SetCheckingResult(v int32) *UpdateInvoiceVerifyStatusRequest {
	s.CheckingResult = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequest) SetCheckingStatus(v int32) *UpdateInvoiceVerifyStatusRequest {
	s.CheckingStatus = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequest) SetCode(v string) *UpdateInvoiceVerifyStatusRequest {
	s.Code = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequest) SetCorpId(v string) *UpdateInvoiceVerifyStatusRequest {
	s.CorpId = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequest) SetExtension(v string) *UpdateInvoiceVerifyStatusRequest {
	s.Extension = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequest) SetInvoiceCode(v string) *UpdateInvoiceVerifyStatusRequest {
	s.InvoiceCode = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequest) SetInvoiceNo(v string) *UpdateInvoiceVerifyStatusRequest {
	s.InvoiceNo = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequest) SetInvoiceStatus(v int32) *UpdateInvoiceVerifyStatusRequest {
	s.InvoiceStatus = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequest) SetInvoiceVerifyId(v string) *UpdateInvoiceVerifyStatusRequest {
	s.InvoiceVerifyId = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequest) SetMsg(v string) *UpdateInvoiceVerifyStatusRequest {
	s.Msg = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequest) SetUnionId(v string) *UpdateInvoiceVerifyStatusRequest {
	s.UnionId = &v
	return s
}

type UpdateInvoiceVerifyStatusResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateInvoiceVerifyStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceVerifyStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceVerifyStatusResponseBody) SetResult(v bool) *UpdateInvoiceVerifyStatusResponseBody {
	s.Result = &v
	return s
}

type UpdateInvoiceVerifyStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInvoiceVerifyStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInvoiceVerifyStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceVerifyStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceVerifyStatusResponse) SetHeaders(v map[string]*string) *UpdateInvoiceVerifyStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateInvoiceVerifyStatusResponse) SetStatusCode(v int32) *UpdateInvoiceVerifyStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusResponse) SetBody(v *UpdateInvoiceVerifyStatusResponseBody) *UpdateInvoiceVerifyStatusResponse {
	s.Body = v
	return s
}

type UploadInvoiceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UploadInvoiceHeaders) String() string {
	return tea.Prettify(s)
}

func (s UploadInvoiceHeaders) GoString() string {
	return s.String()
}

func (s *UploadInvoiceHeaders) SetCommonHeaders(v map[string]*string) *UploadInvoiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UploadInvoiceHeaders) SetXAcsDingtalkAccessToken(v string) *UploadInvoiceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UploadInvoiceRequest struct {
	Extension *UploadInvoiceRequestExtension `json:"extension,omitempty" xml:"extension,omitempty" type:"Struct"`
	// This parameter is required.
	Invoices     []*UploadInvoiceRequestInvoices   `json:"invoices,omitempty" xml:"invoices,omitempty" type:"Repeated"`
	UserIdentity *UploadInvoiceRequestUserIdentity `json:"userIdentity,omitempty" xml:"userIdentity,omitempty" type:"Struct"`
}

func (s UploadInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadInvoiceRequest) GoString() string {
	return s.String()
}

func (s *UploadInvoiceRequest) SetExtension(v *UploadInvoiceRequestExtension) *UploadInvoiceRequest {
	s.Extension = v
	return s
}

func (s *UploadInvoiceRequest) SetInvoices(v []*UploadInvoiceRequestInvoices) *UploadInvoiceRequest {
	s.Invoices = v
	return s
}

func (s *UploadInvoiceRequest) SetUserIdentity(v *UploadInvoiceRequestUserIdentity) *UploadInvoiceRequest {
	s.UserIdentity = v
	return s
}

type UploadInvoiceRequestExtension struct {
	// This parameter is required.
	//
	// example:
	//
	// TRIP
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	// Deprecated
	//
	// example:
	//
	// 111924191922
	OrderNo     *string   `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	OrderNoList []*string `json:"orderNoList,omitempty" xml:"orderNoList,omitempty" type:"Repeated"`
	// example:
	//
	// HOTEL
	OrderType *string `json:"orderType,omitempty" xml:"orderType,omitempty"`
}

func (s UploadInvoiceRequestExtension) String() string {
	return tea.Prettify(s)
}

func (s UploadInvoiceRequestExtension) GoString() string {
	return s.String()
}

func (s *UploadInvoiceRequestExtension) SetBizCode(v string) *UploadInvoiceRequestExtension {
	s.BizCode = &v
	return s
}

func (s *UploadInvoiceRequestExtension) SetOrderNo(v string) *UploadInvoiceRequestExtension {
	s.OrderNo = &v
	return s
}

func (s *UploadInvoiceRequestExtension) SetOrderNoList(v []*string) *UploadInvoiceRequestExtension {
	s.OrderNoList = v
	return s
}

func (s *UploadInvoiceRequestExtension) SetOrderType(v string) *UploadInvoiceRequestExtension {
	s.OrderType = &v
	return s
}

type UploadInvoiceRequestInvoices struct {
	// This parameter is required.
	//
	// example:
	//
	// 100.00
	InvoiceAmount *string `json:"invoiceAmount,omitempty" xml:"invoiceAmount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 033002000712
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-02-21
	InvoiceDate *string `json:"invoiceDate,omitempty" xml:"invoiceDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20532643
	InvoiceNo *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
	// This parameter is required.
	InvoiceType *string `json:"invoiceType,omitempty" xml:"invoiceType,omitempty"`
	LogoUrl     *string `json:"logoUrl,omitempty" xml:"logoUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 小钉科技有限公司
	PayeeName *string `json:"payeeName,omitempty" xml:"payeeName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 91330100MA28XNB274
	PayeeTaxNo *string `json:"payeeTaxNo,omitempty" xml:"payeeTaxNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 小钉科技有限公司
	PayerName *string `json:"payerName,omitempty" xml:"payerName,omitempty"`
	// example:
	//
	// 91330100MA28XNB274
	PayerTaxNo *string `json:"payerTaxNo,omitempty" xml:"payerTaxNo,omitempty"`
	// This parameter is required.
	PdfUrl *string `json:"pdfUrl,omitempty" xml:"pdfUrl,omitempty"`
	// example:
	//
	// 0.50
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	// example:
	//
	// 增值税普通发票必填，示例：52501101414266612380
	VerifyCode *string `json:"verifyCode,omitempty" xml:"verifyCode,omitempty"`
	// example:
	//
	// 99.50
	WithoutTaxAmount *string `json:"withoutTaxAmount,omitempty" xml:"withoutTaxAmount,omitempty"`
}

func (s UploadInvoiceRequestInvoices) String() string {
	return tea.Prettify(s)
}

func (s UploadInvoiceRequestInvoices) GoString() string {
	return s.String()
}

func (s *UploadInvoiceRequestInvoices) SetInvoiceAmount(v string) *UploadInvoiceRequestInvoices {
	s.InvoiceAmount = &v
	return s
}

func (s *UploadInvoiceRequestInvoices) SetInvoiceCode(v string) *UploadInvoiceRequestInvoices {
	s.InvoiceCode = &v
	return s
}

func (s *UploadInvoiceRequestInvoices) SetInvoiceDate(v string) *UploadInvoiceRequestInvoices {
	s.InvoiceDate = &v
	return s
}

func (s *UploadInvoiceRequestInvoices) SetInvoiceNo(v string) *UploadInvoiceRequestInvoices {
	s.InvoiceNo = &v
	return s
}

func (s *UploadInvoiceRequestInvoices) SetInvoiceType(v string) *UploadInvoiceRequestInvoices {
	s.InvoiceType = &v
	return s
}

func (s *UploadInvoiceRequestInvoices) SetLogoUrl(v string) *UploadInvoiceRequestInvoices {
	s.LogoUrl = &v
	return s
}

func (s *UploadInvoiceRequestInvoices) SetPayeeName(v string) *UploadInvoiceRequestInvoices {
	s.PayeeName = &v
	return s
}

func (s *UploadInvoiceRequestInvoices) SetPayeeTaxNo(v string) *UploadInvoiceRequestInvoices {
	s.PayeeTaxNo = &v
	return s
}

func (s *UploadInvoiceRequestInvoices) SetPayerName(v string) *UploadInvoiceRequestInvoices {
	s.PayerName = &v
	return s
}

func (s *UploadInvoiceRequestInvoices) SetPayerTaxNo(v string) *UploadInvoiceRequestInvoices {
	s.PayerTaxNo = &v
	return s
}

func (s *UploadInvoiceRequestInvoices) SetPdfUrl(v string) *UploadInvoiceRequestInvoices {
	s.PdfUrl = &v
	return s
}

func (s *UploadInvoiceRequestInvoices) SetTaxAmount(v string) *UploadInvoiceRequestInvoices {
	s.TaxAmount = &v
	return s
}

func (s *UploadInvoiceRequestInvoices) SetVerifyCode(v string) *UploadInvoiceRequestInvoices {
	s.VerifyCode = &v
	return s
}

func (s *UploadInvoiceRequestInvoices) SetWithoutTaxAmount(v string) *UploadInvoiceRequestInvoices {
	s.WithoutTaxAmount = &v
	return s
}

type UploadInvoiceRequestUserIdentity struct {
	// example:
	//
	// 95188
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// example:
	//
	// 86
	MobileStateCode *string `json:"mobileStateCode,omitempty" xml:"mobileStateCode,omitempty"`
	// example:
	//
	// dinng1123434
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// STAFF_ID
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// akdfwiiw
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// example:
	//
	// 02734930
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UploadInvoiceRequestUserIdentity) String() string {
	return tea.Prettify(s)
}

func (s UploadInvoiceRequestUserIdentity) GoString() string {
	return s.String()
}

func (s *UploadInvoiceRequestUserIdentity) SetMobile(v string) *UploadInvoiceRequestUserIdentity {
	s.Mobile = &v
	return s
}

func (s *UploadInvoiceRequestUserIdentity) SetMobileStateCode(v string) *UploadInvoiceRequestUserIdentity {
	s.MobileStateCode = &v
	return s
}

func (s *UploadInvoiceRequestUserIdentity) SetTargetCorpId(v string) *UploadInvoiceRequestUserIdentity {
	s.TargetCorpId = &v
	return s
}

func (s *UploadInvoiceRequestUserIdentity) SetType(v string) *UploadInvoiceRequestUserIdentity {
	s.Type = &v
	return s
}

func (s *UploadInvoiceRequestUserIdentity) SetUnionId(v string) *UploadInvoiceRequestUserIdentity {
	s.UnionId = &v
	return s
}

func (s *UploadInvoiceRequestUserIdentity) SetUserId(v string) *UploadInvoiceRequestUserIdentity {
	s.UserId = &v
	return s
}

type UploadInvoiceResponseBody struct {
	Result *UploadInvoiceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UploadInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *UploadInvoiceResponseBody) SetResult(v *UploadInvoiceResponseBodyResult) *UploadInvoiceResponseBody {
	s.Result = v
	return s
}

type UploadInvoiceResponseBodyResult struct {
	// This parameter is required.
	Results []*UploadInvoiceResponseBodyResultResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s UploadInvoiceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UploadInvoiceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UploadInvoiceResponseBodyResult) SetResults(v []*UploadInvoiceResponseBodyResultResults) *UploadInvoiceResponseBodyResult {
	s.Results = v
	return s
}

type UploadInvoiceResponseBodyResultResults struct {
	// This parameter is required.
	//
	// example:
	//
	// 20006
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 033002000712
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20532643
	InvoiceNo *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// duplicateInvoice
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UploadInvoiceResponseBodyResultResults) String() string {
	return tea.Prettify(s)
}

func (s UploadInvoiceResponseBodyResultResults) GoString() string {
	return s.String()
}

func (s *UploadInvoiceResponseBodyResultResults) SetErrCode(v string) *UploadInvoiceResponseBodyResultResults {
	s.ErrCode = &v
	return s
}

func (s *UploadInvoiceResponseBodyResultResults) SetInvoiceCode(v string) *UploadInvoiceResponseBodyResultResults {
	s.InvoiceCode = &v
	return s
}

func (s *UploadInvoiceResponseBodyResultResults) SetInvoiceNo(v string) *UploadInvoiceResponseBodyResultResults {
	s.InvoiceNo = &v
	return s
}

func (s *UploadInvoiceResponseBodyResultResults) SetReason(v string) *UploadInvoiceResponseBodyResultResults {
	s.Reason = &v
	return s
}

func (s *UploadInvoiceResponseBodyResultResults) SetSuccess(v bool) *UploadInvoiceResponseBodyResultResults {
	s.Success = &v
	return s
}

type UploadInvoiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadInvoiceResponse) GoString() string {
	return s.String()
}

func (s *UploadInvoiceResponse) SetHeaders(v map[string]*string) *UploadInvoiceResponse {
	s.Headers = v
	return s
}

func (s *UploadInvoiceResponse) SetStatusCode(v int32) *UploadInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadInvoiceResponse) SetBody(v *UploadInvoiceResponseBody) *UploadInvoiceResponse {
	s.Body = v
	return s
}

type UploadInvoiceByAuthHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UploadInvoiceByAuthHeaders) String() string {
	return tea.Prettify(s)
}

func (s UploadInvoiceByAuthHeaders) GoString() string {
	return s.String()
}

func (s *UploadInvoiceByAuthHeaders) SetCommonHeaders(v map[string]*string) *UploadInvoiceByAuthHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UploadInvoiceByAuthHeaders) SetXAcsDingtalkAccessToken(v string) *UploadInvoiceByAuthHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UploadInvoiceByAuthRequest struct {
	Extension *UploadInvoiceByAuthRequestExtension `json:"extension,omitempty" xml:"extension,omitempty" type:"Struct"`
	// This parameter is required.
	Invoices []*UploadInvoiceByAuthRequestInvoices `json:"invoices,omitempty" xml:"invoices,omitempty" type:"Repeated"`
}

func (s UploadInvoiceByAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadInvoiceByAuthRequest) GoString() string {
	return s.String()
}

func (s *UploadInvoiceByAuthRequest) SetExtension(v *UploadInvoiceByAuthRequestExtension) *UploadInvoiceByAuthRequest {
	s.Extension = v
	return s
}

func (s *UploadInvoiceByAuthRequest) SetInvoices(v []*UploadInvoiceByAuthRequestInvoices) *UploadInvoiceByAuthRequest {
	s.Invoices = v
	return s
}

type UploadInvoiceByAuthRequestExtension struct {
	// This parameter is required.
	//
	// example:
	//
	// TRIP
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	// example:
	//
	// 111924191922
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	// example:
	//
	// HOTEL
	OrderType *string `json:"orderType,omitempty" xml:"orderType,omitempty"`
}

func (s UploadInvoiceByAuthRequestExtension) String() string {
	return tea.Prettify(s)
}

func (s UploadInvoiceByAuthRequestExtension) GoString() string {
	return s.String()
}

func (s *UploadInvoiceByAuthRequestExtension) SetBizCode(v string) *UploadInvoiceByAuthRequestExtension {
	s.BizCode = &v
	return s
}

func (s *UploadInvoiceByAuthRequestExtension) SetOrderNo(v string) *UploadInvoiceByAuthRequestExtension {
	s.OrderNo = &v
	return s
}

func (s *UploadInvoiceByAuthRequestExtension) SetOrderType(v string) *UploadInvoiceByAuthRequestExtension {
	s.OrderType = &v
	return s
}

type UploadInvoiceByAuthRequestInvoices struct {
	// This parameter is required.
	//
	// example:
	//
	// 100.00
	InvoiceAmount *string `json:"invoiceAmount,omitempty" xml:"invoiceAmount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 033002000712
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-02-21
	InvoiceDate *string `json:"invoiceDate,omitempty" xml:"invoiceDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20532643
	InvoiceNo *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
	// This parameter is required.
	InvoiceType *string `json:"invoiceType,omitempty" xml:"invoiceType,omitempty"`
	LogoUrl     *string `json:"logoUrl,omitempty" xml:"logoUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 小钉科技有限公司
	PayeeName *string `json:"payeeName,omitempty" xml:"payeeName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 91330100MA28XNB274
	PayeeTaxNo *string `json:"payeeTaxNo,omitempty" xml:"payeeTaxNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 小钉科技有限公司
	PayerName *string `json:"payerName,omitempty" xml:"payerName,omitempty"`
	// example:
	//
	// 91330100MA28XNB274
	PayerTaxNo *string `json:"payerTaxNo,omitempty" xml:"payerTaxNo,omitempty"`
	// This parameter is required.
	PdfUrl *string `json:"pdfUrl,omitempty" xml:"pdfUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.50
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	// example:
	//
	// 增值税普通发票必填，示例：52501101414266612380
	VerifyCode *string `json:"verifyCode,omitempty" xml:"verifyCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 99.50
	WithoutTaxAmount *string `json:"withoutTaxAmount,omitempty" xml:"withoutTaxAmount,omitempty"`
}

func (s UploadInvoiceByAuthRequestInvoices) String() string {
	return tea.Prettify(s)
}

func (s UploadInvoiceByAuthRequestInvoices) GoString() string {
	return s.String()
}

func (s *UploadInvoiceByAuthRequestInvoices) SetInvoiceAmount(v string) *UploadInvoiceByAuthRequestInvoices {
	s.InvoiceAmount = &v
	return s
}

func (s *UploadInvoiceByAuthRequestInvoices) SetInvoiceCode(v string) *UploadInvoiceByAuthRequestInvoices {
	s.InvoiceCode = &v
	return s
}

func (s *UploadInvoiceByAuthRequestInvoices) SetInvoiceDate(v string) *UploadInvoiceByAuthRequestInvoices {
	s.InvoiceDate = &v
	return s
}

func (s *UploadInvoiceByAuthRequestInvoices) SetInvoiceNo(v string) *UploadInvoiceByAuthRequestInvoices {
	s.InvoiceNo = &v
	return s
}

func (s *UploadInvoiceByAuthRequestInvoices) SetInvoiceType(v string) *UploadInvoiceByAuthRequestInvoices {
	s.InvoiceType = &v
	return s
}

func (s *UploadInvoiceByAuthRequestInvoices) SetLogoUrl(v string) *UploadInvoiceByAuthRequestInvoices {
	s.LogoUrl = &v
	return s
}

func (s *UploadInvoiceByAuthRequestInvoices) SetPayeeName(v string) *UploadInvoiceByAuthRequestInvoices {
	s.PayeeName = &v
	return s
}

func (s *UploadInvoiceByAuthRequestInvoices) SetPayeeTaxNo(v string) *UploadInvoiceByAuthRequestInvoices {
	s.PayeeTaxNo = &v
	return s
}

func (s *UploadInvoiceByAuthRequestInvoices) SetPayerName(v string) *UploadInvoiceByAuthRequestInvoices {
	s.PayerName = &v
	return s
}

func (s *UploadInvoiceByAuthRequestInvoices) SetPayerTaxNo(v string) *UploadInvoiceByAuthRequestInvoices {
	s.PayerTaxNo = &v
	return s
}

func (s *UploadInvoiceByAuthRequestInvoices) SetPdfUrl(v string) *UploadInvoiceByAuthRequestInvoices {
	s.PdfUrl = &v
	return s
}

func (s *UploadInvoiceByAuthRequestInvoices) SetTaxAmount(v string) *UploadInvoiceByAuthRequestInvoices {
	s.TaxAmount = &v
	return s
}

func (s *UploadInvoiceByAuthRequestInvoices) SetVerifyCode(v string) *UploadInvoiceByAuthRequestInvoices {
	s.VerifyCode = &v
	return s
}

func (s *UploadInvoiceByAuthRequestInvoices) SetWithoutTaxAmount(v string) *UploadInvoiceByAuthRequestInvoices {
	s.WithoutTaxAmount = &v
	return s
}

type UploadInvoiceByAuthResponseBody struct {
	Result *UploadInvoiceByAuthResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UploadInvoiceByAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadInvoiceByAuthResponseBody) GoString() string {
	return s.String()
}

func (s *UploadInvoiceByAuthResponseBody) SetResult(v *UploadInvoiceByAuthResponseBodyResult) *UploadInvoiceByAuthResponseBody {
	s.Result = v
	return s
}

type UploadInvoiceByAuthResponseBodyResult struct {
	// This parameter is required.
	Results []*UploadInvoiceByAuthResponseBodyResultResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s UploadInvoiceByAuthResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UploadInvoiceByAuthResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UploadInvoiceByAuthResponseBodyResult) SetResults(v []*UploadInvoiceByAuthResponseBodyResultResults) *UploadInvoiceByAuthResponseBodyResult {
	s.Results = v
	return s
}

type UploadInvoiceByAuthResponseBodyResultResults struct {
	// This parameter is required.
	//
	// example:
	//
	// 20006
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 033002000712
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20532643
	InvoiceNo *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// duplicateInvoice
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UploadInvoiceByAuthResponseBodyResultResults) String() string {
	return tea.Prettify(s)
}

func (s UploadInvoiceByAuthResponseBodyResultResults) GoString() string {
	return s.String()
}

func (s *UploadInvoiceByAuthResponseBodyResultResults) SetErrCode(v string) *UploadInvoiceByAuthResponseBodyResultResults {
	s.ErrCode = &v
	return s
}

func (s *UploadInvoiceByAuthResponseBodyResultResults) SetInvoiceCode(v string) *UploadInvoiceByAuthResponseBodyResultResults {
	s.InvoiceCode = &v
	return s
}

func (s *UploadInvoiceByAuthResponseBodyResultResults) SetInvoiceNo(v string) *UploadInvoiceByAuthResponseBodyResultResults {
	s.InvoiceNo = &v
	return s
}

func (s *UploadInvoiceByAuthResponseBodyResultResults) SetReason(v string) *UploadInvoiceByAuthResponseBodyResultResults {
	s.Reason = &v
	return s
}

func (s *UploadInvoiceByAuthResponseBodyResultResults) SetSuccess(v bool) *UploadInvoiceByAuthResponseBodyResultResults {
	s.Success = &v
	return s
}

type UploadInvoiceByAuthResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadInvoiceByAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadInvoiceByAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadInvoiceByAuthResponse) GoString() string {
	return s.String()
}

func (s *UploadInvoiceByAuthResponse) SetHeaders(v map[string]*string) *UploadInvoiceByAuthResponse {
	s.Headers = v
	return s
}

func (s *UploadInvoiceByAuthResponse) SetStatusCode(v int32) *UploadInvoiceByAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadInvoiceByAuthResponse) SetBody(v *UploadInvoiceByAuthResponseBody) *UploadInvoiceByAuthResponse {
	s.Body = v
	return s
}

type UploadInvoiceByMobileHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UploadInvoiceByMobileHeaders) String() string {
	return tea.Prettify(s)
}

func (s UploadInvoiceByMobileHeaders) GoString() string {
	return s.String()
}

func (s *UploadInvoiceByMobileHeaders) SetCommonHeaders(v map[string]*string) *UploadInvoiceByMobileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UploadInvoiceByMobileHeaders) SetXAcsDingtalkAccessToken(v string) *UploadInvoiceByMobileHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UploadInvoiceByMobileRequest struct {
	// This parameter is required.
	Invoices []*UploadInvoiceByMobileRequestInvoices `json:"invoices,omitempty" xml:"invoices,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 13600000000
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 86
	MobileStateCode *string `json:"mobileStateCode,omitempty" xml:"mobileStateCode,omitempty"`
}

func (s UploadInvoiceByMobileRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadInvoiceByMobileRequest) GoString() string {
	return s.String()
}

func (s *UploadInvoiceByMobileRequest) SetInvoices(v []*UploadInvoiceByMobileRequestInvoices) *UploadInvoiceByMobileRequest {
	s.Invoices = v
	return s
}

func (s *UploadInvoiceByMobileRequest) SetMobile(v string) *UploadInvoiceByMobileRequest {
	s.Mobile = &v
	return s
}

func (s *UploadInvoiceByMobileRequest) SetMobileStateCode(v string) *UploadInvoiceByMobileRequest {
	s.MobileStateCode = &v
	return s
}

type UploadInvoiceByMobileRequestInvoices struct {
	// This parameter is required.
	//
	// example:
	//
	// 100.00
	InvoiceAmount *string `json:"invoiceAmount,omitempty" xml:"invoiceAmount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 033002000712
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-02-21
	InvoiceDate *string `json:"invoiceDate,omitempty" xml:"invoiceDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20532643
	InvoiceNo *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
	// This parameter is required.
	InvoiceType *string `json:"invoiceType,omitempty" xml:"invoiceType,omitempty"`
	LogoUrl     *string `json:"logoUrl,omitempty" xml:"logoUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 小钉科技有限公司
	PayeeName *string `json:"payeeName,omitempty" xml:"payeeName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 91330100MA28XNB274
	PayeeTaxNo *string `json:"payeeTaxNo,omitempty" xml:"payeeTaxNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 小钉科技有限公司
	PayerName *string `json:"payerName,omitempty" xml:"payerName,omitempty"`
	// example:
	//
	// 91330100MA28XNB274
	PayerTaxNo *string `json:"payerTaxNo,omitempty" xml:"payerTaxNo,omitempty"`
	// This parameter is required.
	PdfUrl *string `json:"pdfUrl,omitempty" xml:"pdfUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.50
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	// example:
	//
	// 增值税普通发票必填，示例：52501101414266612380
	VerifyCode *string `json:"verifyCode,omitempty" xml:"verifyCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 99.50
	WithoutTaxAmount *string `json:"withoutTaxAmount,omitempty" xml:"withoutTaxAmount,omitempty"`
}

func (s UploadInvoiceByMobileRequestInvoices) String() string {
	return tea.Prettify(s)
}

func (s UploadInvoiceByMobileRequestInvoices) GoString() string {
	return s.String()
}

func (s *UploadInvoiceByMobileRequestInvoices) SetInvoiceAmount(v string) *UploadInvoiceByMobileRequestInvoices {
	s.InvoiceAmount = &v
	return s
}

func (s *UploadInvoiceByMobileRequestInvoices) SetInvoiceCode(v string) *UploadInvoiceByMobileRequestInvoices {
	s.InvoiceCode = &v
	return s
}

func (s *UploadInvoiceByMobileRequestInvoices) SetInvoiceDate(v string) *UploadInvoiceByMobileRequestInvoices {
	s.InvoiceDate = &v
	return s
}

func (s *UploadInvoiceByMobileRequestInvoices) SetInvoiceNo(v string) *UploadInvoiceByMobileRequestInvoices {
	s.InvoiceNo = &v
	return s
}

func (s *UploadInvoiceByMobileRequestInvoices) SetInvoiceType(v string) *UploadInvoiceByMobileRequestInvoices {
	s.InvoiceType = &v
	return s
}

func (s *UploadInvoiceByMobileRequestInvoices) SetLogoUrl(v string) *UploadInvoiceByMobileRequestInvoices {
	s.LogoUrl = &v
	return s
}

func (s *UploadInvoiceByMobileRequestInvoices) SetPayeeName(v string) *UploadInvoiceByMobileRequestInvoices {
	s.PayeeName = &v
	return s
}

func (s *UploadInvoiceByMobileRequestInvoices) SetPayeeTaxNo(v string) *UploadInvoiceByMobileRequestInvoices {
	s.PayeeTaxNo = &v
	return s
}

func (s *UploadInvoiceByMobileRequestInvoices) SetPayerName(v string) *UploadInvoiceByMobileRequestInvoices {
	s.PayerName = &v
	return s
}

func (s *UploadInvoiceByMobileRequestInvoices) SetPayerTaxNo(v string) *UploadInvoiceByMobileRequestInvoices {
	s.PayerTaxNo = &v
	return s
}

func (s *UploadInvoiceByMobileRequestInvoices) SetPdfUrl(v string) *UploadInvoiceByMobileRequestInvoices {
	s.PdfUrl = &v
	return s
}

func (s *UploadInvoiceByMobileRequestInvoices) SetTaxAmount(v string) *UploadInvoiceByMobileRequestInvoices {
	s.TaxAmount = &v
	return s
}

func (s *UploadInvoiceByMobileRequestInvoices) SetVerifyCode(v string) *UploadInvoiceByMobileRequestInvoices {
	s.VerifyCode = &v
	return s
}

func (s *UploadInvoiceByMobileRequestInvoices) SetWithoutTaxAmount(v string) *UploadInvoiceByMobileRequestInvoices {
	s.WithoutTaxAmount = &v
	return s
}

type UploadInvoiceByMobileResponseBody struct {
	Result *UploadInvoiceByMobileResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UploadInvoiceByMobileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadInvoiceByMobileResponseBody) GoString() string {
	return s.String()
}

func (s *UploadInvoiceByMobileResponseBody) SetResult(v *UploadInvoiceByMobileResponseBodyResult) *UploadInvoiceByMobileResponseBody {
	s.Result = v
	return s
}

type UploadInvoiceByMobileResponseBodyResult struct {
	// This parameter is required.
	Results []*UploadInvoiceByMobileResponseBodyResultResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s UploadInvoiceByMobileResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UploadInvoiceByMobileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UploadInvoiceByMobileResponseBodyResult) SetResults(v []*UploadInvoiceByMobileResponseBodyResultResults) *UploadInvoiceByMobileResponseBodyResult {
	s.Results = v
	return s
}

type UploadInvoiceByMobileResponseBodyResultResults struct {
	// This parameter is required.
	//
	// example:
	//
	// 20006
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 033002000712
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20532643
	InvoiceNo *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// duplicateInvoice
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UploadInvoiceByMobileResponseBodyResultResults) String() string {
	return tea.Prettify(s)
}

func (s UploadInvoiceByMobileResponseBodyResultResults) GoString() string {
	return s.String()
}

func (s *UploadInvoiceByMobileResponseBodyResultResults) SetErrCode(v string) *UploadInvoiceByMobileResponseBodyResultResults {
	s.ErrCode = &v
	return s
}

func (s *UploadInvoiceByMobileResponseBodyResultResults) SetInvoiceCode(v string) *UploadInvoiceByMobileResponseBodyResultResults {
	s.InvoiceCode = &v
	return s
}

func (s *UploadInvoiceByMobileResponseBodyResultResults) SetInvoiceNo(v string) *UploadInvoiceByMobileResponseBodyResultResults {
	s.InvoiceNo = &v
	return s
}

func (s *UploadInvoiceByMobileResponseBodyResultResults) SetReason(v string) *UploadInvoiceByMobileResponseBodyResultResults {
	s.Reason = &v
	return s
}

func (s *UploadInvoiceByMobileResponseBodyResultResults) SetSuccess(v bool) *UploadInvoiceByMobileResponseBodyResultResults {
	s.Success = &v
	return s
}

type UploadInvoiceByMobileResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadInvoiceByMobileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadInvoiceByMobileResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadInvoiceByMobileResponse) GoString() string {
	return s.String()
}

func (s *UploadInvoiceByMobileResponse) SetHeaders(v map[string]*string) *UploadInvoiceByMobileResponse {
	s.Headers = v
	return s
}

func (s *UploadInvoiceByMobileResponse) SetStatusCode(v int32) *UploadInvoiceByMobileResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadInvoiceByMobileResponse) SetBody(v *UploadInvoiceByMobileResponseBody) *UploadInvoiceByMobileResponse {
	s.Body = v
	return s
}

type UploadRegisterImageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UploadRegisterImageHeaders) String() string {
	return tea.Prettify(s)
}

func (s UploadRegisterImageHeaders) GoString() string {
	return s.String()
}

func (s *UploadRegisterImageHeaders) SetCommonHeaders(v map[string]*string) *UploadRegisterImageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UploadRegisterImageHeaders) SetXAcsDingtalkAccessToken(v string) *UploadRegisterImageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UploadRegisterImageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// byte[]转Base64
	ImageContent *string `json:"imageContent,omitempty" xml:"imageContent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	ImageName *string `json:"imageName,omitempty" xml:"imageName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// JPG
	ImageType *string `json:"imageType,omitempty" xml:"imageType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12020001
	InstId *string `json:"instId,omitempty" xml:"instId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALIPAY
	PayChannel *string `json:"payChannel,omitempty" xml:"payChannel,omitempty"`
}

func (s UploadRegisterImageRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadRegisterImageRequest) GoString() string {
	return s.String()
}

func (s *UploadRegisterImageRequest) SetImageContent(v string) *UploadRegisterImageRequest {
	s.ImageContent = &v
	return s
}

func (s *UploadRegisterImageRequest) SetImageName(v string) *UploadRegisterImageRequest {
	s.ImageName = &v
	return s
}

func (s *UploadRegisterImageRequest) SetImageType(v string) *UploadRegisterImageRequest {
	s.ImageType = &v
	return s
}

func (s *UploadRegisterImageRequest) SetInstId(v string) *UploadRegisterImageRequest {
	s.InstId = &v
	return s
}

func (s *UploadRegisterImageRequest) SetPayChannel(v string) *UploadRegisterImageRequest {
	s.PayChannel = &v
	return s
}

type UploadRegisterImageResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// ossUrl
	OssUrl *string `json:"ossUrl,omitempty" xml:"ossUrl,omitempty"`
}

func (s UploadRegisterImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadRegisterImageResponseBody) GoString() string {
	return s.String()
}

func (s *UploadRegisterImageResponseBody) SetOssUrl(v string) *UploadRegisterImageResponseBody {
	s.OssUrl = &v
	return s
}

type UploadRegisterImageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadRegisterImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadRegisterImageResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadRegisterImageResponse) GoString() string {
	return s.String()
}

func (s *UploadRegisterImageResponse) SetHeaders(v map[string]*string) *UploadRegisterImageResponse {
	s.Headers = v
	return s
}

func (s *UploadRegisterImageResponse) SetStatusCode(v int32) *UploadRegisterImageResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadRegisterImageResponse) SetBody(v *UploadRegisterImageResponseBody) *UploadRegisterImageResponse {
	s.Body = v
	return s
}

type UserAgreementPageSignHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UserAgreementPageSignHeaders) String() string {
	return tea.Prettify(s)
}

func (s UserAgreementPageSignHeaders) GoString() string {
	return s.String()
}

func (s *UserAgreementPageSignHeaders) SetCommonHeaders(v map[string]*string) *UserAgreementPageSignHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UserAgreementPageSignHeaders) SetXAcsDingtalkAccessToken(v string) *UserAgreementPageSignHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UserAgreementPageSignRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// TRADE
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// WITHHOLDING
	BizScene *string `json:"bizScene,omitempty" xml:"bizScene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 202111090001
	InstId *string `json:"instId,omitempty" xml:"instId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 支付宝
	PayChannel *string `json:"payChannel,omitempty" xml:"payChannel,omitempty"`
	// example:
	//
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// http://****.com
	ReturnUrl *string `json:"returnUrl,omitempty" xml:"returnUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 详见支付宝接口文档https://opendocs.alipay.com/open/20190319114403226822/signscene
	SignScene *string `json:"signScene,omitempty" xml:"signScene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1001
	SubInstId *string `json:"subInstId,omitempty" xml:"subInstId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 滴滴出行
	SubMerchantName *string `json:"subMerchantName,omitempty" xml:"subMerchantName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 免密付车费，单次最高500元
	SubMerchantServiceDesc *string `json:"subMerchantServiceDesc,omitempty" xml:"subMerchantServiceDesc,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 滴滴出行免密支付
	SubMerchantServiceName *string `json:"subMerchantServiceName,omitempty" xml:"subMerchantServiceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2120493284
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UserAgreementPageSignRequest) String() string {
	return tea.Prettify(s)
}

func (s UserAgreementPageSignRequest) GoString() string {
	return s.String()
}

func (s *UserAgreementPageSignRequest) SetBizCode(v string) *UserAgreementPageSignRequest {
	s.BizCode = &v
	return s
}

func (s *UserAgreementPageSignRequest) SetBizScene(v string) *UserAgreementPageSignRequest {
	s.BizScene = &v
	return s
}

func (s *UserAgreementPageSignRequest) SetInstId(v string) *UserAgreementPageSignRequest {
	s.InstId = &v
	return s
}

func (s *UserAgreementPageSignRequest) SetPayChannel(v string) *UserAgreementPageSignRequest {
	s.PayChannel = &v
	return s
}

func (s *UserAgreementPageSignRequest) SetRemark(v string) *UserAgreementPageSignRequest {
	s.Remark = &v
	return s
}

func (s *UserAgreementPageSignRequest) SetReturnUrl(v string) *UserAgreementPageSignRequest {
	s.ReturnUrl = &v
	return s
}

func (s *UserAgreementPageSignRequest) SetSignScene(v string) *UserAgreementPageSignRequest {
	s.SignScene = &v
	return s
}

func (s *UserAgreementPageSignRequest) SetSubInstId(v string) *UserAgreementPageSignRequest {
	s.SubInstId = &v
	return s
}

func (s *UserAgreementPageSignRequest) SetSubMerchantName(v string) *UserAgreementPageSignRequest {
	s.SubMerchantName = &v
	return s
}

func (s *UserAgreementPageSignRequest) SetSubMerchantServiceDesc(v string) *UserAgreementPageSignRequest {
	s.SubMerchantServiceDesc = &v
	return s
}

func (s *UserAgreementPageSignRequest) SetSubMerchantServiceName(v string) *UserAgreementPageSignRequest {
	s.SubMerchantServiceName = &v
	return s
}

func (s *UserAgreementPageSignRequest) SetUserId(v string) *UserAgreementPageSignRequest {
	s.UserId = &v
	return s
}

type UserAgreementPageSignResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 略
	PageData *string `json:"pageData,omitempty" xml:"pageData,omitempty"`
}

func (s UserAgreementPageSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UserAgreementPageSignResponseBody) GoString() string {
	return s.String()
}

func (s *UserAgreementPageSignResponseBody) SetPageData(v string) *UserAgreementPageSignResponseBody {
	s.PageData = &v
	return s
}

type UserAgreementPageSignResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UserAgreementPageSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UserAgreementPageSignResponse) String() string {
	return tea.Prettify(s)
}

func (s UserAgreementPageSignResponse) GoString() string {
	return s.String()
}

func (s *UserAgreementPageSignResponse) SetHeaders(v map[string]*string) *UserAgreementPageSignResponse {
	s.Headers = v
	return s
}

func (s *UserAgreementPageSignResponse) SetStatusCode(v int32) *UserAgreementPageSignResponse {
	s.StatusCode = &v
	return s
}

func (s *UserAgreementPageSignResponse) SetBody(v *UserAgreementPageSignResponseBody) *UserAgreementPageSignResponse {
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
// 批量付款
//
// @param request - ApplyBatchPayRequest
//
// @param headers - ApplyBatchPayHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyBatchPayResponse
func (client *Client) ApplyBatchPayWithOptions(request *ApplyBatchPayRequest, headers *ApplyBatchPayHeaders, runtime *util.RuntimeOptions) (_result *ApplyBatchPayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["accountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNo)) {
		body["orderNo"] = request.OrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.PassBackParams)) {
		body["passBackParams"] = request.PassBackParams
	}

	if !tea.BoolValue(util.IsUnset(request.PayTerminal)) {
		body["payTerminal"] = request.PayTerminal
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnUrl)) {
		body["returnUrl"] = request.ReturnUrl
	}

	if !tea.BoolValue(util.IsUnset(request.StaffId)) {
		body["staffId"] = request.StaffId
	}

	if !tea.BoolValue(util.IsUnset(request.TransAmount)) {
		body["transAmount"] = request.TransAmount
	}

	if !tea.BoolValue(util.IsUnset(request.TransExpireTime)) {
		body["transExpireTime"] = request.TransExpireTime
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
		Action:      tea.String("ApplyBatchPay"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/batchTrades/orders/pay"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyBatchPayResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量付款
//
// @param request - ApplyBatchPayRequest
//
// @return ApplyBatchPayResponse
func (client *Client) ApplyBatchPay(request *ApplyBatchPayRequest) (_result *ApplyBatchPayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ApplyBatchPayHeaders{}
	_result = &ApplyBatchPayResponse{}
	_body, _err := client.ApplyBatchPayWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 助贷业务关闭借款入口
//
// @param request - CloseLoanEntranceRequest
//
// @param headers - CloseLoanEntranceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseLoanEntranceResponse
func (client *Client) CloseLoanEntranceWithOptions(request *CloseLoanEntranceRequest, headers *CloseLoanEntranceHeaders, runtime *util.RuntimeOptions) (_result *CloseLoanEntranceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
		Action:      tea.String("CloseLoanEntrance"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/loans/entrances/close"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseLoanEntranceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 助贷业务关闭借款入口
//
// @param request - CloseLoanEntranceRequest
//
// @return CloseLoanEntranceResponse
func (client *Client) CloseLoanEntrance(request *CloseLoanEntranceRequest) (_result *CloseLoanEntranceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CloseLoanEntranceHeaders{}
	_result = &CloseLoanEntranceResponse{}
	_body, _err := client.CloseLoanEntranceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 子机构创建进件预校验
//
// @param request - ConsultCreateSubInstitutionRequest
//
// @param headers - ConsultCreateSubInstitutionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConsultCreateSubInstitutionResponse
func (client *Client) ConsultCreateSubInstitutionWithOptions(request *ConsultCreateSubInstitutionRequest, headers *ConsultCreateSubInstitutionHeaders, runtime *util.RuntimeOptions) (_result *ConsultCreateSubInstitutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BindingAlipayLogonId)) {
		body["bindingAlipayLogonId"] = request.BindingAlipayLogonId
	}

	if !tea.BoolValue(util.IsUnset(request.ContactInfo)) {
		body["contactInfo"] = request.ContactInfo
	}

	if !tea.BoolValue(util.IsUnset(request.InstId)) {
		body["instId"] = request.InstId
	}

	if !tea.BoolValue(util.IsUnset(request.LegalPersonCertInfo)) {
		body["legalPersonCertInfo"] = request.LegalPersonCertInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OutTradeNo)) {
		body["outTradeNo"] = request.OutTradeNo
	}

	if !tea.BoolValue(util.IsUnset(request.PayChannel)) {
		body["payChannel"] = request.PayChannel
	}

	if !tea.BoolValue(util.IsUnset(request.QualificationInfos)) {
		body["qualificationInfos"] = request.QualificationInfos
	}

	if !tea.BoolValue(util.IsUnset(request.Services)) {
		body["services"] = request.Services
	}

	if !tea.BoolValue(util.IsUnset(request.SettleInfo)) {
		body["settleInfo"] = request.SettleInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Solution)) {
		body["solution"] = request.Solution
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstAddressInfo)) {
		body["subInstAddressInfo"] = request.SubInstAddressInfo
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstAuthInfo)) {
		body["subInstAuthInfo"] = request.SubInstAuthInfo
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstBasicInfo)) {
		body["subInstBasicInfo"] = request.SubInstBasicInfo
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstCertifyInfo)) {
		body["subInstCertifyInfo"] = request.SubInstCertifyInfo
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstId)) {
		body["subInstId"] = request.SubInstId
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstInvoiceInfo)) {
		body["subInstInvoiceInfo"] = request.SubInstInvoiceInfo
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstShopInfo)) {
		body["subInstShopInfo"] = request.SubInstShopInfo
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
		Action:      tea.String("ConsultCreateSubInstitution"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/institutions/subInstitutions/consult"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ConsultCreateSubInstitutionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 子机构创建进件预校验
//
// @param request - ConsultCreateSubInstitutionRequest
//
// @return ConsultCreateSubInstitutionResponse
func (client *Client) ConsultCreateSubInstitution(request *ConsultCreateSubInstitutionRequest) (_result *ConsultCreateSubInstitutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ConsultCreateSubInstitutionHeaders{}
	_result = &ConsultCreateSubInstitutionResponse{}
	_body, _err := client.ConsultCreateSubInstitutionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发起代扣交易
//
// @param request - CreatWithholdingOrderAndPayRequest
//
// @param headers - CreatWithholdingOrderAndPayHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatWithholdingOrderAndPayResponse
func (client *Client) CreatWithholdingOrderAndPayWithOptions(request *CreatWithholdingOrderAndPayRequest, headers *CreatWithholdingOrderAndPayHeaders, runtime *util.RuntimeOptions) (_result *CreatWithholdingOrderAndPayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		body["amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.InstId)) {
		body["instId"] = request.InstId
	}

	if !tea.BoolValue(util.IsUnset(request.OtherPayChannelDetailInfoList)) {
		body["otherPayChannelDetailInfoList"] = request.OtherPayChannelDetailInfoList
	}

	if !tea.BoolValue(util.IsUnset(request.OutTradeNo)) {
		body["outTradeNo"] = request.OutTradeNo
	}

	if !tea.BoolValue(util.IsUnset(request.PayChannel)) {
		body["payChannel"] = request.PayChannel
	}

	if !tea.BoolValue(util.IsUnset(request.PayerUserId)) {
		body["payerUserId"] = request.PayerUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstId)) {
		body["subInstId"] = request.SubInstId
	}

	if !tea.BoolValue(util.IsUnset(request.TimeOutExpress)) {
		body["timeOutExpress"] = request.TimeOutExpress
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
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
		Action:      tea.String("CreatWithholdingOrderAndPay"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/withholdingOrders"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatWithholdingOrderAndPayResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发起代扣交易
//
// @param request - CreatWithholdingOrderAndPayRequest
//
// @return CreatWithholdingOrderAndPayResponse
func (client *Client) CreatWithholdingOrderAndPay(request *CreatWithholdingOrderAndPayRequest) (_result *CreatWithholdingOrderAndPayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreatWithholdingOrderAndPayHeaders{}
	_result = &CreatWithholdingOrderAndPayResponse{}
	_body, _err := client.CreatWithholdingOrderAndPayWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 收单退款交易
//
// @param request - CreateAcquireRefundOrderRequest
//
// @param headers - CreateAcquireRefundOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAcquireRefundOrderResponse
func (client *Client) CreateAcquireRefundOrderWithOptions(request *CreateAcquireRefundOrderRequest, headers *CreateAcquireRefundOrderHeaders, runtime *util.RuntimeOptions) (_result *CreateAcquireRefundOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstId)) {
		body["instId"] = request.InstId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.OriginOutTradeNo)) {
		body["originOutTradeNo"] = request.OriginOutTradeNo
	}

	if !tea.BoolValue(util.IsUnset(request.OtherPayChannelDetailInfoList)) {
		body["otherPayChannelDetailInfoList"] = request.OtherPayChannelDetailInfoList
	}

	if !tea.BoolValue(util.IsUnset(request.OutRefundNo)) {
		body["outRefundNo"] = request.OutRefundNo
	}

	if !tea.BoolValue(util.IsUnset(request.RefundAmount)) {
		body["refundAmount"] = request.RefundAmount
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstId)) {
		body["subInstId"] = request.SubInstId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
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
		Action:      tea.String("CreateAcquireRefundOrder"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/acquireOrders/refund"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAcquireRefundOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 收单退款交易
//
// @param request - CreateAcquireRefundOrderRequest
//
// @return CreateAcquireRefundOrderResponse
func (client *Client) CreateAcquireRefundOrder(request *CreateAcquireRefundOrderRequest) (_result *CreateAcquireRefundOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateAcquireRefundOrderHeaders{}
	_result = &CreateAcquireRefundOrderResponse{}
	_body, _err := client.CreateAcquireRefundOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建批量支付单
//
// @param request - CreateBatchTradeOrderRequest
//
// @param headers - CreateBatchTradeOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBatchTradeOrderResponse
func (client *Client) CreateBatchTradeOrderWithOptions(request *CreateBatchTradeOrderRequest, headers *CreateBatchTradeOrderHeaders, runtime *util.RuntimeOptions) (_result *CreateBatchTradeOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["accountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.AccountNo)) {
		body["accountNo"] = request.AccountNo
	}

	if !tea.BoolValue(util.IsUnset(request.BatchRemark)) {
		body["batchRemark"] = request.BatchRemark
	}

	if !tea.BoolValue(util.IsUnset(request.BatchTradeDetails)) {
		body["batchTradeDetails"] = request.BatchTradeDetails
	}

	if !tea.BoolValue(util.IsUnset(request.OutBatchNo)) {
		body["outBatchNo"] = request.OutBatchNo
	}

	if !tea.BoolValue(util.IsUnset(request.StaffId)) {
		body["staffId"] = request.StaffId
	}

	if !tea.BoolValue(util.IsUnset(request.TotalAmount)) {
		body["totalAmount"] = request.TotalAmount
	}

	if !tea.BoolValue(util.IsUnset(request.TotalCount)) {
		body["totalCount"] = request.TotalCount
	}

	if !tea.BoolValue(util.IsUnset(request.TradeTitle)) {
		body["tradeTitle"] = request.TradeTitle
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
		Action:      tea.String("CreateBatchTradeOrder"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/batchTrades/orders"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBatchTradeOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建批量支付单
//
// @param request - CreateBatchTradeOrderRequest
//
// @return CreateBatchTradeOrderResponse
func (client *Client) CreateBatchTradeOrder(request *CreateBatchTradeOrderRequest) (_result *CreateBatchTradeOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateBatchTradeOrderHeaders{}
	_result = &CreateBatchTradeOrderResponse{}
	_body, _err := client.CreateBatchTradeOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建子机构
//
// @param request - CreateSubInstitutionRequest
//
// @param headers - CreateSubInstitutionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSubInstitutionResponse
func (client *Client) CreateSubInstitutionWithOptions(request *CreateSubInstitutionRequest, headers *CreateSubInstitutionHeaders, runtime *util.RuntimeOptions) (_result *CreateSubInstitutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BindingAlipayLogonId)) {
		body["bindingAlipayLogonId"] = request.BindingAlipayLogonId
	}

	if !tea.BoolValue(util.IsUnset(request.ContactInfo)) {
		body["contactInfo"] = request.ContactInfo
	}

	if !tea.BoolValue(util.IsUnset(request.InstId)) {
		body["instId"] = request.InstId
	}

	if !tea.BoolValue(util.IsUnset(request.LegalPersonCertInfo)) {
		body["legalPersonCertInfo"] = request.LegalPersonCertInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OutTradeNo)) {
		body["outTradeNo"] = request.OutTradeNo
	}

	if !tea.BoolValue(util.IsUnset(request.PayChannel)) {
		body["payChannel"] = request.PayChannel
	}

	if !tea.BoolValue(util.IsUnset(request.QualificationInfos)) {
		body["qualificationInfos"] = request.QualificationInfos
	}

	if !tea.BoolValue(util.IsUnset(request.Services)) {
		body["services"] = request.Services
	}

	if !tea.BoolValue(util.IsUnset(request.SettleInfo)) {
		body["settleInfo"] = request.SettleInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Solution)) {
		body["solution"] = request.Solution
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstAddressInfo)) {
		body["subInstAddressInfo"] = request.SubInstAddressInfo
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstAuthInfo)) {
		body["subInstAuthInfo"] = request.SubInstAuthInfo
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstBasicInfo)) {
		body["subInstBasicInfo"] = request.SubInstBasicInfo
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstCertifyInfo)) {
		body["subInstCertifyInfo"] = request.SubInstCertifyInfo
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstId)) {
		body["subInstId"] = request.SubInstId
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstInvoiceInfo)) {
		body["subInstInvoiceInfo"] = request.SubInstInvoiceInfo
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstShopInfo)) {
		body["subInstShopInfo"] = request.SubInstShopInfo
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
		Action:      tea.String("CreateSubInstitution"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/institutions/subInstitutions"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSubInstitutionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建子机构
//
// @param request - CreateSubInstitutionRequest
//
// @return CreateSubInstitutionResponse
func (client *Client) CreateSubInstitution(request *CreateSubInstitutionRequest) (_result *CreateSubInstitutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateSubInstitutionHeaders{}
	_result = &CreateSubInstitutionResponse{}
	_body, _err := client.CreateSubInstitutionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建用户码实例
//
// @param request - CreateUserCodeInstanceRequest
//
// @param headers - CreateUserCodeInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserCodeInstanceResponse
func (client *Client) CreateUserCodeInstanceWithOptions(request *CreateUserCodeInstanceRequest, headers *CreateUserCodeInstanceHeaders, runtime *util.RuntimeOptions) (_result *CreateUserCodeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AvailableTimes)) {
		body["availableTimes"] = request.AvailableTimes
	}

	if !tea.BoolValue(util.IsUnset(request.CodeIdentity)) {
		body["codeIdentity"] = request.CodeIdentity
	}

	if !tea.BoolValue(util.IsUnset(request.CodeValue)) {
		body["codeValue"] = request.CodeValue
	}

	if !tea.BoolValue(util.IsUnset(request.CodeValueType)) {
		body["codeValueType"] = request.CodeValueType
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["extInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.GmtExpired)) {
		body["gmtExpired"] = request.GmtExpired
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserCorpRelationType)) {
		body["userCorpRelationType"] = request.UserCorpRelationType
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdentity)) {
		body["userIdentity"] = request.UserIdentity
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
		Action:      tea.String("CreateUserCodeInstance"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/payCodes/userInstances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserCodeInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建用户码实例
//
// @param request - CreateUserCodeInstanceRequest
//
// @return CreateUserCodeInstanceResponse
func (client *Client) CreateUserCodeInstance(request *CreateUserCodeInstanceRequest) (_result *CreateUserCodeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateUserCodeInstanceHeaders{}
	_result = &CreateUserCodeInstanceResponse{}
	_body, _err := client.CreateUserCodeInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解码付款码
//
// @param request - DecodePayCodeRequest
//
// @param headers - DecodePayCodeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DecodePayCodeResponse
func (client *Client) DecodePayCodeWithOptions(request *DecodePayCodeRequest, headers *DecodePayCodeHeaders, runtime *util.RuntimeOptions) (_result *DecodePayCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PayCode)) {
		body["payCode"] = request.PayCode
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
		Action:      tea.String("DecodePayCode"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/payCodes/decode"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DecodePayCodeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解码付款码
//
// @param request - DecodePayCodeRequest
//
// @return DecodePayCodeResponse
func (client *Client) DecodePayCode(request *DecodePayCodeRequest) (_result *DecodePayCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DecodePayCodeHeaders{}
	_result = &DecodePayCodeResponse{}
	_body, _err := client.DecodePayCodeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 企业金融助贷业务进件通知接口
//
// @param request - FinanceLoanNotifyRegisterRequest
//
// @param headers - FinanceLoanNotifyRegisterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FinanceLoanNotifyRegisterResponse
func (client *Client) FinanceLoanNotifyRegisterWithOptions(request *FinanceLoanNotifyRegisterRequest, headers *FinanceLoanNotifyRegisterHeaders, runtime *util.RuntimeOptions) (_result *FinanceLoanNotifyRegisterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompleteTime)) {
		body["completeTime"] = request.CompleteTime
	}

	if !tea.BoolValue(util.IsUnset(request.Extension)) {
		body["extension"] = request.Extension
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardNo)) {
		body["idCardNo"] = request.IdCardNo
	}

	if !tea.BoolValue(util.IsUnset(request.OpenChannelName)) {
		body["openChannelName"] = request.OpenChannelName
	}

	if !tea.BoolValue(util.IsUnset(request.OpenProductCode)) {
		body["openProductCode"] = request.OpenProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.OpenProductName)) {
		body["openProductName"] = request.OpenProductName
	}

	if !tea.BoolValue(util.IsUnset(request.OpenProductType)) {
		body["openProductType"] = request.OpenProductType
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessingStatus)) {
		body["processingStatus"] = request.ProcessingStatus
	}

	if !tea.BoolValue(util.IsUnset(request.RefuseCode)) {
		body["refuseCode"] = request.RefuseCode
	}

	if !tea.BoolValue(util.IsUnset(request.RefuseReason)) {
		body["refuseReason"] = request.RefuseReason
	}

	if !tea.BoolValue(util.IsUnset(request.RegisterNo)) {
		body["registerNo"] = request.RegisterNo
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.SubmitTime)) {
		body["submitTime"] = request.SubmitTime
	}

	if !tea.BoolValue(util.IsUnset(request.UserMobile)) {
		body["userMobile"] = request.UserMobile
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
		Action:      tea.String("FinanceLoanNotifyRegister"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/loans/notifications/register"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &FinanceLoanNotifyRegisterResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 企业金融助贷业务进件通知接口
//
// @param request - FinanceLoanNotifyRegisterRequest
//
// @return FinanceLoanNotifyRegisterResponse
func (client *Client) FinanceLoanNotifyRegister(request *FinanceLoanNotifyRegisterRequest) (_result *FinanceLoanNotifyRegisterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &FinanceLoanNotifyRegisterHeaders{}
	_result = &FinanceLoanNotifyRegisterResponse{}
	_body, _err := client.FinanceLoanNotifyRegisterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改子机构信息
//
// @param request - ModifySubInstitutionRequest
//
// @param headers - ModifySubInstitutionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySubInstitutionResponse
func (client *Client) ModifySubInstitutionWithOptions(request *ModifySubInstitutionRequest, headers *ModifySubInstitutionHeaders, runtime *util.RuntimeOptions) (_result *ModifySubInstitutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BindingAlipayLogonId)) {
		body["bindingAlipayLogonId"] = request.BindingAlipayLogonId
	}

	if !tea.BoolValue(util.IsUnset(request.ContactInfo)) {
		body["contactInfo"] = request.ContactInfo
	}

	if !tea.BoolValue(util.IsUnset(request.InstId)) {
		body["instId"] = request.InstId
	}

	if !tea.BoolValue(util.IsUnset(request.LegalPersonCertInfo)) {
		body["legalPersonCertInfo"] = request.LegalPersonCertInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OutTradeNo)) {
		body["outTradeNo"] = request.OutTradeNo
	}

	if !tea.BoolValue(util.IsUnset(request.PayChannel)) {
		body["payChannel"] = request.PayChannel
	}

	if !tea.BoolValue(util.IsUnset(request.QualificationInfos)) {
		body["qualificationInfos"] = request.QualificationInfos
	}

	if !tea.BoolValue(util.IsUnset(request.Services)) {
		body["services"] = request.Services
	}

	if !tea.BoolValue(util.IsUnset(request.SettleInfo)) {
		body["settleInfo"] = request.SettleInfo
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstAddressInfo)) {
		body["subInstAddressInfo"] = request.SubInstAddressInfo
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstAuthInfo)) {
		body["subInstAuthInfo"] = request.SubInstAuthInfo
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstBasicInfo)) {
		body["subInstBasicInfo"] = request.SubInstBasicInfo
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstCertifyInfo)) {
		body["subInstCertifyInfo"] = request.SubInstCertifyInfo
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstId)) {
		body["subInstId"] = request.SubInstId
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstInvoiceInfo)) {
		body["subInstInvoiceInfo"] = request.SubInstInvoiceInfo
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstShopInfo)) {
		body["subInstShopInfo"] = request.SubInstShopInfo
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
		Action:      tea.String("ModifySubInstitution"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/institutions/subInstitutions/modify"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySubInstitutionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改子机构信息
//
// @param request - ModifySubInstitutionRequest
//
// @return ModifySubInstitutionResponse
func (client *Client) ModifySubInstitution(request *ModifySubInstitutionRequest) (_result *ModifySubInstitutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ModifySubInstitutionHeaders{}
	_result = &ModifySubInstitutionResponse{}
	_body, _err := client.ModifySubInstitutionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通知付款码支付结果
//
// @param request - NotifyPayCodePayResultRequest
//
// @param headers - NotifyPayCodePayResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return NotifyPayCodePayResultResponse
func (client *Client) NotifyPayCodePayResultWithOptions(request *NotifyPayCodePayResultRequest, headers *NotifyPayCodePayResultHeaders, runtime *util.RuntimeOptions) (_result *NotifyPayCodePayResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		body["amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeAmount)) {
		body["chargeAmount"] = request.ChargeAmount
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["extInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.GmtTradeCreate)) {
		body["gmtTradeCreate"] = request.GmtTradeCreate
	}

	if !tea.BoolValue(util.IsUnset(request.GmtTradeFinish)) {
		body["gmtTradeFinish"] = request.GmtTradeFinish
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantName)) {
		body["merchantName"] = request.MerchantName
	}

	if !tea.BoolValue(util.IsUnset(request.PayChannelDetailList)) {
		body["payChannelDetailList"] = request.PayChannelDetailList
	}

	if !tea.BoolValue(util.IsUnset(request.PayCode)) {
		body["payCode"] = request.PayCode
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionAmount)) {
		body["promotionAmount"] = request.PromotionAmount
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.TradeErrorCode)) {
		body["tradeErrorCode"] = request.TradeErrorCode
	}

	if !tea.BoolValue(util.IsUnset(request.TradeErrorMsg)) {
		body["tradeErrorMsg"] = request.TradeErrorMsg
	}

	if !tea.BoolValue(util.IsUnset(request.TradeNo)) {
		body["tradeNo"] = request.TradeNo
	}

	if !tea.BoolValue(util.IsUnset(request.TradeStatus)) {
		body["tradeStatus"] = request.TradeStatus
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
		Action:      tea.String("NotifyPayCodePayResult"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/payCodes/payResults/notify"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &NotifyPayCodePayResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通知付款码支付结果
//
// @param request - NotifyPayCodePayResultRequest
//
// @return NotifyPayCodePayResultResponse
func (client *Client) NotifyPayCodePayResult(request *NotifyPayCodePayResultRequest) (_result *NotifyPayCodePayResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &NotifyPayCodePayResultHeaders{}
	_result = &NotifyPayCodePayResultResponse{}
	_body, _err := client.NotifyPayCodePayResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通知付款码退款结果
//
// @param request - NotifyPayCodeRefundResultRequest
//
// @param headers - NotifyPayCodeRefundResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return NotifyPayCodeRefundResultResponse
func (client *Client) NotifyPayCodeRefundResultWithOptions(request *NotifyPayCodeRefundResultRequest, headers *NotifyPayCodeRefundResultHeaders, runtime *util.RuntimeOptions) (_result *NotifyPayCodeRefundResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.GmtRefund)) {
		body["gmtRefund"] = request.GmtRefund
	}

	if !tea.BoolValue(util.IsUnset(request.PayChannelDetailList)) {
		body["payChannelDetailList"] = request.PayChannelDetailList
	}

	if !tea.BoolValue(util.IsUnset(request.PayCode)) {
		body["payCode"] = request.PayCode
	}

	if !tea.BoolValue(util.IsUnset(request.RefundAmount)) {
		body["refundAmount"] = request.RefundAmount
	}

	if !tea.BoolValue(util.IsUnset(request.RefundOrderNo)) {
		body["refundOrderNo"] = request.RefundOrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.RefundPromotionAmount)) {
		body["refundPromotionAmount"] = request.RefundPromotionAmount
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.TradeNo)) {
		body["tradeNo"] = request.TradeNo
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
		Action:      tea.String("NotifyPayCodeRefundResult"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/payCodes/refundResults/notify"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &NotifyPayCodeRefundResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通知付款码退款结果
//
// @param request - NotifyPayCodeRefundResultRequest
//
// @return NotifyPayCodeRefundResultResponse
func (client *Client) NotifyPayCodeRefundResult(request *NotifyPayCodeRefundResultRequest) (_result *NotifyPayCodeRefundResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &NotifyPayCodeRefundResultHeaders{}
	_result = &NotifyPayCodeRefundResultResponse{}
	_body, _err := client.NotifyPayCodeRefundResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上报码验证结果
//
// @param request - NotifyVerifyResultRequest
//
// @param headers - NotifyVerifyResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return NotifyVerifyResultResponse
func (client *Client) NotifyVerifyResultWithOptions(request *NotifyVerifyResultRequest, headers *NotifyVerifyResultHeaders, runtime *util.RuntimeOptions) (_result *NotifyVerifyResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.PayCode)) {
		body["payCode"] = request.PayCode
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.UserCorpRelationType)) {
		body["userCorpRelationType"] = request.UserCorpRelationType
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdentity)) {
		body["userIdentity"] = request.UserIdentity
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyEvent)) {
		body["verifyEvent"] = request.VerifyEvent
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyLocation)) {
		body["verifyLocation"] = request.VerifyLocation
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyNo)) {
		body["verifyNo"] = request.VerifyNo
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyResult)) {
		body["verifyResult"] = request.VerifyResult
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyTime)) {
		body["verifyTime"] = request.VerifyTime
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
		Action:      tea.String("NotifyVerifyResult"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/payCodes/verifyResults/notify"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &NotifyVerifyResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上报码验证结果
//
// @param request - NotifyVerifyResultRequest
//
// @return NotifyVerifyResultResponse
func (client *Client) NotifyVerifyResult(request *NotifyVerifyResultRequest) (_result *NotifyVerifyResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &NotifyVerifyResultHeaders{}
	_result = &NotifyVerifyResultResponse{}
	_body, _err := client.NotifyVerifyResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 预创建群收款订单返回订单号
//
// @param request - PreCreateGroupBillOrderRequest
//
// @param headers - PreCreateGroupBillOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreCreateGroupBillOrderResponse
func (client *Client) PreCreateGroupBillOrderWithOptions(request *PreCreateGroupBillOrderRequest, headers *PreCreateGroupBillOrderHeaders, runtime *util.RuntimeOptions) (_result *PreCreateGroupBillOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillItemList)) {
		body["billItemList"] = request.BillItemList
	}

	if !tea.BoolValue(util.IsUnset(request.ExtParams)) {
		body["extParams"] = request.ExtParams
	}

	if !tea.BoolValue(util.IsUnset(request.HeadCount)) {
		body["headCount"] = request.HeadCount
	}

	if !tea.BoolValue(util.IsUnset(request.IsAverageAmount)) {
		body["isAverageAmount"] = request.IsAverageAmount
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantId)) {
		body["merchantId"] = request.MerchantId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenCid)) {
		body["openCid"] = request.OpenCid
	}

	if !tea.BoolValue(util.IsUnset(request.OutBizNo)) {
		body["outBizNo"] = request.OutBizNo
	}

	if !tea.BoolValue(util.IsUnset(request.PayeeCorpId)) {
		body["payeeCorpId"] = request.PayeeCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.PayeeUnionId)) {
		body["payeeUnionId"] = request.PayeeUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.TotalAmount)) {
		body["totalAmount"] = request.TotalAmount
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
		Action:      tea.String("PreCreateGroupBillOrder"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/groupbills/preCreate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PreCreateGroupBillOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预创建群收款订单返回订单号
//
// @param request - PreCreateGroupBillOrderRequest
//
// @return PreCreateGroupBillOrderResponse
func (client *Client) PreCreateGroupBillOrder(request *PreCreateGroupBillOrderRequest) (_result *PreCreateGroupBillOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PreCreateGroupBillOrderHeaders{}
	_result = &PreCreateGroupBillOrderResponse{}
	_body, _err := client.PreCreateGroupBillOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询收单退款交易
//
// @param request - QueryAcquireRefundOrderRequest
//
// @param headers - QueryAcquireRefundOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAcquireRefundOrderResponse
func (client *Client) QueryAcquireRefundOrderWithOptions(request *QueryAcquireRefundOrderRequest, headers *QueryAcquireRefundOrderHeaders, runtime *util.RuntimeOptions) (_result *QueryAcquireRefundOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutRefundNo)) {
		query["outRefundNo"] = request.OutRefundNo
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
		Action:      tea.String("QueryAcquireRefundOrder"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/acquireOrders/refunds"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAcquireRefundOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询收单退款交易
//
// @param request - QueryAcquireRefundOrderRequest
//
// @return QueryAcquireRefundOrderResponse
func (client *Client) QueryAcquireRefundOrder(request *QueryAcquireRefundOrderRequest) (_result *QueryAcquireRefundOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAcquireRefundOrderHeaders{}
	_result = &QueryAcquireRefundOrderResponse{}
	_body, _err := client.QueryAcquireRefundOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询批量付款明细列表
//
// @param request - QueryBatchTradeDetailListRequest
//
// @param headers - QueryBatchTradeDetailListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBatchTradeDetailListResponse
func (client *Client) QueryBatchTradeDetailListWithOptions(request *QueryBatchTradeDetailListRequest, headers *QueryBatchTradeDetailListHeaders, runtime *util.RuntimeOptions) (_result *QueryBatchTradeDetailListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutBatchNo)) {
		query["outBatchNo"] = request.OutBatchNo
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
		Action:      tea.String("QueryBatchTradeDetailList"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/batchTrades/details"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBatchTradeDetailListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询批量付款明细列表
//
// @param request - QueryBatchTradeDetailListRequest
//
// @return QueryBatchTradeDetailListResponse
func (client *Client) QueryBatchTradeDetailList(request *QueryBatchTradeDetailListRequest) (_result *QueryBatchTradeDetailListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryBatchTradeDetailListHeaders{}
	_result = &QueryBatchTradeDetailListResponse{}
	_body, _err := client.QueryBatchTradeDetailListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据请求对象查询批量交易订单信息
//
// @param request - QueryBatchTradeOrderRequest
//
// @param headers - QueryBatchTradeOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBatchTradeOrderResponse
func (client *Client) QueryBatchTradeOrderWithOptions(request *QueryBatchTradeOrderRequest, headers *QueryBatchTradeOrderHeaders, runtime *util.RuntimeOptions) (_result *QueryBatchTradeOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutBatchNos)) {
		body["outBatchNos"] = request.OutBatchNos
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
		Action:      tea.String("QueryBatchTradeOrder"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/batchTrades/orders/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBatchTradeOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据请求对象查询批量交易订单信息
//
// @param request - QueryBatchTradeOrderRequest
//
// @return QueryBatchTradeOrderResponse
func (client *Client) QueryBatchTradeOrder(request *QueryBatchTradeOrderRequest) (_result *QueryBatchTradeOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryBatchTradeOrderHeaders{}
	_result = &QueryBatchTradeOrderResponse{}
	_body, _err := client.QueryBatchTradeOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询付款账号列表
//
// @param headers - QueryPayAccountListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPayAccountListResponse
func (client *Client) QueryPayAccountListWithOptions(headers *QueryPayAccountListHeaders, runtime *util.RuntimeOptions) (_result *QueryPayAccountListResponse, _err error) {
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
		Action:      tea.String("QueryPayAccountList"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/payAccounts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPayAccountListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询付款账号列表
//
// @return QueryPayAccountListResponse
func (client *Client) QueryPayAccountList() (_result *QueryPayAccountListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryPayAccountListHeaders{}
	_result = &QueryPayAccountListResponse{}
	_body, _err := client.QueryPayAccountListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询子机构申请单状态
//
// @param request - QueryRegisterOrderRequest
//
// @param headers - QueryRegisterOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRegisterOrderResponse
func (client *Client) QueryRegisterOrderWithOptions(request *QueryRegisterOrderRequest, headers *QueryRegisterOrderHeaders, runtime *util.RuntimeOptions) (_result *QueryRegisterOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstId)) {
		query["instId"] = request.InstId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["orderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OutTradeNo)) {
		query["outTradeNo"] = request.OutTradeNo
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstId)) {
		query["subInstId"] = request.SubInstId
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
		Action:      tea.String("QueryRegisterOrder"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/institutions/subInstitutions/orders"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRegisterOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询子机构申请单状态
//
// @param request - QueryRegisterOrderRequest
//
// @return QueryRegisterOrderResponse
func (client *Client) QueryRegisterOrder(request *QueryRegisterOrderRequest) (_result *QueryRegisterOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryRegisterOrderHeaders{}
	_result = &QueryRegisterOrderResponse{}
	_body, _err := client.QueryRegisterOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户协议
//
// @param request - QueryUserAgreementRequest
//
// @param headers - QueryUserAgreementHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserAgreementResponse
func (client *Client) QueryUserAgreementWithOptions(request *QueryUserAgreementRequest, headers *QueryUserAgreementHeaders, runtime *util.RuntimeOptions) (_result *QueryUserAgreementResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["bizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.BizScene)) {
		query["bizScene"] = request.BizScene
	}

	if !tea.BoolValue(util.IsUnset(request.InstId)) {
		query["instId"] = request.InstId
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstId)) {
		query["subInstId"] = request.SubInstId
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
		Action:      tea.String("QueryUserAgreement"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/userAgreements"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserAgreementResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户协议
//
// @param request - QueryUserAgreementRequest
//
// @return QueryUserAgreementResponse
func (client *Client) QueryUserAgreement(request *QueryUserAgreementRequest) (_result *QueryUserAgreementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserAgreementHeaders{}
	_result = &QueryUserAgreementResponse{}
	_body, _err := client.QueryUserAgreementWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户绑定支付宝信息
//
// @param headers - QueryUserAlipayAccountHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserAlipayAccountResponse
func (client *Client) QueryUserAlipayAccountWithOptions(headers *QueryUserAlipayAccountHeaders, runtime *util.RuntimeOptions) (_result *QueryUserAlipayAccountResponse, _err error) {
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
		Action:      tea.String("QueryUserAlipayAccount"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/userAlipayAccounts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserAlipayAccountResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户绑定支付宝信息
//
// @return QueryUserAlipayAccountResponse
func (client *Client) QueryUserAlipayAccount() (_result *QueryUserAlipayAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserAlipayAccountHeaders{}
	_result = &QueryUserAlipayAccountResponse{}
	_body, _err := client.QueryUserAlipayAccountWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询代扣交易订单信息
//
// @param request - QueryWithholdingOrderRequest
//
// @param headers - QueryWithholdingOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryWithholdingOrderResponse
func (client *Client) QueryWithholdingOrderWithOptions(request *QueryWithholdingOrderRequest, headers *QueryWithholdingOrderHeaders, runtime *util.RuntimeOptions) (_result *QueryWithholdingOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutTradeNo)) {
		query["outTradeNo"] = request.OutTradeNo
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
		Action:      tea.String("QueryWithholdingOrder"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/withholdingOrders"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryWithholdingOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询代扣交易订单信息
//
// @param request - QueryWithholdingOrderRequest
//
// @return QueryWithholdingOrderResponse
func (client *Client) QueryWithholdingOrder(request *QueryWithholdingOrderRequest) (_result *QueryWithholdingOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryWithholdingOrderHeaders{}
	_result = &QueryWithholdingOrderResponse{}
	_body, _err := client.QueryWithholdingOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 保存付款码企业配置信息
//
// @param request - SaveCorpPayCodeRequest
//
// @param headers - SaveCorpPayCodeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveCorpPayCodeResponse
func (client *Client) SaveCorpPayCodeWithOptions(request *SaveCorpPayCodeRequest, headers *SaveCorpPayCodeHeaders, runtime *util.RuntimeOptions) (_result *SaveCorpPayCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CodeIdentity)) {
		body["codeIdentity"] = request.CodeIdentity
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["extInfo"] = request.ExtInfo
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
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveCorpPayCode"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/payCodes/corpSettings"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveCorpPayCodeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存付款码企业配置信息
//
// @param request - SaveCorpPayCodeRequest
//
// @return SaveCorpPayCodeResponse
func (client *Client) SaveCorpPayCode(request *SaveCorpPayCodeRequest) (_result *SaveCorpPayCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveCorpPayCodeHeaders{}
	_result = &SaveCorpPayCodeResponse{}
	_body, _err := client.SaveCorpPayCodeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解约用户协议
//
// @param request - UnsignUserAgreementRequest
//
// @param headers - UnsignUserAgreementHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnsignUserAgreementResponse
func (client *Client) UnsignUserAgreementWithOptions(request *UnsignUserAgreementRequest, headers *UnsignUserAgreementHeaders, runtime *util.RuntimeOptions) (_result *UnsignUserAgreementResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgreementNo)) {
		body["agreementNo"] = request.AgreementNo
	}

	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		body["bizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.BizScene)) {
		body["bizScene"] = request.BizScene
	}

	if !tea.BoolValue(util.IsUnset(request.InstId)) {
		body["instId"] = request.InstId
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstId)) {
		body["subInstId"] = request.SubInstId
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
		Action:      tea.String("UnsignUserAgreement"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/userAgreements/unsign"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("none"),
	}
	_result = &UnsignUserAgreementResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解约用户协议
//
// @param request - UnsignUserAgreementRequest
//
// @return UnsignUserAgreementResponse
func (client *Client) UnsignUserAgreement(request *UnsignUserAgreementRequest) (_result *UnsignUserAgreementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UnsignUserAgreementHeaders{}
	_result = &UnsignUserAgreementResponse{}
	_body, _err := client.UnsignUserAgreementWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新用户码实例
//
// @param request - UpateUserCodeInstanceRequest
//
// @param headers - UpateUserCodeInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpateUserCodeInstanceResponse
func (client *Client) UpateUserCodeInstanceWithOptions(request *UpateUserCodeInstanceRequest, headers *UpateUserCodeInstanceHeaders, runtime *util.RuntimeOptions) (_result *UpateUserCodeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AvailableTimes)) {
		body["availableTimes"] = request.AvailableTimes
	}

	if !tea.BoolValue(util.IsUnset(request.CodeId)) {
		body["codeId"] = request.CodeId
	}

	if !tea.BoolValue(util.IsUnset(request.CodeIdentity)) {
		body["codeIdentity"] = request.CodeIdentity
	}

	if !tea.BoolValue(util.IsUnset(request.CodeValue)) {
		body["codeValue"] = request.CodeValue
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["extInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.GmtExpired)) {
		body["gmtExpired"] = request.GmtExpired
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserCorpRelationType)) {
		body["userCorpRelationType"] = request.UserCorpRelationType
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdentity)) {
		body["userIdentity"] = request.UserIdentity
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
		Action:      tea.String("UpateUserCodeInstance"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/payCodes/userInstances"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpateUserCodeInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新用户码实例
//
// @param request - UpateUserCodeInstanceRequest
//
// @return UpateUserCodeInstanceResponse
func (client *Client) UpateUserCodeInstance(request *UpateUserCodeInstanceRequest) (_result *UpateUserCodeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpateUserCodeInstanceHeaders{}
	_result = &UpateUserCodeInstanceResponse{}
	_body, _err := client.UpateUserCodeInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用来给第三方企业提供发票验真结果更新的oapi
//
// @param request - UpdateInvoiceVerifyStatusRequest
//
// @param headers - UpdateInvoiceVerifyStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInvoiceVerifyStatusResponse
func (client *Client) UpdateInvoiceVerifyStatusWithOptions(request *UpdateInvoiceVerifyStatusRequest, headers *UpdateInvoiceVerifyStatusHeaders, runtime *util.RuntimeOptions) (_result *UpdateInvoiceVerifyStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.CheckingResult)) {
		body["checkingResult"] = request.CheckingResult
	}

	if !tea.BoolValue(util.IsUnset(request.CheckingStatus)) {
		body["checkingStatus"] = request.CheckingStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Extension)) {
		body["extension"] = request.Extension
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceCode)) {
		body["invoiceCode"] = request.InvoiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceNo)) {
		body["invoiceNo"] = request.InvoiceNo
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceStatus)) {
		body["invoiceStatus"] = request.InvoiceStatus
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceVerifyId)) {
		body["invoiceVerifyId"] = request.InvoiceVerifyId
	}

	if !tea.BoolValue(util.IsUnset(request.Msg)) {
		body["msg"] = request.Msg
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
		Action:      tea.String("UpdateInvoiceVerifyStatus"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/invoices/verifyStatus"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInvoiceVerifyStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用来给第三方企业提供发票验真结果更新的oapi
//
// @param request - UpdateInvoiceVerifyStatusRequest
//
// @return UpdateInvoiceVerifyStatusResponse
func (client *Client) UpdateInvoiceVerifyStatus(request *UpdateInvoiceVerifyStatusRequest) (_result *UpdateInvoiceVerifyStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateInvoiceVerifyStatusHeaders{}
	_result = &UpdateInvoiceVerifyStatusResponse{}
	_body, _err := client.UpdateInvoiceVerifyStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上传发票
//
// @param request - UploadInvoiceRequest
//
// @param headers - UploadInvoiceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadInvoiceResponse
func (client *Client) UploadInvoiceWithOptions(request *UploadInvoiceRequest, headers *UploadInvoiceHeaders, runtime *util.RuntimeOptions) (_result *UploadInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Extension)) {
		body["extension"] = request.Extension
	}

	if !tea.BoolValue(util.IsUnset(request.Invoices)) {
		body["invoices"] = request.Invoices
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdentity)) {
		body["userIdentity"] = request.UserIdentity
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
		Action:      tea.String("UploadInvoice"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/invoices/upload"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadInvoiceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传发票
//
// @param request - UploadInvoiceRequest
//
// @return UploadInvoiceResponse
func (client *Client) UploadInvoice(request *UploadInvoiceRequest) (_result *UploadInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UploadInvoiceHeaders{}
	_result = &UploadInvoiceResponse{}
	_body, _err := client.UploadInvoiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用户授权上传发票oapi
//
// @param request - UploadInvoiceByAuthRequest
//
// @param headers - UploadInvoiceByAuthHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadInvoiceByAuthResponse
func (client *Client) UploadInvoiceByAuthWithOptions(request *UploadInvoiceByAuthRequest, headers *UploadInvoiceByAuthHeaders, runtime *util.RuntimeOptions) (_result *UploadInvoiceByAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Extension)) {
		body["extension"] = request.Extension
	}

	if !tea.BoolValue(util.IsUnset(request.Invoices)) {
		body["invoices"] = request.Invoices
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
		Action:      tea.String("UploadInvoiceByAuth"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/invoices/authorizations/upload"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadInvoiceByAuthResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用户授权上传发票oapi
//
// @param request - UploadInvoiceByAuthRequest
//
// @return UploadInvoiceByAuthResponse
func (client *Client) UploadInvoiceByAuth(request *UploadInvoiceByAuthRequest) (_result *UploadInvoiceByAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UploadInvoiceByAuthHeaders{}
	_result = &UploadInvoiceByAuthResponse{}
	_body, _err := client.UploadInvoiceByAuthWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过手机号上传发票
//
// @param request - UploadInvoiceByMobileRequest
//
// @param headers - UploadInvoiceByMobileHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadInvoiceByMobileResponse
func (client *Client) UploadInvoiceByMobileWithOptions(request *UploadInvoiceByMobileRequest, headers *UploadInvoiceByMobileHeaders, runtime *util.RuntimeOptions) (_result *UploadInvoiceByMobileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Invoices)) {
		body["invoices"] = request.Invoices
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		body["mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.MobileStateCode)) {
		body["mobileStateCode"] = request.MobileStateCode
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
		Action:      tea.String("UploadInvoiceByMobile"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/invoices/mobiles/upload"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadInvoiceByMobileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过手机号上传发票
//
// @param request - UploadInvoiceByMobileRequest
//
// @return UploadInvoiceByMobileResponse
func (client *Client) UploadInvoiceByMobile(request *UploadInvoiceByMobileRequest) (_result *UploadInvoiceByMobileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UploadInvoiceByMobileHeaders{}
	_result = &UploadInvoiceByMobileResponse{}
	_body, _err := client.UploadInvoiceByMobileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上传进件中的图片获得图片链接
//
// @param request - UploadRegisterImageRequest
//
// @param headers - UploadRegisterImageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadRegisterImageResponse
func (client *Client) UploadRegisterImageWithOptions(request *UploadRegisterImageRequest, headers *UploadRegisterImageHeaders, runtime *util.RuntimeOptions) (_result *UploadRegisterImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageContent)) {
		body["imageContent"] = request.ImageContent
	}

	if !tea.BoolValue(util.IsUnset(request.ImageName)) {
		body["imageName"] = request.ImageName
	}

	if !tea.BoolValue(util.IsUnset(request.ImageType)) {
		body["imageType"] = request.ImageType
	}

	if !tea.BoolValue(util.IsUnset(request.InstId)) {
		body["instId"] = request.InstId
	}

	if !tea.BoolValue(util.IsUnset(request.PayChannel)) {
		body["payChannel"] = request.PayChannel
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
		Action:      tea.String("UploadRegisterImage"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/institutions/images"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadRegisterImageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传进件中的图片获得图片链接
//
// @param request - UploadRegisterImageRequest
//
// @return UploadRegisterImageResponse
func (client *Client) UploadRegisterImage(request *UploadRegisterImageRequest) (_result *UploadRegisterImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UploadRegisterImageHeaders{}
	_result = &UploadRegisterImageResponse{}
	_body, _err := client.UploadRegisterImageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用户协议签约
//
// @param request - UserAgreementPageSignRequest
//
// @param headers - UserAgreementPageSignHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UserAgreementPageSignResponse
func (client *Client) UserAgreementPageSignWithOptions(request *UserAgreementPageSignRequest, headers *UserAgreementPageSignHeaders, runtime *util.RuntimeOptions) (_result *UserAgreementPageSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		body["bizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.BizScene)) {
		body["bizScene"] = request.BizScene
	}

	if !tea.BoolValue(util.IsUnset(request.InstId)) {
		body["instId"] = request.InstId
	}

	if !tea.BoolValue(util.IsUnset(request.PayChannel)) {
		body["payChannel"] = request.PayChannel
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnUrl)) {
		body["returnUrl"] = request.ReturnUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SignScene)) {
		body["signScene"] = request.SignScene
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstId)) {
		body["subInstId"] = request.SubInstId
	}

	if !tea.BoolValue(util.IsUnset(request.SubMerchantName)) {
		body["subMerchantName"] = request.SubMerchantName
	}

	if !tea.BoolValue(util.IsUnset(request.SubMerchantServiceDesc)) {
		body["subMerchantServiceDesc"] = request.SubMerchantServiceDesc
	}

	if !tea.BoolValue(util.IsUnset(request.SubMerchantServiceName)) {
		body["subMerchantServiceName"] = request.SubMerchantServiceName
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
		Action:      tea.String("UserAgreementPageSign"),
		Version:     tea.String("finance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/finance/userAgreements"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UserAgreementPageSignResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用户协议签约
//
// @param request - UserAgreementPageSignRequest
//
// @return UserAgreementPageSignResponse
func (client *Client) UserAgreementPageSign(request *UserAgreementPageSignRequest) (_result *UserAgreementPageSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UserAgreementPageSignHeaders{}
	_result = &UserAgreementPageSignResponse{}
	_body, _err := client.UserAgreementPageSignWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
