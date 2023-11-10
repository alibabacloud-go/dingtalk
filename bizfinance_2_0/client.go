// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package bizfinance_2_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	Operator       *string   `json:"operator,omitempty" xml:"operator,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchDeleteReceiptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Code                 *string   `json:"code,omitempty" xml:"code,omitempty"`
	IsDir                *bool     `json:"isDir,omitempty" xml:"isDir,omitempty"`
	Name                 *string   `json:"name,omitempty" xml:"name,omitempty"`
	ParentCode           *string   `json:"parentCode,omitempty" xml:"parentCode,omitempty"`
	Status               *string   `json:"status,omitempty" xml:"status,omitempty"`
	Type                 *string   `json:"type,omitempty" xml:"type,omitempty"`
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

func (s *GetCategoryResponseBody) SetStatus(v string) *GetCategoryResponseBody {
	s.Status = &v
	return s
}

func (s *GetCategoryResponseBody) SetType(v string) *GetCategoryResponseBody {
	s.Type = &v
	return s
}

type GetCategoryResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AccountCode          *string   `json:"accountCode,omitempty" xml:"accountCode,omitempty"`
	AccountId            *string   `json:"accountId,omitempty" xml:"accountId,omitempty"`
	AccountName          *string   `json:"accountName,omitempty" xml:"accountName,omitempty"`
	AccountRemark        *string   `json:"accountRemark,omitempty" xml:"accountRemark,omitempty"`
	AccountType          *string   `json:"accountType,omitempty" xml:"accountType,omitempty"`
	AccountantBookIdList []*string `json:"accountantBookIdList,omitempty" xml:"accountantBookIdList,omitempty" type:"Repeated"`
	Amount               *string   `json:"amount,omitempty" xml:"amount,omitempty"`
	BankCode             *string   `json:"bankCode,omitempty" xml:"bankCode,omitempty"`
	BankName             *string   `json:"bankName,omitempty" xml:"bankName,omitempty"`
	CreateTime           *int64    `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Creator              *string   `json:"creator,omitempty" xml:"creator,omitempty"`
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

type GetFinanceAccountResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFinanceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CreateTime           *int64    `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Creator              *string   `json:"creator,omitempty" xml:"creator,omitempty"`
	Description          *string   `json:"description,omitempty" xml:"description,omitempty"`
	Name                 *string   `json:"name,omitempty" xml:"name,omitempty"`
	ProjectCode          *string   `json:"projectCode,omitempty" xml:"projectCode,omitempty"`
	ProjectName          *string   `json:"projectName,omitempty" xml:"projectName,omitempty"`
	Status               *string   `json:"status,omitempty" xml:"status,omitempty"`
	UserDefineCode       *string   `json:"userDefineCode,omitempty" xml:"userDefineCode,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Code                 *string   `json:"code,omitempty" xml:"code,omitempty"`
	CreateTime           *int64    `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description          *string   `json:"description,omitempty" xml:"description,omitempty"`
	Name                 *string   `json:"name,omitempty" xml:"name,omitempty"`
	Status               *string   `json:"status,omitempty" xml:"status,omitempty"`
	UserDefineCode       *string   `json:"userDefineCode,omitempty" xml:"userDefineCode,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSupplierResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PageNumber *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
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
	HasMore *bool                                  `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List    []*QueryCategoryByPageResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
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
	Code       *string `json:"code,omitempty" xml:"code,omitempty"`
	IsDir      *bool   `json:"isDir,omitempty" xml:"isDir,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	ParentCode *string `json:"parentCode,omitempty" xml:"parentCode,omitempty"`
	Status     *string `json:"status,omitempty" xml:"status,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCategoryByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
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
	HasMore *bool                                  `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List    []*QueryCustomerByPageResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
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
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	CreateTime     *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCustomerByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
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
	HasMore *bool                                           `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List    []*QueryEnterpriseAccountByPageResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
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
	AccountCode   *string `json:"accountCode,omitempty" xml:"accountCode,omitempty"`
	AccountId     *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	AccountName   *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	AccountRemark *string `json:"accountRemark,omitempty" xml:"accountRemark,omitempty"`
	AccountType   *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	Amount        *string `json:"amount,omitempty" xml:"amount,omitempty"`
	BankCode      *string `json:"bankCode,omitempty" xml:"bankCode,omitempty"`
	BankName      *string `json:"bankName,omitempty" xml:"bankName,omitempty"`
	CreateTime    *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Creator       *string `json:"creator,omitempty" xml:"creator,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryEnterpriseAccountByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Result *QueryInstancePaymentOrderDetailResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryInstancePaymentOrderDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInstancePaymentOrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInstancePaymentOrderDetailResponseBody) SetResult(v *QueryInstancePaymentOrderDetailResponseBodyResult) *QueryInstancePaymentOrderDetailResponseBody {
	s.Result = v
	return s
}

type QueryInstancePaymentOrderDetailResponseBodyResult struct {
	Amount          *string                                                           `json:"amount,omitempty" xml:"amount,omitempty"`
	InstanceId      *string                                                           `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	PayeeAccountDTO *QueryInstancePaymentOrderDetailResponseBodyResultPayeeAccountDTO `json:"payeeAccountDTO,omitempty" xml:"payeeAccountDTO,omitempty" type:"Struct"`
	PayerAccountDTO *QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTO `json:"payerAccountDTO,omitempty" xml:"payerAccountDTO,omitempty" type:"Struct"`
	Remark          *string                                                           `json:"remark,omitempty" xml:"remark,omitempty"`
	Usage           *string                                                           `json:"usage,omitempty" xml:"usage,omitempty"`
	UserId          *string                                                           `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryInstancePaymentOrderDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryInstancePaymentOrderDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryInstancePaymentOrderDetailResponseBodyResult) SetAmount(v string) *QueryInstancePaymentOrderDetailResponseBodyResult {
	s.Amount = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyResult) SetInstanceId(v string) *QueryInstancePaymentOrderDetailResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyResult) SetPayeeAccountDTO(v *QueryInstancePaymentOrderDetailResponseBodyResultPayeeAccountDTO) *QueryInstancePaymentOrderDetailResponseBodyResult {
	s.PayeeAccountDTO = v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyResult) SetPayerAccountDTO(v *QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTO) *QueryInstancePaymentOrderDetailResponseBodyResult {
	s.PayerAccountDTO = v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyResult) SetRemark(v string) *QueryInstancePaymentOrderDetailResponseBodyResult {
	s.Remark = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyResult) SetUsage(v string) *QueryInstancePaymentOrderDetailResponseBodyResult {
	s.Usage = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyResult) SetUserId(v string) *QueryInstancePaymentOrderDetailResponseBodyResult {
	s.UserId = &v
	return s
}

type QueryInstancePaymentOrderDetailResponseBodyResultPayeeAccountDTO struct {
	BankOpenDTO *QueryInstancePaymentOrderDetailResponseBodyResultPayeeAccountDTOBankOpenDTO `json:"bankOpenDTO,omitempty" xml:"bankOpenDTO,omitempty" type:"Struct"`
}

func (s QueryInstancePaymentOrderDetailResponseBodyResultPayeeAccountDTO) String() string {
	return tea.Prettify(s)
}

func (s QueryInstancePaymentOrderDetailResponseBodyResultPayeeAccountDTO) GoString() string {
	return s.String()
}

func (s *QueryInstancePaymentOrderDetailResponseBodyResultPayeeAccountDTO) SetBankOpenDTO(v *QueryInstancePaymentOrderDetailResponseBodyResultPayeeAccountDTOBankOpenDTO) *QueryInstancePaymentOrderDetailResponseBodyResultPayeeAccountDTO {
	s.BankOpenDTO = v
	return s
}

type QueryInstancePaymentOrderDetailResponseBodyResultPayeeAccountDTOBankOpenDTO struct {
	BankBranchCode *string `json:"bankBranchCode,omitempty" xml:"bankBranchCode,omitempty"`
	BankBranchName *string `json:"bankBranchName,omitempty" xml:"bankBranchName,omitempty"`
	BankCardNo     *string `json:"bankCardNo,omitempty" xml:"bankCardNo,omitempty"`
	BankCode       *string `json:"bankCode,omitempty" xml:"bankCode,omitempty"`
	BankName       *string `json:"bankName,omitempty" xml:"bankName,omitempty"`
	Type           *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryInstancePaymentOrderDetailResponseBodyResultPayeeAccountDTOBankOpenDTO) String() string {
	return tea.Prettify(s)
}

func (s QueryInstancePaymentOrderDetailResponseBodyResultPayeeAccountDTOBankOpenDTO) GoString() string {
	return s.String()
}

func (s *QueryInstancePaymentOrderDetailResponseBodyResultPayeeAccountDTOBankOpenDTO) SetBankBranchCode(v string) *QueryInstancePaymentOrderDetailResponseBodyResultPayeeAccountDTOBankOpenDTO {
	s.BankBranchCode = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyResultPayeeAccountDTOBankOpenDTO) SetBankBranchName(v string) *QueryInstancePaymentOrderDetailResponseBodyResultPayeeAccountDTOBankOpenDTO {
	s.BankBranchName = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyResultPayeeAccountDTOBankOpenDTO) SetBankCardNo(v string) *QueryInstancePaymentOrderDetailResponseBodyResultPayeeAccountDTOBankOpenDTO {
	s.BankCardNo = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyResultPayeeAccountDTOBankOpenDTO) SetBankCode(v string) *QueryInstancePaymentOrderDetailResponseBodyResultPayeeAccountDTOBankOpenDTO {
	s.BankCode = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyResultPayeeAccountDTOBankOpenDTO) SetBankName(v string) *QueryInstancePaymentOrderDetailResponseBodyResultPayeeAccountDTOBankOpenDTO {
	s.BankName = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyResultPayeeAccountDTOBankOpenDTO) SetType(v string) *QueryInstancePaymentOrderDetailResponseBodyResultPayeeAccountDTOBankOpenDTO {
	s.Type = &v
	return s
}

type QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTO struct {
	BankOpenDTO           *QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTOBankOpenDTO `json:"bankOpenDTO,omitempty" xml:"bankOpenDTO,omitempty" type:"Struct"`
	EnterpriseAccountCode *string                                                                      `json:"enterpriseAccountCode,omitempty" xml:"enterpriseAccountCode,omitempty"`
}

func (s QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTO) String() string {
	return tea.Prettify(s)
}

func (s QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTO) GoString() string {
	return s.String()
}

func (s *QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTO) SetBankOpenDTO(v *QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTOBankOpenDTO) *QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTO {
	s.BankOpenDTO = v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTO) SetEnterpriseAccountCode(v string) *QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTO {
	s.EnterpriseAccountCode = &v
	return s
}

type QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTOBankOpenDTO struct {
	BankBranchCode *string `json:"bankBranchCode,omitempty" xml:"bankBranchCode,omitempty"`
	BankBranchName *string `json:"bankBranchName,omitempty" xml:"bankBranchName,omitempty"`
	BankCardNo     *string `json:"bankCardNo,omitempty" xml:"bankCardNo,omitempty"`
	BankCode       *string `json:"bankCode,omitempty" xml:"bankCode,omitempty"`
	BankName       *string `json:"bankName,omitempty" xml:"bankName,omitempty"`
	Type           *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTOBankOpenDTO) String() string {
	return tea.Prettify(s)
}

func (s QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTOBankOpenDTO) GoString() string {
	return s.String()
}

func (s *QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTOBankOpenDTO) SetBankBranchCode(v string) *QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTOBankOpenDTO {
	s.BankBranchCode = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTOBankOpenDTO) SetBankBranchName(v string) *QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTOBankOpenDTO {
	s.BankBranchName = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTOBankOpenDTO) SetBankCardNo(v string) *QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTOBankOpenDTO {
	s.BankCardNo = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTOBankOpenDTO) SetBankCode(v string) *QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTOBankOpenDTO {
	s.BankCode = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTOBankOpenDTO) SetBankName(v string) *QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTOBankOpenDTO {
	s.BankName = &v
	return s
}

func (s *QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTOBankOpenDTO) SetType(v string) *QueryInstancePaymentOrderDetailResponseBodyResultPayerAccountDTOBankOpenDTO {
	s.Type = &v
	return s
}

type QueryInstancePaymentOrderDetailResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryInstancePaymentOrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
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
	HasMore *bool                                 `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List    []*QueryProjectByPageResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
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
	Caode          *string `json:"caode,omitempty" xml:"caode,omitempty"`
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	CreateTime     *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Creator        *string `json:"creator,omitempty" xml:"creator,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	ProjectCode    *string `json:"projectCode,omitempty" xml:"projectCode,omitempty"`
	ProjectName    *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryProjectByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
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
	HasMore *bool                                  `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List    []*QuerySupplierByPageResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
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
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	CreateTime     *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QuerySupplierByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	BankCardNo      *string `json:"bankCardNo,omitempty" xml:"bankCardNo,omitempty"`
	Operator        *string `json:"operator,omitempty" xml:"operator,omitempty"`
	SignOperateType *string `json:"signOperateType,omitempty" xml:"signOperateType,omitempty"`
}

func (s SignEnterpriseAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s SignEnterpriseAccountRequest) GoString() string {
	return s.String()
}

func (s *SignEnterpriseAccountRequest) SetBankCardNo(v string) *SignEnterpriseAccountRequest {
	s.BankCardNo = &v
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SignEnterpriseAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	FailReason *string                                  `json:"failReason,omitempty" xml:"failReason,omitempty"`
	OrderNo    *string                                  `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	OutOrderNo *string                                  `json:"outOrderNo,omitempty" xml:"outOrderNo,omitempty"`
	PayerBank  *UpdateInstanceOrderInfoRequestPayerBank `json:"payerBank,omitempty" xml:"payerBank,omitempty" type:"Struct"`
	Status     *string                                  `json:"status,omitempty" xml:"status,omitempty"`
	UserId     *string                                  `json:"userId,omitempty" xml:"userId,omitempty"`
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

func (s *UpdateInstanceOrderInfoRequest) SetStatus(v string) *UpdateInstanceOrderInfoRequest {
	s.Status = &v
	return s
}

func (s *UpdateInstanceOrderInfoRequest) SetUserId(v string) *UpdateInstanceOrderInfoRequest {
	s.UserId = &v
	return s
}

type UpdateInstanceOrderInfoRequestPayerBank struct {
	CardNo *string `json:"cardNo,omitempty" xml:"cardNo,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
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
	FailReason      *string `json:"failReason,omitempty" xml:"failReason,omitempty"`
	OrderNo         *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	OutOrderNo      *string `json:"outOrderNo,omitempty" xml:"outOrderNo,omitempty"`
	PayerBankShrink *string `json:"payerBank,omitempty" xml:"payerBank,omitempty"`
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
	UserId          *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateInstanceOrderInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (client *Client) SignEnterpriseAccountWithOptions(request *SignEnterpriseAccountRequest, headers *SignEnterpriseAccountHeaders, runtime *util.RuntimeOptions) (_result *SignEnterpriseAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BankCardNo)) {
		query["bankCardNo"] = request.BankCardNo
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
