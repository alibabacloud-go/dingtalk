// This file is auto-generated, don't edit it. Thanks.
package datacenter_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CloseDataDeliverHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CloseDataDeliverHeaders) String() string {
	return tea.Prettify(s)
}

func (s CloseDataDeliverHeaders) GoString() string {
	return s.String()
}

func (s *CloseDataDeliverHeaders) SetCommonHeaders(v map[string]*string) *CloseDataDeliverHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CloseDataDeliverHeaders) SetXAcsDingtalkAccessToken(v string) *CloseDataDeliverHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CloseDataDeliverRequest struct {
	// example:
	//
	// DELIVER-3e1a2d2f-fa76-45e8-XXXX-7fd29307c859
	DeliverId *string `json:"deliverId,omitempty" xml:"deliverId,omitempty"`
	// example:
	//
	// RT
	DispatchingItemType *string `json:"dispatchingItemType,omitempty" xml:"dispatchingItemType,omitempty"`
}

func (s CloseDataDeliverRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseDataDeliverRequest) GoString() string {
	return s.String()
}

func (s *CloseDataDeliverRequest) SetDeliverId(v string) *CloseDataDeliverRequest {
	s.DeliverId = &v
	return s
}

func (s *CloseDataDeliverRequest) SetDispatchingItemType(v string) *CloseDataDeliverRequest {
	s.DispatchingItemType = &v
	return s
}

type CloseDataDeliverResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CloseDataDeliverResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseDataDeliverResponseBody) GoString() string {
	return s.String()
}

func (s *CloseDataDeliverResponseBody) SetResult(v string) *CloseDataDeliverResponseBody {
	s.Result = &v
	return s
}

func (s *CloseDataDeliverResponseBody) SetSuccess(v bool) *CloseDataDeliverResponseBody {
	s.Success = &v
	return s
}

type CloseDataDeliverResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseDataDeliverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseDataDeliverResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseDataDeliverResponse) GoString() string {
	return s.String()
}

func (s *CloseDataDeliverResponse) SetHeaders(v map[string]*string) *CloseDataDeliverResponse {
	s.Headers = v
	return s
}

func (s *CloseDataDeliverResponse) SetStatusCode(v int32) *CloseDataDeliverResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseDataDeliverResponse) SetBody(v *CloseDataDeliverResponseBody) *CloseDataDeliverResponse {
	s.Body = v
	return s
}

type CreateDataDeliverHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateDataDeliverHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateDataDeliverHeaders) GoString() string {
	return s.String()
}

func (s *CreateDataDeliverHeaders) SetCommonHeaders(v map[string]*string) *CreateDataDeliverHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateDataDeliverHeaders) SetXAcsDingtalkAccessToken(v string) *CreateDataDeliverHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateDataDeliverRequest struct {
	Bizcode              *string `json:"bizcode,omitempty" xml:"bizcode,omitempty"`
	Param                *string `json:"param,omitempty" xml:"param,omitempty"`
	UserId               *string `json:"userId,omitempty" xml:"userId,omitempty"`
	DispatchingCycle     *string `json:"dispatchingCycle,omitempty" xml:"dispatchingCycle,omitempty"`
	DispatchingItemType  *string `json:"dispatchingItemType,omitempty" xml:"dispatchingItemType,omitempty"`
	DispatchingStartDate *int64  `json:"dispatchingStartDate,omitempty" xml:"dispatchingStartDate,omitempty"`
}

func (s CreateDataDeliverRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataDeliverRequest) GoString() string {
	return s.String()
}

func (s *CreateDataDeliverRequest) SetBizcode(v string) *CreateDataDeliverRequest {
	s.Bizcode = &v
	return s
}

func (s *CreateDataDeliverRequest) SetParam(v string) *CreateDataDeliverRequest {
	s.Param = &v
	return s
}

func (s *CreateDataDeliverRequest) SetUserId(v string) *CreateDataDeliverRequest {
	s.UserId = &v
	return s
}

func (s *CreateDataDeliverRequest) SetDispatchingCycle(v string) *CreateDataDeliverRequest {
	s.DispatchingCycle = &v
	return s
}

func (s *CreateDataDeliverRequest) SetDispatchingItemType(v string) *CreateDataDeliverRequest {
	s.DispatchingItemType = &v
	return s
}

func (s *CreateDataDeliverRequest) SetDispatchingStartDate(v int64) *CreateDataDeliverRequest {
	s.DispatchingStartDate = &v
	return s
}

type CreateDataDeliverResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateDataDeliverResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDataDeliverResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataDeliverResponseBody) SetResult(v string) *CreateDataDeliverResponseBody {
	s.Result = &v
	return s
}

func (s *CreateDataDeliverResponseBody) SetSuccess(v bool) *CreateDataDeliverResponseBody {
	s.Success = &v
	return s
}

type CreateDataDeliverResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataDeliverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataDeliverResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDataDeliverResponse) GoString() string {
	return s.String()
}

func (s *CreateDataDeliverResponse) SetHeaders(v map[string]*string) *CreateDataDeliverResponse {
	s.Headers = v
	return s
}

func (s *CreateDataDeliverResponse) SetStatusCode(v int32) *CreateDataDeliverResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataDeliverResponse) SetBody(v *CreateDataDeliverResponseBody) *CreateDataDeliverResponse {
	s.Body = v
	return s
}

type CreateScreenHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateScreenHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateScreenHeaders) GoString() string {
	return s.String()
}

func (s *CreateScreenHeaders) SetCommonHeaders(v map[string]*string) *CreateScreenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateScreenHeaders) SetXAcsDingtalkAccessToken(v string) *CreateScreenHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateScreenRequest struct {
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s CreateScreenRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScreenRequest) GoString() string {
	return s.String()
}

func (s *CreateScreenRequest) SetOperatorId(v string) *CreateScreenRequest {
	s.OperatorId = &v
	return s
}

func (s *CreateScreenRequest) SetTemplateId(v string) *CreateScreenRequest {
	s.TemplateId = &v
	return s
}

type CreateScreenResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateScreenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScreenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScreenResponseBody) SetResult(v string) *CreateScreenResponseBody {
	s.Result = &v
	return s
}

func (s *CreateScreenResponseBody) SetSuccess(v bool) *CreateScreenResponseBody {
	s.Success = &v
	return s
}

type CreateScreenResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScreenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScreenResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScreenResponse) GoString() string {
	return s.String()
}

func (s *CreateScreenResponse) SetHeaders(v map[string]*string) *CreateScreenResponse {
	s.Headers = v
	return s
}

func (s *CreateScreenResponse) SetStatusCode(v int32) *CreateScreenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScreenResponse) SetBody(v *CreateScreenResponseBody) *CreateScreenResponse {
	s.Body = v
	return s
}

type DataMarketIsvServiceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DataMarketIsvServiceHeaders) String() string {
	return tea.Prettify(s)
}

func (s DataMarketIsvServiceHeaders) GoString() string {
	return s.String()
}

func (s *DataMarketIsvServiceHeaders) SetCommonHeaders(v map[string]*string) *DataMarketIsvServiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DataMarketIsvServiceHeaders) SetXAcsDingtalkAccessToken(v string) *DataMarketIsvServiceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DataMarketIsvServiceRequest struct {
	// This parameter is required.
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// This parameter is required.
	ApiId *string `json:"apiId,omitempty" xml:"apiId,omitempty"`
	// This parameter is required.
	Args *string `json:"args,omitempty" xml:"args,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DataMarketIsvServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DataMarketIsvServiceRequest) GoString() string {
	return s.String()
}

func (s *DataMarketIsvServiceRequest) SetAccountId(v string) *DataMarketIsvServiceRequest {
	s.AccountId = &v
	return s
}

func (s *DataMarketIsvServiceRequest) SetApiId(v string) *DataMarketIsvServiceRequest {
	s.ApiId = &v
	return s
}

func (s *DataMarketIsvServiceRequest) SetArgs(v string) *DataMarketIsvServiceRequest {
	s.Args = &v
	return s
}

func (s *DataMarketIsvServiceRequest) SetUserId(v string) *DataMarketIsvServiceRequest {
	s.UserId = &v
	return s
}

type DataMarketIsvServiceResponseBody struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	Msg  *string `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s DataMarketIsvServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DataMarketIsvServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DataMarketIsvServiceResponseBody) SetCode(v string) *DataMarketIsvServiceResponseBody {
	s.Code = &v
	return s
}

func (s *DataMarketIsvServiceResponseBody) SetData(v string) *DataMarketIsvServiceResponseBody {
	s.Data = &v
	return s
}

func (s *DataMarketIsvServiceResponseBody) SetMsg(v string) *DataMarketIsvServiceResponseBody {
	s.Msg = &v
	return s
}

type DataMarketIsvServiceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DataMarketIsvServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DataMarketIsvServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DataMarketIsvServiceResponse) GoString() string {
	return s.String()
}

func (s *DataMarketIsvServiceResponse) SetHeaders(v map[string]*string) *DataMarketIsvServiceResponse {
	s.Headers = v
	return s
}

func (s *DataMarketIsvServiceResponse) SetStatusCode(v int32) *DataMarketIsvServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DataMarketIsvServiceResponse) SetBody(v *DataMarketIsvServiceResponseBody) *DataMarketIsvServiceResponse {
	s.Body = v
	return s
}

type DataMarketServiceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DataMarketServiceHeaders) String() string {
	return tea.Prettify(s)
}

func (s DataMarketServiceHeaders) GoString() string {
	return s.String()
}

func (s *DataMarketServiceHeaders) SetCommonHeaders(v map[string]*string) *DataMarketServiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DataMarketServiceHeaders) SetXAcsDingtalkAccessToken(v string) *DataMarketServiceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DataMarketServiceRequest struct {
	ApiId *string `json:"apiId,omitempty" xml:"apiId,omitempty"`
	Args  *string `json:"args,omitempty" xml:"args,omitempty"`
}

func (s DataMarketServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DataMarketServiceRequest) GoString() string {
	return s.String()
}

func (s *DataMarketServiceRequest) SetApiId(v string) *DataMarketServiceRequest {
	s.ApiId = &v
	return s
}

func (s *DataMarketServiceRequest) SetArgs(v string) *DataMarketServiceRequest {
	s.Args = &v
	return s
}

type DataMarketServiceResponseBody struct {
	Code       *string `json:"code,omitempty" xml:"code,omitempty"`
	Data       *string `json:"data,omitempty" xml:"data,omitempty"`
	Msg        *string `json:"msg,omitempty" xml:"msg,omitempty"`
	TotalQuota *int64  `json:"totalQuota,omitempty" xml:"totalQuota,omitempty"`
	UsedQuota  *int64  `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
}

func (s DataMarketServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DataMarketServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DataMarketServiceResponseBody) SetCode(v string) *DataMarketServiceResponseBody {
	s.Code = &v
	return s
}

func (s *DataMarketServiceResponseBody) SetData(v string) *DataMarketServiceResponseBody {
	s.Data = &v
	return s
}

func (s *DataMarketServiceResponseBody) SetMsg(v string) *DataMarketServiceResponseBody {
	s.Msg = &v
	return s
}

func (s *DataMarketServiceResponseBody) SetTotalQuota(v int64) *DataMarketServiceResponseBody {
	s.TotalQuota = &v
	return s
}

func (s *DataMarketServiceResponseBody) SetUsedQuota(v int64) *DataMarketServiceResponseBody {
	s.UsedQuota = &v
	return s
}

type DataMarketServiceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DataMarketServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DataMarketServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DataMarketServiceResponse) GoString() string {
	return s.String()
}

func (s *DataMarketServiceResponse) SetHeaders(v map[string]*string) *DataMarketServiceResponse {
	s.Headers = v
	return s
}

func (s *DataMarketServiceResponse) SetStatusCode(v int32) *DataMarketServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DataMarketServiceResponse) SetBody(v *DataMarketServiceResponseBody) *DataMarketServiceResponse {
	s.Body = v
	return s
}

type GetAbnormalOperationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAbnormalOperationHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAbnormalOperationHeaders) GoString() string {
	return s.String()
}

func (s *GetAbnormalOperationHeaders) SetCommonHeaders(v map[string]*string) *GetAbnormalOperationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAbnormalOperationHeaders) SetXAcsDingtalkAccessToken(v string) *GetAbnormalOperationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAbnormalOperationRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey,omitempty"`
}

func (s GetAbnormalOperationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAbnormalOperationRequest) GoString() string {
	return s.String()
}

func (s *GetAbnormalOperationRequest) SetPageNumber(v int32) *GetAbnormalOperationRequest {
	s.PageNumber = &v
	return s
}

func (s *GetAbnormalOperationRequest) SetPageSize(v int32) *GetAbnormalOperationRequest {
	s.PageSize = &v
	return s
}

func (s *GetAbnormalOperationRequest) SetSearchKey(v string) *GetAbnormalOperationRequest {
	s.SearchKey = &v
	return s
}

type GetAbnormalOperationResponseBody struct {
	// example:
	//
	// [     {       "DEPARTMENT": "xx",       "IN_REASON": "xx",       "OUT_DATE": "2006-12-07",       "OUT_DEPARTMENT": "xx",       "IN_DATE": "2006-12-07",       "OUT_REASON": "xx"     }   ]
	Data  *string `json:"data,omitempty" xml:"data,omitempty"`
	Total *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetAbnormalOperationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAbnormalOperationResponseBody) GoString() string {
	return s.String()
}

func (s *GetAbnormalOperationResponseBody) SetData(v string) *GetAbnormalOperationResponseBody {
	s.Data = &v
	return s
}

func (s *GetAbnormalOperationResponseBody) SetTotal(v int64) *GetAbnormalOperationResponseBody {
	s.Total = &v
	return s
}

type GetAbnormalOperationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAbnormalOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAbnormalOperationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAbnormalOperationResponse) GoString() string {
	return s.String()
}

func (s *GetAbnormalOperationResponse) SetHeaders(v map[string]*string) *GetAbnormalOperationResponse {
	s.Headers = v
	return s
}

func (s *GetAbnormalOperationResponse) SetStatusCode(v int32) *GetAbnormalOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAbnormalOperationResponse) SetBody(v *GetAbnormalOperationResponseBody) *GetAbnormalOperationResponse {
	s.Body = v
	return s
}

type GetAdministrativeLicensingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAdministrativeLicensingHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAdministrativeLicensingHeaders) GoString() string {
	return s.String()
}

func (s *GetAdministrativeLicensingHeaders) SetCommonHeaders(v map[string]*string) *GetAdministrativeLicensingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAdministrativeLicensingHeaders) SetXAcsDingtalkAccessToken(v string) *GetAdministrativeLicensingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAdministrativeLicensingRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey,omitempty"`
}

func (s GetAdministrativeLicensingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAdministrativeLicensingRequest) GoString() string {
	return s.String()
}

func (s *GetAdministrativeLicensingRequest) SetPageNumber(v int32) *GetAdministrativeLicensingRequest {
	s.PageNumber = &v
	return s
}

func (s *GetAdministrativeLicensingRequest) SetPageSize(v int32) *GetAdministrativeLicensingRequest {
	s.PageSize = &v
	return s
}

func (s *GetAdministrativeLicensingRequest) SetSearchKey(v string) *GetAdministrativeLicensingRequest {
	s.SearchKey = &v
	return s
}

type GetAdministrativeLicensingResponseBody struct {
	// example:
	//
	// [     {       "LicenseNo": "梯4403331978",       "StartDate": "2022-05-10",       "Department": "深圳市市场监督管理局",       "Content": "注册代码:7;设备种类:电梯",       "LicenseName": "特种设备使用登记",       "EndDate": "2099-12-31"     },     {       "LicenseNo": "东水务审﹝2021﹞8267号",       "StartDate": "2021-06-11",       "Department": "东莞市水务局",       "Content": "水土保持方案审批准予行政许可决定",       "LicenseName": "",       "EndDate": "2026-12-31"     } ]
	Data  *string `json:"data,omitempty" xml:"data,omitempty"`
	Total *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetAdministrativeLicensingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAdministrativeLicensingResponseBody) GoString() string {
	return s.String()
}

func (s *GetAdministrativeLicensingResponseBody) SetData(v string) *GetAdministrativeLicensingResponseBody {
	s.Data = &v
	return s
}

func (s *GetAdministrativeLicensingResponseBody) SetTotal(v int64) *GetAdministrativeLicensingResponseBody {
	s.Total = &v
	return s
}

type GetAdministrativeLicensingResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAdministrativeLicensingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAdministrativeLicensingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAdministrativeLicensingResponse) GoString() string {
	return s.String()
}

func (s *GetAdministrativeLicensingResponse) SetHeaders(v map[string]*string) *GetAdministrativeLicensingResponse {
	s.Headers = v
	return s
}

func (s *GetAdministrativeLicensingResponse) SetStatusCode(v int32) *GetAdministrativeLicensingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAdministrativeLicensingResponse) SetBody(v *GetAdministrativeLicensingResponseBody) *GetAdministrativeLicensingResponse {
	s.Body = v
	return s
}

type GetAdministrativePenaltiesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAdministrativePenaltiesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAdministrativePenaltiesHeaders) GoString() string {
	return s.String()
}

func (s *GetAdministrativePenaltiesHeaders) SetCommonHeaders(v map[string]*string) *GetAdministrativePenaltiesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAdministrativePenaltiesHeaders) SetXAcsDingtalkAccessToken(v string) *GetAdministrativePenaltiesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAdministrativePenaltiesRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey,omitempty"`
}

func (s GetAdministrativePenaltiesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAdministrativePenaltiesRequest) GoString() string {
	return s.String()
}

func (s *GetAdministrativePenaltiesRequest) SetPageNumber(v int32) *GetAdministrativePenaltiesRequest {
	s.PageNumber = &v
	return s
}

func (s *GetAdministrativePenaltiesRequest) SetPageSize(v int32) *GetAdministrativePenaltiesRequest {
	s.PageSize = &v
	return s
}

func (s *GetAdministrativePenaltiesRequest) SetSearchKey(v string) *GetAdministrativePenaltiesRequest {
	s.SearchKey = &v
	return s
}

type GetAdministrativePenaltiesResponseBody struct {
	// example:
	//
	// [     {       "DATE_COL": "xx",       "NUMBER": "xx",       "ILLEGAL_TYPE": "xx",       "DEPARTMENT": "xx",       "PUBLIC_DATE": "xx",       "CONTENT": "xx",       "BASED_ON":"xx",       "DESCRIPTION":"xx"      }   ]
	Data  *string `json:"data,omitempty" xml:"data,omitempty"`
	Total *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetAdministrativePenaltiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAdministrativePenaltiesResponseBody) GoString() string {
	return s.String()
}

func (s *GetAdministrativePenaltiesResponseBody) SetData(v string) *GetAdministrativePenaltiesResponseBody {
	s.Data = &v
	return s
}

func (s *GetAdministrativePenaltiesResponseBody) SetTotal(v int64) *GetAdministrativePenaltiesResponseBody {
	s.Total = &v
	return s
}

type GetAdministrativePenaltiesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAdministrativePenaltiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAdministrativePenaltiesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAdministrativePenaltiesResponse) GoString() string {
	return s.String()
}

func (s *GetAdministrativePenaltiesResponse) SetHeaders(v map[string]*string) *GetAdministrativePenaltiesResponse {
	s.Headers = v
	return s
}

func (s *GetAdministrativePenaltiesResponse) SetStatusCode(v int32) *GetAdministrativePenaltiesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAdministrativePenaltiesResponse) SetBody(v *GetAdministrativePenaltiesResponseBody) *GetAdministrativePenaltiesResponse {
	s.Body = v
	return s
}

type GetBasicInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetBasicInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetBasicInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetBasicInfoHeaders) SetCommonHeaders(v map[string]*string) *GetBasicInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetBasicInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetBasicInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetBasicInfoRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey,omitempty"`
}

func (s GetBasicInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *GetBasicInfoRequest) SetPageNumber(v int32) *GetBasicInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *GetBasicInfoRequest) SetPageSize(v int32) *GetBasicInfoRequest {
	s.PageSize = &v
	return s
}

func (s *GetBasicInfoRequest) SetSearchKey(v string) *GetBasicInfoRequest {
	s.SearchKey = &v
	return s
}

type GetBasicInfoResponseBody struct {
	// example:
	//
	// [     {       "ENT_NAME": "xx",       "LEGAL_NAME": "xx",       "ES_DATE": "2006-12-07",       "ENT_STATUS": "在营",       "REG_CAP": "1000万人民币",       "REC_CAP": "1000万人民币",       "SOCIAL_CREDIT_CODE": "91330108793696828A",       "LICENSE_NUMBER": "330108000000965",       "ORG_NO": "793696828",       "TAX_NUM": "91330108793696828A",       "ENT_TYPE": "有限责任公司(非自然人投资或控股的法人独资)",       "INDUSTRY_NAME_LV1": "租赁和商务服务业",       "INDUSTRY_NAME_LV2": "商务服务业",       "OP_FROM": "2006-12-07",       "OP_TO": "2036-12-06",       "COLLEGUES_NUM": "6",       "ENT_NAME_ENG": "Hangzhou Ali Baba Advertising Co.,Ltd.",       "FORMER_NAMES": "xx",       "REG_ORG": "xx",       "REG_ORG_PROVINCE": "浙江省",       "REG_ORG_CITY": "杭州市",       "REG_ORG_DISTRICT": "滨江区",       "STD_REG_CAP": 10000000     }   ]
	Data  *string `json:"data,omitempty" xml:"data,omitempty"`
	Total *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetBasicInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetBasicInfoResponseBody) SetData(v string) *GetBasicInfoResponseBody {
	s.Data = &v
	return s
}

func (s *GetBasicInfoResponseBody) SetTotal(v int64) *GetBasicInfoResponseBody {
	s.Total = &v
	return s
}

type GetBasicInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBasicInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *GetBasicInfoResponse) SetHeaders(v map[string]*string) *GetBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *GetBasicInfoResponse) SetStatusCode(v int32) *GetBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBasicInfoResponse) SetBody(v *GetBasicInfoResponseBody) *GetBasicInfoResponse {
	s.Body = v
	return s
}

type GetBiddingInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetBiddingInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetBiddingInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetBiddingInfoHeaders) SetCommonHeaders(v map[string]*string) *GetBiddingInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetBiddingInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetBiddingInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetBiddingInfoRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey,omitempty"`
}

func (s GetBiddingInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBiddingInfoRequest) GoString() string {
	return s.String()
}

func (s *GetBiddingInfoRequest) SetPageNumber(v int32) *GetBiddingInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *GetBiddingInfoRequest) SetPageSize(v int32) *GetBiddingInfoRequest {
	s.PageSize = &v
	return s
}

func (s *GetBiddingInfoRequest) SetSearchKey(v string) *GetBiddingInfoRequest {
	s.SearchKey = &v
	return s
}

type GetBiddingInfoResponseBody struct {
	// example:
	//
	// [{ "EntName":"企业名称", "BidTitle":"标文标题", "BidType":"招标方式", "RegionName":"地区", "BidIndustry":"标的所属行业", "PublicDate":"发布时间", "ProjectNum":"项目编号", "ProjectName":"项目名称", "ProjectAmount":"项目金额", "TenderEntName":"招标企业", "AgentEntName":"代理企业", "WinnerEntName":"中标企业", "Content":"正文", "InfoType":"标文类型", "SubType":"子类型", "OpeningTime":"开标时间" }]
	Data  *string `json:"data,omitempty" xml:"data,omitempty"`
	Total *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetBiddingInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBiddingInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetBiddingInfoResponseBody) SetData(v string) *GetBiddingInfoResponseBody {
	s.Data = &v
	return s
}

func (s *GetBiddingInfoResponseBody) SetTotal(v int64) *GetBiddingInfoResponseBody {
	s.Total = &v
	return s
}

type GetBiddingInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBiddingInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBiddingInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBiddingInfoResponse) GoString() string {
	return s.String()
}

func (s *GetBiddingInfoResponse) SetHeaders(v map[string]*string) *GetBiddingInfoResponse {
	s.Headers = v
	return s
}

func (s *GetBiddingInfoResponse) SetStatusCode(v int32) *GetBiddingInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBiddingInfoResponse) SetBody(v *GetBiddingInfoResponseBody) *GetBiddingInfoResponse {
	s.Body = v
	return s
}

type GetBranchInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetBranchInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetBranchInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetBranchInfoHeaders) SetCommonHeaders(v map[string]*string) *GetBranchInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetBranchInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetBranchInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetBranchInfoRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey,omitempty"`
}

func (s GetBranchInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBranchInfoRequest) GoString() string {
	return s.String()
}

func (s *GetBranchInfoRequest) SetPageNumber(v int32) *GetBranchInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *GetBranchInfoRequest) SetPageSize(v int32) *GetBranchInfoRequest {
	s.PageSize = &v
	return s
}

func (s *GetBranchInfoRequest) SetSearchKey(v string) *GetBranchInfoRequest {
	s.SearchKey = &v
	return s
}

type GetBranchInfoResponseBody struct {
	// example:
	//
	// [     {       "OperName": "李柯",       "EntStatus": "",       "EntName": "华为技术有限公司驻广州办事处",       "EsDate": ""     },     {       "OperName": "李实",       "EntStatus": "",       "EntName": "华为技术有限公司重庆分公司",       "EsDate": ""     } ]
	Data  *string `json:"data,omitempty" xml:"data,omitempty"`
	Total *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetBranchInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBranchInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetBranchInfoResponseBody) SetData(v string) *GetBranchInfoResponseBody {
	s.Data = &v
	return s
}

func (s *GetBranchInfoResponseBody) SetTotal(v int64) *GetBranchInfoResponseBody {
	s.Total = &v
	return s
}

type GetBranchInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBranchInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBranchInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBranchInfoResponse) GoString() string {
	return s.String()
}

func (s *GetBranchInfoResponse) SetHeaders(v map[string]*string) *GetBranchInfoResponse {
	s.Headers = v
	return s
}

func (s *GetBranchInfoResponse) SetStatusCode(v int32) *GetBranchInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBranchInfoResponse) SetBody(v *GetBranchInfoResponseBody) *GetBranchInfoResponse {
	s.Body = v
	return s
}

type GetChangeRecordHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetChangeRecordHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetChangeRecordHeaders) GoString() string {
	return s.String()
}

func (s *GetChangeRecordHeaders) SetCommonHeaders(v map[string]*string) *GetChangeRecordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetChangeRecordHeaders) SetXAcsDingtalkAccessToken(v string) *GetChangeRecordHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetChangeRecordRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey,omitempty"`
}

func (s GetChangeRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s GetChangeRecordRequest) GoString() string {
	return s.String()
}

func (s *GetChangeRecordRequest) SetPageNumber(v int32) *GetChangeRecordRequest {
	s.PageNumber = &v
	return s
}

func (s *GetChangeRecordRequest) SetPageSize(v int32) *GetChangeRecordRequest {
	s.PageSize = &v
	return s
}

func (s *GetChangeRecordRequest) SetSearchKey(v string) *GetChangeRecordRequest {
	s.SearchKey = &v
	return s
}

type GetChangeRecordResponseBody struct {
	// example:
	//
	// [         {             "Type":"投资人变更(包括出资额、出资方式、出资日期、投资人名称等)",             "ChangeDate":"2014-12-23",             "AfterContent":"股东名称:华为投资控股有限公司、出资额:3990813.182000、出资比例:100.000000;",             "BeforeContent":"股东名称:华为投资控股有限公司、出资额:3960813.182000、出资比例:100.000000;"         },         {             "Type":"期限变更(经营期限、营业期限、驻在期限、合伙期限等变更)",             "ChangeDate":"1997-12-04",             "AfterContent":"1987-09-15,2040-04-09",             "BeforeContent":"1987-09-15,1998-12-31"         } ]
	Data  *string `json:"data,omitempty" xml:"data,omitempty"`
	Total *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetChangeRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetChangeRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetChangeRecordResponseBody) SetData(v string) *GetChangeRecordResponseBody {
	s.Data = &v
	return s
}

func (s *GetChangeRecordResponseBody) SetTotal(v int64) *GetChangeRecordResponseBody {
	s.Total = &v
	return s
}

type GetChangeRecordResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChangeRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChangeRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s GetChangeRecordResponse) GoString() string {
	return s.String()
}

func (s *GetChangeRecordResponse) SetHeaders(v map[string]*string) *GetChangeRecordResponse {
	s.Headers = v
	return s
}

func (s *GetChangeRecordResponse) SetStatusCode(v int32) *GetChangeRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChangeRecordResponse) SetBody(v *GetChangeRecordResponseBody) *GetChangeRecordResponse {
	s.Body = v
	return s
}

type GetDataDeliverHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDataDeliverHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDataDeliverHeaders) GoString() string {
	return s.String()
}

func (s *GetDataDeliverHeaders) SetCommonHeaders(v map[string]*string) *GetDataDeliverHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDataDeliverHeaders) SetXAcsDingtalkAccessToken(v string) *GetDataDeliverHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDataDeliverRequest struct {
	// example:
	//
	// DELIVER-3e1a2d2f-fa76-45e8-XXXX-7fd29307c859
	DeliverId *string `json:"deliverId,omitempty" xml:"deliverId,omitempty"`
	// example:
	//
	// RT
	DispatchingItemType *string `json:"dispatchingItemType,omitempty" xml:"dispatchingItemType,omitempty"`
}

func (s GetDataDeliverRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDataDeliverRequest) GoString() string {
	return s.String()
}

func (s *GetDataDeliverRequest) SetDeliverId(v string) *GetDataDeliverRequest {
	s.DeliverId = &v
	return s
}

func (s *GetDataDeliverRequest) SetDispatchingItemType(v string) *GetDataDeliverRequest {
	s.DispatchingItemType = &v
	return s
}

type GetDataDeliverResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDataDeliverResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataDeliverResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataDeliverResponseBody) SetResult(v string) *GetDataDeliverResponseBody {
	s.Result = &v
	return s
}

func (s *GetDataDeliverResponseBody) SetSuccess(v bool) *GetDataDeliverResponseBody {
	s.Success = &v
	return s
}

type GetDataDeliverResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataDeliverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataDeliverResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataDeliverResponse) GoString() string {
	return s.String()
}

func (s *GetDataDeliverResponse) SetHeaders(v map[string]*string) *GetDataDeliverResponse {
	s.Headers = v
	return s
}

func (s *GetDataDeliverResponse) SetStatusCode(v int32) *GetDataDeliverResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataDeliverResponse) SetBody(v *GetDataDeliverResponseBody) *GetDataDeliverResponse {
	s.Body = v
	return s
}

type GetDomainInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDomainInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDomainInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetDomainInfoHeaders) SetCommonHeaders(v map[string]*string) *GetDomainInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDomainInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetDomainInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDomainInfoRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey,omitempty"`
}

func (s GetDomainInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDomainInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDomainInfoRequest) SetPageNumber(v int32) *GetDomainInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *GetDomainInfoRequest) SetPageSize(v int32) *GetDomainInfoRequest {
	s.PageSize = &v
	return s
}

func (s *GetDomainInfoRequest) SetSearchKey(v string) *GetDomainInfoRequest {
	s.SearchKey = &v
	return s
}

type GetDomainInfoResponseBody struct {
	// example:
	//
	// [{ "EntName":"企业名称" "Number":"备案号" "Domain":"域名" "SiteName":"网站名称" "HomeUrl":"网站首页链接" "CheckDate":"备案日期" }]
	Data  *string `json:"data,omitempty" xml:"data,omitempty"`
	Total *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetDomainInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDomainInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDomainInfoResponseBody) SetData(v string) *GetDomainInfoResponseBody {
	s.Data = &v
	return s
}

func (s *GetDomainInfoResponseBody) SetTotal(v int64) *GetDomainInfoResponseBody {
	s.Total = &v
	return s
}

type GetDomainInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDomainInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDomainInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDomainInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDomainInfoResponse) SetHeaders(v map[string]*string) *GetDomainInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDomainInfoResponse) SetStatusCode(v int32) *GetDomainInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDomainInfoResponse) SetBody(v *GetDomainInfoResponseBody) *GetDomainInfoResponse {
	s.Body = v
	return s
}

type GetDoubleRandomHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDoubleRandomHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDoubleRandomHeaders) GoString() string {
	return s.String()
}

func (s *GetDoubleRandomHeaders) SetCommonHeaders(v map[string]*string) *GetDoubleRandomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDoubleRandomHeaders) SetXAcsDingtalkAccessToken(v string) *GetDoubleRandomHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDoubleRandomRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey,omitempty"`
}

func (s GetDoubleRandomRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDoubleRandomRequest) GoString() string {
	return s.String()
}

func (s *GetDoubleRandomRequest) SetPageNumber(v int32) *GetDoubleRandomRequest {
	s.PageNumber = &v
	return s
}

func (s *GetDoubleRandomRequest) SetPageSize(v int32) *GetDoubleRandomRequest {
	s.PageSize = &v
	return s
}

func (s *GetDoubleRandomRequest) SetSearchKey(v string) *GetDoubleRandomRequest {
	s.SearchKey = &v
	return s
}

type GetDoubleRandomResponseBody struct {
	// example:
	//
	// [     {       "InspectPlanId": "44030020191021",       "InspectTypeName": "定向",       "InspectPlanName": "2019能效标识生产企业计量监督抽查1",       "InspectItem": "",       "InspectResult": "",       "InspectDepartment": "深圳市市场监督管理局龙岗局",       "InspectDate": "2019-10-14",       "InspectTaskId": "44030020191021",       "InspectTaskName": "2019能效标识生产企业计量监督抽查1"     }   ]
	Data  *string `json:"data,omitempty" xml:"data,omitempty"`
	Total *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetDoubleRandomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDoubleRandomResponseBody) GoString() string {
	return s.String()
}

func (s *GetDoubleRandomResponseBody) SetData(v string) *GetDoubleRandomResponseBody {
	s.Data = &v
	return s
}

func (s *GetDoubleRandomResponseBody) SetTotal(v int64) *GetDoubleRandomResponseBody {
	s.Total = &v
	return s
}

type GetDoubleRandomResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDoubleRandomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDoubleRandomResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDoubleRandomResponse) GoString() string {
	return s.String()
}

func (s *GetDoubleRandomResponse) SetHeaders(v map[string]*string) *GetDoubleRandomResponse {
	s.Headers = v
	return s
}

func (s *GetDoubleRandomResponse) SetStatusCode(v int32) *GetDoubleRandomResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDoubleRandomResponse) SetBody(v *GetDoubleRandomResponseBody) *GetDoubleRandomResponse {
	s.Body = v
	return s
}

type GetEnvironmentalPenaltiesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetEnvironmentalPenaltiesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentalPenaltiesHeaders) GoString() string {
	return s.String()
}

func (s *GetEnvironmentalPenaltiesHeaders) SetCommonHeaders(v map[string]*string) *GetEnvironmentalPenaltiesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetEnvironmentalPenaltiesHeaders) SetXAcsDingtalkAccessToken(v string) *GetEnvironmentalPenaltiesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetEnvironmentalPenaltiesRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey,omitempty"`
}

func (s GetEnvironmentalPenaltiesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentalPenaltiesRequest) GoString() string {
	return s.String()
}

func (s *GetEnvironmentalPenaltiesRequest) SetPageNumber(v int32) *GetEnvironmentalPenaltiesRequest {
	s.PageNumber = &v
	return s
}

func (s *GetEnvironmentalPenaltiesRequest) SetPageSize(v int32) *GetEnvironmentalPenaltiesRequest {
	s.PageSize = &v
	return s
}

func (s *GetEnvironmentalPenaltiesRequest) SetSearchKey(v string) *GetEnvironmentalPenaltiesRequest {
	s.SearchKey = &v
	return s
}

type GetEnvironmentalPenaltiesResponseBody struct {
	// example:
	//
	// [     {       "DEPARTMENT": "xx",       "ENT_NAME": "xx",       "EXEC_STATUS": "xx",       "PUNISH_BASIS": "xx",       "PUNISH_CONTENT": "xx",       "PUNISH_LAW": "xx",       "PUNISH_NUM": "xx",       "PUNISH_RES": "xx",       "PUNISH_DATE": "xx"      }   ]
	Data  *string `json:"data,omitempty" xml:"data,omitempty"`
	Total *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetEnvironmentalPenaltiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentalPenaltiesResponseBody) GoString() string {
	return s.String()
}

func (s *GetEnvironmentalPenaltiesResponseBody) SetData(v string) *GetEnvironmentalPenaltiesResponseBody {
	s.Data = &v
	return s
}

func (s *GetEnvironmentalPenaltiesResponseBody) SetTotal(v int64) *GetEnvironmentalPenaltiesResponseBody {
	s.Total = &v
	return s
}

type GetEnvironmentalPenaltiesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEnvironmentalPenaltiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEnvironmentalPenaltiesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentalPenaltiesResponse) GoString() string {
	return s.String()
}

func (s *GetEnvironmentalPenaltiesResponse) SetHeaders(v map[string]*string) *GetEnvironmentalPenaltiesResponse {
	s.Headers = v
	return s
}

func (s *GetEnvironmentalPenaltiesResponse) SetStatusCode(v int32) *GetEnvironmentalPenaltiesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEnvironmentalPenaltiesResponse) SetBody(v *GetEnvironmentalPenaltiesResponseBody) *GetEnvironmentalPenaltiesResponse {
	s.Body = v
	return s
}

type GetEventDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetEventDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetEventDataHeaders) GoString() string {
	return s.String()
}

func (s *GetEventDataHeaders) SetCommonHeaders(v map[string]*string) *GetEventDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetEventDataHeaders) SetXAcsDingtalkAccessToken(v string) *GetEventDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetEventDataRequest struct {
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 819e50d7c32e9096
	EventUid *string `json:"eventUid,omitempty" xml:"eventUid,omitempty"`
	SubId    *string `json:"subId,omitempty" xml:"subId,omitempty"`
}

func (s GetEventDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEventDataRequest) GoString() string {
	return s.String()
}

func (s *GetEventDataRequest) SetBizId(v string) *GetEventDataRequest {
	s.BizId = &v
	return s
}

func (s *GetEventDataRequest) SetEventUid(v string) *GetEventDataRequest {
	s.EventUid = &v
	return s
}

func (s *GetEventDataRequest) SetSubId(v string) *GetEventDataRequest {
	s.SubId = &v
	return s
}

type GetEventDataResponseBody struct {
	Success *string                `json:"success,omitempty" xml:"success,omitempty"`
	Value   map[string]interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetEventDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEventDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetEventDataResponseBody) SetSuccess(v string) *GetEventDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetEventDataResponseBody) SetValue(v map[string]interface{}) *GetEventDataResponseBody {
	s.Value = v
	return s
}

type GetEventDataResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEventDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEventDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEventDataResponse) GoString() string {
	return s.String()
}

func (s *GetEventDataResponse) SetHeaders(v map[string]*string) *GetEventDataResponse {
	s.Headers = v
	return s
}

func (s *GetEventDataResponse) SetStatusCode(v int32) *GetEventDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEventDataResponse) SetBody(v *GetEventDataResponseBody) *GetEventDataResponse {
	s.Body = v
	return s
}

type GetHolderInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetHolderInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetHolderInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetHolderInfoHeaders) SetCommonHeaders(v map[string]*string) *GetHolderInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHolderInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetHolderInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetHolderInfoRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey,omitempty"`
}

func (s GetHolderInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHolderInfoRequest) GoString() string {
	return s.String()
}

func (s *GetHolderInfoRequest) SetPageNumber(v int32) *GetHolderInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *GetHolderInfoRequest) SetPageSize(v int32) *GetHolderInfoRequest {
	s.PageSize = &v
	return s
}

func (s *GetHolderInfoRequest) SetSearchKey(v string) *GetHolderInfoRequest {
	s.SearchKey = &v
	return s
}

type GetHolderInfoResponseBody struct {
	// example:
	//
	// [     {       "STOCK_TYPE": "企业法人",       "STOCK_NAME": "xxx",       "STOCK_PERCENT": "100.00%",       "SHOULD_CAPI": "1000.0",       "SHOULD_CAPI_TIME": "2007-09-28"     }   ]
	Data  *string `json:"data,omitempty" xml:"data,omitempty"`
	Total *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetHolderInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHolderInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetHolderInfoResponseBody) SetData(v string) *GetHolderInfoResponseBody {
	s.Data = &v
	return s
}

func (s *GetHolderInfoResponseBody) SetTotal(v int64) *GetHolderInfoResponseBody {
	s.Total = &v
	return s
}

type GetHolderInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHolderInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHolderInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHolderInfoResponse) GoString() string {
	return s.String()
}

func (s *GetHolderInfoResponse) SetHeaders(v map[string]*string) *GetHolderInfoResponse {
	s.Headers = v
	return s
}

func (s *GetHolderInfoResponse) SetStatusCode(v int32) *GetHolderInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHolderInfoResponse) SetBody(v *GetHolderInfoResponseBody) *GetHolderInfoResponse {
	s.Body = v
	return s
}

type GetIntellectualPropertyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetIntellectualPropertyHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetIntellectualPropertyHeaders) GoString() string {
	return s.String()
}

func (s *GetIntellectualPropertyHeaders) SetCommonHeaders(v map[string]*string) *GetIntellectualPropertyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetIntellectualPropertyHeaders) SetXAcsDingtalkAccessToken(v string) *GetIntellectualPropertyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetIntellectualPropertyRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey,omitempty"`
}

func (s GetIntellectualPropertyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIntellectualPropertyRequest) GoString() string {
	return s.String()
}

func (s *GetIntellectualPropertyRequest) SetPageNumber(v int32) *GetIntellectualPropertyRequest {
	s.PageNumber = &v
	return s
}

func (s *GetIntellectualPropertyRequest) SetPageSize(v int32) *GetIntellectualPropertyRequest {
	s.PageSize = &v
	return s
}

func (s *GetIntellectualPropertyRequest) SetSearchKey(v string) *GetIntellectualPropertyRequest {
	s.SearchKey = &v
	return s
}

type GetIntellectualPropertyResponseBody struct {
	// example:
	//
	// [     {       "Status": "有效",       "Type": "专利",       "Pledgor": "齐风莲",       "Number": "91611024MA70X17M7E",       "Period": "2015年06月11日至2015年06月11日",       "PublicDate": "2015-06-18 00:00:00",       "Pawnee": "齐风莲",       "entName": "东兰县鸿发摩托车安全技术检验有限公司"     }   ]
	Data  *string `json:"data,omitempty" xml:"data,omitempty"`
	Total *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetIntellectualPropertyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIntellectualPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *GetIntellectualPropertyResponseBody) SetData(v string) *GetIntellectualPropertyResponseBody {
	s.Data = &v
	return s
}

func (s *GetIntellectualPropertyResponseBody) SetTotal(v int64) *GetIntellectualPropertyResponseBody {
	s.Total = &v
	return s
}

type GetIntellectualPropertyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIntellectualPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIntellectualPropertyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIntellectualPropertyResponse) GoString() string {
	return s.String()
}

func (s *GetIntellectualPropertyResponse) SetHeaders(v map[string]*string) *GetIntellectualPropertyResponse {
	s.Headers = v
	return s
}

func (s *GetIntellectualPropertyResponse) SetStatusCode(v int32) *GetIntellectualPropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIntellectualPropertyResponse) SetBody(v *GetIntellectualPropertyResponseBody) *GetIntellectualPropertyResponse {
	s.Body = v
	return s
}

type GetInvestmentAbroadHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetInvestmentAbroadHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetInvestmentAbroadHeaders) GoString() string {
	return s.String()
}

func (s *GetInvestmentAbroadHeaders) SetCommonHeaders(v map[string]*string) *GetInvestmentAbroadHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInvestmentAbroadHeaders) SetXAcsDingtalkAccessToken(v string) *GetInvestmentAbroadHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetInvestmentAbroadRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey,omitempty"`
}

func (s GetInvestmentAbroadRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInvestmentAbroadRequest) GoString() string {
	return s.String()
}

func (s *GetInvestmentAbroadRequest) SetPageNumber(v int32) *GetInvestmentAbroadRequest {
	s.PageNumber = &v
	return s
}

func (s *GetInvestmentAbroadRequest) SetPageSize(v int32) *GetInvestmentAbroadRequest {
	s.PageSize = &v
	return s
}

func (s *GetInvestmentAbroadRequest) SetSearchKey(v string) *GetInvestmentAbroadRequest {
	s.SearchKey = &v
	return s
}

type GetInvestmentAbroadResponseBody struct {
	// example:
	//
	// [     {       "InvestLicenseNo": "440301104818958",       "InvestStatus": "在营",       "InvestEsDate": "1998-11-25",       "InvestCreditCode": "914403007084643962",       "ShouldCap": "2000.0万人民币",       "EntName": "华为技术有限公司",       "InvestLegalName": "汤启兵",       "StockPercentage": "100.0%",       "InvestName": "深圳市华为技术服务有限公司",       "InvestRegCap": "2000.0万人民币"     }   ]
	Data  *string `json:"data,omitempty" xml:"data,omitempty"`
	Total *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetInvestmentAbroadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInvestmentAbroadResponseBody) GoString() string {
	return s.String()
}

func (s *GetInvestmentAbroadResponseBody) SetData(v string) *GetInvestmentAbroadResponseBody {
	s.Data = &v
	return s
}

func (s *GetInvestmentAbroadResponseBody) SetTotal(v int64) *GetInvestmentAbroadResponseBody {
	s.Total = &v
	return s
}

type GetInvestmentAbroadResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInvestmentAbroadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInvestmentAbroadResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInvestmentAbroadResponse) GoString() string {
	return s.String()
}

func (s *GetInvestmentAbroadResponse) SetHeaders(v map[string]*string) *GetInvestmentAbroadResponse {
	s.Headers = v
	return s
}

func (s *GetInvestmentAbroadResponse) SetStatusCode(v int32) *GetInvestmentAbroadResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInvestmentAbroadResponse) SetBody(v *GetInvestmentAbroadResponseBody) *GetInvestmentAbroadResponse {
	s.Body = v
	return s
}

type GetJobInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetJobInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetJobInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetJobInfoHeaders) SetCommonHeaders(v map[string]*string) *GetJobInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetJobInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetJobInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetJobInfoRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey,omitempty"`
}

func (s GetJobInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobInfoRequest) GoString() string {
	return s.String()
}

func (s *GetJobInfoRequest) SetPageNumber(v int32) *GetJobInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *GetJobInfoRequest) SetPageSize(v int32) *GetJobInfoRequest {
	s.PageSize = &v
	return s
}

func (s *GetJobInfoRequest) SetSearchKey(v string) *GetJobInfoRequest {
	s.SearchKey = &v
	return s
}

type GetJobInfoResponseBody struct {
	// example:
	//
	// [     {       "DEPARTMENT": "xx",       "IN_REASON": "xx",       "OUT_DATE": "2006-12-07",       "OUT_DEPARTMENT": "xx",       "IN_DATE": "2006-12-07",       "OUT_REASON": "xx"     }   ]
	Data  *string `json:"data,omitempty" xml:"data,omitempty"`
	Total *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetJobInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponseBody) SetData(v string) *GetJobInfoResponseBody {
	s.Data = &v
	return s
}

func (s *GetJobInfoResponseBody) SetTotal(v int64) *GetJobInfoResponseBody {
	s.Total = &v
	return s
}

type GetJobInfoResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobInfoResponse) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponse) SetHeaders(v map[string]*string) *GetJobInfoResponse {
	s.Headers = v
	return s
}

func (s *GetJobInfoResponse) SetStatusCode(v int32) *GetJobInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobInfoResponse) SetBody(v *GetJobInfoResponseBody) *GetJobInfoResponse {
	s.Body = v
	return s
}

type GetPatentInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPatentInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPatentInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetPatentInfoHeaders) SetCommonHeaders(v map[string]*string) *GetPatentInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPatentInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetPatentInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPatentInfoRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey,omitempty"`
}

func (s GetPatentInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPatentInfoRequest) GoString() string {
	return s.String()
}

func (s *GetPatentInfoRequest) SetPageNumber(v int32) *GetPatentInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *GetPatentInfoRequest) SetPageSize(v int32) *GetPatentInfoRequest {
	s.PageSize = &v
	return s
}

func (s *GetPatentInfoRequest) SetSearchKey(v string) *GetPatentInfoRequest {
	s.SearchKey = &v
	return s
}

type GetPatentInfoResponseBody struct {
	// example:
	//
	// [{"EntName":"企业名称", "PatentType":"专利类型", "PatentName":"专利名", "PatentStatus":"专利状态", "RequestNum":"申请号", "RequestDate":"申请日", "PublicNum":"公开(公告)号", "PublicDate":"公开(公告)日", "InventorList":"发明人", "PatenteeList":"专利权人", "CateNum":"分类号", "PrioNum":"优先权号", "PrioDate":"优先权日", "Agency":"专利代理机构", "Agent":"代理人", "Brief":"简要说明", "MainClaim":"主权项"}]
	Data  *string `json:"data,omitempty" xml:"data,omitempty"`
	Total *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetPatentInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPatentInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetPatentInfoResponseBody) SetData(v string) *GetPatentInfoResponseBody {
	s.Data = &v
	return s
}

func (s *GetPatentInfoResponseBody) SetTotal(v int64) *GetPatentInfoResponseBody {
	s.Total = &v
	return s
}

type GetPatentInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPatentInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPatentInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPatentInfoResponse) GoString() string {
	return s.String()
}

func (s *GetPatentInfoResponse) SetHeaders(v map[string]*string) *GetPatentInfoResponse {
	s.Headers = v
	return s
}

func (s *GetPatentInfoResponse) SetStatusCode(v int32) *GetPatentInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPatentInfoResponse) SetBody(v *GetPatentInfoResponseBody) *GetPatentInfoResponse {
	s.Body = v
	return s
}

type GetPrincipalEmployeeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPrincipalEmployeeHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPrincipalEmployeeHeaders) GoString() string {
	return s.String()
}

func (s *GetPrincipalEmployeeHeaders) SetCommonHeaders(v map[string]*string) *GetPrincipalEmployeeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPrincipalEmployeeHeaders) SetXAcsDingtalkAccessToken(v string) *GetPrincipalEmployeeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPrincipalEmployeeRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey,omitempty"`
}

func (s GetPrincipalEmployeeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPrincipalEmployeeRequest) GoString() string {
	return s.String()
}

func (s *GetPrincipalEmployeeRequest) SetPageNumber(v int32) *GetPrincipalEmployeeRequest {
	s.PageNumber = &v
	return s
}

func (s *GetPrincipalEmployeeRequest) SetPageSize(v int32) *GetPrincipalEmployeeRequest {
	s.PageSize = &v
	return s
}

func (s *GetPrincipalEmployeeRequest) SetSearchKey(v string) *GetPrincipalEmployeeRequest {
	s.SearchKey = &v
	return s
}

type GetPrincipalEmployeeResponseBody struct {
	// example:
	//
	// [     {       "JobTitle": "董事长",       "Name": "梁华"     },     {       "JobTitle": "副董事长",       "Name": "孟晚舟"     },     {       "JobTitle": "副董事长",       "Name": "徐直军"     }   ]
	Data  *string `json:"data,omitempty" xml:"data,omitempty"`
	Total *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetPrincipalEmployeeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPrincipalEmployeeResponseBody) GoString() string {
	return s.String()
}

func (s *GetPrincipalEmployeeResponseBody) SetData(v string) *GetPrincipalEmployeeResponseBody {
	s.Data = &v
	return s
}

func (s *GetPrincipalEmployeeResponseBody) SetTotal(v int64) *GetPrincipalEmployeeResponseBody {
	s.Total = &v
	return s
}

type GetPrincipalEmployeeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPrincipalEmployeeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPrincipalEmployeeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPrincipalEmployeeResponse) GoString() string {
	return s.String()
}

func (s *GetPrincipalEmployeeResponse) SetHeaders(v map[string]*string) *GetPrincipalEmployeeResponse {
	s.Headers = v
	return s
}

func (s *GetPrincipalEmployeeResponse) SetStatusCode(v int32) *GetPrincipalEmployeeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPrincipalEmployeeResponse) SetBody(v *GetPrincipalEmployeeResponseBody) *GetPrincipalEmployeeResponse {
	s.Body = v
	return s
}

type GetQeneralTaxpayerInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetQeneralTaxpayerInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetQeneralTaxpayerInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetQeneralTaxpayerInfoHeaders) SetCommonHeaders(v map[string]*string) *GetQeneralTaxpayerInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetQeneralTaxpayerInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetQeneralTaxpayerInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetQeneralTaxpayerInfoRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey,omitempty"`
}

func (s GetQeneralTaxpayerInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQeneralTaxpayerInfoRequest) GoString() string {
	return s.String()
}

func (s *GetQeneralTaxpayerInfoRequest) SetPageNumber(v int32) *GetQeneralTaxpayerInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *GetQeneralTaxpayerInfoRequest) SetPageSize(v int32) *GetQeneralTaxpayerInfoRequest {
	s.PageSize = &v
	return s
}

func (s *GetQeneralTaxpayerInfoRequest) SetSearchKey(v string) *GetQeneralTaxpayerInfoRequest {
	s.SearchKey = &v
	return s
}

type GetQeneralTaxpayerInfoResponseBody struct {
	// example:
	//
	// [     {       "DEPARTMENT":"xx"       "END_DATE":"2017-01-04"       "ENT_NAME":"xx"       "QUALIFICATION"       "START_DATE":"2017-01-04"       "TAXPAYER_NUM":"11"       "JUDGE_DATE":"2017-05-04"      }   ]
	Data  *string `json:"data,omitempty" xml:"data,omitempty"`
	Total *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetQeneralTaxpayerInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQeneralTaxpayerInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetQeneralTaxpayerInfoResponseBody) SetData(v string) *GetQeneralTaxpayerInfoResponseBody {
	s.Data = &v
	return s
}

func (s *GetQeneralTaxpayerInfoResponseBody) SetTotal(v int64) *GetQeneralTaxpayerInfoResponseBody {
	s.Total = &v
	return s
}

type GetQeneralTaxpayerInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQeneralTaxpayerInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQeneralTaxpayerInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQeneralTaxpayerInfoResponse) GoString() string {
	return s.String()
}

func (s *GetQeneralTaxpayerInfoResponse) SetHeaders(v map[string]*string) *GetQeneralTaxpayerInfoResponse {
	s.Headers = v
	return s
}

func (s *GetQeneralTaxpayerInfoResponse) SetStatusCode(v int32) *GetQeneralTaxpayerInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQeneralTaxpayerInfoResponse) SetBody(v *GetQeneralTaxpayerInfoResponseBody) *GetQeneralTaxpayerInfoResponse {
	s.Body = v
	return s
}

type GetQualificationCertHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetQualificationCertHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetQualificationCertHeaders) GoString() string {
	return s.String()
}

func (s *GetQualificationCertHeaders) SetCommonHeaders(v map[string]*string) *GetQualificationCertHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetQualificationCertHeaders) SetXAcsDingtalkAccessToken(v string) *GetQualificationCertHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetQualificationCertRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey,omitempty"`
}

func (s GetQualificationCertRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQualificationCertRequest) GoString() string {
	return s.String()
}

func (s *GetQualificationCertRequest) SetPageNumber(v int32) *GetQualificationCertRequest {
	s.PageNumber = &v
	return s
}

func (s *GetQualificationCertRequest) SetPageSize(v int32) *GetQualificationCertRequest {
	s.PageSize = &v
	return s
}

func (s *GetQualificationCertRequest) SetSearchKey(v string) *GetQualificationCertRequest {
	s.SearchKey = &v
	return s
}

type GetQualificationCertResponseBody struct {
	// example:
	//
	// [{"EntName":"企业名称", "CertType":"证书类型", "CertNum":"证书认证编号", "ValidStartDate":"有效期开始日期", "ValidEndDate":"有效期截止日期", "AuthorizeDate":"授权日期", "AuthorizeDepartment":"授权部门", "PubDate":"公示日期", "Province":"省份", "CertScope":"认证范围"} ]
	Data  *string `json:"data,omitempty" xml:"data,omitempty"`
	Total *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetQualificationCertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQualificationCertResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualificationCertResponseBody) SetData(v string) *GetQualificationCertResponseBody {
	s.Data = &v
	return s
}

func (s *GetQualificationCertResponseBody) SetTotal(v int64) *GetQualificationCertResponseBody {
	s.Total = &v
	return s
}

type GetQualificationCertResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualificationCertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualificationCertResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQualificationCertResponse) GoString() string {
	return s.String()
}

func (s *GetQualificationCertResponse) SetHeaders(v map[string]*string) *GetQualificationCertResponse {
	s.Headers = v
	return s
}

func (s *GetQualificationCertResponse) SetStatusCode(v int32) *GetQualificationCertResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualificationCertResponse) SetBody(v *GetQualificationCertResponseBody) *GetQualificationCertResponse {
	s.Body = v
	return s
}

type GetSeriousViolationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSeriousViolationHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSeriousViolationHeaders) GoString() string {
	return s.String()
}

func (s *GetSeriousViolationHeaders) SetCommonHeaders(v map[string]*string) *GetSeriousViolationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSeriousViolationHeaders) SetXAcsDingtalkAccessToken(v string) *GetSeriousViolationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSeriousViolationRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey,omitempty"`
}

func (s GetSeriousViolationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSeriousViolationRequest) GoString() string {
	return s.String()
}

func (s *GetSeriousViolationRequest) SetPageNumber(v int32) *GetSeriousViolationRequest {
	s.PageNumber = &v
	return s
}

func (s *GetSeriousViolationRequest) SetPageSize(v int32) *GetSeriousViolationRequest {
	s.PageSize = &v
	return s
}

func (s *GetSeriousViolationRequest) SetSearchKey(v string) *GetSeriousViolationRequest {
	s.SearchKey = &v
	return s
}

type GetSeriousViolationResponseBody struct {
	// example:
	//
	// [     {       "IN_DATE": "xx",       "IN_DEPARTMENT": "xx",       "IN_REASON": "xx"      }   ]
	Data  *string `json:"data,omitempty" xml:"data,omitempty"`
	Total *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetSeriousViolationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSeriousViolationResponseBody) GoString() string {
	return s.String()
}

func (s *GetSeriousViolationResponseBody) SetData(v string) *GetSeriousViolationResponseBody {
	s.Data = &v
	return s
}

func (s *GetSeriousViolationResponseBody) SetTotal(v int64) *GetSeriousViolationResponseBody {
	s.Total = &v
	return s
}

type GetSeriousViolationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSeriousViolationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSeriousViolationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSeriousViolationResponse) GoString() string {
	return s.String()
}

func (s *GetSeriousViolationResponse) SetHeaders(v map[string]*string) *GetSeriousViolationResponse {
	s.Headers = v
	return s
}

func (s *GetSeriousViolationResponse) SetStatusCode(v int32) *GetSeriousViolationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSeriousViolationResponse) SetBody(v *GetSeriousViolationResponseBody) *GetSeriousViolationResponse {
	s.Body = v
	return s
}

type GetSoftwareCopyrightHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSoftwareCopyrightHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSoftwareCopyrightHeaders) GoString() string {
	return s.String()
}

func (s *GetSoftwareCopyrightHeaders) SetCommonHeaders(v map[string]*string) *GetSoftwareCopyrightHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSoftwareCopyrightHeaders) SetXAcsDingtalkAccessToken(v string) *GetSoftwareCopyrightHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSoftwareCopyrightRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey,omitempty"`
}

func (s GetSoftwareCopyrightRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSoftwareCopyrightRequest) GoString() string {
	return s.String()
}

func (s *GetSoftwareCopyrightRequest) SetPageNumber(v int32) *GetSoftwareCopyrightRequest {
	s.PageNumber = &v
	return s
}

func (s *GetSoftwareCopyrightRequest) SetPageSize(v int32) *GetSoftwareCopyrightRequest {
	s.PageSize = &v
	return s
}

func (s *GetSoftwareCopyrightRequest) SetSearchKey(v string) *GetSoftwareCopyrightRequest {
	s.SearchKey = &v
	return s
}

type GetSoftwareCopyrightResponseBody struct {
	// example:
	//
	// [{ "EntName:企业名称", "CopyNum:登记号", "TypeNum:分类号", "ShortName:作品简称", "CopyName:作品全称", "Version:版本号", "SuccessDate:创作完成日期", "FirstDate:首次发表日期", "ApprovalDate:登记批准日期" }]
	Data  *string `json:"data,omitempty" xml:"data,omitempty"`
	Total *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetSoftwareCopyrightResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSoftwareCopyrightResponseBody) GoString() string {
	return s.String()
}

func (s *GetSoftwareCopyrightResponseBody) SetData(v string) *GetSoftwareCopyrightResponseBody {
	s.Data = &v
	return s
}

func (s *GetSoftwareCopyrightResponseBody) SetTotal(v int64) *GetSoftwareCopyrightResponseBody {
	s.Total = &v
	return s
}

type GetSoftwareCopyrightResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSoftwareCopyrightResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSoftwareCopyrightResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSoftwareCopyrightResponse) GoString() string {
	return s.String()
}

func (s *GetSoftwareCopyrightResponse) SetHeaders(v map[string]*string) *GetSoftwareCopyrightResponse {
	s.Headers = v
	return s
}

func (s *GetSoftwareCopyrightResponse) SetStatusCode(v int32) *GetSoftwareCopyrightResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSoftwareCopyrightResponse) SetBody(v *GetSoftwareCopyrightResponseBody) *GetSoftwareCopyrightResponse {
	s.Body = v
	return s
}

type GetTrademarkInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTrademarkInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTrademarkInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetTrademarkInfoHeaders) SetCommonHeaders(v map[string]*string) *GetTrademarkInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTrademarkInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetTrademarkInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTrademarkInfoRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey,omitempty"`
}

func (s GetTrademarkInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTrademarkInfoRequest) GoString() string {
	return s.String()
}

func (s *GetTrademarkInfoRequest) SetPageNumber(v int32) *GetTrademarkInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *GetTrademarkInfoRequest) SetPageSize(v int32) *GetTrademarkInfoRequest {
	s.PageSize = &v
	return s
}

func (s *GetTrademarkInfoRequest) SetSearchKey(v string) *GetTrademarkInfoRequest {
	s.SearchKey = &v
	return s
}

type GetTrademarkInfoResponseBody struct {
	// example:
	//
	// [{ "entName:企业名称", "trademarkName:商标名称", "regNum:商标注册号", "trademarkType:商标类型", "trademarkForm:商标形式", "trademarkStatus:商标状态", "applyDate:申请日期", "imageUrl:图片链接", "typeName:商标类型名", "period:专用权期限", "agent:代理人名称", "regPubNo:注册公告号", "regPubDate:注册公告日期", "firstPubNo:初审公告号", "firstPubDate:初审公告日期"}]
	Data  *string `json:"data,omitempty" xml:"data,omitempty"`
	Total *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetTrademarkInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrademarkInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrademarkInfoResponseBody) SetData(v string) *GetTrademarkInfoResponseBody {
	s.Data = &v
	return s
}

func (s *GetTrademarkInfoResponseBody) SetTotal(v int64) *GetTrademarkInfoResponseBody {
	s.Total = &v
	return s
}

type GetTrademarkInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrademarkInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrademarkInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrademarkInfoResponse) GoString() string {
	return s.String()
}

func (s *GetTrademarkInfoResponse) SetHeaders(v map[string]*string) *GetTrademarkInfoResponse {
	s.Headers = v
	return s
}

func (s *GetTrademarkInfoResponse) SetStatusCode(v int32) *GetTrademarkInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrademarkInfoResponse) SetBody(v *GetTrademarkInfoResponseBody) *GetTrademarkInfoResponse {
	s.Body = v
	return s
}

type GetWorkCopyrightHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetWorkCopyrightHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetWorkCopyrightHeaders) GoString() string {
	return s.String()
}

func (s *GetWorkCopyrightHeaders) SetCommonHeaders(v map[string]*string) *GetWorkCopyrightHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetWorkCopyrightHeaders) SetXAcsDingtalkAccessToken(v string) *GetWorkCopyrightHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetWorkCopyrightRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey,omitempty"`
}

func (s GetWorkCopyrightRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWorkCopyrightRequest) GoString() string {
	return s.String()
}

func (s *GetWorkCopyrightRequest) SetPageNumber(v int32) *GetWorkCopyrightRequest {
	s.PageNumber = &v
	return s
}

func (s *GetWorkCopyrightRequest) SetPageSize(v int32) *GetWorkCopyrightRequest {
	s.PageSize = &v
	return s
}

func (s *GetWorkCopyrightRequest) SetSearchKey(v string) *GetWorkCopyrightRequest {
	s.SearchKey = &v
	return s
}

type GetWorkCopyrightResponseBody struct {
	// example:
	//
	// [{ "EntName":"企业名称", "CopyName":"作品全称", "TypeName":"作品类别", "CopyNum":"登记号", "SuccessDate":"创作完成日期", "FirstDate":"首次发表日期", "ApprovalDate":"登记批准日期" }]
	Data  *string `json:"data,omitempty" xml:"data,omitempty"`
	Total *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetWorkCopyrightResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkCopyrightResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkCopyrightResponseBody) SetData(v string) *GetWorkCopyrightResponseBody {
	s.Data = &v
	return s
}

func (s *GetWorkCopyrightResponseBody) SetTotal(v int64) *GetWorkCopyrightResponseBody {
	s.Total = &v
	return s
}

type GetWorkCopyrightResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkCopyrightResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkCopyrightResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkCopyrightResponse) GoString() string {
	return s.String()
}

func (s *GetWorkCopyrightResponse) SetHeaders(v map[string]*string) *GetWorkCopyrightResponse {
	s.Headers = v
	return s
}

func (s *GetWorkCopyrightResponse) SetStatusCode(v int32) *GetWorkCopyrightResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkCopyrightResponse) SetBody(v *GetWorkCopyrightResponseBody) *GetWorkCopyrightResponse {
	s.Body = v
	return s
}

type ListDataDeliversHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListDataDeliversHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListDataDeliversHeaders) GoString() string {
	return s.String()
}

func (s *ListDataDeliversHeaders) SetCommonHeaders(v map[string]*string) *ListDataDeliversHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDataDeliversHeaders) SetXAcsDingtalkAccessToken(v string) *ListDataDeliversHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListDataDeliversRequest struct {
	// example:
	//
	// RT
	DispatchingItemType *string `json:"dispatchingItemType,omitempty" xml:"dispatchingItemType,omitempty"`
}

func (s ListDataDeliversRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataDeliversRequest) GoString() string {
	return s.String()
}

func (s *ListDataDeliversRequest) SetDispatchingItemType(v string) *ListDataDeliversRequest {
	s.DispatchingItemType = &v
	return s
}

type ListDataDeliversResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListDataDeliversResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataDeliversResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataDeliversResponseBody) SetResult(v string) *ListDataDeliversResponseBody {
	s.Result = &v
	return s
}

func (s *ListDataDeliversResponseBody) SetSuccess(v bool) *ListDataDeliversResponseBody {
	s.Success = &v
	return s
}

type ListDataDeliversResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataDeliversResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataDeliversResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataDeliversResponse) GoString() string {
	return s.String()
}

func (s *ListDataDeliversResponse) SetHeaders(v map[string]*string) *ListDataDeliversResponse {
	s.Headers = v
	return s
}

func (s *ListDataDeliversResponse) SetStatusCode(v int32) *ListDataDeliversResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataDeliversResponse) SetBody(v *ListDataDeliversResponseBody) *ListDataDeliversResponse {
	s.Body = v
	return s
}

type OperateChartConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s OperateChartConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s OperateChartConfigHeaders) GoString() string {
	return s.String()
}

func (s *OperateChartConfigHeaders) SetCommonHeaders(v map[string]*string) *OperateChartConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OperateChartConfigHeaders) SetXAcsDingtalkAccessToken(v string) *OperateChartConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type OperateChartConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	CorpId *string                `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Param  map[string]interface{} `json:"param,omitempty" xml:"param,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8ABvoWxoelSxcxZBsF3MeWBDe5oi8jmFtU790jhpRoLrfJDWO8UDHbUqvTb3pQA5
	Ticket *string `json:"ticket,omitempty" xml:"ticket,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s OperateChartConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateChartConfigRequest) GoString() string {
	return s.String()
}

func (s *OperateChartConfigRequest) SetApiKey(v string) *OperateChartConfigRequest {
	s.ApiKey = &v
	return s
}

func (s *OperateChartConfigRequest) SetCorpId(v string) *OperateChartConfigRequest {
	s.CorpId = &v
	return s
}

func (s *OperateChartConfigRequest) SetParam(v map[string]interface{}) *OperateChartConfigRequest {
	s.Param = v
	return s
}

func (s *OperateChartConfigRequest) SetTicket(v string) *OperateChartConfigRequest {
	s.Ticket = &v
	return s
}

func (s *OperateChartConfigRequest) SetUserId(v string) *OperateChartConfigRequest {
	s.UserId = &v
	return s
}

type OperateChartConfigResponseBody struct {
	Result  map[string]*string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OperateChartConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateChartConfigResponseBody) GoString() string {
	return s.String()
}

func (s *OperateChartConfigResponseBody) SetResult(v map[string]*string) *OperateChartConfigResponseBody {
	s.Result = v
	return s
}

func (s *OperateChartConfigResponseBody) SetSuccess(v bool) *OperateChartConfigResponseBody {
	s.Success = &v
	return s
}

type OperateChartConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateChartConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateChartConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateChartConfigResponse) GoString() string {
	return s.String()
}

func (s *OperateChartConfigResponse) SetHeaders(v map[string]*string) *OperateChartConfigResponse {
	s.Headers = v
	return s
}

func (s *OperateChartConfigResponse) SetStatusCode(v int32) *OperateChartConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateChartConfigResponse) SetBody(v *OperateChartConfigResponseBody) *OperateChartConfigResponse {
	s.Body = v
	return s
}

type PostCorpAuthInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PostCorpAuthInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s PostCorpAuthInfoHeaders) GoString() string {
	return s.String()
}

func (s *PostCorpAuthInfoHeaders) SetCommonHeaders(v map[string]*string) *PostCorpAuthInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PostCorpAuthInfoHeaders) SetXAcsDingtalkAccessToken(v string) *PostCorpAuthInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PostCorpAuthInfoResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PostCorpAuthInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PostCorpAuthInfoResponseBody) GoString() string {
	return s.String()
}

func (s *PostCorpAuthInfoResponseBody) SetSuccess(v bool) *PostCorpAuthInfoResponseBody {
	s.Success = &v
	return s
}

type PostCorpAuthInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostCorpAuthInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostCorpAuthInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s PostCorpAuthInfoResponse) GoString() string {
	return s.String()
}

func (s *PostCorpAuthInfoResponse) SetHeaders(v map[string]*string) *PostCorpAuthInfoResponse {
	s.Headers = v
	return s
}

func (s *PostCorpAuthInfoResponse) SetStatusCode(v int32) *PostCorpAuthInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *PostCorpAuthInfoResponse) SetBody(v *PostCorpAuthInfoResponseBody) *PostCorpAuthInfoResponse {
	s.Body = v
	return s
}

type QueryActiveUserStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryActiveUserStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryActiveUserStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryActiveUserStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryActiveUserStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryActiveUserStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryActiveUserStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryActiveUserStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryActiveUserStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryActiveUserStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryActiveUserStatisticalDataRequest) SetStatDate(v string) *QueryActiveUserStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryActiveUserStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                              `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryActiveUserStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryActiveUserStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryActiveUserStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryActiveUserStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryActiveUserStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryActiveUserStatisticalDataResponseBody) SetMetaList(v []*QueryActiveUserStatisticalDataResponseBodyMetaList) *QueryActiveUserStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryActiveUserStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryActiveUserStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryActiveUserStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryActiveUserStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryActiveUserStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryActiveUserStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryActiveUserStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryActiveUserStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryActiveUserStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryActiveUserStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryActiveUserStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryActiveUserStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryActiveUserStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryActiveUserStatisticalDataResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryActiveUserStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryActiveUserStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryActiveUserStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryActiveUserStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryActiveUserStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryActiveUserStatisticalDataResponse) SetStatusCode(v int32) *QueryActiveUserStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryActiveUserStatisticalDataResponse) SetBody(v *QueryActiveUserStatisticalDataResponseBody) *QueryActiveUserStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryAnhmdStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAnhmdStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAnhmdStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryAnhmdStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryAnhmdStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAnhmdStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAnhmdStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAnhmdStatisticalDataRequest struct {
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryAnhmdStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAnhmdStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryAnhmdStatisticalDataRequest) SetPageNumber(v int64) *QueryAnhmdStatisticalDataRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryAnhmdStatisticalDataRequest) SetPageSize(v int64) *QueryAnhmdStatisticalDataRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAnhmdStatisticalDataRequest) SetStatDate(v string) *QueryAnhmdStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryAnhmdStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                         `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryAnhmdStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryAnhmdStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAnhmdStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAnhmdStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryAnhmdStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryAnhmdStatisticalDataResponseBody) SetMetaList(v []*QueryAnhmdStatisticalDataResponseBodyMetaList) *QueryAnhmdStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryAnhmdStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryAnhmdStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryAnhmdStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryAnhmdStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryAnhmdStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryAnhmdStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryAnhmdStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryAnhmdStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryAnhmdStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryAnhmdStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryAnhmdStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryAnhmdStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryAnhmdStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryAnhmdStatisticalDataResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAnhmdStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAnhmdStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAnhmdStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryAnhmdStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryAnhmdStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryAnhmdStatisticalDataResponse) SetStatusCode(v int32) *QueryAnhmdStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAnhmdStatisticalDataResponse) SetBody(v *QueryAnhmdStatisticalDataResponseBody) *QueryAnhmdStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryApprovalStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryApprovalStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryApprovalStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryApprovalStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryApprovalStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryApprovalStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryApprovalStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryApprovalStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryApprovalStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryApprovalStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryApprovalStatisticalDataRequest) SetStatDate(v string) *QueryApprovalStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryApprovalStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                            `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryApprovalStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryApprovalStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryApprovalStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryApprovalStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryApprovalStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryApprovalStatisticalDataResponseBody) SetMetaList(v []*QueryApprovalStatisticalDataResponseBodyMetaList) *QueryApprovalStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryApprovalStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryApprovalStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryApprovalStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryApprovalStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryApprovalStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryApprovalStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryApprovalStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryApprovalStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryApprovalStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryApprovalStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryApprovalStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryApprovalStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryApprovalStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryApprovalStatisticalDataResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryApprovalStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryApprovalStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryApprovalStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryApprovalStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryApprovalStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryApprovalStatisticalDataResponse) SetStatusCode(v int32) *QueryApprovalStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryApprovalStatisticalDataResponse) SetBody(v *QueryApprovalStatisticalDataResponseBody) *QueryApprovalStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryAttendanceStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAttendanceStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAttendanceStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryAttendanceStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryAttendanceStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAttendanceStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAttendanceStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAttendanceStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryAttendanceStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAttendanceStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryAttendanceStatisticalDataRequest) SetStatDate(v string) *QueryAttendanceStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryAttendanceStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                              `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryAttendanceStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryAttendanceStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAttendanceStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAttendanceStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryAttendanceStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryAttendanceStatisticalDataResponseBody) SetMetaList(v []*QueryAttendanceStatisticalDataResponseBodyMetaList) *QueryAttendanceStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryAttendanceStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryAttendanceStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryAttendanceStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryAttendanceStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryAttendanceStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryAttendanceStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryAttendanceStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryAttendanceStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryAttendanceStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryAttendanceStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryAttendanceStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryAttendanceStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryAttendanceStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryAttendanceStatisticalDataResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAttendanceStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAttendanceStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAttendanceStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryAttendanceStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryAttendanceStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryAttendanceStatisticalDataResponse) SetStatusCode(v int32) *QueryAttendanceStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAttendanceStatisticalDataResponse) SetBody(v *QueryAttendanceStatisticalDataResponseBody) *QueryAttendanceStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryBlackboardStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryBlackboardStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryBlackboardStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryBlackboardStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryBlackboardStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryBlackboardStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryBlackboardStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryBlackboardStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryBlackboardStatisticalDataRequest) SetStatDate(v string) *QueryBlackboardStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryBlackboardStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                              `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryBlackboardStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryBlackboardStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBlackboardStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryBlackboardStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryBlackboardStatisticalDataResponseBody) SetMetaList(v []*QueryBlackboardStatisticalDataResponseBodyMetaList) *QueryBlackboardStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryBlackboardStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryBlackboardStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryBlackboardStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryBlackboardStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryBlackboardStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryBlackboardStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryBlackboardStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryBlackboardStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryBlackboardStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryBlackboardStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryBlackboardStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryBlackboardStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryBlackboardStatisticalDataResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBlackboardStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBlackboardStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryBlackboardStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryBlackboardStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryBlackboardStatisticalDataResponse) SetStatusCode(v int32) *QueryBlackboardStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBlackboardStatisticalDataResponse) SetBody(v *QueryBlackboardStatisticalDataResponseBody) *QueryBlackboardStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryCalendarStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCalendarStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCalendarStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryCalendarStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryCalendarStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCalendarStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCalendarStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCalendarStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryCalendarStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCalendarStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryCalendarStatisticalDataRequest) SetStatDate(v string) *QueryCalendarStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryCalendarStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                            `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryCalendarStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryCalendarStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCalendarStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCalendarStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryCalendarStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryCalendarStatisticalDataResponseBody) SetMetaList(v []*QueryCalendarStatisticalDataResponseBodyMetaList) *QueryCalendarStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryCalendarStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryCalendarStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryCalendarStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryCalendarStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryCalendarStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryCalendarStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryCalendarStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryCalendarStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryCalendarStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryCalendarStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryCalendarStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryCalendarStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryCalendarStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryCalendarStatisticalDataResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCalendarStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCalendarStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCalendarStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryCalendarStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryCalendarStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryCalendarStatisticalDataResponse) SetStatusCode(v int32) *QueryCalendarStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCalendarStatisticalDataResponse) SetBody(v *QueryCalendarStatisticalDataResponseBody) *QueryCalendarStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryChartDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryChartDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryChartDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryChartDataHeaders) SetCommonHeaders(v map[string]*string) *QueryChartDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryChartDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryChartDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryChartDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ABC
	Ticket *string `json:"ticket,omitempty" xml:"ticket,omitempty"`
}

func (s QueryChartDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryChartDataRequest) GoString() string {
	return s.String()
}

func (s *QueryChartDataRequest) SetCode(v string) *QueryChartDataRequest {
	s.Code = &v
	return s
}

func (s *QueryChartDataRequest) SetTicket(v string) *QueryChartDataRequest {
	s.Ticket = &v
	return s
}

type QueryChartDataResponseBody struct {
	Result  []interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success *string       `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryChartDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryChartDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryChartDataResponseBody) SetResult(v []interface{}) *QueryChartDataResponseBody {
	s.Result = v
	return s
}

func (s *QueryChartDataResponseBody) SetSuccess(v string) *QueryChartDataResponseBody {
	s.Success = &v
	return s
}

type QueryChartDataResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryChartDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryChartDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryChartDataResponse) GoString() string {
	return s.String()
}

func (s *QueryChartDataResponse) SetHeaders(v map[string]*string) *QueryChartDataResponse {
	s.Headers = v
	return s
}

func (s *QueryChartDataResponse) SetStatusCode(v int32) *QueryChartDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryChartDataResponse) SetBody(v *QueryChartDataResponseBody) *QueryChartDataResponse {
	s.Body = v
	return s
}

type QueryCheckinStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCheckinStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCheckinStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryCheckinStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryCheckinStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCheckinStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCheckinStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCheckinStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryCheckinStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCheckinStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryCheckinStatisticalDataRequest) SetStatDate(v string) *QueryCheckinStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryCheckinStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                           `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryCheckinStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryCheckinStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCheckinStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCheckinStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryCheckinStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryCheckinStatisticalDataResponseBody) SetMetaList(v []*QueryCheckinStatisticalDataResponseBodyMetaList) *QueryCheckinStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryCheckinStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryCheckinStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryCheckinStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryCheckinStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryCheckinStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryCheckinStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryCheckinStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryCheckinStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryCheckinStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryCheckinStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryCheckinStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryCheckinStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryCheckinStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryCheckinStatisticalDataResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCheckinStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCheckinStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCheckinStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryCheckinStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryCheckinStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryCheckinStatisticalDataResponse) SetStatusCode(v int32) *QueryCheckinStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCheckinStatisticalDataResponse) SetBody(v *QueryCheckinStatisticalDataResponseBody) *QueryCheckinStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryCircleStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCircleStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCircleStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryCircleStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryCircleStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCircleStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCircleStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCircleStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryCircleStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCircleStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryCircleStatisticalDataRequest) SetStatDate(v string) *QueryCircleStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryCircleStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                          `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryCircleStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryCircleStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCircleStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCircleStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryCircleStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryCircleStatisticalDataResponseBody) SetMetaList(v []*QueryCircleStatisticalDataResponseBodyMetaList) *QueryCircleStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryCircleStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryCircleStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryCircleStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryCircleStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryCircleStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryCircleStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryCircleStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryCircleStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryCircleStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryCircleStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryCircleStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryCircleStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryCircleStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryCircleStatisticalDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCircleStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCircleStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCircleStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryCircleStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryCircleStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryCircleStatisticalDataResponse) SetStatusCode(v int32) *QueryCircleStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCircleStatisticalDataResponse) SetBody(v *QueryCircleStatisticalDataResponseBody) *QueryCircleStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryCompanyBasicInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCompanyBasicInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCompanyBasicInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryCompanyBasicInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryCompanyBasicInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCompanyBasicInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCompanyBasicInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCompanyBasicInfoRequest struct {
	// This parameter is required.
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// This parameter is required.
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryCompanyBasicInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCompanyBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryCompanyBasicInfoRequest) SetKeyword(v string) *QueryCompanyBasicInfoRequest {
	s.Keyword = &v
	return s
}

func (s *QueryCompanyBasicInfoRequest) SetPageNumber(v int64) *QueryCompanyBasicInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryCompanyBasicInfoRequest) SetPageSize(v int64) *QueryCompanyBasicInfoRequest {
	s.PageSize = &v
	return s
}

type QueryCompanyBasicInfoResponseBody struct {
	Code      *string              `json:"code,omitempty" xml:"code,omitempty"`
	Data      []map[string]*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Message   *string              `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Total     *int32               `json:"total,omitempty" xml:"total,omitempty"`
}

func (s QueryCompanyBasicInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCompanyBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCompanyBasicInfoResponseBody) SetCode(v string) *QueryCompanyBasicInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCompanyBasicInfoResponseBody) SetData(v []map[string]*string) *QueryCompanyBasicInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryCompanyBasicInfoResponseBody) SetMessage(v string) *QueryCompanyBasicInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCompanyBasicInfoResponseBody) SetRequestId(v string) *QueryCompanyBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCompanyBasicInfoResponseBody) SetTotal(v int32) *QueryCompanyBasicInfoResponseBody {
	s.Total = &v
	return s
}

type QueryCompanyBasicInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCompanyBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCompanyBasicInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCompanyBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryCompanyBasicInfoResponse) SetHeaders(v map[string]*string) *QueryCompanyBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryCompanyBasicInfoResponse) SetStatusCode(v int32) *QueryCompanyBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCompanyBasicInfoResponse) SetBody(v *QueryCompanyBasicInfoResponseBody) *QueryCompanyBasicInfoResponse {
	s.Body = v
	return s
}

type QueryDigitalDistrictOrgInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDigitalDistrictOrgInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDigitalDistrictOrgInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryDigitalDistrictOrgInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryDigitalDistrictOrgInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDigitalDistrictOrgInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDigitalDistrictOrgInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDigitalDistrictOrgInfoRequest struct {
	// This parameter is required.
	CorpIds []*string `json:"corpIds,omitempty" xml:"corpIds,omitempty" type:"Repeated"`
	// This parameter is required.
	StatDates []*string `json:"statDates,omitempty" xml:"statDates,omitempty" type:"Repeated"`
}

func (s QueryDigitalDistrictOrgInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDigitalDistrictOrgInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryDigitalDistrictOrgInfoRequest) SetCorpIds(v []*string) *QueryDigitalDistrictOrgInfoRequest {
	s.CorpIds = v
	return s
}

func (s *QueryDigitalDistrictOrgInfoRequest) SetStatDates(v []*string) *QueryDigitalDistrictOrgInfoRequest {
	s.StatDates = v
	return s
}

type QueryDigitalDistrictOrgInfoResponseBody struct {
	Arguments []*string `json:"arguments,omitempty" xml:"arguments,omitempty" type:"Repeated"`
	Result    *string   `json:"result,omitempty" xml:"result,omitempty"`
}

func (s QueryDigitalDistrictOrgInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDigitalDistrictOrgInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDigitalDistrictOrgInfoResponseBody) SetArguments(v []*string) *QueryDigitalDistrictOrgInfoResponseBody {
	s.Arguments = v
	return s
}

func (s *QueryDigitalDistrictOrgInfoResponseBody) SetResult(v string) *QueryDigitalDistrictOrgInfoResponseBody {
	s.Result = &v
	return s
}

type QueryDigitalDistrictOrgInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDigitalDistrictOrgInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDigitalDistrictOrgInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDigitalDistrictOrgInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryDigitalDistrictOrgInfoResponse) SetHeaders(v map[string]*string) *QueryDigitalDistrictOrgInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryDigitalDistrictOrgInfoResponse) SetStatusCode(v int32) *QueryDigitalDistrictOrgInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDigitalDistrictOrgInfoResponse) SetBody(v *QueryDigitalDistrictOrgInfoResponseBody) *QueryDigitalDistrictOrgInfoResponse {
	s.Body = v
	return s
}

type QueryDingReciveStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDingReciveStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDingReciveStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryDingReciveStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryDingReciveStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDingReciveStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDingReciveStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDingReciveStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryDingReciveStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDingReciveStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryDingReciveStatisticalDataRequest) SetStatDate(v string) *QueryDingReciveStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryDingReciveStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                              `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryDingReciveStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryDingReciveStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDingReciveStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDingReciveStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryDingReciveStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryDingReciveStatisticalDataResponseBody) SetMetaList(v []*QueryDingReciveStatisticalDataResponseBodyMetaList) *QueryDingReciveStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryDingReciveStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryDingReciveStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryDingReciveStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryDingReciveStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryDingReciveStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryDingReciveStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryDingReciveStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryDingReciveStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryDingReciveStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryDingReciveStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryDingReciveStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryDingReciveStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryDingReciveStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryDingReciveStatisticalDataResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDingReciveStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDingReciveStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDingReciveStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryDingReciveStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryDingReciveStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryDingReciveStatisticalDataResponse) SetStatusCode(v int32) *QueryDingReciveStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDingReciveStatisticalDataResponse) SetBody(v *QueryDingReciveStatisticalDataResponseBody) *QueryDingReciveStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryDingSendStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDingSendStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDingSendStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryDingSendStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryDingSendStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDingSendStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDingSendStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDingSendStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryDingSendStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDingSendStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryDingSendStatisticalDataRequest) SetStatDate(v string) *QueryDingSendStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryDingSendStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                            `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryDingSendStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryDingSendStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDingSendStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDingSendStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryDingSendStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryDingSendStatisticalDataResponseBody) SetMetaList(v []*QueryDingSendStatisticalDataResponseBodyMetaList) *QueryDingSendStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryDingSendStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryDingSendStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryDingSendStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryDingSendStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryDingSendStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryDingSendStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryDingSendStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryDingSendStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryDingSendStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryDingSendStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryDingSendStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryDingSendStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryDingSendStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryDingSendStatisticalDataResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDingSendStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDingSendStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDingSendStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryDingSendStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryDingSendStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryDingSendStatisticalDataResponse) SetStatusCode(v int32) *QueryDingSendStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDingSendStatisticalDataResponse) SetBody(v *QueryDingSendStatisticalDataResponseBody) *QueryDingSendStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryDocumentStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDocumentStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDocumentStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryDocumentStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryDocumentStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDocumentStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDocumentStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDocumentStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryDocumentStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDocumentStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryDocumentStatisticalDataRequest) SetStatDate(v string) *QueryDocumentStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryDocumentStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                            `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryDocumentStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryDocumentStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDocumentStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDocumentStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryDocumentStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryDocumentStatisticalDataResponseBody) SetMetaList(v []*QueryDocumentStatisticalDataResponseBodyMetaList) *QueryDocumentStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryDocumentStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryDocumentStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryDocumentStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryDocumentStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryDocumentStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryDocumentStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryDocumentStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryDocumentStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryDocumentStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryDocumentStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryDocumentStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryDocumentStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryDocumentStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryDocumentStatisticalDataResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDocumentStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDocumentStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDocumentStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryDocumentStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryDocumentStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryDocumentStatisticalDataResponse) SetStatusCode(v int32) *QueryDocumentStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDocumentStatisticalDataResponse) SetBody(v *QueryDocumentStatisticalDataResponseBody) *QueryDocumentStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryDpaasDataPackageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDpaasDataPackageHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDpaasDataPackageHeaders) GoString() string {
	return s.String()
}

func (s *QueryDpaasDataPackageHeaders) SetCommonHeaders(v map[string]*string) *QueryDpaasDataPackageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDpaasDataPackageHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDpaasDataPackageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDpaasDataPackageResponseBody struct {
	Buy           *bool   `json:"buy,omitempty" xml:"buy,omitempty"`
	EndDate       *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	Quota         *int64  `json:"quota,omitempty" xml:"quota,omitempty"`
	StartDate     *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	Success       *bool   `json:"success,omitempty" xml:"success,omitempty"`
	UsedNum       *int64  `json:"usedNum,omitempty" xml:"usedNum,omitempty"`
	WhiteCustomer *bool   `json:"whiteCustomer,omitempty" xml:"whiteCustomer,omitempty"`
}

func (s QueryDpaasDataPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDpaasDataPackageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDpaasDataPackageResponseBody) SetBuy(v bool) *QueryDpaasDataPackageResponseBody {
	s.Buy = &v
	return s
}

func (s *QueryDpaasDataPackageResponseBody) SetEndDate(v string) *QueryDpaasDataPackageResponseBody {
	s.EndDate = &v
	return s
}

func (s *QueryDpaasDataPackageResponseBody) SetQuota(v int64) *QueryDpaasDataPackageResponseBody {
	s.Quota = &v
	return s
}

func (s *QueryDpaasDataPackageResponseBody) SetStartDate(v string) *QueryDpaasDataPackageResponseBody {
	s.StartDate = &v
	return s
}

func (s *QueryDpaasDataPackageResponseBody) SetSuccess(v bool) *QueryDpaasDataPackageResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDpaasDataPackageResponseBody) SetUsedNum(v int64) *QueryDpaasDataPackageResponseBody {
	s.UsedNum = &v
	return s
}

func (s *QueryDpaasDataPackageResponseBody) SetWhiteCustomer(v bool) *QueryDpaasDataPackageResponseBody {
	s.WhiteCustomer = &v
	return s
}

type QueryDpaasDataPackageResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDpaasDataPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDpaasDataPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDpaasDataPackageResponse) GoString() string {
	return s.String()
}

func (s *QueryDpaasDataPackageResponse) SetHeaders(v map[string]*string) *QueryDpaasDataPackageResponse {
	s.Headers = v
	return s
}

func (s *QueryDpaasDataPackageResponse) SetStatusCode(v int32) *QueryDpaasDataPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDpaasDataPackageResponse) SetBody(v *QueryDpaasDataPackageResponseBody) *QueryDpaasDataPackageResponse {
	s.Body = v
	return s
}

type QueryDriveStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDriveStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDriveStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryDriveStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryDriveStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDriveStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDriveStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDriveStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryDriveStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDriveStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryDriveStatisticalDataRequest) SetStatDate(v string) *QueryDriveStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryDriveStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                         `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryDriveStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryDriveStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDriveStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDriveStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryDriveStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryDriveStatisticalDataResponseBody) SetMetaList(v []*QueryDriveStatisticalDataResponseBodyMetaList) *QueryDriveStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryDriveStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryDriveStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryDriveStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryDriveStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryDriveStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryDriveStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryDriveStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryDriveStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryDriveStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryDriveStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryDriveStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryDriveStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryDriveStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryDriveStatisticalDataResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDriveStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDriveStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDriveStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryDriveStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryDriveStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryDriveStatisticalDataResponse) SetStatusCode(v int32) *QueryDriveStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDriveStatisticalDataResponse) SetBody(v *QueryDriveStatisticalDataResponseBody) *QueryDriveStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryEduUnionAuthServiceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryEduUnionAuthServiceHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryEduUnionAuthServiceHeaders) GoString() string {
	return s.String()
}

func (s *QueryEduUnionAuthServiceHeaders) SetCommonHeaders(v map[string]*string) *QueryEduUnionAuthServiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryEduUnionAuthServiceHeaders) SetXAcsDingtalkAccessToken(v string) *QueryEduUnionAuthServiceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryEduUnionAuthServiceRequest struct {
	Ds *string `json:"ds,omitempty" xml:"ds,omitempty"`
}

func (s QueryEduUnionAuthServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEduUnionAuthServiceRequest) GoString() string {
	return s.String()
}

func (s *QueryEduUnionAuthServiceRequest) SetDs(v string) *QueryEduUnionAuthServiceRequest {
	s.Ds = &v
	return s
}

type QueryEduUnionAuthServiceResponseBody struct {
	AuthInfoModels []*QueryEduUnionAuthServiceResponseBodyAuthInfoModels `json:"authInfoModels,omitempty" xml:"authInfoModels,omitempty" type:"Repeated"`
}

func (s QueryEduUnionAuthServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEduUnionAuthServiceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEduUnionAuthServiceResponseBody) SetAuthInfoModels(v []*QueryEduUnionAuthServiceResponseBodyAuthInfoModels) *QueryEduUnionAuthServiceResponseBody {
	s.AuthInfoModels = v
	return s
}

type QueryEduUnionAuthServiceResponseBodyAuthInfoModels struct {
	AuthCorpId    *string   `json:"authCorpId,omitempty" xml:"authCorpId,omitempty"`
	AuthCorpName  *string   `json:"authCorpName,omitempty" xml:"authCorpName,omitempty"`
	AuthTime      *string   `json:"authTime,omitempty" xml:"authTime,omitempty"`
	ResourceNames []*string `json:"resourceNames,omitempty" xml:"resourceNames,omitempty" type:"Repeated"`
}

func (s QueryEduUnionAuthServiceResponseBodyAuthInfoModels) String() string {
	return tea.Prettify(s)
}

func (s QueryEduUnionAuthServiceResponseBodyAuthInfoModels) GoString() string {
	return s.String()
}

func (s *QueryEduUnionAuthServiceResponseBodyAuthInfoModels) SetAuthCorpId(v string) *QueryEduUnionAuthServiceResponseBodyAuthInfoModels {
	s.AuthCorpId = &v
	return s
}

func (s *QueryEduUnionAuthServiceResponseBodyAuthInfoModels) SetAuthCorpName(v string) *QueryEduUnionAuthServiceResponseBodyAuthInfoModels {
	s.AuthCorpName = &v
	return s
}

func (s *QueryEduUnionAuthServiceResponseBodyAuthInfoModels) SetAuthTime(v string) *QueryEduUnionAuthServiceResponseBodyAuthInfoModels {
	s.AuthTime = &v
	return s
}

func (s *QueryEduUnionAuthServiceResponseBodyAuthInfoModels) SetResourceNames(v []*string) *QueryEduUnionAuthServiceResponseBodyAuthInfoModels {
	s.ResourceNames = v
	return s
}

type QueryEduUnionAuthServiceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEduUnionAuthServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEduUnionAuthServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEduUnionAuthServiceResponse) GoString() string {
	return s.String()
}

func (s *QueryEduUnionAuthServiceResponse) SetHeaders(v map[string]*string) *QueryEduUnionAuthServiceResponse {
	s.Headers = v
	return s
}

func (s *QueryEduUnionAuthServiceResponse) SetStatusCode(v int32) *QueryEduUnionAuthServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEduUnionAuthServiceResponse) SetBody(v *QueryEduUnionAuthServiceResponseBody) *QueryEduUnionAuthServiceResponse {
	s.Body = v
	return s
}

type QueryEmployeeTypeStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryEmployeeTypeStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryEmployeeTypeStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryEmployeeTypeStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryEmployeeTypeStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryEmployeeTypeStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryEmployeeTypeStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryEmployeeTypeStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryEmployeeTypeStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEmployeeTypeStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryEmployeeTypeStatisticalDataRequest) SetStatDate(v string) *QueryEmployeeTypeStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryEmployeeTypeStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryEmployeeTypeStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryEmployeeTypeStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEmployeeTypeStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEmployeeTypeStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryEmployeeTypeStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryEmployeeTypeStatisticalDataResponseBody) SetMetaList(v []*QueryEmployeeTypeStatisticalDataResponseBodyMetaList) *QueryEmployeeTypeStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryEmployeeTypeStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryEmployeeTypeStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryEmployeeTypeStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryEmployeeTypeStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryEmployeeTypeStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryEmployeeTypeStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryEmployeeTypeStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryEmployeeTypeStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryEmployeeTypeStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryEmployeeTypeStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryEmployeeTypeStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryEmployeeTypeStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryEmployeeTypeStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryEmployeeTypeStatisticalDataResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEmployeeTypeStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEmployeeTypeStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEmployeeTypeStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryEmployeeTypeStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryEmployeeTypeStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryEmployeeTypeStatisticalDataResponse) SetStatusCode(v int32) *QueryEmployeeTypeStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEmployeeTypeStatisticalDataResponse) SetBody(v *QueryEmployeeTypeStatisticalDataResponseBody) *QueryEmployeeTypeStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryGeneralDataServiceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGeneralDataServiceHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGeneralDataServiceHeaders) GoString() string {
	return s.String()
}

func (s *QueryGeneralDataServiceHeaders) SetCommonHeaders(v map[string]*string) *QueryGeneralDataServiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGeneralDataServiceHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGeneralDataServiceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGeneralDataServiceRequest struct {
	// example:
	//
	// 123
	DeptId *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20220803
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize    *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ReturnTotal *bool  `json:"returnTotal,omitempty" xml:"returnTotal,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// API-7fa754fd-f53e-46ee-9b77-898aa6eb590c
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20220801
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0234412313
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryGeneralDataServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGeneralDataServiceRequest) GoString() string {
	return s.String()
}

func (s *QueryGeneralDataServiceRequest) SetDeptId(v string) *QueryGeneralDataServiceRequest {
	s.DeptId = &v
	return s
}

func (s *QueryGeneralDataServiceRequest) SetEndDate(v string) *QueryGeneralDataServiceRequest {
	s.EndDate = &v
	return s
}

func (s *QueryGeneralDataServiceRequest) SetPageNumber(v int64) *QueryGeneralDataServiceRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryGeneralDataServiceRequest) SetPageSize(v int64) *QueryGeneralDataServiceRequest {
	s.PageSize = &v
	return s
}

func (s *QueryGeneralDataServiceRequest) SetReturnTotal(v bool) *QueryGeneralDataServiceRequest {
	s.ReturnTotal = &v
	return s
}

func (s *QueryGeneralDataServiceRequest) SetServiceId(v string) *QueryGeneralDataServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *QueryGeneralDataServiceRequest) SetStartDate(v string) *QueryGeneralDataServiceRequest {
	s.StartDate = &v
	return s
}

func (s *QueryGeneralDataServiceRequest) SetUserId(v string) *QueryGeneralDataServiceRequest {
	s.UserId = &v
	return s
}

type QueryGeneralDataServiceResponseBody struct {
	DataList []map[string]interface{}                       `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryGeneralDataServiceResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
	Total    *string                                        `json:"total,omitempty" xml:"total,omitempty"`
}

func (s QueryGeneralDataServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGeneralDataServiceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGeneralDataServiceResponseBody) SetDataList(v []map[string]interface{}) *QueryGeneralDataServiceResponseBody {
	s.DataList = v
	return s
}

func (s *QueryGeneralDataServiceResponseBody) SetMetaList(v []*QueryGeneralDataServiceResponseBodyMetaList) *QueryGeneralDataServiceResponseBody {
	s.MetaList = v
	return s
}

func (s *QueryGeneralDataServiceResponseBody) SetTotal(v string) *QueryGeneralDataServiceResponseBody {
	s.Total = &v
	return s
}

type QueryGeneralDataServiceResponseBodyMetaList struct {
	// This parameter is required.
	FieldDesc *string `json:"fieldDesc,omitempty" xml:"fieldDesc,omitempty"`
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// This parameter is required.
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
}

func (s QueryGeneralDataServiceResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryGeneralDataServiceResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryGeneralDataServiceResponseBodyMetaList) SetFieldDesc(v string) *QueryGeneralDataServiceResponseBodyMetaList {
	s.FieldDesc = &v
	return s
}

func (s *QueryGeneralDataServiceResponseBodyMetaList) SetFieldId(v string) *QueryGeneralDataServiceResponseBodyMetaList {
	s.FieldId = &v
	return s
}

func (s *QueryGeneralDataServiceResponseBodyMetaList) SetFieldName(v string) *QueryGeneralDataServiceResponseBodyMetaList {
	s.FieldName = &v
	return s
}

func (s *QueryGeneralDataServiceResponseBodyMetaList) SetFieldType(v string) *QueryGeneralDataServiceResponseBodyMetaList {
	s.FieldType = &v
	return s
}

type QueryGeneralDataServiceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGeneralDataServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGeneralDataServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGeneralDataServiceResponse) GoString() string {
	return s.String()
}

func (s *QueryGeneralDataServiceResponse) SetHeaders(v map[string]*string) *QueryGeneralDataServiceResponse {
	s.Headers = v
	return s
}

func (s *QueryGeneralDataServiceResponse) SetStatusCode(v int32) *QueryGeneralDataServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGeneralDataServiceResponse) SetBody(v *QueryGeneralDataServiceResponseBody) *QueryGeneralDataServiceResponse {
	s.Body = v
	return s
}

type QueryGeneralDataServiceBatchHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGeneralDataServiceBatchHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGeneralDataServiceBatchHeaders) GoString() string {
	return s.String()
}

func (s *QueryGeneralDataServiceBatchHeaders) SetCommonHeaders(v map[string]*string) *QueryGeneralDataServiceBatchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGeneralDataServiceBatchHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGeneralDataServiceBatchHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGeneralDataServiceBatchRequest struct {
	DeptIds []*string `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	// This parameter is required.
	EndDate *string                                       `json:"endDate,omitempty" xml:"endDate,omitempty"`
	Filters []*QueryGeneralDataServiceBatchRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Repeated"`
	// This parameter is required.
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	PageSize    *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ReturnTotal *bool  `json:"returnTotal,omitempty" xml:"returnTotal,omitempty"`
	// This parameter is required.
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// This parameter is required.
	StartDate *string   `json:"startDate,omitempty" xml:"startDate,omitempty"`
	UserId    *string   `json:"userId,omitempty" xml:"userId,omitempty"`
	UserIds   []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s QueryGeneralDataServiceBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGeneralDataServiceBatchRequest) GoString() string {
	return s.String()
}

func (s *QueryGeneralDataServiceBatchRequest) SetDeptIds(v []*string) *QueryGeneralDataServiceBatchRequest {
	s.DeptIds = v
	return s
}

func (s *QueryGeneralDataServiceBatchRequest) SetEndDate(v string) *QueryGeneralDataServiceBatchRequest {
	s.EndDate = &v
	return s
}

func (s *QueryGeneralDataServiceBatchRequest) SetFilters(v []*QueryGeneralDataServiceBatchRequestFilters) *QueryGeneralDataServiceBatchRequest {
	s.Filters = v
	return s
}

func (s *QueryGeneralDataServiceBatchRequest) SetPageNumber(v int64) *QueryGeneralDataServiceBatchRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryGeneralDataServiceBatchRequest) SetPageSize(v int64) *QueryGeneralDataServiceBatchRequest {
	s.PageSize = &v
	return s
}

func (s *QueryGeneralDataServiceBatchRequest) SetReturnTotal(v bool) *QueryGeneralDataServiceBatchRequest {
	s.ReturnTotal = &v
	return s
}

func (s *QueryGeneralDataServiceBatchRequest) SetServiceId(v string) *QueryGeneralDataServiceBatchRequest {
	s.ServiceId = &v
	return s
}

func (s *QueryGeneralDataServiceBatchRequest) SetStartDate(v string) *QueryGeneralDataServiceBatchRequest {
	s.StartDate = &v
	return s
}

func (s *QueryGeneralDataServiceBatchRequest) SetUserId(v string) *QueryGeneralDataServiceBatchRequest {
	s.UserId = &v
	return s
}

func (s *QueryGeneralDataServiceBatchRequest) SetUserIds(v []*string) *QueryGeneralDataServiceBatchRequest {
	s.UserIds = v
	return s
}

type QueryGeneralDataServiceBatchRequestFilters struct {
	FieldId  *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s QueryGeneralDataServiceBatchRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s QueryGeneralDataServiceBatchRequestFilters) GoString() string {
	return s.String()
}

func (s *QueryGeneralDataServiceBatchRequestFilters) SetFieldId(v string) *QueryGeneralDataServiceBatchRequestFilters {
	s.FieldId = &v
	return s
}

func (s *QueryGeneralDataServiceBatchRequestFilters) SetOperator(v string) *QueryGeneralDataServiceBatchRequestFilters {
	s.Operator = &v
	return s
}

func (s *QueryGeneralDataServiceBatchRequestFilters) SetValue(v string) *QueryGeneralDataServiceBatchRequestFilters {
	s.Value = &v
	return s
}

type QueryGeneralDataServiceBatchResponseBody struct {
	DataList []map[string]interface{}                            `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryGeneralDataServiceBatchResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
	Total    *int64                                              `json:"total,omitempty" xml:"total,omitempty"`
}

func (s QueryGeneralDataServiceBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGeneralDataServiceBatchResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGeneralDataServiceBatchResponseBody) SetDataList(v []map[string]interface{}) *QueryGeneralDataServiceBatchResponseBody {
	s.DataList = v
	return s
}

func (s *QueryGeneralDataServiceBatchResponseBody) SetMetaList(v []*QueryGeneralDataServiceBatchResponseBodyMetaList) *QueryGeneralDataServiceBatchResponseBody {
	s.MetaList = v
	return s
}

func (s *QueryGeneralDataServiceBatchResponseBody) SetTotal(v int64) *QueryGeneralDataServiceBatchResponseBody {
	s.Total = &v
	return s
}

type QueryGeneralDataServiceBatchResponseBodyMetaList struct {
	// This parameter is required.
	FieldDesc *string `json:"fieldDesc,omitempty" xml:"fieldDesc,omitempty"`
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// This parameter is required.
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
}

func (s QueryGeneralDataServiceBatchResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryGeneralDataServiceBatchResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryGeneralDataServiceBatchResponseBodyMetaList) SetFieldDesc(v string) *QueryGeneralDataServiceBatchResponseBodyMetaList {
	s.FieldDesc = &v
	return s
}

func (s *QueryGeneralDataServiceBatchResponseBodyMetaList) SetFieldId(v string) *QueryGeneralDataServiceBatchResponseBodyMetaList {
	s.FieldId = &v
	return s
}

func (s *QueryGeneralDataServiceBatchResponseBodyMetaList) SetFieldName(v string) *QueryGeneralDataServiceBatchResponseBodyMetaList {
	s.FieldName = &v
	return s
}

func (s *QueryGeneralDataServiceBatchResponseBodyMetaList) SetFieldType(v string) *QueryGeneralDataServiceBatchResponseBodyMetaList {
	s.FieldType = &v
	return s
}

type QueryGeneralDataServiceBatchResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGeneralDataServiceBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGeneralDataServiceBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGeneralDataServiceBatchResponse) GoString() string {
	return s.String()
}

func (s *QueryGeneralDataServiceBatchResponse) SetHeaders(v map[string]*string) *QueryGeneralDataServiceBatchResponse {
	s.Headers = v
	return s
}

func (s *QueryGeneralDataServiceBatchResponse) SetStatusCode(v int32) *QueryGeneralDataServiceBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGeneralDataServiceBatchResponse) SetBody(v *QueryGeneralDataServiceBatchResponseBody) *QueryGeneralDataServiceBatchResponse {
	s.Body = v
	return s
}

type QueryGeneralDataUpdateDateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGeneralDataUpdateDateHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGeneralDataUpdateDateHeaders) GoString() string {
	return s.String()
}

func (s *QueryGeneralDataUpdateDateHeaders) SetCommonHeaders(v map[string]*string) *QueryGeneralDataUpdateDateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGeneralDataUpdateDateHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGeneralDataUpdateDateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGeneralDataUpdateDateRequest struct {
	// This parameter is required.
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
}

func (s QueryGeneralDataUpdateDateRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGeneralDataUpdateDateRequest) GoString() string {
	return s.String()
}

func (s *QueryGeneralDataUpdateDateRequest) SetServiceId(v string) *QueryGeneralDataUpdateDateRequest {
	s.ServiceId = &v
	return s
}

type QueryGeneralDataUpdateDateResponseBody struct {
	Success    *bool   `json:"success,omitempty" xml:"success,omitempty"`
	UpdateDate *string `json:"updateDate,omitempty" xml:"updateDate,omitempty"`
}

func (s QueryGeneralDataUpdateDateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGeneralDataUpdateDateResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGeneralDataUpdateDateResponseBody) SetSuccess(v bool) *QueryGeneralDataUpdateDateResponseBody {
	s.Success = &v
	return s
}

func (s *QueryGeneralDataUpdateDateResponseBody) SetUpdateDate(v string) *QueryGeneralDataUpdateDateResponseBody {
	s.UpdateDate = &v
	return s
}

type QueryGeneralDataUpdateDateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGeneralDataUpdateDateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGeneralDataUpdateDateResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGeneralDataUpdateDateResponse) GoString() string {
	return s.String()
}

func (s *QueryGeneralDataUpdateDateResponse) SetHeaders(v map[string]*string) *QueryGeneralDataUpdateDateResponse {
	s.Headers = v
	return s
}

func (s *QueryGeneralDataUpdateDateResponse) SetStatusCode(v int32) *QueryGeneralDataUpdateDateResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGeneralDataUpdateDateResponse) SetBody(v *QueryGeneralDataUpdateDateResponseBody) *QueryGeneralDataUpdateDateResponse {
	s.Body = v
	return s
}

type QueryGroupLiveStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGroupLiveStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupLiveStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryGroupLiveStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryGroupLiveStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGroupLiveStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGroupLiveStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGroupLiveStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryGroupLiveStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupLiveStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryGroupLiveStatisticalDataRequest) SetStatDate(v string) *QueryGroupLiveStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryGroupLiveStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                             `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryGroupLiveStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryGroupLiveStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupLiveStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupLiveStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryGroupLiveStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryGroupLiveStatisticalDataResponseBody) SetMetaList(v []*QueryGroupLiveStatisticalDataResponseBodyMetaList) *QueryGroupLiveStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryGroupLiveStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryGroupLiveStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupLiveStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryGroupLiveStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryGroupLiveStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryGroupLiveStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryGroupLiveStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryGroupLiveStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryGroupLiveStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryGroupLiveStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryGroupLiveStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryGroupLiveStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryGroupLiveStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryGroupLiveStatisticalDataResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGroupLiveStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGroupLiveStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupLiveStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupLiveStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryGroupLiveStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupLiveStatisticalDataResponse) SetStatusCode(v int32) *QueryGroupLiveStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGroupLiveStatisticalDataResponse) SetBody(v *QueryGroupLiveStatisticalDataResponseBody) *QueryGroupLiveStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryGroupMessageStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGroupMessageStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMessageStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryGroupMessageStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryGroupMessageStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGroupMessageStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGroupMessageStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGroupMessageStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryGroupMessageStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMessageStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryGroupMessageStatisticalDataRequest) SetStatDate(v string) *QueryGroupMessageStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryGroupMessageStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryGroupMessageStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryGroupMessageStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMessageStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupMessageStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryGroupMessageStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryGroupMessageStatisticalDataResponseBody) SetMetaList(v []*QueryGroupMessageStatisticalDataResponseBodyMetaList) *QueryGroupMessageStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryGroupMessageStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryGroupMessageStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMessageStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryGroupMessageStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryGroupMessageStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryGroupMessageStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryGroupMessageStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryGroupMessageStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryGroupMessageStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryGroupMessageStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryGroupMessageStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryGroupMessageStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryGroupMessageStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryGroupMessageStatisticalDataResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGroupMessageStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGroupMessageStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMessageStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupMessageStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryGroupMessageStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupMessageStatisticalDataResponse) SetStatusCode(v int32) *QueryGroupMessageStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGroupMessageStatisticalDataResponse) SetBody(v *QueryGroupMessageStatisticalDataResponseBody) *QueryGroupMessageStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryHealthStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryHealthStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryHealthStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryHealthStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryHealthStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryHealthStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryHealthStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryHealthStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryHealthStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHealthStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryHealthStatisticalDataRequest) SetStatDate(v string) *QueryHealthStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryHealthStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                          `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryHealthStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryHealthStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHealthStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHealthStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryHealthStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryHealthStatisticalDataResponseBody) SetMetaList(v []*QueryHealthStatisticalDataResponseBodyMetaList) *QueryHealthStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryHealthStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryHealthStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryHealthStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryHealthStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryHealthStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryHealthStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryHealthStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryHealthStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryHealthStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryHealthStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryHealthStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryHealthStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryHealthStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryHealthStatisticalDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryHealthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryHealthStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHealthStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryHealthStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryHealthStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryHealthStatisticalDataResponse) SetStatusCode(v int32) *QueryHealthStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHealthStatisticalDataResponse) SetBody(v *QueryHealthStatisticalDataResponseBody) *QueryHealthStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryMailStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMailStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMailStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryMailStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryMailStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMailStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMailStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMailStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryMailStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMailStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryMailStatisticalDataRequest) SetStatDate(v string) *QueryMailStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryMailStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                        `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryMailStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryMailStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMailStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMailStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryMailStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryMailStatisticalDataResponseBody) SetMetaList(v []*QueryMailStatisticalDataResponseBodyMetaList) *QueryMailStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryMailStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryMailStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryMailStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryMailStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryMailStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryMailStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryMailStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryMailStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryMailStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryMailStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryMailStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryMailStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryMailStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryMailStatisticalDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMailStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMailStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMailStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryMailStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryMailStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryMailStatisticalDataResponse) SetStatusCode(v int32) *QueryMailStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMailStatisticalDataResponse) SetBody(v *QueryMailStatisticalDataResponseBody) *QueryMailStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryOfficialDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOfficialDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryOfficialDataHeaders) SetCommonHeaders(v map[string]*string) *QueryOfficialDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOfficialDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOfficialDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOfficialDataRequest struct {
	Param  *string `json:"param,omitempty" xml:"param,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryOfficialDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialDataRequest) GoString() string {
	return s.String()
}

func (s *QueryOfficialDataRequest) SetParam(v string) *QueryOfficialDataRequest {
	s.Param = &v
	return s
}

func (s *QueryOfficialDataRequest) SetUserId(v string) *QueryOfficialDataRequest {
	s.UserId = &v
	return s
}

type QueryOfficialDataResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryOfficialDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOfficialDataResponseBody) SetResult(v string) *QueryOfficialDataResponseBody {
	s.Result = &v
	return s
}

func (s *QueryOfficialDataResponseBody) SetSuccess(v bool) *QueryOfficialDataResponseBody {
	s.Success = &v
	return s
}

type QueryOfficialDataResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOfficialDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOfficialDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialDataResponse) GoString() string {
	return s.String()
}

func (s *QueryOfficialDataResponse) SetHeaders(v map[string]*string) *QueryOfficialDataResponse {
	s.Headers = v
	return s
}

func (s *QueryOfficialDataResponse) SetStatusCode(v int32) *QueryOfficialDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOfficialDataResponse) SetBody(v *QueryOfficialDataResponseBody) *QueryOfficialDataResponse {
	s.Body = v
	return s
}

type QueryOfficialDatasetFieldsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOfficialDatasetFieldsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialDatasetFieldsHeaders) GoString() string {
	return s.String()
}

func (s *QueryOfficialDatasetFieldsHeaders) SetCommonHeaders(v map[string]*string) *QueryOfficialDatasetFieldsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOfficialDatasetFieldsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOfficialDatasetFieldsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOfficialDatasetFieldsRequest struct {
	// example:
	//
	// ding3xxx__-PROC-42FF6625-9692-4003-B13C-307CAACEC354
	DsId *string `json:"dsId,omitempty" xml:"dsId,omitempty"`
	// example:
	//
	// 12345
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryOfficialDatasetFieldsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialDatasetFieldsRequest) GoString() string {
	return s.String()
}

func (s *QueryOfficialDatasetFieldsRequest) SetDsId(v string) *QueryOfficialDatasetFieldsRequest {
	s.DsId = &v
	return s
}

func (s *QueryOfficialDatasetFieldsRequest) SetUserId(v string) *QueryOfficialDatasetFieldsRequest {
	s.UserId = &v
	return s
}

type QueryOfficialDatasetFieldsResponseBody struct {
	Result  *QueryOfficialDatasetFieldsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryOfficialDatasetFieldsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialDatasetFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOfficialDatasetFieldsResponseBody) SetResult(v *QueryOfficialDatasetFieldsResponseBodyResult) *QueryOfficialDatasetFieldsResponseBody {
	s.Result = v
	return s
}

func (s *QueryOfficialDatasetFieldsResponseBody) SetSuccess(v bool) *QueryOfficialDatasetFieldsResponseBody {
	s.Success = &v
	return s
}

type QueryOfficialDatasetFieldsResponseBodyResult struct {
	DisplayName *string                                               `json:"displayName,omitempty" xml:"displayName,omitempty"`
	DsId        *string                                               `json:"dsId,omitempty" xml:"dsId,omitempty"`
	Fields      []*QueryOfficialDatasetFieldsResponseBodyResultFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
}

func (s QueryOfficialDatasetFieldsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialDatasetFieldsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryOfficialDatasetFieldsResponseBodyResult) SetDisplayName(v string) *QueryOfficialDatasetFieldsResponseBodyResult {
	s.DisplayName = &v
	return s
}

func (s *QueryOfficialDatasetFieldsResponseBodyResult) SetDsId(v string) *QueryOfficialDatasetFieldsResponseBodyResult {
	s.DsId = &v
	return s
}

func (s *QueryOfficialDatasetFieldsResponseBodyResult) SetFields(v []*QueryOfficialDatasetFieldsResponseBodyResultFields) *QueryOfficialDatasetFieldsResponseBodyResult {
	s.Fields = v
	return s
}

type QueryOfficialDatasetFieldsResponseBodyResultFields struct {
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	FieldId     *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	FieldType   *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
}

func (s QueryOfficialDatasetFieldsResponseBodyResultFields) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialDatasetFieldsResponseBodyResultFields) GoString() string {
	return s.String()
}

func (s *QueryOfficialDatasetFieldsResponseBodyResultFields) SetDisplayName(v string) *QueryOfficialDatasetFieldsResponseBodyResultFields {
	s.DisplayName = &v
	return s
}

func (s *QueryOfficialDatasetFieldsResponseBodyResultFields) SetFieldId(v string) *QueryOfficialDatasetFieldsResponseBodyResultFields {
	s.FieldId = &v
	return s
}

func (s *QueryOfficialDatasetFieldsResponseBodyResultFields) SetFieldType(v string) *QueryOfficialDatasetFieldsResponseBodyResultFields {
	s.FieldType = &v
	return s
}

type QueryOfficialDatasetFieldsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOfficialDatasetFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOfficialDatasetFieldsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialDatasetFieldsResponse) GoString() string {
	return s.String()
}

func (s *QueryOfficialDatasetFieldsResponse) SetHeaders(v map[string]*string) *QueryOfficialDatasetFieldsResponse {
	s.Headers = v
	return s
}

func (s *QueryOfficialDatasetFieldsResponse) SetStatusCode(v int32) *QueryOfficialDatasetFieldsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOfficialDatasetFieldsResponse) SetBody(v *QueryOfficialDatasetFieldsResponseBody) *QueryOfficialDatasetFieldsResponse {
	s.Body = v
	return s
}

type QueryOfficialDatasetListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOfficialDatasetListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialDatasetListHeaders) GoString() string {
	return s.String()
}

func (s *QueryOfficialDatasetListHeaders) SetCommonHeaders(v map[string]*string) *QueryOfficialDatasetListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOfficialDatasetListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOfficialDatasetListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOfficialDatasetListRequest struct {
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryOfficialDatasetListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialDatasetListRequest) GoString() string {
	return s.String()
}

func (s *QueryOfficialDatasetListRequest) SetKeyword(v string) *QueryOfficialDatasetListRequest {
	s.Keyword = &v
	return s
}

func (s *QueryOfficialDatasetListRequest) SetPageNumber(v int32) *QueryOfficialDatasetListRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryOfficialDatasetListRequest) SetPageSize(v int32) *QueryOfficialDatasetListRequest {
	s.PageSize = &v
	return s
}

type QueryOfficialDatasetListResponseBody struct {
	Result  *QueryOfficialDatasetListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                       `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryOfficialDatasetListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialDatasetListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOfficialDatasetListResponseBody) SetResult(v *QueryOfficialDatasetListResponseBodyResult) *QueryOfficialDatasetListResponseBody {
	s.Result = v
	return s
}

func (s *QueryOfficialDatasetListResponseBody) SetSuccess(v bool) *QueryOfficialDatasetListResponseBody {
	s.Success = &v
	return s
}

type QueryOfficialDatasetListResponseBodyResult struct {
	Rows       []*QueryOfficialDatasetListResponseBodyResultRows `json:"rows,omitempty" xml:"rows,omitempty" type:"Repeated"`
	TotalCount *int64                                            `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryOfficialDatasetListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialDatasetListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryOfficialDatasetListResponseBodyResult) SetRows(v []*QueryOfficialDatasetListResponseBodyResultRows) *QueryOfficialDatasetListResponseBodyResult {
	s.Rows = v
	return s
}

func (s *QueryOfficialDatasetListResponseBodyResult) SetTotalCount(v int64) *QueryOfficialDatasetListResponseBodyResult {
	s.TotalCount = &v
	return s
}

type QueryOfficialDatasetListResponseBodyResultRows struct {
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	DsId        *string `json:"dsId,omitempty" xml:"dsId,omitempty"`
}

func (s QueryOfficialDatasetListResponseBodyResultRows) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialDatasetListResponseBodyResultRows) GoString() string {
	return s.String()
}

func (s *QueryOfficialDatasetListResponseBodyResultRows) SetDisplayName(v string) *QueryOfficialDatasetListResponseBodyResultRows {
	s.DisplayName = &v
	return s
}

func (s *QueryOfficialDatasetListResponseBodyResultRows) SetDsId(v string) *QueryOfficialDatasetListResponseBodyResultRows {
	s.DsId = &v
	return s
}

type QueryOfficialDatasetListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOfficialDatasetListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOfficialDatasetListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialDatasetListResponse) GoString() string {
	return s.String()
}

func (s *QueryOfficialDatasetListResponse) SetHeaders(v map[string]*string) *QueryOfficialDatasetListResponse {
	s.Headers = v
	return s
}

func (s *QueryOfficialDatasetListResponse) SetStatusCode(v int32) *QueryOfficialDatasetListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOfficialDatasetListResponse) SetBody(v *QueryOfficialDatasetListResponseBody) *QueryOfficialDatasetListResponse {
	s.Body = v
	return s
}

type QueryOfficialFormDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOfficialFormDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialFormDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryOfficialFormDataHeaders) SetCommonHeaders(v map[string]*string) *QueryOfficialFormDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOfficialFormDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOfficialFormDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOfficialFormDataRequest struct {
	Param  *string `json:"param,omitempty" xml:"param,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryOfficialFormDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialFormDataRequest) GoString() string {
	return s.String()
}

func (s *QueryOfficialFormDataRequest) SetParam(v string) *QueryOfficialFormDataRequest {
	s.Param = &v
	return s
}

func (s *QueryOfficialFormDataRequest) SetUserId(v string) *QueryOfficialFormDataRequest {
	s.UserId = &v
	return s
}

type QueryOfficialFormDataResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryOfficialFormDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialFormDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOfficialFormDataResponseBody) SetResult(v string) *QueryOfficialFormDataResponseBody {
	s.Result = &v
	return s
}

func (s *QueryOfficialFormDataResponseBody) SetSuccess(v bool) *QueryOfficialFormDataResponseBody {
	s.Success = &v
	return s
}

type QueryOfficialFormDataResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOfficialFormDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOfficialFormDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialFormDataResponse) GoString() string {
	return s.String()
}

func (s *QueryOfficialFormDataResponse) SetHeaders(v map[string]*string) *QueryOfficialFormDataResponse {
	s.Headers = v
	return s
}

func (s *QueryOfficialFormDataResponse) SetStatusCode(v int32) *QueryOfficialFormDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOfficialFormDataResponse) SetBody(v *QueryOfficialFormDataResponseBody) *QueryOfficialFormDataResponse {
	s.Body = v
	return s
}

type QueryOfficialFormDataDirectHoloHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOfficialFormDataDirectHoloHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialFormDataDirectHoloHeaders) GoString() string {
	return s.String()
}

func (s *QueryOfficialFormDataDirectHoloHeaders) SetCommonHeaders(v map[string]*string) *QueryOfficialFormDataDirectHoloHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOfficialFormDataDirectHoloHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOfficialFormDataDirectHoloHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOfficialFormDataDirectHoloRequest struct {
	Param  *string `json:"param,omitempty" xml:"param,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryOfficialFormDataDirectHoloRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialFormDataDirectHoloRequest) GoString() string {
	return s.String()
}

func (s *QueryOfficialFormDataDirectHoloRequest) SetParam(v string) *QueryOfficialFormDataDirectHoloRequest {
	s.Param = &v
	return s
}

func (s *QueryOfficialFormDataDirectHoloRequest) SetUserId(v string) *QueryOfficialFormDataDirectHoloRequest {
	s.UserId = &v
	return s
}

type QueryOfficialFormDataDirectHoloResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryOfficialFormDataDirectHoloResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialFormDataDirectHoloResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOfficialFormDataDirectHoloResponseBody) SetResult(v string) *QueryOfficialFormDataDirectHoloResponseBody {
	s.Result = &v
	return s
}

func (s *QueryOfficialFormDataDirectHoloResponseBody) SetSuccess(v bool) *QueryOfficialFormDataDirectHoloResponseBody {
	s.Success = &v
	return s
}

type QueryOfficialFormDataDirectHoloResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOfficialFormDataDirectHoloResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOfficialFormDataDirectHoloResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialFormDataDirectHoloResponse) GoString() string {
	return s.String()
}

func (s *QueryOfficialFormDataDirectHoloResponse) SetHeaders(v map[string]*string) *QueryOfficialFormDataDirectHoloResponse {
	s.Headers = v
	return s
}

func (s *QueryOfficialFormDataDirectHoloResponse) SetStatusCode(v int32) *QueryOfficialFormDataDirectHoloResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOfficialFormDataDirectHoloResponse) SetBody(v *QueryOfficialFormDataDirectHoloResponseBody) *QueryOfficialFormDataDirectHoloResponse {
	s.Body = v
	return s
}

type QueryOnlineUserStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOnlineUserStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOnlineUserStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryOnlineUserStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryOnlineUserStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOnlineUserStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOnlineUserStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOnlineUserStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryOnlineUserStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOnlineUserStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryOnlineUserStatisticalDataRequest) SetStatDate(v string) *QueryOnlineUserStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryOnlineUserStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                              `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryOnlineUserStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryOnlineUserStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOnlineUserStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOnlineUserStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryOnlineUserStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryOnlineUserStatisticalDataResponseBody) SetMetaList(v []*QueryOnlineUserStatisticalDataResponseBodyMetaList) *QueryOnlineUserStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryOnlineUserStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryOnlineUserStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryOnlineUserStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryOnlineUserStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryOnlineUserStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryOnlineUserStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryOnlineUserStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryOnlineUserStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryOnlineUserStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryOnlineUserStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryOnlineUserStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryOnlineUserStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryOnlineUserStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryOnlineUserStatisticalDataResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOnlineUserStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOnlineUserStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOnlineUserStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryOnlineUserStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryOnlineUserStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryOnlineUserStatisticalDataResponse) SetStatusCode(v int32) *QueryOnlineUserStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOnlineUserStatisticalDataResponse) SetBody(v *QueryOnlineUserStatisticalDataResponseBody) *QueryOnlineUserStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryRedEnvelopeReciveStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryRedEnvelopeReciveStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryRedEnvelopeReciveStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryRedEnvelopeReciveStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryRedEnvelopeReciveStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryRedEnvelopeReciveStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryRedEnvelopeReciveStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryRedEnvelopeReciveStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryRedEnvelopeReciveStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRedEnvelopeReciveStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryRedEnvelopeReciveStatisticalDataRequest) SetStatDate(v string) *QueryRedEnvelopeReciveStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryRedEnvelopeReciveStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                     `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryRedEnvelopeReciveStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRedEnvelopeReciveStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRedEnvelopeReciveStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryRedEnvelopeReciveStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryRedEnvelopeReciveStatisticalDataResponseBody) SetMetaList(v []*QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList) *QueryRedEnvelopeReciveStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryRedEnvelopeReciveStatisticalDataResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRedEnvelopeReciveStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRedEnvelopeReciveStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRedEnvelopeReciveStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryRedEnvelopeReciveStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryRedEnvelopeReciveStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryRedEnvelopeReciveStatisticalDataResponse) SetStatusCode(v int32) *QueryRedEnvelopeReciveStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRedEnvelopeReciveStatisticalDataResponse) SetBody(v *QueryRedEnvelopeReciveStatisticalDataResponseBody) *QueryRedEnvelopeReciveStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryRedEnvelopeSendStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryRedEnvelopeSendStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryRedEnvelopeSendStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryRedEnvelopeSendStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryRedEnvelopeSendStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryRedEnvelopeSendStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryRedEnvelopeSendStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryRedEnvelopeSendStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryRedEnvelopeSendStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRedEnvelopeSendStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryRedEnvelopeSendStatisticalDataRequest) SetStatDate(v string) *QueryRedEnvelopeSendStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryRedEnvelopeSendStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                   `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryRedEnvelopeSendStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRedEnvelopeSendStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRedEnvelopeSendStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryRedEnvelopeSendStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryRedEnvelopeSendStatisticalDataResponseBody) SetMetaList(v []*QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList) *QueryRedEnvelopeSendStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryRedEnvelopeSendStatisticalDataResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRedEnvelopeSendStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRedEnvelopeSendStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRedEnvelopeSendStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryRedEnvelopeSendStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryRedEnvelopeSendStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryRedEnvelopeSendStatisticalDataResponse) SetStatusCode(v int32) *QueryRedEnvelopeSendStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRedEnvelopeSendStatisticalDataResponse) SetBody(v *QueryRedEnvelopeSendStatisticalDataResponseBody) *QueryRedEnvelopeSendStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryReportStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryReportStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryReportStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryReportStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryReportStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryReportStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryReportStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryReportStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryReportStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryReportStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryReportStatisticalDataRequest) SetStatDate(v string) *QueryReportStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryReportStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                          `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryReportStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryReportStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryReportStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryReportStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryReportStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryReportStatisticalDataResponseBody) SetMetaList(v []*QueryReportStatisticalDataResponseBodyMetaList) *QueryReportStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryReportStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryReportStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryReportStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryReportStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryReportStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryReportStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryReportStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryReportStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryReportStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryReportStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryReportStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryReportStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryReportStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryReportStatisticalDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryReportStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryReportStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryReportStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryReportStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryReportStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryReportStatisticalDataResponse) SetStatusCode(v int32) *QueryReportStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryReportStatisticalDataResponse) SetBody(v *QueryReportStatisticalDataResponseBody) *QueryReportStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryScreenHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryScreenHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryScreenHeaders) GoString() string {
	return s.String()
}

func (s *QueryScreenHeaders) SetCommonHeaders(v map[string]*string) *QueryScreenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryScreenHeaders) SetXAcsDingtalkAccessToken(v string) *QueryScreenHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryScreenRequest struct {
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s QueryScreenRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryScreenRequest) GoString() string {
	return s.String()
}

func (s *QueryScreenRequest) SetOperatorId(v string) *QueryScreenRequest {
	s.OperatorId = &v
	return s
}

type QueryScreenResponseBody struct {
	Result  []*QueryScreenResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success *bool                            `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryScreenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryScreenResponseBody) GoString() string {
	return s.String()
}

func (s *QueryScreenResponseBody) SetResult(v []*QueryScreenResponseBodyResult) *QueryScreenResponseBody {
	s.Result = v
	return s
}

func (s *QueryScreenResponseBody) SetSuccess(v bool) *QueryScreenResponseBody {
	s.Success = &v
	return s
}

type QueryScreenResponseBodyResult struct {
	OperatePermission *string `json:"operatePermission,omitempty" xml:"operatePermission,omitempty"`
	ScreenId          *int64  `json:"screenId,omitempty" xml:"screenId,omitempty"`
	ScreenName        *string `json:"screenName,omitempty" xml:"screenName,omitempty"`
	State             *string `json:"state,omitempty" xml:"state,omitempty"`
	ThumbUrl          *string `json:"thumbUrl,omitempty" xml:"thumbUrl,omitempty"`
}

func (s QueryScreenResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryScreenResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryScreenResponseBodyResult) SetOperatePermission(v string) *QueryScreenResponseBodyResult {
	s.OperatePermission = &v
	return s
}

func (s *QueryScreenResponseBodyResult) SetScreenId(v int64) *QueryScreenResponseBodyResult {
	s.ScreenId = &v
	return s
}

func (s *QueryScreenResponseBodyResult) SetScreenName(v string) *QueryScreenResponseBodyResult {
	s.ScreenName = &v
	return s
}

func (s *QueryScreenResponseBodyResult) SetState(v string) *QueryScreenResponseBodyResult {
	s.State = &v
	return s
}

func (s *QueryScreenResponseBodyResult) SetThumbUrl(v string) *QueryScreenResponseBodyResult {
	s.ThumbUrl = &v
	return s
}

type QueryScreenResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryScreenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryScreenResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryScreenResponse) GoString() string {
	return s.String()
}

func (s *QueryScreenResponse) SetHeaders(v map[string]*string) *QueryScreenResponse {
	s.Headers = v
	return s
}

func (s *QueryScreenResponse) SetStatusCode(v int32) *QueryScreenResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryScreenResponse) SetBody(v *QueryScreenResponseBody) *QueryScreenResponse {
	s.Body = v
	return s
}

type QueryScreenTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryScreenTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryScreenTemplateHeaders) GoString() string {
	return s.String()
}

func (s *QueryScreenTemplateHeaders) SetCommonHeaders(v map[string]*string) *QueryScreenTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryScreenTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *QueryScreenTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryScreenTemplateRequest struct {
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	Sample     *bool   `json:"sample,omitempty" xml:"sample,omitempty"`
}

func (s QueryScreenTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryScreenTemplateRequest) GoString() string {
	return s.String()
}

func (s *QueryScreenTemplateRequest) SetOperatorId(v string) *QueryScreenTemplateRequest {
	s.OperatorId = &v
	return s
}

func (s *QueryScreenTemplateRequest) SetSample(v bool) *QueryScreenTemplateRequest {
	s.Sample = &v
	return s
}

type QueryScreenTemplateResponseBody struct {
	Result  []*QueryScreenTemplateResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success *bool                                    `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryScreenTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryScreenTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *QueryScreenTemplateResponseBody) SetResult(v []*QueryScreenTemplateResponseBodyResult) *QueryScreenTemplateResponseBody {
	s.Result = v
	return s
}

func (s *QueryScreenTemplateResponseBody) SetSuccess(v bool) *QueryScreenTemplateResponseBody {
	s.Success = &v
	return s
}

type QueryScreenTemplateResponseBodyResult struct {
	PreviewUrl   *string `json:"previewUrl,omitempty" xml:"previewUrl,omitempty"`
	ScreenSize   *string `json:"screenSize,omitempty" xml:"screenSize,omitempty"`
	TemplateId   *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	ThumbUrl     *string `json:"thumbUrl,omitempty" xml:"thumbUrl,omitempty"`
}

func (s QueryScreenTemplateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryScreenTemplateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryScreenTemplateResponseBodyResult) SetPreviewUrl(v string) *QueryScreenTemplateResponseBodyResult {
	s.PreviewUrl = &v
	return s
}

func (s *QueryScreenTemplateResponseBodyResult) SetScreenSize(v string) *QueryScreenTemplateResponseBodyResult {
	s.ScreenSize = &v
	return s
}

func (s *QueryScreenTemplateResponseBodyResult) SetTemplateId(v string) *QueryScreenTemplateResponseBodyResult {
	s.TemplateId = &v
	return s
}

func (s *QueryScreenTemplateResponseBodyResult) SetTemplateName(v string) *QueryScreenTemplateResponseBodyResult {
	s.TemplateName = &v
	return s
}

func (s *QueryScreenTemplateResponseBodyResult) SetThumbUrl(v string) *QueryScreenTemplateResponseBodyResult {
	s.ThumbUrl = &v
	return s
}

type QueryScreenTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryScreenTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryScreenTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryScreenTemplateResponse) GoString() string {
	return s.String()
}

func (s *QueryScreenTemplateResponse) SetHeaders(v map[string]*string) *QueryScreenTemplateResponse {
	s.Headers = v
	return s
}

func (s *QueryScreenTemplateResponse) SetStatusCode(v int32) *QueryScreenTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryScreenTemplateResponse) SetBody(v *QueryScreenTemplateResponseBody) *QueryScreenTemplateResponse {
	s.Body = v
	return s
}

type QuerySingleMessageStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QuerySingleMessageStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleMessageStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QuerySingleMessageStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QuerySingleMessageStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QuerySingleMessageStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QuerySingleMessageStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QuerySingleMessageStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QuerySingleMessageStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleMessageStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QuerySingleMessageStatisticalDataRequest) SetStatDate(v string) *QuerySingleMessageStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QuerySingleMessageStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                 `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QuerySingleMessageStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QuerySingleMessageStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleMessageStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySingleMessageStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QuerySingleMessageStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QuerySingleMessageStatisticalDataResponseBody) SetMetaList(v []*QuerySingleMessageStatisticalDataResponseBodyMetaList) *QuerySingleMessageStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QuerySingleMessageStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QuerySingleMessageStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleMessageStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QuerySingleMessageStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QuerySingleMessageStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QuerySingleMessageStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QuerySingleMessageStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QuerySingleMessageStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QuerySingleMessageStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QuerySingleMessageStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QuerySingleMessageStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QuerySingleMessageStatisticalDataResponseBodyMetaList) SetUnit(v string) *QuerySingleMessageStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QuerySingleMessageStatisticalDataResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySingleMessageStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySingleMessageStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleMessageStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QuerySingleMessageStatisticalDataResponse) SetHeaders(v map[string]*string) *QuerySingleMessageStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QuerySingleMessageStatisticalDataResponse) SetStatusCode(v int32) *QuerySingleMessageStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySingleMessageStatisticalDataResponse) SetBody(v *QuerySingleMessageStatisticalDataResponseBody) *QuerySingleMessageStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryTelMeetingStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryTelMeetingStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryTelMeetingStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryTelMeetingStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryTelMeetingStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryTelMeetingStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryTelMeetingStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryTelMeetingStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryTelMeetingStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTelMeetingStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryTelMeetingStatisticalDataRequest) SetStatDate(v string) *QueryTelMeetingStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryTelMeetingStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                              `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryTelMeetingStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryTelMeetingStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTelMeetingStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTelMeetingStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryTelMeetingStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryTelMeetingStatisticalDataResponseBody) SetMetaList(v []*QueryTelMeetingStatisticalDataResponseBodyMetaList) *QueryTelMeetingStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryTelMeetingStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryTelMeetingStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryTelMeetingStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryTelMeetingStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryTelMeetingStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryTelMeetingStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryTelMeetingStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryTelMeetingStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryTelMeetingStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryTelMeetingStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryTelMeetingStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryTelMeetingStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryTelMeetingStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryTelMeetingStatisticalDataResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTelMeetingStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTelMeetingStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTelMeetingStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryTelMeetingStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryTelMeetingStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryTelMeetingStatisticalDataResponse) SetStatusCode(v int32) *QueryTelMeetingStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTelMeetingStatisticalDataResponse) SetBody(v *QueryTelMeetingStatisticalDataResponseBody) *QueryTelMeetingStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryTodoStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryTodoStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryTodoStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryTodoStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryTodoStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryTodoStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryTodoStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryTodoStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryTodoStatisticalDataRequest) SetStatDate(v string) *QueryTodoStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryTodoStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                        `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryTodoStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryTodoStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTodoStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryTodoStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryTodoStatisticalDataResponseBody) SetMetaList(v []*QueryTodoStatisticalDataResponseBodyMetaList) *QueryTodoStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryTodoStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryTodoStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryTodoStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryTodoStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryTodoStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryTodoStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryTodoStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryTodoStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryTodoStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryTodoStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryTodoStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryTodoStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryTodoStatisticalDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTodoStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTodoStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryTodoStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryTodoStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryTodoStatisticalDataResponse) SetStatusCode(v int32) *QueryTodoStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTodoStatisticalDataResponse) SetBody(v *QueryTodoStatisticalDataResponseBody) *QueryTodoStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryTotalDataCountServiceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryTotalDataCountServiceHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryTotalDataCountServiceHeaders) GoString() string {
	return s.String()
}

func (s *QueryTotalDataCountServiceHeaders) SetCommonHeaders(v map[string]*string) *QueryTotalDataCountServiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryTotalDataCountServiceHeaders) SetXAcsDingtalkAccessToken(v string) *QueryTotalDataCountServiceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryTotalDataCountServiceRequest struct {
	DeptIds []*string `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	// example:
	//
	// 20240611
	EndDate    *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	PageNumber *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// API-xxxx
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// example:
	//
	// 20240611
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// example:
	//
	// 222
	UserId  *string   `json:"userId,omitempty" xml:"userId,omitempty"`
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s QueryTotalDataCountServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTotalDataCountServiceRequest) GoString() string {
	return s.String()
}

func (s *QueryTotalDataCountServiceRequest) SetDeptIds(v []*string) *QueryTotalDataCountServiceRequest {
	s.DeptIds = v
	return s
}

func (s *QueryTotalDataCountServiceRequest) SetEndDate(v string) *QueryTotalDataCountServiceRequest {
	s.EndDate = &v
	return s
}

func (s *QueryTotalDataCountServiceRequest) SetPageNumber(v int64) *QueryTotalDataCountServiceRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryTotalDataCountServiceRequest) SetPageSize(v int64) *QueryTotalDataCountServiceRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTotalDataCountServiceRequest) SetServiceId(v string) *QueryTotalDataCountServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *QueryTotalDataCountServiceRequest) SetStartDate(v string) *QueryTotalDataCountServiceRequest {
	s.StartDate = &v
	return s
}

func (s *QueryTotalDataCountServiceRequest) SetUserId(v string) *QueryTotalDataCountServiceRequest {
	s.UserId = &v
	return s
}

func (s *QueryTotalDataCountServiceRequest) SetUserIds(v []*string) *QueryTotalDataCountServiceRequest {
	s.UserIds = v
	return s
}

type QueryTotalDataCountServiceResponseBody struct {
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
	Total   *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s QueryTotalDataCountServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTotalDataCountServiceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTotalDataCountServiceResponseBody) SetSuccess(v string) *QueryTotalDataCountServiceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryTotalDataCountServiceResponseBody) SetTotal(v int64) *QueryTotalDataCountServiceResponseBody {
	s.Total = &v
	return s
}

type QueryTotalDataCountServiceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTotalDataCountServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTotalDataCountServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTotalDataCountServiceResponse) GoString() string {
	return s.String()
}

func (s *QueryTotalDataCountServiceResponse) SetHeaders(v map[string]*string) *QueryTotalDataCountServiceResponse {
	s.Headers = v
	return s
}

func (s *QueryTotalDataCountServiceResponse) SetStatusCode(v int32) *QueryTotalDataCountServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTotalDataCountServiceResponse) SetBody(v *QueryTotalDataCountServiceResponseBody) *QueryTotalDataCountServiceResponse {
	s.Body = v
	return s
}

type QueryVedioMeetingStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryVedioMeetingStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryVedioMeetingStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryVedioMeetingStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryVedioMeetingStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryVedioMeetingStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryVedioMeetingStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryVedioMeetingStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryVedioMeetingStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryVedioMeetingStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryVedioMeetingStatisticalDataRequest) SetStatDate(v string) *QueryVedioMeetingStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryVedioMeetingStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryVedioMeetingStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryVedioMeetingStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryVedioMeetingStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVedioMeetingStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryVedioMeetingStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryVedioMeetingStatisticalDataResponseBody) SetMetaList(v []*QueryVedioMeetingStatisticalDataResponseBodyMetaList) *QueryVedioMeetingStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryVedioMeetingStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryVedioMeetingStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryVedioMeetingStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryVedioMeetingStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryVedioMeetingStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryVedioMeetingStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryVedioMeetingStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryVedioMeetingStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryVedioMeetingStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryVedioMeetingStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryVedioMeetingStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryVedioMeetingStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryVedioMeetingStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryVedioMeetingStatisticalDataResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryVedioMeetingStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryVedioMeetingStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryVedioMeetingStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryVedioMeetingStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryVedioMeetingStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryVedioMeetingStatisticalDataResponse) SetStatusCode(v int32) *QueryVedioMeetingStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVedioMeetingStatisticalDataResponse) SetBody(v *QueryVedioMeetingStatisticalDataResponseBody) *QueryVedioMeetingStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydActiveDayStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydActiveDayStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydActiveDayStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydActiveDayStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydActiveDayStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydActiveDayStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydActiveDayStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydActiveDayStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydActiveDayStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydActiveDayStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydActiveDayStatisticalDataRequest) SetStatDate(v string) *QueryYydActiveDayStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydActiveDayStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydActiveDayStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydActiveDayStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydActiveDayStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydActiveDayStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydActiveDayStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydActiveDayStatisticalDataResponseBody) SetMetaList(v []*QueryYydActiveDayStatisticalDataResponseBodyMetaList) *QueryYydActiveDayStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydActiveDayStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydActiveDayStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydActiveDayStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydActiveDayStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydActiveDayStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydActiveDayStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydActiveDayStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydActiveDayStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydActiveDayStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydActiveDayStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydActiveDayStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydActiveDayStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydActiveDayStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydActiveDayStatisticalDataResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydActiveDayStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydActiveDayStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydActiveDayStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydActiveDayStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydActiveDayStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydActiveDayStatisticalDataResponse) SetStatusCode(v int32) *QueryYydActiveDayStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydActiveDayStatisticalDataResponse) SetBody(v *QueryYydActiveDayStatisticalDataResponseBody) *QueryYydActiveDayStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydActiveMonthStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydActiveMonthStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydActiveMonthStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydActiveMonthStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydActiveMonthStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydActiveMonthStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydActiveMonthStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydActiveMonthStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydActiveMonthStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydActiveMonthStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydActiveMonthStatisticalDataRequest) SetStatDate(v string) *QueryYydActiveMonthStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydActiveMonthStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                  `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydActiveMonthStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydActiveMonthStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydActiveMonthStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydActiveMonthStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydActiveMonthStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydActiveMonthStatisticalDataResponseBody) SetMetaList(v []*QueryYydActiveMonthStatisticalDataResponseBodyMetaList) *QueryYydActiveMonthStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydActiveMonthStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydActiveMonthStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydActiveMonthStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydActiveMonthStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydActiveMonthStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydActiveMonthStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydActiveMonthStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydActiveMonthStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydActiveMonthStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydActiveMonthStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydActiveMonthStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydActiveMonthStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydActiveMonthStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydActiveMonthStatisticalDataResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydActiveMonthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydActiveMonthStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydActiveMonthStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydActiveMonthStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydActiveMonthStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydActiveMonthStatisticalDataResponse) SetStatusCode(v int32) *QueryYydActiveMonthStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydActiveMonthStatisticalDataResponse) SetBody(v *QueryYydActiveMonthStatisticalDataResponseBody) *QueryYydActiveMonthStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydActiveWeekStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydActiveWeekStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydActiveWeekStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydActiveWeekStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydActiveWeekStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydActiveWeekStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydActiveWeekStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydActiveWeekStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydActiveWeekStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydActiveWeekStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydActiveWeekStatisticalDataRequest) SetStatDate(v string) *QueryYydActiveWeekStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydActiveWeekStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                 `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydActiveWeekStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydActiveWeekStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydActiveWeekStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydActiveWeekStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydActiveWeekStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydActiveWeekStatisticalDataResponseBody) SetMetaList(v []*QueryYydActiveWeekStatisticalDataResponseBodyMetaList) *QueryYydActiveWeekStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydActiveWeekStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydActiveWeekStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydActiveWeekStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydActiveWeekStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydActiveWeekStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydActiveWeekStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydActiveWeekStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydActiveWeekStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydActiveWeekStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydActiveWeekStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydActiveWeekStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydActiveWeekStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydActiveWeekStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydActiveWeekStatisticalDataResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydActiveWeekStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydActiveWeekStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydActiveWeekStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydActiveWeekStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydActiveWeekStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydActiveWeekStatisticalDataResponse) SetStatusCode(v int32) *QueryYydActiveWeekStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydActiveWeekStatisticalDataResponse) SetBody(v *QueryYydActiveWeekStatisticalDataResponseBody) *QueryYydActiveWeekStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydAppDayStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydAppDayStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydAppDayStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydAppDayStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydAppDayStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydAppDayStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydAppDayStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydAppDayStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydAppDayStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydAppDayStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydAppDayStatisticalDataRequest) SetStatDate(v string) *QueryYydAppDayStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydAppDayStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                             `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydAppDayStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydAppDayStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydAppDayStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydAppDayStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydAppDayStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydAppDayStatisticalDataResponseBody) SetMetaList(v []*QueryYydAppDayStatisticalDataResponseBodyMetaList) *QueryYydAppDayStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydAppDayStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydAppDayStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydAppDayStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydAppDayStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydAppDayStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydAppDayStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydAppDayStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydAppDayStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydAppDayStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydAppDayStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydAppDayStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydAppDayStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydAppDayStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydAppDayStatisticalDataResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydAppDayStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydAppDayStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydAppDayStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydAppDayStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydAppDayStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydAppDayStatisticalDataResponse) SetStatusCode(v int32) *QueryYydAppDayStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydAppDayStatisticalDataResponse) SetBody(v *QueryYydAppDayStatisticalDataResponseBody) *QueryYydAppDayStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydAppMonthStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydAppMonthStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydAppMonthStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydAppMonthStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydAppMonthStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydAppMonthStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydAppMonthStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydAppMonthStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydAppMonthStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydAppMonthStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydAppMonthStatisticalDataRequest) SetStatDate(v string) *QueryYydAppMonthStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydAppMonthStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                               `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydAppMonthStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydAppMonthStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydAppMonthStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydAppMonthStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydAppMonthStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydAppMonthStatisticalDataResponseBody) SetMetaList(v []*QueryYydAppMonthStatisticalDataResponseBodyMetaList) *QueryYydAppMonthStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydAppMonthStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydAppMonthStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydAppMonthStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydAppMonthStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydAppMonthStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydAppMonthStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydAppMonthStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydAppMonthStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydAppMonthStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydAppMonthStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydAppMonthStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydAppMonthStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydAppMonthStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydAppMonthStatisticalDataResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydAppMonthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydAppMonthStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydAppMonthStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydAppMonthStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydAppMonthStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydAppMonthStatisticalDataResponse) SetStatusCode(v int32) *QueryYydAppMonthStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydAppMonthStatisticalDataResponse) SetBody(v *QueryYydAppMonthStatisticalDataResponseBody) *QueryYydAppMonthStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydAppStdStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydAppStdStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydAppStdStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydAppStdStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydAppStdStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydAppStdStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydAppStdStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydAppStdStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydAppStdStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydAppStdStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydAppStdStatisticalDataRequest) SetStatDate(v string) *QueryYydAppStdStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydAppStdStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                             `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydAppStdStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydAppStdStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydAppStdStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydAppStdStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydAppStdStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydAppStdStatisticalDataResponseBody) SetMetaList(v []*QueryYydAppStdStatisticalDataResponseBodyMetaList) *QueryYydAppStdStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydAppStdStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydAppStdStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydAppStdStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydAppStdStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydAppStdStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydAppStdStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydAppStdStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydAppStdStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydAppStdStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydAppStdStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydAppStdStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydAppStdStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydAppStdStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydAppStdStatisticalDataResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydAppStdStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydAppStdStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydAppStdStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydAppStdStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydAppStdStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydAppStdStatisticalDataResponse) SetStatusCode(v int32) *QueryYydAppStdStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydAppStdStatisticalDataResponse) SetBody(v *QueryYydAppStdStatisticalDataResponseBody) *QueryYydAppStdStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydAppWeekStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydAppWeekStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydAppWeekStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydAppWeekStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydAppWeekStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydAppWeekStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydAppWeekStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydAppWeekStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydAppWeekStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydAppWeekStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydAppWeekStatisticalDataRequest) SetStatDate(v string) *QueryYydAppWeekStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydAppWeekStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                              `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydAppWeekStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydAppWeekStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydAppWeekStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydAppWeekStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydAppWeekStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydAppWeekStatisticalDataResponseBody) SetMetaList(v []*QueryYydAppWeekStatisticalDataResponseBodyMetaList) *QueryYydAppWeekStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydAppWeekStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydAppWeekStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydAppWeekStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydAppWeekStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydAppWeekStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydAppWeekStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydAppWeekStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydAppWeekStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydAppWeekStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydAppWeekStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydAppWeekStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydAppWeekStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydAppWeekStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydAppWeekStatisticalDataResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydAppWeekStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydAppWeekStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydAppWeekStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydAppWeekStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydAppWeekStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydAppWeekStatisticalDataResponse) SetStatusCode(v int32) *QueryYydAppWeekStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydAppWeekStatisticalDataResponse) SetBody(v *QueryYydAppWeekStatisticalDataResponseBody) *QueryYydAppWeekStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydCalendarDayStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydCalendarDayStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydCalendarDayStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydCalendarDayStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydCalendarDayStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydCalendarDayStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydCalendarDayStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydCalendarDayStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydCalendarDayStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydCalendarDayStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydCalendarDayStatisticalDataRequest) SetStatDate(v string) *QueryYydCalendarDayStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydCalendarDayStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                  `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydCalendarDayStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydCalendarDayStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydCalendarDayStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydCalendarDayStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydCalendarDayStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydCalendarDayStatisticalDataResponseBody) SetMetaList(v []*QueryYydCalendarDayStatisticalDataResponseBodyMetaList) *QueryYydCalendarDayStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydCalendarDayStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydCalendarDayStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydCalendarDayStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydCalendarDayStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydCalendarDayStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydCalendarDayStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydCalendarDayStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydCalendarDayStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydCalendarDayStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydCalendarDayStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydCalendarDayStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydCalendarDayStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydCalendarDayStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydCalendarDayStatisticalDataResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydCalendarDayStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydCalendarDayStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydCalendarDayStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydCalendarDayStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydCalendarDayStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydCalendarDayStatisticalDataResponse) SetStatusCode(v int32) *QueryYydCalendarDayStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydCalendarDayStatisticalDataResponse) SetBody(v *QueryYydCalendarDayStatisticalDataResponseBody) *QueryYydCalendarDayStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydCalendarMonthStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydCalendarMonthStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydCalendarMonthStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydCalendarMonthStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydCalendarMonthStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydCalendarMonthStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydCalendarMonthStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydCalendarMonthStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydCalendarMonthStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydCalendarMonthStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydCalendarMonthStatisticalDataRequest) SetStatDate(v string) *QueryYydCalendarMonthStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydCalendarMonthStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                    `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydCalendarMonthStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydCalendarMonthStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydCalendarMonthStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydCalendarMonthStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydCalendarMonthStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydCalendarMonthStatisticalDataResponseBody) SetMetaList(v []*QueryYydCalendarMonthStatisticalDataResponseBodyMetaList) *QueryYydCalendarMonthStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydCalendarMonthStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydCalendarMonthStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydCalendarMonthStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydCalendarMonthStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydCalendarMonthStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydCalendarMonthStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydCalendarMonthStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydCalendarMonthStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydCalendarMonthStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydCalendarMonthStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydCalendarMonthStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydCalendarMonthStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydCalendarMonthStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydCalendarMonthStatisticalDataResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydCalendarMonthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydCalendarMonthStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydCalendarMonthStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydCalendarMonthStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydCalendarMonthStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydCalendarMonthStatisticalDataResponse) SetStatusCode(v int32) *QueryYydCalendarMonthStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydCalendarMonthStatisticalDataResponse) SetBody(v *QueryYydCalendarMonthStatisticalDataResponseBody) *QueryYydCalendarMonthStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydCalendarWeekStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydCalendarWeekStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydCalendarWeekStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydCalendarWeekStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydCalendarWeekStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydCalendarWeekStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydCalendarWeekStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydCalendarWeekStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydCalendarWeekStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydCalendarWeekStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydCalendarWeekStatisticalDataRequest) SetStatDate(v string) *QueryYydCalendarWeekStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydCalendarWeekStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                   `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydCalendarWeekStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydCalendarWeekStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydCalendarWeekStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydCalendarWeekStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydCalendarWeekStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydCalendarWeekStatisticalDataResponseBody) SetMetaList(v []*QueryYydCalendarWeekStatisticalDataResponseBodyMetaList) *QueryYydCalendarWeekStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydCalendarWeekStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydCalendarWeekStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydCalendarWeekStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydCalendarWeekStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydCalendarWeekStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydCalendarWeekStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydCalendarWeekStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydCalendarWeekStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydCalendarWeekStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydCalendarWeekStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydCalendarWeekStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydCalendarWeekStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydCalendarWeekStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydCalendarWeekStatisticalDataResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydCalendarWeekStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydCalendarWeekStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydCalendarWeekStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydCalendarWeekStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydCalendarWeekStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydCalendarWeekStatisticalDataResponse) SetStatusCode(v int32) *QueryYydCalendarWeekStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydCalendarWeekStatisticalDataResponse) SetBody(v *QueryYydCalendarWeekStatisticalDataResponseBody) *QueryYydCalendarWeekStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydDingMsgDayStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydDingMsgDayStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydDingMsgDayStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydDingMsgDayStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydDingMsgDayStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydDingMsgDayStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydDingMsgDayStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydDingMsgDayStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydDingMsgDayStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydDingMsgDayStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydDingMsgDayStatisticalDataRequest) SetStatDate(v string) *QueryYydDingMsgDayStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydDingMsgDayStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                 `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydDingMsgDayStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydDingMsgDayStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydDingMsgDayStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydDingMsgDayStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydDingMsgDayStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydDingMsgDayStatisticalDataResponseBody) SetMetaList(v []*QueryYydDingMsgDayStatisticalDataResponseBodyMetaList) *QueryYydDingMsgDayStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydDingMsgDayStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydDingMsgDayStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydDingMsgDayStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydDingMsgDayStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydDingMsgDayStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydDingMsgDayStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydDingMsgDayStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydDingMsgDayStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydDingMsgDayStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydDingMsgDayStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydDingMsgDayStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydDingMsgDayStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydDingMsgDayStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydDingMsgDayStatisticalDataResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydDingMsgDayStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydDingMsgDayStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydDingMsgDayStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydDingMsgDayStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydDingMsgDayStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydDingMsgDayStatisticalDataResponse) SetStatusCode(v int32) *QueryYydDingMsgDayStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydDingMsgDayStatisticalDataResponse) SetBody(v *QueryYydDingMsgDayStatisticalDataResponseBody) *QueryYydDingMsgDayStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydDingMsgMonthStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydDingMsgMonthStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydDingMsgMonthStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydDingMsgMonthStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydDingMsgMonthStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydDingMsgMonthStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydDingMsgMonthStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydDingMsgMonthStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydDingMsgMonthStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydDingMsgMonthStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydDingMsgMonthStatisticalDataRequest) SetStatDate(v string) *QueryYydDingMsgMonthStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydDingMsgMonthStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                   `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydDingMsgMonthStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydDingMsgMonthStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydDingMsgMonthStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydDingMsgMonthStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydDingMsgMonthStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydDingMsgMonthStatisticalDataResponseBody) SetMetaList(v []*QueryYydDingMsgMonthStatisticalDataResponseBodyMetaList) *QueryYydDingMsgMonthStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydDingMsgMonthStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydDingMsgMonthStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydDingMsgMonthStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydDingMsgMonthStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydDingMsgMonthStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydDingMsgMonthStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydDingMsgMonthStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydDingMsgMonthStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydDingMsgMonthStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydDingMsgMonthStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydDingMsgMonthStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydDingMsgMonthStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydDingMsgMonthStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydDingMsgMonthStatisticalDataResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydDingMsgMonthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydDingMsgMonthStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydDingMsgMonthStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydDingMsgMonthStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydDingMsgMonthStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydDingMsgMonthStatisticalDataResponse) SetStatusCode(v int32) *QueryYydDingMsgMonthStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydDingMsgMonthStatisticalDataResponse) SetBody(v *QueryYydDingMsgMonthStatisticalDataResponseBody) *QueryYydDingMsgMonthStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydDingMsgWeekStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydDingMsgWeekStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydDingMsgWeekStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydDingMsgWeekStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydDingMsgWeekStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydDingMsgWeekStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydDingMsgWeekStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydDingMsgWeekStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydDingMsgWeekStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydDingMsgWeekStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydDingMsgWeekStatisticalDataRequest) SetStatDate(v string) *QueryYydDingMsgWeekStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydDingMsgWeekStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                  `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydDingMsgWeekStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydDingMsgWeekStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydDingMsgWeekStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydDingMsgWeekStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydDingMsgWeekStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydDingMsgWeekStatisticalDataResponseBody) SetMetaList(v []*QueryYydDingMsgWeekStatisticalDataResponseBodyMetaList) *QueryYydDingMsgWeekStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydDingMsgWeekStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydDingMsgWeekStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydDingMsgWeekStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydDingMsgWeekStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydDingMsgWeekStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydDingMsgWeekStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydDingMsgWeekStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydDingMsgWeekStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydDingMsgWeekStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydDingMsgWeekStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydDingMsgWeekStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydDingMsgWeekStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydDingMsgWeekStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydDingMsgWeekStatisticalDataResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydDingMsgWeekStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydDingMsgWeekStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydDingMsgWeekStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydDingMsgWeekStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydDingMsgWeekStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydDingMsgWeekStatisticalDataResponse) SetStatusCode(v int32) *QueryYydDingMsgWeekStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydDingMsgWeekStatisticalDataResponse) SetBody(v *QueryYydDingMsgWeekStatisticalDataResponseBody) *QueryYydDingMsgWeekStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydGroupMsgDayStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydGroupMsgDayStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydGroupMsgDayStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydGroupMsgDayStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydGroupMsgDayStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydGroupMsgDayStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydGroupMsgDayStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydGroupMsgDayStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydGroupMsgDayStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydGroupMsgDayStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydGroupMsgDayStatisticalDataRequest) SetStatDate(v string) *QueryYydGroupMsgDayStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydGroupMsgDayStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                  `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydGroupMsgDayStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydGroupMsgDayStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydGroupMsgDayStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydGroupMsgDayStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydGroupMsgDayStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydGroupMsgDayStatisticalDataResponseBody) SetMetaList(v []*QueryYydGroupMsgDayStatisticalDataResponseBodyMetaList) *QueryYydGroupMsgDayStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydGroupMsgDayStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydGroupMsgDayStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydGroupMsgDayStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydGroupMsgDayStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydGroupMsgDayStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydGroupMsgDayStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydGroupMsgDayStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydGroupMsgDayStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydGroupMsgDayStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydGroupMsgDayStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydGroupMsgDayStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydGroupMsgDayStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydGroupMsgDayStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydGroupMsgDayStatisticalDataResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydGroupMsgDayStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydGroupMsgDayStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydGroupMsgDayStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydGroupMsgDayStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydGroupMsgDayStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydGroupMsgDayStatisticalDataResponse) SetStatusCode(v int32) *QueryYydGroupMsgDayStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydGroupMsgDayStatisticalDataResponse) SetBody(v *QueryYydGroupMsgDayStatisticalDataResponseBody) *QueryYydGroupMsgDayStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydGroupMsgMonthStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydGroupMsgMonthStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydGroupMsgMonthStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydGroupMsgMonthStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydGroupMsgMonthStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydGroupMsgMonthStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydGroupMsgMonthStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydGroupMsgMonthStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydGroupMsgMonthStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydGroupMsgMonthStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydGroupMsgMonthStatisticalDataRequest) SetStatDate(v string) *QueryYydGroupMsgMonthStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydGroupMsgMonthStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                    `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydGroupMsgMonthStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydGroupMsgMonthStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydGroupMsgMonthStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydGroupMsgMonthStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydGroupMsgMonthStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydGroupMsgMonthStatisticalDataResponseBody) SetMetaList(v []*QueryYydGroupMsgMonthStatisticalDataResponseBodyMetaList) *QueryYydGroupMsgMonthStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydGroupMsgMonthStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydGroupMsgMonthStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydGroupMsgMonthStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydGroupMsgMonthStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydGroupMsgMonthStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydGroupMsgMonthStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydGroupMsgMonthStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydGroupMsgMonthStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydGroupMsgMonthStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydGroupMsgMonthStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydGroupMsgMonthStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydGroupMsgMonthStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydGroupMsgMonthStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydGroupMsgMonthStatisticalDataResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydGroupMsgMonthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydGroupMsgMonthStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydGroupMsgMonthStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydGroupMsgMonthStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydGroupMsgMonthStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydGroupMsgMonthStatisticalDataResponse) SetStatusCode(v int32) *QueryYydGroupMsgMonthStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydGroupMsgMonthStatisticalDataResponse) SetBody(v *QueryYydGroupMsgMonthStatisticalDataResponseBody) *QueryYydGroupMsgMonthStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydGroupMsgWeekStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydGroupMsgWeekStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydGroupMsgWeekStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydGroupMsgWeekStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydGroupMsgWeekStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydGroupMsgWeekStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydGroupMsgWeekStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydGroupMsgWeekStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydGroupMsgWeekStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydGroupMsgWeekStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydGroupMsgWeekStatisticalDataRequest) SetStatDate(v string) *QueryYydGroupMsgWeekStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydGroupMsgWeekStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                   `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydGroupMsgWeekStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydGroupMsgWeekStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydGroupMsgWeekStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydGroupMsgWeekStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydGroupMsgWeekStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydGroupMsgWeekStatisticalDataResponseBody) SetMetaList(v []*QueryYydGroupMsgWeekStatisticalDataResponseBodyMetaList) *QueryYydGroupMsgWeekStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydGroupMsgWeekStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydGroupMsgWeekStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydGroupMsgWeekStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydGroupMsgWeekStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydGroupMsgWeekStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydGroupMsgWeekStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydGroupMsgWeekStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydGroupMsgWeekStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydGroupMsgWeekStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydGroupMsgWeekStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydGroupMsgWeekStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydGroupMsgWeekStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydGroupMsgWeekStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydGroupMsgWeekStatisticalDataResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydGroupMsgWeekStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydGroupMsgWeekStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydGroupMsgWeekStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydGroupMsgWeekStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydGroupMsgWeekStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydGroupMsgWeekStatisticalDataResponse) SetStatusCode(v int32) *QueryYydGroupMsgWeekStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydGroupMsgWeekStatisticalDataResponse) SetBody(v *QueryYydGroupMsgWeekStatisticalDataResponseBody) *QueryYydGroupMsgWeekStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydLogDayStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydLogDayStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydLogDayStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydLogDayStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydLogDayStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydLogDayStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydLogDayStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydLogDayStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydLogDayStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydLogDayStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydLogDayStatisticalDataRequest) SetStatDate(v string) *QueryYydLogDayStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydLogDayStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                             `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydLogDayStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydLogDayStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydLogDayStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydLogDayStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydLogDayStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydLogDayStatisticalDataResponseBody) SetMetaList(v []*QueryYydLogDayStatisticalDataResponseBodyMetaList) *QueryYydLogDayStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydLogDayStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydLogDayStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydLogDayStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydLogDayStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydLogDayStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydLogDayStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydLogDayStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydLogDayStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydLogDayStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydLogDayStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydLogDayStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydLogDayStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydLogDayStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydLogDayStatisticalDataResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydLogDayStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydLogDayStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydLogDayStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydLogDayStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydLogDayStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydLogDayStatisticalDataResponse) SetStatusCode(v int32) *QueryYydLogDayStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydLogDayStatisticalDataResponse) SetBody(v *QueryYydLogDayStatisticalDataResponseBody) *QueryYydLogDayStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydLogMonthStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydLogMonthStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydLogMonthStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydLogMonthStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydLogMonthStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydLogMonthStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydLogMonthStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydLogMonthStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydLogMonthStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydLogMonthStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydLogMonthStatisticalDataRequest) SetStatDate(v string) *QueryYydLogMonthStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydLogMonthStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                               `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydLogMonthStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydLogMonthStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydLogMonthStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydLogMonthStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydLogMonthStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydLogMonthStatisticalDataResponseBody) SetMetaList(v []*QueryYydLogMonthStatisticalDataResponseBodyMetaList) *QueryYydLogMonthStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydLogMonthStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydLogMonthStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydLogMonthStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydLogMonthStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydLogMonthStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydLogMonthStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydLogMonthStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydLogMonthStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydLogMonthStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydLogMonthStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydLogMonthStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydLogMonthStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydLogMonthStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydLogMonthStatisticalDataResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydLogMonthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydLogMonthStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydLogMonthStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydLogMonthStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydLogMonthStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydLogMonthStatisticalDataResponse) SetStatusCode(v int32) *QueryYydLogMonthStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydLogMonthStatisticalDataResponse) SetBody(v *QueryYydLogMonthStatisticalDataResponseBody) *QueryYydLogMonthStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydLogWeekStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydLogWeekStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydLogWeekStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydLogWeekStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydLogWeekStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydLogWeekStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydLogWeekStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydLogWeekStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydLogWeekStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydLogWeekStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydLogWeekStatisticalDataRequest) SetStatDate(v string) *QueryYydLogWeekStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydLogWeekStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                              `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydLogWeekStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydLogWeekStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydLogWeekStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydLogWeekStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydLogWeekStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydLogWeekStatisticalDataResponseBody) SetMetaList(v []*QueryYydLogWeekStatisticalDataResponseBodyMetaList) *QueryYydLogWeekStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydLogWeekStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydLogWeekStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydLogWeekStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydLogWeekStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydLogWeekStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydLogWeekStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydLogWeekStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydLogWeekStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydLogWeekStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydLogWeekStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydLogWeekStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydLogWeekStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydLogWeekStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydLogWeekStatisticalDataResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydLogWeekStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydLogWeekStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydLogWeekStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydLogWeekStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydLogWeekStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydLogWeekStatisticalDataResponse) SetStatusCode(v int32) *QueryYydLogWeekStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydLogWeekStatisticalDataResponse) SetBody(v *QueryYydLogWeekStatisticalDataResponseBody) *QueryYydLogWeekStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydMeetingDayStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydMeetingDayStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydMeetingDayStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydMeetingDayStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydMeetingDayStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydMeetingDayStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydMeetingDayStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydMeetingDayStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydMeetingDayStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydMeetingDayStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydMeetingDayStatisticalDataRequest) SetStatDate(v string) *QueryYydMeetingDayStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydMeetingDayStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                 `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydMeetingDayStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydMeetingDayStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydMeetingDayStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydMeetingDayStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydMeetingDayStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydMeetingDayStatisticalDataResponseBody) SetMetaList(v []*QueryYydMeetingDayStatisticalDataResponseBodyMetaList) *QueryYydMeetingDayStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydMeetingDayStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydMeetingDayStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydMeetingDayStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydMeetingDayStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydMeetingDayStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydMeetingDayStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydMeetingDayStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydMeetingDayStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydMeetingDayStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydMeetingDayStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydMeetingDayStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydMeetingDayStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydMeetingDayStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydMeetingDayStatisticalDataResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydMeetingDayStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydMeetingDayStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydMeetingDayStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydMeetingDayStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydMeetingDayStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydMeetingDayStatisticalDataResponse) SetStatusCode(v int32) *QueryYydMeetingDayStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydMeetingDayStatisticalDataResponse) SetBody(v *QueryYydMeetingDayStatisticalDataResponseBody) *QueryYydMeetingDayStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydMeetingMonthStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydMeetingMonthStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydMeetingMonthStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydMeetingMonthStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydMeetingMonthStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydMeetingMonthStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydMeetingMonthStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydMeetingMonthStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydMeetingMonthStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydMeetingMonthStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydMeetingMonthStatisticalDataRequest) SetStatDate(v string) *QueryYydMeetingMonthStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydMeetingMonthStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                   `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydMeetingMonthStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydMeetingMonthStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydMeetingMonthStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydMeetingMonthStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydMeetingMonthStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydMeetingMonthStatisticalDataResponseBody) SetMetaList(v []*QueryYydMeetingMonthStatisticalDataResponseBodyMetaList) *QueryYydMeetingMonthStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydMeetingMonthStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydMeetingMonthStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydMeetingMonthStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydMeetingMonthStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydMeetingMonthStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydMeetingMonthStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydMeetingMonthStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydMeetingMonthStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydMeetingMonthStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydMeetingMonthStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydMeetingMonthStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydMeetingMonthStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydMeetingMonthStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydMeetingMonthStatisticalDataResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydMeetingMonthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydMeetingMonthStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydMeetingMonthStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydMeetingMonthStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydMeetingMonthStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydMeetingMonthStatisticalDataResponse) SetStatusCode(v int32) *QueryYydMeetingMonthStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydMeetingMonthStatisticalDataResponse) SetBody(v *QueryYydMeetingMonthStatisticalDataResponseBody) *QueryYydMeetingMonthStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydMeetingWeekStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydMeetingWeekStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydMeetingWeekStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydMeetingWeekStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydMeetingWeekStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydMeetingWeekStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydMeetingWeekStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydMeetingWeekStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydMeetingWeekStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydMeetingWeekStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydMeetingWeekStatisticalDataRequest) SetStatDate(v string) *QueryYydMeetingWeekStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydMeetingWeekStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                  `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydMeetingWeekStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydMeetingWeekStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydMeetingWeekStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydMeetingWeekStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydMeetingWeekStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydMeetingWeekStatisticalDataResponseBody) SetMetaList(v []*QueryYydMeetingWeekStatisticalDataResponseBodyMetaList) *QueryYydMeetingWeekStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydMeetingWeekStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydMeetingWeekStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydMeetingWeekStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydMeetingWeekStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydMeetingWeekStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydMeetingWeekStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydMeetingWeekStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydMeetingWeekStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydMeetingWeekStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydMeetingWeekStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydMeetingWeekStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydMeetingWeekStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydMeetingWeekStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydMeetingWeekStatisticalDataResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydMeetingWeekStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydMeetingWeekStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydMeetingWeekStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydMeetingWeekStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydMeetingWeekStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydMeetingWeekStatisticalDataResponse) SetStatusCode(v int32) *QueryYydMeetingWeekStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydMeetingWeekStatisticalDataResponse) SetBody(v *QueryYydMeetingWeekStatisticalDataResponseBody) *QueryYydMeetingWeekStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydNoticeDayStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydNoticeDayStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydNoticeDayStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydNoticeDayStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydNoticeDayStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydNoticeDayStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydNoticeDayStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydNoticeDayStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydNoticeDayStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydNoticeDayStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydNoticeDayStatisticalDataRequest) SetStatDate(v string) *QueryYydNoticeDayStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydNoticeDayStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydNoticeDayStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydNoticeDayStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydNoticeDayStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydNoticeDayStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydNoticeDayStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydNoticeDayStatisticalDataResponseBody) SetMetaList(v []*QueryYydNoticeDayStatisticalDataResponseBodyMetaList) *QueryYydNoticeDayStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydNoticeDayStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydNoticeDayStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydNoticeDayStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydNoticeDayStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydNoticeDayStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydNoticeDayStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydNoticeDayStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydNoticeDayStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydNoticeDayStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydNoticeDayStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydNoticeDayStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydNoticeDayStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydNoticeDayStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydNoticeDayStatisticalDataResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydNoticeDayStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydNoticeDayStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydNoticeDayStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydNoticeDayStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydNoticeDayStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydNoticeDayStatisticalDataResponse) SetStatusCode(v int32) *QueryYydNoticeDayStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydNoticeDayStatisticalDataResponse) SetBody(v *QueryYydNoticeDayStatisticalDataResponseBody) *QueryYydNoticeDayStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydNoticeMonthStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydNoticeMonthStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydNoticeMonthStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydNoticeMonthStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydNoticeMonthStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydNoticeMonthStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydNoticeMonthStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydNoticeMonthStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydNoticeMonthStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydNoticeMonthStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydNoticeMonthStatisticalDataRequest) SetStatDate(v string) *QueryYydNoticeMonthStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydNoticeMonthStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                  `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydNoticeMonthStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydNoticeMonthStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydNoticeMonthStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydNoticeMonthStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydNoticeMonthStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydNoticeMonthStatisticalDataResponseBody) SetMetaList(v []*QueryYydNoticeMonthStatisticalDataResponseBodyMetaList) *QueryYydNoticeMonthStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydNoticeMonthStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydNoticeMonthStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydNoticeMonthStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydNoticeMonthStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydNoticeMonthStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydNoticeMonthStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydNoticeMonthStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydNoticeMonthStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydNoticeMonthStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydNoticeMonthStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydNoticeMonthStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydNoticeMonthStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydNoticeMonthStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydNoticeMonthStatisticalDataResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydNoticeMonthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydNoticeMonthStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydNoticeMonthStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydNoticeMonthStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydNoticeMonthStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydNoticeMonthStatisticalDataResponse) SetStatusCode(v int32) *QueryYydNoticeMonthStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydNoticeMonthStatisticalDataResponse) SetBody(v *QueryYydNoticeMonthStatisticalDataResponseBody) *QueryYydNoticeMonthStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydNoticeWeekStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydNoticeWeekStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydNoticeWeekStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydNoticeWeekStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydNoticeWeekStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydNoticeWeekStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydNoticeWeekStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydNoticeWeekStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydNoticeWeekStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydNoticeWeekStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydNoticeWeekStatisticalDataRequest) SetStatDate(v string) *QueryYydNoticeWeekStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydNoticeWeekStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                 `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydNoticeWeekStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydNoticeWeekStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydNoticeWeekStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydNoticeWeekStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydNoticeWeekStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydNoticeWeekStatisticalDataResponseBody) SetMetaList(v []*QueryYydNoticeWeekStatisticalDataResponseBodyMetaList) *QueryYydNoticeWeekStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydNoticeWeekStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydNoticeWeekStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydNoticeWeekStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydNoticeWeekStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydNoticeWeekStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydNoticeWeekStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydNoticeWeekStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydNoticeWeekStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydNoticeWeekStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydNoticeWeekStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydNoticeWeekStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydNoticeWeekStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydNoticeWeekStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydNoticeWeekStatisticalDataResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydNoticeWeekStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydNoticeWeekStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydNoticeWeekStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydNoticeWeekStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydNoticeWeekStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydNoticeWeekStatisticalDataResponse) SetStatusCode(v int32) *QueryYydNoticeWeekStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydNoticeWeekStatisticalDataResponse) SetBody(v *QueryYydNoticeWeekStatisticalDataResponseBody) *QueryYydNoticeWeekStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydSingleMsgDayStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydSingleMsgDayStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydSingleMsgDayStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydSingleMsgDayStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydSingleMsgDayStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydSingleMsgDayStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydSingleMsgDayStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydSingleMsgDayStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydSingleMsgDayStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydSingleMsgDayStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydSingleMsgDayStatisticalDataRequest) SetStatDate(v string) *QueryYydSingleMsgDayStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydSingleMsgDayStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                   `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydSingleMsgDayStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydSingleMsgDayStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydSingleMsgDayStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydSingleMsgDayStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydSingleMsgDayStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydSingleMsgDayStatisticalDataResponseBody) SetMetaList(v []*QueryYydSingleMsgDayStatisticalDataResponseBodyMetaList) *QueryYydSingleMsgDayStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydSingleMsgDayStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydSingleMsgDayStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydSingleMsgDayStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydSingleMsgDayStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydSingleMsgDayStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydSingleMsgDayStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydSingleMsgDayStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydSingleMsgDayStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydSingleMsgDayStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydSingleMsgDayStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydSingleMsgDayStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydSingleMsgDayStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydSingleMsgDayStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydSingleMsgDayStatisticalDataResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydSingleMsgDayStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydSingleMsgDayStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydSingleMsgDayStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydSingleMsgDayStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydSingleMsgDayStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydSingleMsgDayStatisticalDataResponse) SetStatusCode(v int32) *QueryYydSingleMsgDayStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydSingleMsgDayStatisticalDataResponse) SetBody(v *QueryYydSingleMsgDayStatisticalDataResponseBody) *QueryYydSingleMsgDayStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydSingleMsgMonthStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydSingleMsgMonthStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydSingleMsgMonthStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydSingleMsgMonthStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydSingleMsgMonthStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydSingleMsgMonthStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydSingleMsgMonthStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydSingleMsgMonthStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydSingleMsgMonthStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydSingleMsgMonthStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydSingleMsgMonthStatisticalDataRequest) SetStatDate(v string) *QueryYydSingleMsgMonthStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydSingleMsgMonthStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                     `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydSingleMsgMonthStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydSingleMsgMonthStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydSingleMsgMonthStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydSingleMsgMonthStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydSingleMsgMonthStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydSingleMsgMonthStatisticalDataResponseBody) SetMetaList(v []*QueryYydSingleMsgMonthStatisticalDataResponseBodyMetaList) *QueryYydSingleMsgMonthStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydSingleMsgMonthStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydSingleMsgMonthStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydSingleMsgMonthStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydSingleMsgMonthStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydSingleMsgMonthStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydSingleMsgMonthStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydSingleMsgMonthStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydSingleMsgMonthStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydSingleMsgMonthStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydSingleMsgMonthStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydSingleMsgMonthStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydSingleMsgMonthStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydSingleMsgMonthStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydSingleMsgMonthStatisticalDataResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydSingleMsgMonthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydSingleMsgMonthStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydSingleMsgMonthStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydSingleMsgMonthStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydSingleMsgMonthStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydSingleMsgMonthStatisticalDataResponse) SetStatusCode(v int32) *QueryYydSingleMsgMonthStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydSingleMsgMonthStatisticalDataResponse) SetBody(v *QueryYydSingleMsgMonthStatisticalDataResponseBody) *QueryYydSingleMsgMonthStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydSingleMsgWeekStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydSingleMsgWeekStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydSingleMsgWeekStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydSingleMsgWeekStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydSingleMsgWeekStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydSingleMsgWeekStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydSingleMsgWeekStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydSingleMsgWeekStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydSingleMsgWeekStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydSingleMsgWeekStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydSingleMsgWeekStatisticalDataRequest) SetStatDate(v string) *QueryYydSingleMsgWeekStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydSingleMsgWeekStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                    `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydSingleMsgWeekStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydSingleMsgWeekStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydSingleMsgWeekStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydSingleMsgWeekStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydSingleMsgWeekStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydSingleMsgWeekStatisticalDataResponseBody) SetMetaList(v []*QueryYydSingleMsgWeekStatisticalDataResponseBodyMetaList) *QueryYydSingleMsgWeekStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydSingleMsgWeekStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydSingleMsgWeekStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydSingleMsgWeekStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydSingleMsgWeekStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydSingleMsgWeekStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydSingleMsgWeekStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydSingleMsgWeekStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydSingleMsgWeekStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydSingleMsgWeekStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydSingleMsgWeekStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydSingleMsgWeekStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydSingleMsgWeekStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydSingleMsgWeekStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydSingleMsgWeekStatisticalDataResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydSingleMsgWeekStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydSingleMsgWeekStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydSingleMsgWeekStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydSingleMsgWeekStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydSingleMsgWeekStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydSingleMsgWeekStatisticalDataResponse) SetStatusCode(v int32) *QueryYydSingleMsgWeekStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydSingleMsgWeekStatisticalDataResponse) SetBody(v *QueryYydSingleMsgWeekStatisticalDataResponseBody) *QueryYydSingleMsgWeekStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydToatlMsgDayStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydToatlMsgDayStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydToatlMsgDayStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydToatlMsgDayStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydToatlMsgDayStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydToatlMsgDayStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydToatlMsgDayStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydToatlMsgDayStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydToatlMsgDayStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydToatlMsgDayStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydToatlMsgDayStatisticalDataRequest) SetStatDate(v string) *QueryYydToatlMsgDayStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydToatlMsgDayStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                  `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydToatlMsgDayStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydToatlMsgDayStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydToatlMsgDayStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydToatlMsgDayStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydToatlMsgDayStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydToatlMsgDayStatisticalDataResponseBody) SetMetaList(v []*QueryYydToatlMsgDayStatisticalDataResponseBodyMetaList) *QueryYydToatlMsgDayStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydToatlMsgDayStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydToatlMsgDayStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydToatlMsgDayStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydToatlMsgDayStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydToatlMsgDayStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydToatlMsgDayStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydToatlMsgDayStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydToatlMsgDayStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydToatlMsgDayStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydToatlMsgDayStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydToatlMsgDayStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydToatlMsgDayStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydToatlMsgDayStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydToatlMsgDayStatisticalDataResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydToatlMsgDayStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydToatlMsgDayStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydToatlMsgDayStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydToatlMsgDayStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydToatlMsgDayStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydToatlMsgDayStatisticalDataResponse) SetStatusCode(v int32) *QueryYydToatlMsgDayStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydToatlMsgDayStatisticalDataResponse) SetBody(v *QueryYydToatlMsgDayStatisticalDataResponseBody) *QueryYydToatlMsgDayStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydToatlMsgMonthStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydToatlMsgMonthStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydToatlMsgMonthStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydToatlMsgMonthStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydToatlMsgMonthStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydToatlMsgMonthStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydToatlMsgMonthStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydToatlMsgMonthStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydToatlMsgMonthStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydToatlMsgMonthStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydToatlMsgMonthStatisticalDataRequest) SetStatDate(v string) *QueryYydToatlMsgMonthStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydToatlMsgMonthStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                    `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydToatlMsgMonthStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydToatlMsgMonthStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydToatlMsgMonthStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydToatlMsgMonthStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydToatlMsgMonthStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydToatlMsgMonthStatisticalDataResponseBody) SetMetaList(v []*QueryYydToatlMsgMonthStatisticalDataResponseBodyMetaList) *QueryYydToatlMsgMonthStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydToatlMsgMonthStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydToatlMsgMonthStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydToatlMsgMonthStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydToatlMsgMonthStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydToatlMsgMonthStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydToatlMsgMonthStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydToatlMsgMonthStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydToatlMsgMonthStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydToatlMsgMonthStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydToatlMsgMonthStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydToatlMsgMonthStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydToatlMsgMonthStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydToatlMsgMonthStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydToatlMsgMonthStatisticalDataResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydToatlMsgMonthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydToatlMsgMonthStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydToatlMsgMonthStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydToatlMsgMonthStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydToatlMsgMonthStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydToatlMsgMonthStatisticalDataResponse) SetStatusCode(v int32) *QueryYydToatlMsgMonthStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydToatlMsgMonthStatisticalDataResponse) SetBody(v *QueryYydToatlMsgMonthStatisticalDataResponseBody) *QueryYydToatlMsgMonthStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydToatlMsgWeekStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydToatlMsgWeekStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydToatlMsgWeekStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydToatlMsgWeekStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydToatlMsgWeekStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydToatlMsgWeekStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydToatlMsgWeekStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydToatlMsgWeekStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydToatlMsgWeekStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydToatlMsgWeekStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydToatlMsgWeekStatisticalDataRequest) SetStatDate(v string) *QueryYydToatlMsgWeekStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydToatlMsgWeekStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                   `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydToatlMsgWeekStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydToatlMsgWeekStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydToatlMsgWeekStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydToatlMsgWeekStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydToatlMsgWeekStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydToatlMsgWeekStatisticalDataResponseBody) SetMetaList(v []*QueryYydToatlMsgWeekStatisticalDataResponseBodyMetaList) *QueryYydToatlMsgWeekStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydToatlMsgWeekStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydToatlMsgWeekStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydToatlMsgWeekStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydToatlMsgWeekStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydToatlMsgWeekStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydToatlMsgWeekStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydToatlMsgWeekStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydToatlMsgWeekStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydToatlMsgWeekStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydToatlMsgWeekStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydToatlMsgWeekStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydToatlMsgWeekStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydToatlMsgWeekStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydToatlMsgWeekStatisticalDataResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydToatlMsgWeekStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydToatlMsgWeekStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydToatlMsgWeekStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydToatlMsgWeekStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydToatlMsgWeekStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydToatlMsgWeekStatisticalDataResponse) SetStatusCode(v int32) *QueryYydToatlMsgWeekStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydToatlMsgWeekStatisticalDataResponse) SetBody(v *QueryYydToatlMsgWeekStatisticalDataResponseBody) *QueryYydToatlMsgWeekStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydTodoDayStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydTodoDayStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTodoDayStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydTodoDayStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydTodoDayStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydTodoDayStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydTodoDayStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydTodoDayStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydTodoDayStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTodoDayStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydTodoDayStatisticalDataRequest) SetStatDate(v string) *QueryYydTodoDayStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydTodoDayStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                              `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydTodoDayStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydTodoDayStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTodoDayStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydTodoDayStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydTodoDayStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydTodoDayStatisticalDataResponseBody) SetMetaList(v []*QueryYydTodoDayStatisticalDataResponseBodyMetaList) *QueryYydTodoDayStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydTodoDayStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydTodoDayStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTodoDayStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydTodoDayStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydTodoDayStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydTodoDayStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydTodoDayStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydTodoDayStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydTodoDayStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydTodoDayStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydTodoDayStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydTodoDayStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydTodoDayStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydTodoDayStatisticalDataResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydTodoDayStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydTodoDayStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTodoDayStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydTodoDayStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydTodoDayStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydTodoDayStatisticalDataResponse) SetStatusCode(v int32) *QueryYydTodoDayStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydTodoDayStatisticalDataResponse) SetBody(v *QueryYydTodoDayStatisticalDataResponseBody) *QueryYydTodoDayStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydTodoMonthStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydTodoMonthStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTodoMonthStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydTodoMonthStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydTodoMonthStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydTodoMonthStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydTodoMonthStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydTodoMonthStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydTodoMonthStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTodoMonthStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydTodoMonthStatisticalDataRequest) SetStatDate(v string) *QueryYydTodoMonthStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydTodoMonthStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydTodoMonthStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydTodoMonthStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTodoMonthStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydTodoMonthStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydTodoMonthStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydTodoMonthStatisticalDataResponseBody) SetMetaList(v []*QueryYydTodoMonthStatisticalDataResponseBodyMetaList) *QueryYydTodoMonthStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydTodoMonthStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydTodoMonthStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTodoMonthStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydTodoMonthStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydTodoMonthStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydTodoMonthStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydTodoMonthStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydTodoMonthStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydTodoMonthStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydTodoMonthStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydTodoMonthStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydTodoMonthStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydTodoMonthStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydTodoMonthStatisticalDataResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydTodoMonthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydTodoMonthStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTodoMonthStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydTodoMonthStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydTodoMonthStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydTodoMonthStatisticalDataResponse) SetStatusCode(v int32) *QueryYydTodoMonthStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydTodoMonthStatisticalDataResponse) SetBody(v *QueryYydTodoMonthStatisticalDataResponseBody) *QueryYydTodoMonthStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydTodoWeekStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydTodoWeekStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTodoWeekStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydTodoWeekStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydTodoWeekStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydTodoWeekStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydTodoWeekStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydTodoWeekStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydTodoWeekStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTodoWeekStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydTodoWeekStatisticalDataRequest) SetStatDate(v string) *QueryYydTodoWeekStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydTodoWeekStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                               `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydTodoWeekStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydTodoWeekStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTodoWeekStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydTodoWeekStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydTodoWeekStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydTodoWeekStatisticalDataResponseBody) SetMetaList(v []*QueryYydTodoWeekStatisticalDataResponseBodyMetaList) *QueryYydTodoWeekStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydTodoWeekStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydTodoWeekStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTodoWeekStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydTodoWeekStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydTodoWeekStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydTodoWeekStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydTodoWeekStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydTodoWeekStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydTodoWeekStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydTodoWeekStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydTodoWeekStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydTodoWeekStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydTodoWeekStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydTodoWeekStatisticalDataResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydTodoWeekStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydTodoWeekStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTodoWeekStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydTodoWeekStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydTodoWeekStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydTodoWeekStatisticalDataResponse) SetStatusCode(v int32) *QueryYydTodoWeekStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydTodoWeekStatisticalDataResponse) SetBody(v *QueryYydTodoWeekStatisticalDataResponseBody) *QueryYydTodoWeekStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydTotalDayStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydTotalDayStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTotalDayStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydTotalDayStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydTotalDayStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydTotalDayStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydTotalDayStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydTotalDayStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydTotalDayStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTotalDayStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydTotalDayStatisticalDataRequest) SetStatDate(v string) *QueryYydTotalDayStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydTotalDayStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                               `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydTotalDayStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydTotalDayStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTotalDayStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydTotalDayStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydTotalDayStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydTotalDayStatisticalDataResponseBody) SetMetaList(v []*QueryYydTotalDayStatisticalDataResponseBodyMetaList) *QueryYydTotalDayStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydTotalDayStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydTotalDayStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTotalDayStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydTotalDayStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydTotalDayStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydTotalDayStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydTotalDayStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydTotalDayStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydTotalDayStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydTotalDayStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydTotalDayStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydTotalDayStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydTotalDayStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydTotalDayStatisticalDataResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydTotalDayStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydTotalDayStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTotalDayStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydTotalDayStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydTotalDayStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydTotalDayStatisticalDataResponse) SetStatusCode(v int32) *QueryYydTotalDayStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydTotalDayStatisticalDataResponse) SetBody(v *QueryYydTotalDayStatisticalDataResponseBody) *QueryYydTotalDayStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydTotalMonthStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydTotalMonthStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTotalMonthStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydTotalMonthStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydTotalMonthStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydTotalMonthStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydTotalMonthStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydTotalMonthStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydTotalMonthStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTotalMonthStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydTotalMonthStatisticalDataRequest) SetStatDate(v string) *QueryYydTotalMonthStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydTotalMonthStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                 `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydTotalMonthStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydTotalMonthStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTotalMonthStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydTotalMonthStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydTotalMonthStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydTotalMonthStatisticalDataResponseBody) SetMetaList(v []*QueryYydTotalMonthStatisticalDataResponseBodyMetaList) *QueryYydTotalMonthStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydTotalMonthStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydTotalMonthStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTotalMonthStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydTotalMonthStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydTotalMonthStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydTotalMonthStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydTotalMonthStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydTotalMonthStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydTotalMonthStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydTotalMonthStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydTotalMonthStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydTotalMonthStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydTotalMonthStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydTotalMonthStatisticalDataResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydTotalMonthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydTotalMonthStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTotalMonthStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydTotalMonthStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydTotalMonthStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydTotalMonthStatisticalDataResponse) SetStatusCode(v int32) *QueryYydTotalMonthStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydTotalMonthStatisticalDataResponse) SetBody(v *QueryYydTotalMonthStatisticalDataResponseBody) *QueryYydTotalMonthStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydTotalStdStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydTotalStdStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTotalStdStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydTotalStdStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydTotalStdStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydTotalStdStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydTotalStdStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydTotalStdStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydTotalStdStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTotalStdStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydTotalStdStatisticalDataRequest) SetStatDate(v string) *QueryYydTotalStdStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydTotalStdStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                               `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydTotalStdStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydTotalStdStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTotalStdStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydTotalStdStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydTotalStdStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydTotalStdStatisticalDataResponseBody) SetMetaList(v []*QueryYydTotalStdStatisticalDataResponseBodyMetaList) *QueryYydTotalStdStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydTotalStdStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydTotalStdStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTotalStdStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydTotalStdStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydTotalStdStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydTotalStdStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydTotalStdStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydTotalStdStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydTotalStdStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydTotalStdStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydTotalStdStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydTotalStdStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydTotalStdStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydTotalStdStatisticalDataResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydTotalStdStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydTotalStdStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTotalStdStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydTotalStdStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydTotalStdStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydTotalStdStatisticalDataResponse) SetStatusCode(v int32) *QueryYydTotalStdStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydTotalStdStatisticalDataResponse) SetBody(v *QueryYydTotalStdStatisticalDataResponseBody) *QueryYydTotalStdStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryYydTotalWeekStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryYydTotalWeekStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTotalWeekStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryYydTotalWeekStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryYydTotalWeekStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryYydTotalWeekStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryYydTotalWeekStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryYydTotalWeekStatisticalDataRequest struct {
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryYydTotalWeekStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTotalWeekStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryYydTotalWeekStatisticalDataRequest) SetStatDate(v string) *QueryYydTotalWeekStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryYydTotalWeekStatisticalDataResponseBody struct {
	DataList []map[string]interface{}                                `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*QueryYydTotalWeekStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryYydTotalWeekStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTotalWeekStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYydTotalWeekStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryYydTotalWeekStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryYydTotalWeekStatisticalDataResponseBody) SetMetaList(v []*QueryYydTotalWeekStatisticalDataResponseBodyMetaList) *QueryYydTotalWeekStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryYydTotalWeekStatisticalDataResponseBodyMetaList struct {
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryYydTotalWeekStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTotalWeekStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryYydTotalWeekStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryYydTotalWeekStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryYydTotalWeekStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryYydTotalWeekStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryYydTotalWeekStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryYydTotalWeekStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryYydTotalWeekStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryYydTotalWeekStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *QueryYydTotalWeekStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryYydTotalWeekStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

type QueryYydTotalWeekStatisticalDataResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYydTotalWeekStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYydTotalWeekStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYydTotalWeekStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryYydTotalWeekStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryYydTotalWeekStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryYydTotalWeekStatisticalDataResponse) SetStatusCode(v int32) *QueryYydTotalWeekStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYydTotalWeekStatisticalDataResponse) SetBody(v *QueryYydTotalWeekStatisticalDataResponseBody) *QueryYydTotalWeekStatisticalDataResponse {
	s.Body = v
	return s
}

type SearchCompanyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchCompanyHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchCompanyHeaders) GoString() string {
	return s.String()
}

func (s *SearchCompanyHeaders) SetCommonHeaders(v map[string]*string) *SearchCompanyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchCompanyHeaders) SetXAcsDingtalkAccessToken(v string) *SearchCompanyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchCompanyRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey,omitempty"`
}

func (s SearchCompanyRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchCompanyRequest) GoString() string {
	return s.String()
}

func (s *SearchCompanyRequest) SetPageNumber(v int32) *SearchCompanyRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchCompanyRequest) SetPageSize(v int32) *SearchCompanyRequest {
	s.PageSize = &v
	return s
}

func (s *SearchCompanyRequest) SetSearchKey(v string) *SearchCompanyRequest {
	s.SearchKey = &v
	return s
}

type SearchCompanyResponseBody struct {
	// example:
	//
	// [     {       "ENT_NAME": "xx",       "SOCIAL_CREDIT_CODE": "xx",       "LICENSE_NUMBER": "xx",       "REG_CAP": "10000000.0",       "ES_DATE": "2006-12-07",       "LEGAL_NAME": "xx",       "ORG_NO": "xx",       "TAX_NUM": "xx",       "ENT_STATUS": "在营"     }   ]
	Data  *string `json:"data,omitempty" xml:"data,omitempty"`
	Total *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s SearchCompanyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchCompanyResponseBody) GoString() string {
	return s.String()
}

func (s *SearchCompanyResponseBody) SetData(v string) *SearchCompanyResponseBody {
	s.Data = &v
	return s
}

func (s *SearchCompanyResponseBody) SetTotal(v int64) *SearchCompanyResponseBody {
	s.Total = &v
	return s
}

type SearchCompanyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchCompanyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchCompanyResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchCompanyResponse) GoString() string {
	return s.String()
}

func (s *SearchCompanyResponse) SetHeaders(v map[string]*string) *SearchCompanyResponse {
	s.Headers = v
	return s
}

func (s *SearchCompanyResponse) SetStatusCode(v int32) *SearchCompanyResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchCompanyResponse) SetBody(v *SearchCompanyResponseBody) *SearchCompanyResponse {
	s.Body = v
	return s
}

type SyncDataScreenHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SyncDataScreenHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncDataScreenHeaders) GoString() string {
	return s.String()
}

func (s *SyncDataScreenHeaders) SetCommonHeaders(v map[string]*string) *SyncDataScreenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncDataScreenHeaders) SetXAcsDingtalkAccessToken(v string) *SyncDataScreenHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SyncDataScreenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	ScreenId *string `json:"screenId,omitempty" xml:"screenId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ABC
	Ticket *string `json:"ticket,omitempty" xml:"ticket,omitempty"`
}

func (s SyncDataScreenRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncDataScreenRequest) GoString() string {
	return s.String()
}

func (s *SyncDataScreenRequest) SetScreenId(v string) *SyncDataScreenRequest {
	s.ScreenId = &v
	return s
}

func (s *SyncDataScreenRequest) SetTicket(v string) *SyncDataScreenRequest {
	s.Ticket = &v
	return s
}

type SyncDataScreenResponseBody struct {
	Result  *bool `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SyncDataScreenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncDataScreenResponseBody) GoString() string {
	return s.String()
}

func (s *SyncDataScreenResponseBody) SetResult(v bool) *SyncDataScreenResponseBody {
	s.Result = &v
	return s
}

func (s *SyncDataScreenResponseBody) SetSuccess(v bool) *SyncDataScreenResponseBody {
	s.Success = &v
	return s
}

type SyncDataScreenResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncDataScreenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncDataScreenResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncDataScreenResponse) GoString() string {
	return s.String()
}

func (s *SyncDataScreenResponse) SetHeaders(v map[string]*string) *SyncDataScreenResponse {
	s.Headers = v
	return s
}

func (s *SyncDataScreenResponse) SetStatusCode(v int32) *SyncDataScreenResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncDataScreenResponse) SetBody(v *SyncDataScreenResponseBody) *SyncDataScreenResponse {
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
// 关闭数据投递任务
//
// @param request - CloseDataDeliverRequest
//
// @param headers - CloseDataDeliverHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseDataDeliverResponse
func (client *Client) CloseDataDeliverWithOptions(request *CloseDataDeliverRequest, headers *CloseDataDeliverHeaders, runtime *util.RuntimeOptions) (_result *CloseDataDeliverResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeliverId)) {
		query["deliverId"] = request.DeliverId
	}

	if !tea.BoolValue(util.IsUnset(request.DispatchingItemType)) {
		query["dispatchingItemType"] = request.DispatchingItemType
	}

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
		Action:      tea.String("CloseDataDeliver"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/dataDeliverServices/close"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseDataDeliverResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关闭数据投递任务
//
// @param request - CloseDataDeliverRequest
//
// @return CloseDataDeliverResponse
func (client *Client) CloseDataDeliver(request *CloseDataDeliverRequest) (_result *CloseDataDeliverResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CloseDataDeliverHeaders{}
	_result = &CloseDataDeliverResponse{}
	_body, _err := client.CloseDataDeliverWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建数据投递
//
// @param request - CreateDataDeliverRequest
//
// @param headers - CreateDataDeliverHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataDeliverResponse
func (client *Client) CreateDataDeliverWithOptions(request *CreateDataDeliverRequest, headers *CreateDataDeliverHeaders, runtime *util.RuntimeOptions) (_result *CreateDataDeliverResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bizcode)) {
		query["bizcode"] = request.Bizcode
	}

	if !tea.BoolValue(util.IsUnset(request.DispatchingCycle)) {
		query["dispatchingCycle"] = request.DispatchingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.DispatchingItemType)) {
		query["dispatchingItemType"] = request.DispatchingItemType
	}

	if !tea.BoolValue(util.IsUnset(request.DispatchingStartDate)) {
		query["dispatchingStartDate"] = request.DispatchingStartDate
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Param)) {
		body["param"] = request.Param
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
		Action:      tea.String("CreateDataDeliver"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/dataDeliveries"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDataDeliverResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据投递
//
// @param request - CreateDataDeliverRequest
//
// @return CreateDataDeliverResponse
func (client *Client) CreateDataDeliver(request *CreateDataDeliverRequest) (_result *CreateDataDeliverResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateDataDeliverHeaders{}
	_result = &CreateDataDeliverResponse{}
	_body, _err := client.CreateDataDeliverWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增数据大屏
//
// @param request - CreateScreenRequest
//
// @param headers - CreateScreenHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScreenResponse
func (client *Client) CreateScreenWithOptions(request *CreateScreenRequest, headers *CreateScreenHeaders, runtime *util.RuntimeOptions) (_result *CreateScreenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["templateId"] = request.TemplateId
	}

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
		Action:      tea.String("CreateScreen"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/screens"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateScreenResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增数据大屏
//
// @param request - CreateScreenRequest
//
// @return CreateScreenResponse
func (client *Client) CreateScreen(request *CreateScreenRequest) (_result *CreateScreenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateScreenHeaders{}
	_result = &CreateScreenResponse{}
	_body, _err := client.CreateScreenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 数据流通中心获取数据服务
//
// @param request - DataMarketIsvServiceRequest
//
// @param headers - DataMarketIsvServiceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DataMarketIsvServiceResponse
func (client *Client) DataMarketIsvServiceWithOptions(request *DataMarketIsvServiceRequest, headers *DataMarketIsvServiceHeaders, runtime *util.RuntimeOptions) (_result *DataMarketIsvServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["accountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		body["apiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.Args)) {
		body["args"] = request.Args
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
		Action:      tea.String("DataMarketIsvService"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/isv/dataMarketServices/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DataMarketIsvServiceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 数据流通中心获取数据服务
//
// @param request - DataMarketIsvServiceRequest
//
// @return DataMarketIsvServiceResponse
func (client *Client) DataMarketIsvService(request *DataMarketIsvServiceRequest) (_result *DataMarketIsvServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DataMarketIsvServiceHeaders{}
	_result = &DataMarketIsvServiceResponse{}
	_body, _err := client.DataMarketIsvServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 数据流通中心获取数据服务
//
// @param request - DataMarketServiceRequest
//
// @param headers - DataMarketServiceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DataMarketServiceResponse
func (client *Client) DataMarketServiceWithOptions(request *DataMarketServiceRequest, headers *DataMarketServiceHeaders, runtime *util.RuntimeOptions) (_result *DataMarketServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		body["apiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.Args)) {
		body["args"] = request.Args
	}

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
		Action:      tea.String("DataMarketService"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/dataMarketServices/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DataMarketServiceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 数据流通中心获取数据服务
//
// @param request - DataMarketServiceRequest
//
// @return DataMarketServiceResponse
func (client *Client) DataMarketService(request *DataMarketServiceRequest) (_result *DataMarketServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DataMarketServiceHeaders{}
	_result = &DataMarketServiceResponse{}
	_body, _err := client.DataMarketServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 工商-经营异常
//
// @param request - GetAbnormalOperationRequest
//
// @param headers - GetAbnormalOperationHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAbnormalOperationResponse
func (client *Client) GetAbnormalOperationWithOptions(request *GetAbnormalOperationRequest, headers *GetAbnormalOperationHeaders, runtime *util.RuntimeOptions) (_result *GetAbnormalOperationResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["searchKey"] = request.SearchKey
	}

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
		Action:      tea.String("GetAbnormalOperation"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/companies/abnormalOperations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAbnormalOperationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 工商-经营异常
//
// @param request - GetAbnormalOperationRequest
//
// @return GetAbnormalOperationResponse
func (client *Client) GetAbnormalOperation(request *GetAbnormalOperationRequest) (_result *GetAbnormalOperationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAbnormalOperationHeaders{}
	_result = &GetAbnormalOperationResponse{}
	_body, _err := client.GetAbnormalOperationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取工商-行政许可
//
// @param request - GetAdministrativeLicensingRequest
//
// @param headers - GetAdministrativeLicensingHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAdministrativeLicensingResponse
func (client *Client) GetAdministrativeLicensingWithOptions(request *GetAdministrativeLicensingRequest, headers *GetAdministrativeLicensingHeaders, runtime *util.RuntimeOptions) (_result *GetAdministrativeLicensingResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["searchKey"] = request.SearchKey
	}

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
		Action:      tea.String("GetAdministrativeLicensing"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/companies/administrativeLicenses"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAdministrativeLicensingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工商-行政许可
//
// @param request - GetAdministrativeLicensingRequest
//
// @return GetAdministrativeLicensingResponse
func (client *Client) GetAdministrativeLicensing(request *GetAdministrativeLicensingRequest) (_result *GetAdministrativeLicensingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAdministrativeLicensingHeaders{}
	_result = &GetAdministrativeLicensingResponse{}
	_body, _err := client.GetAdministrativeLicensingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 负面-行政处罚
//
// @param request - GetAdministrativePenaltiesRequest
//
// @param headers - GetAdministrativePenaltiesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAdministrativePenaltiesResponse
func (client *Client) GetAdministrativePenaltiesWithOptions(request *GetAdministrativePenaltiesRequest, headers *GetAdministrativePenaltiesHeaders, runtime *util.RuntimeOptions) (_result *GetAdministrativePenaltiesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["searchKey"] = request.SearchKey
	}

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
		Action:      tea.String("GetAdministrativePenalties"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/companies/administrativePenalties"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAdministrativePenaltiesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 负面-行政处罚
//
// @param request - GetAdministrativePenaltiesRequest
//
// @return GetAdministrativePenaltiesResponse
func (client *Client) GetAdministrativePenalties(request *GetAdministrativePenaltiesRequest) (_result *GetAdministrativePenaltiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAdministrativePenaltiesHeaders{}
	_result = &GetAdministrativePenaltiesResponse{}
	_body, _err := client.GetAdministrativePenaltiesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 工商-基础信息
//
// @param request - GetBasicInfoRequest
//
// @param headers - GetBasicInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBasicInfoResponse
func (client *Client) GetBasicInfoWithOptions(request *GetBasicInfoRequest, headers *GetBasicInfoHeaders, runtime *util.RuntimeOptions) (_result *GetBasicInfoResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["searchKey"] = request.SearchKey
	}

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
		Action:      tea.String("GetBasicInfo"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/companies/businessBasicInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBasicInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 工商-基础信息
//
// @param request - GetBasicInfoRequest
//
// @return GetBasicInfoResponse
func (client *Client) GetBasicInfo(request *GetBasicInfoRequest) (_result *GetBasicInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetBasicInfoHeaders{}
	_result = &GetBasicInfoResponse{}
	_body, _err := client.GetBasicInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取经营-招投标信息
//
// @param request - GetBiddingInfoRequest
//
// @param headers - GetBiddingInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBiddingInfoResponse
func (client *Client) GetBiddingInfoWithOptions(request *GetBiddingInfoRequest, headers *GetBiddingInfoHeaders, runtime *util.RuntimeOptions) (_result *GetBiddingInfoResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["searchKey"] = request.SearchKey
	}

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
		Action:      tea.String("GetBiddingInfo"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/companies/biddingInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBiddingInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取经营-招投标信息
//
// @param request - GetBiddingInfoRequest
//
// @return GetBiddingInfoResponse
func (client *Client) GetBiddingInfo(request *GetBiddingInfoRequest) (_result *GetBiddingInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetBiddingInfoHeaders{}
	_result = &GetBiddingInfoResponse{}
	_body, _err := client.GetBiddingInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取工商-分支机构
//
// @param request - GetBranchInfoRequest
//
// @param headers - GetBranchInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBranchInfoResponse
func (client *Client) GetBranchInfoWithOptions(request *GetBranchInfoRequest, headers *GetBranchInfoHeaders, runtime *util.RuntimeOptions) (_result *GetBranchInfoResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["searchKey"] = request.SearchKey
	}

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
		Action:      tea.String("GetBranchInfo"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/companies/branchInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBranchInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工商-分支机构
//
// @param request - GetBranchInfoRequest
//
// @return GetBranchInfoResponse
func (client *Client) GetBranchInfo(request *GetBranchInfoRequest) (_result *GetBranchInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetBranchInfoHeaders{}
	_result = &GetBranchInfoResponse{}
	_body, _err := client.GetBranchInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取工商-变更记录
//
// @param request - GetChangeRecordRequest
//
// @param headers - GetChangeRecordHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChangeRecordResponse
func (client *Client) GetChangeRecordWithOptions(request *GetChangeRecordRequest, headers *GetChangeRecordHeaders, runtime *util.RuntimeOptions) (_result *GetChangeRecordResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["searchKey"] = request.SearchKey
	}

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
		Action:      tea.String("GetChangeRecord"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/companies/changeRecords"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetChangeRecordResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工商-变更记录
//
// @param request - GetChangeRecordRequest
//
// @return GetChangeRecordResponse
func (client *Client) GetChangeRecord(request *GetChangeRecordRequest) (_result *GetChangeRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetChangeRecordHeaders{}
	_result = &GetChangeRecordResponse{}
	_body, _err := client.GetChangeRecordWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取投递任务信息
//
// @param request - GetDataDeliverRequest
//
// @param headers - GetDataDeliverHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataDeliverResponse
func (client *Client) GetDataDeliverWithOptions(request *GetDataDeliverRequest, headers *GetDataDeliverHeaders, runtime *util.RuntimeOptions) (_result *GetDataDeliverResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeliverId)) {
		query["deliverId"] = request.DeliverId
	}

	if !tea.BoolValue(util.IsUnset(request.DispatchingItemType)) {
		query["dispatchingItemType"] = request.DispatchingItemType
	}

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
		Action:      tea.String("GetDataDeliver"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/dataDeliverServices/infos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDataDeliverResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取投递任务信息
//
// @param request - GetDataDeliverRequest
//
// @return GetDataDeliverResponse
func (client *Client) GetDataDeliver(request *GetDataDeliverRequest) (_result *GetDataDeliverResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDataDeliverHeaders{}
	_result = &GetDataDeliverResponse{}
	_body, _err := client.GetDataDeliverWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取知识产权-域名信息
//
// @param request - GetDomainInfoRequest
//
// @param headers - GetDomainInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDomainInfoResponse
func (client *Client) GetDomainInfoWithOptions(request *GetDomainInfoRequest, headers *GetDomainInfoHeaders, runtime *util.RuntimeOptions) (_result *GetDomainInfoResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["searchKey"] = request.SearchKey
	}

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
		Action:      tea.String("GetDomainInfo"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/companies/domainInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDomainInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取知识产权-域名信息
//
// @param request - GetDomainInfoRequest
//
// @return GetDomainInfoResponse
func (client *Client) GetDomainInfo(request *GetDomainInfoRequest) (_result *GetDomainInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDomainInfoHeaders{}
	_result = &GetDomainInfoResponse{}
	_body, _err := client.GetDomainInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取工商-双随机抽查结果
//
// @param request - GetDoubleRandomRequest
//
// @param headers - GetDoubleRandomHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDoubleRandomResponse
func (client *Client) GetDoubleRandomWithOptions(request *GetDoubleRandomRequest, headers *GetDoubleRandomHeaders, runtime *util.RuntimeOptions) (_result *GetDoubleRandomResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["searchKey"] = request.SearchKey
	}

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
		Action:      tea.String("GetDoubleRandom"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/companies/doubleRandomness"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDoubleRandomResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工商-双随机抽查结果
//
// @param request - GetDoubleRandomRequest
//
// @return GetDoubleRandomResponse
func (client *Client) GetDoubleRandom(request *GetDoubleRandomRequest) (_result *GetDoubleRandomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDoubleRandomHeaders{}
	_result = &GetDoubleRandomResponse{}
	_body, _err := client.GetDoubleRandomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 负面-环保处罚
//
// @param request - GetEnvironmentalPenaltiesRequest
//
// @param headers - GetEnvironmentalPenaltiesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEnvironmentalPenaltiesResponse
func (client *Client) GetEnvironmentalPenaltiesWithOptions(request *GetEnvironmentalPenaltiesRequest, headers *GetEnvironmentalPenaltiesHeaders, runtime *util.RuntimeOptions) (_result *GetEnvironmentalPenaltiesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["searchKey"] = request.SearchKey
	}

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
		Action:      tea.String("GetEnvironmentalPenalties"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/companies/environmentalPenalties"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEnvironmentalPenaltiesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 负面-环保处罚
//
// @param request - GetEnvironmentalPenaltiesRequest
//
// @return GetEnvironmentalPenaltiesResponse
func (client *Client) GetEnvironmentalPenalties(request *GetEnvironmentalPenaltiesRequest) (_result *GetEnvironmentalPenaltiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetEnvironmentalPenaltiesHeaders{}
	_result = &GetEnvironmentalPenaltiesResponse{}
	_body, _err := client.GetEnvironmentalPenaltiesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取事件订阅的数据
//
// @param request - GetEventDataRequest
//
// @param headers - GetEventDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEventDataResponse
func (client *Client) GetEventDataWithOptions(request *GetEventDataRequest, headers *GetEventDataHeaders, runtime *util.RuntimeOptions) (_result *GetEventDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.EventUid)) {
		body["eventUid"] = request.EventUid
	}

	if !tea.BoolValue(util.IsUnset(request.SubId)) {
		body["subId"] = request.SubId
	}

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
		Action:      tea.String("GetEventData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/eventDatas/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEventDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取事件订阅的数据
//
// @param request - GetEventDataRequest
//
// @return GetEventDataResponse
func (client *Client) GetEventData(request *GetEventDataRequest) (_result *GetEventDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetEventDataHeaders{}
	_result = &GetEventDataResponse{}
	_body, _err := client.GetEventDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 工商-股东信息
//
// @param request - GetHolderInfoRequest
//
// @param headers - GetHolderInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHolderInfoResponse
func (client *Client) GetHolderInfoWithOptions(request *GetHolderInfoRequest, headers *GetHolderInfoHeaders, runtime *util.RuntimeOptions) (_result *GetHolderInfoResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["searchKey"] = request.SearchKey
	}

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
		Action:      tea.String("GetHolderInfo"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/companies/shareholderInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHolderInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 工商-股东信息
//
// @param request - GetHolderInfoRequest
//
// @return GetHolderInfoResponse
func (client *Client) GetHolderInfo(request *GetHolderInfoRequest) (_result *GetHolderInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetHolderInfoHeaders{}
	_result = &GetHolderInfoResponse{}
	_body, _err := client.GetHolderInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取工商-知识产权出质
//
// @param request - GetIntellectualPropertyRequest
//
// @param headers - GetIntellectualPropertyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIntellectualPropertyResponse
func (client *Client) GetIntellectualPropertyWithOptions(request *GetIntellectualPropertyRequest, headers *GetIntellectualPropertyHeaders, runtime *util.RuntimeOptions) (_result *GetIntellectualPropertyResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["searchKey"] = request.SearchKey
	}

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
		Action:      tea.String("GetIntellectualProperty"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/companies/intellectualProperties"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIntellectualPropertyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工商-知识产权出质
//
// @param request - GetIntellectualPropertyRequest
//
// @return GetIntellectualPropertyResponse
func (client *Client) GetIntellectualProperty(request *GetIntellectualPropertyRequest) (_result *GetIntellectualPropertyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetIntellectualPropertyHeaders{}
	_result = &GetIntellectualPropertyResponse{}
	_body, _err := client.GetIntellectualPropertyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取工商-对外投资
//
// @param request - GetInvestmentAbroadRequest
//
// @param headers - GetInvestmentAbroadHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInvestmentAbroadResponse
func (client *Client) GetInvestmentAbroadWithOptions(request *GetInvestmentAbroadRequest, headers *GetInvestmentAbroadHeaders, runtime *util.RuntimeOptions) (_result *GetInvestmentAbroadResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["searchKey"] = request.SearchKey
	}

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
		Action:      tea.String("GetInvestmentAbroad"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/companies/abroadInvestments"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInvestmentAbroadResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工商-对外投资
//
// @param request - GetInvestmentAbroadRequest
//
// @return GetInvestmentAbroadResponse
func (client *Client) GetInvestmentAbroad(request *GetInvestmentAbroadRequest) (_result *GetInvestmentAbroadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetInvestmentAbroadHeaders{}
	_result = &GetInvestmentAbroadResponse{}
	_body, _err := client.GetInvestmentAbroadWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取经营-招聘信息
//
// @param request - GetJobInfoRequest
//
// @param headers - GetJobInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobInfoResponse
func (client *Client) GetJobInfoWithOptions(request *GetJobInfoRequest, headers *GetJobInfoHeaders, runtime *util.RuntimeOptions) (_result *GetJobInfoResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["searchKey"] = request.SearchKey
	}

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
		Action:      tea.String("GetJobInfo"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/companies/jobInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取经营-招聘信息
//
// @param request - GetJobInfoRequest
//
// @return GetJobInfoResponse
func (client *Client) GetJobInfo(request *GetJobInfoRequest) (_result *GetJobInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetJobInfoHeaders{}
	_result = &GetJobInfoResponse{}
	_body, _err := client.GetJobInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取知识产权-专利信息
//
// @param request - GetPatentInfoRequest
//
// @param headers - GetPatentInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPatentInfoResponse
func (client *Client) GetPatentInfoWithOptions(request *GetPatentInfoRequest, headers *GetPatentInfoHeaders, runtime *util.RuntimeOptions) (_result *GetPatentInfoResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["searchKey"] = request.SearchKey
	}

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
		Action:      tea.String("GetPatentInfo"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/companies/patentInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPatentInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取知识产权-专利信息
//
// @param request - GetPatentInfoRequest
//
// @return GetPatentInfoResponse
func (client *Client) GetPatentInfo(request *GetPatentInfoRequest) (_result *GetPatentInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPatentInfoHeaders{}
	_result = &GetPatentInfoResponse{}
	_body, _err := client.GetPatentInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取工商-主要人员
//
// @param request - GetPrincipalEmployeeRequest
//
// @param headers - GetPrincipalEmployeeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrincipalEmployeeResponse
func (client *Client) GetPrincipalEmployeeWithOptions(request *GetPrincipalEmployeeRequest, headers *GetPrincipalEmployeeHeaders, runtime *util.RuntimeOptions) (_result *GetPrincipalEmployeeResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["searchKey"] = request.SearchKey
	}

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
		Action:      tea.String("GetPrincipalEmployee"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/companies/principalEmployees"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPrincipalEmployeeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工商-主要人员
//
// @param request - GetPrincipalEmployeeRequest
//
// @return GetPrincipalEmployeeResponse
func (client *Client) GetPrincipalEmployee(request *GetPrincipalEmployeeRequest) (_result *GetPrincipalEmployeeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPrincipalEmployeeHeaders{}
	_result = &GetPrincipalEmployeeResponse{}
	_body, _err := client.GetPrincipalEmployeeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 经营-一般纳税人
//
// @param request - GetQeneralTaxpayerInfoRequest
//
// @param headers - GetQeneralTaxpayerInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQeneralTaxpayerInfoResponse
func (client *Client) GetQeneralTaxpayerInfoWithOptions(request *GetQeneralTaxpayerInfoRequest, headers *GetQeneralTaxpayerInfoHeaders, runtime *util.RuntimeOptions) (_result *GetQeneralTaxpayerInfoResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["searchKey"] = request.SearchKey
	}

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
		Action:      tea.String("GetQeneralTaxpayerInfo"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/companies/generalTaxpayerInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQeneralTaxpayerInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 经营-一般纳税人
//
// @param request - GetQeneralTaxpayerInfoRequest
//
// @return GetQeneralTaxpayerInfoResponse
func (client *Client) GetQeneralTaxpayerInfo(request *GetQeneralTaxpayerInfoRequest) (_result *GetQeneralTaxpayerInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetQeneralTaxpayerInfoHeaders{}
	_result = &GetQeneralTaxpayerInfoResponse{}
	_body, _err := client.GetQeneralTaxpayerInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取知识产权-资质证书
//
// @param request - GetQualificationCertRequest
//
// @param headers - GetQualificationCertHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualificationCertResponse
func (client *Client) GetQualificationCertWithOptions(request *GetQualificationCertRequest, headers *GetQualificationCertHeaders, runtime *util.RuntimeOptions) (_result *GetQualificationCertResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["searchKey"] = request.SearchKey
	}

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
		Action:      tea.String("GetQualificationCert"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/companies/qualificationCerts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQualificationCertResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取知识产权-资质证书
//
// @param request - GetQualificationCertRequest
//
// @return GetQualificationCertResponse
func (client *Client) GetQualificationCert(request *GetQualificationCertRequest) (_result *GetQualificationCertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetQualificationCertHeaders{}
	_result = &GetQualificationCertResponse{}
	_body, _err := client.GetQualificationCertWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 负面-严重违法
//
// @param request - GetSeriousViolationRequest
//
// @param headers - GetSeriousViolationHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSeriousViolationResponse
func (client *Client) GetSeriousViolationWithOptions(request *GetSeriousViolationRequest, headers *GetSeriousViolationHeaders, runtime *util.RuntimeOptions) (_result *GetSeriousViolationResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["searchKey"] = request.SearchKey
	}

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
		Action:      tea.String("GetSeriousViolation"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/companies/seriousViolations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSeriousViolationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 负面-严重违法
//
// @param request - GetSeriousViolationRequest
//
// @return GetSeriousViolationResponse
func (client *Client) GetSeriousViolation(request *GetSeriousViolationRequest) (_result *GetSeriousViolationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSeriousViolationHeaders{}
	_result = &GetSeriousViolationResponse{}
	_body, _err := client.GetSeriousViolationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取知识产权-软件著作权
//
// @param request - GetSoftwareCopyrightRequest
//
// @param headers - GetSoftwareCopyrightHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSoftwareCopyrightResponse
func (client *Client) GetSoftwareCopyrightWithOptions(request *GetSoftwareCopyrightRequest, headers *GetSoftwareCopyrightHeaders, runtime *util.RuntimeOptions) (_result *GetSoftwareCopyrightResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["searchKey"] = request.SearchKey
	}

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
		Action:      tea.String("GetSoftwareCopyright"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/companies/softwareCopyrights"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSoftwareCopyrightResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取知识产权-软件著作权
//
// @param request - GetSoftwareCopyrightRequest
//
// @return GetSoftwareCopyrightResponse
func (client *Client) GetSoftwareCopyright(request *GetSoftwareCopyrightRequest) (_result *GetSoftwareCopyrightResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSoftwareCopyrightHeaders{}
	_result = &GetSoftwareCopyrightResponse{}
	_body, _err := client.GetSoftwareCopyrightWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取知识产权-商标信息
//
// @param request - GetTrademarkInfoRequest
//
// @param headers - GetTrademarkInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrademarkInfoResponse
func (client *Client) GetTrademarkInfoWithOptions(request *GetTrademarkInfoRequest, headers *GetTrademarkInfoHeaders, runtime *util.RuntimeOptions) (_result *GetTrademarkInfoResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["searchKey"] = request.SearchKey
	}

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
		Action:      tea.String("GetTrademarkInfo"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/companies/trademarkInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTrademarkInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取知识产权-商标信息
//
// @param request - GetTrademarkInfoRequest
//
// @return GetTrademarkInfoResponse
func (client *Client) GetTrademarkInfo(request *GetTrademarkInfoRequest) (_result *GetTrademarkInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTrademarkInfoHeaders{}
	_result = &GetTrademarkInfoResponse{}
	_body, _err := client.GetTrademarkInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取知识产权-作品著作权
//
// @param request - GetWorkCopyrightRequest
//
// @param headers - GetWorkCopyrightHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkCopyrightResponse
func (client *Client) GetWorkCopyrightWithOptions(request *GetWorkCopyrightRequest, headers *GetWorkCopyrightHeaders, runtime *util.RuntimeOptions) (_result *GetWorkCopyrightResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["searchKey"] = request.SearchKey
	}

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
		Action:      tea.String("GetWorkCopyright"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/companies/workCopyrights"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWorkCopyrightResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取知识产权-作品著作权
//
// @param request - GetWorkCopyrightRequest
//
// @return GetWorkCopyrightResponse
func (client *Client) GetWorkCopyright(request *GetWorkCopyrightRequest) (_result *GetWorkCopyrightResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetWorkCopyrightHeaders{}
	_result = &GetWorkCopyrightResponse{}
	_body, _err := client.GetWorkCopyrightWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 数据投递列表
//
// @param request - ListDataDeliversRequest
//
// @param headers - ListDataDeliversHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataDeliversResponse
func (client *Client) ListDataDeliversWithOptions(request *ListDataDeliversRequest, headers *ListDataDeliversHeaders, runtime *util.RuntimeOptions) (_result *ListDataDeliversResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DispatchingItemType)) {
		query["dispatchingItemType"] = request.DispatchingItemType
	}

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
		Action:      tea.String("ListDataDelivers"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/dataDeliverServices/lists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataDeliversResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 数据投递列表
//
// @param request - ListDataDeliversRequest
//
// @return ListDataDeliversResponse
func (client *Client) ListDataDelivers(request *ListDataDeliversRequest) (_result *ListDataDeliversResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListDataDeliversHeaders{}
	_result = &ListDataDeliversResponse{}
	_body, _err := client.ListDataDeliversWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 操作表格配置
//
// @param request - OperateChartConfigRequest
//
// @param headers - OperateChartConfigHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateChartConfigResponse
func (client *Client) OperateChartConfigWithOptions(request *OperateChartConfigRequest, headers *OperateChartConfigHeaders, runtime *util.RuntimeOptions) (_result *OperateChartConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiKey)) {
		body["apiKey"] = request.ApiKey
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Param)) {
		body["param"] = request.Param
	}

	if !tea.BoolValue(util.IsUnset(request.Ticket)) {
		body["ticket"] = request.Ticket
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
		Action:      tea.String("OperateChartConfig"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/chartConfigs/operate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &OperateChartConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 操作表格配置
//
// @param request - OperateChartConfigRequest
//
// @return OperateChartConfigResponse
func (client *Client) OperateChartConfig(request *OperateChartConfigRequest) (_result *OperateChartConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OperateChartConfigHeaders{}
	_result = &OperateChartConfigResponse{}
	_body, _err := client.OperateChartConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 企业授权信息
//
// @param headers - PostCorpAuthInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PostCorpAuthInfoResponse
func (client *Client) PostCorpAuthInfoWithOptions(headers *PostCorpAuthInfoHeaders, runtime *util.RuntimeOptions) (_result *PostCorpAuthInfoResponse, _err error) {
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
		Action:      tea.String("PostCorpAuthInfo"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/corporations/authorize"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PostCorpAuthInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 企业授权信息
//
// @return PostCorpAuthInfoResponse
func (client *Client) PostCorpAuthInfo() (_result *PostCorpAuthInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PostCorpAuthInfoHeaders{}
	_result = &PostCorpAuthInfoResponse{}
	_body, _err := client.PostCorpAuthInfoWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业用户激活状态统计数据
//
// @param request - QueryActiveUserStatisticalDataRequest
//
// @param headers - QueryActiveUserStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryActiveUserStatisticalDataResponse
func (client *Client) QueryActiveUserStatisticalDataWithOptions(request *QueryActiveUserStatisticalDataRequest, headers *QueryActiveUserStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryActiveUserStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryActiveUserStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/activeUserData"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryActiveUserStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业用户激活状态统计数据
//
// @param request - QueryActiveUserStatisticalDataRequest
//
// @return QueryActiveUserStatisticalDataResponse
func (client *Client) QueryActiveUserStatisticalData(request *QueryActiveUserStatisticalDataRequest) (_result *QueryActiveUserStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryActiveUserStatisticalDataHeaders{}
	_result = &QueryActiveUserStatisticalDataResponse{}
	_body, _err := client.QueryActiveUserStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取安恒密盾统计数据
//
// @param request - QueryAnhmdStatisticalDataRequest
//
// @param headers - QueryAnhmdStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAnhmdStatisticalDataResponse
func (client *Client) QueryAnhmdStatisticalDataWithOptions(request *QueryAnhmdStatisticalDataRequest, headers *QueryAnhmdStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryAnhmdStatisticalDataResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryAnhmdStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/statisticDatas/anHmd"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAnhmdStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取安恒密盾统计数据
//
// @param request - QueryAnhmdStatisticalDataRequest
//
// @return QueryAnhmdStatisticalDataResponse
func (client *Client) QueryAnhmdStatisticalData(request *QueryAnhmdStatisticalDataRequest) (_result *QueryAnhmdStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAnhmdStatisticalDataHeaders{}
	_result = &QueryAnhmdStatisticalDataResponse{}
	_body, _err := client.QueryAnhmdStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业审批统计数据
//
// @param request - QueryApprovalStatisticalDataRequest
//
// @param headers - QueryApprovalStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryApprovalStatisticalDataResponse
func (client *Client) QueryApprovalStatisticalDataWithOptions(request *QueryApprovalStatisticalDataRequest, headers *QueryApprovalStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryApprovalStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryApprovalStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/approvalData"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryApprovalStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业审批统计数据
//
// @param request - QueryApprovalStatisticalDataRequest
//
// @return QueryApprovalStatisticalDataResponse
func (client *Client) QueryApprovalStatisticalData(request *QueryApprovalStatisticalDataRequest) (_result *QueryApprovalStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryApprovalStatisticalDataHeaders{}
	_result = &QueryApprovalStatisticalDataResponse{}
	_body, _err := client.QueryApprovalStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业考勤统计数据
//
// @param request - QueryAttendanceStatisticalDataRequest
//
// @param headers - QueryAttendanceStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAttendanceStatisticalDataResponse
func (client *Client) QueryAttendanceStatisticalDataWithOptions(request *QueryAttendanceStatisticalDataRequest, headers *QueryAttendanceStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryAttendanceStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryAttendanceStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/attendanceData"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAttendanceStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业考勤统计数据
//
// @param request - QueryAttendanceStatisticalDataRequest
//
// @return QueryAttendanceStatisticalDataResponse
func (client *Client) QueryAttendanceStatisticalData(request *QueryAttendanceStatisticalDataRequest) (_result *QueryAttendanceStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAttendanceStatisticalDataHeaders{}
	_result = &QueryAttendanceStatisticalDataResponse{}
	_body, _err := client.QueryAttendanceStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业公告统计数据
//
// @param request - QueryBlackboardStatisticalDataRequest
//
// @param headers - QueryBlackboardStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBlackboardStatisticalDataResponse
func (client *Client) QueryBlackboardStatisticalDataWithOptions(request *QueryBlackboardStatisticalDataRequest, headers *QueryBlackboardStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryBlackboardStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryBlackboardStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/blackboardData"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBlackboardStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业公告统计数据
//
// @param request - QueryBlackboardStatisticalDataRequest
//
// @return QueryBlackboardStatisticalDataResponse
func (client *Client) QueryBlackboardStatisticalData(request *QueryBlackboardStatisticalDataRequest) (_result *QueryBlackboardStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryBlackboardStatisticalDataHeaders{}
	_result = &QueryBlackboardStatisticalDataResponse{}
	_body, _err := client.QueryBlackboardStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业日程统计数据
//
// @param request - QueryCalendarStatisticalDataRequest
//
// @param headers - QueryCalendarStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCalendarStatisticalDataResponse
func (client *Client) QueryCalendarStatisticalDataWithOptions(request *QueryCalendarStatisticalDataRequest, headers *QueryCalendarStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryCalendarStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryCalendarStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/calendarData"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCalendarStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业日程统计数据
//
// @param request - QueryCalendarStatisticalDataRequest
//
// @return QueryCalendarStatisticalDataResponse
func (client *Client) QueryCalendarStatisticalData(request *QueryCalendarStatisticalDataRequest) (_result *QueryCalendarStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCalendarStatisticalDataHeaders{}
	_result = &QueryCalendarStatisticalDataResponse{}
	_body, _err := client.QueryCalendarStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取图表数据
//
// @param request - QueryChartDataRequest
//
// @param headers - QueryChartDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryChartDataResponse
func (client *Client) QueryChartDataWithOptions(request *QueryChartDataRequest, headers *QueryChartDataHeaders, runtime *util.RuntimeOptions) (_result *QueryChartDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.Ticket)) {
		body["ticket"] = request.Ticket
	}

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
		Action:      tea.String("QueryChartData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/chartDatas/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryChartDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取图表数据
//
// @param request - QueryChartDataRequest
//
// @return QueryChartDataResponse
func (client *Client) QueryChartData(request *QueryChartDataRequest) (_result *QueryChartDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryChartDataHeaders{}
	_result = &QueryChartDataResponse{}
	_body, _err := client.QueryChartDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业签到统计数据
//
// @param request - QueryCheckinStatisticalDataRequest
//
// @param headers - QueryCheckinStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCheckinStatisticalDataResponse
func (client *Client) QueryCheckinStatisticalDataWithOptions(request *QueryCheckinStatisticalDataRequest, headers *QueryCheckinStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryCheckinStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryCheckinStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/checkinData"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCheckinStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业签到统计数据
//
// @param request - QueryCheckinStatisticalDataRequest
//
// @return QueryCheckinStatisticalDataResponse
func (client *Client) QueryCheckinStatisticalData(request *QueryCheckinStatisticalDataRequest) (_result *QueryCheckinStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCheckinStatisticalDataHeaders{}
	_result = &QueryCheckinStatisticalDataResponse{}
	_body, _err := client.QueryCheckinStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业全员圈统计数据
//
// @param request - QueryCircleStatisticalDataRequest
//
// @param headers - QueryCircleStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCircleStatisticalDataResponse
func (client *Client) QueryCircleStatisticalDataWithOptions(request *QueryCircleStatisticalDataRequest, headers *QueryCircleStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryCircleStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryCircleStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/circleData"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCircleStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业全员圈统计数据
//
// @param request - QueryCircleStatisticalDataRequest
//
// @return QueryCircleStatisticalDataResponse
func (client *Client) QueryCircleStatisticalData(request *QueryCircleStatisticalDataRequest) (_result *QueryCircleStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCircleStatisticalDataHeaders{}
	_result = &QueryCircleStatisticalDataResponse{}
	_body, _err := client.QueryCircleStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过企业名称/社会统一信用代码/工商注册号，查询企业的基本画像信息。
//
// @param request - QueryCompanyBasicInfoRequest
//
// @param headers - QueryCompanyBasicInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCompanyBasicInfoResponse
func (client *Client) QueryCompanyBasicInfoWithOptions(request *QueryCompanyBasicInfoRequest, headers *QueryCompanyBasicInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryCompanyBasicInfoResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("QueryCompanyBasicInfo"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/companies/basicInfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCompanyBasicInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过企业名称/社会统一信用代码/工商注册号，查询企业的基本画像信息。
//
// @param request - QueryCompanyBasicInfoRequest
//
// @return QueryCompanyBasicInfoResponse
func (client *Client) QueryCompanyBasicInfo(request *QueryCompanyBasicInfoRequest) (_result *QueryCompanyBasicInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCompanyBasicInfoHeaders{}
	_result = &QueryCompanyBasicInfoResponse{}
	_body, _err := client.QueryCompanyBasicInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数字区县组织信息
//
// @param request - QueryDigitalDistrictOrgInfoRequest
//
// @param headers - QueryDigitalDistrictOrgInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDigitalDistrictOrgInfoResponse
func (client *Client) QueryDigitalDistrictOrgInfoWithOptions(request *QueryDigitalDistrictOrgInfoRequest, headers *QueryDigitalDistrictOrgInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryDigitalDistrictOrgInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpIds)) {
		body["corpIds"] = request.CorpIds
	}

	if !tea.BoolValue(util.IsUnset(request.StatDates)) {
		body["statDates"] = request.StatDates
	}

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
		Action:      tea.String("QueryDigitalDistrictOrgInfo"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/digitalCounty/orgInfos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDigitalDistrictOrgInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数字区县组织信息
//
// @param request - QueryDigitalDistrictOrgInfoRequest
//
// @return QueryDigitalDistrictOrgInfoResponse
func (client *Client) QueryDigitalDistrictOrgInfo(request *QueryDigitalDistrictOrgInfoRequest) (_result *QueryDigitalDistrictOrgInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDigitalDistrictOrgInfoHeaders{}
	_result = &QueryDigitalDistrictOrgInfoResponse{}
	_body, _err := client.QueryDigitalDistrictOrgInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业DING接收及评论统计数据
//
// @param request - QueryDingReciveStatisticalDataRequest
//
// @param headers - QueryDingReciveStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDingReciveStatisticalDataResponse
func (client *Client) QueryDingReciveStatisticalDataWithOptions(request *QueryDingReciveStatisticalDataRequest, headers *QueryDingReciveStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryDingReciveStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryDingReciveStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/dingReciveData"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDingReciveStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业DING接收及评论统计数据
//
// @param request - QueryDingReciveStatisticalDataRequest
//
// @return QueryDingReciveStatisticalDataResponse
func (client *Client) QueryDingReciveStatisticalData(request *QueryDingReciveStatisticalDataRequest) (_result *QueryDingReciveStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDingReciveStatisticalDataHeaders{}
	_result = &QueryDingReciveStatisticalDataResponse{}
	_body, _err := client.QueryDingReciveStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业DING发送统计数据
//
// @param request - QueryDingSendStatisticalDataRequest
//
// @param headers - QueryDingSendStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDingSendStatisticalDataResponse
func (client *Client) QueryDingSendStatisticalDataWithOptions(request *QueryDingSendStatisticalDataRequest, headers *QueryDingSendStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryDingSendStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryDingSendStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/dingSendData"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDingSendStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业DING发送统计数据
//
// @param request - QueryDingSendStatisticalDataRequest
//
// @return QueryDingSendStatisticalDataResponse
func (client *Client) QueryDingSendStatisticalData(request *QueryDingSendStatisticalDataRequest) (_result *QueryDingSendStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDingSendStatisticalDataHeaders{}
	_result = &QueryDingSendStatisticalDataResponse{}
	_body, _err := client.QueryDingSendStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业文档统计数据
//
// @param request - QueryDocumentStatisticalDataRequest
//
// @param headers - QueryDocumentStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDocumentStatisticalDataResponse
func (client *Client) QueryDocumentStatisticalDataWithOptions(request *QueryDocumentStatisticalDataRequest, headers *QueryDocumentStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryDocumentStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryDocumentStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/documentData"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDocumentStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业文档统计数据
//
// @param request - QueryDocumentStatisticalDataRequest
//
// @return QueryDocumentStatisticalDataResponse
func (client *Client) QueryDocumentStatisticalData(request *QueryDocumentStatisticalDataRequest) (_result *QueryDocumentStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDocumentStatisticalDataHeaders{}
	_result = &QueryDocumentStatisticalDataResponse{}
	_body, _err := client.QueryDocumentStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询数据资产平台增购包购买信息
//
// @param headers - QueryDpaasDataPackageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDpaasDataPackageResponse
func (client *Client) QueryDpaasDataPackageWithOptions(headers *QueryDpaasDataPackageHeaders, runtime *util.RuntimeOptions) (_result *QueryDpaasDataPackageResponse, _err error) {
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
		Action:      tea.String("QueryDpaasDataPackage"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/dpaas/dataPackages/purchaseInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDpaasDataPackageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据资产平台增购包购买信息
//
// @return QueryDpaasDataPackageResponse
func (client *Client) QueryDpaasDataPackage() (_result *QueryDpaasDataPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDpaasDataPackageHeaders{}
	_result = &QueryDpaasDataPackageResponse{}
	_body, _err := client.QueryDpaasDataPackageWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业钉盘统计数据
//
// @param request - QueryDriveStatisticalDataRequest
//
// @param headers - QueryDriveStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDriveStatisticalDataResponse
func (client *Client) QueryDriveStatisticalDataWithOptions(request *QueryDriveStatisticalDataRequest, headers *QueryDriveStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryDriveStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryDriveStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/driveData"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDriveStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业钉盘统计数据
//
// @param request - QueryDriveStatisticalDataRequest
//
// @return QueryDriveStatisticalDataResponse
func (client *Client) QueryDriveStatisticalData(request *QueryDriveStatisticalDataRequest) (_result *QueryDriveStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDriveStatisticalDataHeaders{}
	_result = &QueryDriveStatisticalDataResponse{}
	_body, _err := client.QueryDriveStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 教育组织上下游授权信息查询服务
//
// @param request - QueryEduUnionAuthServiceRequest
//
// @param headers - QueryEduUnionAuthServiceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEduUnionAuthServiceResponse
func (client *Client) QueryEduUnionAuthServiceWithOptions(request *QueryEduUnionAuthServiceRequest, headers *QueryEduUnionAuthServiceHeaders, runtime *util.RuntimeOptions) (_result *QueryEduUnionAuthServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ds)) {
		query["ds"] = request.Ds
	}

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
		Action:      tea.String("QueryEduUnionAuthService"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/eduUnionAuthService"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryEduUnionAuthServiceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 教育组织上下游授权信息查询服务
//
// @param request - QueryEduUnionAuthServiceRequest
//
// @return QueryEduUnionAuthServiceResponse
func (client *Client) QueryEduUnionAuthService(request *QueryEduUnionAuthServiceRequest) (_result *QueryEduUnionAuthServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryEduUnionAuthServiceHeaders{}
	_result = &QueryEduUnionAuthServiceResponse{}
	_body, _err := client.QueryEduUnionAuthServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业员工类型统计数据
//
// @param request - QueryEmployeeTypeStatisticalDataRequest
//
// @param headers - QueryEmployeeTypeStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEmployeeTypeStatisticalDataResponse
func (client *Client) QueryEmployeeTypeStatisticalDataWithOptions(request *QueryEmployeeTypeStatisticalDataRequest, headers *QueryEmployeeTypeStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryEmployeeTypeStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryEmployeeTypeStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/employeeTypeData"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryEmployeeTypeStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业员工类型统计数据
//
// @param request - QueryEmployeeTypeStatisticalDataRequest
//
// @return QueryEmployeeTypeStatisticalDataResponse
func (client *Client) QueryEmployeeTypeStatisticalData(request *QueryEmployeeTypeStatisticalDataRequest) (_result *QueryEmployeeTypeStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryEmployeeTypeStatisticalDataHeaders{}
	_result = &QueryEmployeeTypeStatisticalDataResponse{}
	_body, _err := client.QueryEmployeeTypeStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 数据资产平台数据服务接口
//
// @param request - QueryGeneralDataServiceRequest
//
// @param headers - QueryGeneralDataServiceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGeneralDataServiceResponse
func (client *Client) QueryGeneralDataServiceWithOptions(request *QueryGeneralDataServiceRequest, headers *QueryGeneralDataServiceHeaders, runtime *util.RuntimeOptions) (_result *QueryGeneralDataServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["endDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnTotal)) {
		query["returnTotal"] = request.ReturnTotal
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["serviceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["startDate"] = request.StartDate
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
		Action:      tea.String("QueryGeneralDataService"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/generalDataServices"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGeneralDataServiceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 数据资产平台数据服务接口
//
// @param request - QueryGeneralDataServiceRequest
//
// @return QueryGeneralDataServiceResponse
func (client *Client) QueryGeneralDataService(request *QueryGeneralDataServiceRequest) (_result *QueryGeneralDataServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGeneralDataServiceHeaders{}
	_result = &QueryGeneralDataServiceResponse{}
	_body, _err := client.QueryGeneralDataServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 数据资产平台数据服务接口(支持部门、员工维度批量拉取)
//
// @param request - QueryGeneralDataServiceBatchRequest
//
// @param headers - QueryGeneralDataServiceBatchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGeneralDataServiceBatchResponse
func (client *Client) QueryGeneralDataServiceBatchWithOptions(request *QueryGeneralDataServiceBatchRequest, headers *QueryGeneralDataServiceBatchHeaders, runtime *util.RuntimeOptions) (_result *QueryGeneralDataServiceBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptIds)) {
		body["deptIds"] = request.DeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["endDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		body["filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnTotal)) {
		body["returnTotal"] = request.ReturnTotal
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["serviceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["startDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
		Action:      tea.String("QueryGeneralDataServiceBatch"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/dataServices/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGeneralDataServiceBatchResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 数据资产平台数据服务接口(支持部门、员工维度批量拉取)
//
// @param request - QueryGeneralDataServiceBatchRequest
//
// @return QueryGeneralDataServiceBatchResponse
func (client *Client) QueryGeneralDataServiceBatch(request *QueryGeneralDataServiceBatchRequest) (_result *QueryGeneralDataServiceBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGeneralDataServiceBatchHeaders{}
	_result = &QueryGeneralDataServiceBatchResponse{}
	_body, _err := client.QueryGeneralDataServiceBatchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 数据资产平台数据服务接口(查询数据更新日期)
//
// @param request - QueryGeneralDataUpdateDateRequest
//
// @param headers - QueryGeneralDataUpdateDateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGeneralDataUpdateDateResponse
func (client *Client) QueryGeneralDataUpdateDateWithOptions(request *QueryGeneralDataUpdateDateRequest, headers *QueryGeneralDataUpdateDateHeaders, runtime *util.RuntimeOptions) (_result *QueryGeneralDataUpdateDateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["serviceId"] = request.ServiceId
	}

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
		Action:      tea.String("QueryGeneralDataUpdateDate"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/dataUpdateDates"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGeneralDataUpdateDateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 数据资产平台数据服务接口(查询数据更新日期)
//
// @param request - QueryGeneralDataUpdateDateRequest
//
// @return QueryGeneralDataUpdateDateResponse
func (client *Client) QueryGeneralDataUpdateDate(request *QueryGeneralDataUpdateDateRequest) (_result *QueryGeneralDataUpdateDateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGeneralDataUpdateDateHeaders{}
	_result = &QueryGeneralDataUpdateDateResponse{}
	_body, _err := client.QueryGeneralDataUpdateDateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业群直播统计数据
//
// @param request - QueryGroupLiveStatisticalDataRequest
//
// @param headers - QueryGroupLiveStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGroupLiveStatisticalDataResponse
func (client *Client) QueryGroupLiveStatisticalDataWithOptions(request *QueryGroupLiveStatisticalDataRequest, headers *QueryGroupLiveStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryGroupLiveStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryGroupLiveStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/groupLiveData"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGroupLiveStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业群直播统计数据
//
// @param request - QueryGroupLiveStatisticalDataRequest
//
// @return QueryGroupLiveStatisticalDataResponse
func (client *Client) QueryGroupLiveStatisticalData(request *QueryGroupLiveStatisticalDataRequest) (_result *QueryGroupLiveStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGroupLiveStatisticalDataHeaders{}
	_result = &QueryGroupLiveStatisticalDataResponse{}
	_body, _err := client.QueryGroupLiveStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业群聊统计数据
//
// @param request - QueryGroupMessageStatisticalDataRequest
//
// @param headers - QueryGroupMessageStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGroupMessageStatisticalDataResponse
func (client *Client) QueryGroupMessageStatisticalDataWithOptions(request *QueryGroupMessageStatisticalDataRequest, headers *QueryGroupMessageStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryGroupMessageStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryGroupMessageStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/groupMessageData"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGroupMessageStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业群聊统计数据
//
// @param request - QueryGroupMessageStatisticalDataRequest
//
// @return QueryGroupMessageStatisticalDataResponse
func (client *Client) QueryGroupMessageStatisticalData(request *QueryGroupMessageStatisticalDataRequest) (_result *QueryGroupMessageStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGroupMessageStatisticalDataHeaders{}
	_result = &QueryGroupMessageStatisticalDataResponse{}
	_body, _err := client.QueryGroupMessageStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业钉钉运动统计数据
//
// @param request - QueryHealthStatisticalDataRequest
//
// @param headers - QueryHealthStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryHealthStatisticalDataResponse
func (client *Client) QueryHealthStatisticalDataWithOptions(request *QueryHealthStatisticalDataRequest, headers *QueryHealthStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryHealthStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryHealthStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/healtheUserData"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryHealthStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业钉钉运动统计数据
//
// @param request - QueryHealthStatisticalDataRequest
//
// @return QueryHealthStatisticalDataResponse
func (client *Client) QueryHealthStatisticalData(request *QueryHealthStatisticalDataRequest) (_result *QueryHealthStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryHealthStatisticalDataHeaders{}
	_result = &QueryHealthStatisticalDataResponse{}
	_body, _err := client.QueryHealthStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业邮箱统计数据
//
// @param request - QueryMailStatisticalDataRequest
//
// @param headers - QueryMailStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMailStatisticalDataResponse
func (client *Client) QueryMailStatisticalDataWithOptions(request *QueryMailStatisticalDataRequest, headers *QueryMailStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryMailStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryMailStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/mailData"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMailStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业邮箱统计数据
//
// @param request - QueryMailStatisticalDataRequest
//
// @return QueryMailStatisticalDataResponse
func (client *Client) QueryMailStatisticalData(request *QueryMailStatisticalDataRequest) (_result *QueryMailStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMailStatisticalDataHeaders{}
	_result = &QueryMailStatisticalDataResponse{}
	_body, _err := client.QueryMailStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取官方数据集数据
//
// @param request - QueryOfficialDataRequest
//
// @param headers - QueryOfficialDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOfficialDataResponse
func (client *Client) QueryOfficialDataWithOptions(request *QueryOfficialDataRequest, headers *QueryOfficialDataHeaders, runtime *util.RuntimeOptions) (_result *QueryOfficialDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Param)) {
		query["param"] = request.Param
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
		Action:      tea.String("QueryOfficialData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/datas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOfficialDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取官方数据集数据
//
// @param request - QueryOfficialDataRequest
//
// @return QueryOfficialDataResponse
func (client *Client) QueryOfficialData(request *QueryOfficialDataRequest) (_result *QueryOfficialDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOfficialDataHeaders{}
	_result = &QueryOfficialDataResponse{}
	_body, _err := client.QueryOfficialDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ISV获取官方数据集字段信息
//
// @param request - QueryOfficialDatasetFieldsRequest
//
// @param headers - QueryOfficialDatasetFieldsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOfficialDatasetFieldsResponse
func (client *Client) QueryOfficialDatasetFieldsWithOptions(request *QueryOfficialDatasetFieldsRequest, headers *QueryOfficialDatasetFieldsHeaders, runtime *util.RuntimeOptions) (_result *QueryOfficialDatasetFieldsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DsId)) {
		query["dsId"] = request.DsId
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
		Action:      tea.String("QueryOfficialDatasetFields"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/datasetFields"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOfficialDatasetFieldsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ISV获取官方数据集字段信息
//
// @param request - QueryOfficialDatasetFieldsRequest
//
// @return QueryOfficialDatasetFieldsResponse
func (client *Client) QueryOfficialDatasetFields(request *QueryOfficialDatasetFieldsRequest) (_result *QueryOfficialDatasetFieldsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOfficialDatasetFieldsHeaders{}
	_result = &QueryOfficialDatasetFieldsResponse{}
	_body, _err := client.QueryOfficialDatasetFieldsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ISV获取官方数据集列表
//
// @param request - QueryOfficialDatasetListRequest
//
// @param headers - QueryOfficialDatasetListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOfficialDatasetListResponse
func (client *Client) QueryOfficialDatasetListWithOptions(request *QueryOfficialDatasetListRequest, headers *QueryOfficialDatasetListHeaders, runtime *util.RuntimeOptions) (_result *QueryOfficialDatasetListResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("QueryOfficialDatasetList"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/datasetLists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOfficialDatasetListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ISV获取官方数据集列表
//
// @param request - QueryOfficialDatasetListRequest
//
// @return QueryOfficialDatasetListResponse
func (client *Client) QueryOfficialDatasetList(request *QueryOfficialDatasetListRequest) (_result *QueryOfficialDatasetListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOfficialDatasetListHeaders{}
	_result = &QueryOfficialDatasetListResponse{}
	_body, _err := client.QueryOfficialDatasetListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取官方数据集数据
//
// @param request - QueryOfficialFormDataRequest
//
// @param headers - QueryOfficialFormDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOfficialFormDataResponse
func (client *Client) QueryOfficialFormDataWithOptions(request *QueryOfficialFormDataRequest, headers *QueryOfficialFormDataHeaders, runtime *util.RuntimeOptions) (_result *QueryOfficialFormDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Param)) {
		body["param"] = request.Param
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
		Action:      tea.String("QueryOfficialFormData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/datas/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOfficialFormDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取官方数据集数据
//
// @param request - QueryOfficialFormDataRequest
//
// @return QueryOfficialFormDataResponse
func (client *Client) QueryOfficialFormData(request *QueryOfficialFormDataRequest) (_result *QueryOfficialFormDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOfficialFormDataHeaders{}
	_result = &QueryOfficialFormDataResponse{}
	_body, _err := client.QueryOfficialFormDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取HOLO中官方OA表单数据集数据
//
// @param request - QueryOfficialFormDataDirectHoloRequest
//
// @param headers - QueryOfficialFormDataDirectHoloHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOfficialFormDataDirectHoloResponse
func (client *Client) QueryOfficialFormDataDirectHoloWithOptions(request *QueryOfficialFormDataDirectHoloRequest, headers *QueryOfficialFormDataDirectHoloHeaders, runtime *util.RuntimeOptions) (_result *QueryOfficialFormDataDirectHoloResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Param)) {
		body["param"] = request.Param
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
		Action:      tea.String("QueryOfficialFormDataDirectHolo"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/oaDatas/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOfficialFormDataDirectHoloResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取HOLO中官方OA表单数据集数据
//
// @param request - QueryOfficialFormDataDirectHoloRequest
//
// @return QueryOfficialFormDataDirectHoloResponse
func (client *Client) QueryOfficialFormDataDirectHolo(request *QueryOfficialFormDataDirectHoloRequest) (_result *QueryOfficialFormDataDirectHoloResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOfficialFormDataDirectHoloHeaders{}
	_result = &QueryOfficialFormDataDirectHoloResponse{}
	_body, _err := client.QueryOfficialFormDataDirectHoloWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业用户在线统计数据
//
// @param request - QueryOnlineUserStatisticalDataRequest
//
// @param headers - QueryOnlineUserStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOnlineUserStatisticalDataResponse
func (client *Client) QueryOnlineUserStatisticalDataWithOptions(request *QueryOnlineUserStatisticalDataRequest, headers *QueryOnlineUserStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryOnlineUserStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryOnlineUserStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/onlineUserData"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOnlineUserStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业用户在线统计数据
//
// @param request - QueryOnlineUserStatisticalDataRequest
//
// @return QueryOnlineUserStatisticalDataResponse
func (client *Client) QueryOnlineUserStatisticalData(request *QueryOnlineUserStatisticalDataRequest) (_result *QueryOnlineUserStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOnlineUserStatisticalDataHeaders{}
	_result = &QueryOnlineUserStatisticalDataResponse{}
	_body, _err := client.QueryOnlineUserStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业接收红包统计数据
//
// @param request - QueryRedEnvelopeReciveStatisticalDataRequest
//
// @param headers - QueryRedEnvelopeReciveStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRedEnvelopeReciveStatisticalDataResponse
func (client *Client) QueryRedEnvelopeReciveStatisticalDataWithOptions(request *QueryRedEnvelopeReciveStatisticalDataRequest, headers *QueryRedEnvelopeReciveStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryRedEnvelopeReciveStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryRedEnvelopeReciveStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/redEnvelopeReciveData"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRedEnvelopeReciveStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业接收红包统计数据
//
// @param request - QueryRedEnvelopeReciveStatisticalDataRequest
//
// @return QueryRedEnvelopeReciveStatisticalDataResponse
func (client *Client) QueryRedEnvelopeReciveStatisticalData(request *QueryRedEnvelopeReciveStatisticalDataRequest) (_result *QueryRedEnvelopeReciveStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryRedEnvelopeReciveStatisticalDataHeaders{}
	_result = &QueryRedEnvelopeReciveStatisticalDataResponse{}
	_body, _err := client.QueryRedEnvelopeReciveStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业发送红包统计数据
//
// @param request - QueryRedEnvelopeSendStatisticalDataRequest
//
// @param headers - QueryRedEnvelopeSendStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRedEnvelopeSendStatisticalDataResponse
func (client *Client) QueryRedEnvelopeSendStatisticalDataWithOptions(request *QueryRedEnvelopeSendStatisticalDataRequest, headers *QueryRedEnvelopeSendStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryRedEnvelopeSendStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryRedEnvelopeSendStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/redEnvelopeSendData"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRedEnvelopeSendStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业发送红包统计数据
//
// @param request - QueryRedEnvelopeSendStatisticalDataRequest
//
// @return QueryRedEnvelopeSendStatisticalDataResponse
func (client *Client) QueryRedEnvelopeSendStatisticalData(request *QueryRedEnvelopeSendStatisticalDataRequest) (_result *QueryRedEnvelopeSendStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryRedEnvelopeSendStatisticalDataHeaders{}
	_result = &QueryRedEnvelopeSendStatisticalDataResponse{}
	_body, _err := client.QueryRedEnvelopeSendStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业日志统计数据
//
// @param request - QueryReportStatisticalDataRequest
//
// @param headers - QueryReportStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryReportStatisticalDataResponse
func (client *Client) QueryReportStatisticalDataWithOptions(request *QueryReportStatisticalDataRequest, headers *QueryReportStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryReportStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryReportStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/reportData"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryReportStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业日志统计数据
//
// @param request - QueryReportStatisticalDataRequest
//
// @return QueryReportStatisticalDataResponse
func (client *Client) QueryReportStatisticalData(request *QueryReportStatisticalDataRequest) (_result *QueryReportStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryReportStatisticalDataHeaders{}
	_result = &QueryReportStatisticalDataResponse{}
	_body, _err := client.QueryReportStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询数据大屏
//
// @param request - QueryScreenRequest
//
// @param headers - QueryScreenHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryScreenResponse
func (client *Client) QueryScreenWithOptions(request *QueryScreenRequest, headers *QueryScreenHeaders, runtime *util.RuntimeOptions) (_result *QueryScreenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

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
		Action:      tea.String("QueryScreen"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/screens"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryScreenResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据大屏
//
// @param request - QueryScreenRequest
//
// @return QueryScreenResponse
func (client *Client) QueryScreen(request *QueryScreenRequest) (_result *QueryScreenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryScreenHeaders{}
	_result = &QueryScreenResponse{}
	_body, _err := client.QueryScreenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询数据大屏模版
//
// @param request - QueryScreenTemplateRequest
//
// @param headers - QueryScreenTemplateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryScreenTemplateResponse
func (client *Client) QueryScreenTemplateWithOptions(request *QueryScreenTemplateRequest, headers *QueryScreenTemplateHeaders, runtime *util.RuntimeOptions) (_result *QueryScreenTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.Sample)) {
		query["sample"] = request.Sample
	}

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
		Action:      tea.String("QueryScreenTemplate"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/screenTemplates"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryScreenTemplateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据大屏模版
//
// @param request - QueryScreenTemplateRequest
//
// @return QueryScreenTemplateResponse
func (client *Client) QueryScreenTemplate(request *QueryScreenTemplateRequest) (_result *QueryScreenTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryScreenTemplateHeaders{}
	_result = &QueryScreenTemplateResponse{}
	_body, _err := client.QueryScreenTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业单聊统计数据
//
// @param request - QuerySingleMessageStatisticalDataRequest
//
// @param headers - QuerySingleMessageStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySingleMessageStatisticalDataResponse
func (client *Client) QuerySingleMessageStatisticalDataWithOptions(request *QuerySingleMessageStatisticalDataRequest, headers *QuerySingleMessageStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QuerySingleMessageStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QuerySingleMessageStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/singleMessagerData"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySingleMessageStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业单聊统计数据
//
// @param request - QuerySingleMessageStatisticalDataRequest
//
// @return QuerySingleMessageStatisticalDataResponse
func (client *Client) QuerySingleMessageStatisticalData(request *QuerySingleMessageStatisticalDataRequest) (_result *QuerySingleMessageStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QuerySingleMessageStatisticalDataHeaders{}
	_result = &QuerySingleMessageStatisticalDataResponse{}
	_body, _err := client.QuerySingleMessageStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业电话会议统计数据
//
// @param request - QueryTelMeetingStatisticalDataRequest
//
// @param headers - QueryTelMeetingStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTelMeetingStatisticalDataResponse
func (client *Client) QueryTelMeetingStatisticalDataWithOptions(request *QueryTelMeetingStatisticalDataRequest, headers *QueryTelMeetingStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryTelMeetingStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryTelMeetingStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/telMeetingData"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTelMeetingStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业电话会议统计数据
//
// @param request - QueryTelMeetingStatisticalDataRequest
//
// @return QueryTelMeetingStatisticalDataResponse
func (client *Client) QueryTelMeetingStatisticalData(request *QueryTelMeetingStatisticalDataRequest) (_result *QueryTelMeetingStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryTelMeetingStatisticalDataHeaders{}
	_result = &QueryTelMeetingStatisticalDataResponse{}
	_body, _err := client.QueryTelMeetingStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业待办统计数据
//
// @param request - QueryTodoStatisticalDataRequest
//
// @param headers - QueryTodoStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTodoStatisticalDataResponse
func (client *Client) QueryTodoStatisticalDataWithOptions(request *QueryTodoStatisticalDataRequest, headers *QueryTodoStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryTodoStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryTodoStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/todoUserData"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTodoStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业待办统计数据
//
// @param request - QueryTodoStatisticalDataRequest
//
// @return QueryTodoStatisticalDataResponse
func (client *Client) QueryTodoStatisticalData(request *QueryTodoStatisticalDataRequest) (_result *QueryTodoStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryTodoStatisticalDataHeaders{}
	_result = &QueryTodoStatisticalDataResponse{}
	_body, _err := client.QueryTodoStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 数据资产平台查询数据记录数
//
// @param request - QueryTotalDataCountServiceRequest
//
// @param headers - QueryTotalDataCountServiceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTotalDataCountServiceResponse
func (client *Client) QueryTotalDataCountServiceWithOptions(request *QueryTotalDataCountServiceRequest, headers *QueryTotalDataCountServiceHeaders, runtime *util.RuntimeOptions) (_result *QueryTotalDataCountServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptIds)) {
		body["deptIds"] = request.DeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["endDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["serviceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["startDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
		Action:      tea.String("QueryTotalDataCountService"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/datas/totalCounts/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTotalDataCountServiceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 数据资产平台查询数据记录数
//
// @param request - QueryTotalDataCountServiceRequest
//
// @return QueryTotalDataCountServiceResponse
func (client *Client) QueryTotalDataCountService(request *QueryTotalDataCountServiceRequest) (_result *QueryTotalDataCountServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryTotalDataCountServiceHeaders{}
	_result = &QueryTotalDataCountServiceResponse{}
	_body, _err := client.QueryTotalDataCountServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业视频会议统计数据
//
// @param request - QueryVedioMeetingStatisticalDataRequest
//
// @param headers - QueryVedioMeetingStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryVedioMeetingStatisticalDataResponse
func (client *Client) QueryVedioMeetingStatisticalDataWithOptions(request *QueryVedioMeetingStatisticalDataRequest, headers *QueryVedioMeetingStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryVedioMeetingStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryVedioMeetingStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/vedioMeetingData"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryVedioMeetingStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业视频会议统计数据
//
// @param request - QueryVedioMeetingStatisticalDataRequest
//
// @return QueryVedioMeetingStatisticalDataResponse
func (client *Client) QueryVedioMeetingStatisticalData(request *QueryVedioMeetingStatisticalDataRequest) (_result *QueryVedioMeetingStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryVedioMeetingStatisticalDataHeaders{}
	_result = &QueryVedioMeetingStatisticalDataResponse{}
	_body, _err := client.QueryVedioMeetingStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉参谋活跃分析（按日统计）指标接口
//
// @param request - QueryYydActiveDayStatisticalDataRequest
//
// @param headers - QueryYydActiveDayStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydActiveDayStatisticalDataResponse
func (client *Client) QueryYydActiveDayStatisticalDataWithOptions(request *QueryYydActiveDayStatisticalDataRequest, headers *QueryYydActiveDayStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydActiveDayStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydActiveDayStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydActiveDayDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydActiveDayStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉参谋活跃分析（按日统计）指标接口
//
// @param request - QueryYydActiveDayStatisticalDataRequest
//
// @return QueryYydActiveDayStatisticalDataResponse
func (client *Client) QueryYydActiveDayStatisticalData(request *QueryYydActiveDayStatisticalDataRequest) (_result *QueryYydActiveDayStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydActiveDayStatisticalDataHeaders{}
	_result = &QueryYydActiveDayStatisticalDataResponse{}
	_body, _err := client.QueryYydActiveDayStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉参谋活跃分析（按月统计）指标接口
//
// @param request - QueryYydActiveMonthStatisticalDataRequest
//
// @param headers - QueryYydActiveMonthStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydActiveMonthStatisticalDataResponse
func (client *Client) QueryYydActiveMonthStatisticalDataWithOptions(request *QueryYydActiveMonthStatisticalDataRequest, headers *QueryYydActiveMonthStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydActiveMonthStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydActiveMonthStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydActiveMonthDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydActiveMonthStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉参谋活跃分析（按月统计）指标接口
//
// @param request - QueryYydActiveMonthStatisticalDataRequest
//
// @return QueryYydActiveMonthStatisticalDataResponse
func (client *Client) QueryYydActiveMonthStatisticalData(request *QueryYydActiveMonthStatisticalDataRequest) (_result *QueryYydActiveMonthStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydActiveMonthStatisticalDataHeaders{}
	_result = &QueryYydActiveMonthStatisticalDataResponse{}
	_body, _err := client.QueryYydActiveMonthStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉参谋活跃分析（按周统计）指标接口
//
// @param request - QueryYydActiveWeekStatisticalDataRequest
//
// @param headers - QueryYydActiveWeekStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydActiveWeekStatisticalDataResponse
func (client *Client) QueryYydActiveWeekStatisticalDataWithOptions(request *QueryYydActiveWeekStatisticalDataRequest, headers *QueryYydActiveWeekStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydActiveWeekStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydActiveWeekStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydActiveWeekDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydActiveWeekStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉参谋活跃分析（按周统计）指标接口
//
// @param request - QueryYydActiveWeekStatisticalDataRequest
//
// @return QueryYydActiveWeekStatisticalDataResponse
func (client *Client) QueryYydActiveWeekStatisticalData(request *QueryYydActiveWeekStatisticalDataRequest) (_result *QueryYydActiveWeekStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydActiveWeekStatisticalDataHeaders{}
	_result = &QueryYydActiveWeekStatisticalDataResponse{}
	_body, _err := client.QueryYydActiveWeekStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋应用概况（按日统计）指标接口
//
// @param request - QueryYydAppDayStatisticalDataRequest
//
// @param headers - QueryYydAppDayStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydAppDayStatisticalDataResponse
func (client *Client) QueryYydAppDayStatisticalDataWithOptions(request *QueryYydAppDayStatisticalDataRequest, headers *QueryYydAppDayStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydAppDayStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydAppDayStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydAppDayDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydAppDayStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋应用概况（按日统计）指标接口
//
// @param request - QueryYydAppDayStatisticalDataRequest
//
// @return QueryYydAppDayStatisticalDataResponse
func (client *Client) QueryYydAppDayStatisticalData(request *QueryYydAppDayStatisticalDataRequest) (_result *QueryYydAppDayStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydAppDayStatisticalDataHeaders{}
	_result = &QueryYydAppDayStatisticalDataResponse{}
	_body, _err := client.QueryYydAppDayStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋应用概况（按月统计）指标接口
//
// @param request - QueryYydAppMonthStatisticalDataRequest
//
// @param headers - QueryYydAppMonthStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydAppMonthStatisticalDataResponse
func (client *Client) QueryYydAppMonthStatisticalDataWithOptions(request *QueryYydAppMonthStatisticalDataRequest, headers *QueryYydAppMonthStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydAppMonthStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydAppMonthStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydAppMonthDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydAppMonthStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋应用概况（按月统计）指标接口
//
// @param request - QueryYydAppMonthStatisticalDataRequest
//
// @return QueryYydAppMonthStatisticalDataResponse
func (client *Client) QueryYydAppMonthStatisticalData(request *QueryYydAppMonthStatisticalDataRequest) (_result *QueryYydAppMonthStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydAppMonthStatisticalDataHeaders{}
	_result = &QueryYydAppMonthStatisticalDataResponse{}
	_body, _err := client.QueryYydAppMonthStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋应用概况（累计）指标接口
//
// @param request - QueryYydAppStdStatisticalDataRequest
//
// @param headers - QueryYydAppStdStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydAppStdStatisticalDataResponse
func (client *Client) QueryYydAppStdStatisticalDataWithOptions(request *QueryYydAppStdStatisticalDataRequest, headers *QueryYydAppStdStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydAppStdStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydAppStdStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydAppStdDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydAppStdStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋应用概况（累计）指标接口
//
// @param request - QueryYydAppStdStatisticalDataRequest
//
// @return QueryYydAppStdStatisticalDataResponse
func (client *Client) QueryYydAppStdStatisticalData(request *QueryYydAppStdStatisticalDataRequest) (_result *QueryYydAppStdStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydAppStdStatisticalDataHeaders{}
	_result = &QueryYydAppStdStatisticalDataResponse{}
	_body, _err := client.QueryYydAppStdStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋应用概况（按周统计）指标接口
//
// @param request - QueryYydAppWeekStatisticalDataRequest
//
// @param headers - QueryYydAppWeekStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydAppWeekStatisticalDataResponse
func (client *Client) QueryYydAppWeekStatisticalDataWithOptions(request *QueryYydAppWeekStatisticalDataRequest, headers *QueryYydAppWeekStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydAppWeekStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydAppWeekStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydAppWeekDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydAppWeekStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋应用概况（按周统计）指标接口
//
// @param request - QueryYydAppWeekStatisticalDataRequest
//
// @return QueryYydAppWeekStatisticalDataResponse
func (client *Client) QueryYydAppWeekStatisticalData(request *QueryYydAppWeekStatisticalDataRequest) (_result *QueryYydAppWeekStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydAppWeekStatisticalDataHeaders{}
	_result = &QueryYydAppWeekStatisticalDataResponse{}
	_body, _err := client.QueryYydAppWeekStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋会议日程分析（按日统计）指标接口
//
// @param request - QueryYydCalendarDayStatisticalDataRequest
//
// @param headers - QueryYydCalendarDayStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydCalendarDayStatisticalDataResponse
func (client *Client) QueryYydCalendarDayStatisticalDataWithOptions(request *QueryYydCalendarDayStatisticalDataRequest, headers *QueryYydCalendarDayStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydCalendarDayStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydCalendarDayStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydCalendarDayDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydCalendarDayStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋会议日程分析（按日统计）指标接口
//
// @param request - QueryYydCalendarDayStatisticalDataRequest
//
// @return QueryYydCalendarDayStatisticalDataResponse
func (client *Client) QueryYydCalendarDayStatisticalData(request *QueryYydCalendarDayStatisticalDataRequest) (_result *QueryYydCalendarDayStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydCalendarDayStatisticalDataHeaders{}
	_result = &QueryYydCalendarDayStatisticalDataResponse{}
	_body, _err := client.QueryYydCalendarDayStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋会议日程分析（按月统计）指标接口
//
// @param request - QueryYydCalendarMonthStatisticalDataRequest
//
// @param headers - QueryYydCalendarMonthStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydCalendarMonthStatisticalDataResponse
func (client *Client) QueryYydCalendarMonthStatisticalDataWithOptions(request *QueryYydCalendarMonthStatisticalDataRequest, headers *QueryYydCalendarMonthStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydCalendarMonthStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydCalendarMonthStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydCalendarMonthDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydCalendarMonthStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋会议日程分析（按月统计）指标接口
//
// @param request - QueryYydCalendarMonthStatisticalDataRequest
//
// @return QueryYydCalendarMonthStatisticalDataResponse
func (client *Client) QueryYydCalendarMonthStatisticalData(request *QueryYydCalendarMonthStatisticalDataRequest) (_result *QueryYydCalendarMonthStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydCalendarMonthStatisticalDataHeaders{}
	_result = &QueryYydCalendarMonthStatisticalDataResponse{}
	_body, _err := client.QueryYydCalendarMonthStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋会议日程分析（按周统计）指标接口
//
// @param request - QueryYydCalendarWeekStatisticalDataRequest
//
// @param headers - QueryYydCalendarWeekStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydCalendarWeekStatisticalDataResponse
func (client *Client) QueryYydCalendarWeekStatisticalDataWithOptions(request *QueryYydCalendarWeekStatisticalDataRequest, headers *QueryYydCalendarWeekStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydCalendarWeekStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydCalendarWeekStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydCalendarWeekDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydCalendarWeekStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋会议日程分析（按周统计）指标接口
//
// @param request - QueryYydCalendarWeekStatisticalDataRequest
//
// @return QueryYydCalendarWeekStatisticalDataResponse
func (client *Client) QueryYydCalendarWeekStatisticalData(request *QueryYydCalendarWeekStatisticalDataRequest) (_result *QueryYydCalendarWeekStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydCalendarWeekStatisticalDataHeaders{}
	_result = &QueryYydCalendarWeekStatisticalDataResponse{}
	_body, _err := client.QueryYydCalendarWeekStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋钉消息分析（按日统计）指标接口
//
// @param request - QueryYydDingMsgDayStatisticalDataRequest
//
// @param headers - QueryYydDingMsgDayStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydDingMsgDayStatisticalDataResponse
func (client *Client) QueryYydDingMsgDayStatisticalDataWithOptions(request *QueryYydDingMsgDayStatisticalDataRequest, headers *QueryYydDingMsgDayStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydDingMsgDayStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydDingMsgDayStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydDingMsgDayDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydDingMsgDayStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋钉消息分析（按日统计）指标接口
//
// @param request - QueryYydDingMsgDayStatisticalDataRequest
//
// @return QueryYydDingMsgDayStatisticalDataResponse
func (client *Client) QueryYydDingMsgDayStatisticalData(request *QueryYydDingMsgDayStatisticalDataRequest) (_result *QueryYydDingMsgDayStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydDingMsgDayStatisticalDataHeaders{}
	_result = &QueryYydDingMsgDayStatisticalDataResponse{}
	_body, _err := client.QueryYydDingMsgDayStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋钉消息分析（按月统计）指标接口
//
// @param request - QueryYydDingMsgMonthStatisticalDataRequest
//
// @param headers - QueryYydDingMsgMonthStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydDingMsgMonthStatisticalDataResponse
func (client *Client) QueryYydDingMsgMonthStatisticalDataWithOptions(request *QueryYydDingMsgMonthStatisticalDataRequest, headers *QueryYydDingMsgMonthStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydDingMsgMonthStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydDingMsgMonthStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydDingMsgMonthDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydDingMsgMonthStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋钉消息分析（按月统计）指标接口
//
// @param request - QueryYydDingMsgMonthStatisticalDataRequest
//
// @return QueryYydDingMsgMonthStatisticalDataResponse
func (client *Client) QueryYydDingMsgMonthStatisticalData(request *QueryYydDingMsgMonthStatisticalDataRequest) (_result *QueryYydDingMsgMonthStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydDingMsgMonthStatisticalDataHeaders{}
	_result = &QueryYydDingMsgMonthStatisticalDataResponse{}
	_body, _err := client.QueryYydDingMsgMonthStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋钉消息分析（按周统计）指标接口
//
// @param request - QueryYydDingMsgWeekStatisticalDataRequest
//
// @param headers - QueryYydDingMsgWeekStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydDingMsgWeekStatisticalDataResponse
func (client *Client) QueryYydDingMsgWeekStatisticalDataWithOptions(request *QueryYydDingMsgWeekStatisticalDataRequest, headers *QueryYydDingMsgWeekStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydDingMsgWeekStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydDingMsgWeekStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydDingMsgWeekDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydDingMsgWeekStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋钉消息分析（按周统计）指标接口
//
// @param request - QueryYydDingMsgWeekStatisticalDataRequest
//
// @return QueryYydDingMsgWeekStatisticalDataResponse
func (client *Client) QueryYydDingMsgWeekStatisticalData(request *QueryYydDingMsgWeekStatisticalDataRequest) (_result *QueryYydDingMsgWeekStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydDingMsgWeekStatisticalDataHeaders{}
	_result = &QueryYydDingMsgWeekStatisticalDataResponse{}
	_body, _err := client.QueryYydDingMsgWeekStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋群聊分析（按日统计）指标接口
//
// @param request - QueryYydGroupMsgDayStatisticalDataRequest
//
// @param headers - QueryYydGroupMsgDayStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydGroupMsgDayStatisticalDataResponse
func (client *Client) QueryYydGroupMsgDayStatisticalDataWithOptions(request *QueryYydGroupMsgDayStatisticalDataRequest, headers *QueryYydGroupMsgDayStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydGroupMsgDayStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydGroupMsgDayStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydGroupMsgDayDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydGroupMsgDayStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋群聊分析（按日统计）指标接口
//
// @param request - QueryYydGroupMsgDayStatisticalDataRequest
//
// @return QueryYydGroupMsgDayStatisticalDataResponse
func (client *Client) QueryYydGroupMsgDayStatisticalData(request *QueryYydGroupMsgDayStatisticalDataRequest) (_result *QueryYydGroupMsgDayStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydGroupMsgDayStatisticalDataHeaders{}
	_result = &QueryYydGroupMsgDayStatisticalDataResponse{}
	_body, _err := client.QueryYydGroupMsgDayStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋群聊分析（按月统计）指标接口
//
// @param request - QueryYydGroupMsgMonthStatisticalDataRequest
//
// @param headers - QueryYydGroupMsgMonthStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydGroupMsgMonthStatisticalDataResponse
func (client *Client) QueryYydGroupMsgMonthStatisticalDataWithOptions(request *QueryYydGroupMsgMonthStatisticalDataRequest, headers *QueryYydGroupMsgMonthStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydGroupMsgMonthStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydGroupMsgMonthStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydGroupMsgMonthDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydGroupMsgMonthStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋群聊分析（按月统计）指标接口
//
// @param request - QueryYydGroupMsgMonthStatisticalDataRequest
//
// @return QueryYydGroupMsgMonthStatisticalDataResponse
func (client *Client) QueryYydGroupMsgMonthStatisticalData(request *QueryYydGroupMsgMonthStatisticalDataRequest) (_result *QueryYydGroupMsgMonthStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydGroupMsgMonthStatisticalDataHeaders{}
	_result = &QueryYydGroupMsgMonthStatisticalDataResponse{}
	_body, _err := client.QueryYydGroupMsgMonthStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋群聊分析（按周统计）指标接口
//
// @param request - QueryYydGroupMsgWeekStatisticalDataRequest
//
// @param headers - QueryYydGroupMsgWeekStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydGroupMsgWeekStatisticalDataResponse
func (client *Client) QueryYydGroupMsgWeekStatisticalDataWithOptions(request *QueryYydGroupMsgWeekStatisticalDataRequest, headers *QueryYydGroupMsgWeekStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydGroupMsgWeekStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydGroupMsgWeekStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydGroupMsgWeekDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydGroupMsgWeekStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋群聊分析（按周统计）指标接口
//
// @param request - QueryYydGroupMsgWeekStatisticalDataRequest
//
// @return QueryYydGroupMsgWeekStatisticalDataResponse
func (client *Client) QueryYydGroupMsgWeekStatisticalData(request *QueryYydGroupMsgWeekStatisticalDataRequest) (_result *QueryYydGroupMsgWeekStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydGroupMsgWeekStatisticalDataHeaders{}
	_result = &QueryYydGroupMsgWeekStatisticalDataResponse{}
	_body, _err := client.QueryYydGroupMsgWeekStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋日志分析（按日统计）指标接口
//
// @param request - QueryYydLogDayStatisticalDataRequest
//
// @param headers - QueryYydLogDayStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydLogDayStatisticalDataResponse
func (client *Client) QueryYydLogDayStatisticalDataWithOptions(request *QueryYydLogDayStatisticalDataRequest, headers *QueryYydLogDayStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydLogDayStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydLogDayStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydLogDayDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydLogDayStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋日志分析（按日统计）指标接口
//
// @param request - QueryYydLogDayStatisticalDataRequest
//
// @return QueryYydLogDayStatisticalDataResponse
func (client *Client) QueryYydLogDayStatisticalData(request *QueryYydLogDayStatisticalDataRequest) (_result *QueryYydLogDayStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydLogDayStatisticalDataHeaders{}
	_result = &QueryYydLogDayStatisticalDataResponse{}
	_body, _err := client.QueryYydLogDayStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋日志分析（按月统计）指标接口
//
// @param request - QueryYydLogMonthStatisticalDataRequest
//
// @param headers - QueryYydLogMonthStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydLogMonthStatisticalDataResponse
func (client *Client) QueryYydLogMonthStatisticalDataWithOptions(request *QueryYydLogMonthStatisticalDataRequest, headers *QueryYydLogMonthStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydLogMonthStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydLogMonthStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydLogMonthDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydLogMonthStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋日志分析（按月统计）指标接口
//
// @param request - QueryYydLogMonthStatisticalDataRequest
//
// @return QueryYydLogMonthStatisticalDataResponse
func (client *Client) QueryYydLogMonthStatisticalData(request *QueryYydLogMonthStatisticalDataRequest) (_result *QueryYydLogMonthStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydLogMonthStatisticalDataHeaders{}
	_result = &QueryYydLogMonthStatisticalDataResponse{}
	_body, _err := client.QueryYydLogMonthStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋日志分析（按周统计）指标接口
//
// @param request - QueryYydLogWeekStatisticalDataRequest
//
// @param headers - QueryYydLogWeekStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydLogWeekStatisticalDataResponse
func (client *Client) QueryYydLogWeekStatisticalDataWithOptions(request *QueryYydLogWeekStatisticalDataRequest, headers *QueryYydLogWeekStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydLogWeekStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydLogWeekStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydLogWeekDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydLogWeekStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋日志分析（按周统计）指标接口
//
// @param request - QueryYydLogWeekStatisticalDataRequest
//
// @return QueryYydLogWeekStatisticalDataResponse
func (client *Client) QueryYydLogWeekStatisticalData(request *QueryYydLogWeekStatisticalDataRequest) (_result *QueryYydLogWeekStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydLogWeekStatisticalDataHeaders{}
	_result = &QueryYydLogWeekStatisticalDataResponse{}
	_body, _err := client.QueryYydLogWeekStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋钉会议分析（按日统计）指标接口
//
// @param request - QueryYydMeetingDayStatisticalDataRequest
//
// @param headers - QueryYydMeetingDayStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydMeetingDayStatisticalDataResponse
func (client *Client) QueryYydMeetingDayStatisticalDataWithOptions(request *QueryYydMeetingDayStatisticalDataRequest, headers *QueryYydMeetingDayStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydMeetingDayStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydMeetingDayStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydMeetingDayDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydMeetingDayStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋钉会议分析（按日统计）指标接口
//
// @param request - QueryYydMeetingDayStatisticalDataRequest
//
// @return QueryYydMeetingDayStatisticalDataResponse
func (client *Client) QueryYydMeetingDayStatisticalData(request *QueryYydMeetingDayStatisticalDataRequest) (_result *QueryYydMeetingDayStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydMeetingDayStatisticalDataHeaders{}
	_result = &QueryYydMeetingDayStatisticalDataResponse{}
	_body, _err := client.QueryYydMeetingDayStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋钉会议分析（按月统计）指标接口
//
// @param request - QueryYydMeetingMonthStatisticalDataRequest
//
// @param headers - QueryYydMeetingMonthStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydMeetingMonthStatisticalDataResponse
func (client *Client) QueryYydMeetingMonthStatisticalDataWithOptions(request *QueryYydMeetingMonthStatisticalDataRequest, headers *QueryYydMeetingMonthStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydMeetingMonthStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydMeetingMonthStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydMeetingMonthDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydMeetingMonthStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋钉会议分析（按月统计）指标接口
//
// @param request - QueryYydMeetingMonthStatisticalDataRequest
//
// @return QueryYydMeetingMonthStatisticalDataResponse
func (client *Client) QueryYydMeetingMonthStatisticalData(request *QueryYydMeetingMonthStatisticalDataRequest) (_result *QueryYydMeetingMonthStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydMeetingMonthStatisticalDataHeaders{}
	_result = &QueryYydMeetingMonthStatisticalDataResponse{}
	_body, _err := client.QueryYydMeetingMonthStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋钉会议分析（按周统计）指标接口
//
// @param request - QueryYydMeetingWeekStatisticalDataRequest
//
// @param headers - QueryYydMeetingWeekStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydMeetingWeekStatisticalDataResponse
func (client *Client) QueryYydMeetingWeekStatisticalDataWithOptions(request *QueryYydMeetingWeekStatisticalDataRequest, headers *QueryYydMeetingWeekStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydMeetingWeekStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydMeetingWeekStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydMeetingWeekDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydMeetingWeekStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋钉会议分析（按周统计）指标接口
//
// @param request - QueryYydMeetingWeekStatisticalDataRequest
//
// @return QueryYydMeetingWeekStatisticalDataResponse
func (client *Client) QueryYydMeetingWeekStatisticalData(request *QueryYydMeetingWeekStatisticalDataRequest) (_result *QueryYydMeetingWeekStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydMeetingWeekStatisticalDataHeaders{}
	_result = &QueryYydMeetingWeekStatisticalDataResponse{}
	_body, _err := client.QueryYydMeetingWeekStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋通知分析（按日统计）指标接口
//
// @param request - QueryYydNoticeDayStatisticalDataRequest
//
// @param headers - QueryYydNoticeDayStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydNoticeDayStatisticalDataResponse
func (client *Client) QueryYydNoticeDayStatisticalDataWithOptions(request *QueryYydNoticeDayStatisticalDataRequest, headers *QueryYydNoticeDayStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydNoticeDayStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydNoticeDayStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydNoticeDayDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydNoticeDayStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋通知分析（按日统计）指标接口
//
// @param request - QueryYydNoticeDayStatisticalDataRequest
//
// @return QueryYydNoticeDayStatisticalDataResponse
func (client *Client) QueryYydNoticeDayStatisticalData(request *QueryYydNoticeDayStatisticalDataRequest) (_result *QueryYydNoticeDayStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydNoticeDayStatisticalDataHeaders{}
	_result = &QueryYydNoticeDayStatisticalDataResponse{}
	_body, _err := client.QueryYydNoticeDayStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋通知分析（按月统计）指标接口
//
// @param request - QueryYydNoticeMonthStatisticalDataRequest
//
// @param headers - QueryYydNoticeMonthStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydNoticeMonthStatisticalDataResponse
func (client *Client) QueryYydNoticeMonthStatisticalDataWithOptions(request *QueryYydNoticeMonthStatisticalDataRequest, headers *QueryYydNoticeMonthStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydNoticeMonthStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydNoticeMonthStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydNoticeMonthDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydNoticeMonthStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋通知分析（按月统计）指标接口
//
// @param request - QueryYydNoticeMonthStatisticalDataRequest
//
// @return QueryYydNoticeMonthStatisticalDataResponse
func (client *Client) QueryYydNoticeMonthStatisticalData(request *QueryYydNoticeMonthStatisticalDataRequest) (_result *QueryYydNoticeMonthStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydNoticeMonthStatisticalDataHeaders{}
	_result = &QueryYydNoticeMonthStatisticalDataResponse{}
	_body, _err := client.QueryYydNoticeMonthStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋通知分析（按周统计）指标接口
//
// @param request - QueryYydNoticeWeekStatisticalDataRequest
//
// @param headers - QueryYydNoticeWeekStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydNoticeWeekStatisticalDataResponse
func (client *Client) QueryYydNoticeWeekStatisticalDataWithOptions(request *QueryYydNoticeWeekStatisticalDataRequest, headers *QueryYydNoticeWeekStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydNoticeWeekStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydNoticeWeekStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydNoticeWeekDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydNoticeWeekStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋通知分析（按周统计）指标接口
//
// @param request - QueryYydNoticeWeekStatisticalDataRequest
//
// @return QueryYydNoticeWeekStatisticalDataResponse
func (client *Client) QueryYydNoticeWeekStatisticalData(request *QueryYydNoticeWeekStatisticalDataRequest) (_result *QueryYydNoticeWeekStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydNoticeWeekStatisticalDataHeaders{}
	_result = &QueryYydNoticeWeekStatisticalDataResponse{}
	_body, _err := client.QueryYydNoticeWeekStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋单聊分析（按日统计）指标接口
//
// @param request - QueryYydSingleMsgDayStatisticalDataRequest
//
// @param headers - QueryYydSingleMsgDayStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydSingleMsgDayStatisticalDataResponse
func (client *Client) QueryYydSingleMsgDayStatisticalDataWithOptions(request *QueryYydSingleMsgDayStatisticalDataRequest, headers *QueryYydSingleMsgDayStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydSingleMsgDayStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydSingleMsgDayStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydSingleMsgDayDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydSingleMsgDayStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋单聊分析（按日统计）指标接口
//
// @param request - QueryYydSingleMsgDayStatisticalDataRequest
//
// @return QueryYydSingleMsgDayStatisticalDataResponse
func (client *Client) QueryYydSingleMsgDayStatisticalData(request *QueryYydSingleMsgDayStatisticalDataRequest) (_result *QueryYydSingleMsgDayStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydSingleMsgDayStatisticalDataHeaders{}
	_result = &QueryYydSingleMsgDayStatisticalDataResponse{}
	_body, _err := client.QueryYydSingleMsgDayStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋单聊分析（按月统计）指标接口
//
// @param request - QueryYydSingleMsgMonthStatisticalDataRequest
//
// @param headers - QueryYydSingleMsgMonthStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydSingleMsgMonthStatisticalDataResponse
func (client *Client) QueryYydSingleMsgMonthStatisticalDataWithOptions(request *QueryYydSingleMsgMonthStatisticalDataRequest, headers *QueryYydSingleMsgMonthStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydSingleMsgMonthStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydSingleMsgMonthStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydSingleMsgMonthDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydSingleMsgMonthStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋单聊分析（按月统计）指标接口
//
// @param request - QueryYydSingleMsgMonthStatisticalDataRequest
//
// @return QueryYydSingleMsgMonthStatisticalDataResponse
func (client *Client) QueryYydSingleMsgMonthStatisticalData(request *QueryYydSingleMsgMonthStatisticalDataRequest) (_result *QueryYydSingleMsgMonthStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydSingleMsgMonthStatisticalDataHeaders{}
	_result = &QueryYydSingleMsgMonthStatisticalDataResponse{}
	_body, _err := client.QueryYydSingleMsgMonthStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋单聊分析（按周统计）指标接口
//
// @param request - QueryYydSingleMsgWeekStatisticalDataRequest
//
// @param headers - QueryYydSingleMsgWeekStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydSingleMsgWeekStatisticalDataResponse
func (client *Client) QueryYydSingleMsgWeekStatisticalDataWithOptions(request *QueryYydSingleMsgWeekStatisticalDataRequest, headers *QueryYydSingleMsgWeekStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydSingleMsgWeekStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydSingleMsgWeekStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydSingleMsgWeekDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydSingleMsgWeekStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋单聊分析（按周统计）指标接口
//
// @param request - QueryYydSingleMsgWeekStatisticalDataRequest
//
// @return QueryYydSingleMsgWeekStatisticalDataResponse
func (client *Client) QueryYydSingleMsgWeekStatisticalData(request *QueryYydSingleMsgWeekStatisticalDataRequest) (_result *QueryYydSingleMsgWeekStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydSingleMsgWeekStatisticalDataHeaders{}
	_result = &QueryYydSingleMsgWeekStatisticalDataResponse{}
	_body, _err := client.QueryYydSingleMsgWeekStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋消息概览（按日统计）指标接口
//
// @param request - QueryYydToatlMsgDayStatisticalDataRequest
//
// @param headers - QueryYydToatlMsgDayStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydToatlMsgDayStatisticalDataResponse
func (client *Client) QueryYydToatlMsgDayStatisticalDataWithOptions(request *QueryYydToatlMsgDayStatisticalDataRequest, headers *QueryYydToatlMsgDayStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydToatlMsgDayStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydToatlMsgDayStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydToatlMsgDayDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydToatlMsgDayStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋消息概览（按日统计）指标接口
//
// @param request - QueryYydToatlMsgDayStatisticalDataRequest
//
// @return QueryYydToatlMsgDayStatisticalDataResponse
func (client *Client) QueryYydToatlMsgDayStatisticalData(request *QueryYydToatlMsgDayStatisticalDataRequest) (_result *QueryYydToatlMsgDayStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydToatlMsgDayStatisticalDataHeaders{}
	_result = &QueryYydToatlMsgDayStatisticalDataResponse{}
	_body, _err := client.QueryYydToatlMsgDayStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋消息概览（按月统计）指标接口
//
// @param request - QueryYydToatlMsgMonthStatisticalDataRequest
//
// @param headers - QueryYydToatlMsgMonthStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydToatlMsgMonthStatisticalDataResponse
func (client *Client) QueryYydToatlMsgMonthStatisticalDataWithOptions(request *QueryYydToatlMsgMonthStatisticalDataRequest, headers *QueryYydToatlMsgMonthStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydToatlMsgMonthStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydToatlMsgMonthStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydToatlMsgMonthDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydToatlMsgMonthStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋消息概览（按月统计）指标接口
//
// @param request - QueryYydToatlMsgMonthStatisticalDataRequest
//
// @return QueryYydToatlMsgMonthStatisticalDataResponse
func (client *Client) QueryYydToatlMsgMonthStatisticalData(request *QueryYydToatlMsgMonthStatisticalDataRequest) (_result *QueryYydToatlMsgMonthStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydToatlMsgMonthStatisticalDataHeaders{}
	_result = &QueryYydToatlMsgMonthStatisticalDataResponse{}
	_body, _err := client.QueryYydToatlMsgMonthStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋消息概览（按周统计）指标接口
//
// @param request - QueryYydToatlMsgWeekStatisticalDataRequest
//
// @param headers - QueryYydToatlMsgWeekStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydToatlMsgWeekStatisticalDataResponse
func (client *Client) QueryYydToatlMsgWeekStatisticalDataWithOptions(request *QueryYydToatlMsgWeekStatisticalDataRequest, headers *QueryYydToatlMsgWeekStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydToatlMsgWeekStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydToatlMsgWeekStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydToatlMsgWeekDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydToatlMsgWeekStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋消息概览（按周统计）指标接口
//
// @param request - QueryYydToatlMsgWeekStatisticalDataRequest
//
// @return QueryYydToatlMsgWeekStatisticalDataResponse
func (client *Client) QueryYydToatlMsgWeekStatisticalData(request *QueryYydToatlMsgWeekStatisticalDataRequest) (_result *QueryYydToatlMsgWeekStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydToatlMsgWeekStatisticalDataHeaders{}
	_result = &QueryYydToatlMsgWeekStatisticalDataResponse{}
	_body, _err := client.QueryYydToatlMsgWeekStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋待办分析（按日统计）指标接口
//
// @param request - QueryYydTodoDayStatisticalDataRequest
//
// @param headers - QueryYydTodoDayStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydTodoDayStatisticalDataResponse
func (client *Client) QueryYydTodoDayStatisticalDataWithOptions(request *QueryYydTodoDayStatisticalDataRequest, headers *QueryYydTodoDayStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydTodoDayStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydTodoDayStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydTodoDayDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydTodoDayStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋待办分析（按日统计）指标接口
//
// @param request - QueryYydTodoDayStatisticalDataRequest
//
// @return QueryYydTodoDayStatisticalDataResponse
func (client *Client) QueryYydTodoDayStatisticalData(request *QueryYydTodoDayStatisticalDataRequest) (_result *QueryYydTodoDayStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydTodoDayStatisticalDataHeaders{}
	_result = &QueryYydTodoDayStatisticalDataResponse{}
	_body, _err := client.QueryYydTodoDayStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋待办分析（按月统计）指标接口
//
// @param request - QueryYydTodoMonthStatisticalDataRequest
//
// @param headers - QueryYydTodoMonthStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydTodoMonthStatisticalDataResponse
func (client *Client) QueryYydTodoMonthStatisticalDataWithOptions(request *QueryYydTodoMonthStatisticalDataRequest, headers *QueryYydTodoMonthStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydTodoMonthStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydTodoMonthStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydTodoMonthDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydTodoMonthStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋待办分析（按月统计）指标接口
//
// @param request - QueryYydTodoMonthStatisticalDataRequest
//
// @return QueryYydTodoMonthStatisticalDataResponse
func (client *Client) QueryYydTodoMonthStatisticalData(request *QueryYydTodoMonthStatisticalDataRequest) (_result *QueryYydTodoMonthStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydTodoMonthStatisticalDataHeaders{}
	_result = &QueryYydTodoMonthStatisticalDataResponse{}
	_body, _err := client.QueryYydTodoMonthStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋待办分析（按周统计）指标接口
//
// @param request - QueryYydTodoWeekStatisticalDataRequest
//
// @param headers - QueryYydTodoWeekStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydTodoWeekStatisticalDataResponse
func (client *Client) QueryYydTodoWeekStatisticalDataWithOptions(request *QueryYydTodoWeekStatisticalDataRequest, headers *QueryYydTodoWeekStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydTodoWeekStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydTodoWeekStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydTodoWeekDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydTodoWeekStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉数字参谋待办分析（按周统计）指标接口
//
// @param request - QueryYydTodoWeekStatisticalDataRequest
//
// @return QueryYydTodoWeekStatisticalDataResponse
func (client *Client) QueryYydTodoWeekStatisticalData(request *QueryYydTodoWeekStatisticalDataRequest) (_result *QueryYydTodoWeekStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydTodoWeekStatisticalDataHeaders{}
	_result = &QueryYydTodoWeekStatisticalDataResponse{}
	_body, _err := client.QueryYydTodoWeekStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉参谋全局概览（按日统计）指标接口
//
// @param request - QueryYydTotalDayStatisticalDataRequest
//
// @param headers - QueryYydTotalDayStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydTotalDayStatisticalDataResponse
func (client *Client) QueryYydTotalDayStatisticalDataWithOptions(request *QueryYydTotalDayStatisticalDataRequest, headers *QueryYydTotalDayStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydTotalDayStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydTotalDayStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydTotalDayDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydTotalDayStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉参谋全局概览（按日统计）指标接口
//
// @param request - QueryYydTotalDayStatisticalDataRequest
//
// @return QueryYydTotalDayStatisticalDataResponse
func (client *Client) QueryYydTotalDayStatisticalData(request *QueryYydTotalDayStatisticalDataRequest) (_result *QueryYydTotalDayStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydTotalDayStatisticalDataHeaders{}
	_result = &QueryYydTotalDayStatisticalDataResponse{}
	_body, _err := client.QueryYydTotalDayStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉参谋全局概览（按月统计）指标接口
//
// @param request - QueryYydTotalMonthStatisticalDataRequest
//
// @param headers - QueryYydTotalMonthStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydTotalMonthStatisticalDataResponse
func (client *Client) QueryYydTotalMonthStatisticalDataWithOptions(request *QueryYydTotalMonthStatisticalDataRequest, headers *QueryYydTotalMonthStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydTotalMonthStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydTotalMonthStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydTotalMonthDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydTotalMonthStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉参谋全局概览（按月统计）指标接口
//
// @param request - QueryYydTotalMonthStatisticalDataRequest
//
// @return QueryYydTotalMonthStatisticalDataResponse
func (client *Client) QueryYydTotalMonthStatisticalData(request *QueryYydTotalMonthStatisticalDataRequest) (_result *QueryYydTotalMonthStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydTotalMonthStatisticalDataHeaders{}
	_result = &QueryYydTotalMonthStatisticalDataResponse{}
	_body, _err := client.QueryYydTotalMonthStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉参谋全局概览（累计）指标接口
//
// @param request - QueryYydTotalStdStatisticalDataRequest
//
// @param headers - QueryYydTotalStdStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydTotalStdStatisticalDataResponse
func (client *Client) QueryYydTotalStdStatisticalDataWithOptions(request *QueryYydTotalStdStatisticalDataRequest, headers *QueryYydTotalStdStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydTotalStdStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydTotalStdStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydTotalStdDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydTotalStdStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉参谋全局概览（累计）指标接口
//
// @param request - QueryYydTotalStdStatisticalDataRequest
//
// @return QueryYydTotalStdStatisticalDataResponse
func (client *Client) QueryYydTotalStdStatisticalData(request *QueryYydTotalStdStatisticalDataRequest) (_result *QueryYydTotalStdStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydTotalStdStatisticalDataHeaders{}
	_result = &QueryYydTotalStdStatisticalDataResponse{}
	_body, _err := client.QueryYydTotalStdStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 亚运钉参谋全局概览（按周统计）指标接口
//
// @param request - QueryYydTotalWeekStatisticalDataRequest
//
// @param headers - QueryYydTotalWeekStatisticalDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYydTotalWeekStatisticalDataResponse
func (client *Client) QueryYydTotalWeekStatisticalDataWithOptions(request *QueryYydTotalWeekStatisticalDataRequest, headers *QueryYydTotalWeekStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryYydTotalWeekStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

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
		Action:      tea.String("QueryYydTotalWeekStatisticalData"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/yydTotalWeekDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYydTotalWeekStatisticalDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 亚运钉参谋全局概览（按周统计）指标接口
//
// @param request - QueryYydTotalWeekStatisticalDataRequest
//
// @return QueryYydTotalWeekStatisticalDataResponse
func (client *Client) QueryYydTotalWeekStatisticalData(request *QueryYydTotalWeekStatisticalDataRequest) (_result *QueryYydTotalWeekStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryYydTotalWeekStatisticalDataHeaders{}
	_result = &QueryYydTotalWeekStatisticalDataResponse{}
	_body, _err := client.QueryYydTotalWeekStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过关键词搜索企业
//
// @param request - SearchCompanyRequest
//
// @param headers - SearchCompanyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchCompanyResponse
func (client *Client) SearchCompanyWithOptions(request *SearchCompanyRequest, headers *SearchCompanyHeaders, runtime *util.RuntimeOptions) (_result *SearchCompanyResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["searchKey"] = request.SearchKey
	}

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
		Action:      tea.String("SearchCompany"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/keywords/companies"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchCompanyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过关键词搜索企业
//
// @param request - SearchCompanyRequest
//
// @return SearchCompanyResponse
func (client *Client) SearchCompany(request *SearchCompanyRequest) (_result *SearchCompanyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchCompanyHeaders{}
	_result = &SearchCompanyResponse{}
	_body, _err := client.SearchCompanyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 同步数据大屏
//
// @param request - SyncDataScreenRequest
//
// @param headers - SyncDataScreenHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncDataScreenResponse
func (client *Client) SyncDataScreenWithOptions(request *SyncDataScreenRequest, headers *SyncDataScreenHeaders, runtime *util.RuntimeOptions) (_result *SyncDataScreenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ScreenId)) {
		body["screenId"] = request.ScreenId
	}

	if !tea.BoolValue(util.IsUnset(request.Ticket)) {
		body["ticket"] = request.Ticket
	}

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
		Action:      tea.String("SyncDataScreen"),
		Version:     tea.String("datacenter_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/datacenter/dataScreens/sync"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncDataScreenResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步数据大屏
//
// @param request - SyncDataScreenRequest
//
// @return SyncDataScreenResponse
func (client *Client) SyncDataScreen(request *SyncDataScreenRequest) (_result *SyncDataScreenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncDataScreenHeaders{}
	_result = &SyncDataScreenResponse{}
	_body, _err := client.SyncDataScreenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
