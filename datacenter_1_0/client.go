// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package datacenter_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	// 页数,第几页
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页条数
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 关键词
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
	// 返回结果
	// DEPARTMENT:列入决定机关
	// IN_REASON 列入原因
	// OUT_DATE:移出日期
	// OUT_DEPARTMENT:移出决定机关
	// OUT_REASON:移出原因
	// IN_DATE:列入日期
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// 总条数
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAbnormalOperationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetAbnormalOperationResponse) SetBody(v *GetAbnormalOperationResponseBody) *GetAbnormalOperationResponse {
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
	// 页数,第几页
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页条数
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 关键词
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
	// 返回结果
	// DATE_COL:处罚日期
	// NUMBER:决定书文号
	// ILLEGAL_TYPE:处罚类型
	// DEPARTMENT:处罚机关
	// PUBLIC_DATE 公示日期
	// CONTENT:处罚内容
	// BASED_ON:处罚依据
	// DESCRIPTION:处罚判决书
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// 总条数
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAdministrativePenaltiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 页数,第几页
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页条数
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 关键词
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
	// 返回结果
	// ENT_NAME:企业名称
	// LEGAL_NAME:法定代表人姓名
	// ES_DATE:开业日期
	// ENT_STATUS:经营状态
	// REG_CAP:注册资本
	// REC_CAP:实收资本
	// SOCIAL_CREDIT_CODE:统一社会信用代码
	// LICENSE_NUMBER:工商注册号
	// ORG_NO:组织机构代码
	// TAX_NUM:纳税人识别号
	// ENT_TYPE:企业类型
	// INDUSTRY_NAME_LV1:国民经济行业门类名称
	// INDUSTRY_NAME_LV2:国民经济行业大类名称
	// OP_FROM:经营期限自
	// OP_TO:经营期限至
	// COLLEGUES_NUM:人员规模
	// INSURED_NUM:参保人数
	// ENT_NAME_ENG:英文名称
	// FORMER_NAMES:曾用名
	// REG_ORG:登记机关
	// CHECK_DATE:核准日期
	// OP_SCOPE:经营范围
	// IDENTITY_ID:ID
	// ENT_ADDRESS:企业地址
	// EMPLOYEES_INFO:主要管理人员
	// ENT_BRIEF:公司简介
	// REG_ORG_PROVINCE:注册地址所在省
	// REG_ORG_CITY:注册地址所在城市
	// REG_ORG_DISTRICT:注册地址所在区县
	// STD_REG_CAP:清洗后注册资本
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// 总条数
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetBasicInfoResponse) SetBody(v *GetBasicInfoResponseBody) *GetBasicInfoResponse {
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
	// 页数,第几页
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页条数
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 关键词
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
	// 返回结果
	// DEPARTMENT:处罚单位
	// ENT_NAME:企业名称
	// EXEC_STATUS 执行情况
	// PUNISH_BASIS:处罚依据
	// PUNISH_CONTENT:处罚事由
	// PUNISH_LAW:违反法律
	// PUNISH_NUM:决定文书号
	// PUNISH_RES:处罚结果
	// PUNISH_DATE:处罚时间
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// 总条数
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEnvironmentalPenaltiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetEnvironmentalPenaltiesResponse) SetBody(v *GetEnvironmentalPenaltiesResponseBody) *GetEnvironmentalPenaltiesResponse {
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
	// 页数,第几页
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页条数
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 关键词
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
	// 返回结果
	// STOCK_TYPE:股东类型
	// STOCK_NAME:股东名称
	// STOCK_PERCENT:持股比例
	// SHOULD_CAPI:认缴出资额
	// SHOULD_CAPI_TIME:认缴出资日期
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// 总条数
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHolderInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetHolderInfoResponse) SetBody(v *GetHolderInfoResponseBody) *GetHolderInfoResponse {
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
	// 页数,第几页
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页条数
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 关键词
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
	// 返回结果
	// DEPARTMENT:主管机关
	// END_DATE:有效日期止
	// ENT_NAME:纳税人名称
	// QUALIFICATION 纳税人资格
	// START_DATE:有效日期起
	// TAXPAYER_NUM:纳税人识别号
	// JUDGE_DATE:认定时间
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// 总条数
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetQeneralTaxpayerInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetQeneralTaxpayerInfoResponse) SetBody(v *GetQeneralTaxpayerInfoResponseBody) *GetQeneralTaxpayerInfoResponse {
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
	// 页数,第几页
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页条数
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 关键词
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
	// 返回结果
	// IN_DATE:列入日期
	// IN_DEPARTMENT:列入决定机关
	// IN_REASON:列入严重违法失信企业名单原因
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// 总条数
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSeriousViolationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetSeriousViolationResponse) SetBody(v *GetSeriousViolationResponseBody) *GetSeriousViolationResponse {
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PostCorpAuthInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryActiveUserStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAnhmdStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryApprovalStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAttendanceStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryBlackboardStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCalendarStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryCalendarStatisticalDataResponse) SetBody(v *QueryCalendarStatisticalDataResponseBody) *QueryCalendarStatisticalDataResponse {
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCheckinStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCircleStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Keyword    *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	PageNumber *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
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
	// code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data []map[string]*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// traceId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// total
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCompanyBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CorpIds   []*string `json:"corpIds,omitempty" xml:"corpIds,omitempty" type:"Repeated"`
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
	// arguments
	Arguments []*string `json:"arguments,omitempty" xml:"arguments,omitempty" type:"Repeated"`
	// result
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
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
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDigitalDistrictOrgInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDingReciveStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDingSendStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDocumentStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryDocumentStatisticalDataResponse) SetBody(v *QueryDocumentStatisticalDataResponseBody) *QueryDocumentStatisticalDataResponse {
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDriveStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryDriveStatisticalDataResponse) SetBody(v *QueryDriveStatisticalDataResponseBody) *QueryDriveStatisticalDataResponse {
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryEmployeeTypeStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 部门ID
	DeptId *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 结束日期
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// 分页页码
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页大小
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 服务编码
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// statDate
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// 员工ID
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
	MetaList []*QueryGeneralDataServiceResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
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

type QueryGeneralDataServiceResponseBodyMetaList struct {
	// 指标名称
	FieldDesc *string `json:"fieldDesc,omitempty" xml:"fieldDesc,omitempty"`
	// 指标口径
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// 指标ID
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// 指标单位
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryGeneralDataServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryGeneralDataServiceResponse) SetBody(v *QueryGeneralDataServiceResponseBody) *QueryGeneralDataServiceResponse {
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryGroupLiveStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryGroupMessageStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryHealthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMailStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryOfficialDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 数据集id
	DsId *string `json:"dsId,omitempty" xml:"dsId,omitempty"`
	// 用户id
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryOfficialDatasetFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 关键词搜索
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 起始页，从1开始
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 单页大小，最大100
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryOfficialDatasetListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryOfficialDatasetListResponse) SetBody(v *QueryOfficialDatasetListResponseBody) *QueryOfficialDatasetListResponse {
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryOnlineUserStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRedEnvelopeReciveStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRedEnvelopeSendStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryReportStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryReportStatisticalDataResponse) SetBody(v *QueryReportStatisticalDataResponseBody) *QueryReportStatisticalDataResponse {
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySingleMessageStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTelMeetingStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTodoStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryTodoStatisticalDataResponse) SetBody(v *QueryTodoStatisticalDataResponseBody) *QueryTodoStatisticalDataResponse {
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryVedioMeetingStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydActiveDayStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydActiveMonthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydActiveWeekStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydAppDayStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydAppMonthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydAppStdStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydAppWeekStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydCalendarDayStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydCalendarMonthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydCalendarWeekStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydDingMsgDayStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydDingMsgMonthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydDingMsgWeekStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydGroupMsgDayStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydGroupMsgMonthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydGroupMsgWeekStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydLogDayStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydLogMonthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydLogWeekStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydMeetingDayStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydMeetingMonthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydMeetingWeekStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydNoticeDayStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydNoticeMonthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydNoticeWeekStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydSingleMsgDayStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydSingleMsgMonthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydSingleMsgWeekStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydToatlMsgDayStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydToatlMsgMonthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydToatlMsgWeekStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydTodoDayStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydTodoMonthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydTodoWeekStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydTotalDayStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydTotalMonthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydTotalStdStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// statDate
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
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
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
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// 指标单位
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
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryYydTotalWeekStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 起始页
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 页面大小
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 关键词
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
	// 返回数据结果
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// 总条数
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchCompanyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SearchCompanyResponse) SetBody(v *SearchCompanyResponseBody) *SearchCompanyResponse {
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
	_result = &GetAbnormalOperationResponse{}
	_body, _err := client.DoROARequest(tea.String("GetAbnormalOperation"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/companies/abnormalOperations"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetAdministrativePenaltiesResponse{}
	_body, _err := client.DoROARequest(tea.String("GetAdministrativePenalties"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/companies/administrativePenalties"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetBasicInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetBasicInfo"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/companies/businessBasicInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetEnvironmentalPenaltiesResponse{}
	_body, _err := client.DoROARequest(tea.String("GetEnvironmentalPenalties"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/companies/environmentalPenalties"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetHolderInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetHolderInfo"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/companies/shareholderInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetQeneralTaxpayerInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetQeneralTaxpayerInfo"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/companies/generalTaxpayerInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetSeriousViolationResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSeriousViolation"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/companies/seriousViolations"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &PostCorpAuthInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("PostCorpAuthInfo"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/datacenter/corporations/authorize"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryActiveUserStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryActiveUserStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/activeUserData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryAnhmdStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAnhmdStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/statisticDatas/anHmd"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryApprovalStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryApprovalStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/approvalData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryAttendanceStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAttendanceStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/attendanceData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryBlackboardStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryBlackboardStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/blackboardData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryCalendarStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryCalendarStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/calendarData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryCheckinStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryCheckinStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/checkinData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryCircleStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryCircleStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/circleData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryCompanyBasicInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryCompanyBasicInfo"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/companies/basicInfo"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryDigitalDistrictOrgInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryDigitalDistrictOrgInfo"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/datacenter/digitalCounty/orgInfos/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryDingReciveStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryDingReciveStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/dingReciveData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryDingSendStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryDingSendStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/dingSendData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryDocumentStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryDocumentStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/documentData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryDriveStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryDriveStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/driveData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryEmployeeTypeStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryEmployeeTypeStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/employeeTypeData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryGeneralDataServiceResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryGeneralDataService"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/generalDataServices"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryGroupLiveStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryGroupLiveStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/groupLiveData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryGroupMessageStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryGroupMessageStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/groupMessageData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryHealthStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryHealthStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/healtheUserData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryMailStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryMailStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/mailData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryOfficialDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryOfficialData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/datas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryOfficialDatasetFieldsResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryOfficialDatasetFields"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/datasetFields"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryOfficialDatasetListResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryOfficialDatasetList"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/datasetLists"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryOnlineUserStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryOnlineUserStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/onlineUserData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryRedEnvelopeReciveStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryRedEnvelopeReciveStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/redEnvelopeReciveData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryRedEnvelopeSendStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryRedEnvelopeSendStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/redEnvelopeSendData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryReportStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryReportStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/reportData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QuerySingleMessageStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QuerySingleMessageStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/singleMessagerData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryTelMeetingStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryTelMeetingStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/telMeetingData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryTodoStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryTodoStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/todoUserData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryVedioMeetingStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryVedioMeetingStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/vedioMeetingData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydActiveDayStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydActiveDayStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydActiveDayDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydActiveMonthStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydActiveMonthStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydActiveMonthDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydActiveWeekStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydActiveWeekStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydActiveWeekDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydAppDayStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydAppDayStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydAppDayDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydAppMonthStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydAppMonthStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydAppMonthDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydAppStdStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydAppStdStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydAppStdDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydAppWeekStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydAppWeekStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydAppWeekDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydCalendarDayStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydCalendarDayStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydCalendarDayDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydCalendarMonthStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydCalendarMonthStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydCalendarMonthDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydCalendarWeekStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydCalendarWeekStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydCalendarWeekDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydDingMsgDayStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydDingMsgDayStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydDingMsgDayDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydDingMsgMonthStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydDingMsgMonthStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydDingMsgMonthDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydDingMsgWeekStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydDingMsgWeekStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydDingMsgWeekDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydGroupMsgDayStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydGroupMsgDayStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydGroupMsgDayDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydGroupMsgMonthStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydGroupMsgMonthStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydGroupMsgMonthDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydGroupMsgWeekStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydGroupMsgWeekStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydGroupMsgWeekDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydLogDayStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydLogDayStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydLogDayDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydLogMonthStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydLogMonthStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydLogMonthDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydLogWeekStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydLogWeekStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydLogWeekDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydMeetingDayStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydMeetingDayStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydMeetingDayDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydMeetingMonthStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydMeetingMonthStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydMeetingMonthDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydMeetingWeekStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydMeetingWeekStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydMeetingWeekDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydNoticeDayStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydNoticeDayStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydNoticeDayDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydNoticeMonthStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydNoticeMonthStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydNoticeMonthDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydNoticeWeekStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydNoticeWeekStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydNoticeWeekDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydSingleMsgDayStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydSingleMsgDayStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydSingleMsgDayDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydSingleMsgMonthStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydSingleMsgMonthStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydSingleMsgMonthDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydSingleMsgWeekStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydSingleMsgWeekStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydSingleMsgWeekDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydToatlMsgDayStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydToatlMsgDayStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydToatlMsgDayDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydToatlMsgMonthStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydToatlMsgMonthStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydToatlMsgMonthDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydToatlMsgWeekStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydToatlMsgWeekStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydToatlMsgWeekDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydTodoDayStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydTodoDayStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydTodoDayDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydTodoMonthStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydTodoMonthStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydTodoMonthDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydTodoWeekStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydTodoWeekStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydTodoWeekDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydTotalDayStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydTotalDayStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydTotalDayDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydTotalMonthStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydTotalMonthStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydTotalMonthDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydTotalStdStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydTotalStdStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydTotalStdDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &QueryYydTotalWeekStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryYydTotalWeekStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/yydTotalWeekDatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &SearchCompanyResponse{}
	_body, _err := client.DoROARequest(tea.String("SearchCompany"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/keywords/companies"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
