// This file is auto-generated, don't edit it. Thanks.
package hrm_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ResultValue struct {
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ResultValue) String() string {
	return tea.Prettify(s)
}

func (s ResultValue) GoString() string {
	return s.String()
}

func (s *ResultValue) SetStatus(v string) *ResultValue {
	s.Status = &v
	return s
}

type AddCustomRosterFieldHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddCustomRosterFieldHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddCustomRosterFieldHeaders) GoString() string {
	return s.String()
}

func (s *AddCustomRosterFieldHeaders) SetCommonHeaders(v map[string]*string) *AddCustomRosterFieldHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddCustomRosterFieldHeaders) SetXAcsDingtalkAccessToken(v string) *AddCustomRosterFieldHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddCustomRosterFieldRequest struct {
	// This parameter is required.
	EditFromEmployeeFlag *bool `json:"editFromEmployeeFlag,omitempty" xml:"editFromEmployeeFlag,omitempty"`
	// This parameter is required.
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// This parameter is required.
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	// This parameter is required.
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// This parameter is required.
	HiddenFromEmployeeFlag *bool   `json:"hiddenFromEmployeeFlag,omitempty" xml:"hiddenFromEmployeeFlag,omitempty"`
	Hint                   *string `json:"hint,omitempty" xml:"hint,omitempty"`
	NoWatermark            *bool   `json:"noWatermark,omitempty" xml:"noWatermark,omitempty"`
	NumberDecimalPlace     *int32  `json:"numberDecimalPlace,omitempty" xml:"numberDecimalPlace,omitempty"`
	NumberFormatType       *string `json:"numberFormatType,omitempty" xml:"numberFormatType,omitempty"`
	NumberValueType        *string `json:"numberValueType,omitempty" xml:"numberValueType,omitempty"`
	OptionText             *string `json:"optionText,omitempty" xml:"optionText,omitempty"`
	// This parameter is required.
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// This parameter is required.
	VisibleByEmp *bool `json:"visibleByEmp,omitempty" xml:"visibleByEmp,omitempty"`
}

func (s AddCustomRosterFieldRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCustomRosterFieldRequest) GoString() string {
	return s.String()
}

func (s *AddCustomRosterFieldRequest) SetEditFromEmployeeFlag(v bool) *AddCustomRosterFieldRequest {
	s.EditFromEmployeeFlag = &v
	return s
}

func (s *AddCustomRosterFieldRequest) SetFieldName(v string) *AddCustomRosterFieldRequest {
	s.FieldName = &v
	return s
}

func (s *AddCustomRosterFieldRequest) SetFieldType(v string) *AddCustomRosterFieldRequest {
	s.FieldType = &v
	return s
}

func (s *AddCustomRosterFieldRequest) SetGroupId(v string) *AddCustomRosterFieldRequest {
	s.GroupId = &v
	return s
}

func (s *AddCustomRosterFieldRequest) SetHiddenFromEmployeeFlag(v bool) *AddCustomRosterFieldRequest {
	s.HiddenFromEmployeeFlag = &v
	return s
}

func (s *AddCustomRosterFieldRequest) SetHint(v string) *AddCustomRosterFieldRequest {
	s.Hint = &v
	return s
}

func (s *AddCustomRosterFieldRequest) SetNoWatermark(v bool) *AddCustomRosterFieldRequest {
	s.NoWatermark = &v
	return s
}

func (s *AddCustomRosterFieldRequest) SetNumberDecimalPlace(v int32) *AddCustomRosterFieldRequest {
	s.NumberDecimalPlace = &v
	return s
}

func (s *AddCustomRosterFieldRequest) SetNumberFormatType(v string) *AddCustomRosterFieldRequest {
	s.NumberFormatType = &v
	return s
}

func (s *AddCustomRosterFieldRequest) SetNumberValueType(v string) *AddCustomRosterFieldRequest {
	s.NumberValueType = &v
	return s
}

func (s *AddCustomRosterFieldRequest) SetOptionText(v string) *AddCustomRosterFieldRequest {
	s.OptionText = &v
	return s
}

func (s *AddCustomRosterFieldRequest) SetRequired(v bool) *AddCustomRosterFieldRequest {
	s.Required = &v
	return s
}

func (s *AddCustomRosterFieldRequest) SetVisibleByEmp(v bool) *AddCustomRosterFieldRequest {
	s.VisibleByEmp = &v
	return s
}

type AddCustomRosterFieldResponseBody struct {
	DingOpenErrcode *int32                                  `json:"dingOpenErrcode,omitempty" xml:"dingOpenErrcode,omitempty"`
	ErrorMsg        *string                                 `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result          *AddCustomRosterFieldResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success         *bool                                   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddCustomRosterFieldResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCustomRosterFieldResponseBody) GoString() string {
	return s.String()
}

func (s *AddCustomRosterFieldResponseBody) SetDingOpenErrcode(v int32) *AddCustomRosterFieldResponseBody {
	s.DingOpenErrcode = &v
	return s
}

func (s *AddCustomRosterFieldResponseBody) SetErrorMsg(v string) *AddCustomRosterFieldResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *AddCustomRosterFieldResponseBody) SetResult(v *AddCustomRosterFieldResponseBodyResult) *AddCustomRosterFieldResponseBody {
	s.Result = v
	return s
}

func (s *AddCustomRosterFieldResponseBody) SetSuccess(v bool) *AddCustomRosterFieldResponseBody {
	s.Success = &v
	return s
}

type AddCustomRosterFieldResponseBodyResult struct {
	ContactClientFlag      *bool   `json:"contactClientFlag,omitempty" xml:"contactClientFlag,omitempty"`
	ContactFlag            *bool   `json:"contactFlag,omitempty" xml:"contactFlag,omitempty"`
	ContactSource          *int32  `json:"contactSource,omitempty" xml:"contactSource,omitempty"`
	ContactSystemFlag      *bool   `json:"contactSystemFlag,omitempty" xml:"contactSystemFlag,omitempty"`
	Deleted                *bool   `json:"deleted,omitempty" xml:"deleted,omitempty"`
	Derived                *bool   `json:"derived,omitempty" xml:"derived,omitempty"`
	Disabled               *int32  `json:"disabled,omitempty" xml:"disabled,omitempty"`
	EditFromEmployeeFlag   *bool   `json:"editFromEmployeeFlag,omitempty" xml:"editFromEmployeeFlag,omitempty"`
	EditableByHr           *bool   `json:"editableByHr,omitempty" xml:"editableByHr,omitempty"`
	FieldCode              *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	FieldName              *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	FieldTip               *string `json:"fieldTip,omitempty" xml:"fieldTip,omitempty"`
	FieldType              *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	GroupId                *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	HiddenFromEmployeeFlag *bool   `json:"hiddenFromEmployeeFlag,omitempty" xml:"hiddenFromEmployeeFlag,omitempty"`
	Hint                   *string `json:"hint,omitempty" xml:"hint,omitempty"`
	HistoryField           *bool   `json:"historyField,omitempty" xml:"historyField,omitempty"`
	Index                  *int32  `json:"index,omitempty" xml:"index,omitempty"`
	ModifyOptions          *bool   `json:"modifyOptions,omitempty" xml:"modifyOptions,omitempty"`
	NoWatermark            *bool   `json:"noWatermark,omitempty" xml:"noWatermark,omitempty"`
	NumberDecimalPlace     *string `json:"numberDecimalPlace,omitempty" xml:"numberDecimalPlace,omitempty"`
	NumberFormatType       *string `json:"numberFormatType,omitempty" xml:"numberFormatType,omitempty"`
	NumberValueType        *string `json:"numberValueType,omitempty" xml:"numberValueType,omitempty"`
	OptionText             *string `json:"optionText,omitempty" xml:"optionText,omitempty"`
	Required               *bool   `json:"required,omitempty" xml:"required,omitempty"`
	SourceFieldCode        *string `json:"sourceFieldCode,omitempty" xml:"sourceFieldCode,omitempty"`
	SystemFlag             *bool   `json:"systemFlag,omitempty" xml:"systemFlag,omitempty"`
	TextToSelectField      *bool   `json:"textToSelectField,omitempty" xml:"textToSelectField,omitempty"`
	Value                  *string `json:"value,omitempty" xml:"value,omitempty"`
	VisibleByEmp           *bool   `json:"visibleByEmp,omitempty" xml:"visibleByEmp,omitempty"`
}

func (s AddCustomRosterFieldResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddCustomRosterFieldResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddCustomRosterFieldResponseBodyResult) SetContactClientFlag(v bool) *AddCustomRosterFieldResponseBodyResult {
	s.ContactClientFlag = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetContactFlag(v bool) *AddCustomRosterFieldResponseBodyResult {
	s.ContactFlag = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetContactSource(v int32) *AddCustomRosterFieldResponseBodyResult {
	s.ContactSource = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetContactSystemFlag(v bool) *AddCustomRosterFieldResponseBodyResult {
	s.ContactSystemFlag = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetDeleted(v bool) *AddCustomRosterFieldResponseBodyResult {
	s.Deleted = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetDerived(v bool) *AddCustomRosterFieldResponseBodyResult {
	s.Derived = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetDisabled(v int32) *AddCustomRosterFieldResponseBodyResult {
	s.Disabled = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetEditFromEmployeeFlag(v bool) *AddCustomRosterFieldResponseBodyResult {
	s.EditFromEmployeeFlag = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetEditableByHr(v bool) *AddCustomRosterFieldResponseBodyResult {
	s.EditableByHr = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetFieldCode(v string) *AddCustomRosterFieldResponseBodyResult {
	s.FieldCode = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetFieldName(v string) *AddCustomRosterFieldResponseBodyResult {
	s.FieldName = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetFieldTip(v string) *AddCustomRosterFieldResponseBodyResult {
	s.FieldTip = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetFieldType(v string) *AddCustomRosterFieldResponseBodyResult {
	s.FieldType = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetGroupId(v string) *AddCustomRosterFieldResponseBodyResult {
	s.GroupId = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetHiddenFromEmployeeFlag(v bool) *AddCustomRosterFieldResponseBodyResult {
	s.HiddenFromEmployeeFlag = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetHint(v string) *AddCustomRosterFieldResponseBodyResult {
	s.Hint = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetHistoryField(v bool) *AddCustomRosterFieldResponseBodyResult {
	s.HistoryField = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetIndex(v int32) *AddCustomRosterFieldResponseBodyResult {
	s.Index = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetModifyOptions(v bool) *AddCustomRosterFieldResponseBodyResult {
	s.ModifyOptions = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetNoWatermark(v bool) *AddCustomRosterFieldResponseBodyResult {
	s.NoWatermark = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetNumberDecimalPlace(v string) *AddCustomRosterFieldResponseBodyResult {
	s.NumberDecimalPlace = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetNumberFormatType(v string) *AddCustomRosterFieldResponseBodyResult {
	s.NumberFormatType = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetNumberValueType(v string) *AddCustomRosterFieldResponseBodyResult {
	s.NumberValueType = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetOptionText(v string) *AddCustomRosterFieldResponseBodyResult {
	s.OptionText = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetRequired(v bool) *AddCustomRosterFieldResponseBodyResult {
	s.Required = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetSourceFieldCode(v string) *AddCustomRosterFieldResponseBodyResult {
	s.SourceFieldCode = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetSystemFlag(v bool) *AddCustomRosterFieldResponseBodyResult {
	s.SystemFlag = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetTextToSelectField(v bool) *AddCustomRosterFieldResponseBodyResult {
	s.TextToSelectField = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetValue(v string) *AddCustomRosterFieldResponseBodyResult {
	s.Value = &v
	return s
}

func (s *AddCustomRosterFieldResponseBodyResult) SetVisibleByEmp(v bool) *AddCustomRosterFieldResponseBodyResult {
	s.VisibleByEmp = &v
	return s
}

type AddCustomRosterFieldResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCustomRosterFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCustomRosterFieldResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCustomRosterFieldResponse) GoString() string {
	return s.String()
}

func (s *AddCustomRosterFieldResponse) SetHeaders(v map[string]*string) *AddCustomRosterFieldResponse {
	s.Headers = v
	return s
}

func (s *AddCustomRosterFieldResponse) SetStatusCode(v int32) *AddCustomRosterFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCustomRosterFieldResponse) SetBody(v *AddCustomRosterFieldResponseBody) *AddCustomRosterFieldResponse {
	s.Body = v
	return s
}

type AddHrmLegalEntityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddHrmLegalEntityHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddHrmLegalEntityHeaders) GoString() string {
	return s.String()
}

func (s *AddHrmLegalEntityHeaders) SetCommonHeaders(v map[string]*string) *AddHrmLegalEntityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddHrmLegalEntityHeaders) SetXAcsDingtalkAccessToken(v string) *AddHrmLegalEntityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddHrmLegalEntityRequest struct {
	// This parameter is required.
	CorpId       *string                      `json:"corpId,omitempty" xml:"corpId,omitempty"`
	CreateUserId *string                      `json:"createUserId,omitempty" xml:"createUserId,omitempty"`
	Ext          *AddHrmLegalEntityRequestExt `json:"ext,omitempty" xml:"ext,omitempty" type:"Struct"`
	// This parameter is required.
	LegalEntityName      *string `json:"legalEntityName,omitempty" xml:"legalEntityName,omitempty"`
	LegalEntityShortName *string `json:"legalEntityShortName,omitempty" xml:"legalEntityShortName,omitempty"`
	// This parameter is required.
	LegalEntityStatus *int32  `json:"legalEntityStatus,omitempty" xml:"legalEntityStatus,omitempty"`
	LegalPersonName   *string `json:"legalPersonName,omitempty" xml:"legalPersonName,omitempty"`
	DingTenantId      *int64  `json:"dingTenantId,omitempty" xml:"dingTenantId,omitempty"`
}

func (s AddHrmLegalEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s AddHrmLegalEntityRequest) GoString() string {
	return s.String()
}

func (s *AddHrmLegalEntityRequest) SetCorpId(v string) *AddHrmLegalEntityRequest {
	s.CorpId = &v
	return s
}

func (s *AddHrmLegalEntityRequest) SetCreateUserId(v string) *AddHrmLegalEntityRequest {
	s.CreateUserId = &v
	return s
}

func (s *AddHrmLegalEntityRequest) SetExt(v *AddHrmLegalEntityRequestExt) *AddHrmLegalEntityRequest {
	s.Ext = v
	return s
}

func (s *AddHrmLegalEntityRequest) SetLegalEntityName(v string) *AddHrmLegalEntityRequest {
	s.LegalEntityName = &v
	return s
}

func (s *AddHrmLegalEntityRequest) SetLegalEntityShortName(v string) *AddHrmLegalEntityRequest {
	s.LegalEntityShortName = &v
	return s
}

func (s *AddHrmLegalEntityRequest) SetLegalEntityStatus(v int32) *AddHrmLegalEntityRequest {
	s.LegalEntityStatus = &v
	return s
}

func (s *AddHrmLegalEntityRequest) SetLegalPersonName(v string) *AddHrmLegalEntityRequest {
	s.LegalPersonName = &v
	return s
}

func (s *AddHrmLegalEntityRequest) SetDingTenantId(v int64) *AddHrmLegalEntityRequest {
	s.DingTenantId = &v
	return s
}

type AddHrmLegalEntityRequestExt struct {
	LegalEntityEnName       *string                                         `json:"legalEntityEnName,omitempty" xml:"legalEntityEnName,omitempty"`
	LegalEntityEnShortName  *string                                         `json:"legalEntityEnShortName,omitempty" xml:"legalEntityEnShortName,omitempty"`
	LegalEntityType         *string                                         `json:"legalEntityType,omitempty" xml:"legalEntityType,omitempty"`
	ManageAddress           *AddHrmLegalEntityRequestExtManageAddress       `json:"manageAddress,omitempty" xml:"manageAddress,omitempty" type:"Struct"`
	RegistrationAddress     *AddHrmLegalEntityRequestExtRegistrationAddress `json:"registrationAddress,omitempty" xml:"registrationAddress,omitempty" type:"Struct"`
	RegistrationDate        *int64                                          `json:"registrationDate,omitempty" xml:"registrationDate,omitempty"`
	UnifiedSocialCreditCode *string                                         `json:"unifiedSocialCreditCode,omitempty" xml:"unifiedSocialCreditCode,omitempty"`
	// example:
	//
	// 515200
	ZipCode *string `json:"zipCode,omitempty" xml:"zipCode,omitempty"`
}

func (s AddHrmLegalEntityRequestExt) String() string {
	return tea.Prettify(s)
}

func (s AddHrmLegalEntityRequestExt) GoString() string {
	return s.String()
}

func (s *AddHrmLegalEntityRequestExt) SetLegalEntityEnName(v string) *AddHrmLegalEntityRequestExt {
	s.LegalEntityEnName = &v
	return s
}

func (s *AddHrmLegalEntityRequestExt) SetLegalEntityEnShortName(v string) *AddHrmLegalEntityRequestExt {
	s.LegalEntityEnShortName = &v
	return s
}

func (s *AddHrmLegalEntityRequestExt) SetLegalEntityType(v string) *AddHrmLegalEntityRequestExt {
	s.LegalEntityType = &v
	return s
}

func (s *AddHrmLegalEntityRequestExt) SetManageAddress(v *AddHrmLegalEntityRequestExtManageAddress) *AddHrmLegalEntityRequestExt {
	s.ManageAddress = v
	return s
}

func (s *AddHrmLegalEntityRequestExt) SetRegistrationAddress(v *AddHrmLegalEntityRequestExtRegistrationAddress) *AddHrmLegalEntityRequestExt {
	s.RegistrationAddress = v
	return s
}

func (s *AddHrmLegalEntityRequestExt) SetRegistrationDate(v int64) *AddHrmLegalEntityRequestExt {
	s.RegistrationDate = &v
	return s
}

func (s *AddHrmLegalEntityRequestExt) SetUnifiedSocialCreditCode(v string) *AddHrmLegalEntityRequestExt {
	s.UnifiedSocialCreditCode = &v
	return s
}

func (s *AddHrmLegalEntityRequestExt) SetZipCode(v string) *AddHrmLegalEntityRequestExt {
	s.ZipCode = &v
	return s
}

type AddHrmLegalEntityRequestExtManageAddress struct {
	// example:
	//
	// 110101
	AreaCode *string `json:"areaCode,omitempty" xml:"areaCode,omitempty"`
	// example:
	//
	// 东城区
	AreaName *string `json:"areaName,omitempty" xml:"areaName,omitempty"`
	// example:
	//
	// 123
	CityCode *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	// example:
	//
	// 广州市
	CityName *string `json:"cityName,omitempty" xml:"cityName,omitempty"`
	// example:
	//
	// 123
	CountryCode *string `json:"countryCode,omitempty" xml:"countryCode,omitempty"`
	// example:
	//
	// China
	CountryName *string `json:"countryName,omitempty" xml:"countryName,omitempty"`
	// example:
	//
	// 北京市东城区xx街道xx小区xx楼
	DetailAddress *string `json:"detailAddress,omitempty" xml:"detailAddress,omitempty"`
	// example:
	//
	// 1
	GlobalAreaType *string `json:"globalAreaType,omitempty" xml:"globalAreaType,omitempty"`
	// example:
	//
	// 123
	ProvinceCode *string `json:"provinceCode,omitempty" xml:"provinceCode,omitempty"`
	// example:
	//
	// 广东省
	ProvinceName *string `json:"provinceName,omitempty" xml:"provinceName,omitempty"`
}

func (s AddHrmLegalEntityRequestExtManageAddress) String() string {
	return tea.Prettify(s)
}

func (s AddHrmLegalEntityRequestExtManageAddress) GoString() string {
	return s.String()
}

func (s *AddHrmLegalEntityRequestExtManageAddress) SetAreaCode(v string) *AddHrmLegalEntityRequestExtManageAddress {
	s.AreaCode = &v
	return s
}

func (s *AddHrmLegalEntityRequestExtManageAddress) SetAreaName(v string) *AddHrmLegalEntityRequestExtManageAddress {
	s.AreaName = &v
	return s
}

func (s *AddHrmLegalEntityRequestExtManageAddress) SetCityCode(v string) *AddHrmLegalEntityRequestExtManageAddress {
	s.CityCode = &v
	return s
}

func (s *AddHrmLegalEntityRequestExtManageAddress) SetCityName(v string) *AddHrmLegalEntityRequestExtManageAddress {
	s.CityName = &v
	return s
}

func (s *AddHrmLegalEntityRequestExtManageAddress) SetCountryCode(v string) *AddHrmLegalEntityRequestExtManageAddress {
	s.CountryCode = &v
	return s
}

func (s *AddHrmLegalEntityRequestExtManageAddress) SetCountryName(v string) *AddHrmLegalEntityRequestExtManageAddress {
	s.CountryName = &v
	return s
}

func (s *AddHrmLegalEntityRequestExtManageAddress) SetDetailAddress(v string) *AddHrmLegalEntityRequestExtManageAddress {
	s.DetailAddress = &v
	return s
}

func (s *AddHrmLegalEntityRequestExtManageAddress) SetGlobalAreaType(v string) *AddHrmLegalEntityRequestExtManageAddress {
	s.GlobalAreaType = &v
	return s
}

func (s *AddHrmLegalEntityRequestExtManageAddress) SetProvinceCode(v string) *AddHrmLegalEntityRequestExtManageAddress {
	s.ProvinceCode = &v
	return s
}

func (s *AddHrmLegalEntityRequestExtManageAddress) SetProvinceName(v string) *AddHrmLegalEntityRequestExtManageAddress {
	s.ProvinceName = &v
	return s
}

type AddHrmLegalEntityRequestExtRegistrationAddress struct {
	// example:
	//
	// 110101
	AreaCode *string `json:"areaCode,omitempty" xml:"areaCode,omitempty"`
	// example:
	//
	// 东城区
	AreaName *string `json:"areaName,omitempty" xml:"areaName,omitempty"`
	// example:
	//
	// 123
	CityCode *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	// example:
	//
	// 广州市
	CityName *string `json:"cityName,omitempty" xml:"cityName,omitempty"`
	// example:
	//
	// 123
	CountryCode *string `json:"countryCode,omitempty" xml:"countryCode,omitempty"`
	// example:
	//
	// China
	CountryName *string `json:"countryName,omitempty" xml:"countryName,omitempty"`
	// example:
	//
	// 北京市东城区xx街道xx小区xx楼
	DetailAddress *string `json:"detailAddress,omitempty" xml:"detailAddress,omitempty"`
	// example:
	//
	// 1
	GlobalAreaType *string `json:"globalAreaType,omitempty" xml:"globalAreaType,omitempty"`
	// example:
	//
	// 123
	ProvinceCode *string `json:"provinceCode,omitempty" xml:"provinceCode,omitempty"`
	// example:
	//
	// 广东省
	ProvinceName *string `json:"provinceName,omitempty" xml:"provinceName,omitempty"`
}

func (s AddHrmLegalEntityRequestExtRegistrationAddress) String() string {
	return tea.Prettify(s)
}

func (s AddHrmLegalEntityRequestExtRegistrationAddress) GoString() string {
	return s.String()
}

func (s *AddHrmLegalEntityRequestExtRegistrationAddress) SetAreaCode(v string) *AddHrmLegalEntityRequestExtRegistrationAddress {
	s.AreaCode = &v
	return s
}

func (s *AddHrmLegalEntityRequestExtRegistrationAddress) SetAreaName(v string) *AddHrmLegalEntityRequestExtRegistrationAddress {
	s.AreaName = &v
	return s
}

func (s *AddHrmLegalEntityRequestExtRegistrationAddress) SetCityCode(v string) *AddHrmLegalEntityRequestExtRegistrationAddress {
	s.CityCode = &v
	return s
}

func (s *AddHrmLegalEntityRequestExtRegistrationAddress) SetCityName(v string) *AddHrmLegalEntityRequestExtRegistrationAddress {
	s.CityName = &v
	return s
}

func (s *AddHrmLegalEntityRequestExtRegistrationAddress) SetCountryCode(v string) *AddHrmLegalEntityRequestExtRegistrationAddress {
	s.CountryCode = &v
	return s
}

func (s *AddHrmLegalEntityRequestExtRegistrationAddress) SetCountryName(v string) *AddHrmLegalEntityRequestExtRegistrationAddress {
	s.CountryName = &v
	return s
}

func (s *AddHrmLegalEntityRequestExtRegistrationAddress) SetDetailAddress(v string) *AddHrmLegalEntityRequestExtRegistrationAddress {
	s.DetailAddress = &v
	return s
}

func (s *AddHrmLegalEntityRequestExtRegistrationAddress) SetGlobalAreaType(v string) *AddHrmLegalEntityRequestExtRegistrationAddress {
	s.GlobalAreaType = &v
	return s
}

func (s *AddHrmLegalEntityRequestExtRegistrationAddress) SetProvinceCode(v string) *AddHrmLegalEntityRequestExtRegistrationAddress {
	s.ProvinceCode = &v
	return s
}

func (s *AddHrmLegalEntityRequestExtRegistrationAddress) SetProvinceName(v string) *AddHrmLegalEntityRequestExtRegistrationAddress {
	s.ProvinceName = &v
	return s
}

type AddHrmLegalEntityResponseBody struct {
	Result  *AddHrmLegalEntityResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddHrmLegalEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddHrmLegalEntityResponseBody) GoString() string {
	return s.String()
}

func (s *AddHrmLegalEntityResponseBody) SetResult(v *AddHrmLegalEntityResponseBodyResult) *AddHrmLegalEntityResponseBody {
	s.Result = v
	return s
}

func (s *AddHrmLegalEntityResponseBody) SetSuccess(v bool) *AddHrmLegalEntityResponseBody {
	s.Success = &v
	return s
}

type AddHrmLegalEntityResponseBodyResult struct {
	// example:
	//
	// ding123
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// 2023-01-01
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2023-01-01
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 1234567
	LegalEntityId *string `json:"legalEntityId,omitempty" xml:"legalEntityId,omitempty"`
	// example:
	//
	// 公司1
	LegalEntityName *string `json:"legalEntityName,omitempty" xml:"legalEntityName,omitempty"`
	// example:
	//
	// 公1
	LegalEntityShortName *string `json:"legalEntityShortName,omitempty" xml:"legalEntityShortName,omitempty"`
	// example:
	//
	// 1
	LegalEntityStatus *int32 `json:"legalEntityStatus,omitempty" xml:"legalEntityStatus,omitempty"`
	// example:
	//
	// 法人
	LegalPersonName *string `json:"legalPersonName,omitempty" xml:"legalPersonName,omitempty"`
}

func (s AddHrmLegalEntityResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddHrmLegalEntityResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddHrmLegalEntityResponseBodyResult) SetCorpId(v string) *AddHrmLegalEntityResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *AddHrmLegalEntityResponseBodyResult) SetGmtCreate(v int64) *AddHrmLegalEntityResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *AddHrmLegalEntityResponseBodyResult) SetGmtModified(v int64) *AddHrmLegalEntityResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *AddHrmLegalEntityResponseBodyResult) SetLegalEntityId(v string) *AddHrmLegalEntityResponseBodyResult {
	s.LegalEntityId = &v
	return s
}

func (s *AddHrmLegalEntityResponseBodyResult) SetLegalEntityName(v string) *AddHrmLegalEntityResponseBodyResult {
	s.LegalEntityName = &v
	return s
}

func (s *AddHrmLegalEntityResponseBodyResult) SetLegalEntityShortName(v string) *AddHrmLegalEntityResponseBodyResult {
	s.LegalEntityShortName = &v
	return s
}

func (s *AddHrmLegalEntityResponseBodyResult) SetLegalEntityStatus(v int32) *AddHrmLegalEntityResponseBodyResult {
	s.LegalEntityStatus = &v
	return s
}

func (s *AddHrmLegalEntityResponseBodyResult) SetLegalPersonName(v string) *AddHrmLegalEntityResponseBodyResult {
	s.LegalPersonName = &v
	return s
}

type AddHrmLegalEntityResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddHrmLegalEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddHrmLegalEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s AddHrmLegalEntityResponse) GoString() string {
	return s.String()
}

func (s *AddHrmLegalEntityResponse) SetHeaders(v map[string]*string) *AddHrmLegalEntityResponse {
	s.Headers = v
	return s
}

func (s *AddHrmLegalEntityResponse) SetStatusCode(v int32) *AddHrmLegalEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *AddHrmLegalEntityResponse) SetBody(v *AddHrmLegalEntityResponseBody) *AddHrmLegalEntityResponse {
	s.Body = v
	return s
}

type AddHrmPreentryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddHrmPreentryHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryHeaders) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryHeaders) SetCommonHeaders(v map[string]*string) *AddHrmPreentryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddHrmPreentryHeaders) SetXAcsDingtalkAccessToken(v string) *AddHrmPreentryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddHrmPreentryRequest struct {
	AgentId *int64                         `json:"agentId,omitempty" xml:"agentId,omitempty"`
	Groups  []*AddHrmPreentryRequestGroups `json:"groups,omitempty" xml:"groups,omitempty" type:"Repeated"`
	// This parameter is required.
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// This parameter is required.
	Name                *string `json:"name,omitempty" xml:"name,omitempty"`
	NeedSendPreEntryMsg *bool   `json:"needSendPreEntryMsg,omitempty" xml:"needSendPreEntryMsg,omitempty"`
	PreEntryTime        *int64  `json:"preEntryTime,omitempty" xml:"preEntryTime,omitempty"`
}

func (s AddHrmPreentryRequest) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryRequest) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryRequest) SetAgentId(v int64) *AddHrmPreentryRequest {
	s.AgentId = &v
	return s
}

func (s *AddHrmPreentryRequest) SetGroups(v []*AddHrmPreentryRequestGroups) *AddHrmPreentryRequest {
	s.Groups = v
	return s
}

func (s *AddHrmPreentryRequest) SetMobile(v string) *AddHrmPreentryRequest {
	s.Mobile = &v
	return s
}

func (s *AddHrmPreentryRequest) SetName(v string) *AddHrmPreentryRequest {
	s.Name = &v
	return s
}

func (s *AddHrmPreentryRequest) SetNeedSendPreEntryMsg(v bool) *AddHrmPreentryRequest {
	s.NeedSendPreEntryMsg = &v
	return s
}

func (s *AddHrmPreentryRequest) SetPreEntryTime(v int64) *AddHrmPreentryRequest {
	s.PreEntryTime = &v
	return s
}

type AddHrmPreentryRequestGroups struct {
	// example:
	//
	// sys01
	GroupId  *string                                `json:"groupId,omitempty" xml:"groupId,omitempty"`
	Sections []*AddHrmPreentryRequestGroupsSections `json:"sections,omitempty" xml:"sections,omitempty" type:"Repeated"`
}

func (s AddHrmPreentryRequestGroups) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryRequestGroups) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryRequestGroups) SetGroupId(v string) *AddHrmPreentryRequestGroups {
	s.GroupId = &v
	return s
}

func (s *AddHrmPreentryRequestGroups) SetSections(v []*AddHrmPreentryRequestGroupsSections) *AddHrmPreentryRequestGroups {
	s.Sections = v
	return s
}

type AddHrmPreentryRequestGroupsSections struct {
	EmpFieldVOList []*AddHrmPreentryRequestGroupsSectionsEmpFieldVOList `json:"empFieldVOList,omitempty" xml:"empFieldVOList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	OldIndex *int32 `json:"oldIndex,omitempty" xml:"oldIndex,omitempty"`
}

func (s AddHrmPreentryRequestGroupsSections) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryRequestGroupsSections) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryRequestGroupsSections) SetEmpFieldVOList(v []*AddHrmPreentryRequestGroupsSectionsEmpFieldVOList) *AddHrmPreentryRequestGroupsSections {
	s.EmpFieldVOList = v
	return s
}

func (s *AddHrmPreentryRequestGroupsSections) SetOldIndex(v int32) *AddHrmPreentryRequestGroupsSections {
	s.OldIndex = &v
	return s
}

type AddHrmPreentryRequestGroupsSectionsEmpFieldVOList struct {
	// example:
	//
	// sys01-birthTime
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	// example:
	//
	// 2020-10-10
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AddHrmPreentryRequestGroupsSectionsEmpFieldVOList) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryRequestGroupsSectionsEmpFieldVOList) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryRequestGroupsSectionsEmpFieldVOList) SetFieldCode(v string) *AddHrmPreentryRequestGroupsSectionsEmpFieldVOList {
	s.FieldCode = &v
	return s
}

func (s *AddHrmPreentryRequestGroupsSectionsEmpFieldVOList) SetValue(v string) *AddHrmPreentryRequestGroupsSectionsEmpFieldVOList {
	s.Value = &v
	return s
}

type AddHrmPreentryResponseBody struct {
	// example:
	//
	// manager123
	TmpUserId *string `json:"tmpUserId,omitempty" xml:"tmpUserId,omitempty"`
}

func (s AddHrmPreentryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryResponseBody) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryResponseBody) SetTmpUserId(v string) *AddHrmPreentryResponseBody {
	s.TmpUserId = &v
	return s
}

type AddHrmPreentryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddHrmPreentryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddHrmPreentryResponse) String() string {
	return tea.Prettify(s)
}

func (s AddHrmPreentryResponse) GoString() string {
	return s.String()
}

func (s *AddHrmPreentryResponse) SetHeaders(v map[string]*string) *AddHrmPreentryResponse {
	s.Headers = v
	return s
}

func (s *AddHrmPreentryResponse) SetStatusCode(v int32) *AddHrmPreentryResponse {
	s.StatusCode = &v
	return s
}

func (s *AddHrmPreentryResponse) SetBody(v *AddHrmPreentryResponseBody) *AddHrmPreentryResponse {
	s.Body = v
	return s
}

type AddRosterFieldFormHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddRosterFieldFormHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddRosterFieldFormHeaders) GoString() string {
	return s.String()
}

func (s *AddRosterFieldFormHeaders) SetCommonHeaders(v map[string]*string) *AddRosterFieldFormHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddRosterFieldFormHeaders) SetXAcsDingtalkAccessToken(v string) *AddRosterFieldFormHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddRosterFieldFormRequest struct {
	// This parameter is required.
	Detail *bool `json:"detail,omitempty" xml:"detail,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s AddRosterFieldFormRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRosterFieldFormRequest) GoString() string {
	return s.String()
}

func (s *AddRosterFieldFormRequest) SetDetail(v bool) *AddRosterFieldFormRequest {
	s.Detail = &v
	return s
}

func (s *AddRosterFieldFormRequest) SetName(v string) *AddRosterFieldFormRequest {
	s.Name = &v
	return s
}

type AddRosterFieldFormResponseBody struct {
	DingOpenErrcode *int32                                `json:"dingOpenErrcode,omitempty" xml:"dingOpenErrcode,omitempty"`
	ErrorMsg        *string                               `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result          *AddRosterFieldFormResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success         *bool                                 `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddRosterFieldFormResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRosterFieldFormResponseBody) GoString() string {
	return s.String()
}

func (s *AddRosterFieldFormResponseBody) SetDingOpenErrcode(v int32) *AddRosterFieldFormResponseBody {
	s.DingOpenErrcode = &v
	return s
}

func (s *AddRosterFieldFormResponseBody) SetErrorMsg(v string) *AddRosterFieldFormResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *AddRosterFieldFormResponseBody) SetResult(v *AddRosterFieldFormResponseBodyResult) *AddRosterFieldFormResponseBody {
	s.Result = v
	return s
}

func (s *AddRosterFieldFormResponseBody) SetSuccess(v bool) *AddRosterFieldFormResponseBody {
	s.Success = &v
	return s
}

type AddRosterFieldFormResponseBodyResult struct {
	BizGroupId  *int64      `json:"bizGroupId,omitempty" xml:"bizGroupId,omitempty"`
	Content     *string     `json:"content,omitempty" xml:"content,omitempty"`
	CorpId      *string     `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeleteFlag  *string     `json:"deleteFlag,omitempty" xml:"deleteFlag,omitempty"`
	Detail      *bool       `json:"detail,omitempty" xml:"detail,omitempty"`
	FormId      *string     `json:"formId,omitempty" xml:"formId,omitempty"`
	GmtCreate   interface{} `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified interface{} `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Icon        *string     `json:"icon,omitempty" xml:"icon,omitempty"`
	Id          *int64      `json:"id,omitempty" xml:"id,omitempty"`
	Memo        *string     `json:"memo,omitempty" xml:"memo,omitempty"`
	Name        *string     `json:"name,omitempty" xml:"name,omitempty"`
	SortOrder   *int32      `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	VersionId   *int32      `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s AddRosterFieldFormResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddRosterFieldFormResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddRosterFieldFormResponseBodyResult) SetBizGroupId(v int64) *AddRosterFieldFormResponseBodyResult {
	s.BizGroupId = &v
	return s
}

func (s *AddRosterFieldFormResponseBodyResult) SetContent(v string) *AddRosterFieldFormResponseBodyResult {
	s.Content = &v
	return s
}

func (s *AddRosterFieldFormResponseBodyResult) SetCorpId(v string) *AddRosterFieldFormResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *AddRosterFieldFormResponseBodyResult) SetDeleteFlag(v string) *AddRosterFieldFormResponseBodyResult {
	s.DeleteFlag = &v
	return s
}

func (s *AddRosterFieldFormResponseBodyResult) SetDetail(v bool) *AddRosterFieldFormResponseBodyResult {
	s.Detail = &v
	return s
}

func (s *AddRosterFieldFormResponseBodyResult) SetFormId(v string) *AddRosterFieldFormResponseBodyResult {
	s.FormId = &v
	return s
}

func (s *AddRosterFieldFormResponseBodyResult) SetGmtCreate(v interface{}) *AddRosterFieldFormResponseBodyResult {
	s.GmtCreate = v
	return s
}

func (s *AddRosterFieldFormResponseBodyResult) SetGmtModified(v interface{}) *AddRosterFieldFormResponseBodyResult {
	s.GmtModified = v
	return s
}

func (s *AddRosterFieldFormResponseBodyResult) SetIcon(v string) *AddRosterFieldFormResponseBodyResult {
	s.Icon = &v
	return s
}

func (s *AddRosterFieldFormResponseBodyResult) SetId(v int64) *AddRosterFieldFormResponseBodyResult {
	s.Id = &v
	return s
}

func (s *AddRosterFieldFormResponseBodyResult) SetMemo(v string) *AddRosterFieldFormResponseBodyResult {
	s.Memo = &v
	return s
}

func (s *AddRosterFieldFormResponseBodyResult) SetName(v string) *AddRosterFieldFormResponseBodyResult {
	s.Name = &v
	return s
}

func (s *AddRosterFieldFormResponseBodyResult) SetSortOrder(v int32) *AddRosterFieldFormResponseBodyResult {
	s.SortOrder = &v
	return s
}

func (s *AddRosterFieldFormResponseBodyResult) SetVersionId(v int32) *AddRosterFieldFormResponseBodyResult {
	s.VersionId = &v
	return s
}

type AddRosterFieldFormResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRosterFieldFormResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRosterFieldFormResponse) String() string {
	return tea.Prettify(s)
}

func (s AddRosterFieldFormResponse) GoString() string {
	return s.String()
}

func (s *AddRosterFieldFormResponse) SetHeaders(v map[string]*string) *AddRosterFieldFormResponse {
	s.Headers = v
	return s
}

func (s *AddRosterFieldFormResponse) SetStatusCode(v int32) *AddRosterFieldFormResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRosterFieldFormResponse) SetBody(v *AddRosterFieldFormResponseBody) *AddRosterFieldFormResponse {
	s.Body = v
	return s
}

type CreateRecordHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateRecordHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordHeaders) GoString() string {
	return s.String()
}

func (s *CreateRecordHeaders) SetCommonHeaders(v map[string]*string) *CreateRecordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateRecordHeaders) SetXAcsDingtalkAccessToken(v string) *CreateRecordHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateRecordRequest struct {
	AttachmentList []*CreateRecordRequestAttachmentList `json:"attachmentList,omitempty" xml:"attachmentList,omitempty" type:"Repeated"`
	// example:
	//
	// 908608088
	DeptId    *int64                          `json:"deptId,omitempty" xml:"deptId,omitempty"`
	FieldList []*CreateRecordRequestFieldList `json:"fieldList,omitempty" xml:"fieldList,omitempty" type:"Repeated"`
	GroupList []*CreateRecordRequestGroupList `json:"groupList,omitempty" xml:"groupList,omitempty" type:"Repeated"`
	OuterId   *string                         `json:"outerId,omitempty" xml:"outerId,omitempty"`
	// example:
	//
	// xxx员工劳动合同电子签署
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// xxx有限公司
	SignLastLegalEntityName *string `json:"signLastLegalEntityName,omitempty" xml:"signLastLegalEntityName,omitempty"`
	// example:
	//
	// xxx有限公司
	SignLegalEntityName *string `json:"signLegalEntityName,omitempty" xml:"signLegalEntityName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CONTRACT
	SignSource *string `json:"signSource,omitempty" xml:"signSource,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 48510731071405348944
	SignStartUserId *string `json:"signStartUserId,omitempty" xml:"signStartUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 660658
	SignUserId *string `json:"signUserId,omitempty" xml:"signUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9ad11eb3daa24a9692037079e0732f13
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s CreateRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordRequest) GoString() string {
	return s.String()
}

func (s *CreateRecordRequest) SetAttachmentList(v []*CreateRecordRequestAttachmentList) *CreateRecordRequest {
	s.AttachmentList = v
	return s
}

func (s *CreateRecordRequest) SetDeptId(v int64) *CreateRecordRequest {
	s.DeptId = &v
	return s
}

func (s *CreateRecordRequest) SetFieldList(v []*CreateRecordRequestFieldList) *CreateRecordRequest {
	s.FieldList = v
	return s
}

func (s *CreateRecordRequest) SetGroupList(v []*CreateRecordRequestGroupList) *CreateRecordRequest {
	s.GroupList = v
	return s
}

func (s *CreateRecordRequest) SetOuterId(v string) *CreateRecordRequest {
	s.OuterId = &v
	return s
}

func (s *CreateRecordRequest) SetRemark(v string) *CreateRecordRequest {
	s.Remark = &v
	return s
}

func (s *CreateRecordRequest) SetSignLastLegalEntityName(v string) *CreateRecordRequest {
	s.SignLastLegalEntityName = &v
	return s
}

func (s *CreateRecordRequest) SetSignLegalEntityName(v string) *CreateRecordRequest {
	s.SignLegalEntityName = &v
	return s
}

func (s *CreateRecordRequest) SetSignSource(v string) *CreateRecordRequest {
	s.SignSource = &v
	return s
}

func (s *CreateRecordRequest) SetSignStartUserId(v string) *CreateRecordRequest {
	s.SignStartUserId = &v
	return s
}

func (s *CreateRecordRequest) SetSignUserId(v string) *CreateRecordRequest {
	s.SignUserId = &v
	return s
}

func (s *CreateRecordRequest) SetTemplateId(v string) *CreateRecordRequest {
	s.TemplateId = &v
	return s
}

type CreateRecordRequestAttachmentList struct {
	// example:
	//
	// attachment_profile
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	// example:
	//
	// 简历附件
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// example:
	//
	// DDAttachmentField
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	// example:
	//
	// https://dt-staging-moka-public.oss-cn-zhangjiakou.aliyuncs.com/form/attachment/b32509e4a809cb4e18a72fc4aa75e655.pdf
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
	// example:
	//
	// attachment
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
}

func (s CreateRecordRequestAttachmentList) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordRequestAttachmentList) GoString() string {
	return s.String()
}

func (s *CreateRecordRequestAttachmentList) SetFieldCode(v string) *CreateRecordRequestAttachmentList {
	s.FieldCode = &v
	return s
}

func (s *CreateRecordRequestAttachmentList) SetFieldName(v string) *CreateRecordRequestAttachmentList {
	s.FieldName = &v
	return s
}

func (s *CreateRecordRequestAttachmentList) SetFieldType(v string) *CreateRecordRequestAttachmentList {
	s.FieldType = &v
	return s
}

func (s *CreateRecordRequestAttachmentList) SetFieldValue(v string) *CreateRecordRequestAttachmentList {
	s.FieldValue = &v
	return s
}

func (s *CreateRecordRequestAttachmentList) SetGroupId(v string) *CreateRecordRequestAttachmentList {
	s.GroupId = &v
	return s
}

type CreateRecordRequestFieldList struct {
	// example:
	//
	// contract.contract_type
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	// example:
	//
	// 合同类型
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// example:
	//
	// DDSelectField
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	// example:
	//
	// 劳动合同
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
	// example:
	//
	// contract
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// 劳动合同
	OptionId *string `json:"optionId,omitempty" xml:"optionId,omitempty"`
	// example:
	//
	// [{\"label\":\"劳动合同\",\"value\":\"劳动合同\"},{\"label\":\"保密协议\",\"value\":\"保密协议\"}]
	Options         *string `json:"options,omitempty" xml:"options,omitempty"`
	SignRequired    *bool   `json:"signRequired,omitempty" xml:"signRequired,omitempty"`
	UserCustomField *bool   `json:"userCustomField,omitempty" xml:"userCustomField,omitempty"`
}

func (s CreateRecordRequestFieldList) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordRequestFieldList) GoString() string {
	return s.String()
}

func (s *CreateRecordRequestFieldList) SetFieldCode(v string) *CreateRecordRequestFieldList {
	s.FieldCode = &v
	return s
}

func (s *CreateRecordRequestFieldList) SetFieldName(v string) *CreateRecordRequestFieldList {
	s.FieldName = &v
	return s
}

func (s *CreateRecordRequestFieldList) SetFieldType(v string) *CreateRecordRequestFieldList {
	s.FieldType = &v
	return s
}

func (s *CreateRecordRequestFieldList) SetFieldValue(v string) *CreateRecordRequestFieldList {
	s.FieldValue = &v
	return s
}

func (s *CreateRecordRequestFieldList) SetGroupId(v string) *CreateRecordRequestFieldList {
	s.GroupId = &v
	return s
}

func (s *CreateRecordRequestFieldList) SetOptionId(v string) *CreateRecordRequestFieldList {
	s.OptionId = &v
	return s
}

func (s *CreateRecordRequestFieldList) SetOptions(v string) *CreateRecordRequestFieldList {
	s.Options = &v
	return s
}

func (s *CreateRecordRequestFieldList) SetSignRequired(v bool) *CreateRecordRequestFieldList {
	s.SignRequired = &v
	return s
}

func (s *CreateRecordRequestFieldList) SetUserCustomField(v bool) *CreateRecordRequestFieldList {
	s.UserCustomField = &v
	return s
}

type CreateRecordRequestGroupList struct {
	DetailFlag *bool                                      `json:"detailFlag,omitempty" xml:"detailFlag,omitempty"`
	FieldList  [][]*CreateRecordRequestGroupListFieldList `json:"fieldList,omitempty" xml:"fieldList,omitempty" type:"Repeated"`
	// example:
	//
	// family
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// 家庭成员
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
}

func (s CreateRecordRequestGroupList) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordRequestGroupList) GoString() string {
	return s.String()
}

func (s *CreateRecordRequestGroupList) SetDetailFlag(v bool) *CreateRecordRequestGroupList {
	s.DetailFlag = &v
	return s
}

func (s *CreateRecordRequestGroupList) SetFieldList(v [][]*CreateRecordRequestGroupListFieldList) *CreateRecordRequestGroupList {
	s.FieldList = v
	return s
}

func (s *CreateRecordRequestGroupList) SetGroupId(v string) *CreateRecordRequestGroupList {
	s.GroupId = &v
	return s
}

func (s *CreateRecordRequestGroupList) SetGroupName(v string) *CreateRecordRequestGroupList {
	s.GroupName = &v
	return s
}

type CreateRecordRequestGroupListFieldList struct {
	// example:
	//
	// contract.contract_type
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	// example:
	//
	// 合同类型
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// example:
	//
	// DDSelectField
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	// example:
	//
	// 劳动合同
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
	// example:
	//
	// [{\"label\":\"劳动合同\",\"value\":\"劳动合同\"},{\"label\":\"培训协议\",\"value\":\"培训协议\"}]
	Options *string `json:"options,omitempty" xml:"options,omitempty"`
	// example:
	//
	// 劳动合同
	OptionId *string `json:"optionId,omitempty" xml:"optionId,omitempty"`
	// example:
	//
	// contract
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
}

func (s CreateRecordRequestGroupListFieldList) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordRequestGroupListFieldList) GoString() string {
	return s.String()
}

func (s *CreateRecordRequestGroupListFieldList) SetFieldCode(v string) *CreateRecordRequestGroupListFieldList {
	s.FieldCode = &v
	return s
}

func (s *CreateRecordRequestGroupListFieldList) SetFieldName(v string) *CreateRecordRequestGroupListFieldList {
	s.FieldName = &v
	return s
}

func (s *CreateRecordRequestGroupListFieldList) SetFieldType(v string) *CreateRecordRequestGroupListFieldList {
	s.FieldType = &v
	return s
}

func (s *CreateRecordRequestGroupListFieldList) SetFieldValue(v string) *CreateRecordRequestGroupListFieldList {
	s.FieldValue = &v
	return s
}

func (s *CreateRecordRequestGroupListFieldList) SetOptions(v string) *CreateRecordRequestGroupListFieldList {
	s.Options = &v
	return s
}

func (s *CreateRecordRequestGroupListFieldList) SetOptionId(v string) *CreateRecordRequestGroupListFieldList {
	s.OptionId = &v
	return s
}

func (s *CreateRecordRequestGroupListFieldList) SetGroupId(v string) *CreateRecordRequestGroupListFieldList {
	s.GroupId = &v
	return s
}

type CreateRecordResponseBody struct {
	Result  *CreateRecordResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecordResponseBody) SetResult(v *CreateRecordResponseBodyResult) *CreateRecordResponseBody {
	s.Result = v
	return s
}

func (s *CreateRecordResponseBody) SetSuccess(v bool) *CreateRecordResponseBody {
	s.Success = &v
	return s
}

type CreateRecordResponseBodyResult struct {
	Details *string `json:"details,omitempty" xml:"details,omitempty"`
	ItemId  *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	Type    *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateRecordResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateRecordResponseBodyResult) SetDetails(v string) *CreateRecordResponseBodyResult {
	s.Details = &v
	return s
}

func (s *CreateRecordResponseBodyResult) SetItemId(v string) *CreateRecordResponseBodyResult {
	s.ItemId = &v
	return s
}

func (s *CreateRecordResponseBodyResult) SetType(v string) *CreateRecordResponseBodyResult {
	s.Type = &v
	return s
}

type CreateRecordResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordResponse) GoString() string {
	return s.String()
}

func (s *CreateRecordResponse) SetHeaders(v map[string]*string) *CreateRecordResponse {
	s.Headers = v
	return s
}

func (s *CreateRecordResponse) SetStatusCode(v int32) *CreateRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRecordResponse) SetBody(v *CreateRecordResponseBody) *CreateRecordResponse {
	s.Body = v
	return s
}

type DeleteCustomRosterFieldHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteCustomRosterFieldHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomRosterFieldHeaders) GoString() string {
	return s.String()
}

func (s *DeleteCustomRosterFieldHeaders) SetCommonHeaders(v map[string]*string) *DeleteCustomRosterFieldHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteCustomRosterFieldHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteCustomRosterFieldHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteCustomRosterFieldRequest struct {
	// This parameter is required.
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	// This parameter is required.
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
}

func (s DeleteCustomRosterFieldRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomRosterFieldRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomRosterFieldRequest) SetFieldCode(v string) *DeleteCustomRosterFieldRequest {
	s.FieldCode = &v
	return s
}

func (s *DeleteCustomRosterFieldRequest) SetGroupId(v string) *DeleteCustomRosterFieldRequest {
	s.GroupId = &v
	return s
}

type DeleteCustomRosterFieldResponseBody struct {
	DingOpenErrcode *int32  `json:"dingOpenErrcode,omitempty" xml:"dingOpenErrcode,omitempty"`
	ErrorMsg        *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result          *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success         *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteCustomRosterFieldResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomRosterFieldResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomRosterFieldResponseBody) SetDingOpenErrcode(v int32) *DeleteCustomRosterFieldResponseBody {
	s.DingOpenErrcode = &v
	return s
}

func (s *DeleteCustomRosterFieldResponseBody) SetErrorMsg(v string) *DeleteCustomRosterFieldResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteCustomRosterFieldResponseBody) SetResult(v bool) *DeleteCustomRosterFieldResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteCustomRosterFieldResponseBody) SetSuccess(v bool) *DeleteCustomRosterFieldResponseBody {
	s.Success = &v
	return s
}

type DeleteCustomRosterFieldResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomRosterFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomRosterFieldResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomRosterFieldResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomRosterFieldResponse) SetHeaders(v map[string]*string) *DeleteCustomRosterFieldResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomRosterFieldResponse) SetStatusCode(v int32) *DeleteCustomRosterFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomRosterFieldResponse) SetBody(v *DeleteCustomRosterFieldResponseBody) *DeleteCustomRosterFieldResponse {
	s.Body = v
	return s
}

type DeleteRosterFieldFormHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteRosterFieldFormHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteRosterFieldFormHeaders) GoString() string {
	return s.String()
}

func (s *DeleteRosterFieldFormHeaders) SetCommonHeaders(v map[string]*string) *DeleteRosterFieldFormHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteRosterFieldFormHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteRosterFieldFormHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteRosterFieldFormRequest struct {
	// This parameter is required.
	FormId *string `json:"formId,omitempty" xml:"formId,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeleteRosterFieldFormRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRosterFieldFormRequest) GoString() string {
	return s.String()
}

func (s *DeleteRosterFieldFormRequest) SetFormId(v string) *DeleteRosterFieldFormRequest {
	s.FormId = &v
	return s
}

func (s *DeleteRosterFieldFormRequest) SetUserId(v string) *DeleteRosterFieldFormRequest {
	s.UserId = &v
	return s
}

type DeleteRosterFieldFormResponseBody struct {
	DingOpenErrcode *int32  `json:"dingOpenErrcode,omitempty" xml:"dingOpenErrcode,omitempty"`
	ErrorMsg        *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result          *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success         *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteRosterFieldFormResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRosterFieldFormResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRosterFieldFormResponseBody) SetDingOpenErrcode(v int32) *DeleteRosterFieldFormResponseBody {
	s.DingOpenErrcode = &v
	return s
}

func (s *DeleteRosterFieldFormResponseBody) SetErrorMsg(v string) *DeleteRosterFieldFormResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteRosterFieldFormResponseBody) SetResult(v bool) *DeleteRosterFieldFormResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteRosterFieldFormResponseBody) SetSuccess(v bool) *DeleteRosterFieldFormResponseBody {
	s.Success = &v
	return s
}

type DeleteRosterFieldFormResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRosterFieldFormResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRosterFieldFormResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRosterFieldFormResponse) GoString() string {
	return s.String()
}

func (s *DeleteRosterFieldFormResponse) SetHeaders(v map[string]*string) *DeleteRosterFieldFormResponse {
	s.Headers = v
	return s
}

func (s *DeleteRosterFieldFormResponse) SetStatusCode(v int32) *DeleteRosterFieldFormResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRosterFieldFormResponse) SetBody(v *DeleteRosterFieldFormResponseBody) *DeleteRosterFieldFormResponse {
	s.Body = v
	return s
}

type DeviceMarketManagerResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeviceMarketManagerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeviceMarketManagerResponseBody) GoString() string {
	return s.String()
}

func (s *DeviceMarketManagerResponseBody) SetRequestId(v string) *DeviceMarketManagerResponseBody {
	s.RequestId = &v
	return s
}

type DeviceMarketManagerResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeviceMarketManagerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeviceMarketManagerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeviceMarketManagerResponse) GoString() string {
	return s.String()
}

func (s *DeviceMarketManagerResponse) SetHeaders(v map[string]*string) *DeviceMarketManagerResponse {
	s.Headers = v
	return s
}

func (s *DeviceMarketManagerResponse) SetStatusCode(v int32) *DeviceMarketManagerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeviceMarketManagerResponse) SetBody(v *DeviceMarketManagerResponseBody) *DeviceMarketManagerResponse {
	s.Body = v
	return s
}

type DeviceMarketOrderManagerResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeviceMarketOrderManagerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeviceMarketOrderManagerResponseBody) GoString() string {
	return s.String()
}

func (s *DeviceMarketOrderManagerResponseBody) SetRequestId(v string) *DeviceMarketOrderManagerResponseBody {
	s.RequestId = &v
	return s
}

type DeviceMarketOrderManagerResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeviceMarketOrderManagerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeviceMarketOrderManagerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeviceMarketOrderManagerResponse) GoString() string {
	return s.String()
}

func (s *DeviceMarketOrderManagerResponse) SetHeaders(v map[string]*string) *DeviceMarketOrderManagerResponse {
	s.Headers = v
	return s
}

func (s *DeviceMarketOrderManagerResponse) SetStatusCode(v int32) *DeviceMarketOrderManagerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeviceMarketOrderManagerResponse) SetBody(v *DeviceMarketOrderManagerResponseBody) *DeviceMarketOrderManagerResponse {
	s.Body = v
	return s
}

type ECertQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ECertQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s ECertQueryHeaders) GoString() string {
	return s.String()
}

func (s *ECertQueryHeaders) SetCommonHeaders(v map[string]*string) *ECertQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ECertQueryHeaders) SetXAcsDingtalkAccessToken(v string) *ECertQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ECertQueryRequest struct {
	// if can be null:
	// false
	//
	// example:
	//
	// manager231
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ECertQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s ECertQueryRequest) GoString() string {
	return s.String()
}

func (s *ECertQueryRequest) SetUserId(v string) *ECertQueryRequest {
	s.UserId = &v
	return s
}

type ECertQueryResponseBody struct {
	// example:
	//
	// 3300111192912113
	CertNO *string `json:"certNO,omitempty" xml:"certNO,omitempty"`
	// example:
	//
	// 1123124124124
	EmployJobId *string `json:"employJobId,omitempty" xml:"employJobId,omitempty"`
	// example:
	//
	// 职务
	EmployJobIdLabel *string `json:"employJobIdLabel,omitempty" xml:"employJobIdLabel,omitempty"`
	// example:
	//
	// 1231231232313123
	EmployPositionId *string `json:"employPositionId,omitempty" xml:"employPositionId,omitempty"`
	// example:
	//
	// 职位
	EmployPositionIdLabel *string `json:"employPositionIdLabel,omitempty" xml:"employPositionIdLabel,omitempty"`
	// example:
	//
	// 498192313
	EmployPositionRankId *string `json:"employPositionRankId,omitempty" xml:"employPositionRankId,omitempty"`
	// example:
	//
	// 职级
	EmployPositionRankIdLabel *string `json:"employPositionRankIdLabel,omitempty" xml:"employPositionRankIdLabel,omitempty"`
	// example:
	//
	// 2020-03-14
	HiredDate *string `json:"hiredDate,omitempty" xml:"hiredDate,omitempty"`
	// example:
	//
	// 2021-03-14
	LastWorkDay *string `json:"lastWorkDay,omitempty" xml:"lastWorkDay,omitempty"`
	// example:
	//
	// -1
	MainDeptId *int64 `json:"mainDeptId,omitempty" xml:"mainDeptId,omitempty"`
	// example:
	//
	// 测试部门
	MainDeptName *string `json:"mainDeptName,omitempty" xml:"mainDeptName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 李四
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 张三
	RealName                   *string   `json:"realName,omitempty" xml:"realName,omitempty"`
	TerminationReasonPassive   []*string `json:"terminationReasonPassive,omitempty" xml:"terminationReasonPassive,omitempty" type:"Repeated"`
	TerminationReasonVoluntary []*string `json:"terminationReasonVoluntary,omitempty" xml:"terminationReasonVoluntary,omitempty" type:"Repeated"`
}

func (s ECertQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ECertQueryResponseBody) GoString() string {
	return s.String()
}

func (s *ECertQueryResponseBody) SetCertNO(v string) *ECertQueryResponseBody {
	s.CertNO = &v
	return s
}

func (s *ECertQueryResponseBody) SetEmployJobId(v string) *ECertQueryResponseBody {
	s.EmployJobId = &v
	return s
}

func (s *ECertQueryResponseBody) SetEmployJobIdLabel(v string) *ECertQueryResponseBody {
	s.EmployJobIdLabel = &v
	return s
}

func (s *ECertQueryResponseBody) SetEmployPositionId(v string) *ECertQueryResponseBody {
	s.EmployPositionId = &v
	return s
}

func (s *ECertQueryResponseBody) SetEmployPositionIdLabel(v string) *ECertQueryResponseBody {
	s.EmployPositionIdLabel = &v
	return s
}

func (s *ECertQueryResponseBody) SetEmployPositionRankId(v string) *ECertQueryResponseBody {
	s.EmployPositionRankId = &v
	return s
}

func (s *ECertQueryResponseBody) SetEmployPositionRankIdLabel(v string) *ECertQueryResponseBody {
	s.EmployPositionRankIdLabel = &v
	return s
}

func (s *ECertQueryResponseBody) SetHiredDate(v string) *ECertQueryResponseBody {
	s.HiredDate = &v
	return s
}

func (s *ECertQueryResponseBody) SetLastWorkDay(v string) *ECertQueryResponseBody {
	s.LastWorkDay = &v
	return s
}

func (s *ECertQueryResponseBody) SetMainDeptId(v int64) *ECertQueryResponseBody {
	s.MainDeptId = &v
	return s
}

func (s *ECertQueryResponseBody) SetMainDeptName(v string) *ECertQueryResponseBody {
	s.MainDeptName = &v
	return s
}

func (s *ECertQueryResponseBody) SetName(v string) *ECertQueryResponseBody {
	s.Name = &v
	return s
}

func (s *ECertQueryResponseBody) SetRealName(v string) *ECertQueryResponseBody {
	s.RealName = &v
	return s
}

func (s *ECertQueryResponseBody) SetTerminationReasonPassive(v []*string) *ECertQueryResponseBody {
	s.TerminationReasonPassive = v
	return s
}

func (s *ECertQueryResponseBody) SetTerminationReasonVoluntary(v []*string) *ECertQueryResponseBody {
	s.TerminationReasonVoluntary = v
	return s
}

type ECertQueryResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ECertQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ECertQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s ECertQueryResponse) GoString() string {
	return s.String()
}

func (s *ECertQueryResponse) SetHeaders(v map[string]*string) *ECertQueryResponse {
	s.Headers = v
	return s
}

func (s *ECertQueryResponse) SetStatusCode(v int32) *ECertQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *ECertQueryResponse) SetBody(v *ECertQueryResponseBody) *ECertQueryResponse {
	s.Body = v
	return s
}

type EmpStartDismissionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EmpStartDismissionHeaders) String() string {
	return tea.Prettify(s)
}

func (s EmpStartDismissionHeaders) GoString() string {
	return s.String()
}

func (s *EmpStartDismissionHeaders) SetCommonHeaders(v map[string]*string) *EmpStartDismissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EmpStartDismissionHeaders) SetXAcsDingtalkAccessToken(v string) *EmpStartDismissionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EmpStartDismissionRequest struct {
	// This parameter is required.
	LastWorkDate               *int64    `json:"lastWorkDate,omitempty" xml:"lastWorkDate,omitempty"`
	Partner                    *bool     `json:"partner,omitempty" xml:"partner,omitempty"`
	Remark                     *string   `json:"remark,omitempty" xml:"remark,omitempty"`
	TerminationReasonPassive   []*string `json:"terminationReasonPassive,omitempty" xml:"terminationReasonPassive,omitempty" type:"Repeated"`
	TerminationReasonVoluntary []*string `json:"terminationReasonVoluntary,omitempty" xml:"terminationReasonVoluntary,omitempty" type:"Repeated"`
	ToHireBlackList            *bool     `json:"toHireBlackList,omitempty" xml:"toHireBlackList,omitempty"`
	ToHireDismissionTalent     *bool     `json:"toHireDismissionTalent,omitempty" xml:"toHireDismissionTalent,omitempty"`
	ToHrmBlackList             *bool     `json:"toHrmBlackList,omitempty" xml:"toHrmBlackList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2163515669935611
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s EmpStartDismissionRequest) String() string {
	return tea.Prettify(s)
}

func (s EmpStartDismissionRequest) GoString() string {
	return s.String()
}

func (s *EmpStartDismissionRequest) SetLastWorkDate(v int64) *EmpStartDismissionRequest {
	s.LastWorkDate = &v
	return s
}

func (s *EmpStartDismissionRequest) SetPartner(v bool) *EmpStartDismissionRequest {
	s.Partner = &v
	return s
}

func (s *EmpStartDismissionRequest) SetRemark(v string) *EmpStartDismissionRequest {
	s.Remark = &v
	return s
}

func (s *EmpStartDismissionRequest) SetTerminationReasonPassive(v []*string) *EmpStartDismissionRequest {
	s.TerminationReasonPassive = v
	return s
}

func (s *EmpStartDismissionRequest) SetTerminationReasonVoluntary(v []*string) *EmpStartDismissionRequest {
	s.TerminationReasonVoluntary = v
	return s
}

func (s *EmpStartDismissionRequest) SetToHireBlackList(v bool) *EmpStartDismissionRequest {
	s.ToHireBlackList = &v
	return s
}

func (s *EmpStartDismissionRequest) SetToHireDismissionTalent(v bool) *EmpStartDismissionRequest {
	s.ToHireDismissionTalent = &v
	return s
}

func (s *EmpStartDismissionRequest) SetToHrmBlackList(v bool) *EmpStartDismissionRequest {
	s.ToHrmBlackList = &v
	return s
}

func (s *EmpStartDismissionRequest) SetUserId(v string) *EmpStartDismissionRequest {
	s.UserId = &v
	return s
}

type EmpStartDismissionResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s EmpStartDismissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EmpStartDismissionResponseBody) GoString() string {
	return s.String()
}

func (s *EmpStartDismissionResponseBody) SetRequestId(v string) *EmpStartDismissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *EmpStartDismissionResponseBody) SetResult(v bool) *EmpStartDismissionResponseBody {
	s.Result = &v
	return s
}

func (s *EmpStartDismissionResponseBody) SetSuccess(v bool) *EmpStartDismissionResponseBody {
	s.Success = &v
	return s
}

type EmpStartDismissionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EmpStartDismissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EmpStartDismissionResponse) String() string {
	return tea.Prettify(s)
}

func (s EmpStartDismissionResponse) GoString() string {
	return s.String()
}

func (s *EmpStartDismissionResponse) SetHeaders(v map[string]*string) *EmpStartDismissionResponse {
	s.Headers = v
	return s
}

func (s *EmpStartDismissionResponse) SetStatusCode(v int32) *EmpStartDismissionResponse {
	s.StatusCode = &v
	return s
}

func (s *EmpStartDismissionResponse) SetBody(v *EmpStartDismissionResponseBody) *EmpStartDismissionResponse {
	s.Body = v
	return s
}

type EmployeeAttachmentUpdateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EmployeeAttachmentUpdateHeaders) String() string {
	return tea.Prettify(s)
}

func (s EmployeeAttachmentUpdateHeaders) GoString() string {
	return s.String()
}

func (s *EmployeeAttachmentUpdateHeaders) SetCommonHeaders(v map[string]*string) *EmployeeAttachmentUpdateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EmployeeAttachmentUpdateHeaders) SetXAcsDingtalkAccessToken(v string) *EmployeeAttachmentUpdateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EmployeeAttachmentUpdateRequest struct {
	// This parameter is required.
	AppAgentId *int64 `json:"appAgentId,omitempty" xml:"appAgentId,omitempty"`
	// This parameter is required.
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	// example:
	//
	// .png
	FileSuffix *string `json:"fileSuffix,omitempty" xml:"fileSuffix,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// "$iAEKAqNwbmcDBgTNAk"
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// admin123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s EmployeeAttachmentUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s EmployeeAttachmentUpdateRequest) GoString() string {
	return s.String()
}

func (s *EmployeeAttachmentUpdateRequest) SetAppAgentId(v int64) *EmployeeAttachmentUpdateRequest {
	s.AppAgentId = &v
	return s
}

func (s *EmployeeAttachmentUpdateRequest) SetFieldCode(v string) *EmployeeAttachmentUpdateRequest {
	s.FieldCode = &v
	return s
}

func (s *EmployeeAttachmentUpdateRequest) SetFileSuffix(v string) *EmployeeAttachmentUpdateRequest {
	s.FileSuffix = &v
	return s
}

func (s *EmployeeAttachmentUpdateRequest) SetMediaId(v string) *EmployeeAttachmentUpdateRequest {
	s.MediaId = &v
	return s
}

func (s *EmployeeAttachmentUpdateRequest) SetUserId(v string) *EmployeeAttachmentUpdateRequest {
	s.UserId = &v
	return s
}

type EmployeeAttachmentUpdateResponseBody struct {
	Result  *bool `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s EmployeeAttachmentUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EmployeeAttachmentUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *EmployeeAttachmentUpdateResponseBody) SetResult(v bool) *EmployeeAttachmentUpdateResponseBody {
	s.Result = &v
	return s
}

func (s *EmployeeAttachmentUpdateResponseBody) SetSuccess(v bool) *EmployeeAttachmentUpdateResponseBody {
	s.Success = &v
	return s
}

type EmployeeAttachmentUpdateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EmployeeAttachmentUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EmployeeAttachmentUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s EmployeeAttachmentUpdateResponse) GoString() string {
	return s.String()
}

func (s *EmployeeAttachmentUpdateResponse) SetHeaders(v map[string]*string) *EmployeeAttachmentUpdateResponse {
	s.Headers = v
	return s
}

func (s *EmployeeAttachmentUpdateResponse) SetStatusCode(v int32) *EmployeeAttachmentUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *EmployeeAttachmentUpdateResponse) SetBody(v *EmployeeAttachmentUpdateResponseBody) *EmployeeAttachmentUpdateResponse {
	s.Body = v
	return s
}

type EsignRollbackHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EsignRollbackHeaders) String() string {
	return tea.Prettify(s)
}

func (s EsignRollbackHeaders) GoString() string {
	return s.String()
}

func (s *EsignRollbackHeaders) SetCommonHeaders(v map[string]*string) *EsignRollbackHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EsignRollbackHeaders) SetXAcsDingtalkAccessToken(v string) *EsignRollbackHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EsignRollbackRequest struct {
	// This parameter is required.
	OptUserId *string `json:"optUserId,omitempty" xml:"optUserId,omitempty"`
}

func (s EsignRollbackRequest) String() string {
	return tea.Prettify(s)
}

func (s EsignRollbackRequest) GoString() string {
	return s.String()
}

func (s *EsignRollbackRequest) SetOptUserId(v string) *EsignRollbackRequest {
	s.OptUserId = &v
	return s
}

type EsignRollbackResponseBody struct {
	// This parameter is required.
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s EsignRollbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EsignRollbackResponseBody) GoString() string {
	return s.String()
}

func (s *EsignRollbackResponseBody) SetResult(v bool) *EsignRollbackResponseBody {
	s.Result = &v
	return s
}

type EsignRollbackResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EsignRollbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EsignRollbackResponse) String() string {
	return tea.Prettify(s)
}

func (s EsignRollbackResponse) GoString() string {
	return s.String()
}

func (s *EsignRollbackResponse) SetHeaders(v map[string]*string) *EsignRollbackResponse {
	s.Headers = v
	return s
}

func (s *EsignRollbackResponse) SetStatusCode(v int32) *EsignRollbackResponse {
	s.StatusCode = &v
	return s
}

func (s *EsignRollbackResponse) SetBody(v *EsignRollbackResponseBody) *EsignRollbackResponse {
	s.Body = v
	return s
}

type GetAllDismissionReasonsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAllDismissionReasonsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAllDismissionReasonsHeaders) GoString() string {
	return s.String()
}

func (s *GetAllDismissionReasonsHeaders) SetCommonHeaders(v map[string]*string) *GetAllDismissionReasonsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAllDismissionReasonsHeaders) SetXAcsDingtalkAccessToken(v string) *GetAllDismissionReasonsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAllDismissionReasonsResponseBody struct {
	RequestId *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetAllDismissionReasonsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success   *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetAllDismissionReasonsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAllDismissionReasonsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllDismissionReasonsResponseBody) SetRequestId(v string) *GetAllDismissionReasonsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAllDismissionReasonsResponseBody) SetResult(v *GetAllDismissionReasonsResponseBodyResult) *GetAllDismissionReasonsResponseBody {
	s.Result = v
	return s
}

func (s *GetAllDismissionReasonsResponseBody) SetSuccess(v bool) *GetAllDismissionReasonsResponseBody {
	s.Success = &v
	return s
}

type GetAllDismissionReasonsResponseBodyResult struct {
	PassiveList   []*GetAllDismissionReasonsResponseBodyResultPassiveList   `json:"passiveList,omitempty" xml:"passiveList,omitempty" type:"Repeated"`
	VoluntaryList []*GetAllDismissionReasonsResponseBodyResultVoluntaryList `json:"voluntaryList,omitempty" xml:"voluntaryList,omitempty" type:"Repeated"`
}

func (s GetAllDismissionReasonsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetAllDismissionReasonsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAllDismissionReasonsResponseBodyResult) SetPassiveList(v []*GetAllDismissionReasonsResponseBodyResultPassiveList) *GetAllDismissionReasonsResponseBodyResult {
	s.PassiveList = v
	return s
}

func (s *GetAllDismissionReasonsResponseBodyResult) SetVoluntaryList(v []*GetAllDismissionReasonsResponseBodyResultVoluntaryList) *GetAllDismissionReasonsResponseBodyResult {
	s.VoluntaryList = v
	return s
}

type GetAllDismissionReasonsResponseBodyResultPassiveList struct {
	// example:
	//
	// 1
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 家庭原因
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetAllDismissionReasonsResponseBodyResultPassiveList) String() string {
	return tea.Prettify(s)
}

func (s GetAllDismissionReasonsResponseBodyResultPassiveList) GoString() string {
	return s.String()
}

func (s *GetAllDismissionReasonsResponseBodyResultPassiveList) SetId(v string) *GetAllDismissionReasonsResponseBodyResultPassiveList {
	s.Id = &v
	return s
}

func (s *GetAllDismissionReasonsResponseBodyResultPassiveList) SetName(v string) *GetAllDismissionReasonsResponseBodyResultPassiveList {
	s.Name = &v
	return s
}

type GetAllDismissionReasonsResponseBodyResultVoluntaryList struct {
	// example:
	//
	// 1
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 家庭原因
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetAllDismissionReasonsResponseBodyResultVoluntaryList) String() string {
	return tea.Prettify(s)
}

func (s GetAllDismissionReasonsResponseBodyResultVoluntaryList) GoString() string {
	return s.String()
}

func (s *GetAllDismissionReasonsResponseBodyResultVoluntaryList) SetId(v string) *GetAllDismissionReasonsResponseBodyResultVoluntaryList {
	s.Id = &v
	return s
}

func (s *GetAllDismissionReasonsResponseBodyResultVoluntaryList) SetName(v string) *GetAllDismissionReasonsResponseBodyResultVoluntaryList {
	s.Name = &v
	return s
}

type GetAllDismissionReasonsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAllDismissionReasonsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAllDismissionReasonsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllDismissionReasonsResponse) GoString() string {
	return s.String()
}

func (s *GetAllDismissionReasonsResponse) SetHeaders(v map[string]*string) *GetAllDismissionReasonsResponse {
	s.Headers = v
	return s
}

func (s *GetAllDismissionReasonsResponse) SetStatusCode(v int32) *GetAllDismissionReasonsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAllDismissionReasonsResponse) SetBody(v *GetAllDismissionReasonsResponseBody) *GetAllDismissionReasonsResponse {
	s.Body = v
	return s
}

type GetEmployeeRosterByFieldHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetEmployeeRosterByFieldHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetEmployeeRosterByFieldHeaders) GoString() string {
	return s.String()
}

func (s *GetEmployeeRosterByFieldHeaders) SetCommonHeaders(v map[string]*string) *GetEmployeeRosterByFieldHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetEmployeeRosterByFieldHeaders) SetXAcsDingtalkAccessToken(v string) *GetEmployeeRosterByFieldHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetEmployeeRosterByFieldRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	AppAgentId         *int64    `json:"appAgentId,omitempty" xml:"appAgentId,omitempty"`
	FieldFilterList    []*string `json:"fieldFilterList,omitempty" xml:"fieldFilterList,omitempty" type:"Repeated"`
	Text2SelectConvert *bool     `json:"text2SelectConvert,omitempty" xml:"text2SelectConvert,omitempty"`
	// This parameter is required.
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s GetEmployeeRosterByFieldRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEmployeeRosterByFieldRequest) GoString() string {
	return s.String()
}

func (s *GetEmployeeRosterByFieldRequest) SetAppAgentId(v int64) *GetEmployeeRosterByFieldRequest {
	s.AppAgentId = &v
	return s
}

func (s *GetEmployeeRosterByFieldRequest) SetFieldFilterList(v []*string) *GetEmployeeRosterByFieldRequest {
	s.FieldFilterList = v
	return s
}

func (s *GetEmployeeRosterByFieldRequest) SetText2SelectConvert(v bool) *GetEmployeeRosterByFieldRequest {
	s.Text2SelectConvert = &v
	return s
}

func (s *GetEmployeeRosterByFieldRequest) SetUserIdList(v []*string) *GetEmployeeRosterByFieldRequest {
	s.UserIdList = v
	return s
}

type GetEmployeeRosterByFieldResponseBody struct {
	Result []*GetEmployeeRosterByFieldResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetEmployeeRosterByFieldResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEmployeeRosterByFieldResponseBody) GoString() string {
	return s.String()
}

func (s *GetEmployeeRosterByFieldResponseBody) SetResult(v []*GetEmployeeRosterByFieldResponseBodyResult) *GetEmployeeRosterByFieldResponseBody {
	s.Result = v
	return s
}

type GetEmployeeRosterByFieldResponseBodyResult struct {
	// example:
	//
	// ding20a11xxx
	CorpId        *string                                                    `json:"corpId,omitempty" xml:"corpId,omitempty"`
	FieldDataList []*GetEmployeeRosterByFieldResponseBodyResultFieldDataList `json:"fieldDataList,omitempty" xml:"fieldDataList,omitempty" type:"Repeated"`
	UnionId       *string                                                    `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// example:
	//
	// 042519
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetEmployeeRosterByFieldResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetEmployeeRosterByFieldResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetEmployeeRosterByFieldResponseBodyResult) SetCorpId(v string) *GetEmployeeRosterByFieldResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *GetEmployeeRosterByFieldResponseBodyResult) SetFieldDataList(v []*GetEmployeeRosterByFieldResponseBodyResultFieldDataList) *GetEmployeeRosterByFieldResponseBodyResult {
	s.FieldDataList = v
	return s
}

func (s *GetEmployeeRosterByFieldResponseBodyResult) SetUnionId(v string) *GetEmployeeRosterByFieldResponseBodyResult {
	s.UnionId = &v
	return s
}

func (s *GetEmployeeRosterByFieldResponseBodyResult) SetUserId(v string) *GetEmployeeRosterByFieldResponseBodyResult {
	s.UserId = &v
	return s
}

type GetEmployeeRosterByFieldResponseBodyResultFieldDataList struct {
	// example:
	//
	// sys01-employeeStatus
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	// example:
	//
	// 员工状态
	FieldName      *string                                                                  `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	FieldValueList []*GetEmployeeRosterByFieldResponseBodyResultFieldDataListFieldValueList `json:"fieldValueList,omitempty" xml:"fieldValueList,omitempty" type:"Repeated"`
	// example:
	//
	// sys01
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
}

func (s GetEmployeeRosterByFieldResponseBodyResultFieldDataList) String() string {
	return tea.Prettify(s)
}

func (s GetEmployeeRosterByFieldResponseBodyResultFieldDataList) GoString() string {
	return s.String()
}

func (s *GetEmployeeRosterByFieldResponseBodyResultFieldDataList) SetFieldCode(v string) *GetEmployeeRosterByFieldResponseBodyResultFieldDataList {
	s.FieldCode = &v
	return s
}

func (s *GetEmployeeRosterByFieldResponseBodyResultFieldDataList) SetFieldName(v string) *GetEmployeeRosterByFieldResponseBodyResultFieldDataList {
	s.FieldName = &v
	return s
}

func (s *GetEmployeeRosterByFieldResponseBodyResultFieldDataList) SetFieldValueList(v []*GetEmployeeRosterByFieldResponseBodyResultFieldDataListFieldValueList) *GetEmployeeRosterByFieldResponseBodyResultFieldDataList {
	s.FieldValueList = v
	return s
}

func (s *GetEmployeeRosterByFieldResponseBodyResultFieldDataList) SetGroupId(v string) *GetEmployeeRosterByFieldResponseBodyResultFieldDataList {
	s.GroupId = &v
	return s
}

type GetEmployeeRosterByFieldResponseBodyResultFieldDataListFieldValueList struct {
	// example:
	//
	// 0
	ItemIndex *int32 `json:"itemIndex,omitempty" xml:"itemIndex,omitempty"`
	// example:
	//
	// 正式
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// example:
	//
	// 3
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetEmployeeRosterByFieldResponseBodyResultFieldDataListFieldValueList) String() string {
	return tea.Prettify(s)
}

func (s GetEmployeeRosterByFieldResponseBodyResultFieldDataListFieldValueList) GoString() string {
	return s.String()
}

func (s *GetEmployeeRosterByFieldResponseBodyResultFieldDataListFieldValueList) SetItemIndex(v int32) *GetEmployeeRosterByFieldResponseBodyResultFieldDataListFieldValueList {
	s.ItemIndex = &v
	return s
}

func (s *GetEmployeeRosterByFieldResponseBodyResultFieldDataListFieldValueList) SetLabel(v string) *GetEmployeeRosterByFieldResponseBodyResultFieldDataListFieldValueList {
	s.Label = &v
	return s
}

func (s *GetEmployeeRosterByFieldResponseBodyResultFieldDataListFieldValueList) SetValue(v string) *GetEmployeeRosterByFieldResponseBodyResultFieldDataListFieldValueList {
	s.Value = &v
	return s
}

type GetEmployeeRosterByFieldResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEmployeeRosterByFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEmployeeRosterByFieldResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEmployeeRosterByFieldResponse) GoString() string {
	return s.String()
}

func (s *GetEmployeeRosterByFieldResponse) SetHeaders(v map[string]*string) *GetEmployeeRosterByFieldResponse {
	s.Headers = v
	return s
}

func (s *GetEmployeeRosterByFieldResponse) SetStatusCode(v int32) *GetEmployeeRosterByFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEmployeeRosterByFieldResponse) SetBody(v *GetEmployeeRosterByFieldResponseBody) *GetEmployeeRosterByFieldResponse {
	s.Body = v
	return s
}

type GetFileTemplateListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFileTemplateListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFileTemplateListHeaders) GoString() string {
	return s.String()
}

func (s *GetFileTemplateListHeaders) SetCommonHeaders(v map[string]*string) *GetFileTemplateListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFileTemplateListHeaders) SetXAcsDingtalkAccessToken(v string) *GetFileTemplateListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFileTemplateListRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 0
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CONTRACT
	SignSource *string `json:"signSource,omitempty" xml:"signSource,omitempty"`
	// example:
	//
	// 1
	TemplateStatus   *int32    `json:"templateStatus,omitempty" xml:"templateStatus,omitempty"`
	TemplateTypeList []*string `json:"templateTypeList,omitempty" xml:"templateTypeList,omitempty" type:"Repeated"`
}

func (s GetFileTemplateListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFileTemplateListRequest) GoString() string {
	return s.String()
}

func (s *GetFileTemplateListRequest) SetMaxResults(v int32) *GetFileTemplateListRequest {
	s.MaxResults = &v
	return s
}

func (s *GetFileTemplateListRequest) SetNextToken(v int64) *GetFileTemplateListRequest {
	s.NextToken = &v
	return s
}

func (s *GetFileTemplateListRequest) SetSignSource(v string) *GetFileTemplateListRequest {
	s.SignSource = &v
	return s
}

func (s *GetFileTemplateListRequest) SetTemplateStatus(v int32) *GetFileTemplateListRequest {
	s.TemplateStatus = &v
	return s
}

func (s *GetFileTemplateListRequest) SetTemplateTypeList(v []*string) *GetFileTemplateListRequest {
	s.TemplateTypeList = v
	return s
}

type GetFileTemplateListResponseBody struct {
	Result  *GetFileTemplateListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetFileTemplateListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFileTemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileTemplateListResponseBody) SetResult(v *GetFileTemplateListResponseBodyResult) *GetFileTemplateListResponseBody {
	s.Result = v
	return s
}

func (s *GetFileTemplateListResponseBody) SetSuccess(v bool) *GetFileTemplateListResponseBody {
	s.Success = &v
	return s
}

type GetFileTemplateListResponseBodyResult struct {
	Data []*GetFileTemplateListResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// This parameter is required.
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 10
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetFileTemplateListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFileTemplateListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFileTemplateListResponseBodyResult) SetData(v []*GetFileTemplateListResponseBodyResultData) *GetFileTemplateListResponseBodyResult {
	s.Data = v
	return s
}

func (s *GetFileTemplateListResponseBodyResult) SetHasMore(v bool) *GetFileTemplateListResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResult) SetNextToken(v int64) *GetFileTemplateListResponseBodyResult {
	s.NextToken = &v
	return s
}

type GetFileTemplateListResponseBodyResultData struct {
	AttachmentList []*GetFileTemplateListResponseBodyResultDataAttachmentList `json:"attachmentList,omitempty" xml:"attachmentList,omitempty" type:"Repeated"`
	// example:
	//
	// ding57935b18bfd13e9735c2f4657eb6378f
	CorpId    *string                                               `json:"corpId,omitempty" xml:"corpId,omitempty"`
	FieldList []*GetFileTemplateListResponseBodyResultDataFieldList `json:"fieldList,omitempty" xml:"fieldList,omitempty" type:"Repeated"`
	GroupList []*GetFileTemplateListResponseBodyResultDataGroupList `json:"groupList,omitempty" xml:"groupList,omitempty" type:"Repeated"`
	// example:
	//
	// f3ed5402e3024febaed773b3ac19feb3
	TemplateId       *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	TemplateInstName *string `json:"templateInstName,omitempty" xml:"templateInstName,omitempty"`
	// example:
	//
	// 劳动合同模板
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// example:
	//
	// 1
	TemplateSignStatus *int32 `json:"templateSignStatus,omitempty" xml:"templateSignStatus,omitempty"`
	// example:
	//
	// 1
	TemplateStatus *int32 `json:"templateStatus,omitempty" xml:"templateStatus,omitempty"`
	// example:
	//
	// CONTRACT
	TemplateType *string `json:"templateType,omitempty" xml:"templateType,omitempty"`
	// example:
	//
	// 合同模板
	TemplateTypeName *string `json:"templateTypeName,omitempty" xml:"templateTypeName,omitempty"`
	// example:
	//
	// 24
	TenantId *int64 `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetFileTemplateListResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s GetFileTemplateListResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *GetFileTemplateListResponseBodyResultData) SetAttachmentList(v []*GetFileTemplateListResponseBodyResultDataAttachmentList) *GetFileTemplateListResponseBodyResultData {
	s.AttachmentList = v
	return s
}

func (s *GetFileTemplateListResponseBodyResultData) SetCorpId(v string) *GetFileTemplateListResponseBodyResultData {
	s.CorpId = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultData) SetFieldList(v []*GetFileTemplateListResponseBodyResultDataFieldList) *GetFileTemplateListResponseBodyResultData {
	s.FieldList = v
	return s
}

func (s *GetFileTemplateListResponseBodyResultData) SetGroupList(v []*GetFileTemplateListResponseBodyResultDataGroupList) *GetFileTemplateListResponseBodyResultData {
	s.GroupList = v
	return s
}

func (s *GetFileTemplateListResponseBodyResultData) SetTemplateId(v string) *GetFileTemplateListResponseBodyResultData {
	s.TemplateId = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultData) SetTemplateInstName(v string) *GetFileTemplateListResponseBodyResultData {
	s.TemplateInstName = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultData) SetTemplateName(v string) *GetFileTemplateListResponseBodyResultData {
	s.TemplateName = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultData) SetTemplateSignStatus(v int32) *GetFileTemplateListResponseBodyResultData {
	s.TemplateSignStatus = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultData) SetTemplateStatus(v int32) *GetFileTemplateListResponseBodyResultData {
	s.TemplateStatus = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultData) SetTemplateType(v string) *GetFileTemplateListResponseBodyResultData {
	s.TemplateType = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultData) SetTemplateTypeName(v string) *GetFileTemplateListResponseBodyResultData {
	s.TemplateTypeName = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultData) SetTenantId(v int64) *GetFileTemplateListResponseBodyResultData {
	s.TenantId = &v
	return s
}

type GetFileTemplateListResponseBodyResultDataAttachmentList struct {
	// example:
	//
	// 简历附件
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// attachment_profile
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	// example:
	//
	// 简历附件
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// example:
	//
	// DDAttachmentField
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	// example:
	//
	// attachment
	GroupId         *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	SignRequired    *bool   `json:"signRequired,omitempty" xml:"signRequired,omitempty"`
	UserCustomField *bool   `json:"userCustomField,omitempty" xml:"userCustomField,omitempty"`
}

func (s GetFileTemplateListResponseBodyResultDataAttachmentList) String() string {
	return tea.Prettify(s)
}

func (s GetFileTemplateListResponseBodyResultDataAttachmentList) GoString() string {
	return s.String()
}

func (s *GetFileTemplateListResponseBodyResultDataAttachmentList) SetDesc(v string) *GetFileTemplateListResponseBodyResultDataAttachmentList {
	s.Desc = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultDataAttachmentList) SetFieldCode(v string) *GetFileTemplateListResponseBodyResultDataAttachmentList {
	s.FieldCode = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultDataAttachmentList) SetFieldName(v string) *GetFileTemplateListResponseBodyResultDataAttachmentList {
	s.FieldName = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultDataAttachmentList) SetFieldType(v string) *GetFileTemplateListResponseBodyResultDataAttachmentList {
	s.FieldType = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultDataAttachmentList) SetGroupId(v string) *GetFileTemplateListResponseBodyResultDataAttachmentList {
	s.GroupId = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultDataAttachmentList) SetSignRequired(v bool) *GetFileTemplateListResponseBodyResultDataAttachmentList {
	s.SignRequired = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultDataAttachmentList) SetUserCustomField(v bool) *GetFileTemplateListResponseBodyResultDataAttachmentList {
	s.UserCustomField = &v
	return s
}

type GetFileTemplateListResponseBodyResultDataFieldList struct {
	// example:
	//
	// 真实姓名字段
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// esign_name
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	// example:
	//
	// 真实姓名
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// example:
	//
	// TextField
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	// example:
	//
	// baseInfo
	GroupId         *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	SignRequired    *bool   `json:"signRequired,omitempty" xml:"signRequired,omitempty"`
	UserCustomField *bool   `json:"userCustomField,omitempty" xml:"userCustomField,omitempty"`
}

func (s GetFileTemplateListResponseBodyResultDataFieldList) String() string {
	return tea.Prettify(s)
}

func (s GetFileTemplateListResponseBodyResultDataFieldList) GoString() string {
	return s.String()
}

func (s *GetFileTemplateListResponseBodyResultDataFieldList) SetDesc(v string) *GetFileTemplateListResponseBodyResultDataFieldList {
	s.Desc = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultDataFieldList) SetFieldCode(v string) *GetFileTemplateListResponseBodyResultDataFieldList {
	s.FieldCode = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultDataFieldList) SetFieldName(v string) *GetFileTemplateListResponseBodyResultDataFieldList {
	s.FieldName = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultDataFieldList) SetFieldType(v string) *GetFileTemplateListResponseBodyResultDataFieldList {
	s.FieldType = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultDataFieldList) SetGroupId(v string) *GetFileTemplateListResponseBodyResultDataFieldList {
	s.GroupId = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultDataFieldList) SetSignRequired(v bool) *GetFileTemplateListResponseBodyResultDataFieldList {
	s.SignRequired = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultDataFieldList) SetUserCustomField(v bool) *GetFileTemplateListResponseBodyResultDataFieldList {
	s.UserCustomField = &v
	return s
}

type GetFileTemplateListResponseBodyResultDataGroupList struct {
	DetailFlag *bool                                                          `json:"detailFlag,omitempty" xml:"detailFlag,omitempty"`
	FieldList  []*GetFileTemplateListResponseBodyResultDataGroupListFieldList `json:"fieldList,omitempty" xml:"fieldList,omitempty" type:"Repeated"`
	// example:
	//
	// family
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// 家庭成员
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
}

func (s GetFileTemplateListResponseBodyResultDataGroupList) String() string {
	return tea.Prettify(s)
}

func (s GetFileTemplateListResponseBodyResultDataGroupList) GoString() string {
	return s.String()
}

func (s *GetFileTemplateListResponseBodyResultDataGroupList) SetDetailFlag(v bool) *GetFileTemplateListResponseBodyResultDataGroupList {
	s.DetailFlag = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultDataGroupList) SetFieldList(v []*GetFileTemplateListResponseBodyResultDataGroupListFieldList) *GetFileTemplateListResponseBodyResultDataGroupList {
	s.FieldList = v
	return s
}

func (s *GetFileTemplateListResponseBodyResultDataGroupList) SetGroupId(v string) *GetFileTemplateListResponseBodyResultDataGroupList {
	s.GroupId = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultDataGroupList) SetGroupName(v string) *GetFileTemplateListResponseBodyResultDataGroupList {
	s.GroupName = &v
	return s
}

type GetFileTemplateListResponseBodyResultDataGroupListFieldList struct {
	// example:
	//
	// 家庭成员明细分组
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// family.member_name
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	// example:
	//
	// 成员姓名
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// example:
	//
	// TextField
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	// example:
	//
	// family
	GroupId         *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	SignRequired    *bool   `json:"signRequired,omitempty" xml:"signRequired,omitempty"`
	UserCustomField *bool   `json:"userCustomField,omitempty" xml:"userCustomField,omitempty"`
}

func (s GetFileTemplateListResponseBodyResultDataGroupListFieldList) String() string {
	return tea.Prettify(s)
}

func (s GetFileTemplateListResponseBodyResultDataGroupListFieldList) GoString() string {
	return s.String()
}

func (s *GetFileTemplateListResponseBodyResultDataGroupListFieldList) SetDesc(v string) *GetFileTemplateListResponseBodyResultDataGroupListFieldList {
	s.Desc = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultDataGroupListFieldList) SetFieldCode(v string) *GetFileTemplateListResponseBodyResultDataGroupListFieldList {
	s.FieldCode = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultDataGroupListFieldList) SetFieldName(v string) *GetFileTemplateListResponseBodyResultDataGroupListFieldList {
	s.FieldName = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultDataGroupListFieldList) SetFieldType(v string) *GetFileTemplateListResponseBodyResultDataGroupListFieldList {
	s.FieldType = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultDataGroupListFieldList) SetGroupId(v string) *GetFileTemplateListResponseBodyResultDataGroupListFieldList {
	s.GroupId = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultDataGroupListFieldList) SetSignRequired(v bool) *GetFileTemplateListResponseBodyResultDataGroupListFieldList {
	s.SignRequired = &v
	return s
}

func (s *GetFileTemplateListResponseBodyResultDataGroupListFieldList) SetUserCustomField(v bool) *GetFileTemplateListResponseBodyResultDataGroupListFieldList {
	s.UserCustomField = &v
	return s
}

type GetFileTemplateListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileTemplateListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFileTemplateListResponse) GoString() string {
	return s.String()
}

func (s *GetFileTemplateListResponse) SetHeaders(v map[string]*string) *GetFileTemplateListResponse {
	s.Headers = v
	return s
}

func (s *GetFileTemplateListResponse) SetStatusCode(v int32) *GetFileTemplateListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileTemplateListResponse) SetBody(v *GetFileTemplateListResponseBody) *GetFileTemplateListResponse {
	s.Body = v
	return s
}

type GetSignRecordByIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSignRecordByIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSignRecordByIdHeaders) GoString() string {
	return s.String()
}

func (s *GetSignRecordByIdHeaders) SetCommonHeaders(v map[string]*string) *GetSignRecordByIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSignRecordByIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetSignRecordByIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSignRecordByIdRequest struct {
	// This parameter is required.
	Body []*string `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s GetSignRecordByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSignRecordByIdRequest) GoString() string {
	return s.String()
}

func (s *GetSignRecordByIdRequest) SetBody(v []*string) *GetSignRecordByIdRequest {
	s.Body = v
	return s
}

type GetSignRecordByIdResponseBody struct {
	Result []*GetSignRecordByIdResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// This parameter is required.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetSignRecordByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSignRecordByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetSignRecordByIdResponseBody) SetResult(v []*GetSignRecordByIdResponseBodyResult) *GetSignRecordByIdResponseBody {
	s.Result = v
	return s
}

func (s *GetSignRecordByIdResponseBody) SetSuccess(v bool) *GetSignRecordByIdResponseBody {
	s.Success = &v
	return s
}

type GetSignRecordByIdResponseBodyResult struct {
	// example:
	//
	// ding12345678
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// 员工签署中
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// 1723534265000
	SignExpireTime *int64 `json:"signExpireTime,omitempty" xml:"signExpireTime,omitempty"`
	// example:
	//
	// xxxx-合同文件.pdf
	SignFileName *string `json:"signFileName,omitempty" xml:"signFileName,omitempty"`
	// example:
	//
	// 1723534265000
	SignFinishTime *int64 `json:"signFinishTime,omitempty" xml:"signFinishTime,omitempty"`
	// example:
	//
	// xxxx有限公司
	SignLegalEntityName *string `json:"signLegalEntityName,omitempty" xml:"signLegalEntityName,omitempty"`
	// example:
	//
	// 6fe06f57ab5a45078f3219be8fd829c6
	SignRecordId *string `json:"signRecordId,omitempty" xml:"signRecordId,omitempty"`
	// example:
	//
	// 1723534265000
	SignStartTime *int64 `json:"signStartTime,omitempty" xml:"signStartTime,omitempty"`
	// example:
	//
	// SIGNING
	SignStatus *string `json:"signStatus,omitempty" xml:"signStatus,omitempty"`
	// example:
	//
	// 签署中
	SignStatusRemarks *string `json:"signStatusRemarks,omitempty" xml:"signStatusRemarks,omitempty"`
	// example:
	//
	// CONTRACT
	SignTemplateType *string `json:"signTemplateType,omitempty" xml:"signTemplateType,omitempty"`
	// example:
	//
	// 400873120394
	SignUserId *string `json:"signUserId,omitempty" xml:"signUserId,omitempty"`
	// example:
	//
	// xiaoming
	SignUserName *string `json:"signUserName,omitempty" xml:"signUserName,omitempty"`
	// example:
	//
	// ON_LINE
	SignWay *string `json:"signWay,omitempty" xml:"signWay,omitempty"`
}

func (s GetSignRecordByIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetSignRecordByIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSignRecordByIdResponseBodyResult) SetCorpId(v string) *GetSignRecordByIdResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *GetSignRecordByIdResponseBodyResult) SetRemark(v string) *GetSignRecordByIdResponseBodyResult {
	s.Remark = &v
	return s
}

func (s *GetSignRecordByIdResponseBodyResult) SetSignExpireTime(v int64) *GetSignRecordByIdResponseBodyResult {
	s.SignExpireTime = &v
	return s
}

func (s *GetSignRecordByIdResponseBodyResult) SetSignFileName(v string) *GetSignRecordByIdResponseBodyResult {
	s.SignFileName = &v
	return s
}

func (s *GetSignRecordByIdResponseBodyResult) SetSignFinishTime(v int64) *GetSignRecordByIdResponseBodyResult {
	s.SignFinishTime = &v
	return s
}

func (s *GetSignRecordByIdResponseBodyResult) SetSignLegalEntityName(v string) *GetSignRecordByIdResponseBodyResult {
	s.SignLegalEntityName = &v
	return s
}

func (s *GetSignRecordByIdResponseBodyResult) SetSignRecordId(v string) *GetSignRecordByIdResponseBodyResult {
	s.SignRecordId = &v
	return s
}

func (s *GetSignRecordByIdResponseBodyResult) SetSignStartTime(v int64) *GetSignRecordByIdResponseBodyResult {
	s.SignStartTime = &v
	return s
}

func (s *GetSignRecordByIdResponseBodyResult) SetSignStatus(v string) *GetSignRecordByIdResponseBodyResult {
	s.SignStatus = &v
	return s
}

func (s *GetSignRecordByIdResponseBodyResult) SetSignStatusRemarks(v string) *GetSignRecordByIdResponseBodyResult {
	s.SignStatusRemarks = &v
	return s
}

func (s *GetSignRecordByIdResponseBodyResult) SetSignTemplateType(v string) *GetSignRecordByIdResponseBodyResult {
	s.SignTemplateType = &v
	return s
}

func (s *GetSignRecordByIdResponseBodyResult) SetSignUserId(v string) *GetSignRecordByIdResponseBodyResult {
	s.SignUserId = &v
	return s
}

func (s *GetSignRecordByIdResponseBodyResult) SetSignUserName(v string) *GetSignRecordByIdResponseBodyResult {
	s.SignUserName = &v
	return s
}

func (s *GetSignRecordByIdResponseBodyResult) SetSignWay(v string) *GetSignRecordByIdResponseBodyResult {
	s.SignWay = &v
	return s
}

type GetSignRecordByIdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSignRecordByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSignRecordByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSignRecordByIdResponse) GoString() string {
	return s.String()
}

func (s *GetSignRecordByIdResponse) SetHeaders(v map[string]*string) *GetSignRecordByIdResponse {
	s.Headers = v
	return s
}

func (s *GetSignRecordByIdResponse) SetStatusCode(v int32) *GetSignRecordByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSignRecordByIdResponse) SetBody(v *GetSignRecordByIdResponseBody) *GetSignRecordByIdResponse {
	s.Body = v
	return s
}

type GetSignRecordByUserIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSignRecordByUserIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSignRecordByUserIdHeaders) GoString() string {
	return s.String()
}

func (s *GetSignRecordByUserIdHeaders) SetCommonHeaders(v map[string]*string) *GetSignRecordByUserIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSignRecordByUserIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetSignRecordByUserIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSignRecordByUserIdRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 0
	NextToken  *int64    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	SignStatus []*string `json:"signStatus,omitempty" xml:"signStatus,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 660658
	SignUserId *string `json:"signUserId,omitempty" xml:"signUserId,omitempty"`
}

func (s GetSignRecordByUserIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSignRecordByUserIdRequest) GoString() string {
	return s.String()
}

func (s *GetSignRecordByUserIdRequest) SetMaxResults(v int32) *GetSignRecordByUserIdRequest {
	s.MaxResults = &v
	return s
}

func (s *GetSignRecordByUserIdRequest) SetNextToken(v int64) *GetSignRecordByUserIdRequest {
	s.NextToken = &v
	return s
}

func (s *GetSignRecordByUserIdRequest) SetSignStatus(v []*string) *GetSignRecordByUserIdRequest {
	s.SignStatus = v
	return s
}

func (s *GetSignRecordByUserIdRequest) SetSignUserId(v string) *GetSignRecordByUserIdRequest {
	s.SignUserId = &v
	return s
}

type GetSignRecordByUserIdResponseBody struct {
	Result  *GetSignRecordByUserIdResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                    `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetSignRecordByUserIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSignRecordByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetSignRecordByUserIdResponseBody) SetResult(v *GetSignRecordByUserIdResponseBodyResult) *GetSignRecordByUserIdResponseBody {
	s.Result = v
	return s
}

func (s *GetSignRecordByUserIdResponseBody) SetSuccess(v bool) *GetSignRecordByUserIdResponseBody {
	s.Success = &v
	return s
}

type GetSignRecordByUserIdResponseBodyResult struct {
	Data      []*GetSignRecordByUserIdResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	HasMore   *bool                                          `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken *int64                                         `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetSignRecordByUserIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetSignRecordByUserIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSignRecordByUserIdResponseBodyResult) SetData(v []*GetSignRecordByUserIdResponseBodyResultData) *GetSignRecordByUserIdResponseBodyResult {
	s.Data = v
	return s
}

func (s *GetSignRecordByUserIdResponseBodyResult) SetHasMore(v bool) *GetSignRecordByUserIdResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *GetSignRecordByUserIdResponseBodyResult) SetNextToken(v int64) *GetSignRecordByUserIdResponseBodyResult {
	s.NextToken = &v
	return s
}

type GetSignRecordByUserIdResponseBodyResultData struct {
	// example:
	//
	// ding57935b18bfd13e9735c2f4657eb6378f
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// CONTRACT_123456
	OuterId *string `json:"outerId,omitempty" xml:"outerId,omitempty"`
	// example:
	//
	// 劳动合同电子签签署备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// 1720775436000
	SignExpireTime *int64 `json:"signExpireTime,omitempty" xml:"signExpireTime,omitempty"`
	// example:
	//
	// 小明-劳动合同-20240808.pdf
	SignFileName *string `json:"signFileName,omitempty" xml:"signFileName,omitempty"`
	// example:
	//
	// https://n.dingtalk.com/xxx
	SignFileUrl *string `json:"signFileUrl,omitempty" xml:"signFileUrl,omitempty"`
	// example:
	//
	// 1720775436000
	SignFinishTime *int64 `json:"signFinishTime,omitempty" xml:"signFinishTime,omitempty"`
	// example:
	//
	// xxx有限公司
	SignLegalEntityName *string `json:"signLegalEntityName,omitempty" xml:"signLegalEntityName,omitempty"`
	// example:
	//
	// 38bc7b69bb6a46b8a69b9001d5c0bdf3
	SignRecordId *string `json:"signRecordId,omitempty" xml:"signRecordId,omitempty"`
	// example:
	//
	// 1720775436000
	SignStartTime *int64 `json:"signStartTime,omitempty" xml:"signStartTime,omitempty"`
	// example:
	//
	// FINISHED
	SignStatus *string `json:"signStatus,omitempty" xml:"signStatus,omitempty"`
	// example:
	//
	// 法人公司未开通
	SignStatusRemarks *string `json:"signStatusRemarks,omitempty" xml:"signStatusRemarks,omitempty"`
	// example:
	//
	// CONTRACT
	SignTemplateType *string `json:"signTemplateType,omitempty" xml:"signTemplateType,omitempty"`
	// example:
	//
	// 660658
	SignUserId *string `json:"signUserId,omitempty" xml:"signUserId,omitempty"`
	// example:
	//
	// 小明
	SignUserName *string `json:"signUserName,omitempty" xml:"signUserName,omitempty"`
	// example:
	//
	// ON_LINE
	SignWay *string `json:"signWay,omitempty" xml:"signWay,omitempty"`
}

func (s GetSignRecordByUserIdResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s GetSignRecordByUserIdResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *GetSignRecordByUserIdResponseBodyResultData) SetCorpId(v string) *GetSignRecordByUserIdResponseBodyResultData {
	s.CorpId = &v
	return s
}

func (s *GetSignRecordByUserIdResponseBodyResultData) SetOuterId(v string) *GetSignRecordByUserIdResponseBodyResultData {
	s.OuterId = &v
	return s
}

func (s *GetSignRecordByUserIdResponseBodyResultData) SetRemark(v string) *GetSignRecordByUserIdResponseBodyResultData {
	s.Remark = &v
	return s
}

func (s *GetSignRecordByUserIdResponseBodyResultData) SetSignExpireTime(v int64) *GetSignRecordByUserIdResponseBodyResultData {
	s.SignExpireTime = &v
	return s
}

func (s *GetSignRecordByUserIdResponseBodyResultData) SetSignFileName(v string) *GetSignRecordByUserIdResponseBodyResultData {
	s.SignFileName = &v
	return s
}

func (s *GetSignRecordByUserIdResponseBodyResultData) SetSignFileUrl(v string) *GetSignRecordByUserIdResponseBodyResultData {
	s.SignFileUrl = &v
	return s
}

func (s *GetSignRecordByUserIdResponseBodyResultData) SetSignFinishTime(v int64) *GetSignRecordByUserIdResponseBodyResultData {
	s.SignFinishTime = &v
	return s
}

func (s *GetSignRecordByUserIdResponseBodyResultData) SetSignLegalEntityName(v string) *GetSignRecordByUserIdResponseBodyResultData {
	s.SignLegalEntityName = &v
	return s
}

func (s *GetSignRecordByUserIdResponseBodyResultData) SetSignRecordId(v string) *GetSignRecordByUserIdResponseBodyResultData {
	s.SignRecordId = &v
	return s
}

func (s *GetSignRecordByUserIdResponseBodyResultData) SetSignStartTime(v int64) *GetSignRecordByUserIdResponseBodyResultData {
	s.SignStartTime = &v
	return s
}

func (s *GetSignRecordByUserIdResponseBodyResultData) SetSignStatus(v string) *GetSignRecordByUserIdResponseBodyResultData {
	s.SignStatus = &v
	return s
}

func (s *GetSignRecordByUserIdResponseBodyResultData) SetSignStatusRemarks(v string) *GetSignRecordByUserIdResponseBodyResultData {
	s.SignStatusRemarks = &v
	return s
}

func (s *GetSignRecordByUserIdResponseBodyResultData) SetSignTemplateType(v string) *GetSignRecordByUserIdResponseBodyResultData {
	s.SignTemplateType = &v
	return s
}

func (s *GetSignRecordByUserIdResponseBodyResultData) SetSignUserId(v string) *GetSignRecordByUserIdResponseBodyResultData {
	s.SignUserId = &v
	return s
}

func (s *GetSignRecordByUserIdResponseBodyResultData) SetSignUserName(v string) *GetSignRecordByUserIdResponseBodyResultData {
	s.SignUserName = &v
	return s
}

func (s *GetSignRecordByUserIdResponseBodyResultData) SetSignWay(v string) *GetSignRecordByUserIdResponseBodyResultData {
	s.SignWay = &v
	return s
}

type GetSignRecordByUserIdResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSignRecordByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSignRecordByUserIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSignRecordByUserIdResponse) GoString() string {
	return s.String()
}

func (s *GetSignRecordByUserIdResponse) SetHeaders(v map[string]*string) *GetSignRecordByUserIdResponse {
	s.Headers = v
	return s
}

func (s *GetSignRecordByUserIdResponse) SetStatusCode(v int32) *GetSignRecordByUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSignRecordByUserIdResponse) SetBody(v *GetSignRecordByUserIdResponseBody) *GetSignRecordByUserIdResponse {
	s.Body = v
	return s
}

type GetUserSignedRecordsByOuterIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserSignedRecordsByOuterIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserSignedRecordsByOuterIdHeaders) GoString() string {
	return s.String()
}

func (s *GetUserSignedRecordsByOuterIdHeaders) SetCommonHeaders(v map[string]*string) *GetUserSignedRecordsByOuterIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserSignedRecordsByOuterIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserSignedRecordsByOuterIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserSignedRecordsByOuterIdRequest struct {
	Body []*string `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s GetUserSignedRecordsByOuterIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserSignedRecordsByOuterIdRequest) GoString() string {
	return s.String()
}

func (s *GetUserSignedRecordsByOuterIdRequest) SetBody(v []*string) *GetUserSignedRecordsByOuterIdRequest {
	s.Body = v
	return s
}

type GetUserSignedRecordsByOuterIdResponseBody struct {
	Result  []*GetUserSignedRecordsByOuterIdResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success *bool                                              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetUserSignedRecordsByOuterIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserSignedRecordsByOuterIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserSignedRecordsByOuterIdResponseBody) SetResult(v []*GetUserSignedRecordsByOuterIdResponseBodyResult) *GetUserSignedRecordsByOuterIdResponseBody {
	s.Result = v
	return s
}

func (s *GetUserSignedRecordsByOuterIdResponseBody) SetSuccess(v bool) *GetUserSignedRecordsByOuterIdResponseBody {
	s.Success = &v
	return s
}

type GetUserSignedRecordsByOuterIdResponseBodyResult struct {
	// example:
	//
	// ding33a9d1a6e9647854a39a90f97fcb1e09
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// CONTRACT_123456
	OuterId *string `json:"outerId,omitempty" xml:"outerId,omitempty"`
	// example:
	//
	// 测试
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// 1720775436000
	SignExpireTime *string `json:"signExpireTime,omitempty" xml:"signExpireTime,omitempty"`
	// example:
	//
	// 小明-劳动合同-20240808.pdf
	SignFileName *string `json:"signFileName,omitempty" xml:"signFileName,omitempty"`
	// example:
	//
	// https://n.dingtalk.com/xxx
	SignFileUrl *string `json:"signFileUrl,omitempty" xml:"signFileUrl,omitempty"`
	// example:
	//
	// 1720775436000
	SignFinishTime *string `json:"signFinishTime,omitempty" xml:"signFinishTime,omitempty"`
	// example:
	//
	// xx公司
	SignLegalEntityName *string `json:"signLegalEntityName,omitempty" xml:"signLegalEntityName,omitempty"`
	// example:
	//
	// fb1a9c4adaba4f52b7cab7941008b9dd
	SignRecordId *string `json:"signRecordId,omitempty" xml:"signRecordId,omitempty"`
	// example:
	//
	// 1720775436000
	SignStartTime *string `json:"signStartTime,omitempty" xml:"signStartTime,omitempty"`
	// example:
	//
	// FINISHED
	SignStatus *string `json:"signStatus,omitempty" xml:"signStatus,omitempty"`
	// example:
	//
	// 法人公司未开通
	SignStatusRemarks *string `json:"signStatusRemarks,omitempty" xml:"signStatusRemarks,omitempty"`
	// example:
	//
	// CONTRACT
	SignTemplateType *string `json:"signTemplateType,omitempty" xml:"signTemplateType,omitempty"`
	// example:
	//
	// userId123
	SignUserId *string `json:"signUserId,omitempty" xml:"signUserId,omitempty"`
	// example:
	//
	// userName
	SignUserName *string `json:"signUserName,omitempty" xml:"signUserName,omitempty"`
	// example:
	//
	// ON_LINE
	SignWay *string `json:"signWay,omitempty" xml:"signWay,omitempty"`
}

func (s GetUserSignedRecordsByOuterIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetUserSignedRecordsByOuterIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetUserSignedRecordsByOuterIdResponseBodyResult) SetCorpId(v string) *GetUserSignedRecordsByOuterIdResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *GetUserSignedRecordsByOuterIdResponseBodyResult) SetOuterId(v string) *GetUserSignedRecordsByOuterIdResponseBodyResult {
	s.OuterId = &v
	return s
}

func (s *GetUserSignedRecordsByOuterIdResponseBodyResult) SetRemark(v string) *GetUserSignedRecordsByOuterIdResponseBodyResult {
	s.Remark = &v
	return s
}

func (s *GetUserSignedRecordsByOuterIdResponseBodyResult) SetSignExpireTime(v string) *GetUserSignedRecordsByOuterIdResponseBodyResult {
	s.SignExpireTime = &v
	return s
}

func (s *GetUserSignedRecordsByOuterIdResponseBodyResult) SetSignFileName(v string) *GetUserSignedRecordsByOuterIdResponseBodyResult {
	s.SignFileName = &v
	return s
}

func (s *GetUserSignedRecordsByOuterIdResponseBodyResult) SetSignFileUrl(v string) *GetUserSignedRecordsByOuterIdResponseBodyResult {
	s.SignFileUrl = &v
	return s
}

func (s *GetUserSignedRecordsByOuterIdResponseBodyResult) SetSignFinishTime(v string) *GetUserSignedRecordsByOuterIdResponseBodyResult {
	s.SignFinishTime = &v
	return s
}

func (s *GetUserSignedRecordsByOuterIdResponseBodyResult) SetSignLegalEntityName(v string) *GetUserSignedRecordsByOuterIdResponseBodyResult {
	s.SignLegalEntityName = &v
	return s
}

func (s *GetUserSignedRecordsByOuterIdResponseBodyResult) SetSignRecordId(v string) *GetUserSignedRecordsByOuterIdResponseBodyResult {
	s.SignRecordId = &v
	return s
}

func (s *GetUserSignedRecordsByOuterIdResponseBodyResult) SetSignStartTime(v string) *GetUserSignedRecordsByOuterIdResponseBodyResult {
	s.SignStartTime = &v
	return s
}

func (s *GetUserSignedRecordsByOuterIdResponseBodyResult) SetSignStatus(v string) *GetUserSignedRecordsByOuterIdResponseBodyResult {
	s.SignStatus = &v
	return s
}

func (s *GetUserSignedRecordsByOuterIdResponseBodyResult) SetSignStatusRemarks(v string) *GetUserSignedRecordsByOuterIdResponseBodyResult {
	s.SignStatusRemarks = &v
	return s
}

func (s *GetUserSignedRecordsByOuterIdResponseBodyResult) SetSignTemplateType(v string) *GetUserSignedRecordsByOuterIdResponseBodyResult {
	s.SignTemplateType = &v
	return s
}

func (s *GetUserSignedRecordsByOuterIdResponseBodyResult) SetSignUserId(v string) *GetUserSignedRecordsByOuterIdResponseBodyResult {
	s.SignUserId = &v
	return s
}

func (s *GetUserSignedRecordsByOuterIdResponseBodyResult) SetSignUserName(v string) *GetUserSignedRecordsByOuterIdResponseBodyResult {
	s.SignUserName = &v
	return s
}

func (s *GetUserSignedRecordsByOuterIdResponseBodyResult) SetSignWay(v string) *GetUserSignedRecordsByOuterIdResponseBodyResult {
	s.SignWay = &v
	return s
}

type GetUserSignedRecordsByOuterIdResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserSignedRecordsByOuterIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserSignedRecordsByOuterIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserSignedRecordsByOuterIdResponse) GoString() string {
	return s.String()
}

func (s *GetUserSignedRecordsByOuterIdResponse) SetHeaders(v map[string]*string) *GetUserSignedRecordsByOuterIdResponse {
	s.Headers = v
	return s
}

func (s *GetUserSignedRecordsByOuterIdResponse) SetStatusCode(v int32) *GetUserSignedRecordsByOuterIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserSignedRecordsByOuterIdResponse) SetBody(v *GetUserSignedRecordsByOuterIdResponseBody) *GetUserSignedRecordsByOuterIdResponse {
	s.Body = v
	return s
}

type HrmAuthResourcesQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrmAuthResourcesQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrmAuthResourcesQueryHeaders) GoString() string {
	return s.String()
}

func (s *HrmAuthResourcesQueryHeaders) SetCommonHeaders(v map[string]*string) *HrmAuthResourcesQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrmAuthResourcesQueryHeaders) SetXAcsDingtalkAccessToken(v string) *HrmAuthResourcesQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrmAuthResourcesQueryRequest struct {
	// This parameter is required.
	AuthResourceIds []*string `json:"authResourceIds,omitempty" xml:"authResourceIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1231
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s HrmAuthResourcesQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s HrmAuthResourcesQueryRequest) GoString() string {
	return s.String()
}

func (s *HrmAuthResourcesQueryRequest) SetAuthResourceIds(v []*string) *HrmAuthResourcesQueryRequest {
	s.AuthResourceIds = v
	return s
}

func (s *HrmAuthResourcesQueryRequest) SetUserId(v string) *HrmAuthResourcesQueryRequest {
	s.UserId = &v
	return s
}

type HrmAuthResourcesQueryResponseBody struct {
	Result []*HrmAuthResourcesQueryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s HrmAuthResourcesQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrmAuthResourcesQueryResponseBody) GoString() string {
	return s.String()
}

func (s *HrmAuthResourcesQueryResponseBody) SetResult(v []*HrmAuthResourcesQueryResponseBodyResult) *HrmAuthResourcesQueryResponseBody {
	s.Result = v
	return s
}

type HrmAuthResourcesQueryResponseBodyResult struct {
	Authorized *bool `json:"authorized,omitempty" xml:"authorized,omitempty"`
	// example:
	//
	// /signSetting/manage/*
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
}

func (s HrmAuthResourcesQueryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s HrmAuthResourcesQueryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *HrmAuthResourcesQueryResponseBodyResult) SetAuthorized(v bool) *HrmAuthResourcesQueryResponseBodyResult {
	s.Authorized = &v
	return s
}

func (s *HrmAuthResourcesQueryResponseBodyResult) SetResourceId(v string) *HrmAuthResourcesQueryResponseBodyResult {
	s.ResourceId = &v
	return s
}

type HrmAuthResourcesQueryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrmAuthResourcesQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrmAuthResourcesQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s HrmAuthResourcesQueryResponse) GoString() string {
	return s.String()
}

func (s *HrmAuthResourcesQueryResponse) SetHeaders(v map[string]*string) *HrmAuthResourcesQueryResponse {
	s.Headers = v
	return s
}

func (s *HrmAuthResourcesQueryResponse) SetStatusCode(v int32) *HrmAuthResourcesQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *HrmAuthResourcesQueryResponse) SetBody(v *HrmAuthResourcesQueryResponseBody) *HrmAuthResourcesQueryResponse {
	s.Body = v
	return s
}

type HrmBenefitQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrmBenefitQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrmBenefitQueryHeaders) GoString() string {
	return s.String()
}

func (s *HrmBenefitQueryHeaders) SetCommonHeaders(v map[string]*string) *HrmBenefitQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrmBenefitQueryHeaders) SetXAcsDingtalkAccessToken(v string) *HrmBenefitQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrmBenefitQueryRequest struct {
	// This parameter is required.
	BenefitCodes []*string `json:"benefitCodes,omitempty" xml:"benefitCodes,omitempty" type:"Repeated"`
}

func (s HrmBenefitQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s HrmBenefitQueryRequest) GoString() string {
	return s.String()
}

func (s *HrmBenefitQueryRequest) SetBenefitCodes(v []*string) *HrmBenefitQueryRequest {
	s.BenefitCodes = v
	return s
}

type HrmBenefitQueryResponseBody struct {
	Result interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s HrmBenefitQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrmBenefitQueryResponseBody) GoString() string {
	return s.String()
}

func (s *HrmBenefitQueryResponseBody) SetResult(v interface{}) *HrmBenefitQueryResponseBody {
	s.Result = v
	return s
}

type HrmBenefitQueryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrmBenefitQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrmBenefitQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s HrmBenefitQueryResponse) GoString() string {
	return s.String()
}

func (s *HrmBenefitQueryResponse) SetHeaders(v map[string]*string) *HrmBenefitQueryResponse {
	s.Headers = v
	return s
}

func (s *HrmBenefitQueryResponse) SetStatusCode(v int32) *HrmBenefitQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *HrmBenefitQueryResponse) SetBody(v *HrmBenefitQueryResponseBody) *HrmBenefitQueryResponse {
	s.Body = v
	return s
}

type HrmCorpConfigQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrmCorpConfigQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrmCorpConfigQueryHeaders) GoString() string {
	return s.String()
}

func (s *HrmCorpConfigQueryHeaders) SetCommonHeaders(v map[string]*string) *HrmCorpConfigQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrmCorpConfigQueryHeaders) SetXAcsDingtalkAccessToken(v string) *HrmCorpConfigQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrmCorpConfigQueryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// policy
	SubType *string `json:"subType,omitempty" xml:"subType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hrm_ai
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s HrmCorpConfigQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s HrmCorpConfigQueryRequest) GoString() string {
	return s.String()
}

func (s *HrmCorpConfigQueryRequest) SetSubType(v string) *HrmCorpConfigQueryRequest {
	s.SubType = &v
	return s
}

func (s *HrmCorpConfigQueryRequest) SetType(v string) *HrmCorpConfigQueryRequest {
	s.Type = &v
	return s
}

type HrmCorpConfigQueryResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s HrmCorpConfigQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrmCorpConfigQueryResponseBody) GoString() string {
	return s.String()
}

func (s *HrmCorpConfigQueryResponseBody) SetResult(v string) *HrmCorpConfigQueryResponseBody {
	s.Result = &v
	return s
}

type HrmCorpConfigQueryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrmCorpConfigQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrmCorpConfigQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s HrmCorpConfigQueryResponse) GoString() string {
	return s.String()
}

func (s *HrmCorpConfigQueryResponse) SetHeaders(v map[string]*string) *HrmCorpConfigQueryResponse {
	s.Headers = v
	return s
}

func (s *HrmCorpConfigQueryResponse) SetStatusCode(v int32) *HrmCorpConfigQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *HrmCorpConfigQueryResponse) SetBody(v *HrmCorpConfigQueryResponseBody) *HrmCorpConfigQueryResponse {
	s.Body = v
	return s
}

type HrmMailSendHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrmMailSendHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrmMailSendHeaders) GoString() string {
	return s.String()
}

func (s *HrmMailSendHeaders) SetCommonHeaders(v map[string]*string) *HrmMailSendHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrmMailSendHeaders) SetXAcsDingtalkAccessToken(v string) *HrmMailSendHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrmMailSendRequest struct {
	// This parameter is required.
	Mail *HrmMailSendRequestMail `json:"mail,omitempty" xml:"mail,omitempty" type:"Struct"`
	// This parameter is required.
	Operator *HrmMailSendRequestOperator `json:"operator,omitempty" xml:"operator,omitempty" type:"Struct"`
}

func (s HrmMailSendRequest) String() string {
	return tea.Prettify(s)
}

func (s HrmMailSendRequest) GoString() string {
	return s.String()
}

func (s *HrmMailSendRequest) SetMail(v *HrmMailSendRequestMail) *HrmMailSendRequest {
	s.Mail = v
	return s
}

func (s *HrmMailSendRequest) SetOperator(v *HrmMailSendRequestOperator) *HrmMailSendRequest {
	s.Operator = v
	return s
}

type HrmMailSendRequestMail struct {
	Attachments []*HrmMailSendRequestMailAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// example:
	//
	// abd@aaa.com,bcd@aaa.com,
	BccAddress *string `json:"bccAddress,omitempty" xml:"bccAddress,omitempty"`
	// example:
	//
	// abd@aaa.com,bcd@aaa.com,
	CcAddress *string `json:"ccAddress,omitempty" xml:"ccAddress,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 请及时填写请填写入职登记表
	Content *string                        `json:"content,omitempty" xml:"content,omitempty"`
	Meeting *HrmMailSendRequestMailMeeting `json:"meeting,omitempty" xml:"meeting,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// abd@aaa.com,bcd@aaa.com,
	ReceiverAddress *string `json:"receiverAddress,omitempty" xml:"receiverAddress,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 智能人事入职
	SenderAlias *string `json:"senderAlias,omitempty" xml:"senderAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 请填写入职登记表
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
}

func (s HrmMailSendRequestMail) String() string {
	return tea.Prettify(s)
}

func (s HrmMailSendRequestMail) GoString() string {
	return s.String()
}

func (s *HrmMailSendRequestMail) SetAttachments(v []*HrmMailSendRequestMailAttachments) *HrmMailSendRequestMail {
	s.Attachments = v
	return s
}

func (s *HrmMailSendRequestMail) SetBccAddress(v string) *HrmMailSendRequestMail {
	s.BccAddress = &v
	return s
}

func (s *HrmMailSendRequestMail) SetCcAddress(v string) *HrmMailSendRequestMail {
	s.CcAddress = &v
	return s
}

func (s *HrmMailSendRequestMail) SetContent(v string) *HrmMailSendRequestMail {
	s.Content = &v
	return s
}

func (s *HrmMailSendRequestMail) SetMeeting(v *HrmMailSendRequestMailMeeting) *HrmMailSendRequestMail {
	s.Meeting = v
	return s
}

func (s *HrmMailSendRequestMail) SetReceiverAddress(v string) *HrmMailSendRequestMail {
	s.ReceiverAddress = &v
	return s
}

func (s *HrmMailSendRequestMail) SetSenderAlias(v string) *HrmMailSendRequestMail {
	s.SenderAlias = &v
	return s
}

func (s *HrmMailSendRequestMail) SetSubject(v string) *HrmMailSendRequestMail {
	s.Subject = &v
	return s
}

type HrmMailSendRequestMailAttachments struct {
	// This parameter is required.
	//
	// example:
	//
	// 测试.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// @asdc12312
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// media
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s HrmMailSendRequestMailAttachments) String() string {
	return tea.Prettify(s)
}

func (s HrmMailSendRequestMailAttachments) GoString() string {
	return s.String()
}

func (s *HrmMailSendRequestMailAttachments) SetName(v string) *HrmMailSendRequestMailAttachments {
	s.Name = &v
	return s
}

func (s *HrmMailSendRequestMailAttachments) SetPath(v string) *HrmMailSendRequestMailAttachments {
	s.Path = &v
	return s
}

func (s *HrmMailSendRequestMailAttachments) SetType(v string) *HrmMailSendRequestMailAttachments {
	s.Type = &v
	return s
}

type HrmMailSendRequestMailMeeting struct {
	Alarm     *HrmMailSendRequestMailMeetingAlarm       `json:"alarm,omitempty" xml:"alarm,omitempty" type:"Struct"`
	Attendees []*HrmMailSendRequestMailMeetingAttendees `json:"attendees,omitempty" xml:"attendees,omitempty" type:"Repeated"`
	// example:
	//
	// 会议描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1701692590881
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 会议室
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// REQUEST
	Method    *string                                 `json:"method,omitempty" xml:"method,omitempty"`
	Organizer *HrmMailSendRequestMailMeetingOrganizer `json:"organizer,omitempty" xml:"organizer,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Sequence *int32 `json:"sequence,omitempty" xml:"sequence,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1701692590881
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试会议
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// uuidssssss
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s HrmMailSendRequestMailMeeting) String() string {
	return tea.Prettify(s)
}

func (s HrmMailSendRequestMailMeeting) GoString() string {
	return s.String()
}

func (s *HrmMailSendRequestMailMeeting) SetAlarm(v *HrmMailSendRequestMailMeetingAlarm) *HrmMailSendRequestMailMeeting {
	s.Alarm = v
	return s
}

func (s *HrmMailSendRequestMailMeeting) SetAttendees(v []*HrmMailSendRequestMailMeetingAttendees) *HrmMailSendRequestMailMeeting {
	s.Attendees = v
	return s
}

func (s *HrmMailSendRequestMailMeeting) SetDescription(v string) *HrmMailSendRequestMailMeeting {
	s.Description = &v
	return s
}

func (s *HrmMailSendRequestMailMeeting) SetEndTime(v int64) *HrmMailSendRequestMailMeeting {
	s.EndTime = &v
	return s
}

func (s *HrmMailSendRequestMailMeeting) SetLocation(v string) *HrmMailSendRequestMailMeeting {
	s.Location = &v
	return s
}

func (s *HrmMailSendRequestMailMeeting) SetMethod(v string) *HrmMailSendRequestMailMeeting {
	s.Method = &v
	return s
}

func (s *HrmMailSendRequestMailMeeting) SetOrganizer(v *HrmMailSendRequestMailMeetingOrganizer) *HrmMailSendRequestMailMeeting {
	s.Organizer = v
	return s
}

func (s *HrmMailSendRequestMailMeeting) SetSequence(v int32) *HrmMailSendRequestMailMeeting {
	s.Sequence = &v
	return s
}

func (s *HrmMailSendRequestMailMeeting) SetStartTime(v int64) *HrmMailSendRequestMailMeeting {
	s.StartTime = &v
	return s
}

func (s *HrmMailSendRequestMailMeeting) SetSummary(v string) *HrmMailSendRequestMailMeeting {
	s.Summary = &v
	return s
}

func (s *HrmMailSendRequestMailMeeting) SetUuid(v string) *HrmMailSendRequestMailMeeting {
	s.Uuid = &v
	return s
}

type HrmMailSendRequestMailMeetingAlarm struct {
	// This parameter is required.
	//
	// example:
	//
	// 还有10分钟开始
	AlarmDesc *string `json:"alarmDesc,omitempty" xml:"alarmDesc,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	AlarmMinutes *int32 `json:"alarmMinutes,omitempty" xml:"alarmMinutes,omitempty"`
	// This parameter is required.
	AlarmSummary *string `json:"alarmSummary,omitempty" xml:"alarmSummary,omitempty"`
}

func (s HrmMailSendRequestMailMeetingAlarm) String() string {
	return tea.Prettify(s)
}

func (s HrmMailSendRequestMailMeetingAlarm) GoString() string {
	return s.String()
}

func (s *HrmMailSendRequestMailMeetingAlarm) SetAlarmDesc(v string) *HrmMailSendRequestMailMeetingAlarm {
	s.AlarmDesc = &v
	return s
}

func (s *HrmMailSendRequestMailMeetingAlarm) SetAlarmMinutes(v int32) *HrmMailSendRequestMailMeetingAlarm {
	s.AlarmMinutes = &v
	return s
}

func (s *HrmMailSendRequestMailMeetingAlarm) SetAlarmSummary(v string) *HrmMailSendRequestMailMeetingAlarm {
	s.AlarmSummary = &v
	return s
}

type HrmMailSendRequestMailMeetingAttendees struct {
	// This parameter is required.
	//
	// example:
	//
	// abc@aaa.com
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 参会人1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s HrmMailSendRequestMailMeetingAttendees) String() string {
	return tea.Prettify(s)
}

func (s HrmMailSendRequestMailMeetingAttendees) GoString() string {
	return s.String()
}

func (s *HrmMailSendRequestMailMeetingAttendees) SetAddress(v string) *HrmMailSendRequestMailMeetingAttendees {
	s.Address = &v
	return s
}

func (s *HrmMailSendRequestMailMeetingAttendees) SetName(v string) *HrmMailSendRequestMailMeetingAttendees {
	s.Name = &v
	return s
}

type HrmMailSendRequestMailMeetingOrganizer struct {
	// This parameter is required.
	//
	// example:
	//
	// abc@aaa.com
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 系统
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s HrmMailSendRequestMailMeetingOrganizer) String() string {
	return tea.Prettify(s)
}

func (s HrmMailSendRequestMailMeetingOrganizer) GoString() string {
	return s.String()
}

func (s *HrmMailSendRequestMailMeetingOrganizer) SetAddress(v string) *HrmMailSendRequestMailMeetingOrganizer {
	s.Address = &v
	return s
}

func (s *HrmMailSendRequestMailMeetingOrganizer) SetName(v string) *HrmMailSendRequestMailMeetingOrganizer {
	s.Name = &v
	return s
}

type HrmMailSendRequestOperator struct {
	// This parameter is required.
	//
	// example:
	//
	// biz222ddd
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hrm
	MailAccountType *string `json:"mailAccountType,omitempty" xml:"mailAccountType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// tokenabcd
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s HrmMailSendRequestOperator) String() string {
	return tea.Prettify(s)
}

func (s HrmMailSendRequestOperator) GoString() string {
	return s.String()
}

func (s *HrmMailSendRequestOperator) SetBizId(v string) *HrmMailSendRequestOperator {
	s.BizId = &v
	return s
}

func (s *HrmMailSendRequestOperator) SetMailAccountType(v string) *HrmMailSendRequestOperator {
	s.MailAccountType = &v
	return s
}

func (s *HrmMailSendRequestOperator) SetToken(v string) *HrmMailSendRequestOperator {
	s.Token = &v
	return s
}

type HrmMailSendResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s HrmMailSendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrmMailSendResponseBody) GoString() string {
	return s.String()
}

func (s *HrmMailSendResponseBody) SetResult(v string) *HrmMailSendResponseBody {
	s.Result = &v
	return s
}

type HrmMailSendResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrmMailSendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrmMailSendResponse) String() string {
	return tea.Prettify(s)
}

func (s HrmMailSendResponse) GoString() string {
	return s.String()
}

func (s *HrmMailSendResponse) SetHeaders(v map[string]*string) *HrmMailSendResponse {
	s.Headers = v
	return s
}

func (s *HrmMailSendResponse) SetStatusCode(v int32) *HrmMailSendResponse {
	s.StatusCode = &v
	return s
}

func (s *HrmMailSendResponse) SetBody(v *HrmMailSendResponseBody) *HrmMailSendResponse {
	s.Body = v
	return s
}

type HrmMokaEventHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrmMokaEventHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrmMokaEventHeaders) GoString() string {
	return s.String()
}

func (s *HrmMokaEventHeaders) SetCommonHeaders(v map[string]*string) *HrmMokaEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrmMokaEventHeaders) SetXAcsDingtalkAccessToken(v string) *HrmMokaEventHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrmMokaEventRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// /user/role/get
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"a":"b"}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s HrmMokaEventRequest) String() string {
	return tea.Prettify(s)
}

func (s HrmMokaEventRequest) GoString() string {
	return s.String()
}

func (s *HrmMokaEventRequest) SetBizId(v string) *HrmMokaEventRequest {
	s.BizId = &v
	return s
}

func (s *HrmMokaEventRequest) SetContent(v string) *HrmMokaEventRequest {
	s.Content = &v
	return s
}

type HrmMokaEventResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s HrmMokaEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrmMokaEventResponseBody) GoString() string {
	return s.String()
}

func (s *HrmMokaEventResponseBody) SetResult(v bool) *HrmMokaEventResponseBody {
	s.Result = &v
	return s
}

type HrmMokaEventResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrmMokaEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrmMokaEventResponse) String() string {
	return tea.Prettify(s)
}

func (s HrmMokaEventResponse) GoString() string {
	return s.String()
}

func (s *HrmMokaEventResponse) SetHeaders(v map[string]*string) *HrmMokaEventResponse {
	s.Headers = v
	return s
}

func (s *HrmMokaEventResponse) SetStatusCode(v int32) *HrmMokaEventResponse {
	s.StatusCode = &v
	return s
}

func (s *HrmMokaEventResponse) SetBody(v *HrmMokaEventResponseBody) *HrmMokaEventResponse {
	s.Body = v
	return s
}

type HrmMokaOapiHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrmMokaOapiHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrmMokaOapiHeaders) GoString() string {
	return s.String()
}

func (s *HrmMokaOapiHeaders) SetCommonHeaders(v map[string]*string) *HrmMokaOapiHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrmMokaOapiHeaders) SetXAcsDingtalkAccessToken(v string) *HrmMokaOapiHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrmMokaOapiRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// /user/role/get
	ApiCode *string     `json:"apiCode,omitempty" xml:"apiCode,omitempty"`
	Params  interface{} `json:"params,omitempty" xml:"params,omitempty"`
}

func (s HrmMokaOapiRequest) String() string {
	return tea.Prettify(s)
}

func (s HrmMokaOapiRequest) GoString() string {
	return s.String()
}

func (s *HrmMokaOapiRequest) SetApiCode(v string) *HrmMokaOapiRequest {
	s.ApiCode = &v
	return s
}

func (s *HrmMokaOapiRequest) SetParams(v interface{}) *HrmMokaOapiRequest {
	s.Params = v
	return s
}

type HrmMokaOapiResponseBody struct {
	BizSuccess *bool                  `json:"bizSuccess,omitempty" xml:"bizSuccess,omitempty"`
	ErrorCode  *string                `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg   *string                `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result     map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s HrmMokaOapiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrmMokaOapiResponseBody) GoString() string {
	return s.String()
}

func (s *HrmMokaOapiResponseBody) SetBizSuccess(v bool) *HrmMokaOapiResponseBody {
	s.BizSuccess = &v
	return s
}

func (s *HrmMokaOapiResponseBody) SetErrorCode(v string) *HrmMokaOapiResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *HrmMokaOapiResponseBody) SetErrorMsg(v string) *HrmMokaOapiResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *HrmMokaOapiResponseBody) SetResult(v map[string]interface{}) *HrmMokaOapiResponseBody {
	s.Result = v
	return s
}

type HrmMokaOapiResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrmMokaOapiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrmMokaOapiResponse) String() string {
	return tea.Prettify(s)
}

func (s HrmMokaOapiResponse) GoString() string {
	return s.String()
}

func (s *HrmMokaOapiResponse) SetHeaders(v map[string]*string) *HrmMokaOapiResponse {
	s.Headers = v
	return s
}

func (s *HrmMokaOapiResponse) SetStatusCode(v int32) *HrmMokaOapiResponse {
	s.StatusCode = &v
	return s
}

func (s *HrmMokaOapiResponse) SetBody(v *HrmMokaOapiResponseBody) *HrmMokaOapiResponse {
	s.Body = v
	return s
}

type HrmProcessRegularHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrmProcessRegularHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessRegularHeaders) GoString() string {
	return s.String()
}

func (s *HrmProcessRegularHeaders) SetCommonHeaders(v map[string]*string) *HrmProcessRegularHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrmProcessRegularHeaders) SetXAcsDingtalkAccessToken(v string) *HrmProcessRegularHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrmProcessRegularRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 16690147049882572
	OperationId *string `json:"operationId,omitempty" xml:"operationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1672542359000
	RegularDate *int64 `json:"regularDate,omitempty" xml:"regularDate,omitempty"`
	// example:
	//
	// 同意转正
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 16690147049882572
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s HrmProcessRegularRequest) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessRegularRequest) GoString() string {
	return s.String()
}

func (s *HrmProcessRegularRequest) SetOperationId(v string) *HrmProcessRegularRequest {
	s.OperationId = &v
	return s
}

func (s *HrmProcessRegularRequest) SetRegularDate(v int64) *HrmProcessRegularRequest {
	s.RegularDate = &v
	return s
}

func (s *HrmProcessRegularRequest) SetRemark(v string) *HrmProcessRegularRequest {
	s.Remark = &v
	return s
}

func (s *HrmProcessRegularRequest) SetUserId(v string) *HrmProcessRegularRequest {
	s.UserId = &v
	return s
}

type HrmProcessRegularResponseBody struct {
	// This parameter is required.
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s HrmProcessRegularResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessRegularResponseBody) GoString() string {
	return s.String()
}

func (s *HrmProcessRegularResponseBody) SetResult(v bool) *HrmProcessRegularResponseBody {
	s.Result = &v
	return s
}

type HrmProcessRegularResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrmProcessRegularResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrmProcessRegularResponse) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessRegularResponse) GoString() string {
	return s.String()
}

func (s *HrmProcessRegularResponse) SetHeaders(v map[string]*string) *HrmProcessRegularResponse {
	s.Headers = v
	return s
}

func (s *HrmProcessRegularResponse) SetStatusCode(v int32) *HrmProcessRegularResponse {
	s.StatusCode = &v
	return s
}

func (s *HrmProcessRegularResponse) SetBody(v *HrmProcessRegularResponseBody) *HrmProcessRegularResponse {
	s.Body = v
	return s
}

type HrmProcessTerminationAndHandoverHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrmProcessTerminationAndHandoverHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessTerminationAndHandoverHeaders) GoString() string {
	return s.String()
}

func (s *HrmProcessTerminationAndHandoverHeaders) SetCommonHeaders(v map[string]*string) *HrmProcessTerminationAndHandoverHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrmProcessTerminationAndHandoverHeaders) SetXAcsDingtalkAccessToken(v string) *HrmProcessTerminationAndHandoverHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrmProcessTerminationAndHandoverRequest struct {
	// example:
	//
	// user123
	AflowHandOverUserId *string `json:"aflowHandOverUserId,omitempty" xml:"aflowHandOverUserId,omitempty"`
	// example:
	//
	// user123
	DingPanHandoverUserId *string `json:"dingPanHandoverUserId,omitempty" xml:"dingPanHandoverUserId,omitempty"`
	// example:
	//
	// user123
	DirectSubordinatesHandoverUserId *string `json:"directSubordinatesHandoverUserId,omitempty" xml:"directSubordinatesHandoverUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aefadfadaewedad
	DismissionMemo *string `json:"dismissionMemo,omitempty" xml:"dismissionMemo,omitempty"`
	// example:
	//
	// 1
	DismissionReason *int32 `json:"dismissionReason,omitempty" xml:"dismissionReason,omitempty"`
	// example:
	//
	// user123
	DocNoteHandoverUserId *string `json:"docNoteHandoverUserId,omitempty" xml:"docNoteHandoverUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1704074400000
	LastWorkDate *int64 `json:"lastWorkDate,omitempty" xml:"lastWorkDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 经理
	OptUserId *string `json:"optUserId,omitempty" xml:"optUserId,omitempty"`
	// example:
	//
	// user123
	PermissionHandoverUserId   *string   `json:"permissionHandoverUserId,omitempty" xml:"permissionHandoverUserId,omitempty"`
	TerminationReasonPassive   []*string `json:"terminationReasonPassive,omitempty" xml:"terminationReasonPassive,omitempty" type:"Repeated"`
	TerminationReasonVoluntary []*string `json:"terminationReasonVoluntary,omitempty" xml:"terminationReasonVoluntary,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 2332
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s HrmProcessTerminationAndHandoverRequest) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessTerminationAndHandoverRequest) GoString() string {
	return s.String()
}

func (s *HrmProcessTerminationAndHandoverRequest) SetAflowHandOverUserId(v string) *HrmProcessTerminationAndHandoverRequest {
	s.AflowHandOverUserId = &v
	return s
}

func (s *HrmProcessTerminationAndHandoverRequest) SetDingPanHandoverUserId(v string) *HrmProcessTerminationAndHandoverRequest {
	s.DingPanHandoverUserId = &v
	return s
}

func (s *HrmProcessTerminationAndHandoverRequest) SetDirectSubordinatesHandoverUserId(v string) *HrmProcessTerminationAndHandoverRequest {
	s.DirectSubordinatesHandoverUserId = &v
	return s
}

func (s *HrmProcessTerminationAndHandoverRequest) SetDismissionMemo(v string) *HrmProcessTerminationAndHandoverRequest {
	s.DismissionMemo = &v
	return s
}

func (s *HrmProcessTerminationAndHandoverRequest) SetDismissionReason(v int32) *HrmProcessTerminationAndHandoverRequest {
	s.DismissionReason = &v
	return s
}

func (s *HrmProcessTerminationAndHandoverRequest) SetDocNoteHandoverUserId(v string) *HrmProcessTerminationAndHandoverRequest {
	s.DocNoteHandoverUserId = &v
	return s
}

func (s *HrmProcessTerminationAndHandoverRequest) SetLastWorkDate(v int64) *HrmProcessTerminationAndHandoverRequest {
	s.LastWorkDate = &v
	return s
}

func (s *HrmProcessTerminationAndHandoverRequest) SetOptUserId(v string) *HrmProcessTerminationAndHandoverRequest {
	s.OptUserId = &v
	return s
}

func (s *HrmProcessTerminationAndHandoverRequest) SetPermissionHandoverUserId(v string) *HrmProcessTerminationAndHandoverRequest {
	s.PermissionHandoverUserId = &v
	return s
}

func (s *HrmProcessTerminationAndHandoverRequest) SetTerminationReasonPassive(v []*string) *HrmProcessTerminationAndHandoverRequest {
	s.TerminationReasonPassive = v
	return s
}

func (s *HrmProcessTerminationAndHandoverRequest) SetTerminationReasonVoluntary(v []*string) *HrmProcessTerminationAndHandoverRequest {
	s.TerminationReasonVoluntary = v
	return s
}

func (s *HrmProcessTerminationAndHandoverRequest) SetUserId(v string) *HrmProcessTerminationAndHandoverRequest {
	s.UserId = &v
	return s
}

type HrmProcessTerminationAndHandoverResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s HrmProcessTerminationAndHandoverResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessTerminationAndHandoverResponseBody) GoString() string {
	return s.String()
}

func (s *HrmProcessTerminationAndHandoverResponseBody) SetResult(v bool) *HrmProcessTerminationAndHandoverResponseBody {
	s.Result = &v
	return s
}

type HrmProcessTerminationAndHandoverResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrmProcessTerminationAndHandoverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrmProcessTerminationAndHandoverResponse) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessTerminationAndHandoverResponse) GoString() string {
	return s.String()
}

func (s *HrmProcessTerminationAndHandoverResponse) SetHeaders(v map[string]*string) *HrmProcessTerminationAndHandoverResponse {
	s.Headers = v
	return s
}

func (s *HrmProcessTerminationAndHandoverResponse) SetStatusCode(v int32) *HrmProcessTerminationAndHandoverResponse {
	s.StatusCode = &v
	return s
}

func (s *HrmProcessTerminationAndHandoverResponse) SetBody(v *HrmProcessTerminationAndHandoverResponseBody) *HrmProcessTerminationAndHandoverResponse {
	s.Body = v
	return s
}

type HrmProcessTransferHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrmProcessTransferHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessTransferHeaders) GoString() string {
	return s.String()
}

func (s *HrmProcessTransferHeaders) SetCommonHeaders(v map[string]*string) *HrmProcessTransferHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrmProcessTransferHeaders) SetXAcsDingtalkAccessToken(v string) *HrmProcessTransferHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrmProcessTransferRequest struct {
	DeptIdsAfterTransfer []*int64 `json:"deptIdsAfterTransfer,omitempty" xml:"deptIdsAfterTransfer,omitempty" type:"Repeated"`
	// example:
	//
	// aefadfadaewedad
	JobIdAfterTransfer *string `json:"jobIdAfterTransfer,omitempty" xml:"jobIdAfterTransfer,omitempty"`
	// example:
	//
	// 123L
	MainDeptIdAfterTransfer *int64 `json:"mainDeptIdAfterTransfer,omitempty" xml:"mainDeptIdAfterTransfer,omitempty"`
	// example:
	//
	// 232312312
	OperateUserId *string `json:"operateUserId,omitempty" xml:"operateUserId,omitempty"`
	// example:
	//
	// fasdfaddsadfa
	PositionIdAfterTransfer *string `json:"positionIdAfterTransfer,omitempty" xml:"positionIdAfterTransfer,omitempty"`
	// example:
	//
	// L1
	PositionLevelAfterTransfer *string `json:"positionLevelAfterTransfer,omitempty" xml:"positionLevelAfterTransfer,omitempty"`
	// example:
	//
	// 经理
	PositionNameAfterTransfer *string `json:"positionNameAfterTransfer,omitempty" xml:"positionNameAfterTransfer,omitempty"`
	// example:
	//
	// fasdfaddsadfa
	RankIdAfterTransfer *string `json:"rankIdAfterTransfer,omitempty" xml:"rankIdAfterTransfer,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2332
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s HrmProcessTransferRequest) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessTransferRequest) GoString() string {
	return s.String()
}

func (s *HrmProcessTransferRequest) SetDeptIdsAfterTransfer(v []*int64) *HrmProcessTransferRequest {
	s.DeptIdsAfterTransfer = v
	return s
}

func (s *HrmProcessTransferRequest) SetJobIdAfterTransfer(v string) *HrmProcessTransferRequest {
	s.JobIdAfterTransfer = &v
	return s
}

func (s *HrmProcessTransferRequest) SetMainDeptIdAfterTransfer(v int64) *HrmProcessTransferRequest {
	s.MainDeptIdAfterTransfer = &v
	return s
}

func (s *HrmProcessTransferRequest) SetOperateUserId(v string) *HrmProcessTransferRequest {
	s.OperateUserId = &v
	return s
}

func (s *HrmProcessTransferRequest) SetPositionIdAfterTransfer(v string) *HrmProcessTransferRequest {
	s.PositionIdAfterTransfer = &v
	return s
}

func (s *HrmProcessTransferRequest) SetPositionLevelAfterTransfer(v string) *HrmProcessTransferRequest {
	s.PositionLevelAfterTransfer = &v
	return s
}

func (s *HrmProcessTransferRequest) SetPositionNameAfterTransfer(v string) *HrmProcessTransferRequest {
	s.PositionNameAfterTransfer = &v
	return s
}

func (s *HrmProcessTransferRequest) SetRankIdAfterTransfer(v string) *HrmProcessTransferRequest {
	s.RankIdAfterTransfer = &v
	return s
}

func (s *HrmProcessTransferRequest) SetUserId(v string) *HrmProcessTransferRequest {
	s.UserId = &v
	return s
}

type HrmProcessTransferResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s HrmProcessTransferResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessTransferResponseBody) GoString() string {
	return s.String()
}

func (s *HrmProcessTransferResponseBody) SetResult(v bool) *HrmProcessTransferResponseBody {
	s.Result = &v
	return s
}

type HrmProcessTransferResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrmProcessTransferResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrmProcessTransferResponse) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessTransferResponse) GoString() string {
	return s.String()
}

func (s *HrmProcessTransferResponse) SetHeaders(v map[string]*string) *HrmProcessTransferResponse {
	s.Headers = v
	return s
}

func (s *HrmProcessTransferResponse) SetStatusCode(v int32) *HrmProcessTransferResponse {
	s.StatusCode = &v
	return s
}

func (s *HrmProcessTransferResponse) SetBody(v *HrmProcessTransferResponseBody) *HrmProcessTransferResponse {
	s.Body = v
	return s
}

type HrmProcessUpdateTerminationInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrmProcessUpdateTerminationInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessUpdateTerminationInfoHeaders) GoString() string {
	return s.String()
}

func (s *HrmProcessUpdateTerminationInfoHeaders) SetCommonHeaders(v map[string]*string) *HrmProcessUpdateTerminationInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrmProcessUpdateTerminationInfoHeaders) SetXAcsDingtalkAccessToken(v string) *HrmProcessUpdateTerminationInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrmProcessUpdateTerminationInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 因个人原因离职
	DismissionMemo *string `json:"dismissionMemo,omitempty" xml:"dismissionMemo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1672502400000
	LastWorkDate *int64 `json:"lastWorkDate,omitempty" xml:"lastWorkDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// admin123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s HrmProcessUpdateTerminationInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessUpdateTerminationInfoRequest) GoString() string {
	return s.String()
}

func (s *HrmProcessUpdateTerminationInfoRequest) SetDismissionMemo(v string) *HrmProcessUpdateTerminationInfoRequest {
	s.DismissionMemo = &v
	return s
}

func (s *HrmProcessUpdateTerminationInfoRequest) SetLastWorkDate(v int64) *HrmProcessUpdateTerminationInfoRequest {
	s.LastWorkDate = &v
	return s
}

func (s *HrmProcessUpdateTerminationInfoRequest) SetUserId(v string) *HrmProcessUpdateTerminationInfoRequest {
	s.UserId = &v
	return s
}

type HrmProcessUpdateTerminationInfoResponseBody struct {
	// This parameter is required.
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s HrmProcessUpdateTerminationInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessUpdateTerminationInfoResponseBody) GoString() string {
	return s.String()
}

func (s *HrmProcessUpdateTerminationInfoResponseBody) SetResult(v bool) *HrmProcessUpdateTerminationInfoResponseBody {
	s.Result = &v
	return s
}

type HrmProcessUpdateTerminationInfoResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrmProcessUpdateTerminationInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrmProcessUpdateTerminationInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s HrmProcessUpdateTerminationInfoResponse) GoString() string {
	return s.String()
}

func (s *HrmProcessUpdateTerminationInfoResponse) SetHeaders(v map[string]*string) *HrmProcessUpdateTerminationInfoResponse {
	s.Headers = v
	return s
}

func (s *HrmProcessUpdateTerminationInfoResponse) SetStatusCode(v int32) *HrmProcessUpdateTerminationInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *HrmProcessUpdateTerminationInfoResponse) SetBody(v *HrmProcessUpdateTerminationInfoResponseBody) *HrmProcessUpdateTerminationInfoResponse {
	s.Body = v
	return s
}

type HrmPtsServiceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HrmPtsServiceHeaders) String() string {
	return tea.Prettify(s)
}

func (s HrmPtsServiceHeaders) GoString() string {
	return s.String()
}

func (s *HrmPtsServiceHeaders) SetCommonHeaders(v map[string]*string) *HrmPtsServiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HrmPtsServiceHeaders) SetXAcsDingtalkAccessToken(v string) *HrmPtsServiceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HrmPtsServiceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dev  or online
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// example:
	//
	// GET/POST
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abd123213
	OuterId *string     `json:"outerId,omitempty" xml:"outerId,omitempty"`
	Params  interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /user/role/get
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s HrmPtsServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s HrmPtsServiceRequest) GoString() string {
	return s.String()
}

func (s *HrmPtsServiceRequest) SetEnv(v string) *HrmPtsServiceRequest {
	s.Env = &v
	return s
}

func (s *HrmPtsServiceRequest) SetMethod(v string) *HrmPtsServiceRequest {
	s.Method = &v
	return s
}

func (s *HrmPtsServiceRequest) SetOuterId(v string) *HrmPtsServiceRequest {
	s.OuterId = &v
	return s
}

func (s *HrmPtsServiceRequest) SetParams(v interface{}) *HrmPtsServiceRequest {
	s.Params = v
	return s
}

func (s *HrmPtsServiceRequest) SetPath(v string) *HrmPtsServiceRequest {
	s.Path = &v
	return s
}

type HrmPtsServiceResponseBody struct {
	BizSuccess *bool                  `json:"bizSuccess,omitempty" xml:"bizSuccess,omitempty"`
	ErrorCode  *string                `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg   *string                `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result     map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s HrmPtsServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HrmPtsServiceResponseBody) GoString() string {
	return s.String()
}

func (s *HrmPtsServiceResponseBody) SetBizSuccess(v bool) *HrmPtsServiceResponseBody {
	s.BizSuccess = &v
	return s
}

func (s *HrmPtsServiceResponseBody) SetErrorCode(v string) *HrmPtsServiceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *HrmPtsServiceResponseBody) SetErrorMsg(v string) *HrmPtsServiceResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *HrmPtsServiceResponseBody) SetResult(v map[string]interface{}) *HrmPtsServiceResponseBody {
	s.Result = v
	return s
}

type HrmPtsServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HrmPtsServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HrmPtsServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s HrmPtsServiceResponse) GoString() string {
	return s.String()
}

func (s *HrmPtsServiceResponse) SetHeaders(v map[string]*string) *HrmPtsServiceResponse {
	s.Headers = v
	return s
}

func (s *HrmPtsServiceResponse) SetStatusCode(v int32) *HrmPtsServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *HrmPtsServiceResponse) SetBody(v *HrmPtsServiceResponseBody) *HrmPtsServiceResponse {
	s.Body = v
	return s
}

type InvalidSignRecordsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InvalidSignRecordsHeaders) String() string {
	return tea.Prettify(s)
}

func (s InvalidSignRecordsHeaders) GoString() string {
	return s.String()
}

func (s *InvalidSignRecordsHeaders) SetCommonHeaders(v map[string]*string) *InvalidSignRecordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvalidSignRecordsHeaders) SetXAcsDingtalkAccessToken(v string) *InvalidSignRecordsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InvalidSignRecordsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	InvalidUserId *string `json:"invalidUserId,omitempty" xml:"invalidUserId,omitempty"`
	// This parameter is required.
	SignRecordIds []*string `json:"signRecordIds,omitempty" xml:"signRecordIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 作废测试
	StatusRemark *string `json:"statusRemark,omitempty" xml:"statusRemark,omitempty"`
}

func (s InvalidSignRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s InvalidSignRecordsRequest) GoString() string {
	return s.String()
}

func (s *InvalidSignRecordsRequest) SetInvalidUserId(v string) *InvalidSignRecordsRequest {
	s.InvalidUserId = &v
	return s
}

func (s *InvalidSignRecordsRequest) SetSignRecordIds(v []*string) *InvalidSignRecordsRequest {
	s.SignRecordIds = v
	return s
}

func (s *InvalidSignRecordsRequest) SetStatusRemark(v string) *InvalidSignRecordsRequest {
	s.StatusRemark = &v
	return s
}

type InvalidSignRecordsResponseBody struct {
	Result *InvalidSignRecordsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// This parameter is required.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s InvalidSignRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvalidSignRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *InvalidSignRecordsResponseBody) SetResult(v *InvalidSignRecordsResponseBodyResult) *InvalidSignRecordsResponseBody {
	s.Result = v
	return s
}

func (s *InvalidSignRecordsResponseBody) SetSuccess(v bool) *InvalidSignRecordsResponseBody {
	s.Success = &v
	return s
}

type InvalidSignRecordsResponseBodyResult struct {
	FailItems    []*InvalidSignRecordsResponseBodyResultFailItems    `json:"failItems,omitempty" xml:"failItems,omitempty" type:"Repeated"`
	SuccessItems []*InvalidSignRecordsResponseBodyResultSuccessItems `json:"successItems,omitempty" xml:"successItems,omitempty" type:"Repeated"`
}

func (s InvalidSignRecordsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s InvalidSignRecordsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *InvalidSignRecordsResponseBodyResult) SetFailItems(v []*InvalidSignRecordsResponseBodyResultFailItems) *InvalidSignRecordsResponseBodyResult {
	s.FailItems = v
	return s
}

func (s *InvalidSignRecordsResponseBodyResult) SetSuccessItems(v []*InvalidSignRecordsResponseBodyResultSuccessItems) *InvalidSignRecordsResponseBodyResult {
	s.SuccessItems = v
	return s
}

type InvalidSignRecordsResponseBodyResultFailItems struct {
	// example:
	//
	// 1234566789
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// example:
	//
	// 电子签状态变更不合法
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s InvalidSignRecordsResponseBodyResultFailItems) String() string {
	return tea.Prettify(s)
}

func (s InvalidSignRecordsResponseBodyResultFailItems) GoString() string {
	return s.String()
}

func (s *InvalidSignRecordsResponseBodyResultFailItems) SetItemId(v string) *InvalidSignRecordsResponseBodyResultFailItems {
	s.ItemId = &v
	return s
}

func (s *InvalidSignRecordsResponseBodyResultFailItems) SetType(v string) *InvalidSignRecordsResponseBodyResultFailItems {
	s.Type = &v
	return s
}

type InvalidSignRecordsResponseBodyResultSuccessItems struct {
	// example:
	//
	// 123456789
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
}

func (s InvalidSignRecordsResponseBodyResultSuccessItems) String() string {
	return tea.Prettify(s)
}

func (s InvalidSignRecordsResponseBodyResultSuccessItems) GoString() string {
	return s.String()
}

func (s *InvalidSignRecordsResponseBodyResultSuccessItems) SetItemId(v string) *InvalidSignRecordsResponseBodyResultSuccessItems {
	s.ItemId = &v
	return s
}

type InvalidSignRecordsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvalidSignRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvalidSignRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s InvalidSignRecordsResponse) GoString() string {
	return s.String()
}

func (s *InvalidSignRecordsResponse) SetHeaders(v map[string]*string) *InvalidSignRecordsResponse {
	s.Headers = v
	return s
}

func (s *InvalidSignRecordsResponse) SetStatusCode(v int32) *InvalidSignRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *InvalidSignRecordsResponse) SetBody(v *InvalidSignRecordsResponseBody) *InvalidSignRecordsResponse {
	s.Body = v
	return s
}

type MasterDataDeleteHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MasterDataDeleteHeaders) String() string {
	return tea.Prettify(s)
}

func (s MasterDataDeleteHeaders) GoString() string {
	return s.String()
}

func (s *MasterDataDeleteHeaders) SetCommonHeaders(v map[string]*string) *MasterDataDeleteHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MasterDataDeleteHeaders) SetXAcsDingtalkAccessToken(v string) *MasterDataDeleteHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MasterDataDeleteRequest struct {
	// This parameter is required.
	Body []*MasterDataDeleteRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	TenantId *int64 `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s MasterDataDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s MasterDataDeleteRequest) GoString() string {
	return s.String()
}

func (s *MasterDataDeleteRequest) SetBody(v []*MasterDataDeleteRequestBody) *MasterDataDeleteRequest {
	s.Body = v
	return s
}

func (s *MasterDataDeleteRequest) SetTenantId(v int64) *MasterDataDeleteRequest {
	s.TenantId = &v
	return s
}

type MasterDataDeleteRequestBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 12312
	BizTime *int64 `json:"bizTime,omitempty" xml:"bizTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// uk123
	BizUk *string `json:"bizUk,omitempty" xml:"bizUk,omitempty"`
	// example:
	//
	// base
	EntityCode *string                                 `json:"entityCode,omitempty" xml:"entityCode,omitempty"`
	FieldList  []*MasterDataDeleteRequestBodyFieldList `json:"fieldList,omitempty" xml:"fieldList,omitempty" type:"Repeated"`
	// This parameter is required.
	Scope *MasterDataDeleteRequestBodyScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
}

func (s MasterDataDeleteRequestBody) String() string {
	return tea.Prettify(s)
}

func (s MasterDataDeleteRequestBody) GoString() string {
	return s.String()
}

func (s *MasterDataDeleteRequestBody) SetBizTime(v int64) *MasterDataDeleteRequestBody {
	s.BizTime = &v
	return s
}

func (s *MasterDataDeleteRequestBody) SetBizUk(v string) *MasterDataDeleteRequestBody {
	s.BizUk = &v
	return s
}

func (s *MasterDataDeleteRequestBody) SetEntityCode(v string) *MasterDataDeleteRequestBody {
	s.EntityCode = &v
	return s
}

func (s *MasterDataDeleteRequestBody) SetFieldList(v []*MasterDataDeleteRequestBodyFieldList) *MasterDataDeleteRequestBody {
	s.FieldList = v
	return s
}

func (s *MasterDataDeleteRequestBody) SetScope(v *MasterDataDeleteRequestBodyScope) *MasterDataDeleteRequestBody {
	s.Scope = v
	return s
}

type MasterDataDeleteRequestBodyFieldList struct {
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 123
	ValueStr *string `json:"valueStr,omitempty" xml:"valueStr,omitempty"`
}

func (s MasterDataDeleteRequestBodyFieldList) String() string {
	return tea.Prettify(s)
}

func (s MasterDataDeleteRequestBodyFieldList) GoString() string {
	return s.String()
}

func (s *MasterDataDeleteRequestBodyFieldList) SetName(v string) *MasterDataDeleteRequestBodyFieldList {
	s.Name = &v
	return s
}

func (s *MasterDataDeleteRequestBodyFieldList) SetValueStr(v string) *MasterDataDeleteRequestBodyFieldList {
	s.ValueStr = &v
	return s
}

type MasterDataDeleteRequestBodyScope struct {
	// This parameter is required.
	//
	// example:
	//
	// performance
	ScopeCode *string `json:"scopeCode,omitempty" xml:"scopeCode,omitempty"`
	// example:
	//
	// 1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s MasterDataDeleteRequestBodyScope) String() string {
	return tea.Prettify(s)
}

func (s MasterDataDeleteRequestBodyScope) GoString() string {
	return s.String()
}

func (s *MasterDataDeleteRequestBodyScope) SetScopeCode(v string) *MasterDataDeleteRequestBodyScope {
	s.ScopeCode = &v
	return s
}

func (s *MasterDataDeleteRequestBodyScope) SetVersion(v int32) *MasterDataDeleteRequestBodyScope {
	s.Version = &v
	return s
}

type MasterDataDeleteResponseBody struct {
	// This parameter is required.
	AllSuccess *bool                                     `json:"allSuccess,omitempty" xml:"allSuccess,omitempty"`
	FailResult []*MasterDataDeleteResponseBodyFailResult `json:"failResult,omitempty" xml:"failResult,omitempty" type:"Repeated"`
}

func (s MasterDataDeleteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MasterDataDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *MasterDataDeleteResponseBody) SetAllSuccess(v bool) *MasterDataDeleteResponseBody {
	s.AllSuccess = &v
	return s
}

func (s *MasterDataDeleteResponseBody) SetFailResult(v []*MasterDataDeleteResponseBodyFailResult) *MasterDataDeleteResponseBody {
	s.FailResult = v
	return s
}

type MasterDataDeleteResponseBodyFailResult struct {
	// example:
	//
	// uk123
	BizUK *string `json:"bizUK,omitempty" xml:"bizUK,omitempty"`
	// example:
	//
	// S0005
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// 主键冲突
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Success  *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s MasterDataDeleteResponseBodyFailResult) String() string {
	return tea.Prettify(s)
}

func (s MasterDataDeleteResponseBodyFailResult) GoString() string {
	return s.String()
}

func (s *MasterDataDeleteResponseBodyFailResult) SetBizUK(v string) *MasterDataDeleteResponseBodyFailResult {
	s.BizUK = &v
	return s
}

func (s *MasterDataDeleteResponseBodyFailResult) SetErrorCode(v string) *MasterDataDeleteResponseBodyFailResult {
	s.ErrorCode = &v
	return s
}

func (s *MasterDataDeleteResponseBodyFailResult) SetErrorMsg(v string) *MasterDataDeleteResponseBodyFailResult {
	s.ErrorMsg = &v
	return s
}

func (s *MasterDataDeleteResponseBodyFailResult) SetSuccess(v bool) *MasterDataDeleteResponseBodyFailResult {
	s.Success = &v
	return s
}

type MasterDataDeleteResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MasterDataDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MasterDataDeleteResponse) String() string {
	return tea.Prettify(s)
}

func (s MasterDataDeleteResponse) GoString() string {
	return s.String()
}

func (s *MasterDataDeleteResponse) SetHeaders(v map[string]*string) *MasterDataDeleteResponse {
	s.Headers = v
	return s
}

func (s *MasterDataDeleteResponse) SetStatusCode(v int32) *MasterDataDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *MasterDataDeleteResponse) SetBody(v *MasterDataDeleteResponseBody) *MasterDataDeleteResponse {
	s.Body = v
	return s
}

type MasterDataQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MasterDataQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryHeaders) GoString() string {
	return s.String()
}

func (s *MasterDataQueryHeaders) SetCommonHeaders(v map[string]*string) *MasterDataQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MasterDataQueryHeaders) SetXAcsDingtalkAccessToken(v string) *MasterDataQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MasterDataQueryRequest struct {
	// example:
	//
	// uk_12123
	BizUK *string `json:"bizUK,omitempty" xml:"bizUK,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	NextToken *int32 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// admin1234
	OptUserId   *string                              `json:"optUserId,omitempty" xml:"optUserId,omitempty"`
	QueryParams []*MasterDataQueryRequestQueryParams `json:"queryParams,omitempty" xml:"queryParams,omitempty" type:"Repeated"`
	// This parameter is required.
	RelationIds []*string `json:"relationIds,omitempty" xml:"relationIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// PERFORMANCE
	ScopeCode *string `json:"scopeCode,omitempty" xml:"scopeCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	TenantId *int64 `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// base
	ViewEntityCode *string `json:"viewEntityCode,omitempty" xml:"viewEntityCode,omitempty"`
}

func (s MasterDataQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryRequest) GoString() string {
	return s.String()
}

func (s *MasterDataQueryRequest) SetBizUK(v string) *MasterDataQueryRequest {
	s.BizUK = &v
	return s
}

func (s *MasterDataQueryRequest) SetMaxResults(v int32) *MasterDataQueryRequest {
	s.MaxResults = &v
	return s
}

func (s *MasterDataQueryRequest) SetNextToken(v int32) *MasterDataQueryRequest {
	s.NextToken = &v
	return s
}

func (s *MasterDataQueryRequest) SetOptUserId(v string) *MasterDataQueryRequest {
	s.OptUserId = &v
	return s
}

func (s *MasterDataQueryRequest) SetQueryParams(v []*MasterDataQueryRequestQueryParams) *MasterDataQueryRequest {
	s.QueryParams = v
	return s
}

func (s *MasterDataQueryRequest) SetRelationIds(v []*string) *MasterDataQueryRequest {
	s.RelationIds = v
	return s
}

func (s *MasterDataQueryRequest) SetScopeCode(v string) *MasterDataQueryRequest {
	s.ScopeCode = &v
	return s
}

func (s *MasterDataQueryRequest) SetTenantId(v int64) *MasterDataQueryRequest {
	s.TenantId = &v
	return s
}

func (s *MasterDataQueryRequest) SetViewEntityCode(v string) *MasterDataQueryRequest {
	s.ViewEntityCode = &v
	return s
}

type MasterDataQueryRequestQueryParams struct {
	ConditionList []*MasterDataQueryRequestQueryParamsConditionList `json:"conditionList,omitempty" xml:"conditionList,omitempty" type:"Repeated"`
	// example:
	//
	// performance_code
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	// example:
	//
	// AND
	JoinType *string `json:"joinType,omitempty" xml:"joinType,omitempty"`
}

func (s MasterDataQueryRequestQueryParams) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryRequestQueryParams) GoString() string {
	return s.String()
}

func (s *MasterDataQueryRequestQueryParams) SetConditionList(v []*MasterDataQueryRequestQueryParamsConditionList) *MasterDataQueryRequestQueryParams {
	s.ConditionList = v
	return s
}

func (s *MasterDataQueryRequestQueryParams) SetFieldCode(v string) *MasterDataQueryRequestQueryParams {
	s.FieldCode = &v
	return s
}

func (s *MasterDataQueryRequestQueryParams) SetJoinType(v string) *MasterDataQueryRequestQueryParams {
	s.JoinType = &v
	return s
}

type MasterDataQueryRequestQueryParamsConditionList struct {
	// example:
	//
	// EQUAL
	Operate *string `json:"operate,omitempty" xml:"operate,omitempty"`
	Value   *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s MasterDataQueryRequestQueryParamsConditionList) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryRequestQueryParamsConditionList) GoString() string {
	return s.String()
}

func (s *MasterDataQueryRequestQueryParamsConditionList) SetOperate(v string) *MasterDataQueryRequestQueryParamsConditionList {
	s.Operate = &v
	return s
}

func (s *MasterDataQueryRequestQueryParamsConditionList) SetValue(v string) *MasterDataQueryRequestQueryParamsConditionList {
	s.Value = &v
	return s
}

type MasterDataQueryResponseBody struct {
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 0
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	Result []*MasterDataQueryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 100
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s MasterDataQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryResponseBody) GoString() string {
	return s.String()
}

func (s *MasterDataQueryResponseBody) SetHasMore(v bool) *MasterDataQueryResponseBody {
	s.HasMore = &v
	return s
}

func (s *MasterDataQueryResponseBody) SetNextToken(v int64) *MasterDataQueryResponseBody {
	s.NextToken = &v
	return s
}

func (s *MasterDataQueryResponseBody) SetResult(v []*MasterDataQueryResponseBodyResult) *MasterDataQueryResponseBody {
	s.Result = v
	return s
}

func (s *MasterDataQueryResponseBody) SetSuccess(v bool) *MasterDataQueryResponseBody {
	s.Success = &v
	return s
}

func (s *MasterDataQueryResponseBody) SetTotal(v int64) *MasterDataQueryResponseBody {
	s.Total = &v
	return s
}

type MasterDataQueryResponseBodyResult struct {
	// example:
	//
	// uk123123
	OuterId *string `json:"outerId,omitempty" xml:"outerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// admind123
	RelationId *string `json:"relationId,omitempty" xml:"relationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PERFORMANCE
	ScopeCode *string `json:"scopeCode,omitempty" xml:"scopeCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// base
	ViewEntityCode        *string                                                   `json:"viewEntityCode,omitempty" xml:"viewEntityCode,omitempty"`
	ViewEntityFieldVOList []*MasterDataQueryResponseBodyResultViewEntityFieldVOList `json:"viewEntityFieldVOList,omitempty" xml:"viewEntityFieldVOList,omitempty" type:"Repeated"`
}

func (s MasterDataQueryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *MasterDataQueryResponseBodyResult) SetOuterId(v string) *MasterDataQueryResponseBodyResult {
	s.OuterId = &v
	return s
}

func (s *MasterDataQueryResponseBodyResult) SetRelationId(v string) *MasterDataQueryResponseBodyResult {
	s.RelationId = &v
	return s
}

func (s *MasterDataQueryResponseBodyResult) SetScopeCode(v string) *MasterDataQueryResponseBodyResult {
	s.ScopeCode = &v
	return s
}

func (s *MasterDataQueryResponseBodyResult) SetViewEntityCode(v string) *MasterDataQueryResponseBodyResult {
	s.ViewEntityCode = &v
	return s
}

func (s *MasterDataQueryResponseBodyResult) SetViewEntityFieldVOList(v []*MasterDataQueryResponseBodyResultViewEntityFieldVOList) *MasterDataQueryResponseBodyResult {
	s.ViewEntityFieldVOList = v
	return s
}

type MasterDataQueryResponseBodyResultViewEntityFieldVOList struct {
	// example:
	//
	// performanceValue
	FieldCode   *string                                                            `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	FieldDataVO *MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO `json:"fieldDataVO,omitempty" xml:"fieldDataVO,omitempty" type:"Struct"`
	// example:
	//
	// 绩效等级
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// example:
	//
	// 1
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
}

func (s MasterDataQueryResponseBodyResultViewEntityFieldVOList) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryResponseBodyResultViewEntityFieldVOList) GoString() string {
	return s.String()
}

func (s *MasterDataQueryResponseBodyResultViewEntityFieldVOList) SetFieldCode(v string) *MasterDataQueryResponseBodyResultViewEntityFieldVOList {
	s.FieldCode = &v
	return s
}

func (s *MasterDataQueryResponseBodyResultViewEntityFieldVOList) SetFieldDataVO(v *MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO) *MasterDataQueryResponseBodyResultViewEntityFieldVOList {
	s.FieldDataVO = v
	return s
}

func (s *MasterDataQueryResponseBodyResultViewEntityFieldVOList) SetFieldName(v string) *MasterDataQueryResponseBodyResultViewEntityFieldVOList {
	s.FieldName = &v
	return s
}

func (s *MasterDataQueryResponseBodyResultViewEntityFieldVOList) SetFieldType(v string) *MasterDataQueryResponseBodyResultViewEntityFieldVOList {
	s.FieldType = &v
	return s
}

type MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO struct {
	// example:
	//
	// 100
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// 100
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO) GoString() string {
	return s.String()
}

func (s *MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO) SetKey(v string) *MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO {
	s.Key = &v
	return s
}

func (s *MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO) SetValue(v string) *MasterDataQueryResponseBodyResultViewEntityFieldVOListFieldDataVO {
	s.Value = &v
	return s
}

type MasterDataQueryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MasterDataQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MasterDataQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s MasterDataQueryResponse) GoString() string {
	return s.String()
}

func (s *MasterDataQueryResponse) SetHeaders(v map[string]*string) *MasterDataQueryResponse {
	s.Headers = v
	return s
}

func (s *MasterDataQueryResponse) SetStatusCode(v int32) *MasterDataQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *MasterDataQueryResponse) SetBody(v *MasterDataQueryResponseBody) *MasterDataQueryResponse {
	s.Body = v
	return s
}

type MasterDataSaveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MasterDataSaveHeaders) String() string {
	return tea.Prettify(s)
}

func (s MasterDataSaveHeaders) GoString() string {
	return s.String()
}

func (s *MasterDataSaveHeaders) SetCommonHeaders(v map[string]*string) *MasterDataSaveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MasterDataSaveHeaders) SetXAcsDingtalkAccessToken(v string) *MasterDataSaveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MasterDataSaveRequest struct {
	// This parameter is required.
	Body []*MasterDataSaveRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	TenantId *int64 `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s MasterDataSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s MasterDataSaveRequest) GoString() string {
	return s.String()
}

func (s *MasterDataSaveRequest) SetBody(v []*MasterDataSaveRequestBody) *MasterDataSaveRequest {
	s.Body = v
	return s
}

func (s *MasterDataSaveRequest) SetTenantId(v int64) *MasterDataSaveRequest {
	s.TenantId = &v
	return s
}

type MasterDataSaveRequestBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 12312
	BizTime *int64 `json:"bizTime,omitempty" xml:"bizTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// uk123
	BizUk *string `json:"bizUk,omitempty" xml:"bizUk,omitempty"`
	// example:
	//
	// base
	EntityCode *string `json:"entityCode,omitempty" xml:"entityCode,omitempty"`
	// This parameter is required.
	FieldList []*MasterDataSaveRequestBodyFieldList `json:"fieldList,omitempty" xml:"fieldList,omitempty" type:"Repeated"`
	// This parameter is required.
	Scope *MasterDataSaveRequestBodyScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// user123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s MasterDataSaveRequestBody) String() string {
	return tea.Prettify(s)
}

func (s MasterDataSaveRequestBody) GoString() string {
	return s.String()
}

func (s *MasterDataSaveRequestBody) SetBizTime(v int64) *MasterDataSaveRequestBody {
	s.BizTime = &v
	return s
}

func (s *MasterDataSaveRequestBody) SetBizUk(v string) *MasterDataSaveRequestBody {
	s.BizUk = &v
	return s
}

func (s *MasterDataSaveRequestBody) SetEntityCode(v string) *MasterDataSaveRequestBody {
	s.EntityCode = &v
	return s
}

func (s *MasterDataSaveRequestBody) SetFieldList(v []*MasterDataSaveRequestBodyFieldList) *MasterDataSaveRequestBody {
	s.FieldList = v
	return s
}

func (s *MasterDataSaveRequestBody) SetScope(v *MasterDataSaveRequestBodyScope) *MasterDataSaveRequestBody {
	s.Scope = v
	return s
}

func (s *MasterDataSaveRequestBody) SetUserId(v string) *MasterDataSaveRequestBody {
	s.UserId = &v
	return s
}

type MasterDataSaveRequestBodyFieldList struct {
	// This parameter is required.
	//
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	ValueStr *string `json:"valueStr,omitempty" xml:"valueStr,omitempty"`
}

func (s MasterDataSaveRequestBodyFieldList) String() string {
	return tea.Prettify(s)
}

func (s MasterDataSaveRequestBodyFieldList) GoString() string {
	return s.String()
}

func (s *MasterDataSaveRequestBodyFieldList) SetName(v string) *MasterDataSaveRequestBodyFieldList {
	s.Name = &v
	return s
}

func (s *MasterDataSaveRequestBodyFieldList) SetValueStr(v string) *MasterDataSaveRequestBodyFieldList {
	s.ValueStr = &v
	return s
}

type MasterDataSaveRequestBodyScope struct {
	// This parameter is required.
	//
	// example:
	//
	// performance
	ScopeCode *string `json:"scopeCode,omitempty" xml:"scopeCode,omitempty"`
	// example:
	//
	// 1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s MasterDataSaveRequestBodyScope) String() string {
	return tea.Prettify(s)
}

func (s MasterDataSaveRequestBodyScope) GoString() string {
	return s.String()
}

func (s *MasterDataSaveRequestBodyScope) SetScopeCode(v string) *MasterDataSaveRequestBodyScope {
	s.ScopeCode = &v
	return s
}

func (s *MasterDataSaveRequestBodyScope) SetVersion(v int32) *MasterDataSaveRequestBodyScope {
	s.Version = &v
	return s
}

type MasterDataSaveResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	AllSuccess *bool                                   `json:"allSuccess,omitempty" xml:"allSuccess,omitempty"`
	FailResult []*MasterDataSaveResponseBodyFailResult `json:"failResult,omitempty" xml:"failResult,omitempty" type:"Repeated"`
}

func (s MasterDataSaveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MasterDataSaveResponseBody) GoString() string {
	return s.String()
}

func (s *MasterDataSaveResponseBody) SetAllSuccess(v bool) *MasterDataSaveResponseBody {
	s.AllSuccess = &v
	return s
}

func (s *MasterDataSaveResponseBody) SetFailResult(v []*MasterDataSaveResponseBodyFailResult) *MasterDataSaveResponseBody {
	s.FailResult = v
	return s
}

type MasterDataSaveResponseBodyFailResult struct {
	// example:
	//
	// uk123
	BizUk *string `json:"bizUk,omitempty" xml:"bizUk,omitempty"`
	// example:
	//
	// S0005
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// 主键冲突
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s MasterDataSaveResponseBodyFailResult) String() string {
	return tea.Prettify(s)
}

func (s MasterDataSaveResponseBodyFailResult) GoString() string {
	return s.String()
}

func (s *MasterDataSaveResponseBodyFailResult) SetBizUk(v string) *MasterDataSaveResponseBodyFailResult {
	s.BizUk = &v
	return s
}

func (s *MasterDataSaveResponseBodyFailResult) SetErrorCode(v string) *MasterDataSaveResponseBodyFailResult {
	s.ErrorCode = &v
	return s
}

func (s *MasterDataSaveResponseBodyFailResult) SetErrorMsg(v string) *MasterDataSaveResponseBodyFailResult {
	s.ErrorMsg = &v
	return s
}

func (s *MasterDataSaveResponseBodyFailResult) SetSuccess(v bool) *MasterDataSaveResponseBodyFailResult {
	s.Success = &v
	return s
}

type MasterDataSaveResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MasterDataSaveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MasterDataSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s MasterDataSaveResponse) GoString() string {
	return s.String()
}

func (s *MasterDataSaveResponse) SetHeaders(v map[string]*string) *MasterDataSaveResponse {
	s.Headers = v
	return s
}

func (s *MasterDataSaveResponse) SetStatusCode(v int32) *MasterDataSaveResponse {
	s.StatusCode = &v
	return s
}

func (s *MasterDataSaveResponse) SetBody(v *MasterDataSaveResponseBody) *MasterDataSaveResponse {
	s.Body = v
	return s
}

type MasterDataTenantQueyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MasterDataTenantQueyHeaders) String() string {
	return tea.Prettify(s)
}

func (s MasterDataTenantQueyHeaders) GoString() string {
	return s.String()
}

func (s *MasterDataTenantQueyHeaders) SetCommonHeaders(v map[string]*string) *MasterDataTenantQueyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MasterDataTenantQueyHeaders) SetXAcsDingtalkAccessToken(v string) *MasterDataTenantQueyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MasterDataTenantQueyRequest struct {
	// This parameter is required.
	EntityCode *string `json:"entityCode,omitempty" xml:"entityCode,omitempty"`
	// This parameter is required.
	ScopeCode *string `json:"scopeCode,omitempty" xml:"scopeCode,omitempty"`
}

func (s MasterDataTenantQueyRequest) String() string {
	return tea.Prettify(s)
}

func (s MasterDataTenantQueyRequest) GoString() string {
	return s.String()
}

func (s *MasterDataTenantQueyRequest) SetEntityCode(v string) *MasterDataTenantQueyRequest {
	s.EntityCode = &v
	return s
}

func (s *MasterDataTenantQueyRequest) SetScopeCode(v string) *MasterDataTenantQueyRequest {
	s.ScopeCode = &v
	return s
}

type MasterDataTenantQueyResponseBody struct {
	// This parameter is required.
	Result []*MasterDataTenantQueyResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s MasterDataTenantQueyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MasterDataTenantQueyResponseBody) GoString() string {
	return s.String()
}

func (s *MasterDataTenantQueyResponseBody) SetResult(v []*MasterDataTenantQueyResponseBodyResult) *MasterDataTenantQueyResponseBody {
	s.Result = v
	return s
}

type MasterDataTenantQueyResponseBodyResult struct {
	// example:
	//
	// true
	HasData *bool `json:"hasData,omitempty" xml:"hasData,omitempty"`
	// example:
	//
	// true
	IntegrateDataAuth *bool `json:"integrateDataAuth,omitempty" xml:"integrateDataAuth,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// "智能绩效"
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// true
	ReadAuth *bool `json:"readAuth,omitempty" xml:"readAuth,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	TenantId *int64 `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s MasterDataTenantQueyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s MasterDataTenantQueyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *MasterDataTenantQueyResponseBodyResult) SetHasData(v bool) *MasterDataTenantQueyResponseBodyResult {
	s.HasData = &v
	return s
}

func (s *MasterDataTenantQueyResponseBodyResult) SetIntegrateDataAuth(v bool) *MasterDataTenantQueyResponseBodyResult {
	s.IntegrateDataAuth = &v
	return s
}

func (s *MasterDataTenantQueyResponseBodyResult) SetName(v string) *MasterDataTenantQueyResponseBodyResult {
	s.Name = &v
	return s
}

func (s *MasterDataTenantQueyResponseBodyResult) SetReadAuth(v bool) *MasterDataTenantQueyResponseBodyResult {
	s.ReadAuth = &v
	return s
}

func (s *MasterDataTenantQueyResponseBodyResult) SetTenantId(v int64) *MasterDataTenantQueyResponseBodyResult {
	s.TenantId = &v
	return s
}

func (s *MasterDataTenantQueyResponseBodyResult) SetType(v int32) *MasterDataTenantQueyResponseBodyResult {
	s.Type = &v
	return s
}

type MasterDataTenantQueyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MasterDataTenantQueyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MasterDataTenantQueyResponse) String() string {
	return tea.Prettify(s)
}

func (s MasterDataTenantQueyResponse) GoString() string {
	return s.String()
}

func (s *MasterDataTenantQueyResponse) SetHeaders(v map[string]*string) *MasterDataTenantQueyResponse {
	s.Headers = v
	return s
}

func (s *MasterDataTenantQueyResponse) SetStatusCode(v int32) *MasterDataTenantQueyResponse {
	s.StatusCode = &v
	return s
}

func (s *MasterDataTenantQueyResponse) SetBody(v *MasterDataTenantQueyResponseBody) *MasterDataTenantQueyResponse {
	s.Body = v
	return s
}

type MasterDatasGetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MasterDatasGetHeaders) String() string {
	return tea.Prettify(s)
}

func (s MasterDatasGetHeaders) GoString() string {
	return s.String()
}

func (s *MasterDatasGetHeaders) SetCommonHeaders(v map[string]*string) *MasterDatasGetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MasterDatasGetHeaders) SetXAcsDingtalkAccessToken(v string) *MasterDatasGetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MasterDatasGetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// uk1231
	ObjId *string `json:"objId,omitempty" xml:"objId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PERFORMANCE
	ScopeCode *string `json:"scopeCode,omitempty" xml:"scopeCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	TenantId *int64 `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// base
	ViewEntityCode *string `json:"viewEntityCode,omitempty" xml:"viewEntityCode,omitempty"`
}

func (s MasterDatasGetRequest) String() string {
	return tea.Prettify(s)
}

func (s MasterDatasGetRequest) GoString() string {
	return s.String()
}

func (s *MasterDatasGetRequest) SetObjId(v string) *MasterDatasGetRequest {
	s.ObjId = &v
	return s
}

func (s *MasterDatasGetRequest) SetScopeCode(v string) *MasterDatasGetRequest {
	s.ScopeCode = &v
	return s
}

func (s *MasterDatasGetRequest) SetTenantId(v int64) *MasterDatasGetRequest {
	s.TenantId = &v
	return s
}

func (s *MasterDatasGetRequest) SetViewEntityCode(v string) *MasterDatasGetRequest {
	s.ViewEntityCode = &v
	return s
}

type MasterDatasGetResponseBody struct {
	Result *MasterDatasGetResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s MasterDatasGetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MasterDatasGetResponseBody) GoString() string {
	return s.String()
}

func (s *MasterDatasGetResponseBody) SetResult(v *MasterDatasGetResponseBodyResult) *MasterDatasGetResponseBody {
	s.Result = v
	return s
}

type MasterDatasGetResponseBodyResult struct {
	// example:
	//
	// uk123123
	ObjId *string `json:"objId,omitempty" xml:"objId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// admind123
	RelationId *string `json:"relationId,omitempty" xml:"relationId,omitempty"`
	// example:
	//
	// PERFORMANCE
	ScopeCode *string `json:"scopeCode,omitempty" xml:"scopeCode,omitempty"`
	// example:
	//
	// base
	ViewEntityCode        *string                                                  `json:"viewEntityCode,omitempty" xml:"viewEntityCode,omitempty"`
	ViewEntityFieldVOList []*MasterDatasGetResponseBodyResultViewEntityFieldVOList `json:"viewEntityFieldVOList,omitempty" xml:"viewEntityFieldVOList,omitempty" type:"Repeated"`
}

func (s MasterDatasGetResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s MasterDatasGetResponseBodyResult) GoString() string {
	return s.String()
}

func (s *MasterDatasGetResponseBodyResult) SetObjId(v string) *MasterDatasGetResponseBodyResult {
	s.ObjId = &v
	return s
}

func (s *MasterDatasGetResponseBodyResult) SetRelationId(v string) *MasterDatasGetResponseBodyResult {
	s.RelationId = &v
	return s
}

func (s *MasterDatasGetResponseBodyResult) SetScopeCode(v string) *MasterDatasGetResponseBodyResult {
	s.ScopeCode = &v
	return s
}

func (s *MasterDatasGetResponseBodyResult) SetViewEntityCode(v string) *MasterDatasGetResponseBodyResult {
	s.ViewEntityCode = &v
	return s
}

func (s *MasterDatasGetResponseBodyResult) SetViewEntityFieldVOList(v []*MasterDatasGetResponseBodyResultViewEntityFieldVOList) *MasterDatasGetResponseBodyResult {
	s.ViewEntityFieldVOList = v
	return s
}

type MasterDatasGetResponseBodyResultViewEntityFieldVOList struct {
	// example:
	//
	// performanceValue
	FieldCode   *string                                                           `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	FieldDataVO *MasterDatasGetResponseBodyResultViewEntityFieldVOListFieldDataVO `json:"fieldDataVO,omitempty" xml:"fieldDataVO,omitempty" type:"Struct"`
	// example:
	//
	// 绩效等级
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// example:
	//
	// 1
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
}

func (s MasterDatasGetResponseBodyResultViewEntityFieldVOList) String() string {
	return tea.Prettify(s)
}

func (s MasterDatasGetResponseBodyResultViewEntityFieldVOList) GoString() string {
	return s.String()
}

func (s *MasterDatasGetResponseBodyResultViewEntityFieldVOList) SetFieldCode(v string) *MasterDatasGetResponseBodyResultViewEntityFieldVOList {
	s.FieldCode = &v
	return s
}

func (s *MasterDatasGetResponseBodyResultViewEntityFieldVOList) SetFieldDataVO(v *MasterDatasGetResponseBodyResultViewEntityFieldVOListFieldDataVO) *MasterDatasGetResponseBodyResultViewEntityFieldVOList {
	s.FieldDataVO = v
	return s
}

func (s *MasterDatasGetResponseBodyResultViewEntityFieldVOList) SetFieldName(v string) *MasterDatasGetResponseBodyResultViewEntityFieldVOList {
	s.FieldName = &v
	return s
}

func (s *MasterDatasGetResponseBodyResultViewEntityFieldVOList) SetFieldType(v string) *MasterDatasGetResponseBodyResultViewEntityFieldVOList {
	s.FieldType = &v
	return s
}

type MasterDatasGetResponseBodyResultViewEntityFieldVOListFieldDataVO struct {
	// example:
	//
	// 100
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// 100
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s MasterDatasGetResponseBodyResultViewEntityFieldVOListFieldDataVO) String() string {
	return tea.Prettify(s)
}

func (s MasterDatasGetResponseBodyResultViewEntityFieldVOListFieldDataVO) GoString() string {
	return s.String()
}

func (s *MasterDatasGetResponseBodyResultViewEntityFieldVOListFieldDataVO) SetKey(v string) *MasterDatasGetResponseBodyResultViewEntityFieldVOListFieldDataVO {
	s.Key = &v
	return s
}

func (s *MasterDatasGetResponseBodyResultViewEntityFieldVOListFieldDataVO) SetValue(v string) *MasterDatasGetResponseBodyResultViewEntityFieldVOListFieldDataVO {
	s.Value = &v
	return s
}

type MasterDatasGetResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MasterDatasGetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MasterDatasGetResponse) String() string {
	return tea.Prettify(s)
}

func (s MasterDatasGetResponse) GoString() string {
	return s.String()
}

func (s *MasterDatasGetResponse) SetHeaders(v map[string]*string) *MasterDatasGetResponse {
	s.Headers = v
	return s
}

func (s *MasterDatasGetResponse) SetStatusCode(v int32) *MasterDatasGetResponse {
	s.StatusCode = &v
	return s
}

func (s *MasterDatasGetResponse) SetBody(v *MasterDatasGetResponseBody) *MasterDatasGetResponse {
	s.Body = v
	return s
}

type MasterDatasQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MasterDatasQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s MasterDatasQueryHeaders) GoString() string {
	return s.String()
}

func (s *MasterDatasQueryHeaders) SetCommonHeaders(v map[string]*string) *MasterDatasQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MasterDatasQueryHeaders) SetXAcsDingtalkAccessToken(v string) *MasterDatasQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MasterDatasQueryRequest struct {
	// example:
	//
	// uk_12123
	BizUK *string `json:"bizUK,omitempty" xml:"bizUK,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	NextToken   *int32                                `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	QueryParams []*MasterDatasQueryRequestQueryParams `json:"queryParams,omitempty" xml:"queryParams,omitempty" type:"Repeated"`
	// This parameter is required.
	RelationIds []*string `json:"relationIds,omitempty" xml:"relationIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// PERFORMANCE
	ScopeCode *string `json:"scopeCode,omitempty" xml:"scopeCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	TenantId *int64 `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// base
	ViewEntityCode *string `json:"viewEntityCode,omitempty" xml:"viewEntityCode,omitempty"`
}

func (s MasterDatasQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s MasterDatasQueryRequest) GoString() string {
	return s.String()
}

func (s *MasterDatasQueryRequest) SetBizUK(v string) *MasterDatasQueryRequest {
	s.BizUK = &v
	return s
}

func (s *MasterDatasQueryRequest) SetMaxResults(v int32) *MasterDatasQueryRequest {
	s.MaxResults = &v
	return s
}

func (s *MasterDatasQueryRequest) SetNextToken(v int32) *MasterDatasQueryRequest {
	s.NextToken = &v
	return s
}

func (s *MasterDatasQueryRequest) SetQueryParams(v []*MasterDatasQueryRequestQueryParams) *MasterDatasQueryRequest {
	s.QueryParams = v
	return s
}

func (s *MasterDatasQueryRequest) SetRelationIds(v []*string) *MasterDatasQueryRequest {
	s.RelationIds = v
	return s
}

func (s *MasterDatasQueryRequest) SetScopeCode(v string) *MasterDatasQueryRequest {
	s.ScopeCode = &v
	return s
}

func (s *MasterDatasQueryRequest) SetTenantId(v int64) *MasterDatasQueryRequest {
	s.TenantId = &v
	return s
}

func (s *MasterDatasQueryRequest) SetViewEntityCode(v string) *MasterDatasQueryRequest {
	s.ViewEntityCode = &v
	return s
}

type MasterDatasQueryRequestQueryParams struct {
	ConditionList []*MasterDatasQueryRequestQueryParamsConditionList `json:"conditionList,omitempty" xml:"conditionList,omitempty" type:"Repeated"`
	// example:
	//
	// performance_code
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	// example:
	//
	// AND
	JoinType *string `json:"joinType,omitempty" xml:"joinType,omitempty"`
}

func (s MasterDatasQueryRequestQueryParams) String() string {
	return tea.Prettify(s)
}

func (s MasterDatasQueryRequestQueryParams) GoString() string {
	return s.String()
}

func (s *MasterDatasQueryRequestQueryParams) SetConditionList(v []*MasterDatasQueryRequestQueryParamsConditionList) *MasterDatasQueryRequestQueryParams {
	s.ConditionList = v
	return s
}

func (s *MasterDatasQueryRequestQueryParams) SetFieldCode(v string) *MasterDatasQueryRequestQueryParams {
	s.FieldCode = &v
	return s
}

func (s *MasterDatasQueryRequestQueryParams) SetJoinType(v string) *MasterDatasQueryRequestQueryParams {
	s.JoinType = &v
	return s
}

type MasterDatasQueryRequestQueryParamsConditionList struct {
	// example:
	//
	// EQUAL
	Operate *string `json:"operate,omitempty" xml:"operate,omitempty"`
	Value   *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s MasterDatasQueryRequestQueryParamsConditionList) String() string {
	return tea.Prettify(s)
}

func (s MasterDatasQueryRequestQueryParamsConditionList) GoString() string {
	return s.String()
}

func (s *MasterDatasQueryRequestQueryParamsConditionList) SetOperate(v string) *MasterDatasQueryRequestQueryParamsConditionList {
	s.Operate = &v
	return s
}

func (s *MasterDatasQueryRequestQueryParamsConditionList) SetValue(v string) *MasterDatasQueryRequestQueryParamsConditionList {
	s.Value = &v
	return s
}

type MasterDatasQueryResponseBody struct {
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 0
	NextToken *int64                                `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Result    []*MasterDatasQueryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// This parameter is required.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 100
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s MasterDatasQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MasterDatasQueryResponseBody) GoString() string {
	return s.String()
}

func (s *MasterDatasQueryResponseBody) SetHasMore(v bool) *MasterDatasQueryResponseBody {
	s.HasMore = &v
	return s
}

func (s *MasterDatasQueryResponseBody) SetNextToken(v int64) *MasterDatasQueryResponseBody {
	s.NextToken = &v
	return s
}

func (s *MasterDatasQueryResponseBody) SetResult(v []*MasterDatasQueryResponseBodyResult) *MasterDatasQueryResponseBody {
	s.Result = v
	return s
}

func (s *MasterDatasQueryResponseBody) SetSuccess(v bool) *MasterDatasQueryResponseBody {
	s.Success = &v
	return s
}

func (s *MasterDatasQueryResponseBody) SetTotal(v int64) *MasterDatasQueryResponseBody {
	s.Total = &v
	return s
}

type MasterDatasQueryResponseBodyResult struct {
	// example:
	//
	// uk123123
	ObjId *string `json:"objId,omitempty" xml:"objId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// admind123
	RelationId *string `json:"relationId,omitempty" xml:"relationId,omitempty"`
	// example:
	//
	// PERFORMANCE
	ScopeCode *string `json:"scopeCode,omitempty" xml:"scopeCode,omitempty"`
	// example:
	//
	// base
	ViewEntityCode        *string                                                    `json:"viewEntityCode,omitempty" xml:"viewEntityCode,omitempty"`
	ViewEntityFieldVOList []*MasterDatasQueryResponseBodyResultViewEntityFieldVOList `json:"viewEntityFieldVOList,omitempty" xml:"viewEntityFieldVOList,omitempty" type:"Repeated"`
}

func (s MasterDatasQueryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s MasterDatasQueryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *MasterDatasQueryResponseBodyResult) SetObjId(v string) *MasterDatasQueryResponseBodyResult {
	s.ObjId = &v
	return s
}

func (s *MasterDatasQueryResponseBodyResult) SetRelationId(v string) *MasterDatasQueryResponseBodyResult {
	s.RelationId = &v
	return s
}

func (s *MasterDatasQueryResponseBodyResult) SetScopeCode(v string) *MasterDatasQueryResponseBodyResult {
	s.ScopeCode = &v
	return s
}

func (s *MasterDatasQueryResponseBodyResult) SetViewEntityCode(v string) *MasterDatasQueryResponseBodyResult {
	s.ViewEntityCode = &v
	return s
}

func (s *MasterDatasQueryResponseBodyResult) SetViewEntityFieldVOList(v []*MasterDatasQueryResponseBodyResultViewEntityFieldVOList) *MasterDatasQueryResponseBodyResult {
	s.ViewEntityFieldVOList = v
	return s
}

type MasterDatasQueryResponseBodyResultViewEntityFieldVOList struct {
	// example:
	//
	// performanceValue
	FieldCode   *string                                                             `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	FieldDataVO *MasterDatasQueryResponseBodyResultViewEntityFieldVOListFieldDataVO `json:"fieldDataVO,omitempty" xml:"fieldDataVO,omitempty" type:"Struct"`
	// example:
	//
	// 绩效等级
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// example:
	//
	// 1
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
}

func (s MasterDatasQueryResponseBodyResultViewEntityFieldVOList) String() string {
	return tea.Prettify(s)
}

func (s MasterDatasQueryResponseBodyResultViewEntityFieldVOList) GoString() string {
	return s.String()
}

func (s *MasterDatasQueryResponseBodyResultViewEntityFieldVOList) SetFieldCode(v string) *MasterDatasQueryResponseBodyResultViewEntityFieldVOList {
	s.FieldCode = &v
	return s
}

func (s *MasterDatasQueryResponseBodyResultViewEntityFieldVOList) SetFieldDataVO(v *MasterDatasQueryResponseBodyResultViewEntityFieldVOListFieldDataVO) *MasterDatasQueryResponseBodyResultViewEntityFieldVOList {
	s.FieldDataVO = v
	return s
}

func (s *MasterDatasQueryResponseBodyResultViewEntityFieldVOList) SetFieldName(v string) *MasterDatasQueryResponseBodyResultViewEntityFieldVOList {
	s.FieldName = &v
	return s
}

func (s *MasterDatasQueryResponseBodyResultViewEntityFieldVOList) SetFieldType(v string) *MasterDatasQueryResponseBodyResultViewEntityFieldVOList {
	s.FieldType = &v
	return s
}

type MasterDatasQueryResponseBodyResultViewEntityFieldVOListFieldDataVO struct {
	// example:
	//
	// 100
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// 100
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s MasterDatasQueryResponseBodyResultViewEntityFieldVOListFieldDataVO) String() string {
	return tea.Prettify(s)
}

func (s MasterDatasQueryResponseBodyResultViewEntityFieldVOListFieldDataVO) GoString() string {
	return s.String()
}

func (s *MasterDatasQueryResponseBodyResultViewEntityFieldVOListFieldDataVO) SetKey(v string) *MasterDatasQueryResponseBodyResultViewEntityFieldVOListFieldDataVO {
	s.Key = &v
	return s
}

func (s *MasterDatasQueryResponseBodyResultViewEntityFieldVOListFieldDataVO) SetValue(v string) *MasterDatasQueryResponseBodyResultViewEntityFieldVOListFieldDataVO {
	s.Value = &v
	return s
}

type MasterDatasQueryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MasterDatasQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MasterDatasQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s MasterDatasQueryResponse) GoString() string {
	return s.String()
}

func (s *MasterDatasQueryResponse) SetHeaders(v map[string]*string) *MasterDatasQueryResponse {
	s.Headers = v
	return s
}

func (s *MasterDatasQueryResponse) SetStatusCode(v int32) *MasterDatasQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *MasterDatasQueryResponse) SetBody(v *MasterDatasQueryResponseBody) *MasterDatasQueryResponse {
	s.Body = v
	return s
}

type OpenOemMicroAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s OpenOemMicroAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s OpenOemMicroAppHeaders) GoString() string {
	return s.String()
}

func (s *OpenOemMicroAppHeaders) SetCommonHeaders(v map[string]*string) *OpenOemMicroAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OpenOemMicroAppHeaders) SetXAcsDingtalkAccessToken(v string) *OpenOemMicroAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type OpenOemMicroAppRequest struct {
	// example:
	//
	// 12
	TenantId *int64 `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s OpenOemMicroAppRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenOemMicroAppRequest) GoString() string {
	return s.String()
}

func (s *OpenOemMicroAppRequest) SetTenantId(v int64) *OpenOemMicroAppRequest {
	s.TenantId = &v
	return s
}

type OpenOemMicroAppResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OpenOemMicroAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenOemMicroAppResponseBody) GoString() string {
	return s.String()
}

func (s *OpenOemMicroAppResponseBody) SetRequestId(v string) *OpenOemMicroAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenOemMicroAppResponseBody) SetSuccess(v bool) *OpenOemMicroAppResponseBody {
	s.Success = &v
	return s
}

type OpenOemMicroAppResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenOemMicroAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenOemMicroAppResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenOemMicroAppResponse) GoString() string {
	return s.String()
}

func (s *OpenOemMicroAppResponse) SetHeaders(v map[string]*string) *OpenOemMicroAppResponse {
	s.Headers = v
	return s
}

func (s *OpenOemMicroAppResponse) SetStatusCode(v int32) *OpenOemMicroAppResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenOemMicroAppResponse) SetBody(v *OpenOemMicroAppResponseBody) *OpenOemMicroAppResponse {
	s.Body = v
	return s
}

type QueryCustomEntryProcessesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCustomEntryProcessesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomEntryProcessesHeaders) GoString() string {
	return s.String()
}

func (s *QueryCustomEntryProcessesHeaders) SetCommonHeaders(v map[string]*string) *QueryCustomEntryProcessesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCustomEntryProcessesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCustomEntryProcessesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCustomEntryProcessesRequest struct {
	// example:
	//
	// 20，最大为100，不填默认为100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 默认为0
	NextToken     *int32  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OperateUserId *string `json:"operateUserId,omitempty" xml:"operateUserId,omitempty"`
}

func (s QueryCustomEntryProcessesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomEntryProcessesRequest) GoString() string {
	return s.String()
}

func (s *QueryCustomEntryProcessesRequest) SetMaxResults(v int32) *QueryCustomEntryProcessesRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryCustomEntryProcessesRequest) SetNextToken(v int32) *QueryCustomEntryProcessesRequest {
	s.NextToken = &v
	return s
}

func (s *QueryCustomEntryProcessesRequest) SetOperateUserId(v string) *QueryCustomEntryProcessesRequest {
	s.OperateUserId = &v
	return s
}

type QueryCustomEntryProcessesResponseBody struct {
	HasMore   *bool                                        `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List      []*QueryCustomEntryProcessesResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextToken *int64                                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryCustomEntryProcessesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomEntryProcessesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCustomEntryProcessesResponseBody) SetHasMore(v bool) *QueryCustomEntryProcessesResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryCustomEntryProcessesResponseBody) SetList(v []*QueryCustomEntryProcessesResponseBodyList) *QueryCustomEntryProcessesResponseBody {
	s.List = v
	return s
}

func (s *QueryCustomEntryProcessesResponseBody) SetNextToken(v int64) *QueryCustomEntryProcessesResponseBody {
	s.NextToken = &v
	return s
}

type QueryCustomEntryProcessesResponseBodyList struct {
	FormDesc *string `json:"formDesc,omitempty" xml:"formDesc,omitempty"`
	FormId   *string `json:"formId,omitempty" xml:"formId,omitempty"`
	FormName *string `json:"formName,omitempty" xml:"formName,omitempty"`
	ShortUrl *string `json:"shortUrl,omitempty" xml:"shortUrl,omitempty"`
}

func (s QueryCustomEntryProcessesResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomEntryProcessesResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryCustomEntryProcessesResponseBodyList) SetFormDesc(v string) *QueryCustomEntryProcessesResponseBodyList {
	s.FormDesc = &v
	return s
}

func (s *QueryCustomEntryProcessesResponseBodyList) SetFormId(v string) *QueryCustomEntryProcessesResponseBodyList {
	s.FormId = &v
	return s
}

func (s *QueryCustomEntryProcessesResponseBodyList) SetFormName(v string) *QueryCustomEntryProcessesResponseBodyList {
	s.FormName = &v
	return s
}

func (s *QueryCustomEntryProcessesResponseBodyList) SetShortUrl(v string) *QueryCustomEntryProcessesResponseBodyList {
	s.ShortUrl = &v
	return s
}

type QueryCustomEntryProcessesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCustomEntryProcessesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCustomEntryProcessesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomEntryProcessesResponse) GoString() string {
	return s.String()
}

func (s *QueryCustomEntryProcessesResponse) SetHeaders(v map[string]*string) *QueryCustomEntryProcessesResponse {
	s.Headers = v
	return s
}

func (s *QueryCustomEntryProcessesResponse) SetStatusCode(v int32) *QueryCustomEntryProcessesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCustomEntryProcessesResponse) SetBody(v *QueryCustomEntryProcessesResponseBody) *QueryCustomEntryProcessesResponse {
	s.Body = v
	return s
}

type QueryDismissionStaffIdListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDismissionStaffIdListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDismissionStaffIdListHeaders) GoString() string {
	return s.String()
}

func (s *QueryDismissionStaffIdListHeaders) SetCommonHeaders(v map[string]*string) *QueryDismissionStaffIdListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDismissionStaffIdListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDismissionStaffIdListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDismissionStaffIdListRequest struct {
	// example:
	//
	// 30
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 0
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryDismissionStaffIdListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDismissionStaffIdListRequest) GoString() string {
	return s.String()
}

func (s *QueryDismissionStaffIdListRequest) SetMaxResults(v int32) *QueryDismissionStaffIdListRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryDismissionStaffIdListRequest) SetNextToken(v int64) *QueryDismissionStaffIdListRequest {
	s.NextToken = &v
	return s
}

type QueryDismissionStaffIdListResponseBody struct {
	HasMore    *bool     `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken  *int64    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s QueryDismissionStaffIdListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDismissionStaffIdListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDismissionStaffIdListResponseBody) SetHasMore(v bool) *QueryDismissionStaffIdListResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryDismissionStaffIdListResponseBody) SetNextToken(v int64) *QueryDismissionStaffIdListResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryDismissionStaffIdListResponseBody) SetUserIdList(v []*string) *QueryDismissionStaffIdListResponseBody {
	s.UserIdList = v
	return s
}

type QueryDismissionStaffIdListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDismissionStaffIdListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDismissionStaffIdListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDismissionStaffIdListResponse) GoString() string {
	return s.String()
}

func (s *QueryDismissionStaffIdListResponse) SetHeaders(v map[string]*string) *QueryDismissionStaffIdListResponse {
	s.Headers = v
	return s
}

func (s *QueryDismissionStaffIdListResponse) SetStatusCode(v int32) *QueryDismissionStaffIdListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDismissionStaffIdListResponse) SetBody(v *QueryDismissionStaffIdListResponseBody) *QueryDismissionStaffIdListResponse {
	s.Body = v
	return s
}

type QueryHrmEmployeeDismissionInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryHrmEmployeeDismissionInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryHrmEmployeeDismissionInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryHrmEmployeeDismissionInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryHrmEmployeeDismissionInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryHrmEmployeeDismissionInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryHrmEmployeeDismissionInfoRequest struct {
	// This parameter is required.
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s QueryHrmEmployeeDismissionInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHrmEmployeeDismissionInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryHrmEmployeeDismissionInfoRequest) SetUserIdList(v []*string) *QueryHrmEmployeeDismissionInfoRequest {
	s.UserIdList = v
	return s
}

type QueryHrmEmployeeDismissionInfoShrinkRequest struct {
	// This parameter is required.
	UserIdListShrink *string `json:"userIdList,omitempty" xml:"userIdList,omitempty"`
}

func (s QueryHrmEmployeeDismissionInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHrmEmployeeDismissionInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryHrmEmployeeDismissionInfoShrinkRequest) SetUserIdListShrink(v string) *QueryHrmEmployeeDismissionInfoShrinkRequest {
	s.UserIdListShrink = &v
	return s
}

type QueryHrmEmployeeDismissionInfoResponseBody struct {
	Result []*QueryHrmEmployeeDismissionInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryHrmEmployeeDismissionInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHrmEmployeeDismissionInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHrmEmployeeDismissionInfoResponseBody) SetResult(v []*QueryHrmEmployeeDismissionInfoResponseBodyResult) *QueryHrmEmployeeDismissionInfoResponseBody {
	s.Result = v
	return s
}

type QueryHrmEmployeeDismissionInfoResponseBodyResult struct {
	DeptList        []*QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList `json:"deptList,omitempty" xml:"deptList,omitempty" type:"Repeated"`
	HandoverUserId  *string                                                     `json:"handoverUserId,omitempty" xml:"handoverUserId,omitempty"`
	LastWorkDay     *int64                                                      `json:"lastWorkDay,omitempty" xml:"lastWorkDay,omitempty"`
	MainDeptId      *int64                                                      `json:"mainDeptId,omitempty" xml:"mainDeptId,omitempty"`
	MainDeptName    *string                                                     `json:"mainDeptName,omitempty" xml:"mainDeptName,omitempty"`
	Name            *string                                                     `json:"name,omitempty" xml:"name,omitempty"`
	PassiveReason   []*string                                                   `json:"passiveReason,omitempty" xml:"passiveReason,omitempty" type:"Repeated"`
	PreStatus       *int32                                                      `json:"preStatus,omitempty" xml:"preStatus,omitempty"`
	ReasonMemo      *string                                                     `json:"reasonMemo,omitempty" xml:"reasonMemo,omitempty"`
	Status          *int32                                                      `json:"status,omitempty" xml:"status,omitempty"`
	UserId          *string                                                     `json:"userId,omitempty" xml:"userId,omitempty"`
	VoluntaryReason []*string                                                   `json:"voluntaryReason,omitempty" xml:"voluntaryReason,omitempty" type:"Repeated"`
}

func (s QueryHrmEmployeeDismissionInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryHrmEmployeeDismissionInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetDeptList(v []*QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.DeptList = v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetHandoverUserId(v string) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.HandoverUserId = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetLastWorkDay(v int64) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.LastWorkDay = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetMainDeptId(v int64) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.MainDeptId = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetMainDeptName(v string) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.MainDeptName = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetName(v string) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.Name = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetPassiveReason(v []*string) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.PassiveReason = v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetPreStatus(v int32) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.PreStatus = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetReasonMemo(v string) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.ReasonMemo = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetStatus(v int32) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.Status = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetUserId(v string) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResult) SetVoluntaryReason(v []*string) *QueryHrmEmployeeDismissionInfoResponseBodyResult {
	s.VoluntaryReason = v
	return s
}

type QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList struct {
	DeptId   *int64  `json:"dept_id,omitempty" xml:"dept_id,omitempty"`
	DeptPath *string `json:"dept_path,omitempty" xml:"dept_path,omitempty"`
}

func (s QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList) String() string {
	return tea.Prettify(s)
}

func (s QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList) GoString() string {
	return s.String()
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList) SetDeptId(v int64) *QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList {
	s.DeptId = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList) SetDeptPath(v string) *QueryHrmEmployeeDismissionInfoResponseBodyResultDeptList {
	s.DeptPath = &v
	return s
}

type QueryHrmEmployeeDismissionInfoResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryHrmEmployeeDismissionInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryHrmEmployeeDismissionInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHrmEmployeeDismissionInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryHrmEmployeeDismissionInfoResponse) SetHeaders(v map[string]*string) *QueryHrmEmployeeDismissionInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponse) SetStatusCode(v int32) *QueryHrmEmployeeDismissionInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHrmEmployeeDismissionInfoResponse) SetBody(v *QueryHrmEmployeeDismissionInfoResponseBody) *QueryHrmEmployeeDismissionInfoResponse {
	s.Body = v
	return s
}

type QueryJobRanksHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryJobRanksHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryJobRanksHeaders) GoString() string {
	return s.String()
}

func (s *QueryJobRanksHeaders) SetCommonHeaders(v map[string]*string) *QueryJobRanksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryJobRanksHeaders) SetXAcsDingtalkAccessToken(v string) *QueryJobRanksHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryJobRanksRequest struct {
	// This parameter is required.
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	NextToken      *int32  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RankCategoryId *string `json:"rankCategoryId,omitempty" xml:"rankCategoryId,omitempty"`
	RankCode       *string `json:"rankCode,omitempty" xml:"rankCode,omitempty"`
	RankName       *string `json:"rankName,omitempty" xml:"rankName,omitempty"`
}

func (s QueryJobRanksRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryJobRanksRequest) GoString() string {
	return s.String()
}

func (s *QueryJobRanksRequest) SetMaxResults(v int32) *QueryJobRanksRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryJobRanksRequest) SetNextToken(v int32) *QueryJobRanksRequest {
	s.NextToken = &v
	return s
}

func (s *QueryJobRanksRequest) SetRankCategoryId(v string) *QueryJobRanksRequest {
	s.RankCategoryId = &v
	return s
}

func (s *QueryJobRanksRequest) SetRankCode(v string) *QueryJobRanksRequest {
	s.RankCode = &v
	return s
}

func (s *QueryJobRanksRequest) SetRankName(v string) *QueryJobRanksRequest {
	s.RankName = &v
	return s
}

type QueryJobRanksResponseBody struct {
	HasMore   *bool                            `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List      []*QueryJobRanksResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextToken *int64                           `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryJobRanksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryJobRanksResponseBody) GoString() string {
	return s.String()
}

func (s *QueryJobRanksResponseBody) SetHasMore(v bool) *QueryJobRanksResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryJobRanksResponseBody) SetList(v []*QueryJobRanksResponseBodyList) *QueryJobRanksResponseBody {
	s.List = v
	return s
}

func (s *QueryJobRanksResponseBody) SetNextToken(v int64) *QueryJobRanksResponseBody {
	s.NextToken = &v
	return s
}

type QueryJobRanksResponseBodyList struct {
	// example:
	//
	// 30
	MaxJobGrade *int32 `json:"maxJobGrade,omitempty" xml:"maxJobGrade,omitempty"`
	// example:
	//
	// 1
	MinJobGrade     *int32  `json:"minJobGrade,omitempty" xml:"minJobGrade,omitempty"`
	RankCategoryId  *string `json:"rankCategoryId,omitempty" xml:"rankCategoryId,omitempty"`
	RankCode        *string `json:"rankCode,omitempty" xml:"rankCode,omitempty"`
	RankDescription *string `json:"rankDescription,omitempty" xml:"rankDescription,omitempty"`
	RankId          *string `json:"rankId,omitempty" xml:"rankId,omitempty"`
	RankName        *string `json:"rankName,omitempty" xml:"rankName,omitempty"`
}

func (s QueryJobRanksResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryJobRanksResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryJobRanksResponseBodyList) SetMaxJobGrade(v int32) *QueryJobRanksResponseBodyList {
	s.MaxJobGrade = &v
	return s
}

func (s *QueryJobRanksResponseBodyList) SetMinJobGrade(v int32) *QueryJobRanksResponseBodyList {
	s.MinJobGrade = &v
	return s
}

func (s *QueryJobRanksResponseBodyList) SetRankCategoryId(v string) *QueryJobRanksResponseBodyList {
	s.RankCategoryId = &v
	return s
}

func (s *QueryJobRanksResponseBodyList) SetRankCode(v string) *QueryJobRanksResponseBodyList {
	s.RankCode = &v
	return s
}

func (s *QueryJobRanksResponseBodyList) SetRankDescription(v string) *QueryJobRanksResponseBodyList {
	s.RankDescription = &v
	return s
}

func (s *QueryJobRanksResponseBodyList) SetRankId(v string) *QueryJobRanksResponseBodyList {
	s.RankId = &v
	return s
}

func (s *QueryJobRanksResponseBodyList) SetRankName(v string) *QueryJobRanksResponseBodyList {
	s.RankName = &v
	return s
}

type QueryJobRanksResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryJobRanksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryJobRanksResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryJobRanksResponse) GoString() string {
	return s.String()
}

func (s *QueryJobRanksResponse) SetHeaders(v map[string]*string) *QueryJobRanksResponse {
	s.Headers = v
	return s
}

func (s *QueryJobRanksResponse) SetStatusCode(v int32) *QueryJobRanksResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryJobRanksResponse) SetBody(v *QueryJobRanksResponseBody) *QueryJobRanksResponse {
	s.Body = v
	return s
}

type QueryJobsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryJobsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryJobsHeaders) GoString() string {
	return s.String()
}

func (s *QueryJobsHeaders) SetCommonHeaders(v map[string]*string) *QueryJobsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryJobsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryJobsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryJobsRequest struct {
	// example:
	//
	// 工程师
	JobName *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20，最大为100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	NextToken *int32 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryJobsRequest) GoString() string {
	return s.String()
}

func (s *QueryJobsRequest) SetJobName(v string) *QueryJobsRequest {
	s.JobName = &v
	return s
}

func (s *QueryJobsRequest) SetMaxResults(v int32) *QueryJobsRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryJobsRequest) SetNextToken(v int32) *QueryJobsRequest {
	s.NextToken = &v
	return s
}

type QueryJobsResponseBody struct {
	HasMore   *bool                        `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List      []*QueryJobsResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextToken *int64                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryJobsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryJobsResponseBody) SetHasMore(v bool) *QueryJobsResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryJobsResponseBody) SetList(v []*QueryJobsResponseBodyList) *QueryJobsResponseBody {
	s.List = v
	return s
}

func (s *QueryJobsResponseBody) SetNextToken(v int64) *QueryJobsResponseBody {
	s.NextToken = &v
	return s
}

type QueryJobsResponseBodyList struct {
	// example:
	//
	// 职务描述
	JobDescription *string `json:"jobDescription,omitempty" xml:"jobDescription,omitempty"`
	// example:
	//
	// ac67286db74c48e28d787173ccc1a111
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// example:
	//
	// 总裁
	JobName *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
}

func (s QueryJobsResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryJobsResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryJobsResponseBodyList) SetJobDescription(v string) *QueryJobsResponseBodyList {
	s.JobDescription = &v
	return s
}

func (s *QueryJobsResponseBodyList) SetJobId(v string) *QueryJobsResponseBodyList {
	s.JobId = &v
	return s
}

func (s *QueryJobsResponseBodyList) SetJobName(v string) *QueryJobsResponseBodyList {
	s.JobName = &v
	return s
}

type QueryJobsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryJobsResponse) GoString() string {
	return s.String()
}

func (s *QueryJobsResponse) SetHeaders(v map[string]*string) *QueryJobsResponse {
	s.Headers = v
	return s
}

func (s *QueryJobsResponse) SetStatusCode(v int32) *QueryJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryJobsResponse) SetBody(v *QueryJobsResponseBody) *QueryJobsResponse {
	s.Body = v
	return s
}

type QueryMicroAppStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMicroAppStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMicroAppStatusHeaders) GoString() string {
	return s.String()
}

func (s *QueryMicroAppStatusHeaders) SetCommonHeaders(v map[string]*string) *QueryMicroAppStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMicroAppStatusHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMicroAppStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMicroAppStatusRequest struct {
	TenantIdList []*int64 `json:"tenantIdList,omitempty" xml:"tenantIdList,omitempty" type:"Repeated"`
}

func (s QueryMicroAppStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMicroAppStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryMicroAppStatusRequest) SetTenantIdList(v []*int64) *QueryMicroAppStatusRequest {
	s.TenantIdList = v
	return s
}

type QueryMicroAppStatusResponseBody struct {
	RequestId *string                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]*ResultValue `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool                   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryMicroAppStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMicroAppStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMicroAppStatusResponseBody) SetRequestId(v string) *QueryMicroAppStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMicroAppStatusResponseBody) SetResult(v map[string]*ResultValue) *QueryMicroAppStatusResponseBody {
	s.Result = v
	return s
}

func (s *QueryMicroAppStatusResponseBody) SetSuccess(v bool) *QueryMicroAppStatusResponseBody {
	s.Success = &v
	return s
}

type QueryMicroAppStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMicroAppStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMicroAppStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMicroAppStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryMicroAppStatusResponse) SetHeaders(v map[string]*string) *QueryMicroAppStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryMicroAppStatusResponse) SetStatusCode(v int32) *QueryMicroAppStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMicroAppStatusResponse) SetBody(v *QueryMicroAppStatusResponseBody) *QueryMicroAppStatusResponse {
	s.Body = v
	return s
}

type QueryMicroAppViewHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMicroAppViewHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMicroAppViewHeaders) GoString() string {
	return s.String()
}

func (s *QueryMicroAppViewHeaders) SetCommonHeaders(v map[string]*string) *QueryMicroAppViewHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMicroAppViewHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMicroAppViewHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMicroAppViewRequest struct {
	TenantIdList []*int64 `json:"tenantIdList,omitempty" xml:"tenantIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 2163515669935611
	ViewUserId *string `json:"viewUserId,omitempty" xml:"viewUserId,omitempty"`
}

func (s QueryMicroAppViewRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMicroAppViewRequest) GoString() string {
	return s.String()
}

func (s *QueryMicroAppViewRequest) SetTenantIdList(v []*int64) *QueryMicroAppViewRequest {
	s.TenantIdList = v
	return s
}

func (s *QueryMicroAppViewRequest) SetViewUserId(v string) *QueryMicroAppViewRequest {
	s.ViewUserId = &v
	return s
}

type QueryMicroAppViewResponseBody struct {
	RequestId *string          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]*bool `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool            `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryMicroAppViewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMicroAppViewResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMicroAppViewResponseBody) SetRequestId(v string) *QueryMicroAppViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMicroAppViewResponseBody) SetResult(v map[string]*bool) *QueryMicroAppViewResponseBody {
	s.Result = v
	return s
}

func (s *QueryMicroAppViewResponseBody) SetSuccess(v bool) *QueryMicroAppViewResponseBody {
	s.Success = &v
	return s
}

type QueryMicroAppViewResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMicroAppViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMicroAppViewResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMicroAppViewResponse) GoString() string {
	return s.String()
}

func (s *QueryMicroAppViewResponse) SetHeaders(v map[string]*string) *QueryMicroAppViewResponse {
	s.Headers = v
	return s
}

func (s *QueryMicroAppViewResponse) SetStatusCode(v int32) *QueryMicroAppViewResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMicroAppViewResponse) SetBody(v *QueryMicroAppViewResponseBody) *QueryMicroAppViewResponse {
	s.Body = v
	return s
}

type QueryPositionVersionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryPositionVersionHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryPositionVersionHeaders) GoString() string {
	return s.String()
}

func (s *QueryPositionVersionHeaders) SetCommonHeaders(v map[string]*string) *QueryPositionVersionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryPositionVersionHeaders) SetXAcsDingtalkAccessToken(v string) *QueryPositionVersionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryPositionVersionResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s QueryPositionVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPositionVersionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPositionVersionResponseBody) SetResult(v string) *QueryPositionVersionResponseBody {
	s.Result = &v
	return s
}

type QueryPositionVersionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPositionVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPositionVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPositionVersionResponse) GoString() string {
	return s.String()
}

func (s *QueryPositionVersionResponse) SetHeaders(v map[string]*string) *QueryPositionVersionResponse {
	s.Headers = v
	return s
}

func (s *QueryPositionVersionResponse) SetStatusCode(v int32) *QueryPositionVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPositionVersionResponse) SetBody(v *QueryPositionVersionResponseBody) *QueryPositionVersionResponse {
	s.Body = v
	return s
}

type QueryPositionsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryPositionsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryPositionsHeaders) GoString() string {
	return s.String()
}

func (s *QueryPositionsHeaders) SetCommonHeaders(v map[string]*string) *QueryPositionsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryPositionsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryPositionsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryPositionsRequest struct {
	// example:
	//
	// 部门id
	DeptId        *int64    `json:"deptId,omitempty" xml:"deptId,omitempty"`
	InCategoryIds []*string `json:"inCategoryIds,omitempty" xml:"inCategoryIds,omitempty" type:"Repeated"`
	InPositionIds []*string `json:"inPositionIds,omitempty" xml:"inPositionIds,omitempty" type:"Repeated"`
	// example:
	//
	// 职位名称
	PositionName *string `json:"positionName,omitempty" xml:"positionName,omitempty"`
	// This parameter is required.
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	NextToken *int32 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryPositionsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPositionsRequest) GoString() string {
	return s.String()
}

func (s *QueryPositionsRequest) SetDeptId(v int64) *QueryPositionsRequest {
	s.DeptId = &v
	return s
}

func (s *QueryPositionsRequest) SetInCategoryIds(v []*string) *QueryPositionsRequest {
	s.InCategoryIds = v
	return s
}

func (s *QueryPositionsRequest) SetInPositionIds(v []*string) *QueryPositionsRequest {
	s.InPositionIds = v
	return s
}

func (s *QueryPositionsRequest) SetPositionName(v string) *QueryPositionsRequest {
	s.PositionName = &v
	return s
}

func (s *QueryPositionsRequest) SetMaxResults(v int32) *QueryPositionsRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryPositionsRequest) SetNextToken(v int32) *QueryPositionsRequest {
	s.NextToken = &v
	return s
}

type QueryPositionsResponseBody struct {
	HasMore   *bool                             `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List      []*QueryPositionsResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextToken *int64                            `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryPositionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPositionsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPositionsResponseBody) SetHasMore(v bool) *QueryPositionsResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryPositionsResponseBody) SetList(v []*QueryPositionsResponseBodyList) *QueryPositionsResponseBody {
	s.List = v
	return s
}

func (s *QueryPositionsResponseBody) SetNextToken(v int64) *QueryPositionsResponseBody {
	s.NextToken = &v
	return s
}

type QueryPositionsResponseBodyList struct {
	JobId              *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	PositionCategoryId *string `json:"positionCategoryId,omitempty" xml:"positionCategoryId,omitempty"`
	PositionDes        *string `json:"positionDes,omitempty" xml:"positionDes,omitempty"`
	// example:
	//
	// ac67286db74c48e28d787173ccc1a111
	PositionId   *string   `json:"positionId,omitempty" xml:"positionId,omitempty"`
	PositionName *string   `json:"positionName,omitempty" xml:"positionName,omitempty"`
	RankIdList   []*string `json:"rankIdList,omitempty" xml:"rankIdList,omitempty" type:"Repeated"`
	Status       *int32    `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryPositionsResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryPositionsResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryPositionsResponseBodyList) SetJobId(v string) *QueryPositionsResponseBodyList {
	s.JobId = &v
	return s
}

func (s *QueryPositionsResponseBodyList) SetPositionCategoryId(v string) *QueryPositionsResponseBodyList {
	s.PositionCategoryId = &v
	return s
}

func (s *QueryPositionsResponseBodyList) SetPositionDes(v string) *QueryPositionsResponseBodyList {
	s.PositionDes = &v
	return s
}

func (s *QueryPositionsResponseBodyList) SetPositionId(v string) *QueryPositionsResponseBodyList {
	s.PositionId = &v
	return s
}

func (s *QueryPositionsResponseBodyList) SetPositionName(v string) *QueryPositionsResponseBodyList {
	s.PositionName = &v
	return s
}

func (s *QueryPositionsResponseBodyList) SetRankIdList(v []*string) *QueryPositionsResponseBodyList {
	s.RankIdList = v
	return s
}

func (s *QueryPositionsResponseBodyList) SetStatus(v int32) *QueryPositionsResponseBodyList {
	s.Status = &v
	return s
}

type QueryPositionsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPositionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPositionsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPositionsResponse) GoString() string {
	return s.String()
}

func (s *QueryPositionsResponse) SetHeaders(v map[string]*string) *QueryPositionsResponse {
	s.Headers = v
	return s
}

func (s *QueryPositionsResponse) SetStatusCode(v int32) *QueryPositionsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPositionsResponse) SetBody(v *QueryPositionsResponseBody) *QueryPositionsResponse {
	s.Body = v
	return s
}

type RevokeSignRecordsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RevokeSignRecordsHeaders) String() string {
	return tea.Prettify(s)
}

func (s RevokeSignRecordsHeaders) GoString() string {
	return s.String()
}

func (s *RevokeSignRecordsHeaders) SetCommonHeaders(v map[string]*string) *RevokeSignRecordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RevokeSignRecordsHeaders) SetXAcsDingtalkAccessToken(v string) *RevokeSignRecordsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RevokeSignRecordsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	RevokeUserId *string `json:"revokeUserId,omitempty" xml:"revokeUserId,omitempty"`
	// This parameter is required.
	SignRecordIds []*string `json:"signRecordIds,omitempty" xml:"signRecordIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 撤销签署
	StatusRemark *string `json:"statusRemark,omitempty" xml:"statusRemark,omitempty"`
}

func (s RevokeSignRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeSignRecordsRequest) GoString() string {
	return s.String()
}

func (s *RevokeSignRecordsRequest) SetRevokeUserId(v string) *RevokeSignRecordsRequest {
	s.RevokeUserId = &v
	return s
}

func (s *RevokeSignRecordsRequest) SetSignRecordIds(v []*string) *RevokeSignRecordsRequest {
	s.SignRecordIds = v
	return s
}

func (s *RevokeSignRecordsRequest) SetStatusRemark(v string) *RevokeSignRecordsRequest {
	s.StatusRemark = &v
	return s
}

type RevokeSignRecordsResponseBody struct {
	Result *RevokeSignRecordsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// This parameter is required.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RevokeSignRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeSignRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeSignRecordsResponseBody) SetResult(v *RevokeSignRecordsResponseBodyResult) *RevokeSignRecordsResponseBody {
	s.Result = v
	return s
}

func (s *RevokeSignRecordsResponseBody) SetSuccess(v bool) *RevokeSignRecordsResponseBody {
	s.Success = &v
	return s
}

type RevokeSignRecordsResponseBodyResult struct {
	FailItems    []*RevokeSignRecordsResponseBodyResultFailItems    `json:"failItems,omitempty" xml:"failItems,omitempty" type:"Repeated"`
	SuccessItems []*RevokeSignRecordsResponseBodyResultSuccessItems `json:"successItems,omitempty" xml:"successItems,omitempty" type:"Repeated"`
}

func (s RevokeSignRecordsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s RevokeSignRecordsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RevokeSignRecordsResponseBodyResult) SetFailItems(v []*RevokeSignRecordsResponseBodyResultFailItems) *RevokeSignRecordsResponseBodyResult {
	s.FailItems = v
	return s
}

func (s *RevokeSignRecordsResponseBodyResult) SetSuccessItems(v []*RevokeSignRecordsResponseBodyResultSuccessItems) *RevokeSignRecordsResponseBodyResult {
	s.SuccessItems = v
	return s
}

type RevokeSignRecordsResponseBodyResultFailItems struct {
	// example:
	//
	// 6fe06f57ab5a45078f3219be8fd829c6
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// example:
	//
	// 电子签状态变更不合法
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RevokeSignRecordsResponseBodyResultFailItems) String() string {
	return tea.Prettify(s)
}

func (s RevokeSignRecordsResponseBodyResultFailItems) GoString() string {
	return s.String()
}

func (s *RevokeSignRecordsResponseBodyResultFailItems) SetItemId(v string) *RevokeSignRecordsResponseBodyResultFailItems {
	s.ItemId = &v
	return s
}

func (s *RevokeSignRecordsResponseBodyResultFailItems) SetType(v string) *RevokeSignRecordsResponseBodyResultFailItems {
	s.Type = &v
	return s
}

type RevokeSignRecordsResponseBodyResultSuccessItems struct {
	// example:
	//
	// 6fe06f57ab5a45078f3219be8fd829c6
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
}

func (s RevokeSignRecordsResponseBodyResultSuccessItems) String() string {
	return tea.Prettify(s)
}

func (s RevokeSignRecordsResponseBodyResultSuccessItems) GoString() string {
	return s.String()
}

func (s *RevokeSignRecordsResponseBodyResultSuccessItems) SetItemId(v string) *RevokeSignRecordsResponseBodyResultSuccessItems {
	s.ItemId = &v
	return s
}

type RevokeSignRecordsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeSignRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeSignRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeSignRecordsResponse) GoString() string {
	return s.String()
}

func (s *RevokeSignRecordsResponse) SetHeaders(v map[string]*string) *RevokeSignRecordsResponse {
	s.Headers = v
	return s
}

func (s *RevokeSignRecordsResponse) SetStatusCode(v int32) *RevokeSignRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeSignRecordsResponse) SetBody(v *RevokeSignRecordsResponseBody) *RevokeSignRecordsResponse {
	s.Body = v
	return s
}

type RevokeTerminationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RevokeTerminationHeaders) String() string {
	return tea.Prettify(s)
}

func (s RevokeTerminationHeaders) GoString() string {
	return s.String()
}

func (s *RevokeTerminationHeaders) SetCommonHeaders(v map[string]*string) *RevokeTerminationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RevokeTerminationHeaders) SetXAcsDingtalkAccessToken(v string) *RevokeTerminationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RevokeTerminationRequest struct {
	// example:
	//
	// 2163515669935611
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s RevokeTerminationRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeTerminationRequest) GoString() string {
	return s.String()
}

func (s *RevokeTerminationRequest) SetUserId(v string) *RevokeTerminationRequest {
	s.UserId = &v
	return s
}

type RevokeTerminationResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RevokeTerminationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeTerminationResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeTerminationResponseBody) SetRequestId(v string) *RevokeTerminationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeTerminationResponseBody) SetResult(v bool) *RevokeTerminationResponseBody {
	s.Result = &v
	return s
}

func (s *RevokeTerminationResponseBody) SetSuccess(v bool) *RevokeTerminationResponseBody {
	s.Success = &v
	return s
}

type RevokeTerminationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeTerminationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeTerminationResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeTerminationResponse) GoString() string {
	return s.String()
}

func (s *RevokeTerminationResponse) SetHeaders(v map[string]*string) *RevokeTerminationResponse {
	s.Headers = v
	return s
}

func (s *RevokeTerminationResponse) SetStatusCode(v int32) *RevokeTerminationResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeTerminationResponse) SetBody(v *RevokeTerminationResponseBody) *RevokeTerminationResponse {
	s.Body = v
	return s
}

type RosterMetaAvailableFieldListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RosterMetaAvailableFieldListHeaders) String() string {
	return tea.Prettify(s)
}

func (s RosterMetaAvailableFieldListHeaders) GoString() string {
	return s.String()
}

func (s *RosterMetaAvailableFieldListHeaders) SetCommonHeaders(v map[string]*string) *RosterMetaAvailableFieldListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RosterMetaAvailableFieldListHeaders) SetXAcsDingtalkAccessToken(v string) *RosterMetaAvailableFieldListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RosterMetaAvailableFieldListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1231
	AppAgentId *int64 `json:"appAgentId,omitempty" xml:"appAgentId,omitempty"`
}

func (s RosterMetaAvailableFieldListRequest) String() string {
	return tea.Prettify(s)
}

func (s RosterMetaAvailableFieldListRequest) GoString() string {
	return s.String()
}

func (s *RosterMetaAvailableFieldListRequest) SetAppAgentId(v int64) *RosterMetaAvailableFieldListRequest {
	s.AppAgentId = &v
	return s
}

type RosterMetaAvailableFieldListResponseBody struct {
	Result []*RosterMetaAvailableFieldListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s RosterMetaAvailableFieldListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RosterMetaAvailableFieldListResponseBody) GoString() string {
	return s.String()
}

func (s *RosterMetaAvailableFieldListResponseBody) SetResult(v []*RosterMetaAvailableFieldListResponseBodyResult) *RosterMetaAvailableFieldListResponseBody {
	s.Result = v
	return s
}

type RosterMetaAvailableFieldListResponseBodyResult struct {
	// example:
	//
	// sys01-employeeType
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	// example:
	//
	// 员工类型
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// example:
	//
	// DDSelectField
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	// example:
	//
	// [{"value":"1","label":"男"},{"value":"2","label":"女"}]
	OptionText *string `json:"optionText,omitempty" xml:"optionText,omitempty"`
}

func (s RosterMetaAvailableFieldListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s RosterMetaAvailableFieldListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RosterMetaAvailableFieldListResponseBodyResult) SetFieldCode(v string) *RosterMetaAvailableFieldListResponseBodyResult {
	s.FieldCode = &v
	return s
}

func (s *RosterMetaAvailableFieldListResponseBodyResult) SetFieldName(v string) *RosterMetaAvailableFieldListResponseBodyResult {
	s.FieldName = &v
	return s
}

func (s *RosterMetaAvailableFieldListResponseBodyResult) SetFieldType(v string) *RosterMetaAvailableFieldListResponseBodyResult {
	s.FieldType = &v
	return s
}

func (s *RosterMetaAvailableFieldListResponseBodyResult) SetOptionText(v string) *RosterMetaAvailableFieldListResponseBodyResult {
	s.OptionText = &v
	return s
}

type RosterMetaAvailableFieldListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RosterMetaAvailableFieldListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RosterMetaAvailableFieldListResponse) String() string {
	return tea.Prettify(s)
}

func (s RosterMetaAvailableFieldListResponse) GoString() string {
	return s.String()
}

func (s *RosterMetaAvailableFieldListResponse) SetHeaders(v map[string]*string) *RosterMetaAvailableFieldListResponse {
	s.Headers = v
	return s
}

func (s *RosterMetaAvailableFieldListResponse) SetStatusCode(v int32) *RosterMetaAvailableFieldListResponse {
	s.StatusCode = &v
	return s
}

func (s *RosterMetaAvailableFieldListResponse) SetBody(v *RosterMetaAvailableFieldListResponseBody) *RosterMetaAvailableFieldListResponse {
	s.Body = v
	return s
}

type RosterMetaFieldOptionsUpdateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RosterMetaFieldOptionsUpdateHeaders) String() string {
	return tea.Prettify(s)
}

func (s RosterMetaFieldOptionsUpdateHeaders) GoString() string {
	return s.String()
}

func (s *RosterMetaFieldOptionsUpdateHeaders) SetCommonHeaders(v map[string]*string) *RosterMetaFieldOptionsUpdateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RosterMetaFieldOptionsUpdateHeaders) SetXAcsDingtalkAccessToken(v string) *RosterMetaFieldOptionsUpdateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RosterMetaFieldOptionsUpdateRequest struct {
	AppAgentId *int64 `json:"appAgentId,omitempty" xml:"appAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sys05-contractType
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sys05
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// This parameter is required.
	Labels []*string `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// OPTIONS_ADD
	ModifyType *string `json:"modifyType,omitempty" xml:"modifyType,omitempty"`
}

func (s RosterMetaFieldOptionsUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s RosterMetaFieldOptionsUpdateRequest) GoString() string {
	return s.String()
}

func (s *RosterMetaFieldOptionsUpdateRequest) SetAppAgentId(v int64) *RosterMetaFieldOptionsUpdateRequest {
	s.AppAgentId = &v
	return s
}

func (s *RosterMetaFieldOptionsUpdateRequest) SetFieldCode(v string) *RosterMetaFieldOptionsUpdateRequest {
	s.FieldCode = &v
	return s
}

func (s *RosterMetaFieldOptionsUpdateRequest) SetGroupId(v string) *RosterMetaFieldOptionsUpdateRequest {
	s.GroupId = &v
	return s
}

func (s *RosterMetaFieldOptionsUpdateRequest) SetLabels(v []*string) *RosterMetaFieldOptionsUpdateRequest {
	s.Labels = v
	return s
}

func (s *RosterMetaFieldOptionsUpdateRequest) SetModifyType(v string) *RosterMetaFieldOptionsUpdateRequest {
	s.ModifyType = &v
	return s
}

type RosterMetaFieldOptionsUpdateResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RosterMetaFieldOptionsUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RosterMetaFieldOptionsUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *RosterMetaFieldOptionsUpdateResponseBody) SetResult(v bool) *RosterMetaFieldOptionsUpdateResponseBody {
	s.Result = &v
	return s
}

type RosterMetaFieldOptionsUpdateResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RosterMetaFieldOptionsUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RosterMetaFieldOptionsUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s RosterMetaFieldOptionsUpdateResponse) GoString() string {
	return s.String()
}

func (s *RosterMetaFieldOptionsUpdateResponse) SetHeaders(v map[string]*string) *RosterMetaFieldOptionsUpdateResponse {
	s.Headers = v
	return s
}

func (s *RosterMetaFieldOptionsUpdateResponse) SetStatusCode(v int32) *RosterMetaFieldOptionsUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *RosterMetaFieldOptionsUpdateResponse) SetBody(v *RosterMetaFieldOptionsUpdateResponseBody) *RosterMetaFieldOptionsUpdateResponse {
	s.Body = v
	return s
}

type SendIsvCardMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendIsvCardMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendIsvCardMessageHeaders) GoString() string {
	return s.String()
}

func (s *SendIsvCardMessageHeaders) SetCommonHeaders(v map[string]*string) *SendIsvCardMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendIsvCardMessageHeaders) SetXAcsDingtalkAccessToken(v string) *SendIsvCardMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendIsvCardMessageRequest struct {
	// This parameter is required.
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// This parameter is required.
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// This parameter is required.
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// This parameter is required.
	ReceiverUserIds []*string `json:"receiverUserIds,omitempty" xml:"receiverUserIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 16690147049882572
	SceneType *string `json:"sceneType,omitempty" xml:"sceneType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 同意转正
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 16690147049882572
	SenderUserId *string            `json:"senderUserId,omitempty" xml:"senderUserId,omitempty"`
	ValueMap     map[string]*string `json:"valueMap,omitempty" xml:"valueMap,omitempty"`
}

func (s SendIsvCardMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendIsvCardMessageRequest) GoString() string {
	return s.String()
}

func (s *SendIsvCardMessageRequest) SetAgentId(v int64) *SendIsvCardMessageRequest {
	s.AgentId = &v
	return s
}

func (s *SendIsvCardMessageRequest) SetBizId(v string) *SendIsvCardMessageRequest {
	s.BizId = &v
	return s
}

func (s *SendIsvCardMessageRequest) SetMessageType(v string) *SendIsvCardMessageRequest {
	s.MessageType = &v
	return s
}

func (s *SendIsvCardMessageRequest) SetReceiverUserIds(v []*string) *SendIsvCardMessageRequest {
	s.ReceiverUserIds = v
	return s
}

func (s *SendIsvCardMessageRequest) SetSceneType(v string) *SendIsvCardMessageRequest {
	s.SceneType = &v
	return s
}

func (s *SendIsvCardMessageRequest) SetScope(v string) *SendIsvCardMessageRequest {
	s.Scope = &v
	return s
}

func (s *SendIsvCardMessageRequest) SetSenderUserId(v string) *SendIsvCardMessageRequest {
	s.SenderUserId = &v
	return s
}

func (s *SendIsvCardMessageRequest) SetValueMap(v map[string]*string) *SendIsvCardMessageRequest {
	s.ValueMap = v
	return s
}

type SendIsvCardMessageResponseBody struct {
	ErrorCode                    *string                                                     `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg                     *string                                                     `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	HrmInteractiveCardSendResult *SendIsvCardMessageResponseBodyHrmInteractiveCardSendResult `json:"hrmInteractiveCardSendResult,omitempty" xml:"hrmInteractiveCardSendResult,omitempty" type:"Struct"`
	RequestId                    *string                                                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success                      *bool                                                       `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SendIsvCardMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendIsvCardMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendIsvCardMessageResponseBody) SetErrorCode(v string) *SendIsvCardMessageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SendIsvCardMessageResponseBody) SetErrorMsg(v string) *SendIsvCardMessageResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SendIsvCardMessageResponseBody) SetHrmInteractiveCardSendResult(v *SendIsvCardMessageResponseBodyHrmInteractiveCardSendResult) *SendIsvCardMessageResponseBody {
	s.HrmInteractiveCardSendResult = v
	return s
}

func (s *SendIsvCardMessageResponseBody) SetRequestId(v string) *SendIsvCardMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendIsvCardMessageResponseBody) SetSuccess(v bool) *SendIsvCardMessageResponseBody {
	s.Success = &v
	return s
}

type SendIsvCardMessageResponseBodyHrmInteractiveCardSendResult struct {
	BizId     *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
}

func (s SendIsvCardMessageResponseBodyHrmInteractiveCardSendResult) String() string {
	return tea.Prettify(s)
}

func (s SendIsvCardMessageResponseBodyHrmInteractiveCardSendResult) GoString() string {
	return s.String()
}

func (s *SendIsvCardMessageResponseBodyHrmInteractiveCardSendResult) SetBizId(v string) *SendIsvCardMessageResponseBodyHrmInteractiveCardSendResult {
	s.BizId = &v
	return s
}

func (s *SendIsvCardMessageResponseBodyHrmInteractiveCardSendResult) SetErrorCode(v string) *SendIsvCardMessageResponseBodyHrmInteractiveCardSendResult {
	s.ErrorCode = &v
	return s
}

func (s *SendIsvCardMessageResponseBodyHrmInteractiveCardSendResult) SetErrorMsg(v string) *SendIsvCardMessageResponseBodyHrmInteractiveCardSendResult {
	s.ErrorMsg = &v
	return s
}

type SendIsvCardMessageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendIsvCardMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendIsvCardMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendIsvCardMessageResponse) GoString() string {
	return s.String()
}

func (s *SendIsvCardMessageResponse) SetHeaders(v map[string]*string) *SendIsvCardMessageResponse {
	s.Headers = v
	return s
}

func (s *SendIsvCardMessageResponse) SetStatusCode(v int32) *SendIsvCardMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendIsvCardMessageResponse) SetBody(v *SendIsvCardMessageResponseBody) *SendIsvCardMessageResponse {
	s.Body = v
	return s
}

type SendRealAuthInviteMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendRealAuthInviteMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendRealAuthInviteMessageHeaders) GoString() string {
	return s.String()
}

func (s *SendRealAuthInviteMessageHeaders) SetCommonHeaders(v map[string]*string) *SendRealAuthInviteMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendRealAuthInviteMessageHeaders) SetXAcsDingtalkAccessToken(v string) *SendRealAuthInviteMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendRealAuthInviteMessageRequest struct {
	// This parameter is required.
	InviterId *string `json:"inviterId,omitempty" xml:"inviterId,omitempty"`
	// This parameter is required.
	OnWorkingEmpSearchVO *SendRealAuthInviteMessageRequestOnWorkingEmpSearchVO `json:"onWorkingEmpSearchVO,omitempty" xml:"onWorkingEmpSearchVO,omitempty" type:"Struct"`
}

func (s SendRealAuthInviteMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendRealAuthInviteMessageRequest) GoString() string {
	return s.String()
}

func (s *SendRealAuthInviteMessageRequest) SetInviterId(v string) *SendRealAuthInviteMessageRequest {
	s.InviterId = &v
	return s
}

func (s *SendRealAuthInviteMessageRequest) SetOnWorkingEmpSearchVO(v *SendRealAuthInviteMessageRequestOnWorkingEmpSearchVO) *SendRealAuthInviteMessageRequest {
	s.OnWorkingEmpSearchVO = v
	return s
}

type SendRealAuthInviteMessageRequestOnWorkingEmpSearchVO struct {
	// This parameter is required.
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s SendRealAuthInviteMessageRequestOnWorkingEmpSearchVO) String() string {
	return tea.Prettify(s)
}

func (s SendRealAuthInviteMessageRequestOnWorkingEmpSearchVO) GoString() string {
	return s.String()
}

func (s *SendRealAuthInviteMessageRequestOnWorkingEmpSearchVO) SetUserIds(v []*string) *SendRealAuthInviteMessageRequestOnWorkingEmpSearchVO {
	s.UserIds = v
	return s
}

type SendRealAuthInviteMessageResponseBody struct {
	DingOpenErrcode *int32                                       `json:"dingOpenErrcode,omitempty" xml:"dingOpenErrcode,omitempty"`
	ErrorMsg        *string                                      `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result          *SendRealAuthInviteMessageResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success         *bool                                        `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SendRealAuthInviteMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendRealAuthInviteMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendRealAuthInviteMessageResponseBody) SetDingOpenErrcode(v int32) *SendRealAuthInviteMessageResponseBody {
	s.DingOpenErrcode = &v
	return s
}

func (s *SendRealAuthInviteMessageResponseBody) SetErrorMsg(v string) *SendRealAuthInviteMessageResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SendRealAuthInviteMessageResponseBody) SetResult(v *SendRealAuthInviteMessageResponseBodyResult) *SendRealAuthInviteMessageResponseBody {
	s.Result = v
	return s
}

func (s *SendRealAuthInviteMessageResponseBody) SetSuccess(v bool) *SendRealAuthInviteMessageResponseBody {
	s.Success = &v
	return s
}

type SendRealAuthInviteMessageResponseBodyResult struct {
	HadInvitedNum  *int32 `json:"hadInvitedNum,omitempty" xml:"hadInvitedNum,omitempty"`
	HadRealAuthNum *int32 `json:"hadRealAuthNum,omitempty" xml:"hadRealAuthNum,omitempty"`
	SuccessNum     *int32 `json:"successNum,omitempty" xml:"successNum,omitempty"`
}

func (s SendRealAuthInviteMessageResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SendRealAuthInviteMessageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SendRealAuthInviteMessageResponseBodyResult) SetHadInvitedNum(v int32) *SendRealAuthInviteMessageResponseBodyResult {
	s.HadInvitedNum = &v
	return s
}

func (s *SendRealAuthInviteMessageResponseBodyResult) SetHadRealAuthNum(v int32) *SendRealAuthInviteMessageResponseBodyResult {
	s.HadRealAuthNum = &v
	return s
}

func (s *SendRealAuthInviteMessageResponseBodyResult) SetSuccessNum(v int32) *SendRealAuthInviteMessageResponseBodyResult {
	s.SuccessNum = &v
	return s
}

type SendRealAuthInviteMessageResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendRealAuthInviteMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendRealAuthInviteMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendRealAuthInviteMessageResponse) GoString() string {
	return s.String()
}

func (s *SendRealAuthInviteMessageResponse) SetHeaders(v map[string]*string) *SendRealAuthInviteMessageResponse {
	s.Headers = v
	return s
}

func (s *SendRealAuthInviteMessageResponse) SetStatusCode(v int32) *SendRealAuthInviteMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendRealAuthInviteMessageResponse) SetBody(v *SendRealAuthInviteMessageResponseBody) *SendRealAuthInviteMessageResponse {
	s.Body = v
	return s
}

type SolutionTaskInitHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SolutionTaskInitHeaders) String() string {
	return tea.Prettify(s)
}

func (s SolutionTaskInitHeaders) GoString() string {
	return s.String()
}

func (s *SolutionTaskInitHeaders) SetCommonHeaders(v map[string]*string) *SolutionTaskInitHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SolutionTaskInitHeaders) SetXAcsDingtalkAccessToken(v string) *SolutionTaskInitHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SolutionTaskInitRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// training
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// 时间戳
	ClaimTime *int64 `json:"claimTime,omitempty" xml:"claimTime,omitempty"`
	// example:
	//
	// 这是一个新人培训任务
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 时间戳
	FinishTime *int64 `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fdagshfjhajl
	OuterId *string `json:"outerId,omitempty" xml:"outerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 新人学习任务
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// onboarding
	SolutionType *string `json:"solutionType,omitempty" xml:"solutionType,omitempty"`
}

func (s SolutionTaskInitRequest) String() string {
	return tea.Prettify(s)
}

func (s SolutionTaskInitRequest) GoString() string {
	return s.String()
}

func (s *SolutionTaskInitRequest) SetCategory(v string) *SolutionTaskInitRequest {
	s.Category = &v
	return s
}

func (s *SolutionTaskInitRequest) SetClaimTime(v int64) *SolutionTaskInitRequest {
	s.ClaimTime = &v
	return s
}

func (s *SolutionTaskInitRequest) SetDescription(v string) *SolutionTaskInitRequest {
	s.Description = &v
	return s
}

func (s *SolutionTaskInitRequest) SetFinishTime(v int64) *SolutionTaskInitRequest {
	s.FinishTime = &v
	return s
}

func (s *SolutionTaskInitRequest) SetOuterId(v string) *SolutionTaskInitRequest {
	s.OuterId = &v
	return s
}

func (s *SolutionTaskInitRequest) SetStatus(v string) *SolutionTaskInitRequest {
	s.Status = &v
	return s
}

func (s *SolutionTaskInitRequest) SetTitle(v string) *SolutionTaskInitRequest {
	s.Title = &v
	return s
}

func (s *SolutionTaskInitRequest) SetUserId(v string) *SolutionTaskInitRequest {
	s.UserId = &v
	return s
}

func (s *SolutionTaskInitRequest) SetSolutionType(v string) *SolutionTaskInitRequest {
	s.SolutionType = &v
	return s
}

type SolutionTaskInitResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SolutionTaskInitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SolutionTaskInitResponseBody) GoString() string {
	return s.String()
}

func (s *SolutionTaskInitResponseBody) SetResult(v bool) *SolutionTaskInitResponseBody {
	s.Result = &v
	return s
}

type SolutionTaskInitResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SolutionTaskInitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SolutionTaskInitResponse) String() string {
	return tea.Prettify(s)
}

func (s SolutionTaskInitResponse) GoString() string {
	return s.String()
}

func (s *SolutionTaskInitResponse) SetHeaders(v map[string]*string) *SolutionTaskInitResponse {
	s.Headers = v
	return s
}

func (s *SolutionTaskInitResponse) SetStatusCode(v int32) *SolutionTaskInitResponse {
	s.StatusCode = &v
	return s
}

func (s *SolutionTaskInitResponse) SetBody(v *SolutionTaskInitResponseBody) *SolutionTaskInitResponse {
	s.Body = v
	return s
}

type SolutionTaskSaveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SolutionTaskSaveHeaders) String() string {
	return tea.Prettify(s)
}

func (s SolutionTaskSaveHeaders) GoString() string {
	return s.String()
}

func (s *SolutionTaskSaveHeaders) SetCommonHeaders(v map[string]*string) *SolutionTaskSaveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SolutionTaskSaveHeaders) SetXAcsDingtalkAccessToken(v string) *SolutionTaskSaveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SolutionTaskSaveRequest struct {
	// example:
	//
	// 时间戳
	ClaimTime *int64 `json:"claimTime,omitempty" xml:"claimTime,omitempty"`
	// example:
	//
	// 这是一个新人培训任务
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 时间戳
	FinishTime *int64 `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fdagshfjhajl
	OuterId *string `json:"outerId,omitempty" xml:"outerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qweqweqwe
	SolutionInstanceId *string `json:"solutionInstanceId,omitempty" xml:"solutionInstanceId,omitempty"`
	StartTime          *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PERFORMANCE_TASK、TRAIN_TASK
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// example:
	//
	// sdfasd2323sdaf
	TemplateOuterId *string `json:"templateOuterId,omitempty" xml:"templateOuterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 新人学习任务
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// onboarding
	SolutionType *string `json:"solutionType,omitempty" xml:"solutionType,omitempty"`
}

func (s SolutionTaskSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s SolutionTaskSaveRequest) GoString() string {
	return s.String()
}

func (s *SolutionTaskSaveRequest) SetClaimTime(v int64) *SolutionTaskSaveRequest {
	s.ClaimTime = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetDescription(v string) *SolutionTaskSaveRequest {
	s.Description = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetFinishTime(v int64) *SolutionTaskSaveRequest {
	s.FinishTime = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetOuterId(v string) *SolutionTaskSaveRequest {
	s.OuterId = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetSolutionInstanceId(v string) *SolutionTaskSaveRequest {
	s.SolutionInstanceId = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetStartTime(v int64) *SolutionTaskSaveRequest {
	s.StartTime = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetStatus(v string) *SolutionTaskSaveRequest {
	s.Status = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetTaskType(v string) *SolutionTaskSaveRequest {
	s.TaskType = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetTemplateOuterId(v string) *SolutionTaskSaveRequest {
	s.TemplateOuterId = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetTitle(v string) *SolutionTaskSaveRequest {
	s.Title = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetUserId(v string) *SolutionTaskSaveRequest {
	s.UserId = &v
	return s
}

func (s *SolutionTaskSaveRequest) SetSolutionType(v string) *SolutionTaskSaveRequest {
	s.SolutionType = &v
	return s
}

type SolutionTaskSaveResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SolutionTaskSaveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SolutionTaskSaveResponseBody) GoString() string {
	return s.String()
}

func (s *SolutionTaskSaveResponseBody) SetResult(v bool) *SolutionTaskSaveResponseBody {
	s.Result = &v
	return s
}

type SolutionTaskSaveResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SolutionTaskSaveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SolutionTaskSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s SolutionTaskSaveResponse) GoString() string {
	return s.String()
}

func (s *SolutionTaskSaveResponse) SetHeaders(v map[string]*string) *SolutionTaskSaveResponse {
	s.Headers = v
	return s
}

func (s *SolutionTaskSaveResponse) SetStatusCode(v int32) *SolutionTaskSaveResponse {
	s.StatusCode = &v
	return s
}

func (s *SolutionTaskSaveResponse) SetBody(v *SolutionTaskSaveResponseBody) *SolutionTaskSaveResponse {
	s.Body = v
	return s
}

type SyncSolutionStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SyncSolutionStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncSolutionStatusHeaders) GoString() string {
	return s.String()
}

func (s *SyncSolutionStatusHeaders) SetCommonHeaders(v map[string]*string) *SyncSolutionStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncSolutionStatusHeaders) SetXAcsDingtalkAccessToken(v string) *SyncSolutionStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SyncSolutionStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// start
	SolutionStatus *string `json:"solutionStatus,omitempty" xml:"solutionStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// onboarding_v2
	SolutionType *string `json:"solutionType,omitempty" xml:"solutionType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12
	TenantId *int64 `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// This parameter is required.
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s SyncSolutionStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncSolutionStatusRequest) GoString() string {
	return s.String()
}

func (s *SyncSolutionStatusRequest) SetBizId(v string) *SyncSolutionStatusRequest {
	s.BizId = &v
	return s
}

func (s *SyncSolutionStatusRequest) SetSolutionStatus(v string) *SyncSolutionStatusRequest {
	s.SolutionStatus = &v
	return s
}

func (s *SyncSolutionStatusRequest) SetSolutionType(v string) *SyncSolutionStatusRequest {
	s.SolutionType = &v
	return s
}

func (s *SyncSolutionStatusRequest) SetTenantId(v int64) *SyncSolutionStatusRequest {
	s.TenantId = &v
	return s
}

func (s *SyncSolutionStatusRequest) SetUserIds(v []*string) *SyncSolutionStatusRequest {
	s.UserIds = v
	return s
}

type SyncSolutionStatusResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SyncSolutionStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncSolutionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SyncSolutionStatusResponseBody) SetResult(v bool) *SyncSolutionStatusResponseBody {
	s.Result = &v
	return s
}

type SyncSolutionStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncSolutionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncSolutionStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncSolutionStatusResponse) GoString() string {
	return s.String()
}

func (s *SyncSolutionStatusResponse) SetHeaders(v map[string]*string) *SyncSolutionStatusResponse {
	s.Headers = v
	return s
}

func (s *SyncSolutionStatusResponse) SetStatusCode(v int32) *SyncSolutionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncSolutionStatusResponse) SetBody(v *SyncSolutionStatusResponseBody) *SyncSolutionStatusResponse {
	s.Body = v
	return s
}

type SyncTaskTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SyncTaskTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncTaskTemplateHeaders) GoString() string {
	return s.String()
}

func (s *SyncTaskTemplateHeaders) SetCommonHeaders(v map[string]*string) *SyncTaskTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncTaskTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *SyncTaskTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SyncTaskTemplateRequest struct {
	// if can be null:
	// false
	Delete *bool `json:"delete,omitempty" xml:"delete,omitempty"`
	// example:
	//
	// 培训、薪酬任务模版
	Des *string `json:"des,omitempty" xml:"des,omitempty"`
	// example:
	//
	// {\"key\":value}
	Ext *string `json:"ext,omitempty" xml:"ext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 培训模版
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 23234
	OptUserId *string `json:"optUserId,omitempty" xml:"optUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 232332
	OuterId     *string                             `json:"outerId,omitempty" xml:"outerId,omitempty"`
	TaskScopeVO *SyncTaskTemplateRequestTaskScopeVO `json:"taskScopeVO,omitempty" xml:"taskScopeVO,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// PERFORMANCE_TASK、TRAIN_TASK
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// onboarding
	SolutionType *string `json:"solutionType,omitempty" xml:"solutionType,omitempty"`
}

func (s SyncTaskTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncTaskTemplateRequest) GoString() string {
	return s.String()
}

func (s *SyncTaskTemplateRequest) SetDelete(v bool) *SyncTaskTemplateRequest {
	s.Delete = &v
	return s
}

func (s *SyncTaskTemplateRequest) SetDes(v string) *SyncTaskTemplateRequest {
	s.Des = &v
	return s
}

func (s *SyncTaskTemplateRequest) SetExt(v string) *SyncTaskTemplateRequest {
	s.Ext = &v
	return s
}

func (s *SyncTaskTemplateRequest) SetName(v string) *SyncTaskTemplateRequest {
	s.Name = &v
	return s
}

func (s *SyncTaskTemplateRequest) SetOptUserId(v string) *SyncTaskTemplateRequest {
	s.OptUserId = &v
	return s
}

func (s *SyncTaskTemplateRequest) SetOuterId(v string) *SyncTaskTemplateRequest {
	s.OuterId = &v
	return s
}

func (s *SyncTaskTemplateRequest) SetTaskScopeVO(v *SyncTaskTemplateRequestTaskScopeVO) *SyncTaskTemplateRequest {
	s.TaskScopeVO = v
	return s
}

func (s *SyncTaskTemplateRequest) SetTaskType(v string) *SyncTaskTemplateRequest {
	s.TaskType = &v
	return s
}

func (s *SyncTaskTemplateRequest) SetSolutionType(v string) *SyncTaskTemplateRequest {
	s.SolutionType = &v
	return s
}

type SyncTaskTemplateRequestTaskScopeVO struct {
	DeptIds     []*int64  `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	PositionIds []*string `json:"positionIds,omitempty" xml:"positionIds,omitempty" type:"Repeated"`
	RoleIds     []*string `json:"roleIds,omitempty" xml:"roleIds,omitempty" type:"Repeated"`
	UserIds     []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s SyncTaskTemplateRequestTaskScopeVO) String() string {
	return tea.Prettify(s)
}

func (s SyncTaskTemplateRequestTaskScopeVO) GoString() string {
	return s.String()
}

func (s *SyncTaskTemplateRequestTaskScopeVO) SetDeptIds(v []*int64) *SyncTaskTemplateRequestTaskScopeVO {
	s.DeptIds = v
	return s
}

func (s *SyncTaskTemplateRequestTaskScopeVO) SetPositionIds(v []*string) *SyncTaskTemplateRequestTaskScopeVO {
	s.PositionIds = v
	return s
}

func (s *SyncTaskTemplateRequestTaskScopeVO) SetRoleIds(v []*string) *SyncTaskTemplateRequestTaskScopeVO {
	s.RoleIds = v
	return s
}

func (s *SyncTaskTemplateRequestTaskScopeVO) SetUserIds(v []*string) *SyncTaskTemplateRequestTaskScopeVO {
	s.UserIds = v
	return s
}

type SyncTaskTemplateResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SyncTaskTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncTaskTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *SyncTaskTemplateResponseBody) SetResult(v bool) *SyncTaskTemplateResponseBody {
	s.Result = &v
	return s
}

type SyncTaskTemplateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncTaskTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncTaskTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncTaskTemplateResponse) GoString() string {
	return s.String()
}

func (s *SyncTaskTemplateResponse) SetHeaders(v map[string]*string) *SyncTaskTemplateResponse {
	s.Headers = v
	return s
}

func (s *SyncTaskTemplateResponse) SetStatusCode(v int32) *SyncTaskTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncTaskTemplateResponse) SetBody(v *SyncTaskTemplateResponseBody) *SyncTaskTemplateResponse {
	s.Body = v
	return s
}

type UpdateCustomRosterFieldHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateCustomRosterFieldHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomRosterFieldHeaders) GoString() string {
	return s.String()
}

func (s *UpdateCustomRosterFieldHeaders) SetCommonHeaders(v map[string]*string) *UpdateCustomRosterFieldHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateCustomRosterFieldHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateCustomRosterFieldHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateCustomRosterFieldRequest struct {
	ContactClientFlag *bool  `json:"contactClientFlag,omitempty" xml:"contactClientFlag,omitempty"`
	ContactFlag       *bool  `json:"contactFlag,omitempty" xml:"contactFlag,omitempty"`
	ContactSource     *int32 `json:"contactSource,omitempty" xml:"contactSource,omitempty"`
	ContactSystemFlag *bool  `json:"contactSystemFlag,omitempty" xml:"contactSystemFlag,omitempty"`
	Deleted           *bool  `json:"deleted,omitempty" xml:"deleted,omitempty"`
	Derived           *bool  `json:"derived,omitempty" xml:"derived,omitempty"`
	Disabled          *int32 `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// This parameter is required.
	EditFromEmployeeFlag *bool `json:"editFromEmployeeFlag,omitempty" xml:"editFromEmployeeFlag,omitempty"`
	EditableByHr         *bool `json:"editableByHr,omitempty" xml:"editableByHr,omitempty"`
	// This parameter is required.
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	// This parameter is required.
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	FieldTip  *string `json:"fieldTip,omitempty" xml:"fieldTip,omitempty"`
	// This parameter is required.
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	// This parameter is required.
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// This parameter is required.
	HiddenFromEmployeeFlag *bool   `json:"hiddenFromEmployeeFlag,omitempty" xml:"hiddenFromEmployeeFlag,omitempty"`
	Hint                   *string `json:"hint,omitempty" xml:"hint,omitempty"`
	HistoryField           *bool   `json:"historyField,omitempty" xml:"historyField,omitempty"`
	Index                  *int32  `json:"index,omitempty" xml:"index,omitempty"`
	ModifyOptions          *bool   `json:"modifyOptions,omitempty" xml:"modifyOptions,omitempty"`
	NoWatermark            *bool   `json:"noWatermark,omitempty" xml:"noWatermark,omitempty"`
	NumberDecimalPlace     *int32  `json:"numberDecimalPlace,omitempty" xml:"numberDecimalPlace,omitempty"`
	NumberFormatType       *string `json:"numberFormatType,omitempty" xml:"numberFormatType,omitempty"`
	NumberValueType        *string `json:"numberValueType,omitempty" xml:"numberValueType,omitempty"`
	OptionText             *string `json:"optionText,omitempty" xml:"optionText,omitempty"`
	// This parameter is required.
	Required          *bool   `json:"required,omitempty" xml:"required,omitempty"`
	SourceFieldCode   *string `json:"sourceFieldCode,omitempty" xml:"sourceFieldCode,omitempty"`
	SystemFlag        *bool   `json:"systemFlag,omitempty" xml:"systemFlag,omitempty"`
	TextToSelectField *bool   `json:"textToSelectField,omitempty" xml:"textToSelectField,omitempty"`
	// This parameter is required.
	UpdateUserId *string `json:"updateUserId,omitempty" xml:"updateUserId,omitempty"`
	Value        *string `json:"value,omitempty" xml:"value,omitempty"`
	// This parameter is required.
	VisibleByEmp *bool `json:"visibleByEmp,omitempty" xml:"visibleByEmp,omitempty"`
}

func (s UpdateCustomRosterFieldRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomRosterFieldRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomRosterFieldRequest) SetContactClientFlag(v bool) *UpdateCustomRosterFieldRequest {
	s.ContactClientFlag = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetContactFlag(v bool) *UpdateCustomRosterFieldRequest {
	s.ContactFlag = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetContactSource(v int32) *UpdateCustomRosterFieldRequest {
	s.ContactSource = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetContactSystemFlag(v bool) *UpdateCustomRosterFieldRequest {
	s.ContactSystemFlag = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetDeleted(v bool) *UpdateCustomRosterFieldRequest {
	s.Deleted = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetDerived(v bool) *UpdateCustomRosterFieldRequest {
	s.Derived = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetDisabled(v int32) *UpdateCustomRosterFieldRequest {
	s.Disabled = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetEditFromEmployeeFlag(v bool) *UpdateCustomRosterFieldRequest {
	s.EditFromEmployeeFlag = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetEditableByHr(v bool) *UpdateCustomRosterFieldRequest {
	s.EditableByHr = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetFieldCode(v string) *UpdateCustomRosterFieldRequest {
	s.FieldCode = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetFieldName(v string) *UpdateCustomRosterFieldRequest {
	s.FieldName = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetFieldTip(v string) *UpdateCustomRosterFieldRequest {
	s.FieldTip = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetFieldType(v string) *UpdateCustomRosterFieldRequest {
	s.FieldType = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetGroupId(v string) *UpdateCustomRosterFieldRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetHiddenFromEmployeeFlag(v bool) *UpdateCustomRosterFieldRequest {
	s.HiddenFromEmployeeFlag = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetHint(v string) *UpdateCustomRosterFieldRequest {
	s.Hint = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetHistoryField(v bool) *UpdateCustomRosterFieldRequest {
	s.HistoryField = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetIndex(v int32) *UpdateCustomRosterFieldRequest {
	s.Index = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetModifyOptions(v bool) *UpdateCustomRosterFieldRequest {
	s.ModifyOptions = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetNoWatermark(v bool) *UpdateCustomRosterFieldRequest {
	s.NoWatermark = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetNumberDecimalPlace(v int32) *UpdateCustomRosterFieldRequest {
	s.NumberDecimalPlace = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetNumberFormatType(v string) *UpdateCustomRosterFieldRequest {
	s.NumberFormatType = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetNumberValueType(v string) *UpdateCustomRosterFieldRequest {
	s.NumberValueType = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetOptionText(v string) *UpdateCustomRosterFieldRequest {
	s.OptionText = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetRequired(v bool) *UpdateCustomRosterFieldRequest {
	s.Required = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetSourceFieldCode(v string) *UpdateCustomRosterFieldRequest {
	s.SourceFieldCode = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetSystemFlag(v bool) *UpdateCustomRosterFieldRequest {
	s.SystemFlag = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetTextToSelectField(v bool) *UpdateCustomRosterFieldRequest {
	s.TextToSelectField = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetUpdateUserId(v string) *UpdateCustomRosterFieldRequest {
	s.UpdateUserId = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetValue(v string) *UpdateCustomRosterFieldRequest {
	s.Value = &v
	return s
}

func (s *UpdateCustomRosterFieldRequest) SetVisibleByEmp(v bool) *UpdateCustomRosterFieldRequest {
	s.VisibleByEmp = &v
	return s
}

type UpdateCustomRosterFieldResponseBody struct {
	DingOpenErrcode *int32                                     `json:"dingOpenErrcode,omitempty" xml:"dingOpenErrcode,omitempty"`
	ErrorMsg        *string                                    `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result          *UpdateCustomRosterFieldResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success         *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateCustomRosterFieldResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomRosterFieldResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomRosterFieldResponseBody) SetDingOpenErrcode(v int32) *UpdateCustomRosterFieldResponseBody {
	s.DingOpenErrcode = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBody) SetErrorMsg(v string) *UpdateCustomRosterFieldResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBody) SetResult(v *UpdateCustomRosterFieldResponseBodyResult) *UpdateCustomRosterFieldResponseBody {
	s.Result = v
	return s
}

func (s *UpdateCustomRosterFieldResponseBody) SetSuccess(v bool) *UpdateCustomRosterFieldResponseBody {
	s.Success = &v
	return s
}

type UpdateCustomRosterFieldResponseBodyResult struct {
	ContactClientFlag      *bool   `json:"contactClientFlag,omitempty" xml:"contactClientFlag,omitempty"`
	ContactFlag            *bool   `json:"contactFlag,omitempty" xml:"contactFlag,omitempty"`
	ContactSource          *int32  `json:"contactSource,omitempty" xml:"contactSource,omitempty"`
	ContactSystemFlag      *bool   `json:"contactSystemFlag,omitempty" xml:"contactSystemFlag,omitempty"`
	Deleted                *bool   `json:"deleted,omitempty" xml:"deleted,omitempty"`
	Derived                *bool   `json:"derived,omitempty" xml:"derived,omitempty"`
	Disabled               *int32  `json:"disabled,omitempty" xml:"disabled,omitempty"`
	EditFromEmployeeFlag   *bool   `json:"editFromEmployeeFlag,omitempty" xml:"editFromEmployeeFlag,omitempty"`
	EditableByHr           *bool   `json:"editableByHr,omitempty" xml:"editableByHr,omitempty"`
	FieldCode              *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	FieldName              *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	FieldTip               *string `json:"fieldTip,omitempty" xml:"fieldTip,omitempty"`
	FieldType              *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	GroupId                *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	HiddenFromEmployeeFlag *bool   `json:"hiddenFromEmployeeFlag,omitempty" xml:"hiddenFromEmployeeFlag,omitempty"`
	Hint                   *string `json:"hint,omitempty" xml:"hint,omitempty"`
	HistoryField           *bool   `json:"historyField,omitempty" xml:"historyField,omitempty"`
	Index                  *int32  `json:"index,omitempty" xml:"index,omitempty"`
	ModifyOptions          *bool   `json:"modifyOptions,omitempty" xml:"modifyOptions,omitempty"`
	NoWatermark            *bool   `json:"noWatermark,omitempty" xml:"noWatermark,omitempty"`
	NumberDecimalPlace     *string `json:"numberDecimalPlace,omitempty" xml:"numberDecimalPlace,omitempty"`
	NumberFormatType       *string `json:"numberFormatType,omitempty" xml:"numberFormatType,omitempty"`
	NumberValueType        *string `json:"numberValueType,omitempty" xml:"numberValueType,omitempty"`
	OptionText             *string `json:"optionText,omitempty" xml:"optionText,omitempty"`
	Required               *bool   `json:"required,omitempty" xml:"required,omitempty"`
	SourceFieldCode        *string `json:"sourceFieldCode,omitempty" xml:"sourceFieldCode,omitempty"`
	SystemFlag             *bool   `json:"systemFlag,omitempty" xml:"systemFlag,omitempty"`
	TextToSelectField      *bool   `json:"textToSelectField,omitempty" xml:"textToSelectField,omitempty"`
	Value                  *string `json:"value,omitempty" xml:"value,omitempty"`
	VisibleByEmp           *bool   `json:"visibleByEmp,omitempty" xml:"visibleByEmp,omitempty"`
}

func (s UpdateCustomRosterFieldResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomRosterFieldResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetContactClientFlag(v bool) *UpdateCustomRosterFieldResponseBodyResult {
	s.ContactClientFlag = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetContactFlag(v bool) *UpdateCustomRosterFieldResponseBodyResult {
	s.ContactFlag = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetContactSource(v int32) *UpdateCustomRosterFieldResponseBodyResult {
	s.ContactSource = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetContactSystemFlag(v bool) *UpdateCustomRosterFieldResponseBodyResult {
	s.ContactSystemFlag = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetDeleted(v bool) *UpdateCustomRosterFieldResponseBodyResult {
	s.Deleted = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetDerived(v bool) *UpdateCustomRosterFieldResponseBodyResult {
	s.Derived = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetDisabled(v int32) *UpdateCustomRosterFieldResponseBodyResult {
	s.Disabled = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetEditFromEmployeeFlag(v bool) *UpdateCustomRosterFieldResponseBodyResult {
	s.EditFromEmployeeFlag = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetEditableByHr(v bool) *UpdateCustomRosterFieldResponseBodyResult {
	s.EditableByHr = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetFieldCode(v string) *UpdateCustomRosterFieldResponseBodyResult {
	s.FieldCode = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetFieldName(v string) *UpdateCustomRosterFieldResponseBodyResult {
	s.FieldName = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetFieldTip(v string) *UpdateCustomRosterFieldResponseBodyResult {
	s.FieldTip = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetFieldType(v string) *UpdateCustomRosterFieldResponseBodyResult {
	s.FieldType = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetGroupId(v string) *UpdateCustomRosterFieldResponseBodyResult {
	s.GroupId = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetHiddenFromEmployeeFlag(v bool) *UpdateCustomRosterFieldResponseBodyResult {
	s.HiddenFromEmployeeFlag = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetHint(v string) *UpdateCustomRosterFieldResponseBodyResult {
	s.Hint = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetHistoryField(v bool) *UpdateCustomRosterFieldResponseBodyResult {
	s.HistoryField = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetIndex(v int32) *UpdateCustomRosterFieldResponseBodyResult {
	s.Index = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetModifyOptions(v bool) *UpdateCustomRosterFieldResponseBodyResult {
	s.ModifyOptions = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetNoWatermark(v bool) *UpdateCustomRosterFieldResponseBodyResult {
	s.NoWatermark = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetNumberDecimalPlace(v string) *UpdateCustomRosterFieldResponseBodyResult {
	s.NumberDecimalPlace = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetNumberFormatType(v string) *UpdateCustomRosterFieldResponseBodyResult {
	s.NumberFormatType = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetNumberValueType(v string) *UpdateCustomRosterFieldResponseBodyResult {
	s.NumberValueType = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetOptionText(v string) *UpdateCustomRosterFieldResponseBodyResult {
	s.OptionText = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetRequired(v bool) *UpdateCustomRosterFieldResponseBodyResult {
	s.Required = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetSourceFieldCode(v string) *UpdateCustomRosterFieldResponseBodyResult {
	s.SourceFieldCode = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetSystemFlag(v bool) *UpdateCustomRosterFieldResponseBodyResult {
	s.SystemFlag = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetTextToSelectField(v bool) *UpdateCustomRosterFieldResponseBodyResult {
	s.TextToSelectField = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetValue(v string) *UpdateCustomRosterFieldResponseBodyResult {
	s.Value = &v
	return s
}

func (s *UpdateCustomRosterFieldResponseBodyResult) SetVisibleByEmp(v bool) *UpdateCustomRosterFieldResponseBodyResult {
	s.VisibleByEmp = &v
	return s
}

type UpdateCustomRosterFieldResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomRosterFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomRosterFieldResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomRosterFieldResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomRosterFieldResponse) SetHeaders(v map[string]*string) *UpdateCustomRosterFieldResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomRosterFieldResponse) SetStatusCode(v int32) *UpdateCustomRosterFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomRosterFieldResponse) SetBody(v *UpdateCustomRosterFieldResponseBody) *UpdateCustomRosterFieldResponse {
	s.Body = v
	return s
}

type UpdateEmpDismissionInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateEmpDismissionInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateEmpDismissionInfoHeaders) GoString() string {
	return s.String()
}

func (s *UpdateEmpDismissionInfoHeaders) SetCommonHeaders(v map[string]*string) *UpdateEmpDismissionInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateEmpDismissionInfoHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateEmpDismissionInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateEmpDismissionInfoRequest struct {
	DismissionMemo *string `json:"dismissionMemo,omitempty" xml:"dismissionMemo,omitempty"`
	// This parameter is required.
	LastWorkDate               *int64    `json:"lastWorkDate,omitempty" xml:"lastWorkDate,omitempty"`
	Partner                    *bool     `json:"partner,omitempty" xml:"partner,omitempty"`
	TerminationReasonPassive   []*string `json:"terminationReasonPassive,omitempty" xml:"terminationReasonPassive,omitempty" type:"Repeated"`
	TerminationReasonVoluntary []*string `json:"terminationReasonVoluntary,omitempty" xml:"terminationReasonVoluntary,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 2163515669935611
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateEmpDismissionInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEmpDismissionInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateEmpDismissionInfoRequest) SetDismissionMemo(v string) *UpdateEmpDismissionInfoRequest {
	s.DismissionMemo = &v
	return s
}

func (s *UpdateEmpDismissionInfoRequest) SetLastWorkDate(v int64) *UpdateEmpDismissionInfoRequest {
	s.LastWorkDate = &v
	return s
}

func (s *UpdateEmpDismissionInfoRequest) SetPartner(v bool) *UpdateEmpDismissionInfoRequest {
	s.Partner = &v
	return s
}

func (s *UpdateEmpDismissionInfoRequest) SetTerminationReasonPassive(v []*string) *UpdateEmpDismissionInfoRequest {
	s.TerminationReasonPassive = v
	return s
}

func (s *UpdateEmpDismissionInfoRequest) SetTerminationReasonVoluntary(v []*string) *UpdateEmpDismissionInfoRequest {
	s.TerminationReasonVoluntary = v
	return s
}

func (s *UpdateEmpDismissionInfoRequest) SetUserId(v string) *UpdateEmpDismissionInfoRequest {
	s.UserId = &v
	return s
}

type UpdateEmpDismissionInfoResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateEmpDismissionInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEmpDismissionInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEmpDismissionInfoResponseBody) SetRequestId(v string) *UpdateEmpDismissionInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEmpDismissionInfoResponseBody) SetResult(v bool) *UpdateEmpDismissionInfoResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateEmpDismissionInfoResponseBody) SetSuccess(v bool) *UpdateEmpDismissionInfoResponseBody {
	s.Success = &v
	return s
}

type UpdateEmpDismissionInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEmpDismissionInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEmpDismissionInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEmpDismissionInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateEmpDismissionInfoResponse) SetHeaders(v map[string]*string) *UpdateEmpDismissionInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateEmpDismissionInfoResponse) SetStatusCode(v int32) *UpdateEmpDismissionInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEmpDismissionInfoResponse) SetBody(v *UpdateEmpDismissionInfoResponseBody) *UpdateEmpDismissionInfoResponse {
	s.Body = v
	return s
}

type UpdateHrmLegalEntityNameHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateHrmLegalEntityNameHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateHrmLegalEntityNameHeaders) GoString() string {
	return s.String()
}

func (s *UpdateHrmLegalEntityNameHeaders) SetCommonHeaders(v map[string]*string) *UpdateHrmLegalEntityNameHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateHrmLegalEntityNameHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateHrmLegalEntityNameHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateHrmLegalEntityNameRequest struct {
	// example:
	//
	// 57
	DingTenantId *int64 `json:"dingTenantId,omitempty" xml:"dingTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 公司2
	LegalEntityName *string `json:"legalEntityName,omitempty" xml:"legalEntityName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 公司1
	OriginLegalEntityName *string `json:"originLegalEntityName,omitempty" xml:"originLegalEntityName,omitempty"`
}

func (s UpdateHrmLegalEntityNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHrmLegalEntityNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateHrmLegalEntityNameRequest) SetDingTenantId(v int64) *UpdateHrmLegalEntityNameRequest {
	s.DingTenantId = &v
	return s
}

func (s *UpdateHrmLegalEntityNameRequest) SetLegalEntityName(v string) *UpdateHrmLegalEntityNameRequest {
	s.LegalEntityName = &v
	return s
}

func (s *UpdateHrmLegalEntityNameRequest) SetOriginLegalEntityName(v string) *UpdateHrmLegalEntityNameRequest {
	s.OriginLegalEntityName = &v
	return s
}

type UpdateHrmLegalEntityNameResponseBody struct {
	Result  *UpdateHrmLegalEntityNameResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                       `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateHrmLegalEntityNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateHrmLegalEntityNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHrmLegalEntityNameResponseBody) SetResult(v *UpdateHrmLegalEntityNameResponseBodyResult) *UpdateHrmLegalEntityNameResponseBody {
	s.Result = v
	return s
}

func (s *UpdateHrmLegalEntityNameResponseBody) SetSuccess(v bool) *UpdateHrmLegalEntityNameResponseBody {
	s.Success = &v
	return s
}

type UpdateHrmLegalEntityNameResponseBodyResult struct {
	// example:
	//
	// ding123
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// 2023-08-08
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2023-08-08
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 111233
	LegalEntityId *string `json:"legalEntityId,omitempty" xml:"legalEntityId,omitempty"`
	// example:
	//
	// 公司2
	LegalEntityName *string `json:"legalEntityName,omitempty" xml:"legalEntityName,omitempty"`
	// example:
	//
	// 公2
	LegalEntityShortName *string `json:"legalEntityShortName,omitempty" xml:"legalEntityShortName,omitempty"`
	// example:
	//
	// 1
	LegalEntityStatus *int32 `json:"legalEntityStatus,omitempty" xml:"legalEntityStatus,omitempty"`
	// example:
	//
	// 法人1
	LegalPersonName *string `json:"legalPersonName,omitempty" xml:"legalPersonName,omitempty"`
}

func (s UpdateHrmLegalEntityNameResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateHrmLegalEntityNameResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateHrmLegalEntityNameResponseBodyResult) SetCorpId(v string) *UpdateHrmLegalEntityNameResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *UpdateHrmLegalEntityNameResponseBodyResult) SetGmtCreate(v int64) *UpdateHrmLegalEntityNameResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *UpdateHrmLegalEntityNameResponseBodyResult) SetGmtModified(v int64) *UpdateHrmLegalEntityNameResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *UpdateHrmLegalEntityNameResponseBodyResult) SetLegalEntityId(v string) *UpdateHrmLegalEntityNameResponseBodyResult {
	s.LegalEntityId = &v
	return s
}

func (s *UpdateHrmLegalEntityNameResponseBodyResult) SetLegalEntityName(v string) *UpdateHrmLegalEntityNameResponseBodyResult {
	s.LegalEntityName = &v
	return s
}

func (s *UpdateHrmLegalEntityNameResponseBodyResult) SetLegalEntityShortName(v string) *UpdateHrmLegalEntityNameResponseBodyResult {
	s.LegalEntityShortName = &v
	return s
}

func (s *UpdateHrmLegalEntityNameResponseBodyResult) SetLegalEntityStatus(v int32) *UpdateHrmLegalEntityNameResponseBodyResult {
	s.LegalEntityStatus = &v
	return s
}

func (s *UpdateHrmLegalEntityNameResponseBodyResult) SetLegalPersonName(v string) *UpdateHrmLegalEntityNameResponseBodyResult {
	s.LegalPersonName = &v
	return s
}

type UpdateHrmLegalEntityNameResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHrmLegalEntityNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHrmLegalEntityNameResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateHrmLegalEntityNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateHrmLegalEntityNameResponse) SetHeaders(v map[string]*string) *UpdateHrmLegalEntityNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateHrmLegalEntityNameResponse) SetStatusCode(v int32) *UpdateHrmLegalEntityNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHrmLegalEntityNameResponse) SetBody(v *UpdateHrmLegalEntityNameResponseBody) *UpdateHrmLegalEntityNameResponse {
	s.Body = v
	return s
}

type UpdateHrmLegalEntityWithoutNameHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateHrmLegalEntityWithoutNameHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateHrmLegalEntityWithoutNameHeaders) GoString() string {
	return s.String()
}

func (s *UpdateHrmLegalEntityWithoutNameHeaders) SetCommonHeaders(v map[string]*string) *UpdateHrmLegalEntityWithoutNameHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateHrmLegalEntityWithoutNameHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateHrmLegalEntityWithoutNameRequest struct {
	// This parameter is required.
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// 123
	CreateUserId *string                                    `json:"createUserId,omitempty" xml:"createUserId,omitempty"`
	Ext          *UpdateHrmLegalEntityWithoutNameRequestExt `json:"ext,omitempty" xml:"ext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 公司1
	LegalEntityName *string `json:"legalEntityName,omitempty" xml:"legalEntityName,omitempty"`
	// example:
	//
	// 公1
	LegalEntityShortName *string `json:"legalEntityShortName,omitempty" xml:"legalEntityShortName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	LegalEntityStatus *int32 `json:"legalEntityStatus,omitempty" xml:"legalEntityStatus,omitempty"`
	// example:
	//
	// 法人
	LegalPersonName *string `json:"legalPersonName,omitempty" xml:"legalPersonName,omitempty"`
	// example:
	//
	// 57
	DingTenantId *int64 `json:"dingTenantId,omitempty" xml:"dingTenantId,omitempty"`
}

func (s UpdateHrmLegalEntityWithoutNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHrmLegalEntityWithoutNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateHrmLegalEntityWithoutNameRequest) SetCorpId(v string) *UpdateHrmLegalEntityWithoutNameRequest {
	s.CorpId = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequest) SetCreateUserId(v string) *UpdateHrmLegalEntityWithoutNameRequest {
	s.CreateUserId = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequest) SetExt(v *UpdateHrmLegalEntityWithoutNameRequestExt) *UpdateHrmLegalEntityWithoutNameRequest {
	s.Ext = v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequest) SetLegalEntityName(v string) *UpdateHrmLegalEntityWithoutNameRequest {
	s.LegalEntityName = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequest) SetLegalEntityShortName(v string) *UpdateHrmLegalEntityWithoutNameRequest {
	s.LegalEntityShortName = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequest) SetLegalEntityStatus(v int32) *UpdateHrmLegalEntityWithoutNameRequest {
	s.LegalEntityStatus = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequest) SetLegalPersonName(v string) *UpdateHrmLegalEntityWithoutNameRequest {
	s.LegalPersonName = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequest) SetDingTenantId(v int64) *UpdateHrmLegalEntityWithoutNameRequest {
	s.DingTenantId = &v
	return s
}

type UpdateHrmLegalEntityWithoutNameRequestExt struct {
	// example:
	//
	// company
	LegalEntityEnName *string `json:"legalEntityEnName,omitempty" xml:"legalEntityEnName,omitempty"`
	// example:
	//
	// com
	LegalEntityEnShortName *string `json:"legalEntityEnShortName,omitempty" xml:"legalEntityEnShortName,omitempty"`
	// example:
	//
	// whollyOwned
	LegalEntityType     *string                                                       `json:"legalEntityType,omitempty" xml:"legalEntityType,omitempty"`
	ManageAddress       *UpdateHrmLegalEntityWithoutNameRequestExtManageAddress       `json:"manageAddress,omitempty" xml:"manageAddress,omitempty" type:"Struct"`
	RegistrationAddress *UpdateHrmLegalEntityWithoutNameRequestExtRegistrationAddress `json:"registrationAddress,omitempty" xml:"registrationAddress,omitempty" type:"Struct"`
	// example:
	//
	// 2023-01-01
	RegistrationDate *int64 `json:"registrationDate,omitempty" xml:"registrationDate,omitempty"`
	// example:
	//
	// 123456
	UnifiedSocialCreditCode *string `json:"unifiedSocialCreditCode,omitempty" xml:"unifiedSocialCreditCode,omitempty"`
	// example:
	//
	// 515200
	ZipCode *string `json:"zipCode,omitempty" xml:"zipCode,omitempty"`
}

func (s UpdateHrmLegalEntityWithoutNameRequestExt) String() string {
	return tea.Prettify(s)
}

func (s UpdateHrmLegalEntityWithoutNameRequestExt) GoString() string {
	return s.String()
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExt) SetLegalEntityEnName(v string) *UpdateHrmLegalEntityWithoutNameRequestExt {
	s.LegalEntityEnName = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExt) SetLegalEntityEnShortName(v string) *UpdateHrmLegalEntityWithoutNameRequestExt {
	s.LegalEntityEnShortName = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExt) SetLegalEntityType(v string) *UpdateHrmLegalEntityWithoutNameRequestExt {
	s.LegalEntityType = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExt) SetManageAddress(v *UpdateHrmLegalEntityWithoutNameRequestExtManageAddress) *UpdateHrmLegalEntityWithoutNameRequestExt {
	s.ManageAddress = v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExt) SetRegistrationAddress(v *UpdateHrmLegalEntityWithoutNameRequestExtRegistrationAddress) *UpdateHrmLegalEntityWithoutNameRequestExt {
	s.RegistrationAddress = v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExt) SetRegistrationDate(v int64) *UpdateHrmLegalEntityWithoutNameRequestExt {
	s.RegistrationDate = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExt) SetUnifiedSocialCreditCode(v string) *UpdateHrmLegalEntityWithoutNameRequestExt {
	s.UnifiedSocialCreditCode = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExt) SetZipCode(v string) *UpdateHrmLegalEntityWithoutNameRequestExt {
	s.ZipCode = &v
	return s
}

type UpdateHrmLegalEntityWithoutNameRequestExtManageAddress struct {
	// example:
	//
	// 110101
	AreaCode *string `json:"areaCode,omitempty" xml:"areaCode,omitempty"`
	// example:
	//
	// 东城区
	AreaName *string `json:"areaName,omitempty" xml:"areaName,omitempty"`
	// example:
	//
	// 1234
	CityCode *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	// example:
	//
	// 广州
	CityName *string `json:"cityName,omitempty" xml:"cityName,omitempty"`
	// example:
	//
	// 123
	CountryCode *string `json:"countryCode,omitempty" xml:"countryCode,omitempty"`
	// example:
	//
	// China
	CountryName *string `json:"countryName,omitempty" xml:"countryName,omitempty"`
	// example:
	//
	// 北京市东城区xx街道xx小区xx楼
	DetailAddress *string `json:"detailAddress,omitempty" xml:"detailAddress,omitempty"`
	// example:
	//
	// 1
	GlobalAreaType *string `json:"globalAreaType,omitempty" xml:"globalAreaType,omitempty"`
	// example:
	//
	// 123
	ProvinceCode *string `json:"provinceCode,omitempty" xml:"provinceCode,omitempty"`
	// example:
	//
	// 广东
	ProvinceName *string `json:"provinceName,omitempty" xml:"provinceName,omitempty"`
}

func (s UpdateHrmLegalEntityWithoutNameRequestExtManageAddress) String() string {
	return tea.Prettify(s)
}

func (s UpdateHrmLegalEntityWithoutNameRequestExtManageAddress) GoString() string {
	return s.String()
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExtManageAddress) SetAreaCode(v string) *UpdateHrmLegalEntityWithoutNameRequestExtManageAddress {
	s.AreaCode = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExtManageAddress) SetAreaName(v string) *UpdateHrmLegalEntityWithoutNameRequestExtManageAddress {
	s.AreaName = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExtManageAddress) SetCityCode(v string) *UpdateHrmLegalEntityWithoutNameRequestExtManageAddress {
	s.CityCode = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExtManageAddress) SetCityName(v string) *UpdateHrmLegalEntityWithoutNameRequestExtManageAddress {
	s.CityName = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExtManageAddress) SetCountryCode(v string) *UpdateHrmLegalEntityWithoutNameRequestExtManageAddress {
	s.CountryCode = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExtManageAddress) SetCountryName(v string) *UpdateHrmLegalEntityWithoutNameRequestExtManageAddress {
	s.CountryName = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExtManageAddress) SetDetailAddress(v string) *UpdateHrmLegalEntityWithoutNameRequestExtManageAddress {
	s.DetailAddress = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExtManageAddress) SetGlobalAreaType(v string) *UpdateHrmLegalEntityWithoutNameRequestExtManageAddress {
	s.GlobalAreaType = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExtManageAddress) SetProvinceCode(v string) *UpdateHrmLegalEntityWithoutNameRequestExtManageAddress {
	s.ProvinceCode = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExtManageAddress) SetProvinceName(v string) *UpdateHrmLegalEntityWithoutNameRequestExtManageAddress {
	s.ProvinceName = &v
	return s
}

type UpdateHrmLegalEntityWithoutNameRequestExtRegistrationAddress struct {
	// example:
	//
	// 110101
	AreaCode *string `json:"areaCode,omitempty" xml:"areaCode,omitempty"`
	// example:
	//
	// 东城区
	AreaName *string `json:"areaName,omitempty" xml:"areaName,omitempty"`
	// example:
	//
	// 1234
	CityCode *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	// example:
	//
	// 广州
	CityName *string `json:"cityName,omitempty" xml:"cityName,omitempty"`
	// example:
	//
	// 123
	CountryCode *string `json:"countryCode,omitempty" xml:"countryCode,omitempty"`
	// example:
	//
	// China
	CountryName *string `json:"countryName,omitempty" xml:"countryName,omitempty"`
	// example:
	//
	// 北京市东城区xx街道xx小区xx楼
	DetailAddress *string `json:"detailAddress,omitempty" xml:"detailAddress,omitempty"`
	// example:
	//
	// 1
	GlobalAreaType *string `json:"globalAreaType,omitempty" xml:"globalAreaType,omitempty"`
	// example:
	//
	// 123
	ProvinceCode *string `json:"provinceCode,omitempty" xml:"provinceCode,omitempty"`
	// example:
	//
	// 广东
	ProvinceName *string `json:"provinceName,omitempty" xml:"provinceName,omitempty"`
}

func (s UpdateHrmLegalEntityWithoutNameRequestExtRegistrationAddress) String() string {
	return tea.Prettify(s)
}

func (s UpdateHrmLegalEntityWithoutNameRequestExtRegistrationAddress) GoString() string {
	return s.String()
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExtRegistrationAddress) SetAreaCode(v string) *UpdateHrmLegalEntityWithoutNameRequestExtRegistrationAddress {
	s.AreaCode = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExtRegistrationAddress) SetAreaName(v string) *UpdateHrmLegalEntityWithoutNameRequestExtRegistrationAddress {
	s.AreaName = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExtRegistrationAddress) SetCityCode(v string) *UpdateHrmLegalEntityWithoutNameRequestExtRegistrationAddress {
	s.CityCode = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExtRegistrationAddress) SetCityName(v string) *UpdateHrmLegalEntityWithoutNameRequestExtRegistrationAddress {
	s.CityName = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExtRegistrationAddress) SetCountryCode(v string) *UpdateHrmLegalEntityWithoutNameRequestExtRegistrationAddress {
	s.CountryCode = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExtRegistrationAddress) SetCountryName(v string) *UpdateHrmLegalEntityWithoutNameRequestExtRegistrationAddress {
	s.CountryName = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExtRegistrationAddress) SetDetailAddress(v string) *UpdateHrmLegalEntityWithoutNameRequestExtRegistrationAddress {
	s.DetailAddress = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExtRegistrationAddress) SetGlobalAreaType(v string) *UpdateHrmLegalEntityWithoutNameRequestExtRegistrationAddress {
	s.GlobalAreaType = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExtRegistrationAddress) SetProvinceCode(v string) *UpdateHrmLegalEntityWithoutNameRequestExtRegistrationAddress {
	s.ProvinceCode = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameRequestExtRegistrationAddress) SetProvinceName(v string) *UpdateHrmLegalEntityWithoutNameRequestExtRegistrationAddress {
	s.ProvinceName = &v
	return s
}

type UpdateHrmLegalEntityWithoutNameResponseBody struct {
	Result  *UpdateHrmLegalEntityWithoutNameResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateHrmLegalEntityWithoutNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateHrmLegalEntityWithoutNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHrmLegalEntityWithoutNameResponseBody) SetResult(v *UpdateHrmLegalEntityWithoutNameResponseBodyResult) *UpdateHrmLegalEntityWithoutNameResponseBody {
	s.Result = v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameResponseBody) SetSuccess(v bool) *UpdateHrmLegalEntityWithoutNameResponseBody {
	s.Success = &v
	return s
}

type UpdateHrmLegalEntityWithoutNameResponseBodyResult struct {
	// example:
	//
	// ding123
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// 2023-01-01
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2023-01-01
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 123456
	LegalEntityId *string `json:"legalEntityId,omitempty" xml:"legalEntityId,omitempty"`
	// example:
	//
	// 公司1
	LegalEntityName *string `json:"legalEntityName,omitempty" xml:"legalEntityName,omitempty"`
	// example:
	//
	// 公1
	LegalEntityShortName *string `json:"legalEntityShortName,omitempty" xml:"legalEntityShortName,omitempty"`
	// example:
	//
	// 1
	LegalEntityStatus *int32 `json:"legalEntityStatus,omitempty" xml:"legalEntityStatus,omitempty"`
	// example:
	//
	// 法人
	LegalPersonName *string `json:"legalPersonName,omitempty" xml:"legalPersonName,omitempty"`
}

func (s UpdateHrmLegalEntityWithoutNameResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateHrmLegalEntityWithoutNameResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateHrmLegalEntityWithoutNameResponseBodyResult) SetCorpId(v string) *UpdateHrmLegalEntityWithoutNameResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameResponseBodyResult) SetGmtCreate(v int64) *UpdateHrmLegalEntityWithoutNameResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameResponseBodyResult) SetGmtModified(v int64) *UpdateHrmLegalEntityWithoutNameResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameResponseBodyResult) SetLegalEntityId(v string) *UpdateHrmLegalEntityWithoutNameResponseBodyResult {
	s.LegalEntityId = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameResponseBodyResult) SetLegalEntityName(v string) *UpdateHrmLegalEntityWithoutNameResponseBodyResult {
	s.LegalEntityName = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameResponseBodyResult) SetLegalEntityShortName(v string) *UpdateHrmLegalEntityWithoutNameResponseBodyResult {
	s.LegalEntityShortName = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameResponseBodyResult) SetLegalEntityStatus(v int32) *UpdateHrmLegalEntityWithoutNameResponseBodyResult {
	s.LegalEntityStatus = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameResponseBodyResult) SetLegalPersonName(v string) *UpdateHrmLegalEntityWithoutNameResponseBodyResult {
	s.LegalPersonName = &v
	return s
}

type UpdateHrmLegalEntityWithoutNameResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHrmLegalEntityWithoutNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHrmLegalEntityWithoutNameResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateHrmLegalEntityWithoutNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateHrmLegalEntityWithoutNameResponse) SetHeaders(v map[string]*string) *UpdateHrmLegalEntityWithoutNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameResponse) SetStatusCode(v int32) *UpdateHrmLegalEntityWithoutNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHrmLegalEntityWithoutNameResponse) SetBody(v *UpdateHrmLegalEntityWithoutNameResponseBody) *UpdateHrmLegalEntityWithoutNameResponse {
	s.Body = v
	return s
}

type UpdateHrmVersionRollBackStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateHrmVersionRollBackStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateHrmVersionRollBackStatusHeaders) GoString() string {
	return s.String()
}

func (s *UpdateHrmVersionRollBackStatusHeaders) SetCommonHeaders(v map[string]*string) *UpdateHrmVersionRollBackStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateHrmVersionRollBackStatusHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateHrmVersionRollBackStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateHrmVersionRollBackStatusRequest struct {
	// example:
	//
	// show
	ConfigValue *string `json:"configValue,omitempty" xml:"configValue,omitempty"`
	// example:
	//
	// 1231231232
	OptUserId *string `json:"optUserId,omitempty" xml:"optUserId,omitempty"`
}

func (s UpdateHrmVersionRollBackStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHrmVersionRollBackStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateHrmVersionRollBackStatusRequest) SetConfigValue(v string) *UpdateHrmVersionRollBackStatusRequest {
	s.ConfigValue = &v
	return s
}

func (s *UpdateHrmVersionRollBackStatusRequest) SetOptUserId(v string) *UpdateHrmVersionRollBackStatusRequest {
	s.OptUserId = &v
	return s
}

type UpdateHrmVersionRollBackStatusResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateHrmVersionRollBackStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateHrmVersionRollBackStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHrmVersionRollBackStatusResponseBody) SetRequestId(v string) *UpdateHrmVersionRollBackStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHrmVersionRollBackStatusResponseBody) SetResult(v bool) *UpdateHrmVersionRollBackStatusResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateHrmVersionRollBackStatusResponseBody) SetSuccess(v bool) *UpdateHrmVersionRollBackStatusResponseBody {
	s.Success = &v
	return s
}

type UpdateHrmVersionRollBackStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHrmVersionRollBackStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHrmVersionRollBackStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateHrmVersionRollBackStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateHrmVersionRollBackStatusResponse) SetHeaders(v map[string]*string) *UpdateHrmVersionRollBackStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateHrmVersionRollBackStatusResponse) SetStatusCode(v int32) *UpdateHrmVersionRollBackStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHrmVersionRollBackStatusResponse) SetBody(v *UpdateHrmVersionRollBackStatusResponseBody) *UpdateHrmVersionRollBackStatusResponse {
	s.Body = v
	return s
}

type UpdateIsvCardMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateIsvCardMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateIsvCardMessageHeaders) GoString() string {
	return s.String()
}

func (s *UpdateIsvCardMessageHeaders) SetCommonHeaders(v map[string]*string) *UpdateIsvCardMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateIsvCardMessageHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateIsvCardMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateIsvCardMessageRequest struct {
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// This parameter is required.
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// This parameter is required.
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 16690147049882572
	SceneType *string `json:"sceneType,omitempty" xml:"sceneType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 同意转正
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// This parameter is required.
	ValueMap map[string]*string `json:"valueMap,omitempty" xml:"valueMap,omitempty"`
}

func (s UpdateIsvCardMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIsvCardMessageRequest) GoString() string {
	return s.String()
}

func (s *UpdateIsvCardMessageRequest) SetAgentId(v int64) *UpdateIsvCardMessageRequest {
	s.AgentId = &v
	return s
}

func (s *UpdateIsvCardMessageRequest) SetBizId(v string) *UpdateIsvCardMessageRequest {
	s.BizId = &v
	return s
}

func (s *UpdateIsvCardMessageRequest) SetMessageType(v string) *UpdateIsvCardMessageRequest {
	s.MessageType = &v
	return s
}

func (s *UpdateIsvCardMessageRequest) SetSceneType(v string) *UpdateIsvCardMessageRequest {
	s.SceneType = &v
	return s
}

func (s *UpdateIsvCardMessageRequest) SetScope(v string) *UpdateIsvCardMessageRequest {
	s.Scope = &v
	return s
}

func (s *UpdateIsvCardMessageRequest) SetValueMap(v map[string]*string) *UpdateIsvCardMessageRequest {
	s.ValueMap = v
	return s
}

type UpdateIsvCardMessageResponseBody struct {
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateIsvCardMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIsvCardMessageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIsvCardMessageResponseBody) SetErrorCode(v string) *UpdateIsvCardMessageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateIsvCardMessageResponseBody) SetErrorMsg(v string) *UpdateIsvCardMessageResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateIsvCardMessageResponseBody) SetRequestId(v string) *UpdateIsvCardMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIsvCardMessageResponseBody) SetSuccess(v bool) *UpdateIsvCardMessageResponseBody {
	s.Success = &v
	return s
}

type UpdateIsvCardMessageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIsvCardMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIsvCardMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIsvCardMessageResponse) GoString() string {
	return s.String()
}

func (s *UpdateIsvCardMessageResponse) SetHeaders(v map[string]*string) *UpdateIsvCardMessageResponse {
	s.Headers = v
	return s
}

func (s *UpdateIsvCardMessageResponse) SetStatusCode(v int32) *UpdateIsvCardMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIsvCardMessageResponse) SetBody(v *UpdateIsvCardMessageResponseBody) *UpdateIsvCardMessageResponse {
	s.Body = v
	return s
}

type UpdateRosterFieldFormHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateRosterFieldFormHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateRosterFieldFormHeaders) GoString() string {
	return s.String()
}

func (s *UpdateRosterFieldFormHeaders) SetCommonHeaders(v map[string]*string) *UpdateRosterFieldFormHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateRosterFieldFormHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateRosterFieldFormHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateRosterFieldFormRequest struct {
	// This parameter is required.
	Detail *bool `json:"detail,omitempty" xml:"detail,omitempty"`
	// This parameter is required.
	FormId *string `json:"formId,omitempty" xml:"formId,omitempty"`
	// This parameter is required.
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateRosterFieldFormRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRosterFieldFormRequest) GoString() string {
	return s.String()
}

func (s *UpdateRosterFieldFormRequest) SetDetail(v bool) *UpdateRosterFieldFormRequest {
	s.Detail = &v
	return s
}

func (s *UpdateRosterFieldFormRequest) SetFormId(v string) *UpdateRosterFieldFormRequest {
	s.FormId = &v
	return s
}

func (s *UpdateRosterFieldFormRequest) SetName(v string) *UpdateRosterFieldFormRequest {
	s.Name = &v
	return s
}

func (s *UpdateRosterFieldFormRequest) SetUserId(v string) *UpdateRosterFieldFormRequest {
	s.UserId = &v
	return s
}

type UpdateRosterFieldFormResponseBody struct {
	DingOpenErrcode *int32  `json:"dingOpenErrcode,omitempty" xml:"dingOpenErrcode,omitempty"`
	ErrorMsg        *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result          *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success         *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateRosterFieldFormResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRosterFieldFormResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRosterFieldFormResponseBody) SetDingOpenErrcode(v int32) *UpdateRosterFieldFormResponseBody {
	s.DingOpenErrcode = &v
	return s
}

func (s *UpdateRosterFieldFormResponseBody) SetErrorMsg(v string) *UpdateRosterFieldFormResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateRosterFieldFormResponseBody) SetResult(v bool) *UpdateRosterFieldFormResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateRosterFieldFormResponseBody) SetSuccess(v bool) *UpdateRosterFieldFormResponseBody {
	s.Success = &v
	return s
}

type UpdateRosterFieldFormResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRosterFieldFormResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRosterFieldFormResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRosterFieldFormResponse) GoString() string {
	return s.String()
}

func (s *UpdateRosterFieldFormResponse) SetHeaders(v map[string]*string) *UpdateRosterFieldFormResponse {
	s.Headers = v
	return s
}

func (s *UpdateRosterFieldFormResponse) SetStatusCode(v int32) *UpdateRosterFieldFormResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRosterFieldFormResponse) SetBody(v *UpdateRosterFieldFormResponseBody) *UpdateRosterFieldFormResponse {
	s.Body = v
	return s
}

type UploadAttachmentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UploadAttachmentHeaders) String() string {
	return tea.Prettify(s)
}

func (s UploadAttachmentHeaders) GoString() string {
	return s.String()
}

func (s *UploadAttachmentHeaders) SetCommonHeaders(v map[string]*string) *UploadAttachmentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UploadAttachmentHeaders) SetXAcsDingtalkAccessToken(v string) *UploadAttachmentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UploadAttachmentRequest struct {
	// example:
	//
	// @dsa8d87y7c8d8c
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	// example:
	//
	// 16768800278994283
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UploadAttachmentRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadAttachmentRequest) GoString() string {
	return s.String()
}

func (s *UploadAttachmentRequest) SetMediaId(v string) *UploadAttachmentRequest {
	s.MediaId = &v
	return s
}

func (s *UploadAttachmentRequest) SetUserId(v string) *UploadAttachmentRequest {
	s.UserId = &v
	return s
}

type UploadAttachmentResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UploadAttachmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *UploadAttachmentResponseBody) SetResult(v string) *UploadAttachmentResponseBody {
	s.Result = &v
	return s
}

func (s *UploadAttachmentResponseBody) SetSuccess(v bool) *UploadAttachmentResponseBody {
	s.Success = &v
	return s
}

type UploadAttachmentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadAttachmentResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadAttachmentResponse) GoString() string {
	return s.String()
}

func (s *UploadAttachmentResponse) SetHeaders(v map[string]*string) *UploadAttachmentResponse {
	s.Headers = v
	return s
}

func (s *UploadAttachmentResponse) SetStatusCode(v int32) *UploadAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadAttachmentResponse) SetBody(v *UploadAttachmentResponseBody) *UploadAttachmentResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

// Summary:
//
// 添加自定义花名册字段
//
// @param request - AddCustomRosterFieldRequest
//
// @param headers - AddCustomRosterFieldHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCustomRosterFieldResponse
func (client *Client) AddCustomRosterFieldWithOptions(request *AddCustomRosterFieldRequest, headers *AddCustomRosterFieldHeaders, runtime *util.RuntimeOptions) (_result *AddCustomRosterFieldResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EditFromEmployeeFlag)) {
		body["editFromEmployeeFlag"] = request.EditFromEmployeeFlag
	}

	if !tea.BoolValue(util.IsUnset(request.FieldName)) {
		body["fieldName"] = request.FieldName
	}

	if !tea.BoolValue(util.IsUnset(request.FieldType)) {
		body["fieldType"] = request.FieldType
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["groupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.HiddenFromEmployeeFlag)) {
		body["hiddenFromEmployeeFlag"] = request.HiddenFromEmployeeFlag
	}

	if !tea.BoolValue(util.IsUnset(request.Hint)) {
		body["hint"] = request.Hint
	}

	if !tea.BoolValue(util.IsUnset(request.NoWatermark)) {
		body["noWatermark"] = request.NoWatermark
	}

	if !tea.BoolValue(util.IsUnset(request.NumberDecimalPlace)) {
		body["numberDecimalPlace"] = request.NumberDecimalPlace
	}

	if !tea.BoolValue(util.IsUnset(request.NumberFormatType)) {
		body["numberFormatType"] = request.NumberFormatType
	}

	if !tea.BoolValue(util.IsUnset(request.NumberValueType)) {
		body["numberValueType"] = request.NumberValueType
	}

	if !tea.BoolValue(util.IsUnset(request.OptionText)) {
		body["optionText"] = request.OptionText
	}

	if !tea.BoolValue(util.IsUnset(request.Required)) {
		body["required"] = request.Required
	}

	if !tea.BoolValue(util.IsUnset(request.VisibleByEmp)) {
		body["visibleByEmp"] = request.VisibleByEmp
	}

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
		Action:      tea.String("AddCustomRosterField"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/customRosterField/add"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddCustomRosterFieldResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加自定义花名册字段
//
// @param request - AddCustomRosterFieldRequest
//
// @return AddCustomRosterFieldResponse
func (client *Client) AddCustomRosterField(request *AddCustomRosterFieldRequest) (_result *AddCustomRosterFieldResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddCustomRosterFieldHeaders{}
	_result = &AddCustomRosterFieldResponse{}
	_body, _err := client.AddCustomRosterFieldWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增法人公司
//
// @param request - AddHrmLegalEntityRequest
//
// @param headers - AddHrmLegalEntityHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddHrmLegalEntityResponse
func (client *Client) AddHrmLegalEntityWithOptions(request *AddHrmLegalEntityRequest, headers *AddHrmLegalEntityHeaders, runtime *util.RuntimeOptions) (_result *AddHrmLegalEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingTenantId)) {
		query["dingTenantId"] = request.DingTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.CreateUserId)) {
		body["createUserId"] = request.CreateUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.LegalEntityName)) {
		body["legalEntityName"] = request.LegalEntityName
	}

	if !tea.BoolValue(util.IsUnset(request.LegalEntityShortName)) {
		body["legalEntityShortName"] = request.LegalEntityShortName
	}

	if !tea.BoolValue(util.IsUnset(request.LegalEntityStatus)) {
		body["legalEntityStatus"] = request.LegalEntityStatus
	}

	if !tea.BoolValue(util.IsUnset(request.LegalPersonName)) {
		body["legalPersonName"] = request.LegalPersonName
	}

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
		Action:      tea.String("AddHrmLegalEntity"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/masters/legalEntities/companies"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddHrmLegalEntityResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增法人公司
//
// @param request - AddHrmLegalEntityRequest
//
// @return AddHrmLegalEntityResponse
func (client *Client) AddHrmLegalEntity(request *AddHrmLegalEntityRequest) (_result *AddHrmLegalEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddHrmLegalEntityHeaders{}
	_result = &AddHrmLegalEntityResponse{}
	_body, _err := client.AddHrmLegalEntityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 智能人事添加待入职员工信息(支持花名册数据和分组明细更新)
//
// @param request - AddHrmPreentryRequest
//
// @param headers - AddHrmPreentryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddHrmPreentryResponse
func (client *Client) AddHrmPreentryWithOptions(request *AddHrmPreentryRequest, headers *AddHrmPreentryHeaders, runtime *util.RuntimeOptions) (_result *AddHrmPreentryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		body["agentId"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.Groups)) {
		body["groups"] = request.Groups
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		body["mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NeedSendPreEntryMsg)) {
		body["needSendPreEntryMsg"] = request.NeedSendPreEntryMsg
	}

	if !tea.BoolValue(util.IsUnset(request.PreEntryTime)) {
		body["preEntryTime"] = request.PreEntryTime
	}

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
		Action:      tea.String("AddHrmPreentry"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/preentries"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddHrmPreentryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能人事添加待入职员工信息(支持花名册数据和分组明细更新)
//
// @param request - AddHrmPreentryRequest
//
// @return AddHrmPreentryResponse
func (client *Client) AddHrmPreentry(request *AddHrmPreentryRequest) (_result *AddHrmPreentryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddHrmPreentryHeaders{}
	_result = &AddHrmPreentryResponse{}
	_body, _err := client.AddHrmPreentryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新建花名册字段分组
//
// @param request - AddRosterFieldFormRequest
//
// @param headers - AddRosterFieldFormHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRosterFieldFormResponse
func (client *Client) AddRosterFieldFormWithOptions(request *AddRosterFieldFormRequest, headers *AddRosterFieldFormHeaders, runtime *util.RuntimeOptions) (_result *AddRosterFieldFormResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Detail)) {
		body["detail"] = request.Detail
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

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
		Action:      tea.String("AddRosterFieldForm"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/rosterFieldForm/add"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddRosterFieldFormResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建花名册字段分组
//
// @param request - AddRosterFieldFormRequest
//
// @return AddRosterFieldFormResponse
func (client *Client) AddRosterFieldForm(request *AddRosterFieldFormRequest) (_result *AddRosterFieldFormResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddRosterFieldFormHeaders{}
	_result = &AddRosterFieldFormResponse{}
	_body, _err := client.AddRosterFieldFormWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建电子签签署记录
//
// @param request - CreateRecordRequest
//
// @param headers - CreateRecordHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRecordResponse
func (client *Client) CreateRecordWithOptions(request *CreateRecordRequest, headers *CreateRecordHeaders, runtime *util.RuntimeOptions) (_result *CreateRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttachmentList)) {
		body["attachmentList"] = request.AttachmentList
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.FieldList)) {
		body["fieldList"] = request.FieldList
	}

	if !tea.BoolValue(util.IsUnset(request.GroupList)) {
		body["groupList"] = request.GroupList
	}

	if !tea.BoolValue(util.IsUnset(request.OuterId)) {
		body["outerId"] = request.OuterId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.SignLastLegalEntityName)) {
		body["signLastLegalEntityName"] = request.SignLastLegalEntityName
	}

	if !tea.BoolValue(util.IsUnset(request.SignLegalEntityName)) {
		body["signLegalEntityName"] = request.SignLegalEntityName
	}

	if !tea.BoolValue(util.IsUnset(request.SignSource)) {
		body["signSource"] = request.SignSource
	}

	if !tea.BoolValue(util.IsUnset(request.SignStartUserId)) {
		body["signStartUserId"] = request.SignStartUserId
	}

	if !tea.BoolValue(util.IsUnset(request.SignUserId)) {
		body["signUserId"] = request.SignUserId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
	}

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
		Action:      tea.String("CreateRecord"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/masters/signCenters/records"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRecordResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建电子签签署记录
//
// @param request - CreateRecordRequest
//
// @return CreateRecordResponse
func (client *Client) CreateRecord(request *CreateRecordRequest) (_result *CreateRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateRecordHeaders{}
	_result = &CreateRecordResponse{}
	_body, _err := client.CreateRecordWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除花名册自定义字段
//
// @param request - DeleteCustomRosterFieldRequest
//
// @param headers - DeleteCustomRosterFieldHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomRosterFieldResponse
func (client *Client) DeleteCustomRosterFieldWithOptions(request *DeleteCustomRosterFieldRequest, headers *DeleteCustomRosterFieldHeaders, runtime *util.RuntimeOptions) (_result *DeleteCustomRosterFieldResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FieldCode)) {
		body["fieldCode"] = request.FieldCode
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["groupId"] = request.GroupId
	}

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
		Action:      tea.String("DeleteCustomRosterField"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/customRosterField/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCustomRosterFieldResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除花名册自定义字段
//
// @param request - DeleteCustomRosterFieldRequest
//
// @return DeleteCustomRosterFieldResponse
func (client *Client) DeleteCustomRosterField(request *DeleteCustomRosterFieldRequest) (_result *DeleteCustomRosterFieldResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteCustomRosterFieldHeaders{}
	_result = &DeleteCustomRosterFieldResponse{}
	_body, _err := client.DeleteCustomRosterFieldWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除花名册字段分组
//
// @param request - DeleteRosterFieldFormRequest
//
// @param headers - DeleteRosterFieldFormHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRosterFieldFormResponse
func (client *Client) DeleteRosterFieldFormWithOptions(request *DeleteRosterFieldFormRequest, headers *DeleteRosterFieldFormHeaders, runtime *util.RuntimeOptions) (_result *DeleteRosterFieldFormResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FormId)) {
		body["formId"] = request.FormId
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
		Action:      tea.String("DeleteRosterFieldForm"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/rosterFieldForm/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRosterFieldFormResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除花名册字段分组
//
// @param request - DeleteRosterFieldFormRequest
//
// @return DeleteRosterFieldFormResponse
func (client *Client) DeleteRosterFieldForm(request *DeleteRosterFieldFormRequest) (_result *DeleteRosterFieldFormResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteRosterFieldFormHeaders{}
	_result = &DeleteRosterFieldFormResponse{}
	_body, _err := client.DeleteRosterFieldFormWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 智能人事设备市场管理
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeviceMarketManagerResponse
func (client *Client) DeviceMarketManagerWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeviceMarketManagerResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeviceMarketManager"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/device/market/manager"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeviceMarketManagerResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能人事设备市场管理
//
// @return DeviceMarketManagerResponse
func (client *Client) DeviceMarketManager() (_result *DeviceMarketManagerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeviceMarketManagerResponse{}
	_body, _err := client.DeviceMarketManagerWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 智能人事设备定向管理接口
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeviceMarketOrderManagerResponse
func (client *Client) DeviceMarketOrderManagerWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeviceMarketOrderManagerResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeviceMarketOrderManager"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/device/market/order/manager"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeviceMarketOrderManagerResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能人事设备定向管理接口
//
// @return DeviceMarketOrderManagerResponse
func (client *Client) DeviceMarketOrderManager() (_result *DeviceMarketOrderManagerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeviceMarketOrderManagerResponse{}
	_body, _err := client.DeviceMarketOrderManagerWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// e签宝专有查询证件接口
//
// @param request - ECertQueryRequest
//
// @param headers - ECertQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ECertQueryResponse
func (client *Client) ECertQueryWithOptions(request *ECertQueryRequest, headers *ECertQueryHeaders, runtime *util.RuntimeOptions) (_result *ECertQueryResponse, _err error) {
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
		Action:      tea.String("ECertQuery"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/eCerts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ECertQueryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// e签宝专有查询证件接口
//
// @param request - ECertQueryRequest
//
// @return ECertQueryResponse
func (client *Client) ECertQuery(request *ECertQueryRequest) (_result *ECertQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ECertQueryHeaders{}
	_result = &ECertQueryResponse{}
	_body, _err := client.ECertQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 加入待离职
//
// @param request - EmpStartDismissionRequest
//
// @param headers - EmpStartDismissionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EmpStartDismissionResponse
func (client *Client) EmpStartDismissionWithOptions(request *EmpStartDismissionRequest, headers *EmpStartDismissionHeaders, runtime *util.RuntimeOptions) (_result *EmpStartDismissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LastWorkDate)) {
		body["lastWorkDate"] = request.LastWorkDate
	}

	if !tea.BoolValue(util.IsUnset(request.Partner)) {
		body["partner"] = request.Partner
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.TerminationReasonPassive)) {
		body["terminationReasonPassive"] = request.TerminationReasonPassive
	}

	if !tea.BoolValue(util.IsUnset(request.TerminationReasonVoluntary)) {
		body["terminationReasonVoluntary"] = request.TerminationReasonVoluntary
	}

	if !tea.BoolValue(util.IsUnset(request.ToHireBlackList)) {
		body["toHireBlackList"] = request.ToHireBlackList
	}

	if !tea.BoolValue(util.IsUnset(request.ToHireDismissionTalent)) {
		body["toHireDismissionTalent"] = request.ToHireDismissionTalent
	}

	if !tea.BoolValue(util.IsUnset(request.ToHrmBlackList)) {
		body["toHrmBlackList"] = request.ToHrmBlackList
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
		Action:      tea.String("EmpStartDismission"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/pendingDismission/start"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &EmpStartDismissionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 加入待离职
//
// @param request - EmpStartDismissionRequest
//
// @return EmpStartDismissionResponse
func (client *Client) EmpStartDismission(request *EmpStartDismissionRequest) (_result *EmpStartDismissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EmpStartDismissionHeaders{}
	_result = &EmpStartDismissionResponse{}
	_body, _err := client.EmpStartDismissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 智能人事员工档案附件更新
//
// @param request - EmployeeAttachmentUpdateRequest
//
// @param headers - EmployeeAttachmentUpdateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EmployeeAttachmentUpdateResponse
func (client *Client) EmployeeAttachmentUpdateWithOptions(request *EmployeeAttachmentUpdateRequest, headers *EmployeeAttachmentUpdateHeaders, runtime *util.RuntimeOptions) (_result *EmployeeAttachmentUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppAgentId)) {
		query["appAgentId"] = request.AppAgentId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FieldCode)) {
		body["fieldCode"] = request.FieldCode
	}

	if !tea.BoolValue(util.IsUnset(request.FileSuffix)) {
		body["fileSuffix"] = request.FileSuffix
	}

	if !tea.BoolValue(util.IsUnset(request.MediaId)) {
		body["mediaId"] = request.MediaId
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
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EmployeeAttachmentUpdate"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/employees/attachments"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &EmployeeAttachmentUpdateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能人事员工档案附件更新
//
// @param request - EmployeeAttachmentUpdateRequest
//
// @return EmployeeAttachmentUpdateResponse
func (client *Client) EmployeeAttachmentUpdate(request *EmployeeAttachmentUpdateRequest) (_result *EmployeeAttachmentUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EmployeeAttachmentUpdateHeaders{}
	_result = &EmployeeAttachmentUpdateResponse{}
	_body, _err := client.EmployeeAttachmentUpdateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 人事高级合同管理回退
//
// @param request - EsignRollbackRequest
//
// @param headers - EsignRollbackHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EsignRollbackResponse
func (client *Client) EsignRollbackWithOptions(request *EsignRollbackRequest, headers *EsignRollbackHeaders, runtime *util.RuntimeOptions) (_result *EsignRollbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OptUserId)) {
		query["optUserId"] = request.OptUserId
	}

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
		Action:      tea.String("EsignRollback"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/contracts/esign/rollback"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &EsignRollbackResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人事高级合同管理回退
//
// @param request - EsignRollbackRequest
//
// @return EsignRollbackResponse
func (client *Client) EsignRollback(request *EsignRollbackRequest) (_result *EsignRollbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EsignRollbackHeaders{}
	_result = &EsignRollbackResponse{}
	_body, _err := client.EsignRollbackWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取所有离职原因
//
// @param headers - GetAllDismissionReasonsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAllDismissionReasonsResponse
func (client *Client) GetAllDismissionReasonsWithOptions(headers *GetAllDismissionReasonsHeaders, runtime *util.RuntimeOptions) (_result *GetAllDismissionReasonsResponse, _err error) {
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
		Action:      tea.String("GetAllDismissionReasons"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/dismission/reasons"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAllDismissionReasonsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取所有离职原因
//
// @return GetAllDismissionReasonsResponse
func (client *Client) GetAllDismissionReasons() (_result *GetAllDismissionReasonsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAllDismissionReasonsHeaders{}
	_result = &GetAllDismissionReasonsResponse{}
	_body, _err := client.GetAllDismissionReasonsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取员工花名册指定字段的信息，支持明细分组字段
//
// @param request - GetEmployeeRosterByFieldRequest
//
// @param headers - GetEmployeeRosterByFieldHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEmployeeRosterByFieldResponse
func (client *Client) GetEmployeeRosterByFieldWithOptions(request *GetEmployeeRosterByFieldRequest, headers *GetEmployeeRosterByFieldHeaders, runtime *util.RuntimeOptions) (_result *GetEmployeeRosterByFieldResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppAgentId)) {
		body["appAgentId"] = request.AppAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.FieldFilterList)) {
		body["fieldFilterList"] = request.FieldFilterList
	}

	if !tea.BoolValue(util.IsUnset(request.Text2SelectConvert)) {
		body["text2SelectConvert"] = request.Text2SelectConvert
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
		Action:      tea.String("GetEmployeeRosterByField"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/rosters/lists/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEmployeeRosterByFieldResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取员工花名册指定字段的信息，支持明细分组字段
//
// @param request - GetEmployeeRosterByFieldRequest
//
// @return GetEmployeeRosterByFieldResponse
func (client *Client) GetEmployeeRosterByField(request *GetEmployeeRosterByFieldRequest) (_result *GetEmployeeRosterByFieldResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetEmployeeRosterByFieldHeaders{}
	_result = &GetEmployeeRosterByFieldResponse{}
	_body, _err := client.GetEmployeeRosterByFieldWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询文件模板列表及文件模板内花名册字段
//
// @param request - GetFileTemplateListRequest
//
// @param headers - GetFileTemplateListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileTemplateListResponse
func (client *Client) GetFileTemplateListWithOptions(request *GetFileTemplateListRequest, headers *GetFileTemplateListHeaders, runtime *util.RuntimeOptions) (_result *GetFileTemplateListResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SignSource)) {
		body["signSource"] = request.SignSource
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateStatus)) {
		body["templateStatus"] = request.TemplateStatus
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateTypeList)) {
		body["templateTypeList"] = request.TemplateTypeList
	}

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
		Action:      tea.String("GetFileTemplateList"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/masters/fileTemplates/lists/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFileTemplateListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询文件模板列表及文件模板内花名册字段
//
// @param request - GetFileTemplateListRequest
//
// @return GetFileTemplateListResponse
func (client *Client) GetFileTemplateList(request *GetFileTemplateListRequest) (_result *GetFileTemplateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFileTemplateListHeaders{}
	_result = &GetFileTemplateListResponse{}
	_body, _err := client.GetFileTemplateListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过签署记录id查询指定的电子签署记录
//
// @param request - GetSignRecordByIdRequest
//
// @param headers - GetSignRecordByIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSignRecordByIdResponse
func (client *Client) GetSignRecordByIdWithOptions(request *GetSignRecordByIdRequest, headers *GetSignRecordByIdHeaders, runtime *util.RuntimeOptions) (_result *GetSignRecordByIdResponse, _err error) {
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
		Action:      tea.String("GetSignRecordById"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/masters/signCenters/records/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSignRecordByIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过签署记录id查询指定的电子签署记录
//
// @param request - GetSignRecordByIdRequest
//
// @return GetSignRecordByIdResponse
func (client *Client) GetSignRecordById(request *GetSignRecordByIdRequest) (_result *GetSignRecordByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSignRecordByIdHeaders{}
	_result = &GetSignRecordByIdResponse{}
	_body, _err := client.GetSignRecordByIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询指定用户的电子签署记录，并返回签署记录的基本数据及已签署完成的文件预览地址
//
// @param request - GetSignRecordByUserIdRequest
//
// @param headers - GetSignRecordByUserIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSignRecordByUserIdResponse
func (client *Client) GetSignRecordByUserIdWithOptions(request *GetSignRecordByUserIdRequest, headers *GetSignRecordByUserIdHeaders, runtime *util.RuntimeOptions) (_result *GetSignRecordByUserIdResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SignStatus)) {
		body["signStatus"] = request.SignStatus
	}

	if !tea.BoolValue(util.IsUnset(request.SignUserId)) {
		body["signUserId"] = request.SignUserId
	}

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
		Action:      tea.String("GetSignRecordByUserId"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/masters/signCenters/users/records/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSignRecordByUserIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定用户的电子签署记录，并返回签署记录的基本数据及已签署完成的文件预览地址
//
// @param request - GetSignRecordByUserIdRequest
//
// @return GetSignRecordByUserIdResponse
func (client *Client) GetSignRecordByUserId(request *GetSignRecordByUserIdRequest) (_result *GetSignRecordByUserIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSignRecordByUserIdHeaders{}
	_result = &GetSignRecordByUserIdResponse{}
	_body, _err := client.GetSignRecordByUserIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询指定outerId的电子签署记录详情
//
// @param request - GetUserSignedRecordsByOuterIdRequest
//
// @param headers - GetUserSignedRecordsByOuterIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserSignedRecordsByOuterIdResponse
func (client *Client) GetUserSignedRecordsByOuterIdWithOptions(request *GetUserSignedRecordsByOuterIdRequest, headers *GetUserSignedRecordsByOuterIdHeaders, runtime *util.RuntimeOptions) (_result *GetUserSignedRecordsByOuterIdResponse, _err error) {
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
		Action:      tea.String("GetUserSignedRecordsByOuterId"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/masters/signCenters/outerIds/records/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserSignedRecordsByOuterIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定outerId的电子签署记录详情
//
// @param request - GetUserSignedRecordsByOuterIdRequest
//
// @return GetUserSignedRecordsByOuterIdResponse
func (client *Client) GetUserSignedRecordsByOuterId(request *GetUserSignedRecordsByOuterIdRequest) (_result *GetUserSignedRecordsByOuterIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserSignedRecordsByOuterIdHeaders{}
	_result = &GetUserSignedRecordsByOuterIdResponse{}
	_body, _err := client.GetUserSignedRecordsByOuterIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 智能人事权限查询
//
// @param request - HrmAuthResourcesQueryRequest
//
// @param headers - HrmAuthResourcesQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrmAuthResourcesQueryResponse
func (client *Client) HrmAuthResourcesQueryWithOptions(request *HrmAuthResourcesQueryRequest, headers *HrmAuthResourcesQueryHeaders, runtime *util.RuntimeOptions) (_result *HrmAuthResourcesQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthResourceIds)) {
		body["authResourceIds"] = request.AuthResourceIds
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
		Action:      tea.String("HrmAuthResourcesQuery"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/authResources/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrmAuthResourcesQueryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能人事权限查询
//
// @param request - HrmAuthResourcesQueryRequest
//
// @return HrmAuthResourcesQueryResponse
func (client *Client) HrmAuthResourcesQuery(request *HrmAuthResourcesQueryRequest) (_result *HrmAuthResourcesQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrmAuthResourcesQueryHeaders{}
	_result = &HrmAuthResourcesQueryResponse{}
	_body, _err := client.HrmAuthResourcesQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 智能人事权益查询
//
// @param request - HrmBenefitQueryRequest
//
// @param headers - HrmBenefitQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrmBenefitQueryResponse
func (client *Client) HrmBenefitQueryWithOptions(request *HrmBenefitQueryRequest, headers *HrmBenefitQueryHeaders, runtime *util.RuntimeOptions) (_result *HrmBenefitQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BenefitCodes)) {
		body["benefitCodes"] = request.BenefitCodes
	}

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
		Action:      tea.String("HrmBenefitQuery"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/benefits/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrmBenefitQueryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能人事权益查询
//
// @param request - HrmBenefitQueryRequest
//
// @return HrmBenefitQueryResponse
func (client *Client) HrmBenefitQuery(request *HrmBenefitQueryRequest) (_result *HrmBenefitQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrmBenefitQueryHeaders{}
	_result = &HrmBenefitQueryResponse{}
	_body, _err := client.HrmBenefitQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询企业配置信息
//
// @param request - HrmCorpConfigQueryRequest
//
// @param headers - HrmCorpConfigQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrmCorpConfigQueryResponse
func (client *Client) HrmCorpConfigQueryWithOptions(request *HrmCorpConfigQueryRequest, headers *HrmCorpConfigQueryHeaders, runtime *util.RuntimeOptions) (_result *HrmCorpConfigQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubType)) {
		body["subType"] = request.SubType
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
		Action:      tea.String("HrmCorpConfigQuery"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/corp/configs/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrmCorpConfigQueryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询企业配置信息
//
// @param request - HrmCorpConfigQueryRequest
//
// @return HrmCorpConfigQueryResponse
func (client *Client) HrmCorpConfigQuery(request *HrmCorpConfigQueryRequest) (_result *HrmCorpConfigQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrmCorpConfigQueryHeaders{}
	_result = &HrmCorpConfigQueryResponse{}
	_body, _err := client.HrmCorpConfigQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 智能人事邮件发送
//
// @param request - HrmMailSendRequest
//
// @param headers - HrmMailSendHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrmMailSendResponse
func (client *Client) HrmMailSendWithOptions(request *HrmMailSendRequest, headers *HrmMailSendHeaders, runtime *util.RuntimeOptions) (_result *HrmMailSendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Mail)) {
		body["mail"] = request.Mail
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
		Action:      tea.String("HrmMailSend"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/mails/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrmMailSendResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能人事邮件发送
//
// @param request - HrmMailSendRequest
//
// @return HrmMailSendResponse
func (client *Client) HrmMailSend(request *HrmMailSendRequest) (_result *HrmMailSendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrmMailSendHeaders{}
	_result = &HrmMailSendResponse{}
	_body, _err := client.HrmMailSendWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 人事2.0支持Moka事件转发
//
// @param request - HrmMokaEventRequest
//
// @param headers - HrmMokaEventHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrmMokaEventResponse
func (client *Client) HrmMokaEventWithOptions(request *HrmMokaEventRequest, headers *HrmMokaEventHeaders, runtime *util.RuntimeOptions) (_result *HrmMokaEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

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
		Action:      tea.String("HrmMokaEvent"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/moka/events/forward"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrmMokaEventResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人事2.0支持Moka事件转发
//
// @param request - HrmMokaEventRequest
//
// @return HrmMokaEventResponse
func (client *Client) HrmMokaEvent(request *HrmMokaEventRequest) (_result *HrmMokaEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrmMokaEventHeaders{}
	_result = &HrmMokaEventResponse{}
	_body, _err := client.HrmMokaEventWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 人事2.0支持Moka接口转发
//
// @param request - HrmMokaOapiRequest
//
// @param headers - HrmMokaOapiHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrmMokaOapiResponse
func (client *Client) HrmMokaOapiWithOptions(request *HrmMokaOapiRequest, headers *HrmMokaOapiHeaders, runtime *util.RuntimeOptions) (_result *HrmMokaOapiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiCode)) {
		body["apiCode"] = request.ApiCode
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
		Action:      tea.String("HrmMokaOapi"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/moka/forward"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrmMokaOapiResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人事2.0支持Moka接口转发
//
// @param request - HrmMokaOapiRequest
//
// @return HrmMokaOapiResponse
func (client *Client) HrmMokaOapi(request *HrmMokaOapiRequest) (_result *HrmMokaOapiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrmMokaOapiHeaders{}
	_result = &HrmMokaOapiResponse{}
	_body, _err := client.HrmMokaOapiWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 智能人事转正接口
//
// @param request - HrmProcessRegularRequest
//
// @param headers - HrmProcessRegularHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrmProcessRegularResponse
func (client *Client) HrmProcessRegularWithOptions(request *HrmProcessRegularRequest, headers *HrmProcessRegularHeaders, runtime *util.RuntimeOptions) (_result *HrmProcessRegularResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperationId)) {
		body["operationId"] = request.OperationId
	}

	if !tea.BoolValue(util.IsUnset(request.RegularDate)) {
		body["regularDate"] = request.RegularDate
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
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
		Action:      tea.String("HrmProcessRegular"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/processes/regulars/become"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrmProcessRegularResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能人事转正接口
//
// @param request - HrmProcessRegularRequest
//
// @return HrmProcessRegularResponse
func (client *Client) HrmProcessRegular(request *HrmProcessRegularRequest) (_result *HrmProcessRegularResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrmProcessRegularHeaders{}
	_result = &HrmProcessRegularResponse{}
	_body, _err := client.HrmProcessRegularWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 智能人事离职和交接接口
//
// @param request - HrmProcessTerminationAndHandoverRequest
//
// @param headers - HrmProcessTerminationAndHandoverHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrmProcessTerminationAndHandoverResponse
func (client *Client) HrmProcessTerminationAndHandoverWithOptions(request *HrmProcessTerminationAndHandoverRequest, headers *HrmProcessTerminationAndHandoverHeaders, runtime *util.RuntimeOptions) (_result *HrmProcessTerminationAndHandoverResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AflowHandOverUserId)) {
		body["aflowHandOverUserId"] = request.AflowHandOverUserId
	}

	if !tea.BoolValue(util.IsUnset(request.DingPanHandoverUserId)) {
		body["dingPanHandoverUserId"] = request.DingPanHandoverUserId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectSubordinatesHandoverUserId)) {
		body["directSubordinatesHandoverUserId"] = request.DirectSubordinatesHandoverUserId
	}

	if !tea.BoolValue(util.IsUnset(request.DismissionMemo)) {
		body["dismissionMemo"] = request.DismissionMemo
	}

	if !tea.BoolValue(util.IsUnset(request.DismissionReason)) {
		body["dismissionReason"] = request.DismissionReason
	}

	if !tea.BoolValue(util.IsUnset(request.DocNoteHandoverUserId)) {
		body["docNoteHandoverUserId"] = request.DocNoteHandoverUserId
	}

	if !tea.BoolValue(util.IsUnset(request.LastWorkDate)) {
		body["lastWorkDate"] = request.LastWorkDate
	}

	if !tea.BoolValue(util.IsUnset(request.OptUserId)) {
		body["optUserId"] = request.OptUserId
	}

	if !tea.BoolValue(util.IsUnset(request.PermissionHandoverUserId)) {
		body["permissionHandoverUserId"] = request.PermissionHandoverUserId
	}

	if !tea.BoolValue(util.IsUnset(request.TerminationReasonPassive)) {
		body["terminationReasonPassive"] = request.TerminationReasonPassive
	}

	if !tea.BoolValue(util.IsUnset(request.TerminationReasonVoluntary)) {
		body["terminationReasonVoluntary"] = request.TerminationReasonVoluntary
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
		Action:      tea.String("HrmProcessTerminationAndHandover"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/processes/terminateAndHandOver"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrmProcessTerminationAndHandoverResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能人事离职和交接接口
//
// @param request - HrmProcessTerminationAndHandoverRequest
//
// @return HrmProcessTerminationAndHandoverResponse
func (client *Client) HrmProcessTerminationAndHandover(request *HrmProcessTerminationAndHandoverRequest) (_result *HrmProcessTerminationAndHandoverResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrmProcessTerminationAndHandoverHeaders{}
	_result = &HrmProcessTerminationAndHandoverResponse{}
	_body, _err := client.HrmProcessTerminationAndHandoverWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 智能人事调岗接口
//
// @param request - HrmProcessTransferRequest
//
// @param headers - HrmProcessTransferHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrmProcessTransferResponse
func (client *Client) HrmProcessTransferWithOptions(request *HrmProcessTransferRequest, headers *HrmProcessTransferHeaders, runtime *util.RuntimeOptions) (_result *HrmProcessTransferResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptIdsAfterTransfer)) {
		body["deptIdsAfterTransfer"] = request.DeptIdsAfterTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.JobIdAfterTransfer)) {
		body["jobIdAfterTransfer"] = request.JobIdAfterTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.MainDeptIdAfterTransfer)) {
		body["mainDeptIdAfterTransfer"] = request.MainDeptIdAfterTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.OperateUserId)) {
		body["operateUserId"] = request.OperateUserId
	}

	if !tea.BoolValue(util.IsUnset(request.PositionIdAfterTransfer)) {
		body["positionIdAfterTransfer"] = request.PositionIdAfterTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.PositionLevelAfterTransfer)) {
		body["positionLevelAfterTransfer"] = request.PositionLevelAfterTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.PositionNameAfterTransfer)) {
		body["positionNameAfterTransfer"] = request.PositionNameAfterTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.RankIdAfterTransfer)) {
		body["rankIdAfterTransfer"] = request.RankIdAfterTransfer
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
		Action:      tea.String("HrmProcessTransfer"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/processes/transfer"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrmProcessTransferResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能人事调岗接口
//
// @param request - HrmProcessTransferRequest
//
// @return HrmProcessTransferResponse
func (client *Client) HrmProcessTransfer(request *HrmProcessTransferRequest) (_result *HrmProcessTransferResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrmProcessTransferHeaders{}
	_result = &HrmProcessTransferResponse{}
	_body, _err := client.HrmProcessTransferWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改员工最后一次离职信息
//
// @param request - HrmProcessUpdateTerminationInfoRequest
//
// @param headers - HrmProcessUpdateTerminationInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrmProcessUpdateTerminationInfoResponse
func (client *Client) HrmProcessUpdateTerminationInfoWithOptions(request *HrmProcessUpdateTerminationInfoRequest, headers *HrmProcessUpdateTerminationInfoHeaders, runtime *util.RuntimeOptions) (_result *HrmProcessUpdateTerminationInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DismissionMemo)) {
		body["dismissionMemo"] = request.DismissionMemo
	}

	if !tea.BoolValue(util.IsUnset(request.LastWorkDate)) {
		body["lastWorkDate"] = request.LastWorkDate
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
		Action:      tea.String("HrmProcessUpdateTerminationInfo"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/processes/employees/terminations"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrmProcessUpdateTerminationInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改员工最后一次离职信息
//
// @param request - HrmProcessUpdateTerminationInfoRequest
//
// @return HrmProcessUpdateTerminationInfoResponse
func (client *Client) HrmProcessUpdateTerminationInfo(request *HrmProcessUpdateTerminationInfoRequest) (_result *HrmProcessUpdateTerminationInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrmProcessUpdateTerminationInfoHeaders{}
	_result = &HrmProcessUpdateTerminationInfoResponse{}
	_body, _err := client.HrmProcessUpdateTerminationInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 智能人事pts能力调用
//
// @param request - HrmPtsServiceRequest
//
// @param headers - HrmPtsServiceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HrmPtsServiceResponse
func (client *Client) HrmPtsServiceWithOptions(request *HrmPtsServiceRequest, headers *HrmPtsServiceHeaders, runtime *util.RuntimeOptions) (_result *HrmPtsServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Env)) {
		body["env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.Method)) {
		body["method"] = request.Method
	}

	if !tea.BoolValue(util.IsUnset(request.OuterId)) {
		body["outerId"] = request.OuterId
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		body["path"] = request.Path
	}

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
		Action:      tea.String("HrmPtsService"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/pts/request"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HrmPtsServiceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能人事pts能力调用
//
// @param request - HrmPtsServiceRequest
//
// @return HrmPtsServiceResponse
func (client *Client) HrmPtsService(request *HrmPtsServiceRequest) (_result *HrmPtsServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HrmPtsServiceHeaders{}
	_result = &HrmPtsServiceResponse{}
	_body, _err := client.HrmPtsServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 作废签署记录
//
// @param request - InvalidSignRecordsRequest
//
// @param headers - InvalidSignRecordsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvalidSignRecordsResponse
func (client *Client) InvalidSignRecordsWithOptions(request *InvalidSignRecordsRequest, headers *InvalidSignRecordsHeaders, runtime *util.RuntimeOptions) (_result *InvalidSignRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InvalidUserId)) {
		body["invalidUserId"] = request.InvalidUserId
	}

	if !tea.BoolValue(util.IsUnset(request.SignRecordIds)) {
		body["signRecordIds"] = request.SignRecordIds
	}

	if !tea.BoolValue(util.IsUnset(request.StatusRemark)) {
		body["statusRemark"] = request.StatusRemark
	}

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
		Action:      tea.String("InvalidSignRecords"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/masters/signCenters/records/invalid"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &InvalidSignRecordsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 作废签署记录
//
// @param request - InvalidSignRecordsRequest
//
// @return InvalidSignRecordsResponse
func (client *Client) InvalidSignRecords(request *InvalidSignRecordsRequest) (_result *InvalidSignRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InvalidSignRecordsHeaders{}
	_result = &InvalidSignRecordsResponse{}
	_body, _err := client.InvalidSignRecordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 智能人事主数据删除服务
//
// @param request - MasterDataDeleteRequest
//
// @param headers - MasterDataDeleteHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MasterDataDeleteResponse
func (client *Client) MasterDataDeleteWithOptions(request *MasterDataDeleteRequest, headers *MasterDataDeleteHeaders, runtime *util.RuntimeOptions) (_result *MasterDataDeleteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

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
		Action:      tea.String("MasterDataDelete"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/masters/datas/batchRemove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &MasterDataDeleteResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能人事主数据删除服务
//
// @param request - MasterDataDeleteRequest
//
// @return MasterDataDeleteResponse
func (client *Client) MasterDataDelete(request *MasterDataDeleteRequest) (_result *MasterDataDeleteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MasterDataDeleteHeaders{}
	_result = &MasterDataDeleteResponse{}
	_body, _err := client.MasterDataDeleteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 智能人事主数据查询服务
//
// @param request - MasterDataQueryRequest
//
// @param headers - MasterDataQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MasterDataQueryResponse
func (client *Client) MasterDataQueryWithOptions(request *MasterDataQueryRequest, headers *MasterDataQueryHeaders, runtime *util.RuntimeOptions) (_result *MasterDataQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizUK)) {
		body["bizUK"] = request.BizUK
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OptUserId)) {
		body["optUserId"] = request.OptUserId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryParams)) {
		body["queryParams"] = request.QueryParams
	}

	if !tea.BoolValue(util.IsUnset(request.RelationIds)) {
		body["relationIds"] = request.RelationIds
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeCode)) {
		body["scopeCode"] = request.ScopeCode
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["tenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.ViewEntityCode)) {
		body["viewEntityCode"] = request.ViewEntityCode
	}

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
		Action:      tea.String("MasterDataQuery"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/masters/datas/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &MasterDataQueryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能人事主数据查询服务
//
// @param request - MasterDataQueryRequest
//
// @return MasterDataQueryResponse
func (client *Client) MasterDataQuery(request *MasterDataQueryRequest) (_result *MasterDataQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MasterDataQueryHeaders{}
	_result = &MasterDataQueryResponse{}
	_body, _err := client.MasterDataQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 智能人事主数据保存服务
//
// @param request - MasterDataSaveRequest
//
// @param headers - MasterDataSaveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MasterDataSaveResponse
func (client *Client) MasterDataSaveWithOptions(request *MasterDataSaveRequest, headers *MasterDataSaveHeaders, runtime *util.RuntimeOptions) (_result *MasterDataSaveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

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
		Action:      tea.String("MasterDataSave"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/masters/datas/save"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &MasterDataSaveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能人事主数据保存服务
//
// @param request - MasterDataSaveRequest
//
// @return MasterDataSaveResponse
func (client *Client) MasterDataSave(request *MasterDataSaveRequest) (_result *MasterDataSaveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MasterDataSaveHeaders{}
	_result = &MasterDataSaveResponse{}
	_body, _err := client.MasterDataSaveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 主数据中拥有某个领域数据的租户信息查询
//
// @param request - MasterDataTenantQueyRequest
//
// @param headers - MasterDataTenantQueyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MasterDataTenantQueyResponse
func (client *Client) MasterDataTenantQueyWithOptions(request *MasterDataTenantQueyRequest, headers *MasterDataTenantQueyHeaders, runtime *util.RuntimeOptions) (_result *MasterDataTenantQueyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityCode)) {
		query["entityCode"] = request.EntityCode
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeCode)) {
		query["scopeCode"] = request.ScopeCode
	}

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
		Action:      tea.String("MasterDataTenantQuey"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/masters/tenants"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &MasterDataTenantQueyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 主数据中拥有某个领域数据的租户信息查询
//
// @param request - MasterDataTenantQueyRequest
//
// @return MasterDataTenantQueyResponse
func (client *Client) MasterDataTenantQuey(request *MasterDataTenantQueyRequest) (_result *MasterDataTenantQueyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MasterDataTenantQueyHeaders{}
	_result = &MasterDataTenantQueyResponse{}
	_body, _err := client.MasterDataTenantQueyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 智能人事主数据根据ID获取
//
// @param request - MasterDatasGetRequest
//
// @param headers - MasterDatasGetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MasterDatasGetResponse
func (client *Client) MasterDatasGetWithOptions(request *MasterDatasGetRequest, headers *MasterDatasGetHeaders, runtime *util.RuntimeOptions) (_result *MasterDatasGetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ObjId)) {
		body["objId"] = request.ObjId
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeCode)) {
		body["scopeCode"] = request.ScopeCode
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["tenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.ViewEntityCode)) {
		body["viewEntityCode"] = request.ViewEntityCode
	}

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
		Action:      tea.String("MasterDatasGet"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/masterDatas/objects/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &MasterDatasGetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能人事主数据根据ID获取
//
// @param request - MasterDatasGetRequest
//
// @return MasterDatasGetResponse
func (client *Client) MasterDatasGet(request *MasterDatasGetRequest) (_result *MasterDatasGetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MasterDatasGetHeaders{}
	_result = &MasterDatasGetResponse{}
	_body, _err := client.MasterDatasGetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 智能人事主数据查询服务
//
// @param request - MasterDatasQueryRequest
//
// @param headers - MasterDatasQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MasterDatasQueryResponse
func (client *Client) MasterDatasQueryWithOptions(request *MasterDatasQueryRequest, headers *MasterDatasQueryHeaders, runtime *util.RuntimeOptions) (_result *MasterDatasQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizUK)) {
		body["bizUK"] = request.BizUK
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.QueryParams)) {
		body["queryParams"] = request.QueryParams
	}

	if !tea.BoolValue(util.IsUnset(request.RelationIds)) {
		body["relationIds"] = request.RelationIds
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeCode)) {
		body["scopeCode"] = request.ScopeCode
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["tenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.ViewEntityCode)) {
		body["viewEntityCode"] = request.ViewEntityCode
	}

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
		Action:      tea.String("MasterDatasQuery"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/masterDatas/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &MasterDatasQueryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能人事主数据查询服务
//
// @param request - MasterDatasQueryRequest
//
// @return MasterDatasQueryResponse
func (client *Client) MasterDatasQuery(request *MasterDatasQueryRequest) (_result *MasterDatasQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MasterDatasQueryHeaders{}
	_result = &MasterDatasQueryResponse{}
	_body, _err := client.MasterDatasQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// oem 老用户数据迁移时，开通oem 应用
//
// @param request - OpenOemMicroAppRequest
//
// @param headers - OpenOemMicroAppHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenOemMicroAppResponse
func (client *Client) OpenOemMicroAppWithOptions(request *OpenOemMicroAppRequest, headers *OpenOemMicroAppHeaders, runtime *util.RuntimeOptions) (_result *OpenOemMicroAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

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
		Action:      tea.String("OpenOemMicroApp"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/oem/microApps/open"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenOemMicroAppResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// oem 老用户数据迁移时，开通oem 应用
//
// @param request - OpenOemMicroAppRequest
//
// @return OpenOemMicroAppResponse
func (client *Client) OpenOemMicroApp(request *OpenOemMicroAppRequest) (_result *OpenOemMicroAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OpenOemMicroAppHeaders{}
	_result = &OpenOemMicroAppResponse{}
	_body, _err := client.OpenOemMicroAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 自定义入职流程数据查询
//
// @param request - QueryCustomEntryProcessesRequest
//
// @param headers - QueryCustomEntryProcessesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCustomEntryProcessesResponse
func (client *Client) QueryCustomEntryProcessesWithOptions(request *QueryCustomEntryProcessesRequest, headers *QueryCustomEntryProcessesHeaders, runtime *util.RuntimeOptions) (_result *QueryCustomEntryProcessesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.OperateUserId)) {
		query["operateUserId"] = request.OperateUserId
	}

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
		Action:      tea.String("QueryCustomEntryProcesses"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/customEntryProcesses"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCustomEntryProcessesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 自定义入职流程数据查询
//
// @param request - QueryCustomEntryProcessesRequest
//
// @return QueryCustomEntryProcessesResponse
func (client *Client) QueryCustomEntryProcesses(request *QueryCustomEntryProcessesRequest) (_result *QueryCustomEntryProcessesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCustomEntryProcessesHeaders{}
	_result = &QueryCustomEntryProcessesResponse{}
	_body, _err := client.QueryCustomEntryProcessesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询企业已离职员工列表
//
// @param request - QueryDismissionStaffIdListRequest
//
// @param headers - QueryDismissionStaffIdListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDismissionStaffIdListResponse
func (client *Client) QueryDismissionStaffIdListWithOptions(request *QueryDismissionStaffIdListRequest, headers *QueryDismissionStaffIdListHeaders, runtime *util.RuntimeOptions) (_result *QueryDismissionStaffIdListResponse, _err error) {
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
		Action:      tea.String("QueryDismissionStaffIdList"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/employees/dismissions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDismissionStaffIdListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询企业已离职员工列表
//
// @param request - QueryDismissionStaffIdListRequest
//
// @return QueryDismissionStaffIdListResponse
func (client *Client) QueryDismissionStaffIdList(request *QueryDismissionStaffIdListRequest) (_result *QueryDismissionStaffIdListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDismissionStaffIdListHeaders{}
	_result = &QueryDismissionStaffIdListResponse{}
	_body, _err := client.QueryDismissionStaffIdListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据传入的staffId列表，批量查询员工的离职信息
//
// @param tmpReq - QueryHrmEmployeeDismissionInfoRequest
//
// @param headers - QueryHrmEmployeeDismissionInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryHrmEmployeeDismissionInfoResponse
func (client *Client) QueryHrmEmployeeDismissionInfoWithOptions(tmpReq *QueryHrmEmployeeDismissionInfoRequest, headers *QueryHrmEmployeeDismissionInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryHrmEmployeeDismissionInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryHrmEmployeeDismissionInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UserIdList)) {
		request.UserIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIdList, tea.String("userIdList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserIdListShrink)) {
		query["userIdList"] = request.UserIdListShrink
	}

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
		Action:      tea.String("QueryHrmEmployeeDismissionInfo"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/employees/dimissionInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryHrmEmployeeDismissionInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据传入的staffId列表，批量查询员工的离职信息
//
// @param request - QueryHrmEmployeeDismissionInfoRequest
//
// @return QueryHrmEmployeeDismissionInfoResponse
func (client *Client) QueryHrmEmployeeDismissionInfo(request *QueryHrmEmployeeDismissionInfoRequest) (_result *QueryHrmEmployeeDismissionInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryHrmEmployeeDismissionInfoHeaders{}
	_result = &QueryHrmEmployeeDismissionInfoResponse{}
	_body, _err := client.QueryHrmEmployeeDismissionInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询企业的职级信息
//
// @param request - QueryJobRanksRequest
//
// @param headers - QueryJobRanksHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryJobRanksResponse
func (client *Client) QueryJobRanksWithOptions(request *QueryJobRanksRequest, headers *QueryJobRanksHeaders, runtime *util.RuntimeOptions) (_result *QueryJobRanksResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.RankCategoryId)) {
		query["rankCategoryId"] = request.RankCategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.RankCode)) {
		query["rankCode"] = request.RankCode
	}

	if !tea.BoolValue(util.IsUnset(request.RankName)) {
		query["rankName"] = request.RankName
	}

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
		Action:      tea.String("QueryJobRanks"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/jobRanks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryJobRanksResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询企业的职级信息
//
// @param request - QueryJobRanksRequest
//
// @return QueryJobRanksResponse
func (client *Client) QueryJobRanks(request *QueryJobRanksRequest) (_result *QueryJobRanksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryJobRanksHeaders{}
	_result = &QueryJobRanksResponse{}
	_body, _err := client.QueryJobRanksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询企业职务信息
//
// @param request - QueryJobsRequest
//
// @param headers - QueryJobsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryJobsResponse
func (client *Client) QueryJobsWithOptions(request *QueryJobsRequest, headers *QueryJobsHeaders, runtime *util.RuntimeOptions) (_result *QueryJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobName)) {
		query["jobName"] = request.JobName
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
		Action:      tea.String("QueryJobs"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/jobs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryJobsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询企业职务信息
//
// @param request - QueryJobsRequest
//
// @return QueryJobsResponse
func (client *Client) QueryJobs(request *QueryJobsRequest) (_result *QueryJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryJobsHeaders{}
	_result = &QueryJobsResponse{}
	_body, _err := client.QueryJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 智能人事查询微应用状态
//
// @param request - QueryMicroAppStatusRequest
//
// @param headers - QueryMicroAppStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMicroAppStatusResponse
func (client *Client) QueryMicroAppStatusWithOptions(request *QueryMicroAppStatusRequest, headers *QueryMicroAppStatusHeaders, runtime *util.RuntimeOptions) (_result *QueryMicroAppStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TenantIdList)) {
		body["tenantIdList"] = request.TenantIdList
	}

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
		Action:      tea.String("QueryMicroAppStatus"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/microApps/statuses/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMicroAppStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能人事查询微应用状态
//
// @param request - QueryMicroAppStatusRequest
//
// @return QueryMicroAppStatusResponse
func (client *Client) QueryMicroAppStatus(request *QueryMicroAppStatusRequest) (_result *QueryMicroAppStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMicroAppStatusHeaders{}
	_result = &QueryMicroAppStatusResponse{}
	_body, _err := client.QueryMicroAppStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 智能人事查询微应用可见性
//
// @param request - QueryMicroAppViewRequest
//
// @param headers - QueryMicroAppViewHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMicroAppViewResponse
func (client *Client) QueryMicroAppViewWithOptions(request *QueryMicroAppViewRequest, headers *QueryMicroAppViewHeaders, runtime *util.RuntimeOptions) (_result *QueryMicroAppViewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TenantIdList)) {
		body["tenantIdList"] = request.TenantIdList
	}

	if !tea.BoolValue(util.IsUnset(request.ViewUserId)) {
		body["viewUserId"] = request.ViewUserId
	}

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
		Action:      tea.String("QueryMicroAppView"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/microApps/visibilities/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMicroAppViewResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能人事查询微应用可见性
//
// @param request - QueryMicroAppViewRequest
//
// @return QueryMicroAppViewResponse
func (client *Client) QueryMicroAppView(request *QueryMicroAppViewRequest) (_result *QueryMicroAppViewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMicroAppViewHeaders{}
	_result = &QueryMicroAppViewResponse{}
	_body, _err := client.QueryMicroAppViewWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询企业职位版本
//
// @param headers - QueryPositionVersionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPositionVersionResponse
func (client *Client) QueryPositionVersionWithOptions(headers *QueryPositionVersionHeaders, runtime *util.RuntimeOptions) (_result *QueryPositionVersionResponse, _err error) {
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
		Action:      tea.String("QueryPositionVersion"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/positions/versions/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPositionVersionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询企业职位版本
//
// @return QueryPositionVersionResponse
func (client *Client) QueryPositionVersion() (_result *QueryPositionVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryPositionVersionHeaders{}
	_result = &QueryPositionVersionResponse{}
	_body, _err := client.QueryPositionVersionWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询企业职位信息
//
// @param request - QueryPositionsRequest
//
// @param headers - QueryPositionsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPositionsResponse
func (client *Client) QueryPositionsWithOptions(request *QueryPositionsRequest, headers *QueryPositionsHeaders, runtime *util.RuntimeOptions) (_result *QueryPositionsResponse, _err error) {
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

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.InCategoryIds)) {
		body["inCategoryIds"] = request.InCategoryIds
	}

	if !tea.BoolValue(util.IsUnset(request.InPositionIds)) {
		body["inPositionIds"] = request.InPositionIds
	}

	if !tea.BoolValue(util.IsUnset(request.PositionName)) {
		body["positionName"] = request.PositionName
	}

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
		Action:      tea.String("QueryPositions"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/positions/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPositionsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询企业职位信息
//
// @param request - QueryPositionsRequest
//
// @return QueryPositionsResponse
func (client *Client) QueryPositions(request *QueryPositionsRequest) (_result *QueryPositionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryPositionsHeaders{}
	_result = &QueryPositionsResponse{}
	_body, _err := client.QueryPositionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 撤回电子签署中的签署记录
//
// @param request - RevokeSignRecordsRequest
//
// @param headers - RevokeSignRecordsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeSignRecordsResponse
func (client *Client) RevokeSignRecordsWithOptions(request *RevokeSignRecordsRequest, headers *RevokeSignRecordsHeaders, runtime *util.RuntimeOptions) (_result *RevokeSignRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RevokeUserId)) {
		body["revokeUserId"] = request.RevokeUserId
	}

	if !tea.BoolValue(util.IsUnset(request.SignRecordIds)) {
		body["signRecordIds"] = request.SignRecordIds
	}

	if !tea.BoolValue(util.IsUnset(request.StatusRemark)) {
		body["statusRemark"] = request.StatusRemark
	}

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
		Action:      tea.String("RevokeSignRecords"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/masters/signCenters/records/revoke"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokeSignRecordsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 撤回电子签署中的签署记录
//
// @param request - RevokeSignRecordsRequest
//
// @return RevokeSignRecordsResponse
func (client *Client) RevokeSignRecords(request *RevokeSignRecordsRequest) (_result *RevokeSignRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RevokeSignRecordsHeaders{}
	_result = &RevokeSignRecordsResponse{}
	_body, _err := client.RevokeSignRecordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 撤销待离职
//
// @param request - RevokeTerminationRequest
//
// @param headers - RevokeTerminationHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeTerminationResponse
func (client *Client) RevokeTerminationWithOptions(request *RevokeTerminationRequest, headers *RevokeTerminationHeaders, runtime *util.RuntimeOptions) (_result *RevokeTerminationResponse, _err error) {
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
		Action:      tea.String("RevokeTermination"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/pendingDismission/revoke"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokeTerminationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 撤销待离职
//
// @param request - RevokeTerminationRequest
//
// @return RevokeTerminationResponse
func (client *Client) RevokeTermination(request *RevokeTerminationRequest) (_result *RevokeTerminationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RevokeTerminationHeaders{}
	_result = &RevokeTerminationResponse{}
	_body, _err := client.RevokeTerminationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询花名册中有权限的字段列表
//
// @param request - RosterMetaAvailableFieldListRequest
//
// @param headers - RosterMetaAvailableFieldListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RosterMetaAvailableFieldListResponse
func (client *Client) RosterMetaAvailableFieldListWithOptions(request *RosterMetaAvailableFieldListRequest, headers *RosterMetaAvailableFieldListHeaders, runtime *util.RuntimeOptions) (_result *RosterMetaAvailableFieldListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppAgentId)) {
		query["appAgentId"] = request.AppAgentId
	}

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
		Action:      tea.String("RosterMetaAvailableFieldList"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/rosters/meta/authorities/fields"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RosterMetaAvailableFieldListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询花名册中有权限的字段列表
//
// @param request - RosterMetaAvailableFieldListRequest
//
// @return RosterMetaAvailableFieldListResponse
func (client *Client) RosterMetaAvailableFieldList(request *RosterMetaAvailableFieldListRequest) (_result *RosterMetaAvailableFieldListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RosterMetaAvailableFieldListHeaders{}
	_result = &RosterMetaAvailableFieldListResponse{}
	_body, _err := client.RosterMetaAvailableFieldListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 智能人事花名册字段选项修改
//
// @param request - RosterMetaFieldOptionsUpdateRequest
//
// @param headers - RosterMetaFieldOptionsUpdateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RosterMetaFieldOptionsUpdateResponse
func (client *Client) RosterMetaFieldOptionsUpdateWithOptions(request *RosterMetaFieldOptionsUpdateRequest, headers *RosterMetaFieldOptionsUpdateHeaders, runtime *util.RuntimeOptions) (_result *RosterMetaFieldOptionsUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppAgentId)) {
		query["appAgentId"] = request.AppAgentId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FieldCode)) {
		body["fieldCode"] = request.FieldCode
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["groupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyType)) {
		body["modifyType"] = request.ModifyType
	}

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
		Action:      tea.String("RosterMetaFieldOptionsUpdate"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/rosters/meta/fields/options"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RosterMetaFieldOptionsUpdateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能人事花名册字段选项修改
//
// @param request - RosterMetaFieldOptionsUpdateRequest
//
// @return RosterMetaFieldOptionsUpdateResponse
func (client *Client) RosterMetaFieldOptionsUpdate(request *RosterMetaFieldOptionsUpdateRequest) (_result *RosterMetaFieldOptionsUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RosterMetaFieldOptionsUpdateHeaders{}
	_result = &RosterMetaFieldOptionsUpdateResponse{}
	_body, _err := client.RosterMetaFieldOptionsUpdateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ISV发送卡片消息
//
// @param request - SendIsvCardMessageRequest
//
// @param headers - SendIsvCardMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendIsvCardMessageResponse
func (client *Client) SendIsvCardMessageWithOptions(request *SendIsvCardMessageRequest, headers *SendIsvCardMessageHeaders, runtime *util.RuntimeOptions) (_result *SendIsvCardMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		query["agentId"] = request.AgentId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		body["messageType"] = request.MessageType
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUserIds)) {
		body["receiverUserIds"] = request.ReceiverUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.SceneType)) {
		body["sceneType"] = request.SceneType
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.SenderUserId)) {
		body["senderUserId"] = request.SenderUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ValueMap)) {
		body["valueMap"] = request.ValueMap
	}

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
		Action:      tea.String("SendIsvCardMessage"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/cardMessages/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendIsvCardMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ISV发送卡片消息
//
// @param request - SendIsvCardMessageRequest
//
// @return SendIsvCardMessageResponse
func (client *Client) SendIsvCardMessage(request *SendIsvCardMessageRequest) (_result *SendIsvCardMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendIsvCardMessageHeaders{}
	_result = &SendIsvCardMessageResponse{}
	_body, _err := client.SendIsvCardMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发送实人认证邀请消息
//
// @param request - SendRealAuthInviteMessageRequest
//
// @param headers - SendRealAuthInviteMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendRealAuthInviteMessageResponse
func (client *Client) SendRealAuthInviteMessageWithOptions(request *SendRealAuthInviteMessageRequest, headers *SendRealAuthInviteMessageHeaders, runtime *util.RuntimeOptions) (_result *SendRealAuthInviteMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InviterId)) {
		body["inviterId"] = request.InviterId
	}

	if !tea.BoolValue(util.IsUnset(request.OnWorkingEmpSearchVO)) {
		body["onWorkingEmpSearchVO"] = request.OnWorkingEmpSearchVO
	}

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
		Action:      tea.String("SendRealAuthInviteMessage"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/realAuth/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendRealAuthInviteMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送实人认证邀请消息
//
// @param request - SendRealAuthInviteMessageRequest
//
// @return SendRealAuthInviteMessageResponse
func (client *Client) SendRealAuthInviteMessage(request *SendRealAuthInviteMessageRequest) (_result *SendRealAuthInviteMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendRealAuthInviteMessageHeaders{}
	_result = &SendRealAuthInviteMessageResponse{}
	_body, _err := client.SendRealAuthInviteMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 初始化解决方案任务
//
// @param request - SolutionTaskInitRequest
//
// @param headers - SolutionTaskInitHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SolutionTaskInitResponse
func (client *Client) SolutionTaskInitWithOptions(request *SolutionTaskInitRequest, headers *SolutionTaskInitHeaders, runtime *util.RuntimeOptions) (_result *SolutionTaskInitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SolutionType)) {
		query["solutionType"] = request.SolutionType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.ClaimTime)) {
		body["claimTime"] = request.ClaimTime
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FinishTime)) {
		body["finishTime"] = request.FinishTime
	}

	if !tea.BoolValue(util.IsUnset(request.OuterId)) {
		body["outerId"] = request.OuterId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
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
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SolutionTaskInit"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/solutions/tasks/init"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SolutionTaskInitResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 初始化解决方案任务
//
// @param request - SolutionTaskInitRequest
//
// @return SolutionTaskInitResponse
func (client *Client) SolutionTaskInit(request *SolutionTaskInitRequest) (_result *SolutionTaskInitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SolutionTaskInitHeaders{}
	_result = &SolutionTaskInitResponse{}
	_body, _err := client.SolutionTaskInitWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 保存解决方案任务
//
// @param request - SolutionTaskSaveRequest
//
// @param headers - SolutionTaskSaveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SolutionTaskSaveResponse
func (client *Client) SolutionTaskSaveWithOptions(request *SolutionTaskSaveRequest, headers *SolutionTaskSaveHeaders, runtime *util.RuntimeOptions) (_result *SolutionTaskSaveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SolutionType)) {
		query["solutionType"] = request.SolutionType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClaimTime)) {
		body["claimTime"] = request.ClaimTime
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FinishTime)) {
		body["finishTime"] = request.FinishTime
	}

	if !tea.BoolValue(util.IsUnset(request.OuterId)) {
		body["outerId"] = request.OuterId
	}

	if !tea.BoolValue(util.IsUnset(request.SolutionInstanceId)) {
		body["solutionInstanceId"] = request.SolutionInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		body["taskType"] = request.TaskType
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateOuterId)) {
		body["templateOuterId"] = request.TemplateOuterId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
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
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SolutionTaskSave"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/solutions/tasks/save"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SolutionTaskSaveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存解决方案任务
//
// @param request - SolutionTaskSaveRequest
//
// @return SolutionTaskSaveResponse
func (client *Client) SolutionTaskSave(request *SolutionTaskSaveRequest) (_result *SolutionTaskSaveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SolutionTaskSaveHeaders{}
	_result = &SolutionTaskSaveResponse{}
	_body, _err := client.SolutionTaskSaveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 同步解决方案状态
//
// @param request - SyncSolutionStatusRequest
//
// @param headers - SyncSolutionStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncSolutionStatusResponse
func (client *Client) SyncSolutionStatusWithOptions(request *SyncSolutionStatusRequest, headers *SyncSolutionStatusHeaders, runtime *util.RuntimeOptions) (_result *SyncSolutionStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.SolutionStatus)) {
		body["solutionStatus"] = request.SolutionStatus
	}

	if !tea.BoolValue(util.IsUnset(request.SolutionType)) {
		body["solutionType"] = request.SolutionType
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["tenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
	}

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
		Action:      tea.String("SyncSolutionStatus"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/solutions/statuses/sync"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncSolutionStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步解决方案状态
//
// @param request - SyncSolutionStatusRequest
//
// @return SyncSolutionStatusResponse
func (client *Client) SyncSolutionStatus(request *SyncSolutionStatusRequest) (_result *SyncSolutionStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncSolutionStatusHeaders{}
	_result = &SyncSolutionStatusResponse{}
	_body, _err := client.SyncSolutionStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 同步解决方案任务模版
//
// @param request - SyncTaskTemplateRequest
//
// @param headers - SyncTaskTemplateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncTaskTemplateResponse
func (client *Client) SyncTaskTemplateWithOptions(request *SyncTaskTemplateRequest, headers *SyncTaskTemplateHeaders, runtime *util.RuntimeOptions) (_result *SyncTaskTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SolutionType)) {
		query["solutionType"] = request.SolutionType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Delete)) {
		body["delete"] = request.Delete
	}

	if !tea.BoolValue(util.IsUnset(request.Des)) {
		body["des"] = request.Des
	}

	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OptUserId)) {
		body["optUserId"] = request.OptUserId
	}

	if !tea.BoolValue(util.IsUnset(request.OuterId)) {
		body["outerId"] = request.OuterId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskScopeVO)) {
		body["taskScopeVO"] = request.TaskScopeVO
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		body["taskType"] = request.TaskType
	}

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
		Action:      tea.String("SyncTaskTemplate"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/solutions/tasks/templates/sync"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncTaskTemplateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步解决方案任务模版
//
// @param request - SyncTaskTemplateRequest
//
// @return SyncTaskTemplateResponse
func (client *Client) SyncTaskTemplate(request *SyncTaskTemplateRequest) (_result *SyncTaskTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncTaskTemplateHeaders{}
	_result = &SyncTaskTemplateResponse{}
	_body, _err := client.SyncTaskTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新花名册自定义字段
//
// @param request - UpdateCustomRosterFieldRequest
//
// @param headers - UpdateCustomRosterFieldHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomRosterFieldResponse
func (client *Client) UpdateCustomRosterFieldWithOptions(request *UpdateCustomRosterFieldRequest, headers *UpdateCustomRosterFieldHeaders, runtime *util.RuntimeOptions) (_result *UpdateCustomRosterFieldResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactClientFlag)) {
		body["contactClientFlag"] = request.ContactClientFlag
	}

	if !tea.BoolValue(util.IsUnset(request.ContactFlag)) {
		body["contactFlag"] = request.ContactFlag
	}

	if !tea.BoolValue(util.IsUnset(request.ContactSource)) {
		body["contactSource"] = request.ContactSource
	}

	if !tea.BoolValue(util.IsUnset(request.ContactSystemFlag)) {
		body["contactSystemFlag"] = request.ContactSystemFlag
	}

	if !tea.BoolValue(util.IsUnset(request.Deleted)) {
		body["deleted"] = request.Deleted
	}

	if !tea.BoolValue(util.IsUnset(request.Derived)) {
		body["derived"] = request.Derived
	}

	if !tea.BoolValue(util.IsUnset(request.Disabled)) {
		body["disabled"] = request.Disabled
	}

	if !tea.BoolValue(util.IsUnset(request.EditFromEmployeeFlag)) {
		body["editFromEmployeeFlag"] = request.EditFromEmployeeFlag
	}

	if !tea.BoolValue(util.IsUnset(request.EditableByHr)) {
		body["editableByHr"] = request.EditableByHr
	}

	if !tea.BoolValue(util.IsUnset(request.FieldCode)) {
		body["fieldCode"] = request.FieldCode
	}

	if !tea.BoolValue(util.IsUnset(request.FieldName)) {
		body["fieldName"] = request.FieldName
	}

	if !tea.BoolValue(util.IsUnset(request.FieldTip)) {
		body["fieldTip"] = request.FieldTip
	}

	if !tea.BoolValue(util.IsUnset(request.FieldType)) {
		body["fieldType"] = request.FieldType
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["groupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.HiddenFromEmployeeFlag)) {
		body["hiddenFromEmployeeFlag"] = request.HiddenFromEmployeeFlag
	}

	if !tea.BoolValue(util.IsUnset(request.Hint)) {
		body["hint"] = request.Hint
	}

	if !tea.BoolValue(util.IsUnset(request.HistoryField)) {
		body["historyField"] = request.HistoryField
	}

	if !tea.BoolValue(util.IsUnset(request.Index)) {
		body["index"] = request.Index
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyOptions)) {
		body["modifyOptions"] = request.ModifyOptions
	}

	if !tea.BoolValue(util.IsUnset(request.NoWatermark)) {
		body["noWatermark"] = request.NoWatermark
	}

	if !tea.BoolValue(util.IsUnset(request.NumberDecimalPlace)) {
		body["numberDecimalPlace"] = request.NumberDecimalPlace
	}

	if !tea.BoolValue(util.IsUnset(request.NumberFormatType)) {
		body["numberFormatType"] = request.NumberFormatType
	}

	if !tea.BoolValue(util.IsUnset(request.NumberValueType)) {
		body["numberValueType"] = request.NumberValueType
	}

	if !tea.BoolValue(util.IsUnset(request.OptionText)) {
		body["optionText"] = request.OptionText
	}

	if !tea.BoolValue(util.IsUnset(request.Required)) {
		body["required"] = request.Required
	}

	if !tea.BoolValue(util.IsUnset(request.SourceFieldCode)) {
		body["sourceFieldCode"] = request.SourceFieldCode
	}

	if !tea.BoolValue(util.IsUnset(request.SystemFlag)) {
		body["systemFlag"] = request.SystemFlag
	}

	if !tea.BoolValue(util.IsUnset(request.TextToSelectField)) {
		body["textToSelectField"] = request.TextToSelectField
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateUserId)) {
		body["updateUserId"] = request.UpdateUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["value"] = request.Value
	}

	if !tea.BoolValue(util.IsUnset(request.VisibleByEmp)) {
		body["visibleByEmp"] = request.VisibleByEmp
	}

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
		Action:      tea.String("UpdateCustomRosterField"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/customRosterField/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCustomRosterFieldResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新花名册自定义字段
//
// @param request - UpdateCustomRosterFieldRequest
//
// @return UpdateCustomRosterFieldResponse
func (client *Client) UpdateCustomRosterField(request *UpdateCustomRosterFieldRequest) (_result *UpdateCustomRosterFieldResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateCustomRosterFieldHeaders{}
	_result = &UpdateCustomRosterFieldResponse{}
	_body, _err := client.UpdateCustomRosterFieldWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新待离职信息
//
// @param request - UpdateEmpDismissionInfoRequest
//
// @param headers - UpdateEmpDismissionInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEmpDismissionInfoResponse
func (client *Client) UpdateEmpDismissionInfoWithOptions(request *UpdateEmpDismissionInfoRequest, headers *UpdateEmpDismissionInfoHeaders, runtime *util.RuntimeOptions) (_result *UpdateEmpDismissionInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DismissionMemo)) {
		body["dismissionMemo"] = request.DismissionMemo
	}

	if !tea.BoolValue(util.IsUnset(request.LastWorkDate)) {
		body["lastWorkDate"] = request.LastWorkDate
	}

	if !tea.BoolValue(util.IsUnset(request.Partner)) {
		body["partner"] = request.Partner
	}

	if !tea.BoolValue(util.IsUnset(request.TerminationReasonPassive)) {
		body["terminationReasonPassive"] = request.TerminationReasonPassive
	}

	if !tea.BoolValue(util.IsUnset(request.TerminationReasonVoluntary)) {
		body["terminationReasonVoluntary"] = request.TerminationReasonVoluntary
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
		Action:      tea.String("UpdateEmpDismissionInfo"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/pendingDismission/infos"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEmpDismissionInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新待离职信息
//
// @param request - UpdateEmpDismissionInfoRequest
//
// @return UpdateEmpDismissionInfoResponse
func (client *Client) UpdateEmpDismissionInfo(request *UpdateEmpDismissionInfoRequest) (_result *UpdateEmpDismissionInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateEmpDismissionInfoHeaders{}
	_result = &UpdateEmpDismissionInfoResponse{}
	_body, _err := client.UpdateEmpDismissionInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新法人公司名称
//
// @param request - UpdateHrmLegalEntityNameRequest
//
// @param headers - UpdateHrmLegalEntityNameHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHrmLegalEntityNameResponse
func (client *Client) UpdateHrmLegalEntityNameWithOptions(request *UpdateHrmLegalEntityNameRequest, headers *UpdateHrmLegalEntityNameHeaders, runtime *util.RuntimeOptions) (_result *UpdateHrmLegalEntityNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingTenantId)) {
		query["dingTenantId"] = request.DingTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.LegalEntityName)) {
		query["legalEntityName"] = request.LegalEntityName
	}

	if !tea.BoolValue(util.IsUnset(request.OriginLegalEntityName)) {
		query["originLegalEntityName"] = request.OriginLegalEntityName
	}

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
		Action:      tea.String("UpdateHrmLegalEntityName"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/masters/legalEntities/companyNames"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateHrmLegalEntityNameResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新法人公司名称
//
// @param request - UpdateHrmLegalEntityNameRequest
//
// @return UpdateHrmLegalEntityNameResponse
func (client *Client) UpdateHrmLegalEntityName(request *UpdateHrmLegalEntityNameRequest) (_result *UpdateHrmLegalEntityNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateHrmLegalEntityNameHeaders{}
	_result = &UpdateHrmLegalEntityNameResponse{}
	_body, _err := client.UpdateHrmLegalEntityNameWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新法人公司
//
// @param request - UpdateHrmLegalEntityWithoutNameRequest
//
// @param headers - UpdateHrmLegalEntityWithoutNameHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHrmLegalEntityWithoutNameResponse
func (client *Client) UpdateHrmLegalEntityWithoutNameWithOptions(request *UpdateHrmLegalEntityWithoutNameRequest, headers *UpdateHrmLegalEntityWithoutNameHeaders, runtime *util.RuntimeOptions) (_result *UpdateHrmLegalEntityWithoutNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingTenantId)) {
		query["dingTenantId"] = request.DingTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.CreateUserId)) {
		body["createUserId"] = request.CreateUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.LegalEntityName)) {
		body["legalEntityName"] = request.LegalEntityName
	}

	if !tea.BoolValue(util.IsUnset(request.LegalEntityShortName)) {
		body["legalEntityShortName"] = request.LegalEntityShortName
	}

	if !tea.BoolValue(util.IsUnset(request.LegalEntityStatus)) {
		body["legalEntityStatus"] = request.LegalEntityStatus
	}

	if !tea.BoolValue(util.IsUnset(request.LegalPersonName)) {
		body["legalPersonName"] = request.LegalPersonName
	}

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
		Action:      tea.String("UpdateHrmLegalEntityWithoutName"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/masters/legalEntities/companies"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateHrmLegalEntityWithoutNameResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新法人公司
//
// @param request - UpdateHrmLegalEntityWithoutNameRequest
//
// @return UpdateHrmLegalEntityWithoutNameResponse
func (client *Client) UpdateHrmLegalEntityWithoutName(request *UpdateHrmLegalEntityWithoutNameRequest) (_result *UpdateHrmLegalEntityWithoutNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateHrmLegalEntityWithoutNameHeaders{}
	_result = &UpdateHrmLegalEntityWithoutNameResponse{}
	_body, _err := client.UpdateHrmLegalEntityWithoutNameWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 智能人事更新版本回退按钮状态
//
// @param request - UpdateHrmVersionRollBackStatusRequest
//
// @param headers - UpdateHrmVersionRollBackStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHrmVersionRollBackStatusResponse
func (client *Client) UpdateHrmVersionRollBackStatusWithOptions(request *UpdateHrmVersionRollBackStatusRequest, headers *UpdateHrmVersionRollBackStatusHeaders, runtime *util.RuntimeOptions) (_result *UpdateHrmVersionRollBackStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigValue)) {
		body["configValue"] = request.ConfigValue
	}

	if !tea.BoolValue(util.IsUnset(request.OptUserId)) {
		body["optUserId"] = request.OptUserId
	}

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
		Action:      tea.String("UpdateHrmVersionRollBackStatus"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/versions/rollbackButtons/statuses"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateHrmVersionRollBackStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能人事更新版本回退按钮状态
//
// @param request - UpdateHrmVersionRollBackStatusRequest
//
// @return UpdateHrmVersionRollBackStatusResponse
func (client *Client) UpdateHrmVersionRollBackStatus(request *UpdateHrmVersionRollBackStatusRequest) (_result *UpdateHrmVersionRollBackStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateHrmVersionRollBackStatusHeaders{}
	_result = &UpdateHrmVersionRollBackStatusResponse{}
	_body, _err := client.UpdateHrmVersionRollBackStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ISV更新卡片消息
//
// @param request - UpdateIsvCardMessageRequest
//
// @param headers - UpdateIsvCardMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIsvCardMessageResponse
func (client *Client) UpdateIsvCardMessageWithOptions(request *UpdateIsvCardMessageRequest, headers *UpdateIsvCardMessageHeaders, runtime *util.RuntimeOptions) (_result *UpdateIsvCardMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		query["agentId"] = request.AgentId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		body["messageType"] = request.MessageType
	}

	if !tea.BoolValue(util.IsUnset(request.SceneType)) {
		body["sceneType"] = request.SceneType
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.ValueMap)) {
		body["valueMap"] = request.ValueMap
	}

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
		Action:      tea.String("UpdateIsvCardMessage"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/cardMessages"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIsvCardMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ISV更新卡片消息
//
// @param request - UpdateIsvCardMessageRequest
//
// @return UpdateIsvCardMessageResponse
func (client *Client) UpdateIsvCardMessage(request *UpdateIsvCardMessageRequest) (_result *UpdateIsvCardMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateIsvCardMessageHeaders{}
	_result = &UpdateIsvCardMessageResponse{}
	_body, _err := client.UpdateIsvCardMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新花名册字段分组
//
// @param request - UpdateRosterFieldFormRequest
//
// @param headers - UpdateRosterFieldFormHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRosterFieldFormResponse
func (client *Client) UpdateRosterFieldFormWithOptions(request *UpdateRosterFieldFormRequest, headers *UpdateRosterFieldFormHeaders, runtime *util.RuntimeOptions) (_result *UpdateRosterFieldFormResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Detail)) {
		body["detail"] = request.Detail
	}

	if !tea.BoolValue(util.IsUnset(request.FormId)) {
		body["formId"] = request.FormId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
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
		Action:      tea.String("UpdateRosterFieldForm"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/rosterFieldForm/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRosterFieldFormResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新花名册字段分组
//
// @param request - UpdateRosterFieldFormRequest
//
// @return UpdateRosterFieldFormResponse
func (client *Client) UpdateRosterFieldForm(request *UpdateRosterFieldFormRequest) (_result *UpdateRosterFieldFormResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateRosterFieldFormHeaders{}
	_result = &UpdateRosterFieldFormResponse{}
	_body, _err := client.UpdateRosterFieldFormWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上传附件材料
//
// @param request - UploadAttachmentRequest
//
// @param headers - UploadAttachmentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadAttachmentResponse
func (client *Client) UploadAttachmentWithOptions(request *UploadAttachmentRequest, headers *UploadAttachmentHeaders, runtime *util.RuntimeOptions) (_result *UploadAttachmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MediaId)) {
		body["mediaId"] = request.MediaId
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
		Action:      tea.String("UploadAttachment"),
		Version:     tea.String("hrm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/hrm/attachments/upload"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadAttachmentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传附件材料
//
// @param request - UploadAttachmentRequest
//
// @return UploadAttachmentResponse
func (client *Client) UploadAttachment(request *UploadAttachmentRequest) (_result *UploadAttachmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UploadAttachmentHeaders{}
	_result = &UploadAttachmentResponse{}
	_body, _err := client.UploadAttachmentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
