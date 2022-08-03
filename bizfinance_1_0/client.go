// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package bizfinance_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BatchAddInvoiceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchAddInvoiceHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchAddInvoiceHeaders) GoString() string {
	return s.String()
}

func (s *BatchAddInvoiceHeaders) SetCommonHeaders(v map[string]*string) *BatchAddInvoiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchAddInvoiceHeaders) SetXAcsDingtalkAccessToken(v string) *BatchAddInvoiceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchAddInvoiceRequest struct {
	// 发票模型
	GeneralInvoiceVOList []*BatchAddInvoiceRequestGeneralInvoiceVOList `json:"generalInvoiceVOList,omitempty" xml:"generalInvoiceVOList,omitempty" type:"Repeated"`
	// 操作员
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
}

func (s BatchAddInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchAddInvoiceRequest) GoString() string {
	return s.String()
}

func (s *BatchAddInvoiceRequest) SetGeneralInvoiceVOList(v []*BatchAddInvoiceRequestGeneralInvoiceVOList) *BatchAddInvoiceRequest {
	s.GeneralInvoiceVOList = v
	return s
}

func (s *BatchAddInvoiceRequest) SetOperator(v string) *BatchAddInvoiceRequest {
	s.Operator = &v
	return s
}

type BatchAddInvoiceRequestGeneralInvoiceVOList struct {
	// 账期时间
	AccountPeriod *string `json:"accountPeriod,omitempty" xml:"accountPeriod,omitempty"`
	// 不含税金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 含税金额
	AmountWithTax *string `json:"amountWithTax,omitempty" xml:"amountWithTax,omitempty"`
	// 校验码
	CheckCode *string `json:"checkCode,omitempty" xml:"checkCode,omitempty"`
	// 查验时间
	CheckTime *string `json:"checkTime,omitempty" xml:"checkTime,omitempty"`
	// 开票日期
	DrewDate *string `json:"drewDate,omitempty" xml:"drewDate,omitempty"`
	// 电票版式文件下载地址
	ElectronicUrl *string `json:"electronicUrl,omitempty" xml:"electronicUrl,omitempty"`
	// 财务类型，INPUT-VAT(进项),OUTPUT_VAT(销项)
	FinanceType *string `json:"financeType,omitempty" xml:"financeType,omitempty"`
	// 资金类型 ，RED（红票），（BLUE）蓝票
	FundType *string `json:"fundType,omitempty" xml:"fundType,omitempty"`
	// 常规发票明细
	GeneralInvoiceDetailVOList []*BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList `json:"generalInvoiceDetailVOList,omitempty" xml:"generalInvoiceDetailVOList,omitempty" type:"Repeated"`
	// 发票代码
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	// 发票号码
	InvoiceNo *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
	// 发票状态
	InvoiceStatus *string `json:"invoiceStatus,omitempty" xml:"invoiceStatus,omitempty"`
	// 发票类型
	InvoiceType *string `json:"invoiceType,omitempty" xml:"invoiceType,omitempty"`
	// 机器码
	MachineCode *string `json:"machineCode,omitempty" xml:"machineCode,omitempty"`
	// 成品油标识
	OilFlag *string `json:"oilFlag,omitempty" xml:"oilFlag,omitempty"`
	// 收款人
	Payee *string `json:"payee,omitempty" xml:"payee,omitempty"`
	// 审批单实例
	ProcessInstCode *string `json:"processInstCode,omitempty" xml:"processInstCode,omitempty"`
	// 审批单类型
	ProcessInstType *string `json:"processInstType,omitempty" xml:"processInstType,omitempty"`
	// 购方地址
	PurchaserAddress *string `json:"purchaserAddress,omitempty" xml:"purchaserAddress,omitempty"`
	// 购方银行账户
	PurchaserBankAccount *string `json:"purchaserBankAccount,omitempty" xml:"purchaserBankAccount,omitempty"`
	// 购方银行名称
	PurchaserBankNameAccount *string `json:"purchaserBankNameAccount,omitempty" xml:"purchaserBankNameAccount,omitempty"`
	// 购方名称
	PurchaserName *string `json:"purchaserName,omitempty" xml:"purchaserName,omitempty"`
	// 购方税号
	PurchaserTaxNo *string `json:"purchaserTaxNo,omitempty" xml:"purchaserTaxNo,omitempty"`
	// 购方电话
	PurchaserTel *string `json:"purchaserTel,omitempty" xml:"purchaserTel,omitempty"`
	// 备注
	Remark                         *string                                                                     `json:"remark,omitempty" xml:"remark,omitempty"`
	SecondHandCarInvoiceDetailList []*BatchAddInvoiceRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList `json:"secondHandCarInvoiceDetailList,omitempty" xml:"secondHandCarInvoiceDetailList,omitempty" type:"Repeated"`
	// 销方地址
	SellerAddress *string `json:"sellerAddress,omitempty" xml:"sellerAddress,omitempty"`
	// 销方银行账户
	SellerBankAccount *string `json:"sellerBankAccount,omitempty" xml:"sellerBankAccount,omitempty"`
	// 销方银行名称
	SellerBankNameAccount *string `json:"sellerBankNameAccount,omitempty" xml:"sellerBankNameAccount,omitempty"`
	// 销方名称
	SellerName *string `json:"sellerName,omitempty" xml:"sellerName,omitempty"`
	// 销方税号
	SellerTaxNo *string `json:"sellerTaxNo,omitempty" xml:"sellerTaxNo,omitempty"`
	// 销方电话
	SellerTel *string `json:"sellerTel,omitempty" xml:"sellerTel,omitempty"`
	// 代开发票标识 1-自开，2-代开
	SupplySign *string `json:"supplySign,omitempty" xml:"supplySign,omitempty"`
	// 税额
	TaxAmount                   *string                                                                  `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	UsedVehicleSaleDetailVOList []*BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList `json:"usedVehicleSaleDetailVOList,omitempty" xml:"usedVehicleSaleDetailVOList,omitempty" type:"Repeated"`
	VehicleSaleDetailVOList     []*BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList     `json:"vehicleSaleDetailVOList,omitempty" xml:"vehicleSaleDetailVOList,omitempty" type:"Repeated"`
	// 发票查验状态
	VerifyStatus *string `json:"verifyStatus,omitempty" xml:"verifyStatus,omitempty"`
	// 凭证code
	VoucherCode *string `json:"voucherCode,omitempty" xml:"voucherCode,omitempty"`
	// 生成凭证状态
	VoucherStatus *string `json:"voucherStatus,omitempty" xml:"voucherStatus,omitempty"`
}

func (s BatchAddInvoiceRequestGeneralInvoiceVOList) String() string {
	return tea.Prettify(s)
}

func (s BatchAddInvoiceRequestGeneralInvoiceVOList) GoString() string {
	return s.String()
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetAccountPeriod(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.AccountPeriod = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetAmount(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.Amount = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetAmountWithTax(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.AmountWithTax = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetCheckCode(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.CheckCode = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetCheckTime(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.CheckTime = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetDrewDate(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.DrewDate = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetElectronicUrl(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.ElectronicUrl = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetFinanceType(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.FinanceType = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetFundType(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.FundType = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetGeneralInvoiceDetailVOList(v []*BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.GeneralInvoiceDetailVOList = v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetInvoiceCode(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.InvoiceCode = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetInvoiceNo(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.InvoiceNo = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetInvoiceStatus(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.InvoiceStatus = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetInvoiceType(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.InvoiceType = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetMachineCode(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.MachineCode = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetOilFlag(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.OilFlag = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetPayee(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.Payee = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetProcessInstCode(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.ProcessInstCode = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetProcessInstType(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.ProcessInstType = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetPurchaserAddress(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.PurchaserAddress = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetPurchaserBankAccount(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.PurchaserBankAccount = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetPurchaserBankNameAccount(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.PurchaserBankNameAccount = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetPurchaserName(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.PurchaserName = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetPurchaserTaxNo(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.PurchaserTaxNo = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetPurchaserTel(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.PurchaserTel = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetRemark(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.Remark = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetSecondHandCarInvoiceDetailList(v []*BatchAddInvoiceRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.SecondHandCarInvoiceDetailList = v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetSellerAddress(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.SellerAddress = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetSellerBankAccount(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.SellerBankAccount = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetSellerBankNameAccount(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.SellerBankNameAccount = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetSellerName(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.SellerName = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetSellerTaxNo(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.SellerTaxNo = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetSellerTel(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.SellerTel = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetSupplySign(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.SupplySign = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetTaxAmount(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.TaxAmount = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetUsedVehicleSaleDetailVOList(v []*BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.UsedVehicleSaleDetailVOList = v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetVehicleSaleDetailVOList(v []*BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.VehicleSaleDetailVOList = v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetVerifyStatus(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.VerifyStatus = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetVoucherCode(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.VoucherCode = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOList) SetVoucherStatus(v string) *BatchAddInvoiceRequestGeneralInvoiceVOList {
	s.VoucherStatus = &v
	return s
}

type BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList struct {
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 商品名称
	GoodsName *string `json:"goodsName,omitempty" xml:"goodsName,omitempty"`
	// 数量
	Quantity *string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 税收分类编码
	RevenueCode *string `json:"revenueCode,omitempty" xml:"revenueCode,omitempty"`
	// 行号
	RowNo *string `json:"rowNo,omitempty" xml:"rowNo,omitempty"`
	// 规格型号
	Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
	// 税额
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	// 是否享受税收优惠：0-不享受，1-享受
	TaxPre *string `json:"taxPre,omitempty" xml:"taxPre,omitempty"`
	// 优惠政策类型
	TaxPreType *string `json:"taxPreType,omitempty" xml:"taxPreType,omitempty"`
	// 税率
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// 单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 单价
	UnitPrice *string `json:"unitPrice,omitempty" xml:"unitPrice,omitempty"`
}

func (s BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) String() string {
	return tea.Prettify(s)
}

func (s BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) GoString() string {
	return s.String()
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetAmount(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.Amount = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetGoodsName(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.GoodsName = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetQuantity(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.Quantity = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetRevenueCode(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.RevenueCode = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetRowNo(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.RowNo = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetSpecification(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.Specification = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetTaxAmount(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.TaxAmount = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetTaxPre(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.TaxPre = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetTaxPreType(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.TaxPreType = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetTaxRate(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.TaxRate = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetUnit(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.Unit = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetUnitPrice(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.UnitPrice = &v
	return s
}

type BatchAddInvoiceRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList struct {
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 车牌号
	CardNo *string `json:"cardNo,omitempty" xml:"cardNo,omitempty"`
	// 通行日期止
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// 商品名称
	GoodsName *string `json:"goodsName,omitempty" xml:"goodsName,omitempty"`
	// 税收分类编码
	RevenueCode *string `json:"revenueCode,omitempty" xml:"revenueCode,omitempty"`
	// 行号
	RowNo *string `json:"rowNo,omitempty" xml:"rowNo,omitempty"`
	// 通行日期起
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// 税额
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	// 税率
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// 类型
	VehicleType *string `json:"vehicleType,omitempty" xml:"vehicleType,omitempty"`
}

func (s BatchAddInvoiceRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) String() string {
	return tea.Prettify(s)
}

func (s BatchAddInvoiceRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) GoString() string {
	return s.String()
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetAmount(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.Amount = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetCardNo(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.CardNo = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetEndDate(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.EndDate = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetGoodsName(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.GoodsName = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetRevenueCode(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.RevenueCode = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetRowNo(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.RowNo = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetStartDate(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.StartDate = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetTaxAmount(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.TaxAmount = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetTaxRate(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.TaxRate = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetVehicleType(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.VehicleType = &v
	return s
}

type BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList struct {
	// 经营、拍卖单位
	AuctionUnit *string `json:"auctionUnit,omitempty" xml:"auctionUnit,omitempty"`
	// 经营、拍卖单位地址
	AuctionUnitAddress *string `json:"auctionUnitAddress,omitempty" xml:"auctionUnitAddress,omitempty"`
	// 经营、拍卖单位银行
	AuctionUnitBank *string `json:"auctionUnitBank,omitempty" xml:"auctionUnitBank,omitempty"`
	// 经营、拍卖单位税号
	AuctionUnitTaxNo *string `json:"auctionUnitTaxNo,omitempty" xml:"auctionUnitTaxNo,omitempty"`
	// 经营、拍卖单位电话
	AuctionUtilTel *string `json:"auctionUtilTel,omitempty" xml:"auctionUtilTel,omitempty"`
	// 厂牌型号
	CarModel *string `json:"carModel,omitempty" xml:"carModel,omitempty"`
	// 车牌照号
	CardNo *string `json:"cardNo,omitempty" xml:"cardNo,omitempty"`
	// 登记证号
	Registration *string `json:"registration,omitempty" xml:"registration,omitempty"`
	// 转入地车辆管理所名称
	TransferVehicle *string `json:"transferVehicle,omitempty" xml:"transferVehicle,omitempty"`
	// 二手车市场地址
	UsedCarAddress *string `json:"usedCarAddress,omitempty" xml:"usedCarAddress,omitempty"`
	// 二手车市场
	UsedCarMarket *string `json:"usedCarMarket,omitempty" xml:"usedCarMarket,omitempty"`
	// 二手车市场开户银行、账号
	UsedCarMarketBank *string `json:"usedCarMarketBank,omitempty" xml:"usedCarMarketBank,omitempty"`
	// 二手车市场电话
	UsedCarMarketPhone *string `json:"usedCarMarketPhone,omitempty" xml:"usedCarMarketPhone,omitempty"`
	// 二手车市场纳税人识别号
	UsedCarMarketTaxNo *string `json:"usedCarMarketTaxNo,omitempty" xml:"usedCarMarketTaxNo,omitempty"`
	// 车架号/车辆识别号
	VehicleNo *string `json:"vehicleNo,omitempty" xml:"vehicleNo,omitempty"`
	// 车辆类型
	VehicleType *string `json:"vehicleType,omitempty" xml:"vehicleType,omitempty"`
}

func (s BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) String() string {
	return tea.Prettify(s)
}

func (s BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) GoString() string {
	return s.String()
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetAuctionUnit(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.AuctionUnit = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetAuctionUnitAddress(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.AuctionUnitAddress = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetAuctionUnitBank(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.AuctionUnitBank = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetAuctionUnitTaxNo(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.AuctionUnitTaxNo = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetAuctionUtilTel(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.AuctionUtilTel = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetCarModel(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.CarModel = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetCardNo(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.CardNo = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetRegistration(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.Registration = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetTransferVehicle(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.TransferVehicle = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetUsedCarAddress(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.UsedCarAddress = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetUsedCarMarket(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.UsedCarMarket = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetUsedCarMarketBank(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.UsedCarMarketBank = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetUsedCarMarketPhone(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.UsedCarMarketPhone = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetUsedCarMarketTaxNo(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.UsedCarMarketTaxNo = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetVehicleNo(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.VehicleNo = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetVehicleType(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.VehicleType = &v
	return s
}

type BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList struct {
	// 品牌
	Brand *string `json:"brand,omitempty" xml:"brand,omitempty"`
	// 合格证号
	CertificateNo *string `json:"certificateNo,omitempty" xml:"certificateNo,omitempty"`
	// 发动机号
	EngineNo *string `json:"engineNo,omitempty" xml:"engineNo,omitempty"`
	// 身份证号/组织机构代码
	IdCardNo *string `json:"idCardNo,omitempty" xml:"idCardNo,omitempty"`
	// 进口证书号
	ImportCertificateNo *string `json:"importCertificateNo,omitempty" xml:"importCertificateNo,omitempty"`
	// 商检单号
	InspectionListNo *string `json:"inspectionListNo,omitempty" xml:"inspectionListNo,omitempty"`
	// 限乘人数
	MaxPassengers *string `json:"maxPassengers,omitempty" xml:"maxPassengers,omitempty"`
	// 产地
	OriginPlace *string `json:"originPlace,omitempty" xml:"originPlace,omitempty"`
	// 完税凭证号码
	PaymentVoucherNo *string `json:"paymentVoucherNo,omitempty" xml:"paymentVoucherNo,omitempty"`
	// 主管税务机关名称
	TaxAuthorityName *string `json:"taxAuthorityName,omitempty" xml:"taxAuthorityName,omitempty"`
	// 主管税务机关代码
	TaxAuthorityNo *string `json:"taxAuthorityNo,omitempty" xml:"taxAuthorityNo,omitempty"`
	// 税率
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// 吨位
	Tonnage *string `json:"tonnage,omitempty" xml:"tonnage,omitempty"`
	// 车架号码
	VehicleNo *string `json:"vehicleNo,omitempty" xml:"vehicleNo,omitempty"`
	// 车辆类型
	VehicleType *string `json:"vehicleType,omitempty" xml:"vehicleType,omitempty"`
}

func (s BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList) String() string {
	return tea.Prettify(s)
}

func (s BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList) GoString() string {
	return s.String()
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetBrand(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.Brand = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetCertificateNo(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.CertificateNo = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetEngineNo(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.EngineNo = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetIdCardNo(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.IdCardNo = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetImportCertificateNo(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.ImportCertificateNo = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetInspectionListNo(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.InspectionListNo = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetMaxPassengers(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.MaxPassengers = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetOriginPlace(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.OriginPlace = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetPaymentVoucherNo(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.PaymentVoucherNo = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetTaxAuthorityName(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.TaxAuthorityName = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetTaxAuthorityNo(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.TaxAuthorityNo = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetTaxRate(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.TaxRate = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetTonnage(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.Tonnage = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetVehicleNo(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.VehicleNo = &v
	return s
}

func (s *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetVehicleType(v string) *BatchAddInvoiceRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.VehicleType = &v
	return s
}

type BatchAddInvoiceResponseBody struct {
	// 错误信息
	Result []*BatchAddInvoiceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s BatchAddInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchAddInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *BatchAddInvoiceResponseBody) SetResult(v []*BatchAddInvoiceResponseBodyResult) *BatchAddInvoiceResponseBody {
	s.Result = v
	return s
}

type BatchAddInvoiceResponseBodyResult struct {
	// 错误数据的key
	ErrorKey *string `json:"errorKey,omitempty" xml:"errorKey,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
}

func (s BatchAddInvoiceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s BatchAddInvoiceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *BatchAddInvoiceResponseBodyResult) SetErrorKey(v string) *BatchAddInvoiceResponseBodyResult {
	s.ErrorKey = &v
	return s
}

func (s *BatchAddInvoiceResponseBodyResult) SetErrorMsg(v string) *BatchAddInvoiceResponseBodyResult {
	s.ErrorMsg = &v
	return s
}

type BatchAddInvoiceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchAddInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchAddInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchAddInvoiceResponse) GoString() string {
	return s.String()
}

func (s *BatchAddInvoiceResponse) SetHeaders(v map[string]*string) *BatchAddInvoiceResponse {
	s.Headers = v
	return s
}

func (s *BatchAddInvoiceResponse) SetBody(v *BatchAddInvoiceResponseBody) *BatchAddInvoiceResponse {
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
	// 结束时间
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 进项发票/销项发票
	FinanceType *string `json:"financeType,omitempty" xml:"financeType,omitempty"`
	// 发票编码
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	// 发票号码
	InvoiceNo *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
	// 页号
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 当前页大小
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 税号
	TaxNo *string `json:"taxNo,omitempty" xml:"taxNo,omitempty"`
	// 发票认证状态
	VerifyStatus *string `json:"verifyStatus,omitempty" xml:"verifyStatus,omitempty"`
}

func (s CheckVoucherStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckVoucherStatusRequest) GoString() string {
	return s.String()
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckVoucherStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CheckVoucherStatusResponse) SetBody(v *CheckVoucherStatusResponseBody) *CheckVoucherStatusResponse {
	s.Body = v
	return s
}

type CreateCustomerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateCustomerHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerHeaders) GoString() string {
	return s.String()
}

func (s *CreateCustomerHeaders) SetCommonHeaders(v map[string]*string) *CreateCustomerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateCustomerHeaders) SetXAcsDingtalkAccessToken(v string) *CreateCustomerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateCustomerRequest struct {
	// 创建人userId
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 客户描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 客户名字
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 购方账户
	PurchaserAccount *string `json:"purchaserAccount,omitempty" xml:"purchaserAccount,omitempty"`
	// 购房地址
	PurchaserAddress *string `json:"purchaserAddress,omitempty" xml:"purchaserAddress,omitempty"`
	// 购方银行
	PurchaserBankName *string `json:"purchaserBankName,omitempty" xml:"purchaserBankName,omitempty"`
	// 购方名字
	PurchaserName *string `json:"purchaserName,omitempty" xml:"purchaserName,omitempty"`
	// 购方税号
	PurchaserTaxNo *string `json:"purchaserTaxNo,omitempty" xml:"purchaserTaxNo,omitempty"`
	// 购方电话
	PurchaserTel *string `json:"purchaserTel,omitempty" xml:"purchaserTel,omitempty"`
}

func (s CreateCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomerRequest) SetCreator(v string) *CreateCustomerRequest {
	s.Creator = &v
	return s
}

func (s *CreateCustomerRequest) SetDescription(v string) *CreateCustomerRequest {
	s.Description = &v
	return s
}

func (s *CreateCustomerRequest) SetName(v string) *CreateCustomerRequest {
	s.Name = &v
	return s
}

func (s *CreateCustomerRequest) SetPurchaserAccount(v string) *CreateCustomerRequest {
	s.PurchaserAccount = &v
	return s
}

func (s *CreateCustomerRequest) SetPurchaserAddress(v string) *CreateCustomerRequest {
	s.PurchaserAddress = &v
	return s
}

func (s *CreateCustomerRequest) SetPurchaserBankName(v string) *CreateCustomerRequest {
	s.PurchaserBankName = &v
	return s
}

func (s *CreateCustomerRequest) SetPurchaserName(v string) *CreateCustomerRequest {
	s.PurchaserName = &v
	return s
}

func (s *CreateCustomerRequest) SetPurchaserTaxNo(v string) *CreateCustomerRequest {
	s.PurchaserTaxNo = &v
	return s
}

func (s *CreateCustomerRequest) SetPurchaserTel(v string) *CreateCustomerRequest {
	s.PurchaserTel = &v
	return s
}

type CreateCustomerResponseBody struct {
	// 客户CODE
	CustomerCode *string `json:"customerCode,omitempty" xml:"customerCode,omitempty"`
}

func (s CreateCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomerResponseBody) SetCustomerCode(v string) *CreateCustomerResponseBody {
	s.CustomerCode = &v
	return s
}

type CreateCustomerResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomerResponse) SetHeaders(v map[string]*string) *CreateCustomerResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomerResponse) SetBody(v *CreateCustomerResponseBody) *CreateCustomerResponse {
	s.Body = v
	return s
}

type CreateReceiptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateReceiptHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateReceiptHeaders) GoString() string {
	return s.String()
}

func (s *CreateReceiptHeaders) SetCommonHeaders(v map[string]*string) *CreateReceiptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateReceiptHeaders) SetXAcsDingtalkAccessToken(v string) *CreateReceiptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateReceiptRequest struct {
	// 单据列表，不超过10条数据
	Receipts []*CreateReceiptRequestReceipts `json:"receipts,omitempty" xml:"receipts,omitempty" type:"Repeated"`
}

func (s CreateReceiptRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateReceiptRequest) GoString() string {
	return s.String()
}

func (s *CreateReceiptRequest) SetReceipts(v []*CreateReceiptRequestReceipts) *CreateReceiptRequest {
	s.Receipts = v
	return s
}

type CreateReceiptRequestReceipts struct {
	// 单据金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 关联收支类别code
	CategoryCode *string `json:"categoryCode,omitempty" xml:"categoryCode,omitempty"`
	// 单据唯一编号
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 单据创建时间，默认当前时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建人工号
	CreateUserId *string `json:"createUserId,omitempty" xml:"createUserId,omitempty"`
	// 关联客户code
	CustomerCode *string `json:"customerCode,omitempty" xml:"customerCode,omitempty"`
	// 关联企业账户code
	EnterpriseAcountCode *string `json:"enterpriseAcountCode,omitempty" xml:"enterpriseAcountCode,omitempty"`
	// 业务发生时间，默认当前时间
	OccurDate *int64 `json:"occurDate,omitempty" xml:"occurDate,omitempty"`
	// 负责人工号，传空代表不修改
	PrincipalId *string `json:"principalId,omitempty" xml:"principalId,omitempty"`
	// 关联项目code
	ProjectCode *string `json:"projectCode,omitempty" xml:"projectCode,omitempty"`
	// 单据类型：1付款单，2收款单
	ReceiptType *int64 `json:"receiptType,omitempty" xml:"receiptType,omitempty"`
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 关联供应商code
	SupplierCode *string `json:"supplierCode,omitempty" xml:"supplierCode,omitempty"`
	// 单据标题，不传由系统默认生成
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateReceiptRequestReceipts) String() string {
	return tea.Prettify(s)
}

func (s CreateReceiptRequestReceipts) GoString() string {
	return s.String()
}

func (s *CreateReceiptRequestReceipts) SetAmount(v string) *CreateReceiptRequestReceipts {
	s.Amount = &v
	return s
}

func (s *CreateReceiptRequestReceipts) SetCategoryCode(v string) *CreateReceiptRequestReceipts {
	s.CategoryCode = &v
	return s
}

func (s *CreateReceiptRequestReceipts) SetCode(v string) *CreateReceiptRequestReceipts {
	s.Code = &v
	return s
}

func (s *CreateReceiptRequestReceipts) SetCreateTime(v int64) *CreateReceiptRequestReceipts {
	s.CreateTime = &v
	return s
}

func (s *CreateReceiptRequestReceipts) SetCreateUserId(v string) *CreateReceiptRequestReceipts {
	s.CreateUserId = &v
	return s
}

func (s *CreateReceiptRequestReceipts) SetCustomerCode(v string) *CreateReceiptRequestReceipts {
	s.CustomerCode = &v
	return s
}

func (s *CreateReceiptRequestReceipts) SetEnterpriseAcountCode(v string) *CreateReceiptRequestReceipts {
	s.EnterpriseAcountCode = &v
	return s
}

func (s *CreateReceiptRequestReceipts) SetOccurDate(v int64) *CreateReceiptRequestReceipts {
	s.OccurDate = &v
	return s
}

func (s *CreateReceiptRequestReceipts) SetPrincipalId(v string) *CreateReceiptRequestReceipts {
	s.PrincipalId = &v
	return s
}

func (s *CreateReceiptRequestReceipts) SetProjectCode(v string) *CreateReceiptRequestReceipts {
	s.ProjectCode = &v
	return s
}

func (s *CreateReceiptRequestReceipts) SetReceiptType(v int64) *CreateReceiptRequestReceipts {
	s.ReceiptType = &v
	return s
}

func (s *CreateReceiptRequestReceipts) SetRemark(v string) *CreateReceiptRequestReceipts {
	s.Remark = &v
	return s
}

func (s *CreateReceiptRequestReceipts) SetSupplierCode(v string) *CreateReceiptRequestReceipts {
	s.SupplierCode = &v
	return s
}

func (s *CreateReceiptRequestReceipts) SetTitle(v string) *CreateReceiptRequestReceipts {
	s.Title = &v
	return s
}

type CreateReceiptResponseBody struct {
	// 结果列表
	Results []*CreateReceiptResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s CreateReceiptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateReceiptResponseBody) GoString() string {
	return s.String()
}

func (s *CreateReceiptResponseBody) SetResults(v []*CreateReceiptResponseBodyResults) *CreateReceiptResponseBody {
	s.Results = v
	return s
}

type CreateReceiptResponseBodyResults struct {
	// 数据唯一编号
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateReceiptResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s CreateReceiptResponseBodyResults) GoString() string {
	return s.String()
}

func (s *CreateReceiptResponseBodyResults) SetCode(v string) *CreateReceiptResponseBodyResults {
	s.Code = &v
	return s
}

func (s *CreateReceiptResponseBodyResults) SetErrorCode(v string) *CreateReceiptResponseBodyResults {
	s.ErrorCode = &v
	return s
}

func (s *CreateReceiptResponseBodyResults) SetErrorMsg(v string) *CreateReceiptResponseBodyResults {
	s.ErrorMsg = &v
	return s
}

func (s *CreateReceiptResponseBodyResults) SetSuccess(v bool) *CreateReceiptResponseBodyResults {
	s.Success = &v
	return s
}

type CreateReceiptResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateReceiptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateReceiptResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateReceiptResponse) GoString() string {
	return s.String()
}

func (s *CreateReceiptResponse) SetHeaders(v map[string]*string) *CreateReceiptResponse {
	s.Headers = v
	return s
}

func (s *CreateReceiptResponse) SetBody(v *CreateReceiptResponseBody) *CreateReceiptResponse {
	s.Body = v
	return s
}

type DeleteReceiptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteReceiptHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteReceiptHeaders) GoString() string {
	return s.String()
}

func (s *DeleteReceiptHeaders) SetCommonHeaders(v map[string]*string) *DeleteReceiptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteReceiptHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteReceiptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteReceiptRequest struct {
	// 单据列表，最长不超过10条
	Receipts []*DeleteReceiptRequestReceipts `json:"receipts,omitempty" xml:"receipts,omitempty" type:"Repeated"`
}

func (s DeleteReceiptRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteReceiptRequest) GoString() string {
	return s.String()
}

func (s *DeleteReceiptRequest) SetReceipts(v []*DeleteReceiptRequestReceipts) *DeleteReceiptRequest {
	s.Receipts = v
	return s
}

type DeleteReceiptRequestReceipts struct {
	// 单据唯一编号
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 修改者工号
	DeleteUserId *string `json:"deleteUserId,omitempty" xml:"deleteUserId,omitempty"`
	// 单据类型：1付款单，2收款单
	ReceiptType *int64 `json:"receiptType,omitempty" xml:"receiptType,omitempty"`
}

func (s DeleteReceiptRequestReceipts) String() string {
	return tea.Prettify(s)
}

func (s DeleteReceiptRequestReceipts) GoString() string {
	return s.String()
}

func (s *DeleteReceiptRequestReceipts) SetCode(v string) *DeleteReceiptRequestReceipts {
	s.Code = &v
	return s
}

func (s *DeleteReceiptRequestReceipts) SetDeleteUserId(v string) *DeleteReceiptRequestReceipts {
	s.DeleteUserId = &v
	return s
}

func (s *DeleteReceiptRequestReceipts) SetReceiptType(v int64) *DeleteReceiptRequestReceipts {
	s.ReceiptType = &v
	return s
}

type DeleteReceiptResponseBody struct {
	// 结果列表
	Results []*DeleteReceiptResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s DeleteReceiptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteReceiptResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteReceiptResponseBody) SetResults(v []*DeleteReceiptResponseBodyResults) *DeleteReceiptResponseBody {
	s.Results = v
	return s
}

type DeleteReceiptResponseBodyResults struct {
	// 数据唯一编号
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteReceiptResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s DeleteReceiptResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DeleteReceiptResponseBodyResults) SetCode(v string) *DeleteReceiptResponseBodyResults {
	s.Code = &v
	return s
}

func (s *DeleteReceiptResponseBodyResults) SetErrorCode(v string) *DeleteReceiptResponseBodyResults {
	s.ErrorCode = &v
	return s
}

func (s *DeleteReceiptResponseBodyResults) SetErrorMsg(v string) *DeleteReceiptResponseBodyResults {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteReceiptResponseBodyResults) SetSuccess(v bool) *DeleteReceiptResponseBodyResults {
	s.Success = &v
	return s
}

type DeleteReceiptResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteReceiptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteReceiptResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteReceiptResponse) GoString() string {
	return s.String()
}

func (s *DeleteReceiptResponse) SetHeaders(v map[string]*string) *DeleteReceiptResponse {
	s.Headers = v
	return s
}

func (s *DeleteReceiptResponse) SetBody(v *DeleteReceiptResponseBody) *DeleteReceiptResponse {
	s.Body = v
	return s
}

type GetBookkeepingUserListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetBookkeepingUserListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetBookkeepingUserListHeaders) GoString() string {
	return s.String()
}

func (s *GetBookkeepingUserListHeaders) SetCommonHeaders(v map[string]*string) *GetBookkeepingUserListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetBookkeepingUserListHeaders) SetXAcsDingtalkAccessToken(v string) *GetBookkeepingUserListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetBookkeepingUserListResponseBody struct {
	// staffId列表
	Result []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetBookkeepingUserListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBookkeepingUserListResponseBody) GoString() string {
	return s.String()
}

func (s *GetBookkeepingUserListResponseBody) SetResult(v []*string) *GetBookkeepingUserListResponseBody {
	s.Result = v
	return s
}

type GetBookkeepingUserListResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetBookkeepingUserListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBookkeepingUserListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBookkeepingUserListResponse) GoString() string {
	return s.String()
}

func (s *GetBookkeepingUserListResponse) SetHeaders(v map[string]*string) *GetBookkeepingUserListResponse {
	s.Headers = v
	return s
}

func (s *GetBookkeepingUserListResponse) SetBody(v *GetBookkeepingUserListResponseBody) *GetBookkeepingUserListResponse {
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
	// 类别code
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
	// 类别code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否为目录
	IsDir *bool `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// 名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 父类别code
	ParentCode *string `json:"parentCode,omitempty" xml:"parentCode,omitempty"`
	// 状态:valid,invalid,deleted
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 类型：income收入，expense支出
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCategoryResponseBody) GoString() string {
	return s.String()
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

func (s *GetCategoryResponseBody) SetStatus(v string) *GetCategoryResponseBody {
	s.Status = &v
	return s
}

func (s *GetCategoryResponseBody) SetType(v string) *GetCategoryResponseBody {
	s.Type = &v
	return s
}

type GetCategoryResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetCategoryResponse) SetBody(v *GetCategoryResponseBody) *GetCategoryResponse {
	s.Body = v
	return s
}

type GetCustomerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCustomerHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerHeaders) GoString() string {
	return s.String()
}

func (s *GetCustomerHeaders) SetCommonHeaders(v map[string]*string) *GetCustomerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCustomerHeaders) SetXAcsDingtalkAccessToken(v string) *GetCustomerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCustomerRequest struct {
	// 客户code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s GetCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerRequest) GoString() string {
	return s.String()
}

func (s *GetCustomerRequest) SetCode(v string) *GetCustomerRequest {
	s.Code = &v
	return s
}

type GetCustomerResponseBody struct {
	// 客户Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 创建时间(单位MS)
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 客户描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 客户名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 状态：启用(valid), 停用(invalid), 删除(deleted)
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 用户自定义code
	UserDefineCode *string `json:"userDefineCode,omitempty" xml:"userDefineCode,omitempty"`
}

func (s GetCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomerResponseBody) SetCode(v string) *GetCustomerResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomerResponseBody) SetCreateTime(v int64) *GetCustomerResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetCustomerResponseBody) SetDescription(v string) *GetCustomerResponseBody {
	s.Description = &v
	return s
}

func (s *GetCustomerResponseBody) SetName(v string) *GetCustomerResponseBody {
	s.Name = &v
	return s
}

func (s *GetCustomerResponseBody) SetStatus(v string) *GetCustomerResponseBody {
	s.Status = &v
	return s
}

func (s *GetCustomerResponseBody) SetUserDefineCode(v string) *GetCustomerResponseBody {
	s.UserDefineCode = &v
	return s
}

type GetCustomerResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerResponse) GoString() string {
	return s.String()
}

func (s *GetCustomerResponse) SetHeaders(v map[string]*string) *GetCustomerResponse {
	s.Headers = v
	return s
}

func (s *GetCustomerResponse) SetBody(v *GetCustomerResponseBody) *GetCustomerResponse {
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
	// 账户code
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
	// 账户code
	AccountCode *string `json:"accountCode,omitempty" xml:"accountCode,omitempty"`
	// 关联资金账户id
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// 账户名称
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	// 备注
	AccountRemark *string `json:"accountRemark,omitempty" xml:"accountRemark,omitempty"`
	// 账户类型:ALIPAY, BANKCARD, CASH, WECHAT
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	// 账户总额，保留2位小数
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 银行代号，如果是银行卡类型，有值，其他类型时，为空
	BankCode *string `json:"bankCode,omitempty" xml:"bankCode,omitempty"`
	// 银行名称，如果是银行卡类型，有值，其他类型时，为空
	BankName *string `json:"bankName,omitempty" xml:"bankName,omitempty"`
	// 创建时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建人工号
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
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

type GetFinanceAccountResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFinanceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetFinanceAccountResponse) SetBody(v *GetFinanceAccountResponseBody) *GetFinanceAccountResponse {
	s.Body = v
	return s
}

type GetInvoiceByPageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetInvoiceByPageHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetInvoiceByPageHeaders) GoString() string {
	return s.String()
}

func (s *GetInvoiceByPageHeaders) SetCommonHeaders(v map[string]*string) *GetInvoiceByPageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInvoiceByPageHeaders) SetXAcsDingtalkAccessToken(v string) *GetInvoiceByPageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetInvoiceByPageRequest struct {
	Request *GetInvoiceByPageRequestRequest `json:"request,omitempty" xml:"request,omitempty" type:"Struct"`
}

func (s GetInvoiceByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInvoiceByPageRequest) GoString() string {
	return s.String()
}

func (s *GetInvoiceByPageRequest) SetRequest(v *GetInvoiceByPageRequestRequest) *GetInvoiceByPageRequest {
	s.Request = v
	return s
}

type GetInvoiceByPageRequestRequest struct {
	// 结束时间
	//
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 进项票/销项票
	FinanceType *string `json:"financeType,omitempty" xml:"financeType,omitempty"`
	// 分页参数
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页参数
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 税号
	TaxNo *string `json:"taxNo,omitempty" xml:"taxNo,omitempty"`
	// 认证状态
	VerifyStatus *string `json:"verifyStatus,omitempty" xml:"verifyStatus,omitempty"`
}

func (s GetInvoiceByPageRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInvoiceByPageRequestRequest) GoString() string {
	return s.String()
}

func (s *GetInvoiceByPageRequestRequest) SetEndTime(v int64) *GetInvoiceByPageRequestRequest {
	s.EndTime = &v
	return s
}

func (s *GetInvoiceByPageRequestRequest) SetFinanceType(v string) *GetInvoiceByPageRequestRequest {
	s.FinanceType = &v
	return s
}

func (s *GetInvoiceByPageRequestRequest) SetPageNumber(v int64) *GetInvoiceByPageRequestRequest {
	s.PageNumber = &v
	return s
}

func (s *GetInvoiceByPageRequestRequest) SetPageSize(v int64) *GetInvoiceByPageRequestRequest {
	s.PageSize = &v
	return s
}

func (s *GetInvoiceByPageRequestRequest) SetStartTime(v int64) *GetInvoiceByPageRequestRequest {
	s.StartTime = &v
	return s
}

func (s *GetInvoiceByPageRequestRequest) SetTaxNo(v string) *GetInvoiceByPageRequestRequest {
	s.TaxNo = &v
	return s
}

func (s *GetInvoiceByPageRequestRequest) SetVerifyStatus(v string) *GetInvoiceByPageRequestRequest {
	s.VerifyStatus = &v
	return s
}

type GetInvoiceByPageShrinkRequest struct {
	RequestShrink *string `json:"request,omitempty" xml:"request,omitempty"`
}

func (s GetInvoiceByPageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInvoiceByPageShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetInvoiceByPageShrinkRequest) SetRequestShrink(v string) *GetInvoiceByPageShrinkRequest {
	s.RequestShrink = &v
	return s
}

type GetInvoiceByPageResponseBody struct {
	ErrorCode *string                             `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                             `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result    *GetInvoiceByPageResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success   *bool                               `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetInvoiceByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInvoiceByPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetInvoiceByPageResponseBody) SetErrorCode(v string) *GetInvoiceByPageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetInvoiceByPageResponseBody) SetErrorMsg(v string) *GetInvoiceByPageResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetInvoiceByPageResponseBody) SetResult(v *GetInvoiceByPageResponseBodyResult) *GetInvoiceByPageResponseBody {
	s.Result = v
	return s
}

func (s *GetInvoiceByPageResponseBody) SetSuccess(v bool) *GetInvoiceByPageResponseBody {
	s.Success = &v
	return s
}

type GetInvoiceByPageResponseBodyResult struct {
	HasMore    *string                                   `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*GetInvoiceByPageResponseBodyResultList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextCursor *int64                                    `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	TotalCount *int64                                    `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetInvoiceByPageResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetInvoiceByPageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetInvoiceByPageResponseBodyResult) SetHasMore(v string) *GetInvoiceByPageResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResult) SetList(v []*GetInvoiceByPageResponseBodyResultList) *GetInvoiceByPageResponseBodyResult {
	s.List = v
	return s
}

func (s *GetInvoiceByPageResponseBodyResult) SetNextCursor(v int64) *GetInvoiceByPageResponseBodyResult {
	s.NextCursor = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResult) SetTotalCount(v int64) *GetInvoiceByPageResponseBodyResult {
	s.TotalCount = &v
	return s
}

type GetInvoiceByPageResponseBodyResultList struct {
	// 账期时间
	AccountPeriod *string `json:"accountPeriod,omitempty" xml:"accountPeriod,omitempty"`
	// 不含税金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 含税金额
	AmountWithTax *string `json:"amountWithTax,omitempty" xml:"amountWithTax,omitempty"`
	// 校验码
	CheckCode *string `json:"checkCode,omitempty" xml:"checkCode,omitempty"`
	// 查验时间
	CheckTime *string `json:"checkTime,omitempty" xml:"checkTime,omitempty"`
	// 开票日期
	DrewDate *string `json:"drewDate,omitempty" xml:"drewDate,omitempty"`
	// 电票版式文件下载地址
	ElectronicUrl *string `json:"electronicUrl,omitempty" xml:"electronicUrl,omitempty"`
	// 财务类型，INPUT-VAT(进项),OUTPUT_VAT(销项)
	FinanceType *string `json:"financeType,omitempty" xml:"financeType,omitempty"`
	// 资金类型 ，RED（红票），（BLUE）蓝票
	FundType *string `json:"fundType,omitempty" xml:"fundType,omitempty"`
	// 常规发票明细
	GeneralInvoiceDetailVOList []*GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList `json:"generalInvoiceDetailVOList,omitempty" xml:"generalInvoiceDetailVOList,omitempty" type:"Repeated"`
	// 发票代码
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	// 发票号码
	InvoiceNo *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
	// 发票状态
	InvoiceStatus *string `json:"invoiceStatus,omitempty" xml:"invoiceStatus,omitempty"`
	// 发票类型
	InvoiceType *string `json:"invoiceType,omitempty" xml:"invoiceType,omitempty"`
	// 机器码
	MachineCode *string `json:"machineCode,omitempty" xml:"machineCode,omitempty"`
	// 成品油标识
	OilFlag *string `json:"oilFlag,omitempty" xml:"oilFlag,omitempty"`
	// 收款人
	Payee *string `json:"payee,omitempty" xml:"payee,omitempty"`
	// 审批单实例
	ProcessInstCode *string `json:"processInstCode,omitempty" xml:"processInstCode,omitempty"`
	// 审批单类型
	ProcessInstType *string `json:"processInstType,omitempty" xml:"processInstType,omitempty"`
	// 购方地址
	PurchaserAddress *string `json:"purchaserAddress,omitempty" xml:"purchaserAddress,omitempty"`
	// 购方银行
	PurchaserBankNameAccount *string `json:"purchaserBankNameAccount,omitempty" xml:"purchaserBankNameAccount,omitempty"`
	// 购方名称
	PurchaserName *string `json:"purchaserName,omitempty" xml:"purchaserName,omitempty"`
	// 购方税号
	PurchaserTaxNo *string `json:"purchaserTaxNo,omitempty" xml:"purchaserTaxNo,omitempty"`
	// 购方电话
	PurchaserTel *string `json:"purchaserTel,omitempty" xml:"purchaserTel,omitempty"`
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 销方地址
	SellerAddress *string `json:"sellerAddress,omitempty" xml:"sellerAddress,omitempty"`
	// 销方银行
	SellerBankNameAccount *string `json:"sellerBankNameAccount,omitempty" xml:"sellerBankNameAccount,omitempty"`
	// 销方名称
	SellerName *string `json:"sellerName,omitempty" xml:"sellerName,omitempty"`
	// 销方税号
	SellerTaxNo *string `json:"sellerTaxNo,omitempty" xml:"sellerTaxNo,omitempty"`
	// 销方电话
	SellerTel *string `json:"sellerTel,omitempty" xml:"sellerTel,omitempty"`
	// 发票状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 代开发票标识 1-自开，2-代开
	SupplySign *string `json:"supplySign,omitempty" xml:"supplySign,omitempty"`
	// 税额
	TaxAmount                   *string                                                              `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	TransportFeeDetailVOList    []*GetInvoiceByPageResponseBodyResultListTransportFeeDetailVOList    `json:"transportFeeDetailVOList,omitempty" xml:"transportFeeDetailVOList,omitempty" type:"Repeated"`
	UsedVehicleSaleDetailVOList []*GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList `json:"usedVehicleSaleDetailVOList,omitempty" xml:"usedVehicleSaleDetailVOList,omitempty" type:"Repeated"`
	VehicleSaleDetailVOList     []*GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList     `json:"vehicleSaleDetailVOList,omitempty" xml:"vehicleSaleDetailVOList,omitempty" type:"Repeated"`
	// 发票查验状态
	VerifyStatus *string `json:"verifyStatus,omitempty" xml:"verifyStatus,omitempty"`
	// 凭证code
	VoucherCode *string `json:"voucherCode,omitempty" xml:"voucherCode,omitempty"`
	// 生成凭证状态
	VoucherStatus *string `json:"voucherStatus,omitempty" xml:"voucherStatus,omitempty"`
}

func (s GetInvoiceByPageResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s GetInvoiceByPageResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *GetInvoiceByPageResponseBodyResultList) SetAccountPeriod(v string) *GetInvoiceByPageResponseBodyResultList {
	s.AccountPeriod = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetAmount(v string) *GetInvoiceByPageResponseBodyResultList {
	s.Amount = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetAmountWithTax(v string) *GetInvoiceByPageResponseBodyResultList {
	s.AmountWithTax = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetCheckCode(v string) *GetInvoiceByPageResponseBodyResultList {
	s.CheckCode = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetCheckTime(v string) *GetInvoiceByPageResponseBodyResultList {
	s.CheckTime = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetDrewDate(v string) *GetInvoiceByPageResponseBodyResultList {
	s.DrewDate = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetElectronicUrl(v string) *GetInvoiceByPageResponseBodyResultList {
	s.ElectronicUrl = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetFinanceType(v string) *GetInvoiceByPageResponseBodyResultList {
	s.FinanceType = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetFundType(v string) *GetInvoiceByPageResponseBodyResultList {
	s.FundType = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetGeneralInvoiceDetailVOList(v []*GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList) *GetInvoiceByPageResponseBodyResultList {
	s.GeneralInvoiceDetailVOList = v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetInvoiceCode(v string) *GetInvoiceByPageResponseBodyResultList {
	s.InvoiceCode = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetInvoiceNo(v string) *GetInvoiceByPageResponseBodyResultList {
	s.InvoiceNo = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetInvoiceStatus(v string) *GetInvoiceByPageResponseBodyResultList {
	s.InvoiceStatus = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetInvoiceType(v string) *GetInvoiceByPageResponseBodyResultList {
	s.InvoiceType = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetMachineCode(v string) *GetInvoiceByPageResponseBodyResultList {
	s.MachineCode = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetOilFlag(v string) *GetInvoiceByPageResponseBodyResultList {
	s.OilFlag = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetPayee(v string) *GetInvoiceByPageResponseBodyResultList {
	s.Payee = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetProcessInstCode(v string) *GetInvoiceByPageResponseBodyResultList {
	s.ProcessInstCode = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetProcessInstType(v string) *GetInvoiceByPageResponseBodyResultList {
	s.ProcessInstType = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetPurchaserAddress(v string) *GetInvoiceByPageResponseBodyResultList {
	s.PurchaserAddress = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetPurchaserBankNameAccount(v string) *GetInvoiceByPageResponseBodyResultList {
	s.PurchaserBankNameAccount = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetPurchaserName(v string) *GetInvoiceByPageResponseBodyResultList {
	s.PurchaserName = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetPurchaserTaxNo(v string) *GetInvoiceByPageResponseBodyResultList {
	s.PurchaserTaxNo = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetPurchaserTel(v string) *GetInvoiceByPageResponseBodyResultList {
	s.PurchaserTel = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetRemark(v string) *GetInvoiceByPageResponseBodyResultList {
	s.Remark = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetSellerAddress(v string) *GetInvoiceByPageResponseBodyResultList {
	s.SellerAddress = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetSellerBankNameAccount(v string) *GetInvoiceByPageResponseBodyResultList {
	s.SellerBankNameAccount = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetSellerName(v string) *GetInvoiceByPageResponseBodyResultList {
	s.SellerName = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetSellerTaxNo(v string) *GetInvoiceByPageResponseBodyResultList {
	s.SellerTaxNo = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetSellerTel(v string) *GetInvoiceByPageResponseBodyResultList {
	s.SellerTel = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetStatus(v string) *GetInvoiceByPageResponseBodyResultList {
	s.Status = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetSupplySign(v string) *GetInvoiceByPageResponseBodyResultList {
	s.SupplySign = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetTaxAmount(v string) *GetInvoiceByPageResponseBodyResultList {
	s.TaxAmount = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetTransportFeeDetailVOList(v []*GetInvoiceByPageResponseBodyResultListTransportFeeDetailVOList) *GetInvoiceByPageResponseBodyResultList {
	s.TransportFeeDetailVOList = v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetUsedVehicleSaleDetailVOList(v []*GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList) *GetInvoiceByPageResponseBodyResultList {
	s.UsedVehicleSaleDetailVOList = v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetVehicleSaleDetailVOList(v []*GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList) *GetInvoiceByPageResponseBodyResultList {
	s.VehicleSaleDetailVOList = v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetVerifyStatus(v string) *GetInvoiceByPageResponseBodyResultList {
	s.VerifyStatus = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetVoucherCode(v string) *GetInvoiceByPageResponseBodyResultList {
	s.VoucherCode = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultList) SetVoucherStatus(v string) *GetInvoiceByPageResponseBodyResultList {
	s.VoucherStatus = &v
	return s
}

type GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList struct {
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 商品名称
	GoodsName *string `json:"goodsName,omitempty" xml:"goodsName,omitempty"`
	// 数量
	Quantity *string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 税收分类编码
	RevenueCode *string `json:"revenueCode,omitempty" xml:"revenueCode,omitempty"`
	// 行号
	RowNo *string `json:"rowNo,omitempty" xml:"rowNo,omitempty"`
	// 规格型号
	Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
	// 税额
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	// 是否享受税收优惠：0-不享受，1-享受
	TaxPre *string `json:"taxPre,omitempty" xml:"taxPre,omitempty"`
	// 税率
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// 单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 单价
	UnitPrice *string `json:"unitPrice,omitempty" xml:"unitPrice,omitempty"`
}

func (s GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList) String() string {
	return tea.Prettify(s)
}

func (s GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList) GoString() string {
	return s.String()
}

func (s *GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList) SetAmount(v string) *GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList {
	s.Amount = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList) SetGoodsName(v string) *GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList {
	s.GoodsName = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList) SetQuantity(v string) *GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList {
	s.Quantity = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList) SetRevenueCode(v string) *GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList {
	s.RevenueCode = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList) SetRowNo(v string) *GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList {
	s.RowNo = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList) SetSpecification(v string) *GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList {
	s.Specification = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList) SetTaxAmount(v string) *GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList {
	s.TaxAmount = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList) SetTaxPre(v string) *GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList {
	s.TaxPre = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList) SetTaxRate(v string) *GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList {
	s.TaxRate = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList) SetUnit(v string) *GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList {
	s.Unit = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList) SetUnitPrice(v string) *GetInvoiceByPageResponseBodyResultListGeneralInvoiceDetailVOList {
	s.UnitPrice = &v
	return s
}

type GetInvoiceByPageResponseBodyResultListTransportFeeDetailVOList struct {
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 车牌号
	CardNo *string `json:"cardNo,omitempty" xml:"cardNo,omitempty"`
	// 通行日期止
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// 商品名称
	GoodsName *string `json:"goodsName,omitempty" xml:"goodsName,omitempty"`
	// 税收分类编码
	RevenueCode *string `json:"revenueCode,omitempty" xml:"revenueCode,omitempty"`
	// 行号
	RowNo *string `json:"rowNo,omitempty" xml:"rowNo,omitempty"`
	// 通行日期起
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// 税额
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	// 税率
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// 类型
	VehicleType *string `json:"vehicleType,omitempty" xml:"vehicleType,omitempty"`
}

func (s GetInvoiceByPageResponseBodyResultListTransportFeeDetailVOList) String() string {
	return tea.Prettify(s)
}

func (s GetInvoiceByPageResponseBodyResultListTransportFeeDetailVOList) GoString() string {
	return s.String()
}

func (s *GetInvoiceByPageResponseBodyResultListTransportFeeDetailVOList) SetAmount(v string) *GetInvoiceByPageResponseBodyResultListTransportFeeDetailVOList {
	s.Amount = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListTransportFeeDetailVOList) SetCardNo(v string) *GetInvoiceByPageResponseBodyResultListTransportFeeDetailVOList {
	s.CardNo = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListTransportFeeDetailVOList) SetEndDate(v string) *GetInvoiceByPageResponseBodyResultListTransportFeeDetailVOList {
	s.EndDate = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListTransportFeeDetailVOList) SetGoodsName(v string) *GetInvoiceByPageResponseBodyResultListTransportFeeDetailVOList {
	s.GoodsName = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListTransportFeeDetailVOList) SetRevenueCode(v string) *GetInvoiceByPageResponseBodyResultListTransportFeeDetailVOList {
	s.RevenueCode = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListTransportFeeDetailVOList) SetRowNo(v string) *GetInvoiceByPageResponseBodyResultListTransportFeeDetailVOList {
	s.RowNo = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListTransportFeeDetailVOList) SetStartDate(v string) *GetInvoiceByPageResponseBodyResultListTransportFeeDetailVOList {
	s.StartDate = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListTransportFeeDetailVOList) SetTaxAmount(v string) *GetInvoiceByPageResponseBodyResultListTransportFeeDetailVOList {
	s.TaxAmount = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListTransportFeeDetailVOList) SetTaxRate(v string) *GetInvoiceByPageResponseBodyResultListTransportFeeDetailVOList {
	s.TaxRate = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListTransportFeeDetailVOList) SetVehicleType(v string) *GetInvoiceByPageResponseBodyResultListTransportFeeDetailVOList {
	s.VehicleType = &v
	return s
}

type GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList struct {
	// 经营、拍卖单位
	AuctionUnit *string `json:"auctionUnit,omitempty" xml:"auctionUnit,omitempty"`
	// 经营、拍卖单位地址
	AuctionUnitAddress *string `json:"auctionUnitAddress,omitempty" xml:"auctionUnitAddress,omitempty"`
	// 经营、拍卖单位银行
	AuctionUnitBank *string `json:"auctionUnitBank,omitempty" xml:"auctionUnitBank,omitempty"`
	// 经营、拍卖单位税号
	AuctionUnitTaxNo *string `json:"auctionUnitTaxNo,omitempty" xml:"auctionUnitTaxNo,omitempty"`
	// 经营、拍卖单位电话
	AuctionUtilTel *string `json:"auctionUtilTel,omitempty" xml:"auctionUtilTel,omitempty"`
	// 厂牌型号
	CarModel *string `json:"carModel,omitempty" xml:"carModel,omitempty"`
	// 车牌照号
	CardNo *string `json:"cardNo,omitempty" xml:"cardNo,omitempty"`
	// 登记证号
	Registration *string `json:"registration,omitempty" xml:"registration,omitempty"`
	// 转入地车辆管理所名称
	TransferVehicle *string `json:"transferVehicle,omitempty" xml:"transferVehicle,omitempty"`
	// 二手车市场地址
	UsedCarAddress *string `json:"usedCarAddress,omitempty" xml:"usedCarAddress,omitempty"`
	// 二手车市场
	UsedCarMarket *string `json:"usedCarMarket,omitempty" xml:"usedCarMarket,omitempty"`
	// 二手车市场开户银行、账号
	UsedCarMarketBank *string `json:"usedCarMarketBank,omitempty" xml:"usedCarMarketBank,omitempty"`
	// 二手车市场电话
	UsedCarMarketPhone *string `json:"usedCarMarketPhone,omitempty" xml:"usedCarMarketPhone,omitempty"`
	// 二手车市场纳税人识别号
	UsedCarMarketTaxNo *string `json:"usedCarMarketTaxNo,omitempty" xml:"usedCarMarketTaxNo,omitempty"`
	// 车架号/车辆识别号
	VehicleNo *string `json:"vehicleNo,omitempty" xml:"vehicleNo,omitempty"`
	// 车辆类型
	VehicleType *string `json:"vehicleType,omitempty" xml:"vehicleType,omitempty"`
}

func (s GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList) String() string {
	return tea.Prettify(s)
}

func (s GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList) GoString() string {
	return s.String()
}

func (s *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList) SetAuctionUnit(v string) *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList {
	s.AuctionUnit = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList) SetAuctionUnitAddress(v string) *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList {
	s.AuctionUnitAddress = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList) SetAuctionUnitBank(v string) *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList {
	s.AuctionUnitBank = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList) SetAuctionUnitTaxNo(v string) *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList {
	s.AuctionUnitTaxNo = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList) SetAuctionUtilTel(v string) *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList {
	s.AuctionUtilTel = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList) SetCarModel(v string) *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList {
	s.CarModel = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList) SetCardNo(v string) *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList {
	s.CardNo = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList) SetRegistration(v string) *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList {
	s.Registration = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList) SetTransferVehicle(v string) *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList {
	s.TransferVehicle = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList) SetUsedCarAddress(v string) *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList {
	s.UsedCarAddress = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList) SetUsedCarMarket(v string) *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList {
	s.UsedCarMarket = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList) SetUsedCarMarketBank(v string) *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList {
	s.UsedCarMarketBank = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList) SetUsedCarMarketPhone(v string) *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList {
	s.UsedCarMarketPhone = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList) SetUsedCarMarketTaxNo(v string) *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList {
	s.UsedCarMarketTaxNo = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList) SetVehicleNo(v string) *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList {
	s.VehicleNo = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList) SetVehicleType(v string) *GetInvoiceByPageResponseBodyResultListUsedVehicleSaleDetailVOList {
	s.VehicleType = &v
	return s
}

type GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList struct {
	// 品牌
	Brand *string `json:"brand,omitempty" xml:"brand,omitempty"`
	// 合格证号
	CertificateNo *string `json:"certificateNo,omitempty" xml:"certificateNo,omitempty"`
	// 发动机号
	EngineNo *string `json:"engineNo,omitempty" xml:"engineNo,omitempty"`
	// 身份证号/组织机构代码
	IdCardNo *string `json:"idCardNo,omitempty" xml:"idCardNo,omitempty"`
	// 进口证书号
	ImportCertificateNo *string `json:"importCertificateNo,omitempty" xml:"importCertificateNo,omitempty"`
	// 限乘人数
	MaxPassengers *string `json:"maxPassengers,omitempty" xml:"maxPassengers,omitempty"`
	// 产地
	OriginPlace *string `json:"originPlace,omitempty" xml:"originPlace,omitempty"`
	// 完税凭证号码
	PaymentVoucherNo *string `json:"paymentVoucherNo,omitempty" xml:"paymentVoucherNo,omitempty"`
	// 主管税务机关名称
	TaxAuthorityName *string `json:"taxAuthorityName,omitempty" xml:"taxAuthorityName,omitempty"`
	// 主管税务机关代码
	TaxAuthorityNo *string `json:"taxAuthorityNo,omitempty" xml:"taxAuthorityNo,omitempty"`
	// 税率
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// 吨位
	Tonnage *string `json:"tonnage,omitempty" xml:"tonnage,omitempty"`
	// 车架号码
	VehicleNo *string `json:"vehicleNo,omitempty" xml:"vehicleNo,omitempty"`
	// 车辆类型
	VehicleType *string `json:"vehicleType,omitempty" xml:"vehicleType,omitempty"`
}

func (s GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList) String() string {
	return tea.Prettify(s)
}

func (s GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList) GoString() string {
	return s.String()
}

func (s *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList) SetBrand(v string) *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList {
	s.Brand = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList) SetCertificateNo(v string) *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList {
	s.CertificateNo = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList) SetEngineNo(v string) *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList {
	s.EngineNo = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList) SetIdCardNo(v string) *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList {
	s.IdCardNo = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList) SetImportCertificateNo(v string) *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList {
	s.ImportCertificateNo = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList) SetMaxPassengers(v string) *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList {
	s.MaxPassengers = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList) SetOriginPlace(v string) *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList {
	s.OriginPlace = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList) SetPaymentVoucherNo(v string) *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList {
	s.PaymentVoucherNo = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList) SetTaxAuthorityName(v string) *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList {
	s.TaxAuthorityName = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList) SetTaxAuthorityNo(v string) *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList {
	s.TaxAuthorityNo = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList) SetTaxRate(v string) *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList {
	s.TaxRate = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList) SetTonnage(v string) *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList {
	s.Tonnage = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList) SetVehicleNo(v string) *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList {
	s.VehicleNo = &v
	return s
}

func (s *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList) SetVehicleType(v string) *GetInvoiceByPageResponseBodyResultListVehicleSaleDetailVOList {
	s.VehicleType = &v
	return s
}

type GetInvoiceByPageResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInvoiceByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInvoiceByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInvoiceByPageResponse) GoString() string {
	return s.String()
}

func (s *GetInvoiceByPageResponse) SetHeaders(v map[string]*string) *GetInvoiceByPageResponse {
	s.Headers = v
	return s
}

func (s *GetInvoiceByPageResponse) SetBody(v *GetInvoiceByPageResponseBody) *GetInvoiceByPageResponse {
	s.Body = v
	return s
}

type GetIsNewVersionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetIsNewVersionHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetIsNewVersionHeaders) GoString() string {
	return s.String()
}

func (s *GetIsNewVersionHeaders) SetCommonHeaders(v map[string]*string) *GetIsNewVersionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetIsNewVersionHeaders) SetXAcsDingtalkAccessToken(v string) *GetIsNewVersionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetIsNewVersionResponseBody struct {
	Result  *bool `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetIsNewVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIsNewVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetIsNewVersionResponseBody) SetResult(v bool) *GetIsNewVersionResponseBody {
	s.Result = &v
	return s
}

func (s *GetIsNewVersionResponseBody) SetSuccess(v bool) *GetIsNewVersionResponseBody {
	s.Success = &v
	return s
}

type GetIsNewVersionResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetIsNewVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIsNewVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIsNewVersionResponse) GoString() string {
	return s.String()
}

func (s *GetIsNewVersionResponse) SetHeaders(v map[string]*string) *GetIsNewVersionResponse {
	s.Headers = v
	return s
}

func (s *GetIsNewVersionResponse) SetBody(v *GetIsNewVersionResponseBody) *GetIsNewVersionResponse {
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
	// 项目code
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
	// 项目code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 创建时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建人工号
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 项目描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 项目名字
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 项目code，废弃，请使用code
	ProjectCode *string `json:"projectCode,omitempty" xml:"projectCode,omitempty"`
	// 项目名称，废弃，请使用name
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// 状态:valid, invalid, deleted
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 用户自定义code
	UserDefineCode *string `json:"userDefineCode,omitempty" xml:"userDefineCode,omitempty"`
}

func (s GetProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBody) GoString() string {
	return s.String()
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
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 单据号
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 模型id
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
	// 数据来源于开放时，对应的微应用id
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// 单据数据体json
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// 数据模型id
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// 数据来源：审批(approval)，开放接口(openapi)
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
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetReceiptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 供应商code
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
	// 供应商Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 创建时间(单位MS)
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 供应商描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 供应商名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 状态：启用(valid), 停用(invalid), 删除(deleted)
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 用户自定义code
	UserDefineCode *string `json:"userDefineCode,omitempty" xml:"userDefineCode,omitempty"`
}

func (s GetSupplierResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSupplierResponseBody) GoString() string {
	return s.String()
}

func (s *GetSupplierResponseBody) SetCode(v string) *GetSupplierResponseBody {
	s.Code = &v
	return s
}

func (s *GetSupplierResponseBody) SetCreateTime(v int64) *GetSupplierResponseBody {
	s.CreateTime = &v
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

type GetSupplierResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSupplierResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetSupplierResponse) SetBody(v *GetSupplierResponseBody) *GetSupplierResponse {
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
	// 分页，从1开始
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页大小，默认10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 类型：income收入，expense支出
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
	// 是否还有更多数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// resultList
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
	// 类别code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否为目录
	IsDir *bool `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// 名字
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 父类别code
	ParentCode *string `json:"parentCode,omitempty" xml:"parentCode,omitempty"`
	// 状态:valid,invalid,deleted
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 类型:income收入，expense支出
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

func (s *QueryCategoryByPageResponseBodyList) SetStatus(v string) *QueryCategoryByPageResponseBodyList {
	s.Status = &v
	return s
}

func (s *QueryCategoryByPageResponseBodyList) SetType(v string) *QueryCategoryByPageResponseBodyList {
	s.Type = &v
	return s
}

type QueryCategoryByPageResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCategoryByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryCategoryByPageResponse) SetBody(v *QueryCategoryByPageResponseBody) *QueryCategoryByPageResponse {
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
	// 分页，从1开始
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页大小，默认10
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
	// 是否还有更多数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// resultList
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
	// 客户Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 创建时间(单位MS)
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 客户描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 客户名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 状态：启用(valid), 停用(invalid), 删除(deleted)
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 用户自定义code
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCustomerByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryCustomerByPageResponse) SetBody(v *QueryCustomerByPageResponseBody) *QueryCustomerByPageResponse {
	s.Body = v
	return s
}

type QueryCustomerInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCustomerInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryCustomerInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryCustomerInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCustomerInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCustomerInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCustomerInfoRequest struct {
	// 查询条件，目前支持 名字、税号、购方电话
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 查询页码，从1开始
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页查询数量
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryCustomerInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryCustomerInfoRequest) SetKeyword(v string) *QueryCustomerInfoRequest {
	s.Keyword = &v
	return s
}

func (s *QueryCustomerInfoRequest) SetPageNumber(v int64) *QueryCustomerInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryCustomerInfoRequest) SetPageSize(v int64) *QueryCustomerInfoRequest {
	s.PageSize = &v
	return s
}

type QueryCustomerInfoResponseBody struct {
	// 是否还有数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 客户分页数据
	List []*QueryCustomerInfoResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// 总数
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryCustomerInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCustomerInfoResponseBody) SetHasMore(v bool) *QueryCustomerInfoResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryCustomerInfoResponseBody) SetList(v []*QueryCustomerInfoResponseBodyList) *QueryCustomerInfoResponseBody {
	s.List = v
	return s
}

func (s *QueryCustomerInfoResponseBody) SetTotalCount(v int64) *QueryCustomerInfoResponseBody {
	s.TotalCount = &v
	return s
}

type QueryCustomerInfoResponseBodyList struct {
	// 客户code
	Code                    *string `json:"code,omitempty" xml:"code,omitempty"`
	ContactAddress          *string `json:"contactAddress,omitempty" xml:"contactAddress,omitempty"`
	ContactCompanyTelephone *string `json:"contactCompanyTelephone,omitempty" xml:"contactCompanyTelephone,omitempty"`
	ContactEmail            *string `json:"contactEmail,omitempty" xml:"contactEmail,omitempty"`
	ContactName             *string `json:"contactName,omitempty" xml:"contactName,omitempty"`
	ContactTelephone        *string `json:"contactTelephone,omitempty" xml:"contactTelephone,omitempty"`
	// 客户描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 客户名字
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 购方账户
	PurchaserAccount *string `json:"purchaserAccount,omitempty" xml:"purchaserAccount,omitempty"`
	// 购方地址
	PurchaserAddress *string `json:"purchaserAddress,omitempty" xml:"purchaserAddress,omitempty"`
	// 购方姓名
	PurchaserName *string `json:"purchaserName,omitempty" xml:"purchaserName,omitempty"`
	// 购方税号
	PurchaserTaxNo *string `json:"purchaserTaxNo,omitempty" xml:"purchaserTaxNo,omitempty"`
	// 购方电话
	PurchaserTel *string `json:"purchaserTel,omitempty" xml:"purchaserTel,omitempty"`
	// 购方银行
	PurchaserrBankName *string `json:"purchaserrBankName,omitempty" xml:"purchaserrBankName,omitempty"`
	// 状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 用户自定义code
	UserDefineCode *string `json:"userDefineCode,omitempty" xml:"userDefineCode,omitempty"`
}

func (s QueryCustomerInfoResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerInfoResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryCustomerInfoResponseBodyList) SetCode(v string) *QueryCustomerInfoResponseBodyList {
	s.Code = &v
	return s
}

func (s *QueryCustomerInfoResponseBodyList) SetContactAddress(v string) *QueryCustomerInfoResponseBodyList {
	s.ContactAddress = &v
	return s
}

func (s *QueryCustomerInfoResponseBodyList) SetContactCompanyTelephone(v string) *QueryCustomerInfoResponseBodyList {
	s.ContactCompanyTelephone = &v
	return s
}

func (s *QueryCustomerInfoResponseBodyList) SetContactEmail(v string) *QueryCustomerInfoResponseBodyList {
	s.ContactEmail = &v
	return s
}

func (s *QueryCustomerInfoResponseBodyList) SetContactName(v string) *QueryCustomerInfoResponseBodyList {
	s.ContactName = &v
	return s
}

func (s *QueryCustomerInfoResponseBodyList) SetContactTelephone(v string) *QueryCustomerInfoResponseBodyList {
	s.ContactTelephone = &v
	return s
}

func (s *QueryCustomerInfoResponseBodyList) SetDescription(v string) *QueryCustomerInfoResponseBodyList {
	s.Description = &v
	return s
}

func (s *QueryCustomerInfoResponseBodyList) SetName(v string) *QueryCustomerInfoResponseBodyList {
	s.Name = &v
	return s
}

func (s *QueryCustomerInfoResponseBodyList) SetPurchaserAccount(v string) *QueryCustomerInfoResponseBodyList {
	s.PurchaserAccount = &v
	return s
}

func (s *QueryCustomerInfoResponseBodyList) SetPurchaserAddress(v string) *QueryCustomerInfoResponseBodyList {
	s.PurchaserAddress = &v
	return s
}

func (s *QueryCustomerInfoResponseBodyList) SetPurchaserName(v string) *QueryCustomerInfoResponseBodyList {
	s.PurchaserName = &v
	return s
}

func (s *QueryCustomerInfoResponseBodyList) SetPurchaserTaxNo(v string) *QueryCustomerInfoResponseBodyList {
	s.PurchaserTaxNo = &v
	return s
}

func (s *QueryCustomerInfoResponseBodyList) SetPurchaserTel(v string) *QueryCustomerInfoResponseBodyList {
	s.PurchaserTel = &v
	return s
}

func (s *QueryCustomerInfoResponseBodyList) SetPurchaserrBankName(v string) *QueryCustomerInfoResponseBodyList {
	s.PurchaserrBankName = &v
	return s
}

func (s *QueryCustomerInfoResponseBodyList) SetStatus(v string) *QueryCustomerInfoResponseBodyList {
	s.Status = &v
	return s
}

func (s *QueryCustomerInfoResponseBodyList) SetUserDefineCode(v string) *QueryCustomerInfoResponseBodyList {
	s.UserDefineCode = &v
	return s
}

type QueryCustomerInfoResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCustomerInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCustomerInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryCustomerInfoResponse) SetHeaders(v map[string]*string) *QueryCustomerInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryCustomerInfoResponse) SetBody(v *QueryCustomerInfoResponseBody) *QueryCustomerInfoResponse {
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
	// 分页，从1开始
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页大小，默认10
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
	// 是否还有更多数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// resultList
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
	// 账户code
	AccountCode *string `json:"accountCode,omitempty" xml:"accountCode,omitempty"`
	// 关联资金账号id
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// 账户名称
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	// 备注
	AccountRemark *string `json:"accountRemark,omitempty" xml:"accountRemark,omitempty"`
	// 账户类型:ALIPAY, BANKCARD, CASH, WECHAT
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	// 账户总额，保留2位小数
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 银行代号，如果是银行卡类型，有值，其他类型时，为空
	BankCode *string `json:"bankCode,omitempty" xml:"bankCode,omitempty"`
	// 银行名字，如果是银行卡类型，有值，其他类型时，为空
	BankName *string `json:"bankName,omitempty" xml:"bankName,omitempty"`
	// 创建时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建人工号
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
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

func (s *QueryEnterpriseAccountByPageResponseBodyList) SetCreateTime(v int64) *QueryEnterpriseAccountByPageResponseBodyList {
	s.CreateTime = &v
	return s
}

func (s *QueryEnterpriseAccountByPageResponseBodyList) SetCreator(v string) *QueryEnterpriseAccountByPageResponseBodyList {
	s.Creator = &v
	return s
}

type QueryEnterpriseAccountByPageResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryEnterpriseAccountByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryEnterpriseAccountByPageResponse) SetBody(v *QueryEnterpriseAccountByPageResponseBody) *QueryEnterpriseAccountByPageResponse {
	s.Body = v
	return s
}

type QueryFinanceCompanyInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryFinanceCompanyInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryFinanceCompanyInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryFinanceCompanyInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryFinanceCompanyInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryFinanceCompanyInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryFinanceCompanyInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryFinanceCompanyInfoResponseBody struct {
	CompanyName *string `json:"companyName,omitempty" xml:"companyName,omitempty"`
	TaxNature   *string `json:"taxNature,omitempty" xml:"taxNature,omitempty"`
	TaxNo       *string `json:"taxNo,omitempty" xml:"taxNo,omitempty"`
}

func (s QueryFinanceCompanyInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFinanceCompanyInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFinanceCompanyInfoResponseBody) SetCompanyName(v string) *QueryFinanceCompanyInfoResponseBody {
	s.CompanyName = &v
	return s
}

func (s *QueryFinanceCompanyInfoResponseBody) SetTaxNature(v string) *QueryFinanceCompanyInfoResponseBody {
	s.TaxNature = &v
	return s
}

func (s *QueryFinanceCompanyInfoResponseBody) SetTaxNo(v string) *QueryFinanceCompanyInfoResponseBody {
	s.TaxNo = &v
	return s
}

type QueryFinanceCompanyInfoResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryFinanceCompanyInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryFinanceCompanyInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFinanceCompanyInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryFinanceCompanyInfoResponse) SetHeaders(v map[string]*string) *QueryFinanceCompanyInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryFinanceCompanyInfoResponse) SetBody(v *QueryFinanceCompanyInfoResponseBody) *QueryFinanceCompanyInfoResponse {
	s.Body = v
	return s
}

type QueryInvoiceRelationCountHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryInvoiceRelationCountHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoiceRelationCountHeaders) GoString() string {
	return s.String()
}

func (s *QueryInvoiceRelationCountHeaders) SetCommonHeaders(v map[string]*string) *QueryInvoiceRelationCountHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryInvoiceRelationCountHeaders) SetXAcsDingtalkAccessToken(v string) *QueryInvoiceRelationCountHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryInvoiceRelationCountResponseBody struct {
	RelationCountMap map[string]*int64 `json:"relationCountMap,omitempty" xml:"relationCountMap,omitempty"`
}

func (s QueryInvoiceRelationCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoiceRelationCountResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInvoiceRelationCountResponseBody) SetRelationCountMap(v map[string]*int64) *QueryInvoiceRelationCountResponseBody {
	s.RelationCountMap = v
	return s
}

type QueryInvoiceRelationCountResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryInvoiceRelationCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryInvoiceRelationCountResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoiceRelationCountResponse) GoString() string {
	return s.String()
}

func (s *QueryInvoiceRelationCountResponse) SetHeaders(v map[string]*string) *QueryInvoiceRelationCountResponse {
	s.Headers = v
	return s
}

func (s *QueryInvoiceRelationCountResponse) SetBody(v *QueryInvoiceRelationCountResponseBody) *QueryInvoiceRelationCountResponse {
	s.Body = v
	return s
}

type QueryPermissionByUserIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryPermissionByUserIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryPermissionByUserIdHeaders) GoString() string {
	return s.String()
}

func (s *QueryPermissionByUserIdHeaders) SetCommonHeaders(v map[string]*string) *QueryPermissionByUserIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryPermissionByUserIdHeaders) SetXAcsDingtalkAccessToken(v string) *QueryPermissionByUserIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryPermissionByUserIdRequest struct {
	// 用户ID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryPermissionByUserIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPermissionByUserIdRequest) GoString() string {
	return s.String()
}

func (s *QueryPermissionByUserIdRequest) SetUserId(v string) *QueryPermissionByUserIdRequest {
	s.UserId = &v
	return s
}

type QueryPermissionByUserIdResponseBody struct {
	// 权限信息列表
	PermissionDTOList []*QueryPermissionByUserIdResponseBodyPermissionDTOList `json:"permissionDTOList,omitempty" xml:"permissionDTOList,omitempty" type:"Repeated"`
	// 用户ID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryPermissionByUserIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPermissionByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPermissionByUserIdResponseBody) SetPermissionDTOList(v []*QueryPermissionByUserIdResponseBodyPermissionDTOList) *QueryPermissionByUserIdResponseBody {
	s.PermissionDTOList = v
	return s
}

func (s *QueryPermissionByUserIdResponseBody) SetUserId(v string) *QueryPermissionByUserIdResponseBody {
	s.UserId = &v
	return s
}

type QueryPermissionByUserIdResponseBodyPermissionDTOList struct {
	ActionIdList     []*string `json:"actionIdList,omitempty" xml:"actionIdList,omitempty" type:"Repeated"`
	ResourceIdentity *string   `json:"resourceIdentity,omitempty" xml:"resourceIdentity,omitempty"`
}

func (s QueryPermissionByUserIdResponseBodyPermissionDTOList) String() string {
	return tea.Prettify(s)
}

func (s QueryPermissionByUserIdResponseBodyPermissionDTOList) GoString() string {
	return s.String()
}

func (s *QueryPermissionByUserIdResponseBodyPermissionDTOList) SetActionIdList(v []*string) *QueryPermissionByUserIdResponseBodyPermissionDTOList {
	s.ActionIdList = v
	return s
}

func (s *QueryPermissionByUserIdResponseBodyPermissionDTOList) SetResourceIdentity(v string) *QueryPermissionByUserIdResponseBodyPermissionDTOList {
	s.ResourceIdentity = &v
	return s
}

type QueryPermissionByUserIdResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryPermissionByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPermissionByUserIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPermissionByUserIdResponse) GoString() string {
	return s.String()
}

func (s *QueryPermissionByUserIdResponse) SetHeaders(v map[string]*string) *QueryPermissionByUserIdResponse {
	s.Headers = v
	return s
}

func (s *QueryPermissionByUserIdResponse) SetBody(v *QueryPermissionByUserIdResponseBody) *QueryPermissionByUserIdResponse {
	s.Body = v
	return s
}

type QueryPermissionRoleMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryPermissionRoleMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryPermissionRoleMemberHeaders) GoString() string {
	return s.String()
}

func (s *QueryPermissionRoleMemberHeaders) SetCommonHeaders(v map[string]*string) *QueryPermissionRoleMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryPermissionRoleMemberHeaders) SetXAcsDingtalkAccessToken(v string) *QueryPermissionRoleMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryPermissionRoleMemberRequest struct {
	// 角色的唯一标识列表
	RoleCodeList []*string `json:"roleCodeList,omitempty" xml:"roleCodeList,omitempty" type:"Repeated"`
}

func (s QueryPermissionRoleMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPermissionRoleMemberRequest) GoString() string {
	return s.String()
}

func (s *QueryPermissionRoleMemberRequest) SetRoleCodeList(v []*string) *QueryPermissionRoleMemberRequest {
	s.RoleCodeList = v
	return s
}

type QueryPermissionRoleMemberResponseBody struct {
	// 角色下的成员MAP
	RoleMemberMap map[string]*RoleMemberMapValue `json:"roleMemberMap,omitempty" xml:"roleMemberMap,omitempty"`
}

func (s QueryPermissionRoleMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPermissionRoleMemberResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPermissionRoleMemberResponseBody) SetRoleMemberMap(v map[string]*RoleMemberMapValue) *QueryPermissionRoleMemberResponseBody {
	s.RoleMemberMap = v
	return s
}

type QueryPermissionRoleMemberResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryPermissionRoleMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPermissionRoleMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPermissionRoleMemberResponse) GoString() string {
	return s.String()
}

func (s *QueryPermissionRoleMemberResponse) SetHeaders(v map[string]*string) *QueryPermissionRoleMemberResponse {
	s.Headers = v
	return s
}

func (s *QueryPermissionRoleMemberResponse) SetBody(v *QueryPermissionRoleMemberResponseBody) *QueryPermissionRoleMemberResponse {
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
	// 分页，从1开始
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页大小，默认10
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
	// 是否还有更多数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// resultList
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
	// 项目code
	Caode *string `json:"caode,omitempty" xml:"caode,omitempty"`
	// 创建时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建人工号
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 项目名字
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 项目code，废弃，请使用code
	ProjectCode *string `json:"projectCode,omitempty" xml:"projectCode,omitempty"`
	// 项目名称，废弃，请使用name
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// 状态: valid, invalid, deleted
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 用户自定义code
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryProjectByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 发票状态筛选列表 applied 已生成、unapplied 未生成 、 ignore 已忽略
	ApplyStatusList []*string `json:"applyStatusList,omitempty" xml:"applyStatusList,omitempty" type:"Repeated"`
	EndTime         *int64    `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 分页参数，从1 开始
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页参数，每页查询个数
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 单据状态筛选条件列表，审批中、已通过 RUNNGIN、COMPLETED
	ReceiptStatusList []*string `json:"receiptStatusList,omitempty" xml:"receiptStatusList,omitempty" type:"Repeated"`
	// 开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 单据标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryReceiptForInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiptForInvoiceRequest) GoString() string {
	return s.String()
}

func (s *QueryReceiptForInvoiceRequest) SetApplyStatusList(v []*string) *QueryReceiptForInvoiceRequest {
	s.ApplyStatusList = v
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

func (s *QueryReceiptForInvoiceRequest) SetStartTime(v int64) *QueryReceiptForInvoiceRequest {
	s.StartTime = &v
	return s
}

func (s *QueryReceiptForInvoiceRequest) SetTitle(v string) *QueryReceiptForInvoiceRequest {
	s.Title = &v
	return s
}

type QueryReceiptForInvoiceResponseBody struct {
	// 是否还有数据
	HasMore *string `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 分页数据
	List []*QueryReceiptForInvoiceResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// 总数
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
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 开票状态
	ApplyStatus *string `json:"applyStatus,omitempty" xml:"applyStatus,omitempty"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建人
	Creator *QueryReceiptForInvoiceResponseBodyListCreator `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	// 客户
	Customer *QueryReceiptForInvoiceResponseBodyListCustomer `json:"customer,omitempty" xml:"customer,omitempty" type:"Struct"`
	// 发票种类
	InvoiceType *string `json:"invoiceType,omitempty" xml:"invoiceType,omitempty"`
	// 主数据modelId
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// 购方账户
	PurchaserAccount *string `json:"purchaserAccount,omitempty" xml:"purchaserAccount,omitempty"`
	// 购方地址
	PurchaserAddress *string `json:"purchaserAddress,omitempty" xml:"purchaserAddress,omitempty"`
	// 购方银行
	PurchaserBankName *string `json:"purchaserBankName,omitempty" xml:"purchaserBankName,omitempty"`
	// 购方抬头
	PurchaserName *string `json:"purchaserName,omitempty" xml:"purchaserName,omitempty"`
	// 购方税号
	PurchaserTaxNo *string `json:"purchaserTaxNo,omitempty" xml:"purchaserTaxNo,omitempty"`
	// 购方电话
	PurchaserTel *string `json:"purchaserTel,omitempty" xml:"purchaserTel,omitempty"`
	// 单据ID
	ReceiptId *string `json:"receiptId,omitempty" xml:"receiptId,omitempty"`
	// 记录时间，默认为审批通过时间
	RecordTime *string `json:"recordTime,omitempty" xml:"recordTime,omitempty"`
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 来源
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 状态 agree running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 单据标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryReceiptForInvoiceResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiptForInvoiceResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetAmount(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.Amount = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetApplyStatus(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.ApplyStatus = &v
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

func (s *QueryReceiptForInvoiceResponseBodyList) SetInvoiceType(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.InvoiceType = &v
	return s
}

func (s *QueryReceiptForInvoiceResponseBodyList) SetModelId(v string) *QueryReceiptForInvoiceResponseBodyList {
	s.ModelId = &v
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
	// 创建人头像
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// 创建人昵称
	Nick *string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 创建人工号
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
	// 客户code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 客户名字
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

type QueryReceiptForInvoiceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryReceiptForInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryReceiptForInvoiceResponse) SetBody(v *QueryReceiptForInvoiceResponseBody) *QueryReceiptForInvoiceResponse {
	s.Body = v
	return s
}

type QueryReceiptsBaseInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryReceiptsBaseInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiptsBaseInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryReceiptsBaseInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryReceiptsBaseInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryReceiptsBaseInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryReceiptsBaseInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryReceiptsBaseInfoRequest struct {
	// 结束时间
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 分页参数，从1 开始
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页参数，每页查询个数
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 单据标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 凭证状态
	VoucherStatus *string `json:"voucherStatus,omitempty" xml:"voucherStatus,omitempty"`
}

func (s QueryReceiptsBaseInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiptsBaseInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryReceiptsBaseInfoRequest) SetEndTime(v int64) *QueryReceiptsBaseInfoRequest {
	s.EndTime = &v
	return s
}

func (s *QueryReceiptsBaseInfoRequest) SetPageNumber(v int64) *QueryReceiptsBaseInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryReceiptsBaseInfoRequest) SetPageSize(v int64) *QueryReceiptsBaseInfoRequest {
	s.PageSize = &v
	return s
}

func (s *QueryReceiptsBaseInfoRequest) SetStartTime(v int64) *QueryReceiptsBaseInfoRequest {
	s.StartTime = &v
	return s
}

func (s *QueryReceiptsBaseInfoRequest) SetTitle(v string) *QueryReceiptsBaseInfoRequest {
	s.Title = &v
	return s
}

func (s *QueryReceiptsBaseInfoRequest) SetVoucherStatus(v string) *QueryReceiptsBaseInfoRequest {
	s.VoucherStatus = &v
	return s
}

type QueryReceiptsBaseInfoResponseBody struct {
	// 是否还有数据
	HasMore *string `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 分页数据
	List []*QueryReceiptsBaseInfoResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// 总数
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryReceiptsBaseInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiptsBaseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryReceiptsBaseInfoResponseBody) SetHasMore(v string) *QueryReceiptsBaseInfoResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryReceiptsBaseInfoResponseBody) SetList(v []*QueryReceiptsBaseInfoResponseBodyList) *QueryReceiptsBaseInfoResponseBody {
	s.List = v
	return s
}

func (s *QueryReceiptsBaseInfoResponseBody) SetTotalCount(v int64) *QueryReceiptsBaseInfoResponseBody {
	s.TotalCount = &v
	return s
}

type QueryReceiptsBaseInfoResponseBodyList struct {
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建人
	Creator *QueryReceiptsBaseInfoResponseBodyListCreator `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	// 客户
	Customer *QueryReceiptsBaseInfoResponseBodyListCustomer `json:"customer,omitempty" xml:"customer,omitempty" type:"Struct"`
	// 主数据modelId
	ModelId   *string                                         `json:"modelId,omitempty" xml:"modelId,omitempty"`
	Principal *QueryReceiptsBaseInfoResponseBodyListPrincipal `json:"principal,omitempty" xml:"principal,omitempty" type:"Struct"`
	Project   *QueryReceiptsBaseInfoResponseBodyListProject   `json:"project,omitempty" xml:"project,omitempty" type:"Struct"`
	// 单据ID
	ReceiptId *string `json:"receiptId,omitempty" xml:"receiptId,omitempty"`
	// 记录时间，默认为审批通过时间
	RecordTime *string `json:"recordTime,omitempty" xml:"recordTime,omitempty"`
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 来源
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 状态 agree running
	Status   *string                                        `json:"status,omitempty" xml:"status,omitempty"`
	Supplier *QueryReceiptsBaseInfoResponseBodyListSupplier `json:"supplier,omitempty" xml:"supplier,omitempty" type:"Struct"`
	// 单据标题
	Title         *string `json:"title,omitempty" xml:"title,omitempty"`
	VoucherStatus *string `json:"voucherStatus,omitempty" xml:"voucherStatus,omitempty"`
}

func (s QueryReceiptsBaseInfoResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiptsBaseInfoResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryReceiptsBaseInfoResponseBodyList) SetAmount(v string) *QueryReceiptsBaseInfoResponseBodyList {
	s.Amount = &v
	return s
}

func (s *QueryReceiptsBaseInfoResponseBodyList) SetCreateTime(v string) *QueryReceiptsBaseInfoResponseBodyList {
	s.CreateTime = &v
	return s
}

func (s *QueryReceiptsBaseInfoResponseBodyList) SetCreator(v *QueryReceiptsBaseInfoResponseBodyListCreator) *QueryReceiptsBaseInfoResponseBodyList {
	s.Creator = v
	return s
}

func (s *QueryReceiptsBaseInfoResponseBodyList) SetCustomer(v *QueryReceiptsBaseInfoResponseBodyListCustomer) *QueryReceiptsBaseInfoResponseBodyList {
	s.Customer = v
	return s
}

func (s *QueryReceiptsBaseInfoResponseBodyList) SetModelId(v string) *QueryReceiptsBaseInfoResponseBodyList {
	s.ModelId = &v
	return s
}

func (s *QueryReceiptsBaseInfoResponseBodyList) SetPrincipal(v *QueryReceiptsBaseInfoResponseBodyListPrincipal) *QueryReceiptsBaseInfoResponseBodyList {
	s.Principal = v
	return s
}

func (s *QueryReceiptsBaseInfoResponseBodyList) SetProject(v *QueryReceiptsBaseInfoResponseBodyListProject) *QueryReceiptsBaseInfoResponseBodyList {
	s.Project = v
	return s
}

func (s *QueryReceiptsBaseInfoResponseBodyList) SetReceiptId(v string) *QueryReceiptsBaseInfoResponseBodyList {
	s.ReceiptId = &v
	return s
}

func (s *QueryReceiptsBaseInfoResponseBodyList) SetRecordTime(v string) *QueryReceiptsBaseInfoResponseBodyList {
	s.RecordTime = &v
	return s
}

func (s *QueryReceiptsBaseInfoResponseBodyList) SetRemark(v string) *QueryReceiptsBaseInfoResponseBodyList {
	s.Remark = &v
	return s
}

func (s *QueryReceiptsBaseInfoResponseBodyList) SetSource(v string) *QueryReceiptsBaseInfoResponseBodyList {
	s.Source = &v
	return s
}

func (s *QueryReceiptsBaseInfoResponseBodyList) SetStatus(v string) *QueryReceiptsBaseInfoResponseBodyList {
	s.Status = &v
	return s
}

func (s *QueryReceiptsBaseInfoResponseBodyList) SetSupplier(v *QueryReceiptsBaseInfoResponseBodyListSupplier) *QueryReceiptsBaseInfoResponseBodyList {
	s.Supplier = v
	return s
}

func (s *QueryReceiptsBaseInfoResponseBodyList) SetTitle(v string) *QueryReceiptsBaseInfoResponseBodyList {
	s.Title = &v
	return s
}

func (s *QueryReceiptsBaseInfoResponseBodyList) SetVoucherStatus(v string) *QueryReceiptsBaseInfoResponseBodyList {
	s.VoucherStatus = &v
	return s
}

type QueryReceiptsBaseInfoResponseBodyListCreator struct {
	// 创建人头像
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// 创建人昵称
	Nick *string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 创建人工号
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryReceiptsBaseInfoResponseBodyListCreator) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiptsBaseInfoResponseBodyListCreator) GoString() string {
	return s.String()
}

func (s *QueryReceiptsBaseInfoResponseBodyListCreator) SetAvatarUrl(v string) *QueryReceiptsBaseInfoResponseBodyListCreator {
	s.AvatarUrl = &v
	return s
}

func (s *QueryReceiptsBaseInfoResponseBodyListCreator) SetNick(v string) *QueryReceiptsBaseInfoResponseBodyListCreator {
	s.Nick = &v
	return s
}

func (s *QueryReceiptsBaseInfoResponseBodyListCreator) SetUserId(v string) *QueryReceiptsBaseInfoResponseBodyListCreator {
	s.UserId = &v
	return s
}

type QueryReceiptsBaseInfoResponseBodyListCustomer struct {
	// 客户code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 客户名字
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryReceiptsBaseInfoResponseBodyListCustomer) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiptsBaseInfoResponseBodyListCustomer) GoString() string {
	return s.String()
}

func (s *QueryReceiptsBaseInfoResponseBodyListCustomer) SetCode(v string) *QueryReceiptsBaseInfoResponseBodyListCustomer {
	s.Code = &v
	return s
}

func (s *QueryReceiptsBaseInfoResponseBodyListCustomer) SetName(v string) *QueryReceiptsBaseInfoResponseBodyListCustomer {
	s.Name = &v
	return s
}

type QueryReceiptsBaseInfoResponseBodyListPrincipal struct {
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	Nick      *string `json:"nick,omitempty" xml:"nick,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryReceiptsBaseInfoResponseBodyListPrincipal) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiptsBaseInfoResponseBodyListPrincipal) GoString() string {
	return s.String()
}

func (s *QueryReceiptsBaseInfoResponseBodyListPrincipal) SetAvatarUrl(v string) *QueryReceiptsBaseInfoResponseBodyListPrincipal {
	s.AvatarUrl = &v
	return s
}

func (s *QueryReceiptsBaseInfoResponseBodyListPrincipal) SetNick(v string) *QueryReceiptsBaseInfoResponseBodyListPrincipal {
	s.Nick = &v
	return s
}

func (s *QueryReceiptsBaseInfoResponseBodyListPrincipal) SetUserId(v string) *QueryReceiptsBaseInfoResponseBodyListPrincipal {
	s.UserId = &v
	return s
}

type QueryReceiptsBaseInfoResponseBodyListProject struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryReceiptsBaseInfoResponseBodyListProject) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiptsBaseInfoResponseBodyListProject) GoString() string {
	return s.String()
}

func (s *QueryReceiptsBaseInfoResponseBodyListProject) SetCode(v string) *QueryReceiptsBaseInfoResponseBodyListProject {
	s.Code = &v
	return s
}

func (s *QueryReceiptsBaseInfoResponseBodyListProject) SetName(v string) *QueryReceiptsBaseInfoResponseBodyListProject {
	s.Name = &v
	return s
}

type QueryReceiptsBaseInfoResponseBodyListSupplier struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryReceiptsBaseInfoResponseBodyListSupplier) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiptsBaseInfoResponseBodyListSupplier) GoString() string {
	return s.String()
}

func (s *QueryReceiptsBaseInfoResponseBodyListSupplier) SetCode(v string) *QueryReceiptsBaseInfoResponseBodyListSupplier {
	s.Code = &v
	return s
}

func (s *QueryReceiptsBaseInfoResponseBodyListSupplier) SetName(v string) *QueryReceiptsBaseInfoResponseBodyListSupplier {
	s.Name = &v
	return s
}

type QueryReceiptsBaseInfoResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryReceiptsBaseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryReceiptsBaseInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiptsBaseInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryReceiptsBaseInfoResponse) SetHeaders(v map[string]*string) *QueryReceiptsBaseInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryReceiptsBaseInfoResponse) SetBody(v *QueryReceiptsBaseInfoResponseBody) *QueryReceiptsBaseInfoResponse {
	s.Body = v
	return s
}

type QueryReceiptsByPageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryReceiptsByPageHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiptsByPageHeaders) GoString() string {
	return s.String()
}

func (s *QueryReceiptsByPageHeaders) SetCommonHeaders(v map[string]*string) *QueryReceiptsByPageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryReceiptsByPageHeaders) SetXAcsDingtalkAccessToken(v string) *QueryReceiptsByPageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryReceiptsByPageRequest struct {
	// 检索结束时间，默认当前时间，离开始时间最长不超过180天
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 数据模型id
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// 分页，从1开始
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页大小，默认10，最大100
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 检索开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 检索排序时间类型：创建时间(gmt_create)，更新时间(gmt_modified)
	TimeFilterField *string `json:"timeFilterField,omitempty" xml:"timeFilterField,omitempty"`
}

func (s QueryReceiptsByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiptsByPageRequest) GoString() string {
	return s.String()
}

func (s *QueryReceiptsByPageRequest) SetEndTime(v int64) *QueryReceiptsByPageRequest {
	s.EndTime = &v
	return s
}

func (s *QueryReceiptsByPageRequest) SetModelId(v string) *QueryReceiptsByPageRequest {
	s.ModelId = &v
	return s
}

func (s *QueryReceiptsByPageRequest) SetPageNumber(v int64) *QueryReceiptsByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryReceiptsByPageRequest) SetPageSize(v int64) *QueryReceiptsByPageRequest {
	s.PageSize = &v
	return s
}

func (s *QueryReceiptsByPageRequest) SetStartTime(v int64) *QueryReceiptsByPageRequest {
	s.StartTime = &v
	return s
}

func (s *QueryReceiptsByPageRequest) SetTimeFilterField(v string) *QueryReceiptsByPageRequest {
	s.TimeFilterField = &v
	return s
}

type QueryReceiptsByPageResponseBody struct {
	// 是否还有更多数据
	HasMore *string `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 数据列表
	List []*QueryReceiptsByPageResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s QueryReceiptsByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiptsByPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryReceiptsByPageResponseBody) SetHasMore(v string) *QueryReceiptsByPageResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryReceiptsByPageResponseBody) SetList(v []*QueryReceiptsByPageResponseBodyList) *QueryReceiptsByPageResponseBody {
	s.List = v
	return s
}

type QueryReceiptsByPageResponseBodyList struct {
	// 数据来源于开放时，对应的微应用id
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// 单据数据体json
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// 模型id
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// 数据来源：审批(approval)，开放接口(openapi)
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s QueryReceiptsByPageResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiptsByPageResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryReceiptsByPageResponseBodyList) SetAppId(v string) *QueryReceiptsByPageResponseBodyList {
	s.AppId = &v
	return s
}

func (s *QueryReceiptsByPageResponseBodyList) SetData(v string) *QueryReceiptsByPageResponseBodyList {
	s.Data = &v
	return s
}

func (s *QueryReceiptsByPageResponseBodyList) SetModelId(v string) *QueryReceiptsByPageResponseBodyList {
	s.ModelId = &v
	return s
}

func (s *QueryReceiptsByPageResponseBodyList) SetSource(v string) *QueryReceiptsByPageResponseBodyList {
	s.Source = &v
	return s
}

type QueryReceiptsByPageResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryReceiptsByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryReceiptsByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiptsByPageResponse) GoString() string {
	return s.String()
}

func (s *QueryReceiptsByPageResponse) SetHeaders(v map[string]*string) *QueryReceiptsByPageResponse {
	s.Headers = v
	return s
}

func (s *QueryReceiptsByPageResponse) SetBody(v *QueryReceiptsByPageResponseBody) *QueryReceiptsByPageResponse {
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
	// 分页，从1开始
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页大小，默认10
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
	// 是否还有更多数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// resultList
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
	// 供应商Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 创建时间(单位MS)
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 供应商描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 供应商名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 状态：启用(valid), 停用(invalid), 删除(deleted)
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 用户自定义code
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

type QuerySupplierByPageResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySupplierByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QuerySupplierByPageResponse) SetBody(v *QuerySupplierByPageResponseBody) *QuerySupplierByPageResponse {
	s.Body = v
	return s
}

type UpdateApplyReceiptAndInvoiceRelatedHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateApplyReceiptAndInvoiceRelatedHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplyReceiptAndInvoiceRelatedHeaders) GoString() string {
	return s.String()
}

func (s *UpdateApplyReceiptAndInvoiceRelatedHeaders) SetCommonHeaders(v map[string]*string) *UpdateApplyReceiptAndInvoiceRelatedHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateApplyReceiptAndInvoiceRelatedHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateApplyReceiptAndInvoiceRelatedRequest struct {
	// 发票模型
	GeneralInvoiceVOList []*UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList `json:"generalInvoiceVOList,omitempty" xml:"generalInvoiceVOList,omitempty" type:"Repeated"`
	// 审批单id
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// 操作员
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
}

func (s UpdateApplyReceiptAndInvoiceRelatedRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplyReceiptAndInvoiceRelatedRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequest) SetGeneralInvoiceVOList(v []*UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) *UpdateApplyReceiptAndInvoiceRelatedRequest {
	s.GeneralInvoiceVOList = v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequest) SetInstanceId(v string) *UpdateApplyReceiptAndInvoiceRelatedRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequest) SetOperator(v string) *UpdateApplyReceiptAndInvoiceRelatedRequest {
	s.Operator = &v
	return s
}

type UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList struct {
	// 账期时间
	AccountPeriod *string `json:"accountPeriod,omitempty" xml:"accountPeriod,omitempty"`
	// 不含税金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 含税金额
	AmountWithTax *string `json:"amountWithTax,omitempty" xml:"amountWithTax,omitempty"`
	// 校验码
	CheckCode *string `json:"checkCode,omitempty" xml:"checkCode,omitempty"`
	// 查验时间
	CheckTime *string `json:"checkTime,omitempty" xml:"checkTime,omitempty"`
	// 开票日期
	DrewDate *string `json:"drewDate,omitempty" xml:"drewDate,omitempty"`
	// 电票版式文件下载地址
	ElectronicUrl *string `json:"electronicUrl,omitempty" xml:"electronicUrl,omitempty"`
	// 财务类型，INPUT-VAT(进项),OUTPUT_VAT(销项)
	FinanceType *string `json:"financeType,omitempty" xml:"financeType,omitempty"`
	// 资金类型 ，RED（红票），（BLUE）蓝票
	FundType *string `json:"fundType,omitempty" xml:"fundType,omitempty"`
	// 常规发票明细
	GeneralInvoiceDetailVOList []*UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList `json:"generalInvoiceDetailVOList,omitempty" xml:"generalInvoiceDetailVOList,omitempty" type:"Repeated"`
	// 发票代码
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	// 发票号码
	InvoiceNo *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
	// 发票状态
	InvoiceStatus *string `json:"invoiceStatus,omitempty" xml:"invoiceStatus,omitempty"`
	// 发票类型
	InvoiceType *string `json:"invoiceType,omitempty" xml:"invoiceType,omitempty"`
	// 机器码
	MachineCode *string `json:"machineCode,omitempty" xml:"machineCode,omitempty"`
	// 成品油标识
	OilFlag *string `json:"oilFlag,omitempty" xml:"oilFlag,omitempty"`
	// 收款人
	Payee *string `json:"payee,omitempty" xml:"payee,omitempty"`
	// 审批单实例
	ProcessInstCode *string `json:"processInstCode,omitempty" xml:"processInstCode,omitempty"`
	// 审批单类型
	ProcessInstType *string `json:"processInstType,omitempty" xml:"processInstType,omitempty"`
	// 购方地址
	PurchaserAddress *string `json:"purchaserAddress,omitempty" xml:"purchaserAddress,omitempty"`
	// 购方银行账户
	PurchaserBankAccount *string `json:"purchaserBankAccount,omitempty" xml:"purchaserBankAccount,omitempty"`
	// 购方银行名称
	PurchaserBankNameAccount *string `json:"purchaserBankNameAccount,omitempty" xml:"purchaserBankNameAccount,omitempty"`
	// 购方名称
	PurchaserName *string `json:"purchaserName,omitempty" xml:"purchaserName,omitempty"`
	// 购方税号
	PurchaserTaxNo *string `json:"purchaserTaxNo,omitempty" xml:"purchaserTaxNo,omitempty"`
	// 购方电话
	PurchaserTel *string `json:"purchaserTel,omitempty" xml:"purchaserTel,omitempty"`
	// 备注
	Remark                         *string                                                                                         `json:"remark,omitempty" xml:"remark,omitempty"`
	SecondHandCarInvoiceDetailList []*UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList `json:"secondHandCarInvoiceDetailList,omitempty" xml:"secondHandCarInvoiceDetailList,omitempty" type:"Repeated"`
	// 销方地址
	SellerAddress *string `json:"sellerAddress,omitempty" xml:"sellerAddress,omitempty"`
	// 销方银行账户
	SellerBankAccount *string `json:"sellerBankAccount,omitempty" xml:"sellerBankAccount,omitempty"`
	// 销方银行名称
	SellerBankNameAccount *string `json:"sellerBankNameAccount,omitempty" xml:"sellerBankNameAccount,omitempty"`
	// 销方名称
	SellerName *string `json:"sellerName,omitempty" xml:"sellerName,omitempty"`
	// 销方税号
	SellerTaxNo *string `json:"sellerTaxNo,omitempty" xml:"sellerTaxNo,omitempty"`
	// 销方电话
	SellerTel *string `json:"sellerTel,omitempty" xml:"sellerTel,omitempty"`
	// 代开发票标识 1-自开，2-代开
	SupplySign *string `json:"supplySign,omitempty" xml:"supplySign,omitempty"`
	// 税额
	TaxAmount                   *string                                                                                      `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	UsedVehicleSaleDetailVOList []*UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList `json:"usedVehicleSaleDetailVOList,omitempty" xml:"usedVehicleSaleDetailVOList,omitempty" type:"Repeated"`
	VehicleSaleDetailVOList     []*UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList     `json:"vehicleSaleDetailVOList,omitempty" xml:"vehicleSaleDetailVOList,omitempty" type:"Repeated"`
	// 发票查验状态
	VerifyStatus *string `json:"verifyStatus,omitempty" xml:"verifyStatus,omitempty"`
	// 凭证code
	VoucherCode *string `json:"voucherCode,omitempty" xml:"voucherCode,omitempty"`
	// 生成凭证状态
	VoucherStatus *string `json:"voucherStatus,omitempty" xml:"voucherStatus,omitempty"`
}

func (s UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) GoString() string {
	return s.String()
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetAccountPeriod(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.AccountPeriod = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetAmount(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.Amount = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetAmountWithTax(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.AmountWithTax = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetCheckCode(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.CheckCode = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetCheckTime(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.CheckTime = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetDrewDate(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.DrewDate = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetElectronicUrl(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.ElectronicUrl = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetFinanceType(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.FinanceType = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetFundType(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.FundType = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetGeneralInvoiceDetailVOList(v []*UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.GeneralInvoiceDetailVOList = v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetInvoiceCode(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.InvoiceCode = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetInvoiceNo(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.InvoiceNo = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetInvoiceStatus(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.InvoiceStatus = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetInvoiceType(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.InvoiceType = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetMachineCode(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.MachineCode = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetOilFlag(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.OilFlag = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetPayee(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.Payee = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetProcessInstCode(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.ProcessInstCode = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetProcessInstType(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.ProcessInstType = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetPurchaserAddress(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.PurchaserAddress = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetPurchaserBankAccount(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.PurchaserBankAccount = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetPurchaserBankNameAccount(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.PurchaserBankNameAccount = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetPurchaserName(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.PurchaserName = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetPurchaserTaxNo(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.PurchaserTaxNo = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetPurchaserTel(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.PurchaserTel = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetRemark(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.Remark = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetSecondHandCarInvoiceDetailList(v []*UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.SecondHandCarInvoiceDetailList = v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetSellerAddress(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.SellerAddress = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetSellerBankAccount(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.SellerBankAccount = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetSellerBankNameAccount(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.SellerBankNameAccount = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetSellerName(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.SellerName = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetSellerTaxNo(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.SellerTaxNo = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetSellerTel(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.SellerTel = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetSupplySign(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.SupplySign = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetTaxAmount(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.TaxAmount = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetUsedVehicleSaleDetailVOList(v []*UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.UsedVehicleSaleDetailVOList = v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetVehicleSaleDetailVOList(v []*UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.VehicleSaleDetailVOList = v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetVerifyStatus(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.VerifyStatus = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetVoucherCode(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.VoucherCode = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList) SetVoucherStatus(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOList {
	s.VoucherStatus = &v
	return s
}

type UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList struct {
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 商品名称
	GoodsName *string `json:"goodsName,omitempty" xml:"goodsName,omitempty"`
	// 数量
	Quantity *string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 税收分类编码
	RevenueCode *string `json:"revenueCode,omitempty" xml:"revenueCode,omitempty"`
	// 行号
	RowNo *string `json:"rowNo,omitempty" xml:"rowNo,omitempty"`
	// 规格型号
	Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
	// 税额
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	// 是否享受税收优惠：0-不享受，1-享受
	TaxPre *string `json:"taxPre,omitempty" xml:"taxPre,omitempty"`
	// 优惠政策类型
	TaxPreType *string `json:"taxPreType,omitempty" xml:"taxPreType,omitempty"`
	// 税率
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// 单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 单价
	UnitPrice *string `json:"unitPrice,omitempty" xml:"unitPrice,omitempty"`
}

func (s UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) GoString() string {
	return s.String()
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetAmount(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.Amount = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetGoodsName(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.GoodsName = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetQuantity(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.Quantity = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetRevenueCode(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.RevenueCode = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetRowNo(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.RowNo = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetSpecification(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.Specification = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetTaxAmount(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.TaxAmount = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetTaxPre(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.TaxPre = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetTaxPreType(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.TaxPreType = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetTaxRate(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.TaxRate = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetUnit(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.Unit = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetUnitPrice(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.UnitPrice = &v
	return s
}

type UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList struct {
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 车牌号
	CardNo *string `json:"cardNo,omitempty" xml:"cardNo,omitempty"`
	// 通行日期止
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// 商品名称
	GoodsName *string `json:"goodsName,omitempty" xml:"goodsName,omitempty"`
	// 税收分类编码
	RevenueCode *string `json:"revenueCode,omitempty" xml:"revenueCode,omitempty"`
	// 行号
	RowNo *string `json:"rowNo,omitempty" xml:"rowNo,omitempty"`
	// 通行日期起
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// 税额
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	// 税率
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// 类型
	VehicleType *string `json:"vehicleType,omitempty" xml:"vehicleType,omitempty"`
}

func (s UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) GoString() string {
	return s.String()
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetAmount(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.Amount = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetCardNo(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.CardNo = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetEndDate(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.EndDate = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetGoodsName(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.GoodsName = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetRevenueCode(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.RevenueCode = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetRowNo(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.RowNo = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetStartDate(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.StartDate = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetTaxAmount(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.TaxAmount = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetTaxRate(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.TaxRate = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetVehicleType(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.VehicleType = &v
	return s
}

type UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList struct {
	// 经营、拍卖单位
	AuctionUnit *string `json:"auctionUnit,omitempty" xml:"auctionUnit,omitempty"`
	// 经营、拍卖单位地址
	AuctionUnitAddress *string `json:"auctionUnitAddress,omitempty" xml:"auctionUnitAddress,omitempty"`
	// 经营、拍卖单位银行
	AuctionUnitBank *string `json:"auctionUnitBank,omitempty" xml:"auctionUnitBank,omitempty"`
	// 经营、拍卖单位税号
	AuctionUnitTaxNo *string `json:"auctionUnitTaxNo,omitempty" xml:"auctionUnitTaxNo,omitempty"`
	// 经营、拍卖单位电话
	AuctionUtilTel *string `json:"auctionUtilTel,omitempty" xml:"auctionUtilTel,omitempty"`
	// 厂牌型号
	CarModel *string `json:"carModel,omitempty" xml:"carModel,omitempty"`
	// 车牌照号
	CardNo *string `json:"cardNo,omitempty" xml:"cardNo,omitempty"`
	// 登记证号
	Registration *string `json:"registration,omitempty" xml:"registration,omitempty"`
	// 转入地车辆管理所名称
	TransferVehicle *string `json:"transferVehicle,omitempty" xml:"transferVehicle,omitempty"`
	// 二手车市场地址
	UsedCarAddress *string `json:"usedCarAddress,omitempty" xml:"usedCarAddress,omitempty"`
	// 二手车市场
	UsedCarMarket *string `json:"usedCarMarket,omitempty" xml:"usedCarMarket,omitempty"`
	// 二手车市场开户银行、账号
	UsedCarMarketBank *string `json:"usedCarMarketBank,omitempty" xml:"usedCarMarketBank,omitempty"`
	// 二手车市场电话
	UsedCarMarketPhone *string `json:"usedCarMarketPhone,omitempty" xml:"usedCarMarketPhone,omitempty"`
	// 二手车市场纳税人识别号
	UsedCarMarketTaxNo *string `json:"usedCarMarketTaxNo,omitempty" xml:"usedCarMarketTaxNo,omitempty"`
	// 车架号/车辆识别号
	VehicleNo *string `json:"vehicleNo,omitempty" xml:"vehicleNo,omitempty"`
	// 车辆类型
	VehicleType *string `json:"vehicleType,omitempty" xml:"vehicleType,omitempty"`
}

func (s UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) GoString() string {
	return s.String()
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetAuctionUnit(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.AuctionUnit = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetAuctionUnitAddress(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.AuctionUnitAddress = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetAuctionUnitBank(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.AuctionUnitBank = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetAuctionUnitTaxNo(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.AuctionUnitTaxNo = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetAuctionUtilTel(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.AuctionUtilTel = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetCarModel(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.CarModel = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetCardNo(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.CardNo = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetRegistration(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.Registration = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetTransferVehicle(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.TransferVehicle = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetUsedCarAddress(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.UsedCarAddress = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetUsedCarMarket(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.UsedCarMarket = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetUsedCarMarketBank(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.UsedCarMarketBank = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetUsedCarMarketPhone(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.UsedCarMarketPhone = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetUsedCarMarketTaxNo(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.UsedCarMarketTaxNo = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetVehicleNo(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.VehicleNo = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetVehicleType(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.VehicleType = &v
	return s
}

type UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList struct {
	// 品牌
	Brand *string `json:"brand,omitempty" xml:"brand,omitempty"`
	// 合格证号
	CertificateNo *string `json:"certificateNo,omitempty" xml:"certificateNo,omitempty"`
	// 发动机号
	EngineNo *string `json:"engineNo,omitempty" xml:"engineNo,omitempty"`
	// 身份证号/组织机构代码
	IdCardNo *string `json:"idCardNo,omitempty" xml:"idCardNo,omitempty"`
	// 进口证书号
	ImportCertificateNo *string `json:"importCertificateNo,omitempty" xml:"importCertificateNo,omitempty"`
	// 商检单号
	InspectionListNo *string `json:"inspectionListNo,omitempty" xml:"inspectionListNo,omitempty"`
	// 限乘人数
	MaxPassengers *string `json:"maxPassengers,omitempty" xml:"maxPassengers,omitempty"`
	// 产地
	OriginPlace *string `json:"originPlace,omitempty" xml:"originPlace,omitempty"`
	// 完税凭证号码
	PaymentVoucherNo *string `json:"paymentVoucherNo,omitempty" xml:"paymentVoucherNo,omitempty"`
	// 主管税务机关名称
	TaxAuthorityName *string `json:"taxAuthorityName,omitempty" xml:"taxAuthorityName,omitempty"`
	// 主管税务机关代码
	TaxAuthorityNo *string `json:"taxAuthorityNo,omitempty" xml:"taxAuthorityNo,omitempty"`
	// 税率
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// 吨位
	Tonnage *string `json:"tonnage,omitempty" xml:"tonnage,omitempty"`
	// 车架号码
	VehicleNo *string `json:"vehicleNo,omitempty" xml:"vehicleNo,omitempty"`
	// 车辆类型
	VehicleType *string `json:"vehicleType,omitempty" xml:"vehicleType,omitempty"`
}

func (s UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList) GoString() string {
	return s.String()
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetBrand(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.Brand = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetCertificateNo(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.CertificateNo = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetEngineNo(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.EngineNo = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetIdCardNo(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.IdCardNo = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetImportCertificateNo(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.ImportCertificateNo = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetInspectionListNo(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.InspectionListNo = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetMaxPassengers(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.MaxPassengers = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetOriginPlace(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.OriginPlace = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetPaymentVoucherNo(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.PaymentVoucherNo = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetTaxAuthorityName(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.TaxAuthorityName = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetTaxAuthorityNo(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.TaxAuthorityNo = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetTaxRate(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.TaxRate = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetTonnage(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.Tonnage = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetVehicleNo(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.VehicleNo = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetVehicleType(v string) *UpdateApplyReceiptAndInvoiceRelatedRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.VehicleType = &v
	return s
}

type UpdateApplyReceiptAndInvoiceRelatedResponseBody struct {
	// 批量更新发票返回结果
	//
	BatchUpdateInvoiceResponse *UpdateApplyReceiptAndInvoiceRelatedResponseBodyBatchUpdateInvoiceResponse `json:"batchUpdateInvoiceResponse,omitempty" xml:"batchUpdateInvoiceResponse,omitempty" type:"Struct"`
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateApplyReceiptAndInvoiceRelatedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplyReceiptAndInvoiceRelatedResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplyReceiptAndInvoiceRelatedResponseBody) SetBatchUpdateInvoiceResponse(v *UpdateApplyReceiptAndInvoiceRelatedResponseBodyBatchUpdateInvoiceResponse) *UpdateApplyReceiptAndInvoiceRelatedResponseBody {
	s.BatchUpdateInvoiceResponse = v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedResponseBody) SetSuccess(v bool) *UpdateApplyReceiptAndInvoiceRelatedResponseBody {
	s.Success = &v
	return s
}

type UpdateApplyReceiptAndInvoiceRelatedResponseBodyBatchUpdateInvoiceResponse struct {
	// 错误结果列表
	//
	InvoiceKeyVOList []*UpdateApplyReceiptAndInvoiceRelatedResponseBodyBatchUpdateInvoiceResponseInvoiceKeyVOList `json:"invoiceKeyVOList,omitempty" xml:"invoiceKeyVOList,omitempty" type:"Repeated"`
}

func (s UpdateApplyReceiptAndInvoiceRelatedResponseBodyBatchUpdateInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplyReceiptAndInvoiceRelatedResponseBodyBatchUpdateInvoiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplyReceiptAndInvoiceRelatedResponseBodyBatchUpdateInvoiceResponse) SetInvoiceKeyVOList(v []*UpdateApplyReceiptAndInvoiceRelatedResponseBodyBatchUpdateInvoiceResponseInvoiceKeyVOList) *UpdateApplyReceiptAndInvoiceRelatedResponseBodyBatchUpdateInvoiceResponse {
	s.InvoiceKeyVOList = v
	return s
}

type UpdateApplyReceiptAndInvoiceRelatedResponseBodyBatchUpdateInvoiceResponseInvoiceKeyVOList struct {
	// 发票编码
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	// 发票号码
	InvoiceNo *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
}

func (s UpdateApplyReceiptAndInvoiceRelatedResponseBodyBatchUpdateInvoiceResponseInvoiceKeyVOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplyReceiptAndInvoiceRelatedResponseBodyBatchUpdateInvoiceResponseInvoiceKeyVOList) GoString() string {
	return s.String()
}

func (s *UpdateApplyReceiptAndInvoiceRelatedResponseBodyBatchUpdateInvoiceResponseInvoiceKeyVOList) SetInvoiceCode(v string) *UpdateApplyReceiptAndInvoiceRelatedResponseBodyBatchUpdateInvoiceResponseInvoiceKeyVOList {
	s.InvoiceCode = &v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedResponseBodyBatchUpdateInvoiceResponseInvoiceKeyVOList) SetInvoiceNo(v string) *UpdateApplyReceiptAndInvoiceRelatedResponseBodyBatchUpdateInvoiceResponseInvoiceKeyVOList {
	s.InvoiceNo = &v
	return s
}

type UpdateApplyReceiptAndInvoiceRelatedResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateApplyReceiptAndInvoiceRelatedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateApplyReceiptAndInvoiceRelatedResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplyReceiptAndInvoiceRelatedResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplyReceiptAndInvoiceRelatedResponse) SetHeaders(v map[string]*string) *UpdateApplyReceiptAndInvoiceRelatedResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplyReceiptAndInvoiceRelatedResponse) SetBody(v *UpdateApplyReceiptAndInvoiceRelatedResponseBody) *UpdateApplyReceiptAndInvoiceRelatedResponse {
	s.Body = v
	return s
}

type UpdateFinanceCompanyInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateFinanceCompanyInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateFinanceCompanyInfoHeaders) GoString() string {
	return s.String()
}

func (s *UpdateFinanceCompanyInfoHeaders) SetCommonHeaders(v map[string]*string) *UpdateFinanceCompanyInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateFinanceCompanyInfoHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateFinanceCompanyInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateFinanceCompanyInfoRequest struct {
	// 公司名字
	CompanyName *string `json:"companyName,omitempty" xml:"companyName,omitempty"`
	// 纳税性质 小规模纳税人 一般纳税人
	TaxNature *string `json:"taxNature,omitempty" xml:"taxNature,omitempty"`
	// 税号
	TaxNo *string `json:"taxNo,omitempty" xml:"taxNo,omitempty"`
	// 用户ID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateFinanceCompanyInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFinanceCompanyInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateFinanceCompanyInfoRequest) SetCompanyName(v string) *UpdateFinanceCompanyInfoRequest {
	s.CompanyName = &v
	return s
}

func (s *UpdateFinanceCompanyInfoRequest) SetTaxNature(v string) *UpdateFinanceCompanyInfoRequest {
	s.TaxNature = &v
	return s
}

func (s *UpdateFinanceCompanyInfoRequest) SetTaxNo(v string) *UpdateFinanceCompanyInfoRequest {
	s.TaxNo = &v
	return s
}

func (s *UpdateFinanceCompanyInfoRequest) SetUserId(v string) *UpdateFinanceCompanyInfoRequest {
	s.UserId = &v
	return s
}

type UpdateFinanceCompanyInfoResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateFinanceCompanyInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFinanceCompanyInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFinanceCompanyInfoResponseBody) SetResult(v bool) *UpdateFinanceCompanyInfoResponseBody {
	s.Result = &v
	return s
}

type UpdateFinanceCompanyInfoResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFinanceCompanyInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFinanceCompanyInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFinanceCompanyInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateFinanceCompanyInfoResponse) SetHeaders(v map[string]*string) *UpdateFinanceCompanyInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateFinanceCompanyInfoResponse) SetBody(v *UpdateFinanceCompanyInfoResponseBody) *UpdateFinanceCompanyInfoResponse {
	s.Body = v
	return s
}

type UpdateInvoiceAbandonStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateInvoiceAbandonStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAbandonStatusHeaders) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAbandonStatusHeaders) SetCommonHeaders(v map[string]*string) *UpdateInvoiceAbandonStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateInvoiceAbandonStatusHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateInvoiceAbandonStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateInvoiceAbandonStatusRequest struct {
	// 发票全票面信息（蓝票）
	BlueGeneralInvoiceVO *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO `json:"blueGeneralInvoiceVO,omitempty" xml:"blueGeneralInvoiceVO,omitempty" type:"Struct"`
	// 发票编码（蓝票）
	BlueInvoiceCode *string `json:"blueInvoiceCode,omitempty" xml:"blueInvoiceCode,omitempty"`
	// 发票号码（蓝票）
	BlueInvoiceNo *string `json:"blueInvoiceNo,omitempty" xml:"blueInvoiceNo,omitempty"`
	// 状态-红冲、废弃
	BlueInvoiceStatus *string `json:"blueInvoiceStatus,omitempty" xml:"blueInvoiceStatus,omitempty"`
	// 操作员
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 发票全票面信息（红票）
	RedGeneralInvoiceVO *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO `json:"redGeneralInvoiceVO,omitempty" xml:"redGeneralInvoiceVO,omitempty" type:"Struct"`
	// 红字发票code
	RedInvoiceCode *string `json:"redInvoiceCode,omitempty" xml:"redInvoiceCode,omitempty"`
	// 红字发票编码
	RedInvoiceNo *string `json:"redInvoiceNo,omitempty" xml:"redInvoiceNo,omitempty"`
	// 红字发票状态
	RedInvoiceStatus *string `json:"redInvoiceStatus,omitempty" xml:"redInvoiceStatus,omitempty"`
	// 目标发票
	TargetInvoice *string `json:"targetInvoice,omitempty" xml:"targetInvoice,omitempty"`
}

func (s UpdateInvoiceAbandonStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAbandonStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAbandonStatusRequest) SetBlueGeneralInvoiceVO(v *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) *UpdateInvoiceAbandonStatusRequest {
	s.BlueGeneralInvoiceVO = v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequest) SetBlueInvoiceCode(v string) *UpdateInvoiceAbandonStatusRequest {
	s.BlueInvoiceCode = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequest) SetBlueInvoiceNo(v string) *UpdateInvoiceAbandonStatusRequest {
	s.BlueInvoiceNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequest) SetBlueInvoiceStatus(v string) *UpdateInvoiceAbandonStatusRequest {
	s.BlueInvoiceStatus = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequest) SetOperator(v string) *UpdateInvoiceAbandonStatusRequest {
	s.Operator = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequest) SetRedGeneralInvoiceVO(v *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) *UpdateInvoiceAbandonStatusRequest {
	s.RedGeneralInvoiceVO = v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequest) SetRedInvoiceCode(v string) *UpdateInvoiceAbandonStatusRequest {
	s.RedInvoiceCode = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequest) SetRedInvoiceNo(v string) *UpdateInvoiceAbandonStatusRequest {
	s.RedInvoiceNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequest) SetRedInvoiceStatus(v string) *UpdateInvoiceAbandonStatusRequest {
	s.RedInvoiceStatus = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequest) SetTargetInvoice(v string) *UpdateInvoiceAbandonStatusRequest {
	s.TargetInvoice = &v
	return s
}

type UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO struct {
	// 账期时间
	AccountPeriod *string `json:"accountPeriod,omitempty" xml:"accountPeriod,omitempty"`
	// 不含税金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 含税金额
	AmountWithTax *string `json:"amountWithTax,omitempty" xml:"amountWithTax,omitempty"`
	// 校验码
	CheckCode *string `json:"checkCode,omitempty" xml:"checkCode,omitempty"`
	// 查验时间
	CheckTime *string `json:"checkTime,omitempty" xml:"checkTime,omitempty"`
	// 开票日期
	DrewDate *string `json:"drewDate,omitempty" xml:"drewDate,omitempty"`
	// 电票版式文件下载地址
	ElectronicUrl *string `json:"electronicUrl,omitempty" xml:"electronicUrl,omitempty"`
	// 财务类型，INPUT-VAT(进项),OUTPUT_VAT(销项)
	FinanceType *string `json:"financeType,omitempty" xml:"financeType,omitempty"`
	// 资金类型 ，RED（红票），（BLUE）蓝票
	FundType *string `json:"fundType,omitempty" xml:"fundType,omitempty"`
	// 常规发票明细
	GeneralInvoiceDetailVOList []*UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList `json:"generalInvoiceDetailVOList,omitempty" xml:"generalInvoiceDetailVOList,omitempty" type:"Repeated"`
	// 发票代码
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	// 发票号码
	InvoiceNo *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
	// 发票状态
	InvoiceStatus *string `json:"invoiceStatus,omitempty" xml:"invoiceStatus,omitempty"`
	// 发票类型
	InvoiceType *string `json:"invoiceType,omitempty" xml:"invoiceType,omitempty"`
	// 机器码
	MachineCode *string `json:"machineCode,omitempty" xml:"machineCode,omitempty"`
	// 成品油标识
	OilFlag *string `json:"oilFlag,omitempty" xml:"oilFlag,omitempty"`
	// 收款人
	Payee *string `json:"payee,omitempty" xml:"payee,omitempty"`
	// 审批单实例
	ProcessInstCode *string `json:"processInstCode,omitempty" xml:"processInstCode,omitempty"`
	// 审批单类型
	ProcessInstType *string `json:"processInstType,omitempty" xml:"processInstType,omitempty"`
	// 购方地址
	PurchaserAddress *string `json:"purchaserAddress,omitempty" xml:"purchaserAddress,omitempty"`
	// 购方银行账户
	PurchaserBankAccount *string `json:"purchaserBankAccount,omitempty" xml:"purchaserBankAccount,omitempty"`
	// 购方银行名称
	PurchaserBankNameAccount *string `json:"purchaserBankNameAccount,omitempty" xml:"purchaserBankNameAccount,omitempty"`
	// 购方名称
	PurchaserName *string `json:"purchaserName,omitempty" xml:"purchaserName,omitempty"`
	// 购方税号
	PurchaserTaxNo *string `json:"purchaserTaxNo,omitempty" xml:"purchaserTaxNo,omitempty"`
	// 购方电话
	PurchaserTel *string `json:"purchaserTel,omitempty" xml:"purchaserTel,omitempty"`
	// 备注
	Remark                         *string                                                                                `json:"remark,omitempty" xml:"remark,omitempty"`
	SecondHandCarInvoiceDetailList []*UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOSecondHandCarInvoiceDetailList `json:"secondHandCarInvoiceDetailList,omitempty" xml:"secondHandCarInvoiceDetailList,omitempty" type:"Repeated"`
	// 销方地址
	SellerAddress *string `json:"sellerAddress,omitempty" xml:"sellerAddress,omitempty"`
	// 销方银行账户
	SellerBankAccount *string `json:"sellerBankAccount,omitempty" xml:"sellerBankAccount,omitempty"`
	// 销方银行
	SellerBankNameAccount *string `json:"sellerBankNameAccount,omitempty" xml:"sellerBankNameAccount,omitempty"`
	// 销方名称
	SellerName *string `json:"sellerName,omitempty" xml:"sellerName,omitempty"`
	// 销方税号
	SellerTaxNo *string `json:"sellerTaxNo,omitempty" xml:"sellerTaxNo,omitempty"`
	// 销方电话
	SellerTel *string `json:"sellerTel,omitempty" xml:"sellerTel,omitempty"`
	// 代开发票标识 1-自开，2-代开
	SupplySign *string `json:"supplySign,omitempty" xml:"supplySign,omitempty"`
	// 税额
	TaxAmount                   *string                                                                             `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	UsedVehicleSaleDetailVOList []*UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList `json:"usedVehicleSaleDetailVOList,omitempty" xml:"usedVehicleSaleDetailVOList,omitempty" type:"Repeated"`
	VehicleSaleDetailVOList     []*UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList     `json:"vehicleSaleDetailVOList,omitempty" xml:"vehicleSaleDetailVOList,omitempty" type:"Repeated"`
	// 发票查验状态
	VerifyStatus *string `json:"verifyStatus,omitempty" xml:"verifyStatus,omitempty"`
	// 凭证code
	VoucherCode *string `json:"voucherCode,omitempty" xml:"voucherCode,omitempty"`
	// 生成凭证状态
	VoucherStatus *string `json:"voucherStatus,omitempty" xml:"voucherStatus,omitempty"`
}

func (s UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetAccountPeriod(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.AccountPeriod = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetAmount(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.Amount = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetAmountWithTax(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.AmountWithTax = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetCheckCode(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.CheckCode = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetCheckTime(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.CheckTime = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetDrewDate(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.DrewDate = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetElectronicUrl(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.ElectronicUrl = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetFinanceType(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.FinanceType = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetFundType(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.FundType = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetGeneralInvoiceDetailVOList(v []*UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.GeneralInvoiceDetailVOList = v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetInvoiceCode(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.InvoiceCode = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetInvoiceNo(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.InvoiceNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetInvoiceStatus(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.InvoiceStatus = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetInvoiceType(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.InvoiceType = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetMachineCode(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.MachineCode = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetOilFlag(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.OilFlag = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetPayee(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.Payee = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetProcessInstCode(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.ProcessInstCode = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetProcessInstType(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.ProcessInstType = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetPurchaserAddress(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.PurchaserAddress = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetPurchaserBankAccount(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.PurchaserBankAccount = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetPurchaserBankNameAccount(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.PurchaserBankNameAccount = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetPurchaserName(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.PurchaserName = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetPurchaserTaxNo(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.PurchaserTaxNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetPurchaserTel(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.PurchaserTel = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetRemark(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.Remark = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetSecondHandCarInvoiceDetailList(v []*UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOSecondHandCarInvoiceDetailList) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.SecondHandCarInvoiceDetailList = v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetSellerAddress(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.SellerAddress = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetSellerBankAccount(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.SellerBankAccount = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetSellerBankNameAccount(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.SellerBankNameAccount = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetSellerName(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.SellerName = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetSellerTaxNo(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.SellerTaxNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetSellerTel(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.SellerTel = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetSupplySign(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.SupplySign = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetTaxAmount(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.TaxAmount = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetUsedVehicleSaleDetailVOList(v []*UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.UsedVehicleSaleDetailVOList = v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetVehicleSaleDetailVOList(v []*UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.VehicleSaleDetailVOList = v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetVerifyStatus(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.VerifyStatus = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetVoucherCode(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.VoucherCode = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO) SetVoucherStatus(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVO {
	s.VoucherStatus = &v
	return s
}

type UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList struct {
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 商品名称
	GoodsName *string `json:"goodsName,omitempty" xml:"goodsName,omitempty"`
	// 数量
	Quantity *string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 税收分类编码
	RevenueCode *string `json:"revenueCode,omitempty" xml:"revenueCode,omitempty"`
	// 行号
	RowNo *string `json:"rowNo,omitempty" xml:"rowNo,omitempty"`
	// 规格型号
	Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
	// 税额
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	// 是否享受税收优惠：0-不享受，1-享受
	TaxPre *string `json:"taxPre,omitempty" xml:"taxPre,omitempty"`
	// 优惠政策类型
	TaxPreType *string `json:"taxPreType,omitempty" xml:"taxPreType,omitempty"`
	// 税率
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// 单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 单价
	UnitPrice *string `json:"unitPrice,omitempty" xml:"unitPrice,omitempty"`
}

func (s UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList) SetAmount(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.Amount = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList) SetGoodsName(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.GoodsName = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList) SetQuantity(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.Quantity = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList) SetRevenueCode(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.RevenueCode = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList) SetRowNo(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.RowNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList) SetSpecification(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.Specification = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList) SetTaxAmount(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.TaxAmount = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList) SetTaxPre(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.TaxPre = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList) SetTaxPreType(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.TaxPreType = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList) SetTaxRate(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.TaxRate = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList) SetUnit(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.Unit = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList) SetUnitPrice(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.UnitPrice = &v
	return s
}

type UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOSecondHandCarInvoiceDetailList struct {
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 车牌号
	CardNo *string `json:"cardNo,omitempty" xml:"cardNo,omitempty"`
	// 通行日期止
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// 商品名称
	GoodsName *string `json:"goodsName,omitempty" xml:"goodsName,omitempty"`
	// 税收分类编码
	RevenueCode *string `json:"revenueCode,omitempty" xml:"revenueCode,omitempty"`
	// 行号
	RowNo *string `json:"rowNo,omitempty" xml:"rowNo,omitempty"`
	// 通行日期起
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// 税额
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	// 税率
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// 类型
	VehicleType *string `json:"vehicleType,omitempty" xml:"vehicleType,omitempty"`
}

func (s UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOSecondHandCarInvoiceDetailList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOSecondHandCarInvoiceDetailList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetAmount(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.Amount = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetCardNo(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.CardNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetEndDate(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.EndDate = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetGoodsName(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.GoodsName = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetRevenueCode(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.RevenueCode = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetRowNo(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.RowNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetStartDate(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.StartDate = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetTaxAmount(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.TaxAmount = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetTaxRate(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.TaxRate = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetVehicleType(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.VehicleType = &v
	return s
}

type UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList struct {
	// 经营、拍卖单位
	AuctionUnit *string `json:"auctionUnit,omitempty" xml:"auctionUnit,omitempty"`
	// 经营、拍卖单位地址
	AuctionUnitAddress *string `json:"auctionUnitAddress,omitempty" xml:"auctionUnitAddress,omitempty"`
	// 经营、拍卖单位银行
	AuctionUnitBank *string `json:"auctionUnitBank,omitempty" xml:"auctionUnitBank,omitempty"`
	// 经营、拍卖单位税号
	AuctionUnitTaxNo *string `json:"auctionUnitTaxNo,omitempty" xml:"auctionUnitTaxNo,omitempty"`
	// 经营、拍卖单位电话
	AuctionUtilTel *string `json:"auctionUtilTel,omitempty" xml:"auctionUtilTel,omitempty"`
	// 厂牌型号
	CarModel *string `json:"carModel,omitempty" xml:"carModel,omitempty"`
	// 车牌照号
	CardNo *string `json:"cardNo,omitempty" xml:"cardNo,omitempty"`
	// 登记证号
	Registration *string `json:"registration,omitempty" xml:"registration,omitempty"`
	// 转入地车辆管理所名称
	TransferVehicle *string `json:"transferVehicle,omitempty" xml:"transferVehicle,omitempty"`
	// 二手车市场地址
	UsedCarAddress *string `json:"usedCarAddress,omitempty" xml:"usedCarAddress,omitempty"`
	// 二手车市场
	UsedCarMarket *string `json:"usedCarMarket,omitempty" xml:"usedCarMarket,omitempty"`
	// 二手车市场开户银行、账号
	UsedCarMarketBank *string `json:"usedCarMarketBank,omitempty" xml:"usedCarMarketBank,omitempty"`
	// 二手车市场电话
	UsedCarMarketPhone *string `json:"usedCarMarketPhone,omitempty" xml:"usedCarMarketPhone,omitempty"`
	// 二手车市场纳税人识别号
	UsedCarMarketTaxNo *string `json:"usedCarMarketTaxNo,omitempty" xml:"usedCarMarketTaxNo,omitempty"`
	// 车架号/车辆识别号
	VehicleNo *string `json:"vehicleNo,omitempty" xml:"vehicleNo,omitempty"`
	// 车辆类型
	VehicleType *string `json:"vehicleType,omitempty" xml:"vehicleType,omitempty"`
}

func (s UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetAuctionUnit(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.AuctionUnit = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetAuctionUnitAddress(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.AuctionUnitAddress = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetAuctionUnitBank(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.AuctionUnitBank = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetAuctionUnitTaxNo(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.AuctionUnitTaxNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetAuctionUtilTel(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.AuctionUtilTel = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetCarModel(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.CarModel = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetCardNo(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.CardNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetRegistration(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.Registration = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetTransferVehicle(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.TransferVehicle = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetUsedCarAddress(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.UsedCarAddress = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetUsedCarMarket(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.UsedCarMarket = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetUsedCarMarketBank(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.UsedCarMarketBank = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetUsedCarMarketPhone(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.UsedCarMarketPhone = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetUsedCarMarketTaxNo(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.UsedCarMarketTaxNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetVehicleNo(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.VehicleNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetVehicleType(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.VehicleType = &v
	return s
}

type UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList struct {
	// 品牌
	Brand *string `json:"brand,omitempty" xml:"brand,omitempty"`
	// 合格证号
	CertificateNo *string `json:"certificateNo,omitempty" xml:"certificateNo,omitempty"`
	// 发动机号
	EngineNo *string `json:"engineNo,omitempty" xml:"engineNo,omitempty"`
	// 身份证号/组织机构代码
	IdCardNo *string `json:"idCardNo,omitempty" xml:"idCardNo,omitempty"`
	// 进口证书号
	ImportCertificateNo *string `json:"importCertificateNo,omitempty" xml:"importCertificateNo,omitempty"`
	InspectionListNo    *string `json:"inspectionListNo,omitempty" xml:"inspectionListNo,omitempty"`
	// 限乘人数
	MaxPassengers *string `json:"maxPassengers,omitempty" xml:"maxPassengers,omitempty"`
	// 产地
	OriginPlace *string `json:"originPlace,omitempty" xml:"originPlace,omitempty"`
	// 完税凭证号码
	PaymentVoucherNo *string `json:"paymentVoucherNo,omitempty" xml:"paymentVoucherNo,omitempty"`
	// 主管税务机关名称
	TaxAuthorityName *string `json:"taxAuthorityName,omitempty" xml:"taxAuthorityName,omitempty"`
	// 主管税务机关代码
	TaxAuthorityNo *string `json:"taxAuthorityNo,omitempty" xml:"taxAuthorityNo,omitempty"`
	// 税率
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// 吨位
	Tonnage *string `json:"tonnage,omitempty" xml:"tonnage,omitempty"`
	// 车架号码
	VehicleNo *string `json:"vehicleNo,omitempty" xml:"vehicleNo,omitempty"`
	// 车辆类型
	VehicleType *string `json:"vehicleType,omitempty" xml:"vehicleType,omitempty"`
}

func (s UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList) SetBrand(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList {
	s.Brand = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList) SetCertificateNo(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList {
	s.CertificateNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList) SetEngineNo(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList {
	s.EngineNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList) SetIdCardNo(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList {
	s.IdCardNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList) SetImportCertificateNo(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList {
	s.ImportCertificateNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList) SetInspectionListNo(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList {
	s.InspectionListNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList) SetMaxPassengers(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList {
	s.MaxPassengers = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList) SetOriginPlace(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList {
	s.OriginPlace = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList) SetPaymentVoucherNo(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList {
	s.PaymentVoucherNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList) SetTaxAuthorityName(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList {
	s.TaxAuthorityName = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList) SetTaxAuthorityNo(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList {
	s.TaxAuthorityNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList) SetTaxRate(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList {
	s.TaxRate = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList) SetTonnage(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList {
	s.Tonnage = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList) SetVehicleNo(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList {
	s.VehicleNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList) SetVehicleType(v string) *UpdateInvoiceAbandonStatusRequestBlueGeneralInvoiceVOVehicleSaleDetailVOList {
	s.VehicleType = &v
	return s
}

type UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO struct {
	// 账期时间
	AccountPeriod *string `json:"accountPeriod,omitempty" xml:"accountPeriod,omitempty"`
	// 不含税金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 含税金额
	AmountWithTax *string `json:"amountWithTax,omitempty" xml:"amountWithTax,omitempty"`
	// 校验码
	CheckCode *string `json:"checkCode,omitempty" xml:"checkCode,omitempty"`
	// 查验时间
	CheckTime *string `json:"checkTime,omitempty" xml:"checkTime,omitempty"`
	// 开票日期
	DrewDate *string `json:"drewDate,omitempty" xml:"drewDate,omitempty"`
	// 电票版式文件下载地址
	ElectronicUrl *string `json:"electronicUrl,omitempty" xml:"electronicUrl,omitempty"`
	// 财务类型，INPUT-VAT(进项),OUTPUT_VAT(销项)
	FinanceType *string `json:"financeType,omitempty" xml:"financeType,omitempty"`
	// 资金类型 ，RED（红票），（BLUE）蓝票
	FundType *string `json:"fundType,omitempty" xml:"fundType,omitempty"`
	// 常规发票明细
	GeneralInvoiceDetailVOList []*UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList `json:"generalInvoiceDetailVOList,omitempty" xml:"generalInvoiceDetailVOList,omitempty" type:"Repeated"`
	// 发票代码
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	// 发票号码
	InvoiceNo *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
	// 发票类型
	InvoiceType *string `json:"invoiceType,omitempty" xml:"invoiceType,omitempty"`
	// 机器码
	MachineCode *string `json:"machineCode,omitempty" xml:"machineCode,omitempty"`
	// 成品油标识
	OilFlag *string `json:"oilFlag,omitempty" xml:"oilFlag,omitempty"`
	// 收款人
	Payee *string `json:"payee,omitempty" xml:"payee,omitempty"`
	// 审批单实例
	ProcessInstCode *string `json:"processInstCode,omitempty" xml:"processInstCode,omitempty"`
	// 审批单类型
	ProcessInstType *string `json:"processInstType,omitempty" xml:"processInstType,omitempty"`
	// 购方地址
	PurchaserAddress *string `json:"purchaserAddress,omitempty" xml:"purchaserAddress,omitempty"`
	// 购方银行账户
	PurchaserBankAccount *string `json:"purchaserBankAccount,omitempty" xml:"purchaserBankAccount,omitempty"`
	// 购方银行
	PurchaserBankNameAccount *string `json:"purchaserBankNameAccount,omitempty" xml:"purchaserBankNameAccount,omitempty"`
	// 购方名称
	PurchaserName *string `json:"purchaserName,omitempty" xml:"purchaserName,omitempty"`
	// 购方税号
	PurchaserTaxNo *string `json:"purchaserTaxNo,omitempty" xml:"purchaserTaxNo,omitempty"`
	// 购方电话
	PurchaserTel *string `json:"purchaserTel,omitempty" xml:"purchaserTel,omitempty"`
	// 备注
	Remark                         *string                                                                               `json:"remark,omitempty" xml:"remark,omitempty"`
	SecondHandCarInvoiceDetailList []*UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOSecondHandCarInvoiceDetailList `json:"secondHandCarInvoiceDetailList,omitempty" xml:"secondHandCarInvoiceDetailList,omitempty" type:"Repeated"`
	// 销方地址
	SellerAddress *string `json:"sellerAddress,omitempty" xml:"sellerAddress,omitempty"`
	// 购方银行账户
	SellerBankAccount *string `json:"sellerBankAccount,omitempty" xml:"sellerBankAccount,omitempty"`
	// 销方银行
	SellerBankNameAccount *string `json:"sellerBankNameAccount,omitempty" xml:"sellerBankNameAccount,omitempty"`
	// 销方名称
	SellerName *string `json:"sellerName,omitempty" xml:"sellerName,omitempty"`
	// 销方税号
	SellerTaxNo *string `json:"sellerTaxNo,omitempty" xml:"sellerTaxNo,omitempty"`
	// 销方电话
	SellerTel *string `json:"sellerTel,omitempty" xml:"sellerTel,omitempty"`
	// 发票状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 代开发票标识 1-自开，2-代开
	SupplySign *string `json:"supplySign,omitempty" xml:"supplySign,omitempty"`
	// 税额
	TaxAmount                   *string                                                                            `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	UsedVehicleSaleDetailVOList []*UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList `json:"usedVehicleSaleDetailVOList,omitempty" xml:"usedVehicleSaleDetailVOList,omitempty" type:"Repeated"`
	VehicleSaleDetailVOList     []*UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList     `json:"vehicleSaleDetailVOList,omitempty" xml:"vehicleSaleDetailVOList,omitempty" type:"Repeated"`
	// 发票查验状态
	VerifyStatus *string `json:"verifyStatus,omitempty" xml:"verifyStatus,omitempty"`
	// 凭证code
	VoucherCode *string `json:"voucherCode,omitempty" xml:"voucherCode,omitempty"`
	// 生成凭证状态
	VoucherStatus *string `json:"voucherStatus,omitempty" xml:"voucherStatus,omitempty"`
}

func (s UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetAccountPeriod(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.AccountPeriod = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetAmount(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.Amount = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetAmountWithTax(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.AmountWithTax = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetCheckCode(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.CheckCode = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetCheckTime(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.CheckTime = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetDrewDate(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.DrewDate = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetElectronicUrl(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.ElectronicUrl = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetFinanceType(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.FinanceType = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetFundType(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.FundType = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetGeneralInvoiceDetailVOList(v []*UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.GeneralInvoiceDetailVOList = v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetInvoiceCode(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.InvoiceCode = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetInvoiceNo(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.InvoiceNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetInvoiceType(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.InvoiceType = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetMachineCode(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.MachineCode = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetOilFlag(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.OilFlag = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetPayee(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.Payee = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetProcessInstCode(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.ProcessInstCode = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetProcessInstType(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.ProcessInstType = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetPurchaserAddress(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.PurchaserAddress = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetPurchaserBankAccount(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.PurchaserBankAccount = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetPurchaserBankNameAccount(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.PurchaserBankNameAccount = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetPurchaserName(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.PurchaserName = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetPurchaserTaxNo(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.PurchaserTaxNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetPurchaserTel(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.PurchaserTel = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetRemark(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.Remark = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetSecondHandCarInvoiceDetailList(v []*UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOSecondHandCarInvoiceDetailList) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.SecondHandCarInvoiceDetailList = v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetSellerAddress(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.SellerAddress = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetSellerBankAccount(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.SellerBankAccount = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetSellerBankNameAccount(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.SellerBankNameAccount = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetSellerName(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.SellerName = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetSellerTaxNo(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.SellerTaxNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetSellerTel(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.SellerTel = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetStatus(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.Status = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetSupplySign(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.SupplySign = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetTaxAmount(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.TaxAmount = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetUsedVehicleSaleDetailVOList(v []*UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.UsedVehicleSaleDetailVOList = v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetVehicleSaleDetailVOList(v []*UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.VehicleSaleDetailVOList = v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetVerifyStatus(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.VerifyStatus = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetVoucherCode(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.VoucherCode = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO) SetVoucherStatus(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVO {
	s.VoucherStatus = &v
	return s
}

type UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList struct {
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 商品名称
	GoodsName *string `json:"goodsName,omitempty" xml:"goodsName,omitempty"`
	// 数量
	Quantity *string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 税收分类编码
	RevenueCode *string `json:"revenueCode,omitempty" xml:"revenueCode,omitempty"`
	// 行号
	RowNo *string `json:"rowNo,omitempty" xml:"rowNo,omitempty"`
	// 规格型号
	Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
	// 税额
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	// 是否享受税收优惠：0-不享受，1-享受
	TaxPre *string `json:"taxPre,omitempty" xml:"taxPre,omitempty"`
	// 优惠政策类型
	TaxPreType *string `json:"taxPreType,omitempty" xml:"taxPreType,omitempty"`
	// 税率
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// 单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 单价
	UnitPrice *string `json:"unitPrice,omitempty" xml:"unitPrice,omitempty"`
}

func (s UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList) SetAmount(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.Amount = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList) SetGoodsName(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.GoodsName = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList) SetQuantity(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.Quantity = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList) SetRevenueCode(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.RevenueCode = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList) SetRowNo(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.RowNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList) SetSpecification(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.Specification = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList) SetTaxAmount(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.TaxAmount = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList) SetTaxPre(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.TaxPre = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList) SetTaxPreType(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.TaxPreType = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList) SetTaxRate(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.TaxRate = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList) SetUnit(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.Unit = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList) SetUnitPrice(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.UnitPrice = &v
	return s
}

type UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOSecondHandCarInvoiceDetailList struct {
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 车牌号
	CardNo *string `json:"cardNo,omitempty" xml:"cardNo,omitempty"`
	// 通行日期止
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// 商品名称
	GoodsName *string `json:"goodsName,omitempty" xml:"goodsName,omitempty"`
	// 税收分类编码
	RevenueCode *string `json:"revenueCode,omitempty" xml:"revenueCode,omitempty"`
	// 行号
	RowNo *string `json:"rowNo,omitempty" xml:"rowNo,omitempty"`
	// 通行日期起
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// 税额
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	// 税率
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// 类型
	VehicleType *string `json:"vehicleType,omitempty" xml:"vehicleType,omitempty"`
}

func (s UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOSecondHandCarInvoiceDetailList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOSecondHandCarInvoiceDetailList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetAmount(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.Amount = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetCardNo(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.CardNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetEndDate(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.EndDate = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetGoodsName(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.GoodsName = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetRevenueCode(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.RevenueCode = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetRowNo(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.RowNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetStartDate(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.StartDate = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetTaxAmount(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.TaxAmount = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetTaxRate(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.TaxRate = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetVehicleType(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.VehicleType = &v
	return s
}

type UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList struct {
	// 经营、拍卖单位
	AuctionUnit *string `json:"auctionUnit,omitempty" xml:"auctionUnit,omitempty"`
	// 经营、拍卖单位地址
	AuctionUnitAddress *string `json:"auctionUnitAddress,omitempty" xml:"auctionUnitAddress,omitempty"`
	// 经营、拍卖单位银行
	AuctionUnitBank *string `json:"auctionUnitBank,omitempty" xml:"auctionUnitBank,omitempty"`
	// 经营、拍卖单位税号
	AuctionUnitTaxNo *string `json:"auctionUnitTaxNo,omitempty" xml:"auctionUnitTaxNo,omitempty"`
	// 经营、拍卖单位电话
	AuctionUtilTel *string `json:"auctionUtilTel,omitempty" xml:"auctionUtilTel,omitempty"`
	// 厂牌型号
	CarModel *string `json:"carModel,omitempty" xml:"carModel,omitempty"`
	// 车牌照号
	CardNo *string `json:"cardNo,omitempty" xml:"cardNo,omitempty"`
	// 登记证号
	Registration *string `json:"registration,omitempty" xml:"registration,omitempty"`
	// 转入地车辆管理所名称
	TransferVehicle *string `json:"transferVehicle,omitempty" xml:"transferVehicle,omitempty"`
	// 二手车市场地址
	UsedCarAddress *string `json:"usedCarAddress,omitempty" xml:"usedCarAddress,omitempty"`
	// 二手车市场
	UsedCarMarket *string `json:"usedCarMarket,omitempty" xml:"usedCarMarket,omitempty"`
	// 二手车市场开户银行、账号
	UsedCarMarketBank *string `json:"usedCarMarketBank,omitempty" xml:"usedCarMarketBank,omitempty"`
	// 二手车市场电话
	UsedCarMarketPhone *string `json:"usedCarMarketPhone,omitempty" xml:"usedCarMarketPhone,omitempty"`
	// 二手车市场纳税人识别号
	UsedCarMarketTaxNo *string `json:"usedCarMarketTaxNo,omitempty" xml:"usedCarMarketTaxNo,omitempty"`
	// 车架号/车辆识别号
	VehicleNo *string `json:"vehicleNo,omitempty" xml:"vehicleNo,omitempty"`
	// 车辆类型
	VehicleType *string `json:"vehicleType,omitempty" xml:"vehicleType,omitempty"`
}

func (s UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetAuctionUnit(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.AuctionUnit = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetAuctionUnitAddress(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.AuctionUnitAddress = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetAuctionUnitBank(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.AuctionUnitBank = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetAuctionUnitTaxNo(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.AuctionUnitTaxNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetAuctionUtilTel(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.AuctionUtilTel = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetCarModel(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.CarModel = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetCardNo(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.CardNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetRegistration(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.Registration = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetTransferVehicle(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.TransferVehicle = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetUsedCarAddress(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.UsedCarAddress = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetUsedCarMarket(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.UsedCarMarket = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetUsedCarMarketBank(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.UsedCarMarketBank = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetUsedCarMarketPhone(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.UsedCarMarketPhone = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetUsedCarMarketTaxNo(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.UsedCarMarketTaxNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetVehicleNo(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.VehicleNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetVehicleType(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.VehicleType = &v
	return s
}

type UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList struct {
	// 品牌
	Brand *string `json:"brand,omitempty" xml:"brand,omitempty"`
	// 合格证号
	CertificateNo *string `json:"certificateNo,omitempty" xml:"certificateNo,omitempty"`
	// 发动机号
	EngineNo *string `json:"engineNo,omitempty" xml:"engineNo,omitempty"`
	// 身份证号/组织机构代码
	IdCardNo *string `json:"idCardNo,omitempty" xml:"idCardNo,omitempty"`
	// 进口证书号
	ImportCertificateNo *string `json:"importCertificateNo,omitempty" xml:"importCertificateNo,omitempty"`
	// 商检单号
	InspectionListNo *string `json:"inspectionListNo,omitempty" xml:"inspectionListNo,omitempty"`
	// 限乘人数
	MaxPassengers *string `json:"maxPassengers,omitempty" xml:"maxPassengers,omitempty"`
	// 产地
	OriginPlace *string `json:"originPlace,omitempty" xml:"originPlace,omitempty"`
	// 完税凭证号码
	PaymentVoucherNo *string `json:"paymentVoucherNo,omitempty" xml:"paymentVoucherNo,omitempty"`
	// 主管税务机关名称
	TaxAuthorityName *string `json:"taxAuthorityName,omitempty" xml:"taxAuthorityName,omitempty"`
	// 主管税务机关代码
	TaxAuthorityNo *string `json:"taxAuthorityNo,omitempty" xml:"taxAuthorityNo,omitempty"`
	// 税率
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// 吨位
	Tonnage *string `json:"tonnage,omitempty" xml:"tonnage,omitempty"`
	// 车架号码
	VehicleNo *string `json:"vehicleNo,omitempty" xml:"vehicleNo,omitempty"`
	// 车辆类型
	VehicleType *string `json:"vehicleType,omitempty" xml:"vehicleType,omitempty"`
}

func (s UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList) SetBrand(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList {
	s.Brand = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList) SetCertificateNo(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList {
	s.CertificateNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList) SetEngineNo(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList {
	s.EngineNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList) SetIdCardNo(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList {
	s.IdCardNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList) SetImportCertificateNo(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList {
	s.ImportCertificateNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList) SetInspectionListNo(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList {
	s.InspectionListNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList) SetMaxPassengers(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList {
	s.MaxPassengers = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList) SetOriginPlace(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList {
	s.OriginPlace = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList) SetPaymentVoucherNo(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList {
	s.PaymentVoucherNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList) SetTaxAuthorityName(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList {
	s.TaxAuthorityName = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList) SetTaxAuthorityNo(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList {
	s.TaxAuthorityNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList) SetTaxRate(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList {
	s.TaxRate = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList) SetTonnage(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList {
	s.Tonnage = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList) SetVehicleNo(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList {
	s.VehicleNo = &v
	return s
}

func (s *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList) SetVehicleType(v string) *UpdateInvoiceAbandonStatusRequestRedGeneralInvoiceVOVehicleSaleDetailVOList {
	s.VehicleType = &v
	return s
}

type UpdateInvoiceAbandonStatusResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateInvoiceAbandonStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAbandonStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAbandonStatusResponseBody) SetResult(v bool) *UpdateInvoiceAbandonStatusResponseBody {
	s.Result = &v
	return s
}

type UpdateInvoiceAbandonStatusResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInvoiceAbandonStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInvoiceAbandonStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAbandonStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAbandonStatusResponse) SetHeaders(v map[string]*string) *UpdateInvoiceAbandonStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateInvoiceAbandonStatusResponse) SetBody(v *UpdateInvoiceAbandonStatusResponseBody) *UpdateInvoiceAbandonStatusResponse {
	s.Body = v
	return s
}

type UpdateInvoiceAccountPeriodHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateInvoiceAccountPeriodHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAccountPeriodHeaders) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAccountPeriodHeaders) SetCommonHeaders(v map[string]*string) *UpdateInvoiceAccountPeriodHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateInvoiceAccountPeriodHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateInvoiceAccountPeriodHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateInvoiceAccountPeriodRequest struct {
	// 认证状态
	AccountPeriod *string `json:"accountPeriod,omitempty" xml:"accountPeriod,omitempty"`
	// 发票模型
	GeneralInvoiceVOList []*UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList `json:"generalInvoiceVOList,omitempty" xml:"generalInvoiceVOList,omitempty" type:"Repeated"`
	// 发票主键列表
	InvoiceKeyVOList []*UpdateInvoiceAccountPeriodRequestInvoiceKeyVOList `json:"invoiceKeyVOList,omitempty" xml:"invoiceKeyVOList,omitempty" type:"Repeated"`
	// 操作员
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
}

func (s UpdateInvoiceAccountPeriodRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAccountPeriodRequest) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAccountPeriodRequest) SetAccountPeriod(v string) *UpdateInvoiceAccountPeriodRequest {
	s.AccountPeriod = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequest) SetGeneralInvoiceVOList(v []*UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) *UpdateInvoiceAccountPeriodRequest {
	s.GeneralInvoiceVOList = v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequest) SetInvoiceKeyVOList(v []*UpdateInvoiceAccountPeriodRequestInvoiceKeyVOList) *UpdateInvoiceAccountPeriodRequest {
	s.InvoiceKeyVOList = v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequest) SetOperator(v string) *UpdateInvoiceAccountPeriodRequest {
	s.Operator = &v
	return s
}

type UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList struct {
	// 账期时间
	AccountPeriod *string `json:"accountPeriod,omitempty" xml:"accountPeriod,omitempty"`
	// 不含税金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 含税金额
	AmountWithTax *string `json:"amountWithTax,omitempty" xml:"amountWithTax,omitempty"`
	// 校验码
	CheckCode *string `json:"checkCode,omitempty" xml:"checkCode,omitempty"`
	// 查验时间
	CheckTime *string `json:"checkTime,omitempty" xml:"checkTime,omitempty"`
	// 开票日期
	DrewDate *string `json:"drewDate,omitempty" xml:"drewDate,omitempty"`
	// 电票版式文件下载地址
	ElectronicUrl *string `json:"electronicUrl,omitempty" xml:"electronicUrl,omitempty"`
	// 财务类型，INPUT-VAT(进项),OUTPUT_VAT(销项)
	FinanceType *string `json:"financeType,omitempty" xml:"financeType,omitempty"`
	// 资金类型 ，RED（红票），（BLUE）蓝票
	FundType *string `json:"fundType,omitempty" xml:"fundType,omitempty"`
	// 常规发票明细
	GeneralInvoiceDetailVOList []*UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList `json:"generalInvoiceDetailVOList,omitempty" xml:"generalInvoiceDetailVOList,omitempty" type:"Repeated"`
	// 发票代码
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	// 发票号码
	InvoiceNo *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
	// 发票状态
	InvoiceStatus *string `json:"invoiceStatus,omitempty" xml:"invoiceStatus,omitempty"`
	// 发票类型
	InvoiceType *string `json:"invoiceType,omitempty" xml:"invoiceType,omitempty"`
	// 机器码
	MachineCode *string `json:"machineCode,omitempty" xml:"machineCode,omitempty"`
	// 成品油标识
	OilFlag *string `json:"oilFlag,omitempty" xml:"oilFlag,omitempty"`
	// 收款人
	Payee *string `json:"payee,omitempty" xml:"payee,omitempty"`
	// 审批单实例
	ProcessInstCode *string `json:"processInstCode,omitempty" xml:"processInstCode,omitempty"`
	// 审批单类型
	ProcessInstType *string `json:"processInstType,omitempty" xml:"processInstType,omitempty"`
	// 购方地址
	PurchaserAddress *string `json:"purchaserAddress,omitempty" xml:"purchaserAddress,omitempty"`
	// 购方银行账户
	PurchaserBankAccount *string `json:"purchaserBankAccount,omitempty" xml:"purchaserBankAccount,omitempty"`
	// 购方银行名称
	PurchaserBankNameAccount *string `json:"purchaserBankNameAccount,omitempty" xml:"purchaserBankNameAccount,omitempty"`
	// 购方名称
	PurchaserName *string `json:"purchaserName,omitempty" xml:"purchaserName,omitempty"`
	// 购方税号
	PurchaserTaxNo *string `json:"purchaserTaxNo,omitempty" xml:"purchaserTaxNo,omitempty"`
	// 购方电话
	PurchaserTel *string `json:"purchaserTel,omitempty" xml:"purchaserTel,omitempty"`
	// 备注
	Remark                         *string                                                                                `json:"remark,omitempty" xml:"remark,omitempty"`
	SecondHandCarInvoiceDetailList []*UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList `json:"secondHandCarInvoiceDetailList,omitempty" xml:"secondHandCarInvoiceDetailList,omitempty" type:"Repeated"`
	// 销方地址
	SellerAddress *string `json:"sellerAddress,omitempty" xml:"sellerAddress,omitempty"`
	// 销方银行账户
	SellerBankAccount *string `json:"sellerBankAccount,omitempty" xml:"sellerBankAccount,omitempty"`
	// 销方银行名称
	SellerBankNameAccount *string `json:"sellerBankNameAccount,omitempty" xml:"sellerBankNameAccount,omitempty"`
	// 销方名称
	SellerName *string `json:"sellerName,omitempty" xml:"sellerName,omitempty"`
	// 销方税号
	SellerTaxNo *string `json:"sellerTaxNo,omitempty" xml:"sellerTaxNo,omitempty"`
	// 销方电话
	SellerTel *string `json:"sellerTel,omitempty" xml:"sellerTel,omitempty"`
	// 代开发票标识 1-自开，2-代开
	SupplySign *string `json:"supplySign,omitempty" xml:"supplySign,omitempty"`
	// 税额
	TaxAmount                   *string                                                                             `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	UsedVehicleSaleDetailVOList []*UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList `json:"usedVehicleSaleDetailVOList,omitempty" xml:"usedVehicleSaleDetailVOList,omitempty" type:"Repeated"`
	VehicleSaleDetailVOList     []*UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList     `json:"vehicleSaleDetailVOList,omitempty" xml:"vehicleSaleDetailVOList,omitempty" type:"Repeated"`
	// 发票查验状态
	VerifyStatus *string `json:"verifyStatus,omitempty" xml:"verifyStatus,omitempty"`
	// 凭证code
	VoucherCode *string `json:"voucherCode,omitempty" xml:"voucherCode,omitempty"`
	// 生成凭证状态
	VoucherStatus *string `json:"voucherStatus,omitempty" xml:"voucherStatus,omitempty"`
}

func (s UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetAccountPeriod(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.AccountPeriod = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetAmount(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.Amount = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetAmountWithTax(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.AmountWithTax = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetCheckCode(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.CheckCode = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetCheckTime(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.CheckTime = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetDrewDate(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.DrewDate = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetElectronicUrl(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.ElectronicUrl = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetFinanceType(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.FinanceType = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetFundType(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.FundType = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetGeneralInvoiceDetailVOList(v []*UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.GeneralInvoiceDetailVOList = v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetInvoiceCode(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.InvoiceCode = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetInvoiceNo(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.InvoiceNo = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetInvoiceStatus(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.InvoiceStatus = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetInvoiceType(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.InvoiceType = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetMachineCode(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.MachineCode = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetOilFlag(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.OilFlag = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetPayee(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.Payee = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetProcessInstCode(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.ProcessInstCode = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetProcessInstType(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.ProcessInstType = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetPurchaserAddress(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.PurchaserAddress = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetPurchaserBankAccount(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.PurchaserBankAccount = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetPurchaserBankNameAccount(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.PurchaserBankNameAccount = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetPurchaserName(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.PurchaserName = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetPurchaserTaxNo(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.PurchaserTaxNo = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetPurchaserTel(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.PurchaserTel = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetRemark(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.Remark = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetSecondHandCarInvoiceDetailList(v []*UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.SecondHandCarInvoiceDetailList = v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetSellerAddress(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.SellerAddress = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetSellerBankAccount(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.SellerBankAccount = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetSellerBankNameAccount(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.SellerBankNameAccount = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetSellerName(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.SellerName = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetSellerTaxNo(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.SellerTaxNo = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetSellerTel(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.SellerTel = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetSupplySign(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.SupplySign = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetTaxAmount(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.TaxAmount = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetUsedVehicleSaleDetailVOList(v []*UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.UsedVehicleSaleDetailVOList = v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetVehicleSaleDetailVOList(v []*UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.VehicleSaleDetailVOList = v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetVerifyStatus(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.VerifyStatus = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetVoucherCode(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.VoucherCode = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList) SetVoucherStatus(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOList {
	s.VoucherStatus = &v
	return s
}

type UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList struct {
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 商品名称
	GoodsName *string `json:"goodsName,omitempty" xml:"goodsName,omitempty"`
	// 数量
	Quantity *string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 税收分类编码
	RevenueCode *string `json:"revenueCode,omitempty" xml:"revenueCode,omitempty"`
	// 行号
	RowNo *string `json:"rowNo,omitempty" xml:"rowNo,omitempty"`
	// 规格型号
	Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
	// 税额
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	// 是否享受税收优惠：0-不享受，1-享受
	TaxPre *string `json:"taxPre,omitempty" xml:"taxPre,omitempty"`
	// 优惠政策类型
	TaxPreType *string `json:"taxPreType,omitempty" xml:"taxPreType,omitempty"`
	// 税率
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// 单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 单价
	UnitPrice *string `json:"unitPrice,omitempty" xml:"unitPrice,omitempty"`
}

func (s UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetAmount(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.Amount = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetGoodsName(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.GoodsName = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetQuantity(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.Quantity = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetRevenueCode(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.RevenueCode = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetRowNo(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.RowNo = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetSpecification(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.Specification = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetTaxAmount(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.TaxAmount = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetTaxPre(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.TaxPre = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetTaxPreType(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.TaxPreType = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetTaxRate(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.TaxRate = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetUnit(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.Unit = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetUnitPrice(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.UnitPrice = &v
	return s
}

type UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList struct {
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 车牌号
	CardNo *string `json:"cardNo,omitempty" xml:"cardNo,omitempty"`
	// 通行日期止
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// 商品名称
	GoodsName *string `json:"goodsName,omitempty" xml:"goodsName,omitempty"`
	// 税收分类编码
	RevenueCode *string `json:"revenueCode,omitempty" xml:"revenueCode,omitempty"`
	// 行号
	RowNo *string `json:"rowNo,omitempty" xml:"rowNo,omitempty"`
	// 通行日期起
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// 税额
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	// 税率
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// 类型
	VehicleType *string `json:"vehicleType,omitempty" xml:"vehicleType,omitempty"`
}

func (s UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetAmount(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.Amount = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetCardNo(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.CardNo = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetEndDate(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.EndDate = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetGoodsName(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.GoodsName = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetRevenueCode(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.RevenueCode = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetRowNo(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.RowNo = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetStartDate(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.StartDate = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetTaxAmount(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.TaxAmount = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetTaxRate(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.TaxRate = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetVehicleType(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.VehicleType = &v
	return s
}

type UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList struct {
	// 经营、拍卖单位
	AuctionUnit *string `json:"auctionUnit,omitempty" xml:"auctionUnit,omitempty"`
	// 经营、拍卖单位地址
	AuctionUnitAddress *string `json:"auctionUnitAddress,omitempty" xml:"auctionUnitAddress,omitempty"`
	// 经营、拍卖单位银行
	AuctionUnitBank *string `json:"auctionUnitBank,omitempty" xml:"auctionUnitBank,omitempty"`
	// 经营、拍卖单位税号
	AuctionUnitTaxNo *string `json:"auctionUnitTaxNo,omitempty" xml:"auctionUnitTaxNo,omitempty"`
	// 经营、拍卖单位电话
	AuctionUtilTel *string `json:"auctionUtilTel,omitempty" xml:"auctionUtilTel,omitempty"`
	// 厂牌型号
	CarModel *string `json:"carModel,omitempty" xml:"carModel,omitempty"`
	// 车牌照号
	CardNo *string `json:"cardNo,omitempty" xml:"cardNo,omitempty"`
	// 登记证号
	Registration *string `json:"registration,omitempty" xml:"registration,omitempty"`
	// 转入地车辆管理所名称
	TransferVehicle *string `json:"transferVehicle,omitempty" xml:"transferVehicle,omitempty"`
	// 二手车市场地址
	UsedCarAddress *string `json:"usedCarAddress,omitempty" xml:"usedCarAddress,omitempty"`
	// 二手车市场
	UsedCarMarket *string `json:"usedCarMarket,omitempty" xml:"usedCarMarket,omitempty"`
	// 二手车市场开户银行、账号
	UsedCarMarketBank *string `json:"usedCarMarketBank,omitempty" xml:"usedCarMarketBank,omitempty"`
	// 二手车市场电话
	UsedCarMarketPhone *string `json:"usedCarMarketPhone,omitempty" xml:"usedCarMarketPhone,omitempty"`
	// 二手车市场纳税人识别号
	UsedCarMarketTaxNo *string `json:"usedCarMarketTaxNo,omitempty" xml:"usedCarMarketTaxNo,omitempty"`
	// 车架号/车辆识别号
	VehicleNo *string `json:"vehicleNo,omitempty" xml:"vehicleNo,omitempty"`
	// 车辆类型
	VehicleType *string `json:"vehicleType,omitempty" xml:"vehicleType,omitempty"`
}

func (s UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetAuctionUnit(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.AuctionUnit = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetAuctionUnitAddress(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.AuctionUnitAddress = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetAuctionUnitBank(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.AuctionUnitBank = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetAuctionUnitTaxNo(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.AuctionUnitTaxNo = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetAuctionUtilTel(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.AuctionUtilTel = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetCarModel(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.CarModel = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetCardNo(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.CardNo = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetRegistration(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.Registration = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetTransferVehicle(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.TransferVehicle = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetUsedCarAddress(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.UsedCarAddress = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetUsedCarMarket(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.UsedCarMarket = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetUsedCarMarketBank(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.UsedCarMarketBank = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetUsedCarMarketPhone(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.UsedCarMarketPhone = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetUsedCarMarketTaxNo(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.UsedCarMarketTaxNo = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetVehicleNo(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.VehicleNo = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetVehicleType(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.VehicleType = &v
	return s
}

type UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList struct {
	// 品牌
	Brand *string `json:"brand,omitempty" xml:"brand,omitempty"`
	// 合格证号
	CertificateNo *string `json:"certificateNo,omitempty" xml:"certificateNo,omitempty"`
	// 发动机号
	EngineNo *string `json:"engineNo,omitempty" xml:"engineNo,omitempty"`
	// 身份证号/组织机构代码
	IdCardNo *string `json:"idCardNo,omitempty" xml:"idCardNo,omitempty"`
	// 进口证书号
	ImportCertificateNo *string `json:"importCertificateNo,omitempty" xml:"importCertificateNo,omitempty"`
	// 商检单号
	InspectionListNo *string `json:"inspectionListNo,omitempty" xml:"inspectionListNo,omitempty"`
	// 限乘人数
	MaxPassengers *string `json:"maxPassengers,omitempty" xml:"maxPassengers,omitempty"`
	// 产地
	OriginPlace *string `json:"originPlace,omitempty" xml:"originPlace,omitempty"`
	// 完税凭证号码
	PaymentVoucherNo *string `json:"paymentVoucherNo,omitempty" xml:"paymentVoucherNo,omitempty"`
	// 主管税务机关名称
	TaxAuthorityName *string `json:"taxAuthorityName,omitempty" xml:"taxAuthorityName,omitempty"`
	// 主管税务机关代码
	TaxAuthorityNo *string `json:"taxAuthorityNo,omitempty" xml:"taxAuthorityNo,omitempty"`
	// 税率
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// 吨位
	Tonnage *string `json:"tonnage,omitempty" xml:"tonnage,omitempty"`
	// 车架号码
	VehicleNo *string `json:"vehicleNo,omitempty" xml:"vehicleNo,omitempty"`
	// 车辆类型
	VehicleType *string `json:"vehicleType,omitempty" xml:"vehicleType,omitempty"`
}

func (s UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetBrand(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.Brand = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetCertificateNo(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.CertificateNo = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetEngineNo(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.EngineNo = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetIdCardNo(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.IdCardNo = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetImportCertificateNo(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.ImportCertificateNo = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetInspectionListNo(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.InspectionListNo = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetMaxPassengers(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.MaxPassengers = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetOriginPlace(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.OriginPlace = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetPaymentVoucherNo(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.PaymentVoucherNo = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetTaxAuthorityName(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.TaxAuthorityName = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetTaxAuthorityNo(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.TaxAuthorityNo = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetTaxRate(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.TaxRate = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetTonnage(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.Tonnage = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetVehicleNo(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.VehicleNo = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetVehicleType(v string) *UpdateInvoiceAccountPeriodRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.VehicleType = &v
	return s
}

type UpdateInvoiceAccountPeriodRequestInvoiceKeyVOList struct {
	// 发票编码
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	// 发票号码
	InvoiceNo *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
}

func (s UpdateInvoiceAccountPeriodRequestInvoiceKeyVOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAccountPeriodRequestInvoiceKeyVOList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAccountPeriodRequestInvoiceKeyVOList) SetInvoiceCode(v string) *UpdateInvoiceAccountPeriodRequestInvoiceKeyVOList {
	s.InvoiceCode = &v
	return s
}

func (s *UpdateInvoiceAccountPeriodRequestInvoiceKeyVOList) SetInvoiceNo(v string) *UpdateInvoiceAccountPeriodRequestInvoiceKeyVOList {
	s.InvoiceNo = &v
	return s
}

type UpdateInvoiceAccountPeriodResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateInvoiceAccountPeriodResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAccountPeriodResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAccountPeriodResponseBody) SetResult(v bool) *UpdateInvoiceAccountPeriodResponseBody {
	s.Result = &v
	return s
}

type UpdateInvoiceAccountPeriodResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInvoiceAccountPeriodResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInvoiceAccountPeriodResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAccountPeriodResponse) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAccountPeriodResponse) SetHeaders(v map[string]*string) *UpdateInvoiceAccountPeriodResponse {
	s.Headers = v
	return s
}

func (s *UpdateInvoiceAccountPeriodResponse) SetBody(v *UpdateInvoiceAccountPeriodResponseBody) *UpdateInvoiceAccountPeriodResponse {
	s.Body = v
	return s
}

type UpdateInvoiceAndReceiptRelatedHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateInvoiceAndReceiptRelatedHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAndReceiptRelatedHeaders) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAndReceiptRelatedHeaders) SetCommonHeaders(v map[string]*string) *UpdateInvoiceAndReceiptRelatedHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateInvoiceAndReceiptRelatedHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateInvoiceAndReceiptRelatedRequest struct {
	// 发票全票面信息
	GeneralInvoiceVO *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO `json:"generalInvoiceVO,omitempty" xml:"generalInvoiceVO,omitempty" type:"Struct"`
	// 发票代码
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	// 发票号码
	InvoiceNo *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
	// 操作员
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 钉钉审批单号
	ReceiptCode *string `json:"receiptCode,omitempty" xml:"receiptCode,omitempty"`
}

func (s UpdateInvoiceAndReceiptRelatedRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAndReceiptRelatedRequest) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAndReceiptRelatedRequest) SetGeneralInvoiceVO(v *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) *UpdateInvoiceAndReceiptRelatedRequest {
	s.GeneralInvoiceVO = v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequest) SetInvoiceCode(v string) *UpdateInvoiceAndReceiptRelatedRequest {
	s.InvoiceCode = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequest) SetInvoiceNo(v string) *UpdateInvoiceAndReceiptRelatedRequest {
	s.InvoiceNo = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequest) SetOperator(v string) *UpdateInvoiceAndReceiptRelatedRequest {
	s.Operator = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequest) SetReceiptCode(v string) *UpdateInvoiceAndReceiptRelatedRequest {
	s.ReceiptCode = &v
	return s
}

type UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO struct {
	// 账期时间
	AccountPeriod *string `json:"accountPeriod,omitempty" xml:"accountPeriod,omitempty"`
	// 不含税金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 含税金额
	AmountWithTax *string `json:"amountWithTax,omitempty" xml:"amountWithTax,omitempty"`
	// 校验码
	CheckCode *string `json:"checkCode,omitempty" xml:"checkCode,omitempty"`
	// 查验时间
	CheckTime *string `json:"checkTime,omitempty" xml:"checkTime,omitempty"`
	// 开票日期
	DrewDate *string `json:"drewDate,omitempty" xml:"drewDate,omitempty"`
	// 电票版式文件下载地址
	ElectronicUrl *string `json:"electronicUrl,omitempty" xml:"electronicUrl,omitempty"`
	// 财务类型，INPUT-VAT(进项),OUTPUT_VAT(销项)
	FinanceType *string `json:"financeType,omitempty" xml:"financeType,omitempty"`
	// 资金类型 ，RED（红票），（BLUE）蓝票
	FundType *string `json:"fundType,omitempty" xml:"fundType,omitempty"`
	// 常规发票明细
	GeneralInvoiceDetailVOList []*UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList `json:"generalInvoiceDetailVOList,omitempty" xml:"generalInvoiceDetailVOList,omitempty" type:"Repeated"`
	// 发票代码
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	// 发票号码
	InvoiceNo *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
	// 发票状态
	InvoiceStatus *string `json:"invoiceStatus,omitempty" xml:"invoiceStatus,omitempty"`
	// 发票类型
	InvoiceType *string `json:"invoiceType,omitempty" xml:"invoiceType,omitempty"`
	// 机器码
	MachineCode *string `json:"machineCode,omitempty" xml:"machineCode,omitempty"`
	// 成品油标识
	OilFlag *string `json:"oilFlag,omitempty" xml:"oilFlag,omitempty"`
	// 收款人
	Payee *string `json:"payee,omitempty" xml:"payee,omitempty"`
	// 审批单实例
	ProcessInstCode *string `json:"processInstCode,omitempty" xml:"processInstCode,omitempty"`
	// 审批单类型
	ProcessInstType *string `json:"processInstType,omitempty" xml:"processInstType,omitempty"`
	// 购方地址
	PurchaserAddress *string `json:"purchaserAddress,omitempty" xml:"purchaserAddress,omitempty"`
	// 购方银行账户
	//
	PurchaserBankAccount     *string `json:"purchaserBankAccount,omitempty" xml:"purchaserBankAccount,omitempty"`
	PurchaserBankNameAccount *string `json:"purchaserBankNameAccount,omitempty" xml:"purchaserBankNameAccount,omitempty"`
	// 购方名称
	PurchaserName *string `json:"purchaserName,omitempty" xml:"purchaserName,omitempty"`
	// 购方税号
	PurchaserTaxNo *string `json:"purchaserTaxNo,omitempty" xml:"purchaserTaxNo,omitempty"`
	// 购方电话
	PurchaserTel *string `json:"purchaserTel,omitempty" xml:"purchaserTel,omitempty"`
	// 备注
	Remark                         *string                                                                                `json:"remark,omitempty" xml:"remark,omitempty"`
	SecondHandCarInvoiceDetailList []*UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOSecondHandCarInvoiceDetailList `json:"secondHandCarInvoiceDetailList,omitempty" xml:"secondHandCarInvoiceDetailList,omitempty" type:"Repeated"`
	// 销方地址
	SellerAddress *string `json:"sellerAddress,omitempty" xml:"sellerAddress,omitempty"`
	// 销方银行账户
	SellerBankAccount     *string `json:"sellerBankAccount,omitempty" xml:"sellerBankAccount,omitempty"`
	SellerBankNameAccount *string `json:"sellerBankNameAccount,omitempty" xml:"sellerBankNameAccount,omitempty"`
	// 销方名称
	SellerName *string `json:"sellerName,omitempty" xml:"sellerName,omitempty"`
	// 销方税号
	SellerTaxNo *string `json:"sellerTaxNo,omitempty" xml:"sellerTaxNo,omitempty"`
	// 销方电话
	SellerTel *string `json:"sellerTel,omitempty" xml:"sellerTel,omitempty"`
	// 代开发票标识 1-自开，2-代开
	SupplySign *string `json:"supplySign,omitempty" xml:"supplySign,omitempty"`
	// 税额
	TaxAmount                   *string                                                                             `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	UsedVehicleSaleDetailVOList []*UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList `json:"usedVehicleSaleDetailVOList,omitempty" xml:"usedVehicleSaleDetailVOList,omitempty" type:"Repeated"`
	VehicleSaleDetailVOList     []*UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList     `json:"vehicleSaleDetailVOList,omitempty" xml:"vehicleSaleDetailVOList,omitempty" type:"Repeated"`
	// 发票查验状态
	VerifyStatus *string `json:"verifyStatus,omitempty" xml:"verifyStatus,omitempty"`
	// 凭证code
	VoucherCode *string `json:"voucherCode,omitempty" xml:"voucherCode,omitempty"`
	// 生成凭证状态
	VoucherStatus *string `json:"voucherStatus,omitempty" xml:"voucherStatus,omitempty"`
}

func (s UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetAccountPeriod(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.AccountPeriod = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetAmount(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.Amount = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetAmountWithTax(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.AmountWithTax = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetCheckCode(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.CheckCode = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetCheckTime(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.CheckTime = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetDrewDate(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.DrewDate = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetElectronicUrl(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.ElectronicUrl = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetFinanceType(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.FinanceType = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetFundType(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.FundType = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetGeneralInvoiceDetailVOList(v []*UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.GeneralInvoiceDetailVOList = v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetInvoiceCode(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.InvoiceCode = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetInvoiceNo(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.InvoiceNo = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetInvoiceStatus(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.InvoiceStatus = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetInvoiceType(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.InvoiceType = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetMachineCode(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.MachineCode = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetOilFlag(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.OilFlag = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetPayee(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.Payee = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetProcessInstCode(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.ProcessInstCode = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetProcessInstType(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.ProcessInstType = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetPurchaserAddress(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.PurchaserAddress = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetPurchaserBankAccount(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.PurchaserBankAccount = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetPurchaserBankNameAccount(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.PurchaserBankNameAccount = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetPurchaserName(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.PurchaserName = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetPurchaserTaxNo(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.PurchaserTaxNo = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetPurchaserTel(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.PurchaserTel = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetRemark(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.Remark = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetSecondHandCarInvoiceDetailList(v []*UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOSecondHandCarInvoiceDetailList) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.SecondHandCarInvoiceDetailList = v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetSellerAddress(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.SellerAddress = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetSellerBankAccount(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.SellerBankAccount = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetSellerBankNameAccount(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.SellerBankNameAccount = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetSellerName(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.SellerName = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetSellerTaxNo(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.SellerTaxNo = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetSellerTel(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.SellerTel = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetSupplySign(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.SupplySign = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetTaxAmount(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.TaxAmount = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetUsedVehicleSaleDetailVOList(v []*UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.UsedVehicleSaleDetailVOList = v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetVehicleSaleDetailVOList(v []*UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.VehicleSaleDetailVOList = v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetVerifyStatus(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.VerifyStatus = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetVoucherCode(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.VoucherCode = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO) SetVoucherStatus(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVO {
	s.VoucherStatus = &v
	return s
}

type UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList struct {
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 商品名称
	GoodsName *string `json:"goodsName,omitempty" xml:"goodsName,omitempty"`
	// 数量
	Quantity *string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 税收分类编码
	RevenueCode *string `json:"revenueCode,omitempty" xml:"revenueCode,omitempty"`
	// 行号
	RowNo *string `json:"rowNo,omitempty" xml:"rowNo,omitempty"`
	// 规格型号
	Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
	// 税额
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	// 是否享受税收优惠：0-不享受，1-享受
	TaxPre *string `json:"taxPre,omitempty" xml:"taxPre,omitempty"`
	// 优惠政策类型
	TaxPreType *string `json:"taxPreType,omitempty" xml:"taxPreType,omitempty"`
	// 税率
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// 单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 单价
	UnitPrice *string `json:"unitPrice,omitempty" xml:"unitPrice,omitempty"`
}

func (s UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList) SetAmount(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.Amount = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList) SetGoodsName(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.GoodsName = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList) SetQuantity(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.Quantity = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList) SetRevenueCode(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.RevenueCode = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList) SetRowNo(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.RowNo = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList) SetSpecification(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.Specification = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList) SetTaxAmount(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.TaxAmount = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList) SetTaxPre(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.TaxPre = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList) SetTaxPreType(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.TaxPreType = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList) SetTaxRate(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.TaxRate = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList) SetUnit(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.Unit = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList) SetUnitPrice(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOGeneralInvoiceDetailVOList {
	s.UnitPrice = &v
	return s
}

type UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOSecondHandCarInvoiceDetailList struct {
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 车牌号
	CardNo *string `json:"cardNo,omitempty" xml:"cardNo,omitempty"`
	// 通行日期止
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// 商品名称
	GoodsName *string `json:"goodsName,omitempty" xml:"goodsName,omitempty"`
	// 税收分类编码
	RevenueCode *string `json:"revenueCode,omitempty" xml:"revenueCode,omitempty"`
	// 行号
	RowNo *string `json:"rowNo,omitempty" xml:"rowNo,omitempty"`
	// 通行日期起
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// 税额
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	// 税率
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// 类型
	VehicleType *string `json:"vehicleType,omitempty" xml:"vehicleType,omitempty"`
}

func (s UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOSecondHandCarInvoiceDetailList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOSecondHandCarInvoiceDetailList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetAmount(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.Amount = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetCardNo(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.CardNo = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetEndDate(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.EndDate = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetGoodsName(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.GoodsName = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetRevenueCode(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.RevenueCode = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetRowNo(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.RowNo = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetStartDate(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.StartDate = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetTaxAmount(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.TaxAmount = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetTaxRate(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.TaxRate = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOSecondHandCarInvoiceDetailList) SetVehicleType(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOSecondHandCarInvoiceDetailList {
	s.VehicleType = &v
	return s
}

type UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList struct {
	// 经营、拍卖单位
	AuctionUnit *string `json:"auctionUnit,omitempty" xml:"auctionUnit,omitempty"`
	// 经营、拍卖单位地址
	AuctionUnitAddress *string `json:"auctionUnitAddress,omitempty" xml:"auctionUnitAddress,omitempty"`
	// 经营、拍卖单位银行
	AuctionUnitBank *string `json:"auctionUnitBank,omitempty" xml:"auctionUnitBank,omitempty"`
	// 经营、拍卖单位税号
	AuctionUnitTaxNo *string `json:"auctionUnitTaxNo,omitempty" xml:"auctionUnitTaxNo,omitempty"`
	// 经营、拍卖单位电话
	AuctionUtilTel *string `json:"auctionUtilTel,omitempty" xml:"auctionUtilTel,omitempty"`
	// 厂牌型号
	CarModel *string `json:"carModel,omitempty" xml:"carModel,omitempty"`
	// 车牌照号
	CardNo *string `json:"cardNo,omitempty" xml:"cardNo,omitempty"`
	// 登记证号
	Registration *string `json:"registration,omitempty" xml:"registration,omitempty"`
	// 转入地车辆管理所名称
	TransferVehicle *string `json:"transferVehicle,omitempty" xml:"transferVehicle,omitempty"`
	// 二手车市场地址
	UsedCarAddress *string `json:"usedCarAddress,omitempty" xml:"usedCarAddress,omitempty"`
	// 二手车市场
	UsedCarMarket *string `json:"usedCarMarket,omitempty" xml:"usedCarMarket,omitempty"`
	// 二手车市场开户银行、账号
	UsedCarMarketBank *string `json:"usedCarMarketBank,omitempty" xml:"usedCarMarketBank,omitempty"`
	// 二手车市场电话
	UsedCarMarketPhone *string `json:"usedCarMarketPhone,omitempty" xml:"usedCarMarketPhone,omitempty"`
	// 二手车市场纳税人识别号
	UsedCarMarketTaxNo *string `json:"usedCarMarketTaxNo,omitempty" xml:"usedCarMarketTaxNo,omitempty"`
	// 车架号/车辆识别号
	VehicleNo *string `json:"vehicleNo,omitempty" xml:"vehicleNo,omitempty"`
	// 车辆类型
	VehicleType *string `json:"vehicleType,omitempty" xml:"vehicleType,omitempty"`
}

func (s UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetAuctionUnit(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.AuctionUnit = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetAuctionUnitAddress(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.AuctionUnitAddress = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetAuctionUnitBank(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.AuctionUnitBank = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetAuctionUnitTaxNo(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.AuctionUnitTaxNo = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetAuctionUtilTel(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.AuctionUtilTel = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetCarModel(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.CarModel = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetCardNo(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.CardNo = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetRegistration(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.Registration = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetTransferVehicle(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.TransferVehicle = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetUsedCarAddress(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.UsedCarAddress = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetUsedCarMarket(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.UsedCarMarket = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetUsedCarMarketBank(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.UsedCarMarketBank = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetUsedCarMarketPhone(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.UsedCarMarketPhone = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetUsedCarMarketTaxNo(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.UsedCarMarketTaxNo = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetVehicleNo(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.VehicleNo = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList) SetVehicleType(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOUsedVehicleSaleDetailVOList {
	s.VehicleType = &v
	return s
}

type UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList struct {
	// 品牌
	Brand *string `json:"brand,omitempty" xml:"brand,omitempty"`
	// 合格证号
	CertificateNo *string `json:"certificateNo,omitempty" xml:"certificateNo,omitempty"`
	// 发动机号
	EngineNo *string `json:"engineNo,omitempty" xml:"engineNo,omitempty"`
	// 身份证号/组织机构代码
	IdCardNo *string `json:"idCardNo,omitempty" xml:"idCardNo,omitempty"`
	// 进口证书号
	ImportCertificateNo *string `json:"importCertificateNo,omitempty" xml:"importCertificateNo,omitempty"`
	// 商检单号
	InspectionListNo *string `json:"inspectionListNo,omitempty" xml:"inspectionListNo,omitempty"`
	// 限乘人数
	MaxPassengers *string `json:"maxPassengers,omitempty" xml:"maxPassengers,omitempty"`
	// 产地
	OriginPlace *string `json:"originPlace,omitempty" xml:"originPlace,omitempty"`
	// 完税凭证号码
	PaymentVoucherNo *string `json:"paymentVoucherNo,omitempty" xml:"paymentVoucherNo,omitempty"`
	// 主管税务机关名称
	TaxAuthorityName *string `json:"taxAuthorityName,omitempty" xml:"taxAuthorityName,omitempty"`
	// 主管税务机关代码
	TaxAuthorityNo *string `json:"taxAuthorityNo,omitempty" xml:"taxAuthorityNo,omitempty"`
	// 税率
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// 吨位
	Tonnage *string `json:"tonnage,omitempty" xml:"tonnage,omitempty"`
	// 车架号码
	VehicleNo *string `json:"vehicleNo,omitempty" xml:"vehicleNo,omitempty"`
	// 车辆类型
	VehicleType *string `json:"vehicleType,omitempty" xml:"vehicleType,omitempty"`
}

func (s UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList) SetBrand(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList {
	s.Brand = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList) SetCertificateNo(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList {
	s.CertificateNo = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList) SetEngineNo(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList {
	s.EngineNo = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList) SetIdCardNo(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList {
	s.IdCardNo = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList) SetImportCertificateNo(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList {
	s.ImportCertificateNo = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList) SetInspectionListNo(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList {
	s.InspectionListNo = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList) SetMaxPassengers(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList {
	s.MaxPassengers = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList) SetOriginPlace(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList {
	s.OriginPlace = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList) SetPaymentVoucherNo(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList {
	s.PaymentVoucherNo = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList) SetTaxAuthorityName(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList {
	s.TaxAuthorityName = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList) SetTaxAuthorityNo(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList {
	s.TaxAuthorityNo = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList) SetTaxRate(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList {
	s.TaxRate = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList) SetTonnage(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList {
	s.Tonnage = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList) SetVehicleNo(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList {
	s.VehicleNo = &v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList) SetVehicleType(v string) *UpdateInvoiceAndReceiptRelatedRequestGeneralInvoiceVOVehicleSaleDetailVOList {
	s.VehicleType = &v
	return s
}

type UpdateInvoiceAndReceiptRelatedResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateInvoiceAndReceiptRelatedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAndReceiptRelatedResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAndReceiptRelatedResponseBody) SetResult(v bool) *UpdateInvoiceAndReceiptRelatedResponseBody {
	s.Result = &v
	return s
}

type UpdateInvoiceAndReceiptRelatedResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInvoiceAndReceiptRelatedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInvoiceAndReceiptRelatedResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceAndReceiptRelatedResponse) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceAndReceiptRelatedResponse) SetHeaders(v map[string]*string) *UpdateInvoiceAndReceiptRelatedResponse {
	s.Headers = v
	return s
}

func (s *UpdateInvoiceAndReceiptRelatedResponse) SetBody(v *UpdateInvoiceAndReceiptRelatedResponseBody) *UpdateInvoiceAndReceiptRelatedResponse {
	s.Body = v
	return s
}

type UpdateInvoiceIgnoreStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateInvoiceIgnoreStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceIgnoreStatusHeaders) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceIgnoreStatusHeaders) SetCommonHeaders(v map[string]*string) *UpdateInvoiceIgnoreStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateInvoiceIgnoreStatusHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateInvoiceIgnoreStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateInvoiceIgnoreStatusRequest struct {
	// 审批单id
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// 操作员
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateInvoiceIgnoreStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceIgnoreStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceIgnoreStatusRequest) SetInstanceId(v string) *UpdateInvoiceIgnoreStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInvoiceIgnoreStatusRequest) SetOperator(v string) *UpdateInvoiceIgnoreStatusRequest {
	s.Operator = &v
	return s
}

func (s *UpdateInvoiceIgnoreStatusRequest) SetStatus(v string) *UpdateInvoiceIgnoreStatusRequest {
	s.Status = &v
	return s
}

type UpdateInvoiceIgnoreStatusResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateInvoiceIgnoreStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceIgnoreStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceIgnoreStatusResponseBody) SetResult(v bool) *UpdateInvoiceIgnoreStatusResponseBody {
	s.Result = &v
	return s
}

type UpdateInvoiceIgnoreStatusResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInvoiceIgnoreStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInvoiceIgnoreStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceIgnoreStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceIgnoreStatusResponse) SetHeaders(v map[string]*string) *UpdateInvoiceIgnoreStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateInvoiceIgnoreStatusResponse) SetBody(v *UpdateInvoiceIgnoreStatusResponseBody) *UpdateInvoiceIgnoreStatusResponse {
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
	// 抵扣状态
	//
	DeductStatus *string `json:"deductStatus,omitempty" xml:"deductStatus,omitempty"`
	// 发票模型
	GeneralInvoiceVOList []*UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList `json:"generalInvoiceVOList,omitempty" xml:"generalInvoiceVOList,omitempty" type:"Repeated"`
	// 待更新
	InvoiceKeyVOList []*UpdateInvoiceVerifyStatusRequestInvoiceKeyVOList `json:"invoiceKeyVOList,omitempty" xml:"invoiceKeyVOList,omitempty" type:"Repeated"`
	// 操作员
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 认证状态
	VerifyStatus *string `json:"verifyStatus,omitempty" xml:"verifyStatus,omitempty"`
}

func (s UpdateInvoiceVerifyStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceVerifyStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceVerifyStatusRequest) SetDeductStatus(v string) *UpdateInvoiceVerifyStatusRequest {
	s.DeductStatus = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequest) SetGeneralInvoiceVOList(v []*UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) *UpdateInvoiceVerifyStatusRequest {
	s.GeneralInvoiceVOList = v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequest) SetInvoiceKeyVOList(v []*UpdateInvoiceVerifyStatusRequestInvoiceKeyVOList) *UpdateInvoiceVerifyStatusRequest {
	s.InvoiceKeyVOList = v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequest) SetOperator(v string) *UpdateInvoiceVerifyStatusRequest {
	s.Operator = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequest) SetVerifyStatus(v string) *UpdateInvoiceVerifyStatusRequest {
	s.VerifyStatus = &v
	return s
}

type UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList struct {
	// 账期时间
	AccountPeriod *string `json:"accountPeriod,omitempty" xml:"accountPeriod,omitempty"`
	// 不含税金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 含税金额
	AmountWithTax *string `json:"amountWithTax,omitempty" xml:"amountWithTax,omitempty"`
	// 校验码
	CheckCode *string `json:"checkCode,omitempty" xml:"checkCode,omitempty"`
	// 查验时间
	CheckTime *string `json:"checkTime,omitempty" xml:"checkTime,omitempty"`
	// 开票日期
	DrewDate *string `json:"drewDate,omitempty" xml:"drewDate,omitempty"`
	// 电票版式文件下载地址
	ElectronicUrl *string `json:"electronicUrl,omitempty" xml:"electronicUrl,omitempty"`
	// 财务类型，INPUT-VAT(进项),OUTPUT_VAT(销项)
	FinanceType *string `json:"financeType,omitempty" xml:"financeType,omitempty"`
	// 资金类型 ，RED（红票），（BLUE）蓝票
	FundType *string `json:"fundType,omitempty" xml:"fundType,omitempty"`
	// 常规发票明细
	GeneralInvoiceDetailVOList []*UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList `json:"generalInvoiceDetailVOList,omitempty" xml:"generalInvoiceDetailVOList,omitempty" type:"Repeated"`
	// 发票代码
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	// 发票号码
	InvoiceNo *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
	// 发票状态
	InvoiceStatus *string `json:"invoiceStatus,omitempty" xml:"invoiceStatus,omitempty"`
	// 发票类型
	InvoiceType *string `json:"invoiceType,omitempty" xml:"invoiceType,omitempty"`
	// 机器码
	MachineCode *string `json:"machineCode,omitempty" xml:"machineCode,omitempty"`
	// 成品油标识
	OilFlag *string `json:"oilFlag,omitempty" xml:"oilFlag,omitempty"`
	// 收款人
	Payee *string `json:"payee,omitempty" xml:"payee,omitempty"`
	// 审批单实例
	ProcessInstCode *string `json:"processInstCode,omitempty" xml:"processInstCode,omitempty"`
	// 审批单类型
	ProcessInstType *string `json:"processInstType,omitempty" xml:"processInstType,omitempty"`
	// 购方地址
	PurchaserAddress *string `json:"purchaserAddress,omitempty" xml:"purchaserAddress,omitempty"`
	// 购方银行账户
	PurchaserBankAccount *string `json:"purchaserBankAccount,omitempty" xml:"purchaserBankAccount,omitempty"`
	// 购方银行名称
	PurchaserBankNameAccount *string `json:"purchaserBankNameAccount,omitempty" xml:"purchaserBankNameAccount,omitempty"`
	// 购方名称
	PurchaserName *string `json:"purchaserName,omitempty" xml:"purchaserName,omitempty"`
	// 购方税号
	PurchaserTaxNo *string `json:"purchaserTaxNo,omitempty" xml:"purchaserTaxNo,omitempty"`
	// 购方电话
	PurchaserTel *string `json:"purchaserTel,omitempty" xml:"purchaserTel,omitempty"`
	// 备注
	Remark                         *string                                                                               `json:"remark,omitempty" xml:"remark,omitempty"`
	SecondHandCarInvoiceDetailList []*UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList `json:"secondHandCarInvoiceDetailList,omitempty" xml:"secondHandCarInvoiceDetailList,omitempty" type:"Repeated"`
	// 销方地址
	SellerAddress *string `json:"sellerAddress,omitempty" xml:"sellerAddress,omitempty"`
	// 销方银行账户
	SellerBankAccount *string `json:"sellerBankAccount,omitempty" xml:"sellerBankAccount,omitempty"`
	// 销方银行名称
	SellerBankNameAccount *string `json:"sellerBankNameAccount,omitempty" xml:"sellerBankNameAccount,omitempty"`
	// 销方名称
	SellerName *string `json:"sellerName,omitempty" xml:"sellerName,omitempty"`
	// 销方税号
	SellerTaxNo *string `json:"sellerTaxNo,omitempty" xml:"sellerTaxNo,omitempty"`
	// 销方电话
	SellerTel *string `json:"sellerTel,omitempty" xml:"sellerTel,omitempty"`
	// 代开发票标识 1-自开，2-代开
	SupplySign *string `json:"supplySign,omitempty" xml:"supplySign,omitempty"`
	// 税额
	TaxAmount                   *string                                                                            `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	UsedVehicleSaleDetailVOList []*UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList `json:"usedVehicleSaleDetailVOList,omitempty" xml:"usedVehicleSaleDetailVOList,omitempty" type:"Repeated"`
	VehicleSaleDetailVOList     []*UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList     `json:"vehicleSaleDetailVOList,omitempty" xml:"vehicleSaleDetailVOList,omitempty" type:"Repeated"`
	// 发票查验状态
	VerifyStatus *string `json:"verifyStatus,omitempty" xml:"verifyStatus,omitempty"`
	// 凭证code
	VoucherCode *string `json:"voucherCode,omitempty" xml:"voucherCode,omitempty"`
	// 生成凭证状态
	VoucherStatus *string `json:"voucherStatus,omitempty" xml:"voucherStatus,omitempty"`
}

func (s UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetAccountPeriod(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.AccountPeriod = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetAmount(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.Amount = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetAmountWithTax(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.AmountWithTax = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetCheckCode(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.CheckCode = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetCheckTime(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.CheckTime = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetDrewDate(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.DrewDate = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetElectronicUrl(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.ElectronicUrl = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetFinanceType(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.FinanceType = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetFundType(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.FundType = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetGeneralInvoiceDetailVOList(v []*UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.GeneralInvoiceDetailVOList = v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetInvoiceCode(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.InvoiceCode = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetInvoiceNo(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.InvoiceNo = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetInvoiceStatus(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.InvoiceStatus = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetInvoiceType(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.InvoiceType = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetMachineCode(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.MachineCode = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetOilFlag(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.OilFlag = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetPayee(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.Payee = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetProcessInstCode(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.ProcessInstCode = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetProcessInstType(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.ProcessInstType = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetPurchaserAddress(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.PurchaserAddress = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetPurchaserBankAccount(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.PurchaserBankAccount = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetPurchaserBankNameAccount(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.PurchaserBankNameAccount = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetPurchaserName(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.PurchaserName = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetPurchaserTaxNo(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.PurchaserTaxNo = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetPurchaserTel(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.PurchaserTel = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetRemark(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.Remark = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetSecondHandCarInvoiceDetailList(v []*UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.SecondHandCarInvoiceDetailList = v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetSellerAddress(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.SellerAddress = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetSellerBankAccount(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.SellerBankAccount = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetSellerBankNameAccount(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.SellerBankNameAccount = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetSellerName(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.SellerName = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetSellerTaxNo(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.SellerTaxNo = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetSellerTel(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.SellerTel = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetSupplySign(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.SupplySign = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetTaxAmount(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.TaxAmount = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetUsedVehicleSaleDetailVOList(v []*UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.UsedVehicleSaleDetailVOList = v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetVehicleSaleDetailVOList(v []*UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.VehicleSaleDetailVOList = v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetVerifyStatus(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.VerifyStatus = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetVoucherCode(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.VoucherCode = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList) SetVoucherStatus(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOList {
	s.VoucherStatus = &v
	return s
}

type UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList struct {
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 商品名称
	GoodsName *string `json:"goodsName,omitempty" xml:"goodsName,omitempty"`
	// 数量
	Quantity *string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 税收分类编码
	RevenueCode *string `json:"revenueCode,omitempty" xml:"revenueCode,omitempty"`
	// 行号
	RowNo *string `json:"rowNo,omitempty" xml:"rowNo,omitempty"`
	// 规格型号
	Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
	// 税额
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	// 是否享受税收优惠：0-不享受，1-享受
	TaxPre *string `json:"taxPre,omitempty" xml:"taxPre,omitempty"`
	// 优惠政策类型
	TaxPreType *string `json:"taxPreType,omitempty" xml:"taxPreType,omitempty"`
	// 税率
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// 单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 单价
	UnitPrice *string `json:"unitPrice,omitempty" xml:"unitPrice,omitempty"`
}

func (s UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetAmount(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.Amount = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetGoodsName(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.GoodsName = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetQuantity(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.Quantity = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetRevenueCode(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.RevenueCode = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetRowNo(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.RowNo = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetSpecification(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.Specification = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetTaxAmount(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.TaxAmount = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetTaxPre(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.TaxPre = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetTaxPreType(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.TaxPreType = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetTaxRate(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.TaxRate = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetUnit(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.Unit = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList) SetUnitPrice(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListGeneralInvoiceDetailVOList {
	s.UnitPrice = &v
	return s
}

type UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList struct {
	// 金额
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 车牌号
	CardNo *string `json:"cardNo,omitempty" xml:"cardNo,omitempty"`
	// 通行日期止
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// 商品名称
	GoodsName *string `json:"goodsName,omitempty" xml:"goodsName,omitempty"`
	// 税收分类编码
	RevenueCode *string `json:"revenueCode,omitempty" xml:"revenueCode,omitempty"`
	// 行号
	RowNo *string `json:"rowNo,omitempty" xml:"rowNo,omitempty"`
	// 通行日期起
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// 税额
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	// 税率
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// 类型
	VehicleType *string `json:"vehicleType,omitempty" xml:"vehicleType,omitempty"`
}

func (s UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetAmount(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.Amount = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetCardNo(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.CardNo = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetEndDate(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.EndDate = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetGoodsName(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.GoodsName = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetRevenueCode(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.RevenueCode = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetRowNo(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.RowNo = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetStartDate(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.StartDate = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetTaxAmount(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.TaxAmount = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetTaxRate(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.TaxRate = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList) SetVehicleType(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListSecondHandCarInvoiceDetailList {
	s.VehicleType = &v
	return s
}

type UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList struct {
	// 经营、拍卖单位
	AuctionUnit *string `json:"auctionUnit,omitempty" xml:"auctionUnit,omitempty"`
	// 经营、拍卖单位地址
	AuctionUnitAddress *string `json:"auctionUnitAddress,omitempty" xml:"auctionUnitAddress,omitempty"`
	// 经营、拍卖单位银行
	AuctionUnitBank *string `json:"auctionUnitBank,omitempty" xml:"auctionUnitBank,omitempty"`
	// 经营、拍卖单位税号
	AuctionUnitTaxNo *string `json:"auctionUnitTaxNo,omitempty" xml:"auctionUnitTaxNo,omitempty"`
	// 经营、拍卖单位电话
	AuctionUtilTel *string `json:"auctionUtilTel,omitempty" xml:"auctionUtilTel,omitempty"`
	// 厂牌型号
	CarModel *string `json:"carModel,omitempty" xml:"carModel,omitempty"`
	// 车牌照号
	CardNo *string `json:"cardNo,omitempty" xml:"cardNo,omitempty"`
	// 登记证号
	Registration *string `json:"registration,omitempty" xml:"registration,omitempty"`
	// 转入地车辆管理所名称
	TransferVehicle *string `json:"transferVehicle,omitempty" xml:"transferVehicle,omitempty"`
	// 二手车市场地址
	UsedCarAddress *string `json:"usedCarAddress,omitempty" xml:"usedCarAddress,omitempty"`
	// 二手车市场
	UsedCarMarket *string `json:"usedCarMarket,omitempty" xml:"usedCarMarket,omitempty"`
	// 二手车市场开户银行、账号
	UsedCarMarketBank *string `json:"usedCarMarketBank,omitempty" xml:"usedCarMarketBank,omitempty"`
	// 二手车市场电话
	UsedCarMarketPhone *string `json:"usedCarMarketPhone,omitempty" xml:"usedCarMarketPhone,omitempty"`
	// 二手车市场纳税人识别号
	UsedCarMarketTaxNo *string `json:"usedCarMarketTaxNo,omitempty" xml:"usedCarMarketTaxNo,omitempty"`
	// 车架号/车辆识别号
	VehicleNo *string `json:"vehicleNo,omitempty" xml:"vehicleNo,omitempty"`
	// 车辆类型
	VehicleType *string `json:"vehicleType,omitempty" xml:"vehicleType,omitempty"`
}

func (s UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetAuctionUnit(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.AuctionUnit = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetAuctionUnitAddress(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.AuctionUnitAddress = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetAuctionUnitBank(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.AuctionUnitBank = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetAuctionUnitTaxNo(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.AuctionUnitTaxNo = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetAuctionUtilTel(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.AuctionUtilTel = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetCarModel(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.CarModel = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetCardNo(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.CardNo = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetRegistration(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.Registration = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetTransferVehicle(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.TransferVehicle = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetUsedCarAddress(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.UsedCarAddress = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetUsedCarMarket(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.UsedCarMarket = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetUsedCarMarketBank(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.UsedCarMarketBank = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetUsedCarMarketPhone(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.UsedCarMarketPhone = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetUsedCarMarketTaxNo(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.UsedCarMarketTaxNo = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetVehicleNo(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.VehicleNo = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList) SetVehicleType(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListUsedVehicleSaleDetailVOList {
	s.VehicleType = &v
	return s
}

type UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList struct {
	// 品牌
	Brand *string `json:"brand,omitempty" xml:"brand,omitempty"`
	// 合格证号
	CertificateNo *string `json:"certificateNo,omitempty" xml:"certificateNo,omitempty"`
	// 发动机号
	EngineNo *string `json:"engineNo,omitempty" xml:"engineNo,omitempty"`
	// 身份证号/组织机构代码
	IdCardNo *string `json:"idCardNo,omitempty" xml:"idCardNo,omitempty"`
	// 进口证书号
	ImportCertificateNo *string `json:"importCertificateNo,omitempty" xml:"importCertificateNo,omitempty"`
	// 商检单号
	InspectionListNo *string `json:"inspectionListNo,omitempty" xml:"inspectionListNo,omitempty"`
	// 限乘人数
	MaxPassengers *string `json:"maxPassengers,omitempty" xml:"maxPassengers,omitempty"`
	// 产地
	OriginPlace *string `json:"originPlace,omitempty" xml:"originPlace,omitempty"`
	// 完税凭证号码
	PaymentVoucherNo *string `json:"paymentVoucherNo,omitempty" xml:"paymentVoucherNo,omitempty"`
	// 主管税务机关名称
	TaxAuthorityName *string `json:"taxAuthorityName,omitempty" xml:"taxAuthorityName,omitempty"`
	// 主管税务机关代码
	TaxAuthorityNo *string `json:"taxAuthorityNo,omitempty" xml:"taxAuthorityNo,omitempty"`
	// 税率
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// 吨位
	Tonnage *string `json:"tonnage,omitempty" xml:"tonnage,omitempty"`
	// 车架号码
	VehicleNo *string `json:"vehicleNo,omitempty" xml:"vehicleNo,omitempty"`
	// 车辆类型
	VehicleType *string `json:"vehicleType,omitempty" xml:"vehicleType,omitempty"`
}

func (s UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetBrand(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.Brand = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetCertificateNo(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.CertificateNo = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetEngineNo(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.EngineNo = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetIdCardNo(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.IdCardNo = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetImportCertificateNo(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.ImportCertificateNo = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetInspectionListNo(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.InspectionListNo = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetMaxPassengers(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.MaxPassengers = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetOriginPlace(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.OriginPlace = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetPaymentVoucherNo(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.PaymentVoucherNo = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetTaxAuthorityName(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.TaxAuthorityName = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetTaxAuthorityNo(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.TaxAuthorityNo = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetTaxRate(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.TaxRate = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetTonnage(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.Tonnage = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetVehicleNo(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.VehicleNo = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList) SetVehicleType(v string) *UpdateInvoiceVerifyStatusRequestGeneralInvoiceVOListVehicleSaleDetailVOList {
	s.VehicleType = &v
	return s
}

type UpdateInvoiceVerifyStatusRequestInvoiceKeyVOList struct {
	// 发票编码
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	// 发票号码
	InvoiceNo *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
}

func (s UpdateInvoiceVerifyStatusRequestInvoiceKeyVOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceVerifyStatusRequestInvoiceKeyVOList) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceVerifyStatusRequestInvoiceKeyVOList) SetInvoiceCode(v string) *UpdateInvoiceVerifyStatusRequestInvoiceKeyVOList {
	s.InvoiceCode = &v
	return s
}

func (s *UpdateInvoiceVerifyStatusRequestInvoiceKeyVOList) SetInvoiceNo(v string) *UpdateInvoiceVerifyStatusRequestInvoiceKeyVOList {
	s.InvoiceNo = &v
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInvoiceVerifyStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateInvoiceVerifyStatusResponse) SetBody(v *UpdateInvoiceVerifyStatusResponseBody) *UpdateInvoiceVerifyStatusResponse {
	s.Body = v
	return s
}

type UpdateInvoiceVoucherStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateInvoiceVoucherStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceVoucherStatusHeaders) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceVoucherStatusHeaders) SetCommonHeaders(v map[string]*string) *UpdateInvoiceVoucherStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateInvoiceVoucherStatusHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateInvoiceVoucherStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateInvoiceVoucherStatusRequest struct {
	// 操作类型
	ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// 发票编码
	InvoiceCode *string `json:"invoiceCode,omitempty" xml:"invoiceCode,omitempty"`
	// 发票号码
	InvoiceNo *string `json:"invoiceNo,omitempty" xml:"invoiceNo,omitempty"`
	// 操作员
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 凭证id
	VoucherId *string `json:"voucherId,omitempty" xml:"voucherId,omitempty"`
}

func (s UpdateInvoiceVoucherStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceVoucherStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceVoucherStatusRequest) SetActionType(v string) *UpdateInvoiceVoucherStatusRequest {
	s.ActionType = &v
	return s
}

func (s *UpdateInvoiceVoucherStatusRequest) SetInvoiceCode(v string) *UpdateInvoiceVoucherStatusRequest {
	s.InvoiceCode = &v
	return s
}

func (s *UpdateInvoiceVoucherStatusRequest) SetInvoiceNo(v string) *UpdateInvoiceVoucherStatusRequest {
	s.InvoiceNo = &v
	return s
}

func (s *UpdateInvoiceVoucherStatusRequest) SetOperator(v string) *UpdateInvoiceVoucherStatusRequest {
	s.Operator = &v
	return s
}

func (s *UpdateInvoiceVoucherStatusRequest) SetVoucherId(v string) *UpdateInvoiceVoucherStatusRequest {
	s.VoucherId = &v
	return s
}

type UpdateInvoiceVoucherStatusResponseBody struct {
	// 业务返回结果
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// 系统调用结果
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateInvoiceVoucherStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceVoucherStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceVoucherStatusResponseBody) SetResult(v bool) *UpdateInvoiceVoucherStatusResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateInvoiceVoucherStatusResponseBody) SetSuccess(v bool) *UpdateInvoiceVoucherStatusResponseBody {
	s.Success = &v
	return s
}

type UpdateInvoiceVoucherStatusResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInvoiceVoucherStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInvoiceVoucherStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInvoiceVoucherStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateInvoiceVoucherStatusResponse) SetHeaders(v map[string]*string) *UpdateInvoiceVoucherStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateInvoiceVoucherStatusResponse) SetBody(v *UpdateInvoiceVoucherStatusResponseBody) *UpdateInvoiceVoucherStatusResponse {
	s.Body = v
	return s
}

type UpdateReceiptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateReceiptHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateReceiptHeaders) GoString() string {
	return s.String()
}

func (s *UpdateReceiptHeaders) SetCommonHeaders(v map[string]*string) *UpdateReceiptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateReceiptHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateReceiptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateReceiptRequest struct {
	// 单据列表 ，最长10
	Receipts []*UpdateReceiptRequestReceipts `json:"receipts,omitempty" xml:"receipts,omitempty" type:"Repeated"`
}

func (s UpdateReceiptRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateReceiptRequest) GoString() string {
	return s.String()
}

func (s *UpdateReceiptRequest) SetReceipts(v []*UpdateReceiptRequestReceipts) *UpdateReceiptRequest {
	s.Receipts = v
	return s
}

type UpdateReceiptRequestReceipts struct {
	// 总金额，传空代表不修改
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 关联收支类别，传空代表不修改
	CategoryCode *string `json:"categoryCode,omitempty" xml:"categoryCode,omitempty"`
	// 单据唯一编号
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 关联客户code，传空代表不修改
	CustomerCode *string `json:"customerCode,omitempty" xml:"customerCode,omitempty"`
	// 关联企业账户code，传空代表不修改
	EnterpriseAcountCode *string `json:"enterpriseAcountCode,omitempty" xml:"enterpriseAcountCode,omitempty"`
	// 业务发生时间，传空代表不修改
	OccurDate *int64 `json:"occurDate,omitempty" xml:"occurDate,omitempty"`
	// 负责人工号，传空代表不修改
	PrincipalId *string `json:"principalId,omitempty" xml:"principalId,omitempty"`
	// 关联项目code，传空代表不修改
	ProjectCode *string `json:"projectCode,omitempty" xml:"projectCode,omitempty"`
	// 单据类型：1付款单，2收款单
	ReceiptType *int64 `json:"receiptType,omitempty" xml:"receiptType,omitempty"`
	// 备注，传空代表不修改
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 关联供应商code，传空代表不修改
	SupplierCode *string `json:"supplierCode,omitempty" xml:"supplierCode,omitempty"`
	// 单据标题，传空代表不修改
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 单据更新时间
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// 修改者工号
	UpdateUserId *string `json:"updateUserId,omitempty" xml:"updateUserId,omitempty"`
}

func (s UpdateReceiptRequestReceipts) String() string {
	return tea.Prettify(s)
}

func (s UpdateReceiptRequestReceipts) GoString() string {
	return s.String()
}

func (s *UpdateReceiptRequestReceipts) SetAmount(v string) *UpdateReceiptRequestReceipts {
	s.Amount = &v
	return s
}

func (s *UpdateReceiptRequestReceipts) SetCategoryCode(v string) *UpdateReceiptRequestReceipts {
	s.CategoryCode = &v
	return s
}

func (s *UpdateReceiptRequestReceipts) SetCode(v string) *UpdateReceiptRequestReceipts {
	s.Code = &v
	return s
}

func (s *UpdateReceiptRequestReceipts) SetCustomerCode(v string) *UpdateReceiptRequestReceipts {
	s.CustomerCode = &v
	return s
}

func (s *UpdateReceiptRequestReceipts) SetEnterpriseAcountCode(v string) *UpdateReceiptRequestReceipts {
	s.EnterpriseAcountCode = &v
	return s
}

func (s *UpdateReceiptRequestReceipts) SetOccurDate(v int64) *UpdateReceiptRequestReceipts {
	s.OccurDate = &v
	return s
}

func (s *UpdateReceiptRequestReceipts) SetPrincipalId(v string) *UpdateReceiptRequestReceipts {
	s.PrincipalId = &v
	return s
}

func (s *UpdateReceiptRequestReceipts) SetProjectCode(v string) *UpdateReceiptRequestReceipts {
	s.ProjectCode = &v
	return s
}

func (s *UpdateReceiptRequestReceipts) SetReceiptType(v int64) *UpdateReceiptRequestReceipts {
	s.ReceiptType = &v
	return s
}

func (s *UpdateReceiptRequestReceipts) SetRemark(v string) *UpdateReceiptRequestReceipts {
	s.Remark = &v
	return s
}

func (s *UpdateReceiptRequestReceipts) SetSupplierCode(v string) *UpdateReceiptRequestReceipts {
	s.SupplierCode = &v
	return s
}

func (s *UpdateReceiptRequestReceipts) SetTitle(v string) *UpdateReceiptRequestReceipts {
	s.Title = &v
	return s
}

func (s *UpdateReceiptRequestReceipts) SetUpdateTime(v int64) *UpdateReceiptRequestReceipts {
	s.UpdateTime = &v
	return s
}

func (s *UpdateReceiptRequestReceipts) SetUpdateUserId(v string) *UpdateReceiptRequestReceipts {
	s.UpdateUserId = &v
	return s
}

type UpdateReceiptResponseBody struct {
	// 结果列表
	Results []*UpdateReceiptResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s UpdateReceiptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateReceiptResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateReceiptResponseBody) SetResults(v []*UpdateReceiptResponseBodyResults) *UpdateReceiptResponseBody {
	s.Results = v
	return s
}

type UpdateReceiptResponseBodyResults struct {
	// 数据唯一编号
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateReceiptResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s UpdateReceiptResponseBodyResults) GoString() string {
	return s.String()
}

func (s *UpdateReceiptResponseBodyResults) SetCode(v string) *UpdateReceiptResponseBodyResults {
	s.Code = &v
	return s
}

func (s *UpdateReceiptResponseBodyResults) SetErrorCode(v string) *UpdateReceiptResponseBodyResults {
	s.ErrorCode = &v
	return s
}

func (s *UpdateReceiptResponseBodyResults) SetErrorMsg(v string) *UpdateReceiptResponseBodyResults {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateReceiptResponseBodyResults) SetSuccess(v bool) *UpdateReceiptResponseBodyResults {
	s.Success = &v
	return s
}

type UpdateReceiptResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateReceiptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateReceiptResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateReceiptResponse) GoString() string {
	return s.String()
}

func (s *UpdateReceiptResponse) SetHeaders(v map[string]*string) *UpdateReceiptResponse {
	s.Headers = v
	return s
}

func (s *UpdateReceiptResponse) SetBody(v *UpdateReceiptResponseBody) *UpdateReceiptResponse {
	s.Body = v
	return s
}

type UpdateReceiptVoucherStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateReceiptVoucherStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateReceiptVoucherStatusHeaders) GoString() string {
	return s.String()
}

func (s *UpdateReceiptVoucherStatusHeaders) SetCommonHeaders(v map[string]*string) *UpdateReceiptVoucherStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateReceiptVoucherStatusHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateReceiptVoucherStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateReceiptVoucherStatusRequest struct {
	// 账期
	AccountPeriod *string `json:"accountPeriod,omitempty" xml:"accountPeriod,omitempty"`
	// 操作类型 add 添加凭证关系、delete 删除凭证关系
	ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// 操作人工号
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// 审批单据ID
	ReceiptId *string `json:"receiptId,omitempty" xml:"receiptId,omitempty"`
	// 凭证CODE
	VoucherCode *string `json:"voucherCode,omitempty" xml:"voucherCode,omitempty"`
	// 凭证ID
	VoucherId *string `json:"voucherId,omitempty" xml:"voucherId,omitempty"`
}

func (s UpdateReceiptVoucherStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateReceiptVoucherStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateReceiptVoucherStatusRequest) SetAccountPeriod(v string) *UpdateReceiptVoucherStatusRequest {
	s.AccountPeriod = &v
	return s
}

func (s *UpdateReceiptVoucherStatusRequest) SetActionType(v string) *UpdateReceiptVoucherStatusRequest {
	s.ActionType = &v
	return s
}

func (s *UpdateReceiptVoucherStatusRequest) SetOperatorId(v string) *UpdateReceiptVoucherStatusRequest {
	s.OperatorId = &v
	return s
}

func (s *UpdateReceiptVoucherStatusRequest) SetReceiptId(v string) *UpdateReceiptVoucherStatusRequest {
	s.ReceiptId = &v
	return s
}

func (s *UpdateReceiptVoucherStatusRequest) SetVoucherCode(v string) *UpdateReceiptVoucherStatusRequest {
	s.VoucherCode = &v
	return s
}

func (s *UpdateReceiptVoucherStatusRequest) SetVoucherId(v string) *UpdateReceiptVoucherStatusRequest {
	s.VoucherId = &v
	return s
}

type UpdateReceiptVoucherStatusResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateReceiptVoucherStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateReceiptVoucherStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateReceiptVoucherStatusResponseBody) SetResult(v bool) *UpdateReceiptVoucherStatusResponseBody {
	s.Result = &v
	return s
}

type UpdateReceiptVoucherStatusResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateReceiptVoucherStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateReceiptVoucherStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateReceiptVoucherStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateReceiptVoucherStatusResponse) SetHeaders(v map[string]*string) *UpdateReceiptVoucherStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateReceiptVoucherStatusResponse) SetBody(v *UpdateReceiptVoucherStatusResponseBody) *UpdateReceiptVoucherStatusResponse {
	s.Body = v
	return s
}

type RoleMemberMapValue struct {
	// 角色唯一标识
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	// 成员信息列表
	MemberList []*RoleMemberMapValueMemberList `json:"memberList,omitempty" xml:"memberList,omitempty" type:"Repeated"`
}

func (s RoleMemberMapValue) String() string {
	return tea.Prettify(s)
}

func (s RoleMemberMapValue) GoString() string {
	return s.String()
}

func (s *RoleMemberMapValue) SetRoleCode(v string) *RoleMemberMapValue {
	s.RoleCode = &v
	return s
}

func (s *RoleMemberMapValue) SetMemberList(v []*RoleMemberMapValueMemberList) *RoleMemberMapValue {
	s.MemberList = v
	return s
}

type RoleMemberMapValueMemberList struct {
	// 用户ID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 昵称
	Nick *string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 头像URL
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
}

func (s RoleMemberMapValueMemberList) String() string {
	return tea.Prettify(s)
}

func (s RoleMemberMapValueMemberList) GoString() string {
	return s.String()
}

func (s *RoleMemberMapValueMemberList) SetUserId(v string) *RoleMemberMapValueMemberList {
	s.UserId = &v
	return s
}

func (s *RoleMemberMapValueMemberList) SetNick(v string) *RoleMemberMapValueMemberList {
	s.Nick = &v
	return s
}

func (s *RoleMemberMapValueMemberList) SetAvatarUrl(v string) *RoleMemberMapValueMemberList {
	s.AvatarUrl = &v
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

func (client *Client) BatchAddInvoice(request *BatchAddInvoiceRequest) (_result *BatchAddInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchAddInvoiceHeaders{}
	_result = &BatchAddInvoiceResponse{}
	_body, _err := client.BatchAddInvoiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchAddInvoiceWithOptions(request *BatchAddInvoiceRequest, headers *BatchAddInvoiceHeaders, runtime *util.RuntimeOptions) (_result *BatchAddInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GeneralInvoiceVOList)) {
		body["generalInvoiceVOList"] = request.GeneralInvoiceVOList
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
	_result = &BatchAddInvoiceResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchAddInvoice"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/bizfinance/invoices/batch"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CheckVoucherStatusWithOptions(request *CheckVoucherStatusRequest, headers *CheckVoucherStatusHeaders, runtime *util.RuntimeOptions) (_result *CheckVoucherStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
	_result = &CheckVoucherStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("CheckVoucherStatus"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/bizfinance/invoices/checkVoucherStatus/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCustomer(request *CreateCustomerRequest) (_result *CreateCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateCustomerHeaders{}
	_result = &CreateCustomerResponse{}
	_body, _err := client.CreateCustomerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCustomerWithOptions(request *CreateCustomerRequest, headers *CreateCustomerHeaders, runtime *util.RuntimeOptions) (_result *CreateCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Creator)) {
		body["creator"] = request.Creator
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PurchaserAccount)) {
		body["purchaserAccount"] = request.PurchaserAccount
	}

	if !tea.BoolValue(util.IsUnset(request.PurchaserAddress)) {
		body["purchaserAddress"] = request.PurchaserAddress
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
	_result = &CreateCustomerResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateCustomer"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/bizfinance/auxiliaries/customers"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateReceipt(request *CreateReceiptRequest) (_result *CreateReceiptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateReceiptHeaders{}
	_result = &CreateReceiptResponse{}
	_body, _err := client.CreateReceiptWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateReceiptWithOptions(request *CreateReceiptRequest, headers *CreateReceiptHeaders, runtime *util.RuntimeOptions) (_result *CreateReceiptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Receipts)) {
		body["receipts"] = request.Receipts
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
	_result = &CreateReceiptResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateReceipt"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/bizfinance/receipts"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteReceipt(request *DeleteReceiptRequest) (_result *DeleteReceiptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteReceiptHeaders{}
	_result = &DeleteReceiptResponse{}
	_body, _err := client.DeleteReceiptWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteReceiptWithOptions(request *DeleteReceiptRequest, headers *DeleteReceiptHeaders, runtime *util.RuntimeOptions) (_result *DeleteReceiptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Receipts)) {
		body["receipts"] = request.Receipts
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
	_result = &DeleteReceiptResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteReceipt"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/bizfinance/receipts/remove"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBookkeepingUserList() (_result *GetBookkeepingUserListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetBookkeepingUserListHeaders{}
	_result = &GetBookkeepingUserListResponse{}
	_body, _err := client.GetBookkeepingUserListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBookkeepingUserListWithOptions(headers *GetBookkeepingUserListHeaders, runtime *util.RuntimeOptions) (_result *GetBookkeepingUserListResponse, _err error) {
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
	_result = &GetBookkeepingUserListResponse{}
	_body, _err := client.DoROARequest(tea.String("GetBookkeepingUserList"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/bizfinance/bookkeeping/users"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetCategoryResponse{}
	_body, _err := client.DoROARequest(tea.String("GetCategory"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/bizfinance/categories/get"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCustomer(request *GetCustomerRequest) (_result *GetCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCustomerHeaders{}
	_result = &GetCustomerResponse{}
	_body, _err := client.GetCustomerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCustomerWithOptions(request *GetCustomerRequest, headers *GetCustomerHeaders, runtime *util.RuntimeOptions) (_result *GetCustomerResponse, _err error) {
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
	_result = &GetCustomerResponse{}
	_body, _err := client.DoROARequest(tea.String("GetCustomer"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/bizfinance/customers/details"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetFinanceAccountResponse{}
	_body, _err := client.DoROARequest(tea.String("GetFinanceAccount"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/bizfinance/financeAccounts/get"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInvoiceByPage(request *GetInvoiceByPageRequest) (_result *GetInvoiceByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetInvoiceByPageHeaders{}
	_result = &GetInvoiceByPageResponse{}
	_body, _err := client.GetInvoiceByPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInvoiceByPageWithOptions(tmpReq *GetInvoiceByPageRequest, headers *GetInvoiceByPageHeaders, runtime *util.RuntimeOptions) (_result *GetInvoiceByPageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetInvoiceByPageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Request))) {
		request.RequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Request), tea.String("request"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RequestShrink)) {
		query["request"] = request.RequestShrink
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
	_result = &GetInvoiceByPageResponse{}
	_body, _err := client.DoROARequest(tea.String("GetInvoiceByPage"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/bizfinance/invoices"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIsNewVersion() (_result *GetIsNewVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetIsNewVersionHeaders{}
	_result = &GetIsNewVersionResponse{}
	_body, _err := client.GetIsNewVersionWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIsNewVersionWithOptions(headers *GetIsNewVersionHeaders, runtime *util.RuntimeOptions) (_result *GetIsNewVersionResponse, _err error) {
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
	_result = &GetIsNewVersionResponse{}
	_body, _err := client.DoROARequest(tea.String("GetIsNewVersion"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/bizfinance/accounts/uses"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetProjectResponse{}
	_body, _err := client.DoROARequest(tea.String("GetProject"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/bizfinance/projects/get"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetReceiptResponse{}
	_body, _err := client.DoROARequest(tea.String("GetReceipt"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/bizfinance/receipts/details"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetSupplierResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSupplier"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/bizfinance/suppliers/details"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryCategoryByPageResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryCategoryByPage"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/bizfinance/categories/list"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryCustomerByPageResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryCustomerByPage"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/bizfinance/customers"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCustomerInfo(request *QueryCustomerInfoRequest) (_result *QueryCustomerInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCustomerInfoHeaders{}
	_result = &QueryCustomerInfoResponse{}
	_body, _err := client.QueryCustomerInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCustomerInfoWithOptions(request *QueryCustomerInfoRequest, headers *QueryCustomerInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryCustomerInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
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
	_result = &QueryCustomerInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryCustomerInfo"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/bizfinance/auxiliaries/customers"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryEnterpriseAccountByPageResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryEnterpriseAccountByPage"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/bizfinance/financeAccounts/list"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFinanceCompanyInfo() (_result *QueryFinanceCompanyInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryFinanceCompanyInfoHeaders{}
	_result = &QueryFinanceCompanyInfoResponse{}
	_body, _err := client.QueryFinanceCompanyInfoWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFinanceCompanyInfoWithOptions(headers *QueryFinanceCompanyInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryFinanceCompanyInfoResponse, _err error) {
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
	_result = &QueryFinanceCompanyInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryFinanceCompanyInfo"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/bizfinance/companies"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryInvoiceRelationCount() (_result *QueryInvoiceRelationCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryInvoiceRelationCountHeaders{}
	_result = &QueryInvoiceRelationCountResponse{}
	_body, _err := client.QueryInvoiceRelationCountWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryInvoiceRelationCountWithOptions(headers *QueryInvoiceRelationCountHeaders, runtime *util.RuntimeOptions) (_result *QueryInvoiceRelationCountResponse, _err error) {
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
	_result = &QueryInvoiceRelationCountResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryInvoiceRelationCount"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/bizfinance/invoices/relationReceipts/counts"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPermissionByUserId(request *QueryPermissionByUserIdRequest) (_result *QueryPermissionByUserIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryPermissionByUserIdHeaders{}
	_result = &QueryPermissionByUserIdResponse{}
	_body, _err := client.QueryPermissionByUserIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPermissionByUserIdWithOptions(request *QueryPermissionByUserIdRequest, headers *QueryPermissionByUserIdHeaders, runtime *util.RuntimeOptions) (_result *QueryPermissionByUserIdResponse, _err error) {
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
	_result = &QueryPermissionByUserIdResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryPermissionByUserId"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/bizfinance/permissions"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPermissionRoleMember(request *QueryPermissionRoleMemberRequest) (_result *QueryPermissionRoleMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryPermissionRoleMemberHeaders{}
	_result = &QueryPermissionRoleMemberResponse{}
	_body, _err := client.QueryPermissionRoleMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPermissionRoleMemberWithOptions(request *QueryPermissionRoleMemberRequest, headers *QueryPermissionRoleMemberHeaders, runtime *util.RuntimeOptions) (_result *QueryPermissionRoleMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleCodeList)) {
		body["roleCodeList"] = request.RoleCodeList
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
	_result = &QueryPermissionRoleMemberResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryPermissionRoleMember"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/bizfinance/roles/members/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryProjectByPageResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryProjectByPage"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/bizfinance/projects/list"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) QueryReceiptForInvoiceWithOptions(request *QueryReceiptForInvoiceRequest, headers *QueryReceiptForInvoiceHeaders, runtime *util.RuntimeOptions) (_result *QueryReceiptForInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyStatusList)) {
		body["applyStatusList"] = request.ApplyStatusList
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
	_result = &QueryReceiptForInvoiceResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryReceiptForInvoice"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/bizfinance/invoices/receipts/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryReceiptsBaseInfo(request *QueryReceiptsBaseInfoRequest) (_result *QueryReceiptsBaseInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryReceiptsBaseInfoHeaders{}
	_result = &QueryReceiptsBaseInfoResponse{}
	_body, _err := client.QueryReceiptsBaseInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryReceiptsBaseInfoWithOptions(request *QueryReceiptsBaseInfoRequest, headers *QueryReceiptsBaseInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryReceiptsBaseInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.VoucherStatus)) {
		query["voucherStatus"] = request.VoucherStatus
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
	_result = &QueryReceiptsBaseInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryReceiptsBaseInfo"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/bizfinance/receipts/dataInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryReceiptsByPage(request *QueryReceiptsByPageRequest) (_result *QueryReceiptsByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryReceiptsByPageHeaders{}
	_result = &QueryReceiptsByPageResponse{}
	_body, _err := client.QueryReceiptsByPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryReceiptsByPageWithOptions(request *QueryReceiptsByPageRequest, headers *QueryReceiptsByPageHeaders, runtime *util.RuntimeOptions) (_result *QueryReceiptsByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		query["modelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TimeFilterField)) {
		query["timeFilterField"] = request.TimeFilterField
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
	_result = &QueryReceiptsByPageResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryReceiptsByPage"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/bizfinance/receipts"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QuerySupplierByPageResponse{}
	_body, _err := client.DoROARequest(tea.String("QuerySupplierByPage"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/bizfinance/suppliers"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateApplyReceiptAndInvoiceRelated(request *UpdateApplyReceiptAndInvoiceRelatedRequest) (_result *UpdateApplyReceiptAndInvoiceRelatedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateApplyReceiptAndInvoiceRelatedHeaders{}
	_result = &UpdateApplyReceiptAndInvoiceRelatedResponse{}
	_body, _err := client.UpdateApplyReceiptAndInvoiceRelatedWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateApplyReceiptAndInvoiceRelatedWithOptions(request *UpdateApplyReceiptAndInvoiceRelatedRequest, headers *UpdateApplyReceiptAndInvoiceRelatedHeaders, runtime *util.RuntimeOptions) (_result *UpdateApplyReceiptAndInvoiceRelatedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GeneralInvoiceVOList)) {
		body["generalInvoiceVOList"] = request.GeneralInvoiceVOList
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
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
	_result = &UpdateApplyReceiptAndInvoiceRelatedResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateApplyReceiptAndInvoiceRelated"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/bizfinance/invoices/applyReceipts/relate"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFinanceCompanyInfo(request *UpdateFinanceCompanyInfoRequest) (_result *UpdateFinanceCompanyInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateFinanceCompanyInfoHeaders{}
	_result = &UpdateFinanceCompanyInfoResponse{}
	_body, _err := client.UpdateFinanceCompanyInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFinanceCompanyInfoWithOptions(request *UpdateFinanceCompanyInfoRequest, headers *UpdateFinanceCompanyInfoHeaders, runtime *util.RuntimeOptions) (_result *UpdateFinanceCompanyInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompanyName)) {
		query["companyName"] = request.CompanyName
	}

	if !tea.BoolValue(util.IsUnset(request.TaxNature)) {
		query["taxNature"] = request.TaxNature
	}

	if !tea.BoolValue(util.IsUnset(request.TaxNo)) {
		query["taxNo"] = request.TaxNo
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
	_result = &UpdateFinanceCompanyInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateFinanceCompanyInfo"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/bizfinance/companies"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInvoiceAbandonStatus(request *UpdateInvoiceAbandonStatusRequest) (_result *UpdateInvoiceAbandonStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateInvoiceAbandonStatusHeaders{}
	_result = &UpdateInvoiceAbandonStatusResponse{}
	_body, _err := client.UpdateInvoiceAbandonStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInvoiceAbandonStatusWithOptions(request *UpdateInvoiceAbandonStatusRequest, headers *UpdateInvoiceAbandonStatusHeaders, runtime *util.RuntimeOptions) (_result *UpdateInvoiceAbandonStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.BlueGeneralInvoiceVO))) {
		body["blueGeneralInvoiceVO"] = request.BlueGeneralInvoiceVO
	}

	if !tea.BoolValue(util.IsUnset(request.BlueInvoiceCode)) {
		body["blueInvoiceCode"] = request.BlueInvoiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.BlueInvoiceNo)) {
		body["blueInvoiceNo"] = request.BlueInvoiceNo
	}

	if !tea.BoolValue(util.IsUnset(request.BlueInvoiceStatus)) {
		body["blueInvoiceStatus"] = request.BlueInvoiceStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.RedGeneralInvoiceVO))) {
		body["redGeneralInvoiceVO"] = request.RedGeneralInvoiceVO
	}

	if !tea.BoolValue(util.IsUnset(request.RedInvoiceCode)) {
		body["redInvoiceCode"] = request.RedInvoiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.RedInvoiceNo)) {
		body["redInvoiceNo"] = request.RedInvoiceNo
	}

	if !tea.BoolValue(util.IsUnset(request.RedInvoiceStatus)) {
		body["redInvoiceStatus"] = request.RedInvoiceStatus
	}

	if !tea.BoolValue(util.IsUnset(request.TargetInvoice)) {
		body["targetInvoice"] = request.TargetInvoice
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
	_result = &UpdateInvoiceAbandonStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateInvoiceAbandonStatus"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/bizfinance/invoices/abandonStatus"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInvoiceAccountPeriod(request *UpdateInvoiceAccountPeriodRequest) (_result *UpdateInvoiceAccountPeriodResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateInvoiceAccountPeriodHeaders{}
	_result = &UpdateInvoiceAccountPeriodResponse{}
	_body, _err := client.UpdateInvoiceAccountPeriodWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInvoiceAccountPeriodWithOptions(request *UpdateInvoiceAccountPeriodRequest, headers *UpdateInvoiceAccountPeriodHeaders, runtime *util.RuntimeOptions) (_result *UpdateInvoiceAccountPeriodResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountPeriod)) {
		body["accountPeriod"] = request.AccountPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.GeneralInvoiceVOList)) {
		body["generalInvoiceVOList"] = request.GeneralInvoiceVOList
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
	_result = &UpdateInvoiceAccountPeriodResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateInvoiceAccountPeriod"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/bizfinance/invoices/accountPeriods"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInvoiceAndReceiptRelated(request *UpdateInvoiceAndReceiptRelatedRequest) (_result *UpdateInvoiceAndReceiptRelatedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateInvoiceAndReceiptRelatedHeaders{}
	_result = &UpdateInvoiceAndReceiptRelatedResponse{}
	_body, _err := client.UpdateInvoiceAndReceiptRelatedWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInvoiceAndReceiptRelatedWithOptions(request *UpdateInvoiceAndReceiptRelatedRequest, headers *UpdateInvoiceAndReceiptRelatedHeaders, runtime *util.RuntimeOptions) (_result *UpdateInvoiceAndReceiptRelatedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.GeneralInvoiceVO))) {
		body["generalInvoiceVO"] = request.GeneralInvoiceVO
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceCode)) {
		body["invoiceCode"] = request.InvoiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceNo)) {
		body["invoiceNo"] = request.InvoiceNo
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiptCode)) {
		body["receiptCode"] = request.ReceiptCode
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
	_result = &UpdateInvoiceAndReceiptRelatedResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateInvoiceAndReceiptRelated"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/bizfinance/invoices/approvalReceipts"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInvoiceIgnoreStatus(request *UpdateInvoiceIgnoreStatusRequest) (_result *UpdateInvoiceIgnoreStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateInvoiceIgnoreStatusHeaders{}
	_result = &UpdateInvoiceIgnoreStatusResponse{}
	_body, _err := client.UpdateInvoiceIgnoreStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInvoiceIgnoreStatusWithOptions(request *UpdateInvoiceIgnoreStatusRequest, headers *UpdateInvoiceIgnoreStatusHeaders, runtime *util.RuntimeOptions) (_result *UpdateInvoiceIgnoreStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		query["operator"] = request.Operator
	}

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
	_result = &UpdateInvoiceIgnoreStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateInvoiceIgnoreStatus"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/bizfinance/invoices/ignoreStatus"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) UpdateInvoiceVerifyStatusWithOptions(request *UpdateInvoiceVerifyStatusRequest, headers *UpdateInvoiceVerifyStatusHeaders, runtime *util.RuntimeOptions) (_result *UpdateInvoiceVerifyStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeductStatus)) {
		body["deductStatus"] = request.DeductStatus
	}

	if !tea.BoolValue(util.IsUnset(request.GeneralInvoiceVOList)) {
		body["generalInvoiceVOList"] = request.GeneralInvoiceVOList
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceKeyVOList)) {
		body["invoiceKeyVOList"] = request.InvoiceKeyVOList
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["operator"] = request.Operator
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
	_result = &UpdateInvoiceVerifyStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateInvoiceVerifyStatus"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/bizfinance/invoices/verifyStatus"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInvoiceVoucherStatus(request *UpdateInvoiceVoucherStatusRequest) (_result *UpdateInvoiceVoucherStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateInvoiceVoucherStatusHeaders{}
	_result = &UpdateInvoiceVoucherStatusResponse{}
	_body, _err := client.UpdateInvoiceVoucherStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInvoiceVoucherStatusWithOptions(request *UpdateInvoiceVoucherStatusRequest, headers *UpdateInvoiceVoucherStatusHeaders, runtime *util.RuntimeOptions) (_result *UpdateInvoiceVoucherStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionType)) {
		body["actionType"] = request.ActionType
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceCode)) {
		body["invoiceCode"] = request.InvoiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceNo)) {
		body["invoiceNo"] = request.InvoiceNo
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.VoucherId)) {
		body["voucherId"] = request.VoucherId
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
	_result = &UpdateInvoiceVoucherStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateInvoiceVoucherStatus"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/bizfinance/invoices/vouchers/states"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateReceipt(request *UpdateReceiptRequest) (_result *UpdateReceiptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateReceiptHeaders{}
	_result = &UpdateReceiptResponse{}
	_body, _err := client.UpdateReceiptWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateReceiptWithOptions(request *UpdateReceiptRequest, headers *UpdateReceiptHeaders, runtime *util.RuntimeOptions) (_result *UpdateReceiptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Receipts)) {
		body["receipts"] = request.Receipts
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
	_result = &UpdateReceiptResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateReceipt"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/bizfinance/receipts"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateReceiptVoucherStatus(request *UpdateReceiptVoucherStatusRequest) (_result *UpdateReceiptVoucherStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateReceiptVoucherStatusHeaders{}
	_result = &UpdateReceiptVoucherStatusResponse{}
	_body, _err := client.UpdateReceiptVoucherStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateReceiptVoucherStatusWithOptions(request *UpdateReceiptVoucherStatusRequest, headers *UpdateReceiptVoucherStatusHeaders, runtime *util.RuntimeOptions) (_result *UpdateReceiptVoucherStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountPeriod)) {
		body["accountPeriod"] = request.AccountPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.ActionType)) {
		body["actionType"] = request.ActionType
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiptId)) {
		body["receiptId"] = request.ReceiptId
	}

	if !tea.BoolValue(util.IsUnset(request.VoucherCode)) {
		body["voucherCode"] = request.VoucherCode
	}

	if !tea.BoolValue(util.IsUnset(request.VoucherId)) {
		body["voucherId"] = request.VoucherId
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
	_result = &UpdateReceiptVoucherStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateReceiptVoucherStatus"), tea.String("bizfinance_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/bizfinance/vouchers/recepits"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
