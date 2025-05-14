// This file is auto-generated, don't edit it. Thanks.
package bizfinance_2_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BenefitMapValue struct {
	CanUse      *bool  `json:"canUse,omitempty" xml:"canUse,omitempty"`
	CanUseQuota *int64 `json:"canUseQuota,omitempty" xml:"canUseQuota,omitempty"`
	UsedQuota   *int64 `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
}

func (s BenefitMapValue) String() string {
	return tea.Prettify(s)
}

func (s BenefitMapValue) GoString() string {
	return s.String()
}

func (s *BenefitMapValue) SetCanUse(v bool) *BenefitMapValue {
	s.CanUse = &v
	return s
}

func (s *BenefitMapValue) SetCanUseQuota(v int64) *BenefitMapValue {
	s.CanUseQuota = &v
	return s
}

func (s *BenefitMapValue) SetUsedQuota(v int64) *BenefitMapValue {
	s.UsedQuota = &v
	return s
}

type AddFinanceEnterpriseAccountHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddFinanceEnterpriseAccountHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddFinanceEnterpriseAccountHeaders) GoString() string {
	return s.String()
}

func (s *AddFinanceEnterpriseAccountHeaders) SetCommonHeaders(v map[string]*string) *AddFinanceEnterpriseAccountHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddFinanceEnterpriseAccountHeaders) SetXAcsDingtalkAccessToken(v string) *AddFinanceEnterpriseAccountHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddFinanceEnterpriseAccountRequest struct {
	// example:
	//
	// 钉钉科技
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	// example:
	//
	// BANKCARD
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	BankCardNo  *string `json:"bankCardNo,omitempty" xml:"bankCardNo,omitempty"`
	// example:
	//
	// ICBC
	BankCode *string `json:"bankCode,omitempty" xml:"bankCode,omitempty"`
	// example:
	//
	// 中国工商银行
	BankName *string `json:"bankName,omitempty" xml:"bankName,omitempty"`
	// example:
	//
	// 杭州市
	City *string `json:"city,omitempty" xml:"city,omitempty"`
	// example:
	//
	// 账户描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 中国工商银行余杭分行
	OfficialName *string `json:"officialName,omitempty" xml:"officialName,omitempty"`
	// example:
	//
	// 1128878445
	OfficialNumber *string `json:"officialNumber,omitempty" xml:"officialNumber,omitempty"`
	// example:
	//
	// 浙江省
	Province *string `json:"province,omitempty" xml:"province,omitempty"`
	// example:
	//
	// smallScaleTaxpayer
	TaxNature *string `json:"taxNature,omitempty" xml:"taxNature,omitempty"`
	// example:
	//
	// 147581475814758
	TaxNo *string `json:"taxNo,omitempty" xml:"taxNo,omitempty"`
	// example:
	//
	// 5046195764756652
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AddFinanceEnterpriseAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFinanceEnterpriseAccountRequest) GoString() string {
	return s.String()
}

func (s *AddFinanceEnterpriseAccountRequest) SetAccountName(v string) *AddFinanceEnterpriseAccountRequest {
	s.AccountName = &v
	return s
}

func (s *AddFinanceEnterpriseAccountRequest) SetAccountType(v string) *AddFinanceEnterpriseAccountRequest {
	s.AccountType = &v
	return s
}

func (s *AddFinanceEnterpriseAccountRequest) SetBankCardNo(v string) *AddFinanceEnterpriseAccountRequest {
	s.BankCardNo = &v
	return s
}

func (s *AddFinanceEnterpriseAccountRequest) SetBankCode(v string) *AddFinanceEnterpriseAccountRequest {
	s.BankCode = &v
	return s
}

func (s *AddFinanceEnterpriseAccountRequest) SetBankName(v string) *AddFinanceEnterpriseAccountRequest {
	s.BankName = &v
	return s
}

func (s *AddFinanceEnterpriseAccountRequest) SetCity(v string) *AddFinanceEnterpriseAccountRequest {
	s.City = &v
	return s
}

func (s *AddFinanceEnterpriseAccountRequest) SetDescription(v string) *AddFinanceEnterpriseAccountRequest {
	s.Description = &v
	return s
}

func (s *AddFinanceEnterpriseAccountRequest) SetOfficialName(v string) *AddFinanceEnterpriseAccountRequest {
	s.OfficialName = &v
	return s
}

func (s *AddFinanceEnterpriseAccountRequest) SetOfficialNumber(v string) *AddFinanceEnterpriseAccountRequest {
	s.OfficialNumber = &v
	return s
}

func (s *AddFinanceEnterpriseAccountRequest) SetProvince(v string) *AddFinanceEnterpriseAccountRequest {
	s.Province = &v
	return s
}

func (s *AddFinanceEnterpriseAccountRequest) SetTaxNature(v string) *AddFinanceEnterpriseAccountRequest {
	s.TaxNature = &v
	return s
}

func (s *AddFinanceEnterpriseAccountRequest) SetTaxNo(v string) *AddFinanceEnterpriseAccountRequest {
	s.TaxNo = &v
	return s
}

func (s *AddFinanceEnterpriseAccountRequest) SetUserId(v string) *AddFinanceEnterpriseAccountRequest {
	s.UserId = &v
	return s
}

type AddFinanceEnterpriseAccountResponseBody struct {
	AccountCode *string `json:"accountCode,omitempty" xml:"accountCode,omitempty"`
}

func (s AddFinanceEnterpriseAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFinanceEnterpriseAccountResponseBody) GoString() string {
	return s.String()
}

func (s *AddFinanceEnterpriseAccountResponseBody) SetAccountCode(v string) *AddFinanceEnterpriseAccountResponseBody {
	s.AccountCode = &v
	return s
}

type AddFinanceEnterpriseAccountResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFinanceEnterpriseAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFinanceEnterpriseAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s AddFinanceEnterpriseAccountResponse) GoString() string {
	return s.String()
}

func (s *AddFinanceEnterpriseAccountResponse) SetHeaders(v map[string]*string) *AddFinanceEnterpriseAccountResponse {
	s.Headers = v
	return s
}

func (s *AddFinanceEnterpriseAccountResponse) SetStatusCode(v int32) *AddFinanceEnterpriseAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFinanceEnterpriseAccountResponse) SetBody(v *AddFinanceEnterpriseAccountResponseBody) *AddFinanceEnterpriseAccountResponse {
	s.Body = v
	return s
}

type BankGatewayInvokeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BankGatewayInvokeHeaders) String() string {
	return tea.Prettify(s)
}

func (s BankGatewayInvokeHeaders) GoString() string {
	return s.String()
}

func (s *BankGatewayInvokeHeaders) SetCommonHeaders(v map[string]*string) *BankGatewayInvokeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BankGatewayInvokeHeaders) SetXAcsDingtalkAccessToken(v string) *BankGatewayInvokeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BankGatewayInvokeRequest struct {
	// example:
	//
	// bankHttp
	ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	InputData  *string `json:"inputData,omitempty" xml:"inputData,omitempty"`
	// example:
	//
	// https://cdc.cmbchina.com/cdcserver/api/v2
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s BankGatewayInvokeRequest) String() string {
	return tea.Prettify(s)
}

func (s BankGatewayInvokeRequest) GoString() string {
	return s.String()
}

func (s *BankGatewayInvokeRequest) SetActionType(v string) *BankGatewayInvokeRequest {
	s.ActionType = &v
	return s
}

func (s *BankGatewayInvokeRequest) SetInputData(v string) *BankGatewayInvokeRequest {
	s.InputData = &v
	return s
}

func (s *BankGatewayInvokeRequest) SetUrl(v string) *BankGatewayInvokeRequest {
	s.Url = &v
	return s
}

type BankGatewayInvokeResponseBody struct {
	OutputData *string `json:"outputData,omitempty" xml:"outputData,omitempty"`
}

func (s BankGatewayInvokeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BankGatewayInvokeResponseBody) GoString() string {
	return s.String()
}

func (s *BankGatewayInvokeResponseBody) SetOutputData(v string) *BankGatewayInvokeResponseBody {
	s.OutputData = &v
	return s
}

type BankGatewayInvokeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BankGatewayInvokeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BankGatewayInvokeResponse) String() string {
	return tea.Prettify(s)
}

func (s BankGatewayInvokeResponse) GoString() string {
	return s.String()
}

func (s *BankGatewayInvokeResponse) SetHeaders(v map[string]*string) *BankGatewayInvokeResponse {
	s.Headers = v
	return s
}

func (s *BankGatewayInvokeResponse) SetStatusCode(v int32) *BankGatewayInvokeResponse {
	s.StatusCode = &v
	return s
}

func (s *BankGatewayInvokeResponse) SetBody(v *BankGatewayInvokeResponseBody) *BankGatewayInvokeResponse {
	s.Body = v
	return s
}

type BatchDeleteReceiptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchDeleteReceiptHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteReceiptHeaders) GoString() string {
	return s.String()
}

func (s *BatchDeleteReceiptHeaders) SetCommonHeaders(v map[string]*string) *BatchDeleteReceiptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchDeleteReceiptHeaders) SetXAcsDingtalkAccessToken(v string) *BatchDeleteReceiptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchDeleteReceiptRequest struct {
	InstanceIdList []*string `json:"instanceIdList,omitempty" xml:"instanceIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 504XX
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
}

func (s BatchDeleteReceiptRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteReceiptRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteReceiptRequest) SetInstanceIdList(v []*string) *BatchDeleteReceiptRequest {
	s.InstanceIdList = v
	return s
}

func (s *BatchDeleteReceiptRequest) SetOperator(v string) *BatchDeleteReceiptRequest {
	s.Operator = &v
	return s
}

type BatchDeleteReceiptResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s BatchDeleteReceiptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteReceiptResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteReceiptResponseBody) SetResult(v bool) *BatchDeleteReceiptResponseBody {
	s.Result = &v
	return s
}

type BatchDeleteReceiptResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteReceiptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteReceiptResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteReceiptResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteReceiptResponse) SetHeaders(v map[string]*string) *BatchDeleteReceiptResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteReceiptResponse) SetStatusCode(v int32) *BatchDeleteReceiptResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteReceiptResponse) SetBody(v *BatchDeleteReceiptResponseBody) *BatchDeleteReceiptResponse {
	s.Body = v
	return s
}

type BatchQueryOrgInvoiceUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchQueryOrgInvoiceUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryOrgInvoiceUrlHeaders) GoString() string {
	return s.String()
}

func (s *BatchQueryOrgInvoiceUrlHeaders) SetCommonHeaders(v map[string]*string) *BatchQueryOrgInvoiceUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchQueryOrgInvoiceUrlHeaders) SetXAcsDingtalkAccessToken(v string) *BatchQueryOrgInvoiceUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchQueryOrgInvoiceUrlRequest struct {
	CompanyCode      *string                                           `json:"companyCode,omitempty" xml:"companyCode,omitempty"`
	InvoiceKeyVOList []*BatchQueryOrgInvoiceUrlRequestInvoiceKeyVOList `json:"invoiceKeyVOList,omitempty" xml:"invoiceKeyVOList,omitempty" type:"Repeated"`
	Operator         *string                                           `json:"operator,omitempty" xml:"operator,omitempty"`
}

func (s BatchQueryOrgInvoiceUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryOrgInvoiceUrlRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryOrgInvoiceUrlRequest) SetCompanyCode(v string) *BatchQueryOrgInvoiceUrlRequest {
	s.CompanyCode = &v
	return s
}

func (s *BatchQueryOrgInvoiceUrlRequest) SetInvoiceKeyVOList(v []*BatchQueryOrgInvoiceUrlRequestInvoiceKeyVOList) *BatchQueryOrgInvoiceUrlRequest {
	s.InvoiceKeyVOList = v
	return s
}

func (s *BatchQueryOrgInvoiceUrlRequest) SetOperator(v string) *BatchQueryOrgInvoiceUrlRequest {
	s.Operator = &v
	return s
}

type BatchQueryOrgInvoiceUrlRequestInvoiceKeyVOList struct {
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	InvoiceNo   *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
}

func (s BatchQueryOrgInvoiceUrlRequestInvoiceKeyVOList) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryOrgInvoiceUrlRequestInvoiceKeyVOList) GoString() string {
	return s.String()
}

func (s *BatchQueryOrgInvoiceUrlRequestInvoiceKeyVOList) SetInvoiceCode(v string) *BatchQueryOrgInvoiceUrlRequestInvoiceKeyVOList {
	s.InvoiceCode = &v
	return s
}

func (s *BatchQueryOrgInvoiceUrlRequestInvoiceKeyVOList) SetInvoiceNo(v string) *BatchQueryOrgInvoiceUrlRequestInvoiceKeyVOList {
	s.InvoiceNo = &v
	return s
}

type BatchQueryOrgInvoiceUrlResponseBody struct {
	FailInvoiceList   []*BatchQueryOrgInvoiceUrlResponseBodyFailInvoiceList   `json:"failInvoiceList,omitempty" xml:"failInvoiceList,omitempty" type:"Repeated"`
	OrgInvoiceUrlList []*BatchQueryOrgInvoiceUrlResponseBodyOrgInvoiceUrlList `json:"orgInvoiceUrlList,omitempty" xml:"orgInvoiceUrlList,omitempty" type:"Repeated"`
}

func (s BatchQueryOrgInvoiceUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryOrgInvoiceUrlResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQueryOrgInvoiceUrlResponseBody) SetFailInvoiceList(v []*BatchQueryOrgInvoiceUrlResponseBodyFailInvoiceList) *BatchQueryOrgInvoiceUrlResponseBody {
	s.FailInvoiceList = v
	return s
}

func (s *BatchQueryOrgInvoiceUrlResponseBody) SetOrgInvoiceUrlList(v []*BatchQueryOrgInvoiceUrlResponseBodyOrgInvoiceUrlList) *BatchQueryOrgInvoiceUrlResponseBody {
	s.OrgInvoiceUrlList = v
	return s
}

type BatchQueryOrgInvoiceUrlResponseBodyFailInvoiceList struct {
	ErrorMsg    *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	InvoiceNo   *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
}

func (s BatchQueryOrgInvoiceUrlResponseBodyFailInvoiceList) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryOrgInvoiceUrlResponseBodyFailInvoiceList) GoString() string {
	return s.String()
}

func (s *BatchQueryOrgInvoiceUrlResponseBodyFailInvoiceList) SetErrorMsg(v string) *BatchQueryOrgInvoiceUrlResponseBodyFailInvoiceList {
	s.ErrorMsg = &v
	return s
}

func (s *BatchQueryOrgInvoiceUrlResponseBodyFailInvoiceList) SetInvoiceCode(v string) *BatchQueryOrgInvoiceUrlResponseBodyFailInvoiceList {
	s.InvoiceCode = &v
	return s
}

func (s *BatchQueryOrgInvoiceUrlResponseBodyFailInvoiceList) SetInvoiceNo(v string) *BatchQueryOrgInvoiceUrlResponseBodyFailInvoiceList {
	s.InvoiceNo = &v
	return s
}

type BatchQueryOrgInvoiceUrlResponseBodyOrgInvoiceUrlList struct {
	InvoiceCode    *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	InvoiceNo      *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
	OfdUrl         *string `json:"ofdUrl,omitempty" xml:"ofdUrl,omitempty"`
	OriginFileType *string `json:"originFileType,omitempty" xml:"originFileType,omitempty"`
	OriginFileUrl  *string `json:"originFileUrl,omitempty" xml:"originFileUrl,omitempty"`
	PdfUrl         *string `json:"pdfUrl,omitempty" xml:"pdfUrl,omitempty"`
	XmlUrl         *string `json:"xmlUrl,omitempty" xml:"xmlUrl,omitempty"`
}

func (s BatchQueryOrgInvoiceUrlResponseBodyOrgInvoiceUrlList) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryOrgInvoiceUrlResponseBodyOrgInvoiceUrlList) GoString() string {
	return s.String()
}

func (s *BatchQueryOrgInvoiceUrlResponseBodyOrgInvoiceUrlList) SetInvoiceCode(v string) *BatchQueryOrgInvoiceUrlResponseBodyOrgInvoiceUrlList {
	s.InvoiceCode = &v
	return s
}

func (s *BatchQueryOrgInvoiceUrlResponseBodyOrgInvoiceUrlList) SetInvoiceNo(v string) *BatchQueryOrgInvoiceUrlResponseBodyOrgInvoiceUrlList {
	s.InvoiceNo = &v
	return s
}

func (s *BatchQueryOrgInvoiceUrlResponseBodyOrgInvoiceUrlList) SetOfdUrl(v string) *BatchQueryOrgInvoiceUrlResponseBodyOrgInvoiceUrlList {
	s.OfdUrl = &v
	return s
}

func (s *BatchQueryOrgInvoiceUrlResponseBodyOrgInvoiceUrlList) SetOriginFileType(v string) *BatchQueryOrgInvoiceUrlResponseBodyOrgInvoiceUrlList {
	s.OriginFileType = &v
	return s
}

func (s *BatchQueryOrgInvoiceUrlResponseBodyOrgInvoiceUrlList) SetOriginFileUrl(v string) *BatchQueryOrgInvoiceUrlResponseBodyOrgInvoiceUrlList {
	s.OriginFileUrl = &v
	return s
}

func (s *BatchQueryOrgInvoiceUrlResponseBodyOrgInvoiceUrlList) SetPdfUrl(v string) *BatchQueryOrgInvoiceUrlResponseBodyOrgInvoiceUrlList {
	s.PdfUrl = &v
	return s
}

func (s *BatchQueryOrgInvoiceUrlResponseBodyOrgInvoiceUrlList) SetXmlUrl(v string) *BatchQueryOrgInvoiceUrlResponseBodyOrgInvoiceUrlList {
	s.XmlUrl = &v
	return s
}

type BatchQueryOrgInvoiceUrlResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchQueryOrgInvoiceUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchQueryOrgInvoiceUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryOrgInvoiceUrlResponse) GoString() string {
	return s.String()
}

func (s *BatchQueryOrgInvoiceUrlResponse) SetHeaders(v map[string]*string) *BatchQueryOrgInvoiceUrlResponse {
	s.Headers = v
	return s
}

func (s *BatchQueryOrgInvoiceUrlResponse) SetStatusCode(v int32) *BatchQueryOrgInvoiceUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchQueryOrgInvoiceUrlResponse) SetBody(v *BatchQueryOrgInvoiceUrlResponseBody) *BatchQueryOrgInvoiceUrlResponse {
	s.Body = v
	return s
}

type BatchQueryPaymentRecallFileHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchQueryPaymentRecallFileHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryPaymentRecallFileHeaders) GoString() string {
	return s.String()
}

func (s *BatchQueryPaymentRecallFileHeaders) SetCommonHeaders(v map[string]*string) *BatchQueryPaymentRecallFileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchQueryPaymentRecallFileHeaders) SetXAcsDingtalkAccessToken(v string) *BatchQueryPaymentRecallFileHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchQueryPaymentRecallFileRequest struct {
	DetailIdList []*string `json:"detailIdList,omitempty" xml:"detailIdList,omitempty" type:"Repeated"`
	Operator     *string   `json:"operator,omitempty" xml:"operator,omitempty"`
}

func (s BatchQueryPaymentRecallFileRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryPaymentRecallFileRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryPaymentRecallFileRequest) SetDetailIdList(v []*string) *BatchQueryPaymentRecallFileRequest {
	s.DetailIdList = v
	return s
}

func (s *BatchQueryPaymentRecallFileRequest) SetOperator(v string) *BatchQueryPaymentRecallFileRequest {
	s.Operator = &v
	return s
}

type BatchQueryPaymentRecallFileResponseBody struct {
	PaymentRecallFileList []*BatchQueryPaymentRecallFileResponseBodyPaymentRecallFileList `json:"paymentRecallFileList,omitempty" xml:"paymentRecallFileList,omitempty" type:"Repeated"`
}

func (s BatchQueryPaymentRecallFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryPaymentRecallFileResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQueryPaymentRecallFileResponseBody) SetPaymentRecallFileList(v []*BatchQueryPaymentRecallFileResponseBodyPaymentRecallFileList) *BatchQueryPaymentRecallFileResponseBody {
	s.PaymentRecallFileList = v
	return s
}

type BatchQueryPaymentRecallFileResponseBodyPaymentRecallFileList struct {
	DetailId      *string `json:"detailId,omitempty" xml:"detailId,omitempty"`
	FileId        *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	FileName      *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileSize      *string `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	FileType      *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	RecallFileUrl *string `json:"recallFileUrl,omitempty" xml:"recallFileUrl,omitempty"`
	SpaceId       *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s BatchQueryPaymentRecallFileResponseBodyPaymentRecallFileList) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryPaymentRecallFileResponseBodyPaymentRecallFileList) GoString() string {
	return s.String()
}

func (s *BatchQueryPaymentRecallFileResponseBodyPaymentRecallFileList) SetDetailId(v string) *BatchQueryPaymentRecallFileResponseBodyPaymentRecallFileList {
	s.DetailId = &v
	return s
}

func (s *BatchQueryPaymentRecallFileResponseBodyPaymentRecallFileList) SetFileId(v string) *BatchQueryPaymentRecallFileResponseBodyPaymentRecallFileList {
	s.FileId = &v
	return s
}

func (s *BatchQueryPaymentRecallFileResponseBodyPaymentRecallFileList) SetFileName(v string) *BatchQueryPaymentRecallFileResponseBodyPaymentRecallFileList {
	s.FileName = &v
	return s
}

func (s *BatchQueryPaymentRecallFileResponseBodyPaymentRecallFileList) SetFileSize(v string) *BatchQueryPaymentRecallFileResponseBodyPaymentRecallFileList {
	s.FileSize = &v
	return s
}

func (s *BatchQueryPaymentRecallFileResponseBodyPaymentRecallFileList) SetFileType(v string) *BatchQueryPaymentRecallFileResponseBodyPaymentRecallFileList {
	s.FileType = &v
	return s
}

func (s *BatchQueryPaymentRecallFileResponseBodyPaymentRecallFileList) SetRecallFileUrl(v string) *BatchQueryPaymentRecallFileResponseBodyPaymentRecallFileList {
	s.RecallFileUrl = &v
	return s
}

func (s *BatchQueryPaymentRecallFileResponseBodyPaymentRecallFileList) SetSpaceId(v string) *BatchQueryPaymentRecallFileResponseBodyPaymentRecallFileList {
	s.SpaceId = &v
	return s
}

type BatchQueryPaymentRecallFileResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchQueryPaymentRecallFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchQueryPaymentRecallFileResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryPaymentRecallFileResponse) GoString() string {
	return s.String()
}

func (s *BatchQueryPaymentRecallFileResponse) SetHeaders(v map[string]*string) *BatchQueryPaymentRecallFileResponse {
	s.Headers = v
	return s
}

func (s *BatchQueryPaymentRecallFileResponse) SetStatusCode(v int32) *BatchQueryPaymentRecallFileResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchQueryPaymentRecallFileResponse) SetBody(v *BatchQueryPaymentRecallFileResponseBody) *BatchQueryPaymentRecallFileResponse {
	s.Body = v
	return s
}

type BatchSyncBankReceiptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchSyncBankReceiptHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchSyncBankReceiptHeaders) GoString() string {
	return s.String()
}

func (s *BatchSyncBankReceiptHeaders) SetCommonHeaders(v map[string]*string) *BatchSyncBankReceiptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchSyncBankReceiptHeaders) SetXAcsDingtalkAccessToken(v string) *BatchSyncBankReceiptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchSyncBankReceiptRequest struct {
	Body []*BatchSyncBankReceiptRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s BatchSyncBankReceiptRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSyncBankReceiptRequest) GoString() string {
	return s.String()
}

func (s *BatchSyncBankReceiptRequest) SetBody(v []*BatchSyncBankReceiptRequestBody) *BatchSyncBankReceiptRequest {
	s.Body = v
	return s
}

type BatchSyncBankReceiptRequestBody struct {
	// example:
	//
	// https://xxxxx
	FileDownloadUrl *string `json:"fileDownloadUrl,omitempty" xml:"fileDownloadUrl,omitempty"`
	// example:
	//
	// xxxx回单.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// 2024000001902335
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// example:
	//
	// detailId
	MessageIdType *string `json:"messageIdType,omitempty" xml:"messageIdType,omitempty"`
}

func (s BatchSyncBankReceiptRequestBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSyncBankReceiptRequestBody) GoString() string {
	return s.String()
}

func (s *BatchSyncBankReceiptRequestBody) SetFileDownloadUrl(v string) *BatchSyncBankReceiptRequestBody {
	s.FileDownloadUrl = &v
	return s
}

func (s *BatchSyncBankReceiptRequestBody) SetFileName(v string) *BatchSyncBankReceiptRequestBody {
	s.FileName = &v
	return s
}

func (s *BatchSyncBankReceiptRequestBody) SetMessageId(v string) *BatchSyncBankReceiptRequestBody {
	s.MessageId = &v
	return s
}

func (s *BatchSyncBankReceiptRequestBody) SetMessageIdType(v string) *BatchSyncBankReceiptRequestBody {
	s.MessageIdType = &v
	return s
}

type BatchSyncBankReceiptResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s BatchSyncBankReceiptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSyncBankReceiptResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSyncBankReceiptResponseBody) SetResult(v bool) *BatchSyncBankReceiptResponseBody {
	s.Result = &v
	return s
}

type BatchSyncBankReceiptResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchSyncBankReceiptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchSyncBankReceiptResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchSyncBankReceiptResponse) GoString() string {
	return s.String()
}

func (s *BatchSyncBankReceiptResponse) SetHeaders(v map[string]*string) *BatchSyncBankReceiptResponse {
	s.Headers = v
	return s
}

func (s *BatchSyncBankReceiptResponse) SetStatusCode(v int32) *BatchSyncBankReceiptResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSyncBankReceiptResponse) SetBody(v *BatchSyncBankReceiptResponseBody) *BatchSyncBankReceiptResponse {
	s.Body = v
	return s
}

type CheckVoucherStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CheckVoucherStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s CheckVoucherStatusHeaders) GoString() string {
	return s.String()
}

func (s *CheckVoucherStatusHeaders) SetCommonHeaders(v map[string]*string) *CheckVoucherStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CheckVoucherStatusHeaders) SetXAcsDingtalkAccessToken(v string) *CheckVoucherStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CheckVoucherStatusRequest struct {
	// example:
	//
	// COM_DEFUALT
	CompanyCode *string `json:"companyCode,omitempty" xml:"companyCode,omitempty"`
	// example:
	//
	// 1631526550994
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// VAT_IN
	FinanceType *string `json:"financeType,omitempty" xml:"financeType,omitempty"`
	// example:
	//
	// 3121234560
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	// example:
	//
	// 1234567890
	InvoiceNo *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1631526550994
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 12345678901
	TaxNo *string `json:"taxNo,omitempty" xml:"taxNo,omitempty"`
	// example:
	//
	// VERIFIED
	VerifyStatus *string `json:"verifyStatus,omitempty" xml:"verifyStatus,omitempty"`
}

func (s CheckVoucherStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckVoucherStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckVoucherStatusRequest) SetCompanyCode(v string) *CheckVoucherStatusRequest {
	s.CompanyCode = &v
	return s
}

func (s *CheckVoucherStatusRequest) SetEndTime(v int64) *CheckVoucherStatusRequest {
	s.EndTime = &v
	return s
}

func (s *CheckVoucherStatusRequest) SetFinanceType(v string) *CheckVoucherStatusRequest {
	s.FinanceType = &v
	return s
}

func (s *CheckVoucherStatusRequest) SetInvoiceCode(v string) *CheckVoucherStatusRequest {
	s.InvoiceCode = &v
	return s
}

func (s *CheckVoucherStatusRequest) SetInvoiceNo(v string) *CheckVoucherStatusRequest {
	s.InvoiceNo = &v
	return s
}

func (s *CheckVoucherStatusRequest) SetPageNumber(v int64) *CheckVoucherStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *CheckVoucherStatusRequest) SetPageSize(v int64) *CheckVoucherStatusRequest {
	s.PageSize = &v
	return s
}

func (s *CheckVoucherStatusRequest) SetStartTime(v int64) *CheckVoucherStatusRequest {
	s.StartTime = &v
	return s
}

func (s *CheckVoucherStatusRequest) SetTaxNo(v string) *CheckVoucherStatusRequest {
	s.TaxNo = &v
	return s
}

func (s *CheckVoucherStatusRequest) SetVerifyStatus(v string) *CheckVoucherStatusRequest {
	s.VerifyStatus = &v
	return s
}

type CheckVoucherStatusResponseBody struct {
	Result  *bool `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CheckVoucherStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckVoucherStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckVoucherStatusResponseBody) SetResult(v bool) *CheckVoucherStatusResponseBody {
	s.Result = &v
	return s
}

func (s *CheckVoucherStatusResponseBody) SetSuccess(v bool) *CheckVoucherStatusResponseBody {
	s.Success = &v
	return s
}

type CheckVoucherStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckVoucherStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckVoucherStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckVoucherStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckVoucherStatusResponse) SetHeaders(v map[string]*string) *CheckVoucherStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckVoucherStatusResponse) SetStatusCode(v int32) *CheckVoucherStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckVoucherStatusResponse) SetBody(v *CheckVoucherStatusResponseBody) *CheckVoucherStatusResponse {
	s.Body = v
	return s
}

type ConfirmPaymentOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ConfirmPaymentOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s ConfirmPaymentOrderHeaders) GoString() string {
	return s.String()
}

func (s *ConfirmPaymentOrderHeaders) SetCommonHeaders(v map[string]*string) *ConfirmPaymentOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ConfirmPaymentOrderHeaders) SetXAcsDingtalkAccessToken(v string) *ConfirmPaymentOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ConfirmPaymentOrderRequest struct {
	OutBizNoList []*string `json:"outBizNoList,omitempty" xml:"outBizNoList,omitempty" type:"Repeated"`
	// example:
	//
	// 5041123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ConfirmPaymentOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfirmPaymentOrderRequest) GoString() string {
	return s.String()
}

func (s *ConfirmPaymentOrderRequest) SetOutBizNoList(v []*string) *ConfirmPaymentOrderRequest {
	s.OutBizNoList = v
	return s
}

func (s *ConfirmPaymentOrderRequest) SetUserId(v string) *ConfirmPaymentOrderRequest {
	s.UserId = &v
	return s
}

type ConfirmPaymentOrderResponseBody struct {
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ConfirmPaymentOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfirmPaymentOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmPaymentOrderResponseBody) SetUrl(v string) *ConfirmPaymentOrderResponseBody {
	s.Url = &v
	return s
}

type ConfirmPaymentOrderResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfirmPaymentOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmPaymentOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfirmPaymentOrderResponse) GoString() string {
	return s.String()
}

func (s *ConfirmPaymentOrderResponse) SetHeaders(v map[string]*string) *ConfirmPaymentOrderResponse {
	s.Headers = v
	return s
}

func (s *ConfirmPaymentOrderResponse) SetStatusCode(v int32) *ConfirmPaymentOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmPaymentOrderResponse) SetBody(v *ConfirmPaymentOrderResponseBody) *ConfirmPaymentOrderResponse {
	s.Body = v
	return s
}

type CreateCollectionOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateCollectionOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateCollectionOrderHeaders) GoString() string {
	return s.String()
}

func (s *CreateCollectionOrderHeaders) SetCommonHeaders(v map[string]*string) *CreateCollectionOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateCollectionOrderHeaders) SetXAcsDingtalkAccessToken(v string) *CreateCollectionOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateCollectionOrderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0.01
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a
	CollectionInfoId *string `json:"collectionInfoId,omitempty" xml:"collectionInfoId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 收款
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s CreateCollectionOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCollectionOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateCollectionOrderRequest) SetAmount(v string) *CreateCollectionOrderRequest {
	s.Amount = &v
	return s
}

func (s *CreateCollectionOrderRequest) SetCollectionInfoId(v string) *CreateCollectionOrderRequest {
	s.CollectionInfoId = &v
	return s
}

func (s *CreateCollectionOrderRequest) SetInstanceId(v string) *CreateCollectionOrderRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCollectionOrderRequest) SetRemark(v string) *CreateCollectionOrderRequest {
	s.Remark = &v
	return s
}

type CreateCollectionOrderResponseBody struct {
	CollectionUrl *string `json:"collectionUrl,omitempty" xml:"collectionUrl,omitempty"`
}

func (s CreateCollectionOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCollectionOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCollectionOrderResponseBody) SetCollectionUrl(v string) *CreateCollectionOrderResponseBody {
	s.CollectionUrl = &v
	return s
}

type CreateCollectionOrderResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCollectionOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCollectionOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCollectionOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateCollectionOrderResponse) SetHeaders(v map[string]*string) *CreateCollectionOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateCollectionOrderResponse) SetStatusCode(v int32) *CreateCollectionOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCollectionOrderResponse) SetBody(v *CreateCollectionOrderResponseBody) *CreateCollectionOrderResponse {
	s.Body = v
	return s
}

type CreatePaymentOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreatePaymentOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreatePaymentOrderHeaders) GoString() string {
	return s.String()
}

func (s *CreatePaymentOrderHeaders) SetCommonHeaders(v map[string]*string) *CreatePaymentOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreatePaymentOrderHeaders) SetXAcsDingtalkAccessToken(v string) *CreatePaymentOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreatePaymentOrderRequest struct {
	// example:
	//
	// 100
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// example:
	//
	// 1741941518884
	ExpireTime *int64 `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// example:
	//
	// xxxxabc
	OutBizNo        *string                                   `json:"outBizNo,omitempty" xml:"outBizNo,omitempty"`
	PayeeAccountDTO *CreatePaymentOrderRequestPayeeAccountDTO `json:"payeeAccountDTO,omitempty" xml:"payeeAccountDTO,omitempty" type:"Struct"`
	PayerAccountDTO *CreatePaymentOrderRequestPayerAccountDTO `json:"payerAccountDTO,omitempty" xml:"payerAccountDTO,omitempty" type:"Struct"`
	// example:
	//
	// 日常报销
	PaymentOrderTitle *string `json:"paymentOrderTitle,omitempty" xml:"paymentOrderTitle,omitempty"`
	// example:
	//
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// 付款用途
	Usage *string `json:"usage,omitempty" xml:"usage,omitempty"`
	// example:
	//
	// 5046195764756652
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreatePaymentOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePaymentOrderRequest) GoString() string {
	return s.String()
}

func (s *CreatePaymentOrderRequest) SetAmount(v string) *CreatePaymentOrderRequest {
	s.Amount = &v
	return s
}

func (s *CreatePaymentOrderRequest) SetExpireTime(v int64) *CreatePaymentOrderRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreatePaymentOrderRequest) SetOutBizNo(v string) *CreatePaymentOrderRequest {
	s.OutBizNo = &v
	return s
}

func (s *CreatePaymentOrderRequest) SetPayeeAccountDTO(v *CreatePaymentOrderRequestPayeeAccountDTO) *CreatePaymentOrderRequest {
	s.PayeeAccountDTO = v
	return s
}

func (s *CreatePaymentOrderRequest) SetPayerAccountDTO(v *CreatePaymentOrderRequestPayerAccountDTO) *CreatePaymentOrderRequest {
	s.PayerAccountDTO = v
	return s
}

func (s *CreatePaymentOrderRequest) SetPaymentOrderTitle(v string) *CreatePaymentOrderRequest {
	s.PaymentOrderTitle = &v
	return s
}

func (s *CreatePaymentOrderRequest) SetRemark(v string) *CreatePaymentOrderRequest {
	s.Remark = &v
	return s
}

func (s *CreatePaymentOrderRequest) SetUsage(v string) *CreatePaymentOrderRequest {
	s.Usage = &v
	return s
}

func (s *CreatePaymentOrderRequest) SetUserId(v string) *CreatePaymentOrderRequest {
	s.UserId = &v
	return s
}

type CreatePaymentOrderRequestPayeeAccountDTO struct {
	BankOpenDTO *CreatePaymentOrderRequestPayeeAccountDTOBankOpenDTO `json:"bankOpenDTO,omitempty" xml:"bankOpenDTO,omitempty" type:"Struct"`
}

func (s CreatePaymentOrderRequestPayeeAccountDTO) String() string {
	return tea.Prettify(s)
}

func (s CreatePaymentOrderRequestPayeeAccountDTO) GoString() string {
	return s.String()
}

func (s *CreatePaymentOrderRequestPayeeAccountDTO) SetBankOpenDTO(v *CreatePaymentOrderRequestPayeeAccountDTOBankOpenDTO) *CreatePaymentOrderRequestPayeeAccountDTO {
	s.BankOpenDTO = v
	return s
}

type CreatePaymentOrderRequestPayeeAccountDTOBankOpenDTO struct {
	// example:
	//
	// 钉钉中国
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	// example:
	//
	// 1025884624512
	BankBranchCode *string `json:"bankBranchCode,omitempty" xml:"bankBranchCode,omitempty"`
	// example:
	//
	// 招商银行杭州余杭分行
	BankBranchName *string `json:"bankBranchName,omitempty" xml:"bankBranchName,omitempty"`
	// example:
	//
	// 122215111012
	BankCardNo *string `json:"bankCardNo,omitempty" xml:"bankCardNo,omitempty"`
	// example:
	//
	// ICBC
	BankCode *string `json:"bankCode,omitempty" xml:"bankCode,omitempty"`
	// example:
	//
	// 招商银行
	BankName *string `json:"bankName,omitempty" xml:"bankName,omitempty"`
	// example:
	//
	// PERSONAL_BANK_CARD
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreatePaymentOrderRequestPayeeAccountDTOBankOpenDTO) String() string {
	return tea.Prettify(s)
}

func (s CreatePaymentOrderRequestPayeeAccountDTOBankOpenDTO) GoString() string {
	return s.String()
}

func (s *CreatePaymentOrderRequestPayeeAccountDTOBankOpenDTO) SetAccountName(v string) *CreatePaymentOrderRequestPayeeAccountDTOBankOpenDTO {
	s.AccountName = &v
	return s
}

func (s *CreatePaymentOrderRequestPayeeAccountDTOBankOpenDTO) SetBankBranchCode(v string) *CreatePaymentOrderRequestPayeeAccountDTOBankOpenDTO {
	s.BankBranchCode = &v
	return s
}

func (s *CreatePaymentOrderRequestPayeeAccountDTOBankOpenDTO) SetBankBranchName(v string) *CreatePaymentOrderRequestPayeeAccountDTOBankOpenDTO {
	s.BankBranchName = &v
	return s
}

func (s *CreatePaymentOrderRequestPayeeAccountDTOBankOpenDTO) SetBankCardNo(v string) *CreatePaymentOrderRequestPayeeAccountDTOBankOpenDTO {
	s.BankCardNo = &v
	return s
}

func (s *CreatePaymentOrderRequestPayeeAccountDTOBankOpenDTO) SetBankCode(v string) *CreatePaymentOrderRequestPayeeAccountDTOBankOpenDTO {
	s.BankCode = &v
	return s
}

func (s *CreatePaymentOrderRequestPayeeAccountDTOBankOpenDTO) SetBankName(v string) *CreatePaymentOrderRequestPayeeAccountDTOBankOpenDTO {
	s.BankName = &v
	return s
}

func (s *CreatePaymentOrderRequestPayeeAccountDTOBankOpenDTO) SetType(v string) *CreatePaymentOrderRequestPayeeAccountDTOBankOpenDTO {
	s.Type = &v
	return s
}

type CreatePaymentOrderRequestPayerAccountDTO struct {
	// example:
	//
	// ACC_XXXXX
	EnterpriseAccountCode *string `json:"enterpriseAccountCode,omitempty" xml:"enterpriseAccountCode,omitempty"`
}

func (s CreatePaymentOrderRequestPayerAccountDTO) String() string {
	return tea.Prettify(s)
}

func (s CreatePaymentOrderRequestPayerAccountDTO) GoString() string {
	return s.String()
}

func (s *CreatePaymentOrderRequestPayerAccountDTO) SetEnterpriseAccountCode(v string) *CreatePaymentOrderRequestPayerAccountDTO {
	s.EnterpriseAccountCode = &v
	return s
}

type CreatePaymentOrderResponseBody struct {
	ExpireTime *int64  `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	OrderNo    *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	OutBizNo   *string `json:"outBizNo,omitempty" xml:"outBizNo,omitempty"`
}

func (s CreatePaymentOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePaymentOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePaymentOrderResponseBody) SetExpireTime(v int64) *CreatePaymentOrderResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *CreatePaymentOrderResponseBody) SetOrderNo(v string) *CreatePaymentOrderResponseBody {
	s.OrderNo = &v
	return s
}

func (s *CreatePaymentOrderResponseBody) SetOutBizNo(v string) *CreatePaymentOrderResponseBody {
	s.OutBizNo = &v
	return s
}

type CreatePaymentOrderResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePaymentOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePaymentOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePaymentOrderResponse) GoString() string {
	return s.String()
}

func (s *CreatePaymentOrderResponse) SetHeaders(v map[string]*string) *CreatePaymentOrderResponse {
	s.Headers = v
	return s
}

func (s *CreatePaymentOrderResponse) SetStatusCode(v int32) *CreatePaymentOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePaymentOrderResponse) SetBody(v *CreatePaymentOrderResponseBody) *CreatePaymentOrderResponse {
	s.Body = v
	return s
}

type GetCategoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCategoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCategoryHeaders) GoString() string {
	return s.String()
}

func (s *GetCategoryHeaders) SetCommonHeaders(v map[string]*string) *GetCategoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCategoryHeaders) SetXAcsDingtalkAccessToken(v string) *GetCategoryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCategoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// INCOME_XXX
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s GetCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCategoryRequest) GoString() string {
	return s.String()
}

func (s *GetCategoryRequest) SetCode(v string) *GetCategoryRequest {
	s.Code = &v
	return s
}

type GetCategoryResponseBody struct {
	AccountantBookIdList []*string `json:"accountantBookIdList,omitempty" xml:"accountantBookIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// INCOME_XXX
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	IsDir *bool `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 汽车
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// DIR_XXX
	ParentCode *string `json:"parentCode,omitempty" xml:"parentCode,omitempty"`
	Remark     *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// valid
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// income
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetCategoryResponseBody) SetAccountantBookIdList(v []*string) *GetCategoryResponseBody {
	s.AccountantBookIdList = v
	return s
}

func (s *GetCategoryResponseBody) SetCode(v string) *GetCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *GetCategoryResponseBody) SetIsDir(v bool) *GetCategoryResponseBody {
	s.IsDir = &v
	return s
}

func (s *GetCategoryResponseBody) SetName(v string) *GetCategoryResponseBody {
	s.Name = &v
	return s
}

func (s *GetCategoryResponseBody) SetParentCode(v string) *GetCategoryResponseBody {
	s.ParentCode = &v
	return s
}

func (s *GetCategoryResponseBody) SetRemark(v string) *GetCategoryResponseBody {
	s.Remark = &v
	return s
}

func (s *GetCategoryResponseBody) SetStatus(v string) *GetCategoryResponseBody {
	s.Status = &v
	return s
}

func (s *GetCategoryResponseBody) SetType(v string) *GetCategoryResponseBody {
	s.Type = &v
	return s
}

type GetCategoryResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCategoryResponse) GoString() string {
	return s.String()
}

func (s *GetCategoryResponse) SetHeaders(v map[string]*string) *GetCategoryResponse {
	s.Headers = v
	return s
}

func (s *GetCategoryResponse) SetStatusCode(v int32) *GetCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCategoryResponse) SetBody(v *GetCategoryResponseBody) *GetCategoryResponse {
	s.Body = v
	return s
}

type GetDefineHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDefineHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDefineHeaders) GoString() string {
	return s.String()
}

func (s *GetDefineHeaders) SetCommonHeaders(v map[string]*string) *GetDefineHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDefineHeaders) SetXAcsDingtalkAccessToken(v string) *GetDefineHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDefineRequest struct {
	// example:
	//
	// DEF_123456
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 门店
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetDefineRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDefineRequest) GoString() string {
	return s.String()
}

func (s *GetDefineRequest) SetCode(v string) *GetDefineRequest {
	s.Code = &v
	return s
}

func (s *GetDefineRequest) SetName(v string) *GetDefineRequest {
	s.Name = &v
	return s
}

func (s *GetDefineRequest) SetPageNumber(v int32) *GetDefineRequest {
	s.PageNumber = &v
	return s
}

func (s *GetDefineRequest) SetPageSize(v int32) *GetDefineRequest {
	s.PageSize = &v
	return s
}

type GetDefineResponseBody struct {
	HasMore    *bool                        `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*GetDefineResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	TotalCount *int64                       `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetDefineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDefineResponseBody) GoString() string {
	return s.String()
}

func (s *GetDefineResponseBody) SetHasMore(v bool) *GetDefineResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetDefineResponseBody) SetList(v []*GetDefineResponseBodyList) *GetDefineResponseBody {
	s.List = v
	return s
}

func (s *GetDefineResponseBody) SetTotalCount(v int64) *GetDefineResponseBody {
	s.TotalCount = &v
	return s
}

type GetDefineResponseBodyList struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetDefineResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s GetDefineResponseBodyList) GoString() string {
	return s.String()
}

func (s *GetDefineResponseBodyList) SetCode(v string) *GetDefineResponseBodyList {
	s.Code = &v
	return s
}

func (s *GetDefineResponseBodyList) SetName(v string) *GetDefineResponseBodyList {
	s.Name = &v
	return s
}

type GetDefineResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDefineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDefineResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDefineResponse) GoString() string {
	return s.String()
}

func (s *GetDefineResponse) SetHeaders(v map[string]*string) *GetDefineResponse {
	s.Headers = v
	return s
}

func (s *GetDefineResponse) SetStatusCode(v int32) *GetDefineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDefineResponse) SetBody(v *GetDefineResponseBody) *GetDefineResponse {
	s.Body = v
	return s
}

type GetDefineDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDefineDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDefineDataHeaders) GoString() string {
	return s.String()
}

func (s *GetDefineDataHeaders) SetCommonHeaders(v map[string]*string) *GetDefineDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDefineDataHeaders) SetXAcsDingtalkAccessToken(v string) *GetDefineDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDefineDataRequest struct {
	// example:
	//
	// DEF_123456
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// xx路门店1号
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetDefineDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDefineDataRequest) GoString() string {
	return s.String()
}

func (s *GetDefineDataRequest) SetCode(v string) *GetDefineDataRequest {
	s.Code = &v
	return s
}

func (s *GetDefineDataRequest) SetName(v string) *GetDefineDataRequest {
	s.Name = &v
	return s
}

func (s *GetDefineDataRequest) SetPageNumber(v int32) *GetDefineDataRequest {
	s.PageNumber = &v
	return s
}

func (s *GetDefineDataRequest) SetPageSize(v int32) *GetDefineDataRequest {
	s.PageSize = &v
	return s
}

type GetDefineDataResponseBody struct {
	HasMore    *bool                            `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*GetDefineDataResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	TotalCount *int64                           `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetDefineDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDefineDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetDefineDataResponseBody) SetHasMore(v bool) *GetDefineDataResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetDefineDataResponseBody) SetList(v []*GetDefineDataResponseBodyList) *GetDefineDataResponseBody {
	s.List = v
	return s
}

func (s *GetDefineDataResponseBody) SetTotalCount(v int64) *GetDefineDataResponseBody {
	s.TotalCount = &v
	return s
}

type GetDefineDataResponseBodyList struct {
	// example:
	//
	// DA_123456
	DataCode *string `json:"dataCode,omitempty" xml:"dataCode,omitempty"`
	// example:
	//
	// DEF_123456
	DefineCode *string `json:"defineCode,omitempty" xml:"defineCode,omitempty"`
	// example:
	//
	// xx路1号门店
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// -1
	ParentDataCode *string `json:"parentDataCode,omitempty" xml:"parentDataCode,omitempty"`
	// example:
	//
	// valid
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetDefineDataResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s GetDefineDataResponseBodyList) GoString() string {
	return s.String()
}

func (s *GetDefineDataResponseBodyList) SetDataCode(v string) *GetDefineDataResponseBodyList {
	s.DataCode = &v
	return s
}

func (s *GetDefineDataResponseBodyList) SetDefineCode(v string) *GetDefineDataResponseBodyList {
	s.DefineCode = &v
	return s
}

func (s *GetDefineDataResponseBodyList) SetName(v string) *GetDefineDataResponseBodyList {
	s.Name = &v
	return s
}

func (s *GetDefineDataResponseBodyList) SetParentDataCode(v string) *GetDefineDataResponseBodyList {
	s.ParentDataCode = &v
	return s
}

func (s *GetDefineDataResponseBodyList) SetStatus(v string) *GetDefineDataResponseBodyList {
	s.Status = &v
	return s
}

type GetDefineDataResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDefineDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDefineDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDefineDataResponse) GoString() string {
	return s.String()
}

func (s *GetDefineDataResponse) SetHeaders(v map[string]*string) *GetDefineDataResponse {
	s.Headers = v
	return s
}

func (s *GetDefineDataResponse) SetStatusCode(v int32) *GetDefineDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDefineDataResponse) SetBody(v *GetDefineDataResponseBody) *GetDefineDataResponse {
	s.Body = v
	return s
}

type GetFinanceAccountHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFinanceAccountHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFinanceAccountHeaders) GoString() string {
	return s.String()
}

func (s *GetFinanceAccountHeaders) SetCommonHeaders(v map[string]*string) *GetFinanceAccountHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFinanceAccountHeaders) SetXAcsDingtalkAccessToken(v string) *GetFinanceAccountHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFinanceAccountRequest struct {
	// This parameter is required.
	AccountCode *string `json:"accountCode,omitempty" xml:"accountCode,omitempty"`
}

func (s GetFinanceAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFinanceAccountRequest) GoString() string {
	return s.String()
}

func (s *GetFinanceAccountRequest) SetAccountCode(v string) *GetFinanceAccountRequest {
	s.AccountCode = &v
	return s
}

type GetFinanceAccountResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	AccountCode *string `json:"accountCode,omitempty" xml:"accountCode,omitempty"`
	// example:
	//
	// test@alipay.com
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试
	AccountName   *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	AccountRemark *string `json:"accountRemark,omitempty" xml:"accountRemark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALIPAY
	AccountType          *string   `json:"accountType,omitempty" xml:"accountType,omitempty"`
	AccountantBookIdList []*string `json:"accountantBookIdList,omitempty" xml:"accountantBookIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 50000.55
	Amount   *string `json:"amount,omitempty" xml:"amount,omitempty"`
	BankCode *string `json:"bankCode,omitempty" xml:"bankCode,omitempty"`
	BankName *string `json:"bankName,omitempty" xml:"bankName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1631526550994
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abcdef
	Creator        *string `json:"creator,omitempty" xml:"creator,omitempty"`
	OfficialName   *string `json:"officialName,omitempty" xml:"officialName,omitempty"`
	OfficialNumber *string `json:"officialNumber,omitempty" xml:"officialNumber,omitempty"`
	SignStatus     *string `json:"signStatus,omitempty" xml:"signStatus,omitempty"`
}

func (s GetFinanceAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFinanceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *GetFinanceAccountResponseBody) SetAccountCode(v string) *GetFinanceAccountResponseBody {
	s.AccountCode = &v
	return s
}

func (s *GetFinanceAccountResponseBody) SetAccountId(v string) *GetFinanceAccountResponseBody {
	s.AccountId = &v
	return s
}

func (s *GetFinanceAccountResponseBody) SetAccountName(v string) *GetFinanceAccountResponseBody {
	s.AccountName = &v
	return s
}

func (s *GetFinanceAccountResponseBody) SetAccountRemark(v string) *GetFinanceAccountResponseBody {
	s.AccountRemark = &v
	return s
}

func (s *GetFinanceAccountResponseBody) SetAccountType(v string) *GetFinanceAccountResponseBody {
	s.AccountType = &v
	return s
}

func (s *GetFinanceAccountResponseBody) SetAccountantBookIdList(v []*string) *GetFinanceAccountResponseBody {
	s.AccountantBookIdList = v
	return s
}

func (s *GetFinanceAccountResponseBody) SetAmount(v string) *GetFinanceAccountResponseBody {
	s.Amount = &v
	return s
}

func (s *GetFinanceAccountResponseBody) SetBankCode(v string) *GetFinanceAccountResponseBody {
	s.BankCode = &v
	return s
}

func (s *GetFinanceAccountResponseBody) SetBankName(v string) *GetFinanceAccountResponseBody {
	s.BankName = &v
	return s
}

func (s *GetFinanceAccountResponseBody) SetCreateTime(v int64) *GetFinanceAccountResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetFinanceAccountResponseBody) SetCreator(v string) *GetFinanceAccountResponseBody {
	s.Creator = &v
	return s
}

func (s *GetFinanceAccountResponseBody) SetOfficialName(v string) *GetFinanceAccountResponseBody {
	s.OfficialName = &v
	return s
}

func (s *GetFinanceAccountResponseBody) SetOfficialNumber(v string) *GetFinanceAccountResponseBody {
	s.OfficialNumber = &v
	return s
}

func (s *GetFinanceAccountResponseBody) SetSignStatus(v string) *GetFinanceAccountResponseBody {
	s.SignStatus = &v
	return s
}

type GetFinanceAccountResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFinanceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFinanceAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFinanceAccountResponse) GoString() string {
	return s.String()
}

func (s *GetFinanceAccountResponse) SetHeaders(v map[string]*string) *GetFinanceAccountResponse {
	s.Headers = v
	return s
}

func (s *GetFinanceAccountResponse) SetStatusCode(v int32) *GetFinanceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFinanceAccountResponse) SetBody(v *GetFinanceAccountResponseBody) *GetFinanceAccountResponse {
	s.Body = v
	return s
}

type GetProjectHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetProjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetProjectHeaders) GoString() string {
	return s.String()
}

func (s *GetProjectHeaders) SetCommonHeaders(v map[string]*string) *GetProjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetProjectHeaders) SetXAcsDingtalkAccessToken(v string) *GetProjectHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetProjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// PROJ-xxx
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s GetProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectRequest) GoString() string {
	return s.String()
}

func (s *GetProjectRequest) SetCode(v string) *GetProjectRequest {
	s.Code = &v
	return s
}

type GetProjectResponseBody struct {
	AccountantBookIdList []*string `json:"accountantBookIdList,omitempty" xml:"accountantBookIdList,omitempty" type:"Repeated"`
	Code                 *string   `json:"code,omitempty" xml:"code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1631526550994
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aaa
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 和外部合作
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	ParentCode  *string `json:"parentCode,omitempty" xml:"parentCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROJ-XXX
	ProjectCode *string `json:"projectCode,omitempty" xml:"projectCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 外包项目
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// valid
	Status         *string `json:"status,omitempty" xml:"status,omitempty"`
	UserDefineCode *string `json:"userDefineCode,omitempty" xml:"userDefineCode,omitempty"`
}

func (s GetProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBody) SetAccountantBookIdList(v []*string) *GetProjectResponseBody {
	s.AccountantBookIdList = v
	return s
}

func (s *GetProjectResponseBody) SetCode(v string) *GetProjectResponseBody {
	s.Code = &v
	return s
}

func (s *GetProjectResponseBody) SetCreateTime(v int64) *GetProjectResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetProjectResponseBody) SetCreator(v string) *GetProjectResponseBody {
	s.Creator = &v
	return s
}

func (s *GetProjectResponseBody) SetDescription(v string) *GetProjectResponseBody {
	s.Description = &v
	return s
}

func (s *GetProjectResponseBody) SetName(v string) *GetProjectResponseBody {
	s.Name = &v
	return s
}

func (s *GetProjectResponseBody) SetParentCode(v string) *GetProjectResponseBody {
	s.ParentCode = &v
	return s
}

func (s *GetProjectResponseBody) SetProjectCode(v string) *GetProjectResponseBody {
	s.ProjectCode = &v
	return s
}

func (s *GetProjectResponseBody) SetProjectName(v string) *GetProjectResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetProjectResponseBody) SetStatus(v string) *GetProjectResponseBody {
	s.Status = &v
	return s
}

func (s *GetProjectResponseBody) SetUserDefineCode(v string) *GetProjectResponseBody {
	s.UserDefineCode = &v
	return s
}

type GetProjectResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponse) GoString() string {
	return s.String()
}

func (s *GetProjectResponse) SetHeaders(v map[string]*string) *GetProjectResponse {
	s.Headers = v
	return s
}

func (s *GetProjectResponse) SetStatusCode(v int32) *GetProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectResponse) SetBody(v *GetProjectResponseBody) *GetProjectResponse {
	s.Body = v
	return s
}

type GetReceiptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetReceiptHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetReceiptHeaders) GoString() string {
	return s.String()
}

func (s *GetReceiptHeaders) SetCommonHeaders(v map[string]*string) *GetReceiptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetReceiptHeaders) SetXAcsDingtalkAccessToken(v string) *GetReceiptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetReceiptRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 19b98a1c-5a31-4d78-9da7-0e347593820a
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// EM-1017F28E03350B1738B3000X
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
}

func (s GetReceiptRequest) String() string {
	return tea.Prettify(s)
}

func (s GetReceiptRequest) GoString() string {
	return s.String()
}

func (s *GetReceiptRequest) SetCode(v string) *GetReceiptRequest {
	s.Code = &v
	return s
}

func (s *GetReceiptRequest) SetModelId(v string) *GetReceiptRequest {
	s.ModelId = &v
	return s
}

type GetReceiptResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"operatorUserId":"015865244722391178","data":{"amount":{"amountStr":"566"},"code":"d0d54815-32c5-4b18-8391-e79713bba95e","payeeAt":1637251200000,"departmentCode":"-1","project":{"projectCode":"PROJ_101761F3FF6B21362ECA000N","projectName":"客户合作项目"},"principalId":"015865244722391178","enterpriseAccount":{},"approvedAt":1637305373000,"title":"地狱猫提交的智能财务-收款","createAt":1637305353000,"paymentAt":1637251200000,"supplier":{},"operateUserId":"015865244722391178","applicantEmployeeId":"015865244722391178","comment":"ffff","category":{"categoryCode":"INC_1016D6CB3C181E28F0120009","categoryName":"销售收入"},"customer":{"customerCode":"CUS_10178592ECEC2133C893000F","customerName":"钉钉"},"status":"agree"}}
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// EM-1017F28E03350B1738B3000X
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// approval
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s GetReceiptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetReceiptResponseBody) GoString() string {
	return s.String()
}

func (s *GetReceiptResponseBody) SetAppId(v string) *GetReceiptResponseBody {
	s.AppId = &v
	return s
}

func (s *GetReceiptResponseBody) SetData(v string) *GetReceiptResponseBody {
	s.Data = &v
	return s
}

func (s *GetReceiptResponseBody) SetModelId(v string) *GetReceiptResponseBody {
	s.ModelId = &v
	return s
}

func (s *GetReceiptResponseBody) SetSource(v string) *GetReceiptResponseBody {
	s.Source = &v
	return s
}

type GetReceiptResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetReceiptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetReceiptResponse) String() string {
	return tea.Prettify(s)
}

func (s GetReceiptResponse) GoString() string {
	return s.String()
}

func (s *GetReceiptResponse) SetHeaders(v map[string]*string) *GetReceiptResponse {
	s.Headers = v
	return s
}

func (s *GetReceiptResponse) SetStatusCode(v int32) *GetReceiptResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReceiptResponse) SetBody(v *GetReceiptResponseBody) *GetReceiptResponse {
	s.Body = v
	return s
}

type GetSupplierHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSupplierHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSupplierHeaders) GoString() string {
	return s.String()
}

func (s *GetSupplierHeaders) SetCommonHeaders(v map[string]*string) *GetSupplierHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSupplierHeaders) SetXAcsDingtalkAccessToken(v string) *GetSupplierHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSupplierRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// SUP_XXX
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s GetSupplierRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSupplierRequest) GoString() string {
	return s.String()
}

func (s *GetSupplierRequest) SetCode(v string) *GetSupplierRequest {
	s.Code = &v
	return s
}

type GetSupplierResponseBody struct {
	AccountantBookIdList []*string `json:"accountantBookIdList,omitempty" xml:"accountantBookIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// SUP_XXX
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1634786828686
	CreateTime         *int64                                       `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CustomFormDataList []*GetSupplierResponseBodyCustomFormDataList `json:"customFormDataList,omitempty" xml:"customFormDataList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 原材料供应商
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// XX公司
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// valid
	Status         *string `json:"status,omitempty" xml:"status,omitempty"`
	UserDefineCode *string `json:"userDefineCode,omitempty" xml:"userDefineCode,omitempty"`
}

func (s GetSupplierResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSupplierResponseBody) GoString() string {
	return s.String()
}

func (s *GetSupplierResponseBody) SetAccountantBookIdList(v []*string) *GetSupplierResponseBody {
	s.AccountantBookIdList = v
	return s
}

func (s *GetSupplierResponseBody) SetCode(v string) *GetSupplierResponseBody {
	s.Code = &v
	return s
}

func (s *GetSupplierResponseBody) SetCreateTime(v int64) *GetSupplierResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetSupplierResponseBody) SetCustomFormDataList(v []*GetSupplierResponseBodyCustomFormDataList) *GetSupplierResponseBody {
	s.CustomFormDataList = v
	return s
}

func (s *GetSupplierResponseBody) SetDescription(v string) *GetSupplierResponseBody {
	s.Description = &v
	return s
}

func (s *GetSupplierResponseBody) SetName(v string) *GetSupplierResponseBody {
	s.Name = &v
	return s
}

func (s *GetSupplierResponseBody) SetStatus(v string) *GetSupplierResponseBody {
	s.Status = &v
	return s
}

func (s *GetSupplierResponseBody) SetUserDefineCode(v string) *GetSupplierResponseBody {
	s.UserDefineCode = &v
	return s
}

type GetSupplierResponseBodyCustomFormDataList struct {
	BizAlias      *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	ExtValue      *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	Id            *string `json:"id,omitempty" xml:"id,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	Value         *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetSupplierResponseBodyCustomFormDataList) String() string {
	return tea.Prettify(s)
}

func (s GetSupplierResponseBodyCustomFormDataList) GoString() string {
	return s.String()
}

func (s *GetSupplierResponseBodyCustomFormDataList) SetBizAlias(v string) *GetSupplierResponseBodyCustomFormDataList {
	s.BizAlias = &v
	return s
}

func (s *GetSupplierResponseBodyCustomFormDataList) SetComponentType(v string) *GetSupplierResponseBodyCustomFormDataList {
	s.ComponentType = &v
	return s
}

func (s *GetSupplierResponseBodyCustomFormDataList) SetExtValue(v string) *GetSupplierResponseBodyCustomFormDataList {
	s.ExtValue = &v
	return s
}

func (s *GetSupplierResponseBodyCustomFormDataList) SetId(v string) *GetSupplierResponseBodyCustomFormDataList {
	s.Id = &v
	return s
}

func (s *GetSupplierResponseBodyCustomFormDataList) SetName(v string) *GetSupplierResponseBodyCustomFormDataList {
	s.Name = &v
	return s
}

func (s *GetSupplierResponseBodyCustomFormDataList) SetValue(v string) *GetSupplierResponseBodyCustomFormDataList {
	s.Value = &v
	return s
}

type GetSupplierResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSupplierResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSupplierResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSupplierResponse) GoString() string {
	return s.String()
}

func (s *GetSupplierResponse) SetHeaders(v map[string]*string) *GetSupplierResponse {
	s.Headers = v
	return s
}

func (s *GetSupplierResponse) SetStatusCode(v int32) *GetSupplierResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSupplierResponse) SetBody(v *GetSupplierResponseBody) *GetSupplierResponse {
	s.Body = v
	return s
}

type IssueInvoiceWithOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IssueInvoiceWithOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s IssueInvoiceWithOrderHeaders) GoString() string {
	return s.String()
}

func (s *IssueInvoiceWithOrderHeaders) SetCommonHeaders(v map[string]*string) *IssueInvoiceWithOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IssueInvoiceWithOrderHeaders) SetXAcsDingtalkAccessToken(v string) *IssueInvoiceWithOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IssueInvoiceWithOrderRequest struct {
	Content       *IssueInvoiceWithOrderRequestContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	FinanceAppKey *string                              `json:"financeAppKey,omitempty" xml:"financeAppKey,omitempty"`
	Operator      *string                              `json:"operator,omitempty" xml:"operator,omitempty"`
	Signature     *string                              `json:"signature,omitempty" xml:"signature,omitempty"`
}

func (s IssueInvoiceWithOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s IssueInvoiceWithOrderRequest) GoString() string {
	return s.String()
}

func (s *IssueInvoiceWithOrderRequest) SetContent(v *IssueInvoiceWithOrderRequestContent) *IssueInvoiceWithOrderRequest {
	s.Content = v
	return s
}

func (s *IssueInvoiceWithOrderRequest) SetFinanceAppKey(v string) *IssueInvoiceWithOrderRequest {
	s.FinanceAppKey = &v
	return s
}

func (s *IssueInvoiceWithOrderRequest) SetOperator(v string) *IssueInvoiceWithOrderRequest {
	s.Operator = &v
	return s
}

func (s *IssueInvoiceWithOrderRequest) SetSignature(v string) *IssueInvoiceWithOrderRequest {
	s.Signature = &v
	return s
}

type IssueInvoiceWithOrderRequestContent struct {
	AdditionInfo     []*IssueInvoiceWithOrderRequestContentAdditionInfo `json:"additionInfo,omitempty" xml:"additionInfo,omitempty" type:"Repeated"`
	ApplyPerson      *string                                            `json:"applyPerson,omitempty" xml:"applyPerson,omitempty"`
	BankAccount      *string                                            `json:"bankAccount,omitempty" xml:"bankAccount,omitempty"`
	BankName         *string                                            `json:"bankName,omitempty" xml:"bankName,omitempty"`
	InvoiceRemark    *string                                            `json:"invoiceRemark,omitempty" xml:"invoiceRemark,omitempty"`
	InvoiceType      *int64                                             `json:"invoiceType,omitempty" xml:"invoiceType,omitempty"`
	NaturalPerson    *string                                            `json:"naturalPerson,omitempty" xml:"naturalPerson,omitempty"`
	OrderId          *string                                            `json:"orderId,omitempty" xml:"orderId,omitempty"`
	Payee            *string                                            `json:"payee,omitempty" xml:"payee,omitempty"`
	Phone            *string                                            `json:"phone,omitempty" xml:"phone,omitempty"`
	Products         []*IssueInvoiceWithOrderRequestContentProducts     `json:"products,omitempty" xml:"products,omitempty" type:"Repeated"`
	Purchaser        *string                                            `json:"purchaser,omitempty" xml:"purchaser,omitempty"`
	PurchaserAddress *string                                            `json:"purchaserAddress,omitempty" xml:"purchaserAddress,omitempty"`
	PurchaserTel     *string                                            `json:"purchaserTel,omitempty" xml:"purchaserTel,omitempty"`
	Remark           *string                                            `json:"remark,omitempty" xml:"remark,omitempty"`
	Reviewer         *string                                            `json:"reviewer,omitempty" xml:"reviewer,omitempty"`
	Taxnum           *string                                            `json:"taxnum,omitempty" xml:"taxnum,omitempty"`
}

func (s IssueInvoiceWithOrderRequestContent) String() string {
	return tea.Prettify(s)
}

func (s IssueInvoiceWithOrderRequestContent) GoString() string {
	return s.String()
}

func (s *IssueInvoiceWithOrderRequestContent) SetAdditionInfo(v []*IssueInvoiceWithOrderRequestContentAdditionInfo) *IssueInvoiceWithOrderRequestContent {
	s.AdditionInfo = v
	return s
}

func (s *IssueInvoiceWithOrderRequestContent) SetApplyPerson(v string) *IssueInvoiceWithOrderRequestContent {
	s.ApplyPerson = &v
	return s
}

func (s *IssueInvoiceWithOrderRequestContent) SetBankAccount(v string) *IssueInvoiceWithOrderRequestContent {
	s.BankAccount = &v
	return s
}

func (s *IssueInvoiceWithOrderRequestContent) SetBankName(v string) *IssueInvoiceWithOrderRequestContent {
	s.BankName = &v
	return s
}

func (s *IssueInvoiceWithOrderRequestContent) SetInvoiceRemark(v string) *IssueInvoiceWithOrderRequestContent {
	s.InvoiceRemark = &v
	return s
}

func (s *IssueInvoiceWithOrderRequestContent) SetInvoiceType(v int64) *IssueInvoiceWithOrderRequestContent {
	s.InvoiceType = &v
	return s
}

func (s *IssueInvoiceWithOrderRequestContent) SetNaturalPerson(v string) *IssueInvoiceWithOrderRequestContent {
	s.NaturalPerson = &v
	return s
}

func (s *IssueInvoiceWithOrderRequestContent) SetOrderId(v string) *IssueInvoiceWithOrderRequestContent {
	s.OrderId = &v
	return s
}

func (s *IssueInvoiceWithOrderRequestContent) SetPayee(v string) *IssueInvoiceWithOrderRequestContent {
	s.Payee = &v
	return s
}

func (s *IssueInvoiceWithOrderRequestContent) SetPhone(v string) *IssueInvoiceWithOrderRequestContent {
	s.Phone = &v
	return s
}

func (s *IssueInvoiceWithOrderRequestContent) SetProducts(v []*IssueInvoiceWithOrderRequestContentProducts) *IssueInvoiceWithOrderRequestContent {
	s.Products = v
	return s
}

func (s *IssueInvoiceWithOrderRequestContent) SetPurchaser(v string) *IssueInvoiceWithOrderRequestContent {
	s.Purchaser = &v
	return s
}

func (s *IssueInvoiceWithOrderRequestContent) SetPurchaserAddress(v string) *IssueInvoiceWithOrderRequestContent {
	s.PurchaserAddress = &v
	return s
}

func (s *IssueInvoiceWithOrderRequestContent) SetPurchaserTel(v string) *IssueInvoiceWithOrderRequestContent {
	s.PurchaserTel = &v
	return s
}

func (s *IssueInvoiceWithOrderRequestContent) SetRemark(v string) *IssueInvoiceWithOrderRequestContent {
	s.Remark = &v
	return s
}

func (s *IssueInvoiceWithOrderRequestContent) SetReviewer(v string) *IssueInvoiceWithOrderRequestContent {
	s.Reviewer = &v
	return s
}

func (s *IssueInvoiceWithOrderRequestContent) SetTaxnum(v string) *IssueInvoiceWithOrderRequestContent {
	s.Taxnum = &v
	return s
}

type IssueInvoiceWithOrderRequestContentAdditionInfo struct {
	AdditionContent *string `json:"additionContent,omitempty" xml:"additionContent,omitempty"`
	AdditionName    *string `json:"additionName,omitempty" xml:"additionName,omitempty"`
	DataType        *int64  `json:"dataType,omitempty" xml:"dataType,omitempty"`
}

func (s IssueInvoiceWithOrderRequestContentAdditionInfo) String() string {
	return tea.Prettify(s)
}

func (s IssueInvoiceWithOrderRequestContentAdditionInfo) GoString() string {
	return s.String()
}

func (s *IssueInvoiceWithOrderRequestContentAdditionInfo) SetAdditionContent(v string) *IssueInvoiceWithOrderRequestContentAdditionInfo {
	s.AdditionContent = &v
	return s
}

func (s *IssueInvoiceWithOrderRequestContentAdditionInfo) SetAdditionName(v string) *IssueInvoiceWithOrderRequestContentAdditionInfo {
	s.AdditionName = &v
	return s
}

func (s *IssueInvoiceWithOrderRequestContentAdditionInfo) SetDataType(v int64) *IssueInvoiceWithOrderRequestContentAdditionInfo {
	s.DataType = &v
	return s
}

type IssueInvoiceWithOrderRequestContentProducts struct {
	AmountIncludeTax *string `json:"amountIncludeTax,omitempty" xml:"amountIncludeTax,omitempty"`
	ProductName      *string `json:"productName,omitempty" xml:"productName,omitempty"`
	Quantity         *string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	RevenueCode      *string `json:"revenueCode,omitempty" xml:"revenueCode,omitempty"`
	Specs            *string `json:"specs,omitempty" xml:"specs,omitempty"`
	TaxSign          *string `json:"taxSign,omitempty" xml:"taxSign,omitempty"`
	Unit             *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s IssueInvoiceWithOrderRequestContentProducts) String() string {
	return tea.Prettify(s)
}

func (s IssueInvoiceWithOrderRequestContentProducts) GoString() string {
	return s.String()
}

func (s *IssueInvoiceWithOrderRequestContentProducts) SetAmountIncludeTax(v string) *IssueInvoiceWithOrderRequestContentProducts {
	s.AmountIncludeTax = &v
	return s
}

func (s *IssueInvoiceWithOrderRequestContentProducts) SetProductName(v string) *IssueInvoiceWithOrderRequestContentProducts {
	s.ProductName = &v
	return s
}

func (s *IssueInvoiceWithOrderRequestContentProducts) SetQuantity(v string) *IssueInvoiceWithOrderRequestContentProducts {
	s.Quantity = &v
	return s
}

func (s *IssueInvoiceWithOrderRequestContentProducts) SetRevenueCode(v string) *IssueInvoiceWithOrderRequestContentProducts {
	s.RevenueCode = &v
	return s
}

func (s *IssueInvoiceWithOrderRequestContentProducts) SetSpecs(v string) *IssueInvoiceWithOrderRequestContentProducts {
	s.Specs = &v
	return s
}

func (s *IssueInvoiceWithOrderRequestContentProducts) SetTaxSign(v string) *IssueInvoiceWithOrderRequestContentProducts {
	s.TaxSign = &v
	return s
}

func (s *IssueInvoiceWithOrderRequestContentProducts) SetUnit(v string) *IssueInvoiceWithOrderRequestContentProducts {
	s.Unit = &v
	return s
}

type IssueInvoiceWithOrderResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s IssueInvoiceWithOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IssueInvoiceWithOrderResponseBody) GoString() string {
	return s.String()
}

func (s *IssueInvoiceWithOrderResponseBody) SetSuccess(v bool) *IssueInvoiceWithOrderResponseBody {
	s.Success = &v
	return s
}

type IssueInvoiceWithOrderResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IssueInvoiceWithOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IssueInvoiceWithOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s IssueInvoiceWithOrderResponse) GoString() string {
	return s.String()
}

func (s *IssueInvoiceWithOrderResponse) SetHeaders(v map[string]*string) *IssueInvoiceWithOrderResponse {
	s.Headers = v
	return s
}

func (s *IssueInvoiceWithOrderResponse) SetStatusCode(v int32) *IssueInvoiceWithOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *IssueInvoiceWithOrderResponse) SetBody(v *IssueInvoiceWithOrderResponseBody) *IssueInvoiceWithOrderResponse {
	s.Body = v
	return s
}

type LinkCommonInvokeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s LinkCommonInvokeHeaders) String() string {
	return tea.Prettify(s)
}

func (s LinkCommonInvokeHeaders) GoString() string {
	return s.String()
}

func (s *LinkCommonInvokeHeaders) SetCommonHeaders(v map[string]*string) *LinkCommonInvokeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *LinkCommonInvokeHeaders) SetXAcsDingtalkAccessToken(v string) *LinkCommonInvokeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type LinkCommonInvokeRequest struct {
	// example:
	//
	// TEST
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// example:
	//
	// "{\"key\":\"value\"}"
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// INOVKE_XXX
	InvokeId *string `json:"invokeId,omitempty" xml:"invokeId,omitempty"`
	// example:
	//
	// abc
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s LinkCommonInvokeRequest) String() string {
	return tea.Prettify(s)
}

func (s LinkCommonInvokeRequest) GoString() string {
	return s.String()
}

func (s *LinkCommonInvokeRequest) SetBizType(v string) *LinkCommonInvokeRequest {
	s.BizType = &v
	return s
}

func (s *LinkCommonInvokeRequest) SetData(v string) *LinkCommonInvokeRequest {
	s.Data = &v
	return s
}

func (s *LinkCommonInvokeRequest) SetInvokeId(v string) *LinkCommonInvokeRequest {
	s.InvokeId = &v
	return s
}

func (s *LinkCommonInvokeRequest) SetUserId(v string) *LinkCommonInvokeRequest {
	s.UserId = &v
	return s
}

type LinkCommonInvokeResponseBody struct {
	BizType  *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	Data     *string `json:"data,omitempty" xml:"data,omitempty"`
	InvokeId *string `json:"invokeId,omitempty" xml:"invokeId,omitempty"`
}

func (s LinkCommonInvokeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LinkCommonInvokeResponseBody) GoString() string {
	return s.String()
}

func (s *LinkCommonInvokeResponseBody) SetBizType(v string) *LinkCommonInvokeResponseBody {
	s.BizType = &v
	return s
}

func (s *LinkCommonInvokeResponseBody) SetData(v string) *LinkCommonInvokeResponseBody {
	s.Data = &v
	return s
}

func (s *LinkCommonInvokeResponseBody) SetInvokeId(v string) *LinkCommonInvokeResponseBody {
	s.InvokeId = &v
	return s
}

type LinkCommonInvokeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LinkCommonInvokeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LinkCommonInvokeResponse) String() string {
	return tea.Prettify(s)
}

func (s LinkCommonInvokeResponse) GoString() string {
	return s.String()
}

func (s *LinkCommonInvokeResponse) SetHeaders(v map[string]*string) *LinkCommonInvokeResponse {
	s.Headers = v
	return s
}

func (s *LinkCommonInvokeResponse) SetStatusCode(v int32) *LinkCommonInvokeResponse {
	s.StatusCode = &v
	return s
}

func (s *LinkCommonInvokeResponse) SetBody(v *LinkCommonInvokeResponseBody) *LinkCommonInvokeResponse {
	s.Body = v
	return s
}

type OrderBillingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s OrderBillingHeaders) String() string {
	return tea.Prettify(s)
}

func (s OrderBillingHeaders) GoString() string {
	return s.String()
}

func (s *OrderBillingHeaders) SetCommonHeaders(v map[string]*string) *OrderBillingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OrderBillingHeaders) SetXAcsDingtalkAccessToken(v string) *OrderBillingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type OrderBillingRequest struct {
	AdditionInfos []*OrderBillingRequestAdditionInfos `json:"additionInfos,omitempty" xml:"additionInfos,omitempty" type:"Repeated"`
	// example:
	//
	// abc
	AppKey        *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	ApplyPerson   *string `json:"applyPerson,omitempty" xml:"applyPerson,omitempty"`
	InvoiceRemark *string `json:"invoiceRemark,omitempty" xml:"invoiceRemark,omitempty"`
	// example:
	//
	// VAT_NORMAL_E
	InvoiceType     *string `json:"invoiceType,omitempty" xml:"invoiceType,omitempty"`
	IsNaturalPerson *bool   `json:"isNaturalPerson,omitempty" xml:"isNaturalPerson,omitempty"`
	Operator        *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// example:
	//
	// abc
	OrderId  *string                        `json:"orderId,omitempty" xml:"orderId,omitempty"`
	Payee    *string                        `json:"payee,omitempty" xml:"payee,omitempty"`
	Phone    *string                        `json:"phone,omitempty" xml:"phone,omitempty"`
	Products []*OrderBillingRequestProducts `json:"products,omitempty" xml:"products,omitempty" type:"Repeated"`
	// example:
	//
	// 浙江省杭州市XXX街道
	PurchaserAddress     *string `json:"purchaserAddress,omitempty" xml:"purchaserAddress,omitempty"`
	PurchaserBankAccount *string `json:"purchaserBankAccount,omitempty" xml:"purchaserBankAccount,omitempty"`
	PurchaserBankName    *string `json:"purchaserBankName,omitempty" xml:"purchaserBankName,omitempty"`
	// example:
	//
	// XXX公司
	PurchaserName  *string `json:"purchaserName,omitempty" xml:"purchaserName,omitempty"`
	PurchaserTaxNo *string `json:"purchaserTaxNo,omitempty" xml:"purchaserTaxNo,omitempty"`
	PurchaserTel   *string `json:"purchaserTel,omitempty" xml:"purchaserTel,omitempty"`
	Remark         *string `json:"remark,omitempty" xml:"remark,omitempty"`
	Reviewer       *string `json:"reviewer,omitempty" xml:"reviewer,omitempty"`
	Signature      *string `json:"signature,omitempty" xml:"signature,omitempty"`
	TaxSign        *int32  `json:"taxSign,omitempty" xml:"taxSign,omitempty"`
}

func (s OrderBillingRequest) String() string {
	return tea.Prettify(s)
}

func (s OrderBillingRequest) GoString() string {
	return s.String()
}

func (s *OrderBillingRequest) SetAdditionInfos(v []*OrderBillingRequestAdditionInfos) *OrderBillingRequest {
	s.AdditionInfos = v
	return s
}

func (s *OrderBillingRequest) SetAppKey(v string) *OrderBillingRequest {
	s.AppKey = &v
	return s
}

func (s *OrderBillingRequest) SetApplyPerson(v string) *OrderBillingRequest {
	s.ApplyPerson = &v
	return s
}

func (s *OrderBillingRequest) SetInvoiceRemark(v string) *OrderBillingRequest {
	s.InvoiceRemark = &v
	return s
}

func (s *OrderBillingRequest) SetInvoiceType(v string) *OrderBillingRequest {
	s.InvoiceType = &v
	return s
}

func (s *OrderBillingRequest) SetIsNaturalPerson(v bool) *OrderBillingRequest {
	s.IsNaturalPerson = &v
	return s
}

func (s *OrderBillingRequest) SetOperator(v string) *OrderBillingRequest {
	s.Operator = &v
	return s
}

func (s *OrderBillingRequest) SetOrderId(v string) *OrderBillingRequest {
	s.OrderId = &v
	return s
}

func (s *OrderBillingRequest) SetPayee(v string) *OrderBillingRequest {
	s.Payee = &v
	return s
}

func (s *OrderBillingRequest) SetPhone(v string) *OrderBillingRequest {
	s.Phone = &v
	return s
}

func (s *OrderBillingRequest) SetProducts(v []*OrderBillingRequestProducts) *OrderBillingRequest {
	s.Products = v
	return s
}

func (s *OrderBillingRequest) SetPurchaserAddress(v string) *OrderBillingRequest {
	s.PurchaserAddress = &v
	return s
}

func (s *OrderBillingRequest) SetPurchaserBankAccount(v string) *OrderBillingRequest {
	s.PurchaserBankAccount = &v
	return s
}

func (s *OrderBillingRequest) SetPurchaserBankName(v string) *OrderBillingRequest {
	s.PurchaserBankName = &v
	return s
}

func (s *OrderBillingRequest) SetPurchaserName(v string) *OrderBillingRequest {
	s.PurchaserName = &v
	return s
}

func (s *OrderBillingRequest) SetPurchaserTaxNo(v string) *OrderBillingRequest {
	s.PurchaserTaxNo = &v
	return s
}

func (s *OrderBillingRequest) SetPurchaserTel(v string) *OrderBillingRequest {
	s.PurchaserTel = &v
	return s
}

func (s *OrderBillingRequest) SetRemark(v string) *OrderBillingRequest {
	s.Remark = &v
	return s
}

func (s *OrderBillingRequest) SetReviewer(v string) *OrderBillingRequest {
	s.Reviewer = &v
	return s
}

func (s *OrderBillingRequest) SetSignature(v string) *OrderBillingRequest {
	s.Signature = &v
	return s
}

func (s *OrderBillingRequest) SetTaxSign(v int32) *OrderBillingRequest {
	s.TaxSign = &v
	return s
}

type OrderBillingRequestAdditionInfos struct {
	AdditionContent *string `json:"additionContent,omitempty" xml:"additionContent,omitempty"`
	AdditionName    *string `json:"additionName,omitempty" xml:"additionName,omitempty"`
	DataType        *int32  `json:"dataType,omitempty" xml:"dataType,omitempty"`
}

func (s OrderBillingRequestAdditionInfos) String() string {
	return tea.Prettify(s)
}

func (s OrderBillingRequestAdditionInfos) GoString() string {
	return s.String()
}

func (s *OrderBillingRequestAdditionInfos) SetAdditionContent(v string) *OrderBillingRequestAdditionInfos {
	s.AdditionContent = &v
	return s
}

func (s *OrderBillingRequestAdditionInfos) SetAdditionName(v string) *OrderBillingRequestAdditionInfos {
	s.AdditionName = &v
	return s
}

func (s *OrderBillingRequestAdditionInfos) SetDataType(v int32) *OrderBillingRequestAdditionInfos {
	s.DataType = &v
	return s
}

type OrderBillingRequestProducts struct {
	// example:
	//
	// 12.55
	AmountWithTax *string `json:"amountWithTax,omitempty" xml:"amountWithTax,omitempty"`
	// example:
	//
	// 面包
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// example:
	//
	// 5
	Quantity      *string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	RevenueCode   *string `json:"revenueCode,omitempty" xml:"revenueCode,omitempty"`
	Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
	// example:
	//
	// 个
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// example:
	//
	// 19.99
	UnitPrice *string `json:"unitPrice,omitempty" xml:"unitPrice,omitempty"`
}

func (s OrderBillingRequestProducts) String() string {
	return tea.Prettify(s)
}

func (s OrderBillingRequestProducts) GoString() string {
	return s.String()
}

func (s *OrderBillingRequestProducts) SetAmountWithTax(v string) *OrderBillingRequestProducts {
	s.AmountWithTax = &v
	return s
}

func (s *OrderBillingRequestProducts) SetProductName(v string) *OrderBillingRequestProducts {
	s.ProductName = &v
	return s
}

func (s *OrderBillingRequestProducts) SetQuantity(v string) *OrderBillingRequestProducts {
	s.Quantity = &v
	return s
}

func (s *OrderBillingRequestProducts) SetRevenueCode(v string) *OrderBillingRequestProducts {
	s.RevenueCode = &v
	return s
}

func (s *OrderBillingRequestProducts) SetSpecification(v string) *OrderBillingRequestProducts {
	s.Specification = &v
	return s
}

func (s *OrderBillingRequestProducts) SetUnit(v string) *OrderBillingRequestProducts {
	s.Unit = &v
	return s
}

func (s *OrderBillingRequestProducts) SetUnitPrice(v string) *OrderBillingRequestProducts {
	s.UnitPrice = &v
	return s
}

type OrderBillingResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OrderBillingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OrderBillingResponseBody) GoString() string {
	return s.String()
}

func (s *OrderBillingResponseBody) SetSuccess(v bool) *OrderBillingResponseBody {
	s.Success = &v
	return s
}

type OrderBillingResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OrderBillingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OrderBillingResponse) String() string {
	return tea.Prettify(s)
}

func (s OrderBillingResponse) GoString() string {
	return s.String()
}

func (s *OrderBillingResponse) SetHeaders(v map[string]*string) *OrderBillingResponse {
	s.Headers = v
	return s
}

func (s *OrderBillingResponse) SetStatusCode(v int32) *OrderBillingResponse {
	s.StatusCode = &v
	return s
}

func (s *OrderBillingResponse) SetBody(v *OrderBillingResponseBody) *OrderBillingResponse {
	s.Body = v
	return s
}

type QueryAccountTradeByPageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAccountTradeByPageHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountTradeByPageHeaders) GoString() string {
	return s.String()
}

func (s *QueryAccountTradeByPageHeaders) SetCommonHeaders(v map[string]*string) *QueryAccountTradeByPageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAccountTradeByPageHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAccountTradeByPageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAccountTradeByPageRequest struct {
	// example:
	//
	// ACC_XXXXX
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// 1730778990150
	EndDate *int64                                `json:"endDate,omitempty" xml:"endDate,omitempty"`
	Filter  *QueryAccountTradeByPageRequestFilter `json:"filter,omitempty" xml:"filter,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1730778990150
	StartDate *int64 `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// example:
	//
	// 50423414443123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryAccountTradeByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountTradeByPageRequest) GoString() string {
	return s.String()
}

func (s *QueryAccountTradeByPageRequest) SetAccountId(v string) *QueryAccountTradeByPageRequest {
	s.AccountId = &v
	return s
}

func (s *QueryAccountTradeByPageRequest) SetEndDate(v int64) *QueryAccountTradeByPageRequest {
	s.EndDate = &v
	return s
}

func (s *QueryAccountTradeByPageRequest) SetFilter(v *QueryAccountTradeByPageRequestFilter) *QueryAccountTradeByPageRequest {
	s.Filter = v
	return s
}

func (s *QueryAccountTradeByPageRequest) SetPageNumber(v int32) *QueryAccountTradeByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryAccountTradeByPageRequest) SetPageSize(v int32) *QueryAccountTradeByPageRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAccountTradeByPageRequest) SetStartDate(v int64) *QueryAccountTradeByPageRequest {
	s.StartDate = &v
	return s
}

func (s *QueryAccountTradeByPageRequest) SetUserId(v string) *QueryAccountTradeByPageRequest {
	s.UserId = &v
	return s
}

type QueryAccountTradeByPageRequestFilter struct {
	// example:
	//
	// 10
	EndAmount *string `json:"endAmount,omitempty" xml:"endAmount,omitempty"`
	// example:
	//
	// 钉钉
	OtherAccountName *string `json:"otherAccountName,omitempty" xml:"otherAccountName,omitempty"`
	// example:
	//
	// 1
	StartAmount *string `json:"startAmount,omitempty" xml:"startAmount,omitempty"`
	// example:
	//
	// in
	TradeType *string `json:"tradeType,omitempty" xml:"tradeType,omitempty"`
}

func (s QueryAccountTradeByPageRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountTradeByPageRequestFilter) GoString() string {
	return s.String()
}

func (s *QueryAccountTradeByPageRequestFilter) SetEndAmount(v string) *QueryAccountTradeByPageRequestFilter {
	s.EndAmount = &v
	return s
}

func (s *QueryAccountTradeByPageRequestFilter) SetOtherAccountName(v string) *QueryAccountTradeByPageRequestFilter {
	s.OtherAccountName = &v
	return s
}

func (s *QueryAccountTradeByPageRequestFilter) SetStartAmount(v string) *QueryAccountTradeByPageRequestFilter {
	s.StartAmount = &v
	return s
}

func (s *QueryAccountTradeByPageRequestFilter) SetTradeType(v string) *QueryAccountTradeByPageRequestFilter {
	s.TradeType = &v
	return s
}

type QueryAccountTradeByPageResponseBody struct {
	Result     []*QueryAccountTradeByPageResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int64                                       `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryAccountTradeByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountTradeByPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAccountTradeByPageResponseBody) SetResult(v []*QueryAccountTradeByPageResponseBodyResult) *QueryAccountTradeByPageResponseBody {
	s.Result = v
	return s
}

func (s *QueryAccountTradeByPageResponseBody) SetTotalCount(v int64) *QueryAccountTradeByPageResponseBody {
	s.TotalCount = &v
	return s
}

type QueryAccountTradeByPageResponseBodyResult struct {
	Balance          *string                                               `json:"balance,omitempty" xml:"balance,omitempty"`
	DetailId         *string                                               `json:"detailId,omitempty" xml:"detailId,omitempty"`
	InstanceId       *string                                               `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	InstanceTitle    *string                                               `json:"instanceTitle,omitempty" xml:"instanceTitle,omitempty"`
	InstanceUrl      *string                                               `json:"instanceUrl,omitempty" xml:"instanceUrl,omitempty"`
	OtherAccountName *string                                               `json:"otherAccountName,omitempty" xml:"otherAccountName,omitempty"`
	OtherAccountNo   *string                                               `json:"otherAccountNo,omitempty" xml:"otherAccountNo,omitempty"`
	ReceiptFile      *QueryAccountTradeByPageResponseBodyResultReceiptFile `json:"receiptFile,omitempty" xml:"receiptFile,omitempty" type:"Struct"`
	Remark           *string                                               `json:"remark,omitempty" xml:"remark,omitempty"`
	TradeAmount      *string                                               `json:"tradeAmount,omitempty" xml:"tradeAmount,omitempty"`
	TradeNo          *string                                               `json:"tradeNo,omitempty" xml:"tradeNo,omitempty"`
	TradeTime        *int64                                                `json:"tradeTime,omitempty" xml:"tradeTime,omitempty"`
	TradeType        *string                                               `json:"tradeType,omitempty" xml:"tradeType,omitempty"`
}

func (s QueryAccountTradeByPageResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountTradeByPageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryAccountTradeByPageResponseBodyResult) SetBalance(v string) *QueryAccountTradeByPageResponseBodyResult {
	s.Balance = &v
	return s
}

func (s *QueryAccountTradeByPageResponseBodyResult) SetDetailId(v string) *QueryAccountTradeByPageResponseBodyResult {
	s.DetailId = &v
	return s
}

func (s *QueryAccountTradeByPageResponseBodyResult) SetInstanceId(v string) *QueryAccountTradeByPageResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *QueryAccountTradeByPageResponseBodyResult) SetInstanceTitle(v string) *QueryAccountTradeByPageResponseBodyResult {
	s.InstanceTitle = &v
	return s
}

func (s *QueryAccountTradeByPageResponseBodyResult) SetInstanceUrl(v string) *QueryAccountTradeByPageResponseBodyResult {
	s.InstanceUrl = &v
	return s
}

func (s *QueryAccountTradeByPageResponseBodyResult) SetOtherAccountName(v string) *QueryAccountTradeByPageResponseBodyResult {
	s.OtherAccountName = &v
	return s
}

func (s *QueryAccountTradeByPageResponseBodyResult) SetOtherAccountNo(v string) *QueryAccountTradeByPageResponseBodyResult {
	s.OtherAccountNo = &v
	return s
}

func (s *QueryAccountTradeByPageResponseBodyResult) SetReceiptFile(v *QueryAccountTradeByPageResponseBodyResultReceiptFile) *QueryAccountTradeByPageResponseBodyResult {
	s.ReceiptFile = v
	return s
}

func (s *QueryAccountTradeByPageResponseBodyResult) SetRemark(v string) *QueryAccountTradeByPageResponseBodyResult {
	s.Remark = &v
	return s
}

func (s *QueryAccountTradeByPageResponseBodyResult) SetTradeAmount(v string) *QueryAccountTradeByPageResponseBodyResult {
	s.TradeAmount = &v
	return s
}

func (s *QueryAccountTradeByPageResponseBodyResult) SetTradeNo(v string) *QueryAccountTradeByPageResponseBodyResult {
	s.TradeNo = &v
	return s
}

func (s *QueryAccountTradeByPageResponseBodyResult) SetTradeTime(v int64) *QueryAccountTradeByPageResponseBodyResult {
	s.TradeTime = &v
	return s
}

func (s *QueryAccountTradeByPageResponseBodyResult) SetTradeType(v string) *QueryAccountTradeByPageResponseBodyResult {
	s.TradeType = &v
	return s
}

type QueryAccountTradeByPageResponseBodyResultReceiptFile struct {
	FileId     *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	FileName   *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileSize   *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	FileType   *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	PreviewUrl *string `json:"previewUrl,omitempty" xml:"previewUrl,omitempty"`
	SpaceId    *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s QueryAccountTradeByPageResponseBodyResultReceiptFile) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountTradeByPageResponseBodyResultReceiptFile) GoString() string {
	return s.String()
}

func (s *QueryAccountTradeByPageResponseBodyResultReceiptFile) SetFileId(v string) *QueryAccountTradeByPageResponseBodyResultReceiptFile {
	s.FileId = &v
	return s
}

func (s *QueryAccountTradeByPageResponseBodyResultReceiptFile) SetFileName(v string) *QueryAccountTradeByPageResponseBodyResultReceiptFile {
	s.FileName = &v
	return s
}

func (s *QueryAccountTradeByPageResponseBodyResultReceiptFile) SetFileSize(v int64) *QueryAccountTradeByPageResponseBodyResultReceiptFile {
	s.FileSize = &v
	return s
}

func (s *QueryAccountTradeByPageResponseBodyResultReceiptFile) SetFileType(v string) *QueryAccountTradeByPageResponseBodyResultReceiptFile {
	s.FileType = &v
	return s
}

func (s *QueryAccountTradeByPageResponseBodyResultReceiptFile) SetPreviewUrl(v string) *QueryAccountTradeByPageResponseBodyResultReceiptFile {
	s.PreviewUrl = &v
	return s
}

func (s *QueryAccountTradeByPageResponseBodyResultReceiptFile) SetSpaceId(v string) *QueryAccountTradeByPageResponseBodyResultReceiptFile {
	s.SpaceId = &v
	return s
}

type QueryAccountTradeByPageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAccountTradeByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAccountTradeByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountTradeByPageResponse) GoString() string {
	return s.String()
}

func (s *QueryAccountTradeByPageResponse) SetHeaders(v map[string]*string) *QueryAccountTradeByPageResponse {
	s.Headers = v
	return s
}

func (s *QueryAccountTradeByPageResponse) SetStatusCode(v int32) *QueryAccountTradeByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAccountTradeByPageResponse) SetBody(v *QueryAccountTradeByPageResponseBody) *QueryAccountTradeByPageResponse {
	s.Body = v
	return s
}

type QueryCategoryByPageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCategoryByPageHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCategoryByPageHeaders) GoString() string {
	return s.String()
}

func (s *QueryCategoryByPageHeaders) SetCommonHeaders(v map[string]*string) *QueryCategoryByPageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCategoryByPageHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCategoryByPageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCategoryByPageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryCategoryByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCategoryByPageRequest) GoString() string {
	return s.String()
}

func (s *QueryCategoryByPageRequest) SetPageNumber(v int64) *QueryCategoryByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryCategoryByPageRequest) SetPageSize(v int64) *QueryCategoryByPageRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCategoryByPageRequest) SetType(v string) *QueryCategoryByPageRequest {
	s.Type = &v
	return s
}

type QueryCategoryByPageResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// This parameter is required.
	List []*QueryCategoryByPageResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s QueryCategoryByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCategoryByPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCategoryByPageResponseBody) SetHasMore(v bool) *QueryCategoryByPageResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryCategoryByPageResponseBody) SetList(v []*QueryCategoryByPageResponseBodyList) *QueryCategoryByPageResponseBody {
	s.List = v
	return s
}

type QueryCategoryByPageResponseBodyList struct {
	// This parameter is required.
	//
	// example:
	//
	// INCOME_XXX
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	IsDir *bool `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 汽车
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// INCOM_XXX
	ParentCode *string `json:"parentCode,omitempty" xml:"parentCode,omitempty"`
	Remark     *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// valid
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// income
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryCategoryByPageResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryCategoryByPageResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryCategoryByPageResponseBodyList) SetCode(v string) *QueryCategoryByPageResponseBodyList {
	s.Code = &v
	return s
}

func (s *QueryCategoryByPageResponseBodyList) SetIsDir(v bool) *QueryCategoryByPageResponseBodyList {
	s.IsDir = &v
	return s
}

func (s *QueryCategoryByPageResponseBodyList) SetName(v string) *QueryCategoryByPageResponseBodyList {
	s.Name = &v
	return s
}

func (s *QueryCategoryByPageResponseBodyList) SetParentCode(v string) *QueryCategoryByPageResponseBodyList {
	s.ParentCode = &v
	return s
}

func (s *QueryCategoryByPageResponseBodyList) SetRemark(v string) *QueryCategoryByPageResponseBodyList {
	s.Remark = &v
	return s
}

func (s *QueryCategoryByPageResponseBodyList) SetStatus(v string) *QueryCategoryByPageResponseBodyList {
	s.Status = &v
	return s
}

func (s *QueryCategoryByPageResponseBodyList) SetType(v string) *QueryCategoryByPageResponseBodyList {
	s.Type = &v
	return s
}

type QueryCategoryByPageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCategoryByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCategoryByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCategoryByPageResponse) GoString() string {
	return s.String()
}

func (s *QueryCategoryByPageResponse) SetHeaders(v map[string]*string) *QueryCategoryByPageResponse {
	s.Headers = v
	return s
}

func (s *QueryCategoryByPageResponse) SetStatusCode(v int32) *QueryCategoryByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCategoryByPageResponse) SetBody(v *QueryCategoryByPageResponseBody) *QueryCategoryByPageResponse {
	s.Body = v
	return s
}

type QueryCollectionInfoListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCollectionInfoListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCollectionInfoListHeaders) GoString() string {
	return s.String()
}

func (s *QueryCollectionInfoListHeaders) SetCommonHeaders(v map[string]*string) *QueryCollectionInfoListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCollectionInfoListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCollectionInfoListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCollectionInfoListRequest struct {
	// if can be null:
	// true
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryCollectionInfoListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCollectionInfoListRequest) GoString() string {
	return s.String()
}

func (s *QueryCollectionInfoListRequest) SetStatus(v string) *QueryCollectionInfoListRequest {
	s.Status = &v
	return s
}

type QueryCollectionInfoListResponseBody struct {
	CollectionInfoList []*QueryCollectionInfoListResponseBodyCollectionInfoList `json:"collectionInfoList,omitempty" xml:"collectionInfoList,omitempty" type:"Repeated"`
}

func (s QueryCollectionInfoListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCollectionInfoListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCollectionInfoListResponseBody) SetCollectionInfoList(v []*QueryCollectionInfoListResponseBodyCollectionInfoList) *QueryCollectionInfoListResponseBody {
	s.CollectionInfoList = v
	return s
}

type QueryCollectionInfoListResponseBodyCollectionInfoList struct {
	AccountHolderName *string `json:"accountHolderName,omitempty" xml:"accountHolderName,omitempty"`
	AlipayLogonId     *string `json:"alipayLogonId,omitempty" xml:"alipayLogonId,omitempty"`
	AuditStatus       *string `json:"auditStatus,omitempty" xml:"auditStatus,omitempty"`
	CertNo            *string `json:"certNo,omitempty" xml:"certNo,omitempty"`
	CollectionInfoId  *string `json:"collectionInfoId,omitempty" xml:"collectionInfoId,omitempty"`
	FailReason        *string `json:"failReason,omitempty" xml:"failReason,omitempty"`
	GmtAudit          *int64  `json:"gmtAudit,omitempty" xml:"gmtAudit,omitempty"`
	MerchantName      *string `json:"merchantName,omitempty" xml:"merchantName,omitempty"`
	Type              *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryCollectionInfoListResponseBodyCollectionInfoList) String() string {
	return tea.Prettify(s)
}

func (s QueryCollectionInfoListResponseBodyCollectionInfoList) GoString() string {
	return s.String()
}

func (s *QueryCollectionInfoListResponseBodyCollectionInfoList) SetAccountHolderName(v string) *QueryCollectionInfoListResponseBodyCollectionInfoList {
	s.AccountHolderName = &v
	return s
}

func (s *QueryCollectionInfoListResponseBodyCollectionInfoList) SetAlipayLogonId(v string) *QueryCollectionInfoListResponseBodyCollectionInfoList {
	s.AlipayLogonId = &v
	return s
}

func (s *QueryCollectionInfoListResponseBodyCollectionInfoList) SetAuditStatus(v string) *QueryCollectionInfoListResponseBodyCollectionInfoList {
	s.AuditStatus = &v
	return s
}

func (s *QueryCollectionInfoListResponseBodyCollectionInfoList) SetCertNo(v string) *QueryCollectionInfoListResponseBodyCollectionInfoList {
	s.CertNo = &v
	return s
}

func (s *QueryCollectionInfoListResponseBodyCollectionInfoList) SetCollectionInfoId(v string) *QueryCollectionInfoListResponseBodyCollectionInfoList {
	s.CollectionInfoId = &v
	return s
}

func (s *QueryCollectionInfoListResponseBodyCollectionInfoList) SetFailReason(v string) *QueryCollectionInfoListResponseBodyCollectionInfoList {
	s.FailReason = &v
	return s
}

func (s *QueryCollectionInfoListResponseBodyCollectionInfoList) SetGmtAudit(v int64) *QueryCollectionInfoListResponseBodyCollectionInfoList {
	s.GmtAudit = &v
	return s
}

func (s *QueryCollectionInfoListResponseBodyCollectionInfoList) SetMerchantName(v string) *QueryCollectionInfoListResponseBodyCollectionInfoList {
	s.MerchantName = &v
	return s
}

func (s *QueryCollectionInfoListResponseBodyCollectionInfoList) SetType(v string) *QueryCollectionInfoListResponseBodyCollectionInfoList {
	s.Type = &v
	return s
}

type QueryCollectionInfoListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCollectionInfoListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCollectionInfoListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCollectionInfoListResponse) GoString() string {
	return s.String()
}

func (s *QueryCollectionInfoListResponse) SetHeaders(v map[string]*string) *QueryCollectionInfoListResponse {
	s.Headers = v
	return s
}

func (s *QueryCollectionInfoListResponse) SetStatusCode(v int32) *QueryCollectionInfoListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCollectionInfoListResponse) SetBody(v *QueryCollectionInfoListResponseBody) *QueryCollectionInfoListResponse {
	s.Body = v
	return s
}

type QueryCollectionOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCollectionOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCollectionOrderHeaders) GoString() string {
	return s.String()
}

func (s *QueryCollectionOrderHeaders) SetCommonHeaders(v map[string]*string) *QueryCollectionOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCollectionOrderHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCollectionOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCollectionOrderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s QueryCollectionOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCollectionOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryCollectionOrderRequest) SetInstanceId(v string) *QueryCollectionOrderRequest {
	s.InstanceId = &v
	return s
}

type QueryCollectionOrderResponseBody struct {
	Amount     *string `json:"amount,omitempty" xml:"amount,omitempty"`
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Remark     *string `json:"remark,omitempty" xml:"remark,omitempty"`
	Status     *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryCollectionOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCollectionOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCollectionOrderResponseBody) SetAmount(v string) *QueryCollectionOrderResponseBody {
	s.Amount = &v
	return s
}

func (s *QueryCollectionOrderResponseBody) SetInstanceId(v string) *QueryCollectionOrderResponseBody {
	s.InstanceId = &v
	return s
}

func (s *QueryCollectionOrderResponseBody) SetRemark(v string) *QueryCollectionOrderResponseBody {
	s.Remark = &v
	return s
}

func (s *QueryCollectionOrderResponseBody) SetStatus(v string) *QueryCollectionOrderResponseBody {
	s.Status = &v
	return s
}

type QueryCollectionOrderResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCollectionOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCollectionOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCollectionOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryCollectionOrderResponse) SetHeaders(v map[string]*string) *QueryCollectionOrderResponse {
	s.Headers = v
	return s
}

func (s *QueryCollectionOrderResponse) SetStatusCode(v int32) *QueryCollectionOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCollectionOrderResponse) SetBody(v *QueryCollectionOrderResponseBody) *QueryCollectionOrderResponse {
	s.Body = v
	return s
}

type QueryCustomerByPageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCustomerByPageHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerByPageHeaders) GoString() string {
	return s.String()
}

func (s *QueryCustomerByPageHeaders) SetCommonHeaders(v map[string]*string) *QueryCustomerByPageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCustomerByPageHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCustomerByPageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCustomerByPageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryCustomerByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerByPageRequest) GoString() string {
	return s.String()
}

func (s *QueryCustomerByPageRequest) SetPageNumber(v int64) *QueryCustomerByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryCustomerByPageRequest) SetPageSize(v int64) *QueryCustomerByPageRequest {
	s.PageSize = &v
	return s
}

type QueryCustomerByPageResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// This parameter is required.
	List []*QueryCustomerByPageResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s QueryCustomerByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerByPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCustomerByPageResponseBody) SetHasMore(v bool) *QueryCustomerByPageResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryCustomerByPageResponseBody) SetList(v []*QueryCustomerByPageResponseBodyList) *QueryCustomerByPageResponseBody {
	s.List = v
	return s
}

type QueryCustomerByPageResponseBodyList struct {
	// This parameter is required.
	//
	// example:
	//
	// CUS_XXX
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1634786828686
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 重要客户
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// XX有限公司
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// valid
	Status         *string `json:"status,omitempty" xml:"status,omitempty"`
	UserDefineCode *string `json:"userDefineCode,omitempty" xml:"userDefineCode,omitempty"`
}

func (s QueryCustomerByPageResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerByPageResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryCustomerByPageResponseBodyList) SetCode(v string) *QueryCustomerByPageResponseBodyList {
	s.Code = &v
	return s
}

func (s *QueryCustomerByPageResponseBodyList) SetCreateTime(v int64) *QueryCustomerByPageResponseBodyList {
	s.CreateTime = &v
	return s
}

func (s *QueryCustomerByPageResponseBodyList) SetDescription(v string) *QueryCustomerByPageResponseBodyList {
	s.Description = &v
	return s
}

func (s *QueryCustomerByPageResponseBodyList) SetName(v string) *QueryCustomerByPageResponseBodyList {
	s.Name = &v
	return s
}

func (s *QueryCustomerByPageResponseBodyList) SetStatus(v string) *QueryCustomerByPageResponseBodyList {
	s.Status = &v
	return s
}

func (s *QueryCustomerByPageResponseBodyList) SetUserDefineCode(v string) *QueryCustomerByPageResponseBodyList {
	s.UserDefineCode = &v
	return s
}

type QueryCustomerByPageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCustomerByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCustomerByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerByPageResponse) GoString() string {
	return s.String()
}

func (s *QueryCustomerByPageResponse) SetHeaders(v map[string]*string) *QueryCustomerByPageResponse {
	s.Headers = v
	return s
}

func (s *QueryCustomerByPageResponse) SetStatusCode(v int32) *QueryCustomerByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCustomerByPageResponse) SetBody(v *QueryCustomerByPageResponseBody) *QueryCustomerByPageResponse {
	s.Body = v
	return s
}

type QueryEnterpriseAccountByPageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryEnterpriseAccountByPageHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryEnterpriseAccountByPageHeaders) GoString() string {
	return s.String()
}

func (s *QueryEnterpriseAccountByPageHeaders) SetCommonHeaders(v map[string]*string) *QueryEnterpriseAccountByPageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryEnterpriseAccountByPageHeaders) SetXAcsDingtalkAccessToken(v string) *QueryEnterpriseAccountByPageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryEnterpriseAccountByPageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryEnterpriseAccountByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEnterpriseAccountByPageRequest) GoString() string {
	return s.String()
}

func (s *QueryEnterpriseAccountByPageRequest) SetPageNumber(v int64) *QueryEnterpriseAccountByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryEnterpriseAccountByPageRequest) SetPageSize(v int64) *QueryEnterpriseAccountByPageRequest {
	s.PageSize = &v
	return s
}

type QueryEnterpriseAccountByPageResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// This parameter is required.
	List []*QueryEnterpriseAccountByPageResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s QueryEnterpriseAccountByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEnterpriseAccountByPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEnterpriseAccountByPageResponseBody) SetHasMore(v bool) *QueryEnterpriseAccountByPageResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryEnterpriseAccountByPageResponseBody) SetList(v []*QueryEnterpriseAccountByPageResponseBodyList) *QueryEnterpriseAccountByPageResponseBody {
	s.List = v
	return s
}

type QueryEnterpriseAccountByPageResponseBodyList struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	AccountCode *string `json:"accountCode,omitempty" xml:"accountCode,omitempty"`
	// example:
	//
	// test@alipay.com
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 网商银行
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	// example:
	//
	// test
	AccountRemark *string `json:"accountRemark,omitempty" xml:"accountRemark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALIPAY
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	// example:
	//
	// 10000.33
	Amount      *string `json:"amount,omitempty" xml:"amount,omitempty"`
	BankCode    *string `json:"bankCode,omitempty" xml:"bankCode,omitempty"`
	BankName    *string `json:"bankName,omitempty" xml:"bankName,omitempty"`
	CompanyCode *string `json:"companyCode,omitempty" xml:"companyCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1631526550994
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aaa
	Creator            *string `json:"creator,omitempty" xml:"creator,omitempty"`
	OfficialName       *string `json:"officialName,omitempty" xml:"officialName,omitempty"`
	OfficialNumber     *string `json:"officialNumber,omitempty" xml:"officialNumber,omitempty"`
	SignStatus         *string `json:"signStatus,omitempty" xml:"signStatus,omitempty"`
	SupportReceipt     *bool   `json:"supportReceipt,omitempty" xml:"supportReceipt,omitempty"`
	SupportTradeDetail *bool   `json:"supportTradeDetail,omitempty" xml:"supportTradeDetail,omitempty"`
}

func (s QueryEnterpriseAccountByPageResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryEnterpriseAccountByPageResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryEnterpriseAccountByPageResponseBodyList) SetAccountCode(v string) *QueryEnterpriseAccountByPageResponseBodyList {
	s.AccountCode = &v
	return s
}

func (s *QueryEnterpriseAccountByPageResponseBodyList) SetAccountId(v string) *QueryEnterpriseAccountByPageResponseBodyList {
	s.AccountId = &v
	return s
}

func (s *QueryEnterpriseAccountByPageResponseBodyList) SetAccountName(v string) *QueryEnterpriseAccountByPageResponseBodyList {
	s.AccountName = &v
	return s
}

func (s *QueryEnterpriseAccountByPageResponseBodyList) SetAccountRemark(v string) *QueryEnterpriseAccountByPageResponseBodyList {
	s.AccountRemark = &v
	return s
}

func (s *QueryEnterpriseAccountByPageResponseBodyList) SetAccountType(v string) *QueryEnterpriseAccountByPageResponseBodyList {
	s.AccountType = &v
	return s
}

func (s *QueryEnterpriseAccountByPageResponseBodyList) SetAmount(v string) *QueryEnterpriseAccountByPageResponseBodyList {
	s.Amount = &v
	return s
}

func (s *QueryEnterpriseAccountByPageResponseBodyList) SetBankCode(v string) *QueryEnterpriseAccountByPageResponseBodyList {
	s.BankCode = &v
	return s
}

func (s *QueryEnterpriseAccountByPageResponseBodyList) SetBankName(v string) *QueryEnterpriseAccountByPageResponseBodyList {
	s.BankName = &v
	return s
}

func (s *QueryEnterpriseAccountByPageResponseBodyList) SetCompanyCode(v string) *QueryEnterpriseAccountByPageResponseBodyList {
	s.CompanyCode = &v
	return s
}

func (s *QueryEnterpriseAccountByPageResponseBodyList) SetCreateTime(v int64) *QueryEnterpriseAccountByPageResponseBodyList {
	s.CreateTime = &v
	return s
}

func (s *QueryEnterpriseAccountByPageResponseBodyList) SetCreator(v string) *QueryEnterpriseAccountByPageResponseBodyList {
	s.Creator = &v
	return s
}

func (s *QueryEnterpriseAccountByPageResponseBodyList) SetOfficialName(v string) *QueryEnterpriseAccountByPageResponseBodyList {
	s.OfficialName = &v
	return s
}

func (s *QueryEnterpriseAccountByPageResponseBodyList) SetOfficialNumber(v string) *QueryEnterpriseAccountByPageResponseBodyList {
	s.OfficialNumber = &v
	return s
}

func (s *QueryEnterpriseAccountByPageResponseBodyList) SetSignStatus(v string) *QueryEnterpriseAccountByPageResponseBodyList {
	s.SignStatus = &v
	return s
}

func (s *QueryEnterpriseAccountByPageResponseBodyList) SetSupportReceipt(v bool) *QueryEnterpriseAccountByPageResponseBodyList {
	s.SupportReceipt = &v
	return s
}

func (s *QueryEnterpriseAccountByPageResponseBodyList) SetSupportTradeDetail(v bool) *QueryEnterpriseAccountByPageResponseBodyList {
	s.SupportTradeDetail = &v
	return s
}

type QueryEnterpriseAccountByPageResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEnterpriseAccountByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEnterpriseAccountByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEnterpriseAccountByPageResponse) GoString() string {
	return s.String()
}

func (s *QueryEnterpriseAccountByPageResponse) SetHeaders(v map[string]*string) *QueryEnterpriseAccountByPageResponse {
	s.Headers = v
	return s
}

func (s *QueryEnterpriseAccountByPageResponse) SetStatusCode(v int32) *QueryEnterpriseAccountByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEnterpriseAccountByPageResponse) SetBody(v *QueryEnterpriseAccountByPageResponseBody) *QueryEnterpriseAccountByPageResponse {
	s.Body = v
	return s
}

type QueryEnterpriseAccountSignUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryEnterpriseAccountSignUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryEnterpriseAccountSignUrlHeaders) GoString() string {
	return s.String()
}

func (s *QueryEnterpriseAccountSignUrlHeaders) SetCommonHeaders(v map[string]*string) *QueryEnterpriseAccountSignUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryEnterpriseAccountSignUrlHeaders) SetXAcsDingtalkAccessToken(v string) *QueryEnterpriseAccountSignUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryEnterpriseAccountSignUrlRequest struct {
	// example:
	//
	// ACC_X12133
	AccountCode *string `json:"accountCode,omitempty" xml:"accountCode,omitempty"`
	// example:
	//
	// 5041145245
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryEnterpriseAccountSignUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEnterpriseAccountSignUrlRequest) GoString() string {
	return s.String()
}

func (s *QueryEnterpriseAccountSignUrlRequest) SetAccountCode(v string) *QueryEnterpriseAccountSignUrlRequest {
	s.AccountCode = &v
	return s
}

func (s *QueryEnterpriseAccountSignUrlRequest) SetUserId(v string) *QueryEnterpriseAccountSignUrlRequest {
	s.UserId = &v
	return s
}

type QueryEnterpriseAccountSignUrlResponseBody struct {
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s QueryEnterpriseAccountSignUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEnterpriseAccountSignUrlResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEnterpriseAccountSignUrlResponseBody) SetUrl(v string) *QueryEnterpriseAccountSignUrlResponseBody {
	s.Url = &v
	return s
}

type QueryEnterpriseAccountSignUrlResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEnterpriseAccountSignUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEnterpriseAccountSignUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEnterpriseAccountSignUrlResponse) GoString() string {
	return s.String()
}

func (s *QueryEnterpriseAccountSignUrlResponse) SetHeaders(v map[string]*string) *QueryEnterpriseAccountSignUrlResponse {
	s.Headers = v
	return s
}

func (s *QueryEnterpriseAccountSignUrlResponse) SetStatusCode(v int32) *QueryEnterpriseAccountSignUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEnterpriseAccountSignUrlResponse) SetBody(v *QueryEnterpriseAccountSignUrlResponseBody) *QueryEnterpriseAccountSignUrlResponse {
	s.Body = v
	return s
}

type QueryInstancePaymentOrderDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryInstancePaymentOrderDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryInstancePaymentOrderDetailHeaders) GoString() string {
	return s.String()
}

func (s *QueryInstancePaymentOrderDetailHeaders) SetCommonHeaders(v map[string]*string) *QueryInstancePaymentOrderDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryInstancePaymentOrderDetailHeaders) SetXAcsDingtalkAccessToken(v string) *QueryInstancePaymentOrderDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryInstancePaymentOrderDetailRequest struct {
	// example:
	//
	// abc
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
}

func (s QueryInstancePaymentOrderDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInstancePaymentOrderDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryInstancePaymentOrderDetailRequest) SetOrderNo(v string) *QueryInstancePaymentOrderDetailRequest {
	s.OrderNo = &v
	return s
}

type QueryInstancePaymentOrderDetailResponseBody struct {
	Amount          *string                                                     `json:"amount,omitempty" xml:"amount,omitempty"`
	InstanceId      *string                                                     `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	PayeeAccountDTO *QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTO `json:"payeeAccountDTO,omitempty" xml:"payeeAccountDTO,omitempty" type:"Struct"`
	PayerAccountDTO *QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTO `json:"payerAccountDTO,omitempty" xml:"payerAccountDTO,omitempty" type:"Struct"`
	Remark          *string                                                     `json:"remark,omitempty" xml:"remark,omitempty"`
	Usage           *string                                                     `json:"usage,omitempty" xml:"usage,omitempty"`
	UserId          *string                                                     `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryInstancePaymentOrderDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInstancePaymentOrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInstancePaymentOrderDetailResponseBody) SetAmount(v string) *QueryInstancePaymentOrderDetailResponseBody {
	s.Amount = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBody) SetInstanceId(v string) *QueryInstancePaymentOrderDetailResponseBody {
	s.InstanceId = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBody) SetPayeeAccountDTO(v *QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTO) *QueryInstancePaymentOrderDetailResponseBody {
	s.PayeeAccountDTO = v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBody) SetPayerAccountDTO(v *QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTO) *QueryInstancePaymentOrderDetailResponseBody {
	s.PayerAccountDTO = v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBody) SetRemark(v string) *QueryInstancePaymentOrderDetailResponseBody {
	s.Remark = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBody) SetUsage(v string) *QueryInstancePaymentOrderDetailResponseBody {
	s.Usage = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBody) SetUserId(v string) *QueryInstancePaymentOrderDetailResponseBody {
	s.UserId = &v
	return s
}

type QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTO struct {
	BankOpenDTO *QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTOBankOpenDTO `json:"bankOpenDTO,omitempty" xml:"bankOpenDTO,omitempty" type:"Struct"`
}

func (s QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTO) String() string {
	return tea.Prettify(s)
}

func (s QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTO) GoString() string {
	return s.String()
}

func (s *QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTO) SetBankOpenDTO(v *QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTOBankOpenDTO) *QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTO {
	s.BankOpenDTO = v
	return s
}

type QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTOBankOpenDTO struct {
	AccountName    *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	BankBranchCode *string `json:"bankBranchCode,omitempty" xml:"bankBranchCode,omitempty"`
	BankBranchName *string `json:"bankBranchName,omitempty" xml:"bankBranchName,omitempty"`
	BankCardNo     *string `json:"bankCardNo,omitempty" xml:"bankCardNo,omitempty"`
	BankCode       *string `json:"bankCode,omitempty" xml:"bankCode,omitempty"`
	BankName       *string `json:"bankName,omitempty" xml:"bankName,omitempty"`
	Type           *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTOBankOpenDTO) String() string {
	return tea.Prettify(s)
}

func (s QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTOBankOpenDTO) GoString() string {
	return s.String()
}

func (s *QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTOBankOpenDTO) SetAccountName(v string) *QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTOBankOpenDTO {
	s.AccountName = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTOBankOpenDTO) SetBankBranchCode(v string) *QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTOBankOpenDTO {
	s.BankBranchCode = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTOBankOpenDTO) SetBankBranchName(v string) *QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTOBankOpenDTO {
	s.BankBranchName = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTOBankOpenDTO) SetBankCardNo(v string) *QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTOBankOpenDTO {
	s.BankCardNo = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTOBankOpenDTO) SetBankCode(v string) *QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTOBankOpenDTO {
	s.BankCode = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTOBankOpenDTO) SetBankName(v string) *QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTOBankOpenDTO {
	s.BankName = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTOBankOpenDTO) SetType(v string) *QueryInstancePaymentOrderDetailResponseBodyPayeeAccountDTOBankOpenDTO {
	s.Type = &v
	return s
}

type QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTO struct {
	BankOpenDTO           *QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTOBankOpenDTO `json:"bankOpenDTO,omitempty" xml:"bankOpenDTO,omitempty" type:"Struct"`
	EnterpriseAccountCode *string                                                                `json:"enterpriseAccountCode,omitempty" xml:"enterpriseAccountCode,omitempty"`
}

func (s QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTO) String() string {
	return tea.Prettify(s)
}

func (s QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTO) GoString() string {
	return s.String()
}

func (s *QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTO) SetBankOpenDTO(v *QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTOBankOpenDTO) *QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTO {
	s.BankOpenDTO = v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTO) SetEnterpriseAccountCode(v string) *QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTO {
	s.EnterpriseAccountCode = &v
	return s
}

type QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTOBankOpenDTO struct {
	AccountName    *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	BankBranchCode *string `json:"bankBranchCode,omitempty" xml:"bankBranchCode,omitempty"`
	BankBranchName *string `json:"bankBranchName,omitempty" xml:"bankBranchName,omitempty"`
	BankCardNo     *string `json:"bankCardNo,omitempty" xml:"bankCardNo,omitempty"`
	BankCode       *string `json:"bankCode,omitempty" xml:"bankCode,omitempty"`
	BankName       *string `json:"bankName,omitempty" xml:"bankName,omitempty"`
	Type           *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTOBankOpenDTO) String() string {
	return tea.Prettify(s)
}

func (s QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTOBankOpenDTO) GoString() string {
	return s.String()
}

func (s *QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTOBankOpenDTO) SetAccountName(v string) *QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTOBankOpenDTO {
	s.AccountName = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTOBankOpenDTO) SetBankBranchCode(v string) *QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTOBankOpenDTO {
	s.BankBranchCode = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTOBankOpenDTO) SetBankBranchName(v string) *QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTOBankOpenDTO {
	s.BankBranchName = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTOBankOpenDTO) SetBankCardNo(v string) *QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTOBankOpenDTO {
	s.BankCardNo = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTOBankOpenDTO) SetBankCode(v string) *QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTOBankOpenDTO {
	s.BankCode = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTOBankOpenDTO) SetBankName(v string) *QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTOBankOpenDTO {
	s.BankName = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTOBankOpenDTO) SetType(v string) *QueryInstancePaymentOrderDetailResponseBodyPayerAccountDTOBankOpenDTO {
	s.Type = &v
	return s
}

type QueryInstancePaymentOrderDetailResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInstancePaymentOrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInstancePaymentOrderDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryInstancePaymentOrderDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryInstancePaymentOrderDetailResponse) SetHeaders(v map[string]*string) *QueryInstancePaymentOrderDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponse) SetStatusCode(v int32) *QueryInstancePaymentOrderDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponse) SetBody(v *QueryInstancePaymentOrderDetailResponseBody) *QueryInstancePaymentOrderDetailResponse {
	s.Body = v
	return s
}

type QueryInvoiceTransferDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryInvoiceTransferDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoiceTransferDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryInvoiceTransferDataHeaders) SetCommonHeaders(v map[string]*string) *QueryInvoiceTransferDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryInvoiceTransferDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryInvoiceTransferDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryInvoiceTransferDataRequest struct {
	Body *QueryInvoiceTransferDataRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s QueryInvoiceTransferDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoiceTransferDataRequest) GoString() string {
	return s.String()
}

func (s *QueryInvoiceTransferDataRequest) SetBody(v *QueryInvoiceTransferDataRequestBody) *QueryInvoiceTransferDataRequest {
	s.Body = v
	return s
}

type QueryInvoiceTransferDataRequestBody struct {
	Keys []*string `json:"keys,omitempty" xml:"keys,omitempty" type:"Repeated"`
}

func (s QueryInvoiceTransferDataRequestBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoiceTransferDataRequestBody) GoString() string {
	return s.String()
}

func (s *QueryInvoiceTransferDataRequestBody) SetKeys(v []*string) *QueryInvoiceTransferDataRequestBody {
	s.Keys = v
	return s
}

type QueryInvoiceTransferDataShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInvoiceTransferDataShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoiceTransferDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryInvoiceTransferDataShrinkRequest) SetBodyShrink(v string) *QueryInvoiceTransferDataShrinkRequest {
	s.BodyShrink = &v
	return s
}

type QueryInvoiceTransferDataResponseBody struct {
	KeyToData map[string]*string `json:"keyToData,omitempty" xml:"keyToData,omitempty"`
}

func (s QueryInvoiceTransferDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoiceTransferDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInvoiceTransferDataResponseBody) SetKeyToData(v map[string]*string) *QueryInvoiceTransferDataResponseBody {
	s.KeyToData = v
	return s
}

type QueryInvoiceTransferDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInvoiceTransferDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInvoiceTransferDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoiceTransferDataResponse) GoString() string {
	return s.String()
}

func (s *QueryInvoiceTransferDataResponse) SetHeaders(v map[string]*string) *QueryInvoiceTransferDataResponse {
	s.Headers = v
	return s
}

func (s *QueryInvoiceTransferDataResponse) SetStatusCode(v int32) *QueryInvoiceTransferDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInvoiceTransferDataResponse) SetBody(v *QueryInvoiceTransferDataResponseBody) *QueryInvoiceTransferDataResponse {
	s.Body = v
	return s
}

type QueryPaymentBenefitHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryPaymentBenefitHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentBenefitHeaders) GoString() string {
	return s.String()
}

func (s *QueryPaymentBenefitHeaders) SetCommonHeaders(v map[string]*string) *QueryPaymentBenefitHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryPaymentBenefitHeaders) SetXAcsDingtalkAccessToken(v string) *QueryPaymentBenefitHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryPaymentBenefitRequest struct {
	// example:
	//
	// 5042215136
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryPaymentBenefitRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentBenefitRequest) GoString() string {
	return s.String()
}

func (s *QueryPaymentBenefitRequest) SetUserId(v string) *QueryPaymentBenefitRequest {
	s.UserId = &v
	return s
}

type QueryPaymentBenefitResponseBody struct {
	BenefitMap map[string]*BenefitMapValue `json:"benefitMap,omitempty" xml:"benefitMap,omitempty"`
	CorpId     *string                     `json:"corpId,omitempty" xml:"corpId,omitempty"`
	RequestId  *string                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryPaymentBenefitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentBenefitResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPaymentBenefitResponseBody) SetBenefitMap(v map[string]*BenefitMapValue) *QueryPaymentBenefitResponseBody {
	s.BenefitMap = v
	return s
}

func (s *QueryPaymentBenefitResponseBody) SetCorpId(v string) *QueryPaymentBenefitResponseBody {
	s.CorpId = &v
	return s
}

func (s *QueryPaymentBenefitResponseBody) SetRequestId(v string) *QueryPaymentBenefitResponseBody {
	s.RequestId = &v
	return s
}

type QueryPaymentBenefitResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPaymentBenefitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPaymentBenefitResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentBenefitResponse) GoString() string {
	return s.String()
}

func (s *QueryPaymentBenefitResponse) SetHeaders(v map[string]*string) *QueryPaymentBenefitResponse {
	s.Headers = v
	return s
}

func (s *QueryPaymentBenefitResponse) SetStatusCode(v int32) *QueryPaymentBenefitResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPaymentBenefitResponse) SetBody(v *QueryPaymentBenefitResponseBody) *QueryPaymentBenefitResponse {
	s.Body = v
	return s
}

type QueryPaymentOrderDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryPaymentOrderDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentOrderDetailHeaders) GoString() string {
	return s.String()
}

func (s *QueryPaymentOrderDetailHeaders) SetCommonHeaders(v map[string]*string) *QueryPaymentOrderDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryPaymentOrderDetailHeaders) SetXAcsDingtalkAccessToken(v string) *QueryPaymentOrderDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryPaymentOrderDetailRequest struct {
	OutBizNoList []*string `json:"outBizNoList,omitempty" xml:"outBizNoList,omitempty" type:"Repeated"`
	// example:
	//
	// 50455845112
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryPaymentOrderDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentOrderDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryPaymentOrderDetailRequest) SetOutBizNoList(v []*string) *QueryPaymentOrderDetailRequest {
	s.OutBizNoList = v
	return s
}

func (s *QueryPaymentOrderDetailRequest) SetUserId(v string) *QueryPaymentOrderDetailRequest {
	s.UserId = &v
	return s
}

type QueryPaymentOrderDetailShrinkRequest struct {
	OutBizNoListShrink *string `json:"outBizNoList,omitempty" xml:"outBizNoList,omitempty"`
	// example:
	//
	// 50455845112
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryPaymentOrderDetailShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentOrderDetailShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryPaymentOrderDetailShrinkRequest) SetOutBizNoListShrink(v string) *QueryPaymentOrderDetailShrinkRequest {
	s.OutBizNoListShrink = &v
	return s
}

func (s *QueryPaymentOrderDetailShrinkRequest) SetUserId(v string) *QueryPaymentOrderDetailShrinkRequest {
	s.UserId = &v
	return s
}

type QueryPaymentOrderDetailResponseBody struct {
	OrderList []*QueryPaymentOrderDetailResponseBodyOrderList `json:"orderList,omitempty" xml:"orderList,omitempty" type:"Repeated"`
	RequestId *string                                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryPaymentOrderDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentOrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPaymentOrderDetailResponseBody) SetOrderList(v []*QueryPaymentOrderDetailResponseBodyOrderList) *QueryPaymentOrderDetailResponseBody {
	s.OrderList = v
	return s
}

func (s *QueryPaymentOrderDetailResponseBody) SetRequestId(v string) *QueryPaymentOrderDetailResponseBody {
	s.RequestId = &v
	return s
}

type QueryPaymentOrderDetailResponseBodyOrderList struct {
	Amount          *string                                                      `json:"amount,omitempty" xml:"amount,omitempty"`
	CorpId          *string                                                      `json:"corpId,omitempty" xml:"corpId,omitempty"`
	OrderNo         *string                                                      `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	OutBizNo        *string                                                      `json:"outBizNo,omitempty" xml:"outBizNo,omitempty"`
	PayeeAccountDTO *QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTO `json:"payeeAccountDTO,omitempty" xml:"payeeAccountDTO,omitempty" type:"Struct"`
	PayerAccountDTO *QueryPaymentOrderDetailResponseBodyOrderListPayerAccountDTO `json:"payerAccountDTO,omitempty" xml:"payerAccountDTO,omitempty" type:"Struct"`
	Remark          *string                                                      `json:"remark,omitempty" xml:"remark,omitempty"`
	Status          *string                                                      `json:"status,omitempty" xml:"status,omitempty"`
	SubOrderList    []*QueryPaymentOrderDetailResponseBodyOrderListSubOrderList  `json:"subOrderList,omitempty" xml:"subOrderList,omitempty" type:"Repeated"`
	Usage           *string                                                      `json:"usage,omitempty" xml:"usage,omitempty"`
	UserId          *string                                                      `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryPaymentOrderDetailResponseBodyOrderList) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentOrderDetailResponseBodyOrderList) GoString() string {
	return s.String()
}

func (s *QueryPaymentOrderDetailResponseBodyOrderList) SetAmount(v string) *QueryPaymentOrderDetailResponseBodyOrderList {
	s.Amount = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderList) SetCorpId(v string) *QueryPaymentOrderDetailResponseBodyOrderList {
	s.CorpId = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderList) SetOrderNo(v string) *QueryPaymentOrderDetailResponseBodyOrderList {
	s.OrderNo = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderList) SetOutBizNo(v string) *QueryPaymentOrderDetailResponseBodyOrderList {
	s.OutBizNo = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderList) SetPayeeAccountDTO(v *QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTO) *QueryPaymentOrderDetailResponseBodyOrderList {
	s.PayeeAccountDTO = v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderList) SetPayerAccountDTO(v *QueryPaymentOrderDetailResponseBodyOrderListPayerAccountDTO) *QueryPaymentOrderDetailResponseBodyOrderList {
	s.PayerAccountDTO = v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderList) SetRemark(v string) *QueryPaymentOrderDetailResponseBodyOrderList {
	s.Remark = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderList) SetStatus(v string) *QueryPaymentOrderDetailResponseBodyOrderList {
	s.Status = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderList) SetSubOrderList(v []*QueryPaymentOrderDetailResponseBodyOrderListSubOrderList) *QueryPaymentOrderDetailResponseBodyOrderList {
	s.SubOrderList = v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderList) SetUsage(v string) *QueryPaymentOrderDetailResponseBodyOrderList {
	s.Usage = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderList) SetUserId(v string) *QueryPaymentOrderDetailResponseBodyOrderList {
	s.UserId = &v
	return s
}

type QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTO struct {
	BankDTO *QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTOBankDTO `json:"bankDTO,omitempty" xml:"bankDTO,omitempty" type:"Struct"`
}

func (s QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTO) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTO) GoString() string {
	return s.String()
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTO) SetBankDTO(v *QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTOBankDTO) *QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTO {
	s.BankDTO = v
	return s
}

type QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTOBankDTO struct {
	AccountName    *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	BankBranchCode *string `json:"bankBranchCode,omitempty" xml:"bankBranchCode,omitempty"`
	BankBranchName *string `json:"bankBranchName,omitempty" xml:"bankBranchName,omitempty"`
	BankCardNo     *string `json:"bankCardNo,omitempty" xml:"bankCardNo,omitempty"`
	BankCode       *string `json:"bankCode,omitempty" xml:"bankCode,omitempty"`
	BankName       *string `json:"bankName,omitempty" xml:"bankName,omitempty"`
	Type           *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTOBankDTO) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTOBankDTO) GoString() string {
	return s.String()
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTOBankDTO) SetAccountName(v string) *QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTOBankDTO {
	s.AccountName = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTOBankDTO) SetBankBranchCode(v string) *QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTOBankDTO {
	s.BankBranchCode = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTOBankDTO) SetBankBranchName(v string) *QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTOBankDTO {
	s.BankBranchName = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTOBankDTO) SetBankCardNo(v string) *QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTOBankDTO {
	s.BankCardNo = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTOBankDTO) SetBankCode(v string) *QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTOBankDTO {
	s.BankCode = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTOBankDTO) SetBankName(v string) *QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTOBankDTO {
	s.BankName = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTOBankDTO) SetType(v string) *QueryPaymentOrderDetailResponseBodyOrderListPayeeAccountDTOBankDTO {
	s.Type = &v
	return s
}

type QueryPaymentOrderDetailResponseBodyOrderListPayerAccountDTO struct {
	EnterpriseAccountCode *string `json:"enterpriseAccountCode,omitempty" xml:"enterpriseAccountCode,omitempty"`
}

func (s QueryPaymentOrderDetailResponseBodyOrderListPayerAccountDTO) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentOrderDetailResponseBodyOrderListPayerAccountDTO) GoString() string {
	return s.String()
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListPayerAccountDTO) SetEnterpriseAccountCode(v string) *QueryPaymentOrderDetailResponseBodyOrderListPayerAccountDTO {
	s.EnterpriseAccountCode = &v
	return s
}

type QueryPaymentOrderDetailResponseBodyOrderListSubOrderList struct {
	Amount          *string                                                                  `json:"amount,omitempty" xml:"amount,omitempty"`
	CorpId          *string                                                                  `json:"corpId,omitempty" xml:"corpId,omitempty"`
	OrderNo         *string                                                                  `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	OutBizNo        *string                                                                  `json:"outBizNo,omitempty" xml:"outBizNo,omitempty"`
	PayeeAccountDTO *QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTO `json:"payeeAccountDTO,omitempty" xml:"payeeAccountDTO,omitempty" type:"Struct"`
	PayerAccountDTO *QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayerAccountDTO `json:"payerAccountDTO,omitempty" xml:"payerAccountDTO,omitempty" type:"Struct"`
	Remark          *string                                                                  `json:"remark,omitempty" xml:"remark,omitempty"`
	Status          *string                                                                  `json:"status,omitempty" xml:"status,omitempty"`
	Usage           *string                                                                  `json:"usage,omitempty" xml:"usage,omitempty"`
	UserId          *string                                                                  `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryPaymentOrderDetailResponseBodyOrderListSubOrderList) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentOrderDetailResponseBodyOrderListSubOrderList) GoString() string {
	return s.String()
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListSubOrderList) SetAmount(v string) *QueryPaymentOrderDetailResponseBodyOrderListSubOrderList {
	s.Amount = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListSubOrderList) SetCorpId(v string) *QueryPaymentOrderDetailResponseBodyOrderListSubOrderList {
	s.CorpId = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListSubOrderList) SetOrderNo(v string) *QueryPaymentOrderDetailResponseBodyOrderListSubOrderList {
	s.OrderNo = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListSubOrderList) SetOutBizNo(v string) *QueryPaymentOrderDetailResponseBodyOrderListSubOrderList {
	s.OutBizNo = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListSubOrderList) SetPayeeAccountDTO(v *QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTO) *QueryPaymentOrderDetailResponseBodyOrderListSubOrderList {
	s.PayeeAccountDTO = v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListSubOrderList) SetPayerAccountDTO(v *QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayerAccountDTO) *QueryPaymentOrderDetailResponseBodyOrderListSubOrderList {
	s.PayerAccountDTO = v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListSubOrderList) SetRemark(v string) *QueryPaymentOrderDetailResponseBodyOrderListSubOrderList {
	s.Remark = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListSubOrderList) SetStatus(v string) *QueryPaymentOrderDetailResponseBodyOrderListSubOrderList {
	s.Status = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListSubOrderList) SetUsage(v string) *QueryPaymentOrderDetailResponseBodyOrderListSubOrderList {
	s.Usage = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListSubOrderList) SetUserId(v string) *QueryPaymentOrderDetailResponseBodyOrderListSubOrderList {
	s.UserId = &v
	return s
}

type QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTO struct {
	BankDTO *QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTOBankDTO `json:"bankDTO,omitempty" xml:"bankDTO,omitempty" type:"Struct"`
}

func (s QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTO) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTO) GoString() string {
	return s.String()
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTO) SetBankDTO(v *QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTOBankDTO) *QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTO {
	s.BankDTO = v
	return s
}

type QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTOBankDTO struct {
	AccountName    *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	BankBranchCode *string `json:"bankBranchCode,omitempty" xml:"bankBranchCode,omitempty"`
	BankBranchName *string `json:"bankBranchName,omitempty" xml:"bankBranchName,omitempty"`
	BankCardNo     *string `json:"bankCardNo,omitempty" xml:"bankCardNo,omitempty"`
	BankCode       *string `json:"bankCode,omitempty" xml:"bankCode,omitempty"`
	BankName       *string `json:"bankName,omitempty" xml:"bankName,omitempty"`
	Type           *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTOBankDTO) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTOBankDTO) GoString() string {
	return s.String()
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTOBankDTO) SetAccountName(v string) *QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTOBankDTO {
	s.AccountName = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTOBankDTO) SetBankBranchCode(v string) *QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTOBankDTO {
	s.BankBranchCode = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTOBankDTO) SetBankBranchName(v string) *QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTOBankDTO {
	s.BankBranchName = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTOBankDTO) SetBankCardNo(v string) *QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTOBankDTO {
	s.BankCardNo = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTOBankDTO) SetBankCode(v string) *QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTOBankDTO {
	s.BankCode = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTOBankDTO) SetBankName(v string) *QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTOBankDTO {
	s.BankName = &v
	return s
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTOBankDTO) SetType(v string) *QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayeeAccountDTOBankDTO {
	s.Type = &v
	return s
}

type QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayerAccountDTO struct {
	EnterpriseAccountCode *string `json:"enterpriseAccountCode,omitempty" xml:"enterpriseAccountCode,omitempty"`
}

func (s QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayerAccountDTO) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayerAccountDTO) GoString() string {
	return s.String()
}

func (s *QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayerAccountDTO) SetEnterpriseAccountCode(v string) *QueryPaymentOrderDetailResponseBodyOrderListSubOrderListPayerAccountDTO {
	s.EnterpriseAccountCode = &v
	return s
}

type QueryPaymentOrderDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPaymentOrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPaymentOrderDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentOrderDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryPaymentOrderDetailResponse) SetHeaders(v map[string]*string) *QueryPaymentOrderDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryPaymentOrderDetailResponse) SetStatusCode(v int32) *QueryPaymentOrderDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPaymentOrderDetailResponse) SetBody(v *QueryPaymentOrderDetailResponseBody) *QueryPaymentOrderDetailResponse {
	s.Body = v
	return s
}

type QueryPaymentRecallFileHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryPaymentRecallFileHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentRecallFileHeaders) GoString() string {
	return s.String()
}

func (s *QueryPaymentRecallFileHeaders) SetCommonHeaders(v map[string]*string) *QueryPaymentRecallFileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryPaymentRecallFileHeaders) SetXAcsDingtalkAccessToken(v string) *QueryPaymentRecallFileHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryPaymentRecallFileRequest struct {
	// example:
	//
	// 504
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryPaymentRecallFileRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentRecallFileRequest) GoString() string {
	return s.String()
}

func (s *QueryPaymentRecallFileRequest) SetUserId(v string) *QueryPaymentRecallFileRequest {
	s.UserId = &v
	return s
}

type QueryPaymentRecallFileResponseBody struct {
	CorpId                *string                                                    `json:"corpId,omitempty" xml:"corpId,omitempty"`
	PaymentRecallFileList []*QueryPaymentRecallFileResponseBodyPaymentRecallFileList `json:"paymentRecallFileList,omitempty" xml:"paymentRecallFileList,omitempty" type:"Repeated"`
}

func (s QueryPaymentRecallFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentRecallFileResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPaymentRecallFileResponseBody) SetCorpId(v string) *QueryPaymentRecallFileResponseBody {
	s.CorpId = &v
	return s
}

func (s *QueryPaymentRecallFileResponseBody) SetPaymentRecallFileList(v []*QueryPaymentRecallFileResponseBodyPaymentRecallFileList) *QueryPaymentRecallFileResponseBody {
	s.PaymentRecallFileList = v
	return s
}

type QueryPaymentRecallFileResponseBodyPaymentRecallFileList struct {
	FileId     *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	FileName   *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileSize   *string `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	FileType   *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	OrderNo    *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	PreviewUrl *string `json:"previewUrl,omitempty" xml:"previewUrl,omitempty"`
	SpaceId    *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s QueryPaymentRecallFileResponseBodyPaymentRecallFileList) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentRecallFileResponseBodyPaymentRecallFileList) GoString() string {
	return s.String()
}

func (s *QueryPaymentRecallFileResponseBodyPaymentRecallFileList) SetFileId(v string) *QueryPaymentRecallFileResponseBodyPaymentRecallFileList {
	s.FileId = &v
	return s
}

func (s *QueryPaymentRecallFileResponseBodyPaymentRecallFileList) SetFileName(v string) *QueryPaymentRecallFileResponseBodyPaymentRecallFileList {
	s.FileName = &v
	return s
}

func (s *QueryPaymentRecallFileResponseBodyPaymentRecallFileList) SetFileSize(v string) *QueryPaymentRecallFileResponseBodyPaymentRecallFileList {
	s.FileSize = &v
	return s
}

func (s *QueryPaymentRecallFileResponseBodyPaymentRecallFileList) SetFileType(v string) *QueryPaymentRecallFileResponseBodyPaymentRecallFileList {
	s.FileType = &v
	return s
}

func (s *QueryPaymentRecallFileResponseBodyPaymentRecallFileList) SetInstanceId(v string) *QueryPaymentRecallFileResponseBodyPaymentRecallFileList {
	s.InstanceId = &v
	return s
}

func (s *QueryPaymentRecallFileResponseBodyPaymentRecallFileList) SetOrderNo(v string) *QueryPaymentRecallFileResponseBodyPaymentRecallFileList {
	s.OrderNo = &v
	return s
}

func (s *QueryPaymentRecallFileResponseBodyPaymentRecallFileList) SetPreviewUrl(v string) *QueryPaymentRecallFileResponseBodyPaymentRecallFileList {
	s.PreviewUrl = &v
	return s
}

func (s *QueryPaymentRecallFileResponseBodyPaymentRecallFileList) SetSpaceId(v string) *QueryPaymentRecallFileResponseBodyPaymentRecallFileList {
	s.SpaceId = &v
	return s
}

type QueryPaymentRecallFileResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPaymentRecallFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPaymentRecallFileResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentRecallFileResponse) GoString() string {
	return s.String()
}

func (s *QueryPaymentRecallFileResponse) SetHeaders(v map[string]*string) *QueryPaymentRecallFileResponse {
	s.Headers = v
	return s
}

func (s *QueryPaymentRecallFileResponse) SetStatusCode(v int32) *QueryPaymentRecallFileResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPaymentRecallFileResponse) SetBody(v *QueryPaymentRecallFileResponseBody) *QueryPaymentRecallFileResponse {
	s.Body = v
	return s
}

type QueryPaymentStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryPaymentStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentStatusHeaders) GoString() string {
	return s.String()
}

func (s *QueryPaymentStatusHeaders) SetCommonHeaders(v map[string]*string) *QueryPaymentStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryPaymentStatusHeaders) SetXAcsDingtalkAccessToken(v string) *QueryPaymentStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryPaymentStatusRequest struct {
	// example:
	//
	// ABC
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 20241102231
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	// example:
	//
	// 504323
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryPaymentStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryPaymentStatusRequest) SetInstanceId(v string) *QueryPaymentStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryPaymentStatusRequest) SetOrderNo(v string) *QueryPaymentStatusRequest {
	s.OrderNo = &v
	return s
}

func (s *QueryPaymentStatusRequest) SetUserId(v string) *QueryPaymentStatusRequest {
	s.UserId = &v
	return s
}

type QueryPaymentStatusResponseBody struct {
	// example:
	//
	// dingXXX
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// 收款账户错误
	FailReason *string `json:"failReason,omitempty" xml:"failReason,omitempty"`
	// example:
	//
	// ABC
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 20241112
	OrderNo          *string                                         `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	PayeeAccountInfo *QueryPaymentStatusResponseBodyPayeeAccountInfo `json:"payeeAccountInfo,omitempty" xml:"payeeAccountInfo,omitempty" type:"Struct"`
	PayerAccountInfo *QueryPaymentStatusResponseBodyPayerAccountInfo `json:"payerAccountInfo,omitempty" xml:"payerAccountInfo,omitempty" type:"Struct"`
	// example:
	//
	// SUCCESS
	PaymentStatus *string `json:"paymentStatus,omitempty" xml:"paymentStatus,omitempty"`
	// example:
	//
	// 2024-11-11 00:00:00
	PaymentTime *string `json:"paymentTime,omitempty" xml:"paymentTime,omitempty"`
	// example:
	//
	// ABC
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// 报销
	Usage *string `json:"usage,omitempty" xml:"usage,omitempty"`
	// example:
	//
	// 504
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryPaymentStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPaymentStatusResponseBody) SetCorpId(v string) *QueryPaymentStatusResponseBody {
	s.CorpId = &v
	return s
}

func (s *QueryPaymentStatusResponseBody) SetFailReason(v string) *QueryPaymentStatusResponseBody {
	s.FailReason = &v
	return s
}

func (s *QueryPaymentStatusResponseBody) SetInstanceId(v string) *QueryPaymentStatusResponseBody {
	s.InstanceId = &v
	return s
}

func (s *QueryPaymentStatusResponseBody) SetOrderNo(v string) *QueryPaymentStatusResponseBody {
	s.OrderNo = &v
	return s
}

func (s *QueryPaymentStatusResponseBody) SetPayeeAccountInfo(v *QueryPaymentStatusResponseBodyPayeeAccountInfo) *QueryPaymentStatusResponseBody {
	s.PayeeAccountInfo = v
	return s
}

func (s *QueryPaymentStatusResponseBody) SetPayerAccountInfo(v *QueryPaymentStatusResponseBodyPayerAccountInfo) *QueryPaymentStatusResponseBody {
	s.PayerAccountInfo = v
	return s
}

func (s *QueryPaymentStatusResponseBody) SetPaymentStatus(v string) *QueryPaymentStatusResponseBody {
	s.PaymentStatus = &v
	return s
}

func (s *QueryPaymentStatusResponseBody) SetPaymentTime(v string) *QueryPaymentStatusResponseBody {
	s.PaymentTime = &v
	return s
}

func (s *QueryPaymentStatusResponseBody) SetRemark(v string) *QueryPaymentStatusResponseBody {
	s.Remark = &v
	return s
}

func (s *QueryPaymentStatusResponseBody) SetUsage(v string) *QueryPaymentStatusResponseBody {
	s.Usage = &v
	return s
}

func (s *QueryPaymentStatusResponseBody) SetUserId(v string) *QueryPaymentStatusResponseBody {
	s.UserId = &v
	return s
}

type QueryPaymentStatusResponseBodyPayeeAccountInfo struct {
	BankOpenDTO *QueryPaymentStatusResponseBodyPayeeAccountInfoBankOpenDTO `json:"bankOpenDTO,omitempty" xml:"bankOpenDTO,omitempty" type:"Struct"`
}

func (s QueryPaymentStatusResponseBodyPayeeAccountInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentStatusResponseBodyPayeeAccountInfo) GoString() string {
	return s.String()
}

func (s *QueryPaymentStatusResponseBodyPayeeAccountInfo) SetBankOpenDTO(v *QueryPaymentStatusResponseBodyPayeeAccountInfoBankOpenDTO) *QueryPaymentStatusResponseBodyPayeeAccountInfo {
	s.BankOpenDTO = v
	return s
}

type QueryPaymentStatusResponseBodyPayeeAccountInfoBankOpenDTO struct {
	BankBranchName *string `json:"bankBranchName,omitempty" xml:"bankBranchName,omitempty"`
	BankCardNo     *string `json:"bankCardNo,omitempty" xml:"bankCardNo,omitempty"`
	BankName       *string `json:"bankName,omitempty" xml:"bankName,omitempty"`
}

func (s QueryPaymentStatusResponseBodyPayeeAccountInfoBankOpenDTO) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentStatusResponseBodyPayeeAccountInfoBankOpenDTO) GoString() string {
	return s.String()
}

func (s *QueryPaymentStatusResponseBodyPayeeAccountInfoBankOpenDTO) SetBankBranchName(v string) *QueryPaymentStatusResponseBodyPayeeAccountInfoBankOpenDTO {
	s.BankBranchName = &v
	return s
}

func (s *QueryPaymentStatusResponseBodyPayeeAccountInfoBankOpenDTO) SetBankCardNo(v string) *QueryPaymentStatusResponseBodyPayeeAccountInfoBankOpenDTO {
	s.BankCardNo = &v
	return s
}

func (s *QueryPaymentStatusResponseBodyPayeeAccountInfoBankOpenDTO) SetBankName(v string) *QueryPaymentStatusResponseBodyPayeeAccountInfoBankOpenDTO {
	s.BankName = &v
	return s
}

type QueryPaymentStatusResponseBodyPayerAccountInfo struct {
	AccountType           *string                                                    `json:"accountType,omitempty" xml:"accountType,omitempty"`
	BankOpenDTO           *QueryPaymentStatusResponseBodyPayerAccountInfoBankOpenDTO `json:"bankOpenDTO,omitempty" xml:"bankOpenDTO,omitempty" type:"Struct"`
	EnterpriseAccountCode *string                                                    `json:"enterpriseAccountCode,omitempty" xml:"enterpriseAccountCode,omitempty"`
}

func (s QueryPaymentStatusResponseBodyPayerAccountInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentStatusResponseBodyPayerAccountInfo) GoString() string {
	return s.String()
}

func (s *QueryPaymentStatusResponseBodyPayerAccountInfo) SetAccountType(v string) *QueryPaymentStatusResponseBodyPayerAccountInfo {
	s.AccountType = &v
	return s
}

func (s *QueryPaymentStatusResponseBodyPayerAccountInfo) SetBankOpenDTO(v *QueryPaymentStatusResponseBodyPayerAccountInfoBankOpenDTO) *QueryPaymentStatusResponseBodyPayerAccountInfo {
	s.BankOpenDTO = v
	return s
}

func (s *QueryPaymentStatusResponseBodyPayerAccountInfo) SetEnterpriseAccountCode(v string) *QueryPaymentStatusResponseBodyPayerAccountInfo {
	s.EnterpriseAccountCode = &v
	return s
}

type QueryPaymentStatusResponseBodyPayerAccountInfoBankOpenDTO struct {
	BankBranchName *string `json:"bankBranchName,omitempty" xml:"bankBranchName,omitempty"`
	BankCardNo     *string `json:"bankCardNo,omitempty" xml:"bankCardNo,omitempty"`
	BankName       *string `json:"bankName,omitempty" xml:"bankName,omitempty"`
}

func (s QueryPaymentStatusResponseBodyPayerAccountInfoBankOpenDTO) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentStatusResponseBodyPayerAccountInfoBankOpenDTO) GoString() string {
	return s.String()
}

func (s *QueryPaymentStatusResponseBodyPayerAccountInfoBankOpenDTO) SetBankBranchName(v string) *QueryPaymentStatusResponseBodyPayerAccountInfoBankOpenDTO {
	s.BankBranchName = &v
	return s
}

func (s *QueryPaymentStatusResponseBodyPayerAccountInfoBankOpenDTO) SetBankCardNo(v string) *QueryPaymentStatusResponseBodyPayerAccountInfoBankOpenDTO {
	s.BankCardNo = &v
	return s
}

func (s *QueryPaymentStatusResponseBodyPayerAccountInfoBankOpenDTO) SetBankName(v string) *QueryPaymentStatusResponseBodyPayerAccountInfoBankOpenDTO {
	s.BankName = &v
	return s
}

type QueryPaymentStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPaymentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPaymentStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPaymentStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryPaymentStatusResponse) SetHeaders(v map[string]*string) *QueryPaymentStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryPaymentStatusResponse) SetStatusCode(v int32) *QueryPaymentStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPaymentStatusResponse) SetBody(v *QueryPaymentStatusResponseBody) *QueryPaymentStatusResponse {
	s.Body = v
	return s
}

type QueryProductByPageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryProductByPageHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryProductByPageHeaders) GoString() string {
	return s.String()
}

func (s *QueryProductByPageHeaders) SetCommonHeaders(v map[string]*string) *QueryProductByPageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryProductByPageHeaders) SetXAcsDingtalkAccessToken(v string) *QueryProductByPageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryProductByPageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryProductByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryProductByPageRequest) GoString() string {
	return s.String()
}

func (s *QueryProductByPageRequest) SetPageNumber(v int64) *QueryProductByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryProductByPageRequest) SetPageSize(v int64) *QueryProductByPageRequest {
	s.PageSize = &v
	return s
}

type QueryProductByPageResponseBody struct {
	HasMore *bool                                 `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List    []*QueryProductByPageResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s QueryProductByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryProductByPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryProductByPageResponseBody) SetHasMore(v bool) *QueryProductByPageResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryProductByPageResponseBody) SetList(v []*QueryProductByPageResponseBodyList) *QueryProductByPageResponseBody {
	s.List = v
	return s
}

type QueryProductByPageResponseBodyList struct {
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	CreateTime     *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	Information    *string `json:"information,omitempty" xml:"information,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	Specification  *string `json:"specification,omitempty" xml:"specification,omitempty"`
	Status         *string `json:"status,omitempty" xml:"status,omitempty"`
	Unit           *string `json:"unit,omitempty" xml:"unit,omitempty"`
	UserDefineCode *string `json:"userDefineCode,omitempty" xml:"userDefineCode,omitempty"`
}

func (s QueryProductByPageResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryProductByPageResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryProductByPageResponseBodyList) SetCode(v string) *QueryProductByPageResponseBodyList {
	s.Code = &v
	return s
}

func (s *QueryProductByPageResponseBodyList) SetCreateTime(v int64) *QueryProductByPageResponseBodyList {
	s.CreateTime = &v
	return s
}

func (s *QueryProductByPageResponseBodyList) SetDescription(v string) *QueryProductByPageResponseBodyList {
	s.Description = &v
	return s
}

func (s *QueryProductByPageResponseBodyList) SetInformation(v string) *QueryProductByPageResponseBodyList {
	s.Information = &v
	return s
}

func (s *QueryProductByPageResponseBodyList) SetName(v string) *QueryProductByPageResponseBodyList {
	s.Name = &v
	return s
}

func (s *QueryProductByPageResponseBodyList) SetSpecification(v string) *QueryProductByPageResponseBodyList {
	s.Specification = &v
	return s
}

func (s *QueryProductByPageResponseBodyList) SetStatus(v string) *QueryProductByPageResponseBodyList {
	s.Status = &v
	return s
}

func (s *QueryProductByPageResponseBodyList) SetUnit(v string) *QueryProductByPageResponseBodyList {
	s.Unit = &v
	return s
}

func (s *QueryProductByPageResponseBodyList) SetUserDefineCode(v string) *QueryProductByPageResponseBodyList {
	s.UserDefineCode = &v
	return s
}

type QueryProductByPageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryProductByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryProductByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryProductByPageResponse) GoString() string {
	return s.String()
}

func (s *QueryProductByPageResponse) SetHeaders(v map[string]*string) *QueryProductByPageResponse {
	s.Headers = v
	return s
}

func (s *QueryProductByPageResponse) SetStatusCode(v int32) *QueryProductByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryProductByPageResponse) SetBody(v *QueryProductByPageResponseBody) *QueryProductByPageResponse {
	s.Body = v
	return s
}

type QueryProjectByPageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryProjectByPageHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectByPageHeaders) GoString() string {
	return s.String()
}

func (s *QueryProjectByPageHeaders) SetCommonHeaders(v map[string]*string) *QueryProjectByPageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryProjectByPageHeaders) SetXAcsDingtalkAccessToken(v string) *QueryProjectByPageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryProjectByPageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryProjectByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectByPageRequest) GoString() string {
	return s.String()
}

func (s *QueryProjectByPageRequest) SetPageNumber(v int64) *QueryProjectByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryProjectByPageRequest) SetPageSize(v int64) *QueryProjectByPageRequest {
	s.PageSize = &v
	return s
}

type QueryProjectByPageResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// This parameter is required.
	List []*QueryProjectByPageResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s QueryProjectByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectByPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryProjectByPageResponseBody) SetHasMore(v bool) *QueryProjectByPageResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryProjectByPageResponseBody) SetList(v []*QueryProjectByPageResponseBodyList) *QueryProjectByPageResponseBody {
	s.List = v
	return s
}

type QueryProjectByPageResponseBodyList struct {
	Caode *string `json:"caode,omitempty" xml:"caode,omitempty"`
	Code  *string `json:"code,omitempty" xml:"code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1631524595555
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aaaa
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// 外派项目
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	ParentCode  *string `json:"parentCode,omitempty" xml:"parentCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROJ-xxx
	ProjectCode *string `json:"projectCode,omitempty" xml:"projectCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 外派项目
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// valid
	Status         *string `json:"status,omitempty" xml:"status,omitempty"`
	UserDefineCode *string `json:"userDefineCode,omitempty" xml:"userDefineCode,omitempty"`
}

func (s QueryProjectByPageResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectByPageResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryProjectByPageResponseBodyList) SetCaode(v string) *QueryProjectByPageResponseBodyList {
	s.Caode = &v
	return s
}

func (s *QueryProjectByPageResponseBodyList) SetCode(v string) *QueryProjectByPageResponseBodyList {
	s.Code = &v
	return s
}

func (s *QueryProjectByPageResponseBodyList) SetCreateTime(v int64) *QueryProjectByPageResponseBodyList {
	s.CreateTime = &v
	return s
}

func (s *QueryProjectByPageResponseBodyList) SetCreator(v string) *QueryProjectByPageResponseBodyList {
	s.Creator = &v
	return s
}

func (s *QueryProjectByPageResponseBodyList) SetDescription(v string) *QueryProjectByPageResponseBodyList {
	s.Description = &v
	return s
}

func (s *QueryProjectByPageResponseBodyList) SetName(v string) *QueryProjectByPageResponseBodyList {
	s.Name = &v
	return s
}

func (s *QueryProjectByPageResponseBodyList) SetParentCode(v string) *QueryProjectByPageResponseBodyList {
	s.ParentCode = &v
	return s
}

func (s *QueryProjectByPageResponseBodyList) SetProjectCode(v string) *QueryProjectByPageResponseBodyList {
	s.ProjectCode = &v
	return s
}

func (s *QueryProjectByPageResponseBodyList) SetProjectName(v string) *QueryProjectByPageResponseBodyList {
	s.ProjectName = &v
	return s
}

func (s *QueryProjectByPageResponseBodyList) SetStatus(v string) *QueryProjectByPageResponseBodyList {
	s.Status = &v
	return s
}

func (s *QueryProjectByPageResponseBodyList) SetUserDefineCode(v string) *QueryProjectByPageResponseBodyList {
	s.UserDefineCode = &v
	return s
}

type QueryProjectByPageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryProjectByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryProjectByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectByPageResponse) GoString() string {
	return s.String()
}

func (s *QueryProjectByPageResponse) SetHeaders(v map[string]*string) *QueryProjectByPageResponse {
	s.Headers = v
	return s
}

func (s *QueryProjectByPageResponse) SetStatusCode(v int32) *QueryProjectByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryProjectByPageResponse) SetBody(v *QueryProjectByPageResponseBody) *QueryProjectByPageResponse {
	s.Body = v
	return s
}

type QueryReceiptForInvoiceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryReceiptForInvoiceHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiptForInvoiceHeaders) GoString() string {
	return s.String()
}

func (s *QueryReceiptForInvoiceHeaders) SetCommonHeaders(v map[string]*string) *QueryReceiptForInvoiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryReceiptForInvoiceHeaders) SetXAcsDingtalkAccessToken(v string) *QueryReceiptForInvoiceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryReceiptForInvoiceRequest struct {
	// example:
	//
	// abc
	AccountantBookId *string   `json:"accountantBookId,omitempty" xml:"accountantBookId,omitempty"`
	ApplyStatusList  []*string `json:"applyStatusList,omitempty" xml:"applyStatusList,omitempty" type:"Repeated"`
	BizStatusList    []*string `json:"bizStatusList,omitempty" xml:"bizStatusList,omitempty" type:"Repeated"`
	// example:
	//
	// COM_DEFAULT
	CompanyCode       *string            `json:"companyCode,omitempty" xml:"companyCode,omitempty"`
	EndTime           *int64             `json:"endTime,omitempty" xml:"endTime,omitempty"`
	PageNumber        *int64             `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize          *int64             `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ReceiptStatusList []*string          `json:"receiptStatusList,omitempty" xml:"receiptStatusList,omitempty" type:"Repeated"`
	SearchParams      map[string]*string `json:"searchParams,omitempty" xml:"searchParams,omitempty"`
	StartTime         *int64             `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Title             *string            `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryReceiptForInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiptForInvoiceRequest) GoString() string {
	return s.String()
}

func (s *QueryReceiptForInvoiceRequest) SetAccountantBookId(v string) *QueryReceiptForInvoiceRequest {
	s.AccountantBookId = &v
	return s
}

func (s *QueryReceiptForInvoiceRequest) SetApplyStatusList(v []*string) *QueryReceiptForInvoiceRequest {
	s.ApplyStatusList = v
	return s
}

func (s *QueryReceiptForInvoiceRequest) SetBizStatusList(v []*string) *QueryReceiptForInvoiceRequest {
	s.BizStatusList = v
	return s
}

func (s *QueryReceiptForInvoiceRequest) SetCompanyCode(v string) *QueryReceiptForInvoiceRequest {
	s.CompanyCode = &v
	return s
}

func (s *QueryReceiptForInvoiceRequest) SetEndTime(v int64) *QueryReceiptForInvoiceRequest {
	s.EndTime = &v
	return s
}

func (s *QueryReceiptForInvoiceRequest) SetPageNumber(v int64) *QueryReceiptForInvoiceRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryReceiptForInvoiceRequest) SetPageSize(v int64) *QueryReceiptForInvoiceRequest {
	s.PageSize = &v
	return s
}

func (s *QueryReceiptForInvoiceRequest) SetReceiptStatusList(v []*string) *QueryReceiptForInvoiceRequest {
	s.ReceiptStatusList = v
	return s
}

func (s *QueryReceiptForInvoiceRequest) SetSearchParams(v map[string]*string) *QueryReceiptForInvoiceRequest {
	s.SearchParams = v
	return s
}

func (s *QueryReceiptForInvoiceRequest) SetStartTime(v int64) *QueryReceiptForInvoiceRequest {
	s.StartTime = &v
	return s
}

func (s *QueryReceiptForInvoiceRequest) SetTitle(v string) *QueryReceiptForInvoiceRequest {
	s.Title = &v
	return s
}

type QueryReceiptForInvoiceResponseBody struct {
	// example:
	//
	// true
	HasMore *string                                   `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List    []*QueryReceiptForInvoiceResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// example:
	//
	// 500
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryReceiptForInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiptForInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryReceiptForInvoiceResponseBody) SetHasMore(v string) *QueryReceiptForInvoiceResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBody) SetList(v []*QueryReceiptForInvoiceResponseBodyList) *QueryReceiptForInvoiceResponseBody {
	s.List = v
	return s
}

func (s *QueryReceiptForInvoiceResponseBody) SetTotalCount(v int64) *QueryReceiptForInvoiceResponseBody {
	s.TotalCount = &v
	return s
}

type QueryReceiptForInvoiceResponseBodyList struct {
	// example:
	//
	// abc
	AccountantBookId *string `json:"accountantBookId,omitempty" xml:"accountantBookId,omitempty"`
	// example:
	//
	// 5000
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// example:
	//
	// applied
	ApplyStatus *string `json:"applyStatus,omitempty" xml:"applyStatus,omitempty"`
	// example:
	//
	// invoicing
	BizStatus *string `json:"bizStatus,omitempty" xml:"bizStatus,omitempty"`
	// example:
	//
	// 123
	BusinessId *string `json:"businessId,omitempty" xml:"businessId,omitempty"`
	// example:
	//
	// COM_DEFAULT
	CompanyCode *string                                         `json:"companyCode,omitempty" xml:"companyCode,omitempty"`
	CreateTime  *string                                         `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Creator     *QueryReceiptForInvoiceResponseBodyListCreator  `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	Customer    *QueryReceiptForInvoiceResponseBodyListCustomer `json:"customer,omitempty" xml:"customer,omitempty" type:"Struct"`
	// example:
	//
	// www.abc.com
	DrawerEmail *string `json:"drawerEmail,omitempty" xml:"drawerEmail,omitempty"`
	// example:
	//
	// 12345678901
	DrawerTelephone *string `json:"drawerTelephone,omitempty" xml:"drawerTelephone,omitempty"`
	// example:
	//
	// 增值税发票
	InvoiceType *string `json:"invoiceType,omitempty" xml:"invoiceType,omitempty"`
	JumpUrl     *string `json:"jumpUrl,omitempty" xml:"jumpUrl,omitempty"`
	// example:
	//
	// EM-xxxxx
	ModelId         *string                                                  `json:"modelId,omitempty" xml:"modelId,omitempty"`
	ProductInfoList []*QueryReceiptForInvoiceResponseBodyListProductInfoList `json:"productInfoList,omitempty" xml:"productInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// abc
	PurchaserAccount *string `json:"purchaserAccount,omitempty" xml:"purchaserAccount,omitempty"`
	// example:
	//
	// 杭州市
	PurchaserAddress *string `json:"purchaserAddress,omitempty" xml:"purchaserAddress,omitempty"`
	// example:
	//
	// 建设银行
	PurchaserBankName *string `json:"purchaserBankName,omitempty" xml:"purchaserBankName,omitempty"`
	// example:
	//
	// 钉有限公司
	PurchaserName *string `json:"purchaserName,omitempty" xml:"purchaserName,omitempty"`
	// example:
	//
	// 123456
	PurchaserTaxNo *string `json:"purchaserTaxNo,omitempty" xml:"purchaserTaxNo,omitempty"`
	// example:
	//
	// 13333333333
	PurchaserTel *string `json:"purchaserTel,omitempty" xml:"purchaserTel,omitempty"`
	// example:
	//
	// abc
	ReceiptId *string `json:"receiptId,omitempty" xml:"receiptId,omitempty"`
	// example:
	//
	// 16000000
	RecordTime *string `json:"recordTime,omitempty" xml:"recordTime,omitempty"`
	// example:
	//
	// 备注信息
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// approval
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// agree
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 张三提交的开票申请单
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryReceiptForInvoiceResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiptForInvoiceResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetAccountantBookId(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.AccountantBookId = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetAmount(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.Amount = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetApplyStatus(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.ApplyStatus = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetBizStatus(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.BizStatus = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetBusinessId(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.BusinessId = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetCompanyCode(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.CompanyCode = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetCreateTime(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.CreateTime = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetCreator(v *QueryReceiptForInvoiceResponseBodyListCreator) *QueryReceiptForInvoiceResponseBodyList {
	s.Creator = v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetCustomer(v *QueryReceiptForInvoiceResponseBodyListCustomer) *QueryReceiptForInvoiceResponseBodyList {
	s.Customer = v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetDrawerEmail(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.DrawerEmail = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetDrawerTelephone(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.DrawerTelephone = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetInvoiceType(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.InvoiceType = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetJumpUrl(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.JumpUrl = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetModelId(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.ModelId = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetProductInfoList(v []*QueryReceiptForInvoiceResponseBodyListProductInfoList) *QueryReceiptForInvoiceResponseBodyList {
	s.ProductInfoList = v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetPurchaserAccount(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.PurchaserAccount = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetPurchaserAddress(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.PurchaserAddress = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetPurchaserBankName(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.PurchaserBankName = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetPurchaserName(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.PurchaserName = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetPurchaserTaxNo(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.PurchaserTaxNo = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetPurchaserTel(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.PurchaserTel = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetReceiptId(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.ReceiptId = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetRecordTime(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.RecordTime = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetRemark(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.Remark = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetSource(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.Source = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetStatus(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.Status = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetTitle(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.Title = &v
	return s
}

type QueryReceiptForInvoiceResponseBodyListCreator struct {
	// example:
	//
	// https://xxxx
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// 测试名字
	Nick *string `json:"nick,omitempty" xml:"nick,omitempty"`
	// example:
	//
	// 1231
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryReceiptForInvoiceResponseBodyListCreator) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiptForInvoiceResponseBodyListCreator) GoString() string {
	return s.String()
}

func (s *QueryReceiptForInvoiceResponseBodyListCreator) SetAvatarUrl(v string) *QueryReceiptForInvoiceResponseBodyListCreator {
	s.AvatarUrl = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyListCreator) SetNick(v string) *QueryReceiptForInvoiceResponseBodyListCreator {
	s.Nick = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyListCreator) SetUserId(v string) *QueryReceiptForInvoiceResponseBodyListCreator {
	s.UserId = &v
	return s
}

type QueryReceiptForInvoiceResponseBodyListCustomer struct {
	// example:
	//
	// CUS_xxxxx
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 李四
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryReceiptForInvoiceResponseBodyListCustomer) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiptForInvoiceResponseBodyListCustomer) GoString() string {
	return s.String()
}

func (s *QueryReceiptForInvoiceResponseBodyListCustomer) SetCode(v string) *QueryReceiptForInvoiceResponseBodyListCustomer {
	s.Code = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyListCustomer) SetName(v string) *QueryReceiptForInvoiceResponseBodyListCustomer {
	s.Name = &v
	return s
}

type QueryReceiptForInvoiceResponseBodyListProductInfoList struct {
	// example:
	//
	// 12.3
	AmountWithTax *string `json:"amountWithTax,omitempty" xml:"amountWithTax,omitempty"`
	// example:
	//
	// 100
	AmountWithoutTax *string `json:"amountWithoutTax,omitempty" xml:"amountWithoutTax,omitempty"`
	// example:
	//
	// 10
	DiscountAmount *string `json:"discountAmount,omitempty" xml:"discountAmount,omitempty"`
	// example:
	//
	// 鱼
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 2
	Quantity *string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// example:
	//
	// 大型
	Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
	// example:
	//
	// XXX
	TaxClassificationCode *string `json:"taxClassificationCode,omitempty" xml:"taxClassificationCode,omitempty"`
	// example:
	//
	// 0.3
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// example:
	//
	// 千克
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// example:
	//
	// 12.3
	UnitPriceWithTax *string `json:"unitPriceWithTax,omitempty" xml:"unitPriceWithTax,omitempty"`
	// example:
	//
	// 100
	UnitPriceWithoutTax *string `json:"unitPriceWithoutTax,omitempty" xml:"unitPriceWithoutTax,omitempty"`
	WithTax             *bool   `json:"withTax,omitempty" xml:"withTax,omitempty"`
}

func (s QueryReceiptForInvoiceResponseBodyListProductInfoList) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiptForInvoiceResponseBodyListProductInfoList) GoString() string {
	return s.String()
}

func (s *QueryReceiptForInvoiceResponseBodyListProductInfoList) SetAmountWithTax(v string) *QueryReceiptForInvoiceResponseBodyListProductInfoList {
	s.AmountWithTax = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyListProductInfoList) SetAmountWithoutTax(v string) *QueryReceiptForInvoiceResponseBodyListProductInfoList {
	s.AmountWithoutTax = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyListProductInfoList) SetDiscountAmount(v string) *QueryReceiptForInvoiceResponseBodyListProductInfoList {
	s.DiscountAmount = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyListProductInfoList) SetName(v string) *QueryReceiptForInvoiceResponseBodyListProductInfoList {
	s.Name = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyListProductInfoList) SetQuantity(v string) *QueryReceiptForInvoiceResponseBodyListProductInfoList {
	s.Quantity = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyListProductInfoList) SetSpecification(v string) *QueryReceiptForInvoiceResponseBodyListProductInfoList {
	s.Specification = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyListProductInfoList) SetTaxClassificationCode(v string) *QueryReceiptForInvoiceResponseBodyListProductInfoList {
	s.TaxClassificationCode = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyListProductInfoList) SetTaxRate(v string) *QueryReceiptForInvoiceResponseBodyListProductInfoList {
	s.TaxRate = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyListProductInfoList) SetUnit(v string) *QueryReceiptForInvoiceResponseBodyListProductInfoList {
	s.Unit = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyListProductInfoList) SetUnitPriceWithTax(v string) *QueryReceiptForInvoiceResponseBodyListProductInfoList {
	s.UnitPriceWithTax = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyListProductInfoList) SetUnitPriceWithoutTax(v string) *QueryReceiptForInvoiceResponseBodyListProductInfoList {
	s.UnitPriceWithoutTax = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyListProductInfoList) SetWithTax(v bool) *QueryReceiptForInvoiceResponseBodyListProductInfoList {
	s.WithTax = &v
	return s
}

type QueryReceiptForInvoiceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryReceiptForInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryReceiptForInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiptForInvoiceResponse) GoString() string {
	return s.String()
}

func (s *QueryReceiptForInvoiceResponse) SetHeaders(v map[string]*string) *QueryReceiptForInvoiceResponse {
	s.Headers = v
	return s
}

func (s *QueryReceiptForInvoiceResponse) SetStatusCode(v int32) *QueryReceiptForInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryReceiptForInvoiceResponse) SetBody(v *QueryReceiptForInvoiceResponseBody) *QueryReceiptForInvoiceResponse {
	s.Body = v
	return s
}

type QuerySupplierByPageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QuerySupplierByPageHeaders) String() string {
	return tea.Prettify(s)
}

func (s QuerySupplierByPageHeaders) GoString() string {
	return s.String()
}

func (s *QuerySupplierByPageHeaders) SetCommonHeaders(v map[string]*string) *QuerySupplierByPageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QuerySupplierByPageHeaders) SetXAcsDingtalkAccessToken(v string) *QuerySupplierByPageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QuerySupplierByPageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QuerySupplierByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySupplierByPageRequest) GoString() string {
	return s.String()
}

func (s *QuerySupplierByPageRequest) SetPageNumber(v int64) *QuerySupplierByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *QuerySupplierByPageRequest) SetPageSize(v int64) *QuerySupplierByPageRequest {
	s.PageSize = &v
	return s
}

type QuerySupplierByPageResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// This parameter is required.
	List []*QuerySupplierByPageResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s QuerySupplierByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySupplierByPageResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySupplierByPageResponseBody) SetHasMore(v bool) *QuerySupplierByPageResponseBody {
	s.HasMore = &v
	return s
}

func (s *QuerySupplierByPageResponseBody) SetList(v []*QuerySupplierByPageResponseBodyList) *QuerySupplierByPageResponseBody {
	s.List = v
	return s
}

type QuerySupplierByPageResponseBodyList struct {
	// This parameter is required.
	//
	// example:
	//
	// SUP_XXX
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1634786828686
	CreateTime         *int64                                                   `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CustomFormDataList []*QuerySupplierByPageResponseBodyListCustomFormDataList `json:"customFormDataList,omitempty" xml:"customFormDataList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 原材料供应商
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// XX供应商
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// valid
	Status         *string `json:"status,omitempty" xml:"status,omitempty"`
	UserDefineCode *string `json:"userDefineCode,omitempty" xml:"userDefineCode,omitempty"`
}

func (s QuerySupplierByPageResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QuerySupplierByPageResponseBodyList) GoString() string {
	return s.String()
}

func (s *QuerySupplierByPageResponseBodyList) SetCode(v string) *QuerySupplierByPageResponseBodyList {
	s.Code = &v
	return s
}

func (s *QuerySupplierByPageResponseBodyList) SetCreateTime(v int64) *QuerySupplierByPageResponseBodyList {
	s.CreateTime = &v
	return s
}

func (s *QuerySupplierByPageResponseBodyList) SetCustomFormDataList(v []*QuerySupplierByPageResponseBodyListCustomFormDataList) *QuerySupplierByPageResponseBodyList {
	s.CustomFormDataList = v
	return s
}

func (s *QuerySupplierByPageResponseBodyList) SetDescription(v string) *QuerySupplierByPageResponseBodyList {
	s.Description = &v
	return s
}

func (s *QuerySupplierByPageResponseBodyList) SetName(v string) *QuerySupplierByPageResponseBodyList {
	s.Name = &v
	return s
}

func (s *QuerySupplierByPageResponseBodyList) SetStatus(v string) *QuerySupplierByPageResponseBodyList {
	s.Status = &v
	return s
}

func (s *QuerySupplierByPageResponseBodyList) SetUserDefineCode(v string) *QuerySupplierByPageResponseBodyList {
	s.UserDefineCode = &v
	return s
}

type QuerySupplierByPageResponseBodyListCustomFormDataList struct {
	BizAlias      *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	ExtValue      *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	Id            *string `json:"id,omitempty" xml:"id,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	Value         *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s QuerySupplierByPageResponseBodyListCustomFormDataList) String() string {
	return tea.Prettify(s)
}

func (s QuerySupplierByPageResponseBodyListCustomFormDataList) GoString() string {
	return s.String()
}

func (s *QuerySupplierByPageResponseBodyListCustomFormDataList) SetBizAlias(v string) *QuerySupplierByPageResponseBodyListCustomFormDataList {
	s.BizAlias = &v
	return s
}

func (s *QuerySupplierByPageResponseBodyListCustomFormDataList) SetComponentType(v string) *QuerySupplierByPageResponseBodyListCustomFormDataList {
	s.ComponentType = &v
	return s
}

func (s *QuerySupplierByPageResponseBodyListCustomFormDataList) SetExtValue(v string) *QuerySupplierByPageResponseBodyListCustomFormDataList {
	s.ExtValue = &v
	return s
}

func (s *QuerySupplierByPageResponseBodyListCustomFormDataList) SetId(v string) *QuerySupplierByPageResponseBodyListCustomFormDataList {
	s.Id = &v
	return s
}

func (s *QuerySupplierByPageResponseBodyListCustomFormDataList) SetName(v string) *QuerySupplierByPageResponseBodyListCustomFormDataList {
	s.Name = &v
	return s
}

func (s *QuerySupplierByPageResponseBodyListCustomFormDataList) SetValue(v string) *QuerySupplierByPageResponseBodyListCustomFormDataList {
	s.Value = &v
	return s
}

type QuerySupplierByPageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySupplierByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySupplierByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySupplierByPageResponse) GoString() string {
	return s.String()
}

func (s *QuerySupplierByPageResponse) SetHeaders(v map[string]*string) *QuerySupplierByPageResponse {
	s.Headers = v
	return s
}

func (s *QuerySupplierByPageResponse) SetStatusCode(v int32) *QuerySupplierByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySupplierByPageResponse) SetBody(v *QuerySupplierByPageResponseBody) *QuerySupplierByPageResponse {
	s.Body = v
	return s
}

type QueryUseNewInvoiceAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUseNewInvoiceAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUseNewInvoiceAppHeaders) GoString() string {
	return s.String()
}

func (s *QueryUseNewInvoiceAppHeaders) SetCommonHeaders(v map[string]*string) *QueryUseNewInvoiceAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUseNewInvoiceAppHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUseNewInvoiceAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUseNewInvoiceAppResponseBody struct {
	// This parameter is required.
	UseNew *bool `json:"useNew,omitempty" xml:"useNew,omitempty"`
}

func (s QueryUseNewInvoiceAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUseNewInvoiceAppResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUseNewInvoiceAppResponseBody) SetUseNew(v bool) *QueryUseNewInvoiceAppResponseBody {
	s.UseNew = &v
	return s
}

type QueryUseNewInvoiceAppResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUseNewInvoiceAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUseNewInvoiceAppResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUseNewInvoiceAppResponse) GoString() string {
	return s.String()
}

func (s *QueryUseNewInvoiceAppResponse) SetHeaders(v map[string]*string) *QueryUseNewInvoiceAppResponse {
	s.Headers = v
	return s
}

func (s *QueryUseNewInvoiceAppResponse) SetStatusCode(v int32) *QueryUseNewInvoiceAppResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUseNewInvoiceAppResponse) SetBody(v *QueryUseNewInvoiceAppResponseBody) *QueryUseNewInvoiceAppResponse {
	s.Body = v
	return s
}

type QueryUserRoleListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserRoleListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserRoleListHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserRoleListHeaders) SetCommonHeaders(v map[string]*string) *QueryUserRoleListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserRoleListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserRoleListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserRoleListRequest struct {
	// example:
	//
	// COM_DEFAULT
	CompanyCode *string `json:"companyCode,omitempty" xml:"companyCode,omitempty"`
	// example:
	//
	// 12312231231
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryUserRoleListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserRoleListRequest) GoString() string {
	return s.String()
}

func (s *QueryUserRoleListRequest) SetCompanyCode(v string) *QueryUserRoleListRequest {
	s.CompanyCode = &v
	return s
}

func (s *QueryUserRoleListRequest) SetUserId(v string) *QueryUserRoleListRequest {
	s.UserId = &v
	return s
}

type QueryUserRoleListResponseBody struct {
	// example:
	//
	// COM_DEFAULT
	CompanyCode            *string                                                `json:"companyCode,omitempty" xml:"companyCode,omitempty"`
	FinanceEmpDeptOpenList []*QueryUserRoleListResponseBodyFinanceEmpDeptOpenList `json:"financeEmpDeptOpenList,omitempty" xml:"financeEmpDeptOpenList,omitempty" type:"Repeated"`
	RoleVOList             []*QueryUserRoleListResponseBodyRoleVOList             `json:"roleVOList,omitempty" xml:"roleVOList,omitempty" type:"Repeated"`
}

func (s QueryUserRoleListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserRoleListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserRoleListResponseBody) SetCompanyCode(v string) *QueryUserRoleListResponseBody {
	s.CompanyCode = &v
	return s
}

func (s *QueryUserRoleListResponseBody) SetFinanceEmpDeptOpenList(v []*QueryUserRoleListResponseBodyFinanceEmpDeptOpenList) *QueryUserRoleListResponseBody {
	s.FinanceEmpDeptOpenList = v
	return s
}

func (s *QueryUserRoleListResponseBody) SetRoleVOList(v []*QueryUserRoleListResponseBodyRoleVOList) *QueryUserRoleListResponseBody {
	s.RoleVOList = v
	return s
}

type QueryUserRoleListResponseBodyFinanceEmpDeptOpenList struct {
	CascadeDeptId *string `json:"cascadeDeptId,omitempty" xml:"cascadeDeptId,omitempty"`
	DeptId        *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	SuperDeptId   *int64  `json:"superDeptId,omitempty" xml:"superDeptId,omitempty"`
}

func (s QueryUserRoleListResponseBodyFinanceEmpDeptOpenList) String() string {
	return tea.Prettify(s)
}

func (s QueryUserRoleListResponseBodyFinanceEmpDeptOpenList) GoString() string {
	return s.String()
}

func (s *QueryUserRoleListResponseBodyFinanceEmpDeptOpenList) SetCascadeDeptId(v string) *QueryUserRoleListResponseBodyFinanceEmpDeptOpenList {
	s.CascadeDeptId = &v
	return s
}

func (s *QueryUserRoleListResponseBodyFinanceEmpDeptOpenList) SetDeptId(v int64) *QueryUserRoleListResponseBodyFinanceEmpDeptOpenList {
	s.DeptId = &v
	return s
}

func (s *QueryUserRoleListResponseBodyFinanceEmpDeptOpenList) SetName(v string) *QueryUserRoleListResponseBodyFinanceEmpDeptOpenList {
	s.Name = &v
	return s
}

func (s *QueryUserRoleListResponseBodyFinanceEmpDeptOpenList) SetSuperDeptId(v int64) *QueryUserRoleListResponseBodyFinanceEmpDeptOpenList {
	s.SuperDeptId = &v
	return s
}

type QueryUserRoleListResponseBodyRoleVOList struct {
	// example:
	//
	// applicationManager
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	// example:
	//
	// 应用管理员
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
}

func (s QueryUserRoleListResponseBodyRoleVOList) String() string {
	return tea.Prettify(s)
}

func (s QueryUserRoleListResponseBodyRoleVOList) GoString() string {
	return s.String()
}

func (s *QueryUserRoleListResponseBodyRoleVOList) SetRoleCode(v string) *QueryUserRoleListResponseBodyRoleVOList {
	s.RoleCode = &v
	return s
}

func (s *QueryUserRoleListResponseBodyRoleVOList) SetRoleName(v string) *QueryUserRoleListResponseBodyRoleVOList {
	s.RoleName = &v
	return s
}

type QueryUserRoleListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserRoleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserRoleListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserRoleListResponse) GoString() string {
	return s.String()
}

func (s *QueryUserRoleListResponse) SetHeaders(v map[string]*string) *QueryUserRoleListResponse {
	s.Headers = v
	return s
}

func (s *QueryUserRoleListResponse) SetStatusCode(v int32) *QueryUserRoleListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserRoleListResponse) SetBody(v *QueryUserRoleListResponseBody) *QueryUserRoleListResponse {
	s.Body = v
	return s
}

type SignEnterpriseAccountHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SignEnterpriseAccountHeaders) String() string {
	return tea.Prettify(s)
}

func (s SignEnterpriseAccountHeaders) GoString() string {
	return s.String()
}

func (s *SignEnterpriseAccountHeaders) SetCommonHeaders(v map[string]*string) *SignEnterpriseAccountHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SignEnterpriseAccountHeaders) SetXAcsDingtalkAccessToken(v string) *SignEnterpriseAccountHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SignEnterpriseAccountRequest struct {
	// example:
	//
	// ACC_XXX
	AccountCode *string `json:"accountCode,omitempty" xml:"accountCode,omitempty"`
	// example:
	//
	// 123
	BankCardNo *string `json:"bankCardNo,omitempty" xml:"bankCardNo,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 123
	BankOpenId *string `json:"bankOpenId,omitempty" xml:"bankOpenId,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// COMM_UNIONPAY
	ChannelType *string `json:"channelType,omitempty" xml:"channelType,omitempty"`
	// example:
	//
	// XXX
	FeeItemCode *string `json:"feeItemCode,omitempty" xml:"feeItemCode,omitempty"`
	// example:
	//
	// XXXX
	IssueNo *string `json:"issueNo,omitempty" xml:"issueNo,omitempty"`
	// example:
	//
	// 123
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// example:
	//
	// signed
	SignOperateType *string `json:"signOperateType,omitempty" xml:"signOperateType,omitempty"`
}

func (s SignEnterpriseAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s SignEnterpriseAccountRequest) GoString() string {
	return s.String()
}

func (s *SignEnterpriseAccountRequest) SetAccountCode(v string) *SignEnterpriseAccountRequest {
	s.AccountCode = &v
	return s
}

func (s *SignEnterpriseAccountRequest) SetBankCardNo(v string) *SignEnterpriseAccountRequest {
	s.BankCardNo = &v
	return s
}

func (s *SignEnterpriseAccountRequest) SetBankOpenId(v string) *SignEnterpriseAccountRequest {
	s.BankOpenId = &v
	return s
}

func (s *SignEnterpriseAccountRequest) SetChannelType(v string) *SignEnterpriseAccountRequest {
	s.ChannelType = &v
	return s
}

func (s *SignEnterpriseAccountRequest) SetFeeItemCode(v string) *SignEnterpriseAccountRequest {
	s.FeeItemCode = &v
	return s
}

func (s *SignEnterpriseAccountRequest) SetIssueNo(v string) *SignEnterpriseAccountRequest {
	s.IssueNo = &v
	return s
}

func (s *SignEnterpriseAccountRequest) SetOperator(v string) *SignEnterpriseAccountRequest {
	s.Operator = &v
	return s
}

func (s *SignEnterpriseAccountRequest) SetSignOperateType(v string) *SignEnterpriseAccountRequest {
	s.SignOperateType = &v
	return s
}

type SignEnterpriseAccountResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SignEnterpriseAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SignEnterpriseAccountResponseBody) GoString() string {
	return s.String()
}

func (s *SignEnterpriseAccountResponseBody) SetResult(v bool) *SignEnterpriseAccountResponseBody {
	s.Result = &v
	return s
}

type SignEnterpriseAccountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SignEnterpriseAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SignEnterpriseAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s SignEnterpriseAccountResponse) GoString() string {
	return s.String()
}

func (s *SignEnterpriseAccountResponse) SetHeaders(v map[string]*string) *SignEnterpriseAccountResponse {
	s.Headers = v
	return s
}

func (s *SignEnterpriseAccountResponse) SetStatusCode(v int32) *SignEnterpriseAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *SignEnterpriseAccountResponse) SetBody(v *SignEnterpriseAccountResponseBody) *SignEnterpriseAccountResponse {
	s.Body = v
	return s
}

type SyncReceiptRecallHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SyncReceiptRecallHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncReceiptRecallHeaders) GoString() string {
	return s.String()
}

func (s *SyncReceiptRecallHeaders) SetCommonHeaders(v map[string]*string) *SyncReceiptRecallHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncReceiptRecallHeaders) SetXAcsDingtalkAccessToken(v string) *SyncReceiptRecallHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SyncReceiptRecallRequest struct {
	// example:
	//
	// https:xxxx.pdf
	FileDownloadUrl *string `json:"fileDownloadUrl,omitempty" xml:"fileDownloadUrl,omitempty"`
	// example:
	//
	// 1234.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// 123
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
}

func (s SyncReceiptRecallRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncReceiptRecallRequest) GoString() string {
	return s.String()
}

func (s *SyncReceiptRecallRequest) SetFileDownloadUrl(v string) *SyncReceiptRecallRequest {
	s.FileDownloadUrl = &v
	return s
}

func (s *SyncReceiptRecallRequest) SetFileName(v string) *SyncReceiptRecallRequest {
	s.FileName = &v
	return s
}

func (s *SyncReceiptRecallRequest) SetOrderNo(v string) *SyncReceiptRecallRequest {
	s.OrderNo = &v
	return s
}

type SyncReceiptRecallResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SyncReceiptRecallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncReceiptRecallResponseBody) GoString() string {
	return s.String()
}

func (s *SyncReceiptRecallResponseBody) SetResult(v bool) *SyncReceiptRecallResponseBody {
	s.Result = &v
	return s
}

type SyncReceiptRecallResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncReceiptRecallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncReceiptRecallResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncReceiptRecallResponse) GoString() string {
	return s.String()
}

func (s *SyncReceiptRecallResponse) SetHeaders(v map[string]*string) *SyncReceiptRecallResponse {
	s.Headers = v
	return s
}

func (s *SyncReceiptRecallResponse) SetStatusCode(v int32) *SyncReceiptRecallResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncReceiptRecallResponse) SetBody(v *SyncReceiptRecallResponseBody) *SyncReceiptRecallResponse {
	s.Body = v
	return s
}

type UpdateAuxiliaryStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateAuxiliaryStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuxiliaryStatusHeaders) GoString() string {
	return s.String()
}

func (s *UpdateAuxiliaryStatusHeaders) SetCommonHeaders(v map[string]*string) *UpdateAuxiliaryStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateAuxiliaryStatusHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateAuxiliaryStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateAuxiliaryStatusRequest struct {
	AuxiliaryId   *string `json:"auxiliaryId,omitempty" xml:"auxiliaryId,omitempty"`
	AuxiliaryType *string `json:"auxiliaryType,omitempty" xml:"auxiliaryType,omitempty"`
	OperateType   *string `json:"operateType,omitempty" xml:"operateType,omitempty"`
	UserId        *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateAuxiliaryStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuxiliaryStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateAuxiliaryStatusRequest) SetAuxiliaryId(v string) *UpdateAuxiliaryStatusRequest {
	s.AuxiliaryId = &v
	return s
}

func (s *UpdateAuxiliaryStatusRequest) SetAuxiliaryType(v string) *UpdateAuxiliaryStatusRequest {
	s.AuxiliaryType = &v
	return s
}

func (s *UpdateAuxiliaryStatusRequest) SetOperateType(v string) *UpdateAuxiliaryStatusRequest {
	s.OperateType = &v
	return s
}

func (s *UpdateAuxiliaryStatusRequest) SetUserId(v string) *UpdateAuxiliaryStatusRequest {
	s.UserId = &v
	return s
}

type UpdateAuxiliaryStatusResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateAuxiliaryStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuxiliaryStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAuxiliaryStatusResponseBody) SetResult(v bool) *UpdateAuxiliaryStatusResponseBody {
	s.Result = &v
	return s
}

type UpdateAuxiliaryStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAuxiliaryStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAuxiliaryStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuxiliaryStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateAuxiliaryStatusResponse) SetHeaders(v map[string]*string) *UpdateAuxiliaryStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateAuxiliaryStatusResponse) SetStatusCode(v int32) *UpdateAuxiliaryStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAuxiliaryStatusResponse) SetBody(v *UpdateAuxiliaryStatusResponseBody) *UpdateAuxiliaryStatusResponse {
	s.Body = v
	return s
}

type UpdateFinanceEnterpriseAccountHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateFinanceEnterpriseAccountHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateFinanceEnterpriseAccountHeaders) GoString() string {
	return s.String()
}

func (s *UpdateFinanceEnterpriseAccountHeaders) SetCommonHeaders(v map[string]*string) *UpdateFinanceEnterpriseAccountHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateFinanceEnterpriseAccountHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateFinanceEnterpriseAccountHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateFinanceEnterpriseAccountRequest struct {
	// example:
	//
	// ACC_XXXXXX
	AccountCode *string `json:"accountCode,omitempty" xml:"accountCode,omitempty"`
	// example:
	//
	// 钉钉科技
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	// example:
	//
	// BANKCARD
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	BankCardNo  *string `json:"bankCardNo,omitempty" xml:"bankCardNo,omitempty"`
	// example:
	//
	// ICBC
	BankCode *string `json:"bankCode,omitempty" xml:"bankCode,omitempty"`
	// example:
	//
	// 中国工商银行
	BankName *string `json:"bankName,omitempty" xml:"bankName,omitempty"`
	// example:
	//
	// 杭州市
	City *string `json:"city,omitempty" xml:"city,omitempty"`
	// example:
	//
	// 账户描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 中国工商银行余杭分行
	OfficialName *string `json:"officialName,omitempty" xml:"officialName,omitempty"`
	// example:
	//
	// 1128878445
	OfficialNumber *string `json:"officialNumber,omitempty" xml:"officialNumber,omitempty"`
	// example:
	//
	// 浙江省
	Province *string `json:"province,omitempty" xml:"province,omitempty"`
	// example:
	//
	// smallScaleTaxpayer
	TaxNature *string `json:"taxNature,omitempty" xml:"taxNature,omitempty"`
	// example:
	//
	// 125481254812548
	TaxNo *string `json:"taxNo,omitempty" xml:"taxNo,omitempty"`
	// example:
	//
	// 5046195764756652
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateFinanceEnterpriseAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFinanceEnterpriseAccountRequest) GoString() string {
	return s.String()
}

func (s *UpdateFinanceEnterpriseAccountRequest) SetAccountCode(v string) *UpdateFinanceEnterpriseAccountRequest {
	s.AccountCode = &v
	return s
}

func (s *UpdateFinanceEnterpriseAccountRequest) SetAccountName(v string) *UpdateFinanceEnterpriseAccountRequest {
	s.AccountName = &v
	return s
}

func (s *UpdateFinanceEnterpriseAccountRequest) SetAccountType(v string) *UpdateFinanceEnterpriseAccountRequest {
	s.AccountType = &v
	return s
}

func (s *UpdateFinanceEnterpriseAccountRequest) SetBankCardNo(v string) *UpdateFinanceEnterpriseAccountRequest {
	s.BankCardNo = &v
	return s
}

func (s *UpdateFinanceEnterpriseAccountRequest) SetBankCode(v string) *UpdateFinanceEnterpriseAccountRequest {
	s.BankCode = &v
	return s
}

func (s *UpdateFinanceEnterpriseAccountRequest) SetBankName(v string) *UpdateFinanceEnterpriseAccountRequest {
	s.BankName = &v
	return s
}

func (s *UpdateFinanceEnterpriseAccountRequest) SetCity(v string) *UpdateFinanceEnterpriseAccountRequest {
	s.City = &v
	return s
}

func (s *UpdateFinanceEnterpriseAccountRequest) SetDescription(v string) *UpdateFinanceEnterpriseAccountRequest {
	s.Description = &v
	return s
}

func (s *UpdateFinanceEnterpriseAccountRequest) SetOfficialName(v string) *UpdateFinanceEnterpriseAccountRequest {
	s.OfficialName = &v
	return s
}

func (s *UpdateFinanceEnterpriseAccountRequest) SetOfficialNumber(v string) *UpdateFinanceEnterpriseAccountRequest {
	s.OfficialNumber = &v
	return s
}

func (s *UpdateFinanceEnterpriseAccountRequest) SetProvince(v string) *UpdateFinanceEnterpriseAccountRequest {
	s.Province = &v
	return s
}

func (s *UpdateFinanceEnterpriseAccountRequest) SetTaxNature(v string) *UpdateFinanceEnterpriseAccountRequest {
	s.TaxNature = &v
	return s
}

func (s *UpdateFinanceEnterpriseAccountRequest) SetTaxNo(v string) *UpdateFinanceEnterpriseAccountRequest {
	s.TaxNo = &v
	return s
}

func (s *UpdateFinanceEnterpriseAccountRequest) SetUserId(v string) *UpdateFinanceEnterpriseAccountRequest {
	s.UserId = &v
	return s
}

type UpdateFinanceEnterpriseAccountResponseBody struct {
	AccountCode *string `json:"accountCode,omitempty" xml:"accountCode,omitempty"`
}

func (s UpdateFinanceEnterpriseAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFinanceEnterpriseAccountResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFinanceEnterpriseAccountResponseBody) SetAccountCode(v string) *UpdateFinanceEnterpriseAccountResponseBody {
	s.AccountCode = &v
	return s
}

type UpdateFinanceEnterpriseAccountResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFinanceEnterpriseAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFinanceEnterpriseAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFinanceEnterpriseAccountResponse) GoString() string {
	return s.String()
}

func (s *UpdateFinanceEnterpriseAccountResponse) SetHeaders(v map[string]*string) *UpdateFinanceEnterpriseAccountResponse {
	s.Headers = v
	return s
}

func (s *UpdateFinanceEnterpriseAccountResponse) SetStatusCode(v int32) *UpdateFinanceEnterpriseAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFinanceEnterpriseAccountResponse) SetBody(v *UpdateFinanceEnterpriseAccountResponseBody) *UpdateFinanceEnterpriseAccountResponse {
	s.Body = v
	return s
}

type UpdateInstanceOrderInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateInstanceOrderInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceOrderInfoHeaders) GoString() string {
	return s.String()
}

func (s *UpdateInstanceOrderInfoHeaders) SetCommonHeaders(v map[string]*string) *UpdateInstanceOrderInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateInstanceOrderInfoHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateInstanceOrderInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateInstanceOrderInfoRequest struct {
	// example:
	//
	// xxx错误
	FailReason *string `json:"failReason,omitempty" xml:"failReason,omitempty"`
	// example:
	//
	// abc
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	// example:
	//
	// 123
	OutOrderNo *string                                  `json:"outOrderNo,omitempty" xml:"outOrderNo,omitempty"`
	PayerBank  *UpdateInstanceOrderInfoRequestPayerBank `json:"payerBank,omitempty" xml:"payerBank,omitempty" type:"Struct"`
	// example:
	//
	// 1709691000682
	PaymentTime *int64 `json:"paymentTime,omitempty" xml:"paymentTime,omitempty"`
	// example:
	//
	// PAYING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateInstanceOrderInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceOrderInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceOrderInfoRequest) SetFailReason(v string) *UpdateInstanceOrderInfoRequest {
	s.FailReason = &v
	return s
}

func (s *UpdateInstanceOrderInfoRequest) SetOrderNo(v string) *UpdateInstanceOrderInfoRequest {
	s.OrderNo = &v
	return s
}

func (s *UpdateInstanceOrderInfoRequest) SetOutOrderNo(v string) *UpdateInstanceOrderInfoRequest {
	s.OutOrderNo = &v
	return s
}

func (s *UpdateInstanceOrderInfoRequest) SetPayerBank(v *UpdateInstanceOrderInfoRequestPayerBank) *UpdateInstanceOrderInfoRequest {
	s.PayerBank = v
	return s
}

func (s *UpdateInstanceOrderInfoRequest) SetPaymentTime(v int64) *UpdateInstanceOrderInfoRequest {
	s.PaymentTime = &v
	return s
}

func (s *UpdateInstanceOrderInfoRequest) SetStatus(v string) *UpdateInstanceOrderInfoRequest {
	s.Status = &v
	return s
}

func (s *UpdateInstanceOrderInfoRequest) SetUserId(v string) *UpdateInstanceOrderInfoRequest {
	s.UserId = &v
	return s
}

type UpdateInstanceOrderInfoRequestPayerBank struct {
	// example:
	//
	// 64112222
	CardNo *string `json:"cardNo,omitempty" xml:"cardNo,omitempty"`
	// example:
	//
	// 招商银行
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateInstanceOrderInfoRequestPayerBank) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceOrderInfoRequestPayerBank) GoString() string {
	return s.String()
}

func (s *UpdateInstanceOrderInfoRequestPayerBank) SetCardNo(v string) *UpdateInstanceOrderInfoRequestPayerBank {
	s.CardNo = &v
	return s
}

func (s *UpdateInstanceOrderInfoRequestPayerBank) SetName(v string) *UpdateInstanceOrderInfoRequestPayerBank {
	s.Name = &v
	return s
}

type UpdateInstanceOrderInfoShrinkRequest struct {
	// example:
	//
	// xxx错误
	FailReason *string `json:"failReason,omitempty" xml:"failReason,omitempty"`
	// example:
	//
	// abc
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	// example:
	//
	// 123
	OutOrderNo      *string `json:"outOrderNo,omitempty" xml:"outOrderNo,omitempty"`
	PayerBankShrink *string `json:"payerBank,omitempty" xml:"payerBank,omitempty"`
	// example:
	//
	// 1709691000682
	PaymentTime *int64 `json:"paymentTime,omitempty" xml:"paymentTime,omitempty"`
	// example:
	//
	// PAYING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateInstanceOrderInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceOrderInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceOrderInfoShrinkRequest) SetFailReason(v string) *UpdateInstanceOrderInfoShrinkRequest {
	s.FailReason = &v
	return s
}

func (s *UpdateInstanceOrderInfoShrinkRequest) SetOrderNo(v string) *UpdateInstanceOrderInfoShrinkRequest {
	s.OrderNo = &v
	return s
}

func (s *UpdateInstanceOrderInfoShrinkRequest) SetOutOrderNo(v string) *UpdateInstanceOrderInfoShrinkRequest {
	s.OutOrderNo = &v
	return s
}

func (s *UpdateInstanceOrderInfoShrinkRequest) SetPayerBankShrink(v string) *UpdateInstanceOrderInfoShrinkRequest {
	s.PayerBankShrink = &v
	return s
}

func (s *UpdateInstanceOrderInfoShrinkRequest) SetPaymentTime(v int64) *UpdateInstanceOrderInfoShrinkRequest {
	s.PaymentTime = &v
	return s
}

func (s *UpdateInstanceOrderInfoShrinkRequest) SetStatus(v string) *UpdateInstanceOrderInfoShrinkRequest {
	s.Status = &v
	return s
}

func (s *UpdateInstanceOrderInfoShrinkRequest) SetUserId(v string) *UpdateInstanceOrderInfoShrinkRequest {
	s.UserId = &v
	return s
}

type UpdateInstanceOrderInfoResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateInstanceOrderInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceOrderInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceOrderInfoResponseBody) SetResult(v bool) *UpdateInstanceOrderInfoResponseBody {
	s.Result = &v
	return s
}

type UpdateInstanceOrderInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceOrderInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceOrderInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceOrderInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceOrderInfoResponse) SetHeaders(v map[string]*string) *UpdateInstanceOrderInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceOrderInfoResponse) SetStatusCode(v int32) *UpdateInstanceOrderInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceOrderInfoResponse) SetBody(v *UpdateInstanceOrderInfoResponseBody) *UpdateInstanceOrderInfoResponse {
	s.Body = v
	return s
}

type UpdateInvoiceDataTransferDoneHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateInvoiceDataTransferDoneHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceDataTransferDoneHeaders) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceDataTransferDoneHeaders) SetCommonHeaders(v map[string]*string) *UpdateInvoiceDataTransferDoneHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateInvoiceDataTransferDoneHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateInvoiceDataTransferDoneHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateInvoiceDataTransferDoneResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateInvoiceDataTransferDoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceDataTransferDoneResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceDataTransferDoneResponseBody) SetSuccess(v bool) *UpdateInvoiceDataTransferDoneResponseBody {
	s.Success = &v
	return s
}

type UpdateInvoiceDataTransferDoneResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInvoiceDataTransferDoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInvoiceDataTransferDoneResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceDataTransferDoneResponse) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceDataTransferDoneResponse) SetHeaders(v map[string]*string) *UpdateInvoiceDataTransferDoneResponse {
	s.Headers = v
	return s
}

func (s *UpdateInvoiceDataTransferDoneResponse) SetStatusCode(v int32) *UpdateInvoiceDataTransferDoneResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInvoiceDataTransferDoneResponse) SetBody(v *UpdateInvoiceDataTransferDoneResponseBody) *UpdateInvoiceDataTransferDoneResponse {
	s.Body = v
	return s
}

type UpdateInvoiceUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateInvoiceUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceUrlHeaders) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceUrlHeaders) SetCommonHeaders(v map[string]*string) *UpdateInvoiceUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateInvoiceUrlHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateInvoiceUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateInvoiceUrlRequest struct {
	Body *UpdateInvoiceUrlRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s UpdateInvoiceUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceUrlRequest) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceUrlRequest) SetBody(v *UpdateInvoiceUrlRequestBody) *UpdateInvoiceUrlRequest {
	s.Body = v
	return s
}

type UpdateInvoiceUrlRequestBody struct {
	CompanyCode *string                               `json:"companyCode,omitempty" xml:"companyCode,omitempty"`
	Operator    *string                               `json:"operator,omitempty" xml:"operator,omitempty"`
	UrlList     []*UpdateInvoiceUrlRequestBodyUrlList `json:"urlList,omitempty" xml:"urlList,omitempty" type:"Repeated"`
}

func (s UpdateInvoiceUrlRequestBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceUrlRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceUrlRequestBody) SetCompanyCode(v string) *UpdateInvoiceUrlRequestBody {
	s.CompanyCode = &v
	return s
}

func (s *UpdateInvoiceUrlRequestBody) SetOperator(v string) *UpdateInvoiceUrlRequestBody {
	s.Operator = &v
	return s
}

func (s *UpdateInvoiceUrlRequestBody) SetUrlList(v []*UpdateInvoiceUrlRequestBodyUrlList) *UpdateInvoiceUrlRequestBody {
	s.UrlList = v
	return s
}

type UpdateInvoiceUrlRequestBodyUrlList struct {
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	InvoiceNo   *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
	OfdUrl      *string `json:"ofdUrl,omitempty" xml:"ofdUrl,omitempty"`
	PdfUrl      *string `json:"pdfUrl,omitempty" xml:"pdfUrl,omitempty"`
	XmlUrl      *string `json:"xmlUrl,omitempty" xml:"xmlUrl,omitempty"`
}

func (s UpdateInvoiceUrlRequestBodyUrlList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceUrlRequestBodyUrlList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceUrlRequestBodyUrlList) SetInvoiceCode(v string) *UpdateInvoiceUrlRequestBodyUrlList {
	s.InvoiceCode = &v
	return s
}

func (s *UpdateInvoiceUrlRequestBodyUrlList) SetInvoiceNo(v string) *UpdateInvoiceUrlRequestBodyUrlList {
	s.InvoiceNo = &v
	return s
}

func (s *UpdateInvoiceUrlRequestBodyUrlList) SetOfdUrl(v string) *UpdateInvoiceUrlRequestBodyUrlList {
	s.OfdUrl = &v
	return s
}

func (s *UpdateInvoiceUrlRequestBodyUrlList) SetPdfUrl(v string) *UpdateInvoiceUrlRequestBodyUrlList {
	s.PdfUrl = &v
	return s
}

func (s *UpdateInvoiceUrlRequestBodyUrlList) SetXmlUrl(v string) *UpdateInvoiceUrlRequestBodyUrlList {
	s.XmlUrl = &v
	return s
}

type UpdateInvoiceUrlShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInvoiceUrlShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceUrlShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceUrlShrinkRequest) SetBodyShrink(v string) *UpdateInvoiceUrlShrinkRequest {
	s.BodyShrink = &v
	return s
}

type UpdateInvoiceUrlResponseBody struct {
	Result *UpdateInvoiceUrlResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateInvoiceUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceUrlResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceUrlResponseBody) SetResult(v *UpdateInvoiceUrlResponseBodyResult) *UpdateInvoiceUrlResponseBody {
	s.Result = v
	return s
}

type UpdateInvoiceUrlResponseBodyResult struct {
	FailInvoiceList []*UpdateInvoiceUrlResponseBodyResultFailInvoiceList `json:"failInvoiceList,omitempty" xml:"failInvoiceList,omitempty" type:"Repeated"`
	IsAllSuccess    *string                                              `json:"isAllSuccess,omitempty" xml:"isAllSuccess,omitempty"`
}

func (s UpdateInvoiceUrlResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceUrlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceUrlResponseBodyResult) SetFailInvoiceList(v []*UpdateInvoiceUrlResponseBodyResultFailInvoiceList) *UpdateInvoiceUrlResponseBodyResult {
	s.FailInvoiceList = v
	return s
}

func (s *UpdateInvoiceUrlResponseBodyResult) SetIsAllSuccess(v string) *UpdateInvoiceUrlResponseBodyResult {
	s.IsAllSuccess = &v
	return s
}

type UpdateInvoiceUrlResponseBodyResultFailInvoiceList struct {
	ErrorMsg    *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	InvoiceNo   *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
}

func (s UpdateInvoiceUrlResponseBodyResultFailInvoiceList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceUrlResponseBodyResultFailInvoiceList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceUrlResponseBodyResultFailInvoiceList) SetErrorMsg(v string) *UpdateInvoiceUrlResponseBodyResultFailInvoiceList {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateInvoiceUrlResponseBodyResultFailInvoiceList) SetInvoiceCode(v string) *UpdateInvoiceUrlResponseBodyResultFailInvoiceList {
	s.InvoiceCode = &v
	return s
}

func (s *UpdateInvoiceUrlResponseBodyResultFailInvoiceList) SetInvoiceNo(v string) *UpdateInvoiceUrlResponseBodyResultFailInvoiceList {
	s.InvoiceNo = &v
	return s
}

type UpdateInvoiceUrlResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInvoiceUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInvoiceUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceUrlResponse) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceUrlResponse) SetHeaders(v map[string]*string) *UpdateInvoiceUrlResponse {
	s.Headers = v
	return s
}

func (s *UpdateInvoiceUrlResponse) SetStatusCode(v int32) *UpdateInvoiceUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInvoiceUrlResponse) SetBody(v *UpdateInvoiceUrlResponseBody) *UpdateInvoiceUrlResponse {
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
// 新增智能财务的企业账户
//
// @param request - AddFinanceEnterpriseAccountRequest
//
// @param headers - AddFinanceEnterpriseAccountHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFinanceEnterpriseAccountResponse
func (client *Client) AddFinanceEnterpriseAccountWithOptions(request *AddFinanceEnterpriseAccountRequest, headers *AddFinanceEnterpriseAccountHeaders, runtime *util.RuntimeOptions) (_result *AddFinanceEnterpriseAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["accountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountType)) {
		body["accountType"] = request.AccountType
	}

	if !tea.BoolValue(util.IsUnset(request.BankCardNo)) {
		body["bankCardNo"] = request.BankCardNo
	}

	if !tea.BoolValue(util.IsUnset(request.BankCode)) {
		body["bankCode"] = request.BankCode
	}

	if !tea.BoolValue(util.IsUnset(request.BankName)) {
		body["bankName"] = request.BankName
	}

	if !tea.BoolValue(util.IsUnset(request.City)) {
		body["city"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.OfficialName)) {
		body["officialName"] = request.OfficialName
	}

	if !tea.BoolValue(util.IsUnset(request.OfficialNumber)) {
		body["officialNumber"] = request.OfficialNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Province)) {
		body["province"] = request.Province
	}

	if !tea.BoolValue(util.IsUnset(request.TaxNature)) {
		body["taxNature"] = request.TaxNature
	}

	if !tea.BoolValue(util.IsUnset(request.TaxNo)) {
		body["taxNo"] = request.TaxNo
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
		Action:      tea.String("AddFinanceEnterpriseAccount"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/enterpriseAccounts"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddFinanceEnterpriseAccountResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增智能财务的企业账户
//
// @param request - AddFinanceEnterpriseAccountRequest
//
// @return AddFinanceEnterpriseAccountResponse
func (client *Client) AddFinanceEnterpriseAccount(request *AddFinanceEnterpriseAccountRequest) (_result *AddFinanceEnterpriseAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddFinanceEnterpriseAccountHeaders{}
	_result = &AddFinanceEnterpriseAccountResponse{}
	_body, _err := client.AddFinanceEnterpriseAccountWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 银行接入层通用接口
//
// @param request - BankGatewayInvokeRequest
//
// @param headers - BankGatewayInvokeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BankGatewayInvokeResponse
func (client *Client) BankGatewayInvokeWithOptions(request *BankGatewayInvokeRequest, headers *BankGatewayInvokeHeaders, runtime *util.RuntimeOptions) (_result *BankGatewayInvokeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionType)) {
		body["actionType"] = request.ActionType
	}

	if !tea.BoolValue(util.IsUnset(request.InputData)) {
		body["inputData"] = request.InputData
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["url"] = request.Url
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
		Action:      tea.String("BankGatewayInvoke"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/bankGateways/invoke"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BankGatewayInvokeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 银行接入层通用接口
//
// @param request - BankGatewayInvokeRequest
//
// @return BankGatewayInvokeResponse
func (client *Client) BankGatewayInvoke(request *BankGatewayInvokeRequest) (_result *BankGatewayInvokeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BankGatewayInvokeHeaders{}
	_result = &BankGatewayInvokeResponse{}
	_body, _err := client.BankGatewayInvokeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量删除智能财务单据
//
// @param request - BatchDeleteReceiptRequest
//
// @param headers - BatchDeleteReceiptHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteReceiptResponse
func (client *Client) BatchDeleteReceiptWithOptions(request *BatchDeleteReceiptRequest, headers *BatchDeleteReceiptHeaders, runtime *util.RuntimeOptions) (_result *BatchDeleteReceiptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIdList)) {
		body["instanceIdList"] = request.InstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["operator"] = request.Operator
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
		Action:      tea.String("BatchDeleteReceipt"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/instances/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchDeleteReceiptResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除智能财务单据
//
// @param request - BatchDeleteReceiptRequest
//
// @return BatchDeleteReceiptResponse
func (client *Client) BatchDeleteReceipt(request *BatchDeleteReceiptRequest) (_result *BatchDeleteReceiptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchDeleteReceiptHeaders{}
	_result = &BatchDeleteReceiptResponse{}
	_body, _err := client.BatchDeleteReceiptWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询企业票池发票下载链接
//
// @param request - BatchQueryOrgInvoiceUrlRequest
//
// @param headers - BatchQueryOrgInvoiceUrlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchQueryOrgInvoiceUrlResponse
func (client *Client) BatchQueryOrgInvoiceUrlWithOptions(request *BatchQueryOrgInvoiceUrlRequest, headers *BatchQueryOrgInvoiceUrlHeaders, runtime *util.RuntimeOptions) (_result *BatchQueryOrgInvoiceUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompanyCode)) {
		body["companyCode"] = request.CompanyCode
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceKeyVOList)) {
		body["invoiceKeyVOList"] = request.InvoiceKeyVOList
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["operator"] = request.Operator
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
		Action:      tea.String("BatchQueryOrgInvoiceUrl"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/invoices/urls/batchQuery"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchQueryOrgInvoiceUrlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询企业票池发票下载链接
//
// @param request - BatchQueryOrgInvoiceUrlRequest
//
// @return BatchQueryOrgInvoiceUrlResponse
func (client *Client) BatchQueryOrgInvoiceUrl(request *BatchQueryOrgInvoiceUrlRequest) (_result *BatchQueryOrgInvoiceUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchQueryOrgInvoiceUrlHeaders{}
	_result = &BatchQueryOrgInvoiceUrlResponse{}
	_body, _err := client.BatchQueryOrgInvoiceUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询支付回单文件
//
// @param request - BatchQueryPaymentRecallFileRequest
//
// @param headers - BatchQueryPaymentRecallFileHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchQueryPaymentRecallFileResponse
func (client *Client) BatchQueryPaymentRecallFileWithOptions(request *BatchQueryPaymentRecallFileRequest, headers *BatchQueryPaymentRecallFileHeaders, runtime *util.RuntimeOptions) (_result *BatchQueryPaymentRecallFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DetailIdList)) {
		body["detailIdList"] = request.DetailIdList
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["operator"] = request.Operator
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
		Action:      tea.String("BatchQueryPaymentRecallFile"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/payments/recallFiles/batchQuery"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchQueryPaymentRecallFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询支付回单文件
//
// @param request - BatchQueryPaymentRecallFileRequest
//
// @return BatchQueryPaymentRecallFileResponse
func (client *Client) BatchQueryPaymentRecallFile(request *BatchQueryPaymentRecallFileRequest) (_result *BatchQueryPaymentRecallFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchQueryPaymentRecallFileHeaders{}
	_result = &BatchQueryPaymentRecallFileResponse{}
	_body, _err := client.BatchQueryPaymentRecallFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量同步银行回单
//
// @param request - BatchSyncBankReceiptRequest
//
// @param headers - BatchSyncBankReceiptHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSyncBankReceiptResponse
func (client *Client) BatchSyncBankReceiptWithOptions(request *BatchSyncBankReceiptRequest, headers *BatchSyncBankReceiptHeaders, runtime *util.RuntimeOptions) (_result *BatchSyncBankReceiptResponse, _err error) {
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
		Action:      tea.String("BatchSyncBankReceipt"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/receipts/batchSync"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchSyncBankReceiptResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量同步银行回单
//
// @param request - BatchSyncBankReceiptRequest
//
// @return BatchSyncBankReceiptResponse
func (client *Client) BatchSyncBankReceipt(request *BatchSyncBankReceiptRequest) (_result *BatchSyncBankReceiptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchSyncBankReceiptHeaders{}
	_result = &BatchSyncBankReceiptResponse{}
	_body, _err := client.BatchSyncBankReceiptWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查验发票是否生成凭证
//
// @param request - CheckVoucherStatusRequest
//
// @param headers - CheckVoucherStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckVoucherStatusResponse
func (client *Client) CheckVoucherStatusWithOptions(request *CheckVoucherStatusRequest, headers *CheckVoucherStatusHeaders, runtime *util.RuntimeOptions) (_result *CheckVoucherStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompanyCode)) {
		body["companyCode"] = request.CompanyCode
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FinanceType)) {
		body["financeType"] = request.FinanceType
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceCode)) {
		body["invoiceCode"] = request.InvoiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceNo)) {
		body["invoiceNo"] = request.InvoiceNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaxNo)) {
		body["taxNo"] = request.TaxNo
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyStatus)) {
		body["verifyStatus"] = request.VerifyStatus
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
		Action:      tea.String("CheckVoucherStatus"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/invoices/checkVoucherStatus/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckVoucherStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查验发票是否生成凭证
//
// @param request - CheckVoucherStatusRequest
//
// @return CheckVoucherStatusResponse
func (client *Client) CheckVoucherStatus(request *CheckVoucherStatusRequest) (_result *CheckVoucherStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CheckVoucherStatusHeaders{}
	_result = &CheckVoucherStatusResponse{}
	_body, _err := client.CheckVoucherStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取唤起智能财务收银台的地址
//
// @param request - ConfirmPaymentOrderRequest
//
// @param headers - ConfirmPaymentOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfirmPaymentOrderResponse
func (client *Client) ConfirmPaymentOrderWithOptions(request *ConfirmPaymentOrderRequest, headers *ConfirmPaymentOrderHeaders, runtime *util.RuntimeOptions) (_result *ConfirmPaymentOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutBizNoList)) {
		body["outBizNoList"] = request.OutBizNoList
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
		Action:      tea.String("ConfirmPaymentOrder"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/payments/confirm"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfirmPaymentOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取唤起智能财务收银台的地址
//
// @param request - ConfirmPaymentOrderRequest
//
// @return ConfirmPaymentOrderResponse
func (client *Client) ConfirmPaymentOrder(request *ConfirmPaymentOrderRequest) (_result *ConfirmPaymentOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ConfirmPaymentOrderHeaders{}
	_result = &ConfirmPaymentOrderResponse{}
	_body, _err := client.ConfirmPaymentOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建收款订单
//
// @param request - CreateCollectionOrderRequest
//
// @param headers - CreateCollectionOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCollectionOrderResponse
func (client *Client) CreateCollectionOrderWithOptions(request *CreateCollectionOrderRequest, headers *CreateCollectionOrderHeaders, runtime *util.RuntimeOptions) (_result *CreateCollectionOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		query["amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.CollectionInfoId)) {
		query["collectionInfoId"] = request.CollectionInfoId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["remark"] = request.Remark
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
		Action:      tea.String("CreateCollectionOrder"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/me/collections/orders"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCollectionOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建收款订单
//
// @param request - CreateCollectionOrderRequest
//
// @return CreateCollectionOrderResponse
func (client *Client) CreateCollectionOrder(request *CreateCollectionOrderRequest) (_result *CreateCollectionOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateCollectionOrderHeaders{}
	_result = &CreateCollectionOrderResponse{}
	_body, _err := client.CreateCollectionOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建智能财务待支付订单
//
// @param request - CreatePaymentOrderRequest
//
// @param headers - CreatePaymentOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePaymentOrderResponse
func (client *Client) CreatePaymentOrderWithOptions(request *CreatePaymentOrderRequest, headers *CreatePaymentOrderHeaders, runtime *util.RuntimeOptions) (_result *CreatePaymentOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		body["amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		body["expireTime"] = request.ExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.OutBizNo)) {
		body["outBizNo"] = request.OutBizNo
	}

	if !tea.BoolValue(util.IsUnset(request.PayeeAccountDTO)) {
		body["payeeAccountDTO"] = request.PayeeAccountDTO
	}

	if !tea.BoolValue(util.IsUnset(request.PayerAccountDTO)) {
		body["payerAccountDTO"] = request.PayerAccountDTO
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentOrderTitle)) {
		body["paymentOrderTitle"] = request.PaymentOrderTitle
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Usage)) {
		body["usage"] = request.Usage
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
		Action:      tea.String("CreatePaymentOrder"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/payments/orders"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePaymentOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建智能财务待支付订单
//
// @param request - CreatePaymentOrderRequest
//
// @return CreatePaymentOrderResponse
func (client *Client) CreatePaymentOrder(request *CreatePaymentOrderRequest) (_result *CreatePaymentOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreatePaymentOrderHeaders{}
	_result = &CreatePaymentOrderResponse{}
	_body, _err := client.CreatePaymentOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取费用类别
//
// @param request - GetCategoryRequest
//
// @param headers - GetCategoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCategoryResponse
func (client *Client) GetCategoryWithOptions(request *GetCategoryRequest, headers *GetCategoryHeaders, runtime *util.RuntimeOptions) (_result *GetCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["code"] = request.Code
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
		Action:      tea.String("GetCategory"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/categories"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCategoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取费用类别
//
// @param request - GetCategoryRequest
//
// @return GetCategoryResponse
func (client *Client) GetCategory(request *GetCategoryRequest) (_result *GetCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCategoryHeaders{}
	_result = &GetCategoryResponse{}
	_body, _err := client.GetCategoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询企业内自定义辅助档案信息
//
// @param request - GetDefineRequest
//
// @param headers - GetDefineHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDefineResponse
func (client *Client) GetDefineWithOptions(request *GetDefineRequest, headers *GetDefineHeaders, runtime *util.RuntimeOptions) (_result *GetDefineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
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
		Action:      tea.String("GetDefine"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/customInfos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDefineResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询企业内自定义辅助档案信息
//
// @param request - GetDefineRequest
//
// @return GetDefineResponse
func (client *Client) GetDefine(request *GetDefineRequest) (_result *GetDefineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDefineHeaders{}
	_result = &GetDefineResponse{}
	_body, _err := client.GetDefineWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询企业内自定义辅助档案数据信息
//
// @param request - GetDefineDataRequest
//
// @param headers - GetDefineDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDefineDataResponse
func (client *Client) GetDefineDataWithOptions(request *GetDefineDataRequest, headers *GetDefineDataHeaders, runtime *util.RuntimeOptions) (_result *GetDefineDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
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
		Action:      tea.String("GetDefineData"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/customDatas/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDefineDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询企业内自定义辅助档案数据信息
//
// @param request - GetDefineDataRequest
//
// @return GetDefineDataResponse
func (client *Client) GetDefineData(request *GetDefineDataRequest) (_result *GetDefineDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDefineDataHeaders{}
	_result = &GetDefineDataResponse{}
	_body, _err := client.GetDefineDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业账户
//
// @param request - GetFinanceAccountRequest
//
// @param headers - GetFinanceAccountHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFinanceAccountResponse
func (client *Client) GetFinanceAccountWithOptions(request *GetFinanceAccountRequest, headers *GetFinanceAccountHeaders, runtime *util.RuntimeOptions) (_result *GetFinanceAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountCode)) {
		query["accountCode"] = request.AccountCode
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
		Action:      tea.String("GetFinanceAccount"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/financeAccounts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFinanceAccountResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业账户
//
// @param request - GetFinanceAccountRequest
//
// @return GetFinanceAccountResponse
func (client *Client) GetFinanceAccount(request *GetFinanceAccountRequest) (_result *GetFinanceAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFinanceAccountHeaders{}
	_result = &GetFinanceAccountResponse{}
	_body, _err := client.GetFinanceAccountWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取单条项目
//
// @param request - GetProjectRequest
//
// @param headers - GetProjectHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectResponse
func (client *Client) GetProjectWithOptions(request *GetProjectRequest, headers *GetProjectHeaders, runtime *util.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["code"] = request.Code
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
		Action:      tea.String("GetProject"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/projects"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取单条项目
//
// @param request - GetProjectRequest
//
// @return GetProjectResponse
func (client *Client) GetProject(request *GetProjectRequest) (_result *GetProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetProjectHeaders{}
	_result = &GetProjectResponse{}
	_body, _err := client.GetProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取智能财务单据详情
//
// @param request - GetReceiptRequest
//
// @param headers - GetReceiptHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetReceiptResponse
func (client *Client) GetReceiptWithOptions(request *GetReceiptRequest, headers *GetReceiptHeaders, runtime *util.RuntimeOptions) (_result *GetReceiptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		query["modelId"] = request.ModelId
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
		Action:      tea.String("GetReceipt"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/receipts/details"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetReceiptResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取智能财务单据详情
//
// @param request - GetReceiptRequest
//
// @return GetReceiptResponse
func (client *Client) GetReceipt(request *GetReceiptRequest) (_result *GetReceiptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetReceiptHeaders{}
	_result = &GetReceiptResponse{}
	_body, _err := client.GetReceiptWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取智能财务应用内维护的供应商信息
//
// @param request - GetSupplierRequest
//
// @param headers - GetSupplierHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSupplierResponse
func (client *Client) GetSupplierWithOptions(request *GetSupplierRequest, headers *GetSupplierHeaders, runtime *util.RuntimeOptions) (_result *GetSupplierResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["code"] = request.Code
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
		Action:      tea.String("GetSupplier"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/suppliers/details"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSupplierResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取智能财务应用内维护的供应商信息
//
// @param request - GetSupplierRequest
//
// @return GetSupplierResponse
func (client *Client) GetSupplier(request *GetSupplierRequest) (_result *GetSupplierResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSupplierHeaders{}
	_result = &GetSupplierResponse{}
	_body, _err := client.GetSupplierWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 订单开票
//
// @param request - IssueInvoiceWithOrderRequest
//
// @param headers - IssueInvoiceWithOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IssueInvoiceWithOrderResponse
func (client *Client) IssueInvoiceWithOrderWithOptions(request *IssueInvoiceWithOrderRequest, headers *IssueInvoiceWithOrderHeaders, runtime *util.RuntimeOptions) (_result *IssueInvoiceWithOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.FinanceAppKey)) {
		body["financeAppKey"] = request.FinanceAppKey
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		body["signature"] = request.Signature
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
		Action:      tea.String("IssueInvoiceWithOrder"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/issueInvoices/order"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &IssueInvoiceWithOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 订单开票
//
// @param request - IssueInvoiceWithOrderRequest
//
// @return IssueInvoiceWithOrderResponse
func (client *Client) IssueInvoiceWithOrder(request *IssueInvoiceWithOrderRequest) (_result *IssueInvoiceWithOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IssueInvoiceWithOrderHeaders{}
	_result = &IssueInvoiceWithOrderResponse{}
	_body, _err := client.IssueInvoiceWithOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据不同的bizType查询不同的数据
//
// @param request - LinkCommonInvokeRequest
//
// @param headers - LinkCommonInvokeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LinkCommonInvokeResponse
func (client *Client) LinkCommonInvokeWithOptions(request *LinkCommonInvokeRequest, headers *LinkCommonInvokeHeaders, runtime *util.RuntimeOptions) (_result *LinkCommonInvokeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.InvokeId)) {
		body["invokeId"] = request.InvokeId
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
		Action:      tea.String("LinkCommonInvoke"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/link/bizTypes/invoke"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &LinkCommonInvokeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据不同的bizType查询不同的数据
//
// @param request - LinkCommonInvokeRequest
//
// @return LinkCommonInvokeResponse
func (client *Client) LinkCommonInvoke(request *LinkCommonInvokeRequest) (_result *LinkCommonInvokeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &LinkCommonInvokeHeaders{}
	_result = &LinkCommonInvokeResponse{}
	_body, _err := client.LinkCommonInvokeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 订单开票
//
// @param request - OrderBillingRequest
//
// @param headers - OrderBillingHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OrderBillingResponse
func (client *Client) OrderBillingWithOptions(request *OrderBillingRequest, headers *OrderBillingHeaders, runtime *util.RuntimeOptions) (_result *OrderBillingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdditionInfos)) {
		body["additionInfos"] = request.AdditionInfos
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["appKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.ApplyPerson)) {
		body["applyPerson"] = request.ApplyPerson
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceRemark)) {
		body["invoiceRemark"] = request.InvoiceRemark
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceType)) {
		body["invoiceType"] = request.InvoiceType
	}

	if !tea.BoolValue(util.IsUnset(request.IsNaturalPerson)) {
		body["isNaturalPerson"] = request.IsNaturalPerson
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		body["orderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.Payee)) {
		body["payee"] = request.Payee
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		body["phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.Products)) {
		body["products"] = request.Products
	}

	if !tea.BoolValue(util.IsUnset(request.PurchaserAddress)) {
		body["purchaserAddress"] = request.PurchaserAddress
	}

	if !tea.BoolValue(util.IsUnset(request.PurchaserBankAccount)) {
		body["purchaserBankAccount"] = request.PurchaserBankAccount
	}

	if !tea.BoolValue(util.IsUnset(request.PurchaserBankName)) {
		body["purchaserBankName"] = request.PurchaserBankName
	}

	if !tea.BoolValue(util.IsUnset(request.PurchaserName)) {
		body["purchaserName"] = request.PurchaserName
	}

	if !tea.BoolValue(util.IsUnset(request.PurchaserTaxNo)) {
		body["purchaserTaxNo"] = request.PurchaserTaxNo
	}

	if !tea.BoolValue(util.IsUnset(request.PurchaserTel)) {
		body["purchaserTel"] = request.PurchaserTel
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Reviewer)) {
		body["reviewer"] = request.Reviewer
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		body["signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.TaxSign)) {
		body["taxSign"] = request.TaxSign
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
		Action:      tea.String("OrderBilling"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/billings/order"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &OrderBillingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 订单开票
//
// @param request - OrderBillingRequest
//
// @return OrderBillingResponse
func (client *Client) OrderBilling(request *OrderBillingRequest) (_result *OrderBillingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OrderBillingHeaders{}
	_result = &OrderBillingResponse{}
	_body, _err := client.OrderBillingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询账户的银行交易流水
//
// @param request - QueryAccountTradeByPageRequest
//
// @param headers - QueryAccountTradeByPageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAccountTradeByPageResponse
func (client *Client) QueryAccountTradeByPageWithOptions(request *QueryAccountTradeByPageRequest, headers *QueryAccountTradeByPageHeaders, runtime *util.RuntimeOptions) (_result *QueryAccountTradeByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["accountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["endDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		body["filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["startDate"] = request.StartDate
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
		Action:      tea.String("QueryAccountTradeByPage"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/payments/trades/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAccountTradeByPageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询账户的银行交易流水
//
// @param request - QueryAccountTradeByPageRequest
//
// @return QueryAccountTradeByPageResponse
func (client *Client) QueryAccountTradeByPage(request *QueryAccountTradeByPageRequest) (_result *QueryAccountTradeByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAccountTradeByPageHeaders{}
	_result = &QueryAccountTradeByPageResponse{}
	_body, _err := client.QueryAccountTradeByPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量获取费用类别
//
// @param request - QueryCategoryByPageRequest
//
// @param headers - QueryCategoryByPageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCategoryByPageResponse
func (client *Client) QueryCategoryByPageWithOptions(request *QueryCategoryByPageRequest, headers *QueryCategoryByPageHeaders, runtime *util.RuntimeOptions) (_result *QueryCategoryByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
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
		Action:      tea.String("QueryCategoryByPage"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/categories/batch"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCategoryByPageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量获取费用类别
//
// @param request - QueryCategoryByPageRequest
//
// @return QueryCategoryByPageResponse
func (client *Client) QueryCategoryByPage(request *QueryCategoryByPageRequest) (_result *QueryCategoryByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCategoryByPageHeaders{}
	_result = &QueryCategoryByPageResponse{}
	_body, _err := client.QueryCategoryByPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询进件信息
//
// @param request - QueryCollectionInfoListRequest
//
// @param headers - QueryCollectionInfoListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCollectionInfoListResponse
func (client *Client) QueryCollectionInfoListWithOptions(request *QueryCollectionInfoListRequest, headers *QueryCollectionInfoListHeaders, runtime *util.RuntimeOptions) (_result *QueryCollectionInfoListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
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
		Action:      tea.String("QueryCollectionInfoList"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/me/collections/accounts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCollectionInfoListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询进件信息
//
// @param request - QueryCollectionInfoListRequest
//
// @return QueryCollectionInfoListResponse
func (client *Client) QueryCollectionInfoList(request *QueryCollectionInfoListRequest) (_result *QueryCollectionInfoListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCollectionInfoListHeaders{}
	_result = &QueryCollectionInfoListResponse{}
	_body, _err := client.QueryCollectionInfoListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询收款订单
//
// @param request - QueryCollectionOrderRequest
//
// @param headers - QueryCollectionOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCollectionOrderResponse
func (client *Client) QueryCollectionOrderWithOptions(request *QueryCollectionOrderRequest, headers *QueryCollectionOrderHeaders, runtime *util.RuntimeOptions) (_result *QueryCollectionOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
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
		Action:      tea.String("QueryCollectionOrder"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/me/collections/orders"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCollectionOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询收款订单
//
// @param request - QueryCollectionOrderRequest
//
// @return QueryCollectionOrderResponse
func (client *Client) QueryCollectionOrder(request *QueryCollectionOrderRequest) (_result *QueryCollectionOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCollectionOrderHeaders{}
	_result = &QueryCollectionOrderResponse{}
	_body, _err := client.QueryCollectionOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页批量获取智能财务应用内维护的客户信息
//
// @param request - QueryCustomerByPageRequest
//
// @param headers - QueryCustomerByPageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCustomerByPageResponse
func (client *Client) QueryCustomerByPageWithOptions(request *QueryCustomerByPageRequest, headers *QueryCustomerByPageHeaders, runtime *util.RuntimeOptions) (_result *QueryCustomerByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      tea.String("QueryCustomerByPage"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/customers/batch"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCustomerByPageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页批量获取智能财务应用内维护的客户信息
//
// @param request - QueryCustomerByPageRequest
//
// @return QueryCustomerByPageResponse
func (client *Client) QueryCustomerByPage(request *QueryCustomerByPageRequest) (_result *QueryCustomerByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCustomerByPageHeaders{}
	_result = &QueryCustomerByPageResponse{}
	_body, _err := client.QueryCustomerByPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量获取企业账户
//
// @param request - QueryEnterpriseAccountByPageRequest
//
// @param headers - QueryEnterpriseAccountByPageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEnterpriseAccountByPageResponse
func (client *Client) QueryEnterpriseAccountByPageWithOptions(request *QueryEnterpriseAccountByPageRequest, headers *QueryEnterpriseAccountByPageHeaders, runtime *util.RuntimeOptions) (_result *QueryEnterpriseAccountByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      tea.String("QueryEnterpriseAccountByPage"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/financeAccounts/batch"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryEnterpriseAccountByPageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量获取企业账户
//
// @param request - QueryEnterpriseAccountByPageRequest
//
// @return QueryEnterpriseAccountByPageResponse
func (client *Client) QueryEnterpriseAccountByPage(request *QueryEnterpriseAccountByPageRequest) (_result *QueryEnterpriseAccountByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryEnterpriseAccountByPageHeaders{}
	_result = &QueryEnterpriseAccountByPageResponse{}
	_body, _err := client.QueryEnterpriseAccountByPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取智能财务企业账户签约地址
//
// @param request - QueryEnterpriseAccountSignUrlRequest
//
// @param headers - QueryEnterpriseAccountSignUrlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEnterpriseAccountSignUrlResponse
func (client *Client) QueryEnterpriseAccountSignUrlWithOptions(request *QueryEnterpriseAccountSignUrlRequest, headers *QueryEnterpriseAccountSignUrlHeaders, runtime *util.RuntimeOptions) (_result *QueryEnterpriseAccountSignUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountCode)) {
		query["accountCode"] = request.AccountCode
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
		Action:      tea.String("QueryEnterpriseAccountSignUrl"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/enterpriseAccounts/sign"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryEnterpriseAccountSignUrlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取智能财务企业账户签约地址
//
// @param request - QueryEnterpriseAccountSignUrlRequest
//
// @return QueryEnterpriseAccountSignUrlResponse
func (client *Client) QueryEnterpriseAccountSignUrl(request *QueryEnterpriseAccountSignUrlRequest) (_result *QueryEnterpriseAccountSignUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryEnterpriseAccountSignUrlHeaders{}
	_result = &QueryEnterpriseAccountSignUrlResponse{}
	_body, _err := client.QueryEnterpriseAccountSignUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询支付订单详情
//
// @param request - QueryInstancePaymentOrderDetailRequest
//
// @param headers - QueryInstancePaymentOrderDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInstancePaymentOrderDetailResponse
func (client *Client) QueryInstancePaymentOrderDetailWithOptions(instanceId *string, request *QueryInstancePaymentOrderDetailRequest, headers *QueryInstancePaymentOrderDetailHeaders, runtime *util.RuntimeOptions) (_result *QueryInstancePaymentOrderDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderNo)) {
		query["orderNo"] = request.OrderNo
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
		Action:      tea.String("QueryInstancePaymentOrderDetail"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/instances/" + tea.StringValue(instanceId) + "/paymentOrders/details"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryInstancePaymentOrderDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询支付订单详情
//
// @param request - QueryInstancePaymentOrderDetailRequest
//
// @return QueryInstancePaymentOrderDetailResponse
func (client *Client) QueryInstancePaymentOrderDetail(instanceId *string, request *QueryInstancePaymentOrderDetailRequest) (_result *QueryInstancePaymentOrderDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryInstancePaymentOrderDetailHeaders{}
	_result = &QueryInstancePaymentOrderDetailResponse{}
	_body, _err := client.QueryInstancePaymentOrderDetailWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发票数据迁移，根据数据key查询具体数据data
//
// @param tmpReq - QueryInvoiceTransferDataRequest
//
// @param headers - QueryInvoiceTransferDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInvoiceTransferDataResponse
func (client *Client) QueryInvoiceTransferDataWithOptions(tmpReq *QueryInvoiceTransferDataRequest, headers *QueryInvoiceTransferDataHeaders, runtime *util.RuntimeOptions) (_result *QueryInvoiceTransferDataResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryInvoiceTransferDataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
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
		Action:      tea.String("QueryInvoiceTransferData"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/invoices/transferredDatas/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryInvoiceTransferDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发票数据迁移，根据数据key查询具体数据data
//
// @param request - QueryInvoiceTransferDataRequest
//
// @return QueryInvoiceTransferDataResponse
func (client *Client) QueryInvoiceTransferData(request *QueryInvoiceTransferDataRequest) (_result *QueryInvoiceTransferDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryInvoiceTransferDataHeaders{}
	_result = &QueryInvoiceTransferDataResponse{}
	_body, _err := client.QueryInvoiceTransferDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询支付的权益使用信息
//
// @param request - QueryPaymentBenefitRequest
//
// @param headers - QueryPaymentBenefitHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPaymentBenefitResponse
func (client *Client) QueryPaymentBenefitWithOptions(request *QueryPaymentBenefitRequest, headers *QueryPaymentBenefitHeaders, runtime *util.RuntimeOptions) (_result *QueryPaymentBenefitResponse, _err error) {
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
		Action:      tea.String("QueryPaymentBenefit"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/payments/benefits"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPaymentBenefitResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询支付的权益使用信息
//
// @param request - QueryPaymentBenefitRequest
//
// @return QueryPaymentBenefitResponse
func (client *Client) QueryPaymentBenefit(request *QueryPaymentBenefitRequest) (_result *QueryPaymentBenefitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryPaymentBenefitHeaders{}
	_result = &QueryPaymentBenefitResponse{}
	_body, _err := client.QueryPaymentBenefitWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询智能财务支付订单信息
//
// @param tmpReq - QueryPaymentOrderDetailRequest
//
// @param headers - QueryPaymentOrderDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPaymentOrderDetailResponse
func (client *Client) QueryPaymentOrderDetailWithOptions(tmpReq *QueryPaymentOrderDetailRequest, headers *QueryPaymentOrderDetailHeaders, runtime *util.RuntimeOptions) (_result *QueryPaymentOrderDetailResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryPaymentOrderDetailShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.OutBizNoList)) {
		request.OutBizNoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OutBizNoList, tea.String("outBizNoList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutBizNoListShrink)) {
		query["outBizNoList"] = request.OutBizNoListShrink
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
		Action:      tea.String("QueryPaymentOrderDetail"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/payments/orders"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPaymentOrderDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询智能财务支付订单信息
//
// @param request - QueryPaymentOrderDetailRequest
//
// @return QueryPaymentOrderDetailResponse
func (client *Client) QueryPaymentOrderDetail(request *QueryPaymentOrderDetailRequest) (_result *QueryPaymentOrderDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryPaymentOrderDetailHeaders{}
	_result = &QueryPaymentOrderDetailResponse{}
	_body, _err := client.QueryPaymentOrderDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询支付回单信息
//
// @param request - QueryPaymentRecallFileRequest
//
// @param headers - QueryPaymentRecallFileHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPaymentRecallFileResponse
func (client *Client) QueryPaymentRecallFileWithOptions(instanceId *string, request *QueryPaymentRecallFileRequest, headers *QueryPaymentRecallFileHeaders, runtime *util.RuntimeOptions) (_result *QueryPaymentRecallFileResponse, _err error) {
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
		Action:      tea.String("QueryPaymentRecallFile"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/payments/recallFiles/" + tea.StringValue(instanceId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPaymentRecallFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询支付回单信息
//
// @param request - QueryPaymentRecallFileRequest
//
// @return QueryPaymentRecallFileResponse
func (client *Client) QueryPaymentRecallFile(instanceId *string, request *QueryPaymentRecallFileRequest) (_result *QueryPaymentRecallFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryPaymentRecallFileHeaders{}
	_result = &QueryPaymentRecallFileResponse{}
	_body, _err := client.QueryPaymentRecallFileWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询支付订单的状态
//
// @param request - QueryPaymentStatusRequest
//
// @param headers - QueryPaymentStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPaymentStatusResponse
func (client *Client) QueryPaymentStatusWithOptions(request *QueryPaymentStatusRequest, headers *QueryPaymentStatusHeaders, runtime *util.RuntimeOptions) (_result *QueryPaymentStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNo)) {
		query["orderNo"] = request.OrderNo
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
		Action:      tea.String("QueryPaymentStatus"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/payments/statuses"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPaymentStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询支付订单的状态
//
// @param request - QueryPaymentStatusRequest
//
// @return QueryPaymentStatusResponse
func (client *Client) QueryPaymentStatus(request *QueryPaymentStatusRequest) (_result *QueryPaymentStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryPaymentStatusHeaders{}
	_result = &QueryPaymentStatusResponse{}
	_body, _err := client.QueryPaymentStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量获取商品信息
//
// @param request - QueryProductByPageRequest
//
// @param headers - QueryProductByPageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryProductByPageResponse
func (client *Client) QueryProductByPageWithOptions(request *QueryProductByPageRequest, headers *QueryProductByPageHeaders, runtime *util.RuntimeOptions) (_result *QueryProductByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      tea.String("QueryProductByPage"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/products/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryProductByPageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量获取商品信息
//
// @param request - QueryProductByPageRequest
//
// @return QueryProductByPageResponse
func (client *Client) QueryProductByPage(request *QueryProductByPageRequest) (_result *QueryProductByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryProductByPageHeaders{}
	_result = &QueryProductByPageResponse{}
	_body, _err := client.QueryProductByPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量获取项目信息
//
// @param request - QueryProjectByPageRequest
//
// @param headers - QueryProjectByPageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryProjectByPageResponse
func (client *Client) QueryProjectByPageWithOptions(request *QueryProjectByPageRequest, headers *QueryProjectByPageHeaders, runtime *util.RuntimeOptions) (_result *QueryProjectByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      tea.String("QueryProjectByPage"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/projects/batch"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryProjectByPageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量获取项目信息
//
// @param request - QueryProjectByPageRequest
//
// @return QueryProjectByPageResponse
func (client *Client) QueryProjectByPage(request *QueryProjectByPageRequest) (_result *QueryProjectByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryProjectByPageHeaders{}
	_result = &QueryProjectByPageResponse{}
	_body, _err := client.QueryProjectByPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询智能财务单据主数据信息
//
// @param request - QueryReceiptForInvoiceRequest
//
// @param headers - QueryReceiptForInvoiceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryReceiptForInvoiceResponse
func (client *Client) QueryReceiptForInvoiceWithOptions(request *QueryReceiptForInvoiceRequest, headers *QueryReceiptForInvoiceHeaders, runtime *util.RuntimeOptions) (_result *QueryReceiptForInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountantBookId)) {
		body["accountantBookId"] = request.AccountantBookId
	}

	if !tea.BoolValue(util.IsUnset(request.ApplyStatusList)) {
		body["applyStatusList"] = request.ApplyStatusList
	}

	if !tea.BoolValue(util.IsUnset(request.BizStatusList)) {
		body["bizStatusList"] = request.BizStatusList
	}

	if !tea.BoolValue(util.IsUnset(request.CompanyCode)) {
		body["companyCode"] = request.CompanyCode
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiptStatusList)) {
		body["receiptStatusList"] = request.ReceiptStatusList
	}

	if !tea.BoolValue(util.IsUnset(request.SearchParams)) {
		body["searchParams"] = request.SearchParams
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
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
		Action:      tea.String("QueryReceiptForInvoice"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/invoices/receipts/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryReceiptForInvoiceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询智能财务单据主数据信息
//
// @param request - QueryReceiptForInvoiceRequest
//
// @return QueryReceiptForInvoiceResponse
func (client *Client) QueryReceiptForInvoice(request *QueryReceiptForInvoiceRequest) (_result *QueryReceiptForInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryReceiptForInvoiceHeaders{}
	_result = &QueryReceiptForInvoiceResponse{}
	_body, _err := client.QueryReceiptForInvoiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页批量获取智能财务应用内维护的供应商信息
//
// @param request - QuerySupplierByPageRequest
//
// @param headers - QuerySupplierByPageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySupplierByPageResponse
func (client *Client) QuerySupplierByPageWithOptions(request *QuerySupplierByPageRequest, headers *QuerySupplierByPageHeaders, runtime *util.RuntimeOptions) (_result *QuerySupplierByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      tea.String("QuerySupplierByPage"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/suppliers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySupplierByPageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页批量获取智能财务应用内维护的供应商信息
//
// @param request - QuerySupplierByPageRequest
//
// @return QuerySupplierByPageResponse
func (client *Client) QuerySupplierByPage(request *QuerySupplierByPageRequest) (_result *QuerySupplierByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QuerySupplierByPageHeaders{}
	_result = &QuerySupplierByPageResponse{}
	_body, _err := client.QuerySupplierByPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询组织是否命中走新发票应用
//
// @param headers - QueryUseNewInvoiceAppHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUseNewInvoiceAppResponse
func (client *Client) QueryUseNewInvoiceAppWithOptions(headers *QueryUseNewInvoiceAppHeaders, runtime *util.RuntimeOptions) (_result *QueryUseNewInvoiceAppResponse, _err error) {
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
		Action:      tea.String("QueryUseNewInvoiceApp"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/invoice/appGray"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUseNewInvoiceAppResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询组织是否命中走新发票应用
//
// @return QueryUseNewInvoiceAppResponse
func (client *Client) QueryUseNewInvoiceApp() (_result *QueryUseNewInvoiceAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUseNewInvoiceAppHeaders{}
	_result = &QueryUseNewInvoiceAppResponse{}
	_body, _err := client.QueryUseNewInvoiceAppWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户角色成员，支持分页，可获取某个企业主体下的角色成员
//
// @param request - QueryUserRoleListRequest
//
// @param headers - QueryUserRoleListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserRoleListResponse
func (client *Client) QueryUserRoleListWithOptions(request *QueryUserRoleListRequest, headers *QueryUserRoleListHeaders, runtime *util.RuntimeOptions) (_result *QueryUserRoleListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompanyCode)) {
		query["companyCode"] = request.CompanyCode
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
		Action:      tea.String("QueryUserRoleList"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/users/roles"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserRoleListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户角色成员，支持分页，可获取某个企业主体下的角色成员
//
// @param request - QueryUserRoleListRequest
//
// @return QueryUserRoleListResponse
func (client *Client) QueryUserRoleList(request *QueryUserRoleListRequest) (_result *QueryUserRoleListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserRoleListHeaders{}
	_result = &QueryUserRoleListResponse{}
	_body, _err := client.QueryUserRoleListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 签约企业账户
//
// @param request - SignEnterpriseAccountRequest
//
// @param headers - SignEnterpriseAccountHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SignEnterpriseAccountResponse
func (client *Client) SignEnterpriseAccountWithOptions(request *SignEnterpriseAccountRequest, headers *SignEnterpriseAccountHeaders, runtime *util.RuntimeOptions) (_result *SignEnterpriseAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountCode)) {
		query["accountCode"] = request.AccountCode
	}

	if !tea.BoolValue(util.IsUnset(request.BankCardNo)) {
		query["bankCardNo"] = request.BankCardNo
	}

	if !tea.BoolValue(util.IsUnset(request.BankOpenId)) {
		query["bankOpenId"] = request.BankOpenId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelType)) {
		query["channelType"] = request.ChannelType
	}

	if !tea.BoolValue(util.IsUnset(request.FeeItemCode)) {
		query["feeItemCode"] = request.FeeItemCode
	}

	if !tea.BoolValue(util.IsUnset(request.IssueNo)) {
		query["issueNo"] = request.IssueNo
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		query["operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.SignOperateType)) {
		query["signOperateType"] = request.SignOperateType
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
		Action:      tea.String("SignEnterpriseAccount"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/enterpriseAccounts/sign"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SignEnterpriseAccountResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 签约企业账户
//
// @param request - SignEnterpriseAccountRequest
//
// @return SignEnterpriseAccountResponse
func (client *Client) SignEnterpriseAccount(request *SignEnterpriseAccountRequest) (_result *SignEnterpriseAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SignEnterpriseAccountHeaders{}
	_result = &SignEnterpriseAccountResponse{}
	_body, _err := client.SignEnterpriseAccountWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发送银企支付回单文件信息
//
// @param request - SyncReceiptRecallRequest
//
// @param headers - SyncReceiptRecallHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncReceiptRecallResponse
func (client *Client) SyncReceiptRecallWithOptions(request *SyncReceiptRecallRequest, headers *SyncReceiptRecallHeaders, runtime *util.RuntimeOptions) (_result *SyncReceiptRecallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileDownloadUrl)) {
		query["fileDownloadUrl"] = request.FileDownloadUrl
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNo)) {
		query["orderNo"] = request.OrderNo
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
		Action:      tea.String("SyncReceiptRecall"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/receipts/syncRecall"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncReceiptRecallResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送银企支付回单文件信息
//
// @param request - SyncReceiptRecallRequest
//
// @return SyncReceiptRecallResponse
func (client *Client) SyncReceiptRecall(request *SyncReceiptRecallRequest) (_result *SyncReceiptRecallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncReceiptRecallHeaders{}
	_result = &SyncReceiptRecallResponse{}
	_body, _err := client.SyncReceiptRecallWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新智能财务档案的状态
//
// @param request - UpdateAuxiliaryStatusRequest
//
// @param headers - UpdateAuxiliaryStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAuxiliaryStatusResponse
func (client *Client) UpdateAuxiliaryStatusWithOptions(request *UpdateAuxiliaryStatusRequest, headers *UpdateAuxiliaryStatusHeaders, runtime *util.RuntimeOptions) (_result *UpdateAuxiliaryStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuxiliaryId)) {
		query["auxiliaryId"] = request.AuxiliaryId
	}

	if !tea.BoolValue(util.IsUnset(request.AuxiliaryType)) {
		query["auxiliaryType"] = request.AuxiliaryType
	}

	if !tea.BoolValue(util.IsUnset(request.OperateType)) {
		query["operateType"] = request.OperateType
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
		Action:      tea.String("UpdateAuxiliaryStatus"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/auxiliaries/status"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAuxiliaryStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新智能财务档案的状态
//
// @param request - UpdateAuxiliaryStatusRequest
//
// @return UpdateAuxiliaryStatusResponse
func (client *Client) UpdateAuxiliaryStatus(request *UpdateAuxiliaryStatusRequest) (_result *UpdateAuxiliaryStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateAuxiliaryStatusHeaders{}
	_result = &UpdateAuxiliaryStatusResponse{}
	_body, _err := client.UpdateAuxiliaryStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新智能财务的企业账户
//
// @param request - UpdateFinanceEnterpriseAccountRequest
//
// @param headers - UpdateFinanceEnterpriseAccountHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFinanceEnterpriseAccountResponse
func (client *Client) UpdateFinanceEnterpriseAccountWithOptions(request *UpdateFinanceEnterpriseAccountRequest, headers *UpdateFinanceEnterpriseAccountHeaders, runtime *util.RuntimeOptions) (_result *UpdateFinanceEnterpriseAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountCode)) {
		body["accountCode"] = request.AccountCode
	}

	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["accountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountType)) {
		body["accountType"] = request.AccountType
	}

	if !tea.BoolValue(util.IsUnset(request.BankCardNo)) {
		body["bankCardNo"] = request.BankCardNo
	}

	if !tea.BoolValue(util.IsUnset(request.BankCode)) {
		body["bankCode"] = request.BankCode
	}

	if !tea.BoolValue(util.IsUnset(request.BankName)) {
		body["bankName"] = request.BankName
	}

	if !tea.BoolValue(util.IsUnset(request.City)) {
		body["city"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.OfficialName)) {
		body["officialName"] = request.OfficialName
	}

	if !tea.BoolValue(util.IsUnset(request.OfficialNumber)) {
		body["officialNumber"] = request.OfficialNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Province)) {
		body["province"] = request.Province
	}

	if !tea.BoolValue(util.IsUnset(request.TaxNature)) {
		body["taxNature"] = request.TaxNature
	}

	if !tea.BoolValue(util.IsUnset(request.TaxNo)) {
		body["taxNo"] = request.TaxNo
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
		Action:      tea.String("UpdateFinanceEnterpriseAccount"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/enterpriseAccounts"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFinanceEnterpriseAccountResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新智能财务的企业账户
//
// @param request - UpdateFinanceEnterpriseAccountRequest
//
// @return UpdateFinanceEnterpriseAccountResponse
func (client *Client) UpdateFinanceEnterpriseAccount(request *UpdateFinanceEnterpriseAccountRequest) (_result *UpdateFinanceEnterpriseAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateFinanceEnterpriseAccountHeaders{}
	_result = &UpdateFinanceEnterpriseAccountResponse{}
	_body, _err := client.UpdateFinanceEnterpriseAccountWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新单据的支付状态
//
// @param tmpReq - UpdateInstanceOrderInfoRequest
//
// @param headers - UpdateInstanceOrderInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceOrderInfoResponse
func (client *Client) UpdateInstanceOrderInfoWithOptions(instanceId *string, tmpReq *UpdateInstanceOrderInfoRequest, headers *UpdateInstanceOrderInfoHeaders, runtime *util.RuntimeOptions) (_result *UpdateInstanceOrderInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateInstanceOrderInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.PayerBank)) {
		request.PayerBankShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PayerBank, tea.String("payerBank"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FailReason)) {
		query["failReason"] = request.FailReason
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNo)) {
		query["orderNo"] = request.OrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.OutOrderNo)) {
		query["outOrderNo"] = request.OutOrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.PayerBankShrink)) {
		query["payerBank"] = request.PayerBankShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentTime)) {
		query["paymentTime"] = request.PaymentTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
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
		Action:      tea.String("UpdateInstanceOrderInfo"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/instances/" + tea.StringValue(instanceId) + "/paymentOrders/states"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceOrderInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新单据的支付状态
//
// @param request - UpdateInstanceOrderInfoRequest
//
// @return UpdateInstanceOrderInfoResponse
func (client *Client) UpdateInstanceOrderInfo(instanceId *string, request *UpdateInstanceOrderInfoRequest) (_result *UpdateInstanceOrderInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateInstanceOrderInfoHeaders{}
	_result = &UpdateInstanceOrderInfoResponse{}
	_body, _err := client.UpdateInstanceOrderInfoWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发票数据迁移，新发票应用上报已成功搬移数据
//
// @param headers - UpdateInvoiceDataTransferDoneHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInvoiceDataTransferDoneResponse
func (client *Client) UpdateInvoiceDataTransferDoneWithOptions(headers *UpdateInvoiceDataTransferDoneHeaders, runtime *util.RuntimeOptions) (_result *UpdateInvoiceDataTransferDoneResponse, _err error) {
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
		Action:      tea.String("UpdateInvoiceDataTransferDone"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/invoices/transferredDatas/statuses"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInvoiceDataTransferDoneResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发票数据迁移，新发票应用上报已成功搬移数据
//
// @return UpdateInvoiceDataTransferDoneResponse
func (client *Client) UpdateInvoiceDataTransferDone() (_result *UpdateInvoiceDataTransferDoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateInvoiceDataTransferDoneHeaders{}
	_result = &UpdateInvoiceDataTransferDoneResponse{}
	_body, _err := client.UpdateInvoiceDataTransferDoneWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用于更新智能财务企业票池内对应发票的下载链接
//
// @param tmpReq - UpdateInvoiceUrlRequest
//
// @param headers - UpdateInvoiceUrlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInvoiceUrlResponse
func (client *Client) UpdateInvoiceUrlWithOptions(tmpReq *UpdateInvoiceUrlRequest, headers *UpdateInvoiceUrlHeaders, runtime *util.RuntimeOptions) (_result *UpdateInvoiceUrlResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateInvoiceUrlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
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
		Action:      tea.String("UpdateInvoiceUrl"),
		Version:     tea.String("bizfinance_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/bizfinance/invoices/urls"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInvoiceUrlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用于更新智能财务企业票池内对应发票的下载链接
//
// @param request - UpdateInvoiceUrlRequest
//
// @return UpdateInvoiceUrlResponse
func (client *Client) UpdateInvoiceUrl(request *UpdateInvoiceUrlRequest) (_result *UpdateInvoiceUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateInvoiceUrlHeaders{}
	_result = &UpdateInvoiceUrlResponse{}
	_body, _err := client.UpdateInvoiceUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
