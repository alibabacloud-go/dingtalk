// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package finance_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
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
	// 支付账号唯一id
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// 钉钉订单号(和商户批次号一一对应)
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	// 公用回传参数，如果请求时传递了该参数，则异步通知商户时会回传该参数
	PassBackParams map[string]interface{} `json:"passBackParams,omitempty" xml:"passBackParams,omitempty"`
	// 支付终端
	PayTerminal *string `json:"payTerminal,omitempty" xml:"payTerminal,omitempty"`
	// 回调url
	ReturnUrl *string `json:"returnUrl,omitempty" xml:"returnUrl,omitempty"`
	// 支付发起人staffId
	StaffId *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
	// 订单总金额（必填）, 单位为：元
	TransAmount *string `json:"transAmount,omitempty" xml:"transAmount,omitempty"`
	// 转账过期时间
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
	// 钉钉支付的批次号
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	// 支付确认页数据
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApplyBatchPayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ApplyBatchPayResponse) SetBody(v *ApplyBatchPayResponseBody) *ApplyBatchPayResponse {
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
	// 签约支付宝账户，用于协议确认
	BindingAlipayLogonId *string `json:"bindingAlipayLogonId,omitempty" xml:"bindingAlipayLogonId,omitempty"`
	// 联系人
	ContractInfo *ConsultCreateSubInstitutionRequestContractInfo `json:"contractInfo,omitempty" xml:"contractInfo,omitempty" type:"Struct"`
	// 主机构编号
	InstId              *string                                                `json:"instId,omitempty" xml:"instId,omitempty"`
	LegalPersonCertInfo *ConsultCreateSubInstitutionRequestLegalPersonCertInfo `json:"legalPersonCertInfo,omitempty" xml:"legalPersonCertInfo,omitempty" type:"Struct"`
	// 进件创建外部流水号
	OutTradeNo *string `json:"outTradeNo,omitempty" xml:"outTradeNo,omitempty"`
	// 进件渠道
	PayChannel *string `json:"payChannel,omitempty" xml:"payChannel,omitempty"`
	// 资质信息
	QualificationInfos []*ConsultCreateSubInstitutionRequestQualificationInfos `json:"qualificationInfos,omitempty" xml:"qualificationInfos,omitempty" type:"Repeated"`
	// 开通的服务类型
	Services []*string `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
	// 资金账户信息
	SettleInfo *ConsultCreateSubInstitutionRequestSettleInfo `json:"settleInfo,omitempty" xml:"settleInfo,omitempty" type:"Struct"`
	// 解决方案，包含清算、费率规则
	Solution *string `json:"solution,omitempty" xml:"solution,omitempty"`
	// 子机构地址信息
	SubInstAddressInfo *ConsultCreateSubInstitutionRequestSubInstAddressInfo `json:"subInstAddressInfo,omitempty" xml:"subInstAddressInfo,omitempty" type:"Struct"`
	// 授权信息
	SubInstAuthInfo *ConsultCreateSubInstitutionRequestSubInstAuthInfo `json:"subInstAuthInfo,omitempty" xml:"subInstAuthInfo,omitempty" type:"Struct"`
	// 子机构基本信息
	SubInstBasicInfo *ConsultCreateSubInstitutionRequestSubInstBasicInfo `json:"subInstBasicInfo,omitempty" xml:"subInstBasicInfo,omitempty" type:"Struct"`
	// 子机构认证信息
	SubInstCertifyInfo *ConsultCreateSubInstitutionRequestSubInstCertifyInfo `json:"subInstCertifyInfo,omitempty" xml:"subInstCertifyInfo,omitempty" type:"Struct"`
	// 子机构编号
	SubInstId *string `json:"subInstId,omitempty" xml:"subInstId,omitempty"`
	// 开票信息
	SubInstInvoiceInfo *ConsultCreateSubInstitutionRequestSubInstInvoiceInfo `json:"subInstInvoiceInfo,omitempty" xml:"subInstInvoiceInfo,omitempty" type:"Struct"`
	// 子机构门店信息
	SubInstShopInfo *ConsultCreateSubInstitutionRequestSubInstShopInfo `json:"subInstShopInfo,omitempty" xml:"subInstShopInfo,omitempty" type:"Struct"`
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

func (s *ConsultCreateSubInstitutionRequest) SetContractInfo(v *ConsultCreateSubInstitutionRequestContractInfo) *ConsultCreateSubInstitutionRequest {
	s.ContractInfo = v
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

type ConsultCreateSubInstitutionRequestContractInfo struct {
	// 联系人姓名
	ContractName *string `json:"contractName,omitempty" xml:"contractName,omitempty"`
	// 联系人手机号
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
}

func (s ConsultCreateSubInstitutionRequestContractInfo) String() string {
	return tea.Prettify(s)
}

func (s ConsultCreateSubInstitutionRequestContractInfo) GoString() string {
	return s.String()
}

func (s *ConsultCreateSubInstitutionRequestContractInfo) SetContractName(v string) *ConsultCreateSubInstitutionRequestContractInfo {
	s.ContractName = &v
	return s
}

func (s *ConsultCreateSubInstitutionRequestContractInfo) SetMobile(v string) *ConsultCreateSubInstitutionRequestContractInfo {
	s.Mobile = &v
	return s
}

type ConsultCreateSubInstitutionRequestLegalPersonCertInfo struct {
	// 法人证件反面url
	CertBackImage *string `json:"certBackImage,omitempty" xml:"certBackImage,omitempty"`
	// 法人证件正面url
	CertFrontImage *string `json:"certFrontImage,omitempty" xml:"certFrontImage,omitempty"`
	// 法人姓名
	CertName *string `json:"certName,omitempty" xml:"certName,omitempty"`
	// 法人证件类型 不填默认为身份证
	CertType *string `json:"certType,omitempty" xml:"certType,omitempty"`
	// 法人证件号
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
	// 子机构行业资质图片
	QualificationImage *string `json:"qualificationImage,omitempty" xml:"qualificationImage,omitempty"`
	// 子机构行业资质类型
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
	// 账户账号
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// 账户名称 账号类型银行卡时为卡户名
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	// 卡类型
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	// 支行名称
	BankBranchName *string `json:"bankBranchName,omitempty" xml:"bankBranchName,omitempty"`
	// 开户行所在地 市
	BankCity *string `json:"bankCity,omitempty" xml:"bankCity,omitempty"`
	// 联行号
	BankCode *string `json:"bankCode,omitempty" xml:"bankCode,omitempty"`
	// 银行名称
	BankName *string `json:"bankName,omitempty" xml:"bankName,omitempty"`
	// 开户行所在地 省
	BankProvince *string `json:"bankProvince,omitempty" xml:"bankProvince,omitempty"`
	// 开户行简称缩写
	BankShortNameCode *string `json:"bankShortNameCode,omitempty" xml:"bankShortNameCode,omitempty"`
	// 账号类型
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 账户使用类型
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
	// 详细地址
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// 市码
	CityCode *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	// 区码
	DistrictCode *string `json:"districtCode,omitempty" xml:"districtCode,omitempty"`
	// 省码
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
	// 授权函图片url
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
	// 别名
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// 机构识别码
	Mcc *string `json:"mcc,omitempty" xml:"mcc,omitempty"`
	// 名称
	SubInstName *string `json:"subInstName,omitempty" xml:"subInstName,omitempty"`
	// 类型
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
	// 证件图片, 如果是特殊行业必填
	CertImage *string `json:"certImage,omitempty" xml:"certImage,omitempty"`
	// 证件号码
	CertNo *string `json:"certNo,omitempty" xml:"certNo,omitempty"`
	// 证件类型
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
	// 是否接受电票
	AcceptElectronic *bool `json:"acceptElectronic,omitempty" xml:"acceptElectronic,omitempty"`
	// 开票地址
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// 是否自动开票
	AutoInvoice *bool `json:"autoInvoice,omitempty" xml:"autoInvoice,omitempty"`
	// 银行账户
	BankAccount *string `json:"bankAccount,omitempty" xml:"bankAccount,omitempty"`
	// 银行名称
	BankName *string `json:"bankName,omitempty" xml:"bankName,omitempty"`
	// 收件地址
	MailAddress *ConsultCreateSubInstitutionRequestSubInstInvoiceInfoMailAddress `json:"mailAddress,omitempty" xml:"mailAddress,omitempty" type:"Struct"`
	// 收件人名称
	MailName *string `json:"mailName,omitempty" xml:"mailName,omitempty"`
	// 收件人号码
	MailPhone *string `json:"mailPhone,omitempty" xml:"mailPhone,omitempty"`
	// 纳税人识别号
	TaxNo *string `json:"taxNo,omitempty" xml:"taxNo,omitempty"`
	// 纳税人资质
	TaxPayerQualification *string `json:"taxPayerQualification,omitempty" xml:"taxPayerQualification,omitempty"`
	// 纳税人资格开始时间
	TaxPayerValidDate *string `json:"taxPayerValidDate,omitempty" xml:"taxPayerValidDate,omitempty"`
	// 开票电话
	Telephone *string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	// 纳税人抬头
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
	// 详细地址
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// 市码
	CityCode *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	// 区码
	DistrictCode *string `json:"districtCode,omitempty" xml:"districtCode,omitempty"`
	// 省码
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
	// 内景照
	InDoorImages []*string `json:"inDoorImages,omitempty" xml:"inDoorImages,omitempty" type:"Repeated"`
	// 外景照
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
	// 进件申请单号
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
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConsultCreateSubInstitutionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 扣款金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 主机构编号
	InstId *string `json:"instId,omitempty" xml:"instId,omitempty"`
	// 其他资金渠道付款明细
	OtherPayChannelDetailInfoList []*CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoList `json:"otherPayChannelDetailInfoList,omitempty" xml:"otherPayChannelDetailInfoList,omitempty" type:"Repeated"`
	// 外部订单号
	OutTradeNo *string `json:"outTradeNo,omitempty" xml:"outTradeNo,omitempty"`
	// 支付渠道
	PayChannel *string `json:"payChannel,omitempty" xml:"payChannel,omitempty"`
	// 付款人staffId
	PayerUserId *string `json:"payerUserId,omitempty" xml:"payerUserId,omitempty"`
	// 代扣备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 子机构编号
	SubInstId *string `json:"subInstId,omitempty" xml:"subInstId,omitempty"`
	// 代扣过期时间
	TimeOutExpress *string `json:"timeOutExpress,omitempty" xml:"timeOutExpress,omitempty"`
	// 代扣标题
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
	// 渠道金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 资金明细列表
	FundToolDetailInfoList []*CreatWithholdingOrderAndPayRequestOtherPayChannelDetailInfoListFundToolDetailInfoList `json:"fundToolDetailInfoList,omitempty" xml:"fundToolDetailInfoList,omitempty" type:"Repeated"`
	// 渠道名称
	PayChannelName *string `json:"payChannelName,omitempty" xml:"payChannelName,omitempty"`
	// 支付渠道单号
	PayChannelOrderNo *string `json:"payChannelOrderNo,omitempty" xml:"payChannelOrderNo,omitempty"`
	// 渠道类型
	PayChannelType *string `json:"payChannelType,omitempty" xml:"payChannelType,omitempty"`
	// 总优惠金额
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
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 扩展信息
	ExtInfo *string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// 资金工具名称
	FundToolName *string `json:"fundToolName,omitempty" xml:"fundToolName,omitempty"`
	// 资金明细创建时间
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 资金明细完成时间
	GmtFinish *string `json:"gmtFinish,omitempty" xml:"gmtFinish,omitempty"`
	// 是否是优惠工具
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
	// 代扣金额（元）
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 付款完成日期
	GmtPay *string `json:"gmtPay,omitempty" xml:"gmtPay,omitempty"`
	// 主机构编号
	InstId *string `json:"instId,omitempty" xml:"instId,omitempty"`
	// 钉钉订单号
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	// 外部订单号
	OutTradeNo *string `json:"outTradeNo,omitempty" xml:"outTradeNo,omitempty"`
	// 支付渠道
	PayChannel *string `json:"payChannel,omitempty" xml:"payChannel,omitempty"`
	// 支付渠道支付账号（脱敏后返回）
	PayChannelAccountNo *string `json:"payChannelAccountNo,omitempty" xml:"payChannelAccountNo,omitempty"`
	// 付款人staffId
	PayerStaffId *string `json:"payerStaffId,omitempty" xml:"payerStaffId,omitempty"`
	// 代扣备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 子机构编号
	SubInstId *string `json:"subInstId,omitempty" xml:"subInstId,omitempty"`
	// 代扣标题
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
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatWithholdingOrderAndPayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreatWithholdingOrderAndPayResponse) SetBody(v *CreatWithholdingOrderAndPayResponseBody) *CreatWithholdingOrderAndPayResponse {
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
	// 付款账号唯一id
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// 付款账号(注意：用户上送的是脱敏数据)
	AccountNo *string `json:"accountNo,omitempty" xml:"accountNo,omitempty"`
	// 批次备注
	BatchRemark *string `json:"batchRemark,omitempty" xml:"batchRemark,omitempty"`
	// 交易明细列表
	BatchTradeDetails []*CreateBatchTradeOrderRequestBatchTradeDetails `json:"batchTradeDetails,omitempty" xml:"batchTradeDetails,omitempty" type:"Repeated"`
	// 外部商户批次号
	OutBatchNo *string `json:"outBatchNo,omitempty" xml:"outBatchNo,omitempty"`
	// 员工staffId
	StaffId *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
	// 总金额（必填，单位：元）
	TotalAmount *string `json:"totalAmount,omitempty" xml:"totalAmount,omitempty"`
	// 总笔数（必填）
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 交易抬头
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
	// 金额（必填，单位：元）
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 备注（选填）
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 收款方户名（必填）
	PayeeAccountName *string `json:"payeeAccountName,omitempty" xml:"payeeAccountName,omitempty"`
	// 收款方账号（必填）
	PayeeAccountNo *string `json:"payeeAccountNo,omitempty" xml:"payeeAccountNo,omitempty"`
	// 收款方账号类型（必填）
	PayeeAccountType *string `json:"payeeAccountType,omitempty" xml:"payeeAccountType,omitempty"`
	// 序号（必填）
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
	// 钉钉批次单号
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	// 批次订单状态
	OrderStatus *string `json:"orderStatus,omitempty" xml:"orderStatus,omitempty"`
	// 商户批次号
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateBatchTradeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 签约支付宝账户，用于协议确认
	BindingAlipayLogonId *string `json:"bindingAlipayLogonId,omitempty" xml:"bindingAlipayLogonId,omitempty"`
	// 联系人
	ContractInfo *CreateSubInstitutionRequestContractInfo `json:"contractInfo,omitempty" xml:"contractInfo,omitempty" type:"Struct"`
	// 主机构编号
	InstId              *string                                         `json:"instId,omitempty" xml:"instId,omitempty"`
	LegalPersonCertInfo *CreateSubInstitutionRequestLegalPersonCertInfo `json:"legalPersonCertInfo,omitempty" xml:"legalPersonCertInfo,omitempty" type:"Struct"`
	// 进件创建外部流水号
	OutTradeNo *string `json:"outTradeNo,omitempty" xml:"outTradeNo,omitempty"`
	// 进件渠道
	PayChannel *string `json:"payChannel,omitempty" xml:"payChannel,omitempty"`
	// 资质信息
	QualificationInfos []*CreateSubInstitutionRequestQualificationInfos `json:"qualificationInfos,omitempty" xml:"qualificationInfos,omitempty" type:"Repeated"`
	// 开通的服务类型
	Services []*string `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
	// 资金账户信息
	SettleInfo *CreateSubInstitutionRequestSettleInfo `json:"settleInfo,omitempty" xml:"settleInfo,omitempty" type:"Struct"`
	// 解决方案，包含费率、清算规则等
	Solution *string `json:"solution,omitempty" xml:"solution,omitempty"`
	// 子机构地址信息
	SubInstAddressInfo *CreateSubInstitutionRequestSubInstAddressInfo `json:"subInstAddressInfo,omitempty" xml:"subInstAddressInfo,omitempty" type:"Struct"`
	// 授权信息
	SubInstAuthInfo *CreateSubInstitutionRequestSubInstAuthInfo `json:"subInstAuthInfo,omitempty" xml:"subInstAuthInfo,omitempty" type:"Struct"`
	// 子机构基本信息
	SubInstBasicInfo *CreateSubInstitutionRequestSubInstBasicInfo `json:"subInstBasicInfo,omitempty" xml:"subInstBasicInfo,omitempty" type:"Struct"`
	// 子机构认证信息
	SubInstCertifyInfo *CreateSubInstitutionRequestSubInstCertifyInfo `json:"subInstCertifyInfo,omitempty" xml:"subInstCertifyInfo,omitempty" type:"Struct"`
	// 子机构编号
	SubInstId *string `json:"subInstId,omitempty" xml:"subInstId,omitempty"`
	// 开票信息
	SubInstInvoiceInfo *CreateSubInstitutionRequestSubInstInvoiceInfo `json:"subInstInvoiceInfo,omitempty" xml:"subInstInvoiceInfo,omitempty" type:"Struct"`
	// 子机构门店信息
	SubInstShopInfo *CreateSubInstitutionRequestSubInstShopInfo `json:"subInstShopInfo,omitempty" xml:"subInstShopInfo,omitempty" type:"Struct"`
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

func (s *CreateSubInstitutionRequest) SetContractInfo(v *CreateSubInstitutionRequestContractInfo) *CreateSubInstitutionRequest {
	s.ContractInfo = v
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

type CreateSubInstitutionRequestContractInfo struct {
	// 联系人姓名
	ContractName *string `json:"contractName,omitempty" xml:"contractName,omitempty"`
	// 联系人手机号
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
}

func (s CreateSubInstitutionRequestContractInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateSubInstitutionRequestContractInfo) GoString() string {
	return s.String()
}

func (s *CreateSubInstitutionRequestContractInfo) SetContractName(v string) *CreateSubInstitutionRequestContractInfo {
	s.ContractName = &v
	return s
}

func (s *CreateSubInstitutionRequestContractInfo) SetMobile(v string) *CreateSubInstitutionRequestContractInfo {
	s.Mobile = &v
	return s
}

type CreateSubInstitutionRequestLegalPersonCertInfo struct {
	// 法人证件反面url
	CertBackImage *string `json:"certBackImage,omitempty" xml:"certBackImage,omitempty"`
	// 法人证件正面url
	CertFrontImage *string `json:"certFrontImage,omitempty" xml:"certFrontImage,omitempty"`
	// 法人姓名
	CertName *string `json:"certName,omitempty" xml:"certName,omitempty"`
	// 法人证件类型 不填默认为身份证
	CertType *string `json:"certType,omitempty" xml:"certType,omitempty"`
	// 法人证件号
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
	// 子机构行业资质图片
	QualificationImage *string `json:"qualificationImage,omitempty" xml:"qualificationImage,omitempty"`
	// 子机构行业资质类型
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
	// 账户账号
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// 账户名称 账号类型银行卡时为卡户名
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	// 卡类型, DEBIT_CARD借记卡，CREDIT_CARD信用卡
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	// 支行名称
	BankBranchName *string `json:"bankBranchName,omitempty" xml:"bankBranchName,omitempty"`
	// 开户行所在地 市
	BankCity *string `json:"bankCity,omitempty" xml:"bankCity,omitempty"`
	// 联行号
	BankCode *string `json:"bankCode,omitempty" xml:"bankCode,omitempty"`
	// 银行名称
	BankName *string `json:"bankName,omitempty" xml:"bankName,omitempty"`
	// 开户行所在地 省
	BankProvince *string `json:"bankProvince,omitempty" xml:"bankProvince,omitempty"`
	// 开户行简称缩写
	BankShortNameCode *string `json:"bankShortNameCode,omitempty" xml:"bankShortNameCode,omitempty"`
	// 账号类型
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 账户使用类型
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
	// 详细地址
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// 市码
	CityCode *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	// 区码
	DistrictCode *string `json:"districtCode,omitempty" xml:"districtCode,omitempty"`
	// 省码
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
	// 授权函图片url
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
	// 别名
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// 机构识别码
	Mcc *string `json:"mcc,omitempty" xml:"mcc,omitempty"`
	// 名称
	SubInstName *string `json:"subInstName,omitempty" xml:"subInstName,omitempty"`
	// 类型
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
	// 证件图片, 如果是特殊行业必填
	CertImage *string `json:"certImage,omitempty" xml:"certImage,omitempty"`
	// 证件号码
	CertNo *string `json:"certNo,omitempty" xml:"certNo,omitempty"`
	// 证件类型
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
	// 是否接受电票
	AcceptElectronic *bool `json:"acceptElectronic,omitempty" xml:"acceptElectronic,omitempty"`
	// 开票地址
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// 是否自动开票
	AutoInvoice *bool `json:"autoInvoice,omitempty" xml:"autoInvoice,omitempty"`
	// 银行账户
	BankAccount *string `json:"bankAccount,omitempty" xml:"bankAccount,omitempty"`
	// 银行名称
	BankName *string `json:"bankName,omitempty" xml:"bankName,omitempty"`
	// 收件地址
	MailAddress *CreateSubInstitutionRequestSubInstInvoiceInfoMailAddress `json:"mailAddress,omitempty" xml:"mailAddress,omitempty" type:"Struct"`
	// 收件人名称
	MailName *string `json:"mailName,omitempty" xml:"mailName,omitempty"`
	// 收件人号码
	MailPhone *string `json:"mailPhone,omitempty" xml:"mailPhone,omitempty"`
	// 纳税人识别号
	TaxNo *string `json:"taxNo,omitempty" xml:"taxNo,omitempty"`
	// 纳税人资质
	TaxPayerQualification *string `json:"taxPayerQualification,omitempty" xml:"taxPayerQualification,omitempty"`
	// 纳税人资格开始时间
	TaxPayerValidDate *string `json:"taxPayerValidDate,omitempty" xml:"taxPayerValidDate,omitempty"`
	// 开票电话
	Telephone *string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	// 纳税人抬头
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
	// 详细地址
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// 市码
	CityCode *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	// 区码
	DistrictCode *string `json:"districtCode,omitempty" xml:"districtCode,omitempty"`
	// 省码
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
	// 内景照
	InDoorImages []*string `json:"inDoorImages,omitempty" xml:"inDoorImages,omitempty" type:"Repeated"`
	// 外景照
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
	// 进件申请单号
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSubInstitutionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 有效时间列表，对于连续时间段，只需传入一个对象即可，注意过期时间必须晚于最晚结束时间
	AvailableTimes []*CreateUserCodeInstanceRequestAvailableTimes `json:"availableTimes,omitempty" xml:"availableTimes,omitempty" type:"Repeated"`
	// 码标识，由钉钉颁发
	CodeIdentity *string `json:"codeIdentity,omitempty" xml:"codeIdentity,omitempty"`
	// 码值
	CodeValue *string `json:"codeValue,omitempty" xml:"codeValue,omitempty"`
	// 码值类型，钉钉静态码值：DING_STATIC，访客码或会展码传入
	CodeValueType *string `json:"codeValueType,omitempty" xml:"codeValueType,omitempty"`
	// 企业ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 扩展参数
	ExtInfo map[string]interface{} `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// 临时码，传入过期时间
	GmtExpired *string `json:"gmtExpired,omitempty" xml:"gmtExpired,omitempty"`
	// 业务幂等ID
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 状态，传入关闭状态需要用户手动开启后才会渲染二维
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 用户和企业的关系类型，区分内部员工，外部联系人，无关系普通用户
	UserCorpRelationType *string `json:"userCorpRelationType,omitempty" xml:"userCorpRelationType,omitempty"`
	// 用户身份标识，取值和用户企业关系类型相关，如果企业无关，传入手机号
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
	// 结束时间
	GmtEnd *string `json:"gmtEnd,omitempty" xml:"gmtEnd,omitempty"`
	// 开始时间
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
	// 码详情跳转地址
	CodeDetailUrl *string `json:"codeDetailUrl,omitempty" xml:"codeDetailUrl,omitempty"`
	// 码ID
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUserCodeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// payCode
	PayCode *string `json:"payCode,omitempty" xml:"payCode,omitempty"`
	// requestId
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
	// 支付宝付款码
	AlipayCode *string `json:"alipayCode,omitempty" xml:"alipayCode,omitempty"`
	// 码ID，对于访客或会展码等静态码值返回
	CodeId *string `json:"codeId,omitempty" xml:"codeId,omitempty"`
	// 工牌码：DT_IDENTITY，访客码：DT_VISITOR，会展码：DT_CONFERENCE
	CodeIdentity *string `json:"codeIdentity,omitempty" xml:"codeIdentity,omitempty"`
	// 码类型
	CodeType *string `json:"codeType,omitempty" xml:"codeType,omitempty"`
	// 企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 外部业务ID,其值为调用创建用户码接口传入的requestId
	OutBizId *string `json:"outBizId,omitempty" xml:"outBizId,omitempty"`
	// 用户和企业关系
	UserCorpRelationType *string `json:"userCorpRelationType,omitempty" xml:"userCorpRelationType,omitempty"`
	// 员工id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户是否还在组织内
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DecodePayCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DecodePayCodeResponse) SetBody(v *DecodePayCodeResponseBody) *DecodePayCodeResponse {
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
	// 签约支付宝账户，用于协议确认
	BindingAlipayLogonId *string `json:"bindingAlipayLogonId,omitempty" xml:"bindingAlipayLogonId,omitempty"`
	// 联系人
	ContractInfo *ModifySubInstitutionRequestContractInfo `json:"contractInfo,omitempty" xml:"contractInfo,omitempty" type:"Struct"`
	// 主机构编号
	InstId              *string                                         `json:"instId,omitempty" xml:"instId,omitempty"`
	LegalPersonCertInfo *ModifySubInstitutionRequestLegalPersonCertInfo `json:"legalPersonCertInfo,omitempty" xml:"legalPersonCertInfo,omitempty" type:"Struct"`
	// 进件创建外部流水号
	OutTradeNo *string `json:"outTradeNo,omitempty" xml:"outTradeNo,omitempty"`
	// 进件渠道
	PayChannel *string `json:"payChannel,omitempty" xml:"payChannel,omitempty"`
	// 资质信息
	QualificationInfos []*ModifySubInstitutionRequestQualificationInfos `json:"qualificationInfos,omitempty" xml:"qualificationInfos,omitempty" type:"Repeated"`
	// 开通的服务类型
	Services []*string `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
	// 资金账户信息
	SettleInfo *ModifySubInstitutionRequestSettleInfo `json:"settleInfo,omitempty" xml:"settleInfo,omitempty" type:"Struct"`
	// 子机构地址信息
	SubInstAddressInfo *ModifySubInstitutionRequestSubInstAddressInfo `json:"subInstAddressInfo,omitempty" xml:"subInstAddressInfo,omitempty" type:"Struct"`
	// 授权信息
	SubInstAuthInfo *ModifySubInstitutionRequestSubInstAuthInfo `json:"subInstAuthInfo,omitempty" xml:"subInstAuthInfo,omitempty" type:"Struct"`
	// 子机构基本信息
	SubInstBasicInfo *ModifySubInstitutionRequestSubInstBasicInfo `json:"subInstBasicInfo,omitempty" xml:"subInstBasicInfo,omitempty" type:"Struct"`
	// 子机构认证信息
	SubInstCertifyInfo *ModifySubInstitutionRequestSubInstCertifyInfo `json:"subInstCertifyInfo,omitempty" xml:"subInstCertifyInfo,omitempty" type:"Struct"`
	// 子机构编号
	SubInstId *string `json:"subInstId,omitempty" xml:"subInstId,omitempty"`
	// 开票信息
	SubInstInvoiceInfo *ModifySubInstitutionRequestSubInstInvoiceInfo `json:"subInstInvoiceInfo,omitempty" xml:"subInstInvoiceInfo,omitempty" type:"Struct"`
	// 子机构门店信息
	SubInstShopInfo *ModifySubInstitutionRequestSubInstShopInfo `json:"subInstShopInfo,omitempty" xml:"subInstShopInfo,omitempty" type:"Struct"`
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

func (s *ModifySubInstitutionRequest) SetContractInfo(v *ModifySubInstitutionRequestContractInfo) *ModifySubInstitutionRequest {
	s.ContractInfo = v
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

type ModifySubInstitutionRequestContractInfo struct {
	// 联系人姓名
	ContractName *string `json:"contractName,omitempty" xml:"contractName,omitempty"`
	// 联系人手机号
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
}

func (s ModifySubInstitutionRequestContractInfo) String() string {
	return tea.Prettify(s)
}

func (s ModifySubInstitutionRequestContractInfo) GoString() string {
	return s.String()
}

func (s *ModifySubInstitutionRequestContractInfo) SetContractName(v string) *ModifySubInstitutionRequestContractInfo {
	s.ContractName = &v
	return s
}

func (s *ModifySubInstitutionRequestContractInfo) SetMobile(v string) *ModifySubInstitutionRequestContractInfo {
	s.Mobile = &v
	return s
}

type ModifySubInstitutionRequestLegalPersonCertInfo struct {
	// 法人证件反面url
	CertBackImage *string `json:"certBackImage,omitempty" xml:"certBackImage,omitempty"`
	// 法人证件正面url
	CertFrontImage *string `json:"certFrontImage,omitempty" xml:"certFrontImage,omitempty"`
	// 法人姓名
	CertName *string `json:"certName,omitempty" xml:"certName,omitempty"`
	// 法人证件类型 不填默认为身份证
	CertType *string `json:"certType,omitempty" xml:"certType,omitempty"`
	// 法人证件号
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
	// 子机构行业资质图片
	QualificationImage *string `json:"qualificationImage,omitempty" xml:"qualificationImage,omitempty"`
	// 子机构行业资质类型
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
	// 账户账号
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// 账户名称 账号类型银行卡时为卡户名
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	// 卡类型
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	// 支行名称
	BankBranchName *string `json:"bankBranchName,omitempty" xml:"bankBranchName,omitempty"`
	// 开户行所在地 市
	BankCity *string `json:"bankCity,omitempty" xml:"bankCity,omitempty"`
	// 联行号
	BankCode *string `json:"bankCode,omitempty" xml:"bankCode,omitempty"`
	// 银行名称
	BankName *string `json:"bankName,omitempty" xml:"bankName,omitempty"`
	// 开户行所在地 省
	BankProvince *string `json:"bankProvince,omitempty" xml:"bankProvince,omitempty"`
	// 开户行简称缩写
	BankShortNameCode *string `json:"bankShortNameCode,omitempty" xml:"bankShortNameCode,omitempty"`
	// 账号类型
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 账户使用类型
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
	// 详细地址
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// 市码
	CityCode *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	// 区码
	DistrictCode *string `json:"districtCode,omitempty" xml:"districtCode,omitempty"`
	// 省码
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
	// 授权函图片url
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
	// 别名
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// 机构识别码
	Mcc *string `json:"mcc,omitempty" xml:"mcc,omitempty"`
	// 名称
	SubInstName *string `json:"subInstName,omitempty" xml:"subInstName,omitempty"`
	// 类型
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
	// 证件图片, 如果是特殊行业必填
	CertImage *string `json:"certImage,omitempty" xml:"certImage,omitempty"`
	// 证件号码
	CertNo *string `json:"certNo,omitempty" xml:"certNo,omitempty"`
	// 证件类型
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
	// 是否接受电票
	AcceptElectronic *bool `json:"acceptElectronic,omitempty" xml:"acceptElectronic,omitempty"`
	// 开票地址
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// 是否自动开票
	AutoInvoice *bool `json:"autoInvoice,omitempty" xml:"autoInvoice,omitempty"`
	// 银行账户
	BankAccount *string `json:"bankAccount,omitempty" xml:"bankAccount,omitempty"`
	// 银行名称
	BankName *string `json:"bankName,omitempty" xml:"bankName,omitempty"`
	// 收件地址
	MailAddress *ModifySubInstitutionRequestSubInstInvoiceInfoMailAddress `json:"mailAddress,omitempty" xml:"mailAddress,omitempty" type:"Struct"`
	// 收件人名称
	MailName *string `json:"mailName,omitempty" xml:"mailName,omitempty"`
	// 收件人号码
	MailPhone *string `json:"mailPhone,omitempty" xml:"mailPhone,omitempty"`
	// 纳税人识别号
	TaxNo *string `json:"taxNo,omitempty" xml:"taxNo,omitempty"`
	// 纳税人资质
	TaxPayerQualification *string `json:"taxPayerQualification,omitempty" xml:"taxPayerQualification,omitempty"`
	// 纳税人资格开始时间
	TaxPayerValidDate *string `json:"taxPayerValidDate,omitempty" xml:"taxPayerValidDate,omitempty"`
	// 开票电话
	Telephone *string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	// 纳税人抬头
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
	// 详细地址
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// 市码
	CityCode *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	// 区码
	DistrictCode *string `json:"districtCode,omitempty" xml:"districtCode,omitempty"`
	// 省码
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
	// 内景照
	InDoorImages []*string `json:"inDoorImages,omitempty" xml:"inDoorImages,omitempty" type:"Repeated"`
	// 外景照
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
	// 修改申请单号
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySubInstitutionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 订单金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 收费金额
	ChargeAmount *string `json:"chargeAmount,omitempty" xml:"chargeAmount,omitempty"`
	// 企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 扩展信息
	ExtInfo *string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// 交易开始时间
	GmtTradeCreate *string `json:"gmtTradeCreate,omitempty" xml:"gmtTradeCreate,omitempty"`
	// 交易结束时间
	GmtTradeFinish *string `json:"gmtTradeFinish,omitempty" xml:"gmtTradeFinish,omitempty"`
	// merchantName
	MerchantName *string `json:"merchantName,omitempty" xml:"merchantName,omitempty"`
	// 支付渠道明细信息
	PayChannelDetailList []*NotifyPayCodePayResultRequestPayChannelDetailList `json:"payChannelDetailList,omitempty" xml:"payChannelDetailList,omitempty" type:"Repeated"`
	// 付款码值
	PayCode *string `json:"payCode,omitempty" xml:"payCode,omitempty"`
	// 订单优惠金额
	PromotionAmount *string `json:"promotionAmount,omitempty" xml:"promotionAmount,omitempty"`
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 订单标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 支付失败错误码
	TradeErrorCode *string `json:"tradeErrorCode,omitempty" xml:"tradeErrorCode,omitempty"`
	// 支付失败信息
	TradeErrorMsg *string `json:"tradeErrorMsg,omitempty" xml:"tradeErrorMsg,omitempty"`
	// 交易号
	TradeNo *string `json:"tradeNo,omitempty" xml:"tradeNo,omitempty"`
	// 交易状态
	TradeStatus *string `json:"tradeStatus,omitempty" xml:"tradeStatus,omitempty"`
	// 用户id
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
	// 支付金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 资金工具明细
	FundToolDetailList []*NotifyPayCodePayResultRequestPayChannelDetailListFundToolDetailList `json:"fundToolDetailList,omitempty" xml:"fundToolDetailList,omitempty" type:"Repeated"`
	// 开始时间
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 结束时间
	GmtFinish *string `json:"gmtFinish,omitempty" xml:"gmtFinish,omitempty"`
	// 支付渠道名称
	PayChannelName *string `json:"payChannelName,omitempty" xml:"payChannelName,omitempty"`
	// 支付渠道单号
	PayChannelOrderNo *string `json:"payChannelOrderNo,omitempty" xml:"payChannelOrderNo,omitempty"`
	// 支付渠道类型
	PayChannelType *string `json:"payChannelType,omitempty" xml:"payChannelType,omitempty"`
	// 优惠金额
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
	// 1.00
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 扩展信息
	ExtInfo *string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// 资金渠道名称
	FundToolName *string `json:"fundToolName,omitempty" xml:"fundToolName,omitempty"`
	// 开始时间
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 结束时间
	GmtFinish *string `json:"gmtFinish,omitempty" xml:"gmtFinish,omitempty"`
	// 是否是优惠工具
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
	// 处理结果
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *NotifyPayCodePayResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 退款时间
	GmtRefund *string `json:"gmtRefund,omitempty" xml:"gmtRefund,omitempty"`
	// 支付渠道信息
	PayChannelDetailList []*NotifyPayCodeRefundResultRequestPayChannelDetailList `json:"payChannelDetailList,omitempty" xml:"payChannelDetailList,omitempty" type:"Repeated"`
	// 支付时使用的付款码
	PayCode *string `json:"payCode,omitempty" xml:"payCode,omitempty"`
	// 退款金额
	RefundAmount *string `json:"refundAmount,omitempty" xml:"refundAmount,omitempty"`
	// 本次退款订单号
	RefundOrderNo *string `json:"refundOrderNo,omitempty" xml:"refundOrderNo,omitempty"`
	// 退款的优惠金额
	RefundPromotionAmount *string `json:"refundPromotionAmount,omitempty" xml:"refundPromotionAmount,omitempty"`
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 交易订单号
	TradeNo *string `json:"tradeNo,omitempty" xml:"tradeNo,omitempty"`
	// 用户id
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
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 支付资金列表
	FundToolDetailList []*NotifyPayCodeRefundResultRequestPayChannelDetailListFundToolDetailList `json:"fundToolDetailList,omitempty" xml:"fundToolDetailList,omitempty" type:"Repeated"`
	// 支付渠道名称
	PayChannelName *string `json:"payChannelName,omitempty" xml:"payChannelName,omitempty"`
	// 支付渠道号
	PayChannelOrderNo *string `json:"payChannelOrderNo,omitempty" xml:"payChannelOrderNo,omitempty"`
	// 支付渠道退款号
	PayChannelRefundOrderNo *string `json:"payChannelRefundOrderNo,omitempty" xml:"payChannelRefundOrderNo,omitempty"`
	// 支付渠道类型
	PayChannelType *string `json:"payChannelType,omitempty" xml:"payChannelType,omitempty"`
	// 优惠金额
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
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 扩展信息
	ExtInfo *string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// 资金工具名称
	FundToolName *string `json:"fundToolName,omitempty" xml:"fundToolName,omitempty"`
	// 创建时间
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 完成时间
	GmtFinish *string `json:"gmtFinish,omitempty" xml:"gmtFinish,omitempty"`
	// 是否是优惠工具
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
	// 处理结果
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *NotifyPayCodeRefundResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 企业ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 码值
	PayCode *string `json:"payCode,omitempty" xml:"payCode,omitempty"`
	// 备注信息
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 用户和企业的关系类型，区分内部员工，外部联系人，无关系普通用户
	UserCorpRelationType *string `json:"userCorpRelationType,omitempty" xml:"userCorpRelationType,omitempty"`
	// 用户身份标识
	UserIdentity *string `json:"userIdentity,omitempty" xml:"userIdentity,omitempty"`
	// 验证事件，长度不超过8个中文
	VerifyEvent *string `json:"verifyEvent,omitempty" xml:"verifyEvent,omitempty"`
	// 验证地点，调用时请务必传入，以便生成工牌使用记录
	VerifyLocation *string `json:"verifyLocation,omitempty" xml:"verifyLocation,omitempty"`
	// 验证流水号，长度不超过32位，用户下唯一，调用时请务必传入，以便生成工牌使用记录
	VerifyNo *string `json:"verifyNo,omitempty" xml:"verifyNo,omitempty"`
	// 验证结果
	VerifyResult *bool `json:"verifyResult,omitempty" xml:"verifyResult,omitempty"`
	// 验证时间
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
	// 结果
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *NotifyVerifyResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *NotifyVerifyResultResponse) SetBody(v *NotifyVerifyResultResponseBody) *NotifyVerifyResultResponse {
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
	// 外部商户批次号
	OutBatchNo *string `json:"outBatchNo,omitempty" xml:"outBatchNo,omitempty"`
	// 当前页数
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页记录数
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
	// 明细列表
	BatchTradeDetailList []*QueryBatchTradeDetailListResponseBodyBatchTradeDetailList `json:"batchTradeDetailList,omitempty" xml:"batchTradeDetailList,omitempty" type:"Repeated"`
	// 当前页数
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 单页条数
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 总记录数
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
	// 总页数
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
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 明细单号
	DetailNo *string `json:"detailNo,omitempty" xml:"detailNo,omitempty"`
	// 订单时间时间
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 支付完成时间
	GmtFinish *string `json:"gmtFinish,omitempty" xml:"gmtFinish,omitempty"`
	// 备注
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 收款方电子钱包持有者姓名
	PayeeAccountName *string `json:"payeeAccountName,omitempty" xml:"payeeAccountName,omitempty"`
	// 收款人账号
	PayeeAccountNo *string `json:"payeeAccountNo,omitempty" xml:"payeeAccountNo,omitempty"`
	// 收款账号类型
	PayeeAccountType *string `json:"payeeAccountType,omitempty" xml:"payeeAccountType,omitempty"`
	// 序号
	SerialNo *int64 `json:"serialNo,omitempty" xml:"serialNo,omitempty"`
	// 状态
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryBatchTradeDetailListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 外部商户批次号列表
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
	// 批量交易订单VO
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
	// 支付宝批次订单号
	AlipayTransId *string `json:"alipayTransId,omitempty" xml:"alipayTransId,omitempty"`
	// 明细处理失败的支付汇总金额
	FailAmount *string `json:"failAmount,omitempty" xml:"failAmount,omitempty"`
	// 失败笔数
	FailCount *int64 `json:"failCount,omitempty" xml:"failCount,omitempty"`
	// 失败原因
	FailReason *string `json:"failReason,omitempty" xml:"failReason,omitempty"`
	// 批次完成交易时间
	GmtFinish *string `json:"gmtFinish,omitempty" xml:"gmtFinish,omitempty"`
	// 批次受理交易时间
	GmtSubmit *string `json:"gmtSubmit,omitempty" xml:"gmtSubmit,omitempty"`
	// 批次号
	OutBatchNo *string `json:"outBatchNo,omitempty" xml:"outBatchNo,omitempty"`
	// 付款人staffId
	PayerStaffId *string `json:"payerStaffId,omitempty" xml:"payerStaffId,omitempty"`
	// 付款方需要支付的金额（元）
	PaymentAmount *string `json:"paymentAmount,omitempty" xml:"paymentAmount,omitempty"`
	// 支付币种
	PaymentCurrency *string `json:"paymentCurrency,omitempty" xml:"paymentCurrency,omitempty"`
	// 状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 成功金额（元）
	SuccessAmount *string `json:"successAmount,omitempty" xml:"successAmount,omitempty"`
	// 成功笔数
	SuccessCount *int64 `json:"successCount,omitempty" xml:"successCount,omitempty"`
	// 批次的总金额（元）
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryBatchTradeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 账号列表
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
	// 账户分类
	AccountClass *string `json:"accountClass,omitempty" xml:"accountClass,omitempty"`
	// 账号唯一id
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// 账号名称
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	// 付款账号（脱敏）
	AccountNo *string `json:"accountNo,omitempty" xml:"accountNo,omitempty"`
	// 账户备注
	AccountRemark *string `json:"accountRemark,omitempty" xml:"accountRemark,omitempty"`
	// 账户类型
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryPayAccountListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 主机构编号
	InstId *string `json:"instId,omitempty" xml:"instId,omitempty"`
	// 申请单号，和外部流水号至少一个必填
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// 外部流水号，和申请单编号至少一个必填
	OutTradeNo *string `json:"outTradeNo,omitempty" xml:"outTradeNo,omitempty"`
	// 子机构编号
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
	// 失败原因
	FailReason *string `json:"failReason,omitempty" xml:"failReason,omitempty"`
	// 审核时间
	GmtAudit *string `json:"gmtAudit,omitempty" xml:"gmtAudit,omitempty"`
	// 主机构编号
	InstId *string `json:"instId,omitempty" xml:"instId,omitempty"`
	// 申请单号
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// 外部流水号
	OutTradeNo *string `json:"outTradeNo,omitempty" xml:"outTradeNo,omitempty"`
	// 申请单状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 子机构编号
	SubInstId *string `json:"subInstId,omitempty" xml:"subInstId,omitempty"`
	// 子机构名称
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRegisterOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 业务编码
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	// 业务场景
	BizScene *string `json:"bizScene,omitempty" xml:"bizScene,omitempty"`
	// 主机构编号
	InstId *string `json:"instId,omitempty" xml:"instId,omitempty"`
	// 子机构编号
	SubInstId *string `json:"subInstId,omitempty" xml:"subInstId,omitempty"`
	// 付款人staffId
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
	// 协议编号
	AgreementNo *string `json:"agreementNo,omitempty" xml:"agreementNo,omitempty"`
	// 组织id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 实际过期日期
	GmtExpire *string `json:"gmtExpire,omitempty" xml:"gmtExpire,omitempty"`
	// 实际签约日期
	GmtSign *string `json:"gmtSign,omitempty" xml:"gmtSign,omitempty"`
	// 实际生效日期
	GmtValid *string `json:"gmtValid,omitempty" xml:"gmtValid,omitempty"`
	// 主机构id
	InstId *string `json:"instId,omitempty" xml:"instId,omitempty"`
	// 支付渠道
	PayChannel *string `json:"payChannel,omitempty" xml:"payChannel,omitempty"`
	// 实际支付账户名（脱敏）
	PayChannelAccountName *string `json:"payChannelAccountName,omitempty" xml:"payChannelAccountName,omitempty"`
	// 实际支付账号（脱敏）
	PayChannelAccountNo *string `json:"payChannelAccountNo,omitempty" xml:"payChannelAccountNo,omitempty"`
	// 用户id
	StaffId *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
	// 签约状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 子机构id
	SubInstId *string `json:"subInstId,omitempty" xml:"subInstId,omitempty"`
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

func (s *QueryUserAgreementResponseBody) SetStaffId(v string) *QueryUserAgreementResponseBody {
	s.StaffId = &v
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

type QueryUserAgreementResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUserAgreementResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 支付宝uid
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUserAlipayAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 外部订单流水号
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
	// 代扣金额（元）
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 订单创建日期
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 付款完成日期
	GmtPay *string `json:"gmtPay,omitempty" xml:"gmtPay,omitempty"`
	// 主机构编号
	InstId *string `json:"instId,omitempty" xml:"instId,omitempty"`
	// 钉钉订单号
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	// 外部订单号
	OutTradeNo *string `json:"outTradeNo,omitempty" xml:"outTradeNo,omitempty"`
	// 支付渠道
	PayChannel *string `json:"payChannel,omitempty" xml:"payChannel,omitempty"`
	// 支付渠道支付账号（脱敏后返回）
	PayChannelAccountNo *string `json:"payChannelAccountNo,omitempty" xml:"payChannelAccountNo,omitempty"`
	// 付款人staffId
	PayerStaffId *string `json:"payerStaffId,omitempty" xml:"payerStaffId,omitempty"`
	// 代扣备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 子机构编号
	SubInstId *string `json:"subInstId,omitempty" xml:"subInstId,omitempty"`
	// 代扣标题
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

func (s *QueryWithholdingOrderResponseBody) SetPayerStaffId(v string) *QueryWithholdingOrderResponseBody {
	s.PayerStaffId = &v
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryWithholdingOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 码标识，由钉钉颁发
	CodeIdentity *string `json:"codeIdentity,omitempty" xml:"codeIdentity,omitempty"`
	// 开通的企业ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 扩展参数
	ExtInfo map[string]*string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// 状态，OPEN或CLOSED
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
	// 码标识
	CodeIdentity *string `json:"codeIdentity,omitempty" xml:"codeIdentity,omitempty"`
	// 开通的企业ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 扩展参数
	ExtInfo map[string]*string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// 状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveCorpPayCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 协议编号
	AgreementNo *string `json:"agreementNo,omitempty" xml:"agreementNo,omitempty"`
	// 业务代码
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	// 业务场景
	BizScene *string `json:"bizScene,omitempty" xml:"bizScene,omitempty"`
	// 主机构编号
	InstId *string `json:"instId,omitempty" xml:"instId,omitempty"`
	// 子机构编号
	SubInstId *string `json:"subInstId,omitempty" xml:"subInstId,omitempty"`
	// 付款人staffId
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
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
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
	// 有效时间列表，对于连续时间段，只需传入一个对象即可，注意过期时间必须晚于最晚结束时间
	AvailableTimes []*UpateUserCodeInstanceRequestAvailableTimes `json:"availableTimes,omitempty" xml:"availableTimes,omitempty" type:"Repeated"`
	// 用户码ID
	CodeId *string `json:"codeId,omitempty" xml:"codeId,omitempty"`
	// 码标识
	CodeIdentity *string `json:"codeIdentity,omitempty" xml:"codeIdentity,omitempty"`
	// 码值
	CodeValue *string `json:"codeValue,omitempty" xml:"codeValue,omitempty"`
	// 企业ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 扩展参数
	ExtInfo map[string]interface{} `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// 临时码，传入过期时间
	GmtExpired *string `json:"gmtExpired,omitempty" xml:"gmtExpired,omitempty"`
	// 状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 用户和企业的关系类型，区分内部员工，外部联系人，无关系普通用户
	UserCorpRelationType *string `json:"userCorpRelationType,omitempty" xml:"userCorpRelationType,omitempty"`
	// 用户身份标识，取值和用户企业关系类型相关，如果企业无关，传入手机号
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
	// 结束时间
	GmtEnd *string `json:"gmtEnd,omitempty" xml:"gmtEnd,omitempty"`
	// 开始时间
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
	// 码ID
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpateUserCodeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpateUserCodeInstanceResponse) SetBody(v *UpateUserCodeInstanceResponseBody) *UpateUserCodeInstanceResponse {
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
	// 图片内容
	ImageContent *string `json:"imageContent,omitempty" xml:"imageContent,omitempty"`
	// 图片名称
	ImageName *string `json:"imageName,omitempty" xml:"imageName,omitempty"`
	// 图片类型
	ImageType *string `json:"imageType,omitempty" xml:"imageType,omitempty"`
	// 主机构id
	InstId *string `json:"instId,omitempty" xml:"instId,omitempty"`
	// 进件渠道
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
	// 进件图片上传响应
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadRegisterImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 业务编码
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	// 业务场景
	BizScene *string `json:"bizScene,omitempty" xml:"bizScene,omitempty"`
	// 主机构编号
	InstId *string `json:"instId,omitempty" xml:"instId,omitempty"`
	// 支付渠道
	PayChannel *string `json:"payChannel,omitempty" xml:"payChannel,omitempty"`
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 签约场景
	SignScene *string `json:"signScene,omitempty" xml:"signScene,omitempty"`
	// 子机构编号
	SubInstId *string `json:"subInstId,omitempty" xml:"subInstId,omitempty"`
	// 子商户商户名称
	SubMerchantName *string `json:"subMerchantName,omitempty" xml:"subMerchantName,omitempty"`
	// 子商户服务描述
	SubMerchantServiceDesc *string `json:"subMerchantServiceDesc,omitempty" xml:"subMerchantServiceDesc,omitempty"`
	// 子商户服务名称
	SubMerchantServiceName *string `json:"subMerchantServiceName,omitempty" xml:"subMerchantServiceName,omitempty"`
	// 付款人staffId
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
	// 拉起签约页的字符串
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UserAgreementPageSignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

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
	_result = &ApplyBatchPayResponse{}
	_body, _err := client.DoROARequest(tea.String("ApplyBatchPay"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/finance/batchTrades/orders/pay"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ConsultCreateSubInstitutionWithOptions(request *ConsultCreateSubInstitutionRequest, headers *ConsultCreateSubInstitutionHeaders, runtime *util.RuntimeOptions) (_result *ConsultCreateSubInstitutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BindingAlipayLogonId)) {
		body["bindingAlipayLogonId"] = request.BindingAlipayLogonId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.ContractInfo))) {
		body["contractInfo"] = request.ContractInfo
	}

	if !tea.BoolValue(util.IsUnset(request.InstId)) {
		body["instId"] = request.InstId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.LegalPersonCertInfo))) {
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

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SettleInfo))) {
		body["settleInfo"] = request.SettleInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Solution)) {
		body["solution"] = request.Solution
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SubInstAddressInfo))) {
		body["subInstAddressInfo"] = request.SubInstAddressInfo
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SubInstAuthInfo))) {
		body["subInstAuthInfo"] = request.SubInstAuthInfo
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SubInstBasicInfo))) {
		body["subInstBasicInfo"] = request.SubInstBasicInfo
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SubInstCertifyInfo))) {
		body["subInstCertifyInfo"] = request.SubInstCertifyInfo
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstId)) {
		body["subInstId"] = request.SubInstId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SubInstInvoiceInfo))) {
		body["subInstInvoiceInfo"] = request.SubInstInvoiceInfo
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SubInstShopInfo))) {
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
	_result = &ConsultCreateSubInstitutionResponse{}
	_body, _err := client.DoROARequest(tea.String("ConsultCreateSubInstitution"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/finance/institutions/subInstitutions/consult"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CreatWithholdingOrderAndPayResponse{}
	_body, _err := client.DoROARequest(tea.String("CreatWithholdingOrderAndPay"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/finance/withholdingOrders"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CreateBatchTradeOrderResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateBatchTradeOrder"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/finance/batchTrades/orders"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CreateSubInstitutionWithOptions(request *CreateSubInstitutionRequest, headers *CreateSubInstitutionHeaders, runtime *util.RuntimeOptions) (_result *CreateSubInstitutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BindingAlipayLogonId)) {
		body["bindingAlipayLogonId"] = request.BindingAlipayLogonId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.ContractInfo))) {
		body["contractInfo"] = request.ContractInfo
	}

	if !tea.BoolValue(util.IsUnset(request.InstId)) {
		body["instId"] = request.InstId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.LegalPersonCertInfo))) {
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

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SettleInfo))) {
		body["settleInfo"] = request.SettleInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Solution)) {
		body["solution"] = request.Solution
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SubInstAddressInfo))) {
		body["subInstAddressInfo"] = request.SubInstAddressInfo
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SubInstAuthInfo))) {
		body["subInstAuthInfo"] = request.SubInstAuthInfo
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SubInstBasicInfo))) {
		body["subInstBasicInfo"] = request.SubInstBasicInfo
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SubInstCertifyInfo))) {
		body["subInstCertifyInfo"] = request.SubInstCertifyInfo
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstId)) {
		body["subInstId"] = request.SubInstId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SubInstInvoiceInfo))) {
		body["subInstInvoiceInfo"] = request.SubInstInvoiceInfo
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SubInstShopInfo))) {
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
	_result = &CreateSubInstitutionResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateSubInstitution"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/finance/institutions/subInstitutions"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CreateUserCodeInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateUserCodeInstance"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/finance/payCodes/userInstances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &DecodePayCodeResponse{}
	_body, _err := client.DoROARequest(tea.String("DecodePayCode"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/finance/payCodes/decode"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ModifySubInstitutionWithOptions(request *ModifySubInstitutionRequest, headers *ModifySubInstitutionHeaders, runtime *util.RuntimeOptions) (_result *ModifySubInstitutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BindingAlipayLogonId)) {
		body["bindingAlipayLogonId"] = request.BindingAlipayLogonId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.ContractInfo))) {
		body["contractInfo"] = request.ContractInfo
	}

	if !tea.BoolValue(util.IsUnset(request.InstId)) {
		body["instId"] = request.InstId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.LegalPersonCertInfo))) {
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

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SettleInfo))) {
		body["settleInfo"] = request.SettleInfo
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SubInstAddressInfo))) {
		body["subInstAddressInfo"] = request.SubInstAddressInfo
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SubInstAuthInfo))) {
		body["subInstAuthInfo"] = request.SubInstAuthInfo
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SubInstBasicInfo))) {
		body["subInstBasicInfo"] = request.SubInstBasicInfo
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SubInstCertifyInfo))) {
		body["subInstCertifyInfo"] = request.SubInstCertifyInfo
	}

	if !tea.BoolValue(util.IsUnset(request.SubInstId)) {
		body["subInstId"] = request.SubInstId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SubInstInvoiceInfo))) {
		body["subInstInvoiceInfo"] = request.SubInstInvoiceInfo
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SubInstShopInfo))) {
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
	_result = &ModifySubInstitutionResponse{}
	_body, _err := client.DoROARequest(tea.String("ModifySubInstitution"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/finance/institutions/subInstitutions/modify"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &NotifyPayCodePayResultResponse{}
	_body, _err := client.DoROARequest(tea.String("NotifyPayCodePayResult"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/finance/payCodes/payResults/notify"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &NotifyPayCodeRefundResultResponse{}
	_body, _err := client.DoROARequest(tea.String("NotifyPayCodeRefundResult"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/finance/payCodes/refundResults/notify"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &NotifyVerifyResultResponse{}
	_body, _err := client.DoROARequest(tea.String("NotifyVerifyResult"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/finance/payCodes/verifyResults/notify"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryBatchTradeDetailListResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryBatchTradeDetailList"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/finance/batchTrades/details"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryBatchTradeOrderResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryBatchTradeOrder"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/finance/batchTrades/orders/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryPayAccountListResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryPayAccountList"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/finance/payAccounts"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryRegisterOrderResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryRegisterOrder"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/finance/institutions/subInstitutions/orders"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryUserAgreementResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryUserAgreement"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/finance/userAgreements"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryUserAlipayAccountResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryUserAlipayAccount"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/finance/userAlipayAccounts"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryWithholdingOrderResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryWithholdingOrder"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/finance/withholdingOrders"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &SaveCorpPayCodeResponse{}
	_body, _err := client.DoROARequest(tea.String("SaveCorpPayCode"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/finance/payCodes/corpSettings"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UnsignUserAgreementResponse{}
	_body, _err := client.DoROARequest(tea.String("UnsignUserAgreement"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/finance/userAgreements/unsign"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UpateUserCodeInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("UpateUserCodeInstance"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/finance/payCodes/userInstances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UploadRegisterImageResponse{}
	_body, _err := client.DoROARequest(tea.String("UploadRegisterImage"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/finance/institutions/images"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UserAgreementPageSignResponse{}
	_body, _err := client.DoROARequest(tea.String("UserAgreementPageSign"), tea.String("finance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/finance/userAgreements"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
