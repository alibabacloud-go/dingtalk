// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package trip_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetTravelProcessDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTravelProcessDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTravelProcessDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetTravelProcessDetailHeaders) SetCommonHeaders(v map[string]*string) *GetTravelProcessDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTravelProcessDetailHeaders) SetXAcsDingtalkAccessToken(v string) *GetTravelProcessDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTravelProcessDetailRequest struct {
	ProcessCorpId     *string `json:"processCorpId,omitempty" xml:"processCorpId,omitempty"`
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
}

func (s GetTravelProcessDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTravelProcessDetailRequest) GoString() string {
	return s.String()
}

func (s *GetTravelProcessDetailRequest) SetProcessCorpId(v string) *GetTravelProcessDetailRequest {
	s.ProcessCorpId = &v
	return s
}

func (s *GetTravelProcessDetailRequest) SetProcessInstanceId(v string) *GetTravelProcessDetailRequest {
	s.ProcessInstanceId = &v
	return s
}

type GetTravelProcessDetailResponseBody struct {
	Result  *GetTravelProcessDetailResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTravelProcessDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTravelProcessDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetTravelProcessDetailResponseBody) SetResult(v *GetTravelProcessDetailResponseBodyResult) *GetTravelProcessDetailResponseBody {
	s.Result = v
	return s
}

func (s *GetTravelProcessDetailResponseBody) SetSuccess(v bool) *GetTravelProcessDetailResponseBody {
	s.Success = &v
	return s
}

type GetTravelProcessDetailResponseBodyResult struct {
	BizCategoryId                *string                                                     `json:"bizCategoryId,omitempty" xml:"bizCategoryId,omitempty"`
	BusinessId                   *string                                                     `json:"businessId,omitempty" xml:"businessId,omitempty"`
	CorpId                       *string                                                     `json:"corpId,omitempty" xml:"corpId,omitempty"`
	CostCenter                   *string                                                     `json:"costCenter,omitempty" xml:"costCenter,omitempty"`
	CostCenterId                 *string                                                     `json:"costCenterId,omitempty" xml:"costCenterId,omitempty"`
	CostCenterThirdPartyId       *string                                                     `json:"costCenterThirdPartyId,omitempty" xml:"costCenterThirdPartyId,omitempty"`
	CreateTime                   *string                                                     `json:"createTime,omitempty" xml:"createTime,omitempty"`
	ExtFormComponent             []*GetTravelProcessDetailResponseBodyResultExtFormComponent `json:"extFormComponent,omitempty" xml:"extFormComponent,omitempty" type:"Repeated"`
	FeeType                      *string                                                     `json:"feeType,omitempty" xml:"feeType,omitempty"`
	InvoiceTitle                 *string                                                     `json:"invoiceTitle,omitempty" xml:"invoiceTitle,omitempty"`
	InvoiceTitleId               *string                                                     `json:"invoiceTitleId,omitempty" xml:"invoiceTitleId,omitempty"`
	InvoiceTitleThirdPartyId     *string                                                     `json:"invoiceTitleThirdPartyId,omitempty" xml:"invoiceTitleThirdPartyId,omitempty"`
	ItineraryProject             *string                                                     `json:"itineraryProject,omitempty" xml:"itineraryProject,omitempty"`
	ItineraryProjectThirdPartyId *string                                                     `json:"itineraryProjectThirdPartyId,omitempty" xml:"itineraryProjectThirdPartyId,omitempty"`
	Journeys                     []*GetTravelProcessDetailResponseBodyResultJourneys         `json:"journeys,omitempty" xml:"journeys,omitempty" type:"Repeated"`
	MainProcessInstanceId        *string                                                     `json:"mainProcessInstanceId,omitempty" xml:"mainProcessInstanceId,omitempty"`
	Memo                         *string                                                     `json:"memo,omitempty" xml:"memo,omitempty"`
	OriginatorId                 *string                                                     `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	ProcessInstanceId            *string                                                     `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	ProcessResult                *string                                                     `json:"processResult,omitempty" xml:"processResult,omitempty"`
	ProcessStatus                *string                                                     `json:"processStatus,omitempty" xml:"processStatus,omitempty"`
	Remark                       *string                                                     `json:"remark,omitempty" xml:"remark,omitempty"`
	TravelCategory               *string                                                     `json:"travelCategory,omitempty" xml:"travelCategory,omitempty"`
	Travelers                    []*string                                                   `json:"travelers,omitempty" xml:"travelers,omitempty" type:"Repeated"`
}

func (s GetTravelProcessDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetTravelProcessDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetTravelProcessDetailResponseBodyResult) SetBizCategoryId(v string) *GetTravelProcessDetailResponseBodyResult {
	s.BizCategoryId = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResult) SetBusinessId(v string) *GetTravelProcessDetailResponseBodyResult {
	s.BusinessId = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResult) SetCorpId(v string) *GetTravelProcessDetailResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResult) SetCostCenter(v string) *GetTravelProcessDetailResponseBodyResult {
	s.CostCenter = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResult) SetCostCenterId(v string) *GetTravelProcessDetailResponseBodyResult {
	s.CostCenterId = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResult) SetCostCenterThirdPartyId(v string) *GetTravelProcessDetailResponseBodyResult {
	s.CostCenterThirdPartyId = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResult) SetCreateTime(v string) *GetTravelProcessDetailResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResult) SetExtFormComponent(v []*GetTravelProcessDetailResponseBodyResultExtFormComponent) *GetTravelProcessDetailResponseBodyResult {
	s.ExtFormComponent = v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResult) SetFeeType(v string) *GetTravelProcessDetailResponseBodyResult {
	s.FeeType = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResult) SetInvoiceTitle(v string) *GetTravelProcessDetailResponseBodyResult {
	s.InvoiceTitle = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResult) SetInvoiceTitleId(v string) *GetTravelProcessDetailResponseBodyResult {
	s.InvoiceTitleId = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResult) SetInvoiceTitleThirdPartyId(v string) *GetTravelProcessDetailResponseBodyResult {
	s.InvoiceTitleThirdPartyId = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResult) SetItineraryProject(v string) *GetTravelProcessDetailResponseBodyResult {
	s.ItineraryProject = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResult) SetItineraryProjectThirdPartyId(v string) *GetTravelProcessDetailResponseBodyResult {
	s.ItineraryProjectThirdPartyId = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResult) SetJourneys(v []*GetTravelProcessDetailResponseBodyResultJourneys) *GetTravelProcessDetailResponseBodyResult {
	s.Journeys = v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResult) SetMainProcessInstanceId(v string) *GetTravelProcessDetailResponseBodyResult {
	s.MainProcessInstanceId = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResult) SetMemo(v string) *GetTravelProcessDetailResponseBodyResult {
	s.Memo = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResult) SetOriginatorId(v string) *GetTravelProcessDetailResponseBodyResult {
	s.OriginatorId = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResult) SetProcessInstanceId(v string) *GetTravelProcessDetailResponseBodyResult {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResult) SetProcessResult(v string) *GetTravelProcessDetailResponseBodyResult {
	s.ProcessResult = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResult) SetProcessStatus(v string) *GetTravelProcessDetailResponseBodyResult {
	s.ProcessStatus = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResult) SetRemark(v string) *GetTravelProcessDetailResponseBodyResult {
	s.Remark = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResult) SetTravelCategory(v string) *GetTravelProcessDetailResponseBodyResult {
	s.TravelCategory = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResult) SetTravelers(v []*string) *GetTravelProcessDetailResponseBodyResult {
	s.Travelers = v
	return s
}

type GetTravelProcessDetailResponseBodyResultExtFormComponent struct {
	BizAlias      *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	ExtValue      *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	Id            *string `json:"id,omitempty" xml:"id,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	Value         *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetTravelProcessDetailResponseBodyResultExtFormComponent) String() string {
	return tea.Prettify(s)
}

func (s GetTravelProcessDetailResponseBodyResultExtFormComponent) GoString() string {
	return s.String()
}

func (s *GetTravelProcessDetailResponseBodyResultExtFormComponent) SetBizAlias(v string) *GetTravelProcessDetailResponseBodyResultExtFormComponent {
	s.BizAlias = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultExtFormComponent) SetComponentType(v string) *GetTravelProcessDetailResponseBodyResultExtFormComponent {
	s.ComponentType = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultExtFormComponent) SetExtValue(v string) *GetTravelProcessDetailResponseBodyResultExtFormComponent {
	s.ExtValue = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultExtFormComponent) SetId(v string) *GetTravelProcessDetailResponseBodyResultExtFormComponent {
	s.Id = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultExtFormComponent) SetName(v string) *GetTravelProcessDetailResponseBodyResultExtFormComponent {
	s.Name = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultExtFormComponent) SetValue(v string) *GetTravelProcessDetailResponseBodyResultExtFormComponent {
	s.Value = &v
	return s
}

type GetTravelProcessDetailResponseBodyResultJourneys struct {
	Arrival                      *GetTravelProcessDetailResponseBodyResultJourneysArrival   `json:"arrival,omitempty" xml:"arrival,omitempty" type:"Struct"`
	CostCenter                   *string                                                    `json:"costCenter,omitempty" xml:"costCenter,omitempty"`
	CostCenterId                 *string                                                    `json:"costCenterId,omitempty" xml:"costCenterId,omitempty"`
	CostCenterThirdPartyId       *string                                                    `json:"costCenterThirdPartyId,omitempty" xml:"costCenterThirdPartyId,omitempty"`
	Departure                    *GetTravelProcessDetailResponseBodyResultJourneysDeparture `json:"departure,omitempty" xml:"departure,omitempty" type:"Struct"`
	EndTime                      *string                                                    `json:"endTime,omitempty" xml:"endTime,omitempty"`
	EndTimeAcc                   *string                                                    `json:"endTimeAcc,omitempty" xml:"endTimeAcc,omitempty"`
	InvoiceTitle                 *string                                                    `json:"invoiceTitle,omitempty" xml:"invoiceTitle,omitempty"`
	InvoiceTitleId               *string                                                    `json:"invoiceTitleId,omitempty" xml:"invoiceTitleId,omitempty"`
	InvoiceTitleThirdPartyId     *string                                                    `json:"invoiceTitleThirdPartyId,omitempty" xml:"invoiceTitleThirdPartyId,omitempty"`
	ItineraryProject             *string                                                    `json:"itineraryProject,omitempty" xml:"itineraryProject,omitempty"`
	ItineraryProjectId           *string                                                    `json:"itineraryProjectId,omitempty" xml:"itineraryProjectId,omitempty"`
	ItineraryProjectThirdPartyId *string                                                    `json:"itineraryProjectThirdPartyId,omitempty" xml:"itineraryProjectThirdPartyId,omitempty"`
	JourneyBizNo                 *string                                                    `json:"journeyBizNo,omitempty" xml:"journeyBizNo,omitempty"`
	StartTime                    *string                                                    `json:"startTime,omitempty" xml:"startTime,omitempty"`
	StartTimeAcc                 *string                                                    `json:"startTimeAcc,omitempty" xml:"startTimeAcc,omitempty"`
	TimeUnit                     *string                                                    `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
	TravelType                   *string                                                    `json:"travelType,omitempty" xml:"travelType,omitempty"`
	TripWay                      *string                                                    `json:"tripWay,omitempty" xml:"tripWay,omitempty"`
}

func (s GetTravelProcessDetailResponseBodyResultJourneys) String() string {
	return tea.Prettify(s)
}

func (s GetTravelProcessDetailResponseBodyResultJourneys) GoString() string {
	return s.String()
}

func (s *GetTravelProcessDetailResponseBodyResultJourneys) SetArrival(v *GetTravelProcessDetailResponseBodyResultJourneysArrival) *GetTravelProcessDetailResponseBodyResultJourneys {
	s.Arrival = v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneys) SetCostCenter(v string) *GetTravelProcessDetailResponseBodyResultJourneys {
	s.CostCenter = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneys) SetCostCenterId(v string) *GetTravelProcessDetailResponseBodyResultJourneys {
	s.CostCenterId = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneys) SetCostCenterThirdPartyId(v string) *GetTravelProcessDetailResponseBodyResultJourneys {
	s.CostCenterThirdPartyId = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneys) SetDeparture(v *GetTravelProcessDetailResponseBodyResultJourneysDeparture) *GetTravelProcessDetailResponseBodyResultJourneys {
	s.Departure = v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneys) SetEndTime(v string) *GetTravelProcessDetailResponseBodyResultJourneys {
	s.EndTime = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneys) SetEndTimeAcc(v string) *GetTravelProcessDetailResponseBodyResultJourneys {
	s.EndTimeAcc = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneys) SetInvoiceTitle(v string) *GetTravelProcessDetailResponseBodyResultJourneys {
	s.InvoiceTitle = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneys) SetInvoiceTitleId(v string) *GetTravelProcessDetailResponseBodyResultJourneys {
	s.InvoiceTitleId = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneys) SetInvoiceTitleThirdPartyId(v string) *GetTravelProcessDetailResponseBodyResultJourneys {
	s.InvoiceTitleThirdPartyId = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneys) SetItineraryProject(v string) *GetTravelProcessDetailResponseBodyResultJourneys {
	s.ItineraryProject = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneys) SetItineraryProjectId(v string) *GetTravelProcessDetailResponseBodyResultJourneys {
	s.ItineraryProjectId = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneys) SetItineraryProjectThirdPartyId(v string) *GetTravelProcessDetailResponseBodyResultJourneys {
	s.ItineraryProjectThirdPartyId = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneys) SetJourneyBizNo(v string) *GetTravelProcessDetailResponseBodyResultJourneys {
	s.JourneyBizNo = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneys) SetStartTime(v string) *GetTravelProcessDetailResponseBodyResultJourneys {
	s.StartTime = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneys) SetStartTimeAcc(v string) *GetTravelProcessDetailResponseBodyResultJourneys {
	s.StartTimeAcc = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneys) SetTimeUnit(v string) *GetTravelProcessDetailResponseBodyResultJourneys {
	s.TimeUnit = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneys) SetTravelType(v string) *GetTravelProcessDetailResponseBodyResultJourneys {
	s.TravelType = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneys) SetTripWay(v string) *GetTravelProcessDetailResponseBodyResultJourneys {
	s.TripWay = &v
	return s
}

type GetTravelProcessDetailResponseBodyResultJourneysArrival struct {
	Code             *string `json:"code,omitempty" xml:"code,omitempty"`
	CountryCode      *string `json:"countryCode,omitempty" xml:"countryCode,omitempty"`
	CountryName      *string `json:"countryName,omitempty" xml:"countryName,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	NationalCityCode *string `json:"nationalCityCode,omitempty" xml:"nationalCityCode,omitempty"`
}

func (s GetTravelProcessDetailResponseBodyResultJourneysArrival) String() string {
	return tea.Prettify(s)
}

func (s GetTravelProcessDetailResponseBodyResultJourneysArrival) GoString() string {
	return s.String()
}

func (s *GetTravelProcessDetailResponseBodyResultJourneysArrival) SetCode(v string) *GetTravelProcessDetailResponseBodyResultJourneysArrival {
	s.Code = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneysArrival) SetCountryCode(v string) *GetTravelProcessDetailResponseBodyResultJourneysArrival {
	s.CountryCode = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneysArrival) SetCountryName(v string) *GetTravelProcessDetailResponseBodyResultJourneysArrival {
	s.CountryName = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneysArrival) SetName(v string) *GetTravelProcessDetailResponseBodyResultJourneysArrival {
	s.Name = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneysArrival) SetNationalCityCode(v string) *GetTravelProcessDetailResponseBodyResultJourneysArrival {
	s.NationalCityCode = &v
	return s
}

type GetTravelProcessDetailResponseBodyResultJourneysDeparture struct {
	Code             *string `json:"code,omitempty" xml:"code,omitempty"`
	CountryCode      *string `json:"countryCode,omitempty" xml:"countryCode,omitempty"`
	CountryName      *string `json:"countryName,omitempty" xml:"countryName,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	NationalCityCode *string `json:"nationalCityCode,omitempty" xml:"nationalCityCode,omitempty"`
}

func (s GetTravelProcessDetailResponseBodyResultJourneysDeparture) String() string {
	return tea.Prettify(s)
}

func (s GetTravelProcessDetailResponseBodyResultJourneysDeparture) GoString() string {
	return s.String()
}

func (s *GetTravelProcessDetailResponseBodyResultJourneysDeparture) SetCode(v string) *GetTravelProcessDetailResponseBodyResultJourneysDeparture {
	s.Code = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneysDeparture) SetCountryCode(v string) *GetTravelProcessDetailResponseBodyResultJourneysDeparture {
	s.CountryCode = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneysDeparture) SetCountryName(v string) *GetTravelProcessDetailResponseBodyResultJourneysDeparture {
	s.CountryName = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneysDeparture) SetName(v string) *GetTravelProcessDetailResponseBodyResultJourneysDeparture {
	s.Name = &v
	return s
}

func (s *GetTravelProcessDetailResponseBodyResultJourneysDeparture) SetNationalCityCode(v string) *GetTravelProcessDetailResponseBodyResultJourneysDeparture {
	s.NationalCityCode = &v
	return s
}

type GetTravelProcessDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTravelProcessDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTravelProcessDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTravelProcessDetailResponse) GoString() string {
	return s.String()
}

func (s *GetTravelProcessDetailResponse) SetHeaders(v map[string]*string) *GetTravelProcessDetailResponse {
	s.Headers = v
	return s
}

func (s *GetTravelProcessDetailResponse) SetStatusCode(v int32) *GetTravelProcessDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTravelProcessDetailResponse) SetBody(v *GetTravelProcessDetailResponseBody) *GetTravelProcessDetailResponse {
	s.Body = v
	return s
}

type PreCheckTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PreCheckTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s PreCheckTemplateHeaders) GoString() string {
	return s.String()
}

func (s *PreCheckTemplateHeaders) SetCommonHeaders(v map[string]*string) *PreCheckTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PreCheckTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *PreCheckTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PreCheckTemplateRequest struct {
	CustomerCorpId *string `json:"customerCorpId,omitempty" xml:"customerCorpId,omitempty"`
}

func (s PreCheckTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s PreCheckTemplateRequest) GoString() string {
	return s.String()
}

func (s *PreCheckTemplateRequest) SetCustomerCorpId(v string) *PreCheckTemplateRequest {
	s.CustomerCorpId = &v
	return s
}

type PreCheckTemplateResponseBody struct {
	Result  *PreCheckTemplateResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                               `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PreCheckTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PreCheckTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *PreCheckTemplateResponseBody) SetResult(v *PreCheckTemplateResponseBodyResult) *PreCheckTemplateResponseBody {
	s.Result = v
	return s
}

func (s *PreCheckTemplateResponseBody) SetSuccess(v bool) *PreCheckTemplateResponseBody {
	s.Success = &v
	return s
}

type PreCheckTemplateResponseBodyResult struct {
	BlockRecords []*PreCheckTemplateResponseBodyResultBlockRecords `json:"blockRecords,omitempty" xml:"blockRecords,omitempty" type:"Repeated"`
	Pass         *bool                                             `json:"pass,omitempty" xml:"pass,omitempty"`
}

func (s PreCheckTemplateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PreCheckTemplateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PreCheckTemplateResponseBodyResult) SetBlockRecords(v []*PreCheckTemplateResponseBodyResultBlockRecords) *PreCheckTemplateResponseBodyResult {
	s.BlockRecords = v
	return s
}

func (s *PreCheckTemplateResponseBodyResult) SetPass(v bool) *PreCheckTemplateResponseBodyResult {
	s.Pass = &v
	return s
}

type PreCheckTemplateResponseBodyResultBlockRecords struct {
	BlockType *string `json:"blockType,omitempty" xml:"blockType,omitempty"`
	Reason    *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s PreCheckTemplateResponseBodyResultBlockRecords) String() string {
	return tea.Prettify(s)
}

func (s PreCheckTemplateResponseBodyResultBlockRecords) GoString() string {
	return s.String()
}

func (s *PreCheckTemplateResponseBodyResultBlockRecords) SetBlockType(v string) *PreCheckTemplateResponseBodyResultBlockRecords {
	s.BlockType = &v
	return s
}

func (s *PreCheckTemplateResponseBodyResultBlockRecords) SetReason(v string) *PreCheckTemplateResponseBodyResultBlockRecords {
	s.Reason = &v
	return s
}

type PreCheckTemplateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreCheckTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreCheckTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s PreCheckTemplateResponse) GoString() string {
	return s.String()
}

func (s *PreCheckTemplateResponse) SetHeaders(v map[string]*string) *PreCheckTemplateResponse {
	s.Headers = v
	return s
}

func (s *PreCheckTemplateResponse) SetStatusCode(v int32) *PreCheckTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *PreCheckTemplateResponse) SetBody(v *PreCheckTemplateResponseBody) *PreCheckTemplateResponse {
	s.Body = v
	return s
}

type QueryTripProcessTemplatesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryTripProcessTemplatesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryTripProcessTemplatesHeaders) GoString() string {
	return s.String()
}

func (s *QueryTripProcessTemplatesHeaders) SetCommonHeaders(v map[string]*string) *QueryTripProcessTemplatesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryTripProcessTemplatesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryTripProcessTemplatesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryTripProcessTemplatesRequest struct {
	CustomerCorpId *string `json:"customerCorpId,omitempty" xml:"customerCorpId,omitempty"`
}

func (s QueryTripProcessTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTripProcessTemplatesRequest) GoString() string {
	return s.String()
}

func (s *QueryTripProcessTemplatesRequest) SetCustomerCorpId(v string) *QueryTripProcessTemplatesRequest {
	s.CustomerCorpId = &v
	return s
}

type QueryTripProcessTemplatesResponseBody struct {
	Result  *QueryTripProcessTemplatesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                        `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryTripProcessTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTripProcessTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTripProcessTemplatesResponseBody) SetResult(v *QueryTripProcessTemplatesResponseBodyResult) *QueryTripProcessTemplatesResponseBody {
	s.Result = v
	return s
}

func (s *QueryTripProcessTemplatesResponseBody) SetSuccess(v bool) *QueryTripProcessTemplatesResponseBody {
	s.Success = &v
	return s
}

type QueryTripProcessTemplatesResponseBodyResult struct {
	Schemas []*QueryTripProcessTemplatesResponseBodyResultSchemas `json:"schemas,omitempty" xml:"schemas,omitempty" type:"Repeated"`
}

func (s QueryTripProcessTemplatesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryTripProcessTemplatesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryTripProcessTemplatesResponseBodyResult) SetSchemas(v []*QueryTripProcessTemplatesResponseBodyResultSchemas) *QueryTripProcessTemplatesResponseBodyResult {
	s.Schemas = v
	return s
}

type QueryTripProcessTemplatesResponseBodyResultSchemas struct {
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	ProcessName *string `json:"processName,omitempty" xml:"processName,omitempty"`
	Type        *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryTripProcessTemplatesResponseBodyResultSchemas) String() string {
	return tea.Prettify(s)
}

func (s QueryTripProcessTemplatesResponseBodyResultSchemas) GoString() string {
	return s.String()
}

func (s *QueryTripProcessTemplatesResponseBodyResultSchemas) SetProcessCode(v string) *QueryTripProcessTemplatesResponseBodyResultSchemas {
	s.ProcessCode = &v
	return s
}

func (s *QueryTripProcessTemplatesResponseBodyResultSchemas) SetProcessName(v string) *QueryTripProcessTemplatesResponseBodyResultSchemas {
	s.ProcessName = &v
	return s
}

func (s *QueryTripProcessTemplatesResponseBodyResultSchemas) SetType(v string) *QueryTripProcessTemplatesResponseBodyResultSchemas {
	s.Type = &v
	return s
}

type QueryTripProcessTemplatesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTripProcessTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTripProcessTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTripProcessTemplatesResponse) GoString() string {
	return s.String()
}

func (s *QueryTripProcessTemplatesResponse) SetHeaders(v map[string]*string) *QueryTripProcessTemplatesResponse {
	s.Headers = v
	return s
}

func (s *QueryTripProcessTemplatesResponse) SetStatusCode(v int32) *QueryTripProcessTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTripProcessTemplatesResponse) SetBody(v *QueryTripProcessTemplatesResponseBody) *QueryTripProcessTemplatesResponse {
	s.Body = v
	return s
}

type SyncBusinessSignInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SyncBusinessSignInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncBusinessSignInfoHeaders) GoString() string {
	return s.String()
}

func (s *SyncBusinessSignInfoHeaders) SetCommonHeaders(v map[string]*string) *SyncBusinessSignInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncBusinessSignInfoHeaders) SetXAcsDingtalkAccessToken(v string) *SyncBusinessSignInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SyncBusinessSignInfoRequest struct {
	BizTypeList          []*string                                          `json:"bizTypeList,omitempty" xml:"bizTypeList,omitempty" type:"Repeated"`
	GmtOrgPay            *string                                            `json:"gmtOrgPay,omitempty" xml:"gmtOrgPay,omitempty"`
	GmtSign              *string                                            `json:"gmtSign,omitempty" xml:"gmtSign,omitempty"`
	OrgPayStatus         *string                                            `json:"orgPayStatus,omitempty" xml:"orgPayStatus,omitempty"`
	SignStatus           *string                                            `json:"signStatus,omitempty" xml:"signStatus,omitempty"`
	TargetCorpId         *string                                            `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
	TmcProductDetailList []*SyncBusinessSignInfoRequestTmcProductDetailList `json:"tmcProductDetailList,omitempty" xml:"tmcProductDetailList,omitempty" type:"Repeated"`
	TmcProductList       []*SyncBusinessSignInfoRequestTmcProductList       `json:"tmcProductList,omitempty" xml:"tmcProductList,omitempty" type:"Repeated"`
}

func (s SyncBusinessSignInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncBusinessSignInfoRequest) GoString() string {
	return s.String()
}

func (s *SyncBusinessSignInfoRequest) SetBizTypeList(v []*string) *SyncBusinessSignInfoRequest {
	s.BizTypeList = v
	return s
}

func (s *SyncBusinessSignInfoRequest) SetGmtOrgPay(v string) *SyncBusinessSignInfoRequest {
	s.GmtOrgPay = &v
	return s
}

func (s *SyncBusinessSignInfoRequest) SetGmtSign(v string) *SyncBusinessSignInfoRequest {
	s.GmtSign = &v
	return s
}

func (s *SyncBusinessSignInfoRequest) SetOrgPayStatus(v string) *SyncBusinessSignInfoRequest {
	s.OrgPayStatus = &v
	return s
}

func (s *SyncBusinessSignInfoRequest) SetSignStatus(v string) *SyncBusinessSignInfoRequest {
	s.SignStatus = &v
	return s
}

func (s *SyncBusinessSignInfoRequest) SetTargetCorpId(v string) *SyncBusinessSignInfoRequest {
	s.TargetCorpId = &v
	return s
}

func (s *SyncBusinessSignInfoRequest) SetTmcProductDetailList(v []*SyncBusinessSignInfoRequestTmcProductDetailList) *SyncBusinessSignInfoRequest {
	s.TmcProductDetailList = v
	return s
}

func (s *SyncBusinessSignInfoRequest) SetTmcProductList(v []*SyncBusinessSignInfoRequestTmcProductList) *SyncBusinessSignInfoRequest {
	s.TmcProductList = v
	return s
}

type SyncBusinessSignInfoRequestTmcProductDetailList struct {
	GmtOrgPay *string `json:"gmtOrgPay,omitempty" xml:"gmtOrgPay,omitempty"`
	PayType   *string `json:"payType,omitempty" xml:"payType,omitempty"`
	Product   *string `json:"product,omitempty" xml:"product,omitempty"`
}

func (s SyncBusinessSignInfoRequestTmcProductDetailList) String() string {
	return tea.Prettify(s)
}

func (s SyncBusinessSignInfoRequestTmcProductDetailList) GoString() string {
	return s.String()
}

func (s *SyncBusinessSignInfoRequestTmcProductDetailList) SetGmtOrgPay(v string) *SyncBusinessSignInfoRequestTmcProductDetailList {
	s.GmtOrgPay = &v
	return s
}

func (s *SyncBusinessSignInfoRequestTmcProductDetailList) SetPayType(v string) *SyncBusinessSignInfoRequestTmcProductDetailList {
	s.PayType = &v
	return s
}

func (s *SyncBusinessSignInfoRequestTmcProductDetailList) SetProduct(v string) *SyncBusinessSignInfoRequestTmcProductDetailList {
	s.Product = &v
	return s
}

type SyncBusinessSignInfoRequestTmcProductList struct {
	ProductDetailList []*SyncBusinessSignInfoRequestTmcProductListProductDetailList `json:"productDetailList,omitempty" xml:"productDetailList,omitempty" type:"Repeated"`
	TmcCorpId         *string                                                       `json:"tmcCorpId,omitempty" xml:"tmcCorpId,omitempty"`
}

func (s SyncBusinessSignInfoRequestTmcProductList) String() string {
	return tea.Prettify(s)
}

func (s SyncBusinessSignInfoRequestTmcProductList) GoString() string {
	return s.String()
}

func (s *SyncBusinessSignInfoRequestTmcProductList) SetProductDetailList(v []*SyncBusinessSignInfoRequestTmcProductListProductDetailList) *SyncBusinessSignInfoRequestTmcProductList {
	s.ProductDetailList = v
	return s
}

func (s *SyncBusinessSignInfoRequestTmcProductList) SetTmcCorpId(v string) *SyncBusinessSignInfoRequestTmcProductList {
	s.TmcCorpId = &v
	return s
}

type SyncBusinessSignInfoRequestTmcProductListProductDetailList struct {
	CategoryType *string `json:"categoryType,omitempty" xml:"categoryType,omitempty"`
	GmtOrgPay    *string `json:"gmtOrgPay,omitempty" xml:"gmtOrgPay,omitempty"`
	OpenStatus   *bool   `json:"openStatus,omitempty" xml:"openStatus,omitempty"`
	PayType      *string `json:"payType,omitempty" xml:"payType,omitempty"`
	Product      *string `json:"product,omitempty" xml:"product,omitempty"`
}

func (s SyncBusinessSignInfoRequestTmcProductListProductDetailList) String() string {
	return tea.Prettify(s)
}

func (s SyncBusinessSignInfoRequestTmcProductListProductDetailList) GoString() string {
	return s.String()
}

func (s *SyncBusinessSignInfoRequestTmcProductListProductDetailList) SetCategoryType(v string) *SyncBusinessSignInfoRequestTmcProductListProductDetailList {
	s.CategoryType = &v
	return s
}

func (s *SyncBusinessSignInfoRequestTmcProductListProductDetailList) SetGmtOrgPay(v string) *SyncBusinessSignInfoRequestTmcProductListProductDetailList {
	s.GmtOrgPay = &v
	return s
}

func (s *SyncBusinessSignInfoRequestTmcProductListProductDetailList) SetOpenStatus(v bool) *SyncBusinessSignInfoRequestTmcProductListProductDetailList {
	s.OpenStatus = &v
	return s
}

func (s *SyncBusinessSignInfoRequestTmcProductListProductDetailList) SetPayType(v string) *SyncBusinessSignInfoRequestTmcProductListProductDetailList {
	s.PayType = &v
	return s
}

func (s *SyncBusinessSignInfoRequestTmcProductListProductDetailList) SetProduct(v string) *SyncBusinessSignInfoRequestTmcProductListProductDetailList {
	s.Product = &v
	return s
}

type SyncBusinessSignInfoResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SyncBusinessSignInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncBusinessSignInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SyncBusinessSignInfoResponseBody) SetRequestId(v string) *SyncBusinessSignInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncBusinessSignInfoResponseBody) SetSuccess(v bool) *SyncBusinessSignInfoResponseBody {
	s.Success = &v
	return s
}

type SyncBusinessSignInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncBusinessSignInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncBusinessSignInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncBusinessSignInfoResponse) GoString() string {
	return s.String()
}

func (s *SyncBusinessSignInfoResponse) SetHeaders(v map[string]*string) *SyncBusinessSignInfoResponse {
	s.Headers = v
	return s
}

func (s *SyncBusinessSignInfoResponse) SetStatusCode(v int32) *SyncBusinessSignInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncBusinessSignInfoResponse) SetBody(v *SyncBusinessSignInfoResponseBody) *SyncBusinessSignInfoResponse {
	s.Body = v
	return s
}

type SyncCostCenterHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SyncCostCenterHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncCostCenterHeaders) GoString() string {
	return s.String()
}

func (s *SyncCostCenterHeaders) SetCommonHeaders(v map[string]*string) *SyncCostCenterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncCostCenterHeaders) SetXAcsDingtalkAccessToken(v string) *SyncCostCenterHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SyncCostCenterRequest struct {
	ChannelCorpId *string `json:"channelCorpId,omitempty" xml:"channelCorpId,omitempty"`
	CostCenterId  *string `json:"costCenterId,omitempty" xml:"costCenterId,omitempty"`
	DeleteFlag    *bool   `json:"deleteFlag,omitempty" xml:"deleteFlag,omitempty"`
	Extension     *string `json:"extension,omitempty" xml:"extension,omitempty"`
	GmtAction     *string `json:"gmtAction,omitempty" xml:"gmtAction,omitempty"`
	Number        *string `json:"number,omitempty" xml:"number,omitempty"`
	Scope         *int32  `json:"scope,omitempty" xml:"scope,omitempty"`
	Source        *string `json:"source,omitempty" xml:"source,omitempty"`
	ThirdPartId   *string `json:"thirdPartId,omitempty" xml:"thirdPartId,omitempty"`
	Title         *string `json:"title,omitempty" xml:"title,omitempty"`
	UserId        *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SyncCostCenterRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncCostCenterRequest) GoString() string {
	return s.String()
}

func (s *SyncCostCenterRequest) SetChannelCorpId(v string) *SyncCostCenterRequest {
	s.ChannelCorpId = &v
	return s
}

func (s *SyncCostCenterRequest) SetCostCenterId(v string) *SyncCostCenterRequest {
	s.CostCenterId = &v
	return s
}

func (s *SyncCostCenterRequest) SetDeleteFlag(v bool) *SyncCostCenterRequest {
	s.DeleteFlag = &v
	return s
}

func (s *SyncCostCenterRequest) SetExtension(v string) *SyncCostCenterRequest {
	s.Extension = &v
	return s
}

func (s *SyncCostCenterRequest) SetGmtAction(v string) *SyncCostCenterRequest {
	s.GmtAction = &v
	return s
}

func (s *SyncCostCenterRequest) SetNumber(v string) *SyncCostCenterRequest {
	s.Number = &v
	return s
}

func (s *SyncCostCenterRequest) SetScope(v int32) *SyncCostCenterRequest {
	s.Scope = &v
	return s
}

func (s *SyncCostCenterRequest) SetSource(v string) *SyncCostCenterRequest {
	s.Source = &v
	return s
}

func (s *SyncCostCenterRequest) SetThirdPartId(v string) *SyncCostCenterRequest {
	s.ThirdPartId = &v
	return s
}

func (s *SyncCostCenterRequest) SetTitle(v string) *SyncCostCenterRequest {
	s.Title = &v
	return s
}

func (s *SyncCostCenterRequest) SetUserId(v string) *SyncCostCenterRequest {
	s.UserId = &v
	return s
}

type SyncCostCenterResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SyncCostCenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncCostCenterResponseBody) GoString() string {
	return s.String()
}

func (s *SyncCostCenterResponseBody) SetSuccess(v bool) *SyncCostCenterResponseBody {
	s.Success = &v
	return s
}

type SyncCostCenterResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncCostCenterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncCostCenterResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncCostCenterResponse) GoString() string {
	return s.String()
}

func (s *SyncCostCenterResponse) SetHeaders(v map[string]*string) *SyncCostCenterResponse {
	s.Headers = v
	return s
}

func (s *SyncCostCenterResponse) SetStatusCode(v int32) *SyncCostCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncCostCenterResponse) SetBody(v *SyncCostCenterResponseBody) *SyncCostCenterResponse {
	s.Body = v
	return s
}

type SyncCostCenterEntityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SyncCostCenterEntityHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncCostCenterEntityHeaders) GoString() string {
	return s.String()
}

func (s *SyncCostCenterEntityHeaders) SetCommonHeaders(v map[string]*string) *SyncCostCenterEntityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncCostCenterEntityHeaders) SetXAcsDingtalkAccessToken(v string) *SyncCostCenterEntityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SyncCostCenterEntityRequest struct {
	ChannelCorpId *string                                  `json:"channelCorpId,omitempty" xml:"channelCorpId,omitempty"`
	CostCenterId  *string                                  `json:"costCenterId,omitempty" xml:"costCenterId,omitempty"`
	DelAll        *bool                                    `json:"delAll,omitempty" xml:"delAll,omitempty"`
	EntityList    []*SyncCostCenterEntityRequestEntityList `json:"entityList,omitempty" xml:"entityList,omitempty" type:"Repeated"`
	UserId        *string                                  `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SyncCostCenterEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncCostCenterEntityRequest) GoString() string {
	return s.String()
}

func (s *SyncCostCenterEntityRequest) SetChannelCorpId(v string) *SyncCostCenterEntityRequest {
	s.ChannelCorpId = &v
	return s
}

func (s *SyncCostCenterEntityRequest) SetCostCenterId(v string) *SyncCostCenterEntityRequest {
	s.CostCenterId = &v
	return s
}

func (s *SyncCostCenterEntityRequest) SetDelAll(v bool) *SyncCostCenterEntityRequest {
	s.DelAll = &v
	return s
}

func (s *SyncCostCenterEntityRequest) SetEntityList(v []*SyncCostCenterEntityRequestEntityList) *SyncCostCenterEntityRequest {
	s.EntityList = v
	return s
}

func (s *SyncCostCenterEntityRequest) SetUserId(v string) *SyncCostCenterEntityRequest {
	s.UserId = &v
	return s
}

type SyncCostCenterEntityRequestEntityList struct {
	EntityId   *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	EntityType *string `json:"entityType,omitempty" xml:"entityType,omitempty"`
}

func (s SyncCostCenterEntityRequestEntityList) String() string {
	return tea.Prettify(s)
}

func (s SyncCostCenterEntityRequestEntityList) GoString() string {
	return s.String()
}

func (s *SyncCostCenterEntityRequestEntityList) SetEntityId(v string) *SyncCostCenterEntityRequestEntityList {
	s.EntityId = &v
	return s
}

func (s *SyncCostCenterEntityRequestEntityList) SetEntityType(v string) *SyncCostCenterEntityRequestEntityList {
	s.EntityType = &v
	return s
}

type SyncCostCenterEntityResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SyncCostCenterEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncCostCenterEntityResponseBody) GoString() string {
	return s.String()
}

func (s *SyncCostCenterEntityResponseBody) SetSuccess(v bool) *SyncCostCenterEntityResponseBody {
	s.Success = &v
	return s
}

type SyncCostCenterEntityResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncCostCenterEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncCostCenterEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncCostCenterEntityResponse) GoString() string {
	return s.String()
}

func (s *SyncCostCenterEntityResponse) SetHeaders(v map[string]*string) *SyncCostCenterEntityResponse {
	s.Headers = v
	return s
}

func (s *SyncCostCenterEntityResponse) SetStatusCode(v int32) *SyncCostCenterEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncCostCenterEntityResponse) SetBody(v *SyncCostCenterEntityResponseBody) *SyncCostCenterEntityResponse {
	s.Body = v
	return s
}

type SyncInvoiceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SyncInvoiceHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncInvoiceHeaders) GoString() string {
	return s.String()
}

func (s *SyncInvoiceHeaders) SetCommonHeaders(v map[string]*string) *SyncInvoiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncInvoiceHeaders) SetXAcsDingtalkAccessToken(v string) *SyncInvoiceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SyncInvoiceRequest struct {
	Address       *string   `json:"address,omitempty" xml:"address,omitempty"`
	BankName      *string   `json:"bankName,omitempty" xml:"bankName,omitempty"`
	BankNo        *string   `json:"bankNo,omitempty" xml:"bankNo,omitempty"`
	ChannelCorpId *string   `json:"channelCorpId,omitempty" xml:"channelCorpId,omitempty"`
	DeleteFlag    *bool     `json:"deleteFlag,omitempty" xml:"deleteFlag,omitempty"`
	GmtAction     *string   `json:"gmtAction,omitempty" xml:"gmtAction,omitempty"`
	InvoiceId     *string   `json:"invoiceId,omitempty" xml:"invoiceId,omitempty"`
	ProjectIds    []*string `json:"projectIds,omitempty" xml:"projectIds,omitempty" type:"Repeated"`
	Scope         *int32    `json:"scope,omitempty" xml:"scope,omitempty"`
	Source        *string   `json:"source,omitempty" xml:"source,omitempty"`
	TaxNo         *string   `json:"taxNo,omitempty" xml:"taxNo,omitempty"`
	Tel           *string   `json:"tel,omitempty" xml:"tel,omitempty"`
	ThirdPartId   *string   `json:"thirdPartId,omitempty" xml:"thirdPartId,omitempty"`
	Title         *string   `json:"title,omitempty" xml:"title,omitempty"`
	Type          *int32    `json:"type,omitempty" xml:"type,omitempty"`
	UnitType      *int32    `json:"unitType,omitempty" xml:"unitType,omitempty"`
	UserId        *string   `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SyncInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncInvoiceRequest) GoString() string {
	return s.String()
}

func (s *SyncInvoiceRequest) SetAddress(v string) *SyncInvoiceRequest {
	s.Address = &v
	return s
}

func (s *SyncInvoiceRequest) SetBankName(v string) *SyncInvoiceRequest {
	s.BankName = &v
	return s
}

func (s *SyncInvoiceRequest) SetBankNo(v string) *SyncInvoiceRequest {
	s.BankNo = &v
	return s
}

func (s *SyncInvoiceRequest) SetChannelCorpId(v string) *SyncInvoiceRequest {
	s.ChannelCorpId = &v
	return s
}

func (s *SyncInvoiceRequest) SetDeleteFlag(v bool) *SyncInvoiceRequest {
	s.DeleteFlag = &v
	return s
}

func (s *SyncInvoiceRequest) SetGmtAction(v string) *SyncInvoiceRequest {
	s.GmtAction = &v
	return s
}

func (s *SyncInvoiceRequest) SetInvoiceId(v string) *SyncInvoiceRequest {
	s.InvoiceId = &v
	return s
}

func (s *SyncInvoiceRequest) SetProjectIds(v []*string) *SyncInvoiceRequest {
	s.ProjectIds = v
	return s
}

func (s *SyncInvoiceRequest) SetScope(v int32) *SyncInvoiceRequest {
	s.Scope = &v
	return s
}

func (s *SyncInvoiceRequest) SetSource(v string) *SyncInvoiceRequest {
	s.Source = &v
	return s
}

func (s *SyncInvoiceRequest) SetTaxNo(v string) *SyncInvoiceRequest {
	s.TaxNo = &v
	return s
}

func (s *SyncInvoiceRequest) SetTel(v string) *SyncInvoiceRequest {
	s.Tel = &v
	return s
}

func (s *SyncInvoiceRequest) SetThirdPartId(v string) *SyncInvoiceRequest {
	s.ThirdPartId = &v
	return s
}

func (s *SyncInvoiceRequest) SetTitle(v string) *SyncInvoiceRequest {
	s.Title = &v
	return s
}

func (s *SyncInvoiceRequest) SetType(v int32) *SyncInvoiceRequest {
	s.Type = &v
	return s
}

func (s *SyncInvoiceRequest) SetUnitType(v int32) *SyncInvoiceRequest {
	s.UnitType = &v
	return s
}

func (s *SyncInvoiceRequest) SetUserId(v string) *SyncInvoiceRequest {
	s.UserId = &v
	return s
}

type SyncInvoiceResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SyncInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *SyncInvoiceResponseBody) SetSuccess(v bool) *SyncInvoiceResponseBody {
	s.Success = &v
	return s
}

type SyncInvoiceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncInvoiceResponse) GoString() string {
	return s.String()
}

func (s *SyncInvoiceResponse) SetHeaders(v map[string]*string) *SyncInvoiceResponse {
	s.Headers = v
	return s
}

func (s *SyncInvoiceResponse) SetStatusCode(v int32) *SyncInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncInvoiceResponse) SetBody(v *SyncInvoiceResponseBody) *SyncInvoiceResponse {
	s.Body = v
	return s
}

type SyncInvoiceEntityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SyncInvoiceEntityHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncInvoiceEntityHeaders) GoString() string {
	return s.String()
}

func (s *SyncInvoiceEntityHeaders) SetCommonHeaders(v map[string]*string) *SyncInvoiceEntityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncInvoiceEntityHeaders) SetXAcsDingtalkAccessToken(v string) *SyncInvoiceEntityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SyncInvoiceEntityRequest struct {
	ChannelCorpId *string                               `json:"channelCorpId,omitempty" xml:"channelCorpId,omitempty"`
	DelAll        *bool                                 `json:"delAll,omitempty" xml:"delAll,omitempty"`
	EntityList    []*SyncInvoiceEntityRequestEntityList `json:"entityList,omitempty" xml:"entityList,omitempty" type:"Repeated"`
	InvoiceId     *string                               `json:"invoiceId,omitempty" xml:"invoiceId,omitempty"`
	UserId        *string                               `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SyncInvoiceEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncInvoiceEntityRequest) GoString() string {
	return s.String()
}

func (s *SyncInvoiceEntityRequest) SetChannelCorpId(v string) *SyncInvoiceEntityRequest {
	s.ChannelCorpId = &v
	return s
}

func (s *SyncInvoiceEntityRequest) SetDelAll(v bool) *SyncInvoiceEntityRequest {
	s.DelAll = &v
	return s
}

func (s *SyncInvoiceEntityRequest) SetEntityList(v []*SyncInvoiceEntityRequestEntityList) *SyncInvoiceEntityRequest {
	s.EntityList = v
	return s
}

func (s *SyncInvoiceEntityRequest) SetInvoiceId(v string) *SyncInvoiceEntityRequest {
	s.InvoiceId = &v
	return s
}

func (s *SyncInvoiceEntityRequest) SetUserId(v string) *SyncInvoiceEntityRequest {
	s.UserId = &v
	return s
}

type SyncInvoiceEntityRequestEntityList struct {
	EntityId   *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	EntityType *string `json:"entityType,omitempty" xml:"entityType,omitempty"`
}

func (s SyncInvoiceEntityRequestEntityList) String() string {
	return tea.Prettify(s)
}

func (s SyncInvoiceEntityRequestEntityList) GoString() string {
	return s.String()
}

func (s *SyncInvoiceEntityRequestEntityList) SetEntityId(v string) *SyncInvoiceEntityRequestEntityList {
	s.EntityId = &v
	return s
}

func (s *SyncInvoiceEntityRequestEntityList) SetEntityType(v string) *SyncInvoiceEntityRequestEntityList {
	s.EntityType = &v
	return s
}

type SyncInvoiceEntityResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SyncInvoiceEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncInvoiceEntityResponseBody) GoString() string {
	return s.String()
}

func (s *SyncInvoiceEntityResponseBody) SetSuccess(v bool) *SyncInvoiceEntityResponseBody {
	s.Success = &v
	return s
}

type SyncInvoiceEntityResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncInvoiceEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncInvoiceEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncInvoiceEntityResponse) GoString() string {
	return s.String()
}

func (s *SyncInvoiceEntityResponse) SetHeaders(v map[string]*string) *SyncInvoiceEntityResponse {
	s.Headers = v
	return s
}

func (s *SyncInvoiceEntityResponse) SetStatusCode(v int32) *SyncInvoiceEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncInvoiceEntityResponse) SetBody(v *SyncInvoiceEntityResponseBody) *SyncInvoiceEntityResponse {
	s.Body = v
	return s
}

type SyncProjectHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SyncProjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncProjectHeaders) GoString() string {
	return s.String()
}

func (s *SyncProjectHeaders) SetCommonHeaders(v map[string]*string) *SyncProjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncProjectHeaders) SetXAcsDingtalkAccessToken(v string) *SyncProjectHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SyncProjectRequest struct {
	ChannelCorpId *string   `json:"channelCorpId,omitempty" xml:"channelCorpId,omitempty"`
	Code          *string   `json:"code,omitempty" xml:"code,omitempty"`
	CostCenterId  *string   `json:"costCenterId,omitempty" xml:"costCenterId,omitempty"`
	DeleteFlag    *bool     `json:"deleteFlag,omitempty" xml:"deleteFlag,omitempty"`
	Extension     *string   `json:"extension,omitempty" xml:"extension,omitempty"`
	GmtAction     *string   `json:"gmtAction,omitempty" xml:"gmtAction,omitempty"`
	InvoiceId     *string   `json:"invoiceId,omitempty" xml:"invoiceId,omitempty"`
	ManagerIds    []*string `json:"managerIds,omitempty" xml:"managerIds,omitempty" type:"Repeated"`
	ProjectId     *string   `json:"projectId,omitempty" xml:"projectId,omitempty"`
	ProjectName   *string   `json:"projectName,omitempty" xml:"projectName,omitempty"`
	Scope         *int32    `json:"scope,omitempty" xml:"scope,omitempty"`
	Source        *string   `json:"source,omitempty" xml:"source,omitempty"`
	ThirdPartId   *string   `json:"thirdPartId,omitempty" xml:"thirdPartId,omitempty"`
	UserId        *string   `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SyncProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncProjectRequest) GoString() string {
	return s.String()
}

func (s *SyncProjectRequest) SetChannelCorpId(v string) *SyncProjectRequest {
	s.ChannelCorpId = &v
	return s
}

func (s *SyncProjectRequest) SetCode(v string) *SyncProjectRequest {
	s.Code = &v
	return s
}

func (s *SyncProjectRequest) SetCostCenterId(v string) *SyncProjectRequest {
	s.CostCenterId = &v
	return s
}

func (s *SyncProjectRequest) SetDeleteFlag(v bool) *SyncProjectRequest {
	s.DeleteFlag = &v
	return s
}

func (s *SyncProjectRequest) SetExtension(v string) *SyncProjectRequest {
	s.Extension = &v
	return s
}

func (s *SyncProjectRequest) SetGmtAction(v string) *SyncProjectRequest {
	s.GmtAction = &v
	return s
}

func (s *SyncProjectRequest) SetInvoiceId(v string) *SyncProjectRequest {
	s.InvoiceId = &v
	return s
}

func (s *SyncProjectRequest) SetManagerIds(v []*string) *SyncProjectRequest {
	s.ManagerIds = v
	return s
}

func (s *SyncProjectRequest) SetProjectId(v string) *SyncProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *SyncProjectRequest) SetProjectName(v string) *SyncProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *SyncProjectRequest) SetScope(v int32) *SyncProjectRequest {
	s.Scope = &v
	return s
}

func (s *SyncProjectRequest) SetSource(v string) *SyncProjectRequest {
	s.Source = &v
	return s
}

func (s *SyncProjectRequest) SetThirdPartId(v string) *SyncProjectRequest {
	s.ThirdPartId = &v
	return s
}

func (s *SyncProjectRequest) SetUserId(v string) *SyncProjectRequest {
	s.UserId = &v
	return s
}

type SyncProjectResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SyncProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncProjectResponseBody) GoString() string {
	return s.String()
}

func (s *SyncProjectResponseBody) SetSuccess(v bool) *SyncProjectResponseBody {
	s.Success = &v
	return s
}

type SyncProjectResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncProjectResponse) GoString() string {
	return s.String()
}

func (s *SyncProjectResponse) SetHeaders(v map[string]*string) *SyncProjectResponse {
	s.Headers = v
	return s
}

func (s *SyncProjectResponse) SetStatusCode(v int32) *SyncProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncProjectResponse) SetBody(v *SyncProjectResponseBody) *SyncProjectResponse {
	s.Body = v
	return s
}

type SyncProjectEntityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SyncProjectEntityHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncProjectEntityHeaders) GoString() string {
	return s.String()
}

func (s *SyncProjectEntityHeaders) SetCommonHeaders(v map[string]*string) *SyncProjectEntityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncProjectEntityHeaders) SetXAcsDingtalkAccessToken(v string) *SyncProjectEntityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SyncProjectEntityRequest struct {
	ChannelCorpId *string                               `json:"channelCorpId,omitempty" xml:"channelCorpId,omitempty"`
	DelAll        *bool                                 `json:"delAll,omitempty" xml:"delAll,omitempty"`
	EntityList    []*SyncProjectEntityRequestEntityList `json:"entityList,omitempty" xml:"entityList,omitempty" type:"Repeated"`
	ProjectId     *string                               `json:"projectId,omitempty" xml:"projectId,omitempty"`
	UserId        *string                               `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SyncProjectEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncProjectEntityRequest) GoString() string {
	return s.String()
}

func (s *SyncProjectEntityRequest) SetChannelCorpId(v string) *SyncProjectEntityRequest {
	s.ChannelCorpId = &v
	return s
}

func (s *SyncProjectEntityRequest) SetDelAll(v bool) *SyncProjectEntityRequest {
	s.DelAll = &v
	return s
}

func (s *SyncProjectEntityRequest) SetEntityList(v []*SyncProjectEntityRequestEntityList) *SyncProjectEntityRequest {
	s.EntityList = v
	return s
}

func (s *SyncProjectEntityRequest) SetProjectId(v string) *SyncProjectEntityRequest {
	s.ProjectId = &v
	return s
}

func (s *SyncProjectEntityRequest) SetUserId(v string) *SyncProjectEntityRequest {
	s.UserId = &v
	return s
}

type SyncProjectEntityRequestEntityList struct {
	EntityId   *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	EntityType *string `json:"entityType,omitempty" xml:"entityType,omitempty"`
}

func (s SyncProjectEntityRequestEntityList) String() string {
	return tea.Prettify(s)
}

func (s SyncProjectEntityRequestEntityList) GoString() string {
	return s.String()
}

func (s *SyncProjectEntityRequestEntityList) SetEntityId(v string) *SyncProjectEntityRequestEntityList {
	s.EntityId = &v
	return s
}

func (s *SyncProjectEntityRequestEntityList) SetEntityType(v string) *SyncProjectEntityRequestEntityList {
	s.EntityType = &v
	return s
}

type SyncProjectEntityResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SyncProjectEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncProjectEntityResponseBody) GoString() string {
	return s.String()
}

func (s *SyncProjectEntityResponseBody) SetSuccess(v bool) *SyncProjectEntityResponseBody {
	s.Success = &v
	return s
}

type SyncProjectEntityResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncProjectEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncProjectEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncProjectEntityResponse) GoString() string {
	return s.String()
}

func (s *SyncProjectEntityResponse) SetHeaders(v map[string]*string) *SyncProjectEntityResponse {
	s.Headers = v
	return s
}

func (s *SyncProjectEntityResponse) SetStatusCode(v int32) *SyncProjectEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncProjectEntityResponse) SetBody(v *SyncProjectEntityResponseBody) *SyncProjectEntityResponse {
	s.Body = v
	return s
}

type SyncSecretKeyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SyncSecretKeyHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncSecretKeyHeaders) GoString() string {
	return s.String()
}

func (s *SyncSecretKeyHeaders) SetCommonHeaders(v map[string]*string) *SyncSecretKeyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncSecretKeyHeaders) SetXAcsDingtalkAccessToken(v string) *SyncSecretKeyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SyncSecretKeyRequest struct {
	ActionType      *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	SecretString    *string `json:"secretString,omitempty" xml:"secretString,omitempty"`
	TargetCorpId    *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
	TripAppKey      *string `json:"tripAppKey,omitempty" xml:"tripAppKey,omitempty"`
	TripAppSecurity *string `json:"tripAppSecurity,omitempty" xml:"tripAppSecurity,omitempty"`
	TripCorpId      *string `json:"tripCorpId,omitempty" xml:"tripCorpId,omitempty"`
}

func (s SyncSecretKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncSecretKeyRequest) GoString() string {
	return s.String()
}

func (s *SyncSecretKeyRequest) SetActionType(v string) *SyncSecretKeyRequest {
	s.ActionType = &v
	return s
}

func (s *SyncSecretKeyRequest) SetSecretString(v string) *SyncSecretKeyRequest {
	s.SecretString = &v
	return s
}

func (s *SyncSecretKeyRequest) SetTargetCorpId(v string) *SyncSecretKeyRequest {
	s.TargetCorpId = &v
	return s
}

func (s *SyncSecretKeyRequest) SetTripAppKey(v string) *SyncSecretKeyRequest {
	s.TripAppKey = &v
	return s
}

func (s *SyncSecretKeyRequest) SetTripAppSecurity(v string) *SyncSecretKeyRequest {
	s.TripAppSecurity = &v
	return s
}

func (s *SyncSecretKeyRequest) SetTripCorpId(v string) *SyncSecretKeyRequest {
	s.TripCorpId = &v
	return s
}

type SyncSecretKeyResponseBody struct {
	Result  *SyncSecretKeyResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *string                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SyncSecretKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncSecretKeyResponseBody) GoString() string {
	return s.String()
}

func (s *SyncSecretKeyResponseBody) SetResult(v *SyncSecretKeyResponseBodyResult) *SyncSecretKeyResponseBody {
	s.Result = v
	return s
}

func (s *SyncSecretKeyResponseBody) SetSuccess(v string) *SyncSecretKeyResponseBody {
	s.Success = &v
	return s
}

type SyncSecretKeyResponseBodyResult struct {
	SecretString    *string `json:"secretString,omitempty" xml:"secretString,omitempty"`
	TargetCorpId    *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
	TripAppKey      *string `json:"tripAppKey,omitempty" xml:"tripAppKey,omitempty"`
	TripAppSecurity *string `json:"tripAppSecurity,omitempty" xml:"tripAppSecurity,omitempty"`
	TripCorpId      *string `json:"tripCorpId,omitempty" xml:"tripCorpId,omitempty"`
}

func (s SyncSecretKeyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SyncSecretKeyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SyncSecretKeyResponseBodyResult) SetSecretString(v string) *SyncSecretKeyResponseBodyResult {
	s.SecretString = &v
	return s
}

func (s *SyncSecretKeyResponseBodyResult) SetTargetCorpId(v string) *SyncSecretKeyResponseBodyResult {
	s.TargetCorpId = &v
	return s
}

func (s *SyncSecretKeyResponseBodyResult) SetTripAppKey(v string) *SyncSecretKeyResponseBodyResult {
	s.TripAppKey = &v
	return s
}

func (s *SyncSecretKeyResponseBodyResult) SetTripAppSecurity(v string) *SyncSecretKeyResponseBodyResult {
	s.TripAppSecurity = &v
	return s
}

func (s *SyncSecretKeyResponseBodyResult) SetTripCorpId(v string) *SyncSecretKeyResponseBodyResult {
	s.TripCorpId = &v
	return s
}

type SyncSecretKeyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncSecretKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncSecretKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncSecretKeyResponse) GoString() string {
	return s.String()
}

func (s *SyncSecretKeyResponse) SetHeaders(v map[string]*string) *SyncSecretKeyResponse {
	s.Headers = v
	return s
}

func (s *SyncSecretKeyResponse) SetStatusCode(v int32) *SyncSecretKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncSecretKeyResponse) SetBody(v *SyncSecretKeyResponseBody) *SyncSecretKeyResponse {
	s.Body = v
	return s
}

type SyncTripOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SyncTripOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncTripOrderHeaders) GoString() string {
	return s.String()
}

func (s *SyncTripOrderHeaders) SetCommonHeaders(v map[string]*string) *SyncTripOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncTripOrderHeaders) SetXAcsDingtalkAccessToken(v string) *SyncTripOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SyncTripOrderRequest struct {
	BizExtension    *string                             `json:"bizExtension,omitempty" xml:"bizExtension,omitempty"`
	ChannelType     *string                             `json:"channelType,omitempty" xml:"channelType,omitempty"`
	Currency        *string                             `json:"currency,omitempty" xml:"currency,omitempty"`
	DingUserId      *string                             `json:"dingUserId,omitempty" xml:"dingUserId,omitempty"`
	DiscountAmount  *string                             `json:"discountAmount,omitempty" xml:"discountAmount,omitempty"`
	EndorseFlag     *bool                               `json:"endorseFlag,omitempty" xml:"endorseFlag,omitempty"`
	Event           *SyncTripOrderRequestEvent          `json:"event,omitempty" xml:"event,omitempty" type:"Struct"`
	GmtOrder        *string                             `json:"gmtOrder,omitempty" xml:"gmtOrder,omitempty"`
	GmtPay          *string                             `json:"gmtPay,omitempty" xml:"gmtPay,omitempty"`
	GmtRefund       *string                             `json:"gmtRefund,omitempty" xml:"gmtRefund,omitempty"`
	InvoiceApplyUrl *string                             `json:"invoiceApplyUrl,omitempty" xml:"invoiceApplyUrl,omitempty"`
	JourneyBizNo    *string                             `json:"journeyBizNo,omitempty" xml:"journeyBizNo,omitempty"`
	OrderDetails    []*SyncTripOrderRequestOrderDetails `json:"orderDetails,omitempty" xml:"orderDetails,omitempty" type:"Repeated"`
	OrderNo         *string                             `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	OrderUrl        *string                             `json:"orderUrl,omitempty" xml:"orderUrl,omitempty"`
	ProcessId       *string                             `json:"processId,omitempty" xml:"processId,omitempty"`
	RealAmount      *string                             `json:"realAmount,omitempty" xml:"realAmount,omitempty"`
	RefundAmount    *string                             `json:"refundAmount,omitempty" xml:"refundAmount,omitempty"`
	RelativeOrderNo *string                             `json:"relativeOrderNo,omitempty" xml:"relativeOrderNo,omitempty"`
	Source          *string                             `json:"source,omitempty" xml:"source,omitempty"`
	SupplyLogo      *string                             `json:"supplyLogo,omitempty" xml:"supplyLogo,omitempty"`
	SupplyName      *string                             `json:"supplyName,omitempty" xml:"supplyName,omitempty"`
	TargetCorpId    *string                             `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
	TmcCorpId       *string                             `json:"tmcCorpId,omitempty" xml:"tmcCorpId,omitempty"`
	TotalAmount     *string                             `json:"totalAmount,omitempty" xml:"totalAmount,omitempty"`
	Type            *string                             `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SyncTripOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncTripOrderRequest) GoString() string {
	return s.String()
}

func (s *SyncTripOrderRequest) SetBizExtension(v string) *SyncTripOrderRequest {
	s.BizExtension = &v
	return s
}

func (s *SyncTripOrderRequest) SetChannelType(v string) *SyncTripOrderRequest {
	s.ChannelType = &v
	return s
}

func (s *SyncTripOrderRequest) SetCurrency(v string) *SyncTripOrderRequest {
	s.Currency = &v
	return s
}

func (s *SyncTripOrderRequest) SetDingUserId(v string) *SyncTripOrderRequest {
	s.DingUserId = &v
	return s
}

func (s *SyncTripOrderRequest) SetDiscountAmount(v string) *SyncTripOrderRequest {
	s.DiscountAmount = &v
	return s
}

func (s *SyncTripOrderRequest) SetEndorseFlag(v bool) *SyncTripOrderRequest {
	s.EndorseFlag = &v
	return s
}

func (s *SyncTripOrderRequest) SetEvent(v *SyncTripOrderRequestEvent) *SyncTripOrderRequest {
	s.Event = v
	return s
}

func (s *SyncTripOrderRequest) SetGmtOrder(v string) *SyncTripOrderRequest {
	s.GmtOrder = &v
	return s
}

func (s *SyncTripOrderRequest) SetGmtPay(v string) *SyncTripOrderRequest {
	s.GmtPay = &v
	return s
}

func (s *SyncTripOrderRequest) SetGmtRefund(v string) *SyncTripOrderRequest {
	s.GmtRefund = &v
	return s
}

func (s *SyncTripOrderRequest) SetInvoiceApplyUrl(v string) *SyncTripOrderRequest {
	s.InvoiceApplyUrl = &v
	return s
}

func (s *SyncTripOrderRequest) SetJourneyBizNo(v string) *SyncTripOrderRequest {
	s.JourneyBizNo = &v
	return s
}

func (s *SyncTripOrderRequest) SetOrderDetails(v []*SyncTripOrderRequestOrderDetails) *SyncTripOrderRequest {
	s.OrderDetails = v
	return s
}

func (s *SyncTripOrderRequest) SetOrderNo(v string) *SyncTripOrderRequest {
	s.OrderNo = &v
	return s
}

func (s *SyncTripOrderRequest) SetOrderUrl(v string) *SyncTripOrderRequest {
	s.OrderUrl = &v
	return s
}

func (s *SyncTripOrderRequest) SetProcessId(v string) *SyncTripOrderRequest {
	s.ProcessId = &v
	return s
}

func (s *SyncTripOrderRequest) SetRealAmount(v string) *SyncTripOrderRequest {
	s.RealAmount = &v
	return s
}

func (s *SyncTripOrderRequest) SetRefundAmount(v string) *SyncTripOrderRequest {
	s.RefundAmount = &v
	return s
}

func (s *SyncTripOrderRequest) SetRelativeOrderNo(v string) *SyncTripOrderRequest {
	s.RelativeOrderNo = &v
	return s
}

func (s *SyncTripOrderRequest) SetSource(v string) *SyncTripOrderRequest {
	s.Source = &v
	return s
}

func (s *SyncTripOrderRequest) SetSupplyLogo(v string) *SyncTripOrderRequest {
	s.SupplyLogo = &v
	return s
}

func (s *SyncTripOrderRequest) SetSupplyName(v string) *SyncTripOrderRequest {
	s.SupplyName = &v
	return s
}

func (s *SyncTripOrderRequest) SetTargetCorpId(v string) *SyncTripOrderRequest {
	s.TargetCorpId = &v
	return s
}

func (s *SyncTripOrderRequest) SetTmcCorpId(v string) *SyncTripOrderRequest {
	s.TmcCorpId = &v
	return s
}

func (s *SyncTripOrderRequest) SetTotalAmount(v string) *SyncTripOrderRequest {
	s.TotalAmount = &v
	return s
}

func (s *SyncTripOrderRequest) SetType(v string) *SyncTripOrderRequest {
	s.Type = &v
	return s
}

type SyncTripOrderRequestEvent struct {
	Action    *string `json:"action,omitempty" xml:"action,omitempty"`
	GmtAction *string `json:"gmtAction,omitempty" xml:"gmtAction,omitempty"`
}

func (s SyncTripOrderRequestEvent) String() string {
	return tea.Prettify(s)
}

func (s SyncTripOrderRequestEvent) GoString() string {
	return s.String()
}

func (s *SyncTripOrderRequestEvent) SetAction(v string) *SyncTripOrderRequestEvent {
	s.Action = &v
	return s
}

func (s *SyncTripOrderRequestEvent) SetGmtAction(v string) *SyncTripOrderRequestEvent {
	s.GmtAction = &v
	return s
}

type SyncTripOrderRequestOrderDetails struct {
	ArrivalTime         *string                                             `json:"arrivalTime,omitempty" xml:"arrivalTime,omitempty"`
	CarColor            *string                                             `json:"carColor,omitempty" xml:"carColor,omitempty"`
	CarModel            *string                                             `json:"carModel,omitempty" xml:"carModel,omitempty"`
	CarNumber           *string                                             `json:"carNumber,omitempty" xml:"carNumber,omitempty"`
	CateringType        *string                                             `json:"cateringType,omitempty" xml:"cateringType,omitempty"`
	CheckInTime         *string                                             `json:"checkInTime,omitempty" xml:"checkInTime,omitempty"`
	CheckOutTime        *string                                             `json:"checkOutTime,omitempty" xml:"checkOutTime,omitempty"`
	DepartTime          *string                                             `json:"departTime,omitempty" xml:"departTime,omitempty"`
	DestinationCity     *string                                             `json:"destinationCity,omitempty" xml:"destinationCity,omitempty"`
	DestinationCityCode *string                                             `json:"destinationCityCode,omitempty" xml:"destinationCityCode,omitempty"`
	DestinationStation  *string                                             `json:"destinationStation,omitempty" xml:"destinationStation,omitempty"`
	DetailAmount        *string                                             `json:"detailAmount,omitempty" xml:"detailAmount,omitempty"`
	HotelAddress        *string                                             `json:"hotelAddress,omitempty" xml:"hotelAddress,omitempty"`
	HotelCity           *string                                             `json:"hotelCity,omitempty" xml:"hotelCity,omitempty"`
	HotelLocation       *SyncTripOrderRequestOrderDetailsHotelLocation      `json:"hotelLocation,omitempty" xml:"hotelLocation,omitempty" type:"Struct"`
	HotelName           *string                                             `json:"hotelName,omitempty" xml:"hotelName,omitempty"`
	OpenConsumerInfo    []*SyncTripOrderRequestOrderDetailsOpenConsumerInfo `json:"openConsumerInfo,omitempty" xml:"openConsumerInfo,omitempty" type:"Repeated"`
	OriginCity          *string                                             `json:"originCity,omitempty" xml:"originCity,omitempty"`
	OriginCityCode      *string                                             `json:"originCityCode,omitempty" xml:"originCityCode,omitempty"`
	OriginStation       *string                                             `json:"originStation,omitempty" xml:"originStation,omitempty"`
	RoomCount           *int32                                              `json:"roomCount,omitempty" xml:"roomCount,omitempty"`
	SeatInfo            *string                                             `json:"seatInfo,omitempty" xml:"seatInfo,omitempty"`
	ServiceType         *string                                             `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	SubSupplyLogo       *string                                             `json:"subSupplyLogo,omitempty" xml:"subSupplyLogo,omitempty"`
	SubSupplyName       *string                                             `json:"subSupplyName,omitempty" xml:"subSupplyName,omitempty"`
	TaxiType            *string                                             `json:"taxiType,omitempty" xml:"taxiType,omitempty"`
	Telephone           *string                                             `json:"telephone,omitempty" xml:"telephone,omitempty"`
	TransportNumber     *string                                             `json:"transportNumber,omitempty" xml:"transportNumber,omitempty"`
	TypeDescription     *string                                             `json:"typeDescription,omitempty" xml:"typeDescription,omitempty"`
}

func (s SyncTripOrderRequestOrderDetails) String() string {
	return tea.Prettify(s)
}

func (s SyncTripOrderRequestOrderDetails) GoString() string {
	return s.String()
}

func (s *SyncTripOrderRequestOrderDetails) SetArrivalTime(v string) *SyncTripOrderRequestOrderDetails {
	s.ArrivalTime = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetCarColor(v string) *SyncTripOrderRequestOrderDetails {
	s.CarColor = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetCarModel(v string) *SyncTripOrderRequestOrderDetails {
	s.CarModel = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetCarNumber(v string) *SyncTripOrderRequestOrderDetails {
	s.CarNumber = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetCateringType(v string) *SyncTripOrderRequestOrderDetails {
	s.CateringType = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetCheckInTime(v string) *SyncTripOrderRequestOrderDetails {
	s.CheckInTime = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetCheckOutTime(v string) *SyncTripOrderRequestOrderDetails {
	s.CheckOutTime = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetDepartTime(v string) *SyncTripOrderRequestOrderDetails {
	s.DepartTime = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetDestinationCity(v string) *SyncTripOrderRequestOrderDetails {
	s.DestinationCity = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetDestinationCityCode(v string) *SyncTripOrderRequestOrderDetails {
	s.DestinationCityCode = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetDestinationStation(v string) *SyncTripOrderRequestOrderDetails {
	s.DestinationStation = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetDetailAmount(v string) *SyncTripOrderRequestOrderDetails {
	s.DetailAmount = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetHotelAddress(v string) *SyncTripOrderRequestOrderDetails {
	s.HotelAddress = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetHotelCity(v string) *SyncTripOrderRequestOrderDetails {
	s.HotelCity = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetHotelLocation(v *SyncTripOrderRequestOrderDetailsHotelLocation) *SyncTripOrderRequestOrderDetails {
	s.HotelLocation = v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetHotelName(v string) *SyncTripOrderRequestOrderDetails {
	s.HotelName = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetOpenConsumerInfo(v []*SyncTripOrderRequestOrderDetailsOpenConsumerInfo) *SyncTripOrderRequestOrderDetails {
	s.OpenConsumerInfo = v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetOriginCity(v string) *SyncTripOrderRequestOrderDetails {
	s.OriginCity = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetOriginCityCode(v string) *SyncTripOrderRequestOrderDetails {
	s.OriginCityCode = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetOriginStation(v string) *SyncTripOrderRequestOrderDetails {
	s.OriginStation = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetRoomCount(v int32) *SyncTripOrderRequestOrderDetails {
	s.RoomCount = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetSeatInfo(v string) *SyncTripOrderRequestOrderDetails {
	s.SeatInfo = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetServiceType(v string) *SyncTripOrderRequestOrderDetails {
	s.ServiceType = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetSubSupplyLogo(v string) *SyncTripOrderRequestOrderDetails {
	s.SubSupplyLogo = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetSubSupplyName(v string) *SyncTripOrderRequestOrderDetails {
	s.SubSupplyName = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetTaxiType(v string) *SyncTripOrderRequestOrderDetails {
	s.TaxiType = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetTelephone(v string) *SyncTripOrderRequestOrderDetails {
	s.Telephone = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetTransportNumber(v string) *SyncTripOrderRequestOrderDetails {
	s.TransportNumber = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetails) SetTypeDescription(v string) *SyncTripOrderRequestOrderDetails {
	s.TypeDescription = &v
	return s
}

type SyncTripOrderRequestOrderDetailsHotelLocation struct {
	Lat    *string `json:"lat,omitempty" xml:"lat,omitempty"`
	Lon    *string `json:"lon,omitempty" xml:"lon,omitempty"`
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	Url    *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SyncTripOrderRequestOrderDetailsHotelLocation) String() string {
	return tea.Prettify(s)
}

func (s SyncTripOrderRequestOrderDetailsHotelLocation) GoString() string {
	return s.String()
}

func (s *SyncTripOrderRequestOrderDetailsHotelLocation) SetLat(v string) *SyncTripOrderRequestOrderDetailsHotelLocation {
	s.Lat = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetailsHotelLocation) SetLon(v string) *SyncTripOrderRequestOrderDetailsHotelLocation {
	s.Lon = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetailsHotelLocation) SetSource(v string) *SyncTripOrderRequestOrderDetailsHotelLocation {
	s.Source = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetailsHotelLocation) SetUrl(v string) *SyncTripOrderRequestOrderDetailsHotelLocation {
	s.Url = &v
	return s
}

type SyncTripOrderRequestOrderDetailsOpenConsumerInfo struct {
	CorpId       *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	StaffFlag    *bool   `json:"staffFlag,omitempty" xml:"staffFlag,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
	TicketAmount *string `json:"ticketAmount,omitempty" xml:"ticketAmount,omitempty"`
	TicketNo     *string `json:"ticketNo,omitempty" xml:"ticketNo,omitempty"`
	UserId       *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SyncTripOrderRequestOrderDetailsOpenConsumerInfo) String() string {
	return tea.Prettify(s)
}

func (s SyncTripOrderRequestOrderDetailsOpenConsumerInfo) GoString() string {
	return s.String()
}

func (s *SyncTripOrderRequestOrderDetailsOpenConsumerInfo) SetCorpId(v string) *SyncTripOrderRequestOrderDetailsOpenConsumerInfo {
	s.CorpId = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetailsOpenConsumerInfo) SetName(v string) *SyncTripOrderRequestOrderDetailsOpenConsumerInfo {
	s.Name = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetailsOpenConsumerInfo) SetStaffFlag(v bool) *SyncTripOrderRequestOrderDetailsOpenConsumerInfo {
	s.StaffFlag = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetailsOpenConsumerInfo) SetStatus(v string) *SyncTripOrderRequestOrderDetailsOpenConsumerInfo {
	s.Status = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetailsOpenConsumerInfo) SetTicketAmount(v string) *SyncTripOrderRequestOrderDetailsOpenConsumerInfo {
	s.TicketAmount = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetailsOpenConsumerInfo) SetTicketNo(v string) *SyncTripOrderRequestOrderDetailsOpenConsumerInfo {
	s.TicketNo = &v
	return s
}

func (s *SyncTripOrderRequestOrderDetailsOpenConsumerInfo) SetUserId(v string) *SyncTripOrderRequestOrderDetailsOpenConsumerInfo {
	s.UserId = &v
	return s
}

type SyncTripOrderResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SyncTripOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncTripOrderResponseBody) GoString() string {
	return s.String()
}

func (s *SyncTripOrderResponseBody) SetRequestId(v string) *SyncTripOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncTripOrderResponseBody) SetSuccess(v bool) *SyncTripOrderResponseBody {
	s.Success = &v
	return s
}

type SyncTripOrderResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncTripOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncTripOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncTripOrderResponse) GoString() string {
	return s.String()
}

func (s *SyncTripOrderResponse) SetHeaders(v map[string]*string) *SyncTripOrderResponse {
	s.Headers = v
	return s
}

func (s *SyncTripOrderResponse) SetStatusCode(v int32) *SyncTripOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncTripOrderResponse) SetBody(v *SyncTripOrderResponseBody) *SyncTripOrderResponse {
	s.Body = v
	return s
}

type SyncTripProductConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SyncTripProductConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncTripProductConfigHeaders) GoString() string {
	return s.String()
}

func (s *SyncTripProductConfigHeaders) SetCommonHeaders(v map[string]*string) *SyncTripProductConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncTripProductConfigHeaders) SetXAcsDingtalkAccessToken(v string) *SyncTripProductConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SyncTripProductConfigRequest struct {
	TargetCorpId          *string                                              `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
	TripProductConfigList []*SyncTripProductConfigRequestTripProductConfigList `json:"tripProductConfigList,omitempty" xml:"tripProductConfigList,omitempty" type:"Repeated"`
}

func (s SyncTripProductConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncTripProductConfigRequest) GoString() string {
	return s.String()
}

func (s *SyncTripProductConfigRequest) SetTargetCorpId(v string) *SyncTripProductConfigRequest {
	s.TargetCorpId = &v
	return s
}

func (s *SyncTripProductConfigRequest) SetTripProductConfigList(v []*SyncTripProductConfigRequestTripProductConfigList) *SyncTripProductConfigRequest {
	s.TripProductConfigList = v
	return s
}

type SyncTripProductConfigRequestTripProductConfigList struct {
	AllVisible         *bool                                                        `json:"allVisible,omitempty" xml:"allVisible,omitempty"`
	DeptVisibleScopes  []*string                                                    `json:"deptVisibleScopes,omitempty" xml:"deptVisibleScopes,omitempty" type:"Repeated"`
	OpenStatus         *bool                                                        `json:"openStatus,omitempty" xml:"openStatus,omitempty"`
	ProductType        *string                                                      `json:"productType,omitempty" xml:"productType,omitempty"`
	RoleVisibleScopes  []*string                                                    `json:"roleVisibleScopes,omitempty" xml:"roleVisibleScopes,omitempty" type:"Repeated"`
	StaffVisibleScopes []*string                                                    `json:"staffVisibleScopes,omitempty" xml:"staffVisibleScopes,omitempty" type:"Repeated"`
	TmcInfos           []*SyncTripProductConfigRequestTripProductConfigListTmcInfos `json:"tmcInfos,omitempty" xml:"tmcInfos,omitempty" type:"Repeated"`
}

func (s SyncTripProductConfigRequestTripProductConfigList) String() string {
	return tea.Prettify(s)
}

func (s SyncTripProductConfigRequestTripProductConfigList) GoString() string {
	return s.String()
}

func (s *SyncTripProductConfigRequestTripProductConfigList) SetAllVisible(v bool) *SyncTripProductConfigRequestTripProductConfigList {
	s.AllVisible = &v
	return s
}

func (s *SyncTripProductConfigRequestTripProductConfigList) SetDeptVisibleScopes(v []*string) *SyncTripProductConfigRequestTripProductConfigList {
	s.DeptVisibleScopes = v
	return s
}

func (s *SyncTripProductConfigRequestTripProductConfigList) SetOpenStatus(v bool) *SyncTripProductConfigRequestTripProductConfigList {
	s.OpenStatus = &v
	return s
}

func (s *SyncTripProductConfigRequestTripProductConfigList) SetProductType(v string) *SyncTripProductConfigRequestTripProductConfigList {
	s.ProductType = &v
	return s
}

func (s *SyncTripProductConfigRequestTripProductConfigList) SetRoleVisibleScopes(v []*string) *SyncTripProductConfigRequestTripProductConfigList {
	s.RoleVisibleScopes = v
	return s
}

func (s *SyncTripProductConfigRequestTripProductConfigList) SetStaffVisibleScopes(v []*string) *SyncTripProductConfigRequestTripProductConfigList {
	s.StaffVisibleScopes = v
	return s
}

func (s *SyncTripProductConfigRequestTripProductConfigList) SetTmcInfos(v []*SyncTripProductConfigRequestTripProductConfigListTmcInfos) *SyncTripProductConfigRequestTripProductConfigList {
	s.TmcInfos = v
	return s
}

type SyncTripProductConfigRequestTripProductConfigListTmcInfos struct {
	CategoryType *string `json:"categoryType,omitempty" xml:"categoryType,omitempty"`
	GmtOrgPay    *string `json:"gmtOrgPay,omitempty" xml:"gmtOrgPay,omitempty"`
	PayType      *string `json:"payType,omitempty" xml:"payType,omitempty"`
	TmcCorpId    *string `json:"tmcCorpId,omitempty" xml:"tmcCorpId,omitempty"`
}

func (s SyncTripProductConfigRequestTripProductConfigListTmcInfos) String() string {
	return tea.Prettify(s)
}

func (s SyncTripProductConfigRequestTripProductConfigListTmcInfos) GoString() string {
	return s.String()
}

func (s *SyncTripProductConfigRequestTripProductConfigListTmcInfos) SetCategoryType(v string) *SyncTripProductConfigRequestTripProductConfigListTmcInfos {
	s.CategoryType = &v
	return s
}

func (s *SyncTripProductConfigRequestTripProductConfigListTmcInfos) SetGmtOrgPay(v string) *SyncTripProductConfigRequestTripProductConfigListTmcInfos {
	s.GmtOrgPay = &v
	return s
}

func (s *SyncTripProductConfigRequestTripProductConfigListTmcInfos) SetPayType(v string) *SyncTripProductConfigRequestTripProductConfigListTmcInfos {
	s.PayType = &v
	return s
}

func (s *SyncTripProductConfigRequestTripProductConfigListTmcInfos) SetTmcCorpId(v string) *SyncTripProductConfigRequestTripProductConfigListTmcInfos {
	s.TmcCorpId = &v
	return s
}

type SyncTripProductConfigResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SyncTripProductConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncTripProductConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SyncTripProductConfigResponseBody) SetSuccess(v bool) *SyncTripProductConfigResponseBody {
	s.Success = &v
	return s
}

type SyncTripProductConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncTripProductConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncTripProductConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncTripProductConfigResponse) GoString() string {
	return s.String()
}

func (s *SyncTripProductConfigResponse) SetHeaders(v map[string]*string) *SyncTripProductConfigResponse {
	s.Headers = v
	return s
}

func (s *SyncTripProductConfigResponse) SetStatusCode(v int32) *SyncTripProductConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncTripProductConfigResponse) SetBody(v *SyncTripProductConfigResponseBody) *SyncTripProductConfigResponse {
	s.Body = v
	return s
}

type UpgradeTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpgradeTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpgradeTemplateHeaders) GoString() string {
	return s.String()
}

func (s *UpgradeTemplateHeaders) SetCommonHeaders(v map[string]*string) *UpgradeTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpgradeTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *UpgradeTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpgradeTemplateRequest struct {
	ChannelCorpId *string `json:"channelCorpId,omitempty" xml:"channelCorpId,omitempty"`
	ForceUpgrade  *bool   `json:"forceUpgrade,omitempty" xml:"forceUpgrade,omitempty"`
	TmcCorpId     *string `json:"tmcCorpId,omitempty" xml:"tmcCorpId,omitempty"`
}

func (s UpgradeTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpgradeTemplateRequest) SetChannelCorpId(v string) *UpgradeTemplateRequest {
	s.ChannelCorpId = &v
	return s
}

func (s *UpgradeTemplateRequest) SetForceUpgrade(v bool) *UpgradeTemplateRequest {
	s.ForceUpgrade = &v
	return s
}

func (s *UpgradeTemplateRequest) SetTmcCorpId(v string) *UpgradeTemplateRequest {
	s.TmcCorpId = &v
	return s
}

type UpgradeTemplateResponseBody struct {
	Result  *UpgradeTemplateResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpgradeTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeTemplateResponseBody) SetResult(v *UpgradeTemplateResponseBodyResult) *UpgradeTemplateResponseBody {
	s.Result = v
	return s
}

func (s *UpgradeTemplateResponseBody) SetSuccess(v bool) *UpgradeTemplateResponseBody {
	s.Success = &v
	return s
}

type UpgradeTemplateResponseBodyResult struct {
	UpgradeResult *bool `json:"upgradeResult,omitempty" xml:"upgradeResult,omitempty"`
}

func (s UpgradeTemplateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpgradeTemplateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpgradeTemplateResponseBodyResult) SetUpgradeResult(v bool) *UpgradeTemplateResponseBodyResult {
	s.UpgradeResult = &v
	return s
}

type UpgradeTemplateResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpgradeTemplateResponse) SetHeaders(v map[string]*string) *UpgradeTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpgradeTemplateResponse) SetStatusCode(v int32) *UpgradeTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeTemplateResponse) SetBody(v *UpgradeTemplateResponseBody) *UpgradeTemplateResponse {
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

func (client *Client) GetTravelProcessDetailWithOptions(request *GetTravelProcessDetailRequest, headers *GetTravelProcessDetailHeaders, runtime *util.RuntimeOptions) (_result *GetTravelProcessDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProcessCorpId)) {
		query["processCorpId"] = request.ProcessCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		query["processInstanceId"] = request.ProcessInstanceId
	}

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
		Action:      tea.String("GetTravelProcessDetail"),
		Version:     tea.String("trip_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/trip/processes/details"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTravelProcessDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTravelProcessDetail(request *GetTravelProcessDetailRequest) (_result *GetTravelProcessDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTravelProcessDetailHeaders{}
	_result = &GetTravelProcessDetailResponse{}
	_body, _err := client.GetTravelProcessDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PreCheckTemplateWithOptions(request *PreCheckTemplateRequest, headers *PreCheckTemplateHeaders, runtime *util.RuntimeOptions) (_result *PreCheckTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomerCorpId)) {
		body["customerCorpId"] = request.CustomerCorpId
	}

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
		Action:      tea.String("PreCheckTemplate"),
		Version:     tea.String("trip_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/trip/processes/templateUpgrades/preCheck"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PreCheckTemplateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PreCheckTemplate(request *PreCheckTemplateRequest) (_result *PreCheckTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PreCheckTemplateHeaders{}
	_result = &PreCheckTemplateResponse{}
	_body, _err := client.PreCheckTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTripProcessTemplatesWithOptions(request *QueryTripProcessTemplatesRequest, headers *QueryTripProcessTemplatesHeaders, runtime *util.RuntimeOptions) (_result *QueryTripProcessTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomerCorpId)) {
		query["customerCorpId"] = request.CustomerCorpId
	}

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
		Action:      tea.String("QueryTripProcessTemplates"),
		Version:     tea.String("trip_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/trip/processes/templatesDetails"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTripProcessTemplatesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTripProcessTemplates(request *QueryTripProcessTemplatesRequest) (_result *QueryTripProcessTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryTripProcessTemplatesHeaders{}
	_result = &QueryTripProcessTemplatesResponse{}
	_body, _err := client.QueryTripProcessTemplatesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncBusinessSignInfoWithOptions(request *SyncBusinessSignInfoRequest, headers *SyncBusinessSignInfoHeaders, runtime *util.RuntimeOptions) (_result *SyncBusinessSignInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizTypeList)) {
		body["bizTypeList"] = request.BizTypeList
	}

	if !tea.BoolValue(util.IsUnset(request.GmtOrgPay)) {
		body["gmtOrgPay"] = request.GmtOrgPay
	}

	if !tea.BoolValue(util.IsUnset(request.GmtSign)) {
		body["gmtSign"] = request.GmtSign
	}

	if !tea.BoolValue(util.IsUnset(request.OrgPayStatus)) {
		body["orgPayStatus"] = request.OrgPayStatus
	}

	if !tea.BoolValue(util.IsUnset(request.SignStatus)) {
		body["signStatus"] = request.SignStatus
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		body["targetCorpId"] = request.TargetCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.TmcProductDetailList)) {
		body["tmcProductDetailList"] = request.TmcProductDetailList
	}

	if !tea.BoolValue(util.IsUnset(request.TmcProductList)) {
		body["tmcProductList"] = request.TmcProductList
	}

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
		Action:      tea.String("SyncBusinessSignInfo"),
		Version:     tea.String("trip_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/trip/businessSignInfos/sync"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncBusinessSignInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncBusinessSignInfo(request *SyncBusinessSignInfoRequest) (_result *SyncBusinessSignInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncBusinessSignInfoHeaders{}
	_result = &SyncBusinessSignInfoResponse{}
	_body, _err := client.SyncBusinessSignInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncCostCenterWithOptions(request *SyncCostCenterRequest, headers *SyncCostCenterHeaders, runtime *util.RuntimeOptions) (_result *SyncCostCenterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelCorpId)) {
		body["channelCorpId"] = request.ChannelCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.CostCenterId)) {
		body["costCenterId"] = request.CostCenterId
	}

	if !tea.BoolValue(util.IsUnset(request.DeleteFlag)) {
		body["deleteFlag"] = request.DeleteFlag
	}

	if !tea.BoolValue(util.IsUnset(request.Extension)) {
		body["extension"] = request.Extension
	}

	if !tea.BoolValue(util.IsUnset(request.GmtAction)) {
		body["gmtAction"] = request.GmtAction
	}

	if !tea.BoolValue(util.IsUnset(request.Number)) {
		body["number"] = request.Number
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartId)) {
		body["thirdPartId"] = request.ThirdPartId
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
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SyncCostCenter"),
		Version:     tea.String("trip_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/trip/processes/costCenters/sync"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncCostCenterResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncCostCenter(request *SyncCostCenterRequest) (_result *SyncCostCenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncCostCenterHeaders{}
	_result = &SyncCostCenterResponse{}
	_body, _err := client.SyncCostCenterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncCostCenterEntityWithOptions(request *SyncCostCenterEntityRequest, headers *SyncCostCenterEntityHeaders, runtime *util.RuntimeOptions) (_result *SyncCostCenterEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelCorpId)) {
		body["channelCorpId"] = request.ChannelCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.CostCenterId)) {
		body["costCenterId"] = request.CostCenterId
	}

	if !tea.BoolValue(util.IsUnset(request.DelAll)) {
		body["delAll"] = request.DelAll
	}

	if !tea.BoolValue(util.IsUnset(request.EntityList)) {
		body["entityList"] = request.EntityList
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
		Action:      tea.String("SyncCostCenterEntity"),
		Version:     tea.String("trip_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/trip/processes/costCenters/applicableScopes/sync"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncCostCenterEntityResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncCostCenterEntity(request *SyncCostCenterEntityRequest) (_result *SyncCostCenterEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncCostCenterEntityHeaders{}
	_result = &SyncCostCenterEntityResponse{}
	_body, _err := client.SyncCostCenterEntityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncInvoiceWithOptions(request *SyncInvoiceRequest, headers *SyncInvoiceHeaders, runtime *util.RuntimeOptions) (_result *SyncInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		body["address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.BankName)) {
		body["bankName"] = request.BankName
	}

	if !tea.BoolValue(util.IsUnset(request.BankNo)) {
		body["bankNo"] = request.BankNo
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelCorpId)) {
		body["channelCorpId"] = request.ChannelCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DeleteFlag)) {
		body["deleteFlag"] = request.DeleteFlag
	}

	if !tea.BoolValue(util.IsUnset(request.GmtAction)) {
		body["gmtAction"] = request.GmtAction
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceId)) {
		body["invoiceId"] = request.InvoiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectIds)) {
		body["projectIds"] = request.ProjectIds
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.TaxNo)) {
		body["taxNo"] = request.TaxNo
	}

	if !tea.BoolValue(util.IsUnset(request.Tel)) {
		body["tel"] = request.Tel
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartId)) {
		body["thirdPartId"] = request.ThirdPartId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UnitType)) {
		body["unitType"] = request.UnitType
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
		Action:      tea.String("SyncInvoice"),
		Version:     tea.String("trip_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/trip/processes/invoiceTitles/sync"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncInvoiceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncInvoice(request *SyncInvoiceRequest) (_result *SyncInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncInvoiceHeaders{}
	_result = &SyncInvoiceResponse{}
	_body, _err := client.SyncInvoiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncInvoiceEntityWithOptions(request *SyncInvoiceEntityRequest, headers *SyncInvoiceEntityHeaders, runtime *util.RuntimeOptions) (_result *SyncInvoiceEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelCorpId)) {
		body["channelCorpId"] = request.ChannelCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DelAll)) {
		body["delAll"] = request.DelAll
	}

	if !tea.BoolValue(util.IsUnset(request.EntityList)) {
		body["entityList"] = request.EntityList
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceId)) {
		body["invoiceId"] = request.InvoiceId
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
		Action:      tea.String("SyncInvoiceEntity"),
		Version:     tea.String("trip_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/trip/processes/invoiceTitles/applicableScopes/sync"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncInvoiceEntityResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncInvoiceEntity(request *SyncInvoiceEntityRequest) (_result *SyncInvoiceEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncInvoiceEntityHeaders{}
	_result = &SyncInvoiceEntityResponse{}
	_body, _err := client.SyncInvoiceEntityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncProjectWithOptions(request *SyncProjectRequest, headers *SyncProjectHeaders, runtime *util.RuntimeOptions) (_result *SyncProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelCorpId)) {
		body["channelCorpId"] = request.ChannelCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.CostCenterId)) {
		body["costCenterId"] = request.CostCenterId
	}

	if !tea.BoolValue(util.IsUnset(request.DeleteFlag)) {
		body["deleteFlag"] = request.DeleteFlag
	}

	if !tea.BoolValue(util.IsUnset(request.Extension)) {
		body["extension"] = request.Extension
	}

	if !tea.BoolValue(util.IsUnset(request.GmtAction)) {
		body["gmtAction"] = request.GmtAction
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceId)) {
		body["invoiceId"] = request.InvoiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerIds)) {
		body["managerIds"] = request.ManagerIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["projectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["projectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartId)) {
		body["thirdPartId"] = request.ThirdPartId
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
		Action:      tea.String("SyncProject"),
		Version:     tea.String("trip_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/trip/processes/projects/sync"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncProjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncProject(request *SyncProjectRequest) (_result *SyncProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncProjectHeaders{}
	_result = &SyncProjectResponse{}
	_body, _err := client.SyncProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncProjectEntityWithOptions(request *SyncProjectEntityRequest, headers *SyncProjectEntityHeaders, runtime *util.RuntimeOptions) (_result *SyncProjectEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelCorpId)) {
		body["channelCorpId"] = request.ChannelCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DelAll)) {
		body["delAll"] = request.DelAll
	}

	if !tea.BoolValue(util.IsUnset(request.EntityList)) {
		body["entityList"] = request.EntityList
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["projectId"] = request.ProjectId
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
		Action:      tea.String("SyncProjectEntity"),
		Version:     tea.String("trip_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/trip/processes/projects/applicableScopes/sync"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncProjectEntityResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncProjectEntity(request *SyncProjectEntityRequest) (_result *SyncProjectEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncProjectEntityHeaders{}
	_result = &SyncProjectEntityResponse{}
	_body, _err := client.SyncProjectEntityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncSecretKeyWithOptions(request *SyncSecretKeyRequest, headers *SyncSecretKeyHeaders, runtime *util.RuntimeOptions) (_result *SyncSecretKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionType)) {
		body["actionType"] = request.ActionType
	}

	if !tea.BoolValue(util.IsUnset(request.SecretString)) {
		body["secretString"] = request.SecretString
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		body["targetCorpId"] = request.TargetCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.TripAppKey)) {
		body["tripAppKey"] = request.TripAppKey
	}

	if !tea.BoolValue(util.IsUnset(request.TripAppSecurity)) {
		body["tripAppSecurity"] = request.TripAppSecurity
	}

	if !tea.BoolValue(util.IsUnset(request.TripCorpId)) {
		body["tripCorpId"] = request.TripCorpId
	}

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
		Action:      tea.String("SyncSecretKey"),
		Version:     tea.String("trip_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/trip/secretKeys/sync"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncSecretKeyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncSecretKey(request *SyncSecretKeyRequest) (_result *SyncSecretKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncSecretKeyHeaders{}
	_result = &SyncSecretKeyResponse{}
	_body, _err := client.SyncSecretKeyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncTripOrderWithOptions(request *SyncTripOrderRequest, headers *SyncTripOrderHeaders, runtime *util.RuntimeOptions) (_result *SyncTripOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizExtension)) {
		body["bizExtension"] = request.BizExtension
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelType)) {
		body["channelType"] = request.ChannelType
	}

	if !tea.BoolValue(util.IsUnset(request.Currency)) {
		body["currency"] = request.Currency
	}

	if !tea.BoolValue(util.IsUnset(request.DingUserId)) {
		body["dingUserId"] = request.DingUserId
	}

	if !tea.BoolValue(util.IsUnset(request.DiscountAmount)) {
		body["discountAmount"] = request.DiscountAmount
	}

	if !tea.BoolValue(util.IsUnset(request.EndorseFlag)) {
		body["endorseFlag"] = request.EndorseFlag
	}

	if !tea.BoolValue(util.IsUnset(request.Event)) {
		body["event"] = request.Event
	}

	if !tea.BoolValue(util.IsUnset(request.GmtOrder)) {
		body["gmtOrder"] = request.GmtOrder
	}

	if !tea.BoolValue(util.IsUnset(request.GmtPay)) {
		body["gmtPay"] = request.GmtPay
	}

	if !tea.BoolValue(util.IsUnset(request.GmtRefund)) {
		body["gmtRefund"] = request.GmtRefund
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceApplyUrl)) {
		body["invoiceApplyUrl"] = request.InvoiceApplyUrl
	}

	if !tea.BoolValue(util.IsUnset(request.JourneyBizNo)) {
		body["journeyBizNo"] = request.JourneyBizNo
	}

	if !tea.BoolValue(util.IsUnset(request.OrderDetails)) {
		body["orderDetails"] = request.OrderDetails
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNo)) {
		body["orderNo"] = request.OrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.OrderUrl)) {
		body["orderUrl"] = request.OrderUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessId)) {
		body["processId"] = request.ProcessId
	}

	if !tea.BoolValue(util.IsUnset(request.RealAmount)) {
		body["realAmount"] = request.RealAmount
	}

	if !tea.BoolValue(util.IsUnset(request.RefundAmount)) {
		body["refundAmount"] = request.RefundAmount
	}

	if !tea.BoolValue(util.IsUnset(request.RelativeOrderNo)) {
		body["relativeOrderNo"] = request.RelativeOrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SupplyLogo)) {
		body["supplyLogo"] = request.SupplyLogo
	}

	if !tea.BoolValue(util.IsUnset(request.SupplyName)) {
		body["supplyName"] = request.SupplyName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		body["targetCorpId"] = request.TargetCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.TmcCorpId)) {
		body["tmcCorpId"] = request.TmcCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.TotalAmount)) {
		body["totalAmount"] = request.TotalAmount
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
		Action:      tea.String("SyncTripOrder"),
		Version:     tea.String("trip_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/trip/tripOrders/sync"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncTripOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncTripOrder(request *SyncTripOrderRequest) (_result *SyncTripOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncTripOrderHeaders{}
	_result = &SyncTripOrderResponse{}
	_body, _err := client.SyncTripOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncTripProductConfigWithOptions(request *SyncTripProductConfigRequest, headers *SyncTripProductConfigHeaders, runtime *util.RuntimeOptions) (_result *SyncTripProductConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		body["targetCorpId"] = request.TargetCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.TripProductConfigList)) {
		body["tripProductConfigList"] = request.TripProductConfigList
	}

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
		Action:      tea.String("SyncTripProductConfig"),
		Version:     tea.String("trip_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/trip/productConfigs/sync"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncTripProductConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncTripProductConfig(request *SyncTripProductConfigRequest) (_result *SyncTripProductConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncTripProductConfigHeaders{}
	_result = &SyncTripProductConfigResponse{}
	_body, _err := client.SyncTripProductConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeTemplateWithOptions(request *UpgradeTemplateRequest, headers *UpgradeTemplateHeaders, runtime *util.RuntimeOptions) (_result *UpgradeTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelCorpId)) {
		body["channelCorpId"] = request.ChannelCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.ForceUpgrade)) {
		body["forceUpgrade"] = request.ForceUpgrade
	}

	if !tea.BoolValue(util.IsUnset(request.TmcCorpId)) {
		body["tmcCorpId"] = request.TmcCorpId
	}

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
		Action:      tea.String("UpgradeTemplate"),
		Version:     tea.String("trip_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/trip/process/templates/upgrade"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeTemplateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeTemplate(request *UpgradeTemplateRequest) (_result *UpgradeTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpgradeTemplateHeaders{}
	_result = &UpgradeTemplateResponse{}
	_body, _err := client.UpgradeTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
